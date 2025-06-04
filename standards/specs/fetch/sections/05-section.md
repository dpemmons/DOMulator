## [5. ]{.secno}[Fetch API]{.content}[](#fetch-api){.self-link} {#fetch-api .heading .settled level="5"}

The [`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch④ .idl-code
link-type="method"} method is relatively low-level API for
[fetching](#concept-fetch){#ref-for-concept-fetch②⑥ link-type="dfn"}
resources. It covers slightly more ground than
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest⑤
link-type="idl"}, although it is currently lacking when it comes to
request progression (not response progression).

::: {#fetch-blob-example .example}
[](#fetch-blob-example){.self-link}

The [`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch⑤ .idl-code
link-type="method"} method makes it quite straightforward to
[fetch](#concept-fetch){#ref-for-concept-fetch②⑦ link-type="dfn"} a
resource and extract its contents as a
[`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob②
link-type="idl"}:

``` {.lang-javascript .highlight}
fetch("/music/pk/altes-kamuffel.flac")
  .then(res => res.blob()).then(playBlob)
```

If you just care to log a particular response header:

``` {.lang-javascript .highlight}
fetch("/", {method:"HEAD"})
  .then(res => log(res.headers.get("strict-transport-security")))
```

If you want to check a particular response header and then process the
response of a cross-origin resource:

``` {.lang-javascript .highlight}
fetch("https://pk.example/berlin-calling.json", {mode:"cors"})
  .then(res => {
    if(res.headers.get("content-type") &&
       res.headers.get("content-type").toLowerCase().indexOf("application/json") >= 0) {
      return res.json()
    } else {
      throw new TypeError()
    }
  }).then(processJSON)
```

If you want to work with URL query parameters:

``` {.lang-javascript .highlight}
var url = new URL("https://geo.example.org/api"),
    params = {lat:35.696233, long:139.570431}
Object.keys(params).forEach(key => url.searchParams.append(key, params[key]))
fetch(url).then(/* … */)
```

If you want to receive the body data progressively:

``` {.lang-javascript .highlight}
function consume(reader) {
  var total = 0
  return pump()
  function pump() {
    return reader.read().then(({done, value}) => {
      if (done) {
        return
      }
      total += value.byteLength
      log(`received ${value.byteLength} bytes (${total} bytes in total)`)
      return pump()
    })
  }
}

fetch("/music/pk/altes-kamuffel.flac")
  .then(res => consume(res.body.getReader()))
  .then(() => log("consumed the entire body without keeping the whole thing in memory!"))
  .catch(e => log("something went wrong: " + e))
```
:::

### [5.1. ]{.secno}[Headers class]{.content}[](#headers-class){.self-link} {#headers-class .heading .settled level="5.1"}

``` {.idl .highlight .def}
typedef (sequence<sequence<ByteString>> or record<ByteString, ByteString>) HeadersInit;

[Exposed=(Window,Worker)]
interface Headers {
  constructor(optional HeadersInit init);

  undefined append(ByteString name, ByteString value);
  undefined delete(ByteString name);
  ByteString? get(ByteString name);
  sequence<ByteString> getSetCookie();
  boolean has(ByteString name);
  undefined set(ByteString name, ByteString value);
  iterable<ByteString, ByteString>;
};
```

A [`Headers`{.idl}](#headers){#ref-for-headers① link-type="idl"} object
has an associated [header list]{#concept-headers-header-list .dfn
.dfn-paneled dfn-for="Headers" dfn-type="dfn" export=""} (a [header
list](#concept-header-list){#ref-for-concept-header-list②②
link-type="dfn"}), which is initially empty. [This can be a pointer to
the [header list](#concept-header-list){#ref-for-concept-header-list②③
link-type="dfn"} of something else, e.g., of a
[request](#concept-request){#ref-for-concept-request①⓪⑨ link-type="dfn"}
as demonstrated by [`Request`{.idl}](#request){#ref-for-request
link-type="idl"} objects.]{.note}

A [`Headers`{.idl}](#headers){#ref-for-headers② link-type="idl"} object
also has an associated [guard]{#concept-headers-guard .dfn .dfn-paneled
dfn-for="Headers" dfn-type="dfn" export=""}, which is a [headers
guard]{#headers-guard .dfn .dfn-paneled dfn-type="dfn" noexport=""}. A
[headers guard](#headers-guard){#ref-for-headers-guard link-type="dfn"}
is \"`immutable`\", \"`request`\", \"`request-no-cors`\", \"`response`\"
or \"`none`\".

------------------------------------------------------------------------

`headers`{.variable}` = new `[`Headers`](#dom-headers){#ref-for-dom-headers① .idl-code link-type="constructor"}`([``init`{.variable}`])`

:   Creates a new [`Headers`{.idl}](#headers){#ref-for-headers③
    link-type="idl"} object. `init`{.variable} can be used to fill its
    internal header list, as per the example below.

    ::: {#example-headers-class .example}
    [](#example-headers-class){.self-link}
    ``` {.lang-javascript .highlight}
    const meta = { "Content-Type": "text/xml", "Breaking-Bad": "<3" };
    new Headers(meta);

    // The above is equivalent to
    const meta2 = [
      [ "Content-Type", "text/xml" ],
      [ "Breaking-Bad", "<3" ]
    ];
    new Headers(meta2);
    ```
    :::

`headers`{.variable}` . `[`append`](#dom-headers-append){#ref-for-dom-headers-append① .idl-code link-type="method"}`(``name`{.variable}`, ``value`{.variable}`)`

:   Appends a header to `headers`{.variable}.

`headers`{.variable}` . `[`delete`](#dom-headers-delete){#ref-for-dom-headers-delete① .idl-code link-type="method"}`(``name`{.variable}`)`

:   Removes a header from `headers`{.variable}.

`headers`{.variable}` . `[`get`](#dom-headers-get){#ref-for-dom-headers-get① .idl-code link-type="method"}`(``name`{.variable}`)`

:   Returns as a string the values of all headers whose name is
    `name`{.variable}, separated by a comma and a space.

`headers`{.variable}` . `[`getSetCookie`](#dom-headers-getsetcookie){#ref-for-dom-headers-getsetcookie① .idl-code link-type="method"}`()`

:   Returns a list of the values for all headers whose name is
    \``Set-Cookie`\`.

`headers`{.variable}` . `[`has`](#dom-headers-has){#ref-for-dom-headers-has① .idl-code link-type="method"}`(``name`{.variable}`)`

:   Returns whether there is a header whose name is `name`{.variable}.

`headers`{.variable}` . `[`set`](#dom-headers-set){#ref-for-dom-headers-set① .idl-code link-type="method"}`(``name`{.variable}`, ``value`{.variable}`)`

:   Replaces the value of the first header whose name is
    `name`{.variable} with `value`{.variable} and removes any remaining
    headers whose name is `name`{.variable}.

`for(const [``name`{.variable}`, ``value`{.variable}`] of ``headers`{.variable}`)`

:   `headers`{.variable} can be iterated over.

------------------------------------------------------------------------

::: {.algorithm algorithm="validate" algorithm-for="Headers"}
To [validate]{#headers-validate .dfn .dfn-paneled dfn-for="Headers"
dfn-type="dfn" noexport=""} a
[header](#concept-header){#ref-for-concept-header⑤⑥ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) for a
[`Headers`{.idl}](#headers){#ref-for-headers④ link-type="idl"} object
`headers`{.variable}:

1.  If `name`{.variable} is not a [header
    name](#header-name){#ref-for-header-name①⑨ link-type="dfn"} or
    `value`{.variable} is not a [header
    value](#header-value){#ref-for-header-value⑨ link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②
    link-type="idl"}.

2.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard
    link-type="dfn"} is \"`immutable`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror③
    link-type="idl"}.

3.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①
    link-type="dfn"} is \"`request`\" and (`name`{.variable},
    `value`{.variable}) is a [forbidden
    request-header](#forbidden-request-header){#ref-for-forbidden-request-header①
    link-type="dfn"}, then return false.

4.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard②
    link-type="dfn"} is \"`response`\" and `name`{.variable} is a
    [forbidden response-header
    name](#forbidden-response-header-name){#ref-for-forbidden-response-header-name②
    link-type="dfn"}, then return false.

5.  Return true.
:::

Steps for \"`request-no-cors`\" are not shared as you cannot have a fake
value (for
[`delete()`{.idl}](#dom-headers-delete){#ref-for-dom-headers-delete②
link-type="idl"}) that always succeeds in [CORS-safelisted
request-header](#cors-safelisted-request-header){#ref-for-cors-safelisted-request-header③
link-type="dfn"}.

::: {.algorithm algorithm="append" algorithm-for="Headers"}
To [append]{#concept-headers-append .dfn .dfn-paneled dfn-for="Headers"
dfn-type="dfn" export=""} a
[header](#concept-header){#ref-for-concept-header⑤⑦ link-type="dfn"}
(`name`{.variable}, `value`{.variable}) to a
[`Headers`{.idl}](#headers){#ref-for-headers⑤ link-type="idl"} object
`headers`{.variable}, run these steps:

1.  [Normalize](#concept-header-value-normalize){#ref-for-concept-header-value-normalize
    link-type="dfn"} `value`{.variable}.

2.  If [validating](#headers-validate){#ref-for-headers-validate
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) for
    `headers`{.variable} returns false, then return.

3.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard③
    link-type="dfn"} is \"`request-no-cors`\":

    1.  Let `temporaryValue`{.variable} be the result of
        [getting](#concept-header-list-get){#ref-for-concept-header-list-get⑦
        link-type="dfn"} `name`{.variable} from `headers`{.variable}'s
        [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list
        link-type="dfn"}.

    2.  If `temporaryValue`{.variable} is null, then set
        `temporaryValue`{.variable} to `value`{.variable}.

    3.  Otherwise, set `temporaryValue`{.variable} to
        `temporaryValue`{.variable}, followed by 0x2C 0x20, followed by
        `value`{.variable}.

    4.  If (`name`{.variable}, `temporaryValue`{.variable}) is not a
        [no-CORS-safelisted
        request-header](#no-cors-safelisted-request-header){#ref-for-no-cors-safelisted-request-header
        link-type="dfn"}, then return.

4.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append②①
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
    `headers`{.variable}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①
    link-type="dfn"}.

5.  If `headers`{.variable}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard④
    link-type="dfn"} is \"`request-no-cors`\", then [remove privileged
    no-CORS
    request-headers](#concept-headers-remove-privileged-no-cors-request-headers){#ref-for-concept-headers-remove-privileged-no-cors-request-headers
    link-type="dfn"} from `headers`{.variable}.
:::

::: {.algorithm algorithm="fill" algorithm-for="Headers"}
To [fill]{#concept-headers-fill .dfn .dfn-paneled dfn-for="Headers"
dfn-type="dfn" export=""} a
[`Headers`{.idl}](#headers){#ref-for-headers⑥ link-type="idl"} object
`headers`{.variable} with a given object `object`{.variable}, run these
steps:

1.  If `object`{.variable} is a
    [sequence](https://webidl.spec.whatwg.org/#idl-sequence){#ref-for-idl-sequence③
    link-type="dfn"}, then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①④
    link-type="dfn"} `header`{.variable} of `object`{.variable}:

    1.  If `header`{.variable}'s
        [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size①
        link-type="dfn"} is not 2, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror④
        link-type="idl"}.

    2.  [Append](#concept-headers-append){#ref-for-concept-headers-append
        link-type="dfn"} (`header`{.variable}\[0\],
        `header`{.variable}\[1\]) to `headers`{.variable}.

2.  Otherwise, `object`{.variable} is a
    [record](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#ref-for-sec-list-and-record-specification-type②
    link-type="dfn"}, then [for
    each](https://infra.spec.whatwg.org/#map-iterate){#ref-for-map-iterate
    link-type="dfn"} `key`{.variable} → `value`{.variable} of
    `object`{.variable},
    [append](#concept-headers-append){#ref-for-concept-headers-append①
    link-type="dfn"} (`key`{.variable}, `value`{.variable}) to
    `headers`{.variable}.
:::

::: {.algorithm algorithm="remove privileged no-CORS request-headers" algorithm-for="Headers"}
To [remove privileged no-CORS
request-headers]{#concept-headers-remove-privileged-no-cors-request-headers
.dfn .dfn-paneled dfn-for="Headers" dfn-type="dfn" noexport=""} from a
[`Headers`{.idl}](#headers){#ref-for-headers⑦ link-type="idl"} object
(`headers`{.variable}), run these steps:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑤
    link-type="dfn"} `headerName`{.variable} of [privileged no-CORS
    request-header
    names](#privileged-no-cors-request-header-name){#ref-for-privileged-no-cors-request-header-name
    link-type="dfn"}:

    1.  [Delete](#concept-header-list-delete){#ref-for-concept-header-list-delete②
        link-type="dfn"} `headerName`{.variable} from
        `headers`{.variable}'s [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list②
        link-type="dfn"}.

This is called when headers are modified by unprivileged code.
:::

::: {.algorithm algorithm="Headers(init)" algorithm-for="Headers"}
The [`new Headers(``init`{.variable}`)`]{#dom-headers .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="constructor" export=""
lt="Headers(init)|constructor(init)|Headers()|constructor()"}
constructor steps are:

1.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑤
    link-type="dfn"} to \"`none`\".

2.  If `init`{.variable} is given, then
    [fill](#concept-headers-fill){#ref-for-concept-headers-fill
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①
    link-type="dfn"} with `init`{.variable}.
:::

::: {.algorithm algorithm="append(name, value)" algorithm-for="Headers"}
The
[`append(``name`{.variable}`, ``value`{.variable}`)`]{#dom-headers-append
.dfn .dfn-paneled .idl-code dfn-for="Headers" dfn-type="method"
export=""} method steps are to
[append](#concept-headers-append){#ref-for-concept-headers-append②
link-type="dfn"} (`name`{.variable}, `value`{.variable}) to
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②
link-type="dfn"}.
:::

::: {.algorithm algorithm="delete(name)" algorithm-for="Headers"}
The [`delete(``name`{.variable}`)`]{#dom-headers-delete .dfn
.dfn-paneled .idl-code dfn-for="Headers" dfn-type="method" export=""}
method steps are:

1.  If [validating](#headers-validate){#ref-for-headers-validate①
    link-type="dfn"} (`name`{.variable}, \`\`) for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③
    link-type="dfn"} returns false, then return.

    Passing a dummy [header
    value](#header-value){#ref-for-header-value①⓪ link-type="dfn"} ought
    not to have any negative repercussions.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑥
    link-type="dfn"} is \"`request-no-cors`\", `name`{.variable} is not
    a [no-CORS-safelisted request-header
    name](#no-cors-safelisted-request-header-name){#ref-for-no-cors-safelisted-request-header-name①
    link-type="dfn"}, and `name`{.variable} is not a [privileged no-CORS
    request-header
    name](#privileged-no-cors-request-header-name){#ref-for-privileged-no-cors-request-header-name①
    link-type="dfn"}, then return.

3.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list③
    link-type="dfn"} [does not
    contain](#header-list-contains){#ref-for-header-list-contains②③
    link-type="dfn"} `name`{.variable}, then return.

4.  [Delete](#concept-header-list-delete){#ref-for-concept-header-list-delete③
    link-type="dfn"} `name`{.variable} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list④
    link-type="dfn"}.

5.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑦
    link-type="dfn"} is \"`request-no-cors`\", then [remove privileged
    no-CORS
    request-headers](#concept-headers-remove-privileged-no-cors-request-headers){#ref-for-concept-headers-remove-privileged-no-cors-request-headers①
    link-type="dfn"} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧
    link-type="dfn"}.
:::

::: {.algorithm algorithm="get(name)" algorithm-for="Headers"}
The [`get(``name`{.variable}`)`]{#dom-headers-get .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="method" export=""} method steps
are:

1.  If `name`{.variable} is not a [header
    name](#header-name){#ref-for-header-name②⓪ link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑤
    link-type="idl"}.

2.  Return the result of
    [getting](#concept-header-list-get){#ref-for-concept-header-list-get⑧
    link-type="dfn"} `name`{.variable} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑤
    link-type="dfn"}.
:::

::: {.algorithm algorithm="getSetCookie()" algorithm-for="Headers"}
The [`getSetCookie()`]{#dom-headers-getsetcookie .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="method" export=""} method steps
are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑥
    link-type="dfn"} [does not
    contain](#header-list-contains){#ref-for-header-list-contains②④
    link-type="dfn"} \``Set-Cookie`\`, then return « ».

2.  Return the
    [values](#concept-header-value){#ref-for-concept-header-value②①
    link-type="dfn"} of all
    [headers](#concept-header){#ref-for-concept-header⑤⑧
    link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑦
    link-type="dfn"} whose
    [name](#concept-header-name){#ref-for-concept-header-name②③
    link-type="dfn"} is a
    [byte-case-insensitive](https://infra.spec.whatwg.org/#byte-case-insensitive){#ref-for-byte-case-insensitive①⑨
    link-type="dfn"} match for \``Set-Cookie`\`, in order.
:::

::: {.algorithm algorithm="has(name)" algorithm-for="Headers"}
The [`has(``name`{.variable}`)`]{#dom-headers-has .dfn .dfn-paneled
.idl-code dfn-for="Headers" dfn-type="method" export=""} method steps
are:

1.  If `name`{.variable} is not a [header
    name](#header-name){#ref-for-header-name②① link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑥
    link-type="idl"}.

2.  Return true if
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑧
    link-type="dfn"}
    [contains](#header-list-contains){#ref-for-header-list-contains②⑤
    link-type="dfn"} `name`{.variable}; otherwise false.
:::

::: {.algorithm algorithm="set(name, value)" algorithm-for="Headers"}
The [`set(``name`{.variable}`, ``value`{.variable}`)`]{#dom-headers-set
.dfn .dfn-paneled .idl-code dfn-for="Headers" dfn-type="method"
export=""} method steps are:

1.  [Normalize](#concept-header-value-normalize){#ref-for-concept-header-value-normalize①
    link-type="dfn"} `value`{.variable}.

2.  If [validating](#headers-validate){#ref-for-headers-validate②
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③
    link-type="dfn"} returns false, then return.

3.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑧
    link-type="dfn"} is \"`request-no-cors`\" and (`name`{.variable},
    `value`{.variable}) is not a [no-CORS-safelisted
    request-header](#no-cors-safelisted-request-header){#ref-for-no-cors-safelisted-request-header①
    link-type="dfn"}, then return.

4.  [Set](#concept-header-list-set){#ref-for-concept-header-list-set①
    link-type="dfn"} (`name`{.variable}, `value`{.variable}) in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤
    link-type="dfn"}'s [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list⑨
    link-type="dfn"}.

5.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard⑨
    link-type="dfn"} is \"`request-no-cors`\", then [remove privileged
    no-CORS
    request-headers](#concept-headers-remove-privileged-no-cors-request-headers){#ref-for-concept-headers-remove-privileged-no-cors-request-headers②
    link-type="dfn"} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦
    link-type="dfn"}.
:::

The [value pairs to iterate
over](https://webidl.spec.whatwg.org/#dfn-value-pairs-to-iterate-over){#ref-for-dfn-value-pairs-to-iterate-over
link-type="dfn"} are the return value of running [sort and
combine](#concept-header-list-sort-and-combine){#ref-for-concept-header-list-sort-and-combine
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧
link-type="dfn"}'s [header
list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⓪
link-type="dfn"}.

### [5.2. ]{.secno}[BodyInit unions]{.content}[](#bodyinit-unions){.self-link} {#bodyinit-unions .heading .settled level="5.2"}

``` {.idl .highlight .def}
typedef (Blob or BufferSource or FormData or URLSearchParams or USVString) XMLHttpRequestBodyInit;

typedef (ReadableStream or XMLHttpRequestBodyInit) BodyInit;
```

To [safely extract]{#bodyinit-safely-extract .dfn .dfn-paneled
dfn-for="BodyInit" dfn-type="dfn" export=""} a [body with
type](#body-with-type){#ref-for-body-with-type link-type="dfn"} from a
[byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence①⑨
link-type="dfn"} or [`BodyInit`{.idl}](#bodyinit){#ref-for-bodyinit
link-type="idl"} object `object`{.variable}, run these steps:

1.  If `object`{.variable} is a
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑤
    link-type="idl"} object, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②③
        link-type="dfn"}: `object`{.variable} is neither
        [disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed
        link-type="dfn"} nor
        [locked](https://streams.spec.whatwg.org/#readablestream-locked){#ref-for-readablestream-locked
        link-type="dfn"}.

2.  Return the result of
    [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract
    link-type="dfn"} `object`{.variable}.

The [safely
extract](#bodyinit-safely-extract){#ref-for-bodyinit-safely-extract⑥
link-type="dfn"} operation is a subset of the
[extract](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract①
link-type="dfn"} operation that is guaranteed to not throw an exception.

To [extract]{#concept-bodyinit-extract .dfn .dfn-paneled
dfn-for="BodyInit" dfn-type="dfn" export=""} a [body with
type](#body-with-type){#ref-for-body-with-type① link-type="dfn"} from a
[byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⓪
link-type="dfn"} or [`BodyInit`{.idl}](#bodyinit){#ref-for-bodyinit①
link-type="idl"} object `object`{.variable}, with an optional boolean
[`keepalive`{.variable}]{#keepalive .dfn .dfn-paneled
dfn-for="BodyInit/extract" dfn-type="dfn" export=""} (default false),
run these steps:

1.  Let `stream`{.variable} be null.

2.  If `object`{.variable} is a
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑥
    link-type="idl"} object, then set `stream`{.variable} to
    `object`{.variable}.

3.  Otherwise, if `object`{.variable} is a
    [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob④
    link-type="idl"} object, set `stream`{.variable} to the result of
    running `object`{.variable}'s [get
    stream](https://w3c.github.io/FileAPI/#blob-get-stream){#ref-for-blob-get-stream
    link-type="dfn"}.

4.  Otherwise, set `stream`{.variable} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new①
    link-type="dfn"}
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑦
    link-type="idl"} object, and [set
    up](https://streams.spec.whatwg.org/#readablestream-set-up-with-byte-reading-support){#ref-for-readablestream-set-up-with-byte-reading-support①
    link-type="dfn"} `stream`{.variable} with byte reading support.

5.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②④
    link-type="dfn"}: `stream`{.variable} is a
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑧
    link-type="idl"} object.

6.  Let `action`{.variable} be null.

7.  Let `source`{.variable} be null.

8.  Let `length`{.variable} be null.

9.  Let `type`{.variable} be null.

10. Switch on `object`{.variable}:

    [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob⑤ link-type="idl"}

    :   Set `source`{.variable} to `object`{.variable}.

        Set `length`{.variable} to `object`{.variable}'s
        [`size`{.idl}](https://w3c.github.io/FileAPI/#dfn-size){#ref-for-dfn-size②
        link-type="idl"}.

        If `object`{.variable}'s
        [`type`{.idl}](https://w3c.github.io/FileAPI/#dfn-type){#ref-for-dfn-type①
        link-type="idl"} attribute is not the empty byte sequence, set
        `type`{.variable} to its value.

    [byte sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②① link-type="dfn"}

    :   Set `source`{.variable} to `object`{.variable}.

    [`BufferSource`{.idl}](https://webidl.spec.whatwg.org/#BufferSource){#ref-for-BufferSource① link-type="idl"}

    :   Set `source`{.variable} to a [copy of the
        bytes](https://webidl.spec.whatwg.org/#dfn-get-buffer-source-copy){#ref-for-dfn-get-buffer-source-copy①
        link-type="dfn"} held by `object`{.variable}.

    [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata② link-type="idl"}

    :   Set `action`{.variable} to this step: run the
        [`multipart/form-data` encoding
        algorithm](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#multipart%2Fform-data-encoding-algorithm){#ref-for-multipart%2Fform-data-encoding-algorithm
        link-type="dfn"}, with `object`{.variable}'s [entry
        list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#ref-for-concept-formdata-entry-list
        link-type="dfn"} and
        [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8①
        link-type="dfn"}.

        Set `source`{.variable} to `object`{.variable}.

        Set `length`{.variable} to [unclear, see
        [html/6424](https://github.com/whatwg/html/issues/6424) for
        improving this]{.XXX}.

        Set `type`{.variable} to \``multipart/form-data; boundary=`\`,
        followed by the [`multipart/form-data` boundary
        string](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#multipart%2Fform-data-boundary-string){#ref-for-multipart%2Fform-data-boundary-string
        link-type="dfn"} generated by the [`multipart/form-data`
        encoding
        algorithm](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#multipart%2Fform-data-encoding-algorithm){#ref-for-multipart%2Fform-data-encoding-algorithm①
        link-type="dfn"}.

    [`URLSearchParams`{.idl}](https://url.spec.whatwg.org/#urlsearchparams){#ref-for-urlsearchparams① link-type="idl"}

    :   Set `source`{.variable} to the result of running the
        [`application/x-www-form-urlencoded`
        serializer](https://url.spec.whatwg.org/#concept-urlencoded-serializer){#ref-for-concept-urlencoded-serializer
        link-type="dfn"} with `object`{.variable}'s
        [list](https://url.spec.whatwg.org/#concept-urlsearchparams-list){#ref-for-concept-urlsearchparams-list
        link-type="dfn"}.

        Set `type`{.variable} to
        \``application/x-www-form-urlencoded;charset=UTF-8`\`.

    [scalar value string](https://infra.spec.whatwg.org/#scalar-value-string){#ref-for-scalar-value-string link-type="dfn"}

    :   Set `source`{.variable} to the [UTF-8
        encoding](https://encoding.spec.whatwg.org/#utf-8-encode){#ref-for-utf-8-encode
        link-type="dfn"} of `object`{.variable}.

        Set `type`{.variable} to \``text/plain;charset=UTF-8`\`.

    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream⑨ link-type="idl"}

    :   If `keepalive`{.variable} is true, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑦
        link-type="idl"}.

        If `object`{.variable} is
        [disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed①
        link-type="dfn"} or
        [locked](https://streams.spec.whatwg.org/#readablestream-locked){#ref-for-readablestream-locked①
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑧
        link-type="idl"}.

11. If `source`{.variable} is a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②②
    link-type="dfn"}, then set `action`{.variable} to a step that
    returns `source`{.variable} and `length`{.variable} to
    `source`{.variable}'s
    [length](https://infra.spec.whatwg.org/#byte-sequence-length){#ref-for-byte-sequence-length⑤
    link-type="dfn"}.

12. If `action`{.variable} is non-null, then run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel⑦
    link-type="dfn"}:

    1.  Run `action`{.variable}.

        Whenever one or more bytes are available and `stream`{.variable}
        is not
        [errored](https://streams.spec.whatwg.org/#readablestream-errored){#ref-for-readablestream-errored①
        link-type="dfn"},
        [enqueue](https://streams.spec.whatwg.org/#readablestream-enqueue){#ref-for-readablestream-enqueue③
        link-type="dfn"} the result of
        [creating](https://webidl.spec.whatwg.org/#arraybufferview-create){#ref-for-arraybufferview-create
        link-type="dfn"} a
        [`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array②
        link-type="idl"} from the available bytes into
        `stream`{.variable}.

        When running `action`{.variable} is done,
        [close](https://streams.spec.whatwg.org/#readablestream-close){#ref-for-readablestream-close①
        link-type="dfn"} `stream`{.variable}.

13. Let `body`{.variable} be a
    [body](#concept-body){#ref-for-concept-body①② link-type="dfn"} whose
    [stream](#concept-body-stream){#ref-for-concept-body-stream①④
    link-type="dfn"} is `stream`{.variable},
    [source](#concept-body-source){#ref-for-concept-body-source①③
    link-type="dfn"} is `source`{.variable}, and
    [length](#concept-body-total-bytes){#ref-for-concept-body-total-bytes③
    link-type="dfn"} is `length`{.variable}.

14. Return (`body`{.variable}, `type`{.variable}).

### [5.3. ]{.secno}[Body mixin]{.content}[](#body-mixin){.self-link} {#body-mixin .heading .settled level="5.3"}

``` {.idl .highlight .def}
interface mixin Body {
  readonly attribute ReadableStream? body;
  readonly attribute boolean bodyUsed;
  [NewObject] Promise<ArrayBuffer> arrayBuffer();
  [NewObject] Promise<Blob> blob();
  [NewObject] Promise<Uint8Array> bytes();
  [NewObject] Promise<FormData> formData();
  [NewObject] Promise<any> json();
  [NewObject] Promise<USVString> text();
};
```

Formats you would not want a network layer to be dependent upon, such as
HTML, will likely not be exposed here. Rather, an HTML parser API might
accept a stream in due course.

Objects including the [`Body`{.idl}](#body){#ref-for-body
link-type="idl"} interface mixin have an associated
[body]{#concept-body-body .dfn .dfn-paneled dfn-for="Body"
dfn-type="dfn" noexport=""} (null or a
[body](#concept-body){#ref-for-concept-body①③ link-type="dfn"}).

An object including the [`Body`{.idl}](#body){#ref-for-body①
link-type="idl"} interface mixin is said to be [unusable]{#body-unusable
.dfn .dfn-paneled dfn-for="Body" dfn-type="dfn" export=""} if its
[body](#concept-body-body){#ref-for-concept-body-body link-type="dfn"}
is non-null and its
[body](#concept-body-body){#ref-for-concept-body-body①
link-type="dfn"}'s
[stream](#concept-body-stream){#ref-for-concept-body-stream①⑤
link-type="dfn"} is
[disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed②
link-type="dfn"} or
[locked](https://streams.spec.whatwg.org/#readablestream-locked){#ref-for-readablestream-locked②
link-type="dfn"}.

------------------------------------------------------------------------

`requestOrResponse`{.variable}` . `[`body`](#dom-body-body){#ref-for-dom-body-body① .idl-code link-type="attribute"}

:   Returns `requestOrResponse`{.variable}'s body as
    [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream①①
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`bodyUsed`](#dom-body-bodyused){#ref-for-dom-body-bodyused① .idl-code link-type="attribute"}

:   Returns whether `requestOrResponse`{.variable}'s body has been read
    from.

`requestOrResponse`{.variable}` . `[`arrayBuffer`](#dom-body-arraybuffer){#ref-for-dom-body-arraybuffer① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`ArrayBuffer`{.idl}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){#ref-for-idl-ArrayBuffer①
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`blob`](#dom-body-blob){#ref-for-dom-body-blob① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob⑦
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`bytes`](#dom-body-bytes){#ref-for-dom-body-bytes① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array④
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`formData`](#dom-body-formdata){#ref-for-dom-body-formdata① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as
    [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata④
    link-type="idl"}.

`requestOrResponse`{.variable}` . `[`json`](#dom-body-json){#ref-for-dom-body-json① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body parsed as JSON.

`requestOrResponse`{.variable}` . `[`text`](#dom-body-text){#ref-for-dom-body-text① .idl-code link-type="method"}`()`

:   Returns a promise fulfilled with `requestOrResponse`{.variable}'s
    body as string.

------------------------------------------------------------------------

::: {.algorithm algorithm="get the MIME type" algorithm-for="Body"}
To [get the MIME type]{#concept-body-mime-type .dfn .dfn-paneled
dfn-for="Body" dfn-type="dfn" noexport=""}, given a
[`Request`{.idl}](#request){#ref-for-request① link-type="idl"} or
[`Response`{.idl}](#response){#ref-for-response① link-type="idl"} object
`requestOrResponse`{.variable}:

1.  Let `headers`{.variable} be null.

2.  If `requestOrResponse`{.variable} is a
    [`Request`{.idl}](#request){#ref-for-request② link-type="idl"}
    object, then set `headers`{.variable} to
    `requestOrResponse`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request
    link-type="dfn"}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list④⑤
    link-type="dfn"}.

3.  Otherwise, set `headers`{.variable} to
    `requestOrResponse`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③②
    link-type="dfn"}.

4.  Let `mimeType`{.variable} be the result of [extracting a MIME
    type](#concept-header-extract-mime-type){#ref-for-concept-header-extract-mime-type⑧
    link-type="dfn"} from `headers`{.variable}.

5.  If `mimeType`{.variable} is failure, then return null.

6.  Return `mimeType`{.variable}.
:::

::: {.algorithm algorithm="body" algorithm-for="Body"}
The [`body`]{#dom-body-body .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="attribute" export=""} getter steps are to return null if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body② link-type="dfn"}
is null; otherwise
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body③
link-type="dfn"}'s
[stream](#concept-body-stream){#ref-for-concept-body-stream①⑥
link-type="dfn"}.
:::

::: {.algorithm algorithm="bodyUsed" algorithm-for="Body"}
The [`bodyUsed`]{#dom-body-bodyused .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body④ link-type="dfn"}
is non-null and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②
link-type="dfn"}'s
[body](#concept-body-body){#ref-for-concept-body-body⑤
link-type="dfn"}'s
[stream](#concept-body-stream){#ref-for-concept-body-stream①⑦
link-type="dfn"} is
[disturbed](https://streams.spec.whatwg.org/#is-readable-stream-disturbed){#ref-for-is-readable-stream-disturbed③
link-type="dfn"}; otherwise false.
:::

::: {.algorithm algorithm="consume body" algorithm-for="Body"}
The [consume body]{#concept-body-consume-body .dfn .dfn-paneled
dfn-for="Body" dfn-type="dfn" noexport=""} algorithm, given an object
that includes [`Body`{.idl}](#body){#ref-for-body② link-type="idl"}
`object`{.variable} and an algorithm that takes a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②③
link-type="dfn"} and returns a JavaScript value or throws an exception
`convertBytesToJSValue`{.variable}, runs these steps:

1.  If `object`{.variable} is
    [unusable](#body-unusable){#ref-for-body-unusable link-type="dfn"},
    then return [a promise rejected
    with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#ref-for-a-promise-rejected-with
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror⑨
    link-type="idl"}.

2.  Let `promise`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise①
    link-type="dfn"}.

3.  Let `errorSteps`{.variable} given `error`{.variable} be to
    [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject
    link-type="dfn"} `promise`{.variable} with `error`{.variable}.

4.  Let `successSteps`{.variable} given a [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②④
    link-type="dfn"} `data`{.variable} be to
    [resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve①
    link-type="dfn"} `promise`{.variable} with the result of running
    `convertBytesToJSValue`{.variable} with `data`{.variable}. If that
    threw an exception, then run `errorSteps`{.variable} with that
    exception.

5.  If `object`{.variable}'s
    [body](#concept-body-body){#ref-for-concept-body-body⑥
    link-type="dfn"} is null, then run `successSteps`{.variable} with an
    empty [byte
    sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑤
    link-type="dfn"}.

6.  Otherwise, [fully read](#body-fully-read){#ref-for-body-fully-read②
    link-type="dfn"} `object`{.variable}'s
    [body](#concept-body-body){#ref-for-concept-body-body⑦
    link-type="dfn"} given `successSteps`{.variable},
    `errorSteps`{.variable}, and `object`{.variable}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global
    link-type="dfn"}.

7.  Return `promise`{.variable}.
:::

::: {.algorithm algorithm="arrayBuffer()" algorithm-for="Body"}
The [`arrayBuffer()`]{#dom-body-arraybuffer .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="method" export=""} method steps are to return
the result of running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③
link-type="dfn"} and the following step given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑥
link-type="dfn"} `bytes`{.variable}: return the result of
[creating](https://webidl.spec.whatwg.org/#arraybuffer-create){#ref-for-arraybuffer-create
link-type="dfn"} an
[`ArrayBuffer`{.idl}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){#ref-for-idl-ArrayBuffer②
link-type="idl"} from `bytes`{.variable} in
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④
link-type="dfn"}'s [relevant
realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm
link-type="dfn"}.

The above method can reject with a
[`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror
link-type="idl"}.
:::

::: {.algorithm algorithm="blob()" algorithm-for="Body"}
The [`blob()`]{#dom-body-blob .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="method" export=""} method steps are to return the result of
running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body①
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤
link-type="dfn"} and the following step given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑦
link-type="dfn"} `bytes`{.variable}: return a
[`Blob`{.idl}](https://w3c.github.io/FileAPI/#dfn-Blob){#ref-for-dfn-Blob⑧
link-type="idl"} whose contents are `bytes`{.variable} and whose
[`type`{.idl}](https://w3c.github.io/FileAPI/#dfn-type){#ref-for-dfn-type②
link-type="idl"} attribute is the result of [get the MIME
type](#concept-body-mime-type){#ref-for-concept-body-mime-type
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥
link-type="dfn"}.
:::

::: {.algorithm algorithm="bytes()" algorithm-for="Body"}
The [`bytes()`]{#dom-body-bytes .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="method" export=""} method steps are to return
the result of running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body②
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦
link-type="dfn"} and the following step given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑧
link-type="dfn"} `bytes`{.variable}: return the result of
[creating](https://webidl.spec.whatwg.org/#arraybufferview-create){#ref-for-arraybufferview-create①
link-type="dfn"} a
[`Uint8Array`{.idl}](https://webidl.spec.whatwg.org/#idl-Uint8Array){#ref-for-idl-Uint8Array⑤
link-type="idl"} from `bytes`{.variable} in
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧
link-type="dfn"}'s [relevant
realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm①
link-type="dfn"}.

The above method can reject with a
[`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror①
link-type="idl"}.
:::

::: {.algorithm algorithm="formData()" algorithm-for="Body"}
The [`formData()`]{#dom-body-formdata .dfn .dfn-paneled .idl-code
dfn-for="Body" dfn-type="method" export=""} method steps are to return
the result of running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body③
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨
link-type="dfn"} and the following steps given a [byte
sequence](https://infra.spec.whatwg.org/#byte-sequence){#ref-for-byte-sequence②⑨
link-type="dfn"} `bytes`{.variable}:

1.  Let `mimeType`{.variable} be the result of [get the MIME
    type](#concept-body-mime-type){#ref-for-concept-body-mime-type①
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪
    link-type="dfn"}.

2.  If `mimeType`{.variable} is non-null, then switch on
    `mimeType`{.variable}'s
    [essence](https://mimesniff.spec.whatwg.org/#mime-type-essence){#ref-for-mime-type-essence⑧
    link-type="dfn"} and run the corresponding steps:

    \"`multipart/form-data`\"

    :   1.  Parse `bytes`{.variable}, using the value of the
            \``boundary`\` parameter from `mimeType`{.variable}, per the
            rules set forth in Returning Values from Forms:
            multipart/form-data.
            [\[RFC7578\]](#biblio-rfc7578 "Returning Values from Forms: multipart/form-data"){link-type="biblio"}

            Each part whose \``Content-Disposition`\` header contains a
            \``filename`\` parameter must be parsed into an
            [entry](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#form-entry){#ref-for-form-entry
            link-type="dfn"} whose value is a
            [`File`{.idl}](https://w3c.github.io/FileAPI/#dfn-file){#ref-for-dfn-file
            link-type="idl"} object whose contents are the contents of
            the part. The
            [`name`{.idl}](https://w3c.github.io/FileAPI/#dfn-name){#ref-for-dfn-name
            link-type="idl"} attribute of the
            [`File`{.idl}](https://w3c.github.io/FileAPI/#dfn-file){#ref-for-dfn-file①
            link-type="idl"} object must have the value of the
            \``filename`\` parameter of the part. The
            [`type`{.idl}](https://w3c.github.io/FileAPI/#dfn-type){#ref-for-dfn-type③
            link-type="idl"} attribute of the
            [`File`{.idl}](https://w3c.github.io/FileAPI/#dfn-file){#ref-for-dfn-file②
            link-type="idl"} object must have the value of the
            \``Content-Type`\` header of the part if the part has such
            header, and \``text/plain`\` (the default defined by
            [\[RFC7578\]](#biblio-rfc7578 "Returning Values from Forms: multipart/form-data"){link-type="biblio"}
            section 4.4) otherwise.

            Each part whose \``Content-Disposition`\` header does not
            contain a \``filename`\` parameter must be parsed into an
            [entry](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#form-entry){#ref-for-form-entry①
            link-type="dfn"} whose value is the [UTF-8 decoded without
            BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom){#ref-for-utf-8-decode-without-bom
            link-type="dfn"} content of the part. [This is done
            regardless of the presence or the value of a
            \``Content-Type`\` header and regardless of the presence or
            the value of a \``charset`\` parameter.]{.note}

            A part whose \``Content-Disposition`\` header contains a
            \``name`\` parameter whose value is \``_charset_`\` is
            parsed like any other part. It does not change the encoding.

        2.  If that fails for some reason, then
            [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦
            link-type="dfn"} a
            [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⓪
            link-type="idl"}.

        3.  Return a new
            [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata⑤
            link-type="idl"} object, appending each
            [entry](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#form-entry){#ref-for-form-entry②
            link-type="dfn"}, resulting from the parsing operation, to
            its [entry
            list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#ref-for-concept-formdata-entry-list①
            link-type="dfn"}.

        The above is a rough approximation of what is needed for
        \``multipart/form-data`\`, a more detailed parsing specification
        is to be written. Volunteers welcome.

    \"`application/x-www-form-urlencoded`\"

    :   1.  Let `entries`{.variable} be the result of
            [parsing](https://url.spec.whatwg.org/#concept-urlencoded-parser){#ref-for-concept-urlencoded-parser
            link-type="dfn"} `bytes`{.variable}.

        2.  Return a new
            [`FormData`{.idl}](https://xhr.spec.whatwg.org/#formdata){#ref-for-formdata⑥
            link-type="idl"} object whose [entry
            list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list){#ref-for-concept-formdata-entry-list②
            link-type="dfn"} is `entries`{.variable}.

3.  [Throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①①
    link-type="idl"}.
:::

::: {.algorithm algorithm="json()" algorithm-for="Body"}
The [`json()`]{#dom-body-json .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="method" export=""} method steps are to return the result of
running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body④
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①
link-type="dfn"} and [parse JSON from
bytes](https://infra.spec.whatwg.org/#parse-json-bytes-to-a-javascript-value){#ref-for-parse-json-bytes-to-a-javascript-value
link-type="dfn"}.

The above method can reject with a
[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror
link-type="idl"}.
:::

::: {.algorithm algorithm="text()" algorithm-for="Body"}
The [`text()`]{#dom-body-text .dfn .dfn-paneled .idl-code dfn-for="Body"
dfn-type="method" export=""} method steps are to return the result of
running [consume
body](#concept-body-consume-body){#ref-for-concept-body-consume-body⑤
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②
link-type="dfn"} and [UTF-8
decode](https://encoding.spec.whatwg.org/#utf-8-decode){#ref-for-utf-8-decode
link-type="dfn"}.
:::

### [5.4. ]{.secno}[Request class]{.content}[](#request-class){.self-link} {#request-class .heading .settled level="5.4"}

``` {.idl .highlight .def}
typedef (Request or USVString) RequestInfo;

[Exposed=(Window,Worker)]
interface Request {
  constructor(RequestInfo input, optional RequestInit init = {});

  readonly attribute ByteString method;
  readonly attribute USVString url;
  [SameObject] readonly attribute Headers headers;

  readonly attribute RequestDestination destination;
  readonly attribute USVString referrer;
  readonly attribute ReferrerPolicy referrerPolicy;
  readonly attribute RequestMode mode;
  readonly attribute RequestCredentials credentials;
  readonly attribute RequestCache cache;
  readonly attribute RequestRedirect redirect;
  readonly attribute DOMString integrity;
  readonly attribute boolean keepalive;
  readonly attribute boolean isReloadNavigation;
  readonly attribute boolean isHistoryNavigation;
  readonly attribute AbortSignal signal;
  readonly attribute RequestDuplex duplex;

  [NewObject] Request clone();
};
Request includes Body;

dictionary RequestInit {
  ByteString method;
  HeadersInit headers;
  BodyInit? body;
  USVString referrer;
  ReferrerPolicy referrerPolicy;
  RequestMode mode;
  RequestCredentials credentials;
  RequestCache cache;
  RequestRedirect redirect;
  DOMString integrity;
  boolean keepalive;
  AbortSignal? signal;
  RequestDuplex duplex;
  RequestPriority priority;
  any window; // can only be set to null
};

enum RequestDestination { "", "audio", "audioworklet", "document", "embed", "font", "frame", "iframe", "image", "json", "manifest", "object", "paintworklet", "report", "script", "sharedworker", "style",  "track", "video", "worker", "xslt" };
enum RequestMode { "navigate", "same-origin", "no-cors", "cors" };
enum RequestCredentials { "omit", "same-origin", "include" };
enum RequestCache { "default", "no-store", "reload", "no-cache", "force-cache", "only-if-cached" };
enum RequestRedirect { "follow", "error", "manual" };
enum RequestDuplex { "half" };
enum RequestPriority { "high", "low", "auto" };
```

\"`serviceworker`\" is omitted from
[`RequestDestination`](#requestdestination){#ref-for-requestdestination②
.idl-code link-type="enum"} as it cannot be observed from JavaScript.
Implementations will still need to support it as a
[destination](#concept-request-destination){#ref-for-concept-request-destination②⓪
link-type="dfn"}. \"`websocket`\" is omitted from
[`RequestMode`](#requestmode){#ref-for-requestmode② .idl-code
link-type="enum"} as it cannot be used nor observed from JavaScript.

A [`Request`{.idl}](#request){#ref-for-request⑥ link-type="idl"} object
has an associated [request]{#concept-request-request .dfn .dfn-paneled
dfn-for="Request" dfn-type="dfn" export=""} (a
[request](#concept-request){#ref-for-concept-request①①⓪
link-type="dfn"}).

A [`Request`{.idl}](#request){#ref-for-request⑦ link-type="idl"} object
also has an associated [headers]{#request-headers .dfn .dfn-paneled
dfn-for="Request" dfn-type="dfn" export=""} (null or a
[`Headers`{.idl}](#headers){#ref-for-headers⑨ link-type="idl"} object),
initially null.

A [`Request`{.idl}](#request){#ref-for-request⑧ link-type="idl"} object
has an associated [signal]{#request-signal .dfn .dfn-paneled
dfn-for="Request" dfn-type="dfn" noexport=""} (null or an
[`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal②
link-type="idl"} object), initially null.

A [`Request`{.idl}](#request){#ref-for-request⑨ link-type="idl"}
object's [body](#concept-body-body){#ref-for-concept-body-body⑧
link-type="dfn"} is its
[request](#concept-request-request){#ref-for-concept-request-request①
link-type="dfn"}'s
[body](#concept-request-body){#ref-for-concept-request-body④④
link-type="dfn"}.

------------------------------------------------------------------------

`request`{.variable}` = new `[`Request`](#dom-request){#ref-for-dom-request① .idl-code link-type="constructor"}`(``input`{.variable}` [, ``init`{.variable}`])`

:   Returns a new `request`{.variable} whose
    [`url`{.idl}](#dom-request-url){#ref-for-dom-request-url①
    link-type="idl"} property is `input`{.variable} if
    `input`{.variable} is a string, and `input`{.variable}'s
    [`url`{.idl}](#dom-request-url){#ref-for-dom-request-url②
    link-type="idl"} if `input`{.variable} is a
    [`Request`{.idl}](#request){#ref-for-request①⓪ link-type="idl"}
    object.

    The `init`{.variable} argument is an object whose properties can be
    set as follows:

    [`method`{.idl}](#dom-requestinit-method){#ref-for-dom-requestinit-method link-type="idl"}
    :   A string to set `request`{.variable}'s
        [`method`{.idl}](#dom-request-method){#ref-for-dom-request-method①
        link-type="idl"}.

    [`headers`{.idl}](#dom-requestinit-headers){#ref-for-dom-requestinit-headers link-type="idl"}
    :   A [`Headers`{.idl}](#headers){#ref-for-headers①⓪
        link-type="idl"} object, an object literal, or an array of
        two-item arrays to set `request`{.variable}'s
        [`headers`{.idl}](#dom-request-headers){#ref-for-dom-request-headers①
        link-type="idl"}.

    [`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body link-type="idl"}
    :   A [`BodyInit`{.idl}](#bodyinit){#ref-for-bodyinit③
        link-type="idl"} object or null to set `request`{.variable}'s
        [body](#concept-request-body){#ref-for-concept-request-body④⑤
        link-type="dfn"}.

    [`referrer`{.idl}](#dom-requestinit-referrer){#ref-for-dom-requestinit-referrer link-type="idl"}
    :   A string whose value is a same-origin URL, \"`about:client`\",
        or the empty string, to set `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer⑧
        link-type="dfn"}.

    [`referrerPolicy`{.idl}](#dom-requestinit-referrerpolicy){#ref-for-dom-requestinit-referrerpolicy link-type="idl"}
    :   A [referrer
        policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#ref-for-referrer-policy①
        link-type="dfn"} to set `request`{.variable}'s
        [`referrerPolicy`{.idl}](#dom-request-referrerpolicy){#ref-for-dom-request-referrerpolicy①
        link-type="idl"}.

    [`mode`{.idl}](#dom-requestinit-mode){#ref-for-dom-requestinit-mode link-type="idl"}
    :   A string to indicate whether the request will use CORS, or will
        be restricted to same-origin URLs. Sets `request`{.variable}'s
        [`mode`{.idl}](#dom-request-mode){#ref-for-dom-request-mode①
        link-type="idl"}. If `input`{.variable} is a string, it defaults
        to \"`cors`\".

    [`credentials`{.idl}](#dom-requestinit-credentials){#ref-for-dom-requestinit-credentials link-type="idl"}
    :   A string indicating whether credentials will be sent with the
        request always, never, or only when sent to a same-origin URL
        --- as well as whether any credentials sent back in the response
        will be used always, never, or only when received from a
        same-origin URL. Sets `request`{.variable}'s
        [`credentials`{.idl}](#dom-request-credentials){#ref-for-dom-request-credentials①
        link-type="idl"}. If `input`{.variable} is a string, it defaults
        to \"`same-origin`\".

    [`cache`{.idl}](#dom-requestinit-cache){#ref-for-dom-requestinit-cache link-type="idl"}
    :   A string indicating how the request will interact with the
        browser's cache to set `request`{.variable}'s
        [`cache`{.idl}](#dom-request-cache){#ref-for-dom-request-cache①
        link-type="idl"}.

    [`redirect`{.idl}](#dom-requestinit-redirect){#ref-for-dom-requestinit-redirect link-type="idl"}
    :   A string indicating whether `request`{.variable} follows
        redirects, results in an error upon encountering a redirect, or
        returns the redirect (in an opaque fashion). Sets
        `request`{.variable}'s
        [`redirect`{.idl}](#dom-request-redirect){#ref-for-dom-request-redirect①
        link-type="idl"}.

    [`integrity`{.idl}](#dom-requestinit-integrity){#ref-for-dom-requestinit-integrity link-type="idl"}
    :   A cryptographic hash of the resource to be fetched by
        `request`{.variable}. Sets `request`{.variable}'s
        [`integrity`{.idl}](#dom-request-integrity){#ref-for-dom-request-integrity①
        link-type="idl"}.

    [`keepalive`{.idl}](#dom-requestinit-keepalive){#ref-for-dom-requestinit-keepalive link-type="idl"}
    :   A boolean to set `request`{.variable}'s
        [`keepalive`{.idl}](#dom-request-keepalive){#ref-for-dom-request-keepalive①
        link-type="idl"}.

    [`signal`{.idl}](#dom-requestinit-signal){#ref-for-dom-requestinit-signal link-type="idl"}
    :   An
        [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal③
        link-type="idl"} to set `request`{.variable}'s
        [`signal`{.idl}](#dom-request-signal){#ref-for-dom-request-signal①
        link-type="idl"}.

    [`window`{.idl}](#dom-requestinit-window){#ref-for-dom-requestinit-window link-type="idl"}
    :   Can only be null. Used to disassociate `request`{.variable} from
        any
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window②
        link-type="idl"}.

    [`duplex`{.idl}](#dom-requestinit-duplex){#ref-for-dom-requestinit-duplex link-type="idl"}
    :   \"`half`\" is the only valid value and it is for initiating a
        half-duplex fetch (i.e., the user agent sends the entire request
        before processing the response). \"`full`\" is reserved for
        future use, for initiating a full-duplex fetch (i.e., the user
        agent can process the response before sending the entire
        request). This member needs to be set when
        [`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body①
        link-type="idl"} is a
        [`ReadableStream`{.idl}](https://streams.spec.whatwg.org/#readablestream){#ref-for-readablestream①②
        link-type="idl"} object. [See [issue
        #1254](https://github.com/whatwg/fetch/issues/1254) for defining
        \"`full`\".]{.note}

    [`priority`{.idl}](#dom-requestinit-priority){#ref-for-dom-requestinit-priority link-type="idl"}
    :   A string to set `request`{.variable}'s
        [priority](#request-priority){#ref-for-request-priority①
        link-type="dfn"}.

`request`{.variable}` . `[`method`](#dom-request-method){#ref-for-dom-request-method② .idl-code link-type="attribute"}
:   Returns `request`{.variable}'s HTTP method, which is \"`GET`\" by
    default.

`request`{.variable}` . `[`url`](#dom-request-url){#ref-for-dom-request-url③ .idl-code link-type="attribute"}
:   Returns the URL of `request`{.variable} as a string.

`request`{.variable}` . `[`headers`](#dom-request-headers){#ref-for-dom-request-headers② .idl-code link-type="attribute"}
:   Returns a [`Headers`{.idl}](#headers){#ref-for-headers①①
    link-type="idl"} object consisting of the headers associated with
    `request`{.variable}. Note that headers added in the network layer
    by the user agent will not be accounted for in this object, e.g.,
    the \"`Host`\" header.

`request`{.variable}` . `[`destination`](#dom-request-destination){#ref-for-dom-request-destination① .idl-code link-type="attribute"}
:   Returns the kind of resource requested by `request`{.variable},
    e.g., \"`document`\" or \"`script`\".

`request`{.variable}` . `[`referrer`](#dom-request-referrer){#ref-for-dom-request-referrer① .idl-code link-type="attribute"}
:   Returns the referrer of `request`{.variable}. Its value can be a
    same-origin URL if explicitly set in `init`{.variable}, the empty
    string to indicate no referrer, and \"`about:client`\" when
    defaulting to the global's default. This is used during fetching to
    determine the value of the \``Referer`\` header of the request being
    made.

`request`{.variable}` . `[`referrerPolicy`](#dom-request-referrerpolicy){#ref-for-dom-request-referrerpolicy② .idl-code link-type="attribute"}
:   Returns the referrer policy associated with `request`{.variable}.
    This is used during fetching to compute the value of the
    `request`{.variable}'s referrer.

`request`{.variable}` . `[`mode`](#dom-request-mode){#ref-for-dom-request-mode② .idl-code link-type="attribute"}
:   Returns the
    [mode](#concept-request-mode){#ref-for-concept-request-mode②④
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating whether the request will use CORS, or will be
    restricted to same-origin URLs.

`request`{.variable}` . `[`credentials`](#dom-request-credentials){#ref-for-dom-request-credentials② .idl-code link-type="attribute"}
:   Returns the [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②⓪
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating whether credentials will be sent with the request
    always, never, or only when sent to a same-origin URL.

`request`{.variable}` . `[`cache`](#dom-request-cache){#ref-for-dom-request-cache② .idl-code link-type="attribute"}
:   Returns the [cache
    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①①
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating how the request will interact with the browser's
    cache when fetching.

`request`{.variable}` . `[`redirect`](#dom-request-redirect){#ref-for-dom-request-redirect② .idl-code link-type="attribute"}
:   Returns the [redirect
    mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode⑧
    link-type="dfn"} associated with `request`{.variable}, which is a
    string indicating how redirects for the request will be handled
    during fetching. A
    [request](#concept-request){#ref-for-concept-request①①①
    link-type="dfn"} will follow redirects by default.

`request`{.variable}` . `[`integrity`](#dom-request-integrity){#ref-for-dom-request-integrity② .idl-code link-type="attribute"}
:   Returns `request`{.variable}'s subresource integrity metadata, which
    is a cryptographic hash of the resource being fetched. Its value
    consists of multiple hashes separated by whitespace.
    [\[SRI\]](#biblio-sri "Subresource Integrity"){link-type="biblio"}

`request`{.variable}` . `[`keepalive`](#dom-request-keepalive){#ref-for-dom-request-keepalive② .idl-code link-type="attribute"}
:   Returns a boolean indicating whether or not `request`{.variable} can
    outlive the global in which it was created.

`request`{.variable}` . `[`isReloadNavigation`](#dom-request-isreloadnavigation){#ref-for-dom-request-isreloadnavigation① .idl-code link-type="attribute"}
:   Returns a boolean indicating whether or not `request`{.variable} is
    for a reload navigation.

`request`{.variable}` . `[`isHistoryNavigation`](#dom-request-ishistorynavigation){#ref-for-dom-request-ishistorynavigation① .idl-code link-type="attribute"}
:   Returns a boolean indicating whether or not `request`{.variable} is
    for a history navigation (a.k.a. back-foward navigation).

`request`{.variable}` . `[`signal`](#dom-request-signal){#ref-for-dom-request-signal② .idl-code link-type="attribute"}
:   Returns the signal associated with `request`{.variable}, which is an
    [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal④
    link-type="idl"} object indicating whether or not
    `request`{.variable} has been aborted, and its abort event handler.

`request`{.variable}` . `[`duplex`](#dom-request-duplex){#ref-for-dom-request-duplex① .idl-code link-type="attribute"}
:   Returns \"`half`\", meaning the fetch will be half-duplex (i.e., the
    user agent sends the entire request before processing the response).
    In future, it could also return \"`full`\", meaning the fetch will
    be full-duplex (i.e., the user agent can process the response before
    sending the entire request) to indicate that the fetch will be
    full-duplex. [See [issue
    #1254](https://github.com/whatwg/fetch/issues/1254) for defining
    \"`full`\".]{.note}

`request`{.variable}` . `[`clone`](#dom-request-clone){#ref-for-dom-request-clone① .idl-code link-type="method"}`()`

:   Returns a clone of `request`{.variable}.

------------------------------------------------------------------------

::: {.algorithm algorithm="create" algorithm-for="Request"}
To [create]{#request-create .dfn .dfn-paneled dfn-for="Request"
dfn-type="dfn" export="" lt="create|creating"} a
[`Request`{.idl}](#request){#ref-for-request①① link-type="idl"} object,
given a [request](#concept-request){#ref-for-concept-request①①②
link-type="dfn"} `request`{.variable}, [headers
guard](#headers-guard){#ref-for-headers-guard① link-type="dfn"}
`guard`{.variable},
[`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal⑤
link-type="idl"} object `signal`{.variable}, and
[realm](https://tc39.es/ecma262/#realm){#ref-for-realm② link-type="dfn"}
`realm`{.variable}:

1.  Let `requestObject`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new②
    link-type="dfn"} [`Request`{.idl}](#request){#ref-for-request①②
    link-type="idl"} object with `realm`{.variable}.

2.  Set `requestObject`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request②
    link-type="dfn"} to `request`{.variable}.

3.  Set `requestObject`{.variable}'s
    [headers](#request-headers){#ref-for-request-headers
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new③
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①②
    link-type="idl"} object with `realm`{.variable}, whose [headers
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①①
    link-type="dfn"} is `request`{.variable}'s [headers
    list](#concept-request-header-list){#ref-for-concept-request-header-list④⑥
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①⓪
    link-type="dfn"} is `guard`{.variable}.

4.  Set `requestObject`{.variable}'s
    [signal](#request-signal){#ref-for-request-signal link-type="dfn"}
    to `signal`{.variable}.

5.  Return `requestObject`{.variable}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="Request(input, init)" algorithm-for="Request"}
The
[`new Request(``input`{.variable}`, ``init`{.variable}`)`]{#dom-request
.dfn .dfn-paneled .idl-code dfn-for="Request" dfn-type="constructor"
export=""
lt="Request(input, init)|constructor(input, init)|Request(input)|constructor(input)"}
constructor steps are:

1.  Let `request`{.variable} be null.

2.  Let `fallbackMode`{.variable} be null.

3.  Let `baseURL`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③③
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object
    link-type="dfn"}'s [API base
    URL](https://html.spec.whatwg.org/multipage/webappapis.html#api-base-url){#ref-for-api-base-url
    link-type="dfn"}.

4.  Let `signal`{.variable} be null.

5.  If `input`{.variable} is a string, then:

    1.  Let `parsedURL`{.variable} be the result of
        [parsing](https://url.spec.whatwg.org/#concept-url-parser){#ref-for-concept-url-parser①
        link-type="dfn"} `input`{.variable} with `baseURL`{.variable}.

    2.  If `parsedURL`{.variable} is failure, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①②
        link-type="idl"}.

    3.  If `parsedURL`{.variable} [includes
        credentials](https://url.spec.whatwg.org/#include-credentials){#ref-for-include-credentials④
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⓪
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①③
        link-type="idl"}.

    4.  Set `request`{.variable} to a new
        [request](#concept-request){#ref-for-concept-request①①③
        link-type="dfn"} whose
        [URL](#concept-request-url){#ref-for-concept-request-url⑥
        link-type="dfn"} is `parsedURL`{.variable}.

    5.  Set `fallbackMode`{.variable} to \"`cors`\".

6.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑤
        link-type="dfn"}: `input`{.variable} is a
        [`Request`{.idl}](#request){#ref-for-request①③ link-type="idl"}
        object.

    2.  Set `request`{.variable} to `input`{.variable}'s
        [request](#concept-request-request){#ref-for-concept-request-request③
        link-type="dfn"}.

    3.  Set `signal`{.variable} to `input`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal①
        link-type="dfn"}.

7.  Let `origin`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③④
    link-type="dfn"}'s [relevant settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object①
    link-type="dfn"}'s
    [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin②
    link-type="dfn"}.

8.  Let `traversableForUserPrompts`{.variable} be \"`client`\".

9.  If `request`{.variable}'s [traversable for user
    prompts](#concept-request-window){#ref-for-concept-request-window①③
    link-type="dfn"} is an [environment settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#environment-settings-object){#ref-for-environment-settings-object⑨
    link-type="dfn"} and its
    [origin](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-origin){#ref-for-concept-settings-object-origin③
    link-type="dfn"} is [same
    origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin①⓪
    link-type="dfn"} with `origin`{.variable}, then set
    `traversableForUserPrompts`{.variable} to `request`{.variable}'s
    [traversable for user
    prompts](#concept-request-window){#ref-for-concept-request-window①④
    link-type="dfn"}.

10. If
    `init`{.variable}\[\"[`window`{.idl}](#dom-requestinit-window){#ref-for-dom-requestinit-window①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists③
    link-type="dfn"} and is non-null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①①
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①④
    link-type="idl"}.

11. If
    `init`{.variable}\[\"[`window`{.idl}](#dom-requestinit-window){#ref-for-dom-requestinit-window②
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists④
    link-type="dfn"}, then set `traversableForUserPrompts`{.variable} to
    \"`no-traversable`\".

12. Set `request`{.variable} to a new
    [request](#concept-request){#ref-for-concept-request①①④
    link-type="dfn"} with the following properties:

    [URL](#concept-request-url){#ref-for-concept-request-url⑦ link-type="dfn"}
    :   `request`{.variable}'s
        [URL](#concept-request-url){#ref-for-concept-request-url⑧
        link-type="dfn"}.

    [method](#concept-request-method){#ref-for-concept-request-method②⓪ link-type="dfn"}
    :   `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method②①
        link-type="dfn"}.

    [header list](#concept-request-header-list){#ref-for-concept-request-header-list④⑦ link-type="dfn"}
    :   A copy of `request`{.variable}'s [header
        list](#concept-request-header-list){#ref-for-concept-request-header-list④⑧
        link-type="dfn"}.

    [unsafe-request flag](#unsafe-request-flag){#ref-for-unsafe-request-flag③ link-type="dfn"}
    :   Set.

    [client](#concept-request-client){#ref-for-concept-request-client②⑨ link-type="dfn"}
    :   [This](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑤
        link-type="dfn"}'s [relevant settings
        object](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-settings-object){#ref-for-relevant-settings-object②
        link-type="dfn"}.

    [traversable for user prompts](#concept-request-window){#ref-for-concept-request-window①⑤ link-type="dfn"}
    :   `traversableForUserPrompts`{.variable}.

    [internal priority](#request-internal-priority){#ref-for-request-internal-priority② link-type="dfn"}
    :   `request`{.variable}'s [internal
        priority](#request-internal-priority){#ref-for-request-internal-priority③
        link-type="dfn"}.

    [origin](#concept-request-origin){#ref-for-concept-request-origin②② link-type="dfn"}
    :   `request`{.variable}'s
        [origin](#concept-request-origin){#ref-for-concept-request-origin②③
        link-type="dfn"}. [The propagation of the
        [origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin⑧
        link-type="dfn"} is only significant for navigation requests
        being handled by a service worker. In this scenario a request
        can have an origin that is different from the current
        client.]{.note}

    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer⑨ link-type="dfn"}
    :   `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⓪
        link-type="dfn"}.

    [referrer policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑥ link-type="dfn"}
    :   `request`{.variable}'s [referrer
        policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑦
        link-type="dfn"}.

    [mode](#concept-request-mode){#ref-for-concept-request-mode②⑤ link-type="dfn"}
    :   `request`{.variable}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode②⑥
        link-type="dfn"}.

    [credentials mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②① link-type="dfn"}
    :   `request`{.variable}'s [credentials
        mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②②
        link-type="dfn"}.

    [cache mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①② link-type="dfn"}
    :   `request`{.variable}'s [cache
        mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①③
        link-type="dfn"}.

    [redirect mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode⑨ link-type="dfn"}
    :   `request`{.variable}'s [redirect
        mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①⓪
        link-type="dfn"}.

    [integrity metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata③ link-type="dfn"}
    :   `request`{.variable}'s [integrity
        metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata④
        link-type="dfn"}.

    [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag③ link-type="dfn"}
    :   `request`{.variable}'s
        [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag④
        link-type="dfn"}.

    [reload-navigation flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag link-type="dfn"}
    :   `request`{.variable}'s [reload-navigation
        flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag①
        link-type="dfn"}.

    [history-navigation flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag link-type="dfn"}
    :   `request`{.variable}'s [history-navigation
        flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag①
        link-type="dfn"}.

    [URL list](#concept-request-url-list){#ref-for-concept-request-url-list⑨ link-type="dfn"}
    :   A
        [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone③
        link-type="dfn"} of `request`{.variable}'s [URL
        list](#concept-request-url-list){#ref-for-concept-request-url-list①⓪
        link-type="dfn"}.

    [initiator type](#request-initiator-type){#ref-for-request-initiator-type③ link-type="dfn"}
    :   \"`fetch`\".

13. If `init`{.variable} [is not
    empty](https://infra.spec.whatwg.org/#map-is-empty){#ref-for-map-is-empty
    link-type="dfn"}, then:

    1.  If `request`{.variable}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode②⑦
        link-type="dfn"} is \"`navigate`\", then set it to
        \"`same-origin`\".

    2.  Unset `request`{.variable}'s [reload-navigation
        flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag②
        link-type="dfn"}.

    3.  Unset `request`{.variable}'s [history-navigation
        flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag②
        link-type="dfn"}.

    4.  Set `request`{.variable}'s
        [origin](#concept-request-origin){#ref-for-concept-request-origin②④
        link-type="dfn"} to \"`client`\".

    5.  Set `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①①
        link-type="dfn"} to \"`client`\".

    6.  Set `request`{.variable}'s [referrer
        policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑧
        link-type="dfn"} to the empty string.

    7.  Set `request`{.variable}'s
        [URL](#concept-request-url){#ref-for-concept-request-url⑨
        link-type="dfn"} to `request`{.variable}'s [current
        URL](#concept-request-current-url){#ref-for-concept-request-current-url③⑦
        link-type="dfn"}.

    8.  Set `request`{.variable}'s [URL
        list](#concept-request-url-list){#ref-for-concept-request-url-list①①
        link-type="dfn"} to « `request`{.variable}'s
        [URL](#concept-request-url){#ref-for-concept-request-url①⓪
        link-type="dfn"} ».

    This is done to ensure that when a service worker \"redirects\" a
    request, e.g., from an image in a cross-origin style sheet, and
    makes modifications, it no longer appears to come from the original
    source (i.e., the cross-origin style sheet), but instead from the
    service worker that \"redirected\" the request. This is important as
    the original source might not even be able to generate the same kind
    of requests as the service worker. Services that trust the original
    source could therefore be exploited were this not done, although
    that is somewhat farfetched.

14. If
    `init`{.variable}\[\"[`referrer`{.idl}](#dom-requestinit-referrer){#ref-for-dom-requestinit-referrer①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑤
    link-type="dfn"}, then:

    1.  Let `referrer`{.variable} be
        `init`{.variable}\[\"[`referrer`{.idl}](#dom-requestinit-referrer){#ref-for-dom-requestinit-referrer②
        link-type="idl"}\"\].

    2.  If `referrer`{.variable} is the empty string, then set
        `request`{.variable}'s
        [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①②
        link-type="dfn"} to \"`no-referrer`\".

    3.  Otherwise:

        1.  Let `parsedReferrer`{.variable} be the result of
            [parsing](https://url.spec.whatwg.org/#concept-url-parser){#ref-for-concept-url-parser②
            link-type="dfn"} `referrer`{.variable} with
            `baseURL`{.variable}.

        2.  If `parsedReferrer`{.variable} is failure, then
            [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①②
            link-type="dfn"} a
            [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑤
            link-type="idl"}.

        3.  If one of the following is true

            - `parsedReferrer`{.variable}'s
              [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#ref-for-concept-url-scheme①⑤
              link-type="dfn"} is \"`about`\" and
              [path](https://url.spec.whatwg.org/#concept-url-path){#ref-for-concept-url-path②
              link-type="dfn"} is the string \"`client`\"

            - `parsedReferrer`{.variable}'s
              [origin](https://url.spec.whatwg.org/#concept-url-origin){#ref-for-concept-url-origin②⓪
              link-type="dfn"} is not [same
              origin](https://html.spec.whatwg.org/multipage/browsers.html#same-origin){#ref-for-same-origin①①
              link-type="dfn"} with `origin`{.variable}

            then set `request`{.variable}'s
            [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①③
            link-type="dfn"} to \"`client`\".

        4.  Otherwise, set `request`{.variable}'s
            [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①④
            link-type="dfn"} to `parsedReferrer`{.variable}.

15. If
    `init`{.variable}\[\"[`referrerPolicy`{.idl}](#dom-requestinit-referrerpolicy){#ref-for-dom-requestinit-referrerpolicy①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑥
    link-type="dfn"}, then set `request`{.variable}'s [referrer
    policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy⑨
    link-type="dfn"} to it.

16. Let `mode`{.variable} be
    `init`{.variable}\[\"[`mode`{.idl}](#dom-requestinit-mode){#ref-for-dom-requestinit-mode①
    link-type="idl"}\"\] if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑦
    link-type="dfn"}, and `fallbackMode`{.variable} otherwise.

17. If `mode`{.variable} is \"`navigate`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①③
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑥
    link-type="idl"}.

18. If `mode`{.variable} is non-null, set `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode②⑧
    link-type="dfn"} to `mode`{.variable}.

19. If
    `init`{.variable}\[\"[`credentials`{.idl}](#dom-requestinit-credentials){#ref-for-dom-requestinit-credentials①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑧
    link-type="dfn"}, then set `request`{.variable}'s [credentials
    mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②③
    link-type="dfn"} to it.

20. If
    `init`{.variable}\[\"[`cache`{.idl}](#dom-requestinit-cache){#ref-for-dom-requestinit-cache①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑨
    link-type="dfn"}, then set `request`{.variable}'s [cache
    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①④
    link-type="dfn"} to it.

21. If `request`{.variable}'s [cache
    mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①⑤
    link-type="dfn"} is \"`only-if-cached`\" and `request`{.variable}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode②⑨
    link-type="dfn"} is *not* \"`same-origin`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①④
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑦
    link-type="idl"}.

22. If
    `init`{.variable}\[\"[`redirect`{.idl}](#dom-requestinit-redirect){#ref-for-dom-requestinit-redirect①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⓪
    link-type="dfn"}, then set `request`{.variable}'s [redirect
    mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①①
    link-type="dfn"} to it.

23. If
    `init`{.variable}\[\"[`integrity`{.idl}](#dom-requestinit-integrity){#ref-for-dom-requestinit-integrity①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①①
    link-type="dfn"}, then set `request`{.variable}'s [integrity
    metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata⑤
    link-type="dfn"} to it.

24. If
    `init`{.variable}\[\"[`keepalive`{.idl}](#dom-requestinit-keepalive){#ref-for-dom-requestinit-keepalive①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①②
    link-type="dfn"}, then set `request`{.variable}'s
    [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag⑤
    link-type="dfn"} to it.

25. If
    `init`{.variable}\[\"[`method`{.idl}](#dom-requestinit-method){#ref-for-dom-requestinit-method①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①③
    link-type="dfn"}, then:

    1.  Let `method`{.variable} be
        `init`{.variable}\[\"[`method`{.idl}](#dom-requestinit-method){#ref-for-dom-requestinit-method②
        link-type="idl"}\"\].

    2.  If `method`{.variable} is not a
        [method](#concept-method){#ref-for-concept-method①②
        link-type="dfn"} or `method`{.variable} is a [forbidden
        method](#forbidden-method){#ref-for-forbidden-method②
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑤
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑧
        link-type="idl"}.

    3.  [Normalize](#concept-method-normalize){#ref-for-concept-method-normalize②
        link-type="dfn"} `method`{.variable}.

    4.  Set `request`{.variable}'s
        [method](#concept-request-method){#ref-for-concept-request-method②②
        link-type="dfn"} to `method`{.variable}.

26. If
    `init`{.variable}\[\"[`signal`{.idl}](#dom-requestinit-signal){#ref-for-dom-requestinit-signal①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①④
    link-type="dfn"}, then set `signal`{.variable} to it.

27. If
    `init`{.variable}\[\"[`priority`{.idl}](#dom-requestinit-priority){#ref-for-dom-requestinit-priority①
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑤
    link-type="dfn"}, then:

    1.  If `request`{.variable}'s [internal
        priority](#request-internal-priority){#ref-for-request-internal-priority④
        link-type="dfn"} is not null, then update `request`{.variable}'s
        [internal
        priority](#request-internal-priority){#ref-for-request-internal-priority⑤
        link-type="dfn"} in an
        [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#ref-for-implementation-defined①⑨
        link-type="dfn"} manner.

    2.  Otherwise, set `request`{.variable}'s
        [priority](#request-priority){#ref-for-request-priority②
        link-type="dfn"} to
        `init`{.variable}\[\"[`priority`{.idl}](#dom-requestinit-priority){#ref-for-dom-requestinit-priority②
        link-type="idl"}\"\].

28. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑥
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request④
    link-type="dfn"} to `request`{.variable}.

29. Let `signals`{.variable} be « `signal`{.variable} » if
    `signal`{.variable} is non-null; otherwise « ».

30. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑦
    link-type="dfn"}'s
    [signal](#request-signal){#ref-for-request-signal② link-type="dfn"}
    to the result of [creating a dependent abort
    signal](https://dom.spec.whatwg.org/#create-a-dependent-abort-signal){#ref-for-create-a-dependent-abort-signal
    link-type="dfn"} from `signals`{.variable}, using
    [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal⑥
    link-type="idl"} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑧
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm②
    link-type="dfn"}.

31. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑨
    link-type="dfn"}'s
    [headers](#request-headers){#ref-for-request-headers①
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new④
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①③
    link-type="idl"} object with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⓪
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm③
    link-type="dfn"}, whose [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①②
    link-type="dfn"} is `request`{.variable}'s [header
    list](#concept-request-header-list){#ref-for-concept-request-header-list④⑨
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①①
    link-type="dfn"} is \"`request`\".

32. If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④①
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request⑤
    link-type="dfn"}'s
    [mode](#concept-request-mode){#ref-for-concept-request-mode③⓪
    link-type="dfn"} is \"`no-cors`\", then:

    1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④②
        link-type="dfn"}'s
        [request](#concept-request-request){#ref-for-concept-request-request⑥
        link-type="dfn"}'s
        [method](#concept-request-method){#ref-for-concept-request-method②③
        link-type="dfn"} is not a [CORS-safelisted
        method](#cors-safelisted-method){#ref-for-cors-safelisted-method④
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑥
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror①⑨
        link-type="idl"}.

    2.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④③
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers②
        link-type="dfn"}'s
        [guard](#concept-headers-guard){#ref-for-concept-headers-guard①②
        link-type="dfn"} to \"`request-no-cors`\".

33. If `init`{.variable} [is not
    empty](https://infra.spec.whatwg.org/#map-is-empty){#ref-for-map-is-empty①
    link-type="dfn"}, then:

    The headers are sanitized as they might contain headers that are not
    allowed by this mode. Otherwise, they were previously sanitized or
    are unmodified since they were set by a privileged API.

    1.  Let `headers`{.variable} be a copy of
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④④
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers③
        link-type="dfn"} and its associated [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①③
        link-type="dfn"}.

    2.  If
        `init`{.variable}\[\"[`headers`{.idl}](#dom-requestinit-headers){#ref-for-dom-requestinit-headers①
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑥
        link-type="dfn"}, then set `headers`{.variable} to
        `init`{.variable}\[\"[`headers`{.idl}](#dom-requestinit-headers){#ref-for-dom-requestinit-headers②
        link-type="idl"}\"\].

    3.  Empty
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑤
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers④
        link-type="dfn"}'s [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①④
        link-type="dfn"}.

    4.  If `headers`{.variable} is a
        [`Headers`{.idl}](#headers){#ref-for-headers①④ link-type="idl"}
        object, then [for
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑥
        link-type="dfn"} `header`{.variable} of its [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑤
        link-type="dfn"},
        [append](#concept-headers-append){#ref-for-concept-headers-append③
        link-type="dfn"} `header`{.variable} to
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑥
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑤
        link-type="dfn"}.

    5.  Otherwise,
        [fill](#concept-headers-fill){#ref-for-concept-headers-fill①
        link-type="dfn"}
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑦
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑥
        link-type="dfn"} with `headers`{.variable}.

34. Let `inputBody`{.variable} be `input`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request⑦
    link-type="dfn"}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑥
    link-type="dfn"} if `input`{.variable} is a
    [`Request`{.idl}](#request){#ref-for-request①④ link-type="idl"}
    object; otherwise null.

35. If either
    `init`{.variable}\[\"[`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body②
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑦
    link-type="dfn"} and is non-null or `inputBody`{.variable} is
    non-null, and `request`{.variable}'s
    [method](#concept-request-method){#ref-for-concept-request-method②④
    link-type="dfn"} is \``GET`\` or \``HEAD`\`, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑦
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⓪
    link-type="idl"}.

36. Let `initBody`{.variable} be null.

37. If
    `init`{.variable}\[\"[`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body③
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑧
    link-type="dfn"} and is non-null, then:

    1.  Let `bodyWithType`{.variable} be the result of
        [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract②
        link-type="dfn"}
        `init`{.variable}\[\"[`body`{.idl}](#dom-requestinit-body){#ref-for-dom-requestinit-body④
        link-type="idl"}\"\], with
        [*keepalive*](#keepalive){#ref-for-keepalive link-type="dfn"}
        set to `request`{.variable}'s
        [keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag⑥
        link-type="dfn"}.

    2.  Set `initBody`{.variable} to `bodyWithType`{.variable}'s
        [body](#body-with-type-body){#ref-for-body-with-type-body⑤
        link-type="dfn"}.

    3.  Let `type`{.variable} be `bodyWithType`{.variable}'s
        [type](#body-with-type-type){#ref-for-body-with-type-type
        link-type="dfn"}.

    4.  If `type`{.variable} is non-null and
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑧
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑦
        link-type="dfn"}'s [header
        list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑥
        link-type="dfn"} [does not
        contain](#header-list-contains){#ref-for-header-list-contains②⑥
        link-type="dfn"} \``Content-Type`\`, then
        [append](#concept-headers-append){#ref-for-concept-headers-append④
        link-type="dfn"} (\``Content-Type`\`, `type`{.variable}) to
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑨
        link-type="dfn"}'s
        [headers](#request-headers){#ref-for-request-headers⑧
        link-type="dfn"}.

38. Let `inputOrInitBody`{.variable} be `initBody`{.variable} if it is
    non-null; otherwise `inputBody`{.variable}.

39. If `inputOrInitBody`{.variable} is non-null and
    `inputOrInitBody`{.variable}'s
    [source](#concept-body-source){#ref-for-concept-body-source①④
    link-type="dfn"} is null, then:

    1.  If `initBody`{.variable} is non-null and
        `init`{.variable}\[\"[`duplex`{.idl}](#dom-requestinit-duplex){#ref-for-dom-requestinit-duplex①
        link-type="idl"}\"\] does not
        [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⑨
        link-type="dfn"}, then throw a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②①
        link-type="idl"}.

    2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⓪
        link-type="dfn"}'s
        [request](#concept-request-request){#ref-for-concept-request-request⑧
        link-type="dfn"}'s
        [mode](#concept-request-mode){#ref-for-concept-request-mode③①
        link-type="dfn"} is neither \"`same-origin`\" nor \"`cors`\",
        then throw a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②②
        link-type="idl"}.

    3.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤①
        link-type="dfn"}'s
        [request](#concept-request-request){#ref-for-concept-request-request⑨
        link-type="dfn"}'s [use-CORS-preflight
        flag](#use-cors-preflight-flag){#ref-for-use-cors-preflight-flag⑥
        link-type="dfn"}.

40. Let `finalBody`{.variable} be `inputOrInitBody`{.variable}.

41. If `initBody`{.variable} is null and `inputBody`{.variable} is
    non-null, then:

    1.  If `input`{.variable} is
        [unusable](#body-unusable){#ref-for-body-unusable①
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑧
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②③
        link-type="idl"}.

    2.  Set `finalBody`{.variable} to the result of [creating a
        proxy](https://streams.spec.whatwg.org/#readablestream-create-a-proxy){#ref-for-readablestream-create-a-proxy
        link-type="dfn"} for `inputBody`{.variable}.

42. Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤②
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①⓪
    link-type="dfn"}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑦
    link-type="dfn"} to `finalBody`{.variable}.
:::

The [`method`]{#dom-request-method .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤③
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①①
link-type="dfn"}'s
[method](#concept-request-method){#ref-for-concept-request-method②⑤
link-type="dfn"}.

The [`url`]{#dom-request-url .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤④
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①②
link-type="dfn"}'s
[URL](#concept-request-url){#ref-for-concept-request-url①①
link-type="dfn"},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer②
link-type="dfn"}.

The [`headers`]{#dom-request-headers .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑤
link-type="dfn"}'s [headers](#request-headers){#ref-for-request-headers⑨
link-type="dfn"}.

The [`destination`]{#dom-request-destination .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑥
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①③
link-type="dfn"}'s
[destination](#concept-request-destination){#ref-for-concept-request-destination②①
link-type="dfn"}.

::: {.algorithm algorithm="referrer" algorithm-for="Request"}
The [`referrer`]{#dom-request-referrer .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑦
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①④
    link-type="dfn"}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⑤
    link-type="dfn"} is \"`no-referrer`\", then return the empty string.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑧
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①⑤
    link-type="dfn"}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⑥
    link-type="dfn"} is \"`client`\", then return \"`about:client`\".

3.  Return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑨
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request①⑥
    link-type="dfn"}'s
    [referrer](#concept-request-referrer){#ref-for-concept-request-referrer①⑦
    link-type="dfn"},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer③
    link-type="dfn"}.
:::

The [`referrerPolicy`]{#dom-request-referrerpolicy .dfn .dfn-paneled
.idl-code dfn-for="Request" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⓪
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①⑦
link-type="dfn"}'s [referrer
policy](#concept-request-referrer-policy){#ref-for-concept-request-referrer-policy①⓪
link-type="dfn"}.

The [`mode`]{#dom-request-mode .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥①
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①⑧
link-type="dfn"}'s
[mode](#concept-request-mode){#ref-for-concept-request-mode③②
link-type="dfn"}.

The [`credentials`]{#dom-request-credentials .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥②
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request①⑨
link-type="dfn"}'s [credentials
mode](#concept-request-credentials-mode){#ref-for-concept-request-credentials-mode②④
link-type="dfn"}.

The [`cache`]{#dom-request-cache .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥③
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②⓪
link-type="dfn"}'s [cache
mode](#concept-request-cache-mode){#ref-for-concept-request-cache-mode①⑥
link-type="dfn"}.

The [`redirect`]{#dom-request-redirect .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥④
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②①
link-type="dfn"}'s [redirect
mode](#concept-request-redirect-mode){#ref-for-concept-request-redirect-mode①②
link-type="dfn"}.

The [`integrity`]{#dom-request-integrity .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑤
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②②
link-type="dfn"}'s [integrity
metadata](#concept-request-integrity-metadata){#ref-for-concept-request-integrity-metadata⑥
link-type="dfn"}.

The [`keepalive`]{#dom-request-keepalive .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑥
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②③
link-type="dfn"}'s
[keepalive](#request-keepalive-flag){#ref-for-request-keepalive-flag⑦
link-type="dfn"}.

The [`isReloadNavigation`]{#dom-request-isreloadnavigation .dfn
.dfn-paneled .idl-code dfn-for="Request" dfn-type="attribute" export=""}
getter steps are to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑦
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②④
link-type="dfn"}'s [reload-navigation
flag](#concept-request-reload-navigation-flag){#ref-for-concept-request-reload-navigation-flag③
link-type="dfn"} is set; otherwise false.

The [`isHistoryNavigation`]{#dom-request-ishistorynavigation .dfn
.dfn-paneled .idl-code dfn-for="Request" dfn-type="attribute" export=""}
getter steps are to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑧
link-type="dfn"}'s
[request](#concept-request-request){#ref-for-concept-request-request②⑤
link-type="dfn"}'s [history-navigation
flag](#concept-request-history-navigation-flag){#ref-for-concept-request-history-navigation-flag③
link-type="dfn"} is set; otherwise false.

The [`signal`]{#dom-request-signal .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑨
link-type="dfn"}'s [signal](#request-signal){#ref-for-request-signal③
link-type="dfn"}.

The [`duplex`]{#dom-request-duplex .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="attribute" export=""} getter steps are to
return \"`half`\".

------------------------------------------------------------------------

::: {.algorithm algorithm="clone()" algorithm-for="Request"}
The [`clone()`]{#dom-request-clone .dfn .dfn-paneled .idl-code
dfn-for="Request" dfn-type="method" export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⓪
    link-type="dfn"} is
    [unusable](#body-unusable){#ref-for-body-unusable② link-type="dfn"},
    then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑨
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②④
    link-type="idl"}.

2.  Let `clonedRequest`{.variable} be the result of
    [cloning](#concept-request-clone){#ref-for-concept-request-clone③
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦①
    link-type="dfn"}'s
    [request](#concept-request-request){#ref-for-concept-request-request②⑥
    link-type="dfn"}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑥
    link-type="dfn"}:
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦②
    link-type="dfn"}'s
    [signal](#request-signal){#ref-for-request-signal④ link-type="dfn"}
    is non-null.

4.  Let `clonedSignal`{.variable} be the result of [creating a dependent
    abort
    signal](https://dom.spec.whatwg.org/#create-a-dependent-abort-signal){#ref-for-create-a-dependent-abort-signal①
    link-type="dfn"} from «
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦③
    link-type="dfn"}'s
    [signal](#request-signal){#ref-for-request-signal⑤ link-type="dfn"}
    », using
    [`AbortSignal`{.idl}](https://dom.spec.whatwg.org/#abortsignal){#ref-for-abortsignal⑦
    link-type="idl"} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦④
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm④
    link-type="dfn"}.

5.  Let `clonedRequestObject`{.variable} be the result of
    [creating](#request-create){#ref-for-request-create link-type="dfn"}
    a [`Request`{.idl}](#request){#ref-for-request①⑤ link-type="idl"}
    object, given `clonedRequest`{.variable},
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑤
    link-type="dfn"}'s
    [headers](#request-headers){#ref-for-request-headers①⓪
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①③
    link-type="dfn"}, `clonedSignal`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑥
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑤
    link-type="dfn"}.

6.  Return `clonedRequestObject`{.variable}.
:::

### [5.5. ]{.secno}[Response class]{.content}[](#response-class){.self-link} {#response-class .heading .settled level="5.5"}

``` {.idl .highlight .def}
[Exposed=(Window,Worker)]interface Response {
  constructor(optional BodyInit? body = null, optional ResponseInit init = {});

  [NewObject] static Response error();
  [NewObject] static Response redirect(USVString url, optional unsigned short status = 302);
  [NewObject] static Response json(any data, optional ResponseInit init = {});

  readonly attribute ResponseType type;

  readonly attribute USVString url;
  readonly attribute boolean redirected;
  readonly attribute unsigned short status;
  readonly attribute boolean ok;
  readonly attribute ByteString statusText;
  [SameObject] readonly attribute Headers headers;

  [NewObject] Response clone();
};
Response includes Body;

dictionary ResponseInit {
  unsigned short status = 200;
  ByteString statusText = "";
  HeadersInit headers;
};

enum ResponseType { "basic", "cors", "default", "error", "opaque", "opaqueredirect" };
```

A [`Response`{.idl}](#response){#ref-for-response⑦ link-type="idl"}
object has an associated [response]{#concept-response-response .dfn
.dfn-paneled dfn-for="Response" dfn-type="dfn" export=""} (a
[response](#concept-response){#ref-for-concept-response⑥①
link-type="dfn"}).

A [`Response`{.idl}](#response){#ref-for-response⑧ link-type="idl"}
object also has an associated [headers]{#response-headers .dfn
.dfn-paneled dfn-for="Response" dfn-type="dfn" export=""} (null or a
[`Headers`{.idl}](#headers){#ref-for-headers①⑥ link-type="idl"} object),
initially null.

A [`Response`{.idl}](#response){#ref-for-response⑨ link-type="idl"}
object's [body](#concept-body-body){#ref-for-concept-body-body⑨
link-type="dfn"} is its
[response](#concept-response-response){#ref-for-concept-response-response①
link-type="dfn"}'s
[body](#concept-response-body){#ref-for-concept-response-body②⑦
link-type="dfn"}.

------------------------------------------------------------------------

`response`{.variable}` = new `[`Response`](#dom-response){#ref-for-dom-response① .idl-code link-type="constructor"}`(``body`{.variable}` = null [, ``init`{.variable}`])`

:   Creates a [`Response`{.idl}](#response){#ref-for-response①⓪
    link-type="idl"} whose body is `body`{.variable}, and status, status
    message, and headers are provided by `init`{.variable}.

`response`{.variable}` = `[`Response`](#response){#ref-for-response①① link-type="idl"}` . `[`error`](#dom-response-error){#ref-for-dom-response-error① .idl-code link-type="method"}`()`

:   Creates network error
    [`Response`{.idl}](#response){#ref-for-response①② link-type="idl"}.

`response`{.variable}` = `[`Response`](#response){#ref-for-response①③ link-type="idl"}` . `[`redirect`](#dom-response-redirect){#ref-for-dom-response-redirect① .idl-code link-type="method"}`(``url`{.variable}`, ``status`{.variable}` = 302)`

:   Creates a redirect [`Response`{.idl}](#response){#ref-for-response①④
    link-type="idl"} that redirects to `url`{.variable} with status
    `status`{.variable}.

`response`{.variable}` = `[`Response`](#response){#ref-for-response①⑤ link-type="idl"}` . `[`json`](#dom-response-json){#ref-for-dom-response-json① .idl-code link-type="method"}`(``data`{.variable}` [, ``init`{.variable}`])`

:   Creates a [`Response`{.idl}](#response){#ref-for-response①⑥
    link-type="idl"} whose body is the JSON-encoded `data`{.variable},
    and status, status message, and headers are provided by
    `init`{.variable}.

`response`{.variable}` . `[`type`](#dom-response-type){#ref-for-dom-response-type② .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s type, e.g., \"`cors`\".

`response`{.variable}` . `[`url`](#dom-response-url){#ref-for-dom-response-url① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s URL, if it has one; otherwise the
    empty string.

`response`{.variable}` . `[`redirected`](#dom-response-redirected){#ref-for-dom-response-redirected① .idl-code link-type="attribute"}

:   Returns whether `response`{.variable} was obtained through a
    redirect.

`response`{.variable}` . `[`status`](#dom-response-status){#ref-for-dom-response-status① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s status.

`response`{.variable}` . `[`ok`](#dom-response-ok){#ref-for-dom-response-ok② .idl-code link-type="attribute"}

:   Returns whether `response`{.variable}'s status is an [ok
    status](#ok-status){#ref-for-ok-status② link-type="dfn"}.

`response`{.variable}` . `[`statusText`](#dom-response-statustext){#ref-for-dom-response-statustext① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s status message.

`response`{.variable}` . `[`headers`](#dom-response-headers){#ref-for-dom-response-headers① .idl-code link-type="attribute"}

:   Returns `response`{.variable}'s headers as
    [`Headers`{.idl}](#headers){#ref-for-headers①⑦ link-type="idl"}.

`response`{.variable}` . `[`clone`](#dom-response-clone){#ref-for-dom-response-clone① .idl-code link-type="method"}`()`

:   Returns a clone of `response`{.variable}.

------------------------------------------------------------------------

::: {.algorithm algorithm="create" algorithm-for="Response"}
To [create]{#response-create .dfn .dfn-paneled dfn-for="Response"
dfn-type="dfn" export="" lt="create|creating"} a
[`Response`{.idl}](#response){#ref-for-response①⑦ link-type="idl"}
object, given a
[response](#concept-response){#ref-for-concept-response⑥②
link-type="dfn"} `response`{.variable}, [headers
guard](#headers-guard){#ref-for-headers-guard② link-type="dfn"}
`guard`{.variable}, and
[realm](https://tc39.es/ecma262/#realm){#ref-for-realm③ link-type="dfn"}
`realm`{.variable}, run these steps:

1.  Let `responseObject`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new⑤
    link-type="dfn"} [`Response`{.idl}](#response){#ref-for-response①⑧
    link-type="idl"} object with `realm`{.variable}.

2.  Set `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response②
    link-type="dfn"} to `response`{.variable}.

3.  Set `responseObject`{.variable}'s
    [headers](#response-headers){#ref-for-response-headers
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new⑥
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①⑧
    link-type="idl"} object with `realm`{.variable}, whose [headers
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑦
    link-type="dfn"} is `response`{.variable}'s [headers
    list](#concept-response-header-list){#ref-for-concept-response-header-list③③
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①④
    link-type="dfn"} is `guard`{.variable}.

4.  Return `responseObject`{.variable}.
:::

::: {.algorithm algorithm="initialize a response"}
To [initialize a response]{#initialize-a-response .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a
[`Response`{.idl}](#response){#ref-for-response①⑨ link-type="idl"}
object `response`{.variable},
[`ResponseInit`{.idl}](#responseinit){#ref-for-responseinit②
link-type="idl"} `init`{.variable}, and null or a [body with
type](#body-with-type){#ref-for-body-with-type② link-type="dfn"}
`body`{.variable}:

1.  If
    `init`{.variable}\[\"[`status`{.idl}](#dom-responseinit-status){#ref-for-dom-responseinit-status
    link-type="idl"}\"\] is not in the range 200 to 599, inclusive, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⓪
    link-type="dfn"} a
    [`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror②
    link-type="idl"}.

2.  If
    `init`{.variable}\[\"[`statusText`{.idl}](#dom-responseinit-statustext){#ref-for-dom-responseinit-statustext
    link-type="idl"}\"\] is not the empty string and does not match the
    [reason-phrase](https://httpwg.org/specs/rfc9112.html#status.line){#ref-for-status.line
    link-type="dfn"} token production, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②①
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑤
    link-type="idl"}.

3.  Set `response`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response③
    link-type="dfn"}'s
    [status](#concept-response-status){#ref-for-concept-response-status②①
    link-type="dfn"} to
    `init`{.variable}\[\"[`status`{.idl}](#dom-responseinit-status){#ref-for-dom-responseinit-status①
    link-type="idl"}\"\].

4.  Set `response`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response④
    link-type="dfn"}'s [status
    message](#concept-response-status-message){#ref-for-concept-response-status-message⑦
    link-type="dfn"} to
    `init`{.variable}\[\"[`statusText`{.idl}](#dom-responseinit-statustext){#ref-for-dom-responseinit-statustext①
    link-type="idl"}\"\].

5.  If
    `init`{.variable}\[\"[`headers`{.idl}](#dom-responseinit-headers){#ref-for-dom-responseinit-headers
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists②⓪
    link-type="dfn"}, then
    [fill](#concept-headers-fill){#ref-for-concept-headers-fill②
    link-type="dfn"} `response`{.variable}'s
    [headers](#response-headers){#ref-for-response-headers①
    link-type="dfn"} with
    `init`{.variable}\[\"[`headers`{.idl}](#dom-responseinit-headers){#ref-for-dom-responseinit-headers①
    link-type="idl"}\"\].

6.  If `body`{.variable} is non-null, then:

    1.  If `response`{.variable}'s
        [status](#concept-response-status){#ref-for-concept-response-status②②
        link-type="dfn"} is a [null body
        status](#null-body-status){#ref-for-null-body-status①
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②②
        link-type="dfn"} a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑥
        link-type="idl"}.

        101 and 103 are included in [null body
        status](#null-body-status){#ref-for-null-body-status②
        link-type="dfn"} due to their use elsewhere. They do not affect
        this step.

    2.  Set `response`{.variable}'s
        [body](#concept-response-body){#ref-for-concept-response-body②⑧
        link-type="dfn"} to `body`{.variable}'s
        [body](#body-with-type-body){#ref-for-body-with-type-body⑥
        link-type="dfn"}.

    3.  If `body`{.variable}'s
        [type](#body-with-type-type){#ref-for-body-with-type-type①
        link-type="dfn"} is non-null and `response`{.variable}'s [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list③④
        link-type="dfn"} [does not
        contain](#header-list-contains){#ref-for-header-list-contains②⑦
        link-type="dfn"} \``Content-Type`\`, then
        [append](#concept-header-list-append){#ref-for-concept-header-list-append②②
        link-type="dfn"} (\``Content-Type`\`, `body`{.variable}'s
        [type](#body-with-type-type){#ref-for-body-with-type-type②
        link-type="dfn"}) to `response`{.variable}'s [header
        list](#concept-response-header-list){#ref-for-concept-response-header-list③⑤
        link-type="dfn"}.
:::

------------------------------------------------------------------------

::: {.algorithm algorithm="Response(body, init)" algorithm-for="Response"}
The
[`new Response(``body`{.variable}`, ``init`{.variable}`)`]{#dom-response
.dfn .dfn-paneled .idl-code dfn-for="Response" dfn-type="constructor"
export=""
lt="Response(body, init)|constructor(body, init)|Response(body)|constructor(body)|Response()|constructor()"}
constructor steps are:

1.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑦
    link-type="dfn"}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑤
    link-type="dfn"} to a new
    [response](#concept-response){#ref-for-concept-response⑥③
    link-type="dfn"}.

2.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑧
    link-type="dfn"}'s
    [headers](#response-headers){#ref-for-response-headers②
    link-type="dfn"} to a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new⑦
    link-type="dfn"} [`Headers`{.idl}](#headers){#ref-for-headers①⑨
    link-type="idl"} object with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑨
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑥
    link-type="dfn"}, whose [header
    list](#concept-headers-header-list){#ref-for-concept-headers-header-list①⑧
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⓪
    link-type="dfn"}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑥
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③⑥
    link-type="dfn"} and
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①⑤
    link-type="dfn"} is \"`response`\".

3.  Let `bodyWithType`{.variable} be null.

4.  If `body`{.variable} is non-null, then set `bodyWithType`{.variable}
    to the result of
    [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract③
    link-type="dfn"} `body`{.variable}.

5.  Perform [initialize a
    response](#initialize-a-response){#ref-for-initialize-a-response
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧①
    link-type="dfn"}, `init`{.variable}, and `bodyWithType`{.variable}.
:::

The static [`error()`]{#dom-response-error .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="method" export=""} method steps are to
return the result of
[creating](#response-create){#ref-for-response-create link-type="dfn"} a
[`Response`{.idl}](#response){#ref-for-response②⓪ link-type="idl"}
object, given a new [network
error](#concept-network-error){#ref-for-concept-network-error⑤⑨
link-type="dfn"}, \"`immutable`\", and the [current
realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm
link-type="dfn"}.

::: {.algorithm algorithm="redirect(url, status)" algorithm-for="Response"}
The static
[`redirect(``url`{.variable}`, ``status`{.variable}`)`]{#dom-response-redirect
.dfn .dfn-paneled .idl-code dfn-for="Response" dfn-type="method"
export="" lt="redirect(url, status)|redirect(url)"} method steps are:

1.  Let `parsedURL`{.variable} be the result of
    [parsing](https://url.spec.whatwg.org/#concept-url-parser){#ref-for-concept-url-parser③
    link-type="dfn"} `url`{.variable} with [current settings
    object](https://html.spec.whatwg.org/multipage/webappapis.html#current-settings-object){#ref-for-current-settings-object
    link-type="dfn"}'s [API base
    URL](https://html.spec.whatwg.org/multipage/webappapis.html#api-base-url){#ref-for-api-base-url①
    link-type="dfn"}.

2.  If `parsedURL`{.variable} is failure, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②③
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑦
    link-type="idl"}.

3.  If `status`{.variable} is not a [redirect
    status](#redirect-status){#ref-for-redirect-status③
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②④
    link-type="dfn"} a
    [`RangeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-rangeerror){#ref-for-exceptiondef-rangeerror③
    link-type="idl"}.

4.  Let `responseObject`{.variable} be the result of
    [creating](#response-create){#ref-for-response-create①
    link-type="dfn"} a [`Response`{.idl}](#response){#ref-for-response②①
    link-type="idl"} object, given a new
    [response](#concept-response){#ref-for-concept-response⑥④
    link-type="dfn"}, \"`immutable`\", and the [current
    realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm①
    link-type="dfn"}.

5.  Set `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑦
    link-type="dfn"}'s
    [status](#concept-response-status){#ref-for-concept-response-status②③
    link-type="dfn"} to `status`{.variable}.

6.  Let `value`{.variable} be `parsedURL`{.variable},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer④
    link-type="dfn"} and [isomorphic
    encoded](https://infra.spec.whatwg.org/#isomorphic-encode){#ref-for-isomorphic-encode①①
    link-type="dfn"}.

7.  [Append](#concept-header-list-append){#ref-for-concept-header-list-append②③
    link-type="dfn"} (\``Location`\`, `value`{.variable}) to
    `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response⑧
    link-type="dfn"}'s [header
    list](#concept-response-header-list){#ref-for-concept-response-header-list③⑦
    link-type="dfn"}.

8.  Return `responseObject`{.variable}.
:::

::: {.algorithm algorithm="json(data, init)" algorithm-for="Response"}
The static
[`json(``data`{.variable}`, ``init`{.variable}`)`]{#dom-response-json
.dfn .dfn-paneled .idl-code dfn-for="Response" dfn-type="method"
export="" lt="json(data, init)|json(data)"} method steps are:

1.  Let `bytes`{.variable} the result of running [serialize a JavaScript
    value to JSON
    bytes](https://infra.spec.whatwg.org/#serialize-a-javascript-value-to-json-bytes){#ref-for-serialize-a-javascript-value-to-json-bytes
    link-type="dfn"} on `data`{.variable}.

2.  Let `body`{.variable} be the result of
    [extracting](#concept-bodyinit-extract){#ref-for-concept-bodyinit-extract④
    link-type="dfn"} `bytes`{.variable}.

3.  Let `responseObject`{.variable} be the result of
    [creating](#response-create){#ref-for-response-create②
    link-type="dfn"} a [`Response`{.idl}](#response){#ref-for-response②②
    link-type="idl"} object, given a new
    [response](#concept-response){#ref-for-concept-response⑥⑤
    link-type="dfn"}, \"`response`\", and the [current
    realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm②
    link-type="dfn"}.

4.  Perform [initialize a
    response](#initialize-a-response){#ref-for-initialize-a-response①
    link-type="dfn"} given `responseObject`{.variable},
    `init`{.variable}, and (`body`{.variable}, \"`application/json`\").

5.  Return `responseObject`{.variable}.
:::

The [`type`]{#dom-response-type .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧②
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response⑨
link-type="dfn"}'s
[type](#concept-response-type){#ref-for-concept-response-type①②
link-type="dfn"}.

The [`url`]{#dom-response-url .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return the empty string if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧③
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①⓪
link-type="dfn"}'s
[URL](#concept-response-url){#ref-for-concept-response-url⑧
link-type="dfn"} is null; otherwise
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧④
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①①
link-type="dfn"}'s
[URL](#concept-response-url){#ref-for-concept-response-url⑨
link-type="dfn"},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer⑤
link-type="dfn"} with [*exclude
fragment*](https://url.spec.whatwg.org/#url-serializer-exclude-fragment){#ref-for-url-serializer-exclude-fragment①
link-type="dfn"} set to true.

The [`redirected`]{#dom-response-redirected .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑤
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①②
link-type="dfn"}'s [URL
list](#concept-response-url-list){#ref-for-concept-response-url-list①①
link-type="dfn"}'s
[size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size②
link-type="dfn"} is greater than 1; otherwise false.

To filter out [responses](#concept-response){#ref-for-concept-response⑥⑥
link-type="dfn"} that are the result of a redirect, do this directly
through the API, e.g., `fetch(url, { redirect:"error" })`. This way a
potentially unsafe
[response](#concept-response){#ref-for-concept-response⑥⑦
link-type="dfn"} cannot accidentally leak.

The [`status`]{#dom-response-status .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑥
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①③
link-type="dfn"}'s
[status](#concept-response-status){#ref-for-concept-response-status②④
link-type="dfn"}.

The [`ok`]{#dom-response-ok .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑦
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①④
link-type="dfn"}'s
[status](#concept-response-status){#ref-for-concept-response-status②⑤
link-type="dfn"} is an [ok status](#ok-status){#ref-for-ok-status③
link-type="dfn"}; otherwise false.

The [`statusText`]{#dom-response-statustext .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑧
link-type="dfn"}'s
[response](#concept-response-response){#ref-for-concept-response-response①⑤
link-type="dfn"}'s [status
message](#concept-response-status-message){#ref-for-concept-response-status-message⑧
link-type="dfn"}.

The [`headers`]{#dom-response-headers .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑨
link-type="dfn"}'s
[headers](#response-headers){#ref-for-response-headers③
link-type="dfn"}.

------------------------------------------------------------------------

::: {.algorithm algorithm="clone()" algorithm-for="Response"}
The [`clone()`]{#dom-response-clone .dfn .dfn-paneled .idl-code
dfn-for="Response" dfn-type="method" export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⓪
    link-type="dfn"} is
    [unusable](#body-unusable){#ref-for-body-unusable③ link-type="dfn"},
    then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⑤
    link-type="dfn"} a
    [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑧
    link-type="idl"}.

2.  Let `clonedResponse`{.variable} be the result of
    [cloning](#concept-response-clone){#ref-for-concept-response-clone①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨①
    link-type="dfn"}'s
    [response](#concept-response-response){#ref-for-concept-response-response①⑥
    link-type="dfn"}.

3.  Return the result of
    [creating](#response-create){#ref-for-response-create③
    link-type="dfn"} a [`Response`{.idl}](#response){#ref-for-response②③
    link-type="idl"} object, given `clonedResponse`{.variable},
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨②
    link-type="dfn"}'s
    [headers](#response-headers){#ref-for-response-headers④
    link-type="dfn"}'s
    [guard](#concept-headers-guard){#ref-for-concept-headers-guard①⑥
    link-type="dfn"}, and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨③
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑦
    link-type="dfn"}.
:::

### [5.6. ]{.secno}[Fetch method]{.content}[](#fetch-method){.self-link} {#fetch-method .heading .settled level="5.6"}

``` {.idl .highlight .def}
partial interface mixin WindowOrWorkerGlobalScope {
  [NewObject] Promise<Response> fetch(RequestInfo input, optional RequestInit init = {});
};
```

::: {.algorithm algorithm="fetch(input, init)" algorithm-for="WindowOrWorkerGlobalScope"}
The
[`fetch(``input`{.variable}`, ``init`{.variable}`)`]{#dom-global-fetch
.dfn .dfn-paneled .idl-code dfn-for="WindowOrWorkerGlobalScope"
dfn-type="method" export="" lt="fetch(input, init)|fetch(input)"} method
steps are:

1.  Let `p`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise②
    link-type="dfn"}.

2.  Let `requestObject`{.variable} be the result of invoking the initial
    value of [`Request`{.idl}](#request){#ref-for-request①⑥
    link-type="idl"} as constructor with `input`{.variable} and
    `init`{.variable} as arguments. If this throws an exception,
    [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject①
    link-type="dfn"} `p`{.variable} with it and return `p`{.variable}.

3.  Let `request`{.variable} be `requestObject`{.variable}'s
    [request](#concept-request-request){#ref-for-concept-request-request②⑦
    link-type="dfn"}.

4.  If `requestObject`{.variable}'s
    [signal](#request-signal){#ref-for-request-signal⑥ link-type="dfn"}
    is
    [aborted](https://dom.spec.whatwg.org/#abortsignal-aborted){#ref-for-abortsignal-aborted
    link-type="dfn"}, then:

    1.  [Abort the `fetch()` call](#abort-fetch){#ref-for-abort-fetch
        link-type="dfn"} with `p`{.variable}, `request`{.variable},
        null, and `requestObject`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal⑦
        link-type="dfn"}'s [abort
        reason](https://dom.spec.whatwg.org/#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason
        link-type="dfn"}.

    2.  Return `p`{.variable}.

5.  Let `globalObject`{.variable} be `request`{.variable}'s
    [client](#concept-request-client){#ref-for-concept-request-client③⓪
    link-type="dfn"}'s [global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-settings-object-global){#ref-for-concept-settings-object-global⑦
    link-type="dfn"}.

6.  If `globalObject`{.variable} is a
    [`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope
    link-type="idl"} object, then set `request`{.variable}'s
    [service-workers
    mode](#request-service-workers-mode){#ref-for-request-service-workers-mode④
    link-type="dfn"} to \"`none`\".

7.  Let `responseObject`{.variable} be null.

8.  Let `relevantRealm`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨④
    link-type="dfn"}'s [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm⑧
    link-type="dfn"}.

9.  Let `locallyAborted`{.variable} be false.

    This lets us reject promises with predictable timing, when the
    request to abort comes from the same thread as the call to fetch.

10. Let `controller`{.variable} be null.

11. [Add the following abort
    steps](https://dom.spec.whatwg.org/#abortsignal-add){#ref-for-abortsignal-add
    link-type="dfn"} to `requestObject`{.variable}'s
    [signal](#request-signal){#ref-for-request-signal⑧ link-type="dfn"}:

    1.  Set `locallyAborted`{.variable} to true.

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②⑦
        link-type="dfn"}: `controller`{.variable} is non-null.

    3.  [Abort](#fetch-controller-abort){#ref-for-fetch-controller-abort②
        link-type="dfn"} `controller`{.variable} with
        `requestObject`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal⑨
        link-type="dfn"}'s [abort
        reason](https://dom.spec.whatwg.org/#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①
        link-type="dfn"}.

    4.  [Abort the `fetch()` call](#abort-fetch){#ref-for-abort-fetch①
        link-type="dfn"} with `p`{.variable}, `request`{.variable},
        `responseObject`{.variable}, and `requestObject`{.variable}'s
        [signal](#request-signal){#ref-for-request-signal①⓪
        link-type="dfn"}'s [abort
        reason](https://dom.spec.whatwg.org/#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason②
        link-type="dfn"}.

12. Set `controller`{.variable} to the result of calling
    [fetch](#concept-fetch){#ref-for-concept-fetch②⑧ link-type="dfn"}
    given `request`{.variable} and
    [*processResponse*](#process-response){#ref-for-process-response①
    link-type="dfn"} given `response`{.variable} being these steps:

    1.  If `locallyAborted`{.variable} is true, then abort these steps.

    2.  If `response`{.variable}'s [aborted
        flag](#concept-response-aborted){#ref-for-concept-response-aborted②
        link-type="dfn"} is set, then:

        1.  Let `deserializedError`{.variable} be the result of
            [deserialize a serialized abort
            reason](#deserialize-a-serialized-abort-reason){#ref-for-deserialize-a-serialized-abort-reason①
            link-type="dfn"} given `controller`{.variable}'s [serialized
            abort
            reason](#fetch-controller-serialized-abort-reason){#ref-for-fetch-controller-serialized-abort-reason②
            link-type="dfn"} and `relevantRealm`{.variable}.

        2.  [Abort the `fetch()`
            call](#abort-fetch){#ref-for-abort-fetch② link-type="dfn"}
            with `p`{.variable}, `request`{.variable},
            `responseObject`{.variable}, and
            `deserializedError`{.variable}.

        3.  Abort these steps.

    3.  If `response`{.variable} is a [network
        error](#concept-network-error){#ref-for-concept-network-error⑥⓪
        link-type="dfn"}, then
        [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject②
        link-type="dfn"} `p`{.variable} with a
        [`TypeError`{.idl}](https://webidl.spec.whatwg.org/#exceptiondef-typeerror){#ref-for-exceptiondef-typeerror②⑨
        link-type="idl"} and abort these steps.

    4.  Set `responseObject`{.variable} to the result of
        [creating](#response-create){#ref-for-response-create④
        link-type="dfn"} a
        [`Response`{.idl}](#response){#ref-for-response②⑤
        link-type="idl"} object, given `response`{.variable},
        \"`immutable`\", and `relevantRealm`{.variable}.

    5.  [Resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve②
        link-type="dfn"} `p`{.variable} with
        `responseObject`{.variable}.

13. Return `p`{.variable}.
:::

::: {.algorithm algorithm="Abort the fetch() call"}
To [abort a `fetch()` call]{#abort-fetch .dfn .dfn-paneled
dfn-type="dfn" export="" lt="Abort the fetch() call"} with a
`promise`{.variable}, `request`{.variable}, `responseObject`{.variable},
and an `error`{.variable}:

1.  [Reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject③
    link-type="dfn"} `promise`{.variable} with `error`{.variable}.

    This is a no-op if `promise`{.variable} has already fulfilled.

2.  If `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑧
    link-type="dfn"} is non-null and is
    [readable](https://streams.spec.whatwg.org/#readablestream-readable){#ref-for-readablestream-readable③
    link-type="dfn"}, then
    [cancel](https://streams.spec.whatwg.org/#readablestream-cancel){#ref-for-readablestream-cancel①
    link-type="dfn"} `request`{.variable}'s
    [body](#concept-request-body){#ref-for-concept-request-body④⑨
    link-type="dfn"} with `error`{.variable}.

3.  If `responseObject`{.variable} is null, then return.

4.  Let `response`{.variable} be `responseObject`{.variable}'s
    [response](#concept-response-response){#ref-for-concept-response-response①⑦
    link-type="dfn"}.

5.  If `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body②⑨
    link-type="dfn"} is non-null and is
    [readable](https://streams.spec.whatwg.org/#readablestream-readable){#ref-for-readablestream-readable④
    link-type="dfn"}, then
    [error](https://streams.spec.whatwg.org/#readablestream-error){#ref-for-readablestream-error②
    link-type="dfn"} `response`{.variable}'s
    [body](#concept-response-body){#ref-for-concept-response-body③⓪
    link-type="dfn"} with `error`{.variable}.
:::

### [5.7. ]{.secno}[Garbage collection]{.content}[](#garbage-collection){.self-link} {#garbage-collection .heading .settled level="5.7"}

The user agent may
[terminate](#fetch-controller-terminate){#ref-for-fetch-controller-terminate⑤
link-type="dfn"} an ongoing fetch if that termination is not observable
through script.

\"Observable through script\" means observable through
[`fetch()`](#dom-global-fetch){#ref-for-dom-global-fetch⑦ .idl-code
link-type="method"}'s arguments and return value. Other ways, such as
communicating with the server through a side-channel are not included.

The server being able to observe garbage collection has precedent, e.g.,
with
[`WebSocket`{.idl}](https://websockets.spec.whatwg.org/#websocket){#ref-for-websocket①
link-type="idl"} and
[`XMLHttpRequest`{.idl}](https://xhr.spec.whatwg.org/#xmlhttprequest){#ref-for-xmlhttprequest⑥
link-type="idl"} objects.

::: {#terminate-examples .example}
[](#terminate-examples){.self-link}

The user agent can terminate the fetch because the termination cannot be
observed.

``` {.lang-javascript .highlight}
fetch("https://www.example.com/")
```

The user agent cannot terminate the fetch because the termination can be
observed through the promise.

``` {.lang-javascript .highlight}
window.promise = fetch("https://www.example.com/")
```

The user agent can terminate the fetch because the associated body is
not observable.

``` {.lang-javascript .highlight}
window.promise = fetch("https://www.example.com/").then(res => res.headers)
```

The user agent can terminate the fetch because the termination cannot be
observed.

``` {.lang-javascript .highlight}
fetch("https://www.example.com/").then(res => res.body.getReader().closed)
```

The user agent cannot terminate the fetch because one can observe the
termination by registering a handler for the promise object.

``` {.lang-javascript .highlight}
window.promise = fetch("https://www.example.com/")
  .then(res => res.body.getReader().closed)
```

The user agent cannot terminate the fetch as termination would be
observable via the registered handler.

``` {.lang-javascript .highlight}
fetch("https://www.example.com/")
  .then(res => {
    res.body.getReader().closed.then(() => console.log("stream closed!"))
  })
```

(The above examples of non-observability assume that built-in properties
and functions, such as
[`body.getReader()`{.idl}](https://streams.spec.whatwg.org/#rs-get-reader){#ref-for-rs-get-reader
link-type="idl"}, have not been overwritten.)
:::

