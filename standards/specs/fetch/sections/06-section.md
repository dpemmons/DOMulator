## [6. ]{.secno}[`data:` URLs]{.content}[](#data-urls){.self-link} {#data-urls .heading .settled level="6"}

For an informative description of `data:` URLs, see RFC 2397. This
section replaces that RFC's normative processing requirements to be
compatible with deployed content.
[\[RFC2397\]](#biblio-rfc2397 "The "data" URL scheme"){link-type="biblio"}

A [`data:` URL struct]{#data-url-struct .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct⑤
link-type="dfn"} that consists of a [MIME
type]{#data-url-struct-mime-type .dfn .dfn-paneled
dfn-for="data: URL struct" dfn-type="dfn" noexport=""} (a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#ref-for-mime-type⑤
link-type="dfn"}) and a [body]{#data-url-struct-body .dfn .dfn-paneled
dfn-for="data: URL struct" dfn-type="dfn" noexport=""} (a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence③⓪
link-type="dfn"}).

::: {.algorithm algorithm="data: URL processor"}
The [`data:` URL processor]{#data-url-processor .dfn .dfn-paneled
dfn-type="dfn" export=""} takes a
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url②①
link-type="dfn"} `dataURL`{.variable} and then runs these steps:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑧
    link-type="dfn"}: `dataURL`{.variable}'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①⑥
    link-type="dfn"} is \"`data`\".

2.  Let `input`{.variable} be the result of running the [URL
    serializer](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer⑥
    link-type="dfn"} on `dataURL`{.variable} with [*exclude
    fragment*](https://url.spec.whatwg.org/#url-serializer-exclude-fragment){#ref-for-url-serializer-exclude-fragment②
    link-type="dfn"} set to true.

3.  Remove the leading \"`data:`\" from `input`{.variable}.

4.  Let `position`{.variable} point at the start of `input`{.variable}.

5.  Let `mimeType`{.variable} be the result of [collecting a sequence of
    code
    points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#ref-for-collect-a-sequence-of-code-points⑧
    link-type="dfn"} that are not equal to U+002C (,), given
    `position`{.variable}.

6.  [Strip leading and trailing ASCII
    whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#ref-for-strip-leading-and-trailing-ascii-whitespace
    link-type="dfn"} from `mimeType`{.variable}.

    This will only remove U+0020 SPACE [code
    points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①⓪
    link-type="dfn"}, if any.

7.  If `position`{.variable} is past the end of `input`{.variable}, then
    return failure.

8.  Advance `position`{.variable} by 1.

9.  Let `encodedBody`{.variable} be the remainder of `input`{.variable}.

10. Let `body`{.variable} be the
    [percent-decoding](https://url.spec.whatwg.org/#string-percent-decode){#ref-for-string-percent-decode
    link-type="dfn"} of `encodedBody`{.variable}.

11. If `mimeType`{.variable} ends with U+003B (;), followed by zero or
    more U+0020 SPACE, followed by an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive①
    link-type="dfn"} match for \"`base64`\", then:

    1.  Let `stringBody`{.variable} be the [isomorphic
        decode](https://infra.spec.whatwg.org/#isomorphic-decode){#ref-for-isomorphic-decode③
        link-type="dfn"} of `body`{.variable}.

    2.  Set `body`{.variable} to the [forgiving-base64
        decode](https://infra.spec.whatwg.org/#forgiving-base64-decode){#ref-for-forgiving-base64-decode
        link-type="dfn"} of `stringBody`{.variable}.

    3.  If `body`{.variable} is failure, then return failure.

    4.  Remove the last 6 [code
        points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①①
        link-type="dfn"} from `mimeType`{.variable}.

    5.  Remove trailing U+0020 SPACE [code
        points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①②
        link-type="dfn"} from `mimeType`{.variable}, if any.

    6.  Remove the last U+003B (;) from `mimeType`{.variable}.

12. If `mimeType`{.variable} [starts
    with](https://infra.spec.whatwg.org/#string-starts-with){#ref-for-string-starts-with②
    link-type="dfn"} \"`;`\", then prepend \"`text/plain`\" to
    `mimeType`{.variable}.

13. Let `mimeTypeRecord`{.variable} be the result of
    [parsing](https://mimesniff.spec.whatwg.org/#parse-a-mime-type){#ref-for-parse-a-mime-type②
    link-type="dfn"} `mimeType`{.variable}.

14. If `mimeTypeRecord`{.variable} is failure, then set
    `mimeTypeRecord`{.variable} to `text/plain;charset=US-ASCII`.

15. Return a new [`data:` URL
    struct](#data-url-struct){#ref-for-data-url-struct link-type="dfn"}
    whose [MIME
    type](#data-url-struct-mime-type){#ref-for-data-url-struct-mime-type①
    link-type="dfn"} is `mimeTypeRecord`{.variable} and
    [body](#data-url-struct-body){#ref-for-data-url-struct-body①
    link-type="dfn"} is `body`{.variable}.
:::

