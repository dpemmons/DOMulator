## [7. ]{.secno}[Sets]{.content}[](#sets){.self-link} {#sets .heading .settled level="7"}

Yes, the name
[`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist④
link-type="idl"} is an unfortunate legacy mishap.

### [7.1. ]{.secno}[Interface [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist⑤ link-type="idl"}]{.content}[](#interface-domtokenlist){.self-link} {#interface-domtokenlist .heading .settled level="7.1"}

``` {.idl .highlight .def}
[Exposed=Window]
interface DOMTokenList {
  readonly attribute unsigned long length;
  getter DOMString? item(unsigned long index);
  boolean contains(DOMString token);
  [CEReactions] undefined add(DOMString... tokens);
  [CEReactions] undefined remove(DOMString... tokens);
  [CEReactions] boolean toggle(DOMString token, optional boolean force);
  [CEReactions] boolean replace(DOMString token, DOMString newToken);
  boolean supports(DOMString token);
  [CEReactions] stringifier attribute DOMString value;
  iterable<DOMString>;
};
```

A [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist⑥
link-type="idl"} object has an associated [token
set]{#concept-dtl-tokens .dfn .dfn-paneled dfn-for="DOMTokenList"
dfn-type="dfn" export=""} (a
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set⑦
link-type="dfn"}), which is initially empty.

A [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist⑦
link-type="idl"} object also has an associated
[element](#concept-element){#ref-for-concept-element①①⓪ link-type="dfn"}
and an [attribute](#concept-attribute){#ref-for-concept-attribute⑥⑦
link-type="dfn"}'s [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②⑧
link-type="dfn"}.

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications①⓪
link-type="dfn"} may define [supported tokens]{#concept-supported-tokens
.dfn .dfn-paneled dfn-for="Node" dfn-type="dfn" export=""} for a
[`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist⑧
link-type="idl"}'s associated
[attribute](#concept-attribute){#ref-for-concept-attribute⑥⑧
link-type="dfn"}'s [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②⑨
link-type="dfn"}.

A [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist⑨
link-type="idl"} object's [validation
steps]{#concept-domtokenlist-validation .dfn .dfn-paneled
dfn-for="DOMTokenList" dfn-type="dfn" export=""} for a given
`token`{.variable} are:

1.  If the associated
    [attribute](#concept-attribute){#ref-for-concept-attribute⑥⑨
    link-type="dfn"}'s [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③⓪
    link-type="dfn"} does not define [supported
    tokens](#concept-supported-tokens){#ref-for-concept-supported-tokens
    link-type="dfn"},
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨②
    link-type="dfn"} a `TypeError`.

2.  Let `lowercase token`{.variable} be a copy of `token`{.variable}, in
    [ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase⑧
    link-type="dfn"}.

3.  If `lowercase token`{.variable} is present in [supported
    tokens](#concept-supported-tokens){#ref-for-concept-supported-tokens①
    link-type="dfn"}, return true.

4.  Return false.

A [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist①⓪
link-type="idl"} object's [update steps]{#concept-dtl-update .dfn
.dfn-paneled dfn-for="DOMTokenList" dfn-type="dfn" noexport=""} are:

1.  If the associated
    [element](#concept-element){#ref-for-concept-element①①①
    link-type="dfn"} does not have an associated
    [attribute](#concept-attribute){#ref-for-concept-attribute⑦⓪
    link-type="dfn"} and [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①
    link-type="dfn"} is empty, then return.

2.  [Set an attribute
    value](#concept-element-attributes-set-value){#ref-for-concept-element-attributes-set-value②
    link-type="dfn"} for the associated
    [element](#concept-element){#ref-for-concept-element①①②
    link-type="dfn"} using associated
    [attribute](#concept-attribute){#ref-for-concept-attribute⑦①
    link-type="dfn"}'s [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③①
    link-type="dfn"} and the result of running the [ordered set
    serializer](#concept-ordered-set-serializer){#ref-for-concept-ordered-set-serializer
    link-type="dfn"} for [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens②
    link-type="dfn"}.

A [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist①①
link-type="idl"} object's [serialize steps]{#concept-dtl-serialize .dfn
.dfn-paneled dfn-for="DOMTokenList" dfn-type="dfn" noexport=""} are to
return the result of running [get an attribute
value](#concept-element-attributes-get-value){#ref-for-concept-element-attributes-get-value①
link-type="dfn"} given the associated
[element](#concept-element){#ref-for-concept-element①①③ link-type="dfn"}
and the associated
[attribute](#concept-attribute){#ref-for-concept-attribute⑦②
link-type="dfn"}'s [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③②
link-type="dfn"}.

------------------------------------------------------------------------

A [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist①②
link-type="idl"} object has these [attribute change
steps](#concept-element-attributes-change-ext){#ref-for-concept-element-attributes-change-ext④
link-type="dfn"} for its associated
[element](#concept-element){#ref-for-concept-element①①④
link-type="dfn"}:

1.  If `localName`{.variable} is associated attribute's [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③③
    link-type="dfn"}, `namespace`{.variable} is null, and
    `value`{.variable} is null, then
    [empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty⑥
    link-type="dfn"} [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens③
    link-type="dfn"}.

2.  Otherwise, if `localName`{.variable} is associated attribute's
    [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③④
    link-type="dfn"}, `namespace`{.variable} is null, then set [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens④
    link-type="dfn"} to `value`{.variable},
    [parsed](#concept-ordered-set-parser){#ref-for-concept-ordered-set-parser①
    link-type="dfn"}.

When a [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist①③
link-type="idl"} object is created, then:

1.  Let `element`{.variable} be associated
    [element](#concept-element){#ref-for-concept-element①①⑤
    link-type="dfn"}.

2.  Let `localName`{.variable} be associated attribute's [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③⑤
    link-type="dfn"}.

3.  Let `value`{.variable} be the result of [getting an attribute
    value](#concept-element-attributes-get-value){#ref-for-concept-element-attributes-get-value②
    link-type="dfn"} given `element`{.variable} and
    `localName`{.variable}.

4.  Run the [attribute change
    steps](#concept-element-attributes-change-ext){#ref-for-concept-element-attributes-change-ext⑤
    link-type="dfn"} for `element`{.variable}, `localName`{.variable},
    `value`{.variable}, `value`{.variable}, and null.

`tokenlist`{.variable}` . `[`length`{.idl}](#dom-domtokenlist-length){#ref-for-dom-domtokenlist-length① link-type="idl"}

:   Returns the number of tokens.

`tokenlist`{.variable}` . `[`item(index)`{.idl}](#dom-domtokenlist-item){#ref-for-dom-domtokenlist-item① link-type="idl"}\
`tokenlist`{.variable}`[``index`{.variable}`]`

:   Returns the token with index `index`{.variable}.

`tokenlist`{.variable}` . `[`contains(token)`{.idl}](#dom-domtokenlist-contains){#ref-for-dom-domtokenlist-contains① link-type="idl"}

:   Returns true if `token`{.variable} is present; otherwise false.

`tokenlist`{.variable}` . `[`add(``tokens`{.variable}`…)`](#dom-domtokenlist-add){#ref-for-dom-domtokenlist-add① link-type="functionish"}

:   Adds all arguments passed, except those already present.

    Throws a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪⑦
    link-type="idl"} if one of the arguments is the empty string.

    Throws an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪⑧
    link-type="idl"} if one of the arguments contains any [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace①
    link-type="dfn"}.

`tokenlist`{.variable}` . `[`remove(``tokens`{.variable}`…)`](#dom-domtokenlist-remove){#ref-for-dom-domtokenlist-remove① link-type="functionish"}

:   Removes arguments passed, if they are present.

    Throws a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪⑨
    link-type="idl"} if one of the arguments is the empty string.

    Throws an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①⓪
    link-type="idl"} if one of the arguments contains any [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace②
    link-type="dfn"}.

`tokenlist`{.variable}` . `[`toggle(``token`{.variable}` [, ``force`{.variable}`])`](#dom-domtokenlist-toggle){#ref-for-dom-domtokenlist-toggle① .idl-code link-type="method"}

:   If `force`{.variable} is not given, \"toggles\" `token`{.variable},
    removing it if it's present and adding it if it's not present. If
    `force`{.variable} is true, adds `token`{.variable} (same as
    [`add()`{.idl}](#dom-domtokenlist-add){#ref-for-dom-domtokenlist-add②
    link-type="idl"}). If `force`{.variable} is false, removes
    `token`{.variable} (same as
    [`remove()`{.idl}](#dom-domtokenlist-remove){#ref-for-dom-domtokenlist-remove②
    link-type="idl"}).

    Returns true if `token`{.variable} is now present; otherwise false.

    Throws a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①①
    link-type="idl"} if `token`{.variable} is empty.

    Throws an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①②
    link-type="idl"} if `token`{.variable} contains any spaces.

`tokenlist`{.variable}` . `[`replace(``token`{.variable}`, ``newToken`{.variable}`)`](#dom-domtokenlist-replace){#ref-for-dom-domtokenlist-replace① .idl-code link-type="method"}

:   Replaces `token`{.variable} with `newToken`{.variable}.

    Returns true if `token`{.variable} was replaced with
    `newToken`{.variable}; otherwise false.

    Throws a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror⑦
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①③
    link-type="idl"} if one of the arguments is the empty string.

    Throws an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①④
    link-type="idl"} if one of the arguments contains any [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace③
    link-type="dfn"}.

`tokenlist`{.variable}` . `[`supports(``token`{.variable}`)`](#dom-domtokenlist-supports){#ref-for-dom-domtokenlist-supports① .idl-code link-type="method"}

:   Returns true if `token`{.variable} is in the associated attribute's
    supported tokens. Returns false otherwise.

    Throws a `TypeError` if the associated attribute has no supported
    tokens defined.

`tokenlist`{.variable}` . `[`value`{.idl}](#dom-domtokenlist-value){#ref-for-dom-domtokenlist-value② link-type="idl"}

:   Returns the associated set as string.

    Can be set, to change the associated attribute.

The [`length`]{#dom-domtokenlist-length .dfn .dfn-paneled .idl-code
dfn-for="DOMTokenList" dfn-type="attribute" export=""} attribute\'
getter must return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①⑨
link-type="dfn"}'s [token
set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens⑤
link-type="dfn"}'s
[size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑦
link-type="dfn"}.

The object's [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices⑥
link-type="dfn"} are the numbers in the range zero to object's [token
set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens⑥
link-type="dfn"}'s
[size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑧
link-type="dfn"} minus one, unless [token
set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens⑦ link-type="dfn"}
[is
empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty①①
link-type="dfn"}, in which case there are no [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices⑦
link-type="dfn"}.

The [`item(``index`{.variable}`)`]{#dom-domtokenlist-item .dfn
.dfn-paneled .idl-code dfn-for="DOMTokenList" dfn-type="method"
export=""} method steps are:

1.  If `index`{.variable} is equal to or greater than
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②⓪
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens⑧
    link-type="dfn"}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑨
    link-type="dfn"}, then return null.

2.  Return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②①
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens⑨
    link-type="dfn"}\[`index`{.variable}\].

The [`contains(``token`{.variable}`)`]{#dom-domtokenlist-contains .dfn
.dfn-paneled .idl-code dfn-for="DOMTokenList" dfn-type="method"
export=""} method steps are to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②②
link-type="dfn"}'s [token
set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①⓪
link-type="dfn"}\[`token`{.variable}\]
[exists](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain⑦
link-type="dfn"}; otherwise false.

The [`add(``tokens`{.variable}`…)`]{#dom-domtokenlist-add .dfn
.dfn-paneled .idl-code dfn-for="DOMTokenList" dfn-type="method"
export="" lt="add(...tokens)|add()|add(tokens)"} method steps are:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②⑦
    link-type="dfn"} `token`{.variable} of `tokens`{.variable}:

    1.  If `token`{.variable} is the empty string, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨③
        link-type="dfn"} a
        \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror⑧
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①⑤
        link-type="idl"}.

    2.  If `token`{.variable} contains any [ASCII
        whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace④
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨④
        link-type="dfn"} an
        \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①⑦
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①⑥
        link-type="idl"}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②⑧
    link-type="dfn"} `token`{.variable} of `tokens`{.variable},
    [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append①②
    link-type="dfn"} `token`{.variable} to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②③
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①①
    link-type="dfn"}.

3.  Run the [update
    steps](#concept-dtl-update){#ref-for-concept-dtl-update
    link-type="dfn"}.

The [`remove(``tokens`{.variable}`…)`]{#dom-domtokenlist-remove .dfn
.dfn-paneled .idl-code dfn-for="DOMTokenList" dfn-type="method"
export="" lt="remove(...tokens)|remove()|remove(tokens)"} method steps
are:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②⑨
    link-type="dfn"} `token`{.variable} of `tokens`{.variable}:

    1.  If `token`{.variable} is the empty string, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨⑤
        link-type="dfn"} a
        \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror⑨
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①⑦
        link-type="idl"}.

    2.  If `token`{.variable} contains any [ASCII
        whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace⑤
        link-type="dfn"}, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨⑥
        link-type="dfn"} an
        \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①⑧
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①⑧
        link-type="idl"}.

2.  For each `token`{.variable} in `tokens`{.variable},
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove⑧
    link-type="dfn"} `token`{.variable} from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②④
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①②
    link-type="dfn"}.

3.  Run the [update
    steps](#concept-dtl-update){#ref-for-concept-dtl-update①
    link-type="dfn"}.

The
[`toggle(``token`{.variable}`, ``force`{.variable}`)`]{#dom-domtokenlist-toggle
.dfn .dfn-paneled .idl-code dfn-for="DOMTokenList" dfn-type="method"
export="" lt="toggle(token, force)|toggle(token)"} method steps are:

1.  If `token`{.variable} is the empty string, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨⑦
    link-type="dfn"} a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror①⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①⑨
    link-type="idl"}.

2.  If `token`{.variable} contains any [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace⑥
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨⑧
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①⑨
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①②⓪
    link-type="idl"}.

3.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②⑤
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①③
    link-type="dfn"}\[`token`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain⑧
    link-type="dfn"}:

    1.  If `force`{.variable} is either not given or is false, then
        [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove⑨
        link-type="dfn"} `token`{.variable} from
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②⑥
        link-type="dfn"}'s [token
        set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①④
        link-type="dfn"}, run the [update
        steps](#concept-dtl-update){#ref-for-concept-dtl-update②
        link-type="dfn"} and return false.

    2.  Return true.

4.  Otherwise, if `force`{.variable} not given or is true,
    [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append①③
    link-type="dfn"} `token`{.variable} to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②⑦
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①⑤
    link-type="dfn"}, run the [update
    steps](#concept-dtl-update){#ref-for-concept-dtl-update③
    link-type="dfn"}, and return true.

5.  Return false.

The [update steps](#concept-dtl-update){#ref-for-concept-dtl-update④
link-type="dfn"} are not always run for
[`toggle()`{.idl}](#dom-domtokenlist-toggle){#ref-for-dom-domtokenlist-toggle②
link-type="idl"} for web compatibility.

The
[`replace(``token`{.variable}`, ``newToken`{.variable}`)`]{#dom-domtokenlist-replace
.dfn .dfn-paneled .idl-code dfn-for="DOMTokenList" dfn-type="method"
export=""} method steps are:

1.  If either `token`{.variable} or `newToken`{.variable} is the empty
    string, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨⑨
    link-type="dfn"} a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror①①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①②①
    link-type="idl"}.

2.  If either `token`{.variable} or `newToken`{.variable} contains any
    [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace⑦
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⓪⓪
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror②⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①②②
    link-type="idl"}.

3.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②⑧
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①⑥
    link-type="dfn"} does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain⑨
    link-type="dfn"} `token`{.variable}, then return false.

4.  [Replace](https://infra.spec.whatwg.org/#set-replace){#ref-for-set-replace
    link-type="dfn"} `token`{.variable} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②⑨
    link-type="dfn"}'s [token
    set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens①⑦
    link-type="dfn"} with `newToken`{.variable}.

5.  Run the [update
    steps](#concept-dtl-update){#ref-for-concept-dtl-update⑤
    link-type="dfn"}.

6.  Return true.

The [update steps](#concept-dtl-update){#ref-for-concept-dtl-update⑥
link-type="dfn"} are not always run for
[`replace()`{.idl}](#dom-domtokenlist-replace){#ref-for-dom-domtokenlist-replace②
link-type="idl"} for web compatibility.

The [`supports(``token`{.variable}`)`]{#dom-domtokenlist-supports .dfn
.dfn-paneled .idl-code dfn-for="DOMTokenList" dfn-type="method"
export=""} method steps are:

1.  Let `result`{.variable} be the return value of [validation
    steps](#concept-domtokenlist-validation){#ref-for-concept-domtokenlist-validation
    link-type="dfn"} called with `token`{.variable}.

2.  Return `result`{.variable}.

The [`value`]{#dom-domtokenlist-value .dfn .dfn-paneled .idl-code
dfn-for="DOMTokenList" dfn-type="attribute" export=""} attribute must
return the result of running
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③③⓪
link-type="dfn"}'s [serialize
steps](#concept-dtl-serialize){#ref-for-concept-dtl-serialize
link-type="dfn"}.

Setting the
[`value`{.idl}](#dom-domtokenlist-value){#ref-for-dom-domtokenlist-value③
link-type="idl"} attribute must [set an attribute
value](#concept-element-attributes-set-value){#ref-for-concept-element-attributes-set-value③
link-type="dfn"} for the associated
[element](#concept-element){#ref-for-concept-element①①⑥ link-type="dfn"}
using associated
[attribute](#concept-attribute){#ref-for-concept-attribute⑦③
link-type="dfn"}'s [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③⑥
link-type="dfn"} and the given value.

