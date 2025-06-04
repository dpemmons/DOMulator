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

