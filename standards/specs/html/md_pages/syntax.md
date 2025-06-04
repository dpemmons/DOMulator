HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 12 Web storage](webstorage.html) --- [Table of Contents](./) ---
[13.2 Parsing HTML documents →](parsing.html)

1.  [[[13]{.secno} The HTML syntax](syntax.html#syntax)]{#toc-syntax}
    1.  [[13.1]{.secno} Writing HTML documents](syntax.html#writing)
        1.  [[13.1.1]{.secno} The DOCTYPE](syntax.html#the-doctype)
        2.  [[13.1.2]{.secno} Elements](syntax.html#elements-2)
            1.  [[13.1.2.1]{.secno} Start tags](syntax.html#start-tags)
            2.  [[13.1.2.2]{.secno} End tags](syntax.html#end-tags)
            3.  [[13.1.2.3]{.secno}
                Attributes](syntax.html#attributes-2)
            4.  [[13.1.2.4]{.secno} Optional
                tags](syntax.html#optional-tags)
            5.  [[13.1.2.5]{.secno} Restrictions on content
                models](syntax.html#element-restrictions)
            6.  [[13.1.2.6]{.secno} Restrictions on the contents of raw
                text and escapable raw text
                elements](syntax.html#cdata-rcdata-restrictions)
        3.  [[13.1.3]{.secno} Text](syntax.html#text-2)
            1.  [[13.1.3.1]{.secno} Newlines](syntax.html#newlines)
        4.  [[13.1.4]{.secno} Character
            references](syntax.html#character-references)
        5.  [[13.1.5]{.secno} CDATA
            sections](syntax.html#cdata-sections)
        6.  [[13.1.6]{.secno} Comments](syntax.html#comments)

## [13]{.secno} [The HTML syntax]{.dfn}[](#syntax){.self-link} {#syntax}

This section only describes the rules for resources labeled with an
[HTML MIME
type](https://mimesniff.spec.whatwg.org/#html-mime-type){#syntax:html-mime-type
x-internal="html-mime-type"}. Rules for XML resources are discussed in
the section below entitled \"[The XML
syntax](xhtml.html#the-xhtml-syntax){#syntax:the-xhtml-syntax}\".

### [13.1]{.secno} Writing HTML documents[](#writing){.self-link} {#writing}

*This section only applies to documents, authoring tools, and markup
generators. In particular, it does not apply to conformance checkers;
conformance checkers must use the requirements given in the next section
(\"parsing HTML documents\").*

Documents must consist of the following parts, in the given order:

1.  Optionally, a single U+FEFF BYTE ORDER MARK (BOM) character.
2.  Any number of [comments](#syntax-comments){#writing:syntax-comments}
    and [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#writing:space-characters
    x-internal="space-characters"}.
3.  A [DOCTYPE](#syntax-doctype){#writing:syntax-doctype}.
4.  Any number of
    [comments](#syntax-comments){#writing:syntax-comments-2} and [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#writing:space-characters-2
    x-internal="space-characters"}.
5.  The [document
    element](https://dom.spec.whatwg.org/#document-element){#writing:document-element
    x-internal="document-element"}, in the form of an
    [`html`{#writing:the-html-element}](semantics.html#the-html-element)
    [element](#syntax-elements){#writing:syntax-elements}.
6.  Any number of
    [comments](#syntax-comments){#writing:syntax-comments-3} and [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#writing:space-characters-3
    x-internal="space-characters"}.

The various types of content mentioned above are described in the next
few sections.

In addition, there are some restrictions on how [character encoding
declarations](semantics.html#character-encoding-declaration){#writing:character-encoding-declaration}
are to be serialized, as discussed in the section on that topic.

::: note
[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#writing:space-characters-4
x-internal="space-characters"} before the
[`html`{#writing:the-html-element-2}](semantics.html#the-html-element)
element, at the start of the
[`html`{#writing:the-html-element-3}](semantics.html#the-html-element)
element and before the
[`head`{#writing:the-head-element}](semantics.html#the-head-element)
element, will be dropped when the document is parsed; [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#writing:space-characters-5
x-internal="space-characters"} *after* the
[`html`{#writing:the-html-element-4}](semantics.html#the-html-element)
element will be parsed as if it were at the end of the
[`body`{#writing:the-body-element}](sections.html#the-body-element)
element. Thus, [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#writing:space-characters-6
x-internal="space-characters"} around the [document
element](https://dom.spec.whatwg.org/#document-element){#writing:document-element-2
x-internal="document-element"} does not round-trip.

It is suggested that newlines be inserted after the DOCTYPE, after any
comments that are before the document element, after the
[`html`{#writing:the-html-element-5}](semantics.html#the-html-element)
element\'s start tag (if it is not
[omitted](#syntax-tag-omission){#writing:syntax-tag-omission}), and
after any comments that are inside the
[`html`{#writing:the-html-element-6}](semantics.html#the-html-element)
element but before the
[`head`{#writing:the-head-element-2}](semantics.html#the-head-element)
element.
:::

Many strings in the HTML syntax (e.g. the names of elements and their
attributes) are case-insensitive, but only for [ASCII upper
alphas](https://infra.spec.whatwg.org/#ascii-upper-alpha){#writing:uppercase-ascii-letters
x-internal="uppercase-ascii-letters"} and [ASCII lower
alphas](https://infra.spec.whatwg.org/#ascii-lower-alpha){#writing:lowercase-ascii-letters
x-internal="lowercase-ascii-letters"}. For convenience, in this section
this is just referred to as \"case-insensitive\".

#### [13.1.1]{.secno} The DOCTYPE[](#the-doctype){.self-link}

A [DOCTYPE]{#syntax-doctype .dfn} is a required preamble.

DOCTYPEs are required for legacy reasons. When omitted, browsers tend to
use a different rendering mode that is incompatible with some
specifications. Including the DOCTYPE in a document ensures that the
browser makes a best-effort attempt at following the relevant
specifications.

A DOCTYPE must consist of the following components, in this order:

1.  A string that is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-doctype:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for the string
    \"`<!DOCTYPE`\".
2.  One or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-doctype:space-characters
    x-internal="space-characters"}.
3.  A string that is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-doctype:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for the string
    \"`html`\".
4.  Optionally, a [DOCTYPE legacy
    string](#doctype-legacy-string){#the-doctype:doctype-legacy-string}.
5.  Zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-doctype:space-characters-2
    x-internal="space-characters"}.
6.  A U+003E GREATER-THAN SIGN character (\>).

In other words, `<!DOCTYPE html>`, case-insensitively.

------------------------------------------------------------------------

For the purposes of HTML generators that cannot output HTML markup with
the short DOCTYPE \"`<!DOCTYPE html>`\", a [DOCTYPE legacy
string]{#doctype-legacy-string .dfn} may be inserted into the DOCTYPE
(in the position defined above). This string must consist of:

1.  One or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-doctype:space-characters-3
    x-internal="space-characters"}.
2.  A string that is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-doctype:ascii-case-insensitive-3
    x-internal="ascii-case-insensitive"} match for the string
    \"`SYSTEM`\".
3.  One or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-doctype:space-characters-4
    x-internal="space-characters"}.
4.  A U+0022 QUOTATION MARK or U+0027 APOSTROPHE character (the *quote
    mark*).
5.  The literal string
    \"[`about:legacy-compat`{#the-doctype:about:legacy-compat}](urls-and-fetching.html#about:legacy-compat)\".
6.  A matching U+0022 QUOTATION MARK or U+0027 APOSTROPHE character
    (i.e. the same character as in the earlier step labeled *quote
    mark*).

In other words, `<!DOCTYPE html SYSTEM "about:legacy-compat">` or
`<!DOCTYPE html SYSTEM 'about:legacy-compat'>`, case-insensitively
except for the part in single or double quotes.

The [DOCTYPE legacy
string](#doctype-legacy-string){#the-doctype:doctype-legacy-string-2}
should not be used unless the document is generated from a system that
cannot output the shorter string.

#### [13.1.2]{.secno} Elements[](#elements-2){.self-link} {#elements-2}

There are six different kinds of [elements]{#syntax-elements .dfn}:
[void elements](#void-elements){#elements-2:void-elements}, [the
`template`
element](#the-template-element-2){#elements-2:the-template-element-2},
[raw text elements](#raw-text-elements){#elements-2:raw-text-elements},
[escapable raw text
elements](#escapable-raw-text-elements){#elements-2:escapable-raw-text-elements},
[foreign elements](#foreign-elements){#elements-2:foreign-elements}, and
[normal elements](#normal-elements){#elements-2:normal-elements}.

[Void elements]{#void-elements .dfn}
:   [`area`{#elements-2:the-area-element}](image-maps.html#the-area-element),
    [`base`{#elements-2:the-base-element}](semantics.html#the-base-element),
    [`br`{#elements-2:the-br-element}](text-level-semantics.html#the-br-element),
    [`col`{#elements-2:the-col-element}](tables.html#the-col-element),
    [`embed`{#elements-2:the-embed-element}](iframe-embed-object.html#the-embed-element),
    [`hr`{#elements-2:the-hr-element}](grouping-content.html#the-hr-element),
    [`img`{#elements-2:the-img-element}](embedded-content.html#the-img-element),
    [`input`{#elements-2:the-input-element}](input.html#the-input-element),
    [`link`{#elements-2:the-link-element}](semantics.html#the-link-element),
    [`meta`{#elements-2:the-meta-element}](semantics.html#the-meta-element),
    [`source`{#elements-2:the-source-element}](embedded-content.html#the-source-element),
    [`track`{#elements-2:the-track-element}](media.html#the-track-element),
    [`wbr`{#elements-2:the-wbr-element}](text-level-semantics.html#the-wbr-element)

[The `template` element]{#the-template-element-2 .dfn}
:   [`template`{#elements-2:the-template-element}](scripting.html#the-template-element)

[Raw text elements]{#raw-text-elements .dfn}
:   [`script`{#elements-2:the-script-element}](scripting.html#the-script-element),
    [`style`{#elements-2:the-style-element}](semantics.html#the-style-element)

[Escapable raw text elements]{#escapable-raw-text-elements .dfn}
:   [`textarea`{#elements-2:the-textarea-element}](form-elements.html#the-textarea-element),
    [`title`{#elements-2:the-title-element}](semantics.html#the-title-element)

[Foreign elements]{#foreign-elements .dfn}
:   Elements from the [MathML
    namespace](https://infra.spec.whatwg.org/#mathml-namespace){#elements-2:mathml-namespace
    x-internal="mathml-namespace"} and the [SVG
    namespace](https://infra.spec.whatwg.org/#svg-namespace){#elements-2:svg-namespace
    x-internal="svg-namespace"}.

[Normal elements]{#normal-elements .dfn}
:   All other allowed [HTML
    elements](infrastructure.html#html-elements){#elements-2:html-elements}
    are normal elements.

[Tags]{#syntax-tags .dfn} are used to delimit the start and end of
elements in the markup. [Raw
text](#raw-text-elements){#elements-2:raw-text-elements-2}, [escapable
raw
text](#escapable-raw-text-elements){#elements-2:escapable-raw-text-elements-2},
and [normal](#normal-elements){#elements-2:normal-elements-2} elements
have a [start tag](#syntax-start-tag){#elements-2:syntax-start-tag} to
indicate where they begin, and an [end
tag](#syntax-end-tag){#elements-2:syntax-end-tag} to indicate where they
end. The start and end tags of certain [normal
elements](#normal-elements){#elements-2:normal-elements-3} can be
[omitted](#syntax-tag-omission){#elements-2:syntax-tag-omission}, as
described below in the section on [optional
tags](#syntax-tag-omission){#elements-2:syntax-tag-omission-2}. Those
that cannot be omitted must not be omitted. [Void
elements](#void-elements){#elements-2:void-elements-2} only have a start
tag; end tags must not be specified for [void
elements](#void-elements){#elements-2:void-elements-3}. [Foreign
elements](#foreign-elements){#elements-2:foreign-elements-2} must either
have a start tag and an end tag, or a start tag that is marked as
self-closing, in which case they must not have an end tag.

The
[contents](dom.html#concept-html-contents){#elements-2:concept-html-contents}
of the element must be placed between just after the start tag (which
[might be implied, in certain
cases](#syntax-tag-omission){#elements-2:syntax-tag-omission-3}) and
just before the end tag (which again, [might be implied in certain
cases](#syntax-tag-omission){#elements-2:syntax-tag-omission-4}). The
exact allowed contents of each individual element depend on the [content
model](dom.html#content-models){#elements-2:content-models} of that
element, as described earlier in this specification. Elements must not
contain content that their content model disallows. In addition to the
restrictions placed on the contents by those content models, however,
the five types of elements have additional *syntactic* requirements.

[Void elements](#void-elements){#elements-2:void-elements-4} can\'t have
any contents (since there\'s no end tag, no content can be put between
the start tag and the end tag).

[The `template`
element](#the-template-element-2){#elements-2:the-template-element-2-2}
can have [template
contents](scripting.html#template-contents){#elements-2:template-contents},
but such [template
contents](scripting.html#template-contents){#elements-2:template-contents-2}
are not children of the
[`template`{#elements-2:the-template-element-3}](scripting.html#the-template-element)
element itself. Instead, they are stored in a
[`DocumentFragment`{#elements-2:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
associated with a different
[`Document`{#elements-2:document}](dom.html#document) --- without a
[browsing
context](document-sequences.html#browsing-context){#elements-2:browsing-context}
--- so as to avoid the
[`template`{#elements-2:the-template-element-4}](scripting.html#the-template-element)
contents interfering with the main
[`Document`{#elements-2:document-2}](dom.html#document). The markup for
the [template
contents](scripting.html#template-contents){#elements-2:template-contents-3}
of a
[`template`{#elements-2:the-template-element-5}](scripting.html#the-template-element)
element is placed just after the
[`template`{#elements-2:the-template-element-6}](scripting.html#the-template-element)
element\'s start tag and just before
[`template`{#elements-2:the-template-element-7}](scripting.html#the-template-element)
element\'s end tag (as with other elements), and may consist of any
[text](#syntax-text){#elements-2:syntax-text}, [character
references](#syntax-charref){#elements-2:syntax-charref},
[elements](#syntax-elements){#elements-2:syntax-elements}, and
[comments](#syntax-comments){#elements-2:syntax-comments}, but the text
must not contain the character U+003C LESS-THAN SIGN (\<) or an
[ambiguous
ampersand](#syntax-ambiguous-ampersand){#elements-2:syntax-ambiguous-ampersand}.

[Raw text elements](#raw-text-elements){#elements-2:raw-text-elements-3}
can have [text](#syntax-text){#elements-2:syntax-text-2}, though it has
[restrictions](#cdata-rcdata-restrictions) described below.

[Escapable raw text
elements](#escapable-raw-text-elements){#elements-2:escapable-raw-text-elements-3}
can have [text](#syntax-text){#elements-2:syntax-text-3} and [character
references](#syntax-charref){#elements-2:syntax-charref-2}, but the text
must not contain an [ambiguous
ampersand](#syntax-ambiguous-ampersand){#elements-2:syntax-ambiguous-ampersand-2}.
There are also [further restrictions](#cdata-rcdata-restrictions)
described below.

[Foreign elements](#foreign-elements){#elements-2:foreign-elements-3}
whose start tag is marked as self-closing can\'t have any contents
(since, again, as there\'s no end tag, no content can be put between the
start tag and the end tag). [Foreign
elements](#foreign-elements){#elements-2:foreign-elements-4} whose start
tag is *not* marked as self-closing can have
[text](#syntax-text){#elements-2:syntax-text-4}, [character
references](#syntax-charref){#elements-2:syntax-charref-3}, [CDATA
sections](#syntax-cdata){#elements-2:syntax-cdata}, other
[elements](#syntax-elements){#elements-2:syntax-elements-2}, and
[comments](#syntax-comments){#elements-2:syntax-comments-2}, but the
text must not contain the character U+003C LESS-THAN SIGN (\<) or an
[ambiguous
ampersand](#syntax-ambiguous-ampersand){#elements-2:syntax-ambiguous-ampersand-3}.

::: note
The HTML syntax does not support namespace declarations, even in
[foreign elements](#foreign-elements){#elements-2:foreign-elements-5}.

For instance, consider the following HTML fragment:

``` html
<p>
 <svg>
  <metadata>
   <!-- this is invalid -->
   <cdr:license xmlns:cdr="https://www.example.com/cdr/metadata" name="MIT"/>
  </metadata>
 </svg>
</p>
```

The innermost element, `cdr:license`, is actually in the SVG namespace,
as the \"`xmlns:cdr`\" attribute has no effect (unlike in XML). In fact,
as the comment in the fragment above says, the fragment is actually
non-conforming. This is because SVG 2 does not define any elements
called \"`cdr:license`\" in the SVG namespace.
:::

[Normal elements](#normal-elements){#elements-2:normal-elements-4} can
have [text](#syntax-text){#elements-2:syntax-text-5}, [character
references](#syntax-charref){#elements-2:syntax-charref-4}, other
[elements](#syntax-elements){#elements-2:syntax-elements-3}, and
[comments](#syntax-comments){#elements-2:syntax-comments-3}, but the
text must not contain the character U+003C LESS-THAN SIGN (\<) or an
[ambiguous
ampersand](#syntax-ambiguous-ampersand){#elements-2:syntax-ambiguous-ampersand-4}.
Some [normal elements](#normal-elements){#elements-2:normal-elements-5}
also have [yet more restrictions](#element-restrictions) on what content
they are allowed to hold, beyond the restrictions imposed by the content
model and those described in this paragraph. Those restrictions are
described below.

Tags contain a [tag name]{#syntax-tag-name .dfn}, giving the element\'s
name. HTML elements all have names that only use [ASCII
alphanumerics](https://infra.spec.whatwg.org/#ascii-alphanumeric){#elements-2:alphanumeric-ascii-characters
x-internal="alphanumeric-ascii-characters"}. In the HTML syntax, tag
names, even those for [foreign
elements](#foreign-elements){#elements-2:foreign-elements-6}, may be
written with any mix of lower- and uppercase letters that, when
converted to all-lowercase, matches the element\'s tag name; tag names
are case-insensitive.

##### [13.1.2.1]{.secno} Start tags[](#start-tags){.self-link}

[Start tags]{#syntax-start-tag .dfn} must have the following format:

1.  The first character of a start tag must be a U+003C LESS-THAN SIGN
    character (\<).
2.  The next few characters of a start tag must be the element\'s [tag
    name](#syntax-tag-name){#start-tags:syntax-tag-name}.
3.  If there are to be any attributes in the next step, there must first
    be one or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#start-tags:space-characters
    x-internal="space-characters"}.
4.  Then, the start tag may have a number of attributes, the [syntax for
    which](#syntax-attributes){#start-tags:syntax-attributes} is
    described below. Attributes must be separated from each other by one
    or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#start-tags:space-characters-2
    x-internal="space-characters"}.
5.  After the attributes, or after the [tag
    name](#syntax-tag-name){#start-tags:syntax-tag-name-2} if there are
    no attributes, there may be one or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#start-tags:space-characters-3
    x-internal="space-characters"}. (Some attributes are required to be
    followed by a space. See the [attributes
    section](#syntax-attributes){#start-tags:syntax-attributes-2}
    below.)
6.  Then, if the element is one of the [void
    elements](#void-elements){#start-tags:void-elements}, or if the
    element is a [foreign
    element](#foreign-elements){#start-tags:foreign-elements}, then
    there may be a single U+002F SOLIDUS character (/), which on
    [foreign
    elements](#foreign-elements){#start-tags:foreign-elements-2} marks
    the start tag as self-closing. On [void
    elements](#void-elements){#start-tags:void-elements-2}, it does not
    mark the start tag as self-closing but instead is unnecessary and
    has no effect of any kind. For such void elements, it should be used
    only with caution --- especially since, if directly preceded by an
    [unquoted attribute value](#unquoted), it becomes part of the
    attribute value rather than being discarded by the parser.
7.  Finally, start tags must be closed by a U+003E GREATER-THAN SIGN
    character (\>).

##### [13.1.2.2]{.secno} End tags[](#end-tags){.self-link}

[End tags]{#syntax-end-tag .dfn} must have the following format:

1.  The first character of an end tag must be a U+003C LESS-THAN SIGN
    character (\<).
2.  The second character of an end tag must be a U+002F SOLIDUS
    character (/).
3.  The next few characters of an end tag must be the element\'s [tag
    name](#syntax-tag-name){#end-tags:syntax-tag-name}.
4.  After the tag name, there may be one or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#end-tags:space-characters
    x-internal="space-characters"}.
5.  Finally, end tags must be closed by a U+003E GREATER-THAN SIGN
    character (\>).

##### [13.1.2.3]{.secno} Attributes[](#attributes-2){.self-link} {#attributes-2}

[Attributes]{#syntax-attributes .dfn} for an element are expressed
inside the element\'s start tag.

Attributes have a name and a value. [Attribute
names]{#syntax-attribute-name .dfn} must consist of one or more
characters other than
[controls](https://infra.spec.whatwg.org/#control){#attributes-2:control
x-internal="control"}, U+0020 SPACE, U+0022 (\"), U+0027 (\'), U+003E
(\>), U+002F (/), U+003D (=), and
[noncharacters](https://infra.spec.whatwg.org/#noncharacter){#attributes-2:noncharacter
x-internal="noncharacter"}. In the HTML syntax, attribute names, even
those for [foreign
elements](#foreign-elements){#attributes-2:foreign-elements}, may be
written with any mix of [ASCII
lower](https://infra.spec.whatwg.org/#ascii-lower-alpha){#attributes-2:lowercase-ascii-letters
x-internal="lowercase-ascii-letters"} and [ASCII upper
alphas](https://infra.spec.whatwg.org/#ascii-upper-alpha){#attributes-2:uppercase-ascii-letters
x-internal="uppercase-ascii-letters"}.

[Attribute values]{#syntax-attribute-value .dfn} are a mixture of
[text](#syntax-text){#attributes-2:syntax-text} and [character
references](#syntax-charref){#attributes-2:syntax-charref}, except with
the additional restriction that the text cannot contain an [ambiguous
ampersand](#syntax-ambiguous-ampersand){#attributes-2:syntax-ambiguous-ampersand}.

Attributes can be specified in four different ways:

Empty attribute syntax

:   Just the [attribute
    name](#syntax-attribute-name){#attributes-2:syntax-attribute-name}.
    The value is implicitly the empty string.

    ::: example
    In the following example, the
    [`disabled`{#attributes-2:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
    attribute is given with the empty attribute syntax:

    ``` html
    <input disabled>
    ```
    :::

    If an attribute using the empty attribute syntax is to be followed
    by another attribute, then there must be [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters
    x-internal="space-characters"} separating the two.

Unquoted attribute value syntax

:   The [attribute
    name](#syntax-attribute-name){#attributes-2:syntax-attribute-name-2},
    followed by zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-2
    x-internal="space-characters"}, followed by a single U+003D EQUALS
    SIGN character, followed by zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-3
    x-internal="space-characters"}, followed by the [attribute
    value](#syntax-attribute-value){#attributes-2:syntax-attribute-value},
    which, in addition to the requirements given above for attribute
    values, must not contain any literal [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-4
    x-internal="space-characters"}, any U+0022 QUOTATION MARK characters
    (\"), U+0027 APOSTROPHE characters (\'), U+003D EQUALS SIGN
    characters (=), U+003C LESS-THAN SIGN characters (\<), U+003E
    GREATER-THAN SIGN characters (\>), or U+0060 GRAVE ACCENT characters
    (\`), and must not be the empty string.

    ::: example
    In the following example, the
    [`value`{#attributes-2:attr-input-value}](input.html#attr-input-value)
    attribute is given with the unquoted attribute value syntax:

    ``` html
    <input value=yes>
    ```
    :::

    If an attribute using the unquoted attribute syntax is to be
    followed by another attribute or by the optional U+002F SOLIDUS
    character (/) allowed in step 6 of the [start
    tag](#syntax-start-tag){#attributes-2:syntax-start-tag} syntax
    above, then there must be [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-5
    x-internal="space-characters"} separating the two.

Single-quoted attribute value syntax

:   The [attribute
    name](#syntax-attribute-name){#attributes-2:syntax-attribute-name-3},
    followed by zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-6
    x-internal="space-characters"}, followed by a single U+003D EQUALS
    SIGN character, followed by zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-7
    x-internal="space-characters"}, followed by a single U+0027
    APOSTROPHE character (\'), followed by the [attribute
    value](#syntax-attribute-value){#attributes-2:syntax-attribute-value-2},
    which, in addition to the requirements given above for attribute
    values, must not contain any literal U+0027 APOSTROPHE characters
    (\'), and finally followed by a second single U+0027 APOSTROPHE
    character (\').

    ::: example
    In the following example, the
    [`type`{#attributes-2:attr-input-type}](input.html#attr-input-type)
    attribute is given with the single-quoted attribute value syntax:

    ``` html
    <input type='checkbox'>
    ```
    :::

    If an attribute using the single-quoted attribute syntax is to be
    followed by another attribute, then there must be [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-8
    x-internal="space-characters"} separating the two.

Double-quoted attribute value syntax

:   The [attribute
    name](#syntax-attribute-name){#attributes-2:syntax-attribute-name-4},
    followed by zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-9
    x-internal="space-characters"}, followed by a single U+003D EQUALS
    SIGN character, followed by zero or more [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-10
    x-internal="space-characters"}, followed by a single U+0022
    QUOTATION MARK character (\"), followed by the [attribute
    value](#syntax-attribute-value){#attributes-2:syntax-attribute-value-3},
    which, in addition to the requirements given above for attribute
    values, must not contain any literal U+0022 QUOTATION MARK
    characters (\"), and finally followed by a second single U+0022
    QUOTATION MARK character (\").

    ::: example
    In the following example, the
    [`name`{#attributes-2:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    attribute is given with the double-quoted attribute value syntax:

    ``` html
    <input name="be evil">
    ```
    :::

    If an attribute using the double-quoted attribute syntax is to be
    followed by another attribute, then there must be [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#attributes-2:space-characters-11
    x-internal="space-characters"} separating the two.

There must never be two or more attributes on the same start tag whose
names are an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#attributes-2:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for each other.

------------------------------------------------------------------------

When a [foreign
element](#foreign-elements){#attributes-2:foreign-elements-2} has one of
the namespaced attributes given by the local name and namespace of the
first and second cells of a row from the following table, it must be
written using the name given by the third cell from the same row.

Local name

Namespace

Attribute name

`actuate`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#attributes-2:xlink-namespace
x-internal="xlink-namespace"}

`xlink:actuate`

`arcrole`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#attributes-2:xlink-namespace-2
x-internal="xlink-namespace"}

`xlink:arcrole`

`href`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#attributes-2:xlink-namespace-3
x-internal="xlink-namespace"}

`xlink:href`

`role`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#attributes-2:xlink-namespace-4
x-internal="xlink-namespace"}

`xlink:role`

`show`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#attributes-2:xlink-namespace-5
x-internal="xlink-namespace"}

`xlink:show`

`title`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#attributes-2:xlink-namespace-6
x-internal="xlink-namespace"}

`xlink:title`

`type`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#attributes-2:xlink-namespace-7
x-internal="xlink-namespace"}

`xlink:type`

`lang`

[XML
namespace](https://infra.spec.whatwg.org/#xml-namespace){#attributes-2:xml-namespace
x-internal="xml-namespace"}

`xml:lang`

`space`

[XML
namespace](https://infra.spec.whatwg.org/#xml-namespace){#attributes-2:xml-namespace-2
x-internal="xml-namespace"}

`xml:space`

`xmlns`

[XMLNS
namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#attributes-2:xmlns-namespace
x-internal="xmlns-namespace"}

`xmlns`

`xlink`

[XMLNS
namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#attributes-2:xmlns-namespace-2
x-internal="xmlns-namespace"}

`xmlns:xlink`

No other namespaced attribute can be expressed in [the HTML
syntax](#syntax){#attributes-2:syntax}.

Whether the attributes in the table above are conforming or not is
defined by other specifications (e.g. SVG 2 and MathML); this section
only describes the syntax rules if the attributes are serialized using
the HTML syntax.

##### [13.1.2.4]{.secno} Optional tags[](#optional-tags){.self-link}

Certain tags can be [omitted]{#syntax-tag-omission .dfn}.

Omitting an element\'s [start
tag](#syntax-start-tag){#optional-tags:syntax-start-tag} in the
situations described below does not mean the element is not present; it
is implied, but it is still there. For example, an HTML document always
has a root
[`html`{#optional-tags:the-html-element}](semantics.html#the-html-element)
element, even if the string `<html>` doesn\'t appear anywhere in the
markup.

An
[`html`{#optional-tags:the-html-element-2}](semantics.html#the-html-element)
element\'s [start
tag](#syntax-start-tag){#optional-tags:syntax-start-tag-2} may be
omitted if the first thing inside the
[`html`{#optional-tags:the-html-element-3}](semantics.html#the-html-element)
element is not a
[comment](#syntax-comments){#optional-tags:syntax-comments}.

::: example
For example, in the following case it\'s ok to remove the \"`<html>`\"
tag:

``` html
<!DOCTYPE HTML>
<html>
  <head>
    <title>Hello</title>
  </head>
  <body>
    <p>Welcome to this example.</p>
  </body>
</html>
```

Doing so would make the document look like this:

``` html
<!DOCTYPE HTML>

  <head>
    <title>Hello</title>
  </head>
  <body>
    <p>Welcome to this example.</p>
  </body>
</html>
```

This has the exact same DOM. In particular, note that whitespace around
the [document
element](https://dom.spec.whatwg.org/#document-element){#optional-tags:document-element
x-internal="document-element"} is ignored by the parser. The following
example would also have the exact same DOM:

``` html
<!DOCTYPE HTML><head>
    <title>Hello</title>
  </head>
  <body>
    <p>Welcome to this example.</p>
  </body>
</html>
```

However, in the following example, removing the start tag moves the
comment to before the
[`html`{#optional-tags:the-html-element-4}](semantics.html#the-html-element)
element:

``` html
<!DOCTYPE HTML>
<html>
  <!-- where is this comment in the DOM? -->
  <head>
    <title>Hello</title>
  </head>
  <body>
    <p>Welcome to this example.</p>
  </body>
</html>
```

With the tag removed, the document actually turns into the same as this:

``` html
<!DOCTYPE HTML>
<!-- where is this comment in the DOM? -->
<html>
  <head>
    <title>Hello</title>
  </head>
  <body>
    <p>Welcome to this example.</p>
  </body>
</html>
```

This is why the tag can only be removed if it is not followed by a
comment: removing the tag when there is a comment there changes the
document\'s resulting parse tree. Of course, if the position of the
comment does not matter, then the tag can be omitted, as if the comment
had been moved to before the start tag in the first place.
:::

An
[`html`{#optional-tags:the-html-element-5}](semantics.html#the-html-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag} may
be omitted if the
[`html`{#optional-tags:the-html-element-6}](semantics.html#the-html-element)
element is not immediately followed by a
[comment](#syntax-comments){#optional-tags:syntax-comments-2}.

A
[`head`{#optional-tags:the-head-element}](semantics.html#the-head-element)
element\'s [start
tag](#syntax-start-tag){#optional-tags:syntax-start-tag-3} may be
omitted if the element is empty, or if the first thing inside the
[`head`{#optional-tags:the-head-element-2}](semantics.html#the-head-element)
element is an element.

A
[`head`{#optional-tags:the-head-element-3}](semantics.html#the-head-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-2}
may be omitted if the
[`head`{#optional-tags:the-head-element-4}](semantics.html#the-head-element)
element is not immediately followed by [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#optional-tags:space-characters
x-internal="space-characters"} or a
[comment](#syntax-comments){#optional-tags:syntax-comments-3}.

A
[`body`{#optional-tags:the-body-element}](sections.html#the-body-element)
element\'s [start
tag](#syntax-start-tag){#optional-tags:syntax-start-tag-4} may be
omitted if the element is empty, or if the first thing inside the
[`body`{#optional-tags:the-body-element-2}](sections.html#the-body-element)
element is not [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#optional-tags:space-characters-2
x-internal="space-characters"} or a
[comment](#syntax-comments){#optional-tags:syntax-comments-4}, except if
the first thing inside the
[`body`{#optional-tags:the-body-element-3}](sections.html#the-body-element)
element is a
[`meta`{#optional-tags:the-meta-element}](semantics.html#the-meta-element),
[`noscript`{#optional-tags:the-noscript-element}](scripting.html#the-noscript-element),
[`link`{#optional-tags:the-link-element}](semantics.html#the-link-element),
[`script`{#optional-tags:the-script-element}](scripting.html#the-script-element),
[`style`{#optional-tags:the-style-element}](semantics.html#the-style-element),
or
[`template`{#optional-tags:the-template-element}](scripting.html#the-template-element)
element.

A
[`body`{#optional-tags:the-body-element-4}](sections.html#the-body-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-3}
may be omitted if the
[`body`{#optional-tags:the-body-element-5}](sections.html#the-body-element)
element is not immediately followed by a
[comment](#syntax-comments){#optional-tags:syntax-comments-5}.

::: example
Note that in the example above, the
[`head`{#optional-tags:the-head-element-5}](semantics.html#the-head-element)
element start and end tags, and the
[`body`{#optional-tags:the-body-element-6}](sections.html#the-body-element)
element start tag, can\'t be omitted, because they are surrounded by
whitespace:

``` html
<!DOCTYPE HTML>
<html>
  <head>
    <title>Hello</title>
  </head>
  <body>
    <p>Welcome to this example.</p>
  </body>
</html>
```

(The
[`body`{#optional-tags:the-body-element-7}](sections.html#the-body-element)
and
[`html`{#optional-tags:the-html-element-7}](semantics.html#the-html-element)
element end tags could be omitted without trouble; any spaces after
those get parsed into the
[`body`{#optional-tags:the-body-element-8}](sections.html#the-body-element)
element anyway.)

Usually, however, whitespace isn\'t an issue. If we first remove the
whitespace we don\'t care about:

``` html
<!DOCTYPE HTML><html><head><title>Hello</title></head><body><p>Welcome to this example.</p></body></html>
```

Then we can omit a number of tags without affecting the DOM:

``` html
<!DOCTYPE HTML><title>Hello</title><p>Welcome to this example.</p>
```

At that point, we can also add some whitespace back:

``` html
<!DOCTYPE HTML>
<title>Hello</title>
<p>Welcome to this example.</p>
```

This would be equivalent to this document, with the omitted tags shown
in their parser-implied positions; the only whitespace text node that
results from this is the newline at the end of the
[`head`{#optional-tags:the-head-element-6}](semantics.html#the-head-element)
element:

``` html
<!DOCTYPE HTML>
<html><head><title>Hello</title>
</head><body><p>Welcome to this example.</p></body></html>
```
:::

An
[`li`{#optional-tags:the-li-element}](grouping-content.html#the-li-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-4}
may be omitted if the
[`li`{#optional-tags:the-li-element-2}](grouping-content.html#the-li-element)
element is immediately followed by another
[`li`{#optional-tags:the-li-element-3}](grouping-content.html#the-li-element)
element or if there is no more content in the parent element.

A
[`dt`{#optional-tags:the-dt-element}](grouping-content.html#the-dt-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-5}
may be omitted if the
[`dt`{#optional-tags:the-dt-element-2}](grouping-content.html#the-dt-element)
element is immediately followed by another
[`dt`{#optional-tags:the-dt-element-3}](grouping-content.html#the-dt-element)
element or a
[`dd`{#optional-tags:the-dd-element}](grouping-content.html#the-dd-element)
element.

A
[`dd`{#optional-tags:the-dd-element-2}](grouping-content.html#the-dd-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-6}
may be omitted if the
[`dd`{#optional-tags:the-dd-element-3}](grouping-content.html#the-dd-element)
element is immediately followed by another
[`dd`{#optional-tags:the-dd-element-4}](grouping-content.html#the-dd-element)
element or a
[`dt`{#optional-tags:the-dt-element-4}](grouping-content.html#the-dt-element)
element, or if there is no more content in the parent element.

A
[`p`{#optional-tags:the-p-element}](grouping-content.html#the-p-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-7}
may be omitted if the
[`p`{#optional-tags:the-p-element-2}](grouping-content.html#the-p-element)
element is immediately followed by an
[`address`{#optional-tags:the-address-element}](sections.html#the-address-element),
[`article`{#optional-tags:the-article-element}](sections.html#the-article-element),
[`aside`{#optional-tags:the-aside-element}](sections.html#the-aside-element),
[`blockquote`{#optional-tags:the-blockquote-element}](grouping-content.html#the-blockquote-element),
[`details`{#optional-tags:the-details-element}](interactive-elements.html#the-details-element),
[`dialog`{#optional-tags:the-dialog-element}](interactive-elements.html#the-dialog-element),
[`div`{#optional-tags:the-div-element}](grouping-content.html#the-div-element),
[`dl`{#optional-tags:the-dl-element}](grouping-content.html#the-dl-element),
[`fieldset`{#optional-tags:the-fieldset-element}](form-elements.html#the-fieldset-element),
[`figcaption`{#optional-tags:the-figcaption-element}](grouping-content.html#the-figcaption-element),
[`figure`{#optional-tags:the-figure-element}](grouping-content.html#the-figure-element),
[`footer`{#optional-tags:the-footer-element}](sections.html#the-footer-element),
[`form`{#optional-tags:the-form-element}](forms.html#the-form-element),
[`h1`{#optional-tags:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h2`{#optional-tags:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h3`{#optional-tags:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h4`{#optional-tags:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h5`{#optional-tags:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`h6`{#optional-tags:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
[`header`{#optional-tags:the-header-element}](sections.html#the-header-element),
[`hgroup`{#optional-tags:the-hgroup-element}](sections.html#the-hgroup-element),
[`hr`{#optional-tags:the-hr-element}](grouping-content.html#the-hr-element),
[`main`{#optional-tags:the-main-element}](grouping-content.html#the-main-element),
[`menu`{#optional-tags:the-menu-element}](grouping-content.html#the-menu-element),
[`nav`{#optional-tags:the-nav-element}](sections.html#the-nav-element),
[`ol`{#optional-tags:the-ol-element}](grouping-content.html#the-ol-element),
[`p`{#optional-tags:the-p-element-3}](grouping-content.html#the-p-element),
[`pre`{#optional-tags:the-pre-element}](grouping-content.html#the-pre-element),
[`search`{#optional-tags:the-search-element}](grouping-content.html#the-search-element),
[`section`{#optional-tags:the-section-element}](sections.html#the-section-element),
[`table`{#optional-tags:the-table-element}](tables.html#the-table-element),
or
[`ul`{#optional-tags:the-ul-element}](grouping-content.html#the-ul-element)
element, or if there is no more content in the parent element and the
parent element is an [HTML
element](infrastructure.html#html-elements){#optional-tags:html-elements}
that is not an
[`a`{#optional-tags:the-a-element}](text-level-semantics.html#the-a-element),
[`audio`{#optional-tags:the-audio-element}](media.html#the-audio-element),
[`del`{#optional-tags:the-del-element}](edits.html#the-del-element),
[`ins`{#optional-tags:the-ins-element}](edits.html#the-ins-element),
[`map`{#optional-tags:the-map-element}](image-maps.html#the-map-element),
[`noscript`{#optional-tags:the-noscript-element-2}](scripting.html#the-noscript-element),
or
[`video`{#optional-tags:the-video-element}](media.html#the-video-element)
element, or an [autonomous custom
element](custom-elements.html#autonomous-custom-element){#optional-tags:autonomous-custom-element}.

::: example
We can thus simplify the earlier example further:

``` html
<!DOCTYPE HTML><title>Hello</title><p>Welcome to this example.
```
:::

An
[`rt`{#optional-tags:the-rt-element}](text-level-semantics.html#the-rt-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-8}
may be omitted if the
[`rt`{#optional-tags:the-rt-element-2}](text-level-semantics.html#the-rt-element)
element is immediately followed by an
[`rt`{#optional-tags:the-rt-element-3}](text-level-semantics.html#the-rt-element)
or
[`rp`{#optional-tags:the-rp-element}](text-level-semantics.html#the-rp-element)
element, or if there is no more content in the parent element.

An
[`rp`{#optional-tags:the-rp-element-2}](text-level-semantics.html#the-rp-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-9}
may be omitted if the
[`rp`{#optional-tags:the-rp-element-3}](text-level-semantics.html#the-rp-element)
element is immediately followed by an
[`rt`{#optional-tags:the-rt-element-4}](text-level-semantics.html#the-rt-element)
or
[`rp`{#optional-tags:the-rp-element-4}](text-level-semantics.html#the-rp-element)
element, or if there is no more content in the parent element.

An
[`optgroup`{#optional-tags:the-optgroup-element}](form-elements.html#the-optgroup-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-10}
may be omitted if the
[`optgroup`{#optional-tags:the-optgroup-element-2}](form-elements.html#the-optgroup-element)
element is immediately followed by another
[`optgroup`{#optional-tags:the-optgroup-element-3}](form-elements.html#the-optgroup-element)
element, if it is immediately followed by an
[`hr`{#optional-tags:the-hr-element-2}](grouping-content.html#the-hr-element)
element, or if there is no more content in the parent element.

An
[`option`{#optional-tags:the-option-element}](form-elements.html#the-option-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-11}
may be omitted if the
[`option`{#optional-tags:the-option-element-2}](form-elements.html#the-option-element)
element is immediately followed by another
[`option`{#optional-tags:the-option-element-3}](form-elements.html#the-option-element)
element, if it is immediately followed by an
[`optgroup`{#optional-tags:the-optgroup-element-4}](form-elements.html#the-optgroup-element)
element, if it is immediately followed by an
[`hr`{#optional-tags:the-hr-element-3}](grouping-content.html#the-hr-element)
element, or if there is no more content in the parent element.

A
[`colgroup`{#optional-tags:the-colgroup-element}](tables.html#the-colgroup-element)
element\'s [start
tag](#syntax-start-tag){#optional-tags:syntax-start-tag-5} may be
omitted if the first thing inside the
[`colgroup`{#optional-tags:the-colgroup-element-2}](tables.html#the-colgroup-element)
element is a
[`col`{#optional-tags:the-col-element}](tables.html#the-col-element)
element, and if the element is not immediately preceded by another
[`colgroup`{#optional-tags:the-colgroup-element-3}](tables.html#the-colgroup-element)
element whose [end
tag](#syntax-end-tag){#optional-tags:syntax-end-tag-12} has been
omitted. (It can\'t be omitted if the element is empty.)

A
[`colgroup`{#optional-tags:the-colgroup-element-4}](tables.html#the-colgroup-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-13}
may be omitted if the
[`colgroup`{#optional-tags:the-colgroup-element-5}](tables.html#the-colgroup-element)
element is not immediately followed by [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#optional-tags:space-characters-3
x-internal="space-characters"} or a
[comment](#syntax-comments){#optional-tags:syntax-comments-6}.

A
[`caption`{#optional-tags:the-caption-element}](tables.html#the-caption-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-14}
may be omitted if the
[`caption`{#optional-tags:the-caption-element-2}](tables.html#the-caption-element)
element is not immediately followed by [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#optional-tags:space-characters-4
x-internal="space-characters"} or a
[comment](#syntax-comments){#optional-tags:syntax-comments-7}.

A
[`thead`{#optional-tags:the-thead-element}](tables.html#the-thead-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-15}
may be omitted if the
[`thead`{#optional-tags:the-thead-element-2}](tables.html#the-thead-element)
element is immediately followed by a
[`tbody`{#optional-tags:the-tbody-element}](tables.html#the-tbody-element)
or
[`tfoot`{#optional-tags:the-tfoot-element}](tables.html#the-tfoot-element)
element.

A
[`tbody`{#optional-tags:the-tbody-element-2}](tables.html#the-tbody-element)
element\'s [start
tag](#syntax-start-tag){#optional-tags:syntax-start-tag-6} may be
omitted if the first thing inside the
[`tbody`{#optional-tags:the-tbody-element-3}](tables.html#the-tbody-element)
element is a
[`tr`{#optional-tags:the-tr-element}](tables.html#the-tr-element)
element, and if the element is not immediately preceded by a
[`tbody`{#optional-tags:the-tbody-element-4}](tables.html#the-tbody-element),
[`thead`{#optional-tags:the-thead-element-3}](tables.html#the-thead-element),
or
[`tfoot`{#optional-tags:the-tfoot-element-2}](tables.html#the-tfoot-element)
element whose [end
tag](#syntax-end-tag){#optional-tags:syntax-end-tag-16} has been
omitted. (It can\'t be omitted if the element is empty.)

A
[`tbody`{#optional-tags:the-tbody-element-5}](tables.html#the-tbody-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-17}
may be omitted if the
[`tbody`{#optional-tags:the-tbody-element-6}](tables.html#the-tbody-element)
element is immediately followed by a
[`tbody`{#optional-tags:the-tbody-element-7}](tables.html#the-tbody-element)
or
[`tfoot`{#optional-tags:the-tfoot-element-3}](tables.html#the-tfoot-element)
element, or if there is no more content in the parent element.

A
[`tfoot`{#optional-tags:the-tfoot-element-4}](tables.html#the-tfoot-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-18}
may be omitted if there is no more content in the parent element.

A [`tr`{#optional-tags:the-tr-element-2}](tables.html#the-tr-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-19}
may be omitted if the
[`tr`{#optional-tags:the-tr-element-3}](tables.html#the-tr-element)
element is immediately followed by another
[`tr`{#optional-tags:the-tr-element-4}](tables.html#the-tr-element)
element, or if there is no more content in the parent element.

A [`td`{#optional-tags:the-td-element}](tables.html#the-td-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-20}
may be omitted if the
[`td`{#optional-tags:the-td-element-2}](tables.html#the-td-element)
element is immediately followed by a
[`td`{#optional-tags:the-td-element-3}](tables.html#the-td-element) or
[`th`{#optional-tags:the-th-element}](tables.html#the-th-element)
element, or if there is no more content in the parent element.

A [`th`{#optional-tags:the-th-element-2}](tables.html#the-th-element)
element\'s [end tag](#syntax-end-tag){#optional-tags:syntax-end-tag-21}
may be omitted if the
[`th`{#optional-tags:the-th-element-3}](tables.html#the-th-element)
element is immediately followed by a
[`td`{#optional-tags:the-td-element-4}](tables.html#the-td-element) or
[`th`{#optional-tags:the-th-element-4}](tables.html#the-th-element)
element, or if there is no more content in the parent element.

::: example
The ability to omit all these table-related tags makes table markup much
terser.

Take this example:

``` html
<table>
 <caption>37547 TEE Electric Powered Rail Car Train Functions (Abbreviated)</caption>
 <colgroup><col><col><col></colgroup>
 <thead>
  <tr>
   <th>Function</th>
   <th>Control Unit</th>
   <th>Central Station</th>
  </tr>
 </thead>
 <tbody>
  <tr>
   <td>Headlights</td>
   <td>✔</td>
   <td>✔</td>
  </tr>
  <tr>
   <td>Interior Lights</td>
   <td>✔</td>
   <td>✔</td>
  </tr>
  <tr>
   <td>Electric locomotive operating sounds</td>
   <td>✔</td>
   <td>✔</td>
  </tr>
  <tr>
   <td>Engineer's cab lighting</td>
   <td></td>
   <td>✔</td>
  </tr>
  <tr>
   <td>Station Announcements - Swiss</td>
   <td></td>
   <td>✔</td>
  </tr>
 </tbody>
</table>
```

The exact same table, modulo some whitespace differences, could be
marked up as follows:

``` html
<table>
 <caption>37547 TEE Electric Powered Rail Car Train Functions (Abbreviated)
 <colgroup><col><col><col>
 <thead>
  <tr>
   <th>Function
   <th>Control Unit
   <th>Central Station
 <tbody>
  <tr>
   <td>Headlights
   <td>✔
   <td>✔
  <tr>
   <td>Interior Lights
   <td>✔
   <td>✔
  <tr>
   <td>Electric locomotive operating sounds
   <td>✔
   <td>✔
  <tr>
   <td>Engineer's cab lighting
   <td>
   <td>✔
  <tr>
   <td>Station Announcements - Swiss
   <td>
   <td>✔
</table>
```

Since the cells take up much less room this way, this can be made even
terser by having each row on one line:

``` html
<table>
 <caption>37547 TEE Electric Powered Rail Car Train Functions (Abbreviated)
 <colgroup><col><col><col>
 <thead>
  <tr> <th>Function                              <th>Control Unit     <th>Central Station
 <tbody>
  <tr> <td>Headlights                            <td>✔                <td>✔
  <tr> <td>Interior Lights                       <td>✔                <td>✔
  <tr> <td>Electric locomotive operating sounds  <td>✔                <td>✔
  <tr> <td>Engineer's cab lighting               <td>                 <td>✔
  <tr> <td>Station Announcements - Swiss         <td>                 <td>✔
</table>
```

The only differences between these tables, at the DOM level, is with the
precise position of the (in any case semantically-neutral) whitespace.
:::

**However**, a [start
tag](#syntax-start-tag){#optional-tags:syntax-start-tag-7} must never be
omitted if it has any attributes.

::: example
Returning to the earlier example with all the whitespace removed and
then all the optional tags removed:

``` html
<!DOCTYPE HTML><title>Hello</title><p>Welcome to this example.
```

If the
[`body`{#optional-tags:the-body-element-9}](sections.html#the-body-element)
element in this example had to have a
[`class`{#optional-tags:classes}](dom.html#classes) attribute and the
[`html`{#optional-tags:the-html-element-8}](semantics.html#the-html-element)
element had to have a
[`lang`{#optional-tags:attr-lang}](dom.html#attr-lang) attribute, the
markup would have to become:

``` html
<!DOCTYPE HTML><html lang="en"><title>Hello</title><body class="demo"><p>Welcome to this example.
```
:::

This section assumes that the document is conforming, in particular,
that there are no [content
model](dom.html#content-models){#optional-tags:content-models}
violations. Omitting tags in the fashion described in this section in a
document that does not conform to the [content
models](dom.html#content-models){#optional-tags:content-models-2}
described in this specification is likely to result in unexpected DOM
differences (this is, in part, what the content models are designed to
avoid).

##### [13.1.2.5]{.secno} Restrictions on content models[](#element-restrictions){.self-link} {#element-restrictions}

For historical reasons, certain elements have extra restrictions beyond
even the restrictions given by their content model.

A
[`table`{#element-restrictions:the-table-element}](tables.html#the-table-element)
element must not contain
[`tr`{#element-restrictions:the-tr-element}](tables.html#the-tr-element)
elements, even though these elements are technically allowed inside
[`table`{#element-restrictions:the-table-element-2}](tables.html#the-table-element)
elements according to the content models described in this
specification. (If a
[`tr`{#element-restrictions:the-tr-element-2}](tables.html#the-tr-element)
element is put inside a
[`table`{#element-restrictions:the-table-element-3}](tables.html#the-table-element)
in the markup, it will in fact imply a
[`tbody`{#element-restrictions:the-tbody-element}](tables.html#the-tbody-element)
start tag before it.)

A single
[newline](#syntax-newlines){#element-restrictions:syntax-newlines} may
be placed immediately after the [start
tag](#syntax-start-tag){#element-restrictions:syntax-start-tag} of
[`pre`{#element-restrictions:the-pre-element}](grouping-content.html#the-pre-element)
and
[`textarea`{#element-restrictions:the-textarea-element}](form-elements.html#the-textarea-element)
elements. This does not affect the processing of the element. The
otherwise optional
[newline](#syntax-newlines){#element-restrictions:syntax-newlines-2}
*must* be included if the element\'s contents themselves start with a
[newline](#syntax-newlines){#element-restrictions:syntax-newlines-3}
(because otherwise the leading newline in the contents would be treated
like the optional newline, and ignored).

::: example
The following two
[`pre`{#element-restrictions:the-pre-element-2}](grouping-content.html#the-pre-element)
blocks are equivalent:

``` html
<pre>Hello</pre>
```

``` html
<pre>
Hello</pre>
```
:::

##### [13.1.2.6]{.secno} Restrictions on the contents of raw text and escapable raw text elements[](#cdata-rcdata-restrictions){.self-link} {#cdata-rcdata-restrictions}

The text in [raw
text](#raw-text-elements){#cdata-rcdata-restrictions:raw-text-elements}
and [escapable raw text
elements](#escapable-raw-text-elements){#cdata-rcdata-restrictions:escapable-raw-text-elements}
must not contain any occurrences of the string \"`</`\" (U+003C
LESS-THAN SIGN, U+002F SOLIDUS) followed by characters that
case-insensitively match the tag name of the element followed by one of
U+0009 CHARACTER TABULATION (tab), U+000A LINE FEED (LF), U+000C FORM
FEED (FF), U+000D CARRIAGE RETURN (CR), U+0020 SPACE, U+003E
GREATER-THAN SIGN (\>), or U+002F SOLIDUS (/).

#### [13.1.3]{.secno} Text[](#text-2){.self-link} {#text-2}

[Text]{#syntax-text .dfn} is allowed inside elements, attribute values,
and comments. Extra constraints are placed on what is and what is not
allowed in text based on where the text is to be put, as described in
the other sections.

##### [13.1.3.1]{.secno} Newlines[](#newlines){.self-link}

[Newlines]{#syntax-newlines .dfn} in HTML may be represented either as
U+000D CARRIAGE RETURN (CR) characters, U+000A LINE FEED (LF)
characters, or pairs of U+000D CARRIAGE RETURN (CR), U+000A LINE FEED
(LF) characters in that order.

Where [character references](#syntax-charref){#newlines:syntax-charref}
are allowed, a character reference of a U+000A LINE FEED (LF) character
(but not a U+000D CARRIAGE RETURN (CR) character) also represents a
[newline](#syntax-newlines){#newlines:syntax-newlines}.

#### [13.1.4]{.secno} Character references[](#character-references){.self-link}

In certain cases described in other sections,
[text](#syntax-text){#character-references:syntax-text} may be mixed
with [character references]{#syntax-charref .dfn}. These can be used to
escape characters that couldn\'t otherwise legally be included in
[text](#syntax-text){#character-references:syntax-text-2}.

Character references must start with a U+0026 AMPERSAND character (&).
Following this, there are three possible kinds of character references:

Named character references
:   The ampersand must be followed by one of the names given in the
    [named character
    references](named-characters.html#named-character-references){#character-references:named-character-references}
    section, using the same case. The name must be one that is
    terminated by a U+003B SEMICOLON character (;).

Decimal numeric character reference
:   The ampersand must be followed by a U+0023 NUMBER SIGN character
    (#), followed by one or more [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#character-references:ascii-digits
    x-internal="ascii-digits"}, representing a base-ten integer that
    corresponds to a code point that is allowed according to the
    definition below. The digits must then be followed by a U+003B
    SEMICOLON character (;).

Hexadecimal numeric character reference
:   The ampersand must be followed by a U+0023 NUMBER SIGN character
    (#), which must be followed by either a U+0078 LATIN SMALL LETTER X
    character (x) or a U+0058 LATIN CAPITAL LETTER X character (X),
    which must then be followed by one or more [ASCII hex
    digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#character-references:ascii-hex-digits
    x-internal="ascii-hex-digits"}, representing a hexadecimal integer
    that corresponds to a code point that is allowed according to the
    definition below. The digits must then be followed by a U+003B
    SEMICOLON character (;).

The numeric character reference forms described above are allowed to
reference any code point excluding U+000D CR,
[noncharacters](https://infra.spec.whatwg.org/#noncharacter){#character-references:noncharacter
x-internal="noncharacter"}, and
[controls](https://infra.spec.whatwg.org/#control){#character-references:control
x-internal="control"} other than [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#character-references:space-characters
x-internal="space-characters"}.

An [ambiguous ampersand]{#syntax-ambiguous-ampersand .dfn} is a U+0026
AMPERSAND character (&) that is followed by one or more [ASCII
alphanumerics](https://infra.spec.whatwg.org/#ascii-alphanumeric){#character-references:alphanumeric-ascii-characters
x-internal="alphanumeric-ascii-characters"}, followed by a U+003B
SEMICOLON character (;), where these characters do not match any of the
names given in the [named character
references](named-characters.html#named-character-references){#character-references:named-character-references-2}
section.

#### [13.1.5]{.secno} CDATA sections[](#cdata-sections){.self-link}

[CDATA sections]{#syntax-cdata .dfn} must consist of the following
components, in this order:

1.  The string \"`<![CDATA[`\".
2.  Optionally, [text](#syntax-text){#cdata-sections:syntax-text}, with
    the additional restriction that the text must not contain the string
    \"`]]>`\".
3.  The string \"`]]>`\".

::: example
CDATA sections can only be used in foreign content (MathML or SVG). In
this example, a CDATA section is used to escape the contents of a
[MathML
`ms`](https://w3c.github.io/mathml-core/#string-literal-ms){#cdata-sections:mathml-ms
x-internal="mathml-ms"} element:

``` html
<p>You can add a string to a number, but this stringifies the number:</p>
<math>
 <ms><![CDATA[x<y]]></ms>
 <mo>+</mo>
 <mn>3</mn>
 <mo>=</mo>
 <ms><![CDATA[x<y3]]></ms>
</math>
```
:::

#### [13.1.6]{.secno} Comments[](#comments){.self-link}

[Comments]{#syntax-comments .dfn} must have the following format:

1.  The string \"`<!--`\".
2.  Optionally, [text](#syntax-text){#comments:syntax-text}, with the
    additional restriction that the text must not start with the string
    \"`>`\", nor start with the string \"`->`\", nor contain the strings
    \"`<!--`\", \"`-->`\", or \"`--!>`\", nor end with the string
    \"`<!-`\".
3.  The string \"`-->`\".

The [text](#syntax-text){#comments:syntax-text-2} is allowed to end with
the string \"`<!`\", as in `<!--My favorite operators are > and <!-->`.

[← 12 Web storage](webstorage.html) --- [Table of Contents](./) ---
[13.2 Parsing HTML documents →](parsing.html)
