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

