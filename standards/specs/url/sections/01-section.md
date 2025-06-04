## [1. ]{.secno}[Infrastructure]{.content}[](#infrastructure){.self-link} {#infrastructure .heading .settled level="1"}

This specification depends on Infra.
[\[INFRA\]](#biblio-infra "Infra Standard"){link-type="biblio"}

Some terms used in this specification are defined in the following
standards and specifications:

- Encoding
  [\[ENCODING\]](#biblio-encoding "Encoding Standard"){link-type="biblio"}
- File API [\[FILEAPI\]](#biblio-fileapi "File API"){link-type="biblio"}
- HTML [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
- Unicode IDNA Compatibility Processing
  [\[UTS46\]](#biblio-uts46 "Unicode IDNA Compatibility Processing"){link-type="biblio"}
- Web IDL
  [\[WEBIDL\]](#biblio-webidl "Web IDL Standard"){link-type="biblio"}

------------------------------------------------------------------------

To [serialize an integer]{#serialize-an-integer .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, represent it as the shortest possible
decimal number.

### [1.1. ]{.secno}[Writing]{.content}[](#writing){.self-link} {#writing .heading .settled level="1.1"}

A [[]{#syntax-violation .bs-old-id}validation error]{#validation-error
.dfn .dfn-paneled dfn-type="dfn" noexport=""} indicates a mismatch
between input and valid input. User agents, especially conformance
checkers, are encouraged to report them somewhere.

::: {.note role="note"}
A [validation error](#validation-error){#ref-for-validation-error
link-type="dfn"} does not mean that the parser terminates. Termination
of a parser is always stated explicitly, e.g., through a return
statement.

It is useful to signal [validation
errors](#validation-error){#ref-for-validation-error① link-type="dfn"}
as error-handling can be non-intuitive, legacy user agents might not
implement correct error-handling, and the intent of what is written
might be unclear to other developers.
:::

Error type

Error description

Failure

[IDNA](#idna)

[domain-to-ASCII]{#validation-error-domain-to-ascii .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

[Unicode
ToASCII](https://www.unicode.org/reports/tr46/#ToASCII){#ref-for-ToASCII
link-type="abstract-op"} records an error or returns the empty string.
[\[UTS46\]](#biblio-uts46 "Unicode IDNA Compatibility Processing"){link-type="biblio"}

If details about [Unicode
ToASCII](https://www.unicode.org/reports/tr46/#ToASCII){#ref-for-ToASCII①
link-type="abstract-op"} errors are recorded, user agents are encouraged
to pass those along.

Yes

[domain-invalid-code-point]{#domain-invalid-code-point .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

The input's [host](#concept-host){#ref-for-concept-host link-type="dfn"}
contains a [forbidden domain code
point](#forbidden-domain-code-point){#ref-for-forbidden-domain-code-point
link-type="dfn"}.

::: {#example-domain-invalid-code-point .example}
[](#example-domain-invalid-code-point){.self-link}

Hosts are
[percent-decoded](#string-percent-decode){#ref-for-string-percent-decode
link-type="dfn"} before being processed when the URL [is
special](#is-special){#ref-for-is-special link-type="dfn"}, which would
result in the following host portion becoming \"`exa#mple.org`\" and
thus triggering this error.

\"`https://exa%23mple.org`\"
:::

Yes

[domain-to-Unicode]{#domain-to-unicode .dfn .dfn-paneled dfn-type="dfn"
noexport=""}

[Unicode
ToUnicode](https://www.unicode.org/reports/tr46/#ToUnicode){#ref-for-ToUnicode
link-type="abstract-op"} records an error.
[\[UTS46\]](#biblio-uts46 "Unicode IDNA Compatibility Processing"){link-type="biblio"}

The same considerations as with
[domain-to-ASCII](#validation-error-domain-to-ascii){#ref-for-validation-error-domain-to-ascii
link-type="dfn"} apply.

·

[Host parsing](#host-parsing)

[host-invalid-code-point]{#host-invalid-code-point .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [opaque host](#opaque-host){#ref-for-opaque-host link-type="dfn"} (in
a URL that [is not special](#is-not-special){#ref-for-is-not-special
link-type="dfn"}) contains a [forbidden host code
point](#forbidden-host-code-point){#ref-for-forbidden-host-code-point
link-type="dfn"}.

[](#example-host-invalid-code-point){.self-link}\"`foo://exa[mple.org`\"

Yes

[IPv4-empty-part]{#ipv4-empty-part .dfn .dfn-paneled dfn-type="dfn"
noexport=""}

An [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4 link-type="dfn"}
ends with a U+002E (.).

[](#example-ipv4-empty-part){.self-link}\"`https://127.0.0.1./`\"

·

[IPv4-too-many-parts]{#ipv4-too-many-parts .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4① link-type="dfn"}
does not consist of exactly 4 parts.

[](#example-ipv4-too-many-parts){.self-link}\"`https://1.2.3.4.5/`\"

Yes

[IPv4-non-numeric-part]{#ipv4-non-numeric-part .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4② link-type="dfn"}
part is not numeric.

[](#example-ipv4-non-numeric-part){.self-link}\"`https://test.42`\"

Yes

[IPv4-non-decimal-part]{#ipv4-non-decimal-part .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

The [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4③
link-type="dfn"} contains numbers expressed using hexadecimal or octal
digits.

[](#example-ipv4-non-decimal-part){.self-link}\"`https://127.0.0x0.1`\"

·

[IPv4-out-of-range-part]{#ipv4-out-of-range-part .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4④ link-type="dfn"}
part exceeds 255.

[](#example-ipv4-out-of-range-part){.self-link}\"`https://255.255.4000.1`\"

Yes\
(only if applicable to the last part)

[IPv6-unclosed]{#ipv6-unclosed .dfn .dfn-paneled dfn-type="dfn"
noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6 link-type="dfn"}
is missing the closing U+005D (\]).

[](#example-ipv6-unclosed){.self-link}\"`https://[::1`\"

Yes

[IPv6-invalid-compression]{#ipv6-invalid-compression .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6① link-type="dfn"}
begins with improper compression.

[](#example-ipv6-invalid-compression){.self-link}\"`https://[:1]`\"

Yes

[IPv6-too-many-pieces]{#ipv6-too-many-pieces .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6② link-type="dfn"}
contains more than 8 pieces.

[](#example-ipv6-too-many-pieces){.self-link}\"`https://[1:2:3:4:5:6:7:8:9]`\"

Yes

[IPv6-multiple-compression]{#ipv6-multiple-compression .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6③ link-type="dfn"}
is compressed in more than one spot.

[](#example-ipv6-multiple-compression){.self-link}\"`https://[1::1::1]`\"

Yes

[IPv6-invalid-code-point]{#ipv6-invalid-code-point .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6④ link-type="dfn"}
contains a code point that is neither an [ASCII hex
digit](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit
link-type="dfn"} nor a U+003A (:). Or it unexpectedly ends.

::: {#example-ipv6-invalid-code-point .example}
[](#example-ipv6-invalid-code-point){.self-link}

\"`https://[1:2:3!:4]`\"

\"`https://[1:2:3:]`\"
:::

Yes

[IPv6-too-few-pieces]{#ipv6-too-few-pieces .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

An uncompressed [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6⑤
link-type="dfn"} contains fewer than 8 pieces.

[](#example-ipv6-too-few-pieces){.self-link}\"`https://[1:2:3]`\"

Yes

[IPv4-in-IPv6-too-many-pieces]{#ipv4-in-ipv6-too-many-pieces .dfn
.dfn-paneled dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6⑥ link-type="dfn"}
with [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4⑤
link-type="dfn"} syntax: the IPv6 address has more than 6 pieces.

[](#example-ipv4-in-ipv6-too-many-pieces){.self-link}\"`https://[1:1:1:1:1:1:1:127.0.0.1]`\"

Yes

[IPv4-in-IPv6-invalid-code-point]{#ipv4-in-ipv6-invalid-code-point .dfn
.dfn-paneled dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6⑦ link-type="dfn"}
with [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4⑥
link-type="dfn"} syntax:

- An IPv4 part is empty or contains a non-[ASCII
  digit](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit
  link-type="dfn"}.
- An IPv4 part contains a leading 0.
- There are too many IPv4 parts.

::: {#example-ipv4-in-ipv6-invalid-code-point .example}
[](#example-ipv4-in-ipv6-invalid-code-point){.self-link}

\"`https://[ffff::.0.0.1]`\"

\"`https://[ffff::127.0.xyz.1]`\"

\"`https://[ffff::127.0xyz]`\"

\"`https://[ffff::127.00.0.1]`\"

\"`https://[ffff::127.0.0.1.2]`\"
:::

Yes

[IPv4-in-IPv6-out-of-range-part]{#ipv4-in-ipv6-out-of-range-part .dfn
.dfn-paneled dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6⑧ link-type="dfn"}
with [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4⑦
link-type="dfn"} syntax: an IPv4 part exceeds 255.

[](#example-ipv4-in-ipv6-out-of-range-part){.self-link}\"`https://[ffff::127.0.0.4000]`\"

Yes

[IPv4-in-IPv6-too-few-parts]{#ipv4-in-ipv6-too-few-parts .dfn
.dfn-paneled dfn-type="dfn" noexport=""}

An [IPv6 address](#concept-ipv6){#ref-for-concept-ipv6⑨ link-type="dfn"}
with [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4⑧
link-type="dfn"} syntax: an IPv4 address contains too few parts.

[](#example-ipv4-in-ipv6-too-few-parts){.self-link}\"`https://[ffff::127.0.0]`\"

Yes

[URL parsing](#url-parsing)

[invalid-URL-unit]{#invalid-url-unit .dfn .dfn-paneled dfn-type="dfn"
noexport=""}

A code point is found that is not a [URL
unit](#url-units){#ref-for-url-units link-type="dfn"}.

::: {#example-invalid-url-unit .example}
[](#example-invalid-url-unit){.self-link}

\"`https://example.org/>`\"

\"` https://example.org `\"

\"`ht`\
`tps://example.org`\"

\"`https://example.org/%s`\"
:::

·

[special-scheme-missing-following-solidus]{#special-scheme-missing-following-solidus
.dfn .dfn-paneled dfn-type="dfn" noexport=""}

The input's scheme is not followed by \"`//`\".

::: {#example-special-scheme-missing-following-solidus .example}
[](#example-special-scheme-missing-following-solidus){.self-link}

\"`file:c:/my-secret-folder`\"

\"`https:example.org`\"

``` {.lang-javascript .highlight}
const url = new URL("https:foo.html", "https://example.org/");
```
:::

·

[missing-scheme-non-relative-URL]{#missing-scheme-non-relative-url .dfn
.dfn-paneled dfn-type="dfn" noexport=""}

The input is missing a
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme
link-type="dfn"}, because it does not begin with an [ASCII
alpha](https://infra.spec.whatwg.org/#ascii-alpha){#ref-for-ascii-alpha
link-type="dfn"}, and either no [base
URL](#concept-base-url){#ref-for-concept-base-url link-type="dfn"} was
provided or the [base URL](#concept-base-url){#ref-for-concept-base-url①
link-type="dfn"} cannot be used as a [base
URL](#concept-base-url){#ref-for-concept-base-url② link-type="dfn"}
because it has an [opaque
path](#url-opaque-path){#ref-for-url-opaque-path link-type="dfn"}.

::: {#example-missing-scheme-non-relative-url .example}
[](#example-missing-scheme-non-relative-url){.self-link}

Input's [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①
link-type="dfn"} is missing and no [base
URL](#concept-base-url){#ref-for-concept-base-url③ link-type="dfn"} is
given:

``` {.lang-javascript .highlight}
const url = new URL("💩");
```

Input's [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②
link-type="dfn"} is missing, but the [base
URL](#concept-base-url){#ref-for-concept-base-url④ link-type="dfn"} has
an [opaque path](#url-opaque-path){#ref-for-url-opaque-path①
link-type="dfn"}.

``` {.lang-javascript .highlight}
const url = new URL("💩", "mailto:user@example.org");
```
:::

Yes

[invalid-reverse-solidus]{#invalid-reverse-solidus .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

The URL has a [special scheme](#special-scheme){#ref-for-special-scheme
link-type="dfn"} and it uses U+005C (\\) instead of U+002F (/).

[](#example-invalid-reverse-solidus){.self-link}\"`https://example.org\path\to\file`\"

·

[invalid-credentials]{#invalid-credentials .dfn .dfn-paneled
dfn-type="dfn" noexport=""}

The input [includes
credentials](#include-credentials){#ref-for-include-credentials
link-type="dfn"}.

::: {#example-invalid-credentials .example}
[](#example-invalid-credentials){.self-link}

\"`https://user@example.org`\"

\"`ssh://user@example.org`\"
:::

·

[host-missing]{#host-missing .dfn .dfn-paneled dfn-type="dfn"
noexport=""}

The input has a [special
scheme](#special-scheme){#ref-for-special-scheme① link-type="dfn"}, but
does not contain a [host](#concept-host){#ref-for-concept-host①
link-type="dfn"}.

::: {#example-host-missing .example}
[](#example-host-missing){.self-link}

\"`https://#fragment`\"

\"`https://:443`\"

\"`https://user:pass@`\"
:::

Yes

[port-out-of-range]{#port-out-of-range .dfn .dfn-paneled dfn-type="dfn"
noexport=""}

The input's port is too big.

[](#example-port-out-of-range){.self-link}\"`https://example.org:70000`\"

Yes

[port-invalid]{#port-invalid .dfn .dfn-paneled dfn-type="dfn"
noexport=""}

The input's port is invalid.

[](#example-port-invalid){.self-link}\"`https://example.org:7z`\"

Yes

[file-invalid-Windows-drive-letter]{#file-invalid-windows-drive-letter
.dfn .dfn-paneled dfn-type="dfn" noexport=""}

The input is a [relative-URL
string](#relative-url-string){#ref-for-relative-url-string
link-type="dfn"} that [starts with a Windows drive
letter](#start-with-a-windows-drive-letter){#ref-for-start-with-a-windows-drive-letter
link-type="dfn"} and the [base
URL](#concept-base-url){#ref-for-concept-base-url⑤ link-type="dfn"}'s
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme③
link-type="dfn"} is \"`file`\".

``` {#example-file-invalid-windows-drive-letter .example}
const url = new URL("/c:/path/to/file", "file:///c:/");
```

·

[file-invalid-Windows-drive-letter-host]{#file-invalid-windows-drive-letter-host
.dfn .dfn-paneled dfn-type="dfn" noexport=""}

A `file:` URL's host is a Windows drive letter.

[](#example-file-invalid-windows-drive-letter-host){.self-link}\"`file://c:`\"

·

### [1.2. ]{.secno}[Parsers]{.content}[](#parsers){.self-link} {#parsers .heading .settled level="1.2"}

The [EOF code point]{#eof-code-point .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is a conceptual code point that signifies the end of a
string or code point stream.

A [pointer]{#pointer .dfn .dfn-paneled dfn-type="dfn" noexport=""} for a
[string](https://infra.spec.whatwg.org/#string){#ref-for-string
link-type="dfn"} `input`{.variable} is an integer that points to a [code
point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point
link-type="dfn"} within `input`{.variable}. Initially it points to the
start of `input`{.variable}. If it is −1 it points nowhere. If it is
greater than or equal to `input`{.variable}'s [code point
length](https://infra.spec.whatwg.org/#string-code-point-length){#ref-for-string-code-point-length
link-type="dfn"}, it points to the [EOF code
point](#eof-code-point){#ref-for-eof-code-point link-type="dfn"}.

When a [pointer](#pointer){#ref-for-pointer link-type="dfn"} is used,
[c]{#c .dfn .dfn-paneled dfn-type="dfn" noexport=""} references the
[code
point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point①
link-type="dfn"} the [pointer](#pointer){#ref-for-pointer①
link-type="dfn"} points to as long as it does not point nowhere. When
the [pointer](#pointer){#ref-for-pointer② link-type="dfn"} points to
nowhere [c](#c){#ref-for-c link-type="dfn"} cannot be used.

When a [pointer](#pointer){#ref-for-pointer③ link-type="dfn"} is used,
[remaining]{#remaining .dfn .dfn-paneled dfn-type="dfn" noexport=""}
references the [code point
substring](https://infra.spec.whatwg.org/#code-point-substring-to-the-end-of-the-string){#ref-for-code-point-substring-to-the-end-of-the-string
link-type="dfn"} from the [pointer](#pointer){#ref-for-pointer④
link-type="dfn"} + 1 to the end of the string, as long as
[c](#c){#ref-for-c① link-type="dfn"} is not the [EOF code
point](#eof-code-point){#ref-for-eof-code-point① link-type="dfn"}. When
[c](#c){#ref-for-c② link-type="dfn"} is the [EOF code
point](#eof-code-point){#ref-for-eof-code-point② link-type="dfn"}
[remaining](#remaining){#ref-for-remaining link-type="dfn"} cannot be
used.

[](#example-12672b6a){.self-link}If \"`mailto:username@example`\" is a
[string](https://infra.spec.whatwg.org/#string){#ref-for-string①
link-type="dfn"} being processed and a
[pointer](#pointer){#ref-for-pointer⑤ link-type="dfn"} points to @,
[c](#c){#ref-for-c③ link-type="dfn"} is U+0040 (@) and
[remaining](#remaining){#ref-for-remaining① link-type="dfn"} is
\"`example`\".

[](#example-empty-string){.self-link}If the empty string is being
processed and a [pointer](#pointer){#ref-for-pointer⑥ link-type="dfn"}
points to the start and is then decreased by 1, using
[c](#c){#ref-for-c④ link-type="dfn"} or
[remaining](#remaining){#ref-for-remaining② link-type="dfn"} would be an
error.

### [1.3. ]{.secno}[Percent-encoded bytes]{.content}[](#percent-encoded-bytes){.self-link} {#percent-encoded-bytes .heading .settled level="1.3"}

A [percent-encoded byte]{#percent-encoded-byte .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is U+0025 (%), followed by two [ASCII hex
digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit①
link-type="dfn"}.

It is generally a good idea for sequences of [percent-encoded
bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte
link-type="dfn"} to be such that, when
[percent-decoded](#string-percent-decode){#ref-for-string-percent-decode①
link-type="dfn"} and then passed to [UTF-8 decode without BOM or
fail](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom-or-fail){#ref-for-utf-8-decode-without-bom-or-fail
link-type="dfn"}, they do not end up as failure. How important this is
depends on where the [percent-encoded
bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte①
link-type="dfn"} are used. E.g., for the [host
parser](#concept-host-parser){#ref-for-concept-host-parser
link-type="dfn"} not following this advice is fatal, whereas for [URL
rendering](#url-rendering-i18n) the [percent-encoded
bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte②
link-type="dfn"} would not be rendered
[percent-decoded](#string-percent-decode){#ref-for-string-percent-decode②
link-type="dfn"}.

::: {.algorithm algorithm="percent-encode" algorithm-for="byte"}
To [percent-encode]{#percent-encode .dfn .dfn-paneled dfn-for="byte"
dfn-type="dfn" noexport=""} a
[byte](https://infra.spec.whatwg.org/#byte){#ref-for-byte
link-type="dfn"} `byte`{.variable}, return a
[string](https://infra.spec.whatwg.org/#string){#ref-for-string②
link-type="dfn"} consisting of U+0025 (%), followed by two [ASCII upper
hex
digits](https://infra.spec.whatwg.org/#ascii-upper-hex-digit){#ref-for-ascii-upper-hex-digit
link-type="dfn"} representing `byte`{.variable}.
:::

::: {.algorithm algorithm="percent-decode" algorithm-for="byte sequence"}
To [percent-decode]{#percent-decode .dfn .dfn-paneled
dfn-for="byte sequence" dfn-type="dfn" export=""} a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence
link-type="dfn"} `input`{.variable}, run these steps:

Using anything but [UTF-8 decode without
BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#ref-for-utf-8-decode-without-bom
link-type="dfn"} when `input`{.variable} contains bytes that are not
[ASCII
bytes](https://infra.spec.whatwg.org/#ascii-byte){#ref-for-ascii-byte
link-type="dfn"} might be insecure and is not recommended.

1.  Let `output`{.variable} be an empty [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①
    link-type="dfn"}.

2.  For each byte `byte`{.variable} in `input`{.variable}:

    1.  If `byte`{.variable} is not 0x25 (%), then append
        `byte`{.variable} to `output`{.variable}.

    2.  Otherwise, if `byte`{.variable} is 0x25 (%) and the next two
        bytes after `byte`{.variable} in `input`{.variable} are not in
        the ranges 0x30 (0) to 0x39 (9), 0x41 (A) to 0x46 (F), and
        0x61 (a) to 0x66 (f), all inclusive, append `byte`{.variable} to
        `output`{.variable}.

    3.  Otherwise:

        1.  Let `bytePoint`{.variable} be the two bytes after
            `byte`{.variable} in `input`{.variable},
            [decoded](https://infra.spec.whatwg.org/#isomorphic-decode){#ref-for-isomorphic-decode
            link-type="dfn"}, and then interpreted as hexadecimal
            number.

        2.  Append a byte whose value is `bytePoint`{.variable} to
            `output`{.variable}.

        3.  Skip the next two bytes in `input`{.variable}.

3.  Return `output`{.variable}.
:::

::: {.algorithm algorithm="percent-decode" algorithm-for="string"}
To [percent-decode]{#string-percent-decode .dfn .dfn-paneled
dfn-for="string" dfn-type="dfn" export=""} a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string
link-type="dfn"} `input`{.variable}:

1.  Let `bytes`{.variable} be the [UTF-8
    encoding](https://encoding.spec.whatwg.org/#utf-8-encode){#ref-for-utf-8-encode
    link-type="dfn"} of `input`{.variable}.

2.  Return the
    [percent-decoding](#percent-decode){#ref-for-percent-decode
    link-type="dfn"} of `bytes`{.variable}.

In general, percent-encoding results in a string with more U+0025 (%)
code points than the input, and percent-decoding results in a byte
sequence with less 0x25 (%) bytes than the input.
:::

------------------------------------------------------------------------

The [[]{#simple-encode-set .bs-old-id}C0 control percent-encode
set]{#c0-control-percent-encode-set .dfn .dfn-paneled dfn-type="dfn"
noexport=""} are the [C0
controls](https://infra.spec.whatwg.org/#c0-control){#ref-for-c0-control
link-type="dfn"} and all [code
points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point②
link-type="dfn"} greater than U+007E (\~).

The [fragment percent-encode set]{#fragment-percent-encode-set .dfn
.dfn-paneled dfn-type="dfn" noexport=""} is the [C0 control
percent-encode
set](#c0-control-percent-encode-set){#ref-for-c0-control-percent-encode-set
link-type="dfn"} and U+0020 SPACE, U+0022 (\"), U+003C (\<), U+003E
(\>), and U+0060 (\`).

The [query percent-encode set]{#query-percent-encode-set .dfn
.dfn-paneled dfn-type="dfn" noexport=""} is the [C0 control
percent-encode
set](#c0-control-percent-encode-set){#ref-for-c0-control-percent-encode-set①
link-type="dfn"} and U+0020 SPACE, U+0022 (\"), U+0023 (#), U+003C (\<),
and U+003E (\>).

The [query percent-encode
set](#query-percent-encode-set){#ref-for-query-percent-encode-set
link-type="dfn"} cannot be defined in terms of the [fragment
percent-encode
set](#fragment-percent-encode-set){#ref-for-fragment-percent-encode-set
link-type="dfn"} due to the omission of U+0060 (\`).

The [special-query percent-encode set]{#special-query-percent-encode-set
.dfn .dfn-paneled dfn-type="dfn" noexport=""} is the [query
percent-encode
set](#query-percent-encode-set){#ref-for-query-percent-encode-set①
link-type="dfn"} and U+0027 (\').

The [[]{#default-encode-set .bs-old-id}path percent-encode
set]{#path-percent-encode-set .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is the [query percent-encode
set](#query-percent-encode-set){#ref-for-query-percent-encode-set②
link-type="dfn"} and U+003F (?), U+005E (\^), U+0060 (\`), U+007B ({),
and U+007D (}).

The [[]{#userinfo-encode-set .bs-old-id}userinfo percent-encode
set]{#userinfo-percent-encode-set .dfn .dfn-paneled dfn-type="dfn"
noexport=""} is the [path percent-encode
set](#path-percent-encode-set){#ref-for-path-percent-encode-set
link-type="dfn"} and U+002F (/), U+003A (:), U+003B (;), U+003D (=),
U+0040 (@), U+005B (\[) to U+005D (\]), inclusive, and U+007C (\|).

The [component percent-encode set]{#component-percent-encode-set .dfn
.dfn-paneled dfn-type="dfn" export=""} is the [userinfo percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set
link-type="dfn"} and U+0024 (\$) to U+0026 (&), inclusive, U+002B (+),
and U+002C (,).

This is used by HTML for
[`registerProtocolHandler()`{.idl}](https://html.spec.whatwg.org/multipage/system-state.html#dom-navigator-registerprotocolhandler){#ref-for-dom-navigator-registerprotocolhandler
link-type="idl"}, and could also be used by other standards to
percent-encode data that can then be embedded in a
[URL](#concept-url){#ref-for-concept-url link-type="dfn"}'s
[path](#concept-url-path){#ref-for-concept-url-path link-type="dfn"},
[query](#concept-url-query){#ref-for-concept-url-query link-type="dfn"},
or [fragment](#concept-url-fragment){#ref-for-concept-url-fragment
link-type="dfn"}; or in an [opaque
host](#opaque-host){#ref-for-opaque-host① link-type="dfn"}. Using it
with [UTF-8
percent-encode](#string-utf-8-percent-encode){#ref-for-string-utf-8-percent-encode
link-type="dfn"} gives identical results to JavaScript's
[`encodeURIComponent()`
\[sic\]](https://tc39.es/ecma262/#sec-encodeuricomponent-uricomponent){#ref-for-sec-encodeuricomponent-uricomponent
.idl-code link-type="method"}.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
[\[ECMA-262\]](#biblio-ecma-262 "ECMAScript Language Specification"){link-type="biblio"}

The [`application/x-www-form-urlencoded` percent-encode
set]{#application-x-www-form-urlencoded-percent-encode-set .dfn
.dfn-paneled dfn-type="dfn" noexport=""} is the [component
percent-encode
set](#component-percent-encode-set){#ref-for-component-percent-encode-set
link-type="dfn"} and U+0021 (!), U+0027 (\') to U+0029 RIGHT
PARENTHESIS, inclusive, and U+007E (\~).

The [`application/x-www-form-urlencoded` percent-encode
set](#application-x-www-form-urlencoded-percent-encode-set){#ref-for-application-x-www-form-urlencoded-percent-encode-set
link-type="dfn"} contains all code points, except the [ASCII
alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric){#ref-for-ascii-alphanumeric
link-type="dfn"}, U+002A (\*), U+002D (-), U+002E (.), and U+005F (\_).

::: {.algorithm algorithm="percent-encode after encoding" algorithm-for="string"}
To [percent-encode after encoding]{#string-percent-encode-after-encoding
.dfn .dfn-paneled dfn-for="string" dfn-type="dfn" noexport=""}, given an
[encoding](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding
link-type="dfn"} `encoding`{.variable}, [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string①
link-type="dfn"} `input`{.variable}, a `percentEncodeSet`{.variable},
and an optional boolean `spaceAsPlus`{.variable} (default false):

1.  Let `encoder`{.variable} be the result of [getting an
    encoder](https://encoding.spec.whatwg.org/#get-an-encoder){#ref-for-get-an-encoder
    link-type="dfn"} from `encoding`{.variable}.

2.  Let `inputQueue`{.variable} be `input`{.variable} converted to an
    [I/O
    queue](https://encoding.spec.whatwg.org/#concept-stream){#ref-for-concept-stream
    link-type="dfn"}.

3.  Let `output`{.variable} be the empty string.

4.  Let `potentialError`{.variable} be 0.

    This needs to be a non-null value to initiate the subsequent while
    loop.

5.  While `potentialError`{.variable} is non-null:

    1.  Let `encodeOutput`{.variable} be an empty [I/O
        queue](https://encoding.spec.whatwg.org/#concept-stream){#ref-for-concept-stream①
        link-type="dfn"}.

    2.  Set `potentialError`{.variable} to the result of running [encode
        or
        fail](https://encoding.spec.whatwg.org/#encode-or-fail){#ref-for-encode-or-fail
        link-type="dfn"} with `inputQueue`{.variable},
        `encoder`{.variable}, and `encodeOutput`{.variable}.

    3.  For each `byte`{.variable} of `encodeOutput`{.variable}
        converted to a byte sequence:

        1.  If `spaceAsPlus`{.variable} is true and `byte`{.variable} is
            0x20 (SP), then append U+002B (+) to `output`{.variable} and
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue
            link-type="dfn"}.

        2.  Let `isomorph`{.variable} be a [code
            point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point③
            link-type="dfn"} whose
            [value](https://infra.spec.whatwg.org/#code-point-value){#ref-for-code-point-value
            link-type="dfn"} is `byte`{.variable}'s
            [value](https://infra.spec.whatwg.org/#byte-value){#ref-for-byte-value
            link-type="dfn"}.

        3.  Assert: `percentEncodeSet`{.variable} includes all
            non-[ASCII code
            points](https://infra.spec.whatwg.org/#ascii-code-point){#ref-for-ascii-code-point
            link-type="dfn"}.

        4.  If `isomorph`{.variable} is not in
            `percentEncodeSet`{.variable}, then append
            `isomorph`{.variable} to `output`{.variable}.

        5.  Otherwise,
            [percent-encode](#percent-encode){#ref-for-percent-encode
            link-type="dfn"} `byte`{.variable} and append the result to
            `output`{.variable}.

    4.  If `potentialError`{.variable} is non-null, then append
        \"`%26%23`\", followed by the shortest sequence of [ASCII
        digits](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit①
        link-type="dfn"} representing `potentialError`{.variable} in
        base ten, followed by \"`%3B`\", to `output`{.variable}.

        This can happen when `encoding`{.variable} is not
        [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8
        link-type="dfn"}.

6.  Return `output`{.variable}.

Of the possible values for the `percentEncodeSet`{.variable} argument
only two end up encoding U+0025 (%) and thus give "roundtripable data":
[component percent-encode
set](#component-percent-encode-set){#ref-for-component-percent-encode-set①
link-type="dfn"} and [`application/x-www-form-urlencoded` percent-encode
set](#application-x-www-form-urlencoded-percent-encode-set){#ref-for-application-x-www-form-urlencoded-percent-encode-set①
link-type="dfn"}. The other values for the `percentEncodeSet`{.variable}
argument --- which happen to be used by the [URL
parser](#concept-url-parser){#ref-for-concept-url-parser
link-type="dfn"} --- leave U+0025 (%) untouched and as such it needs to
be
[percent-encoded](#utf-8-percent-encode){#ref-for-utf-8-percent-encode
link-type="dfn"} first in order to be properly represented.
:::

::: {.algorithm algorithm="UTF-8 percent-encode" algorithm-for="code point"}
To [UTF-8 percent-encode]{#utf-8-percent-encode .dfn .dfn-paneled
dfn-for="code point" dfn-type="dfn" noexport=""} a [scalar
value](https://infra.spec.whatwg.org/#scalar-value){#ref-for-scalar-value
link-type="dfn"} `scalarValue`{.variable} using a
`percentEncodeSet`{.variable}, return the result of running
[percent-encode after
encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding
link-type="dfn"} with
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8①
link-type="dfn"}, `scalarValue`{.variable} as a
[string](https://infra.spec.whatwg.org/#string){#ref-for-string③
link-type="dfn"}, and `percentEncodeSet`{.variable}.
:::

::: {.algorithm algorithm="UTF-8 percent-encode" algorithm-for="string"}
To [UTF-8 percent-encode]{#string-utf-8-percent-encode .dfn .dfn-paneled
dfn-for="string" dfn-type="dfn" export=""} a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string②
link-type="dfn"} `input`{.variable} using a
`percentEncodeSet`{.variable}, return the result of running
[percent-encode after
encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding①
link-type="dfn"} with
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8②
link-type="dfn"}, `input`{.variable}, and `percentEncodeSet`{.variable}.
:::

------------------------------------------------------------------------

::: {#example-percent-encode-operations .example}
[](#example-percent-encode-operations){.self-link}

Here is a summary, by way of example, of the operations defined above:

Operation

Input

Output

[Percent-encode](#percent-encode){#ref-for-percent-encode①
link-type="dfn"} `input`{.variable}

0x23

\"`%23`\"

0x7F

\"`%7F`\"

[Percent-decode](#percent-decode){#ref-for-percent-decode①
link-type="dfn"} `input`{.variable}

\``%25%s%1G`\`

\``%%s%1G`\`

[Percent-decode](#string-percent-decode){#ref-for-string-percent-decode③
link-type="dfn"} `input`{.variable}

\"`‽%25%2E`\"

0xE2 0x80 0xBD 0x25 0x2E

[Percent-encode after
encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding②
link-type="dfn"} with
[Shift_JIS](https://encoding.spec.whatwg.org/#shift_jis){#ref-for-shift_jis
link-type="dfn"}, `input`{.variable}, and the [userinfo percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set①
link-type="dfn"}

\"` `\"

\"`%20`\"

\"`≡`\"

\"`%81%DF`\"

\"`‽`\"

\"`%26%238253%3B`\"

[Percent-encode after
encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding③
link-type="dfn"} with
[ISO-2022-JP](https://encoding.spec.whatwg.org/#iso-2022-jp){#ref-for-iso-2022-jp
link-type="dfn"}, `input`{.variable}, and the [userinfo percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set②
link-type="dfn"}

\"`¥`\"

\"`%1B(J\%1B(B`\"

[Percent-encode after
encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding④
link-type="dfn"} with
[Shift_JIS](https://encoding.spec.whatwg.org/#shift_jis){#ref-for-shift_jis①
link-type="dfn"}, `input`{.variable}, the [userinfo percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set③
link-type="dfn"}, and true

\"`1+1 ≡ 2%20‽`\"

\"`1+1+%81%DF+2%20%26%238253%3B`\"

[UTF-8
percent-encode](#utf-8-percent-encode){#ref-for-utf-8-percent-encode①
link-type="dfn"} `input`{.variable} using the [userinfo percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set④
link-type="dfn"}

U+2261 (≡)

\"`%E2%89%A1`\"

U+203D (‽)

\"`%E2%80%BD`\"

[UTF-8
percent-encode](#string-utf-8-percent-encode){#ref-for-string-utf-8-percent-encode①
link-type="dfn"} `input`{.variable} using the [userinfo percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set⑤
link-type="dfn"}

\"`Say what‽`\"

\"`Say%20what%E2%80%BD`\"
:::

