HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 13 The HTML syntax](syntax.html) --- [Table of Contents](./) ---
[13.5 Named character references →](named-characters.html)

1.  ::: {#toc-syntax}
    1.  [[13.2]{.secno} Parsing HTML documents](parsing.html#parsing)
        1.  [[13.2.1]{.secno} Overview of the parsing
            model](parsing.html#overview-of-the-parsing-model)
        2.  [[13.2.2]{.secno} Parse errors](parsing.html#parse-errors)
        3.  [[13.2.3]{.secno} The input byte
            stream](parsing.html#the-input-byte-stream)
            1.  [[13.2.3.1]{.secno} Parsing with a known character
                encoding](parsing.html#parsing-with-a-known-character-encoding)
            2.  [[13.2.3.2]{.secno} Determining the character
                encoding](parsing.html#determining-the-character-encoding)
            3.  [[13.2.3.3]{.secno} Character
                encodings](parsing.html#character-encodings)
            4.  [[13.2.3.4]{.secno} Changing the encoding while
                parsing](parsing.html#changing-the-encoding-while-parsing)
            5.  [[13.2.3.5]{.secno} Preprocessing the input
                stream](parsing.html#preprocessing-the-input-stream)
        4.  [[13.2.4]{.secno} Parse state](parsing.html#parse-state)
            1.  [[13.2.4.1]{.secno} The insertion
                mode](parsing.html#the-insertion-mode)
            2.  [[13.2.4.2]{.secno} The stack of open
                elements](parsing.html#the-stack-of-open-elements)
            3.  [[13.2.4.3]{.secno} The list of active formatting
                elements](parsing.html#the-list-of-active-formatting-elements)
            4.  [[13.2.4.4]{.secno} The element
                pointers](parsing.html#the-element-pointers)
            5.  [[13.2.4.5]{.secno} Other parsing state
                flags](parsing.html#other-parsing-state-flags)
        5.  [[13.2.5]{.secno} Tokenization](parsing.html#tokenization)
            1.  [[13.2.5.1]{.secno} Data state](parsing.html#data-state)
            2.  [[13.2.5.2]{.secno} RCDATA
                state](parsing.html#rcdata-state)
            3.  [[13.2.5.3]{.secno} RAWTEXT
                state](parsing.html#rawtext-state)
            4.  [[13.2.5.4]{.secno} Script data
                state](parsing.html#script-data-state)
            5.  [[13.2.5.5]{.secno} PLAINTEXT
                state](parsing.html#plaintext-state)
            6.  [[13.2.5.6]{.secno} Tag open
                state](parsing.html#tag-open-state)
            7.  [[13.2.5.7]{.secno} End tag open
                state](parsing.html#end-tag-open-state)
            8.  [[13.2.5.8]{.secno} Tag name
                state](parsing.html#tag-name-state)
            9.  [[13.2.5.9]{.secno} RCDATA less-than sign
                state](parsing.html#rcdata-less-than-sign-state)
            10. [[13.2.5.10]{.secno} RCDATA end tag open
                state](parsing.html#rcdata-end-tag-open-state)
            11. [[13.2.5.11]{.secno} RCDATA end tag name
                state](parsing.html#rcdata-end-tag-name-state)
            12. [[13.2.5.12]{.secno} RAWTEXT less-than sign
                state](parsing.html#rawtext-less-than-sign-state)
            13. [[13.2.5.13]{.secno} RAWTEXT end tag open
                state](parsing.html#rawtext-end-tag-open-state)
            14. [[13.2.5.14]{.secno} RAWTEXT end tag name
                state](parsing.html#rawtext-end-tag-name-state)
            15. [[13.2.5.15]{.secno} Script data less-than sign
                state](parsing.html#script-data-less-than-sign-state)
            16. [[13.2.5.16]{.secno} Script data end tag open
                state](parsing.html#script-data-end-tag-open-state)
            17. [[13.2.5.17]{.secno} Script data end tag name
                state](parsing.html#script-data-end-tag-name-state)
            18. [[13.2.5.18]{.secno} Script data escape start
                state](parsing.html#script-data-escape-start-state)
            19. [[13.2.5.19]{.secno} Script data escape start dash
                state](parsing.html#script-data-escape-start-dash-state)
            20. [[13.2.5.20]{.secno} Script data escaped
                state](parsing.html#script-data-escaped-state)
            21. [[13.2.5.21]{.secno} Script data escaped dash
                state](parsing.html#script-data-escaped-dash-state)
            22. [[13.2.5.22]{.secno} Script data escaped dash dash
                state](parsing.html#script-data-escaped-dash-dash-state)
            23. [[13.2.5.23]{.secno} Script data escaped less-than sign
                state](parsing.html#script-data-escaped-less-than-sign-state)
            24. [[13.2.5.24]{.secno} Script data escaped end tag open
                state](parsing.html#script-data-escaped-end-tag-open-state)
            25. [[13.2.5.25]{.secno} Script data escaped end tag name
                state](parsing.html#script-data-escaped-end-tag-name-state)
            26. [[13.2.5.26]{.secno} Script data double escape start
                state](parsing.html#script-data-double-escape-start-state)
            27. [[13.2.5.27]{.secno} Script data double escaped
                state](parsing.html#script-data-double-escaped-state)
            28. [[13.2.5.28]{.secno} Script data double escaped dash
                state](parsing.html#script-data-double-escaped-dash-state)
            29. [[13.2.5.29]{.secno} Script data double escaped dash
                dash
                state](parsing.html#script-data-double-escaped-dash-dash-state)
            30. [[13.2.5.30]{.secno} Script data double escaped
                less-than sign
                state](parsing.html#script-data-double-escaped-less-than-sign-state)
            31. [[13.2.5.31]{.secno} Script data double escape end
                state](parsing.html#script-data-double-escape-end-state)
            32. [[13.2.5.32]{.secno} Before attribute name
                state](parsing.html#before-attribute-name-state)
            33. [[13.2.5.33]{.secno} Attribute name
                state](parsing.html#attribute-name-state)
            34. [[13.2.5.34]{.secno} After attribute name
                state](parsing.html#after-attribute-name-state)
            35. [[13.2.5.35]{.secno} Before attribute value
                state](parsing.html#before-attribute-value-state)
            36. [[13.2.5.36]{.secno} Attribute value (double-quoted)
                state](parsing.html#attribute-value-(double-quoted)-state)
            37. [[13.2.5.37]{.secno} Attribute value (single-quoted)
                state](parsing.html#attribute-value-(single-quoted)-state)
            38. [[13.2.5.38]{.secno} Attribute value (unquoted)
                state](parsing.html#attribute-value-(unquoted)-state)
            39. [[13.2.5.39]{.secno} After attribute value (quoted)
                state](parsing.html#after-attribute-value-(quoted)-state)
            40. [[13.2.5.40]{.secno} Self-closing start tag
                state](parsing.html#self-closing-start-tag-state)
            41. [[13.2.5.41]{.secno} Bogus comment
                state](parsing.html#bogus-comment-state)
            42. [[13.2.5.42]{.secno} Markup declaration open
                state](parsing.html#markup-declaration-open-state)
            43. [[13.2.5.43]{.secno} Comment start
                state](parsing.html#comment-start-state)
            44. [[13.2.5.44]{.secno} Comment start dash
                state](parsing.html#comment-start-dash-state)
            45. [[13.2.5.45]{.secno} Comment
                state](parsing.html#comment-state)
            46. [[13.2.5.46]{.secno} Comment less-than sign
                state](parsing.html#comment-less-than-sign-state)
            47. [[13.2.5.47]{.secno} Comment less-than sign bang
                state](parsing.html#comment-less-than-sign-bang-state)
            48. [[13.2.5.48]{.secno} Comment less-than sign bang dash
                state](parsing.html#comment-less-than-sign-bang-dash-state)
            49. [[13.2.5.49]{.secno} Comment less-than sign bang dash
                dash
                state](parsing.html#comment-less-than-sign-bang-dash-dash-state)
            50. [[13.2.5.50]{.secno} Comment end dash
                state](parsing.html#comment-end-dash-state)
            51. [[13.2.5.51]{.secno} Comment end
                state](parsing.html#comment-end-state)
            52. [[13.2.5.52]{.secno} Comment end bang
                state](parsing.html#comment-end-bang-state)
            53. [[13.2.5.53]{.secno} DOCTYPE
                state](parsing.html#doctype-state)
            54. [[13.2.5.54]{.secno} Before DOCTYPE name
                state](parsing.html#before-doctype-name-state)
            55. [[13.2.5.55]{.secno} DOCTYPE name
                state](parsing.html#doctype-name-state)
            56. [[13.2.5.56]{.secno} After DOCTYPE name
                state](parsing.html#after-doctype-name-state)
            57. [[13.2.5.57]{.secno} After DOCTYPE public keyword
                state](parsing.html#after-doctype-public-keyword-state)
            58. [[13.2.5.58]{.secno} Before DOCTYPE public identifier
                state](parsing.html#before-doctype-public-identifier-state)
            59. [[13.2.5.59]{.secno} DOCTYPE public identifier
                (double-quoted)
                state](parsing.html#doctype-public-identifier-(double-quoted)-state)
            60. [[13.2.5.60]{.secno} DOCTYPE public identifier
                (single-quoted)
                state](parsing.html#doctype-public-identifier-(single-quoted)-state)
            61. [[13.2.5.61]{.secno} After DOCTYPE public identifier
                state](parsing.html#after-doctype-public-identifier-state)
            62. [[13.2.5.62]{.secno} Between DOCTYPE public and system
                identifiers
                state](parsing.html#between-doctype-public-and-system-identifiers-state)
            63. [[13.2.5.63]{.secno} After DOCTYPE system keyword
                state](parsing.html#after-doctype-system-keyword-state)
            64. [[13.2.5.64]{.secno} Before DOCTYPE system identifier
                state](parsing.html#before-doctype-system-identifier-state)
            65. [[13.2.5.65]{.secno} DOCTYPE system identifier
                (double-quoted)
                state](parsing.html#doctype-system-identifier-(double-quoted)-state)
            66. [[13.2.5.66]{.secno} DOCTYPE system identifier
                (single-quoted)
                state](parsing.html#doctype-system-identifier-(single-quoted)-state)
            67. [[13.2.5.67]{.secno} After DOCTYPE system identifier
                state](parsing.html#after-doctype-system-identifier-state)
            68. [[13.2.5.68]{.secno} Bogus DOCTYPE
                state](parsing.html#bogus-doctype-state)
            69. [[13.2.5.69]{.secno} CDATA section
                state](parsing.html#cdata-section-state)
            70. [[13.2.5.70]{.secno} CDATA section bracket
                state](parsing.html#cdata-section-bracket-state)
            71. [[13.2.5.71]{.secno} CDATA section end
                state](parsing.html#cdata-section-end-state)
            72. [[13.2.5.72]{.secno} Character reference
                state](parsing.html#character-reference-state)
            73. [[13.2.5.73]{.secno} Named character reference
                state](parsing.html#named-character-reference-state)
            74. [[13.2.5.74]{.secno} Ambiguous ampersand
                state](parsing.html#ambiguous-ampersand-state)
            75. [[13.2.5.75]{.secno} Numeric character reference
                state](parsing.html#numeric-character-reference-state)
            76. [[13.2.5.76]{.secno} Hexadecimal character reference
                start
                state](parsing.html#hexadecimal-character-reference-start-state)
            77. [[13.2.5.77]{.secno} Decimal character reference start
                state](parsing.html#decimal-character-reference-start-state)
            78. [[13.2.5.78]{.secno} Hexadecimal character reference
                state](parsing.html#hexadecimal-character-reference-state)
            79. [[13.2.5.79]{.secno} Decimal character reference
                state](parsing.html#decimal-character-reference-state)
            80. [[13.2.5.80]{.secno} Numeric character reference end
                state](parsing.html#numeric-character-reference-end-state)
        6.  [[13.2.6]{.secno} Tree
            construction](parsing.html#tree-construction)
            1.  [[13.2.6.1]{.secno} Creating and inserting
                nodes](parsing.html#creating-and-inserting-nodes)
            2.  [[13.2.6.2]{.secno} Parsing elements that contain only
                text](parsing.html#parsing-elements-that-contain-only-text)
            3.  [[13.2.6.3]{.secno} Closing elements that have implied
                end
                tags](parsing.html#closing-elements-that-have-implied-end-tags)
            4.  [[13.2.6.4]{.secno} The rules for parsing tokens in HTML
                content](parsing.html#parsing-main-inhtml)
                1.  [[13.2.6.4.1]{.secno} The \"initial\" insertion
                    mode](parsing.html#the-initial-insertion-mode)
                2.  [[13.2.6.4.2]{.secno} The \"before html\" insertion
                    mode](parsing.html#the-before-html-insertion-mode)
                3.  [[13.2.6.4.3]{.secno} The \"before head\" insertion
                    mode](parsing.html#the-before-head-insertion-mode)
                4.  [[13.2.6.4.4]{.secno} The \"in head\" insertion
                    mode](parsing.html#parsing-main-inhead)
                5.  [[13.2.6.4.5]{.secno} The \"in head noscript\"
                    insertion
                    mode](parsing.html#parsing-main-inheadnoscript)
                6.  [[13.2.6.4.6]{.secno} The \"after head\" insertion
                    mode](parsing.html#the-after-head-insertion-mode)
                7.  [[13.2.6.4.7]{.secno} The \"in body\" insertion
                    mode](parsing.html#parsing-main-inbody)
                8.  [[13.2.6.4.8]{.secno} The \"text\" insertion
                    mode](parsing.html#parsing-main-incdata)
                9.  [[13.2.6.4.9]{.secno} The \"in table\" insertion
                    mode](parsing.html#parsing-main-intable)
                10. [[13.2.6.4.10]{.secno} The \"in table text\"
                    insertion
                    mode](parsing.html#parsing-main-intabletext)
                11. [[13.2.6.4.11]{.secno} The \"in caption\" insertion
                    mode](parsing.html#parsing-main-incaption)
                12. [[13.2.6.4.12]{.secno} The \"in column group\"
                    insertion
                    mode](parsing.html#parsing-main-incolgroup)
                13. [[13.2.6.4.13]{.secno} The \"in table body\"
                    insertion mode](parsing.html#parsing-main-intbody)
                14. [[13.2.6.4.14]{.secno} The \"in row\" insertion
                    mode](parsing.html#parsing-main-intr)
                15. [[13.2.6.4.15]{.secno} The \"in cell\" insertion
                    mode](parsing.html#parsing-main-intd)
                16. [[13.2.6.4.16]{.secno} The \"in select\" insertion
                    mode](parsing.html#parsing-main-inselect)
                17. [[13.2.6.4.17]{.secno} The \"in select in table\"
                    insertion
                    mode](parsing.html#parsing-main-inselectintable)
                18. [[13.2.6.4.18]{.secno} The \"in template\" insertion
                    mode](parsing.html#parsing-main-intemplate)
                19. [[13.2.6.4.19]{.secno} The \"after body\" insertion
                    mode](parsing.html#parsing-main-afterbody)
                20. [[13.2.6.4.20]{.secno} The \"in frameset\" insertion
                    mode](parsing.html#parsing-main-inframeset)
                21. [[13.2.6.4.21]{.secno} The \"after frameset\"
                    insertion
                    mode](parsing.html#parsing-main-afterframeset)
                22. [[13.2.6.4.22]{.secno} The \"after after body\"
                    insertion
                    mode](parsing.html#the-after-after-body-insertion-mode)
                23. [[13.2.6.4.23]{.secno} The \"after after frameset\"
                    insertion
                    mode](parsing.html#the-after-after-frameset-insertion-mode)
            5.  [[13.2.6.5]{.secno} The rules for parsing tokens in
                foreign content](parsing.html#parsing-main-inforeign)
        7.  [[13.2.7]{.secno} The end](parsing.html#the-end)
        8.  [[13.2.8]{.secno} Speculative HTML
            parsing](parsing.html#speculative-html-parsing)
        9.  [[13.2.9]{.secno} Coercing an HTML DOM into an
            infoset](parsing.html#coercing-an-html-dom-into-an-infoset)
        10. [[13.2.10]{.secno} An introduction to error handling and
            strange cases in the
            parser](parsing.html#an-introduction-to-error-handling-and-strange-cases-in-the-parser)
            1.  [[13.2.10.1]{.secno} Misnested tags:
                \<b\>\<i\>\</b\>\</i\>](parsing.html#misnested-tags:-b-i-/b-/i)
            2.  [[13.2.10.2]{.secno} Misnested tags:
                \<b\>\<p\>\</b\>\</p\>](parsing.html#misnested-tags:-b-p-/b-/p)
            3.  [[13.2.10.3]{.secno} Unexpected markup in
                tables](parsing.html#unexpected-markup-in-tables)
            4.  [[13.2.10.4]{.secno} Scripts that modify the page as it
                is being
                parsed](parsing.html#scripts-that-modify-the-page-as-it-is-being-parsed)
            5.  [[13.2.10.5]{.secno} The execution of scripts that are
                moving across multiple
                documents](parsing.html#the-execution-of-scripts-that-are-moving-across-multiple-documents)
            6.  [[13.2.10.6]{.secno} Unclosed formatting
                elements](parsing.html#unclosed-formatting-elements)
    2.  [[13.3]{.secno} Serializing HTML
        fragments](parsing.html#serialising-html-fragments)
    3.  [[13.4]{.secno} Parsing HTML
        fragments](parsing.html#parsing-html-fragments)
    :::

### [13.2]{.secno} Parsing HTML documents[](#parsing){.self-link} {#parsing}

*This section only applies to user agents, data mining tools, and
conformance checkers.*

The rules for parsing XML documents into DOM trees are covered by the
next section, entitled \"[The XML
syntax](xhtml.html#the-xhtml-syntax){#parsing:the-xhtml-syntax}\".

User agents must use the parsing rules described in this section to
generate the DOM trees from
[`text/html`{#parsing:text/html}](iana.html#text/html) resources.
Together, these rules define what is referred to as the [HTML
parser]{#html-parser .dfn export=""}.

::: note
While the HTML syntax described in this specification bears a close
resemblance to SGML and XML, it is a separate language with its own
parsing rules.

Some earlier versions of HTML (in particular from HTML2 to HTML4) were
based on SGML and used SGML parsing rules. However, few (if any) web
browsers ever implemented true SGML parsing for HTML documents; the only
user agents to strictly handle HTML as an SGML application have
historically been validators. The resulting confusion --- with
validators claiming documents to have one representation while widely
deployed web browsers interoperably implemented a different
representation --- has wasted decades of productivity. This version of
HTML thus returns to a non-SGML basis.
:::

For the purposes of conformance checkers, if a resource is determined to
be in [the HTML syntax](syntax.html#syntax){#parsing:syntax}, then it is
an [HTML
document](https://dom.spec.whatwg.org/#html-document){#parsing:html-documents
x-internal="html-documents"}.

As stated [in the terminology
section](infrastructure.html#html-elements){#parsing:html-elements},
references to [element
types](infrastructure.html#element-type){#parsing:element-type} that do
not explicitly specify a namespace always refer to elements in the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing:html-namespace-2
x-internal="html-namespace-2"}. For example, if the spec talks about \"a
[`menu`{#parsing:the-menu-element}](grouping-content.html#the-menu-element)
element\", then that is an element with the local name \"`menu`\", the
namespace \"`http://www.w3.org/1999/xhtml`\", and the interface
[`HTMLMenuElement`{#parsing:htmlmenuelement}](grouping-content.html#htmlmenuelement).
Where possible, references to such elements are hyperlinked to their
definition.

#### [13.2.1]{.secno} Overview of the parsing model[](#overview-of-the-parsing-model){.self-link}

![](/images/parsing-model-overview.svg){.darkmode-aware width="345"
height="535"}

The input to the HTML parsing process consists of a stream of [code
points](https://infra.spec.whatwg.org/#code-point){#overview-of-the-parsing-model:code-point
x-internal="code-point"}, which is passed through a
[tokenization](#tokenization){#overview-of-the-parsing-model:tokenization}
stage followed by a [tree
construction](#tree-construction){#overview-of-the-parsing-model:tree-construction}
stage. The output is a
[`Document`{#overview-of-the-parsing-model:document}](dom.html#document)
object.

Implementations that [do not support
scripting](infrastructure.html#non-scripted) do not have to actually
create a DOM
[`Document`{#overview-of-the-parsing-model:document-2}](dom.html#document)
object, but the DOM tree in such cases is still used as the model for
the rest of the specification.

In the common case, the data handled by the tokenization stage comes
from the network, but [it can also come from
script](dynamic-markup-insertion.html#dynamic-markup-insertion){#overview-of-the-parsing-model:dynamic-markup-insertion}
running in the user agent, e.g. using the
[`document.write()`{#overview-of-the-parsing-model:dom-document-write}](dynamic-markup-insertion.html#dom-document-write)
API.

There is only one set of states for the tokenizer stage and the tree
construction stage, but the tree construction stage is reentrant,
meaning that while the tree construction stage is handling one token,
the tokenizer might be resumed, causing further tokens to be emitted and
processed before the first token\'s processing is complete.

::: example
In the following example, the tree construction stage will be called
upon to handle a \"p\" start tag token while handling the \"script\" end
tag token:

``` html
...
<script>
 document.write('<p>');
</script>
...
```
:::

To handle these cases, parsers have a [script nesting
level]{#script-nesting-level .dfn}, which must be initially set to zero,
and a [parser pause flag]{#parser-pause-flag .dfn}, which must be
initially set to false.

#### [13.2.2]{.secno} [Parse errors]{.dfn}[](#parse-errors){.self-link}

This specification defines the parsing rules for HTML documents, whether
they are syntactically correct or not. Certain points in the parsing
algorithm are said to be [parse
errors](#parse-errors){#parse-errors:parse-errors}. The error handling
for parse errors is well-defined (that\'s the processing rules described
throughout this specification), but user agents, while parsing an HTML
document, may [abort the
parser](#abort-a-parser){#parse-errors:abort-a-parser} at the first
[parse error](#parse-errors){#parse-errors:parse-errors-2} that they
encounter for which they do not wish to apply the rules described in
this specification.

Conformance checkers must report at least one parse error condition to
the user if one or more parse error conditions exist in the document and
must not report parse error conditions if none exist in the document.
Conformance checkers may report more than one parse error condition if
more than one parse error condition exists in the document.

Parse errors are only errors with the *syntax* of HTML. In addition to
checking for parse errors, conformance checkers will also verify that
the document obeys all the other conformance requirements described in
this specification.

Some parse errors have dedicated codes outlined in the table below that
should be used by conformance checkers in reports.

*Error descriptions in the table below are non-normative.*

Code

Description

[abrupt-closing-of-empty-comment]{#parse-error-abrupt-closing-of-empty-comment
.dfn}

This error occurs if the parser encounters an empty
[comment](syntax.html#syntax-comments){#parse-errors:syntax-comments}
that is abruptly closed by a U+003E (\>) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point
x-internal="code-point"} (i.e., `<!-->` or `<!--->`). The parser behaves
as if the comment is closed correctly.

[abrupt-doctype-public-identifier]{#parse-error-abrupt-doctype-public-identifier
.dfn}

This error occurs if the parser encounters a U+003E (\>) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-2
x-internal="code-point"} in the
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype}
public identifier (e.g., `<!DOCTYPE html PUBLIC "foo>`). In such a case,
if the DOCTYPE is correctly placed as a document preamble, the parser
sets the [`Document`{#parse-errors:document}](dom.html#document) to
[quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode
x-internal="quirks-mode"}.

[abrupt-doctype-system-identifier]{#parse-error-abrupt-doctype-system-identifier
.dfn}

This error occurs if the parser encounters a U+003E (\>) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-3
x-internal="code-point"} in the
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-2}
system identifier (e.g.,
`<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01//EN" "foo>`). In such a
case, if the DOCTYPE is correctly placed as a document preamble, the
parser sets the
[`Document`{#parse-errors:document-2}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-2
x-internal="quirks-mode"}.

[absence-of-digits-in-numeric-character-reference]{#parse-error-absence-of-digits-in-numeric-character-reference
.dfn}

This error occurs if the parser encounters a numeric [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref}
that doesn\'t contain any digits (e.g., `&#qux;`). In this case the
parser doesn\'t resolve the character reference.

[cdata-in-html-content]{#parse-error-cdata-in-html-content .dfn}

This error occurs if the parser encounters a [CDATA
section](syntax.html#syntax-cdata){#parse-errors:syntax-cdata} outside
of foreign content (SVG or MathML). The parser treats such CDATA
sections (including leading \"`[CDATA[`\" and trailing \"`]]`\" strings)
as comments.

[character-reference-outside-unicode-range]{#parse-error-character-reference-outside-unicode-range
.dfn}

This error occurs if the parser encounters a numeric [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref-2}
that references a [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-4
x-internal="code-point"} that is greater than the valid Unicode range.
The parser resolves such a character reference to a U+FFFD REPLACEMENT
CHARACTER.

[control-character-in-input-stream]{#parse-error-control-character-in-input-stream
.dfn}

This error occurs if the [input
stream](#input-stream){#parse-errors:input-stream} contains a
[control](https://infra.spec.whatwg.org/#control){#parse-errors:control
x-internal="control"} [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-5
x-internal="code-point"} that is not [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters
x-internal="space-characters"} or U+0000 NULL. Such code points are
parsed as-is and usually, where parsing rules don\'t apply any
additional restrictions, make their way into the DOM.

[control-character-reference]{#parse-error-control-character-reference
.dfn}

This error occurs if the parser encounters a numeric [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref-3}
that references a
[control](https://infra.spec.whatwg.org/#control){#parse-errors:control-2
x-internal="control"} [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-6
x-internal="code-point"} that is not [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-2
x-internal="space-characters"} or is a U+000D CARRIAGE RETURN. The
parser resolves such character references as-is except C1 control
references that are replaced according to the [numeric character
reference end
state](#numeric-character-reference-end-state){#parse-errors:numeric-character-reference-end-state}.

[duplicate-attribute]{#parse-error-duplicate-attribute .dfn}

This error occurs if the parser encounters an
[attribute](syntax.html#syntax-attributes){#parse-errors:syntax-attributes}
in a tag that already has an attribute with the same name. The parser
ignores all such duplicate occurrences of the attribute.

[end-tag-with-attributes]{#parse-error-end-tag-with-attributes .dfn}

This error occurs if the parser encounters an [end
tag](syntax.html#syntax-end-tag){#parse-errors:syntax-end-tag} with
[attributes](syntax.html#syntax-attributes){#parse-errors:syntax-attributes-2}.
Attributes in end tags are ignored and do not make their way into the
DOM.

[end-tag-with-trailing-solidus]{#parse-error-end-tag-with-trailing-solidus
.dfn}

This error occurs if the parser encounters an [end
tag](syntax.html#syntax-end-tag){#parse-errors:syntax-end-tag-2} that
has a U+002F (/) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-7
x-internal="code-point"} right before the closing U+003E (\>) code point
(e.g., `</div/>`). Such a tag is treated as a regular end tag.

[eof-before-tag-name]{#parse-error-eof-before-tag-name .dfn}

This error occurs if the parser encounters the end of the [input
stream](#input-stream){#parse-errors:input-stream-2} where a tag name is
expected. In this case the parser treats the beginning of a [start
tag](syntax.html#syntax-start-tag){#parse-errors:syntax-start-tag}
(i.e., `<`) or an [end
tag](syntax.html#syntax-end-tag){#parse-errors:syntax-end-tag-3} (i.e.,
`</`) as text content.

[eof-in-cdata]{#parse-error-eof-in-cdata .dfn}

This error occurs if the parser encounters the end of the [input
stream](#input-stream){#parse-errors:input-stream-3} in a [CDATA
section](syntax.html#syntax-cdata){#parse-errors:syntax-cdata-2}. The
parser treats such CDATA sections as if they are closed immediately
before the end of the input stream.

[eof-in-comment]{#parse-error-eof-in-comment .dfn}

This error occurs if the parser encounters the end of the [input
stream](#input-stream){#parse-errors:input-stream-4} in a
[comment](syntax.html#syntax-comments){#parse-errors:syntax-comments-2}.
The parser treats such comments as if they are closed immediately before
the end of the input stream.

[eof-in-doctype]{#parse-error-eof-in-doctype .dfn}

This error occurs if the parser encounters the end of the input stream
in a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-3}.
In such a case, if the DOCTYPE is correctly placed as a document
preamble, the parser sets the
[`Document`{#parse-errors:document-3}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-3
x-internal="quirks-mode"}.

[eof-in-script-html-comment-like-text]{#parse-error-eof-in-script-html-comment-like-text
.dfn}

This error occurs if the parser encounters the end of the [input
stream](#input-stream){#parse-errors:input-stream-5} in text that
resembles an [HTML
comment](syntax.html#syntax-comments){#parse-errors:syntax-comments-3}
inside
[`script`{#parse-errors:the-script-element}](scripting.html#the-script-element)
element content (e.g., `<script><!-- foo`).

Syntactic structures that resemble HTML comments in
[`script`{#parse-errors:the-script-element-2}](scripting.html#the-script-element)
elements are parsed as text content. They can be a part of a scripting
language-specific syntactic structure or be treated as an HTML-like
comment, if the scripting language supports them (e.g., parsing rules
for HTML-like comments can be found in Annex B of the JavaScript
specification). The common reason for this error is a violation of the
[restrictions for contents of `script`
elements](scripting.html#restrictions-for-contents-of-script-elements){#parse-errors:restrictions-for-contents-of-script-elements}.
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

[eof-in-tag]{#parse-error-eof-in-tag .dfn}

This error occurs if the parser encounters the end of the [input
stream](#input-stream){#parse-errors:input-stream-6} in a [start
tag](syntax.html#syntax-start-tag){#parse-errors:syntax-start-tag-2} or
an [end tag](syntax.html#syntax-end-tag){#parse-errors:syntax-end-tag-4}
(e.g., `<div id=`). Such a tag is ignored.

[incorrectly-closed-comment]{#parse-error-incorrectly-closed-comment
.dfn}

This error occurs if the parser encounters a
[comment](syntax.html#syntax-comments){#parse-errors:syntax-comments-4}
that is closed by the \"`--!>`\" [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-8
x-internal="code-point"} sequence. The parser treats such comments as if
they are correctly closed by the \"`-->`\" code point sequence.

[incorrectly-opened-comment]{#parse-error-incorrectly-opened-comment
.dfn}

This error occurs if the parser encounters the \"`<!`\" [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-9
x-internal="code-point"} sequence that is not immediately followed by
two U+002D (-) code points and that is not the start of a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-4} or
a [CDATA
section](syntax.html#syntax-cdata){#parse-errors:syntax-cdata-3}. All
content that follows the \"`<!`\" code point sequence up to a U+003E
(\>) code point (if present) or to the end of the [input
stream](#input-stream){#parse-errors:input-stream-7} is treated as a
comment.

One possible cause of this error is using an XML markup declaration
(e.g., `<!ELEMENT br EMPTY>`) in HTML.

[invalid-character-sequence-after-doctype-name]{#parse-error-invalid-character-sequence-after-doctype-name
.dfn}

This error occurs if the parser encounters any [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-10
x-internal="code-point"} sequence other than \"`PUBLIC`\" and
\"`SYSTEM`\" keywords after a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-5}
name. In such a case, the parser ignores any following public or system
identifiers, and if the DOCTYPE is correctly placed as a document
preamble, and if the [parser cannot change the mode
flag](#parser-cannot-change-the-mode-flag){#parse-errors:parser-cannot-change-the-mode-flag}
is false, sets the
[`Document`{#parse-errors:document-4}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-4
x-internal="quirks-mode"}.

[invalid-first-character-of-tag-name]{#parse-error-invalid-first-character-of-tag-name
.dfn}

This error occurs if the parser encounters a [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-11
x-internal="code-point"} that is not an [ASCII
alpha](https://infra.spec.whatwg.org/#ascii-alpha){#parse-errors:ascii-letters
x-internal="ascii-letters"} where first code point of a [start
tag](syntax.html#syntax-start-tag){#parse-errors:syntax-start-tag-3}
name or an [end
tag](syntax.html#syntax-end-tag){#parse-errors:syntax-end-tag-5} name is
expected. If a start tag was expected such code point and a preceding
U+003C (\<) is treated as text content, and all content that follows is
treated as markup. Whereas, if an end tag was expected, such code point
and all content that follows up to a U+003E (\>) code point (if present)
or to the end of the [input
stream](#input-stream){#parse-errors:input-stream-8} is treated as a
comment.

::: example
For example, consider the following markup:

``` html
<42></42>
```

This will be parsed into:

- [`html`{#parse-errors:the-html-element}](semantics.html#the-html-element)
  - [`head`{#parse-errors:the-head-element}](semantics.html#the-head-element)
  - [`body`{#parse-errors:the-body-element}](sections.html#the-body-element)
    - [`#text`{#parse-errors:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
      \<42\>
    - [`#comment`{#parse-errors:comment-2}](https://dom.spec.whatwg.org/#interface-comment){x-internal="comment-2"}:
      42
:::

While the first code point of a tag name is limited to an [ASCII
alpha](https://infra.spec.whatwg.org/#ascii-alpha){#parse-errors:ascii-letters-2
x-internal="ascii-letters"}, a wide range of code points (including
[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#parse-errors:ascii-digits
x-internal="ascii-digits"}) is allowed in subsequent positions.

[missing-attribute-value]{#parse-error-missing-attribute-value .dfn}

This error occurs if the parser encounters a U+003E (\>) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-12
x-internal="code-point"} where an
[attribute](syntax.html#syntax-attributes){#parse-errors:syntax-attributes-3}
value is expected (e.g., `<div id=>`). The parser treats the attribute
as having an empty value.

[missing-doctype-name]{#parse-error-missing-doctype-name .dfn}

This error occurs if the parser encounters a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-6}
that is missing a name (e.g., `<!DOCTYPE>`). In such a case, if the
DOCTYPE is correctly placed as a document preamble, the parser sets the
[`Document`{#parse-errors:document-5}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-5
x-internal="quirks-mode"}.

[missing-doctype-public-identifier]{#parse-error-missing-doctype-public-identifier
.dfn}

This error occurs if the parser encounters a U+003E (\>) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-13
x-internal="code-point"} where start of the
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-7}
public identifier is expected (e.g., `<!DOCTYPE html PUBLIC >`). In such
a case, if the DOCTYPE is correctly placed as a document preamble, the
parser sets the
[`Document`{#parse-errors:document-6}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-6
x-internal="quirks-mode"}.

[missing-doctype-system-identifier]{#parse-error-missing-doctype-system-identifier
.dfn}

This error occurs if the parser encounters a U+003E (\>) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-14
x-internal="code-point"} where start of the
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-8}
system identifier is expected (e.g., `<!DOCTYPE html SYSTEM >`). In such
a case, if the DOCTYPE is correctly placed as a document preamble, the
parser sets the
[`Document`{#parse-errors:document-7}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-7
x-internal="quirks-mode"}.

[missing-end-tag-name]{#parse-error-missing-end-tag-name .dfn}

This error occurs if the parser encounters a U+003E (\>) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-15
x-internal="code-point"} where an [end
tag](syntax.html#syntax-end-tag){#parse-errors:syntax-end-tag-6} name is
expected, i.e., `</>`. The parser ignores the whole \"`</>`\" code point
sequence.

[missing-quote-before-doctype-public-identifier]{#parse-error-missing-quote-before-doctype-public-identifier
.dfn}

This error occurs if the parser encounters the
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-9}
public identifier that is not preceded by a quote (e.g.,
`<!DOCTYPE html PUBLIC -//W3C//DTD HTML 4.01//EN">`). In such a case,
the parser ignores the public identifier, and if the DOCTYPE is
correctly placed as a document preamble, sets the
[`Document`{#parse-errors:document-8}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-8
x-internal="quirks-mode"}.

[missing-quote-before-doctype-system-identifier]{#parse-error-missing-quote-before-doctype-system-identifier
.dfn}

This error occurs if the parser encounters the
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-10}
system identifier that is not preceded by a quote (e.g.,
`<!DOCTYPE html SYSTEM http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`).
In such a case, the parser ignores the system identifier, and if the
DOCTYPE is correctly placed as a document preamble, sets the
[`Document`{#parse-errors:document-9}](dom.html#document) to [quirks
mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parse-errors:quirks-mode-9
x-internal="quirks-mode"}.

[missing-semicolon-after-character-reference]{#parse-error-missing-semicolon-after-character-reference
.dfn}

This error occurs if the parser encounters a [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref-4}
that is not terminated by a U+003B (;) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-16
x-internal="code-point"}. Usually the parser behaves as if character
reference is terminated by the U+003B (;) code point; however, there are
some ambiguous cases in which the parser includes subsequent code points
in the character reference.

For example, `&not;in` will be parsed as \"`¬in`\" whereas `&notin` will
be parsed as \"`∉`\".

[missing-whitespace-after-doctype-public-keyword]{#parse-error-missing-whitespace-after-doctype-public-keyword
.dfn}

This error occurs if the parser encounters a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-11}
whose \"`PUBLIC`\" keyword and public identifier are not separated by
[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-3
x-internal="space-characters"}. In this case the parser behaves as if
ASCII whitespace is present.

[missing-whitespace-after-doctype-system-keyword]{#parse-error-missing-whitespace-after-doctype-system-keyword
.dfn}

This error occurs if the parser encounters a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-12}
whose \"`SYSTEM`\" keyword and system identifier are not separated by
[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-4
x-internal="space-characters"}. In this case the parser behaves as if
ASCII whitespace is present.

[missing-whitespace-before-doctype-name]{#parse-error-missing-whitespace-before-doctype-name
.dfn}

This error occurs if the parser encounters a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-13}
whose \"`DOCTYPE`\" keyword and name are not separated by [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-5
x-internal="space-characters"}. In this case the parser behaves as if
ASCII whitespace is present.

[missing-whitespace-between-attributes]{#parse-error-missing-whitespace-between-attributes
.dfn}

This error occurs if the parser encounters
[attributes](syntax.html#syntax-attributes){#parse-errors:syntax-attributes-4}
that are not separated by [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-6
x-internal="space-characters"} (e.g., `<div id="foo"class="bar">`). In
this case the parser behaves as if ASCII whitespace is present.

[missing-whitespace-between-doctype-public-and-system-identifiers]{#parse-error-missing-whitespace-between-doctype-public-and-system-identifiers
.dfn}

This error occurs if the parser encounters a
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-14}
whose public and system identifiers are not separated by [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-7
x-internal="space-characters"}. In this case the parser behaves as if
ASCII whitespace is present.

[nested-comment]{#parse-error-nested-comment .dfn}

This error occurs if the parser encounters a nested
[comment](syntax.html#syntax-comments){#parse-errors:syntax-comments-5}
(e.g., `<!-- <!-- nested --> -->`). Such a comment will be closed by the
first occurring \"`-->`\" [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-17
x-internal="code-point"} sequence and everything that follows will be
treated as markup.

[noncharacter-character-reference]{#parse-error-noncharacter-character-reference
.dfn}

This error occurs if the parser encounters a numeric [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref-5}
that references a
[noncharacter](https://infra.spec.whatwg.org/#noncharacter){#parse-errors:noncharacter
x-internal="noncharacter"}. The parser resolves such character
references as-is.

[noncharacter-in-input-stream]{#parse-error-noncharacter-in-input-stream
.dfn}

This error occurs if the [input
stream](#input-stream){#parse-errors:input-stream-9} contains a
[noncharacter](https://infra.spec.whatwg.org/#noncharacter){#parse-errors:noncharacter-2
x-internal="noncharacter"}. Such [code
points](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-18
x-internal="code-point"} are parsed as-is and usually, where parsing
rules don\'t apply any additional restrictions, make their way into the
DOM.

[non-void-html-element-start-tag-with-trailing-solidus]{#parse-error-non-void-html-element-start-tag-with-trailing-solidus
.dfn}

This error occurs if the parser encounters a [start
tag](syntax.html#syntax-start-tag){#parse-errors:syntax-start-tag-4} for
an element that is not in the list of [void
elements](syntax.html#void-elements){#parse-errors:void-elements} or is
not a part of foreign content (i.e., not an SVG or MathML element) that
has a U+002F (/) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-19
x-internal="code-point"} right before the closing U+003E (\>) code
point. The parser behaves as if the U+002F (/) is not present.

::: example
For example, consider the following markup:

``` html
<div/><span></span><span></span>
```

This will be parsed into:

- [`html`{#parse-errors:the-html-element-2}](semantics.html#the-html-element)
  - [`head`{#parse-errors:the-head-element-2}](semantics.html#the-head-element)
  - [`body`{#parse-errors:the-body-element-2}](sections.html#the-body-element)
    - [`div`{#parse-errors:the-div-element}](grouping-content.html#the-div-element)
      - [`span`{#parse-errors:the-span-element}](text-level-semantics.html#the-span-element)
      - [`span`{#parse-errors:the-span-element-2}](text-level-semantics.html#the-span-element)
:::

The trailing U+002F (/) in a start tag name can be used only in foreign
content to specify self-closing tags. (Self-closing tags don\'t exist in
HTML.) It is also allowed for void elements, but doesn\'t have any
effect in this case.

[null-character-reference]{#parse-error-null-character-reference .dfn}

This error occurs if the parser encounters a numeric [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref-6}
that references a U+0000 NULL [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-20
x-internal="code-point"}. The parser resolves such character references
to a U+FFFD REPLACEMENT CHARACTER.

[surrogate-character-reference]{#parse-error-surrogate-character-reference
.dfn}

This error occurs if the parser encounters a numeric [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref-7}
that references a
[surrogate](https://infra.spec.whatwg.org/#surrogate){#parse-errors:surrogate
x-internal="surrogate"}. The parser resolves such character references
to a U+FFFD REPLACEMENT CHARACTER.

[surrogate-in-input-stream]{#parse-error-surrogate-in-input-stream .dfn}

This error occurs if the [input
stream](#input-stream){#parse-errors:input-stream-10} contains a
[surrogate](https://infra.spec.whatwg.org/#surrogate){#parse-errors:surrogate-2
x-internal="surrogate"}. Such [code
points](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-21
x-internal="code-point"} are parsed as-is and usually, where parsing
rules don\'t apply any additional restrictions, make their way into the
DOM.

Surrogates can only find their way into the input stream via script APIs
such as
[`document.write()`{#parse-errors:dom-document-write}](dynamic-markup-insertion.html#dom-document-write).

[unexpected-character-after-doctype-system-identifier]{#parse-error-unexpected-character-after-doctype-system-identifier
.dfn}

This error occurs if the parser encounters any [code
points](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-22
x-internal="code-point"} other than [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-8
x-internal="space-characters"} or closing U+003E (\>) after the
[DOCTYPE](syntax.html#syntax-doctype){#parse-errors:syntax-doctype-15}
system identifier. The parser ignores these code points.

[unexpected-character-in-attribute-name]{#parse-error-unexpected-character-in-attribute-name
.dfn}

This error occurs if the parser encounters a U+0022 (\"), U+0027 (\'),
or U+003C (\<) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-23
x-internal="code-point"} in an [attribute
name](syntax.html#syntax-attribute-name){#parse-errors:syntax-attribute-name}.
The parser includes such code points in the attribute name.

Code points that trigger this error are usually a part of another
syntactic construct and can be a sign of a typo around the attribute
name.

::: example
For example, consider the following markup:

``` html
<div foo<div>
```

Due to a forgotten U+003E (\>) code point after `foo` the parser treats
this markup as a single
[`div`{#parse-errors:the-div-element-2}](grouping-content.html#the-div-element)
element with a \"`foo<div`\" attribute.

As another example of this error, consider the following markup:

``` html
<div id'bar'>
```

Due to a forgotten U+003D (=) code point between an attribute name and
value the parser treats this markup as a
[`div`{#parse-errors:the-div-element-3}](grouping-content.html#the-div-element)
element with the attribute \"`id'bar'`\" that has an empty value.
:::

[unexpected-character-in-unquoted-attribute-value]{#parse-error-unexpected-character-in-unquoted-attribute-value
.dfn}

This error occurs if the parser encounters a U+0022 (\"), U+0027 (\'),
U+003C (\<), U+003D (=), or U+0060 (\`) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-24
x-internal="code-point"} in an unquoted [attribute
value](syntax.html#syntax-attribute-value){#parse-errors:syntax-attribute-value}.
The parser includes such code points in the attribute value.

Code points that trigger this error are usually a part of another
syntactic construct and can be a sign of a typo around the attribute
value.

U+0060 (\`) is in the list of code points that trigger this error
because certain legacy user agents treat it as a quote.

::: example
For example, consider the following markup:

``` html
<div foo=b'ar'>
```

Due to a misplaced U+0027 (\') code point the parser sets the value of
the \"`foo`\" attribute to \"`b'ar'`\".
:::

[unexpected-equals-sign-before-attribute-name]{#parse-error-unexpected-equals-sign-before-attribute-name
.dfn}

This error occurs if the parser encounters a U+003D (=) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-25
x-internal="code-point"} before an attribute name. In this case the
parser treats U+003D (=) as the first code point of the attribute name.

The common reason for this error is a forgotten attribute name.

::: example
For example, consider the following markup:

``` html
<div foo="bar" ="baz">
```

Due to a forgotten attribute name the parser treats this markup as a
[`div`{#parse-errors:the-div-element-4}](grouping-content.html#the-div-element)
element with two attributes: a \"`foo`\" attribute with a \"`bar`\"
value and a \"`="baz"`\" attribute with an empty value.
:::

[unexpected-null-character]{#parse-error-unexpected-null-character .dfn}

This error occurs if the parser encounters a U+0000 NULL [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-26
x-internal="code-point"} in the [input
stream](#input-stream){#parse-errors:input-stream-11} in certain
positions. In general, such code points are either ignored or, for
security reasons, replaced with a U+FFFD REPLACEMENT CHARACTER.

[unexpected-question-mark-instead-of-tag-name]{#parse-error-unexpected-question-mark-instead-of-tag-name
.dfn}

This error occurs if the parser encounters a U+003F (?) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-27
x-internal="code-point"} where first code point of a [start
tag](syntax.html#syntax-start-tag){#parse-errors:syntax-start-tag-5}
name is expected. The U+003F (?) and all content that follows up to a
U+003E (\>) code point (if present) or to the end of the [input
stream](#input-stream){#parse-errors:input-stream-12} is treated as a
comment.

::: example
For example, consider the following markup:

``` html
<?xml-stylesheet type="text/css" href="style.css"?>
```

This will be parsed into:

- [`#comment`{#parse-errors:comment-2-2}](https://dom.spec.whatwg.org/#interface-comment){x-internal="comment-2"}:
  ?xml-stylesheet type=\"text/css\" href=\"style.css\"?
- [`html`{#parse-errors:the-html-element-3}](semantics.html#the-html-element)
  - [`head`{#parse-errors:the-head-element-3}](semantics.html#the-head-element)
  - [`body`{#parse-errors:the-body-element-3}](sections.html#the-body-element)
:::

The common reason for this error is an XML processing instruction (e.g.,
`<?xml-stylesheet type="text/css" href="style.css"?>`) or an XML
declaration (e.g., `<?xml version="1.0" encoding="UTF-8"?>`) being used
in HTML.

[unexpected-solidus-in-tag]{#parse-error-unexpected-solidus-in-tag .dfn}

This error occurs if the parser encounters a U+002F (/) [code
point](https://infra.spec.whatwg.org/#code-point){#parse-errors:code-point-28
x-internal="code-point"} that is not a part of a quoted
[attribute](syntax.html#syntax-attributes){#parse-errors:syntax-attributes-5}
value and not immediately followed by a U+003E (\>) code point in a tag
(e.g., `<div / id="foo">`). In this case the parser behaves as if it
encountered [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parse-errors:space-characters-9
x-internal="space-characters"}.

[unknown-named-character-reference]{#parse-error-unknown-named-character-reference
.dfn}

This error occurs if the parser encounters an [ambiguous
ampersand](syntax.html#syntax-ambiguous-ampersand){#parse-errors:syntax-ambiguous-ampersand}.
In this case the parser doesn\'t resolve the [character
reference](syntax.html#syntax-charref){#parse-errors:syntax-charref-8}.

#### [13.2.3]{.secno} The [input byte stream]{.dfn}[](#the-input-byte-stream){.self-link}

The stream of code points that comprises the input to the tokenization
stage will be initially seen by the user agent as a stream of bytes
(typically coming over the network or from the local file system). The
bytes encode the actual characters according to a particular *character
encoding*, which the user agent uses to decode the bytes into
characters.

For XML documents, the algorithm user agents are required to use to
determine the character encoding is given by XML. This section does not
apply to XML documents. [\[XML\]](references.html#refsXML)

Usually, the [encoding sniffing
algorithm](#encoding-sniffing-algorithm){#the-input-byte-stream:encoding-sniffing-algorithm}
defined below is used to determine the character encoding.

Given a character encoding, the bytes in the [input byte
stream](#the-input-byte-stream){#the-input-byte-stream:the-input-byte-stream}
must be converted to characters for the tokenizer\'s [input
stream](#input-stream){#the-input-byte-stream:input-stream}, by passing
the [input byte
stream](#the-input-byte-stream){#the-input-byte-stream:the-input-byte-stream-2}
and character encoding to
[decode](https://encoding.spec.whatwg.org/#decode){#the-input-byte-stream:decode
x-internal="decode"}.

A leading Byte Order Mark (BOM) causes the character encoding argument
to be ignored and will itself be skipped.

Bytes or sequences of bytes in the original byte stream that did not
conform to the Encoding standard (e.g. invalid UTF-8 byte sequences in a
UTF-8 input byte stream) are errors that conformance checkers are
expected to report. [\[ENCODING\]](references.html#refsENCODING)

The decoder algorithms describe how to handle invalid input; for
security reasons, it is imperative that those rules be followed
precisely. Differences in how invalid byte sequences are handled can
result in, amongst other problems, script injection vulnerabilities
(\"XSS\").

When the HTML parser is decoding an input byte stream, it uses a
character encoding and a [confidence]{#concept-encoding-confidence
.dfn}. The confidence is either *tentative*, *certain*, or *irrelevant*.
The encoding used, and whether the confidence in that encoding is
*tentative* or *certain*, is [used during the
parsing](#meta-charset-during-parse) to determine whether to [change the
encoding](#change-the-encoding){#the-input-byte-stream:change-the-encoding}.
If no encoding is necessary, e.g. because the parser is operating on a
Unicode stream and doesn\'t have to use a character encoding at all,
then the
[confidence](#concept-encoding-confidence){#the-input-byte-stream:concept-encoding-confidence}
is *irrelevant*.

Some algorithms feed the parser by directly adding characters to the
[input stream](#input-stream){#the-input-byte-stream:input-stream-2}
rather than adding bytes to the [input byte
stream](#the-input-byte-stream){#the-input-byte-stream:the-input-byte-stream-3}.

##### [13.2.3.1]{.secno} Parsing with a known character encoding[](#parsing-with-a-known-character-encoding){.self-link}

When the HTML parser is to operate on an input byte stream that has [a
known definite encoding]{#a-known-definite-encoding .dfn export=""},
then the character encoding is that encoding and the
[confidence](#concept-encoding-confidence){#parsing-with-a-known-character-encoding:concept-encoding-confidence}
is *certain*.

##### [13.2.3.2]{.secno} Determining the character encoding[](#determining-the-character-encoding){.self-link}

In some cases, it might be impractical to unambiguously determine the
encoding before parsing the document. Because of this, this
specification provides for a two-pass mechanism with an optional
pre-scan. Implementations are allowed, as described below, to apply a
simplified parsing algorithm to whatever bytes they have available
before beginning to parse the document. Then, the real parser is
started, using a tentative encoding derived from this pre-parse and
other out-of-band metadata. If, while the document is being loaded, the
user agent discovers a character encoding declaration that conflicts
with this information, then the parser can get reinvoked to perform a
parse of the document with the real encoding.

User agents must use the following algorithm, called the [encoding
sniffing algorithm]{#encoding-sniffing-algorithm .dfn}, to determine the
character encoding to use when decoding a document in the first pass.
This algorithm takes as input any out-of-band metadata available to the
user agent (e.g. the [Content-Type
metadata](urls-and-fetching.html#content-type){#determining-the-character-encoding:content-type}
of the document) and all the bytes available so far, and returns a
character encoding and a
[confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence}
that is either *tentative* or *certain*.

1.  If the result of [BOM
    sniffing](https://encoding.spec.whatwg.org/#bom-sniff){#determining-the-character-encoding:bom-sniff
    x-internal="bom-sniff"} is an encoding, return that encoding with
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-2}
    *certain*.

    Although the
    [decode](https://encoding.spec.whatwg.org/#decode){#determining-the-character-encoding:decode
    x-internal="decode"} algorithm will itself change the encoding to
    use based on the presence of a byte order mark, this algorithm
    sniffs the BOM as well in order to set the correct [document\'s
    character
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#determining-the-character-encoding:document's-character-encoding
    x-internal="document's-character-encoding"} and
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-3}.

2.  If the user has explicitly instructed the user agent to override the
    document\'s character encoding with a specific encoding, optionally
    return that encoding with the
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-4}
    *certain*.

    Typically, user agents remember such user requests across sessions,
    and in some cases apply them to documents in
    [`iframe`{#determining-the-character-encoding:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s
    as well.

3.  The user agent may wait for more bytes of the resource to be
    available, either in this step or at any later step in this
    algorithm. For instance, a user agent might wait 500ms or 1024
    bytes, whichever came first. In general preparsing the source to
    find the encoding improves performance, as it reduces the need to
    throw away the data structures used when parsing upon finding the
    encoding information. However, if the user agent delays too long to
    obtain data to determine the encoding, then the cost of the delay
    could outweigh any performance improvements from the preparse.

    The authoring conformance requirements for character encoding
    declarations limit them to only appearing [in the first 1024
    bytes](semantics.html#charset1024). User agents are therefore
    encouraged to use the prescan algorithm below (as invoked by these
    steps) on the first 1024 bytes, but not to stall beyond that.

4.  If the transport layer specifies a character encoding, and it is
    supported, return that encoding with the
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-5}
    *certain*.

5.  Optionally, [prescan the byte stream to determine its
    encoding](#prescan-a-byte-stream-to-determine-its-encoding){#determining-the-character-encoding:prescan-a-byte-stream-to-determine-its-encoding},
    with the *[end
    condition](#prescan-end-condition){#determining-the-character-encoding:prescan-end-condition}*
    being when the user agent decides that scanning further bytes would
    not be efficient. User agents are encouraged to only prescan the
    first 1024 bytes. User agents may decide that scanning *any* bytes
    is not efficient, in which case these substeps are entirely skipped.

    The aforementioned algorithm returns either a character encoding or
    failure. If it returns a character encoding, then return the same
    encoding, with
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-6}
    *tentative*.

6.  If the [HTML
    parser](#html-parser){#determining-the-character-encoding:html-parser}
    for which this algorithm is being run is associated with a
    [`Document`{#determining-the-character-encoding:document}](dom.html#document)
    `d`{.variable} whose [container
    document](document-sequences.html#doc-container-document){#determining-the-character-encoding:doc-container-document}
    is non-null, then:

    1.  Let `parentDocument`{.variable} be `d`{.variable}\'s [container
        document](document-sequences.html#doc-container-document){#determining-the-character-encoding:doc-container-document-2}.

    2.  If `parentDocument`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#determining-the-character-encoding:concept-document-origin
        x-internal="concept-document-origin"} is [same
        origin](browsers.html#same-origin){#determining-the-character-encoding:same-origin}
        with `d`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#determining-the-character-encoding:concept-document-origin-2
        x-internal="concept-document-origin"} and
        `parentDocument`{.variable}\'s [character
        encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#determining-the-character-encoding:document's-character-encoding-2
        x-internal="document's-character-encoding"} is not
        [UTF-16BE/LE](https://encoding.spec.whatwg.org/#utf-16be-le){#determining-the-character-encoding:utf-16-encoding
        x-internal="utf-16-encoding"}, then return
        `parentDocument`{.variable}\'s [character
        encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#determining-the-character-encoding:document's-character-encoding-3
        x-internal="document's-character-encoding"}, with the
        [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-7}
        *tentative*.

7.  Otherwise, if the user agent has information on the likely encoding
    for this page, e.g. based on the encoding of the page when it was
    last visited, then return that encoding, with the
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-8}
    *tentative*.

8.  The user agent may attempt to autodetect the character encoding from
    applying frequency analysis or other algorithms to the data stream.
    Such algorithms may use information about the resource other than
    the resource\'s contents, including the address of the resource. If
    autodetection succeeds in determining a character encoding, and that
    encoding is a supported encoding, then return that encoding, with
    the
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-9}
    *tentative*. [\[UNIVCHARDET\]](references.html#refsUNIVCHARDET)

    User agents are generally discouraged from attempting to autodetect
    encodings for resources obtained over the network, since doing so
    involves inherently non-interoperable heuristics. Attempting to
    detect encodings based on an HTML document\'s preamble is especially
    tricky since HTML markup typically uses only ASCII characters, and
    HTML documents tend to begin with a lot of markup rather than with
    text content.

    The UTF-8 encoding has a highly detectable bit pattern. Files from
    the local file system that contain bytes with values greater than
    0x7F which match the UTF-8 pattern are very likely to be UTF-8,
    while documents with byte sequences that do not match it are very
    likely not. When a user agent can examine the whole file, rather
    than just the preamble, detecting for UTF-8 specifically can be
    especially effective. [\[PPUTF8\]](references.html#refsPPUTF8)
    [\[UTF8DET\]](references.html#refsUTF8DET)

9.  Otherwise, return an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#determining-the-character-encoding:implementation-defined
    x-internal="implementation-defined"} or user-specified default
    character encoding, with the
    [confidence](#concept-encoding-confidence){#determining-the-character-encoding:concept-encoding-confidence-10}
    *tentative*.

    In controlled environments or in environments where the encoding of
    documents can be prescribed (for example, for user agents intended
    for dedicated use in new networks), the comprehensive `UTF-8`
    encoding is suggested.

    In other environments, the default encoding is typically dependent
    on the user\'s locale (an approximation of the languages, and thus
    often encodings, of the pages that the user is likely to frequent).
    The following table gives suggested defaults based on the user\'s
    locale, for compatibility with legacy content. Locales are
    identified by BCP 47 language tags.
    [\[BCP47\]](references.html#refsBCP47)
    [\[ENCODING\]](references.html#refsENCODING)

    Locale language

    Suggested default encoding

    ar

    Arabic

    [windows-1256](https://encoding.spec.whatwg.org/#windows-1256){#determining-the-character-encoding:windows-1256
    x-internal="windows-1256"}

    az

    Azeri

    [windows-1254](https://encoding.spec.whatwg.org/#windows-1254){#determining-the-character-encoding:windows-1254
    x-internal="windows-1254"}

    ba

    Bashkir

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251
    x-internal="windows-1251"}

    be

    Belarusian

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-2
    x-internal="windows-1251"}

    bg

    Bulgarian

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-3
    x-internal="windows-1251"}

    cs

    Czech

    [windows-1250](https://encoding.spec.whatwg.org/#windows-1250){#determining-the-character-encoding:windows-1250
    x-internal="windows-1250"}

    el

    Greek

    [ISO-8859-7](https://encoding.spec.whatwg.org/#iso-8859-7){#determining-the-character-encoding:iso-8859-7
    x-internal="iso-8859-7"}

    et

    Estonian

    [windows-1257](https://encoding.spec.whatwg.org/#windows-1257){#determining-the-character-encoding:windows-1257
    x-internal="windows-1257"}

    fa

    Persian

    [windows-1256](https://encoding.spec.whatwg.org/#windows-1256){#determining-the-character-encoding:windows-1256-2
    x-internal="windows-1256"}

    he

    Hebrew

    [windows-1255](https://encoding.spec.whatwg.org/#windows-1255){#determining-the-character-encoding:windows-1255
    x-internal="windows-1255"}

    hr

    Croatian

    [windows-1250](https://encoding.spec.whatwg.org/#windows-1250){#determining-the-character-encoding:windows-1250-2
    x-internal="windows-1250"}

    hu

    Hungarian

    [ISO-8859-2](https://encoding.spec.whatwg.org/#iso-8859-2){#determining-the-character-encoding:iso-8859-2
    x-internal="iso-8859-2"}

    ja

    Japanese

    [Shift_JIS](https://encoding.spec.whatwg.org/#shift_jis){#determining-the-character-encoding:shift_jis
    x-internal="shift_jis"}

    kk

    Kazakh

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-4
    x-internal="windows-1251"}

    ko

    Korean

    [EUC-KR](https://encoding.spec.whatwg.org/#euc-kr){#determining-the-character-encoding:euc-kr
    x-internal="euc-kr"}

    ku

    Kurdish

    [windows-1254](https://encoding.spec.whatwg.org/#windows-1254){#determining-the-character-encoding:windows-1254-2
    x-internal="windows-1254"}

    ky

    Kyrgyz

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-5
    x-internal="windows-1251"}

    lt

    Lithuanian

    [windows-1257](https://encoding.spec.whatwg.org/#windows-1257){#determining-the-character-encoding:windows-1257-2
    x-internal="windows-1257"}

    lv

    Latvian

    [windows-1257](https://encoding.spec.whatwg.org/#windows-1257){#determining-the-character-encoding:windows-1257-3
    x-internal="windows-1257"}

    mk

    Macedonian

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-6
    x-internal="windows-1251"}

    pl

    Polish

    [ISO-8859-2](https://encoding.spec.whatwg.org/#iso-8859-2){#determining-the-character-encoding:iso-8859-2-2
    x-internal="iso-8859-2"}

    ru

    Russian

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-7
    x-internal="windows-1251"}

    sah

    Yakut

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-8
    x-internal="windows-1251"}

    sk

    Slovak

    [windows-1250](https://encoding.spec.whatwg.org/#windows-1250){#determining-the-character-encoding:windows-1250-3
    x-internal="windows-1250"}

    sl

    Slovenian

    [ISO-8859-2](https://encoding.spec.whatwg.org/#iso-8859-2){#determining-the-character-encoding:iso-8859-2-3
    x-internal="iso-8859-2"}

    sr

    Serbian

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-9
    x-internal="windows-1251"}

    tg

    Tajik

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-10
    x-internal="windows-1251"}

    th

    Thai

    [windows-874](https://encoding.spec.whatwg.org/#windows-874){#determining-the-character-encoding:windows-874
    x-internal="windows-874"}

    tr

    Turkish

    [windows-1254](https://encoding.spec.whatwg.org/#windows-1254){#determining-the-character-encoding:windows-1254-3
    x-internal="windows-1254"}

    tt

    Tatar

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-11
    x-internal="windows-1251"}

    uk

    Ukrainian

    [windows-1251](https://encoding.spec.whatwg.org/#windows-1251){#determining-the-character-encoding:windows-1251-12
    x-internal="windows-1251"}

    vi

    Vietnamese

    [windows-1258](https://encoding.spec.whatwg.org/#windows-1258){#determining-the-character-encoding:windows-1258
    x-internal="windows-1258"}

    zh-Hans, zh-CN, zh-SG

    Chinese, Simplified

    [GBK](https://encoding.spec.whatwg.org/#gbk){#determining-the-character-encoding:gbk
    x-internal="gbk"}

    zh-Hant, zh-HK, zh-MO, zh-TW

    Chinese, Traditional

    [Big5](https://encoding.spec.whatwg.org/#big5){#determining-the-character-encoding:big5
    x-internal="big5"}

    All other locales

    [windows-1252](https://encoding.spec.whatwg.org/#windows-1252){#determining-the-character-encoding:windows-1252
    x-internal="windows-1252"}

    [The contents of this table are derived from the intersection of
    Windows, Chrome, and Firefox defaults.]{.small}

The [document\'s character
encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#determining-the-character-encoding:document's-character-encoding-4
x-internal="document's-character-encoding"} must immediately be set to
the value returned from this algorithm, at the same time as the user
agent uses the returned value to select the decoder to use for the input
byte stream.

------------------------------------------------------------------------

When an algorithm requires a user agent to [prescan a byte stream to
determine its encoding]{#prescan-a-byte-stream-to-determine-its-encoding
.dfn export=""}, given some defined
[`end condition`{.variable}]{#prescan-end-condition .dfn
dfn-for="prescan a byte
  stream to determine its encoding" export=""}, then it must run the
following steps. If at any point during these steps (including during
instances of the [get an
attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing}
algorithm invoked by this one) the user agent either runs out of bytes
(meaning the `position`{.variable} pointer created in the first step
below goes beyond the end of the byte stream obtained so far) or reaches
its `end condition`{.variable}, then abort the [prescan a byte stream to
determine its
encoding](#prescan-a-byte-stream-to-determine-its-encoding){#determining-the-character-encoding:prescan-a-byte-stream-to-determine-its-encoding-2}
algorithm and return the result [get an XML
encoding](#concept-get-xml-encoding-when-sniffing){#determining-the-character-encoding:concept-get-xml-encoding-when-sniffing}
applied to the same bytes that the [prescan a byte stream to determine
its
encoding](#prescan-a-byte-stream-to-determine-its-encoding){#determining-the-character-encoding:prescan-a-byte-stream-to-determine-its-encoding-3}
algorithm was applied to. Otherwise, these steps will return a character
encoding.

1.  Let `position`{.variable} be a pointer to a byte in the input byte
    stream, initially pointing at the first byte.

2.  Prescan for UTF-16 XML declarations: If `position`{.variable} points
    to:

    A sequence of bytes starting with: 0x3C, 0x0, 0x3F, 0x0, 0x78, 0x0 (case-sensitive UTF-16 little-endian \'\<?x\')
    :   Return
        [UTF-16LE](https://encoding.spec.whatwg.org/#utf-16le){#determining-the-character-encoding:utf-16le
        x-internal="utf-16le"}.

    A sequence of bytes starting with: 0x0, 0x3C, 0x0, 0x3F, 0x0, 0x78 (case-sensitive UTF-16 big-endian \'\<?x\')

    :   Return
        [UTF-16BE](https://encoding.spec.whatwg.org/#utf-16be){#determining-the-character-encoding:utf-16be
        x-internal="utf-16be"}.

    For historical reasons, the prefix is two bytes longer than in
    [Appendix F](https://www.w3.org/TR/REC-xml/#sec-guessing) of XML and
    the encoding name is not checked.

3.  *Loop*: If `position`{.variable} points to:

    A sequence of bytes starting with: 0x3C 0x21 0x2D 0x2D (\``<!--`\`)

    :   Advance the `position`{.variable} pointer so that it points at
        the first 0x3E byte which is preceded by two 0x2D bytes (i.e. at
        the end of an ASCII \'\--\>\' sequence) and comes after the 0x3C
        byte that was found. (The two 0x2D bytes can be the same as
        those in the \'\<!\--\' sequence.)

    A sequence of bytes starting with: 0x3C, 0x4D or 0x6D, 0x45 or 0x65, 0x54 or 0x74, 0x41 or 0x61, and one of 0x09, 0x0A, 0x0C, 0x0D, 0x20, 0x2F (case-insensitive ASCII \'\<meta\' followed by a space or slash)

    :   1.  Advance the `position`{.variable} pointer so that it points
            at the next 0x09, 0x0A, 0x0C, 0x0D, 0x20, or 0x2F byte (the
            one in sequence of characters matched above).

        2.  Let `attribute list`{.variable} be an empty list of strings.

        3.  Let `got pragma`{.variable} be false.

        4.  Let `need pragma`{.variable} be null.

        5.  Let `charset`{.variable} be the null value (which, for the
            purposes of this algorithm, is distinct from an unrecognized
            encoding or the empty string).

        6.  *Attributes*: [Get an
            attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing-2}
            and its value. If no attribute was sniffed, then jump to the
            *processing* step below.

        7.  If the attribute\'s name is already in
            `attribute list`{.variable}, then return to the step labeled
            *attributes*.

        8.  Add the attribute\'s name to `attribute list`{.variable}.

        9.  Run the appropriate step from the following list, if one
            applies:

            If the attribute\'s name is \"`http-equiv`\"
            :   If the attribute\'s value is \"`content-type`\", then
                set `got pragma`{.variable} to true.

            If the attribute\'s name is \"`content`\"
            :   Apply the [algorithm for extracting a character encoding
                from a `meta`
                element](urls-and-fetching.html#algorithm-for-extracting-a-character-encoding-from-a-meta-element){#determining-the-character-encoding:algorithm-for-extracting-a-character-encoding-from-a-meta-element},
                giving the attribute\'s value as the string to parse. If
                a character encoding is returned, and if
                `charset`{.variable} is still set to null, let
                `charset`{.variable} be the encoding returned, and set
                `need pragma`{.variable} to true.

            If the attribute\'s name is \"`charset`\"

            :   Let `charset`{.variable} be the result of [getting an
                encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#determining-the-character-encoding:getting-an-encoding
                x-internal="getting-an-encoding"} from the attribute\'s
                value, and set `need pragma`{.variable} to false.

        10. Return to the step labeled *attributes*.

        11. *Processing*: If `need pragma`{.variable} is null, then jump
            to the step below labeled *next byte*.

        12. If `need pragma`{.variable} is true but
            `got pragma`{.variable} is false, then jump to the step
            below labeled *next byte*.

        13. If `charset`{.variable} is failure, then jump to the step
            below labeled *next byte*.

        14. If `charset`{.variable} is
            [UTF-16BE/LE](https://encoding.spec.whatwg.org/#utf-16be-le){#determining-the-character-encoding:utf-16-encoding-2
            x-internal="utf-16-encoding"}, then set `charset`{.variable}
            to
            [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#determining-the-character-encoding:utf-8
            x-internal="utf-8"}.

        15. If `charset`{.variable} is
            [x-user-defined](https://encoding.spec.whatwg.org/#x-user-defined){#determining-the-character-encoding:x-user-defined
            x-internal="x-user-defined"}, then set `charset`{.variable}
            to
            [windows-1252](https://encoding.spec.whatwg.org/#windows-1252){#determining-the-character-encoding:windows-1252-2
            x-internal="windows-1252"}.

        16. Return `charset`{.variable}.

    A sequence of bytes starting with a 0x3C byte (\<), optionally a 0x2F byte (/), and finally a byte in the range 0x41-0x5A or 0x61-0x7A (A-Z or a-z)

    :   1.  Advance the `position`{.variable} pointer so that it points
            at the next 0x09 (HT), 0x0A (LF), 0x0C (FF), 0x0D (CR), 0x20
            (SP), or 0x3E (\>) byte.

        2.  Repeatedly [get an
            attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing-3}
            until no further attributes can be found, then jump to the
            step below labeled *next byte*.

    A sequence of bytes starting with: 0x3C 0x21 (\``<!`\`)\
    A sequence of bytes starting with: 0x3C 0x2F (\``</`\`)\
    A sequence of bytes starting with: 0x3C 0x3F (\``<?`\`)

    :   Advance the `position`{.variable} pointer so that it points at
        the first 0x3E byte (\>) that comes after the 0x3C byte that was
        found.

    Any other byte

    :   Do nothing with that byte.

4.  *Next byte*: Move `position`{.variable} so it points at the next
    byte in the input byte stream, and return to the step above labeled
    *loop*.

When the [prescan a byte stream to determine its
encoding](#prescan-a-byte-stream-to-determine-its-encoding){#determining-the-character-encoding:prescan-a-byte-stream-to-determine-its-encoding-4}
algorithm says to [get an
attribute]{#concept-get-attributes-when-sniffing .dfn}, it means doing
this:

1.  If the byte at `position`{.variable} is one of 0x09 (HT), 0x0A (LF),
    0x0C (FF), 0x0D (CR), 0x20 (SP), or 0x2F (/), then advance
    `position`{.variable} to the next byte and redo this step.

2.  If the byte at `position`{.variable} is 0x3E (\>), then abort the
    [get an
    attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing-4}
    algorithm. There isn\'t one.

3.  Otherwise, the byte at `position`{.variable} is the start of the
    attribute name. Let `attribute name`{.variable} and
    `attribute value`{.variable} be the empty string.

4.  Process the byte at `position`{.variable} as follows:

    If it is 0x3D (=), and the `attribute name`{.variable} is longer than the empty string
    :   Advance `position`{.variable} to the next byte and jump to the
        step below labeled *value*.

    If it is 0x09 (HT), 0x0A (LF), 0x0C (FF), 0x0D (CR), or 0x20 (SP)
    :   Jump to the step below labeled *spaces*.

    If it is 0x2F (/) or 0x3E (\>)
    :   Abort the [get an
        attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing-5}
        algorithm. The attribute\'s name is the value of
        `attribute name`{.variable}, its value is the empty string.

    If it is in the range 0x41 (A) to 0x5A (Z)
    :   Append the code point `b`{.variable}+0x20 to
        `attribute name`{.variable} (where `b`{.variable} is the value
        of the byte at `position`{.variable}). (This converts the input
        to lowercase.)

    Anything else
    :   Append the code point with the same value as the byte at
        `position`{.variable} to `attribute name`{.variable}. (It
        doesn\'t actually matter how bytes outside the ASCII range are
        handled here, since only ASCII bytes can contribute to the
        detection of a character encoding.)

5.  Advance `position`{.variable} to the next byte and return to the
    previous step.

6.  *Spaces*: If the byte at `position`{.variable} is one of 0x09 (HT),
    0x0A (LF), 0x0C (FF), 0x0D (CR), or 0x20 (SP), then advance
    `position`{.variable} to the next byte, then, repeat this step.

7.  If the byte at `position`{.variable} is *not* 0x3D (=), abort the
    [get an
    attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing-6}
    algorithm. The attribute\'s name is the value of
    `attribute name`{.variable}, its value is the empty string.

8.  Advance `position`{.variable} past the 0x3D (=) byte.

9.  *Value*: If the byte at `position`{.variable} is one of 0x09 (HT),
    0x0A (LF), 0x0C (FF), 0x0D (CR), or 0x20 (SP), then advance
    `position`{.variable} to the next byte, then, repeat this step.

10. Process the byte at `position`{.variable} as follows:

    If it is 0x22 (\") or 0x27 (\')

    :   1.  Let `b`{.variable} be the value of the byte at
            `position`{.variable}.
        2.  *Quote loop*: Advance `position`{.variable} to the next
            byte.
        3.  If the value of the byte at `position`{.variable} is the
            value of `b`{.variable}, then advance `position`{.variable}
            to the next byte and abort the \"get an attribute\"
            algorithm. The attribute\'s name is the value of
            `attribute name`{.variable}, and its value is the value of
            `attribute value`{.variable}.
        4.  Otherwise, if the value of the byte at `position`{.variable}
            is in the range 0x41 (A) to 0x5A (Z), then append a code
            point to `attribute value`{.variable} whose value is 0x20
            more than the value of the byte at `position`{.variable}.
        5.  Otherwise, append a code point to
            `attribute value`{.variable} whose value is the same as the
            value of the byte at `position`{.variable}.
        6.  Return to the step above labeled *quote loop*.

    If it is 0x3E (\>)
    :   Abort the [get an
        attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing-7}
        algorithm. The attribute\'s name is the value of
        `attribute name`{.variable}, its value is the empty string.

    If it is in the range 0x41 (A) to 0x5A (Z)
    :   Append a code point `b`{.variable}+0x20 to
        `attribute value`{.variable} (where `b`{.variable} is the value
        of the byte at `position`{.variable}). Advance
        `position`{.variable} to the next byte.

    Anything else
    :   Append a code point with the same value as the byte at
        `position`{.variable} to `attribute value`{.variable}. Advance
        `position`{.variable} to the next byte.

11. Process the byte at `position`{.variable} as follows:

    If it is 0x09 (HT), 0x0A (LF), 0x0C (FF), 0x0D (CR), 0x20 (SP), or 0x3E (\>)
    :   Abort the [get an
        attribute](#concept-get-attributes-when-sniffing){#determining-the-character-encoding:concept-get-attributes-when-sniffing-8}
        algorithm. The attribute\'s name is the value of
        `attribute name`{.variable} and its value is the value of
        `attribute value`{.variable}.

    If it is in the range 0x41 (A) to 0x5A (Z)
    :   Append a code point `b`{.variable}+0x20 to
        `attribute value`{.variable} (where `b`{.variable} is the value
        of the byte at `position`{.variable}).

    Anything else
    :   Append a code point with the same value as the byte at
        `position`{.variable} to `attribute value`{.variable}.

12. Advance `position`{.variable} to the next byte and return to the
    previous step.

When the [prescan a byte stream to determine its
encoding](#prescan-a-byte-stream-to-determine-its-encoding){#determining-the-character-encoding:prescan-a-byte-stream-to-determine-its-encoding-5}
algorithm is aborted without returning an encoding, [get an XML
encoding]{#concept-get-xml-encoding-when-sniffing .dfn} means doing
this.

Looking for syntax resembling an XML declaration, even in
[`text/html`{#determining-the-character-encoding:text/html}](iana.html#text/html),
is necessary for compatibility with existing content.

1.  Let `encodingPosition`{.variable} be a pointer to the start of the
    stream.

2.  If `encodingPosition`{.variable} does not point to the start of a
    byte sequence 0x3C, 0x3F, 0x78, 0x6D, 0x6C (\``<?xml`\`), then
    return failure.

3.  Let `xmlDeclarationEnd`{.variable} be a pointer to the next byte in
    the input byte stream which is 0x3E (\>). If there is no such byte,
    then return failure.

4.  Set `encodingPosition`{.variable} to the position of the first
    occurrence of the subsequence of bytes 0x65, 0x6E, 0x63, 0x6F, 0x64,
    0x69, 0x6E, 0x67 (\``encoding`\`) at or after the current
    `encodingPosition`{.variable}. If there is no such sequence, then
    return failure.

5.  Advance `encodingPosition`{.variable} past the 0x67 (g) byte.

6.  While the byte at `encodingPosition`{.variable} is less than or
    equal to 0x20 (i.e., it is either an ASCII space or control
    character), advance `encodingPosition`{.variable} to the next byte.

7.  If the byte at `encodingPosition`{.variable} is not 0x3D (=), then
    return failure.

8.  Advance `encodingPosition`{.variable} to the next byte.

9.  While the byte at `encodingPosition`{.variable} is less than or
    equal to 0x20 (i.e., it is either an ASCII space or control
    character), advance `encodingPosition`{.variable} to the next byte.

10. Let `quoteMark`{.variable} be the byte at
    `encodingPosition`{.variable}.

11. If `quoteMark`{.variable} is not either 0x22 (\") or 0x27 (\'), then
    return failure.

12. Advance `encodingPosition`{.variable} to the next byte.

13. Let `encodingEndPosition`{.variable} be the position of the next
    occurrence of `quoteMark`{.variable} at or after
    `encodingPosition`{.variable}. If `quoteMark`{.variable} does not
    occur again, then return failure.

14. Let `potentialEncoding`{.variable} be the sequence of the bytes
    between `encodingPosition`{.variable} (inclusive) and
    `encodingEndPosition`{.variable} (exclusive).

15. If `potentialEncoding`{.variable} contains one or more bytes whose
    byte value is 0x20 or below, then return failure.

16. Let `encoding`{.variable} be the result of [getting an
    encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#determining-the-character-encoding:getting-an-encoding-2
    x-internal="getting-an-encoding"} given
    `potentialEncoding`{.variable} [isomorphic
    decoded](https://infra.spec.whatwg.org/#isomorphic-decode){#determining-the-character-encoding:isomorphic-decode
    x-internal="isomorphic-decode"}.

17. If the `encoding`{.variable} is
    [UTF-16BE/LE](https://encoding.spec.whatwg.org/#utf-16be-le){#determining-the-character-encoding:utf-16-encoding-3
    x-internal="utf-16-encoding"}, then change it to
    [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#determining-the-character-encoding:utf-8-2
    x-internal="utf-8"}.

18. Return `encoding`{.variable}.

For the sake of interoperability, user agents should not use a pre-scan
algorithm that returns different results than the one described above.
(But, if you do, please at least let us know, so that we can improve
this algorithm and benefit everyone\...)

##### [13.2.3.3]{.secno} Character encodings[](#character-encodings){.self-link}

User agents must support the encodings defined in Encoding, including,
but not limited to,
[[UTF-8](https://encoding.spec.whatwg.org/#utf-8)]{#utf-8 .dfn},
[[ISO-8859-2](https://encoding.spec.whatwg.org/#iso-8859-2)]{#iso-8859-2
.dfn},
[[ISO-8859-7](https://encoding.spec.whatwg.org/#iso-8859-7)]{#iso-8859-7
.dfn},
[[ISO-8859-8](https://encoding.spec.whatwg.org/#iso-8859-8)]{#iso-8859-8
.dfn},
[[windows-874](https://encoding.spec.whatwg.org/#windows-874)]{#windows-874
.dfn},
[[windows-1250](https://encoding.spec.whatwg.org/#windows-1250)]{#windows-1250
.dfn},
[[windows-1251](https://encoding.spec.whatwg.org/#windows-1251)]{#windows-1251
.dfn},
[[windows-1252](https://encoding.spec.whatwg.org/#windows-1252)]{#windows-1252
.dfn},
[[windows-1254](https://encoding.spec.whatwg.org/#windows-1254)]{#windows-1254
.dfn},
[[windows-1255](https://encoding.spec.whatwg.org/#windows-1255)]{#windows-1255
.dfn},
[[windows-1256](https://encoding.spec.whatwg.org/#windows-1256)]{#windows-1256
.dfn},
[[windows-1257](https://encoding.spec.whatwg.org/#windows-1257)]{#windows-1257
.dfn},
[[windows-1258](https://encoding.spec.whatwg.org/#windows-1258)]{#windows-1258
.dfn}, [[GBK](https://encoding.spec.whatwg.org/#gbk)]{#gbk .dfn},
[[Big5](https://encoding.spec.whatwg.org/#big5)]{#big5 .dfn},
[[ISO-2022-JP](https://encoding.spec.whatwg.org/#iso-2022-jp)]{#iso-2022-jp
.dfn},
[[Shift_JIS](https://encoding.spec.whatwg.org/#shift_jis)]{#shift_jis
.dfn}, [[EUC-KR](https://encoding.spec.whatwg.org/#euc-kr)]{#euc-kr
.dfn},
[[UTF-16BE](https://encoding.spec.whatwg.org/#utf-16be)]{#utf-16be
.dfn},
[[UTF-16LE](https://encoding.spec.whatwg.org/#utf-16le)]{#utf-16le
.dfn},
[[UTF-16BE/LE](https://encoding.spec.whatwg.org/#utf-16be-le)]{#utf-16-encoding
.dfn}, and
[[x-user-defined](https://encoding.spec.whatwg.org/#x-user-defined)]{#x-user-defined
.dfn}. User agents must not support other encodings.

The above prohibits supporting, for example, CESU-8, UTF-7, BOCU-1,
SCSU, EBCDIC, and UTF-32. This specification does not make any attempt
to support prohibited encodings in its algorithms; support and use of
prohibited encodings would thus lead to unexpected behavior.
[\[CESU8\]](references.html#refsCESU8)
[\[UTF7\]](references.html#refsUTF7)
[\[BOCU1\]](references.html#refsBOCU1)
[\[SCSU\]](references.html#refsSCSU)

##### [13.2.3.4]{.secno} Changing the encoding while parsing[](#changing-the-encoding-while-parsing){.self-link}

When the parser requires the user agent to [change the
encoding]{#change-the-encoding .dfn}, it must run the following steps.
This might happen if the [encoding sniffing
algorithm](#encoding-sniffing-algorithm){#changing-the-encoding-while-parsing:encoding-sniffing-algorithm}
described above failed to find a character encoding, or if it found a
character encoding that was not the actual encoding of the file.

1.  If the encoding that is already being used to interpret the input
    stream is
    [UTF-16BE/LE](https://encoding.spec.whatwg.org/#utf-16be-le){#changing-the-encoding-while-parsing:utf-16-encoding
    x-internal="utf-16-encoding"}, then set the
    [confidence](#concept-encoding-confidence){#changing-the-encoding-while-parsing:concept-encoding-confidence}
    to *certain* and return. The new encoding is ignored; if it was
    anything but the same encoding, then it would be clearly incorrect.

2.  If the new encoding is
    [UTF-16BE/LE](https://encoding.spec.whatwg.org/#utf-16be-le){#changing-the-encoding-while-parsing:utf-16-encoding-2
    x-internal="utf-16-encoding"}, then change it to
    [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#changing-the-encoding-while-parsing:utf-8
    x-internal="utf-8"}.

3.  If the new encoding is
    [x-user-defined](https://encoding.spec.whatwg.org/#x-user-defined){#changing-the-encoding-while-parsing:x-user-defined
    x-internal="x-user-defined"}, then change it to
    [windows-1252](https://encoding.spec.whatwg.org/#windows-1252){#changing-the-encoding-while-parsing:windows-1252
    x-internal="windows-1252"}.

4.  If the new encoding is identical or equivalent to the encoding that
    is already being used to interpret the input stream, then set the
    [confidence](#concept-encoding-confidence){#changing-the-encoding-while-parsing:concept-encoding-confidence-2}
    to *certain* and return. This happens when the encoding information
    found in the file matches what the [encoding sniffing
    algorithm](#encoding-sniffing-algorithm){#changing-the-encoding-while-parsing:encoding-sniffing-algorithm-2}
    determined to be the encoding, and in the second pass through the
    parser if the first pass found that the encoding sniffing algorithm
    described in the earlier section failed to find the right encoding.

5.  If all the bytes up to the last byte converted by the current
    decoder have the same Unicode interpretations in both the current
    encoding and the new encoding, and if the user agent supports
    changing the converter on the fly, then the user agent may change to
    the new converter for the encoding on the fly. Set the [document\'s
    character
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#changing-the-encoding-while-parsing:document's-character-encoding
    x-internal="document's-character-encoding"} and the encoding used to
    convert the input stream to the new encoding, set the
    [confidence](#concept-encoding-confidence){#changing-the-encoding-while-parsing:concept-encoding-confidence-3}
    to *certain*, and return.

6.  Otherwise, restart the
    [navigate](browsing-the-web.html#navigate){#changing-the-encoding-while-parsing:navigate}
    algorithm, with
    *[historyHandling](browsing-the-web.html#navigation-hh)* set to
    \"[`replace`{#changing-the-encoding-while-parsing:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\"
    and other inputs kept the same, but this time skip the [encoding
    sniffing
    algorithm](#encoding-sniffing-algorithm){#changing-the-encoding-while-parsing:encoding-sniffing-algorithm-3}
    and instead just set the encoding to the new encoding and the
    [confidence](#concept-encoding-confidence){#changing-the-encoding-while-parsing:concept-encoding-confidence-4}
    to *certain*. Whenever possible, this should be done without
    actually contacting the network layer (the bytes should be re-parsed
    from memory), even if, e.g., the document is marked as not being
    cacheable. If this is not possible and contacting the network layer
    would involve repeating a request that uses a method other than
    \``GET`\`, then instead set the
    [confidence](#concept-encoding-confidence){#changing-the-encoding-while-parsing:concept-encoding-confidence-5}
    to *certain* and ignore the new encoding. The resource will be
    misinterpreted. User agents may notify the user of the situation, to
    aid in application development.

This algorithm is only invoked when a new encoding is found declared on
a
[`meta`{#changing-the-encoding-while-parsing:the-meta-element}](semantics.html#the-meta-element)
element.

##### [13.2.3.5]{.secno} Preprocessing the input stream[](#preprocessing-the-input-stream){.self-link}

The [input stream]{#input-stream .dfn} consists of the characters pushed
into it as the [input byte
stream](#the-input-byte-stream){#preprocessing-the-input-stream:the-input-byte-stream}
is decoded or from the various APIs that directly manipulate the input
stream.

Any occurrences of
[surrogates](https://infra.spec.whatwg.org/#surrogate){#preprocessing-the-input-stream:surrogate
x-internal="surrogate"} are
[surrogate-in-input-stream](#parse-error-surrogate-in-input-stream){#preprocessing-the-input-stream:parse-error-surrogate-in-input-stream}
[parse
errors](#parse-errors){#preprocessing-the-input-stream:parse-errors}.
Any occurrences of
[noncharacters](https://infra.spec.whatwg.org/#noncharacter){#preprocessing-the-input-stream:noncharacter
x-internal="noncharacter"} are
[noncharacter-in-input-stream](#parse-error-noncharacter-in-input-stream){#preprocessing-the-input-stream:parse-error-noncharacter-in-input-stream}
[parse
errors](#parse-errors){#preprocessing-the-input-stream:parse-errors-2}
and any occurrences of
[controls](https://infra.spec.whatwg.org/#control){#preprocessing-the-input-stream:control
x-internal="control"} other than [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#preprocessing-the-input-stream:space-characters
x-internal="space-characters"} and U+0000 NULL characters are
[control-character-in-input-stream](#parse-error-control-character-in-input-stream){#preprocessing-the-input-stream:parse-error-control-character-in-input-stream}
[parse
errors](#parse-errors){#preprocessing-the-input-stream:parse-errors-3}.

The handling of U+0000 NULL characters varies based on where the
characters are found and happens at the later stages of the parsing.
They are either ignored or, for security reasons, replaced with a U+FFFD
REPLACEMENT CHARACTER. This handling is, by necessity, spread across
both the tokenization stage and the tree construction stage.

Before the
[tokenization](#tokenization){#preprocessing-the-input-stream:tokenization}
stage, the input stream must be preprocessed by [normalizing
newlines](https://infra.spec.whatwg.org/#normalize-newlines){#preprocessing-the-input-stream:normalize-newlines
x-internal="normalize-newlines"}. Thus, newlines in HTML DOMs are
represented by U+000A LF characters, and there are never any U+000D CR
characters in the input to the
[tokenization](#tokenization){#preprocessing-the-input-stream:tokenization-2}
stage.

The [next input character]{#next-input-character .dfn} is the first
character in the [input
stream](#input-stream){#preprocessing-the-input-stream:input-stream}
that has not yet been [consumed]{.dfn} or explicitly ignored by the
requirements in this section. Initially, the *[next input
character](#next-input-character)* is the first character in the input.
The [current input character]{#current-input-character .dfn} is the last
character to have been *consumed*.

The [insertion point]{#insertion-point .dfn} is the position (just
before a character or just before the end of the input stream) where
content inserted using
[`document.write()`{#preprocessing-the-input-stream:dom-document-write}](dynamic-markup-insertion.html#dom-document-write)
is actually inserted. The insertion point is relative to the position of
the character immediately after it, it is not an absolute offset into
the input stream. Initially, the insertion point is undefined.

The \"EOF\" character in the tables below is a conceptual character
representing the end of the [input
stream](#input-stream){#preprocessing-the-input-stream:input-stream-2}.
If the parser is a [script-created
parser](dynamic-markup-insertion.html#script-created-parser){#preprocessing-the-input-stream:script-created-parser},
then the end of the [input
stream](#input-stream){#preprocessing-the-input-stream:input-stream-3}
is reached when an [explicit \"EOF\" character]{#explicit-eof-character
.dfn} (inserted by the
[`document.close()`{#preprocessing-the-input-stream:dom-document-close}](dynamic-markup-insertion.html#dom-document-close)
method) is consumed. Otherwise, the \"EOF\" character is not a real
character in the stream, but rather the lack of any further characters.

#### [13.2.4]{.secno} Parse state[](#parse-state){.self-link}

##### [13.2.4.1]{.secno} The insertion mode[](#the-insertion-mode){.self-link}

The [insertion mode]{#insertion-mode .dfn} is a state variable that
controls the primary operation of the tree construction stage.

Initially, the [insertion
mode](#insertion-mode){#the-insertion-mode:insertion-mode} is
\"[initial](#the-initial-insertion-mode){#the-insertion-mode:the-initial-insertion-mode}\".
It can change to \"[before
html](#the-before-html-insertion-mode){#the-insertion-mode:the-before-html-insertion-mode}\",
\"[before
head](#the-before-head-insertion-mode){#the-insertion-mode:the-before-head-insertion-mode}\",
\"[in
head](#parsing-main-inhead){#the-insertion-mode:parsing-main-inhead}\",
\"[in head
noscript](#parsing-main-inheadnoscript){#the-insertion-mode:parsing-main-inheadnoscript}\",
\"[after
head](#the-after-head-insertion-mode){#the-insertion-mode:the-after-head-insertion-mode}\",
\"[in
body](#parsing-main-inbody){#the-insertion-mode:parsing-main-inbody}\",
\"[text](#parsing-main-incdata){#the-insertion-mode:parsing-main-incdata}\",
\"[in
table](#parsing-main-intable){#the-insertion-mode:parsing-main-intable}\",
\"[in table
text](#parsing-main-intabletext){#the-insertion-mode:parsing-main-intabletext}\",
\"[in
caption](#parsing-main-incaption){#the-insertion-mode:parsing-main-incaption}\",
\"[in column
group](#parsing-main-incolgroup){#the-insertion-mode:parsing-main-incolgroup}\",
\"[in table
body](#parsing-main-intbody){#the-insertion-mode:parsing-main-intbody}\",
\"[in row](#parsing-main-intr){#the-insertion-mode:parsing-main-intr}\",
\"[in
cell](#parsing-main-intd){#the-insertion-mode:parsing-main-intd}\",
\"[in
select](#parsing-main-inselect){#the-insertion-mode:parsing-main-inselect}\",
\"[in select in
table](#parsing-main-inselectintable){#the-insertion-mode:parsing-main-inselectintable}\",
\"[in
template](#parsing-main-intemplate){#the-insertion-mode:parsing-main-intemplate}\",
\"[after
body](#parsing-main-afterbody){#the-insertion-mode:parsing-main-afterbody}\",
\"[in
frameset](#parsing-main-inframeset){#the-insertion-mode:parsing-main-inframeset}\",
\"[after
frameset](#parsing-main-afterframeset){#the-insertion-mode:parsing-main-afterframeset}\",
\"[after after
body](#the-after-after-body-insertion-mode){#the-insertion-mode:the-after-after-body-insertion-mode}\",
and \"[after after
frameset](#the-after-after-frameset-insertion-mode){#the-insertion-mode:the-after-after-frameset-insertion-mode}\"
during the course of the parsing, as described in the [tree
construction](#tree-construction){#the-insertion-mode:tree-construction}
stage. The insertion mode affects how tokens are processed and whether
CDATA sections are supported.

Several of these modes, namely \"[in
head](#parsing-main-inhead){#the-insertion-mode:parsing-main-inhead-2}\",
\"[in
body](#parsing-main-inbody){#the-insertion-mode:parsing-main-inbody-2}\",
\"[in
table](#parsing-main-intable){#the-insertion-mode:parsing-main-intable-2}\",
and \"[in
select](#parsing-main-inselect){#the-insertion-mode:parsing-main-inselect-2}\",
are special, in that the other modes defer to them at various times.
When the algorithm below says that the user agent is to do something
\"[using the rules for]{#using-the-rules-for .dfn} the `m`{.variable}
insertion mode\", where `m`{.variable} is one of these modes, the user
agent must use the rules described under the `m`{.variable} [insertion
mode](#insertion-mode){#the-insertion-mode:insertion-mode-2}\'s section,
but must leave the [insertion
mode](#insertion-mode){#the-insertion-mode:insertion-mode-3} unchanged
unless the rules in `m`{.variable} themselves switch the [insertion
mode](#insertion-mode){#the-insertion-mode:insertion-mode-4} to a new
value.

When the insertion mode is switched to
\"[text](#parsing-main-incdata){#the-insertion-mode:parsing-main-incdata-2}\"
or \"[in table
text](#parsing-main-intabletext){#the-insertion-mode:parsing-main-intabletext-2}\",
the [original insertion mode]{#original-insertion-mode .dfn} is also
set. This is the insertion mode to which the tree construction stage
will return.

Similarly, to parse nested
[`template`{#the-insertion-mode:the-template-element}](scripting.html#the-template-element)
elements, a [stack of template insertion
modes]{#stack-of-template-insertion-modes .dfn} is used. It is initially
empty. The [current template insertion
mode]{#current-template-insertion-mode .dfn} is the insertion mode that
was most recently added to the [stack of template insertion
modes](#stack-of-template-insertion-modes){#the-insertion-mode:stack-of-template-insertion-modes}.
The algorithms in the sections below will *push* insertion modes onto
this stack, meaning that the specified insertion mode is to be added to
the stack, and *pop* insertion modes from the stack, which means that
the most recently added insertion mode must be removed from the stack.

------------------------------------------------------------------------

When the steps below require the UA to [reset the insertion mode
appropriately]{#reset-the-insertion-mode-appropriately .dfn}, it means
the UA must follow these steps:

1.  Let `last`{.variable} be false.

2.  Let `node`{.variable} be the last node in the [stack of open
    elements](#stack-of-open-elements){#the-insertion-mode:stack-of-open-elements}.

3.  *Loop*: If `node`{.variable} is the first node in the stack of open
    elements, then set `last`{.variable} to true, and, if the parser was
    created as part of the [HTML fragment parsing
    algorithm](#html-fragment-parsing-algorithm){#the-insertion-mode:html-fragment-parsing-algorithm}
    ([fragment
    case](#fragment-case){#the-insertion-mode:fragment-case}), set
    `node`{.variable} to the
    [`context`{#the-insertion-mode:concept-frag-parse-context
    .variable}](#concept-frag-parse-context) element passed to that
    algorithm.

4.  If `node`{.variable} is a
    [`select`{#the-insertion-mode:the-select-element}](form-elements.html#the-select-element)
    element, run these substeps:

    1.  If `last`{.variable} is true, jump to the step below labeled
        *done*.

    2.  Let `ancestor`{.variable} be `node`{.variable}.

    3.  *Loop*: If `ancestor`{.variable} is the first node in the [stack
        of open
        elements](#stack-of-open-elements){#the-insertion-mode:stack-of-open-elements-2},
        jump to the step below labeled *done*.

    4.  Let `ancestor`{.variable} be the node before
        `ancestor`{.variable} in the [stack of open
        elements](#stack-of-open-elements){#the-insertion-mode:stack-of-open-elements-3}.

    5.  If `ancestor`{.variable} is a
        [`template`{#the-insertion-mode:the-template-element-2}](scripting.html#the-template-element)
        node, jump to the step below labeled *done*.

    6.  If `ancestor`{.variable} is a
        [`table`{#the-insertion-mode:the-table-element}](tables.html#the-table-element)
        node, switch the [insertion
        mode](#insertion-mode){#the-insertion-mode:insertion-mode-5} to
        \"[in select in
        table](#parsing-main-inselectintable){#the-insertion-mode:parsing-main-inselectintable-2}\"
        and return.

    7.  Jump back to the step labeled *loop*.

    8.  *Done*: Switch the [insertion
        mode](#insertion-mode){#the-insertion-mode:insertion-mode-6} to
        \"[in
        select](#parsing-main-inselect){#the-insertion-mode:parsing-main-inselect-3}\"
        and return.

5.  If `node`{.variable} is a
    [`td`{#the-insertion-mode:the-td-element}](tables.html#the-td-element)
    or
    [`th`{#the-insertion-mode:the-th-element}](tables.html#the-th-element)
    element and `last`{.variable} is false, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-7} to
    \"[in
    cell](#parsing-main-intd){#the-insertion-mode:parsing-main-intd-2}\"
    and return.

6.  If `node`{.variable} is a
    [`tr`{#the-insertion-mode:the-tr-element}](tables.html#the-tr-element)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-8} to
    \"[in
    row](#parsing-main-intr){#the-insertion-mode:parsing-main-intr-2}\"
    and return.

7.  If `node`{.variable} is a
    [`tbody`{#the-insertion-mode:the-tbody-element}](tables.html#the-tbody-element),
    [`thead`{#the-insertion-mode:the-thead-element}](tables.html#the-thead-element),
    or
    [`tfoot`{#the-insertion-mode:the-tfoot-element}](tables.html#the-tfoot-element)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-9} to
    \"[in table
    body](#parsing-main-intbody){#the-insertion-mode:parsing-main-intbody-2}\"
    and return.

8.  If `node`{.variable} is a
    [`caption`{#the-insertion-mode:the-caption-element}](tables.html#the-caption-element)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-10} to
    \"[in
    caption](#parsing-main-incaption){#the-insertion-mode:parsing-main-incaption-2}\"
    and return.

9.  If `node`{.variable} is a
    [`colgroup`{#the-insertion-mode:the-colgroup-element}](tables.html#the-colgroup-element)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-11} to
    \"[in column
    group](#parsing-main-incolgroup){#the-insertion-mode:parsing-main-incolgroup-2}\"
    and return.

10. If `node`{.variable} is a
    [`table`{#the-insertion-mode:the-table-element-2}](tables.html#the-table-element)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-12} to
    \"[in
    table](#parsing-main-intable){#the-insertion-mode:parsing-main-intable-3}\"
    and return.

11. If `node`{.variable} is a
    [`template`{#the-insertion-mode:the-template-element-3}](scripting.html#the-template-element)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-13} to the
    [current template insertion
    mode](#current-template-insertion-mode){#the-insertion-mode:current-template-insertion-mode}
    and return.

12. If `node`{.variable} is a
    [`head`{#the-insertion-mode:the-head-element}](semantics.html#the-head-element)
    element and `last`{.variable} is false, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-14} to
    \"[in
    head](#parsing-main-inhead){#the-insertion-mode:parsing-main-inhead-3}\"
    and return.

13. If `node`{.variable} is a
    [`body`{#the-insertion-mode:the-body-element}](sections.html#the-body-element)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-15} to
    \"[in
    body](#parsing-main-inbody){#the-insertion-mode:parsing-main-inbody-3}\"
    and return.

14. If `node`{.variable} is a
    [`frameset`{#the-insertion-mode:frameset}](obsolete.html#frameset)
    element, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-16} to
    \"[in
    frameset](#parsing-main-inframeset){#the-insertion-mode:parsing-main-inframeset-2}\"
    and return. ([fragment
    case](#fragment-case){#the-insertion-mode:fragment-case-2})

15. If `node`{.variable} is an
    [`html`{#the-insertion-mode:the-html-element}](semantics.html#the-html-element)
    element, run these substeps:

    1.  If the [`head` element
        pointer](#head-element-pointer){#the-insertion-mode:head-element-pointer}
        is null, switch the [insertion
        mode](#insertion-mode){#the-insertion-mode:insertion-mode-17} to
        \"[before
        head](#the-before-head-insertion-mode){#the-insertion-mode:the-before-head-insertion-mode-2}\"
        and return. ([fragment
        case](#fragment-case){#the-insertion-mode:fragment-case-3})

    2.  Otherwise, the [`head` element
        pointer](#head-element-pointer){#the-insertion-mode:head-element-pointer-2}
        is not null, switch the [insertion
        mode](#insertion-mode){#the-insertion-mode:insertion-mode-18} to
        \"[after
        head](#the-after-head-insertion-mode){#the-insertion-mode:the-after-head-insertion-mode-2}\"
        and return.

16. If `last`{.variable} is true, then switch the [insertion
    mode](#insertion-mode){#the-insertion-mode:insertion-mode-19} to
    \"[in
    body](#parsing-main-inbody){#the-insertion-mode:parsing-main-inbody-4}\"
    and return. ([fragment
    case](#fragment-case){#the-insertion-mode:fragment-case-4})

17. Let `node`{.variable} now be the node before `node`{.variable} in
    the [stack of open
    elements](#stack-of-open-elements){#the-insertion-mode:stack-of-open-elements-4}.

18. Return to the step labeled *loop*.

##### [13.2.4.2]{.secno} The stack of open elements[](#the-stack-of-open-elements){.self-link}

Initially, the [stack of open elements]{#stack-of-open-elements .dfn} is
empty. The stack grows downwards; the topmost node on the stack is the
first one added to the stack, and the bottommost node of the stack is
the most recently added node in the stack (notwithstanding when the
stack is manipulated in a random access fashion as part of [the handling
for misnested tags](#adoptionAgency)).

The \"[before
html](#the-before-html-insertion-mode){#the-stack-of-open-elements:the-before-html-insertion-mode}\"
[insertion
mode](#insertion-mode){#the-stack-of-open-elements:insertion-mode}
creates the
[`html`{#the-stack-of-open-elements:the-html-element}](semantics.html#the-html-element)
[document
element](https://dom.spec.whatwg.org/#document-element){#the-stack-of-open-elements:document-element
x-internal="document-element"}, which is then added to the stack.

In the [fragment
case](#fragment-case){#the-stack-of-open-elements:fragment-case}, the
[stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements}
is initialized to contain an
[`html`{#the-stack-of-open-elements:the-html-element-2}](semantics.html#the-html-element)
element that is created as part of [that
algorithm](#html-fragment-parsing-algorithm){#the-stack-of-open-elements:html-fragment-parsing-algorithm}.
(The [fragment
case](#fragment-case){#the-stack-of-open-elements:fragment-case-2} skips
the \"[before
html](#the-before-html-insertion-mode){#the-stack-of-open-elements:the-before-html-insertion-mode-2}\"
[insertion
mode](#insertion-mode){#the-stack-of-open-elements:insertion-mode-2}.)

The
[`html`{#the-stack-of-open-elements:the-html-element-3}](semantics.html#the-html-element)
node, however it is created, is the topmost node of the stack. It only
gets popped off the stack when the parser
[finishes](#stop-parsing){#the-stack-of-open-elements:stop-parsing}.

The [current node]{#current-node .dfn} is the bottommost node in this
[stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-2}.

The [adjusted current node]{#adjusted-current-node .dfn} is the
*[context](#concept-frag-parse-context)* element if the parser was
created as part of the [HTML fragment parsing
algorithm](#html-fragment-parsing-algorithm){#the-stack-of-open-elements:html-fragment-parsing-algorithm-2}
and the [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-3}
has only one element in it ([fragment
case](#fragment-case){#the-stack-of-open-elements:fragment-case-3});
otherwise, the [adjusted current
node](#adjusted-current-node){#the-stack-of-open-elements:adjusted-current-node}
is the [current
node](#current-node){#the-stack-of-open-elements:current-node}.

When the [current
node](#current-node){#the-stack-of-open-elements:current-node-2} is
removed from the [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-4},
[process internal resource
links](links.html#process-internal-resource-links){#the-stack-of-open-elements:process-internal-resource-links}
given the [current
node](#current-node){#the-stack-of-open-elements:current-node-3}\'s
[node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-stack-of-open-elements:node-document
x-internal="node-document"}.

Elements in the [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-5}
fall into the following categories:

[Special]{#special .dfn}

:   The following elements have varying levels of special parsing rules:
    HTML\'s
    [`address`{#the-stack-of-open-elements:the-address-element}](sections.html#the-address-element),
    [`applet`{#the-stack-of-open-elements:applet}](obsolete.html#applet),
    [`area`{#the-stack-of-open-elements:the-area-element}](image-maps.html#the-area-element),
    [`article`{#the-stack-of-open-elements:the-article-element}](sections.html#the-article-element),
    [`aside`{#the-stack-of-open-elements:the-aside-element}](sections.html#the-aside-element),
    [`base`{#the-stack-of-open-elements:the-base-element}](semantics.html#the-base-element),
    [`basefont`{#the-stack-of-open-elements:basefont}](obsolete.html#basefont),
    [`bgsound`{#the-stack-of-open-elements:bgsound}](obsolete.html#bgsound),
    [`blockquote`{#the-stack-of-open-elements:the-blockquote-element}](grouping-content.html#the-blockquote-element),
    [`body`{#the-stack-of-open-elements:the-body-element}](sections.html#the-body-element),
    [`br`{#the-stack-of-open-elements:the-br-element}](text-level-semantics.html#the-br-element),
    [`button`{#the-stack-of-open-elements:the-button-element}](form-elements.html#the-button-element),
    [`caption`{#the-stack-of-open-elements:the-caption-element}](tables.html#the-caption-element),
    [`center`{#the-stack-of-open-elements:center}](obsolete.html#center),
    [`col`{#the-stack-of-open-elements:the-col-element}](tables.html#the-col-element),
    [`colgroup`{#the-stack-of-open-elements:the-colgroup-element}](tables.html#the-colgroup-element),
    [`dd`{#the-stack-of-open-elements:the-dd-element}](grouping-content.html#the-dd-element),
    [`details`{#the-stack-of-open-elements:the-details-element}](interactive-elements.html#the-details-element),
    [`dir`{#the-stack-of-open-elements:dir}](obsolete.html#dir),
    [`div`{#the-stack-of-open-elements:the-div-element}](grouping-content.html#the-div-element),
    [`dl`{#the-stack-of-open-elements:the-dl-element}](grouping-content.html#the-dl-element),
    [`dt`{#the-stack-of-open-elements:the-dt-element}](grouping-content.html#the-dt-element),
    [`embed`{#the-stack-of-open-elements:the-embed-element}](iframe-embed-object.html#the-embed-element),
    [`fieldset`{#the-stack-of-open-elements:the-fieldset-element}](form-elements.html#the-fieldset-element),
    [`figcaption`{#the-stack-of-open-elements:the-figcaption-element}](grouping-content.html#the-figcaption-element),
    [`figure`{#the-stack-of-open-elements:the-figure-element}](grouping-content.html#the-figure-element),
    [`footer`{#the-stack-of-open-elements:the-footer-element}](sections.html#the-footer-element),
    [`form`{#the-stack-of-open-elements:the-form-element}](forms.html#the-form-element),
    [`frame`{#the-stack-of-open-elements:frame}](obsolete.html#frame),
    [`frameset`{#the-stack-of-open-elements:frameset}](obsolete.html#frameset),
    [`h1`{#the-stack-of-open-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h2`{#the-stack-of-open-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-2}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h3`{#the-stack-of-open-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-3}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h4`{#the-stack-of-open-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-4}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h5`{#the-stack-of-open-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-5}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`h6`{#the-stack-of-open-elements:the-h1,-h2,-h3,-h4,-h5,-and-h6-elements-6}](sections.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements),
    [`head`{#the-stack-of-open-elements:the-head-element}](semantics.html#the-head-element),
    [`header`{#the-stack-of-open-elements:the-header-element}](sections.html#the-header-element),
    [`hgroup`{#the-stack-of-open-elements:the-hgroup-element}](sections.html#the-hgroup-element),
    [`hr`{#the-stack-of-open-elements:the-hr-element}](grouping-content.html#the-hr-element),
    [`html`{#the-stack-of-open-elements:the-html-element-4}](semantics.html#the-html-element),
    [`iframe`{#the-stack-of-open-elements:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
    [`img`{#the-stack-of-open-elements:the-img-element}](embedded-content.html#the-img-element),
    [`input`{#the-stack-of-open-elements:the-input-element}](input.html#the-input-element),
    [`keygen`{#the-stack-of-open-elements:keygen}](obsolete.html#keygen),
    [`li`{#the-stack-of-open-elements:the-li-element}](grouping-content.html#the-li-element),
    [`link`{#the-stack-of-open-elements:the-link-element}](semantics.html#the-link-element),
    [`listing`{#the-stack-of-open-elements:listing}](obsolete.html#listing),
    [`main`{#the-stack-of-open-elements:the-main-element}](grouping-content.html#the-main-element),
    [`marquee`{#the-stack-of-open-elements:the-marquee-element}](obsolete.html#the-marquee-element),
    [`menu`{#the-stack-of-open-elements:the-menu-element}](grouping-content.html#the-menu-element),
    [`meta`{#the-stack-of-open-elements:the-meta-element}](semantics.html#the-meta-element),
    [`nav`{#the-stack-of-open-elements:the-nav-element}](sections.html#the-nav-element),
    [`noembed`{#the-stack-of-open-elements:noembed}](obsolete.html#noembed),
    [`noframes`{#the-stack-of-open-elements:noframes}](obsolete.html#noframes),
    [`noscript`{#the-stack-of-open-elements:the-noscript-element}](scripting.html#the-noscript-element),
    [`object`{#the-stack-of-open-elements:the-object-element}](iframe-embed-object.html#the-object-element),
    [`ol`{#the-stack-of-open-elements:the-ol-element}](grouping-content.html#the-ol-element),
    [`p`{#the-stack-of-open-elements:the-p-element}](grouping-content.html#the-p-element),
    [`param`{#the-stack-of-open-elements:param}](obsolete.html#param),
    [`plaintext`{#the-stack-of-open-elements:plaintext}](obsolete.html#plaintext),
    [`pre`{#the-stack-of-open-elements:the-pre-element}](grouping-content.html#the-pre-element),
    [`script`{#the-stack-of-open-elements:the-script-element}](scripting.html#the-script-element),
    [`search`{#the-stack-of-open-elements:the-search-element}](grouping-content.html#the-search-element),
    [`section`{#the-stack-of-open-elements:the-section-element}](sections.html#the-section-element),
    [`select`{#the-stack-of-open-elements:the-select-element}](form-elements.html#the-select-element),
    [`source`{#the-stack-of-open-elements:the-source-element}](embedded-content.html#the-source-element),
    [`style`{#the-stack-of-open-elements:the-style-element}](semantics.html#the-style-element),
    [`summary`{#the-stack-of-open-elements:the-summary-element}](interactive-elements.html#the-summary-element),
    [`table`{#the-stack-of-open-elements:the-table-element}](tables.html#the-table-element),
    [`tbody`{#the-stack-of-open-elements:the-tbody-element}](tables.html#the-tbody-element),
    [`td`{#the-stack-of-open-elements:the-td-element}](tables.html#the-td-element),
    [`template`{#the-stack-of-open-elements:the-template-element}](scripting.html#the-template-element),
    [`textarea`{#the-stack-of-open-elements:the-textarea-element}](form-elements.html#the-textarea-element),
    [`tfoot`{#the-stack-of-open-elements:the-tfoot-element}](tables.html#the-tfoot-element),
    [`th`{#the-stack-of-open-elements:the-th-element}](tables.html#the-th-element),
    [`thead`{#the-stack-of-open-elements:the-thead-element}](tables.html#the-thead-element),
    [`title`{#the-stack-of-open-elements:the-title-element}](semantics.html#the-title-element),
    [`tr`{#the-stack-of-open-elements:the-tr-element}](tables.html#the-tr-element),
    [`track`{#the-stack-of-open-elements:the-track-element}](media.html#the-track-element),
    [`ul`{#the-stack-of-open-elements:the-ul-element}](grouping-content.html#the-ul-element),
    [`wbr`{#the-stack-of-open-elements:the-wbr-element}](text-level-semantics.html#the-wbr-element),
    [`xmp`{#the-stack-of-open-elements:xmp}](obsolete.html#xmp); [MathML
    `mi`](https://w3c.github.io/mathml-core/#the-mi-element){#the-stack-of-open-elements:mathml-mi
    x-internal="mathml-mi"}, [MathML
    `mo`](https://w3c.github.io/mathml-core/#operator-fence-separator-or-accent-mo){#the-stack-of-open-elements:mathml-mo
    x-internal="mathml-mo"}, [MathML
    `mn`](https://w3c.github.io/mathml-core/#number-mn){#the-stack-of-open-elements:mathml-mn
    x-internal="mathml-mn"}, [MathML
    `ms`](https://w3c.github.io/mathml-core/#string-literal-ms){#the-stack-of-open-elements:mathml-ms
    x-internal="mathml-ms"}, [MathML
    `mtext`](https://w3c.github.io/mathml-core/#text-mtext){#the-stack-of-open-elements:mathml-mtext
    x-internal="mathml-mtext"}, and [MathML
    `annotation-xml`](https://w3c.github.io/mathml-core/#dfn-annotation-xml){#the-stack-of-open-elements:mathml-annotation-xml
    x-internal="mathml-annotation-xml"}; and [SVG
    `foreignObject`](https://svgwg.org/svg2-draft/embedded.html#ForeignObjectElement){#the-stack-of-open-elements:svg-foreignobject
    x-internal="svg-foreignobject"}, [SVG
    `desc`](https://svgwg.org/svg2-draft/struct.html#DescElement){#the-stack-of-open-elements:svg-desc
    x-internal="svg-desc"}, and [SVG
    `title`](https://svgwg.org/svg2-draft/struct.html#TitleElement){#the-stack-of-open-elements:svg-title
    x-internal="svg-title"}.

    An `image` start tag token is handled by the tree builder, but it is
    not in this list because it is not an element; it gets turned into
    an
    [`img`{#the-stack-of-open-elements:the-img-element-2}](embedded-content.html#the-img-element)
    element.

[Formatting]{#formatting .dfn}
:   The following HTML elements are those that end up in the [list of
    active formatting
    elements](#list-of-active-formatting-elements){#the-stack-of-open-elements:list-of-active-formatting-elements}:
    [`a`{#the-stack-of-open-elements:the-a-element}](text-level-semantics.html#the-a-element),
    [`b`{#the-stack-of-open-elements:the-b-element}](text-level-semantics.html#the-b-element),
    [`big`{#the-stack-of-open-elements:big}](obsolete.html#big),
    [`code`{#the-stack-of-open-elements:the-code-element}](text-level-semantics.html#the-code-element),
    [`em`{#the-stack-of-open-elements:the-em-element}](text-level-semantics.html#the-em-element),
    [`font`{#the-stack-of-open-elements:font}](obsolete.html#font),
    [`i`{#the-stack-of-open-elements:the-i-element}](text-level-semantics.html#the-i-element),
    [`nobr`{#the-stack-of-open-elements:nobr}](obsolete.html#nobr),
    [`s`{#the-stack-of-open-elements:the-s-element}](text-level-semantics.html#the-s-element),
    [`small`{#the-stack-of-open-elements:the-small-element}](text-level-semantics.html#the-small-element),
    [`strike`{#the-stack-of-open-elements:strike}](obsolete.html#strike),
    [`strong`{#the-stack-of-open-elements:the-strong-element}](text-level-semantics.html#the-strong-element),
    [`tt`{#the-stack-of-open-elements:tt}](obsolete.html#tt), and
    [`u`{#the-stack-of-open-elements:the-u-element}](text-level-semantics.html#the-u-element).

[Ordinary]{#ordinary .dfn}

:   All other elements found while parsing an HTML document.

Typically, the [special](#special){#the-stack-of-open-elements:special}
elements have the start and end tag tokens handled specifically, while
[ordinary](#ordinary){#the-stack-of-open-elements:ordinary} elements\'
tokens fall into \"any other start tag\" and \"any other end tag\"
clauses, and some parts of the tree builder check if a particular
element in the [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-6}
is in the [special](#special){#the-stack-of-open-elements:special-2}
category. However, some elements (e.g., the
[`option`{#the-stack-of-open-elements:the-option-element}](form-elements.html#the-option-element)
element) have their start or end tag tokens handled specifically, but
are still not in the
[special](#special){#the-stack-of-open-elements:special-3} category, so
that they get the
[ordinary](#ordinary){#the-stack-of-open-elements:ordinary-2} handling
elsewhere.

The [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-7}
is said to [have an element `target node`{.variable} in a specific
scope]{#has-an-element-in-the-specific-scope .dfn} consisting of a list
of element types `list`{.variable} when the following algorithm
terminates in a match state:

1.  Initialize `node`{.variable} to be the [current
    node](#current-node){#the-stack-of-open-elements:current-node-4}
    (the bottommost node of the stack).

2.  If `node`{.variable} is `target node`{.variable}, terminate in a
    match state.

3.  Otherwise, if `node`{.variable} is one of the element types in
    `list`{.variable}, terminate in a failure state.

4.  Otherwise, set `node`{.variable} to the previous entry in the [stack
    of open
    elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-8}
    and return to step 2. (This will never fail, since the loop will
    always terminate in the previous step if the top of the stack --- an
    [`html`{#the-stack-of-open-elements:the-html-element-5}](semantics.html#the-html-element)
    element --- is reached.)

The [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-9}
is said to [have a particular element in scope]{#has-an-element-in-scope
.dfn} when it [has that element in the specific
scope](#has-an-element-in-the-specific-scope){#the-stack-of-open-elements:has-an-element-in-the-specific-scope}
consisting of the following element types:

- [`applet`{#the-stack-of-open-elements:applet-2}](obsolete.html#applet)
- [`caption`{#the-stack-of-open-elements:the-caption-element-2}](tables.html#the-caption-element)
- [`html`{#the-stack-of-open-elements:the-html-element-6}](semantics.html#the-html-element)
- [`table`{#the-stack-of-open-elements:the-table-element-2}](tables.html#the-table-element)
- [`td`{#the-stack-of-open-elements:the-td-element-2}](tables.html#the-td-element)
- [`th`{#the-stack-of-open-elements:the-th-element-2}](tables.html#the-th-element)
- [`marquee`{#the-stack-of-open-elements:the-marquee-element-2}](obsolete.html#the-marquee-element)
- [`object`{#the-stack-of-open-elements:the-object-element-2}](iframe-embed-object.html#the-object-element)
- [`template`{#the-stack-of-open-elements:the-template-element-2}](scripting.html#the-template-element)
- [MathML
  `mi`](https://w3c.github.io/mathml-core/#the-mi-element){#the-stack-of-open-elements:mathml-mi-2
  x-internal="mathml-mi"}
- [MathML
  `mo`](https://w3c.github.io/mathml-core/#operator-fence-separator-or-accent-mo){#the-stack-of-open-elements:mathml-mo-2
  x-internal="mathml-mo"}
- [MathML
  `mn`](https://w3c.github.io/mathml-core/#number-mn){#the-stack-of-open-elements:mathml-mn-2
  x-internal="mathml-mn"}
- [MathML
  `ms`](https://w3c.github.io/mathml-core/#string-literal-ms){#the-stack-of-open-elements:mathml-ms-2
  x-internal="mathml-ms"}
- [MathML
  `mtext`](https://w3c.github.io/mathml-core/#text-mtext){#the-stack-of-open-elements:mathml-mtext-2
  x-internal="mathml-mtext"}
- [MathML
  `annotation-xml`](https://w3c.github.io/mathml-core/#dfn-annotation-xml){#the-stack-of-open-elements:mathml-annotation-xml-2
  x-internal="mathml-annotation-xml"}
- [SVG
  `foreignObject`](https://svgwg.org/svg2-draft/embedded.html#ForeignObjectElement){#the-stack-of-open-elements:svg-foreignobject-2
  x-internal="svg-foreignobject"}
- [SVG
  `desc`](https://svgwg.org/svg2-draft/struct.html#DescElement){#the-stack-of-open-elements:svg-desc-2
  x-internal="svg-desc"}
- [SVG
  `title`](https://svgwg.org/svg2-draft/struct.html#TitleElement){#the-stack-of-open-elements:svg-title-2
  x-internal="svg-title"}

The [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-10}
is said to [have a particular element in list item
scope]{#has-an-element-in-list-item-scope .dfn} when it [has that
element in the specific
scope](#has-an-element-in-the-specific-scope){#the-stack-of-open-elements:has-an-element-in-the-specific-scope-2}
consisting of the following element types:

- All the element types listed above for the *[has an element in
  scope](#has-an-element-in-scope)* algorithm.
- [`ol`{#the-stack-of-open-elements:the-ol-element-2}](grouping-content.html#the-ol-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2
  x-internal="html-namespace-2"}
- [`ul`{#the-stack-of-open-elements:the-ul-element-2}](grouping-content.html#the-ul-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2-2
  x-internal="html-namespace-2"}

The [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-11}
is said to [have a particular element in button
scope]{#has-an-element-in-button-scope .dfn} when it [has that element
in the specific
scope](#has-an-element-in-the-specific-scope){#the-stack-of-open-elements:has-an-element-in-the-specific-scope-3}
consisting of the following element types:

- All the element types listed above for the *[has an element in
  scope](#has-an-element-in-scope)* algorithm.
- [`button`{#the-stack-of-open-elements:the-button-element-2}](form-elements.html#the-button-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2-3
  x-internal="html-namespace-2"}

The [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-12}
is said to [have a particular element in table
scope]{#has-an-element-in-table-scope .dfn} when it [has that element in
the specific
scope](#has-an-element-in-the-specific-scope){#the-stack-of-open-elements:has-an-element-in-the-specific-scope-4}
consisting of the following element types:

- [`html`{#the-stack-of-open-elements:the-html-element-7}](semantics.html#the-html-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2-4
  x-internal="html-namespace-2"}
- [`table`{#the-stack-of-open-elements:the-table-element-3}](tables.html#the-table-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2-5
  x-internal="html-namespace-2"}
- [`template`{#the-stack-of-open-elements:the-template-element-3}](scripting.html#the-template-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2-6
  x-internal="html-namespace-2"}

The [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-13}
is said to [have a particular element in select
scope]{#has-an-element-in-select-scope .dfn} when it [has that element
in the specific
scope](#has-an-element-in-the-specific-scope){#the-stack-of-open-elements:has-an-element-in-the-specific-scope-5}
consisting of all element types *except* the following:

- [`optgroup`{#the-stack-of-open-elements:the-optgroup-element}](form-elements.html#the-optgroup-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2-7
  x-internal="html-namespace-2"}
- [`option`{#the-stack-of-open-elements:the-option-element-2}](form-elements.html#the-option-element)
  in the [HTML
  namespace](https://infra.spec.whatwg.org/#html-namespace){#the-stack-of-open-elements:html-namespace-2-8
  x-internal="html-namespace-2"}

Nothing happens if at any time any of the elements in the [stack of open
elements](#stack-of-open-elements){#the-stack-of-open-elements:stack-of-open-elements-14}
are moved to a new location in, or removed from, the
[`Document`{#the-stack-of-open-elements:document}](dom.html#document)
tree. In particular, the stack is not changed in this situation. This
can cause, amongst other strange effects, content to be appended to
nodes that are no longer in the DOM.

In some cases (namely, when [closing misnested formatting
elements](#adoptionAgency)), the stack is manipulated in a random-access
fashion.

##### [13.2.4.3]{.secno} The list of active formatting elements[](#the-list-of-active-formatting-elements){.self-link}

Initially, the [list of active formatting
elements]{#list-of-active-formatting-elements .dfn} is empty. It is used
to handle mis-nested [formatting element
tags](#formatting){#the-list-of-active-formatting-elements:formatting}.

The list contains elements in the
[formatting](#formatting){#the-list-of-active-formatting-elements:formatting-2}
category, and
[markers](#concept-parser-marker){#the-list-of-active-formatting-elements:concept-parser-marker}.
The [markers]{#concept-parser-marker .dfn} are inserted when entering
[`applet`{#the-list-of-active-formatting-elements:applet}](obsolete.html#applet),
[`object`{#the-list-of-active-formatting-elements:the-object-element}](iframe-embed-object.html#the-object-element),
[`marquee`{#the-list-of-active-formatting-elements:the-marquee-element}](obsolete.html#the-marquee-element),
[`template`{#the-list-of-active-formatting-elements:the-template-element}](scripting.html#the-template-element),
[`td`{#the-list-of-active-formatting-elements:the-td-element}](tables.html#the-td-element),
[`th`{#the-list-of-active-formatting-elements:the-th-element}](tables.html#the-th-element),
and
[`caption`{#the-list-of-active-formatting-elements:the-caption-element}](tables.html#the-caption-element)
elements, and are used to prevent formatting from \"leaking\" *into*
[`applet`{#the-list-of-active-formatting-elements:applet-2}](obsolete.html#applet),
[`object`{#the-list-of-active-formatting-elements:the-object-element-2}](iframe-embed-object.html#the-object-element),
[`marquee`{#the-list-of-active-formatting-elements:the-marquee-element-2}](obsolete.html#the-marquee-element),
[`template`{#the-list-of-active-formatting-elements:the-template-element-2}](scripting.html#the-template-element),
[`td`{#the-list-of-active-formatting-elements:the-td-element-2}](tables.html#the-td-element),
[`th`{#the-list-of-active-formatting-elements:the-th-element-2}](tables.html#the-th-element),
and
[`caption`{#the-list-of-active-formatting-elements:the-caption-element-2}](tables.html#the-caption-element)
elements.

In addition, each element in the [list of active formatting
elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements}
is associated with the token for which it was created, so that further
elements can be created for that token if necessary.

When the steps below require the UA to [push onto the list of active
formatting elements]{#push-onto-the-list-of-active-formatting-elements
.dfn} an element `element`{.variable}, the UA must perform the following
steps:

1.  If there are already three elements in the [list of active
    formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-2}
    after the last
    [marker](#concept-parser-marker){#the-list-of-active-formatting-elements:concept-parser-marker-2},
    if any, or anywhere in the list if there are no
    [markers](#concept-parser-marker){#the-list-of-active-formatting-elements:concept-parser-marker-3},
    that have the same tag name, namespace, and attributes as
    `element`{.variable}, then remove the earliest such element from the
    [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-3}.
    For these purposes, the attributes must be compared as they were
    when the elements were created by the parser; two elements have the
    same attributes if all their parsed attributes can be paired such
    that the two attributes in each pair have identical names,
    namespaces, and values (the order of the attributes does not
    matter).

    This is the Noah\'s Ark clause. But with three per family instead of
    two.

2.  Add `element`{.variable} to the [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-4}.

When the steps below require the UA to [reconstruct the active
formatting elements]{#reconstruct-the-active-formatting-elements .dfn},
the UA must perform the following steps:

1.  If there are no entries in the [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-5},
    then there is nothing to reconstruct; stop this algorithm.

2.  If the last (most recently added) entry in the [list of active
    formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-6}
    is a
    [marker](#concept-parser-marker){#the-list-of-active-formatting-elements:concept-parser-marker-4},
    or if it is an element that is in the [stack of open
    elements](#stack-of-open-elements){#the-list-of-active-formatting-elements:stack-of-open-elements},
    then there is nothing to reconstruct; stop this algorithm.

3.  Let `entry`{.variable} be the last (most recently added) element in
    the [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-7}.

4.  *Rewind*: If there are no entries before `entry`{.variable} in the
    [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-8},
    then jump to the step labeled *create*.

5.  Let `entry`{.variable} be the entry one earlier than
    `entry`{.variable} in the [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-9}.

6.  If `entry`{.variable} is neither a
    [marker](#concept-parser-marker){#the-list-of-active-formatting-elements:concept-parser-marker-5}
    nor an element that is also in the [stack of open
    elements](#stack-of-open-elements){#the-list-of-active-formatting-elements:stack-of-open-elements-2},
    go to the step labeled *rewind*.

7.  *Advance*: Let `entry`{.variable} be the element one later than
    `entry`{.variable} in the [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-10}.

8.  *Create*: [Insert an HTML
    element](#insert-an-html-element){#the-list-of-active-formatting-elements:insert-an-html-element}
    for the token for which the element `entry`{.variable} was created,
    to obtain `new element`{.variable}.

9.  Replace the entry for `entry`{.variable} in the list with an entry
    for `new element`{.variable}.

10. If the entry for `new element`{.variable} in the [list of active
    formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-11}
    is not the last entry in the list, return to the step labeled
    *advance*.

This has the effect of reopening all the formatting elements that were
opened in the current body, cell, or caption (whichever is youngest)
that haven\'t been explicitly closed.

The way this specification is written, the [list of active formatting
elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-12}
always consists of elements in chronological order with the least
recently added element first and the most recently added element last
(except for while steps 7 to 10 of the above algorithm are being
executed, of course).

When the steps below require the UA to [clear the list of active
formatting elements up to the last
marker]{#clear-the-list-of-active-formatting-elements-up-to-the-last-marker
.dfn}, the UA must perform the following steps:

1.  Let `entry`{.variable} be the last (most recently added) entry in
    the [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-13}.

2.  Remove `entry`{.variable} from the [list of active formatting
    elements](#list-of-active-formatting-elements){#the-list-of-active-formatting-elements:list-of-active-formatting-elements-14}.

3.  If `entry`{.variable} was a
    [marker](#concept-parser-marker){#the-list-of-active-formatting-elements:concept-parser-marker-6},
    then stop the algorithm at this point. The list has been cleared up
    to the last
    [marker](#concept-parser-marker){#the-list-of-active-formatting-elements:concept-parser-marker-7}.

4.  Go to step 1.

##### [13.2.4.4]{.secno} The element pointers[](#the-element-pointers){.self-link}

Initially, the [`head` element pointer]{#head-element-pointer .dfn} and
the [`form` element pointer]{#form-element-pointer .dfn} are both null.

Once a
[`head`{#the-element-pointers:the-head-element}](semantics.html#the-head-element)
element has been parsed (whether implicitly or explicitly) the [`head`
element
pointer](#head-element-pointer){#the-element-pointers:head-element-pointer}
gets set to point to this node.

The [`form` element
pointer](#form-element-pointer){#the-element-pointers:form-element-pointer}
points to the last
[`form`{#the-element-pointers:the-form-element}](forms.html#the-form-element)
element that was opened and whose end tag has not yet been seen. It is
used to make form controls associate with forms in the face of
dramatically bad markup, for historical reasons. It is ignored inside
[`template`{#the-element-pointers:the-template-element}](scripting.html#the-template-element)
elements.

##### [13.2.4.5]{.secno} Other parsing state flags[](#other-parsing-state-flags){.self-link}

The [scripting flag]{#scripting-flag .dfn} is set to \"enabled\" if
[scripting was
enabled](webappapis.html#concept-n-script){#other-parsing-state-flags:concept-n-script}
for the
[`Document`{#other-parsing-state-flags:document}](dom.html#document)
with which the parser is associated when the parser was created, and
\"disabled\" otherwise.

The [scripting
flag](#scripting-flag){#other-parsing-state-flags:scripting-flag} can be
enabled even when the parser was created as part of the [HTML fragment
parsing
algorithm](#html-fragment-parsing-algorithm){#other-parsing-state-flags:html-fragment-parsing-algorithm},
even though
[`script`{#other-parsing-state-flags:the-script-element}](scripting.html#the-script-element)
elements don\'t execute in that case.

The [frameset-ok flag]{#frameset-ok-flag .dfn} is set to \"ok\" when the
parser is created. It is set to \"not ok\" after certain tokens are
seen.

#### [13.2.5]{.secno} [Tokenization]{.dfn}[](#tokenization){.self-link}

Implementations must act as if they used the following state machine to
tokenize HTML. The state machine must start in the [data
state](#data-state){#tokenization:data-state}. Most states consume a
single character, which may have various side-effects, and either
switches the state machine to a new state to
[reconsume](#reconsume){#tokenization:reconsume} the [current input
character](#current-input-character){#tokenization:current-input-character},
or switches it to a new state to consume the [next
character](#next-input-character){#tokenization:next-input-character},
or stays in the same state to consume the next character. Some states
have more complicated behavior and can consume several characters before
switching to another state. In some cases, the tokenizer state is also
changed by the tree construction stage.

When a state says to [reconsume]{#reconsume .dfn} a matched character in
a specified state, that means to switch to that state, but when it
attempts to consume the [next input
character](#next-input-character){#tokenization:next-input-character-2},
provide it with the [current input
character](#current-input-character){#tokenization:current-input-character-2}
instead.

The exact behavior of certain states depends on the [insertion
mode](#insertion-mode){#tokenization:insertion-mode} and the [stack of
open
elements](#stack-of-open-elements){#tokenization:stack-of-open-elements}.
Certain states also use a
[`temporary buffer`{.variable}]{#temporary-buffer .dfn} to track
progress, and the [character reference
state](#character-reference-state){#tokenization:character-reference-state}
uses a [`return state`{.variable}]{#return-state .dfn} to return to the
state it was invoked from.

The output of the tokenization step is a series of zero or more of the
following tokens: DOCTYPE, start tag, end tag, comment, character,
end-of-file. DOCTYPE tokens have a name, a public identifier, a system
identifier, and a [*force-quirks flag*]{#force-quirks-flag .dfn}. When a
DOCTYPE token is created, its name, public identifier, and system
identifier must be marked as missing (which is a distinct state from the
empty string), and the *[force-quirks flag](#force-quirks-flag)* must be
set to *off* (its other state is *on*). Start and end tag tokens have a
tag name, a [self-closing flag]{#self-closing-flag .dfn}, and a list of
attributes, each of which has a name and a value. When a start or end
tag token is created, its *[self-closing flag](#self-closing-flag)* must
be unset (its other state is that it be set), and its attributes list
must be empty. Comment and character tokens have data.

When a token is emitted, it must immediately be handled by the [tree
construction](#tree-construction){#tokenization:tree-construction}
stage. The tree construction stage can affect the state of the
tokenization stage, and can insert additional characters into the
stream. (For example, the
[`script`{#tokenization:the-script-element}](scripting.html#the-script-element)
element can result in scripts executing and using the [dynamic markup
insertion](dynamic-markup-insertion.html#dynamic-markup-insertion){#tokenization:dynamic-markup-insertion}
APIs to insert characters into the stream being tokenized.)

Creating a token and emitting it are distinct actions. It is possible
for a token to be created but implicitly abandoned (never emitted), e.g.
if the file ends unexpectedly while processing the characters that are
being parsed into a start tag token.

When a start tag token is emitted with its *[self-closing
flag](#self-closing-flag)* set, if the flag is not
[acknowledged]{#acknowledge-self-closing-flag .dfn} when it is processed
by the tree construction stage, that is a
[non-void-html-element-start-tag-with-trailing-solidus](#parse-error-non-void-html-element-start-tag-with-trailing-solidus){#tokenization:parse-error-non-void-html-element-start-tag-with-trailing-solidus}
[parse error](#parse-errors){#tokenization:parse-errors}.

When an end tag token is emitted with attributes, that is an
[end-tag-with-attributes](#parse-error-end-tag-with-attributes){#tokenization:parse-error-end-tag-with-attributes}
[parse error](#parse-errors){#tokenization:parse-errors-2}.

When an end tag token is emitted with its *[self-closing
flag](#self-closing-flag)* set, that is an
[end-tag-with-trailing-solidus](#parse-error-end-tag-with-trailing-solidus){#tokenization:parse-error-end-tag-with-trailing-solidus}
[parse error](#parse-errors){#tokenization:parse-errors-3}.

An [appropriate end tag token]{#appropriate-end-tag-token .dfn} is an
end tag token whose tag name matches the tag name of the last start tag
to have been emitted from this tokenizer, if any. If no start tag has
been emitted from this tokenizer, then no end tag token is appropriate.

A [character
reference](syntax.html#syntax-charref){#tokenization:syntax-charref} is
said to be [consumed as part of an attribute]{#charref-in-attribute
.dfn} if the [`return state`{#tokenization:return-state
.variable}](#return-state) is either [attribute value (double-quoted)
state](#attribute-value-(double-quoted)-state){#tokenization:attribute-value-(double-quoted)-state},
[attribute value (single-quoted)
state](#attribute-value-(single-quoted)-state){#tokenization:attribute-value-(single-quoted)-state},
or [attribute value (unquoted)
state](#attribute-value-(unquoted)-state){#tokenization:attribute-value-(unquoted)-state}.

When a state says to [flush code points consumed as a character
reference]{#flush-code-points-consumed-as-a-character-reference .dfn},
it means that for each [code
point](https://infra.spec.whatwg.org/#code-point){#tokenization:code-point
x-internal="code-point"} in the
[`temporary buffer`{#tokenization:temporary-buffer
.variable}](#temporary-buffer) (in the order they were added to the
buffer) user agent must append the code point from the buffer to the
current attribute\'s value if the character reference was [consumed as
part of an
attribute](#charref-in-attribute){#tokenization:charref-in-attribute},
or emit the code point as a character token otherwise.

Before each step of the tokenizer, the user agent must first check the
[parser pause
flag](#parser-pause-flag){#tokenization:parser-pause-flag}. If it is
true, then the tokenizer must abort the processing of any nested
invocations of the tokenizer, yielding control back to the caller.

The tokenizer state machine consists of the states defined in the
following subsections.

##### [13.2.5.1]{.secno} [Data state]{.dfn}[](#data-state){.self-link}

Consume the [next input
character](#next-input-character){#data-state:next-input-character}:

U+0026 AMPERSAND (&)
:   Set the [`return state`{#data-state:return-state
    .variable}](#return-state) to the [data
    state](#data-state){#data-state:data-state}. Switch to the
    [character reference
    state](#character-reference-state){#data-state:character-reference-state}.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [tag open
    state](#tag-open-state){#data-state:tag-open-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#data-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#data-state:parse-errors}. Emit the
    [current input
    character](#current-input-character){#data-state:current-input-character}
    as a character token.

EOF
:   Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#data-state:current-input-character-2}
    as a character token.

##### [13.2.5.2]{.secno} [RCDATA state]{.dfn}[](#rcdata-state){.self-link}

Consume the [next input
character](#next-input-character){#rcdata-state:next-input-character}:

U+0026 AMPERSAND (&)
:   Set the [`return state`{#rcdata-state:return-state
    .variable}](#return-state) to the [RCDATA
    state](#rcdata-state){#rcdata-state:rcdata-state}. Switch to the
    [character reference
    state](#character-reference-state){#rcdata-state:character-reference-state}.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [RCDATA less-than sign
    state](#rcdata-less-than-sign-state){#rcdata-state:rcdata-less-than-sign-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#rcdata-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#rcdata-state:parse-errors}. Emit a
    U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#rcdata-state:current-input-character}
    as a character token.

##### [13.2.5.3]{.secno} [RAWTEXT state]{.dfn}[](#rawtext-state){.self-link}

Consume the [next input
character](#next-input-character){#rawtext-state:next-input-character}:

U+003C LESS-THAN SIGN (\<)
:   Switch to the [RAWTEXT less-than sign
    state](#rawtext-less-than-sign-state){#rawtext-state:rawtext-less-than-sign-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#rawtext-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#rawtext-state:parse-errors}. Emit a
    U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#rawtext-state:current-input-character}
    as a character token.

##### [13.2.5.4]{.secno} [Script data state]{.dfn}[](#script-data-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-state:next-input-character}:

U+003C LESS-THAN SIGN (\<)
:   Switch to the [script data less-than sign
    state](#script-data-less-than-sign-state){#script-data-state:script-data-less-than-sign-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#script-data-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#script-data-state:parse-errors}. Emit
    a U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#script-data-state:current-input-character}
    as a character token.

##### [13.2.5.5]{.secno} [PLAINTEXT state]{.dfn}[](#plaintext-state){.self-link}

Consume the [next input
character](#next-input-character){#plaintext-state:next-input-character}:

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#plaintext-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#plaintext-state:parse-errors}. Emit a
    U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#plaintext-state:current-input-character}
    as a character token.

##### [13.2.5.6]{.secno} [Tag open state]{.dfn}[](#tag-open-state){.self-link}

Consume the [next input
character](#next-input-character){#tag-open-state:next-input-character}:

U+0021 EXCLAMATION MARK (!)
:   Switch to the [markup declaration open
    state](#markup-declaration-open-state){#tag-open-state:markup-declaration-open-state}.

U+002F SOLIDUS (/)
:   Switch to the [end tag open
    state](#end-tag-open-state){#tag-open-state:end-tag-open-state}.

[ASCII alpha](https://infra.spec.whatwg.org/#ascii-alpha){#tag-open-state:ascii-letters x-internal="ascii-letters"}
:   Create a new start tag token, set its tag name to the empty string.
    [Reconsume](#reconsume){#tag-open-state:reconsume} in the [tag name
    state](#tag-name-state){#tag-open-state:tag-name-state}.

U+003F QUESTION MARK (?)
:   This is an
    [unexpected-question-mark-instead-of-tag-name](#parse-error-unexpected-question-mark-instead-of-tag-name){#tag-open-state:parse-error-unexpected-question-mark-instead-of-tag-name}
    [parse error](#parse-errors){#tag-open-state:parse-errors}. Create a
    comment token whose data is the empty string.
    [Reconsume](#reconsume){#tag-open-state:reconsume-2} in the [bogus
    comment
    state](#bogus-comment-state){#tag-open-state:bogus-comment-state}.

EOF
:   This is an
    [eof-before-tag-name](#parse-error-eof-before-tag-name){#tag-open-state:parse-error-eof-before-tag-name}
    [parse error](#parse-errors){#tag-open-state:parse-errors-2}. Emit a
    U+003C LESS-THAN SIGN character token and an end-of-file token.

Anything else
:   This is an
    [invalid-first-character-of-tag-name](#parse-error-invalid-first-character-of-tag-name){#tag-open-state:parse-error-invalid-first-character-of-tag-name}
    [parse error](#parse-errors){#tag-open-state:parse-errors-3}. Emit a
    U+003C LESS-THAN SIGN character token.
    [Reconsume](#reconsume){#tag-open-state:reconsume-3} in the [data
    state](#data-state){#tag-open-state:data-state}.

##### [13.2.5.7]{.secno} [End tag open state]{.dfn}[](#end-tag-open-state){.self-link}

Consume the [next input
character](#next-input-character){#end-tag-open-state:next-input-character}:

[ASCII alpha](https://infra.spec.whatwg.org/#ascii-alpha){#end-tag-open-state:ascii-letters x-internal="ascii-letters"}
:   Create a new end tag token, set its tag name to the empty string.
    [Reconsume](#reconsume){#end-tag-open-state:reconsume} in the [tag
    name state](#tag-name-state){#end-tag-open-state:tag-name-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is a
    [missing-end-tag-name](#parse-error-missing-end-tag-name){#end-tag-open-state:parse-error-missing-end-tag-name}
    [parse error](#parse-errors){#end-tag-open-state:parse-errors}.
    Switch to the [data
    state](#data-state){#end-tag-open-state:data-state}.

EOF
:   This is an
    [eof-before-tag-name](#parse-error-eof-before-tag-name){#end-tag-open-state:parse-error-eof-before-tag-name}
    [parse error](#parse-errors){#end-tag-open-state:parse-errors-2}.
    Emit a U+003C LESS-THAN SIGN character token, a U+002F SOLIDUS
    character token and an end-of-file token.

Anything else
:   This is an
    [invalid-first-character-of-tag-name](#parse-error-invalid-first-character-of-tag-name){#end-tag-open-state:parse-error-invalid-first-character-of-tag-name}
    [parse error](#parse-errors){#end-tag-open-state:parse-errors-3}.
    Create a comment token whose data is the empty string.
    [Reconsume](#reconsume){#end-tag-open-state:reconsume-2} in the
    [bogus comment
    state](#bogus-comment-state){#end-tag-open-state:bogus-comment-state}.

##### [13.2.5.8]{.secno} [Tag name state]{.dfn}[](#tag-name-state){.self-link}

Consume the [next input
character](#next-input-character){#tag-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [before attribute name
    state](#before-attribute-name-state){#tag-name-state:before-attribute-name-state}.

U+002F SOLIDUS (/)
:   Switch to the [self-closing start tag
    state](#self-closing-start-tag-state){#tag-name-state:self-closing-start-tag-state}.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data state](#data-state){#tag-name-state:data-state}.
    Emit the current tag token.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#tag-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#tag-name-state:current-input-character}
    (add 0x0020 to the character\'s code point) to the current tag
    token\'s tag name.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#tag-name-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#tag-name-state:parse-errors}. Append a
    U+FFFD REPLACEMENT CHARACTER character to the current tag token\'s
    tag name.

EOF
:   This is an
    [eof-in-tag](#parse-error-eof-in-tag){#tag-name-state:parse-error-eof-in-tag}
    [parse error](#parse-errors){#tag-name-state:parse-errors-2}. Emit
    an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#tag-name-state:current-input-character-2}
    to the current tag token\'s tag name.

##### [13.2.5.9]{.secno} [RCDATA less-than sign state]{.dfn}[](#rcdata-less-than-sign-state){.self-link}

Consume the [next input
character](#next-input-character){#rcdata-less-than-sign-state:next-input-character}:

U+002F SOLIDUS (/)
:   Set the
    [`temporary buffer`{#rcdata-less-than-sign-state:temporary-buffer
    .variable}](#temporary-buffer) to the empty string. Switch to the
    [RCDATA end tag open
    state](#rcdata-end-tag-open-state){#rcdata-less-than-sign-state:rcdata-end-tag-open-state}.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token.
    [Reconsume](#reconsume){#rcdata-less-than-sign-state:reconsume} in
    the [RCDATA
    state](#rcdata-state){#rcdata-less-than-sign-state:rcdata-state}.

##### [13.2.5.10]{.secno} [RCDATA end tag open state]{.dfn}[](#rcdata-end-tag-open-state){.self-link}

Consume the [next input
character](#next-input-character){#rcdata-end-tag-open-state:next-input-character}:

[ASCII alpha](https://infra.spec.whatwg.org/#ascii-alpha){#rcdata-end-tag-open-state:ascii-letters x-internal="ascii-letters"}
:   Create a new end tag token, set its tag name to the empty string.
    [Reconsume](#reconsume){#rcdata-end-tag-open-state:reconsume} in the
    [RCDATA end tag name
    state](#rcdata-end-tag-name-state){#rcdata-end-tag-open-state:rcdata-end-tag-name-state}.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token and a U+002F SOLIDUS
    character token.
    [Reconsume](#reconsume){#rcdata-end-tag-open-state:reconsume-2} in
    the [RCDATA
    state](#rcdata-state){#rcdata-end-tag-open-state:rcdata-state}.

##### [13.2.5.11]{.secno} [RCDATA end tag name state]{.dfn}[](#rcdata-end-tag-name-state){.self-link}

Consume the [next input
character](#next-input-character){#rcdata-end-tag-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#rcdata-end-tag-name-state:appropriate-end-tag-token},
    then switch to the [before attribute name
    state](#before-attribute-name-state){#rcdata-end-tag-name-state:before-attribute-name-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+002F SOLIDUS (/)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#rcdata-end-tag-name-state:appropriate-end-tag-token-2},
    then switch to the [self-closing start tag
    state](#self-closing-start-tag-state){#rcdata-end-tag-name-state:self-closing-start-tag-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+003E GREATER-THAN SIGN (\>)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#rcdata-end-tag-name-state:appropriate-end-tag-token-3},
    then switch to the [data
    state](#data-state){#rcdata-end-tag-name-state:data-state} and emit
    the current tag token. Otherwise, treat it as per the \"anything
    else\" entry below.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#rcdata-end-tag-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#rcdata-end-tag-name-state:current-input-character}
    (add 0x0020 to the character\'s code point) to the current tag
    token\'s tag name. Append the [current input
    character](#current-input-character){#rcdata-end-tag-name-state:current-input-character-2}
    to the
    [`temporary buffer`{#rcdata-end-tag-name-state:temporary-buffer
    .variable}](#temporary-buffer).

[ASCII lower alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#rcdata-end-tag-name-state:lowercase-ascii-letters x-internal="lowercase-ascii-letters"}
:   Append the [current input
    character](#current-input-character){#rcdata-end-tag-name-state:current-input-character-3}
    to the current tag token\'s tag name. Append the [current input
    character](#current-input-character){#rcdata-end-tag-name-state:current-input-character-4}
    to the
    [`temporary buffer`{#rcdata-end-tag-name-state:temporary-buffer-2
    .variable}](#temporary-buffer).

Anything else
:   Emit a U+003C LESS-THAN SIGN character token, a U+002F SOLIDUS
    character token, and a character token for each of the characters in
    the
    [`temporary buffer`{#rcdata-end-tag-name-state:temporary-buffer-3
    .variable}](#temporary-buffer) (in the order they were added to the
    buffer).
    [Reconsume](#reconsume){#rcdata-end-tag-name-state:reconsume} in the
    [RCDATA
    state](#rcdata-state){#rcdata-end-tag-name-state:rcdata-state}.

##### [13.2.5.12]{.secno} [RAWTEXT less-than sign state]{.dfn}[](#rawtext-less-than-sign-state){.self-link}

Consume the [next input
character](#next-input-character){#rawtext-less-than-sign-state:next-input-character}:

U+002F SOLIDUS (/)
:   Set the
    [`temporary buffer`{#rawtext-less-than-sign-state:temporary-buffer
    .variable}](#temporary-buffer) to the empty string. Switch to the
    [RAWTEXT end tag open
    state](#rawtext-end-tag-open-state){#rawtext-less-than-sign-state:rawtext-end-tag-open-state}.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token.
    [Reconsume](#reconsume){#rawtext-less-than-sign-state:reconsume} in
    the [RAWTEXT
    state](#rawtext-state){#rawtext-less-than-sign-state:rawtext-state}.

##### [13.2.5.13]{.secno} [RAWTEXT end tag open state]{.dfn}[](#rawtext-end-tag-open-state){.self-link}

Consume the [next input
character](#next-input-character){#rawtext-end-tag-open-state:next-input-character}:

[ASCII alpha](https://infra.spec.whatwg.org/#ascii-alpha){#rawtext-end-tag-open-state:ascii-letters x-internal="ascii-letters"}
:   Create a new end tag token, set its tag name to the empty string.
    [Reconsume](#reconsume){#rawtext-end-tag-open-state:reconsume} in
    the [RAWTEXT end tag name
    state](#rawtext-end-tag-name-state){#rawtext-end-tag-open-state:rawtext-end-tag-name-state}.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token and a U+002F SOLIDUS
    character token.
    [Reconsume](#reconsume){#rawtext-end-tag-open-state:reconsume-2} in
    the [RAWTEXT
    state](#rawtext-state){#rawtext-end-tag-open-state:rawtext-state}.

##### [13.2.5.14]{.secno} [RAWTEXT end tag name state]{.dfn}[](#rawtext-end-tag-name-state){.self-link}

Consume the [next input
character](#next-input-character){#rawtext-end-tag-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#rawtext-end-tag-name-state:appropriate-end-tag-token},
    then switch to the [before attribute name
    state](#before-attribute-name-state){#rawtext-end-tag-name-state:before-attribute-name-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+002F SOLIDUS (/)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#rawtext-end-tag-name-state:appropriate-end-tag-token-2},
    then switch to the [self-closing start tag
    state](#self-closing-start-tag-state){#rawtext-end-tag-name-state:self-closing-start-tag-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+003E GREATER-THAN SIGN (\>)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#rawtext-end-tag-name-state:appropriate-end-tag-token-3},
    then switch to the [data
    state](#data-state){#rawtext-end-tag-name-state:data-state} and emit
    the current tag token. Otherwise, treat it as per the \"anything
    else\" entry below.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#rawtext-end-tag-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#rawtext-end-tag-name-state:current-input-character}
    (add 0x0020 to the character\'s code point) to the current tag
    token\'s tag name. Append the [current input
    character](#current-input-character){#rawtext-end-tag-name-state:current-input-character-2}
    to the
    [`temporary buffer`{#rawtext-end-tag-name-state:temporary-buffer
    .variable}](#temporary-buffer).

[ASCII lower alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#rawtext-end-tag-name-state:lowercase-ascii-letters x-internal="lowercase-ascii-letters"}
:   Append the [current input
    character](#current-input-character){#rawtext-end-tag-name-state:current-input-character-3}
    to the current tag token\'s tag name. Append the [current input
    character](#current-input-character){#rawtext-end-tag-name-state:current-input-character-4}
    to the
    [`temporary buffer`{#rawtext-end-tag-name-state:temporary-buffer-2
    .variable}](#temporary-buffer).

Anything else
:   Emit a U+003C LESS-THAN SIGN character token, a U+002F SOLIDUS
    character token, and a character token for each of the characters in
    the
    [`temporary buffer`{#rawtext-end-tag-name-state:temporary-buffer-3
    .variable}](#temporary-buffer) (in the order they were added to the
    buffer).
    [Reconsume](#reconsume){#rawtext-end-tag-name-state:reconsume} in
    the [RAWTEXT
    state](#rawtext-state){#rawtext-end-tag-name-state:rawtext-state}.

##### [13.2.5.15]{.secno} [Script data less-than sign state]{.dfn}[](#script-data-less-than-sign-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-less-than-sign-state:next-input-character}:

U+002F SOLIDUS (/)
:   Set the
    [`temporary buffer`{#script-data-less-than-sign-state:temporary-buffer
    .variable}](#temporary-buffer) to the empty string. Switch to the
    [script data end tag open
    state](#script-data-end-tag-open-state){#script-data-less-than-sign-state:script-data-end-tag-open-state}.

U+0021 EXCLAMATION MARK (!)
:   Switch to the [script data escape start
    state](#script-data-escape-start-state){#script-data-less-than-sign-state:script-data-escape-start-state}.
    Emit a U+003C LESS-THAN SIGN character token and a U+0021
    EXCLAMATION MARK character token.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token.
    [Reconsume](#reconsume){#script-data-less-than-sign-state:reconsume}
    in the [script data
    state](#script-data-state){#script-data-less-than-sign-state:script-data-state}.

##### [13.2.5.16]{.secno} [Script data end tag open state]{.dfn}[](#script-data-end-tag-open-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-end-tag-open-state:next-input-character}:

[ASCII alpha](https://infra.spec.whatwg.org/#ascii-alpha){#script-data-end-tag-open-state:ascii-letters x-internal="ascii-letters"}
:   Create a new end tag token, set its tag name to the empty string.
    [Reconsume](#reconsume){#script-data-end-tag-open-state:reconsume}
    in the [script data end tag name
    state](#script-data-end-tag-name-state){#script-data-end-tag-open-state:script-data-end-tag-name-state}.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token and a U+002F SOLIDUS
    character token.
    [Reconsume](#reconsume){#script-data-end-tag-open-state:reconsume-2}
    in the [script data
    state](#script-data-state){#script-data-end-tag-open-state:script-data-state}.

##### [13.2.5.17]{.secno} [Script data end tag name state]{.dfn}[](#script-data-end-tag-name-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-end-tag-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#script-data-end-tag-name-state:appropriate-end-tag-token},
    then switch to the [before attribute name
    state](#before-attribute-name-state){#script-data-end-tag-name-state:before-attribute-name-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+002F SOLIDUS (/)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#script-data-end-tag-name-state:appropriate-end-tag-token-2},
    then switch to the [self-closing start tag
    state](#self-closing-start-tag-state){#script-data-end-tag-name-state:self-closing-start-tag-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+003E GREATER-THAN SIGN (\>)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#script-data-end-tag-name-state:appropriate-end-tag-token-3},
    then switch to the [data
    state](#data-state){#script-data-end-tag-name-state:data-state} and
    emit the current tag token. Otherwise, treat it as per the
    \"anything else\" entry below.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#script-data-end-tag-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#script-data-end-tag-name-state:current-input-character}
    (add 0x0020 to the character\'s code point) to the current tag
    token\'s tag name. Append the [current input
    character](#current-input-character){#script-data-end-tag-name-state:current-input-character-2}
    to the
    [`temporary buffer`{#script-data-end-tag-name-state:temporary-buffer
    .variable}](#temporary-buffer).

[ASCII lower alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#script-data-end-tag-name-state:lowercase-ascii-letters x-internal="lowercase-ascii-letters"}
:   Append the [current input
    character](#current-input-character){#script-data-end-tag-name-state:current-input-character-3}
    to the current tag token\'s tag name. Append the [current input
    character](#current-input-character){#script-data-end-tag-name-state:current-input-character-4}
    to the
    [`temporary buffer`{#script-data-end-tag-name-state:temporary-buffer-2
    .variable}](#temporary-buffer).

Anything else
:   Emit a U+003C LESS-THAN SIGN character token, a U+002F SOLIDUS
    character token, and a character token for each of the characters in
    the
    [`temporary buffer`{#script-data-end-tag-name-state:temporary-buffer-3
    .variable}](#temporary-buffer) (in the order they were added to the
    buffer).
    [Reconsume](#reconsume){#script-data-end-tag-name-state:reconsume}
    in the [script data
    state](#script-data-state){#script-data-end-tag-name-state:script-data-state}.

##### [13.2.5.18]{.secno} [Script data escape start state]{.dfn}[](#script-data-escape-start-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escape-start-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [script data escape start dash
    state](#script-data-escape-start-dash-state){#script-data-escape-start-state:script-data-escape-start-dash-state}.
    Emit a U+002D HYPHEN-MINUS character token.

Anything else
:   [Reconsume](#reconsume){#script-data-escape-start-state:reconsume}
    in the [script data
    state](#script-data-state){#script-data-escape-start-state:script-data-state}.

##### [13.2.5.19]{.secno} [Script data escape start dash state]{.dfn}[](#script-data-escape-start-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escape-start-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [script data escaped dash dash
    state](#script-data-escaped-dash-dash-state){#script-data-escape-start-dash-state:script-data-escaped-dash-dash-state}.
    Emit a U+002D HYPHEN-MINUS character token.

Anything else
:   [Reconsume](#reconsume){#script-data-escape-start-dash-state:reconsume}
    in the [script data
    state](#script-data-state){#script-data-escape-start-dash-state:script-data-state}.

##### [13.2.5.20]{.secno} [Script data escaped state]{.dfn}[](#script-data-escaped-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escaped-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [script data escaped dash
    state](#script-data-escaped-dash-state){#script-data-escaped-state:script-data-escaped-dash-state}.
    Emit a U+002D HYPHEN-MINUS character token.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [script data escaped less-than sign
    state](#script-data-escaped-less-than-sign-state){#script-data-escaped-state:script-data-escaped-less-than-sign-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#script-data-escaped-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#script-data-escaped-state:parse-errors}. Emit
    a U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   This is an
    [eof-in-script-html-comment-like-text](#parse-error-eof-in-script-html-comment-like-text){#script-data-escaped-state:parse-error-eof-in-script-html-comment-like-text}
    [parse
    error](#parse-errors){#script-data-escaped-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#script-data-escaped-state:current-input-character}
    as a character token.

##### [13.2.5.21]{.secno} [Script data escaped dash state]{.dfn}[](#script-data-escaped-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escaped-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [script data escaped dash dash
    state](#script-data-escaped-dash-dash-state){#script-data-escaped-dash-state:script-data-escaped-dash-dash-state}.
    Emit a U+002D HYPHEN-MINUS character token.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [script data escaped less-than sign
    state](#script-data-escaped-less-than-sign-state){#script-data-escaped-dash-state:script-data-escaped-less-than-sign-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#script-data-escaped-dash-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#script-data-escaped-dash-state:parse-errors}.
    Switch to the [script data escaped
    state](#script-data-escaped-state){#script-data-escaped-dash-state:script-data-escaped-state}.
    Emit a U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   This is an
    [eof-in-script-html-comment-like-text](#parse-error-eof-in-script-html-comment-like-text){#script-data-escaped-dash-state:parse-error-eof-in-script-html-comment-like-text}
    [parse
    error](#parse-errors){#script-data-escaped-dash-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Switch to the [script data escaped
    state](#script-data-escaped-state){#script-data-escaped-dash-state:script-data-escaped-state-2}.
    Emit the [current input
    character](#current-input-character){#script-data-escaped-dash-state:current-input-character}
    as a character token.

##### [13.2.5.22]{.secno} [Script data escaped dash dash state]{.dfn}[](#script-data-escaped-dash-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escaped-dash-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Emit a U+002D HYPHEN-MINUS character token.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [script data escaped less-than sign
    state](#script-data-escaped-less-than-sign-state){#script-data-escaped-dash-dash-state:script-data-escaped-less-than-sign-state}.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [script data
    state](#script-data-state){#script-data-escaped-dash-dash-state:script-data-state}.
    Emit a U+003E GREATER-THAN SIGN character token.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#script-data-escaped-dash-dash-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#script-data-escaped-dash-dash-state:parse-errors}.
    Switch to the [script data escaped
    state](#script-data-escaped-state){#script-data-escaped-dash-dash-state:script-data-escaped-state}.
    Emit a U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   This is an
    [eof-in-script-html-comment-like-text](#parse-error-eof-in-script-html-comment-like-text){#script-data-escaped-dash-dash-state:parse-error-eof-in-script-html-comment-like-text}
    [parse
    error](#parse-errors){#script-data-escaped-dash-dash-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Switch to the [script data escaped
    state](#script-data-escaped-state){#script-data-escaped-dash-dash-state:script-data-escaped-state-2}.
    Emit the [current input
    character](#current-input-character){#script-data-escaped-dash-dash-state:current-input-character}
    as a character token.

##### [13.2.5.23]{.secno} [Script data escaped less-than sign state]{.dfn}[](#script-data-escaped-less-than-sign-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escaped-less-than-sign-state:next-input-character}:

U+002F SOLIDUS (/)
:   Set the
    [`temporary buffer`{#script-data-escaped-less-than-sign-state:temporary-buffer
    .variable}](#temporary-buffer) to the empty string. Switch to the
    [script data escaped end tag open
    state](#script-data-escaped-end-tag-open-state){#script-data-escaped-less-than-sign-state:script-data-escaped-end-tag-open-state}.

[ASCII alpha](https://infra.spec.whatwg.org/#ascii-alpha){#script-data-escaped-less-than-sign-state:ascii-letters x-internal="ascii-letters"}
:   Set the
    [`temporary buffer`{#script-data-escaped-less-than-sign-state:temporary-buffer-2
    .variable}](#temporary-buffer) to the empty string. Emit a U+003C
    LESS-THAN SIGN character token.
    [Reconsume](#reconsume){#script-data-escaped-less-than-sign-state:reconsume}
    in the [script data double escape start
    state](#script-data-double-escape-start-state){#script-data-escaped-less-than-sign-state:script-data-double-escape-start-state}.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token.
    [Reconsume](#reconsume){#script-data-escaped-less-than-sign-state:reconsume-2}
    in the [script data escaped
    state](#script-data-escaped-state){#script-data-escaped-less-than-sign-state:script-data-escaped-state}.

##### [13.2.5.24]{.secno} [Script data escaped end tag open state]{.dfn}[](#script-data-escaped-end-tag-open-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escaped-end-tag-open-state:next-input-character}:

[ASCII alpha](https://infra.spec.whatwg.org/#ascii-alpha){#script-data-escaped-end-tag-open-state:ascii-letters x-internal="ascii-letters"}
:   Create a new end tag token, set its tag name to the empty string.
    [Reconsume](#reconsume){#script-data-escaped-end-tag-open-state:reconsume}
    in the [script data escaped end tag name
    state](#script-data-escaped-end-tag-name-state){#script-data-escaped-end-tag-open-state:script-data-escaped-end-tag-name-state}.

Anything else
:   Emit a U+003C LESS-THAN SIGN character token and a U+002F SOLIDUS
    character token.
    [Reconsume](#reconsume){#script-data-escaped-end-tag-open-state:reconsume-2}
    in the [script data escaped
    state](#script-data-escaped-state){#script-data-escaped-end-tag-open-state:script-data-escaped-state}.

##### [13.2.5.25]{.secno} [Script data escaped end tag name state]{.dfn}[](#script-data-escaped-end-tag-name-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-escaped-end-tag-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#script-data-escaped-end-tag-name-state:appropriate-end-tag-token},
    then switch to the [before attribute name
    state](#before-attribute-name-state){#script-data-escaped-end-tag-name-state:before-attribute-name-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+002F SOLIDUS (/)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#script-data-escaped-end-tag-name-state:appropriate-end-tag-token-2},
    then switch to the [self-closing start tag
    state](#self-closing-start-tag-state){#script-data-escaped-end-tag-name-state:self-closing-start-tag-state}.
    Otherwise, treat it as per the \"anything else\" entry below.

U+003E GREATER-THAN SIGN (\>)
:   If the current end tag token is an [appropriate end tag
    token](#appropriate-end-tag-token){#script-data-escaped-end-tag-name-state:appropriate-end-tag-token-3},
    then switch to the [data
    state](#data-state){#script-data-escaped-end-tag-name-state:data-state}
    and emit the current tag token. Otherwise, treat it as per the
    \"anything else\" entry below.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#script-data-escaped-end-tag-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#script-data-escaped-end-tag-name-state:current-input-character}
    (add 0x0020 to the character\'s code point) to the current tag
    token\'s tag name. Append the [current input
    character](#current-input-character){#script-data-escaped-end-tag-name-state:current-input-character-2}
    to the
    [`temporary buffer`{#script-data-escaped-end-tag-name-state:temporary-buffer
    .variable}](#temporary-buffer).

[ASCII lower alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#script-data-escaped-end-tag-name-state:lowercase-ascii-letters x-internal="lowercase-ascii-letters"}
:   Append the [current input
    character](#current-input-character){#script-data-escaped-end-tag-name-state:current-input-character-3}
    to the current tag token\'s tag name. Append the [current input
    character](#current-input-character){#script-data-escaped-end-tag-name-state:current-input-character-4}
    to the
    [`temporary buffer`{#script-data-escaped-end-tag-name-state:temporary-buffer-2
    .variable}](#temporary-buffer).

Anything else
:   Emit a U+003C LESS-THAN SIGN character token, a U+002F SOLIDUS
    character token, and a character token for each of the characters in
    the
    [`temporary buffer`{#script-data-escaped-end-tag-name-state:temporary-buffer-3
    .variable}](#temporary-buffer)` `{#script-data-escaped-end-tag-name-state:temporary-buffer-3
    .variable} (in the order they were added to the buffer).
    [Reconsume](#reconsume){#script-data-escaped-end-tag-name-state:reconsume}
    in the [script data escaped
    state](#script-data-escaped-state){#script-data-escaped-end-tag-name-state:script-data-escaped-state}.

##### [13.2.5.26]{.secno} [Script data double escape start state]{.dfn}[](#script-data-double-escape-start-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-double-escape-start-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE\
U+002F SOLIDUS (/)\
U+003E GREATER-THAN SIGN (\>)
:   If the
    [`temporary buffer`{#script-data-double-escape-start-state:temporary-buffer
    .variable}](#temporary-buffer) is the string \"`script`\", then
    switch to the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escape-start-state:script-data-double-escaped-state}.
    Otherwise, switch to the [script data escaped
    state](#script-data-escaped-state){#script-data-double-escape-start-state:script-data-escaped-state}.
    Emit the [current input
    character](#current-input-character){#script-data-double-escape-start-state:current-input-character}
    as a character token.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#script-data-double-escape-start-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#script-data-double-escape-start-state:current-input-character-2}
    (add 0x0020 to the character\'s code point) to the
    [`temporary buffer`{#script-data-double-escape-start-state:temporary-buffer-2
    .variable}](#temporary-buffer). Emit the [current input
    character](#current-input-character){#script-data-double-escape-start-state:current-input-character-3}
    as a character token.

[ASCII lower alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#script-data-double-escape-start-state:lowercase-ascii-letters x-internal="lowercase-ascii-letters"}
:   Append the [current input
    character](#current-input-character){#script-data-double-escape-start-state:current-input-character-4}
    to the
    [`temporary buffer`{#script-data-double-escape-start-state:temporary-buffer-3
    .variable}](#temporary-buffer). Emit the [current input
    character](#current-input-character){#script-data-double-escape-start-state:current-input-character-5}
    as a character token.

Anything else
:   [Reconsume](#reconsume){#script-data-double-escape-start-state:reconsume}
    in the [script data escaped
    state](#script-data-escaped-state){#script-data-double-escape-start-state:script-data-escaped-state-2}.

##### [13.2.5.27]{.secno} [Script data double escaped state]{.dfn}[](#script-data-double-escaped-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-double-escaped-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [script data double escaped dash
    state](#script-data-double-escaped-dash-state){#script-data-double-escaped-state:script-data-double-escaped-dash-state}.
    Emit a U+002D HYPHEN-MINUS character token.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [script data double escaped less-than sign
    state](#script-data-double-escaped-less-than-sign-state){#script-data-double-escaped-state:script-data-double-escaped-less-than-sign-state}.
    Emit a U+003C LESS-THAN SIGN character token.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#script-data-double-escaped-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#script-data-double-escaped-state:parse-errors}.
    Emit a U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   This is an
    [eof-in-script-html-comment-like-text](#parse-error-eof-in-script-html-comment-like-text){#script-data-double-escaped-state:parse-error-eof-in-script-html-comment-like-text}
    [parse
    error](#parse-errors){#script-data-double-escaped-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#script-data-double-escaped-state:current-input-character}
    as a character token.

##### [13.2.5.28]{.secno} [Script data double escaped dash state]{.dfn}[](#script-data-double-escaped-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-double-escaped-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [script data double escaped dash dash
    state](#script-data-double-escaped-dash-dash-state){#script-data-double-escaped-dash-state:script-data-double-escaped-dash-dash-state}.
    Emit a U+002D HYPHEN-MINUS character token.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [script data double escaped less-than sign
    state](#script-data-double-escaped-less-than-sign-state){#script-data-double-escaped-dash-state:script-data-double-escaped-less-than-sign-state}.
    Emit a U+003C LESS-THAN SIGN character token.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#script-data-double-escaped-dash-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#script-data-double-escaped-dash-state:parse-errors}.
    Switch to the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escaped-dash-state:script-data-double-escaped-state}.
    Emit a U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   This is an
    [eof-in-script-html-comment-like-text](#parse-error-eof-in-script-html-comment-like-text){#script-data-double-escaped-dash-state:parse-error-eof-in-script-html-comment-like-text}
    [parse
    error](#parse-errors){#script-data-double-escaped-dash-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Switch to the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escaped-dash-state:script-data-double-escaped-state-2}.
    Emit the [current input
    character](#current-input-character){#script-data-double-escaped-dash-state:current-input-character}
    as a character token.

##### [13.2.5.29]{.secno} [Script data double escaped dash dash state]{.dfn}[](#script-data-double-escaped-dash-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-double-escaped-dash-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Emit a U+002D HYPHEN-MINUS character token.

U+003C LESS-THAN SIGN (\<)
:   Switch to the [script data double escaped less-than sign
    state](#script-data-double-escaped-less-than-sign-state){#script-data-double-escaped-dash-dash-state:script-data-double-escaped-less-than-sign-state}.
    Emit a U+003C LESS-THAN SIGN character token.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [script data
    state](#script-data-state){#script-data-double-escaped-dash-dash-state:script-data-state}.
    Emit a U+003E GREATER-THAN SIGN character token.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#script-data-double-escaped-dash-dash-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#script-data-double-escaped-dash-dash-state:parse-errors}.
    Switch to the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escaped-dash-dash-state:script-data-double-escaped-state}.
    Emit a U+FFFD REPLACEMENT CHARACTER character token.

EOF
:   This is an
    [eof-in-script-html-comment-like-text](#parse-error-eof-in-script-html-comment-like-text){#script-data-double-escaped-dash-dash-state:parse-error-eof-in-script-html-comment-like-text}
    [parse
    error](#parse-errors){#script-data-double-escaped-dash-dash-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Switch to the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escaped-dash-dash-state:script-data-double-escaped-state-2}.
    Emit the [current input
    character](#current-input-character){#script-data-double-escaped-dash-dash-state:current-input-character}
    as a character token.

##### [13.2.5.30]{.secno} [Script data double escaped less-than sign state]{.dfn}[](#script-data-double-escaped-less-than-sign-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-double-escaped-less-than-sign-state:next-input-character}:

U+002F SOLIDUS (/)
:   Set the
    [`temporary buffer`{#script-data-double-escaped-less-than-sign-state:temporary-buffer
    .variable}](#temporary-buffer) to the empty string. Switch to the
    [script data double escape end
    state](#script-data-double-escape-end-state){#script-data-double-escaped-less-than-sign-state:script-data-double-escape-end-state}.
    Emit a U+002F SOLIDUS character token.

Anything else
:   [Reconsume](#reconsume){#script-data-double-escaped-less-than-sign-state:reconsume}
    in the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escaped-less-than-sign-state:script-data-double-escaped-state}.

##### [13.2.5.31]{.secno} [Script data double escape end state]{.dfn}[](#script-data-double-escape-end-state){.self-link}

Consume the [next input
character](#next-input-character){#script-data-double-escape-end-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE\
U+002F SOLIDUS (/)\
U+003E GREATER-THAN SIGN (\>)
:   If the
    [`temporary buffer`{#script-data-double-escape-end-state:temporary-buffer
    .variable}](#temporary-buffer) is the string \"`script`\", then
    switch to the [script data escaped
    state](#script-data-escaped-state){#script-data-double-escape-end-state:script-data-escaped-state}.
    Otherwise, switch to the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escape-end-state:script-data-double-escaped-state}.
    Emit the [current input
    character](#current-input-character){#script-data-double-escape-end-state:current-input-character}
    as a character token.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#script-data-double-escape-end-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#script-data-double-escape-end-state:current-input-character-2}
    (add 0x0020 to the character\'s code point) to the
    [`temporary buffer`{#script-data-double-escape-end-state:temporary-buffer-2
    .variable}](#temporary-buffer). Emit the [current input
    character](#current-input-character){#script-data-double-escape-end-state:current-input-character-3}
    as a character token.

[ASCII lower alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha){#script-data-double-escape-end-state:lowercase-ascii-letters x-internal="lowercase-ascii-letters"}
:   Append the [current input
    character](#current-input-character){#script-data-double-escape-end-state:current-input-character-4}
    to the
    [`temporary buffer`{#script-data-double-escape-end-state:temporary-buffer-3
    .variable}](#temporary-buffer). Emit the [current input
    character](#current-input-character){#script-data-double-escape-end-state:current-input-character-5}
    as a character token.

Anything else
:   [Reconsume](#reconsume){#script-data-double-escape-end-state:reconsume}
    in the [script data double escaped
    state](#script-data-double-escaped-state){#script-data-double-escape-end-state:script-data-double-escaped-state-2}.

##### [13.2.5.32]{.secno} [Before attribute name state]{.dfn}[](#before-attribute-name-state){.self-link}

Consume the [next input
character](#next-input-character){#before-attribute-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+002F SOLIDUS (/)\
U+003E GREATER-THAN SIGN (\>)\
EOF
:   [Reconsume](#reconsume){#before-attribute-name-state:reconsume} in
    the [after attribute name
    state](#after-attribute-name-state){#before-attribute-name-state:after-attribute-name-state}.

U+003D EQUALS SIGN (=)
:   This is an
    [unexpected-equals-sign-before-attribute-name](#parse-error-unexpected-equals-sign-before-attribute-name){#before-attribute-name-state:parse-error-unexpected-equals-sign-before-attribute-name}
    [parse
    error](#parse-errors){#before-attribute-name-state:parse-errors}.
    Start a new attribute in the current tag token. Set that
    attribute\'s name to the [current input
    character](#current-input-character){#before-attribute-name-state:current-input-character},
    and its value to the empty string. Switch to the [attribute name
    state](#attribute-name-state){#before-attribute-name-state:attribute-name-state}.

Anything else
:   Start a new attribute in the current tag token. Set that attribute
    name and value to the empty string.
    [Reconsume](#reconsume){#before-attribute-name-state:reconsume-2} in
    the [attribute name
    state](#attribute-name-state){#before-attribute-name-state:attribute-name-state-2}.

##### [13.2.5.33]{.secno} [Attribute name state]{.dfn}[](#attribute-name-state){.self-link}

Consume the [next input
character](#next-input-character){#attribute-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE\
U+002F SOLIDUS (/)\
U+003E GREATER-THAN SIGN (\>)\
EOF
:   [Reconsume](#reconsume){#attribute-name-state:reconsume} in the
    [after attribute name
    state](#after-attribute-name-state){#attribute-name-state:after-attribute-name-state}.

U+003D EQUALS SIGN (=)
:   Switch to the [before attribute value
    state](#before-attribute-value-state){#attribute-name-state:before-attribute-value-state}.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#attribute-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#attribute-name-state:current-input-character}
    (add 0x0020 to the character\'s code point) to the current
    attribute\'s name.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#attribute-name-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#attribute-name-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    attribute\'s name.

U+0022 QUOTATION MARK (\")\
U+0027 APOSTROPHE (\')\
U+003C LESS-THAN SIGN (\<)
:   This is an
    [unexpected-character-in-attribute-name](#parse-error-unexpected-character-in-attribute-name){#attribute-name-state:parse-error-unexpected-character-in-attribute-name}
    [parse error](#parse-errors){#attribute-name-state:parse-errors-2}.
    Treat it as per the \"anything else\" entry below.

Anything else
:   Append the [current input
    character](#current-input-character){#attribute-name-state:current-input-character-2}
    to the current attribute\'s name.

When the user agent leaves the attribute name state (and before emitting
the tag token, if appropriate), the complete attribute\'s name must be
compared to the other attributes on the same token; if there is already
an attribute on the token with the exact same name, then this is a
[duplicate-attribute](#parse-error-duplicate-attribute){#attribute-name-state:parse-error-duplicate-attribute}
[parse error](#parse-errors){#attribute-name-state:parse-errors-3} and
the new attribute must be removed from the token.

If an attribute is so removed from a token, it, and the value that gets
associated with it, if any, are never subsequently used by the parser,
and are therefore effectively discarded. Removing the attribute in this
way does not change its status as the \"current attribute\" for the
purposes of the tokenizer, however.

##### [13.2.5.34]{.secno} [After attribute name state]{.dfn}[](#after-attribute-name-state){.self-link}

Consume the [next input
character](#next-input-character){#after-attribute-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+002F SOLIDUS (/)
:   Switch to the [self-closing start tag
    state](#self-closing-start-tag-state){#after-attribute-name-state:self-closing-start-tag-state}.

U+003D EQUALS SIGN (=)
:   Switch to the [before attribute value
    state](#before-attribute-value-state){#after-attribute-name-state:before-attribute-value-state}.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#after-attribute-name-state:data-state}. Emit
    the current tag token.

EOF
:   This is an
    [eof-in-tag](#parse-error-eof-in-tag){#after-attribute-name-state:parse-error-eof-in-tag}
    [parse
    error](#parse-errors){#after-attribute-name-state:parse-errors}.
    Emit an end-of-file token.

Anything else
:   Start a new attribute in the current tag token. Set that attribute
    name and value to the empty string.
    [Reconsume](#reconsume){#after-attribute-name-state:reconsume} in
    the [attribute name
    state](#attribute-name-state){#after-attribute-name-state:attribute-name-state}.

##### [13.2.5.35]{.secno} [Before attribute value state]{.dfn}[](#before-attribute-value-state){.self-link}

Consume the [next input
character](#next-input-character){#before-attribute-value-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+0022 QUOTATION MARK (\")
:   Switch to the [attribute value (double-quoted)
    state](#attribute-value-(double-quoted)-state){#before-attribute-value-state:attribute-value-(double-quoted)-state}.

U+0027 APOSTROPHE (\')
:   Switch to the [attribute value (single-quoted)
    state](#attribute-value-(single-quoted)-state){#before-attribute-value-state:attribute-value-(single-quoted)-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is a
    [missing-attribute-value](#parse-error-missing-attribute-value){#before-attribute-value-state:parse-error-missing-attribute-value}
    [parse
    error](#parse-errors){#before-attribute-value-state:parse-errors}.
    Switch to the [data
    state](#data-state){#before-attribute-value-state:data-state}. Emit
    the current tag token.

Anything else
:   [Reconsume](#reconsume){#before-attribute-value-state:reconsume} in
    the [attribute value (unquoted)
    state](#attribute-value-(unquoted)-state){#before-attribute-value-state:attribute-value-(unquoted)-state}.

##### [13.2.5.36]{.secno} [Attribute value (double-quoted) state]{.dfn}[](#attribute-value-(double-quoted)-state){.self-link} {#attribute-value-(double-quoted)-state}

Consume the [next input
character](#next-input-character){#attribute-value-(double-quoted)-state:next-input-character}:

U+0022 QUOTATION MARK (\")
:   Switch to the [after attribute value (quoted)
    state](#after-attribute-value-(quoted)-state){#attribute-value-(double-quoted)-state:after-attribute-value-(quoted)-state}.

U+0026 AMPERSAND (&)
:   Set the
    [`return state`{#attribute-value-(double-quoted)-state:return-state
    .variable}](#return-state) to the [attribute value (double-quoted)
    state](#attribute-value-(double-quoted)-state){#attribute-value-(double-quoted)-state:attribute-value-(double-quoted)-state}.
    Switch to the [character reference
    state](#character-reference-state){#attribute-value-(double-quoted)-state:character-reference-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#attribute-value-(double-quoted)-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#attribute-value-(double-quoted)-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    attribute\'s value.

EOF
:   This is an
    [eof-in-tag](#parse-error-eof-in-tag){#attribute-value-(double-quoted)-state:parse-error-eof-in-tag}
    [parse
    error](#parse-errors){#attribute-value-(double-quoted)-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#attribute-value-(double-quoted)-state:current-input-character}
    to the current attribute\'s value.

##### [13.2.5.37]{.secno} [Attribute value (single-quoted) state]{.dfn}[](#attribute-value-(single-quoted)-state){.self-link} {#attribute-value-(single-quoted)-state}

Consume the [next input
character](#next-input-character){#attribute-value-(single-quoted)-state:next-input-character}:

U+0027 APOSTROPHE (\')
:   Switch to the [after attribute value (quoted)
    state](#after-attribute-value-(quoted)-state){#attribute-value-(single-quoted)-state:after-attribute-value-(quoted)-state}.

U+0026 AMPERSAND (&)
:   Set the
    [`return state`{#attribute-value-(single-quoted)-state:return-state
    .variable}](#return-state) to the [attribute value (single-quoted)
    state](#attribute-value-(single-quoted)-state){#attribute-value-(single-quoted)-state:attribute-value-(single-quoted)-state}.
    Switch to the [character reference
    state](#character-reference-state){#attribute-value-(single-quoted)-state:character-reference-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#attribute-value-(single-quoted)-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#attribute-value-(single-quoted)-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    attribute\'s value.

EOF
:   This is an
    [eof-in-tag](#parse-error-eof-in-tag){#attribute-value-(single-quoted)-state:parse-error-eof-in-tag}
    [parse
    error](#parse-errors){#attribute-value-(single-quoted)-state:parse-errors-2}.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#attribute-value-(single-quoted)-state:current-input-character}
    to the current attribute\'s value.

##### [13.2.5.38]{.secno} [Attribute value (unquoted) state]{.dfn}[](#attribute-value-(unquoted)-state){.self-link} {#attribute-value-(unquoted)-state}

Consume the [next input
character](#next-input-character){#attribute-value-(unquoted)-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [before attribute name
    state](#before-attribute-name-state){#attribute-value-(unquoted)-state:before-attribute-name-state}.

U+0026 AMPERSAND (&)
:   Set the
    [`return state`{#attribute-value-(unquoted)-state:return-state
    .variable}](#return-state) to the [attribute value (unquoted)
    state](#attribute-value-(unquoted)-state){#attribute-value-(unquoted)-state:attribute-value-(unquoted)-state}.
    Switch to the [character reference
    state](#character-reference-state){#attribute-value-(unquoted)-state:character-reference-state}.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#attribute-value-(unquoted)-state:data-state}.
    Emit the current tag token.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#attribute-value-(unquoted)-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#attribute-value-(unquoted)-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    attribute\'s value.

U+0022 QUOTATION MARK (\")\
U+0027 APOSTROPHE (\')\
U+003C LESS-THAN SIGN (\<)\
U+003D EQUALS SIGN (=)\
U+0060 GRAVE ACCENT (\`)
:   This is an
    [unexpected-character-in-unquoted-attribute-value](#parse-error-unexpected-character-in-unquoted-attribute-value){#attribute-value-(unquoted)-state:parse-error-unexpected-character-in-unquoted-attribute-value}
    [parse
    error](#parse-errors){#attribute-value-(unquoted)-state:parse-errors-2}.
    Treat it as per the \"anything else\" entry below.

EOF
:   This is an
    [eof-in-tag](#parse-error-eof-in-tag){#attribute-value-(unquoted)-state:parse-error-eof-in-tag}
    [parse
    error](#parse-errors){#attribute-value-(unquoted)-state:parse-errors-3}.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#attribute-value-(unquoted)-state:current-input-character}
    to the current attribute\'s value.

##### [13.2.5.39]{.secno} [After attribute value (quoted) state]{.dfn}[](#after-attribute-value-(quoted)-state){.self-link} {#after-attribute-value-(quoted)-state}

Consume the [next input
character](#next-input-character){#after-attribute-value-(quoted)-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [before attribute name
    state](#before-attribute-name-state){#after-attribute-value-(quoted)-state:before-attribute-name-state}.

U+002F SOLIDUS (/)
:   Switch to the [self-closing start tag
    state](#self-closing-start-tag-state){#after-attribute-value-(quoted)-state:self-closing-start-tag-state}.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#after-attribute-value-(quoted)-state:data-state}.
    Emit the current tag token.

EOF
:   This is an
    [eof-in-tag](#parse-error-eof-in-tag){#after-attribute-value-(quoted)-state:parse-error-eof-in-tag}
    [parse
    error](#parse-errors){#after-attribute-value-(quoted)-state:parse-errors}.
    Emit an end-of-file token.

Anything else
:   This is a
    [missing-whitespace-between-attributes](#parse-error-missing-whitespace-between-attributes){#after-attribute-value-(quoted)-state:parse-error-missing-whitespace-between-attributes}
    [parse
    error](#parse-errors){#after-attribute-value-(quoted)-state:parse-errors-2}.
    [Reconsume](#reconsume){#after-attribute-value-(quoted)-state:reconsume}
    in the [before attribute name
    state](#before-attribute-name-state){#after-attribute-value-(quoted)-state:before-attribute-name-state-2}.

##### [13.2.5.40]{.secno} [Self-closing start tag state]{.dfn}[](#self-closing-start-tag-state){.self-link}

Consume the [next input
character](#next-input-character){#self-closing-start-tag-state:next-input-character}:

U+003E GREATER-THAN SIGN (\>)
:   Set the *[self-closing flag](#self-closing-flag)* of the current tag
    token. Switch to the [data
    state](#data-state){#self-closing-start-tag-state:data-state}. Emit
    the current tag token.

EOF
:   This is an
    [eof-in-tag](#parse-error-eof-in-tag){#self-closing-start-tag-state:parse-error-eof-in-tag}
    [parse
    error](#parse-errors){#self-closing-start-tag-state:parse-errors}.
    Emit an end-of-file token.

Anything else
:   This is an
    [unexpected-solidus-in-tag](#parse-error-unexpected-solidus-in-tag){#self-closing-start-tag-state:parse-error-unexpected-solidus-in-tag}
    [parse
    error](#parse-errors){#self-closing-start-tag-state:parse-errors-2}.
    [Reconsume](#reconsume){#self-closing-start-tag-state:reconsume} in
    the [before attribute name
    state](#before-attribute-name-state){#self-closing-start-tag-state:before-attribute-name-state}.

##### [13.2.5.41]{.secno} [Bogus comment state]{.dfn}[](#bogus-comment-state){.self-link}

Consume the [next input
character](#next-input-character){#bogus-comment-state:next-input-character}:

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#bogus-comment-state:data-state}. Emit the
    current comment token.

EOF
:   Emit the comment. Emit an end-of-file token.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#bogus-comment-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#bogus-comment-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the comment
    token\'s data.

Anything else
:   Append the [current input
    character](#current-input-character){#bogus-comment-state:current-input-character}
    to the comment token\'s data.

##### [13.2.5.42]{.secno} [Markup declaration open state]{.dfn}[](#markup-declaration-open-state){.self-link}

If the next few characters are:

Two U+002D HYPHEN-MINUS characters (-)
:   Consume those two characters, create a comment token whose data is
    the empty string, and switch to the [comment start
    state](#comment-start-state){#markup-declaration-open-state:comment-start-state}.

[ASCII case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#markup-declaration-open-state:ascii-case-insensitive x-internal="ascii-case-insensitive"} match for the word \"DOCTYPE\"
:   Consume those characters and switch to the [DOCTYPE
    state](#doctype-state){#markup-declaration-open-state:doctype-state}.

The string \"\[CDATA\[\" (the five uppercase letters \"CDATA\" with a U+005B LEFT SQUARE BRACKET character before and after)
:   Consume those characters. If there is an [adjusted current
    node](#adjusted-current-node){#markup-declaration-open-state:adjusted-current-node}
    and it is not an element in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#markup-declaration-open-state:html-namespace-2
    x-internal="html-namespace-2"}, then switch to the [CDATA section
    state](#cdata-section-state){#markup-declaration-open-state:cdata-section-state}.
    Otherwise, this is a
    [cdata-in-html-content](#parse-error-cdata-in-html-content){#markup-declaration-open-state:parse-error-cdata-in-html-content}
    [parse
    error](#parse-errors){#markup-declaration-open-state:parse-errors}.
    Create a comment token whose data is the \"\[CDATA\[\" string.
    Switch to the [bogus comment
    state](#bogus-comment-state){#markup-declaration-open-state:bogus-comment-state}.

Anything else
:   This is an
    [incorrectly-opened-comment](#parse-error-incorrectly-opened-comment){#markup-declaration-open-state:parse-error-incorrectly-opened-comment}
    [parse
    error](#parse-errors){#markup-declaration-open-state:parse-errors-2}.
    Create a comment token whose data is the empty string. Switch to the
    [bogus comment
    state](#bogus-comment-state){#markup-declaration-open-state:bogus-comment-state-2}
    (don\'t consume anything in the current state).

##### [13.2.5.43]{.secno} [Comment start state]{.dfn}[](#comment-start-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-start-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [comment start dash
    state](#comment-start-dash-state){#comment-start-state:comment-start-dash-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is an
    [abrupt-closing-of-empty-comment](#parse-error-abrupt-closing-of-empty-comment){#comment-start-state:parse-error-abrupt-closing-of-empty-comment}
    [parse error](#parse-errors){#comment-start-state:parse-errors}.
    Switch to the [data
    state](#data-state){#comment-start-state:data-state}. Emit the
    current comment token.

Anything else
:   [Reconsume](#reconsume){#comment-start-state:reconsume} in the
    [comment state](#comment-state){#comment-start-state:comment-state}.

##### [13.2.5.44]{.secno} [Comment start dash state]{.dfn}[](#comment-start-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-start-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [comment end
    state](#comment-end-state){#comment-start-dash-state:comment-end-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is an
    [abrupt-closing-of-empty-comment](#parse-error-abrupt-closing-of-empty-comment){#comment-start-dash-state:parse-error-abrupt-closing-of-empty-comment}
    [parse
    error](#parse-errors){#comment-start-dash-state:parse-errors}.
    Switch to the [data
    state](#data-state){#comment-start-dash-state:data-state}. Emit the
    current comment token.

EOF
:   This is an
    [eof-in-comment](#parse-error-eof-in-comment){#comment-start-dash-state:parse-error-eof-in-comment}
    [parse
    error](#parse-errors){#comment-start-dash-state:parse-errors-2}.
    Emit the current comment token. Emit an end-of-file token.

Anything else
:   Append a U+002D HYPHEN-MINUS character (-) to the comment token\'s
    data. [Reconsume](#reconsume){#comment-start-dash-state:reconsume}
    in the [comment
    state](#comment-state){#comment-start-dash-state:comment-state}.

##### [13.2.5.45]{.secno} [Comment state]{#comment .dfn}[](#comment-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-state:next-input-character}:

U+003C LESS-THAN SIGN (\<)
:   Append the [current input
    character](#current-input-character){#comment-state:current-input-character}
    to the comment token\'s data. Switch to the [comment less-than sign
    state](#comment-less-than-sign-state){#comment-state:comment-less-than-sign-state}.

U+002D HYPHEN-MINUS (-)
:   Switch to the [comment end dash
    state](#comment-end-dash-state){#comment-state:comment-end-dash-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#comment-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#comment-state:parse-errors}. Append a
    U+FFFD REPLACEMENT CHARACTER character to the comment token\'s data.

EOF
:   This is an
    [eof-in-comment](#parse-error-eof-in-comment){#comment-state:parse-error-eof-in-comment}
    [parse error](#parse-errors){#comment-state:parse-errors-2}. Emit
    the current comment token. Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#comment-state:current-input-character-2}
    to the comment token\'s data.

##### [13.2.5.46]{.secno} [Comment less-than sign state]{.dfn}[](#comment-less-than-sign-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-less-than-sign-state:next-input-character}:

U+0021 EXCLAMATION MARK (!)
:   Append the [current input
    character](#current-input-character){#comment-less-than-sign-state:current-input-character}
    to the comment token\'s data. Switch to the [comment less-than sign
    bang
    state](#comment-less-than-sign-bang-state){#comment-less-than-sign-state:comment-less-than-sign-bang-state}.

U+003C LESS-THAN SIGN (\<)
:   Append the [current input
    character](#current-input-character){#comment-less-than-sign-state:current-input-character-2}
    to the comment token\'s data.

Anything else
:   [Reconsume](#reconsume){#comment-less-than-sign-state:reconsume} in
    the [comment
    state](#comment-state){#comment-less-than-sign-state:comment-state}.

##### [13.2.5.47]{.secno} [Comment less-than sign bang state]{.dfn}[](#comment-less-than-sign-bang-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-less-than-sign-bang-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [comment less-than sign bang dash
    state](#comment-less-than-sign-bang-dash-state){#comment-less-than-sign-bang-state:comment-less-than-sign-bang-dash-state}.

Anything else
:   [Reconsume](#reconsume){#comment-less-than-sign-bang-state:reconsume}
    in the [comment
    state](#comment-state){#comment-less-than-sign-bang-state:comment-state}.

##### [13.2.5.48]{.secno} [Comment less-than sign bang dash state]{.dfn}[](#comment-less-than-sign-bang-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-less-than-sign-bang-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [comment less-than sign bang dash dash
    state](#comment-less-than-sign-bang-dash-dash-state){#comment-less-than-sign-bang-dash-state:comment-less-than-sign-bang-dash-dash-state}.

Anything else
:   [Reconsume](#reconsume){#comment-less-than-sign-bang-dash-state:reconsume}
    in the [comment end dash
    state](#comment-end-dash-state){#comment-less-than-sign-bang-dash-state:comment-end-dash-state}.

##### [13.2.5.49]{.secno} [Comment less-than sign bang dash dash state]{.dfn}[](#comment-less-than-sign-bang-dash-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-less-than-sign-bang-dash-dash-state:next-input-character}:

U+003E GREATER-THAN SIGN (\>)\
EOF
:   [Reconsume](#reconsume){#comment-less-than-sign-bang-dash-dash-state:reconsume}
    in the [comment end
    state](#comment-end-state){#comment-less-than-sign-bang-dash-dash-state:comment-end-state}.

Anything else
:   This is a
    [nested-comment](#parse-error-nested-comment){#comment-less-than-sign-bang-dash-dash-state:parse-error-nested-comment}
    [parse
    error](#parse-errors){#comment-less-than-sign-bang-dash-dash-state:parse-errors}.
    [Reconsume](#reconsume){#comment-less-than-sign-bang-dash-dash-state:reconsume-2}
    in the [comment end
    state](#comment-end-state){#comment-less-than-sign-bang-dash-dash-state:comment-end-state-2}.

##### [13.2.5.50]{.secno} [Comment end dash state]{.dfn}[](#comment-end-dash-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-end-dash-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Switch to the [comment end
    state](#comment-end-state){#comment-end-dash-state:comment-end-state}.

EOF
:   This is an
    [eof-in-comment](#parse-error-eof-in-comment){#comment-end-dash-state:parse-error-eof-in-comment}
    [parse error](#parse-errors){#comment-end-dash-state:parse-errors}.
    Emit the current comment token. Emit an end-of-file token.

Anything else
:   Append a U+002D HYPHEN-MINUS character (-) to the comment token\'s
    data. [Reconsume](#reconsume){#comment-end-dash-state:reconsume} in
    the [comment
    state](#comment-state){#comment-end-dash-state:comment-state}.

##### [13.2.5.51]{.secno} [Comment end state]{.dfn}[](#comment-end-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-end-state:next-input-character}:

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#comment-end-state:data-state}. Emit the current
    comment token.

U+0021 EXCLAMATION MARK (!)
:   Switch to the [comment end bang
    state](#comment-end-bang-state){#comment-end-state:comment-end-bang-state}.

U+002D HYPHEN-MINUS (-)
:   Append a U+002D HYPHEN-MINUS character (-) to the comment token\'s
    data.

EOF
:   This is an
    [eof-in-comment](#parse-error-eof-in-comment){#comment-end-state:parse-error-eof-in-comment}
    [parse error](#parse-errors){#comment-end-state:parse-errors}. Emit
    the current comment token. Emit an end-of-file token.

Anything else
:   Append two U+002D HYPHEN-MINUS characters (-) to the comment
    token\'s data. [Reconsume](#reconsume){#comment-end-state:reconsume}
    in the [comment
    state](#comment-state){#comment-end-state:comment-state}.

##### [13.2.5.52]{.secno} [Comment end bang state]{.dfn}[](#comment-end-bang-state){.self-link}

Consume the [next input
character](#next-input-character){#comment-end-bang-state:next-input-character}:

U+002D HYPHEN-MINUS (-)
:   Append two U+002D HYPHEN-MINUS characters (-) and a U+0021
    EXCLAMATION MARK character (!) to the comment token\'s data. Switch
    to the [comment end dash
    state](#comment-end-dash-state){#comment-end-bang-state:comment-end-dash-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is an
    [incorrectly-closed-comment](#parse-error-incorrectly-closed-comment){#comment-end-bang-state:parse-error-incorrectly-closed-comment}
    [parse error](#parse-errors){#comment-end-bang-state:parse-errors}.
    Switch to the [data
    state](#data-state){#comment-end-bang-state:data-state}. Emit the
    current comment token.

EOF
:   This is an
    [eof-in-comment](#parse-error-eof-in-comment){#comment-end-bang-state:parse-error-eof-in-comment}
    [parse
    error](#parse-errors){#comment-end-bang-state:parse-errors-2}. Emit
    the current comment token. Emit an end-of-file token.

Anything else
:   Append two U+002D HYPHEN-MINUS characters (-) and a U+0021
    EXCLAMATION MARK character (!) to the comment token\'s data.
    [Reconsume](#reconsume){#comment-end-bang-state:reconsume} in the
    [comment
    state](#comment-state){#comment-end-bang-state:comment-state}.

##### [13.2.5.53]{.secno} [DOCTYPE state]{.dfn}[](#doctype-state){.self-link}

Consume the [next input
character](#next-input-character){#doctype-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [before DOCTYPE name
    state](#before-doctype-name-state){#doctype-state:before-doctype-name-state}.

U+003E GREATER-THAN SIGN (\>)
:   [Reconsume](#reconsume){#doctype-state:reconsume} in the [before
    DOCTYPE name
    state](#before-doctype-name-state){#doctype-state:before-doctype-name-state-2}.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#doctype-state:parse-error-eof-in-doctype}
    [parse error](#parse-errors){#doctype-state:parse-errors}. Create a
    new DOCTYPE token. Set its *[force-quirks flag](#force-quirks-flag)*
    to *on*. Emit the current token. Emit an end-of-file token.

Anything else
:   This is a
    [missing-whitespace-before-doctype-name](#parse-error-missing-whitespace-before-doctype-name){#doctype-state:parse-error-missing-whitespace-before-doctype-name}
    [parse error](#parse-errors){#doctype-state:parse-errors-2}.
    [Reconsume](#reconsume){#doctype-state:reconsume-2} in the [before
    DOCTYPE name
    state](#before-doctype-name-state){#doctype-state:before-doctype-name-state-3}.

##### [13.2.5.54]{.secno} [Before DOCTYPE name state]{.dfn}[](#before-doctype-name-state){.self-link}

Consume the [next input
character](#next-input-character){#before-doctype-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#before-doctype-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Create a new DOCTYPE token. Set the token\'s name to the lowercase
    version of the [current input
    character](#current-input-character){#before-doctype-name-state:current-input-character}
    (add 0x0020 to the character\'s code point). Switch to the [DOCTYPE
    name
    state](#doctype-name-state){#before-doctype-name-state:doctype-name-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#before-doctype-name-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#before-doctype-name-state:parse-errors}.
    Create a new DOCTYPE token. Set the token\'s name to a U+FFFD
    REPLACEMENT CHARACTER character. Switch to the [DOCTYPE name
    state](#doctype-name-state){#before-doctype-name-state:doctype-name-state-2}.

U+003E GREATER-THAN SIGN (\>)
:   This is a
    [missing-doctype-name](#parse-error-missing-doctype-name){#before-doctype-name-state:parse-error-missing-doctype-name}
    [parse
    error](#parse-errors){#before-doctype-name-state:parse-errors-2}.
    Create a new DOCTYPE token. Set its *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#before-doctype-name-state:data-state}. Emit the
    current token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#before-doctype-name-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#before-doctype-name-state:parse-errors-3}.
    Create a new DOCTYPE token. Set its *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current token. Emit an
    end-of-file token.

Anything else
:   Create a new DOCTYPE token. Set the token\'s name to the [current
    input
    character](#current-input-character){#before-doctype-name-state:current-input-character-2}.
    Switch to the [DOCTYPE name
    state](#doctype-name-state){#before-doctype-name-state:doctype-name-state-3}.

##### [13.2.5.55]{.secno} [DOCTYPE name state]{.dfn}[](#doctype-name-state){.self-link}

Consume the [next input
character](#next-input-character){#doctype-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [after DOCTYPE name
    state](#after-doctype-name-state){#doctype-name-state:after-doctype-name-state}.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#doctype-name-state:data-state}. Emit the
    current DOCTYPE token.

[ASCII upper alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha){#doctype-name-state:uppercase-ascii-letters x-internal="uppercase-ascii-letters"}
:   Append the lowercase version of the [current input
    character](#current-input-character){#doctype-name-state:current-input-character}
    (add 0x0020 to the character\'s code point) to the current DOCTYPE
    token\'s name.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#doctype-name-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#doctype-name-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    DOCTYPE token\'s name.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#doctype-name-state:parse-error-eof-in-doctype}
    [parse error](#parse-errors){#doctype-name-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#doctype-name-state:current-input-character-2}
    to the current DOCTYPE token\'s name.

##### [13.2.5.56]{.secno} [After DOCTYPE name state]{.dfn}[](#after-doctype-name-state){.self-link}

Consume the [next input
character](#next-input-character){#after-doctype-name-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#after-doctype-name-state:data-state}. Emit the
    current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#after-doctype-name-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#after-doctype-name-state:parse-errors}. Set
    the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else

:   If the six characters starting from the [current input
    character](#current-input-character){#after-doctype-name-state:current-input-character}
    are an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#after-doctype-name-state:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for the word \"PUBLIC\",
    then consume those characters and switch to the [after DOCTYPE
    public keyword
    state](#after-doctype-public-keyword-state){#after-doctype-name-state:after-doctype-public-keyword-state}.

    Otherwise, if the six characters starting from the [current input
    character](#current-input-character){#after-doctype-name-state:current-input-character-2}
    are an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#after-doctype-name-state:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for the word \"SYSTEM\",
    then consume those characters and switch to the [after DOCTYPE
    system keyword
    state](#after-doctype-system-keyword-state){#after-doctype-name-state:after-doctype-system-keyword-state}.

    Otherwise, this is an
    [invalid-character-sequence-after-doctype-name](#parse-error-invalid-character-sequence-after-doctype-name){#after-doctype-name-state:parse-error-invalid-character-sequence-after-doctype-name}
    [parse
    error](#parse-errors){#after-doctype-name-state:parse-errors-2}. Set
    the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.
    [Reconsume](#reconsume){#after-doctype-name-state:reconsume} in the
    [bogus DOCTYPE
    state](#bogus-doctype-state){#after-doctype-name-state:bogus-doctype-state}.

##### [13.2.5.57]{.secno} [After DOCTYPE public keyword state]{.dfn}[](#after-doctype-public-keyword-state){.self-link}

Consume the [next input
character](#next-input-character){#after-doctype-public-keyword-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [before DOCTYPE public identifier
    state](#before-doctype-public-identifier-state){#after-doctype-public-keyword-state:before-doctype-public-identifier-state}.

U+0022 QUOTATION MARK (\")
:   This is a
    [missing-whitespace-after-doctype-public-keyword](#parse-error-missing-whitespace-after-doctype-public-keyword){#after-doctype-public-keyword-state:parse-error-missing-whitespace-after-doctype-public-keyword}
    [parse
    error](#parse-errors){#after-doctype-public-keyword-state:parse-errors}.
    Set the current DOCTYPE token\'s public identifier to the empty
    string (not missing), then switch to the [DOCTYPE public identifier
    (double-quoted)
    state](#doctype-public-identifier-(double-quoted)-state){#after-doctype-public-keyword-state:doctype-public-identifier-(double-quoted)-state}.

U+0027 APOSTROPHE (\')
:   This is a
    [missing-whitespace-after-doctype-public-keyword](#parse-error-missing-whitespace-after-doctype-public-keyword){#after-doctype-public-keyword-state:parse-error-missing-whitespace-after-doctype-public-keyword-2}
    [parse
    error](#parse-errors){#after-doctype-public-keyword-state:parse-errors-2}.
    Set the current DOCTYPE token\'s public identifier to the empty
    string (not missing), then switch to the [DOCTYPE public identifier
    (single-quoted)
    state](#doctype-public-identifier-(single-quoted)-state){#after-doctype-public-keyword-state:doctype-public-identifier-(single-quoted)-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is a
    [missing-doctype-public-identifier](#parse-error-missing-doctype-public-identifier){#after-doctype-public-keyword-state:parse-error-missing-doctype-public-identifier}
    [parse
    error](#parse-errors){#after-doctype-public-keyword-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#after-doctype-public-keyword-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#after-doctype-public-keyword-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#after-doctype-public-keyword-state:parse-errors-4}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   This is a
    [missing-quote-before-doctype-public-identifier](#parse-error-missing-quote-before-doctype-public-identifier){#after-doctype-public-keyword-state:parse-error-missing-quote-before-doctype-public-identifier}
    [parse
    error](#parse-errors){#after-doctype-public-keyword-state:parse-errors-5}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.
    [Reconsume](#reconsume){#after-doctype-public-keyword-state:reconsume}
    in the [bogus DOCTYPE
    state](#bogus-doctype-state){#after-doctype-public-keyword-state:bogus-doctype-state}.

##### [13.2.5.58]{.secno} [Before DOCTYPE public identifier state]{.dfn}[](#before-doctype-public-identifier-state){.self-link}

Consume the [next input
character](#next-input-character){#before-doctype-public-identifier-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+0022 QUOTATION MARK (\")
:   Set the current DOCTYPE token\'s public identifier to the empty
    string (not missing), then switch to the [DOCTYPE public identifier
    (double-quoted)
    state](#doctype-public-identifier-(double-quoted)-state){#before-doctype-public-identifier-state:doctype-public-identifier-(double-quoted)-state}.

U+0027 APOSTROPHE (\')
:   Set the current DOCTYPE token\'s public identifier to the empty
    string (not missing), then switch to the [DOCTYPE public identifier
    (single-quoted)
    state](#doctype-public-identifier-(single-quoted)-state){#before-doctype-public-identifier-state:doctype-public-identifier-(single-quoted)-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is a
    [missing-doctype-public-identifier](#parse-error-missing-doctype-public-identifier){#before-doctype-public-identifier-state:parse-error-missing-doctype-public-identifier}
    [parse
    error](#parse-errors){#before-doctype-public-identifier-state:parse-errors}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#before-doctype-public-identifier-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#before-doctype-public-identifier-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#before-doctype-public-identifier-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   This is a
    [missing-quote-before-doctype-public-identifier](#parse-error-missing-quote-before-doctype-public-identifier){#before-doctype-public-identifier-state:parse-error-missing-quote-before-doctype-public-identifier}
    [parse
    error](#parse-errors){#before-doctype-public-identifier-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.
    [Reconsume](#reconsume){#before-doctype-public-identifier-state:reconsume}
    in the [bogus DOCTYPE
    state](#bogus-doctype-state){#before-doctype-public-identifier-state:bogus-doctype-state}.

##### [13.2.5.59]{.secno} [DOCTYPE public identifier (double-quoted) state]{.dfn}[](#doctype-public-identifier-(double-quoted)-state){.self-link} {#doctype-public-identifier-(double-quoted)-state}

Consume the [next input
character](#next-input-character){#doctype-public-identifier-(double-quoted)-state:next-input-character}:

U+0022 QUOTATION MARK (\")
:   Switch to the [after DOCTYPE public identifier
    state](#after-doctype-public-identifier-state){#doctype-public-identifier-(double-quoted)-state:after-doctype-public-identifier-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#doctype-public-identifier-(double-quoted)-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#doctype-public-identifier-(double-quoted)-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    DOCTYPE token\'s public identifier.

U+003E GREATER-THAN SIGN (\>)
:   This is an
    [abrupt-doctype-public-identifier](#parse-error-abrupt-doctype-public-identifier){#doctype-public-identifier-(double-quoted)-state:parse-error-abrupt-doctype-public-identifier}
    [parse
    error](#parse-errors){#doctype-public-identifier-(double-quoted)-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#doctype-public-identifier-(double-quoted)-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#doctype-public-identifier-(double-quoted)-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#doctype-public-identifier-(double-quoted)-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#doctype-public-identifier-(double-quoted)-state:current-input-character}
    to the current DOCTYPE token\'s public identifier.

##### [13.2.5.60]{.secno} [DOCTYPE public identifier (single-quoted) state]{.dfn}[](#doctype-public-identifier-(single-quoted)-state){.self-link} {#doctype-public-identifier-(single-quoted)-state}

Consume the [next input
character](#next-input-character){#doctype-public-identifier-(single-quoted)-state:next-input-character}:

U+0027 APOSTROPHE (\')
:   Switch to the [after DOCTYPE public identifier
    state](#after-doctype-public-identifier-state){#doctype-public-identifier-(single-quoted)-state:after-doctype-public-identifier-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#doctype-public-identifier-(single-quoted)-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#doctype-public-identifier-(single-quoted)-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    DOCTYPE token\'s public identifier.

U+003E GREATER-THAN SIGN (\>)
:   This is an
    [abrupt-doctype-public-identifier](#parse-error-abrupt-doctype-public-identifier){#doctype-public-identifier-(single-quoted)-state:parse-error-abrupt-doctype-public-identifier}
    [parse
    error](#parse-errors){#doctype-public-identifier-(single-quoted)-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#doctype-public-identifier-(single-quoted)-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#doctype-public-identifier-(single-quoted)-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#doctype-public-identifier-(single-quoted)-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#doctype-public-identifier-(single-quoted)-state:current-input-character}
    to the current DOCTYPE token\'s public identifier.

##### [13.2.5.61]{.secno} [After DOCTYPE public identifier state]{.dfn}[](#after-doctype-public-identifier-state){.self-link}

Consume the [next input
character](#next-input-character){#after-doctype-public-identifier-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [between DOCTYPE public and system identifiers
    state](#between-doctype-public-and-system-identifiers-state){#after-doctype-public-identifier-state:between-doctype-public-and-system-identifiers-state}.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#after-doctype-public-identifier-state:data-state}.
    Emit the current DOCTYPE token.

U+0022 QUOTATION MARK (\")
:   This is a
    [missing-whitespace-between-doctype-public-and-system-identifiers](#parse-error-missing-whitespace-between-doctype-public-and-system-identifiers){#after-doctype-public-identifier-state:parse-error-missing-whitespace-between-doctype-public-and-system-identifiers}
    [parse
    error](#parse-errors){#after-doctype-public-identifier-state:parse-errors}.
    Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (double-quoted)
    state](#doctype-system-identifier-(double-quoted)-state){#after-doctype-public-identifier-state:doctype-system-identifier-(double-quoted)-state}.

U+0027 APOSTROPHE (\')
:   This is a
    [missing-whitespace-between-doctype-public-and-system-identifiers](#parse-error-missing-whitespace-between-doctype-public-and-system-identifiers){#after-doctype-public-identifier-state:parse-error-missing-whitespace-between-doctype-public-and-system-identifiers-2}
    [parse
    error](#parse-errors){#after-doctype-public-identifier-state:parse-errors-2}.
    Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (single-quoted)
    state](#doctype-system-identifier-(single-quoted)-state){#after-doctype-public-identifier-state:doctype-system-identifier-(single-quoted)-state}.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#after-doctype-public-identifier-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#after-doctype-public-identifier-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   This is a
    [missing-quote-before-doctype-system-identifier](#parse-error-missing-quote-before-doctype-system-identifier){#after-doctype-public-identifier-state:parse-error-missing-quote-before-doctype-system-identifier}
    [parse
    error](#parse-errors){#after-doctype-public-identifier-state:parse-errors-4}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.
    [Reconsume](#reconsume){#after-doctype-public-identifier-state:reconsume}
    in the [bogus DOCTYPE
    state](#bogus-doctype-state){#after-doctype-public-identifier-state:bogus-doctype-state}.

##### [13.2.5.62]{.secno} [Between DOCTYPE public and system identifiers state]{.dfn}[](#between-doctype-public-and-system-identifiers-state){.self-link}

Consume the [next input
character](#next-input-character){#between-doctype-public-and-system-identifiers-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#between-doctype-public-and-system-identifiers-state:data-state}.
    Emit the current DOCTYPE token.

U+0022 QUOTATION MARK (\")
:   Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (double-quoted)
    state](#doctype-system-identifier-(double-quoted)-state){#between-doctype-public-and-system-identifiers-state:doctype-system-identifier-(double-quoted)-state}.

U+0027 APOSTROPHE (\')
:   Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (single-quoted)
    state](#doctype-system-identifier-(single-quoted)-state){#between-doctype-public-and-system-identifiers-state:doctype-system-identifier-(single-quoted)-state}.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#between-doctype-public-and-system-identifiers-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#between-doctype-public-and-system-identifiers-state:parse-errors}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   This is a
    [missing-quote-before-doctype-system-identifier](#parse-error-missing-quote-before-doctype-system-identifier){#between-doctype-public-and-system-identifiers-state:parse-error-missing-quote-before-doctype-system-identifier}
    [parse
    error](#parse-errors){#between-doctype-public-and-system-identifiers-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.
    [Reconsume](#reconsume){#between-doctype-public-and-system-identifiers-state:reconsume}
    in the [bogus DOCTYPE
    state](#bogus-doctype-state){#between-doctype-public-and-system-identifiers-state:bogus-doctype-state}.

##### [13.2.5.63]{.secno} [After DOCTYPE system keyword state]{.dfn}[](#after-doctype-system-keyword-state){.self-link}

Consume the [next input
character](#next-input-character){#after-doctype-system-keyword-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Switch to the [before DOCTYPE system identifier
    state](#before-doctype-system-identifier-state){#after-doctype-system-keyword-state:before-doctype-system-identifier-state}.

U+0022 QUOTATION MARK (\")
:   This is a
    [missing-whitespace-after-doctype-system-keyword](#parse-error-missing-whitespace-after-doctype-system-keyword){#after-doctype-system-keyword-state:parse-error-missing-whitespace-after-doctype-system-keyword}
    [parse
    error](#parse-errors){#after-doctype-system-keyword-state:parse-errors}.
    Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (double-quoted)
    state](#doctype-system-identifier-(double-quoted)-state){#after-doctype-system-keyword-state:doctype-system-identifier-(double-quoted)-state}.

U+0027 APOSTROPHE (\')
:   This is a
    [missing-whitespace-after-doctype-system-keyword](#parse-error-missing-whitespace-after-doctype-system-keyword){#after-doctype-system-keyword-state:parse-error-missing-whitespace-after-doctype-system-keyword-2}
    [parse
    error](#parse-errors){#after-doctype-system-keyword-state:parse-errors-2}.
    Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (single-quoted)
    state](#doctype-system-identifier-(single-quoted)-state){#after-doctype-system-keyword-state:doctype-system-identifier-(single-quoted)-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is a
    [missing-doctype-system-identifier](#parse-error-missing-doctype-system-identifier){#after-doctype-system-keyword-state:parse-error-missing-doctype-system-identifier}
    [parse
    error](#parse-errors){#after-doctype-system-keyword-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#after-doctype-system-keyword-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#after-doctype-system-keyword-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#after-doctype-system-keyword-state:parse-errors-4}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   This is a
    [missing-quote-before-doctype-system-identifier](#parse-error-missing-quote-before-doctype-system-identifier){#after-doctype-system-keyword-state:parse-error-missing-quote-before-doctype-system-identifier}
    [parse
    error](#parse-errors){#after-doctype-system-keyword-state:parse-errors-5}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.
    [Reconsume](#reconsume){#after-doctype-system-keyword-state:reconsume}
    in the [bogus DOCTYPE
    state](#bogus-doctype-state){#after-doctype-system-keyword-state:bogus-doctype-state}.

##### [13.2.5.64]{.secno} [Before DOCTYPE system identifier state]{.dfn}[](#before-doctype-system-identifier-state){.self-link}

Consume the [next input
character](#next-input-character){#before-doctype-system-identifier-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+0022 QUOTATION MARK (\")
:   Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (double-quoted)
    state](#doctype-system-identifier-(double-quoted)-state){#before-doctype-system-identifier-state:doctype-system-identifier-(double-quoted)-state}.

U+0027 APOSTROPHE (\')
:   Set the current DOCTYPE token\'s system identifier to the empty
    string (not missing), then switch to the [DOCTYPE system identifier
    (single-quoted)
    state](#doctype-system-identifier-(single-quoted)-state){#before-doctype-system-identifier-state:doctype-system-identifier-(single-quoted)-state}.

U+003E GREATER-THAN SIGN (\>)
:   This is a
    [missing-doctype-system-identifier](#parse-error-missing-doctype-system-identifier){#before-doctype-system-identifier-state:parse-error-missing-doctype-system-identifier}
    [parse
    error](#parse-errors){#before-doctype-system-identifier-state:parse-errors}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#before-doctype-system-identifier-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#before-doctype-system-identifier-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#before-doctype-system-identifier-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   This is a
    [missing-quote-before-doctype-system-identifier](#parse-error-missing-quote-before-doctype-system-identifier){#before-doctype-system-identifier-state:parse-error-missing-quote-before-doctype-system-identifier}
    [parse
    error](#parse-errors){#before-doctype-system-identifier-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.
    [Reconsume](#reconsume){#before-doctype-system-identifier-state:reconsume}
    in the [bogus DOCTYPE
    state](#bogus-doctype-state){#before-doctype-system-identifier-state:bogus-doctype-state}.

##### [13.2.5.65]{.secno} [DOCTYPE system identifier (double-quoted) state]{.dfn}[](#doctype-system-identifier-(double-quoted)-state){.self-link} {#doctype-system-identifier-(double-quoted)-state}

Consume the [next input
character](#next-input-character){#doctype-system-identifier-(double-quoted)-state:next-input-character}:

U+0022 QUOTATION MARK (\")
:   Switch to the [after DOCTYPE system identifier
    state](#after-doctype-system-identifier-state){#doctype-system-identifier-(double-quoted)-state:after-doctype-system-identifier-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#doctype-system-identifier-(double-quoted)-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#doctype-system-identifier-(double-quoted)-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    DOCTYPE token\'s system identifier.

U+003E GREATER-THAN SIGN (\>)
:   This is an
    [abrupt-doctype-system-identifier](#parse-error-abrupt-doctype-system-identifier){#doctype-system-identifier-(double-quoted)-state:parse-error-abrupt-doctype-system-identifier}
    [parse
    error](#parse-errors){#doctype-system-identifier-(double-quoted)-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#doctype-system-identifier-(double-quoted)-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#doctype-system-identifier-(double-quoted)-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#doctype-system-identifier-(double-quoted)-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#doctype-system-identifier-(double-quoted)-state:current-input-character}
    to the current DOCTYPE token\'s system identifier.

##### [13.2.5.66]{.secno} [DOCTYPE system identifier (single-quoted) state]{.dfn}[](#doctype-system-identifier-(single-quoted)-state){.self-link} {#doctype-system-identifier-(single-quoted)-state}

Consume the [next input
character](#next-input-character){#doctype-system-identifier-(single-quoted)-state:next-input-character}:

U+0027 APOSTROPHE (\')
:   Switch to the [after DOCTYPE system identifier
    state](#after-doctype-system-identifier-state){#doctype-system-identifier-(single-quoted)-state:after-doctype-system-identifier-state}.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#doctype-system-identifier-(single-quoted)-state:parse-error-unexpected-null-character}
    [parse
    error](#parse-errors){#doctype-system-identifier-(single-quoted)-state:parse-errors}.
    Append a U+FFFD REPLACEMENT CHARACTER character to the current
    DOCTYPE token\'s system identifier.

U+003E GREATER-THAN SIGN (\>)
:   This is an
    [abrupt-doctype-system-identifier](#parse-error-abrupt-doctype-system-identifier){#doctype-system-identifier-(single-quoted)-state:parse-error-abrupt-doctype-system-identifier}
    [parse
    error](#parse-errors){#doctype-system-identifier-(single-quoted)-state:parse-errors-2}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Switch to the [data
    state](#data-state){#doctype-system-identifier-(single-quoted)-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#doctype-system-identifier-(single-quoted)-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#doctype-system-identifier-(single-quoted)-state:parse-errors-3}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   Append the [current input
    character](#current-input-character){#doctype-system-identifier-(single-quoted)-state:current-input-character}
    to the current DOCTYPE token\'s system identifier.

##### [13.2.5.67]{.secno} [After DOCTYPE system identifier state]{.dfn}[](#after-doctype-system-identifier-state){.self-link}

Consume the [next input
character](#next-input-character){#after-doctype-system-identifier-state:next-input-character}:

U+0009 CHARACTER TABULATION (tab)\
U+000A LINE FEED (LF)\
U+000C FORM FEED (FF)\
U+0020 SPACE
:   Ignore the character.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#after-doctype-system-identifier-state:data-state}.
    Emit the current DOCTYPE token.

EOF
:   This is an
    [eof-in-doctype](#parse-error-eof-in-doctype){#after-doctype-system-identifier-state:parse-error-eof-in-doctype}
    [parse
    error](#parse-errors){#after-doctype-system-identifier-state:parse-errors}.
    Set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*. Emit the current DOCTYPE token.
    Emit an end-of-file token.

Anything else
:   This is an
    [unexpected-character-after-doctype-system-identifier](#parse-error-unexpected-character-after-doctype-system-identifier){#after-doctype-system-identifier-state:parse-error-unexpected-character-after-doctype-system-identifier}
    [parse
    error](#parse-errors){#after-doctype-system-identifier-state:parse-errors-2}.
    [Reconsume](#reconsume){#after-doctype-system-identifier-state:reconsume}
    in the [bogus DOCTYPE
    state](#bogus-doctype-state){#after-doctype-system-identifier-state:bogus-doctype-state}.
    (This does *not* set the current DOCTYPE token\'s *[force-quirks
    flag](#force-quirks-flag)* to *on*.)

##### [13.2.5.68]{.secno} [Bogus DOCTYPE state]{.dfn}[](#bogus-doctype-state){.self-link}

Consume the [next input
character](#next-input-character){#bogus-doctype-state:next-input-character}:

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#bogus-doctype-state:data-state}. Emit the
    DOCTYPE token.

U+0000 NULL
:   This is an
    [unexpected-null-character](#parse-error-unexpected-null-character){#bogus-doctype-state:parse-error-unexpected-null-character}
    [parse error](#parse-errors){#bogus-doctype-state:parse-errors}.
    Ignore the character.

EOF
:   Emit the DOCTYPE token. Emit an end-of-file token.

Anything else
:   Ignore the character.

##### [13.2.5.69]{.secno} [CDATA section state]{.dfn}[](#cdata-section-state){.self-link}

Consume the [next input
character](#next-input-character){#cdata-section-state:next-input-character}:

U+005D RIGHT SQUARE BRACKET (\])
:   Switch to the [CDATA section bracket
    state](#cdata-section-bracket-state){#cdata-section-state:cdata-section-bracket-state}.

EOF
:   This is an
    [eof-in-cdata](#parse-error-eof-in-cdata){#cdata-section-state:parse-error-eof-in-cdata}
    [parse error](#parse-errors){#cdata-section-state:parse-errors}.
    Emit an end-of-file token.

Anything else
:   Emit the [current input
    character](#current-input-character){#cdata-section-state:current-input-character}
    as a character token.

U+0000 NULL characters are handled in the tree construction stage, as
part of the [in foreign
content](#parsing-main-inforeign){#cdata-section-state:parsing-main-inforeign}
insertion mode, which is the only place where CDATA sections can appear.

##### [13.2.5.70]{.secno} [CDATA section bracket state]{.dfn}[](#cdata-section-bracket-state){.self-link}

Consume the [next input
character](#next-input-character){#cdata-section-bracket-state:next-input-character}:

U+005D RIGHT SQUARE BRACKET (\])
:   Switch to the [CDATA section end
    state](#cdata-section-end-state){#cdata-section-bracket-state:cdata-section-end-state}.

Anything else
:   Emit a U+005D RIGHT SQUARE BRACKET character token.
    [Reconsume](#reconsume){#cdata-section-bracket-state:reconsume} in
    the [CDATA section
    state](#cdata-section-state){#cdata-section-bracket-state:cdata-section-state}.

##### [13.2.5.71]{.secno} [CDATA section end state]{.dfn}[](#cdata-section-end-state){.self-link}

Consume the [next input
character](#next-input-character){#cdata-section-end-state:next-input-character}:

U+005D RIGHT SQUARE BRACKET (\])
:   Emit a U+005D RIGHT SQUARE BRACKET character token.

U+003E GREATER-THAN SIGN (\>)
:   Switch to the [data
    state](#data-state){#cdata-section-end-state:data-state}.

Anything else
:   Emit two U+005D RIGHT SQUARE BRACKET character tokens.
    [Reconsume](#reconsume){#cdata-section-end-state:reconsume} in the
    [CDATA section
    state](#cdata-section-state){#cdata-section-end-state:cdata-section-state}.

##### [13.2.5.72]{.secno} [Character reference state]{.dfn}[](#character-reference-state){.self-link}

Set the [`temporary buffer`{#character-reference-state:temporary-buffer
.variable}](#temporary-buffer) to the empty string. Append a U+0026
AMPERSAND (&) character to the
[`temporary buffer`{#character-reference-state:temporary-buffer-2
.variable}](#temporary-buffer). Consume the [next input
character](#next-input-character){#character-reference-state:next-input-character}:

[ASCII alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric){#character-reference-state:alphanumeric-ascii-characters x-internal="alphanumeric-ascii-characters"}
:   [Reconsume](#reconsume){#character-reference-state:reconsume} in the
    [named character reference
    state](#named-character-reference-state){#character-reference-state:named-character-reference-state}.

U+0023 NUMBER SIGN (#)
:   Append the [current input
    character](#current-input-character){#character-reference-state:current-input-character}
    to the
    [`temporary buffer`{#character-reference-state:temporary-buffer-3
    .variable}](#temporary-buffer). Switch to the [numeric character
    reference
    state](#numeric-character-reference-state){#character-reference-state:numeric-character-reference-state}.

Anything else
:   [Flush code points consumed as a character
    reference](#flush-code-points-consumed-as-a-character-reference){#character-reference-state:flush-code-points-consumed-as-a-character-reference}.
    [Reconsume](#reconsume){#character-reference-state:reconsume-2} in
    the [`return state`{#character-reference-state:return-state
    .variable}](#return-state).

##### [13.2.5.73]{.secno} [Named character reference state]{.dfn}[](#named-character-reference-state){.self-link}

Consume the maximum number of characters possible, where the consumed
characters are one of the identifiers in the first column of the [named
character
references](named-characters.html#named-character-references){#named-character-reference-state:named-character-references}
table. Append each character to the
[`temporary buffer`{#named-character-reference-state:temporary-buffer
.variable}](#temporary-buffer) when it\'s consumed.

If there is a match

:   If the character reference was [consumed as part of an
    attribute](#charref-in-attribute){#named-character-reference-state:charref-in-attribute},
    and the last character matched is not a U+003B SEMICOLON character
    (;), and the [next input
    character](#next-input-character){#named-character-reference-state:next-input-character}
    is either a U+003D EQUALS SIGN character (=) or an [ASCII
    alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric){#named-character-reference-state:alphanumeric-ascii-characters
    x-internal="alphanumeric-ascii-characters"}, then, for historical
    reasons, [flush code points consumed as a character
    reference](#flush-code-points-consumed-as-a-character-reference){#named-character-reference-state:flush-code-points-consumed-as-a-character-reference}
    and switch to the
    [`return state`{#named-character-reference-state:return-state
    .variable}](#return-state).

    Otherwise:

    1.  If the last character matched is not a U+003B SEMICOLON
        character (;), then this is a
        [missing-semicolon-after-character-reference](#parse-error-missing-semicolon-after-character-reference){#named-character-reference-state:parse-error-missing-semicolon-after-character-reference}
        [parse
        error](#parse-errors){#named-character-reference-state:parse-errors}.
    2.  Set the
        [`temporary buffer`{#named-character-reference-state:temporary-buffer-2
        .variable}](#temporary-buffer) to the empty string. Append one
        or two characters corresponding to the character reference name
        (as given by the second column of the [named character
        references](named-characters.html#named-character-references){#named-character-reference-state:named-character-references-2}
        table) to the
        [`temporary buffer`{#named-character-reference-state:temporary-buffer-3
        .variable}](#temporary-buffer).
    3.  [Flush code points consumed as a character
        reference](#flush-code-points-consumed-as-a-character-reference){#named-character-reference-state:flush-code-points-consumed-as-a-character-reference-2}.
        Switch to the
        [`return state`{#named-character-reference-state:return-state-2
        .variable}](#return-state).

Otherwise
:   [Flush code points consumed as a character
    reference](#flush-code-points-consumed-as-a-character-reference){#named-character-reference-state:flush-code-points-consumed-as-a-character-reference-3}.
    Switch to the [ambiguous ampersand
    state](#ambiguous-ampersand-state){#named-character-reference-state:ambiguous-ampersand-state}.

::: example
If the markup contains (not in an attribute) the string
`I'm &notit; I tell you`, the character reference is parsed as \"not\",
as in, `I'm ¬it; I tell you` (and this is a parse error). But if the
markup was `I'm &notin; I tell you`, the character reference would be
parsed as \"notin;\", resulting in `I'm ∉ I tell you` (and no parse
error).

However, if the markup contains the string `I'm &notit; I tell you` in
an attribute, no character reference is parsed and string remains intact
(and there is no parse error).
:::

##### [13.2.5.74]{.secno} [Ambiguous ampersand state]{.dfn}[](#ambiguous-ampersand-state){.self-link}

Consume the [next input
character](#next-input-character){#ambiguous-ampersand-state:next-input-character}:

[ASCII alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric){#ambiguous-ampersand-state:alphanumeric-ascii-characters x-internal="alphanumeric-ascii-characters"}
:   If the character reference was [consumed as part of an
    attribute](#charref-in-attribute){#ambiguous-ampersand-state:charref-in-attribute},
    then append the [current input
    character](#current-input-character){#ambiguous-ampersand-state:current-input-character}
    to the current attribute\'s value. Otherwise, emit the [current
    input
    character](#current-input-character){#ambiguous-ampersand-state:current-input-character-2}
    as a character token.

U+003B SEMICOLON (;)
:   This is an
    [unknown-named-character-reference](#parse-error-unknown-named-character-reference){#ambiguous-ampersand-state:parse-error-unknown-named-character-reference}
    [parse
    error](#parse-errors){#ambiguous-ampersand-state:parse-errors}.
    [Reconsume](#reconsume){#ambiguous-ampersand-state:reconsume} in the
    [`return state`{#ambiguous-ampersand-state:return-state
    .variable}](#return-state).

Anything else
:   [Reconsume](#reconsume){#ambiguous-ampersand-state:reconsume-2} in
    the [`return state`{#ambiguous-ampersand-state:return-state-2
    .variable}](#return-state).

##### [13.2.5.75]{.secno} [Numeric character reference state]{.dfn}[](#numeric-character-reference-state){.self-link}

Set the
[`character reference code`{.variable}]{#character-reference-code .dfn}
to zero (0).

Consume the [next input
character](#next-input-character){#numeric-character-reference-state:next-input-character}:

U+0078 LATIN SMALL LETTER X\
U+0058 LATIN CAPITAL LETTER X
:   Append the [current input
    character](#current-input-character){#numeric-character-reference-state:current-input-character}
    to the
    [`temporary buffer`{#numeric-character-reference-state:temporary-buffer
    .variable}](#temporary-buffer). Switch to the [hexadecimal character
    reference start
    state](#hexadecimal-character-reference-start-state){#numeric-character-reference-state:hexadecimal-character-reference-start-state}.

Anything else
:   [Reconsume](#reconsume){#numeric-character-reference-state:reconsume}
    in the [decimal character reference start
    state](#decimal-character-reference-start-state){#numeric-character-reference-state:decimal-character-reference-start-state}.

##### [13.2.5.76]{.secno} [Hexadecimal character reference start state]{.dfn}[](#hexadecimal-character-reference-start-state){.self-link}

Consume the [next input
character](#next-input-character){#hexadecimal-character-reference-start-state:next-input-character}:

[ASCII hex digit](https://infra.spec.whatwg.org/#ascii-hex-digit){#hexadecimal-character-reference-start-state:ascii-hex-digits x-internal="ascii-hex-digits"}
:   [Reconsume](#reconsume){#hexadecimal-character-reference-start-state:reconsume}
    in the [hexadecimal character reference
    state](#hexadecimal-character-reference-state){#hexadecimal-character-reference-start-state:hexadecimal-character-reference-state}.

Anything else
:   This is an
    [absence-of-digits-in-numeric-character-reference](#parse-error-absence-of-digits-in-numeric-character-reference){#hexadecimal-character-reference-start-state:parse-error-absence-of-digits-in-numeric-character-reference}
    [parse
    error](#parse-errors){#hexadecimal-character-reference-start-state:parse-errors}.
    [Flush code points consumed as a character
    reference](#flush-code-points-consumed-as-a-character-reference){#hexadecimal-character-reference-start-state:flush-code-points-consumed-as-a-character-reference}.
    [Reconsume](#reconsume){#hexadecimal-character-reference-start-state:reconsume-2}
    in the
    [`return state`{#hexadecimal-character-reference-start-state:return-state
    .variable}](#return-state).

##### [13.2.5.77]{.secno} [Decimal character reference start state]{.dfn}[](#decimal-character-reference-start-state){.self-link}

Consume the [next input
character](#next-input-character){#decimal-character-reference-start-state:next-input-character}:

[ASCII digit](https://infra.spec.whatwg.org/#ascii-digit){#decimal-character-reference-start-state:ascii-digits x-internal="ascii-digits"}
:   [Reconsume](#reconsume){#decimal-character-reference-start-state:reconsume}
    in the [decimal character reference
    state](#decimal-character-reference-state){#decimal-character-reference-start-state:decimal-character-reference-state}.

Anything else
:   This is an
    [absence-of-digits-in-numeric-character-reference](#parse-error-absence-of-digits-in-numeric-character-reference){#decimal-character-reference-start-state:parse-error-absence-of-digits-in-numeric-character-reference}
    [parse
    error](#parse-errors){#decimal-character-reference-start-state:parse-errors}.
    [Flush code points consumed as a character
    reference](#flush-code-points-consumed-as-a-character-reference){#decimal-character-reference-start-state:flush-code-points-consumed-as-a-character-reference}.
    [Reconsume](#reconsume){#decimal-character-reference-start-state:reconsume-2}
    in the
    [`return state`{#decimal-character-reference-start-state:return-state
    .variable}](#return-state).

##### [13.2.5.78]{.secno} [Hexadecimal character reference state]{.dfn}[](#hexadecimal-character-reference-state){.self-link}

Consume the [next input
character](#next-input-character){#hexadecimal-character-reference-state:next-input-character}:

[ASCII digit](https://infra.spec.whatwg.org/#ascii-digit){#hexadecimal-character-reference-state:ascii-digits x-internal="ascii-digits"}
:   Multiply the
    [`character reference code`{#hexadecimal-character-reference-state:character-reference-code
    .variable}](#character-reference-code) by 16. Add a numeric version
    of the [current input
    character](#current-input-character){#hexadecimal-character-reference-state:current-input-character}
    (subtract 0x0030 from the character\'s code point) to the
    [`character reference code`{#hexadecimal-character-reference-state:character-reference-code-2
    .variable}](#character-reference-code).

[ASCII upper hex digit](https://infra.spec.whatwg.org/#ascii-upper-hex-digit){#hexadecimal-character-reference-state:uppercase-ascii-hex-digits x-internal="uppercase-ascii-hex-digits"}
:   Multiply the
    [`character reference code`{#hexadecimal-character-reference-state:character-reference-code-3
    .variable}](#character-reference-code) by 16. Add a numeric version
    of the [current input
    character](#current-input-character){#hexadecimal-character-reference-state:current-input-character-2}
    as a hexadecimal digit (subtract 0x0037 from the character\'s code
    point) to the
    [`character reference code`{#hexadecimal-character-reference-state:character-reference-code-4
    .variable}](#character-reference-code).

[ASCII lower hex digit](https://infra.spec.whatwg.org/#ascii-lower-hex-digit){#hexadecimal-character-reference-state:lowercase-ascii-hex-digits x-internal="lowercase-ascii-hex-digits"}
:   Multiply the
    [`character reference code`{#hexadecimal-character-reference-state:character-reference-code-5
    .variable}](#character-reference-code) by 16. Add a numeric version
    of the [current input
    character](#current-input-character){#hexadecimal-character-reference-state:current-input-character-3}
    as a hexadecimal digit (subtract 0x0057 from the character\'s code
    point) to the
    [`character reference code`{#hexadecimal-character-reference-state:character-reference-code-6
    .variable}](#character-reference-code).

U+003B SEMICOLON (;)
:   Switch to the [numeric character reference end
    state](#numeric-character-reference-end-state){#hexadecimal-character-reference-state:numeric-character-reference-end-state}.

Anything else
:   This is a
    [missing-semicolon-after-character-reference](#parse-error-missing-semicolon-after-character-reference){#hexadecimal-character-reference-state:parse-error-missing-semicolon-after-character-reference}
    [parse
    error](#parse-errors){#hexadecimal-character-reference-state:parse-errors}.
    [Reconsume](#reconsume){#hexadecimal-character-reference-state:reconsume}
    in the [numeric character reference end
    state](#numeric-character-reference-end-state){#hexadecimal-character-reference-state:numeric-character-reference-end-state-2}.

##### [13.2.5.79]{.secno} [Decimal character reference state]{.dfn}[](#decimal-character-reference-state){.self-link}

Consume the [next input
character](#next-input-character){#decimal-character-reference-state:next-input-character}:

[ASCII digit](https://infra.spec.whatwg.org/#ascii-digit){#decimal-character-reference-state:ascii-digits x-internal="ascii-digits"}
:   Multiply the
    [`character reference code`{#decimal-character-reference-state:character-reference-code
    .variable}](#character-reference-code) by 10. Add a numeric version
    of the [current input
    character](#current-input-character){#decimal-character-reference-state:current-input-character}
    (subtract 0x0030 from the character\'s code point) to the
    [`character reference code`{#decimal-character-reference-state:character-reference-code-2
    .variable}](#character-reference-code).

U+003B SEMICOLON (;)
:   Switch to the [numeric character reference end
    state](#numeric-character-reference-end-state){#decimal-character-reference-state:numeric-character-reference-end-state}.

Anything else
:   This is a
    [missing-semicolon-after-character-reference](#parse-error-missing-semicolon-after-character-reference){#decimal-character-reference-state:parse-error-missing-semicolon-after-character-reference}
    [parse
    error](#parse-errors){#decimal-character-reference-state:parse-errors}.
    [Reconsume](#reconsume){#decimal-character-reference-state:reconsume}
    in the [numeric character reference end
    state](#numeric-character-reference-end-state){#decimal-character-reference-state:numeric-character-reference-end-state-2}.

##### [13.2.5.80]{.secno} [Numeric character reference end state]{.dfn}[](#numeric-character-reference-end-state){.self-link}

Check the
[`character reference code`{#numeric-character-reference-end-state:character-reference-code
.variable}](#character-reference-code):

- If the number is 0x00, then this is a
  [null-character-reference](#parse-error-null-character-reference){#numeric-character-reference-end-state:parse-error-null-character-reference}
  [parse
  error](#parse-errors){#numeric-character-reference-end-state:parse-errors}.
  Set the
  [`character reference code`{#numeric-character-reference-end-state:character-reference-code-2
  .variable}](#character-reference-code) to 0xFFFD.

- If the number is greater than 0x10FFFF, then this is a
  [character-reference-outside-unicode-range](#parse-error-character-reference-outside-unicode-range){#numeric-character-reference-end-state:parse-error-character-reference-outside-unicode-range}
  [parse
  error](#parse-errors){#numeric-character-reference-end-state:parse-errors-2}.
  Set the
  [`character reference code`{#numeric-character-reference-end-state:character-reference-code-3
  .variable}](#character-reference-code) to 0xFFFD.

- If the number is a
  [surrogate](https://infra.spec.whatwg.org/#surrogate){#numeric-character-reference-end-state:surrogate
  x-internal="surrogate"}, then this is a
  [surrogate-character-reference](#parse-error-surrogate-character-reference){#numeric-character-reference-end-state:parse-error-surrogate-character-reference}
  [parse
  error](#parse-errors){#numeric-character-reference-end-state:parse-errors-3}.
  Set the
  [`character reference code`{#numeric-character-reference-end-state:character-reference-code-4
  .variable}](#character-reference-code) to 0xFFFD.

- If the number is a
  [noncharacter](https://infra.spec.whatwg.org/#noncharacter){#numeric-character-reference-end-state:noncharacter
  x-internal="noncharacter"}, then this is a
  [noncharacter-character-reference](#parse-error-noncharacter-character-reference){#numeric-character-reference-end-state:parse-error-noncharacter-character-reference}
  [parse
  error](#parse-errors){#numeric-character-reference-end-state:parse-errors-4}.

- If the number is 0x0D, or a
  [control](https://infra.spec.whatwg.org/#control){#numeric-character-reference-end-state:control
  x-internal="control"} that\'s not [ASCII
  whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#numeric-character-reference-end-state:space-characters
  x-internal="space-characters"}, then this is a
  [control-character-reference](#parse-error-control-character-reference){#numeric-character-reference-end-state:parse-error-control-character-reference}
  [parse
  error](#parse-errors){#numeric-character-reference-end-state:parse-errors-5}.
  If the number is one of the numbers in the first column of the
  following table, then find the row with that number in the first
  column, and set the
  [`character reference code`{#numeric-character-reference-end-state:character-reference-code-5
  .variable}](#character-reference-code) to the number in the second
  column of that row.

  Number

  Code point

  0x80

  0x20AC

  EURO SIGN (€)

  0x82

  0x201A

  SINGLE LOW-9 QUOTATION MARK (‚)

  0x83

  0x0192

  LATIN SMALL LETTER F WITH HOOK (ƒ)

  0x84

  0x201E

  DOUBLE LOW-9 QUOTATION MARK („)

  0x85

  0x2026

  HORIZONTAL ELLIPSIS (...)

  0x86

  0x2020

  DAGGER (†)

  0x87

  0x2021

  DOUBLE DAGGER (‡)

  0x88

  0x02C6

  MODIFIER LETTER CIRCUMFLEX ACCENT (ˆ)

  0x89

  0x2030

  PER MILLE SIGN (‰)

  0x8A

  0x0160

  LATIN CAPITAL LETTER S WITH CARON (Š)

  0x8B

  0x2039

  SINGLE LEFT-POINTING ANGLE QUOTATION MARK (‹)

  0x8C

  0x0152

  LATIN CAPITAL LIGATURE OE (Œ)

  0x8E

  0x017D

  LATIN CAPITAL LETTER Z WITH CARON (Ž)

  0x91

  0x2018

  LEFT SINGLE QUOTATION MARK (')

  0x92

  0x2019

  RIGHT SINGLE QUOTATION MARK (')

  0x93

  0x201C

  LEFT DOUBLE QUOTATION MARK (")

  0x94

  0x201D

  RIGHT DOUBLE QUOTATION MARK (")

  0x95

  0x2022

  BULLET (•)

  0x96

  0x2013

  EN DASH (--)

  0x97

  0x2014

  EM DASH (---)

  0x98

  0x02DC

  SMALL TILDE (˜)

  0x99

  0x2122

  TRADE MARK SIGN (™)

  0x9A

  0x0161

  LATIN SMALL LETTER S WITH CARON (š)

  0x9B

  0x203A

  SINGLE RIGHT-POINTING ANGLE QUOTATION MARK (›)

  0x9C

  0x0153

  LATIN SMALL LIGATURE OE (œ)

  0x9E

  0x017E

  LATIN SMALL LETTER Z WITH CARON (ž)

  0x9F

  0x0178

  LATIN CAPITAL LETTER Y WITH DIAERESIS (Ÿ)

Set the
[`temporary buffer`{#numeric-character-reference-end-state:temporary-buffer
.variable}](#temporary-buffer) to the empty string. Append a code point
equal to the
[`character reference code`{#numeric-character-reference-end-state:character-reference-code-6
.variable}](#character-reference-code) to the
[`temporary buffer`{#numeric-character-reference-end-state:temporary-buffer-2
.variable}](#temporary-buffer). [Flush code points consumed as a
character
reference](#flush-code-points-consumed-as-a-character-reference){#numeric-character-reference-end-state:flush-code-points-consumed-as-a-character-reference}.
Switch to the
[`return state`{#numeric-character-reference-end-state:return-state
.variable}](#return-state).

#### [13.2.6]{.secno} [Tree construction]{.dfn}[](#tree-construction){.self-link}

The input to the tree construction stage is a sequence of tokens from
the [tokenization](#tokenization){#tree-construction:tokenization}
stage. The tree construction stage is associated with a DOM
[`Document`{#tree-construction:document}](dom.html#document) object when
a parser is created. The \"output\" of this stage consists of
dynamically modifying or extending that document\'s DOM tree.

This specification does not define when an interactive user agent has to
render the
[`Document`{#tree-construction:document-2}](dom.html#document) so that
it is available to the user, or when it has to begin accepting user
input.

------------------------------------------------------------------------

As each token is emitted from the tokenizer, the user agent must follow
the appropriate steps from the following list, known as the [tree
construction dispatcher]{#tree-construction-dispatcher .dfn}:

If the [stack of open elements](#stack-of-open-elements){#tree-construction:stack-of-open-elements} is empty\
If the [adjusted current node](#adjusted-current-node){#tree-construction:adjusted-current-node} is an element in the [HTML namespace](https://infra.spec.whatwg.org/#html-namespace){#tree-construction:html-namespace-2 x-internal="html-namespace-2"}\
If the [adjusted current node](#adjusted-current-node){#tree-construction:adjusted-current-node-2} is a [MathML text integration point](#mathml-text-integration-point){#tree-construction:mathml-text-integration-point} and the token is a start tag whose tag name is neither \"mglyph\" nor \"malignmark\"\
If the [adjusted current node](#adjusted-current-node){#tree-construction:adjusted-current-node-3} is a [MathML text integration point](#mathml-text-integration-point){#tree-construction:mathml-text-integration-point-2} and the token is a character token\
If the [adjusted current node](#adjusted-current-node){#tree-construction:adjusted-current-node-4} is a [MathML `annotation-xml`](https://w3c.github.io/mathml-core/#dfn-annotation-xml){#tree-construction:mathml-annotation-xml x-internal="mathml-annotation-xml"} element and the token is a start tag whose tag name is \"svg\"\
If the [adjusted current node](#adjusted-current-node){#tree-construction:adjusted-current-node-5} is an [HTML integration point](#html-integration-point){#tree-construction:html-integration-point} and the token is a start tag\
If the [adjusted current node](#adjusted-current-node){#tree-construction:adjusted-current-node-6} is an [HTML integration point](#html-integration-point){#tree-construction:html-integration-point-2} and the token is a character token\
If the token is an end-of-file token
:   Process the token according to the rules given in the section
    corresponding to the current [insertion
    mode](#insertion-mode){#tree-construction:insertion-mode} in HTML
    content.

Otherwise
:   Process the token according to the rules given in the section for
    parsing tokens [in foreign
    content](#parsing-main-inforeign){#tree-construction:parsing-main-inforeign}.

The [next token]{#next-token .dfn} is the token that is about to be
processed by the [tree construction
dispatcher](#tree-construction-dispatcher){#tree-construction:tree-construction-dispatcher}
(even if the token is subsequently just ignored).

A node is a [MathML text integration
point]{#mathml-text-integration-point .dfn} if it is one of the
following elements:

- A [MathML
  `mi`](https://w3c.github.io/mathml-core/#the-mi-element){#tree-construction:mathml-mi
  x-internal="mathml-mi"} element
- A [MathML
  `mo`](https://w3c.github.io/mathml-core/#operator-fence-separator-or-accent-mo){#tree-construction:mathml-mo
  x-internal="mathml-mo"} element
- A [MathML
  `mn`](https://w3c.github.io/mathml-core/#number-mn){#tree-construction:mathml-mn
  x-internal="mathml-mn"} element
- A [MathML
  `ms`](https://w3c.github.io/mathml-core/#string-literal-ms){#tree-construction:mathml-ms
  x-internal="mathml-ms"} element
- A [MathML
  `mtext`](https://w3c.github.io/mathml-core/#text-mtext){#tree-construction:mathml-mtext
  x-internal="mathml-mtext"} element

A node is an [HTML integration point]{#html-integration-point .dfn} if
it is one of the following elements:

- A [MathML
  `annotation-xml`](https://w3c.github.io/mathml-core/#dfn-annotation-xml){#tree-construction:mathml-annotation-xml-2
  x-internal="mathml-annotation-xml"} element whose start tag token had
  an attribute with the name \"encoding\" whose value was an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#tree-construction:ascii-case-insensitive
  x-internal="ascii-case-insensitive"} match for the string
  \"`text/html`\"
- A [MathML
  `annotation-xml`](https://w3c.github.io/mathml-core/#dfn-annotation-xml){#tree-construction:mathml-annotation-xml-3
  x-internal="mathml-annotation-xml"} element whose start tag token had
  an attribute with the name \"encoding\" whose value was an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#tree-construction:ascii-case-insensitive-2
  x-internal="ascii-case-insensitive"} match for the string
  \"`application/xhtml+xml`\"
- An [SVG
  `foreignObject`](https://svgwg.org/svg2-draft/embedded.html#ForeignObjectElement){#tree-construction:svg-foreignobject
  x-internal="svg-foreignobject"} element
- An [SVG
  `desc`](https://svgwg.org/svg2-draft/struct.html#DescElement){#tree-construction:svg-desc
  x-internal="svg-desc"} element
- An [SVG
  `title`](https://svgwg.org/svg2-draft/struct.html#TitleElement){#tree-construction:svg-title
  x-internal="svg-title"} element

If the node in question is the
[`context`{#tree-construction:concept-frag-parse-context
.variable}](#concept-frag-parse-context) element passed to the [HTML
fragment parsing
algorithm](#html-fragment-parsing-algorithm){#tree-construction:html-fragment-parsing-algorithm},
then the start tag token for that element is the \"fake\" token created
during by that [HTML fragment parsing
algorithm](#html-fragment-parsing-algorithm){#tree-construction:html-fragment-parsing-algorithm-2}.

------------------------------------------------------------------------

Not all of the tag names mentioned below are conformant tag names in
this specification; many are included to handle legacy content. They
still form part of the algorithm that implementations are required to
implement to claim conformance.

The algorithm described below places no limit on the depth of the DOM
tree generated, or on the length of tag names, attribute names,
attribute values,
[`Text`{#tree-construction:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes, etc. While implementers are encouraged to [avoid arbitrary
limits](https://infra.spec.whatwg.org/#algorithm-limits), it is
recognized that practical concerns will likely force user agents to
impose nesting depth constraints.

##### [13.2.6.1]{.secno} Creating and inserting nodes[](#creating-and-inserting-nodes){.self-link}

While the parser is processing a token, it can enable or disable [foster
parenting]{#foster-parent .dfn}. This affects the following algorithm.

The [appropriate place for inserting a
node]{#appropriate-place-for-inserting-a-node .dfn}, optionally using a
particular *override target*, is the position in an element returned by
running the following steps:

1.  If there was an *override target* specified, then let
    `target`{.variable} be the *override target*.

    Otherwise, let `target`{.variable} be the [current
    node](#current-node){#creating-and-inserting-nodes:current-node}.

2.  Determine the `adjusted insertion location`{.variable} using the
    first matching steps from the following list:

    If [foster parenting](#foster-parent){#creating-and-inserting-nodes:foster-parent} is enabled and `target`{.variable} is a [`table`{#creating-and-inserting-nodes:the-table-element}](tables.html#the-table-element), [`tbody`{#creating-and-inserting-nodes:the-tbody-element}](tables.html#the-tbody-element), [`tfoot`{#creating-and-inserting-nodes:the-tfoot-element}](tables.html#the-tfoot-element), [`thead`{#creating-and-inserting-nodes:the-thead-element}](tables.html#the-thead-element), or [`tr`{#creating-and-inserting-nodes:the-tr-element}](tables.html#the-tr-element) element

    :   Foster parenting happens when content is misnested in tables.

        Run these substeps:

        1.  Let `last template`{.variable} be the last
            [`template`{#creating-and-inserting-nodes:the-template-element}](scripting.html#the-template-element)
            element in the [stack of open
            elements](#stack-of-open-elements){#creating-and-inserting-nodes:stack-of-open-elements},
            if any.

        2.  Let `last table`{.variable} be the last
            [`table`{#creating-and-inserting-nodes:the-table-element-2}](tables.html#the-table-element)
            element in the [stack of open
            elements](#stack-of-open-elements){#creating-and-inserting-nodes:stack-of-open-elements-2},
            if any.

        3.  If there is a `last template`{.variable} and either there is
            no `last table`{.variable}, or there is one, but
            `last template`{.variable} is lower (more recently added)
            than `last table`{.variable} in the [stack of open
            elements](#stack-of-open-elements){#creating-and-inserting-nodes:stack-of-open-elements-3},
            then: let `adjusted insertion location`{.variable} be inside
            `last template`{.variable}\'s [template
            contents](scripting.html#template-contents){#creating-and-inserting-nodes:template-contents},
            after its last child (if any), and abort these steps.

        4.  If there is no `last table`{.variable}, then let
            `adjusted insertion location`{.variable} be inside the first
            element in the [stack of open
            elements](#stack-of-open-elements){#creating-and-inserting-nodes:stack-of-open-elements-4}
            (the
            [`html`{#creating-and-inserting-nodes:the-html-element}](semantics.html#the-html-element)
            element), after its last child (if any), and abort these
            steps. ([fragment
            case](#fragment-case){#creating-and-inserting-nodes:fragment-case})

        5.  If `last table`{.variable} has a parent node, then let
            `adjusted insertion location`{.variable} be inside
            `last table`{.variable}\'s parent node, immediately before
            `last table`{.variable}, and abort these steps.

        6.  Let `previous element`{.variable} be the element immediately
            above `last table`{.variable} in the [stack of open
            elements](#stack-of-open-elements){#creating-and-inserting-nodes:stack-of-open-elements-5}.

        7.  Let `adjusted insertion location`{.variable} be inside
            `previous element`{.variable}, after its last child (if
            any).

        These steps are involved in part because it\'s possible for
        elements, the
        [`table`{#creating-and-inserting-nodes:the-table-element-3}](tables.html#the-table-element)
        element in this case in particular, to have been moved by a
        script around in the DOM, or indeed removed from the DOM
        entirely, after the element was inserted by the parser.

    Otherwise

    :   Let `adjusted insertion location`{.variable} be inside
        `target`{.variable}, after its last child (if any).

3.  If the `adjusted insertion location`{.variable} is inside a
    [`template`{#creating-and-inserting-nodes:the-template-element-2}](scripting.html#the-template-element)
    element, let it instead be inside the
    [`template`{#creating-and-inserting-nodes:the-template-element-3}](scripting.html#the-template-element)
    element\'s [template
    contents](scripting.html#template-contents){#creating-and-inserting-nodes:template-contents-2},
    after its last child (if any).

4.  Return the `adjusted insertion location`{.variable}.

------------------------------------------------------------------------

To [create an element for a token]{#create-an-element-for-the-token
.dfn}, given a token `token`{.variable}, a string `namespace`{.variable}
and a
[`Node`{#creating-and-inserting-nodes:node}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
object `intendedParent`{.variable}:

1.  If the [active speculative HTML
    parser](#active-speculative-html-parser){#creating-and-inserting-nodes:active-speculative-html-parser}
    is not null, then return the result of [creating a speculative mock
    element](#create-a-speculative-mock-element){#creating-and-inserting-nodes:create-a-speculative-mock-element}
    given `namespace`{.variable}, `token`{.variable}\'s tag name, and
    `token`{.variable}\'s attributes.

2.  Otherwise, optionally [create a speculative mock
    element](#create-a-speculative-mock-element){#creating-and-inserting-nodes:create-a-speculative-mock-element-2}
    given `namespace`{.variable}, `token`{.variable}\'s tag name, and
    `token`{.variable}\'s attributes.

    The result is not used. This step allows for a [speculative
    fetch](#speculative-fetch){#creating-and-inserting-nodes:speculative-fetch}
    to be initiated from non-speculative parsing. The fetch is still
    speculative at this point, because, for example, by the time the
    element is inserted, `intended parent`{.variable} might have been
    removed from the document.

3.  Let `document`{.variable} be `intendedParent`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#creating-and-inserting-nodes:node-document
    x-internal="node-document"}.

4.  Let `localName`{.variable} be `token`{.variable}\'s tag name.

5.  Let `is`{.variable} be the value of the
    \"[`is`{#creating-and-inserting-nodes:attr-is}](custom-elements.html#attr-is)\"
    attribute in `token`{.variable}, if such an attribute exists;
    otherwise null.

6.  Let `registry`{.variable} be the result of [looking up a custom
    element
    registry](custom-elements.html#look-up-a-custom-element-registry){#creating-and-inserting-nodes:look-up-a-custom-element-registry}
    given `intendedParent`{.variable}.

7.  Let `definition`{.variable} be the result of [looking up a custom
    element
    definition](custom-elements.html#look-up-a-custom-element-definition){#creating-and-inserting-nodes:look-up-a-custom-element-definition}
    given `registry`{.variable}, `namespace`{.variable},
    `localName`{.variable}, and `is`{.variable}.

8.  Let `willExecuteScript`{.variable} be true if
    `definition`{.variable} is non-null and the parser was not created
    as part of the [HTML fragment parsing
    algorithm](#html-fragment-parsing-algorithm){#creating-and-inserting-nodes:html-fragment-parsing-algorithm};
    otherwise false.

9.  If `willExecuteScript`{.variable} is true:

    1.  Increment `document`{.variable}\'s
        [throw-on-dynamic-markup-insertion
        counter](dynamic-markup-insertion.html#throw-on-dynamic-markup-insertion-counter){#creating-and-inserting-nodes:throw-on-dynamic-markup-insertion-counter}.

    2.  If the [JavaScript execution context
        stack](https://tc39.es/ecma262/#execution-context-stack){#creating-and-inserting-nodes:javascript-execution-context-stack
        x-internal="javascript-execution-context-stack"} is empty, then
        [perform a microtask
        checkpoint](webappapis.html#perform-a-microtask-checkpoint){#creating-and-inserting-nodes:perform-a-microtask-checkpoint}.

    3.  Push a new [element
        queue](custom-elements.html#element-queue){#creating-and-inserting-nodes:element-queue}
        onto `document`{.variable}\'s [relevant
        agent](webappapis.html#relevant-agent){#creating-and-inserting-nodes:relevant-agent}\'s
        [custom element reactions
        stack](custom-elements.html#custom-element-reactions-stack){#creating-and-inserting-nodes:custom-element-reactions-stack}.

10. Let `element`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#creating-and-inserting-nodes:create-an-element
    x-internal="create-an-element"} given `document`{.variable},
    `localName`{.variable}, `namespace`{.variable}, null,
    `is`{.variable}, `willExecuteScript`{.variable}, and
    `registry`{.variable}.

    This will cause [custom element
    constructors](custom-elements.html#custom-element-constructor){#creating-and-inserting-nodes:custom-element-constructor}
    to run, if `willExecuteScript`{.variable} is true. However, since we
    incremented the [throw-on-dynamic-markup-insertion
    counter](dynamic-markup-insertion.html#throw-on-dynamic-markup-insertion-counter){#creating-and-inserting-nodes:throw-on-dynamic-markup-insertion-counter-2},
    this cannot cause [new characters to be inserted into the
    tokenizer](dynamic-markup-insertion.html#dom-document-write){#creating-and-inserting-nodes:dom-document-write},
    or [the document to be blown
    away](dynamic-markup-insertion.html#dom-document-open){#creating-and-inserting-nodes:dom-document-open}.

11. [Append](https://dom.spec.whatwg.org/#concept-element-attributes-append){#creating-and-inserting-nodes:concept-element-attributes-append
    x-internal="concept-element-attributes-append"} each attribute in
    the given token to `element`{.variable}.

    This can [enqueue a custom element callback
    reaction](custom-elements.html#enqueue-a-custom-element-callback-reaction){#creating-and-inserting-nodes:enqueue-a-custom-element-callback-reaction}
    for the `attributeChangedCallback`, which might run immediately (in
    the next step).

    Even though the
    [`is`{#creating-and-inserting-nodes:attr-is-2}](custom-elements.html#attr-is)
    attribute governs the
    [creation](https://dom.spec.whatwg.org/#concept-create-element){#creating-and-inserting-nodes:create-an-element-2
    x-internal="create-an-element"} of a [customized built-in
    element](custom-elements.html#customized-built-in-element){#creating-and-inserting-nodes:customized-built-in-element},
    it is not present during the execution of the relevant [custom
    element
    constructor](custom-elements.html#custom-element-constructor){#creating-and-inserting-nodes:custom-element-constructor-2};
    it is appended in this step, along with all other attributes.

12. If `willExecuteScript`{.variable} is true:

    1.  Let `queue`{.variable} be the result of popping from
        `document`{.variable}\'s [relevant
        agent](webappapis.html#relevant-agent){#creating-and-inserting-nodes:relevant-agent-2}\'s
        [custom element reactions
        stack](custom-elements.html#custom-element-reactions-stack){#creating-and-inserting-nodes:custom-element-reactions-stack-2}.
        (This will be the same [element
        queue](custom-elements.html#element-queue){#creating-and-inserting-nodes:element-queue-2}
        as was pushed above.)

    2.  [Invoke custom element
        reactions](custom-elements.html#invoke-custom-element-reactions){#creating-and-inserting-nodes:invoke-custom-element-reactions}
        in `queue`{.variable}.

    3.  Decrement `document`{.variable}\'s
        [throw-on-dynamic-markup-insertion
        counter](dynamic-markup-insertion.html#throw-on-dynamic-markup-insertion-counter){#creating-and-inserting-nodes:throw-on-dynamic-markup-insertion-counter-3}.

13. If `element`{.variable} has an `xmlns` attribute *in the [XMLNS
    namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#creating-and-inserting-nodes:xmlns-namespace
    x-internal="xmlns-namespace"}* whose value is not exactly the same
    as the element\'s namespace, that is a [parse
    error](#parse-errors){#creating-and-inserting-nodes:parse-errors}.
    Similarly, if `element`{.variable} has an `xmlns:xlink` attribute in
    the [XMLNS
    namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#creating-and-inserting-nodes:xmlns-namespace-2
    x-internal="xmlns-namespace"} whose value is not the [XLink
    Namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace
    x-internal="xlink-namespace"}, that is a [parse
    error](#parse-errors){#creating-and-inserting-nodes:parse-errors-2}.

14. If `element`{.variable} is a [resettable
    element](forms.html#category-reset){#creating-and-inserting-nodes:category-reset}
    and not a [form-associated custom
    element](custom-elements.html#form-associated-custom-element){#creating-and-inserting-nodes:form-associated-custom-element},
    then invoke its [reset
    algorithm](form-control-infrastructure.html#concept-form-reset-control){#creating-and-inserting-nodes:concept-form-reset-control}.
    (This initializes the element\'s
    [value](form-control-infrastructure.html#concept-fe-value){#creating-and-inserting-nodes:concept-fe-value}
    and
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#creating-and-inserting-nodes:concept-fe-checked}
    based on the element\'s attributes.)

15. If `element`{.variable} is a [form-associated
    element](forms.html#form-associated-element){#creating-and-inserting-nodes:form-associated-element}
    and not a [form-associated custom
    element](custom-elements.html#form-associated-custom-element){#creating-and-inserting-nodes:form-associated-custom-element-2},
    the [`form` element
    pointer](#form-element-pointer){#creating-and-inserting-nodes:form-element-pointer}
    is not null, there is no
    [`template`{#creating-and-inserting-nodes:the-template-element-4}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#creating-and-inserting-nodes:stack-of-open-elements-6},
    `element`{.variable} is either not
    [listed](forms.html#category-listed){#creating-and-inserting-nodes:category-listed}
    or doesn\'t have a
    [`form`{#creating-and-inserting-nodes:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    attribute, and the `intendedParent`{.variable} is in the same
    [tree](https://dom.spec.whatwg.org/#concept-tree){#creating-and-inserting-nodes:tree
    x-internal="tree"} as the element pointed to by the [`form` element
    pointer](#form-element-pointer){#creating-and-inserting-nodes:form-element-pointer-2},
    then
    [associate](form-control-infrastructure.html#concept-form-association){#creating-and-inserting-nodes:concept-form-association}
    `element`{.variable} with the
    [`form`{#creating-and-inserting-nodes:the-form-element}](forms.html#the-form-element)
    element pointed to by the [`form` element
    pointer](#form-element-pointer){#creating-and-inserting-nodes:form-element-pointer-3}
    and set `element`{.variable}\'s [parser inserted
    flag](form-control-infrastructure.html#parser-inserted-flag){#creating-and-inserting-nodes:parser-inserted-flag}.

16. Return `element`{.variable}.

------------------------------------------------------------------------

To [insert an element at the adjusted insertion
location]{#insert-an-element-at-the-adjusted-insertion-location .dfn}
with an element `element`{.variable}:

1.  Let the `adjusted insertion location`{.variable} be the [appropriate
    place for inserting a
    node](#appropriate-place-for-inserting-a-node){#creating-and-inserting-nodes:appropriate-place-for-inserting-a-node}.

2.  If it is not possible to insert `element`{.variable} at the
    `adjusted insertion location`{.variable}, abort these steps.

3.  If the parser was not created as part of the [HTML fragment parsing
    algorithm](#html-fragment-parsing-algorithm){#creating-and-inserting-nodes:html-fragment-parsing-algorithm-2},
    then push a new [element
    queue](custom-elements.html#element-queue){#creating-and-inserting-nodes:element-queue-3}
    onto `element`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#creating-and-inserting-nodes:relevant-agent-3}\'s
    [custom element reactions
    stack](custom-elements.html#custom-element-reactions-stack){#creating-and-inserting-nodes:custom-element-reactions-stack-3}.

4.  Insert `element`{.variable} at the
    `adjusted insertion location`{.variable}.

5.  If the parser was not created as part of the [HTML fragment parsing
    algorithm](#html-fragment-parsing-algorithm){#creating-and-inserting-nodes:html-fragment-parsing-algorithm-3},
    then pop the [element
    queue](custom-elements.html#element-queue){#creating-and-inserting-nodes:element-queue-4}
    from `element`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#creating-and-inserting-nodes:relevant-agent-4}\'s
    [custom element reactions
    stack](custom-elements.html#custom-element-reactions-stack){#creating-and-inserting-nodes:custom-element-reactions-stack-4},
    and [invoke custom element
    reactions](custom-elements.html#invoke-custom-element-reactions){#creating-and-inserting-nodes:invoke-custom-element-reactions-2}
    in that queue.

If the `adjusted insertion location`{.variable} cannot accept more
elements, e.g., because it\'s a
[`Document`{#creating-and-inserting-nodes:document}](dom.html#document)
that already has an element child, then `element`{.variable} is dropped
on the floor.

To [insert a foreign element]{#insert-a-foreign-element .dfn}, given a
token `token`{.variable}, a string `namespace`{.variable}, and a boolean
`onlyAddToElementStack`{.variable}:

1.  Let the `adjustedInsertionLocation`{.variable} be the [appropriate
    place for inserting a
    node](#appropriate-place-for-inserting-a-node){#creating-and-inserting-nodes:appropriate-place-for-inserting-a-node-2}.

2.  Let `element`{.variable} be the result of [creating an element for
    the
    token](#create-an-element-for-the-token){#creating-and-inserting-nodes:create-an-element-for-the-token}
    given `token`{.variable}, `namespace`{.variable}, and the element in
    which the `adjustedInsertionLocation`{.variable} finds itself.

3.  If `onlyAddToElementStack`{.variable} is false, then run [insert an
    element at the adjusted insertion
    location](#insert-an-element-at-the-adjusted-insertion-location){#creating-and-inserting-nodes:insert-an-element-at-the-adjusted-insertion-location}
    with `element`{.variable}.

4.  Push `element`{.variable} onto the [stack of open
    elements](#stack-of-open-elements){#creating-and-inserting-nodes:stack-of-open-elements-7}
    so that it is the new [current
    node](#current-node){#creating-and-inserting-nodes:current-node-2}.

5.  Return `element`{.variable}.

To [insert an HTML element]{#insert-an-html-element .dfn} given a token
`token`{.variable}: [insert a foreign
element](#insert-a-foreign-element){#creating-and-inserting-nodes:insert-a-foreign-element}
given `token`{.variable}, the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#creating-and-inserting-nodes:html-namespace-2
x-internal="html-namespace-2"}, and false.

------------------------------------------------------------------------

When the steps below require the user agent to [adjust MathML
attributes]{#adjust-mathml-attributes .dfn} for a token, then, if the
token has an attribute named `definitionurl`, change its name to
`definitionURL` (note the case difference).

When the steps below require the user agent to [adjust SVG
attributes]{#adjust-svg-attributes .dfn} for a token, then, for each
attribute on the token whose attribute name is one of the ones in the
first column of the following table, change the attribute\'s name to the
name given in the corresponding cell in the second column. (This fixes
the case of SVG attributes that are not all lowercase.)

Attribute name on token

Attribute name on element

`attributename`

`attributeName`

`attributetype`

`attributeType`

`basefrequency`

`baseFrequency`

`baseprofile`

`baseProfile`

`calcmode`

`calcMode`

`clippathunits`

`clipPathUnits`

`diffuseconstant`

`diffuseConstant`

`edgemode`

`edgeMode`

`filterunits`

`filterUnits`

`glyphref`

`glyphRef`

`gradienttransform`

`gradientTransform`

`gradientunits`

`gradientUnits`

`kernelmatrix`

`kernelMatrix`

`kernelunitlength`

`kernelUnitLength`

`keypoints`

`keyPoints`

`keysplines`

`keySplines`

`keytimes`

`keyTimes`

`lengthadjust`

`lengthAdjust`

`limitingconeangle`

`limitingConeAngle`

`markerheight`

`markerHeight`

`markerunits`

`markerUnits`

`markerwidth`

`markerWidth`

`maskcontentunits`

`maskContentUnits`

`maskunits`

`maskUnits`

`numoctaves`

`numOctaves`

`pathlength`

`pathLength`

`patterncontentunits`

`patternContentUnits`

`patterntransform`

`patternTransform`

`patternunits`

`patternUnits`

`pointsatx`

`pointsAtX`

`pointsaty`

`pointsAtY`

`pointsatz`

`pointsAtZ`

`preservealpha`

`preserveAlpha`

`preserveaspectratio`

`preserveAspectRatio`

`primitiveunits`

`primitiveUnits`

`refx`

`refX`

`refy`

`refY`

`repeatcount`

`repeatCount`

`repeatdur`

`repeatDur`

`requiredextensions`

`requiredExtensions`

`requiredfeatures`

`requiredFeatures`

`specularconstant`

`specularConstant`

`specularexponent`

`specularExponent`

`spreadmethod`

`spreadMethod`

`startoffset`

`startOffset`

`stddeviation`

`stdDeviation`

`stitchtiles`

`stitchTiles`

`surfacescale`

`surfaceScale`

`systemlanguage`

`systemLanguage`

`tablevalues`

`tableValues`

`targetx`

`targetX`

`targety`

`targetY`

`textlength`

`textLength`

`viewbox`

`viewBox`

`viewtarget`

`viewTarget`

`xchannelselector`

`xChannelSelector`

`ychannelselector`

`yChannelSelector`

`zoomandpan`

`zoomAndPan`

When the steps below require the user agent to [adjust foreign
attributes]{#adjust-foreign-attributes .dfn} for a token, then, if any
of the attributes on the token match the strings given in the first
column of the following table, let the attribute be a namespaced
attribute, with the prefix being the string given in the corresponding
cell in the second column, the local name being the string given in the
corresponding cell in the third column, and the namespace being the
namespace given in the corresponding cell in the fourth column. (This
fixes the use of namespaced attributes, in particular [`lang` attributes
in the XML
namespace](https://www.w3.org/TR/xml/#sec-lang-tag){#creating-and-inserting-nodes:attr-xml-lang
x-internal="attr-xml-lang"}.)

Attribute name

Prefix

Local name

Namespace

`xlink:actuate`

`xlink`

`actuate`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace-2
x-internal="xlink-namespace"}

`xlink:arcrole`

`xlink`

`arcrole`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace-3
x-internal="xlink-namespace"}

`xlink:href`

`xlink`

`href`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace-4
x-internal="xlink-namespace"}

`xlink:role`

`xlink`

`role`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace-5
x-internal="xlink-namespace"}

`xlink:show`

`xlink`

`show`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace-6
x-internal="xlink-namespace"}

`xlink:title`

`xlink`

`title`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace-7
x-internal="xlink-namespace"}

`xlink:type`

`xlink`

`type`

[XLink
namespace](https://infra.spec.whatwg.org/#xlink-namespace){#creating-and-inserting-nodes:xlink-namespace-8
x-internal="xlink-namespace"}

`xml:lang`

`xml`

`lang`

[XML
namespace](https://infra.spec.whatwg.org/#xml-namespace){#creating-and-inserting-nodes:xml-namespace
x-internal="xml-namespace"}

`xml:space`

`xml`

`space`

[XML
namespace](https://infra.spec.whatwg.org/#xml-namespace){#creating-and-inserting-nodes:xml-namespace-2
x-internal="xml-namespace"}

`xmlns`

(none)

`xmlns`

[XMLNS
namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#creating-and-inserting-nodes:xmlns-namespace-3
x-internal="xmlns-namespace"}

`xmlns:xlink`

`xmlns`

`xlink`

[XMLNS
namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#creating-and-inserting-nodes:xmlns-namespace-4
x-internal="xmlns-namespace"}

------------------------------------------------------------------------

When the steps below require the user agent to [insert a
character]{#insert-a-character .dfn} while processing a token, the user
agent must run the following steps:

1.  Let `data`{.variable} be the characters passed to the algorithm, or,
    if no characters were explicitly specified, the character of the
    character token being processed.

2.  Let the `adjusted insertion location`{.variable} be the [appropriate
    place for inserting a
    node](#appropriate-place-for-inserting-a-node){#creating-and-inserting-nodes:appropriate-place-for-inserting-a-node-3}.

3.  If the `adjusted insertion location`{.variable} is in a
    [`Document`{#creating-and-inserting-nodes:document-2}](dom.html#document)
    node, then return.

    The DOM will not let
    [`Document`{#creating-and-inserting-nodes:document-3}](dom.html#document)
    nodes have
    [`Text`{#creating-and-inserting-nodes:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node children, so they are dropped on the floor.

4.  If there is a
    [`Text`{#creating-and-inserting-nodes:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node immediately before the
    `adjusted insertion location`{.variable}, then append
    `data`{.variable} to that
    [`Text`{#creating-and-inserting-nodes:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node\'s
    [data](https://dom.spec.whatwg.org/#concept-cd-data){#creating-and-inserting-nodes:concept-cd-data
    x-internal="concept-cd-data"}.

    Otherwise, create a new
    [`Text`{#creating-and-inserting-nodes:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node whose
    [data](https://dom.spec.whatwg.org/#concept-cd-data){#creating-and-inserting-nodes:concept-cd-data-2
    x-internal="concept-cd-data"} is `data`{.variable} and whose [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#creating-and-inserting-nodes:node-document-2
    x-internal="node-document"} is the same as that of the element in
    which the `adjusted insertion location`{.variable} finds itself, and
    insert the newly created node at the
    `adjusted insertion location`{.variable}.

::: example
Here are some sample inputs to the parser and the corresponding number
of
[`Text`{#creating-and-inserting-nodes:text-5}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes that they result in, assuming a user agent that executes scripts.

Input

Number of
[`Text`{#creating-and-inserting-nodes:text-6}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes

``` html
A<script>
var script = document.getElementsByTagName('script')[0];
document.body.removeChild(script);
</script>B
```

One
[`Text`{#creating-and-inserting-nodes:text-7}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node in the document, containing \"AB\".

``` html
A<script>
var text = document.createTextNode('B');
document.body.appendChild(text);
</script>C
```

Three
[`Text`{#creating-and-inserting-nodes:text-8}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes; \"A\" before the script, the script\'s contents, and \"BC\" after
the script (the parser appends to the
[`Text`{#creating-and-inserting-nodes:text-9}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node created by the script).

``` html
A<script>
var text = document.getElementsByTagName('script')[0].firstChild;
text.data = 'B';
document.body.appendChild(text);
</script>C
```

Two adjacent
[`Text`{#creating-and-inserting-nodes:text-10}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes in the document, containing \"A\" and \"BC\".

``` html
A<table>B<tr>C</tr>D</table>
```

One
[`Text`{#creating-and-inserting-nodes:text-11}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node before the table, containing \"ABCD\". (This is caused by [foster
parenting](#foster-parent){#creating-and-inserting-nodes:foster-parent-2}.)

``` html
A<table><tr> B</tr> C</table>
```

One
[`Text`{#creating-and-inserting-nodes:text-12}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node before the table, containing \"A B C\" (A-space-B-space-C). (This
is caused by [foster
parenting](#foster-parent){#creating-and-inserting-nodes:foster-parent-3}.)

``` html
A<table><tr> B</tr> </em>C</table>
```

One
[`Text`{#creating-and-inserting-nodes:text-13}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node before the table, containing \"A BC\" (A-space-B-C), and one
[`Text`{#creating-and-inserting-nodes:text-14}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node inside the table (as a child of a
[`tbody`{#creating-and-inserting-nodes:the-tbody-element-2}](tables.html#the-tbody-element))
with a single space character. (Space characters separated from
non-space characters by non-character tokens are not affected by [foster
parenting](#foster-parent){#creating-and-inserting-nodes:foster-parent-4},
even if those other tokens then get ignored.)
:::

------------------------------------------------------------------------

When the steps below require the user agent to [insert a
comment]{#insert-a-comment .dfn} while processing a comment token,
optionally with an explicit insertion position `position`{.variable},
the user agent must run the following steps:

1.  Let `data`{.variable} be the data given in the comment token being
    processed.

2.  If `position`{.variable} was specified, then let the
    `adjusted insertion location`{.variable} be `position`{.variable}.
    Otherwise, let `adjusted insertion location`{.variable} be the
    [appropriate place for inserting a
    node](#appropriate-place-for-inserting-a-node){#creating-and-inserting-nodes:appropriate-place-for-inserting-a-node-4}.

3.  Create a
    [`Comment`{#creating-and-inserting-nodes:comment-2}](https://dom.spec.whatwg.org/#interface-comment){x-internal="comment-2"}
    node whose `data` attribute is set to `data`{.variable} and whose
    [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#creating-and-inserting-nodes:node-document-3
    x-internal="node-document"} is the same as that of the node in which
    the `adjusted insertion location`{.variable} finds itself.

4.  Insert the newly created node at the
    `adjusted insertion location`{.variable}.

##### [13.2.6.2]{.secno} Parsing elements that contain only text[](#parsing-elements-that-contain-only-text){.self-link}

The [generic raw text element parsing
algorithm]{#generic-raw-text-element-parsing-algorithm .dfn} and the
[generic RCDATA element parsing
algorithm]{#generic-rcdata-element-parsing-algorithm .dfn} consist of
the following steps. These algorithms are always invoked in response to
a start tag token.

1.  [Insert an HTML
    element](#insert-an-html-element){#parsing-elements-that-contain-only-text:insert-an-html-element}
    for the token.

2.  If the algorithm that was invoked is the [generic raw text element
    parsing
    algorithm](#generic-raw-text-element-parsing-algorithm){#parsing-elements-that-contain-only-text:generic-raw-text-element-parsing-algorithm},
    switch the tokenizer to the [RAWTEXT
    state](#rawtext-state){#parsing-elements-that-contain-only-text:rawtext-state};
    otherwise the algorithm invoked was the [generic RCDATA element
    parsing
    algorithm](#generic-rcdata-element-parsing-algorithm){#parsing-elements-that-contain-only-text:generic-rcdata-element-parsing-algorithm},
    switch the tokenizer to the [RCDATA
    state](#rcdata-state){#parsing-elements-that-contain-only-text:rcdata-state}.

3.  Set the [original insertion
    mode](#original-insertion-mode){#parsing-elements-that-contain-only-text:original-insertion-mode}
    to the current [insertion
    mode](#insertion-mode){#parsing-elements-that-contain-only-text:insertion-mode}.

4.  Then, switch the [insertion
    mode](#insertion-mode){#parsing-elements-that-contain-only-text:insertion-mode-2}
    to
    \"[text](#parsing-main-incdata){#parsing-elements-that-contain-only-text:parsing-main-incdata}\".

##### [13.2.6.3]{.secno} Closing elements that have implied end tags[](#closing-elements-that-have-implied-end-tags){.self-link}

When the steps below require the UA to [generate implied end
tags]{#generate-implied-end-tags .dfn}, then, while the [current
node](#current-node){#closing-elements-that-have-implied-end-tags:current-node}
is a
[`dd`{#closing-elements-that-have-implied-end-tags:the-dd-element}](grouping-content.html#the-dd-element)
element, a
[`dt`{#closing-elements-that-have-implied-end-tags:the-dt-element}](grouping-content.html#the-dt-element)
element, an
[`li`{#closing-elements-that-have-implied-end-tags:the-li-element}](grouping-content.html#the-li-element)
element, an
[`optgroup`{#closing-elements-that-have-implied-end-tags:the-optgroup-element}](form-elements.html#the-optgroup-element)
element, an
[`option`{#closing-elements-that-have-implied-end-tags:the-option-element}](form-elements.html#the-option-element)
element, a
[`p`{#closing-elements-that-have-implied-end-tags:the-p-element}](grouping-content.html#the-p-element)
element, an
[`rb`{#closing-elements-that-have-implied-end-tags:rb}](obsolete.html#rb)
element, an
[`rp`{#closing-elements-that-have-implied-end-tags:the-rp-element}](text-level-semantics.html#the-rp-element)
element, an
[`rt`{#closing-elements-that-have-implied-end-tags:the-rt-element}](text-level-semantics.html#the-rt-element)
element, or an
[`rtc`{#closing-elements-that-have-implied-end-tags:rtc}](obsolete.html#rtc)
element, the UA must pop the [current
node](#current-node){#closing-elements-that-have-implied-end-tags:current-node-2}
off the [stack of open
elements](#stack-of-open-elements){#closing-elements-that-have-implied-end-tags:stack-of-open-elements}.

If a step requires the UA to generate implied end tags but lists an
element to exclude from the process, then the UA must perform the above
steps as if that element was not in the above list.

When the steps below require the UA to [generate all implied end tags
thoroughly]{#generate-all-implied-end-tags-thoroughly .dfn}, then, while
the [current
node](#current-node){#closing-elements-that-have-implied-end-tags:current-node-3}
is a
[`caption`{#closing-elements-that-have-implied-end-tags:the-caption-element}](tables.html#the-caption-element)
element, a
[`colgroup`{#closing-elements-that-have-implied-end-tags:the-colgroup-element}](tables.html#the-colgroup-element)
element, a
[`dd`{#closing-elements-that-have-implied-end-tags:the-dd-element-2}](grouping-content.html#the-dd-element)
element, a
[`dt`{#closing-elements-that-have-implied-end-tags:the-dt-element-2}](grouping-content.html#the-dt-element)
element, an
[`li`{#closing-elements-that-have-implied-end-tags:the-li-element-2}](grouping-content.html#the-li-element)
element, an
[`optgroup`{#closing-elements-that-have-implied-end-tags:the-optgroup-element-2}](form-elements.html#the-optgroup-element)
element, an
[`option`{#closing-elements-that-have-implied-end-tags:the-option-element-2}](form-elements.html#the-option-element)
element, a
[`p`{#closing-elements-that-have-implied-end-tags:the-p-element-2}](grouping-content.html#the-p-element)
element, an
[`rb`{#closing-elements-that-have-implied-end-tags:rb-2}](obsolete.html#rb)
element, an
[`rp`{#closing-elements-that-have-implied-end-tags:the-rp-element-2}](text-level-semantics.html#the-rp-element)
element, an
[`rt`{#closing-elements-that-have-implied-end-tags:the-rt-element-2}](text-level-semantics.html#the-rt-element)
element, an
[`rtc`{#closing-elements-that-have-implied-end-tags:rtc-2}](obsolete.html#rtc)
element, a
[`tbody`{#closing-elements-that-have-implied-end-tags:the-tbody-element}](tables.html#the-tbody-element)
element, a
[`td`{#closing-elements-that-have-implied-end-tags:the-td-element}](tables.html#the-td-element)
element, a
[`tfoot`{#closing-elements-that-have-implied-end-tags:the-tfoot-element}](tables.html#the-tfoot-element)
element, a
[`th`{#closing-elements-that-have-implied-end-tags:the-th-element}](tables.html#the-th-element)
element, a
[`thead`{#closing-elements-that-have-implied-end-tags:the-thead-element}](tables.html#the-thead-element)
element, or a
[`tr`{#closing-elements-that-have-implied-end-tags:the-tr-element}](tables.html#the-tr-element)
element, the UA must pop the [current
node](#current-node){#closing-elements-that-have-implied-end-tags:current-node-4}
off the [stack of open
elements](#stack-of-open-elements){#closing-elements-that-have-implied-end-tags:stack-of-open-elements-2}.

##### [13.2.6.4]{.secno} The rules for parsing tokens in HTML content[](#parsing-main-inhtml){.self-link} {#parsing-main-inhtml}

###### [13.2.6.4.1]{.secno} The \"[initial]{.dfn}\" insertion mode[](#the-initial-insertion-mode){.self-link}

A [`Document`{#the-initial-insertion-mode:document}](dom.html#document)
object has an associated [parser cannot change the mode
flag]{#parser-cannot-change-the-mode-flag .dfn} (a boolean). It is
initially false.

When the user agent is to apply the rules for the
\"[initial](#the-initial-insertion-mode){#the-initial-insertion-mode:the-initial-insertion-mode}\"
[insertion
mode](#insertion-mode){#the-initial-insertion-mode:insertion-mode}, the
user agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   Ignore the token.

A comment token

:   [Insert a
    comment](#insert-a-comment){#the-initial-insertion-mode:insert-a-comment}
    as the last child of the
    [`Document`{#the-initial-insertion-mode:document-2}](dom.html#document)
    object.

A DOCTYPE token

:   If the DOCTYPE token\'s name is not \"`html`\", or the token\'s
    public identifier is not missing, or the token\'s system identifier
    is neither missing nor
    \"[`about:legacy-compat`{#the-initial-insertion-mode:about:legacy-compat}](urls-and-fetching.html#about:legacy-compat)\",
    then there is a [parse
    error](#parse-errors){#the-initial-insertion-mode:parse-errors}.

    Append a
    [`DocumentType`{#the-initial-insertion-mode:documenttype}](https://dom.spec.whatwg.org/#interface-documenttype){x-internal="documenttype"}
    node to the
    [`Document`{#the-initial-insertion-mode:document-3}](dom.html#document)
    node, with its
    [name](https://dom.spec.whatwg.org/#concept-doctype-name){#the-initial-insertion-mode:concept-doctype-name
    x-internal="concept-doctype-name"} set to the name given in the
    DOCTYPE token, or the empty string if the name was missing; its
    [public
    ID](https://dom.spec.whatwg.org/#concept-doctype-publicid){#the-initial-insertion-mode:concept-doctype-publicid
    x-internal="concept-doctype-publicid"} set to the public identifier
    given in the DOCTYPE token, or the empty string if the public
    identifier was missing; and its [system
    ID](https://dom.spec.whatwg.org/#concept-doctype-systemid){#the-initial-insertion-mode:concept-doctype-systemid
    x-internal="concept-doctype-systemid"} set to the system identifier
    given in the DOCTYPE token, or the empty string if the system
    identifier was missing.

    This also ensures that the
    [`DocumentType`{#the-initial-insertion-mode:documenttype-2}](https://dom.spec.whatwg.org/#interface-documenttype){x-internal="documenttype"}
    node is returned as the value of the
    [`doctype`{#the-initial-insertion-mode:dom-document-doctype}](https://dom.spec.whatwg.org/#dom-document-doctype){x-internal="dom-document-doctype"}
    attribute of the
    [`Document`{#the-initial-insertion-mode:document-4}](dom.html#document)
    object.

    Then, if the document is *not* [an `iframe` `srcdoc`
    document](iframe-embed-object.html#an-iframe-srcdoc-document){#the-initial-insertion-mode:an-iframe-srcdoc-document},
    and the [parser cannot change the mode
    flag](#parser-cannot-change-the-mode-flag){#the-initial-insertion-mode:parser-cannot-change-the-mode-flag}
    is false, and the DOCTYPE token matches one of the conditions in the
    following list, then set the
    [`Document`{#the-initial-insertion-mode:document-5}](dom.html#document)
    to [quirks
    mode](https://dom.spec.whatwg.org/#concept-document-quirks){#the-initial-insertion-mode:quirks-mode
    x-internal="quirks-mode"}:

    - The *[force-quirks flag](#force-quirks-flag)* is set to *on*.
    - The name is not \"`html`\".
    - The public identifier is set to:
      \"`-//W3O//DTD W3 HTML Strict 3.0//EN//`\"
    - The public identifier is set to:
      \"`-/W3C/DTD HTML 4.0 Transitional/EN`\"
    - The public identifier is set to: \"`HTML`\"
    - The system identifier is set to:
      \"`http://www.ibm.com/data/dtd/v11/ibmxhtml1-transitional.dtd`\"
    - The public identifier starts with:
      \"`+//Silmaril//dtd html Pro v0r11 19970101//`\"
    - The public identifier starts with:
      \"`-//AS//DTD HTML 3.0 asWedit + extensions//`\"
    - The public identifier starts with:
      \"`-//AdvaSoft Ltd//DTD HTML 3.0 asWedit + extensions//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML 2.0 Level 1//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML 2.0 Level 2//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML 2.0 Strict Level 1//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML 2.0 Strict Level 2//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML 2.0 Strict//`\"
    - The public identifier starts with: \"`-//IETF//DTD HTML 2.0//`\"
    - The public identifier starts with: \"`-//IETF//DTD HTML 2.1E//`\"
    - The public identifier starts with: \"`-//IETF//DTD HTML 3.0//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML 3.2 Final//`\"
    - The public identifier starts with: \"`-//IETF//DTD HTML 3.2//`\"
    - The public identifier starts with: \"`-//IETF//DTD HTML 3//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Level 0//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Level 1//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Level 2//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Level 3//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Strict Level 0//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Strict Level 1//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Strict Level 2//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Strict Level 3//`\"
    - The public identifier starts with:
      \"`-//IETF//DTD HTML Strict//`\"
    - The public identifier starts with: \"`-//IETF//DTD HTML//`\"
    - The public identifier starts with:
      \"`-//Metrius//DTD Metrius Presentational//`\"
    - The public identifier starts with:
      \"`-//Microsoft//DTD Internet Explorer 2.0 HTML Strict//`\"
    - The public identifier starts with:
      \"`-//Microsoft//DTD Internet Explorer 2.0 HTML//`\"
    - The public identifier starts with:
      \"`-//Microsoft//DTD Internet Explorer 2.0 Tables//`\"
    - The public identifier starts with:
      \"`-//Microsoft//DTD Internet Explorer 3.0 HTML Strict//`\"
    - The public identifier starts with:
      \"`-//Microsoft//DTD Internet Explorer 3.0 HTML//`\"
    - The public identifier starts with:
      \"`-//Microsoft//DTD Internet Explorer 3.0 Tables//`\"
    - The public identifier starts with:
      \"`-//Netscape Comm. Corp.//DTD HTML//`\"
    - The public identifier starts with:
      \"`-//Netscape Comm. Corp.//DTD Strict HTML//`\"
    - The public identifier starts with:
      \"`-//O'Reilly and Associates//DTD HTML 2.0//`\"
    - The public identifier starts with:
      \"`-//O'Reilly and Associates//DTD HTML Extended 1.0//`\"
    - The public identifier starts with:
      \"`-//O'Reilly and Associates//DTD HTML Extended Relaxed 1.0//`\"
    - The public identifier starts with:
      \"`-//SQ//DTD HTML 2.0 HoTMetaL + extensions//`\"
    - The public identifier starts with:
      \"`-//SoftQuad Software//DTD HoTMetaL PRO 6.0::19990601::extensions to HTML 4.0//`\"
    - The public identifier starts with:
      \"`-//SoftQuad//DTD HoTMetaL PRO 4.0::19971010::extensions to HTML 4.0//`\"
    - The public identifier starts with:
      \"`-//Spyglass//DTD HTML 2.0 Extended//`\"
    - The public identifier starts with:
      \"`-//Sun Microsystems Corp.//DTD HotJava HTML//`\"
    - The public identifier starts with:
      \"`-//Sun Microsystems Corp.//DTD HotJava Strict HTML//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML 3 1995-03-24//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML 3.2 Draft//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML 3.2 Final//`\"
    - The public identifier starts with: \"`-//W3C//DTD HTML 3.2//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML 3.2S Draft//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML 4.0 Frameset//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML 4.0 Transitional//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML Experimental 19960712//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD HTML Experimental 970421//`\"
    - The public identifier starts with: \"`-//W3C//DTD W3 HTML//`\"
    - The public identifier starts with: \"`-//W3O//DTD W3 HTML 3.0//`\"
    - The public identifier starts with:
      \"`-//WebTechs//DTD Mozilla HTML 2.0//`\"
    - The public identifier starts with:
      \"`-//WebTechs//DTD Mozilla HTML//`\"
    - The system identifier is missing and the public identifier starts
      with: \"`-//W3C//DTD HTML 4.01 Frameset//`\"
    - The system identifier is missing and the public identifier starts
      with: \"`-//W3C//DTD HTML 4.01 Transitional//`\"

    Otherwise, if the document is *not* [an `iframe` `srcdoc`
    document](iframe-embed-object.html#an-iframe-srcdoc-document){#the-initial-insertion-mode:an-iframe-srcdoc-document-2},
    and the [parser cannot change the mode
    flag](#parser-cannot-change-the-mode-flag){#the-initial-insertion-mode:parser-cannot-change-the-mode-flag-2}
    is false, and the DOCTYPE token matches one of the conditions in the
    following list, then set the
    [`Document`{#the-initial-insertion-mode:document-6}](dom.html#document)
    to [limited-quirks
    mode](https://dom.spec.whatwg.org/#concept-document-limited-quirks){#the-initial-insertion-mode:limited-quirks-mode
    x-internal="limited-quirks-mode"}:

    - The public identifier starts with:
      \"`-//W3C//DTD XHTML 1.0 Frameset//`\"
    - The public identifier starts with:
      \"`-//W3C//DTD XHTML 1.0 Transitional//`\"
    - The system identifier is not missing and the public identifier
      starts with: \"`-//W3C//DTD HTML 4.01 Frameset//`\"
    - The system identifier is not missing and the public identifier
      starts with: \"`-//W3C//DTD HTML 4.01 Transitional//`\"

    The system identifier and public identifier strings must be compared
    to the values given in the lists above in an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-initial-insertion-mode:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} manner. A system identifier
    whose value is the empty string is not considered missing for the
    purposes of the conditions above.

    Then, switch the [insertion
    mode](#insertion-mode){#the-initial-insertion-mode:insertion-mode-2}
    to \"[before
    html](#the-before-html-insertion-mode){#the-initial-insertion-mode:the-before-html-insertion-mode}\".

Anything else

:   If the document is *not* [an `iframe` `srcdoc`
    document](iframe-embed-object.html#an-iframe-srcdoc-document){#the-initial-insertion-mode:an-iframe-srcdoc-document-3},
    then this is a [parse
    error](#parse-errors){#the-initial-insertion-mode:parse-errors-2};
    if the [parser cannot change the mode
    flag](#parser-cannot-change-the-mode-flag){#the-initial-insertion-mode:parser-cannot-change-the-mode-flag-3}
    is false, set the
    [`Document`{#the-initial-insertion-mode:document-7}](dom.html#document)
    to [quirks
    mode](https://dom.spec.whatwg.org/#concept-document-quirks){#the-initial-insertion-mode:quirks-mode-2
    x-internal="quirks-mode"}.

    In any case, switch the [insertion
    mode](#insertion-mode){#the-initial-insertion-mode:insertion-mode-3}
    to \"[before
    html](#the-before-html-insertion-mode){#the-initial-insertion-mode:the-before-html-insertion-mode-2}\",
    then reprocess the token.

###### [13.2.6.4.2]{.secno} The \"[before html]{.dfn}\" insertion mode[](#the-before-html-insertion-mode){.self-link}

When the user agent is to apply the rules for the \"[before
html](#the-before-html-insertion-mode){#the-before-html-insertion-mode:the-before-html-insertion-mode}\"
[insertion
mode](#insertion-mode){#the-before-html-insertion-mode:insertion-mode},
the user agent must handle the token as follows:

A DOCTYPE token

:   [Parse
    error](#parse-errors){#the-before-html-insertion-mode:parse-errors}.
    Ignore the token.

A comment token

:   [Insert a
    comment](#insert-a-comment){#the-before-html-insertion-mode:insert-a-comment}
    as the last child of the
    [`Document`{#the-before-html-insertion-mode:document}](dom.html#document)
    object.

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   Ignore the token.

A start tag whose tag name is \"html\"

:   [Create an element for the
    token](#create-an-element-for-the-token){#the-before-html-insertion-mode:create-an-element-for-the-token}
    in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-before-html-insertion-mode:html-namespace-2
    x-internal="html-namespace-2"}, with the
    [`Document`{#the-before-html-insertion-mode:document-2}](dom.html#document)
    as the intended parent. Append it to the
    [`Document`{#the-before-html-insertion-mode:document-3}](dom.html#document)
    object. Put this element in the [stack of open
    elements](#stack-of-open-elements){#the-before-html-insertion-mode:stack-of-open-elements}.

    Switch the [insertion
    mode](#insertion-mode){#the-before-html-insertion-mode:insertion-mode-2}
    to \"[before
    head](#the-before-head-insertion-mode){#the-before-html-insertion-mode:the-before-head-insertion-mode}\".

An end tag whose tag name is one of: \"head\", \"body\", \"html\", \"br\"

:   Act as described in the \"anything else\" entry below.

Any other end tag

:   [Parse
    error](#parse-errors){#the-before-html-insertion-mode:parse-errors-2}.
    Ignore the token.

Anything else

:   Create an
    [`html`{#the-before-html-insertion-mode:the-html-element}](semantics.html#the-html-element)
    element whose [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-before-html-insertion-mode:node-document
    x-internal="node-document"} is the
    [`Document`{#the-before-html-insertion-mode:document-4}](dom.html#document)
    object. Append it to the
    [`Document`{#the-before-html-insertion-mode:document-5}](dom.html#document)
    object. Put this element in the [stack of open
    elements](#stack-of-open-elements){#the-before-html-insertion-mode:stack-of-open-elements-2}.

    Switch the [insertion
    mode](#insertion-mode){#the-before-html-insertion-mode:insertion-mode-3}
    to \"[before
    head](#the-before-head-insertion-mode){#the-before-html-insertion-mode:the-before-head-insertion-mode-2}\",
    then reprocess the token.

The [document
element](https://dom.spec.whatwg.org/#document-element){#the-before-html-insertion-mode:document-element
x-internal="document-element"} can end up being removed from the
[`Document`{#the-before-html-insertion-mode:document-6}](dom.html#document)
object, e.g. by scripts; nothing in particular happens in such cases,
content continues being appended to the nodes as described in the next
section.

###### [13.2.6.4.3]{.secno} The \"[before head]{.dfn}\" insertion mode[](#the-before-head-insertion-mode){.self-link}

When the user agent is to apply the rules for the \"[before
head](#the-before-head-insertion-mode){#the-before-head-insertion-mode:the-before-head-insertion-mode}\"
[insertion
mode](#insertion-mode){#the-before-head-insertion-mode:insertion-mode},
the user agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   Ignore the token.

A comment token

:   [Insert a
    comment](#insert-a-comment){#the-before-head-insertion-mode:insert-a-comment}.

A DOCTYPE token

:   [Parse
    error](#parse-errors){#the-before-head-insertion-mode:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#the-before-head-insertion-mode:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#the-before-head-insertion-mode:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#the-before-head-insertion-mode:insertion-mode-2}.

A start tag whose tag name is \"head\"

:   [Insert an HTML
    element](#insert-an-html-element){#the-before-head-insertion-mode:insert-an-html-element}
    for the token.

    Set the [`head` element
    pointer](#head-element-pointer){#the-before-head-insertion-mode:head-element-pointer}
    to the newly created
    [`head`{#the-before-head-insertion-mode:the-head-element}](semantics.html#the-head-element)
    element.

    Switch the [insertion
    mode](#insertion-mode){#the-before-head-insertion-mode:insertion-mode-3}
    to \"[in
    head](#parsing-main-inhead){#the-before-head-insertion-mode:parsing-main-inhead}\".

An end tag whose tag name is one of: \"head\", \"body\", \"html\", \"br\"

:   Act as described in the \"anything else\" entry below.

Any other end tag

:   [Parse
    error](#parse-errors){#the-before-head-insertion-mode:parse-errors-2}.
    Ignore the token.

Anything else

:   [Insert an HTML
    element](#insert-an-html-element){#the-before-head-insertion-mode:insert-an-html-element-2}
    for a \"head\" start tag token with no attributes.

    Set the [`head` element
    pointer](#head-element-pointer){#the-before-head-insertion-mode:head-element-pointer-2}
    to the newly created
    [`head`{#the-before-head-insertion-mode:the-head-element-2}](semantics.html#the-head-element)
    element.

    Switch the [insertion
    mode](#insertion-mode){#the-before-head-insertion-mode:insertion-mode-4}
    to \"[in
    head](#parsing-main-inhead){#the-before-head-insertion-mode:parsing-main-inhead-2}\".

    Reprocess the current token.

###### [13.2.6.4.4]{.secno} The \"[in head]{.dfn}\" insertion mode[](#parsing-main-inhead){.self-link} {#parsing-main-inhead}

When the user agent is to apply the rules for the \"[in
head](#parsing-main-inhead){#parsing-main-inhead:parsing-main-inhead}\"
[insertion mode](#insertion-mode){#parsing-main-inhead:insertion-mode},
the user agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   [Insert the
    character](#insert-a-character){#parsing-main-inhead:insert-a-character}.

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-inhead:insert-a-comment}.

A DOCTYPE token

:   [Parse error](#parse-errors){#parsing-main-inhead:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inhead:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-inhead:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inhead:insertion-mode-2}.

A start tag whose tag name is one of: \"base\", \"basefont\", \"bgsound\", \"link\"

:   [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inhead:insert-an-html-element}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inhead:current-node} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inhead:acknowledge-self-closing-flag},
    if it is set.

A start tag whose tag name is \"meta\"

:   [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inhead:insert-an-html-element-2}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inhead:current-node-2} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements-2}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inhead:acknowledge-self-closing-flag-2},
    if it is set.

    If the [active speculative HTML
    parser](#active-speculative-html-parser){#parsing-main-inhead:active-speculative-html-parser}
    is null, then:

    1.  If the element has a
        [`charset`{#parsing-main-inhead:attr-meta-charset}](semantics.html#attr-meta-charset)
        attribute, and [getting an
        encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#parsing-main-inhead:getting-an-encoding
        x-internal="getting-an-encoding"} from its value results in an
        [encoding](https://encoding.spec.whatwg.org/#encoding){#parsing-main-inhead:encoding
        x-internal="encoding"}, and the
        [confidence](#concept-encoding-confidence){#parsing-main-inhead:concept-encoding-confidence}
        is currently *tentative*, then [change the
        encoding](#change-the-encoding){#parsing-main-inhead:change-the-encoding}
        to the resulting encoding.

    2.  Otherwise, if the element has an
        [`http-equiv`{#parsing-main-inhead:attr-meta-http-equiv}](semantics.html#attr-meta-http-equiv)
        attribute whose value is an [ASCII
        case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#parsing-main-inhead:ascii-case-insensitive
        x-internal="ascii-case-insensitive"} match for the string
        \"`Content-Type`\", and the element has a
        [`content`{#parsing-main-inhead:attr-meta-content}](semantics.html#attr-meta-content)
        attribute, and applying the [algorithm for extracting a
        character encoding from a `meta`
        element](urls-and-fetching.html#algorithm-for-extracting-a-character-encoding-from-a-meta-element){#parsing-main-inhead:algorithm-for-extracting-a-character-encoding-from-a-meta-element}
        to that attribute\'s value returns an
        [encoding](https://encoding.spec.whatwg.org/#encoding){#parsing-main-inhead:encoding-2
        x-internal="encoding"}, and the
        [confidence](#concept-encoding-confidence){#parsing-main-inhead:concept-encoding-confidence-2}
        is currently *tentative*, then [change the
        encoding](#change-the-encoding){#parsing-main-inhead:change-the-encoding-2}
        to the extracted encoding.

    The [speculative HTML
    parser](#speculative-html-parser){#parsing-main-inhead:speculative-html-parser}
    doesn\'t speculatively apply character encoding declarations in
    order to reduce implementation complexity.

A start tag whose tag name is \"title\"

:   Follow the [generic RCDATA element parsing
    algorithm](#generic-rcdata-element-parsing-algorithm){#parsing-main-inhead:generic-rcdata-element-parsing-algorithm}.

A start tag whose tag name is \"noscript\", if the [scripting flag](#scripting-flag){#parsing-main-inhead:scripting-flag} is enabled\
A start tag whose tag name is one of: \"noframes\", \"style\"

:   Follow the [generic raw text element parsing
    algorithm](#generic-raw-text-element-parsing-algorithm){#parsing-main-inhead:generic-raw-text-element-parsing-algorithm}.

A start tag whose tag name is \"noscript\", if the [scripting flag](#scripting-flag){#parsing-main-inhead:scripting-flag-2} is disabled

:   [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inhead:insert-an-html-element-3}
    for the token.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inhead:insertion-mode-3} to
    \"[in head
    noscript](#parsing-main-inheadnoscript){#parsing-main-inhead:parsing-main-inheadnoscript}\".

A start tag whose tag name is \"script\"

:   Run these steps:

    1.  Let the `adjusted insertion location`{.variable} be the
        [appropriate place for inserting a
        node](#appropriate-place-for-inserting-a-node){#parsing-main-inhead:appropriate-place-for-inserting-a-node}.

    2.  [Create an element for the
        token](#create-an-element-for-the-token){#parsing-main-inhead:create-an-element-for-the-token}
        in the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing-main-inhead:html-namespace-2
        x-internal="html-namespace-2"}, with the intended parent being
        the element in which the
        `adjusted insertion location`{.variable} finds itself.

    3.  Set the element\'s [parser
        document](scripting.html#parser-document){#parsing-main-inhead:parser-document}
        to the
        [`Document`{#parsing-main-inhead:document}](dom.html#document),
        and set the element\'s [force
        async](scripting.html#script-force-async){#parsing-main-inhead:script-force-async}
        to false.

        This ensures that, if the script is external, any
        [`document.write()`{#parsing-main-inhead:dom-document-write}](dynamic-markup-insertion.html#dom-document-write)
        calls in the script will execute in-line, instead of blowing the
        document away, as would happen in most other cases. It also
        prevents the script from executing until the end tag is seen.

    4.  If the parser was created as part of the [HTML fragment parsing
        algorithm](#html-fragment-parsing-algorithm){#parsing-main-inhead:html-fragment-parsing-algorithm},
        then set the
        [`script`{#parsing-main-inhead:the-script-element}](scripting.html#the-script-element)
        element\'s [already
        started](scripting.html#already-started){#parsing-main-inhead:already-started}
        to true. ([fragment
        case](#fragment-case){#parsing-main-inhead:fragment-case})

    5.  [If the parser was invoked via the
        [`document.write()`{#parsing-main-inhead:dom-document-write-2}](dynamic-markup-insertion.html#dom-document-write)
        or
        [`document.writeln()`{#parsing-main-inhead:dom-document-writeln}](dynamic-markup-insertion.html#dom-document-writeln)
        methods, then optionally set the
        [`script`{#parsing-main-inhead:the-script-element-2}](scripting.html#the-script-element)
        element\'s [already
        started](scripting.html#already-started){#parsing-main-inhead:already-started-2}
        to true. (For example, the user agent might use this clause to
        prevent execution of
        [cross-origin](browsers.html#concept-origin){#parsing-main-inhead:concept-origin}
        scripts inserted via
        [`document.write()`{#parsing-main-inhead:dom-document-write-3}](dynamic-markup-insertion.html#dom-document-write)
        under slow network conditions, or when the page has already
        taken a long time to
        load.)]{#document-written-scripts-intervention}

    6.  Insert the newly created element at the
        `adjusted insertion location`{.variable}.

    7.  Push the element onto the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements-3}
        so that it is the new [current
        node](#current-node){#parsing-main-inhead:current-node-3}.

    8.  Switch the tokenizer to the [script data
        state](#script-data-state){#parsing-main-inhead:script-data-state}.

    9.  Set the [original insertion
        mode](#original-insertion-mode){#parsing-main-inhead:original-insertion-mode}
        to the current [insertion
        mode](#insertion-mode){#parsing-main-inhead:insertion-mode-4}.

    10. Switch the [insertion
        mode](#insertion-mode){#parsing-main-inhead:insertion-mode-5} to
        \"[text](#parsing-main-incdata){#parsing-main-inhead:parsing-main-incdata}\".

An end tag whose tag name is \"head\"

:   Pop the [current
    node](#current-node){#parsing-main-inhead:current-node-4} (which
    will be the
    [`head`{#parsing-main-inhead:the-head-element}](semantics.html#the-head-element)
    element) off the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements-4}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inhead:insertion-mode-6} to
    \"[after
    head](#the-after-head-insertion-mode){#parsing-main-inhead:the-after-head-insertion-mode}\".

An end tag whose tag name is one of: \"body\", \"html\", \"br\"

:   Act as described in the \"anything else\" entry below.

A start tag whose tag name is \"template\"

:   Run these steps:

    1.  Let `templateStartTag`{.variable} be the start tag.

    2.  Insert a
        [marker](#concept-parser-marker){#parsing-main-inhead:concept-parser-marker}
        at the end of the [list of active formatting
        elements](#list-of-active-formatting-elements){#parsing-main-inhead:list-of-active-formatting-elements}.

    3.  Set the [frameset-ok
        flag](#frameset-ok-flag){#parsing-main-inhead:frameset-ok-flag}
        to \"not ok\".

    4.  Switch the [insertion
        mode](#insertion-mode){#parsing-main-inhead:insertion-mode-7} to
        \"[in
        template](#parsing-main-intemplate){#parsing-main-inhead:parsing-main-intemplate}\".

    5.  Push \"[in
        template](#parsing-main-intemplate){#parsing-main-inhead:parsing-main-intemplate-2}\"
        onto the [stack of template insertion
        modes](#stack-of-template-insertion-modes){#parsing-main-inhead:stack-of-template-insertion-modes}
        so that it is the new [current template insertion
        mode](#current-template-insertion-mode){#parsing-main-inhead:current-template-insertion-mode}.

    6.  Let the `adjustedInsertionLocation`{.variable} be the
        [appropriate place for inserting a
        node](#appropriate-place-for-inserting-a-node){#parsing-main-inhead:appropriate-place-for-inserting-a-node-2}.

    7.  Let `intendedParent`{.variable} be the element in which the
        `adjustedInsertionLocation`{.variable} finds itself.

    8.  Let `document`{.variable} be `intendedParent`{.variable}\'s
        [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#parsing-main-inhead:node-document
        x-internal="node-document"}.

    9.  If any of the following are false:

        - `templateStartTag`{.variable}\'s
          [`shadowrootmode`{#parsing-main-inhead:attr-template-shadowrootmode}](scripting.html#attr-template-shadowrootmode)
          is not in the
          [None](scripting.html#attr-shadowrootmode-none-state){#parsing-main-inhead:attr-shadowrootmode-none-state}
          state;
        - `document`{.variable}\'s [allow declarative shadow
          roots](https://dom.spec.whatwg.org/#concept-document-allow-declarative-shadow-roots){#parsing-main-inhead:concept-document-allow-declarative-shadow-roots
          x-internal="concept-document-allow-declarative-shadow-roots"}
          is true; or
        - the [adjusted current
          node](#adjusted-current-node){#parsing-main-inhead:adjusted-current-node}
          is not the topmost element in the [stack of open
          elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements-5},

        then [insert an HTML
        element](#insert-an-html-element){#parsing-main-inhead:insert-an-html-element-4}
        for the token.

    10. Otherwise:

        1.  Let `declarativeShadowHostElement`{.variable} be [adjusted
            current
            node](#adjusted-current-node){#parsing-main-inhead:adjusted-current-node-2}.

        2.  Let `template`{.variable} be the result of [insert a foreign
            element](#insert-a-foreign-element){#parsing-main-inhead:insert-a-foreign-element}
            for `templateStartTag`{.variable}, with [HTML
            namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing-main-inhead:html-namespace-2-2
            x-internal="html-namespace-2"} and true.

        3.  Let `mode`{.variable} be `templateStartTag`{.variable}\'s
            [`shadowrootmode`{#parsing-main-inhead:attr-template-shadowrootmode-2}](scripting.html#attr-template-shadowrootmode)
            attribute\'s value.

        4.  Let `clonable`{.variable} be true if
            `templateStartTag`{.variable} has a
            [`shadowrootclonable`{#parsing-main-inhead:attr-template-shadowrootclonable}](scripting.html#attr-template-shadowrootclonable)
            attribute; otherwise false.

        5.  Let `serializable`{.variable} be true if
            `templateStartTag`{.variable} has a
            [`shadowrootserializable`{#parsing-main-inhead:attr-template-shadowrootserializable}](scripting.html#attr-template-shadowrootserializable)
            attribute; otherwise false.

        6.  Let `delegatesFocus`{.variable} be true if
            `templateStartTag`{.variable} has a
            [`shadowrootdelegatesfocus`{#parsing-main-inhead:attr-template-shadowrootdelegatesfocus}](scripting.html#attr-template-shadowrootdelegatesfocus)
            attribute; otherwise false.

        7.  If `declarativeShadowHostElement`{.variable} is a [shadow
            host](https://dom.spec.whatwg.org/#element-shadow-host){#parsing-main-inhead:shadow-host
            x-internal="shadow-host"}, then [insert an element at the
            adjusted insertion
            location](#insert-an-element-at-the-adjusted-insertion-location){#parsing-main-inhead:insert-an-element-at-the-adjusted-insertion-location}
            with `template`{.variable}.

        8.  Otherwise:

            1.  Let `registry`{.variable} be
                `declarativeShadowHostElement`{.variable}\'s [custom
                element
                registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#parsing-main-inhead:element-custom-element-registry
                x-internal="element-custom-element-registry"}.

            2.  If `templateStartTag`{.variable} has a
                [`shadowrootcustomelementregistry`{#parsing-main-inhead:attr-template-shadowrootcustomelementregistry}](scripting.html#attr-template-shadowrootcustomelementregistry)
                attribute, then set `registry`{.variable} to null.

            3.  [Attach a shadow
                root](https://dom.spec.whatwg.org/#concept-attach-a-shadow-root){#parsing-main-inhead:concept-attach-a-shadow-root
                x-internal="concept-attach-a-shadow-root"} with
                `declarativeShadowHostElement`{.variable},
                `mode`{.variable}, `clonable`{.variable},
                `serializable`{.variable}, `delegatesFocus`{.variable},
                \"`named`\", and `registry`{.variable}.

                If an exception is thrown, then catch it and:

                1.  [Insert an element at the adjusted insertion
                    location](#insert-an-element-at-the-adjusted-insertion-location){#parsing-main-inhead:insert-an-element-at-the-adjusted-insertion-location-2}
                    with `template`{.variable}.

                2.  The user agent may report an error to the developer
                    console.

                3.  Return.

            4.  Let `shadow`{.variable} be
                `declarativeShadowHostElement`{.variable}\'s [shadow
                root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#parsing-main-inhead:concept-element-shadow-root
                x-internal="concept-element-shadow-root"}.

            5.  Set `shadow`{.variable}\'s
                [declarative](https://dom.spec.whatwg.org/#shadowroot-declarative){#parsing-main-inhead:concept-shadow-root-declarative
                x-internal="concept-shadow-root-declarative"} to true.

            6.  Set
                [template](scripting.html#the-template-element){#parsing-main-inhead:the-template-element}\'s
                [template
                contents](scripting.html#template-contents){#parsing-main-inhead:template-contents}
                property to `shadow`{.variable}.

            7.  Set `shadow`{.variable}\'s [available to element
                internals](https://dom.spec.whatwg.org/#shadowroot-available-to-element-internals){#parsing-main-inhead:available-to-element-internals
                x-internal="available-to-element-internals"} to true.

            8.  If `templateStartTag`{.variable} has a
                [`shadowrootcustomelementregistry`{#parsing-main-inhead:attr-template-shadowrootcustomelementregistry-2}](scripting.html#attr-template-shadowrootcustomelementregistry)
                attribute, then set `shadow`{.variable}\'s [keep custom
                element registry
                null](https://dom.spec.whatwg.org/#shadowroot-keep-custom-element-registry-null){#parsing-main-inhead:keep-custom-element-registry-null
                x-internal="keep-custom-element-registry-null"} to true.

An end tag whose tag name is \"template\"

:   If there is no
    [`template`{#parsing-main-inhead:the-template-element-2}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements-6},
    then this is a [parse
    error](#parse-errors){#parsing-main-inhead:parse-errors-2}; ignore
    the token.

    Otherwise, run these steps:

    1.  [Generate all implied end tags
        thoroughly](#generate-all-implied-end-tags-thoroughly){#parsing-main-inhead:generate-all-implied-end-tags-thoroughly}.

    2.  If the [current
        node](#current-node){#parsing-main-inhead:current-node-5} is not
        a
        [`template`{#parsing-main-inhead:the-template-element-3}](scripting.html#the-template-element)
        element, then this is a [parse
        error](#parse-errors){#parsing-main-inhead:parse-errors-3}.

    3.  Pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements-7}
        until a
        [`template`{#parsing-main-inhead:the-template-element-4}](scripting.html#the-template-element)
        element has been popped from the stack.

    4.  [Clear the list of active formatting elements up to the last
        marker](#clear-the-list-of-active-formatting-elements-up-to-the-last-marker){#parsing-main-inhead:clear-the-list-of-active-formatting-elements-up-to-the-last-marker}.

    5.  Pop the [current template insertion
        mode](#current-template-insertion-mode){#parsing-main-inhead:current-template-insertion-mode-2}
        off the [stack of template insertion
        modes](#stack-of-template-insertion-modes){#parsing-main-inhead:stack-of-template-insertion-modes-2}.

    6.  [Reset the insertion mode
        appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-inhead:reset-the-insertion-mode-appropriately}.

A start tag whose tag name is \"head\"\
Any other end tag

:   [Parse error](#parse-errors){#parsing-main-inhead:parse-errors-4}.
    Ignore the token.

Anything else

:   Pop the [current
    node](#current-node){#parsing-main-inhead:current-node-6} (which
    will be the
    [`head`{#parsing-main-inhead:the-head-element-2}](semantics.html#the-head-element)
    element) off the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inhead:stack-of-open-elements-8}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inhead:insertion-mode-8} to
    \"[after
    head](#the-after-head-insertion-mode){#parsing-main-inhead:the-after-head-insertion-mode-2}\".

    Reprocess the token.

###### [13.2.6.4.5]{.secno} The \"[in head noscript]{.dfn}\" insertion mode[](#parsing-main-inheadnoscript){.self-link} {#parsing-main-inheadnoscript}

When the user agent is to apply the rules for the \"[in head
noscript](#parsing-main-inheadnoscript){#parsing-main-inheadnoscript:parsing-main-inheadnoscript}\"
[insertion
mode](#insertion-mode){#parsing-main-inheadnoscript:insertion-mode}, the
user agent must handle the token as follows:

A DOCTYPE token

:   [Parse
    error](#parse-errors){#parsing-main-inheadnoscript:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inheadnoscript:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-inheadnoscript:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inheadnoscript:insertion-mode-2}.

An end tag whose tag name is \"noscript\"

:   Pop the [current
    node](#current-node){#parsing-main-inheadnoscript:current-node}
    (which will be a
    [`noscript`{#parsing-main-inheadnoscript:the-noscript-element}](scripting.html#the-noscript-element)
    element) from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inheadnoscript:stack-of-open-elements};
    the new [current
    node](#current-node){#parsing-main-inheadnoscript:current-node-2}
    will be a
    [`head`{#parsing-main-inheadnoscript:the-head-element}](semantics.html#the-head-element)
    element.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inheadnoscript:insertion-mode-3}
    to \"[in
    head](#parsing-main-inhead){#parsing-main-inheadnoscript:parsing-main-inhead}\".

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE\
A comment token\
A start tag whose tag name is one of: \"basefont\", \"bgsound\", \"link\", \"meta\", \"noframes\", \"style\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inheadnoscript:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-inheadnoscript:parsing-main-inhead-2}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inheadnoscript:insertion-mode-4}.

An end tag whose tag name is \"br\"

:   Act as described in the \"anything else\" entry below.

A start tag whose tag name is one of: \"head\", \"noscript\"\
Any other end tag

:   [Parse
    error](#parse-errors){#parsing-main-inheadnoscript:parse-errors-2}.
    Ignore the token.

Anything else

:   [Parse
    error](#parse-errors){#parsing-main-inheadnoscript:parse-errors-3}.

    Pop the [current
    node](#current-node){#parsing-main-inheadnoscript:current-node-3}
    (which will be a
    [`noscript`{#parsing-main-inheadnoscript:the-noscript-element-2}](scripting.html#the-noscript-element)
    element) from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inheadnoscript:stack-of-open-elements-2};
    the new [current
    node](#current-node){#parsing-main-inheadnoscript:current-node-4}
    will be a
    [`head`{#parsing-main-inheadnoscript:the-head-element-2}](semantics.html#the-head-element)
    element.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inheadnoscript:insertion-mode-5}
    to \"[in
    head](#parsing-main-inhead){#parsing-main-inheadnoscript:parsing-main-inhead-3}\".

    Reprocess the token.

###### [13.2.6.4.6]{.secno} The \"[after head]{.dfn}\" insertion mode[](#the-after-head-insertion-mode){.self-link}

When the user agent is to apply the rules for the \"[after
head](#the-after-head-insertion-mode){#the-after-head-insertion-mode:the-after-head-insertion-mode}\"
[insertion
mode](#insertion-mode){#the-after-head-insertion-mode:insertion-mode},
the user agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   [Insert the
    character](#insert-a-character){#the-after-head-insertion-mode:insert-a-character}.

A comment token

:   [Insert a
    comment](#insert-a-comment){#the-after-head-insertion-mode:insert-a-comment}.

A DOCTYPE token

:   [Parse
    error](#parse-errors){#the-after-head-insertion-mode:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#the-after-head-insertion-mode:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#the-after-head-insertion-mode:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#the-after-head-insertion-mode:insertion-mode-2}.

A start tag whose tag name is \"body\"

:   [Insert an HTML
    element](#insert-an-html-element){#the-after-head-insertion-mode:insert-an-html-element}
    for the token.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#the-after-head-insertion-mode:frameset-ok-flag}
    to \"not ok\".

    Switch the [insertion
    mode](#insertion-mode){#the-after-head-insertion-mode:insertion-mode-3}
    to \"[in
    body](#parsing-main-inbody){#the-after-head-insertion-mode:parsing-main-inbody-2}\".

A start tag whose tag name is \"frameset\"

:   [Insert an HTML
    element](#insert-an-html-element){#the-after-head-insertion-mode:insert-an-html-element-2}
    for the token.

    Switch the [insertion
    mode](#insertion-mode){#the-after-head-insertion-mode:insertion-mode-4}
    to \"[in
    frameset](#parsing-main-inframeset){#the-after-head-insertion-mode:parsing-main-inframeset}\".

A start tag whose tag name is one of: \"base\", \"basefont\", \"bgsound\", \"link\", \"meta\", \"noframes\", \"script\", \"style\", \"template\", \"title\"

:   [Parse
    error](#parse-errors){#the-after-head-insertion-mode:parse-errors-2}.

    Push the node pointed to by the [`head` element
    pointer](#head-element-pointer){#the-after-head-insertion-mode:head-element-pointer}
    onto the [stack of open
    elements](#stack-of-open-elements){#the-after-head-insertion-mode:stack-of-open-elements}.

    Process the token [using the rules
    for](#using-the-rules-for){#the-after-head-insertion-mode:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#the-after-head-insertion-mode:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#the-after-head-insertion-mode:insertion-mode-5}.

    Remove the node pointed to by the [`head` element
    pointer](#head-element-pointer){#the-after-head-insertion-mode:head-element-pointer-2}
    from the [stack of open
    elements](#stack-of-open-elements){#the-after-head-insertion-mode:stack-of-open-elements-2}.
    (It might not be the [current
    node](#current-node){#the-after-head-insertion-mode:current-node} at
    this point.)

    The [`head` element
    pointer](#head-element-pointer){#the-after-head-insertion-mode:head-element-pointer-3}
    cannot be null at this point.

An end tag whose tag name is \"template\"

:   Process the token [using the rules
    for](#using-the-rules-for){#the-after-head-insertion-mode:using-the-rules-for-3}
    the \"[in
    head](#parsing-main-inhead){#the-after-head-insertion-mode:parsing-main-inhead-2}\"
    [insertion
    mode](#insertion-mode){#the-after-head-insertion-mode:insertion-mode-6}.

An end tag whose tag name is one of: \"body\", \"html\", \"br\"

:   Act as described in the \"anything else\" entry below.

A start tag whose tag name is \"head\"\
Any other end tag

:   [Parse
    error](#parse-errors){#the-after-head-insertion-mode:parse-errors-3}.
    Ignore the token.

Anything else

:   [Insert an HTML
    element](#insert-an-html-element){#the-after-head-insertion-mode:insert-an-html-element-3}
    for a \"body\" start tag token with no attributes.

    Switch the [insertion
    mode](#insertion-mode){#the-after-head-insertion-mode:insertion-mode-7}
    to \"[in
    body](#parsing-main-inbody){#the-after-head-insertion-mode:parsing-main-inbody-3}\".

    Reprocess the current token.

###### [13.2.6.4.7]{.secno} The \"[in body]{.dfn}\" insertion mode[](#parsing-main-inbody){.self-link} {#parsing-main-inbody}

When the user agent is to apply the rules for the \"[in
body](#parsing-main-inbody){#parsing-main-inbody:parsing-main-inbody}\"
[insertion mode](#insertion-mode){#parsing-main-inbody:insertion-mode},
the user agent must handle the token as follows:

A character token that is U+0000 NULL

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors}.
    Ignore the token.

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements},
    if any.

    [Insert the token\'s
    character](#insert-a-character){#parsing-main-inbody:insert-a-character}.

Any other character token

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-2},
    if any.

    [Insert the token\'s
    character](#insert-a-character){#parsing-main-inbody:insert-a-character-2}.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag} to
    \"not ok\".

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-inbody:insert-a-comment}.

A DOCTYPE token

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors-2}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors-3}.

    If there is a
    [`template`{#parsing-main-inbody:the-template-element}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements},
    then ignore the token.

    Otherwise, for each attribute on the token, check to see if the
    attribute is already present on the top element of the [stack of
    open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-2}.
    If it is not, add the attribute and its corresponding value to that
    element.

A start tag whose tag name is one of: \"base\", \"basefont\", \"bgsound\", \"link\", \"meta\", \"noframes\", \"script\", \"style\", \"template\", \"title\"\
An end tag whose tag name is \"template\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inbody:using-the-rules-for}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-inbody:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-2}.

A start tag whose tag name is \"body\"

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors-4}.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-3}
    has only one node on it, if the second element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-4}
    is not a
    [`body`{#parsing-main-inbody:the-body-element}](sections.html#the-body-element)
    element, or if there is a
    [`template`{#parsing-main-inbody:the-template-element-2}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-5},
    then ignore the token. ([fragment
    case](#fragment-case){#parsing-main-inbody:fragment-case} or there
    is a
    [`template`{#parsing-main-inbody:the-template-element-3}](scripting.html#the-template-element)
    element on the stack)

    Otherwise, set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-2} to
    \"not ok\"; then, for each attribute on the token, check to see if
    the attribute is already present on the
    [`body`{#parsing-main-inbody:the-body-element-2}](sections.html#the-body-element)
    element (the second element) on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-6},
    and if it is not, add the attribute and its corresponding value to
    that element.

A start tag whose tag name is \"frameset\"

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors-5}.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-7}
    has only one node on it, or if the second element on the [stack of
    open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-8}
    is not a
    [`body`{#parsing-main-inbody:the-body-element-3}](sections.html#the-body-element)
    element, then ignore the token. ([fragment
    case](#fragment-case){#parsing-main-inbody:fragment-case-2} or there
    is a
    [`template`{#parsing-main-inbody:the-template-element-4}](scripting.html#the-template-element)
    element on the stack)

    If the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-3} is
    set to \"not ok\", ignore the token.

    Otherwise, run the following steps:

    1.  Remove the second element on the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-9}
        from its parent node, if it has one.

    2.  Pop all the nodes from the bottom of the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-10},
        from the [current
        node](#current-node){#parsing-main-inbody:current-node} up to,
        but not including, the root
        [`html`{#parsing-main-inbody:the-html-element}](semantics.html#the-html-element)
        element.

    3.  [Insert an HTML
        element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element}
        for the token.

    4.  Switch the [insertion
        mode](#insertion-mode){#parsing-main-inbody:insertion-mode-3} to
        \"[in
        frameset](#parsing-main-inframeset){#parsing-main-inbody:parsing-main-inframeset}\".

An end-of-file token

:   If the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-inbody:stack-of-template-insertion-modes}
    is not empty, then process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inbody:using-the-rules-for-2}
    the \"[in
    template](#parsing-main-intemplate){#parsing-main-inbody:parsing-main-intemplate}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-4}.

    Otherwise, follow these steps:

    1.  If there is a node in the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-11}
        that is not either a
        [`dd`{#parsing-main-inbody:the-dd-element}](grouping-content.html#the-dd-element)
        element, a
        [`dt`{#parsing-main-inbody:the-dt-element}](grouping-content.html#the-dt-element)
        element, an
        [`li`{#parsing-main-inbody:the-li-element}](grouping-content.html#the-li-element)
        element, an
        [`optgroup`{#parsing-main-inbody:the-optgroup-element}](form-elements.html#the-optgroup-element)
        element, an
        [`option`{#parsing-main-inbody:the-option-element}](form-elements.html#the-option-element)
        element, a
        [`p`{#parsing-main-inbody:the-p-element}](grouping-content.html#the-p-element)
        element, an [`rb`{#parsing-main-inbody:rb}](obsolete.html#rb)
        element, an
        [`rp`{#parsing-main-inbody:the-rp-element}](text-level-semantics.html#the-rp-element)
        element, an
        [`rt`{#parsing-main-inbody:the-rt-element}](text-level-semantics.html#the-rt-element)
        element, an [`rtc`{#parsing-main-inbody:rtc}](obsolete.html#rtc)
        element, a
        [`tbody`{#parsing-main-inbody:the-tbody-element}](tables.html#the-tbody-element)
        element, a
        [`td`{#parsing-main-inbody:the-td-element}](tables.html#the-td-element)
        element, a
        [`tfoot`{#parsing-main-inbody:the-tfoot-element}](tables.html#the-tfoot-element)
        element, a
        [`th`{#parsing-main-inbody:the-th-element}](tables.html#the-th-element)
        element, a
        [`thead`{#parsing-main-inbody:the-thead-element}](tables.html#the-thead-element)
        element, a
        [`tr`{#parsing-main-inbody:the-tr-element}](tables.html#the-tr-element)
        element, the
        [`body`{#parsing-main-inbody:the-body-element-4}](sections.html#the-body-element)
        element, or the
        [`html`{#parsing-main-inbody:the-html-element-2}](semantics.html#the-html-element)
        element, then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-6}.

    2.  [Stop
        parsing](#stop-parsing){#parsing-main-inbody:stop-parsing}.

An end tag whose tag name is \"body\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-12}
    does not [have a `body` element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope},
    this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-7}; ignore
    the token.

    Otherwise, if there is a node in the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-13}
    that is not either a
    [`dd`{#parsing-main-inbody:the-dd-element-2}](grouping-content.html#the-dd-element)
    element, a
    [`dt`{#parsing-main-inbody:the-dt-element-2}](grouping-content.html#the-dt-element)
    element, an
    [`li`{#parsing-main-inbody:the-li-element-2}](grouping-content.html#the-li-element)
    element, an
    [`optgroup`{#parsing-main-inbody:the-optgroup-element-2}](form-elements.html#the-optgroup-element)
    element, an
    [`option`{#parsing-main-inbody:the-option-element-2}](form-elements.html#the-option-element)
    element, a
    [`p`{#parsing-main-inbody:the-p-element-2}](grouping-content.html#the-p-element)
    element, an [`rb`{#parsing-main-inbody:rb-2}](obsolete.html#rb)
    element, an
    [`rp`{#parsing-main-inbody:the-rp-element-2}](text-level-semantics.html#the-rp-element)
    element, an
    [`rt`{#parsing-main-inbody:the-rt-element-2}](text-level-semantics.html#the-rt-element)
    element, an [`rtc`{#parsing-main-inbody:rtc-2}](obsolete.html#rtc)
    element, a
    [`tbody`{#parsing-main-inbody:the-tbody-element-2}](tables.html#the-tbody-element)
    element, a
    [`td`{#parsing-main-inbody:the-td-element-2}](tables.html#the-td-element)
    element, a
    [`tfoot`{#parsing-main-inbody:the-tfoot-element-2}](tables.html#the-tfoot-element)
    element, a
    [`th`{#parsing-main-inbody:the-th-element-2}](tables.html#the-th-element)
    element, a
    [`thead`{#parsing-main-inbody:the-thead-element-2}](tables.html#the-thead-element)
    element, a
    [`tr`{#parsing-main-inbody:the-tr-element-2}](tables.html#the-tr-element)
    element, the
    [`body`{#parsing-main-inbody:the-body-element-5}](sections.html#the-body-element)
    element, or the
    [`html`{#parsing-main-inbody:the-html-element-3}](semantics.html#the-html-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-8}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-5} to
    \"[after
    body](#parsing-main-afterbody){#parsing-main-inbody:parsing-main-afterbody}\".

An end tag whose tag name is \"html\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-14}
    does not [have a `body` element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-2},
    this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-9}; ignore
    the token.

    Otherwise, if there is a node in the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-15}
    that is not either a
    [`dd`{#parsing-main-inbody:the-dd-element-3}](grouping-content.html#the-dd-element)
    element, a
    [`dt`{#parsing-main-inbody:the-dt-element-3}](grouping-content.html#the-dt-element)
    element, an
    [`li`{#parsing-main-inbody:the-li-element-3}](grouping-content.html#the-li-element)
    element, an
    [`optgroup`{#parsing-main-inbody:the-optgroup-element-3}](form-elements.html#the-optgroup-element)
    element, an
    [`option`{#parsing-main-inbody:the-option-element-3}](form-elements.html#the-option-element)
    element, a
    [`p`{#parsing-main-inbody:the-p-element-3}](grouping-content.html#the-p-element)
    element, an [`rb`{#parsing-main-inbody:rb-3}](obsolete.html#rb)
    element, an
    [`rp`{#parsing-main-inbody:the-rp-element-3}](text-level-semantics.html#the-rp-element)
    element, an
    [`rt`{#parsing-main-inbody:the-rt-element-3}](text-level-semantics.html#the-rt-element)
    element, an [`rtc`{#parsing-main-inbody:rtc-3}](obsolete.html#rtc)
    element, a
    [`tbody`{#parsing-main-inbody:the-tbody-element-3}](tables.html#the-tbody-element)
    element, a
    [`td`{#parsing-main-inbody:the-td-element-3}](tables.html#the-td-element)
    element, a
    [`tfoot`{#parsing-main-inbody:the-tfoot-element-3}](tables.html#the-tfoot-element)
    element, a
    [`th`{#parsing-main-inbody:the-th-element-3}](tables.html#the-th-element)
    element, a
    [`thead`{#parsing-main-inbody:the-thead-element-3}](tables.html#the-thead-element)
    element, a
    [`tr`{#parsing-main-inbody:the-tr-element-3}](tables.html#the-tr-element)
    element, the
    [`body`{#parsing-main-inbody:the-body-element-6}](sections.html#the-body-element)
    element, or the
    [`html`{#parsing-main-inbody:the-html-element-4}](semantics.html#the-html-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-10}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-6} to
    \"[after
    body](#parsing-main-afterbody){#parsing-main-inbody:parsing-main-afterbody-2}\".

    Reprocess the token.

A start tag whose tag name is one of: \"address\", \"article\", \"aside\", \"blockquote\", \"center\", \"details\", \"dialog\", \"dir\", \"div\", \"dl\", \"fieldset\", \"figcaption\", \"figure\", \"footer\", \"header\", \"hgroup\", \"main\", \"menu\", \"nav\", \"ol\", \"p\", \"search\", \"section\", \"summary\", \"ul\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-16}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-2}
    for the token.

A start tag whose tag name is one of: \"h1\", \"h2\", \"h3\", \"h4\", \"h5\", \"h6\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-17}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-2},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-2}.

    If the [current
    node](#current-node){#parsing-main-inbody:current-node-2} is an
    [HTML
    element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements}
    whose tag name is one of \"h1\", \"h2\", \"h3\", \"h4\", \"h5\", or
    \"h6\", then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-11}; pop the
    [current node](#current-node){#parsing-main-inbody:current-node-3}
    off the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-18}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-3}
    for the token.

A start tag whose tag name is one of: \"pre\", \"listing\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-19}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-3},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-3}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-4}
    for the token.

    If the [next token](#next-token){#parsing-main-inbody:next-token} is
    a U+000A LINE FEED (LF) character token, then ignore that token and
    move on to the next one. (Newlines at the start of
    [`pre`{#parsing-main-inbody:the-pre-element}](grouping-content.html#the-pre-element)
    blocks are ignored as an authoring convenience.)

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-4} to
    \"not ok\".

A start tag whose tag name is \"form\"

:   If the [`form` element
    pointer](#form-element-pointer){#parsing-main-inbody:form-element-pointer}
    is not null, and there is no
    [`template`{#parsing-main-inbody:the-template-element-5}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-20},
    then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-12}; ignore
    the token.

    Otherwise:

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-21}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-4},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-4}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-5}
    for the token, and, if there is no
    [`template`{#parsing-main-inbody:the-template-element-6}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-22},
    set the [`form` element
    pointer](#form-element-pointer){#parsing-main-inbody:form-element-pointer-2}
    to point to the element created.

A start tag whose tag name is \"li\"

:   Run these steps:

    1.  Set the [frameset-ok
        flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-5}
        to \"not ok\".

    2.  Initialize `node`{.variable} to be the [current
        node](#current-node){#parsing-main-inbody:current-node-4} (the
        bottommost node of the stack).

    3.  *Loop*: If `node`{.variable} is an
        [`li`{#parsing-main-inbody:the-li-element-4}](grouping-content.html#the-li-element)
        element, then run these substeps:

        1.  [Generate implied end
            tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags},
            except for
            [`li`{#parsing-main-inbody:the-li-element-5}](grouping-content.html#the-li-element)
            elements.

        2.  If the [current
            node](#current-node){#parsing-main-inbody:current-node-5} is
            not an
            [`li`{#parsing-main-inbody:the-li-element-6}](grouping-content.html#the-li-element)
            element, then this is a [parse
            error](#parse-errors){#parsing-main-inbody:parse-errors-13}.

        3.  Pop elements from the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-23}
            until an
            [`li`{#parsing-main-inbody:the-li-element-7}](grouping-content.html#the-li-element)
            element has been popped from the stack.

        4.  Jump to the step labeled *done* below.

    4.  If `node`{.variable} is in the
        [special](#special){#parsing-main-inbody:special} category, but
        is not an
        [`address`{#parsing-main-inbody:the-address-element}](sections.html#the-address-element),
        [`div`{#parsing-main-inbody:the-div-element}](grouping-content.html#the-div-element),
        or
        [`p`{#parsing-main-inbody:the-p-element-4}](grouping-content.html#the-p-element)
        element, then jump to the step labeled *done* below.

    5.  Otherwise, set `node`{.variable} to the previous entry in the
        [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-24}
        and return to the step labeled *loop*.

    6.  *Done*: If the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-25}
        [has a `p` element in button
        scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-5},
        then [close a `p`
        element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-5}.

    7.  Finally, [insert an HTML
        element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-6}
        for the token.

A start tag whose tag name is one of: \"dd\", \"dt\"

:   Run these steps:

    1.  Set the [frameset-ok
        flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-6}
        to \"not ok\".

    2.  Initialize `node`{.variable} to be the [current
        node](#current-node){#parsing-main-inbody:current-node-6} (the
        bottommost node of the stack).

    3.  *Loop*: If `node`{.variable} is a
        [`dd`{#parsing-main-inbody:the-dd-element-4}](grouping-content.html#the-dd-element)
        element, then run these substeps:

        1.  [Generate implied end
            tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-2},
            except for
            [`dd`{#parsing-main-inbody:the-dd-element-5}](grouping-content.html#the-dd-element)
            elements.

        2.  If the [current
            node](#current-node){#parsing-main-inbody:current-node-7} is
            not a
            [`dd`{#parsing-main-inbody:the-dd-element-6}](grouping-content.html#the-dd-element)
            element, then this is a [parse
            error](#parse-errors){#parsing-main-inbody:parse-errors-14}.

        3.  Pop elements from the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-26}
            until a
            [`dd`{#parsing-main-inbody:the-dd-element-7}](grouping-content.html#the-dd-element)
            element has been popped from the stack.

        4.  Jump to the step labeled *done* below.

    4.  If `node`{.variable} is a
        [`dt`{#parsing-main-inbody:the-dt-element-4}](grouping-content.html#the-dt-element)
        element, then run these substeps:

        1.  [Generate implied end
            tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-3},
            except for
            [`dt`{#parsing-main-inbody:the-dt-element-5}](grouping-content.html#the-dt-element)
            elements.

        2.  If the [current
            node](#current-node){#parsing-main-inbody:current-node-8} is
            not a
            [`dt`{#parsing-main-inbody:the-dt-element-6}](grouping-content.html#the-dt-element)
            element, then this is a [parse
            error](#parse-errors){#parsing-main-inbody:parse-errors-15}.

        3.  Pop elements from the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-27}
            until a
            [`dt`{#parsing-main-inbody:the-dt-element-7}](grouping-content.html#the-dt-element)
            element has been popped from the stack.

        4.  Jump to the step labeled *done* below.

    5.  If `node`{.variable} is in the
        [special](#special){#parsing-main-inbody:special-2} category,
        but is not an
        [`address`{#parsing-main-inbody:the-address-element-2}](sections.html#the-address-element),
        [`div`{#parsing-main-inbody:the-div-element-2}](grouping-content.html#the-div-element),
        or
        [`p`{#parsing-main-inbody:the-p-element-5}](grouping-content.html#the-p-element)
        element, then jump to the step labeled *done* below.

    6.  Otherwise, set `node`{.variable} to the previous entry in the
        [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-28}
        and return to the step labeled *loop*.

    7.  *Done*: If the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-29}
        [has a `p` element in button
        scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-6},
        then [close a `p`
        element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-6}.

    8.  Finally, [insert an HTML
        element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-7}
        for the token.

A start tag whose tag name is \"plaintext\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-30}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-7},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-7}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-8}
    for the token.

    Switch the tokenizer to the [PLAINTEXT
    state](#plaintext-state){#parsing-main-inbody:plaintext-state}.

    Once a start tag with the tag name \"plaintext\" has been seen, all
    remaining tokens will be character tokens (and a final end-of-file
    token) because there is no way to switch the tokenizer out of the
    [PLAINTEXT
    state](#plaintext-state){#parsing-main-inbody:plaintext-state-2}.
    However, as the tree builder remains in its existing insertion mode,
    it might [reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-3}
    while processing those character tokens. This means that the parser
    can insert other elements into the
    [`plaintext`{#parsing-main-inbody:plaintext}](obsolete.html#plaintext)
    element.

A start tag whose tag name is \"button\"

:   1.  If the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-31}
        [has a `button` element in
        scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-3},
        then run these substeps:

        1.  [Parse
            error](#parse-errors){#parsing-main-inbody:parse-errors-16}.

        2.  [Generate implied end
            tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-4}.

        3.  Pop elements from the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-32}
            until a
            [`button`{#parsing-main-inbody:the-button-element}](form-elements.html#the-button-element)
            element has been popped from the stack.

    2.  [Reconstruct the active formatting
        elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-4},
        if any.

    3.  [Insert an HTML
        element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-9}
        for the token.

    4.  Set the [frameset-ok
        flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-7}
        to \"not ok\".

An end tag whose tag name is one of: \"address\", \"article\", \"aside\", \"blockquote\", \"button\", \"center\", \"details\", \"dialog\", \"dir\", \"div\", \"dl\", \"fieldset\", \"figcaption\", \"figure\", \"footer\", \"header\", \"hgroup\", \"listing\", \"main\", \"menu\", \"nav\", \"ol\", \"pre\", \"search\", \"section\", \"summary\", \"ul\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-33}
    does not [have an element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-4}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-2}
    with the same tag name as that of the token, then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-17}; ignore
    the token.

    Otherwise, run these steps:

    1.  [Generate implied end
        tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-5}.

    2.  If the [current
        node](#current-node){#parsing-main-inbody:current-node-9} is not
        an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-3}
        with the same tag name as that of the token, then this is a
        [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-18}.

    3.  Pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-34}
        until an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-4}
        with the same tag name as the token has been popped from the
        stack.

An end tag whose tag name is \"form\"

:   If there is no
    [`template`{#parsing-main-inbody:the-template-element-7}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-35},
    then run these substeps:

    1.  Let `node`{.variable} be the element that the [`form` element
        pointer](#form-element-pointer){#parsing-main-inbody:form-element-pointer-3}
        is set to, or null if it is not set to an element.

    2.  Set the [`form` element
        pointer](#form-element-pointer){#parsing-main-inbody:form-element-pointer-4}
        to null.

    3.  If `node`{.variable} is null or if the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-36}
        does not [have `node`{.variable} in
        scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-5},
        then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-19};
        return and ignore the token.

    4.  [Generate implied end
        tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-6}.

    5.  If the [current
        node](#current-node){#parsing-main-inbody:current-node-10} is
        not `node`{.variable}, then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-20}.

    6.  Remove `node`{.variable} from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-37}.

    If there *is* a
    [`template`{#parsing-main-inbody:the-template-element-8}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-38},
    then run these substeps instead:

    1.  If the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-39}
        does not [have a `form` element in
        scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-6},
        then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-21};
        return and ignore the token.

    2.  [Generate implied end
        tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-7}.

    3.  If the [current
        node](#current-node){#parsing-main-inbody:current-node-11} is
        not a
        [`form`{#parsing-main-inbody:the-form-element}](forms.html#the-form-element)
        element, then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-22}.

    4.  Pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-40}
        until a
        [`form`{#parsing-main-inbody:the-form-element-2}](forms.html#the-form-element)
        element has been popped from the stack.

An end tag whose tag name is \"p\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-41}
    does not [have a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-8},
    then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-23}; [insert
    an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-10}
    for a \"p\" start tag token with no attributes.

    [Close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-8}.

An end tag whose tag name is \"li\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-42}
    does not [have an `li` element in list item
    scope](#has-an-element-in-list-item-scope){#parsing-main-inbody:has-an-element-in-list-item-scope},
    then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-24}; ignore
    the token.

    Otherwise, run these steps:

    1.  [Generate implied end
        tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-8},
        except for
        [`li`{#parsing-main-inbody:the-li-element-8}](grouping-content.html#the-li-element)
        elements.

    2.  If the [current
        node](#current-node){#parsing-main-inbody:current-node-12} is
        not an
        [`li`{#parsing-main-inbody:the-li-element-9}](grouping-content.html#the-li-element)
        element, then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-25}.

    3.  Pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-43}
        until an
        [`li`{#parsing-main-inbody:the-li-element-10}](grouping-content.html#the-li-element)
        element has been popped from the stack.

An end tag whose tag name is one of: \"dd\", \"dt\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-44}
    does not [have an element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-7}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-5}
    with the same tag name as that of the token, then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-26}; ignore
    the token.

    Otherwise, run these steps:

    1.  [Generate implied end
        tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-9},
        except for [HTML
        elements](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-6}
        with the same tag name as the token.

    2.  If the [current
        node](#current-node){#parsing-main-inbody:current-node-13} is
        not an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-7}
        with the same tag name as that of the token, then this is a
        [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-27}.

    3.  Pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-45}
        until an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-8}
        with the same tag name as the token has been popped from the
        stack.

An end tag whose tag name is one of: \"h1\", \"h2\", \"h3\", \"h4\", \"h5\", \"h6\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-46}
    does not [have an element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-8}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-9}
    and whose tag name is one of \"h1\", \"h2\", \"h3\", \"h4\", \"h5\",
    or \"h6\", then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-28}; ignore
    the token.

    Otherwise, run these steps:

    1.  [Generate implied end
        tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-10}.

    2.  If the [current
        node](#current-node){#parsing-main-inbody:current-node-14} is
        not an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-10}
        with the same tag name as that of the token, then this is a
        [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-29}.

    3.  Pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-47}
        until an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-11}
        whose tag name is one of \"h1\", \"h2\", \"h3\", \"h4\", \"h5\",
        or \"h6\" has been popped from the stack.

An end tag whose tag name is \"sarcasm\"

:   Take a deep breath, then act as described in the \"any other end
    tag\" entry below.

A start tag whose tag name is \"a\"

:   If the [list of active formatting
    elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements}
    contains an
    [`a`{#parsing-main-inbody:the-a-element}](text-level-semantics.html#the-a-element)
    element between the end of the list and the last
    [marker](#concept-parser-marker){#parsing-main-inbody:concept-parser-marker}
    on the list (or the start of the list if there is no
    [marker](#concept-parser-marker){#parsing-main-inbody:concept-parser-marker-2}
    on the list), then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-30}; run the
    [adoption agency
    algorithm](#adoption-agency-algorithm){#parsing-main-inbody:adoption-agency-algorithm}
    for the token, then remove that element from the [list of active
    formatting
    elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-2}
    and the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-48}
    if the [adoption agency
    algorithm](#adoption-agency-algorithm){#parsing-main-inbody:adoption-agency-algorithm-2}
    didn\'t already remove it (it might not have if the element is not
    [in table
    scope](#has-an-element-in-table-scope){#parsing-main-inbody:has-an-element-in-table-scope}).

    In the non-conforming stream
    `<a href="a">a<table><a href="b">b</table>x`, the first
    [`a`{#parsing-main-inbody:the-a-element-2}](text-level-semantics.html#the-a-element)
    element would be closed upon seeing the second one, and the \"x\"
    character would be inside a link to \"b\", not to \"a\". This is
    despite the fact that the outer
    [`a`{#parsing-main-inbody:the-a-element-3}](text-level-semantics.html#the-a-element)
    element is not in table scope (meaning that a regular `</a>` end tag
    at the start of the table wouldn\'t close the outer
    [`a`{#parsing-main-inbody:the-a-element-4}](text-level-semantics.html#the-a-element)
    element). The result is that the two
    [`a`{#parsing-main-inbody:the-a-element-5}](text-level-semantics.html#the-a-element)
    elements are indirectly nested inside each other --- non-conforming
    markup will often result in non-conforming DOMs when parsed.

    [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-5},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-11}
    for the token. [Push onto the list of active formatting
    elements](#push-onto-the-list-of-active-formatting-elements){#parsing-main-inbody:push-onto-the-list-of-active-formatting-elements}
    that element.

A start tag whose tag name is one of: \"b\", \"big\", \"code\", \"em\", \"font\", \"i\", \"s\", \"small\", \"strike\", \"strong\", \"tt\", \"u\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-6},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-12}
    for the token. [Push onto the list of active formatting
    elements](#push-onto-the-list-of-active-formatting-elements){#parsing-main-inbody:push-onto-the-list-of-active-formatting-elements-2}
    that element.

A start tag whose tag name is \"nobr\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-7},
    if any.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-49}
    [has a `nobr` element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-9},
    then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-31}; run the
    [adoption agency
    algorithm](#adoption-agency-algorithm){#parsing-main-inbody:adoption-agency-algorithm-3}
    for the token, then once again [reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-8},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-13}
    for the token. [Push onto the list of active formatting
    elements](#push-onto-the-list-of-active-formatting-elements){#parsing-main-inbody:push-onto-the-list-of-active-formatting-elements-3}
    that element.

An end tag whose tag name is one of: \"a\", \"b\", \"big\", \"code\", \"em\", \"font\", \"i\", \"nobr\", \"s\", \"small\", \"strike\", \"strong\", \"tt\", \"u\"

:   Run the [adoption agency
    algorithm](#adoption-agency-algorithm){#parsing-main-inbody:adoption-agency-algorithm-4}
    for the token.

A start tag whose tag name is one of: \"applet\", \"marquee\", \"object\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-9},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-14}
    for the token.

    Insert a
    [marker](#concept-parser-marker){#parsing-main-inbody:concept-parser-marker-3}
    at the end of the [list of active formatting
    elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-3}.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-8} to
    \"not ok\".

An end tag token whose tag name is one of: \"applet\", \"marquee\", \"object\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-50}
    does not [have an element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-10}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-12}
    with the same tag name as that of the token, then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-32}; ignore
    the token.

    Otherwise, run these steps:

    1.  [Generate implied end
        tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-11}.
    2.  If the [current
        node](#current-node){#parsing-main-inbody:current-node-15} is
        not an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-13}
        with the same tag name as that of the token, then this is a
        [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-33}.
    3.  Pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-51}
        until an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-14}
        with the same tag name as the token has been popped from the
        stack.
    4.  [Clear the list of active formatting elements up to the last
        marker](#clear-the-list-of-active-formatting-elements-up-to-the-last-marker){#parsing-main-inbody:clear-the-list-of-active-formatting-elements-up-to-the-last-marker}.

A start tag whose tag name is \"table\"

:   If the
    [`Document`{#parsing-main-inbody:document}](dom.html#document) is
    *not* set to [quirks
    mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parsing-main-inbody:quirks-mode
    x-internal="quirks-mode"}, and the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-52}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-9},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-9}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-15}
    for the token.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-9} to
    \"not ok\".

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-7} to
    \"[in
    table](#parsing-main-intable){#parsing-main-inbody:parsing-main-intable}\".

An end tag whose tag name is \"br\"

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors-34}.
    Drop the attributes from the token, and act as described in the next
    entry; i.e. act as if this was a \"br\" start tag token with no
    attributes, rather than the end tag token that it actually is.

A start tag whose tag name is one of: \"area\", \"br\", \"embed\", \"img\", \"keygen\", \"wbr\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-10},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-16}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inbody:current-node-16} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-53}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inbody:acknowledge-self-closing-flag},
    if it is set.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-10}
    to \"not ok\".

A start tag whose tag name is \"input\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-11},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-17}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inbody:current-node-17} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-54}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inbody:acknowledge-self-closing-flag-2},
    if it is set.

    If the token does not have an attribute with the name \"type\", or
    if it does, but that attribute\'s value is not an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#parsing-main-inbody:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for the string
    \"`hidden`\", then: set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-11}
    to \"not ok\".

A start tag whose tag name is one of: \"param\", \"source\", \"track\"

:   [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-18}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inbody:current-node-18} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-55}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inbody:acknowledge-self-closing-flag-3},
    if it is set.

A start tag whose tag name is \"hr\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-56}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-10},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-10}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-19}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inbody:current-node-19} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-57}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inbody:acknowledge-self-closing-flag-4},
    if it is set.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-12}
    to \"not ok\".

A start tag whose tag name is \"image\"

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors-35}.
    Change the token\'s tag name to \"img\" and reprocess it. (Don\'t
    ask.)

A start tag whose tag name is \"textarea\"

:   Run these steps:

    1.  [Insert an HTML
        element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-20}
        for the token.

    2.  If the [next
        token](#next-token){#parsing-main-inbody:next-token-2} is a
        U+000A LINE FEED (LF) character token, then ignore that token
        and move on to the next one. (Newlines at the start of
        [`textarea`{#parsing-main-inbody:the-textarea-element}](form-elements.html#the-textarea-element)
        elements are ignored as an authoring convenience.)

    3.  Switch the tokenizer to the [RCDATA
        state](#rcdata-state){#parsing-main-inbody:rcdata-state}.

    4.  Set the [original insertion
        mode](#original-insertion-mode){#parsing-main-inbody:original-insertion-mode}
        to the current [insertion
        mode](#insertion-mode){#parsing-main-inbody:insertion-mode-8}.

    5.  Set the [frameset-ok
        flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-13}
        to \"not ok\".

    6.  Switch the [insertion
        mode](#insertion-mode){#parsing-main-inbody:insertion-mode-9} to
        \"[text](#parsing-main-incdata){#parsing-main-inbody:parsing-main-incdata}\".

A start tag whose tag name is \"xmp\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-58}
    [has a `p` element in button
    scope](#has-an-element-in-button-scope){#parsing-main-inbody:has-an-element-in-button-scope-11},
    then [close a `p`
    element](#close-a-p-element){#parsing-main-inbody:close-a-p-element-11}.

    [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-12},
    if any.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-14}
    to \"not ok\".

    Follow the [generic raw text element parsing
    algorithm](#generic-raw-text-element-parsing-algorithm){#parsing-main-inbody:generic-raw-text-element-parsing-algorithm}.

A start tag whose tag name is \"iframe\"

:   Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-15}
    to \"not ok\".

    Follow the [generic raw text element parsing
    algorithm](#generic-raw-text-element-parsing-algorithm){#parsing-main-inbody:generic-raw-text-element-parsing-algorithm-2}.

A start tag whose tag name is \"noembed\"\
A start tag whose tag name is \"noscript\", if the [scripting flag](#scripting-flag){#parsing-main-inbody:scripting-flag} is enabled

:   Follow the [generic raw text element parsing
    algorithm](#generic-raw-text-element-parsing-algorithm){#parsing-main-inbody:generic-raw-text-element-parsing-algorithm-3}.

A start tag whose tag name is \"select\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-13},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-21}
    for the token.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inbody:frameset-ok-flag-16}
    to \"not ok\".

    If the [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-10} is
    one of \"[in
    table](#parsing-main-intable){#parsing-main-inbody:parsing-main-intable-2}\",
    \"[in
    caption](#parsing-main-incaption){#parsing-main-inbody:parsing-main-incaption}\",
    \"[in table
    body](#parsing-main-intbody){#parsing-main-inbody:parsing-main-intbody}\",
    \"[in
    row](#parsing-main-intr){#parsing-main-inbody:parsing-main-intr}\",
    or \"[in
    cell](#parsing-main-intd){#parsing-main-inbody:parsing-main-intd}\",
    then switch the [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-11} to
    \"[in select in
    table](#parsing-main-inselectintable){#parsing-main-inbody:parsing-main-inselectintable}\".
    Otherwise, switch the [insertion
    mode](#insertion-mode){#parsing-main-inbody:insertion-mode-12} to
    \"[in
    select](#parsing-main-inselect){#parsing-main-inbody:parsing-main-inselect}\".

A start tag whose tag name is one of: \"optgroup\", \"option\"

:   If the [current
    node](#current-node){#parsing-main-inbody:current-node-20} is an
    [`option`{#parsing-main-inbody:the-option-element-4}](form-elements.html#the-option-element)
    element, then pop the [current
    node](#current-node){#parsing-main-inbody:current-node-21} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-59}.

    [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-14},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-22}
    for the token.

A start tag whose tag name is one of: \"rb\", \"rtc\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-60}
    [has a `ruby` element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-11},
    then [generate implied end
    tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-12}.
    If the [current
    node](#current-node){#parsing-main-inbody:current-node-22} is not
    now a
    [`ruby`{#parsing-main-inbody:the-ruby-element}](text-level-semantics.html#the-ruby-element)
    element, this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-36}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-23}
    for the token.

A start tag whose tag name is one of: \"rp\", \"rt\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-61}
    [has a `ruby` element in
    scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-12},
    then [generate implied end
    tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-13},
    except for [`rtc`{#parsing-main-inbody:rtc-4}](obsolete.html#rtc)
    elements. If the [current
    node](#current-node){#parsing-main-inbody:current-node-23} is not
    now a [`rtc`{#parsing-main-inbody:rtc-5}](obsolete.html#rtc) element
    or a
    [`ruby`{#parsing-main-inbody:the-ruby-element-2}](text-level-semantics.html#the-ruby-element)
    element, this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-37}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-24}
    for the token.

A start tag whose tag name is \"math\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-15},
    if any.

    [Adjust MathML
    attributes](#adjust-mathml-attributes){#parsing-main-inbody:adjust-mathml-attributes}
    for the token. (This fixes the case of MathML attributes that are
    not all lowercase.)

    [Adjust foreign
    attributes](#adjust-foreign-attributes){#parsing-main-inbody:adjust-foreign-attributes}
    for the token. (This fixes the use of namespaced attributes, in
    particular XLink.)

    [Insert a foreign
    element](#insert-a-foreign-element){#parsing-main-inbody:insert-a-foreign-element}
    for the token, with [MathML
    namespace](https://infra.spec.whatwg.org/#mathml-namespace){#parsing-main-inbody:mathml-namespace
    x-internal="mathml-namespace"} and false.

    If the token has its *[self-closing flag](#self-closing-flag)* set,
    pop the [current
    node](#current-node){#parsing-main-inbody:current-node-24} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-62}
    and [acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inbody:acknowledge-self-closing-flag-5}.

A start tag whose tag name is \"svg\"

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-16},
    if any.

    [Adjust SVG
    attributes](#adjust-svg-attributes){#parsing-main-inbody:adjust-svg-attributes}
    for the token. (This fixes the case of SVG attributes that are not
    all lowercase.)

    [Adjust foreign
    attributes](#adjust-foreign-attributes){#parsing-main-inbody:adjust-foreign-attributes-2}
    for the token. (This fixes the use of namespaced attributes, in
    particular XLink in SVG.)

    [Insert a foreign
    element](#insert-a-foreign-element){#parsing-main-inbody:insert-a-foreign-element-2}
    for the token, with [SVG
    namespace](https://infra.spec.whatwg.org/#svg-namespace){#parsing-main-inbody:svg-namespace
    x-internal="svg-namespace"} and false.

    If the token has its *[self-closing flag](#self-closing-flag)* set,
    pop the [current
    node](#current-node){#parsing-main-inbody:current-node-25} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-63}
    and [acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inbody:acknowledge-self-closing-flag-6}.

A start tag whose tag name is one of: \"caption\", \"col\", \"colgroup\", \"frame\", \"head\", \"tbody\", \"td\", \"tfoot\", \"th\", \"thead\", \"tr\"

:   [Parse error](#parse-errors){#parsing-main-inbody:parse-errors-38}.
    Ignore the token.

Any other start tag

:   [Reconstruct the active formatting
    elements](#reconstruct-the-active-formatting-elements){#parsing-main-inbody:reconstruct-the-active-formatting-elements-17},
    if any.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inbody:insert-an-html-element-25}
    for the token.

    This element will be an
    [ordinary](#ordinary){#parsing-main-inbody:ordinary} element. With
    one exception: if the [scripting
    flag](#scripting-flag){#parsing-main-inbody:scripting-flag-2} is
    disabled, it can also be a
    [`noscript`{#parsing-main-inbody:the-noscript-element}](scripting.html#the-noscript-element)
    element.

Any other end tag

:   Run these steps:

    1.  Initialize `node`{.variable} to be the [current
        node](#current-node){#parsing-main-inbody:current-node-26} (the
        bottommost node of the stack).

    2.  *Loop*: If `node`{.variable} is an [HTML
        element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-15}
        with the same tag name as the token, then:

        1.  [Generate implied end
            tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-14},
            except for [HTML
            elements](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-16}
            with the same tag name as the token.

        2.  If `node`{.variable} is not the [current
            node](#current-node){#parsing-main-inbody:current-node-27},
            then this is a [parse
            error](#parse-errors){#parsing-main-inbody:parse-errors-39}.

        3.  Pop all the nodes from the [current
            node](#current-node){#parsing-main-inbody:current-node-28}
            up to `node`{.variable}, including `node`{.variable}, then
            stop these steps.

    3.  Otherwise, if `node`{.variable} is in the
        [special](#special){#parsing-main-inbody:special-3} category,
        then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-40};
        ignore the token, and return.

    4.  Set `node`{.variable} to the previous entry in the [stack of
        open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-64}.

    5.  Return to the step labeled *loop*.

When the steps above say the user agent is to [close a `p`
element]{#close-a-p-element .dfn}, it means that the user agent must run
the following steps:

1.  [Generate implied end
    tags](#generate-implied-end-tags){#parsing-main-inbody:generate-implied-end-tags-15},
    except for
    [`p`{#parsing-main-inbody:the-p-element-6}](grouping-content.html#the-p-element)
    elements.

2.  If the [current
    node](#current-node){#parsing-main-inbody:current-node-29} is not a
    [`p`{#parsing-main-inbody:the-p-element-7}](grouping-content.html#the-p-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-inbody:parse-errors-41}.

3.  Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-65}
    until a
    [`p`{#parsing-main-inbody:the-p-element-8}](grouping-content.html#the-p-element)
    element has been popped from the stack.

The [adoption agency algorithm]{#adoption-agency-algorithm .dfn}, which
takes as its only argument a token `token`{.variable} for which the
algorithm is being run, consists of the following steps:

1.  Let `subject`{.variable} be `token`{.variable}\'s tag name.

2.  If the [current
    node](#current-node){#parsing-main-inbody:current-node-30} is an
    [HTML
    element](infrastructure.html#html-elements){#parsing-main-inbody:html-elements-17}
    whose tag name is `subject`{.variable}, and the [current
    node](#current-node){#parsing-main-inbody:current-node-31} is not in
    the [list of active formatting
    elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-4},
    then pop the [current
    node](#current-node){#parsing-main-inbody:current-node-32} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-66}
    and return.

3.  Let `outerLoopCounter`{.variable} be 0.

4.  While true:

    1.  If `outerLoopCounter`{.variable} is greater than or equal to 8,
        then return.

    2.  Increment `outerLoopCounter`{.variable} by 1.

    3.  Let `formattingElement`{.variable} be the last element in the
        [list of active formatting
        elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-5}
        that:

        - is between the end of the list and the last
          [marker](#concept-parser-marker){#parsing-main-inbody:concept-parser-marker-4}
          in the list, if any, or the start of the list otherwise, and
        - has the tag name `subject`{.variable}.

        If there is no such element, then return and instead act as
        described in the \"any other end tag\" entry above.

    4.  If `formattingElement`{.variable} is not in the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-67},
        then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-42};
        remove the element from the list, and return.

    5.  If `formattingElement`{.variable} is in the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-68},
        but the element is not [in
        scope](#has-an-element-in-scope){#parsing-main-inbody:has-an-element-in-scope-13},
        then this is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-43};
        return.

    6.  If `formattingElement`{.variable} is not the [current
        node](#current-node){#parsing-main-inbody:current-node-33}, this
        is a [parse
        error](#parse-errors){#parsing-main-inbody:parse-errors-44}.
        (But do not return.)

    7.  Let `furthestBlock`{.variable} be the topmost node in the [stack
        of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-69}
        that is lower in the stack than `formattingElement`{.variable},
        and is an element in the
        [special](#special){#parsing-main-inbody:special-4} category.
        There might not be one.

    8.  If there is no `furthestBlock`{.variable}, then the UA must
        first pop all the nodes from the bottom of the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-70},
        from the [current
        node](#current-node){#parsing-main-inbody:current-node-34} up to
        and including `formattingElement`{.variable}, then remove
        `formattingElement`{.variable} from the [list of active
        formatting
        elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-6},
        and finally return.

    9.  Let `commonAncestor`{.variable} be the element immediately above
        `formattingElement`{.variable} in the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-71}.

    10. Let a bookmark note the position of
        `formattingElement`{.variable} in the [list of active formatting
        elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-7}
        relative to the elements on either side of it in the list.

    11. Let `node`{.variable} and `lastNode`{.variable} be
        `furthestBlock`{.variable}.

    12. Let `innerLoopCounter`{.variable} be 0.

    13. While true:

        1.  Increment `innerLoopCounter`{.variable} by 1.

        2.  Let `node`{.variable} be the element immediately above
            `node`{.variable} in the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-72},
            or if `node`{.variable} is no longer in the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-73}
            (e.g. because it got removed by this algorithm), the element
            that was immediately above `node`{.variable} in the [stack
            of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-74}
            before `node`{.variable} was removed.

        3.  If `node`{.variable} is `formattingElement`{.variable}, then
            [break](https://infra.spec.whatwg.org/#iteration-break){#parsing-main-inbody:break
            x-internal="break"}.

        4.  If `innerLoopCounter`{.variable} is greater than 3 and
            `node`{.variable} is in the [list of active formatting
            elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-8},
            then remove `node`{.variable} from the [list of active
            formatting
            elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-9}.

        5.  If `node`{.variable} is not in the [list of active
            formatting
            elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-10},
            then remove `node`{.variable} from the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-75}
            and
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#parsing-main-inbody:continue
            x-internal="continue"}.

        6.  [Create an element for the
            token](#create-an-element-for-the-token){#parsing-main-inbody:create-an-element-for-the-token}
            for which the element `node`{.variable} was created, in the
            [HTML
            namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing-main-inbody:html-namespace-2
            x-internal="html-namespace-2"}, with
            `commonAncestor`{.variable} as the intended parent; replace
            the entry for `node`{.variable} in the [list of active
            formatting
            elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-11}
            with an entry for the new element, replace the entry for
            `node`{.variable} in the [stack of open
            elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-76}
            with an entry for the new element, and let `node`{.variable}
            be the new element.

        7.  If `lastNode`{.variable} is `furthestBlock`{.variable}, then
            move the aforementioned bookmark to be immediately after the
            new `node`{.variable} in the [list of active formatting
            elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-12}.

        8.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#parsing-main-inbody:concept-node-append
            x-internal="concept-node-append"} `lastNode`{.variable} to
            `node`{.variable}.

        9.  Set `lastNode`{.variable} to `node`{.variable}.

    14. Insert whatever `lastNode`{.variable} ended up being in the
        previous step at the [appropriate place for inserting a
        node](#appropriate-place-for-inserting-a-node){#parsing-main-inbody:appropriate-place-for-inserting-a-node},
        but using `commonAncestor`{.variable} as the *override target*.

    15. [Create an element for the
        token](#create-an-element-for-the-token){#parsing-main-inbody:create-an-element-for-the-token-2}
        for which `formattingElement`{.variable} was created, in the
        [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing-main-inbody:html-namespace-2-2
        x-internal="html-namespace-2"}, with `furthestBlock`{.variable}
        as the intended parent.

    16. Take all of the child nodes of `furthestBlock`{.variable} and
        append them to the element created in the last step.

    17. Append that new element to `furthestBlock`{.variable}.

    18. Remove `formattingElement`{.variable} from the [list of active
        formatting
        elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-13},
        and insert the new element into the [list of active formatting
        elements](#list-of-active-formatting-elements){#parsing-main-inbody:list-of-active-formatting-elements-14}
        at the position of the aforementioned bookmark.

    19. Remove `formattingElement`{.variable} from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-77},
        and insert the new element into the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inbody:stack-of-open-elements-78}
        immediately below the position of `furthestBlock`{.variable} in
        that stack.

This algorithm\'s name, the \"adoption agency algorithm\", comes from
the way it causes elements to change parents, and is in contrast with
[other possible
algorithms](https://ln.hixie.ch/?start=1037910467&count=1) for dealing
with misnested content.

###### [13.2.6.4.8]{.secno} The \"[text]{.dfn}\" insertion mode[](#parsing-main-incdata){.self-link} {#parsing-main-incdata}

When the user agent is to apply the rules for the
\"[text](#parsing-main-incdata){#parsing-main-incdata:parsing-main-incdata}\"
[insertion mode](#insertion-mode){#parsing-main-incdata:insertion-mode},
the user agent must handle the token as follows:

A character token

:   [Insert the token\'s
    character](#insert-a-character){#parsing-main-incdata:insert-a-character}.

    This can never be a U+0000 NULL character; the tokenizer converts
    those to U+FFFD REPLACEMENT CHARACTER characters.

An end-of-file token

:   [Parse error](#parse-errors){#parsing-main-incdata:parse-errors}.

    If the [current
    node](#current-node){#parsing-main-incdata:current-node} is a
    [`script`{#parsing-main-incdata:the-script-element}](scripting.html#the-script-element)
    element, then set its [already
    started](scripting.html#already-started){#parsing-main-incdata:already-started}
    to true.

    Pop the [current
    node](#current-node){#parsing-main-incdata:current-node-2} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-incdata:stack-of-open-elements}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-incdata:insertion-mode-2} to
    the [original insertion
    mode](#original-insertion-mode){#parsing-main-incdata:original-insertion-mode}
    and reprocess the token.

An end tag whose tag name is \"script\"

:   If the [active speculative HTML
    parser](#active-speculative-html-parser){#parsing-main-incdata:active-speculative-html-parser}
    is null and the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#parsing-main-incdata:javascript-execution-context-stack
    x-internal="javascript-execution-context-stack"} is empty, then
    [perform a microtask
    checkpoint](webappapis.html#perform-a-microtask-checkpoint){#parsing-main-incdata:perform-a-microtask-checkpoint}.

    Let `script`{.variable} be the [current
    node](#current-node){#parsing-main-incdata:current-node-3} (which
    will be a
    [`script`{#parsing-main-incdata:the-script-element-2}](scripting.html#the-script-element)
    element).

    Pop the [current
    node](#current-node){#parsing-main-incdata:current-node-4} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-incdata:stack-of-open-elements-2}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-incdata:insertion-mode-3} to
    the [original insertion
    mode](#original-insertion-mode){#parsing-main-incdata:original-insertion-mode-2}.

    Let the `old insertion point`{.variable} have the same value as the
    current [insertion
    point](#insertion-point){#parsing-main-incdata:insertion-point}. Let
    the [insertion
    point](#insertion-point){#parsing-main-incdata:insertion-point-2} be
    just before the [next input
    character](#next-input-character){#parsing-main-incdata:next-input-character}.

    Increment the parser\'s [script nesting
    level](#script-nesting-level){#parsing-main-incdata:script-nesting-level}
    by one.

    If the [active speculative HTML
    parser](#active-speculative-html-parser){#parsing-main-incdata:active-speculative-html-parser-2}
    is null, then [prepare the script
    element](scripting.html#prepare-the-script-element){#parsing-main-incdata:prepare-the-script-element}
    `script`{.variable}. This might cause some script to execute, which
    might cause [new characters to be inserted into the
    tokenizer](dynamic-markup-insertion.html#dom-document-write){#parsing-main-incdata:dom-document-write},
    and might cause the tokenizer to output more tokens, resulting in a
    [reentrant invocation of the parser](#nestedParsing).

    Decrement the parser\'s [script nesting
    level](#script-nesting-level){#parsing-main-incdata:script-nesting-level-2}
    by one. If the parser\'s [script nesting
    level](#script-nesting-level){#parsing-main-incdata:script-nesting-level-3}
    is zero, then set the [parser pause
    flag](#parser-pause-flag){#parsing-main-incdata:parser-pause-flag}
    to false.

    Let the [insertion
    point](#insertion-point){#parsing-main-incdata:insertion-point-3}
    have the value of the `old insertion point`{.variable}. (In other
    words, restore the [insertion
    point](#insertion-point){#parsing-main-incdata:insertion-point-4} to
    its previous value. This value might be the \"undefined\" value.)

    At this stage, if the [pending parsing-blocking
    script](scripting.html#pending-parsing-blocking-script){#parsing-main-incdata:pending-parsing-blocking-script}
    is not null, then:

    If the [script nesting level](#script-nesting-level){#parsing-main-incdata:script-nesting-level-4} is not zero:

    :   Set the [parser pause
        flag](#parser-pause-flag){#parsing-main-incdata:parser-pause-flag-2}
        to true, and abort the processing of any nested invocations of
        the tokenizer, yielding control back to the caller.
        (Tokenization will resume when the caller returns to the
        \"outer\" tree construction stage.)

        The tree construction stage of this particular parser is [being
        called reentrantly](#nestedParsing), say from a call to
        [`document.write()`{#parsing-main-incdata:dom-document-write-2}](dynamic-markup-insertion.html#dom-document-write).

    Otherwise:

    :   While the [pending parsing-blocking
        script](scripting.html#pending-parsing-blocking-script){#parsing-main-incdata:pending-parsing-blocking-script-2}
        is not null:

        1.  Let `the script`{.variable} be the [pending parsing-blocking
            script](scripting.html#pending-parsing-blocking-script){#parsing-main-incdata:pending-parsing-blocking-script-3}.

        2.  Set the [pending parsing-blocking
            script](scripting.html#pending-parsing-blocking-script){#parsing-main-incdata:pending-parsing-blocking-script-4}
            to null.

        3.  [Start the speculative HTML
            parser](#start-the-speculative-html-parser){#parsing-main-incdata:start-the-speculative-html-parser}
            for this instance of the HTML parser.

        4.  Block the
            [tokenizer](#tokenization){#parsing-main-incdata:tokenization}
            for this instance of the [HTML
            parser](#html-parser){#parsing-main-incdata:html-parser},
            such that the [event
            loop](webappapis.html#event-loop){#parsing-main-incdata:event-loop}
            will not run
            [tasks](webappapis.html#concept-task){#parsing-main-incdata:concept-task}
            that invoke the
            [tokenizer](#tokenization){#parsing-main-incdata:tokenization-2}.

        5.  If the parser\'s
            [`Document`{#parsing-main-incdata:document}](dom.html#document)
            [has a style sheet that is blocking
            scripts](semantics.html#has-a-style-sheet-that-is-blocking-scripts){#parsing-main-incdata:has-a-style-sheet-that-is-blocking-scripts}
            or `the script`{.variable}\'s [ready to be
            parser-executed](scripting.html#ready-to-be-parser-executed){#parsing-main-incdata:ready-to-be-parser-executed}
            is false: [spin the event
            loop](webappapis.html#spin-the-event-loop){#parsing-main-incdata:spin-the-event-loop}
            until the parser\'s
            [`Document`{#parsing-main-incdata:document-2}](dom.html#document)
            [has no style sheet that is blocking
            scripts](semantics.html#has-no-style-sheet-that-is-blocking-scripts){#parsing-main-incdata:has-no-style-sheet-that-is-blocking-scripts}
            and `the script`{.variable}\'s [ready to be
            parser-executed](scripting.html#ready-to-be-parser-executed){#parsing-main-incdata:ready-to-be-parser-executed-2}
            becomes true.

        6.  If this [parser has been
            aborted](#abort-a-parser){#parsing-main-incdata:abort-a-parser}
            in the meantime, return.

            This could happen if, e.g., while the [spin the event
            loop](webappapis.html#spin-the-event-loop){#parsing-main-incdata:spin-the-event-loop-2}
            algorithm is running, the
            [`Document`{#parsing-main-incdata:document-3}](dom.html#document)
            gets
            [destroyed](document-lifecycle.html#destroy-a-document){#parsing-main-incdata:destroy-a-document},
            or the
            [`document.open()`{#parsing-main-incdata:dom-document-open}](dynamic-markup-insertion.html#dom-document-open)
            method gets invoked on the
            [`Document`{#parsing-main-incdata:document-4}](dom.html#document).

        7.  [Stop the speculative HTML
            parser](#stop-the-speculative-html-parser){#parsing-main-incdata:stop-the-speculative-html-parser}
            for this instance of the HTML parser.

        8.  Unblock the
            [tokenizer](#tokenization){#parsing-main-incdata:tokenization-3}
            for this instance of the [HTML
            parser](#html-parser){#parsing-main-incdata:html-parser-2},
            such that
            [tasks](webappapis.html#concept-task){#parsing-main-incdata:concept-task-2}
            that invoke the
            [tokenizer](#tokenization){#parsing-main-incdata:tokenization-4}
            can again be run.

        9.  Let the [insertion
            point](#insertion-point){#parsing-main-incdata:insertion-point-5}
            be just before the [next input
            character](#next-input-character){#parsing-main-incdata:next-input-character-2}.

        10. Increment the parser\'s [script nesting
            level](#script-nesting-level){#parsing-main-incdata:script-nesting-level-5}
            by one (it should be zero before this step, so this sets it
            to one).

        11. [Execute the script
            element](scripting.html#execute-the-script-element){#parsing-main-incdata:execute-the-script-element}
            `the script`{.variable}.

        12. Decrement the parser\'s [script nesting
            level](#script-nesting-level){#parsing-main-incdata:script-nesting-level-6}
            by one. If the parser\'s [script nesting
            level](#script-nesting-level){#parsing-main-incdata:script-nesting-level-7}
            is zero (which it always should be at this point), then set
            the [parser pause
            flag](#parser-pause-flag){#parsing-main-incdata:parser-pause-flag-3}
            to false.

        13. Let the [insertion
            point](#insertion-point){#parsing-main-incdata:insertion-point-6}
            be undefined again.

Any other end tag

:   Pop the [current
    node](#current-node){#parsing-main-incdata:current-node-5} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-incdata:stack-of-open-elements-3}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-incdata:insertion-mode-4} to
    the [original insertion
    mode](#original-insertion-mode){#parsing-main-incdata:original-insertion-mode-3}.

###### [13.2.6.4.9]{.secno} The \"[in table]{.dfn}\" insertion mode[](#parsing-main-intable){.self-link} {#parsing-main-intable}

When the user agent is to apply the rules for the \"[in
table](#parsing-main-intable){#parsing-main-intable:parsing-main-intable}\"
[insertion mode](#insertion-mode){#parsing-main-intable:insertion-mode},
the user agent must handle the token as follows:

A character token, if the [current node](#current-node){#parsing-main-intable:current-node} is [`table`{#parsing-main-intable:the-table-element}](tables.html#the-table-element), [`tbody`{#parsing-main-intable:the-tbody-element}](tables.html#the-tbody-element), [`template`{#parsing-main-intable:the-template-element}](scripting.html#the-template-element), [`tfoot`{#parsing-main-intable:the-tfoot-element}](tables.html#the-tfoot-element), [`thead`{#parsing-main-intable:the-thead-element}](tables.html#the-thead-element), or [`tr`{#parsing-main-intable:the-tr-element}](tables.html#the-tr-element) element

:   Let the
    [`pending table character tokens`{.variable}]{#concept-pending-table-char-tokens
    .dfn} be an empty list of tokens.

    Set the [original insertion
    mode](#original-insertion-mode){#parsing-main-intable:original-insertion-mode}
    to the current [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-2}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-3} to
    \"[in table
    text](#parsing-main-intabletext){#parsing-main-intable:parsing-main-intabletext}\"
    and reprocess the token.

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-intable:insert-a-comment}.

A DOCTYPE token

:   [Parse error](#parse-errors){#parsing-main-intable:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"caption\"

:   [Clear the stack back to a table
    context](#clear-the-stack-back-to-a-table-context){#parsing-main-intable:clear-the-stack-back-to-a-table-context}.
    (See below.)

    Insert a
    [marker](#concept-parser-marker){#parsing-main-intable:concept-parser-marker}
    at the end of the [list of active formatting
    elements](#list-of-active-formatting-elements){#parsing-main-intable:list-of-active-formatting-elements}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intable:insert-an-html-element}
    for the token, then switch the [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-4} to
    \"[in
    caption](#parsing-main-incaption){#parsing-main-intable:parsing-main-incaption}\".

A start tag whose tag name is \"colgroup\"

:   [Clear the stack back to a table
    context](#clear-the-stack-back-to-a-table-context){#parsing-main-intable:clear-the-stack-back-to-a-table-context-2}.
    (See below.)

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intable:insert-an-html-element-2}
    for the token, then switch the [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-5} to
    \"[in column
    group](#parsing-main-incolgroup){#parsing-main-intable:parsing-main-incolgroup}\".

A start tag whose tag name is \"col\"

:   [Clear the stack back to a table
    context](#clear-the-stack-back-to-a-table-context){#parsing-main-intable:clear-the-stack-back-to-a-table-context-3}.
    (See below.)

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intable:insert-an-html-element-3}
    for a \"colgroup\" start tag token with no attributes, then switch
    the [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-6} to
    \"[in column
    group](#parsing-main-incolgroup){#parsing-main-intable:parsing-main-incolgroup-2}\".

    Reprocess the current token.

A start tag whose tag name is one of: \"tbody\", \"tfoot\", \"thead\"

:   [Clear the stack back to a table
    context](#clear-the-stack-back-to-a-table-context){#parsing-main-intable:clear-the-stack-back-to-a-table-context-4}.
    (See below.)

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intable:insert-an-html-element-4}
    for the token, then switch the [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-7} to
    \"[in table
    body](#parsing-main-intbody){#parsing-main-intable:parsing-main-intbody}\".

A start tag whose tag name is one of: \"td\", \"th\", \"tr\"

:   [Clear the stack back to a table
    context](#clear-the-stack-back-to-a-table-context){#parsing-main-intable:clear-the-stack-back-to-a-table-context-5}.
    (See below.)

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intable:insert-an-html-element-5}
    for a \"tbody\" start tag token with no attributes, then switch the
    [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-8} to
    \"[in table
    body](#parsing-main-intbody){#parsing-main-intable:parsing-main-intbody-2}\".

    Reprocess the current token.

A start tag whose tag name is \"table\"

:   [Parse error](#parse-errors){#parsing-main-intable:parse-errors-2}.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intable:stack-of-open-elements}
    does not [have a `table` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intable:has-an-element-in-table-scope},
    ignore the token.

    Otherwise:

    Pop elements from this stack until a
    [`table`{#parsing-main-intable:the-table-element-2}](tables.html#the-table-element)
    element has been popped from the stack.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-intable:reset-the-insertion-mode-appropriately}.

    Reprocess the token.

An end tag whose tag name is \"table\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intable:stack-of-open-elements-2}
    does not [have a `table` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intable:has-an-element-in-table-scope-2},
    this is a [parse
    error](#parse-errors){#parsing-main-intable:parse-errors-3}; ignore
    the token.

    Otherwise:

    Pop elements from this stack until a
    [`table`{#parsing-main-intable:the-table-element-3}](tables.html#the-table-element)
    element has been popped from the stack.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-intable:reset-the-insertion-mode-appropriately-2}.

An end tag whose tag name is one of: \"body\", \"caption\", \"col\", \"colgroup\", \"html\", \"tbody\", \"td\", \"tfoot\", \"th\", \"thead\", \"tr\"

:   [Parse error](#parse-errors){#parsing-main-intable:parse-errors-4}.
    Ignore the token.

A start tag whose tag name is one of: \"style\", \"script\", \"template\"\
An end tag whose tag name is \"template\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intable:using-the-rules-for}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-intable:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-9}.

A start tag whose tag name is \"input\"

:   If the token does not have an attribute with the name \"type\", or
    if it does, but that attribute\'s value is not an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#parsing-main-intable:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for the string
    \"`hidden`\", then: act as described in the \"anything else\" entry
    below.

    Otherwise:

    [Parse error](#parse-errors){#parsing-main-intable:parse-errors-5}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intable:insert-an-html-element-6}
    for the token.

    Pop that
    [`input`{#parsing-main-intable:the-input-element}](input.html#the-input-element)
    element off the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intable:stack-of-open-elements-3}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-intable:acknowledge-self-closing-flag},
    if it is set.

A start tag whose tag name is \"form\"

:   [Parse error](#parse-errors){#parsing-main-intable:parse-errors-6}.

    If there is a
    [`template`{#parsing-main-intable:the-template-element-2}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intable:stack-of-open-elements-4},
    or if the [`form` element
    pointer](#form-element-pointer){#parsing-main-intable:form-element-pointer}
    is not null, ignore the token.

    Otherwise:

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intable:insert-an-html-element-7}
    for the token, and set the [`form` element
    pointer](#form-element-pointer){#parsing-main-intable:form-element-pointer-2}
    to point to the element created.

    Pop that
    [`form`{#parsing-main-intable:the-form-element}](forms.html#the-form-element)
    element off the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intable:stack-of-open-elements-5}.

An end-of-file token

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intable:using-the-rules-for-2}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-intable:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-10}.

Anything else

:   [Parse error](#parse-errors){#parsing-main-intable:parse-errors-7}.
    Enable [foster
    parenting](#foster-parent){#parsing-main-intable:foster-parent},
    process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intable:using-the-rules-for-3}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-intable:parsing-main-inbody-2}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intable:insertion-mode-11}, and
    then disable [foster
    parenting](#foster-parent){#parsing-main-intable:foster-parent-2}.

When the steps above require the UA to [clear the stack back to a table
context]{#clear-the-stack-back-to-a-table-context .dfn}, it means that
the UA must, while the [current
node](#current-node){#parsing-main-intable:current-node-2} is not a
[`table`{#parsing-main-intable:the-table-element-4}](tables.html#the-table-element),
[`template`{#parsing-main-intable:the-template-element-3}](scripting.html#the-template-element),
or
[`html`{#parsing-main-intable:the-html-element}](semantics.html#the-html-element)
element, pop elements from the [stack of open
elements](#stack-of-open-elements){#parsing-main-intable:stack-of-open-elements-6}.

This is the same list of elements as used in the *[has an element in
table scope](#has-an-element-in-table-scope)* steps.

The [current node](#current-node){#parsing-main-intable:current-node-3}
being an
[`html`{#parsing-main-intable:the-html-element-2}](semantics.html#the-html-element)
element after this process is a [fragment
case](#fragment-case){#parsing-main-intable:fragment-case}.

###### [13.2.6.4.10]{.secno} The \"[in table text]{.dfn}\" insertion mode[](#parsing-main-intabletext){.self-link} {#parsing-main-intabletext}

When the user agent is to apply the rules for the \"[in table
text](#parsing-main-intabletext){#parsing-main-intabletext:parsing-main-intabletext}\"
[insertion
mode](#insertion-mode){#parsing-main-intabletext:insertion-mode}, the
user agent must handle the token as follows:

A character token that is U+0000 NULL

:   [Parse
    error](#parse-errors){#parsing-main-intabletext:parse-errors}.
    Ignore the token.

Any other character token

:   Append the character token to the
    [`pending table character tokens`{#parsing-main-intabletext:concept-pending-table-char-tokens
    .variable}](#concept-pending-table-char-tokens) list.

Anything else

:   If any of the tokens in the
    [`pending table character tokens`{#parsing-main-intabletext:concept-pending-table-char-tokens-2
    .variable}](#concept-pending-table-char-tokens) list are character
    tokens that are not [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#parsing-main-intabletext:space-characters
    x-internal="space-characters"}, then this is a [parse
    error](#parse-errors){#parsing-main-intabletext:parse-errors-2}:
    reprocess the character tokens in the
    [`pending table character tokens`{#parsing-main-intabletext:concept-pending-table-char-tokens-3
    .variable}](#concept-pending-table-char-tokens) list using the rules
    given in the \"anything else\" entry in the \"[in
    table](#parsing-main-intable){#parsing-main-intabletext:parsing-main-intable}\"
    insertion mode.

    Otherwise, [insert the
    characters](#insert-a-character){#parsing-main-intabletext:insert-a-character}
    given by the
    [`pending table character tokens`{#parsing-main-intabletext:concept-pending-table-char-tokens-4
    .variable}](#concept-pending-table-char-tokens) list.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intabletext:insertion-mode-2}
    to the [original insertion
    mode](#original-insertion-mode){#parsing-main-intabletext:original-insertion-mode}
    and reprocess the token.

###### [13.2.6.4.11]{.secno} The \"[in caption]{.dfn}\" insertion mode[](#parsing-main-incaption){.self-link} {#parsing-main-incaption}

When the user agent is to apply the rules for the \"[in
caption](#parsing-main-incaption){#parsing-main-incaption:parsing-main-incaption}\"
[insertion
mode](#insertion-mode){#parsing-main-incaption:insertion-mode}, the user
agent must handle the token as follows:

An end tag whose tag name is \"caption\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-incaption:stack-of-open-elements}
    does not [have a `caption` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-incaption:has-an-element-in-table-scope},
    this is a [parse
    error](#parse-errors){#parsing-main-incaption:parse-errors}; ignore
    the token. ([fragment
    case](#fragment-case){#parsing-main-incaption:fragment-case})

    Otherwise:

    [Generate implied end
    tags](#generate-implied-end-tags){#parsing-main-incaption:generate-implied-end-tags}.

    Now, if the [current
    node](#current-node){#parsing-main-incaption:current-node} is not a
    [`caption`{#parsing-main-incaption:the-caption-element}](tables.html#the-caption-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-incaption:parse-errors-2}.

    Pop elements from this stack until a
    [`caption`{#parsing-main-incaption:the-caption-element-2}](tables.html#the-caption-element)
    element has been popped from the stack.

    [Clear the list of active formatting elements up to the last
    marker](#clear-the-list-of-active-formatting-elements-up-to-the-last-marker){#parsing-main-incaption:clear-the-list-of-active-formatting-elements-up-to-the-last-marker}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-incaption:insertion-mode-2} to
    \"[in
    table](#parsing-main-intable){#parsing-main-incaption:parsing-main-intable}\".

A start tag whose tag name is one of: \"caption\", \"col\", \"colgroup\", \"tbody\", \"td\", \"tfoot\", \"th\", \"thead\", \"tr\"\
An end tag whose tag name is \"table\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-incaption:stack-of-open-elements-2}
    does not [have a `caption` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-incaption:has-an-element-in-table-scope-2},
    this is a [parse
    error](#parse-errors){#parsing-main-incaption:parse-errors-3};
    ignore the token. ([fragment
    case](#fragment-case){#parsing-main-incaption:fragment-case-2})

    Otherwise:

    [Generate implied end
    tags](#generate-implied-end-tags){#parsing-main-incaption:generate-implied-end-tags-2}.

    Now, if the [current
    node](#current-node){#parsing-main-incaption:current-node-2} is not
    a
    [`caption`{#parsing-main-incaption:the-caption-element-3}](tables.html#the-caption-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-incaption:parse-errors-4}.

    Pop elements from this stack until a
    [`caption`{#parsing-main-incaption:the-caption-element-4}](tables.html#the-caption-element)
    element has been popped from the stack.

    [Clear the list of active formatting elements up to the last
    marker](#clear-the-list-of-active-formatting-elements-up-to-the-last-marker){#parsing-main-incaption:clear-the-list-of-active-formatting-elements-up-to-the-last-marker-2}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-incaption:insertion-mode-3} to
    \"[in
    table](#parsing-main-intable){#parsing-main-incaption:parsing-main-intable-2}\".

    Reprocess the token.

An end tag whose tag name is one of: \"body\", \"col\", \"colgroup\", \"html\", \"tbody\", \"td\", \"tfoot\", \"th\", \"thead\", \"tr\"

:   [Parse
    error](#parse-errors){#parsing-main-incaption:parse-errors-5}.
    Ignore the token.

Anything else

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-incaption:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-incaption:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-incaption:insertion-mode-4}.

###### [13.2.6.4.12]{.secno} The \"[in column group]{.dfn}\" insertion mode[](#parsing-main-incolgroup){.self-link} {#parsing-main-incolgroup}

When the user agent is to apply the rules for the \"[in column
group](#parsing-main-incolgroup){#parsing-main-incolgroup:parsing-main-incolgroup}\"
[insertion
mode](#insertion-mode){#parsing-main-incolgroup:insertion-mode}, the
user agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   [Insert the
    character](#insert-a-character){#parsing-main-incolgroup:insert-a-character}.

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-incolgroup:insert-a-comment}.

A DOCTYPE token

:   [Parse error](#parse-errors){#parsing-main-incolgroup:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-incolgroup:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-incolgroup:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-incolgroup:insertion-mode-2}.

A start tag whose tag name is \"col\"

:   [Insert an HTML
    element](#insert-an-html-element){#parsing-main-incolgroup:insert-an-html-element}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-incolgroup:current-node} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-incolgroup:stack-of-open-elements}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-incolgroup:acknowledge-self-closing-flag},
    if it is set.

An end tag whose tag name is \"colgroup\"

:   If the [current
    node](#current-node){#parsing-main-incolgroup:current-node-2} is not
    a
    [`colgroup`{#parsing-main-incolgroup:the-colgroup-element}](tables.html#the-colgroup-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-incolgroup:parse-errors-2};
    ignore the token.

    Otherwise, pop the [current
    node](#current-node){#parsing-main-incolgroup:current-node-3} from
    the [stack of open
    elements](#stack-of-open-elements){#parsing-main-incolgroup:stack-of-open-elements-2}.
    Switch the [insertion
    mode](#insertion-mode){#parsing-main-incolgroup:insertion-mode-3} to
    \"[in
    table](#parsing-main-intable){#parsing-main-incolgroup:parsing-main-intable}\".

An end tag whose tag name is \"col\"

:   [Parse
    error](#parse-errors){#parsing-main-incolgroup:parse-errors-3}.
    Ignore the token.

A start tag whose tag name is \"template\"\
An end tag whose tag name is \"template\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-incolgroup:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-incolgroup:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#parsing-main-incolgroup:insertion-mode-4}.

An end-of-file token

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-incolgroup:using-the-rules-for-3}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-incolgroup:parsing-main-inbody-2}\"
    [insertion
    mode](#insertion-mode){#parsing-main-incolgroup:insertion-mode-5}.

Anything else

:   If the [current
    node](#current-node){#parsing-main-incolgroup:current-node-4} is not
    a
    [`colgroup`{#parsing-main-incolgroup:the-colgroup-element-2}](tables.html#the-colgroup-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-incolgroup:parse-errors-4};
    ignore the token.

    Otherwise, pop the [current
    node](#current-node){#parsing-main-incolgroup:current-node-5} from
    the [stack of open
    elements](#stack-of-open-elements){#parsing-main-incolgroup:stack-of-open-elements-3}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-incolgroup:insertion-mode-6} to
    \"[in
    table](#parsing-main-intable){#parsing-main-incolgroup:parsing-main-intable-2}\".

    Reprocess the token.

###### [13.2.6.4.13]{.secno} The \"[in table body]{.dfn}\" insertion mode[](#parsing-main-intbody){.self-link} {#parsing-main-intbody}

When the user agent is to apply the rules for the \"[in table
body](#parsing-main-intbody){#parsing-main-intbody:parsing-main-intbody}\"
[insertion mode](#insertion-mode){#parsing-main-intbody:insertion-mode},
the user agent must handle the token as follows:

A start tag whose tag name is \"tr\"

:   [Clear the stack back to a table body
    context](#clear-the-stack-back-to-a-table-body-context){#parsing-main-intbody:clear-the-stack-back-to-a-table-body-context}.
    (See below.)

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intbody:insert-an-html-element}
    for the token, then switch the [insertion
    mode](#insertion-mode){#parsing-main-intbody:insertion-mode-2} to
    \"[in
    row](#parsing-main-intr){#parsing-main-intbody:parsing-main-intr}\".

A start tag whose tag name is one of: \"th\", \"td\"

:   [Parse error](#parse-errors){#parsing-main-intbody:parse-errors}.

    [Clear the stack back to a table body
    context](#clear-the-stack-back-to-a-table-body-context){#parsing-main-intbody:clear-the-stack-back-to-a-table-body-context-2}.
    (See below.)

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intbody:insert-an-html-element-2}
    for a \"tr\" start tag token with no attributes, then switch the
    [insertion
    mode](#insertion-mode){#parsing-main-intbody:insertion-mode-3} to
    \"[in
    row](#parsing-main-intr){#parsing-main-intbody:parsing-main-intr-2}\".

    Reprocess the current token.

An end tag whose tag name is one of: \"tbody\", \"tfoot\", \"thead\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intbody:stack-of-open-elements}
    does not [have an element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intbody:has-an-element-in-table-scope}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-intbody:html-elements}
    with the same tag name as the token, this is a [parse
    error](#parse-errors){#parsing-main-intbody:parse-errors-2}; ignore
    the token.

    Otherwise:

    [Clear the stack back to a table body
    context](#clear-the-stack-back-to-a-table-body-context){#parsing-main-intbody:clear-the-stack-back-to-a-table-body-context-3}.
    (See below.)

    Pop the [current
    node](#current-node){#parsing-main-intbody:current-node} from the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-intbody:stack-of-open-elements-2}.
    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intbody:insertion-mode-4} to
    \"[in
    table](#parsing-main-intable){#parsing-main-intbody:parsing-main-intable}\".

A start tag whose tag name is one of: \"caption\", \"col\", \"colgroup\", \"tbody\", \"tfoot\", \"thead\"\
An end tag whose tag name is \"table\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intbody:stack-of-open-elements-3}
    does not [have a `tbody`, `thead`, or `tfoot` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intbody:has-an-element-in-table-scope-2},
    this is a [parse
    error](#parse-errors){#parsing-main-intbody:parse-errors-3}; ignore
    the token.

    Otherwise:

    [Clear the stack back to a table body
    context](#clear-the-stack-back-to-a-table-body-context){#parsing-main-intbody:clear-the-stack-back-to-a-table-body-context-4}.
    (See below.)

    Pop the [current
    node](#current-node){#parsing-main-intbody:current-node-2} from the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-intbody:stack-of-open-elements-4}.
    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intbody:insertion-mode-5} to
    \"[in
    table](#parsing-main-intable){#parsing-main-intbody:parsing-main-intable-2}\".

    Reprocess the token.

An end tag whose tag name is one of: \"body\", \"caption\", \"col\", \"colgroup\", \"html\", \"td\", \"th\", \"tr\"

:   [Parse error](#parse-errors){#parsing-main-intbody:parse-errors-4}.
    Ignore the token.

Anything else

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intbody:using-the-rules-for}
    the \"[in
    table](#parsing-main-intable){#parsing-main-intbody:parsing-main-intable-3}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intbody:insertion-mode-6}.

When the steps above require the UA to [clear the stack back to a table
body context]{#clear-the-stack-back-to-a-table-body-context .dfn}, it
means that the UA must, while the [current
node](#current-node){#parsing-main-intbody:current-node-3} is not a
[`tbody`{#parsing-main-intbody:the-tbody-element}](tables.html#the-tbody-element),
[`tfoot`{#parsing-main-intbody:the-tfoot-element}](tables.html#the-tfoot-element),
[`thead`{#parsing-main-intbody:the-thead-element}](tables.html#the-thead-element),
[`template`{#parsing-main-intbody:the-template-element}](scripting.html#the-template-element),
or
[`html`{#parsing-main-intbody:the-html-element}](semantics.html#the-html-element)
element, pop elements from the [stack of open
elements](#stack-of-open-elements){#parsing-main-intbody:stack-of-open-elements-5}.

The [current node](#current-node){#parsing-main-intbody:current-node-4}
being an
[`html`{#parsing-main-intbody:the-html-element-2}](semantics.html#the-html-element)
element after this process is a [fragment
case](#fragment-case){#parsing-main-intbody:fragment-case}.

###### [13.2.6.4.14]{.secno} The \"[in row]{.dfn}\" insertion mode[](#parsing-main-intr){.self-link} {#parsing-main-intr}

When the user agent is to apply the rules for the \"[in
row](#parsing-main-intr){#parsing-main-intr:parsing-main-intr}\"
[insertion mode](#insertion-mode){#parsing-main-intr:insertion-mode},
the user agent must handle the token as follows:

A start tag whose tag name is one of: \"th\", \"td\"

:   [Clear the stack back to a table row
    context](#clear-the-stack-back-to-a-table-row-context){#parsing-main-intr:clear-the-stack-back-to-a-table-row-context}.
    (See below.)

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-intr:insert-an-html-element}
    for the token, then switch the [insertion
    mode](#insertion-mode){#parsing-main-intr:insertion-mode-2} to \"[in
    cell](#parsing-main-intd){#parsing-main-intr:parsing-main-intd}\".

    Insert a
    [marker](#concept-parser-marker){#parsing-main-intr:concept-parser-marker}
    at the end of the [list of active formatting
    elements](#list-of-active-formatting-elements){#parsing-main-intr:list-of-active-formatting-elements}.

An end tag whose tag name is \"tr\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements}
    does not [have a `tr` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intr:has-an-element-in-table-scope},
    this is a [parse
    error](#parse-errors){#parsing-main-intr:parse-errors}; ignore the
    token.

    Otherwise:

    [Clear the stack back to a table row
    context](#clear-the-stack-back-to-a-table-row-context){#parsing-main-intr:clear-the-stack-back-to-a-table-row-context-2}.
    (See below.)

    Pop the [current
    node](#current-node){#parsing-main-intr:current-node} (which will be
    a
    [`tr`{#parsing-main-intr:the-tr-element}](tables.html#the-tr-element)
    element) from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements-2}.
    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intr:insertion-mode-3} to \"[in
    table
    body](#parsing-main-intbody){#parsing-main-intr:parsing-main-intbody}\".

A start tag whose tag name is one of: \"caption\", \"col\", \"colgroup\", \"tbody\", \"tfoot\", \"thead\", \"tr\"\
An end tag whose tag name is \"table\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements-3}
    does not [have a `tr` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intr:has-an-element-in-table-scope-2},
    this is a [parse
    error](#parse-errors){#parsing-main-intr:parse-errors-2}; ignore the
    token.

    Otherwise:

    [Clear the stack back to a table row
    context](#clear-the-stack-back-to-a-table-row-context){#parsing-main-intr:clear-the-stack-back-to-a-table-row-context-3}.
    (See below.)

    Pop the [current
    node](#current-node){#parsing-main-intr:current-node-2} (which will
    be a
    [`tr`{#parsing-main-intr:the-tr-element-2}](tables.html#the-tr-element)
    element) from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements-4}.
    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intr:insertion-mode-4} to \"[in
    table
    body](#parsing-main-intbody){#parsing-main-intr:parsing-main-intbody-2}\".

    Reprocess the token.

An end tag whose tag name is one of: \"tbody\", \"tfoot\", \"thead\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements-5}
    does not [have an element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intr:has-an-element-in-table-scope-3}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-intr:html-elements}
    with the same tag name as the token, this is a [parse
    error](#parse-errors){#parsing-main-intr:parse-errors-3}; ignore the
    token.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements-6}
    does not [have a `tr` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intr:has-an-element-in-table-scope-4},
    ignore the token.

    Otherwise:

    [Clear the stack back to a table row
    context](#clear-the-stack-back-to-a-table-row-context){#parsing-main-intr:clear-the-stack-back-to-a-table-row-context-4}.
    (See below.)

    Pop the [current
    node](#current-node){#parsing-main-intr:current-node-3} (which will
    be a
    [`tr`{#parsing-main-intr:the-tr-element-3}](tables.html#the-tr-element)
    element) from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements-7}.
    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intr:insertion-mode-5} to \"[in
    table
    body](#parsing-main-intbody){#parsing-main-intr:parsing-main-intbody-3}\".

    Reprocess the token.

An end tag whose tag name is one of: \"body\", \"caption\", \"col\", \"colgroup\", \"html\", \"td\", \"th\"

:   [Parse error](#parse-errors){#parsing-main-intr:parse-errors-4}.
    Ignore the token.

Anything else

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intr:using-the-rules-for}
    the \"[in
    table](#parsing-main-intable){#parsing-main-intr:parsing-main-intable}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intr:insertion-mode-6}.

When the steps above require the UA to [clear the stack back to a table
row context]{#clear-the-stack-back-to-a-table-row-context .dfn}, it
means that the UA must, while the [current
node](#current-node){#parsing-main-intr:current-node-4} is not a
[`tr`{#parsing-main-intr:the-tr-element-4}](tables.html#the-tr-element),
[`template`{#parsing-main-intr:the-template-element}](scripting.html#the-template-element),
or
[`html`{#parsing-main-intr:the-html-element}](semantics.html#the-html-element)
element, pop elements from the [stack of open
elements](#stack-of-open-elements){#parsing-main-intr:stack-of-open-elements-8}.

The [current node](#current-node){#parsing-main-intr:current-node-5}
being an
[`html`{#parsing-main-intr:the-html-element-2}](semantics.html#the-html-element)
element after this process is a [fragment
case](#fragment-case){#parsing-main-intr:fragment-case}.

###### [13.2.6.4.15]{.secno} The \"[in cell]{.dfn}\" insertion mode[](#parsing-main-intd){.self-link} {#parsing-main-intd}

When the user agent is to apply the rules for the \"[in
cell](#parsing-main-intd){#parsing-main-intd:parsing-main-intd}\"
[insertion mode](#insertion-mode){#parsing-main-intd:insertion-mode},
the user agent must handle the token as follows:

An end tag whose tag name is one of: \"td\", \"th\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intd:stack-of-open-elements}
    does not [have an element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intd:has-an-element-in-table-scope}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-intd:html-elements}
    with the same tag name as that of the token, then this is a [parse
    error](#parse-errors){#parsing-main-intd:parse-errors}; ignore the
    token.

    Otherwise:

    [Generate implied end
    tags](#generate-implied-end-tags){#parsing-main-intd:generate-implied-end-tags}.

    Now, if the [current
    node](#current-node){#parsing-main-intd:current-node} is not an
    [HTML
    element](infrastructure.html#html-elements){#parsing-main-intd:html-elements-2}
    with the same tag name as the token, then this is a [parse
    error](#parse-errors){#parsing-main-intd:parse-errors-2}.

    Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intd:stack-of-open-elements-2}
    until an [HTML
    element](infrastructure.html#html-elements){#parsing-main-intd:html-elements-3}
    with the same tag name as the token has been popped from the stack.

    [Clear the list of active formatting elements up to the last
    marker](#clear-the-list-of-active-formatting-elements-up-to-the-last-marker){#parsing-main-intd:clear-the-list-of-active-formatting-elements-up-to-the-last-marker}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intd:insertion-mode-2} to \"[in
    row](#parsing-main-intr){#parsing-main-intd:parsing-main-intr}\".

A start tag whose tag name is one of: \"caption\", \"col\", \"colgroup\", \"tbody\", \"td\", \"tfoot\", \"th\", \"thead\", \"tr\"

:   [Assert](https://infra.spec.whatwg.org/#assert){#parsing-main-intd:assert
    x-internal="assert"}: The [stack of open
    elements](#stack-of-open-elements){#parsing-main-intd:stack-of-open-elements-3}
    [has a `td` or `th` element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intd:has-an-element-in-table-scope-2}.

    [Close the cell](#close-the-cell){#parsing-main-intd:close-the-cell}
    (see below) and reprocess the token.

An end tag whose tag name is one of: \"body\", \"caption\", \"col\", \"colgroup\", \"html\"

:   [Parse error](#parse-errors){#parsing-main-intd:parse-errors-3}.
    Ignore the token.

An end tag whose tag name is one of: \"table\", \"tbody\", \"tfoot\", \"thead\", \"tr\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intd:stack-of-open-elements-4}
    does not [have an element in table
    scope](#has-an-element-in-table-scope){#parsing-main-intd:has-an-element-in-table-scope-3}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-intd:html-elements-4}
    with the same tag name as that of the token, then this is a [parse
    error](#parse-errors){#parsing-main-intd:parse-errors-4}; ignore the
    token.

    Otherwise, [close the
    cell](#close-the-cell){#parsing-main-intd:close-the-cell-2} (see
    below) and reprocess the token.

Anything else

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intd:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-intd:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intd:insertion-mode-3}.

Where the steps above say to [close the cell]{#close-the-cell .dfn},
they mean to run the following algorithm:

1.  [Generate implied end
    tags](#generate-implied-end-tags){#parsing-main-intd:generate-implied-end-tags-2}.

2.  If the [current
    node](#current-node){#parsing-main-intd:current-node-2} is not now a
    [`td`{#parsing-main-intd:the-td-element}](tables.html#the-td-element)
    element or a
    [`th`{#parsing-main-intd:the-th-element}](tables.html#the-th-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-intd:parse-errors-5}.

3.  Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intd:stack-of-open-elements-5}
    until a
    [`td`{#parsing-main-intd:the-td-element-2}](tables.html#the-td-element)
    element or a
    [`th`{#parsing-main-intd:the-th-element-2}](tables.html#the-th-element)
    element has been popped from the stack.

4.  [Clear the list of active formatting elements up to the last
    marker](#clear-the-list-of-active-formatting-elements-up-to-the-last-marker){#parsing-main-intd:clear-the-list-of-active-formatting-elements-up-to-the-last-marker-2}.

5.  Switch the [insertion
    mode](#insertion-mode){#parsing-main-intd:insertion-mode-4} to \"[in
    row](#parsing-main-intr){#parsing-main-intd:parsing-main-intr-2}\".

The [stack of open
elements](#stack-of-open-elements){#parsing-main-intd:stack-of-open-elements-6}
cannot have both a
[`td`{#parsing-main-intd:the-td-element-3}](tables.html#the-td-element)
and a
[`th`{#parsing-main-intd:the-th-element-3}](tables.html#the-th-element)
element [in table
scope](#has-an-element-in-table-scope){#parsing-main-intd:has-an-element-in-table-scope-4}
at the same time, nor can it have neither when the [close the
cell](#close-the-cell){#parsing-main-intd:close-the-cell-3} algorithm is
invoked.

###### [13.2.6.4.16]{.secno} The \"[in select]{.dfn}\" insertion mode[](#parsing-main-inselect){.self-link} {#parsing-main-inselect}

When the user agent is to apply the rules for the \"[in
select](#parsing-main-inselect){#parsing-main-inselect:parsing-main-inselect}\"
[insertion
mode](#insertion-mode){#parsing-main-inselect:insertion-mode}, the user
agent must handle the token as follows:

A character token that is U+0000 NULL

:   [Parse error](#parse-errors){#parsing-main-inselect:parse-errors}.
    Ignore the token.

Any other character token

:   [Insert the token\'s
    character](#insert-a-character){#parsing-main-inselect:insert-a-character}.

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-inselect:insert-a-comment}.

A DOCTYPE token

:   [Parse error](#parse-errors){#parsing-main-inselect:parse-errors-2}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inselect:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-inselect:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inselect:insertion-mode-2}.

A start tag whose tag name is \"option\"

:   If the [current
    node](#current-node){#parsing-main-inselect:current-node} is an
    [`option`{#parsing-main-inselect:the-option-element}](form-elements.html#the-option-element)
    element, pop that node from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inselect:insert-an-html-element}
    for the token.

A start tag whose tag name is \"optgroup\"

:   If the [current
    node](#current-node){#parsing-main-inselect:current-node-2} is an
    [`option`{#parsing-main-inselect:the-option-element-2}](form-elements.html#the-option-element)
    element, pop that node from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-2}.

    If the [current
    node](#current-node){#parsing-main-inselect:current-node-3} is an
    [`optgroup`{#parsing-main-inselect:the-optgroup-element}](form-elements.html#the-optgroup-element)
    element, pop that node from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-3}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inselect:insert-an-html-element-2}
    for the token.

A start tag whose tag name is \"hr\"

:   If the [current
    node](#current-node){#parsing-main-inselect:current-node-4} is an
    [`option`{#parsing-main-inselect:the-option-element-3}](form-elements.html#the-option-element)
    element, pop that node from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-4}.

    If the [current
    node](#current-node){#parsing-main-inselect:current-node-5} is an
    [`optgroup`{#parsing-main-inselect:the-optgroup-element-2}](form-elements.html#the-optgroup-element)
    element, pop that node from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-5}.

    [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inselect:insert-an-html-element-3}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inselect:current-node-6} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-6}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inselect:acknowledge-self-closing-flag},
    if it is set.

An end tag whose tag name is \"optgroup\"

:   First, if the [current
    node](#current-node){#parsing-main-inselect:current-node-7} is an
    [`option`{#parsing-main-inselect:the-option-element-4}](form-elements.html#the-option-element)
    element, and the node immediately before it in the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-7}
    is an
    [`optgroup`{#parsing-main-inselect:the-optgroup-element-3}](form-elements.html#the-optgroup-element)
    element, then pop the [current
    node](#current-node){#parsing-main-inselect:current-node-8} from the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-8}.

    If the [current
    node](#current-node){#parsing-main-inselect:current-node-9} is an
    [`optgroup`{#parsing-main-inselect:the-optgroup-element-4}](form-elements.html#the-optgroup-element)
    element, then pop that node from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-9}.
    Otherwise, this is a [parse
    error](#parse-errors){#parsing-main-inselect:parse-errors-3}; ignore
    the token.

An end tag whose tag name is \"option\"

:   If the [current
    node](#current-node){#parsing-main-inselect:current-node-10} is an
    [`option`{#parsing-main-inselect:the-option-element-5}](form-elements.html#the-option-element)
    element, then pop that node from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-10}.
    Otherwise, this is a [parse
    error](#parse-errors){#parsing-main-inselect:parse-errors-4}; ignore
    the token.

An end tag whose tag name is \"select\"

:   If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-11}
    does not [have a `select` element in select
    scope](#has-an-element-in-select-scope){#parsing-main-inselect:has-an-element-in-select-scope},
    this is a [parse
    error](#parse-errors){#parsing-main-inselect:parse-errors-5}; ignore
    the token. ([fragment
    case](#fragment-case){#parsing-main-inselect:fragment-case})

    Otherwise:

    Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-12}
    until a
    [`select`{#parsing-main-inselect:the-select-element}](form-elements.html#the-select-element)
    element has been popped from the stack.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-inselect:reset-the-insertion-mode-appropriately}.

A start tag whose tag name is \"select\"

:   [Parse error](#parse-errors){#parsing-main-inselect:parse-errors-6}.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-13}
    does not [have a `select` element in select
    scope](#has-an-element-in-select-scope){#parsing-main-inselect:has-an-element-in-select-scope-2},
    ignore the token. ([fragment
    case](#fragment-case){#parsing-main-inselect:fragment-case-2})

    Otherwise:

    Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-14}
    until a
    [`select`{#parsing-main-inselect:the-select-element-2}](form-elements.html#the-select-element)
    element has been popped from the stack.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-inselect:reset-the-insertion-mode-appropriately-2}.

    It just gets treated like an end tag.

A start tag whose tag name is one of: \"input\", \"keygen\", \"textarea\"

:   [Parse error](#parse-errors){#parsing-main-inselect:parse-errors-7}.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-15}
    does not [have a `select` element in select
    scope](#has-an-element-in-select-scope){#parsing-main-inselect:has-an-element-in-select-scope-3},
    ignore the token. ([fragment
    case](#fragment-case){#parsing-main-inselect:fragment-case-3})

    Otherwise:

    Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselect:stack-of-open-elements-16}
    until a
    [`select`{#parsing-main-inselect:the-select-element-3}](form-elements.html#the-select-element)
    element has been popped from the stack.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-inselect:reset-the-insertion-mode-appropriately-3}.

    Reprocess the token.

A start tag whose tag name is one of: \"script\", \"template\"\
An end tag whose tag name is \"template\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inselect:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-inselect:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inselect:insertion-mode-3}.

An end-of-file token

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inselect:using-the-rules-for-3}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-inselect:parsing-main-inbody-2}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inselect:insertion-mode-4}.

Anything else

:   [Parse error](#parse-errors){#parsing-main-inselect:parse-errors-8}.
    Ignore the token.

###### [13.2.6.4.17]{.secno} The \"[in select in table]{.dfn}\" insertion mode[](#parsing-main-inselectintable){.self-link} {#parsing-main-inselectintable}

When the user agent is to apply the rules for the \"[in select in
table](#parsing-main-inselectintable){#parsing-main-inselectintable:parsing-main-inselectintable}\"
[insertion
mode](#insertion-mode){#parsing-main-inselectintable:insertion-mode},
the user agent must handle the token as follows:

A start tag whose tag name is one of: \"caption\", \"table\", \"tbody\", \"tfoot\", \"thead\", \"tr\", \"td\", \"th\"

:   [Parse
    error](#parse-errors){#parsing-main-inselectintable:parse-errors}.

    Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselectintable:stack-of-open-elements}
    until a
    [`select`{#parsing-main-inselectintable:the-select-element}](form-elements.html#the-select-element)
    element has been popped from the stack.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-inselectintable:reset-the-insertion-mode-appropriately}.

    Reprocess the token.

An end tag whose tag name is one of: \"caption\", \"table\", \"tbody\", \"tfoot\", \"thead\", \"tr\", \"td\", \"th\"

:   [Parse
    error](#parse-errors){#parsing-main-inselectintable:parse-errors-2}.

    If the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselectintable:stack-of-open-elements-2}
    does not [have an element in table
    scope](#has-an-element-in-table-scope){#parsing-main-inselectintable:has-an-element-in-table-scope}
    that is an [HTML
    element](infrastructure.html#html-elements){#parsing-main-inselectintable:html-elements}
    with the same tag name as that of the token, then ignore the token.

    Otherwise:

    Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inselectintable:stack-of-open-elements-3}
    until a
    [`select`{#parsing-main-inselectintable:the-select-element-2}](form-elements.html#the-select-element)
    element has been popped from the stack.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-inselectintable:reset-the-insertion-mode-appropriately-2}.

    Reprocess the token.

Anything else

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inselectintable:using-the-rules-for}
    the \"[in
    select](#parsing-main-inselect){#parsing-main-inselectintable:parsing-main-inselect}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inselectintable:insertion-mode-2}.

###### [13.2.6.4.18]{.secno} The \"[in template]{.dfn}\" insertion mode[](#parsing-main-intemplate){.self-link} {#parsing-main-intemplate}

When the user agent is to apply the rules for the \"[in
template](#parsing-main-intemplate){#parsing-main-intemplate:parsing-main-intemplate}\"
[insertion
mode](#insertion-mode){#parsing-main-intemplate:insertion-mode}, the
user agent must handle the token as follows:

A character token\
A comment token\
A DOCTYPE token

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intemplate:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-intemplate:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intemplate:insertion-mode-2}.

A start tag whose tag name is one of: \"base\", \"basefont\", \"bgsound\", \"link\", \"meta\", \"noframes\", \"script\", \"style\", \"template\", \"title\"\
An end tag whose tag name is \"template\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-intemplate:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-intemplate:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#parsing-main-intemplate:insertion-mode-3}.

A start tag whose tag name is one of: \"caption\", \"colgroup\", \"tbody\", \"tfoot\", \"thead\"

:   Pop the [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode}
    off the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes}.

    Push \"[in
    table](#parsing-main-intable){#parsing-main-intemplate:parsing-main-intable}\"
    onto the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-2}
    so that it is the new [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-2}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intemplate:insertion-mode-4} to
    \"[in
    table](#parsing-main-intable){#parsing-main-intemplate:parsing-main-intable-2}\",
    and reprocess the token.

A start tag whose tag name is \"col\"

:   Pop the [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-3}
    off the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-3}.

    Push \"[in column
    group](#parsing-main-incolgroup){#parsing-main-intemplate:parsing-main-incolgroup}\"
    onto the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-4}
    so that it is the new [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-4}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intemplate:insertion-mode-5} to
    \"[in column
    group](#parsing-main-incolgroup){#parsing-main-intemplate:parsing-main-incolgroup-2}\",
    and reprocess the token.

A start tag whose tag name is \"tr\"

:   Pop the [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-5}
    off the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-5}.

    Push \"[in table
    body](#parsing-main-intbody){#parsing-main-intemplate:parsing-main-intbody}\"
    onto the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-6}
    so that it is the new [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-6}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intemplate:insertion-mode-6} to
    \"[in table
    body](#parsing-main-intbody){#parsing-main-intemplate:parsing-main-intbody-2}\",
    and reprocess the token.

A start tag whose tag name is one of: \"td\", \"th\"

:   Pop the [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-7}
    off the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-7}.

    Push \"[in
    row](#parsing-main-intr){#parsing-main-intemplate:parsing-main-intr}\"
    onto the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-8}
    so that it is the new [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-8}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intemplate:insertion-mode-7} to
    \"[in
    row](#parsing-main-intr){#parsing-main-intemplate:parsing-main-intr-2}\",
    and reprocess the token.

Any other start tag

:   Pop the [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-9}
    off the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-9}.

    Push \"[in
    body](#parsing-main-inbody){#parsing-main-intemplate:parsing-main-inbody-2}\"
    onto the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-10}
    so that it is the new [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-10}.

    Switch the [insertion
    mode](#insertion-mode){#parsing-main-intemplate:insertion-mode-8} to
    \"[in
    body](#parsing-main-inbody){#parsing-main-intemplate:parsing-main-inbody-3}\",
    and reprocess the token.

Any other end tag

:   [Parse error](#parse-errors){#parsing-main-intemplate:parse-errors}.
    Ignore the token.

An end-of-file token

:   If there is no
    [`template`{#parsing-main-intemplate:the-template-element}](scripting.html#the-template-element)
    element on the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intemplate:stack-of-open-elements},
    then [stop
    parsing](#stop-parsing){#parsing-main-intemplate:stop-parsing}.
    ([fragment
    case](#fragment-case){#parsing-main-intemplate:fragment-case})

    Otherwise, this is a [parse
    error](#parse-errors){#parsing-main-intemplate:parse-errors-2}.

    Pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-intemplate:stack-of-open-elements-2}
    until a
    [`template`{#parsing-main-intemplate:the-template-element-2}](scripting.html#the-template-element)
    element has been popped from the stack.

    [Clear the list of active formatting elements up to the last
    marker](#clear-the-list-of-active-formatting-elements-up-to-the-last-marker){#parsing-main-intemplate:clear-the-list-of-active-formatting-elements-up-to-the-last-marker}.

    Pop the [current template insertion
    mode](#current-template-insertion-mode){#parsing-main-intemplate:current-template-insertion-mode-11}
    off the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-main-intemplate:stack-of-template-insertion-modes-11}.

    [Reset the insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-main-intemplate:reset-the-insertion-mode-appropriately}.

    Reprocess the token.

###### [13.2.6.4.19]{.secno} The \"[after body]{.dfn}\" insertion mode[](#parsing-main-afterbody){.self-link} {#parsing-main-afterbody}

When the user agent is to apply the rules for the \"[after
body](#parsing-main-afterbody){#parsing-main-afterbody:parsing-main-afterbody}\"
[insertion
mode](#insertion-mode){#parsing-main-afterbody:insertion-mode}, the user
agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-afterbody:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-afterbody:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-afterbody:insertion-mode-2}.

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-afterbody:insert-a-comment}
    as the last child of the first element in the [stack of open
    elements](#stack-of-open-elements){#parsing-main-afterbody:stack-of-open-elements}
    (the
    [`html`{#parsing-main-afterbody:the-html-element}](semantics.html#the-html-element)
    element).

A DOCTYPE token

:   [Parse error](#parse-errors){#parsing-main-afterbody:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-afterbody:using-the-rules-for-2}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-afterbody:parsing-main-inbody-2}\"
    [insertion
    mode](#insertion-mode){#parsing-main-afterbody:insertion-mode-3}.

An end tag whose tag name is \"html\"

:   If the parser was created as part of the [HTML fragment parsing
    algorithm](#html-fragment-parsing-algorithm){#parsing-main-afterbody:html-fragment-parsing-algorithm},
    this is a [parse
    error](#parse-errors){#parsing-main-afterbody:parse-errors-2};
    ignore the token. ([fragment
    case](#fragment-case){#parsing-main-afterbody:fragment-case})

    Otherwise, switch the [insertion
    mode](#insertion-mode){#parsing-main-afterbody:insertion-mode-4} to
    \"[after after
    body](#the-after-after-body-insertion-mode){#parsing-main-afterbody:the-after-after-body-insertion-mode}\".

An end-of-file token

:   [Stop parsing](#stop-parsing){#parsing-main-afterbody:stop-parsing}.

Anything else

:   [Parse
    error](#parse-errors){#parsing-main-afterbody:parse-errors-3}.
    Switch the [insertion
    mode](#insertion-mode){#parsing-main-afterbody:insertion-mode-5} to
    \"[in
    body](#parsing-main-inbody){#parsing-main-afterbody:parsing-main-inbody-3}\"
    and reprocess the token.

###### [13.2.6.4.20]{.secno} The \"[in frameset]{.dfn}\" insertion mode[](#parsing-main-inframeset){.self-link} {#parsing-main-inframeset}

When the user agent is to apply the rules for the \"[in
frameset](#parsing-main-inframeset){#parsing-main-inframeset:parsing-main-inframeset}\"
[insertion
mode](#insertion-mode){#parsing-main-inframeset:insertion-mode}, the
user agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   [Insert the
    character](#insert-a-character){#parsing-main-inframeset:insert-a-character}.

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-inframeset:insert-a-comment}.

A DOCTYPE token

:   [Parse error](#parse-errors){#parsing-main-inframeset:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inframeset:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-inframeset:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inframeset:insertion-mode-2}.

A start tag whose tag name is \"frameset\"

:   [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inframeset:insert-an-html-element}
    for the token.

An end tag whose tag name is \"frameset\"

:   If the [current
    node](#current-node){#parsing-main-inframeset:current-node} is the
    root
    [`html`{#parsing-main-inframeset:the-html-element}](semantics.html#the-html-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-inframeset:parse-errors-2};
    ignore the token. ([fragment
    case](#fragment-case){#parsing-main-inframeset:fragment-case})

    Otherwise, pop the [current
    node](#current-node){#parsing-main-inframeset:current-node-2} from
    the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inframeset:stack-of-open-elements}.

    If the parser was not created as part of the [HTML fragment parsing
    algorithm](#html-fragment-parsing-algorithm){#parsing-main-inframeset:html-fragment-parsing-algorithm}
    ([fragment
    case](#fragment-case){#parsing-main-inframeset:fragment-case-2}),
    and the [current
    node](#current-node){#parsing-main-inframeset:current-node-3} is no
    longer a
    [`frameset`{#parsing-main-inframeset:frameset}](obsolete.html#frameset)
    element, then switch the [insertion
    mode](#insertion-mode){#parsing-main-inframeset:insertion-mode-3} to
    \"[after
    frameset](#parsing-main-afterframeset){#parsing-main-inframeset:parsing-main-afterframeset}\".

A start tag whose tag name is \"frame\"

:   [Insert an HTML
    element](#insert-an-html-element){#parsing-main-inframeset:insert-an-html-element-2}
    for the token. Immediately pop the [current
    node](#current-node){#parsing-main-inframeset:current-node-4} off
    the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inframeset:stack-of-open-elements-2}.

    [Acknowledge the token\'s *self-closing
    flag*](#acknowledge-self-closing-flag){#parsing-main-inframeset:acknowledge-self-closing-flag},
    if it is set.

A start tag whose tag name is \"noframes\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-inframeset:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-inframeset:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#parsing-main-inframeset:insertion-mode-4}.

An end-of-file token

:   If the [current
    node](#current-node){#parsing-main-inframeset:current-node-5} is not
    the root
    [`html`{#parsing-main-inframeset:the-html-element-2}](semantics.html#the-html-element)
    element, then this is a [parse
    error](#parse-errors){#parsing-main-inframeset:parse-errors-3}.

    The [current
    node](#current-node){#parsing-main-inframeset:current-node-6} can
    only be the root
    [`html`{#parsing-main-inframeset:the-html-element-3}](semantics.html#the-html-element)
    element in the [fragment
    case](#fragment-case){#parsing-main-inframeset:fragment-case-3}.

    [Stop
    parsing](#stop-parsing){#parsing-main-inframeset:stop-parsing}.

Anything else

:   [Parse
    error](#parse-errors){#parsing-main-inframeset:parse-errors-4}.
    Ignore the token.

###### [13.2.6.4.21]{.secno} The \"[after frameset]{.dfn}\" insertion mode[](#parsing-main-afterframeset){.self-link} {#parsing-main-afterframeset}

When the user agent is to apply the rules for the \"[after
frameset](#parsing-main-afterframeset){#parsing-main-afterframeset:parsing-main-afterframeset}\"
[insertion
mode](#insertion-mode){#parsing-main-afterframeset:insertion-mode}, the
user agent must handle the token as follows:

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   [Insert the
    character](#insert-a-character){#parsing-main-afterframeset:insert-a-character}.

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-afterframeset:insert-a-comment}.

A DOCTYPE token

:   [Parse
    error](#parse-errors){#parsing-main-afterframeset:parse-errors}.
    Ignore the token.

A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-afterframeset:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#parsing-main-afterframeset:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#parsing-main-afterframeset:insertion-mode-2}.

An end tag whose tag name is \"html\"

:   Switch the [insertion
    mode](#insertion-mode){#parsing-main-afterframeset:insertion-mode-3}
    to \"[after after
    frameset](#the-after-after-frameset-insertion-mode){#parsing-main-afterframeset:the-after-after-frameset-insertion-mode}\".

A start tag whose tag name is \"noframes\"

:   Process the token [using the rules
    for](#using-the-rules-for){#parsing-main-afterframeset:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#parsing-main-afterframeset:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#parsing-main-afterframeset:insertion-mode-4}.

An end-of-file token

:   [Stop
    parsing](#stop-parsing){#parsing-main-afterframeset:stop-parsing}.

Anything else

:   [Parse
    error](#parse-errors){#parsing-main-afterframeset:parse-errors-2}.
    Ignore the token.

###### [13.2.6.4.22]{.secno} The \"[after after body]{.dfn}\" insertion mode[](#the-after-after-body-insertion-mode){.self-link}

When the user agent is to apply the rules for the \"[after after
body](#the-after-after-body-insertion-mode){#the-after-after-body-insertion-mode:the-after-after-body-insertion-mode}\"
[insertion
mode](#insertion-mode){#the-after-after-body-insertion-mode:insertion-mode},
the user agent must handle the token as follows:

A comment token

:   [Insert a
    comment](#insert-a-comment){#the-after-after-body-insertion-mode:insert-a-comment}
    as the last child of the
    [`Document`{#the-after-after-body-insertion-mode:document}](dom.html#document)
    object.

A DOCTYPE token\
A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE\
A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#the-after-after-body-insertion-mode:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#the-after-after-body-insertion-mode:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#the-after-after-body-insertion-mode:insertion-mode-2}.

An end-of-file token

:   [Stop
    parsing](#stop-parsing){#the-after-after-body-insertion-mode:stop-parsing}.

Anything else

:   [Parse
    error](#parse-errors){#the-after-after-body-insertion-mode:parse-errors}.
    Switch the [insertion
    mode](#insertion-mode){#the-after-after-body-insertion-mode:insertion-mode-3}
    to \"[in
    body](#parsing-main-inbody){#the-after-after-body-insertion-mode:parsing-main-inbody-2}\"
    and reprocess the token.

###### [13.2.6.4.23]{.secno} The \"[after after frameset]{.dfn}\" insertion mode[](#the-after-after-frameset-insertion-mode){.self-link}

When the user agent is to apply the rules for the \"[after after
frameset](#the-after-after-frameset-insertion-mode){#the-after-after-frameset-insertion-mode:the-after-after-frameset-insertion-mode}\"
[insertion
mode](#insertion-mode){#the-after-after-frameset-insertion-mode:insertion-mode},
the user agent must handle the token as follows:

A comment token

:   [Insert a
    comment](#insert-a-comment){#the-after-after-frameset-insertion-mode:insert-a-comment}
    as the last child of the
    [`Document`{#the-after-after-frameset-insertion-mode:document}](dom.html#document)
    object.

A DOCTYPE token\
A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE\
A start tag whose tag name is \"html\"

:   Process the token [using the rules
    for](#using-the-rules-for){#the-after-after-frameset-insertion-mode:using-the-rules-for}
    the \"[in
    body](#parsing-main-inbody){#the-after-after-frameset-insertion-mode:parsing-main-inbody}\"
    [insertion
    mode](#insertion-mode){#the-after-after-frameset-insertion-mode:insertion-mode-2}.

An end-of-file token

:   [Stop
    parsing](#stop-parsing){#the-after-after-frameset-insertion-mode:stop-parsing}.

A start tag whose tag name is \"noframes\"

:   Process the token [using the rules
    for](#using-the-rules-for){#the-after-after-frameset-insertion-mode:using-the-rules-for-2}
    the \"[in
    head](#parsing-main-inhead){#the-after-after-frameset-insertion-mode:parsing-main-inhead}\"
    [insertion
    mode](#insertion-mode){#the-after-after-frameset-insertion-mode:insertion-mode-3}.

Anything else

:   [Parse
    error](#parse-errors){#the-after-after-frameset-insertion-mode:parse-errors}.
    Ignore the token.

##### [13.2.6.5]{.secno} The rules for parsing tokens [in foreign content]{.dfn}[](#parsing-main-inforeign){.self-link} {#parsing-main-inforeign}

When the user agent is to apply the rules for parsing tokens in foreign
content, the user agent must handle the token as follows:

A character token that is U+0000 NULL

:   [Parse error](#parse-errors){#parsing-main-inforeign:parse-errors}.
    [Insert a U+FFFD REPLACEMENT CHARACTER
    character](#insert-a-character){#parsing-main-inforeign:insert-a-character}.

A character token that is one of U+0009 CHARACTER TABULATION, U+000A LINE FEED (LF), U+000C FORM FEED (FF), U+000D CARRIAGE RETURN (CR), or U+0020 SPACE

:   [Insert the token\'s
    character](#insert-a-character){#parsing-main-inforeign:insert-a-character-2}.

Any other character token

:   [Insert the token\'s
    character](#insert-a-character){#parsing-main-inforeign:insert-a-character-3}.

    Set the [frameset-ok
    flag](#frameset-ok-flag){#parsing-main-inforeign:frameset-ok-flag}
    to \"not ok\".

A comment token

:   [Insert a
    comment](#insert-a-comment){#parsing-main-inforeign:insert-a-comment}.

A DOCTYPE token

:   [Parse
    error](#parse-errors){#parsing-main-inforeign:parse-errors-2}.
    Ignore the token.

A start tag whose tag name is one of: \"b\", \"big\", \"blockquote\", \"body\", \"br\", \"center\", \"code\", \"dd\", \"div\", \"dl\", \"dt\", \"em\", \"embed\", \"h1\", \"h2\", \"h3\", \"h4\", \"h5\", \"h6\", \"head\", \"hr\", \"i\", \"img\", \"li\", \"listing\", \"menu\", \"meta\", \"nobr\", \"ol\", \"p\", \"pre\", \"ruby\", \"s\", \"small\", \"span\", \"strong\", \"strike\", \"sub\", \"sup\", \"table\", \"tt\", \"u\", \"ul\", \"var\"\
A start tag whose tag name is \"font\", if the token has any attributes named \"color\", \"face\", or \"size\"\
An end tag whose tag name is \"br\", \"p\"

:   [Parse
    error](#parse-errors){#parsing-main-inforeign:parse-errors-3}.

    While the [current
    node](#current-node){#parsing-main-inforeign:current-node} is not a
    [MathML text integration
    point](#mathml-text-integration-point){#parsing-main-inforeign:mathml-text-integration-point},
    an [HTML integration
    point](#html-integration-point){#parsing-main-inforeign:html-integration-point},
    or an element in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing-main-inforeign:html-namespace-2
    x-internal="html-namespace-2"}, pop elements from the [stack of open
    elements](#stack-of-open-elements){#parsing-main-inforeign:stack-of-open-elements}.

    Reprocess the token according to the rules given in the section
    corresponding to the current [insertion
    mode](#insertion-mode){#parsing-main-inforeign:insertion-mode} in
    HTML content.

Any other start tag

:   If the [adjusted current
    node](#adjusted-current-node){#parsing-main-inforeign:adjusted-current-node}
    is an element in the [MathML
    namespace](https://infra.spec.whatwg.org/#mathml-namespace){#parsing-main-inforeign:mathml-namespace
    x-internal="mathml-namespace"}, [adjust MathML
    attributes](#adjust-mathml-attributes){#parsing-main-inforeign:adjust-mathml-attributes}
    for the token. (This fixes the case of MathML attributes that are
    not all lowercase.)

    If the [adjusted current
    node](#adjusted-current-node){#parsing-main-inforeign:adjusted-current-node-2}
    is an element in the [SVG
    namespace](https://infra.spec.whatwg.org/#svg-namespace){#parsing-main-inforeign:svg-namespace
    x-internal="svg-namespace"}, and the token\'s tag name is one of the
    ones in the first column of the following table, change the tag name
    to the name given in the corresponding cell in the second column.
    (This fixes the case of SVG elements that are not all lowercase.)

    Tag name

    Element name

    `altglyph`

    `altGlyph`

    `altglyphdef`

    `altGlyphDef`

    `altglyphitem`

    `altGlyphItem`

    `animatecolor`

    `animateColor`

    `animatemotion`

    `animateMotion`

    `animatetransform`

    `animateTransform`

    `clippath`

    `clipPath`

    `feblend`

    `feBlend`

    `fecolormatrix`

    `feColorMatrix`

    `fecomponenttransfer`

    `feComponentTransfer`

    `fecomposite`

    `feComposite`

    `feconvolvematrix`

    `feConvolveMatrix`

    `fediffuselighting`

    `feDiffuseLighting`

    `fedisplacementmap`

    `feDisplacementMap`

    `fedistantlight`

    `feDistantLight`

    `fedropshadow`

    `feDropShadow`

    `feflood`

    `feFlood`

    `fefunca`

    `feFuncA`

    `fefuncb`

    `feFuncB`

    `fefuncg`

    `feFuncG`

    `fefuncr`

    `feFuncR`

    `fegaussianblur`

    `feGaussianBlur`

    `feimage`

    `feImage`

    `femerge`

    `feMerge`

    `femergenode`

    `feMergeNode`

    `femorphology`

    `feMorphology`

    `feoffset`

    `feOffset`

    `fepointlight`

    `fePointLight`

    `fespecularlighting`

    `feSpecularLighting`

    `fespotlight`

    `feSpotLight`

    `fetile`

    `feTile`

    `feturbulence`

    `feTurbulence`

    `foreignobject`

    `foreignObject`

    `glyphref`

    `glyphRef`

    `lineargradient`

    `linearGradient`

    `radialgradient`

    `radialGradient`

    `textpath`

    `textPath`

    If the [adjusted current
    node](#adjusted-current-node){#parsing-main-inforeign:adjusted-current-node-3}
    is an element in the [SVG
    namespace](https://infra.spec.whatwg.org/#svg-namespace){#parsing-main-inforeign:svg-namespace-2
    x-internal="svg-namespace"}, [adjust SVG
    attributes](#adjust-svg-attributes){#parsing-main-inforeign:adjust-svg-attributes}
    for the token. (This fixes the case of SVG attributes that are not
    all lowercase.)

    [Adjust foreign
    attributes](#adjust-foreign-attributes){#parsing-main-inforeign:adjust-foreign-attributes}
    for the token. (This fixes the use of namespaced attributes, in
    particular XLink in SVG.)

    [Insert a foreign
    element](#insert-a-foreign-element){#parsing-main-inforeign:insert-a-foreign-element}
    for the token, with [adjusted current
    node](#adjusted-current-node){#parsing-main-inforeign:adjusted-current-node-4}\'s
    namespace and false.

    If the token has its *[self-closing flag](#self-closing-flag)* set,
    then run the appropriate steps from the following list:

    If the token\'s tag name is \"script\", and the new [current node](#current-node){#parsing-main-inforeign:current-node-2} is in the [SVG namespace](https://infra.spec.whatwg.org/#svg-namespace){#parsing-main-inforeign:svg-namespace-3 x-internal="svg-namespace"}

    :   [Acknowledge the token\'s *self-closing
        flag*](#acknowledge-self-closing-flag){#parsing-main-inforeign:acknowledge-self-closing-flag},
        and then act as described in the steps for a \"script\" end tag
        below.

    Otherwise

    :   Pop the [current
        node](#current-node){#parsing-main-inforeign:current-node-3} off
        the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inforeign:stack-of-open-elements-2}
        and [acknowledge the token\'s *self-closing
        flag*](#acknowledge-self-closing-flag){#parsing-main-inforeign:acknowledge-self-closing-flag-2}.

An end tag whose tag name is \"script\", if the [current node](#current-node){#parsing-main-inforeign:current-node-4} is an [SVG `script`](https://svgwg.org/svg2-draft/interact.html#ScriptElement){#parsing-main-inforeign:svg-script x-internal="svg-script"} element

:   Pop the [current
    node](#current-node){#parsing-main-inforeign:current-node-5} off the
    [stack of open
    elements](#stack-of-open-elements){#parsing-main-inforeign:stack-of-open-elements-3}.

    Let the `old insertion point`{.variable} have the same value as the
    current [insertion
    point](#insertion-point){#parsing-main-inforeign:insertion-point}.
    Let the [insertion
    point](#insertion-point){#parsing-main-inforeign:insertion-point-2}
    be just before the [next input
    character](#next-input-character){#parsing-main-inforeign:next-input-character}.

    Increment the parser\'s [script nesting
    level](#script-nesting-level){#parsing-main-inforeign:script-nesting-level}
    by one. Set the [parser pause
    flag](#parser-pause-flag){#parsing-main-inforeign:parser-pause-flag}
    to true.

    If the [active speculative HTML
    parser](#active-speculative-html-parser){#parsing-main-inforeign:active-speculative-html-parser}
    is null and the user agent supports SVG, then [Process the SVG
    `script`
    element](https://www.w3.org/TR/SVGMobile12/script.html#ScriptContentProcessing)
    according to the SVG rules. [\[SVG\]](references.html#refsSVG)

    Even if this causes [new characters to be inserted into the
    tokenizer](dynamic-markup-insertion.html#dom-document-write){#parsing-main-inforeign:dom-document-write},
    the parser will not be executed reentrantly, since the [parser pause
    flag](#parser-pause-flag){#parsing-main-inforeign:parser-pause-flag-2}
    is true.

    Decrement the parser\'s [script nesting
    level](#script-nesting-level){#parsing-main-inforeign:script-nesting-level-2}
    by one. If the parser\'s [script nesting
    level](#script-nesting-level){#parsing-main-inforeign:script-nesting-level-3}
    is zero, then set the [parser pause
    flag](#parser-pause-flag){#parsing-main-inforeign:parser-pause-flag-3}
    to false.

    Let the [insertion
    point](#insertion-point){#parsing-main-inforeign:insertion-point-3}
    have the value of the `old insertion point`{.variable}. (In other
    words, restore the [insertion
    point](#insertion-point){#parsing-main-inforeign:insertion-point-4}
    to its previous value. This value might be the \"undefined\" value.)

Any other end tag

:   Run these steps:

    1.  Initialize `node`{.variable} to be the [current
        node](#current-node){#parsing-main-inforeign:current-node-6}
        (the bottommost node of the stack).

    2.  If `node`{.variable}\'s tag name, [converted to ASCII
        lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#parsing-main-inforeign:converted-to-ascii-lowercase
        x-internal="converted-to-ascii-lowercase"}, is not the same as
        the tag name of the token, then this is a [parse
        error](#parse-errors){#parsing-main-inforeign:parse-errors-4}.

    3.  *Loop*: If `node`{.variable} is the topmost element in the
        [stack of open
        elements](#stack-of-open-elements){#parsing-main-inforeign:stack-of-open-elements-4},
        then return. ([fragment
        case](#fragment-case){#parsing-main-inforeign:fragment-case})

    4.  If `node`{.variable}\'s tag name, [converted to ASCII
        lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#parsing-main-inforeign:converted-to-ascii-lowercase-2
        x-internal="converted-to-ascii-lowercase"}, is the same as the
        tag name of the token, pop elements from the [stack of open
        elements](#stack-of-open-elements){#parsing-main-inforeign:stack-of-open-elements-5}
        until `node`{.variable} has been popped from the stack, and then
        return.

    5.  Set `node`{.variable} to the previous entry in the [stack of
        open
        elements](#stack-of-open-elements){#parsing-main-inforeign:stack-of-open-elements-6}.

    6.  If `node`{.variable} is not an element in the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing-main-inforeign:html-namespace-2-2
        x-internal="html-namespace-2"}, return to the step labeled
        *loop*.

    7.  Otherwise, process the token according to the rules given in the
        section corresponding to the current [insertion
        mode](#insertion-mode){#parsing-main-inforeign:insertion-mode-2}
        in HTML content.

#### [13.2.7]{.secno} The end[](#the-end){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/DOMContentLoaded_event](https://developer.mozilla.org/en-US/docs/Web/API/Document/DOMContentLoaded_event "The DOMContentLoaded event fires when the initial HTML document has been completely loaded and parsed, without waiting for stylesheets, images, and subframes to finish loading.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
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

Once the user agent [stops parsing]{#stop-parsing .dfn} the document,
the user agent must run the following steps:

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/load_event](https://developer.mozilla.org/en-US/docs/Web/API/Window/load_event "The load event is fired when the whole page has loaded, including all dependent resources such as stylesheets, scripts, iframes, and images. This is in contrast to DOMContentLoaded, which is fired as soon as the page DOM has been loaded, without waiting for resources to finish loading.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera4+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

1.  If the [active speculative HTML
    parser](#active-speculative-html-parser){#the-end:active-speculative-html-parser}
    is not null, then [stop the speculative HTML
    parser](#stop-the-speculative-html-parser){#the-end:stop-the-speculative-html-parser}
    and return.

2.  Set the [insertion
    point](#insertion-point){#the-end:insertion-point} to undefined.

3.  [Update the current document
    readiness](dom.html#update-the-current-document-readiness){#the-end:update-the-current-document-readiness}
    to \"`interactive`\".

4.  Pop *all* the nodes off the [stack of open
    elements](#stack-of-open-elements){#the-end:stack-of-open-elements}.

5.  While the [list of scripts that will execute when the document has
    finished
    parsing](scripting.html#list-of-scripts-that-will-execute-when-the-document-has-finished-parsing){#the-end:list-of-scripts-that-will-execute-when-the-document-has-finished-parsing}
    is not empty:

    1.  [Spin the event
        loop](webappapis.html#spin-the-event-loop){#the-end:spin-the-event-loop}
        until the first
        [`script`{#the-end:the-script-element}](scripting.html#the-script-element)
        in the [list of scripts that will execute when the document has
        finished
        parsing](scripting.html#list-of-scripts-that-will-execute-when-the-document-has-finished-parsing){#the-end:list-of-scripts-that-will-execute-when-the-document-has-finished-parsing-2}
        has its [ready to be
        parser-executed](scripting.html#ready-to-be-parser-executed){#the-end:ready-to-be-parser-executed}
        set to true *and* the parser\'s
        [`Document`{#the-end:document}](dom.html#document) [has no style
        sheet that is blocking
        scripts](semantics.html#has-no-style-sheet-that-is-blocking-scripts){#the-end:has-no-style-sheet-that-is-blocking-scripts}.

    2.  [Execute the script
        element](scripting.html#execute-the-script-element){#the-end:execute-the-script-element}
        given by the first
        [`script`{#the-end:the-script-element-2}](scripting.html#the-script-element)
        in the [list of scripts that will execute when the document has
        finished
        parsing](scripting.html#list-of-scripts-that-will-execute-when-the-document-has-finished-parsing){#the-end:list-of-scripts-that-will-execute-when-the-document-has-finished-parsing-3}.

    3.  Remove the first
        [`script`{#the-end:the-script-element-3}](scripting.html#the-script-element)
        element from the [list of scripts that will execute when the
        document has finished
        parsing](scripting.html#list-of-scripts-that-will-execute-when-the-document-has-finished-parsing){#the-end:list-of-scripts-that-will-execute-when-the-document-has-finished-parsing-4}
        (i.e. shift out the first entry in the list).

6.  [Queue a global
    task](webappapis.html#queue-a-global-task){#the-end:queue-a-global-task}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#the-end:dom-manipulation-task-source}
    given the [`Document`{#the-end:document-2}](dom.html#document)\'s
    [relevant global
    object](webappapis.html#concept-relevant-global){#the-end:concept-relevant-global}
    to run the following substeps:

    1.  Set the [`Document`{#the-end:document-3}](dom.html#document)\'s
        [load timing
        info](dom.html#load-timing-info){#the-end:load-timing-info}\'s
        [DOM content loaded event start
        time](dom.html#dom-content-loaded-event-start-time){#the-end:dom-content-loaded-event-start-time}
        to the [current high resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#the-end:current-high-resolution-time
        x-internal="current-high-resolution-time"} given the
        [`Document`{#the-end:document-4}](dom.html#document)\'s
        [relevant global
        object](webappapis.html#concept-relevant-global){#the-end:concept-relevant-global-2}.

    2.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-end:concept-event-fire
        x-internal="concept-event-fire"} named
        [`DOMContentLoaded`{#the-end:event-domcontentloaded}](indices.html#event-domcontentloaded)
        at the [`Document`{#the-end:document-5}](dom.html#document)
        object, with its
        [`bubbles`{#the-end:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
        attribute initialized to true.

    3.  Set the [`Document`{#the-end:document-6}](dom.html#document)\'s
        [load timing
        info](dom.html#load-timing-info){#the-end:load-timing-info-2}\'s
        [DOM content loaded event end
        time](dom.html#dom-content-loaded-event-end-time){#the-end:dom-content-loaded-event-end-time}
        to the [current high resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#the-end:current-high-resolution-time-2
        x-internal="current-high-resolution-time"} given the
        [`Document`{#the-end:document-7}](dom.html#document)\'s
        [relevant global
        object](webappapis.html#concept-relevant-global){#the-end:concept-relevant-global-3}.

    4.  Enable the [client message
        queue](https://w3c.github.io/ServiceWorker/#dfn-client-message-queue){#the-end:dfn-client-message-queue
        x-internal="dfn-client-message-queue"} of the
        [`ServiceWorkerContainer`{#the-end:serviceworkercontainer}](https://w3c.github.io/ServiceWorker/#serviceworkercontainer){x-internal="serviceworkercontainer"}
        object whose associated [service worker
        client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client){#the-end:serviceworkercontainer-service-worker-client
        x-internal="serviceworkercontainer-service-worker-client"} is
        the [`Document`{#the-end:document-8}](dom.html#document)
        object\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#the-end:relevant-settings-object}.

    5.  Invoke [WebDriver BiDi DOM content
        loaded](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-dom-content-loaded){#the-end:webdriver-bidi-dom-content-loaded
        x-internal="webdriver-bidi-dom-content-loaded"} with the
        [`Document`{#the-end:document-9}](dom.html#document)\'s
        [browsing
        context](document-sequences.html#concept-document-bc){#the-end:concept-document-bc},
        and a new [WebDriver BiDi navigation
        status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#the-end:webdriver-bidi-navigation-status
        x-internal="webdriver-bidi-navigation-status"} whose
        [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#the-end:navigation-status-id
        x-internal="navigation-status-id"} is the
        [`Document`{#the-end:document-10}](dom.html#document) object\'s
        [during-loading navigation ID for WebDriver
        BiDi](dom.html#concept-document-navigation-id){#the-end:concept-document-navigation-id},
        [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#the-end:navigation-status-status
        x-internal="navigation-status-status"} is
        \"[`pending`{#the-end:navigation-status-pending}](https://w3c.github.io/webdriver-bidi/#navigation-status-pending){x-internal="navigation-status-pending"}\",
        and
        [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#the-end:navigation-status-url
        x-internal="navigation-status-url"} is the
        [`Document`{#the-end:document-11}](dom.html#document) object\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-end:the-document's-address
        x-internal="the-document's-address"}.

7.  [Spin the event
    loop](webappapis.html#spin-the-event-loop){#the-end:spin-the-event-loop-2}
    until the [set of scripts that will execute as soon as
    possible](scripting.html#set-of-scripts-that-will-execute-as-soon-as-possible){#the-end:set-of-scripts-that-will-execute-as-soon-as-possible}
    and the [list of scripts that will execute in order as soon as
    possible](scripting.html#list-of-scripts-that-will-execute-in-order-as-soon-as-possible){#the-end:list-of-scripts-that-will-execute-in-order-as-soon-as-possible}
    are empty.

8.  [Spin the event
    loop](webappapis.html#spin-the-event-loop){#the-end:spin-the-event-loop-3}
    until there is nothing that [delays the load
    event]{#delay-the-load-event .dfn} in the
    [`Document`{#the-end:document-12}](dom.html#document).

9.  [Queue a global
    task](webappapis.html#queue-a-global-task){#the-end:queue-a-global-task-2}
    on the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#the-end:dom-manipulation-task-source-2}
    given the [`Document`{#the-end:document-13}](dom.html#document)\'s
    [relevant global
    object](webappapis.html#concept-relevant-global){#the-end:concept-relevant-global-4}
    to run the following steps:

    1.  [Update the current document
        readiness](dom.html#update-the-current-document-readiness){#the-end:update-the-current-document-readiness-2}
        to \"`complete`\".

    2.  If the [`Document`{#the-end:document-14}](dom.html#document)
        object\'s [browsing
        context](document-sequences.html#concept-document-bc){#the-end:concept-document-bc-2}
        is null, then abort these steps.

    3.  Let `window`{.variable} be the
        [`Document`{#the-end:document-15}](dom.html#document)\'s
        [relevant global
        object](webappapis.html#concept-relevant-global){#the-end:concept-relevant-global-5}.

    4.  Set the [`Document`{#the-end:document-16}](dom.html#document)\'s
        [load timing
        info](dom.html#load-timing-info){#the-end:load-timing-info-3}\'s
        [load event start
        time](dom.html#load-event-start-time){#the-end:load-event-start-time}
        to the [current high resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#the-end:current-high-resolution-time-3
        x-internal="current-high-resolution-time"} given
        `window`{.variable}.

    5.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-end:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`load`{#the-end:event-load}](indices.html#event-load) at
        `window`{.variable}, with *legacy target override flag* set.

    6.  Invoke [WebDriver BiDi load
        complete](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-load-complete){#the-end:webdriver-bidi-load-complete
        x-internal="webdriver-bidi-load-complete"} with the
        [`Document`{#the-end:document-17}](dom.html#document)\'s
        [browsing
        context](document-sequences.html#concept-document-bc){#the-end:concept-document-bc-3},
        and a new [WebDriver BiDi navigation
        status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status){#the-end:webdriver-bidi-navigation-status-2
        x-internal="webdriver-bidi-navigation-status"} whose
        [id](https://w3c.github.io/webdriver-bidi/#navigation-status-id){#the-end:navigation-status-id-2
        x-internal="navigation-status-id"} is the
        [`Document`{#the-end:document-18}](dom.html#document) object\'s
        [during-loading navigation ID for WebDriver
        BiDi](dom.html#concept-document-navigation-id){#the-end:concept-document-navigation-id-2},
        [status](https://w3c.github.io/webdriver-bidi/#navigation-status-status){#the-end:navigation-status-status-2
        x-internal="navigation-status-status"} is
        \"[`complete`{#the-end:navigation-status-complete}](https://w3c.github.io/webdriver-bidi/#navigation-status-complete){x-internal="navigation-status-complete"}\",
        and
        [url](https://w3c.github.io/webdriver-bidi/#navigation-status-url){#the-end:navigation-status-url-2
        x-internal="navigation-status-url"} is the
        [`Document`{#the-end:document-19}](dom.html#document) object\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-end:the-document's-address-2
        x-internal="the-document's-address"}.

    7.  Set the [`Document`{#the-end:document-20}](dom.html#document)
        object\'s [during-loading navigation ID for WebDriver
        BiDi](dom.html#concept-document-navigation-id){#the-end:concept-document-navigation-id-3}
        to null.

    8.  Set the [`Document`{#the-end:document-21}](dom.html#document)\'s
        [load timing
        info](dom.html#load-timing-info){#the-end:load-timing-info-4}\'s
        [load event end
        time](dom.html#load-event-end-time){#the-end:load-event-end-time}
        to the [current high resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#the-end:current-high-resolution-time-4
        x-internal="current-high-resolution-time"} given
        `window`{.variable}.

    9.  [Assert](https://infra.spec.whatwg.org/#assert){#the-end:assert
        x-internal="assert"}:
        [`Document`{#the-end:document-22}](dom.html#document)\'s [page
        showing](document-lifecycle.html#page-showing){#the-end:page-showing}
        is false.

    10. Set the [`Document`{#the-end:document-23}](dom.html#document)\'s
        [page
        showing](document-lifecycle.html#page-showing){#the-end:page-showing-2}
        to true.

    11. [Fire a page transition
        event](nav-history-apis.html#fire-a-page-transition-event){#the-end:fire-a-page-transition-event}
        named
        [`pageshow`{#the-end:event-pageshow}](indices.html#event-pageshow)
        at `window`{.variable} with false.

    12. [Completely finish
        loading](document-lifecycle.html#completely-finish-loading){#the-end:completely-finish-loading}
        the [`Document`{#the-end:document-24}](dom.html#document).

    13. [Queue the navigation timing
        entry](https://w3c.github.io/navigation-timing/#dfn-queue-the-navigation-timing-entry){#the-end:queue-the-navigation-timing-entry
        x-internal="queue-the-navigation-timing-entry"} for the
        [`Document`{#the-end:document-25}](dom.html#document).

10. If the [`Document`{#the-end:document-26}](dom.html#document)\'s
    [print when
    loaded](timers-and-user-prompts.html#print-when-loaded){#the-end:print-when-loaded}
    flag is set, then run the [printing
    steps](timers-and-user-prompts.html#printing-steps){#the-end:printing-steps}.

11. The [`Document`{#the-end:document-27}](dom.html#document) is now
    [ready for post-load tasks]{#ready-for-post-load-tasks .dfn}.

When the user agent is to [abort a parser]{#abort-a-parser .dfn}, it
must run the following steps:

1.  Throw away any pending content in the [input
    stream](#input-stream){#the-end:input-stream}, and discard any
    future content that would have been added to it.

2.  [Stop the speculative HTML
    parser](#stop-the-speculative-html-parser){#the-end:stop-the-speculative-html-parser-2}
    for this HTML parser.

3.  [Update the current document
    readiness](dom.html#update-the-current-document-readiness){#the-end:update-the-current-document-readiness-3}
    to \"`interactive`\".

4.  Pop *all* the nodes off the [stack of open
    elements](#stack-of-open-elements){#the-end:stack-of-open-elements-2}.

5.  [Update the current document
    readiness](dom.html#update-the-current-document-readiness){#the-end:update-the-current-document-readiness-4}
    to \"`complete`\".

#### [13.2.8]{.secno} Speculative HTML parsing[](#speculative-html-parsing){.self-link}

User agents may implement an optimization, as described in this section,
to speculatively fetch resources that are declared in the HTML markup
while the HTML parser is waiting for a [pending parsing-blocking
script](scripting.html#pending-parsing-blocking-script){#speculative-html-parsing:pending-parsing-blocking-script}
to be fetched and executed, or during normal parsing, at the time [an
element is created for a
token](#create-an-element-for-the-token){#speculative-html-parsing:create-an-element-for-the-token}.
While this optimization is not defined in precise detail, there are some
rules to consider for interoperability.

Each [HTML parser](#html-parser){#speculative-html-parsing:html-parser}
can have an [active speculative HTML
parser]{#active-speculative-html-parser .dfn export=""}. It is initially
null.

The [speculative HTML parser]{#speculative-html-parser .dfn export=""}
must act like the normal HTML parser (e.g., the tree builder rules
apply), with some exceptions:

- The state of the normal HTML parser and the document itself must not
  be affected.

  For example, the [next input
  character](#next-input-character){#speculative-html-parsing:next-input-character}
  or the [stack of open
  elements](#stack-of-open-elements){#speculative-html-parsing:stack-of-open-elements}
  for the normal HTML parser is not affected by the [speculative HTML
  parser](#speculative-html-parser){#speculative-html-parsing:speculative-html-parser}.

- Bytes pushed into the HTML parser\'s [input byte
  stream](#the-input-byte-stream){#speculative-html-parsing:the-input-byte-stream}
  must also be pushed into the speculative HTML parser\'s [input byte
  stream](#the-input-byte-stream){#speculative-html-parsing:the-input-byte-stream-2}.
  Bytes read from the streams must be independent.

- The result of the speculative parsing is primarily a series of
  [speculative
  fetches](#speculative-fetch){#speculative-html-parsing:speculative-fetch}.
  Which kinds of resources to speculatively fetch is
  [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#speculative-html-parsing:implementation-defined
  x-internal="implementation-defined"}, but user agents must not
  speculatively fetch resources that would not be fetched with the
  normal HTML parser, under the assumption that the script that is
  blocking the HTML parser does nothing.

  It is possible that the same markup is seen multiple times from the
  [speculative HTML
  parser](#speculative-html-parser){#speculative-html-parsing:speculative-html-parser-2}
  and then the normal HTML parser. It is expected that duplicated
  fetches will be prevented by caching rules, which are not yet fully
  specified.

A [speculative fetch]{#speculative-fetch .dfn} for a [speculative mock
element](#speculative-mock-element){#speculative-html-parsing:speculative-mock-element}
`element`{.variable} must follow these rules:

Should some of these things be applied to the document \"for real\",
even though they are found speculatively?

- If the [speculative HTML
  parser](#speculative-html-parser){#speculative-html-parsing:speculative-html-parser-3}
  encounters one of the following elements, then act as if that element
  is processed for the purpose of its effect of subsequent speculative
  fetches.

  - A
    [`base`{#speculative-html-parsing:the-base-element}](semantics.html#the-base-element)
    element.
  - A
    [`meta`{#speculative-html-parsing:the-meta-element}](semantics.html#the-meta-element)
    element whose
    [`http-equiv`{#speculative-html-parsing:attr-meta-http-equiv}](semantics.html#attr-meta-http-equiv)
    attribute is in the [Content security
    policy](semantics.html#attr-meta-http-equiv-content-security-policy){#speculative-html-parsing:attr-meta-http-equiv-content-security-policy}
    state.
  - A
    [`meta`{#speculative-html-parsing:the-meta-element-2}](semantics.html#the-meta-element)
    element whose
    [`name`{#speculative-html-parsing:attr-meta-name}](semantics.html#attr-meta-name)
    attribute is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#speculative-html-parsing:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for
    \"[`referrer`{#speculative-html-parsing:meta-referrer}](semantics.html#meta-referrer)\".
  - A
    [`meta`{#speculative-html-parsing:the-meta-element-3}](semantics.html#the-meta-element)
    element whose
    [`name`{#speculative-html-parsing:attr-meta-name-2}](semantics.html#attr-meta-name)
    attribute is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#speculative-html-parsing:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for \"`viewport`\". (This
    can affect whether a media query list [matches the
    environment](common-microsyntaxes.html#matches-the-environment){#speculative-html-parsing:matches-the-environment}.)
    [\[CSSDEVICEADAPT\]](references.html#refsCSSDEVICEADAPT)

- Let `url`{.variable} be the
  [URL](https://url.spec.whatwg.org/#concept-url){#speculative-html-parsing:url
  x-internal="url"} that `element`{.variable} would fetch if it was
  processed normally. If there is no such
  [URL](https://url.spec.whatwg.org/#concept-url){#speculative-html-parsing:url-2
  x-internal="url"} or if it is the empty string, then do nothing.
  Otherwise, if
  [url](https://url.spec.whatwg.org/#concept-url){#speculative-html-parsing:url-3
  x-internal="url"} is already in the [list of speculative fetch
  URLs](#list-of-speculative-fetch-urls){#speculative-html-parsing:list-of-speculative-fetch-urls},
  then do nothing. Otherwise, fetch
  [url](https://url.spec.whatwg.org/#concept-url){#speculative-html-parsing:url-4
  x-internal="url"} as if the element was processed normally, and add
  `url`{.variable} to the [list of speculative fetch
  URLs](#list-of-speculative-fetch-urls){#speculative-html-parsing:list-of-speculative-fetch-urls-2}.

Each [`Document`{#speculative-html-parsing:document}](dom.html#document)
has a [list of speculative fetch URLs]{#list-of-speculative-fetch-urls
.dfn}, which is a
[list](https://infra.spec.whatwg.org/#list){#speculative-html-parsing:list
x-internal="list"} of
[URLs](https://url.spec.whatwg.org/#concept-url){#speculative-html-parsing:url-5
x-internal="url"}, initially empty.

To [start the speculative HTML
parser]{#start-the-speculative-html-parser .dfn} for an instance of an
HTML parser `parser`{.variable}:

1.  Optionally, return.

    This step allows user agents to opt out of speculative HTML parsing.

2.  If `parser`{.variable}\'s [active speculative HTML
    parser](#active-speculative-html-parser){#speculative-html-parsing:active-speculative-html-parser}
    is not null, then [stop the speculative HTML
    parser](#stop-the-speculative-html-parser){#speculative-html-parsing:stop-the-speculative-html-parser}
    for `parser`{.variable}.

    This can happen when
    [`document.write()`{#speculative-html-parsing:dom-document-write}](dynamic-markup-insertion.html#dom-document-write)
    writes another parser-blocking script. For simplicity, this
    specification always restarts speculative parsing, but user agents
    can implement a more efficient strategy, so long as the end result
    is equivalent.

3.  Let `speculativeParser`{.variable} be a new [speculative HTML
    parser](#speculative-html-parser){#speculative-html-parsing:speculative-html-parser-4},
    with the same state as `parser`{.variable}.

4.  Let `speculativeDoc`{.variable} be a new isomorphic representation
    of `parser`{.variable}\'s
    [`Document`{#speculative-html-parsing:document-2}](dom.html#document),
    where all elements are instead [speculative mock
    elements](#speculative-mock-element){#speculative-html-parsing:speculative-mock-element-2}.
    Let `speculativeParser`{.variable} parse into
    `speculativeDoc`{.variable}.

5.  Set `parser`{.variable}\'s [active speculative HTML
    parser](#active-speculative-html-parser){#speculative-html-parsing:active-speculative-html-parser-2}
    to `speculativeParser`{.variable}.

6.  [In
    parallel](infrastructure.html#in-parallel){#speculative-html-parsing:in-parallel},
    run `speculativeParser`{.variable} until it is stopped or until it
    reaches the end of its [input
    stream](#input-stream){#speculative-html-parsing:input-stream}.

To [stop the speculative HTML parser]{#stop-the-speculative-html-parser
.dfn} for an instance of an HTML parser `parser`{.variable}:

1.  Let `speculativeParser`{.variable} be `parser`{.variable}\'s [active
    speculative HTML
    parser](#active-speculative-html-parser){#speculative-html-parsing:active-speculative-html-parser-3}.

2.  If `speculativeParser`{.variable} is null, then return.

3.  Throw away any pending content in `speculativeParser`{.variable}\'s
    [input
    stream](#input-stream){#speculative-html-parsing:input-stream-2},
    and discard any future content that would have been added to it.

4.  Set `parser`{.variable}\'s [active speculative HTML
    parser](#active-speculative-html-parser){#speculative-html-parsing:active-speculative-html-parser-4}
    to null.

The [speculative HTML
parser](#speculative-html-parser){#speculative-html-parsing:speculative-html-parser-5}
will create [speculative mock
elements](#speculative-mock-element){#speculative-html-parsing:speculative-mock-element-3}
instead of normal elements. DOM operations that the tree builder
normally does on elements are expected to work appropriately on
speculative mock elements.

A [speculative mock element]{#speculative-mock-element .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#speculative-html-parsing:struct
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#speculative-html-parsing:struct-item
x-internal="struct-item"}:

- A
  [string](https://infra.spec.whatwg.org/#string){#speculative-html-parsing:string
  x-internal="string"} [namespace]{#concept-mock-namespace .dfn},
  corresponding to an element\'s
  [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#speculative-html-parsing:concept-element-namespace
  x-internal="concept-element-namespace"}.

- A
  [string](https://infra.spec.whatwg.org/#string){#speculative-html-parsing:string-2
  x-internal="string"} [local name]{#concept-mock-local-name .dfn},
  corresponding to an element\'s [local
  name](https://dom.spec.whatwg.org/#concept-element-local-name){#speculative-html-parsing:concept-element-local-name
  x-internal="concept-element-local-name"}.

- A
  [list](https://infra.spec.whatwg.org/#list){#speculative-html-parsing:list-2
  x-internal="list"} [attribute list]{#concept-mock-attribute-list
  .dfn}, corresponding to an element\'s [attribute
  list](https://dom.spec.whatwg.org/#concept-element-attribute){#speculative-html-parsing:attribute-list
  x-internal="attribute-list"}.

- A
  [list](https://infra.spec.whatwg.org/#list){#speculative-html-parsing:list-3
  x-internal="list"} [children]{#concept-mock-children .dfn},
  corresponding to an element\'s
  [children](https://dom.spec.whatwg.org/#concept-tree-child){#speculative-html-parsing:concept-tree-child
  x-internal="concept-tree-child"}.

To [create a speculative mock
element]{#create-a-speculative-mock-element .dfn} given a
`namespace`{.variable}, `tagName`{.variable}, and
`attributes`{.variable}:

1.  Let `element`{.variable} be a new [speculative mock
    element](#speculative-mock-element){#speculative-html-parsing:speculative-mock-element-4}.

2.  Set `element`{.variable}\'s
    [namespace](#concept-mock-namespace){#speculative-html-parsing:concept-mock-namespace}
    to `namespace`{.variable}.

3.  Set `element`{.variable}\'s [local
    name](#concept-mock-local-name){#speculative-html-parsing:concept-mock-local-name}
    to `tagName`{.variable}.

4.  Set `element`{.variable}\'s [attribute
    list](#concept-mock-attribute-list){#speculative-html-parsing:concept-mock-attribute-list}
    to `attributes`{.variable}.

5.  Set `element`{.variable}\'s
    [children](#concept-mock-children){#speculative-html-parsing:concept-mock-children}
    to a new empty
    [list](https://infra.spec.whatwg.org/#list){#speculative-html-parsing:list-4
    x-internal="list"}.

6.  Optionally, perform a [speculative
    fetch](#speculative-fetch){#speculative-html-parsing:speculative-fetch-2}
    for `element`{.variable}.

7.  Return `element`{.variable}.

When the tree builder says to insert an element into a
[`template`{#speculative-html-parsing:the-template-element}](scripting.html#the-template-element)
element\'s [template
contents](scripting.html#template-contents){#speculative-html-parsing:template-contents},
if that is a [speculative mock
element](#speculative-mock-element){#speculative-html-parsing:speculative-mock-element-5},
and the
[`template`{#speculative-html-parsing:the-template-element-2}](scripting.html#the-template-element)
element\'s [template
contents](scripting.html#template-contents){#speculative-html-parsing:template-contents-2}
is not a
[`ShadowRoot`{#speculative-html-parsing:shadowroot}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}
node, instead do nothing. URLs found speculatively inside
non-declarative-shadow-root
[`template`{#speculative-html-parsing:the-template-element-3}](scripting.html#the-template-element)
elements might themselves be templates, and must not be speculatively
fetched.

#### [13.2.9]{.secno} Coercing an HTML DOM into an infoset[](#coercing-an-html-dom-into-an-infoset){.self-link}

When an application uses an [HTML
parser](#html-parser){#coercing-an-html-dom-into-an-infoset:html-parser}
in conjunction with an XML pipeline, it is possible that the constructed
DOM is not compatible with the XML tool chain in certain subtle ways.
For example, an XML toolchain might not be able to represent attributes
with the name `xmlns`, since they conflict with the Namespaces in XML
syntax. There is also some data that the [HTML
parser](#html-parser){#coercing-an-html-dom-into-an-infoset:html-parser-2}
generates that isn\'t included in the DOM itself. This section specifies
some rules for handling these issues.

If the XML API being used doesn\'t support DOCTYPEs, the tool may drop
DOCTYPEs altogether.

If the XML API doesn\'t support attributes in no namespace that are
named \"`xmlns`\", attributes whose names start with \"`xmlns:`\", or
attributes in the [XMLNS
namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#coercing-an-html-dom-into-an-infoset:xmlns-namespace
x-internal="xmlns-namespace"}, then the tool may drop such attributes.

The tool may annotate the output with any namespace declarations
required for proper operation.

If the XML API being used restricts the allowable characters in the
local names of elements and attributes, then the tool may map all
element and attribute local names that the API wouldn\'t support to a
set of names that *are* allowed, by replacing any character that isn\'t
supported with the uppercase letter U and the six digits of the
character\'s code point when expressed in hexadecimal, using digits 0-9
and capital letters A-F as the symbols, in increasing numeric order.

For example, the element name `foo<bar`, which can be output by the
[HTML
parser](#html-parser){#coercing-an-html-dom-into-an-infoset:html-parser-3},
though it is neither a legal HTML element name nor a well-formed XML
element name, would be converted into `fooU00003Cbar`, which *is* a
well-formed XML element name (though it\'s still not legal in HTML by
any means).

As another example, consider the attribute `xlink:href`. Used on a
MathML element, it becomes, after being
[adjusted](#adjust-foreign-attributes){#coercing-an-html-dom-into-an-infoset:adjust-foreign-attributes},
an attribute with a prefix \"`xlink`\" and a local name \"`href`\".
However, used on an HTML element, it becomes an attribute with no prefix
and the local name \"`xlink:href`\", which is not a valid NCName, and
thus might not be accepted by an XML API. It could thus get converted,
becoming \"`xlinkU00003Ahref`\".

The resulting names from this conversion conveniently can\'t clash with
any attribute generated by the [HTML
parser](#html-parser){#coercing-an-html-dom-into-an-infoset:html-parser-4},
since those are all either lowercase or those listed in the [adjust
foreign
attributes](#adjust-foreign-attributes){#coercing-an-html-dom-into-an-infoset:adjust-foreign-attributes-2}
algorithm\'s table.

If the XML API restricts comments from having two consecutive U+002D
HYPHEN-MINUS characters (\--), the tool may insert a single U+0020 SPACE
character between any such offending characters.

If the XML API restricts comments from ending in a U+002D HYPHEN-MINUS
character (-), the tool may insert a single U+0020 SPACE character at
the end of such comments.

If the XML API restricts allowed characters in character data, attribute
values, or comments, the tool may replace any U+000C FORM FEED (FF)
character with a U+0020 SPACE character, and any other literal non-XML
character with a U+FFFD REPLACEMENT CHARACTER.

If the tool has no way to convey out-of-band information, then the tool
may drop the following information:

- Whether the document is set to *[no-quirks
  mode](https://dom.spec.whatwg.org/#concept-document-no-quirks){x-internal="no-quirks-mode"}*,
  *[limited-quirks
  mode](https://dom.spec.whatwg.org/#concept-document-limited-quirks){x-internal="limited-quirks-mode"}*,
  or *[quirks
  mode](https://dom.spec.whatwg.org/#concept-document-quirks){x-internal="quirks-mode"}*
- The association between form controls and forms that aren\'t their
  nearest
  [`form`{#coercing-an-html-dom-into-an-infoset:the-form-element}](forms.html#the-form-element)
  element ancestor (use of the [`form` element
  pointer](#form-element-pointer){#coercing-an-html-dom-into-an-infoset:form-element-pointer}
  in the parser)
- The [template
  contents](scripting.html#template-contents){#coercing-an-html-dom-into-an-infoset:template-contents}
  of any
  [`template`{#coercing-an-html-dom-into-an-infoset:the-template-element}](scripting.html#the-template-element)
  elements.

The mutations allowed by this section apply *after* the [HTML
parser](#html-parser){#coercing-an-html-dom-into-an-infoset:html-parser-5}\'s
rules have been applied. For example, a `<a::>` start tag will be closed
by a `</a::>` end tag, and never by a `</aU00003AU00003A>` end tag, even
if the user agent is using the rules above to then generate an actual
element in the DOM with the name `aU00003AU00003A` for that start tag.

#### [13.2.10]{.secno} An introduction to error handling and strange cases in the parser[](#an-introduction-to-error-handling-and-strange-cases-in-the-parser){.self-link}

*This section is non-normative.*

This section examines some erroneous markup and discusses how the [HTML
parser](#html-parser){#an-introduction-to-error-handling-and-strange-cases-in-the-parser:html-parser}
handles these cases.

##### [13.2.10.1]{.secno} Misnested tags: \<b\>\<i\>\</b\>\</i\>[](#misnested-tags:-b-i-/b-/i){.self-link} {#misnested-tags:-b-i-/b-/i}

*This section is non-normative.*

The most-often discussed example of erroneous markup is as follows:

``` html
<p>1<b>2<i>3</b>4</i>5</p>
```

The parsing of this markup is straightforward up to the \"3\". At this
point, the DOM looks like this:

- [`html`{#misnested-tags:-b-i-/b-/i:the-html-element}](semantics.html#the-html-element)
  - [`head`{#misnested-tags:-b-i-/b-/i:the-head-element}](semantics.html#the-head-element)
  - [`body`{#misnested-tags:-b-i-/b-/i:the-body-element}](sections.html#the-body-element)
    - [`p`{#misnested-tags:-b-i-/b-/i:the-p-element}](grouping-content.html#the-p-element)
      - [`#text`{#misnested-tags:-b-i-/b-/i:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        1
      - [`b`{#misnested-tags:-b-i-/b-/i:the-b-element}](text-level-semantics.html#the-b-element)
        - [`#text`{#misnested-tags:-b-i-/b-/i:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
          2
        - [`i`{#misnested-tags:-b-i-/b-/i:the-i-element}](text-level-semantics.html#the-i-element)
          - [`#text`{#misnested-tags:-b-i-/b-/i:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
            3

Here, the [stack of open
elements](#stack-of-open-elements){#misnested-tags:-b-i-/b-/i:stack-of-open-elements}
has five elements on it:
[`html`{#misnested-tags:-b-i-/b-/i:the-html-element-2}](semantics.html#the-html-element),
[`body`{#misnested-tags:-b-i-/b-/i:the-body-element-2}](sections.html#the-body-element),
[`p`{#misnested-tags:-b-i-/b-/i:the-p-element-2}](grouping-content.html#the-p-element),
[`b`{#misnested-tags:-b-i-/b-/i:the-b-element-2}](text-level-semantics.html#the-b-element),
and
[`i`{#misnested-tags:-b-i-/b-/i:the-i-element-2}](text-level-semantics.html#the-i-element).
The [list of active formatting
elements](#list-of-active-formatting-elements){#misnested-tags:-b-i-/b-/i:list-of-active-formatting-elements}
just has two:
[`b`{#misnested-tags:-b-i-/b-/i:the-b-element-3}](text-level-semantics.html#the-b-element)
and
[`i`{#misnested-tags:-b-i-/b-/i:the-i-element-3}](text-level-semantics.html#the-i-element).
The [insertion
mode](#insertion-mode){#misnested-tags:-b-i-/b-/i:insertion-mode} is
\"[in
body](#parsing-main-inbody){#misnested-tags:-b-i-/b-/i:parsing-main-inbody}\".

Upon receiving the end tag token with the tag name \"b\", the
\"[adoption agency algorithm](#adoptionAgency)\" is invoked. This is a
simple case, in that the `formattingElement`{.variable} is the
[`b`{#misnested-tags:-b-i-/b-/i:the-b-element-4}](text-level-semantics.html#the-b-element)
element, and there is no `furthest block`{.variable}. Thus, the [stack
of open
elements](#stack-of-open-elements){#misnested-tags:-b-i-/b-/i:stack-of-open-elements-2}
ends up with just three elements:
[`html`{#misnested-tags:-b-i-/b-/i:the-html-element-3}](semantics.html#the-html-element),
[`body`{#misnested-tags:-b-i-/b-/i:the-body-element-3}](sections.html#the-body-element),
and
[`p`{#misnested-tags:-b-i-/b-/i:the-p-element-3}](grouping-content.html#the-p-element),
while the [list of active formatting
elements](#list-of-active-formatting-elements){#misnested-tags:-b-i-/b-/i:list-of-active-formatting-elements-2}
has just one:
[`i`{#misnested-tags:-b-i-/b-/i:the-i-element-4}](text-level-semantics.html#the-i-element).
The DOM tree is unmodified at this point.

The next token is a character (\"4\"), triggers the [reconstruction of
the active formatting
elements](#reconstruct-the-active-formatting-elements){#misnested-tags:-b-i-/b-/i:reconstruct-the-active-formatting-elements},
in this case just the
[`i`{#misnested-tags:-b-i-/b-/i:the-i-element-5}](text-level-semantics.html#the-i-element)
element. A new
[`i`{#misnested-tags:-b-i-/b-/i:the-i-element-6}](text-level-semantics.html#the-i-element)
element is thus created for the \"4\"
[`Text`{#misnested-tags:-b-i-/b-/i:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node. After the end tag token for the \"i\" is also received, and the
\"5\"
[`Text`{#misnested-tags:-b-i-/b-/i:text-5}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node is inserted, the DOM looks as follows:

- [`html`{#misnested-tags:-b-i-/b-/i:the-html-element-4}](semantics.html#the-html-element)
  - [`head`{#misnested-tags:-b-i-/b-/i:the-head-element-2}](semantics.html#the-head-element)
  - [`body`{#misnested-tags:-b-i-/b-/i:the-body-element-4}](sections.html#the-body-element)
    - [`p`{#misnested-tags:-b-i-/b-/i:the-p-element-4}](grouping-content.html#the-p-element)
      - [`#text`{#misnested-tags:-b-i-/b-/i:text-6}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        1
      - [`b`{#misnested-tags:-b-i-/b-/i:the-b-element-5}](text-level-semantics.html#the-b-element)
        - [`#text`{#misnested-tags:-b-i-/b-/i:text-7}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
          2
        - [`i`{#misnested-tags:-b-i-/b-/i:the-i-element-7}](text-level-semantics.html#the-i-element)
          - [`#text`{#misnested-tags:-b-i-/b-/i:text-8}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
            3
      - [`i`{#misnested-tags:-b-i-/b-/i:the-i-element-8}](text-level-semantics.html#the-i-element)
        - [`#text`{#misnested-tags:-b-i-/b-/i:text-9}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
          4
      - [`#text`{#misnested-tags:-b-i-/b-/i:text-10}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        5

##### [13.2.10.2]{.secno} Misnested tags: \<b\>\<p\>\</b\>\</p\>[](#misnested-tags:-b-p-/b-/p){.self-link} {#misnested-tags:-b-p-/b-/p}

*This section is non-normative.*

A case similar to the previous one is the following:

``` html
<b>1<p>2</b>3</p>
```

Up to the \"2\" the parsing here is straightforward:

- [`html`{#misnested-tags:-b-p-/b-/p:the-html-element}](semantics.html#the-html-element)
  - [`head`{#misnested-tags:-b-p-/b-/p:the-head-element}](semantics.html#the-head-element)
  - [`body`{#misnested-tags:-b-p-/b-/p:the-body-element}](sections.html#the-body-element)
    - [`b`{#misnested-tags:-b-p-/b-/p:the-b-element}](text-level-semantics.html#the-b-element)
      - [`#text`{#misnested-tags:-b-p-/b-/p:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        1
      - [`p`{#misnested-tags:-b-p-/b-/p:the-p-element}](grouping-content.html#the-p-element)
        - [`#text`{#misnested-tags:-b-p-/b-/p:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
          2

The interesting part is when the end tag token with the tag name \"b\"
is parsed.

Before that token is seen, the [stack of open
elements](#stack-of-open-elements){#misnested-tags:-b-p-/b-/p:stack-of-open-elements}
has four elements on it:
[`html`{#misnested-tags:-b-p-/b-/p:the-html-element-2}](semantics.html#the-html-element),
[`body`{#misnested-tags:-b-p-/b-/p:the-body-element-2}](sections.html#the-body-element),
[`b`{#misnested-tags:-b-p-/b-/p:the-b-element-2}](text-level-semantics.html#the-b-element),
and
[`p`{#misnested-tags:-b-p-/b-/p:the-p-element-2}](grouping-content.html#the-p-element).
The [list of active formatting
elements](#list-of-active-formatting-elements){#misnested-tags:-b-p-/b-/p:list-of-active-formatting-elements}
just has the one:
[`b`{#misnested-tags:-b-p-/b-/p:the-b-element-3}](text-level-semantics.html#the-b-element).
The [insertion
mode](#insertion-mode){#misnested-tags:-b-p-/b-/p:insertion-mode} is
\"[in
body](#parsing-main-inbody){#misnested-tags:-b-p-/b-/p:parsing-main-inbody}\".

Upon receiving the end tag token with the tag name \"b\", the
\"[adoption agency algorithm](#adoptionAgency)\" is invoked, as in the
previous example. However, in this case, there *is* a
`furthest block`{.variable}, namely the
[`p`{#misnested-tags:-b-p-/b-/p:the-p-element-3}](grouping-content.html#the-p-element)
element. Thus, this time the adoption agency algorithm isn\'t skipped
over.

The `common ancestor`{.variable} is the
[`body`{#misnested-tags:-b-p-/b-/p:the-body-element-3}](sections.html#the-body-element)
element. A conceptual \"bookmark\" marks the position of the
[`b`{#misnested-tags:-b-p-/b-/p:the-b-element-4}](text-level-semantics.html#the-b-element)
in the [list of active formatting
elements](#list-of-active-formatting-elements){#misnested-tags:-b-p-/b-/p:list-of-active-formatting-elements-2},
but since that list has only one element in it, the bookmark won\'t have
much effect.

As the algorithm progresses, `node`{.variable} ends up set to the
formatting element
([`b`{#misnested-tags:-b-p-/b-/p:the-b-element-5}](text-level-semantics.html#the-b-element)),
and `last node`{.variable} ends up set to the
`furthest block`{.variable}
([`p`{#misnested-tags:-b-p-/b-/p:the-p-element-4}](grouping-content.html#the-p-element)).

The `last node`{.variable} gets appended (moved) to the
`common ancestor`{.variable}, so that the DOM looks like:

- [`html`{#misnested-tags:-b-p-/b-/p:the-html-element-3}](semantics.html#the-html-element)
  - [`head`{#misnested-tags:-b-p-/b-/p:the-head-element-2}](semantics.html#the-head-element)
  - [`body`{#misnested-tags:-b-p-/b-/p:the-body-element-4}](sections.html#the-body-element)
    - [`b`{#misnested-tags:-b-p-/b-/p:the-b-element-6}](text-level-semantics.html#the-b-element)
      - [`#text`{#misnested-tags:-b-p-/b-/p:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        1
    - [`p`{#misnested-tags:-b-p-/b-/p:the-p-element-5}](grouping-content.html#the-p-element)
      - [`#text`{#misnested-tags:-b-p-/b-/p:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        2

A new
[`b`{#misnested-tags:-b-p-/b-/p:the-b-element-7}](text-level-semantics.html#the-b-element)
element is created, and the children of the
[`p`{#misnested-tags:-b-p-/b-/p:the-p-element-6}](grouping-content.html#the-p-element)
element are moved to it:

- [`html`{#misnested-tags:-b-p-/b-/p:the-html-element-4}](semantics.html#the-html-element)
  - [`head`{#misnested-tags:-b-p-/b-/p:the-head-element-3}](semantics.html#the-head-element)
  - [`body`{#misnested-tags:-b-p-/b-/p:the-body-element-5}](sections.html#the-body-element)
    - [`b`{#misnested-tags:-b-p-/b-/p:the-b-element-8}](text-level-semantics.html#the-b-element)
      - [`#text`{#misnested-tags:-b-p-/b-/p:text-5}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        1
    - [`p`{#misnested-tags:-b-p-/b-/p:the-p-element-7}](grouping-content.html#the-p-element)

<!-- -->

- [`b`{#misnested-tags:-b-p-/b-/p:the-b-element-9}](text-level-semantics.html#the-b-element)
  - [`#text`{#misnested-tags:-b-p-/b-/p:text-6}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
    2

Finally, the new
[`b`{#misnested-tags:-b-p-/b-/p:the-b-element-10}](text-level-semantics.html#the-b-element)
element is appended to the
[`p`{#misnested-tags:-b-p-/b-/p:the-p-element-8}](grouping-content.html#the-p-element)
element, so that the DOM looks like:

- [`html`{#misnested-tags:-b-p-/b-/p:the-html-element-5}](semantics.html#the-html-element)
  - [`head`{#misnested-tags:-b-p-/b-/p:the-head-element-4}](semantics.html#the-head-element)
  - [`body`{#misnested-tags:-b-p-/b-/p:the-body-element-6}](sections.html#the-body-element)
    - [`b`{#misnested-tags:-b-p-/b-/p:the-b-element-11}](text-level-semantics.html#the-b-element)
      - [`#text`{#misnested-tags:-b-p-/b-/p:text-7}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        1
    - [`p`{#misnested-tags:-b-p-/b-/p:the-p-element-9}](grouping-content.html#the-p-element)
      - [`b`{#misnested-tags:-b-p-/b-/p:the-b-element-12}](text-level-semantics.html#the-b-element)
        - [`#text`{#misnested-tags:-b-p-/b-/p:text-8}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
          2

The
[`b`{#misnested-tags:-b-p-/b-/p:the-b-element-13}](text-level-semantics.html#the-b-element)
element is removed from the [list of active formatting
elements](#list-of-active-formatting-elements){#misnested-tags:-b-p-/b-/p:list-of-active-formatting-elements-3}
and the [stack of open
elements](#stack-of-open-elements){#misnested-tags:-b-p-/b-/p:stack-of-open-elements-2},
so that when the \"3\" is parsed, it is appended to the
[`p`{#misnested-tags:-b-p-/b-/p:the-p-element-10}](grouping-content.html#the-p-element)
element:

- [`html`{#misnested-tags:-b-p-/b-/p:the-html-element-6}](semantics.html#the-html-element)
  - [`head`{#misnested-tags:-b-p-/b-/p:the-head-element-5}](semantics.html#the-head-element)
  - [`body`{#misnested-tags:-b-p-/b-/p:the-body-element-7}](sections.html#the-body-element)
    - [`b`{#misnested-tags:-b-p-/b-/p:the-b-element-14}](text-level-semantics.html#the-b-element)
      - [`#text`{#misnested-tags:-b-p-/b-/p:text-9}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        1
    - [`p`{#misnested-tags:-b-p-/b-/p:the-p-element-11}](grouping-content.html#the-p-element)
      - [`b`{#misnested-tags:-b-p-/b-/p:the-b-element-15}](text-level-semantics.html#the-b-element)
        - [`#text`{#misnested-tags:-b-p-/b-/p:text-10}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
          2
      - [`#text`{#misnested-tags:-b-p-/b-/p:text-11}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        3

##### [13.2.10.3]{.secno} Unexpected markup in tables[](#unexpected-markup-in-tables){.self-link}

*This section is non-normative.*

Error handling in tables is, for historical reasons, especially strange.
For example, consider the following markup:

``` html
<table><b><tr><td>aaa</td></tr>bbb</table>ccc
```

The highlighted
[`b`{#unexpected-markup-in-tables:the-b-element}](text-level-semantics.html#the-b-element)
element start tag is not allowed directly inside a table like that, and
the parser handles this case by placing the element *before* the table.
(This is called *[foster parenting](#foster-parent)*.) This can be seen
by examining the DOM tree as it stands just after the
[`table`{#unexpected-markup-in-tables:the-table-element}](tables.html#the-table-element)
element\'s start tag has been seen:

- [`html`{#unexpected-markup-in-tables:the-html-element}](semantics.html#the-html-element)
  - [`head`{#unexpected-markup-in-tables:the-head-element}](semantics.html#the-head-element)
  - [`body`{#unexpected-markup-in-tables:the-body-element}](sections.html#the-body-element)
    - [`table`{#unexpected-markup-in-tables:the-table-element-2}](tables.html#the-table-element)

\...and then immediately after the
[`b`{#unexpected-markup-in-tables:the-b-element-2}](text-level-semantics.html#the-b-element)
element start tag has been seen:

- [`html`{#unexpected-markup-in-tables:the-html-element-2}](semantics.html#the-html-element)
  - [`head`{#unexpected-markup-in-tables:the-head-element-2}](semantics.html#the-head-element)
  - [`body`{#unexpected-markup-in-tables:the-body-element-2}](sections.html#the-body-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-3}](text-level-semantics.html#the-b-element)
    - [`table`{#unexpected-markup-in-tables:the-table-element-3}](tables.html#the-table-element)

At this point, the [stack of open
elements](#stack-of-open-elements){#unexpected-markup-in-tables:stack-of-open-elements}
has on it the elements
[`html`{#unexpected-markup-in-tables:the-html-element-3}](semantics.html#the-html-element),
[`body`{#unexpected-markup-in-tables:the-body-element-3}](sections.html#the-body-element),
[`table`{#unexpected-markup-in-tables:the-table-element-4}](tables.html#the-table-element),
and
[`b`{#unexpected-markup-in-tables:the-b-element-4}](text-level-semantics.html#the-b-element)
(in that order, despite the resulting DOM tree); the [list of active
formatting
elements](#list-of-active-formatting-elements){#unexpected-markup-in-tables:list-of-active-formatting-elements}
just has the
[`b`{#unexpected-markup-in-tables:the-b-element-5}](text-level-semantics.html#the-b-element)
element in it; and the [insertion
mode](#insertion-mode){#unexpected-markup-in-tables:insertion-mode} is
\"[in
table](#parsing-main-intable){#unexpected-markup-in-tables:parsing-main-intable}\".

The
[`tr`{#unexpected-markup-in-tables:the-tr-element}](tables.html#the-tr-element)
start tag causes the
[`b`{#unexpected-markup-in-tables:the-b-element-6}](text-level-semantics.html#the-b-element)
element to be popped off the stack and a
[`tbody`{#unexpected-markup-in-tables:the-tbody-element}](tables.html#the-tbody-element)
start tag to be implied; the
[`tbody`{#unexpected-markup-in-tables:the-tbody-element-2}](tables.html#the-tbody-element)
and
[`tr`{#unexpected-markup-in-tables:the-tr-element-2}](tables.html#the-tr-element)
elements are then handled in a rather straight-forward manner, taking
the parser through the \"[in table
body](#parsing-main-intbody){#unexpected-markup-in-tables:parsing-main-intbody}\"
and \"[in
row](#parsing-main-intr){#unexpected-markup-in-tables:parsing-main-intr}\"
insertion modes, after which the DOM looks as follows:

- [`html`{#unexpected-markup-in-tables:the-html-element-4}](semantics.html#the-html-element)
  - [`head`{#unexpected-markup-in-tables:the-head-element-3}](semantics.html#the-head-element)
  - [`body`{#unexpected-markup-in-tables:the-body-element-4}](sections.html#the-body-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-7}](text-level-semantics.html#the-b-element)
    - [`table`{#unexpected-markup-in-tables:the-table-element-5}](tables.html#the-table-element)
      - [`tbody`{#unexpected-markup-in-tables:the-tbody-element-3}](tables.html#the-tbody-element)
        - [`tr`{#unexpected-markup-in-tables:the-tr-element-3}](tables.html#the-tr-element)

Here, the [stack of open
elements](#stack-of-open-elements){#unexpected-markup-in-tables:stack-of-open-elements-2}
has on it the elements
[`html`{#unexpected-markup-in-tables:the-html-element-5}](semantics.html#the-html-element),
[`body`{#unexpected-markup-in-tables:the-body-element-5}](sections.html#the-body-element),
[`table`{#unexpected-markup-in-tables:the-table-element-6}](tables.html#the-table-element),
[`tbody`{#unexpected-markup-in-tables:the-tbody-element-4}](tables.html#the-tbody-element),
and
[`tr`{#unexpected-markup-in-tables:the-tr-element-4}](tables.html#the-tr-element);
the [list of active formatting
elements](#list-of-active-formatting-elements){#unexpected-markup-in-tables:list-of-active-formatting-elements-2}
still has the
[`b`{#unexpected-markup-in-tables:the-b-element-8}](text-level-semantics.html#the-b-element)
element in it; and the [insertion
mode](#insertion-mode){#unexpected-markup-in-tables:insertion-mode-2} is
\"[in
row](#parsing-main-intr){#unexpected-markup-in-tables:parsing-main-intr-2}\".

The
[`td`{#unexpected-markup-in-tables:the-td-element}](tables.html#the-td-element)
element start tag token, after putting a
[`td`{#unexpected-markup-in-tables:the-td-element-2}](tables.html#the-td-element)
element on the tree, puts a
[marker](#concept-parser-marker){#unexpected-markup-in-tables:concept-parser-marker}
on the [list of active formatting
elements](#list-of-active-formatting-elements){#unexpected-markup-in-tables:list-of-active-formatting-elements-3}
(it also switches to the \"[in
cell](#parsing-main-intd){#unexpected-markup-in-tables:parsing-main-intd}\"
[insertion
mode](#insertion-mode){#unexpected-markup-in-tables:insertion-mode-3}).

- [`html`{#unexpected-markup-in-tables:the-html-element-6}](semantics.html#the-html-element)
  - [`head`{#unexpected-markup-in-tables:the-head-element-4}](semantics.html#the-head-element)
  - [`body`{#unexpected-markup-in-tables:the-body-element-6}](sections.html#the-body-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-9}](text-level-semantics.html#the-b-element)
    - [`table`{#unexpected-markup-in-tables:the-table-element-7}](tables.html#the-table-element)
      - [`tbody`{#unexpected-markup-in-tables:the-tbody-element-5}](tables.html#the-tbody-element)
        - [`tr`{#unexpected-markup-in-tables:the-tr-element-5}](tables.html#the-tr-element)
          - [`td`{#unexpected-markup-in-tables:the-td-element-3}](tables.html#the-td-element)

The
[marker](#concept-parser-marker){#unexpected-markup-in-tables:concept-parser-marker-2}
means that when the \"aaa\" character tokens are seen, no
[`b`{#unexpected-markup-in-tables:the-b-element-10}](text-level-semantics.html#the-b-element)
element is created to hold the resulting
[`Text`{#unexpected-markup-in-tables:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node:

- [`html`{#unexpected-markup-in-tables:the-html-element-7}](semantics.html#the-html-element)
  - [`head`{#unexpected-markup-in-tables:the-head-element-5}](semantics.html#the-head-element)
  - [`body`{#unexpected-markup-in-tables:the-body-element-7}](sections.html#the-body-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-11}](text-level-semantics.html#the-b-element)
    - [`table`{#unexpected-markup-in-tables:the-table-element-8}](tables.html#the-table-element)
      - [`tbody`{#unexpected-markup-in-tables:the-tbody-element-6}](tables.html#the-tbody-element)
        - [`tr`{#unexpected-markup-in-tables:the-tr-element-6}](tables.html#the-tr-element)
          - [`td`{#unexpected-markup-in-tables:the-td-element-4}](tables.html#the-td-element)
            - [`#text`{#unexpected-markup-in-tables:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
              aaa

The end tags are handled in a straight-forward manner; after handling
them, the [stack of open
elements](#stack-of-open-elements){#unexpected-markup-in-tables:stack-of-open-elements-3}
has on it the elements
[`html`{#unexpected-markup-in-tables:the-html-element-8}](semantics.html#the-html-element),
[`body`{#unexpected-markup-in-tables:the-body-element-8}](sections.html#the-body-element),
[`table`{#unexpected-markup-in-tables:the-table-element-9}](tables.html#the-table-element),
and
[`tbody`{#unexpected-markup-in-tables:the-tbody-element-7}](tables.html#the-tbody-element);
the [list of active formatting
elements](#list-of-active-formatting-elements){#unexpected-markup-in-tables:list-of-active-formatting-elements-4}
still has the
[`b`{#unexpected-markup-in-tables:the-b-element-12}](text-level-semantics.html#the-b-element)
element in it (the
[marker](#concept-parser-marker){#unexpected-markup-in-tables:concept-parser-marker-3}
having been removed by the \"td\" end tag token); and the [insertion
mode](#insertion-mode){#unexpected-markup-in-tables:insertion-mode-4} is
\"[in table
body](#parsing-main-intbody){#unexpected-markup-in-tables:parsing-main-intbody-2}\".

Thus it is that the \"bbb\" character tokens are found. These trigger
the \"[in table
text](#parsing-main-intabletext){#unexpected-markup-in-tables:parsing-main-intabletext}\"
insertion mode to be used (with the [original insertion
mode](#original-insertion-mode){#unexpected-markup-in-tables:original-insertion-mode}
set to \"[in table
body](#parsing-main-intbody){#unexpected-markup-in-tables:parsing-main-intbody-3}\").
The character tokens are collected, and when the next token (the
[`table`{#unexpected-markup-in-tables:the-table-element-10}](tables.html#the-table-element)
element end tag) is seen, they are processed as a group. Since they are
not all spaces, they are handled as per the \"anything else\" rules in
the \"[in
table](#parsing-main-intable){#unexpected-markup-in-tables:parsing-main-intable-2}\"
insertion mode, which defer to the \"[in
body](#parsing-main-inbody){#unexpected-markup-in-tables:parsing-main-inbody}\"
insertion mode but with [foster
parenting](#foster-parent){#unexpected-markup-in-tables:foster-parent-2}.

When [the active formatting elements are
reconstructed](#reconstruct-the-active-formatting-elements){#unexpected-markup-in-tables:reconstruct-the-active-formatting-elements},
a
[`b`{#unexpected-markup-in-tables:the-b-element-13}](text-level-semantics.html#the-b-element)
element is created and [foster
parented](#foster-parent){#unexpected-markup-in-tables:foster-parent-3},
and then the \"bbb\"
[`Text`{#unexpected-markup-in-tables:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node is appended to it:

- [`html`{#unexpected-markup-in-tables:the-html-element-9}](semantics.html#the-html-element)
  - [`head`{#unexpected-markup-in-tables:the-head-element-6}](semantics.html#the-head-element)
  - [`body`{#unexpected-markup-in-tables:the-body-element-9}](sections.html#the-body-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-14}](text-level-semantics.html#the-b-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-15}](text-level-semantics.html#the-b-element)
      - [`#text`{#unexpected-markup-in-tables:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        bbb
    - [`table`{#unexpected-markup-in-tables:the-table-element-11}](tables.html#the-table-element)
      - [`tbody`{#unexpected-markup-in-tables:the-tbody-element-8}](tables.html#the-tbody-element)
        - [`tr`{#unexpected-markup-in-tables:the-tr-element-7}](tables.html#the-tr-element)
          - [`td`{#unexpected-markup-in-tables:the-td-element-5}](tables.html#the-td-element)
            - [`#text`{#unexpected-markup-in-tables:text-5}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
              aaa

The [stack of open
elements](#stack-of-open-elements){#unexpected-markup-in-tables:stack-of-open-elements-4}
has on it the elements
[`html`{#unexpected-markup-in-tables:the-html-element-10}](semantics.html#the-html-element),
[`body`{#unexpected-markup-in-tables:the-body-element-10}](sections.html#the-body-element),
[`table`{#unexpected-markup-in-tables:the-table-element-12}](tables.html#the-table-element),
[`tbody`{#unexpected-markup-in-tables:the-tbody-element-9}](tables.html#the-tbody-element),
and the new
[`b`{#unexpected-markup-in-tables:the-b-element-16}](text-level-semantics.html#the-b-element)
(again, note that this doesn\'t match the resulting tree!); the [list of
active formatting
elements](#list-of-active-formatting-elements){#unexpected-markup-in-tables:list-of-active-formatting-elements-5}
has the new
[`b`{#unexpected-markup-in-tables:the-b-element-17}](text-level-semantics.html#the-b-element)
element in it; and the [insertion
mode](#insertion-mode){#unexpected-markup-in-tables:insertion-mode-5} is
still \"[in table
body](#parsing-main-intbody){#unexpected-markup-in-tables:parsing-main-intbody-4}\".

Had the character tokens been only [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#unexpected-markup-in-tables:space-characters
x-internal="space-characters"} instead of \"bbb\", then that [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#unexpected-markup-in-tables:space-characters-2
x-internal="space-characters"} would just be appended to the
[`tbody`{#unexpected-markup-in-tables:the-tbody-element-10}](tables.html#the-tbody-element)
element.

Finally, the
[`table`{#unexpected-markup-in-tables:the-table-element-13}](tables.html#the-table-element)
is closed by a \"table\" end tag. This pops all the nodes from the
[stack of open
elements](#stack-of-open-elements){#unexpected-markup-in-tables:stack-of-open-elements-5}
up to and including the
[`table`{#unexpected-markup-in-tables:the-table-element-14}](tables.html#the-table-element)
element, but it doesn\'t affect the [list of active formatting
elements](#list-of-active-formatting-elements){#unexpected-markup-in-tables:list-of-active-formatting-elements-6},
so the \"ccc\" character tokens after the table result in yet another
[`b`{#unexpected-markup-in-tables:the-b-element-18}](text-level-semantics.html#the-b-element)
element being created, this time after the table:

- [`html`{#unexpected-markup-in-tables:the-html-element-11}](semantics.html#the-html-element)
  - [`head`{#unexpected-markup-in-tables:the-head-element-7}](semantics.html#the-head-element)
  - [`body`{#unexpected-markup-in-tables:the-body-element-11}](sections.html#the-body-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-19}](text-level-semantics.html#the-b-element)
    - [`b`{#unexpected-markup-in-tables:the-b-element-20}](text-level-semantics.html#the-b-element)
      - [`#text`{#unexpected-markup-in-tables:text-6}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        bbb
    - [`table`{#unexpected-markup-in-tables:the-table-element-15}](tables.html#the-table-element)
      - [`tbody`{#unexpected-markup-in-tables:the-tbody-element-11}](tables.html#the-tbody-element)
        - [`tr`{#unexpected-markup-in-tables:the-tr-element-8}](tables.html#the-tr-element)
          - [`td`{#unexpected-markup-in-tables:the-td-element-6}](tables.html#the-td-element)
            - [`#text`{#unexpected-markup-in-tables:text-7}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
              aaa
    - [`b`{#unexpected-markup-in-tables:the-b-element-21}](text-level-semantics.html#the-b-element)
      - [`#text`{#unexpected-markup-in-tables:text-8}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        ccc

##### [13.2.10.4]{.secno} Scripts that modify the page as it is being parsed[](#scripts-that-modify-the-page-as-it-is-being-parsed){.self-link}

*This section is non-normative.*

Consider the following markup, which for this example we will assume is
the document with
[URL](https://url.spec.whatwg.org/#concept-url){#scripts-that-modify-the-page-as-it-is-being-parsed:url
x-internal="url"} `https://example.com/inner`, being rendered as the
content of an
[`iframe`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
in another document with the
[URL](https://url.spec.whatwg.org/#concept-url){#scripts-that-modify-the-page-as-it-is-being-parsed:url-2
x-internal="url"} `https://example.com/outer`:

``` html
<div id=a>
 <script>
  var div = document.getElementById('a');
  parent.document.body.appendChild(div);
 </script>
 <script>
  alert(document.URL);
 </script>
</div>
<script>
 alert(document.URL);
</script>
```

Up to the first \"script\" end tag, before the script is parsed, the
result is relatively straightforward:

- [`html`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-html-element}](semantics.html#the-html-element)
  - [`head`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-head-element}](semantics.html#the-head-element)
  - [`body`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-body-element}](sections.html#the-body-element)
    - [`div`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-div-element}](grouping-content.html#the-div-element)
      [[`id`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-id-attribute
      .attribute .name}](dom.html#the-id-attribute)=\"`a`{.attribute
      .value}\"]{.t2}
      - [`#text`{#scripts-that-modify-the-page-as-it-is-being-parsed:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
      - [`script`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-script-element}](scripting.html#the-script-element)
        - [`#text`{#scripts-that-modify-the-page-as-it-is-being-parsed:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
          var div = document.getElementById(\'a\'); ⏎
          parent.document.body.appendChild(div);

After the script is parsed, though, the
[`div`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-div-element-2}](grouping-content.html#the-div-element)
element and its child
[`script`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-script-element-2}](scripting.html#the-script-element)
element are gone:

- [`html`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-html-element-2}](semantics.html#the-html-element)
  - [`head`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-head-element-2}](semantics.html#the-head-element)
  - [`body`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-body-element-2}](sections.html#the-body-element)

They are, at this point, in the
[`Document`{#scripts-that-modify-the-page-as-it-is-being-parsed:document}](dom.html#document)
of the aforementioned outer [browsing
context](document-sequences.html#browsing-context){#scripts-that-modify-the-page-as-it-is-being-parsed:browsing-context}.
However, the [stack of open
elements](#stack-of-open-elements){#scripts-that-modify-the-page-as-it-is-being-parsed:stack-of-open-elements}
*still contains the
[`div`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-div-element-3}](grouping-content.html#the-div-element)
element*.

Thus, when the second
[`script`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-script-element-3}](scripting.html#the-script-element)
element is parsed, it is inserted *into the outer
[`Document`{#scripts-that-modify-the-page-as-it-is-being-parsed:document-2}](dom.html#document)
object*.

Those parsed into different
[`Document`{#scripts-that-modify-the-page-as-it-is-being-parsed:document-3}](dom.html#document)s
than the one the parser was created for do not execute, so the first
alert does not show.

Once the
[`div`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-div-element-4}](grouping-content.html#the-div-element)
element\'s end tag is parsed, the
[`div`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-div-element-5}](grouping-content.html#the-div-element)
element is popped off the stack, and so the next
[`script`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-script-element-4}](scripting.html#the-script-element)
element is in the inner
[`Document`{#scripts-that-modify-the-page-as-it-is-being-parsed:document-4}](dom.html#document):

- [`html`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-html-element-3}](semantics.html#the-html-element)
  - [`head`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-head-element-3}](semantics.html#the-head-element)
  - [`body`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-body-element-3}](sections.html#the-body-element)
    - [`script`{#scripts-that-modify-the-page-as-it-is-being-parsed:the-script-element-5}](scripting.html#the-script-element)
      - [`#text`{#scripts-that-modify-the-page-as-it-is-being-parsed:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        alert(document.URL);

This script does execute, resulting in an alert that says
\"https://example.com/inner\".

##### [13.2.10.5]{.secno} The execution of scripts that are moving across multiple documents[](#the-execution-of-scripts-that-are-moving-across-multiple-documents){.self-link}

*This section is non-normative.*

Elaborating on the example in the previous section, consider the case
where the second
[`script`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:the-script-element}](scripting.html#the-script-element)
element is an external script (i.e. one with a
[`src`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:attr-script-src}](scripting.html#attr-script-src)
attribute). Since the element was not in the parser\'s
[`Document`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:document}](dom.html#document)
when it was created, that external script is not even downloaded.

In a case where a
[`script`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:the-script-element-2}](scripting.html#the-script-element)
element with a
[`src`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:attr-script-src-2}](scripting.html#attr-script-src)
attribute is parsed normally into its parser\'s
[`Document`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:document-2}](dom.html#document),
but while the external script is being downloaded, the element is moved
to another document, the script continues to download, but does not
execute.

In general, moving
[`script`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:the-script-element-3}](scripting.html#the-script-element)
elements between
[`Document`{#the-execution-of-scripts-that-are-moving-across-multiple-documents:document-3}](dom.html#document)s
is considered a bad practice.

##### [13.2.10.6]{.secno} Unclosed formatting elements[](#unclosed-formatting-elements){.self-link}

*This section is non-normative.*

The following markup shows how nested formatting elements (such as
[`b`{#unclosed-formatting-elements:the-b-element}](text-level-semantics.html#the-b-element))
get collected and continue to be applied even as the elements they are
contained in are closed, but that excessive duplicates are thrown away.

``` html
<!DOCTYPE html>
<p><b class=x><b class=x><b><b class=x><b class=x><b>X
<p>X
<p><b><b class=x><b>X
<p></b></b></b></b></b></b>X
```

The resulting DOM tree is as follows:

- DOCTYPE: `html`
- [`html`{#unclosed-formatting-elements:the-html-element}](semantics.html#the-html-element)
  - [`head`{#unclosed-formatting-elements:the-head-element}](semantics.html#the-head-element)
  - [`body`{#unclosed-formatting-elements:the-body-element}](sections.html#the-body-element)
    - [`p`{#unclosed-formatting-elements:the-p-element}](grouping-content.html#the-p-element)
      - [`b`{#unclosed-formatting-elements:the-b-element-2}](text-level-semantics.html#the-b-element)
        [[`class`{#unclosed-formatting-elements:classes .attribute
        .name}](dom.html#classes)=\"`x`{.attribute .value}\"]{.t2}
        - [`b`{#unclosed-formatting-elements:the-b-element-3}](text-level-semantics.html#the-b-element)
          [[`class`{#unclosed-formatting-elements:classes-2 .attribute
          .name}](dom.html#classes)=\"`x`{.attribute .value}\"]{.t2}
          - [`b`{#unclosed-formatting-elements:the-b-element-4}](text-level-semantics.html#the-b-element)
            - [`b`{#unclosed-formatting-elements:the-b-element-5}](text-level-semantics.html#the-b-element)
              [[`class`{#unclosed-formatting-elements:classes-3
              .attribute .name}](dom.html#classes)=\"`x`{.attribute
              .value}\"]{.t2}
              - [`b`{#unclosed-formatting-elements:the-b-element-6}](text-level-semantics.html#the-b-element)
                [[`class`{#unclosed-formatting-elements:classes-4
                .attribute .name}](dom.html#classes)=\"`x`{.attribute
                .value}\"]{.t2}
                - [`b`{#unclosed-formatting-elements:the-b-element-7}](text-level-semantics.html#the-b-element)
                  - [`#text`{#unclosed-formatting-elements:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
                    X⏎
    - [`p`{#unclosed-formatting-elements:the-p-element-2}](grouping-content.html#the-p-element)
      - [`b`{#unclosed-formatting-elements:the-b-element-8}](text-level-semantics.html#the-b-element)
        [[`class`{#unclosed-formatting-elements:classes-5 .attribute
        .name}](dom.html#classes)=\"`x`{.attribute .value}\"]{.t2}
        - [`b`{#unclosed-formatting-elements:the-b-element-9}](text-level-semantics.html#the-b-element)
          - [`b`{#unclosed-formatting-elements:the-b-element-10}](text-level-semantics.html#the-b-element)
            [[`class`{#unclosed-formatting-elements:classes-6 .attribute
            .name}](dom.html#classes)=\"`x`{.attribute .value}\"]{.t2}
            - [`b`{#unclosed-formatting-elements:the-b-element-11}](text-level-semantics.html#the-b-element)
              [[`class`{#unclosed-formatting-elements:classes-7
              .attribute .name}](dom.html#classes)=\"`x`{.attribute
              .value}\"]{.t2}
              - [`b`{#unclosed-formatting-elements:the-b-element-12}](text-level-semantics.html#the-b-element)
                - [`#text`{#unclosed-formatting-elements:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
                  X⏎
    - [`p`{#unclosed-formatting-elements:the-p-element-3}](grouping-content.html#the-p-element)
      - [`b`{#unclosed-formatting-elements:the-b-element-13}](text-level-semantics.html#the-b-element)
        [[`class`{#unclosed-formatting-elements:classes-8 .attribute
        .name}](dom.html#classes)=\"`x`{.attribute .value}\"]{.t2}
        - [`b`{#unclosed-formatting-elements:the-b-element-14}](text-level-semantics.html#the-b-element)
          - [`b`{#unclosed-formatting-elements:the-b-element-15}](text-level-semantics.html#the-b-element)
            [[`class`{#unclosed-formatting-elements:classes-9 .attribute
            .name}](dom.html#classes)=\"`x`{.attribute .value}\"]{.t2}
            - [`b`{#unclosed-formatting-elements:the-b-element-16}](text-level-semantics.html#the-b-element)
              [[`class`{#unclosed-formatting-elements:classes-10
              .attribute .name}](dom.html#classes)=\"`x`{.attribute
              .value}\"]{.t2}
              - [`b`{#unclosed-formatting-elements:the-b-element-17}](text-level-semantics.html#the-b-element)
                - [`b`{#unclosed-formatting-elements:the-b-element-18}](text-level-semantics.html#the-b-element)
                  - [`b`{#unclosed-formatting-elements:the-b-element-19}](text-level-semantics.html#the-b-element)
                    [[`class`{#unclosed-formatting-elements:classes-11
                    .attribute
                    .name}](dom.html#classes)=\"`x`{.attribute
                    .value}\"]{.t2}
                    - [`b`{#unclosed-formatting-elements:the-b-element-20}](text-level-semantics.html#the-b-element)
                      - [`#text`{#unclosed-formatting-elements:text-3}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
                        X⏎
    - [`p`{#unclosed-formatting-elements:the-p-element-4}](grouping-content.html#the-p-element)
      - [`#text`{#unclosed-formatting-elements:text-4}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}:
        X⏎

Note how the second
[`p`{#unclosed-formatting-elements:the-p-element-5}](grouping-content.html#the-p-element)
element in the markup has no explicit
[`b`{#unclosed-formatting-elements:the-b-element-21}](text-level-semantics.html#the-b-element)
elements, but in the resulting DOM, up to three of each kind of
formatting element (in this case three
[`b`{#unclosed-formatting-elements:the-b-element-22}](text-level-semantics.html#the-b-element)
elements with the class attribute, and two unadorned
[`b`{#unclosed-formatting-elements:the-b-element-23}](text-level-semantics.html#the-b-element)
elements) get reconstructed before the element\'s \"X\".

Also note how this means that in the final paragraph only six
[`b`{#unclosed-formatting-elements:the-b-element-24}](text-level-semantics.html#the-b-element)
end tags are needed to completely clear the [list of active formatting
elements](#list-of-active-formatting-elements){#unclosed-formatting-elements:list-of-active-formatting-elements},
even though nine
[`b`{#unclosed-formatting-elements:the-b-element-25}](text-level-semantics.html#the-b-element)
start tags have been seen up to this point.

### [13.3]{.secno} Serializing HTML fragments[](#serialising-html-fragments){.self-link} {#serialising-html-fragments}

For the purposes of the following algorithm, an element [serializes as
void]{#serializes-as-void .dfn} if its element type is one of the [void
elements](syntax.html#void-elements){#serialising-html-fragments:void-elements},
or is
[`basefont`{#serialising-html-fragments:basefont}](obsolete.html#basefont),
[`bgsound`{#serialising-html-fragments:bgsound}](obsolete.html#bgsound),
[`frame`{#serialising-html-fragments:frame}](obsolete.html#frame),
[`keygen`{#serialising-html-fragments:keygen}](obsolete.html#keygen), or
[`param`{#serialising-html-fragments:param}](obsolete.html#param).

The following steps form the [HTML fragment serialization
algorithm]{#html-fragment-serialisation-algorithm .dfn export=""}. The
algorithm takes as input a DOM
[`Element`{#serialising-html-fragments:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
[`Document`{#serialising-html-fragments:document}](dom.html#document),
or
[`DocumentFragment`{#serialising-html-fragments:documentfragment}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
referred to as `the node`{.variable}, a boolean
`serializableShadowRoots`{.variable}, and a `sequence<ShadowRoot>`
`shadowRoots`{.variable}, and returns a string.

This algorithm serializes the *children* of the node being serialized,
not the node itself.

1.  If `the node`{.variable} [serializes as
    void](#serializes-as-void){#serialising-html-fragments:serializes-as-void},
    then return the empty string.

2.  Let `s`{.variable} be a string, and initialize it to the empty
    string.

3.  If `the node`{.variable} is a
    [`template`{#serialising-html-fragments:the-template-element}](scripting.html#the-template-element)
    element, then let `the node`{.variable} instead be the
    [`template`{#serialising-html-fragments:the-template-element-2}](scripting.html#the-template-element)
    element\'s [template
    contents](scripting.html#template-contents){#serialising-html-fragments:template-contents}
    (a
    [`DocumentFragment`{#serialising-html-fragments:documentfragment-2}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    node).

4.  If `current node`{.variable} is a [shadow
    host](https://dom.spec.whatwg.org/#element-shadow-host){#serialising-html-fragments:shadow-host
    x-internal="shadow-host"}, then:

    1.  Let `shadow`{.variable} be `current node`{.variable}\'s [shadow
        root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#serialising-html-fragments:concept-element-shadow-root
        x-internal="concept-element-shadow-root"}.

    2.  If one of the following is true:

        - `serializableShadowRoots`{.variable} is true and
          `shadow`{.variable}\'s
          [serializable](https://dom.spec.whatwg.org/#shadowroot-serializable){#serialising-html-fragments:shadow-serializable
          x-internal="shadow-serializable"} is true; or

        - `shadowRoots`{.variable} contains `shadow`{.variable},

        then:

        1.  Append \"`<template shadowrootmode="`\".

        2.  If `shadow`{.variable}\'s
            [mode](https://dom.spec.whatwg.org/#shadowroot-mode){#serialising-html-fragments:concept-shadow-root-mode
            x-internal="concept-shadow-root-mode"} is \"`open`\", then
            append \"`open`\". Otherwise, append \"`closed`\".

        3.  Append \"`"`\".

        4.  If `shadow`{.variable}\'s [delegates
            focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus){#serialising-html-fragments:delegates-focus
            x-internal="delegates-focus"} is set, then append
            \"` shadowrootdelegatesfocus=""`\".

        5.  If `shadow`{.variable}\'s
            [serializable](structured-data.html#serializable){#serialising-html-fragments:serializable}
            is set, then append \"` shadowrootserializable=""`\".

        6.  If `shadow`{.variable}\'s
            [clonable](https://dom.spec.whatwg.org/#shadowroot-clonable){#serialising-html-fragments:clonable
            x-internal="clonable"} is set, then append
            \"` shadowrootclonable=""`\".

        7.  If `current node`{.variable}\'s [custom element
            registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#serialising-html-fragments:element-custom-element-registry
            x-internal="element-custom-element-registry"} is not
            `shadow`{.variable}\'s [custom element
            registry](https://dom.spec.whatwg.org/#shadowroot-custom-element-registry){#serialising-html-fragments:shadow-root-custom-element-registry
            x-internal="shadow-root-custom-element-registry"}, then
            append \"` shadowrootcustomelementregistry=""`\".

        8.  Append \"`>`\".

        9.  Append the value of running the [HTML fragment serialization
            algorithm](#html-fragment-serialisation-algorithm){#serialising-html-fragments:html-fragment-serialisation-algorithm}
            with `shadow`{.variable},
            `serializableShadowRoots`{.variable}, and
            `shadowRoots`{.variable} (thus recursing into this algorithm
            for that element).

        10. Append \"`</template>`\".

5.  For each child node of `the node`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#serialising-html-fragments:tree-order
    x-internal="tree-order"}, run the following steps:

    1.  Let `current node`{.variable} be the child node being processed.

    2.  Append the appropriate string from the following list to
        `s`{.variable}:

        If `current node`{.variable} is an [`Element`{#serialising-html-fragments:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}

        :   If `current node`{.variable} is an element in the [HTML
            namespace](https://infra.spec.whatwg.org/#html-namespace){#serialising-html-fragments:html-namespace-2
            x-internal="html-namespace-2"}, the [MathML
            namespace](https://infra.spec.whatwg.org/#mathml-namespace){#serialising-html-fragments:mathml-namespace
            x-internal="mathml-namespace"}, or the [SVG
            namespace](https://infra.spec.whatwg.org/#svg-namespace){#serialising-html-fragments:svg-namespace
            x-internal="svg-namespace"}, then let `tagname`{.variable}
            be `current node`{.variable}\'s local name. Otherwise, let
            `tagname`{.variable} be `current node`{.variable}\'s
            qualified name.

            Append a U+003C LESS-THAN SIGN character (\<), followed by
            `tagname`{.variable}.

            For [HTML
            elements](infrastructure.html#html-elements){#serialising-html-fragments:html-elements}
            created by the [HTML
            parser](#html-parser){#serialising-html-fragments:html-parser}
            or
            [`createElement()`{#serialising-html-fragments:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"},
            `tagname`{.variable} will be lowercase.

            If `current node`{.variable}\'s [`is`
            value](https://dom.spec.whatwg.org/#concept-element-is-value){#serialising-html-fragments:concept-element-is-value
            x-internal="concept-element-is-value"} is not null, and the
            element does not have an
            [`is`{#serialising-html-fragments:attr-is}](custom-elements.html#attr-is)
            attribute in its attribute list, then append the string
            \"` is="`\", followed by `current node`{.variable}\'s [`is`
            value](https://dom.spec.whatwg.org/#concept-element-is-value){#serialising-html-fragments:concept-element-is-value-2
            x-internal="concept-element-is-value"} [escaped as described
            below](#escapingString){#serialising-html-fragments:escapingString}
            in *attribute mode*, followed by a U+0022 QUOTATION MARK
            character (\").

            For each attribute that the element has, append a U+0020
            SPACE character, the [attribute\'s serialized name as
            described
            below](#attribute's-serialised-name){#serialising-html-fragments:attribute's-serialised-name},
            a U+003D EQUALS SIGN character (=), a U+0022 QUOTATION MARK
            character (\"), the attribute\'s value, [escaped as
            described
            below](#escapingString){#serialising-html-fragments:escapingString-2}
            in *attribute mode*, and a second U+0022 QUOTATION MARK
            character (\").

            An [attribute\'s serialized
            name]{#attribute's-serialised-name .dfn} for the purposes of
            the previous paragraph must be determined as follows:

            If the attribute has no namespace

            :   The attribute\'s serialized name is the attribute\'s
                local name.

                For attributes on [HTML
                elements](infrastructure.html#html-elements){#serialising-html-fragments:html-elements-2}
                set by the [HTML
                parser](#html-parser){#serialising-html-fragments:html-parser-2}
                or by
                [`setAttribute()`{#serialising-html-fragments:dom-element-setattribute}](https://dom.spec.whatwg.org/#dom-element-setattribute){x-internal="dom-element-setattribute"},
                the local name will be lowercase.

            If the attribute is in the [XML namespace](https://infra.spec.whatwg.org/#xml-namespace){#serialising-html-fragments:xml-namespace x-internal="xml-namespace"}
            :   The attribute\'s serialized name is the string
                \"`xml:`\" followed by the attribute\'s local name.

            If the attribute is in the [XMLNS namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#serialising-html-fragments:xmlns-namespace x-internal="xmlns-namespace"} and the attribute\'s local name is `xmlns`
            :   The attribute\'s serialized name is the string
                \"`xmlns`\".

            If the attribute is in the [XMLNS namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#serialising-html-fragments:xmlns-namespace-2 x-internal="xmlns-namespace"} and the attribute\'s local name is not `xmlns`
            :   The attribute\'s serialized name is the string
                \"`xmlns:`\" followed by the attribute\'s local name.

            If the attribute is in the [XLink namespace](https://infra.spec.whatwg.org/#xlink-namespace){#serialising-html-fragments:xlink-namespace x-internal="xlink-namespace"}
            :   The attribute\'s serialized name is the string
                \"`xlink:`\" followed by the attribute\'s local name.

            If the attribute is in some other namespace

            :   The attribute\'s serialized name is the attribute\'s
                qualified name.

            While the exact order of attributes is
            [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#serialising-html-fragments:implementation-defined
            x-internal="implementation-defined"}, and may depend on
            factors such as the order that the attributes were given in
            the original markup, the sort order must be stable, such
            that consecutive invocations of this algorithm serialize an
            element\'s attributes in the same order.

            Append a U+003E GREATER-THAN SIGN character (\>).

            If `current node`{.variable} [serializes as
            void](#serializes-as-void){#serialising-html-fragments:serializes-as-void-2},
            then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#serialising-html-fragments:continue
            x-internal="continue"} on to the next child node at this
            point.

            Append the value of running the [HTML fragment serialization
            algorithm](#html-fragment-serialisation-algorithm){#serialising-html-fragments:html-fragment-serialisation-algorithm-2}
            with `current node`{.variable},
            `serializableShadowRoots`{.variable}, and
            `shadowRoots`{.variable} (thus recursing into this algorithm
            for that node), followed by a U+003C LESS-THAN SIGN
            character (\<), a U+002F SOLIDUS character (/),
            `tagname`{.variable} again, and finally a U+003E
            GREATER-THAN SIGN character (\>).

        If `current node`{.variable} is a [`Text`{#serialising-html-fragments:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"} node

        :   If the parent of `current node`{.variable} is a
            [`style`{#serialising-html-fragments:the-style-element}](semantics.html#the-style-element),
            [`script`{#serialising-html-fragments:the-script-element}](scripting.html#the-script-element),
            [`xmp`{#serialising-html-fragments:xmp}](obsolete.html#xmp),
            [`iframe`{#serialising-html-fragments:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
            [`noembed`{#serialising-html-fragments:noembed}](obsolete.html#noembed),
            [`noframes`{#serialising-html-fragments:noframes}](obsolete.html#noframes),
            or
            [`plaintext`{#serialising-html-fragments:plaintext}](obsolete.html#plaintext)
            element, or if the parent of `current node`{.variable} is a
            [`noscript`{#serialising-html-fragments:the-noscript-element}](scripting.html#the-noscript-element)
            element and [scripting is
            enabled](webappapis.html#concept-n-script){#serialising-html-fragments:concept-n-script}
            for the node, then append the value of
            `current node`{.variable}\'s
            [data](https://dom.spec.whatwg.org/#concept-cd-data){#serialising-html-fragments:concept-cd-data
            x-internal="concept-cd-data"} literally.

            Otherwise, append the value of `current node`{.variable}\'s
            [data](https://dom.spec.whatwg.org/#concept-cd-data){#serialising-html-fragments:concept-cd-data-2
            x-internal="concept-cd-data"}, [escaped as described
            below](#escapingString){#serialising-html-fragments:escapingString-3}.

        If `current node`{.variable} is a [`Comment`{#serialising-html-fragments:comment-2}](https://dom.spec.whatwg.org/#interface-comment){x-internal="comment-2"}

        :   Append \"`<!--`\" (U+003C LESS-THAN SIGN, U+0021 EXCLAMATION
            MARK, U+002D HYPHEN-MINUS, U+002D HYPHEN-MINUS), followed by
            the value of `current node`{.variable}\'s
            [data](https://dom.spec.whatwg.org/#concept-cd-data){#serialising-html-fragments:concept-cd-data-3
            x-internal="concept-cd-data"}, followed by the literal
            string \"`-->`\" (U+002D HYPHEN-MINUS, U+002D HYPHEN-MINUS,
            U+003E GREATER-THAN SIGN).

        If `current node`{.variable} is a [`ProcessingInstruction`{#serialising-html-fragments:processinginstruction}](https://dom.spec.whatwg.org/#interface-processinginstruction){x-internal="processinginstruction"}

        :   Append \"`<?`\" (U+003C LESS-THAN SIGN, U+003F QUESTION
            MARK), followed by the value of `current node`{.variable}\'s
            `target` IDL attribute, followed by a single U+0020 SPACE
            character, followed by the value of
            `current node`{.variable}\'s
            [data](https://dom.spec.whatwg.org/#concept-cd-data){#serialising-html-fragments:concept-cd-data-4
            x-internal="concept-cd-data"}, followed by a single U+003E
            GREATER-THAN SIGN character (\>).

        If `current node`{.variable} is a [`DocumentType`{#serialising-html-fragments:documenttype}](https://dom.spec.whatwg.org/#interface-documenttype){x-internal="documenttype"}

        :   Append \"`<!DOCTYPE`\" (U+003C LESS-THAN SIGN, U+0021
            EXCLAMATION MARK, U+0044 LATIN CAPITAL LETTER D, U+004F
            LATIN CAPITAL LETTER O, U+0043 LATIN CAPITAL LETTER C,
            U+0054 LATIN CAPITAL LETTER T, U+0059 LATIN CAPITAL LETTER
            Y, U+0050 LATIN CAPITAL LETTER P, U+0045 LATIN CAPITAL
            LETTER E), followed by a space (U+0020 SPACE), followed by
            the value of `current node`{.variable}\'s
            [name](https://dom.spec.whatwg.org/#concept-doctype-name){#serialising-html-fragments:concept-doctype-name
            x-internal="concept-doctype-name"}, followed by \"`>`\"
            (U+003E GREATER-THAN SIGN).

6.  Return `s`{.variable}.

It is possible that the output of this algorithm, if parsed with an
[HTML parser](#html-parser){#serialising-html-fragments:html-parser-3},
will not return the original tree structure. Tree structures that do not
roundtrip a serialize and reparse step can also be produced by the [HTML
parser](#html-parser){#serialising-html-fragments:html-parser-4} itself,
although such cases are typically non-conforming.

::: example
For instance, if a
[`textarea`{#serialising-html-fragments:the-textarea-element}](form-elements.html#the-textarea-element)
element to which a `Comment` node has been appended is serialized and
the output is then reparsed, the comment will end up being displayed in
the text control. Similarly, if, as a result of DOM manipulation, an
element contains a comment that contains \"`-->`\", then when the result
of serializing the element is parsed, the comment will be truncated at
that point and the rest of the comment will be interpreted as markup.
More examples would be making a
[`script`{#serialising-html-fragments:the-script-element-2}](scripting.html#the-script-element)
element contain a
[`Text`{#serialising-html-fragments:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node with the text string \"`</script>`\", or having a
[`p`{#serialising-html-fragments:the-p-element}](grouping-content.html#the-p-element)
element that contains a
[`ul`{#serialising-html-fragments:the-ul-element}](grouping-content.html#the-ul-element)
element (as the
[`ul`{#serialising-html-fragments:the-ul-element-2}](grouping-content.html#the-ul-element)
element\'s [start
tag](syntax.html#syntax-start-tag){#serialising-html-fragments:syntax-start-tag}
would imply the end tag for the
[`p`{#serialising-html-fragments:the-p-element-2}](grouping-content.html#the-p-element)).

This can enable cross-site scripting attacks. An example of this would
be a page that lets the user enter some font family names that are then
inserted into a CSS
[`style`{#serialising-html-fragments:the-style-element-2}](semantics.html#the-style-element)
block via the DOM and which then uses the
[`innerHTML`{#serialising-html-fragments:dom-element-innerhtml}](dynamic-markup-insertion.html#dom-element-innerhtml)
IDL attribute to get the HTML serialization of that
[`style`{#serialising-html-fragments:the-style-element-3}](semantics.html#the-style-element)
element: if the user enters \"`</style><script>attack</script>`\" as a
font family name,
[`innerHTML`{#serialising-html-fragments:dom-element-innerhtml-2}](dynamic-markup-insertion.html#dom-element-innerhtml)
will return markup that, if parsed in a different context, would contain
a
[`script`{#serialising-html-fragments:the-script-element-3}](scripting.html#the-script-element)
node, even though no
[`script`{#serialising-html-fragments:the-script-element-4}](scripting.html#the-script-element)
node existed in the original DOM.
:::

::: example
For example, consider the following markup:

``` html
<form id="outer"><div></form><form id="inner"><input>
```

This will be parsed into:

- [`html`{#serialising-html-fragments:the-html-element}](semantics.html#the-html-element)
  - [`head`{#serialising-html-fragments:the-head-element}](semantics.html#the-head-element)
  - [`body`{#serialising-html-fragments:the-body-element}](sections.html#the-body-element)
    - [`form`{#serialising-html-fragments:the-form-element}](forms.html#the-form-element)
      [[`id`{#serialising-html-fragments:the-id-attribute .attribute
      .name}](dom.html#the-id-attribute)=\"`outer`{.attribute
      .value}\"]{.t2}
      - [`div`{#serialising-html-fragments:the-div-element}](grouping-content.html#the-div-element)
        - [`form`{#serialising-html-fragments:the-form-element-2}](forms.html#the-form-element)
          [[`id`{#serialising-html-fragments:the-id-attribute-2
          .attribute
          .name}](dom.html#the-id-attribute)=\"`inner`{.attribute
          .value}\"]{.t2}
          - [`input`{#serialising-html-fragments:the-input-element}](input.html#the-input-element)

The
[`input`{#serialising-html-fragments:the-input-element-2}](input.html#the-input-element)
element will be associated with the inner
[`form`{#serialising-html-fragments:the-form-element-3}](forms.html#the-form-element)
element. Now, if this tree structure is serialized and reparsed, the
`<form id="inner">` start tag will be ignored, and so the
[`input`{#serialising-html-fragments:the-input-element-3}](input.html#the-input-element)
element will be associated with the outer
[`form`{#serialising-html-fragments:the-form-element-4}](forms.html#the-form-element)
element instead.

``` html
<html><head></head><body><form id="outer"><div><form id="inner"><input></form></div></form></body></html>
```

- [`html`{#serialising-html-fragments:the-html-element-2}](semantics.html#the-html-element)
  - [`head`{#serialising-html-fragments:the-head-element-2}](semantics.html#the-head-element)
  - [`body`{#serialising-html-fragments:the-body-element-2}](sections.html#the-body-element)
    - [`form`{#serialising-html-fragments:the-form-element-5}](forms.html#the-form-element)
      [[`id`{#serialising-html-fragments:the-id-attribute-3 .attribute
      .name}](dom.html#the-id-attribute)=\"`outer`{.attribute
      .value}\"]{.t2}
      - [`div`{#serialising-html-fragments:the-div-element-2}](grouping-content.html#the-div-element)
        - [`input`{#serialising-html-fragments:the-input-element-4}](input.html#the-input-element)
:::

::: example
As another example, consider the following markup:

``` html
<a><table><a>
```

This will be parsed into:

- [`html`{#serialising-html-fragments:the-html-element-3}](semantics.html#the-html-element)
  - [`head`{#serialising-html-fragments:the-head-element-3}](semantics.html#the-head-element)
  - [`body`{#serialising-html-fragments:the-body-element-3}](sections.html#the-body-element)
    - [`a`{#serialising-html-fragments:the-a-element}](text-level-semantics.html#the-a-element)
      - [`a`{#serialising-html-fragments:the-a-element-2}](text-level-semantics.html#the-a-element)
      - [`table`{#serialising-html-fragments:the-table-element}](tables.html#the-table-element)

That is, the
[`a`{#serialising-html-fragments:the-a-element-3}](text-level-semantics.html#the-a-element)
elements are nested, because the second
[`a`{#serialising-html-fragments:the-a-element-4}](text-level-semantics.html#the-a-element)
element is [foster
parented](#foster-parent){#serialising-html-fragments:foster-parent}.
After a serialize-reparse roundtrip, the
[`a`{#serialising-html-fragments:the-a-element-5}](text-level-semantics.html#the-a-element)
elements and the
[`table`{#serialising-html-fragments:the-table-element-2}](tables.html#the-table-element)
element would all be siblings, because the second `<a>` start tag
implicitly closes the first
[`a`{#serialising-html-fragments:the-a-element-6}](text-level-semantics.html#the-a-element)
element.

``` html
<html><head></head><body><a><a></a><table></table></a></body></html>
```

- [`html`{#serialising-html-fragments:the-html-element-4}](semantics.html#the-html-element)
  - [`head`{#serialising-html-fragments:the-head-element-4}](semantics.html#the-head-element)
  - [`body`{#serialising-html-fragments:the-body-element-4}](sections.html#the-body-element)
    - [`a`{#serialising-html-fragments:the-a-element-7}](text-level-semantics.html#the-a-element)
    - [`a`{#serialising-html-fragments:the-a-element-8}](text-level-semantics.html#the-a-element)
    - [`table`{#serialising-html-fragments:the-table-element-3}](tables.html#the-table-element)
:::

For historical reasons, this algorithm does not round-trip an initial
U+000A LINE FEED (LF) character in
[`pre`{#serialising-html-fragments:the-pre-element}](grouping-content.html#the-pre-element),
[`textarea`{#serialising-html-fragments:the-textarea-element-2}](form-elements.html#the-textarea-element),
or
[`listing`{#serialising-html-fragments:listing}](obsolete.html#listing)
elements, even though (in the first two cases) the markup being
round-tripped can be conforming. The [HTML
parser](#html-parser){#serialising-html-fragments:html-parser-5} will
drop such a character during parsing, but this algorithm does *not*
serialize an extra U+000A LINE FEED (LF) character.

::: example
For example, consider the following markup:

``` html
<pre>

Hello.</pre>
```

When this document is first parsed, the
[`pre`{#serialising-html-fragments:the-pre-element-2}](grouping-content.html#the-pre-element)
element\'s [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#serialising-html-fragments:child-text-content
x-internal="child-text-content"} starts with a single newline character.
After a serialize-reparse roundtrip, the
[`pre`{#serialising-html-fragments:the-pre-element-3}](grouping-content.html#the-pre-element)
element\'s [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#serialising-html-fragments:child-text-content-2
x-internal="child-text-content"} is simply \"`Hello.`\".
:::

Because of the special role of the
[`is`{#serialising-html-fragments:attr-is-2}](custom-elements.html#attr-is)
attribute in signaling the creation of [customized built-in
elements](custom-elements.html#customized-built-in-element){#serialising-html-fragments:customized-built-in-element},
in that it provides a mechanism for parsed HTML to set the element\'s
[`is`
value](https://dom.spec.whatwg.org/#concept-element-is-value){#serialising-html-fragments:concept-element-is-value-3
x-internal="concept-element-is-value"}, we special-case its handling
during serialization. This ensures that an element\'s [`is`
value](https://dom.spec.whatwg.org/#concept-element-is-value){#serialising-html-fragments:concept-element-is-value-4
x-internal="concept-element-is-value"} is preserved through
serialize-parse roundtrips.

::: example
When creating a [customized built-in
element](custom-elements.html#customized-built-in-element){#serialising-html-fragments:customized-built-in-element-2}
via the parser, a developer uses the
[`is`{#serialising-html-fragments:attr-is-3}](custom-elements.html#attr-is)
attribute directly; in such cases serialize-parse roundtrips work fine.

``` html
<script>
window.SuperP = class extends HTMLParagraphElement {};
customElements.define("super-p", SuperP, { extends: "p" });
</script>

<div id="container"><p is="super-p">Superb!</p></div>

<script>
console.log(container.innerHTML); // <p is="super-p">
container.innerHTML = container.innerHTML;
console.log(container.innerHTML); // <p is="super-p">
console.assert(container.firstChild instanceof SuperP);
</script>
```

But when creating a customized built-in element via its
[constructor](custom-elements.html#custom-element-constructor){#serialising-html-fragments:custom-element-constructor}
or via
[`createElement()`{#serialising-html-fragments:dom-document-createelement-2}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"},
the
[`is`{#serialising-html-fragments:attr-is-4}](custom-elements.html#attr-is)
attribute is not added. Instead, the [`is`
value](https://dom.spec.whatwg.org/#concept-element-is-value){#serialising-html-fragments:concept-element-is-value-5
x-internal="concept-element-is-value"} (which is what the custom
elements machinery uses) is set without intermediating through an
attribute.

``` html
<script>
container.innerHTML = "";
const p = document.createElement("p", { is: "super-p" });
container.appendChild(p);

// The is attribute is not present in the DOM:
console.assert(!p.hasAttribute("is"));

// But the element is still a super-p:
console.assert(p instanceof SuperP);
</script>
```

To ensure that serialize-parse roundtrips still work, the serialization
process explicitly writes out the element\'s [`is`
value](https://dom.spec.whatwg.org/#concept-element-is-value){#serialising-html-fragments:concept-element-is-value-6
x-internal="concept-element-is-value"} as an
[`is`{#serialising-html-fragments:attr-is-5}](custom-elements.html#attr-is)
attribute:

``` html
<script>
console.log(container.innerHTML); // <p is="super-p">
container.innerHTML = container.innerHTML;
console.log(container.innerHTML); // <p is="super-p">
console.assert(container.firstChild instanceof SuperP);
</script>
```
:::

[Escaping a string]{#escapingString .dfn} (for the purposes of the
algorithm above) consists of running the following steps:

1.  Replace any occurrence of the \"`&`\" character by the string
    \"`&amp;`\".

2.  Replace any occurrences of the U+00A0 NO-BREAK SPACE character by
    the string \"`&nbsp;`\".

3.  Replace any occurrences of the \"`<`\" character by the string
    \"`&lt;`\".

4.  Replace any occurrences of the \"`>`\" character by the string
    \"`&gt;`\".

5.  If the algorithm was invoked in the *attribute mode*, then replace
    any occurrences of the \"`"`\" character by the string \"`&quot;`\".

### [13.4]{.secno} Parsing HTML fragments[](#parsing-html-fragments){.self-link}

The [HTML fragment parsing algorithm]{#html-fragment-parsing-algorithm
.dfn}, given an
[`Element`{#parsing-html-fragments:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
node [`context`{.variable}]{#concept-frag-parse-context .dfn}, string
`input`{.variable}, and an optional boolean
`allowDeclarativeShadowRoots`{.variable} (default false) is the
following steps. They return a list of zero or more nodes.

Parts marked [fragment case]{#fragment-case .dfn} in algorithms in the
[HTML parser](#html-parser){#parsing-html-fragments:html-parser} section
are parts that only occur if the parser was created for the purposes of
this algorithm. The algorithms have been annotated with such markings
for informational purposes only; such markings have no normative weight.
If it is possible for a condition described as a [fragment
case](#fragment-case){#parsing-html-fragments:fragment-case} to occur
even when the parser wasn\'t created for the purposes of handling this
algorithm, then that is an error in the specification.

1.  Let `document`{.variable} be a
    [`Document`{#parsing-html-fragments:document}](dom.html#document)
    node whose
    [type](https://dom.spec.whatwg.org/#concept-document-type){#parsing-html-fragments:concept-document-type
    x-internal="concept-document-type"} is \"`html`\".

2.  If [`context`{#parsing-html-fragments:concept-frag-parse-context
    .variable}](#concept-frag-parse-context)\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#parsing-html-fragments:node-document
    x-internal="node-document"} is in [quirks
    mode](https://dom.spec.whatwg.org/#concept-document-quirks){#parsing-html-fragments:quirks-mode
    x-internal="quirks-mode"}, then set `document`{.variable}\'s
    [mode](https://dom.spec.whatwg.org/#concept-document-mode){#parsing-html-fragments:concept-document-mode
    x-internal="concept-document-mode"} to \"`quirks`\".

3.  Otherwise, if
    [`context`{#parsing-html-fragments:concept-frag-parse-context-2
    .variable}](#concept-frag-parse-context)\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#parsing-html-fragments:node-document-2
    x-internal="node-document"} is in [limited-quirks
    mode](https://dom.spec.whatwg.org/#concept-document-limited-quirks){#parsing-html-fragments:limited-quirks-mode
    x-internal="limited-quirks-mode"}, then set `document`{.variable}\'s
    [mode](https://dom.spec.whatwg.org/#concept-document-mode){#parsing-html-fragments:concept-document-mode-2
    x-internal="concept-document-mode"} to \"`limited-quirks`\".

4.  If `allowDeclarativeShadowRoots`{.variable} is true, then set
    `document`{.variable}\'s [allow declarative shadow
    roots](https://dom.spec.whatwg.org/#concept-document-allow-declarative-shadow-roots){#parsing-html-fragments:concept-document-allow-declarative-shadow-roots
    x-internal="concept-document-allow-declarative-shadow-roots"} to
    true.

5.  Create a new [HTML
    parser](#html-parser){#parsing-html-fragments:html-parser-2}, and
    associate it with `document`{.variable}.

6.  Set the state of the [HTML
    parser](#html-parser){#parsing-html-fragments:html-parser-3}\'s
    [tokenization](#tokenization){#parsing-html-fragments:tokenization}
    stage as follows, switching on the
    [`context`{#parsing-html-fragments:concept-frag-parse-context-3
    .variable}](#concept-frag-parse-context) element:

    [`title`{#parsing-html-fragments:the-title-element}](semantics.html#the-title-element)\
    [`textarea`{#parsing-html-fragments:the-textarea-element}](form-elements.html#the-textarea-element)
    :   Switch the tokenizer to the [RCDATA
        state](#rcdata-state){#parsing-html-fragments:rcdata-state}.

    [`style`{#parsing-html-fragments:the-style-element}](semantics.html#the-style-element)\
    [`xmp`{#parsing-html-fragments:xmp}](obsolete.html#xmp)\
    [`iframe`{#parsing-html-fragments:the-iframe-element}](iframe-embed-object.html#the-iframe-element)\
    [`noembed`{#parsing-html-fragments:noembed}](obsolete.html#noembed)\
    [`noframes`{#parsing-html-fragments:noframes}](obsolete.html#noframes)
    :   Switch the tokenizer to the [RAWTEXT
        state](#rawtext-state){#parsing-html-fragments:rawtext-state}.

    [`script`{#parsing-html-fragments:the-script-element}](scripting.html#the-script-element)
    :   Switch the tokenizer to the [script data
        state](#script-data-state){#parsing-html-fragments:script-data-state}.

    [`noscript`{#parsing-html-fragments:the-noscript-element}](scripting.html#the-noscript-element)
    :   If the [scripting
        flag](#scripting-flag){#parsing-html-fragments:scripting-flag}
        is enabled, switch the tokenizer to the [RAWTEXT
        state](#rawtext-state){#parsing-html-fragments:rawtext-state-2}.
        Otherwise, leave the tokenizer in the [data
        state](#data-state){#parsing-html-fragments:data-state}.

    [`plaintext`{#parsing-html-fragments:plaintext}](obsolete.html#plaintext)
    :   Switch the tokenizer to the [PLAINTEXT
        state](#plaintext-state){#parsing-html-fragments:plaintext-state}.

    Any other element
    :   Leave the tokenizer in the [data
        state](#data-state){#parsing-html-fragments:data-state-2}.

    For performance reasons, an implementation that does not report
    errors and that uses the actual state machine described in this
    specification directly could use the PLAINTEXT state instead of the
    RAWTEXT and script data states where those are mentioned in the list
    above. Except for rules regarding parse errors, they are equivalent,
    since there is no [appropriate end tag
    token](#appropriate-end-tag-token){#parsing-html-fragments:appropriate-end-tag-token}
    in the fragment case, yet they involve far fewer state transitions.

7.  Let `root`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#parsing-html-fragments:create-an-element
    x-internal="create-an-element"} given `document`{.variable},
    \"`html`\", the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#parsing-html-fragments:html-namespace-2
    x-internal="html-namespace-2"}, null, null, false, and
    `context`{.variable}\'s [custom element
    registry](https://dom.spec.whatwg.org/#element-custom-element-registry){#parsing-html-fragments:element-custom-element-registry
    x-internal="element-custom-element-registry"}.

8.  [Append](https://dom.spec.whatwg.org/#concept-node-append){#parsing-html-fragments:concept-node-append
    x-internal="concept-node-append"} `root`{.variable} to
    `document`{.variable}.

9.  Set up the [HTML
    parser](#html-parser){#parsing-html-fragments:html-parser-4}\'s
    [stack of open
    elements](#stack-of-open-elements){#parsing-html-fragments:stack-of-open-elements}
    so that it contains just the single element `root`{.variable}.

10. If [`context`{#parsing-html-fragments:concept-frag-parse-context-4
    .variable}](#concept-frag-parse-context) is a
    [`template`{#parsing-html-fragments:the-template-element}](scripting.html#the-template-element)
    element, then push \"[in
    template](#parsing-main-intemplate){#parsing-html-fragments:parsing-main-intemplate}\"
    onto the [stack of template insertion
    modes](#stack-of-template-insertion-modes){#parsing-html-fragments:stack-of-template-insertion-modes}
    so that it is the new [current template insertion
    mode](#current-template-insertion-mode){#parsing-html-fragments:current-template-insertion-mode}.

11. Create a start tag token whose name is the local name of
    [`context`{#parsing-html-fragments:concept-frag-parse-context-5
    .variable}](#concept-frag-parse-context) and whose attributes are
    the attributes of
    [`context`{#parsing-html-fragments:concept-frag-parse-context-6
    .variable}](#concept-frag-parse-context).

    Let this start tag token be the start tag token of
    [`context`{#parsing-html-fragments:concept-frag-parse-context-7
    .variable}](#concept-frag-parse-context); e.g. for the purposes of
    determining if it is an [HTML integration
    point](#html-integration-point){#parsing-html-fragments:html-integration-point}.

12. [Reset the parser\'s insertion mode
    appropriately](#reset-the-insertion-mode-appropriately){#parsing-html-fragments:reset-the-insertion-mode-appropriately}.

    The parser will reference the
    [`context`{#parsing-html-fragments:concept-frag-parse-context-8
    .variable}](#concept-frag-parse-context) element as part of that
    algorithm.

13. Set the [HTML
    parser](#html-parser){#parsing-html-fragments:html-parser-5}\'s
    [`form` element
    pointer](#form-element-pointer){#parsing-html-fragments:form-element-pointer}
    to the nearest node to
    [`context`{#parsing-html-fragments:concept-frag-parse-context-9
    .variable}](#concept-frag-parse-context) that is a
    [`form`{#parsing-html-fragments:the-form-element}](forms.html#the-form-element)
    element (going straight up the ancestor chain, and including the
    element itself, if it is a
    [`form`{#parsing-html-fragments:the-form-element-2}](forms.html#the-form-element)
    element), if any. (If there is no such
    [`form`{#parsing-html-fragments:the-form-element-3}](forms.html#the-form-element)
    element, the [`form` element
    pointer](#form-element-pointer){#parsing-html-fragments:form-element-pointer-2}
    keeps its initial value, null.)

14. Place the `input`{.variable} into the [input
    stream](#input-stream){#parsing-html-fragments:input-stream} for the
    [HTML parser](#html-parser){#parsing-html-fragments:html-parser-6}
    just created. The encoding
    [confidence](#concept-encoding-confidence){#parsing-html-fragments:concept-encoding-confidence}
    is *irrelevant*.

15. Start the [HTML
    parser](#html-parser){#parsing-html-fragments:html-parser-7} and let
    it run until it has consumed all the characters just inserted into
    the input stream.

16. Return `root`{.variable}\'s
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#parsing-html-fragments:concept-tree-child
    x-internal="concept-tree-child"}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#parsing-html-fragments:tree-order
    x-internal="tree-order"}.

[← 13 The HTML syntax](syntax.html) --- [Table of Contents](./) ---
[13.5 Named character references →](named-characters.html)
