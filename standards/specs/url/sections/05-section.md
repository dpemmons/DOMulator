## [5. ]{.secno}[`application/x-www-form-urlencoded`]{.content}[](#application/x-www-form-urlencoded){.self-link} {#application/x-www-form-urlencoded .heading .settled level="5"}

The [`application/x-www-form-urlencoded`]{#concept-urlencoded .dfn
.dfn-paneled dfn-type="dfn" export=""} format provides a way to encode a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list④
link-type="dfn"} of
[tuples](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple①
link-type="dfn"}, each consisting of a name and a value.

The `application/x-www-form-urlencoded` format is in many ways an
aberrant monstrosity, the result of many years of implementation
accidents and compromises leading to a set of requirements necessary for
interoperability, but in no way representing good design practices. In
particular, readers are cautioned to pay close attention to the twisted
details involving repeated (and in some cases nested) conversions
between character encodings and byte sequences. Unfortunately the format
is in widespread use due to the prevalence of HTML forms.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

### [5.1. ]{.secno}[`application/x-www-form-urlencoded` parsing]{.content}[](#urlencoded-parsing){.self-link} {#urlencoded-parsing .heading .settled level="5.1"}

A legacy server-oriented implementation might have to support
[encodings](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding③
link-type="dfn"} other than
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8⑧
link-type="dfn"} as well as have special logic for tuples of which the
name is \``_charset`\`. Such logic is not described here as only
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8⑨
link-type="dfn"} is conforming.

::: {.algorithm algorithm="urlencoded parser"}
The [`application/x-www-form-urlencoded`
parser]{#concept-urlencoded-parser .dfn .dfn-paneled dfn-type="dfn"
export="" lt="urlencoded parser"} takes a byte sequence
`input`{.variable}, and then runs these steps:

1.  Let `sequences`{.variable} be the result of splitting
    `input`{.variable} on 0x26 (&).

2.  Let `output`{.variable} be an initially empty
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑤
    link-type="dfn"} of name-value tuples where both name and value hold
    a string.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑥
    link-type="dfn"} byte sequence `bytes`{.variable} in
    `sequences`{.variable}:

    1.  If `bytes`{.variable} is the empty byte sequence, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue⑤
        link-type="dfn"}.

    2.  If `bytes`{.variable} contains a 0x3D (=), then let
        `name`{.variable} be the bytes from the start of
        `bytes`{.variable} up to but excluding its first 0x3D (=), and
        let `value`{.variable} be the bytes, if any, after the first
        0x3D (=) up to the end of `bytes`{.variable}. If 0x3D (=) is the
        first byte, then `name`{.variable} will be the empty byte
        sequence. If it is the last, then `value`{.variable} will be the
        empty byte sequence.

    3.  Otherwise, let `name`{.variable} have the value of
        `bytes`{.variable} and let `value`{.variable} be the empty byte
        sequence.

    4.  Replace any 0x2B (+) in `name`{.variable} and `value`{.variable}
        with 0x20 (SP).

    5.  Let `nameString`{.variable} and `valueString`{.variable} be the
        result of running [UTF-8 decode without
        BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#ref-for-utf-8-decode-without-bom③
        link-type="dfn"} on the
        [percent-decoding](#percent-decode){#ref-for-percent-decode②
        link-type="dfn"} of `name`{.variable} and `value`{.variable},
        respectively.

    6.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑥
        link-type="dfn"} (`nameString`{.variable},
        `valueString`{.variable}) to `output`{.variable}.

4.  Return `output`{.variable}.
:::

### [5.2. ]{.secno}[`application/x-www-form-urlencoded` serializing]{.content}[](#urlencoded-serializing){.self-link} {#urlencoded-serializing .heading .settled level="5.2"}

::: {.algorithm algorithm="urlencoded serializer"}
The [`application/x-www-form-urlencoded`
serializer]{#concept-urlencoded-serializer .dfn .dfn-paneled
dfn-type="dfn" export="" lt="urlencoded serializer"} takes a list of
name-value tuples `tuples`{.variable}, with an optional
[encoding](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding④
link-type="dfn"} `encoding`{.variable} (default
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8①⓪
link-type="dfn"}), and then runs these steps. They return an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string②③
link-type="dfn"}.

1.  Set `encoding`{.variable} to the result of [getting an output
    encoding](https://encoding.spec.whatwg.org/#get-an-output-encoding){#ref-for-get-an-output-encoding①
    link-type="dfn"} from `encoding`{.variable}.

2.  Let `output`{.variable} be the empty string.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑦
    link-type="dfn"} `tuple`{.variable} of `tuples`{.variable}:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert③
        link-type="dfn"}: `tuple`{.variable}'s name and
        `tuple`{.variable}'s value are [scalar value
        strings](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string①⓪
        link-type="dfn"}.

    2.  Let `name`{.variable} be the result of running [percent-encode
        after
        encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding⑥
        link-type="dfn"} with `encoding`{.variable},
        `tuple`{.variable}'s name, the
        [`application/x-www-form-urlencoded` percent-encode
        set](#application-x-www-form-urlencoded-percent-encode-set){#ref-for-application-x-www-form-urlencoded-percent-encode-set②
        link-type="dfn"}, and true.

    3.  Let `value`{.variable} be the result of running [percent-encode
        after
        encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding⑦
        link-type="dfn"} with `encoding`{.variable},
        `tuple`{.variable}'s value, the
        [`application/x-www-form-urlencoded` percent-encode
        set](#application-x-www-form-urlencoded-percent-encode-set){#ref-for-application-x-www-form-urlencoded-percent-encode-set③
        link-type="dfn"}, and true.

    4.  If `output`{.variable} is not the empty string, then append
        U+0026 (&) to `output`{.variable}.

    5.  Append `name`{.variable}, followed by U+003D (=), followed by
        `value`{.variable}, to `output`{.variable}.

4.  Return `output`{.variable}.
:::

### [5.3. ]{.secno}[Hooks]{.content}[](#urlencoded-hooks){.self-link} {#urlencoded-hooks .heading .settled level="5.3"}

The [`application/x-www-form-urlencoded` string
parser]{#concept-urlencoded-string-parser .dfn .dfn-paneled
dfn-type="dfn" lt="urlencoded string parser" noexport=""} takes a
[scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string①①
link-type="dfn"} `input`{.variable}, [UTF-8
encodes](https://encoding.spec.whatwg.org/#utf-8-encode){#ref-for-utf-8-encode①
link-type="dfn"} it, and then returns the result of
[`application/x-www-form-urlencoded`
parsing](#concept-urlencoded-parser){#ref-for-concept-urlencoded-parser
link-type="dfn"} it.

