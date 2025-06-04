HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 15 Rendering](rendering.html) --- [Table of Contents](./) --- [17
IANA considerations →](iana.html)

1.  [[[16]{.secno} Obsolete
    features](obsolete.html#obsolete)]{#toc-obsolete}
    1.  [[16.1]{.secno} Obsolete but conforming
        features](obsolete.html#obsolete-but-conforming-features)
        1.  [[16.1.1]{.secno} Warnings for obsolete but conforming
            features](obsolete.html#warnings-for-obsolete-but-conforming-features)
    2.  [[16.2]{.secno} Non-conforming
        features](obsolete.html#non-conforming-features)
    3.  [[16.3]{.secno} Requirements for
        implementations](obsolete.html#requirements-for-implementations)
        1.  [[16.3.1]{.secno} The `marquee`
            element](obsolete.html#the-marquee-element)
        2.  [[16.3.2]{.secno} Frames](obsolete.html#frames)
        3.  [[16.3.3]{.secno} Other elements, attributes and
            APIs](obsolete.html#other-elements,-attributes-and-apis)

## [16]{.secno} Obsolete features[](#obsolete){.self-link} {#obsolete}

### [16.1]{.secno} Obsolete but conforming features[](#obsolete-but-conforming-features){.self-link}

Features listed in this section will trigger warnings in conformance
checkers.

Authors should not specify a
[`border`{#obsolete-but-conforming-features:attr-img-border}](#attr-img-border)
attribute on an
[`img`{#obsolete-but-conforming-features:the-img-element}](embedded-content.html#the-img-element)
element. If the attribute is present, its value must be the string
\"`0`\". CSS should be used instead.

Authors should not specify a
[`charset`{#obsolete-but-conforming-features:attr-script-charset}](#attr-script-charset)
attribute on a
[`script`{#obsolete-but-conforming-features:the-script-element}](scripting.html#the-script-element)
element. If the attribute is present, its value must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#obsolete-but-conforming-features:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for \"`utf-8`\". (This has no
effect in a document that conforms to the requirements elsewhere in this
standard of being encoded as
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#obsolete-but-conforming-features:utf-8
x-internal="utf-8"}.)

Authors should not specify a
[`language`{#obsolete-but-conforming-features:attr-script-language}](#attr-script-language)
attribute on a
[`script`{#obsolete-but-conforming-features:the-script-element-2}](scripting.html#the-script-element)
element. If the attribute is present, its value must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#obsolete-but-conforming-features:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for the string
\"`JavaScript`\" and either the
[`type`{#obsolete-but-conforming-features:attr-script-type}](scripting.html#attr-script-type)
attribute must be omitted or its value must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#obsolete-but-conforming-features:ascii-case-insensitive-3
x-internal="ascii-case-insensitive"} match for the string
\"`text/javascript`\". The attribute should be entirely omitted instead
(with the value \"`JavaScript`\", it has no effect), or replaced with
use of the
[`type`{#obsolete-but-conforming-features:attr-script-type-2}](scripting.html#attr-script-type)
attribute.

Authors should not specify a value for the
[`type`{#obsolete-but-conforming-features:attr-script-type-3}](scripting.html#attr-script-type)
attribute on
[`script`{#obsolete-but-conforming-features:the-script-element-3}](scripting.html#the-script-element)
elements that is the empty string or a [JavaScript MIME type essence
match](https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match){#obsolete-but-conforming-features:javascript-mime-type-essence-match
x-internal="javascript-mime-type-essence-match"}. Instead, they should
omit the attribute, which has the same effect.

Authors should not specify a
[`type`{#obsolete-but-conforming-features:attr-style-type}](#attr-style-type)
attribute on a
[`style`{#obsolete-but-conforming-features:the-style-element}](semantics.html#the-style-element)
element. If the attribute is present, its value must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#obsolete-but-conforming-features:ascii-case-insensitive-4
x-internal="ascii-case-insensitive"} match for
\"[`text/css`{#obsolete-but-conforming-features:text/css}](indices.html#text/css)\".

Authors should not specify the
[`name`{#obsolete-but-conforming-features:attr-a-name}](#attr-a-name)
attribute on
[`a`{#obsolete-but-conforming-features:the-a-element}](text-level-semantics.html#the-a-element)
elements. If the attribute is present, its value must not be the empty
string and must neither be equal to the value of any of the
[IDs](https://dom.spec.whatwg.org/#concept-id){#obsolete-but-conforming-features:concept-id
x-internal="concept-id"} in the element\'s
[tree](https://dom.spec.whatwg.org/#concept-tree){#obsolete-but-conforming-features:tree
x-internal="tree"} other than the element\'s own
[ID](https://dom.spec.whatwg.org/#concept-id){#obsolete-but-conforming-features:concept-id-2
x-internal="concept-id"}, if any, nor be equal to the value of any of
the other
[`name`{#obsolete-but-conforming-features:attr-a-name-2}](#attr-a-name)
attributes on
[`a`{#obsolete-but-conforming-features:the-a-element-2}](text-level-semantics.html#the-a-element)
elements in the element\'s
[tree](https://dom.spec.whatwg.org/#concept-tree){#obsolete-but-conforming-features:tree-2
x-internal="tree"}. If this attribute is present and the element has an
[ID](https://dom.spec.whatwg.org/#concept-id){#obsolete-but-conforming-features:concept-id-3
x-internal="concept-id"}, then the attribute\'s value must be equal to
the element\'s
[ID](https://dom.spec.whatwg.org/#concept-id){#obsolete-but-conforming-features:concept-id-4
x-internal="concept-id"}. In earlier versions of the language, this
attribute was intended as a way to specify possible targets for
[fragments](https://url.spec.whatwg.org/#concept-url-fragment){#obsolete-but-conforming-features:concept-url-fragment
x-internal="concept-url-fragment"} in
[URLs](https://url.spec.whatwg.org/#concept-url){#obsolete-but-conforming-features:url
x-internal="url"}. The
[`id`{#obsolete-but-conforming-features:the-id-attribute}](dom.html#the-id-attribute)
attribute should be used instead.

Authors should not, but may despite requirements to the contrary
elsewhere in this specification, specify the
[`maxlength`{#obsolete-but-conforming-features:attr-input-maxlength}](input.html#attr-input-maxlength)
and
[`size`{#obsolete-but-conforming-features:attr-input-size}](input.html#attr-input-size)
attributes on
[`input`{#obsolete-but-conforming-features:the-input-element}](input.html#the-input-element)
elements whose
[`type`{#obsolete-but-conforming-features:attr-input-type}](input.html#attr-input-type)
attributes are in the
[Number](input.html#number-state-(type=number)){#obsolete-but-conforming-features:number-state-(type=number)}
state. One valid reason for using these attributes regardless is to help
legacy user agents that do not support
[`input`{#obsolete-but-conforming-features:the-input-element-2}](input.html#the-input-element)
elements with `type="number"` to still render the text control with a
useful width.

#### [16.1.1]{.secno} Warnings for obsolete but conforming features[](#warnings-for-obsolete-but-conforming-features){.self-link}

To ease the transition from HTML4 Transitional documents to the language
defined in *this* specification, and to discourage certain features that
are only allowed in very few circumstances, conformance checkers must
warn the user when the following features are used in a document. These
are generally old obsolete features that have no effect, and are allowed
only to distinguish between likely mistakes (regular conformance errors)
and mere vestigial markup or unusual and discouraged practices (these
warnings).

The following features must be categorized as described above:

- The presence of a
  [`border`{#warnings-for-obsolete-but-conforming-features:attr-img-border}](#attr-img-border)
  attribute on an
  [`img`{#warnings-for-obsolete-but-conforming-features:the-img-element}](embedded-content.html#the-img-element)
  element if its value is the string \"`0`\".

- The presence of a
  [`charset`{#warnings-for-obsolete-but-conforming-features:attr-script-charset}](#attr-script-charset)
  attribute on a
  [`script`{#warnings-for-obsolete-but-conforming-features:the-script-element}](scripting.html#the-script-element)
  element if its value is an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#warnings-for-obsolete-but-conforming-features:ascii-case-insensitive
  x-internal="ascii-case-insensitive"} match for \"`utf-8`\".

- The presence of a
  [`language`{#warnings-for-obsolete-but-conforming-features:attr-script-language}](#attr-script-language)
  attribute on a
  [`script`{#warnings-for-obsolete-but-conforming-features:the-script-element-2}](scripting.html#the-script-element)
  element if its value is an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#warnings-for-obsolete-but-conforming-features:ascii-case-insensitive-2
  x-internal="ascii-case-insensitive"} match for the string
  \"`JavaScript`\" and if there is no
  [`type`{#warnings-for-obsolete-but-conforming-features:attr-script-type}](scripting.html#attr-script-type)
  attribute or there is and its value is an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#warnings-for-obsolete-but-conforming-features:ascii-case-insensitive-3
  x-internal="ascii-case-insensitive"} match for the string
  \"`text/javascript`\".

- The presence of a
  [`type`{#warnings-for-obsolete-but-conforming-features:attr-style-type}](#attr-style-type)
  attribute on a
  [`script`{#warnings-for-obsolete-but-conforming-features:the-script-element-3}](scripting.html#the-script-element)
  element if its value is a [JavaScript MIME type essence
  match](https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match){#warnings-for-obsolete-but-conforming-features:javascript-mime-type-essence-match
  x-internal="javascript-mime-type-essence-match"}.

- The presence of a
  [`type`{#warnings-for-obsolete-but-conforming-features:attr-style-type-2}](#attr-style-type)
  attribute on a
  [`style`{#warnings-for-obsolete-but-conforming-features:the-style-element}](semantics.html#the-style-element)
  element if its value is an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#warnings-for-obsolete-but-conforming-features:ascii-case-insensitive-4
  x-internal="ascii-case-insensitive"} match for
  \"[`text/css`{#warnings-for-obsolete-but-conforming-features:text/css}](indices.html#text/css)\".

- The presence of a
  [`name`{#warnings-for-obsolete-but-conforming-features:attr-a-name}](#attr-a-name)
  attribute on an
  [`a`{#warnings-for-obsolete-but-conforming-features:the-a-element}](text-level-semantics.html#the-a-element)
  element, if its value is not the empty string.

- The presence of a
  [`maxlength`{#warnings-for-obsolete-but-conforming-features:attr-input-maxlength}](input.html#attr-input-maxlength)
  attribute on an
  [`input`{#warnings-for-obsolete-but-conforming-features:the-input-element}](input.html#the-input-element)
  element whose
  [`type`{#warnings-for-obsolete-but-conforming-features:attr-input-type}](input.html#attr-input-type)
  attribute is in the
  [Number](input.html#number-state-(type=number)){#warnings-for-obsolete-but-conforming-features:number-state-(type=number)}
  state.

- The presence of a
  [`size`{#warnings-for-obsolete-but-conforming-features:attr-input-size}](input.html#attr-input-size)
  attribute on an
  [`input`{#warnings-for-obsolete-but-conforming-features:the-input-element-2}](input.html#the-input-element)
  element whose
  [`type`{#warnings-for-obsolete-but-conforming-features:attr-input-type-2}](input.html#attr-input-type)
  attribute is in the
  [Number](input.html#number-state-(type=number)){#warnings-for-obsolete-but-conforming-features:number-state-(type=number)-2}
  state.

Conformance checkers must distinguish between pages that have no
conformance errors and have none of these obsolete features, and pages
that have no conformance errors but do have some of these obsolete
features.

For example, a validator could report some pages as \"Valid HTML\" and
others as \"Valid HTML with warnings\".

### [16.2]{.secno} Non-conforming features[](#non-conforming-features){.self-link}

Elements in the following list are entirely obsolete, and must not be
used by authors:

[`applet`]{#applet .dfn dfn-type="element"}
:   Use
    [`embed`{#non-conforming-features:the-embed-element}](iframe-embed-object.html#the-embed-element)
    or
    [`object`{#non-conforming-features:the-object-element}](iframe-embed-object.html#the-object-element)
    instead.

[`acronym`]{#acronym .dfn dfn-type="element"}
:   Use
    [`abbr`{#non-conforming-features:the-abbr-element}](text-level-semantics.html#the-abbr-element)
    instead.

[`bgsound`]{#bgsound .dfn dfn-type="element"}
:   Use
    [`audio`{#non-conforming-features:the-audio-element}](media.html#the-audio-element)
    instead.

[`dir`]{#dir .dfn dfn-type="element"}
:   Use
    [`ul`{#non-conforming-features:the-ul-element}](grouping-content.html#the-ul-element)
    instead.

[`frame`{#non-conforming-features:frame}](#frame)\
[`frameset`{#non-conforming-features:frameset}](#frameset)\
[`noframes`]{#noframes .dfn dfn-type="element"}
:   Either use
    [`iframe`{#non-conforming-features:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
    and CSS instead, or use server-side includes to generate complete
    pages with the various invariant parts merged in.

[`isindex`]{#isindex .dfn dfn-type="element"}
:   Use an explicit
    [`form`{#non-conforming-features:the-form-element}](forms.html#the-form-element)
    and [text
    control](input.html#text-(type=text)-state-and-search-state-(type=search)){#non-conforming-features:text-(type=text)-state-and-search-state-(type=search)}
    combination instead.

[`keygen`]{#keygen .dfn dfn-type="element"}

:   For enterprise device management use cases, use native on-device
    management capabilities.

    For certificate enrollment use cases, use the Web Cryptography API
    to generate a keypair for the certificate, and then export the
    certificate and key to allow the user to install them manually.
    [\[WEBCRYPTO\]](references.html#refsWEBCRYPTO)

[`listing`]{#listing .dfn dfn-type="element"}
:   Use
    [`pre`{#non-conforming-features:the-pre-element}](grouping-content.html#the-pre-element)
    and
    [`code`{#non-conforming-features:the-code-element}](text-level-semantics.html#the-code-element)
    instead.

[`menuitem`]{#menuitem .dfn dfn-type="element"}
:   To implement a custom context menu, use script to handle the
    [`contextmenu`{#non-conforming-features:event-contextmenu}](https://w3c.github.io/uievents/#event-type-contextmenu){x-internal="event-contextmenu"}
    event.

[`nextid`]{#nextid .dfn dfn-type="element"}
:   Use GUIDs instead.

[`noembed`]{#noembed .dfn dfn-type="element"}
:   Use
    [`object`{#non-conforming-features:the-object-element-2}](iframe-embed-object.html#the-object-element)
    instead of
    [`embed`{#non-conforming-features:the-embed-element-2}](iframe-embed-object.html#the-embed-element)
    when fallback is necessary.

[`param`]{#param .dfn dfn-type="element"}
:   Use the
    [`data`{#non-conforming-features:attr-object-data}](iframe-embed-object.html#attr-object-data)
    attribute of the
    [`object`{#non-conforming-features:the-object-element-3}](iframe-embed-object.html#the-object-element)
    element to set the URL of the external resource.

[`plaintext`]{#plaintext .dfn dfn-type="element"}
:   Use the
    \"[`text/plain`{#non-conforming-features:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\"
    [MIME
    type](https://mimesniff.spec.whatwg.org/#mime-type){#non-conforming-features:mime-type
    x-internal="mime-type"} instead.

[`rb`]{#rb .dfn dfn-type="element"}\
[`rtc`]{#rtc .dfn dfn-type="element"}
:   Providing the ruby base directly inside the
    [`ruby`{#non-conforming-features:the-ruby-element}](text-level-semantics.html#the-ruby-element)
    element or using nested
    [`ruby`{#non-conforming-features:the-ruby-element-2}](text-level-semantics.html#the-ruby-element)
    elements is sufficient.

[`strike`]{#strike .dfn dfn-type="element"}
:   Use
    [`del`{#non-conforming-features:the-del-element}](edits.html#the-del-element)
    instead if the element is marking an edit, otherwise use
    [`s`{#non-conforming-features:the-s-element}](text-level-semantics.html#the-s-element)
    instead.

[`xmp`]{#xmp .dfn dfn-type="element"}
:   Use
    [`pre`{#non-conforming-features:the-pre-element-2}](grouping-content.html#the-pre-element)
    and
    [`code`{#non-conforming-features:the-code-element-2}](text-level-semantics.html#the-code-element)
    instead, and escape \"`<`\" and \"`&`\" characters as \"`&lt;`\" and
    \"`&amp;`\" respectively.

[`basefont`]{#basefont .dfn dfn-type="element"}\
[`big`]{#big .dfn dfn-type="element"}\
[`blink`]{#blink .dfn dfn-type="element"}\
[`center`]{#center .dfn dfn-type="element"}\
[`font`]{#font .dfn dfn-type="element"}\
[`marquee`{#non-conforming-features:the-marquee-element}](#the-marquee-element)\
[`multicol`]{#multicol .dfn dfn-type="element"}\
[`nobr`]{#nobr .dfn dfn-type="element"}\
[`spacer`]{#spacer .dfn dfn-type="element"}\
[`tt`]{#tt .dfn dfn-type="element"}

:   Use appropriate elements or CSS instead.

    Where the [`tt`{#non-conforming-features:tt}](#tt) element would
    have been used for marking up keyboard input, consider the
    [`kbd`{#non-conforming-features:the-kbd-element}](text-level-semantics.html#the-kbd-element)
    element; for variables, consider the
    [`var`{#non-conforming-features:the-var-element}](text-level-semantics.html#the-var-element)
    element; for computer code, consider the
    [`code`{#non-conforming-features:the-code-element-3}](text-level-semantics.html#the-code-element)
    element; and for computer output, consider the
    [`samp`{#non-conforming-features:the-samp-element}](text-level-semantics.html#the-samp-element)
    element.

    Similarly, if the [`big`{#non-conforming-features:big}](#big)
    element is being used to denote a heading, consider using the
    [`h1`{#non-conforming-features:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
    element; if it is being used for marking up important passages,
    consider the
    [`strong`{#non-conforming-features:the-strong-element}](text-level-semantics.html#the-strong-element)
    element; and if it is being used for highlighting text for reference
    purposes, consider the
    [`mark`{#non-conforming-features:the-mark-element}](text-level-semantics.html#the-mark-element)
    element.

    See also the [text-level semantics usage
    summary](text-level-semantics.html#usage-summary) for more
    suggestions with examples.

------------------------------------------------------------------------

The following attributes are obsolete (though the elements are still
part of the language), and must not be used by authors:

[`charset`]{#attr-a-charset .dfn dfn-for="a" dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element}](text-level-semantics.html#the-a-element) elements\
[`charset`]{#attr-link-charset .dfn dfn-for="link" dfn-type="element-attr"} on [`link`{#non-conforming-features:the-link-element}](semantics.html#the-link-element) elements
:   Use an HTTP
    \`[`Content-Type`{#non-conforming-features:content-type}](urls-and-fetching.html#content-type)\`
    header on the linked resource instead.

[`charset`]{#attr-script-charset .dfn dfn-for="script" dfn-type="element-attr"} on [`script`{#non-conforming-features:the-script-element}](scripting.html#the-script-element) elements (except as noted in the previous section)
:   Omit the attribute. Both documents and scripts are required to use
    [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#non-conforming-features:utf-8
    x-internal="utf-8"}, so it is redundant to specify it on the
    [`script`{#non-conforming-features:the-script-element-2}](scripting.html#the-script-element)
    element since it inherits from the document.

[`coords`]{#attr-a-coords .dfn dfn-for="a" dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-2}](text-level-semantics.html#the-a-element) elements\
[`shape`]{#attr-a-shape .dfn dfn-for="a" dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-3}](text-level-semantics.html#the-a-element) elements
:   Use
    [`area`{#non-conforming-features:the-area-element}](image-maps.html#the-area-element)
    instead of
    [`a`{#non-conforming-features:the-a-element-4}](text-level-semantics.html#the-a-element)
    for image maps.

[`methods`]{#attr-a-methods .dfn dfn-for="a" dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-5}](text-level-semantics.html#the-a-element) elements\
[`methods`]{#attr-link-methods .dfn dfn-for="link" dfn-type="element-attr"} on [`link`{#non-conforming-features:the-link-element-2}](semantics.html#the-link-element) elements
:   Use the HTTP OPTIONS feature instead.

[`name`]{#attr-a-name .dfn dfn-for="a" dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-6}](text-level-semantics.html#the-a-element) elements (except as noted in the previous section)\
[`name`]{#attr-embed-name .dfn dfn-for="embed" dfn-type="element-attr"} on [`embed`{#non-conforming-features:the-embed-element-3}](iframe-embed-object.html#the-embed-element) elements\
[`name`]{#attr-img-name .dfn dfn-for="img" dfn-type="element-attr"} on [`img`{#non-conforming-features:the-img-element}](embedded-content.html#the-img-element) elements\
[`name`]{#attr-option-name .dfn dfn-for="option" dfn-type="element-attr"} on [`option`{#non-conforming-features:the-option-element}](form-elements.html#the-option-element) elements
:   Use the
    [`id`{#non-conforming-features:the-id-attribute}](dom.html#the-id-attribute)
    attribute instead.

[`rev`]{#attr-a-rev .dfn dfn-for="a" dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-7}](text-level-semantics.html#the-a-element) elements\
[`rev`]{#attr-link-rev .dfn dfn-for="link" dfn-type="element-attr"} on [`link`{#non-conforming-features:the-link-element-3}](semantics.html#the-link-element) elements
:   Use the
    [`rel`{#non-conforming-features:attr-hyperlink-rel}](links.html#attr-hyperlink-rel)
    attribute instead, with an opposite term. (For example, instead of
    `rev="made"`, use `rel="author"`.)

[`urn`]{#attr-a-urn .dfn dfn-for="a" dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-8}](text-level-semantics.html#the-a-element) elements\
[`urn`]{#attr-link-urn .dfn dfn-for="link" dfn-type="element-attr"} on [`link`{#non-conforming-features:the-link-element-4}](semantics.html#the-link-element) elements
:   Specify the preferred persistent identifier using the
    [`href`{#non-conforming-features:attr-hyperlink-href}](links.html#attr-hyperlink-href)
    attribute instead.

[`accept`]{#attr-form-accept .dfn dfn-for="form" dfn-type="element-attr"} on [`form`{#non-conforming-features:the-form-element-2}](forms.html#the-form-element) elements
:   Use the
    [`accept`{#non-conforming-features:attr-input-accept}](input.html#attr-input-accept)
    attribute directly on the
    [`input`{#non-conforming-features:the-input-element}](input.html#the-input-element)
    elements instead.

[`hreflang`]{#attr-area-hreflang .dfn dfn-for="area" dfn-type="element-attr"} on [`area`{#non-conforming-features:the-area-element-2}](image-maps.html#the-area-element) elements\
[`type`]{#attr-area-type .dfn dfn-for="area" dfn-type="element-attr"} on [`area`{#non-conforming-features:the-area-element-3}](image-maps.html#the-area-element) elements
:   These attributes do not do anything useful, and for historical
    reasons there are no corresponding IDL attributes on
    [`area`{#non-conforming-features:the-area-element-4}](image-maps.html#the-area-element)
    elements. Omit them altogether.

[`nohref`]{#attr-area-nohref .dfn dfn-for="area" dfn-type="element-attr"} on [`area`{#non-conforming-features:the-area-element-5}](image-maps.html#the-area-element) elements
:   Omitting the
    [`href`{#non-conforming-features:attr-hyperlink-href-2}](links.html#attr-hyperlink-href)
    attribute is sufficient; the
    [`nohref`{#non-conforming-features:attr-area-nohref}](#attr-area-nohref)
    attribute is unnecessary. Omit it altogether.

[`profile`]{#attr-head-profile .dfn dfn-for="head" dfn-type="element-attr"} on [`head`{#non-conforming-features:the-head-element}](semantics.html#the-head-element) elements
:   Unnecessary. Omit it altogether.

[`manifest`]{#attr-html-manifest .dfn dfn-for="html" dfn-type="element-attr"} on [`html`{#non-conforming-features:the-html-element}](semantics.html#the-html-element) elements
:   Use service workers instead. [\[SW\]](references.html#refsSW)

[`version`]{#attr-html-version .dfn dfn-for="html" dfn-type="element-attr"} on [`html`{#non-conforming-features:the-html-element-2}](semantics.html#the-html-element) elements
:   Unnecessary. Omit it altogether.

[`ismap`]{#attr-input-ismap .dfn dfn-for="input" dfn-type="element-attr"} on [`input`{#non-conforming-features:the-input-element-2}](input.html#the-input-element) elements
:   Unnecessary. Omit it altogether. All
    [`input`{#non-conforming-features:the-input-element-3}](input.html#the-input-element)
    elements with a
    [`type`{#non-conforming-features:attr-input-type}](input.html#attr-input-type)
    attribute in the [Image
    Button](input.html#image-button-state-(type=image)){#non-conforming-features:image-button-state-(type=image)}
    state are processed as server-side image maps.

[`usemap`]{#attr-input-usemap .dfn dfn-for="input" dfn-type="element-attr"} on [`input`{#non-conforming-features:the-input-element-4}](input.html#the-input-element) elements\
[`usemap`]{#attr-object-usemap .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-4}](iframe-embed-object.html#the-object-element) elements
:   Use the
    [`img`{#non-conforming-features:the-img-element-2}](embedded-content.html#the-img-element)
    element for image maps.

[`longdesc`]{#attr-iframe-longdesc .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element) elements\
[`longdesc`]{#attr-img-longdesc .dfn dfn-for="img" dfn-type="element-attr"} on [`img`{#non-conforming-features:the-img-element-3}](embedded-content.html#the-img-element) elements
:   Use a regular
    [`a`{#non-conforming-features:the-a-element-9}](text-level-semantics.html#the-a-element)
    element to link to the description, or (in the case of images) use
    an [image
    map](image-maps.html#image-map){#non-conforming-features:image-map}
    to provide a link from the image to the image\'s description.

[`lowsrc`]{#attr-img-lowsrc .dfn dfn-for="img" dfn-type="element-attr"} on [`img`{#non-conforming-features:the-img-element-4}](embedded-content.html#the-img-element) elements
:   Use a progressive JPEG image (given in the
    [`src`{#non-conforming-features:attr-img-src}](embedded-content.html#attr-img-src)
    attribute), instead of using two separate images.

[`target`]{#attr-link-target .dfn dfn-for="link" dfn-type="element-attr"} on [`link`{#non-conforming-features:the-link-element-5}](semantics.html#the-link-element) elements
:   Unnecessary. Omit it altogether.

[`type`]{#attr-menu-type .dfn dfn-for="menu" dfn-type="element-attr"} on [`menu`{#non-conforming-features:the-menu-element}](grouping-content.html#the-menu-element) elements
:   To implement a custom context menu, use script to handle the
    [`contextmenu`{#non-conforming-features:event-contextmenu-2}](https://w3c.github.io/uievents/#event-type-contextmenu){x-internal="event-contextmenu"}
    event. For toolbar menus, omit the attribute.

[`label`]{#attr-menu-label .dfn dfn-for="menu" dfn-type="element-attr"} on [`menu`{#non-conforming-features:the-menu-element-2}](grouping-content.html#the-menu-element) elements\
[`contextmenu`]{#attr-contextmenu .dfn dfn-type="element-attr"} on all elements\
[`onshow`]{#handler-onshow .dfn dfn-type="element-attr"} on all elements
:   To implement a custom context menu, use script to handle the
    [`contextmenu`{#non-conforming-features:event-contextmenu-3}](https://w3c.github.io/uievents/#event-type-contextmenu){x-internal="event-contextmenu"}
    event.

[`scheme`]{#attr-meta-scheme .dfn dfn-for="meta" dfn-type="element-attr"} on [`meta`{#non-conforming-features:the-meta-element}](semantics.html#the-meta-element) elements
:   Use only one scheme per field, or make the scheme declaration part
    of the value.

[`archive`]{#attr-object-archive .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-5}](iframe-embed-object.html#the-object-element) elements\
[`classid`]{#attr-object-classid .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-6}](iframe-embed-object.html#the-object-element) elements\
[`code`]{#attr-object-code .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-7}](iframe-embed-object.html#the-object-element) elements\
[`codebase`]{#attr-object-codebase .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-8}](iframe-embed-object.html#the-object-element) elements\
[`codetype`]{#attr-object-codetype .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-9}](iframe-embed-object.html#the-object-element) elements
:   Use the
    [`data`{#non-conforming-features:attr-object-data-2}](iframe-embed-object.html#attr-object-data)
    and
    [`type`{#non-conforming-features:attr-object-type}](iframe-embed-object.html#attr-object-type)
    attributes to invoke
    [plugins](infrastructure.html#plugin){#non-conforming-features:plugin}.

[`declare`]{#attr-object-declare .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-10}](iframe-embed-object.html#the-object-element) elements
:   Repeat the
    [`object`{#non-conforming-features:the-object-element-11}](iframe-embed-object.html#the-object-element)
    element completely each time the resource is to be reused.

[`standby`]{#attr-object-standby .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-12}](iframe-embed-object.html#the-object-element) elements
:   Optimize the linked resource so that it loads quickly or, at least,
    incrementally.

[`typemustmatch`]{#attr-object-typemustmatch .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-13}](iframe-embed-object.html#the-object-element) elements
:   Avoid using
    [`object`{#non-conforming-features:the-object-element-14}](iframe-embed-object.html#the-object-element)
    elements with untrusted resources.

[`language`]{#attr-script-language .dfn dfn-for="script" dfn-type="element-attr"} on [`script`{#non-conforming-features:the-script-element-3}](scripting.html#the-script-element) elements (except as noted in the previous section)
:   Omit the attribute for JavaScript; for [data
    blocks](scripting.html#data-block){#non-conforming-features:data-block},
    use the
    [`type`{#non-conforming-features:attr-script-type}](scripting.html#attr-script-type)
    attribute instead.

[`event`]{#attr-script-event .dfn dfn-for="script" dfn-type="element-attr"} on [`script`{#non-conforming-features:the-script-element-4}](scripting.html#the-script-element) elements\
[`for`]{#attr-script-for .dfn dfn-for="script" dfn-type="element-attr"} on [`script`{#non-conforming-features:the-script-element-5}](scripting.html#the-script-element) elements
:   Use DOM events mechanisms to register event listeners.
    [\[DOM\]](references.html#refsDOM)

[`type`]{#attr-style-type .dfn dfn-for="style" dfn-type="element-attr"} on [`style`{#non-conforming-features:the-style-element}](semantics.html#the-style-element) elements (except as noted in the previous section)
:   Omit the attribute for CSS; for [data
    blocks](scripting.html#data-block){#non-conforming-features:data-block-2},
    use
    [`script`{#non-conforming-features:the-script-element-6}](scripting.html#the-script-element)
    as the container instead of
    [`style`{#non-conforming-features:the-style-element-2}](semantics.html#the-style-element).

[`datapagesize`]{#attr-table-datapagesize .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element}](tables.html#the-table-element) elements
:   Unnecessary. Omit it altogether.

[`summary`]{#attr-table-summary .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-2}](tables.html#the-table-element) elements
:   Use one of the [techniques for describing
    tables](tables.html#table-descriptions-techniques) given in the
    [`table`{#non-conforming-features:the-table-element-3}](tables.html#the-table-element)
    section instead.

[`abbr`]{#attr-td-abbr .dfn dfn-for="td" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element}](tables.html#the-td-element) elements

:   Use text that begins in an unambiguous and terse manner, and include
    any more elaborate text after that. The
    [`title`{#non-conforming-features:attr-title}](dom.html#attr-title)
    attribute can also be useful in including more detailed text, so
    that the cell\'s contents can be made terse. If it\'s a heading, use
    [`th`{#non-conforming-features:the-th-element}](tables.html#the-th-element)
    (which has an
    [`abbr`{#non-conforming-features:attr-th-abbr}](tables.html#attr-th-abbr)
    attribute).

[`axis`]{#attr-tdth-axis .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-2}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-2}](tables.html#the-th-element) elements

:   Use the
    [`scope`{#non-conforming-features:attr-th-scope}](tables.html#attr-th-scope)
    attribute on the relevant
    [`th`{#non-conforming-features:the-th-element-3}](tables.html#the-th-element).

[`scope`]{#attr-td-scope .dfn dfn-for="td" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-3}](tables.html#the-td-element) elements

:   Use
    [`th`{#non-conforming-features:the-th-element-4}](tables.html#the-th-element)
    elements for heading cells.

[`datasrc`]{#attr-datasrc .dfn dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-10}](text-level-semantics.html#the-a-element), [`button`{#non-conforming-features:the-button-element}](form-elements.html#the-button-element), [`div`{#non-conforming-features:the-div-element}](grouping-content.html#the-div-element), [`frame`{#non-conforming-features:frame-2}](#frame), [`iframe`{#non-conforming-features:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element), [`img`{#non-conforming-features:the-img-element-5}](embedded-content.html#the-img-element), [`input`{#non-conforming-features:the-input-element-5}](input.html#the-input-element), [`label`{#non-conforming-features:the-label-element}](forms.html#the-label-element), [`legend`{#non-conforming-features:the-legend-element}](form-elements.html#the-legend-element), [`marquee`{#non-conforming-features:the-marquee-element-2}](#the-marquee-element), [`object`{#non-conforming-features:the-object-element-15}](iframe-embed-object.html#the-object-element), [`option`{#non-conforming-features:the-option-element-2}](form-elements.html#the-option-element), [`select`{#non-conforming-features:the-select-element}](form-elements.html#the-select-element), [`span`{#non-conforming-features:the-span-element}](text-level-semantics.html#the-span-element), [`table`{#non-conforming-features:the-table-element-4}](tables.html#the-table-element), and [`textarea`{#non-conforming-features:the-textarea-element}](form-elements.html#the-textarea-element) elements\
[`datafld`]{#attr-datafld .dfn dfn-type="element-attr"} on [`a`{#non-conforming-features:the-a-element-11}](text-level-semantics.html#the-a-element), [`button`{#non-conforming-features:the-button-element-2}](form-elements.html#the-button-element), [`div`{#non-conforming-features:the-div-element-2}](grouping-content.html#the-div-element), [`fieldset`{#non-conforming-features:the-fieldset-element}](form-elements.html#the-fieldset-element), [`frame`{#non-conforming-features:frame-3}](#frame), [`iframe`{#non-conforming-features:the-iframe-element-4}](iframe-embed-object.html#the-iframe-element), [`img`{#non-conforming-features:the-img-element-6}](embedded-content.html#the-img-element), [`input`{#non-conforming-features:the-input-element-6}](input.html#the-input-element), [`label`{#non-conforming-features:the-label-element-2}](forms.html#the-label-element), [`legend`{#non-conforming-features:the-legend-element-2}](form-elements.html#the-legend-element), [`marquee`{#non-conforming-features:the-marquee-element-3}](#the-marquee-element), [`object`{#non-conforming-features:the-object-element-16}](iframe-embed-object.html#the-object-element), [`select`{#non-conforming-features:the-select-element-2}](form-elements.html#the-select-element), [`span`{#non-conforming-features:the-span-element-2}](text-level-semantics.html#the-span-element), and [`textarea`{#non-conforming-features:the-textarea-element-2}](form-elements.html#the-textarea-element) elements\
[`dataformatas`]{#attr-dataformatas .dfn dfn-type="element-attr"} on [`button`{#non-conforming-features:the-button-element-3}](form-elements.html#the-button-element), [`div`{#non-conforming-features:the-div-element-3}](grouping-content.html#the-div-element), [`input`{#non-conforming-features:the-input-element-7}](input.html#the-input-element), [`label`{#non-conforming-features:the-label-element-3}](forms.html#the-label-element), [`legend`{#non-conforming-features:the-legend-element-3}](form-elements.html#the-legend-element), [`marquee`{#non-conforming-features:the-marquee-element-4}](#the-marquee-element), [`object`{#non-conforming-features:the-object-element-17}](iframe-embed-object.html#the-object-element), [`option`{#non-conforming-features:the-option-element-3}](form-elements.html#the-option-element), [`select`{#non-conforming-features:the-select-element-3}](form-elements.html#the-select-element), [`span`{#non-conforming-features:the-span-element-3}](text-level-semantics.html#the-span-element), and [`table`{#non-conforming-features:the-table-element-5}](tables.html#the-table-element) elements
:   Use script and a mechanism such as
    [`XMLHttpRequest`{#non-conforming-features:xmlhttprequest}](https://xhr.spec.whatwg.org/#xmlhttprequest){x-internal="xmlhttprequest"}
    to populate the page dynamically. [\[XHR\]](references.html#refsXHR)

[`dropzone`]{#attr-dropzone .dfn dfn-type="element-attr"} on all elements
:   Use script to handle the
    [`dragenter`{#non-conforming-features:event-dnd-dragenter}](dnd.html#event-dnd-dragenter)
    and
    [`dragover`{#non-conforming-features:event-dnd-dragover}](dnd.html#event-dnd-dragover)
    events instead.

[`alink`]{#attr-body-alink .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element}](sections.html#the-body-element) elements\
[`bgcolor`]{#attr-body-bgcolor .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-2}](sections.html#the-body-element) elements\
[`bottommargin`]{#attr-body-bottommargin .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-3}](sections.html#the-body-element) elements\
[`leftmargin`]{#attr-body-leftmargin .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-4}](sections.html#the-body-element) elements\
[`link`]{#attr-body-link .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-5}](sections.html#the-body-element) elements\
[`marginheight`]{#attr-body-marginheight .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-6}](sections.html#the-body-element) elements\
[`marginwidth`]{#attr-body-marginwidth .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-7}](sections.html#the-body-element) elements\
[`rightmargin`]{#attr-body-rightmargin .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-8}](sections.html#the-body-element) elements\
[`text`]{#attr-body-text .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-9}](sections.html#the-body-element) elements\
[`topmargin`]{#attr-body-topmargin .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-10}](sections.html#the-body-element) elements\
[`vlink`]{#attr-body-vlink .dfn dfn-for="body" dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-11}](sections.html#the-body-element) elements\
[`clear`]{#attr-br-clear .dfn dfn-for="br" dfn-type="element-attr"} on [`br`{#non-conforming-features:the-br-element}](text-level-semantics.html#the-br-element) elements\
[`align`]{#attr-caption-align .dfn dfn-for="caption" dfn-type="element-attr"} on [`caption`{#non-conforming-features:the-caption-element}](tables.html#the-caption-element) elements\
[`align`]{#attr-col-align .dfn dfn-for="col" dfn-type="element-attr"} on [`col`{#non-conforming-features:the-col-element}](tables.html#the-col-element) elements\
[`char`]{#attr-col-char .dfn dfn-for="col" dfn-type="element-attr"} on [`col`{#non-conforming-features:the-col-element-2}](tables.html#the-col-element) elements\
[`charoff`]{#attr-col-charoff .dfn dfn-for="col" dfn-type="element-attr"} on [`col`{#non-conforming-features:the-col-element-3}](tables.html#the-col-element) elements\
[`valign`]{#attr-col-valign .dfn dfn-for="col" dfn-type="element-attr"} on [`col`{#non-conforming-features:the-col-element-4}](tables.html#the-col-element) elements\
[`width`]{#attr-col-width .dfn dfn-for="col" dfn-type="element-attr"} on [`col`{#non-conforming-features:the-col-element-5}](tables.html#the-col-element) elements\
[`align`]{#attr-div-align .dfn dfn-for="div" dfn-type="element-attr"} on [`div`{#non-conforming-features:the-div-element-4}](grouping-content.html#the-div-element) elements\
[`compact`]{#attr-dl-compact .dfn dfn-for="dl" dfn-type="element-attr"} on [`dl`{#non-conforming-features:the-dl-element}](grouping-content.html#the-dl-element) elements\
[`align`]{#attr-embed-align .dfn dfn-for="embed" dfn-type="element-attr"} on [`embed`{#non-conforming-features:the-embed-element-4}](iframe-embed-object.html#the-embed-element) elements\
[`hspace`]{#attr-embed-hspace .dfn dfn-for="embed" dfn-type="element-attr"} on [`embed`{#non-conforming-features:the-embed-element-5}](iframe-embed-object.html#the-embed-element) elements\
[`vspace`]{#attr-embed-vspace .dfn dfn-for="embed" dfn-type="element-attr"} on [`embed`{#non-conforming-features:the-embed-element-6}](iframe-embed-object.html#the-embed-element) elements\
[`align`]{#attr-hr-align .dfn dfn-for="hr" dfn-type="element-attr"} on [`hr`{#non-conforming-features:the-hr-element}](grouping-content.html#the-hr-element) elements\
[`color`]{#attr-hr-color .dfn dfn-for="hr" dfn-type="element-attr"} on [`hr`{#non-conforming-features:the-hr-element-2}](grouping-content.html#the-hr-element) elements\
[`noshade`]{#attr-hr-noshade .dfn dfn-for="hr" dfn-type="element-attr"} on [`hr`{#non-conforming-features:the-hr-element-3}](grouping-content.html#the-hr-element) elements\
[`size`]{#attr-hr-size .dfn dfn-for="hr" dfn-type="element-attr"} on [`hr`{#non-conforming-features:the-hr-element-4}](grouping-content.html#the-hr-element) elements\
[`width`]{#attr-hr-width .dfn dfn-for="hr" dfn-type="element-attr"} on [`hr`{#non-conforming-features:the-hr-element-5}](grouping-content.html#the-hr-element) elements\
[`align`]{#attr-hx-align .dfn dfn-for="h1,h2,h3,h4,h5,h6" dfn-type="element-attr"} on [`h1`{#non-conforming-features:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)---[`h6`{#non-conforming-features:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements) elements\
[`align`]{#attr-iframe-align .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-5}](iframe-embed-object.html#the-iframe-element) elements\
[`allowtransparency`]{#attr-iframe-allowtransparency .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-6}](iframe-embed-object.html#the-iframe-element) elements\
[`frameborder`]{#attr-iframe-frameborder .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-7}](iframe-embed-object.html#the-iframe-element) elements\
[`framespacing`]{#attr-iframe-framespacing .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-8}](iframe-embed-object.html#the-iframe-element) elements\
[`hspace`]{#attr-iframe-hspace .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-9}](iframe-embed-object.html#the-iframe-element) elements\
[`marginheight`]{#attr-iframe-marginheight .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-10}](iframe-embed-object.html#the-iframe-element) elements\
[`marginwidth`]{#attr-iframe-marginwidth .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-11}](iframe-embed-object.html#the-iframe-element) elements\
[`scrolling`]{#attr-iframe-scrolling .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-12}](iframe-embed-object.html#the-iframe-element) elements\
[`vspace`]{#attr-iframe-vspace .dfn dfn-for="iframe" dfn-type="element-attr"} on [`iframe`{#non-conforming-features:the-iframe-element-13}](iframe-embed-object.html#the-iframe-element) elements\
[`align`]{#attr-input-align .dfn dfn-for="input" dfn-type="element-attr"} on [`input`{#non-conforming-features:the-input-element-8}](input.html#the-input-element) elements\
[`border`]{#attr-input-border .dfn dfn-for="input" dfn-type="element-attr"} on [`input`{#non-conforming-features:the-input-element-9}](input.html#the-input-element) elements\
[`hspace`]{#attr-input-hspace .dfn dfn-for="input" dfn-type="element-attr"} on [`input`{#non-conforming-features:the-input-element-10}](input.html#the-input-element) elements\
[`vspace`]{#attr-input-vspace .dfn dfn-for="input" dfn-type="element-attr"} on [`input`{#non-conforming-features:the-input-element-11}](input.html#the-input-element) elements\
[`align`]{#attr-img-align .dfn dfn-for="img" dfn-type="element-attr"} on [`img`{#non-conforming-features:the-img-element-7}](embedded-content.html#the-img-element) elements\
[`border`]{#attr-img-border .dfn dfn-for="img" dfn-type="element-attr"} on [`img`{#non-conforming-features:the-img-element-8}](embedded-content.html#the-img-element) elements (except as noted in the previous section)\
[`hspace`]{#attr-img-hspace .dfn dfn-for="img" dfn-type="element-attr"} on [`img`{#non-conforming-features:the-img-element-9}](embedded-content.html#the-img-element) elements\
[`vspace`]{#attr-img-vspace .dfn dfn-for="img" dfn-type="element-attr"} on [`img`{#non-conforming-features:the-img-element-10}](embedded-content.html#the-img-element) elements\
[`align`]{#attr-legend-align .dfn dfn-for="legend" dfn-type="element-attr"} on [`legend`{#non-conforming-features:the-legend-element-4}](form-elements.html#the-legend-element) elements\
[`type`]{#attr-li-type .dfn dfn-for="li" dfn-type="element-attr"} on [`li`{#non-conforming-features:the-li-element}](grouping-content.html#the-li-element) elements\
[`compact`]{#attr-menu-compact .dfn dfn-for="menu" dfn-type="element-attr"} on [`menu`{#non-conforming-features:the-menu-element-3}](grouping-content.html#the-menu-element) elements\
[`align`]{#attr-object-align .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-18}](iframe-embed-object.html#the-object-element) elements\
[`border`]{#attr-object-border .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-19}](iframe-embed-object.html#the-object-element) elements\
[`hspace`]{#attr-object-hspace .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-20}](iframe-embed-object.html#the-object-element) elements\
[`vspace`]{#attr-object-vspace .dfn dfn-for="object" dfn-type="element-attr"} on [`object`{#non-conforming-features:the-object-element-21}](iframe-embed-object.html#the-object-element) elements\
[`compact`]{#attr-ol-compact .dfn dfn-for="ol" dfn-type="element-attr"} on [`ol`{#non-conforming-features:the-ol-element}](grouping-content.html#the-ol-element) elements\
[`align`]{#attr-p-align .dfn dfn-for="p" dfn-type="element-attr"} on [`p`{#non-conforming-features:the-p-element}](grouping-content.html#the-p-element) elements\
[`width`]{#attr-pre-width .dfn dfn-for="pre" dfn-type="element-attr"} on [`pre`{#non-conforming-features:the-pre-element-3}](grouping-content.html#the-pre-element) elements\
[`align`]{#attr-table-align .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-6}](tables.html#the-table-element) elements\
[`bgcolor`]{#attr-table-bgcolor .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-7}](tables.html#the-table-element) elements\
[`border`]{#attr-table-border .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-8}](tables.html#the-table-element) elements\
[`bordercolor`]{#attr-table-bordercolor .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-9}](tables.html#the-table-element) elements\
[`cellpadding`]{#attr-table-cellpadding .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-10}](tables.html#the-table-element) elements\
[`cellspacing`]{#attr-table-cellspacing .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-11}](tables.html#the-table-element) elements\
[`frame`]{#attr-table-frame .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-12}](tables.html#the-table-element) elements\
[`height`]{#attr-table-height .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-13}](tables.html#the-table-element) elements\
[`rules`]{#attr-table-rules .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-14}](tables.html#the-table-element) elements\
[`width`]{#attr-table-width .dfn dfn-for="table" dfn-type="element-attr"} on [`table`{#non-conforming-features:the-table-element-15}](tables.html#the-table-element) elements\
[`align`]{#attr-tbody-align .dfn dfn-for="tbody" dfn-type="element-attr"} on [`tbody`{#non-conforming-features:the-tbody-element}](tables.html#the-tbody-element), [`thead`{#non-conforming-features:the-thead-element}](tables.html#the-thead-element), and [`tfoot`{#non-conforming-features:the-tfoot-element}](tables.html#the-tfoot-element) elements\
[`char`]{#attr-tbody-char .dfn dfn-for="tbody" dfn-type="element-attr"} on [`tbody`{#non-conforming-features:the-tbody-element-2}](tables.html#the-tbody-element), [`thead`{#non-conforming-features:the-thead-element-2}](tables.html#the-thead-element), and [`tfoot`{#non-conforming-features:the-tfoot-element-2}](tables.html#the-tfoot-element) elements\
[`charoff`]{#attr-tbody-charoff .dfn dfn-for="tbody" dfn-type="element-attr"} on [`tbody`{#non-conforming-features:the-tbody-element-3}](tables.html#the-tbody-element), [`thead`{#non-conforming-features:the-thead-element-3}](tables.html#the-thead-element), and [`tfoot`{#non-conforming-features:the-tfoot-element-3}](tables.html#the-tfoot-element) elements\
[`height`]{#attr-tbody-height .dfn dfn-for="tbody" dfn-type="element-attr"} on [`thead`{#non-conforming-features:the-thead-element-4}](tables.html#the-thead-element), [`tbody`{#non-conforming-features:the-tbody-element-4}](tables.html#the-tbody-element), and [`tfoot`{#non-conforming-features:the-tfoot-element-4}](tables.html#the-tfoot-element) elements\
[`valign`]{#attr-tbody-valign .dfn dfn-for="tbody" dfn-type="element-attr"} on [`tbody`{#non-conforming-features:the-tbody-element-5}](tables.html#the-tbody-element), [`thead`{#non-conforming-features:the-thead-element-5}](tables.html#the-thead-element), and [`tfoot`{#non-conforming-features:the-tfoot-element-5}](tables.html#the-tfoot-element) elements\
[`align`]{#attr-tdth-align .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-4}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-5}](tables.html#the-th-element) elements\
[`bgcolor`]{#attr-tdth-bgcolor .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-5}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-6}](tables.html#the-th-element) elements\
[`char`]{#attr-tdth-char .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-6}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-7}](tables.html#the-th-element) elements\
[`charoff`]{#attr-tdth-charoff .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-7}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-8}](tables.html#the-th-element) elements\
[`height`]{#attr-tdth-height .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-8}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-9}](tables.html#the-th-element) elements\
[`nowrap`]{#attr-tdth-nowrap .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-9}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-10}](tables.html#the-th-element) elements\
[`valign`]{#attr-tdth-valign .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-10}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-11}](tables.html#the-th-element) elements\
[`width`]{#attr-tdth-width .dfn dfn-for="td,th" dfn-type="element-attr"} on [`td`{#non-conforming-features:the-td-element-11}](tables.html#the-td-element) and [`th`{#non-conforming-features:the-th-element-12}](tables.html#the-th-element) elements\
[`align`]{#attr-tr-align .dfn dfn-for="tr" dfn-type="element-attr"} on [`tr`{#non-conforming-features:the-tr-element}](tables.html#the-tr-element) elements\
[`bgcolor`]{#attr-tr-bgcolor .dfn dfn-for="tr" dfn-type="element-attr"} on [`tr`{#non-conforming-features:the-tr-element-2}](tables.html#the-tr-element) elements\
[`char`]{#attr-tr-char .dfn dfn-for="tr" dfn-type="element-attr"} on [`tr`{#non-conforming-features:the-tr-element-3}](tables.html#the-tr-element) elements\
[`charoff`]{#attr-tr-charoff .dfn dfn-for="tr" dfn-type="element-attr"} on [`tr`{#non-conforming-features:the-tr-element-4}](tables.html#the-tr-element) elements\
[`height`]{#attr-tr-height .dfn dfn-for="tr" dfn-type="element-attr"} on [`tr`{#non-conforming-features:the-tr-element-5}](tables.html#the-tr-element) elements\
[`valign`]{#attr-tr-valign .dfn dfn-for="tr" dfn-type="element-attr"} on [`tr`{#non-conforming-features:the-tr-element-6}](tables.html#the-tr-element) elements\
[`compact`]{#attr-ul-compact .dfn dfn-for="ul" dfn-type="element-attr"} on [`ul`{#non-conforming-features:the-ul-element-2}](grouping-content.html#the-ul-element) elements\
[`type`]{#attr-ul-type .dfn dfn-for="ul" dfn-type="element-attr"} on [`ul`{#non-conforming-features:the-ul-element-3}](grouping-content.html#the-ul-element) elements\
[`background`]{#attr-background .dfn dfn-type="element-attr"} on [`body`{#non-conforming-features:the-body-element-12}](sections.html#the-body-element), [`table`{#non-conforming-features:the-table-element-16}](tables.html#the-table-element), [`thead`{#non-conforming-features:the-thead-element-6}](tables.html#the-thead-element), [`tbody`{#non-conforming-features:the-tbody-element-6}](tables.html#the-tbody-element), [`tfoot`{#non-conforming-features:the-tfoot-element-6}](tables.html#the-tfoot-element), [`tr`{#non-conforming-features:the-tr-element-7}](tables.html#the-tr-element), [`td`{#non-conforming-features:the-td-element-12}](tables.html#the-td-element), and [`th`{#non-conforming-features:the-th-element-13}](tables.html#the-th-element) elements

:   Use CSS instead.

### [16.3]{.secno} Requirements for implementations[](#requirements-for-implementations){.self-link}

#### [16.3.1]{.secno} The [`marquee`]{.dfn dfn-type="element"} element[](#the-marquee-element){.self-link}

The
[`marquee`{#the-marquee-element:the-marquee-element}](#the-marquee-element)
element is a presentational element that animates content. CSS
transitions and animations are a more appropriate mechanism.
[\[CSSANIMATIONS\]](references.html#refsCSSANIMATIONS)
[\[CSSTRANSITIONS\]](references.html#refsCSSTRANSITIONS)

The
[`marquee`{#the-marquee-element:the-marquee-element-2}](#the-marquee-element)
element must implement the
[`HTMLMarqueeElement`{#the-marquee-element:htmlmarqueeelement}](#htmlmarqueeelement)
interface.

``` idl
[Exposed=Window]
interface HTMLMarqueeElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions] attribute DOMString behavior;
  [CEReactions] attribute DOMString bgColor;
  [CEReactions] attribute DOMString direction;
  [CEReactions] attribute DOMString height;
  [CEReactions] attribute unsigned long hspace;
  [CEReactions] attribute long loop;
  [CEReactions] attribute unsigned long scrollAmount;
  [CEReactions] attribute unsigned long scrollDelay;
  [CEReactions] attribute boolean trueSpeed;
  [CEReactions] attribute unsigned long vspace;
  [CEReactions] attribute DOMString width;

  undefined start();
  undefined stop();
};
```

A
[`marquee`{#the-marquee-element:the-marquee-element-3}](#the-marquee-element)
element can be [turned on]{#concept-marquee-on .dfn} or [turned
off]{#concept-marquee-off .dfn}. When it is created, it is [turned
on](#concept-marquee-on){#the-marquee-element:concept-marquee-on}.

When the [`start()`]{#dom-marquee-start .dfn
dfn-for="HTMLMarqueeElement" dfn-type="method"} method is called, the
[`marquee`{#the-marquee-element:the-marquee-element-4}](#the-marquee-element)
element must be [turned
on](#concept-marquee-on){#the-marquee-element:concept-marquee-on-2}.

When the [`stop()`]{#dom-marquee-stop .dfn dfn-for="HTMLMarqueeElement"
dfn-type="method"} method is called, the
[`marquee`{#the-marquee-element:the-marquee-element-5}](#the-marquee-element)
element must be [turned
off](#concept-marquee-off){#the-marquee-element:concept-marquee-off}.

------------------------------------------------------------------------

The [`behavior`]{#attr-marquee-behavior .dfn dfn-for="marquee"
dfn-type="element-attr"} content attribute on
[`marquee`{#the-marquee-element:the-marquee-element-6}](#the-marquee-element)
elements is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-marquee-element:enumerated-attribute}
with the following keywords and states (all non-conforming):

Keyword

State

`scroll`

[scroll]{#attr-marquee-behavior-scroll .dfn}

`slide`

[slide]{#attr-marquee-behavior-slide .dfn}

`alternate`

[alternate]{#attr-marquee-behavior-alternate .dfn}

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[scroll](#attr-marquee-behavior-scroll){#the-marquee-element:attr-marquee-behavior-scroll}
state.

------------------------------------------------------------------------

The [`direction`]{#attr-marquee-direction .dfn dfn-for="marquee"
dfn-type="element-attr"} content attribute on
[`marquee`{#the-marquee-element:the-marquee-element-7}](#the-marquee-element)
elements is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-marquee-element:enumerated-attribute-2}
with the following keywords and states (all non-conforming):

Keyword

State

`left`

[left]{#attr-marquee-direction-left .dfn}

`right`

[right]{#attr-marquee-direction-right .dfn}

`up`

[up]{#attr-marquee-direction-up .dfn}

`down`

[down]{#attr-marquee-direction-down .dfn}

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[left](#attr-marquee-direction-left){#the-marquee-element:attr-marquee-direction-left}
state.

------------------------------------------------------------------------

The [`truespeed`]{#attr-marquee-truespeed .dfn dfn-for="marquee"
dfn-type="element-attr"} content attribute on
[`marquee`{#the-marquee-element:the-marquee-element-8}](#the-marquee-element)
elements is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-marquee-element:boolean-attribute}.

------------------------------------------------------------------------

A
[`marquee`{#the-marquee-element:the-marquee-element-9}](#the-marquee-element)
element has a [marquee scroll interval]{#marquee-scroll-interval .dfn},
which is obtained as follows:

1.  If the element has a `scrolldelay` attribute, and parsing its value
    using the [rules for parsing non-negative
    integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-marquee-element:rules-for-parsing-non-negative-integers}
    does not return an error, then let `delay`{.variable} be the parsed
    value. Otherwise, let `delay`{.variable} be 85.

2.  If the element does not have a
    [`truespeed`{#the-marquee-element:attr-marquee-truespeed}](#attr-marquee-truespeed)
    attribute, and the `delay`{.variable} value is less than 60, then
    let `delay`{.variable} be 60 instead.

3.  The [marquee scroll
    interval](#marquee-scroll-interval){#the-marquee-element:marquee-scroll-interval}
    is `delay`{.variable}, interpreted in milliseconds.

------------------------------------------------------------------------

A
[`marquee`{#the-marquee-element:the-marquee-element-10}](#the-marquee-element)
element has a [marquee scroll distance]{#marquee-scroll-distance .dfn},
which, if the element has a `scrollamount` attribute, and parsing its
value using the [rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-marquee-element:rules-for-parsing-non-negative-integers-2}
does not return an error, is the parsed value interpreted in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-marquee-element:'px'
x-internal="'px'"}, and otherwise is 6 [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-marquee-element:'px'-2
x-internal="'px'"}.

------------------------------------------------------------------------

A
[`marquee`{#the-marquee-element:the-marquee-element-11}](#the-marquee-element)
element has a [marquee loop count]{#marquee-loop-count .dfn}, which, if
the element has a [`loop`]{#attr-marquee-loop .dfn dfn-for="marquee"
dfn-type="element-attr"} attribute, and parsing its value using the
[rules for parsing
integers](common-microsyntaxes.html#rules-for-parsing-integers){#the-marquee-element:rules-for-parsing-integers}
does not return an error or a number less than 1, is the parsed value,
and otherwise is −1.

The [`loop`]{#dom-marquee-loop .dfn dfn-for="HTMLMarqueeElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
element\'s [marquee loop
count](#marquee-loop-count){#the-marquee-element:marquee-loop-count};
and on setting, if the new value is different than the element\'s
[marquee loop
count](#marquee-loop-count){#the-marquee-element:marquee-loop-count-2}
and either greater than zero or equal to −1, must set the element\'s
[`loop`{#the-marquee-element:attr-marquee-loop}](#attr-marquee-loop)
content attribute (adding it if necessary) to the [valid
integer](common-microsyntaxes.html#valid-integer){#the-marquee-element:valid-integer}
that represents the new value. (Other values are ignored.)

A
[`marquee`{#the-marquee-element:the-marquee-element-12}](#the-marquee-element)
element also has a [marquee current loop
index]{#marquee-current-loop-index .dfn}, which is zero when the element
is created.

The rendering layer will occasionally [increment the marquee current
loop index]{#increment-the-marquee-current-loop-index .dfn}, which must
cause the following steps to be run:

1.  If the [marquee loop
    count](#marquee-loop-count){#the-marquee-element:marquee-loop-count-3}
    is −1, then return.

2.  Increment the [marquee current loop
    index](#marquee-current-loop-index){#the-marquee-element:marquee-current-loop-index}
    by one.

3.  If the [marquee current loop
    index](#marquee-current-loop-index){#the-marquee-element:marquee-current-loop-index-2}
    is now greater than or equal to the element\'s [marquee loop
    count](#marquee-loop-count){#the-marquee-element:marquee-loop-count-4},
    [turn
    off](#concept-marquee-off){#the-marquee-element:concept-marquee-off-2}
    the
    [`marquee`{#the-marquee-element:the-marquee-element-13}](#the-marquee-element)
    element.

------------------------------------------------------------------------

The [`behavior`]{#dom-marquee-behavior .dfn dfn-for="HTMLMarqueeElement"
dfn-type="attribute"}, [`direction`]{#dom-marquee-direction .dfn
dfn-for="HTMLMarqueeElement" dfn-type="attribute"},
[`height`]{#dom-marquee-height .dfn dfn-for="HTMLMarqueeElement"
dfn-type="attribute"}, [`hspace`]{#dom-marquee-hspace .dfn
dfn-for="HTMLMarqueeElement" dfn-type="attribute"},
[`vspace`]{#dom-marquee-vspace .dfn dfn-for="HTMLMarqueeElement"
dfn-type="attribute"}, and [`width`]{#dom-marquee-width .dfn
dfn-for="HTMLMarqueeElement" dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-marquee-element:reflect}
the respective content attributes of the same name.

The [`bgColor`]{#dom-marquee-bgcolor .dfn dfn-for="HTMLMarqueeElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-marquee-element:reflect-2}
the `bgcolor` content attribute.

The [`scrollAmount`]{#dom-marquee-scrollamount .dfn
dfn-for="HTMLMarqueeElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-marquee-element:reflect-3}
the `scrollamount` content attribute. The [default
value](common-dom-interfaces.html#default-value){#the-marquee-element:default-value}
is 6.

The [`scrollDelay`]{#dom-marquee-scrolldelay .dfn
dfn-for="HTMLMarqueeElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-marquee-element:reflect-4}
the `scrolldelay` content attribute. The [default
value](common-dom-interfaces.html#default-value){#the-marquee-element:default-value-2}
is 85.

The [`trueSpeed`]{#dom-marquee-truespeed .dfn
dfn-for="HTMLMarqueeElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-marquee-element:reflect-5}
the
[`truespeed`{#the-marquee-element:attr-marquee-truespeed-2}](#attr-marquee-truespeed)
content attribute.

#### [16.3.2]{.secno} Frames[](#frames){.self-link}

The [`frameset`]{#frameset .dfn dfn-type="element"} element acts as [the
body element](dom.html#the-body-element-2){#frames:the-body-element-2}
in documents that use frames.

The [`frameset`{#frames:frameset}](#frameset) element must implement the
[`HTMLFrameSetElement`{#frames:htmlframesetelement}](#htmlframesetelement)
interface.

``` idl
[Exposed=Window]
interface HTMLFrameSetElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions] attribute DOMString cols;
  [CEReactions] attribute DOMString rows;
};
HTMLFrameSetElement includes WindowEventHandlers;
```

The [`cols`]{#dom-frameset-cols .dfn dfn-for="HTMLFrameSetElement"
dfn-type="attribute"} and [`rows`]{#dom-frameset-rows .dfn
dfn-for="HTMLFrameSetElement" dfn-type="attribute"} IDL attributes of
the [`frameset`{#frames:frameset-2}](#frameset) element must
[reflect](common-dom-interfaces.html#reflect){#frames:reflect} the
respective content attributes of the same name.

The [`frameset`{#frames:frameset-3}](#frameset) element exposes as
[event handler content
attributes](webappapis.html#event-handler-content-attributes){#frames:event-handler-content-attributes}
a number of the [event
handlers](webappapis.html#event-handlers){#frames:event-handlers} of the
[`Window`{#frames:window}](nav-history-apis.html#window) object. It also
mirrors their [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#frames:event-handler-idl-attributes}.

The [event
handlers](webappapis.html#event-handlers){#frames:event-handlers-2} of
the [`Window`{#frames:window-2}](nav-history-apis.html#window) object
named by the [`Window`-reflecting body element event handler
set](webappapis.html#window-reflecting-body-element-event-handler-set){#frames:window-reflecting-body-element-event-handler-set},
exposed on the [`frameset`{#frames:frameset-4}](#frameset) element,
replace the generic [event
handlers](webappapis.html#event-handlers){#frames:event-handlers-3} with
the same names normally supported by [HTML
elements](infrastructure.html#html-elements){#frames:html-elements}.

------------------------------------------------------------------------

The [`frame`]{#frame .dfn dfn-type="element"} element has a [content
navigable](document-sequences.html#content-navigable){#frames:content-navigable}
similar to the
[`iframe`{#frames:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
element, but rendered within a
[`frameset`{#frames:frameset-5}](#frameset) element.

The [`frame`{#frames:frame}](#frame) [HTML element insertion
steps](infrastructure.html#html-element-insertion-steps){#frames:html-element-insertion-steps},
given `insertedNode`{.variable}, are:

1.  If `insertedNode`{.variable} is not [in a document
    tree](https://dom.spec.whatwg.org/#in-a-document-tree){#frames:in-a-document-tree
    x-internal="in-a-document-tree"}, then return.

2.  If `insertedNode`{.variable}\'s
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#frames:root
    x-internal="root"}\'s [browsing
    context](document-sequences.html#concept-document-bc){#frames:concept-document-bc}
    is null, then return.

3.  [Create a new child
    navigable](document-sequences.html#create-a-new-child-navigable){#frames:create-a-new-child-navigable}
    for `insertedNode`{.variable}.

4.  [Process the `frame`
    attributes](#process-the-frame-attributes){#frames:process-the-frame-attributes}
    for `insertedNode`{.variable}, with
    *[initialInsertion](#process-frame-initial-insertion)* set to true.

The [`frame`{#frames:frame-2}](#frame) [HTML element removing
steps](infrastructure.html#html-element-removing-steps){#frames:html-element-removing-steps},
given `removedNode`{.variable}, are to [destroy a child
navigable](document-sequences.html#destroy-a-child-navigable){#frames:destroy-a-child-navigable}
given `removedNode`{.variable}.

Whenever a [`frame`{#frames:frame-3}](#frame) element with a non-null
[content
navigable](document-sequences.html#content-navigable){#frames:content-navigable-2}
has its `src` attribute set, changed, or removed, the user agent must
[process the `frame`
attributes](#process-the-frame-attributes){#frames:process-the-frame-attributes-2}.

To [process the `frame` attributes]{#process-the-frame-attributes .dfn}
for an element `element`{.variable}, with an optional boolean
[`initialInsertion`{.variable}]{#process-frame-initial-insertion .dfn}:

1.  Let `url`{.variable} be the result of running the [shared attribute
    processing steps for `iframe` and `frame`
    elements](iframe-embed-object.html#shared-attribute-processing-steps-for-iframe-and-frame-elements){#frames:shared-attribute-processing-steps-for-iframe-and-frame-elements}
    given `element`{.variable} and `initialInsertion`{.variable}.

2.  If `url`{.variable} is null, then return.

3.  If `url`{.variable} [matches
    `about:blank`](urls-and-fetching.html#matches-about:blank){#frames:matches-about:blank}
    and `initialInsertion`{.variable} is true, then:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#frames:concept-event-fire
        x-internal="concept-event-fire"} named
        [`load`{#frames:event-load}](indices.html#event-load) at
        `element`{.variable}.

    2.  Return.

4.  [Navigate an `iframe` or
    `frame`](iframe-embed-object.html#navigate-an-iframe-or-frame){#frames:navigate-an-iframe-or-frame}
    given `element`{.variable}, `url`{.variable}, the empty string, and
    `initialInsertion`{.variable}.

The [`frame`{#frames:frame-4}](#frame) element [potentially delays the
load
event](iframe-embed-object.html#potentially-delays-the-load-event){#frames:potentially-delays-the-load-event}.

The [`frame`{#frames:frame-5}](#frame) element must implement the
[`HTMLFrameElement`{#frames:htmlframeelement}](#htmlframeelement)
interface.

``` idl
[Exposed=Window]
interface HTMLFrameElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions] attribute DOMString name;
  [CEReactions] attribute DOMString scrolling;
  [CEReactions] attribute USVString src;
  [CEReactions] attribute DOMString frameBorder;
  [CEReactions] attribute USVString longDesc;
  [CEReactions] attribute boolean noResize;
  readonly attribute Document? contentDocument;
  readonly attribute WindowProxy? contentWindow;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString marginHeight;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString marginWidth;
};
```

The [`name`]{#dom-frame-name .dfn dfn-for="HTMLFrameElement"
dfn-type="attribute"}, [`scrolling`]{#dom-frame-scrolling .dfn
dfn-for="HTMLFrameElement" dfn-type="attribute"}, and
[`src`]{#dom-frame-src .dfn dfn-for="HTMLFrameElement"
dfn-type="attribute"} IDL attributes of the
[`frame`{#frames:frame-6}](#frame) element must
[reflect](common-dom-interfaces.html#reflect){#frames:reflect-2} the
respective content attributes of the same name. For the purposes of
reflection, the [`frame`{#frames:frame-7}](#frame) element\'s `src`
content attribute is defined as containing a
[URL](https://url.spec.whatwg.org/#concept-url){#frames:url
x-internal="url"}.

The [`frameBorder`]{#dom-frame-frameborder .dfn
dfn-for="HTMLFrameElement" dfn-type="attribute"} IDL attribute of the
[`frame`{#frames:frame-8}](#frame) element must
[reflect](common-dom-interfaces.html#reflect){#frames:reflect-3} the
element\'s `frameborder` content attribute.

The [`longDesc`]{#dom-frame-longdesc .dfn dfn-for="HTMLFrameElement"
dfn-type="attribute"} IDL attribute of the
[`frame`{#frames:frame-9}](#frame) element must
[reflect](common-dom-interfaces.html#reflect){#frames:reflect-4} the
element\'s `longdesc` content attribute, which for the purposes of
reflection is defined as containing a
[URL](https://url.spec.whatwg.org/#concept-url){#frames:url-2
x-internal="url"}.

The [`noResize`]{#dom-frame-noresize .dfn dfn-for="HTMLFrameElement"
dfn-type="attribute"} IDL attribute of the
[`frame`{#frames:frame-10}](#frame) element must
[reflect](common-dom-interfaces.html#reflect){#frames:reflect-5} the
element\'s `noresize` content attribute.

The [`marginHeight`]{#dom-frame-marginheight .dfn
dfn-for="HTMLFrameElement" dfn-type="attribute"} IDL attribute of the
[`frame`{#frames:frame-11}](#frame) element must
[reflect](common-dom-interfaces.html#reflect){#frames:reflect-6} the
element\'s `marginheight` content attribute.

The [`marginWidth`]{#dom-frame-marginwidth .dfn
dfn-for="HTMLFrameElement" dfn-type="attribute"} IDL attribute of the
[`frame`{#frames:frame-12}](#frame) element must
[reflect](common-dom-interfaces.html#reflect){#frames:reflect-7} the
element\'s `marginwidth` content attribute.

The [`contentDocument`]{#dom-frame-contentdocument .dfn
dfn-for="HTMLFrameElement" dfn-type="attribute"} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#frames:this
x-internal="this"}\'s [content
document](document-sequences.html#concept-bcc-content-document){#frames:concept-bcc-content-document}.

The [`contentWindow`]{#dom-frame-contentwindow .dfn
dfn-for="HTMLFrameElement" dfn-type="attribute"} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#frames:this-2
x-internal="this"}\'s [content
window](document-sequences.html#content-window){#frames:content-window}.

#### [16.3.3]{.secno} Other elements, attributes and APIs[](#other-elements,-attributes-and-apis){.self-link} {#other-elements,-attributes-and-apis}

User agents must treat
[`acronym`{#other-elements,-attributes-and-apis:acronym}](#acronym)
elements in a manner equivalent to
[`abbr`{#other-elements,-attributes-and-apis:the-abbr-element}](text-level-semantics.html#the-abbr-element)
elements in terms of semantics and for purposes of rendering.

------------------------------------------------------------------------

``` idl
partial interface HTMLAnchorElement {
  [CEReactions] attribute DOMString coords;
  [CEReactions] attribute DOMString charset;
  [CEReactions] attribute DOMString name;
  [CEReactions] attribute DOMString rev;
  [CEReactions] attribute DOMString shape;
};
```

The [`coords`]{#dom-a-coords .dfn dfn-for="HTMLAnchorElement"
dfn-type="attribute"}, [`charset`]{#dom-a-charset .dfn
dfn-for="HTMLAnchorElement" dfn-type="attribute"}, [`name`]{#dom-a-name
.dfn dfn-for="HTMLAnchorElement" dfn-type="attribute"},
[`rev`]{#dom-a-rev .dfn dfn-for="HTMLAnchorElement"
dfn-type="attribute"}, and [`shape`]{#dom-a-shape .dfn
dfn-for="HTMLAnchorElement" dfn-type="attribute"} IDL attributes of the
[`a`{#other-elements,-attributes-and-apis:the-a-element}](text-level-semantics.html#the-a-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect}
the respective content attributes of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLAreaElement {
  [CEReactions] attribute boolean noHref;
};
```

The [`noHref`]{#dom-area-nohref .dfn dfn-for="HTMLAreaElement"
dfn-type="attribute"} IDL attribute of the
[`area`{#other-elements,-attributes-and-apis:the-area-element}](image-maps.html#the-area-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-2}
the element\'s
[`nohref`{#other-elements,-attributes-and-apis:attr-area-nohref}](#attr-area-nohref)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLBodyElement {
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString text;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString link;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString vLink;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString aLink;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString bgColor;
  [CEReactions] attribute DOMString background;
};
```

The [`text`]{#dom-body-text .dfn dfn-for="HTMLBodyElement"
dfn-type="attribute"} IDL attribute of the
[`body`{#other-elements,-attributes-and-apis:the-body-element}](sections.html#the-body-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-3}
the element\'s
[`text`{#other-elements,-attributes-and-apis:attr-body-text}](#attr-body-text)
content attribute.

The [`link`]{#dom-body-link .dfn dfn-for="HTMLBodyElement"
dfn-type="attribute"} IDL attribute of the
[`body`{#other-elements,-attributes-and-apis:the-body-element-2}](sections.html#the-body-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-4}
the element\'s
[`link`{#other-elements,-attributes-and-apis:attr-body-link}](#attr-body-link)
content attribute.

The [`aLink`]{#dom-body-alink .dfn dfn-for="HTMLBodyElement"
dfn-type="attribute"} IDL attribute of the
[`body`{#other-elements,-attributes-and-apis:the-body-element-3}](sections.html#the-body-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-5}
the element\'s
[`alink`{#other-elements,-attributes-and-apis:attr-body-alink}](#attr-body-alink)
content attribute.

The [`vLink`]{#dom-body-vlink .dfn dfn-for="HTMLBodyElement"
dfn-type="attribute"} IDL attribute of the
[`body`{#other-elements,-attributes-and-apis:the-body-element-4}](sections.html#the-body-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-6}
the element\'s
[`vlink`{#other-elements,-attributes-and-apis:attr-body-vlink}](#attr-body-vlink)
content attribute.

The [`bgColor`]{#dom-body-bgcolor .dfn dfn-for="HTMLBodyElement"
dfn-type="attribute"} IDL attribute of the
[`body`{#other-elements,-attributes-and-apis:the-body-element-5}](sections.html#the-body-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-7}
the element\'s
[`bgcolor`{#other-elements,-attributes-and-apis:attr-body-bgcolor}](#attr-body-bgcolor)
content attribute.

The [`background`]{#dom-body-background .dfn dfn-for="HTMLBodyElement"
dfn-type="attribute"} IDL attribute of the
[`body`{#other-elements,-attributes-and-apis:the-body-element-6}](sections.html#the-body-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-8}
the element\'s
[`background`{#other-elements,-attributes-and-apis:attr-background}](#attr-background)
content attribute. (The
[`background`{#other-elements,-attributes-and-apis:attr-background-2}](#attr-background)
content is *not* defined to contain a
[URL](https://url.spec.whatwg.org/#concept-url){#other-elements,-attributes-and-apis:url
x-internal="url"}, despite rules regarding its handling in the Rendering
section above.)

------------------------------------------------------------------------

``` idl
partial interface HTMLBRElement {
  [CEReactions] attribute DOMString clear;
};
```

The [`clear`]{#dom-br-clear .dfn dfn-for="HTMLBRElement"
dfn-type="attribute"} IDL attribute of the
[`br`{#other-elements,-attributes-and-apis:the-br-element}](text-level-semantics.html#the-br-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-9}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLTableCaptionElement {
  [CEReactions] attribute DOMString align;
};
```

The [`align`]{#dom-caption-align .dfn dfn-for="HTMLTableCaptionElement"
dfn-type="attribute"} IDL attribute of the
[`caption`{#other-elements,-attributes-and-apis:the-caption-element}](tables.html#the-caption-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-10}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLTableColElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString ch;
  [CEReactions] attribute DOMString chOff;
  [CEReactions] attribute DOMString vAlign;
  [CEReactions] attribute DOMString width;
};
```

The [`align`]{#dom-col-align .dfn dfn-for="HTMLTableColElement"
dfn-type="attribute"} and [`width`]{#dom-col-width .dfn
dfn-for="HTMLTableColElement" dfn-type="attribute"} IDL attributes of
the
[`col`{#other-elements,-attributes-and-apis:the-col-element}](tables.html#the-col-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-11}
the respective content attributes of the same name.

The [`ch`]{#dom-col-ch .dfn dfn-for="HTMLTableColElement"
dfn-type="attribute"} IDL attribute of the
[`col`{#other-elements,-attributes-and-apis:the-col-element-2}](tables.html#the-col-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-12}
the element\'s
[`char`{#other-elements,-attributes-and-apis:attr-col-char}](#attr-col-char)
content attribute.

The [`chOff`]{#dom-col-choff .dfn dfn-for="HTMLTableColElement"
dfn-type="attribute"} IDL attribute of the
[`col`{#other-elements,-attributes-and-apis:the-col-element-3}](tables.html#the-col-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-13}
the element\'s
[`charoff`{#other-elements,-attributes-and-apis:attr-col-charoff}](#attr-col-charoff)
content attribute.

The [`vAlign`]{#dom-col-valign .dfn dfn-for="HTMLTableColElement"
dfn-type="attribute"} IDL attribute of the
[`col`{#other-elements,-attributes-and-apis:the-col-element-4}](tables.html#the-col-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-14}
the element\'s
[`valign`{#other-elements,-attributes-and-apis:attr-col-valign}](#attr-col-valign)
content attribute.

------------------------------------------------------------------------

User agents must treat
[`dir`{#other-elements,-attributes-and-apis:dir}](#dir) elements in a
manner equivalent to
[`ul`{#other-elements,-attributes-and-apis:the-ul-element}](grouping-content.html#the-ul-element)
elements in terms of semantics and for purposes of rendering.

The [`dir`{#other-elements,-attributes-and-apis:dir-2}](#dir) element
must implement the
[`HTMLDirectoryElement`{#other-elements,-attributes-and-apis:htmldirectoryelement}](#htmldirectoryelement)
interface.

``` idl
[Exposed=Window]
interface HTMLDirectoryElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions] attribute boolean compact;
};
```

The [`compact`]{#dom-dir-compact .dfn dfn-for="HTMLDirectoryElement"
dfn-type="attribute"} IDL attribute of the
[`dir`{#other-elements,-attributes-and-apis:dir-3}](#dir) element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-15}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLDivElement {
  [CEReactions] attribute DOMString align;
};
```

The [`align`]{#dom-div-align .dfn dfn-for="HTMLDivElement"
dfn-type="attribute"} IDL attribute of the
[`div`{#other-elements,-attributes-and-apis:the-div-element}](grouping-content.html#the-div-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-16}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLDListElement {
  [CEReactions] attribute boolean compact;
};
```

The [`compact`]{#dom-dl-compact .dfn dfn-for="HTMLDListElement"
dfn-type="attribute"} IDL attribute of the
[`dl`{#other-elements,-attributes-and-apis:the-dl-element}](grouping-content.html#the-dl-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-17}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLEmbedElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString name;
};
```

The [`name`]{#dom-embed-name .dfn dfn-for="HTMLEmbedElement"
dfn-type="attribute"} and [`align`]{#dom-embed-align .dfn
dfn-for="HTMLEmbedElement" dfn-type="attribute"} IDL attributes of the
[`embed`{#other-elements,-attributes-and-apis:the-embed-element}](iframe-embed-object.html#the-embed-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-18}
the respective content attributes of the same name.

------------------------------------------------------------------------

The [`font`{#other-elements,-attributes-and-apis:font}](#font) element
must implement the
[`HTMLFontElement`{#other-elements,-attributes-and-apis:htmlfontelement}](#htmlfontelement)
interface.

``` idl
[Exposed=Window]
interface HTMLFontElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString color;
  [CEReactions] attribute DOMString face;
  [CEReactions] attribute DOMString size; 
};
```

The [`color`]{#dom-font-color .dfn dfn-for="HTMLFontElement"
dfn-type="attribute"}, [`face`]{#dom-font-face .dfn
dfn-for="HTMLFontElement" dfn-type="attribute"}, and
[`size`]{#dom-font-size .dfn dfn-for="HTMLFontElement"
dfn-type="attribute"} IDL attributes of the
[`font`{#other-elements,-attributes-and-apis:font-2}](#font) element
must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-19}
the respective content attributes of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLHeadingElement {
  [CEReactions] attribute DOMString align;
};
```

The [`align`]{#dom-hx-align .dfn dfn-for="HTMLHeadingElement"
dfn-type="attribute"} IDL attribute of the
[`h1`{#other-elements,-attributes-and-apis:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)--[`h6`{#other-elements,-attributes-and-apis:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-20}
the content attribute of the same name.

------------------------------------------------------------------------

The [`profile`]{#dom-head-profile .dfn dfn-for="HTMLHeadElement"
dfn-type="attribute"} IDL attribute on
[`head`{#other-elements,-attributes-and-apis:the-head-element}](semantics.html#the-head-element)
elements (with the
[`HTMLHeadElement`{#other-elements,-attributes-and-apis:htmlheadelement}](semantics.html#htmlheadelement)
interface) is intentionally omitted. Unless so required by [another
applicable
specification](infrastructure.html#other-applicable-specifications){#other-elements,-attributes-and-apis:other-applicable-specifications},
implementations would therefore not support this attribute. (It is
mentioned here as it was defined in a previous version of DOM.)

------------------------------------------------------------------------

``` idl
partial interface HTMLHRElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString color;
  [CEReactions] attribute boolean noShade;
  [CEReactions] attribute DOMString size;
  [CEReactions] attribute DOMString width;
};
```

The [`align`]{#dom-hr-align .dfn dfn-for="HTMLHRElement"
dfn-type="attribute"}, [`color`]{#dom-hr-color .dfn
dfn-for="HTMLHRElement" dfn-type="attribute"}, [`size`]{#dom-hr-size
.dfn dfn-for="HTMLHRElement" dfn-type="attribute"}, and
[`width`]{#dom-hr-width .dfn dfn-for="HTMLHRElement"
dfn-type="attribute"} IDL attributes of the
[`hr`{#other-elements,-attributes-and-apis:the-hr-element}](grouping-content.html#the-hr-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-21}
the respective content attributes of the same name.

The [`noShade`]{#dom-hr-noshade .dfn dfn-for="HTMLHRElement"
dfn-type="attribute"} IDL attribute of the
[`hr`{#other-elements,-attributes-and-apis:the-hr-element-2}](grouping-content.html#the-hr-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-22}
the element\'s
[`noshade`{#other-elements,-attributes-and-apis:attr-hr-noshade}](#attr-hr-noshade)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLHtmlElement {
  [CEReactions] attribute DOMString version;
};
```

The [`version`]{#dom-html-version .dfn dfn-for="HTMLHtmlElement"
dfn-type="attribute"} IDL attribute of the
[`html`{#other-elements,-attributes-and-apis:the-html-element}](semantics.html#the-html-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-23}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLIFrameElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString scrolling;
  [CEReactions] attribute DOMString frameBorder;
  [CEReactions] attribute USVString longDesc;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString marginHeight;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString marginWidth;
};
```

The [`align`]{#dom-iframe-align .dfn dfn-for="HTMLIFrameElement"
dfn-type="attribute"} and [`scrolling`]{#dom-iframe-scrolling .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} IDL attributes of the
[`iframe`{#other-elements,-attributes-and-apis:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-24}
the respective content attributes of the same name.

The [`frameBorder`]{#dom-iframe-frameborder .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} IDL attribute of the
[`iframe`{#other-elements,-attributes-and-apis:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-25}
the element\'s
[`frameborder`{#other-elements,-attributes-and-apis:attr-iframe-frameborder}](#attr-iframe-frameborder)
content attribute.

The [`longDesc`]{#dom-iframe-longdesc .dfn dfn-for="HTMLIFrameElement"
dfn-type="attribute"} IDL attribute of the
[`iframe`{#other-elements,-attributes-and-apis:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-26}
the element\'s
[`longdesc`{#other-elements,-attributes-and-apis:attr-iframe-longdesc}](#attr-iframe-longdesc)
content attribute, which for the purposes of reflection is defined as
containing a
[URL](https://url.spec.whatwg.org/#concept-url){#other-elements,-attributes-and-apis:url-2
x-internal="url"}.

The [`marginHeight`]{#dom-iframe-marginheight .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} IDL attribute of the
[`iframe`{#other-elements,-attributes-and-apis:the-iframe-element-4}](iframe-embed-object.html#the-iframe-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-27}
the element\'s
[`marginheight`{#other-elements,-attributes-and-apis:attr-iframe-marginheight}](#attr-iframe-marginheight)
content attribute.

The [`marginWidth`]{#dom-iframe-marginwidth .dfn
dfn-for="HTMLIFrameElement" dfn-type="attribute"} IDL attribute of the
[`iframe`{#other-elements,-attributes-and-apis:the-iframe-element-5}](iframe-embed-object.html#the-iframe-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-28}
the element\'s
[`marginwidth`{#other-elements,-attributes-and-apis:attr-iframe-marginwidth}](#attr-iframe-marginwidth)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLImageElement {
  [CEReactions] attribute DOMString name;
  [CEReactions] attribute USVString lowsrc;
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute unsigned long hspace;
  [CEReactions] attribute unsigned long vspace;
  [CEReactions] attribute USVString longDesc;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString border;
};
```

The [`name`]{#dom-img-name .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"}, [`align`]{#dom-img-align .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"},
[`border`]{#dom-img-border .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"}, [`hspace`]{#dom-img-hspace .dfn
dfn-for="HTMLImageElement" dfn-type="attribute"}, and
[`vspace`]{#dom-img-vspace .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attributes of the
[`img`{#other-elements,-attributes-and-apis:the-img-element}](embedded-content.html#the-img-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-29}
the respective content attributes of the same name.

The [`longDesc`]{#dom-img-longdesc .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute of the
[`img`{#other-elements,-attributes-and-apis:the-img-element-2}](embedded-content.html#the-img-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-30}
the element\'s
[`longdesc`{#other-elements,-attributes-and-apis:attr-img-longdesc}](#attr-img-longdesc)
content attribute, which for the purposes of reflection is defined as
containing a
[URL](https://url.spec.whatwg.org/#concept-url){#other-elements,-attributes-and-apis:url-3
x-internal="url"}.

The [`lowsrc`]{#dom-img-lowsrc .dfn dfn-for="HTMLImageElement"
dfn-type="attribute"} IDL attribute of the
[`img`{#other-elements,-attributes-and-apis:the-img-element-3}](embedded-content.html#the-img-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-31}
the element\'s
[`lowsrc`{#other-elements,-attributes-and-apis:attr-img-lowsrc}](#attr-img-lowsrc)
content attribute, which for the purposes of reflection is defined as
containing a
[URL](https://url.spec.whatwg.org/#concept-url){#other-elements,-attributes-and-apis:url-4
x-internal="url"}.

------------------------------------------------------------------------

``` idl
partial interface HTMLInputElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString useMap;
};
```

The [`align`]{#dom-input-align .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute of the
[`input`{#other-elements,-attributes-and-apis:the-input-element}](input.html#the-input-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-32}
the content attribute of the same name.

The [`useMap`]{#dom-input-usemap .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute of the
[`input`{#other-elements,-attributes-and-apis:the-input-element-2}](input.html#the-input-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-33}
the element\'s
[`usemap`{#other-elements,-attributes-and-apis:attr-input-usemap}](#attr-input-usemap)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLLegendElement {
  [CEReactions] attribute DOMString align;
};
```

The [`align`]{#dom-legend-align .dfn dfn-for="HTMLLegendElement"
dfn-type="attribute"} IDL attribute of the
[`legend`{#other-elements,-attributes-and-apis:the-legend-element}](form-elements.html#the-legend-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-34}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLLIElement {
  [CEReactions] attribute DOMString type;
};
```

The [`type`]{#dom-li-type .dfn dfn-for="HTMLLIElement"
dfn-type="attribute"} IDL attribute of the
[`li`{#other-elements,-attributes-and-apis:the-li-element}](grouping-content.html#the-li-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-35}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLLinkElement {
  [CEReactions] attribute DOMString charset;
  [CEReactions] attribute DOMString rev;
  [CEReactions] attribute DOMString target;
};
```

The [`charset`]{#dom-link-charset .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"}, [`rev`]{#dom-link-rev .dfn
dfn-for="HTMLLinkElement" dfn-type="attribute"}, and
[`target`]{#dom-link-target .dfn dfn-for="HTMLLinkElement"
dfn-type="attribute"} IDL attributes of the
[`link`{#other-elements,-attributes-and-apis:the-link-element}](semantics.html#the-link-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-36}
the respective content attributes of the same name.

------------------------------------------------------------------------

User agents must treat
[`listing`{#other-elements,-attributes-and-apis:listing}](#listing)
elements in a manner equivalent to
[`pre`{#other-elements,-attributes-and-apis:the-pre-element}](grouping-content.html#the-pre-element)
elements in terms of semantics and for purposes of rendering.

------------------------------------------------------------------------

``` idl
partial interface HTMLMenuElement {
  [CEReactions] attribute boolean compact;
};
```

The [`compact`]{#dom-menu-compact .dfn dfn-for="HTMLMenuElement"
dfn-type="attribute"} IDL attribute of the
[`menu`{#other-elements,-attributes-and-apis:the-menu-element}](grouping-content.html#the-menu-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-37}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLMetaElement {
  [CEReactions] attribute DOMString scheme;
};
```

User agents may treat the
[`scheme`{#other-elements,-attributes-and-apis:attr-meta-scheme}](#attr-meta-scheme)
content attribute on the
[`meta`{#other-elements,-attributes-and-apis:the-meta-element}](semantics.html#the-meta-element)
element as an extension of the element\'s
[`name`{#other-elements,-attributes-and-apis:attr-meta-name}](semantics.html#attr-meta-name)
content attribute when processing a
[`meta`{#other-elements,-attributes-and-apis:the-meta-element-2}](semantics.html#the-meta-element)
element with a
[`name`{#other-elements,-attributes-and-apis:attr-meta-name-2}](semantics.html#attr-meta-name)
attribute whose value is one that the user agent recognizes as
supporting the
[`scheme`{#other-elements,-attributes-and-apis:attr-meta-scheme-2}](#attr-meta-scheme)
attribute.

User agents are encouraged to ignore the
[`scheme`{#other-elements,-attributes-and-apis:attr-meta-scheme-3}](#attr-meta-scheme)
attribute and instead process the value given to the metadata name as if
it had been specified for each expected value of the
[`scheme`{#other-elements,-attributes-and-apis:attr-meta-scheme-4}](#attr-meta-scheme)
attribute.

::: example
For example, if the user agent acts on
[`meta`{#other-elements,-attributes-and-apis:the-meta-element-3}](semantics.html#the-meta-element)
elements with
[`name`{#other-elements,-attributes-and-apis:attr-meta-name-3}](semantics.html#attr-meta-name)
attributes having the value \"eGMS.subject.keyword\", and knows that the
[`scheme`{#other-elements,-attributes-and-apis:attr-meta-scheme-5}](#attr-meta-scheme)
attribute is used with this metadata name, then it could take the
[`scheme`{#other-elements,-attributes-and-apis:attr-meta-scheme-6}](#attr-meta-scheme)
attribute into account, acting as if it was an extension of the
[`name`{#other-elements,-attributes-and-apis:attr-meta-name-4}](semantics.html#attr-meta-name)
attribute. Thus the following two
[`meta`{#other-elements,-attributes-and-apis:the-meta-element-4}](semantics.html#the-meta-element)
elements could be treated as two elements giving values for two
different metadata names, one consisting of a combination of
\"eGMS.subject.keyword\" and \"LGCL\", and the other consisting of a
combination of \"eGMS.subject.keyword\" and \"ORLY\":

``` bad
<!-- this markup is invalid -->
<meta name="eGMS.subject.keyword" scheme="LGCL" content="Abandoned vehicles">
<meta name="eGMS.subject.keyword" scheme="ORLY" content="Mah car: kthxbye">
```

The suggested processing of this markup, however, would be equivalent to
the following:

``` html
<meta name="eGMS.subject.keyword" content="Abandoned vehicles">
<meta name="eGMS.subject.keyword" content="Mah car: kthxbye">
```
:::

The [`scheme`]{#dom-meta-scheme .dfn dfn-for="HTMLMetaElement"
dfn-type="attribute"} IDL attribute of the
[`meta`{#other-elements,-attributes-and-apis:the-meta-element-5}](semantics.html#the-meta-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-38}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLObjectElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString archive;
  [CEReactions] attribute DOMString code;
  [CEReactions] attribute boolean declare;
  [CEReactions] attribute unsigned long hspace;
  [CEReactions] attribute DOMString standby;
  [CEReactions] attribute unsigned long vspace;
  [CEReactions] attribute DOMString codeBase;
  [CEReactions] attribute DOMString codeType;
  [CEReactions] attribute DOMString useMap;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString border;
};
```

The [`align`]{#dom-object-align .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"}, [`archive`]{#dom-object-archive .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"},
[`border`]{#dom-object-border .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"}, [`code`]{#dom-object-code .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"},
[`declare`]{#dom-object-declare .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"}, [`hspace`]{#dom-object-hspace .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"},
[`standby`]{#dom-object-standby .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"}, and [`vspace`]{#dom-object-vspace .dfn
dfn-for="HTMLObjectElement" dfn-type="attribute"} IDL attributes of the
[`object`{#other-elements,-attributes-and-apis:the-object-element}](iframe-embed-object.html#the-object-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-39}
the respective content attributes of the same name.

The [`codeBase`]{#dom-object-codebase .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"} IDL attribute of the
[`object`{#other-elements,-attributes-and-apis:the-object-element-2}](iframe-embed-object.html#the-object-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-40}
the element\'s
[`codebase`{#other-elements,-attributes-and-apis:attr-object-codebase}](#attr-object-codebase)
content attribute, which for the purposes of reflection is defined as
containing a
[URL](https://url.spec.whatwg.org/#concept-url){#other-elements,-attributes-and-apis:url-5
x-internal="url"}.

The [`codeType`]{#dom-object-codetype .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"} IDL attribute of the
[`object`{#other-elements,-attributes-and-apis:the-object-element-3}](iframe-embed-object.html#the-object-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-41}
the element\'s
[`codetype`{#other-elements,-attributes-and-apis:attr-object-codetype}](#attr-object-codetype)
content attribute.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/useMap](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/useMap "The useMap property of the HTMLObjectElement interface returns a string that reflects the usemap HTML attribute, specifying a <map> element to use.")

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

The [`useMap`]{#dom-object-usemap .dfn dfn-for="HTMLObjectElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-42}
the
[`usemap`{#other-elements,-attributes-and-apis:attr-hyperlink-usemap}](image-maps.html#attr-hyperlink-usemap)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLOListElement {
  [CEReactions] attribute boolean compact;
};
```

The [`compact`]{#dom-ol-compact .dfn dfn-for="HTMLOListElement"
dfn-type="attribute"} IDL attribute of the
[`ol`{#other-elements,-attributes-and-apis:the-ol-element}](grouping-content.html#the-ol-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-43}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLParagraphElement {
  [CEReactions] attribute DOMString align;
};
```

The [`align`]{#dom-p-align .dfn dfn-for="HTMLParagraphElement"
dfn-type="attribute"} IDL attribute of the
[`p`{#other-elements,-attributes-and-apis:the-p-element}](grouping-content.html#the-p-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-44}
the content attribute of the same name.

------------------------------------------------------------------------

The [`param`{#other-elements,-attributes-and-apis:param}](#param)
element must implement the
[`HTMLParamElement`{#other-elements,-attributes-and-apis:htmlparamelement}](#htmlparamelement)
interface.

``` idl
[Exposed=Window]
interface HTMLParamElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions] attribute DOMString name;
  [CEReactions] attribute DOMString value;
  [CEReactions] attribute DOMString type;
  [CEReactions] attribute DOMString valueType;
};
```

The [`name`]{#dom-param-name .dfn dfn-for="HTMLParamElement"
dfn-type="attribute"}, [`value`]{#dom-param-value .dfn
dfn-for="HTMLParamElement" dfn-type="attribute"}, and
[`type`]{#dom-param-type .dfn dfn-for="HTMLParamElement"
dfn-type="attribute"} IDL attributes of the
[`param`{#other-elements,-attributes-and-apis:param-2}](#param) element
must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-45}
the respective content attributes of the same name.

The [`valueType`]{#dom-param-valuetype .dfn dfn-for="HTMLParamElement"
dfn-type="attribute"} IDL attribute of the
[`param`{#other-elements,-attributes-and-apis:param-3}](#param) element
must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-46}
the element\'s `valuetype` content attribute.

------------------------------------------------------------------------

User agents must treat
[`plaintext`{#other-elements,-attributes-and-apis:plaintext}](#plaintext)
elements in a manner equivalent to
[`pre`{#other-elements,-attributes-and-apis:the-pre-element-2}](grouping-content.html#the-pre-element)
elements in terms of semantics and for purposes of rendering. (The
parser has special behavior for this element, though.)

------------------------------------------------------------------------

``` idl
partial interface HTMLPreElement {
  [CEReactions] attribute long width;
};
```

The [`width`]{#dom-pre-width .dfn dfn-for="HTMLPreElement"
dfn-type="attribute"} IDL attribute of the
[`pre`{#other-elements,-attributes-and-apis:the-pre-element-3}](grouping-content.html#the-pre-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-47}
the content attribute of the same name.

------------------------------------------------------------------------

``` idl
partial interface HTMLStyleElement {
  [CEReactions] attribute DOMString type;
};
```

The [`type`]{#dom-style-type .dfn dfn-for="HTMLStyleElement"
dfn-type="attribute"} IDL attribute of the
[`style`{#other-elements,-attributes-and-apis:the-style-element}](semantics.html#the-style-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-48}
the element\'s
[`type`{#other-elements,-attributes-and-apis:attr-style-type}](#attr-style-type)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLScriptElement {
  [CEReactions] attribute DOMString charset;
  [CEReactions] attribute DOMString event;
  [CEReactions] attribute DOMString htmlFor;
};
```

The [`charset`]{#dom-script-charset .dfn dfn-for="HTMLScriptElement"
dfn-type="attribute"} and [`event`]{#dom-script-event .dfn
dfn-for="HTMLScriptElement" dfn-type="attribute"} IDL attributes of the
[`script`{#other-elements,-attributes-and-apis:the-script-element}](scripting.html#the-script-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-49}
the respective content attributes of the same name.

The [`htmlFor`]{#dom-script-htmlfor .dfn dfn-for="HTMLScriptElement"
dfn-type="attribute"} IDL attribute of the
[`script`{#other-elements,-attributes-and-apis:the-script-element-2}](scripting.html#the-script-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-50}
the element\'s
[`for`{#other-elements,-attributes-and-apis:attr-script-for}](#attr-script-for)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLTableElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString border;
  [CEReactions] attribute DOMString frame;
  [CEReactions] attribute DOMString rules;
  [CEReactions] attribute DOMString summary;
  [CEReactions] attribute DOMString width;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString bgColor;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString cellPadding;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString cellSpacing;
};
```

The [`align`]{#dom-table-align .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"}, [`border`]{#dom-table-border .dfn
dfn-for="HTMLTableElement" dfn-type="attribute"},
[`frame`]{#dom-table-frame .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"}, [`summary`]{#dom-table-summary .dfn
dfn-for="HTMLTableElement" dfn-type="attribute"},
[`rules`]{#dom-table-rules .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"}, and [`width`]{#dom-table-width .dfn
dfn-for="HTMLTableElement" dfn-type="attribute"} IDL attributes of the
[`table`{#other-elements,-attributes-and-apis:the-table-element}](tables.html#the-table-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-51}
the respective content attributes of the same name.

The [`bgColor`]{#dom-table-bgcolor .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"} IDL attribute of the
[`table`{#other-elements,-attributes-and-apis:the-table-element-2}](tables.html#the-table-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-52}
the element\'s
[`bgcolor`{#other-elements,-attributes-and-apis:attr-table-bgcolor}](#attr-table-bgcolor)
content attribute.

The [`cellPadding`]{#dom-table-cellpadding .dfn
dfn-for="HTMLTableElement" dfn-type="attribute"} IDL attribute of the
[`table`{#other-elements,-attributes-and-apis:the-table-element-3}](tables.html#the-table-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-53}
the element\'s
[`cellpadding`{#other-elements,-attributes-and-apis:attr-table-cellpadding}](#attr-table-cellpadding)
content attribute.

The [`cellSpacing`]{#dom-table-cellspacing .dfn
dfn-for="HTMLTableElement" dfn-type="attribute"} IDL attribute of the
[`table`{#other-elements,-attributes-and-apis:the-table-element-4}](tables.html#the-table-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-54}
the element\'s
[`cellspacing`{#other-elements,-attributes-and-apis:attr-table-cellspacing}](#attr-table-cellspacing)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLTableSectionElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString ch;
  [CEReactions] attribute DOMString chOff;
  [CEReactions] attribute DOMString vAlign;
};
```

The [`align`]{#dom-tbody-align .dfn dfn-for="HTMLTableSectionElement"
dfn-type="attribute"} IDL attribute of the
[`tbody`{#other-elements,-attributes-and-apis:the-tbody-element}](tables.html#the-tbody-element),
[`thead`{#other-elements,-attributes-and-apis:the-thead-element}](tables.html#the-thead-element),
and
[`tfoot`{#other-elements,-attributes-and-apis:the-tfoot-element}](tables.html#the-tfoot-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-55}
the content attribute of the same name.

The [`ch`]{#dom-tbody-ch .dfn dfn-for="HTMLTableSectionElement"
dfn-type="attribute"} IDL attribute of the
[`tbody`{#other-elements,-attributes-and-apis:the-tbody-element-2}](tables.html#the-tbody-element),
[`thead`{#other-elements,-attributes-and-apis:the-thead-element-2}](tables.html#the-thead-element),
and
[`tfoot`{#other-elements,-attributes-and-apis:the-tfoot-element-2}](tables.html#the-tfoot-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-56}
the elements\'
[`char`{#other-elements,-attributes-and-apis:attr-tbody-char}](#attr-tbody-char)
content attributes.

The [`chOff`]{#dom-tbody-choff .dfn dfn-for="HTMLTableSectionElement"
dfn-type="attribute"} IDL attribute of the
[`tbody`{#other-elements,-attributes-and-apis:the-tbody-element-3}](tables.html#the-tbody-element),
[`thead`{#other-elements,-attributes-and-apis:the-thead-element-3}](tables.html#the-thead-element),
and
[`tfoot`{#other-elements,-attributes-and-apis:the-tfoot-element-3}](tables.html#the-tfoot-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-57}
the elements\'
[`charoff`{#other-elements,-attributes-and-apis:attr-tbody-charoff}](#attr-tbody-charoff)
content attributes.

The [`vAlign`]{#dom-tbody-valign .dfn dfn-for="HTMLTableSectionElement"
dfn-type="attribute"} IDL attribute of the
[`tbody`{#other-elements,-attributes-and-apis:the-tbody-element-4}](tables.html#the-tbody-element),
[`thead`{#other-elements,-attributes-and-apis:the-thead-element-4}](tables.html#the-thead-element),
and
[`tfoot`{#other-elements,-attributes-and-apis:the-tfoot-element-4}](tables.html#the-tfoot-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-58}
the elements\'
[`valign`{#other-elements,-attributes-and-apis:attr-tbody-valign}](#attr-tbody-valign)
content attributes.

------------------------------------------------------------------------

``` idl
partial interface HTMLTableCellElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString axis;
  [CEReactions] attribute DOMString height;
  [CEReactions] attribute DOMString width;

  [CEReactions] attribute DOMString ch;
  [CEReactions] attribute DOMString chOff;
  [CEReactions] attribute boolean noWrap;
  [CEReactions] attribute DOMString vAlign;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString bgColor;
};
```

The [`align`]{#dom-tdth-align .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"}, [`axis`]{#dom-tdth-axis .dfn
dfn-for="HTMLTableCellElement" dfn-type="attribute"},
[`height`]{#dom-tdth-height .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"}, and [`width`]{#dom-tdth-width .dfn
dfn-for="HTMLTableCellElement" dfn-type="attribute"} IDL attributes of
the
[`td`{#other-elements,-attributes-and-apis:the-td-element}](tables.html#the-td-element)
and
[`th`{#other-elements,-attributes-and-apis:the-th-element}](tables.html#the-th-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-59}
the respective content attributes of the same name.

The [`ch`]{#dom-tdth-ch .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute of the
[`td`{#other-elements,-attributes-and-apis:the-td-element-2}](tables.html#the-td-element)
and
[`th`{#other-elements,-attributes-and-apis:the-th-element-2}](tables.html#the-th-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-60}
the elements\'
[`char`{#other-elements,-attributes-and-apis:attr-tdth-char}](#attr-tdth-char)
content attributes.

The [`chOff`]{#dom-tdth-choff .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute of the
[`td`{#other-elements,-attributes-and-apis:the-td-element-3}](tables.html#the-td-element)
and
[`th`{#other-elements,-attributes-and-apis:the-th-element-3}](tables.html#the-th-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-61}
the elements\'
[`charoff`{#other-elements,-attributes-and-apis:attr-tdth-charoff}](#attr-tdth-charoff)
content attributes.

The [`noWrap`]{#dom-tdth-nowrap .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute of the
[`td`{#other-elements,-attributes-and-apis:the-td-element-4}](tables.html#the-td-element)
and
[`th`{#other-elements,-attributes-and-apis:the-th-element-4}](tables.html#the-th-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-62}
the elements\'
[`nowrap`{#other-elements,-attributes-and-apis:attr-tdth-nowrap}](#attr-tdth-nowrap)
content attributes.

The [`vAlign`]{#dom-tdth-valign .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute of the
[`td`{#other-elements,-attributes-and-apis:the-td-element-5}](tables.html#the-td-element)
and
[`th`{#other-elements,-attributes-and-apis:the-th-element-5}](tables.html#the-th-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-63}
the elements\'
[`valign`{#other-elements,-attributes-and-apis:attr-tdth-valign}](#attr-tdth-valign)
content attributes.

The [`bgColor`]{#dom-tdth-bgcolor .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute of the
[`td`{#other-elements,-attributes-and-apis:the-td-element-6}](tables.html#the-td-element)
and
[`th`{#other-elements,-attributes-and-apis:the-th-element-6}](tables.html#the-th-element)
elements must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-64}
the elements\'
[`bgcolor`{#other-elements,-attributes-and-apis:attr-tdth-bgcolor}](#attr-tdth-bgcolor)
content attributes.

------------------------------------------------------------------------

``` idl
partial interface HTMLTableRowElement {
  [CEReactions] attribute DOMString align;
  [CEReactions] attribute DOMString ch;
  [CEReactions] attribute DOMString chOff;
  [CEReactions] attribute DOMString vAlign;

  [CEReactions] attribute [LegacyNullToEmptyString] DOMString bgColor;
};
```

The [`align`]{#dom-tr-align .dfn dfn-for="HTMLTableRowElement"
dfn-type="attribute"} IDL attribute of the
[`tr`{#other-elements,-attributes-and-apis:the-tr-element}](tables.html#the-tr-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-65}
the content attribute of the same name.

The [`ch`]{#dom-tr-ch .dfn dfn-for="HTMLTableRowElement"
dfn-type="attribute"} IDL attribute of the
[`tr`{#other-elements,-attributes-and-apis:the-tr-element-2}](tables.html#the-tr-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-66}
the element\'s
[`char`{#other-elements,-attributes-and-apis:attr-tr-char}](#attr-tr-char)
content attribute.

The [`chOff`]{#dom-tr-choff .dfn dfn-for="HTMLTableRowElement"
dfn-type="attribute"} IDL attribute of the
[`tr`{#other-elements,-attributes-and-apis:the-tr-element-3}](tables.html#the-tr-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-67}
the element\'s
[`charoff`{#other-elements,-attributes-and-apis:attr-tr-charoff}](#attr-tr-charoff)
content attribute.

The [`vAlign`]{#dom-tr-valign .dfn dfn-for="HTMLTableRowElement"
dfn-type="attribute"} IDL attribute of the
[`tr`{#other-elements,-attributes-and-apis:the-tr-element-4}](tables.html#the-tr-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-68}
the element\'s
[`valign`{#other-elements,-attributes-and-apis:attr-tr-valign}](#attr-tr-valign)
content attribute.

The [`bgColor`]{#dom-tr-bgcolor .dfn dfn-for="HTMLTableRowElement"
dfn-type="attribute"} IDL attribute of the
[`tr`{#other-elements,-attributes-and-apis:the-tr-element-5}](tables.html#the-tr-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-69}
the element\'s
[`bgcolor`{#other-elements,-attributes-and-apis:attr-tr-bgcolor}](#attr-tr-bgcolor)
content attribute.

------------------------------------------------------------------------

``` idl
partial interface HTMLUListElement {
  [CEReactions] attribute boolean compact;
  [CEReactions] attribute DOMString type;
};
```

The [`compact`]{#dom-ul-compact .dfn dfn-for="HTMLUListElement"
dfn-type="attribute"} and [`type`]{#dom-ul-type .dfn
dfn-for="HTMLUListElement" dfn-type="attribute"} IDL attributes of the
[`ul`{#other-elements,-attributes-and-apis:the-ul-element-2}](grouping-content.html#the-ul-element)
element must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-70}
the respective content attributes of the same name.

------------------------------------------------------------------------

User agents must treat
[`xmp`{#other-elements,-attributes-and-apis:xmp}](#xmp) elements in a
manner equivalent to
[`pre`{#other-elements,-attributes-and-apis:the-pre-element-4}](grouping-content.html#the-pre-element)
elements in terms of semantics and for purposes of rendering. (The
parser has special behavior for this element though.)

------------------------------------------------------------------------

``` idl
partial interface Document {
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString fgColor;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString linkColor;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString vlinkColor;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString alinkColor;
  [CEReactions] attribute [LegacyNullToEmptyString] DOMString bgColor;

  [SameObject] readonly attribute HTMLCollection anchors;
  [SameObject] readonly attribute HTMLCollection applets;

  undefined clear();
  undefined captureEvents();
  undefined releaseEvents();

  [SameObject] readonly attribute HTMLAllCollection all;
};
```

The attributes of the
[`Document`{#other-elements,-attributes-and-apis:document}](dom.html#document)
object listed in the first column of the following table must
[reflect](common-dom-interfaces.html#reflect){#other-elements,-attributes-and-apis:reflect-71}
the content attribute on [the body
element](dom.html#the-body-element-2){#other-elements,-attributes-and-apis:the-body-element-2-2}
with the name given in the corresponding cell in the second column on
the same row, if [the body
element](dom.html#the-body-element-2){#other-elements,-attributes-and-apis:the-body-element-2-3}
is a
[`body`{#other-elements,-attributes-and-apis:the-body-element-7}](sections.html#the-body-element)
element (as opposed to a
[`frameset`{#other-elements,-attributes-and-apis:frameset}](#frameset)
element). When there is no [body
element](dom.html#the-body-element-2){#other-elements,-attributes-and-apis:the-body-element-2-4}
or if it is a
[`frameset`{#other-elements,-attributes-and-apis:frameset-2}](#frameset)
element, the attributes must instead return the empty string on getting
and do nothing on setting.

IDL attribute

Content attribute

[`fgColor`]{#dom-document-fgcolor .dfn dfn-for="Document"
dfn-type="attribute"}

[`text`{#other-elements,-attributes-and-apis:attr-body-text-2}](#attr-body-text)

[`linkColor`]{#dom-document-linkcolor .dfn dfn-for="Document"
dfn-type="attribute"}

[`link`{#other-elements,-attributes-and-apis:attr-body-link-2}](#attr-body-link)

[`vlinkColor`]{#dom-document-vlinkcolor .dfn dfn-for="Document"
dfn-type="attribute"}

[`vlink`{#other-elements,-attributes-and-apis:attr-body-vlink-2}](#attr-body-vlink)

[`alinkColor`]{#dom-document-alinkcolor .dfn dfn-for="Document"
dfn-type="attribute"}

[`alink`{#other-elements,-attributes-and-apis:attr-body-alink-2}](#attr-body-alink)

[`bgColor`]{#dom-document-bgcolor .dfn dfn-for="Document"
dfn-type="attribute"}

[`bgcolor`{#other-elements,-attributes-and-apis:attr-body-bgcolor-2}](#attr-body-bgcolor)

------------------------------------------------------------------------

The [`anchors`]{#dom-document-anchors .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#other-elements,-attributes-and-apis:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`Document`{#other-elements,-attributes-and-apis:document-2}](dom.html#document)
node, whose filter matches only
[`a`{#other-elements,-attributes-and-apis:the-a-element-2}](text-level-semantics.html#the-a-element)
elements with
[`name`{#other-elements,-attributes-and-apis:attr-a-name}](#attr-a-name)
attributes.

The [`applets`]{#dom-document-applets .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#other-elements,-attributes-and-apis:htmlcollection-4}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`Document`{#other-elements,-attributes-and-apis:document-3}](dom.html#document)
node, whose filter matches nothing. (It exists for historical reasons.)

The [`clear()`]{#dom-document-clear .dfn dfn-for="Document"
dfn-type="method"}, [`captureEvents()`]{#dom-document-captureevents .dfn
dfn-for="Document" dfn-type="method"}, and
[`releaseEvents()`]{#dom-document-releaseevents .dfn dfn-for="Document"
dfn-type="method"} methods must do nothing.

------------------------------------------------------------------------

The [`all`]{#dom-document-all .dfn dfn-for="Document"
dfn-type="attribute"} attribute must return an
[`HTMLAllCollection`{#other-elements,-attributes-and-apis:htmlallcollection-2}](common-dom-interfaces.html#htmlallcollection)
rooted at the
[`Document`{#other-elements,-attributes-and-apis:document-4}](dom.html#document)
node, whose filter matches all elements.

------------------------------------------------------------------------

``` idl
partial interface Window {
  undefined captureEvents();
  undefined releaseEvents();

  [Replaceable, SameObject] readonly attribute External external;
};
```

The [`captureEvents()`]{#dom-window-captureevents .dfn dfn-for="Window"
dfn-type="method"} and [`releaseEvents()`]{#dom-window-releaseevents
.dfn dfn-for="Window" dfn-type="method"} methods must do nothing.

The [`external`]{#dom-external .dfn dfn-for="Window"
dfn-type="attribute"} attribute of the
[`Window`{#other-elements,-attributes-and-apis:window}](nav-history-apis.html#window)
interface must return an instance of the
[`External`{#other-elements,-attributes-and-apis:external-2}](#external)
interface:

``` idl
[Exposed=Window]
interface External {
  undefined AddSearchProvider();
  undefined IsSearchProviderInstalled();
};
```

The [`AddSearchProvider()`]{#dom-external-addsearchprovider .dfn
dfn-for="External" dfn-type="method"} and
[`IsSearchProviderInstalled()`]{#dom-external-issearchproviderinstalled
.dfn dfn-for="External" dfn-type="method"} methods must do nothing.

[← 15 Rendering](rendering.html) --- [Table of Contents](./) --- [17
IANA considerations →](iana.html)
