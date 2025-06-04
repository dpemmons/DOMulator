## [1. ]{.secno}[]{#dependencies .bs-old-id}[]{#terminology .bs-old-id}[Infrastructure]{.content}[](#infrastructure){.self-link} {#infrastructure .heading .settled level="1"}

This specification depends on the Infra Standard.
[\[INFRA\]](#biblio-infra "Infra Standard"){link-type="biblio"}

Some of the terms used in this specification are defined in Encoding,
Selectors, Web IDL, XML, and Namespaces in XML.
[\[ENCODING\]](#biblio-encoding "Encoding Standard"){link-type="biblio"}
[\[SELECTORS4\]](#biblio-selectors4 "Selectors Level 4"){link-type="biblio"}
[\[WEBIDL\]](#biblio-webidl "Web IDL Standard"){link-type="biblio"}
[\[XML\]](#biblio-xml "Extensible Markup Language (XML) 1.0 (Fifth Edition)"){link-type="biblio"}
[\[XML-NAMES\]](#biblio-xml-names "Namespaces in XML 1.0 (Third Edition)"){link-type="biblio"}

When extensions are needed, the DOM Standard can be updated accordingly,
or a new standard can be written that hooks into the provided
extensibility hooks for [applicable
specifications]{#other-applicable-specifications .dfn .dfn-paneled
dfn-type="dfn" export="" lt="other applicable specifications"}.

### [1.1. ]{.secno}[Trees]{.content}[](#trees){.self-link} {#trees .heading .settled level="1.1"}

A [tree]{#concept-tree .dfn .dfn-paneled dfn-type="dfn" export=""} is a
finite hierarchical tree structure. In [tree order]{#concept-tree-order
.dfn .dfn-paneled dfn-type="dfn" export=""} is preorder, depth-first
traversal of a [tree](#concept-tree){#ref-for-concept-tree
link-type="dfn"}.

An object that [participates]{#concept-tree-participate .dfn
.dfn-paneled dfn-for="tree" dfn-type="dfn" export=""
lt="participate|participate in a tree|participates in a tree"} in a
[tree](#concept-tree){#ref-for-concept-tree① link-type="dfn"} has a
[parent]{#concept-tree-parent .dfn .dfn-paneled dfn-for="tree"
dfn-type="dfn" export=""}, which is either null or an object, and has
[children]{#concept-tree-child .dfn .dfn-paneled dfn-for="tree"
dfn-type="dfn" export="" lt="child|children"}, which is an [ordered
set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set
link-type="dfn"} of objects. An object `A`{.variable} whose
[parent](#concept-tree-parent){#ref-for-concept-tree-parent
link-type="dfn"} is object `B`{.variable} is a
[child](#concept-tree-child){#ref-for-concept-tree-child
link-type="dfn"} of `B`{.variable}.

The [root]{#concept-tree-root .dfn .dfn-paneled dfn-for="tree"
dfn-type="dfn" export=""} of an object is itself, if its
[parent](#concept-tree-parent){#ref-for-concept-tree-parent①
link-type="dfn"} is null, or else it is the
[root](#concept-tree-root){#ref-for-concept-tree-root link-type="dfn"}
of its [parent](#concept-tree-parent){#ref-for-concept-tree-parent②
link-type="dfn"}. The
[root](#concept-tree-root){#ref-for-concept-tree-root① link-type="dfn"}
of a [tree](#concept-tree){#ref-for-concept-tree② link-type="dfn"} is
any object
[participating](#concept-tree-participate){#ref-for-concept-tree-participate
link-type="dfn"} in that [tree](#concept-tree){#ref-for-concept-tree③
link-type="dfn"} whose
[parent](#concept-tree-parent){#ref-for-concept-tree-parent③
link-type="dfn"} is null.

An object `A`{.variable} is called a
[descendant]{#concept-tree-descendant .dfn .dfn-paneled dfn-for="tree"
dfn-type="dfn" export=""} of an object `B`{.variable}, if either
`A`{.variable} is a
[child](#concept-tree-child){#ref-for-concept-tree-child①
link-type="dfn"} of `B`{.variable} or `A`{.variable} is a
[child](#concept-tree-child){#ref-for-concept-tree-child②
link-type="dfn"} of an object `C`{.variable} that is a
[descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant
link-type="dfn"} of `B`{.variable}.

An [inclusive descendant]{#concept-tree-inclusive-descendant .dfn
.dfn-paneled dfn-for="tree" dfn-type="dfn" export=""} is an object or
one of its
[descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant①
link-type="dfn"}.

An object `A`{.variable} is called an [ancestor]{#concept-tree-ancestor
.dfn .dfn-paneled dfn-for="tree" dfn-type="dfn" export=""} of an object
`B`{.variable} if and only if `B`{.variable} is a
[descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②
link-type="dfn"} of `A`{.variable}.

An [inclusive ancestor]{#concept-tree-inclusive-ancestor .dfn
.dfn-paneled dfn-for="tree" dfn-type="dfn" export=""} is an object or
one of its
[ancestors](#concept-tree-ancestor){#ref-for-concept-tree-ancestor
link-type="dfn"}.

An object `A`{.variable} is called a [sibling]{#concept-tree-sibling
.dfn .dfn-paneled dfn-for="tree" dfn-type="dfn" export=""} of an object
`B`{.variable}, if and only if `B`{.variable} and `A`{.variable} share
the same non-null
[parent](#concept-tree-parent){#ref-for-concept-tree-parent④
link-type="dfn"}.

An [inclusive sibling]{#concept-tree-inclusive-sibling .dfn .dfn-paneled
dfn-for="tree" dfn-type="dfn" export=""} is an object or one of its
[siblings](#concept-tree-sibling){#ref-for-concept-tree-sibling
link-type="dfn"}.

An object `A`{.variable} is [preceding]{#concept-tree-preceding .dfn
.dfn-paneled dfn-for="tree" dfn-type="dfn" export=""} an object
`B`{.variable} if `A`{.variable} and `B`{.variable} are in the same
[tree](#concept-tree){#ref-for-concept-tree④ link-type="dfn"} and
`A`{.variable} comes before `B`{.variable} in [tree
order](#concept-tree-order){#ref-for-concept-tree-order
link-type="dfn"}.

An object `A`{.variable} is [following]{#concept-tree-following .dfn
.dfn-paneled dfn-for="tree" dfn-type="dfn" export=""} an object
`B`{.variable} if `A`{.variable} and `B`{.variable} are in the same
[tree](#concept-tree){#ref-for-concept-tree⑤ link-type="dfn"} and
`A`{.variable} comes after `B`{.variable} in [tree
order](#concept-tree-order){#ref-for-concept-tree-order①
link-type="dfn"}.

The [first child]{#concept-tree-first-child .dfn .dfn-paneled
dfn-for="tree" dfn-type="dfn" export=""} of an object is its first
[child](#concept-tree-child){#ref-for-concept-tree-child③
link-type="dfn"} or null if it has no
[children](#concept-tree-child){#ref-for-concept-tree-child④
link-type="dfn"}.

The [last child]{#concept-tree-last-child .dfn .dfn-paneled
dfn-for="tree" dfn-type="dfn" export=""} of an object is its last
[child](#concept-tree-child){#ref-for-concept-tree-child⑤
link-type="dfn"} or null if it has no
[children](#concept-tree-child){#ref-for-concept-tree-child⑥
link-type="dfn"}.

The [previous sibling]{#concept-tree-previous-sibling .dfn .dfn-paneled
dfn-for="tree" dfn-type="dfn" export=""} of an object is its first
[preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding
link-type="dfn"}
[sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling①
link-type="dfn"} or null if it has no
[preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding①
link-type="dfn"}
[sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling②
link-type="dfn"}.

The [next sibling]{#concept-tree-next-sibling .dfn .dfn-paneled
dfn-for="tree" dfn-type="dfn" export=""} of an object is its first
[following](#concept-tree-following){#ref-for-concept-tree-following
link-type="dfn"}
[sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling③
link-type="dfn"} or null if it has no
[following](#concept-tree-following){#ref-for-concept-tree-following①
link-type="dfn"}
[sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling④
link-type="dfn"}.

The [index]{#concept-tree-index .dfn .dfn-paneled dfn-for="tree"
dfn-type="dfn" export=""} of an object is its number of
[preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding②
link-type="dfn"}
[siblings](#concept-tree-sibling){#ref-for-concept-tree-sibling⑤
link-type="dfn"}, or 0 if it has none.

### [1.2. ]{.secno}[Ordered sets]{.content}[](#ordered-sets){.self-link} {#ordered-sets .heading .settled level="1.2"}

The [ordered set parser]{#concept-ordered-set-parser .dfn .dfn-paneled
dfn-type="dfn" export=""} takes a string `input`{.variable} and then
runs these steps:

1.  Let `inputTokens`{.variable} be the result of [splitting
    `input`{.variable} on ASCII
    whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#ref-for-split-on-ascii-whitespace
    link-type="dfn"}.

2.  Let `tokens`{.variable} be a new [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set①
    link-type="dfn"}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate
    link-type="dfn"} `token`{.variable} of `inputTokens`{.variable},
    [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append
    link-type="dfn"} `token`{.variable} to `tokens`{.variable}.

4.  Return `tokens`{.variable}.

The [ordered set serializer]{#concept-ordered-set-serializer .dfn
.dfn-paneled dfn-type="dfn" export=""} takes a `set`{.variable} and
returns the
[concatenation](https://infra.spec.whatwg.org/#string-concatenate){#ref-for-string-concatenate
link-type="dfn"} of `set`{.variable} using U+0020 SPACE.

### [1.3. ]{.secno}[Selectors]{.content}[](#selectors){.self-link} {#selectors .heading .settled level="1.3"}

To [scope-match a selectors string]{#scope-match-a-selectors-string .dfn
.dfn-paneled dfn-type="dfn" export=""} `selectors`{.variable} against a
`node`{.variable}, run these steps:

1.  Let `s`{.variable} be the result of [parse a
    selector](https://drafts.csswg.org/selectors-4/#parse-a-selector){#ref-for-parse-a-selector
    link-type="dfn"} `selectors`{.variable}.
    [\[SELECTORS4\]](#biblio-selectors4 "Selectors Level 4"){link-type="biblio"}

2.  If `s`{.variable} is failure, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw
    link-type="dfn"} a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException
    link-type="idl"}.

3.  Return the result of [match a selector against a
    tree](https://drafts.csswg.org/selectors-4/#match-a-selector-against-a-tree){#ref-for-match-a-selector-against-a-tree
    link-type="dfn"} with `s`{.variable} and `node`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root②
    link-type="dfn"} using [scoping
    root](https://drafts.csswg.org/selectors-4/#scoping-root){#ref-for-scoping-root
    link-type="dfn"} `node`{.variable}.
    [\[SELECTORS4\]](#biblio-selectors4 "Selectors Level 4"){link-type="biblio"}.

Support for namespaces within selectors is not planned and will not be
added.

### [1.4. ]{.secno}[Namespaces]{.content}[](#namespaces){.self-link} {#namespaces .heading .settled level="1.4"}

To [validate]{#validate .dfn .dfn-paneled dfn-type="dfn" export=""} a
`qualifiedName`{.variable},
[throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①
link-type="dfn"} an
\"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror
.idl-code link-type="exception"}\"
[`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①
link-type="idl"} if `qualifiedName`{.variable} does not match the
[`QName`](https://www.w3.org/TR/xml-names/#NT-QName){#ref-for-NT-QName
.css link-type="type"} production.

To [validate and extract]{#validate-and-extract .dfn .dfn-paneled
dfn-type="dfn" export=""} a `namespace`{.variable} and
`qualifiedName`{.variable}, run these steps:

1.  If `namespace`{.variable} is the empty string, then set it to null.

2.  [Validate](#validate){#ref-for-validate link-type="dfn"}
    `qualifiedName`{.variable}.

3.  Let `prefix`{.variable} be null.

4.  Let `localName`{.variable} be `qualifiedName`{.variable}.

5.  If `qualifiedName`{.variable} contains a U+003A (:):

    1.  Let `splitResult`{.variable} be the result of running [strictly
        split](https://infra.spec.whatwg.org/#strictly-split){#ref-for-strictly-split
        link-type="dfn"} given `qualifiedName`{.variable} and U+003A
        (:).

    2.  Set `prefix`{.variable} to `splitResult`{.variable}\[0\].

    3.  Set `localName`{.variable} to `splitResult`{.variable}\[1\].

6.  If `prefix`{.variable} is non-null and `namespace`{.variable} is
    null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②
    link-type="dfn"} a
    \"[`NamespaceError`{.idl}](https://webidl.spec.whatwg.org/#namespaceerror){#ref-for-namespaceerror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②
    link-type="idl"}.

7.  If `prefix`{.variable} is \"`xml`\" and `namespace`{.variable} is
    not the [XML
    namespace](https://infra.spec.whatwg.org/#xml-namespace){#ref-for-xml-namespace
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③
    link-type="dfn"} a
    \"[`NamespaceError`{.idl}](https://webidl.spec.whatwg.org/#namespaceerror){#ref-for-namespaceerror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③
    link-type="idl"}.

8.  If either `qualifiedName`{.variable} or `prefix`{.variable} is
    \"`xmlns`\" and `namespace`{.variable} is not the [XMLNS
    namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#ref-for-xmlns-namespace
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④
    link-type="dfn"} a
    \"[`NamespaceError`{.idl}](https://webidl.spec.whatwg.org/#namespaceerror){#ref-for-namespaceerror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④
    link-type="idl"}.

9.  If `namespace`{.variable} is the [XMLNS
    namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#ref-for-xmlns-namespace①
    link-type="dfn"} and neither `qualifiedName`{.variable} nor
    `prefix`{.variable} is \"`xmlns`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤
    link-type="dfn"} a
    \"[`NamespaceError`{.idl}](https://webidl.spec.whatwg.org/#namespaceerror){#ref-for-namespaceerror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤
    link-type="idl"}.

10. Return `namespace`{.variable}, `prefix`{.variable}, and
    `localName`{.variable}.

