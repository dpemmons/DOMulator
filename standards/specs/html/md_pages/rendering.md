HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 14 The XML syntax](xhtml.html) --- [Table of Contents](./) --- [16
Obsolete features →](obsolete.html)

1.  [[[15]{.secno} Rendering](rendering.html#rendering)]{#toc-rendering}
    1.  [[15.1]{.secno} Introduction](rendering.html#introduction-16)
    2.  [[15.2]{.secno} The CSS user agent style sheet and
        presentational
        hints](rendering.html#the-css-user-agent-style-sheet-and-presentational-hints)
    3.  [[15.3]{.secno} Non-replaced
        elements](rendering.html#non-replaced-elements)
        1.  [[15.3.1]{.secno} Hidden
            elements](rendering.html#hidden-elements)
        2.  [[15.3.2]{.secno} The page](rendering.html#the-page)
        3.  [[15.3.3]{.secno} Flow
            content](rendering.html#flow-content-3)
        4.  [[15.3.4]{.secno} Phrasing
            content](rendering.html#phrasing-content-3)
        5.  [[15.3.5]{.secno} Bidirectional
            text](rendering.html#bidi-rendering)
        6.  [[15.3.6]{.secno} Sections and
            headings](rendering.html#sections-and-headings)
        7.  [[15.3.7]{.secno} Lists](rendering.html#lists)
        8.  [[15.3.8]{.secno} Tables](rendering.html#tables-2)
        9.  [[15.3.9]{.secno} Margin collapsing
            quirks](rendering.html#margin-collapsing-quirks)
        10. [[15.3.10]{.secno} Form
            controls](rendering.html#form-controls)
        11. [[15.3.11]{.secno} The `hr`
            element](rendering.html#the-hr-element-2)
        12. [[15.3.12]{.secno} The `fieldset` and `legend`
            elements](rendering.html#the-fieldset-and-legend-elements)
    4.  [[15.4]{.secno} Replaced
        elements](rendering.html#replaced-elements)
        1.  [[15.4.1]{.secno} Embedded
            content](rendering.html#embedded-content-rendering-rules)
        2.  [[15.4.2]{.secno} Images](rendering.html#images-3)
        3.  [[15.4.3]{.secno} Attributes for embedded content and
            images](rendering.html#attributes-for-embedded-content-and-images)
        4.  [[15.4.4]{.secno} Image maps](rendering.html#image-maps-2)
    5.  [[15.5]{.secno} Widgets](rendering.html#widgets)
        1.  [[15.5.1]{.secno} Native
            appearance](rendering.html#native-appearance-2)
        2.  [[15.5.2]{.secno} Writing mode](rendering.html#writing-mode)
        3.  [[15.5.3]{.secno} Button
            layout](rendering.html#button-layout)
        4.  [[15.5.4]{.secno} The `button`
            element](rendering.html#the-button-element-2)
        5.  [[15.5.5]{.secno} The `details` and `summary`
            elements](rendering.html#the-details-and-summary-elements)
        6.  [[15.5.6]{.secno} The `input` element as a text entry
            widget](rendering.html#the-input-element-as-a-text-entry-widget)
        7.  [[15.5.7]{.secno} The `input` element as domain-specific
            widgets](rendering.html#the-input-element-as-domain-specific-widgets)
        8.  [[15.5.8]{.secno} The `input` element as a range
            control](rendering.html#the-input-element-as-a-range-control)
        9.  [[15.5.9]{.secno} The `input` element as a color
            well](rendering.html#the-input-element-as-a-colour-well)
        10. [[15.5.10]{.secno} The `input` element as a checkbox and
            radio button
            widgets](rendering.html#the-input-element-as-a-checkbox-and-radio-button-widgets)
        11. [[15.5.11]{.secno} The `input` element as a file upload
            control](rendering.html#the-input-element-as-a-file-upload-control)
        12. [[15.5.12]{.secno} The `input` element as a
            button](rendering.html#the-input-element-as-a-button)
        13. [[15.5.13]{.secno} The `marquee`
            element](rendering.html#the-marquee-element-2)
        14. [[15.5.14]{.secno} The `meter`
            element](rendering.html#the-meter-element-2)
        15. [[15.5.15]{.secno} The `progress`
            element](rendering.html#the-progress-element-2)
        16. [[15.5.16]{.secno} The `select`
            element](rendering.html#the-select-element-2)
        17. [[15.5.17]{.secno} The `textarea`
            element](rendering.html#the-textarea-element-2)
    6.  [[15.6]{.secno} Frames and
        framesets](rendering.html#frames-and-framesets)
    7.  [[15.7]{.secno} Interactive
        media](rendering.html#interactive-media)
        1.  [[15.7.1]{.secno} Links, forms, and
            navigation](rendering.html#links,-forms,-and-navigation)
        2.  [[15.7.2]{.secno} The `title`
            attribute](rendering.html#the-title-attribute-2)
        3.  [[15.7.3]{.secno} Editing
            hosts](rendering.html#editing-hosts)
        4.  [[15.7.4]{.secno} Text rendered in native user
            interfaces](rendering.html#text-rendered-in-native-user-interfaces)
    8.  [[15.8]{.secno} Print media](rendering.html#print-media)
    9.  [[15.9]{.secno} Unstyled XML
        documents](rendering.html#unstyled-xml-documents)

## [15]{.secno} Rendering[](#rendering){.self-link}

*User agents are not required to present HTML documents in any
particular way. However, this section provides a set of suggestions for
rendering HTML documents that, if followed, are likely to lead to a user
experience that closely resembles the experience intended by the
documents\' authors. So as to avoid confusion regarding the normativity
of this section, \"must\" has not been used. Instead, the term
\"expected\" is used to indicate behavior that will lead to this
experience. For the purposes of conformance for user agents designated
as [supporting the suggested default
rendering](infrastructure.html#renderingUA), the term \"expected\" in
this section has the same conformance implications as \"must\".*

### [15.1]{.secno} Introduction[](#introduction-16){.self-link} {#introduction-16}

The suggestions in this section are generally expressed in CSS terms.
User agents are expected to either support CSS, or translate from the
CSS rules given in this section to approximations for other presentation
mechanisms.

In the absence of style-layer rules to the contrary (e.g. author style
sheets), user agents are expected to render an element so that it
conveys to the user the meaning that the element
[represents](dom.html#represents){#introduction-16:represents}, as
described by this specification.

The suggestions in this section generally assume a visual output medium
with a resolution of 96dpi or greater, but HTML is intended to apply to
multiple media (it is a *media-independent* language). User agent
implementers are encouraged to adapt the suggestions in this section to
their target media.

------------------------------------------------------------------------

An element is [being rendered]{#being-rendered .dfn export=""} if it has
any associated CSS layout boxes, SVG layout boxes, or some equivalent in
other styling languages.

Just being off-screen does not mean the element is not [being
rendered](#being-rendered){#introduction-16:being-rendered}. The
presence of the
[`hidden`{#introduction-16:attr-hidden}](interaction.html#attr-hidden)
attribute normally means the element is not [being
rendered](#being-rendered){#introduction-16:being-rendered-2}, though
this might be overridden by the style sheets.

The [fully
active](document-sequences.html#fully-active){#introduction-16:fully-active}
state does not affect whether an element is [being
rendered](#being-rendered){#introduction-16:being-rendered-3} or not.
Even if a document is not [fully
active](document-sequences.html#fully-active){#introduction-16:fully-active-2}
and not shown at all to the user, elements within it can still qualify
as \"being rendered\".

An element is said to [intersect the viewport]{#intersect-the-viewport
.dfn} when it is [being
rendered](#being-rendered){#introduction-16:being-rendered-4} and its
associated CSS layout box intersects the
[viewport](https://drafts.csswg.org/css2/#viewport){#introduction-16:viewport
x-internal="viewport"}.

Similar to the [being
rendered](#being-rendered){#introduction-16:being-rendered-5} state,
elements in non-[fully
active](document-sequences.html#fully-active){#introduction-16:fully-active-3}
documents can still [intersect the
viewport](#intersect-the-viewport){#introduction-16:intersect-the-viewport}.
The
[viewport](https://drafts.csswg.org/css2/#viewport){#introduction-16:viewport-2
x-internal="viewport"} is not shared between documents and might not
always be shown to the user, so an element in a non-[fully
active](document-sequences.html#fully-active){#introduction-16:fully-active-4}
document can still intersect the
[viewport](https://drafts.csswg.org/css2/#viewport){#introduction-16:viewport-3
x-internal="viewport"} associated with its document.

This specification does not define the precise timing for when the
intersection is tested, but it is suggested that the timing match that
of the Intersection Observer API.
[\[INTERSECTIONOBSERVER\]](references.html#refsINTERSECTIONOBSERVER)

An element is [delegating its rendering to its
children]{#delegating-its-rendering-to-its-children .dfn} if it is not
[being rendered](#being-rendered){#introduction-16:being-rendered-6} but
its children (if any) could [be
rendered](#being-rendered){#introduction-16:being-rendered-7}, as a
result of CSS \'display: contents\', or some equivalent in other styling
languages. [\[CSSDISPLAY\]](references.html#refsCSSDISPLAY)

------------------------------------------------------------------------

User agents that do not honor author-level CSS style sheets are
nonetheless expected to act as if they applied the CSS rules given in
these sections in a manner consistent with this specification and the
relevant CSS and Unicode specifications.
[\[CSS\]](references.html#refsCSS)
[\[UNICODE\]](references.html#refsUNICODE)
[\[BIDI\]](references.html#refsBIDI)

This is especially important for issues relating to the
[\'display\'](https://drafts.csswg.org/css2/#display-prop){#introduction-16:'display'
x-internal="'display'"},
[\'unicode-bidi\'](https://drafts.csswg.org/css-writing-modes/#unicode-bidi){#introduction-16:'unicode-bidi'
x-internal="'unicode-bidi'"}, and
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#introduction-16:'direction'
x-internal="'direction'"} properties.

### [15.2]{.secno} The CSS user agent style sheet and presentational hints[](#the-css-user-agent-style-sheet-and-presentational-hints){.self-link}

The CSS rules given in these subsections are, except where otherwise
specified, expected to be used as part of the user-agent level style
sheet defaults for all documents that contain [HTML
elements](infrastructure.html#html-elements){#the-css-user-agent-style-sheet-and-presentational-hints:html-elements}.

Some rules are intended for the author-level zero-specificity
presentational hints part of the CSS cascade; these are explicitly
called out as [presentational hints]{#presentational-hints .dfn
lt="presentational hint" export=""}.

------------------------------------------------------------------------

When the text below says that an attribute `attribute`{.variable} on an
element `element`{.variable} [maps to the pixel length
property]{#maps-to-the-pixel-length-property .dfn} (or properties)
`properties`{.variable}, it means that if `element`{.variable} has an
attribute `attribute`{.variable} set, and parsing that attribute\'s
value using the [rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-css-user-agent-style-sheet-and-presentational-hints:rules-for-parsing-non-negative-integers}
doesn\'t generate an error, then the user agent is expected to use the
parsed value as a pixel length for a [presentational
hint](#presentational-hints){#the-css-user-agent-style-sheet-and-presentational-hints:presentational-hints}
for `properties`{.variable}.

When the text below says that an attribute `attribute`{.variable} on an
element `element`{.variable} [maps to the dimension
property]{#maps-to-the-dimension-property .dfn} (or properties)
`properties`{.variable}, it means that if `element`{.variable} has an
attribute `attribute`{.variable} set, and parsing that attribute\'s
value using the [rules for parsing dimension
values](common-microsyntaxes.html#rules-for-parsing-dimension-values){#the-css-user-agent-style-sheet-and-presentational-hints:rules-for-parsing-dimension-values}
doesn\'t generate an error, then the user agent is expected to use the
parsed dimension as the value for a [presentational
hint](#presentational-hints){#the-css-user-agent-style-sheet-and-presentational-hints:presentational-hints-2}
for `properties`{.variable}, with the value given as a pixel length if
the dimension was a length, and with the value given as a percentage if
the dimension was a percentage.

When the text below says that an attribute `attribute`{.variable} on an
element `element`{.variable} [maps to the dimension property (ignoring
zero)]{#maps-to-the-dimension-property-(ignoring-zero) .dfn} (or
properties) `properties`{.variable}, it means that if
`element`{.variable} has an attribute `attribute`{.variable} set, and
parsing that attribute\'s value using the [rules for parsing nonzero
dimension
values](common-microsyntaxes.html#rules-for-parsing-non-zero-dimension-values){#the-css-user-agent-style-sheet-and-presentational-hints:rules-for-parsing-non-zero-dimension-values}
doesn\'t generate an error, then the user agent is expected to use the
parsed dimension as the value for a [presentational
hint](#presentational-hints){#the-css-user-agent-style-sheet-and-presentational-hints:presentational-hints-3}
for `properties`{.variable}, with the value given as a pixel length if
the dimension was a length, and with the value given as a percentage if
the dimension was a percentage.

When the text below says that a pair of attributes `w`{.variable} and
`h`{.variable} on an element `element`{.variable} [map to the
aspect-ratio property]{#map-to-the-aspect-ratio-property .dfn}, it means
that if `element`{.variable} has both attributes `w`{.variable} and
`h`{.variable}, and parsing those attributes\' values using the [rules
for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-css-user-agent-style-sheet-and-presentational-hints:rules-for-parsing-non-negative-integers-2}
doesn\'t generate an error for either, then the user agent is expected
to use the parsed integers as a [presentational
hint](#presentational-hints){#the-css-user-agent-style-sheet-and-presentational-hints:presentational-hints-4}
for the
[\'aspect-ratio\'](https://drafts.csswg.org/css-sizing-4/#aspect-ratio){#the-css-user-agent-style-sheet-and-presentational-hints:'aspect-ratio'
x-internal="'aspect-ratio'"} property of the form
`auto ``w`{.variable}` / ``h`{.variable}.

When the text below says that a pair of attributes `w`{.variable} and
`h`{.variable} on an element `element`{.variable} [map to the
aspect-ratio property (using dimension
rules)]{#map-to-the-aspect-ratio-property-(using-dimension-rules) .dfn},
it means that if `element`{.variable} has both attributes `w`{.variable}
and `h`{.variable}, and parsing those attributes\' values using the
[rules for parsing dimension
values](common-microsyntaxes.html#rules-for-parsing-dimension-values){#the-css-user-agent-style-sheet-and-presentational-hints:rules-for-parsing-dimension-values-2}
doesn\'t generate an error or return a percentage for either, then the
user agent is expected to use the parsed dimensions as a [presentational
hint](#presentational-hints){#the-css-user-agent-style-sheet-and-presentational-hints:presentational-hints-5}
for the
[\'aspect-ratio\'](https://drafts.csswg.org/css-sizing-4/#aspect-ratio){#the-css-user-agent-style-sheet-and-presentational-hints:'aspect-ratio'-2
x-internal="'aspect-ratio'"} property of the form
`auto ``w`{.variable}` / ``h`{.variable}.

When a user agent is to [align descendants]{#align-descendants .dfn} of
a node, the user agent is expected to align only those descendants that
have both their
[\'margin-inline-start\'](https://drafts.csswg.org/css-logical/#propdef-margin-inline-start){#the-css-user-agent-style-sheet-and-presentational-hints:'margin-inline-start'
x-internal="'margin-inline-start'"} and
[\'margin-inline-end\'](https://drafts.csswg.org/css-logical/#propdef-margin-inline-end){#the-css-user-agent-style-sheet-and-presentational-hints:'margin-inline-end'
x-internal="'margin-inline-end'"} properties computing to a value other
than \'auto\', that are over-constrained and that have one of those two
margins with a [used
value](https://drafts.csswg.org/css-cascade/#used-value){#the-css-user-agent-style-sheet-and-presentational-hints:used-value
x-internal="used-value"} forced to a greater value, and that do not
themselves have an applicable `align` attribute. When multiple elements
are to
[align](#align-descendants){#the-css-user-agent-style-sheet-and-presentational-hints:align-descendants}
a particular descendant, the most deeply nested such element is expected
to override the others. Aligned elements are expected to be aligned by
having the [used
values](https://drafts.csswg.org/css-cascade/#used-value){#the-css-user-agent-style-sheet-and-presentational-hints:used-value-2
x-internal="used-value"} of their margins on the
[line-left](https://drafts.csswg.org/css-writing-modes/#line-left){#the-css-user-agent-style-sheet-and-presentational-hints:line-left
x-internal="line-left"} and
[line-right](https://drafts.csswg.org/css-writing-modes/#line-right){#the-css-user-agent-style-sheet-and-presentational-hints:line-right
x-internal="line-right"} sides be set accordingly.
[\[CSSLOGICAL\]](references.html#refsCSSLOGICAL)
[\[CSSWM\]](references.html#refsCSSWM)

### [15.3]{.secno} Non-replaced elements[](#non-replaced-elements){.self-link}

#### [15.3.1]{.secno} Hidden elements[](#hidden-elements){.self-link}

``` css
@namespace "http://www.w3.org/1999/xhtml";

area, base, basefont, datalist, head, link, meta, noembed,
noframes, param, rp, script, style, template, title {
  display: none;
}

[hidden]:not([hidden=until-found i]):not(embed) {
  display: none;
}

[hidden=until-found i]:not(embed) {
  content-visibility: hidden;
}

embed[hidden] { display: inline; height: 0; width: 0; } 

input[type=hidden i] { display: none !important; }

@media (scripting) {
  noscript { display: none !important; }
}
```

#### [15.3.2]{.secno} The page[](#the-page){.self-link}

``` css
@namespace "http://www.w3.org/1999/xhtml";

html, body { display: block; }
```

For each property in the table below, given a
[`body`{#the-page:the-body-element}](sections.html#the-body-element)
element, the first attribute that exists [maps to the pixel length
property](#maps-to-the-pixel-length-property){#the-page:maps-to-the-pixel-length-property}
on the
[`body`{#the-page:the-body-element-2}](sections.html#the-body-element)
element. If none of the attributes for a property are found, or if the
value of the attribute that was found cannot be parsed successfully,
then a default value of 8px is expected to be used for that property
instead.

Property

Source

[\'margin-top\'](https://drafts.csswg.org/css-box/#propdef-margin-top){#the-page:'margin-top'
x-internal="'margin-top'"}

The
[`body`{#the-page:the-body-element-3}](sections.html#the-body-element)
element\'s
[`marginheight`{#the-page:attr-body-marginheight}](obsolete.html#attr-body-marginheight)
attribute

The
[`body`{#the-page:the-body-element-4}](sections.html#the-body-element)
element\'s
[`topmargin`{#the-page:attr-body-topmargin}](obsolete.html#attr-body-topmargin)
attribute

The
[`body`{#the-page:the-body-element-5}](sections.html#the-body-element)
element\'s [container frame
element](#container-frame-element){#the-page:container-frame-element}\'s
[`marginheight`{#the-page:attr-iframe-marginheight}](obsolete.html#attr-iframe-marginheight)
attribute

[\'margin-right\'](https://drafts.csswg.org/css-box/#propdef-margin-right){#the-page:'margin-right'
x-internal="'margin-right'"}

The
[`body`{#the-page:the-body-element-6}](sections.html#the-body-element)
element\'s
[`marginwidth`{#the-page:attr-body-marginwidth}](obsolete.html#attr-body-marginwidth)
attribute

The
[`body`{#the-page:the-body-element-7}](sections.html#the-body-element)
element\'s
[`rightmargin`{#the-page:attr-body-rightmargin}](obsolete.html#attr-body-rightmargin)
attribute

The
[`body`{#the-page:the-body-element-8}](sections.html#the-body-element)
element\'s [container frame
element](#container-frame-element){#the-page:container-frame-element-2}\'s
[`marginwidth`{#the-page:attr-iframe-marginwidth}](obsolete.html#attr-iframe-marginwidth)
attribute

[\'margin-bottom\'](https://drafts.csswg.org/css-box/#propdef-margin-bottom){#the-page:'margin-bottom'
x-internal="'margin-bottom'"}

The
[`body`{#the-page:the-body-element-9}](sections.html#the-body-element)
element\'s
[`marginheight`{#the-page:attr-body-marginheight-2}](obsolete.html#attr-body-marginheight)
attribute

The
[`body`{#the-page:the-body-element-10}](sections.html#the-body-element)
element\'s
[`bottommargin`{#the-page:attr-body-bottommargin}](obsolete.html#attr-body-bottommargin)
attribute

The
[`body`{#the-page:the-body-element-11}](sections.html#the-body-element)
element\'s [container frame
element](#container-frame-element){#the-page:container-frame-element-3}\'s
[`marginheight`{#the-page:attr-iframe-marginheight-2}](obsolete.html#attr-iframe-marginheight)
attribute

[\'margin-left\'](https://drafts.csswg.org/css-box/#propdef-margin-left){#the-page:'margin-left'
x-internal="'margin-left'"}

The
[`body`{#the-page:the-body-element-12}](sections.html#the-body-element)
element\'s
[`marginwidth`{#the-page:attr-body-marginwidth-2}](obsolete.html#attr-body-marginwidth)
attribute

The
[`body`{#the-page:the-body-element-13}](sections.html#the-body-element)
element\'s
[`leftmargin`{#the-page:attr-body-leftmargin}](obsolete.html#attr-body-leftmargin)
attribute

The
[`body`{#the-page:the-body-element-14}](sections.html#the-body-element)
element\'s [container frame
element](#container-frame-element){#the-page:container-frame-element-4}\'s
[`marginwidth`{#the-page:attr-iframe-marginwidth-2}](obsolete.html#attr-iframe-marginwidth)
attribute

If the
[`body`{#the-page:the-body-element-15}](sections.html#the-body-element)
element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-page:node-document
x-internal="node-document"}\'s [node
navigable](document-sequences.html#node-navigable){#the-page:node-navigable}
is a [child
navigable](document-sequences.html#child-navigable){#the-page:child-navigable},
and the
[container](document-sequences.html#nav-container){#the-page:nav-container}
of that
[navigable](document-sequences.html#navigable){#the-page:navigable} is a
[`frame`{#the-page:frame}](obsolete.html#frame) or
[`iframe`{#the-page:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
element, then the [container frame element]{#container-frame-element
.dfn} of the
[`body`{#the-page:the-body-element-16}](sections.html#the-body-element)
element is that [`frame`{#the-page:frame-2}](obsolete.html#frame) or
[`iframe`{#the-page:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
element. Otherwise, there is no [container frame
element](#container-frame-element){#the-page:container-frame-element-5}.

The above requirements imply that a page can change the margins of
another page (including one from another
[origin](browsers.html#concept-origin){#the-page:concept-origin}) using,
for example, an
[`iframe`{#the-page:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element).
This is potentially a security risk, as it might in some cases allow an
attack to contrive a situation in which a page is rendered not as the
author intended, possibly for the purposes of phishing or otherwise
misleading the user.

------------------------------------------------------------------------

If a [`Document`{#the-page:document}](dom.html#document)\'s [node
navigable](document-sequences.html#node-navigable){#the-page:node-navigable-2}
is a [child
navigable](document-sequences.html#child-navigable){#the-page:child-navigable-2},
then it is expected to be positioned and sized to fit inside the
[content
box](https://drafts.csswg.org/css-box/#content-box){#the-page:content-box
x-internal="content-box"} of the
[container](document-sequences.html#nav-container){#the-page:nav-container-2}
of that
[navigable](document-sequences.html#navigable){#the-page:navigable-2}.
If the
[container](document-sequences.html#nav-container){#the-page:nav-container-3}
is not [being rendered](#being-rendered){#the-page:being-rendered}, the
[navigable](document-sequences.html#navigable){#the-page:navigable-3} is
expected to have a
[viewport](https://drafts.csswg.org/css2/#viewport){#the-page:viewport
x-internal="viewport"} with zero width and zero height.

If a [`Document`{#the-page:document-2}](dom.html#document)\'s [node
navigable](document-sequences.html#node-navigable){#the-page:node-navigable-3}
is a [child
navigable](document-sequences.html#child-navigable){#the-page:child-navigable-3},
the
[container](document-sequences.html#nav-container){#the-page:nav-container-4}
of that
[navigable](document-sequences.html#navigable){#the-page:navigable-4} is
a [`frame`{#the-page:frame-3}](obsolete.html#frame) or
[`iframe`{#the-page:the-iframe-element-4}](iframe-embed-object.html#the-iframe-element)
element, that element has a `scrolling` attribute, and that attribute\'s
value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-page:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string \"`off`\",
\"`noscroll`\", or \"`no`\", then the user agent is expected to prevent
any scrollbars from being shown for the
[viewport](https://drafts.csswg.org/css2/#viewport){#the-page:viewport-2
x-internal="viewport"} of the
[`Document`{#the-page:document-3}](dom.html#document)\'s [node
navigable](document-sequences.html#node-navigable){#the-page:node-navigable-4},
regardless of the
[\'overflow\'](https://drafts.csswg.org/css-overflow/#propdef-overflow){#the-page:'overflow'
x-internal="'overflow'"} property that applies to that
[viewport](https://drafts.csswg.org/css2/#viewport){#the-page:viewport-3
x-internal="viewport"}.

------------------------------------------------------------------------

When a
[`body`{#the-page:the-body-element-17}](sections.html#the-body-element)
element has a
[`background`{#the-page:attr-background}](obsolete.html#attr-background)
attribute set to a non-empty value, the new value is expected to be
[encoding-parsed-and-serialized](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#the-page:encoding-parsing-and-serializing-a-url}
relative to the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-page:node-document-2
x-internal="node-document"}, and if that does not return failure, the
user agent is expected to treat the attribute as a [presentational
hint](#presentational-hints){#the-page:presentational-hints} setting the
element\'s
[\'background-image\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-image){#the-page:'background-image'
x-internal="'background-image'"} property to the return value.

When a
[`body`{#the-page:the-body-element-18}](sections.html#the-body-element)
element has a
[`bgcolor`{#the-page:attr-body-bgcolor}](obsolete.html#attr-body-bgcolor)
attribute set, the new value is expected to be parsed using the [rules
for parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#the-page:rules-for-parsing-a-legacy-colour-value},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#the-page:presentational-hints-2} setting
the element\'s
[\'background-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-color){#the-page:'background-color'
x-internal="'background-color'"} property to the resulting color.

When a
[`body`{#the-page:the-body-element-19}](sections.html#the-body-element)
element has a
[`text`{#the-page:attr-body-text}](obsolete.html#attr-body-text)
attribute, its value is expected to be parsed using the [rules for
parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#the-page:rules-for-parsing-a-legacy-colour-value-2},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#the-page:presentational-hints-3} setting
the element\'s
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#the-page:'color'
x-internal="'color'"} property to the resulting color.

When a
[`body`{#the-page:the-body-element-20}](sections.html#the-body-element)
element has a
[`link`{#the-page:attr-body-link}](obsolete.html#attr-body-link)
attribute, its value is expected to be parsed using the [rules for
parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#the-page:rules-for-parsing-a-legacy-colour-value-3},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#the-page:presentational-hints-4} setting
the
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#the-page:'color'-2
x-internal="'color'"} property of any element in the
[`Document`{#the-page:document-4}](dom.html#document) matching the
[`:link`{#the-page:selector-link}](semantics-other.html#selector-link)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#the-page:pseudo-class
x-internal="pseudo-class"} to the resulting color.

When a
[`body`{#the-page:the-body-element-21}](sections.html#the-body-element)
element has a
[`vlink`{#the-page:attr-body-vlink}](obsolete.html#attr-body-vlink)
attribute, its value is expected to be parsed using the [rules for
parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#the-page:rules-for-parsing-a-legacy-colour-value-4},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#the-page:presentational-hints-5} setting
the
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#the-page:'color'-3
x-internal="'color'"} property of any element in the
[`Document`{#the-page:document-5}](dom.html#document) matching the
[`:visited`{#the-page:selector-visited}](semantics-other.html#selector-visited)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#the-page:pseudo-class-2
x-internal="pseudo-class"} to the resulting color.

When a
[`body`{#the-page:the-body-element-22}](sections.html#the-body-element)
element has an
[`alink`{#the-page:attr-body-alink}](obsolete.html#attr-body-alink)
attribute, its value is expected to be parsed using the [rules for
parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#the-page:rules-for-parsing-a-legacy-colour-value-5},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#the-page:presentational-hints-6} setting
the
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#the-page:'color'-4
x-internal="'color'"} property of any element in the
[`Document`{#the-page:document-6}](dom.html#document) matching the
[`:active`{#the-page:selector-active}](semantics-other.html#selector-active)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#the-page:pseudo-class-3
x-internal="pseudo-class"} and either the
[`:link`{#the-page:selector-link-2}](semantics-other.html#selector-link)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#the-page:pseudo-class-4
x-internal="pseudo-class"} or the
[`:visited`{#the-page:selector-visited-2}](semantics-other.html#selector-visited)
[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class){#the-page:pseudo-class-5
x-internal="pseudo-class"} to the resulting color.

#### [15.3.3]{.secno} Flow content[](#flow-content-3){.self-link} {#flow-content-3}

``` css
@namespace "http://www.w3.org/1999/xhtml";

address, blockquote, center, dialog, div, figure, figcaption, footer, form,
header, hr, legend, listing, main, p, plaintext, pre, search, xmp {
  display: block;
}

blockquote, figure, listing, p, plaintext, pre, xmp {
  margin-block: 1em;
}

blockquote, figure { margin-inline: 40px; }

address { font-style: italic; }
listing, plaintext, pre, xmp {
  font-family: monospace; white-space: pre;
}

dialog:not([open]) { display: none; }
dialog {
  position: absolute;
  inset-inline-start: 0; inset-inline-end: 0;
  width: fit-content;
  height: fit-content;
  margin: auto;
  border: solid;
  padding: 1em;
  background-color: Canvas;
  color: CanvasText;
}
dialog:modal {
  position: fixed;
  overflow: auto;
  inset-block: 0;
  max-width: calc(100% - 6px - 2em);
  max-height: calc(100% - 6px - 2em);
}
dialog::backdrop {
  background: rgba(0,0,0,0.1);
}

[popover]:not(:popover-open):not(dialog[open]) {
  display:none;
}

dialog:popover-open {
  display:block;
}

[popover] {
  position: fixed;
  inset: 0;
  width: fit-content;
  height: fit-content;
  margin: auto;
  border: solid;
  padding: 0.25em;
  overflow: auto;
  color: CanvasText;
  background-color: Canvas;
}

:popover-open::backdrop {
  position: fixed;
  inset: 0;
  pointer-events: none !important;
  background-color: transparent;
}

slot {
  display: contents;
}
```

The following rules are also expected to apply, as [presentational
hints](#presentational-hints){#flow-content-3:presentational-hints}:

``` css
@namespace "http://www.w3.org/1999/xhtml";

pre[wrap] { white-space: pre-wrap; }
```

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#flow-content-3:quirks-mode
x-internal="quirks-mode"}, the following rules are also expected to
apply:

``` css
@namespace "http://www.w3.org/1999/xhtml";

form { margin-block-end: 1em; }
```

------------------------------------------------------------------------

The [`center`{#flow-content-3:center}](obsolete.html#center) element,
and the
[`div`{#flow-content-3:the-div-element}](grouping-content.html#the-div-element)
element when it has an
[`align`{#flow-content-3:attr-div-align}](obsolete.html#attr-div-align)
attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#flow-content-3:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for either the string
\"`center`\" or the string \"`middle`\", are expected to center text
within themselves, as if they had their
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#flow-content-3:'text-align'
x-internal="'text-align'"} property set to \'center\' in a
[presentational
hint](#presentational-hints){#flow-content-3:presentational-hints-2},
and to [align
descendants](#align-descendants){#flow-content-3:align-descendants} to
the center.

The
[`div`{#flow-content-3:the-div-element-2}](grouping-content.html#the-div-element)
element, when it has an
[`align`{#flow-content-3:attr-div-align-2}](obsolete.html#attr-div-align)
attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#flow-content-3:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for the string \"`left`\", is
expected to left-align text within itself, as if it had its
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#flow-content-3:'text-align'-2
x-internal="'text-align'"} property set to \'left\' in a [presentational
hint](#presentational-hints){#flow-content-3:presentational-hints-3},
and to [align
descendants](#align-descendants){#flow-content-3:align-descendants-2} to
the left.

The
[`div`{#flow-content-3:the-div-element-3}](grouping-content.html#the-div-element)
element, when it has an
[`align`{#flow-content-3:attr-div-align-3}](obsolete.html#attr-div-align)
attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#flow-content-3:ascii-case-insensitive-3
x-internal="ascii-case-insensitive"} match for the string \"`right`\",
is expected to right-align text within itself, as if it had its
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#flow-content-3:'text-align'-3
x-internal="'text-align'"} property set to \'right\' in a
[presentational
hint](#presentational-hints){#flow-content-3:presentational-hints-4},
and to [align
descendants](#align-descendants){#flow-content-3:align-descendants-3} to
the right.

The
[`div`{#flow-content-3:the-div-element-4}](grouping-content.html#the-div-element)
element, when it has an
[`align`{#flow-content-3:attr-div-align-4}](obsolete.html#attr-div-align)
attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#flow-content-3:ascii-case-insensitive-4
x-internal="ascii-case-insensitive"} match for the string \"`justify`\",
is expected to full-justify text within itself, as if it had its
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#flow-content-3:'text-align'-4
x-internal="'text-align'"} property set to \'justify\' in a
[presentational
hint](#presentational-hints){#flow-content-3:presentational-hints-5},
and to [align
descendants](#align-descendants){#flow-content-3:align-descendants-4} to
the left.

#### [15.3.4]{.secno} Phrasing content[](#phrasing-content-3){.self-link} {#phrasing-content-3}

``` css
@namespace "http://www.w3.org/1999/xhtml";

cite, dfn, em, i, var { font-style: italic; }
b, strong { font-weight: bolder; }
code, kbd, samp, tt { font-family: monospace; }
big { font-size: larger; }
small { font-size: smaller; }

sub { vertical-align: sub; }
sup { vertical-align: super; }
sub, sup { line-height: normal; font-size: smaller; }

ruby { display: ruby; }
rt { display: ruby-text; }

:link { color: #0000EE; }
:visited { color: #551A8B; }
:link:active, :visited:active { color: #FF0000; }
:link, :visited { text-decoration: underline; cursor: pointer; }

:focus-visible { outline: auto; }

mark { background: yellow; color: black; } /* this color is just a suggestion and can be changed based on implementation feedback */

abbr[title], acronym[title] { text-decoration: dotted underline; }
ins, u { text-decoration: underline; }
del, s, strike { text-decoration: line-through; }

q::before { content: open-quote; }
q::after { content: close-quote; }

br { display-outside: newline; } /* this also has bidi implications */
nobr { white-space: nowrap; }
wbr { display-outside: break-opportunity; } /* this also has bidi implications */
nobr wbr { white-space: normal; }
```

The following rules are also expected to apply, as [presentational
hints](#presentational-hints){#phrasing-content-3:presentational-hints}:

``` css
@namespace "http://www.w3.org/1999/xhtml";

br[clear=left i] { clear: left; }
br[clear=right i] { clear: right; }
br[clear=all i], br[clear=both i] { clear: both; }
```

For the purposes of the CSS ruby model, runs of children of
[`ruby`{#phrasing-content-3:the-ruby-element}](text-level-semantics.html#the-ruby-element)
elements that are not
[`rt`{#phrasing-content-3:the-rt-element}](text-level-semantics.html#the-rt-element)
or
[`rp`{#phrasing-content-3:the-rp-element}](text-level-semantics.html#the-rp-element)
elements are expected to be wrapped in anonymous boxes whose
[\'display\'](https://drafts.csswg.org/css2/#display-prop){#phrasing-content-3:'display'
x-internal="'display'"} property has the value
[\'ruby-base\'](https://drafts.csswg.org/css-ruby/#valdef-display-ruby-base){#phrasing-content-3:'ruby-base'
x-internal="'ruby-base'"}. [\[CSSRUBY\]](references.html#refsCSSRUBY)

When a particular part of a ruby has more than one annotation, the
annotations should be distributed on both sides of the base text so as
to minimize the stacking of ruby annotations on one side.

When it becomes possible to do so, the preceding requirement will be
updated to be expressed in terms of CSS ruby. (Currently, CSS ruby does
not handle nested
[`ruby`{#phrasing-content-3:the-ruby-element-2}](text-level-semantics.html#the-ruby-element)
elements or multiple sequential
[`rt`{#phrasing-content-3:the-rt-element-2}](text-level-semantics.html#the-rt-element)
elements, which is how this semantic is expressed.)

User agents that do not support correct ruby rendering are expected to
render parentheses around the text of
[`rt`{#phrasing-content-3:the-rt-element-3}](text-level-semantics.html#the-rt-element)
elements in the absence of
[`rp`{#phrasing-content-3:the-rp-element-2}](text-level-semantics.html#the-rp-element)
elements.

------------------------------------------------------------------------

User agents are expected to support the
[\'clear\'](https://drafts.csswg.org/css2/#flow-control){#phrasing-content-3:'clear'
x-internal="'clear'"} property on inline elements (in order to render
[`br`{#phrasing-content-3:the-br-element}](text-level-semantics.html#the-br-element)
elements with
[`clear`{#phrasing-content-3:attr-br-clear}](obsolete.html#attr-br-clear)
attributes) in the manner described in the non-normative note to this
effect in CSS.

The initial value for the
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#phrasing-content-3:'color'
x-internal="'color'"} property is expected to be black. The initial
value for the
[\'background-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-color){#phrasing-content-3:'background-color'
x-internal="'background-color'"} property is expected to be
\'transparent\'. The canvas\'s background is expected to be white.

------------------------------------------------------------------------

When a [`font`{#phrasing-content-3:font}](obsolete.html#font) element
has a `color` attribute, its value is expected to be parsed using the
[rules for parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#phrasing-content-3:rules-for-parsing-a-legacy-colour-value},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#phrasing-content-3:presentational-hints-2}
setting the element\'s
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#phrasing-content-3:'color'-2
x-internal="'color'"} property to the resulting color.

When a [`font`{#phrasing-content-3:font-2}](obsolete.html#font) element
has a `face` attribute, the user agent is expected to treat the
attribute as a [presentational
hint](#presentational-hints){#phrasing-content-3:presentational-hints-3}
setting the element\'s
[\'font-family\'](https://drafts.csswg.org/css-fonts/#font-family-prop){#phrasing-content-3:'font-family'
x-internal="'font-family'"} property to the attribute\'s value.

When a [`font`{#phrasing-content-3:font-3}](obsolete.html#font) element
has a `size` attribute, the user agent is expected to use the following
steps, known as the [rules for parsing a legacy font
size]{#rules-for-parsing-a-legacy-font-size .dfn}, to treat the
attribute as a [presentational
hint](#presentational-hints){#phrasing-content-3:presentational-hints-4}
setting the element\'s
[\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop){#phrasing-content-3:'font-size'
x-internal="'font-size'"} property:

1.  Let `input`{.variable} be the attribute\'s value.

2.  Let `position`{.variable} be a pointer into `input`{.variable},
    initially pointing at the start of the string.

3.  [Skip ASCII
    whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace){#phrasing-content-3:skip-ascii-whitespace
    x-internal="skip-ascii-whitespace"} within `input`{.variable} given
    `position`{.variable}.

4.  If `position`{.variable} is past the end of `input`{.variable},
    there is no [presentational
    hint](#presentational-hints){#phrasing-content-3:presentational-hints-5}.
    Return.

5.  If the character at `position`{.variable} is a U+002B PLUS SIGN
    character (+), then let `mode`{.variable} be *relative-plus*, and
    advance `position`{.variable} to the next character. Otherwise, if
    the character at `position`{.variable} is a U+002D HYPHEN-MINUS
    character (-), then let `mode`{.variable} be *relative-minus*, and
    advance `position`{.variable} to the next character. Otherwise, let
    `mode`{.variable} be *absolute*.

6.  [Collect a sequence of code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#phrasing-content-3:collect-a-sequence-of-code-points
    x-internal="collect-a-sequence-of-code-points"} that are [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#phrasing-content-3:ascii-digits
    x-internal="ascii-digits"} from `input`{.variable} given
    `position`{.variable}, and let the resulting sequence be
    `digits`{.variable}.

7.  If `digits`{.variable} is the empty string, there is no
    [presentational
    hint](#presentational-hints){#phrasing-content-3:presentational-hints-6}.
    Return.

8.  Interpret `digits`{.variable} as a base-ten integer. Let
    `value`{.variable} be the resulting number.

9.  If `mode`{.variable} is *relative-plus*, then increment
    `value`{.variable} by 3. If `mode`{.variable} is *relative-minus*,
    then let `value`{.variable} be the result of subtracting
    `value`{.variable} from 3.

10. If `value`{.variable} is greater than 7, let it be 7.

11. If `value`{.variable} is less than 1, let it be 1.

12. Set
    [\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop){#phrasing-content-3:'font-size'-2
    x-internal="'font-size'"} to the keyword corresponding to the value
    of `value`{.variable} according to the following table:

    `value`{.variable}

    [\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop){#phrasing-content-3:'font-size'-3
    x-internal="'font-size'"} keyword

    1

    \'x-small\'

    2

    \'small\'

    3

    \'medium\'

    4

    \'large\'

    5

    \'x-large\'

    6

    \'xx-large\'

    7

    \'xxx-large\'

#### [15.3.5]{.secno} Bidirectional text[](#bidi-rendering){.self-link} {#bidi-rendering}

``` css
@namespace "http://www.w3.org/1999/xhtml";

[dir]:dir(ltr), bdi:dir(ltr), input[type=tel i]:dir(ltr) { direction: ltr; }
[dir]:dir(rtl), bdi:dir(rtl) { direction: rtl; }

address, blockquote, center, div, figure, figcaption, footer, form, header, hr,
legend, listing, main, p, plaintext, pre, summary, xmp, article, aside, h1, h2,
h3, h4, h5, h6, hgroup, nav, section, search, table, caption, colgroup, col,
thead, tbody, tfoot, tr, td, th, dir, dd, dl, dt, menu, ol, ul, li, bdi, output,
[dir=ltr i], [dir=rtl i], [dir=auto i] {
  unicode-bidi: isolate; 
}

bdo, bdo[dir] { unicode-bidi: isolate-override; } 

input[dir=auto i]:is([type=search i], [type=tel i], [type=url i],
[type=email i]), textarea[dir=auto i], pre[dir=auto i] {
  unicode-bidi: plaintext;
}
/* see prose for input elements whose type attribute is in the Text state */

/* the rules setting the 'content' property on br and wbr elements also has bidi implications */
```

When an
[`input`{#bidi-rendering:the-input-element}](input.html#the-input-element)
element\'s [`dir`{#bidi-rendering:attr-dir}](dom.html#attr-dir)
attribute is in the
[auto](dom.html#attr-dir-auto){#bidi-rendering:attr-dir-auto} state and
its
[`type`{#bidi-rendering:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#bidi-rendering:text-(type=text)-state-and-search-state-(type=search)}
state, then the user agent is expected to act as if it had a
user-agent-level style sheet rule setting the
[\'unicode-bidi\'](https://drafts.csswg.org/css-writing-modes/#unicode-bidi){#bidi-rendering:'unicode-bidi'
x-internal="'unicode-bidi'"} property to \'plaintext\'.

Input fields (i.e.
[`textarea`{#bidi-rendering:the-textarea-element}](form-elements.html#the-textarea-element)
elements, and
[`input`{#bidi-rendering:the-input-element-2}](input.html#the-input-element)
elements when their
[`type`{#bidi-rendering:attr-input-type-2}](input.html#attr-input-type)
attribute is in the
[Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#bidi-rendering:text-(type=text)-state-and-search-state-(type=search)-2},
[Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#bidi-rendering:text-(type=text)-state-and-search-state-(type=search)-3},
[Telephone](input.html#telephone-state-(type=tel)){#bidi-rendering:telephone-state-(type=tel)},
[URL](input.html#url-state-(type=url)){#bidi-rendering:url-state-(type=url)},
or
[Email](input.html#email-state-(type=email)){#bidi-rendering:email-state-(type=email)}
state) are expected to present an editing user interface with a
directionality that matches the element\'s
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#bidi-rendering:'direction'
x-internal="'direction'"} property.

When the document\'s character encoding is
[ISO-8859-8](https://encoding.spec.whatwg.org/#iso-8859-8){#bidi-rendering:iso-8859-8
x-internal="iso-8859-8"}, the following rules are additionally expected
to apply, following those above:
[\[ENCODING\]](references.html#refsENCODING)

``` css
@namespace "http://www.w3.org/1999/xhtml";

address, blockquote, center, div, figure, figcaption, footer, form, header, hr,
legend, listing, main, p, plaintext, pre, summary, xmp, article, aside, h1, h2,
h3, h4, h5, h6, hgroup, nav, section, search, table, caption, colgroup, col,
thead, tbody, tfoot, tr, td, th, dir, dd, dl, dt, menu, ol, ul, li, [dir=ltr i],
[dir=rtl i], [dir=auto i], *|* {
  unicode-bidi: bidi-override;
}
input:not([type=submit i]):not([type=reset i]):not([type=button i]),
textarea {
  unicode-bidi: normal;
}
```

#### [15.3.6]{.secno} Sections and headings[](#sections-and-headings){.self-link}

``` css
@namespace "http://www.w3.org/1999/xhtml";

article, aside, h1, h2, h3, h4, h5, h6, hgroup, nav, section {
  display: block;
}

h1 { margin-block: 0.67em; font-size: 2.00em; font-weight: bold; }
h2 { margin-block: 0.83em; font-size: 1.50em; font-weight: bold; }
h3 { margin-block: 1.00em; font-size: 1.17em; font-weight: bold; }
h4 { margin-block: 1.33em; font-size: 1.00em; font-weight: bold; }
h5 { margin-block: 1.67em; font-size: 0.83em; font-weight: bold; }
h6 { margin-block: 2.33em; font-size: 0.67em; font-weight: bold; }
```

In the following CSS block, `x`{.variable} is shorthand for the
following selector: `:is(article, aside, nav, section)`

``` css
@namespace "http://www.w3.org/1999/xhtml";

x h1 { margin-block: 0.83em; font-size: 1.50em; }
x x h1 { margin-block: 1.00em; font-size: 1.17em; }
x x x h1 { margin-block: 1.33em; font-size: 1.00em; }
x x x x h1 { margin-block: 1.67em; font-size: 0.83em; }
x x x x x h1 { margin-block: 2.33em; font-size: 0.67em; }
```

The shorthand is used to keep this block at least mildly readable.

#### [15.3.7]{.secno} Lists[](#lists){.self-link}

``` css
@namespace "http://www.w3.org/1999/xhtml";

dir, dd, dl, dt, menu, ol, ul { display: block; }
li { display: list-item; text-align: match-parent; }

dir, dl, menu, ol, ul { margin-block: 1em; }

:is(dir, dl, menu, ol, ul) :is(dir, dl, menu, ol, ul) {
  margin-block: 0;
}

dd { margin-inline-start: 40px; }
dir, menu, ol, ul { padding-inline-start: 40px; }

ol, ul, menu { counter-reset: list-item; }
ol { list-style-type: decimal; }

dir, menu, ul {
  list-style-type: disc;
}
:is(dir, menu, ol, ul) :is(dir, menu, ul) {
  list-style-type: circle;
}
:is(dir, menu, ol, ul) :is(dir, menu, ol, ul) :is(dir, menu, ul) {
  list-style-type: square;
}
```

The following rules are also expected to apply, as [presentational
hints](#presentational-hints){#lists:presentational-hints}:

``` css
@namespace "http://www.w3.org/1999/xhtml";

ol[type="1"], li[type="1"] { list-style-type: decimal; }
ol[type=a s], li[type=a s] { list-style-type: lower-alpha; }
ol[type=A s], li[type=A s] { list-style-type: upper-alpha; }
ol[type=i s], li[type=i s] { list-style-type: lower-roman; }
ol[type=I s], li[type=I s] { list-style-type: upper-roman; }
ul[type=none i], li[type=none i] { list-style-type: none; }
ul[type=disc i], li[type=disc i] { list-style-type: disc; }
ul[type=circle i], li[type=circle i] { list-style-type: circle; }
ul[type=square i], li[type=square i] { list-style-type: square; }
```

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#lists:quirks-mode
x-internal="quirks-mode"}, the following rules are also expected to
apply:

``` css
@namespace "http://www.w3.org/1999/xhtml";

li { list-style-position: inside; }
li :is(dir, menu, ol, ul) { list-style-position: outside; }
:is(dir, menu, ol, ul) :is(dir, menu, ol, ul, li) { list-style-position: unset; }
```

When rendering
[`li`{#lists:the-li-element}](grouping-content.html#the-li-element)
elements, non-CSS user agents are expected to use the [ordinal
value](grouping-content.html#ordinal-value){#lists:ordinal-value} of the
[`li`{#lists:the-li-element-2}](grouping-content.html#the-li-element)
element to render the counter in the list item marker.

For CSS user agents, some aspects of rendering [list
items](https://drafts.csswg.org/css-lists/#list-item){#lists:css-list-item
x-internal="css-list-item"} are defined by the CSS Lists specification.
Additionally, the following attribute mappings are expected to apply:
[\[CSSLISTS\]](references.html#refsCSSLISTS)

When an
[`li`{#lists:the-li-element-3}](grouping-content.html#the-li-element)
element has a
[`value`{#lists:attr-li-value}](grouping-content.html#attr-li-value)
attribute, and parsing that attribute\'s value using the [rules for
parsing
integers](common-microsyntaxes.html#rules-for-parsing-integers){#lists:rules-for-parsing-integers}
doesn\'t generate an error, the user agent is expected to use the parsed
value `value`{.variable} as a [presentational
hint](#presentational-hints){#lists:presentational-hints-2} for the
[\'counter-set\'](https://drafts.csswg.org/css-lists/#propdef-counter-set){#lists:'counter-set'
x-internal="'counter-set'"} property of the form
`list-item ``value`{.variable}.

When an
[`ol`{#lists:the-ol-element}](grouping-content.html#the-ol-element)
element has a
[`start`{#lists:attr-ol-start}](grouping-content.html#attr-ol-start)
attribute or a
[`reversed`{#lists:attr-ol-reversed}](grouping-content.html#attr-ol-reversed)
attribute, or both, the user agent is expected to use the following
steps to treat the attributes as a [presentational
hint](#presentational-hints){#lists:presentational-hints-3} for the
[\'counter-reset\'](https://drafts.csswg.org/css-lists/#propdef-counter-reset){#lists:'counter-reset'
x-internal="'counter-reset'"} property:

1.  Let `value`{.variable} be null.

2.  If the element has a
    [`start`{#lists:attr-ol-start-2}](grouping-content.html#attr-ol-start)
    attribute, then set `value`{.variable} to the result of parsing the
    attribute\'s value using the [rules for parsing
    integers](common-microsyntaxes.html#rules-for-parsing-integers){#lists:rules-for-parsing-integers-2}.

3.  If the element has a
    [`reversed`{#lists:attr-ol-reversed-2}](grouping-content.html#attr-ol-reversed)
    attribute, then:

    1.  If `value`{.variable} is an integer, then increment
        `value`{.variable} by 1 and return
        `reversed(list-item) ``value`{.variable}.

    2.  Otherwise, return `reversed(list-item)`.

        Either the
        [`start`{#lists:attr-ol-start-3}](grouping-content.html#attr-ol-start)
        attribute was absent, or parsing its value resulted in an error.

4.  Otherwise:

    1.  If `value`{.variable} is an integer, then decrement
        `value`{.variable} by 1 and return
        `list-item ``value`{.variable}.

    2.  Otherwise, there is no [presentational
        hint](#presentational-hints){#lists:presentational-hints-4}.

#### [15.3.8]{.secno} Tables[](#tables-2){.self-link} {#tables-2}

``` css
@namespace "http://www.w3.org/1999/xhtml";

table { display: table; }
caption { display: table-caption; }
colgroup, colgroup[hidden] { display: table-column-group; }
col, col[hidden] { display: table-column; }
thead, thead[hidden] { display: table-header-group; }
tbody, tbody[hidden] { display: table-row-group; }
tfoot, tfoot[hidden] { display: table-footer-group; }
tr, tr[hidden] { display: table-row; }
td, th { display: table-cell; }

colgroup[hidden], col[hidden], thead[hidden], tbody[hidden],
tfoot[hidden], tr[hidden] {
  visibility: collapse;
}

table {
  box-sizing: border-box;
  border-spacing: 2px;
  border-collapse: separate;
  text-indent: initial;
}
td, th { padding: 1px; }
th { font-weight: bold; }

caption { text-align: center; }
thead, tbody, tfoot, table > tr { vertical-align: middle; }
tr, td, th { vertical-align: inherit; }

thead, tbody, tfoot, tr { border-color: inherit; }
table[rules=none i], table[rules=groups i], table[rules=rows i],
table[rules=cols i], table[rules=all i], table[frame=void i],
table[frame=above i], table[frame=below i], table[frame=hsides i],
table[frame=lhs i], table[frame=rhs i], table[frame=vsides i],
table[frame=box i], table[frame=border i],
table[rules=none i] > tr > td, table[rules=none i] > tr > th,
table[rules=groups i] > tr > td, table[rules=groups i] > tr > th,
table[rules=rows i] > tr > td, table[rules=rows i] > tr > th,
table[rules=cols i] > tr > td, table[rules=cols i] > tr > th,
table[rules=all i] > tr > td, table[rules=all i] > tr > th,
table[rules=none i] > thead > tr > td, table[rules=none i] > thead > tr > th,
table[rules=groups i] > thead > tr > td, table[rules=groups i] > thead > tr > th,
table[rules=rows i] > thead > tr > td, table[rules=rows i] > thead > tr > th,
table[rules=cols i] > thead > tr > td, table[rules=cols i] > thead > tr > th,
table[rules=all i] > thead > tr > td, table[rules=all i] > thead > tr > th,
table[rules=none i] > tbody > tr > td, table[rules=none i] > tbody > tr > th,
table[rules=groups i] > tbody > tr > td, table[rules=groups i] > tbody > tr > th,
table[rules=rows i] > tbody > tr > td, table[rules=rows i] > tbody > tr > th,
table[rules=cols i] > tbody > tr > td, table[rules=cols i] > tbody > tr > th,
table[rules=all i] > tbody > tr > td, table[rules=all i] > tbody > tr > th,
table[rules=none i] > tfoot > tr > td, table[rules=none i] > tfoot > tr > th,
table[rules=groups i] > tfoot > tr > td, table[rules=groups i] > tfoot > tr > th,
table[rules=rows i] > tfoot > tr > td, table[rules=rows i] > tfoot > tr > th,
table[rules=cols i] > tfoot > tr > td, table[rules=cols i] > tfoot > tr > th,
table[rules=all i] > tfoot > tr > td, table[rules=all i] > tfoot > tr > th {
  border-color: black;
}
```

The following rules are also expected to apply, as [presentational
hints](#presentational-hints){#tables-2:presentational-hints}:

``` css
@namespace "http://www.w3.org/1999/xhtml";

table[align=left i] { float: left; }
table[align=right i] { float: right; }
table[align=center i] { margin-inline: auto; }
thead[align=absmiddle i], tbody[align=absmiddle i], tfoot[align=absmiddle i],
tr[align=absmiddle i], td[align=absmiddle i], th[align=absmiddle i] {
  text-align: center;
}

caption[align=bottom i] { caption-side: bottom; }
p[align=left i], h1[align=left i], h2[align=left i], h3[align=left i],
h4[align=left i], h5[align=left i], h6[align=left i] {
  text-align: left;
}
p[align=right i], h1[align=right i], h2[align=right i], h3[align=right i],
h4[align=right i], h5[align=right i], h6[align=right i] {
  text-align: right;
}
p[align=center i], h1[align=center i], h2[align=center i], h3[align=center i],
h4[align=center i], h5[align=center i], h6[align=center i] {
  text-align: center;
}
p[align=justify i], h1[align=justify i], h2[align=justify i], h3[align=justify i],
h4[align=justify i], h5[align=justify i], h6[align=justify i] {
  text-align: justify;
}
thead[valign=top i], tbody[valign=top i], tfoot[valign=top i],
tr[valign=top i], td[valign=top i], th[valign=top i] {
  vertical-align: top;
}
thead[valign=middle i], tbody[valign=middle i], tfoot[valign=middle i],
tr[valign=middle i], td[valign=middle i], th[valign=middle i] {
  vertical-align: middle;
}
thead[valign=bottom i], tbody[valign=bottom i], tfoot[valign=bottom i],
tr[valign=bottom i], td[valign=bottom i], th[valign=bottom i] {
  vertical-align: bottom;
}
thead[valign=baseline i], tbody[valign=baseline i], tfoot[valign=baseline i],
tr[valign=baseline i], td[valign=baseline i], th[valign=baseline i] {
  vertical-align: baseline;
}

td[nowrap], th[nowrap] { white-space: nowrap; }

table[rules=none i], table[rules=groups i], table[rules=rows i],
table[rules=cols i], table[rules=all i] {
  border-style: hidden;
  border-collapse: collapse;
}
table[border] { border-style: outset; } /* only if border is not equivalent to zero */
table[frame=void i] { border-style: hidden; }
table[frame=above i] { border-style: outset hidden hidden hidden; }
table[frame=below i] { border-style: hidden hidden outset hidden; }
table[frame=hsides i] { border-style: outset hidden outset hidden; }
table[frame=lhs i] { border-style: hidden hidden hidden outset; }
table[frame=rhs i] { border-style: hidden outset hidden hidden; }
table[frame=vsides i] { border-style: hidden outset; }
table[frame=box i], table[frame=border i] { border-style: outset; }

table[border] > tr > td, table[border] > tr > th,
table[border] > thead > tr > td, table[border] > thead > tr > th,
table[border] > tbody > tr > td, table[border] > tbody > tr > th,
table[border] > tfoot > tr > td, table[border] > tfoot > tr > th {
  /* only if border is not equivalent to zero */
  border-width: 1px;
  border-style: inset;
}
table[rules=none i] > tr > td, table[rules=none i] > tr > th,
table[rules=none i] > thead > tr > td, table[rules=none i] > thead > tr > th,
table[rules=none i] > tbody > tr > td, table[rules=none i] > tbody > tr > th,
table[rules=none i] > tfoot > tr > td, table[rules=none i] > tfoot > tr > th,
table[rules=groups i] > tr > td, table[rules=groups i] > tr > th,
table[rules=groups i] > thead > tr > td, table[rules=groups i] > thead > tr > th,
table[rules=groups i] > tbody > tr > td, table[rules=groups i] > tbody > tr > th,
table[rules=groups i] > tfoot > tr > td, table[rules=groups i] > tfoot > tr > th,
table[rules=rows i] > tr > td, table[rules=rows i] > tr > th,
table[rules=rows i] > thead > tr > td, table[rules=rows i] > thead > tr > th,
table[rules=rows i] > tbody > tr > td, table[rules=rows i] > tbody > tr > th,
table[rules=rows i] > tfoot > tr > td, table[rules=rows i] > tfoot > tr > th {
  border-width: 1px;
  border-style: none;
}
table[rules=cols i] > tr > td, table[rules=cols i] > tr > th,
table[rules=cols i] > thead > tr > td, table[rules=cols i] > thead > tr > th,
table[rules=cols i] > tbody > tr > td, table[rules=cols i] > tbody > tr > th,
table[rules=cols i] > tfoot > tr > td, table[rules=cols i] > tfoot > tr > th {
  border-width: 1px;
  border-block-style: none;
  border-inline-style: solid;
}
table[rules=all i] > tr > td, table[rules=all i] > tr > th,
table[rules=all i] > thead > tr > td, table[rules=all i] > thead > tr > th,
table[rules=all i] > tbody > tr > td, table[rules=all i] > tbody > tr > th,
table[rules=all i] > tfoot > tr > td, table[rules=all i] > tfoot > tr > th {
  border-width: 1px;
  border-style: solid;
}

table[rules=groups i] > colgroup {
  border-inline-width: 1px;
  border-inline-style: solid;
}
table[rules=groups i] > thead,
table[rules=groups i] > tbody,
table[rules=groups i] > tfoot {
  border-block-width: 1px;
  border-block-style: solid;
}

table[rules=rows i] > tr, table[rules=rows i] > thead > tr,
table[rules=rows i] > tbody > tr, table[rules=rows i] > tfoot > tr {
  border-block-width: 1px;
  border-block-style: solid;
}
```

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#tables-2:quirks-mode
x-internal="quirks-mode"}, the following rules are also expected to
apply:

``` css
@namespace "http://www.w3.org/1999/xhtml";

table {
  font-weight: initial;
  font-style: initial;
  font-variant: initial;
  font-size: initial;
  line-height: initial;
  white-space: initial;
  text-align: initial;
}
```

------------------------------------------------------------------------

For the purposes of the CSS table model, the
[`col`{#tables-2:the-col-element}](tables.html#the-col-element) element
is expected to be treated as if it was present as many times as its
[`span`{#tables-2:attr-col-span}](tables.html#attr-col-span) attribute
[specifies](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#tables-2:rules-for-parsing-non-negative-integers}.

For the purposes of the CSS table model, the
[`colgroup`{#tables-2:the-colgroup-element}](tables.html#the-colgroup-element)
element, if it contains no
[`col`{#tables-2:the-col-element-2}](tables.html#the-col-element)
element, is expected to be treated as if it had as many such children as
its
[`span`{#tables-2:attr-colgroup-span}](tables.html#attr-colgroup-span)
attribute
[specifies](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#tables-2:rules-for-parsing-non-negative-integers-2}.

For the purposes of the CSS table model, the
[`colspan`{#tables-2:attr-tdth-colspan}](tables.html#attr-tdth-colspan)
and
[`rowspan`{#tables-2:attr-tdth-rowspan}](tables.html#attr-tdth-rowspan)
attributes on
[`td`{#tables-2:the-td-element}](tables.html#the-td-element) and
[`th`{#tables-2:the-th-element}](tables.html#the-th-element) elements
are expected to
[provide](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#tables-2:rules-for-parsing-non-negative-integers-3}
the *special knowledge* regarding cells spanning rows and columns.

In [HTML
documents](https://dom.spec.whatwg.org/#html-document){#tables-2:html-documents
x-internal="html-documents"}, the following rules are also expected to
apply:

``` css
@namespace "http://www.w3.org/1999/xhtml";

:is(table, thead, tbody, tfoot, tr) > form { display: none !important; }
```

------------------------------------------------------------------------

The
[`table`{#tables-2:the-table-element}](tables.html#the-table-element)
element\'s
[`cellspacing`{#tables-2:attr-table-cellspacing}](obsolete.html#attr-table-cellspacing)
attribute [maps to the pixel length
property](#maps-to-the-pixel-length-property){#tables-2:maps-to-the-pixel-length-property}
[\'border-spacing\'](https://drafts.csswg.org/css-tables/#propdef-border-spacing){#tables-2:'border-spacing'
x-internal="'border-spacing'"} on the element.

The
[`table`{#tables-2:the-table-element-2}](tables.html#the-table-element)
element\'s
[`cellpadding`{#tables-2:attr-table-cellpadding}](obsolete.html#attr-table-cellpadding)
attribute [maps to the pixel length
properties](#maps-to-the-pixel-length-property){#tables-2:maps-to-the-pixel-length-property-2}
[\'padding-top\'](https://drafts.csswg.org/css-box/#propdef-padding-top){#tables-2:'padding-top'
x-internal="'padding-top'"},
[\'padding-right\'](https://drafts.csswg.org/css-box/#propdef-padding-right){#tables-2:'padding-right'
x-internal="'padding-right'"},
[\'padding-bottom\'](https://drafts.csswg.org/css-box/#propdef-padding-bottom){#tables-2:'padding-bottom'
x-internal="'padding-bottom'"}, and
[\'padding-left\'](https://drafts.csswg.org/css-box/#propdef-padding-left){#tables-2:'padding-left'
x-internal="'padding-left'"} of any
[`td`{#tables-2:the-td-element-2}](tables.html#the-td-element) and
[`th`{#tables-2:the-th-element-2}](tables.html#the-th-element) elements
that have corresponding
[cells](tables.html#concept-cell){#tables-2:concept-cell} in the
[table](tables.html#concept-table){#tables-2:concept-table}
corresponding to the
[`table`{#tables-2:the-table-element-3}](tables.html#the-table-element)
element.

The
[`table`{#tables-2:the-table-element-4}](tables.html#the-table-element)
element\'s
[`height`{#tables-2:attr-table-height}](obsolete.html#attr-table-height)
attribute [maps to the dimension
property](#maps-to-the-dimension-property){#tables-2:maps-to-the-dimension-property}
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#tables-2:'height'
x-internal="'height'"} on the
[`table`{#tables-2:the-table-element-5}](tables.html#the-table-element)
element.

The
[`table`{#tables-2:the-table-element-6}](tables.html#the-table-element)
element\'s
[`width`{#tables-2:attr-table-width}](obsolete.html#attr-table-width)
attribute [maps to the dimension property (ignoring
zero)](#maps-to-the-dimension-property-(ignoring-zero)){#tables-2:maps-to-the-dimension-property-(ignoring-zero)}
[\'width\'](https://drafts.csswg.org/css2/#the-width-property){#tables-2:'width'
x-internal="'width'"} on the
[`table`{#tables-2:the-table-element-7}](tables.html#the-table-element)
element.

The [`col`{#tables-2:the-col-element-3}](tables.html#the-col-element)
element\'s
[`width`{#tables-2:attr-col-width}](obsolete.html#attr-col-width)
attribute [maps to the dimension
property](#maps-to-the-dimension-property){#tables-2:maps-to-the-dimension-property-2}
[\'width\'](https://drafts.csswg.org/css2/#the-width-property){#tables-2:'width'-2
x-internal="'width'"} on the
[`col`{#tables-2:the-col-element-4}](tables.html#the-col-element)
element.

The
[`thead`{#tables-2:the-thead-element}](tables.html#the-thead-element),
[`tbody`{#tables-2:the-tbody-element}](tables.html#the-tbody-element),
and
[`tfoot`{#tables-2:the-tfoot-element}](tables.html#the-tfoot-element)
elements\'
[`height`{#tables-2:attr-tbody-height}](obsolete.html#attr-tbody-height)
attribute [maps to the dimension
property](#maps-to-the-dimension-property){#tables-2:maps-to-the-dimension-property-3}
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#tables-2:'height'-2
x-internal="'height'"} on the element.

The [`tr`{#tables-2:the-tr-element}](tables.html#the-tr-element)
element\'s
[`height`{#tables-2:attr-tr-height}](obsolete.html#attr-tr-height)
attribute [maps to the dimension
property](#maps-to-the-dimension-property){#tables-2:maps-to-the-dimension-property-4}
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#tables-2:'height'-3
x-internal="'height'"} on the
[`tr`{#tables-2:the-tr-element-2}](tables.html#the-tr-element) element.

The [`td`{#tables-2:the-td-element-3}](tables.html#the-td-element) and
[`th`{#tables-2:the-th-element-3}](tables.html#the-th-element)
elements\'
[`height`{#tables-2:attr-tdth-height}](obsolete.html#attr-tdth-height)
attributes [map to the dimension property (ignoring
zero)](#maps-to-the-dimension-property-(ignoring-zero)){#tables-2:maps-to-the-dimension-property-(ignoring-zero)-2}
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#tables-2:'height'-4
x-internal="'height'"} on the element.

The [`td`{#tables-2:the-td-element-4}](tables.html#the-td-element) and
[`th`{#tables-2:the-th-element-4}](tables.html#the-th-element)
elements\'
[`width`{#tables-2:attr-tdth-width}](obsolete.html#attr-tdth-width)
attributes [map to the dimension property (ignoring
zero)](#maps-to-the-dimension-property-(ignoring-zero)){#tables-2:maps-to-the-dimension-property-(ignoring-zero)-3}
[\'width\'](https://drafts.csswg.org/css2/#the-width-property){#tables-2:'width'-3
x-internal="'width'"} on the element.

------------------------------------------------------------------------

The
[`thead`{#tables-2:the-thead-element-2}](tables.html#the-thead-element),
[`tbody`{#tables-2:the-tbody-element-2}](tables.html#the-tbody-element),
[`tfoot`{#tables-2:the-tfoot-element-2}](tables.html#the-tfoot-element),
[`tr`{#tables-2:the-tr-element-3}](tables.html#the-tr-element),
[`td`{#tables-2:the-td-element-5}](tables.html#the-td-element), and
[`th`{#tables-2:the-th-element-5}](tables.html#the-th-element) elements,
when they have an `align` attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#tables-2:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for either the string
\"`center`\" or the string \"`middle`\", are expected to center text
within themselves, as if they had their
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#tables-2:'text-align'
x-internal="'text-align'"} property set to \'center\' in a
[presentational
hint](#presentational-hints){#tables-2:presentational-hints-2}, and to
[align descendants](#align-descendants){#tables-2:align-descendants} to
the center.

The
[`thead`{#tables-2:the-thead-element-3}](tables.html#the-thead-element),
[`tbody`{#tables-2:the-tbody-element-3}](tables.html#the-tbody-element),
[`tfoot`{#tables-2:the-tfoot-element-3}](tables.html#the-tfoot-element),
[`tr`{#tables-2:the-tr-element-4}](tables.html#the-tr-element),
[`td`{#tables-2:the-td-element-6}](tables.html#the-td-element), and
[`th`{#tables-2:the-th-element-6}](tables.html#the-th-element) elements,
when they have an `align` attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#tables-2:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for the string \"`left`\",
are expected to left-align text within themselves, as if they had their
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#tables-2:'text-align'-2
x-internal="'text-align'"} property set to \'left\' in a [presentational
hint](#presentational-hints){#tables-2:presentational-hints-3}, and to
[align descendants](#align-descendants){#tables-2:align-descendants-2}
to the left.

The
[`thead`{#tables-2:the-thead-element-4}](tables.html#the-thead-element),
[`tbody`{#tables-2:the-tbody-element-4}](tables.html#the-tbody-element),
[`tfoot`{#tables-2:the-tfoot-element-4}](tables.html#the-tfoot-element),
[`tr`{#tables-2:the-tr-element-5}](tables.html#the-tr-element),
[`td`{#tables-2:the-td-element-7}](tables.html#the-td-element), and
[`th`{#tables-2:the-th-element-7}](tables.html#the-th-element) elements,
when they have an `align` attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#tables-2:ascii-case-insensitive-3
x-internal="ascii-case-insensitive"} match for the string \"`right`\",
are expected to right-align text within themselves, as if they had their
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#tables-2:'text-align'-3
x-internal="'text-align'"} property set to \'right\' in a
[presentational
hint](#presentational-hints){#tables-2:presentational-hints-4}, and to
[align descendants](#align-descendants){#tables-2:align-descendants-3}
to the right.

The
[`thead`{#tables-2:the-thead-element-5}](tables.html#the-thead-element),
[`tbody`{#tables-2:the-tbody-element-5}](tables.html#the-tbody-element),
[`tfoot`{#tables-2:the-tfoot-element-5}](tables.html#the-tfoot-element),
[`tr`{#tables-2:the-tr-element-6}](tables.html#the-tr-element),
[`td`{#tables-2:the-td-element-8}](tables.html#the-td-element), and
[`th`{#tables-2:the-th-element-8}](tables.html#the-th-element) elements,
when they have an `align` attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#tables-2:ascii-case-insensitive-4
x-internal="ascii-case-insensitive"} match for the string \"`justify`\",
are expected to full-justify text within themselves, as if they had
their
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#tables-2:'text-align'-4
x-internal="'text-align'"} property set to \'justify\' in a
[presentational
hint](#presentational-hints){#tables-2:presentational-hints-5}, and to
[align descendants](#align-descendants){#tables-2:align-descendants-4}
to the left.

User agents are expected to have a rule in their user agent style sheet
that matches
[`th`{#tables-2:the-th-element-9}](tables.html#the-th-element) elements
that have a parent node whose [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#tables-2:computed-value
x-internal="computed-value"} for the
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#tables-2:'text-align'-5
x-internal="'text-align'"} property is its initial value, whose
declaration block consists of just a single declaration that sets the
[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property){#tables-2:'text-align'-6
x-internal="'text-align'"} property to the value \'center\'.

------------------------------------------------------------------------

When a
[`table`{#tables-2:the-table-element-8}](tables.html#the-table-element),
[`thead`{#tables-2:the-thead-element-6}](tables.html#the-thead-element),
[`tbody`{#tables-2:the-tbody-element-6}](tables.html#the-tbody-element),
[`tfoot`{#tables-2:the-tfoot-element-6}](tables.html#the-tfoot-element),
[`tr`{#tables-2:the-tr-element-7}](tables.html#the-tr-element),
[`td`{#tables-2:the-td-element-9}](tables.html#the-td-element), or
[`th`{#tables-2:the-th-element-10}](tables.html#the-th-element) element
has a
[`background`{#tables-2:attr-background}](obsolete.html#attr-background)
attribute set to a non-empty value, the new value is expected to be
[encoding-parsed-and-serialized](urls-and-fetching.html#encoding-parsing-and-serializing-a-url){#tables-2:encoding-parsing-and-serializing-a-url}
relative to the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#tables-2:node-document
x-internal="node-document"}, and if that does not return failure, the
user agent is expected to treat the attribute as a [presentational
hint](#presentational-hints){#tables-2:presentational-hints-6} setting
the element\'s
[\'background-image\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-image){#tables-2:'background-image'
x-internal="'background-image'"} property to the return value.

When a
[`table`{#tables-2:the-table-element-9}](tables.html#the-table-element),
[`thead`{#tables-2:the-thead-element-7}](tables.html#the-thead-element),
[`tbody`{#tables-2:the-tbody-element-7}](tables.html#the-tbody-element),
[`tfoot`{#tables-2:the-tfoot-element-7}](tables.html#the-tfoot-element),
[`tr`{#tables-2:the-tr-element-8}](tables.html#the-tr-element),
[`td`{#tables-2:the-td-element-10}](tables.html#the-td-element), or
[`th`{#tables-2:the-th-element-11}](tables.html#the-th-element) element
has a `bgcolor` attribute set, the new value is expected to be parsed
using the [rules for parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#tables-2:rules-for-parsing-a-legacy-colour-value},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#tables-2:presentational-hints-7} setting
the element\'s
[\'background-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-color){#tables-2:'background-color'
x-internal="'background-color'"} property to the resulting color.

When a
[`table`{#tables-2:the-table-element-10}](tables.html#the-table-element)
element has a
[`bordercolor`{#tables-2:attr-table-bordercolor}](obsolete.html#attr-table-bordercolor)
attribute, its value is expected to be parsed using the [rules for
parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#tables-2:rules-for-parsing-a-legacy-colour-value-2},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#tables-2:presentational-hints-8} setting
the element\'s
[\'border-top-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-color){#tables-2:'border-top-color'
x-internal="'border-top-color'"},
[\'border-right-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-color){#tables-2:'border-right-color'
x-internal="'border-right-color'"},
[\'border-bottom-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-color){#tables-2:'border-bottom-color'
x-internal="'border-bottom-color'"}, and
[\'border-left-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-color){#tables-2:'border-left-color'
x-internal="'border-left-color'"} properties to the resulting color.

------------------------------------------------------------------------

The
[`table`{#tables-2:the-table-element-11}](tables.html#the-table-element)
element\'s
[`border`{#tables-2:attr-table-border}](obsolete.html#attr-table-border)
attribute [maps to the pixel length
properties](#maps-to-the-pixel-length-property){#tables-2:maps-to-the-pixel-length-property-3}
[\'border-top-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-width){#tables-2:'border-top-width'
x-internal="'border-top-width'"},
[\'border-right-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-width){#tables-2:'border-right-width'
x-internal="'border-right-width'"},
[\'border-bottom-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-width){#tables-2:'border-bottom-width'
x-internal="'border-bottom-width'"},
[\'border-left-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-width){#tables-2:'border-left-width'
x-internal="'border-left-width'"} on the element. If the attribute is
present but parsing the attribute\'s value using the [rules for parsing
non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#tables-2:rules-for-parsing-non-negative-integers-4}
generates an error, a default value of 1px is expected to be used for
that property instead.

Rules marked \"[only if border is not equivalent to
zero]{#magic-border-selector .dfn}\" in the CSS block above is expected
to only be applied if the
[`border`{#tables-2:attr-table-border-2}](obsolete.html#attr-table-border)
attribute mentioned in the selectors for the rule is not only present
but, when parsed using the [rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#tables-2:rules-for-parsing-non-negative-integers-5},
is also found to have a value other than zero or to generate an error.

------------------------------------------------------------------------

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#tables-2:quirks-mode-2
x-internal="quirks-mode"}, a
[`td`{#tables-2:the-td-element-11}](tables.html#the-td-element) element
or a [`th`{#tables-2:the-th-element-12}](tables.html#the-th-element)
element that has a
[`nowrap`{#tables-2:attr-tdth-nowrap}](obsolete.html#attr-tdth-nowrap)
attribute but also has a
[`width`{#tables-2:attr-tdth-width-2}](obsolete.html#attr-tdth-width)
attribute whose value, when parsed using the [rules for parsing nonzero
dimension
values](common-microsyntaxes.html#rules-for-parsing-non-zero-dimension-values){#tables-2:rules-for-parsing-non-zero-dimension-values},
is found to be a length (not an error or a number classified as a
percentage), is expected to have a [presentational
hint](#presentational-hints){#tables-2:presentational-hints-9} setting
the element\'s
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#tables-2:'white-space'
x-internal="'white-space'"} property to \'normal\', overriding the rule
in the CSS block above that sets it to \'nowrap\'.

#### [15.3.9]{.secno} Margin collapsing quirks[](#margin-collapsing-quirks){.self-link}

A node is [substantial]{#concept-rendering-substantial .dfn} if it is a
text node that is not [inter-element
whitespace](dom.html#inter-element-whitespace){#margin-collapsing-quirks:inter-element-whitespace},
or if it is an element node.

A node is [blank]{#concept-rendering-blank .dfn} if it is an element
that contains no
[substantial](#concept-rendering-substantial){#margin-collapsing-quirks:concept-rendering-substantial}
nodes.

The [elements with default
margins]{#concept-rendering-elements-with-margins .dfn} are the
following elements:
[`blockquote`{#margin-collapsing-quirks:the-blockquote-element}](grouping-content.html#the-blockquote-element),
[`dir`{#margin-collapsing-quirks:dir}](obsolete.html#dir),
[`dl`{#margin-collapsing-quirks:the-dl-element}](grouping-content.html#the-dl-element),
[`h1`{#margin-collapsing-quirks:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h2`{#margin-collapsing-quirks:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h3`{#margin-collapsing-quirks:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h4`{#margin-collapsing-quirks:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h5`{#margin-collapsing-quirks:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h6`{#margin-collapsing-quirks:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`listing`{#margin-collapsing-quirks:listing}](obsolete.html#listing),
[`menu`{#margin-collapsing-quirks:the-menu-element}](grouping-content.html#the-menu-element),
[`ol`{#margin-collapsing-quirks:the-ol-element}](grouping-content.html#the-ol-element),
[`p`{#margin-collapsing-quirks:the-p-element}](grouping-content.html#the-p-element),
[`plaintext`{#margin-collapsing-quirks:plaintext}](obsolete.html#plaintext),
[`pre`{#margin-collapsing-quirks:the-pre-element}](grouping-content.html#the-pre-element),
[`ul`{#margin-collapsing-quirks:the-ul-element}](grouping-content.html#the-ul-element),
[`xmp`{#margin-collapsing-quirks:xmp}](obsolete.html#xmp)

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#margin-collapsing-quirks:quirks-mode
x-internal="quirks-mode"}, any [element with default
margins](#concept-rendering-elements-with-margins){#margin-collapsing-quirks:concept-rendering-elements-with-margins}
that is the
[child](https://dom.spec.whatwg.org/#concept-tree-child){#margin-collapsing-quirks:concept-tree-child
x-internal="concept-tree-child"} of a
[`body`{#margin-collapsing-quirks:the-body-element}](sections.html#the-body-element),
[`td`{#margin-collapsing-quirks:the-td-element}](tables.html#the-td-element),
or
[`th`{#margin-collapsing-quirks:the-th-element}](tables.html#the-th-element)
element and has no
[substantial](#concept-rendering-substantial){#margin-collapsing-quirks:concept-rendering-substantial-2}
previous siblings is expected to have a user-agent level style sheet
rule that sets its
[\'margin-block-start\'](https://drafts.csswg.org/css-logical/#propdef-margin-block-start){#margin-collapsing-quirks:'margin-block-start'
x-internal="'margin-block-start'"} property to zero.

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#margin-collapsing-quirks:quirks-mode-2
x-internal="quirks-mode"}, any [element with default
margins](#concept-rendering-elements-with-margins){#margin-collapsing-quirks:concept-rendering-elements-with-margins-2}
that is the
[child](https://dom.spec.whatwg.org/#concept-tree-child){#margin-collapsing-quirks:concept-tree-child-2
x-internal="concept-tree-child"} of a
[`body`{#margin-collapsing-quirks:the-body-element-2}](sections.html#the-body-element),
[`td`{#margin-collapsing-quirks:the-td-element-2}](tables.html#the-td-element),
or
[`th`{#margin-collapsing-quirks:the-th-element-2}](tables.html#the-th-element)
element, has no
[substantial](#concept-rendering-substantial){#margin-collapsing-quirks:concept-rendering-substantial-3}
previous siblings, and is
[blank](#concept-rendering-blank){#margin-collapsing-quirks:concept-rendering-blank},
is expected to have a user-agent level style sheet rule that sets its
[\'margin-block-end\'](https://drafts.csswg.org/css-logical/#propdef-margin-block-end){#margin-collapsing-quirks:'margin-block-end'
x-internal="'margin-block-end'"} property to zero also.

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#margin-collapsing-quirks:quirks-mode-3
x-internal="quirks-mode"}, any [element with default
margins](#concept-rendering-elements-with-margins){#margin-collapsing-quirks:concept-rendering-elements-with-margins-3}
that is the
[child](https://dom.spec.whatwg.org/#concept-tree-child){#margin-collapsing-quirks:concept-tree-child-3
x-internal="concept-tree-child"} of a
[`td`{#margin-collapsing-quirks:the-td-element-3}](tables.html#the-td-element)
or
[`th`{#margin-collapsing-quirks:the-th-element-3}](tables.html#the-th-element)
element, has no
[substantial](#concept-rendering-substantial){#margin-collapsing-quirks:concept-rendering-substantial-4}
following siblings, and is
[blank](#concept-rendering-blank){#margin-collapsing-quirks:concept-rendering-blank-2},
is expected to have a user-agent level style sheet rule that sets its
[\'margin-block-start\'](https://drafts.csswg.org/css-logical/#propdef-margin-block-start){#margin-collapsing-quirks:'margin-block-start'-2
x-internal="'margin-block-start'"} property to zero.

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#margin-collapsing-quirks:quirks-mode-4
x-internal="quirks-mode"}, any
[`p`{#margin-collapsing-quirks:the-p-element-2}](grouping-content.html#the-p-element)
element that is the
[child](https://dom.spec.whatwg.org/#concept-tree-child){#margin-collapsing-quirks:concept-tree-child-4
x-internal="concept-tree-child"} of a
[`td`{#margin-collapsing-quirks:the-td-element-4}](tables.html#the-td-element)
or
[`th`{#margin-collapsing-quirks:the-th-element-4}](tables.html#the-th-element)
element and has no
[substantial](#concept-rendering-substantial){#margin-collapsing-quirks:concept-rendering-substantial-5}
following siblings, is expected to have a user-agent level style sheet
rule that sets its
[\'margin-block-end\'](https://drafts.csswg.org/css-logical/#propdef-margin-block-end){#margin-collapsing-quirks:'margin-block-end'-2
x-internal="'margin-block-end'"} property to zero.

#### [15.3.10]{.secno} Form controls[](#form-controls){.self-link}

``` css
@namespace "http://www.w3.org/1999/xhtml";

input, select, button, textarea {
  letter-spacing: initial;
  word-spacing: initial;
  line-height: initial;
  text-transform: initial;
  text-indent: initial;
  text-shadow: initial;
  appearance: auto;
}

input:not([type=image i], [type=range i], [type=checkbox i], [type=radio i]) {
  overflow: clip !important;
  overflow-clip-margin: 0 !important;
}

input, select, textarea {
  text-align: initial;
}

:autofill {
  field-sizing: fixed !important;
}

input:is([type=reset i], [type=button i], [type=submit i]), button {
  text-align: center;
}

input, button {
  display: inline-block;
}

input[type=hidden i], input[type=file i], input[type=image i] {
  appearance: none;
}

input:is([type=radio i], [type=checkbox i], [type=reset i], [type=button i],
[type=submit i], [type=color i], [type=search i]), select, button {
  box-sizing: border-box;
}

textarea { white-space: pre-wrap; }
```

In [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#form-controls:quirks-mode
x-internal="quirks-mode"}, the following rules are also expected to
apply:

``` css
@namespace "http://www.w3.org/1999/xhtml";

input:not([type=image i]), textarea { box-sizing: border-box; }
```

Each kind of form control is also described in the [Widgets](#widgets)
section, which describes the look and feel of the control.

For
[`input`{#form-controls:the-input-element}](input.html#the-input-element)
elements where the
[`type`{#form-controls:attr-input-type}](input.html#attr-input-type)
attribute is not in the
[Hidden](input.html#hidden-state-(type=hidden)){#form-controls:hidden-state-(type=hidden)}
state or the [Image
Button](input.html#image-button-state-(type=image)){#form-controls:image-button-state-(type=image)}
state, and that are [being
rendered](#being-rendered){#form-controls:being-rendered}, are expected
to act as follows:

- The [inner display
  type](https://drafts.csswg.org/css-display/#inner-display-type){#form-controls:inner-display-type
  x-internal="inner-display-type"} is always \'flow-root\'.

#### [15.3.11]{.secno} The [`hr`{#the-hr-element-2:the-hr-element}](grouping-content.html#the-hr-element) element[](#the-hr-element-2){.self-link} {#the-hr-element-2}

``` css
@namespace "http://www.w3.org/1999/xhtml";

hr {
  color: gray;
  border-style: inset;
  border-width: 1px;
  margin-block: 0.5em;
  margin-inline: auto;
  overflow: hidden;
}
```

The following rules are also expected to apply, as [presentational
hints](#presentational-hints){#the-hr-element-2:presentational-hints}:

``` css
@namespace "http://www.w3.org/1999/xhtml";

hr[align=left i] { margin-left: 0; margin-right: auto; }
hr[align=right i] { margin-left: auto; margin-right: 0; }
hr[align=center i] { margin-left: auto; margin-right: auto; }
hr[color], hr[noshade] { border-style: solid; }
```

If an
[`hr`{#the-hr-element-2:the-hr-element-2}](grouping-content.html#the-hr-element)
element has either a
[`color`{#the-hr-element-2:attr-hr-color}](obsolete.html#attr-hr-color)
attribute or a
[`noshade`{#the-hr-element-2:attr-hr-noshade}](obsolete.html#attr-hr-noshade)
attribute, and furthermore also has a
[`size`{#the-hr-element-2:attr-hr-size}](obsolete.html#attr-hr-size)
attribute, and parsing that attribute\'s value using the [rules for
parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-hr-element-2:rules-for-parsing-non-negative-integers}
doesn\'t generate an error, then the user agent is expected to use the
parsed value divided by two as a pixel length for [presentational
hints](#presentational-hints){#the-hr-element-2:presentational-hints-2}
for the properties
[\'border-top-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-width){#the-hr-element-2:'border-top-width'
x-internal="'border-top-width'"},
[\'border-right-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-width){#the-hr-element-2:'border-right-width'
x-internal="'border-right-width'"},
[\'border-bottom-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-width){#the-hr-element-2:'border-bottom-width'
x-internal="'border-bottom-width'"}, and
[\'border-left-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-width){#the-hr-element-2:'border-left-width'
x-internal="'border-left-width'"} on the element.

Otherwise, if an
[`hr`{#the-hr-element-2:the-hr-element-3}](grouping-content.html#the-hr-element)
element has neither a
[`color`{#the-hr-element-2:attr-hr-color-2}](obsolete.html#attr-hr-color)
attribute nor a
[`noshade`{#the-hr-element-2:attr-hr-noshade-2}](obsolete.html#attr-hr-noshade)
attribute, but does have a
[`size`{#the-hr-element-2:attr-hr-size-2}](obsolete.html#attr-hr-size)
attribute, and parsing that attribute\'s value using the [rules for
parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-hr-element-2:rules-for-parsing-non-negative-integers-2}
doesn\'t generate an error, then: if the parsed value is one, then the
user agent is expected to use the attribute as a [presentational
hint](#presentational-hints){#the-hr-element-2:presentational-hints-3}
setting the element\'s
[\'border-bottom-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-width){#the-hr-element-2:'border-bottom-width'-2
x-internal="'border-bottom-width'"} to 0; otherwise, if the parsed value
is greater than one, then the user agent is expected to use the parsed
value minus two as a pixel length for [presentational
hints](#presentational-hints){#the-hr-element-2:presentational-hints-4}
for the
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#the-hr-element-2:'height'
x-internal="'height'"} property on the element.

The
[`width`{#the-hr-element-2:attr-hr-width}](obsolete.html#attr-hr-width)
attribute on an
[`hr`{#the-hr-element-2:the-hr-element-4}](grouping-content.html#the-hr-element)
element [maps to the dimension
property](#maps-to-the-dimension-property){#the-hr-element-2:maps-to-the-dimension-property}
[\'width\'](https://drafts.csswg.org/css2/#the-width-property){#the-hr-element-2:'width'
x-internal="'width'"} on the element.

When an
[`hr`{#the-hr-element-2:the-hr-element-5}](grouping-content.html#the-hr-element)
element has a
[`color`{#the-hr-element-2:attr-hr-color-3}](obsolete.html#attr-hr-color)
attribute, its value is expected to be parsed using the [rules for
parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#the-hr-element-2:rules-for-parsing-a-legacy-colour-value},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#the-hr-element-2:presentational-hints-5}
setting the element\'s
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#the-hr-element-2:'color'
x-internal="'color'"} property to the resulting color.

#### [15.3.12]{.secno} The [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element}](form-elements.html#the-fieldset-element) and [`legend`{#the-fieldset-and-legend-elements:the-legend-element}](form-elements.html#the-legend-element) elements[](#the-fieldset-and-legend-elements){.self-link}

``` css
@namespace "http://www.w3.org/1999/xhtml";

fieldset {
  display: block;
  margin-inline: 2px;
  border: groove 2px ThreeDFace;
  padding-block: 0.35em 0.625em;
  padding-inline: 0.75em;
  min-inline-size: min-content;
}

legend {
  padding-inline: 2px;
}

legend[align=left i] {
  justify-self: left;
}

legend[align=center i] {
  justify-self: center;
}

legend[align=right i] {
  justify-self: right;
}
```

The
[`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-2}](form-elements.html#the-fieldset-element)
element, when it generates a [CSS
box](https://drafts.csswg.org/css-display/#css-box){#the-fieldset-and-legend-elements:css-box
x-internal="css-box"}, is expected to act as follows:

- The element is expected to establish a new [block formatting
  context](https://drafts.csswg.org/css-display/#block-formatting-context){#the-fieldset-and-legend-elements:block-formatting-context
  x-internal="block-formatting-context"}.

- The
  [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-fieldset-and-legend-elements:'display'
  x-internal="'display'"} property is expected to act as follows:

  - If the computed value of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-fieldset-and-legend-elements:'display'-2
    x-internal="'display'"} is a value such that the [outer display
    type](https://drafts.csswg.org/css-display/#outer-display-type){#the-fieldset-and-legend-elements:outer-display-type
    x-internal="outer-display-type"} is \'inline\', then behave as
    \'inline-block\'.

  - Otherwise, behave as \'flow-root\'.

  This does not change the computed value.

- If the element\'s box has a child box that matches the conditions in
  the list below, then the first such child box is the \'fieldset\'
  element\'s [rendered legend]{#rendered-legend .dfn}:

  - The child is a
    [`legend`{#the-fieldset-and-legend-elements:the-legend-element-2}](form-elements.html#the-legend-element)
    element.
  - The child\'s used value of
    [\'float\'](https://drafts.csswg.org/css2/#float-position){#the-fieldset-and-legend-elements:'float'
    x-internal="'float'"} is \'none\'.
  - The child\'s used value of
    [\'position\'](https://drafts.csswg.org/css-position/#position-property){#the-fieldset-and-legend-elements:'position'
    x-internal="'position'"} is not \'absolute\' or \'fixed\'.

- If the element has a [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend},
  then the border is expected to not be painted behind the rectangle
  defined as follows, using the writing mode of the fieldset:

  1.  The block-start edge of the rectangle is the smaller of the
      block-start edge of the [rendered
      legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-2}\'s
      margin rectangle at its static position (ignoring transforms), and
      the block-start outer edge of the
      [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-3}](form-elements.html#the-fieldset-element)\'s
      border.

  2.  The block-end edge of the rectangle is the larger of the block-end
      edge of the [rendered
      legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-3}\'s
      margin rectangle at its static position (ignoring transforms), and
      the block-end outer edge of the
      [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-4}](form-elements.html#the-fieldset-element)\'s
      border.

  3.  The inline-start edge of the rectangle is the smaller of the
      inline-start edge of the [rendered
      legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-4}\'s
      border rectangle at its static position (ignoring transforms), and
      the inline-start outer edge of the
      [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-5}](form-elements.html#the-fieldset-element)\'s
      border.

  4.  The inline-end edge of the rectangle is the larger of the
      inline-end edge of the [rendered
      legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-5}\'s
      border rectangle at its static position (ignoring transforms), and
      the inline-end outer edge of the
      [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-6}](form-elements.html#the-fieldset-element)\'s
      border.

- The space allocated for the element\'s border on the block-start side
  is expected to be the element\'s
  [\'border-block-start-width\'](https://drafts.csswg.org/css-logical/#propdef-border-block-start-width){#the-fieldset-and-legend-elements:'border-block-start-width'
  x-internal="'border-block-start-width'"} or the [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-6}\'s
  margin box size in the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-7}](form-elements.html#the-fieldset-element)\'s
  block-flow direction, whichever is greater.

- For the purpose of calculating the used
  [\'block-size\'](https://drafts.csswg.org/css-logical/#propdef-block-size){#the-fieldset-and-legend-elements:'block-size'
  x-internal="'block-size'"}, if the computed
  [\'block-size\'](https://drafts.csswg.org/css-logical/#propdef-block-size){#the-fieldset-and-legend-elements:'block-size'-2
  x-internal="'block-size'"} is not \'auto\', the space allocated for
  the [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-7}\'s
  margin box that spills out past the border, if any, is expected to be
  subtracted from the
  [\'block-size\'](https://drafts.csswg.org/css-logical/#propdef-block-size){#the-fieldset-and-legend-elements:'block-size'-3
  x-internal="'block-size'"}. If the content box\'s block-size would be
  negative, then let the content box\'s block-size be zero instead.

- If the element has a [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-8},
  then that element is expected to be the first child box.

- The [anonymous fieldset content
  box](#anonymous-fieldset-content-box){#the-fieldset-and-legend-elements:anonymous-fieldset-content-box}
  is expected to appear after the [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-9}
  and is expected to contain the content (including the \'::before\' and
  \'::after\' pseudo-elements) of the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-8}](form-elements.html#the-fieldset-element)
  element except for the [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-10},
  if there is one.

- The used value of the
  [\'padding-top\'](https://drafts.csswg.org/css-box/#propdef-padding-top){#the-fieldset-and-legend-elements:'padding-top'
  x-internal="'padding-top'"},
  [\'padding-right\'](https://drafts.csswg.org/css-box/#propdef-padding-right){#the-fieldset-and-legend-elements:'padding-right'
  x-internal="'padding-right'"},
  [\'padding-bottom\'](https://drafts.csswg.org/css-box/#propdef-padding-bottom){#the-fieldset-and-legend-elements:'padding-bottom'
  x-internal="'padding-bottom'"}, and
  [\'padding-left\'](https://drafts.csswg.org/css-box/#propdef-padding-left){#the-fieldset-and-legend-elements:'padding-left'
  x-internal="'padding-left'"} properties are expected to be zero.

- For the purpose of calculating the min-content inline size, use the
  greater of the min-content inline size of the [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-11}
  and the min-content inline size of the [anonymous fieldset content
  box](#anonymous-fieldset-content-box){#the-fieldset-and-legend-elements:anonymous-fieldset-content-box-2}.

- For the purpose of calculating the max-content inline size, use the
  greater of the max-content inline size of the [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-12}
  and the max-content inline size of the [anonymous fieldset content
  box](#anonymous-fieldset-content-box){#the-fieldset-and-legend-elements:anonymous-fieldset-content-box-3}.

A
[`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-9}](form-elements.html#the-fieldset-element)
element\'s [rendered
legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-13},
if any, is expected to act as follows:

- The element is expected to establish a new [formatting
  context](https://drafts.csswg.org/css-display/#formatting-context){#the-fieldset-and-legend-elements:formatting-context
  x-internal="formatting-context"} for its contents. The type of this
  [formatting
  context](https://drafts.csswg.org/css-display/#formatting-context){#the-fieldset-and-legend-elements:formatting-context-2
  x-internal="formatting-context"} is determined by its
  [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-fieldset-and-legend-elements:'display'-3
  x-internal="'display'"} value, as usual.

- The
  [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-fieldset-and-legend-elements:'display'-4
  x-internal="'display'"} property is expected to behave as if its
  computed value was blockified.

  This does not change the computed value.

- If the [computed
  value](https://drafts.csswg.org/css-cascade/#computed-value){#the-fieldset-and-legend-elements:computed-value
  x-internal="computed-value"} of
  [\'inline-size\'](https://drafts.csswg.org/css-logical/#propdef-inline-size){#the-fieldset-and-legend-elements:'inline-size'
  x-internal="'inline-size'"} is \'auto\', then the [used
  value](https://drafts.csswg.org/css-cascade/#used-value){#the-fieldset-and-legend-elements:used-value
  x-internal="used-value"} is the [fit-content inline
  size](https://drafts.csswg.org/css-sizing/#fit-content-inline-size){#the-fieldset-and-legend-elements:fit-content-inline-size
  x-internal="fit-content-inline-size"}.

- The element is expected to be positioned in the inline direction as is
  normal for blocks (e.g., taking into account margins and the
  [\'justify-self\'](https://drafts.csswg.org/css-align/#propdef-justify-self){#the-fieldset-and-legend-elements:'justify-self'
  x-internal="'justify-self'"} property).

- The element\'s box is expected to be constrained in the inline
  direction by the inline content size of the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-10}](form-elements.html#the-fieldset-element)
  as if it had used its computed inline padding.

  ::: example
  For example, if the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-11}](form-elements.html#the-fieldset-element)
  has a specified padding of 50px, then the [rendered
  legend](#rendered-legend){#the-fieldset-and-legend-elements:rendered-legend-14}
  will be positioned 50px in from the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-12}](form-elements.html#the-fieldset-element)\'s
  border. The padding will further apply to the [anonymous fieldset
  content
  box](#anonymous-fieldset-content-box){#the-fieldset-and-legend-elements:anonymous-fieldset-content-box-4}
  instead of the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-13}](form-elements.html#the-fieldset-element)
  element itself.
  :::

- The element is expected to be positioned in the block-flow direction
  such that its border box is centered over the border on the
  block-start side of the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-14}](form-elements.html#the-fieldset-element)
  element.

A
[`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-15}](form-elements.html#the-fieldset-element)
element\'s [anonymous fieldset content
box]{#anonymous-fieldset-content-box .dfn} is expected to act as
follows:

- The
  [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-fieldset-and-legend-elements:'display'-5
  x-internal="'display'"} property is expected to act as follows:

  - If the computed value of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-fieldset-and-legend-elements:'display'-6
    x-internal="'display'"} on the
    [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-16}](form-elements.html#the-fieldset-element)
    element is \'grid\' or \'inline-grid\', then set the used value to
    \'grid\'.

  - If the computed value of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#the-fieldset-and-legend-elements:'display'-7
    x-internal="'display'"} on the
    [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-17}](form-elements.html#the-fieldset-element)
    element is \'flex\' or \'inline-flex\', then set the used value to
    \'flex\'.

  - Otherwise, set the used value to \'flow-root\'.

- The following properties are expected to inherit from the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-18}](form-elements.html#the-fieldset-element)
  element:

  - [\'align-content\'](https://drafts.csswg.org/css-align/#propdef-align-content){#the-fieldset-and-legend-elements:'align-content'
    x-internal="'align-content'"}
  - [\'align-items\'](https://drafts.csswg.org/css-align/#propdef-align-items){#the-fieldset-and-legend-elements:'align-items'
    x-internal="'align-items'"}
  - [\'border-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-radius){#the-fieldset-and-legend-elements:'border-radius'
    x-internal="'border-radius'"}
  - [\'column-count\'](https://drafts.csswg.org/css-multicol/#propdef-column-count){#the-fieldset-and-legend-elements:'column-count'
    x-internal="'column-count'"}
  - [\'column-fill\'](https://drafts.csswg.org/css-multicol/#propdef-column-fill){#the-fieldset-and-legend-elements:'column-fill'
    x-internal="'column-fill'"}
  - [\'column-gap\'](https://drafts.csswg.org/css-multicol/#propdef-column-gap){#the-fieldset-and-legend-elements:'column-gap'
    x-internal="'column-gap'"}
  - [\'column-rule\'](https://drafts.csswg.org/css-multicol/#propdef-column-rule){#the-fieldset-and-legend-elements:'column-rule'
    x-internal="'column-rule'"}
  - [\'column-width\'](https://drafts.csswg.org/css-multicol/#propdef-column-width){#the-fieldset-and-legend-elements:'column-width'
    x-internal="'column-width'"}
  - [\'flex-direction\'](https://drafts.csswg.org/css-flexbox/#propdef-flex-direction){#the-fieldset-and-legend-elements:'flex-direction'
    x-internal="'flex-direction'"}
  - [\'flex-wrap\'](https://drafts.csswg.org/css-flexbox/#propdef-flex-wrap){#the-fieldset-and-legend-elements:'flex-wrap'
    x-internal="'flex-wrap'"}
  - [\'grid-auto-columns\'](https://drafts.csswg.org/css-grid/#propdef-grid-auto-columns){#the-fieldset-and-legend-elements:'grid-auto-columns'
    x-internal="'grid-auto-columns'"}
  - [\'grid-auto-flow\'](https://drafts.csswg.org/css-grid/#propdef-grid-auto-flow){#the-fieldset-and-legend-elements:'grid-auto-flow'
    x-internal="'grid-auto-flow'"}
  - [\'grid-auto-rows\'](https://drafts.csswg.org/css-grid/#propdef-grid-auto-rows){#the-fieldset-and-legend-elements:'grid-auto-rows'
    x-internal="'grid-auto-rows'"}
  - [\'grid-column-gap\'](https://drafts.csswg.org/css-grid/#propdef-grid-column-gap){#the-fieldset-and-legend-elements:'grid-column-gap'
    x-internal="'grid-column-gap'"}
  - [\'grid-row-gap\'](https://drafts.csswg.org/css-grid/#propdef-grid-row-gap){#the-fieldset-and-legend-elements:'grid-row-gap'
    x-internal="'grid-row-gap'"}
  - [\'grid-template-areas\'](https://drafts.csswg.org/css-grid/#propdef-grid-template-areas){#the-fieldset-and-legend-elements:'grid-template-areas'
    x-internal="'grid-template-areas'"}
  - [\'grid-template-columns\'](https://drafts.csswg.org/css-grid/#propdef-grid-template-columns){#the-fieldset-and-legend-elements:'grid-template-columns'
    x-internal="'grid-template-columns'"}
  - [\'grid-template-rows\'](https://drafts.csswg.org/css-grid/#propdef-grid-template-rows){#the-fieldset-and-legend-elements:'grid-template-rows'
    x-internal="'grid-template-rows'"}
  - [\'justify-content\'](https://drafts.csswg.org/css-align/#propdef-propdef-justify-content){#the-fieldset-and-legend-elements:'justify-content'
    x-internal="'justify-content'"}
  - [\'justify-items\'](https://drafts.csswg.org/css-align/#propdef-propdef-justify-items){#the-fieldset-and-legend-elements:'justify-items'
    x-internal="'justify-items'"}
  - [\'overflow\'](https://drafts.csswg.org/css-overflow/#propdef-overflow){#the-fieldset-and-legend-elements:'overflow'
    x-internal="'overflow'"}
  - [\'padding-bottom\'](https://drafts.csswg.org/css-box/#propdef-padding-bottom){#the-fieldset-and-legend-elements:'padding-bottom'-2
    x-internal="'padding-bottom'"}
  - [\'padding-left\'](https://drafts.csswg.org/css-box/#propdef-padding-left){#the-fieldset-and-legend-elements:'padding-left'-2
    x-internal="'padding-left'"}
  - [\'padding-right\'](https://drafts.csswg.org/css-box/#propdef-padding-right){#the-fieldset-and-legend-elements:'padding-right'-2
    x-internal="'padding-right'"}
  - [\'padding-top\'](https://drafts.csswg.org/css-box/#propdef-padding-top){#the-fieldset-and-legend-elements:'padding-top'-2
    x-internal="'padding-top'"}
  - [\'text-overflow\'](https://drafts.csswg.org/css-overflow/#propdef-text-overflow){#the-fieldset-and-legend-elements:'text-overflow'
    x-internal="'text-overflow'"}
  - [\'unicode-bidi\'](https://drafts.csswg.org/css-writing-modes/#unicode-bidi){#the-fieldset-and-legend-elements:'unicode-bidi'
    x-internal="'unicode-bidi'"}

- The
  [\'block-size\'](https://drafts.csswg.org/css-logical/#propdef-block-size){#the-fieldset-and-legend-elements:'block-size'-4
  x-internal="'block-size'"} property is expected to be set to \'100%\'.

- For the purpose of calculating percentage padding, act as if the
  padding was calculated for the
  [`fieldset`{#the-fieldset-and-legend-elements:the-fieldset-element-19}](form-elements.html#the-fieldset-element)
  element.

::: {#fieldset-layout-model .note}
![The legend is rendered over the top border, and the top border area
reserves vertical space for the legend. The fieldset\'s top margin
starts at the top margin edge of the legend. The legend\'s horizontal
margins, or the
[\'justify-self\'](https://drafts.csswg.org/css-align/#propdef-justify-self){#the-fieldset-and-legend-elements:'justify-self'-2
x-internal="'justify-self'"} property, gives its horizontal position.
The [anonymous fieldset content
box](#anonymous-fieldset-content-box){#the-fieldset-and-legend-elements:anonymous-fieldset-content-box-5}
appears below the
legend.](data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDgwIiByb2xlPSJpbWciIHZpZXdib3g9IjAgMCA0MDAgMjcwIiBoZWlnaHQ9IjMyNCIgYXJpYS1sYWJlbD0iRmllbGRzZXQgbGF5b3V0CiAgICBtb2RlbCI+CgogICAgIDxkZWZzPgogICAgICA8bWFya2VyIG1hcmtlcmhlaWdodD0iNCIgaWQ9ImFycm93LXJlZCIgb3JpZW50PSJhdXRvIiBtYXJrZXJ3aWR0aD0iNSIgcmVmeD0iMC4xIiByZWZ5PSIyIj4KICAgICAgIDxwYXRoIGZpbGw9InJlZCIgZD0iTTAsMCBWNCBMNSwyIFoiIC8+CiAgICAgIDwvbWFya2VyPgogICAgICA8bWFya2VyIG1hcmtlcmhlaWdodD0iNCIgaWQ9ImFycm93LWJsdWUiIG9yaWVudD0iYXV0byIgbWFya2Vyd2lkdGg9IjUiIHJlZng9IjAuMSIgcmVmeT0iMiI+CiAgICAgICA8cGF0aCBmaWxsPSJibHVlIiBkPSJNMCwwIFY0IEw1LDIgWiIgLz4KICAgICAgPC9tYXJrZXI+CiAgICAgPC9kZWZzPgoKICAgICAKICAgICA8cmVjdCB3aWR0aD0iMzA0IiBzdHJva2UtZGFzaGFycmF5PSI2IiBzdHJva2Utd2lkdGg9IjEiIGZpbGw9Im5vbmUiIHg9IjQ4IiBzdHJva2U9ImJsYWNrIiB5PSI0OCIgaGVpZ2h0PSIxNzYiIC8+CgogICAgIDxsaW5lIHkyPSIyMCIgc3Ryb2tlLXdpZHRoPSIyIiBtYXJrZXItZW5kPSJ1cmwoI2Fycm93LWJsdWUpIiBzdHJva2U9ImJsdWUiIHgxPSIyMDAiIHgyPSIyMDAiIHkxPSI0OCI+PC9saW5lPgogICAgIDxsaW5lIHkyPSIxMzYiIHN0cm9rZS13aWR0aD0iMiIgbWFya2VyLWVuZD0idXJsKCNhcnJvdy1ibHVlKSIgc3Ryb2tlPSJibHVlIiB4MT0iNDgiIHgyPSIyMCIgeTE9IjEzNiI+PC9saW5lPgogICAgIDxsaW5lIHkyPSIxMzYiIHN0cm9rZS13aWR0aD0iMiIgbWFya2VyLWVuZD0idXJsKCNhcnJvdy1ibHVlKSIgc3Ryb2tlPSJibHVlIiB4MT0iMzUyIiB4Mj0iMzgwIiB5MT0iMTM2Ij48L2xpbmU+CiAgICAgPGxpbmUgeTI9IjI1MCIgc3Ryb2tlLXdpZHRoPSIyIiBtYXJrZXItZW5kPSJ1cmwoI2Fycm93LWJsdWUpIiBzdHJva2U9ImJsdWUiIHgxPSIyMDAiIHgyPSIyMDAiIHkxPSIyMjQiPjwvbGluZT4KICAgICA8dGV4dCBmb250LXNpemU9IjEyIiBmb250LXN0eWxlPSJub3JtYWwiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBmaWxsPSJibHVlIiB4PSIyMTAiIHk9IjM1Ij5maWVsZHNldCYjMzk7cyBtYXJnaW48L3RleHQ+CgogICAgIAogICAgIDxyZWN0IHdpZHRoPSIzMDAiIHN0cm9rZS13aWR0aD0iMiIgZmlsbD0id2hpdGUiIHN0cm9rZT0iYmx1ZSIgeD0iNTAiIHk9IjEwNSIgaGVpZ2h0PSIxMTYuNSIgLz4KCiAgICAgCiAgICAgPHJlY3Qgd2lkdGg9Ijg1IiBzdHJva2Utd2lkdGg9IjIiIGZpbGw9IndoaXRlIiBzdHJva2U9InJlZCIgeD0iMTAwIiB5PSI3NSIgaGVpZ2h0PSI2MCIgLz4KICAgICA8dGV4dCBmb250LXN0eWxlPSJub3JtYWwiIGZvbnQtZmFtaWx5PSJzZXJpZiIgZmlsbD0icmVkIiB4PSIxMjAiIHk9IjEwNSIgZG9taW5hbnQtYmFzZWxpbmU9ImNlbnRyYWwiPmxlZ2VuZDwvdGV4dD4KCiAgICAgCiAgICAgPHJlY3Qgd2lkdGg9IjY1IiBzdHJva2Utd2lkdGg9IjE4IiBmaWxsPSJub25lIiBzdHJva2U9IiNlZWUiIHg9IjExMCIgeT0iODUiIGhlaWdodD0iNDAiIC8+CiAgICAgPHRleHQgZm9udC1zaXplPSIxMiIgZm9udC1zdHlsZT0ibm9ybWFsIiBmb250LWZhbWlseT0ic2Fucy1zZXJpZiIgZmlsbD0iZ3JheSIgeD0iMTMzIiB5PSI4OCI+cGFkZGluZzwvdGV4dD4KCiAgICAgPGxpbmUgeTI9IjkwIiBzdHJva2Utd2lkdGg9IjIiIG1hcmtlci1lbmQ9InVybCgjYXJyb3ctcmVkKSIgc3Ryb2tlPSJyZWQiIHgxPSIxODUiIHgyPSIzMjUiIHkxPSI5MCI+PC9saW5lPgogICAgIDx0ZXh0IGZvbnQtc2l6ZT0iMTIiIGZvbnQtc3R5bGU9Im5vcm1hbCIgZm9udC1mYW1pbHk9InNhbnMtc2VyaWYiIGZpbGw9InJlZCIgeD0iMjEwIiB5PSI4NSI+bGVnZW5kJiMzOTtzIG1hcmdpbjwvdGV4dD4KCiAgICAgPGxpbmUgeTI9IjkwIiBzdHJva2Utd2lkdGg9IjIiIG1hcmtlci1lbmQ9InVybCgjYXJyb3ctcmVkKSIgc3Ryb2tlPSJyZWQiIHgxPSIxMDAiIHgyPSI3NSIgeTE9IjkwIj48L2xpbmU+CiAgICAgPGxpbmUgeTI9IjYwIiBzdHJva2Utd2lkdGg9IjIiIG1hcmtlci1lbmQ9InVybCgjYXJyb3ctcmVkKSIgc3Ryb2tlPSJyZWQiIHgxPSIxNDMiIHgyPSIxNDMiIHkxPSI3NSI+PC9saW5lPgogICAgIDxsaW5lIHkyPSIxNTAiIHN0cm9rZS13aWR0aD0iMiIgbWFya2VyLWVuZD0idXJsKCNhcnJvdy1yZWQpIiBzdHJva2U9InJlZCIgeDE9IjE0MyIgeDI9IjE0MyIgeTE9IjEzNSI+PC9saW5lPgoKICAgICAKICAgICA8cmVjdCB3aWR0aD0iMjgwIiBzdHJva2Utd2lkdGg9IjE4IiBmaWxsPSJub25lIiBzdHJva2U9IiNlZWUiIHg9IjYwIiB5PSIxNzEiIGhlaWdodD0iNDAiIC8+CiAgICAgPHRleHQgZm9udC1zaXplPSIxMiIgZm9udC1zdHlsZT0ibm9ybWFsIiBmb250LWZhbWlseT0ic2Fucy1zZXJpZiIgZmlsbD0iZ3JheSIgeD0iMjk3IiB5PSIxNzQiPnBhZGRpbmc8L3RleHQ+CgogICAgIAogICAgIDxyZWN0IHdpZHRoPSIyOTYiIHN0cm9rZS1kYXNoYXJyYXk9IjYiIHN0cm9rZS13aWR0aD0iMSIgZmlsbD0ibm9uZSIgeD0iNTIiIHN0cm9rZT0iYmxhY2siIHk9IjE2MSIgaGVpZ2h0PSI1OCIgLz4KICAgICA8dGV4dCBmb250LXN0eWxlPSJub3JtYWwiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBmb250LXNpemU9IjEyIiB4PSIxNzAiIHk9IjE1NiI+YW5vbnltb3VzCiAgICAgZmllbGRzZXQgY29udGVudCBib3g8L3RleHQ+CiAgICAgPHRleHQgZm9udC1zdHlsZT0ibm9ybWFsIiBmb250LWZhbWlseT0ic2VyaWYiIGZpbGw9ImJsdWUiIHg9IjcwIiB5PSIxOTEiIGRvbWluYW50LWJhc2VsaW5lPSJjZW50cmFsIj5jb250ZW50PC90ZXh0PgoKICAgIDwvc3ZnPg==)
:::

### [15.4]{.secno} Replaced elements[](#replaced-elements){.self-link}

The following elements can be [replaced
elements](https://drafts.csswg.org/css-display/#replaced-element){#replaced-elements:replaced-element
x-internal="replaced-element"}:
[`audio`{#replaced-elements:the-audio-element}](media.html#the-audio-element),
[`canvas`{#replaced-elements:the-canvas-element}](canvas.html#the-canvas-element),
[`embed`{#replaced-elements:the-embed-element}](iframe-embed-object.html#the-embed-element),
[`iframe`{#replaced-elements:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
[`img`{#replaced-elements:the-img-element}](embedded-content.html#the-img-element),
[`input`{#replaced-elements:the-input-element}](input.html#the-input-element),
[`object`{#replaced-elements:the-object-element}](iframe-embed-object.html#the-object-element),
and
[`video`{#replaced-elements:the-video-element}](media.html#the-video-element).

#### [15.4.1]{.secno} Embedded content[](#embedded-content-rendering-rules){.self-link} {#embedded-content-rendering-rules}

The
[`embed`{#embedded-content-rendering-rules:the-embed-element}](iframe-embed-object.html#the-embed-element),
[`iframe`{#embedded-content-rendering-rules:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
and
[`video`{#embedded-content-rendering-rules:the-video-element}](media.html#the-video-element)
elements are expected to be treated as [replaced
elements](https://drafts.csswg.org/css-display/#replaced-element){#embedded-content-rendering-rules:replaced-element
x-internal="replaced-element"}.

A
[`canvas`{#embedded-content-rendering-rules:the-canvas-element}](canvas.html#the-canvas-element)
element that
[represents](dom.html#represents){#embedded-content-rendering-rules:represents}
[embedded
content](dom.html#embedded-content-category){#embedded-content-rendering-rules:embedded-content-category}
is expected to be treated as a [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#embedded-content-rendering-rules:replaced-element-2
x-internal="replaced-element"}; the contents of such elements are the
element\'s bitmap, if any, or else a [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#embedded-content-rendering-rules:transparent-black
x-internal="transparent-black"} bitmap with the same [natural
dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#embedded-content-rendering-rules:natural-dimensions
x-internal="natural-dimensions"} as the element. Other
[`canvas`{#embedded-content-rendering-rules:the-canvas-element-2}](canvas.html#the-canvas-element)
elements are expected to be treated as ordinary elements in the
rendering model.

An
[`object`{#embedded-content-rendering-rules:the-object-element}](iframe-embed-object.html#the-object-element)
element that
[represents](dom.html#represents){#embedded-content-rendering-rules:represents-2}
an image, plugin, or its [content
navigable](document-sequences.html#content-navigable){#embedded-content-rendering-rules:content-navigable}
is expected to be treated as a [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#embedded-content-rendering-rules:replaced-element-3
x-internal="replaced-element"}. Other
[`object`{#embedded-content-rendering-rules:the-object-element-2}](iframe-embed-object.html#the-object-element)
elements are expected to be treated as ordinary elements in the
rendering model.

The
[`audio`{#embedded-content-rendering-rules:the-audio-element}](media.html#the-audio-element)
element, when it is [exposing a user
interface](media.html#expose-a-user-interface-to-the-user){#embedded-content-rendering-rules:expose-a-user-interface-to-the-user},
is expected to be treated as a [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#embedded-content-rendering-rules:replaced-element-4
x-internal="replaced-element"} about one line high, as wide as is
necessary to expose the user agent\'s user interface features. When an
[`audio`{#embedded-content-rendering-rules:the-audio-element-2}](media.html#the-audio-element)
element is not [exposing a user
interface](media.html#expose-a-user-interface-to-the-user){#embedded-content-rendering-rules:expose-a-user-interface-to-the-user-2},
the user agent is expected to force its
[\'display\'](https://drafts.csswg.org/css2/#display-prop){#embedded-content-rendering-rules:'display'
x-internal="'display'"} property to compute to \'none\', irrespective of
CSS rules.

Whether a
[`video`{#embedded-content-rendering-rules:the-video-element-2}](media.html#the-video-element)
element is [exposing a user
interface](media.html#expose-a-user-interface-to-the-user){#embedded-content-rendering-rules:expose-a-user-interface-to-the-user-3}
is not expected to affect the size of the rendering; controls are
expected to be overlaid above the page content without causing any
layout changes, and are expected to disappear when the user does not
need them.

When a
[`video`{#embedded-content-rendering-rules:the-video-element-3}](media.html#the-video-element)
element represents a poster frame or frame of video, the poster frame or
frame of video is expected to be rendered at the largest size that
maintains the aspect ratio of that poster frame or frame of video
without being taller or wider than the
[`video`{#embedded-content-rendering-rules:the-video-element-4}](media.html#the-video-element)
element itself, and is expected to be centered in the
[`video`{#embedded-content-rendering-rules:the-video-element-5}](media.html#the-video-element)
element.

Any subtitles or captions are expected to be overlaid directly on top of
their
[`video`{#embedded-content-rendering-rules:the-video-element-6}](media.html#the-video-element)
element, as defined by the relevant rendering rules; for WebVTT, those
are the [rules for updating the display of WebVTT text
tracks](https://w3c.github.io/webvtt/#rules-for-updating-the-display-of-webvtt-text-tracks){#embedded-content-rendering-rules:rules-for-updating-the-display-of-webvtt-text-tracks
x-internal="rules-for-updating-the-display-of-webvtt-text-tracks"}.
[\[WEBVTT\]](references.html#refsWEBVTT)

When the user agent starts [exposing a user
interface](media.html#expose-a-user-interface-to-the-user){#embedded-content-rendering-rules:expose-a-user-interface-to-the-user-4}
for a
[`video`{#embedded-content-rendering-rules:the-video-element-7}](media.html#the-video-element)
element, the user agent should run the [rules for updating the text
track
rendering](media.html#rules-for-updating-the-text-track-rendering){#embedded-content-rendering-rules:rules-for-updating-the-text-track-rendering}
of each of the [text
tracks](media.html#text-track){#embedded-content-rendering-rules:text-track}
in the
[`video`{#embedded-content-rendering-rules:the-video-element-8}](media.html#the-video-element)
element\'s [list of text
tracks](media.html#list-of-text-tracks){#embedded-content-rendering-rules:list-of-text-tracks}
that are
[showing](media.html#text-track-showing){#embedded-content-rendering-rules:text-track-showing}
and whose [text track
kind](media.html#text-track-kind){#embedded-content-rendering-rules:text-track-kind}
is one of
[`subtitles`{#embedded-content-rendering-rules:dom-texttrack-kind-subtitles}](media.html#dom-texttrack-kind-subtitles)
or
[`captions`{#embedded-content-rendering-rules:dom-texttrack-kind-captions}](media.html#dom-texttrack-kind-captions)
(e.g., for [text
tracks](media.html#text-track){#embedded-content-rendering-rules:text-track-2}
based on WebVTT, the [rules for updating the display of WebVTT text
tracks](https://w3c.github.io/webvtt/#rules-for-updating-the-display-of-webvtt-text-tracks){#embedded-content-rendering-rules:rules-for-updating-the-display-of-webvtt-text-tracks-2
x-internal="rules-for-updating-the-display-of-webvtt-text-tracks"}).
[\[WEBVTT\]](references.html#refsWEBVTT)

Resizing
[`video`{#embedded-content-rendering-rules:the-video-element-9}](media.html#the-video-element)
and
[`canvas`{#embedded-content-rendering-rules:the-canvas-element-3}](canvas.html#the-canvas-element)
elements does not interrupt video playback or clear the canvas.

------------------------------------------------------------------------

The following CSS rules are expected to apply:

``` css
@namespace "http://www.w3.org/1999/xhtml";

iframe { border: 2px inset; }
video { object-fit: contain; }
```

#### [15.4.2]{.secno} Images[](#images-3){.self-link} {#images-3}

User agents are expected to render
[`img`{#images-3:the-img-element}](embedded-content.html#the-img-element)
elements and
[`input`{#images-3:the-input-element}](input.html#the-input-element)
elements whose
[`type`{#images-3:attr-input-type}](input.html#attr-input-type)
attributes are in the [Image
Button](input.html#image-button-state-(type=image)){#images-3:image-button-state-(type=image)}
state, according to the first applicable rules from the following list:

If the element [represents](dom.html#represents){#images-3:represents}
an image

The user agent is expected to treat the element as a [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#images-3:replaced-element
x-internal="replaced-element"} and render the image according to the
rules for doing so defined in CSS.

If the element does not
[represent](dom.html#represents){#images-3:represents-2} an image and
either:

- the user agent has reason to believe that the image will become
  *[available](images.html#img-available)* and be rendered in due
  course, or
- the element has no `alt` attribute, or
- the [`Document`{#images-3:document}](dom.html#document) is in [quirks
  mode](https://dom.spec.whatwg.org/#concept-document-quirks){#images-3:quirks-mode
  x-internal="quirks-mode"}, and the element already has [natural
  dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#images-3:natural-dimensions
  x-internal="natural-dimensions"} (e.g., from the [dimension
  attributes](embedded-content-other.html#dimension-attributes){#images-3:dimension-attributes}
  or CSS rules)

The user agent is expected to treat the element as a [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#images-3:replaced-element-2
x-internal="replaced-element"} whose content is the text that the
element represents, if any, optionally alongside an icon indicating that
the image is being obtained (if applicable). For
[`input`{#images-3:the-input-element-2}](input.html#the-input-element)
elements, the element is expected to appear button-like to indicate that
the element is a
[button](forms.html#concept-button){#images-3:concept-button}.

If the element is an
[`img`{#images-3:the-img-element-2}](embedded-content.html#the-img-element)
element that [represents](dom.html#represents){#images-3:represents-3}
some text and the user agent does not expect this to change

The user agent is expected to treat the element as a non-replaced
phrasing element whose content is the text, optionally with an icon
indicating that an image is missing, so that the user can request the
image be displayed or investigate why it is not rendering. In
non-graphical contexts, such an icon should be omitted.

If the element is an
[`img`{#images-3:the-img-element-3}](embedded-content.html#the-img-element)
element that [represents](dom.html#represents){#images-3:represents-4}
nothing and the user agent does not expect this to change

The user agent is expected to treat the element as a [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#images-3:replaced-element-3
x-internal="replaced-element"} whose [natural
dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#images-3:natural-dimensions-2
x-internal="natural-dimensions"} are 0. (In the absence of further
styles, this will cause the element to essentially not be rendered.)

If the element is an
[`input`{#images-3:the-input-element-3}](input.html#the-input-element)
element that does not
[represent](dom.html#represents){#images-3:represents-5} an image and
the user agent does not expect this to change

The user agent is expected to treat the element as a [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#images-3:replaced-element-4
x-internal="replaced-element"} consisting of a button whose content is
the element\'s alternative text. The [natural
dimensions](https://drafts.csswg.org/css-images/#natural-dimensions){#images-3:natural-dimensions-3
x-internal="natural-dimensions"} of the button are expected to be about
one line in height and whatever width is necessary to render the text on
one line.

The icons mentioned above are expected to be relatively small so as not
to disrupt most text but be easily clickable. In a visual environment,
for instance, icons could be 16 pixels by 16 pixels square, or 1em by
1em if the images are scalable. In an audio environment, the icon could
be a short bleep. The icons are intended to indicate to the user that
they can be used to get to whatever options the UA provides for images,
and, where appropriate, are expected to provide access to the context
menu that would have come up if the user interacted with the actual
image.

------------------------------------------------------------------------

All animated images with the same [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#images-3:absolute-url
x-internal="absolute-url"} and the same image data are expected to be
rendered synchronized to the same timeline as a group, with the timeline
starting at the time of the least recent addition to the group.

In other words, when a second image with the same [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#images-3:absolute-url-2
x-internal="absolute-url"} and animated image data is inserted into a
document, it jumps to the point in the animation cycle that is currently
being displayed by the first image.

When a user agent is to [restart the animation]{#restart-the-animation
.dfn} for an
[`img`{#images-3:the-img-element-4}](embedded-content.html#the-img-element)
element showing an animated image, all animated images with the same
[absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#images-3:absolute-url-3
x-internal="absolute-url"} and the same image data in that
[`img`{#images-3:the-img-element-5}](embedded-content.html#the-img-element)
element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#images-3:node-document
x-internal="node-document"} are expected to restart their animation from
the beginning.

------------------------------------------------------------------------

The following CSS rules are expected to apply:

``` css
@namespace "http://www.w3.org/1999/xhtml";

img:is([sizes="auto" i], [sizes^="auto," i]) {
  contain: size !important;
  contain-intrinsic-size: 300px 150px;
}
```

The following CSS rules are expected to apply when the
[`Document`{#images-3:document-2}](dom.html#document) is in [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#images-3:quirks-mode-2
x-internal="quirks-mode"}:

``` css
@namespace "http://www.w3.org/1999/xhtml";

img[align=left i] { margin-right: 3px; }
img[align=right i] { margin-left: 3px; }
```

#### [15.4.3]{.secno} Attributes for embedded content and images[](#attributes-for-embedded-content-and-images){.self-link}

The following CSS rules are expected to apply as [presentational
hints](#presentational-hints){#attributes-for-embedded-content-and-images:presentational-hints}:

``` css
@namespace "http://www.w3.org/1999/xhtml";

embed[align=left i], iframe[align=left i], img[align=left i],
input[type=image i][align=left i], object[align=left i] {
  float: left;
}

embed[align=right i], iframe[align=right i], img[align=right i],
input[type=image i][align=right i], object[align=right i] {
  float: right;
}

embed[align=top i], iframe[align=top i], img[align=top i],
input[type=image i][align=top i], object[align=top i] {
  vertical-align: top;
}

embed[align=baseline i], iframe[align=baseline i], img[align=baseline i],
input[type=image i][align=baseline i], object[align=baseline i] {
  vertical-align: baseline;
}

embed[align=texttop i], iframe[align=texttop i], img[align=texttop i],
input[type=image i][align=texttop i], object[align=texttop i] {
  vertical-align: text-top;
}

embed[align=absmiddle i], iframe[align=absmiddle i], img[align=absmiddle i],
input[type=image i][align=absmiddle i], object[align=absmiddle i],
embed[align=abscenter i], iframe[align=abscenter i], img[align=abscenter i],
input[type=image i][align=abscenter i], object[align=abscenter i] {
  vertical-align: middle;
}

embed[align=bottom i], iframe[align=bottom i], img[align=bottom i],
input[type=image i][align=bottom i], object[align=bottom i] {
  vertical-align: bottom;
}
```

When an
[`embed`{#attributes-for-embedded-content-and-images:the-embed-element}](iframe-embed-object.html#the-embed-element),
[`iframe`{#attributes-for-embedded-content-and-images:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
[`img`{#attributes-for-embedded-content-and-images:the-img-element}](embedded-content.html#the-img-element),
or
[`object`{#attributes-for-embedded-content-and-images:the-object-element}](iframe-embed-object.html#the-object-element)
element, or an
[`input`{#attributes-for-embedded-content-and-images:the-input-element}](input.html#the-input-element)
element whose
[`type`{#attributes-for-embedded-content-and-images:attr-input-type}](input.html#attr-input-type)
attribute is in the [Image
Button](input.html#image-button-state-(type=image)){#attributes-for-embedded-content-and-images:image-button-state-(type=image)}
state, has an `align` attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#attributes-for-embedded-content-and-images:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string \"`center`\"
or the string \"`middle`\", the user agent is expected to act as if the
element\'s
[\'vertical-align\'](https://drafts.csswg.org/css2/#propdef-vertical-align){#attributes-for-embedded-content-and-images:'vertical-align'
x-internal="'vertical-align'"} property was set to a value that aligns
the vertical middle of the element with the parent element\'s baseline.

The `hspace` attribute of
[`embed`{#attributes-for-embedded-content-and-images:the-embed-element-2}](iframe-embed-object.html#the-embed-element),
[`img`{#attributes-for-embedded-content-and-images:the-img-element-2}](embedded-content.html#the-img-element),
or
[`object`{#attributes-for-embedded-content-and-images:the-object-element-2}](iframe-embed-object.html#the-object-element)
elements, and
[`input`{#attributes-for-embedded-content-and-images:the-input-element-2}](input.html#the-input-element)
elements with a
[`type`{#attributes-for-embedded-content-and-images:attr-input-type-2}](input.html#attr-input-type)
attribute in the [Image
Button](input.html#image-button-state-(type=image)){#attributes-for-embedded-content-and-images:image-button-state-(type=image)-2}
state, [maps to the dimension
properties](#maps-to-the-dimension-property){#attributes-for-embedded-content-and-images:maps-to-the-dimension-property}
[\'margin-left\'](https://drafts.csswg.org/css-box/#propdef-margin-left){#attributes-for-embedded-content-and-images:'margin-left'
x-internal="'margin-left'"} and
[\'margin-right\'](https://drafts.csswg.org/css-box/#propdef-margin-right){#attributes-for-embedded-content-and-images:'margin-right'
x-internal="'margin-right'"} on the element.

The `vspace` attribute of
[`embed`{#attributes-for-embedded-content-and-images:the-embed-element-3}](iframe-embed-object.html#the-embed-element),
[`img`{#attributes-for-embedded-content-and-images:the-img-element-3}](embedded-content.html#the-img-element),
or
[`object`{#attributes-for-embedded-content-and-images:the-object-element-3}](iframe-embed-object.html#the-object-element)
elements, and
[`input`{#attributes-for-embedded-content-and-images:the-input-element-3}](input.html#the-input-element)
elements with a
[`type`{#attributes-for-embedded-content-and-images:attr-input-type-3}](input.html#attr-input-type)
attribute in the [Image
Button](input.html#image-button-state-(type=image)){#attributes-for-embedded-content-and-images:image-button-state-(type=image)-3}
state, [maps to the dimension
properties](#maps-to-the-dimension-property){#attributes-for-embedded-content-and-images:maps-to-the-dimension-property-2}
[\'margin-top\'](https://drafts.csswg.org/css-box/#propdef-margin-top){#attributes-for-embedded-content-and-images:'margin-top'
x-internal="'margin-top'"} and
[\'margin-bottom\'](https://drafts.csswg.org/css-box/#propdef-margin-bottom){#attributes-for-embedded-content-and-images:'margin-bottom'
x-internal="'margin-bottom'"} on the element.

When an
[`iframe`{#attributes-for-embedded-content-and-images:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
element has a
[`frameborder`{#attributes-for-embedded-content-and-images:attr-iframe-frameborder}](obsolete.html#attr-iframe-frameborder)
attribute whose value, when parsed using the [rules for parsing
integers](common-microsyntaxes.html#rules-for-parsing-integers){#attributes-for-embedded-content-and-images:rules-for-parsing-integers},
is zero or an error, the user agent is expected to have [presentational
hints](#presentational-hints){#attributes-for-embedded-content-and-images:presentational-hints-2}
setting the element\'s
[\'border-top-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-width){#attributes-for-embedded-content-and-images:'border-top-width'
x-internal="'border-top-width'"},
[\'border-right-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-width){#attributes-for-embedded-content-and-images:'border-right-width'
x-internal="'border-right-width'"},
[\'border-bottom-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-width){#attributes-for-embedded-content-and-images:'border-bottom-width'
x-internal="'border-bottom-width'"}, and
[\'border-left-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-width){#attributes-for-embedded-content-and-images:'border-left-width'
x-internal="'border-left-width'"} properties to zero.

When an
[`img`{#attributes-for-embedded-content-and-images:the-img-element-4}](embedded-content.html#the-img-element)
element,
[`object`{#attributes-for-embedded-content-and-images:the-object-element-4}](iframe-embed-object.html#the-object-element)
element, or
[`input`{#attributes-for-embedded-content-and-images:the-input-element-4}](input.html#the-input-element)
element with a
[`type`{#attributes-for-embedded-content-and-images:attr-input-type-4}](input.html#attr-input-type)
attribute in the [Image
Button](input.html#image-button-state-(type=image)){#attributes-for-embedded-content-and-images:image-button-state-(type=image)-4}
state has a `border` attribute whose value, when parsed using the [rules
for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#attributes-for-embedded-content-and-images:rules-for-parsing-non-negative-integers},
is found to be a number greater than zero, the user agent is expected to
use the parsed value for eight [presentational
hints](#presentational-hints){#attributes-for-embedded-content-and-images:presentational-hints-3}:
four setting the parsed value as a pixel length for the element\'s
[\'border-top-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-width){#attributes-for-embedded-content-and-images:'border-top-width'-2
x-internal="'border-top-width'"},
[\'border-right-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-width){#attributes-for-embedded-content-and-images:'border-right-width'-2
x-internal="'border-right-width'"},
[\'border-bottom-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-width){#attributes-for-embedded-content-and-images:'border-bottom-width'-2
x-internal="'border-bottom-width'"}, and
[\'border-left-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-width){#attributes-for-embedded-content-and-images:'border-left-width'-2
x-internal="'border-left-width'"} properties, and four setting the
element\'s
[\'border-top-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-style){#attributes-for-embedded-content-and-images:'border-top-style'
x-internal="'border-top-style'"},
[\'border-right-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-style){#attributes-for-embedded-content-and-images:'border-right-style'
x-internal="'border-right-style'"},
[\'border-bottom-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-style){#attributes-for-embedded-content-and-images:'border-bottom-style'
x-internal="'border-bottom-style'"}, and
[\'border-left-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-style){#attributes-for-embedded-content-and-images:'border-left-style'
x-internal="'border-left-style'"} properties to the value \'solid\'.

The
[`width`{#attributes-for-embedded-content-and-images:attr-dim-width}](embedded-content-other.html#attr-dim-width)
and
[`height`{#attributes-for-embedded-content-and-images:attr-dim-height}](embedded-content-other.html#attr-dim-height)
attributes on an
[`img`{#attributes-for-embedded-content-and-images:the-img-element-5}](embedded-content.html#the-img-element)
element\'s [dimension attribute
source](embedded-content.html#concept-img-dimension-attribute-source){#attributes-for-embedded-content-and-images:concept-img-dimension-attribute-source}
[map to the dimension
properties](#maps-to-the-dimension-property){#attributes-for-embedded-content-and-images:maps-to-the-dimension-property-3}
[\'width\'](https://drafts.csswg.org/css2/#the-width-property){#attributes-for-embedded-content-and-images:'width'
x-internal="'width'"} and
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#attributes-for-embedded-content-and-images:'height'
x-internal="'height'"} on the
[`img`{#attributes-for-embedded-content-and-images:the-img-element-6}](embedded-content.html#the-img-element)
element respectively. They similarly [map to the aspect-ratio property
(using dimension
rules)](#map-to-the-aspect-ratio-property-(using-dimension-rules)){#attributes-for-embedded-content-and-images:map-to-the-aspect-ratio-property-(using-dimension-rules)}
of the
[`img`{#attributes-for-embedded-content-and-images:the-img-element-7}](embedded-content.html#the-img-element)
element.

The
[`width`{#attributes-for-embedded-content-and-images:attr-dim-width-2}](embedded-content-other.html#attr-dim-width)
and
[`height`{#attributes-for-embedded-content-and-images:attr-dim-height-2}](embedded-content-other.html#attr-dim-height)
attributes on
[`embed`{#attributes-for-embedded-content-and-images:the-embed-element-4}](iframe-embed-object.html#the-embed-element),
[`iframe`{#attributes-for-embedded-content-and-images:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element),
[`object`{#attributes-for-embedded-content-and-images:the-object-element-5}](iframe-embed-object.html#the-object-element),
and
[`video`{#attributes-for-embedded-content-and-images:the-video-element}](media.html#the-video-element)
elements, and
[`input`{#attributes-for-embedded-content-and-images:the-input-element-5}](input.html#the-input-element)
elements with a
[`type`{#attributes-for-embedded-content-and-images:attr-input-type-5}](input.html#attr-input-type)
attribute in the [Image
Button](input.html#image-button-state-(type=image)){#attributes-for-embedded-content-and-images:image-button-state-(type=image)-5}
state and that either represents an image or that the user expects will
eventually represent an image, [map to the dimension
properties](#maps-to-the-dimension-property){#attributes-for-embedded-content-and-images:maps-to-the-dimension-property-4}
[\'width\'](https://drafts.csswg.org/css2/#the-width-property){#attributes-for-embedded-content-and-images:'width'-2
x-internal="'width'"} and
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#attributes-for-embedded-content-and-images:'height'-2
x-internal="'height'"} on the element respectively.

The
[`width`{#attributes-for-embedded-content-and-images:attr-dim-width-3}](embedded-content-other.html#attr-dim-width)
and
[`height`{#attributes-for-embedded-content-and-images:attr-dim-height-3}](embedded-content-other.html#attr-dim-height)
attributes [map to the aspect-ratio property (using dimension
rules)](#map-to-the-aspect-ratio-property-(using-dimension-rules)){#attributes-for-embedded-content-and-images:map-to-the-aspect-ratio-property-(using-dimension-rules)-2}
on
[`img`{#attributes-for-embedded-content-and-images:the-img-element-8}](embedded-content.html#the-img-element)
and
[`video`{#attributes-for-embedded-content-and-images:the-video-element-2}](media.html#the-video-element)
elements, and
[`input`{#attributes-for-embedded-content-and-images:the-input-element-6}](input.html#the-input-element)
elements with a
[`type`{#attributes-for-embedded-content-and-images:attr-input-type-6}](input.html#attr-input-type)
attribute in the [Image
Button](input.html#image-button-state-(type=image)){#attributes-for-embedded-content-and-images:image-button-state-(type=image)-6}
state.

The
[`width`{#attributes-for-embedded-content-and-images:attr-canvas-width}](canvas.html#attr-canvas-width)
and
[`height`{#attributes-for-embedded-content-and-images:attr-canvas-height}](canvas.html#attr-canvas-height)
attributes [map to the aspect-ratio
property](#map-to-the-aspect-ratio-property){#attributes-for-embedded-content-and-images:map-to-the-aspect-ratio-property}
on
[`canvas`{#attributes-for-embedded-content-and-images:the-canvas-element}](canvas.html#the-canvas-element)
elements.

#### [15.4.4]{.secno} Image maps[](#image-maps-2){.self-link} {#image-maps-2}

Shapes on an [image
map](image-maps.html#image-map){#image-maps-2:image-map} are expected to
act, for the purpose of the CSS cascade, as elements independent of the
original
[`area`{#image-maps-2:the-area-element}](image-maps.html#the-area-element)
element that happen to match the same style rules but inherit from the
[`img`{#image-maps-2:the-img-element}](embedded-content.html#the-img-element)
or
[`object`{#image-maps-2:the-object-element}](iframe-embed-object.html#the-object-element)
element.

For the purposes of the rendering, only the
[\'cursor\'](https://drafts.csswg.org/css-ui/#cursor){#image-maps-2:'cursor'
x-internal="'cursor'"} property is expected to have any effect on the
shape.

Thus, for example, if an
[`area`{#image-maps-2:the-area-element-2}](image-maps.html#the-area-element)
element has a [`style`{#image-maps-2:attr-style}](dom.html#attr-style)
attribute that sets the
[\'cursor\'](https://drafts.csswg.org/css-ui/#cursor){#image-maps-2:'cursor'-2
x-internal="'cursor'"} property to \'help\', then when the user
designates that shape, the cursor would change to a Help cursor.

Similarly, if an
[`area`{#image-maps-2:the-area-element-3}](image-maps.html#the-area-element)
element had a CSS rule that set its
[\'cursor\'](https://drafts.csswg.org/css-ui/#cursor){#image-maps-2:'cursor'-3
x-internal="'cursor'"} property to \'inherit\' (or if no rule setting
the
[\'cursor\'](https://drafts.csswg.org/css-ui/#cursor){#image-maps-2:'cursor'-4
x-internal="'cursor'"} property matched the element at all), the
shape\'s cursor would be inherited from the
[`img`{#image-maps-2:the-img-element-2}](embedded-content.html#the-img-element)
or
[`object`{#image-maps-2:the-object-element-2}](iframe-embed-object.html#the-object-element)
element of the [image
map](image-maps.html#image-map){#image-maps-2:image-map-2}, not from the
parent of the
[`area`{#image-maps-2:the-area-element-4}](image-maps.html#the-area-element)
element.

### [15.5]{.secno} []{#bindings}Widgets[](#widgets){.self-link}

#### [15.5.1]{.secno} Native appearance[](#native-appearance-2){.self-link} {#native-appearance-2}

The CSS Basic User Interface specification calls elements that can have
a [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#native-appearance-2:native-appearance
x-internal="native-appearance"}
[widgets](https://drafts.csswg.org/css-ui/#widget){#native-appearance-2:widget
x-internal="widget"}, and defines whether to use that [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#native-appearance-2:native-appearance-2
x-internal="native-appearance"} depending on the
[\'appearance\'](https://drafts.csswg.org/css-ui/#appearance-switching){#native-appearance-2:'appearance'
x-internal="'appearance'"} property. That logic, in turn, depends on
whether each the element is classified as a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#native-appearance-2:devolvable-widget
x-internal="devolvable-widget"} or [non-devolvable
widget](https://drafts.csswg.org/css-ui/#non-devolvable){#native-appearance-2:non-devolvable-widget
x-internal="non-devolvable-widget"}. This section defines which elements
match these concepts for HTML, what their [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#native-appearance-2:native-appearance-3
x-internal="native-appearance"} is, and any particularity of their
[devolved](https://drafts.csswg.org/css-ui/#devolved){#native-appearance-2:devolved-widget
x-internal="devolved-widget"} state or [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#native-appearance-2:primitive-appearance
x-internal="primitive-appearance"}.
[\[CSSUI\]](references.html#refsCSSUI)

The following elements can have a [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#native-appearance-2:native-appearance-4
x-internal="native-appearance"} for the purpose of the CSS
[\'appearance\'](https://drafts.csswg.org/css-ui/#appearance-switching){#native-appearance-2:'appearance'-2
x-internal="'appearance'"} property.

- [`button`{#native-appearance-2:the-button-element}](form-elements.html#the-button-element)
- [`input`{#native-appearance-2:the-input-element}](input.html#the-input-element)
- [`meter`{#native-appearance-2:the-meter-element}](form-elements.html#the-meter-element)
- [`progress`{#native-appearance-2:the-progress-element}](form-elements.html#the-progress-element)
- [`select`{#native-appearance-2:the-select-element}](form-elements.html#the-select-element)
- [`textarea`{#native-appearance-2:the-textarea-element}](form-elements.html#the-textarea-element)

#### [15.5.2]{.secno} Writing mode[](#writing-mode){.self-link}

Several widgets have their rendering controlled by the
[\'writing-mode\'](https://drafts.csswg.org/css-writing-modes/#propdef-writing-mode){#writing-mode:'writing-mode'
x-internal="'writing-mode'"} CSS property. For the purposes of those
widgets, we have the following definitions.

A [horizontal writing mode]{#horizontal-writing-mode .dfn} is when
resolving the
[\'writing-mode\'](https://drafts.csswg.org/css-writing-modes/#propdef-writing-mode){#writing-mode:'writing-mode'-2
x-internal="'writing-mode'"} property of the control results in a
computed value of \'horizontal-tb\'.

A [vertical writing mode]{#vertical-writing-mode .dfn} is when resolving
the
[\'writing-mode\'](https://drafts.csswg.org/css-writing-modes/#propdef-writing-mode){#writing-mode:'writing-mode'-3
x-internal="'writing-mode'"} property of the control results in a
computed value of either \'vertical-rl\', \'vertical-lr\',
\'sideways-rl\' or \'sideways-lr\'.

#### [15.5.3]{.secno} Button layout[](#button-layout){.self-link}

When an element uses [button
layout](#button-layout-2){#button-layout:button-layout-2}, it is a
[devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#button-layout:devolvable-widget
x-internal="devolvable-widget"}, and it\'s [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#button-layout:native-appearance
x-internal="native-appearance"} is that of a button.

[Button layout]{#button-layout-2 .dfn} is as follows:

- If the element is a
  [`button`{#button-layout:the-button-element}](form-elements.html#the-button-element)
  element, then the
  [\'display\'](https://drafts.csswg.org/css2/#display-prop){#button-layout:'display'
  x-internal="'display'"} property is expected to act as follows:

  - If the computed value of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#button-layout:'display'-2
    x-internal="'display'"} is \'inline-grid\', \'grid\',
    \'inline-flex\', \'flex\', \'none\', or \'contents\', then behave as
    the computed value.

  - Otherwise, if the computed value of
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#button-layout:'display'-3
    x-internal="'display'"} is a value such that the [outer display
    type](https://drafts.csswg.org/css-display/#outer-display-type){#button-layout:outer-display-type
    x-internal="outer-display-type"} is \'inline\', then behave as
    \'inline-block\'.

  - Otherwise, behave as \'flow-root\'.

- The element is expected to establish a new [formatting
  context](https://drafts.csswg.org/css-display/#formatting-context){#button-layout:formatting-context
  x-internal="formatting-context"} for its contents. The type of this
  formatting context is determined by its
  [\'display\'](https://drafts.csswg.org/css2/#display-prop){#button-layout:'display'-4
  x-internal="'display'"} value, as usual.

- If the element is
  [absolutely-positioned](https://drafts.csswg.org/css-position/#absolute-position){#button-layout:absolutely-positioned
  x-internal="absolutely-positioned"}, then for the purpose of the [CSS
  visual formatting model](https://drafts.csswg.org/css2/#visuren), act
  as if the element is a [replaced
  element](https://drafts.csswg.org/css-display/#replaced-element){#button-layout:replaced-element
  x-internal="replaced-element"}. [\[CSS\]](references.html#refsCSS)

- If the [computed
  value](https://drafts.csswg.org/css-cascade/#computed-value){#button-layout:computed-value
  x-internal="computed-value"} of
  [\'inline-size\'](https://drafts.csswg.org/css-logical/#propdef-inline-size){#button-layout:'inline-size'
  x-internal="'inline-size'"} is \'auto\', then the [used
  value](https://drafts.csswg.org/css-cascade/#used-value){#button-layout:used-value
  x-internal="used-value"} is the [fit-content inline
  size](https://drafts.csswg.org/css-sizing/#fit-content-inline-size){#button-layout:fit-content-inline-size
  x-internal="fit-content-inline-size"}.

- For the purpose of the \'normal\' keyword of the
  [\'align-self\'](https://drafts.csswg.org/css-align/#propdef-align-self){#button-layout:'align-self'
  x-internal="'align-self'"} property, act as if the element is a
  replaced element.

- If the element is an
  [`input`{#button-layout:the-input-element}](input.html#the-input-element)
  element, or if it is a
  [`button`{#button-layout:the-button-element-2}](form-elements.html#the-button-element)
  element and its computed value for
  [\'display\'](https://drafts.csswg.org/css2/#display-prop){#button-layout:'display'-5
  x-internal="'display'"} is not \'inline-grid\', \'grid\',
  \'inline-flex\', or \'flex\', then the element\'s box has a child
  [anonymous button content box]{#anonymous-button-content-box .dfn}
  with the following behaviors:

  - The box is a
    [block-level](https://drafts.csswg.org/css-display/#block-level){#button-layout:block-level
    x-internal="block-level"} [block
    container](https://drafts.csswg.org/css-display/#block-container){#button-layout:block-container
    x-internal="block-container"} that establishes a new [block
    formatting
    context](https://drafts.csswg.org/css-display/#block-formatting-context){#button-layout:block-formatting-context
    x-internal="block-formatting-context"} (i.e.,
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#button-layout:'display'-6
    x-internal="'display'"} is \'flow-root\').

  - If the box does not overflow in the horizontal axis, then it is
    centered horizontally.

  - If the box does not overflow in the vertical axis, then it is
    centered vertically.

  Otherwise, there is no [anonymous button content
  box](#anonymous-button-content-box){#button-layout:anonymous-button-content-box}.

Need to define the expected [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#button-layout:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.4]{.secno} The [`button`{#the-button-element-2:the-button-element}](form-elements.html#the-button-element) element[](#the-button-element-2){.self-link} {#the-button-element-2}

The
[`button`{#the-button-element-2:the-button-element-2}](form-elements.html#the-button-element)
element, when it generates a [CSS
box](https://drafts.csswg.org/css-display/#css-box){#the-button-element-2:css-box
x-internal="css-box"}, is expected to depict a button and to use [button
layout](#button-layout-2){#the-button-element-2:button-layout-2} whose
[anonymous button content
box](#anonymous-button-content-box){#the-button-element-2:anonymous-button-content-box}\'s
contents (if there is an [anonymous button content
box](#anonymous-button-content-box){#the-button-element-2:anonymous-button-content-box-2})
are the child boxes the element\'s box would otherwise have.

#### [15.5.5]{.secno} The [`details`{#the-details-and-summary-elements:the-details-element}](interactive-elements.html#the-details-element) and [`summary`{#the-details-and-summary-elements:the-summary-element}](interactive-elements.html#the-summary-element) elements[](#the-details-and-summary-elements){.self-link}

``` css
@namespace "http://www.w3.org/1999/xhtml";

details, summary {
  display: block;
}
details > summary:first-of-type {
  display: list-item;
  counter-increment: list-item 0;
  list-style: disclosure-closed inside;
}
details[open] > summary:first-of-type {
  list-style-type: disclosure-open;
}
```

The
[`details`{#the-details-and-summary-elements:the-details-element-2}](interactive-elements.html#the-details-element)
element is expected to have an internal [shadow
tree](https://dom.spec.whatwg.org/#concept-shadow-tree){#the-details-and-summary-elements:shadow-tree
x-internal="shadow-tree"} with three child elements:

1.  The first child element is a
    [`slot`{#the-details-and-summary-elements:the-slot-element}](scripting.html#the-slot-element)
    that is expected to take the
    [`details`{#the-details-and-summary-elements:the-details-element-3}](interactive-elements.html#the-details-element)
    element\'s first
    [`summary`{#the-details-and-summary-elements:the-summary-element-2}](interactive-elements.html#the-summary-element)
    element child, if any. This element has a single child
    [`summary`{#the-details-and-summary-elements:the-summary-element-3}](interactive-elements.html#the-summary-element)
    element called the [default summary]{#default-summary .dfn} which
    has text content that is
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-details-and-summary-elements:implementation-defined
    x-internal="implementation-defined"} (and probably locale-specific).

    The
    [`summary`{#the-details-and-summary-elements:the-summary-element-4}](interactive-elements.html#the-summary-element)
    element that this slot
    [represents](dom.html#represents){#the-details-and-summary-elements:represents}
    is expected to allow the user to request the details be shown or
    hidden.

2.  The second child element is a
    [`slot`{#the-details-and-summary-elements:the-slot-element-2}](scripting.html#the-slot-element)
    that is expected to take the
    [`details`{#the-details-and-summary-elements:the-details-element-4}](interactive-elements.html#the-details-element)
    element\'s remaining descendants, if any. This element has no
    contents.

    This element is expected to match the
    [\'::details-content\'](https://drafts.csswg.org/css-pseudo/#details-content-pseudo){#the-details-and-summary-elements:'::details-content'
    x-internal="'::details-content'"} pseudo-element.

    This element is expected to have its
    [`style`{#the-details-and-summary-elements:attr-style}](dom.html#attr-style)
    attribute set to \"`display: block; content-visibility: hidden;`\"
    when the
    [`details`{#the-details-and-summary-elements:the-details-element-5}](interactive-elements.html#the-details-element)
    element does not have an
    [`open`{#the-details-and-summary-elements:attr-details-open}](interactive-elements.html#attr-details-open)
    attribute. When it does have the
    [`open`{#the-details-and-summary-elements:attr-details-open-2}](interactive-elements.html#attr-details-open)
    attribute, the
    [`style`{#the-details-and-summary-elements:attr-style-2}](dom.html#attr-style)
    attribute is expected to be set to \"`display: block;`\".

    Because the slots are hidden inside a shadow tree, this
    [`style`{#the-details-and-summary-elements:attr-style-3}](dom.html#attr-style)
    attribute is not directly visible to author code. Its impacts,
    however, are visible. Notably, the choice of
    `content-visibility: hidden` instead of, e.g., `display: none`,
    impacts the results of various APIs that query layout information.

3.  The third child element is either a
    [`link`{#the-details-and-summary-elements:the-link-element}](semantics.html#the-link-element)
    or
    [`style`{#the-details-and-summary-elements:the-style-element}](semantics.html#the-style-element)
    element with the following styles for the [default
    summary](#default-summary){#the-details-and-summary-elements:default-summary}:

    ``` css
    :host summary {
      display: list-item;
      counter-increment: list-item 0;
      list-style: disclosure-closed inside;
    }
    :host([open]) summary {
      list-style-type: disclosure-open;
    }
    ```

    The position of this child element relative to the other two is not
    observable. This means that implementations might have it in a
    different order relative to its siblings. Implementations might even
    associate the style with the shadow tree using a mechanism that is
    not an element.

The structure of this shadow tree is observable through the ways that
the children of the
[`details`{#the-details-and-summary-elements:the-details-element-6}](interactive-elements.html#the-details-element)
element and the \'::details-content\' pseudo-element respond to CSS
styles.

#### [15.5.6]{.secno} The [`input`{#the-input-element-as-a-text-entry-widget:the-input-element}](input.html#the-input-element) element as a text entry widget[](#the-input-element-as-a-text-entry-widget){.self-link}

An
[`input`{#the-input-element-as-a-text-entry-widget:the-input-element-2}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-text-entry-widget:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#the-input-element-as-a-text-entry-widget:text-(type=text)-state-and-search-state-(type=search)},
[Telephone](input.html#telephone-state-(type=tel)){#the-input-element-as-a-text-entry-widget:telephone-state-(type=tel)},
[URL](input.html#url-state-(type=url)){#the-input-element-as-a-text-entry-widget:url-state-(type=url)},
or
[Email](input.html#email-state-(type=email)){#the-input-element-as-a-text-entry-widget:email-state-(type=email)}
state, is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-a-text-entry-widget:devolvable-widget
x-internal="devolvable-widget"}. Its expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-text-entry-widget:native-appearance
x-internal="native-appearance"} is to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-a-text-entry-widget:'inline-block'
x-internal="'inline-block'"} box depicting a one-line text control.

An
[`input`{#the-input-element-as-a-text-entry-widget:the-input-element-3}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-text-entry-widget:attr-input-type-2}](input.html#attr-input-type)
attribute is in the
[Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#the-input-element-as-a-text-entry-widget:text-(type=text)-state-and-search-state-(type=search)-2}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-a-text-entry-widget:devolvable-widget-2
x-internal="devolvable-widget"}. Its expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-text-entry-widget:native-appearance-2
x-internal="native-appearance"} is to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-a-text-entry-widget:'inline-block'-2
x-internal="'inline-block'"} box depicting a one-line text control. If
the [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-input-element-as-a-text-entry-widget:computed-value
x-internal="computed-value"} of the element\'s
[\'appearance\'](https://drafts.csswg.org/css-ui/#appearance-switching){#the-input-element-as-a-text-entry-widget:'appearance'
x-internal="'appearance'"} property is not
[`'textfield'`{#the-input-element-as-a-text-entry-widget:'textfield'}](https://drafts.csswg.org/css-ui/#valdef-appearance-textfield){x-internal="'textfield'"},
it may have a distinct style indicating that it is a search field.

An
[`input`{#the-input-element-as-a-text-entry-widget:the-input-element-4}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-text-entry-widget:attr-input-type-3}](input.html#attr-input-type)
attribute is in the
[Password](input.html#password-state-(type=password)){#the-input-element-as-a-text-entry-widget:password-state-(type=password)}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-a-text-entry-widget:devolvable-widget-3
x-internal="devolvable-widget"}. Its expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-text-entry-widget:native-appearance-3
x-internal="native-appearance"} is to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-a-text-entry-widget:'inline-block'-3
x-internal="'inline-block'"} box depicting a one-line text control that
obscures data entry.

For
[`input`{#the-input-element-as-a-text-entry-widget:the-input-element-5}](input.html#the-input-element)
elements whose
[`type`{#the-input-element-as-a-text-entry-widget:attr-input-type-4}](input.html#attr-input-type)
attribute is in one of the above states, the [used
value](https://drafts.csswg.org/css-cascade/#used-value){#the-input-element-as-a-text-entry-widget:used-value
x-internal="used-value"} of the
[\'line-height\'](https://drafts.csswg.org/css2/#propdef-line-height){#the-input-element-as-a-text-entry-widget:'line-height'
x-internal="'line-height'"} property must be a length value that is no
smaller than what the [used
value](https://drafts.csswg.org/css-cascade/#used-value){#the-input-element-as-a-text-entry-widget:used-value-2
x-internal="used-value"} would be for \'line-height: normal\'.

The [used
value](https://drafts.csswg.org/css-cascade/#used-value){#the-input-element-as-a-text-entry-widget:used-value-3
x-internal="used-value"} will not be the actual keyword \'normal\'.
Also, this rule does not affect the [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-input-element-as-a-text-entry-widget:computed-value-2
x-internal="computed-value"}.

If these text controls provide a text selection, then, when the user
changes the current selection, the user agent is expected to [queue an
element
task](webappapis.html#queue-an-element-task){#the-input-element-as-a-text-entry-widget:queue-an-element-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#the-input-element-as-a-text-entry-widget:user-interaction-task-source}
given the
[`input`{#the-input-element-as-a-text-entry-widget:the-input-element-6}](input.html#the-input-element)
element to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#the-input-element-as-a-text-entry-widget:concept-event-fire
x-internal="concept-event-fire"} named
[`select`{#the-input-element-as-a-text-entry-widget:event-select}](indices.html#event-select)
at the element, with the
[`bubbles`{#the-input-element-as-a-text-entry-widget:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
attribute initialized to true.

An
[`input`{#the-input-element-as-a-text-entry-widget:the-input-element-7}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-text-entry-widget:attr-input-type-5}](input.html#attr-input-type)
attribute is in one of the above states is an [element with default
preferred
size](https://drafts.csswg.org/css-ui/#element-with-default-preferred-size){#the-input-element-as-a-text-entry-widget:element-with-default-preferred-size
x-internal="element-with-default-preferred-size"}, and user agents are
expected to apply the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-input-element-as-a-text-entry-widget:'field-sizing'
x-internal="'field-sizing'"} CSS property to the element. User agents
are expected to determine the [inline
size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-input-element-as-a-text-entry-widget:inline-size
x-internal="inline-size"} of its [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-input-element-as-a-text-entry-widget:intrinsic-size
x-internal="intrinsic-size"} by the following steps:

1.  If the
    [\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-input-element-as-a-text-entry-widget:'field-sizing'-2
    x-internal="'field-sizing'"} property on the element has a [computed
    value](https://drafts.csswg.org/css-cascade/#computed-value){#the-input-element-as-a-text-entry-widget:computed-value-3
    x-internal="computed-value"} of
    [\'content\'](https://drafts.csswg.org/css-ui/#valdef-field-sizing-content){#the-input-element-as-a-text-entry-widget:field-sizing-content
    x-internal="field-sizing-content"}, the [inline
    size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-input-element-as-a-text-entry-widget:inline-size-2
    x-internal="inline-size"} is determined by the text which the
    element shows. The text is either a
    [value](form-control-infrastructure.html#concept-fe-value){#the-input-element-as-a-text-entry-widget:concept-fe-value}
    or a short hint specified by the
    [`placeholder`{#the-input-element-as-a-text-entry-widget:attr-input-placeholder}](input.html#attr-input-placeholder)
    attribute. User agents may take the text caret size into account in
    the [inline
    size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-input-element-as-a-text-entry-widget:inline-size-3
    x-internal="inline-size"}.

2.  If the element has a
    [`size`{#the-input-element-as-a-text-entry-widget:attr-input-size}](input.html#attr-input-size)
    attribute, and parsing that attribute\'s value using the [rules for
    parsing non-negative
    integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-input-element-as-a-text-entry-widget:rules-for-parsing-non-negative-integers}
    doesn\'t generate an error, return the value obtained from applying
    the [converting a character width to
    pixels](#converting-a-character-width-to-pixels){#the-input-element-as-a-text-entry-widget:converting-a-character-width-to-pixels}
    algorithm to the value of the attribute.

3.  Otherwise, return the value obtained from applying the [converting a
    character width to
    pixels](#converting-a-character-width-to-pixels){#the-input-element-as-a-text-entry-widget:converting-a-character-width-to-pixels-2}
    algorithm to the number 20.

The [converting a character width to
pixels]{#converting-a-character-width-to-pixels .dfn} algorithm returns
(`size`{.variable}-1)×`avg`{.variable} + `max`{.variable}, where
`size`{.variable} is the character width to convert, `avg`{.variable} is
the average character width of the primary font for the element for
which the algorithm is being run, in pixels, and `max`{.variable} is the
maximum character width of that same font, also in pixels. (The
element\'s
[\'letter-spacing\'](https://drafts.csswg.org/css-text/#letter-spacing-property){#the-input-element-as-a-text-entry-widget:'letter-spacing'
x-internal="'letter-spacing'"} property does not affect the result.)

These text controls are expected to be [scroll
containers](https://drafts.csswg.org/css-overflow/#scroll-container){#the-input-element-as-a-text-entry-widget:scroll-container
x-internal="scroll-container"} and support scrolling in the [inline
axis](https://drafts.csswg.org/css-writing-modes/#inline-axis){#the-input-element-as-a-text-entry-widget:inline-axis
x-internal="inline-axis"}, but not the [block
axis](https://drafts.csswg.org/css-writing-modes/#block-axis){#the-input-element-as-a-text-entry-widget:block-axis
x-internal="block-axis"}.

Need to detail the expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-text-entry-widget:native-appearance-4
x-internal="native-appearance"} and [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-input-element-as-a-text-entry-widget:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.7]{.secno} The [`input`{#the-input-element-as-domain-specific-widgets:the-input-element}](input.html#the-input-element) element as domain-specific widgets[](#the-input-element-as-domain-specific-widgets){.self-link}

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-2}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Date](input.html#date-state-(type=date)){#the-input-element-as-domain-specific-widgets:date-state-(type=date)}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-domain-specific-widgets:devolvable-widget
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-domain-specific-widgets:'inline-block'
x-internal="'inline-block'"} box depicting a date control.

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-3}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type-2}](input.html#attr-input-type)
attribute is in the
[Month](input.html#month-state-(type=month)){#the-input-element-as-domain-specific-widgets:month-state-(type=month)}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-domain-specific-widgets:devolvable-widget-2
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-domain-specific-widgets:'inline-block'-2
x-internal="'inline-block'"} box depicting a month control.

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-4}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type-3}](input.html#attr-input-type)
attribute is in the
[Week](input.html#week-state-(type=week)){#the-input-element-as-domain-specific-widgets:week-state-(type=week)}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-domain-specific-widgets:devolvable-widget-3
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-domain-specific-widgets:'inline-block'-3
x-internal="'inline-block'"} box depicting a week control.

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-5}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type-4}](input.html#attr-input-type)
attribute is in the
[Time](input.html#time-state-(type=time)){#the-input-element-as-domain-specific-widgets:time-state-(type=time)}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-domain-specific-widgets:devolvable-widget-4
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-domain-specific-widgets:'inline-block'-4
x-internal="'inline-block'"} box depicting a time control.

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-6}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type-5}](input.html#attr-input-type)
attribute is in the [Local Date and
Time](input.html#local-date-and-time-state-(type=datetime-local)){#the-input-element-as-domain-specific-widgets:local-date-and-time-state-(type=datetime-local)}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-domain-specific-widgets:devolvable-widget-5
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-domain-specific-widgets:'inline-block'-5
x-internal="'inline-block'"} box depicting a local date and time
control.

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-7}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type-6}](input.html#attr-input-type)
attribute is in the
[Number](input.html#number-state-(type=number)){#the-input-element-as-domain-specific-widgets:number-state-(type=number)}
state is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-input-element-as-domain-specific-widgets:devolvable-widget-6
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-domain-specific-widgets:'inline-block'-6
x-internal="'inline-block'"} box depicting a number control.

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-8}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type-7}](input.html#attr-input-type)
attribute is in the
[Number](input.html#number-state-(type=number)){#the-input-element-as-domain-specific-widgets:number-state-(type=number)-2}
state is an [element with default preferred
size](https://drafts.csswg.org/css-ui/#element-with-default-preferred-size){#the-input-element-as-domain-specific-widgets:element-with-default-preferred-size
x-internal="element-with-default-preferred-size"}, and user agents are
expected to apply the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-input-element-as-domain-specific-widgets:'field-sizing'
x-internal="'field-sizing'"} CSS property to the element. The [block
size](https://drafts.csswg.org/css-writing-modes/#block-size){#the-input-element-as-domain-specific-widgets:block-size
x-internal="block-size"} of the [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-input-element-as-domain-specific-widgets:intrinsic-size
x-internal="intrinsic-size"} is about one line high. If the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-input-element-as-domain-specific-widgets:'field-sizing'-2
x-internal="'field-sizing'"} property on the element has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-input-element-as-domain-specific-widgets:computed-value
x-internal="computed-value"} of
[\'content\'](https://drafts.csswg.org/css-ui/#valdef-field-sizing-content){#the-input-element-as-domain-specific-widgets:field-sizing-content
x-internal="field-sizing-content"}, the [inline
size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-input-element-as-domain-specific-widgets:inline-size
x-internal="inline-size"} of the [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-input-element-as-domain-specific-widgets:intrinsic-size-2
x-internal="intrinsic-size"} is expected to be about as wide as
necessary to show the current
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element-as-domain-specific-widgets:concept-fe-value}.
Otherwise, the [inline
size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-input-element-as-domain-specific-widgets:inline-size-2
x-internal="inline-size"} of the [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-input-element-as-domain-specific-widgets:intrinsic-size-3
x-internal="intrinsic-size"} is expected to be about as wide as
necessary to show the widest possible value.

An
[`input`{#the-input-element-as-domain-specific-widgets:the-input-element-9}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-domain-specific-widgets:attr-input-type-8}](input.html#attr-input-type)
attribute is in the
[Date](input.html#date-state-(type=date)){#the-input-element-as-domain-specific-widgets:date-state-(type=date)-2},
[Month](input.html#month-state-(type=month)){#the-input-element-as-domain-specific-widgets:month-state-(type=month)-2},
[Week](input.html#week-state-(type=week)){#the-input-element-as-domain-specific-widgets:week-state-(type=week)-2},
[Time](input.html#time-state-(type=time)){#the-input-element-as-domain-specific-widgets:time-state-(type=time)-2},
or [Local Date and
Time](input.html#local-date-and-time-state-(type=datetime-local)){#the-input-element-as-domain-specific-widgets:local-date-and-time-state-(type=datetime-local)-2}
state, is expected to be about one line high, and about as wide as
necessary to show the widest possible value.

Need to detail the expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-domain-specific-widgets:native-appearance
x-internal="native-appearance"} and [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-input-element-as-domain-specific-widgets:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.8]{.secno} The [`input`{#the-input-element-as-a-range-control:the-input-element}](input.html#the-input-element) element as a range control[](#the-input-element-as-a-range-control){.self-link}

An
[`input`{#the-input-element-as-a-range-control:the-input-element-2}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-range-control:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Range](input.html#range-state-(type=range)){#the-input-element-as-a-range-control:range-state-(type=range)}
state is a [non-devolvable
widget](https://drafts.csswg.org/css-ui/#non-devolvable){#the-input-element-as-a-range-control:non-devolvable-widget
x-internal="non-devolvable-widget"}. Its expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-range-control:native-appearance
x-internal="native-appearance"} is to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-a-range-control:'inline-block'
x-internal="'inline-block'"} box depicting a slider control.

When this control has a [horizontal writing
mode](#horizontal-writing-mode){#the-input-element-as-a-range-control:horizontal-writing-mode},
the control is expected to be a horizontal slider. Its lowest value is
on the right if the
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#the-input-element-as-a-range-control:'direction'
x-internal="'direction'"} property has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-input-element-as-a-range-control:computed-value
x-internal="computed-value"} of \'rtl\', and on the left otherwise. When
this control has a [vertical writing
mode](#vertical-writing-mode){#the-input-element-as-a-range-control:vertical-writing-mode},
it is expected to be a vertical slider. Its lowest value is on the
bottom if the
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#the-input-element-as-a-range-control:'direction'-2
x-internal="'direction'"} property has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-input-element-as-a-range-control:computed-value-2
x-internal="computed-value"} of \'rtl\', and on the top otherwise.

Predefined suggested values (provided by the
[`list`{#the-input-element-as-a-range-control:attr-input-list}](input.html#attr-input-list)
attribute) are expected to be shown as tick marks on the slider, which
the slider can snap to.

Need to detail the expected [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-input-element-as-a-range-control:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.9]{.secno} The [`input`{#the-input-element-as-a-colour-well:the-input-element}](input.html#the-input-element) element as a color well[](#the-input-element-as-a-colour-well){.self-link} {#the-input-element-as-a-colour-well}

An
[`input`{#the-input-element-as-a-colour-well:the-input-element-2}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-colour-well:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Color](input.html#color-state-(type=color)){#the-input-element-as-a-colour-well:color-state-(type=color)}
state is expected to depict a color well, which, when activated,
provides the user with a color picker (e.g. a color wheel or color
palette) from which the color can be changed. The element, when it
generates a [CSS
box](https://drafts.csswg.org/css-display/#css-box){#the-input-element-as-a-colour-well:css-box
x-internal="css-box"}, is expected to use [button
layout](#button-layout-2){#the-input-element-as-a-colour-well:button-layout-2},
that has no child boxes of the [anonymous button content
box](#anonymous-button-content-box){#the-input-element-as-a-colour-well:anonymous-button-content-box}.
The [anonymous button content
box](#anonymous-button-content-box){#the-input-element-as-a-colour-well:anonymous-button-content-box-2}
is expected to have a [presentational
hint](#presentational-hints){#the-input-element-as-a-colour-well:presentational-hints}
setting the
[\'background-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-color){#the-input-element-as-a-colour-well:'background-color'
x-internal="'background-color'"} property to the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element-as-a-colour-well:concept-fe-value}.

Predefined suggested values (provided by the
[`list`{#the-input-element-as-a-colour-well:attr-input-list}](input.html#attr-input-list)
attribute) are expected to be shown in the color picker interface, not
on the color well itself.

Need to detail the expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-colour-well:native-appearance
x-internal="native-appearance"} and [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-input-element-as-a-colour-well:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.10]{.secno} The [`input`{#the-input-element-as-a-checkbox-and-radio-button-widgets:the-input-element}](input.html#the-input-element) element as a checkbox and radio button widgets[](#the-input-element-as-a-checkbox-and-radio-button-widgets){.self-link}

An
[`input`{#the-input-element-as-a-checkbox-and-radio-button-widgets:the-input-element-2}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-checkbox-and-radio-button-widgets:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Checkbox](input.html#checkbox-state-(type=checkbox)){#the-input-element-as-a-checkbox-and-radio-button-widgets:checkbox-state-(type=checkbox)}
state is a [non-devolvable
widget](https://drafts.csswg.org/css-ui/#non-devolvable){#the-input-element-as-a-checkbox-and-radio-button-widgets:non-devolvable-widget
x-internal="non-devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-a-checkbox-and-radio-button-widgets:'inline-block'
x-internal="'inline-block'"} box containing a single checkbox control,
with no label.

Need to detail the expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-checkbox-and-radio-button-widgets:native-appearance
x-internal="native-appearance"} and [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-input-element-as-a-checkbox-and-radio-button-widgets:primitive-appearance
x-internal="primitive-appearance"}.

An
[`input`{#the-input-element-as-a-checkbox-and-radio-button-widgets:the-input-element-3}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-checkbox-and-radio-button-widgets:attr-input-type-2}](input.html#attr-input-type)
attribute is in the [Radio
Button](input.html#radio-button-state-(type=radio)){#the-input-element-as-a-checkbox-and-radio-button-widgets:radio-button-state-(type=radio)}
state is a [non-devolvable
widget](https://drafts.csswg.org/css-ui/#non-devolvable){#the-input-element-as-a-checkbox-and-radio-button-widgets:non-devolvable-widget-2
x-internal="non-devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-a-checkbox-and-radio-button-widgets:'inline-block'-2
x-internal="'inline-block'"} box containing a single radio button
control, with no label.

Need to detail the expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-input-element-as-a-checkbox-and-radio-button-widgets:native-appearance-2
x-internal="native-appearance"} and [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-input-element-as-a-checkbox-and-radio-button-widgets:primitive-appearance-2
x-internal="primitive-appearance"}.

#### [15.5.11]{.secno} The [`input`{#the-input-element-as-a-file-upload-control:the-input-element}](input.html#the-input-element) element as a file upload control[](#the-input-element-as-a-file-upload-control){.self-link}

An
[`input`{#the-input-element-as-a-file-upload-control:the-input-element-2}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-file-upload-control:attr-input-type}](input.html#attr-input-type)
attribute is in the [File
Upload](input.html#file-upload-state-(type=file)){#the-input-element-as-a-file-upload-control:file-upload-state-(type=file)}
state, when it generates a [CSS
box](https://drafts.csswg.org/css-display/#css-box){#the-input-element-as-a-file-upload-control:css-box
x-internal="css-box"}, is expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-input-element-as-a-file-upload-control:'inline-block'
x-internal="'inline-block'"} box containing a span of text giving the
filename(s) of the [selected
files](input.html#concept-input-type-file-selected){#the-input-element-as-a-file-upload-control:concept-input-type-file-selected},
if any, followed by a button that, when activated, provides the user
with a file picker from which the selection can be changed. The button
is expected to use [button
layout](#button-layout-2){#the-input-element-as-a-file-upload-control:button-layout-2}
and match the
[\'::file-selector-button\'](https://drafts.csswg.org/css-pseudo/#file-selector-button-pseudo){#the-input-element-as-a-file-upload-control:'::file-selector-button'
x-internal="'::file-selector-button'"} pseudo-element. The contents of
its [anonymous button content
box](#anonymous-button-content-box){#the-input-element-as-a-file-upload-control:anonymous-button-content-box}
are expected to be
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-input-element-as-a-file-upload-control:implementation-defined
x-internal="implementation-defined"} (and possibly locale-specific)
text, for example \"Choose file\".

User agents may handle an
[`input`{#the-input-element-as-a-file-upload-control:the-input-element-3}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-file-upload-control:attr-input-type-2}](input.html#attr-input-type)
attribute is in the [File
Upload](input.html#file-upload-state-(type=file)){#the-input-element-as-a-file-upload-control:file-upload-state-(type=file)-2}
state as an [element with default preferred
size](https://drafts.csswg.org/css-ui/#element-with-default-preferred-size){#the-input-element-as-a-file-upload-control:element-with-default-preferred-size
x-internal="element-with-default-preferred-size"}, and user agents may
apply the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-input-element-as-a-file-upload-control:'field-sizing'
x-internal="'field-sizing'"} CSS property to the element. If the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-input-element-as-a-file-upload-control:'field-sizing'-2
x-internal="'field-sizing'"} property on the element has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-input-element-as-a-file-upload-control:computed-value
x-internal="computed-value"} of
[\'content\'](https://drafts.csswg.org/css-ui/#valdef-field-sizing-content){#the-input-element-as-a-file-upload-control:field-sizing-content
x-internal="field-sizing-content"}, the [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-input-element-as-a-file-upload-control:intrinsic-size
x-internal="intrinsic-size"} of the element is expected to depend on its
content such as the
[\'::file-selector-button\'](https://drafts.csswg.org/css-pseudo/#file-selector-button-pseudo){#the-input-element-as-a-file-upload-control:'::file-selector-button'-2
x-internal="'::file-selector-button'"} pseudo-element and chosen file
names.

#### [15.5.12]{.secno} The [`input`{#the-input-element-as-a-button:the-input-element}](input.html#the-input-element) element as a button[](#the-input-element-as-a-button){.self-link}

An
[`input`{#the-input-element-as-a-button:the-input-element-2}](input.html#the-input-element)
element whose
[`type`{#the-input-element-as-a-button:attr-input-type}](input.html#attr-input-type)
attribute is in the [Submit
Button](input.html#submit-button-state-(type=submit)){#the-input-element-as-a-button:submit-button-state-(type=submit)},
[Reset
Button](input.html#reset-button-state-(type=reset)){#the-input-element-as-a-button:reset-button-state-(type=reset)},
or
[Button](input.html#button-state-(type=button)){#the-input-element-as-a-button:button-state-(type=button)}
state, when it generates a [CSS
box](https://drafts.csswg.org/css-display/#css-box){#the-input-element-as-a-button:css-box
x-internal="css-box"}, is expected to depict a button and use [button
layout](#button-layout-2){#the-input-element-as-a-button:button-layout-2}
and the contents of the [anonymous button content
box](#anonymous-button-content-box){#the-input-element-as-a-button:anonymous-button-content-box}
are expected to be the text of the element\'s
[`value`{#the-input-element-as-a-button:attr-input-value}](input.html#attr-input-value)
attribute, if any, or text derived from the element\'s
[`type`{#the-input-element-as-a-button:attr-input-type-2}](input.html#attr-input-type)
attribute in an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-input-element-as-a-button:implementation-defined
x-internal="implementation-defined"} (and probably locale-specific)
fashion, if not.

#### [15.5.13]{.secno} The [`marquee`{#the-marquee-element-2:the-marquee-element}](obsolete.html#the-marquee-element) element[](#the-marquee-element-2){.self-link} {#the-marquee-element-2}

``` css
@namespace "http://www.w3.org/1999/xhtml";

marquee {
  display: inline-block;
  text-align: initial;
  overflow: hidden !important;
}
```

The
[`marquee`{#the-marquee-element-2:the-marquee-element-2}](obsolete.html#the-marquee-element)
element, while [turned
on](obsolete.html#concept-marquee-on){#the-marquee-element-2:concept-marquee-on},
is expected to render in an animated fashion according to its attributes
as follows:

If the element\'s [`behavior`{#the-marquee-element-2:attr-marquee-behavior}](obsolete.html#attr-marquee-behavior) attribute is in the [scroll](obsolete.html#attr-marquee-behavior-scroll){#the-marquee-element-2:attr-marquee-behavior-scroll} state

:   Slide the contents of the element in the direction described by the
    [`direction`{#the-marquee-element-2:attr-marquee-direction}](obsolete.html#attr-marquee-direction)
    attribute as defined below, such that it begins off the start side
    of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-3}](obsolete.html#the-marquee-element),
    and ends flush with the inner end side.

    For example, if the
    [`direction`{#the-marquee-element-2:attr-marquee-direction-2}](obsolete.html#attr-marquee-direction)
    attribute is
    [left](obsolete.html#attr-marquee-direction-left){#the-marquee-element-2:attr-marquee-direction-left}
    (the default), then the contents would start such that their left
    edge are off the side of the right edge of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-4}](obsolete.html#the-marquee-element)\'s
    [content
    area](https://drafts.csswg.org/css-box/#content-area){#the-marquee-element-2:content-area
    x-internal="content-area"}, and the contents would then slide up to
    the point where the left edge of the contents are flush with the
    left inner edge of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-5}](obsolete.html#the-marquee-element)\'s
    [content
    area](https://drafts.csswg.org/css-box/#content-area){#the-marquee-element-2:content-area-2
    x-internal="content-area"}.

    Once the animation has ended, the user agent is expected to
    [increment the marquee current loop
    index](obsolete.html#increment-the-marquee-current-loop-index){#the-marquee-element-2:increment-the-marquee-current-loop-index}.
    If the element is still [turned
    on](obsolete.html#concept-marquee-on){#the-marquee-element-2:concept-marquee-on-2}
    after this, then the user agent is expected to restart the
    animation.

If the element\'s [`behavior`{#the-marquee-element-2:attr-marquee-behavior-2}](obsolete.html#attr-marquee-behavior) attribute is in the [slide](obsolete.html#attr-marquee-behavior-slide){#the-marquee-element-2:attr-marquee-behavior-slide} state

:   Slide the contents of the element in the direction described by the
    [`direction`{#the-marquee-element-2:attr-marquee-direction-3}](obsolete.html#attr-marquee-direction)
    attribute as defined below, such that it begins off the start side
    of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-6}](obsolete.html#the-marquee-element),
    and ends off the end side of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-7}](obsolete.html#the-marquee-element).

    For example, if the
    [`direction`{#the-marquee-element-2:attr-marquee-direction-4}](obsolete.html#attr-marquee-direction)
    attribute is
    [left](obsolete.html#attr-marquee-direction-left){#the-marquee-element-2:attr-marquee-direction-left-2}
    (the default), then the contents would start such that their left
    edge are off the side of the right edge of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-8}](obsolete.html#the-marquee-element)\'s
    [content
    area](https://drafts.csswg.org/css-box/#content-area){#the-marquee-element-2:content-area-3
    x-internal="content-area"}, and the contents would then slide up to
    the point where the *right* edge of the contents are flush with the
    left inner edge of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-9}](obsolete.html#the-marquee-element)\'s
    [content
    area](https://drafts.csswg.org/css-box/#content-area){#the-marquee-element-2:content-area-4
    x-internal="content-area"}.

    Once the animation has ended, the user agent is expected to
    [increment the marquee current loop
    index](obsolete.html#increment-the-marquee-current-loop-index){#the-marquee-element-2:increment-the-marquee-current-loop-index-2}.
    If the element is still [turned
    on](obsolete.html#concept-marquee-on){#the-marquee-element-2:concept-marquee-on-3}
    after this, then the user agent is expected to restart the
    animation.

If the element\'s [`behavior`{#the-marquee-element-2:attr-marquee-behavior-3}](obsolete.html#attr-marquee-behavior) attribute is in the [alternate](obsolete.html#attr-marquee-behavior-alternate){#the-marquee-element-2:attr-marquee-behavior-alternate} state

:   When the [marquee current loop
    index](obsolete.html#marquee-current-loop-index){#the-marquee-element-2:marquee-current-loop-index}
    is even (or zero), slide the contents of the element in the
    direction described by the
    [`direction`{#the-marquee-element-2:attr-marquee-direction-5}](obsolete.html#attr-marquee-direction)
    attribute as defined below, such that it begins flush with the start
    side of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-10}](obsolete.html#the-marquee-element),
    and ends flush with the end side of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-11}](obsolete.html#the-marquee-element).

    When the [marquee current loop
    index](obsolete.html#marquee-current-loop-index){#the-marquee-element-2:marquee-current-loop-index-2}
    is odd, slide the contents of the element in the opposite direction
    than that described by the
    [`direction`{#the-marquee-element-2:attr-marquee-direction-6}](obsolete.html#attr-marquee-direction)
    attribute as defined below, such that it begins flush with the end
    side of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-12}](obsolete.html#the-marquee-element),
    and ends flush with the start side of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-13}](obsolete.html#the-marquee-element).

    For example, if the
    [`direction`{#the-marquee-element-2:attr-marquee-direction-7}](obsolete.html#attr-marquee-direction)
    attribute is
    [left](obsolete.html#attr-marquee-direction-left){#the-marquee-element-2:attr-marquee-direction-left-3}
    (the default), then the contents would with their right edge flush
    with the right inner edge of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-14}](obsolete.html#the-marquee-element)\'s
    [content
    area](https://drafts.csswg.org/css-box/#content-area){#the-marquee-element-2:content-area-5
    x-internal="content-area"}, and the contents would then slide up to
    the point where the *left* edge of the contents are flush with the
    left inner edge of the
    [`marquee`{#the-marquee-element-2:the-marquee-element-15}](obsolete.html#the-marquee-element)\'s
    [content
    area](https://drafts.csswg.org/css-box/#content-area){#the-marquee-element-2:content-area-6
    x-internal="content-area"}.

    Once the animation has ended, the user agent is expected to
    [increment the marquee current loop
    index](obsolete.html#increment-the-marquee-current-loop-index){#the-marquee-element-2:increment-the-marquee-current-loop-index-3}.
    If the element is still [turned
    on](obsolete.html#concept-marquee-on){#the-marquee-element-2:concept-marquee-on-4}
    after this, then the user agent is expected to continue the
    animation.

The
[`direction`{#the-marquee-element-2:attr-marquee-direction-8}](obsolete.html#attr-marquee-direction)
attribute has the meanings described in the following table:

[`direction`{#the-marquee-element-2:attr-marquee-direction-9}](obsolete.html#attr-marquee-direction)
attribute state

Direction of animation

Start edge

End edge

Opposite direction

[left](obsolete.html#attr-marquee-direction-left){#the-marquee-element-2:attr-marquee-direction-left-4}

← Right to left

Right

Left

→ Left to Right

[right](obsolete.html#attr-marquee-direction-right){#the-marquee-element-2:attr-marquee-direction-right}

→ Left to Right

Left

Right

← Right to left

[up](obsolete.html#attr-marquee-direction-up){#the-marquee-element-2:attr-marquee-direction-up}

↑ Up (Bottom to Top)

Bottom

Top

↓ Down (Top to Bottom)

[down](obsolete.html#attr-marquee-direction-down){#the-marquee-element-2:attr-marquee-direction-down}

↓ Down (Top to Bottom)

Top

Bottom

↑ Up (Bottom to Top)

In any case, the animation should proceed such that there is a delay
given by the [marquee scroll
interval](obsolete.html#marquee-scroll-interval){#the-marquee-element-2:marquee-scroll-interval}
between each frame, and such that the content moves at most the distance
given by the [marquee scroll
distance](obsolete.html#marquee-scroll-distance){#the-marquee-element-2:marquee-scroll-distance}
with each frame.

When a
[`marquee`{#the-marquee-element-2:the-marquee-element-16}](obsolete.html#the-marquee-element)
element has a `bgcolor` attribute set, the value is expected to be
parsed using the [rules for parsing a legacy color
value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#the-marquee-element-2:rules-for-parsing-a-legacy-colour-value},
and if that does not return failure, the user agent is expected to treat
the attribute as a [presentational
hint](#presentational-hints){#the-marquee-element-2:presentational-hints}
setting the element\'s
[\'background-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-color){#the-marquee-element-2:'background-color'
x-internal="'background-color'"} property to the resulting color.

The `width` and `height` attributes on a
[`marquee`{#the-marquee-element-2:the-marquee-element-17}](obsolete.html#the-marquee-element)
element [map to the dimension
properties](#maps-to-the-dimension-property){#the-marquee-element-2:maps-to-the-dimension-property}
[\'width\'](https://drafts.csswg.org/css2/#the-width-property){#the-marquee-element-2:'width'
x-internal="'width'"} and
[\'height\'](https://drafts.csswg.org/css2/#the-height-property){#the-marquee-element-2:'height'
x-internal="'height'"} on the element respectively.

The [natural
height](https://drafts.csswg.org/css-images/#natural-height){#the-marquee-element-2:natural-height
x-internal="natural-height"} of a
[`marquee`{#the-marquee-element-2:the-marquee-element-18}](obsolete.html#the-marquee-element)
element with its
[`direction`{#the-marquee-element-2:attr-marquee-direction-10}](obsolete.html#attr-marquee-direction)
attribute in the
[up](obsolete.html#attr-marquee-direction-up){#the-marquee-element-2:attr-marquee-direction-up-2}
or
[down](obsolete.html#attr-marquee-direction-down){#the-marquee-element-2:attr-marquee-direction-down-2}
states is 200 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-marquee-element-2:'px'
x-internal="'px'"}.

The `vspace` attribute of a
[`marquee`{#the-marquee-element-2:the-marquee-element-19}](obsolete.html#the-marquee-element)
element [maps to the dimension
properties](#maps-to-the-dimension-property){#the-marquee-element-2:maps-to-the-dimension-property-2}
[\'margin-top\'](https://drafts.csswg.org/css-box/#propdef-margin-top){#the-marquee-element-2:'margin-top'
x-internal="'margin-top'"} and
[\'margin-bottom\'](https://drafts.csswg.org/css-box/#propdef-margin-bottom){#the-marquee-element-2:'margin-bottom'
x-internal="'margin-bottom'"} on the element. The `hspace` attribute of
a
[`marquee`{#the-marquee-element-2:the-marquee-element-20}](obsolete.html#the-marquee-element)
element [maps to the dimension
properties](#maps-to-the-dimension-property){#the-marquee-element-2:maps-to-the-dimension-property-3}
[\'margin-left\'](https://drafts.csswg.org/css-box/#propdef-margin-left){#the-marquee-element-2:'margin-left'
x-internal="'margin-left'"} and
[\'margin-right\'](https://drafts.csswg.org/css-box/#propdef-margin-right){#the-marquee-element-2:'margin-right'
x-internal="'margin-right'"} on the element.

#### [15.5.14]{.secno} The [`meter`{#the-meter-element-2:the-meter-element}](form-elements.html#the-meter-element) element[](#the-meter-element-2){.self-link} {#the-meter-element-2}

``` css
@namespace "http://www.w3.org/1999/xhtml";

meter { appearance: auto; }
```

The
[`meter`{#the-meter-element-2:the-meter-element-2}](form-elements.html#the-meter-element)
element is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-meter-element-2:devolvable-widget
x-internal="devolvable-widget"}. Its expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-meter-element-2:native-appearance
x-internal="native-appearance"} is to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-meter-element-2:'inline-block'
x-internal="'inline-block'"} box with a
[\'block-size\'](https://drafts.csswg.org/css-logical/#propdef-block-size){#the-meter-element-2:'block-size'
x-internal="'block-size'"} of \'1em\' and a
[\'inline-size\'](https://drafts.csswg.org/css-logical/#propdef-inline-size){#the-meter-element-2:'inline-size'
x-internal="'inline-size'"} of \'5em\', a
[\'vertical-align\'](https://drafts.csswg.org/css2/#propdef-vertical-align){#the-meter-element-2:'vertical-align'
x-internal="'vertical-align'"} of \'-0.2em\', and with its contents
depicting a gauge.

When this element has a [horizontal writing
mode](#horizontal-writing-mode){#the-meter-element-2:horizontal-writing-mode},
the depiction is expected to be of a horizontal gauge. Its minimum value
is on the right if the
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#the-meter-element-2:'direction'
x-internal="'direction'"} property has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-meter-element-2:computed-value
x-internal="computed-value"} of \'rtl\', and on the left otherwise. When
this element has a [vertical writing
mode](#vertical-writing-mode){#the-meter-element-2:vertical-writing-mode},
it is expected to depict a vertical gauge. Its minimum value is on the
bottom if the
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#the-meter-element-2:'direction'-2
x-internal="'direction'"} property has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-meter-element-2:computed-value-2
x-internal="computed-value"} of \'rtl\', and on the top otherwise.

User agents are expected to use a presentation consistent with platform
conventions for gauges, if any.

Requirements for what must be depicted in the gauge are included in the
definition of the
[`meter`{#the-meter-element-2:the-meter-element-3}](form-elements.html#the-meter-element)
element.

Need to detail the expected [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-meter-element-2:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.15]{.secno} The [`progress`{#the-progress-element-2:the-progress-element}](form-elements.html#the-progress-element) element[](#the-progress-element-2){.self-link} {#the-progress-element-2}

``` css
@namespace "http://www.w3.org/1999/xhtml";

progress { appearance: auto; }
```

The
[`progress`{#the-progress-element-2:the-progress-element-2}](form-elements.html#the-progress-element)
element is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-progress-element-2:devolvable-widget
x-internal="devolvable-widget"}. Its expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-progress-element-2:native-appearance
x-internal="native-appearance"} is to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-progress-element-2:'inline-block'
x-internal="'inline-block'"} box with a
[\'block-size\'](https://drafts.csswg.org/css-logical/#propdef-block-size){#the-progress-element-2:'block-size'
x-internal="'block-size'"} of \'1em\' and a
[\'inline-size\'](https://drafts.csswg.org/css-logical/#propdef-inline-size){#the-progress-element-2:'inline-size'
x-internal="'inline-size'"} of \'10em\', and a
[\'vertical-align\'](https://drafts.csswg.org/css2/#propdef-vertical-align){#the-progress-element-2:'vertical-align'
x-internal="'vertical-align'"} of \'-0.2em\'.

![](/images/sample-progress.png){.extra width="276" height="115"} When
the this element has a [horizontal writing
mode](#horizontal-writing-mode){#the-progress-element-2:horizontal-writing-mode},
the element is expected to be depicted as a horizontal progress bar. The
start is on the right and the end is on the left if the
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#the-progress-element-2:'direction'
x-internal="'direction'"} property on this element has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-progress-element-2:computed-value
x-internal="computed-value"} of \'rtl\', and with the start on the left
and the end on the right otherwise. When this element has a [vertical
writing
mode](#vertical-writing-mode){#the-progress-element-2:vertical-writing-mode},
it is expected to be depicted as a vertical progress bar. The start is
on the bottom and the end is on the top if the
[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction){#the-progress-element-2:'direction'-2
x-internal="'direction'"} property on this element has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-progress-element-2:computed-value-2
x-internal="computed-value"} of \'rtl\', and with the start on the top
and the end on the bottom otherwise.

User agents are expected to use a presentation consistent with platform
conventions for progress bars. In particular, user agents are expected
to use different presentations for determinate and indeterminate
progress bars. User agents are also expected to vary the presentation
based on the dimensions of the element.

Requirements for how to determine if the progress bar is determinate or
indeterminate, and what progress a determinate progress bar is to show,
are included in the definition of the
[`progress`{#the-progress-element-2:the-progress-element-3}](form-elements.html#the-progress-element)
element.

Need to detail the expected [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-progress-element-2:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.16]{.secno} The [`select`{#the-select-element-2:the-select-element}](form-elements.html#the-select-element) element[](#the-select-element-2){.self-link} {#the-select-element-2}

The
[`select`{#the-select-element-2:the-select-element-2}](form-elements.html#the-select-element)
element is an [element with default preferred
size](https://drafts.csswg.org/css-ui/#element-with-default-preferred-size){#the-select-element-2:element-with-default-preferred-size
x-internal="element-with-default-preferred-size"}, and user agents are
expected to apply the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-select-element-2:'field-sizing'
x-internal="'field-sizing'"} CSS property to
[`select`{#the-select-element-2:the-select-element-3}](form-elements.html#the-select-element)
elements.

A
[`select`{#the-select-element-2:the-select-element-4}](form-elements.html#the-select-element)
element is either a [list box]{#list-box .dfn dfn-for="select"
export=""} or a [drop-down box]{#drop-down-box .dfn dfn-for="select"
export=""}, depending on its attributes.

A
[`select`{#the-select-element-2:the-select-element-5}](form-elements.html#the-select-element)
element whose
[`multiple`{#the-select-element-2:attr-select-multiple}](form-elements.html#attr-select-multiple)
attribute is present is expected to render as a multi-select [list
box](#list-box){#the-select-element-2:list-box}.

A
[`select`{#the-select-element-2:the-select-element-6}](form-elements.html#the-select-element)
element whose
[`multiple`{#the-select-element-2:attr-select-multiple-2}](form-elements.html#attr-select-multiple)
attribute is absent, and whose [display
size](form-elements.html#concept-select-size){#the-select-element-2:concept-select-size}
is greater than 1, is expected to render as a single-select [list
box](#list-box){#the-select-element-2:list-box-2}.

When the element renders as a [list
box](#list-box){#the-select-element-2:list-box-3}, it is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-select-element-2:devolvable-widget
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-select-element-2:'inline-block'
x-internal="'inline-block'"} box. The [inline
size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-select-element-2:inline-size
x-internal="inline-size"} of its [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-select-element-2:intrinsic-size
x-internal="intrinsic-size"} is the [width of the `select`\'s
labels](#width-of-the-select's-labels){#the-select-element-2:width-of-the-select's-labels}
plus the width of a scrollbar. The [block
size](https://drafts.csswg.org/css-writing-modes/#block-size){#the-select-element-2:block-size
x-internal="block-size"} of its [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-select-element-2:intrinsic-size-2
x-internal="intrinsic-size"} is determined by the following steps:

1.  If the
    [\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-select-element-2:'field-sizing'-2
    x-internal="'field-sizing'"} property on the element has a [computed
    value](https://drafts.csswg.org/css-cascade/#computed-value){#the-select-element-2:computed-value
    x-internal="computed-value"} of
    [\'content\'](https://drafts.csswg.org/css-ui/#valdef-field-sizing-content){#the-select-element-2:field-sizing-content
    x-internal="field-sizing-content"}, return the height necessary to
    contain all rows for items.

2.  If the
    [`size`{#the-select-element-2:attr-select-size}](form-elements.html#attr-select-size)
    attribute is absent or it has no valid value, return the height
    necessary to contain four rows.

3.  Otherwise, return the height necessary to contain as many rows for
    items as given by the element\'s [display
    size](form-elements.html#concept-select-size){#the-select-element-2:concept-select-size-2}.

A
[`select`{#the-select-element-2:the-select-element-7}](form-elements.html#the-select-element)
element whose
[`multiple`{#the-select-element-2:attr-select-multiple-3}](form-elements.html#attr-select-multiple)
attribute is absent, and whose [display
size](form-elements.html#concept-select-size){#the-select-element-2:concept-select-size-3}
is 1, is expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-select-element-2:'inline-block'-2
x-internal="'inline-block'"} one-line [drop-down
box](#drop-down-box){#the-select-element-2:drop-down-box}. The [inline
size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-select-element-2:inline-size-2
x-internal="inline-size"} of its [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-select-element-2:intrinsic-size-3
x-internal="intrinsic-size"} is the [width of the `select`\'s
labels](#width-of-the-select's-labels){#the-select-element-2:width-of-the-select's-labels-2}.
If the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-select-element-2:'field-sizing'-3
x-internal="'field-sizing'"} property on the element has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-select-element-2:computed-value-2
x-internal="computed-value"} of
[\'content\'](https://drafts.csswg.org/css-ui/#valdef-field-sizing-content){#the-select-element-2:field-sizing-content-2
x-internal="field-sizing-content"}, the [inline
size](https://drafts.csswg.org/css-writing-modes/#inline-size){#the-select-element-2:inline-size-3
x-internal="inline-size"} of the [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-select-element-2:intrinsic-size-4
x-internal="intrinsic-size"} depends on the shown text. The shown text
is typically the label of an
[`option`{#the-select-element-2:the-option-element}](form-elements.html#the-option-element)
of which
[selectedness](form-elements.html#concept-option-selectedness){#the-select-element-2:concept-option-selectedness}
is set to true.

When the element renders as a [drop-down
box](#drop-down-box){#the-select-element-2:drop-down-box-2}, it is a
[devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-select-element-2:devolvable-widget-2
x-internal="devolvable-widget"}. Its appearance in the devolved state,
as well as its appearance when the [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-select-element-2:computed-value-3
x-internal="computed-value"} of the element\'s
[\'appearance\'](https://drafts.csswg.org/css-ui/#appearance-switching){#the-select-element-2:'appearance'
x-internal="'appearance'"} property is
[`'menulist-button'`{#the-select-element-2:'menulist-button'}](https://drafts.csswg.org/css-ui/#valdef-appearance-menulist-button){x-internal="'menulist-button'"},
is that of a drop-down box, including a \"drop-down button\", but not
necessarily rendered using a native control of the host operating
system. In such a state, CSS properties such as
[\'color\'](https://drafts.csswg.org/css-color/#the-color-property){#the-select-element-2:'color'
x-internal="'color'"},
[\'background-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-color){#the-select-element-2:'background-color'
x-internal="'background-color'"}, and \'border\' should not be
disregarded (as is generally permissible when rendering an element
according to its [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-select-element-2:native-appearance
x-internal="native-appearance"}).

In either case ([list box](#list-box){#the-select-element-2:list-box-4}
or [drop-down
box](#drop-down-box){#the-select-element-2:drop-down-box-3}), the
element\'s items are expected to be the element\'s [list of
options](form-elements.html#concept-select-option-list){#the-select-element-2:concept-select-option-list},
with the element\'s
[`optgroup`{#the-select-element-2:the-optgroup-element}](form-elements.html#the-optgroup-element)
element
[children](https://dom.spec.whatwg.org/#concept-tree-child){#the-select-element-2:concept-tree-child
x-internal="concept-tree-child"} providing headers for groups of options
where applicable.

An
[`optgroup`{#the-select-element-2:the-optgroup-element-2}](form-elements.html#the-optgroup-element)
element is expected to be rendered by displaying the element\'s
[`label`{#the-select-element-2:attr-optgroup-label}](form-elements.html#attr-optgroup-label)
attribute.

An
[`option`{#the-select-element-2:the-option-element-2}](form-elements.html#the-option-element)
element is expected to be rendered by displaying the element\'s
[label](form-elements.html#concept-option-label){#the-select-element-2:concept-option-label},
indented under its
[`optgroup`{#the-select-element-2:the-optgroup-element-3}](form-elements.html#the-optgroup-element)
element if it has one.

Each sequence of one or more child
[`hr`{#the-select-element-2:the-hr-element}](grouping-content.html#the-hr-element)
element siblings may be rendered as a single separator.

The [width of the `select`\'s labels]{#width-of-the-select's-labels
.dfn} is the wider of the width necessary to render the widest
[`optgroup`{#the-select-element-2:the-optgroup-element-4}](form-elements.html#the-optgroup-element),
and the width necessary to render the widest
[`option`{#the-select-element-2:the-option-element-3}](form-elements.html#the-option-element)
element in the element\'s [list of
options](form-elements.html#concept-select-option-list){#the-select-element-2:concept-select-option-list-2}
(including its indent, if any).

If a
[`select`{#the-select-element-2:the-select-element-8}](form-elements.html#the-select-element)
element contains a [placeholder label
option](form-elements.html#placeholder-label-option){#the-select-element-2:placeholder-label-option},
the user agent is expected to render that
[`option`{#the-select-element-2:the-option-element-4}](form-elements.html#the-option-element)
in a manner that conveys that it is a label, rather than a valid option
of the control. This can include preventing the [placeholder label
option](form-elements.html#placeholder-label-option){#the-select-element-2:placeholder-label-option-2}
from being explicitly selected by the user. When the [placeholder label
option](form-elements.html#placeholder-label-option){#the-select-element-2:placeholder-label-option-3}\'s
[selectedness](form-elements.html#concept-option-selectedness){#the-select-element-2:concept-option-selectedness-2}
is true, the control is expected to be displayed in a fashion that
indicates that no valid option is currently selected.

User agents are expected to render the labels in a
[`select`{#the-select-element-2:the-select-element-9}](form-elements.html#the-select-element)
in such a manner that any alignment remains consistent whether the label
is being displayed as part of the page or in a menu control.

Need to detail the expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-select-element-2:native-appearance-2
x-internal="native-appearance"} and [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-select-element-2:primitive-appearance
x-internal="primitive-appearance"}.

#### [15.5.17]{.secno} The [`textarea`{#the-textarea-element-2:the-textarea-element}](form-elements.html#the-textarea-element) element[](#the-textarea-element-2){.self-link} {#the-textarea-element-2}

The
[`textarea`{#the-textarea-element-2:the-textarea-element-2}](form-elements.html#the-textarea-element)
element is a [devolvable
widget](https://drafts.csswg.org/css-ui/#devolvable){#the-textarea-element-2:devolvable-widget
x-internal="devolvable-widget"} expected to render as an
[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block){#the-textarea-element-2:'inline-block'
x-internal="'inline-block'"} box depicting a multiline text control. If
this multiline text control provides a selection, then, when the user
changes the current selection, the user agent is expected to [queue an
element
task](webappapis.html#queue-an-element-task){#the-textarea-element-2:queue-an-element-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#the-textarea-element-2:user-interaction-task-source}
given the
[`textarea`{#the-textarea-element-2:the-textarea-element-3}](form-elements.html#the-textarea-element)
element to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#the-textarea-element-2:concept-event-fire
x-internal="concept-event-fire"} named
[`select`{#the-textarea-element-2:event-select}](indices.html#event-select)
at the element, with the
[`bubbles`{#the-textarea-element-2:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
attribute initialized to true.

The
[`textarea`{#the-textarea-element-2:the-textarea-element-4}](form-elements.html#the-textarea-element)
element is an [element with default preferred
size](https://drafts.csswg.org/css-ui/#element-with-default-preferred-size){#the-textarea-element-2:element-with-default-preferred-size
x-internal="element-with-default-preferred-size"}, and user agents are
expected to apply the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-textarea-element-2:'field-sizing'
x-internal="'field-sizing'"} CSS property to
[`textarea`{#the-textarea-element-2:the-textarea-element-5}](form-elements.html#the-textarea-element)
elements.

If the
[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing){#the-textarea-element-2:'field-sizing'-2
x-internal="'field-sizing'"} property on the element has a [computed
value](https://drafts.csswg.org/css-cascade/#computed-value){#the-textarea-element-2:computed-value
x-internal="computed-value"} of
[\'content\'](https://drafts.csswg.org/css-ui/#valdef-field-sizing-content){#the-textarea-element-2:field-sizing-content
x-internal="field-sizing-content"}, the [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-textarea-element-2:intrinsic-size
x-internal="intrinsic-size"} is determined from the text which the
element shows. The text is either a [raw
value](form-elements.html#concept-textarea-raw-value){#the-textarea-element-2:concept-textarea-raw-value}
or a short hint specified by the
[`placeholder`{#the-textarea-element-2:attr-textarea-placeholder}](form-elements.html#attr-textarea-placeholder)
attribute. User agents may take the text caret size into account in the
[intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-textarea-element-2:intrinsic-size-2
x-internal="intrinsic-size"}. Otherwise, its [intrinsic
size](https://drafts.csswg.org/css-sizing/#intrinsic-size){#the-textarea-element-2:intrinsic-size-3
x-internal="intrinsic-size"} is computed from [textarea effective
width](#textarea-effective-width){#the-textarea-element-2:textarea-effective-width}
and [textarea effective
height](#textarea-effective-height){#the-textarea-element-2:textarea-effective-height}
(as defined below).

The [textarea effective width]{#textarea-effective-width .dfn} of a
[`textarea`{#the-textarea-element-2:the-textarea-element-6}](form-elements.html#the-textarea-element)
element is `size`{.variable}×`avg`{.variable} + `sbw`{.variable}, where
`size`{.variable} is the element\'s [character
width](form-elements.html#attr-textarea-cols-value){#the-textarea-element-2:attr-textarea-cols-value},
`avg`{.variable} is the average character width of the primary font of
the element, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-textarea-element-2:'px'
x-internal="'px'"}, and `sbw`{.variable} is the width of a scrollbar, in
[CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-textarea-element-2:'px'-2
x-internal="'px'"}. (The element\'s
[\'letter-spacing\'](https://drafts.csswg.org/css-text/#letter-spacing-property){#the-textarea-element-2:'letter-spacing'
x-internal="'letter-spacing'"} property does not affect the result.)

The [textarea effective height]{#textarea-effective-height .dfn} of a
[`textarea`{#the-textarea-element-2:the-textarea-element-7}](form-elements.html#the-textarea-element)
element is the height in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-textarea-element-2:'px'-3
x-internal="'px'"} of the number of lines specified the element\'s
[character
height](form-elements.html#attr-textarea-rows-value){#the-textarea-element-2:attr-textarea-rows-value},
plus the height of a scrollbar in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-textarea-element-2:'px'-4
x-internal="'px'"}.

User agents are expected to apply the
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#the-textarea-element-2:'white-space'
x-internal="'white-space'"} CSS property to
[`textarea`{#the-textarea-element-2:the-textarea-element-8}](form-elements.html#the-textarea-element)
elements. For historical reasons, if the element has a
[`wrap`{#the-textarea-element-2:attr-textarea-wrap}](form-elements.html#attr-textarea-wrap)
attribute whose value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-textarea-element-2:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string \"`off`\",
then the user agent is expected to treat the attribute as a
[presentational
hint](#presentational-hints){#the-textarea-element-2:presentational-hints}
setting the element\'s
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#the-textarea-element-2:'white-space'-2
x-internal="'white-space'"} property to \'pre\'.

Need to detail the expected [native
appearance](https://drafts.csswg.org/css-ui/#native-appearance){#the-textarea-element-2:native-appearance
x-internal="native-appearance"} and [primitive
appearance](https://drafts.csswg.org/css-ui/#primitive-appearance){#the-textarea-element-2:primitive-appearance
x-internal="primitive-appearance"}.

### [15.6]{.secno} Frames and framesets[](#frames-and-framesets){.self-link}

User agents are expected to render
[`frameset`{#frames-and-framesets:frameset}](obsolete.html#frameset)
elements as a box with the height and width of the
[viewport](https://drafts.csswg.org/css2/#viewport){#frames-and-framesets:viewport
x-internal="viewport"}, with a surface rendered according to the
following layout algorithm:

1.  The `cols`{.variable} and `rows`{.variable} variables are lists of
    zero or more pairs consisting of a number and a unit, the unit being
    one of *percentage*, *relative*, and *absolute*.

    Use the [rules for parsing a list of
    dimensions](common-microsyntaxes.html#rules-for-parsing-a-list-of-dimensions){#frames-and-framesets:rules-for-parsing-a-list-of-dimensions}
    to parse the value of the element\'s `cols` attribute, if there is
    one. Let `cols`{.variable} be the result, or an empty list if there
    is no such attribute.

    Use the [rules for parsing a list of
    dimensions](common-microsyntaxes.html#rules-for-parsing-a-list-of-dimensions){#frames-and-framesets:rules-for-parsing-a-list-of-dimensions-2}
    to parse the value of the element\'s `rows` attribute, if there is
    one. Let `rows`{.variable} be the result, or an empty list if there
    is no such attribute.

2.  For any of the entries in `cols`{.variable} or `rows`{.variable}
    that have the number zero and the unit *relative*, change the
    entry\'s number to one.

3.  If `cols`{.variable} has no entries, then add a single entry
    consisting of the value 1 and the unit *relative* to
    `cols`{.variable}.

    If `rows`{.variable} has no entries, then add a single entry
    consisting of the value 1 and the unit *relative* to
    `rows`{.variable}.

4.  Invoke the algorithm defined below to [convert a list of dimensions
    to a list of pixel
    values](#convert-a-list-of-dimensions-to-a-list-of-pixel-values){#frames-and-framesets:convert-a-list-of-dimensions-to-a-list-of-pixel-values}
    using `cols`{.variable} as the input list, and the width of the
    surface that the
    [`frameset`{#frames-and-framesets:frameset-2}](obsolete.html#frameset)
    is being rendered into, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#frames-and-framesets:'px'
    x-internal="'px'"}, as the input dimension. Let
    `sized cols`{.variable} be the resulting list.

    Invoke the algorithm defined below to [convert a list of dimensions
    to a list of pixel
    values](#convert-a-list-of-dimensions-to-a-list-of-pixel-values){#frames-and-framesets:convert-a-list-of-dimensions-to-a-list-of-pixel-values-2}
    using `rows`{.variable} as the input list, and the height of the
    surface that the
    [`frameset`{#frames-and-framesets:frameset-3}](obsolete.html#frameset)
    is being rendered into, in [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#frames-and-framesets:'px'-2
    x-internal="'px'"}, as the input dimension. Let
    `sized rows`{.variable} be the resulting list.

5.  Split the surface into a grid of `w`{.variable}×`h`{.variable}
    rectangles, where `w`{.variable} is the number of entries in
    `sized cols`{.variable} and `h`{.variable} is the number of entries
    in `sized rows`{.variable}.

    Size the columns so that each column in the grid is as many [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#frames-and-framesets:'px'-3
    x-internal="'px'"} wide as the corresponding entry in the
    `sized cols`{.variable} list.

    Size the rows so that each row in the grid is as many [CSS
    pixels](https://drafts.csswg.org/css-values/#px){#frames-and-framesets:'px'-4
    x-internal="'px'"} high as the corresponding entry in the
    `sized rows`{.variable} list.

6.  Let `children`{.variable} be the list of
    [`frame`{#frames-and-framesets:frame}](obsolete.html#frame) and
    [`frameset`{#frames-and-framesets:frameset-4}](obsolete.html#frameset)
    elements that are
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#frames-and-framesets:concept-tree-child
    x-internal="concept-tree-child"} of the
    [`frameset`{#frames-and-framesets:frameset-5}](obsolete.html#frameset)
    element for which the algorithm was invoked.

7.  For each row of the grid of rectangles created in the previous step,
    from top to bottom, run these substeps:

    1.  For each rectangle in the row, from left to right, run these
        substeps:

        1.  If there are any elements left in `children`{.variable},
            take the first element in the list, and assign it to the
            rectangle.

            If this is a
            [`frameset`{#frames-and-framesets:frameset-6}](obsolete.html#frameset)
            element, then recurse the entire
            [`frameset`{#frames-and-framesets:frameset-7}](obsolete.html#frameset)
            layout algorithm for that
            [`frameset`{#frames-and-framesets:frameset-8}](obsolete.html#frameset)
            element, with the rectangle as the surface.

            Otherwise, it is a
            [`frame`{#frames-and-framesets:frame-2}](obsolete.html#frame)
            element; render its [content
            navigable](document-sequences.html#content-navigable){#frames-and-framesets:content-navigable},
            positioned and sized to fit the rectangle.

        2.  If there are any elements left in `children`{.variable},
            remove the first element from `children`{.variable}.

8.  If the
    [`frameset`{#frames-and-framesets:frameset-9}](obsolete.html#frameset)
    element [has a
    border](#has-a-border){#frames-and-framesets:has-a-border}, draw an
    outer set of borders around the rectangles, using the element\'s
    [frame border
    color](#frame-border-colour){#frames-and-framesets:frame-border-colour}.

    For each rectangle, if there is an element assigned to that
    rectangle, and that element [has a
    border](#has-a-border){#frames-and-framesets:has-a-border-2}, draw
    an inner set of borders around that rectangle, using the element\'s
    [frame border
    color](#frame-border-colour){#frames-and-framesets:frame-border-colour-2}.

    For each (visible) border that does not abut a rectangle that is
    assigned a
    [`frame`{#frames-and-framesets:frame-3}](obsolete.html#frame)
    element with a `noresize` attribute (including rectangles in further
    nested
    [`frameset`{#frames-and-framesets:frameset-10}](obsolete.html#frameset)
    elements), the user agent is expected to allow the user to move the
    border, resizing the rectangles within, keeping the proportions of
    any nested
    [`frameset`{#frames-and-framesets:frameset-11}](obsolete.html#frameset)
    grids.

    A
    [`frameset`{#frames-and-framesets:frameset-12}](obsolete.html#frameset)
    or [`frame`{#frames-and-framesets:frame-4}](obsolete.html#frame)
    element [has a border]{#has-a-border .dfn} if the following
    algorithm returns true:

    1.  If the element has a `frameborder` attribute whose value is not
        the empty string and whose first character is either a U+0031
        DIGIT ONE (1) character, a U+0079 LATIN SMALL LETTER Y character
        (y), or a U+0059 LATIN CAPITAL LETTER Y character (Y), then
        return true.

    2.  Otherwise, if the element has a `frameborder` attribute, return
        false.

    3.  Otherwise, if the element has a parent element that is a
        [`frameset`{#frames-and-framesets:frameset-13}](obsolete.html#frameset)
        element, then return true if *that* element [has a
        border](#has-a-border){#frames-and-framesets:has-a-border-3},
        and false if it does not.

    4.  Otherwise, return true.

    The [frame border color]{#frame-border-colour .dfn} of a
    [`frameset`{#frames-and-framesets:frameset-14}](obsolete.html#frameset)
    or [`frame`{#frames-and-framesets:frame-5}](obsolete.html#frame)
    element is the color obtained from the following algorithm:

    1.  If the element has a `bordercolor` attribute, and applying the
        [rules for parsing a legacy color
        value](common-microsyntaxes.html#rules-for-parsing-a-legacy-colour-value){#frames-and-framesets:rules-for-parsing-a-legacy-colour-value}
        to that attribute\'s value does not return failure, then return
        the color so obtained.

    2.  Otherwise, if the element has a parent element that is a
        [`frameset`{#frames-and-framesets:frameset-15}](obsolete.html#frameset)
        element, then return the [frame border
        color](#frame-border-colour){#frames-and-framesets:frame-border-colour-3}
        of that element.

    3.  Otherwise, return gray.

The algorithm to [convert a list of dimensions to a list of pixel
values]{#convert-a-list-of-dimensions-to-a-list-of-pixel-values .dfn}
consists of the following steps:

1.  Let `input list`{.variable} be the list of numbers and units passed
    to the algorithm.

    Let `output list`{.variable} be a list of numbers the same length as
    `input list`{.variable}, all zero.

    Entries in `output list`{.variable} correspond to the entries in
    `input list`{.variable} that have the same position.

2.  Let `input dimension`{.variable} be the size passed to the
    algorithm.

3.  Let `total percentage`{.variable} be the sum of all the numbers in
    `input list`{.variable} whose unit is *percentage*.

    Let `total relative`{.variable} be the sum of all the numbers in
    `input list`{.variable} whose unit is *relative*.

    Let `total absolute`{.variable} be the sum of all the numbers in
    `input list`{.variable} whose unit is *absolute*.

    Let `remaining space`{.variable} be the value of
    `input dimension`{.variable}.

4.  If `total absolute`{.variable} is greater than
    `remaining space`{.variable}, then for each entry in
    `input list`{.variable} whose unit is *absolute*, set the
    corresponding value in `output list`{.variable} to the number of the
    entry in `input list`{.variable} multiplied by
    `remaining space`{.variable} and divided by
    `total absolute`{.variable}. Then, set `remaining space`{.variable}
    to zero.

    Otherwise, for each entry in `input list`{.variable} whose unit is
    *absolute*, set the corresponding value in `output list`{.variable}
    to the number of the entry in `input list`{.variable}. Then,
    decrement `remaining space`{.variable} by
    `total absolute`{.variable}.

5.  If `total percentage`{.variable} multiplied by the
    `input dimension`{.variable} and divided by 100 is greater than
    `remaining space`{.variable}, then for each entry in
    `input list`{.variable} whose unit is *percentage*, set the
    corresponding value in `output list`{.variable} to the number of the
    entry in `input list`{.variable} multiplied by
    `remaining space`{.variable} and divided by
    `total percentage`{.variable}. Then, set
    `remaining space`{.variable} to zero.

    Otherwise, for each entry in `input list`{.variable} whose unit is
    *percentage*, set the corresponding value in
    `output list`{.variable} to the number of the entry in
    `input list`{.variable} multiplied by the
    `input dimension`{.variable} and divided by 100. Then, decrement
    `remaining space`{.variable} by `total percentage`{.variable}
    multiplied by the `input dimension`{.variable} and divided by 100.

6.  For each entry in `input list`{.variable} whose unit is *relative*,
    set the corresponding value in `output list`{.variable} to the
    number of the entry in `input list`{.variable} multiplied by
    `remaining space`{.variable} and divided by
    `total relative`{.variable}.

7.  Return `output list`{.variable}.

User agents working with integer values for frame widths (as opposed to
user agents that can lay frames out with subpixel accuracy) are expected
to distribute the remainder first to the last entry whose unit is
*relative*, then equally (not proportionally) to each entry whose unit
is *percentage*, then equally (not proportionally) to each entry whose
unit is *absolute*, and finally, failing all else, to the last entry.

------------------------------------------------------------------------

The contents of a
[`frame`{#frames-and-framesets:frame-6}](obsolete.html#frame) element
that does not have a
[`frameset`{#frames-and-framesets:frameset-16}](obsolete.html#frameset)
parent are expected to be rendered as [transparent
black](https://drafts.csswg.org/css-color/#transparent-black){#frames-and-framesets:transparent-black
x-internal="transparent-black"}; the user agent is expected to not
render its [content
navigable](document-sequences.html#content-navigable){#frames-and-framesets:content-navigable-2}
in this case, and its [content
navigable](document-sequences.html#content-navigable){#frames-and-framesets:content-navigable-3}
is expected to have a
[viewport](https://drafts.csswg.org/css2/#viewport){#frames-and-framesets:viewport-2
x-internal="viewport"} with zero width and zero height.

### [15.7]{.secno} Interactive media[](#interactive-media){.self-link}

#### [15.7.1]{.secno} Links, forms, and navigation[](#links,-forms,-and-navigation){.self-link} {#links,-forms,-and-navigation}

User agents are expected to allow the user to control aspects of
[hyperlink](links.html#hyperlink){#links,-forms,-and-navigation:hyperlink}
activation and [form
submission](form-control-infrastructure.html#form-submission-2){#links,-forms,-and-navigation:form-submission-2},
such as which
[navigable](document-sequences.html#navigable){#links,-forms,-and-navigation:navigable}
is to be used for the subsequent
[navigation](browsing-the-web.html#navigate){#links,-forms,-and-navigation:navigate}.

User agents are expected to allow users to discover the destination of
[hyperlinks](links.html#hyperlink){#links,-forms,-and-navigation:hyperlink-2}
and of
[forms](forms.html#the-form-element){#links,-forms,-and-navigation:the-form-element}
before triggering their
[navigation](browsing-the-web.html#navigate){#links,-forms,-and-navigation:navigate-2}.

User agents are expected to inform the user of whether a
[hyperlink](links.html#hyperlink){#links,-forms,-and-navigation:hyperlink-3}
includes [hyperlink
auditing](links.html#hyperlink-auditing){#links,-forms,-and-navigation:hyperlink-auditing},
and to let them know at a minimum which domains will be contacted as
part of such auditing.

User agents may allow users to
[navigate](browsing-the-web.html#navigate){#links,-forms,-and-navigation:navigate-3}
[navigables](document-sequences.html#navigable){#links,-forms,-and-navigation:navigable-2}
to the URLs
[indicated](urls-and-fetching.html#encoding-parsing-a-url){#links,-forms,-and-navigation:encoding-parsing-a-url}
by the `cite` attributes on
[`q`{#links,-forms,-and-navigation:the-q-element}](text-level-semantics.html#the-q-element),
[`blockquote`{#links,-forms,-and-navigation:the-blockquote-element}](grouping-content.html#the-blockquote-element),
[`ins`{#links,-forms,-and-navigation:the-ins-element}](edits.html#the-ins-element),
and
[`del`{#links,-forms,-and-navigation:the-del-element}](edits.html#the-del-element)
elements.

User agents may surface
[hyperlinks](links.html#hyperlink){#links,-forms,-and-navigation:hyperlink-4}
created by
[`link`{#links,-forms,-and-navigation:the-link-element}](semantics.html#the-link-element)
elements in their user interface, as discussed
[previously](semantics.html#providing-users-with-a-means-to-follow-hyperlinks-created-using-the-link-element).

#### [15.7.2]{.secno} The [`title`{#the-title-attribute-2:attr-title}](dom.html#attr-title) attribute[](#the-title-attribute-2){.self-link} {#the-title-attribute-2}

User agents are expected to expose the [advisory
information](dom.html#advisory-information){#the-title-attribute-2:advisory-information}
of elements upon user request, and to make the user aware of the
presence of such information.

On interactive graphical systems where the user can use a pointing
device, this could take the form of a tooltip. When the user is unable
to use a pointing device, then the user agent is expected to make the
content available in some other fashion, e.g. by making the element a
[focusable
area](interaction.html#focusable-area){#the-title-attribute-2:focusable-area}
and always displaying the [advisory
information](dom.html#advisory-information){#the-title-attribute-2:advisory-information-2}
of the currently
[focused](interaction.html#focused){#the-title-attribute-2:focused}
element, or by showing the [advisory
information](dom.html#advisory-information){#the-title-attribute-2:advisory-information-3}
of the elements under the user\'s finger on a touch device as the user
pans around the screen.

U+000A LINE FEED (LF) characters are expected to cause line breaks in
the tooltip; U+0009 CHARACTER TABULATION (tab) characters are expected
to render as a nonzero horizontal shift that lines up the next glyph
with the next tab stop, with tab stops occurring at points that are
multiples of 8 times the width of a U+0020 SPACE character.

::: example
For example, a visual user agent could make elements with a
[`title`{#the-title-attribute-2:attr-title-2}](dom.html#attr-title)
attribute
[focusable](interaction.html#focusable){#the-title-attribute-2:focusable},
and could make any
[focused](interaction.html#focused){#the-title-attribute-2:focused-2}
element with a
[`title`{#the-title-attribute-2:attr-title-3}](dom.html#attr-title)
attribute show its tooltip under the element while the element has
focus. This would allow a user to tab around the document to find all
the advisory text.
:::

::: example
As another example, a screen reader could provide an audio cue when
reading an element with a tooltip, with an associated key to read the
last tooltip for which a cue was played.
:::

#### [15.7.3]{.secno} Editing hosts[](#editing-hosts){.self-link}

The current text editing caret (i.e. the [active
range](https://w3c.github.io/editing/docs/execCommand/#active-range){#editing-hosts:active-range
x-internal="active-range"}, if it is empty and in an [editing
host](interaction.html#editing-host){#editing-hosts:editing-host}), if
any, is expected to act like an inline [replaced
element](https://drafts.csswg.org/css-display/#replaced-element){#editing-hosts:replaced-element
x-internal="replaced-element"} with the vertical dimensions of the caret
and with zero width for the purposes of the CSS rendering model.

This means that even an empty block can have the caret inside it, and
that when the caret is in such an element, it prevents [margins from
collapsing](https://drafts.csswg.org/css2/#collapsing-margins){#editing-hosts:collapsing-margins
x-internal="collapsing-margins"} through the element.

#### [15.7.4]{.secno} Text rendered in native user interfaces[](#text-rendered-in-native-user-interfaces){.self-link}

User agents are expected to honor the Unicode semantics of text that is
exposed in user interfaces, for example supporting the bidirectional
algorithm in text shown in dialogs, title bars, popup menus, and
tooltips. Text from the contents of elements is expected to be rendered
in a manner that honors [the
directionality](dom.html#the-directionality){#text-rendered-in-native-user-interfaces:the-directionality}
of the element from which the text was obtained. Text from attributes is
expected to be rendered in a manner that honours the [directionality of
the
attribute](dom.html#directionality-of-the-attribute){#text-rendered-in-native-user-interfaces:directionality-of-the-attribute}.

::: example
Consider the following markup, which has Hebrew text asking for a
programming language, the languages being text for which a left-to-right
direction is important given the punctuation in some of their names:

``` html
<p dir="rtl" lang="he">
 <label>
  בחר שפת תכנות:
  <select>
   <option dir="ltr">C++</option>
   <option dir="ltr">C#</option>
   <option dir="ltr">FreePascal</option>
   <option dir="ltr">F#</option>
  </select>
 </label>
</p>
```

If the
[`select`{#text-rendered-in-native-user-interfaces:the-select-element}](form-elements.html#the-select-element)
element was rendered as a drop down box, a correct rendering would
ensure that the punctuation was the same both in the drop down, and in
the box showing the current selection.

![](/images/bidiselect.png){width="206" height="105"}
:::

::: example
The directionality of attributes depends on the attribute and on the
element\'s
[`dir`{#text-rendered-in-native-user-interfaces:attr-dir}](dom.html#attr-dir)
attribute, as the following example demonstrates. Consider this markup:

``` html
<table>
 <tr>
  <th abbr="(א" dir=ltr>A
  <th abbr="(א" dir=rtl>A
  <th abbr="(א" dir=auto>A
</table>
```

If the
[`abbr`{#text-rendered-in-native-user-interfaces:attr-th-abbr}](tables.html#attr-th-abbr)
attributes are rendered, e.g. in a tooltip or other user interface, the
first will have a left parenthesis (because the direction is \'ltr\'),
the second will have a right parenthesis (because the direction is
\'rtl\'), and the third will have a right parenthesis (because the
direction is determined *from the attribute value* to be \'rtl\').

However, if instead the attribute was not a [directionality-capable
attribute](dom.html#directionality-capable-attribute){#text-rendered-in-native-user-interfaces:directionality-capable-attribute},
the results would be different:

``` html
<table>
 <tr>
  <th data-abbr="(א" dir=ltr>A
  <th data-abbr="(א" dir=rtl>A
  <th data-abbr="(א" dir=auto>A
</table>
```

In this case, if the user agent were to expose the `data-abbr` attribute
in the user interface (e.g. in a debugging environment), the last case
would be rendered with a *left* parenthesis, because the direction would
be determined from the element\'s contents.
:::

A string provided by a script (e.g. the argument to
[`window.alert()`{#text-rendered-in-native-user-interfaces:dom-alert}](timers-and-user-prompts.html#dom-alert))
is expected to be treated as an independent set of one or more
bidirectional algorithm paragraphs when displayed, as defined by the
bidirectional algorithm, including, for instance, supporting the
paragraph-breaking behavior of U+000A LINE FEED (LF) characters. For the
purposes of determining the paragraph level of such text in the
bidirectional algorithm, this specification does *not* provide a
higher-level override of rules P2 and P3.
[\[BIDI\]](references.html#refsBIDI)

When necessary, authors can enforce a particular direction for a given
paragraph by starting it with the Unicode U+200E LEFT-TO-RIGHT MARK or
U+200F RIGHT-TO-LEFT MARK characters.

::: example
Thus, the following script:

``` js
alert('\u05DC\u05DE\u05D3 HTML \u05D4\u05D9\u05D5\u05DD!')
```

\...would always result in a message reading
\"[למד LMTH היום!]{dir="rtl"}\" (not \"[דמל HTML םויה!]{dir="ltr"}\"),
regardless of the language of the user agent interface or the direction
of the page or any of its elements.
:::

::: example
For a more complex example, consider the following script:

``` bad
/* Warning: this script does not handle right-to-left scripts correctly */
var s;
if (s = prompt('What is your name?')) {
  alert(s + '! Ok, Fred, ' + s + ', and Wilma will get the car.');
}
```

When the user enters \"[Kitty]{.kbd}\", the user agent would alert
\"`Kitty! Ok, Fred, Kitty, and Wilma will get the car.`{.sample}\".
However, if the user enters \"[لا أفهم]{.kbd dir="rtl" lang="ar"}\",
then the bidirectional algorithm will determine that the direction of
the paragraph is right-to-left, and so the output will be the following
unintended mess:
\"[`لا أفهم! derF ,kO, لا أفهم, rac eht teg lliw amliW dna.`{.sample
lang=""}]{dir="rtl"}\"

To force an alert that starts with user-provided text (or other text of
unknown directionality) to render left-to-right, the string can be
prefixed with a U+200E LEFT-TO-RIGHT MARK character:

``` js
var s;
if (s = prompt('What is your name?')) {
  alert('\u200E' + s + '! Ok, Fred, ' + s + ', and Wilma will get the car.');
}
```
:::

### [15.8]{.secno} Print media[](#print-media){.self-link}

User agents are expected to allow the user to request the opportunity to
[obtain a physical form]{#obtain-a-physical-form .dfn} (or a
representation of a physical form) of a
[`Document`{#print-media:document}](dom.html#document). For example,
selecting the option to print a page or convert it to PDF format.
[\[PDF\]](references.html#refsPDF)

When the user actually [obtains a physical
form](#obtain-a-physical-form){#print-media:obtain-a-physical-form} (or
a representation of a physical form) of a
[`Document`{#print-media:document-2}](dom.html#document), the user agent
is expected to create a new rendering of the
[`Document`{#print-media:document-3}](dom.html#document) for the print
media.

### [15.9]{.secno} Unstyled XML documents[](#unstyled-xml-documents){.self-link}

HTML user agents may, in certain circumstances, find themselves
rendering non-HTML documents that use vocabularies for which they lack
any built-in knowledge. This section provides for a way for user agents
to handle such documents in a somewhat useful manner.

While a
[`Document`{#unstyled-xml-documents:document}](dom.html#document) is an
[unstyled
document](#unstyled-document){#unstyled-xml-documents:unstyled-document},
the user agent is expected to render [an unstyled document
view](#an-unstyled-document-view){#unstyled-xml-documents:an-unstyled-document-view}.

A [`Document`{#unstyled-xml-documents:document-2}](dom.html#document) is
an [unstyled document]{#unstyled-document .dfn} while it matches the
following conditions:

- The
  [`Document`{#unstyled-xml-documents:document-3}](dom.html#document)
  has no author style sheets (whether referenced by HTTP headers,
  processing instructions, elements like
  [`link`{#unstyled-xml-documents:the-link-element}](semantics.html#the-link-element),
  inline elements like
  [`style`{#unstyled-xml-documents:the-style-element}](semantics.html#the-style-element),
  or any other mechanism).
- None of the elements in the
  [`Document`{#unstyled-xml-documents:document-4}](dom.html#document)
  have any [presentational
  hints](#presentational-hints){#unstyled-xml-documents:presentational-hints}.
- None of the elements in the
  [`Document`{#unstyled-xml-documents:document-5}](dom.html#document)
  have any [style
  attributes](https://drafts.csswg.org/css-style-attr/#style-attribute){#unstyled-xml-documents:css-styling-attribute
  x-internal="css-styling-attribute"}.
- None of the elements in the
  [`Document`{#unstyled-xml-documents:document-6}](dom.html#document)
  are in any of the following namespaces: [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#unstyled-xml-documents:html-namespace-2
  x-internal="html-namespace-2"}, [SVG
  namespace](https://infra.spec.whatwg.org/#svg-namespace){#unstyled-xml-documents:svg-namespace
  x-internal="svg-namespace"}, [MathML
  namespace](https://infra.spec.whatwg.org/#mathml-namespace){#unstyled-xml-documents:mathml-namespace
  x-internal="mathml-namespace"}
- The
  [`Document`{#unstyled-xml-documents:document-7}](dom.html#document)
  has no [focusable
  area](interaction.html#focusable-area){#unstyled-xml-documents:focusable-area}
  (e.g. from XLink) other than the
  [viewport](https://drafts.csswg.org/css2/#viewport){#unstyled-xml-documents:viewport
  x-internal="viewport"}.
- The
  [`Document`{#unstyled-xml-documents:document-8}](dom.html#document)
  has no
  [hyperlinks](links.html#hyperlink){#unstyled-xml-documents:hyperlink}
  (e.g. from XLink).
- There exists no
  [script](webappapis.html#concept-script){#unstyled-xml-documents:concept-script}
  whose [settings
  object](webappapis.html#settings-object){#unstyled-xml-documents:settings-object}\'s
  [global
  object](webappapis.html#concept-settings-object-global){#unstyled-xml-documents:concept-settings-object-global}
  is a
  [`Window`{#unstyled-xml-documents:window}](nav-history-apis.html#window)
  object with this
  [`Document`{#unstyled-xml-documents:document-9}](dom.html#document) as
  its [associated
  `Document`](nav-history-apis.html#concept-document-window){#unstyled-xml-documents:concept-document-window}.
- None of the elements in the
  [`Document`{#unstyled-xml-documents:document-10}](dom.html#document)
  have any registered event listeners.

[An unstyled document view]{#an-unstyled-document-view .dfn} is one
where the DOM is not rendered according to CSS (which would, since there
are no applicable styles in this context, just result in a wall of
text), but is instead rendered in a manner that is useful for a
developer. This could consist of just showing the
[`Document`{#unstyled-xml-documents:document-11}](dom.html#document)
object\'s source, maybe with syntax highlighting, or it could consist of
displaying just the DOM tree, or simply a message saying that the page
is not a styled document.

If a
[`Document`{#unstyled-xml-documents:document-12}](dom.html#document)
stops being an [unstyled
document](#unstyled-document){#unstyled-xml-documents:unstyled-document-2},
then the conditions above stop applying, and thus a user agent following
these requirements will switch to using the regular CSS rendering.

[← 14 The XML syntax](xhtml.html) --- [Table of Contents](./) --- [16
Obsolete features →](obsolete.html)
