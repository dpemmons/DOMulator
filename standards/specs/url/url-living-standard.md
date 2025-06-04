## [Goals]{.content} {#goals .no-num .heading .settled}

The URL standard takes the following approach towards making URLs fully
interoperable:

- Align RFC 3986 and RFC 3987 with contemporary implementations and
  obsolete the RFCs in the process. (E.g., spaces, other \"illegal\"
  code points, query encoding, equality, canonicalization, are all
  concepts not entirely shared, or defined.) URL parsing needs to become
  as solid as HTML parsing.
  [\[RFC3986\]](#biblio-rfc3986 "Uniform Resource Identifier (URI): Generic Syntax"){link-type="biblio"}
  [\[RFC3987\]](#biblio-rfc3987 "Internationalized Resource Identifiers (IRIs)"){link-type="biblio"}

- Standardize on the term URL. URI and IRI are just confusing. In
  practice a single algorithm is used for both so keeping them distinct
  is not helping anyone. URL also easily wins the [search result
  popularity
  contest](https://trends.google.com/trends/explore?q=url,uri).

- Supplanting [Origin of a URI
  \[sic\]](https://tools.ietf.org/html/rfc6454#section-4).
  [\[RFC6454\]](#biblio-rfc6454 "The Web Origin Concept"){link-type="biblio"}

- Define URL's existing JavaScript API in full detail and add
  enhancements to make it easier to work with. Add a new
  [`URL`](#url){#ref-for-url .idl-code link-type="interface"} object as
  well for URL manipulation without usage of HTML elements. (Useful for
  JavaScript worker environments.)

- Ensure the combination of parser, serializer, and API guarantee
  idempotence. For example, a non-failure result of a
  parse-then-serialize operation will not change with any further
  parse-then-serialize operations applied to it. Similarly, manipulating
  a non-failure result through the API will not change from applying any
  number of serialize-then-parse operations to it.

As the editors learn more about the subject matter the goals might
increase in scope somewhat.

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

## [2. ]{.secno}[Security considerations]{.content}[](#security-considerations){.self-link} {#security-considerations .heading .settled level="2"}

The security of a [URL](#concept-url){#ref-for-concept-url①
link-type="dfn"} is a function of its environment. Care is to be taken
when rendering, interpreting, and passing
[URLs](#concept-url){#ref-for-concept-url② link-type="dfn"} around.

When rendering and allocating new
[URLs](#concept-url){#ref-for-concept-url③ link-type="dfn"} \"spoofing\"
needs to be considered. An attack whereby one
[host](#concept-host){#ref-for-concept-host② link-type="dfn"} or
[URL](#concept-url){#ref-for-concept-url④ link-type="dfn"} can be
confused for another. For instance, consider how 1/l/I, m/rn/rri, 0/O,
and а/a can all appear eerily similar. Or worse, consider how U+202A
LEFT-TO-RIGHT EMBEDDING and similar [code
points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point④
link-type="dfn"} are invisible.
[\[UTR36\]](#biblio-utr36 "Unicode Security Considerations"){link-type="biblio"}

When passing a [URL](#concept-url){#ref-for-concept-url⑤
link-type="dfn"} from party `A`{.variable} to `B`{.variable}, both need
to carefully consider what is happening. `A`{.variable} might end up
leaking data it does not want to leak. `B`{.variable} might receive
input it did not expect and take an action that harms the user. In
particular, `B`{.variable} should never trust `A`{.variable}, as at some
point [URLs](#concept-url){#ref-for-concept-url⑥ link-type="dfn"} from
`A`{.variable} can come from untrusted sources.

## [3. ]{.secno}[Hosts (domains and IP addresses)]{.content}[](#hosts-(domains-and-ip-addresses)){.self-link} {#hosts-(domains-and-ip-addresses) .heading .settled level="3"}

At a high level, a [host](#concept-host){#ref-for-concept-host③
link-type="dfn"}, [valid host
string](#valid-host-string){#ref-for-valid-host-string link-type="dfn"},
[host parser](#concept-host-parser){#ref-for-concept-host-parser①
link-type="dfn"}, and [host
serializer](#concept-host-serializer){#ref-for-concept-host-serializer
link-type="dfn"} relate as follows:

- The [host parser](#concept-host-parser){#ref-for-concept-host-parser②
  link-type="dfn"} takes an arbitrary [scalar value
  string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string③
  link-type="dfn"} and returns either failure or a
  [host](#concept-host){#ref-for-concept-host④ link-type="dfn"}.

- A [host](#concept-host){#ref-for-concept-host⑤ link-type="dfn"} can be
  seen as the in-memory representation.

- A [valid host string](#valid-host-string){#ref-for-valid-host-string①
  link-type="dfn"} defines what input would not trigger a [validation
  error](#validation-error){#ref-for-validation-error② link-type="dfn"}
  or failure when given to the [host
  parser](#concept-host-parser){#ref-for-concept-host-parser③
  link-type="dfn"}. I.e., input that would be considered conforming or
  valid.

- The [host
  serializer](#concept-host-serializer){#ref-for-concept-host-serializer①
  link-type="dfn"} takes a [host](#concept-host){#ref-for-concept-host⑥
  link-type="dfn"} and returns an [ASCII
  string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string
  link-type="dfn"}. (If that string is then
  [parsed](#concept-host-parser){#ref-for-concept-host-parser④
  link-type="dfn"}, the result will
  [equal](#concept-host-equals){#ref-for-concept-host-equals
  link-type="dfn"} the [host](#concept-host){#ref-for-concept-host⑦
  link-type="dfn"} that was
  [serialized](#concept-host-serializer){#ref-for-concept-host-serializer②
  link-type="dfn"}.)

::: {#example-host-parsing .example}
[](#example-host-parsing){.self-link}

A [parse](#concept-host-parser){#ref-for-concept-host-parser⑤
link-type="dfn"}-[serialize](#concept-host-serializer){#ref-for-concept-host-serializer③
link-type="dfn"} roundtrip gives the following results, depending on the
`isOpaque`{.variable} argument to the [host
parser](#concept-host-parser){#ref-for-concept-host-parser⑥
link-type="dfn"}:

Input

Output (`isOpaque`{.variable} = false)

Output (`isOpaque`{.variable} = true)

`EXAMPLE.COM`

`example.com` ([domain](#concept-domain){#ref-for-concept-domain
link-type="dfn"})

`EXAMPLE.COM` ([opaque host](#opaque-host){#ref-for-opaque-host②
link-type="dfn"})

`example%2Ecom`

`example%2Ecom` ([opaque host](#opaque-host){#ref-for-opaque-host③
link-type="dfn"})

`faß.example`

`xn--fa-hia.example` ([domain](#concept-domain){#ref-for-concept-domain①
link-type="dfn"})

`fa%C3%9F.example` ([opaque host](#opaque-host){#ref-for-opaque-host④
link-type="dfn"})

`0`

`0.0.0.0` ([IPv4](#concept-ipv4){#ref-for-concept-ipv4⑨
link-type="dfn"})

`0` ([opaque host](#opaque-host){#ref-for-opaque-host⑤ link-type="dfn"})

`%30`

`%30` ([opaque host](#opaque-host){#ref-for-opaque-host⑥
link-type="dfn"})

`0x`

`0x` ([opaque host](#opaque-host){#ref-for-opaque-host⑦
link-type="dfn"})

`0xffffffff`

`255.255.255.255` ([IPv4](#concept-ipv4){#ref-for-concept-ipv4①⓪
link-type="dfn"})

`0xffffffff` ([opaque host](#opaque-host){#ref-for-opaque-host⑧
link-type="dfn"})

`[0:0::1]`

`[::1]` ([IPv6](#concept-ipv6){#ref-for-concept-ipv6①⓪ link-type="dfn"})

`[0:0::1%5D`

Failure

`[0:0::%31]`

`09`

Failure

`09` ([opaque host](#opaque-host){#ref-for-opaque-host⑨
link-type="dfn"})

`example.255`

`example.255` ([opaque host](#opaque-host){#ref-for-opaque-host①⓪
link-type="dfn"})

`example^example`

Failure
:::

### [3.1. ]{.secno}[Host representation]{.content}[](#host-representation){.self-link} {#host-representation .heading .settled level="3.1"}

A [host]{#concept-host .dfn .dfn-paneled dfn-type="dfn" export=""} is a
[domain](#concept-domain){#ref-for-concept-domain② link-type="dfn"}, an
[IP address](#ip-address){#ref-for-ip-address link-type="dfn"}, an
[opaque host](#opaque-host){#ref-for-opaque-host①① link-type="dfn"}, or
an [empty host](#empty-host){#ref-for-empty-host link-type="dfn"}.
Typically a [host](#concept-host){#ref-for-concept-host⑧
link-type="dfn"} serves as a network address, but it is sometimes used
as opaque identifier in [URLs](#concept-url){#ref-for-concept-url⑦
link-type="dfn"} where a network address is not necessary.

[](#example-opaque-host-url){.self-link}A typical
[URL](#concept-url){#ref-for-concept-url⑧ link-type="dfn"} whose
[host](#concept-url-host){#ref-for-concept-url-host link-type="dfn"} is
an [opaque host](#opaque-host){#ref-for-opaque-host①② link-type="dfn"}
is `git://github.com/whatwg/url.git`.

The RFCs referenced in the paragraphs below are for informative purposes
only. They have no influence on
[host](#concept-host){#ref-for-concept-host⑨ link-type="dfn"} writing,
parsing, and serialization. Unless stated otherwise in the sections that
follow.

A [domain]{#concept-domain .dfn .dfn-paneled dfn-type="dfn" export=""}
is a non-empty [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①
link-type="dfn"} that identifies a realm within a network.
[\[RFC1034\]](#biblio-rfc1034 "Domain names - concepts and facilities"){link-type="biblio"}

The [domain labels]{#domain-label .dfn .dfn-paneled dfn-type="dfn"
export="" lt="domain label"} of a
[domain](#concept-domain){#ref-for-concept-domain③ link-type="dfn"}
`domain`{.variable} are the result of [strictly
splitting](https://infra.spec.whatwg.org/#strictly-split){#ref-for-strictly-split
link-type="dfn"} `domain`{.variable} on U+002E (.).

The `example.com` and `example.com.`
[domains](#concept-domain){#ref-for-concept-domain④ link-type="dfn"} are
not equivalent and typically treated as distinct.

An [IP address]{#ip-address .dfn .dfn-paneled dfn-type="dfn" export=""}
is an [IPv4 address](#concept-ipv4){#ref-for-concept-ipv4①①
link-type="dfn"} or an [IPv6
address](#concept-ipv6){#ref-for-concept-ipv6①① link-type="dfn"}.

An [IPv4 address]{#concept-ipv4 .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [32-bit unsigned
integer](https://infra.spec.whatwg.org/#32-bit-unsigned-integer){#ref-for-32-bit-unsigned-integer
link-type="dfn"} that identifies a network address.
[\[RFC791\]](#biblio-rfc791 "Internet Protocol"){link-type="biblio"}

An [IPv6 address]{#concept-ipv6 .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [128-bit unsigned
integer](https://infra.spec.whatwg.org/#128-bit-unsigned-integer){#ref-for-128-bit-unsigned-integer
link-type="dfn"} that identifies a network address. This integer is
composed of a [list](https://infra.spec.whatwg.org/#list){#ref-for-list
link-type="dfn"} of 8 [16-bit unsigned
integers](https://infra.spec.whatwg.org/#16-bit-unsigned-integer){#ref-for-16-bit-unsigned-integer
link-type="dfn"}, also known as an [IPv6
address](#concept-ipv6){#ref-for-concept-ipv6①② link-type="dfn"}'s
[pieces]{#concept-ipv6-piece .dfn .dfn-paneled dfn-for="IPv6 address"
dfn-type="dfn" export=""}.
[\[RFC4291\]](#biblio-rfc4291 "IP Version 6 Addressing Architecture"){link-type="biblio"}

Support for `<zone_id>` is [intentionally
omitted](https://www.w3.org/Bugs/Public/show_bug.cgi?id=27234#c2).

An [opaque host]{#opaque-host .dfn .dfn-paneled dfn-type="dfn"
export=""} is a non-empty [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string②
link-type="dfn"} that can be used for further processing.

An [empty host]{#empty-host .dfn .dfn-paneled dfn-type="dfn" export=""}
is the empty string.

### [3.2. ]{.secno}[Host miscellaneous]{.content}[](#host-miscellaneous){.self-link} {#host-miscellaneous .heading .settled level="3.2"}

A [forbidden host code point]{#forbidden-host-code-point .dfn
.dfn-paneled dfn-type="dfn" export=""} is U+0000 NULL, U+0009 TAB,
U+000A LF, U+000D CR, U+0020 SPACE, U+0023 (#), U+002F (/), U+003A (:),
U+003C (\<), U+003E (\>), U+003F (?), U+0040 (@), U+005B (\[), U+005C
(\\), U+005D (\]), U+005E (\^), or U+007C (\|).

A [forbidden domain code point]{#forbidden-domain-code-point .dfn
.dfn-paneled dfn-type="dfn" export=""} is a [forbidden host code
point](#forbidden-host-code-point){#ref-for-forbidden-host-code-point①
link-type="dfn"}, a [C0
control](https://infra.spec.whatwg.org/#c0-control){#ref-for-c0-control①
link-type="dfn"}, U+0025 (%), or U+007F DELETE.

::: {.algorithm algorithm="public suffix" algorithm-for="host"}
To obtain the [public suffix]{#host-public-suffix .dfn .dfn-paneled
dfn-for="host" dfn-type="dfn" export=""} of a
[host](#concept-host){#ref-for-concept-host①⓪ link-type="dfn"}
`host`{.variable}, run these steps. They return null or a
[domain](#concept-domain){#ref-for-concept-domain⑤ link-type="dfn"}
representing a portion of `host`{.variable} that is included on the
Public Suffix List.
[\[PSL\]](#biblio-psl "Public Suffix List"){link-type="biblio"}

1.  If `host`{.variable} is not a
    [domain](#concept-domain){#ref-for-concept-domain⑥ link-type="dfn"},
    then return null.

2.  Let `trailingDot`{.variable} be \"`.`\" if `host`{.variable} [ends
    with](https://infra.spec.whatwg.org/#string-ends-with){#ref-for-string-ends-with
    link-type="dfn"} \"`.`\"; otherwise the empty string.

3.  Let `publicSuffix`{.variable} be the public suffix determined by
    running the [Public Suffix List
    algorithm](https://github.com/publicsuffix/list/wiki/Format#formal-algorithm)
    with `host`{.variable} as domain.
    [\[PSL\]](#biblio-psl "Public Suffix List"){link-type="biblio"}

4.  Assert: `publicSuffix`{.variable} is an [ASCII
    string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string③
    link-type="dfn"} that does not [end
    with](https://infra.spec.whatwg.org/#string-ends-with){#ref-for-string-ends-with①
    link-type="dfn"} \"`.`\".

5.  Return `publicSuffix`{.variable} and `trailingDot`{.variable}
    concatenated.
:::

::: {.algorithm algorithm="registrable domain" algorithm-for="host"}
To obtain the [registrable domain]{#host-registrable-domain .dfn
.dfn-paneled dfn-for="host" dfn-type="dfn" export=""} of a
[host](#concept-host){#ref-for-concept-host①① link-type="dfn"}
`host`{.variable}, run these steps. They return null or a
[domain](#concept-domain){#ref-for-concept-domain⑦ link-type="dfn"}
formed by `host`{.variable}'s [public
suffix](#host-public-suffix){#ref-for-host-public-suffix
link-type="dfn"} and the [domain
label](#domain-label){#ref-for-domain-label link-type="dfn"} preceding
it, if any.

1.  If `host`{.variable}'s [public
    suffix](#host-public-suffix){#ref-for-host-public-suffix①
    link-type="dfn"} is null or `host`{.variable}'s [public
    suffix](#host-public-suffix){#ref-for-host-public-suffix②
    link-type="dfn"}
    [equals](#concept-host-equals){#ref-for-concept-host-equals①
    link-type="dfn"} `host`{.variable}, then return null.

2.  Let `trailingDot`{.variable} be \"`.`\" if `host`{.variable} [ends
    with](https://infra.spec.whatwg.org/#string-ends-with){#ref-for-string-ends-with②
    link-type="dfn"} \"`.`\"; otherwise the empty string.

3.  Let `registrableDomain`{.variable} be the registrable domain
    determined by running the [Public Suffix List
    algorithm](https://github.com/publicsuffix/list/wiki/Format#formal-algorithm)
    with `host`{.variable} as domain.
    [\[PSL\]](#biblio-psl "Public Suffix List"){link-type="biblio"}

4.  Assert: `registrableDomain`{.variable} is an [ASCII
    string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string④
    link-type="dfn"} that does not [end
    with](https://infra.spec.whatwg.org/#string-ends-with){#ref-for-string-ends-with③
    link-type="dfn"} \"`.`\".

5.  Return `registrableDomain`{.variable} and `trailingDot`{.variable}
    concatenated.
:::

::: {#example-host-psl .example}
[](#example-host-psl){.self-link}

Host input

Public suffix

Registrable domain

`com`

`com`

null

`example.com`

`com`

`example.com`

`www.example.com`

`com`

`example.com`

`sub.www.example.com`

`com`

`example.com`

`EXAMPLE.COM`

`com`

`example.com`

`example.com.`

`com.`

`example.com.`

`github.io`

`github.io`

null

`whatwg.github.io`

`github.io`

`whatwg.github.io`

`إختبار`

`xn--kgbechtv`

null

`example.إختبار`

`xn--kgbechtv`

`example.xn--kgbechtv`

`sub.example.إختبار`

`xn--kgbechtv`

`example.xn--kgbechtv`

`[2001:0db8:85a3:0000:0000:8a2e:0370:7334]`

null

null
:::

Specifications should prefer the
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin
link-type="dfn"} concept for security decisions. The notion of \"[public
suffix](#host-public-suffix){#ref-for-host-public-suffix③
link-type="dfn"}\" and \"[registrable
domain](#host-registrable-domain){#ref-for-host-registrable-domain
link-type="dfn"}\" cannot be relied-upon to provide a hard security
boundary, as the public suffix list will diverge from client to client.
Specifications which ignore this advice are encouraged to carefully
consider whether URLs\' schemes ought to be incorporated into any
decisions made, i.e. whether to use the [same
site](https://html.spec.whatwg.org/multipage/browsers.html#same-site){#ref-for-same-site
link-type="dfn"} or [schemelessly same
site](https://html.spec.whatwg.org/multipage/browsers.html#schemelessly-same-site){#ref-for-schemelessly-same-site
link-type="dfn"} concepts.

### [3.3. ]{.secno}[IDNA]{.content}[](#idna){.self-link} {#idna .heading .settled level="3.3"}

::: {.algorithm algorithm="domain to ASCII"}
The [domain to ASCII]{#concept-domain-to-ascii .dfn .dfn-paneled
dfn-type="dfn" noexport=""} algorithm, given a
[string](https://infra.spec.whatwg.org/#string){#ref-for-string④
link-type="dfn"} `domain`{.variable} and a boolean
`beStrict`{.variable}, runs these steps:

1.  Let `result`{.variable} be the result of running [Unicode
    ToASCII](https://www.unicode.org/reports/tr46/#ToASCII){#ref-for-ToASCII②
    link-type="abstract-op"} with *domain_name* set to
    `domain`{.variable}, *CheckHyphens* set to `beStrict`{.variable},
    *CheckBidi* set to true, *CheckJoiners* set to true,
    *UseSTD3ASCIIRules* set to `beStrict`{.variable},
    *Transitional_Processing* set to false, *VerifyDnsLength* set to
    `beStrict`{.variable}, and *IgnoreInvalidPunycode* set to false.
    [\[UTS46\]](#biblio-uts46 "Unicode IDNA Compatibility Processing"){link-type="biblio"}

    If `beStrict`{.variable} is false, `domain`{.variable} is an [ASCII
    string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string⑤
    link-type="dfn"}, and [strictly
    splitting](https://infra.spec.whatwg.org/#strictly-split){#ref-for-strictly-split①
    link-type="dfn"} `domain`{.variable} on U+002E (.) does not produce
    any
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item
    link-type="dfn"} that [starts
    with](https://infra.spec.whatwg.org/#string-starts-with){#ref-for-string-starts-with
    link-type="dfn"} an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive
    link-type="dfn"} match for \"`xn--`\", this step is equivalent to
    [ASCII
    lowercasing](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase
    link-type="dfn"} `domain`{.variable}.

2.  If `result`{.variable} is a failure value,
    [domain-to-ASCII](#validation-error-domain-to-ascii){#ref-for-validation-error-domain-to-ascii①
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error③
    link-type="dfn"}, return failure.

3.  If `beStrict`{.variable} is false:

    1.  If `result`{.variable} is the empty string,
        [domain-to-ASCII](#validation-error-domain-to-ascii){#ref-for-validation-error-domain-to-ascii②
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error④
        link-type="dfn"}, return failure.

    2.  If `result`{.variable} contains a [forbidden domain code
        point](#forbidden-domain-code-point){#ref-for-forbidden-domain-code-point①
        link-type="dfn"},
        [domain-invalid-code-point](#domain-invalid-code-point){#ref-for-domain-invalid-code-point
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error⑤
        link-type="dfn"}, return failure.

        Due to web compatibility and compatibility with non-DNS-based
        systems the [forbidden domain code
        points](#forbidden-domain-code-point){#ref-for-forbidden-domain-code-point②
        link-type="dfn"} are a subset of those disallowed when
        *UseSTD3ASCIIRules* is true. See also [issue
        #397](https://github.com/whatwg/url/issues/397).

4.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert
    link-type="dfn"}: `result`{.variable} is not the empty string and
    does not contain a [forbidden domain code
    point](#forbidden-domain-code-point){#ref-for-forbidden-domain-code-point③
    link-type="dfn"}.

    Unicode IDNA Compatibility Processing guarantees this holds when
    `beStrict`{.variable} is true.
    [\[UTS46\]](#biblio-uts46 "Unicode IDNA Compatibility Processing"){link-type="biblio"}

5.  Return `result`{.variable}.

This document and the web platform at large use Unicode IDNA
Compatibility Processing and not IDNA2008. For instance, `☕.example`
becomes `xn--53h.example` and not failure.
[\[UTS46\]](#biblio-uts46 "Unicode IDNA Compatibility Processing"){link-type="biblio"}
[\[RFC5890\]](#biblio-rfc5890 "Internationalized Domain Names for Applications (IDNA): Definitions and Document Framework"){link-type="biblio"}
:::

::: {.algorithm algorithm="domain to Unicode"}
The [domain to Unicode]{#concept-domain-to-unicode .dfn .dfn-paneled
dfn-type="dfn" noexport=""} algorithm, given a
[domain](#concept-domain){#ref-for-concept-domain⑧ link-type="dfn"}
`domain`{.variable} and a boolean `beStrict`{.variable}, runs these
steps:

1.  Let `result`{.variable} be the result of running [Unicode
    ToUnicode](https://www.unicode.org/reports/tr46/#ToUnicode){#ref-for-ToUnicode①
    link-type="abstract-op"} with *domain_name* set to
    `domain`{.variable}, *CheckHyphens* set to `beStrict`{.variable},
    *CheckBidi* set to true, *CheckJoiners* set to true,
    *UseSTD3ASCIIRules* set to `beStrict`{.variable},
    *Transitional_Processing* set to false, and *IgnoreInvalidPunycode*
    set to false.
    [\[UTS46\]](#biblio-uts46 "Unicode IDNA Compatibility Processing"){link-type="biblio"}

2.  Signify
    [domain-to-Unicode](#domain-to-unicode){#ref-for-domain-to-unicode
    link-type="dfn"} [validation
    errors](#validation-error){#ref-for-validation-error⑥
    link-type="dfn"} for any returned errors, and then, return
    `result`{.variable}.
:::

### [3.4. ]{.secno}[]{#host-syntax .bs-old-id}[Host writing]{.content}[](#host-writing){.self-link} {#host-writing .heading .settled level="3.4"}

A [[]{#syntax-host .bs-old-id}valid host string]{#valid-host-string .dfn
.dfn-paneled dfn-type="dfn" export=""} must be a [valid domain
string](#valid-domain-string){#ref-for-valid-domain-string
link-type="dfn"}, a [valid IPv4-address
string](#valid-ipv4-address-string){#ref-for-valid-ipv4-address-string
link-type="dfn"}, or: U+005B (\[), followed by a [valid IPv6-address
string](#valid-ipv6-address-string){#ref-for-valid-ipv6-address-string
link-type="dfn"}, followed by U+005D (\]).

A [string](https://infra.spec.whatwg.org/#string){#ref-for-string⑤
link-type="dfn"} `input`{.variable} is a [valid domain]{#valid-domain
.dfn .dfn-paneled dfn-type="dfn" noexport=""} if these steps return
true:

1.  Let `domain`{.variable} be the result of running [domain to
    ASCII](#concept-domain-to-ascii){#ref-for-concept-domain-to-ascii
    link-type="dfn"} with `input`{.variable} and true.

2.  Return false if `domain`{.variable} is failure; otherwise true.

Ideally we define this in terms of a sequence of code points that make
up a [valid domain](#valid-domain){#ref-for-valid-domain
link-type="dfn"} rather than through a whack-a-mole: [issue
245](https://github.com/whatwg/url/issues/245).

A [[]{#syntax-host-domain .bs-old-id}valid domain
string]{#valid-domain-string .dfn .dfn-paneled dfn-type="dfn" export=""}
must be a string that is a [valid
domain](#valid-domain){#ref-for-valid-domain① link-type="dfn"}.

A [[]{#syntax-host-ipv4 .bs-old-id}valid IPv4-address
string]{#valid-ipv4-address-string .dfn .dfn-paneled dfn-type="dfn"
export=""} must be four shortest possible strings of [ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit②
link-type="dfn"}, representing a decimal number in the range 0 to 255,
inclusive, separated from each other by U+002E (.).

A [[]{#syntax-host-ipv6 .bs-old-id}valid IPv6-address
string]{#valid-ipv6-address-string .dfn .dfn-paneled dfn-type="dfn"
export=""} is defined in the [\"Text Representation of Addresses\"
chapter of IP Version 6 Addressing
Architecture](https://tools.ietf.org/html/rfc4291#section-2.2).
[\[RFC4291\]](#biblio-rfc4291 "IP Version 6 Addressing Architecture"){link-type="biblio"}

A [valid opaque-host string]{#valid-opaque-host-string .dfn .dfn-paneled
dfn-type="dfn" export=""} must be one of the following:

- one or more [URL units](#url-units){#ref-for-url-units①
  link-type="dfn"} excluding [forbidden host code
  points](#forbidden-host-code-point){#ref-for-forbidden-host-code-point②
  link-type="dfn"}

- U+005B (\[), followed by a [valid IPv6-address
  string](#valid-ipv6-address-string){#ref-for-valid-ipv6-address-string①
  link-type="dfn"}, followed by U+005D (\]).

This is not part of the definition of [valid host
string](#valid-host-string){#ref-for-valid-host-string② link-type="dfn"}
as it requires context to be distinguished.

### [3.5. ]{.secno}[Host parsing]{.content}[](#host-parsing){.self-link} {#host-parsing .heading .settled level="3.5"}

::: {.algorithm algorithm="host parser"}
The [host parser]{#concept-host-parser .dfn .dfn-paneled dfn-type="dfn"
export="" lt="host parser|host parsing"} takes a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string④
link-type="dfn"} `input`{.variable} with an optional boolean
`isOpaque`{.variable} (default false), and then runs these steps. They
return failure or a [host](#concept-host){#ref-for-concept-host①②
link-type="dfn"}.

1.  If `input`{.variable} starts with U+005B (\[), then:

    1.  If `input`{.variable} does not end with U+005D (\]),
        [IPv6-unclosed](#ipv6-unclosed){#ref-for-ipv6-unclosed
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error⑦
        link-type="dfn"}, return failure.

    2.  Return the result of [IPv6
        parsing](#concept-ipv6-parser){#ref-for-concept-ipv6-parser
        link-type="dfn"} `input`{.variable} with its leading U+005B (\[)
        and trailing U+005D (\]) removed.

2.  If `isOpaque`{.variable} is true, then return the result of
    [opaque-host
    parsing](#concept-opaque-host-parser){#ref-for-concept-opaque-host-parser
    link-type="dfn"} `input`{.variable}.

3.  Assert: `input`{.variable} is not the empty string.

4.  Let `domain`{.variable} be the result of running [UTF-8 decode
    without
    BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#ref-for-utf-8-decode-without-bom①
    link-type="dfn"} on the
    [percent-decoding](#string-percent-decode){#ref-for-string-percent-decode④
    link-type="dfn"} of `input`{.variable}.

    Alternatively [UTF-8 decode without BOM or
    fail](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom-or-fail){#ref-for-utf-8-decode-without-bom-or-fail①
    link-type="dfn"} can be used, coupled with an early return for
    failure, as [domain to
    ASCII](#concept-domain-to-ascii){#ref-for-concept-domain-to-ascii①
    link-type="dfn"} fails on U+FFFD (�).

5.  Let `asciiDomain`{.variable} be the result of running [domain to
    ASCII](#concept-domain-to-ascii){#ref-for-concept-domain-to-ascii②
    link-type="dfn"} with `domain`{.variable} and false.

6.  If `asciiDomain`{.variable} is failure, then return failure.

7.  If `asciiDomain`{.variable} [ends in a
    number](#ends-in-a-number-checker){#ref-for-ends-in-a-number-checker
    link-type="dfn"}, then return the result of [IPv4
    parsing](#concept-ipv4-parser){#ref-for-concept-ipv4-parser
    link-type="dfn"} `asciiDomain`{.variable}.

8.  Return `asciiDomain`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="ends in a number checker"}
The [ends in a number checker]{#ends-in-a-number-checker .dfn
.dfn-paneled dfn-type="dfn" noexport=""} takes an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string⑥
link-type="dfn"} `input`{.variable} and then runs these steps. They
return a boolean.

1.  Let `parts`{.variable} be the result of [strictly
    splitting](https://infra.spec.whatwg.org/#strictly-split){#ref-for-strictly-split②
    link-type="dfn"} `input`{.variable} on U+002E (.).

2.  If the last
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item①
    link-type="dfn"} in `parts`{.variable} is the empty string, then:

    1.  If `parts`{.variable}'s
        [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size
        link-type="dfn"} is 1, then return false.

    2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove
        link-type="dfn"} the last
        [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item②
        link-type="dfn"} from `parts`{.variable}.

3.  Let `last`{.variable} be the last
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item③
    link-type="dfn"} in `parts`{.variable}.

4.  If `last`{.variable} is non-empty and contains only [ASCII
    digits](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit③
    link-type="dfn"}, then return true.

    The erroneous input \"`09`\" will be caught by the [IPv4
    parser](#concept-ipv4-parser){#ref-for-concept-ipv4-parser①
    link-type="dfn"} at a later stage.

5.  If parsing `last`{.variable} as an [IPv4
    number](#ipv4-number-parser){#ref-for-ipv4-number-parser
    link-type="dfn"} does not return failure, then return true.

    This is equivalent to checking that `last`{.variable} is \"`0X`\" or
    \"`0x`\", followed by zero or more [ASCII hex
    digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit②
    link-type="dfn"}.

6.  Return false.
:::

::: {.algorithm algorithm="IPv4 parser"}
The [IPv4 parser]{#concept-ipv4-parser .dfn .dfn-paneled dfn-type="dfn"
noexport=""} takes an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string⑦
link-type="dfn"} `input`{.variable} and then runs these steps. They
return failure or an [IPv4
address](#concept-ipv4){#ref-for-concept-ipv4①② link-type="dfn"}.

The [IPv4 parser](#concept-ipv4-parser){#ref-for-concept-ipv4-parser②
link-type="dfn"} is not to be invoked directly. Instead check that the
return value of the [host
parser](#concept-host-parser){#ref-for-concept-host-parser⑦
link-type="dfn"} is an [IPv4
address](#concept-ipv4){#ref-for-concept-ipv4①③ link-type="dfn"}.

1.  Let `parts`{.variable} be the result of [strictly
    splitting](https://infra.spec.whatwg.org/#strictly-split){#ref-for-strictly-split③
    link-type="dfn"} `input`{.variable} on U+002E (.).

2.  If the last
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item④
    link-type="dfn"} in `parts`{.variable} is the empty string, then:

    1.  [IPv4-empty-part](#ipv4-empty-part){#ref-for-ipv4-empty-part
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error⑧
        link-type="dfn"}.

    2.  If `parts`{.variable}'s
        [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size①
        link-type="dfn"} is greater than 1, then
        [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove①
        link-type="dfn"} the last
        [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item⑤
        link-type="dfn"} from `parts`{.variable}.

3.  If `parts`{.variable}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size②
    link-type="dfn"} is greater than 4,
    [IPv4-too-many-parts](#ipv4-too-many-parts){#ref-for-ipv4-too-many-parts
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error⑨
    link-type="dfn"}, return failure.

4.  Let `numbers`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list①
    link-type="dfn"}.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate
    link-type="dfn"} `part`{.variable} of `parts`{.variable}:

    1.  Let `result`{.variable} be the result of
        [parsing](#ipv4-number-parser){#ref-for-ipv4-number-parser①
        link-type="dfn"} `part`{.variable}.

    2.  If `result`{.variable} is failure,
        [IPv4-non-numeric-part](#ipv4-non-numeric-part){#ref-for-ipv4-non-numeric-part
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error①⓪
        link-type="dfn"}, return failure.

    3.  If `result`{.variable}\[1\] is true,
        [IPv4-non-decimal-part](#ipv4-non-decimal-part){#ref-for-ipv4-non-decimal-part
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error①①
        link-type="dfn"}.

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append
        link-type="dfn"} `result`{.variable}\[0\] to
        `numbers`{.variable}.

6.  If any item in `numbers`{.variable} is greater than 255,
    [IPv4-out-of-range-part](#ipv4-out-of-range-part){#ref-for-ipv4-out-of-range-part
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error①②
    link-type="dfn"}.

7.  If any but the last
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item⑥
    link-type="dfn"} in `numbers`{.variable} is greater than 255, then
    return failure.

8.  If the last
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item⑦
    link-type="dfn"} in `numbers`{.variable} is greater than or equal to
    256^(5\ −\ `numbers`{.variable}'s\ [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size③
    link-type="dfn"})^, then return failure.

9.  Let `ipv4`{.variable} be the last
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item⑧
    link-type="dfn"} in `numbers`{.variable}.

10. [Remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove②
    link-type="dfn"} the last
    [item](https://infra.spec.whatwg.org/#list-item){#ref-for-list-item⑨
    link-type="dfn"} from `numbers`{.variable}.

11. Let `counter`{.variable} be 0.

12. [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①
    link-type="dfn"} `n`{.variable} of `numbers`{.variable}:

    1.  Increment `ipv4`{.variable} by `n`{.variable} ×
        256^(3\ −\ `counter`{.variable})^.

    2.  Increment `counter`{.variable} by 1.

13. Return `ipv4`{.variable}.
:::

::: {.algorithm algorithm="IPv4 number parser"}
The [IPv4 number parser]{#ipv4-number-parser .dfn .dfn-paneled
dfn-type="dfn" noexport=""} takes an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string⑧
link-type="dfn"} `input`{.variable} and then runs these steps. They
return failure or a
[tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple
link-type="dfn"} of a number and a boolean.

1.  If `input`{.variable} is the empty string, then return failure.

2.  Let `validationError`{.variable} be false.

3.  Let `R`{.variable} be 10.

4.  If `input`{.variable} contains at least two code points and the
    first two code points are either \"`0X`\" or \"`0x`\", then:

    1.  Set `validationError`{.variable} to true.

    2.  Remove the first two code points from `input`{.variable}.

    3.  Set `R`{.variable} to 16.

5.  Otherwise, if `input`{.variable} contains at least two code points
    and the first code point is U+0030 (0), then:

    1.  Set `validationError`{.variable} to true.

    2.  Remove the first code point from `input`{.variable}.

    3.  Set `R`{.variable} to 8.

6.  If `input`{.variable} is the empty string, then return (0, true).

7.  If `input`{.variable} contains a code point that is not a
    radix-`R`{.variable} digit, then return failure.

8.  Let `output`{.variable} be the mathematical integer value that is
    represented by `input`{.variable} in radix-`R`{.variable} notation,
    using [ASCII hex
    digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit③
    link-type="dfn"} for digits with values 0 through 15.

9.  Return (`output`{.variable}, `validationError`{.variable}).
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="IPv6 parser"}
The [IPv6 parser]{#concept-ipv6-parser .dfn .dfn-paneled dfn-type="dfn"
noexport=""} takes a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string⑤
link-type="dfn"} `input`{.variable} and then runs these steps. They
return failure or an [IPv6
address](#concept-ipv6){#ref-for-concept-ipv6①③ link-type="dfn"}.

The [IPv6 parser](#concept-ipv6-parser){#ref-for-concept-ipv6-parser①
link-type="dfn"} could in theory be invoked directly, but please discuss
actually doing that with the editors of this document first.

1.  Let `address`{.variable} be a new [IPv6
    address](#concept-ipv6){#ref-for-concept-ipv6①④ link-type="dfn"}
    whose [pieces](#concept-ipv6-piece){#ref-for-concept-ipv6-piece
    link-type="dfn"} are all 0.

2.  Let `pieceIndex`{.variable} be 0.

3.  Let `compress`{.variable} be null.

4.  Let `pointer`{.variable} be a [pointer](#pointer){#ref-for-pointer⑦
    link-type="dfn"} for `input`{.variable}.

5.  If [c](#c){#ref-for-c⑤ link-type="dfn"} is U+003A (:), then:

    1.  If [remaining](#remaining){#ref-for-remaining③ link-type="dfn"}
        does not start with U+003A (:),
        [IPv6-invalid-compression](#ipv6-invalid-compression){#ref-for-ipv6-invalid-compression
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error①③
        link-type="dfn"}, return failure.

    2.  Increase `pointer`{.variable} by 2.

    3.  Increase `pieceIndex`{.variable} by 1 and then set
        `compress`{.variable} to `pieceIndex`{.variable}.

6.  While [c](#c){#ref-for-c⑥ link-type="dfn"} is not the [EOF code
    point](#eof-code-point){#ref-for-eof-code-point③ link-type="dfn"}:

    1.  If `pieceIndex`{.variable} is 8,
        [IPv6-too-many-pieces](#ipv6-too-many-pieces){#ref-for-ipv6-too-many-pieces
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error①④
        link-type="dfn"}, return failure.

    2.  If [c](#c){#ref-for-c⑦ link-type="dfn"} is U+003A (:), then:

        1.  If `compress`{.variable} is non-null,
            [IPv6-multiple-compression](#ipv6-multiple-compression){#ref-for-ipv6-multiple-compression
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error①⑤
            link-type="dfn"}, return failure.

        2.  Increase `pointer`{.variable} and `pieceIndex`{.variable} by
            1, set `compress`{.variable} to `pieceIndex`{.variable}, and
            then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue①
            link-type="dfn"}.

    3.  Let `value`{.variable} and `length`{.variable} be 0.

    4.  While `length`{.variable} is less than 4 and [c](#c){#ref-for-c⑧
        link-type="dfn"} is an [ASCII hex
        digit](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit④
        link-type="dfn"}, set `value`{.variable} to `value`{.variable} ×
        0x10 + [c](#c){#ref-for-c⑨ link-type="dfn"} interpreted as
        hexadecimal number, and increase `pointer`{.variable} and
        `length`{.variable} by 1.

    5.  If [c](#c){#ref-for-c①⓪ link-type="dfn"} is U+002E (.), then:

        1.  If `length`{.variable} is 0,
            [IPv4-in-IPv6-invalid-code-point](#ipv4-in-ipv6-invalid-code-point){#ref-for-ipv4-in-ipv6-invalid-code-point
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error①⑥
            link-type="dfn"}, return failure.

        2.  Decrease `pointer`{.variable} by `length`{.variable}.

        3.  If `pieceIndex`{.variable} is greater than 6,
            [IPv4-in-IPv6-too-many-pieces](#ipv4-in-ipv6-too-many-pieces){#ref-for-ipv4-in-ipv6-too-many-pieces
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error①⑦
            link-type="dfn"}, return failure.

        4.  Let `numbersSeen`{.variable} be 0.

        5.  While [c](#c){#ref-for-c①① link-type="dfn"} is not the [EOF
            code point](#eof-code-point){#ref-for-eof-code-point④
            link-type="dfn"}:

            1.  Let `ipv4Piece`{.variable} be null.

            2.  If `numbersSeen`{.variable} is greater than 0, then:

                1.  If [c](#c){#ref-for-c①② link-type="dfn"} is a U+002E
                    (.) and `numbersSeen`{.variable} is less than 4,
                    then increase `pointer`{.variable} by 1.

                2.  Otherwise,
                    [IPv4-in-IPv6-invalid-code-point](#ipv4-in-ipv6-invalid-code-point){#ref-for-ipv4-in-ipv6-invalid-code-point①
                    link-type="dfn"} [validation
                    error](#validation-error){#ref-for-validation-error①⑧
                    link-type="dfn"}, return failure.

            3.  If [c](#c){#ref-for-c①③ link-type="dfn"} is not an
                [ASCII
                digit](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit④
                link-type="dfn"},
                [IPv4-in-IPv6-invalid-code-point](#ipv4-in-ipv6-invalid-code-point){#ref-for-ipv4-in-ipv6-invalid-code-point②
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error①⑨
                link-type="dfn"}, return failure.

            4.  While [c](#c){#ref-for-c①④ link-type="dfn"} is an [ASCII
                digit](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit⑤
                link-type="dfn"}:

                1.  Let `number`{.variable} be [c](#c){#ref-for-c①⑤
                    link-type="dfn"} interpreted as decimal number.

                2.  If `ipv4Piece`{.variable} is null, then set
                    `ipv4Piece`{.variable} to `number`{.variable}.

                    Otherwise, if `ipv4Piece`{.variable} is 0,
                    [IPv4-in-IPv6-invalid-code-point](#ipv4-in-ipv6-invalid-code-point){#ref-for-ipv4-in-ipv6-invalid-code-point③
                    link-type="dfn"} [validation
                    error](#validation-error){#ref-for-validation-error②⓪
                    link-type="dfn"}, return failure.

                    Otherwise, set `ipv4Piece`{.variable} to
                    `ipv4Piece`{.variable} × 10 + `number`{.variable}.

                3.  If `ipv4Piece`{.variable} is greater than 255,
                    [IPv4-in-IPv6-out-of-range-part](#ipv4-in-ipv6-out-of-range-part){#ref-for-ipv4-in-ipv6-out-of-range-part
                    link-type="dfn"} [validation
                    error](#validation-error){#ref-for-validation-error②①
                    link-type="dfn"}, return failure.

                4.  Increase `pointer`{.variable} by 1.

            5.  Set `address`{.variable}\[`pieceIndex`{.variable}\] to
                `address`{.variable}\[`pieceIndex`{.variable}\] ×
                0x100 + `ipv4Piece`{.variable}.

            6.  Increase `numbersSeen`{.variable} by 1.

            7.  If `numbersSeen`{.variable} is 2 or 4, then increase
                `pieceIndex`{.variable} by 1.

        6.  If `numbersSeen`{.variable} is not 4,
            [IPv4-in-IPv6-too-few-parts](#ipv4-in-ipv6-too-few-parts){#ref-for-ipv4-in-ipv6-too-few-parts
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error②②
            link-type="dfn"}, return failure.

        7.  [Break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break
            link-type="dfn"}.

    6.  Otherwise, if [c](#c){#ref-for-c①⑥ link-type="dfn"} is U+003A
        (:):

        1.  Increase `pointer`{.variable} by 1.

        2.  If [c](#c){#ref-for-c①⑦ link-type="dfn"} is the [EOF code
            point](#eof-code-point){#ref-for-eof-code-point⑤
            link-type="dfn"},
            [IPv6-invalid-code-point](#ipv6-invalid-code-point){#ref-for-ipv6-invalid-code-point
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error②③
            link-type="dfn"}, return failure.

    7.  Otherwise, if [c](#c){#ref-for-c①⑧ link-type="dfn"} is not the
        [EOF code point](#eof-code-point){#ref-for-eof-code-point⑥
        link-type="dfn"},
        [IPv6-invalid-code-point](#ipv6-invalid-code-point){#ref-for-ipv6-invalid-code-point①
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error②④
        link-type="dfn"}, return failure.

    8.  Set `address`{.variable}\[`pieceIndex`{.variable}\] to
        `value`{.variable}.

    9.  Increase `pieceIndex`{.variable} by 1.

7.  If `compress`{.variable} is non-null, then:

    1.  Let `swaps`{.variable} be `pieceIndex`{.variable} −
        `compress`{.variable}.

    2.  Set `pieceIndex`{.variable} to 7.

    3.  While `pieceIndex`{.variable} is not 0 and `swaps`{.variable} is
        greater than 0, swap
        `address`{.variable}\[`pieceIndex`{.variable}\] with
        `address`{.variable}\[`compress`{.variable} + `swaps`{.variable}
        − 1\], and then decrease both `pieceIndex`{.variable} and
        `swaps`{.variable} by 1.

8.  Otherwise, if `compress`{.variable} is null and
    `pieceIndex`{.variable} is not 8,
    [IPv6-too-few-pieces](#ipv6-too-few-pieces){#ref-for-ipv6-too-few-pieces
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error②⑤
    link-type="dfn"}, return failure.

9.  Return `address`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="opaque-host parser"}
The [opaque-host parser]{#concept-opaque-host-parser .dfn .dfn-paneled
dfn-type="dfn" export=""} takes a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string⑥
link-type="dfn"} `input`{.variable}, and then runs these steps. They
return failure or an [opaque host](#opaque-host){#ref-for-opaque-host①③
link-type="dfn"}.

1.  If `input`{.variable} contains a [forbidden host code
    point](#forbidden-host-code-point){#ref-for-forbidden-host-code-point③
    link-type="dfn"},
    [host-invalid-code-point](#host-invalid-code-point){#ref-for-host-invalid-code-point
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error②⑥
    link-type="dfn"}, return failure.

2.  If `input`{.variable} contains a [code
    point](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑤
    link-type="dfn"} that is not a [URL code
    point](#url-code-points){#ref-for-url-code-points link-type="dfn"}
    and not U+0025 (%),
    [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error②⑦
    link-type="dfn"}.

3.  If `input`{.variable} contains a U+0025 (%) and the two [code
    points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑥
    link-type="dfn"} following it are not [ASCII hex
    digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit⑤
    link-type="dfn"},
    [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit①
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error②⑧
    link-type="dfn"}.

4.  Return the result of running [UTF-8
    percent-encode](#string-utf-8-percent-encode){#ref-for-string-utf-8-percent-encode②
    link-type="dfn"} on `input`{.variable} using the [C0 control
    percent-encode
    set](#c0-control-percent-encode-set){#ref-for-c0-control-percent-encode-set②
    link-type="dfn"}.
:::

### [3.6. ]{.secno}[Host serializing]{.content}[](#host-serializing){.self-link} {#host-serializing .heading .settled level="3.6"}

::: {.algorithm algorithm="host serializer"}
The [host serializer]{#concept-host-serializer .dfn .dfn-paneled
dfn-type="dfn" noexport=""} takes a
[host](#concept-host){#ref-for-concept-host①③ link-type="dfn"}
`host`{.variable} and then runs these steps. They return an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string⑨
link-type="dfn"}.

1.  If `host`{.variable} is an [IPv4
    address](#concept-ipv4){#ref-for-concept-ipv4①④ link-type="dfn"},
    return the result of running the [IPv4
    serializer](#concept-ipv4-serializer){#ref-for-concept-ipv4-serializer
    link-type="dfn"} on `host`{.variable}.

2.  Otherwise, if `host`{.variable} is an [IPv6
    address](#concept-ipv6){#ref-for-concept-ipv6①⑤ link-type="dfn"},
    return U+005B (\[), followed by the result of running the [IPv6
    serializer](#concept-ipv6-serializer){#ref-for-concept-ipv6-serializer
    link-type="dfn"} on `host`{.variable}, followed by U+005D (\]).

3.  Otherwise, `host`{.variable} is a
    [domain](#concept-domain){#ref-for-concept-domain⑨ link-type="dfn"},
    [opaque host](#opaque-host){#ref-for-opaque-host①④ link-type="dfn"},
    or [empty host](#empty-host){#ref-for-empty-host① link-type="dfn"},
    return `host`{.variable}.
:::

::: {.algorithm algorithm="IPv4 serializer"}
The [IPv4 serializer]{#concept-ipv4-serializer .dfn .dfn-paneled
dfn-type="dfn" noexport=""} takes an [IPv4
address](#concept-ipv4){#ref-for-concept-ipv4①⑤ link-type="dfn"}
`address`{.variable} and then runs these steps. They return an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①⓪
link-type="dfn"}.

1.  Let `output`{.variable} be the empty string.

2.  Let `n`{.variable} be the value of `address`{.variable}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②
    link-type="dfn"} `i`{.variable} in the range 1 to 4, inclusive:

    1.  Prepend `n`{.variable} % 256,
        [serialized](#serialize-an-integer){#ref-for-serialize-an-integer
        link-type="dfn"}, to `output`{.variable}.

    2.  If `i`{.variable} is not 4, then prepend U+002E (.) to
        `output`{.variable}.

    3.  Set `n`{.variable} to floor(`n`{.variable} / 256).

4.  Return `output`{.variable}.
:::

::: {.algorithm algorithm="IPv6 serializer"}
The [IPv6 serializer]{#concept-ipv6-serializer .dfn .dfn-paneled
dfn-type="dfn" noexport=""} takes an [IPv6
address](#concept-ipv6){#ref-for-concept-ipv6①⑥ link-type="dfn"}
`address`{.variable} and then runs these steps. They return an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①①
link-type="dfn"}.

1.  Let `output`{.variable} be the empty string.

2.  Let `compress`{.variable} be the result of [finding the IPv6 address
    compressed piece
    index](#find-the-ipv6-address-compressed-piece-index){#ref-for-find-the-ipv6-address-compressed-piece-index
    link-type="dfn"} given `address`{.variable}.

3.  Let `ignore0`{.variable} be false.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate③
    link-type="dfn"} `pieceIndex`{.variable} of `address`{.variable}'s
    [pieces](#concept-ipv6-piece){#ref-for-concept-ipv6-piece①
    link-type="dfn"}'s
    [indices](https://infra.spec.whatwg.org/#list-get-the-indices){#ref-for-list-get-the-indices
    link-type="dfn"}:

    1.  If `ignore0`{.variable} is true and
        `address`{.variable}\[`pieceIndex`{.variable}\] is 0, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue②
        link-type="dfn"}.

    2.  Otherwise, if `ignore0`{.variable} is true, set
        `ignore0`{.variable} to false.

    3.  If `compress`{.variable} is `pieceIndex`{.variable}, then:

        1.  Let `separator`{.variable} be \"`::`\" if
            `pieceIndex`{.variable} is 0; otherwise U+003A (:).

        2.  Append `separator`{.variable} to `output`{.variable}.

        3.  Set `ignore0`{.variable} to true and
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue③
            link-type="dfn"}.

    4.  Append `address`{.variable}\[`pieceIndex`{.variable}\],
        represented as the shortest possible lowercase hexadecimal
        number, to `output`{.variable}.

    5.  If `pieceIndex`{.variable} is not 7, then append U+003A (:) to
        `output`{.variable}.

5.  Return `output`{.variable}.

This algorithm requires the recommendation from A Recommendation for
IPv6 Address Text Representation.
[\[RFC5952\]](#biblio-rfc5952 "A Recommendation for IPv6 Address Text Representation"){link-type="biblio"}
:::

::: {.algorithm algorithm="find the IPv6 address compressed piece index"}
To [find the IPv6 address compressed piece
index]{#find-the-ipv6-address-compressed-piece-index .dfn .dfn-paneled
dfn-type="dfn" noexport=""} given an [IPv6
address](#concept-ipv6){#ref-for-concept-ipv6①⑦ link-type="dfn"}
`address`{.variable}:

1.  Let `longestIndex`{.variable} be null.

2.  Let `longestSize`{.variable} be 1.

3.  Let `foundIndex`{.variable} be null.

4.  Let `foundSize`{.variable} be 0.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate④
    link-type="dfn"} `pieceIndex`{.variable} of `address`{.variable}'s
    [pieces](#concept-ipv6-piece){#ref-for-concept-ipv6-piece②
    link-type="dfn"}'s
    [indices](https://infra.spec.whatwg.org/#list-get-the-indices){#ref-for-list-get-the-indices①
    link-type="dfn"}:

    1.  If `address`{.variable}'s
        [pieces](#concept-ipv6-piece){#ref-for-concept-ipv6-piece③
        link-type="dfn"}\[`pieceIndex`{.variable}\] is not 0:

        1.  If `foundSize`{.variable} is greater than
            `longestSize`{.variable}, then set `longestIndex`{.variable}
            to `foundIndex`{.variable} and `longestSize`{.variable} to
            `foundSize`{.variable}.

        2.  Set `foundIndex`{.variable} to null.

        3.  Set `foundSize`{.variable} to 0.

    2.  Otherwise:

        1.  If `foundIndex`{.variable} is null, then set
            `foundIndex`{.variable} to `pieceIndex`{.variable}.

        2.  Increment `foundSize`{.variable} by 1.

6.  If `foundSize`{.variable} is greater than `longestSize`{.variable},
    then return `foundIndex`{.variable}.

7.  Return `longestIndex`{.variable}.

[](#example-e2b3492e){.self-link}In `0:f:0:0:f:f:0:0` it would point to
the second 0.
:::

### [3.7. ]{.secno}[Host equivalence]{.content}[](#host-equivalence){.self-link} {#host-equivalence .heading .settled level="3.7"}

::: {.algorithm algorithm="equal" algorithm-for="host"}
To determine whether a [host](#concept-host){#ref-for-concept-host①④
link-type="dfn"} `A`{.variable} [equals]{#concept-host-equals .dfn
.dfn-paneled dfn-for="host" dfn-type="dfn" export="" lt="equal"}
[host](#concept-host){#ref-for-concept-host①⑤ link-type="dfn"}
`B`{.variable}, return true if `A`{.variable} is `B`{.variable}, and
false otherwise.
:::

Certificate comparison requires a host equivalence check that ignores
the trailing dot of a domain (if any). However, those hosts have also
various other facets enforced, such as DNS length, that are not enforced
here, as URLs do not enforce them. If anyone has a good suggestion for
how to bring these two closer together, or what a good unified model
would be, please file an issue.

## [4. ]{.secno}[URLs]{.content}[](#urls){.self-link} {#urls .heading .settled level="4"}

At a high level, a [URL](#concept-url){#ref-for-concept-url⑨
link-type="dfn"}, [valid URL
string](#valid-url-string){#ref-for-valid-url-string link-type="dfn"},
[URL parser](#concept-url-parser){#ref-for-concept-url-parser①
link-type="dfn"}, and [URL
serializer](#concept-url-serializer){#ref-for-concept-url-serializer
link-type="dfn"} relate as follows:

- The [URL parser](#concept-url-parser){#ref-for-concept-url-parser②
  link-type="dfn"} takes an arbitrary [scalar value
  string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string⑦
  link-type="dfn"} and returns either failure or a
  [URL](#concept-url){#ref-for-concept-url①⓪ link-type="dfn"}. It might
  also record zero or more [validation
  errors](#validation-error){#ref-for-validation-error②⑨
  link-type="dfn"}.

- A [URL](#concept-url){#ref-for-concept-url①① link-type="dfn"} can be
  seen as the in-memory representation.

- A [valid URL string](#valid-url-string){#ref-for-valid-url-string①
  link-type="dfn"} defines what input would not trigger a [validation
  error](#validation-error){#ref-for-validation-error③⓪ link-type="dfn"}
  or failure when given to the [URL
  parser](#concept-url-parser){#ref-for-concept-url-parser③
  link-type="dfn"}. I.e., input that would be considered conforming or
  valid.

- The [URL
  serializer](#concept-url-serializer){#ref-for-concept-url-serializer①
  link-type="dfn"} takes a [URL](#concept-url){#ref-for-concept-url①②
  link-type="dfn"} and returns an [ASCII
  string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①②
  link-type="dfn"}. (If that string is then
  [parsed](#concept-url-parser){#ref-for-concept-url-parser④
  link-type="dfn"}, the result will
  [equal](#concept-url-equals){#ref-for-concept-url-equals
  link-type="dfn"} the [URL](#concept-url){#ref-for-concept-url①③
  link-type="dfn"} that was
  [serialized](#concept-url-serializer){#ref-for-concept-url-serializer②
  link-type="dfn"}.) The output of the [URL
  serializer](#concept-url-serializer){#ref-for-concept-url-serializer③
  link-type="dfn"} is not always a [valid URL
  string](#valid-url-string){#ref-for-valid-url-string②
  link-type="dfn"}.

::: {#example-url-parsing .example}
[](#example-url-parsing){.self-link}

Input

Base

Valid

Output

`https:example.org`

❌

`https://example.org/`

`https://////example.com///`

❌

`https://example.com///`

`https://example.com/././foo`

✅

`https://example.com/foo`

`hello:world`

`https://example.com/`

✅

`hello:world`

`https:example.org`

`https://example.com/`

❌

`https://example.com/example.org`

`\example\..\demo/.\`

`https://example.com/`

❌

`https://example.com/demo/`

`example`

`https://example.com/demo`

✅

`https://example.com/example`

`file:///C|/demo`

❌

`file:///C:/demo`

`..`

`file:///C:/demo`

✅

`file:///C:/`

`file://loc%61lhost/`

✅

`file:///`

`https://user:password@example.org/`

❌

`https://user:password@example.org/`

`https://example.org/foo bar`

❌

`https://example.org/foo%20bar`

`https://EXAMPLE.com/../x`

✅

`https://example.com/x`

`https://ex ample.org/`

❌

Failure

`example`

❌, due to lack of base

Failure

`https://example.com:demo`

❌

Failure

`http://[www.example.com]/`

❌

Failure

`https://example.org//`

✅

`https://example.org//`

`https://example.com/[]?[]#[]`

❌

`https://example.com/[]?[]#[]`

`https://example/%?%#%`

❌

`https://example/%?%#%`

`https://example/%25?%25#%25`

✅

`https://example/%25?%25#%25`

The base and output [URL](#concept-url){#ref-for-concept-url①④
link-type="dfn"} are represented in
[serialized](#concept-url-serializer){#ref-for-concept-url-serializer④
link-type="dfn"} form for brevity.
:::

### [4.1. ]{.secno}[URL representation]{.content}[](#url-representation){.self-link} {#url-representation .heading .settled level="4.1"}

A [URL]{#concept-url .dfn .dfn-paneled dfn-type="dfn" export=""
lt="URL|URL record"} is a
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct
link-type="dfn"} that represents a universal identifier. To disambiguate
from a [valid URL string](#valid-url-string){#ref-for-valid-url-string③
link-type="dfn"} it can also be referred to as a [URL
record](#concept-url){#ref-for-concept-url①⑤ link-type="dfn"}.

A [URL](#concept-url){#ref-for-concept-url①⑥ link-type="dfn"}'s
[scheme]{#concept-url-scheme .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} is an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①③
link-type="dfn"} that identifies the type of
[URL](#concept-url){#ref-for-concept-url①⑦ link-type="dfn"} and can be
used to dispatch a [URL](#concept-url){#ref-for-concept-url①⑧
link-type="dfn"} for further processing after
[parsing](#concept-url-parser){#ref-for-concept-url-parser⑤
link-type="dfn"}. It is initially the empty string.

A [URL](#concept-url){#ref-for-concept-url①⑨ link-type="dfn"}'s
[username]{#concept-url-username .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} is an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①④
link-type="dfn"} identifying a username. It is initially the empty
string.

A [URL](#concept-url){#ref-for-concept-url②⓪ link-type="dfn"}'s
[password]{#concept-url-password .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} is an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①⑤
link-type="dfn"} identifying a password. It is initially the empty
string.

A [URL](#concept-url){#ref-for-concept-url②① link-type="dfn"}'s
[host]{#concept-url-host .dfn .dfn-paneled dfn-for="url" dfn-type="dfn"
export=""} is null or a [host](#concept-host){#ref-for-concept-host①⑥
link-type="dfn"}. It is initially null.

::: {.note role="note"}
The following table lists allowed
[URL](#concept-url){#ref-for-concept-url②② link-type="dfn"}'s
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme④
link-type="dfn"} / [host](#concept-url-host){#ref-for-concept-url-host①
link-type="dfn"} combinations.

[scheme](#concept-url-scheme){#ref-for-concept-url-scheme⑤
link-type="dfn"}

[host](#concept-url-host){#ref-for-concept-url-host② link-type="dfn"}

[domain](#concept-domain){#ref-for-concept-domain①⓪ link-type="dfn"}

[IPv4 address](#concept-ipv4){#ref-for-concept-ipv4①⑥ link-type="dfn"}

[IPv6 address](#concept-ipv6){#ref-for-concept-ipv6①⑧ link-type="dfn"}

[opaque host](#opaque-host){#ref-for-opaque-host①⑤ link-type="dfn"}

[empty host](#empty-host){#ref-for-empty-host② link-type="dfn"}

null

[Special schemes](#special-scheme){#ref-for-special-scheme②
link-type="dfn"} excluding \"`file`\"

✅

✅

✅

❌

❌

❌

\"`file`\"

✅

✅

✅

❌

✅

❌

Others

❌

❌

✅

✅

✅

✅
:::

A [URL](#concept-url){#ref-for-concept-url②③ link-type="dfn"}'s
[port]{#concept-url-port .dfn .dfn-paneled dfn-for="url" dfn-type="dfn"
export=""} is either null or a [16-bit unsigned
integer](https://infra.spec.whatwg.org/#16-bit-unsigned-integer){#ref-for-16-bit-unsigned-integer①
link-type="dfn"} that identifies a networking port. It is initially
null.

A [URL](#concept-url){#ref-for-concept-url②④ link-type="dfn"}'s
[[]{#url-cannot-be-a-base-url-flag .bs-old-id}[]{#non-relative-flag
.bs-old-id}path]{#concept-url-path .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} is a [URL path](#url-path){#ref-for-url-path
link-type="dfn"}, usually identifying a location. It is initially « ».

A [special](#is-special){#ref-for-is-special① link-type="dfn"}
[URL](#concept-url){#ref-for-concept-url②⑤ link-type="dfn"}'s
[path](#concept-url-path){#ref-for-concept-url-path① link-type="dfn"} is
always a [list](https://infra.spec.whatwg.org/#list){#ref-for-list②
link-type="dfn"}, i.e., it is never
[opaque](#url-opaque-path){#ref-for-url-opaque-path② link-type="dfn"}.

A [URL](#concept-url){#ref-for-concept-url②⑥ link-type="dfn"}'s
[query]{#concept-url-query .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} is either null or an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①⑥
link-type="dfn"}. It is initially null.

A [URL](#concept-url){#ref-for-concept-url②⑦ link-type="dfn"}'s
[fragment]{#concept-url-fragment .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} is either null or an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①⑦
link-type="dfn"} that can be used for further processing on the resource
the [URL](#concept-url){#ref-for-concept-url②⑧ link-type="dfn"}'s other
components identify. It is initially null.

A [URL](#concept-url){#ref-for-concept-url②⑨ link-type="dfn"} also has
an associated [blob URL entry]{#concept-url-blob-entry .dfn .dfn-paneled
dfn-for="url" dfn-type="dfn" export=""} that is either null or a [blob
URL
entry](https://w3c.github.io/FileAPI/#blob-url-entry){#ref-for-blob-url-entry
link-type="dfn"}. It is initially null.

This is used to support caching the object a \"`blob`\" URL refers to as
well as its origin. It is important that these are cached as the
[URL](#concept-url){#ref-for-concept-url③⓪ link-type="dfn"} might be
removed from the [blob URL
store](https://w3c.github.io/FileAPI/#BlobURLStore){#ref-for-BlobURLStore
link-type="dfn"} between parsing and fetching, while fetching will still
need to succeed.

::: {#example-url-components .example}
[](#example-url-components){.self-link}

The following table lists how [valid URL
strings](#valid-url-string){#ref-for-valid-url-string④ link-type="dfn"},
when [parsed](#concept-url-parser){#ref-for-concept-url-parser⑥
link-type="dfn"}, map to a [URL](#concept-url){#ref-for-concept-url③①
link-type="dfn"}'s components.
[Username](#concept-url-username){#ref-for-concept-url-username
link-type="dfn"},
[password](#concept-url-password){#ref-for-concept-url-password
link-type="dfn"}, and [blob URL
entry](#concept-url-blob-entry){#ref-for-concept-url-blob-entry
link-type="dfn"} are omitted; in the examples below they are the empty
string, the empty string, and null, respectively.

Input

[Scheme](#concept-url-scheme){#ref-for-concept-url-scheme⑥
link-type="dfn"}

[Host](#concept-url-host){#ref-for-concept-url-host③ link-type="dfn"}

[Port](#concept-url-port){#ref-for-concept-url-port link-type="dfn"}

[Path](#concept-url-path){#ref-for-concept-url-path② link-type="dfn"}

[Query](#concept-url-query){#ref-for-concept-url-query① link-type="dfn"}

[Fragment](#concept-url-fragment){#ref-for-concept-url-fragment①
link-type="dfn"}

`https://example.com/`

\"`https`\"

\"`example.com`\"

null

« the empty string »

null

null

`https://localhost:8000/search?q=text#hello`

\"`https`\"

\"`localhost`\"

8000

« \"`search`\" »

\"`q=text`\"

\"`hello`\"

`urn:isbn:9780307476463`

\"`urn`\"

null

null

\"`isbn:9780307476463`\"

null

null

`file:///ada/Analytical%20Engine/README.md`

\"`file`\"

null

null

« \"`ada`\", \"`Analytical%20Engine`\", \"`README.md`\" »

null

null
:::

------------------------------------------------------------------------

A [URL path]{#url-path .dfn .dfn-paneled dfn-type="dfn" export=""} is
either a [URL path segment](#url-path-segment){#ref-for-url-path-segment
link-type="dfn"} or a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list③
link-type="dfn"} of zero or more [URL path
segments](#url-path-segment){#ref-for-url-path-segment①
link-type="dfn"}.

A [URL path segment]{#url-path-segment .dfn .dfn-paneled dfn-type="dfn"
export=""} is an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①⑧
link-type="dfn"}. It commonly refers to a directory or a file, but has
no predefined meaning.

A [[]{#syntax-url-path-segment-dot .bs-old-id}single-dot URL path
segment]{#single-dot-path-segment .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [URL path
segment](#url-path-segment){#ref-for-url-path-segment② link-type="dfn"}
that is \"`.`\" or an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive①
link-type="dfn"} match for \"`%2e`\".

A [[]{#syntax-url-path-segment-dotdot .bs-old-id}double-dot URL path
segment]{#double-dot-path-segment .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [URL path
segment](#url-path-segment){#ref-for-url-path-segment③ link-type="dfn"}
that is \"`..`\" or an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive②
link-type="dfn"} match for \"`.%2e`\", \"`%2e.`\", or \"`%2e%2e`\".

### [4.2. ]{.secno}[URL miscellaneous]{.content}[](#url-miscellaneous){.self-link} {#url-miscellaneous .heading .settled level="4.2"}

A [special scheme]{#special-scheme .dfn .dfn-paneled dfn-type="dfn"
export=""} is an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string①⑨
link-type="dfn"} that is listed in the first column of the following
table. The [default port]{#default-port .dfn .dfn-paneled dfn-type="dfn"
export=""} for a [special
scheme](#special-scheme){#ref-for-special-scheme③ link-type="dfn"} is
listed in the second column on the same row. The [default
port](#default-port){#ref-for-default-port link-type="dfn"} for any
other [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string②⓪
link-type="dfn"} is null.

[Special scheme](#special-scheme){#ref-for-special-scheme④
link-type="dfn"}

[Default port](#default-port){#ref-for-default-port① link-type="dfn"}

\"`ftp`\"

21

\"`file`\"

null

\"`http`\"

80

\"`https`\"

443

\"`ws`\"

80

\"`wss`\"

443

A [URL](#concept-url){#ref-for-concept-url③② link-type="dfn"} [is
special]{#is-special .dfn .dfn-paneled dfn-type="dfn" export=""} if its
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme⑦
link-type="dfn"} is a [special
scheme](#special-scheme){#ref-for-special-scheme⑤ link-type="dfn"}. A
[URL](#concept-url){#ref-for-concept-url③③ link-type="dfn"} [is not
special]{#is-not-special .dfn .dfn-paneled dfn-type="dfn" noexport=""}
if its [scheme](#concept-url-scheme){#ref-for-concept-url-scheme⑧
link-type="dfn"} is not a [special
scheme](#special-scheme){#ref-for-special-scheme⑥ link-type="dfn"}.

A [URL](#concept-url){#ref-for-concept-url③④ link-type="dfn"} [includes
credentials]{#include-credentials .dfn .dfn-paneled dfn-type="dfn"
export="" lt="include credentials|includes credentials"} if its
[username](#concept-url-username){#ref-for-concept-url-username①
link-type="dfn"} or
[password](#concept-url-password){#ref-for-concept-url-password①
link-type="dfn"} is not the empty string.

A [URL](#concept-url){#ref-for-concept-url③⑤ link-type="dfn"} has an
[opaque path]{#url-opaque-path .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} if its
[path](#concept-url-path){#ref-for-concept-url-path③ link-type="dfn"} is
a [URL path segment](#url-path-segment){#ref-for-url-path-segment④
link-type="dfn"}.

A [URL](#concept-url){#ref-for-concept-url③⑥ link-type="dfn"} [cannot
have a username/password/port]{#cannot-have-a-username-password-port
.dfn .dfn-paneled dfn-type="dfn" export=""} if its
[host](#concept-url-host){#ref-for-concept-url-host④ link-type="dfn"} is
null or the empty string, or its
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme⑨
link-type="dfn"} is \"`file`\".

A [URL](#concept-url){#ref-for-concept-url③⑦ link-type="dfn"} can be
designated as [base URL]{#concept-base-url .dfn .dfn-paneled
dfn-type="dfn" noexport=""}.

A [base URL](#concept-base-url){#ref-for-concept-base-url⑥
link-type="dfn"} is useful for the [URL
parser](#concept-url-parser){#ref-for-concept-url-parser⑦
link-type="dfn"} when the input might be a [relative-URL
string](#relative-url-string){#ref-for-relative-url-string①
link-type="dfn"}.

------------------------------------------------------------------------

A [Windows drive letter]{#windows-drive-letter .dfn .dfn-paneled
dfn-type="dfn" noexport=""} is two code points, of which the first is an
[ASCII
alpha](https://infra.spec.whatwg.org/#ascii-alpha){#ref-for-ascii-alpha①
link-type="dfn"} and the second is either U+003A (:) or U+007C (\|).

A [normalized Windows drive letter]{#normalized-windows-drive-letter
.dfn .dfn-paneled dfn-type="dfn" noexport=""} is a [Windows drive
letter](#windows-drive-letter){#ref-for-windows-drive-letter
link-type="dfn"} of which the second code point is U+003A (:).

As per the [URL writing](#url-writing) section, only a [normalized
Windows drive
letter](#normalized-windows-drive-letter){#ref-for-normalized-windows-drive-letter
link-type="dfn"} is conforming.

A string [starts with a Windows drive
letter]{#start-with-a-windows-drive-letter .dfn .dfn-paneled
dfn-type="dfn"
lt="start with a Windows drive letter|starts with a Windows drive letter"
noexport=""} if all of the following are true:

- its
  [length](https://infra.spec.whatwg.org/#string-length){#ref-for-string-length
  link-type="dfn"} is greater than or equal to 2
- its first two code points are a [Windows drive
  letter](#windows-drive-letter){#ref-for-windows-drive-letter①
  link-type="dfn"}
- its
  [length](https://infra.spec.whatwg.org/#string-length){#ref-for-string-length①
  link-type="dfn"} is 2 or its third code point is U+002F (/), U+005C
  (\\), U+003F (?), or U+0023 (#).

::: {#example-start-with-a-widows-drive-letter .example}
[](#example-start-with-a-widows-drive-letter){.self-link}

String

Starts with a Windows drive letter

\"`c:`\"

✅

\"`c:/`\"

✅

\"`c:a`\"

❌
:::

::: {.algorithm algorithm="shorten a url’s path"}
To [shorten a `url`{.variable}'s path]{#shorten-a-urls-path .dfn
.dfn-paneled dfn-type="dfn" local-lt="shorten" noexport=""}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①
    link-type="dfn"}: `url`{.variable} does not have an [opaque
    path](#url-opaque-path){#ref-for-url-opaque-path③ link-type="dfn"}.

2.  Let `path`{.variable} be `url`{.variable}'s
    [path](#concept-url-path){#ref-for-concept-url-path④
    link-type="dfn"}.

3.  If `url`{.variable}'s
    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①⓪
    link-type="dfn"} is \"`file`\", `path`{.variable}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size④
    link-type="dfn"} is 1, and `path`{.variable}\[0\] is a [normalized
    Windows drive
    letter](#normalized-windows-drive-letter){#ref-for-normalized-windows-drive-letter①
    link-type="dfn"}, then return.

4.  [Remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove③
    link-type="dfn"} `path`{.variable}'s last item, if any.
:::

### [4.3. ]{.secno}[]{#url-syntax .bs-old-id}[URL writing]{.content}[](#url-writing){.self-link} {#url-writing .heading .settled level="4.3"}

A [[]{#syntax-url .bs-old-id}valid URL string]{#valid-url-string .dfn
.dfn-paneled dfn-type="dfn" export=""} must be either a
[relative-URL-with-fragment
string](#relative-url-with-fragment-string){#ref-for-relative-url-with-fragment-string
link-type="dfn"} or an [absolute-URL-with-fragment
string](#absolute-url-with-fragment-string){#ref-for-absolute-url-with-fragment-string
link-type="dfn"}.

An [[]{#syntax-url-absolute-with-fragment
.bs-old-id}absolute-URL-with-fragment
string]{#absolute-url-with-fragment-string .dfn .dfn-paneled
dfn-type="dfn" export=""} must be an [absolute-URL
string](#absolute-url-string){#ref-for-absolute-url-string
link-type="dfn"}, optionally followed by U+0023 (#) and a [URL-fragment
string](#url-fragment-string){#ref-for-url-fragment-string
link-type="dfn"}.

An [[]{#syntax-url-absolute .bs-old-id}absolute-URL
string]{#absolute-url-string .dfn .dfn-paneled dfn-type="dfn" export=""}
must be one of the following:

- a [URL-scheme string](#url-scheme-string){#ref-for-url-scheme-string
  link-type="dfn"} that is an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive③
  link-type="dfn"} match for a [special
  scheme](#special-scheme){#ref-for-special-scheme⑦ link-type="dfn"} and
  not an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive④
  link-type="dfn"} match for \"`file`\", followed by U+003A (:) and a
  [scheme-relative-special-URL
  string](#scheme-relative-special-url-string){#ref-for-scheme-relative-special-url-string
  link-type="dfn"}

- a [URL-scheme string](#url-scheme-string){#ref-for-url-scheme-string①
  link-type="dfn"} that is *not* an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive⑤
  link-type="dfn"} match for a [special
  scheme](#special-scheme){#ref-for-special-scheme⑧ link-type="dfn"},
  followed by U+003A (:) and a [relative-URL
  string](#relative-url-string){#ref-for-relative-url-string②
  link-type="dfn"}

- a [URL-scheme string](#url-scheme-string){#ref-for-url-scheme-string②
  link-type="dfn"} that is an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive⑥
  link-type="dfn"} match for \"`file`\", followed by U+003A (:) and a
  [scheme-relative-file-URL
  string](#scheme-relative-file-url-string){#ref-for-scheme-relative-file-url-string
  link-type="dfn"}

any optionally followed by U+003F (?) and a [URL-query
string](#url-query-string){#ref-for-url-query-string link-type="dfn"}.

A [[]{#syntax-url-scheme .bs-old-id}URL-scheme
string]{#url-scheme-string .dfn .dfn-paneled dfn-type="dfn" export=""}
must be one [ASCII
alpha](https://infra.spec.whatwg.org/#ascii-alpha){#ref-for-ascii-alpha②
link-type="dfn"}, followed by zero or more of [ASCII
alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric){#ref-for-ascii-alphanumeric①
link-type="dfn"}, U+002B (+), U+002D (-), and U+002E (.).
[Schemes](#url-scheme-string){#ref-for-url-scheme-string③
link-type="dfn"} should be registered in the IANA URI \[sic\] Schemes
registry.
[\[IANA-URI-SCHEMES\]](#biblio-iana-uri-schemes "Uniform Resource Identifier (URI) Schemes"){link-type="biblio"}
[\[RFC7595\]](#biblio-rfc7595 "Guidelines and Registration Procedures for URI Schemes"){link-type="biblio"}

A [[]{#syntax-url-relative-with-fragment
.bs-old-id}relative-URL-with-fragment
string]{#relative-url-with-fragment-string .dfn .dfn-paneled
dfn-type="dfn" export=""} must be a [relative-URL
string](#relative-url-string){#ref-for-relative-url-string③
link-type="dfn"}, optionally followed by U+0023 (#) and a [URL-fragment
string](#url-fragment-string){#ref-for-url-fragment-string①
link-type="dfn"}.

A [[]{#syntax-url-relative .bs-old-id}relative-URL
string]{#relative-url-string .dfn .dfn-paneled dfn-type="dfn" export=""}
must be one of the following, switching on [base
URL](#concept-base-url){#ref-for-concept-base-url⑦ link-type="dfn"}'s
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme①①
link-type="dfn"}:

A [special scheme](#special-scheme){#ref-for-special-scheme⑨ link-type="dfn"} that is not \"`file`\"

:   a [scheme-relative-special-URL
    string](#scheme-relative-special-url-string){#ref-for-scheme-relative-special-url-string①
    link-type="dfn"}

:   a [path-absolute-URL
    string](#path-absolute-url-string){#ref-for-path-absolute-url-string
    link-type="dfn"}

:   a [path-relative-scheme-less-URL
    string](#path-relative-scheme-less-url-string){#ref-for-path-relative-scheme-less-url-string
    link-type="dfn"}

\"`file`\"

:   a [scheme-relative-file-URL
    string](#scheme-relative-file-url-string){#ref-for-scheme-relative-file-url-string①
    link-type="dfn"}

:   a [path-absolute-URL
    string](#path-absolute-url-string){#ref-for-path-absolute-url-string①
    link-type="dfn"} if [base
    URL](#concept-base-url){#ref-for-concept-base-url⑧
    link-type="dfn"}'s
    [host](#concept-url-host){#ref-for-concept-url-host⑤
    link-type="dfn"} is an [empty
    host](#empty-host){#ref-for-empty-host③ link-type="dfn"}

:   a [path-absolute-non-Windows-file-URL
    string](#path-absolute-non-windows-file-url-string){#ref-for-path-absolute-non-windows-file-url-string
    link-type="dfn"} if [base
    URL](#concept-base-url){#ref-for-concept-base-url⑨
    link-type="dfn"}'s
    [host](#concept-url-host){#ref-for-concept-url-host⑥
    link-type="dfn"} is not an [empty
    host](#empty-host){#ref-for-empty-host④ link-type="dfn"}

:   a [path-relative-scheme-less-URL
    string](#path-relative-scheme-less-url-string){#ref-for-path-relative-scheme-less-url-string①
    link-type="dfn"}

Otherwise

:   a [scheme-relative-URL
    string](#scheme-relative-url-string){#ref-for-scheme-relative-url-string
    link-type="dfn"}

:   a [path-absolute-URL
    string](#path-absolute-url-string){#ref-for-path-absolute-url-string②
    link-type="dfn"}

:   a [path-relative-scheme-less-URL
    string](#path-relative-scheme-less-url-string){#ref-for-path-relative-scheme-less-url-string②
    link-type="dfn"}

any optionally followed by U+003F (?) and a [URL-query
string](#url-query-string){#ref-for-url-query-string① link-type="dfn"}.

A non-null [base URL](#concept-base-url){#ref-for-concept-base-url①⓪
link-type="dfn"} is necessary when
[parsing](#concept-url-parser){#ref-for-concept-url-parser⑧
link-type="dfn"} a [relative-URL
string](#relative-url-string){#ref-for-relative-url-string④
link-type="dfn"}.

A [scheme-relative-special-URL
string]{#scheme-relative-special-url-string .dfn .dfn-paneled
dfn-type="dfn" export=""} must be \"`//`\", followed by a [valid host
string](#valid-host-string){#ref-for-valid-host-string③
link-type="dfn"}, optionally followed by U+003A (:) and a [URL-port
string](#url-port-string){#ref-for-url-port-string link-type="dfn"},
optionally followed by a [path-absolute-URL
string](#path-absolute-url-string){#ref-for-path-absolute-url-string③
link-type="dfn"}.

A [[]{#syntax-url-port .bs-old-id}URL-port string]{#url-port-string .dfn
.dfn-paneled dfn-type="dfn" export=""} must be one of the following:

- the empty string

- one or more [ASCII
  digits](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit⑥
  link-type="dfn"} representing a decimal number that is a [16-bit
  unsigned
  integer](https://infra.spec.whatwg.org/#16-bit-unsigned-integer){#ref-for-16-bit-unsigned-integer②
  link-type="dfn"}.

A [[]{#syntax-url-scheme-relative .bs-old-id}scheme-relative-URL
string]{#scheme-relative-url-string .dfn .dfn-paneled dfn-type="dfn"
export=""} must be \"`//`\", followed by an [opaque-host-and-port
string](#opaque-host-and-port-string){#ref-for-opaque-host-and-port-string
link-type="dfn"}, optionally followed by a [path-absolute-URL
string](#path-absolute-url-string){#ref-for-path-absolute-url-string④
link-type="dfn"}.

An [opaque-host-and-port string]{#opaque-host-and-port-string .dfn
.dfn-paneled dfn-type="dfn" export=""} must be either the empty string
or: a [valid opaque-host
string](#valid-opaque-host-string){#ref-for-valid-opaque-host-string
link-type="dfn"}, optionally followed by U+003A (:) and a [URL-port
string](#url-port-string){#ref-for-url-port-string① link-type="dfn"}.

A [[]{#syntax-url-file-scheme-relative
.bs-old-id}scheme-relative-file-URL
string]{#scheme-relative-file-url-string .dfn .dfn-paneled
dfn-type="dfn" export=""} must be \"`//`\", followed by one of the
following:

- a [valid host string](#valid-host-string){#ref-for-valid-host-string④
  link-type="dfn"}, optionally followed by a
  [path-absolute-non-Windows-file-URL
  string](#path-absolute-non-windows-file-url-string){#ref-for-path-absolute-non-windows-file-url-string①
  link-type="dfn"}

- a [path-absolute-URL
  string](#path-absolute-url-string){#ref-for-path-absolute-url-string⑤
  link-type="dfn"}.

A [[]{#syntax-url-path-absolute .bs-old-id}path-absolute-URL
string]{#path-absolute-url-string .dfn .dfn-paneled dfn-type="dfn"
export=""} must be U+002F (/) followed by a [path-relative-URL
string](#path-relative-url-string){#ref-for-path-relative-url-string
link-type="dfn"}.

A [[]{#syntax-url-file-path-absolute
.bs-old-id}path-absolute-non-Windows-file-URL
string]{#path-absolute-non-windows-file-url-string .dfn .dfn-paneled
dfn-type="dfn" export=""} must be a [path-absolute-URL
string](#path-absolute-url-string){#ref-for-path-absolute-url-string⑥
link-type="dfn"} that does not start with: U+002F (/), followed by a
[Windows drive
letter](#windows-drive-letter){#ref-for-windows-drive-letter②
link-type="dfn"}, followed by U+002F (/).

A [[]{#syntax-url-path-relative .bs-old-id}path-relative-URL
string]{#path-relative-url-string .dfn .dfn-paneled dfn-type="dfn"
export=""} must be zero or more [URL-path-segment
strings](#url-path-segment-string){#ref-for-url-path-segment-string
link-type="dfn"}, separated from each other by U+002F (/), and not start
with U+002F (/).

A [[]{#syntax-url-path-relative-scheme-less
.bs-old-id}path-relative-scheme-less-URL
string]{#path-relative-scheme-less-url-string .dfn .dfn-paneled
dfn-type="dfn" export=""} must be a [path-relative-URL
string](#path-relative-url-string){#ref-for-path-relative-url-string①
link-type="dfn"} that does not start with: a [URL-scheme
string](#url-scheme-string){#ref-for-url-scheme-string④
link-type="dfn"}, followed by U+003A (:).

A [[]{#syntax-url-path-segment .bs-old-id}URL-path-segment
string]{#url-path-segment-string .dfn .dfn-paneled dfn-type="dfn"
export=""} must be one of the following:

- zero or more [URL units](#url-units){#ref-for-url-units②
  link-type="dfn"} excluding U+002F (/) and U+003F (?), that together
  are not a [single-dot URL path
  segment](#single-dot-path-segment){#ref-for-single-dot-path-segment
  link-type="dfn"} or a [double-dot URL path
  segment](#double-dot-path-segment){#ref-for-double-dot-path-segment
  link-type="dfn"}.

- a [single-dot URL path
  segment](#single-dot-path-segment){#ref-for-single-dot-path-segment①
  link-type="dfn"}

- a [double-dot URL path
  segment](#double-dot-path-segment){#ref-for-double-dot-path-segment①
  link-type="dfn"}.

A [[]{#syntax-url-query .bs-old-id}URL-query string]{#url-query-string
.dfn .dfn-paneled dfn-type="dfn" export=""} must be zero or more [URL
units](#url-units){#ref-for-url-units③ link-type="dfn"}.

A [[]{#syntax-url-fragment .bs-old-id}URL-fragment
string]{#url-fragment-string .dfn .dfn-paneled dfn-type="dfn" export=""}
must be zero or more [URL units](#url-units){#ref-for-url-units④
link-type="dfn"}.

The [URL code points]{#url-code-points .dfn .dfn-paneled dfn-type="dfn"
export="" lt="URL code point"} are [ASCII
alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric){#ref-for-ascii-alphanumeric②
link-type="dfn"}, U+0021 (!), U+0024 (\$), U+0026 (&), U+0027 (\'),
U+0028 LEFT PARENTHESIS, U+0029 RIGHT PARENTHESIS, U+002A (\*), U+002B
(+), U+002C (,), U+002D (-), U+002E (.), U+002F (/), U+003A (:), U+003B
(;), U+003D (=), U+003F (?), U+0040 (@), U+005F (\_), U+007E (\~), and
[code
points](https://infra.spec.whatwg.org/#code-point){#ref-for-code-point⑦
link-type="dfn"} in the range U+00A0 to U+10FFFD, inclusive, excluding
[surrogates](https://infra.spec.whatwg.org/#surrogate){#ref-for-surrogate
link-type="dfn"} and
[noncharacters](https://infra.spec.whatwg.org/#noncharacter){#ref-for-noncharacter
link-type="dfn"}.

Code points greater than U+007F DELETE will be converted to
[percent-encoded
bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte③
link-type="dfn"} by the [URL
parser](#concept-url-parser){#ref-for-concept-url-parser⑨
link-type="dfn"}.

In HTML, when the document encoding is a legacy encoding, code points in
the [URL-query string](#url-query-string){#ref-for-url-query-string②
link-type="dfn"} that are higher than U+007F DELETE will be converted to
[percent-encoded
bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte④
link-type="dfn"} *using the document's encoding*. This can cause
problems if a URL that works in one document is copied to another
document that uses a different document encoding. Using the
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8③
link-type="dfn"} encoding everywhere solves this problem.

::: {#query-encoding-example .example}
[](#query-encoding-example){.self-link}

For example, consider this HTML document:

``` {.lang-html .highlight}
<!doctype html>
<meta charset="windows-1252">
<a href="?sm&ouml;rg&aring;sbord">Test</a>
```

Since the document encoding is windows-1252, the link's
[URL](#concept-url){#ref-for-concept-url③⑧ link-type="dfn"}'s
[query](#concept-url-query){#ref-for-concept-url-query② link-type="dfn"}
will be \"`sm%F6rg%E5sbord`\". If the document encoding had been UTF-8,
it would instead be \"`sm%C3%B6rg%C3%A5sbord`\".
:::

The [URL units]{#url-units .dfn .dfn-paneled dfn-type="dfn" noexport=""}
are [URL code points](#url-code-points){#ref-for-url-code-points①
link-type="dfn"} and [percent-encoded
bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte⑤
link-type="dfn"}.

[Percent-encoded
bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte⑥
link-type="dfn"} can be used to encode code points that are not [URL
code points](#url-code-points){#ref-for-url-code-points②
link-type="dfn"} or are excluded from being written.

------------------------------------------------------------------------

There is no way to express a
[username](#concept-url-username){#ref-for-concept-url-username②
link-type="dfn"} or
[password](#concept-url-password){#ref-for-concept-url-password②
link-type="dfn"} of a [URL record](#concept-url){#ref-for-concept-url③⑨
link-type="dfn"} within a [valid URL
string](#valid-url-string){#ref-for-valid-url-string⑤ link-type="dfn"}.

### [4.4. ]{.secno}[URL parsing]{.content}[](#url-parsing){.self-link} {#url-parsing .heading .settled level="4.4"}

::: {.algorithm algorithm="URL parser"}
The [URL parser]{#concept-url-parser .dfn .dfn-paneled dfn-type="dfn"
export=""} takes a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string⑧
link-type="dfn"} `input`{.variable}, with an optional null or [base
URL](#concept-base-url){#ref-for-concept-base-url①① link-type="dfn"}
`base`{.variable} (default null) and an optional
[encoding](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding①
link-type="dfn"} `encoding`{.variable} (default
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8④
link-type="dfn"}), and then runs these steps:

Non-web-browser implementations only need to implement the [basic URL
parser](#concept-basic-url-parser){#ref-for-concept-basic-url-parser
link-type="dfn"}.

How user input in the web browser's address bar is converted to a [URL
record](#concept-url){#ref-for-concept-url④⓪ link-type="dfn"} is
out-of-scope of this standard. This standard does include [URL rendering
requirements](#url-rendering) as they pertain trust decisions.

1.  Let `url`{.variable} be the result of running the [basic URL
    parser](#concept-basic-url-parser){#ref-for-concept-basic-url-parser①
    link-type="dfn"} on `input`{.variable} with `base`{.variable} and
    `encoding`{.variable}.

2.  If `url`{.variable} is failure, return failure.

3.  If `url`{.variable}'s
    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①②
    link-type="dfn"} is not \"`blob`\", return `url`{.variable}.

4.  Set `url`{.variable}'s [blob URL
    entry](#concept-url-blob-entry){#ref-for-concept-url-blob-entry①
    link-type="dfn"} to the result of [resolving the blob
    URL](https://w3c.github.io/FileAPI/#blob-url-resolve){#ref-for-blob-url-resolve
    link-type="dfn"} `url`{.variable}, if that did not return failure,
    and null otherwise.

5.  Return `url`{.variable}.
:::

------------------------------------------------------------------------

:::: {.algorithm algorithm="basic URL parser"}
The [basic URL parser]{#concept-basic-url-parser .dfn .dfn-paneled
dfn-type="dfn" export=""} takes a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string⑨
link-type="dfn"} `input`{.variable}, with an optional null or [base
URL](#concept-base-url){#ref-for-concept-base-url①② link-type="dfn"}
`base`{.variable} (default null), an optional
[encoding](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding②
link-type="dfn"} `encoding`{.variable} (default
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8⑤
link-type="dfn"}), an optional
[URL](#concept-url){#ref-for-concept-url④① link-type="dfn"}
[`url`{.variable}]{#basic-url-parser-url .dfn .dfn-paneled
dfn-for="basic URL parser" dfn-type="dfn" export=""}, and an optional
state override
[`state override`{.variable}]{#basic-url-parser-state-override .dfn
.dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}, and
then runs these steps:

::: {.note role="note"}
The `encoding`{.variable} argument is a legacy concept only relevant for
HTML. The `url`{.variable} and `state override`{.variable} arguments are
only for use by various APIs.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

When the `url`{.variable} and `state override`{.variable} arguments are
not passed, the [basic URL
parser](#concept-basic-url-parser){#ref-for-concept-basic-url-parser②
link-type="dfn"} returns either a new
[URL](#concept-url){#ref-for-concept-url④② link-type="dfn"} or failure.
If they are passed, the algorithm modifies the passed `url`{.variable}
and can terminate without returning anything.
:::

1.  If `url`{.variable} is not given:

    1.  Set `url`{.variable} to a new
        [URL](#concept-url){#ref-for-concept-url④③ link-type="dfn"}.

    2.  If `input`{.variable} contains any leading or trailing [C0
        control or
        space](https://infra.spec.whatwg.org/#c0-control-or-space){#ref-for-c0-control-or-space
        link-type="dfn"},
        [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit②
        link-type="dfn"} [validation
        error](#validation-error){#ref-for-validation-error③①
        link-type="dfn"}.

    3.  Remove any leading and trailing [C0 control or
        space](https://infra.spec.whatwg.org/#c0-control-or-space){#ref-for-c0-control-or-space①
        link-type="dfn"} from `input`{.variable}.

2.  If `input`{.variable} contains any [ASCII tab or
    newline](https://infra.spec.whatwg.org/#ascii-tab-or-newline){#ref-for-ascii-tab-or-newline
    link-type="dfn"},
    [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit③
    link-type="dfn"} [validation
    error](#validation-error){#ref-for-validation-error③②
    link-type="dfn"}.

3.  Remove all [ASCII tab or
    newline](https://infra.spec.whatwg.org/#ascii-tab-or-newline){#ref-for-ascii-tab-or-newline①
    link-type="dfn"} from `input`{.variable}.

4.  Let `state`{.variable} be `state override`{.variable} if given, or
    [scheme start
    state](#scheme-start-state){#ref-for-scheme-start-state
    link-type="dfn"} otherwise.

5.  Set `encoding`{.variable} to the result of [getting an output
    encoding](https://encoding.spec.whatwg.org/#get-an-output-encoding){#ref-for-get-an-output-encoding
    link-type="dfn"} from `encoding`{.variable}.

6.  Let `buffer`{.variable} be the empty string.

7.  Let `atSignSeen`{.variable}, `insideBrackets`{.variable}, and
    `passwordTokenSeen`{.variable} be false.

8.  Let `pointer`{.variable} be a [pointer](#pointer){#ref-for-pointer⑧
    link-type="dfn"} for `input`{.variable}.

9.  Keep running the following state machine by switching on
    `state`{.variable}. If after a run `pointer`{.variable} points to
    the [EOF code point](#eof-code-point){#ref-for-eof-code-point⑦
    link-type="dfn"}, go to the next step. Otherwise, increase
    `pointer`{.variable} by 1 and continue with the state machine.

    [scheme start state]{#scheme-start-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c①⑨ link-type="dfn"} is an [ASCII
            alpha](https://infra.spec.whatwg.org/#ascii-alpha){#ref-for-ascii-alpha③
            link-type="dfn"}, append [c](#c){#ref-for-c②⓪
            link-type="dfn"},
            [lowercased](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase①
            link-type="dfn"}, to `buffer`{.variable}, and set
            `state`{.variable} to [scheme
            state](#scheme-state){#ref-for-scheme-state
            link-type="dfn"}.

        2.  Otherwise, if `state override`{.variable} is not given, set
            `state`{.variable} to [no scheme
            state](#no-scheme-state){#ref-for-no-scheme-state
            link-type="dfn"} and decrease `pointer`{.variable} by 1.

        3.  Otherwise, return failure.

            This indication of failure is used exclusively by the
            [`Location`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#location){#ref-for-location
            link-type="idl"} object's
            [`protocol`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#dom-location-protocol){#ref-for-dom-location-protocol
            link-type="idl"} setter.

    [scheme state]{#scheme-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c②① link-type="dfn"} is an [ASCII
            alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric){#ref-for-ascii-alphanumeric③
            link-type="dfn"}, U+002B (+), U+002D (-), or U+002E (.),
            append [c](#c){#ref-for-c②② link-type="dfn"},
            [lowercased](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase②
            link-type="dfn"}, to `buffer`{.variable}.

        2.  Otherwise, if [c](#c){#ref-for-c②③ link-type="dfn"} is
            U+003A (:), then:

            1.  If `state override`{.variable} is given, then:

                1.  If `url`{.variable}'s
                    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①③
                    link-type="dfn"} is a [special
                    scheme](#special-scheme){#ref-for-special-scheme①⓪
                    link-type="dfn"} and `buffer`{.variable} is not a
                    [special
                    scheme](#special-scheme){#ref-for-special-scheme①①
                    link-type="dfn"}, then return.

                2.  If `url`{.variable}'s
                    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①④
                    link-type="dfn"} is not a [special
                    scheme](#special-scheme){#ref-for-special-scheme①②
                    link-type="dfn"} and `buffer`{.variable} is a
                    [special
                    scheme](#special-scheme){#ref-for-special-scheme①③
                    link-type="dfn"}, then return.

                3.  If `url`{.variable} [includes
                    credentials](#include-credentials){#ref-for-include-credentials①
                    link-type="dfn"} or has a non-null
                    [port](#concept-url-port){#ref-for-concept-url-port①
                    link-type="dfn"}, and `buffer`{.variable} is
                    \"`file`\", then return.

                4.  If `url`{.variable}'s
                    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①⑤
                    link-type="dfn"} is \"`file`\" and its
                    [host](#concept-url-host){#ref-for-concept-url-host⑦
                    link-type="dfn"} is an [empty
                    host](#empty-host){#ref-for-empty-host⑤
                    link-type="dfn"}, then return.

            2.  Set `url`{.variable}'s
                [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①⑥
                link-type="dfn"} to `buffer`{.variable}.

            3.  If `state override`{.variable} is given, then:

                1.  If `url`{.variable}'s
                    [port](#concept-url-port){#ref-for-concept-url-port②
                    link-type="dfn"} is `url`{.variable}'s
                    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①⑦
                    link-type="dfn"}'s [default
                    port](#default-port){#ref-for-default-port②
                    link-type="dfn"}, then set `url`{.variable}'s
                    [port](#concept-url-port){#ref-for-concept-url-port③
                    link-type="dfn"} to null.

                2.  Return.

            4.  Set `buffer`{.variable} to the empty string.

            5.  If `url`{.variable}'s
                [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①⑧
                link-type="dfn"} is \"`file`\", then:

                1.  If [remaining](#remaining){#ref-for-remaining④
                    link-type="dfn"} does not start with \"`//`\",
                    [special-scheme-missing-following-solidus](#special-scheme-missing-following-solidus){#ref-for-special-scheme-missing-following-solidus
                    link-type="dfn"} [validation
                    error](#validation-error){#ref-for-validation-error③③
                    link-type="dfn"}.

                2.  Set `state`{.variable} to [file
                    state](#file-state){#ref-for-file-state
                    link-type="dfn"}.

            6.  Otherwise, if `url`{.variable} [is
                special](#is-special){#ref-for-is-special②
                link-type="dfn"}, `base`{.variable} is non-null, and
                `base`{.variable}'s
                [scheme](#concept-url-scheme){#ref-for-concept-url-scheme①⑨
                link-type="dfn"} is `url`{.variable}'s
                [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②⓪
                link-type="dfn"}:

                1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②
                    link-type="dfn"}: `base`{.variable} [is
                    special](#is-special){#ref-for-is-special③
                    link-type="dfn"} (and therefore does not have an
                    [opaque
                    path](#url-opaque-path){#ref-for-url-opaque-path④
                    link-type="dfn"}).

                2.  Set `state`{.variable} to [special relative or
                    authority
                    state](#special-relative-or-authority-state){#ref-for-special-relative-or-authority-state
                    link-type="dfn"}.

            7.  Otherwise, if `url`{.variable} [is
                special](#is-special){#ref-for-is-special④
                link-type="dfn"}, set `state`{.variable} to [special
                authority slashes
                state](#special-authority-slashes-state){#ref-for-special-authority-slashes-state
                link-type="dfn"}.

            8.  Otherwise, if
                [remaining](#remaining){#ref-for-remaining⑤
                link-type="dfn"} starts with an U+002F (/), set
                `state`{.variable} to [path or authority
                state](#path-or-authority-state){#ref-for-path-or-authority-state
                link-type="dfn"} and increase `pointer`{.variable} by 1.

            9.  Otherwise, set `url`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path⑤
                link-type="dfn"} to the empty string and set
                `state`{.variable} to [opaque path
                state](#cannot-be-a-base-url-path-state){#ref-for-cannot-be-a-base-url-path-state
                link-type="dfn"}.

        3.  Otherwise, if `state override`{.variable} is not given, set
            `buffer`{.variable} to the empty string, `state`{.variable}
            to [no scheme
            state](#no-scheme-state){#ref-for-no-scheme-state①
            link-type="dfn"}, and start over (from the first code point
            in `input`{.variable}).

        4.  Otherwise, return failure.

            This indication of failure is used exclusively by the
            [`Location`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#location){#ref-for-location①
            link-type="idl"} object's
            [`protocol`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#dom-location-protocol){#ref-for-dom-location-protocol①
            link-type="idl"} setter. Furthermore, the non-failure
            termination earlier in this state is an intentional
            difference for defining that setter.

    [no scheme state]{#no-scheme-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If `base`{.variable} is null, or `base`{.variable} has an
            [opaque path](#url-opaque-path){#ref-for-url-opaque-path⑤
            link-type="dfn"} and [c](#c){#ref-for-c②④ link-type="dfn"}
            is not U+0023 (#),
            [missing-scheme-non-relative-URL](#missing-scheme-non-relative-url){#ref-for-missing-scheme-non-relative-url
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error③④
            link-type="dfn"}, return failure.

        2.  Otherwise, if `base`{.variable} has an [opaque
            path](#url-opaque-path){#ref-for-url-opaque-path⑥
            link-type="dfn"} and [c](#c){#ref-for-c②⑤ link-type="dfn"}
            is U+0023 (#), set `url`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②①
            link-type="dfn"} to `base`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②②
            link-type="dfn"}, `url`{.variable}'s
            [path](#concept-url-path){#ref-for-concept-url-path⑥
            link-type="dfn"} to `base`{.variable}'s
            [path](#concept-url-path){#ref-for-concept-url-path⑦
            link-type="dfn"}, `url`{.variable}'s
            [query](#concept-url-query){#ref-for-concept-url-query③
            link-type="dfn"} to `base`{.variable}'s
            [query](#concept-url-query){#ref-for-concept-url-query④
            link-type="dfn"}, `url`{.variable}'s
            [fragment](#concept-url-fragment){#ref-for-concept-url-fragment②
            link-type="dfn"} to the empty string, and set
            `state`{.variable} to [fragment
            state](#fragment-state){#ref-for-fragment-state
            link-type="dfn"}.

        3.  Otherwise, if `base`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②③
            link-type="dfn"} is not \"`file`\", set `state`{.variable}
            to [relative state](#relative-state){#ref-for-relative-state
            link-type="dfn"} and decrease `pointer`{.variable} by 1.

        4.  Otherwise, set `state`{.variable} to [file
            state](#file-state){#ref-for-file-state① link-type="dfn"}
            and decrease `pointer`{.variable} by 1.

    [special relative or authority state]{#special-relative-or-authority-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c②⑥ link-type="dfn"} is U+002F (/) and
            [remaining](#remaining){#ref-for-remaining⑥ link-type="dfn"}
            starts with U+002F (/), then set `state`{.variable} to
            [special authority ignore slashes
            state](#special-authority-ignore-slashes-state){#ref-for-special-authority-ignore-slashes-state
            link-type="dfn"} and increase `pointer`{.variable} by 1.

        2.  Otherwise,
            [special-scheme-missing-following-solidus](#special-scheme-missing-following-solidus){#ref-for-special-scheme-missing-following-solidus①
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error③⑤
            link-type="dfn"}, set `state`{.variable} to [relative
            state](#relative-state){#ref-for-relative-state①
            link-type="dfn"} and decrease `pointer`{.variable} by 1.

    [path or authority state]{#path-or-authority-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c②⑦ link-type="dfn"} is U+002F (/), then
            set `state`{.variable} to [authority
            state](#authority-state){#ref-for-authority-state
            link-type="dfn"}.

        2.  Otherwise, set `state`{.variable} to [path
            state](#path-state){#ref-for-path-state link-type="dfn"},
            and decrease `pointer`{.variable} by 1.

    [relative state]{#relative-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  Assert: `base`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②④
            link-type="dfn"} is not \"`file`\".

        2.  Set `url`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②⑤
            link-type="dfn"} to `base`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②⑥
            link-type="dfn"}.

        3.  If [c](#c){#ref-for-c②⑧ link-type="dfn"} is U+002F (/), then
            set `state`{.variable} to [relative slash
            state](#relative-slash-state){#ref-for-relative-slash-state
            link-type="dfn"}.

        4.  Otherwise, if `url`{.variable} [is
            special](#is-special){#ref-for-is-special⑤ link-type="dfn"}
            and [c](#c){#ref-for-c②⑨ link-type="dfn"} is U+005C (\\),
            [invalid-reverse-solidus](#invalid-reverse-solidus){#ref-for-invalid-reverse-solidus
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error③⑥
            link-type="dfn"}, set `state`{.variable} to [relative slash
            state](#relative-slash-state){#ref-for-relative-slash-state①
            link-type="dfn"}.

        5.  Otherwise:

            1.  Set `url`{.variable}'s
                [username](#concept-url-username){#ref-for-concept-url-username③
                link-type="dfn"} to `base`{.variable}'s
                [username](#concept-url-username){#ref-for-concept-url-username④
                link-type="dfn"}, `url`{.variable}'s
                [password](#concept-url-password){#ref-for-concept-url-password③
                link-type="dfn"} to `base`{.variable}'s
                [password](#concept-url-password){#ref-for-concept-url-password④
                link-type="dfn"}, `url`{.variable}'s
                [host](#concept-url-host){#ref-for-concept-url-host⑧
                link-type="dfn"} to `base`{.variable}'s
                [host](#concept-url-host){#ref-for-concept-url-host⑨
                link-type="dfn"}, `url`{.variable}'s
                [port](#concept-url-port){#ref-for-concept-url-port④
                link-type="dfn"} to `base`{.variable}'s
                [port](#concept-url-port){#ref-for-concept-url-port⑤
                link-type="dfn"}, `url`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path⑧
                link-type="dfn"} to a
                [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone
                link-type="dfn"} of `base`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path⑨
                link-type="dfn"}, and `url`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query⑤
                link-type="dfn"} to `base`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query⑥
                link-type="dfn"}.

            2.  If [c](#c){#ref-for-c③⓪ link-type="dfn"} is U+003F (?),
                then set `url`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query⑦
                link-type="dfn"} to the empty string, and
                `state`{.variable} to [query
                state](#query-state){#ref-for-query-state
                link-type="dfn"}.

            3.  Otherwise, if [c](#c){#ref-for-c③① link-type="dfn"} is
                U+0023 (#), set `url`{.variable}'s
                [fragment](#concept-url-fragment){#ref-for-concept-url-fragment③
                link-type="dfn"} to the empty string and
                `state`{.variable} to [fragment
                state](#fragment-state){#ref-for-fragment-state①
                link-type="dfn"}.

            4.  Otherwise, if [c](#c){#ref-for-c③② link-type="dfn"} is
                not the [EOF code
                point](#eof-code-point){#ref-for-eof-code-point⑧
                link-type="dfn"}:

                1.  Set `url`{.variable}'s
                    [query](#concept-url-query){#ref-for-concept-url-query⑧
                    link-type="dfn"} to null.

                2.  [Shorten](#shorten-a-urls-path){#ref-for-shorten-a-urls-path
                    link-type="dfn"} `url`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path①⓪
                    link-type="dfn"}.

                3.  Set `state`{.variable} to [path
                    state](#path-state){#ref-for-path-state①
                    link-type="dfn"} and decrease `pointer`{.variable}
                    by 1.

    [relative slash state]{#relative-slash-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If `url`{.variable} [is
            special](#is-special){#ref-for-is-special⑥ link-type="dfn"}
            and [c](#c){#ref-for-c③③ link-type="dfn"} is U+002F (/) or
            U+005C (\\), then:

            1.  If [c](#c){#ref-for-c③④ link-type="dfn"} is U+005C (\\),
                [invalid-reverse-solidus](#invalid-reverse-solidus){#ref-for-invalid-reverse-solidus①
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error③⑦
                link-type="dfn"}.

            2.  Set `state`{.variable} to [special authority ignore
                slashes
                state](#special-authority-ignore-slashes-state){#ref-for-special-authority-ignore-slashes-state①
                link-type="dfn"}.

        2.  Otherwise, if [c](#c){#ref-for-c③⑤ link-type="dfn"} is
            U+002F (/), then set `state`{.variable} to [authority
            state](#authority-state){#ref-for-authority-state①
            link-type="dfn"}.

        3.  Otherwise, set `url`{.variable}'s
            [username](#concept-url-username){#ref-for-concept-url-username⑤
            link-type="dfn"} to `base`{.variable}'s
            [username](#concept-url-username){#ref-for-concept-url-username⑥
            link-type="dfn"}, `url`{.variable}'s
            [password](#concept-url-password){#ref-for-concept-url-password⑤
            link-type="dfn"} to `base`{.variable}'s
            [password](#concept-url-password){#ref-for-concept-url-password⑥
            link-type="dfn"}, `url`{.variable}'s
            [host](#concept-url-host){#ref-for-concept-url-host①⓪
            link-type="dfn"} to `base`{.variable}'s
            [host](#concept-url-host){#ref-for-concept-url-host①①
            link-type="dfn"}, `url`{.variable}'s
            [port](#concept-url-port){#ref-for-concept-url-port⑥
            link-type="dfn"} to `base`{.variable}'s
            [port](#concept-url-port){#ref-for-concept-url-port⑦
            link-type="dfn"}, `state`{.variable} to [path
            state](#path-state){#ref-for-path-state② link-type="dfn"},
            and then, decrease `pointer`{.variable} by 1.

    [special authority slashes state]{#special-authority-slashes-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c③⑥ link-type="dfn"} is U+002F (/) and
            [remaining](#remaining){#ref-for-remaining⑦ link-type="dfn"}
            starts with U+002F (/), then set `state`{.variable} to
            [special authority ignore slashes
            state](#special-authority-ignore-slashes-state){#ref-for-special-authority-ignore-slashes-state②
            link-type="dfn"} and increase `pointer`{.variable} by 1.

        2.  Otherwise,
            [special-scheme-missing-following-solidus](#special-scheme-missing-following-solidus){#ref-for-special-scheme-missing-following-solidus②
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error③⑧
            link-type="dfn"}, set `state`{.variable} to [special
            authority ignore slashes
            state](#special-authority-ignore-slashes-state){#ref-for-special-authority-ignore-slashes-state③
            link-type="dfn"} and decrease `pointer`{.variable} by 1.

    [special authority ignore slashes state]{#special-authority-ignore-slashes-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c③⑦ link-type="dfn"} is neither U+002F
            (/) nor U+005C (\\), then set `state`{.variable} to
            [authority
            state](#authority-state){#ref-for-authority-state②
            link-type="dfn"} and decrease `pointer`{.variable} by 1.

        2.  Otherwise,
            [special-scheme-missing-following-solidus](#special-scheme-missing-following-solidus){#ref-for-special-scheme-missing-following-solidus③
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error③⑨
            link-type="dfn"}.

    [authority state]{#authority-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c③⑧ link-type="dfn"} is U+0040 (@),
            then:

            1.  [Invalid-credentials](#invalid-credentials){#ref-for-invalid-credentials
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error④⓪
                link-type="dfn"}.

            2.  If `atSignSeen`{.variable} is true, then prepend
                \"`%40`\" to `buffer`{.variable}.

            3.  Set `atSignSeen`{.variable} to true.

            4.  For each `codePoint`{.variable} in `buffer`{.variable}:

                1.  If `codePoint`{.variable} is U+003A (:) and
                    `passwordTokenSeen`{.variable} is false, then set
                    `passwordTokenSeen`{.variable} to true and
                    [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue④
                    link-type="dfn"}.

                2.  Let `encodedCodePoints`{.variable} be the result of
                    running [UTF-8
                    percent-encode](#utf-8-percent-encode){#ref-for-utf-8-percent-encode②
                    link-type="dfn"} `codePoint`{.variable} using the
                    [userinfo percent-encode
                    set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set⑥
                    link-type="dfn"}.

                3.  If `passwordTokenSeen`{.variable} is true, then
                    append `encodedCodePoints`{.variable} to
                    `url`{.variable}'s
                    [password](#concept-url-password){#ref-for-concept-url-password⑦
                    link-type="dfn"}.

                4.  Otherwise, append `encodedCodePoints`{.variable} to
                    `url`{.variable}'s
                    [username](#concept-url-username){#ref-for-concept-url-username⑦
                    link-type="dfn"}.

            5.  Set `buffer`{.variable} to the empty string.

        2.  Otherwise, if one of the following is true:

            - [c](#c){#ref-for-c③⑨ link-type="dfn"} is the [EOF code
              point](#eof-code-point){#ref-for-eof-code-point⑨
              link-type="dfn"}, U+002F (/), U+003F (?), or U+0023 (#)

            - `url`{.variable} [is
              special](#is-special){#ref-for-is-special⑦
              link-type="dfn"} and [c](#c){#ref-for-c④⓪ link-type="dfn"}
              is U+005C (\\)

            then:

            1.  If `atSignSeen`{.variable} is true and
                `buffer`{.variable} is the empty string,
                [host-missing](#host-missing){#ref-for-host-missing
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error④①
                link-type="dfn"}, return failure.

            2.  Decrease `pointer`{.variable} by `buffer`{.variable}'s
                [code point
                length](https://infra.spec.whatwg.org/#string-code-point-length){#ref-for-string-code-point-length①
                link-type="dfn"} + 1, set `buffer`{.variable} to the
                empty string, and set `state`{.variable} to [host
                state](#host-state){#ref-for-host-state
                link-type="dfn"}.

        3.  Otherwise, append [c](#c){#ref-for-c④① link-type="dfn"} to
            `buffer`{.variable}.

    [host state]{#host-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}\
    [hostname state]{#hostname-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If `state override`{.variable} is given and
            `url`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②⑦
            link-type="dfn"} is \"`file`\", then decrease
            `pointer`{.variable} by 1 and set `state`{.variable} to
            [file host state](#file-host-state){#ref-for-file-host-state
            link-type="dfn"}.

        2.  Otherwise, if [c](#c){#ref-for-c④② link-type="dfn"} is
            U+003A (:) and `insideBrackets`{.variable} is false:

            1.  If `buffer`{.variable} is the empty string,
                [host-missing](#host-missing){#ref-for-host-missing①
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error④②
                link-type="dfn"}, return failure.

            2.  If `state override`{.variable} is given and
                `state override`{.variable} is [hostname
                state](#hostname-state){#ref-for-hostname-state
                link-type="dfn"}, then return failure.

            3.  Let `host`{.variable} be the result of [host
                parsing](#concept-host-parser){#ref-for-concept-host-parser⑧
                link-type="dfn"} `buffer`{.variable} with
                `url`{.variable} [is not
                special](#is-not-special){#ref-for-is-not-special①
                link-type="dfn"}.

            4.  If `host`{.variable} is failure, then return failure.

            5.  Set `url`{.variable}'s
                [host](#concept-url-host){#ref-for-concept-url-host①②
                link-type="dfn"} to `host`{.variable},
                `buffer`{.variable} to the empty string, and
                `state`{.variable} to [port
                state](#port-state){#ref-for-port-state
                link-type="dfn"}.

        3.  Otherwise, if one of the following is true:

            - [c](#c){#ref-for-c④③ link-type="dfn"} is the [EOF code
              point](#eof-code-point){#ref-for-eof-code-point①⓪
              link-type="dfn"}, U+002F (/), U+003F (?), or U+0023 (#)

            - `url`{.variable} [is
              special](#is-special){#ref-for-is-special⑧
              link-type="dfn"} and [c](#c){#ref-for-c④④ link-type="dfn"}
              is U+005C (\\)

            then decrease `pointer`{.variable} by 1, and:

            1.  If `url`{.variable} [is
                special](#is-special){#ref-for-is-special⑨
                link-type="dfn"} and `buffer`{.variable} is the empty
                string,
                [host-missing](#host-missing){#ref-for-host-missing②
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error④③
                link-type="dfn"}, return failure.

            2.  Otherwise, if `state override`{.variable} is given,
                `buffer`{.variable} is the empty string, and either
                `url`{.variable} [includes
                credentials](#include-credentials){#ref-for-include-credentials②
                link-type="dfn"} or `url`{.variable}'s
                [port](#concept-url-port){#ref-for-concept-url-port⑧
                link-type="dfn"} is non-null, then return failure.

            3.  Let `host`{.variable} be the result of [host
                parsing](#concept-host-parser){#ref-for-concept-host-parser⑨
                link-type="dfn"} `buffer`{.variable} with
                `url`{.variable} [is not
                special](#is-not-special){#ref-for-is-not-special②
                link-type="dfn"}.

            4.  If `host`{.variable} is failure, then return failure.

            5.  Set `url`{.variable}'s
                [host](#concept-url-host){#ref-for-concept-url-host①③
                link-type="dfn"} to `host`{.variable},
                `buffer`{.variable} to the empty string, and
                `state`{.variable} to [path start
                state](#path-start-state){#ref-for-path-start-state
                link-type="dfn"}.

            6.  If `state override`{.variable} is given, then return.

        4.  Otherwise:

            1.  If [c](#c){#ref-for-c④⑤ link-type="dfn"} is U+005B (\[),
                then set `insideBrackets`{.variable} to true.

            2.  If [c](#c){#ref-for-c④⑥ link-type="dfn"} is U+005D (\]),
                then set `insideBrackets`{.variable} to false.

            3.  Append [c](#c){#ref-for-c④⑦ link-type="dfn"} to
                `buffer`{.variable}.

    [port state]{#port-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c④⑧ link-type="dfn"} is an [ASCII
            digit](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit⑦
            link-type="dfn"}, append [c](#c){#ref-for-c④⑨
            link-type="dfn"} to `buffer`{.variable}.

        2.  Otherwise, if one of the following is true:

            - [c](#c){#ref-for-c⑤⓪ link-type="dfn"} is the [EOF code
              point](#eof-code-point){#ref-for-eof-code-point①①
              link-type="dfn"}, U+002F (/), U+003F (?), or U+0023 (#);

            - `url`{.variable} [is
              special](#is-special){#ref-for-is-special①⓪
              link-type="dfn"} and [c](#c){#ref-for-c⑤① link-type="dfn"}
              is U+005C (\\); or

            - `state override`{.variable} is given,

            then:

            1.  If `buffer`{.variable} is not the empty string:

                1.  Let `port`{.variable} be the mathematical integer
                    value that is represented by `buffer`{.variable} in
                    radix-10 using [ASCII
                    digits](https://infra.spec.whatwg.org/#ascii-digit){#ref-for-ascii-digit⑧
                    link-type="dfn"} for digits with values 0 through 9.

                2.  If `port`{.variable} is not a [16-bit unsigned
                    integer](https://infra.spec.whatwg.org/#16-bit-unsigned-integer){#ref-for-16-bit-unsigned-integer③
                    link-type="dfn"},
                    [port-out-of-range](#port-out-of-range){#ref-for-port-out-of-range
                    link-type="dfn"} [validation
                    error](#validation-error){#ref-for-validation-error④④
                    link-type="dfn"}, return failure.

                3.  Set `url`{.variable}'s
                    [port](#concept-url-port){#ref-for-concept-url-port⑨
                    link-type="dfn"} to null, if `port`{.variable} is
                    `url`{.variable}'s
                    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②⑧
                    link-type="dfn"}'s [default
                    port](#default-port){#ref-for-default-port③
                    link-type="dfn"}; otherwise to `port`{.variable}.

                4.  Set `buffer`{.variable} to the empty string.

                5.  If `state override`{.variable} is given, then
                    return.

            2.  If `state override`{.variable} is given, then return
                failure.

            3.  Set `state`{.variable} to [path start
                state](#path-start-state){#ref-for-path-start-state①
                link-type="dfn"} and decrease `pointer`{.variable} by 1.

        3.  Otherwise,
            [port-invalid](#port-invalid){#ref-for-port-invalid
            link-type="dfn"} [validation
            error](#validation-error){#ref-for-validation-error④⑤
            link-type="dfn"}, return failure.

    [file state]{#file-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  Set `url`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme②⑨
            link-type="dfn"} to \"`file`\".

        2.  Set `url`{.variable}'s
            [host](#concept-url-host){#ref-for-concept-url-host①④
            link-type="dfn"} to the empty string.

        3.  If [c](#c){#ref-for-c⑤② link-type="dfn"} is U+002F (/) or
            U+005C (\\), then:

            1.  If [c](#c){#ref-for-c⑤③ link-type="dfn"} is U+005C (\\),
                [invalid-reverse-solidus](#invalid-reverse-solidus){#ref-for-invalid-reverse-solidus②
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error④⑥
                link-type="dfn"}.

            2.  Set `state`{.variable} to [file slash
                state](#file-slash-state){#ref-for-file-slash-state
                link-type="dfn"}.

        4.  Otherwise, if `base`{.variable} is non-null and
            `base`{.variable}'s
            [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③⓪
            link-type="dfn"} is \"`file`\":

            1.  Set `url`{.variable}'s
                [host](#concept-url-host){#ref-for-concept-url-host①⑤
                link-type="dfn"} to `base`{.variable}'s
                [host](#concept-url-host){#ref-for-concept-url-host①⑥
                link-type="dfn"}, `url`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path①①
                link-type="dfn"} to a
                [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone①
                link-type="dfn"} of `base`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path①②
                link-type="dfn"}, and `url`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query⑨
                link-type="dfn"} to `base`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query①⓪
                link-type="dfn"}.

            2.  If [c](#c){#ref-for-c⑤④ link-type="dfn"} is U+003F (?),
                then set `url`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query①①
                link-type="dfn"} to the empty string and
                `state`{.variable} to [query
                state](#query-state){#ref-for-query-state①
                link-type="dfn"}.

            3.  Otherwise, if [c](#c){#ref-for-c⑤⑤ link-type="dfn"} is
                U+0023 (#), set `url`{.variable}'s
                [fragment](#concept-url-fragment){#ref-for-concept-url-fragment④
                link-type="dfn"} to the empty string and
                `state`{.variable} to [fragment
                state](#fragment-state){#ref-for-fragment-state②
                link-type="dfn"}.

            4.  Otherwise, if [c](#c){#ref-for-c⑤⑥ link-type="dfn"} is
                not the [EOF code
                point](#eof-code-point){#ref-for-eof-code-point①②
                link-type="dfn"}:

                1.  Set `url`{.variable}'s
                    [query](#concept-url-query){#ref-for-concept-url-query①②
                    link-type="dfn"} to null.

                2.  If the [code point
                    substring](https://infra.spec.whatwg.org/#code-point-substring-to-the-end-of-the-string){#ref-for-code-point-substring-to-the-end-of-the-string①
                    link-type="dfn"} from `pointer`{.variable} to the
                    end of `input`{.variable} does not [start with a
                    Windows drive
                    letter](#start-with-a-windows-drive-letter){#ref-for-start-with-a-windows-drive-letter①
                    link-type="dfn"}, then
                    [shorten](#shorten-a-urls-path){#ref-for-shorten-a-urls-path①
                    link-type="dfn"} `url`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path①③
                    link-type="dfn"}.

                3.  Otherwise:

                    1.  [File-invalid-Windows-drive-letter](#file-invalid-windows-drive-letter){#ref-for-file-invalid-windows-drive-letter
                        link-type="dfn"} [validation
                        error](#validation-error){#ref-for-validation-error④⑦
                        link-type="dfn"}.

                    2.  Set `url`{.variable}'s
                        [path](#concept-url-path){#ref-for-concept-url-path①④
                        link-type="dfn"} to « ».

                    This is a (platform-independent) Windows drive
                    letter quirk.

                4.  Set `state`{.variable} to [path
                    state](#path-state){#ref-for-path-state③
                    link-type="dfn"} and decrease `pointer`{.variable}
                    by 1.

        5.  Otherwise, set `state`{.variable} to [path
            state](#path-state){#ref-for-path-state④ link-type="dfn"},
            and decrease `pointer`{.variable} by 1.

    [file slash state]{#file-slash-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c⑤⑦ link-type="dfn"} is U+002F (/) or
            U+005C (\\), then:

            1.  If [c](#c){#ref-for-c⑤⑧ link-type="dfn"} is U+005C (\\),
                [invalid-reverse-solidus](#invalid-reverse-solidus){#ref-for-invalid-reverse-solidus③
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error④⑧
                link-type="dfn"}.

            2.  Set `state`{.variable} to [file host
                state](#file-host-state){#ref-for-file-host-state①
                link-type="dfn"}.

        2.  Otherwise:

            1.  If `base`{.variable} is non-null and `base`{.variable}'s
                [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③①
                link-type="dfn"} is \"`file`\", then:

                1.  Set `url`{.variable}'s
                    [host](#concept-url-host){#ref-for-concept-url-host①⑦
                    link-type="dfn"} to `base`{.variable}'s
                    [host](#concept-url-host){#ref-for-concept-url-host①⑧
                    link-type="dfn"}.

                2.  If the [code point
                    substring](https://infra.spec.whatwg.org/#code-point-substring-to-the-end-of-the-string){#ref-for-code-point-substring-to-the-end-of-the-string②
                    link-type="dfn"} from `pointer`{.variable} to the
                    end of `input`{.variable} does not [start with a
                    Windows drive
                    letter](#start-with-a-windows-drive-letter){#ref-for-start-with-a-windows-drive-letter②
                    link-type="dfn"} and `base`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path①⑤
                    link-type="dfn"}\[0\] is a [normalized Windows drive
                    letter](#normalized-windows-drive-letter){#ref-for-normalized-windows-drive-letter②
                    link-type="dfn"}, then
                    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①
                    link-type="dfn"} `base`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path①⑥
                    link-type="dfn"}\[0\] to `url`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path①⑦
                    link-type="dfn"}.

                    This is a (platform-independent) Windows drive
                    letter quirk.

            2.  Set `state`{.variable} to [path
                state](#path-state){#ref-for-path-state⑤
                link-type="dfn"}, and decrease `pointer`{.variable} by
                1.

    [file host state]{#file-host-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c⑤⑨ link-type="dfn"} is the [EOF code
            point](#eof-code-point){#ref-for-eof-code-point①③
            link-type="dfn"}, U+002F (/), U+005C (\\), U+003F (?), or
            U+0023 (#), then decrease `pointer`{.variable} by 1 and
            then:

            1.  If `state override`{.variable} is not given and
                `buffer`{.variable} is a [Windows drive
                letter](#windows-drive-letter){#ref-for-windows-drive-letter③
                link-type="dfn"},
                [file-invalid-Windows-drive-letter-host](#file-invalid-windows-drive-letter-host){#ref-for-file-invalid-windows-drive-letter-host
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error④⑨
                link-type="dfn"}, set `state`{.variable} to [path
                state](#path-state){#ref-for-path-state⑥
                link-type="dfn"}.

                This is a (platform-independent) Windows drive letter
                quirk. `buffer`{.variable} is not reset here and instead
                used in the [path
                state](#path-state){#ref-for-path-state⑦
                link-type="dfn"}.

            2.  Otherwise, if `buffer`{.variable} is the empty string,
                then:

                1.  Set `url`{.variable}'s
                    [host](#concept-url-host){#ref-for-concept-url-host①⑨
                    link-type="dfn"} to the empty string.

                2.  If `state override`{.variable} is given, then
                    return.

                3.  Set `state`{.variable} to [path start
                    state](#path-start-state){#ref-for-path-start-state②
                    link-type="dfn"}.

            3.  Otherwise, run these steps:

                1.  Let `host`{.variable} be the result of [host
                    parsing](#concept-host-parser){#ref-for-concept-host-parser①⓪
                    link-type="dfn"} `buffer`{.variable} with
                    `url`{.variable} [is not
                    special](#is-not-special){#ref-for-is-not-special③
                    link-type="dfn"}.

                2.  If `host`{.variable} is failure, then return
                    failure.

                3.  If `host`{.variable} is \"`localhost`{title=""}\",
                    then set `host`{.variable} to the empty string.

                4.  Set `url`{.variable}'s
                    [host](#concept-url-host){#ref-for-concept-url-host②⓪
                    link-type="dfn"} to `host`{.variable}.

                5.  If `state override`{.variable} is given, then
                    return.

                6.  Set `buffer`{.variable} to the empty string and
                    `state`{.variable} to [path start
                    state](#path-start-state){#ref-for-path-start-state③
                    link-type="dfn"}.

        2.  Otherwise, append [c](#c){#ref-for-c⑥⓪ link-type="dfn"} to
            `buffer`{.variable}.

    [path start state]{#path-start-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If `url`{.variable} [is
            special](#is-special){#ref-for-is-special①①
            link-type="dfn"}, then:

            1.  If [c](#c){#ref-for-c⑥① link-type="dfn"} is U+005C (\\),
                [invalid-reverse-solidus](#invalid-reverse-solidus){#ref-for-invalid-reverse-solidus④
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤⓪
                link-type="dfn"}.

            2.  Set `state`{.variable} to [path
                state](#path-state){#ref-for-path-state⑧
                link-type="dfn"}.

            3.  If [c](#c){#ref-for-c⑥② link-type="dfn"} is neither
                U+002F (/) nor U+005C (\\), then decrease
                `pointer`{.variable} by 1.

        2.  Otherwise, if `state override`{.variable} is not given and
            [c](#c){#ref-for-c⑥③ link-type="dfn"} is U+003F (?), set
            `url`{.variable}'s
            [query](#concept-url-query){#ref-for-concept-url-query①③
            link-type="dfn"} to the empty string and `state`{.variable}
            to [query state](#query-state){#ref-for-query-state②
            link-type="dfn"}.

        3.  Otherwise, if `state override`{.variable} is not given and
            [c](#c){#ref-for-c⑥④ link-type="dfn"} is U+0023 (#), set
            `url`{.variable}'s
            [fragment](#concept-url-fragment){#ref-for-concept-url-fragment⑤
            link-type="dfn"} to the empty string and `state`{.variable}
            to [fragment
            state](#fragment-state){#ref-for-fragment-state③
            link-type="dfn"}.

        4.  Otherwise, if [c](#c){#ref-for-c⑥⑤ link-type="dfn"} is not
            the [EOF code
            point](#eof-code-point){#ref-for-eof-code-point①④
            link-type="dfn"}:

            1.  Set `state`{.variable} to [path
                state](#path-state){#ref-for-path-state⑨
                link-type="dfn"}.

            2.  If [c](#c){#ref-for-c⑥⑥ link-type="dfn"} is not U+002F
                (/), then decrease `pointer`{.variable} by 1.

        5.  Otherwise, if `state override`{.variable} is given and
            `url`{.variable}'s
            [host](#concept-url-host){#ref-for-concept-url-host②①
            link-type="dfn"} is null,
            [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append②
            link-type="dfn"} the empty string to `url`{.variable}'s
            [path](#concept-url-path){#ref-for-concept-url-path①⑧
            link-type="dfn"}.

    [path state]{#path-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If one of the following is true:

            - [c](#c){#ref-for-c⑥⑦ link-type="dfn"} is the [EOF code
              point](#eof-code-point){#ref-for-eof-code-point①⑤
              link-type="dfn"} or U+002F (/)

            - `url`{.variable} [is
              special](#is-special){#ref-for-is-special①②
              link-type="dfn"} and [c](#c){#ref-for-c⑥⑧ link-type="dfn"}
              is U+005C (\\)

            - `state override`{.variable} is not given and
              [c](#c){#ref-for-c⑥⑨ link-type="dfn"} is U+003F (?) or
              U+0023 (#)

            then:

            1.  If `url`{.variable} [is
                special](#is-special){#ref-for-is-special①③
                link-type="dfn"} and [c](#c){#ref-for-c⑦⓪
                link-type="dfn"} is U+005C (\\),
                [invalid-reverse-solidus](#invalid-reverse-solidus){#ref-for-invalid-reverse-solidus⑤
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤①
                link-type="dfn"}.

            2.  If `buffer`{.variable} is a [double-dot URL path
                segment](#double-dot-path-segment){#ref-for-double-dot-path-segment②
                link-type="dfn"}, then:

                1.  [Shorten](#shorten-a-urls-path){#ref-for-shorten-a-urls-path②
                    link-type="dfn"} `url`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path①⑨
                    link-type="dfn"}.

                2.  If neither [c](#c){#ref-for-c⑦① link-type="dfn"} is
                    U+002F (/), nor `url`{.variable} [is
                    special](#is-special){#ref-for-is-special①④
                    link-type="dfn"} and [c](#c){#ref-for-c⑦②
                    link-type="dfn"} is U+005C (\\),
                    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append③
                    link-type="dfn"} the empty string to
                    `url`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path②⓪
                    link-type="dfn"}.

                    This means that for input `/usr/..` the result is
                    `/` and not a lack of a path.

            3.  Otherwise, if `buffer`{.variable} is a [single-dot URL
                path
                segment](#single-dot-path-segment){#ref-for-single-dot-path-segment②
                link-type="dfn"} and if neither [c](#c){#ref-for-c⑦③
                link-type="dfn"} is U+002F (/), nor `url`{.variable} [is
                special](#is-special){#ref-for-is-special①⑤
                link-type="dfn"} and [c](#c){#ref-for-c⑦④
                link-type="dfn"} is U+005C (\\),
                [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append④
                link-type="dfn"} the empty string to `url`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path②①
                link-type="dfn"}.

            4.  Otherwise, if `buffer`{.variable} is not a [single-dot
                URL path
                segment](#single-dot-path-segment){#ref-for-single-dot-path-segment③
                link-type="dfn"}, then:

                1.  If `url`{.variable}'s
                    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③②
                    link-type="dfn"} is \"`file`\", `url`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path②②
                    link-type="dfn"} [is
                    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty
                    link-type="dfn"}, and `buffer`{.variable} is a
                    [Windows drive
                    letter](#windows-drive-letter){#ref-for-windows-drive-letter④
                    link-type="dfn"}, then replace the second code point
                    in `buffer`{.variable} with U+003A (:).

                    This is a (platform-independent) Windows drive
                    letter quirk.

                2.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑤
                    link-type="dfn"} `buffer`{.variable} to
                    `url`{.variable}'s
                    [path](#concept-url-path){#ref-for-concept-url-path②③
                    link-type="dfn"}.

            5.  Set `buffer`{.variable} to the empty string.

            6.  If [c](#c){#ref-for-c⑦⑤ link-type="dfn"} is U+003F (?),
                then set `url`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query①④
                link-type="dfn"} to the empty string and
                `state`{.variable} to [query
                state](#query-state){#ref-for-query-state③
                link-type="dfn"}.

            7.  If [c](#c){#ref-for-c⑦⑥ link-type="dfn"} is U+0023 (#),
                then set `url`{.variable}'s
                [fragment](#concept-url-fragment){#ref-for-concept-url-fragment⑥
                link-type="dfn"} to the empty string and
                `state`{.variable} to [fragment
                state](#fragment-state){#ref-for-fragment-state④
                link-type="dfn"}.

        2.  Otherwise, run these steps:

            1.  If [c](#c){#ref-for-c⑦⑦ link-type="dfn"} is not a [URL
                code point](#url-code-points){#ref-for-url-code-points③
                link-type="dfn"} and not U+0025 (%),
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit④
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤②
                link-type="dfn"}.

            2.  If [c](#c){#ref-for-c⑦⑧ link-type="dfn"} is U+0025 (%)
                and [remaining](#remaining){#ref-for-remaining⑧
                link-type="dfn"} does not start with two [ASCII hex
                digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit⑥
                link-type="dfn"},
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit⑤
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤③
                link-type="dfn"}.

            3.  [UTF-8
                percent-encode](#utf-8-percent-encode){#ref-for-utf-8-percent-encode③
                link-type="dfn"} [c](#c){#ref-for-c⑦⑨ link-type="dfn"}
                using the [path percent-encode
                set](#path-percent-encode-set){#ref-for-path-percent-encode-set①
                link-type="dfn"} and append the result to
                `buffer`{.variable}.

    [opaque path state]{#cannot-be-a-base-url-path-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c⑧⓪ link-type="dfn"} is U+003F (?), then
            set `url`{.variable}'s
            [query](#concept-url-query){#ref-for-concept-url-query①⑤
            link-type="dfn"} to the empty string and `state`{.variable}
            to [query state](#query-state){#ref-for-query-state④
            link-type="dfn"}.

        2.  Otherwise, if [c](#c){#ref-for-c⑧① link-type="dfn"} is
            U+0023 (#), then set `url`{.variable}'s
            [fragment](#concept-url-fragment){#ref-for-concept-url-fragment⑦
            link-type="dfn"} to the empty string and `state`{.variable}
            to [fragment
            state](#fragment-state){#ref-for-fragment-state⑤
            link-type="dfn"}.

        3.  Otherwise, if [c](#c){#ref-for-c⑧② link-type="dfn"} is
            U+0020 SPACE:

            1.  If [remaining](#remaining){#ref-for-remaining⑨
                link-type="dfn"} starts with U+003F (?) or U+003F (#),
                then append \"`%20`\" to `url`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path②④
                link-type="dfn"}.

            2.  Otherwise, append U+0020 SPACE to `url`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path②⑤
                link-type="dfn"}.

        4.  Otherwise, if [c](#c){#ref-for-c⑧③ link-type="dfn"} is not
            the [EOF code
            point](#eof-code-point){#ref-for-eof-code-point①⑥
            link-type="dfn"}:

            1.  If [c](#c){#ref-for-c⑧④ link-type="dfn"} is not a [URL
                code point](#url-code-points){#ref-for-url-code-points④
                link-type="dfn"} and not U+0025 (%),
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit⑥
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤④
                link-type="dfn"}.

            2.  If [c](#c){#ref-for-c⑧⑤ link-type="dfn"} is U+0025 (%)
                and [remaining](#remaining){#ref-for-remaining①⓪
                link-type="dfn"} does not start with two [ASCII hex
                digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit⑦
                link-type="dfn"},
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit⑦
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤⑤
                link-type="dfn"}.

            3.  [UTF-8
                percent-encode](#utf-8-percent-encode){#ref-for-utf-8-percent-encode④
                link-type="dfn"} [c](#c){#ref-for-c⑧⑥ link-type="dfn"}
                using the [C0 control percent-encode
                set](#c0-control-percent-encode-set){#ref-for-c0-control-percent-encode-set③
                link-type="dfn"} and append the result to
                `url`{.variable}'s
                [path](#concept-url-path){#ref-for-concept-url-path②⑥
                link-type="dfn"}.

    [query state]{#query-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If `encoding`{.variable} is not
            [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8⑥
            link-type="dfn"} and one of the following is true:

            - `url`{.variable} [is not
              special](#is-not-special){#ref-for-is-not-special④
              link-type="dfn"}

            - `url`{.variable}'s
              [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③③
              link-type="dfn"} is \"`ws`\" or \"`wss`\"

            then set `encoding`{.variable} to
            [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8⑦
            link-type="dfn"}.

        2.  If one of the following is true:

            - `state override`{.variable} is not given and
              [c](#c){#ref-for-c⑧⑦ link-type="dfn"} is U+0023 (#)

            - [c](#c){#ref-for-c⑧⑧ link-type="dfn"} is the [EOF code
              point](#eof-code-point){#ref-for-eof-code-point①⑦
              link-type="dfn"}

            then:

            1.  Let `queryPercentEncodeSet`{.variable} be the
                [special-query percent-encode
                set](#special-query-percent-encode-set){#ref-for-special-query-percent-encode-set
                link-type="dfn"} if `url`{.variable} [is
                special](#is-special){#ref-for-is-special①⑥
                link-type="dfn"}; otherwise the [query percent-encode
                set](#query-percent-encode-set){#ref-for-query-percent-encode-set③
                link-type="dfn"}.

            2.  [Percent-encode after
                encoding](#string-percent-encode-after-encoding){#ref-for-string-percent-encode-after-encoding⑤
                link-type="dfn"}, with `encoding`{.variable},
                `buffer`{.variable}, and
                `queryPercentEncodeSet`{.variable}, and append the
                result to `url`{.variable}'s
                [query](#concept-url-query){#ref-for-concept-url-query①⑥
                link-type="dfn"}.

                This operation cannot be invoked
                code-point-for-code-point due to the stateful
                [ISO-2022-JP
                encoder](https://encoding.spec.whatwg.org/#iso-2022-jp-encoder){#ref-for-iso-2022-jp-encoder
                link-type="dfn"}.

            3.  Set `buffer`{.variable} to the empty string.

            4.  If [c](#c){#ref-for-c⑧⑨ link-type="dfn"} is U+0023 (#),
                then set `url`{.variable}'s
                [fragment](#concept-url-fragment){#ref-for-concept-url-fragment⑧
                link-type="dfn"} to the empty string and state to
                [fragment
                state](#fragment-state){#ref-for-fragment-state⑥
                link-type="dfn"}.

        3.  Otherwise, if [c](#c){#ref-for-c⑨⓪ link-type="dfn"} is not
            the [EOF code
            point](#eof-code-point){#ref-for-eof-code-point①⑧
            link-type="dfn"}:

            1.  If [c](#c){#ref-for-c⑨① link-type="dfn"} is not a [URL
                code point](#url-code-points){#ref-for-url-code-points⑤
                link-type="dfn"} and not U+0025 (%),
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit⑧
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤⑥
                link-type="dfn"}.

            2.  If [c](#c){#ref-for-c⑨② link-type="dfn"} is U+0025 (%)
                and [remaining](#remaining){#ref-for-remaining①①
                link-type="dfn"} does not start with two [ASCII hex
                digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit⑧
                link-type="dfn"},
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit⑨
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤⑦
                link-type="dfn"}.

            3.  Append [c](#c){#ref-for-c⑨③ link-type="dfn"} to
                `buffer`{.variable}.

    [fragment state]{#fragment-state .dfn .dfn-paneled dfn-for="basic URL parser" dfn-type="dfn" export=""}

    :   1.  If [c](#c){#ref-for-c⑨④ link-type="dfn"} is not the [EOF
            code point](#eof-code-point){#ref-for-eof-code-point①⑨
            link-type="dfn"}, then:

            1.  If [c](#c){#ref-for-c⑨⑤ link-type="dfn"} is not a [URL
                code point](#url-code-points){#ref-for-url-code-points⑥
                link-type="dfn"} and not U+0025 (%),
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit①⓪
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤⑧
                link-type="dfn"}.

            2.  If [c](#c){#ref-for-c⑨⑥ link-type="dfn"} is U+0025 (%)
                and [remaining](#remaining){#ref-for-remaining①②
                link-type="dfn"} does not start with two [ASCII hex
                digits](https://infra.spec.whatwg.org/#ascii-hex-digit){#ref-for-ascii-hex-digit⑨
                link-type="dfn"},
                [invalid-URL-unit](#invalid-url-unit){#ref-for-invalid-url-unit①①
                link-type="dfn"} [validation
                error](#validation-error){#ref-for-validation-error⑤⑨
                link-type="dfn"}.

            3.  [UTF-8
                percent-encode](#utf-8-percent-encode){#ref-for-utf-8-percent-encode⑤
                link-type="dfn"} [c](#c){#ref-for-c⑨⑦ link-type="dfn"}
                using the [fragment percent-encode
                set](#fragment-percent-encode-set){#ref-for-fragment-percent-encode-set①
                link-type="dfn"} and append the result to
                `url`{.variable}'s
                [fragment](#concept-url-fragment){#ref-for-concept-url-fragment⑨
                link-type="dfn"}.

10. Return `url`{.variable}.
::::

------------------------------------------------------------------------

::: {.algorithm algorithm="set the username" algorithm-for="url"}
To [set the username]{#set-the-username .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} given a `url`{.variable} and
`username`{.variable}, set `url`{.variable}'s
[username](#concept-url-username){#ref-for-concept-url-username⑧
link-type="dfn"} to the result of running [UTF-8
percent-encode](#string-utf-8-percent-encode){#ref-for-string-utf-8-percent-encode③
link-type="dfn"} on `username`{.variable} using the [userinfo
percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set⑦
link-type="dfn"}.
:::

::: {.algorithm algorithm="set the password" algorithm-for="url"}
To [set the password]{#set-the-password .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} given a `url`{.variable} and
`password`{.variable}, set `url`{.variable}'s
[password](#concept-url-password){#ref-for-concept-url-password⑧
link-type="dfn"} to the result of running [UTF-8
percent-encode](#string-utf-8-percent-encode){#ref-for-string-utf-8-percent-encode④
link-type="dfn"} on `password`{.variable} using the [userinfo
percent-encode
set](#userinfo-percent-encode-set){#ref-for-userinfo-percent-encode-set⑧
link-type="dfn"}.
:::

### [4.5. ]{.secno}[URL serializing]{.content}[](#url-serializing){.self-link} {#url-serializing .heading .settled level="4.5"}

::: {.algorithm algorithm="URL serializer"}
The [URL serializer]{#concept-url-serializer .dfn .dfn-paneled
dfn-type="dfn" export=""} takes a
[URL](#concept-url){#ref-for-concept-url④④ link-type="dfn"}
`url`{.variable}, with an optional boolean
[`exclude fragment`{.variable}]{#url-serializer-exclude-fragment .dfn
.dfn-paneled dfn-for="URL serializer" dfn-type="dfn" export=""} (default
false), and then runs these steps. They return an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string②①
link-type="dfn"}.

1.  Let `output`{.variable} be `url`{.variable}'s
    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③④
    link-type="dfn"} and U+003A (:) concatenated.

2.  If `url`{.variable}'s
    [host](#concept-url-host){#ref-for-concept-url-host②②
    link-type="dfn"} is non-null:

    1.  Append \"`//`\" to `output`{.variable}.

    2.  If `url`{.variable} [includes
        credentials](#include-credentials){#ref-for-include-credentials③
        link-type="dfn"}, then:

        1.  Append `url`{.variable}'s
            [username](#concept-url-username){#ref-for-concept-url-username⑨
            link-type="dfn"} to `output`{.variable}.

        2.  If `url`{.variable}'s
            [password](#concept-url-password){#ref-for-concept-url-password⑨
            link-type="dfn"} is not the empty string, then append U+003A
            (:), followed by `url`{.variable}'s
            [password](#concept-url-password){#ref-for-concept-url-password①⓪
            link-type="dfn"}, to `output`{.variable}.

        3.  Append U+0040 (@) to `output`{.variable}.

    3.  Append `url`{.variable}'s
        [host](#concept-url-host){#ref-for-concept-url-host②③
        link-type="dfn"},
        [serialized](#concept-host-serializer){#ref-for-concept-host-serializer④
        link-type="dfn"}, to `output`{.variable}.

    4.  If `url`{.variable}'s
        [port](#concept-url-port){#ref-for-concept-url-port①⓪
        link-type="dfn"} is non-null, append U+003A (:) followed by
        `url`{.variable}'s
        [port](#concept-url-port){#ref-for-concept-url-port①①
        link-type="dfn"},
        [serialized](#serialize-an-integer){#ref-for-serialize-an-integer①
        link-type="dfn"}, to `output`{.variable}.

3.  If `url`{.variable}'s
    [host](#concept-url-host){#ref-for-concept-url-host②④
    link-type="dfn"} is null, `url`{.variable} does not have an [opaque
    path](#url-opaque-path){#ref-for-url-opaque-path⑦ link-type="dfn"},
    `url`{.variable}'s
    [path](#concept-url-path){#ref-for-concept-url-path②⑦
    link-type="dfn"}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑤
    link-type="dfn"} is greater than 1, and `url`{.variable}'s
    [path](#concept-url-path){#ref-for-concept-url-path②⑧
    link-type="dfn"}\[0\] is the empty string, then append U+002F (/)
    followed by U+002E (.) to `output`{.variable}.

    This prevents `web+demo:/.//not-a-host/` or
    `web+demo:/path/..//not-a-host/`, when
    [parsed](#concept-url-parser){#ref-for-concept-url-parser①⓪
    link-type="dfn"} and then
    [serialized](#concept-url-serializer){#ref-for-concept-url-serializer⑤
    link-type="dfn"}, from ending up as `web+demo://not-a-host/` (they
    end up as `web+demo:/.//not-a-host/`).

4.  Append the result of [URL path
    serializing](#url-path-serializer){#ref-for-url-path-serializer
    link-type="dfn"} `url`{.variable} to `output`{.variable}.

5.  If `url`{.variable}'s
    [query](#concept-url-query){#ref-for-concept-url-query①⑦
    link-type="dfn"} is non-null, append U+003F (?), followed by
    `url`{.variable}'s
    [query](#concept-url-query){#ref-for-concept-url-query①⑧
    link-type="dfn"}, to `output`{.variable}.

6.  If `exclude fragment`{.variable} is false and `url`{.variable}'s
    [fragment](#concept-url-fragment){#ref-for-concept-url-fragment①⓪
    link-type="dfn"} is non-null, then append U+0023 (#), followed by
    `url`{.variable}'s
    [fragment](#concept-url-fragment){#ref-for-concept-url-fragment①①
    link-type="dfn"}, to `output`{.variable}.

7.  Return `output`{.variable}.
:::

::: {.algorithm algorithm="URL path serializer"}
The [URL path serializer]{#url-path-serializer .dfn .dfn-paneled
dfn-type="dfn" export="" lt="URL path serializer|URL path serializing"}
takes a [URL](#concept-url){#ref-for-concept-url④⑤ link-type="dfn"}
`url`{.variable} and then runs these steps. They return an [ASCII
string](https://infra.spec.whatwg.org/#ascii-string){#ref-for-ascii-string②②
link-type="dfn"}.

1.  If `url`{.variable} has an [opaque
    path](#url-opaque-path){#ref-for-url-opaque-path⑧ link-type="dfn"},
    then return `url`{.variable}'s
    [path](#concept-url-path){#ref-for-concept-url-path②⑨
    link-type="dfn"}.

2.  Let `output`{.variable} be the empty string.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑤
    link-type="dfn"} `segment`{.variable} of `url`{.variable}'s
    [path](#concept-url-path){#ref-for-concept-url-path③⓪
    link-type="dfn"}: append U+002F (/) followed by `segment`{.variable}
    to `output`{.variable}.

4.  Return `output`{.variable}.
:::

### [4.6. ]{.secno}[URL equivalence]{.content}[](#url-equivalence){.self-link} {#url-equivalence .heading .settled level="4.6"}

::: {.algorithm algorithm="equal"}
To determine whether a [URL](#concept-url){#ref-for-concept-url④⑥
link-type="dfn"} `A`{.variable} [equals]{#concept-url-equals .dfn
.dfn-paneled dfn-for="url" dfn-type="dfn" export="" lt="equal"}
[URL](#concept-url){#ref-for-concept-url④⑦ link-type="dfn"}
`B`{.variable}, with an optional boolean
[`exclude fragments`{.variable}]{#url-equals-exclude-fragments .dfn
.dfn-paneled dfn-for="url/equals" dfn-type="dfn" export=""} (default
false), run these steps:

1.  Let `serializedA`{.variable} be the result of
    [serializing](#concept-url-serializer){#ref-for-concept-url-serializer⑥
    link-type="dfn"} `A`{.variable}, with [*exclude
    fragment*](#url-serializer-exclude-fragment){#ref-for-url-serializer-exclude-fragment
    link-type="dfn"} set to `exclude fragments`{.variable}.

2.  Let `serializedB`{.variable} be the result of
    [serializing](#concept-url-serializer){#ref-for-concept-url-serializer⑦
    link-type="dfn"} `B`{.variable}, with [*exclude
    fragment*](#url-serializer-exclude-fragment){#ref-for-url-serializer-exclude-fragment①
    link-type="dfn"} set to `exclude fragments`{.variable}.

3.  Return true if `serializedA`{.variable} is `serializedB`{.variable};
    otherwise false.
:::

### [4.7. ]{.secno}[Origin]{.content}[](#origin){.self-link} {#origin .heading .settled level="4.7"}

See
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin①
link-type="dfn"}'s definition in HTML for the necessary background
information.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

::: {.algorithm algorithm="origin" algorithm-for="url"}
The [origin]{#concept-url-origin .dfn .dfn-paneled dfn-for="url"
dfn-type="dfn" export=""} of a
[URL](#concept-url){#ref-for-concept-url④⑧ link-type="dfn"}
`url`{.variable} is the
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin②
link-type="dfn"} returned by running these steps, switching on
`url`{.variable}'s
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme③⑤
link-type="dfn"}:

\"`blob`\"

:   1.  If `url`{.variable}'s [blob URL
        entry](#concept-url-blob-entry){#ref-for-concept-url-blob-entry②
        link-type="dfn"} is non-null, then return `url`{.variable}'s
        [blob URL
        entry](#concept-url-blob-entry){#ref-for-concept-url-blob-entry③
        link-type="dfn"}'s
        [environment](https://w3c.github.io/FileAPI/#blob-url-entry-environment){#ref-for-blob-url-entry-environment
        link-type="dfn"}'s
        [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin
        link-type="dfn"}.

    2.  Let `pathURL`{.variable} be the result of
        [parsing](#concept-basic-url-parser){#ref-for-concept-basic-url-parser③
        link-type="dfn"} the result of [URL path
        serializing](#url-path-serializer){#ref-for-url-path-serializer①
        link-type="dfn"} `url`{.variable}.

    3.  If `pathURL`{.variable} is failure, then return a new [opaque
        origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-opaque){#ref-for-concept-origin-opaque
        link-type="dfn"}.

    4.  If `pathURL`{.variable}'s
        [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③⑥
        link-type="dfn"} is \"`http`\", \"`https`\", or \"`file`\", then
        return `pathURL`{.variable}'s
        [origin](#concept-url-origin){#ref-for-concept-url-origin
        link-type="dfn"}.

    5.  Return a new [opaque
        origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-opaque){#ref-for-concept-origin-opaque①
        link-type="dfn"}.

    [](#example-43b5cea5){.self-link}The
    [origin](#concept-url-origin){#ref-for-concept-url-origin①
    link-type="dfn"} of
    `blob:https://whatwg.org/d0360e2f-caee-469f-9a2f-87d5b0456f6f` is
    the [tuple
    origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-tuple){#ref-for-concept-origin-tuple
    link-type="dfn"} (\"`https`\", \"`whatwg.org`\", null, null).

\"`ftp`\"\
\"`http`\"\
\"`https`\"\
\"`ws`\"\
\"`wss`\"

:   Return the [tuple
    origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-tuple){#ref-for-concept-origin-tuple①
    link-type="dfn"} (`url`{.variable}'s
    [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③⑦
    link-type="dfn"}, `url`{.variable}'s
    [host](#concept-url-host){#ref-for-concept-url-host②⑤
    link-type="dfn"}, `url`{.variable}'s
    [port](#concept-url-port){#ref-for-concept-url-port①②
    link-type="dfn"}, null).

\"`file`\"

:   Unfortunate as it is, this is left as an exercise to the reader.
    When in doubt, return a new [opaque
    origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-opaque){#ref-for-concept-origin-opaque②
    link-type="dfn"}.

Otherwise

:   Return a new [opaque
    origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-opaque){#ref-for-concept-origin-opaque③
    link-type="dfn"}.

    This does indeed mean that these
    [URLs](#concept-url){#ref-for-concept-url④⑨ link-type="dfn"} cannot
    be [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin
    link-type="dfn"} with themselves.
:::

### [4.8. ]{.secno}[URL rendering]{.content}[](#url-rendering){.self-link} {#url-rendering .heading .settled level="4.8"}

A [URL](#concept-url){#ref-for-concept-url⑤⓪ link-type="dfn"} should be
rendered in its
[serialized](#concept-url-serializer){#ref-for-concept-url-serializer⑧
link-type="dfn"} form, with modifications described below, when the
primary purpose of displaying a URL is to have the user make a security
or trust decision. For example, users are expected to make trust
decisions based on a URL rendered in the browser address bar.

#### [4.8.1. ]{.secno}[Simplify non-human-readable or irrelevant components]{.content}[](#url-rendering-simplification){.self-link} {#url-rendering-simplification .heading .settled level="4.8.1"}

Remove components that can provide opportunities for spoofing or
distract from security-relevant information:

- Browsers may render only a URL's
  [host](#concept-url-host){#ref-for-concept-url-host②⑥ link-type="dfn"}
  in places where it is important for end users to distinguish between
  the host and other parts of the URL such as the
  [path](#concept-url-path){#ref-for-concept-url-path③①
  link-type="dfn"}. Browsers may consider simplifying the host further
  to draw attention to its [registrable
  domain](#host-registrable-domain){#ref-for-host-registrable-domain①
  link-type="dfn"}. For example, browsers may omit a leading `www` or
  `m` [domain label](#domain-label){#ref-for-domain-label①
  link-type="dfn"} to simplify the host, or display its registrable
  domain only to remove spoofing opportunities posted by subdomains
  (e.g., `https://examplecorp.attacker.com/`).

- Browsers should not render a
  [URL](#concept-url){#ref-for-concept-url⑤① link-type="dfn"}'s
  [username](#concept-url-username){#ref-for-concept-url-username①⓪
  link-type="dfn"} and
  [password](#concept-url-password){#ref-for-concept-url-password①①
  link-type="dfn"}, as they can be mistaken for a
  [URL](#concept-url){#ref-for-concept-url⑤② link-type="dfn"}'s
  [host](#concept-url-host){#ref-for-concept-url-host②⑦ link-type="dfn"}
  (e.g., `https://examplecorp.com@attacker.example/`).

- Browsers may render a URL without its
  [scheme](#concept-url-scheme){#ref-for-concept-url-scheme③⑧
  link-type="dfn"} if the display surface only ever permits a single
  scheme (such as a browser feature that omits `https://` because it is
  only enabled for secure origins). Otherwise, the scheme may be
  replaced or supplemented with a human-readable string (e.g., \"Not
  secure\"), a security indicator icon, or both.

#### [4.8.2. ]{.secno}[Elision]{.content}[](#url-rendering-elision){.self-link} {#url-rendering-elision .heading .settled level="4.8.2"}

In a space-constrained display, URLs should be elided carefully to avoid
misleading the user when making a security decision:

- Browsers should ensure that at least the [registrable
  domain](#host-registrable-domain){#ref-for-host-registrable-domain②
  link-type="dfn"} can be shown when the URL is rendered (to avoid
  showing, e.g., `...examplecorp.com` when loading
  `https://not-really-examplecorp.com/`).

- When the full [host](#concept-url-host){#ref-for-concept-url-host②⑧
  link-type="dfn"} cannot be rendered, browsers should elide [domain
  labels](#domain-label){#ref-for-domain-label② link-type="dfn"}
  starting from the lowest-level domain label. For example,
  `examplecorp.com.evil.com` should be elided as `...com.evil.com`, not
  `examplecorp.com...`. (Note that bidirectional text means that the
  lowest-level domain label may not appear on the left.)

#### [4.8.3. ]{.secno}[Internationalization and special characters]{.content}[](#url-rendering-i18n){.self-link} {#url-rendering-i18n .heading .settled level="4.8.3"}

Internationalized domain names (IDNs), special characters, and
bidirectional text should be handled with care to prevent spoofing:

- Browsers should render a [URL](#concept-url){#ref-for-concept-url⑤③
  link-type="dfn"}'s
  [host](#concept-url-host){#ref-for-concept-url-host②⑨ link-type="dfn"}
  by running [domain to
  Unicode](#concept-domain-to-unicode){#ref-for-concept-domain-to-unicode
  link-type="dfn"} with the [URL](#concept-url){#ref-for-concept-url⑤④
  link-type="dfn"}'s
  [host](#concept-url-host){#ref-for-concept-url-host③⓪ link-type="dfn"}
  and false.

  Various characters can be used in homograph spoofing attacks. Consider
  detecting confusable characters and warning when they are in use.
  [\[IDNFAQ\]](#biblio-idnfaq "Internationalized Domain Names (IDN) FAQ"){link-type="biblio"}
  [\[UTS39\]](#biblio-uts39 "Unicode Security Mechanisms"){link-type="biblio"}

- URLs are particularly prone to confusion between host and path when
  they contain bidirectional text, so in this case it is particularly
  advisable to only render a URL's
  [host](#concept-url-host){#ref-for-concept-url-host③①
  link-type="dfn"}. For readability, other parts of the
  [URL](#concept-url){#ref-for-concept-url⑤⑤ link-type="dfn"}, if
  rendered, should have their sequences of [percent-encoded
  bytes](#percent-encoded-byte){#ref-for-percent-encoded-byte⑦
  link-type="dfn"} replaced with code points resulting from running
  [UTF-8 decode without
  BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#ref-for-utf-8-decode-without-bom②
  link-type="dfn"} on the
  [percent-decoding](#string-percent-decode){#ref-for-string-percent-decode⑤
  link-type="dfn"} of those sequences, unless that renders those
  sequences invisible. Browsers may choose to not decode certain
  sequences that present spoofing risks (e.g., U+1F512 (🔒)).

- Browsers should render bidirectional text as if it were in a
  left-to-right embedding.
  [\[BIDI\]](#biblio-bidi "Unicode Bidirectional Algorithm"){link-type="biblio"}

  Unfortunately, as rendered [URLs](#concept-url){#ref-for-concept-url⑤⑥
  link-type="dfn"} are strings and can appear anywhere, a specific
  bidirectional algorithm for rendered
  [URLs](#concept-url){#ref-for-concept-url⑤⑦ link-type="dfn"} would not
  see wide adoption. Bidirectional text interacts with the parts of a
  [URL](#concept-url){#ref-for-concept-url⑤⑧ link-type="dfn"} in ways
  that can cause the rendering to be different from the model. Users of
  bidirectional languages can come to expect this, particularly in plain
  text environments.

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

## [6. ]{.secno}[API]{.content}[](#api){.self-link} {#api .heading .settled level="6"}

This section uses terminology from Web IDL. Browser user agents must
support this API. JavaScript implementations should support this API.
Other user agents or programming languages are encouraged to use an API
suitable to their needs, which might not be this one.
[\[WEBIDL\]](#biblio-webidl "Web IDL Standard"){link-type="biblio"}

### [6.1. ]{.secno}[URL class]{.content}[](#url-class){.self-link} {#url-class .heading .settled level="6.1"}

``` {.idl .highlight .def}
[Exposed=*,
 LegacyWindowAlias=webkitURL]
interface URL {
  constructor(USVString url, optional USVString base);

  static URL? parse(USVString url, optional USVString base);
  static boolean canParse(USVString url, optional USVString base);

  stringifier attribute USVString href;
  readonly attribute USVString origin;
           attribute USVString protocol;
           attribute USVString username;
           attribute USVString password;
           attribute USVString host;
           attribute USVString hostname;
           attribute USVString port;
           attribute USVString pathname;
           attribute USVString search;
  [SameObject] readonly attribute URLSearchParams searchParams;
           attribute USVString hash;

  USVString toJSON();
};
```

A [`URL`{.idl}](#url){#ref-for-url② link-type="idl"} object has an
associated:

- [URL]{#concept-url-url .dfn .dfn-paneled dfn-for="URL" dfn-type="dfn"
  noexport=""}: a [URL](#concept-url){#ref-for-concept-url⑤⑨
  link-type="dfn"}.
- [query object]{#concept-url-query-object .dfn .dfn-paneled
  dfn-for="URL" dfn-type="dfn" noexport=""}: a
  [`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams①
  link-type="idl"} object.

::: {.algorithm algorithm="API URL parser"}
The [API URL parser]{#api-url-parser .dfn .dfn-paneled dfn-type="dfn"
noexport=""} takes a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string①②
link-type="dfn"} `url`{.variable} and an optional null-or-[scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string①③
link-type="dfn"} `base`{.variable} (default null), and then runs these
steps:

1.  Let `parsedBase`{.variable} be null.

2.  If `base`{.variable} is non-null:

    1.  Set `parsedBase`{.variable} to the result of running the [basic
        URL
        parser](#concept-basic-url-parser){#ref-for-concept-basic-url-parser④
        link-type="dfn"} on `base`{.variable}.

    2.  If `parsedBase`{.variable} is failure, then return failure.

3.  Return the result of running the [basic URL
    parser](#concept-basic-url-parser){#ref-for-concept-basic-url-parser⑤
    link-type="dfn"} on `url`{.variable} with `parsedBase`{.variable}.
:::

::: {.algorithm algorithm="initialize" algorithm-for="URL"}
To [initialize]{#url-initialize .dfn .dfn-paneled dfn-for="URL"
dfn-type="dfn" noexport=""} a [`URL`{.idl}](#url){#ref-for-url③
link-type="idl"} object `url`{.variable} with a
[URL](#concept-url){#ref-for-concept-url⑥⓪ link-type="dfn"}
`urlRecord`{.variable}:

1.  Let `query`{.variable} be `urlRecord`{.variable}'s
    [query](#concept-url-query){#ref-for-concept-url-query①⑨
    link-type="dfn"}, if that is non-null; otherwise the empty string.

2.  Set `url`{.variable}'s
    [URL](#concept-url-url){#ref-for-concept-url-url link-type="dfn"} to
    `urlRecord`{.variable}.

3.  Set `url`{.variable}'s [query
    object](#concept-url-query-object){#ref-for-concept-url-query-object
    link-type="dfn"} to a new
    [`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams②
    link-type="idl"} object.

4.  [Initialize](#urlsearchparams-initialize){#ref-for-urlsearchparams-initialize
    link-type="dfn"} `url`{.variable}'s [query
    object](#concept-url-query-object){#ref-for-concept-url-query-object①
    link-type="dfn"} with `query`{.variable}.

5.  Set `url`{.variable}'s [query
    object](#concept-url-query-object){#ref-for-concept-url-query-object②
    link-type="dfn"}'s [URL
    object](#concept-urlsearchparams-url-object){#ref-for-concept-urlsearchparams-url-object
    link-type="dfn"} to `url`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="URL(url, base)" algorithm-for="URL"}
The [`new URL(``url`{.variable}`, ``base`{.variable}`)`]{#dom-url-url
.dfn .dfn-paneled .idl-code dfn-for="URL" dfn-type="constructor"
export=""
lt="URL(url, base)|constructor(url, base)|URL(url)|constructor(url)"}
constructor steps are:

1.  Let `parsedURL`{.variable} be the result of running the [API URL
    parser](#api-url-parser){#ref-for-api-url-parser link-type="dfn"} on
    `url`{.variable} with `base`{.variable}, if given.

2.  If `parsedURL`{.variable} is failure, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror
    link-type="idl"}.

3.  [Initialize](#url-initialize){#ref-for-url-initialize
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this
    link-type="dfn"} with `parsedURL`{.variable}.
:::

::: {#example-5434421b .example}
[](#example-5434421b){.self-link}

To [parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser⑥
link-type="dfn"} a string into a
[URL](#concept-url){#ref-for-concept-url⑥① link-type="dfn"} without
using a [base URL](#concept-base-url){#ref-for-concept-base-url①③
link-type="dfn"}, invoke the [`URL`{.idl}](#url){#ref-for-url④
link-type="idl"} constructor with a single argument:

``` {.lang-javascript .highlight}
var input = "https://example.org/💩",
    url = new URL(input)
url.pathname // "/%F0%9F%92%A9"
```

This throws an exception if the input is a [relative-URL
string](#relative-url-string){#ref-for-relative-url-string⑤
link-type="dfn"}:

``` {.lang-javascript .highlight}
try {
  var url = new URL("/🍣🍺")
} catch(e) {
  // that happened
}
```

For those cases a [base
URL](#concept-base-url){#ref-for-concept-base-url①④ link-type="dfn"} is
necessary:

``` {.lang-javascript .highlight}
var input = "/🍣🍺",
    url = new URL(input, document.baseURI)
url.href // "https://url.spec.whatwg.org/%F0%9F%8D%A3%F0%9F%8D%BA"
```

A [`URL`{.idl}](#url){#ref-for-url⑤ link-type="idl"} object can be used
as a [base URL](#concept-base-url){#ref-for-concept-base-url①⑤
link-type="dfn"} (as the IDL requires a string as argument, a
[`URL`{.idl}](#url){#ref-for-url⑥ link-type="idl"} object stringifies to
its [`href`{.idl}](#dom-url-href){#ref-for-dom-url-href①
link-type="idl"} getter return value):

``` {.lang-javascript .highlight}
var url = new URL("🏳️‍🌈", new URL("https://pride.example/hello-world"))
url.pathname // "/%F0%9F%8F%B3%EF%B8%8F%E2%80%8D%F0%9F%8C%88"
```
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="parse(url, base)" algorithm-for="URL"}
The static
[`parse(``url`{.variable}`, ``base`{.variable}`)`]{#dom-url-parse .dfn
.dfn-paneled .idl-code dfn-for="URL" dfn-type="method" export=""
lt="parse(url, base)|parse(url)"} method steps are:

1.  Let `parsedURL`{.variable} be the result of running the [API URL
    parser](#api-url-parser){#ref-for-api-url-parser① link-type="dfn"}
    on `url`{.variable} with `base`{.variable}, if given.

2.  If `parsedURL`{.variable} is failure, then return null.

3.  Let `url`{.variable} be a new [`URL`{.idl}](#url){#ref-for-url⑦
    link-type="idl"} object.

4.  [Initialize](#url-initialize){#ref-for-url-initialize①
    link-type="dfn"} `url`{.variable} with `parsedURL`{.variable}.

5.  Return `url`{.variable}.
:::

::: {.algorithm algorithm="canParse(url, base)" algorithm-for="URL"}
The static
[`canParse(``url`{.variable}`, ``base`{.variable}`)`]{#dom-url-canparse
.dfn .dfn-paneled .idl-code dfn-for="URL" dfn-type="method" export=""
lt="canParse(url, base)|canParse(url)"} method steps are:

1.  Let `parsedURL`{.variable} be the result of running the [API URL
    parser](#api-url-parser){#ref-for-api-url-parser② link-type="dfn"}
    on `url`{.variable} with `base`{.variable}, if given.

2.  If `parsedURL`{.variable} is failure, then return false.

3.  Return true.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="href getter"}
The [`href`]{#dom-url-href .dfn .dfn-paneled .idl-code dfn-for="URL"
dfn-type="attribute" export=""} getter steps and the
[`toJSON()`]{#dom-url-tojson .dfn .dfn-paneled .idl-code dfn-for="URL"
dfn-type="method" export=""} method steps are to return the
[serialization](#concept-url-serializer){#ref-for-concept-url-serializer⑨
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url①
link-type="dfn"}.
:::

::: {.algorithm algorithm="href setter"}
The [`href`](#dom-url-href){#ref-for-dom-url-href② .idl-code
link-type="attribute"} setter steps are:

1.  Let `parsedURL`{.variable} be the result of running the [basic URL
    parser](#concept-basic-url-parser){#ref-for-concept-basic-url-parser⑦
    link-type="dfn"} on the given value.

2.  If `parsedURL`{.variable} is failure, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①
    link-type="idl"}.

3.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②
    link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url②
    link-type="dfn"} to `parsedURL`{.variable}.

4.  Empty [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③
    link-type="dfn"}'s [query
    object](#concept-url-query-object){#ref-for-concept-url-query-object③
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list
    link-type="dfn"}.

5.  Let `query`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④
    link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url③
    link-type="dfn"}'s
    [query](#concept-url-query){#ref-for-concept-url-query②⓪
    link-type="dfn"}.

6.  If `query`{.variable} is non-null, then set
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤
    link-type="dfn"}'s [query
    object](#concept-url-query-object){#ref-for-concept-url-query-object④
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①
    link-type="dfn"} to the result of
    [parsing](#concept-urlencoded-string-parser){#ref-for-concept-urlencoded-string-parser
    link-type="dfn"} `query`{.variable}.
:::

::: {.algorithm algorithm="origin" algorithm-for="URL"}
The [`origin`]{#dom-url-origin .dfn .dfn-paneled .idl-code dfn-for="URL"
dfn-type="attribute" export=""} getter steps are to return the
[serialization](https://html.spec.whatwg.org/multipage/browsers.html#ascii-serialisation-of-an-origin){#ref-for-ascii-serialisation-of-an-origin
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url④
link-type="dfn"}'s
[origin](#concept-url-origin){#ref-for-concept-url-origin②
link-type="dfn"}.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="protocol" algorithm-for="URL"}
The [`protocol`]{#dom-url-protocol .dfn .dfn-paneled .idl-code
dfn-for="URL" dfn-type="attribute" export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url⑤
link-type="dfn"}'s
[scheme](#concept-url-scheme){#ref-for-concept-url-scheme③⑨
link-type="dfn"}, followed by U+003A (:).
:::

::: {.algorithm algorithm="protocol setter"}
The [`protocol`](#dom-url-protocol){#ref-for-dom-url-protocol① .idl-code
link-type="attribute"} setter steps are to [basic URL
parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser⑧
link-type="dfn"} the given value, followed by U+003A (:), with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url⑥
link-type="dfn"} as
[*url*](#basic-url-parser-url){#ref-for-basic-url-parser-url
link-type="dfn"} and [scheme start
state](#scheme-start-state){#ref-for-scheme-start-state①
link-type="dfn"} as [*state
override*](#basic-url-parser-state-override){#ref-for-basic-url-parser-state-override
link-type="dfn"}.
:::

::: {.algorithm algorithm="username" algorithm-for="URL"}
The [`username`]{#dom-url-username .dfn .dfn-paneled .idl-code
dfn-for="URL" dfn-type="attribute" export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url⑦
link-type="dfn"}'s
[username](#concept-url-username){#ref-for-concept-url-username①①
link-type="dfn"}.
:::

::: {.algorithm algorithm="username setter"}
The [`username`](#dom-url-username){#ref-for-dom-url-username① .idl-code
link-type="attribute"} setter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪
    link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url⑧
    link-type="dfn"} [cannot have a
    username/password/port](#cannot-have-a-username-password-port){#ref-for-cannot-have-a-username-password-port
    link-type="dfn"}, then return.

2.  [Set the username](#set-the-username){#ref-for-set-the-username
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①
    link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url⑨
    link-type="dfn"} and the given value.
:::

::: {.algorithm algorithm="password" algorithm-for="URL"}
The [`password`]{#dom-url-password .dfn .dfn-paneled .idl-code
dfn-for="URL" dfn-type="attribute" export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url①⓪
link-type="dfn"}'s
[password](#concept-url-password){#ref-for-concept-url-password①②
link-type="dfn"}.
:::

::: {.algorithm algorithm="password setter"}
The [`password`](#dom-url-password){#ref-for-dom-url-password① .idl-code
link-type="attribute"} setter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①① link-type="dfn"}
    [cannot have a
    username/password/port](#cannot-have-a-username-password-port){#ref-for-cannot-have-a-username-password-port①
    link-type="dfn"}, then return.

2.  [Set the password](#set-the-password){#ref-for-set-the-password
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①② link-type="dfn"}
    and the given value.
:::

::: {.algorithm algorithm="host" algorithm-for="URL"}
The [`host`]{#dom-url-host .dfn .dfn-paneled .idl-code dfn-for="URL"
dfn-type="attribute" export=""} getter steps are:

1.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①③ link-type="dfn"}.

2.  If `url`{.variable}'s
    [host](#concept-url-host){#ref-for-concept-url-host③②
    link-type="dfn"} is null, then return the empty string.

3.  If `url`{.variable}'s
    [port](#concept-url-port){#ref-for-concept-url-port①③
    link-type="dfn"} is null, return `url`{.variable}'s
    [host](#concept-url-host){#ref-for-concept-url-host③③
    link-type="dfn"},
    [serialized](#concept-host-serializer){#ref-for-concept-host-serializer⑤
    link-type="dfn"}.

4.  Return `url`{.variable}'s
    [host](#concept-url-host){#ref-for-concept-url-host③④
    link-type="dfn"},
    [serialized](#concept-host-serializer){#ref-for-concept-host-serializer⑥
    link-type="dfn"}, followed by U+003A (:) and `url`{.variable}'s
    [port](#concept-url-port){#ref-for-concept-url-port①④
    link-type="dfn"},
    [serialized](#serialize-an-integer){#ref-for-serialize-an-integer②
    link-type="dfn"}.
:::

::: {.algorithm algorithm="host setter"}
The [`host`](#dom-url-host){#ref-for-dom-url-host① .idl-code
link-type="attribute"} setter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①④ link-type="dfn"}
    has an [opaque path](#url-opaque-path){#ref-for-url-opaque-path⑨
    link-type="dfn"}, then return.

2.  [Basic URL
    parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser⑨
    link-type="dfn"} the given value with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①⑤ link-type="dfn"}
    as [*url*](#basic-url-parser-url){#ref-for-basic-url-parser-url①
    link-type="dfn"} and [host state](#host-state){#ref-for-host-state①
    link-type="dfn"} as [*state
    override*](#basic-url-parser-state-override){#ref-for-basic-url-parser-state-override①
    link-type="dfn"}.

If the given value for the
[`host`](#dom-url-host){#ref-for-dom-url-host② .idl-code
link-type="attribute"} setter lacks a
[port](#url-port-string){#ref-for-url-port-string② link-type="dfn"},
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url①⑥
link-type="dfn"}'s [port](#concept-url-port){#ref-for-concept-url-port①⑤
link-type="dfn"} will not change. This can be unexpected as `host`
getter does return a [URL-port
string](#url-port-string){#ref-for-url-port-string③ link-type="dfn"} so
one might have assumed the setter to always \"reset\" both.
:::

::: {.algorithm algorithm="hostname" algorithm-for="URL"}
The [`hostname`]{#dom-url-hostname .dfn .dfn-paneled .idl-code
dfn-for="URL" dfn-type="attribute" export=""} getter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①⑦
    link-type="dfn"}'s
    [host](#concept-url-host){#ref-for-concept-url-host③⑤
    link-type="dfn"} is null, then return the empty string.

2.  Return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①⑧
    link-type="dfn"}'s
    [host](#concept-url-host){#ref-for-concept-url-host③⑥
    link-type="dfn"},
    [serialized](#concept-host-serializer){#ref-for-concept-host-serializer⑦
    link-type="dfn"}.
:::

::: {.algorithm algorithm="hostname setter"}
The [`hostname`](#dom-url-hostname){#ref-for-dom-url-hostname① .idl-code
link-type="attribute"} setter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url①⑨ link-type="dfn"}
    has an [opaque path](#url-opaque-path){#ref-for-url-opaque-path①⓪
    link-type="dfn"}, then return.

2.  [Basic URL
    parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser①⓪
    link-type="dfn"} the given value with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②⓪ link-type="dfn"}
    as [*url*](#basic-url-parser-url){#ref-for-basic-url-parser-url②
    link-type="dfn"} and [hostname
    state](#hostname-state){#ref-for-hostname-state① link-type="dfn"} as
    [*state
    override*](#basic-url-parser-state-override){#ref-for-basic-url-parser-state-override②
    link-type="dfn"}.
:::

::: {.algorithm algorithm="port" algorithm-for="URL"}
The [`port`]{#dom-url-port .dfn .dfn-paneled .idl-code dfn-for="URL"
dfn-type="attribute" export=""} getter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②①
    link-type="dfn"}'s
    [port](#concept-url-port){#ref-for-concept-url-port①⑥
    link-type="dfn"} is null, then return the empty string.

2.  Return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②②
    link-type="dfn"}'s
    [port](#concept-url-port){#ref-for-concept-url-port①⑦
    link-type="dfn"},
    [serialized](#serialize-an-integer){#ref-for-serialize-an-integer③
    link-type="dfn"}.
:::

::: {.algorithm algorithm="port setter"}
The [`port`](#dom-url-port){#ref-for-dom-url-port① .idl-code
link-type="attribute"} setter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②③ link-type="dfn"}
    [cannot have a
    username/password/port](#cannot-have-a-username-password-port){#ref-for-cannot-have-a-username-password-port②
    link-type="dfn"}, then return.

2.  If the given value is the empty string, then set
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②④
    link-type="dfn"}'s
    [port](#concept-url-port){#ref-for-concept-url-port①⑧
    link-type="dfn"} to null.

3.  Otherwise, [basic URL
    parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser①①
    link-type="dfn"} the given value with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②⑤ link-type="dfn"}
    as [*url*](#basic-url-parser-url){#ref-for-basic-url-parser-url③
    link-type="dfn"} and [port state](#port-state){#ref-for-port-state①
    link-type="dfn"} as [*state
    override*](#basic-url-parser-state-override){#ref-for-basic-url-parser-state-override③
    link-type="dfn"}.
:::

::: {.algorithm algorithm="pathname" algorithm-for="URL"}
The [`pathname`]{#dom-url-pathname .dfn .dfn-paneled .idl-code
dfn-for="URL" dfn-type="attribute" export=""} getter steps are to return
the result of [URL path
serializing](#url-path-serializer){#ref-for-url-path-serializer②
link-type="dfn"}
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧
link-type="dfn"}'s [URL](#concept-url-url){#ref-for-concept-url-url②⑥
link-type="dfn"}.
:::

::: {.algorithm algorithm="pathname setter"}
The [`pathname`](#dom-url-pathname){#ref-for-dom-url-pathname① .idl-code
link-type="attribute"} setter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②⑦ link-type="dfn"}
    has an [opaque path](#url-opaque-path){#ref-for-url-opaque-path①①
    link-type="dfn"}, then return.

2.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②⑧
    link-type="dfn"}'s
    [path](#concept-url-path){#ref-for-concept-url-path③②
    link-type="dfn"}.

3.  [Basic URL
    parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser①②
    link-type="dfn"} the given value with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url②⑨ link-type="dfn"}
    as [*url*](#basic-url-parser-url){#ref-for-basic-url-parser-url④
    link-type="dfn"} and [path start
    state](#path-start-state){#ref-for-path-start-state④
    link-type="dfn"} as [*state
    override*](#basic-url-parser-state-override){#ref-for-basic-url-parser-state-override④
    link-type="dfn"}.
:::

::: {.algorithm algorithm="search" algorithm-for="URL"}
The [`search`]{#dom-url-search .dfn .dfn-paneled .idl-code dfn-for="URL"
dfn-type="attribute" export=""} getter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③⓪
    link-type="dfn"}'s
    [query](#concept-url-query){#ref-for-concept-url-query②①
    link-type="dfn"} is either null or the empty string, then return the
    empty string.

2.  Return U+003F (?), followed by
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③③
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③①
    link-type="dfn"}'s
    [query](#concept-url-query){#ref-for-concept-url-query②②
    link-type="dfn"}.
:::

::: {.algorithm algorithm="search setter"}
The [`search`](#dom-url-search){#ref-for-dom-url-search① .idl-code
link-type="attribute"} setter steps are:

1.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③④
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③② link-type="dfn"}.

2.  If the given value is the empty string, then set `url`{.variable}'s
    [query](#concept-url-query){#ref-for-concept-url-query②③
    link-type="dfn"} to null,
    [empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑤
    link-type="dfn"}'s [query
    object](#concept-url-query-object){#ref-for-concept-url-query-object⑤
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list②
    link-type="dfn"}, and return.

3.  Let `input`{.variable} be the given value with a single leading
    U+003F (?) removed, if any.

4.  Set `url`{.variable}'s
    [query](#concept-url-query){#ref-for-concept-url-query②④
    link-type="dfn"} to the empty string.

5.  [Basic URL
    parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser①③
    link-type="dfn"} `input`{.variable} with `url`{.variable} as
    [*url*](#basic-url-parser-url){#ref-for-basic-url-parser-url⑤
    link-type="dfn"} and [query
    state](#query-state){#ref-for-query-state⑤ link-type="dfn"} as
    [*state
    override*](#basic-url-parser-state-override){#ref-for-basic-url-parser-state-override⑤
    link-type="dfn"}.

6.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑥
    link-type="dfn"}'s [query
    object](#concept-url-query-object){#ref-for-concept-url-query-object⑥
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list③
    link-type="dfn"} to the result of
    [parsing](#concept-urlencoded-string-parser){#ref-for-concept-urlencoded-string-parser①
    link-type="dfn"} `input`{.variable}.
:::

::: {.algorithm algorithm="searchParams" algorithm-for="URL"}
The [`searchParams`]{#dom-url-searchparams .dfn .dfn-paneled .idl-code
dfn-for="URL" dfn-type="attribute" export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑦
link-type="dfn"}'s [query
object](#concept-url-query-object){#ref-for-concept-url-query-object⑦
link-type="dfn"}.
:::

::: {.algorithm algorithm="hash" algorithm-for="URL"}
The [`hash`]{#dom-url-hash .dfn .dfn-paneled .idl-code dfn-for="URL"
dfn-type="attribute" export=""} getter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑧
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③③
    link-type="dfn"}'s
    [fragment](#concept-url-fragment){#ref-for-concept-url-fragment①②
    link-type="dfn"} is either null or the empty string, then return the
    empty string.

2.  Return U+0023 (#), followed by
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑨
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③④
    link-type="dfn"}'s
    [fragment](#concept-url-fragment){#ref-for-concept-url-fragment①③
    link-type="dfn"}.
:::

::: {.algorithm algorithm="hash setter"}
The [`hash`](#dom-url-hash){#ref-for-dom-url-hash① .idl-code
link-type="attribute"} setter steps are:

1.  If the given value is the empty string, then set
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⓪
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③⑤
    link-type="dfn"}'s
    [fragment](#concept-url-fragment){#ref-for-concept-url-fragment①④
    link-type="dfn"} to null and return.

2.  Let `input`{.variable} be the given value with a single leading
    U+0023 (#) removed, if any.

3.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④①
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③⑥
    link-type="dfn"}'s
    [fragment](#concept-url-fragment){#ref-for-concept-url-fragment①⑤
    link-type="dfn"} to the empty string.

4.  [Basic URL
    parse](#concept-basic-url-parser){#ref-for-concept-basic-url-parser①④
    link-type="dfn"} `input`{.variable} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④②
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③⑦ link-type="dfn"}
    as [*url*](#basic-url-parser-url){#ref-for-basic-url-parser-url⑥
    link-type="dfn"} and [fragment
    state](#fragment-state){#ref-for-fragment-state⑦ link-type="dfn"} as
    [*state
    override*](#basic-url-parser-state-override){#ref-for-basic-url-parser-state-override⑥
    link-type="dfn"}.
:::

### [6.2. ]{.secno}[URLSearchParams class]{.content}[](#interface-urlsearchparams){.self-link} {#interface-urlsearchparams .heading .settled level="6.2"}

``` {.idl .highlight .def}
[Exposed=*]
interface URLSearchParams {
  constructor(optional (sequence<sequence<USVString>> or record<USVString, USVString> or USVString) init = "");

  readonly attribute unsigned long size;

  undefined append(USVString name, USVString value);
  undefined delete(USVString name, optional USVString value);
  USVString? get(USVString name);
  sequence<USVString> getAll(USVString name);
  boolean has(USVString name, optional USVString value);
  undefined set(USVString name, USVString value);

  undefined sort();

  iterable<USVString, USVString>;
  stringifier;
};
```

::: {#example-constructing-urlsearchparams .example}
[](#example-constructing-urlsearchparams){.self-link}

Constructing and stringifying a
[`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams③
link-type="idl"} object is fairly straightforward:

``` {.lang-javascript .highlight}
let params = new URLSearchParams({key: "730d67"})
params.toString() // "key=730d67"
```
:::

::: {.note role="note"}
As a
[`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams④
link-type="idl"} object uses the
[`application/x-www-form-urlencoded`](#concept-urlencoded){#ref-for-concept-urlencoded
link-type="dfn"} format underneath there are some difference with how it
encodes certain code points compared to a
[`URL`{.idl}](#url){#ref-for-url⑧ link-type="idl"} object (including
[`href`{.idl}](#dom-url-href){#ref-for-dom-url-href③ link-type="idl"}
and [`search`{.idl}](#dom-url-search){#ref-for-dom-url-search②
link-type="idl"}). This can be especially surprising when using
[`searchParams`{.idl}](#dom-url-searchparams){#ref-for-dom-url-searchparams①
link-type="idl"} to operate on a
[URL](#concept-url){#ref-for-concept-url⑥② link-type="dfn"}'s
[query](#concept-url-query){#ref-for-concept-url-query②⑤
link-type="dfn"}.

``` {.lang-javascript .highlight}
const url = new URL('https://example.com/?a=b ~');
console.log(url.href);   // "https://example.com/?a=b%20~"
url.searchParams.sort();
console.log(url.href);   // "https://example.com/?a=b+%7E"
```

``` {.lang-javascript .highlight}
const url = new URL('https://example.com/?a=~&b=%7E');
console.log(url.search);                // "?a=~&b=%7E"
console.log(url.searchParams.get('a')); // "~"
console.log(url.searchParams.get('b')); // "~"
```

[`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams⑤
link-type="idl"} objects will percent-encode anything in the
[`application/x-www-form-urlencoded` percent-encode
set](#application-x-www-form-urlencoded-percent-encode-set){#ref-for-application-x-www-form-urlencoded-percent-encode-set④
link-type="dfn"}, and will encode U+0020 SPACE as U+002B (+).

Ignoring encodings (use
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8①①
link-type="dfn"}),
[`search`{.idl}](#dom-url-search){#ref-for-dom-url-search③
link-type="idl"} will percent-encode anything in the [query
percent-encode
set](#query-percent-encode-set){#ref-for-query-percent-encode-set④
link-type="dfn"} or the [special-query percent-encode
set](#special-query-percent-encode-set){#ref-for-special-query-percent-encode-set①
link-type="dfn"} (depending on whether or not the
[URL](#concept-url){#ref-for-concept-url⑥③ link-type="dfn"} [is
special](#is-special){#ref-for-is-special①⑦ link-type="dfn"}).
:::

A [`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams⑥
link-type="idl"} object has an associated:

- [list]{#concept-urlsearchparams-list .dfn .dfn-paneled
  dfn-for="URLSearchParams" dfn-type="dfn" export=""}: a
  [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑥
  link-type="dfn"} of
  [tuples](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple②
  link-type="dfn"} each consisting of a name and a value, initially
  empty.
- [URL object]{#concept-urlsearchparams-url-object .dfn .dfn-paneled
  dfn-for="URLSearchParams" dfn-type="dfn" export=""}: null or a
  [`URL`{.idl}](#url){#ref-for-url⑨ link-type="idl"} object, initially
  null.

::: {.algorithm algorithm="initialize" algorithm-for="URLSearchParams"}
To [[]{#concept-urlsearchparams-new
.bs-old-id}initialize]{#urlsearchparams-initialize .dfn .dfn-paneled
dfn-for="URLSearchParams" dfn-type="dfn" noexport=""} a
[`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams⑦
link-type="idl"} object `query`{.variable} with `init`{.variable}:

1.  If `init`{.variable} is a
    [sequence](https://webidl.spec.whatwg.org/#idl-sequence){#ref-for-idl-sequence③
    link-type="dfn"}, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑧
    link-type="dfn"} `innerSequence`{.variable} of `init`{.variable}:

    1.  If `innerSequence`{.variable}'s
        [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑥
        link-type="dfn"} is not 2, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②
        link-type="idl"}.

    2.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑦
        link-type="dfn"} (`innerSequence`{.variable}\[0\],
        `innerSequence`{.variable}\[1\]) to `query`{.variable}'s
        [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list④
        link-type="dfn"}.

2.  Otherwise, if `init`{.variable} is a
    [record](https://webidl.spec.whatwg.org/#idl-record){#ref-for-idl-record①
    link-type="dfn"}, then [for
    each](https://infra.spec.whatwg.org/#map-iterate){#ref-for-map-iterate
    link-type="dfn"} `name`{.variable} → `value`{.variable} of
    `init`{.variable},
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑧
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    `query`{.variable}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list⑤
    link-type="dfn"}.

3.  Otherwise:

    1.  Assert: `init`{.variable} is a string.

    2.  Set `query`{.variable}'s
        [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list⑥
        link-type="dfn"} to the result of
        [parsing](#concept-urlencoded-string-parser){#ref-for-concept-urlencoded-string-parser②
        link-type="dfn"} `init`{.variable}.
:::

::: {.algorithm algorithm="update" algorithm-for="URLSearchParams"}
To [update]{#concept-urlsearchparams-update .dfn .dfn-paneled
dfn-for="URLSearchParams" dfn-type="dfn" noexport=""} a
[`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams⑧
link-type="idl"} object `query`{.variable}:

1.  If `query`{.variable}'s [URL
    object](#concept-urlsearchparams-url-object){#ref-for-concept-urlsearchparams-url-object①
    link-type="dfn"} is null, then return.

2.  Let `serializedQuery`{.variable} be the
    [serialization](#concept-urlencoded-serializer){#ref-for-concept-urlencoded-serializer
    link-type="dfn"} of `query`{.variable}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list⑦
    link-type="dfn"}.

3.  If `serializedQuery`{.variable} is the empty string, then set
    `serializedQuery`{.variable} to null.

4.  Set `query`{.variable}'s [URL
    object](#concept-urlsearchparams-url-object){#ref-for-concept-urlsearchparams-url-object②
    link-type="dfn"}'s
    [URL](#concept-url-url){#ref-for-concept-url-url③⑧
    link-type="dfn"}'s
    [query](#concept-url-query){#ref-for-concept-url-query②⑥
    link-type="dfn"} to `serializedQuery`{.variable}.
:::

::: {.algorithm algorithm="URLSearchParams(init)" algorithm-for="URLSearchParams"}
The
[`new URLSearchParams(``init`{.variable}`)`]{#dom-urlsearchparams-urlsearchparams
.dfn .dfn-paneled .idl-code dfn-for="URLSearchParams"
dfn-type="constructor" export=""
lt="URLSearchParams(init)|constructor(init)|URLSearchParams()|constructor()"}
constructor steps are:

1.  If `init`{.variable} is a string and starts with U+003F (?), then
    remove the first code point from `init`{.variable}.

2.  [Initialize](#urlsearchparams-initialize){#ref-for-urlsearchparams-initialize①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④③
    link-type="dfn"} with `init`{.variable}.
:::

::: {.algorithm algorithm="size" algorithm-for="URLSearchParams"}
The [`size`]{#dom-urlsearchparams-size .dfn .dfn-paneled .idl-code
dfn-for="URLSearchParams" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this④④
link-type="dfn"}'s
[list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list⑧
link-type="dfn"}'s
[size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑦
link-type="dfn"}.
:::

::: {.algorithm algorithm="append(name, value)" algorithm-for="URLSearchParams"}
The
[`append(``name`{.variable}`, ``value`{.variable}`)`]{#dom-urlsearchparams-append
.dfn .dfn-paneled .idl-code dfn-for="URLSearchParams" dfn-type="method"
export=""} method steps are:

1.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑨
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑤
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list⑨
    link-type="dfn"}.

2.  [Update](#concept-urlsearchparams-update){#ref-for-concept-urlsearchparams-update
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑥
    link-type="dfn"}.
:::

::: {.algorithm algorithm="delete(name, value)" algorithm-for="URLSearchParams"}
The
[`delete(``name`{.variable}`, ``value`{.variable}`)`]{#dom-urlsearchparams-delete
.dfn .dfn-paneled .idl-code dfn-for="URLSearchParams" dfn-type="method"
export="" lt="delete(name, value)|delete(name)"} method steps are:

1.  If `value`{.variable} is given, then
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove④
    link-type="dfn"} all
    [tuples](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple③
    link-type="dfn"} whose name is `name`{.variable} and value is
    `value`{.variable} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑦
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①⓪
    link-type="dfn"}.

2.  Otherwise,
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove⑤
    link-type="dfn"} all
    [tuples](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple④
    link-type="dfn"} whose name is `name`{.variable} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑧
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①①
    link-type="dfn"}.

3.  [Update](#concept-urlsearchparams-update){#ref-for-concept-urlsearchparams-update①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑨
    link-type="dfn"}.
:::

::: {.algorithm algorithm="get(name)" algorithm-for="URLSearchParams"}
The [`get(``name`{.variable}`)`]{#dom-urlsearchparams-get .dfn
.dfn-paneled .idl-code dfn-for="URLSearchParams" dfn-type="method"
export=""} method steps are to return the value of the first
[tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple⑤
link-type="dfn"} whose name is `name`{.variable} in
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⓪
link-type="dfn"}'s
[list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①②
link-type="dfn"}, if there is such a
[tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple⑥
link-type="dfn"}; otherwise null.
:::

::: {.algorithm algorithm="getAll(name)" algorithm-for="URLSearchParams"}
The [`getAll(``name`{.variable}`)`]{#dom-urlsearchparams-getall .dfn
.dfn-paneled .idl-code dfn-for="URLSearchParams" dfn-type="method"
export=""} method steps are to return the values of all
[tuples](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple⑦
link-type="dfn"} whose name is `name`{.variable} in
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤①
link-type="dfn"}'s
[list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①③
link-type="dfn"}, in list order; otherwise the empty sequence.
:::

::: {.algorithm algorithm="has(name, value)" algorithm-for="URLSearchParams"}
The
[`has(``name`{.variable}`, ``value`{.variable}`)`]{#dom-urlsearchparams-has
.dfn .dfn-paneled .idl-code dfn-for="URLSearchParams" dfn-type="method"
export="" lt="has(name, value)|has(name)"} method steps are:

1.  If `value`{.variable} is given and there is a
    [tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple⑧
    link-type="dfn"} whose name is `name`{.variable} and value is
    `value`{.variable} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤②
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①④
    link-type="dfn"}, then return true.

2.  If `value`{.variable} is not given and there is a
    [tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple⑨
    link-type="dfn"} whose name is `name`{.variable} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤③
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①⑤
    link-type="dfn"}, then return true.

3.  Return false.
:::

::: {.algorithm algorithm="set(name, value)" algorithm-for="URLSearchParams"}
The
[`set(``name`{.variable}`, ``value`{.variable}`)`]{#dom-urlsearchparams-set
.dfn .dfn-paneled .idl-code dfn-for="URLSearchParams" dfn-type="method"
export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤④
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①⑥
    link-type="dfn"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain
    link-type="dfn"} any
    [tuples](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple①⓪
    link-type="dfn"} whose name is `name`{.variable}, then set the value
    of the first such
    [tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple①①
    link-type="dfn"} to `value`{.variable} and
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove⑥
    link-type="dfn"} the others.

2.  Otherwise,
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①⓪
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑤
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①⑦
    link-type="dfn"}.

3.  [Update](#concept-urlsearchparams-update){#ref-for-concept-urlsearchparams-update②
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑥
    link-type="dfn"}.
:::

------------------------------------------------------------------------

::: {#example-searchparams-sort .example}
[](#example-searchparams-sort){.self-link}

It can be useful to sort the name-value tuples in a
[`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams⑨
link-type="idl"} object, in particular to increase cache hits. This can
be accomplished through invoking the
[`sort()`{.idl}](#dom-urlsearchparams-sort){#ref-for-dom-urlsearchparams-sort①
link-type="idl"} method:

``` {.lang-javascript .highlight}
const url = new URL("https://example.org/?q=🏳️‍🌈&key=e1f7bc78");
url.searchParams.sort();
url.search; // "?key=e1f7bc78&q=%F0%9F%8F%B3%EF%B8%8F%E2%80%8D%F0%9F%8C%88"
```

To avoid altering the original input, e.g., for comparison purposes,
construct a new
[`URLSearchParams`{.idl}](#urlsearchparams){#ref-for-urlsearchparams①⓪
link-type="idl"} object:

``` {.lang-javascript .highlight}
const sorted = new URLSearchParams(url.search)
sorted.sort()
```
:::

::: {.algorithm algorithm="sort()" algorithm-for="URLSearchParams"}
The [`sort()`]{#dom-urlsearchparams-sort .dfn .dfn-paneled .idl-code
dfn-for="URLSearchParams" dfn-type="method" export=""} method steps are:

1.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑦
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①⑧
    link-type="dfn"} to the result of [sorting in ascending
    order](https://infra.spec.whatwg.org/#list-sort-in-ascending-order){#ref-for-list-sort-in-ascending-order
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑧
    link-type="dfn"}'s
    [list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list①⑨
    link-type="dfn"}, with `a`{.variable} being less than `b`{.variable}
    if `a`{.variable}'s name is [code unit less
    than](https://infra.spec.whatwg.org/#code-unit-less-than){#ref-for-code-unit-less-than
    link-type="dfn"} `b`{.variable}'s name.

2.  [Update](#concept-urlsearchparams-update){#ref-for-concept-urlsearchparams-update③
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑨
    link-type="dfn"}.
:::

------------------------------------------------------------------------

The [value pairs to iterate
over](https://webidl.spec.whatwg.org/#dfn-value-pairs-to-iterate-over){#ref-for-dfn-value-pairs-to-iterate-over
link-type="dfn"} are
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⓪
link-type="dfn"}'s
[list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list②⓪
link-type="dfn"}'s
[tuples](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple①②
link-type="dfn"} with the key being the name and the value being the
value.

The [stringification behavior]{#urlsearchparams-stringification-behavior
.dfn .dfn-paneled dfn-for="URLSearchParams" dfn-type="dfn"
lt="stringificationbehavior" noexport=""} steps are to return the
[serialization](#concept-urlencoded-serializer){#ref-for-concept-urlencoded-serializer①
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥①
link-type="dfn"}'s
[list](#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list②①
link-type="dfn"}.

### [6.3. ]{.secno}[URL APIs elsewhere]{.content}[](#url-apis-elsewhere){.self-link} {#url-apis-elsewhere .heading .settled level="6.3"}

A standard that exposes [URLs](#concept-url){#ref-for-concept-url⑥④
link-type="dfn"}, should expose the
[URL](#concept-url){#ref-for-concept-url⑥⑤ link-type="dfn"} as a string
(by
[serializing](#concept-url-serializer){#ref-for-concept-url-serializer①⓪
link-type="dfn"} an internal [URL](#concept-url){#ref-for-concept-url⑥⑥
link-type="dfn"}). A standard should not expose a
[URL](#concept-url){#ref-for-concept-url⑥⑦ link-type="dfn"} using a
[`URL`{.idl}](#url){#ref-for-url①⓪ link-type="idl"} object.
[`URL`{.idl}](#url){#ref-for-url①① link-type="idl"} objects are meant
for [URL](#concept-url){#ref-for-concept-url⑥⑧ link-type="dfn"}
manipulation. In IDL the USVString type should be used.

The higher-level notion here is that values are to be exposed as
immutable data structures.

If a standard decides to use a variant of the name \"URL\" for a feature
it defines, it should name such a feature \"url\" (i.e., lowercase and
with an \"l\" at the end). Names such as \"URL\", \"URI\", and \"IRI\"
should not be used. However, if the name is a compound, \"URL\" (i.e.,
uppercase) is preferred, e.g., \"newURL\" and \"oldURL\".

The
[`EventSource`{.idl}](https://html.spec.whatwg.org/multipage/server-sent-events.html#eventsource){#ref-for-eventsource
link-type="idl"} and
[`HashChangeEvent`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#hashchangeevent){#ref-for-hashchangeevent
link-type="idl"} interfaces in HTML are examples of proper naming.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

## [Acknowledgments]{.content}[](#acknowledgments){.self-link} {#acknowledgments .no-num .heading .settled}

There have been a lot of people that have helped make
[URLs](#concept-url){#ref-for-concept-url⑥⑨ link-type="dfn"} more
interoperable over the years and thereby furthered the goals of this
standard. Likewise many people have helped making this standard what it
is today.

With that, many thanks to 100の人, Adam Barth, Addison Phillips, Adrián
Chaves, Adrien Ricciardi, Albert Wiersch, Alex Christensen, Alexis Hunt,
Alexandre Morgaut, Alexis Hunt, Alwin Blok, Andrew Sullivan, Arkadiusz
Michalski, Behnam Esfahbod, Bobby Holley, Boris Zbarsky, Brad Hill,
Brandon Ross, Cailyn Hansen, Chris Dumez, Chris Rebert, Corey Farwell,
Dan Appelquist, Daniel Bratell, Daniel Stenberg, David Burns, David
Håsäther, David Sheets, David Singer, David Walp, Domenic Denicola,
Emily Schechter, Emily Stark, Eric Lawrence, Erik Arvidsson, Gavin
Carothers, Geoff Richards, Glenn Maynard, Gordon P. Hemsley, hemanth,
Henri Sivonen, Ian Hickson, Ilya Grigorik, Italo A. Casas, Jakub
Gieryluk, James Graham, James Manger, James Ross, Jeff Hodges, Jeffrey
Posnick, Jeffrey Yasskin, Joe Duarte, Joshua Bell, Jxck, Karl Wagner,
Kemal Zebari, 田村健人 (Kent TAMURA), Kevin Grandon, Kornel Lesiński,
Larry Masinter, Leif Halvard Silli, Mark Amery, Mark Davis, Marcos
Cáceres, Marijn Kruisselbrink, Martin Dürst, Mathias Bynens, Matt
Falkenhagen, Matt Giuca, Michael Peick, Michael™ Smith, Michal Bukovský,
Michel Suignard, Mikaël Geljić, Noah Levitt, Peter Occil, Philip
Jägenstedt, Philippe Ombredanne, Prayag Verma, Rimas Misevičius, Robert
Kieffer, Rodney Rehm, Roy Fielding, Ryan Sleevi, Sam Ruby, Sam Sneddon,
Santiago M. Mola, Sebastian Mayr, Shannon Booth, Simon Pieters, Simon
Sapin, Steven Vachon, Stuart Cook, Sven Uhlig, Tab Atkins, 吉野剛史
(Takeshi Yoshino), Tantek Çelik, Tiancheng \"Timothy\" Gu, Tim
Berners-Lee, 簡冠庭 (Tim Guan-tin Chien), Titi_Alone, Tomek Wytrębowicz,
Trevor Rowbotham, Tristan Seligmann, Valentin Gosu, Vyacheslav Matva,
Wei Wang, Wolf Lammen, 山岸和利 (Yamagishi Kazutoshi), Yongsheng Zhang,
成瀬ゆい (Yui Naruse), and zealousidealroll for being awesome!

This standard is written by [Anne van
Kesteren](https://annevankesteren.nl/){lang="nl"}
([Apple](https://www.apple.com/), <annevk@annevk.nl>).

## [Intellectual property rights]{.content}[](#ipr){.self-link} {#ipr .no-num .heading .settled}

Copyright © WHATWG (Apple, Google, Mozilla, Microsoft). This work is
licensed under a [Creative Commons Attribution 4.0 International
License](https://creativecommons.org/licenses/by/4.0/){rel="license"}.
To the extent portions of it are incorporated into source code, such
portions in the source code are licensed under the [BSD 3-Clause
License](https://opensource.org/licenses/BSD-3-Clause){rel="license"}
instead.

This is the Living Standard. Those interested in the patent-review
version should view the [Living Standard Review
Draft](/review-drafts/2025-02/).
