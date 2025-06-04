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

