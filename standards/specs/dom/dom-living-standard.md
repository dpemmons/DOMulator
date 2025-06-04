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

## [2. ]{.secno}[Events]{.content}[](#events){.self-link} {#events .heading .settled level="2"}

### [2.1. ]{.secno}[Introduction to \"DOM Events\"]{.content}[](#introduction-to-dom-events){.self-link} {#introduction-to-dom-events .heading .settled level="2.1"}

Throughout the web platform
[events](#concept-event){#ref-for-concept-event link-type="dfn"} are
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch
link-type="dfn"} to objects to signal an occurrence, such as network
activity or user interaction. These objects implement the
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget
link-type="idl"} interface and can therefore add [event
listeners](#concept-event-listener){#ref-for-concept-event-listener
link-type="dfn"} to observe
[events](#concept-event){#ref-for-concept-event① link-type="dfn"} by
calling
[`addEventListener()`{.idl}](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener
link-type="idl"}:

``` {.lang-javascript .highlight}
obj.addEventListener("load", imgFetched)

function imgFetched(ev) {
  // great success
  …
}
```

[Event
listeners](#concept-event-listener){#ref-for-concept-event-listener①
link-type="dfn"} can be removed by utilizing the
[`removeEventListener()`{.idl}](#dom-eventtarget-removeeventlistener){#ref-for-dom-eventtarget-removeeventlistener
link-type="idl"} method, passing the same arguments.

Alternatively, [event
listeners](#concept-event-listener){#ref-for-concept-event-listener②
link-type="dfn"} can be removed by passing an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal
link-type="idl"} to
[`addEventListener()`{.idl}](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener①
link-type="idl"} and calling
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort
link-type="idl"} on the controller owning the signal.

[Events](#concept-event){#ref-for-concept-event② link-type="dfn"} are
objects too and implement the [`Event`{.idl}](#event){#ref-for-event
link-type="idl"} interface (or a derived interface). In the example
above `ev`{.variable} is the
[event](#concept-event){#ref-for-concept-event③ link-type="dfn"}.
`ev`{.variable} is passed as an argument to the [event
listener](#concept-event-listener){#ref-for-concept-event-listener③
link-type="dfn"}'s
[callback](#event-listener-callback){#ref-for-event-listener-callback
link-type="dfn"} (typically a JavaScript Function as shown above).
[Event
listeners](#concept-event-listener){#ref-for-concept-event-listener④
link-type="dfn"} key off the
[event](#concept-event){#ref-for-concept-event④ link-type="dfn"}'s
[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type link-type="idl"}
attribute value (\"`load`\" in the above example). The
[event](#concept-event){#ref-for-concept-event⑤ link-type="dfn"}'s
[`target`{.idl}](#dom-event-target){#ref-for-dom-event-target
link-type="idl"} attribute value returns the object to which the
[event](#concept-event){#ref-for-concept-event⑥ link-type="dfn"} was
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①
link-type="dfn"} (`obj`{.variable} above).

Although [events](#concept-event){#ref-for-concept-event⑦
link-type="dfn"} are typically
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch②
link-type="dfn"} by the user agent as the result of user interaction or
the completion of some task, applications can
[dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch③
link-type="dfn"} [events](#concept-event){#ref-for-concept-event⑧
link-type="dfn"} themselves by using what are commonly known as
synthetic events:

``` {.lang-javascript .highlight}
// add an appropriate event listener
obj.addEventListener("cat", function(e) { process(e.detail) })

// create and dispatch the event
var event = new CustomEvent("cat", {"detail":{"hazcheeseburger":true}})
obj.dispatchEvent(event)
```

Apart from signaling, [events](#concept-event){#ref-for-concept-event⑨
link-type="dfn"} are sometimes also used to let an application control
what happens next in an operation. For instance as part of form
submission an [event](#concept-event){#ref-for-concept-event①⓪
link-type="dfn"} whose
[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①
link-type="idl"} attribute value is \"`submit`\" is
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch④
link-type="dfn"}. If this
[event](#concept-event){#ref-for-concept-event①① link-type="dfn"}'s
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault
link-type="idl"} method is invoked, form submission will be terminated.
Applications who wish to make use of this functionality through
[events](#concept-event){#ref-for-concept-event①② link-type="dfn"}
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑤
link-type="dfn"} by the application (synthetic events) can make use of
the return value of the
[`dispatchEvent()`{.idl}](#dom-eventtarget-dispatchevent){#ref-for-dom-eventtarget-dispatchevent
link-type="idl"} method:

``` {.lang-javascript .highlight}
if(obj.dispatchEvent(event)) {
  // event was not canceled, time for some magic
  …
}
```

When an [event](#concept-event){#ref-for-concept-event①③
link-type="dfn"} is
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑥
link-type="dfn"} to an object that
[participates](#concept-tree-participate){#ref-for-concept-tree-participate①
link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree⑥
link-type="dfn"} (e.g., an
[element](#concept-element){#ref-for-concept-element link-type="dfn"}),
it can reach [event
listeners](#concept-event-listener){#ref-for-concept-event-listener⑤
link-type="dfn"} on that object's
[ancestors](#concept-tree-ancestor){#ref-for-concept-tree-ancestor①
link-type="dfn"} too. Effectively, all the object's [inclusive
ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor
link-type="dfn"} [event
listeners](#concept-event-listener){#ref-for-concept-event-listener⑥
link-type="dfn"} whose
[capture](#event-listener-capture){#ref-for-event-listener-capture
link-type="dfn"} is true are invoked, in [tree
order](#concept-tree-order){#ref-for-concept-tree-order②
link-type="dfn"}. And then, if
[event](#concept-event){#ref-for-concept-event①④ link-type="dfn"}'s
[`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles
link-type="idl"} is true, all the object's [inclusive
ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①
link-type="dfn"} [event
listeners](#concept-event-listener){#ref-for-concept-event-listener⑦
link-type="dfn"} whose
[capture](#event-listener-capture){#ref-for-event-listener-capture①
link-type="dfn"} is false are invoked, now in reverse [tree
order](#concept-tree-order){#ref-for-concept-tree-order③
link-type="dfn"}.

Let's look at an example of how
[events](#concept-event){#ref-for-concept-event①⑤ link-type="dfn"} work
in a [tree](#concept-tree){#ref-for-concept-tree⑦ link-type="dfn"}:

``` {.lang-markup .highlight}
<!doctype html>
<html>
 <head>
  <title>Boring example</title>
 </head>
 <body>
  <p>Hello <span id=x>world</span>!</p>
  <script>
   function test(e) {
     debug(e.target, e.currentTarget, e.eventPhase)
   }
   document.addEventListener("hey", test, {capture: true})
   document.body.addEventListener("hey", test)
   var ev = new Event("hey", {bubbles:true})
   document.getElementById("x").dispatchEvent(ev)
  </script>
 </body>
</html>
```

The `debug` function will be invoked twice. Each time the
[event](#concept-event){#ref-for-concept-event①⑥ link-type="dfn"}'s
[`target`{.idl}](#dom-event-target){#ref-for-dom-event-target①
link-type="idl"} attribute value will be the `span`
[element](#concept-element){#ref-for-concept-element① link-type="dfn"}.
The first time
[`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget
link-type="idl"} attribute's value will be the
[document](#concept-document){#ref-for-concept-document
link-type="dfn"}, the second time the `body`
[element](#concept-element){#ref-for-concept-element② link-type="dfn"}.
[`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase
link-type="idl"} attribute's value switches from
[`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase
link-type="idl"} to
[`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase
link-type="idl"}. If an [event
listener](#concept-event-listener){#ref-for-concept-event-listener⑧
link-type="dfn"} was registered for the `span`
[element](#concept-element){#ref-for-concept-element③ link-type="dfn"},
[`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase①
link-type="idl"} attribute's value would have been
[`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target
link-type="idl"}.

### [2.2. ]{.secno}[Interface [`Event`{.idl}](#event){#ref-for-event① link-type="idl"}]{.content}[](#interface-event){.self-link} {#interface-event .heading .settled level="2.2"}

``` {.idl .highlight .def}
[Exposed=*]
interface Event {
  constructor(DOMString type, optional EventInit eventInitDict = {});

  readonly attribute DOMString type;
  readonly attribute EventTarget? target;
  readonly attribute EventTarget? srcElement; // legacy
  readonly attribute EventTarget? currentTarget;
  sequence<EventTarget> composedPath();

  const unsigned short NONE = 0;
  const unsigned short CAPTURING_PHASE = 1;
  const unsigned short AT_TARGET = 2;
  const unsigned short BUBBLING_PHASE = 3;
  readonly attribute unsigned short eventPhase;

  undefined stopPropagation();
           attribute boolean cancelBubble; // legacy alias of .stopPropagation()
  undefined stopImmediatePropagation();

  readonly attribute boolean bubbles;
  readonly attribute boolean cancelable;
           attribute boolean returnValue;  // legacy
  undefined preventDefault();
  readonly attribute boolean defaultPrevented;
  readonly attribute boolean composed;

  [LegacyUnforgeable] readonly attribute boolean isTrusted;
  readonly attribute DOMHighResTimeStamp timeStamp;

  undefined initEvent(DOMString type, optional boolean bubbles = false, optional boolean cancelable = false); // legacy
};

dictionary EventInit {
  boolean bubbles = false;
  boolean cancelable = false;
  boolean composed = false;
};
```

An [`Event`{.idl}](#event){#ref-for-event② link-type="idl"} object is
simply named an [event]{#concept-event .dfn .dfn-paneled dfn-type="dfn"
export=""}. It allows for signaling that something has occurred, e.g.,
that an image has completed downloading.

A [potential event target]{#potential-event-target .dfn .dfn-paneled
dfn-type="dfn" export=""} is null or an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑤
link-type="idl"} object.

An [event](#concept-event){#ref-for-concept-event①⑦ link-type="dfn"} has
an associated [target]{#event-target .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""} (a [potential event
target](#potential-event-target){#ref-for-potential-event-target
link-type="dfn"}). Unless stated otherwise it is null.

An [event](#concept-event){#ref-for-concept-event①⑧ link-type="dfn"} has
an associated [relatedTarget]{#event-relatedtarget .dfn .dfn-paneled
dfn-for="Event" dfn-type="dfn" export=""} (a [potential event
target](#potential-event-target){#ref-for-potential-event-target①
link-type="dfn"}). Unless stated otherwise it is null.

Other specifications use
[relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget
link-type="dfn"} to define a `relatedTarget` attribute.
[\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

An [event](#concept-event){#ref-for-concept-event①⑨ link-type="dfn"} has
an associated [touch target list]{#event-touch-target-list .dfn
.dfn-paneled dfn-for="Event" dfn-type="dfn" export=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list
link-type="dfn"} of zero or more [potential event
targets](#potential-event-target){#ref-for-potential-event-target②
link-type="dfn"}). Unless stated otherwise it is the empty list.

The [touch target
list](#event-touch-target-list){#ref-for-event-touch-target-list
link-type="dfn"} is for the exclusive use of defining the
[`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent
link-type="idl"} interface and related interfaces.
[\[TOUCH-EVENTS\]](#biblio-touch-events "Touch Events"){link-type="biblio"}

An [event](#concept-event){#ref-for-concept-event②⓪ link-type="dfn"} has
an associated [path]{#event-path .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""}. A [path](#event-path){#ref-for-event-path
link-type="dfn"} is a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①
link-type="dfn"} of
[structs](https://infra.spec.whatwg.org/#struct){#ref-for-struct
link-type="dfn"}. Each
[struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct①
link-type="dfn"} consists of an [invocation
target]{#event-path-invocation-target .dfn .dfn-paneled
dfn-for="Event/path" dfn-type="dfn" noexport=""} (an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑥
link-type="idl"} object), an
[invocation-target-in-shadow-tree]{#event-path-invocation-target-in-shadow-tree
.dfn .dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
boolean), a [shadow-adjusted target]{#event-path-shadow-adjusted-target
.dfn .dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
[potential event
target](#potential-event-target){#ref-for-potential-event-target③
link-type="dfn"}), a [relatedTarget]{#event-path-relatedtarget .dfn
.dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
[potential event
target](#potential-event-target){#ref-for-potential-event-target④
link-type="dfn"}), a [touch target list]{#event-path-touch-target-list
.dfn .dfn-paneled dfn-for="Event/path" dfn-type="dfn" noexport=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list②
link-type="dfn"} of [potential event
targets](#potential-event-target){#ref-for-potential-event-target⑤
link-type="dfn"}), a
[root-of-closed-tree]{#event-path-root-of-closed-tree .dfn .dfn-paneled
dfn-for="Event/path" dfn-type="dfn" noexport=""} (a boolean), and a
[slot-in-closed-tree]{#event-path-slot-in-closed-tree .dfn .dfn-paneled
dfn-for="Event/path" dfn-type="dfn" noexport=""} (a boolean). A
[path](#event-path){#ref-for-event-path① link-type="dfn"} is initially
the empty list.

`event`{.variable}` = new `[`Event`](#dom-event-event){#ref-for-dom-event-event .idl-code link-type="constructor"}`(``type`{.variable}` [, ``eventInitDict`{.variable}`])`
:   Returns a new `event`{.variable} whose
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type③
    link-type="idl"} attribute value is set to `type`{.variable}. The
    `eventInitDict`{.variable} argument allows for setting the
    [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles②
    link-type="idl"} and
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable①
    link-type="idl"} attributes via object members of the same name.

`event`{.variable}` . `[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type④ link-type="idl"}
:   Returns the type of `event`{.variable}, e.g. \"`click`\",
    \"`hashchange`\", or \"`submit`\".

`event`{.variable}` . `[`target`{.idl}](#dom-event-target){#ref-for-dom-event-target③ link-type="idl"}
:   Returns the object to which `event`{.variable} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑦
    link-type="dfn"} (its [target](#event-target){#ref-for-event-target
    link-type="dfn"}).

`event`{.variable}` . `[`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget② link-type="idl"}
:   Returns the object whose [event
    listener](#concept-event-listener){#ref-for-concept-event-listener⑨
    link-type="dfn"}'s
    [callback](#event-listener-callback){#ref-for-event-listener-callback①
    link-type="dfn"} is currently being invoked.

`event`{.variable}` . `[`composedPath()`{.idl}](#dom-event-composedpath){#ref-for-dom-event-composedpath① link-type="idl"}
:   Returns the [invocation
    target](#event-path-invocation-target){#ref-for-event-path-invocation-target
    link-type="dfn"} objects of `event`{.variable}'s
    [path](#event-path){#ref-for-event-path② link-type="dfn"} (objects
    on which listeners will be invoked), except for any
    [nodes](#concept-node){#ref-for-concept-node link-type="dfn"} in
    [shadow trees](#concept-shadow-tree){#ref-for-concept-shadow-tree
    link-type="dfn"} of which the [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root
    link-type="dfn"}'s [mode](#shadowroot-mode){#ref-for-shadowroot-mode
    link-type="dfn"} is \"`closed`\" that are not reachable from
    `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget③
    link-type="idl"}.

`event`{.variable}` . `[`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase③ link-type="idl"}
:   Returns the [event](#concept-event){#ref-for-concept-event②①
    link-type="dfn"}'s phase, which is one of
    [`NONE`{.idl}](#dom-event-none){#ref-for-dom-event-none①
    link-type="idl"},
    [`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase②
    link-type="idl"},
    [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target②
    link-type="idl"}, and
    [`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase②
    link-type="idl"}.

`event`{.variable}` . `[`stopPropagation`](#dom-event-stoppropagation){#ref-for-dom-event-stoppropagation① .idl-code link-type="method"}`()`
:   When
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑧
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree⑧
    link-type="dfn"}, invoking this method prevents `event`{.variable}
    from reaching any objects other than the current object.

`event`{.variable}` . `[`stopImmediatePropagation`](#dom-event-stopimmediatepropagation){#ref-for-dom-event-stopimmediatepropagation① .idl-code link-type="method"}`()`
:   Invoking this method prevents `event`{.variable} from reaching any
    registered [event
    listeners](#concept-event-listener){#ref-for-concept-event-listener①⓪
    link-type="dfn"} after the current one finishes running and, when
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch⑨
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree⑨
    link-type="dfn"}, also prevents `event`{.variable} from reaching any
    other objects.

`event`{.variable}` . `[`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles③ link-type="idl"}
:   Returns true or false depending on how `event`{.variable} was
    initialized. True if `event`{.variable} goes through its
    [target](#event-target){#ref-for-event-target① link-type="dfn"}'s
    [ancestors](#concept-tree-ancestor){#ref-for-concept-tree-ancestor②
    link-type="dfn"} in reverse [tree
    order](#concept-tree-order){#ref-for-concept-tree-order④
    link-type="dfn"}; otherwise false.

`event`{.variable}` . `[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable② link-type="idl"}
:   Returns true or false depending on how `event`{.variable} was
    initialized. Its return value does not always carry meaning, but
    true can indicate that part of the operation during which
    `event`{.variable} was
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⓪
    link-type="dfn"}, can be canceled by invoking the
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault②
    link-type="idl"} method.

`event`{.variable}` . `[`preventDefault`](#dom-event-preventdefault){#ref-for-dom-event-preventdefault③ .idl-code link-type="method"}`()`
:   If invoked when the
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable③
    link-type="idl"} attribute value is true, and while executing a
    listener for the `event`{.variable} with
    [`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive
    link-type="idl"} set to false, signals to the operation that caused
    `event`{.variable} to be
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①①
    link-type="dfn"} that it needs to be canceled.

`event`{.variable}` . `[`defaultPrevented`{.idl}](#dom-event-defaultprevented){#ref-for-dom-event-defaultprevented① link-type="idl"}
:   Returns true if
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault④
    link-type="idl"} was invoked successfully to indicate cancelation;
    otherwise false.

`event`{.variable}` . `[`composed`{.idl}](#dom-event-composed){#ref-for-dom-event-composed① link-type="idl"}
:   Returns true or false depending on how `event`{.variable} was
    initialized. True if `event`{.variable} invokes listeners past a
    [`ShadowRoot`{.idl}](#shadowroot){#ref-for-shadowroot
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①
    link-type="dfn"} that is the
    [root](#concept-tree-root){#ref-for-concept-tree-root③
    link-type="dfn"} of its
    [target](#event-target){#ref-for-event-target② link-type="dfn"};
    otherwise false.

`event`{.variable}` . `[`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted① link-type="idl"}
:   Returns true if `event`{.variable} was
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①②
    link-type="dfn"} by the user agent, and false otherwise.

`event`{.variable}` . `[`timeStamp`{.idl}](#dom-event-timestamp){#ref-for-dom-event-timestamp① link-type="idl"}
:   Returns the `event`{.variable}'s timestamp as the number of
    milliseconds measured relative to the occurrence.

The [`type`]{#dom-event-type .dfn .dfn-paneled .idl-code dfn-for="Event"
dfn-type="attribute" export=""} attribute must return the value it was
initialized to. When an [event](#concept-event){#ref-for-concept-event②②
link-type="dfn"} is created the attribute must be initialized to the
empty string.

The [`target`]{#dom-event-target .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this
link-type="dfn"}'s [target](#event-target){#ref-for-event-target③
link-type="dfn"}.

The [`srcElement`]{#dom-event-srcelement .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①
link-type="dfn"}'s [target](#event-target){#ref-for-event-target④
link-type="dfn"}.

The [`currentTarget`]{#dom-event-currenttarget .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="attribute" export=""} attribute must
return the value it was initialized to. When an
[event](#concept-event){#ref-for-concept-event②③ link-type="dfn"} is
created the attribute must be initialized to null.

::: {.algorithm algorithm="composedPath()" algorithm-for="Event"}
The [`composedPath()`]{#dom-event-composedpath .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="method" export=""} method steps are:

1.  Let `composedPath`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list③
    link-type="dfn"}.

2.  Let `path`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②
    link-type="dfn"}'s [path](#event-path){#ref-for-event-path③
    link-type="dfn"}.

3.  If `path`{.variable} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty
    link-type="dfn"}, then return `composedPath`{.variable}.

4.  Let `currentTarget`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③
    link-type="dfn"}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget④
    link-type="idl"} attribute value.

5.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert
    link-type="dfn"}: `currentTarget`{.variable} is an
    [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑦
    link-type="idl"} object.

6.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append
    link-type="dfn"} `currentTarget`{.variable} to
    `composedPath`{.variable}.

7.  Let `currentTargetIndex`{.variable} be 0.

8.  Let `currentTargetHiddenSubtreeLevel`{.variable} be 0.

9.  Let `index`{.variable} be `path`{.variable}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size
    link-type="dfn"} − 1.

10. While `index`{.variable} is greater than or equal to 0:

    1.  If `path`{.variable}\[`index`{.variable}\]\'s
        [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree
        link-type="dfn"} is true, then increase
        `currentTargetHiddenSubtreeLevel`{.variable} by 1.

    2.  If `path`{.variable}\[`index`{.variable}\]\'s [invocation
        target](#event-path-invocation-target){#ref-for-event-path-invocation-target①
        link-type="dfn"} is `currentTarget`{.variable}, then set
        `currentTargetIndex`{.variable} to `index`{.variable} and
        [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break
        link-type="dfn"}.

    3.  If `path`{.variable}\[`index`{.variable}\]\'s
        [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree
        link-type="dfn"} is true, then decrease
        `currentTargetHiddenSubtreeLevel`{.variable} by 1.

    4.  Decrease `index`{.variable} by 1.

11. Let `currentHiddenLevel`{.variable} and `maxHiddenLevel`{.variable}
    be `currentTargetHiddenSubtreeLevel`{.variable}.

12. Set `index`{.variable} to `currentTargetIndex`{.variable} − 1.

13. While `index`{.variable} is greater than or equal to 0:

    1.  If `path`{.variable}\[`index`{.variable}\]\'s
        [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree①
        link-type="dfn"} is true, then increase
        `currentHiddenLevel`{.variable} by 1.

    2.  If `currentHiddenLevel`{.variable} is less than or equal to
        `maxHiddenLevel`{.variable}, then
        [prepend](https://infra.spec.whatwg.org/#list-prepend){#ref-for-list-prepend
        link-type="dfn"} `path`{.variable}\[`index`{.variable}\]\'s
        [invocation
        target](#event-path-invocation-target){#ref-for-event-path-invocation-target②
        link-type="dfn"} to `composedPath`{.variable}.

    3.  If `path`{.variable}\[`index`{.variable}\]\'s
        [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree①
        link-type="dfn"} is true:

        1.  Decrease `currentHiddenLevel`{.variable} by 1.

        2.  If `currentHiddenLevel`{.variable} is less than
            `maxHiddenLevel`{.variable}, then set
            `maxHiddenLevel`{.variable} to
            `currentHiddenLevel`{.variable}.

    4.  Decrease `index`{.variable} by 1.

14. Set `currentHiddenLevel`{.variable} and `maxHiddenLevel`{.variable}
    to `currentTargetHiddenSubtreeLevel`{.variable}.

15. Set `index`{.variable} to `currentTargetIndex`{.variable} + 1.

16. While `index`{.variable} is less than `path`{.variable}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size①
    link-type="dfn"}:

    1.  If `path`{.variable}\[`index`{.variable}\]\'s
        [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree②
        link-type="dfn"} is true, then increase
        `currentHiddenLevel`{.variable} by 1.

    2.  If `currentHiddenLevel`{.variable} is less than or equal to
        `maxHiddenLevel`{.variable}, then
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①
        link-type="dfn"} `path`{.variable}\[`index`{.variable}\]\'s
        [invocation
        target](#event-path-invocation-target){#ref-for-event-path-invocation-target③
        link-type="dfn"} to `composedPath`{.variable}.

    3.  If `path`{.variable}\[`index`{.variable}\]\'s
        [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree②
        link-type="dfn"} is true:

        1.  Decrease `currentHiddenLevel`{.variable} by 1.

        2.  If `currentHiddenLevel`{.variable} is less than
            `maxHiddenLevel`{.variable}, then set
            `maxHiddenLevel`{.variable} to
            `currentHiddenLevel`{.variable}.

    4.  Increase `index`{.variable} by 1.

17. Return `composedPath`{.variable}.
:::

The [`eventPhase`]{#dom-event-eventphase .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attribute must return
the value it was initialized to, which must be one of the following:

[`NONE`]{#dom-event-none .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 0)
:   [Events](#concept-event){#ref-for-concept-event②④ link-type="dfn"}
    not currently
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①③
    link-type="dfn"} are in this phase.

[`CAPTURING_PHASE`]{#dom-event-capturing_phase .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 1)
:   When an [event](#concept-event){#ref-for-concept-event②⑤
    link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①④
    link-type="dfn"} to an object that
    [participates](#concept-tree-participate){#ref-for-concept-tree-participate②
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree①⓪
    link-type="dfn"} it will be in this phase before it reaches its
    [target](#event-target){#ref-for-event-target⑤ link-type="dfn"}.

[`AT_TARGET`]{#dom-event-at_target .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 2)
:   When an [event](#concept-event){#ref-for-concept-event②⑥
    link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑤
    link-type="dfn"} it will be in this phase on its
    [target](#event-target){#ref-for-event-target⑥ link-type="dfn"}.

[`BUBBLING_PHASE`]{#dom-event-bubbling_phase .dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="const" export=""} (numeric value 3)
:   When an [event](#concept-event){#ref-for-concept-event②⑦
    link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑥
    link-type="dfn"} to an object that
    [participates](#concept-tree-participate){#ref-for-concept-tree-participate③
    link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree①①
    link-type="dfn"} it will be in this phase after it reaches its
    [target](#event-target){#ref-for-event-target⑦ link-type="dfn"}.

Initially the attribute must be initialized to
[`NONE`{.idl}](#dom-event-none){#ref-for-dom-event-none②
link-type="idl"}.

------------------------------------------------------------------------

Each [event](#concept-event){#ref-for-concept-event②⑧ link-type="dfn"}
has the following associated flags that are all initially unset:

- [stop propagation flag]{#stop-propagation-flag .dfn .dfn-paneled
  dfn-for="Event" dfn-type="dfn" export=""}
- [stop immediate propagation flag]{#stop-immediate-propagation-flag
  .dfn .dfn-paneled dfn-for="Event" dfn-type="dfn" export=""}
- [canceled flag]{#canceled-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}
- [in passive listener flag]{#in-passive-listener-flag .dfn .dfn-paneled
  dfn-for="Event" dfn-type="dfn" export=""}
- [composed flag]{#composed-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}
- [initialized flag]{#initialized-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}
- [dispatch flag]{#dispatch-flag .dfn .dfn-paneled dfn-for="Event"
  dfn-type="dfn" export=""}

The [`stopPropagation()`]{#dom-event-stoppropagation .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="method" export=""} method steps are
to set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag
link-type="dfn"}.

The [`cancelBubble`]{#dom-event-cancelbubble .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag①
link-type="dfn"} is set; otherwise false.

The
[`cancelBubble`{.idl}](#dom-event-cancelbubble){#ref-for-dom-event-cancelbubble①
link-type="idl"} setter steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag②
link-type="dfn"} if the given value is true; otherwise do nothing.

The [`stopImmediatePropagation()`]{#dom-event-stopimmediatepropagation
.dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="method" export=""}
method steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦
link-type="dfn"}'s [stop propagation
flag](#stop-propagation-flag){#ref-for-stop-propagation-flag③
link-type="dfn"} and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧
link-type="dfn"}'s [stop immediate propagation
flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag
link-type="dfn"}.

The [`bubbles`]{#dom-event-bubbles .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} and
[`cancelable`]{#dom-event-cancelable .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attributes must return
the values they were initialized to.

To [set the canceled flag]{#set-the-canceled-flag .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an
[event](#concept-event){#ref-for-concept-event②⑨ link-type="dfn"}
`event`{.variable}, if `event`{.variable}'s
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable④
link-type="idl"} attribute value is true and `event`{.variable}'s [in
passive listener
flag](#in-passive-listener-flag){#ref-for-in-passive-listener-flag
link-type="dfn"} is unset, then set `event`{.variable}'s [canceled
flag](#canceled-flag){#ref-for-canceled-flag link-type="dfn"}, and do
nothing otherwise.

The [`returnValue`]{#dom-event-returnvalue .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return false if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨
link-type="dfn"}'s [canceled
flag](#canceled-flag){#ref-for-canceled-flag① link-type="dfn"} is set;
otherwise true.

The
[`returnValue`{.idl}](#dom-event-returnvalue){#ref-for-dom-event-returnvalue①
link-type="idl"} setter steps are to [set the canceled
flag](#set-the-canceled-flag){#ref-for-set-the-canceled-flag
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪
link-type="dfn"} if the given value is false; otherwise do nothing.

The [`preventDefault()`]{#dom-event-preventdefault .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="method" export=""} method steps are
to [set the canceled
flag](#set-the-canceled-flag){#ref-for-set-the-canceled-flag①
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①
link-type="dfn"}.

There are scenarios where invoking
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑤
link-type="idl"} has no effect. User agents are encouraged to log the
precise cause in a developer console, to aid debugging.

The [`defaultPrevented`]{#dom-event-defaultprevented .dfn .dfn-paneled
.idl-code dfn-for="Event" dfn-type="attribute" export=""} getter steps
are to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②
link-type="dfn"}'s [canceled
flag](#canceled-flag){#ref-for-canceled-flag② link-type="dfn"} is set;
otherwise false.

The [`composed`]{#dom-event-composed .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③
link-type="dfn"}'s [composed
flag](#composed-flag){#ref-for-composed-flag link-type="dfn"} is set;
otherwise false.

------------------------------------------------------------------------

The [`isTrusted`]{#dom-event-istrusted .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attribute must return
the value it was initialized to. When an
[event](#concept-event){#ref-for-concept-event③⓪ link-type="dfn"} is
created the attribute must be initialized to false.

[`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted②
link-type="idl"} is a convenience that indicates whether an
[event](#concept-event){#ref-for-concept-event③① link-type="dfn"} is
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑦
link-type="dfn"} by the user agent (as opposed to using
[`dispatchEvent()`{.idl}](#dom-eventtarget-dispatchevent){#ref-for-dom-eventtarget-dispatchevent①
link-type="idl"}). The sole legacy exception is
[`click()`{.idl}](https://html.spec.whatwg.org/multipage/interaction.html#dom-click){#ref-for-dom-click
link-type="idl"}, which causes the user agent to dispatch an
[event](#concept-event){#ref-for-concept-event③② link-type="dfn"} whose
[`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted③
link-type="idl"} attribute is initialized to false.

The [`timeStamp`]{#dom-event-timestamp .dfn .dfn-paneled .idl-code
dfn-for="Event" dfn-type="attribute" export=""} attribute must return
the value it was initialized to.

------------------------------------------------------------------------

::: {.algorithm algorithm="initialize" algorithm-for="Event"}
To [initialize]{#concept-event-initialize .dfn .dfn-paneled
dfn-for="Event" dfn-type="dfn" export=""} an `event`{.variable}, with
`type`{.variable}, `bubbles`{.variable}, and `cancelable`{.variable},
run these steps:

1.  Set `event`{.variable}'s [initialized
    flag](#initialized-flag){#ref-for-initialized-flag link-type="dfn"}.

2.  Unset `event`{.variable}'s [stop propagation
    flag](#stop-propagation-flag){#ref-for-stop-propagation-flag④
    link-type="dfn"}, [stop immediate propagation
    flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag①
    link-type="dfn"}, and [canceled
    flag](#canceled-flag){#ref-for-canceled-flag③ link-type="dfn"}.

3.  Set `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted④
    link-type="idl"} attribute to false.

4.  Set `event`{.variable}'s
    [target](#event-target){#ref-for-event-target⑧ link-type="dfn"} to
    null.

5.  Set `event`{.variable}'s
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑤
    link-type="idl"} attribute to `type`{.variable}.

6.  Set `event`{.variable}'s
    [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles④
    link-type="idl"} attribute to `bubbles`{.variable}.

7.  Set `event`{.variable}'s
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑤
    link-type="idl"} attribute to `cancelable`{.variable}.
:::

::: {.algorithm algorithm="initEvent(type, bubbles, cancelable)" algorithm-for="Event"}
The
[`initEvent(``type`{.variable}`, ``bubbles`{.variable}`, ``cancelable`{.variable}`)`]{#dom-event-initevent
.dfn .dfn-paneled .idl-code dfn-for="Event" dfn-type="method" export=""
lt="initEvent(type, bubbles, cancelable)|initEvent(type, bubbles)|initEvent(type)"}
method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④
    link-type="dfn"}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag link-type="dfn"} is
    set, then return.

2.  [Initialize](#concept-event-initialize){#ref-for-concept-event-initialize
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤
    link-type="dfn"} with `type`{.variable}, `bubbles`{.variable}, and
    `cancelable`{.variable}.
:::

[`initEvent()`{.idl}](#dom-event-initevent){#ref-for-dom-event-initevent①
link-type="idl"} is redundant with
[event](#concept-event){#ref-for-concept-event③③ link-type="dfn"}
constructors and incapable of setting
[`composed`{.idl}](#dom-event-composed){#ref-for-dom-event-composed②
link-type="idl"}. It has to be supported for legacy content.

### [2.3. ]{.secno}[Legacy extensions to the [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window link-type="idl"} interface]{.content}[](#interface-window-extensions){.self-link} {#interface-window-extensions .heading .settled level="2.3"}

``` {.idl .highlight .def}
partial interface Window {
  [Replaceable] readonly attribute (Event or undefined) event; // legacy
};
```

Each
[`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window②
link-type="idl"} object has an associated [current
event]{#window-current-event .dfn .dfn-paneled dfn-for="Window"
dfn-type="dfn" noexport=""} (undefined or an
[`Event`{.idl}](#event){#ref-for-event④ link-type="idl"} object). Unless
stated otherwise it is undefined.

The [`event`]{#dom-window-event .dfn .dfn-paneled .idl-code
dfn-for="Window" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥
link-type="dfn"}'s [current
event](#window-current-event){#ref-for-window-current-event
link-type="dfn"}.

Web developers are strongly encouraged to instead rely on the
[`Event`{.idl}](#event){#ref-for-event⑤ link-type="idl"} object passed
to event listeners, as that will result in more portable code. This
attribute is not available in workers or worklets, and is inaccurate for
events dispatched in [shadow
trees](#concept-shadow-tree){#ref-for-concept-shadow-tree①
link-type="dfn"}.

### [2.4. ]{.secno}[Interface [`CustomEvent`{.idl}](#customevent){#ref-for-customevent link-type="idl"}]{.content}[](#interface-customevent){.self-link} {#interface-customevent .heading .settled level="2.4"}

``` {.idl .highlight .def}
[Exposed=*]
interface CustomEvent : Event {
  constructor(DOMString type, optional CustomEventInit eventInitDict = {});

  readonly attribute any detail;

  undefined initCustomEvent(DOMString type, optional boolean bubbles = false, optional boolean cancelable = false, optional any detail = null); // legacy
};

dictionary CustomEventInit : EventInit {
  any detail = null;
};
```

[Events](#concept-event){#ref-for-concept-event③④ link-type="dfn"} using
the [`CustomEvent`{.idl}](#customevent){#ref-for-customevent①
link-type="idl"} interface can be used to carry custom data.

`event`{.variable}` = new `[`CustomEvent`](#dom-customevent-customevent){#ref-for-dom-customevent-customevent .idl-code link-type="constructor"}`(``type`{.variable}` [, ``eventInitDict`{.variable}`])`
:   Works analogously to the constructor for
    [`Event`{.idl}](#event){#ref-for-event⑦ link-type="idl"} except that
    the `eventInitDict`{.variable} argument now allows for setting the
    [`detail`{.idl}](#dom-customevent-detail){#ref-for-dom-customevent-detail①
    link-type="idl"} attribute too.

`event`{.variable}` . `[`detail`{.idl}](#dom-customevent-detail){#ref-for-dom-customevent-detail② link-type="idl"}
:   Returns any custom data `event`{.variable} was created with.
    Typically used for synthetic events.

The [`detail`]{#dom-customevent-detail .dfn .dfn-paneled .idl-code
dfn-for="CustomEvent" dfn-type="attribute" export=""} attribute must
return the value it was initialized to.

::: {.algorithm algorithm="initCustomEvent(type, bubbles, cancelable, detail)" algorithm-for="CustomEvent"}
The
[`initCustomEvent(``type`{.variable}`, ``bubbles`{.variable}`, ``cancelable`{.variable}`, ``detail`{.variable}`)`]{#dom-customevent-initcustomevent
.dfn .dfn-paneled .idl-code dfn-for="CustomEvent" dfn-type="method"
export=""
lt="initCustomEvent(type, bubbles, cancelable, detail)|initCustomEvent(type, bubbles, cancelable)|initCustomEvent(type, bubbles)|initCustomEvent(type)"}
method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦
    link-type="dfn"}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag① link-type="dfn"} is
    set, then return.

2.  [Initialize](#concept-event-initialize){#ref-for-concept-event-initialize①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧
    link-type="dfn"} with `type`{.variable}, `bubbles`{.variable}, and
    `cancelable`{.variable}.

3.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨
    link-type="dfn"}'s
    [`detail`{.idl}](#dom-customevent-detail){#ref-for-dom-customevent-detail③
    link-type="idl"} attribute to `detail`{.variable}.
:::

### [2.5. ]{.secno}[Constructing events]{.content}[](#constructing-events){.self-link} {#constructing-events .heading .settled level="2.5"}

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications
link-type="dfn"} may define [event constructing
steps]{#concept-event-constructor-ext .dfn .dfn-paneled dfn-type="dfn"
export=""} for all or some
[events](#concept-event){#ref-for-concept-event③⑤ link-type="dfn"}. The
algorithm is passed an [event](#concept-event){#ref-for-concept-event③⑥
link-type="dfn"} `event`{.variable} and an
[`EventInit`{.idl}](#dictdef-eventinit){#ref-for-dictdef-eventinit②
link-type="idl"} `eventInitDict`{.variable} as indicated in the [inner
event creation
steps](#inner-event-creation-steps){#ref-for-inner-event-creation-steps
link-type="dfn"}.

This construct can be used by [`Event`{.idl}](#event){#ref-for-event⑧
link-type="idl"} subclasses that have a more complex structure than a
simple 1:1 mapping between their initializing dictionary members and IDL
attributes.

::: {.algorithm algorithm="constructor" algorithm-for="Event"}
When a [constructor]{#concept-event-constructor .dfn .dfn-paneled
dfn-for="Event" dfn-type="dfn" export=""} of the
[`Event`{.idl}](#event){#ref-for-event⑨ link-type="idl"} interface, or
of an interface that inherits from the
[`Event`{.idl}](#event){#ref-for-event①⓪ link-type="idl"} interface, is
invoked, these steps must be run, given the arguments `type`{.variable}
and `eventInitDict`{.variable}:

1.  Let `event`{.variable} be the result of running the [inner event
    creation
    steps](#inner-event-creation-steps){#ref-for-inner-event-creation-steps①
    link-type="dfn"} with this interface, null, now, and
    `eventInitDict`{.variable}.

2.  Initialize `event`{.variable}'s
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑥
    link-type="idl"} attribute to `type`{.variable}.

3.  Return `event`{.variable}.
:::

::: {.algorithm algorithm="creating an event"}
To [create an event]{#concept-event-create .dfn .dfn-paneled
dfn-type="dfn" export="" lt="creating an event|create an event"} using
`eventInterface`{.variable}, which must be either
[`Event`{.idl}](#event){#ref-for-event①① link-type="idl"} or an
interface that inherits from it, and optionally given a
[realm](https://tc39.es/ecma262/#realm){#ref-for-realm link-type="dfn"}
`realm`{.variable}, run these steps:

1.  If `realm`{.variable} is not given, then set it to null.

2.  Let `dictionary`{.variable} be the result of
    [converting](https://webidl.spec.whatwg.org/#dfn-convert-ecmascript-to-idl-value){#ref-for-dfn-convert-ecmascript-to-idl-value
    link-type="dfn"} the JavaScript value undefined to the dictionary
    type accepted by `eventInterface`{.variable}'s constructor. (This
    dictionary type will either be
    [`EventInit`{.idl}](#dictdef-eventinit){#ref-for-dictdef-eventinit③
    link-type="idl"} or a dictionary that inherits from it.)

    This does not work if members are required; see
    [whatwg/dom#600](https://github.com/whatwg/dom/issues/600).

3.  Let `event`{.variable} be the result of running the [inner event
    creation
    steps](#inner-event-creation-steps){#ref-for-inner-event-creation-steps②
    link-type="dfn"} with `eventInterface`{.variable},
    `realm`{.variable}, the time of the occurrence that the event is
    signaling, and `dictionary`{.variable}.

    [](#example-timestamp-initialization){.self-link}In macOS the time
    of the occurrence for input actions is available via the `timestamp`
    property of `NSEvent` objects.

4.  Initialize `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑤
    link-type="idl"} attribute to true.

5.  Return `event`{.variable}.
:::

[Create an event](#concept-event-create){#ref-for-concept-event-create
link-type="dfn"} is meant to be used by other specifications which need
to separately
[create](#concept-event-create){#ref-for-concept-event-create①
link-type="dfn"} and
[dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑧
link-type="dfn"} events, instead of simply
[firing](#concept-event-fire){#ref-for-concept-event-fire
link-type="dfn"} them. It ensures the event's attributes are initialized
to the correct defaults.

::: {.algorithm algorithm="inner event creation steps"}
The [inner event creation steps]{#inner-event-creation-steps .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given an
`eventInterface`{.variable}, `realm`{.variable}, `time`{.variable}, and
`dictionary`{.variable}, are as follows:

1.  Let `event`{.variable} be the result of creating a new object using
    `eventInterface`{.variable}. If `realm`{.variable} is non-null, then
    use that realm; otherwise, use the default behavior defined in Web
    IDL.

    As of the time of this writing Web IDL does not yet define any
    default behavior; see
    [whatwg/webidl#135](https://github.com/whatwg/webidl/issues/135).

2.  Set `event`{.variable}'s [initialized
    flag](#initialized-flag){#ref-for-initialized-flag①
    link-type="dfn"}.

3.  Initialize `event`{.variable}'s
    [`timeStamp`{.idl}](#dom-event-timestamp){#ref-for-dom-event-timestamp②
    link-type="idl"} attribute to the [relative high resolution coarse
    time](https://w3c.github.io/hr-time/#dfn-relative-high-resolution-coarse-time){#ref-for-dfn-relative-high-resolution-coarse-time
    link-type="dfn"} given `time`{.variable} and `event`{.variable}'s
    [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global
    link-type="dfn"}.

4.  [For
    each](https://infra.spec.whatwg.org/#map-iterate){#ref-for-map-iterate
    link-type="dfn"} `member`{.variable} → `value`{.variable} of
    `dictionary`{.variable}, if `event`{.variable} has an attribute
    whose
    [identifier](https://webidl.spec.whatwg.org/#dfn-identifier){#ref-for-dfn-identifier
    link-type="dfn"} is `member`{.variable}, then initialize that
    attribute to `value`{.variable}.

5.  Run the [event constructing
    steps](#concept-event-constructor-ext){#ref-for-concept-event-constructor-ext
    link-type="dfn"} with `event`{.variable} and
    `dictionary`{.variable}.

6.  Return `event`{.variable}.
:::

### [2.6. ]{.secno}[Defining event interfaces]{.content}[](#defining-event-interfaces){.self-link} {#defining-event-interfaces .heading .settled level="2.6"}

In general, when defining a new interface that inherits from
[`Event`{.idl}](#event){#ref-for-event①② link-type="idl"} please always
ask feedback from the [WHATWG](https://whatwg.org/) or the [W3C WebApps
WG](https://www.w3.org/2008/webapps/) community.

The [`CustomEvent`{.idl}](#customevent){#ref-for-customevent②
link-type="idl"} interface can be used as starting point. However, do
not introduce any `init``*`{.variable}`Event()` methods as they are
redundant with constructors. Interfaces that inherit from the
[`Event`{.idl}](#event){#ref-for-event①③ link-type="idl"} interface that
have such a method only have it for historical reasons.

### [2.7. ]{.secno}[Interface [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑧ link-type="idl"}]{.content}[](#interface-eventtarget){.self-link} {#interface-eventtarget .heading .settled level="2.7"}

``` {.idl .highlight .def}
[Exposed=*]
interface EventTarget {
  constructor();

  undefined addEventListener(DOMString type, EventListener? callback, optional (AddEventListenerOptions or boolean) options = {});
  undefined removeEventListener(DOMString type, EventListener? callback, optional (EventListenerOptions or boolean) options = {});
  boolean dispatchEvent(Event event);
};

callback interface EventListener {
  undefined handleEvent(Event event);
};

dictionary EventListenerOptions {
  boolean capture = false;
};

dictionary AddEventListenerOptions : EventListenerOptions {
  boolean passive;
  boolean once = false;
  AbortSignal signal;
};
```

An [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget⑨
link-type="idl"} object represents a target to which an
[event](#concept-event){#ref-for-concept-event③⑦ link-type="dfn"} can be
[dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch①⑨
link-type="dfn"} when something has occurred.

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⓪
link-type="idl"} object has an associated [event listener
list]{#eventtarget-event-listener-list .dfn .dfn-paneled
dfn-for="EventTarget" dfn-type="dfn" noexport=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list④
link-type="dfn"} of zero or more [event
listeners](#concept-event-listener){#ref-for-concept-event-listener①①
link-type="dfn"}). It is initially the empty list.

An [event listener]{#concept-event-listener .dfn .dfn-paneled
dfn-type="dfn" export=""} can be used to observe a specific
[event](#concept-event){#ref-for-concept-event③⑧ link-type="dfn"} and
consists of:

- [type]{#event-listener-type .dfn .dfn-paneled dfn-for="event listener"
  dfn-type="dfn" noexport=""} (a string)
- [callback]{#event-listener-callback .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" export=""} (null or an
  [`EventListener`{.idl}](#callbackdef-eventlistener){#ref-for-callbackdef-eventlistener②
  link-type="idl"} object)
- [capture]{#event-listener-capture .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (a boolean,
  initially false)
- [passive]{#event-listener-passive .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (null or a
  boolean, initially null)
- [once]{#event-listener-once .dfn .dfn-paneled dfn-for="event listener"
  dfn-type="dfn" noexport=""} (a boolean, initially false)
- [signal]{#event-listener-signal .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (null or an
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②
  link-type="idl"} object)
- [removed]{#event-listener-removed .dfn .dfn-paneled
  dfn-for="event listener" dfn-type="dfn" noexport=""} (a boolean for
  bookkeeping purposes, initially false)

Although
[callback](#event-listener-callback){#ref-for-event-listener-callback②
link-type="dfn"} is an
[`EventListener`{.idl}](#callbackdef-eventlistener){#ref-for-callbackdef-eventlistener③
link-type="idl"} object, an [event
listener](#concept-event-listener){#ref-for-concept-event-listener①②
link-type="dfn"} is a broader concept as can be seen above.

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①①
link-type="idl"} object also has an associated [get the
parent]{#get-the-parent .dfn .dfn-paneled dfn-type="dfn" export=""}
algorithm, which takes an
[event](#concept-event){#ref-for-concept-event③⑨ link-type="dfn"}
`event`{.variable}, and returns an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①②
link-type="idl"} object. Unless specified otherwise it returns null.

[Nodes](#concept-node){#ref-for-concept-node② link-type="dfn"}, [shadow
roots](#concept-shadow-root){#ref-for-concept-shadow-root①
link-type="dfn"}, and
[documents](#concept-document){#ref-for-concept-document①
link-type="dfn"} override the [get the
parent](#get-the-parent){#ref-for-get-the-parent link-type="dfn"}
algorithm.

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①③
link-type="idl"} object can have an associated [activation
behavior]{#eventtarget-activation-behavior .dfn .dfn-paneled
dfn-for="EventTarget" dfn-type="dfn" export=""} algorithm. The
[activation
behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior
link-type="dfn"} algorithm is passed an
[event](#concept-event){#ref-for-concept-event④⓪ link-type="dfn"}, as
indicated in the
[dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch②⓪
link-type="dfn"} algorithm.

This exists because user agents perform certain actions for certain
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①④
link-type="idl"} objects, e.g., the
[`area`](https://html.spec.whatwg.org/multipage/image-maps.html#the-area-element){#ref-for-the-area-element
link-type="element"} element, in response to synthetic
[`MouseEvent`{.idl}](https://w3c.github.io/uievents/#mouseevent){#ref-for-mouseevent
link-type="idl"} [events](#concept-event){#ref-for-concept-event④①
link-type="dfn"} whose
[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑦
link-type="idl"} attribute is `click`. Web compatibility prevented it
from being removed and it is now the enshrined way of defining an
activation of something.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

Each [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑤
link-type="idl"} object that has [activation
behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior①
link-type="dfn"}, can additionally have both (not either) a
[legacy-pre-activation
behavior]{#eventtarget-legacy-pre-activation-behavior .dfn .dfn-paneled
dfn-for="EventTarget" dfn-type="dfn" export=""} algorithm and a
[legacy-canceled-activation
behavior]{#eventtarget-legacy-canceled-activation-behavior .dfn
.dfn-paneled dfn-for="EventTarget" dfn-type="dfn" export=""} algorithm.

These algorithms only exist for checkbox and radio
[`input`](https://html.spec.whatwg.org/multipage/input.html#the-input-element){#ref-for-the-input-element
link-type="element"} elements and are not to be used for anything else.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

`target`{.variable}` = new `[`EventTarget`](#dom-eventtarget-eventtarget){#ref-for-dom-eventtarget-eventtarget① .idl-code link-type="constructor"}`();`

:   Creates a new
    [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑥
    link-type="idl"} object, which can be used by developers to
    [dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch②①
    link-type="dfn"} and listen for
    [events](#concept-event){#ref-for-concept-event④② link-type="dfn"}.

`target`{.variable}` . `[`addEventListener`](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener③ .idl-code link-type="method"}`(``type`{.variable}`, ``callback`{.variable}` [, ``options`{.variable}`])`

:   Appends an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①③
    link-type="dfn"} for
    [events](#concept-event){#ref-for-concept-event④③ link-type="dfn"}
    whose [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑧
    link-type="idl"} attribute value is `type`{.variable}. The
    `callback`{.variable} argument sets the
    [callback](#event-listener-callback){#ref-for-event-listener-callback③
    link-type="dfn"} that will be invoked when the
    [event](#concept-event){#ref-for-concept-event④④ link-type="dfn"} is
    [dispatched](#concept-event-dispatch){#ref-for-concept-event-dispatch②②
    link-type="dfn"}.

    The `options`{.variable} argument sets listener-specific options.
    For compatibility this can be a boolean, in which case the method
    behaves exactly as if the value was specified as
    `options`{.variable}'s
    [`capture`{.idl}](#dom-eventlisteneroptions-capture){#ref-for-dom-eventlisteneroptions-capture
    link-type="idl"}.

    When set to true, `options`{.variable}'s
    [`capture`{.idl}](#dom-eventlisteneroptions-capture){#ref-for-dom-eventlisteneroptions-capture①
    link-type="idl"} prevents
    [callback](#event-listener-callback){#ref-for-event-listener-callback④
    link-type="dfn"} from being invoked when the
    [event](#concept-event){#ref-for-concept-event④⑤ link-type="dfn"}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase④
    link-type="idl"} attribute value is
    [`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase③
    link-type="idl"}. When false (or not present),
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑤
    link-type="dfn"} will not be invoked when
    [event](#concept-event){#ref-for-concept-event④⑥ link-type="dfn"}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑤
    link-type="idl"} attribute value is
    [`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase③
    link-type="idl"}. Either way,
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑥
    link-type="dfn"} will be invoked if
    [event](#concept-event){#ref-for-concept-event④⑦ link-type="dfn"}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑥
    link-type="idl"} attribute value is
    [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target③
    link-type="idl"}.

    When set to true, `options`{.variable}'s
    [`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive①
    link-type="idl"} indicates that the
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑦
    link-type="dfn"} will not cancel the event by invoking
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑥
    link-type="idl"}. This is used to enable performance optimizations
    described in [§ 2.8 Observing event
    listeners](#observing-event-listeners).

    When set to true, `options`{.variable}'s
    [`once`{.idl}](#dom-addeventlisteneroptions-once){#ref-for-dom-addeventlisteneroptions-once
    link-type="idl"} indicates that the
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑧
    link-type="dfn"} will only be invoked once after which the event
    listener will be removed.

    If an [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③
    link-type="idl"} is passed for `options`{.variable}'s
    [`signal`{.idl}](#dom-addeventlisteneroptions-signal){#ref-for-dom-addeventlisteneroptions-signal
    link-type="idl"}, then the event listener will be removed when
    signal is aborted.

    The [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①④
    link-type="dfn"} is appended to `target`{.variable}'s [event
    listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list
    link-type="dfn"} and is not appended if it has the same
    [type](#event-listener-type){#ref-for-event-listener-type
    link-type="dfn"},
    [callback](#event-listener-callback){#ref-for-event-listener-callback⑨
    link-type="dfn"}, and
    [capture](#event-listener-capture){#ref-for-event-listener-capture②
    link-type="dfn"}.

`target`{.variable}` . `[`removeEventListener`](#dom-eventtarget-removeeventlistener){#ref-for-dom-eventtarget-removeeventlistener② .idl-code link-type="method"}`(``type`{.variable}`, ``callback`{.variable}` [, ``options`{.variable}`])`

:   Removes the [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①⑤
    link-type="dfn"} in `target`{.variable}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list①
    link-type="dfn"} with the same `type`{.variable},
    `callback`{.variable}, and `options`{.variable}.

`target`{.variable}` . `[`dispatchEvent`](#dom-eventtarget-dispatchevent){#ref-for-dom-eventtarget-dispatchevent③ .idl-code link-type="method"}`(``event`{.variable}`)`

:   [Dispatches](#concept-event-dispatch){#ref-for-concept-event-dispatch②③
    link-type="dfn"} a synthetic event `event`{.variable} to
    `target`{.variable} and returns true if either `event`{.variable}'s
    [`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑥
    link-type="idl"} attribute value is false or its
    [`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑦
    link-type="idl"} method was not invoked; otherwise false.

::: {.algorithm algorithm="flatten" algorithm-for="Event"}
To [flatten]{#concept-flatten-options .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""} `options`{.variable}, run these steps:

1.  If `options`{.variable} is a boolean, then return
    `options`{.variable}.

2.  Return
    `options`{.variable}\[\"[`capture`{.idl}](#dom-eventlisteneroptions-capture){#ref-for-dom-eventlisteneroptions-capture②
    link-type="idl"}\"\].
:::

::: {.algorithm algorithm="flatten more" algorithm-for="Event"}
To [flatten more]{#event-flatten-more .dfn .dfn-paneled dfn-for="Event"
dfn-type="dfn" export=""} `options`{.variable}, run these steps:

1.  Let `capture`{.variable} be the result of
    [flattening](#concept-flatten-options){#ref-for-concept-flatten-options
    link-type="dfn"} `options`{.variable}.

2.  Let `once`{.variable} be false.

3.  Let `passive`{.variable} and `signal`{.variable} be null.

4.  If `options`{.variable} is a
    [dictionary](https://webidl.spec.whatwg.org/#dfn-dictionary){#ref-for-dfn-dictionary
    link-type="dfn"}:

    1.  Set `once`{.variable} to
        `options`{.variable}\[\"[`once`{.idl}](#dom-addeventlisteneroptions-once){#ref-for-dom-addeventlisteneroptions-once①
        link-type="idl"}\"\].

    2.  If
        `options`{.variable}\[\"[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive②
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists
        link-type="dfn"}, then set `passive`{.variable} to
        `options`{.variable}\[\"[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive③
        link-type="idl"}\"\].

    3.  If
        `options`{.variable}\[\"[`signal`{.idl}](#dom-addeventlisteneroptions-signal){#ref-for-dom-addeventlisteneroptions-signal①
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①
        link-type="dfn"}, then set `signal`{.variable} to
        `options`{.variable}\[\"[`signal`{.idl}](#dom-addeventlisteneroptions-signal){#ref-for-dom-addeventlisteneroptions-signal②
        link-type="idl"}\"\].

5.  Return `capture`{.variable}, `passive`{.variable},
    `once`{.variable}, and `signal`{.variable}.
:::

The [`new EventTarget()`]{#dom-eventtarget-eventtarget .dfn .dfn-paneled
.idl-code dfn-for="EventTarget" dfn-type="constructor" export=""
lt="EventTarget()|constructor()"} constructor steps are to do nothing.

Because of the defaults stated elsewhere, the returned
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑦
link-type="idl"}'s [get the
parent](#get-the-parent){#ref-for-get-the-parent① link-type="dfn"}
algorithm will return null, and it will have no [activation
behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior②
link-type="dfn"}, [legacy-pre-activation
behavior](#eventtarget-legacy-pre-activation-behavior){#ref-for-eventtarget-legacy-pre-activation-behavior
link-type="dfn"}, or [legacy-canceled-activation
behavior](#eventtarget-legacy-canceled-activation-behavior){#ref-for-eventtarget-legacy-canceled-activation-behavior
link-type="dfn"}.

In the future we could allow custom [get the
parent](#get-the-parent){#ref-for-get-the-parent② link-type="dfn"}
algorithms. Let us know if this would be useful for your programs. For
now, all author-created
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑧
link-type="idl"}s do not participate in a tree structure.

::: {.algorithm algorithm="default passive value"}
The [default passive value]{#default-passive-value .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an event type `type`{.variable} and
an [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget①⑨
link-type="idl"} `eventTarget`{.variable}, is determined as follows:

1.  Return true if all of the following are true:

    - `type`{.variable} is one of \"`touchstart`\", \"`touchmove`\",
      \"`wheel`\", or \"`mousewheel`\".
      [\[TOUCH-EVENTS\]](#biblio-touch-events "Touch Events"){link-type="biblio"}
      [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    - `eventTarget`{.variable} is a
      [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window③
      link-type="idl"} object, or is a
      [node](#concept-node){#ref-for-concept-node③ link-type="dfn"}
      whose [node
      document](#concept-node-document){#ref-for-concept-node-document
      link-type="dfn"} is `eventTarget`{.variable}, or is a
      [node](#concept-node){#ref-for-concept-node④ link-type="dfn"}
      whose [node
      document](#concept-node-document){#ref-for-concept-node-document①
      link-type="dfn"}'s [document
      element](#document-element){#ref-for-document-element
      link-type="dfn"} is `eventTarget`{.variable}, or is a
      [node](#concept-node){#ref-for-concept-node⑤ link-type="dfn"}
      whose [node
      document](#concept-node-document){#ref-for-concept-node-document②
      link-type="dfn"}'s [body
      element](https://html.spec.whatwg.org/multipage/dom.html#the-body-element-2){#ref-for-the-body-element-2
      link-type="dfn"} is `eventTarget`{.variable}.
      [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

2.  Return false.
:::

::: {.algorithm algorithm="add an event listener"}
To [add an event listener]{#add-an-event-listener .dfn .dfn-paneled
dfn-type="dfn" export=""}, given an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②⓪
link-type="idl"} object `eventTarget`{.variable} and an [event
listener](#concept-event-listener){#ref-for-concept-event-listener①⑥
link-type="dfn"} `listener`{.variable}, run these steps:

1.  If `eventTarget`{.variable} is a
    [`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope
    link-type="idl"} object, its [service
    worker](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope-service-worker){#ref-for-serviceworkerglobalscope-service-worker
    link-type="dfn"}'s [script
    resource](https://w3c.github.io/ServiceWorker/#dfn-script-resource){#ref-for-dfn-script-resource
    link-type="dfn"}'s [has ever been evaluated
    flag](https://w3c.github.io/ServiceWorker/#dfn-has-ever-been-evaluated-flag){#ref-for-dfn-has-ever-been-evaluated-flag
    link-type="dfn"} is set, and `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type①
    link-type="dfn"} matches the
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type⑨
    link-type="idl"} attribute value of any of the [service worker
    events](https://w3c.github.io/ServiceWorker/#dfn-service-worker-events){#ref-for-dfn-service-worker-events
    link-type="dfn"}, then [report a warning to the
    console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#ref-for-report-a-warning-to-the-console
    link-type="dfn"} that this might not give the expected results.
    [\[SERVICE-WORKERS\]](#biblio-service-workers "Service Workers"){link-type="biblio"}

2.  If `listener`{.variable}'s
    [signal](#event-listener-signal){#ref-for-event-listener-signal
    link-type="dfn"} is not null and is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted
    link-type="dfn"}, then return.

3.  If `listener`{.variable}'s
    [callback](#event-listener-callback){#ref-for-event-listener-callback①⓪
    link-type="dfn"} is null, then return.

4.  If `listener`{.variable}'s
    [passive](#event-listener-passive){#ref-for-event-listener-passive
    link-type="dfn"} is null, then set it to the [default passive
    value](#default-passive-value){#ref-for-default-passive-value
    link-type="dfn"} given `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type②
    link-type="dfn"} and `eventTarget`{.variable}.

5.  If `eventTarget`{.variable}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list②
    link-type="dfn"} does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain
    link-type="dfn"} an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①⑦
    link-type="dfn"} whose
    [type](#event-listener-type){#ref-for-event-listener-type③
    link-type="dfn"} is `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type④
    link-type="dfn"},
    [callback](#event-listener-callback){#ref-for-event-listener-callback①①
    link-type="dfn"} is `listener`{.variable}'s
    [callback](#event-listener-callback){#ref-for-event-listener-callback①②
    link-type="dfn"}, and
    [capture](#event-listener-capture){#ref-for-event-listener-capture③
    link-type="dfn"} is `listener`{.variable}'s
    [capture](#event-listener-capture){#ref-for-event-listener-capture④
    link-type="dfn"}, then
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append②
    link-type="dfn"} `listener`{.variable} to `eventTarget`{.variable}'s
    [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list③
    link-type="dfn"}.

6.  If `listener`{.variable}'s
    [signal](#event-listener-signal){#ref-for-event-listener-signal①
    link-type="dfn"} is not null, then [add the
    following](#abortsignal-add){#ref-for-abortsignal-add
    link-type="dfn"} abort steps to it:

    1.  [Remove an event
        listener](#remove-an-event-listener){#ref-for-remove-an-event-listener
        link-type="dfn"} with `eventTarget`{.variable} and
        `listener`{.variable}.

The [add an event
listener](#add-an-event-listener){#ref-for-add-an-event-listener
link-type="dfn"} concept exists to ensure [event
handlers](https://html.spec.whatwg.org/multipage/webappapis.html#event-handlers){#ref-for-event-handlers
link-type="dfn"} use the same code path.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="addEventListener(type, callback, options)" algorithm-for="EventTarget"}
The
[`addEventListener(``type`{.variable}`, ``callback`{.variable}`, ``options`{.variable}`)`]{#dom-eventtarget-addeventlistener
.dfn .dfn-paneled .idl-code dfn-for="EventTarget" dfn-type="method"
export=""
lt="addEventListener(type, callback, options)|addEventListener(type, callback)"}
method steps are:

1.  Let `capture`{.variable}, `passive`{.variable}, `once`{.variable},
    and `signal`{.variable} be the result of [flattening
    more](#event-flatten-more){#ref-for-event-flatten-more
    link-type="dfn"} `options`{.variable}.

2.  [Add an event
    listener](#add-an-event-listener){#ref-for-add-an-event-listener①
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪
    link-type="dfn"} and an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener①⑧
    link-type="dfn"} whose
    [type](#event-listener-type){#ref-for-event-listener-type⑤
    link-type="dfn"} is `type`{.variable},
    [callback](#event-listener-callback){#ref-for-event-listener-callback①③
    link-type="dfn"} is `callback`{.variable},
    [capture](#event-listener-capture){#ref-for-event-listener-capture⑤
    link-type="dfn"} is `capture`{.variable},
    [passive](#event-listener-passive){#ref-for-event-listener-passive①
    link-type="dfn"} is `passive`{.variable},
    [once](#event-listener-once){#ref-for-event-listener-once
    link-type="dfn"} is `once`{.variable}, and
    [signal](#event-listener-signal){#ref-for-event-listener-signal②
    link-type="dfn"} is `signal`{.variable}.
:::

::: {.algorithm algorithm="remove an event listener"}
To [remove an event listener]{#remove-an-event-listener .dfn
.dfn-paneled dfn-type="dfn" export=""}, given an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②①
link-type="idl"} object `eventTarget`{.variable} and an [event
listener](#concept-event-listener){#ref-for-concept-event-listener①⑨
link-type="dfn"} `listener`{.variable}, run these steps:

1.  If `eventTarget`{.variable} is a
    [`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope①
    link-type="idl"} object and its [service
    worker](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope-service-worker){#ref-for-serviceworkerglobalscope-service-worker①
    link-type="dfn"}'s [set of event types to
    handle](https://w3c.github.io/ServiceWorker/#dfn-set-of-event-types-to-handle){#ref-for-dfn-set-of-event-types-to-handle
    link-type="dfn"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain①
    link-type="dfn"} `listener`{.variable}'s
    [type](#event-listener-type){#ref-for-event-listener-type⑥
    link-type="dfn"}, then [report a warning to the
    console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#ref-for-report-a-warning-to-the-console①
    link-type="dfn"} that this might not give the expected results.
    [\[SERVICE-WORKERS\]](#biblio-service-workers "Service Workers"){link-type="biblio"}

2.  Set `listener`{.variable}'s
    [removed](#event-listener-removed){#ref-for-event-listener-removed
    link-type="dfn"} to true and
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove
    link-type="dfn"} `listener`{.variable} from
    `eventTarget`{.variable}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list④
    link-type="dfn"}.

HTML needs this to define event handlers.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="remove all event listeners"}
To [remove all event listeners]{#remove-all-event-listeners .dfn
.dfn-paneled dfn-type="dfn" export=""}, given an
[`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②②
link-type="idl"} object `eventTarget`{.variable}, [for
each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①
link-type="dfn"} `listener`{.variable} of `eventTarget`{.variable}'s
[event listener
list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑤
link-type="dfn"}, [remove an event
listener](#remove-an-event-listener){#ref-for-remove-an-event-listener①
link-type="dfn"} with `eventTarget`{.variable} and
`listener`{.variable}.

HTML needs this to define `document.open()`.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

::: {.algorithm algorithm="removeEventListener(type, callback, options)" algorithm-for="EventTarget"}
The
[`removeEventListener(``type`{.variable}`, ``callback`{.variable}`, ``options`{.variable}`)`]{#dom-eventtarget-removeeventlistener
.dfn .dfn-paneled .idl-code dfn-for="EventTarget" dfn-type="method"
export=""
lt="removeEventListener(type, callback, options)|removeEventListener(type, callback)"}
method steps are:

1.  Let `capture`{.variable} be the result of
    [flattening](#concept-flatten-options){#ref-for-concept-flatten-options①
    link-type="dfn"} `options`{.variable}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①
    link-type="dfn"}'s [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑥
    link-type="dfn"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain②
    link-type="dfn"} an [event
    listener](#concept-event-listener){#ref-for-concept-event-listener②⓪
    link-type="dfn"} whose
    [type](#event-listener-type){#ref-for-event-listener-type⑦
    link-type="dfn"} is `type`{.variable},
    [callback](#event-listener-callback){#ref-for-event-listener-callback①④
    link-type="dfn"} is `callback`{.variable}, and
    [capture](#event-listener-capture){#ref-for-event-listener-capture⑥
    link-type="dfn"} is `capture`{.variable}, then [remove an event
    listener](#remove-an-event-listener){#ref-for-remove-an-event-listener②
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②
    link-type="dfn"} and that [event
    listener](#concept-event-listener){#ref-for-concept-event-listener②①
    link-type="dfn"}.

The event listener list will not contain multiple event listeners with
equal `type`{.variable}, `callback`{.variable}, and
`capture`{.variable}, as [add an event
listener](#add-an-event-listener){#ref-for-add-an-event-listener②
link-type="dfn"} prevents that.
:::

::: {.algorithm algorithm="dispatchEvent(event)" algorithm-for="EventTarget"}
The
[`dispatchEvent(``event`{.variable}`)`]{#dom-eventtarget-dispatchevent
.dfn .dfn-paneled .idl-code dfn-for="EventTarget" dfn-type="method"
export=""} method steps are:

1.  If `event`{.variable}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag② link-type="dfn"} is
    set, or if its [initialized
    flag](#initialized-flag){#ref-for-initialized-flag② link-type="dfn"}
    is not set, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥
    link-type="dfn"} an
    \"[`InvalidStateError`{.idl}](https://webidl.spec.whatwg.org/#invalidstateerror){#ref-for-invalidstateerror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥
    link-type="idl"}.

2.  Initialize `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑥
    link-type="idl"} attribute to false.

3.  Return the result of
    [dispatching](#concept-event-dispatch){#ref-for-concept-event-dispatch②④
    link-type="dfn"} `event`{.variable} to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③
    link-type="dfn"}.
:::

### [2.8. ]{.secno}[Observing event listeners]{.content}[](#observing-event-listeners){.self-link} {#observing-event-listeners .heading .settled level="2.8"}

In general, developers do not expect the presence of an [event
listener](#concept-event-listener){#ref-for-concept-event-listener②②
link-type="dfn"} to be observable. The impact of an [event
listener](#concept-event-listener){#ref-for-concept-event-listener②③
link-type="dfn"} is determined by its
[callback](#event-listener-callback){#ref-for-event-listener-callback①⑤
link-type="dfn"}. That is, a developer adding a no-op [event
listener](#concept-event-listener){#ref-for-concept-event-listener②④
link-type="dfn"} would not expect it to have any side effects.

Unfortunately, some event APIs have been designed such that implementing
them efficiently requires observing [event
listeners](#concept-event-listener){#ref-for-concept-event-listener②⑤
link-type="dfn"}. This can make the presence of listeners observable in
that even empty listeners can have a dramatic performance impact on the
behavior of the application. For example, touch and wheel events which
can be used to block asynchronous scrolling. In some cases this problem
can be mitigated by specifying the event to be
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑦
link-type="idl"} only when there is at least one
non-[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive④
link-type="idl"} listener. For example,
non-[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive⑤
link-type="idl"}
[`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent①
link-type="idl"} listeners must block scrolling, but if all listeners
are
[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive⑥
link-type="idl"} then scrolling can be allowed to start [in
parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel
link-type="dfn"} by making the
[`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent②
link-type="idl"} uncancelable (so that calls to
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑧
link-type="idl"} are ignored). So code dispatching an event is able to
observe the absence of
non-[`passive`{.idl}](#dom-addeventlisteneroptions-passive){#ref-for-dom-addeventlisteneroptions-passive⑦
link-type="idl"} listeners, and use that to clear the
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑧
link-type="idl"} property of the event being dispatched.

Ideally, any new event APIs are defined such that they do not need this
property. (Use [whatwg/dom](https://github.com/whatwg/dom/issues) for
discussion.)

::: {.algorithm algorithm="legacy-obtain service worker fetch event listener callbacks"}
To [legacy-obtain service worker fetch event listener
callbacks]{#legacy-obtain-service-worker-fetch-event-listener-callbacks
.dfn .dfn-paneled dfn-type="dfn" export=""} given a
[`ServiceWorkerGlobalScope`{.idl}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){#ref-for-serviceworkerglobalscope②
link-type="idl"} `global`{.variable}, run these steps. They return a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list⑤
link-type="dfn"} of
[`EventListener`{.idl}](#callbackdef-eventlistener){#ref-for-callbackdef-eventlistener④
link-type="idl"} objects.

1.  Let `callbacks`{.variable} be « ».

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②
    link-type="dfn"} `listener`{.variable} of `global`{.variable}'s
    [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑦
    link-type="dfn"}:

    1.  If `listener`{.variable}'s
        [type](#event-listener-type){#ref-for-event-listener-type⑧
        link-type="dfn"} is \"`fetch`\", and `listener`{.variable}'s
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑥
        link-type="dfn"} is not null, then
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append③
        link-type="dfn"} `listener`{.variable}'s
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑦
        link-type="dfn"} to `callbacks`{.variable}.

3.  Return `callbacks`{.variable}.
:::

### [2.9. ]{.secno}[Dispatching events]{.content}[](#dispatching-events){.self-link} {#dispatching-events .heading .settled level="2.9"}

::: {.algorithm algorithm="dispatch"}
To [dispatch]{#concept-event-dispatch .dfn .dfn-paneled dfn-type="dfn"
export=""} an `event`{.variable} to a `target`{.variable}, with an
optional `legacy target override flag`{.variable} and an optional
`legacyOutputDidListenersThrowFlag`{.variable}, run these steps:

1.  Set `event`{.variable}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag③ link-type="dfn"}.

2.  Let `targetOverride`{.variable} be `target`{.variable}, if
    `legacy target override flag`{.variable} is not given, and
    `target`{.variable}'s [associated
    `Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window
    link-type="dfn"} otherwise.
    [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    `legacy target override flag`{.variable} is only used by HTML and
    only when `target`{.variable} is a
    [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window④
    link-type="idl"} object.

3.  Let `activationTarget`{.variable} be null.

4.  Let `relatedTarget`{.variable} be the result of
    [retargeting](#retarget){#ref-for-retarget link-type="dfn"}
    `event`{.variable}'s
    [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget①
    link-type="dfn"} against `target`{.variable}.

5.  Let `clearTargets`{.variable} be false.

6.  If `target`{.variable} is not `relatedTarget`{.variable} or
    `target`{.variable} is `event`{.variable}'s
    [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget②
    link-type="dfn"}:

    1.  Let `touchTargets`{.variable} be a new
        [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑥
        link-type="dfn"}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate③
        link-type="dfn"} `touchTarget`{.variable} of
        `event`{.variable}'s [touch target
        list](#event-touch-target-list){#ref-for-event-touch-target-list①
        link-type="dfn"},
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append④
        link-type="dfn"} the result of
        [retargeting](#retarget){#ref-for-retarget① link-type="dfn"}
        `touchTarget`{.variable} against `target`{.variable} to
        `touchTargets`{.variable}.

    3.  [Append to an event
        path](#concept-event-path-append){#ref-for-concept-event-path-append
        link-type="dfn"} with `event`{.variable}, `target`{.variable},
        `targetOverride`{.variable}, `relatedTarget`{.variable},
        `touchTargets`{.variable}, and false.

    4.  Let `isActivationEvent`{.variable} be true, if
        `event`{.variable} is a
        [`MouseEvent`{.idl}](https://w3c.github.io/uievents/#mouseevent){#ref-for-mouseevent①
        link-type="idl"} object and `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⓪
        link-type="idl"} attribute is \"`click`\"; otherwise false.

    5.  If `isActivationEvent`{.variable} is true and
        `target`{.variable} has [activation
        behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior③
        link-type="dfn"}, then set `activationTarget`{.variable} to
        `target`{.variable}.

    6.  Let `slottable`{.variable} be `target`{.variable}, if
        `target`{.variable} is a
        [slottable](#concept-slotable){#ref-for-concept-slotable
        link-type="dfn"} and is
        [assigned](#slotable-assigned){#ref-for-slotable-assigned
        link-type="dfn"}, and null otherwise.

    7.  Let `slot-in-closed-tree`{.variable} be false.

    8.  Let `parent`{.variable} be the result of invoking
        `target`{.variable}'s [get the
        parent](#get-the-parent){#ref-for-get-the-parent③
        link-type="dfn"} with `event`{.variable}.

    9.  While `parent`{.variable} is non-null:

        1.  If `slottable`{.variable} is non-null:

            1.  Assert: `parent`{.variable} is a
                [slot](#concept-slot){#ref-for-concept-slot
                link-type="dfn"}.

            2.  Set `slottable`{.variable} to null.

            3.  If `parent`{.variable}'s
                [root](#concept-tree-root){#ref-for-concept-tree-root④
                link-type="dfn"} is a [shadow
                root](#concept-shadow-root){#ref-for-concept-shadow-root②
                link-type="dfn"} whose
                [mode](#shadowroot-mode){#ref-for-shadowroot-mode①
                link-type="dfn"} is \"`closed`\", then set
                `slot-in-closed-tree`{.variable} to true.

        2.  If `parent`{.variable} is a
            [slottable](#concept-slotable){#ref-for-concept-slotable①
            link-type="dfn"} and is
            [assigned](#slotable-assigned){#ref-for-slotable-assigned①
            link-type="dfn"}, then set `slottable`{.variable} to
            `parent`{.variable}.

        3.  Let `relatedTarget`{.variable} be the result of
            [retargeting](#retarget){#ref-for-retarget② link-type="dfn"}
            `event`{.variable}'s
            [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget③
            link-type="dfn"} against `parent`{.variable}.

        4.  Let `touchTargets`{.variable} be a new
            [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑦
            link-type="dfn"}.

        5.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate④
            link-type="dfn"} `touchTarget`{.variable} of
            `event`{.variable}'s [touch target
            list](#event-touch-target-list){#ref-for-event-touch-target-list②
            link-type="dfn"},
            [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑤
            link-type="dfn"} the result of
            [retargeting](#retarget){#ref-for-retarget③ link-type="dfn"}
            `touchTarget`{.variable} against `parent`{.variable} to
            `touchTargets`{.variable}.

        6.  If `parent`{.variable} is a
            [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑤
            link-type="idl"} object, or `parent`{.variable} is a
            [node](#concept-node){#ref-for-concept-node⑥
            link-type="dfn"} and `target`{.variable}'s
            [root](#concept-tree-root){#ref-for-concept-tree-root⑤
            link-type="dfn"} is a [shadow-including inclusive
            ancestor](#concept-shadow-including-inclusive-ancestor){#ref-for-concept-shadow-including-inclusive-ancestor
            link-type="dfn"} of `parent`{.variable}:

            1.  If `isActivationEvent`{.variable} is true,
                `event`{.variable}'s
                [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles⑤
                link-type="idl"} attribute is true,
                `activationTarget`{.variable} is null, and
                `parent`{.variable} has [activation
                behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior④
                link-type="dfn"}, then set `activationTarget`{.variable}
                to `parent`{.variable}.

            2.  [Append to an event
                path](#concept-event-path-append){#ref-for-concept-event-path-append①
                link-type="dfn"} with `event`{.variable},
                `parent`{.variable}, null, `relatedTarget`{.variable},
                `touchTargets`{.variable}, and
                `slot-in-closed-tree`{.variable}.

        7.  Otherwise, if `parent`{.variable} is
            `relatedTarget`{.variable}, then set `parent`{.variable} to
            null.

        8.  Otherwise:

            1.  Set `target`{.variable} to `parent`{.variable}.

            2.  If `isActivationEvent`{.variable} is true,
                `activationTarget`{.variable} is null, and
                `target`{.variable} has [activation
                behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior⑤
                link-type="dfn"}, then set `activationTarget`{.variable}
                to `target`{.variable}.

            3.  [Append to an event
                path](#concept-event-path-append){#ref-for-concept-event-path-append②
                link-type="dfn"} with `event`{.variable},
                `parent`{.variable}, `target`{.variable},
                `relatedTarget`{.variable}, `touchTargets`{.variable},
                and `slot-in-closed-tree`{.variable}.

        9.  If `parent`{.variable} is non-null, then set
            `parent`{.variable} to the result of invoking
            `parent`{.variable}'s [get the
            parent](#get-the-parent){#ref-for-get-the-parent④
            link-type="dfn"} with `event`{.variable}.

        10. Set `slot-in-closed-tree`{.variable} to false.

    10. Let `clearTargetsStruct`{.variable} be the last struct in
        `event`{.variable}'s [path](#event-path){#ref-for-event-path④
        link-type="dfn"} whose [shadow-adjusted
        target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target
        link-type="dfn"} is non-null.

    11. If `clearTargetsStruct`{.variable}'s [shadow-adjusted
        target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target①
        link-type="dfn"}, `clearTargetsStruct`{.variable}'s
        [relatedTarget](#event-path-relatedtarget){#ref-for-event-path-relatedtarget
        link-type="dfn"}, or an
        [`EventTarget`{.idl}](#eventtarget){#ref-for-eventtarget②③
        link-type="idl"} object in `clearTargetsStruct`{.variable}'s
        [touch target
        list](#event-path-touch-target-list){#ref-for-event-path-touch-target-list
        link-type="dfn"} is a
        [node](#concept-node){#ref-for-concept-node⑦ link-type="dfn"}
        whose [root](#concept-tree-root){#ref-for-concept-tree-root⑥
        link-type="dfn"} is a [shadow
        root](#concept-shadow-root){#ref-for-concept-shadow-root③
        link-type="dfn"}: set `clearTargets`{.variable} to true.

    12. If `activationTarget`{.variable} is non-null and
        `activationTarget`{.variable} has [legacy-pre-activation
        behavior](#eventtarget-legacy-pre-activation-behavior){#ref-for-eventtarget-legacy-pre-activation-behavior①
        link-type="dfn"}, then run `activationTarget`{.variable}'s
        [legacy-pre-activation
        behavior](#eventtarget-legacy-pre-activation-behavior){#ref-for-eventtarget-legacy-pre-activation-behavior②
        link-type="dfn"}.

    13. [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑤
        link-type="dfn"} `struct`{.variable} of `event`{.variable}'s
        [path](#event-path){#ref-for-event-path⑤ link-type="dfn"}, in
        reverse order:

        1.  If `struct`{.variable}'s [shadow-adjusted
            target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target②
            link-type="dfn"} is non-null, then set `event`{.variable}'s
            [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑦
            link-type="idl"} attribute to
            [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target④
            link-type="idl"}.

        2.  Otherwise, set `event`{.variable}'s
            [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑧
            link-type="idl"} attribute to
            [`CAPTURING_PHASE`{.idl}](#dom-event-capturing_phase){#ref-for-dom-event-capturing_phase④
            link-type="idl"}.

        3.  [Invoke](#concept-event-listener-invoke){#ref-for-concept-event-listener-invoke
            link-type="dfn"} with `struct`{.variable},
            `event`{.variable}, \"`capturing`\", and
            `legacyOutputDidListenersThrowFlag`{.variable} if given.

    14. [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑥
        link-type="dfn"} `struct`{.variable} of `event`{.variable}'s
        [path](#event-path){#ref-for-event-path⑥ link-type="dfn"}:

        1.  If `struct`{.variable}'s [shadow-adjusted
            target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target③
            link-type="dfn"} is non-null, then set `event`{.variable}'s
            [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase⑨
            link-type="idl"} attribute to
            [`AT_TARGET`{.idl}](#dom-event-at_target){#ref-for-dom-event-at_target⑤
            link-type="idl"}.

        2.  Otherwise:

            1.  If `event`{.variable}'s
                [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles⑥
                link-type="idl"} attribute is false, then
                [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue
                link-type="dfn"}.

            2.  Set `event`{.variable}'s
                [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase①⓪
                link-type="idl"} attribute to
                [`BUBBLING_PHASE`{.idl}](#dom-event-bubbling_phase){#ref-for-dom-event-bubbling_phase④
                link-type="idl"}.

        3.  [Invoke](#concept-event-listener-invoke){#ref-for-concept-event-listener-invoke①
            link-type="dfn"} with `struct`{.variable},
            `event`{.variable}, \"`bubbling`\", and
            `legacyOutputDidListenersThrowFlag`{.variable} if given.

7.  Set `event`{.variable}'s
    [`eventPhase`{.idl}](#dom-event-eventphase){#ref-for-dom-event-eventphase①①
    link-type="idl"} attribute to
    [`NONE`{.idl}](#dom-event-none){#ref-for-dom-event-none③
    link-type="idl"}.

8.  Set `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑤
    link-type="idl"} attribute to null.

9.  Set `event`{.variable}'s [path](#event-path){#ref-for-event-path⑦
    link-type="dfn"} to the empty list.

10. Unset `event`{.variable}'s [dispatch
    flag](#dispatch-flag){#ref-for-dispatch-flag④ link-type="dfn"},
    [stop propagation
    flag](#stop-propagation-flag){#ref-for-stop-propagation-flag⑤
    link-type="dfn"}, and [stop immediate propagation
    flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag②
    link-type="dfn"}.

11. If `clearTargets`{.variable} is true:

    1.  Set `event`{.variable}'s
        [target](#event-target){#ref-for-event-target⑨ link-type="dfn"}
        to null.

    2.  Set `event`{.variable}'s
        [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget④
        link-type="dfn"} to null.

    3.  Set `event`{.variable}'s [touch target
        list](#event-touch-target-list){#ref-for-event-touch-target-list③
        link-type="dfn"} to the empty list.

12. If `activationTarget`{.variable} is non-null:

    1.  If `event`{.variable}'s [canceled
        flag](#canceled-flag){#ref-for-canceled-flag④ link-type="dfn"}
        is unset, then run `activationTarget`{.variable}'s [activation
        behavior](#eventtarget-activation-behavior){#ref-for-eventtarget-activation-behavior⑥
        link-type="dfn"} with `event`{.variable}.

    2.  Otherwise, if `activationTarget`{.variable} has
        [legacy-canceled-activation
        behavior](#eventtarget-legacy-canceled-activation-behavior){#ref-for-eventtarget-legacy-canceled-activation-behavior①
        link-type="dfn"}, then run `activationTarget`{.variable}'s
        [legacy-canceled-activation
        behavior](#eventtarget-legacy-canceled-activation-behavior){#ref-for-eventtarget-legacy-canceled-activation-behavior②
        link-type="dfn"}.

13. Return false if `event`{.variable}'s [canceled
    flag](#canceled-flag){#ref-for-canceled-flag⑤ link-type="dfn"} is
    set; otherwise true.
:::

::: {.algorithm algorithm="append to an event path"}
To [append to an event path]{#concept-event-path-append .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given an `event`{.variable},
`invocationTarget`{.variable}, `shadowAdjustedTarget`{.variable},
`relatedTarget`{.variable}, `touchTargets`{.variable}, and a
`slot-in-closed-tree`{.variable}, run these steps:

1.  Let `invocationTargetInShadowTree`{.variable} be false.

2.  If `invocationTarget`{.variable} is a
    [node](#concept-node){#ref-for-concept-node⑧ link-type="dfn"} and
    its [root](#concept-tree-root){#ref-for-concept-tree-root⑦
    link-type="dfn"} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root④
    link-type="dfn"}, then set `invocationTargetInShadowTree`{.variable}
    to true.

3.  Let `root-of-closed-tree`{.variable} be false.

4.  If `invocationTarget`{.variable} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root⑤
    link-type="dfn"} whose
    [mode](#shadowroot-mode){#ref-for-shadowroot-mode② link-type="dfn"}
    is \"`closed`\", then set `root-of-closed-tree`{.variable} to true.

5.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑥
    link-type="dfn"} a new
    [struct](https://infra.spec.whatwg.org/#struct){#ref-for-struct②
    link-type="dfn"} to `event`{.variable}'s
    [path](#event-path){#ref-for-event-path⑧ link-type="dfn"} whose
    [invocation
    target](#event-path-invocation-target){#ref-for-event-path-invocation-target④
    link-type="dfn"} is `invocationTarget`{.variable},
    [invocation-target-in-shadow-tree](#event-path-invocation-target-in-shadow-tree){#ref-for-event-path-invocation-target-in-shadow-tree
    link-type="dfn"} is `invocationTargetInShadowTree`{.variable},
    [shadow-adjusted
    target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target④
    link-type="dfn"} is `shadowAdjustedTarget`{.variable},
    [relatedTarget](#event-path-relatedtarget){#ref-for-event-path-relatedtarget①
    link-type="dfn"} is `relatedTarget`{.variable}, [touch target
    list](#event-path-touch-target-list){#ref-for-event-path-touch-target-list①
    link-type="dfn"} is `touchTargets`{.variable},
    [root-of-closed-tree](#event-path-root-of-closed-tree){#ref-for-event-path-root-of-closed-tree③
    link-type="dfn"} is `root-of-closed-tree`{.variable}, and
    [slot-in-closed-tree](#event-path-slot-in-closed-tree){#ref-for-event-path-slot-in-closed-tree③
    link-type="dfn"} is `slot-in-closed-tree`{.variable}.
:::

::: {.algorithm algorithm="invoke"}
To [invoke]{#concept-event-listener-invoke .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a `struct`{.variable},
`event`{.variable}, `phase`{.variable}, and an optional
`legacyOutputDidListenersThrowFlag`{.variable}, run these steps:

1.  Set `event`{.variable}'s
    [target](#event-target){#ref-for-event-target①⓪ link-type="dfn"} to
    the [shadow-adjusted
    target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target⑤
    link-type="dfn"} of the last struct in `event`{.variable}'s
    [path](#event-path){#ref-for-event-path⑨ link-type="dfn"}, that is
    either `struct`{.variable} or preceding `struct`{.variable}, whose
    [shadow-adjusted
    target](#event-path-shadow-adjusted-target){#ref-for-event-path-shadow-adjusted-target⑥
    link-type="dfn"} is non-null.

2.  Set `event`{.variable}'s
    [relatedTarget](#event-relatedtarget){#ref-for-event-relatedtarget⑤
    link-type="dfn"} to `struct`{.variable}'s
    [relatedTarget](#event-path-relatedtarget){#ref-for-event-path-relatedtarget②
    link-type="dfn"}.

3.  Set `event`{.variable}'s [touch target
    list](#event-touch-target-list){#ref-for-event-touch-target-list④
    link-type="dfn"} to `struct`{.variable}'s [touch target
    list](#event-path-touch-target-list){#ref-for-event-path-touch-target-list②
    link-type="dfn"}.

4.  If `event`{.variable}'s [stop propagation
    flag](#stop-propagation-flag){#ref-for-stop-propagation-flag⑥
    link-type="dfn"} is set, then return.

5.  Initialize `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑥
    link-type="idl"} attribute to `struct`{.variable}'s [invocation
    target](#event-path-invocation-target){#ref-for-event-path-invocation-target⑤
    link-type="dfn"}.

6.  Let `listeners`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone
    link-type="dfn"} of `event`{.variable}'s
    [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑦
    link-type="idl"} attribute value's [event listener
    list](#eventtarget-event-listener-list){#ref-for-eventtarget-event-listener-list⑧
    link-type="dfn"}.

    This avoids [event
    listeners](#concept-event-listener){#ref-for-concept-event-listener②⑥
    link-type="dfn"} added after this point from being run. Note that
    removal still has an effect due to the
    [removed](#event-listener-removed){#ref-for-event-listener-removed①
    link-type="dfn"} field.

7.  Let `invocationTargetInShadowTree`{.variable} be
    `struct`{.variable}'s
    [invocation-target-in-shadow-tree](#event-path-invocation-target-in-shadow-tree){#ref-for-event-path-invocation-target-in-shadow-tree①
    link-type="dfn"}.

8.  Let `found`{.variable} be the result of running [inner
    invoke](#concept-event-listener-inner-invoke){#ref-for-concept-event-listener-inner-invoke
    link-type="dfn"} with `event`{.variable}, `listeners`{.variable},
    `phase`{.variable}, `invocationTargetInShadowTree`{.variable}, and
    `legacyOutputDidListenersThrowFlag`{.variable} if given.

9.  If `found`{.variable} is false and `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑦
    link-type="idl"} attribute is true:

    1.  Let `originalEventType`{.variable} be `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①①
        link-type="idl"} attribute value.

    2.  If `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①②
        link-type="idl"} attribute value is a match for any of the
        strings in the first column in the following table, set
        `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①③
        link-type="idl"} attribute value to the string in the second
        column on the same row as the matching string, and return
        otherwise.

        Event type

        Legacy event type

        \"`animationend`\"

        \"`webkitAnimationEnd`\"

        \"`animationiteration`\"

        \"`webkitAnimationIteration`\"

        \"`animationstart`\"

        \"`webkitAnimationStart`\"

        \"`transitionend`\"

        \"`webkitTransitionEnd`\"

    3.  [Inner
        invoke](#concept-event-listener-inner-invoke){#ref-for-concept-event-listener-inner-invoke①
        link-type="dfn"} with `event`{.variable},
        `listeners`{.variable}, `phase`{.variable},
        `invocationTargetInShadowTree`{.variable}, and
        `legacyOutputDidListenersThrowFlag`{.variable} if given.

    4.  Set `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①④
        link-type="idl"} attribute value to
        `originalEventType`{.variable}.
:::

::: {.algorithm algorithm="inner invoke"}
To [inner invoke]{#concept-event-listener-inner-invoke .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given an `event`{.variable},
`listeners`{.variable}, `phase`{.variable},
`invocationTargetInShadowTree`{.variable}, and an optional
`legacyOutputDidListenersThrowFlag`{.variable}, run these steps:

1.  Let `found`{.variable} be false.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑦
    link-type="dfn"} `listener`{.variable} of `listeners`{.variable},
    whose
    [removed](#event-listener-removed){#ref-for-event-listener-removed②
    link-type="dfn"} is false:

    1.  If `event`{.variable}'s
        [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⑤
        link-type="idl"} attribute value is not `listener`{.variable}'s
        [type](#event-listener-type){#ref-for-event-listener-type⑨
        link-type="dfn"}, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue①
        link-type="dfn"}.

    2.  Set `found`{.variable} to true.

    3.  If `phase`{.variable} is \"`capturing`\" and
        `listener`{.variable}'s
        [capture](#event-listener-capture){#ref-for-event-listener-capture⑦
        link-type="dfn"} is false, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue②
        link-type="dfn"}.

    4.  If `phase`{.variable} is \"`bubbling`\" and
        `listener`{.variable}'s
        [capture](#event-listener-capture){#ref-for-event-listener-capture⑧
        link-type="dfn"} is true, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue③
        link-type="dfn"}.

    5.  If `listener`{.variable}'s
        [once](#event-listener-once){#ref-for-event-listener-once①
        link-type="dfn"} is true, then [remove an event
        listener](#remove-an-event-listener){#ref-for-remove-an-event-listener③
        link-type="dfn"} given `event`{.variable}'s
        [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑧
        link-type="idl"} attribute value and `listener`{.variable}.

    6.  Let `global`{.variable} be `listener`{.variable}
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑧
        link-type="dfn"}'s [associated
        realm](https://webidl.spec.whatwg.org/#dfn-associated-realm){#ref-for-dfn-associated-realm
        link-type="dfn"}'s [global
        object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-realm-global){#ref-for-concept-realm-global
        link-type="dfn"}.

    7.  Let `currentEvent`{.variable} be undefined.

    8.  If `global`{.variable} is a
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑥
        link-type="idl"} object:

        1.  Set `currentEvent`{.variable} to `global`{.variable}'s
            [current
            event](#window-current-event){#ref-for-window-current-event①
            link-type="dfn"}.

        2.  If `invocationTargetInShadowTree`{.variable} is false, then
            set `global`{.variable}'s [current
            event](#window-current-event){#ref-for-window-current-event②
            link-type="dfn"} to `event`{.variable}.

    9.  If `listener`{.variable}'s
        [passive](#event-listener-passive){#ref-for-event-listener-passive②
        link-type="dfn"} is true, then set `event`{.variable}'s [in
        passive listener
        flag](#in-passive-listener-flag){#ref-for-in-passive-listener-flag①
        link-type="dfn"}.

    10. If `global`{.variable} is a
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑦
        link-type="idl"} object, then [record timing info for event
        listener](https://w3c.github.io/long-animation-frames/#record-timing-info-for-event-listener){#ref-for-record-timing-info-for-event-listener
        link-type="dfn"} given `event`{.variable} and
        `listener`{.variable}.

    11. [Call a user object's
        operation](https://webidl.spec.whatwg.org/#call-a-user-objects-operation){#ref-for-call-a-user-objects-operation
        link-type="dfn"} with `listener`{.variable}'s
        [callback](#event-listener-callback){#ref-for-event-listener-callback①⑨
        link-type="dfn"}, \"`handleEvent`\", « `event`{.variable} », and
        `event`{.variable}'s
        [`currentTarget`{.idl}](#dom-event-currenttarget){#ref-for-dom-event-currenttarget⑨
        link-type="idl"} attribute value. If this throws an exception
        `exception`{.variable}:

        1.  [Report](https://html.spec.whatwg.org/multipage/webappapis.html#report-an-exception){#ref-for-report-an-exception
            link-type="dfn"} `exception`{.variable} for
            `listener`{.variable}'s
            [callback](#event-listener-callback){#ref-for-event-listener-callback②⓪
            link-type="dfn"}'s corresponding JavaScript object's
            [associated
            realm](https://webidl.spec.whatwg.org/#dfn-associated-realm){#ref-for-dfn-associated-realm①
            link-type="dfn"}'s [global
            object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-realm-global){#ref-for-concept-realm-global①
            link-type="dfn"}.

        2.  Set `legacyOutputDidListenersThrowFlag`{.variable} if given.

            The `legacyOutputDidListenersThrowFlag`{.variable} is only
            used by Indexed Database API.
            [\[INDEXEDDB\]](#biblio-indexeddb "Indexed Database API"){link-type="biblio"}

    12. Unset `event`{.variable}'s [in passive listener
        flag](#in-passive-listener-flag){#ref-for-in-passive-listener-flag②
        link-type="dfn"}.

    13. If `global`{.variable} is a
        [`Window`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#window){#ref-for-window⑧
        link-type="idl"} object, then set `global`{.variable}'s [current
        event](#window-current-event){#ref-for-window-current-event③
        link-type="dfn"} to `currentEvent`{.variable}.

    14. If `event`{.variable}'s [stop immediate propagation
        flag](#stop-immediate-propagation-flag){#ref-for-stop-immediate-propagation-flag③
        link-type="dfn"} is set, then
        [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break①
        link-type="dfn"}.

3.  Return `found`{.variable}.
:::

### [2.10. ]{.secno}[Firing events]{.content}[](#firing-events){.self-link} {#firing-events .heading .settled level="2.10"}

::: {.algorithm algorithm="fire an event"}
To [fire an event]{#concept-event-fire .dfn .dfn-paneled dfn-type="dfn"
export=""} named `e`{.variable} at `target`{.variable}, optionally using
an `eventConstructor`{.variable}, with a description of how IDL
attributes are to be initialized, and a
`legacy target override flag`{.variable}, run these steps:

1.  If `eventConstructor`{.variable} is not given, then let
    `eventConstructor`{.variable} be
    [`Event`{.idl}](#event){#ref-for-event①⑥ link-type="idl"}.

2.  Let `event`{.variable} be the result of [creating an
    event](#concept-event-create){#ref-for-concept-event-create②
    link-type="dfn"} given `eventConstructor`{.variable}, in the
    [relevant
    realm](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-realm){#ref-for-concept-relevant-realm
    link-type="dfn"} of `target`{.variable}.

3.  Initialize `event`{.variable}'s
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⑥
    link-type="idl"} attribute to `e`{.variable}.

4.  Initialize any other IDL attributes of `event`{.variable} as
    described in the invocation of this algorithm.

    This also allows for the
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑧
    link-type="idl"} attribute to be set to false.

5.  Return the result of
    [dispatching](#concept-event-dispatch){#ref-for-concept-event-dispatch②⑤
    link-type="dfn"} `event`{.variable} at `target`{.variable}, with
    `legacy target override flag`{.variable} set if set.
:::

Fire in the context of DOM is short for
[creating](#concept-event-create){#ref-for-concept-event-create③
link-type="dfn"}, initializing, and
[dispatching](#concept-event-dispatch){#ref-for-concept-event-dispatch②⑥
link-type="dfn"} an [event](#concept-event){#ref-for-concept-event④⑧
link-type="dfn"}. [Fire an
event](#concept-event-fire){#ref-for-concept-event-fire①
link-type="dfn"} makes that process easier to write down.

::: {#firing-events-example .example}
[](#firing-events-example){.self-link}

If the [event](#concept-event){#ref-for-concept-event④⑨ link-type="dfn"}
needs its
[`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles⑦
link-type="idl"} or
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable⑨
link-type="idl"} attribute initialized, one could write \"[fire an
event](#concept-event-fire){#ref-for-concept-event-fire②
link-type="dfn"} named `submit` at `target`{.variable} with its
[`cancelable`{.idl}](#dom-event-cancelable){#ref-for-dom-event-cancelable①⓪
link-type="idl"} attribute initialized to true\".

Or, when a custom constructor is needed, \"[fire an
event](#concept-event-fire){#ref-for-concept-event-fire③
link-type="dfn"} named `click` at `target`{.variable} using
[`MouseEvent`{.idl}](https://w3c.github.io/uievents/#mouseevent){#ref-for-mouseevent②
link-type="idl"} with its
[`detail`{.idl}](https://w3c.github.io/uievents/#dom-uievent-detail){#ref-for-dom-uievent-detail
link-type="idl"} attribute initialized to 1\".

Occasionally the return value is important:

1.  Let `doAction`{.variable} be the result of [firing an
    event](#concept-event-fire){#ref-for-concept-event-fire④
    link-type="dfn"} named `like` at `target`{.variable}.

2.  If `doAction`{.variable} is true, then ...
:::

### [2.11. ]{.secno}[]{#action-versus-occurance .bs-old-id}[Action versus occurrence]{.content}[](#action-versus-occurrence){.self-link} {#action-versus-occurrence .heading .settled level="2.11"}

An [event](#concept-event){#ref-for-concept-event⑤⓪ link-type="dfn"}
signifies an occurrence, not an action. Phrased differently, it
represents a notification from an algorithm and can be used to influence
the future course of that algorithm (e.g., through invoking
[`preventDefault()`{.idl}](#dom-event-preventdefault){#ref-for-dom-event-preventdefault⑨
link-type="idl"}). [Events](#concept-event){#ref-for-concept-event⑤①
link-type="dfn"} must not be used as actions or initiators that cause
some algorithm to start running. That is not what they are for.

This is called out here specifically because previous iterations of the
DOM had a concept of \"default actions\" associated with
[events](#concept-event){#ref-for-concept-event⑤② link-type="dfn"} that
gave folks all the wrong ideas.
[Events](#concept-event){#ref-for-concept-event⑤③ link-type="dfn"} do
not represent or cause actions, they can only be used to influence an
ongoing one.

## [3. ]{.secno}[Aborting ongoing activities]{.content}[](#aborting-ongoing-activities){.self-link} {#aborting-ongoing-activities .heading .settled level="3"}

Though promises do not have a built-in aborting mechanism, many APIs
using them require abort semantics.
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller
link-type="idl"} is meant to support these requirements by providing an
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort①
link-type="idl"} method that toggles the state of a corresponding
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④
link-type="idl"} object. The API which wishes to support aborting can
accept an [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑤
link-type="idl"} object, and use its state to determine how to proceed.

APIs that rely upon
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller①
link-type="idl"} are encouraged to respond to
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort②
link-type="idl"} by rejecting any unsettled promise with the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑥
link-type="idl"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason
link-type="dfn"}.

::: {#aborting-ongoing-activities-example .example}
[](#aborting-ongoing-activities-example){.self-link}

A hypothetical `doAmazingness({ ... })` method could accept an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑦
link-type="idl"} object to support aborting as follows:

``` {.lang-javascript .highlight}
const controller = new AbortController();
const signal = controller.signal;

startSpinner();

doAmazingness({ ..., signal })
  .then(result => ...)
  .catch(err => {
    if (err.name == 'AbortError') return;
    showUserErrorMessage();
  })
  .then(() => stopSpinner());

// …

controller.abort();
```

`doAmazingness` could be implemented as follows:

``` {.lang-javascript .highlight}
function doAmazingness({signal}) {
  return new Promise((resolve, reject) => {
    signal.throwIfAborted();

    // Begin doing amazingness, and call resolve(result) when done.
    // But also, watch for signals:
    signal.addEventListener('abort', () => {
      // Stop doing amazingness, and:
      reject(signal.reason);
    });
  });
}
```
:::

APIs that do not return promises can either react in an equivalent
manner or opt to not surface the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑧
link-type="idl"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①
link-type="dfn"} at all.
[`addEventListener()`{.idl}](#dom-eventtarget-addeventlistener){#ref-for-dom-eventtarget-addeventlistener④
link-type="idl"} is an example of an API where the latter made sense.

APIs that require more granular control could extend both
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller②
link-type="idl"} and
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑨
link-type="idl"} objects according to their needs.

### [3.1. ]{.secno}[Interface [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller③ link-type="idl"}]{.content}[](#interface-abortcontroller){.self-link} {#interface-abortcontroller .heading .settled level="3.1"}

``` {.idl .highlight .def}
[Exposed=*]
interface AbortController {
  constructor();

  [SameObject] readonly attribute AbortSignal signal;

  undefined abort(optional any reason);
};
```

`controller`{.variable}` = new `[`AbortController`](#dom-abortcontroller-abortcontroller){#ref-for-dom-abortcontroller-abortcontroller① .idl-code link-type="constructor"}`()`
:   Returns a new `controller`{.variable} whose
    [`signal`{.idl}](#dom-abortcontroller-signal){#ref-for-dom-abortcontroller-signal①
    link-type="idl"} is set to a newly created
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①①
    link-type="idl"} object.

`controller`{.variable}` . `[`signal`](#dom-abortcontroller-signal){#ref-for-dom-abortcontroller-signal② .idl-code link-type="attribute"}
:   Returns the
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①②
    link-type="idl"} object associated with this object.

`controller`{.variable}` . `[`abort`](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort④ .idl-code link-type="method"}`(``reason`{.variable}`)`
:   Invoking this method will store `reason`{.variable} in this object's
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①③
    link-type="idl"}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason②
    link-type="dfn"}, and signal to any observers that the associated
    activity is to be aborted. If `reason`{.variable} is undefined, then
    an
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦
    link-type="idl"} will be stored.

An [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller④
link-type="idl"} object has an associated
[signal]{#abortcontroller-signal .dfn .dfn-paneled
dfn-for="AbortController" dfn-type="dfn" export=""} (an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①④
link-type="idl"} object).

::: {.algorithm algorithm="AbortController()" algorithm-for="AbortController"}
The [`new AbortController()`]{#dom-abortcontroller-abortcontroller .dfn
.dfn-paneled .idl-code dfn-for="AbortController" dfn-type="constructor"
export="" lt="AbortController()|constructor()"} constructor steps are:

1.  Let `signal`{.variable} be a new
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①⑤
    link-type="idl"} object.

2.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④
    link-type="dfn"}'s
    [signal](#abortcontroller-signal){#ref-for-abortcontroller-signal
    link-type="dfn"} to `signal`{.variable}.
:::

The [`signal`]{#dom-abortcontroller-signal .dfn .dfn-paneled .idl-code
dfn-for="AbortController" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤
link-type="dfn"}'s
[signal](#abortcontroller-signal){#ref-for-abortcontroller-signal①
link-type="dfn"}.

::: {.algorithm algorithm="abort(reason)" algorithm-for="AbortController"}
The [`abort(``reason`{.variable}`)`]{#dom-abortcontroller-abort .dfn
.dfn-paneled .idl-code dfn-for="AbortController" dfn-type="method"
export="" lt="abort(reason)|abort()"} method steps are to [signal
abort](#abortcontroller-signal-abort){#ref-for-abortcontroller-signal-abort
link-type="dfn"} on
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥
link-type="dfn"} with `reason`{.variable} if it is given.
:::

::: {.algorithm algorithm="signal abort" algorithm-for="AbortController"}
To [signal abort]{#abortcontroller-signal-abort .dfn .dfn-paneled
dfn-for="AbortController" dfn-type="dfn" export=""} on an
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑤
link-type="idl"} `controller`{.variable} with an optional
`reason`{.variable}, [signal
abort](#abortsignal-signal-abort){#ref-for-abortsignal-signal-abort
link-type="dfn"} on `controller`{.variable}'s
[signal](#abortcontroller-signal){#ref-for-abortcontroller-signal②
link-type="dfn"} with `reason`{.variable} if it is given.
:::

### [3.2. ]{.secno}[Interface [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal①⑥ link-type="idl"}]{.content}[](#interface-AbortSignal){.self-link} {#interface-AbortSignal .heading .settled level="3.2"}

``` {.idl .highlight .def}
[Exposed=*]
interface AbortSignal : EventTarget {
  [NewObject] static AbortSignal abort(optional any reason);
  [Exposed=(Window,Worker), NewObject] static AbortSignal timeout([EnforceRange] unsigned long long milliseconds);
  [NewObject] static AbortSignal _any(sequence<AbortSignal> signals);

  readonly attribute boolean aborted;
  readonly attribute any reason;
  undefined throwIfAborted();

  attribute EventHandler onabort;
};
```

`AbortSignal . `[`abort`](#dom-abortsignal-abort){#ref-for-dom-abortsignal-abort① .idl-code link-type="method"}`(``reason`{.variable}`)`
:   Returns an
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②①
    link-type="idl"} instance whose [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason③
    link-type="dfn"} is set to `reason`{.variable} if not undefined;
    otherwise to an
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧
    link-type="idl"}.

`AbortSignal . `[`any`](#dom-abortsignal-any){#ref-for-dom-abortsignal-any① .idl-code link-type="method"}`(``signals`{.variable}`)`
:   Returns an
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②②
    link-type="idl"} instance which will be aborted once any of
    `signals`{.variable} is aborted. Its [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason④
    link-type="dfn"} will be set to whichever one of
    `signals`{.variable} caused it to be aborted.

`AbortSignal . `[`timeout`](#dom-abortsignal-timeout){#ref-for-dom-abortsignal-timeout① .idl-code link-type="method"}`(``milliseconds`{.variable}`)`
:   Returns an
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②③
    link-type="idl"} instance which will be aborted in
    `milliseconds`{.variable} milliseconds. Its [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑤
    link-type="dfn"} will be set to a
    \"[`TimeoutError`{.idl}](https://webidl.spec.whatwg.org/#timeouterror){#ref-for-timeouterror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨
    link-type="idl"}.

`signal`{.variable}` . `[`aborted`](#dom-abortsignal-aborted){#ref-for-dom-abortsignal-aborted① .idl-code link-type="attribute"}
:   Returns true if `signal`{.variable}'s
    [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑥
    link-type="idl"} has signaled to abort; otherwise false.

`signal`{.variable}` . `[`reason`](#dom-abortsignal-reason){#ref-for-dom-abortsignal-reason① .idl-code link-type="attribute"}
:   Returns `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑥
    link-type="dfn"}.

`signal`{.variable}` . `[`throwIfAborted`](#dom-abortsignal-throwifaborted){#ref-for-dom-abortsignal-throwifaborted① .idl-code link-type="method"}`()`
:   Throws `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑦
    link-type="dfn"}, if `signal`{.variable}'s
    [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑦
    link-type="idl"} has signaled to abort; otherwise, does nothing.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②④
link-type="idl"} object has an associated [abort
reason]{#abortsignal-abort-reason .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" export=""} (a JavaScript value),
which is initially undefined.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑤
link-type="idl"} object has associated [abort
algorithms]{#abortsignal-abort-algorithms .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""}, (a
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set②
link-type="dfn"} of algorithms which are to be executed when it is
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①
link-type="dfn"}), which is initially empty.

The [abort
algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms
link-type="dfn"} enable APIs with complex requirements to react in a
reasonable way to
[`abort()`{.idl}](#dom-abortcontroller-abort){#ref-for-dom-abortcontroller-abort⑤
link-type="idl"}. For example, a given API's [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑧
link-type="dfn"} might need to be propagated to a cross-thread
environment, such as a service worker.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑥
link-type="idl"} object has a [dependent]{#abortsignal-dependent .dfn
.dfn-paneled dfn-for="AbortSignal" dfn-type="dfn" noexport=""} (a
boolean), which is initially false.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑦
link-type="idl"} object has associated [source
signals]{#abortsignal-source-signals .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""} (a weak
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set③
link-type="dfn"} of
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑧
link-type="idl"} objects that the object is dependent on for its
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted②
link-type="dfn"} state), which is initially empty.

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal②⑨
link-type="idl"} object has associated [dependent
signals]{#abortsignal-dependent-signals .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""} (a weak
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set④
link-type="dfn"} of
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⓪
link-type="idl"} objects that are dependent on the object for their
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted③
link-type="dfn"} state), which is initially empty.

------------------------------------------------------------------------

::: {.algorithm algorithm="abort(reason)" algorithm-for="AbortSignal"}
The static [`abort(``reason`{.variable}`)`]{#dom-abortsignal-abort .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method" export=""
lt="abort(reason)|abort()"} method steps are:

1.  Let `signal`{.variable} be a new
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③①
    link-type="idl"} object.

2.  Set `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason⑨
    link-type="dfn"} to `reason`{.variable} if it is given; otherwise to
    a new
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪
    link-type="idl"}.

3.  Return `signal`{.variable}.
:::

::: {.algorithm algorithm="timeout(milliseconds)" algorithm-for="AbortSignal"}
The static
[`timeout(``milliseconds`{.variable}`)`]{#dom-abortsignal-timeout .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method"
export=""} method steps are:

1.  Let `signal`{.variable} be a new
    [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③②
    link-type="idl"} object.

2.  Let `global`{.variable} be `signal`{.variable}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global①
    link-type="dfn"}.

3.  [Run steps after a
    timeout](https://html.spec.whatwg.org/multipage/timers-and-user-prompts.html#run-steps-after-a-timeout){#ref-for-run-steps-after-a-timeout
    link-type="dfn"} given `global`{.variable},
    \"`AbortSignal-timeout`\", `milliseconds`{.variable}, and the
    following step:

    1.  [Queue a global
        task](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-global-task){#ref-for-queue-a-global-task
        link-type="dfn"} on the [timer task
        source](https://html.spec.whatwg.org/multipage/timers-and-user-prompts.html#timer-task-source){#ref-for-timer-task-source
        link-type="dfn"} given `global`{.variable} to [signal
        abort](#abortsignal-signal-abort){#ref-for-abortsignal-signal-abort①
        link-type="dfn"} given `signal`{.variable} and a new
        \"[`TimeoutError`{.idl}](https://webidl.spec.whatwg.org/#timeouterror){#ref-for-timeouterror①
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①①
        link-type="idl"}.

    For the duration of this timeout, if `signal`{.variable} has any
    event listeners registered for its
    [`abort`{.idl}](#eventdef-abortsignal-abort){#ref-for-eventdef-abortsignal-abort
    link-type="idl"} event, there must be a strong reference from
    `global`{.variable} to `signal`{.variable}.

4.  Return `signal`{.variable}.
:::

::: {.algorithm algorithm="any(signals)" algorithm-for="AbortSignal"}
The static [`any(``signals`{.variable}`)`]{#dom-abortsignal-any .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method"
export=""} method steps are to return the result of [creating a
dependent abort
signal](#create-a-dependent-abort-signal){#ref-for-create-a-dependent-abort-signal
link-type="dfn"} from `signals`{.variable} using
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③③
link-type="idl"} and the [current
realm](https://tc39.es/ecma262/#current-realm){#ref-for-current-realm
link-type="dfn"}.
:::

The [`aborted`]{#dom-abortsignal-aborted .dfn .dfn-paneled .idl-code
dfn-for="AbortSignal" dfn-type="attribute" export=""} getter steps are
to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦
link-type="dfn"} is
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted④
link-type="dfn"}; otherwise false.

The [`reason`]{#dom-abortsignal-reason .dfn .dfn-paneled .idl-code
dfn-for="AbortSignal" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧
link-type="dfn"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⓪
link-type="dfn"}.

The [`throwIfAborted()`]{#dom-abortsignal-throwifaborted .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="method"
export=""} method steps are to throw
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨
link-type="dfn"}'s [abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①①
link-type="dfn"}, if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪
link-type="dfn"} is
[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑤
link-type="dfn"}.

::: {#example-throwifaborted .example}
[](#example-throwifaborted){.self-link}

This method is primarily useful for when functions accepting
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③④
link-type="idl"}s want to throw (or return a rejected promise) at
specific checkpoints, instead of passing along the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑤
link-type="idl"} to other methods. For example, the following function
allows aborting in between each attempt to poll for a condition. This
gives opportunities to abort the polling process, even though the actual
asynchronous operation (i.e., `await func()`{.lang-javascript
.highlight}) does not accept an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑥
link-type="idl"}.

``` {.lang-javascript .highlight}
async function waitForCondition(func, targetValue, { signal } = {}) {
  while (true) {
    signal?.throwIfAborted();

    const result = await func();
    if (result === targetValue) {
      return;
    }
  }
}
```
:::

The [`onabort`]{#dom-abortsignal-onabort .dfn .dfn-paneled .idl-code
dfn-for="AbortSignal" dfn-type="attribute" export=""} attribute is an
[event handler IDL
attribute](https://html.spec.whatwg.org/multipage/webappapis.html#event-handler-idl-attributes){#ref-for-event-handler-idl-attributes
link-type="dfn"} for the [`onabort`]{#abortsignal-onabort .dfn
.dfn-paneled dfn-for="AbortSignal" dfn-type="dfn" export=""} [event
handler](https://html.spec.whatwg.org/multipage/webappapis.html#event-handlers){#ref-for-event-handlers①
link-type="dfn"}, whose [event handler event
type](https://html.spec.whatwg.org/multipage/webappapis.html#event-handler-event-type){#ref-for-event-handler-event-type
link-type="dfn"} is [`abort`]{#eventdef-abortsignal-abort .dfn
.dfn-paneled .idl-code dfn-for="AbortSignal" dfn-type="event"
export=""}.

Changes to an [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑦
link-type="idl"} object represent the wishes of the corresponding
[`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑧
link-type="idl"} object, but an API observing the
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑧
link-type="idl"} object can choose to ignore them. For instance, if the
operation has already completed.

------------------------------------------------------------------------

An [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal③⑨
link-type="idl"} object is [aborted]{#abortsignal-aborted .dfn
.dfn-paneled dfn-for="AbortSignal" dfn-type="dfn" export=""} when its
[abort
reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①②
link-type="dfn"} is not undefined.

::: {.algorithm algorithm="add" algorithm-for="AbortSignal"}
To [add]{#abortsignal-add .dfn .dfn-paneled dfn-for="AbortSignal"
dfn-type="dfn" export=""} an algorithm `algorithm`{.variable} to an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⓪
link-type="idl"} object `signal`{.variable}:

1.  If `signal`{.variable} is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑥
    link-type="dfn"}, then return.

2.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append①
    link-type="dfn"} `algorithm`{.variable} to `signal`{.variable}'s
    [abort
    algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms①
    link-type="dfn"}.
:::

::: {.algorithm algorithm="remove" algorithm-for="AbortSignal"}
To [remove]{#abortsignal-remove .dfn .dfn-paneled dfn-for="AbortSignal"
dfn-type="dfn" export=""} an algorithm `algorithm`{.variable} from an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④①
link-type="idl"} `signal`{.variable},
[remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove①
link-type="dfn"} `algorithm`{.variable} from `signal`{.variable}'s
[abort
algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms②
link-type="dfn"}.
:::

::: {.algorithm algorithm="signal abort" algorithm-for="AbortSignal"}
To [signal abort]{#abortsignal-signal-abort .dfn .dfn-paneled
dfn-for="AbortSignal" dfn-type="dfn" noexport=""}, given an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④②
link-type="idl"} object `signal`{.variable} and an optional
`reason`{.variable}:

1.  If `signal`{.variable} is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑦
    link-type="dfn"}, then return.

2.  Set `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①③
    link-type="dfn"} to `reason`{.variable} if it is given; otherwise to
    a new
    \"[`AbortError`{.idl}](https://webidl.spec.whatwg.org/#aborterror){#ref-for-aborterror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①②
    link-type="idl"}.

3.  Let `dependentSignalsToAbort`{.variable} be a new
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑧
    link-type="dfn"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑧
    link-type="dfn"} `dependentSignal`{.variable} of
    `signal`{.variable}'s [dependent
    signals](#abortsignal-dependent-signals){#ref-for-abortsignal-dependent-signals
    link-type="dfn"}:

    1.  If `dependentSignal`{.variable} is not
        [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑧
        link-type="dfn"}:

        1.  Set `dependentSignal`{.variable}'s [abort
            reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①④
            link-type="dfn"} to `signal`{.variable}'s [abort
            reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑤
            link-type="dfn"}.

        2.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑦
            link-type="dfn"} `dependentSignal`{.variable} to
            `dependentSignalsToAbort`{.variable}.

5.  [Run the abort
    steps](#run-the-abort-steps){#ref-for-run-the-abort-steps
    link-type="dfn"} for `signal`{.variable}.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate⑨
    link-type="dfn"} `dependentSignal`{.variable} of
    `dependentSignalsToAbort`{.variable}, [run the abort
    steps](#run-the-abort-steps){#ref-for-run-the-abort-steps①
    link-type="dfn"} for `dependentSignal`{.variable}.
:::

::: {.algorithm algorithm="run the abort steps"}
To [run the abort steps]{#run-the-abort-steps .dfn .dfn-paneled
dfn-type="dfn" noexport=""} for an
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④③
link-type="idl"} `signal`{.variable}:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⓪
    link-type="dfn"} `algorithm`{.variable} of `signal`{.variable}'s
    [abort
    algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms③
    link-type="dfn"}: run `algorithm`{.variable}.

2.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty
    link-type="dfn"} `signal`{.variable}'s [abort
    algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms④
    link-type="dfn"}.

3.  [Fire an event](#concept-event-fire){#ref-for-concept-event-fire⑤
    link-type="dfn"} named
    [`abort`{.idl}](#eventdef-abortsignal-abort){#ref-for-eventdef-abortsignal-abort①
    link-type="idl"} at `signal`{.variable}.
:::

::: {.algorithm algorithm="create a dependent abort signal"}
To [create a dependent abort signal]{#create-a-dependent-abort-signal
.dfn .dfn-paneled dfn-type="dfn" export=""} from a list of
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④④
link-type="idl"} objects `signals`{.variable}, using
`signalInterface`{.variable}, which must be either
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑤
link-type="idl"} or an interface that inherits from it, and a
`realm`{.variable}:

1.  Let `resultSignal`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#ref-for-new
    link-type="dfn"} object implementing `signalInterface`{.variable}
    using `realm`{.variable}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①①
    link-type="dfn"} `signal`{.variable} of `signals`{.variable}: if
    `signal`{.variable} is
    [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted⑨
    link-type="dfn"}, then set `resultSignal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑥
    link-type="dfn"} to `signal`{.variable}'s [abort
    reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑦
    link-type="dfn"} and return `resultSignal`{.variable}.

3.  Set `resultSignal`{.variable}'s
    [dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent
    link-type="dfn"} to true.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①②
    link-type="dfn"} `signal`{.variable} of `signals`{.variable}:

    1.  If `signal`{.variable}'s
        [dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent①
        link-type="dfn"} is false:

        1.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append②
            link-type="dfn"} `signal`{.variable} to
            `resultSignal`{.variable}'s [source
            signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals
            link-type="dfn"}.

        2.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append③
            link-type="dfn"} `resultSignal`{.variable} to
            `signal`{.variable}'s [dependent
            signals](#abortsignal-dependent-signals){#ref-for-abortsignal-dependent-signals①
            link-type="dfn"}.

    2.  Otherwise, [for
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①③
        link-type="dfn"} `sourceSignal`{.variable} of
        `signal`{.variable}'s [source
        signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals①
        link-type="dfn"}:

        1.  Assert: `sourceSignal`{.variable} is not
            [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①⓪
            link-type="dfn"} and not
            [dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent②
            link-type="dfn"}.

        2.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append④
            link-type="dfn"} `sourceSignal`{.variable} to
            `resultSignal`{.variable}'s [source
            signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals②
            link-type="dfn"}.

        3.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append⑤
            link-type="dfn"} `resultSignal`{.variable} to
            `sourceSignal`{.variable}'s [dependent
            signals](#abortsignal-dependent-signals){#ref-for-abortsignal-dependent-signals②
            link-type="dfn"}.

5.  Return `resultSignal`{.variable}.
:::

#### [3.2.1. ]{.secno}[Garbage collection]{.content}[](#abort-signal-garbage-collection){.self-link} {#abort-signal-garbage-collection .heading .settled level="3.2.1"}

A non-[aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①①
link-type="dfn"}
[dependent](#abortsignal-dependent){#ref-for-abortsignal-dependent③
link-type="dfn"}
[`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑥
link-type="idl"} object must not be garbage collected while its [source
signals](#abortsignal-source-signals){#ref-for-abortsignal-source-signals③
link-type="dfn"} is non-empty and it has registered event listeners for
its
[`abort`{.idl}](#eventdef-abortsignal-abort){#ref-for-eventdef-abortsignal-abort②
link-type="idl"} event or its [abort
algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms⑤
link-type="dfn"} is non-empty.

### [3.3. ]{.secno}[Using [`AbortController`{.idl}](#abortcontroller){#ref-for-abortcontroller⑨ link-type="idl"} and [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑦ link-type="idl"} objects in APIs]{.content}[](#abortcontroller-api-integration){.self-link} {#abortcontroller-api-integration .heading .settled level="3.3"}

Any web platform API using promises to represent operations that can be
aborted must adhere to the following:

- Accept [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑧
  link-type="idl"} objects through a `signal` dictionary member.
- Convey that the operation got aborted by rejecting the promise with
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal④⑨
  link-type="idl"} object's [abort
  reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑧
  link-type="dfn"}.
- Reject immediately if the
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑤⓪
  link-type="idl"} is already
  [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①②
  link-type="dfn"}, otherwise:
- Use the [abort
  algorithms](#abortsignal-abort-algorithms){#ref-for-abortsignal-abort-algorithms⑥
  link-type="dfn"} mechanism to observe changes to the
  [`AbortSignal`{.idl}](#abortsignal){#ref-for-abortsignal⑤①
  link-type="idl"} object and do so in a manner that does not lead to
  clashes with other observers.

::: {#aborting-ongoing-activities-spec-example .example}
[](#aborting-ongoing-activities-spec-example){.self-link}

The method steps for a promise-returning method
`doAmazingness(``options`{.variable}`)` could be as follows:

1.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①
    link-type="dfn"}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global②
    link-type="dfn"}.

2.  Let `p`{.variable} be [a new
    promise](https://webidl.spec.whatwg.org/#a-new-promise){#ref-for-a-new-promise
    link-type="dfn"}.

3.  If `options`{.variable}\[\"`signal`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists②
    link-type="dfn"}:

    1.  Let `signal`{.variable} be `options`{.variable}\[\"`signal`\"\].

    2.  If `signal`{.variable} is
        [aborted](#abortsignal-aborted){#ref-for-abortsignal-aborted①③
        link-type="dfn"}, then
        [reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject
        link-type="dfn"} `p`{.variable} with `signal`{.variable}'s
        [abort
        reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason①⑨
        link-type="dfn"} and return `p`{.variable}.

    3.  [Add the following abort
        steps](#abortsignal-add){#ref-for-abortsignal-add①
        link-type="dfn"} to `signal`{.variable}:

        1.  Stop doing amazing things.

        2.  [Reject](https://webidl.spec.whatwg.org/#reject){#ref-for-reject①
            link-type="dfn"} `p`{.variable} with `signal`{.variable}'s
            [abort
            reason](#abortsignal-abort-reason){#ref-for-abortsignal-abort-reason②⓪
            link-type="dfn"}.

4.  Run these steps [in
    parallel](https://html.spec.whatwg.org/multipage/infrastructure.html#in-parallel){#ref-for-in-parallel①
    link-type="dfn"}:

    1.  Let `amazingResult`{.variable} be the result of doing some
        amazing things.

    2.  [Queue a global
        task](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-global-task){#ref-for-queue-a-global-task①
        link-type="dfn"} on the amazing task source given
        `global`{.variable} to
        [resolve](https://webidl.spec.whatwg.org/#resolve){#ref-for-resolve
        link-type="dfn"} `p`{.variable} with `amazingResult`{.variable}.

5.  Return `p`{.variable}.
:::

APIs not using promises should still adhere to the above as much as
possible.

## [4. ]{.secno}[Nodes]{.content}[](#nodes){.self-link} {#nodes .heading .settled level="4"}

### [4.1. ]{.secno}[Introduction to \"The DOM\"]{.content}[](#introduction-to-the-dom){.self-link} {#introduction-to-the-dom .heading .settled level="4.1"}

In its original sense, \"The DOM\" is an API for accessing and
manipulating documents (in particular, HTML and XML documents). In this
specification, the term \"document\" is used for any markup-based
resource, ranging from short static documents to long essays or reports
with rich multimedia, as well as to fully-fledged interactive
applications.

Each such document is represented as a [node
tree](#concept-node-tree){#ref-for-concept-node-tree link-type="dfn"}.
Some of the [nodes](#concept-node){#ref-for-concept-node⑨
link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree①②
link-type="dfn"} can have
[children](#concept-tree-child){#ref-for-concept-tree-child⑦
link-type="dfn"}, while others are always leaves.

To illustrate, consider this HTML document:

``` {.lang-markup .highlight}
<!DOCTYPE html>
<html class=e>
 <head><title>Aliens?</title></head>
 <body>Why yes.</body>
</html>
```

It is represented as follows:

- [Document](#concept-document){#ref-for-concept-document②
  link-type="dfn"}
  - [Doctype](#concept-doctype){#ref-for-concept-doctype
    link-type="dfn"}: `html`
  - [`Element`{.idl}](#element){#ref-for-element link-type="idl"}:
    `html` [`class`{.attribute .name}=\"`e`{.attribute .value}\"]{.t2}
    - [`Element`{.idl}](#element){#ref-for-element① link-type="idl"}:
      `head`
      - [`Element`{.idl}](#element){#ref-for-element② link-type="idl"}:
        `title`
        - [`Text`{.idl}](#text){#ref-for-text link-type="idl"}: Aliens?
    - [`Text`{.idl}](#text){#ref-for-text① link-type="idl"}: ⏎␣
    - [`Element`{.idl}](#element){#ref-for-element③ link-type="idl"}:
      `body`
      - [`Text`{.idl}](#text){#ref-for-text② link-type="idl"}: Why yes.⏎

Note that, due to the magic that is [HTML
parsing](https://html.spec.whatwg.org/multipage/parsing.html#html-parser){#ref-for-html-parser
link-type="dfn"}, not all [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#ref-for-ascii-whitespace
link-type="dfn"} were turned into [`Text`{.idl}](#text){#ref-for-text③
link-type="idl"} [nodes](#concept-node){#ref-for-concept-node①⓪
link-type="dfn"}, but the general concept is clear. Markup goes in, a
[tree](#concept-tree){#ref-for-concept-tree①③ link-type="dfn"} of
[nodes](#concept-node){#ref-for-concept-node①① link-type="dfn"} comes
out.

The most excellent [Live DOM
Viewer](https://software.hixie.ch/utilities/js/live-dom-viewer/) can be
used to explore this matter in more detail.

### [4.2. ]{.secno}[Node tree]{.content}[](#node-trees){.self-link} {#node-trees .heading .settled level="4.2"}

[Nodes]{#concept-node .dfn .dfn-paneled dfn-type="dfn" export=""} are
objects that
[implement](https://webidl.spec.whatwg.org/#implements){#ref-for-implements
link-type="dfn"} [`Node`{.idl}](#node){#ref-for-node link-type="idl"}.
[Nodes](#concept-node){#ref-for-concept-node①② link-type="dfn"}
[participate](#concept-tree-participate){#ref-for-concept-tree-participate④
link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree①④
link-type="dfn"}, which is known as the [node tree]{#concept-node-tree
.dfn .dfn-paneled dfn-type="dfn" export=""}.

::: {.note role="note"}
In practice you deal with more specific objects.

Objects that
[implement](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①
link-type="dfn"} [`Node`{.idl}](#node){#ref-for-node① link-type="idl"}
also implement an inherited interface:
[`Document`{.idl}](#document){#ref-for-document link-type="idl"},
[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype
link-type="idl"},
[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment
link-type="idl"}, [`Element`{.idl}](#element){#ref-for-element④
link-type="idl"},
[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata
link-type="idl"}, or [`Attr`{.idl}](#attr){#ref-for-attr
link-type="idl"}.

Objects that implement
[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①
link-type="idl"} sometimes implement
[`ShadowRoot`{.idl}](#shadowroot){#ref-for-shadowroot① link-type="idl"}.

Objects that implement [`Element`{.idl}](#element){#ref-for-element⑤
link-type="idl"} also typically implement an inherited interface, such
as
[`HTMLAnchorElement`{.idl}](https://html.spec.whatwg.org/multipage/text-level-semantics.html#htmlanchorelement){#ref-for-htmlanchorelement
link-type="idl"}.

Objects that implement
[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①
link-type="idl"} also implement an inherited interface:
[`Text`{.idl}](#text){#ref-for-text④ link-type="idl"},
[`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction
link-type="idl"}, or [`Comment`{.idl}](#comment){#ref-for-comment
link-type="idl"}.

Objects that implement [`Text`{.idl}](#text){#ref-for-text⑤
link-type="idl"} sometimes implement
[`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection
link-type="idl"}.

Thus, every [node](#boundary-point-node){#ref-for-boundary-point-node
link-type="dfn"}'s [primary
interface](https://webidl.spec.whatwg.org/#dfn-primary-interface){#ref-for-dfn-primary-interface
link-type="dfn"} is one of:
[`Document`{.idl}](#document){#ref-for-document① link-type="idl"},
[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①
link-type="idl"},
[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②
link-type="idl"}, [`ShadowRoot`{.idl}](#shadowroot){#ref-for-shadowroot②
link-type="idl"}, [`Element`{.idl}](#element){#ref-for-element⑥
link-type="idl"} or an inherited interface of
[`Element`{.idl}](#element){#ref-for-element⑦ link-type="idl"},
[`Attr`{.idl}](#attr){#ref-for-attr① link-type="idl"},
[`Text`{.idl}](#text){#ref-for-text⑥ link-type="idl"},
[`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection①
link-type="idl"},
[`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①
link-type="idl"}, or [`Comment`{.idl}](#comment){#ref-for-comment①
link-type="idl"}.
:::

For brevity, this specification refers to an object that
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements②
link-type="dfn"} [`Node`{.idl}](#node){#ref-for-node② link-type="idl"}
and an inherited interface `NodeInterface`{.variable}, as a
`NodeInterface`{.variable} [node](#concept-node){#ref-for-concept-node①③
link-type="dfn"}.

A [node tree](#concept-node-tree){#ref-for-concept-node-tree①
link-type="dfn"} is constrained as follows, expressed as a relationship
between a [node](#concept-node){#ref-for-concept-node①④ link-type="dfn"}
and its potential
[children](#concept-tree-child){#ref-for-concept-tree-child⑧
link-type="dfn"}:

[`Document`{.idl}](#document){#ref-for-document② link-type="idl"}

:   In [tree order](#concept-tree-order){#ref-for-concept-tree-order⑤
    link-type="dfn"}:

    1.  Zero or more
        [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction②
        link-type="idl"} or
        [`Comment`{.idl}](#comment){#ref-for-comment② link-type="idl"}
        [nodes](#concept-node){#ref-for-concept-node①⑤ link-type="dfn"}.

    2.  Optionally one
        [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype②
        link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑥
        link-type="dfn"}.

    3.  Zero or more
        [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction③
        link-type="idl"} or
        [`Comment`{.idl}](#comment){#ref-for-comment③ link-type="idl"}
        [nodes](#concept-node){#ref-for-concept-node①⑦ link-type="dfn"}.

    4.  Optionally one [`Element`{.idl}](#element){#ref-for-element⑧
        link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑧
        link-type="dfn"}.

    5.  Zero or more
        [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction④
        link-type="idl"} or
        [`Comment`{.idl}](#comment){#ref-for-comment④ link-type="idl"}
        [nodes](#concept-node){#ref-for-concept-node①⑨ link-type="dfn"}.

[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③ link-type="idl"}\
[`Element`{.idl}](#element){#ref-for-element⑨ link-type="idl"}

:   Zero or more [`Element`{.idl}](#element){#ref-for-element①⓪
    link-type="idl"} or
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②
    link-type="idl"} [nodes](#concept-node){#ref-for-concept-node②⓪
    link-type="dfn"}.

[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype③ link-type="idl"}\
[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata③ link-type="idl"}\
[`Attr`{.idl}](#attr){#ref-for-attr② link-type="idl"}

:   No [children](#concept-tree-child){#ref-for-concept-tree-child⑨
    link-type="dfn"}.

[`Attr`{.idl}](#attr){#ref-for-attr③ link-type="idl"}
[nodes](#concept-node){#ref-for-concept-node②① link-type="dfn"}
[participate](#concept-tree-participate){#ref-for-concept-tree-participate⑤
link-type="dfn"} in a [tree](#concept-tree){#ref-for-concept-tree①⑤
link-type="dfn"} for historical reasons; they never have a (non-null)
[parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤
link-type="dfn"} or any
[children](#concept-tree-child){#ref-for-concept-tree-child①⓪
link-type="dfn"} and are therefore alone in a
[tree](#concept-tree){#ref-for-concept-tree①⑥ link-type="dfn"}.

To determine the [length]{#concept-node-length .dfn .dfn-paneled
dfn-for="Node" dfn-type="dfn" export=""} of a
[node](#concept-node){#ref-for-concept-node②② link-type="dfn"}
`node`{.variable}, run these steps:

1.  If `node`{.variable} is a
    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype④
    link-type="idl"} or [`Attr`{.idl}](#attr){#ref-for-attr④
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②③
    link-type="dfn"}, then return 0.

2.  If `node`{.variable} is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata④
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②④
    link-type="dfn"}, then return `node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data link-type="dfn"}'s
    [length](https://infra.spec.whatwg.org/#string-length){#ref-for-string-length
    link-type="dfn"}.

3.  Return the number of `node`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child①①
    link-type="dfn"}.

A [node](#concept-node){#ref-for-concept-node②⑤ link-type="dfn"} is
considered [empty]{#concept-node-empty .dfn .dfn-paneled dfn-for="Node"
dfn-type="dfn" export=""} if its
[length](#concept-node-length){#ref-for-concept-node-length
link-type="dfn"} is 0.

#### [4.2.1. ]{.secno}[Document tree]{.content}[](#document-trees){.self-link} {#document-trees .heading .settled level="4.2.1"}

A [document tree]{#concept-document-tree .dfn .dfn-paneled
dfn-type="dfn" export=""} is a [node
tree](#concept-node-tree){#ref-for-concept-node-tree② link-type="dfn"}
whose [root](#concept-tree-root){#ref-for-concept-tree-root⑧
link-type="dfn"} is a
[document](#concept-document){#ref-for-concept-document③
link-type="dfn"}.

The [document element]{#document-element .dfn .dfn-paneled
dfn-type="dfn" export=""} of a
[document](#concept-document){#ref-for-concept-document④
link-type="dfn"} is the
[element](#concept-element){#ref-for-concept-element④ link-type="dfn"}
whose [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥
link-type="dfn"} is that
[document](#concept-document){#ref-for-concept-document⑤
link-type="dfn"}, if it exists; otherwise null.

Per the [node tree](#concept-node-tree){#ref-for-concept-node-tree③
link-type="dfn"} constraints, there can be only one such
[element](#concept-element){#ref-for-concept-element⑤ link-type="dfn"}.

A [node](#concept-node){#ref-for-concept-node②⑥ link-type="dfn"} is [in
a document tree]{#in-a-document-tree .dfn .dfn-paneled dfn-type="dfn"
export=""} if its [root](#concept-tree-root){#ref-for-concept-tree-root⑨
link-type="dfn"} is a
[document](#concept-document){#ref-for-concept-document⑥
link-type="dfn"}.

A [node](#concept-node){#ref-for-concept-node②⑦ link-type="dfn"} is [in
a document]{#in-a-document .dfn .dfn-paneled dfn-type="dfn" export=""}
if it is [in a document
tree](#in-a-document-tree){#ref-for-in-a-document-tree link-type="dfn"}.
[The term [in a document](#in-a-document){#ref-for-in-a-document
link-type="dfn"} is no longer supposed to be used. It indicates that the
standard using it has not been updated to account for [shadow
trees](#concept-shadow-tree){#ref-for-concept-shadow-tree②
link-type="dfn"}.]{.note}

#### [4.2.2. ]{.secno}[Shadow tree]{.content}[](#shadow-trees){.self-link} {#shadow-trees .heading .settled level="4.2.2"}

A [shadow tree]{#concept-shadow-tree .dfn .dfn-paneled dfn-type="dfn"
export=""} is a [node
tree](#concept-node-tree){#ref-for-concept-node-tree④ link-type="dfn"}
whose [root](#concept-tree-root){#ref-for-concept-tree-root①⓪
link-type="dfn"} is a [shadow
root](#concept-shadow-root){#ref-for-concept-shadow-root⑥
link-type="dfn"}.

A [shadow root](#concept-shadow-root){#ref-for-concept-shadow-root⑦
link-type="dfn"} is always attached to another [node
tree](#concept-node-tree){#ref-for-concept-node-tree⑤ link-type="dfn"}
through its
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host
link-type="dfn"}. A [shadow
tree](#concept-shadow-tree){#ref-for-concept-shadow-tree③
link-type="dfn"} is therefore never alone. The [node
tree](#concept-node-tree){#ref-for-concept-node-tree⑥ link-type="dfn"}
of a [shadow root](#concept-shadow-root){#ref-for-concept-shadow-root⑧
link-type="dfn"}'s
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host①
link-type="dfn"} is sometimes referred to as the [light
tree]{#concept-light-tree .dfn .dfn-paneled dfn-type="dfn" export=""}.

A [shadow tree](#concept-shadow-tree){#ref-for-concept-shadow-tree④
link-type="dfn"}'s corresponding [light
tree](#concept-light-tree){#ref-for-concept-light-tree link-type="dfn"}
can be a [shadow
tree](#concept-shadow-tree){#ref-for-concept-shadow-tree⑤
link-type="dfn"} itself.

A [node](#concept-node){#ref-for-concept-node②⑧ link-type="dfn"} is
[connected]{#connected .dfn .dfn-paneled dfn-type="dfn" export=""} if
its [shadow-including
root](#concept-shadow-including-root){#ref-for-concept-shadow-including-root
link-type="dfn"} is a
[document](#concept-document){#ref-for-concept-document⑦
link-type="dfn"}.

##### [4.2.2.1. ]{.secno}[Slots]{.content}[](#shadow-tree-slots){.self-link} {#shadow-tree-slots .heading .settled level="4.2.2.1"}

A [shadow tree](#concept-shadow-tree){#ref-for-concept-shadow-tree⑥
link-type="dfn"} contains zero or more
[elements](#concept-element){#ref-for-concept-element⑥ link-type="dfn"}
that are [slots]{#concept-slot .dfn .dfn-paneled dfn-type="dfn"
export="" lt="slot"}.

A [slot](#concept-slot){#ref-for-concept-slot① link-type="dfn"} can only
be created through HTML's
[`slot`](https://html.spec.whatwg.org/multipage/scripting.html#the-slot-element){#ref-for-the-slot-element
link-type="element"} element.

A [slot](#concept-slot){#ref-for-concept-slot② link-type="dfn"} has an
associated [name]{#slot-name .dfn .dfn-paneled dfn-for="slot"
dfn-type="dfn" export=""} (a string). Unless stated otherwise it is the
empty string.

Use these [attribute change
steps](#concept-element-attributes-change-ext){#ref-for-concept-element-attributes-change-ext
link-type="dfn"} to update a
[slot](#concept-slot){#ref-for-concept-slot③ link-type="dfn"}'s
[name](#slot-name){#ref-for-slot-name link-type="dfn"}:

1.  If `element`{.variable} is a
    [slot](#concept-slot){#ref-for-concept-slot④ link-type="dfn"},
    `localName`{.variable} is `name`, and `namespace`{.variable} is
    null:

    1.  If `value`{.variable} is `oldValue`{.variable}, then return.

    2.  If `value`{.variable} is null and `oldValue`{.variable} is the
        empty string, then return.

    3.  If `value`{.variable} is the empty string and
        `oldValue`{.variable} is null, then return.

    4.  If `value`{.variable} is null or the empty string, then set
        `element`{.variable}'s [name](#slot-name){#ref-for-slot-name①
        link-type="dfn"} to the empty string.

    5.  Otherwise, set `element`{.variable}'s
        [name](#slot-name){#ref-for-slot-name② link-type="dfn"} to
        `value`{.variable}.

    6.  Run [assign slottables for a
        tree](#assign-slotables-for-a-tree){#ref-for-assign-slotables-for-a-tree
        link-type="dfn"} with `element`{.variable}'s
        [root](#concept-tree-root){#ref-for-concept-tree-root①①
        link-type="dfn"}.

The first [slot](#concept-slot){#ref-for-concept-slot⑤ link-type="dfn"}
in a [shadow tree](#concept-shadow-tree){#ref-for-concept-shadow-tree⑦
link-type="dfn"}, in [tree
order](#concept-tree-order){#ref-for-concept-tree-order⑥
link-type="dfn"}, whose [name](#slot-name){#ref-for-slot-name③
link-type="dfn"} is the empty string, is sometimes known as the
\"default slot\".

A [slot](#concept-slot){#ref-for-concept-slot⑥ link-type="dfn"} has an
associated [assigned nodes]{#slot-assigned-nodes .dfn .dfn-paneled
dfn-for="slot" dfn-type="dfn" export=""} (a list of
[slottables](#concept-slotable){#ref-for-concept-slotable②
link-type="dfn"}). Unless stated otherwise it is empty.

##### [4.2.2.2. ]{.secno}[Slottables]{.content}[](#light-tree-slotables){.self-link} {#light-tree-slotables .heading .settled level="4.2.2.2"}

[`Element`{.idl}](#element){#ref-for-element①① link-type="idl"} and
[`Text`{.idl}](#text){#ref-for-text⑦ link-type="idl"}
[nodes](#concept-node){#ref-for-concept-node②⑨ link-type="dfn"} are
[slottables]{#concept-slotable .dfn .dfn-paneled dfn-type="dfn"
export="" lt="slottable"}.

A [slot](#concept-slot){#ref-for-concept-slot⑦ link-type="dfn"} can be a
[slottable](#concept-slotable){#ref-for-concept-slotable③
link-type="dfn"}.

A [slottable](#concept-slotable){#ref-for-concept-slotable④
link-type="dfn"} has an associated [name]{#slotable-name .dfn
.dfn-paneled dfn-for="slottable" dfn-type="dfn" export=""} (a string).
Unless stated otherwise it is the empty string.

Use these [attribute change
steps](#concept-element-attributes-change-ext){#ref-for-concept-element-attributes-change-ext①
link-type="dfn"} to update a
[slottable](#concept-slotable){#ref-for-concept-slotable⑤
link-type="dfn"}'s [name](#slotable-name){#ref-for-slotable-name
link-type="dfn"}:

1.  If `localName`{.variable} is `slot` and `namespace`{.variable} is
    null:

    1.  If `value`{.variable} is `oldValue`{.variable}, then return.

    2.  If `value`{.variable} is null and `oldValue`{.variable} is the
        empty string, then return.

    3.  If `value`{.variable} is the empty string and
        `oldValue`{.variable} is null, then return.

    4.  If `value`{.variable} is null or the empty string, then set
        `element`{.variable}'s
        [name](#slotable-name){#ref-for-slotable-name① link-type="dfn"}
        to the empty string.

    5.  Otherwise, set `element`{.variable}'s
        [name](#slotable-name){#ref-for-slotable-name② link-type="dfn"}
        to `value`{.variable}.

    6.  If `element`{.variable} is
        [assigned](#slotable-assigned){#ref-for-slotable-assigned②
        link-type="dfn"}, then run [assign
        slottables](#assign-slotables){#ref-for-assign-slotables
        link-type="dfn"} for `element`{.variable}'s [assigned
        slot](#slotable-assigned-slot){#ref-for-slotable-assigned-slot
        link-type="dfn"}.

    7.  Run [assign a slot](#assign-a-slot){#ref-for-assign-a-slot
        link-type="dfn"} for `element`{.variable}.

A [slottable](#concept-slotable){#ref-for-concept-slotable⑥
link-type="dfn"} has an associated [assigned
slot]{#slotable-assigned-slot .dfn .dfn-paneled dfn-for="slottable"
dfn-type="dfn" export=""} (null or a
[slot](#concept-slot){#ref-for-concept-slot⑧ link-type="dfn"}). Unless
stated otherwise it is null. A
[slottable](#concept-slotable){#ref-for-concept-slotable⑦
link-type="dfn"} is [assigned]{#slotable-assigned .dfn .dfn-paneled
dfn-for="slottable" dfn-type="dfn" export=""} if its [assigned
slot](#slotable-assigned-slot){#ref-for-slotable-assigned-slot①
link-type="dfn"} is non-null.

A [slottable](#concept-slotable){#ref-for-concept-slotable⑧
link-type="dfn"} has an associated [manual slot
assignment]{#slottable-manual-slot-assignment .dfn .dfn-paneled
dfn-for="slottable" dfn-type="dfn" export=""} (null or a
[slot](#concept-slot){#ref-for-concept-slot⑨ link-type="dfn"}). Unless
stated otherwise, it is null.

A [slottable](#concept-slotable){#ref-for-concept-slotable⑨
link-type="dfn"}'s [manual slot
assignment](#slottable-manual-slot-assignment){#ref-for-slottable-manual-slot-assignment
link-type="dfn"} can be implemented using a weak reference to the
[slot](#concept-slot){#ref-for-concept-slot①⓪ link-type="dfn"}, because
this variable is not directly accessible from script.

##### [4.2.2.3. ]{.secno}[Finding slots and slottables]{.content}[](#finding-slots-and-slotables){.self-link} {#finding-slots-and-slotables .heading .settled level="4.2.2.3"}

::: {.algorithm algorithm="find a slot"}
To [find a slot]{#find-a-slot .dfn .dfn-paneled dfn-type="dfn"
noexport=""} for a given
[slottable](#concept-slotable){#ref-for-concept-slotable①⓪
link-type="dfn"} `slottable`{.variable} and an optional boolean
`open`{.variable} (default false):

1.  If `slottable`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑦
    link-type="dfn"} is null, then return null.

2.  Let `shadow`{.variable} be `slottable`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑧
    link-type="dfn"}'s [shadow
    root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root
    link-type="dfn"}.

3.  If `shadow`{.variable} is null, then return null.

4.  If `open`{.variable} is true and `shadow`{.variable}'s
    [mode](#shadowroot-mode){#ref-for-shadowroot-mode③ link-type="dfn"}
    is not \"`open`\", then return null.

5.  If `shadow`{.variable}'s [slot
    assignment](#shadowroot-slot-assignment){#ref-for-shadowroot-slot-assignment
    link-type="dfn"} is \"`manual`\", then return the
    [slot](#concept-slot){#ref-for-concept-slot①① link-type="dfn"} in
    `shadow`{.variable}'s
    [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant③
    link-type="dfn"} whose [manually assigned
    nodes](https://html.spec.whatwg.org/multipage/scripting.html#manually-assigned-nodes){#ref-for-manually-assigned-nodes
    link-type="dfn"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain③
    link-type="dfn"} `slottable`{.variable}, if any; otherwise null.

6.  Return the first [slot](#concept-slot){#ref-for-concept-slot①②
    link-type="dfn"} in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order⑦
    link-type="dfn"} in `shadow`{.variable}'s
    [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant④
    link-type="dfn"} whose [name](#slot-name){#ref-for-slot-name④
    link-type="dfn"} is `slottable`{.variable}'s
    [name](#slotable-name){#ref-for-slotable-name③ link-type="dfn"}, if
    any; otherwise null.
:::

::: {.algorithm algorithm="find slottables"}
To [find slottables]{#find-slotables .dfn .dfn-paneled dfn-type="dfn"
export=""} for a given [slot](#concept-slot){#ref-for-concept-slot①③
link-type="dfn"} `slot`{.variable}:

1.  Let `result`{.variable} be « ».

2.  Let `root`{.variable} be `slot`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root①②
    link-type="dfn"}.

3.  If `root`{.variable} is not a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root⑨
    link-type="dfn"}, then return `result`{.variable}.

4.  Let `host`{.variable} be `root`{.variable}'s
    [host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host②
    link-type="dfn"}.

5.  If `root`{.variable}'s [slot
    assignment](#shadowroot-slot-assignment){#ref-for-shadowroot-slot-assignment①
    link-type="dfn"} is \"`manual`\":

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①④
        link-type="dfn"}
        [slottable](#concept-slotable){#ref-for-concept-slotable①①
        link-type="dfn"} `slottable`{.variable} of `slot`{.variable}'s
        [manually assigned
        nodes](https://html.spec.whatwg.org/multipage/scripting.html#manually-assigned-nodes){#ref-for-manually-assigned-nodes①
        link-type="dfn"}, if `slottable`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑨
        link-type="dfn"} is `host`{.variable},
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑧
        link-type="dfn"} `slottable`{.variable} to `result`{.variable}.

6.  Otherwise, for each
    [slottable](#concept-slotable){#ref-for-concept-slotable①②
    link-type="dfn"}
    [child](#concept-tree-child){#ref-for-concept-tree-child①②
    link-type="dfn"} `slottable`{.variable} of `host`{.variable}, in
    [tree order](#concept-tree-order){#ref-for-concept-tree-order⑧
    link-type="dfn"}:

    1.  Let `foundSlot`{.variable} be the result of [finding a
        slot](#find-a-slot){#ref-for-find-a-slot link-type="dfn"} given
        `slottable`{.variable}.

    2.  If `foundSlot`{.variable} is `slot`{.variable}, then
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append⑨
        link-type="dfn"} `slottable`{.variable} to `result`{.variable}.

7.  Return `result`{.variable}.
:::

::: {.algorithm algorithm="find flattened slottables"}
To [find flattened slottables]{#find-flattened-slotables .dfn
.dfn-paneled dfn-type="dfn" export=""} for a given
[slot](#concept-slot){#ref-for-concept-slot①④ link-type="dfn"}
`slot`{.variable}:

1.  Let `result`{.variable} be « ».

2.  If `slot`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root①③
    link-type="dfn"} is not a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root①⓪
    link-type="dfn"}, then return `result`{.variable}.

3.  Let `slottables`{.variable} be the result of [finding
    slottables](#find-slotables){#ref-for-find-slotables
    link-type="dfn"} given `slot`{.variable}.

4.  If `slottables`{.variable} is the empty list, then append each
    [slottable](#concept-slotable){#ref-for-concept-slotable①③
    link-type="dfn"}
    [child](#concept-tree-child){#ref-for-concept-tree-child①③
    link-type="dfn"} of `slot`{.variable}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order⑨
    link-type="dfn"}, to `slottables`{.variable}.

5.  For each `node`{.variable} of `slottables`{.variable}:

    1.  If `node`{.variable} is a
        [slot](#concept-slot){#ref-for-concept-slot①⑤ link-type="dfn"}
        whose [root](#concept-tree-root){#ref-for-concept-tree-root①④
        link-type="dfn"} is a [shadow
        root](#concept-shadow-root){#ref-for-concept-shadow-root①①
        link-type="dfn"}:

        1.  Let `temporaryResult`{.variable} be the result of [finding
            flattened
            slottables](#find-flattened-slotables){#ref-for-find-flattened-slotables
            link-type="dfn"} given `node`{.variable}.

        2.  Append each
            [slottable](#concept-slotable){#ref-for-concept-slotable①④
            link-type="dfn"} in `temporaryResult`{.variable}, in order,
            to `result`{.variable}.

    2.  Otherwise, append `node`{.variable} to `result`{.variable}.

6.  Return `result`{.variable}.
:::

##### [4.2.2.4. ]{.secno}[Assigning slottables and slots]{.content}[](#assigning-slotables-and-slots){.self-link} {#assigning-slotables-and-slots .heading .settled level="4.2.2.4"}

::: {.algorithm algorithm="assign slottables"}
To [assign slottables]{#assign-slotables .dfn .dfn-paneled
dfn-type="dfn" noexport=""} for a
[slot](#concept-slot){#ref-for-concept-slot①⑥ link-type="dfn"}
`slot`{.variable}:

1.  Let `slottables`{.variable} be the result of [finding
    slottables](#find-slotables){#ref-for-find-slotables①
    link-type="dfn"} for `slot`{.variable}.

2.  If `slottables`{.variable} and `slot`{.variable}'s [assigned
    nodes](#slot-assigned-nodes){#ref-for-slot-assigned-nodes
    link-type="dfn"} are not identical, then run [signal a slot
    change](#signal-a-slot-change){#ref-for-signal-a-slot-change
    link-type="dfn"} for `slot`{.variable}.

3.  Set `slot`{.variable}'s [assigned
    nodes](#slot-assigned-nodes){#ref-for-slot-assigned-nodes①
    link-type="dfn"} to `slottables`{.variable}.

4.  For each `slottable`{.variable} of `slottables`{.variable}, set
    `slottable`{.variable}'s [assigned
    slot](#slotable-assigned-slot){#ref-for-slotable-assigned-slot②
    link-type="dfn"} to `slot`{.variable}.
:::

::: {.algorithm algorithm="assign slottables for a tree"}
To [assign slottables for a tree]{#assign-slotables-for-a-tree .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given a
[node](#concept-node){#ref-for-concept-node③⓪ link-type="dfn"}
`root`{.variable}, run [assign
slottables](#assign-slotables){#ref-for-assign-slotables①
link-type="dfn"} for each [slot](#concept-slot){#ref-for-concept-slot①⑦
link-type="dfn"} of `root`{.variable}'s [inclusive
descendants](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant
link-type="dfn"}, in [tree
order](#concept-tree-order){#ref-for-concept-tree-order①⓪
link-type="dfn"}.
:::

::: {.algorithm algorithm="assign a slot"}
To [assign a slot]{#assign-a-slot .dfn .dfn-paneled dfn-type="dfn"
noexport=""}, given a
[slottable](#concept-slotable){#ref-for-concept-slotable①⑤
link-type="dfn"} `slottable`{.variable}:

1.  Let `slot`{.variable} be the result of [finding a
    slot](#find-a-slot){#ref-for-find-a-slot① link-type="dfn"} with
    `slottable`{.variable}.

2.  If `slot`{.variable} is non-null, then run [assign
    slottables](#assign-slotables){#ref-for-assign-slotables②
    link-type="dfn"} for `slot`{.variable}.
:::

##### [4.2.2.5. ]{.secno}[Signaling slot change]{.content}[](#signaling-slot-change){.self-link} {#signaling-slot-change .heading .settled level="4.2.2.5"}

Each [similar-origin window
agent](https://html.spec.whatwg.org/multipage/webappapis.html#similar-origin-window-agent){#ref-for-similar-origin-window-agent
link-type="dfn"} has [signal slots]{#signal-slot-list .dfn .dfn-paneled
dfn-type="dfn" noexport=""} (a
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set⑤
link-type="dfn"} of [slots](#concept-slot){#ref-for-concept-slot①⑧
link-type="dfn"}), which is initially empty.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

::: {.algorithm algorithm="signal a slot change"}
To [signal a slot change]{#signal-a-slot-change .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, for a
[slot](#concept-slot){#ref-for-concept-slot①⑨ link-type="dfn"}
`slot`{.variable}:

1.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append⑥
    link-type="dfn"} `slot`{.variable} to `slot`{.variable}'s [relevant
    agent](https://html.spec.whatwg.org/multipage/webappapis.html#relevant-agent){#ref-for-relevant-agent
    link-type="dfn"}'s [signal
    slots](#signal-slot-list){#ref-for-signal-slot-list
    link-type="dfn"}.

2.  [Queue a mutation observer
    microtask](#queue-a-mutation-observer-compound-microtask){#ref-for-queue-a-mutation-observer-compound-microtask
    link-type="dfn"}.
:::

#### [4.2.3. ]{.secno}[Mutation algorithms]{.content}[](#mutation-algorithms){.self-link} {#mutation-algorithms .heading .settled level="4.2.3"}

To [ensure pre-insert
validity]{#concept-node-ensure-pre-insertion-validity .dfn .dfn-paneled
dfn-type="dfn" noexport=""} of a
[node](#concept-node){#ref-for-concept-node③① link-type="dfn"}
`node`{.variable} into a [node](#concept-node){#ref-for-concept-node③②
link-type="dfn"} `parent`{.variable} before a
[node](#concept-node){#ref-for-concept-node③③ link-type="dfn"}
`child`{.variable}:

1.  If `parent`{.variable} is not a
    [`Document`{.idl}](#document){#ref-for-document③ link-type="idl"},
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment④
    link-type="idl"}, or [`Element`{.idl}](#element){#ref-for-element①②
    link-type="idl"} [node](#concept-node){#ref-for-concept-node③④
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①③
    link-type="idl"}.

2.  If `node`{.variable} is a [host-including inclusive
    ancestor](#concept-tree-host-including-inclusive-ancestor){#ref-for-concept-tree-host-including-inclusive-ancestor
    link-type="dfn"} of `parent`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①④
    link-type="idl"}.

3.  If `child`{.variable} is non-null and its
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①⓪
    link-type="dfn"} is not `parent`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨
    link-type="dfn"} a
    \"[`NotFoundError`{.idl}](https://webidl.spec.whatwg.org/#notfounderror){#ref-for-notfounderror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⑤
    link-type="idl"}.

4.  If `node`{.variable} is not a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment⑤
    link-type="idl"},
    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype⑤
    link-type="idl"}, [`Element`{.idl}](#element){#ref-for-element①③
    link-type="idl"}, or
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata⑤
    link-type="idl"} [node](#concept-node){#ref-for-concept-node③⑤
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⓪
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⑥
    link-type="idl"}.

5.  If either `node`{.variable} is a
    [`Text`{.idl}](#text){#ref-for-text⑧ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node③⑥ link-type="dfn"} and
    `parent`{.variable} is a
    [document](#concept-document){#ref-for-concept-document⑧
    link-type="dfn"}, or `node`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype①
    link-type="dfn"} and `parent`{.variable} is not a
    [document](#concept-document){#ref-for-concept-document⑨
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①①
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⑦
    link-type="idl"}.

6.  If `parent`{.variable} is a
    [document](#concept-document){#ref-for-concept-document①⓪
    link-type="dfn"}, and any of the statements below, switched on the
    interface `node`{.variable}
    [implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements③
    link-type="dfn"}, are true, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①②
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⑧
    link-type="idl"}.

    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment⑥ link-type="idl"}

    :   If `node`{.variable} has more than one
        [element](#concept-element){#ref-for-concept-element⑦
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child①④
        link-type="dfn"} or has a [`Text`{.idl}](#text){#ref-for-text⑨
        link-type="idl"} [node](#concept-node){#ref-for-concept-node③⑦
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child①⑤
        link-type="dfn"}.

        Otherwise, if `node`{.variable} has one
        [element](#concept-element){#ref-for-concept-element⑧
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child①⑥
        link-type="dfn"} and either `parent`{.variable} has an
        [element](#concept-element){#ref-for-concept-element⑨
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child①⑦
        link-type="dfn"}, `child`{.variable} is a
        [doctype](#concept-doctype){#ref-for-concept-doctype②
        link-type="dfn"}, or `child`{.variable} is non-null and a
        [doctype](#concept-doctype){#ref-for-concept-doctype③
        link-type="dfn"} is
        [following](#concept-tree-following){#ref-for-concept-tree-following②
        link-type="dfn"} `child`{.variable}.

    [`Element`{.idl}](#element){#ref-for-element①④ link-type="idl"}

    :   `parent`{.variable} has an
        [element](#concept-element){#ref-for-concept-element①⓪
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child①⑧
        link-type="dfn"}, `child`{.variable} is a
        [doctype](#concept-doctype){#ref-for-concept-doctype④
        link-type="dfn"}, or `child`{.variable} is non-null and a
        [doctype](#concept-doctype){#ref-for-concept-doctype⑤
        link-type="dfn"} is
        [following](#concept-tree-following){#ref-for-concept-tree-following③
        link-type="dfn"} `child`{.variable}.

    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype⑥ link-type="idl"}

    :   `parent`{.variable} has a
        [doctype](#concept-doctype){#ref-for-concept-doctype⑥
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child①⑨
        link-type="dfn"}, `child`{.variable} is non-null and an
        [element](#concept-element){#ref-for-concept-element①①
        link-type="dfn"} is
        [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding③
        link-type="dfn"} `child`{.variable}, or `child`{.variable} is
        null and `parent`{.variable} has an
        [element](#concept-element){#ref-for-concept-element①②
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child②⓪
        link-type="dfn"}.

To [pre-insert]{#concept-node-pre-insert .dfn .dfn-paneled
dfn-type="dfn" export=""} a `node`{.variable} into a `parent`{.variable}
before a `child`{.variable}, run these steps:

1.  [Ensure pre-insert
    validity](#concept-node-ensure-pre-insertion-validity){#ref-for-concept-node-ensure-pre-insertion-validity
    link-type="dfn"} of `node`{.variable} into `parent`{.variable}
    before `child`{.variable}.

2.  Let `referenceChild`{.variable} be `child`{.variable}.

3.  If `referenceChild`{.variable} is `node`{.variable}, then set
    `referenceChild`{.variable} to `node`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling
    link-type="dfn"}.

4.  [Insert](#concept-node-insert){#ref-for-concept-node-insert
    link-type="dfn"} `node`{.variable} into `parent`{.variable} before
    `referenceChild`{.variable}.

5.  Return `node`{.variable}.

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications①
link-type="dfn"} may define [insertion steps]{#concept-node-insert-ext
.dfn .dfn-paneled dfn-type="dfn" export=""} for all or some
[nodes](#concept-node){#ref-for-concept-node③⑧ link-type="dfn"}. The
algorithm is passed `insertedNode`{.variable}, as indicated in the
[insert](#concept-node-insert){#ref-for-concept-node-insert①
link-type="dfn"} algorithm below. These steps must not modify the [node
tree](#concept-node-tree){#ref-for-concept-node-tree⑦ link-type="dfn"}
that `insertedNode`{.variable}
[participates](#concept-tree-participate){#ref-for-concept-tree-participate⑥
link-type="dfn"} in, create [browsing
contexts](https://html.spec.whatwg.org/multipage/document-sequences.html#browsing-context){#ref-for-browsing-context
link-type="dfn"}, [fire
events](#concept-event-fire){#ref-for-concept-event-fire⑥
link-type="dfn"}, or otherwise execute JavaScript. These steps may
[queue
tasks](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-global-task){#ref-for-queue-a-global-task②
link-type="dfn"} to do these things asynchronously, however.

::: {#example-foo-what-do-i-put-here .example}
[](#example-foo-what-do-i-put-here){.self-link}

While the [insertion
steps](#concept-node-insert-ext){#ref-for-concept-node-insert-ext
link-type="dfn"} cannot execute JavaScript (among other things), they
will indeed have script-observable consequences. Consider the below
example:

``` {.lang-javascript .highlight}
const h1 = document.querySelector('h1');

const fragment = new DocumentFragment();
const script = fragment.appendChild(document.createElement('script'));
const style = fragment.appendChild(document.createElement('style'));

script.innerText= 'console.log(getComputedStyle(h1).color)'; // Logs 'rgb(255, 0, 0)'
style.innerText = 'h1 {color: rgb(255, 0, 0);}';

document.body.append(fragment);
```

The script in the above example logs `'rgb(255, 0, 0)'`{.lang-javascript
.highlight} because the following happen in order:

1.  The [insert](#concept-node-insert){#ref-for-concept-node-insert②
    link-type="dfn"} algorithm runs, which will insert the
    [`script`](https://html.spec.whatwg.org/multipage/scripting.html#script){#ref-for-script
    link-type="element"} and
    ` `[`style`](https://html.spec.whatwg.org/multipage/semantics.html#the-style-element){#ref-for-the-style-element
    link-type="element"} elements in order.

    1.  The HTML Standard's [insertion
        steps](#concept-node-insert-ext){#ref-for-concept-node-insert-ext①
        link-type="dfn"} run for the
        [`script`](https://html.spec.whatwg.org/multipage/scripting.html#script){#ref-for-script①
        link-type="element"} element; they do nothing.
        [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    2.  The HTML Standard's [insertion
        steps](#concept-node-insert-ext){#ref-for-concept-node-insert-ext②
        link-type="dfn"} run for the
        ` `[`style`](https://html.spec.whatwg.org/multipage/semantics.html#the-style-element){#ref-for-the-style-element①
        link-type="element"} element; they immediately apply its style
        rules to the document.
        [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    3.  The HTML Standard's [post-connection
        steps](#concept-node-post-connection-ext){#ref-for-concept-node-post-connection-ext
        link-type="dfn"} run for the
        [`script`](https://html.spec.whatwg.org/multipage/scripting.html#script){#ref-for-script②
        link-type="element"} element; they run the script, which
        immediately observes the style rules that were applied in the
        above step.
        [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}
:::

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications②
link-type="dfn"} may also define [post-connection
steps]{#concept-node-post-connection-ext .dfn .dfn-paneled
dfn-type="dfn" export=""} for all or some
[nodes](#concept-node){#ref-for-concept-node③⑨ link-type="dfn"}. The
algorithm is passed `connectedNode`{.variable}, as indicated in the
[insert](#concept-node-insert){#ref-for-concept-node-insert③
link-type="dfn"} algorithm below.

The purpose of the [post-connection
steps](#concept-node-post-connection-ext){#ref-for-concept-node-post-connection-ext①
link-type="dfn"} is to provide an opportunity for
[nodes](#concept-node){#ref-for-concept-node④⓪ link-type="dfn"} to
perform any connection-related operations that modify the [node
tree](#concept-node-tree){#ref-for-concept-node-tree⑧ link-type="dfn"}
that `connectedNode`{.variable}
[participates](#concept-tree-participate){#ref-for-concept-tree-participate⑦
link-type="dfn"} in, create [browsing
contexts](https://html.spec.whatwg.org/multipage/document-sequences.html#browsing-context){#ref-for-browsing-context①
link-type="dfn"}, or otherwise execute JavaScript. These steps allow a
batch of [nodes](#concept-node){#ref-for-concept-node④① link-type="dfn"}
to be [inserted](#concept-node-insert){#ref-for-concept-node-insert④
link-type="dfn"} *atomically* with respect to script, with all major
side effects occurring *after* the batch insertions into the [node
tree](#concept-node-tree){#ref-for-concept-node-tree⑨ link-type="dfn"}
is complete. This ensures that all pending [node
tree](#concept-node-tree){#ref-for-concept-node-tree①⓪ link-type="dfn"}
insertions completely finish before more insertions can occur.

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications③
link-type="dfn"} may define [children changed
steps]{#concept-node-children-changed-ext .dfn .dfn-paneled
dfn-type="dfn" export=""} for all or some
[nodes](#concept-node){#ref-for-concept-node④② link-type="dfn"}. The
algorithm is passed no argument and is called from
[insert](#concept-node-insert){#ref-for-concept-node-insert⑤
link-type="dfn"},
[remove](#concept-node-remove){#ref-for-concept-node-remove
link-type="dfn"}, and [replace
data](#concept-cd-replace){#ref-for-concept-cd-replace link-type="dfn"}.

To [insert]{#concept-node-insert .dfn .dfn-paneled dfn-type="dfn"
export=""} a `node`{.variable} into a `parent`{.variable} before a
`child`{.variable}, with an optional *suppress observers flag*, run
these steps:

1.  Let `nodes`{.variable} be `node`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child②①
    link-type="dfn"}, if `node`{.variable} is a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment⑦
    link-type="idl"} [node](#concept-node){#ref-for-concept-node④③
    link-type="dfn"}; otherwise « `node`{.variable} ».

2.  Let `count`{.variable} be `nodes`{.variable}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size②
    link-type="dfn"}.

3.  If `count`{.variable} is 0, then return.

4.  If `node`{.variable} is a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node④④
    link-type="dfn"}:

    1.  [Remove](#concept-node-remove){#ref-for-concept-node-remove①
        link-type="dfn"} its
        [children](#concept-tree-child){#ref-for-concept-tree-child②②
        link-type="dfn"} with the *suppress observers flag* set.

    2.  [Queue a tree mutation
        record](#queue-a-tree-mutation-record){#ref-for-queue-a-tree-mutation-record
        link-type="dfn"} for `node`{.variable} with « »,
        `nodes`{.variable}, null, and null.

        This step intentionally does not pay attention to the *suppress
        observers flag*.

5.  If `child`{.variable} is non-null:

    1.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range
        link-type="dfn"} whose [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node
        link-type="dfn"} is `parent`{.variable} and [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset
        link-type="dfn"} is greater than `child`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index
        link-type="dfn"}, increase its [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①
        link-type="dfn"} by `count`{.variable}.

    2.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range①
        link-type="dfn"} whose [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node
        link-type="dfn"} is `parent`{.variable} and [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset
        link-type="dfn"} is greater than `child`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index①
        link-type="dfn"}, increase its [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①
        link-type="dfn"} by `count`{.variable}.

6.  Let `previousSibling`{.variable} be `child`{.variable}'s [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling
    link-type="dfn"} or `parent`{.variable}'s [last
    child](#concept-tree-last-child){#ref-for-concept-tree-last-child
    link-type="dfn"} if `child`{.variable} is null.

7.  For each `node`{.variable} in `nodes`{.variable}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order①①
    link-type="dfn"}:

    1.  [Adopt](#concept-node-adopt){#ref-for-concept-node-adopt
        link-type="dfn"} `node`{.variable} into `parent`{.variable}'s
        [node
        document](#concept-node-document){#ref-for-concept-node-document③
        link-type="dfn"}.

    2.  If `child`{.variable} is null, then
        [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append⑦
        link-type="dfn"} `node`{.variable} to `parent`{.variable}'s
        [children](#concept-tree-child){#ref-for-concept-tree-child②③
        link-type="dfn"}.

    3.  Otherwise,
        [insert](https://infra.spec.whatwg.org/#list-insert){#ref-for-list-insert
        link-type="dfn"} `node`{.variable} into `parent`{.variable}'s
        [children](#concept-tree-child){#ref-for-concept-tree-child②④
        link-type="dfn"} before `child`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index②
        link-type="dfn"}.

    4.  If `parent`{.variable} is a [shadow
        host](#element-shadow-host){#ref-for-element-shadow-host
        link-type="dfn"} whose [shadow
        root](#concept-shadow-root){#ref-for-concept-shadow-root①②
        link-type="dfn"}'s [slot
        assignment](#shadowroot-slot-assignment){#ref-for-shadowroot-slot-assignment②
        link-type="dfn"} is \"`named`\" and `node`{.variable} is a
        [slottable](#concept-slotable){#ref-for-concept-slotable①⑥
        link-type="dfn"}, then [assign a
        slot](#assign-a-slot){#ref-for-assign-a-slot① link-type="dfn"}
        for `node`{.variable}.

    5.  If `parent`{.variable}'s
        [root](#concept-tree-root){#ref-for-concept-tree-root①⑤
        link-type="dfn"} is a [shadow
        root](#concept-shadow-root){#ref-for-concept-shadow-root①③
        link-type="dfn"}, and `parent`{.variable} is a
        [slot](#concept-slot){#ref-for-concept-slot②⓪ link-type="dfn"}
        whose [assigned
        nodes](#slot-assigned-nodes){#ref-for-slot-assigned-nodes②
        link-type="dfn"} is the empty list, then run [signal a slot
        change](#signal-a-slot-change){#ref-for-signal-a-slot-change①
        link-type="dfn"} for `parent`{.variable}.

    6.  Run [assign slottables for a
        tree](#assign-slotables-for-a-tree){#ref-for-assign-slotables-for-a-tree①
        link-type="dfn"} with `node`{.variable}'s
        [root](#concept-tree-root){#ref-for-concept-tree-root①⑥
        link-type="dfn"}.

    7.  For each [shadow-including inclusive
        descendant](#concept-shadow-including-inclusive-descendant){#ref-for-concept-shadow-including-inclusive-descendant
        link-type="dfn"} `inclusiveDescendant`{.variable} of
        `node`{.variable}, in [shadow-including tree
        order](#concept-shadow-including-tree-order){#ref-for-concept-shadow-including-tree-order
        link-type="dfn"}:

        1.  Run the [insertion
            steps](#concept-node-insert-ext){#ref-for-concept-node-insert-ext③
            link-type="dfn"} with `inclusiveDescendant`{.variable}.

        2.  If `inclusiveDescendant`{.variable} is not
            [connected](#connected){#ref-for-connected link-type="dfn"},
            then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue④
            link-type="dfn"}.

        3.  If `inclusiveDescendant`{.variable} is an
            [element](#concept-element){#ref-for-concept-element①③
            link-type="dfn"}:

            1.  If `inclusiveDescendant`{.variable}'s [custom element
                registry](#element-custom-element-registry){#ref-for-element-custom-element-registry
                link-type="dfn"} is null, then set
                `inclusiveDescendant`{.variable}'s [custom element
                registry](#element-custom-element-registry){#ref-for-element-custom-element-registry①
                link-type="dfn"} to the result of [looking up a custom
                element
                registry](https://html.spec.whatwg.org/multipage/custom-elements.html#look-up-a-custom-element-registry){#ref-for-look-up-a-custom-element-registry
                link-type="dfn"} given
                `inclusiveDescendant`{.variable}'s
                [parent](#concept-tree-parent){#ref-for-concept-tree-parent①①
                link-type="dfn"}.

            2.  Otherwise, if `inclusiveDescendant`{.variable}'s [custom
                element
                registry](#element-custom-element-registry){#ref-for-element-custom-element-registry②
                link-type="dfn"}'s [is
                scoped](https://html.spec.whatwg.org/multipage/custom-elements.html#is-scoped){#ref-for-is-scoped
                link-type="dfn"} is true,
                [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append⑧
                link-type="dfn"} `inclusiveDescendant`{.variable}'s
                [node
                document](#concept-node-document){#ref-for-concept-node-document④
                link-type="dfn"} to `inclusiveDescendant`{.variable}'s
                [custom element
                registry](#element-custom-element-registry){#ref-for-element-custom-element-registry③
                link-type="dfn"}'s [scoped document
                set](https://html.spec.whatwg.org/multipage/custom-elements.html#scoped-document-set){#ref-for-scoped-document-set
                link-type="dfn"}.

            3.  If `inclusiveDescendant`{.variable} is
                [custom](#concept-element-custom){#ref-for-concept-element-custom
                link-type="dfn"}, then [enqueue a custom element
                callback
                reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-callback-reaction){#ref-for-enqueue-a-custom-element-callback-reaction
                link-type="dfn"} with `inclusiveDescendant`{.variable},
                callback name \"`connectedCallback`\", and « ».

            4.  Otherwise, [try to
                upgrade](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-try-upgrade){#ref-for-concept-try-upgrade
                link-type="dfn"} `inclusiveDescendant`{.variable}.

                If this successfully upgrades
                `inclusiveDescendant`{.variable}, its
                `connectedCallback` will be enqueued automatically
                during the [upgrade an
                element](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-upgrade-an-element){#ref-for-concept-upgrade-an-element
                link-type="dfn"} algorithm.

        4.  Otherwise, if `inclusiveDescendant`{.variable} is a [shadow
            root](#concept-shadow-root){#ref-for-concept-shadow-root①④
            link-type="dfn"}:

            1.  If `inclusiveDescendant`{.variable}'s [custom element
                registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry
                link-type="dfn"} is null and
                `inclusiveDescendant`{.variable}'s [keep custom element
                registry
                null](#shadowroot-keep-custom-element-registry-null){#ref-for-shadowroot-keep-custom-element-registry-null
                link-type="dfn"} is false, then set
                `inclusiveDescendant`{.variable}'s [custom element
                registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry①
                link-type="dfn"} to the result of [looking up a custom
                element
                registry](https://html.spec.whatwg.org/multipage/custom-elements.html#look-up-a-custom-element-registry){#ref-for-look-up-a-custom-element-registry①
                link-type="dfn"} given
                `inclusiveDescendant`{.variable}'s
                [host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host③
                link-type="dfn"}.

            2.  Otherwise, if `inclusiveDescendant`{.variable}'s [custom
                element
                registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry②
                link-type="dfn"} is non-null and
                `inclusiveDescendant`{.variable}'s [custom element
                registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry③
                link-type="dfn"}'s [is
                scoped](https://html.spec.whatwg.org/multipage/custom-elements.html#is-scoped){#ref-for-is-scoped①
                link-type="dfn"} is true,
                [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append⑨
                link-type="dfn"} `inclusiveDescendant`{.variable}'s
                [node
                document](#concept-node-document){#ref-for-concept-node-document⑤
                link-type="dfn"} to `inclusiveDescendant`{.variable}'s
                [custom element
                registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry④
                link-type="dfn"}'s [scoped document
                set](https://html.spec.whatwg.org/multipage/custom-elements.html#scoped-document-set){#ref-for-scoped-document-set①
                link-type="dfn"}.

8.  If *suppress observers flag* is unset, then [queue a tree mutation
    record](#queue-a-tree-mutation-record){#ref-for-queue-a-tree-mutation-record①
    link-type="dfn"} for `parent`{.variable} with `nodes`{.variable}, «
    », `previousSibling`{.variable}, and `child`{.variable}.

9.  Run the [children changed
    steps](#concept-node-children-changed-ext){#ref-for-concept-node-children-changed-ext
    link-type="dfn"} for `parent`{.variable}.

10. Let `staticNodeList`{.variable} be a
    [list](https://infra.spec.whatwg.org/#list){#ref-for-list⑨
    link-type="dfn"} of [nodes](#concept-node){#ref-for-concept-node④⑤
    link-type="dfn"}, initially « ».

    We collect all [nodes](#concept-node){#ref-for-concept-node④⑥
    link-type="dfn"} *before* calling the [post-connection
    steps](#concept-node-post-connection-ext){#ref-for-concept-node-post-connection-ext②
    link-type="dfn"} on any one of them, instead of calling the
    [post-connection
    steps](#concept-node-post-connection-ext){#ref-for-concept-node-post-connection-ext③
    link-type="dfn"} *while* we're traversing the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree①①
    link-type="dfn"}. This is because the [post-connection
    steps](#concept-node-post-connection-ext){#ref-for-concept-node-post-connection-ext④
    link-type="dfn"} can modify the tree's structure, making live
    traversal unsafe, possibly leading to the [post-connection
    steps](#concept-node-post-connection-ext){#ref-for-concept-node-post-connection-ext⑤
    link-type="dfn"} being called multiple times on the same
    [node](#boundary-point-node){#ref-for-boundary-point-node①
    link-type="dfn"}.

11. For each `node`{.variable} of `nodes`{.variable}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order①②
    link-type="dfn"}:

    1.  For each [shadow-including inclusive
        descendant](#concept-shadow-including-inclusive-descendant){#ref-for-concept-shadow-including-inclusive-descendant①
        link-type="dfn"} `inclusiveDescendant`{.variable} of
        `node`{.variable}, in [shadow-including tree
        order](#concept-shadow-including-tree-order){#ref-for-concept-shadow-including-tree-order①
        link-type="dfn"},
        [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①⓪
        link-type="dfn"} `inclusiveDescendant`{.variable} to
        `staticNodeList`{.variable}.

12. [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑤
    link-type="dfn"} `node`{.variable} of `staticNodeList`{.variable},
    if `node`{.variable} is [connected](#connected){#ref-for-connected①
    link-type="dfn"}, then run the [post-connection
    steps](#concept-node-post-connection-ext){#ref-for-concept-node-post-connection-ext⑥
    link-type="dfn"} with `node`{.variable}.

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications④
link-type="dfn"} may define [moving steps]{#concept-node-move-ext .dfn
.dfn-paneled dfn-type="dfn" export=""} for all or some
[nodes](#concept-node){#ref-for-concept-node④⑦ link-type="dfn"}. The
algorithm is passed a [node](#concept-node){#ref-for-concept-node④⑧
link-type="dfn"} `movedNode`{.variable}, and a
[node](#concept-node){#ref-for-concept-node④⑨ link-type="dfn"}-or-null
`oldParent`{.variable} as indicated in the [move](#move){#ref-for-move
link-type="dfn"} algorithm below. Like the [insertion
steps](#concept-node-insert-ext){#ref-for-concept-node-insert-ext④
link-type="dfn"}, these steps must not modify the [node
tree](#concept-node-tree){#ref-for-concept-node-tree①② link-type="dfn"}
that `movedNode`{.variable}
[participates](#concept-tree-participate){#ref-for-concept-tree-participate⑧
link-type="dfn"} in, create [browsing
contexts](https://html.spec.whatwg.org/multipage/document-sequences.html#browsing-context){#ref-for-browsing-context②
link-type="dfn"}, [fire
events](#concept-event-fire){#ref-for-concept-event-fire⑦
link-type="dfn"}, or otherwise execute JavaScript. These steps may queue
tasks to do these things asynchronously, however.

To [move]{#move .dfn .dfn-paneled dfn-type="dfn" noexport=""} a
[node](#concept-node){#ref-for-concept-node⑤⓪ link-type="dfn"}
`node`{.variable} into a [node](#concept-node){#ref-for-concept-node⑤①
link-type="dfn"} `newParent`{.variable} before a
[node](#concept-node){#ref-for-concept-node⑤② link-type="dfn"}-or-null
`child`{.variable}:

1.  If `newParent`{.variable}'s [shadow-including
    root](#concept-shadow-including-root){#ref-for-concept-shadow-including-root①
    link-type="dfn"} is not the same as `node`{.variable}'s
    [shadow-including
    root](#concept-shadow-including-root){#ref-for-concept-shadow-including-root②
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①③
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⑨
    link-type="idl"}.

    This has the side effect of ensuring that a move is only performed
    if `newParent`{.variable}'s
    [connected](#connected){#ref-for-connected② link-type="dfn"} is
    `node`{.variable}'s [connected](#connected){#ref-for-connected③
    link-type="dfn"}.

2.  If `node`{.variable} is a [host-including inclusive
    ancestor](#concept-tree-host-including-inclusive-ancestor){#ref-for-concept-tree-host-including-inclusive-ancestor①
    link-type="dfn"} of `newParent`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①④
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②⓪
    link-type="idl"}.

3.  If `child`{.variable} is non-null and its
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①②
    link-type="dfn"} is not `newParent`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑤
    link-type="dfn"} a
    \"[`NotFoundError`{.idl}](https://webidl.spec.whatwg.org/#notfounderror){#ref-for-notfounderror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②①
    link-type="idl"}.

4.  If `node`{.variable} is not an
    [`Element`{.idl}](#element){#ref-for-element①⑤ link-type="idl"} or a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata⑥
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑤③
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑥
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror⑦
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②②
    link-type="idl"}.

5.  If `node`{.variable} is a [`Text`{.idl}](#text){#ref-for-text①⓪
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑤④
    link-type="dfn"} and `newParent`{.variable} is a
    [document](#concept-document){#ref-for-concept-document①①
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑦
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror⑧
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②③
    link-type="idl"}.

6.  If `newParent`{.variable} is a
    [document](#concept-document){#ref-for-concept-document①②
    link-type="dfn"}, `node`{.variable} is an
    [`Element`{.idl}](#element){#ref-for-element①⑥ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node⑤⑤ link-type="dfn"}, and
    either `newParent`{.variable} has an
    [element](#concept-element){#ref-for-concept-element①④
    link-type="dfn"}
    [child](#concept-tree-child){#ref-for-concept-tree-child②⑤
    link-type="dfn"}, `child`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype⑦
    link-type="dfn"}, or `child`{.variable} is non-null and a
    [doctype](#concept-doctype){#ref-for-concept-doctype⑧
    link-type="dfn"} is
    [following](#concept-tree-following){#ref-for-concept-tree-following④
    link-type="dfn"} `child`{.variable} then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑧
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror⑨
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②④
    link-type="idl"}.

7.  Let `oldParent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①③
    link-type="dfn"}.

8.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert①
    link-type="dfn"}: `oldParent`{.variable} is non-null.

9.  Run the [live range pre-remove
    steps](#live-range-pre-remove-steps){#ref-for-live-range-pre-remove-steps
    link-type="dfn"}, given `node`{.variable}.

10. For each [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator
    link-type="idl"} object `iterator`{.variable} whose
    [root](#concept-traversal-root){#ref-for-concept-traversal-root
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑥
    link-type="dfn"} is `node`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑦
    link-type="dfn"}, run the [`NodeIterator` pre-remove
    steps](#nodeiterator-pre-removing-steps){#ref-for-nodeiterator-pre-removing-steps
    link-type="dfn"} given `node`{.variable} and `iterator`{.variable}.

11. Let `oldPreviousSibling`{.variable} be `node`{.variable}'s [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①
    link-type="dfn"}.

12. Let `oldNextSibling`{.variable} be `node`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①
    link-type="dfn"}.

13. [Remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove②
    link-type="dfn"} `node`{.variable} from `oldParent`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child②⑥
    link-type="dfn"}.

14. If `node`{.variable} is
    [assigned](#slotable-assigned){#ref-for-slotable-assigned③
    link-type="dfn"}, then run [assign
    slottables](#assign-slotables){#ref-for-assign-slotables③
    link-type="dfn"} for `node`{.variable}'s [assigned
    slot](#slotable-assigned-slot){#ref-for-slotable-assigned-slot③
    link-type="dfn"}.

15. If `oldParent`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root①⑦
    link-type="dfn"} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root①⑤
    link-type="dfn"}, and `oldParent`{.variable} is a
    [slot](#concept-slot){#ref-for-concept-slot②① link-type="dfn"} whose
    [assigned nodes](#slot-assigned-nodes){#ref-for-slot-assigned-nodes③
    link-type="dfn"} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty①
    link-type="dfn"}, then run [signal a slot
    change](#signal-a-slot-change){#ref-for-signal-a-slot-change②
    link-type="dfn"} for `oldParent`{.variable}.

16. If `node`{.variable} has an [inclusive
    descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant①
    link-type="dfn"} that is a
    [slot](#concept-slot){#ref-for-concept-slot②② link-type="dfn"}:

    1.  Run [assign slottables for a
        tree](#assign-slotables-for-a-tree){#ref-for-assign-slotables-for-a-tree②
        link-type="dfn"} with `oldParent`{.variable}'s
        [root](#concept-tree-root){#ref-for-concept-tree-root①⑧
        link-type="dfn"}.

    2.  Run [assign slottables for a
        tree](#assign-slotables-for-a-tree){#ref-for-assign-slotables-for-a-tree③
        link-type="dfn"} with `node`{.variable}.

17. If `child`{.variable} is non-null:

    1.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range②
        link-type="dfn"} whose [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node①
        link-type="dfn"} is `newParent`{.variable} and [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②
        link-type="dfn"} is greater than `child`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index③
        link-type="dfn"}, increase its [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset③
        link-type="dfn"} by 1.

    2.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range③
        link-type="dfn"} whose [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node①
        link-type="dfn"} is `newParent`{.variable} and [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset②
        link-type="dfn"} is greater than `child`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index④
        link-type="dfn"}, increase its [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset③
        link-type="dfn"} by 1.

18. Let `newPreviousSibling`{.variable} be `child`{.variable}'s
    [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling②
    link-type="dfn"} if `child`{.variable} is non-null, and
    `newParent`{.variable}'s [last
    child](#concept-tree-last-child){#ref-for-concept-tree-last-child①
    link-type="dfn"} otherwise.

19. If `child`{.variable} is null, then
    [append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append①⓪
    link-type="dfn"} `node`{.variable} to `newParent`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child②⑦
    link-type="dfn"}.

20. Otherwise,
    [insert](https://infra.spec.whatwg.org/#list-insert){#ref-for-list-insert①
    link-type="dfn"} `node`{.variable} into `newParent`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child②⑧
    link-type="dfn"} before `child`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index⑤
    link-type="dfn"}.

21. If `newParent`{.variable} is a [shadow
    host](#element-shadow-host){#ref-for-element-shadow-host①
    link-type="dfn"} whose [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root①⑥
    link-type="dfn"}'s [slot
    assignment](#shadowroot-slot-assignment){#ref-for-shadowroot-slot-assignment③
    link-type="dfn"} is \"`named`\" and `node`{.variable} is a
    [slottable](#concept-slotable){#ref-for-concept-slotable①⑦
    link-type="dfn"}, then [assign a
    slot](#assign-a-slot){#ref-for-assign-a-slot② link-type="dfn"} for
    `node`{.variable}.

22. If `newParent`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root①⑨
    link-type="dfn"} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root①⑦
    link-type="dfn"}, and `newParent`{.variable} is a
    [slot](#concept-slot){#ref-for-concept-slot②③ link-type="dfn"} whose
    [assigned nodes](#slot-assigned-nodes){#ref-for-slot-assigned-nodes④
    link-type="dfn"} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty②
    link-type="dfn"}, then run [signal a slot
    change](#signal-a-slot-change){#ref-for-signal-a-slot-change③
    link-type="dfn"} for `newParent`{.variable}.

23. Run [assign slottables for a
    tree](#assign-slotables-for-a-tree){#ref-for-assign-slotables-for-a-tree④
    link-type="dfn"} with `node`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root②⓪
    link-type="dfn"}.

24. For each [shadow-including inclusive
    descendant](#concept-shadow-including-inclusive-descendant){#ref-for-concept-shadow-including-inclusive-descendant②
    link-type="dfn"} `inclusiveDescendant`{.variable} of
    `node`{.variable}, in [shadow-including tree
    order](#concept-shadow-including-tree-order){#ref-for-concept-shadow-including-tree-order②
    link-type="dfn"}:

    1.  If `inclusiveDescendant`{.variable} is `node`{.variable}, then
        run the [moving
        steps](#concept-node-move-ext){#ref-for-concept-node-move-ext
        link-type="dfn"} with `inclusiveDescendant`{.variable} and
        `oldParent`{.variable}. Otherwise, run the [moving
        steps](#concept-node-move-ext){#ref-for-concept-node-move-ext①
        link-type="dfn"} with `inclusiveDescendant`{.variable} and null.

        Because the [move](#move){#ref-for-move① link-type="dfn"}
        algorithm is a separate primitive from
        [insert](#concept-node-insert){#ref-for-concept-node-insert⑥
        link-type="dfn"} and
        [remove](#concept-node-remove){#ref-for-concept-node-remove②
        link-type="dfn"}, it does not invoke the traditional [insertion
        steps](#concept-node-insert-ext){#ref-for-concept-node-insert-ext⑤
        link-type="dfn"} or [removing
        steps](#concept-node-remove-ext){#ref-for-concept-node-remove-ext
        link-type="dfn"} for `inclusiveDescendant`{.variable}.

    2.  If `inclusiveDescendant`{.variable} is
        [custom](#concept-element-custom){#ref-for-concept-element-custom①
        link-type="dfn"} and `newParent`{.variable} is
        [connected](#connected){#ref-for-connected④ link-type="dfn"},
        then [enqueue a custom element callback
        reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-callback-reaction){#ref-for-enqueue-a-custom-element-callback-reaction①
        link-type="dfn"} with `inclusiveDescendant`{.variable}, callback
        name \"`connectedMoveCallback`\", and « ».

25. [Queue a tree mutation
    record](#queue-a-tree-mutation-record){#ref-for-queue-a-tree-mutation-record②
    link-type="dfn"} for `oldParent`{.variable} with « », «
    `node`{.variable} », `oldPreviousSibling`{.variable}, and
    `oldNextSibling`{.variable}.

26. [Queue a tree mutation
    record](#queue-a-tree-mutation-record){#ref-for-queue-a-tree-mutation-record③
    link-type="dfn"} for `newParent`{.variable} with « `node`{.variable}
    », « », `newPreviousSibling`{.variable}, and `child`{.variable}.

To [append]{#concept-node-append .dfn .dfn-paneled dfn-type="dfn"
export=""} a `node`{.variable} to a `parent`{.variable},
[pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert
link-type="dfn"} `node`{.variable} into `parent`{.variable} before null.

To [replace]{#concept-node-replace .dfn .dfn-paneled dfn-type="dfn"
export=""} a `child`{.variable} with `node`{.variable} within a
`parent`{.variable}, run these steps:

1.  If `parent`{.variable} is not a
    [`Document`{.idl}](#document){#ref-for-document④ link-type="idl"},
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment⑨
    link-type="idl"}, or [`Element`{.idl}](#element){#ref-for-element①⑦
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑤⑥
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw①⑨
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②⑤
    link-type="idl"}.

2.  If `node`{.variable} is a [host-including inclusive
    ancestor](#concept-tree-host-including-inclusive-ancestor){#ref-for-concept-tree-host-including-inclusive-ancestor②
    link-type="dfn"} of `parent`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⓪
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②⑥
    link-type="idl"}.

3.  If `child`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①④
    link-type="dfn"} is not `parent`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②①
    link-type="dfn"} a
    \"[`NotFoundError`{.idl}](https://webidl.spec.whatwg.org/#notfounderror){#ref-for-notfounderror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②⑦
    link-type="idl"}.

4.  If `node`{.variable} is not a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①⓪
    link-type="idl"},
    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype⑦
    link-type="idl"}, [`Element`{.idl}](#element){#ref-for-element①⑧
    link-type="idl"}, or
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata⑦
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑤⑦
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②②
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②⑧
    link-type="idl"}.

5.  If either `node`{.variable} is a
    [`Text`{.idl}](#text){#ref-for-text①① link-type="idl"}
    [node](#concept-node){#ref-for-concept-node⑤⑧ link-type="dfn"} and
    `parent`{.variable} is a
    [document](#concept-document){#ref-for-concept-document①③
    link-type="dfn"}, or `node`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype⑨
    link-type="dfn"} and `parent`{.variable} is not a
    [document](#concept-document){#ref-for-concept-document①④
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②③
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException②⑨
    link-type="idl"}.

6.  If `parent`{.variable} is a
    [document](#concept-document){#ref-for-concept-document①⑤
    link-type="dfn"}, and any of the statements below, switched on the
    interface `node`{.variable}
    [implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements④
    link-type="dfn"}, are true, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②④
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③⓪
    link-type="idl"}.

    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①① link-type="idl"}

    :   If `node`{.variable} has more than one
        [element](#concept-element){#ref-for-concept-element①⑤
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child②⑨
        link-type="dfn"} or has a [`Text`{.idl}](#text){#ref-for-text①②
        link-type="idl"} [node](#concept-node){#ref-for-concept-node⑤⑨
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child③⓪
        link-type="dfn"}.

        Otherwise, if `node`{.variable} has one
        [element](#concept-element){#ref-for-concept-element①⑥
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child③①
        link-type="dfn"} and either `parent`{.variable} has an
        [element](#concept-element){#ref-for-concept-element①⑦
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child③②
        link-type="dfn"} that is not `child`{.variable} or a
        [doctype](#concept-doctype){#ref-for-concept-doctype①⓪
        link-type="dfn"} is
        [following](#concept-tree-following){#ref-for-concept-tree-following⑤
        link-type="dfn"} `child`{.variable}.

    [`Element`{.idl}](#element){#ref-for-element①⑨ link-type="idl"}

    :   `parent`{.variable} has an
        [element](#concept-element){#ref-for-concept-element①⑧
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child③③
        link-type="dfn"} that is not `child`{.variable} or a
        [doctype](#concept-doctype){#ref-for-concept-doctype①①
        link-type="dfn"} is
        [following](#concept-tree-following){#ref-for-concept-tree-following⑥
        link-type="dfn"} `child`{.variable}.

    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype⑧ link-type="idl"}

    :   `parent`{.variable} has a
        [doctype](#concept-doctype){#ref-for-concept-doctype①②
        link-type="dfn"}
        [child](#concept-tree-child){#ref-for-concept-tree-child③④
        link-type="dfn"} that is not `child`{.variable}, or an
        [element](#concept-element){#ref-for-concept-element①⑨
        link-type="dfn"} is
        [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding④
        link-type="dfn"} `child`{.variable}.

    The above statements differ from the
    [pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert①
    link-type="dfn"} algorithm.

7.  Let `referenceChild`{.variable} be `child`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling②
    link-type="dfn"}.

8.  If `referenceChild`{.variable} is `node`{.variable}, then set
    `referenceChild`{.variable} to `node`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling③
    link-type="dfn"}.

9.  Let `previousSibling`{.variable} be `child`{.variable}'s [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling③
    link-type="dfn"}.

10. Let `removedNodes`{.variable} be the empty set.

11. If `child`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①⑤
    link-type="dfn"} is non-null:

    1.  Set `removedNodes`{.variable} to « `child`{.variable} ».

    2.  [Remove](#concept-node-remove){#ref-for-concept-node-remove③
        link-type="dfn"} `child`{.variable} with the *suppress observers
        flag* set.

    The above can only be false if `child`{.variable} is
    `node`{.variable}.

12. Let `nodes`{.variable} be `node`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child③⑤
    link-type="dfn"} if `node`{.variable} is a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①②
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑥⓪
    link-type="dfn"}; otherwise « `node`{.variable} ».

13. [Insert](#concept-node-insert){#ref-for-concept-node-insert⑦
    link-type="dfn"} `node`{.variable} into `parent`{.variable} before
    `referenceChild`{.variable} with the *suppress observers flag* set.

14. [Queue a tree mutation
    record](#queue-a-tree-mutation-record){#ref-for-queue-a-tree-mutation-record④
    link-type="dfn"} for `parent`{.variable} with `nodes`{.variable},
    `removedNodes`{.variable}, `previousSibling`{.variable}, and
    `referenceChild`{.variable}.

15. Return `child`{.variable}.

To [replace all]{#concept-node-replace-all .dfn .dfn-paneled
dfn-for="Node" dfn-type="dfn" export=""} with a `node`{.variable} within
a `parent`{.variable}, run these steps:

1.  Let `removedNodes`{.variable} be `parent`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child③⑥
    link-type="dfn"}.

2.  Let `addedNodes`{.variable} be the empty set.

3.  If `node`{.variable} is a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①③
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑥①
    link-type="dfn"}, then set `addedNodes`{.variable} to
    `node`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child③⑦
    link-type="dfn"}.

4.  Otherwise, if `node`{.variable} is non-null, set
    `addedNodes`{.variable} to « `node`{.variable} ».

5.  [Remove](#concept-node-remove){#ref-for-concept-node-remove④
    link-type="dfn"} all `parent`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child③⑧
    link-type="dfn"}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order①③
    link-type="dfn"}, with the *suppress observers flag* set.

6.  If `node`{.variable} is non-null, then
    [insert](#concept-node-insert){#ref-for-concept-node-insert⑧
    link-type="dfn"} `node`{.variable} into `parent`{.variable} before
    null with the *suppress observers flag* set.

7.  If either `addedNodes`{.variable} or `removedNodes`{.variable} [is
    not
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty③
    link-type="dfn"}, then [queue a tree mutation
    record](#queue-a-tree-mutation-record){#ref-for-queue-a-tree-mutation-record⑤
    link-type="dfn"} for `parent`{.variable} with
    `addedNodes`{.variable}, `removedNodes`{.variable}, null, and null.

This algorithm does not make any checks with regards to the [node
tree](#concept-node-tree){#ref-for-concept-node-tree①③ link-type="dfn"}
constraints. Specification authors need to use it wisely.

To [pre-remove]{#concept-node-pre-remove .dfn .dfn-paneled
dfn-type="dfn" export=""} a `child`{.variable} from a
`parent`{.variable}, run these steps:

1.  If `child`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①⑥
    link-type="dfn"} is not `parent`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⑤
    link-type="dfn"} a
    \"[`NotFoundError`{.idl}](https://webidl.spec.whatwg.org/#notfounderror){#ref-for-notfounderror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③①
    link-type="idl"}.

2.  [Remove](#concept-node-remove){#ref-for-concept-node-remove⑤
    link-type="dfn"} `child`{.variable}.

3.  Return `child`{.variable}.

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications⑤
link-type="dfn"} may define [removing steps]{#concept-node-remove-ext
.dfn .dfn-paneled dfn-type="dfn" export=""} for all or some
[nodes](#concept-node){#ref-for-concept-node⑥② link-type="dfn"}. The
algorithm is passed a [node](#concept-node){#ref-for-concept-node⑥③
link-type="dfn"} `removedNode`{.variable} and a
[node](#concept-node){#ref-for-concept-node⑥④ link-type="dfn"}-or-null
`oldParent`{.variable}, as indicated in the
[remove](#concept-node-remove){#ref-for-concept-node-remove⑥
link-type="dfn"} algorithm below.

To [remove]{#concept-node-remove .dfn .dfn-paneled dfn-type="dfn"
export=""} a [node](#concept-node){#ref-for-concept-node⑥⑤
link-type="dfn"} `node`{.variable}, with an optional *suppress observers
flag*, run these steps:

1.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①⑦
    link-type="dfn"}.

2.  Assert: `parent`{.variable} is non-null.

3.  Run the [live range pre-remove
    steps](#live-range-pre-remove-steps){#ref-for-live-range-pre-remove-steps①
    link-type="dfn"}, given `node`{.variable}.

4.  For each
    [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①
    link-type="idl"} object `iterator`{.variable} whose
    [root](#concept-traversal-root){#ref-for-concept-traversal-root①
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑧
    link-type="dfn"} is `node`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑨
    link-type="dfn"}, run the [`NodeIterator` pre-remove
    steps](#nodeiterator-pre-removing-steps){#ref-for-nodeiterator-pre-removing-steps①
    link-type="dfn"} given `node`{.variable} and `iterator`{.variable}.

5.  Let `oldPreviousSibling`{.variable} be `node`{.variable}'s [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling④
    link-type="dfn"}.

6.  Let `oldNextSibling`{.variable} be `node`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling④
    link-type="dfn"}.

7.  [Remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove③
    link-type="dfn"} `node`{.variable} from its `parent`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child③⑨
    link-type="dfn"}.

8.  If `node`{.variable} is
    [assigned](#slotable-assigned){#ref-for-slotable-assigned④
    link-type="dfn"}, then run [assign
    slottables](#assign-slotables){#ref-for-assign-slotables④
    link-type="dfn"} for `node`{.variable}'s [assigned
    slot](#slotable-assigned-slot){#ref-for-slotable-assigned-slot④
    link-type="dfn"}.

9.  If `parent`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root②①
    link-type="dfn"} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root①⑧
    link-type="dfn"}, and `parent`{.variable} is a
    [slot](#concept-slot){#ref-for-concept-slot②④ link-type="dfn"} whose
    [assigned nodes](#slot-assigned-nodes){#ref-for-slot-assigned-nodes⑤
    link-type="dfn"} is the empty list, then run [signal a slot
    change](#signal-a-slot-change){#ref-for-signal-a-slot-change④
    link-type="dfn"} for `parent`{.variable}.

10. If `node`{.variable} has an [inclusive
    descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant②
    link-type="dfn"} that is a
    [slot](#concept-slot){#ref-for-concept-slot②⑤ link-type="dfn"}:

    1.  Run [assign slottables for a
        tree](#assign-slotables-for-a-tree){#ref-for-assign-slotables-for-a-tree⑤
        link-type="dfn"} with `parent`{.variable}'s
        [root](#concept-tree-root){#ref-for-concept-tree-root②②
        link-type="dfn"}.

    2.  Run [assign slottables for a
        tree](#assign-slotables-for-a-tree){#ref-for-assign-slotables-for-a-tree⑥
        link-type="dfn"} with `node`{.variable}.

11. Run the [removing
    steps](#concept-node-remove-ext){#ref-for-concept-node-remove-ext①
    link-type="dfn"} with `node`{.variable} and `parent`{.variable}.

12. Let `isParentConnected`{.variable} be `parent`{.variable}'s
    [connected](#connected){#ref-for-connected⑤ link-type="dfn"}.

13. If `node`{.variable} is
    [custom](#concept-element-custom){#ref-for-concept-element-custom②
    link-type="dfn"} and `isParentConnected`{.variable} is true, then
    [enqueue a custom element callback
    reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-callback-reaction){#ref-for-enqueue-a-custom-element-callback-reaction②
    link-type="dfn"} with `node`{.variable}, callback name
    \"`disconnectedCallback`\", and « ».

    It is intentional for now that
    [custom](#concept-element-custom){#ref-for-concept-element-custom③
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element②⓪
    link-type="dfn"} do not get `parent`{.variable} passed. This might
    change in the future if there is a need.

14. For each [shadow-including
    descendant](#concept-shadow-including-descendant){#ref-for-concept-shadow-including-descendant
    link-type="dfn"} `descendant`{.variable} of `node`{.variable}, in
    [shadow-including tree
    order](#concept-shadow-including-tree-order){#ref-for-concept-shadow-including-tree-order③
    link-type="dfn"}:

    1.  Run the [removing
        steps](#concept-node-remove-ext){#ref-for-concept-node-remove-ext②
        link-type="dfn"} with `descendant`{.variable} and null.

    2.  If `descendant`{.variable} is
        [custom](#concept-element-custom){#ref-for-concept-element-custom④
        link-type="dfn"} and `isParentConnected`{.variable} is true,
        then [enqueue a custom element callback
        reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-callback-reaction){#ref-for-enqueue-a-custom-element-callback-reaction③
        link-type="dfn"} with `descendant`{.variable}, callback name
        \"`disconnectedCallback`\", and « ».

15. For each [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②
    link-type="dfn"} `inclusiveAncestor`{.variable} of
    `parent`{.variable}, and then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑥
    link-type="dfn"} `registered`{.variable} of
    `inclusiveAncestor`{.variable}'s [registered observer
    list](#registered-observer-list){#ref-for-registered-observer-list
    link-type="dfn"}, if `registered`{.variable}'s
    [options](#registered-observer-options){#ref-for-registered-observer-options
    link-type="dfn"}\[\"[`subtree`{.idl}](#dom-mutationobserverinit-subtree){#ref-for-dom-mutationobserverinit-subtree
    link-type="idl"}\"\] is true, then
    [append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①①
    link-type="dfn"} a new [transient registered
    observer](#transient-registered-observer){#ref-for-transient-registered-observer
    link-type="dfn"} whose
    [observer](#registered-observer-observer){#ref-for-registered-observer-observer
    link-type="dfn"} is `registered`{.variable}'s
    [observer](#registered-observer-observer){#ref-for-registered-observer-observer①
    link-type="dfn"},
    [options](#registered-observer-options){#ref-for-registered-observer-options①
    link-type="dfn"} is `registered`{.variable}'s
    [options](#registered-observer-options){#ref-for-registered-observer-options②
    link-type="dfn"}, and
    [source](#transient-registered-observer-source){#ref-for-transient-registered-observer-source
    link-type="dfn"} is `registered`{.variable} to `node`{.variable}'s
    [registered observer
    list](#registered-observer-list){#ref-for-registered-observer-list①
    link-type="dfn"}.

16. If *suppress observers flag* is unset, then [queue a tree mutation
    record](#queue-a-tree-mutation-record){#ref-for-queue-a-tree-mutation-record⑥
    link-type="dfn"} for `parent`{.variable} with « », «
    `node`{.variable} », `oldPreviousSibling`{.variable}, and
    `oldNextSibling`{.variable}.

17. Run the [children changed
    steps](#concept-node-children-changed-ext){#ref-for-concept-node-children-changed-ext①
    link-type="dfn"} for `parent`{.variable}.

#### [4.2.4. ]{.secno}[Mixin [`NonElementParentNode`{.idl}](#nonelementparentnode){#ref-for-nonelementparentnode link-type="idl"}]{.content}[](#interface-nonelementparentnode){.self-link} {#interface-nonelementparentnode .heading .settled level="4.2.4"}

Web compatibility prevents the
[`getElementById()`{.idl}](#dom-nonelementparentnode-getelementbyid){#ref-for-dom-nonelementparentnode-getelementbyid
link-type="idl"} method from being exposed on
[elements](#concept-element){#ref-for-concept-element②① link-type="dfn"}
(and therefore on [`ParentNode`{.idl}](#parentnode){#ref-for-parentnode
link-type="idl"}).

``` {.idl .highlight .def}
interface mixin NonElementParentNode {
  Element? getElementById(DOMString elementId);
};
Document includes NonElementParentNode;
DocumentFragment includes NonElementParentNode;
```

`node`{.variable}` . `[`getElementById`](#dom-nonelementparentnode-getelementbyid){#ref-for-dom-nonelementparentnode-getelementbyid② .idl-code link-type="method"}`(``elementId`{.variable}`)`

:   Returns the first
    [element](#concept-element){#ref-for-concept-element②②
    link-type="dfn"} within `node`{.variable}'s
    [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant⑤
    link-type="dfn"} whose [ID](#concept-id){#ref-for-concept-id
    link-type="dfn"} is `elementId`{.variable}.

The
[`getElementById(``elementId`{.variable}`)`]{#dom-nonelementparentnode-getelementbyid
.dfn .dfn-paneled .idl-code dfn-for="NonElementParentNode"
dfn-type="method" export=""} method steps are to return the first
[element](#concept-element){#ref-for-concept-element②③ link-type="dfn"},
in [tree order](#concept-tree-order){#ref-for-concept-tree-order①④
link-type="dfn"}, within
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③②
link-type="dfn"}'s
[descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant⑥
link-type="dfn"}, whose [ID](#concept-id){#ref-for-concept-id①
link-type="dfn"} is `elementId`{.variable}; otherwise, if there is no
such [element](#concept-element){#ref-for-concept-element②④
link-type="dfn"}, null.

#### [4.2.5. ]{.secno}[Mixin [`DocumentOrShadowRoot`{.idl}](#documentorshadowroot){#ref-for-documentorshadowroot link-type="idl"}]{.content}[](#mixin-documentorshadowroot){.self-link} {#mixin-documentorshadowroot .heading .settled level="4.2.5"}

``` {.idl .highlight .def}
interface mixin DocumentOrShadowRoot {
  readonly attribute CustomElementRegistry? customElementRegistry;
};
Document includes DocumentOrShadowRoot;
ShadowRoot includes DocumentOrShadowRoot;
```

`registry`{.variable}` = ``documentOrShadowRoot`{.variable}` . `[`customElementRegistry`{.idl}](#dom-documentorshadowroot-customelementregistry){#ref-for-dom-documentorshadowroot-customelementregistry① link-type="idl"}

:   Returns `documentOrShadowRoot`{.variable}'s
    [`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry①
    link-type="idl"} object, if any; otherwise null.

The
[`customElementRegistry`]{#dom-documentorshadowroot-customelementregistry
.dfn .dfn-paneled .idl-code dfn-for="DocumentOrShadowRoot"
dfn-type="attribute" export=""} getter steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③③
    link-type="dfn"} is a
    [document](#concept-document){#ref-for-concept-document①⑥
    link-type="dfn"}, then return
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③④
    link-type="dfn"}'s [custom element
    registry](#document-custom-element-registry){#ref-for-document-custom-element-registry
    link-type="dfn"}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert②
    link-type="dfn"}:
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑤
    link-type="dfn"} is a
    [`ShadowRoot`{.idl}](#shadowroot){#ref-for-shadowroot④
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑥⑥
    link-type="dfn"}.

3.  Return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑥
    link-type="dfn"}'s [custom element
    registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry⑤
    link-type="dfn"}.

The
[`DocumentOrShadowRoot`{.idl}](#documentorshadowroot){#ref-for-documentorshadowroot③
link-type="idl"} mixin is also expected to be used by other standards
that want to define APIs shared between
[documents](#concept-document){#ref-for-concept-document①⑦
link-type="dfn"} and [shadow
roots](#concept-shadow-root){#ref-for-concept-shadow-root①⑨
link-type="dfn"}.

#### [4.2.6. ]{.secno}[Mixin [`ParentNode`{.idl}](#parentnode){#ref-for-parentnode① link-type="idl"}]{.content}[](#interface-parentnode){.self-link} {#interface-parentnode .heading .settled level="4.2.6"}

To [convert nodes into a node]{#converting-nodes-into-a-node .dfn
.dfn-paneled dfn-type="dfn" export=""
lt="converting nodes into a node"}, given `nodes`{.variable} and
`document`{.variable}, run these steps:

1.  Let `node`{.variable} be null.

2.  Replace each string in `nodes`{.variable} with a new
    [`Text`{.idl}](#text){#ref-for-text①③ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node⑥⑦ link-type="dfn"} whose
    [data](#concept-cd-data){#ref-for-concept-cd-data① link-type="dfn"}
    is the string and [node
    document](#concept-node-document){#ref-for-concept-node-document①⓪
    link-type="dfn"} is `document`{.variable}.

3.  If `nodes`{.variable} contains one
    [node](#concept-node){#ref-for-concept-node⑥⑧ link-type="dfn"}, then
    set `node`{.variable} to `nodes`{.variable}\[0\].

4.  Otherwise, set `node`{.variable} to a new
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①⑤
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑥⑨
    link-type="dfn"} whose [node
    document](#concept-node-document){#ref-for-concept-node-document①①
    link-type="dfn"} is `document`{.variable}, and then
    [append](#concept-node-append){#ref-for-concept-node-append
    link-type="dfn"} each [node](#concept-node){#ref-for-concept-node⑦⓪
    link-type="dfn"} in `nodes`{.variable}, if any, to it.

5.  Return `node`{.variable}.

``` {.idl .highlight .def}
interface mixin ParentNode {
  [SameObject] readonly attribute HTMLCollection children;
  readonly attribute Element? firstElementChild;
  readonly attribute Element? lastElementChild;
  readonly attribute unsigned long childElementCount;

  [CEReactions, Unscopable] undefined prepend((Node or DOMString)... nodes);
  [CEReactions, Unscopable] undefined append((Node or DOMString)... nodes);
  [CEReactions, Unscopable] undefined replaceChildren((Node or DOMString)... nodes);

  [CEReactions] undefined moveBefore(Node node, Node? child);

  Element? querySelector(DOMString selectors);
  [NewObject] NodeList querySelectorAll(DOMString selectors);
};
Document includes ParentNode;
DocumentFragment includes ParentNode;
Element includes ParentNode;
```

`collection`{.variable}` = ``node`{.variable}` . `[`children`{.idl}](#dom-parentnode-children){#ref-for-dom-parentnode-children① link-type="idl"}
:   Returns the
    [child](#concept-tree-child){#ref-for-concept-tree-child④⓪
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element②⑤
    link-type="dfn"}.

`element`{.variable}` = ``node`{.variable}` . `[`firstElementChild`{.idl}](#dom-parentnode-firstelementchild){#ref-for-dom-parentnode-firstelementchild① link-type="idl"}
:   Returns the first
    [child](#concept-tree-child){#ref-for-concept-tree-child④①
    link-type="dfn"} that is an
    [element](#concept-element){#ref-for-concept-element②⑥
    link-type="dfn"}; otherwise null.

`element`{.variable}` = ``node`{.variable}` . `[`lastElementChild`{.idl}](#dom-parentnode-lastelementchild){#ref-for-dom-parentnode-lastelementchild① link-type="idl"}
:   Returns the last
    [child](#concept-tree-child){#ref-for-concept-tree-child④②
    link-type="dfn"} that is an
    [element](#concept-element){#ref-for-concept-element②⑦
    link-type="dfn"}; otherwise null.

`node`{.variable}` . `[`prepend`](#dom-parentnode-prepend){#ref-for-dom-parentnode-prepend① .idl-code link-type="method"}`(``nodes`{.variable}`)`

:   Inserts `nodes`{.variable} before the [first
    child](#concept-tree-first-child){#ref-for-concept-tree-first-child
    link-type="dfn"} of `node`{.variable}, while replacing strings in
    `nodes`{.variable} with equivalent
    [`Text`{.idl}](#text){#ref-for-text①④ link-type="idl"}
    [nodes](#concept-node){#ref-for-concept-node⑦① link-type="dfn"}.

    [Throws](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⑥
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③②
    link-type="idl"} if the constraints of the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree①④
    link-type="dfn"} are violated.

`node`{.variable}` . `[`append`](#dom-parentnode-append){#ref-for-dom-parentnode-append① .idl-code link-type="method"}`(``nodes`{.variable}`)`

:   Inserts `nodes`{.variable} after the [last
    child](#concept-tree-last-child){#ref-for-concept-tree-last-child②
    link-type="dfn"} of `node`{.variable}, while replacing strings in
    `nodes`{.variable} with equivalent
    [`Text`{.idl}](#text){#ref-for-text①⑤ link-type="idl"}
    [nodes](#concept-node){#ref-for-concept-node⑦② link-type="dfn"}.

    [Throws](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⑦
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③③
    link-type="idl"} if the constraints of the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree①⑤
    link-type="dfn"} are violated.

`node`{.variable}` . `[`replaceChildren`](#dom-parentnode-replacechildren){#ref-for-dom-parentnode-replacechildren① .idl-code link-type="method"}`(``nodes`{.variable}`)`

:   Replace all
    [children](#concept-tree-child){#ref-for-concept-tree-child④③
    link-type="dfn"} of `node`{.variable} with `nodes`{.variable}, while
    replacing strings in `nodes`{.variable} with equivalent
    [`Text`{.idl}](#text){#ref-for-text①⑥ link-type="idl"}
    [nodes](#concept-node){#ref-for-concept-node⑦③ link-type="dfn"}.

    [Throws](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⑧
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①⑦
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③④
    link-type="idl"} if the constraints of the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree①⑥
    link-type="dfn"} are violated.

`node`{.variable}` . `[`moveBefore`](#dom-parentnode-movebefore){#ref-for-dom-parentnode-movebefore① .idl-code link-type="method"}`(``movedNode`{.variable}`, ``child`{.variable}`)`

:   Moves, without first removing, `movedNode`{.variable} into
    `node`{.variable} after `child`{.variable} if `child`{.variable} is
    non-null; otherwise after the [last
    child](#concept-tree-last-child){#ref-for-concept-tree-last-child③
    link-type="dfn"} of `node`{.variable}. This method preserves state
    associated with `movedNode`{.variable}.

    [Throws](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw②⑨
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①⑧
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③⑤
    link-type="idl"} if the constraints of the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree①⑦
    link-type="dfn"} are violated, or the state associated with the
    moved node cannot be preserved.

`node`{.variable}` . `[`querySelector`](#dom-parentnode-queryselector){#ref-for-dom-parentnode-queryselector① .idl-code link-type="method"}`(``selectors`{.variable}`)`

:   Returns the first
    [element](#concept-element){#ref-for-concept-element②⑧
    link-type="dfn"} that is a
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant⑦
    link-type="dfn"} of `node`{.variable} that matches
    `selectors`{.variable}.

`node`{.variable}` . `[`querySelectorAll`](#dom-parentnode-queryselectorall){#ref-for-dom-parentnode-queryselectorall① .idl-code link-type="method"}`(``selectors`{.variable}`)`

:   Returns all [element](#concept-element){#ref-for-concept-element②⑨
    link-type="dfn"}
    [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant⑧
    link-type="dfn"} of `node`{.variable} that match
    `selectors`{.variable}.

The [`children`]{#dom-parentnode-children .dfn .dfn-paneled .idl-code
dfn-for="ParentNode" dfn-type="attribute" export=""} getter steps are to
return an
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①
link-type="idl"}
[collection](#concept-collection){#ref-for-concept-collection
link-type="dfn"} rooted at
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑦
link-type="dfn"} matching only
[element](#concept-element){#ref-for-concept-element③⓪ link-type="dfn"}
[children](#concept-tree-child){#ref-for-concept-tree-child④④
link-type="dfn"}.

The [`firstElementChild`]{#dom-parentnode-firstelementchild .dfn
.dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="attribute"
export=""} getter steps are to return the first
[child](#concept-tree-child){#ref-for-concept-tree-child④⑤
link-type="dfn"} that is an
[element](#concept-element){#ref-for-concept-element③① link-type="dfn"};
otherwise null.

The [`lastElementChild`]{#dom-parentnode-lastelementchild .dfn
.dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="attribute"
export=""} getter steps are to return the last
[child](#concept-tree-child){#ref-for-concept-tree-child④⑥
link-type="dfn"} that is an
[element](#concept-element){#ref-for-concept-element③② link-type="dfn"};
otherwise null.

The [`childElementCount`]{#dom-parentnode-childelementcount .dfn
.dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="attribute"
export=""} getter steps are to return the number of
[children](#concept-tree-child){#ref-for-concept-tree-child④⑦
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑧
link-type="dfn"} that are
[elements](#concept-element){#ref-for-concept-element③③
link-type="dfn"}.

The [`prepend(``nodes`{.variable}`)`]{#dom-parentnode-prepend .dfn
.dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="method" export=""
lt="prepend(...nodes)|prepend()"} method steps are:

1.  Let `node`{.variable} be the result of [converting nodes into a
    node](#converting-nodes-into-a-node){#ref-for-converting-nodes-into-a-node
    link-type="dfn"} given `nodes`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⑨
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document①②
    link-type="dfn"}.

2.  [Pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert②
    link-type="dfn"} `node`{.variable} into
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⓪
    link-type="dfn"} before
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④①
    link-type="dfn"}'s [first
    child](#concept-tree-first-child){#ref-for-concept-tree-first-child①
    link-type="dfn"}.

The [`append(``nodes`{.variable}`)`]{#dom-parentnode-append .dfn
.dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="method" export=""
lt="append(...nodes)|append()"} method steps are:

1.  Let `node`{.variable} be the result of [converting nodes into a
    node](#converting-nodes-into-a-node){#ref-for-converting-nodes-into-a-node①
    link-type="dfn"} given `nodes`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④②
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document①③
    link-type="dfn"}.

2.  [Append](#concept-node-append){#ref-for-concept-node-append①
    link-type="dfn"} `node`{.variable} to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④③
    link-type="dfn"}.

The
[`replaceChildren(``nodes`{.variable}`)`]{#dom-parentnode-replacechildren
.dfn .dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="method"
export="" lt="replaceChildren(...nodes)|replaceChildren()"} method steps
are:

1.  Let `node`{.variable} be the result of [converting nodes into a
    node](#converting-nodes-into-a-node){#ref-for-converting-nodes-into-a-node②
    link-type="dfn"} given `nodes`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④④
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document①④
    link-type="dfn"}.

2.  [Ensure pre-insert
    validity](#concept-node-ensure-pre-insertion-validity){#ref-for-concept-node-ensure-pre-insertion-validity①
    link-type="dfn"} of `node`{.variable} into
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑤
    link-type="dfn"} before null.

3.  [Replace
    all](#concept-node-replace-all){#ref-for-concept-node-replace-all
    link-type="dfn"} with `node`{.variable} within
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑥
    link-type="dfn"}.

The
[`moveBefore(``node`{.variable}`, ``child`{.variable}`)`]{#dom-parentnode-movebefore
.dfn .dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="method"
export=""} method steps are:

1.  Let `referenceChild`{.variable} be `child`{.variable}.

2.  If `referenceChild`{.variable} is `node`{.variable}, then set
    `referenceChild`{.variable} to `node`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling⑤
    link-type="dfn"}.

3.  [Move](#move){#ref-for-move② link-type="dfn"} `node`{.variable} into
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑦
    link-type="dfn"} before `referenceChild`{.variable}.

The
[`querySelector(``selectors`{.variable}`)`]{#dom-parentnode-queryselector
.dfn .dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="method"
export=""} method steps are to return the first result of running
[scope-match a selectors
string](#scope-match-a-selectors-string){#ref-for-scope-match-a-selectors-string
link-type="dfn"} `selectors`{.variable} against
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑧
link-type="dfn"}, if the result is not an empty list; otherwise null.

The
[`querySelectorAll(``selectors`{.variable}`)`]{#dom-parentnode-queryselectorall
.dfn .dfn-paneled .idl-code dfn-for="ParentNode" dfn-type="method"
export=""} method steps are to return the
[static](#concept-collection-static){#ref-for-concept-collection-static
link-type="dfn"} result of running [scope-match a selectors
string](#scope-match-a-selectors-string){#ref-for-scope-match-a-selectors-string①
link-type="dfn"} `selectors`{.variable} against
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this④⑨
link-type="dfn"}.

#### [4.2.7. ]{.secno}[Mixin [`NonDocumentTypeChildNode`{.idl}](#nondocumenttypechildnode){#ref-for-nondocumenttypechildnode link-type="idl"}]{.content}[](#interface-nondocumenttypechildnode){.self-link} {#interface-nondocumenttypechildnode .heading .settled level="4.2.7"}

Web compatibility prevents the
[`previousElementSibling`{.idl}](#dom-nondocumenttypechildnode-previouselementsibling){#ref-for-dom-nondocumenttypechildnode-previouselementsibling
link-type="idl"} and
[`nextElementSibling`{.idl}](#dom-nondocumenttypechildnode-nextelementsibling){#ref-for-dom-nondocumenttypechildnode-nextelementsibling
link-type="idl"} attributes from being exposed on
[doctypes](#concept-doctype){#ref-for-concept-doctype①③ link-type="dfn"}
(and therefore on [`ChildNode`{.idl}](#childnode){#ref-for-childnode
link-type="idl"}).

``` {.idl .highlight .def}
interface mixin NonDocumentTypeChildNode {
  readonly attribute Element? previousElementSibling;
  readonly attribute Element? nextElementSibling;
};
Element includes NonDocumentTypeChildNode;
CharacterData includes NonDocumentTypeChildNode;
```

`element`{.variable}` = ``node`{.variable}` . `[`previousElementSibling`{.idl}](#dom-nondocumenttypechildnode-previouselementsibling){#ref-for-dom-nondocumenttypechildnode-previouselementsibling② link-type="idl"}
:   Returns the first
    [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding⑤
    link-type="dfn"}
    [sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling⑥
    link-type="dfn"} that is an
    [element](#concept-element){#ref-for-concept-element③④
    link-type="dfn"}; otherwise null.

`element`{.variable}` = ``node`{.variable}` . `[`nextElementSibling`{.idl}](#dom-nondocumenttypechildnode-nextelementsibling){#ref-for-dom-nondocumenttypechildnode-nextelementsibling② link-type="idl"}
:   Returns the first
    [following](#concept-tree-following){#ref-for-concept-tree-following⑦
    link-type="dfn"}
    [sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling⑦
    link-type="dfn"} that is an
    [element](#concept-element){#ref-for-concept-element③⑤
    link-type="dfn"}; otherwise null.

The
[`previousElementSibling`]{#dom-nondocumenttypechildnode-previouselementsibling
.dfn .dfn-paneled .idl-code dfn-for="NonDocumentTypeChildNode"
dfn-type="attribute" export=""} getter steps are to return the first
[preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding⑥
link-type="dfn"}
[sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling⑧
link-type="dfn"} that is an
[element](#concept-element){#ref-for-concept-element③⑥ link-type="dfn"};
otherwise null.

The
[`nextElementSibling`]{#dom-nondocumenttypechildnode-nextelementsibling
.dfn .dfn-paneled .idl-code dfn-for="NonDocumentTypeChildNode"
dfn-type="attribute" export=""} getter steps are to return the first
[following](#concept-tree-following){#ref-for-concept-tree-following⑧
link-type="dfn"}
[sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling⑨
link-type="dfn"} that is an
[element](#concept-element){#ref-for-concept-element③⑦ link-type="dfn"};
otherwise null.

#### [4.2.8. ]{.secno}[Mixin [`ChildNode`{.idl}](#childnode){#ref-for-childnode① link-type="idl"}]{.content}[](#interface-childnode){.self-link} {#interface-childnode .heading .settled level="4.2.8"}

``` {.idl .highlight .def}
interface mixin ChildNode {
  [CEReactions, Unscopable] undefined before((Node or DOMString)... nodes);
  [CEReactions, Unscopable] undefined after((Node or DOMString)... nodes);
  [CEReactions, Unscopable] undefined replaceWith((Node or DOMString)... nodes);
  [CEReactions, Unscopable] undefined remove();
};
DocumentType includes ChildNode;
Element includes ChildNode;
CharacterData includes ChildNode;
```

`node`{.variable}` . `[`before(...nodes)`{.idl}](#dom-childnode-before){#ref-for-dom-childnode-before① link-type="idl"}

:   Inserts `nodes`{.variable} just before `node`{.variable}, while
    replacing strings in `nodes`{.variable} with equivalent
    [`Text`{.idl}](#text){#ref-for-text①⑦ link-type="idl"}
    [nodes](#concept-node){#ref-for-concept-node⑦④ link-type="dfn"}.

    [Throws](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③⓪
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror①⑨
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③⑥
    link-type="idl"} if the constraints of the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree①⑧
    link-type="dfn"} are violated.

`node`{.variable}` . `[`after(...nodes)`{.idl}](#dom-childnode-after){#ref-for-dom-childnode-after① link-type="idl"}

:   Inserts `nodes`{.variable} just after `node`{.variable}, while
    replacing strings in `nodes`{.variable} with equivalent
    [`Text`{.idl}](#text){#ref-for-text①⑧ link-type="idl"}
    [nodes](#concept-node){#ref-for-concept-node⑦⑤ link-type="dfn"}.

    [Throws](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③①
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③⑦
    link-type="idl"} if the constraints of the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree①⑨
    link-type="dfn"} are violated.

`node`{.variable}` . `[`replaceWith(...nodes)`{.idl}](#dom-childnode-replacewith){#ref-for-dom-childnode-replacewith① link-type="idl"}

:   Replaces `node`{.variable} with `nodes`{.variable}, while replacing
    strings in `nodes`{.variable} with equivalent
    [`Text`{.idl}](#text){#ref-for-text①⑨ link-type="idl"}
    [nodes](#concept-node){#ref-for-concept-node⑦⑥ link-type="dfn"}.

    [Throws](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③②
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③⑧
    link-type="idl"} if the constraints of the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree②⓪
    link-type="dfn"} are violated.

`node`{.variable}` . `[`remove()`{.idl}](#dom-childnode-remove){#ref-for-dom-childnode-remove① link-type="idl"}
:   Removes `node`{.variable}.

The [`before(``nodes`{.variable}`)`]{#dom-childnode-before .dfn
.dfn-paneled .idl-code dfn-for="ChildNode" dfn-type="method" export=""
lt="before(...nodes)|before()"} method steps are:

1.  Let `parent`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⓪
    link-type="dfn"}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①⑧
    link-type="dfn"}.

2.  If `parent`{.variable} is null, then return.

3.  Let `viablePreviousSibling`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤①
    link-type="dfn"}'s first
    [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding⑦
    link-type="dfn"}
    [sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling①⓪
    link-type="dfn"} not in `nodes`{.variable}; otherwise null.

4.  Let `node`{.variable} be the result of [converting nodes into a
    node](#converting-nodes-into-a-node){#ref-for-converting-nodes-into-a-node③
    link-type="dfn"}, given `nodes`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤②
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document①⑤
    link-type="dfn"}.

5.  If `viablePreviousSibling`{.variable} is null, then set it to
    `parent`{.variable}'s [first
    child](#concept-tree-first-child){#ref-for-concept-tree-first-child②
    link-type="dfn"}; otherwise to `viablePreviousSibling`{.variable}'s
    [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling⑥
    link-type="dfn"}.

6.  [Pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert③
    link-type="dfn"} `node`{.variable} into `parent`{.variable} before
    `viablePreviousSibling`{.variable}.

The [`after(``nodes`{.variable}`)`]{#dom-childnode-after .dfn
.dfn-paneled .idl-code dfn-for="ChildNode" dfn-type="method" export=""
lt="after(...nodes)|after()"} method steps are:

1.  Let `parent`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤③
    link-type="dfn"}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent①⑨
    link-type="dfn"}.

2.  If `parent`{.variable} is null, then return.

3.  Let `viableNextSibling`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤④
    link-type="dfn"}'s first
    [following](#concept-tree-following){#ref-for-concept-tree-following⑨
    link-type="dfn"}
    [sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling①①
    link-type="dfn"} not in `nodes`{.variable}; otherwise null.

4.  Let `node`{.variable} be the result of [converting nodes into a
    node](#converting-nodes-into-a-node){#ref-for-converting-nodes-into-a-node④
    link-type="dfn"}, given `nodes`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑤
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document①⑥
    link-type="dfn"}.

5.  [Pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert④
    link-type="dfn"} `node`{.variable} into `parent`{.variable} before
    `viableNextSibling`{.variable}.

The [`replaceWith(``nodes`{.variable}`)`]{#dom-childnode-replacewith
.dfn .dfn-paneled .idl-code dfn-for="ChildNode" dfn-type="method"
export="" lt="replaceWith(...nodes)|replaceWith()"} method steps are:

1.  Let `parent`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑥
    link-type="dfn"}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent②⓪
    link-type="dfn"}.

2.  If `parent`{.variable} is null, then return.

3.  Let `viableNextSibling`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑦
    link-type="dfn"}'s first
    [following](#concept-tree-following){#ref-for-concept-tree-following①⓪
    link-type="dfn"}
    [sibling](#concept-tree-sibling){#ref-for-concept-tree-sibling①②
    link-type="dfn"} not in `nodes`{.variable}; otherwise null.

4.  Let `node`{.variable} be the result of [converting nodes into a
    node](#converting-nodes-into-a-node){#ref-for-converting-nodes-into-a-node⑤
    link-type="dfn"}, given `nodes`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑧
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document①⑦
    link-type="dfn"}.

5.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑤⑨
    link-type="dfn"}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent②①
    link-type="dfn"} is `parent`{.variable},
    [replace](#concept-node-replace){#ref-for-concept-node-replace
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⓪
    link-type="dfn"} with `node`{.variable} within `parent`{.variable}.

    [This](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥①
    link-type="dfn"} could have been inserted into `node`{.variable}.

6.  Otherwise,
    [pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert⑤
    link-type="dfn"} `node`{.variable} into `parent`{.variable} before
    `viableNextSibling`{.variable}.

The [`remove()`]{#dom-childnode-remove .dfn .dfn-paneled .idl-code
dfn-for="ChildNode" dfn-type="method" export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥②
    link-type="dfn"}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent②②
    link-type="dfn"} is null, then return.

2.  [Remove](#concept-node-remove){#ref-for-concept-node-remove⑦
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥③
    link-type="dfn"}.

#### [4.2.9. ]{.secno}[Mixin [`Slottable`{.idl}](#slotable){#ref-for-slotable link-type="idl"}]{.content}[](#mixin-slotable){.self-link} {#mixin-slotable .heading .settled level="4.2.9"}

``` {.idl .highlight .def}
interface mixin Slottable {
  readonly attribute HTMLSlotElement? assignedSlot;
};
Element includes Slottable;
Text includes Slottable;
```

The [`assignedSlot`]{#dom-slotable-assignedslot .dfn .dfn-paneled
.idl-code dfn-for="Slottable" dfn-type="attribute" export=""} getter
steps are to return the result of [find a
slot](#find-a-slot){#ref-for-find-a-slot② link-type="dfn"} given
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥④
link-type="dfn"} and true.

#### [4.2.10. ]{.secno}[Old-style collections: [`NodeList`{.idl}](#nodelist){#ref-for-nodelist① link-type="idl"} and [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection② link-type="idl"}]{.content}[](#old-style-collections){.self-link} {#old-style-collections .heading .settled level="4.2.10"}

A [collection]{#concept-collection .dfn .dfn-paneled dfn-type="dfn"
export=""} is an object that represents a list of
[nodes](#concept-node){#ref-for-concept-node⑦⑦ link-type="dfn"}. A
[collection](#concept-collection){#ref-for-concept-collection①
link-type="dfn"} can be either [live]{#concept-collection-live .dfn
.dfn-paneled dfn-for="collection" dfn-type="dfn" export=""
local-lt="live" lt="live collection"} or
[static]{#concept-collection-static .dfn .dfn-paneled
dfn-for="collection" dfn-type="dfn" export="" lt="static collection"}.
Unless otherwise stated, a
[collection](#concept-collection){#ref-for-concept-collection②
link-type="dfn"} must be
[live](#concept-collection-live){#ref-for-concept-collection-live
link-type="dfn"}.

If a [collection](#concept-collection){#ref-for-concept-collection③
link-type="dfn"} is
[live](#concept-collection-live){#ref-for-concept-collection-live①
link-type="dfn"}, then the attributes and methods on that object must
operate on the actual underlying data, not a snapshot of the data.

When a [collection](#concept-collection){#ref-for-concept-collection④
link-type="dfn"} is created, a filter and a root are associated with it.

The [collection](#concept-collection){#ref-for-concept-collection⑤
link-type="dfn"} then [represents]{#represented-by-the-collection .dfn
.dfn-paneled dfn-for="collection" dfn-type="dfn" export=""
lt="represented by the collection"} a view of the subtree rooted at the
[collection's](#concept-collection){#ref-for-concept-collection⑥
link-type="dfn"} root, containing only nodes that match the given
filter. The view is linear. In the absence of specific requirements to
the contrary, the nodes within the
[collection](#concept-collection){#ref-for-concept-collection⑦
link-type="dfn"} must be sorted in [tree
order](#concept-tree-order){#ref-for-concept-tree-order①⑤
link-type="dfn"}.

##### [4.2.10.1. ]{.secno}[Interface [`NodeList`{.idl}](#nodelist){#ref-for-nodelist② link-type="idl"}]{.content}[](#interface-nodelist){.self-link} {#interface-nodelist .heading .settled level="4.2.10.1"}

A [`NodeList`{.idl}](#nodelist){#ref-for-nodelist③ link-type="idl"}
object is a
[collection](#concept-collection){#ref-for-concept-collection⑧
link-type="dfn"} of [nodes](#concept-node){#ref-for-concept-node⑦⑧
link-type="dfn"}.

``` {.idl .highlight .def}
[Exposed=Window]
interface NodeList {
  getter Node? item(unsigned long index);
  readonly attribute unsigned long length;
  iterable<Node>;
};
```

`collection`{.variable} . [`length`{.idl}](#dom-nodelist-length){#ref-for-dom-nodelist-length① link-type="idl"}
:   Returns the number of [nodes](#concept-node){#ref-for-concept-node⑦⑨
    link-type="dfn"} in the
    [collection](#concept-collection){#ref-for-concept-collection⑨
    link-type="dfn"}.

`element`{.variable} = `collection`{.variable} . [`item(index)`{.idl}](#dom-nodelist-item){#ref-for-dom-nodelist-item① link-type="idl"}\
`element`{.variable} = `collection`{.variable}\[`index`{.variable}\]
:   Returns the [node](#concept-node){#ref-for-concept-node⑧⓪
    link-type="dfn"} with index `index`{.variable} from the
    [collection](#concept-collection){#ref-for-concept-collection①⓪
    link-type="dfn"}. The [nodes](#concept-node){#ref-for-concept-node⑧①
    link-type="dfn"} are sorted in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order①⑥
    link-type="dfn"}.

::: impl
The object's [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices
link-type="dfn"} are the numbers in the range zero to one less than the
number of nodes [represented by the
collection](#represented-by-the-collection){#ref-for-represented-by-the-collection
link-type="dfn"}. If there are no such elements, then there are no
[supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices①
link-type="dfn"}.

The [`length`]{#dom-nodelist-length .dfn .dfn-paneled .idl-code
dfn-for="NodeList" dfn-type="attribute" export=""} attribute must return
the number of nodes [represented by the
collection](#represented-by-the-collection){#ref-for-represented-by-the-collection①
link-type="dfn"}.

The [`item(``index`{.variable}`)`]{#dom-nodelist-item .dfn .dfn-paneled
.idl-code dfn-for="NodeList" dfn-type="method" export=""} method must
return the `index`{.variable}^th^
[node](#concept-node){#ref-for-concept-node⑧② link-type="dfn"} in the
[collection](#concept-collection){#ref-for-concept-collection①①
link-type="dfn"}. If there is no `index`{.variable}^th^
[node](#concept-node){#ref-for-concept-node⑧③ link-type="dfn"} in the
[collection](#concept-collection){#ref-for-concept-collection①②
link-type="dfn"}, then the method must return null.
:::

##### [4.2.10.2. ]{.secno}[Interface [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection③ link-type="idl"}]{.content}[](#interface-htmlcollection){.self-link} {#interface-htmlcollection .heading .settled level="4.2.10.2"}

``` {.idl .highlight .def}
[Exposed=Window, LegacyUnenumerableNamedProperties]
interface HTMLCollection {
  readonly attribute unsigned long length;
  getter Element? item(unsigned long index);
  getter Element? namedItem(DOMString name);
};
```

An [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection④
link-type="idl"} object is a
[collection](#concept-collection){#ref-for-concept-collection①③
link-type="dfn"} of
[elements](#concept-element){#ref-for-concept-element③⑧
link-type="dfn"}.

[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection⑤
link-type="idl"} is a historical artifact we cannot rid the web of.
While developers are of course welcome to keep using it, new API
standard designers ought not to use it (use `sequence<T>` in IDL
instead).

`collection`{.variable} . [`length`{.idl}](#dom-htmlcollection-length){#ref-for-dom-htmlcollection-length① link-type="idl"}
:   Returns the number of
    [elements](#concept-element){#ref-for-concept-element③⑨
    link-type="dfn"} in the
    [collection](#concept-collection){#ref-for-concept-collection①④
    link-type="dfn"}.

`element`{.variable} = `collection`{.variable} . [`item(index)`{.idl}](#dom-htmlcollection-item){#ref-for-dom-htmlcollection-item① link-type="idl"}\
`element`{.variable} = `collection`{.variable}\[`index`{.variable}\]
:   Returns the [element](#concept-element){#ref-for-concept-element④⓪
    link-type="dfn"} with index `index`{.variable} from the
    [collection](#concept-collection){#ref-for-concept-collection①⑤
    link-type="dfn"}. The
    [elements](#concept-element){#ref-for-concept-element④①
    link-type="dfn"} are sorted in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order①⑦
    link-type="dfn"}.

`element`{.variable} = `collection`{.variable} . [`namedItem(name)`{.idl}](#dom-htmlcollection-nameditem){#ref-for-dom-htmlcollection-nameditem link-type="idl"}\
`element`{.variable} = `collection`{.variable}\[`name`{.variable}\]
:   Returns the first
    [element](#concept-element){#ref-for-concept-element④②
    link-type="dfn"} with [ID](#concept-id){#ref-for-concept-id②
    link-type="dfn"} or name `name`{.variable} from the collection.

The object's [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices②
link-type="dfn"} are the numbers in the range zero to one less than the
number of elements [represented by the
collection](#represented-by-the-collection){#ref-for-represented-by-the-collection②
link-type="dfn"}. If there are no such elements, then there are no
[supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices③
link-type="dfn"}.

The [`length`]{#dom-htmlcollection-length .dfn .dfn-paneled .idl-code
dfn-for="HTMLCollection" dfn-type="attribute" export=""} getter steps
are to return the number of nodes [represented by the
collection](#represented-by-the-collection){#ref-for-represented-by-the-collection③
link-type="dfn"}.

The [`item(``index`{.variable}`)`]{#dom-htmlcollection-item .dfn
.dfn-paneled .idl-code dfn-for="HTMLCollection" dfn-type="method"
export=""} method steps are to return the `index`{.variable}^th^
[element](#concept-element){#ref-for-concept-element④③ link-type="dfn"}
in the [collection](#concept-collection){#ref-for-concept-collection①⑥
link-type="dfn"}. If there is no `index`{.variable}^th^
[element](#concept-element){#ref-for-concept-element④④ link-type="dfn"}
in the [collection](#concept-collection){#ref-for-concept-collection①⑦
link-type="dfn"}, then the method must return null.

The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#ref-for-dfn-supported-property-names
link-type="dfn"} are the values from the list returned by these steps:

1.  Let `result`{.variable} be an empty list.

2.  For each `element`{.variable} [represented by the
    collection](#represented-by-the-collection){#ref-for-represented-by-the-collection④
    link-type="dfn"}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order①⑧
    link-type="dfn"}:

    1.  If `element`{.variable} has an
        [ID](#concept-id){#ref-for-concept-id③ link-type="dfn"} which is
        not in `result`{.variable}, append `element`{.variable}'s
        [ID](#concept-id){#ref-for-concept-id④ link-type="dfn"} to
        `result`{.variable}.

    2.  If `element`{.variable} is in the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace
        link-type="dfn"} and
        [has](#concept-element-attribute-has){#ref-for-concept-element-attribute-has
        link-type="dfn"} a [`name`
        attribute](#concept-named-attribute){#ref-for-concept-named-attribute
        link-type="dfn"} whose
        [value](#concept-attribute-value){#ref-for-concept-attribute-value
        link-type="dfn"} is neither the empty string nor is in
        `result`{.variable}, append `element`{.variable}'s [`name`
        attribute](#concept-named-attribute){#ref-for-concept-named-attribute①
        link-type="dfn"}
        [value](#concept-attribute-value){#ref-for-concept-attribute-value①
        link-type="dfn"} to `result`{.variable}.

3.  Return `result`{.variable}.

The [`namedItem(``key`{.variable}`)`]{#dom-htmlcollection-nameditem-key
.dfn .dfn-paneled .idl-code dfn-for="HTMLCollection" dfn-type="method"
export=""} method steps are:

1.  If `key`{.variable} is the empty string, return null.

2.  Return the first
    [element](#concept-element){#ref-for-concept-element④⑤
    link-type="dfn"} in the
    [collection](#concept-collection){#ref-for-concept-collection①⑧
    link-type="dfn"} for which at least one of the following is true:

    - it has an [ID](#concept-id){#ref-for-concept-id⑤ link-type="dfn"}
      which is `key`{.variable};
    - it is in the [HTML
      namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①
      link-type="dfn"} and
      [has](#concept-element-attribute-has){#ref-for-concept-element-attribute-has①
      link-type="dfn"} a [`name`
      attribute](#concept-named-attribute){#ref-for-concept-named-attribute②
      link-type="dfn"} whose
      [value](#concept-attribute-value){#ref-for-concept-attribute-value②
      link-type="dfn"} is `key`{.variable};

    or null if there is no such
    [element](#concept-element){#ref-for-concept-element④⑥
    link-type="dfn"}.

### [4.3. ]{.secno}[Mutation observers]{.content}[](#mutation-observers){.self-link} {#mutation-observers .heading .settled level="4.3"}

Each [similar-origin window
agent](https://html.spec.whatwg.org/multipage/webappapis.html#similar-origin-window-agent){#ref-for-similar-origin-window-agent①
link-type="dfn"} has a [mutation observer microtask
queued]{#mutation-observer-compound-microtask-queued-flag .dfn
.dfn-paneled dfn-type="dfn" noexport=""} (a boolean), which is initially
false. [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

Each [similar-origin window
agent](https://html.spec.whatwg.org/multipage/webappapis.html#similar-origin-window-agent){#ref-for-similar-origin-window-agent②
link-type="dfn"} also has [pending mutation
observers]{#mutation-observer-list .dfn .dfn-paneled dfn-type="dfn"
noexport=""} (a
[set](https://infra.spec.whatwg.org/#ordered-set){#ref-for-ordered-set⑥
link-type="dfn"} of zero or more
[`MutationObserver`{.idl}](#mutationobserver){#ref-for-mutationobserver
link-type="idl"} objects), which is initially empty.

To [queue a mutation observer
microtask]{#queue-a-mutation-observer-compound-microtask .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, run these steps:

1.  If the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent
    link-type="dfn"}'s [mutation observer microtask
    queued](#mutation-observer-compound-microtask-queued-flag){#ref-for-mutation-observer-compound-microtask-queued-flag
    link-type="dfn"} is true, then return.

2.  Set the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent①
    link-type="dfn"}'s [mutation observer microtask
    queued](#mutation-observer-compound-microtask-queued-flag){#ref-for-mutation-observer-compound-microtask-queued-flag①
    link-type="dfn"} to true.

3.  [Queue](https://html.spec.whatwg.org/multipage/webappapis.html#queue-a-microtask){#ref-for-queue-a-microtask
    link-type="dfn"} a
    [microtask](https://html.spec.whatwg.org/multipage/webappapis.html#microtask){#ref-for-microtask
    link-type="dfn"} to [notify mutation
    observers](#notify-mutation-observers){#ref-for-notify-mutation-observers
    link-type="dfn"}.

To [notify mutation observers]{#notify-mutation-observers .dfn
.dfn-paneled dfn-type="dfn" export=""}, run these steps:

1.  Set the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent②
    link-type="dfn"}'s [mutation observer microtask
    queued](#mutation-observer-compound-microtask-queued-flag){#ref-for-mutation-observer-compound-microtask-queued-flag②
    link-type="dfn"} to false.

2.  Let `notifySet`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone①
    link-type="dfn"} of the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent③
    link-type="dfn"}'s [pending mutation
    observers](#mutation-observer-list){#ref-for-mutation-observer-list
    link-type="dfn"}.

3.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty①
    link-type="dfn"} the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent④
    link-type="dfn"}'s [pending mutation
    observers](#mutation-observer-list){#ref-for-mutation-observer-list①
    link-type="dfn"}.

4.  Let `signalSet`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone②
    link-type="dfn"} of the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent⑤
    link-type="dfn"}'s [signal
    slots](#signal-slot-list){#ref-for-signal-slot-list①
    link-type="dfn"}.

5.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty②
    link-type="dfn"} the [surrounding
    agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent⑥
    link-type="dfn"}'s [signal
    slots](#signal-slot-list){#ref-for-signal-slot-list②
    link-type="dfn"}.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑦
    link-type="dfn"} `mo`{.variable} of `notifySet`{.variable}:

    1.  Let `records`{.variable} be a
        [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone③
        link-type="dfn"} of `mo`{.variable}'s [record
        queue](#concept-mo-queue){#ref-for-concept-mo-queue
        link-type="dfn"}.

    2.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty③
        link-type="dfn"} `mo`{.variable}'s [record
        queue](#concept-mo-queue){#ref-for-concept-mo-queue①
        link-type="dfn"}.

    3.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑧
        link-type="dfn"} `node`{.variable} of `mo`{.variable}'s [node
        list](#mutationobserver-node-list){#ref-for-mutationobserver-node-list
        link-type="dfn"},
        [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove④
        link-type="dfn"} all [transient registered
        observers](#transient-registered-observer){#ref-for-transient-registered-observer①
        link-type="dfn"} whose
        [observer](#registered-observer-observer){#ref-for-registered-observer-observer②
        link-type="dfn"} is `mo`{.variable} from `node`{.variable}'s
        [registered observer
        list](#registered-observer-list){#ref-for-registered-observer-list②
        link-type="dfn"}.

    4.  If `records`{.variable} [is not
        empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty④
        link-type="dfn"}, then
        [invoke](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#ref-for-invoke-a-callback-function
        link-type="dfn"} `mo`{.variable}'s
        [callback](#concept-mo-callback){#ref-for-concept-mo-callback
        link-type="dfn"} with « `records`{.variable}, `mo`{.variable} »
        and \"`report`\", and with [callback this
        value](https://webidl.spec.whatwg.org/#dfn-callback-this-value){#ref-for-dfn-callback-this-value
        link-type="dfn"} `mo`{.variable}.

7.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate①⑨
    link-type="dfn"} `slot`{.variable} of `signalSet`{.variable}, [fire
    an event](#concept-event-fire){#ref-for-concept-event-fire⑧
    link-type="dfn"} named
    [`slotchange`]{#eventdef-htmlslotelement-slotchange .dfn
    .dfn-paneled .idl-code dfn-for="HTMLSlotElement" dfn-type="event"
    export=""}, with its
    [`bubbles`{.idl}](#dom-event-bubbles){#ref-for-dom-event-bubbles⑧
    link-type="idl"} attribute set to true, at `slot`{.variable}.

------------------------------------------------------------------------

Each [node](#concept-node){#ref-for-concept-node⑧④ link-type="dfn"} has
a [registered observer list]{#registered-observer-list .dfn .dfn-paneled
dfn-type="dfn" noexport=""} (a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①⓪
link-type="dfn"} of zero or more [registered
observers](#registered-observer){#ref-for-registered-observer
link-type="dfn"}), which is initially empty.

A [registered observer]{#registered-observer .dfn .dfn-paneled
dfn-type="dfn" noexport=""} consists of an
[observer]{#registered-observer-observer .dfn .dfn-paneled
dfn-for="registered observer" dfn-type="dfn" noexport=""} (a
[`MutationObserver`{.idl}](#mutationobserver){#ref-for-mutationobserver①
link-type="idl"} object) and [options]{#registered-observer-options .dfn
.dfn-paneled dfn-for="registered observer" dfn-type="dfn" noexport=""}
(a
[`MutationObserverInit`{.idl}](#dictdef-mutationobserverinit){#ref-for-dictdef-mutationobserverinit
link-type="idl"} dictionary).

A [transient registered observer]{#transient-registered-observer .dfn
.dfn-paneled dfn-type="dfn" noexport=""} is a [registered
observer](#registered-observer){#ref-for-registered-observer①
link-type="dfn"} that also consists of a
[source]{#transient-registered-observer-source .dfn .dfn-paneled
dfn-for="transient registered observer" dfn-type="dfn" noexport=""} (a
[registered
observer](#registered-observer){#ref-for-registered-observer②
link-type="dfn"}).

[Transient registered
observers](#transient-registered-observer){#ref-for-transient-registered-observer②
link-type="dfn"} are used to track mutations within a given
[node](#concept-node){#ref-for-concept-node⑧⑤ link-type="dfn"}'s
[descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant⑨
link-type="dfn"} after [node](#concept-node){#ref-for-concept-node⑧⑥
link-type="dfn"} has been removed so they do not get lost when
[`subtree`{.idl}](#dom-mutationobserverinit-subtree){#ref-for-dom-mutationobserverinit-subtree①
link-type="idl"} is set to true on
[node](#concept-node){#ref-for-concept-node⑧⑦ link-type="dfn"}'s
[parent](#concept-tree-parent){#ref-for-concept-tree-parent②③
link-type="dfn"}.

#### [4.3.1. ]{.secno}[Interface [`MutationObserver`{.idl}](#mutationobserver){#ref-for-mutationobserver② link-type="idl"}]{.content}[](#interface-mutationobserver){.self-link} {#interface-mutationobserver .heading .settled level="4.3.1"}

``` {.idl .highlight .def}
[Exposed=Window]
interface MutationObserver {
  constructor(MutationCallback callback);

  undefined observe(Node target, optional MutationObserverInit options = {});
  undefined disconnect();
  sequence<MutationRecord> takeRecords();
};

callback MutationCallback = undefined (sequence<MutationRecord> mutations, MutationObserver observer);

dictionary MutationObserverInit {
  boolean childList = false;
  boolean attributes;
  boolean characterData;
  boolean subtree = false;
  boolean attributeOldValue;
  boolean characterDataOldValue;
  sequence<DOMString> attributeFilter;
};
```

A
[`MutationObserver`{.idl}](#mutationobserver){#ref-for-mutationobserver④
link-type="idl"} object can be used to observe mutations to the
[tree](#concept-tree){#ref-for-concept-tree①⑦ link-type="dfn"} of
[nodes](#concept-node){#ref-for-concept-node⑧⑧ link-type="dfn"}.

Each
[`MutationObserver`{.idl}](#mutationobserver){#ref-for-mutationobserver⑤
link-type="idl"} object has these associated concepts:

- A [callback]{#concept-mo-callback .dfn .dfn-paneled
  dfn-for="MutationObserver" dfn-type="dfn" noexport=""} set on
  creation.
- A [node list]{#mutationobserver-node-list .dfn .dfn-paneled
  dfn-for="MutationObserver" dfn-type="dfn" noexport=""} (a
  [list](https://infra.spec.whatwg.org/#list){#ref-for-list①①
  link-type="dfn"} of weak references to
  [nodes](#concept-node){#ref-for-concept-node⑧⑨ link-type="dfn"}),
  which is initially empty.
- A [record queue]{#concept-mo-queue .dfn .dfn-paneled
  dfn-for="MutationObserver" dfn-type="dfn" noexport=""} (a
  [queue](https://infra.spec.whatwg.org/#queue){#ref-for-queue
  link-type="dfn"} of zero or more
  [`MutationRecord`{.idl}](#mutationrecord){#ref-for-mutationrecord②
  link-type="idl"} objects), which is initially empty.

`observer`{.variable}` = new `[`MutationObserver(callback)`{.idl}](#dom-mutationobserver-mutationobserver){#ref-for-dom-mutationobserver-mutationobserver① link-type="idl"}
:   Constructs a
    [`MutationObserver`{.idl}](#mutationobserver){#ref-for-mutationobserver⑥
    link-type="idl"} object and sets its
    [callback](#concept-mo-callback){#ref-for-concept-mo-callback①
    link-type="dfn"} to `callback`{.variable}. The `callback`{.variable}
    is invoked with a list of
    [`MutationRecord`{.idl}](#mutationrecord){#ref-for-mutationrecord③
    link-type="idl"} objects as first argument and the constructed
    [`MutationObserver`{.idl}](#mutationobserver){#ref-for-mutationobserver⑦
    link-type="idl"} object as second argument. It is invoked after
    [nodes](#concept-node){#ref-for-concept-node⑨⓪ link-type="dfn"}
    registered with the
    [`observe()`{.idl}](#dom-mutationobserver-observe){#ref-for-dom-mutationobserver-observe①
    link-type="idl"} method, are mutated.

`observer`{.variable}` . `[`observe(target, options)`{.idl}](#dom-mutationobserver-observe){#ref-for-dom-mutationobserver-observe② link-type="idl"}

:   Instructs the user agent to observe a given `target`{.variable} (a
    [node](#concept-node){#ref-for-concept-node⑨① link-type="dfn"}) and
    report any mutations based on the criteria given by
    `options`{.variable} (an object).

    The `options`{.variable} argument allows for setting mutation
    observation options via object members. These are the object members
    that can be used:

    [`childList`{.idl}](#dom-mutationobserverinit-childlist){#ref-for-dom-mutationobserverinit-childlist link-type="idl"}
    :   Set to true if mutations to `target`{.variable}'s
        [children](#concept-tree-child){#ref-for-concept-tree-child④⑧
        link-type="dfn"} are to be observed.

    [`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes link-type="idl"}
    :   Set to true if mutations to `target`{.variable}'s
        [attributes](#concept-attribute){#ref-for-concept-attribute
        link-type="dfn"} are to be observed. Can be omitted if
        [`attributeOldValue`{.idl}](#dom-mutationobserverinit-attributeoldvalue){#ref-for-dom-mutationobserverinit-attributeoldvalue
        link-type="idl"} or
        [`attributeFilter`{.idl}](#dom-mutationobserverinit-attributefilter){#ref-for-dom-mutationobserverinit-attributefilter
        link-type="idl"} is specified.

    [`characterData`{.idl}](#dom-mutationobserverinit-characterdata){#ref-for-dom-mutationobserverinit-characterdata link-type="idl"}
    :   Set to true if mutations to `target`{.variable}'s
        [data](#concept-cd-data){#ref-for-concept-cd-data②
        link-type="dfn"} are to be observed. Can be omitted if
        [`characterDataOldValue`{.idl}](#dom-mutationobserverinit-characterdataoldvalue){#ref-for-dom-mutationobserverinit-characterdataoldvalue
        link-type="idl"} is specified.

    [`subtree`{.idl}](#dom-mutationobserverinit-subtree){#ref-for-dom-mutationobserverinit-subtree② link-type="idl"}
    :   Set to true if mutations to not just `target`{.variable}, but
        also `target`{.variable}'s
        [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant①⓪
        link-type="dfn"} are to be observed.

    [`attributeOldValue`{.idl}](#dom-mutationobserverinit-attributeoldvalue){#ref-for-dom-mutationobserverinit-attributeoldvalue① link-type="idl"}
    :   Set to true if
        [`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes①
        link-type="idl"} is true or omitted and `target`{.variable}'s
        [attribute](#concept-attribute){#ref-for-concept-attribute①
        link-type="dfn"}
        [value](#concept-attribute-value){#ref-for-concept-attribute-value③
        link-type="dfn"} before the mutation needs to be recorded.

    [`characterDataOldValue`{.idl}](#dom-mutationobserverinit-characterdataoldvalue){#ref-for-dom-mutationobserverinit-characterdataoldvalue① link-type="idl"}
    :   Set to true if
        [`characterData`{.idl}](#dom-mutationobserverinit-characterdata){#ref-for-dom-mutationobserverinit-characterdata①
        link-type="idl"} is set to true or omitted and
        `target`{.variable}'s
        [data](#concept-cd-data){#ref-for-concept-cd-data③
        link-type="dfn"} before the mutation needs to be recorded.

    [`attributeFilter`{.idl}](#dom-mutationobserverinit-attributefilter){#ref-for-dom-mutationobserverinit-attributefilter① link-type="idl"}
    :   Set to a list of
        [attribute](#concept-attribute){#ref-for-concept-attribute②
        link-type="dfn"} [local
        names](#concept-attribute-local-name){#ref-for-concept-attribute-local-name
        link-type="dfn"} (without
        [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace
        link-type="dfn"}) if not all
        [attribute](#concept-attribute){#ref-for-concept-attribute③
        link-type="dfn"} mutations need to be observed and
        [`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes②
        link-type="idl"} is true or omitted.

`observer`{.variable}` . `[`disconnect()`{.idl}](#dom-mutationobserver-disconnect){#ref-for-dom-mutationobserver-disconnect① link-type="idl"}
:   Stops `observer`{.variable} from observing any mutations. Until the
    [`observe()`{.idl}](#dom-mutationobserver-observe){#ref-for-dom-mutationobserver-observe③
    link-type="idl"} method is used again, `observer`{.variable}'s
    [callback](#concept-mo-callback){#ref-for-concept-mo-callback②
    link-type="dfn"} will not be invoked.

`observer`{.variable}` . `[`takeRecords()`{.idl}](#dom-mutationobserver-takerecords){#ref-for-dom-mutationobserver-takerecords① link-type="idl"}
:   Empties the [record
    queue](#concept-mo-queue){#ref-for-concept-mo-queue②
    link-type="dfn"} and returns what was in there.

The
[`new MutationObserver(``callback`{.variable}`)`]{#dom-mutationobserver-mutationobserver
.dfn .dfn-paneled .idl-code dfn-for="MutationObserver"
dfn-type="constructor" export=""
lt="MutationObserver(callback)|constructor(callback)"} constructor steps
are to set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑤
link-type="dfn"}'s
[callback](#concept-mo-callback){#ref-for-concept-mo-callback③
link-type="dfn"} to `callback`{.variable}.

The
[`observe(``target`{.variable}`, ``options`{.variable}`)`]{#dom-mutationobserver-observe
.dfn .dfn-paneled .idl-code dfn-for="MutationObserver" dfn-type="method"
export="" lt="observe(target, options)|observe(target)"} method steps
are:

1.  If either
    `options`{.variable}\[\"[`attributeOldValue`{.idl}](#dom-mutationobserverinit-attributeoldvalue){#ref-for-dom-mutationobserverinit-attributeoldvalue②
    link-type="idl"}\"\] or
    `options`{.variable}\[\"[`attributeFilter`{.idl}](#dom-mutationobserverinit-attributefilter){#ref-for-dom-mutationobserverinit-attributefilter②
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists③
    link-type="dfn"}, and
    `options`{.variable}\[\"[`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes③
    link-type="idl"}\"\] does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists④
    link-type="dfn"}, then set
    `options`{.variable}\[\"[`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes④
    link-type="idl"}\"\] to true.

2.  If
    `options`{.variable}\[\"[`characterDataOldValue`{.idl}](#dom-mutationobserverinit-characterdataoldvalue){#ref-for-dom-mutationobserverinit-characterdataoldvalue②
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑤
    link-type="dfn"} and
    `options`{.variable}\[\"[`characterData`{.idl}](#dom-mutationobserverinit-characterdata){#ref-for-dom-mutationobserverinit-characterdata②
    link-type="idl"}\"\] does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑥
    link-type="dfn"}, then set
    `options`{.variable}\[\"[`characterData`{.idl}](#dom-mutationobserverinit-characterdata){#ref-for-dom-mutationobserverinit-characterdata③
    link-type="idl"}\"\] to true.

3.  If none of
    `options`{.variable}\[\"[`childList`{.idl}](#dom-mutationobserverinit-childlist){#ref-for-dom-mutationobserverinit-childlist①
    link-type="idl"}\"\],
    `options`{.variable}\[\"[`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes⑤
    link-type="idl"}\"\], and
    `options`{.variable}\[\"[`characterData`{.idl}](#dom-mutationobserverinit-characterdata){#ref-for-dom-mutationobserverinit-characterdata④
    link-type="idl"}\"\] is true, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③③
    link-type="dfn"} a `TypeError`.

4.  If
    `options`{.variable}\[\"[`attributeOldValue`{.idl}](#dom-mutationobserverinit-attributeoldvalue){#ref-for-dom-mutationobserverinit-attributeoldvalue③
    link-type="idl"}\"\] is true and
    `options`{.variable}\[\"[`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes⑥
    link-type="idl"}\"\] is false, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③④
    link-type="dfn"} a `TypeError`.

5.  If
    `options`{.variable}\[\"[`attributeFilter`{.idl}](#dom-mutationobserverinit-attributefilter){#ref-for-dom-mutationobserverinit-attributefilter③
    link-type="idl"}\"\] is present and
    `options`{.variable}\[\"[`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes⑦
    link-type="idl"}\"\] is false, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③⑤
    link-type="dfn"} a `TypeError`.

6.  If
    `options`{.variable}\[\"[`characterDataOldValue`{.idl}](#dom-mutationobserverinit-characterdataoldvalue){#ref-for-dom-mutationobserverinit-characterdataoldvalue③
    link-type="idl"}\"\] is true and
    `options`{.variable}\[\"[`characterData`{.idl}](#dom-mutationobserverinit-characterdata){#ref-for-dom-mutationobserverinit-characterdata⑤
    link-type="idl"}\"\] is false, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③⑥
    link-type="dfn"} a `TypeError`.

7.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②⓪
    link-type="dfn"} `registered`{.variable} of `target`{.variable}'s
    [registered observer
    list](#registered-observer-list){#ref-for-registered-observer-list③
    link-type="dfn"}, if `registered`{.variable}'s
    [observer](#registered-observer-observer){#ref-for-registered-observer-observer③
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑥
    link-type="dfn"}:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②①
        link-type="dfn"} `node`{.variable} of
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑦
        link-type="dfn"}'s [node
        list](#mutationobserver-node-list){#ref-for-mutationobserver-node-list①
        link-type="dfn"},
        [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove⑤
        link-type="dfn"} all [transient registered
        observers](#transient-registered-observer){#ref-for-transient-registered-observer③
        link-type="dfn"} whose
        [source](#transient-registered-observer-source){#ref-for-transient-registered-observer-source①
        link-type="dfn"} is `registered`{.variable} from
        `node`{.variable}'s [registered observer
        list](#registered-observer-list){#ref-for-registered-observer-list④
        link-type="dfn"}.

    2.  Set `registered`{.variable}'s
        [options](#registered-observer-options){#ref-for-registered-observer-options③
        link-type="dfn"} to `options`{.variable}.

8.  Otherwise:

    1.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①②
        link-type="dfn"} a new [registered
        observer](#registered-observer){#ref-for-registered-observer③
        link-type="dfn"} whose
        [observer](#registered-observer-observer){#ref-for-registered-observer-observer④
        link-type="dfn"} is
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑧
        link-type="dfn"} and
        [options](#registered-observer-options){#ref-for-registered-observer-options④
        link-type="dfn"} is `options`{.variable} to
        `target`{.variable}'s [registered observer
        list](#registered-observer-list){#ref-for-registered-observer-list⑤
        link-type="dfn"}.

    2.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①③
        link-type="dfn"} a weak reference to `target`{.variable} to
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑥⑨
        link-type="dfn"}'s [node
        list](#mutationobserver-node-list){#ref-for-mutationobserver-node-list②
        link-type="dfn"}.

The [`disconnect()`]{#dom-mutationobserver-disconnect .dfn .dfn-paneled
.idl-code dfn-for="MutationObserver" dfn-type="method" export=""} method
steps are:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②②
    link-type="dfn"} `node`{.variable} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⓪
    link-type="dfn"}'s [node
    list](#mutationobserver-node-list){#ref-for-mutationobserver-node-list③
    link-type="dfn"},
    [remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove⑥
    link-type="dfn"} any [registered
    observer](#registered-observer){#ref-for-registered-observer④
    link-type="dfn"} from `node`{.variable}'s [registered observer
    list](#registered-observer-list){#ref-for-registered-observer-list⑥
    link-type="dfn"} for which
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦①
    link-type="dfn"} is the
    [observer](#registered-observer-observer){#ref-for-registered-observer-observer⑤
    link-type="dfn"}.

2.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty④
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦②
    link-type="dfn"}'s [record
    queue](#concept-mo-queue){#ref-for-concept-mo-queue③
    link-type="dfn"}.

The [`takeRecords()`]{#dom-mutationobserver-takerecords .dfn
.dfn-paneled .idl-code dfn-for="MutationObserver" dfn-type="method"
export=""} method steps are:

1.  Let `records`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ref-for-list-clone④
    link-type="dfn"} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦③
    link-type="dfn"}'s [record
    queue](#concept-mo-queue){#ref-for-concept-mo-queue④
    link-type="dfn"}.

2.  [Empty](https://infra.spec.whatwg.org/#list-empty){#ref-for-list-empty⑤
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦④
    link-type="dfn"}'s [record
    queue](#concept-mo-queue){#ref-for-concept-mo-queue⑤
    link-type="dfn"}.

3.  Return `records`{.variable}.

#### [4.3.2. ]{.secno}[Queuing a mutation record]{.content}[](#queueing-a-mutation-record){.self-link} {#queueing-a-mutation-record .heading .settled level="4.3.2"}

To [queue a mutation record]{#queue-a-mutation-record .dfn .dfn-paneled
dfn-type="dfn" noexport=""} of `type`{.variable} for `target`{.variable}
with `name`{.variable}, `namespace`{.variable}, `oldValue`{.variable},
`addedNodes`{.variable}, `removedNodes`{.variable},
`previousSibling`{.variable}, and `nextSibling`{.variable}, run these
steps:

1.  Let `interestedObservers`{.variable} be an empty
    [map](https://infra.spec.whatwg.org/#ordered-map){#ref-for-ordered-map
    link-type="dfn"}.

2.  Let `nodes`{.variable} be the [inclusive
    ancestors](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor③
    link-type="dfn"} of `target`{.variable}.

3.  For each `node`{.variable} in `nodes`{.variable}, and then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②③
    link-type="dfn"} `registered`{.variable} of `node`{.variable}'s
    [registered observer
    list](#registered-observer-list){#ref-for-registered-observer-list⑦
    link-type="dfn"}:

    1.  Let `options`{.variable} be `registered`{.variable}'s
        [options](#registered-observer-options){#ref-for-registered-observer-options⑤
        link-type="dfn"}.

    2.  If none of the following are true

        - `node`{.variable} is not `target`{.variable} and
          `options`{.variable}\[\"[`subtree`{.idl}](#dom-mutationobserverinit-subtree){#ref-for-dom-mutationobserverinit-subtree③
          link-type="idl"}\"\] is false
        - `type`{.variable} is \"`attributes`\" and
          `options`{.variable}\[\"[`attributes`{.idl}](#dom-mutationobserverinit-attributes){#ref-for-dom-mutationobserverinit-attributes⑧
          link-type="idl"}\"\] either does not
          [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑦
          link-type="dfn"} or is false
        - `type`{.variable} is \"`attributes`\",
          `options`{.variable}\[\"[`attributeFilter`{.idl}](#dom-mutationobserverinit-attributefilter){#ref-for-dom-mutationobserverinit-attributefilter④
          link-type="idl"}\"\]
          [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑧
          link-type="dfn"}, and
          `options`{.variable}\[\"[`attributeFilter`{.idl}](#dom-mutationobserverinit-attributefilter){#ref-for-dom-mutationobserverinit-attributefilter⑤
          link-type="idl"}\"\] does not
          [contain](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain④
          link-type="dfn"} `name`{.variable} or `namespace`{.variable}
          is non-null
        - `type`{.variable} is \"`characterData`\" and
          `options`{.variable}\[\"[`characterData`{.idl}](#dom-mutationobserverinit-characterdata){#ref-for-dom-mutationobserverinit-characterdata⑥
          link-type="idl"}\"\] either does not
          [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists⑨
          link-type="dfn"} or is false
        - `type`{.variable} is \"`childList`\" and
          `options`{.variable}\[\"[`childList`{.idl}](#dom-mutationobserverinit-childlist){#ref-for-dom-mutationobserverinit-childlist②
          link-type="idl"}\"\] is false

        then:

        1.  Let `mo`{.variable} be `registered`{.variable}'s
            [observer](#registered-observer-observer){#ref-for-registered-observer-observer⑥
            link-type="dfn"}.

        2.  If `interestedObservers`{.variable}\[`mo`{.variable}\] does
            not
            [exist](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①⓪
            link-type="dfn"}, then
            [set](https://infra.spec.whatwg.org/#map-set){#ref-for-map-set
            link-type="dfn"}
            `interestedObservers`{.variable}\[`mo`{.variable}\] to null.

        3.  If either `type`{.variable} is \"`attributes`\" and
            `options`{.variable}\[\"[`attributeOldValue`{.idl}](#dom-mutationobserverinit-attributeoldvalue){#ref-for-dom-mutationobserverinit-attributeoldvalue④
            link-type="idl"}\"\] is true, or `type`{.variable} is
            \"`characterData`\" and
            `options`{.variable}\[\"[`characterDataOldValue`{.idl}](#dom-mutationobserverinit-characterdataoldvalue){#ref-for-dom-mutationobserverinit-characterdataoldvalue④
            link-type="idl"}\"\] is true, then
            [set](https://infra.spec.whatwg.org/#map-set){#ref-for-map-set①
            link-type="dfn"}
            `interestedObservers`{.variable}\[`mo`{.variable}\] to
            `oldValue`{.variable}.

4.  [For
    each](https://infra.spec.whatwg.org/#map-iterate){#ref-for-map-iterate①
    link-type="dfn"} `observer`{.variable} → `mappedOldValue`{.variable}
    of `interestedObservers`{.variable}:

    1.  Let `record`{.variable} be a new
        [`MutationRecord`{.idl}](#mutationrecord){#ref-for-mutationrecord④
        link-type="idl"} object with its
        [`type`{.idl}](#dom-mutationrecord-type){#ref-for-dom-mutationrecord-type
        link-type="idl"} set to `type`{.variable},
        [`target`{.idl}](#dom-mutationrecord-target){#ref-for-dom-mutationrecord-target
        link-type="idl"} set to `target`{.variable},
        [`attributeName`{.idl}](#dom-mutationrecord-attributename){#ref-for-dom-mutationrecord-attributename
        link-type="idl"} set to `name`{.variable},
        [`attributeNamespace`{.idl}](#dom-mutationrecord-attributenamespace){#ref-for-dom-mutationrecord-attributenamespace
        link-type="idl"} set to `namespace`{.variable},
        [`oldValue`{.idl}](#dom-mutationrecord-oldvalue){#ref-for-dom-mutationrecord-oldvalue
        link-type="idl"} set to `mappedOldValue`{.variable},
        [`addedNodes`{.idl}](#dom-mutationrecord-addednodes){#ref-for-dom-mutationrecord-addednodes
        link-type="idl"} set to `addedNodes`{.variable},
        [`removedNodes`{.idl}](#dom-mutationrecord-removednodes){#ref-for-dom-mutationrecord-removednodes
        link-type="idl"} set to `removedNodes`{.variable},
        [`previousSibling`{.idl}](#dom-mutationrecord-previoussibling){#ref-for-dom-mutationrecord-previoussibling
        link-type="idl"} set to `previousSibling`{.variable}, and
        [`nextSibling`{.idl}](#dom-mutationrecord-nextsibling){#ref-for-dom-mutationrecord-nextsibling
        link-type="idl"} set to `nextSibling`{.variable}.

    2.  [Enqueue](https://infra.spec.whatwg.org/#queue-enqueue){#ref-for-queue-enqueue
        link-type="dfn"} `record`{.variable} to `observer`{.variable}'s
        [record queue](#concept-mo-queue){#ref-for-concept-mo-queue⑥
        link-type="dfn"}.

    3.  [Append](https://infra.spec.whatwg.org/#set-append){#ref-for-set-append①①
        link-type="dfn"} `observer`{.variable} to the [surrounding
        agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent⑦
        link-type="dfn"}'s [pending mutation
        observers](#mutation-observer-list){#ref-for-mutation-observer-list②
        link-type="dfn"}.

5.  [Queue a mutation observer
    microtask](#queue-a-mutation-observer-compound-microtask){#ref-for-queue-a-mutation-observer-compound-microtask①
    link-type="dfn"}.

To [queue a tree mutation record]{#queue-a-tree-mutation-record .dfn
.dfn-paneled dfn-type="dfn" noexport=""} for `target`{.variable} with
`addedNodes`{.variable}, `removedNodes`{.variable},
`previousSibling`{.variable}, and `nextSibling`{.variable}, run these
steps:

1.  Assert: either `addedNodes`{.variable} or `removedNodes`{.variable}
    [is not
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty⑤
    link-type="dfn"}.

2.  [Queue a mutation
    record](#queue-a-mutation-record){#ref-for-queue-a-mutation-record
    link-type="dfn"} of \"`childList`\" for `target`{.variable} with
    null, null, null, `addedNodes`{.variable},
    `removedNodes`{.variable}, `previousSibling`{.variable}, and
    `nextSibling`{.variable}.

#### [4.3.3. ]{.secno}[Interface [`MutationRecord`{.idl}](#mutationrecord){#ref-for-mutationrecord⑤ link-type="idl"}]{.content}[](#interface-mutationrecord){.self-link} {#interface-mutationrecord .heading .settled level="4.3.3"}

``` {.idl .highlight .def}
[Exposed=Window]
interface MutationRecord {
  readonly attribute DOMString type;
  [SameObject] readonly attribute Node target;
  [SameObject] readonly attribute NodeList addedNodes;
  [SameObject] readonly attribute NodeList removedNodes;
  readonly attribute Node? previousSibling;
  readonly attribute Node? nextSibling;
  readonly attribute DOMString? attributeName;
  readonly attribute DOMString? attributeNamespace;
  readonly attribute DOMString? oldValue;
};
```

`record`{.variable}` . `[`type`{.idl}](#dom-mutationrecord-type){#ref-for-dom-mutationrecord-type② link-type="idl"}
:   Returns \"`attributes`\" if it was an
    [attribute](#concept-attribute){#ref-for-concept-attribute④
    link-type="dfn"} mutation. \"`characterData`\" if it was a mutation
    to a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①⓪
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑨②
    link-type="dfn"}. And \"`childList`\" if it was a mutation to the
    [tree](#concept-tree){#ref-for-concept-tree①⑧ link-type="dfn"} of
    [nodes](#concept-node){#ref-for-concept-node⑨③ link-type="dfn"}.

`record`{.variable}` . `[`target`{.idl}](#dom-mutationrecord-target){#ref-for-dom-mutationrecord-target② link-type="idl"}
:   Returns the [node](#concept-node){#ref-for-concept-node⑨④
    link-type="dfn"} the mutation affected, depending on the
    [`type`{.idl}](#dom-mutationrecord-type){#ref-for-dom-mutationrecord-type③
    link-type="idl"}. For \"`attributes`\", it is the
    [element](#concept-element){#ref-for-concept-element④⑦
    link-type="dfn"} whose
    [attribute](#concept-attribute){#ref-for-concept-attribute⑤
    link-type="dfn"} changed. For \"`characterData`\", it is the
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①①
    link-type="idl"} [node](#concept-node){#ref-for-concept-node⑨⑤
    link-type="dfn"}. For \"`childList`\", it is the
    [node](#concept-node){#ref-for-concept-node⑨⑥ link-type="dfn"} whose
    [children](#concept-tree-child){#ref-for-concept-tree-child④⑨
    link-type="dfn"} changed.

`record`{.variable}` . `[`addedNodes`{.idl}](#dom-mutationrecord-addednodes){#ref-for-dom-mutationrecord-addednodes② link-type="idl"}\
`record`{.variable}` . `[`removedNodes`{.idl}](#dom-mutationrecord-removednodes){#ref-for-dom-mutationrecord-removednodes② link-type="idl"}
:   Return the [nodes](#concept-node){#ref-for-concept-node⑨⑦
    link-type="dfn"} added and removed respectively.

`record`{.variable}` . `[`previousSibling`{.idl}](#dom-mutationrecord-previoussibling){#ref-for-dom-mutationrecord-previoussibling② link-type="idl"}\
`record`{.variable}` . `[`nextSibling`{.idl}](#dom-mutationrecord-nextsibling){#ref-for-dom-mutationrecord-nextsibling② link-type="idl"}
:   Return the
    [previous](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling⑤
    link-type="dfn"} and [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling⑦
    link-type="dfn"} respectively of the added or removed
    [nodes](#concept-node){#ref-for-concept-node⑨⑧ link-type="dfn"};
    otherwise null.

`record`{.variable}` . `[`attributeName`{.idl}](#dom-mutationrecord-attributename){#ref-for-dom-mutationrecord-attributename② link-type="idl"}
:   Returns the [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①
    link-type="dfn"} of the changed
    [attribute](#concept-attribute){#ref-for-concept-attribute⑥
    link-type="dfn"}; otherwise null.

`record`{.variable}` . `[`attributeNamespace`{.idl}](#dom-mutationrecord-attributenamespace){#ref-for-dom-mutationrecord-attributenamespace② link-type="idl"}
:   Returns the
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①
    link-type="dfn"} of the changed
    [attribute](#concept-attribute){#ref-for-concept-attribute⑦
    link-type="dfn"}; otherwise null.

`record`{.variable}` . `[`oldValue`{.idl}](#dom-mutationrecord-oldvalue){#ref-for-dom-mutationrecord-oldvalue② link-type="idl"}
:   The return value depends on
    [`type`{.idl}](#dom-mutationrecord-type){#ref-for-dom-mutationrecord-type④
    link-type="idl"}. For \"`attributes`\", it is the
    [value](#concept-attribute-value){#ref-for-concept-attribute-value④
    link-type="dfn"} of the changed
    [attribute](#concept-attribute){#ref-for-concept-attribute⑧
    link-type="dfn"} before the change. For \"`characterData`\", it is
    the [data](#concept-cd-data){#ref-for-concept-cd-data④
    link-type="dfn"} of the changed
    [node](#concept-node){#ref-for-concept-node⑨⑨ link-type="dfn"}
    before the change. For \"`childList`\", it is null.

The [`type`]{#dom-mutationrecord-type .dfn .dfn-paneled .idl-code
dfn-for="MutationRecord" dfn-type="attribute" export=""},
[`target`]{#dom-mutationrecord-target .dfn .dfn-paneled .idl-code
dfn-for="MutationRecord" dfn-type="attribute" export=""},
[`addedNodes`]{#dom-mutationrecord-addednodes .dfn .dfn-paneled
.idl-code dfn-for="MutationRecord" dfn-type="attribute" export=""},
[`removedNodes`]{#dom-mutationrecord-removednodes .dfn .dfn-paneled
.idl-code dfn-for="MutationRecord" dfn-type="attribute" export=""},
[`previousSibling`]{#dom-mutationrecord-previoussibling .dfn
.dfn-paneled .idl-code dfn-for="MutationRecord" dfn-type="attribute"
export=""}, [`nextSibling`]{#dom-mutationrecord-nextsibling .dfn
.dfn-paneled .idl-code dfn-for="MutationRecord" dfn-type="attribute"
export=""}, [`attributeName`]{#dom-mutationrecord-attributename .dfn
.dfn-paneled .idl-code dfn-for="MutationRecord" dfn-type="attribute"
export=""},
[`attributeNamespace`]{#dom-mutationrecord-attributenamespace .dfn
.dfn-paneled .idl-code dfn-for="MutationRecord" dfn-type="attribute"
export=""}, and [`oldValue`]{#dom-mutationrecord-oldvalue .dfn
.dfn-paneled .idl-code dfn-for="MutationRecord" dfn-type="attribute"
export=""} attributes must return the values they were initialized to.

### [4.4. ]{.secno}[Interface [`Node`{.idl}](#node){#ref-for-node①⑦ link-type="idl"}]{.content}[](#interface-node){.self-link} {#interface-node .heading .settled level="4.4"}

``` {.idl .highlight .def}
[Exposed=Window]
interface Node : EventTarget {
  const unsigned short ELEMENT_NODE = 1;
  const unsigned short ATTRIBUTE_NODE = 2;
  const unsigned short TEXT_NODE = 3;
  const unsigned short CDATA_SECTION_NODE = 4;
  const unsigned short ENTITY_REFERENCE_NODE = 5; // legacy
  const unsigned short ENTITY_NODE = 6; // legacy
  const unsigned short PROCESSING_INSTRUCTION_NODE = 7;
  const unsigned short COMMENT_NODE = 8;
  const unsigned short DOCUMENT_NODE = 9;
  const unsigned short DOCUMENT_TYPE_NODE = 10;
  const unsigned short DOCUMENT_FRAGMENT_NODE = 11;
  const unsigned short NOTATION_NODE = 12; // legacy
  readonly attribute unsigned short nodeType;
  readonly attribute DOMString nodeName;

  readonly attribute USVString baseURI;

  readonly attribute boolean isConnected;
  readonly attribute Document? ownerDocument;
  Node getRootNode(optional GetRootNodeOptions options = {});
  readonly attribute Node? parentNode;
  readonly attribute Element? parentElement;
  boolean hasChildNodes();
  [SameObject] readonly attribute NodeList childNodes;
  readonly attribute Node? firstChild;
  readonly attribute Node? lastChild;
  readonly attribute Node? previousSibling;
  readonly attribute Node? nextSibling;

  [CEReactions] attribute DOMString? nodeValue;
  [CEReactions] attribute DOMString? textContent;
  [CEReactions] undefined normalize();

  [CEReactions, NewObject] Node cloneNode(optional boolean subtree = false);
  boolean isEqualNode(Node? otherNode);
  boolean isSameNode(Node? otherNode); // legacy alias of ===

  const unsigned short DOCUMENT_POSITION_DISCONNECTED = 0x01;
  const unsigned short DOCUMENT_POSITION_PRECEDING = 0x02;
  const unsigned short DOCUMENT_POSITION_FOLLOWING = 0x04;
  const unsigned short DOCUMENT_POSITION_CONTAINS = 0x08;
  const unsigned short DOCUMENT_POSITION_CONTAINED_BY = 0x10;
  const unsigned short DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 0x20;
  unsigned short compareDocumentPosition(Node other);
  boolean contains(Node? other);

  DOMString? lookupPrefix(DOMString? namespace);
  DOMString? lookupNamespaceURI(DOMString? prefix);
  boolean isDefaultNamespace(DOMString? namespace);

  [CEReactions] Node insertBefore(Node node, Node? child);
  [CEReactions] Node appendChild(Node node);
  [CEReactions] Node replaceChild(Node node, Node child);
  [CEReactions] Node removeChild(Node child);
};

dictionary GetRootNodeOptions {
  boolean composed = false;
};
```

[`Node`{.idl}](#node){#ref-for-node③⑨ link-type="idl"} is an abstract
interface that is used by all
[nodes](#concept-node){#ref-for-concept-node①⓪⓪ link-type="dfn"}. You
cannot get a direct instance of it.

Each [node](#concept-node){#ref-for-concept-node①⓪① link-type="dfn"} has
an associated [node document]{#concept-node-document .dfn .dfn-paneled
dfn-for="Node" dfn-type="dfn" export=""}, set upon creation, that is a
[document](#concept-document){#ref-for-concept-document①⑧
link-type="dfn"}.

A [node](#concept-node){#ref-for-concept-node①⓪② link-type="dfn"}'s
[node document](#concept-node-document){#ref-for-concept-node-document①⑧
link-type="dfn"} can be changed by the
[adopt](#concept-node-adopt){#ref-for-concept-node-adopt①
link-type="dfn"} algorithm.

A [node](#concept-node){#ref-for-concept-node①⓪③ link-type="dfn"}'s [get
the parent](#get-the-parent){#ref-for-get-the-parent⑤ link-type="dfn"}
algorithm, given an `event`{.variable}, returns the
[node](#concept-node){#ref-for-concept-node①⓪④ link-type="dfn"}'s
[assigned
slot](#slotable-assigned-slot){#ref-for-slotable-assigned-slot⑤
link-type="dfn"}, if [node](#concept-node){#ref-for-concept-node①⓪⑤
link-type="dfn"} is
[assigned](#slotable-assigned){#ref-for-slotable-assigned⑤
link-type="dfn"}; otherwise
[node](#concept-node){#ref-for-concept-node①⓪⑥ link-type="dfn"}'s
[parent](#concept-tree-parent){#ref-for-concept-tree-parent②④
link-type="dfn"}.

Each [node](#concept-node){#ref-for-concept-node①⓪⑦ link-type="dfn"}
also has a [registered observer
list](#registered-observer-list){#ref-for-registered-observer-list⑧
link-type="dfn"}.

------------------------------------------------------------------------

`node`{.variable}` . `[`nodeType`{.idl}](#dom-node-nodetype){#ref-for-dom-node-nodetype① link-type="idl"}

:   Returns a number appropriate for the type of `node`{.variable}, as
    follows:

    [`Element`{.idl}](#element){#ref-for-element③③ link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④⓪
        link-type="idl"}` . `[`ELEMENT_NODE`{.idl}](#dom-node-element_node){#ref-for-dom-node-element_node①
        link-type="idl"} (1).

    [`Attr`{.idl}](#attr){#ref-for-attr⑤ link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④①
        link-type="idl"}` . `[`ATTRIBUTE_NODE`{.idl}](#dom-node-attribute_node){#ref-for-dom-node-attribute_node①
        link-type="idl"} (2).

    An [exclusive `Text` node](#exclusive-text-node){#ref-for-exclusive-text-node link-type="dfn"}
    :   [`Node`{.idl}](#node){#ref-for-node④②
        link-type="idl"}` . `[`TEXT_NODE`{.idl}](#dom-node-text_node){#ref-for-dom-node-text_node①
        link-type="idl"} (3).

    [`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection② link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④③
        link-type="idl"}` . `[`CDATA_SECTION_NODE`{.idl}](#dom-node-cdata_section_node){#ref-for-dom-node-cdata_section_node①
        link-type="idl"} (4).

    [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction⑤ link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④④
        link-type="idl"}` . `[`PROCESSING_INSTRUCTION_NODE`{.idl}](#dom-node-processing_instruction_node){#ref-for-dom-node-processing_instruction_node①
        link-type="idl"} (7).

    [`Comment`{.idl}](#comment){#ref-for-comment⑤ link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④⑤
        link-type="idl"}` . `[`COMMENT_NODE`{.idl}](#dom-node-comment_node){#ref-for-dom-node-comment_node①
        link-type="idl"} (8).

    [`Document`{.idl}](#document){#ref-for-document⑨ link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④⑥
        link-type="idl"}` . `[`DOCUMENT_NODE`{.idl}](#dom-node-document_node){#ref-for-dom-node-document_node①
        link-type="idl"} (9).

    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①⓪ link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④⑦
        link-type="idl"}` . `[`DOCUMENT_TYPE_NODE`{.idl}](#dom-node-document_type_node){#ref-for-dom-node-document_type_node①
        link-type="idl"} (10).

    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①⑦ link-type="idl"}
    :   [`Node`{.idl}](#node){#ref-for-node④⑧
        link-type="idl"}` . `[`DOCUMENT_FRAGMENT_NODE`{.idl}](#dom-node-document_fragment_node){#ref-for-dom-node-document_fragment_node①
        link-type="idl"} (11).

`node`{.variable}` . `[`nodeName`{.idl}](#dom-node-nodename){#ref-for-dom-node-nodename① link-type="idl"}

:   Returns a string appropriate for the type of `node`{.variable}, as
    follows:

    [`Element`{.idl}](#element){#ref-for-element③④ link-type="idl"}
    :   Its [HTML-uppercased qualified
        name](#element-html-uppercased-qualified-name){#ref-for-element-html-uppercased-qualified-name
        link-type="dfn"}.

    [`Attr`{.idl}](#attr){#ref-for-attr⑥ link-type="idl"}
    :   Its [qualified
        name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name
        link-type="dfn"}.

    An [exclusive `Text` node](#exclusive-text-node){#ref-for-exclusive-text-node① link-type="dfn"}
    :   \"`#text`\".

    [`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection③ link-type="idl"}
    :   \"`#cdata-section`\".

    [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction⑥ link-type="idl"}
    :   Its [target](#concept-pi-target){#ref-for-concept-pi-target
        link-type="dfn"}.

    [`Comment`{.idl}](#comment){#ref-for-comment⑥ link-type="idl"}
    :   \"`#comment`\".

    [`Document`{.idl}](#document){#ref-for-document①⓪ link-type="idl"}
    :   \"`#document`\".

    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①① link-type="idl"}
    :   Its [name](#concept-doctype-name){#ref-for-concept-doctype-name
        link-type="dfn"}.

    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①⑧ link-type="idl"}
    :   \"`#document-fragment`\".

The [`nodeType`]{#dom-node-nodetype .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return the first matching statement, switching on the interface
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑤
link-type="dfn"}
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements⑤
link-type="dfn"}:

[`Element`{.idl}](#element){#ref-for-element③⑤ link-type="idl"}
:   [`ELEMENT_NODE`]{#dom-node-element_node .dfn .dfn-paneled .idl-code
    dfn-for="Node" dfn-type="const" export=""} (1)

[`Attr`{.idl}](#attr){#ref-for-attr⑦ link-type="idl"}
:   [`ATTRIBUTE_NODE`]{#dom-node-attribute_node .dfn .dfn-paneled
    .idl-code dfn-for="Node" dfn-type="const" export=""} (2);

An [exclusive `Text` node](#exclusive-text-node){#ref-for-exclusive-text-node② link-type="dfn"}
:   [`TEXT_NODE`]{#dom-node-text_node .dfn .dfn-paneled .idl-code
    dfn-for="Node" dfn-type="const" export=""} (3);

[`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection④ link-type="idl"}
:   [`CDATA_SECTION_NODE`]{#dom-node-cdata_section_node .dfn
    .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
    (4);

[`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction⑦ link-type="idl"}
:   [`PROCESSING_INSTRUCTION_NODE`]{#dom-node-processing_instruction_node
    .dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="const"
    export=""} (7);

[`Comment`{.idl}](#comment){#ref-for-comment⑦ link-type="idl"}
:   [`COMMENT_NODE`]{#dom-node-comment_node .dfn .dfn-paneled .idl-code
    dfn-for="Node" dfn-type="const" export=""} (8);

[`Document`{.idl}](#document){#ref-for-document①① link-type="idl"}
:   [`DOCUMENT_NODE`]{#dom-node-document_node .dfn .dfn-paneled
    .idl-code dfn-for="Node" dfn-type="const" export=""} (9);

[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①② link-type="idl"}
:   [`DOCUMENT_TYPE_NODE`]{#dom-node-document_type_node .dfn
    .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
    (10);

[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment①⑨ link-type="idl"}
:   [`DOCUMENT_FRAGMENT_NODE`]{#dom-node-document_fragment_node .dfn
    .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
    (11).

The [`nodeName`]{#dom-node-nodename .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return the first matching statement, switching on the interface
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑥
link-type="dfn"}
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements⑥
link-type="dfn"}:

[`Element`{.idl}](#element){#ref-for-element③⑥ link-type="idl"}
:   Its [HTML-uppercased qualified
    name](#element-html-uppercased-qualified-name){#ref-for-element-html-uppercased-qualified-name①
    link-type="dfn"}.

[`Attr`{.idl}](#attr){#ref-for-attr⑧ link-type="idl"}
:   Its [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name①
    link-type="dfn"}.

An [exclusive `Text` node](#exclusive-text-node){#ref-for-exclusive-text-node③ link-type="dfn"}
:   \"`#text`\".

[`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection⑤ link-type="idl"}
:   \"`#cdata-section`\".

[`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction⑧ link-type="idl"}
:   Its [target](#concept-pi-target){#ref-for-concept-pi-target①
    link-type="dfn"}.

[`Comment`{.idl}](#comment){#ref-for-comment⑧ link-type="idl"}
:   \"`#comment`\".

[`Document`{.idl}](#document){#ref-for-document①② link-type="idl"}
:   \"`#document`\".

[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①③ link-type="idl"}
:   Its [name](#concept-doctype-name){#ref-for-concept-doctype-name①
    link-type="dfn"}.

[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②⓪ link-type="idl"}
:   \"`#document-fragment`\".

------------------------------------------------------------------------

`node`{.variable}` . `[`baseURI`{.idl}](#dom-node-baseuri){#ref-for-dom-node-baseuri① link-type="idl"}
:   Returns `node`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document①⑨
    link-type="dfn"}'s [document base
    URL](https://html.spec.whatwg.org/multipage/urls-and-fetching.html#document-base-url){#ref-for-document-base-url
    link-type="dfn"}.

The [`baseURI`]{#dom-node-baseuri .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑦
link-type="dfn"}'s [node
document](#concept-node-document){#ref-for-concept-node-document②⓪
link-type="dfn"}'s [document base
URL](https://html.spec.whatwg.org/multipage/urls-and-fetching.html#document-base-url){#ref-for-document-base-url①
link-type="dfn"},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer
link-type="dfn"}.

------------------------------------------------------------------------

`node`{.variable}` . `[`isConnected`{.idl}](#dom-node-isconnected){#ref-for-dom-node-isconnected① link-type="idl"}

:   Returns true if `node`{.variable} is
    [connected](#connected){#ref-for-connected⑥ link-type="dfn"};
    otherwise false.

`node`{.variable}` . `[`ownerDocument`{.idl}](#dom-node-ownerdocument){#ref-for-dom-node-ownerdocument① link-type="idl"}
:   Returns the [node
    document](#concept-node-document){#ref-for-concept-node-document②①
    link-type="dfn"}. Returns null for
    [documents](#concept-document){#ref-for-concept-document①⑨
    link-type="dfn"}.

`node`{.variable}` . `[`getRootNode()`{.idl}](#dom-node-getrootnode){#ref-for-dom-node-getrootnode① link-type="idl"}
:   Returns `node`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root②③
    link-type="dfn"}.

`node`{.variable}` . `[`getRootNode`](#dom-node-getrootnode){#ref-for-dom-node-getrootnode② link-type="idl"}`({ composed:true })`
:   Returns `node`{.variable}'s [shadow-including
    root](#concept-shadow-including-root){#ref-for-concept-shadow-including-root③
    link-type="dfn"}.

`node`{.variable}` . `[`parentNode`{.idl}](#dom-node-parentnode){#ref-for-dom-node-parentnode① link-type="idl"}
:   Returns the
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent②⑤
    link-type="dfn"}.

`node`{.variable}` . `[`parentElement`{.idl}](#dom-node-parentelement){#ref-for-dom-node-parentelement① link-type="idl"}
:   Returns the [parent
    element](#parent-element){#ref-for-parent-element link-type="dfn"}.

`node`{.variable}` . `[`hasChildNodes()`{.idl}](#dom-node-haschildnodes){#ref-for-dom-node-haschildnodes① link-type="idl"}
:   Returns whether `node`{.variable} has
    [children](#concept-tree-child){#ref-for-concept-tree-child⑤⓪
    link-type="dfn"}.

`node`{.variable}` . `[`childNodes`{.idl}](#dom-node-childnodes){#ref-for-dom-node-childnodes① link-type="idl"}
:   Returns the
    [children](#concept-tree-child){#ref-for-concept-tree-child⑤①
    link-type="dfn"}.

`node`{.variable}` . `[`firstChild`{.idl}](#dom-node-firstchild){#ref-for-dom-node-firstchild① link-type="idl"}
:   Returns the [first
    child](#concept-tree-first-child){#ref-for-concept-tree-first-child③
    link-type="dfn"}.

`node`{.variable}` . `[`lastChild`{.idl}](#dom-node-lastchild){#ref-for-dom-node-lastchild① link-type="idl"}
:   Returns the [last
    child](#concept-tree-last-child){#ref-for-concept-tree-last-child④
    link-type="dfn"}.

`node`{.variable}` . `[`previousSibling`{.idl}](#dom-node-previoussibling){#ref-for-dom-node-previoussibling① link-type="idl"}
:   Returns the [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling⑥
    link-type="dfn"}.

`node`{.variable}` . `[`nextSibling`{.idl}](#dom-node-nextsibling){#ref-for-dom-node-nextsibling① link-type="idl"}
:   Returns the [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling⑧
    link-type="dfn"}.

The [`isConnected`]{#dom-node-isconnected .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return true, if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑧
link-type="dfn"} is [connected](#connected){#ref-for-connected⑦
link-type="dfn"}; otherwise false.

The [`ownerDocument`]{#dom-node-ownerdocument .dfn .dfn-paneled
.idl-code dfn-for="Node" dfn-type="attribute" export=""} getter steps
are to return null, if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑦⑨
link-type="dfn"} is a
[document](#concept-document){#ref-for-concept-document②⓪
link-type="dfn"}; otherwise
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⓪
link-type="dfn"}'s [node
document](#concept-node-document){#ref-for-concept-node-document②②
link-type="dfn"}.

The [node
document](#concept-node-document){#ref-for-concept-node-document②③
link-type="dfn"} of a
[document](#concept-document){#ref-for-concept-document②①
link-type="dfn"} is that
[document](#concept-document){#ref-for-concept-document②②
link-type="dfn"} itself. All
[nodes](#concept-node){#ref-for-concept-node①⓪⑧ link-type="dfn"} have a
[node document](#concept-node-document){#ref-for-concept-node-document②④
link-type="dfn"} at all times.

The [`getRootNode(``options`{.variable}`)`]{#dom-node-getrootnode .dfn
.dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""
lt="getRootNode(options)|getRootNode()"} method steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧①
link-type="dfn"}'s [shadow-including
root](#concept-shadow-including-root){#ref-for-concept-shadow-including-root④
link-type="dfn"} if
`options`{.variable}\[\"[`composed`{.idl}](#dom-getrootnodeoptions-composed){#ref-for-dom-getrootnodeoptions-composed
link-type="idl"}\"\] is true; otherwise
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧②
link-type="dfn"}'s
[root](#concept-tree-root){#ref-for-concept-tree-root②④
link-type="dfn"}.

The [`parentNode`]{#dom-node-parentnode .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧③
link-type="dfn"}'s
[parent](#concept-tree-parent){#ref-for-concept-tree-parent②⑥
link-type="dfn"}.

The [`parentElement`]{#dom-node-parentelement .dfn .dfn-paneled
.idl-code dfn-for="Node" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧④
link-type="dfn"}'s [parent
element](#parent-element){#ref-for-parent-element① link-type="dfn"}.

The [`hasChildNodes()`]{#dom-node-haschildnodes .dfn .dfn-paneled
.idl-code dfn-for="Node" dfn-type="method" export=""} method steps are
to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑤
link-type="dfn"} has
[children](#concept-tree-child){#ref-for-concept-tree-child⑤②
link-type="dfn"}; otherwise false.

The [`childNodes`]{#dom-node-childnodes .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return a [`NodeList`{.idl}](#nodelist){#ref-for-nodelist⑦
link-type="idl"} rooted at
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑥
link-type="dfn"} matching only
[children](#concept-tree-child){#ref-for-concept-tree-child⑤③
link-type="dfn"}.

The [`firstChild` ]{#dom-node-firstchild .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export="" lt="firstChild"} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑦
link-type="dfn"}'s [first
child](#concept-tree-first-child){#ref-for-concept-tree-first-child④
link-type="dfn"}.

The [`lastChild`]{#dom-node-lastchild .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑧
link-type="dfn"}'s [last
child](#concept-tree-last-child){#ref-for-concept-tree-last-child⑤
link-type="dfn"}.

The [`previousSibling`]{#dom-node-previoussibling .dfn .dfn-paneled
.idl-code dfn-for="Node" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑧⑨
link-type="dfn"}'s [previous
sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling⑦
link-type="dfn"}.

The [`nextSibling`]{#dom-node-nextsibling .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⓪
link-type="dfn"}'s [next
sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling⑨
link-type="dfn"}.

------------------------------------------------------------------------

The [`nodeValue`]{#dom-node-nodevalue .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return the following, switching on the interface
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨①
link-type="dfn"}
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements⑦
link-type="dfn"}:

[`Attr`{.idl}](#attr){#ref-for-attr⑨ link-type="idl"}
:   [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨②
    link-type="dfn"}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value⑤
    link-type="dfn"}.

[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①② link-type="idl"}
:   [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨③
    link-type="dfn"}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data⑤ link-type="dfn"}.

Otherwise
:   Null.

The
[`nodeValue`{.idl}](#dom-node-nodevalue){#ref-for-dom-node-nodevalue①
link-type="idl"} setter steps are to, if the given value is null, act as
if it was the empty string instead, and then do as described below,
switching on the interface
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨④
link-type="dfn"}
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements⑧
link-type="dfn"}:

[`Attr`{.idl}](#attr){#ref-for-attr①⓪ link-type="idl"}

:   [Set an existing attribute
    value](#set-an-existing-attribute-value){#ref-for-set-an-existing-attribute-value
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑤
    link-type="dfn"} and the given value.

[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①③ link-type="idl"}

:   [Replace data](#concept-cd-replace){#ref-for-concept-cd-replace①
    link-type="dfn"} with node
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑥
    link-type="dfn"}, offset 0, count
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑦
    link-type="dfn"}'s
    [length](#concept-node-length){#ref-for-concept-node-length①
    link-type="dfn"}, and data the given value.

Otherwise

:   Do nothing.

To [get text content]{#get-text-content .dfn .dfn-paneled dfn-type="dfn"
export=""} with a [node](#concept-node){#ref-for-concept-node①⓪⑨
link-type="dfn"} `node`{.variable}, return the following, switching on
the interface `node`{.variable}
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements⑨
link-type="dfn"}:

[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②① link-type="idl"}\
[`Element`{.idl}](#element){#ref-for-element③⑦ link-type="idl"}
:   The [descendant text
    content](#concept-descendant-text-content){#ref-for-concept-descendant-text-content
    link-type="dfn"} of `node`{.variable}.

[`Attr`{.idl}](#attr){#ref-for-attr①① link-type="idl"}
:   `node`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value⑥
    link-type="dfn"}.

[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①④ link-type="idl"}
:   `node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data⑥ link-type="dfn"}.

Otherwise
:   Null.

The [`textContent`]{#dom-node-textcontent .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="attribute" export=""} getter steps are to
return the result of running [get text
content](#get-text-content){#ref-for-get-text-content link-type="dfn"}
with [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑧
link-type="dfn"}.

To [string replace all]{#string-replace-all .dfn .dfn-paneled
dfn-type="dfn" export=""} with a string `string`{.variable} within a
[node](#concept-node){#ref-for-concept-node①①⓪ link-type="dfn"}
`parent`{.variable}, run these steps:

1.  Let `node`{.variable} be null.

2.  If `string`{.variable} is not the empty string, then set
    `node`{.variable} to a new [`Text`{.idl}](#text){#ref-for-text②①
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①①①
    link-type="dfn"} whose
    [data](#concept-cd-data){#ref-for-concept-cd-data⑦ link-type="dfn"}
    is `string`{.variable} and [node
    document](#concept-node-document){#ref-for-concept-node-document②⑤
    link-type="dfn"} is `parent`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document②⑥
    link-type="dfn"}.

3.  [Replace
    all](#concept-node-replace-all){#ref-for-concept-node-replace-all①
    link-type="dfn"} with `node`{.variable} within `parent`{.variable}.

To [set text content]{#set-text-content .dfn .dfn-paneled dfn-type="dfn"
export=""} with a [node](#concept-node){#ref-for-concept-node①①②
link-type="dfn"} `node`{.variable} and a string `value`{.variable}, do
as defined below, switching on the interface `node`{.variable}
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①⓪
link-type="dfn"}:

[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②② link-type="idl"}\
[`Element`{.idl}](#element){#ref-for-element③⑧ link-type="idl"}

:   [String replace
    all](#string-replace-all){#ref-for-string-replace-all
    link-type="dfn"} with `value`{.variable} within `node`{.variable}.

[`Attr`{.idl}](#attr){#ref-for-attr①② link-type="idl"}

:   [Set an existing attribute
    value](#set-an-existing-attribute-value){#ref-for-set-an-existing-attribute-value①
    link-type="dfn"} with `node`{.variable} and `value`{.variable}.

[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①⑤ link-type="idl"}

:   [Replace data](#concept-cd-replace){#ref-for-concept-cd-replace②
    link-type="dfn"} with node `node`{.variable}, offset 0, count
    `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length②
    link-type="dfn"}, and data `value`{.variable}.

Otherwise

:   Do nothing.

The
[`textContent`{.idl}](#dom-node-textcontent){#ref-for-dom-node-textcontent①
link-type="idl"} setter steps are to, if the given value is null, act as
if it was the empty string instead, and then run [set text
content](#set-text-content){#ref-for-set-text-content link-type="dfn"}
with [this](https://webidl.spec.whatwg.org/#this){#ref-for-this⑨⑨
link-type="dfn"} and the given value.

------------------------------------------------------------------------

`node`{.variable}` . `[`normalize()`{.idl}](#dom-node-normalize){#ref-for-dom-node-normalize① link-type="idl"}
:   Removes [empty](#concept-node-empty){#ref-for-concept-node-empty
    link-type="dfn"} [exclusive `Text`
    nodes](#exclusive-text-node){#ref-for-exclusive-text-node④
    link-type="dfn"} and concatenates the
    [data](#concept-cd-data){#ref-for-concept-cd-data⑧ link-type="dfn"}
    of remaining [contiguous exclusive `Text`
    nodes](#contiguous-exclusive-text-nodes){#ref-for-contiguous-exclusive-text-nodes
    link-type="dfn"} into the first of their
    [nodes](#concept-node){#ref-for-concept-node①①③ link-type="dfn"}.

The [`normalize()`]{#dom-node-normalize .dfn .dfn-paneled .idl-code
dfn-for="Node" dfn-type="method" export=""} method steps are to run
these steps for each
[descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①①
link-type="dfn"} [exclusive `Text`
node](#exclusive-text-node){#ref-for-exclusive-text-node⑤
link-type="dfn"} `node`{.variable} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪⓪
link-type="dfn"}:

1.  Let `length`{.variable} be `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length③
    link-type="dfn"}.

2.  If `length`{.variable} is zero, then
    [remove](#concept-node-remove){#ref-for-concept-node-remove⑧
    link-type="dfn"} `node`{.variable} and continue with the next
    [exclusive `Text`
    node](#exclusive-text-node){#ref-for-exclusive-text-node⑥
    link-type="dfn"}, if any.

3.  Let `data`{.variable} be the
    [concatenation](https://infra.spec.whatwg.org/#string-concatenate){#ref-for-string-concatenate①
    link-type="dfn"} of the
    [data](#concept-cd-data){#ref-for-concept-cd-data⑨ link-type="dfn"}
    of `node`{.variable}'s [contiguous exclusive `Text`
    nodes](#contiguous-exclusive-text-nodes){#ref-for-contiguous-exclusive-text-nodes①
    link-type="dfn"} (excluding itself), in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order①⑨
    link-type="dfn"}.

4.  [Replace data](#concept-cd-replace){#ref-for-concept-cd-replace③
    link-type="dfn"} with node `node`{.variable}, offset
    `length`{.variable}, count 0, and data `data`{.variable}.

5.  Let `currentNode`{.variable} be `node`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①⓪
    link-type="dfn"}.

6.  While `currentNode`{.variable} is an [exclusive `Text`
    node](#exclusive-text-node){#ref-for-exclusive-text-node⑦
    link-type="dfn"}:

    1.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range④
        link-type="dfn"} whose [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node②
        link-type="dfn"} is `currentNode`{.variable}, add
        `length`{.variable} to its [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset④
        link-type="dfn"} and set its [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node③
        link-type="dfn"} to `node`{.variable}.

    2.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range⑤
        link-type="dfn"} whose [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node②
        link-type="dfn"} is `currentNode`{.variable}, add
        `length`{.variable} to its [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset④
        link-type="dfn"} and set its [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node③
        link-type="dfn"} to `node`{.variable}.

    3.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range⑥
        link-type="dfn"} whose [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node④
        link-type="dfn"} is `currentNode`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent②⑦
        link-type="dfn"} and [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset⑤
        link-type="dfn"} is `currentNode`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index⑥
        link-type="dfn"}, set its [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node⑤
        link-type="dfn"} to `node`{.variable} and its [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset⑥
        link-type="dfn"} to `length`{.variable}.

    4.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range⑦
        link-type="dfn"} whose [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node④
        link-type="dfn"} is `currentNode`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent②⑧
        link-type="dfn"} and [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset⑤
        link-type="dfn"} is `currentNode`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index⑦
        link-type="dfn"}, set its [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node⑤
        link-type="dfn"} to `node`{.variable} and its [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset⑥
        link-type="dfn"} to `length`{.variable}.

    5.  Add `currentNode`{.variable}'s
        [length](#concept-node-length){#ref-for-concept-node-length④
        link-type="dfn"} to `length`{.variable}.

    6.  Set `currentNode`{.variable} to its [next
        sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①①
        link-type="dfn"}.

7.  [Remove](#concept-node-remove){#ref-for-concept-node-remove⑨
    link-type="dfn"} `node`{.variable}'s [contiguous exclusive `Text`
    nodes](#contiguous-exclusive-text-nodes){#ref-for-contiguous-exclusive-text-nodes②
    link-type="dfn"} (excluding itself), in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order②⓪
    link-type="dfn"}.

------------------------------------------------------------------------

`node`{.variable}` . `[`cloneNode([``subtree`{.variable}` = false])`](#dom-node-clonenode){#ref-for-dom-node-clonenode① .idl-code link-type="method"}
:   Returns a copy of `node`{.variable}. If `subtree`{.variable} is
    true, the copy also includes the `node`{.variable}'s
    [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant①②
    link-type="dfn"}.

`node`{.variable}` . `[`isEqualNode(otherNode)`{.idl}](#dom-node-isequalnode){#ref-for-dom-node-isequalnode① link-type="idl"}
:   Returns whether `node`{.variable} and `otherNode`{.variable} have
    the same properties.

::::::: impl
::: {.algorithm algorithm="cloning steps"}
[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications⑥
link-type="dfn"} may define [cloning steps]{#concept-node-clone-ext .dfn
.dfn-paneled dfn-type="dfn" export=""} for all or some
[nodes](#concept-node){#ref-for-concept-node①①④ link-type="dfn"}. The
algorithm is passed `node`{.variable}, `copy`{.variable}, and
`subtree`{.variable} as indicated in the [clone a
node](#concept-node-clone){#ref-for-concept-node-clone link-type="dfn"}
algorithm.

HTML defines [cloning
steps](#concept-node-clone-ext){#ref-for-concept-node-clone-ext
link-type="dfn"} for several elements, such as
[`input`](https://html.spec.whatwg.org/multipage/input.html#the-input-element){#ref-for-the-input-element①
link-type="element"},
[`script`](https://html.spec.whatwg.org/multipage/scripting.html#script){#ref-for-script③
link-type="element"}, and
[`template`](https://html.spec.whatwg.org/multipage/scripting.html#the-template-element){#ref-for-the-template-element
link-type="element"}. SVG ought to do the same for its
[`script`](https://html.spec.whatwg.org/multipage/scripting.html#script){#ref-for-script④
link-type="element"} elements, but does not.
:::

::: {.algorithm algorithm="clone a node"}
To [clone a node]{#concept-node-clone .dfn .dfn-paneled dfn-type="dfn"
export=""} given a [node](#concept-node){#ref-for-concept-node①①⑤
link-type="dfn"} `node`{.variable} and an optional
[document](#concept-document){#ref-for-concept-document②③
link-type="dfn"} [`document`{.variable}]{#clone-a-node-document .dfn
.dfn-paneled dfn-for="clone a node" dfn-type="dfn" export=""} (default
`node`{.variable}'s [node
document](#concept-node-document){#ref-for-concept-node-document②⑦
link-type="dfn"}), boolean [`subtree`{.variable}]{#clone-a-node-subtree
.dfn .dfn-paneled dfn-for="clone a node" dfn-type="dfn" export=""}
(default false), [node](#concept-node){#ref-for-concept-node①①⑥
link-type="dfn"}-or-null [`parent`{.variable}]{#clone-a-node-parent .dfn
.dfn-paneled dfn-for="clone a node" dfn-type="dfn" export=""} (default
null), and null or a
[`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry②
link-type="idl"} object
[`fallbackRegistry`{.variable}]{#clone-a-node-fallbackregistry .dfn
.dfn-paneled dfn-for="clone a node" dfn-type="dfn" export=""} (default
null):

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert③
    link-type="dfn"}: `node`{.variable} is not a
    [document](#concept-document){#ref-for-concept-document②④
    link-type="dfn"} or `node`{.variable} is `document`{.variable}.

2.  Let `copy`{.variable} be the result of [cloning a single
    node](#clone-a-single-node){#ref-for-clone-a-single-node
    link-type="dfn"} given `node`{.variable}, `document`{.variable}, and
    `fallbackRegistry`{.variable}.

3.  Run any [cloning
    steps](#concept-node-clone-ext){#ref-for-concept-node-clone-ext①
    link-type="dfn"} defined for `node`{.variable} in [other applicable
    specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications⑦
    link-type="dfn"} and pass `node`{.variable}, `copy`{.variable}, and
    `subtree`{.variable} as parameters.

4.  If `parent`{.variable} is non-null, then
    [append](#concept-node-append){#ref-for-concept-node-append②
    link-type="dfn"} `copy`{.variable} to `parent`{.variable}.

5.  If `subtree`{.variable} is true, then for each `child`{.variable} of
    `node`{.variable}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child⑤④
    link-type="dfn"}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order②①
    link-type="dfn"}: [clone a
    node](#concept-node-clone){#ref-for-concept-node-clone①
    link-type="dfn"} given `child`{.variable} with
    [*document*](#clone-a-node-document){#ref-for-clone-a-node-document
    link-type="dfn"} set to `document`{.variable},
    [*subtree*](#clone-a-node-subtree){#ref-for-clone-a-node-subtree
    link-type="dfn"} set to `subtree`{.variable},
    [*parent*](#clone-a-node-parent){#ref-for-clone-a-node-parent
    link-type="dfn"} set to `copy`{.variable}, and
    [*fallbackRegistry*](#clone-a-node-fallbackregistry){#ref-for-clone-a-node-fallbackregistry
    link-type="dfn"} set to `fallbackRegistry`{.variable}.

6.  If `node`{.variable} is an
    [element](#concept-element){#ref-for-concept-element④⑧
    link-type="dfn"}, `node`{.variable} is a [shadow
    host](#element-shadow-host){#ref-for-element-shadow-host②
    link-type="dfn"}, and `node`{.variable}'s [shadow
    root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①
    link-type="dfn"}'s
    [clonable](#shadowroot-clonable){#ref-for-shadowroot-clonable
    link-type="dfn"} is true:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert④
        link-type="dfn"}: `copy`{.variable} is not a [shadow
        host](#element-shadow-host){#ref-for-element-shadow-host③
        link-type="dfn"}.

    2.  Let `shadowRootRegistry`{.variable} be `node`{.variable}'s
        [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root②
        link-type="dfn"}'s [custom element
        registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry⑥
        link-type="dfn"}.

    3.  [Attach a shadow
        root](#concept-attach-a-shadow-root){#ref-for-concept-attach-a-shadow-root
        link-type="dfn"} with `copy`{.variable}, `node`{.variable}'s
        [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root③
        link-type="dfn"}'s
        [mode](#shadowroot-mode){#ref-for-shadowroot-mode④
        link-type="dfn"}, true, `node`{.variable}'s [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root④
        link-type="dfn"}'s
        [serializable](#shadowroot-serializable){#ref-for-shadowroot-serializable
        link-type="dfn"}, `node`{.variable}'s [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root⑤
        link-type="dfn"}'s [delegates
        focus](#shadowroot-delegates-focus){#ref-for-shadowroot-delegates-focus
        link-type="dfn"}, `node`{.variable}'s [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root⑥
        link-type="dfn"}'s [slot
        assignment](#shadowroot-slot-assignment){#ref-for-shadowroot-slot-assignment④
        link-type="dfn"}, and `shadowRootRegistry`{.variable}.

    4.  Set `copy`{.variable}'s [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root⑦
        link-type="dfn"}'s
        [declarative](#shadowroot-declarative){#ref-for-shadowroot-declarative
        link-type="dfn"} to `node`{.variable}'s [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root⑧
        link-type="dfn"}'s
        [declarative](#shadowroot-declarative){#ref-for-shadowroot-declarative①
        link-type="dfn"}.

    5.  For each `child`{.variable} of `node`{.variable}'s [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root⑨
        link-type="dfn"}'s
        [children](#concept-tree-child){#ref-for-concept-tree-child⑤⑤
        link-type="dfn"}, in [tree
        order](#concept-tree-order){#ref-for-concept-tree-order②②
        link-type="dfn"}: [clone a
        node](#concept-node-clone){#ref-for-concept-node-clone②
        link-type="dfn"} given `child`{.variable} with
        [*document*](#clone-a-node-document){#ref-for-clone-a-node-document①
        link-type="dfn"} set to `document`{.variable},
        [*subtree*](#clone-a-node-subtree){#ref-for-clone-a-node-subtree①
        link-type="dfn"} set to `subtree`{.variable}, and
        [*parent*](#clone-a-node-parent){#ref-for-clone-a-node-parent①
        link-type="dfn"} set to `copy`{.variable}'s [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①⓪
        link-type="dfn"}.

        This intentionally does not pass the
        [*fallbackRegistry*](#clone-a-node-fallbackregistry){#ref-for-clone-a-node-fallbackregistry①
        link-type="dfn"} argument.

7.  Return `copy`{.variable}.
:::

::: {.algorithm algorithm="clone a single node"}
To [clone a single node]{#clone-a-single-node .dfn .dfn-paneled
dfn-type="dfn" noexport=""} given a
[node](#concept-node){#ref-for-concept-node①①⑦ link-type="dfn"}
`node`{.variable},
[document](#concept-document){#ref-for-concept-document②⑤
link-type="dfn"} `document`{.variable}, and null or a
[`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry③
link-type="idl"} object `fallbackRegistry`{.variable}:

1.  Let `copy`{.variable} be null.

2.  If `node`{.variable} is an
    [element](#concept-element){#ref-for-concept-element④⑨
    link-type="dfn"}:

    1.  Let `registry`{.variable} be `node`{.variable}'s [custom element
        registry](#element-custom-element-registry){#ref-for-element-custom-element-registry④
        link-type="dfn"}.

    2.  If `registry`{.variable} is null, then set `registry`{.variable}
        to `fallbackRegistry`{.variable}.

    3.  Set `copy`{.variable} to the result of [creating an
        element](#concept-create-element){#ref-for-concept-create-element
        link-type="dfn"}, given `document`{.variable},
        `node`{.variable}'s [local
        name](#concept-element-local-name){#ref-for-concept-element-local-name
        link-type="dfn"}, `node`{.variable}'s
        [namespace](#concept-element-namespace){#ref-for-concept-element-namespace
        link-type="dfn"}, `node`{.variable}'s [namespace
        prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix
        link-type="dfn"}, `node`{.variable}'s [`is`
        value](#concept-element-is-value){#ref-for-concept-element-is-value
        link-type="dfn"}, false, and `registry`{.variable}.

    4.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②④
        link-type="dfn"} `attribute`{.variable} of `node`{.variable}'s
        [attribute
        list](#concept-element-attribute){#ref-for-concept-element-attribute
        link-type="dfn"}:

        1.  Let `copyAttribute`{.variable} be the result of [cloning a
            single
            node](#clone-a-single-node){#ref-for-clone-a-single-node①
            link-type="dfn"} given `attribute`{.variable},
            `document`{.variable}, and null.

        2.  [Append](#concept-element-attributes-append){#ref-for-concept-element-attributes-append
            link-type="dfn"} `copyAttribute`{.variable} to
            `copy`{.variable}.

3.  Otherwise, set `copy`{.variable} to a
    [node](#concept-node){#ref-for-concept-node①①⑧ link-type="dfn"} that
    [implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①①
    link-type="dfn"} the same interfaces as `node`{.variable}, and
    fulfills these additional requirements, switching on the interface
    `node`{.variable}
    [implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①②
    link-type="dfn"}:

    [`Document`{.idl}](#document){#ref-for-document①③ link-type="idl"}

    :   Set `copy`{.variable}'s
        [encoding](#concept-document-encoding){#ref-for-concept-document-encoding
        link-type="dfn"}, [content
        type](#concept-document-content-type){#ref-for-concept-document-content-type
        link-type="dfn"},
        [URL](#concept-document-url){#ref-for-concept-document-url
        link-type="dfn"},
        [origin](#concept-document-origin){#ref-for-concept-document-origin
        link-type="dfn"},
        [type](#concept-document-type){#ref-for-concept-document-type
        link-type="dfn"},
        [mode](#concept-document-mode){#ref-for-concept-document-mode
        link-type="dfn"}, and [custom element
        registry](#document-custom-element-registry){#ref-for-document-custom-element-registry①
        link-type="dfn"} to those of `node`{.variable}.

    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①④ link-type="idl"}

    :   Set `copy`{.variable}'s
        [name](#concept-doctype-name){#ref-for-concept-doctype-name②
        link-type="dfn"}, [public
        ID](#concept-doctype-publicid){#ref-for-concept-doctype-publicid
        link-type="dfn"}, and [system
        ID](#concept-doctype-systemid){#ref-for-concept-doctype-systemid
        link-type="dfn"} to those of `node`{.variable}.

    [`Attr`{.idl}](#attr){#ref-for-attr①③ link-type="idl"}

    :   Set `copy`{.variable}'s
        [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace②
        link-type="dfn"}, [namespace
        prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix
        link-type="dfn"}, [local
        name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②
        link-type="dfn"}, and
        [value](#concept-attribute-value){#ref-for-concept-attribute-value⑦
        link-type="dfn"} to those of `node`{.variable}.

    [`Text`{.idl}](#text){#ref-for-text②② link-type="idl"}\
    [`Comment`{.idl}](#comment){#ref-for-comment⑨ link-type="idl"}

    :   Set `copy`{.variable}'s
        [data](#concept-cd-data){#ref-for-concept-cd-data①⓪
        link-type="dfn"} to that of `node`{.variable}.

    [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction⑨ link-type="idl"}

    :   Set `copy`{.variable}'s
        [target](#concept-pi-target){#ref-for-concept-pi-target②
        link-type="dfn"} and
        [data](#concept-cd-data){#ref-for-concept-cd-data①①
        link-type="dfn"} to those of `node`{.variable}.

    Otherwise

    :   Do nothing.

4.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑤
    link-type="dfn"}: `copy`{.variable} is a
    [node](#concept-node){#ref-for-concept-node①①⑨ link-type="dfn"}.

5.  If `node`{.variable} is a
    [document](#concept-document){#ref-for-concept-document②⑥
    link-type="dfn"}, then set `document`{.variable} to
    `copy`{.variable}.

6.  Set `copy`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document②⑧
    link-type="dfn"} to `document`{.variable}.

7.  Return `copy`{.variable}.
:::

::: {.algorithm algorithm="cloneNode(subtree)" algorithm-for="Node"}
The [`cloneNode(``subtree`{.variable}`)`]{#dom-node-clonenode .dfn
.dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""
lt="cloneNode(subtree)|cloneNode()"} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪①
    link-type="dfn"} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root②⓪
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③⑦
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException③⑨
    link-type="idl"}.

2.  Return the result of [cloning a
    node](#concept-node-clone){#ref-for-concept-node-clone③
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪②
    link-type="dfn"} with
    [*subtree*](#clone-a-node-subtree){#ref-for-clone-a-node-subtree②
    link-type="dfn"} set to `subtree`{.variable}.
:::

A [node](#concept-node){#ref-for-concept-node①②⓪ link-type="dfn"}
`A`{.variable} [equals]{#concept-node-equals .dfn .dfn-paneled
dfn-for="Node" dfn-type="dfn" export=""} a
[node](#concept-node){#ref-for-concept-node①②① link-type="dfn"}
`B`{.variable} if all of the following conditions are true:

- `A`{.variable} and `B`{.variable}
  [implement](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①③
  link-type="dfn"} the same interfaces.

- The following are equal, switching on the interface `A`{.variable}
  [implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①④
  link-type="dfn"}:

  [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①⑤ link-type="idl"}
  :   Its [name](#concept-doctype-name){#ref-for-concept-doctype-name③
      link-type="dfn"}, [public
      ID](#concept-doctype-publicid){#ref-for-concept-doctype-publicid①
      link-type="dfn"}, and [system
      ID](#concept-doctype-systemid){#ref-for-concept-doctype-systemid①
      link-type="dfn"}.

  [`Element`{.idl}](#element){#ref-for-element③⑨ link-type="idl"}
  :   Its
      [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①
      link-type="dfn"}, [namespace
      prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix①
      link-type="dfn"}, [local
      name](#concept-element-local-name){#ref-for-concept-element-local-name①
      link-type="dfn"}, and its [attribute
      list](#concept-element-attribute){#ref-for-concept-element-attribute①
      link-type="dfn"}'s
      [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size③
      link-type="dfn"}.

  [`Attr`{.idl}](#attr){#ref-for-attr①④ link-type="idl"}
  :   Its
      [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace③
      link-type="dfn"}, [local
      name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name③
      link-type="dfn"}, and
      [value](#concept-attribute-value){#ref-for-concept-attribute-value⑧
      link-type="dfn"}.

  [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①⓪ link-type="idl"}
  :   Its [target](#concept-pi-target){#ref-for-concept-pi-target③
      link-type="dfn"} and
      [data](#concept-cd-data){#ref-for-concept-cd-data①②
      link-type="dfn"}.

  [`Text`{.idl}](#text){#ref-for-text②③ link-type="idl"}\
  [`Comment`{.idl}](#comment){#ref-for-comment①⓪ link-type="idl"}
  :   Its [data](#concept-cd-data){#ref-for-concept-cd-data①③
      link-type="dfn"}.

  Otherwise
  :   ---

- If `A`{.variable} is an
  [element](#concept-element){#ref-for-concept-element⑤⓪
  link-type="dfn"}, each
  [attribute](#concept-attribute){#ref-for-concept-attribute⑨
  link-type="dfn"} in its [attribute
  list](#concept-element-attribute){#ref-for-concept-element-attribute②
  link-type="dfn"} has an
  [attribute](#concept-attribute){#ref-for-concept-attribute①⓪
  link-type="dfn"} that
  [equals](#concept-node-equals){#ref-for-concept-node-equals
  link-type="dfn"} an
  [attribute](#concept-attribute){#ref-for-concept-attribute①①
  link-type="dfn"} in `B`{.variable}'s [attribute
  list](#concept-element-attribute){#ref-for-concept-element-attribute③
  link-type="dfn"}.

- `A`{.variable} and `B`{.variable} have the same number of
  [children](#concept-tree-child){#ref-for-concept-tree-child⑤⑥
  link-type="dfn"}.

- Each [child](#concept-tree-child){#ref-for-concept-tree-child⑤⑦
  link-type="dfn"} of `A`{.variable}
  [equals](#concept-node-equals){#ref-for-concept-node-equals①
  link-type="dfn"} the
  [child](#concept-tree-child){#ref-for-concept-tree-child⑤⑧
  link-type="dfn"} of `B`{.variable} at the identical
  [index](#concept-tree-index){#ref-for-concept-tree-index⑧
  link-type="dfn"}.

The [`isEqualNode(``otherNode`{.variable}`)`]{#dom-node-isequalnode .dfn
.dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are to return true if `otherNode`{.variable} is non-null
and [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪③
link-type="dfn"}
[equals](#concept-node-equals){#ref-for-concept-node-equals②
link-type="dfn"} `otherNode`{.variable}; otherwise false.

The [`isSameNode(``otherNode`{.variable}`)`]{#dom-node-issamenode .dfn
.dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are to return true if `otherNode`{.variable} is
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪④
link-type="dfn"}; otherwise false.
:::::::

------------------------------------------------------------------------

`node`{.variable}` . `[`compareDocumentPosition(other)`{.idl}](#dom-node-comparedocumentposition){#ref-for-dom-node-comparedocumentposition① link-type="idl"}
:   Returns a bitmask indicating the position of `other`{.variable}
    relative to `node`{.variable}. These are the bits that can be set:

    [`Node`{.idl}](#node){#ref-for-node④⑨ link-type="idl"}` . `[`DOCUMENT_POSITION_DISCONNECTED`{.idl}](#dom-node-document_position_disconnected){#ref-for-dom-node-document_position_disconnected① link-type="idl"} (1)
    :   Set when `node`{.variable} and `other`{.variable} are not in the
        same [tree](#concept-tree){#ref-for-concept-tree①⑨
        link-type="dfn"}.

    [`Node`{.idl}](#node){#ref-for-node⑤⓪ link-type="idl"}` . `[`DOCUMENT_POSITION_PRECEDING`{.idl}](#dom-node-document_position_preceding){#ref-for-dom-node-document_position_preceding① link-type="idl"} (2)
    :   Set when `other`{.variable} is
        [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding⑧
        link-type="dfn"} `node`{.variable}.

    [`Node`{.idl}](#node){#ref-for-node⑤① link-type="idl"}` . `[`DOCUMENT_POSITION_FOLLOWING`{.idl}](#dom-node-document_position_following){#ref-for-dom-node-document_position_following① link-type="idl"} (4)
    :   Set when `other`{.variable} is
        [following](#concept-tree-following){#ref-for-concept-tree-following①①
        link-type="dfn"} `node`{.variable}.

    [`Node`{.idl}](#node){#ref-for-node⑤② link-type="idl"}` . `[`DOCUMENT_POSITION_CONTAINS`{.idl}](#dom-node-document_position_contains){#ref-for-dom-node-document_position_contains① link-type="idl"} (8)
    :   Set when `other`{.variable} is an
        [ancestor](#concept-tree-ancestor){#ref-for-concept-tree-ancestor③
        link-type="dfn"} of `node`{.variable}.

    [`Node`{.idl}](#node){#ref-for-node⑤③ link-type="idl"}` . `[`DOCUMENT_POSITION_CONTAINED_BY`{.idl}](#dom-node-document_position_contained_by){#ref-for-dom-node-document_position_contained_by① link-type="idl"} (16, 10 in hexadecimal)
    :   Set when `other`{.variable} is a
        [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①③
        link-type="dfn"} of `node`{.variable}.

`node`{.variable}` . `[`contains(other)`{.idl}](#dom-node-contains){#ref-for-dom-node-contains① link-type="idl"}
:   Returns true if `other`{.variable} is an [inclusive
    descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant③
    link-type="dfn"} of `node`{.variable}; otherwise false.

These are the constants
[`compareDocumentPosition()`{.idl}](#dom-node-comparedocumentposition){#ref-for-dom-node-comparedocumentposition②
link-type="idl"} returns as mask:

- [`DOCUMENT_POSITION_DISCONNECTED`]{#dom-node-document_position_disconnected
  .dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
  (1);
- [`DOCUMENT_POSITION_PRECEDING`]{#dom-node-document_position_preceding
  .dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
  (2);
- [`DOCUMENT_POSITION_FOLLOWING`]{#dom-node-document_position_following
  .dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
  (4);
- [`DOCUMENT_POSITION_CONTAINS`]{#dom-node-document_position_contains
  .dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
  (8);
- [`DOCUMENT_POSITION_CONTAINED_BY`]{#dom-node-document_position_contained_by
  .dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
  (16, 10 in hexadecimal);
- [`DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC`]{#dom-node-document_position_implementation_specific
  .dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="const" export=""}
  (32, 20 in hexadecimal).

The
[`compareDocumentPosition(``other`{.variable}`)`]{#dom-node-comparedocumentposition
.dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪⑤
    link-type="dfn"} is `other`{.variable}, then return zero.

2.  Let `node1`{.variable} be `other`{.variable} and `node2`{.variable}
    be [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪⑥
    link-type="dfn"}.

3.  Let `attr1`{.variable} and `attr2`{.variable} be null.

4.  If `node1`{.variable} is an
    [attribute](#concept-attribute){#ref-for-concept-attribute①②
    link-type="dfn"}, then set `attr1`{.variable} to `node1`{.variable}
    and `node1`{.variable} to `attr1`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element
    link-type="dfn"}.

5.  If `node2`{.variable} is an
    [attribute](#concept-attribute){#ref-for-concept-attribute①③
    link-type="dfn"}:

    1.  Set `attr2`{.variable} to `node2`{.variable} and
        `node2`{.variable} to `attr2`{.variable}'s
        [element](#concept-attribute-element){#ref-for-concept-attribute-element①
        link-type="dfn"}.

    2.  If `attr1`{.variable} and `node1`{.variable} are non-null, and
        `node2`{.variable} is `node1`{.variable}:

        1.  [For
            each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②⑤
            link-type="dfn"} `attr`{.variable} of `node2`{.variable}'s
            [attribute
            list](#concept-element-attribute){#ref-for-concept-element-attribute④
            link-type="dfn"}:

            1.  If `attr`{.variable}
                [equals](#concept-node-equals){#ref-for-concept-node-equals③
                link-type="dfn"} `attr1`{.variable}, then return the
                result of adding
                [`DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC`{.idl}](#dom-node-document_position_implementation_specific){#ref-for-dom-node-document_position_implementation_specific①
                link-type="idl"} and
                [`DOCUMENT_POSITION_PRECEDING`{.idl}](#dom-node-document_position_preceding){#ref-for-dom-node-document_position_preceding②
                link-type="idl"}.

            2.  If `attr`{.variable}
                [equals](#concept-node-equals){#ref-for-concept-node-equals④
                link-type="dfn"} `attr2`{.variable}, then return the
                result of adding
                [`DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC`{.idl}](#dom-node-document_position_implementation_specific){#ref-for-dom-node-document_position_implementation_specific②
                link-type="idl"} and
                [`DOCUMENT_POSITION_FOLLOWING`{.idl}](#dom-node-document_position_following){#ref-for-dom-node-document_position_following②
                link-type="idl"}.

6.  If `node1`{.variable} or `node2`{.variable} is null, or
    `node1`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root②⑤
    link-type="dfn"} is not `node2`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root②⑥
    link-type="dfn"}, then return the result of adding
    [`DOCUMENT_POSITION_DISCONNECTED`{.idl}](#dom-node-document_position_disconnected){#ref-for-dom-node-document_position_disconnected②
    link-type="idl"},
    [`DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC`{.idl}](#dom-node-document_position_implementation_specific){#ref-for-dom-node-document_position_implementation_specific③
    link-type="idl"}, and either
    [`DOCUMENT_POSITION_PRECEDING`{.idl}](#dom-node-document_position_preceding){#ref-for-dom-node-document_position_preceding③
    link-type="idl"} or
    [`DOCUMENT_POSITION_FOLLOWING`{.idl}](#dom-node-document_position_following){#ref-for-dom-node-document_position_following③
    link-type="idl"}, with the constraint that this is to be consistent,
    together.

    Whether to return
    [`DOCUMENT_POSITION_PRECEDING`{.idl}](#dom-node-document_position_preceding){#ref-for-dom-node-document_position_preceding④
    link-type="idl"} or
    [`DOCUMENT_POSITION_FOLLOWING`{.idl}](#dom-node-document_position_following){#ref-for-dom-node-document_position_following④
    link-type="idl"} is typically implemented via pointer comparison. In
    JavaScript implementations a cached `Math.random()`{.lang-javascript
    .highlight} value can be used.

7.  If `node1`{.variable} is an
    [ancestor](#concept-tree-ancestor){#ref-for-concept-tree-ancestor④
    link-type="dfn"} of `node2`{.variable} and `attr1`{.variable} is
    null, or `node1`{.variable} is `node2`{.variable} and
    `attr2`{.variable} is non-null, then return the result of adding
    [`DOCUMENT_POSITION_CONTAINS`{.idl}](#dom-node-document_position_contains){#ref-for-dom-node-document_position_contains②
    link-type="idl"} to
    [`DOCUMENT_POSITION_PRECEDING`{.idl}](#dom-node-document_position_preceding){#ref-for-dom-node-document_position_preceding⑤
    link-type="idl"}.

8.  If `node1`{.variable} is a
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①④
    link-type="dfn"} of `node2`{.variable} and `attr2`{.variable} is
    null, or `node1`{.variable} is `node2`{.variable} and
    `attr1`{.variable} is non-null, then return the result of adding
    [`DOCUMENT_POSITION_CONTAINED_BY`{.idl}](#dom-node-document_position_contained_by){#ref-for-dom-node-document_position_contained_by②
    link-type="idl"} to
    [`DOCUMENT_POSITION_FOLLOWING`{.idl}](#dom-node-document_position_following){#ref-for-dom-node-document_position_following⑤
    link-type="idl"}.

9.  If `node1`{.variable} is
    [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding⑨
    link-type="dfn"} `node2`{.variable}, then return
    [`DOCUMENT_POSITION_PRECEDING`{.idl}](#dom-node-document_position_preceding){#ref-for-dom-node-document_position_preceding⑥
    link-type="idl"}.

    Due to the way
    [attributes](#concept-attribute){#ref-for-concept-attribute①④
    link-type="dfn"} are handled in this algorithm this results in a
    [node](#concept-node){#ref-for-concept-node①②② link-type="dfn"}'s
    [attributes](#concept-attribute){#ref-for-concept-attribute①⑤
    link-type="dfn"} counting as
    [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding①⓪
    link-type="dfn"} that [node](#concept-node){#ref-for-concept-node①②③
    link-type="dfn"}'s
    [children](#concept-tree-child){#ref-for-concept-tree-child⑤⑨
    link-type="dfn"}, despite
    [attributes](#concept-attribute){#ref-for-concept-attribute①⑥
    link-type="dfn"} not
    [participating](#concept-tree-participate){#ref-for-concept-tree-participate⑨
    link-type="dfn"} in the same
    [tree](#concept-tree){#ref-for-concept-tree②⓪ link-type="dfn"}.

10. Return
    [`DOCUMENT_POSITION_FOLLOWING`{.idl}](#dom-node-document_position_following){#ref-for-dom-node-document_position_following⑥
    link-type="idl"}.

The [`contains(``other`{.variable}`)`]{#dom-node-contains .dfn
.dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are to return true if `other`{.variable} is an [inclusive
descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant④
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪⑦
link-type="dfn"}; otherwise false (including when `other`{.variable} is
null).

------------------------------------------------------------------------

To [locate a namespace prefix]{#locate-a-namespace-prefix .dfn
.dfn-paneled dfn-type="dfn" export=""
lt="locate a namespace prefix|locating a namespace prefix"} for an
`element`{.variable} using `namespace`{.variable}, run these steps:

1.  If `element`{.variable}'s
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace②
    link-type="dfn"} is `namespace`{.variable} and its [namespace
    prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix②
    link-type="dfn"} is non-null, then return its [namespace
    prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix③
    link-type="dfn"}.

2.  If `element`{.variable}
    [has](#concept-element-attribute-has){#ref-for-concept-element-attribute-has②
    link-type="dfn"} an
    [attribute](#concept-attribute){#ref-for-concept-attribute①⑦
    link-type="dfn"} whose [namespace
    prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix①
    link-type="dfn"} is \"`xmlns`\" and
    [value](#concept-attribute-value){#ref-for-concept-attribute-value⑨
    link-type="dfn"} is `namespace`{.variable}, then return
    `element`{.variable}'s first such
    [attribute](#concept-attribute){#ref-for-concept-attribute①⑧
    link-type="dfn"}'s [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name④
    link-type="dfn"}.

3.  If `element`{.variable}'s [parent
    element](#parent-element){#ref-for-parent-element② link-type="dfn"}
    is not null, then return the result of running [locate a namespace
    prefix](#locate-a-namespace-prefix){#ref-for-locate-a-namespace-prefix
    link-type="dfn"} on that
    [element](#concept-element){#ref-for-concept-element⑤①
    link-type="dfn"} using `namespace`{.variable}.

4.  Return null.

To [locate a namespace]{#locate-a-namespace .dfn .dfn-paneled
dfn-type="dfn" export=""} for a `node`{.variable} using
`prefix`{.variable}, switch on the interface `node`{.variable}
[implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①⑤
link-type="dfn"}:

[`Element`{.idl}](#element){#ref-for-element④⓪ link-type="idl"}

:   1.  If `prefix`{.variable} is \"`xml`\", then return the [XML
        namespace](https://infra.spec.whatwg.org/#xml-namespace){#ref-for-xml-namespace①
        link-type="dfn"}.

    2.  If `prefix`{.variable} is \"`xmlns`\", then return the [XMLNS
        namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#ref-for-xmlns-namespace②
        link-type="dfn"}.

    3.  If its
        [namespace](#concept-element-namespace){#ref-for-concept-element-namespace③
        link-type="dfn"} is non-null and its [namespace
        prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix④
        link-type="dfn"} is `prefix`{.variable}, then return
        [namespace](#concept-element-namespace){#ref-for-concept-element-namespace④
        link-type="dfn"}.

    4.  If it
        [has](#concept-element-attribute-has){#ref-for-concept-element-attribute-has③
        link-type="dfn"} an
        [attribute](#concept-attribute){#ref-for-concept-attribute①⑨
        link-type="dfn"} whose
        [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace④
        link-type="dfn"} is the [XMLNS
        namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#ref-for-xmlns-namespace③
        link-type="dfn"}, [namespace
        prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix②
        link-type="dfn"} is \"`xmlns`\", and [local
        name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name⑤
        link-type="dfn"} is `prefix`{.variable}, or if
        `prefix`{.variable} is null and it
        [has](#concept-element-attribute-has){#ref-for-concept-element-attribute-has④
        link-type="dfn"} an
        [attribute](#concept-attribute){#ref-for-concept-attribute②⓪
        link-type="dfn"} whose
        [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace⑤
        link-type="dfn"} is the [XMLNS
        namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#ref-for-xmlns-namespace④
        link-type="dfn"}, [namespace
        prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix③
        link-type="dfn"} is null, and [local
        name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name⑥
        link-type="dfn"} is \"`xmlns`\", then return its
        [value](#concept-attribute-value){#ref-for-concept-attribute-value①⓪
        link-type="dfn"} if it is not the empty string, and null
        otherwise.

    5.  If its [parent
        element](#parent-element){#ref-for-parent-element③
        link-type="dfn"} is null, then return null.

    6.  Return the result of running [locate a
        namespace](#locate-a-namespace){#ref-for-locate-a-namespace
        link-type="dfn"} on its [parent
        element](#parent-element){#ref-for-parent-element④
        link-type="dfn"} using `prefix`{.variable}.

[`Document`{.idl}](#document){#ref-for-document①④ link-type="idl"}

:   1.  If its [document
        element](#document-element){#ref-for-document-element①
        link-type="dfn"} is null, then return null.

    2.  Return the result of running [locate a
        namespace](#locate-a-namespace){#ref-for-locate-a-namespace①
        link-type="dfn"} on its [document
        element](#document-element){#ref-for-document-element②
        link-type="dfn"} using `prefix`{.variable}.

[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①⑥ link-type="idl"}\
[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②③ link-type="idl"}

:   Return null.

[`Attr`{.idl}](#attr){#ref-for-attr①⑤ link-type="idl"}

:   1.  If its
        [element](#concept-attribute-element){#ref-for-concept-attribute-element②
        link-type="dfn"} is null, then return null.

    2.  Return the result of running [locate a
        namespace](#locate-a-namespace){#ref-for-locate-a-namespace②
        link-type="dfn"} on its
        [element](#concept-attribute-element){#ref-for-concept-attribute-element③
        link-type="dfn"} using `prefix`{.variable}.

Otherwise

:   1.  If its [parent
        element](#parent-element){#ref-for-parent-element⑤
        link-type="dfn"} is null, then return null.

    2.  Return the result of running [locate a
        namespace](#locate-a-namespace){#ref-for-locate-a-namespace③
        link-type="dfn"} on its [parent
        element](#parent-element){#ref-for-parent-element⑥
        link-type="dfn"} using `prefix`{.variable}.

The [`lookupPrefix(``namespace`{.variable}`)`]{#dom-node-lookupprefix
.dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are:

1.  If `namespace`{.variable} is null or the empty string, then return
    null.

2.  Switch on the interface
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪⑧
    link-type="dfn"}
    [implements](https://webidl.spec.whatwg.org/#implements){#ref-for-implements①⑥
    link-type="dfn"}:

    [`Element`{.idl}](#element){#ref-for-element④① link-type="idl"}

    :   Return the result of [locating a namespace
        prefix](#locate-a-namespace-prefix){#ref-for-locate-a-namespace-prefix①
        link-type="dfn"} for
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⓪⑨
        link-type="dfn"} using `namespace`{.variable}.

    [`Document`{.idl}](#document){#ref-for-document①⑤ link-type="idl"}

    :   1.  If
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①⓪
            link-type="dfn"}'s [document
            element](#document-element){#ref-for-document-element③
            link-type="dfn"} is null, then return null.

        2.  Return the result of [locating a namespace
            prefix](#locate-a-namespace-prefix){#ref-for-locate-a-namespace-prefix②
            link-type="dfn"} for
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①①
            link-type="dfn"}'s [document
            element](#document-element){#ref-for-document-element④
            link-type="dfn"} using `namespace`{.variable}.

    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype①⑦ link-type="idl"}\
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②④ link-type="idl"}

    :   Return null.

    [`Attr`{.idl}](#attr){#ref-for-attr①⑥ link-type="idl"}

    :   1.  If
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①②
            link-type="dfn"}'s
            [element](#concept-attribute-element){#ref-for-concept-attribute-element④
            link-type="dfn"} is null, then return null.

        2.  Return the result of [locating a namespace
            prefix](#locate-a-namespace-prefix){#ref-for-locate-a-namespace-prefix③
            link-type="dfn"} for
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①③
            link-type="dfn"}'s
            [element](#concept-attribute-element){#ref-for-concept-attribute-element⑤
            link-type="dfn"} using `namespace`{.variable}.

    Otherwise

    :   1.  If
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①④
            link-type="dfn"}'s [parent
            element](#parent-element){#ref-for-parent-element⑦
            link-type="dfn"} is null, then return null.

        2.  Return the result of [locating a namespace
            prefix](#locate-a-namespace-prefix){#ref-for-locate-a-namespace-prefix④
            link-type="dfn"} for
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①⑤
            link-type="dfn"}'s [parent
            element](#parent-element){#ref-for-parent-element⑧
            link-type="dfn"} using `namespace`{.variable}.

The
[`lookupNamespaceURI(``prefix`{.variable}`)`]{#dom-node-lookupnamespaceuri
.dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are:

1.  If `prefix`{.variable} is the empty string, then set it to null.

2.  Return the result of running [locate a
    namespace](#locate-a-namespace){#ref-for-locate-a-namespace④
    link-type="dfn"} for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①⑥
    link-type="dfn"} using `prefix`{.variable}.

The
[`isDefaultNamespace(``namespace`{.variable}`)`]{#dom-node-isdefaultnamespace
.dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are:

1.  If `namespace`{.variable} is the empty string, then set it to null.

2.  Let `defaultNamespace`{.variable} be the result of running [locate a
    namespace](#locate-a-namespace){#ref-for-locate-a-namespace⑤
    link-type="dfn"} for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①⑦
    link-type="dfn"} using null.

3.  Return true if `defaultNamespace`{.variable} is the same as
    `namespace`{.variable}; otherwise false.

------------------------------------------------------------------------

The
[`insertBefore(``node`{.variable}`, ``child`{.variable}`)`]{#dom-node-insertbefore
.dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are to return the result of
[pre-inserting](#concept-node-pre-insert){#ref-for-concept-node-pre-insert⑥
link-type="dfn"} `node`{.variable} into
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①⑧
link-type="dfn"} before `child`{.variable}.

The [`appendChild(``node`{.variable}`)`]{#dom-node-appendchild .dfn
.dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are to return the result of
[appending](#concept-node-append){#ref-for-concept-node-append③
link-type="dfn"} `node`{.variable} to
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①①⑨
link-type="dfn"}.

The
[`replaceChild(``node`{.variable}`, ``child`{.variable}`)`]{#dom-node-replacechild
.dfn .dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are to return the result of
[replacing](#concept-node-replace){#ref-for-concept-node-replace①
link-type="dfn"} `child`{.variable} with `node`{.variable} within
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②⓪
link-type="dfn"}.

The [`removeChild(``child`{.variable}`)`]{#dom-node-removechild .dfn
.dfn-paneled .idl-code dfn-for="Node" dfn-type="method" export=""}
method steps are to return the result of
[pre-removing](#concept-node-pre-remove){#ref-for-concept-node-pre-remove
link-type="dfn"} `child`{.variable} from
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②①
link-type="dfn"}.

------------------------------------------------------------------------

The [list of elements with qualified name
`qualifiedName`{.variable}]{#concept-getelementsbytagname .dfn
.dfn-paneled dfn-type="dfn" export=""} for a
[node](#concept-node){#ref-for-concept-node①②④ link-type="dfn"}
`root`{.variable} is the
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection⑥
link-type="idl"} returned by the following algorithm:

1.  If `qualifiedName`{.variable} is U+002A (\*), then return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection⑦
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    only
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①⑤
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤②
    link-type="dfn"}.

2.  Otherwise, if `root`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document②⑨
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document link-type="dfn"},
    return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection⑧
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    the following
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①⑥
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤③
    link-type="dfn"}:

    - Whose
      [namespace](#concept-element-namespace){#ref-for-concept-element-namespace⑤
      link-type="dfn"} is the [HTML
      namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②
      link-type="dfn"} and whose [qualified
      name](#concept-element-qualified-name){#ref-for-concept-element-qualified-name
      link-type="dfn"} is `qualifiedName`{.variable}, in [ASCII
      lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase
      link-type="dfn"}.

    - Whose
      [namespace](#concept-element-namespace){#ref-for-concept-element-namespace⑥
      link-type="dfn"} is *not* the [HTML
      namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace③
      link-type="dfn"} and whose [qualified
      name](#concept-element-qualified-name){#ref-for-concept-element-qualified-name①
      link-type="dfn"} is `qualifiedName`{.variable}.

3.  Otherwise, return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection⑨
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①⑦
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤④
    link-type="dfn"} whose [qualified
    name](#concept-element-qualified-name){#ref-for-concept-element-qualified-name②
    link-type="dfn"} is `qualifiedName`{.variable}.

When invoked with the same argument, and as long as `root`{.variable}'s
[node document](#concept-node-document){#ref-for-concept-node-document③⓪
link-type="dfn"}'s
[type](#concept-document-type){#ref-for-concept-document-type①
link-type="dfn"} has not changed, the same
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①⓪
link-type="idl"} object may be returned as returned by an earlier call.

The [list of elements with namespace `namespace`{.variable} and local
name `localName`{.variable}]{#concept-getelementsbytagnamens .dfn
.dfn-paneled dfn-type="dfn" export=""} for a
[node](#concept-node){#ref-for-concept-node①②⑤ link-type="dfn"}
`root`{.variable} is the
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①①
link-type="idl"} returned by the following algorithm:

1.  If `namespace`{.variable} is the empty string, then set it to null.

2.  If both `namespace`{.variable} and `localName`{.variable} are U+002A
    (\*), then return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①②
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①⑧
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤⑤
    link-type="dfn"}.

3.  If `namespace`{.variable} is U+002A (\*), then return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①③
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant①⑨
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤⑥
    link-type="dfn"} whose [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name②
    link-type="dfn"} is `localName`{.variable}.

4.  If `localName`{.variable} is U+002A (\*), then return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①④
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②⓪
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤⑦
    link-type="dfn"} whose
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace⑦
    link-type="dfn"} is `namespace`{.variable}.

5.  Return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①⑤
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②①
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤⑧
    link-type="dfn"} whose
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace⑧
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name③
    link-type="dfn"} is `localName`{.variable}.

When invoked with the same arguments, the same
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①⑥
link-type="idl"} object may be returned as returned by an earlier call.

The [list of elements with class names
`classNames`{.variable}]{#concept-getelementsbyclassname .dfn
.dfn-paneled dfn-type="dfn" export=""} for a
[node](#concept-node){#ref-for-concept-node①②⑥ link-type="dfn"}
`root`{.variable} is the
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①⑦
link-type="idl"} returned by the following algorithm:

1.  Let `classes`{.variable} be the result of running the [ordered set
    parser](#concept-ordered-set-parser){#ref-for-concept-ordered-set-parser
    link-type="dfn"} on `classNames`{.variable}.

2.  If `classes`{.variable} is the empty set, return an empty
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①⑧
    link-type="idl"}.

3.  Return an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection①⑨
    link-type="idl"} rooted at `root`{.variable}, whose filter matches
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②②
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑤⑨
    link-type="dfn"} that have all their
    [classes](#concept-class){#ref-for-concept-class link-type="dfn"} in
    `classes`{.variable}.

    The comparisons for the
    [classes](#concept-class){#ref-for-concept-class① link-type="dfn"}
    must be done in an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive
    link-type="dfn"} manner if `root`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document③①
    link-type="dfn"}'s
    [mode](#concept-document-mode){#ref-for-concept-document-mode①
    link-type="dfn"} is \"`quirks`\"; otherwise in an [identical
    to](https://infra.spec.whatwg.org/#string-is){#ref-for-string-is
    link-type="dfn"} manner.

When invoked with the same argument, the same
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection②⓪
link-type="idl"} object may be returned as returned by an earlier call.

### [4.5. ]{.secno}[Interface [`Document`{.idl}](#document){#ref-for-document①⑥ link-type="idl"}]{.content}[](#interface-document){.self-link} {#interface-document .heading .settled level="4.5"}

``` {.idl .highlight .def dfn-force="Document"}
[Exposed=Window]
interface Document : Node {
  constructor();

  [SameObject] readonly attribute DOMImplementation implementation;
  readonly attribute USVString URL;
  readonly attribute USVString documentURI;
  readonly attribute DOMString compatMode;
  readonly attribute DOMString characterSet;
  readonly attribute DOMString charset; // legacy alias of .characterSet
  readonly attribute DOMString inputEncoding; // legacy alias of .characterSet
  readonly attribute DOMString contentType;

  readonly attribute DocumentType? doctype;
  readonly attribute Element? documentElement;
  HTMLCollection getElementsByTagName(DOMString qualifiedName);
  HTMLCollection getElementsByTagNameNS(DOMString? namespace, DOMString localName);
  HTMLCollection getElementsByClassName(DOMString classNames);

  [CEReactions, NewObject] Element createElement(DOMString localName, optional (DOMString or ElementCreationOptions) options = {});
  [CEReactions, NewObject] Element createElementNS(DOMString? namespace, DOMString qualifiedName, optional (DOMString or ElementCreationOptions) options = {});
  [NewObject] DocumentFragment createDocumentFragment();
  [NewObject] Text createTextNode(DOMString data);
  [NewObject] CDATASection createCDATASection(DOMString data);
  [NewObject] Comment createComment(DOMString data);
  [NewObject] ProcessingInstruction createProcessingInstruction(DOMString target, DOMString data);

  [CEReactions, NewObject] Node importNode(Node node, optional (boolean or ImportNodeOptions) options = false);
  [CEReactions] Node adoptNode(Node node);

  [NewObject] Attr createAttribute(DOMString localName);
  [NewObject] Attr createAttributeNS(DOMString? namespace, DOMString qualifiedName);

  [NewObject] Event createEvent(DOMString interface); // legacy

  [NewObject] Range createRange();

  // NodeFilter.SHOW_ALL = 0xFFFFFFFF
  [NewObject] NodeIterator createNodeIterator(Node root, optional unsigned long whatToShow = 0xFFFFFFFF, optional NodeFilter? filter = null);
  [NewObject] TreeWalker createTreeWalker(Node root, optional unsigned long whatToShow = 0xFFFFFFFF, optional NodeFilter? filter = null);
};

[Exposed=Window]
interface XMLDocument : Document {};

dictionary ElementCreationOptions {
  CustomElementRegistry customElementRegistry;
  DOMString is;
};

dictionary ImportNodeOptions {
  CustomElementRegistry customElementRegistry;
  boolean selfOnly = false;
};
```

[`Document`{.idl}](#document){#ref-for-document①⑧ link-type="idl"}
[nodes](#concept-node){#ref-for-concept-node①②⑦ link-type="dfn"} are
simply known as [documents]{#concept-document .dfn .dfn-paneled
dfn-type="dfn" export="" lt="document"}.

A [document](#concept-document){#ref-for-concept-document②⑦
link-type="dfn"}'s [node
document](#concept-node-document){#ref-for-concept-node-document③②
link-type="dfn"} is itself.

Each [document](#concept-document){#ref-for-concept-document②⑧
link-type="dfn"} has an associated [encoding]{#concept-document-encoding
.dfn .dfn-paneled dfn-for="Document" dfn-type="dfn" export=""} (an
[encoding](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding
link-type="dfn"}), [content type]{#concept-document-content-type .dfn
.dfn-paneled dfn-for="Document" dfn-type="dfn" export=""} (a string),
[URL]{#concept-document-url .dfn .dfn-paneled dfn-for="Document"
dfn-type="dfn" export=""} (a
[URL](https://url.spec.whatwg.org/#concept-url){#ref-for-concept-url
link-type="dfn"}), [origin]{#concept-document-origin .dfn .dfn-paneled
dfn-for="Document" dfn-type="dfn" export=""} (an
[origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin){#ref-for-concept-origin
link-type="dfn"}), [type]{#concept-document-type .dfn .dfn-paneled
dfn-for="Document" dfn-type="dfn" export=""} (\"`xml`\" or \"`html`\"),
[mode]{#concept-document-mode .dfn .dfn-paneled dfn-for="Document"
dfn-type="dfn" export=""} (\"`no-quirks`\", \"`quirks`\", or
\"`limited-quirks`\"), [allow declarative shadow
roots]{#document-allow-declarative-shadow-roots .dfn .dfn-paneled
dfn-for="Document" dfn-type="dfn" export=""} (a boolean), and [custom
element registry]{#document-custom-element-registry .dfn .dfn-paneled
dfn-for="Document" dfn-type="dfn" export=""} (null or a
[`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry⑥
link-type="idl"} object).
[\[ENCODING\]](#biblio-encoding "Encoding Standard"){link-type="biblio"}
[\[URL\]](#biblio-url "URL Standard"){link-type="biblio"}
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

Unless stated otherwise, a
[document](#concept-document){#ref-for-concept-document②⑨
link-type="dfn"}'s
[encoding](#concept-document-encoding){#ref-for-concept-document-encoding①
link-type="dfn"} is the
[utf-8](https://encoding.spec.whatwg.org/#utf-8){#ref-for-utf-8
link-type="dfn"}
[encoding](https://encoding.spec.whatwg.org/#encoding){#ref-for-encoding①
link-type="dfn"}, [content
type](#concept-document-content-type){#ref-for-concept-document-content-type①
link-type="dfn"} is \"`application/xml`\",
[URL](#concept-document-url){#ref-for-concept-document-url①
link-type="dfn"} is \"`about:blank`\",
[origin](#concept-document-origin){#ref-for-concept-document-origin①
link-type="dfn"} is an [opaque
origin](https://html.spec.whatwg.org/multipage/browsers.html#concept-origin-opaque){#ref-for-concept-origin-opaque
link-type="dfn"},
[type](#concept-document-type){#ref-for-concept-document-type②
link-type="dfn"} is \"`xml`\",
[mode](#concept-document-mode){#ref-for-concept-document-mode②
link-type="dfn"} is \"`no-quirks`\", [allow declarative shadow
roots](#document-allow-declarative-shadow-roots){#ref-for-document-allow-declarative-shadow-roots
link-type="dfn"} is false, and [custom element
registry](#document-custom-element-registry){#ref-for-document-custom-element-registry②
link-type="dfn"} is null.

A [document](#concept-document){#ref-for-concept-document③⓪
link-type="dfn"} is said to be an [XML document]{#xml-document .dfn
.dfn-paneled dfn-type="dfn" export=""} if its
[type](#concept-document-type){#ref-for-concept-document-type③
link-type="dfn"} is \"`xml`\"; otherwise an [HTML
document]{#html-document .dfn .dfn-paneled dfn-type="dfn" export=""}.
Whether a [document](#concept-document){#ref-for-concept-document③①
link-type="dfn"} is an [HTML
document](#html-document){#ref-for-html-document① link-type="dfn"} or an
[XML document](#xml-document){#ref-for-xml-document link-type="dfn"}
affects the behavior of certain APIs.

A [document](#concept-document){#ref-for-concept-document③②
link-type="dfn"} is said to be in [no-quirks
mode]{#concept-document-no-quirks .dfn .dfn-paneled dfn-type="dfn"
export=""} if its
[mode](#concept-document-mode){#ref-for-concept-document-mode③
link-type="dfn"} is \"`no-quirks`\", [quirks
mode]{#concept-document-quirks .dfn .dfn-paneled dfn-type="dfn"
export=""} if its
[mode](#concept-document-mode){#ref-for-concept-document-mode④
link-type="dfn"} is \"`quirks`\", and [limited-quirks
mode]{#concept-document-limited-quirks .dfn .dfn-paneled dfn-type="dfn"
export=""} if its
[mode](#concept-document-mode){#ref-for-concept-document-mode⑤
link-type="dfn"} is \"`limited-quirks`\".

::: {.note role="note"}
The [mode](#concept-document-mode){#ref-for-concept-document-mode⑥
link-type="dfn"} is only ever changed from the default for
[documents](#concept-document){#ref-for-concept-document③③
link-type="dfn"} created by the [HTML
parser](https://html.spec.whatwg.org/multipage/parsing.html#html-parser){#ref-for-html-parser①
link-type="dfn"} based on the presence, absence, or value of the DOCTYPE
string, and by a new [browsing
context](https://html.spec.whatwg.org/multipage/document-sequences.html#browsing-context){#ref-for-browsing-context③
link-type="dfn"} (initial \"`about:blank`\").
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

[No-quirks
mode](#concept-document-no-quirks){#ref-for-concept-document-no-quirks
link-type="dfn"} was originally known as \"standards mode\" and
[limited-quirks
mode](#concept-document-limited-quirks){#ref-for-concept-document-limited-quirks
link-type="dfn"} was once known as \"almost standards mode\". They have
been renamed because their details are now defined by standards. (And
because Ian Hickson vetoed their original names on the basis that they
are nonsensical.)
:::

A [document](#concept-document){#ref-for-concept-document③④
link-type="dfn"}'s [get the
parent](#get-the-parent){#ref-for-get-the-parent⑥ link-type="dfn"}
algorithm, given an `event`{.variable}, returns null if
`event`{.variable}'s
[`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⑦
link-type="idl"} attribute value is \"`load`\" or
[document](#concept-document){#ref-for-concept-document③⑤
link-type="dfn"} does not have a [browsing
context](https://html.spec.whatwg.org/multipage/document-sequences.html#concept-document-bc){#ref-for-concept-document-bc
link-type="dfn"}; otherwise the
[document](#concept-document){#ref-for-concept-document③⑥
link-type="dfn"}'s [relevant global
object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global③
link-type="dfn"}.

------------------------------------------------------------------------

`document`{.variable}` = new `[`Document()`{.idl}](#dom-document-document){#ref-for-dom-document-document① link-type="idl"}
:   Returns a new
    [document](#concept-document){#ref-for-concept-document③⑦
    link-type="dfn"}.

`document`{.variable}` . `[`implementation`{.idl}](#dom-document-implementation){#ref-for-dom-document-implementation① link-type="idl"}
:   Returns `document`{.variable}'s
    [`DOMImplementation`{.idl}](#domimplementation){#ref-for-domimplementation①
    link-type="idl"} object.

`document`{.variable}` . `[`URL`{.idl}](#dom-document-url){#ref-for-dom-document-url① link-type="idl"}\
`document`{.variable}` . `[`documentURI`{.idl}](#dom-document-documenturi){#ref-for-dom-document-documenturi① link-type="idl"}
:   Returns `document`{.variable}'s
    [URL](#concept-document-url){#ref-for-concept-document-url②
    link-type="dfn"}.

`document`{.variable}` . `[`compatMode`{.idl}](#dom-document-compatmode){#ref-for-dom-document-compatmode① link-type="idl"}
:   Returns the string \"`BackCompat`\" if `document`{.variable}'s
    [mode](#concept-document-mode){#ref-for-concept-document-mode⑦
    link-type="dfn"} is \"`quirks`\"; otherwise \"`CSS1Compat`\".

`document`{.variable}` . `[`characterSet`{.idl}](#dom-document-characterset){#ref-for-dom-document-characterset① link-type="idl"}
:   Returns `document`{.variable}'s
    [encoding](#concept-document-encoding){#ref-for-concept-document-encoding②
    link-type="dfn"}.

`document`{.variable}` . `[`contentType`{.idl}](#dom-document-contenttype){#ref-for-dom-document-contenttype① link-type="idl"}
:   Returns `document`{.variable}'s [content
    type](#concept-document-content-type){#ref-for-concept-document-content-type②
    link-type="dfn"}.

The [`new Document()`]{#dom-document-document .dfn .dfn-paneled
.idl-code dfn-for="Document" dfn-type="constructor" export=""
lt="Document()|constructor()"} constructor steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②②
link-type="dfn"}'s
[origin](#concept-document-origin){#ref-for-concept-document-origin②
link-type="dfn"} to the
[origin](#concept-document-origin){#ref-for-concept-document-origin③
link-type="dfn"} of [current global
object](https://html.spec.whatwg.org/multipage/webappapis.html#current-global-object){#ref-for-current-global-object
link-type="dfn"}'s [associated
`Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window①
link-type="dfn"}.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

Unlike
[`createDocument()`{.idl}](#dom-domimplementation-createdocument){#ref-for-dom-domimplementation-createdocument
link-type="idl"}, this constructor does not return an
[`XMLDocument`{.idl}](#xmldocument){#ref-for-xmldocument
link-type="idl"} object, but a
[document](#concept-document){#ref-for-concept-document③⑧
link-type="dfn"} ([`Document`{.idl}](#document){#ref-for-document①⑨
link-type="idl"} object).

The [`implementation`]{#dom-document-implementation .dfn .dfn-paneled
.idl-code dfn-for="Document" dfn-type="attribute" export=""} getter
steps are to return the
[`DOMImplementation`{.idl}](#domimplementation){#ref-for-domimplementation②
link-type="idl"} object that is associated with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②③
link-type="dfn"}.

The [`URL`]{#dom-document-url .dfn .dfn-paneled .idl-code
dfn-for="Document" dfn-type="attribute" export=""} and
[`documentURI`]{#dom-document-documenturi .dfn .dfn-paneled .idl-code
dfn-for="Document" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②④
link-type="dfn"}'s
[URL](#concept-document-url){#ref-for-concept-document-url③
link-type="dfn"},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#ref-for-concept-url-serializer①
link-type="dfn"}.

The [`compatMode`]{#dom-document-compatmode .dfn .dfn-paneled .idl-code
dfn-for="Document" dfn-type="attribute" export=""} getter steps are to
return \"`BackCompat`\" if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②⑤
link-type="dfn"}'s
[mode](#concept-document-mode){#ref-for-concept-document-mode⑧
link-type="dfn"} is \"`quirks`\"; otherwise \"`CSS1Compat`\".

The [`characterSet`]{#dom-document-characterset .dfn .dfn-paneled
.idl-code dfn-for="Document" dfn-type="attribute" export=""},
[`charset`]{#dom-document-charset .dfn .dfn-paneled .idl-code
dfn-for="Document" dfn-type="attribute" export=""}, and
[`inputEncoding`]{#dom-document-inputencoding .dfn .dfn-paneled
.idl-code dfn-for="Document" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②⑥
link-type="dfn"}'s
[encoding](#concept-document-encoding){#ref-for-concept-document-encoding③
link-type="dfn"}'s
[name](https://encoding.spec.whatwg.org/#name){#ref-for-name
link-type="dfn"}.

The [`contentType`]{#dom-document-contenttype .dfn .dfn-paneled
.idl-code dfn-for="Document" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②⑦
link-type="dfn"}'s [content
type](#concept-document-content-type){#ref-for-concept-document-content-type③
link-type="dfn"}.

------------------------------------------------------------------------

`document`{.variable} . [`doctype`{.idl}](#dom-document-doctype){#ref-for-dom-document-doctype① link-type="idl"}
:   Returns the [doctype](#concept-doctype){#ref-for-concept-doctype①④
    link-type="dfn"} or null if there is none.

`document`{.variable} . [`documentElement`{.idl}](#dom-document-documentelement){#ref-for-dom-document-documentelement① link-type="idl"}
:   Returns the [document
    element](#document-element){#ref-for-document-element⑤
    link-type="dfn"}.

`collection`{.variable}` = ``document`{.variable}` . `[`getElementsByTagName`](#dom-document-getelementsbytagname){#ref-for-dom-document-getelementsbytagname① .idl-code link-type="method"}`(``qualifiedName`{.variable}`)`

:   If `qualifiedName`{.variable} is \"`*`\" returns an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection②④
    link-type="idl"} of all
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②③
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑥⓪
    link-type="dfn"}.

    Otherwise, returns an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection②⑤
    link-type="idl"} of all
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②④
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑥①
    link-type="dfn"} whose [qualified
    name](#concept-element-qualified-name){#ref-for-concept-element-qualified-name③
    link-type="dfn"} is `qualifiedName`{.variable}. (Matches
    case-insensitively against
    [elements](#concept-element){#ref-for-concept-element⑥②
    link-type="dfn"} in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace④
    link-type="dfn"} within an [HTML
    document](#html-document){#ref-for-html-document② link-type="dfn"}.)

`collection`{.variable}` = ``document`{.variable}` . `[`getElementsByTagNameNS`](#dom-document-getelementsbytagnamens){#ref-for-dom-document-getelementsbytagnamens① .idl-code link-type="method"}`(``namespace`{.variable}`, ``localName`{.variable}`)`

:   If `namespace`{.variable} and `localName`{.variable} are \"`*`\",
    returns an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection②⑥
    link-type="idl"} of all
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②⑤
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑥③
    link-type="dfn"}.

    If only `namespace`{.variable} is \"`*`\", returns an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection②⑦
    link-type="idl"} of all
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②⑥
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑥④
    link-type="dfn"} whose [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name④
    link-type="dfn"} is `localName`{.variable}.

    If only `localName`{.variable} is \"`*`\", returns an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection②⑧
    link-type="idl"} of all
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②⑦
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑥⑤
    link-type="dfn"} whose
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace⑨
    link-type="dfn"} is `namespace`{.variable}.

    Otherwise, returns an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection②⑨
    link-type="idl"} of all
    [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant②⑧
    link-type="dfn"}
    [elements](#concept-element){#ref-for-concept-element⑥⑥
    link-type="dfn"} whose
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①⓪
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name⑤
    link-type="dfn"} is `localName`{.variable}.

`collection`{.variable}` = ``document`{.variable}` . `[`getElementsByClassName`](#dom-document-getelementsbyclassname){#ref-for-dom-document-getelementsbyclassname① .idl-code link-type="method"}`(``classNames`{.variable}`)`\
`collection`{.variable}` = ``element`{.variable}` . `[`getElementsByClassName`](#dom-element-getelementsbyclassname){#ref-for-dom-element-getelementsbyclassname .idl-code link-type="method"}`(``classNames`{.variable}`)`

:   Returns an
    [`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection③⓪
    link-type="idl"} of the
    [elements](#concept-element){#ref-for-concept-element⑥⑦
    link-type="dfn"} in the object on which the method was invoked (a
    [document](#concept-document){#ref-for-concept-document③⑨
    link-type="dfn"} or an
    [element](#concept-element){#ref-for-concept-element⑥⑧
    link-type="dfn"}) that have all the classes given by
    `classNames`{.variable}. The `classNames`{.variable} argument is
    interpreted as a space-separated list of classes.

The [`doctype`]{#dom-document-doctype .dfn .dfn-paneled .idl-code
dfn-for="Document" dfn-type="attribute" export=""} getter steps are to
return the [child](#concept-tree-child){#ref-for-concept-tree-child⑥⓪
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②⑧
link-type="dfn"} that is a
[doctype](#concept-doctype){#ref-for-concept-doctype①⑤ link-type="dfn"};
otherwise null.

The [`documentElement`]{#dom-document-documentelement .dfn .dfn-paneled
.idl-code dfn-for="Document" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①②⑨
link-type="dfn"}'s [document
element](#document-element){#ref-for-document-element⑥ link-type="dfn"}.

The
[`getElementsByTagName(``qualifiedName`{.variable}`)`]{#dom-document-getelementsbytagname
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are to return the [list of elements with
qualified name
`qualifiedName`{.variable}](#concept-getelementsbytagname){#ref-for-concept-getelementsbytagname
link-type="dfn"} for
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③⓪
link-type="dfn"}.

Thus, in an [HTML document](#html-document){#ref-for-html-document③
link-type="dfn"},
`document.getElementsByTagName("FOO")`{.lang-javascript .highlight} will
match `<FOO>` elements that are not in the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace⑤
link-type="dfn"}, and `<foo>` elements that are in the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace⑥
link-type="dfn"}, but not `<FOO>` elements that are in the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace⑦
link-type="dfn"}.

The
[`getElementsByTagNameNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-document-getelementsbytagnamens
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are to return the [list of elements with
namespace `namespace`{.variable} and local name
`localName`{.variable}](#concept-getelementsbytagnamens){#ref-for-concept-getelementsbytagnamens
link-type="dfn"} for
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③①
link-type="dfn"}.

The
[`getElementsByClassName(``classNames`{.variable}`)`]{#dom-document-getelementsbyclassname
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are to return the [list of elements with class
names
`classNames`{.variable}](#concept-getelementsbyclassname){#ref-for-concept-getelementsbyclassname
link-type="dfn"} for
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③②
link-type="dfn"}.

::: {#example-5ffcda00 .example}
[](#example-5ffcda00){.self-link} Given the following XHTML fragment:

``` {.lang-html .highlight}
<div id="example">
  <p id="p1" class="aaa bbb"/>
  <p id="p2" class="aaa ccc"/>
  <p id="p3" class="bbb ccc"/>
</div>
```

A call to
`document.getElementById("example").getElementsByClassName("aaa")`{.lang-javascript
.highlight} would return an
[`HTMLCollection`{.idl}](#htmlcollection){#ref-for-htmlcollection③①
link-type="idl"} with the two paragraphs `p1` and `p2` in it.

A call to `getElementsByClassName("ccc bbb")`{.lang-javascript
.highlight} would only return one node, however, namely `p3`. A call to
`document.getElementById("example").getElementsByClassName("bbb  ccc ")`{.lang-javascript
.highlight} would return the same thing.

A call to `getElementsByClassName("aaa,bbb")`{.lang-javascript
.highlight} would return no nodes; none of the elements above are in the
`aaa,bbb` class.
:::

------------------------------------------------------------------------

`element`{.variable}` = ``document`{.variable}` . `[`createElement(localName [, options])`](#dom-document-createelement){#ref-for-dom-document-createelement① .idl-code link-type="method"}

:   Returns an [element](#concept-element){#ref-for-concept-element⑥⑨
    link-type="dfn"} with `localName`{.variable} as [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name⑥
    link-type="dfn"} (if `document`{.variable} is an [HTML
    document](#html-document){#ref-for-html-document④ link-type="dfn"},
    `localName`{.variable} gets lowercased). The
    [element](#concept-element){#ref-for-concept-element⑦⓪
    link-type="dfn"}'s
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①①
    link-type="dfn"} is the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace⑧
    link-type="dfn"} when `document`{.variable} is an [HTML
    document](#html-document){#ref-for-html-document⑤ link-type="dfn"}
    or `document`{.variable}'s [content
    type](#concept-document-content-type){#ref-for-concept-document-content-type④
    link-type="dfn"} is \"`application/xhtml+xml`\"; otherwise null.

    When supplied, `options`{.variable}'s
    [`customElementRegistry`{.idl}](#dom-elementcreationoptions-customelementregistry){#ref-for-dom-elementcreationoptions-customelementregistry
    link-type="idl"} can be used to set the
    [`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry⑦
    link-type="idl"}.

    When supplied, `options`{.variable}'s
    [`is`{.idl}](#dom-elementcreationoptions-is){#ref-for-dom-elementcreationoptions-is
    link-type="idl"} can be used to create a [customized built-in
    element](https://html.spec.whatwg.org/multipage/custom-elements.html#customized-built-in-element){#ref-for-customized-built-in-element
    link-type="dfn"}.

    If `localName`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name .css
    link-type="type"} production an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④⓪
    link-type="idl"} will be thrown.

    When both `options`{.variable}'s
    [`customElementRegistry`{.idl}](#dom-elementcreationoptions-customelementregistry){#ref-for-dom-elementcreationoptions-customelementregistry①
    link-type="idl"} and `options`{.variable}'s
    [`is`{.idl}](#dom-elementcreationoptions-is){#ref-for-dom-elementcreationoptions-is①
    link-type="idl"} are supplied, a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④①
    link-type="idl"} will be thrown.

`element`{.variable}` = ``document`{.variable}` . `[`createElementNS(namespace, qualifiedName [, options])`](#dom-document-createelementns){#ref-for-dom-document-createelementns① .idl-code link-type="method"}

:   Returns an [element](#concept-element){#ref-for-concept-element⑦①
    link-type="dfn"} with
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①②
    link-type="dfn"} `namespace`{.variable}. Its [namespace
    prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix⑤
    link-type="dfn"} will be everything before U+003A (:) in
    `qualifiedName`{.variable} or null. Its [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name⑦
    link-type="dfn"} will be everything after U+003A (:) in
    `qualifiedName`{.variable} or `qualifiedName`{.variable}.

    When supplied, `options`{.variable}'s
    [`customElementRegistry`{.idl}](#dom-elementcreationoptions-customelementregistry){#ref-for-dom-elementcreationoptions-customelementregistry②
    link-type="idl"} can be used to set the
    [`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry⑧
    link-type="idl"}.

    When supplied, `options`{.variable}'s
    [`is`{.idl}](#dom-elementcreationoptions-is){#ref-for-dom-elementcreationoptions-is②
    link-type="idl"} can be used to create a [customized built-in
    element](https://html.spec.whatwg.org/multipage/custom-elements.html#customized-built-in-element){#ref-for-customized-built-in-element①
    link-type="dfn"}.

    If `qualifiedName`{.variable} does not match the
    [`QName`](https://www.w3.org/TR/xml-names/#NT-QName){#ref-for-NT-QName①
    .css link-type="type"} production an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④②
    link-type="idl"} will be thrown.

    If one of the following conditions is true a
    \"[`NamespaceError`{.idl}](https://webidl.spec.whatwg.org/#namespaceerror){#ref-for-namespaceerror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④③
    link-type="idl"} will be thrown:

    - [Namespace
      prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix⑥
      link-type="dfn"} is not null and `namespace`{.variable} is the
      empty string.
    - [Namespace
      prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix⑦
      link-type="dfn"} is \"`xml`\" and `namespace`{.variable} is not
      the [XML
      namespace](https://infra.spec.whatwg.org/#xml-namespace){#ref-for-xml-namespace②
      link-type="dfn"}.
    - `qualifiedName`{.variable} or [namespace
      prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix⑧
      link-type="dfn"} is \"`xmlns`\" and `namespace`{.variable} is not
      the [XMLNS
      namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#ref-for-xmlns-namespace⑤
      link-type="dfn"}.
    - `namespace`{.variable} is the [XMLNS
      namespace](https://infra.spec.whatwg.org/#xmlns-namespace){#ref-for-xmlns-namespace⑥
      link-type="dfn"} and neither `qualifiedName`{.variable} nor
      [namespace
      prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix⑨
      link-type="dfn"} is \"`xmlns`\".

    When both `options`{.variable}'s
    [`customElementRegistry`{.idl}](#dom-elementcreationoptions-customelementregistry){#ref-for-dom-elementcreationoptions-customelementregistry③
    link-type="idl"} and `options`{.variable}'s
    [`is`{.idl}](#dom-elementcreationoptions-is){#ref-for-dom-elementcreationoptions-is③
    link-type="idl"} are supplied, a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④④
    link-type="idl"} will be thrown.

`documentFragment`{.variable}` = ``document`{.variable}` . `[`createDocumentFragment()`{.idl}](#dom-document-createdocumentfragment){#ref-for-dom-document-createdocumentfragment① link-type="idl"}
:   Returns a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②⑥
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①②⑧
    link-type="dfn"}.

`text`{.variable}` = ``document`{.variable}` . `[`createTextNode(data)`{.idl}](#dom-document-createtextnode){#ref-for-dom-document-createtextnode① link-type="idl"}
:   Returns a [`Text`{.idl}](#text){#ref-for-text②⑤ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node①②⑨ link-type="dfn"}
    whose [data](#concept-cd-data){#ref-for-concept-cd-data①④
    link-type="dfn"} is `data`{.variable}.

`text`{.variable}` = ``document`{.variable}` . `[`createCDATASection(data)`{.idl}](#dom-document-createcdatasection){#ref-for-dom-document-createcdatasection① link-type="idl"}
:   Returns a
    [`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection⑦
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①③⓪
    link-type="dfn"} whose
    [data](#concept-cd-data){#ref-for-concept-cd-data①⑤ link-type="dfn"}
    is `data`{.variable}.

`comment`{.variable}` = ``document`{.variable}` . `[`createComment(data)`{.idl}](#dom-document-createcomment){#ref-for-dom-document-createcomment① link-type="idl"}
:   Returns a [`Comment`{.idl}](#comment){#ref-for-comment①②
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①③①
    link-type="dfn"} whose
    [data](#concept-cd-data){#ref-for-concept-cd-data①⑥ link-type="dfn"}
    is `data`{.variable}.

`processingInstruction`{.variable}` = ``document`{.variable}` . `[`createProcessingInstruction(target, data)`{.idl}](#dom-document-createprocessinginstruction){#ref-for-dom-document-createprocessinginstruction① link-type="idl"}
:   Returns a
    [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①②
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①③②
    link-type="dfn"} whose
    [target](#concept-pi-target){#ref-for-concept-pi-target④
    link-type="dfn"} is `target`{.variable} and
    [data](#concept-cd-data){#ref-for-concept-cd-data①⑦ link-type="dfn"}
    is `data`{.variable}. If `target`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name① .css
    link-type="type"} production an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④⑤
    link-type="idl"} will be thrown. If `data`{.variable} contains
    \"`?>`\" an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④⑥
    link-type="idl"} will be thrown.

The [element interface]{#concept-element-interface .dfn .dfn-paneled
dfn-type="dfn" export=""} for any `name`{.variable} and
`namespace`{.variable} is [`Element`{.idl}](#element){#ref-for-element④⑤
link-type="idl"}, unless stated otherwise.

The HTML Standard will, e.g., define that for `html` and the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace⑨
link-type="dfn"}, the
[`HTMLHtmlElement`{.idl}](https://html.spec.whatwg.org/multipage/semantics.html#htmlhtmlelement){#ref-for-htmlhtmlelement
link-type="idl"} interface is used.
[\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

The
[`createElement(``localName`{.variable}`, ``options`{.variable}`)`]{#dom-document-createelement
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""
lt="createElement(localName, options)|createElement(localName)"} method
steps are:

1.  If `localName`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name② .css
    link-type="type"} production, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③⑧
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④⑦
    link-type="idl"}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③③
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document⑥ link-type="dfn"},
    then set `localName`{.variable} to `localName`{.variable} in [ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase①
    link-type="dfn"}.

3.  Let `registry`{.variable} and `is`{.variable} be the result of
    [flattening element creation
    options](#flatten-element-creation-options){#ref-for-flatten-element-creation-options
    link-type="dfn"} given `options`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③④
    link-type="dfn"}.

4.  Let `namespace`{.variable} be the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①⓪
    link-type="dfn"}, if
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③⑤
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document⑦ link-type="dfn"}
    or [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③⑥
    link-type="dfn"}'s [content
    type](#concept-document-content-type){#ref-for-concept-document-content-type⑤
    link-type="dfn"} is \"`application/xhtml+xml`\"; otherwise null.

5.  Return the result of [creating an
    element](#concept-create-element){#ref-for-concept-create-element①
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③⑦
    link-type="dfn"}, `localName`{.variable}, `namespace`{.variable},
    null, `is`{.variable}, true, and `registry`{.variable}.

The [internal `createElementNS` steps]{#internal-createelementns-steps
.dfn .dfn-paneled dfn-type="dfn" noexport=""}, given
`document`{.variable}, `namespace`{.variable},
`qualifiedName`{.variable}, and `options`{.variable}, are as follows:

1.  Let `namespace`{.variable}, `prefix`{.variable}, and
    `localName`{.variable} be the result of passing
    `namespace`{.variable} and `qualifiedName`{.variable} to [validate
    and extract](#validate-and-extract){#ref-for-validate-and-extract
    link-type="dfn"}.

2.  Let `registry`{.variable} and `is`{.variable} be the result of
    [flattening element creation
    options](#flatten-element-creation-options){#ref-for-flatten-element-creation-options①
    link-type="dfn"} given `options`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③⑧
    link-type="dfn"}.

3.  Return the result of [creating an
    element](#concept-create-element){#ref-for-concept-create-element②
    link-type="dfn"} given `document`{.variable},
    `localName`{.variable}, `namespace`{.variable}, `prefix`{.variable},
    `is`{.variable}, true, and `registry`{.variable}.

The
[`createElementNS(``namespace`{.variable}`, ``qualifiedName`{.variable}`, ``options`{.variable}`)`]{#dom-document-createelementns
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""
lt="createElementNS(namespace, qualifiedName, options)|createElementNS(namespace, qualifiedName)"}
method steps are to return the result of running the [internal
`createElementNS`
steps](#internal-createelementns-steps){#ref-for-internal-createelementns-steps
link-type="dfn"}, given
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①③⑨
link-type="dfn"}, `namespace`{.variable}, `qualifiedName`{.variable},
and `options`{.variable}.

To [flatten element creation options]{#flatten-element-creation-options
.dfn .dfn-paneled dfn-type="dfn" noexport=""}, given a string or
[`ElementCreationOptions`{.idl}](#dictdef-elementcreationoptions){#ref-for-dictdef-elementcreationoptions②
link-type="idl"} dictionary `options`{.variable} and a
[document](#concept-document){#ref-for-concept-document④⓪
link-type="dfn"} `document`{.variable}:

1.  Let `registry`{.variable} be null.

2.  Let `is`{.variable} be null.

3.  If `options`{.variable} is a dictionary:

    1.  If
        `options`{.variable}\[\"[`customElementRegistry`{.idl}](#dom-elementcreationoptions-customelementregistry){#ref-for-dom-elementcreationoptions-customelementregistry④
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①①
        link-type="dfn"}, then set `registry`{.variable} to it.

    2.  If
        `options`{.variable}\[\"[`is`{.idl}](#dom-elementcreationoptions-is){#ref-for-dom-elementcreationoptions-is④
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①②
        link-type="dfn"}, then set `is`{.variable} to it.

    3.  If `registry`{.variable} is non-null and `is`{.variable} is
        non-null, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw③⑨
        link-type="dfn"} a
        \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror③
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④⑧
        link-type="idl"}.

4.  If `registry`{.variable} is null, then set `registry`{.variable} to
    the result of [looking up a custom element
    registry](https://html.spec.whatwg.org/multipage/custom-elements.html#look-up-a-custom-element-registry){#ref-for-look-up-a-custom-element-registry②
    link-type="dfn"} given `document`{.variable}.

5.  Return `registry`{.variable} and `is`{.variable}.

[`createElement()`{.idl}](#dom-document-createelement){#ref-for-dom-document-createelement②
link-type="idl"} and
[`createElementNS()`{.idl}](#dom-document-createelementns){#ref-for-dom-document-createelementns②
link-type="idl"}'s `options`{.variable} parameter is allowed to be a
string for web compatibility.

The [`createDocumentFragment()`]{#dom-document-createdocumentfragment
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are to return a new
[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②⑦
link-type="idl"} [node](#concept-node){#ref-for-concept-node①③③
link-type="dfn"} whose [node
document](#concept-node-document){#ref-for-concept-node-document③③
link-type="dfn"} is
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④⓪
link-type="dfn"}.

The [`createTextNode(``data`{.variable}`)`]{#dom-document-createtextnode
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are to return a new
[`Text`{.idl}](#text){#ref-for-text②⑥ link-type="idl"}
[node](#concept-node){#ref-for-concept-node①③④ link-type="dfn"} whose
[data](#concept-cd-data){#ref-for-concept-cd-data①⑧ link-type="dfn"} is
`data`{.variable} and [node
document](#concept-node-document){#ref-for-concept-node-document③④
link-type="dfn"} is
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④①
link-type="dfn"}.

No check is performed that `data`{.variable} consists of characters that
match the [`Char`](https://www.w3.org/TR/xml/#NT-Char){#ref-for-NT-Char
.css link-type="type"} production.

The
[`createCDATASection(``data`{.variable}`)`]{#dom-document-createcdatasection
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④②
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document⑧ link-type="dfn"},
    then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④⓪
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException④⑨
    link-type="idl"}.

2.  If `data`{.variable} contains the string \"`]]>`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④①
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤⓪
    link-type="idl"}.

3.  Return a new
    [`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①③⑤
    link-type="dfn"} with its
    [data](#concept-cd-data){#ref-for-concept-cd-data①⑨ link-type="dfn"}
    set to `data`{.variable} and [node
    document](#concept-node-document){#ref-for-concept-node-document③⑤
    link-type="dfn"} set to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④③
    link-type="dfn"}.

The [`createComment(``data`{.variable}`)`]{#dom-document-createcomment
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are to return a new
[`Comment`{.idl}](#comment){#ref-for-comment①③ link-type="idl"}
[node](#concept-node){#ref-for-concept-node①③⑥ link-type="dfn"} whose
[data](#concept-cd-data){#ref-for-concept-cd-data②⓪ link-type="dfn"} is
`data`{.variable} and [node
document](#concept-node-document){#ref-for-concept-node-document③⑥
link-type="dfn"} is
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④④
link-type="dfn"}.

No check is performed that `data`{.variable} consists of characters that
match the [`Char`](https://www.w3.org/TR/xml/#NT-Char){#ref-for-NT-Char①
.css link-type="type"} production or that it contains two adjacent
hyphens or ends with a hyphen.

The
[`createProcessingInstruction(``target`{.variable}`, ``data`{.variable}`)`]{#dom-document-createprocessinginstruction
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are:

1.  If `target`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name③ .css
    link-type="type"} production, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④②
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror⑦
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤①
    link-type="idl"}.
2.  If `data`{.variable} contains the string \"`?>`\", then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④③
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror⑧
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤②
    link-type="idl"}.
3.  Return a new
    [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①③
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①③⑦
    link-type="dfn"}, with
    [target](#concept-pi-target){#ref-for-concept-pi-target⑤
    link-type="dfn"} set to `target`{.variable},
    [data](#concept-cd-data){#ref-for-concept-cd-data②① link-type="dfn"}
    set to `data`{.variable}, and [node
    document](#concept-node-document){#ref-for-concept-node-document③⑦
    link-type="dfn"} set to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④⑤
    link-type="dfn"}.

No check is performed that `target`{.variable} contains \"`xml`\" or
\"`:`\", or that `data`{.variable} contains characters that match the
[`Char`](https://www.w3.org/TR/xml/#NT-Char){#ref-for-NT-Char② .css
link-type="type"} production.

------------------------------------------------------------------------

`clone`{.variable} = `document`{.variable} . [importNode(`node`{.variable} \[, `options`{.variable} = false\])](#dom-document-importnode){#ref-for-dom-document-importnode① .idl-code link-type="method"}

:   Returns a copy of `node`{.variable}. If `options`{.variable} is true
    or `options`{.variable} is a dictionary whose
    [`selfOnly`{.idl}](#dom-importnodeoptions-selfonly){#ref-for-dom-importnodeoptions-selfonly
    link-type="idl"} is false, the copy also includes the
    `node`{.variable}'s
    [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant②⑨
    link-type="dfn"}.

    `options`{.variable}'s
    [`customElementRegistry`{.idl}](#dom-importnodeoptions-customelementregistry){#ref-for-dom-importnodeoptions-customelementregistry
    link-type="idl"} can be used to set the
    [`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry⑨
    link-type="idl"} of elements that have none.

    If `node`{.variable} is a
    [document](#concept-document){#ref-for-concept-document④①
    link-type="dfn"} or a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root②①
    link-type="dfn"}, throws a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤③
    link-type="idl"}.

`node`{.variable} = `document`{.variable} . [`adoptNode(node)`{.idl}](#dom-document-adoptnode){#ref-for-dom-document-adoptnode① link-type="idl"}

:   Moves `node`{.variable} from another
    [document](#concept-document){#ref-for-concept-document④②
    link-type="dfn"} and returns it.

    If `node`{.variable} is a
    [document](#concept-document){#ref-for-concept-document④③
    link-type="dfn"}, throws a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤④
    link-type="idl"} or, if `node`{.variable} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root②②
    link-type="dfn"}, throws a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤⑤
    link-type="idl"}.

::: {.algorithm algorithm="importNode(node, options)" algorithm-for="Document"}
The
[`importNode(``node`{.variable}`, ``options`{.variable}`)`]{#dom-document-importnode
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export="" lt="importNode(node, options)|importNode(node)"} method steps
are:

1.  If `node`{.variable} is a
    [document](#concept-document){#ref-for-concept-document④④
    link-type="dfn"} or [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root②③
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④④
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror⑦
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤⑥
    link-type="idl"}.

2.  Let `subtree`{.variable} be false.

3.  Let `registry`{.variable} be null.

4.  If `options`{.variable} is a boolean, then set `subtree`{.variable}
    to `options`{.variable}.

5.  Otherwise:

    1.  Set `subtree`{.variable} to the negation of
        `options`{.variable}\[\"[`selfOnly`{.idl}](#dom-importnodeoptions-selfonly){#ref-for-dom-importnodeoptions-selfonly①
        link-type="idl"}\"\].

    2.  If
        `options`{.variable}\[\"[`customElementRegistry`{.idl}](#dom-importnodeoptions-customelementregistry){#ref-for-dom-importnodeoptions-customelementregistry①
        link-type="idl"}\"\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①③
        link-type="dfn"}, then set `registry`{.variable} to it.

6.  If `registry`{.variable} is null, then set `registry`{.variable} to
    the result of [looking up a custom element
    registry](https://html.spec.whatwg.org/multipage/custom-elements.html#look-up-a-custom-element-registry){#ref-for-look-up-a-custom-element-registry③
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④⑥
    link-type="dfn"}.

7.  Return the result of [cloning a
    node](#concept-node-clone){#ref-for-concept-node-clone④
    link-type="dfn"} given `node`{.variable} with
    [*document*](#clone-a-node-document){#ref-for-clone-a-node-document②
    link-type="dfn"} set to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④⑦
    link-type="dfn"},
    [*subtree*](#clone-a-node-subtree){#ref-for-clone-a-node-subtree③
    link-type="dfn"} set to `subtree`{.variable}, and
    [*fallbackRegistry*](#clone-a-node-fallbackregistry){#ref-for-clone-a-node-fallbackregistry②
    link-type="dfn"} set to `registry`{.variable}.
:::

[Specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications⑧
link-type="dfn"} may define [adopting steps]{#concept-node-adopt-ext
.dfn .dfn-paneled dfn-type="dfn" export=""} for all or some
[nodes](#concept-node){#ref-for-concept-node①③⑧ link-type="dfn"}. The
algorithm is passed `node`{.variable} and `oldDocument`{.variable}, as
indicated in the
[adopt](#concept-node-adopt){#ref-for-concept-node-adopt②
link-type="dfn"} algorithm.

To [adopt]{#concept-node-adopt .dfn .dfn-paneled dfn-type="dfn"
export=""} a `node`{.variable} into a `document`{.variable}, run these
steps:

1.  Let `oldDocument`{.variable} be `node`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document③⑧
    link-type="dfn"}.

2.  If `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent②⑨
    link-type="dfn"} is non-null, then
    [remove](#concept-node-remove){#ref-for-concept-node-remove①⓪
    link-type="dfn"} `node`{.variable}.

3.  If `document`{.variable} is not `oldDocument`{.variable}:

    1.  For each `inclusiveDescendant`{.variable} in `node`{.variable}'s
        [shadow-including inclusive
        descendants](#concept-shadow-including-inclusive-descendant){#ref-for-concept-shadow-including-inclusive-descendant③
        link-type="dfn"}:

        1.  Set `inclusiveDescendant`{.variable}'s [node
            document](#concept-node-document){#ref-for-concept-node-document③⑨
            link-type="dfn"} to `document`{.variable}.

        2.  If `inclusiveDescendant`{.variable} is an
            [element](#concept-element){#ref-for-concept-element⑦②
            link-type="dfn"}, then set the [node
            document](#concept-node-document){#ref-for-concept-node-document④⓪
            link-type="dfn"} of each
            [attribute](#concept-attribute){#ref-for-concept-attribute②①
            link-type="dfn"} in `inclusiveDescendant`{.variable}'s
            [attribute
            list](#concept-element-attribute){#ref-for-concept-element-attribute⑤
            link-type="dfn"} to `document`{.variable}.

    2.  For each `inclusiveDescendant`{.variable} in `node`{.variable}'s
        [shadow-including inclusive
        descendants](#concept-shadow-including-inclusive-descendant){#ref-for-concept-shadow-including-inclusive-descendant④
        link-type="dfn"} that is
        [custom](#concept-element-custom){#ref-for-concept-element-custom⑤
        link-type="dfn"}, [enqueue a custom element callback
        reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-callback-reaction){#ref-for-enqueue-a-custom-element-callback-reaction④
        link-type="dfn"} with `inclusiveDescendant`{.variable}, callback
        name \"`adoptedCallback`\", and « `oldDocument`{.variable},
        `document`{.variable} ».

    3.  For each `inclusiveDescendant`{.variable} in `node`{.variable}'s
        [shadow-including inclusive
        descendants](#concept-shadow-including-inclusive-descendant){#ref-for-concept-shadow-including-inclusive-descendant⑤
        link-type="dfn"}, in [shadow-including tree
        order](#concept-shadow-including-tree-order){#ref-for-concept-shadow-including-tree-order④
        link-type="dfn"}, run the [adopting
        steps](#concept-node-adopt-ext){#ref-for-concept-node-adopt-ext
        link-type="dfn"} with `inclusiveDescendant`{.variable} and
        `oldDocument`{.variable}.

The [`adoptNode(``node`{.variable}`)`]{#dom-document-adoptnode .dfn
.dfn-paneled .idl-code dfn-for="Document" dfn-type="method" export=""}
method steps are:

1.  If `node`{.variable} is a
    [document](#concept-document){#ref-for-concept-document④⑤
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④⑤
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror⑧
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤⑦
    link-type="idl"}.

2.  If `node`{.variable} is a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root②④
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④⑥
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤⑧
    link-type="idl"}.

3.  If `node`{.variable} is a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①③⑨
    link-type="dfn"} whose
    [host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host④
    link-type="dfn"} is non-null, then return.

4.  [Adopt](#concept-node-adopt){#ref-for-concept-node-adopt③
    link-type="dfn"} `node`{.variable} into
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④⑧
    link-type="dfn"}.

5.  Return `node`{.variable}.

------------------------------------------------------------------------

The
[`createAttribute(``localName`{.variable}`)`]{#dom-document-createattribute
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are:

1.  If `localName`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name④ .css
    link-type="type"} production in XML, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④⑦
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror⑨
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑤⑨
    link-type="idl"}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①④⑨
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document⑨ link-type="dfn"},
    then set `localName`{.variable} to `localName`{.variable} in [ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase②
    link-type="dfn"}.

3.  Return a new
    [attribute](#concept-attribute){#ref-for-concept-attribute②②
    link-type="dfn"} whose [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name⑦
    link-type="dfn"} is `localName`{.variable} and [node
    document](#concept-node-document){#ref-for-concept-node-document④①
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤⓪
    link-type="dfn"}.

The
[`createAttributeNS(``namespace`{.variable}`, ``qualifiedName`{.variable}`)`]{#dom-document-createattributens
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are:

1.  Let `namespace`{.variable}, `prefix`{.variable}, and
    `localName`{.variable} be the result of passing
    `namespace`{.variable} and `qualifiedName`{.variable} to [validate
    and extract](#validate-and-extract){#ref-for-validate-and-extract①
    link-type="dfn"}.

2.  Return a new
    [attribute](#concept-attribute){#ref-for-concept-attribute②③
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace⑥
    link-type="dfn"} is `namespace`{.variable}, [namespace
    prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix④
    link-type="dfn"} is `prefix`{.variable}, [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name⑧
    link-type="dfn"} is `localName`{.variable}, and [node
    document](#concept-node-document){#ref-for-concept-node-document④②
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤①
    link-type="dfn"}.

------------------------------------------------------------------------

The [`createEvent(``interface`{.variable}`)`]{#dom-document-createevent
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""} method steps are:

1.  Let `constructor`{.variable} be null.

2.  If `interface`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive①
    link-type="dfn"} match for any of the strings in the first column in
    the following table, then set `constructor`{.variable} to the
    interface in the second column on the same row as the matching
    string:

    String

    Interface

    Notes

    \"`beforeunloadevent`\"

    [`BeforeUnloadEvent`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#beforeunloadevent){#ref-for-beforeunloadevent
    link-type="idl"}

    [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    \"`compositionevent`\"

    [`CompositionEvent`{.idl}](https://w3c.github.io/uievents/#compositionevent){#ref-for-compositionevent
    link-type="idl"}

    [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    \"`customevent`\"

    [`CustomEvent`{.idl}](#customevent){#ref-for-customevent③
    link-type="idl"}

    \"`devicemotionevent`\"

    [`DeviceMotionEvent`{.idl}](https://w3c.github.io/deviceorientation/spec-source-orientation.html#devicemotion){#ref-for-devicemotion
    link-type="idl"}

    [\[DEVICE-ORIENTATION\]](#biblio-device-orientation "Device Orientation and Motion"){link-type="biblio"}

    \"`deviceorientationevent`\"

    [`DeviceOrientationEvent`{.idl}](https://w3c.github.io/deviceorientation/spec-source-orientation.html#devicemotion){#ref-for-devicemotion①
    link-type="idl" refhint-key="116e25eb"}

    \"`dragevent`\"

    [`DragEvent`{.idl}](https://html.spec.whatwg.org/multipage/dnd.html#dragevent){#ref-for-dragevent
    link-type="idl"}

    [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    \"`event`\"

    [`Event`{.idl}](#event){#ref-for-event①⑧ link-type="idl"}

    \"`events`\"

    \"`focusevent`\"

    [`FocusEvent`{.idl}](https://w3c.github.io/uievents/#focusevent){#ref-for-focusevent
    link-type="idl"}

    [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    \"`hashchangeevent`\"

    [`HashChangeEvent`{.idl}](https://html.spec.whatwg.org/multipage/nav-history-apis.html#hashchangeevent){#ref-for-hashchangeevent
    link-type="idl"}

    [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    \"`htmlevents`\"

    [`Event`{.idl}](#event){#ref-for-event①⑨ link-type="idl"}

    \"`keyboardevent`\"

    [`KeyboardEvent`{.idl}](https://w3c.github.io/uievents/#keyboardevent){#ref-for-keyboardevent
    link-type="idl"}

    [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    \"`messageevent`\"

    [`MessageEvent`{.idl}](https://html.spec.whatwg.org/multipage/comms.html#messageevent){#ref-for-messageevent
    link-type="idl"}

    [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    \"`mouseevent`\"

    [`MouseEvent`{.idl}](https://w3c.github.io/uievents/#mouseevent){#ref-for-mouseevent③
    link-type="idl"}

    [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    \"`mouseevents`\"

    \"`storageevent`\"

    [`StorageEvent`{.idl}](https://html.spec.whatwg.org/multipage/webstorage.html#storageevent){#ref-for-storageevent
    link-type="idl"}

    [\[HTML\]](#biblio-html "HTML Standard"){link-type="biblio"}

    \"`svgevents`\"

    [`Event`{.idl}](#event){#ref-for-event②⓪ link-type="idl"}

    \"`textevent`\"

    [`TextEvent`{.idl}](https://w3c.github.io/uievents/#textevent){#ref-for-textevent
    link-type="idl"}

    [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    \"`touchevent`\"

    [`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent③
    link-type="idl"}

    [\[TOUCH-EVENTS\]](#biblio-touch-events "Touch Events"){link-type="biblio"}

    \"`uievent`\"

    [`UIEvent`{.idl}](https://w3c.github.io/uievents/#uievent){#ref-for-uievent
    link-type="idl"}

    [\[UIEVENTS\]](#biblio-uievents "UI Events"){link-type="biblio"}

    \"`uievents`\"

3.  If `constructor`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④⑧
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror⑨
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥⓪
    link-type="idl"}.

4.  If the interface indicated by `constructor`{.variable} is not
    exposed on the [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global④
    link-type="dfn"} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤②
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw④⑨
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥①
    link-type="idl"}.

    Typically user agents disable support for touch events in some
    configurations, in which case this clause would be triggered for the
    interface
    [`TouchEvent`{.idl}](https://w3c.github.io/touch-events/#idl-def-touchevent){#ref-for-idl-def-touchevent④
    link-type="idl"}.

5.  Let `event`{.variable} be the result of [creating an
    event](#concept-event-create){#ref-for-concept-event-create④
    link-type="dfn"} given `constructor`{.variable}.

6.  Initialize `event`{.variable}'s
    [`type`{.idl}](#dom-event-type){#ref-for-dom-event-type①⑧
    link-type="idl"} attribute to the empty string.

7.  Initialize `event`{.variable}'s
    [`timeStamp`{.idl}](#dom-event-timestamp){#ref-for-dom-event-timestamp③
    link-type="idl"} attribute to the result of calling [current high
    resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#ref-for-dfn-current-high-resolution-time
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤③
    link-type="dfn"}'s [relevant global
    object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-relevant-global){#ref-for-concept-relevant-global⑤
    link-type="dfn"}.

8.  Initialize `event`{.variable}'s
    [`isTrusted`{.idl}](#dom-event-istrusted){#ref-for-dom-event-istrusted⑨
    link-type="idl"} attribute to false.

9.  Unset `event`{.variable}'s [initialized
    flag](#initialized-flag){#ref-for-initialized-flag③
    link-type="dfn"}.

10. Return `event`{.variable}.

[Event](#concept-event){#ref-for-concept-event⑤④ link-type="dfn"}
constructors ought to be used instead.

------------------------------------------------------------------------

The [`createRange()`]{#dom-document-createrange .dfn .dfn-paneled
.idl-code dfn-for="Document" dfn-type="method" export=""} method steps
are to return a new [live
range](#concept-live-range){#ref-for-concept-live-range⑧
link-type="dfn"} with
([this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤④
link-type="dfn"}, 0) as its
[start](#concept-range-start){#ref-for-concept-range-start
link-type="dfn"} an [end](#concept-range-end){#ref-for-concept-range-end
link-type="dfn"}.

The [`Range()`{.idl}](#dom-range-range){#ref-for-dom-range-range
link-type="idl"} constructor can be used instead.

------------------------------------------------------------------------

The
[`createNodeIterator(``root`{.variable}`, ``whatToShow`{.variable}`, ``filter`{.variable}`)`]{#dom-document-createnodeiterator
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""
lt="createNodeIterator(root, whatToShow, filter)|createNodeIterator(root, whatToShow)|createNodeIterator(root)"}
method steps are:

1.  Let `iterator`{.variable} be a new
    [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator③
    link-type="idl"} object.

2.  Set `iterator`{.variable}'s
    [root](#concept-traversal-root){#ref-for-concept-traversal-root②
    link-type="dfn"} and `iterator`{.variable}'s
    [reference](#nodeiterator-reference){#ref-for-nodeiterator-reference
    link-type="dfn"} to `root`{.variable}.

3.  Set `iterator`{.variable}'s [pointer before
    reference](#nodeiterator-pointer-before-reference){#ref-for-nodeiterator-pointer-before-reference
    link-type="dfn"} to true.

4.  Set `iterator`{.variable}'s
    [whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow
    link-type="dfn"} to `whatToShow`{.variable}.

5.  Set `iterator`{.variable}'s
    [filter](#concept-traversal-filter){#ref-for-concept-traversal-filter
    link-type="dfn"} to `filter`{.variable}.

6.  Return `iterator`{.variable}.

The
[`createTreeWalker(``root`{.variable}`, ``whatToShow`{.variable}`, ``filter`{.variable}`)`]{#dom-document-createtreewalker
.dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
export=""
lt="createTreeWalker(root, whatToShow, filter)|createTreeWalker(root, whatToShow)|createTreeWalker(root)"}
method steps are:

1.  Let `walker`{.variable} be a new
    [`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker①
    link-type="idl"} object.

2.  Set `walker`{.variable}'s
    [root](#concept-traversal-root){#ref-for-concept-traversal-root③
    link-type="dfn"} and `walker`{.variable}'s
    [current](#treewalker-current){#ref-for-treewalker-current
    link-type="dfn"} to `root`{.variable}.

3.  Set `walker`{.variable}'s
    [whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow①
    link-type="dfn"} to `whatToShow`{.variable}.

4.  Set `walker`{.variable}'s
    [filter](#concept-traversal-filter){#ref-for-concept-traversal-filter①
    link-type="dfn"} to `filter`{.variable}.

5.  Return `walker`{.variable}.

#### [4.5.1. ]{.secno}[Interface [`DOMImplementation`{.idl}](#domimplementation){#ref-for-domimplementation③ link-type="idl"}]{.content}[](#interface-domimplementation){.self-link} {#interface-domimplementation .heading .settled level="4.5.1"}

User agents must create a
[`DOMImplementation`{.idl}](#domimplementation){#ref-for-domimplementation④
link-type="idl"} object whenever a
[document](#concept-document){#ref-for-concept-document④⑥
link-type="dfn"} is created and associate it with that
[document](#concept-document){#ref-for-concept-document④⑦
link-type="dfn"}.

``` {.idl .highlight .def}
[Exposed=Window]
interface DOMImplementation {
  [NewObject] DocumentType createDocumentType(DOMString qualifiedName, DOMString publicId, DOMString systemId);
  [NewObject] XMLDocument createDocument(DOMString? namespace, [LegacyNullToEmptyString] DOMString qualifiedName, optional DocumentType? doctype = null);
  [NewObject] Document createHTMLDocument(optional DOMString title);

  boolean hasFeature(); // useless; always returns true
};
```

`doctype`{.variable}` = ``document`{.variable}` . `[`implementation`{.idl}](#dom-document-implementation){#ref-for-dom-document-implementation② link-type="idl"}` . `[`createDocumentType(qualifiedName, publicId, systemId)`{.idl}](#dom-domimplementation-createdocumenttype){#ref-for-dom-domimplementation-createdocumenttype① link-type="idl"}
:   Returns a [doctype](#concept-doctype){#ref-for-concept-doctype①⑥
    link-type="dfn"}, with the given `qualifiedName`{.variable},
    `publicId`{.variable}, and `systemId`{.variable}. If
    `qualifiedName`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name⑤ .css
    link-type="type"} production, an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥②
    link-type="idl"} is thrown, and if it does not match the
    [`QName`](https://www.w3.org/TR/xml-names/#NT-QName){#ref-for-NT-QName②
    .css link-type="type"} production, a
    \"[`NamespaceError`{.idl}](https://webidl.spec.whatwg.org/#namespaceerror){#ref-for-namespaceerror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥③
    link-type="idl"} is thrown.

`doc`{.variable}` = ``document`{.variable}` . `[`implementation`{.idl}](#dom-document-implementation){#ref-for-dom-document-implementation③ link-type="idl"}` . `[`createDocument(``namespace`{.variable}`, ``qualifiedName`{.variable}` [, ``doctype`{.variable}` = null])`](#dom-domimplementation-createdocument){#ref-for-dom-domimplementation-createdocument② .idl-code link-type="method"}

:   Returns an [`XMLDocument`{.idl}](#xmldocument){#ref-for-xmldocument②
    link-type="idl"}, with a [document
    element](#document-element){#ref-for-document-element⑦
    link-type="dfn"} whose [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name⑧
    link-type="dfn"} is `qualifiedName`{.variable} and whose
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①③
    link-type="dfn"} is `namespace`{.variable} (unless
    `qualifiedName`{.variable} is the empty string), and with
    `doctype`{.variable}, if it is given, as its
    [doctype](#concept-doctype){#ref-for-concept-doctype①⑦
    link-type="dfn"}.

    This method throws the same exceptions as the
    [`createElementNS()`{.idl}](#dom-document-createelementns){#ref-for-dom-document-createelementns③
    link-type="idl"} method, when invoked with `namespace`{.variable}
    and `qualifiedName`{.variable}.

`doc`{.variable}` = ``document`{.variable}` . `[`implementation`{.idl}](#dom-document-implementation){#ref-for-dom-document-implementation④ link-type="idl"}` . `[`createHTMLDocument([``title`{.variable}`])`](#dom-domimplementation-createhtmldocument){#ref-for-dom-domimplementation-createhtmldocument① .idl-code link-type="method"}
:   Returns a [document](#concept-document){#ref-for-concept-document④⑧
    link-type="dfn"}, with a basic
    [tree](#concept-tree){#ref-for-concept-tree②① link-type="dfn"}
    already constructed including a
    [`title`](https://html.spec.whatwg.org/multipage/semantics.html#the-title-element){#ref-for-the-title-element
    link-type="element"} element, unless the `title`{.variable} argument
    is omitted.

::: impl
The
[`createDocumentType(``qualifiedName`{.variable}`, ``publicId`{.variable}`, ``systemId`{.variable}`)`]{#dom-domimplementation-createdocumenttype
.dfn .dfn-paneled .idl-code dfn-for="DOMImplementation"
dfn-type="method" export=""} method steps are:

1.  [Validate](#validate){#ref-for-validate① link-type="dfn"}
    `qualifiedName`{.variable}.

2.  Return a new [doctype](#concept-doctype){#ref-for-concept-doctype①⑧
    link-type="dfn"}, with `qualifiedName`{.variable} as its
    [name](#concept-doctype-name){#ref-for-concept-doctype-name④
    link-type="dfn"}, `publicId`{.variable} as its [public
    ID](#concept-doctype-publicid){#ref-for-concept-doctype-publicid②
    link-type="dfn"}, and `systemId`{.variable} as its [system
    ID](#concept-doctype-systemid){#ref-for-concept-doctype-systemid②
    link-type="dfn"}, and with its [node
    document](#concept-node-document){#ref-for-concept-node-document④③
    link-type="dfn"} set to the associated
    [document](#concept-document){#ref-for-concept-document④⑨
    link-type="dfn"} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤⑤
    link-type="dfn"}.

No check is performed that `publicId`{.variable} code points match the
[`PubidChar`](https://www.w3.org/TR/xml/#NT-PubidChar){#ref-for-NT-PubidChar
.css link-type="type"} production or that `systemId`{.variable} does not
contain both a \'`"`\' and a \"`'`\".

The
[`createDocument(``namespace`{.variable}`, ``qualifiedName`{.variable}`, ``doctype`{.variable}`)`]{#dom-domimplementation-createdocument
.dfn .dfn-paneled .idl-code dfn-for="DOMImplementation"
dfn-type="method" export=""
lt="createDocument(namespace, qualifiedName, doctype)|createDocument(namespace, qualifiedName)"}
method steps are:

1.  Let `document`{.variable} be a new
    [`XMLDocument`{.idl}](#xmldocument){#ref-for-xmldocument③
    link-type="idl"}.

2.  Let `element`{.variable} be null.

3.  If `qualifiedName`{.variable} is not the empty string, then set
    `element`{.variable} to the result of running the [internal
    `createElementNS`
    steps](#internal-createelementns-steps){#ref-for-internal-createelementns-steps①
    link-type="dfn"}, given `document`{.variable},
    `namespace`{.variable}, `qualifiedName`{.variable}, and an empty
    dictionary.

4.  If `doctype`{.variable} is non-null,
    [append](#concept-node-append){#ref-for-concept-node-append④
    link-type="dfn"} `doctype`{.variable} to `document`{.variable}.

5.  If `element`{.variable} is non-null,
    [append](#concept-node-append){#ref-for-concept-node-append⑤
    link-type="dfn"} `element`{.variable} to `document`{.variable}.

6.  `document`{.variable}'s
    [origin](#concept-document-origin){#ref-for-concept-document-origin④
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤⑥
    link-type="dfn"}'s associated
    [document](#concept-document){#ref-for-concept-document⑤⓪
    link-type="dfn"}'s
    [origin](#concept-document-origin){#ref-for-concept-document-origin⑤
    link-type="dfn"}.

7.  `document`{.variable}'s [content
    type](#concept-document-content-type){#ref-for-concept-document-content-type⑥
    link-type="dfn"} is determined by `namespace`{.variable}:

    [HTML namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①① link-type="dfn"}
    :   `application/xhtml+xml`

    [SVG namespace](https://infra.spec.whatwg.org/#svg-namespace){#ref-for-svg-namespace link-type="dfn"}
    :   `image/svg+xml`

    Any other namespace
    :   `application/xml`

8.  Return `document`{.variable}.

The
[`createHTMLDocument(``title`{.variable}`)`]{#dom-domimplementation-createhtmldocument
.dfn .dfn-paneled .idl-code dfn-for="DOMImplementation"
dfn-type="method" export=""
lt="createHTMLDocument(title)|createHTMLDocument()"} method steps are:

1.  Let `doc`{.variable} be a new
    [document](#concept-document){#ref-for-concept-document⑤①
    link-type="dfn"} that is an [HTML
    document](#html-document){#ref-for-html-document①⓪ link-type="dfn"}.

2.  Set `doc`{.variable}'s [content
    type](#concept-document-content-type){#ref-for-concept-document-content-type⑦
    link-type="dfn"} to \"`text/html`\".

3.  [Append](#concept-node-append){#ref-for-concept-node-append⑥
    link-type="dfn"} a new
    [doctype](#concept-doctype){#ref-for-concept-doctype①⑨
    link-type="dfn"}, with \"`html`\" as its
    [name](#concept-doctype-name){#ref-for-concept-doctype-name⑤
    link-type="dfn"} and with its [node
    document](#concept-node-document){#ref-for-concept-node-document④④
    link-type="dfn"} set to `doc`{.variable}, to `doc`{.variable}.

4.  [Append](#concept-node-append){#ref-for-concept-node-append⑦
    link-type="dfn"} the result of [creating an
    element](#concept-create-element){#ref-for-concept-create-element③
    link-type="dfn"} given `doc`{.variable}, \"`html`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①②
    link-type="dfn"}, to `doc`{.variable}.

5.  [Append](#concept-node-append){#ref-for-concept-node-append⑧
    link-type="dfn"} the result of [creating an
    element](#concept-create-element){#ref-for-concept-create-element④
    link-type="dfn"} given `doc`{.variable}, \"`head`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①③
    link-type="dfn"}, to the
    [`html`](https://html.spec.whatwg.org/multipage/semantics.html#the-html-element){#ref-for-the-html-element
    link-type="element"} element created earlier.

6.  If `title`{.variable} is given:

    1.  [Append](#concept-node-append){#ref-for-concept-node-append⑨
        link-type="dfn"} the result of [creating an
        element](#concept-create-element){#ref-for-concept-create-element⑤
        link-type="dfn"} given `doc`{.variable}, \"`title`\", and the
        [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①④
        link-type="dfn"}, to the
        [`head`](https://html.spec.whatwg.org/multipage/semantics.html#the-head-element){#ref-for-the-head-element
        link-type="element"} element created earlier.

    2.  [Append](#concept-node-append){#ref-for-concept-node-append①⓪
        link-type="dfn"} a new [`Text`{.idl}](#text){#ref-for-text②⑦
        link-type="idl"} [node](#concept-node){#ref-for-concept-node①④⓪
        link-type="dfn"}, with its
        [data](#concept-cd-data){#ref-for-concept-cd-data②②
        link-type="dfn"} set to `title`{.variable} (which could be the
        empty string) and its [node
        document](#concept-node-document){#ref-for-concept-node-document④⑤
        link-type="dfn"} set to `doc`{.variable}, to the
        [`title`](https://html.spec.whatwg.org/multipage/semantics.html#the-title-element){#ref-for-the-title-element①
        link-type="element"} element created earlier.

7.  [Append](#concept-node-append){#ref-for-concept-node-append①①
    link-type="dfn"} the result of [creating an
    element](#concept-create-element){#ref-for-concept-create-element⑥
    link-type="dfn"} given `doc`{.variable}, \"`body`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①⑤
    link-type="dfn"}, to the
    [`html`](https://html.spec.whatwg.org/multipage/semantics.html#the-html-element){#ref-for-the-html-element①
    link-type="element"} element created earlier.

8.  `doc`{.variable}'s
    [origin](#concept-document-origin){#ref-for-concept-document-origin⑥
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤⑦
    link-type="dfn"}'s associated
    [document](#concept-document){#ref-for-concept-document⑤②
    link-type="dfn"}'s
    [origin](#concept-document-origin){#ref-for-concept-document-origin⑦
    link-type="dfn"}.

9.  Return `doc`{.variable}.

The [`hasFeature()`]{#dom-domimplementation-hasfeature .dfn .dfn-paneled
.idl-code dfn-for="DOMImplementation" dfn-type="method" export=""}
method steps are to return true.

[`hasFeature()`{.idl}](#dom-domimplementation-hasfeature){#ref-for-dom-domimplementation-hasfeature①
link-type="idl"} originally would report whether the user agent claimed
to support a given DOM feature, but experience proved it was not nearly
as reliable or granular as simply checking whether the desired objects,
attributes, or methods existed. As such, it is no longer to be used, but
continues to exist (and simply returns true) so that old pages don't
stop working.
:::

### [4.6. ]{.secno}[Interface [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype②① link-type="idl"}]{.content}[](#interface-documenttype){.self-link} {#interface-documenttype .heading .settled level="4.6"}

``` {.idl .highlight .def}
[Exposed=Window]
interface DocumentType : Node {
  readonly attribute DOMString name;
  readonly attribute DOMString publicId;
  readonly attribute DOMString systemId;
};
```

[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype②②
link-type="idl"} [nodes](#concept-node){#ref-for-concept-node①④①
link-type="dfn"} are simply known as [doctypes]{#concept-doctype .dfn
.dfn-paneled dfn-type="dfn" export="" lt="doctype"}.

[Doctypes](#concept-doctype){#ref-for-concept-doctype②⓪ link-type="dfn"}
have an associated [name]{#concept-doctype-name .dfn .dfn-paneled
dfn-for="DocumentType" dfn-type="dfn" export=""}, [public
ID]{#concept-doctype-publicid .dfn .dfn-paneled dfn-for="DocumentType"
dfn-type="dfn" export=""}, and [system ID]{#concept-doctype-systemid
.dfn .dfn-paneled dfn-for="DocumentType" dfn-type="dfn" export=""}.

When a [doctype](#concept-doctype){#ref-for-concept-doctype②①
link-type="dfn"} is created, its
[name](#concept-doctype-name){#ref-for-concept-doctype-name⑥
link-type="dfn"} is always given. Unless explicitly given when a
[doctype](#concept-doctype){#ref-for-concept-doctype②② link-type="dfn"}
is created, its [public
ID](#concept-doctype-publicid){#ref-for-concept-doctype-publicid③
link-type="dfn"} and [system
ID](#concept-doctype-systemid){#ref-for-concept-doctype-systemid③
link-type="dfn"} are the empty string.

The [`name`]{#dom-documenttype-name .dfn .dfn-paneled .idl-code
dfn-for="DocumentType" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤⑧
link-type="dfn"}'s
[name](#concept-doctype-name){#ref-for-concept-doctype-name⑦
link-type="dfn"}.

The [`publicId`]{#dom-documenttype-publicid .dfn .dfn-paneled .idl-code
dfn-for="DocumentType" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑤⑨
link-type="dfn"}'s [public
ID](#concept-doctype-publicid){#ref-for-concept-doctype-publicid④
link-type="dfn"}.

The [`systemId`]{#dom-documenttype-systemid .dfn .dfn-paneled .idl-code
dfn-for="DocumentType" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥⓪
link-type="dfn"}'s [system
ID](#concept-doctype-systemid){#ref-for-concept-doctype-systemid④
link-type="dfn"}.

### [4.7. ]{.secno}[Interface [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment②⑨ link-type="idl"}]{.content}[](#interface-documentfragment){.self-link} {#interface-documentfragment .heading .settled level="4.7"}

``` {.idl .highlight .def}
[Exposed=Window]
interface DocumentFragment : Node {
  constructor();
};
```

A
[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③⓪
link-type="idl"} [node](#concept-node){#ref-for-concept-node①④②
link-type="dfn"} has an associated [host]{#concept-documentfragment-host
.dfn .dfn-paneled dfn-for="DocumentFragment" dfn-type="dfn" export=""}
(null or an [element](#concept-element){#ref-for-concept-element⑦③
link-type="dfn"} in a different [node
tree](#concept-node-tree){#ref-for-concept-node-tree②①
link-type="dfn"}). It is null unless otherwise stated.

An object `A`{.variable} is a [host-including inclusive
ancestor]{#concept-tree-host-including-inclusive-ancestor .dfn
.dfn-paneled dfn-type="dfn" export=""} of an object `B`{.variable}, if
either `A`{.variable} is an [inclusive
ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor④
link-type="dfn"} of `B`{.variable}, or if `B`{.variable}'s
[root](#concept-tree-root){#ref-for-concept-tree-root②⑦ link-type="dfn"}
has a non-null
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host⑤
link-type="dfn"} and `A`{.variable} is a [host-including inclusive
ancestor](#concept-tree-host-including-inclusive-ancestor){#ref-for-concept-tree-host-including-inclusive-ancestor③
link-type="dfn"} of `B`{.variable}'s
[root](#concept-tree-root){#ref-for-concept-tree-root②⑧
link-type="dfn"}'s
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host⑥
link-type="dfn"}.

The
[`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③①
link-type="idl"} [node](#concept-node){#ref-for-concept-node①④③
link-type="dfn"}'s
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host⑦
link-type="dfn"} concept is useful for HTML's
[`template`](https://html.spec.whatwg.org/multipage/scripting.html#the-template-element){#ref-for-the-template-element①
link-type="element"} element and for [shadow
roots](#concept-shadow-root){#ref-for-concept-shadow-root②⑤
link-type="dfn"}, and impacts the
[pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert⑦
link-type="dfn"} and
[replace](#concept-node-replace){#ref-for-concept-node-replace②
link-type="dfn"} algorithms.

`tree`{.variable}` = new `[`DocumentFragment()`{.idl}](#dom-documentfragment-documentfragment){#ref-for-dom-documentfragment-documentfragment① link-type="idl"}
:   Returns a new
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③②
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①④④
    link-type="dfn"}.

The [`new DocumentFragment()`]{#dom-documentfragment-documentfragment
.dfn .dfn-paneled .idl-code dfn-for="DocumentFragment"
dfn-type="constructor" export="" lt="DocumentFragment()|constructor()"}
constructor steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥①
link-type="dfn"}'s [node
document](#concept-node-document){#ref-for-concept-node-document④⑥
link-type="dfn"} to [current global
object](https://html.spec.whatwg.org/multipage/webappapis.html#current-global-object){#ref-for-current-global-object①
link-type="dfn"}'s [associated
`Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window②
link-type="dfn"}.

### [4.8. ]{.secno}[Interface [`ShadowRoot`{.idl}](#shadowroot){#ref-for-shadowroot⑤ link-type="idl"}]{.content}[](#interface-shadowroot){.self-link} {#interface-shadowroot .heading .settled level="4.8"}

``` {.idl .highlight .def}
[Exposed=Window]
interface ShadowRoot : DocumentFragment {
  readonly attribute ShadowRootMode mode;
  readonly attribute boolean delegatesFocus;
  readonly attribute SlotAssignmentMode slotAssignment;
  readonly attribute boolean clonable;
  readonly attribute boolean serializable;
  readonly attribute Element host;

  attribute EventHandler onslotchange;
};

enum ShadowRootMode { "open", "closed" };
enum SlotAssignmentMode { "manual", "named" };
```

[`ShadowRoot`{.idl}](#shadowroot){#ref-for-shadowroot⑥ link-type="idl"}
[nodes](#concept-node){#ref-for-concept-node①④⑤ link-type="dfn"} are
simply known as [shadow roots]{#concept-shadow-root .dfn .dfn-paneled
dfn-type="dfn" export="" lt="shadow root"}.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root②⑥
link-type="dfn"}'s associated
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host⑧
link-type="dfn"} is never null.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root②⑦
link-type="dfn"} have an associated [mode]{#shadowroot-mode .dfn
.dfn-paneled dfn-for="ShadowRoot" dfn-type="dfn" noexport=""}
(\"`open`\" or \"`closed`\").

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root②⑧
link-type="dfn"} have an associated [delegates
focus]{#shadowroot-delegates-focus .dfn .dfn-paneled
dfn-for="ShadowRoot" dfn-type="dfn" export=""} (a boolean). It is
initially set to false.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root②⑨
link-type="dfn"} have an associated [available to element
internals]{#shadowroot-available-to-element-internals .dfn .dfn-paneled
dfn-for="ShadowRoot" dfn-type="dfn" export=""} (a boolean). It is
initially set to false.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root③⓪
link-type="dfn"} have an associated
[declarative]{#shadowroot-declarative .dfn .dfn-paneled
dfn-for="ShadowRoot" dfn-type="dfn" export=""} (a boolean). It is
initially set to false.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root③①
link-type="dfn"} have an associated [slot
assignment]{#shadowroot-slot-assignment .dfn .dfn-paneled
dfn-for="ShadowRoot" dfn-type="dfn" noexport=""} (\"`manual`\" or
\"`named`\").

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root③②
link-type="dfn"} have an associated [clonable]{#shadowroot-clonable .dfn
.dfn-paneled dfn-for="ShadowRoot" dfn-type="dfn" noexport=""} (a
boolean). It is initially set to false.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root③③
link-type="dfn"} have an associated
[serializable]{#shadowroot-serializable .dfn .dfn-paneled
dfn-for="ShadowRoot" dfn-type="dfn" noexport=""} (a boolean). It is
initially set to false.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root③④
link-type="dfn"} have an associated [custom element
registry]{#shadowroot-custom-element-registry .dfn .dfn-paneled
dfn-for="ShadowRoot" dfn-type="dfn" noexport=""} (null or a
[`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry①⓪
link-type="idl"} object). It is initially null.

[Shadow roots](#concept-shadow-root){#ref-for-concept-shadow-root③⑤
link-type="dfn"} have an associated [keep custom element registry
null]{#shadowroot-keep-custom-element-registry-null .dfn .dfn-paneled
dfn-for="ShadowRoot" dfn-type="dfn" noexport=""} (a boolean). It is
initially false.

This can only ever be true in combination with declarative shadow roots.
And it only matters for as long as the [shadow
root](#concept-shadow-root){#ref-for-concept-shadow-root③⑥
link-type="dfn"}'s [custom element
registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry⑦
link-type="dfn"} is null.

------------------------------------------------------------------------

A [shadow root](#concept-shadow-root){#ref-for-concept-shadow-root③⑦
link-type="dfn"}'s [get the
parent](#get-the-parent){#ref-for-get-the-parent⑦ link-type="dfn"}
algorithm, given an `event`{.variable}, returns null if
`event`{.variable}'s [composed
flag](#composed-flag){#ref-for-composed-flag① link-type="dfn"} is unset
and [shadow root](#concept-shadow-root){#ref-for-concept-shadow-root③⑧
link-type="dfn"} is the
[root](#concept-tree-root){#ref-for-concept-tree-root②⑨ link-type="dfn"}
of `event`{.variable}'s [path](#event-path){#ref-for-event-path①⓪
link-type="dfn"}'s first struct's [invocation
target](#event-path-invocation-target){#ref-for-event-path-invocation-target⑥
link-type="dfn"}; otherwise [shadow
root](#concept-shadow-root){#ref-for-concept-shadow-root③⑨
link-type="dfn"}'s
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host⑨
link-type="dfn"}.

------------------------------------------------------------------------

The [`mode`]{#dom-shadowroot-mode .dfn .dfn-paneled .idl-code
dfn-for="ShadowRoot" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥②
link-type="dfn"}'s [mode](#shadowroot-mode){#ref-for-shadowroot-mode⑤
link-type="dfn"}.

The [`delegatesFocus`]{#dom-shadowroot-delegatesfocus .dfn .dfn-paneled
.idl-code dfn-for="ShadowRoot" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥③
link-type="dfn"}'s [delegates
focus](#shadowroot-delegates-focus){#ref-for-shadowroot-delegates-focus①
link-type="dfn"}.

The [`slotAssignment`]{#dom-shadowroot-slotassignment .dfn .dfn-paneled
.idl-code dfn-for="ShadowRoot" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥④
link-type="dfn"}'s [slot
assignment](#shadowroot-slot-assignment){#ref-for-shadowroot-slot-assignment⑤
link-type="dfn"}.

The [`clonable`]{#dom-shadowroot-clonable .dfn .dfn-paneled .idl-code
dfn-for="ShadowRoot" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥⑤
link-type="dfn"}'s
[clonable](#shadowroot-clonable){#ref-for-shadowroot-clonable①
link-type="dfn"}.

The [`serializable`]{#dom-shadowroot-serializable .dfn .dfn-paneled
.idl-code dfn-for="ShadowRoot" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥⑥
link-type="dfn"}'s
[serializable](#shadowroot-serializable){#ref-for-shadowroot-serializable①
link-type="dfn"}.

The [`host`]{#dom-shadowroot-host .dfn .dfn-paneled .idl-code
dfn-for="ShadowRoot" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥⑦
link-type="dfn"}'s
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host①⓪
link-type="dfn"}.

------------------------------------------------------------------------

The [`onslotchange`]{#dom-shadowroot-onslotchange .dfn .dfn-paneled
.idl-code dfn-for="ShadowRoot" dfn-type="attribute" export=""} attribute
is an [event handler IDL
attribute](https://html.spec.whatwg.org/multipage/webappapis.html#event-handler-idl-attributes){#ref-for-event-handler-idl-attributes①
link-type="dfn"} for the [`onslotchange`]{#shadowroot-onslotchange .dfn
.dfn-paneled dfn-for="ShadowRoot" dfn-type="dfn" export=""} [event
handler](https://html.spec.whatwg.org/multipage/webappapis.html#event-handlers){#ref-for-event-handlers②
link-type="dfn"}, whose [event handler event
type](https://html.spec.whatwg.org/multipage/webappapis.html#event-handler-event-type){#ref-for-event-handler-event-type①
link-type="dfn"} is
[`slotchange`{.idl}](#eventdef-htmlslotelement-slotchange){#ref-for-eventdef-htmlslotelement-slotchange
link-type="idl"}.

------------------------------------------------------------------------

In [shadow-including tree order]{#concept-shadow-including-tree-order
.dfn .dfn-paneled dfn-type="dfn" export=""} is [shadow-including
preorder, depth-first
traversal](#shadow-including-preorder-depth-first-traversal){#ref-for-shadow-including-preorder-depth-first-traversal
link-type="dfn"} of a [node
tree](#concept-node-tree){#ref-for-concept-node-tree②② link-type="dfn"}.
[Shadow-including preorder, depth-first
traversal]{#shadow-including-preorder-depth-first-traversal .dfn
.dfn-paneled dfn-type="dfn" noexport=""} of a [node
tree](#concept-node-tree){#ref-for-concept-node-tree②③ link-type="dfn"}
`tree`{.variable} is preorder, depth-first traversal of
`tree`{.variable}, with for each [shadow
host](#element-shadow-host){#ref-for-element-shadow-host④
link-type="dfn"} encountered in `tree`{.variable}, [shadow-including
preorder, depth-first
traversal](#shadow-including-preorder-depth-first-traversal){#ref-for-shadow-including-preorder-depth-first-traversal①
link-type="dfn"} of that
[element](#concept-element){#ref-for-concept-element⑦④
link-type="dfn"}'s [shadow
root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①①
link-type="dfn"}'s [node
tree](#concept-node-tree){#ref-for-concept-node-tree②④ link-type="dfn"}
just after it is encountered.

The [shadow-including root]{#concept-shadow-including-root .dfn
.dfn-paneled dfn-type="dfn" export=""} of an object is its
[root](#concept-tree-root){#ref-for-concept-tree-root③⓪
link-type="dfn"}'s
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host①①
link-type="dfn"}'s [shadow-including
root](#concept-shadow-including-root){#ref-for-concept-shadow-including-root⑤
link-type="dfn"}, if the object's
[root](#concept-tree-root){#ref-for-concept-tree-root③① link-type="dfn"}
is a [shadow root](#concept-shadow-root){#ref-for-concept-shadow-root④⓪
link-type="dfn"}; otherwise its
[root](#concept-tree-root){#ref-for-concept-tree-root③②
link-type="dfn"}.

An object `A`{.variable} is a [shadow-including
descendant]{#concept-shadow-including-descendant .dfn .dfn-paneled
dfn-type="dfn" export=""} of an object `B`{.variable}, if `A`{.variable}
is a
[descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant③⓪
link-type="dfn"} of `B`{.variable}, or `A`{.variable}'s
[root](#concept-tree-root){#ref-for-concept-tree-root③③ link-type="dfn"}
is a [shadow root](#concept-shadow-root){#ref-for-concept-shadow-root④①
link-type="dfn"} and `A`{.variable}'s
[root](#concept-tree-root){#ref-for-concept-tree-root③④
link-type="dfn"}'s
[host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host①②
link-type="dfn"} is a [shadow-including inclusive
descendant](#concept-shadow-including-inclusive-descendant){#ref-for-concept-shadow-including-inclusive-descendant⑥
link-type="dfn"} of `B`{.variable}.

A [shadow-including inclusive
descendant]{#concept-shadow-including-inclusive-descendant .dfn
.dfn-paneled dfn-type="dfn" export=""} is an object or one of its
[shadow-including
descendants](#concept-shadow-including-descendant){#ref-for-concept-shadow-including-descendant①
link-type="dfn"}.

An object `A`{.variable} is a [shadow-including
ancestor]{#concept-shadow-including-ancestor .dfn .dfn-paneled
dfn-type="dfn" export=""} of an object `B`{.variable}, if and only if
`B`{.variable} is a [shadow-including
descendant](#concept-shadow-including-descendant){#ref-for-concept-shadow-including-descendant②
link-type="dfn"} of `A`{.variable}.

A [shadow-including inclusive
ancestor]{#concept-shadow-including-inclusive-ancestor .dfn .dfn-paneled
dfn-type="dfn" export=""} is an object or one of its [shadow-including
ancestors](#concept-shadow-including-ancestor){#ref-for-concept-shadow-including-ancestor
link-type="dfn"}.

A [node](#concept-node){#ref-for-concept-node①④⑥ link-type="dfn"}
`A`{.variable} is [closed-shadow-hidden]{#concept-closed-shadow-hidden
.dfn .dfn-paneled dfn-type="dfn" export=""} from a
[node](#concept-node){#ref-for-concept-node①④⑦ link-type="dfn"}
`B`{.variable} if all of the following conditions are true:

- `A`{.variable}'s
  [root](#concept-tree-root){#ref-for-concept-tree-root③⑤
  link-type="dfn"} is a [shadow
  root](#concept-shadow-root){#ref-for-concept-shadow-root④②
  link-type="dfn"}.

- `A`{.variable}'s
  [root](#concept-tree-root){#ref-for-concept-tree-root③⑥
  link-type="dfn"} is not a [shadow-including inclusive
  ancestor](#concept-shadow-including-inclusive-ancestor){#ref-for-concept-shadow-including-inclusive-ancestor①
  link-type="dfn"} of `B`{.variable}.

- `A`{.variable}'s
  [root](#concept-tree-root){#ref-for-concept-tree-root③⑦
  link-type="dfn"} is a [shadow
  root](#concept-shadow-root){#ref-for-concept-shadow-root④③
  link-type="dfn"} whose
  [mode](#shadowroot-mode){#ref-for-shadowroot-mode⑥ link-type="dfn"} is
  \"`closed`\" or `A`{.variable}'s
  [root](#concept-tree-root){#ref-for-concept-tree-root③⑧
  link-type="dfn"}'s
  [host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host①③
  link-type="dfn"} is
  [closed-shadow-hidden](#concept-closed-shadow-hidden){#ref-for-concept-closed-shadow-hidden
  link-type="dfn"} from `B`{.variable}.

To [retarget]{#retarget .dfn .dfn-paneled dfn-type="dfn" export=""
lt="retarget|retargeting"} an object `A`{.variable} against an object
`B`{.variable}, repeat these steps until they return an object:

1.  If one of the following is true

    - `A`{.variable} is not a
      [node](#concept-node){#ref-for-concept-node①④⑧ link-type="dfn"}
    - `A`{.variable}'s
      [root](#concept-tree-root){#ref-for-concept-tree-root③⑨
      link-type="dfn"} is not a [shadow
      root](#concept-shadow-root){#ref-for-concept-shadow-root④④
      link-type="dfn"}
    - `B`{.variable} is a [node](#concept-node){#ref-for-concept-node①④⑨
      link-type="dfn"} and `A`{.variable}'s
      [root](#concept-tree-root){#ref-for-concept-tree-root④⓪
      link-type="dfn"} is a [shadow-including inclusive
      ancestor](#concept-shadow-including-inclusive-ancestor){#ref-for-concept-shadow-including-inclusive-ancestor②
      link-type="dfn"} of `B`{.variable}

    then return `A`{.variable}.

2.  Set `A`{.variable} to `A`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root④①
    link-type="dfn"}'s
    [host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host①④
    link-type="dfn"}.

The [retargeting](#retarget){#ref-for-retarget④ link-type="dfn"}
algorithm is used by [event
dispatch](#concept-event-dispatch){#ref-for-concept-event-dispatch②⑦
link-type="dfn"} as well as other specifications, such as Fullscreen.
[\[FULLSCREEN\]](#biblio-fullscreen "Fullscreen API Standard"){link-type="biblio"}

### [4.9. ]{.secno}[Interface [`Element`{.idl}](#element){#ref-for-element④⑦ link-type="idl"}]{.content}[](#interface-element){.self-link} {#interface-element .heading .settled level="4.9"}

``` {.idl .highlight .def}
[Exposed=Window]
interface Element : Node {
  readonly attribute DOMString? namespaceURI;
  readonly attribute DOMString? prefix;
  readonly attribute DOMString localName;
  readonly attribute DOMString tagName;

  [CEReactions] attribute DOMString id;
  [CEReactions] attribute DOMString className;
  [SameObject, PutForwards=value] readonly attribute DOMTokenList classList;
  [CEReactions, Unscopable] attribute DOMString slot;

  boolean hasAttributes();
  [SameObject] readonly attribute NamedNodeMap attributes;
  sequence<DOMString> getAttributeNames();
  DOMString? getAttribute(DOMString qualifiedName);
  DOMString? getAttributeNS(DOMString? namespace, DOMString localName);
  [CEReactions] undefined setAttribute(DOMString qualifiedName, DOMString value);
  [CEReactions] undefined setAttributeNS(DOMString? namespace, DOMString qualifiedName, DOMString value);
  [CEReactions] undefined removeAttribute(DOMString qualifiedName);
  [CEReactions] undefined removeAttributeNS(DOMString? namespace, DOMString localName);
  [CEReactions] boolean toggleAttribute(DOMString qualifiedName, optional boolean force);
  boolean hasAttribute(DOMString qualifiedName);
  boolean hasAttributeNS(DOMString? namespace, DOMString localName);

  Attr? getAttributeNode(DOMString qualifiedName);
  Attr? getAttributeNodeNS(DOMString? namespace, DOMString localName);
  [CEReactions] Attr? setAttributeNode(Attr attr);
  [CEReactions] Attr? setAttributeNodeNS(Attr attr);
  [CEReactions] Attr removeAttributeNode(Attr attr);

  ShadowRoot attachShadow(ShadowRootInit init);
  readonly attribute ShadowRoot? shadowRoot;

  readonly attribute CustomElementRegistry? customElementRegistry;

  Element? closest(DOMString selectors);
  boolean matches(DOMString selectors);
  boolean webkitMatchesSelector(DOMString selectors); // legacy alias of .matches

  HTMLCollection getElementsByTagName(DOMString qualifiedName);
  HTMLCollection getElementsByTagNameNS(DOMString? namespace, DOMString localName);
  HTMLCollection getElementsByClassName(DOMString classNames);

  [CEReactions] Element? insertAdjacentElement(DOMString where, Element element); // legacy
  undefined insertAdjacentText(DOMString where, DOMString data); // legacy
};

dictionary ShadowRootInit {
  required ShadowRootMode mode;
  boolean delegatesFocus = false;
  SlotAssignmentMode slotAssignment = "named";
  boolean clonable = false;
  boolean serializable = false;
  CustomElementRegistry customElementRegistry;
};
```

[`Element`{.idl}](#element){#ref-for-element⑤① link-type="idl"}
[nodes](#concept-node){#ref-for-concept-node①⑤⓪ link-type="dfn"} are
simply known as [elements]{#concept-element .dfn .dfn-paneled
dfn-type="dfn" export="" lt="element"}.

[Elements](#concept-element){#ref-for-concept-element⑦⑤ link-type="dfn"}
have an associated:

[namespace]{#concept-element-namespace .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}
:   Null or a non-empty string.

[namespace prefix]{#concept-element-namespace-prefix .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}
:   Null or a non-empty string.

[local name]{#concept-element-local-name .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}
:   A non-empty string.

[custom element registry]{#element-custom-element-registry .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}
:   Null or a
    [`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry①③
    link-type="idl"} object.

[custom element state]{#concept-element-custom-element-state .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}
:   \"`undefined`\", \"`failed`\", \"`uncustomized`\",
    \"`precustomized`\", or \"`custom`\".

[custom element definition]{#concept-element-custom-element-definition .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}
:   Null or a [custom element
    definition](https://html.spec.whatwg.org/multipage/custom-elements.html#custom-element-definition){#ref-for-custom-element-definition
    link-type="dfn"}.

[`is` value]{#concept-element-is-value .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}
:   Null or a [valid custom element
    name](https://html.spec.whatwg.org/multipage/custom-elements.html#valid-custom-element-name){#ref-for-valid-custom-element-name
    link-type="dfn"}.

When an [element](#concept-element){#ref-for-concept-element⑦⑥
link-type="dfn"} is
[created](#concept-create-element){#ref-for-concept-create-element⑦
link-type="dfn"}, all of these values are initialized.

An [element](#concept-element){#ref-for-concept-element⑦⑦
link-type="dfn"} whose [custom element
state](#concept-element-custom-element-state){#ref-for-concept-element-custom-element-state
link-type="dfn"} is \"`uncustomized`\" or \"`custom`\" is said to be
[defined]{#concept-element-defined .dfn .dfn-paneled dfn-for="Element"
dfn-type="dfn" export=""}. An
[element](#concept-element){#ref-for-concept-element⑦⑧ link-type="dfn"}
whose [custom element
state](#concept-element-custom-element-state){#ref-for-concept-element-custom-element-state①
link-type="dfn"} is \"`custom`\" is said to be
[custom]{#concept-element-custom .dfn .dfn-paneled dfn-for="Element"
dfn-type="dfn" export=""}.

Whether or not an element is
[defined](#concept-element-defined){#ref-for-concept-element-defined
link-type="dfn"} is used to determine the behavior of the
[:defined](https://drafts.csswg.org/selectors-4/#defined-pseudo){#ref-for-defined-pseudo
.css link-type="maybe"} pseudo-class. Whether or not an element is
[custom](#concept-element-custom){#ref-for-concept-element-custom⑥
link-type="dfn"} is used to determine the behavior of the [mutation
algorithms](#mutation-algorithms). The \"`failed`\" and
\"`precustomized`\" states are used to ensure that if a [custom element
constructor](https://html.spec.whatwg.org/multipage/custom-elements.html#custom-element-constructor){#ref-for-custom-element-constructor
link-type="dfn"} fails to execute correctly the first time, it is not
executed again by an
[upgrade](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-upgrade-an-element){#ref-for-concept-upgrade-an-element①
link-type="dfn"}.

::: {#example-c5b21302 .example}
[](#example-c5b21302){.self-link}

The following code illustrates elements in each of these four states:

``` {.lang-html .highlight}
<!DOCTYPE html>
<script>
  window.customElements.define("sw-rey", class extends HTMLElement {})
  window.customElements.define("sw-finn", class extends HTMLElement {}, { extends: "p" })
  window.customElements.define("sw-kylo", class extends HTMLElement {
    constructor() {
      // super() intentionally omitted for this example
    }
  })
</script>

<!-- "undefined" (not defined, not custom) -->
<sw-han></sw-han>
<p is="sw-luke"></p>
<p is="asdf"></p>

<!-- "failed" (not defined, not custom) -->
<sw-kylo></sw-kylo>

<!-- "uncustomized" (defined, not custom) -->
<p></p>
<asdf></asdf>

<!-- "custom" (defined, custom) -->
<sw-rey></sw-rey>
<p is="sw-finn"></p>
```
:::

[Elements](#concept-element){#ref-for-concept-element⑦⑨ link-type="dfn"}
also have an associated [shadow root]{#concept-element-shadow-root .dfn
.dfn-paneled dfn-for="Element" dfn-type="dfn" export=""} (null or a
[shadow root](#concept-shadow-root){#ref-for-concept-shadow-root④⑤
link-type="dfn"}). It is null unless otherwise stated. An
[element](#concept-element){#ref-for-concept-element⑧⓪ link-type="dfn"}
is a [shadow host]{#element-shadow-host .dfn .dfn-paneled
dfn-for="Element" dfn-type="dfn" export=""} if its [shadow
root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①②
link-type="dfn"} is non-null.

An [element](#concept-element){#ref-for-concept-element⑧①
link-type="dfn"}'s [qualified name]{#concept-element-qualified-name .dfn
.dfn-paneled dfn-for="Element" dfn-type="dfn" export=""} is its [local
name](#concept-element-local-name){#ref-for-concept-element-local-name⑨
link-type="dfn"} if its [namespace
prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix①⓪
link-type="dfn"} is null; otherwise its [namespace
prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix①①
link-type="dfn"}, followed by \"`:`\", followed by its [local
name](#concept-element-local-name){#ref-for-concept-element-local-name①⓪
link-type="dfn"}.

An [element](#concept-element){#ref-for-concept-element⑧②
link-type="dfn"}'s [HTML-uppercased qualified
name]{#element-html-uppercased-qualified-name .dfn .dfn-paneled
dfn-for="Element" dfn-type="dfn" noexport=""} is the return value of
these steps:

1.  Let `qualifiedName`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥⑧
    link-type="dfn"}'s [qualified
    name](#concept-element-qualified-name){#ref-for-concept-element-qualified-name④
    link-type="dfn"}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑥⑨
    link-type="dfn"} is in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①⑥
    link-type="dfn"} and its [node
    document](#concept-node-document){#ref-for-concept-node-document④⑦
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document①① link-type="dfn"},
    then set `qualifiedName`{.variable} to `qualifiedName`{.variable} in
    [ASCII
    uppercase](https://infra.spec.whatwg.org/#ascii-uppercase){#ref-for-ascii-uppercase
    link-type="dfn"}.

3.  Return `qualifiedName`{.variable}.

User agents could optimize [qualified
name](#concept-element-qualified-name){#ref-for-concept-element-qualified-name⑤
link-type="dfn"} and [HTML-uppercased qualified
name](#element-html-uppercased-qualified-name){#ref-for-element-html-uppercased-qualified-name②
link-type="dfn"} by storing them in internal slots.

To [create an element]{#concept-create-element .dfn .dfn-paneled
dfn-type="dfn" export=""}, given a
[document](#concept-document){#ref-for-concept-document⑤③
link-type="dfn"} `document`{.variable}, string `localName`{.variable},
string-or-null `namespace`{.variable}, and optionally a string-or-null
`prefix`{.variable} (default null), string-or-null `is`{.variable}
(default null), boolean `synchronousCustomElements`{.variable} (default
false), and \"`default`\", null, or a
[`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry①④
link-type="idl"} object `registry`{.variable} (default \"`default`\"):

1.  Let `result`{.variable} be null.

2.  If `registry`{.variable} is \"`default`\", then set
    `registry`{.variable} to the result of [looking up a custom element
    registry](https://html.spec.whatwg.org/multipage/custom-elements.html#look-up-a-custom-element-registry){#ref-for-look-up-a-custom-element-registry④
    link-type="dfn"} given `document`{.variable}.

3.  Let `definition`{.variable} be the result of [looking up a custom
    element
    definition](https://html.spec.whatwg.org/multipage/custom-elements.html#look-up-a-custom-element-definition){#ref-for-look-up-a-custom-element-definition
    link-type="dfn"} given `registry`{.variable},
    `namespace`{.variable}, `localName`{.variable}, and `is`{.variable}.

4.  If `definition`{.variable} is non-null, and
    `definition`{.variable}'s
    [name](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-custom-element-definition-name){#ref-for-concept-custom-element-definition-name
    link-type="dfn"} is not equal to its [local
    name](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-custom-element-definition-local-name){#ref-for-concept-custom-element-definition-local-name
    link-type="dfn"} (i.e., `definition`{.variable} represents a
    [customized built-in
    element](https://html.spec.whatwg.org/multipage/custom-elements.html#customized-built-in-element){#ref-for-customized-built-in-element②
    link-type="dfn"}):

    1.  Let `interface`{.variable} be the [element
        interface](#concept-element-interface){#ref-for-concept-element-interface
        link-type="dfn"} for `localName`{.variable} and the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①⑦
        link-type="dfn"}.

    2.  Set `result`{.variable} to the result of [creating an element
        internal](#create-an-element-internal){#ref-for-create-an-element-internal
        link-type="dfn"} given `document`{.variable},
        `interface`{.variable}, `localName`{.variable}, the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①⑧
        link-type="dfn"}, `prefix`{.variable}, \"`undefined`\",
        `is`{.variable}, and `registry`{.variable}.

    3.  If `synchronousCustomElements`{.variable} is true, then run this
        step while catching any exceptions:

        1.  [Upgrade](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-upgrade-an-element){#ref-for-concept-upgrade-an-element②
            link-type="dfn"} `result`{.variable} using
            `definition`{.variable}.

        If this step threw an exception `exception`{.variable}:

        1.  [Report](https://html.spec.whatwg.org/multipage/webappapis.html#report-an-exception){#ref-for-report-an-exception①
            link-type="dfn"} `exception`{.variable} for
            `definition`{.variable}'s
            [constructor](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-custom-element-definition-constructor){#ref-for-concept-custom-element-definition-constructor
            link-type="dfn"}'s corresponding JavaScript object's
            [associated
            realm](https://webidl.spec.whatwg.org/#dfn-associated-realm){#ref-for-dfn-associated-realm②
            link-type="dfn"}'s [global
            object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-realm-global){#ref-for-concept-realm-global②
            link-type="dfn"}.

        2.  Set `result`{.variable}'s [custom element
            state](#concept-element-custom-element-state){#ref-for-concept-element-custom-element-state②
            link-type="dfn"} to \"`failed`\".

    4.  Otherwise, [enqueue a custom element upgrade
        reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-upgrade-reaction){#ref-for-enqueue-a-custom-element-upgrade-reaction
        link-type="dfn"} given `result`{.variable} and
        `definition`{.variable}.

5.  Otherwise, if `definition`{.variable} is non-null:

    1.  If `synchronousCustomElements`{.variable} is true:

        1.  Let `C`{.variable} be `definition`{.variable}'s
            [constructor](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-custom-element-definition-constructor){#ref-for-concept-custom-element-definition-constructor①
            link-type="dfn"}.

        2.  [Set](https://infra.spec.whatwg.org/#map-set){#ref-for-map-set②
            link-type="dfn"} the [surrounding
            agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent⑧
            link-type="dfn"}'s [active custom element constructor
            map](https://html.spec.whatwg.org/multipage/custom-elements.html#active-custom-element-constructor-map){#ref-for-active-custom-element-constructor-map
            link-type="dfn"}\[`C`{.variable}\] to `registry`{.variable}.

        3.  Run these steps while catching any exceptions:

            1.  Set `result`{.variable} to the result of
                [constructing](https://webidl.spec.whatwg.org/#construct-a-callback-function){#ref-for-construct-a-callback-function
                link-type="dfn"} `C`{.variable}, with no arguments.

            2.  Assert: `result`{.variable}'s [custom element
                state](#concept-element-custom-element-state){#ref-for-concept-element-custom-element-state③
                link-type="dfn"} and [custom element
                definition](#concept-element-custom-element-definition){#ref-for-concept-element-custom-element-definition
                link-type="dfn"} are initialized.

            3.  Assert: `result`{.variable}'s
                [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①④
                link-type="dfn"} is the [HTML
                namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace①⑨
                link-type="dfn"}.

                IDL enforces that `result`{.variable} is an
                [`HTMLElement`{.idl}](https://html.spec.whatwg.org/multipage/dom.html#htmlelement){#ref-for-htmlelement
                link-type="idl"} object, which all use the [HTML
                namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②⓪
                link-type="dfn"}.

            4.  If `result`{.variable}'s [attribute
                list](#concept-element-attribute){#ref-for-concept-element-attribute⑥
                link-type="dfn"} [is not
                empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty⑥
                link-type="dfn"}, then
                [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤⓪
                link-type="dfn"} a
                \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①①
                .idl-code link-type="exception"}\"
                [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥④
                link-type="idl"}.

            5.  If `result`{.variable} has
                [children](#concept-tree-child){#ref-for-concept-tree-child⑥①
                link-type="dfn"}, then
                [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤①
                link-type="dfn"} a
                \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①②
                .idl-code link-type="exception"}\"
                [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥⑤
                link-type="idl"}.

            6.  If `result`{.variable}'s
                [parent](#concept-tree-parent){#ref-for-concept-tree-parent③⓪
                link-type="dfn"} is not null, then
                [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤②
                link-type="dfn"} a
                \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①③
                .idl-code link-type="exception"}\"
                [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥⑥
                link-type="idl"}.

            7.  If `result`{.variable}'s [node
                document](#concept-node-document){#ref-for-concept-node-document④⑧
                link-type="dfn"} is not `document`{.variable}, then
                [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤③
                link-type="dfn"} a
                \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①④
                .idl-code link-type="exception"}\"
                [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥⑦
                link-type="idl"}.

            8.  If `result`{.variable}'s [local
                name](#concept-element-local-name){#ref-for-concept-element-local-name①①
                link-type="dfn"} is not equal to `localName`{.variable},
                then
                [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤④
                link-type="dfn"} a
                \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①⑤
                .idl-code link-type="exception"}\"
                [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥⑧
                link-type="idl"}.

            9.  Set `result`{.variable}'s [namespace
                prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix①②
                link-type="dfn"} to `prefix`{.variable}.

            10. Set `result`{.variable}'s [`is`
                value](#concept-element-is-value){#ref-for-concept-element-is-value①
                link-type="dfn"} to null.

            11. Set `result`{.variable}'s [custom element
                registry](#element-custom-element-registry){#ref-for-element-custom-element-registry⑤
                link-type="dfn"} to `registry`{.variable}.

            If any of these steps threw an exception
            `exception`{.variable}:

            1.  [Report](https://html.spec.whatwg.org/multipage/webappapis.html#report-an-exception){#ref-for-report-an-exception②
                link-type="dfn"} `exception`{.variable} for
                `definition`{.variable}'s
                [constructor](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-custom-element-definition-constructor){#ref-for-concept-custom-element-definition-constructor②
                link-type="dfn"}'s corresponding JavaScript object's
                [associated
                realm](https://webidl.spec.whatwg.org/#dfn-associated-realm){#ref-for-dfn-associated-realm③
                link-type="dfn"}'s [global
                object](https://html.spec.whatwg.org/multipage/webappapis.html#concept-realm-global){#ref-for-concept-realm-global③
                link-type="dfn"}.

            2.  Set `result`{.variable} to the result of [creating an
                element
                internal](#create-an-element-internal){#ref-for-create-an-element-internal①
                link-type="dfn"} given `document`{.variable},
                [`HTMLUnknownElement`{.idl}](https://html.spec.whatwg.org/multipage/dom.html#htmlunknownelement){#ref-for-htmlunknownelement
                link-type="idl"}, `localName`{.variable}, the [HTML
                namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②①
                link-type="dfn"}, `prefix`{.variable}, \"`failed`\",
                null, and `registry`{.variable}.

        4.  [Remove](https://infra.spec.whatwg.org/#map-remove){#ref-for-map-remove
            link-type="dfn"} the [surrounding
            agent](https://tc39.es/ecma262/#surrounding-agent){#ref-for-surrounding-agent⑨
            link-type="dfn"}'s [active custom element constructor
            map](https://html.spec.whatwg.org/multipage/custom-elements.html#active-custom-element-constructor-map){#ref-for-active-custom-element-constructor-map①
            link-type="dfn"}\[`C`{.variable}\].

            Under normal circumstances it will already have been removed
            at this point.

    2.  Otherwise:

        1.  Set `result`{.variable} to the result of [creating an
            element
            internal](#create-an-element-internal){#ref-for-create-an-element-internal②
            link-type="dfn"} given `document`{.variable},
            [`HTMLElement`{.idl}](https://html.spec.whatwg.org/multipage/dom.html#htmlelement){#ref-for-htmlelement①
            link-type="idl"}, `localName`{.variable}, the [HTML
            namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②②
            link-type="dfn"}, `prefix`{.variable}, \"`undefined`\",
            null, and `registry`{.variable}.

        2.  [Enqueue a custom element upgrade
            reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-upgrade-reaction){#ref-for-enqueue-a-custom-element-upgrade-reaction①
            link-type="dfn"} given `result`{.variable} and
            `definition`{.variable}.

6.  Otherwise:

    1.  Let `interface`{.variable} be the [element
        interface](#concept-element-interface){#ref-for-concept-element-interface①
        link-type="dfn"} for `localName`{.variable} and
        `namespace`{.variable}.

    2.  Set `result`{.variable} to the result of [creating an element
        internal](#create-an-element-internal){#ref-for-create-an-element-internal③
        link-type="dfn"} given `document`{.variable},
        `interface`{.variable}, `localName`{.variable},
        `namespace`{.variable}, `prefix`{.variable}, \"`uncustomized`\",
        `is`{.variable}, and `registry`{.variable}.

    3.  If `namespace`{.variable} is the [HTML
        namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②③
        link-type="dfn"}, and either `localName`{.variable} is a [valid
        custom element
        name](https://html.spec.whatwg.org/multipage/custom-elements.html#valid-custom-element-name){#ref-for-valid-custom-element-name①
        link-type="dfn"} or `is`{.variable} is non-null, then set
        `result`{.variable}'s [custom element
        state](#concept-element-custom-element-state){#ref-for-concept-element-custom-element-state④
        link-type="dfn"} to \"`undefined`\".

7.  Return `result`{.variable}.

To [create an element internal]{#create-an-element-internal .dfn
.dfn-paneled dfn-type="dfn" noexport=""} given a
[document](#concept-document){#ref-for-concept-document⑤④
link-type="dfn"} `document`{.variable}, an interface
`interface`{.variable} a string `localName`{.variable}, a string-or-null
`namespace`{.variable}, a string-or-null `prefix`{.variable}, a string
`state`{.variable}, a string-or-null `is`{.variable}, and null or a
[`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry①⑤
link-type="idl"} object `registry`{.variable}:

1.  Let `element`{.variable} be a new
    [element](#concept-element){#ref-for-concept-element⑧③
    link-type="dfn"} that implements `interface`{.variable}, with
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①⑤
    link-type="dfn"} set to `namespace`{.variable}, [namespace
    prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix①③
    link-type="dfn"} set to `prefix`{.variable}, [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name①②
    link-type="dfn"} set to `localName`{.variable}, [custom element
    registry](#element-custom-element-registry){#ref-for-element-custom-element-registry⑥
    link-type="dfn"} set to `registry`{.variable}, [custom element
    state](#concept-element-custom-element-state){#ref-for-concept-element-custom-element-state⑤
    link-type="dfn"} set to `state`{.variable}, [custom element
    definition](#concept-element-custom-element-definition){#ref-for-concept-element-custom-element-definition①
    link-type="dfn"} set to null, [`is`
    value](#concept-element-is-value){#ref-for-concept-element-is-value②
    link-type="dfn"} set to `is`{.variable}, and [node
    document](#concept-node-document){#ref-for-concept-node-document④⑨
    link-type="dfn"} set to `document`{.variable}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑥
    link-type="dfn"}: `element`{.variable}'s [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute⑦
    link-type="dfn"} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty⑦
    link-type="dfn"}.

3.  Return `element`{.variable}.

[Elements](#concept-element){#ref-for-concept-element⑧④ link-type="dfn"}
also have an [attribute list]{#concept-element-attribute .dfn
.dfn-paneled dfn-for="Element" dfn-type="dfn" export=""}, which is a
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①②
link-type="dfn"} exposed through a
[`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap①
link-type="idl"}. Unless explicitly given when an
[element](#concept-element){#ref-for-concept-element⑧⑤ link-type="dfn"}
is created, its [attribute
list](#concept-element-attribute){#ref-for-concept-element-attribute⑧
link-type="dfn"} [is
empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty⑧
link-type="dfn"}.

An [element](#concept-element){#ref-for-concept-element⑧⑥
link-type="dfn"} [has]{#concept-element-attribute-has .dfn .dfn-paneled
dfn-type="dfn" export="" lt="has an attribute"} an
[attribute](#concept-attribute){#ref-for-concept-attribute②④
link-type="dfn"} `A`{.variable} if its [attribute
list](#concept-element-attribute){#ref-for-concept-element-attribute⑨
link-type="dfn"}
[contains](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain⑤
link-type="dfn"} `A`{.variable}.

This and [other
specifications](#other-applicable-specifications){#ref-for-other-applicable-specifications⑨
link-type="dfn"} may define [attribute change
steps]{#concept-element-attributes-change-ext .dfn .dfn-paneled
dfn-type="dfn" export=""} for
[elements](#concept-element){#ref-for-concept-element⑧⑦
link-type="dfn"}. The algorithm is passed `element`{.variable},
`localName`{.variable}, `oldValue`{.variable}, `value`{.variable}, and
`namespace`{.variable}.

To [handle attribute changes]{#handle-attribute-changes .dfn
.dfn-paneled dfn-type="dfn" noexport=""} for an
[attribute](#concept-attribute){#ref-for-concept-attribute②⑤
link-type="dfn"} `attribute`{.variable} with `element`{.variable},
`oldValue`{.variable}, and `newValue`{.variable}, run these steps:

1.  [Queue a mutation
    record](#queue-a-mutation-record){#ref-for-queue-a-mutation-record①
    link-type="dfn"} of \"`attributes`\" for `element`{.variable} with
    `attribute`{.variable}'s [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name⑨
    link-type="dfn"}, `attribute`{.variable}'s
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace⑦
    link-type="dfn"}, `oldValue`{.variable}, « », « », null, and null.

2.  If `element`{.variable} is
    [custom](#concept-element-custom){#ref-for-concept-element-custom⑦
    link-type="dfn"}, then [enqueue a custom element callback
    reaction](https://html.spec.whatwg.org/multipage/custom-elements.html#enqueue-a-custom-element-callback-reaction){#ref-for-enqueue-a-custom-element-callback-reaction⑤
    link-type="dfn"} with `element`{.variable}, callback name
    \"`attributeChangedCallback`\", and « `attribute`{.variable}'s
    [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①⓪
    link-type="dfn"}, `oldValue`{.variable}, `newValue`{.variable},
    `attribute`{.variable}'s
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace⑧
    link-type="dfn"} ».

3.  Run the [attribute change
    steps](#concept-element-attributes-change-ext){#ref-for-concept-element-attributes-change-ext②
    link-type="dfn"} with `element`{.variable}, `attribute`{.variable}'s
    [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①①
    link-type="dfn"}, `oldValue`{.variable}, `newValue`{.variable}, and
    `attribute`{.variable}'s
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace⑨
    link-type="dfn"}.

To [change]{#concept-element-attributes-change .dfn .dfn-paneled
dfn-type="dfn" export="" lt="change an attribute"} an
[attribute](#concept-attribute){#ref-for-concept-attribute②⑥
link-type="dfn"} `attribute`{.variable} to `value`{.variable}, run these
steps:

1.  Let `oldValue`{.variable} be `attribute`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①①
    link-type="dfn"}.

2.  Set `attribute`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①②
    link-type="dfn"} to `value`{.variable}.

3.  [Handle attribute
    changes](#handle-attribute-changes){#ref-for-handle-attribute-changes
    link-type="dfn"} for `attribute`{.variable} with
    `attribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element⑥
    link-type="dfn"}, `oldValue`{.variable}, and `value`{.variable}.

To [append]{#concept-element-attributes-append .dfn .dfn-paneled
dfn-type="dfn" export="" lt="append an attribute"} an
[attribute](#concept-attribute){#ref-for-concept-attribute②⑦
link-type="dfn"} `attribute`{.variable} to an
[element](#concept-element){#ref-for-concept-element⑧⑧ link-type="dfn"}
`element`{.variable}, run these steps:

1.  [Append](https://infra.spec.whatwg.org/#list-append){#ref-for-list-append①④
    link-type="dfn"} `attribute`{.variable} to `element`{.variable}'s
    [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①⓪
    link-type="dfn"}.

2.  Set `attribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element⑦
    link-type="dfn"} to `element`{.variable}.

3.  Set `attribute`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑤⓪
    link-type="dfn"} to `element`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑤①
    link-type="dfn"}.

4.  [Handle attribute
    changes](#handle-attribute-changes){#ref-for-handle-attribute-changes①
    link-type="dfn"} for `attribute`{.variable} with
    `element`{.variable}, null, and `attribute`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①③
    link-type="dfn"}.

To [remove]{#concept-element-attributes-remove .dfn .dfn-paneled
dfn-type="dfn" export="" lt="remove an attribute"} an
[attribute](#concept-attribute){#ref-for-concept-attribute②⑧
link-type="dfn"} `attribute`{.variable}, run these steps:

1.  Let `element`{.variable} be `attribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element⑧
    link-type="dfn"}.

2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#ref-for-list-remove⑦
    link-type="dfn"} `attribute`{.variable} from `element`{.variable}'s
    [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①①
    link-type="dfn"}.

3.  Set `attribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element⑨
    link-type="dfn"} to null.

4.  [Handle attribute
    changes](#handle-attribute-changes){#ref-for-handle-attribute-changes②
    link-type="dfn"} for `attribute`{.variable} with
    `element`{.variable}, `attribute`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①④
    link-type="dfn"}, and null.

To [replace]{#concept-element-attributes-replace .dfn .dfn-paneled
dfn-type="dfn" export="" lt="replace an attribute"} an
[attribute](#concept-attribute){#ref-for-concept-attribute②⑨
link-type="dfn"} `oldAttribute`{.variable} with an
[attribute](#concept-attribute){#ref-for-concept-attribute③⓪
link-type="dfn"} `newAttribute`{.variable}:

1.  Let `element`{.variable} be `oldAttribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element①⓪
    link-type="dfn"}.

2.  [Replace](https://infra.spec.whatwg.org/#list-replace){#ref-for-list-replace
    link-type="dfn"} `oldAttribute`{.variable} by
    `newAttribute`{.variable} in `element`{.variable}'s [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①②
    link-type="dfn"}.

3.  Set `newAttribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element①①
    link-type="dfn"} to `element`{.variable}.

4.  Set `newAttribute`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑤②
    link-type="dfn"} to `element`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑤③
    link-type="dfn"}.

5.  Set `oldAttribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element①②
    link-type="dfn"} to null.

6.  [Handle attribute
    changes](#handle-attribute-changes){#ref-for-handle-attribute-changes③
    link-type="dfn"} for `oldAttribute`{.variable} with
    `element`{.variable}, `oldAttribute`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①⑤
    link-type="dfn"}, and `newAttribute`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①⑥
    link-type="dfn"}.

------------------------------------------------------------------------

::: {.algorithm algorithm="get an attribute by name"}
To [get an attribute by name]{#concept-element-attributes-get-by-name
.dfn .dfn-paneled dfn-type="dfn" export=""} given a string
`qualifiedName`{.variable} and an
[element](#concept-element){#ref-for-concept-element⑧⑨ link-type="dfn"}
`element`{.variable}:

1.  If `element`{.variable} is in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②④
    link-type="dfn"} and its [node
    document](#concept-node-document){#ref-for-concept-node-document⑤④
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document①② link-type="dfn"},
    then set `qualifiedName`{.variable} to `qualifiedName`{.variable} in
    [ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase③
    link-type="dfn"}.

2.  Return the first
    [attribute](#concept-attribute){#ref-for-concept-attribute③①
    link-type="dfn"} in `element`{.variable}'s [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①③
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name②
    link-type="dfn"} is `qualifiedName`{.variable}; otherwise null.
:::

::: {.algorithm algorithm="get an attribute by namespace and local name"}
To [get an attribute by namespace and local
name]{#concept-element-attributes-get-by-namespace .dfn .dfn-paneled
dfn-type="dfn" export=""} given null or a string `namespace`{.variable},
a string `localName`{.variable}, and an
[element](#concept-element){#ref-for-concept-element⑨⓪ link-type="dfn"}
`element`{.variable}:

1.  If `namespace`{.variable} is the empty string, then set it to null.

2.  Return the
    [attribute](#concept-attribute){#ref-for-concept-attribute③②
    link-type="dfn"} in `element`{.variable}'s [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①④
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①⓪
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①②
    link-type="dfn"} is `localName`{.variable}, if any; otherwise null.
:::

::: {.algorithm algorithm="get an attribute value"}
To [get an attribute value]{#concept-element-attributes-get-value .dfn
.dfn-paneled dfn-type="dfn" export=""} given an
[element](#concept-element){#ref-for-concept-element⑨① link-type="dfn"}
`element`{.variable}, a string `localName`{.variable}, and an optional
null or string `namespace`{.variable} (default null):

1.  Let `attr`{.variable} be the result of [getting an
    attribute](#concept-element-attributes-get-by-namespace){#ref-for-concept-element-attributes-get-by-namespace
    link-type="dfn"} given `namespace`{.variable},
    `localName`{.variable}, and `element`{.variable}.

2.  If `attr`{.variable} is null, then return the empty string.

3.  Return `attr`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①⑦
    link-type="dfn"}.
:::

::: {.algorithm algorithm="set an attribute"}
To [set an attribute]{#concept-element-attributes-set .dfn .dfn-paneled
dfn-type="dfn" export=""} given an
[attribute](#concept-attribute){#ref-for-concept-attribute③③
link-type="dfn"} `attr`{.variable} and an
[element](#concept-element){#ref-for-concept-element⑨② link-type="dfn"}
`element`{.variable}:

1.  If `attr`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element①③
    link-type="dfn"} is neither null nor `element`{.variable},
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤⑤
    link-type="dfn"} an
    \"[`InUseAttributeError`{.idl}](https://webidl.spec.whatwg.org/#inuseattributeerror){#ref-for-inuseattributeerror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑥⑨
    link-type="idl"}.

2.  Let `oldAttr`{.variable} be the result of [getting an
    attribute](#concept-element-attributes-get-by-namespace){#ref-for-concept-element-attributes-get-by-namespace①
    link-type="dfn"} given `attr`{.variable}'s
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①①
    link-type="dfn"}, `attr`{.variable}'s [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①③
    link-type="dfn"}, and `element`{.variable}.

3.  If `oldAttr`{.variable} is `attr`{.variable}, return
    `attr`{.variable}.

4.  If `oldAttr`{.variable} is non-null, then
    [replace](#concept-element-attributes-replace){#ref-for-concept-element-attributes-replace
    link-type="dfn"} `oldAttr`{.variable} with `attr`{.variable}.

5.  Otherwise,
    [append](#concept-element-attributes-append){#ref-for-concept-element-attributes-append①
    link-type="dfn"} `attr`{.variable} to `element`{.variable}.

6.  Return `oldAttr`{.variable}.
:::

::: {.algorithm algorithm="set an attribute value"}
To [set an attribute value]{#concept-element-attributes-set-value .dfn
.dfn-paneled dfn-type="dfn" export=""} given an
[element](#concept-element){#ref-for-concept-element⑨③ link-type="dfn"}
`element`{.variable}, a string `localName`{.variable}, a string
`value`{.variable}, an optional null or string `prefix`{.variable}
(default null), and an optional null or string `namespace`{.variable}
(default null):

1.  Let `attribute`{.variable} be the result of [getting an
    attribute](#concept-element-attributes-get-by-namespace){#ref-for-concept-element-attributes-get-by-namespace②
    link-type="dfn"} given `namespace`{.variable},
    `localName`{.variable}, and `element`{.variable}.

2.  If `attribute`{.variable} is null, create an
    [attribute](#concept-attribute){#ref-for-concept-attribute③④
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①②
    link-type="dfn"} is `namespace`{.variable}, [namespace
    prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix⑤
    link-type="dfn"} is `prefix`{.variable}, [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①④
    link-type="dfn"} is `localName`{.variable},
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①⑧
    link-type="dfn"} is `value`{.variable}, and [node
    document](#concept-node-document){#ref-for-concept-node-document⑤⑤
    link-type="dfn"} is `element`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑤⑥
    link-type="dfn"}, then
    [append](#concept-element-attributes-append){#ref-for-concept-element-attributes-append②
    link-type="dfn"} this
    [attribute](#concept-attribute){#ref-for-concept-attribute③⑤
    link-type="dfn"} to `element`{.variable}, and then return.

3.  [Change](#concept-element-attributes-change){#ref-for-concept-element-attributes-change
    link-type="dfn"} `attribute`{.variable} to `value`{.variable}.
:::

::: {.algorithm algorithm="remove an attribute by name"}
To [remove an attribute by
name]{#concept-element-attributes-remove-by-name .dfn .dfn-paneled
dfn-type="dfn" export=""} given a string `qualifiedName`{.variable} and
an [element](#concept-element){#ref-for-concept-element⑨④
link-type="dfn"} `element`{.variable}:

1.  Let `attr`{.variable} be the result of [getting an
    attribute](#concept-element-attributes-get-by-name){#ref-for-concept-element-attributes-get-by-name
    link-type="dfn"} given `qualifiedName`{.variable} and
    `element`{.variable}.

2.  If `attr`{.variable} is non-null, then
    [remove](#concept-element-attributes-remove){#ref-for-concept-element-attributes-remove
    link-type="dfn"} `attr`{.variable}.

3.  Return `attr`{.variable}.
:::

::: {.algorithm algorithm="remove an attribute by namespace and local name"}
To [remove an attribute by namespace and local
name]{#concept-element-attributes-remove-by-namespace .dfn .dfn-paneled
dfn-type="dfn" export=""} given null or a string `namespace`{.variable},
a string `localName`{.variable}, and an
[element](#concept-element){#ref-for-concept-element⑨⑤ link-type="dfn"}
`element`{.variable}:

1.  Let `attr`{.variable} be the result of [getting an
    attribute](#concept-element-attributes-get-by-namespace){#ref-for-concept-element-attributes-get-by-namespace③
    link-type="dfn"} given `namespace`{.variable},
    `localName`{.variable}, and `element`{.variable}.

2.  If `attr`{.variable} is non-null, then
    [remove](#concept-element-attributes-remove){#ref-for-concept-element-attributes-remove①
    link-type="dfn"} `attr`{.variable}.

3.  Return `attr`{.variable}.
:::

------------------------------------------------------------------------

An [element](#concept-element){#ref-for-concept-element⑨⑥
link-type="dfn"} can have an associated [unique identifier
(ID)]{#concept-id .dfn .dfn-paneled dfn-for="Element" dfn-type="dfn"
export="" lt="ID"}

Historically [elements](#concept-element){#ref-for-concept-element⑨⑦
link-type="dfn"} could have multiple identifiers e.g., by using the HTML
`id` [attribute](#concept-attribute){#ref-for-concept-attribute③⑥
link-type="dfn"} and a DTD. This specification makes
[ID](#concept-id){#ref-for-concept-id⑥ link-type="dfn"} a concept of the
DOM and allows for only one per
[element](#concept-element){#ref-for-concept-element⑨⑧ link-type="dfn"},
given by an [`id`
attribute](#concept-named-attribute){#ref-for-concept-named-attribute③
link-type="dfn"}.

Use these [attribute change
steps](#concept-element-attributes-change-ext){#ref-for-concept-element-attributes-change-ext③
link-type="dfn"} to update an
[element](#concept-element){#ref-for-concept-element⑨⑨
link-type="dfn"}'s [ID](#concept-id){#ref-for-concept-id⑦
link-type="dfn"}:

1.  If `localName`{.variable} is `id`, `namespace`{.variable} is null,
    and `value`{.variable} is null or the empty string, then unset
    `element`{.variable}'s [ID](#concept-id){#ref-for-concept-id⑧
    link-type="dfn"}.

2.  Otherwise, if `localName`{.variable} is `id`, `namespace`{.variable}
    is null, then set `element`{.variable}'s
    [ID](#concept-id){#ref-for-concept-id⑨ link-type="dfn"} to
    `value`{.variable}.

While this specification defines requirements for `class`, `id`, and
`slot` [attributes](#concept-attribute){#ref-for-concept-attribute③⑦
link-type="dfn"} on any
[element](#concept-element){#ref-for-concept-element①⓪⓪
link-type="dfn"}, it makes no claims as to whether using them is
conforming or not.

------------------------------------------------------------------------

A [node](#concept-node){#ref-for-concept-node①⑤① link-type="dfn"}'s
[parent](#concept-tree-parent){#ref-for-concept-tree-parent③①
link-type="dfn"} of type [`Element`{.idl}](#element){#ref-for-element⑤②
link-type="idl"} is known as its [parent element]{#parent-element .dfn
.dfn-paneled dfn-type="dfn" export=""}. If the
[node](#concept-node){#ref-for-concept-node①⑤② link-type="dfn"} has a
[parent](#concept-tree-parent){#ref-for-concept-tree-parent③②
link-type="dfn"} of a different type, its [parent
element](#parent-element){#ref-for-parent-element⑨ link-type="dfn"} is
null.

------------------------------------------------------------------------

`namespace`{.variable} = `element`{.variable} . [`namespaceURI`{.idl}](#dom-element-namespaceuri){#ref-for-dom-element-namespaceuri① link-type="idl"}
:   Returns the
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①⑥
    link-type="dfn"}.

`prefix`{.variable} = `element`{.variable} . [`prefix`{.idl}](#dom-element-prefix){#ref-for-dom-element-prefix① link-type="idl"}
:   Returns the [namespace
    prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix①④
    link-type="dfn"}.

`localName`{.variable} = `element`{.variable} . [`localName`{.idl}](#dom-element-localname){#ref-for-dom-element-localname① link-type="idl"}
:   Returns the [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name①③
    link-type="dfn"}.

`qualifiedName`{.variable} = `element`{.variable} . [`tagName`{.idl}](#dom-element-tagname){#ref-for-dom-element-tagname① link-type="idl"}
:   Returns the [HTML-uppercased qualified
    name](#element-html-uppercased-qualified-name){#ref-for-element-html-uppercased-qualified-name③
    link-type="dfn"}.

The [`namespaceURI`]{#dom-element-namespaceuri .dfn .dfn-paneled
.idl-code dfn-for="Element" dfn-type="attribute" export=""} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦⓪
link-type="dfn"}'s
[namespace](#concept-element-namespace){#ref-for-concept-element-namespace①⑦
link-type="dfn"}.

The [`prefix`]{#dom-element-prefix .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦①
link-type="dfn"}'s [namespace
prefix](#concept-element-namespace-prefix){#ref-for-concept-element-namespace-prefix①⑤
link-type="dfn"}.

The [`localName`]{#dom-element-localname .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦②
link-type="dfn"}'s [local
name](#concept-element-local-name){#ref-for-concept-element-local-name①④
link-type="dfn"}.

The [`tagName`]{#dom-element-tagname .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦③
link-type="dfn"}'s [HTML-uppercased qualified
name](#element-html-uppercased-qualified-name){#ref-for-element-html-uppercased-qualified-name④
link-type="dfn"}.

------------------------------------------------------------------------

`element`{.variable}` . `[`id`](#dom-element-id){#ref-for-dom-element-id① .idl-code link-type="attribute"}` [ = ``value`{.variable}` ]`

:   Returns the value of `element`{.variable}'s `id` content attribute.
    Can be set to change it.

`element`{.variable}` . `[`className`](#dom-element-classname){#ref-for-dom-element-classname① .idl-code link-type="attribute"}` [ = ``value`{.variable}` ]`

:   Returns the value of `element`{.variable}'s `class` content
    attribute. Can be set to change it.

`element`{.variable}` . `[`classList`](#dom-element-classlist){#ref-for-dom-element-classlist① .idl-code link-type="attribute"}

:   Allows for manipulation of `element`{.variable}'s `class` content
    attribute as a set of whitespace-separated tokens through a
    [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist①
    link-type="idl"} object.

`element`{.variable}` . `[`slot`](#dom-element-slot){#ref-for-dom-element-slot① .idl-code link-type="attribute"}` [ = ``value`{.variable}` ]`

:   Returns the value of `element`{.variable}'s `slot` content
    attribute. Can be set to change it.

IDL attributes that are defined to [reflect]{#concept-reflect .dfn
.dfn-paneled dfn-for="Attr" dfn-type="dfn" noexport=""} a string
`name`{.variable}, must have these getter and setter steps:

getter steps

:   Return the result of running [get an attribute
    value](#concept-element-attributes-get-value){#ref-for-concept-element-attributes-get-value
    link-type="dfn"} given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦④
    link-type="dfn"} and `name`{.variable}.

setter steps

:   [Set an attribute
    value](#concept-element-attributes-set-value){#ref-for-concept-element-attributes-set-value
    link-type="dfn"} for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦⑤
    link-type="dfn"} using `name`{.variable} and the given value.

The [`id`]{#dom-element-id .dfn .dfn-paneled .idl-code dfn-for="Element"
dfn-type="attribute" export=""} attribute must
[reflect](#concept-reflect){#ref-for-concept-reflect link-type="dfn"}
\"`id`\".

The [`className`]{#dom-element-classname .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} attribute must
[reflect](#concept-reflect){#ref-for-concept-reflect① link-type="dfn"}
\"`class`\".

The [`classList`]{#dom-element-classlist .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} getter steps are to
return a [`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist②
link-type="idl"} object whose associated
[element](#concept-element){#ref-for-concept-element①⓪① link-type="dfn"}
is [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦⑥
link-type="dfn"} and whose associated
[attribute](#concept-attribute){#ref-for-concept-attribute③⑧
link-type="dfn"}'s [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①⑤
link-type="dfn"} is `class`. The [token
set](#concept-dtl-tokens){#ref-for-concept-dtl-tokens link-type="dfn"}
of this particular
[`DOMTokenList`{.idl}](#domtokenlist){#ref-for-domtokenlist③
link-type="idl"} object are also known as the
[element](#concept-element){#ref-for-concept-element①⓪②
link-type="dfn"}'s [classes]{#concept-class .dfn .dfn-paneled
dfn-for="Element" dfn-type="dfn" export="" lt="class"}.

The [`slot`]{#dom-element-slot .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} attribute must
[reflect](#concept-reflect){#ref-for-concept-reflect② link-type="dfn"}
\"`slot`\".

`id`, `class`, and `slot` are effectively superglobal attributes as they
can appear on any element, regardless of that element's namespace.

------------------------------------------------------------------------

`element`{.variable}` . `[`hasAttributes`](#dom-element-hasattributes){#ref-for-dom-element-hasattributes① .idl-code link-type="method"}`()`

:   Returns true if `element`{.variable} has attributes; otherwise
    false.

`element`{.variable}` . `[`getAttributeNames`](#dom-element-getattributenames){#ref-for-dom-element-getattributenames① .idl-code link-type="method"}`()`

:   Returns the [qualified
    names](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name③
    link-type="dfn"} of all `element`{.variable}'s
    [attributes](#concept-attribute){#ref-for-concept-attribute③⑨
    link-type="dfn"}. Can contain duplicates.

`element`{.variable}` . `[`getAttribute`](#dom-element-getattribute){#ref-for-dom-element-getattribute① .idl-code link-type="method"}`(``qualifiedName`{.variable}`)`

:   Returns `element`{.variable}'s first
    [attribute](#concept-attribute){#ref-for-concept-attribute④⓪
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name④
    link-type="dfn"} is `qualifiedName`{.variable}, and null if there is
    no such [attribute](#concept-attribute){#ref-for-concept-attribute④①
    link-type="dfn"} otherwise.

`element`{.variable}` . `[`getAttributeNS`](#dom-element-getattributens){#ref-for-dom-element-getattributens① .idl-code link-type="method"}`(``namespace`{.variable}`, ``localName`{.variable}`)`

:   Returns `element`{.variable}'s
    [attribute](#concept-attribute){#ref-for-concept-attribute④②
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①③
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①⑥
    link-type="dfn"} is `localName`{.variable}, and null if there is no
    such [attribute](#concept-attribute){#ref-for-concept-attribute④③
    link-type="dfn"} otherwise.

`element`{.variable}` . `[`setAttribute`](#dom-element-setattribute){#ref-for-dom-element-setattribute① .idl-code link-type="method"}`(``qualifiedName`{.variable}`, ``value`{.variable}`)`

:   Sets the
    [value](#concept-attribute-value){#ref-for-concept-attribute-value①⑨
    link-type="dfn"} of `element`{.variable}'s first
    [attribute](#concept-attribute){#ref-for-concept-attribute④④
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name⑤
    link-type="dfn"} is `qualifiedName`{.variable} to
    `value`{.variable}.

`element`{.variable}` . `[`setAttributeNS`](#dom-element-setattributens){#ref-for-dom-element-setattributens① .idl-code link-type="method"}`(``namespace`{.variable}`, ``localName`{.variable}`, ``value`{.variable}`)`

:   Sets the
    [value](#concept-attribute-value){#ref-for-concept-attribute-value②⓪
    link-type="dfn"} of `element`{.variable}'s
    [attribute](#concept-attribute){#ref-for-concept-attribute④⑤
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①④
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①⑦
    link-type="dfn"} is `localName`{.variable} to `value`{.variable}.

`element`{.variable}` . `[`removeAttribute`](#dom-element-removeattribute){#ref-for-dom-element-removeattribute① .idl-code link-type="method"}`(``qualifiedName`{.variable}`)`

:   Removes `element`{.variable}'s first
    [attribute](#concept-attribute){#ref-for-concept-attribute④⑥
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name⑥
    link-type="dfn"} is `qualifiedName`{.variable}.

`element`{.variable}` . `[`removeAttributeNS`](#dom-element-removeattributens){#ref-for-dom-element-removeattributens① .idl-code link-type="method"}`(``namespace`{.variable}`, ``localName`{.variable}`)`

:   Removes `element`{.variable}'s
    [attribute](#concept-attribute){#ref-for-concept-attribute④⑦
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①⑤
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①⑧
    link-type="dfn"} is `localName`{.variable}.

`element`{.variable}` . `[`toggleAttribute`](#dom-element-toggleattribute){#ref-for-dom-element-toggleattribute① .idl-code link-type="method"}`(``qualifiedName`{.variable}` [, ``force`{.variable}`])`

:   If `force`{.variable} is not given, \"toggles\"
    `qualifiedName`{.variable}, removing it if it is present and adding
    it if it is not present. If `force`{.variable} is true, adds
    `qualifiedName`{.variable}. If `force`{.variable} is false, removes
    `qualifiedName`{.variable}.

    Returns true if `qualifiedName`{.variable} is now present; otherwise
    false.

`element`{.variable}` . `[`hasAttribute`](#dom-element-hasattribute){#ref-for-dom-element-hasattribute① .idl-code link-type="method"}`(``qualifiedName`{.variable}`)`

:   Returns true if `element`{.variable} has an
    [attribute](#concept-attribute){#ref-for-concept-attribute④⑧
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name⑦
    link-type="dfn"} is `qualifiedName`{.variable}; otherwise false.

`element`{.variable}` . `[`hasAttributeNS`](#dom-element-hasattributens){#ref-for-dom-element-hasattributens① .idl-code link-type="method"}`(``namespace`{.variable}`, ``localName`{.variable}`)`

:   Returns true if `element`{.variable} has an
    [attribute](#concept-attribute){#ref-for-concept-attribute④⑨
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①⑥
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name①⑨
    link-type="dfn"} is `localName`{.variable}.

The [`hasAttributes()`]{#dom-element-hasattributes .dfn .dfn-paneled
.idl-code dfn-for="Element" dfn-type="method" export=""} method steps
are to return false if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦⑦
link-type="dfn"}'s [attribute
list](#concept-element-attribute){#ref-for-concept-element-attribute①⑤
link-type="dfn"} [is
empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty⑨
link-type="dfn"}; otherwise true.

The [`attributes`]{#dom-element-attributes .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} getter steps are to
return the associated
[`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap②
link-type="idl"}.

The [`getAttributeNames()`]{#dom-element-getattributenames .dfn
.dfn-paneled .idl-code dfn-for="Element" dfn-type="method" export=""}
method steps are to return the [qualified
names](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name⑧
link-type="dfn"} of the
[attributes](#concept-attribute){#ref-for-concept-attribute⑤⓪
link-type="dfn"} in
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦⑧
link-type="dfn"}'s [attribute
list](#concept-element-attribute){#ref-for-concept-element-attribute①⑥
link-type="dfn"}, in order; otherwise a new
[list](https://infra.spec.whatwg.org/#list){#ref-for-list①③
link-type="dfn"}.

These are not guaranteed to be unique.

The
[`getAttribute(``qualifiedName`{.variable}`)`]{#dom-element-getattribute
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  Let `attr`{.variable} be the result of [getting an
    attribute](#concept-element-attributes-get-by-name){#ref-for-concept-element-attributes-get-by-name①
    link-type="dfn"} given `qualifiedName`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑦⑨
    link-type="dfn"}.

2.  If `attr`{.variable} is null, return null.

3.  Return `attr`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value②①
    link-type="dfn"}.

The
[`getAttributeNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-element-getattributens
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  Let `attr`{.variable} be the result of [getting an
    attribute](#concept-element-attributes-get-by-namespace){#ref-for-concept-element-attributes-get-by-namespace④
    link-type="dfn"} given `namespace`{.variable},
    `localName`{.variable}, and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧⓪
    link-type="dfn"}.

2.  If `attr`{.variable} is null, return null.

3.  Return `attr`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value②②
    link-type="dfn"}.

The
[`setAttribute(``qualifiedName`{.variable}`, ``value`{.variable}`)`]{#dom-element-setattribute
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  If `qualifiedName`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name⑥ .css
    link-type="type"} production in XML, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤⑥
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦⓪
    link-type="idl"}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧①
    link-type="dfn"} is in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②⑤
    link-type="dfn"} and its [node
    document](#concept-node-document){#ref-for-concept-node-document⑤⑦
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document①③ link-type="dfn"},
    then set `qualifiedName`{.variable} to `qualifiedName`{.variable} in
    [ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase④
    link-type="dfn"}.

3.  Let `attribute`{.variable} be the first
    [attribute](#concept-attribute){#ref-for-concept-attribute⑤①
    link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧②
    link-type="dfn"}'s [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①⑦
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name⑨
    link-type="dfn"} is `qualifiedName`{.variable}, and null otherwise.

4.  If `attribute`{.variable} is null, create an
    [attribute](#concept-attribute){#ref-for-concept-attribute⑤②
    link-type="dfn"} whose [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②⓪
    link-type="dfn"} is `qualifiedName`{.variable},
    [value](#concept-attribute-value){#ref-for-concept-attribute-value②③
    link-type="dfn"} is `value`{.variable}, and [node
    document](#concept-node-document){#ref-for-concept-node-document⑤⑧
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧③
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑤⑨
    link-type="dfn"}, then
    [append](#concept-element-attributes-append){#ref-for-concept-element-attributes-append③
    link-type="dfn"} this
    [attribute](#concept-attribute){#ref-for-concept-attribute⑤③
    link-type="dfn"} to
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧④
    link-type="dfn"}, and then return.

5.  [Change](#concept-element-attributes-change){#ref-for-concept-element-attributes-change①
    link-type="dfn"} `attribute`{.variable} to `value`{.variable}.

The
[`setAttributeNS(``namespace`{.variable}`, ``qualifiedName`{.variable}`, ``value`{.variable}`)`]{#dom-element-setattributens
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  Let `namespace`{.variable}, `prefix`{.variable}, and
    `localName`{.variable} be the result of passing
    `namespace`{.variable} and `qualifiedName`{.variable} to [validate
    and extract](#validate-and-extract){#ref-for-validate-and-extract②
    link-type="dfn"}.

2.  [Set an attribute
    value](#concept-element-attributes-set-value){#ref-for-concept-element-attributes-set-value①
    link-type="dfn"} for
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧⑤
    link-type="dfn"} using `localName`{.variable}, `value`{.variable},
    and also `prefix`{.variable} and `namespace`{.variable}.

The
[`removeAttribute(``qualifiedName`{.variable}`)`]{#dom-element-removeattribute
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to [remove an
attribute](#concept-element-attributes-remove-by-name){#ref-for-concept-element-attributes-remove-by-name
link-type="dfn"} given `qualifiedName`{.variable} and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧⑥
link-type="dfn"}, and then return undefined.

The
[`removeAttributeNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-element-removeattributens
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to [remove an
attribute](#concept-element-attributes-remove-by-namespace){#ref-for-concept-element-attributes-remove-by-namespace
link-type="dfn"} given `namespace`{.variable}, `localName`{.variable},
and [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧⑦
link-type="dfn"}, and then return undefined.

The
[`hasAttribute(``qualifiedName`{.variable}`)`]{#dom-element-hasattribute
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧⑧
    link-type="dfn"} is in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②⑥
    link-type="dfn"} and its [node
    document](#concept-node-document){#ref-for-concept-node-document⑥⓪
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document①④ link-type="dfn"},
    then set `qualifiedName`{.variable} to `qualifiedName`{.variable} in
    [ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase⑤
    link-type="dfn"}.

2.  Return true if
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑧⑨
    link-type="dfn"}
    [has](#concept-element-attribute-has){#ref-for-concept-element-attribute-has⑤
    link-type="dfn"} an
    [attribute](#concept-attribute){#ref-for-concept-attribute⑤④
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name①⓪
    link-type="dfn"} is `qualifiedName`{.variable}; otherwise false.

The
[`toggleAttribute(``qualifiedName`{.variable}`, ``force`{.variable}`)`]{#dom-element-toggleattribute
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""
lt="toggleAttribute(qualifiedName, force)|toggleAttribute(qualifiedName)"}
method steps are:

1.  If `qualifiedName`{.variable} does not match the
    [`Name`](https://www.w3.org/TR/xml/#NT-Name){#ref-for-NT-Name⑦ .css
    link-type="type"} production in XML, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤⑦
    link-type="dfn"} an
    \"[`InvalidCharacterError`{.idl}](https://webidl.spec.whatwg.org/#invalidcharactererror){#ref-for-invalidcharactererror①②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦①
    link-type="idl"}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨⓪
    link-type="dfn"} is in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②⑦
    link-type="dfn"} and its [node
    document](#concept-node-document){#ref-for-concept-node-document⑥①
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document①⑤ link-type="dfn"},
    then set `qualifiedName`{.variable} to `qualifiedName`{.variable} in
    [ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase⑥
    link-type="dfn"}.

3.  Let `attribute`{.variable} be the first
    [attribute](#concept-attribute){#ref-for-concept-attribute⑤⑤
    link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨①
    link-type="dfn"}'s [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①⑧
    link-type="dfn"} whose [qualified
    name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name①①
    link-type="dfn"} is `qualifiedName`{.variable}, and null otherwise.

4.  If `attribute`{.variable} is null:

    1.  If `force`{.variable} is not given or is true, create an
        [attribute](#concept-attribute){#ref-for-concept-attribute⑤⑥
        link-type="dfn"} whose [local
        name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②①
        link-type="dfn"} is `qualifiedName`{.variable},
        [value](#concept-attribute-value){#ref-for-concept-attribute-value②④
        link-type="dfn"} is the empty string, and [node
        document](#concept-node-document){#ref-for-concept-node-document⑥②
        link-type="dfn"} is
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨②
        link-type="dfn"}'s [node
        document](#concept-node-document){#ref-for-concept-node-document⑥③
        link-type="dfn"}, then
        [append](#concept-element-attributes-append){#ref-for-concept-element-attributes-append④
        link-type="dfn"} this
        [attribute](#concept-attribute){#ref-for-concept-attribute⑤⑦
        link-type="dfn"} to
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨③
        link-type="dfn"}, and then return true.

    2.  Return false.

5.  Otherwise, if `force`{.variable} is not given or is false, [remove
    an
    attribute](#concept-element-attributes-remove-by-name){#ref-for-concept-element-attributes-remove-by-name①
    link-type="dfn"} given `qualifiedName`{.variable} and
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨④
    link-type="dfn"}, and then return false.

6.  Return true.

The
[`hasAttributeNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-element-hasattributens
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  If `namespace`{.variable} is the empty string, then set it to null.

2.  Return true if
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨⑤
    link-type="dfn"}
    [has](#concept-element-attribute-has){#ref-for-concept-element-attribute-has⑥
    link-type="dfn"} an
    [attribute](#concept-attribute){#ref-for-concept-attribute⑤⑧
    link-type="dfn"} whose
    [namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①⑦
    link-type="dfn"} is `namespace`{.variable} and [local
    name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②②
    link-type="dfn"} is `localName`{.variable}; otherwise false.

------------------------------------------------------------------------

The
[`getAttributeNode(``qualifiedName`{.variable}`)`]{#dom-element-getattributenode
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to return the result of [getting an
attribute](#concept-element-attributes-get-by-name){#ref-for-concept-element-attributes-get-by-name②
link-type="dfn"} given `qualifiedName`{.variable} and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨⑥
link-type="dfn"}.

The
[`getAttributeNodeNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-element-getattributenodens
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to return the result of [getting an
attribute](#concept-element-attributes-get-by-namespace){#ref-for-concept-element-attributes-get-by-namespace⑤
link-type="dfn"} given `namespace`{.variable}, `localName`{.variable},
and [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨⑦
link-type="dfn"}.

The
[`setAttributeNode(``attr`{.variable}`)`]{#dom-element-setattributenode
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} and
[`setAttributeNodeNS(``attr`{.variable}`)`]{#dom-element-setattributenodens
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} methods steps are to return the result of [setting an
attribute](#concept-element-attributes-set){#ref-for-concept-element-attributes-set
link-type="dfn"} given `attr`{.variable} and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨⑧
link-type="dfn"}.

The
[`removeAttributeNode(``attr`{.variable}`)`]{#dom-element-removeattributenode
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this①⑨⑨
    link-type="dfn"}'s [attribute
    list](#concept-element-attribute){#ref-for-concept-element-attribute①⑨
    link-type="dfn"} does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#ref-for-list-contain⑥
    link-type="dfn"} `attr`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤⑧
    link-type="dfn"} a
    \"[`NotFoundError`{.idl}](https://webidl.spec.whatwg.org/#notfounderror){#ref-for-notfounderror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦②
    link-type="idl"}.

2.  [Remove](#concept-element-attributes-remove){#ref-for-concept-element-attributes-remove②
    link-type="dfn"} `attr`{.variable}.

3.  Return `attr`{.variable}.

------------------------------------------------------------------------

`shadow`{.variable}` = ``element`{.variable}` . `[`attachShadow(init)`{.idl}](#dom-element-attachshadow){#ref-for-dom-element-attachshadow① link-type="idl"}

:   Creates a [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root④⑥
    link-type="dfn"} for `element`{.variable} and returns it.

`shadow`{.variable}` = ``element`{.variable}` . `[`shadowRoot`{.idl}](#dom-element-shadowroot){#ref-for-dom-element-shadowroot① link-type="idl"}

:   Returns `element`{.variable}'s [shadow
    root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①③
    link-type="dfn"}, if any, and if [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root④⑦
    link-type="dfn"}'s
    [mode](#shadowroot-mode){#ref-for-shadowroot-mode⑦ link-type="dfn"}
    is \"`open`\"; otherwise null.

A [valid shadow host name]{#valid-shadow-host-name .dfn .dfn-paneled
dfn-type="dfn" export=""} is:

- a [valid custom element
  name](https://html.spec.whatwg.org/multipage/custom-elements.html#valid-custom-element-name){#ref-for-valid-custom-element-name②
  link-type="dfn"}
- \"`article`\", \"`aside`\", \"`blockquote`\", \"`body`\", \"`div`\",
  \"`footer`\", \"`h1`\", \"`h2`\", \"`h3`\", \"`h4`\", \"`h5`\",
  \"`h6`\", \"`header`\", \"`main`\", \"`nav`\", \"`p`\", \"`section`\",
  or \"`span`\"

::: {.algorithm algorithm="attachShadow(init)" algorithm-for="Element"}
The [`attachShadow(``init`{.variable}`)`]{#dom-element-attachshadow .dfn
.dfn-paneled .idl-code dfn-for="Element" dfn-type="method" export=""}
method steps are:

1.  Let `registry`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪⓪
    link-type="dfn"}'s [custom element
    registry](#element-custom-element-registry){#ref-for-element-custom-element-registry⑦
    link-type="dfn"}.

2.  If
    `init`{.variable}\[\"[`customElementRegistry`{.idl}](#dom-shadowrootinit-customelementregistry){#ref-for-dom-shadowrootinit-customelementregistry
    link-type="idl"}\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#ref-for-map-exists①④
    link-type="dfn"}, then set `registry`{.variable} to it.

3.  Run [attach a shadow
    root](#concept-attach-a-shadow-root){#ref-for-concept-attach-a-shadow-root①
    link-type="dfn"} with
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪①
    link-type="dfn"},
    `init`{.variable}\[\"[`mode`{.idl}](#dom-shadowrootinit-mode){#ref-for-dom-shadowrootinit-mode
    link-type="idl"}\"\],
    `init`{.variable}\[\"[`clonable`{.idl}](#dom-shadowrootinit-clonable){#ref-for-dom-shadowrootinit-clonable
    link-type="idl"}\"\],
    `init`{.variable}\[\"[`serializable`{.idl}](#dom-shadowrootinit-serializable){#ref-for-dom-shadowrootinit-serializable
    link-type="idl"}\"\],
    `init`{.variable}\[\"[`delegatesFocus`{.idl}](#dom-shadowrootinit-delegatesfocus){#ref-for-dom-shadowrootinit-delegatesfocus
    link-type="idl"}\"\],
    `init`{.variable}\[\"[`slotAssignment`{.idl}](#dom-shadowrootinit-slotassignment){#ref-for-dom-shadowrootinit-slotassignment
    link-type="idl"}\"\], and `registry`{.variable}.

4.  Return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪②
    link-type="dfn"}'s [shadow
    root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①④
    link-type="dfn"}.
:::

::: {.algorithm algorithm="attach a shadow root"}
To [attach a shadow root]{#concept-attach-a-shadow-root .dfn
.dfn-paneled dfn-type="dfn" noexport=""}, given an
[element](#concept-element){#ref-for-concept-element①⓪③ link-type="dfn"}
`element`{.variable}, a string `mode`{.variable}, a boolean
`clonable`{.variable}, a boolean `serializable`{.variable}, a boolean
`delegatesFocus`{.variable}, a string `slotAssignment`{.variable}, and
null or a
[`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry①⑥
link-type="idl"} object `registry`{.variable}:

1.  If `element`{.variable}'s
    [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①⑧
    link-type="dfn"} is not the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②⑧
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑤⑨
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦③
    link-type="idl"}.

2.  If `element`{.variable}'s [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name①⑤
    link-type="dfn"} is not a [valid shadow host
    name](#valid-shadow-host-name){#ref-for-valid-shadow-host-name
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥⓪
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①⑦
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦④
    link-type="idl"}.

3.  If `element`{.variable}'s [local
    name](#concept-element-local-name){#ref-for-concept-element-local-name①⑥
    link-type="dfn"} is a [valid custom element
    name](https://html.spec.whatwg.org/multipage/custom-elements.html#valid-custom-element-name){#ref-for-valid-custom-element-name③
    link-type="dfn"}, or `element`{.variable}'s [`is`
    value](#concept-element-is-value){#ref-for-concept-element-is-value③
    link-type="dfn"} is non-null:

    1.  Let `definition`{.variable} be the result of [looking up a
        custom element
        definition](https://html.spec.whatwg.org/multipage/custom-elements.html#look-up-a-custom-element-definition){#ref-for-look-up-a-custom-element-definition①
        link-type="dfn"} given `element`{.variable}'s [custom element
        registry](#element-custom-element-registry){#ref-for-element-custom-element-registry⑧
        link-type="dfn"}, its
        [namespace](#concept-element-namespace){#ref-for-concept-element-namespace①⑨
        link-type="dfn"}, its [local
        name](#concept-element-local-name){#ref-for-concept-element-local-name①⑦
        link-type="dfn"}, and its [`is`
        value](#concept-element-is-value){#ref-for-concept-element-is-value④
        link-type="dfn"}.

    2.  If `definition`{.variable} is not null and
        `definition`{.variable}'s [disable
        shadow](https://html.spec.whatwg.org/multipage/custom-elements.html#concept-custom-element-definition-disable-shadow){#ref-for-concept-custom-element-definition-disable-shadow
        link-type="dfn"} is true, then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥①
        link-type="dfn"} a
        \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①⑧
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦⑤
        link-type="idl"}.

4.  If `element`{.variable} is a [shadow
    host](#element-shadow-host){#ref-for-element-shadow-host⑤
    link-type="dfn"}:

    1.  Let `currentShadowRoot`{.variable} be `element`{.variable}'s
        [shadow
        root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①⑤
        link-type="dfn"}.

    2.  If any of the following are true:

        - `currentShadowRoot`{.variable}'s
          [declarative](#shadowroot-declarative){#ref-for-shadowroot-declarative②
          link-type="dfn"} is false; or

        - `currentShadowRoot`{.variable}'s
          [mode](#shadowroot-mode){#ref-for-shadowroot-mode⑧
          link-type="dfn"} is not `mode`{.variable},

        then
        [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥②
        link-type="dfn"} a
        \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror①⑨
        .idl-code link-type="exception"}\"
        [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦⑥
        link-type="idl"}.

    3.  Otherwise:

        1.  [Remove](#concept-node-remove){#ref-for-concept-node-remove①①
            link-type="dfn"} all of `currentShadowRoot`{.variable}'s
            [children](#concept-tree-child){#ref-for-concept-tree-child⑥②
            link-type="dfn"}, in [tree
            order](#concept-tree-order){#ref-for-concept-tree-order②③
            link-type="dfn"}.

        2.  Set `currentShadowRoot`{.variable}'s
            [declarative](#shadowroot-declarative){#ref-for-shadowroot-declarative③
            link-type="dfn"} to false.

        3.  Return.

5.  Let `shadow`{.variable} be a new [shadow
    root](#concept-shadow-root){#ref-for-concept-shadow-root④⑧
    link-type="dfn"} whose [node
    document](#concept-node-document){#ref-for-concept-node-document⑥④
    link-type="dfn"} is `element`{.variable}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑥⑤
    link-type="dfn"},
    [host](#concept-documentfragment-host){#ref-for-concept-documentfragment-host①⑤
    link-type="dfn"} is `element`{.variable}, and
    [mode](#shadowroot-mode){#ref-for-shadowroot-mode⑨ link-type="dfn"}
    is `mode`{.variable}.

6.  Set `shadow`{.variable}'s [delegates
    focus](#shadowroot-delegates-focus){#ref-for-shadowroot-delegates-focus②
    link-type="dfn"} to `delegatesFocus`{.variable}.

7.  If `element`{.variable}'s [custom element
    state](#concept-element-custom-element-state){#ref-for-concept-element-custom-element-state⑥
    link-type="dfn"} is \"`precustomized`\" or \"`custom`\", then set
    `shadow`{.variable}'s [available to element
    internals](#shadowroot-available-to-element-internals){#ref-for-shadowroot-available-to-element-internals
    link-type="dfn"} to true.

8.  Set `shadow`{.variable}'s [slot
    assignment](#shadowroot-slot-assignment){#ref-for-shadowroot-slot-assignment⑥
    link-type="dfn"} to `slotAssignment`{.variable}.

9.  Set `shadow`{.variable}'s
    [declarative](#shadowroot-declarative){#ref-for-shadowroot-declarative④
    link-type="dfn"} to false.

10. Set `shadow`{.variable}'s
    [clonable](#shadowroot-clonable){#ref-for-shadowroot-clonable②
    link-type="dfn"} to `clonable`{.variable}.

11. Set `shadow`{.variable}'s
    [serializable](#shadowroot-serializable){#ref-for-shadowroot-serializable②
    link-type="dfn"} to `serializable`{.variable}.

12. Set `shadow`{.variable}'s [custom element
    registry](#shadowroot-custom-element-registry){#ref-for-shadowroot-custom-element-registry⑧
    link-type="dfn"} to `registry`{.variable}.

13. Set `element`{.variable}'s [shadow
    root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①⑥
    link-type="dfn"} to `shadow`{.variable}.
:::

::: {.algorithm algorithm="shadowRoot" algorithm-for="Element"}
The [`shadowRoot`]{#dom-element-shadowroot .dfn .dfn-paneled .idl-code
dfn-for="Element" dfn-type="attribute" export=""} getter steps are:

1.  Let `shadow`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪③
    link-type="dfn"}'s [shadow
    root](#concept-element-shadow-root){#ref-for-concept-element-shadow-root①⑦
    link-type="dfn"}.

2.  If `shadow`{.variable} is null or its
    [mode](#shadowroot-mode){#ref-for-shadowroot-mode①⓪ link-type="dfn"}
    is \"`closed`\", then return null.

3.  Return `shadow`{.variable}.
:::

------------------------------------------------------------------------

`registry`{.variable}` = ``element`{.variable}` . `[`customElementRegistry`{.idl}](#dom-element-customelementregistry){#ref-for-dom-element-customelementregistry① link-type="idl"}

:   Returns `element`{.variable}'s
    [`CustomElementRegistry`{.idl}](https://html.spec.whatwg.org/multipage/custom-elements.html#customelementregistry){#ref-for-customelementregistry①⑦
    link-type="idl"} object, if any; otherwise null.

::: {.algorithm algorithm="customElementRegistry" algorithm-for="Element"}
The [`customElementRegistry`]{#dom-element-customelementregistry .dfn
.dfn-paneled .idl-code dfn-for="Element" dfn-type="attribute" export=""}
getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪④
link-type="dfn"}'s [custom element
registry](#element-custom-element-registry){#ref-for-element-custom-element-registry⑨
link-type="dfn"}.
:::

------------------------------------------------------------------------

`element`{.variable}` . `[`closest(selectors)`{.idl}](#dom-element-closest){#ref-for-dom-element-closest① link-type="idl"}
:   Returns the first (starting at `element`{.variable}) [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor⑤
    link-type="dfn"} that matches `selectors`{.variable}, and null
    otherwise.

`element`{.variable}` . `[`matches(selectors)`{.idl}](#dom-element-matches){#ref-for-dom-element-matches① link-type="idl"}
:   Returns true if matching `selectors`{.variable} against
    `element`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root④②
    link-type="dfn"} yields `element`{.variable}; otherwise false.

The [`closest(``selectors`{.variable}`)`]{#dom-element-closest .dfn
.dfn-paneled .idl-code dfn-for="Element" dfn-type="method" export=""}
method steps are:

1.  Let `s`{.variable} be the result of [parse a
    selector](https://drafts.csswg.org/selectors-4/#parse-a-selector){#ref-for-parse-a-selector①
    link-type="dfn"} from `selectors`{.variable}.
    [\[SELECTORS4\]](#biblio-selectors4 "Selectors Level 4"){link-type="biblio"}

2.  If `s`{.variable} is failure, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥③
    link-type="dfn"} a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦⑦
    link-type="idl"}.

3.  Let `elements`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪⑤
    link-type="dfn"}'s [inclusive
    ancestors](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor⑥
    link-type="dfn"} that are
    [elements](#concept-element){#ref-for-concept-element①⓪④
    link-type="dfn"}, in reverse [tree
    order](#concept-tree-order){#ref-for-concept-tree-order②④
    link-type="dfn"}.

4.  For each `element`{.variable} in `elements`{.variable}, if [match a
    selector against an
    element](https://drafts.csswg.org/selectors-4/#match-a-selector-against-an-element){#ref-for-match-a-selector-against-an-element
    link-type="dfn"}, using `s`{.variable}, `element`{.variable}, and
    [scoping
    root](https://drafts.csswg.org/selectors-4/#scoping-root){#ref-for-scoping-root①
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪⑥
    link-type="dfn"}, returns success, return `element`{.variable}.
    [\[SELECTORS4\]](#biblio-selectors4 "Selectors Level 4"){link-type="biblio"}

5.  Return null.

The [`matches(``selectors`{.variable}`)`]{#dom-element-matches .dfn
.dfn-paneled .idl-code dfn-for="Element" dfn-type="method" export=""}
and
[`webkitMatchesSelector(``selectors`{.variable}`)`]{#dom-element-webkitmatchesselector
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  Let `s`{.variable} be the result of [parse a
    selector](https://drafts.csswg.org/selectors-4/#parse-a-selector){#ref-for-parse-a-selector②
    link-type="dfn"} from `selectors`{.variable}.
    [\[SELECTORS4\]](#biblio-selectors4 "Selectors Level 4"){link-type="biblio"}

2.  If `s`{.variable} is failure, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥④
    link-type="dfn"} a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦⑧
    link-type="idl"}.

3.  If the result of [match a selector against an
    element](https://drafts.csswg.org/selectors-4/#match-a-selector-against-an-element){#ref-for-match-a-selector-against-an-element①
    link-type="dfn"}, using `s`{.variable},
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪⑦
    link-type="dfn"}, and [scoping
    root](https://drafts.csswg.org/selectors-4/#scoping-root){#ref-for-scoping-root②
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪⑧
    link-type="dfn"}, returns success, then return true; otherwise,
    return false.
    [\[SELECTORS4\]](#biblio-selectors4 "Selectors Level 4"){link-type="biblio"}

------------------------------------------------------------------------

The
[`getElementsByTagName(``qualifiedName`{.variable}`)`]{#dom-element-getelementsbytagname
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to return the [list of elements with
qualified name
`qualifiedName`{.variable}](#concept-getelementsbytagname){#ref-for-concept-getelementsbytagname①
link-type="dfn"} for
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⓪⑨
link-type="dfn"}.

The
[`getElementsByTagNameNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-element-getelementsbytagnamens
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to return the [list of elements with
namespace `namespace`{.variable} and local name
`localName`{.variable}](#concept-getelementsbytagnamens){#ref-for-concept-getelementsbytagnamens①
link-type="dfn"} for
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①⓪
link-type="dfn"}.

The
[`getElementsByClassName(``classNames`{.variable}`)`]{#dom-element-getelementsbyclassname
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to return the [list of elements with class
names
`classNames`{.variable}](#concept-getelementsbyclassname){#ref-for-concept-getelementsbyclassname①
link-type="dfn"} for
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①①
link-type="dfn"}.

------------------------------------------------------------------------

To [insert adjacent]{#insert-adjacent .dfn .dfn-paneled dfn-type="dfn"
noexport=""}, given an
[element](#concept-element){#ref-for-concept-element①⓪⑤ link-type="dfn"}
`element`{.variable}, string `where`{.variable}, and a
[node](#concept-node){#ref-for-concept-node①⑤③ link-type="dfn"}
`node`{.variable}, run the steps associated with the first [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#ref-for-ascii-case-insensitive②
link-type="dfn"} match for `where`{.variable}:

\"`beforebegin`\"

:   If `element`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent③③
    link-type="dfn"} is null, return null.

    Return the result of
    [pre-inserting](#concept-node-pre-insert){#ref-for-concept-node-pre-insert⑧
    link-type="dfn"} `node`{.variable} into `element`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent③④
    link-type="dfn"} before `element`{.variable}.

\"`afterbegin`\"

:   Return the result of
    [pre-inserting](#concept-node-pre-insert){#ref-for-concept-node-pre-insert⑨
    link-type="dfn"} `node`{.variable} into `element`{.variable} before
    `element`{.variable}'s [first
    child](#concept-tree-first-child){#ref-for-concept-tree-first-child⑤
    link-type="dfn"}.

\"`beforeend`\"

:   Return the result of
    [pre-inserting](#concept-node-pre-insert){#ref-for-concept-node-pre-insert①⓪
    link-type="dfn"} `node`{.variable} into `element`{.variable} before
    null.

\"`afterend`\"

:   If `element`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent③⑤
    link-type="dfn"} is null, return null.

    Return the result of
    [pre-inserting](#concept-node-pre-insert){#ref-for-concept-node-pre-insert①①
    link-type="dfn"} `node`{.variable} into `element`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent③⑥
    link-type="dfn"} before `element`{.variable}'s [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①②
    link-type="dfn"}.

Otherwise

:   [Throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥⑤
    link-type="dfn"} a
    \"[`SyntaxError`{.idl}](https://webidl.spec.whatwg.org/#syntaxerror){#ref-for-syntaxerror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑦⑨
    link-type="idl"}.

The
[`insertAdjacentElement(``where`{.variable}`, ``element`{.variable}`)`]{#dom-element-insertadjacentelement
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are to return the result of running [insert
adjacent](#insert-adjacent){#ref-for-insert-adjacent link-type="dfn"},
give [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①②
link-type="dfn"}, `where`{.variable}, and `element`{.variable}.

The
[`insertAdjacentText(``where`{.variable}`, ``data`{.variable}`)`]{#dom-element-insertadjacenttext
.dfn .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
export=""} method steps are:

1.  Let `text`{.variable} be a new [`Text`{.idl}](#text){#ref-for-text②⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑤④
    link-type="dfn"} whose
    [data](#concept-cd-data){#ref-for-concept-cd-data②③ link-type="dfn"}
    is `data`{.variable} and [node
    document](#concept-node-document){#ref-for-concept-node-document⑥⑥
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①③
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑥⑦
    link-type="dfn"}.

2.  Run [insert adjacent](#insert-adjacent){#ref-for-insert-adjacent①
    link-type="dfn"}, given
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①④
    link-type="dfn"}, `where`{.variable}, and `text`{.variable}.

This method returns nothing because it existed before we had a chance to
design it.

#### [4.9.1. ]{.secno}[Interface [`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap③ link-type="idl"}]{.content}[](#interface-namednodemap){.self-link} {#interface-namednodemap .heading .settled level="4.9.1"}

``` {.idl .highlight .def}
[Exposed=Window,
 LegacyUnenumerableNamedProperties]
interface NamedNodeMap {
  readonly attribute unsigned long length;
  getter Attr? item(unsigned long index);
  getter Attr? getNamedItem(DOMString qualifiedName);
  Attr? getNamedItemNS(DOMString? namespace, DOMString localName);
  [CEReactions] Attr? setNamedItem(Attr attr);
  [CEReactions] Attr? setNamedItemNS(Attr attr);
  [CEReactions] Attr removeNamedItem(DOMString qualifiedName);
  [CEReactions] Attr removeNamedItemNS(DOMString? namespace, DOMString localName);
};
```

A [`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap④
link-type="idl"} has an associated
[element]{#concept-namednodemap-element .dfn .dfn-paneled
dfn-for="NamedNodeMap" dfn-type="dfn" export=""} (an
[element](#concept-element){#ref-for-concept-element①⓪⑥
link-type="dfn"}).

A [`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap⑤
link-type="idl"} object's [attribute
list]{#concept-namednodemap-attribute .dfn .dfn-paneled
dfn-for="NamedNodeMap" dfn-type="dfn" export=""} is its
[element](#concept-namednodemap-element){#ref-for-concept-namednodemap-element
link-type="dfn"}'s [attribute
list](#concept-element-attribute){#ref-for-concept-element-attribute②⓪
link-type="dfn"}.

------------------------------------------------------------------------

A [`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap⑥
link-type="idl"} object's [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices④
link-type="dfn"} are the numbers in the range zero to its [attribute
list](#concept-namednodemap-attribute){#ref-for-concept-namednodemap-attribute
link-type="dfn"}'s
[size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size④
link-type="dfn"} minus one, unless the [attribute
list](#concept-namednodemap-attribute){#ref-for-concept-namednodemap-attribute①
link-type="dfn"} [is
empty](https://infra.spec.whatwg.org/#list-is-empty){#ref-for-list-is-empty①⓪
link-type="dfn"}, in which case there are no [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#ref-for-dfn-supported-property-indices⑤
link-type="dfn"}.

The [`length`]{#dom-namednodemap-length .dfn .dfn-paneled .idl-code
dfn-for="NamedNodeMap" dfn-type="attribute" export=""} getter steps are
to return the [attribute
list](#concept-namednodemap-attribute){#ref-for-concept-namednodemap-attribute②
link-type="dfn"}'s
[size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑤
link-type="dfn"}.

The [`item(``index`{.variable}`)`]{#dom-namednodemap-item .dfn
.dfn-paneled .idl-code dfn-for="NamedNodeMap" dfn-type="method"
export=""} method steps are:

1.  If `index`{.variable} is equal to or greater than
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①⑤
    link-type="dfn"}'s [attribute
    list](#concept-namednodemap-attribute){#ref-for-concept-namednodemap-attribute③
    link-type="dfn"}'s
    [size](https://infra.spec.whatwg.org/#list-size){#ref-for-list-size⑥
    link-type="dfn"}, then return null.

2.  Otherwise, return
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①⑥
    link-type="dfn"}'s [attribute
    list](#concept-namednodemap-attribute){#ref-for-concept-namednodemap-attribute④
    link-type="dfn"}\[`index`{.variable}\].

A [`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap⑦
link-type="idl"} object's [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#ref-for-dfn-supported-property-names①
link-type="dfn"} are the return value of running these steps:

1.  Let `names`{.variable} be the [qualified
    names](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name①②
    link-type="dfn"} of the
    [attributes](#concept-attribute){#ref-for-concept-attribute⑤⑨
    link-type="dfn"} in this
    [`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap⑧
    link-type="idl"} object's [attribute
    list](#concept-namednodemap-attribute){#ref-for-concept-namednodemap-attribute⑤
    link-type="dfn"}, with duplicates omitted, in order.

2.  If this [`NamedNodeMap`{.idl}](#namednodemap){#ref-for-namednodemap⑨
    link-type="idl"} object's
    [element](#concept-namednodemap-element){#ref-for-concept-namednodemap-element①
    link-type="dfn"} is in the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#ref-for-html-namespace②⑨
    link-type="dfn"} and its [node
    document](#concept-node-document){#ref-for-concept-node-document⑥⑧
    link-type="dfn"} is an [HTML
    document](#html-document){#ref-for-html-document①⑥ link-type="dfn"},
    then [for
    each](https://infra.spec.whatwg.org/#list-iterate){#ref-for-list-iterate②⑥
    link-type="dfn"} `name`{.variable} of `names`{.variable}:

    1.  Let `lowercaseName`{.variable} be `name`{.variable}, in [ASCII
        lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#ref-for-ascii-lowercase⑦
        link-type="dfn"}.

    2.  If `lowercaseName`{.variable} is not equal to `name`{.variable},
        remove `name`{.variable} from `names`{.variable}.

3.  Return `names`{.variable}.

The
[`getNamedItem(``qualifiedName`{.variable}`)`]{#dom-namednodemap-getnameditem
.dfn .dfn-paneled .idl-code dfn-for="NamedNodeMap" dfn-type="method"
export=""} method steps are to return the result of [getting an
attribute](#concept-element-attributes-get-by-name){#ref-for-concept-element-attributes-get-by-name③
link-type="dfn"} given `qualifiedName`{.variable} and
[element](#concept-namednodemap-element){#ref-for-concept-namednodemap-element②
link-type="dfn"}.

The
[`getNamedItemNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-namednodemap-getnameditemns
.dfn .dfn-paneled .idl-code dfn-for="NamedNodeMap" dfn-type="method"
export=""} method steps are to return the result of [getting an
attribute](#concept-element-attributes-get-by-namespace){#ref-for-concept-element-attributes-get-by-namespace⑥
link-type="dfn"} given `namespace`{.variable}, `localName`{.variable},
and
[element](#concept-namednodemap-element){#ref-for-concept-namednodemap-element③
link-type="dfn"}.

The [`setNamedItem(``attr`{.variable}`)`]{#dom-namednodemap-setnameditem
.dfn .dfn-paneled .idl-code dfn-for="NamedNodeMap" dfn-type="method"
export=""} and
[`setNamedItemNS(``attr`{.variable}`)`]{#dom-namednodemap-setnameditemns
.dfn .dfn-paneled .idl-code dfn-for="NamedNodeMap" dfn-type="method"
export=""} method steps are to return the result of [setting an
attribute](#concept-element-attributes-set){#ref-for-concept-element-attributes-set①
link-type="dfn"} given `attr`{.variable} and
[element](#concept-namednodemap-element){#ref-for-concept-namednodemap-element④
link-type="dfn"}.

The
[`removeNamedItem(``qualifiedName`{.variable}`)`]{#dom-namednodemap-removenameditem
.dfn .dfn-paneled .idl-code dfn-for="NamedNodeMap" dfn-type="method"
export=""} method steps are:

1.  Let `attr`{.variable} be the result of [removing an
    attribute](#concept-element-attributes-remove-by-name){#ref-for-concept-element-attributes-remove-by-name②
    link-type="dfn"} given `qualifiedName`{.variable} and
    [element](#concept-namednodemap-element){#ref-for-concept-namednodemap-element⑤
    link-type="dfn"}.

2.  If `attr`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥⑥
    link-type="dfn"} a
    \"[`NotFoundError`{.idl}](https://webidl.spec.whatwg.org/#notfounderror){#ref-for-notfounderror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧⓪
    link-type="idl"}.

3.  Return `attr`{.variable}.

The
[`removeNamedItemNS(``namespace`{.variable}`, ``localName`{.variable}`)`]{#dom-namednodemap-removenameditemns
.dfn .dfn-paneled .idl-code dfn-for="NamedNodeMap" dfn-type="method"
export=""} method steps are:

1.  Let `attr`{.variable} be the result of [removing an
    attribute](#concept-element-attributes-remove-by-namespace){#ref-for-concept-element-attributes-remove-by-namespace①
    link-type="dfn"} given `namespace`{.variable},
    `localName`{.variable}, and
    [element](#concept-namednodemap-element){#ref-for-concept-namednodemap-element⑥
    link-type="dfn"}.

2.  If `attr`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥⑦
    link-type="dfn"} a
    \"[`NotFoundError`{.idl}](https://webidl.spec.whatwg.org/#notfounderror){#ref-for-notfounderror⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧①
    link-type="idl"}.

3.  Return `attr`{.variable}.

#### [4.9.2. ]{.secno}[Interface [`Attr`{.idl}](#attr){#ref-for-attr③⑥ link-type="idl"}]{.content}[](#interface-attr){.self-link} {#interface-attr .heading .settled level="4.9.2"}

``` {.idl .highlight .def}
[Exposed=Window]
interface Attr : Node {
  readonly attribute DOMString? namespaceURI;
  readonly attribute DOMString? prefix;
  readonly attribute DOMString localName;
  readonly attribute DOMString name;
  [CEReactions] attribute DOMString value;

  readonly attribute Element? ownerElement;

  readonly attribute boolean specified; // useless; always returns true
};
```

[`Attr`{.idl}](#attr){#ref-for-attr③⑦ link-type="idl"}
[nodes](#concept-node){#ref-for-concept-node①⑤⑤ link-type="dfn"} are
simply known as [attributes]{#concept-attribute .dfn .dfn-paneled
dfn-type="dfn" export="" lt="attribute"}. They are sometimes referred to
as *content attributes* to avoid confusion with IDL attributes.

[Attributes](#concept-attribute){#ref-for-concept-attribute⑥⓪
link-type="dfn"} have a [namespace]{#concept-attribute-namespace .dfn
.dfn-paneled dfn-for="Attr" dfn-type="dfn" export=""} (null or a
non-empty string), [namespace
prefix]{#concept-attribute-namespace-prefix .dfn .dfn-paneled
dfn-for="Attr" dfn-type="dfn" export=""} (null or a non-empty string),
[local name]{#concept-attribute-local-name .dfn .dfn-paneled
dfn-for="Attr" dfn-type="dfn" export=""} (a non-empty string),
[value]{#concept-attribute-value .dfn .dfn-paneled dfn-for="Attr"
dfn-type="dfn" export=""} (a string), and
[element]{#concept-attribute-element .dfn .dfn-paneled dfn-for="Attr"
dfn-type="dfn" export=""} (null or an
[element](#concept-element){#ref-for-concept-element①⓪⑦
link-type="dfn"}).

If designed today they would just have a name and value. ☹

An [attribute](#concept-attribute){#ref-for-concept-attribute⑥①
link-type="dfn"}'s [qualified name]{#concept-attribute-qualified-name
.dfn .dfn-paneled dfn-for="Attr" dfn-type="dfn" export=""} is its [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②③
link-type="dfn"} if its [namespace
prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix⑥
link-type="dfn"} is null, and its [namespace
prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix⑦
link-type="dfn"}, followed by \"`:`\", followed by its [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②④
link-type="dfn"}, otherwise.

User agents could have this as an internal slot as an optimization.

When an [attribute](#concept-attribute){#ref-for-concept-attribute⑥②
link-type="dfn"} is created, its [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②⑤
link-type="dfn"} is given. Unless explicitly given when an
[attribute](#concept-attribute){#ref-for-concept-attribute⑥③
link-type="dfn"} is created, its
[namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①⑧
link-type="dfn"}, [namespace
prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix⑧
link-type="dfn"}, and
[element](#concept-attribute-element){#ref-for-concept-attribute-element①④
link-type="dfn"} are set to null, and its
[value](#concept-attribute-value){#ref-for-concept-attribute-value②⑤
link-type="dfn"} is set to the empty string.

An [`A`{.variable} attribute]{#concept-named-attribute .dfn .dfn-paneled
dfn-type="dfn" export="" lt="named attribute"} is an
[attribute](#concept-attribute){#ref-for-concept-attribute⑥④
link-type="dfn"} whose [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②⑥
link-type="dfn"} is `A`{.variable} and whose
[namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace①⑨
link-type="dfn"} and [namespace
prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix⑨
link-type="dfn"} are null.

------------------------------------------------------------------------

The [`namespaceURI`]{#dom-attr-namespaceuri .dfn .dfn-paneled .idl-code
dfn-for="Attr" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①⑦
link-type="dfn"}'s
[namespace](#concept-attribute-namespace){#ref-for-concept-attribute-namespace②⓪
link-type="dfn"}.

The [`prefix`]{#dom-attr-prefix .dfn .dfn-paneled .idl-code
dfn-for="Attr" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①⑧
link-type="dfn"}'s [namespace
prefix](#concept-attribute-namespace-prefix){#ref-for-concept-attribute-namespace-prefix①⓪
link-type="dfn"}.

The [`localName`]{#dom-attr-localname .dfn .dfn-paneled .idl-code
dfn-for="Attr" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②①⑨
link-type="dfn"}'s [local
name](#concept-attribute-local-name){#ref-for-concept-attribute-local-name②⑦
link-type="dfn"}.

The [`name`]{#dom-attr-name .dfn .dfn-paneled .idl-code dfn-for="Attr"
dfn-type="attribute" export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②⓪
link-type="dfn"}'s [qualified
name](#concept-attribute-qualified-name){#ref-for-concept-attribute-qualified-name①③
link-type="dfn"}.

The [`value`]{#dom-attr-value .dfn .dfn-paneled .idl-code dfn-for="Attr"
dfn-type="attribute" export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②①
link-type="dfn"}'s
[value](#concept-attribute-value){#ref-for-concept-attribute-value②⑥
link-type="dfn"}.

To [set an existing attribute value]{#set-an-existing-attribute-value
.dfn .dfn-paneled dfn-type="dfn" noexport=""}, given an
[attribute](#concept-attribute){#ref-for-concept-attribute⑥⑤
link-type="dfn"} `attribute`{.variable} and string `value`{.variable},
run these steps:

1.  If `attribute`{.variable}'s
    [element](#concept-attribute-element){#ref-for-concept-attribute-element①⑤
    link-type="dfn"} is null, then set `attribute`{.variable}'s
    [value](#concept-attribute-value){#ref-for-concept-attribute-value②⑦
    link-type="dfn"} to `value`{.variable}.

2.  Otherwise,
    [change](#concept-element-attributes-change){#ref-for-concept-element-attributes-change②
    link-type="dfn"} `attribute`{.variable} to `value`{.variable}.

The [`value`{.idl}](#dom-attr-value){#ref-for-dom-attr-value①
link-type="idl"} setter steps are to [set an existing attribute
value](#set-an-existing-attribute-value){#ref-for-set-an-existing-attribute-value②
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②②
link-type="dfn"} and the given value.

------------------------------------------------------------------------

The [`ownerElement`]{#dom-attr-ownerelement .dfn .dfn-paneled .idl-code
dfn-for="Attr" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②③
link-type="dfn"}'s
[element](#concept-attribute-element){#ref-for-concept-attribute-element①⑥
link-type="dfn"}.

------------------------------------------------------------------------

The [`specified`]{#dom-attr-specified .dfn .dfn-paneled .idl-code
dfn-for="Attr" dfn-type="attribute" export=""} getter steps are to
return true.

### [4.10. ]{.secno}[Interface [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①⑥ link-type="idl"}]{.content}[](#interface-characterdata){.self-link} {#interface-characterdata .heading .settled level="4.10"}

``` {.idl .highlight .def}
[Exposed=Window]
interface CharacterData : Node {
  attribute [LegacyNullToEmptyString] DOMString data;
  readonly attribute unsigned long length;
  DOMString substringData(unsigned long offset, unsigned long count);
  undefined appendData(DOMString data);
  undefined insertData(unsigned long offset, DOMString data);
  undefined deleteData(unsigned long offset, unsigned long count);
  undefined replaceData(unsigned long offset, unsigned long count, DOMString data);
};
```

[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①⑦
link-type="idl"} is an abstract interface. You cannot get a direct
instance of it. It is used by [`Text`{.idl}](#text){#ref-for-text②⑨
link-type="idl"},
[`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①④
link-type="idl"}, and [`Comment`{.idl}](#comment){#ref-for-comment①④
link-type="idl"} [nodes](#concept-node){#ref-for-concept-node①⑤⑥
link-type="dfn"}.

Each [node](#concept-node){#ref-for-concept-node①⑤⑦ link-type="dfn"}
inheriting from the
[`CharacterData`{.idl}](#characterdata){#ref-for-characterdata①⑧
link-type="idl"} interface has an associated mutable string called
[data]{#concept-cd-data .dfn .dfn-paneled dfn-for="CharacterData"
dfn-type="dfn" export=""}.

To [replace data]{#concept-cd-replace .dfn .dfn-paneled dfn-type="dfn"
export=""} of node `node`{.variable} with offset `offset`{.variable},
count `count`{.variable}, and data `data`{.variable}, run these steps:

1.  Let `length`{.variable} be `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length⑤
    link-type="dfn"}.

2.  If `offset`{.variable} is greater than `length`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥⑧
    link-type="dfn"} an
    \"[`IndexSizeError`{.idl}](https://webidl.spec.whatwg.org/#indexsizeerror){#ref-for-indexsizeerror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧②
    link-type="idl"}.

3.  If `offset`{.variable} plus `count`{.variable} is greater than
    `length`{.variable}, then set `count`{.variable} to
    `length`{.variable} minus `offset`{.variable}.

4.  [Queue a mutation
    record](#queue-a-mutation-record){#ref-for-queue-a-mutation-record②
    link-type="dfn"} of \"`characterData`\" for `node`{.variable} with
    null, null, `node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data②④
    link-type="dfn"}, « », « », null, and null.

5.  Insert `data`{.variable} into `node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data②⑤ link-type="dfn"}
    after `offset`{.variable} [code
    units](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit
    link-type="dfn"}.

6.  Let `delete offset`{.variable} be `offset`{.variable} +
    `data`{.variable}'s
    [length](https://infra.spec.whatwg.org/#string-length){#ref-for-string-length①
    link-type="dfn"}.

7.  Starting from `delete offset`{.variable} [code
    units](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit①
    link-type="dfn"}, remove `count`{.variable} [code
    units](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit②
    link-type="dfn"} from `node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data②⑥
    link-type="dfn"}.

8.  For each [live
    range](#concept-live-range){#ref-for-concept-live-range⑨
    link-type="dfn"} whose [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node⑥
    link-type="dfn"} is `node`{.variable} and [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset⑦
    link-type="dfn"} is greater than `offset`{.variable} but less than
    or equal to `offset`{.variable} plus `count`{.variable}, set its
    [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset⑧
    link-type="dfn"} to `offset`{.variable}.

9.  For each [live
    range](#concept-live-range){#ref-for-concept-live-range①⓪
    link-type="dfn"} whose [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node⑥
    link-type="dfn"} is `node`{.variable} and [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset⑦
    link-type="dfn"} is greater than `offset`{.variable} but less than
    or equal to `offset`{.variable} plus `count`{.variable}, set its
    [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset⑧
    link-type="dfn"} to `offset`{.variable}.

10. For each [live
    range](#concept-live-range){#ref-for-concept-live-range①①
    link-type="dfn"} whose [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node⑦
    link-type="dfn"} is `node`{.variable} and [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset⑨
    link-type="dfn"} is greater than `offset`{.variable} plus
    `count`{.variable}, increase its [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①⓪
    link-type="dfn"} by `data`{.variable}'s
    [length](https://infra.spec.whatwg.org/#string-length){#ref-for-string-length②
    link-type="dfn"} and decrease it by `count`{.variable}.

11. For each [live
    range](#concept-live-range){#ref-for-concept-live-range①②
    link-type="dfn"} whose [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node⑦
    link-type="dfn"} is `node`{.variable} and [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset⑨
    link-type="dfn"} is greater than `offset`{.variable} plus
    `count`{.variable}, increase its [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①⓪
    link-type="dfn"} by `data`{.variable}'s
    [length](https://infra.spec.whatwg.org/#string-length){#ref-for-string-length③
    link-type="dfn"} and decrease it by `count`{.variable}.

12. If `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent③⑦
    link-type="dfn"} is non-null, then run the [children changed
    steps](#concept-node-children-changed-ext){#ref-for-concept-node-children-changed-ext②
    link-type="dfn"} for `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent③⑧
    link-type="dfn"}.

To [substring data]{#concept-cd-substring .dfn .dfn-paneled
dfn-for="CharacterData, Text, Comment, ProcessingInstruction"
dfn-type="dfn" export=""} with node `node`{.variable}, offset
`offset`{.variable}, and count `count`{.variable}, run these steps:

1.  Let `length`{.variable} be `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length⑥
    link-type="dfn"}.
2.  If `offset`{.variable} is greater than `length`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑥⑨
    link-type="dfn"} an
    \"[`IndexSizeError`{.idl}](https://webidl.spec.whatwg.org/#indexsizeerror){#ref-for-indexsizeerror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧③
    link-type="idl"}.
3.  If `offset`{.variable} plus `count`{.variable} is greater than
    `length`{.variable}, return a string whose value is the [code
    units](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit③
    link-type="dfn"} from the `offset`{.variable}^th^ [code
    unit](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit④
    link-type="dfn"} to the end of `node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data②⑦
    link-type="dfn"}, and then return.
4.  Return a string whose value is the [code
    units](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit⑤
    link-type="dfn"} from the `offset`{.variable}^th^ [code
    unit](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit⑥
    link-type="dfn"} to the `offset`{.variable}+`count`{.variable}^th^
    [code
    unit](https://infra.spec.whatwg.org/#code-unit){#ref-for-code-unit⑦
    link-type="dfn"} in `node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data②⑧
    link-type="dfn"}.

The [`data`]{#dom-characterdata-data .dfn .dfn-paneled .idl-code
dfn-for="CharacterData" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②④
link-type="dfn"}'s [data](#concept-cd-data){#ref-for-concept-cd-data②⑨
link-type="dfn"}. Its setter must [replace
data](#concept-cd-replace){#ref-for-concept-cd-replace④ link-type="dfn"}
with node [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②⑤
link-type="dfn"}, offset 0, count
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②⑥
link-type="dfn"}'s
[length](#concept-node-length){#ref-for-concept-node-length⑦
link-type="dfn"}, and data new value.

The [`length`]{#dom-characterdata-length .dfn .dfn-paneled .idl-code
dfn-for="CharacterData" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②⑦
link-type="dfn"}'s
[length](#concept-node-length){#ref-for-concept-node-length⑧
link-type="dfn"}.

The
[`substringData(``offset`{.variable}`, ``count`{.variable}`)`]{#dom-characterdata-substringdata
.dfn .dfn-paneled .idl-code dfn-for="CharacterData" dfn-type="method"
export=""} method steps are to return the result of running [substring
data](#concept-cd-substring){#ref-for-concept-cd-substring
link-type="dfn"} with node
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②⑧
link-type="dfn"}, offset `offset`{.variable}, and count
`count`{.variable}.

The [`appendData(``data`{.variable}`)`]{#dom-characterdata-appenddata
.dfn .dfn-paneled .idl-code dfn-for="CharacterData" dfn-type="method"
export=""} method steps are to [replace
data](#concept-cd-replace){#ref-for-concept-cd-replace⑤ link-type="dfn"}
with node [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②②⑨
link-type="dfn"}, offset
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③⓪
link-type="dfn"}'s
[length](#concept-node-length){#ref-for-concept-node-length⑨
link-type="dfn"}, count 0, and data `data`{.variable}.

The
[`insertData(``offset`{.variable}`, ``data`{.variable}`)`]{#dom-characterdata-insertdata
.dfn .dfn-paneled .idl-code dfn-for="CharacterData" dfn-type="method"
export=""} method steps are to [replace
data](#concept-cd-replace){#ref-for-concept-cd-replace⑥ link-type="dfn"}
with node [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③①
link-type="dfn"}, offset `offset`{.variable}, count 0, and data
`data`{.variable}.

The
[`deleteData(``offset`{.variable}`, ``count`{.variable}`)`]{#dom-characterdata-deletedata
.dfn .dfn-paneled .idl-code dfn-for="CharacterData" dfn-type="method"
export=""} method steps are to [replace
data](#concept-cd-replace){#ref-for-concept-cd-replace⑦ link-type="dfn"}
with node [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③②
link-type="dfn"}, offset `offset`{.variable}, count `count`{.variable},
and data the empty string.

The
[`replaceData(``offset`{.variable}`, ``count`{.variable}`, ``data`{.variable}`)`]{#dom-characterdata-replacedata
.dfn .dfn-paneled .idl-code dfn-for="CharacterData" dfn-type="method"
export=""} method steps are to [replace
data](#concept-cd-replace){#ref-for-concept-cd-replace⑧ link-type="dfn"}
with node [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③③
link-type="dfn"}, offset `offset`{.variable}, count `count`{.variable},
and data `data`{.variable}.

### [4.11. ]{.secno}[Interface [`Text`{.idl}](#text){#ref-for-text③⓪ link-type="idl"}]{.content}[](#interface-text){.self-link} {#interface-text .heading .settled level="4.11"}

``` {.idl .highlight .def}
[Exposed=Window]
interface Text : CharacterData {
  constructor(optional DOMString data = "");

  [NewObject] Text splitText(unsigned long offset);
  readonly attribute DOMString wholeText;
};
```

`text`{.variable}` = new `[`Text([``data`{.variable}` = ""])`](#dom-text-text){#ref-for-dom-text-text① .idl-code link-type="constructor"}
:   Returns a new [`Text`{.idl}](#text){#ref-for-text③② link-type="idl"}
    [node](#concept-node){#ref-for-concept-node①⑤⑧ link-type="dfn"}
    whose [data](#concept-cd-data){#ref-for-concept-cd-data③⓪
    link-type="dfn"} is `data`{.variable}.

`text`{.variable}` . `[`splitText(offset)`{.idl}](#dom-text-splittext){#ref-for-dom-text-splittext① link-type="idl"}
:   Splits [data](#concept-cd-data){#ref-for-concept-cd-data③①
    link-type="dfn"} at the given `offset`{.variable} and returns the
    remainder as [`Text`{.idl}](#text){#ref-for-text③③ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node①⑤⑨ link-type="dfn"}.

`text`{.variable}` . `[`wholeText`{.idl}](#dom-text-wholetext){#ref-for-dom-text-wholetext① link-type="idl"}
:   Returns the combined
    [data](#concept-cd-data){#ref-for-concept-cd-data③② link-type="dfn"}
    of all direct [`Text`{.idl}](#text){#ref-for-text③④ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node①⑥⓪ link-type="dfn"}
    [siblings](#concept-tree-sibling){#ref-for-concept-tree-sibling①③
    link-type="dfn"}.

------------------------------------------------------------------------

An [exclusive [`Text`{.idl}](#text){#ref-for-text③⑤ link-type="idl"}
node]{#exclusive-text-node .dfn .dfn-paneled dfn-type="dfn" export=""}
is a [`Text`{.idl}](#text){#ref-for-text③⑥ link-type="idl"}
[node](#concept-node){#ref-for-concept-node①⑥① link-type="dfn"} that is
not a [`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection⑨
link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑥②
link-type="dfn"}.

The [contiguous [`Text`{.idl}](#text){#ref-for-text③⑦ link-type="idl"}
nodes]{#contiguous-text-nodes .dfn .dfn-paneled dfn-type="dfn"
export=""} of a [node](#concept-node){#ref-for-concept-node①⑥③
link-type="dfn"} `node`{.variable} are `node`{.variable},
`node`{.variable}'s [previous
sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling⑧
link-type="dfn"} [`Text`{.idl}](#text){#ref-for-text③⑧ link-type="idl"}
[node](#concept-node){#ref-for-concept-node①⑥④ link-type="dfn"}, if any,
and its [contiguous `Text`
nodes](#contiguous-text-nodes){#ref-for-contiguous-text-nodes
link-type="dfn"}, and `node`{.variable}'s [next
sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①③
link-type="dfn"} [`Text`{.idl}](#text){#ref-for-text③⑨ link-type="idl"}
[node](#concept-node){#ref-for-concept-node①⑥⑤ link-type="dfn"}, if any,
and its [contiguous `Text`
nodes](#contiguous-text-nodes){#ref-for-contiguous-text-nodes①
link-type="dfn"}, avoiding any duplicates.

The [contiguous exclusive [`Text`{.idl}](#text){#ref-for-text④⓪
link-type="idl"} nodes]{#contiguous-exclusive-text-nodes .dfn
.dfn-paneled dfn-type="dfn" export=""} of a
[node](#concept-node){#ref-for-concept-node①⑥⑥ link-type="dfn"}
`node`{.variable} are `node`{.variable}, `node`{.variable}'s [previous
sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling⑨
link-type="dfn"} [exclusive `Text`
node](#exclusive-text-node){#ref-for-exclusive-text-node⑧
link-type="dfn"}, if any, and its [contiguous exclusive `Text`
nodes](#contiguous-exclusive-text-nodes){#ref-for-contiguous-exclusive-text-nodes③
link-type="dfn"}, and `node`{.variable}'s [next
sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①④
link-type="dfn"} [exclusive `Text`
node](#exclusive-text-node){#ref-for-exclusive-text-node⑨
link-type="dfn"}, if any, and its [contiguous exclusive `Text`
nodes](#contiguous-exclusive-text-nodes){#ref-for-contiguous-exclusive-text-nodes④
link-type="dfn"}, avoiding any duplicates.

The [child text content]{#concept-child-text-content .dfn .dfn-paneled
dfn-type="dfn" export=""} of a
[node](#concept-node){#ref-for-concept-node①⑥⑦ link-type="dfn"}
`node`{.variable} is the
[concatenation](https://infra.spec.whatwg.org/#string-concatenate){#ref-for-string-concatenate②
link-type="dfn"} of the
[data](#concept-cd-data){#ref-for-concept-cd-data③③ link-type="dfn"} of
all the [`Text`{.idl}](#text){#ref-for-text④① link-type="idl"}
[node](#concept-node){#ref-for-concept-node①⑥⑧ link-type="dfn"}
[children](#concept-tree-child){#ref-for-concept-tree-child⑥③
link-type="dfn"} of `node`{.variable}, in [tree
order](#concept-tree-order){#ref-for-concept-tree-order②⑤
link-type="dfn"}.

The [descendant text content]{#concept-descendant-text-content .dfn
.dfn-paneled dfn-type="dfn" export=""} of a
[node](#concept-node){#ref-for-concept-node①⑥⑨ link-type="dfn"}
`node`{.variable} is the
[concatenation](https://infra.spec.whatwg.org/#string-concatenate){#ref-for-string-concatenate③
link-type="dfn"} of the
[data](#concept-cd-data){#ref-for-concept-cd-data③④ link-type="dfn"} of
all the [`Text`{.idl}](#text){#ref-for-text④② link-type="idl"}
[node](#concept-node){#ref-for-concept-node①⑦⓪ link-type="dfn"}
[descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant③①
link-type="dfn"} of `node`{.variable}, in [tree
order](#concept-tree-order){#ref-for-concept-tree-order②⑥
link-type="dfn"}.

------------------------------------------------------------------------

The [`new Text(``data`{.variable}`)`]{#dom-text-text .dfn .dfn-paneled
.idl-code dfn-for="Text" dfn-type="constructor" export=""
lt="Text(data)|constructor(data)|Text()|constructor()"} constructor
steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③④
link-type="dfn"}'s [data](#concept-cd-data){#ref-for-concept-cd-data③⑤
link-type="dfn"} to `data`{.variable} and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③⑤
link-type="dfn"}'s [node
document](#concept-node-document){#ref-for-concept-node-document⑥⑨
link-type="dfn"} to [current global
object](https://html.spec.whatwg.org/multipage/webappapis.html#current-global-object){#ref-for-current-global-object②
link-type="dfn"}'s [associated
`Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window③
link-type="dfn"}.

To [split]{#concept-text-split .dfn .dfn-paneled dfn-type="dfn"
export="" lt="split a Text node"} a
[`Text`{.idl}](#text){#ref-for-text④③ link-type="idl"}
[node](#concept-node){#ref-for-concept-node①⑦① link-type="dfn"}
`node`{.variable} with offset `offset`{.variable}, run these steps:

1.  Let `length`{.variable} be `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length①⓪
    link-type="dfn"}.

2.  If `offset`{.variable} is greater than `length`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦⓪
    link-type="dfn"} an
    \"[`IndexSizeError`{.idl}](https://webidl.spec.whatwg.org/#indexsizeerror){#ref-for-indexsizeerror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧④
    link-type="idl"}.

3.  Let `count`{.variable} be `length`{.variable} minus
    `offset`{.variable}.

4.  Let `new data`{.variable} be the result of [substringing
    data](#concept-cd-substring){#ref-for-concept-cd-substring①
    link-type="dfn"} with node `node`{.variable}, offset
    `offset`{.variable}, and count `count`{.variable}.

5.  Let `new node`{.variable} be a new
    [`Text`{.idl}](#text){#ref-for-text④④ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node①⑦② link-type="dfn"},
    with the same [node
    document](#concept-node-document){#ref-for-concept-node-document⑦⓪
    link-type="dfn"} as `node`{.variable}. Set `new node`{.variable}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data③⑥ link-type="dfn"}
    to `new data`{.variable}.

6.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent③⑨
    link-type="dfn"}.

7.  If `parent`{.variable} is not null:

    1.  [Insert](#concept-node-insert){#ref-for-concept-node-insert⑨
        link-type="dfn"} `new node`{.variable} into `parent`{.variable}
        before `node`{.variable}'s [next
        sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①⑤
        link-type="dfn"}.

    2.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range①③
        link-type="dfn"} whose [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node⑧
        link-type="dfn"} is `node`{.variable} and [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①①
        link-type="dfn"} is greater than `offset`{.variable}, set its
        [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node⑨
        link-type="dfn"} to `new node`{.variable} and decrease its
        [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①②
        link-type="dfn"} by `offset`{.variable}.

    3.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range①④
        link-type="dfn"} whose [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node⑧
        link-type="dfn"} is `node`{.variable} and [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①①
        link-type="dfn"} is greater than `offset`{.variable}, set its
        [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node⑨
        link-type="dfn"} to `new node`{.variable} and decrease its [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①②
        link-type="dfn"} by `offset`{.variable}.

    4.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range①⑤
        link-type="dfn"} whose [start
        node](#concept-range-start-node){#ref-for-concept-range-start-node①⓪
        link-type="dfn"} is `parent`{.variable} and [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①③
        link-type="dfn"} is equal to the
        [index](#concept-tree-index){#ref-for-concept-tree-index⑨
        link-type="dfn"} of `node`{.variable} plus 1, increase its
        [start
        offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①④
        link-type="dfn"} by 1.

    5.  For each [live
        range](#concept-live-range){#ref-for-concept-live-range①⑥
        link-type="dfn"} whose [end
        node](#concept-range-end-node){#ref-for-concept-range-end-node①⓪
        link-type="dfn"} is `parent`{.variable} and [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①③
        link-type="dfn"} is equal to the
        [index](#concept-tree-index){#ref-for-concept-tree-index①⓪
        link-type="dfn"} of `node`{.variable} plus 1, increase its [end
        offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①④
        link-type="dfn"} by 1.

8.  [Replace data](#concept-cd-replace){#ref-for-concept-cd-replace⑨
    link-type="dfn"} with node `node`{.variable}, offset
    `offset`{.variable}, count `count`{.variable}, and data the empty
    string.

9.  Return `new node`{.variable}.

The [`splitText(``offset`{.variable}`)`]{#dom-text-splittext .dfn
.dfn-paneled .idl-code dfn-for="Text" dfn-type="method" export=""}
method steps are to
[split](#concept-text-split){#ref-for-concept-text-split
link-type="dfn"}
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③⑥
link-type="dfn"} with offset `offset`{.variable}.

The [`wholeText`]{#dom-text-wholetext .dfn .dfn-paneled .idl-code
dfn-for="Text" dfn-type="attribute" export=""} getter steps are to
return the
[concatenation](https://infra.spec.whatwg.org/#string-concatenate){#ref-for-string-concatenate④
link-type="dfn"} of the
[data](#concept-cd-data){#ref-for-concept-cd-data③⑦ link-type="dfn"} of
the [contiguous `Text`
nodes](#contiguous-text-nodes){#ref-for-contiguous-text-nodes②
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③⑦
link-type="dfn"}, in [tree
order](#concept-tree-order){#ref-for-concept-tree-order②⑦
link-type="dfn"}.

### [4.12. ]{.secno}[Interface [`CDATASection`{.idl}](#cdatasection){#ref-for-cdatasection①⓪ link-type="idl"}]{.content}[](#interface-cdatasection){.self-link} {#interface-cdatasection .heading .settled level="4.12"}

``` {.idl .highlight .def}
[Exposed=Window]
interface CDATASection : Text {
};
```

### [4.13. ]{.secno}[Interface [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①⑤ link-type="idl"}]{.content}[](#interface-processinginstruction){.self-link} {#interface-processinginstruction .heading .settled level="4.13"}

``` {.idl .highlight .def}
[Exposed=Window]
interface ProcessingInstruction : CharacterData {
  readonly attribute DOMString target;
};
```

[`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①⑥
link-type="idl"} [nodes](#concept-node){#ref-for-concept-node①⑦③
link-type="dfn"} have an associated [target]{#concept-pi-target .dfn
.dfn-paneled dfn-for="ProcessingInstruction" dfn-type="dfn" export=""}.

The [`target`]{#dom-processinginstruction-target .dfn .dfn-paneled
.idl-code dfn-for="ProcessingInstruction" dfn-type="attribute"
export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③⑧
link-type="dfn"}'s
[target](#concept-pi-target){#ref-for-concept-pi-target⑥
link-type="dfn"}.

### [4.14. ]{.secno}[Interface [`Comment`{.idl}](#comment){#ref-for-comment①⑤ link-type="idl"}]{.content}[](#interface-comment){.self-link} {#interface-comment .heading .settled level="4.14"}

``` {.idl .highlight .def}
[Exposed=Window]
interface Comment : CharacterData {
  constructor(optional DOMString data = "");
};
```

`comment`{.variable}` = new `[`Comment([``data`{.variable}` = ""])`](#dom-comment-comment){#ref-for-dom-comment-comment① .idl-code link-type="constructor"}
:   Returns a new [`Comment`{.idl}](#comment){#ref-for-comment①⑥
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑦④
    link-type="dfn"} whose
    [data](#concept-cd-data){#ref-for-concept-cd-data③⑧ link-type="dfn"}
    is `data`{.variable}.

The [`new Comment(``data`{.variable}`)`]{#dom-comment-comment .dfn
.dfn-paneled .idl-code dfn-for="Comment" dfn-type="constructor"
export="" lt="Comment(data)|constructor(data)|Comment()|constructor()"}
constructor steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②③⑨
link-type="dfn"}'s [data](#concept-cd-data){#ref-for-concept-cd-data③⑨
link-type="dfn"} to `data`{.variable} and
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④⓪
link-type="dfn"}'s [node
document](#concept-node-document){#ref-for-concept-node-document⑦①
link-type="dfn"} to [current global
object](https://html.spec.whatwg.org/multipage/webappapis.html#current-global-object){#ref-for-current-global-object③
link-type="dfn"}'s [associated
`Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window④
link-type="dfn"}.

## [5. ]{.secno}[Ranges]{.content}[](#ranges){.self-link} {#ranges .heading .settled level="5"}

### [5.1. ]{.secno}[Introduction to \"DOM Ranges\"]{.content}[](#introduction-to-dom-ranges){.self-link} {#introduction-to-dom-ranges .heading .settled level="5.1"}

[`StaticRange`{.idl}](#staticrange){#ref-for-staticrange
link-type="idl"} and [`Range`{.idl}](#range){#ref-for-range①
link-type="idl"} objects
([ranges](#concept-range){#ref-for-concept-range link-type="dfn"})
represent a sequence of content within a [node
tree](#concept-node-tree){#ref-for-concept-node-tree②⑤ link-type="dfn"}.
Each [range](#concept-range){#ref-for-concept-range① link-type="dfn"}
has a [start](#concept-range-start){#ref-for-concept-range-start①
link-type="dfn"} and an
[end](#concept-range-end){#ref-for-concept-range-end① link-type="dfn"}
which are [boundary points](#concept-range-bp){#ref-for-concept-range-bp
link-type="dfn"}. A [boundary
point](#concept-range-bp){#ref-for-concept-range-bp① link-type="dfn"} is
a [tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple
link-type="dfn"} consisting of a
[node](#boundary-point-node){#ref-for-boundary-point-node②
link-type="dfn"} and an
[offset](#concept-range-bp-offset){#ref-for-concept-range-bp-offset
link-type="dfn"}. So in other words, a
[range](#concept-range){#ref-for-concept-range② link-type="dfn"}
represents a piece of content within a [node
tree](#concept-node-tree){#ref-for-concept-node-tree②⑥ link-type="dfn"}
between two [boundary
points](#concept-range-bp){#ref-for-concept-range-bp② link-type="dfn"}.

[Ranges](#concept-range){#ref-for-concept-range③ link-type="dfn"} are
frequently used in editing for selecting and copying content.

- [`Element`{.idl}](#element){#ref-for-element⑤④ link-type="idl"}: `p`
  - [`Element`{.idl}](#element){#ref-for-element⑤⑤ link-type="idl"}:
    `<img src="insanity-wolf" alt="Little-endian BOM; decode as big-endian!">`{.lang-markup
    .highlight}
  - [`Text`{.idl}](#text){#ref-for-text④⑥ link-type="idl"}:  CSS 2.1
    syndata is 
  - [`Element`{.idl}](#element){#ref-for-element⑤⑥ link-type="idl"}:
    `<em>`{.lang-markup .highlight}
    - [`Text`{.idl}](#text){#ref-for-text④⑦ link-type="idl"}: awesome
  - [`Text`{.idl}](#text){#ref-for-text④⑧ link-type="idl"}: !

In the [node tree](#concept-node-tree){#ref-for-concept-node-tree②⑦
link-type="dfn"} above, a
[range](#concept-range){#ref-for-concept-range④ link-type="dfn"} can be
used to represent the sequence "syndata is awes". Assuming
`p`{.variable} is assigned to the `p`
[element](#concept-element){#ref-for-concept-element①⓪⑧
link-type="dfn"}, and `em`{.variable} to the `em`
[element](#concept-element){#ref-for-concept-element①⓪⑨
link-type="dfn"}, this would be done as follows:

``` {.lang-javascript .highlight}
var range = new Range(),
    firstText = p.childNodes[1],
    secondText = em.firstChild
range.setStart(firstText, 9) // do not forget the leading space
range.setEnd(secondText, 4)
// range now stringifies to the aforementioned quote
```

[Attributes](#concept-attribute){#ref-for-concept-attribute⑥⑥
link-type="dfn"} such as `src` and `alt` in the [node
tree](#concept-node-tree){#ref-for-concept-node-tree②⑧ link-type="dfn"}
above cannot be represented by a
[range](#concept-range){#ref-for-concept-range⑤ link-type="dfn"}.
[Ranges](#concept-range){#ref-for-concept-range⑥ link-type="dfn"} are
only useful for [nodes](#concept-node){#ref-for-concept-node①⑦⑤
link-type="dfn"}.

[`Range`{.idl}](#range){#ref-for-range② link-type="idl"} objects, unlike
[`StaticRange`{.idl}](#staticrange){#ref-for-staticrange①
link-type="idl"} objects, are affected by mutations to the [node
tree](#concept-node-tree){#ref-for-concept-node-tree②⑨ link-type="dfn"}.
Therefore they are also known as [live
ranges](#concept-live-range){#ref-for-concept-live-range①⑦
link-type="dfn"}. Such mutations will not invalidate them and will try
to ensure that it still represents the same piece of content.
Necessarily, a [live
range](#concept-live-range){#ref-for-concept-live-range①⑧
link-type="dfn"} might itself be modified as part of the mutation to the
[node tree](#concept-node-tree){#ref-for-concept-node-tree③⓪
link-type="dfn"} when, e.g., part of the content it represents is
mutated.

See the [insert](#concept-node-insert){#ref-for-concept-node-insert①⓪
link-type="dfn"} and
[remove](#concept-node-remove){#ref-for-concept-node-remove①②
link-type="dfn"} algorithms, the
[`normalize()`{.idl}](#dom-node-normalize){#ref-for-dom-node-normalize②
link-type="idl"} method, and the [replace
data](#concept-cd-replace){#ref-for-concept-cd-replace①⓪
link-type="dfn"} and
[split](#concept-text-split){#ref-for-concept-text-split①
link-type="dfn"} algorithms for details.

Updating [live
ranges](#concept-live-range){#ref-for-concept-live-range①⑨
link-type="dfn"} in response to [node
tree](#concept-node-tree){#ref-for-concept-node-tree③① link-type="dfn"}
mutations can be expensive. For every [node
tree](#concept-node-tree){#ref-for-concept-node-tree③② link-type="dfn"}
change, all affected [`Range`{.idl}](#range){#ref-for-range③
link-type="idl"} objects need to be updated. Even if the application is
uninterested in some [live
ranges](#concept-live-range){#ref-for-concept-live-range②⓪
link-type="dfn"}, it still has to pay the cost of keeping them
up-to-date when a mutation occurs.

A [`StaticRange`{.idl}](#staticrange){#ref-for-staticrange②
link-type="idl"} object is a lightweight
[range](#concept-range){#ref-for-concept-range⑦ link-type="dfn"} that
does not update when the [node
tree](#concept-node-tree){#ref-for-concept-node-tree③③ link-type="dfn"}
mutates. It is therefore not subject to the same maintenance cost as
[live ranges](#concept-live-range){#ref-for-concept-live-range②①
link-type="dfn"}.

### [5.2. ]{.secno}[Boundary points]{.content}[](#boundary-points){.self-link} {#boundary-points .heading .settled level="5.2"}

A [boundary point]{#concept-range-bp .dfn .dfn-paneled dfn-type="dfn"
export=""} is a
[tuple](https://infra.spec.whatwg.org/#tuple){#ref-for-tuple①
link-type="dfn"} consisting of a [node]{#boundary-point-node .dfn
.dfn-paneled dfn-for="boundary point" dfn-type="dfn" export=""} (a
[node](#concept-node){#ref-for-concept-node①⑦⑥ link-type="dfn"}) and an
[offset]{#concept-range-bp-offset .dfn .dfn-paneled
dfn-for="boundary point" dfn-type="dfn" export=""} (a non-negative
integer).

A correct [boundary point](#concept-range-bp){#ref-for-concept-range-bp③
link-type="dfn"}'s
[offset](#concept-range-bp-offset){#ref-for-concept-range-bp-offset①
link-type="dfn"} will be between 0 and the [boundary
point](#concept-range-bp){#ref-for-concept-range-bp④ link-type="dfn"}'s
[node](#boundary-point-node){#ref-for-boundary-point-node③
link-type="dfn"}'s
[length](#concept-node-length){#ref-for-concept-node-length①①
link-type="dfn"}, inclusive.

The [position]{#concept-range-bp-position .dfn .dfn-paneled
dfn-for="boundary point" dfn-type="dfn" export=""} of a [boundary
point](#concept-range-bp){#ref-for-concept-range-bp⑤ link-type="dfn"}
(`nodeA`{.variable}, `offsetA`{.variable}) relative to a [boundary
point](#concept-range-bp){#ref-for-concept-range-bp⑥ link-type="dfn"}
(`nodeB`{.variable}, `offsetB`{.variable}) is
[before]{#concept-range-bp-before .dfn .dfn-paneled
dfn-for="boundary point" dfn-type="dfn" export=""},
[equal]{#concept-range-bp-equal .dfn .dfn-paneled
dfn-for="boundary point" dfn-type="dfn" export=""}, or
[after]{#concept-range-bp-after .dfn .dfn-paneled
dfn-for="boundary point" dfn-type="dfn" export=""}, as returned by these
steps:

1.  Assert: `nodeA`{.variable} and `nodeB`{.variable} have the same
    [root](#concept-tree-root){#ref-for-concept-tree-root④③
    link-type="dfn"}.

2.  If `nodeA`{.variable} is `nodeB`{.variable}, then return
    [equal](#concept-range-bp-equal){#ref-for-concept-range-bp-equal
    link-type="dfn"} if `offsetA`{.variable} is `offsetB`{.variable},
    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before
    link-type="dfn"} if `offsetA`{.variable} is less than
    `offsetB`{.variable}, and
    [after](#concept-range-bp-after){#ref-for-concept-range-bp-after
    link-type="dfn"} if `offsetA`{.variable} is greater than
    `offsetB`{.variable}.

3.  If `nodeA`{.variable} is
    [following](#concept-tree-following){#ref-for-concept-tree-following①②
    link-type="dfn"} `nodeB`{.variable}, then if the
    [position](#concept-range-bp-position){#ref-for-concept-range-bp-position
    link-type="dfn"} of (`nodeB`{.variable}, `offsetB`{.variable})
    relative to (`nodeA`{.variable}, `offsetA`{.variable}) is
    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before①
    link-type="dfn"}, return
    [after](#concept-range-bp-after){#ref-for-concept-range-bp-after①
    link-type="dfn"}, and if it is
    [after](#concept-range-bp-after){#ref-for-concept-range-bp-after②
    link-type="dfn"}, return
    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before②
    link-type="dfn"}.

4.  If `nodeA`{.variable} is an
    [ancestor](#concept-tree-ancestor){#ref-for-concept-tree-ancestor⑤
    link-type="dfn"} of `nodeB`{.variable}:

    1.  Let `child`{.variable} be `nodeB`{.variable}.

    2.  While `child`{.variable} is not a
        [child](#concept-tree-child){#ref-for-concept-tree-child⑥④
        link-type="dfn"} of `nodeA`{.variable}, set `child`{.variable}
        to its
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent④⓪
        link-type="dfn"}.

    3.  If `child`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index①①
        link-type="dfn"} is less than `offsetA`{.variable}, then return
        [after](#concept-range-bp-after){#ref-for-concept-range-bp-after③
        link-type="dfn"}.

5.  Return
    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before③
    link-type="dfn"}.

### [5.3. ]{.secno}[Interface [`AbstractRange`{.idl}](#abstractrange){#ref-for-abstractrange link-type="idl"}]{.content}[](#interface-abstractrange){.self-link} {#interface-abstractrange .heading .settled level="5.3"}

``` {.idl .highlight .def}
[Exposed=Window]
interface AbstractRange {
  readonly attribute Node startContainer;
  readonly attribute unsigned long startOffset;
  readonly attribute Node endContainer;
  readonly attribute unsigned long endOffset;
  readonly attribute boolean collapsed;
};
```

Objects implementing the
[`AbstractRange`{.idl}](#abstractrange){#ref-for-abstractrange①
link-type="idl"} interface are known as [ranges]{#concept-range .dfn
.dfn-paneled dfn-type="dfn" export="" lt="range"}.

A [range](#concept-range){#ref-for-concept-range⑧ link-type="dfn"} has
two associated [boundary
points](#concept-range-bp){#ref-for-concept-range-bp⑦ link-type="dfn"}
--- a [start]{#concept-range-start .dfn .dfn-paneled dfn-for="range"
dfn-type="dfn" export=""} and [end]{#concept-range-end .dfn .dfn-paneled
dfn-for="range" dfn-type="dfn" export=""}.

For convenience, a [range](#concept-range){#ref-for-concept-range⑨
link-type="dfn"}'s [start node]{#concept-range-start-node .dfn
.dfn-paneled dfn-for="range" dfn-type="dfn" export=""} is its
[start](#concept-range-start){#ref-for-concept-range-start②
link-type="dfn"}'s
[node](#boundary-point-node){#ref-for-boundary-point-node④
link-type="dfn"}, its [start offset]{#concept-range-start-offset .dfn
.dfn-paneled dfn-for="range" dfn-type="dfn" export=""} is its
[start](#concept-range-start){#ref-for-concept-range-start③
link-type="dfn"}'s
[offset](#concept-range-bp-offset){#ref-for-concept-range-bp-offset②
link-type="dfn"}, its [end node]{#concept-range-end-node .dfn
.dfn-paneled dfn-for="range" dfn-type="dfn" export=""} is its
[end](#concept-range-end){#ref-for-concept-range-end② link-type="dfn"}'s
[node](#boundary-point-node){#ref-for-boundary-point-node⑤
link-type="dfn"}, and its [end offset]{#concept-range-end-offset .dfn
.dfn-paneled dfn-for="range" dfn-type="dfn" export=""} is its
[end](#concept-range-end){#ref-for-concept-range-end③ link-type="dfn"}'s
[offset](#concept-range-bp-offset){#ref-for-concept-range-bp-offset③
link-type="dfn"}.

A [range](#concept-range){#ref-for-concept-range①⓪ link-type="dfn"} is
[collapsed]{#range-collapsed .dfn .dfn-paneled dfn-for="range"
dfn-type="dfn" export=""} if its [start
node](#concept-range-start-node){#ref-for-concept-range-start-node①①
link-type="dfn"} is its [end
node](#concept-range-end-node){#ref-for-concept-range-end-node①①
link-type="dfn"} and its [start
offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①⑤
link-type="dfn"} is its [end
offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①⑤
link-type="dfn"}.

`node`{.variable}` = ``range`{.variable}` . `[`startContainer`](#dom-range-startcontainer){#ref-for-dom-range-startcontainer① .idl-code link-type="attribute"}
:   Returns `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node①②
    link-type="dfn"}.

`offset`{.variable}` = ``range`{.variable}` . `[`startOffset`](#dom-range-startoffset){#ref-for-dom-range-startoffset① .idl-code link-type="attribute"}
:   Returns `range`{.variable}'s [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①⑥
    link-type="dfn"}.

`node`{.variable}` = ``range`{.variable}` . `[`endContainer`](#dom-range-endcontainer){#ref-for-dom-range-endcontainer① .idl-code link-type="attribute"}
:   Returns `range`{.variable}'s [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node①②
    link-type="dfn"}.

`offset`{.variable}` = ``range`{.variable}` . `[`endOffset`](#dom-range-endoffset){#ref-for-dom-range-endoffset① .idl-code link-type="attribute"}
:   Returns `range`{.variable}'s [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①⑥
    link-type="dfn"}.

`collapsed`{.variable}` = ``range`{.variable}` . `[`collapsed`](#dom-range-collapsed){#ref-for-dom-range-collapsed① .idl-code link-type="attribute"}
:   Returns true if `range`{.variable} is
    [collapsed](#range-collapsed){#ref-for-range-collapsed
    link-type="dfn"}; otherwise false.

The [`startContainer`]{#dom-range-startcontainer .dfn .dfn-paneled
.idl-code dfn-for="AbstractRange" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④①
link-type="dfn"}'s [start
node](#concept-range-start-node){#ref-for-concept-range-start-node①③
link-type="dfn"}.

The [`startOffset`]{#dom-range-startoffset .dfn .dfn-paneled .idl-code
dfn-for="AbstractRange" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④②
link-type="dfn"}'s [start
offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①⑦
link-type="dfn"}.

The [`endContainer`]{#dom-range-endcontainer .dfn .dfn-paneled .idl-code
dfn-for="AbstractRange" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④③
link-type="dfn"}'s [end
node](#concept-range-end-node){#ref-for-concept-range-end-node①③
link-type="dfn"}.

The [`endOffset`]{#dom-range-endoffset .dfn .dfn-paneled .idl-code
dfn-for="AbstractRange" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④④
link-type="dfn"}'s [end
offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①⑦
link-type="dfn"}.

The [`collapsed`]{#dom-range-collapsed .dfn .dfn-paneled .idl-code
dfn-for="AbstractRange" dfn-type="attribute" export=""} getter steps are
to return true if
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④⑤
link-type="dfn"} is
[collapsed](#range-collapsed){#ref-for-range-collapsed①
link-type="dfn"}; otherwise false.

### [5.4. ]{.secno}[Interface [`StaticRange`{.idl}](#staticrange){#ref-for-staticrange③ link-type="idl"}]{.content}[](#interface-staticrange){.self-link} {#interface-staticrange .heading .settled level="5.4"}

``` {.idl .highlight .def}
dictionary StaticRangeInit {
  required Node startContainer;
  required unsigned long startOffset;
  required Node endContainer;
  required unsigned long endOffset;
};

[Exposed=Window]
interface StaticRange : AbstractRange {
  constructor(StaticRangeInit init);
};
```

`staticRange`{.variable}` = new `[`StaticRange`](#dom-staticrange-staticrange){#ref-for-dom-staticrange-staticrange① .idl-code link-type="constructor"}`(init)`

:   Returns a new [range](#concept-range){#ref-for-concept-range①①
    link-type="dfn"} object that does not update when the [node
    tree](#concept-node-tree){#ref-for-concept-node-tree③④
    link-type="dfn"} mutates.

The
[`new StaticRange(``init`{.variable}`)`]{#dom-staticrange-staticrange
.dfn .dfn-paneled .idl-code dfn-for="StaticRange" dfn-type="constructor"
export="" lt="StaticRange(init)|constructor(init)"} constructor steps
are:

1.  If
    `init`{.variable}\[\"[`startContainer`{.idl}](#dom-staticrangeinit-startcontainer){#ref-for-dom-staticrangeinit-startcontainer
    link-type="idl"}\"\] or
    `init`{.variable}\[\"[`endContainer`{.idl}](#dom-staticrangeinit-endcontainer){#ref-for-dom-staticrangeinit-endcontainer
    link-type="idl"}\"\] is a
    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype②③
    link-type="idl"} or [`Attr`{.idl}](#attr){#ref-for-attr③⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑦⑦
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦①
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧⑤
    link-type="idl"}.

2.  Set [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④⑥
    link-type="dfn"}'s
    [start](#concept-range-start){#ref-for-concept-range-start④
    link-type="dfn"} to
    (`init`{.variable}\[\"[`startContainer`{.idl}](#dom-staticrangeinit-startcontainer){#ref-for-dom-staticrangeinit-startcontainer①
    link-type="idl"}\"\],
    `init`{.variable}\[\"[`startOffset`{.idl}](#dom-staticrangeinit-startoffset){#ref-for-dom-staticrangeinit-startoffset
    link-type="idl"}\"\]) and
    [end](#concept-range-end){#ref-for-concept-range-end④
    link-type="dfn"} to
    (`init`{.variable}\[\"[`endContainer`{.idl}](#dom-staticrangeinit-endcontainer){#ref-for-dom-staticrangeinit-endcontainer①
    link-type="idl"}\"\],
    `init`{.variable}\[\"[`endOffset`{.idl}](#dom-staticrangeinit-endoffset){#ref-for-dom-staticrangeinit-endoffset
    link-type="idl"}\"\]).

A [`StaticRange`{.idl}](#staticrange){#ref-for-staticrange④
link-type="idl"} is [valid]{#staticrange-valid .dfn .dfn-paneled
dfn-for="StaticRange" dfn-type="dfn" export=""} if all of the following
are true:

- Its [start](#concept-range-start){#ref-for-concept-range-start⑤
  link-type="dfn"} and
  [end](#concept-range-end){#ref-for-concept-range-end⑤ link-type="dfn"}
  are in the same [node
  tree](#concept-node-tree){#ref-for-concept-node-tree③⑤
  link-type="dfn"}.

- Its [start
  offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①⑧
  link-type="dfn"} is between 0 and its [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node①④
  link-type="dfn"}'s
  [length](#concept-node-length){#ref-for-concept-node-length①②
  link-type="dfn"}, inclusive.

- Its [end
  offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①⑧
  link-type="dfn"} is between 0 and its [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node①④
  link-type="dfn"}'s
  [length](#concept-node-length){#ref-for-concept-node-length①③
  link-type="dfn"}, inclusive.

- Its [start](#concept-range-start){#ref-for-concept-range-start⑥
  link-type="dfn"} is
  [before](#concept-range-bp-before){#ref-for-concept-range-bp-before④
  link-type="dfn"} or
  [equal](#concept-range-bp-equal){#ref-for-concept-range-bp-equal①
  link-type="dfn"} to its
  [end](#concept-range-end){#ref-for-concept-range-end⑥
  link-type="dfn"}.

### [5.5. ]{.secno}[Interface [`Range`{.idl}](#range){#ref-for-range④ link-type="idl"}]{.content}[](#interface-range){.self-link} {#interface-range .heading .settled level="5.5"}

``` {.idl .highlight .def}
[Exposed=Window]
interface Range : AbstractRange {
  constructor();

  readonly attribute Node commonAncestorContainer;

  undefined setStart(Node node, unsigned long offset);
  undefined setEnd(Node node, unsigned long offset);
  undefined setStartBefore(Node node);
  undefined setStartAfter(Node node);
  undefined setEndBefore(Node node);
  undefined setEndAfter(Node node);
  undefined collapse(optional boolean toStart = false);
  undefined selectNode(Node node);
  undefined selectNodeContents(Node node);

  const unsigned short START_TO_START = 0;
  const unsigned short START_TO_END = 1;
  const unsigned short END_TO_END = 2;
  const unsigned short END_TO_START = 3;
  short compareBoundaryPoints(unsigned short how, Range sourceRange);

  [CEReactions] undefined deleteContents();
  [CEReactions, NewObject] DocumentFragment extractContents();
  [CEReactions, NewObject] DocumentFragment cloneContents();
  [CEReactions] undefined insertNode(Node node);
  [CEReactions] undefined surroundContents(Node newParent);

  [NewObject] Range cloneRange();
  undefined detach();

  boolean isPointInRange(Node node, unsigned long offset);
  short comparePoint(Node node, unsigned long offset);

  boolean intersectsNode(Node node);

  stringifier;
};
```

Objects implementing the [`Range`{.idl}](#range){#ref-for-range⑦
link-type="idl"} interface are known as [live
ranges]{#concept-live-range .dfn .dfn-paneled dfn-type="dfn" export=""}.

Algorithms that modify a [tree](#concept-tree){#ref-for-concept-tree②②
link-type="dfn"} (in particular the
[insert](#concept-node-insert){#ref-for-concept-node-insert①①
link-type="dfn"},
[remove](#concept-node-remove){#ref-for-concept-node-remove①③
link-type="dfn"}, [move](#move){#ref-for-move③ link-type="dfn"},
[replace data](#concept-cd-replace){#ref-for-concept-cd-replace①①
link-type="dfn"}, and
[split](#concept-text-split){#ref-for-concept-text-split②
link-type="dfn"} algorithms) modify [live
ranges](#concept-live-range){#ref-for-concept-live-range②②
link-type="dfn"} associated with that
[tree](#concept-tree){#ref-for-concept-tree②③ link-type="dfn"}.

The [root]{#concept-range-root .dfn .dfn-paneled dfn-for="live range"
dfn-type="dfn" export=""} of a [live
range](#concept-live-range){#ref-for-concept-live-range②③
link-type="dfn"} is the
[root](#concept-tree-root){#ref-for-concept-tree-root④④ link-type="dfn"}
of its [start
node](#concept-range-start-node){#ref-for-concept-range-start-node①⑤
link-type="dfn"}.

A [node](#concept-node){#ref-for-concept-node①⑦⑧ link-type="dfn"}
`node`{.variable} is [contained]{#contained .dfn .dfn-paneled
dfn-for="live range" dfn-type="dfn" export=""} in a [live
range](#concept-live-range){#ref-for-concept-live-range②④
link-type="dfn"} `range`{.variable} if `node`{.variable}'s
[root](#concept-tree-root){#ref-for-concept-tree-root④⑤ link-type="dfn"}
is `range`{.variable}'s
[root](#concept-range-root){#ref-for-concept-range-root
link-type="dfn"}, and (`node`{.variable}, 0) is
[after](#concept-range-bp-after){#ref-for-concept-range-bp-after④
link-type="dfn"} `range`{.variable}'s
[start](#concept-range-start){#ref-for-concept-range-start⑦
link-type="dfn"}, and (`node`{.variable}, `node`{.variable}'s
[length](#concept-node-length){#ref-for-concept-node-length①④
link-type="dfn"}) is
[before](#concept-range-bp-before){#ref-for-concept-range-bp-before⑤
link-type="dfn"} `range`{.variable}'s
[end](#concept-range-end){#ref-for-concept-range-end⑦ link-type="dfn"}.

A [node](#concept-node){#ref-for-concept-node①⑦⑨ link-type="dfn"} is
[partially contained]{#partially-contained .dfn .dfn-paneled
dfn-for="live range" dfn-type="dfn" export=""} in a [live
range](#concept-live-range){#ref-for-concept-live-range②⑤
link-type="dfn"} if it's an [inclusive
ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor⑦
link-type="dfn"} of the [live
range](#concept-live-range){#ref-for-concept-live-range②⑥
link-type="dfn"}'s [start
node](#concept-range-start-node){#ref-for-concept-range-start-node①⑥
link-type="dfn"} but not its [end
node](#concept-range-end-node){#ref-for-concept-range-end-node①⑤
link-type="dfn"}, or vice versa.

::: {.note role="note"}
Some facts to better understand these definitions:

- The content that one would think of as being within the [live
  range](#concept-live-range){#ref-for-concept-live-range②⑦
  link-type="dfn"} consists of all
  [contained](#contained){#ref-for-contained link-type="dfn"}
  [nodes](#concept-node){#ref-for-concept-node①⑧⓪ link-type="dfn"}, plus
  possibly some of the contents of the [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node①⑦
  link-type="dfn"} and [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node①⑥
  link-type="dfn"} if those are
  [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②②
  link-type="idl"} [nodes](#concept-node){#ref-for-concept-node①⑧①
  link-type="dfn"}.

- The [nodes](#concept-node){#ref-for-concept-node①⑧② link-type="dfn"}
  that are contained in a [live
  range](#concept-live-range){#ref-for-concept-live-range②⑧
  link-type="dfn"} will generally not be contiguous, because the
  [parent](#concept-tree-parent){#ref-for-concept-tree-parent④①
  link-type="dfn"} of a [contained](#contained){#ref-for-contained①
  link-type="dfn"} [node](#concept-node){#ref-for-concept-node①⑧③
  link-type="dfn"} will not always be
  [contained](#contained){#ref-for-contained② link-type="dfn"}.

- However, the
  [descendants](#concept-tree-descendant){#ref-for-concept-tree-descendant③②
  link-type="dfn"} of a [contained](#contained){#ref-for-contained③
  link-type="dfn"} [node](#concept-node){#ref-for-concept-node①⑧④
  link-type="dfn"} are [contained](#contained){#ref-for-contained④
  link-type="dfn"}, and if two
  [siblings](#concept-tree-sibling){#ref-for-concept-tree-sibling①④
  link-type="dfn"} are [contained](#contained){#ref-for-contained⑤
  link-type="dfn"}, so are any
  [siblings](#concept-tree-sibling){#ref-for-concept-tree-sibling①⑤
  link-type="dfn"} that lie between them.

- The [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node①⑧
  link-type="dfn"} and [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node①⑦
  link-type="dfn"} of a [live
  range](#concept-live-range){#ref-for-concept-live-range②⑨
  link-type="dfn"} are never [contained](#contained){#ref-for-contained⑥
  link-type="dfn"} within it.

- The first [contained](#contained){#ref-for-contained⑦ link-type="dfn"}
  [node](#concept-node){#ref-for-concept-node①⑧⑤ link-type="dfn"} (if
  there are any) will always be after the [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node①⑨
  link-type="dfn"}, and the last
  [contained](#contained){#ref-for-contained⑧ link-type="dfn"}
  [node](#concept-node){#ref-for-concept-node①⑧⑥ link-type="dfn"} will
  always be equal to or before the [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node①⑧
  link-type="dfn"}'s last
  [descendant](#concept-tree-descendant){#ref-for-concept-tree-descendant③③
  link-type="dfn"}.

- There exists a [partially
  contained](#partially-contained){#ref-for-partially-contained
  link-type="dfn"} [node](#concept-node){#ref-for-concept-node①⑧⑦
  link-type="dfn"} if and only if the [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node②⓪
  link-type="dfn"} and [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node①⑨
  link-type="dfn"} are different.

- The
  [`commonAncestorContainer`{.idl}](#dom-range-commonancestorcontainer){#ref-for-dom-range-commonancestorcontainer①
  link-type="idl"} attribute value is neither
  [contained](#contained){#ref-for-contained⑨ link-type="dfn"} nor
  [partially
  contained](#partially-contained){#ref-for-partially-contained①
  link-type="dfn"}.

- If the [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node②①
  link-type="dfn"} is an
  [ancestor](#concept-tree-ancestor){#ref-for-concept-tree-ancestor⑥
  link-type="dfn"} of the [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node②⓪
  link-type="dfn"}, the common [inclusive
  ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor⑧
  link-type="dfn"} will be the [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node②②
  link-type="dfn"}. Exactly one of its
  [children](#concept-tree-child){#ref-for-concept-tree-child⑥⑤
  link-type="dfn"} will be [partially
  contained](#partially-contained){#ref-for-partially-contained②
  link-type="dfn"}, and a
  [child](#concept-tree-child){#ref-for-concept-tree-child⑥⑥
  link-type="dfn"} will be [contained](#contained){#ref-for-contained①⓪
  link-type="dfn"} if and only if it
  [precedes](#concept-tree-preceding){#ref-for-concept-tree-preceding①①
  link-type="dfn"} the [partially
  contained](#partially-contained){#ref-for-partially-contained③
  link-type="dfn"}
  [child](#concept-tree-child){#ref-for-concept-tree-child⑥⑦
  link-type="dfn"}. If the [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node②①
  link-type="dfn"} is an
  [ancestor](#concept-tree-ancestor){#ref-for-concept-tree-ancestor⑦
  link-type="dfn"} of the [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node②③
  link-type="dfn"}, the opposite holds.

- If the [start
  node](#concept-range-start-node){#ref-for-concept-range-start-node②④
  link-type="dfn"} is not an [inclusive
  ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor⑨
  link-type="dfn"} of the [end
  node](#concept-range-end-node){#ref-for-concept-range-end-node②②
  link-type="dfn"}, nor vice versa, the common [inclusive
  ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①⓪
  link-type="dfn"} will be distinct from both of them. Exactly two of
  its [children](#concept-tree-child){#ref-for-concept-tree-child⑥⑧
  link-type="dfn"} will be [partially
  contained](#partially-contained){#ref-for-partially-contained④
  link-type="dfn"}, and a
  [child](#concept-tree-child){#ref-for-concept-tree-child⑥⑨
  link-type="dfn"} will be contained if and only if it lies between
  those two.
:::

The [live range pre-remove steps]{#live-range-pre-remove-steps .dfn
.dfn-paneled dfn-type="dfn" noexport=""} given a
[node](#concept-node){#ref-for-concept-node①⑧⑧ link-type="dfn"}
`node`{.variable}, are as follows:

1.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent④②
    link-type="dfn"}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#ref-for-assert⑦
    link-type="dfn"}: `parent`{.variable} is not null.

3.  Let `index`{.variable} be `node`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index①②
    link-type="dfn"}.

4.  For each [live
    range](#concept-live-range){#ref-for-concept-live-range③⓪
    link-type="dfn"} whose [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node②⑤
    link-type="dfn"} is an [inclusive
    descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant⑤
    link-type="dfn"} of `node`{.variable}, set its
    [start](#concept-range-start){#ref-for-concept-range-start⑧
    link-type="dfn"} to (`parent`{.variable}, `index`{.variable}).

5.  For each [live
    range](#concept-live-range){#ref-for-concept-live-range③①
    link-type="dfn"} whose [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node②③
    link-type="dfn"} is an [inclusive
    descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant⑥
    link-type="dfn"} of `node`{.variable}, set its
    [end](#concept-range-end){#ref-for-concept-range-end⑧
    link-type="dfn"} to (`parent`{.variable}, `index`{.variable}).

6.  For each [live
    range](#concept-live-range){#ref-for-concept-live-range③②
    link-type="dfn"} whose [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node②⑥
    link-type="dfn"} is `parent`{.variable} and [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset①⑨
    link-type="dfn"} is greater than `index`{.variable}, decrease its
    [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②⓪
    link-type="dfn"} by 1.

7.  For each [live
    range](#concept-live-range){#ref-for-concept-live-range③③
    link-type="dfn"} whose [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node②④
    link-type="dfn"} is `parent`{.variable} and [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset①⑨
    link-type="dfn"} is greater than `index`{.variable}, decrease its
    [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset②⓪
    link-type="dfn"} by 1.

------------------------------------------------------------------------

`range`{.variable}` = new `[`Range()`](#dom-range-range){#ref-for-dom-range-range② .idl-code link-type="constructor"}
:   Returns a new [live
    range](#concept-live-range){#ref-for-concept-live-range③④
    link-type="dfn"}.

The [`new Range()`]{#dom-range-range .dfn .dfn-paneled .idl-code
dfn-for="Range" dfn-type="constructor" export=""
lt="Range()|constructor()"} constructor steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④⑦
link-type="dfn"}'s
[start](#concept-range-start){#ref-for-concept-range-start⑨
link-type="dfn"} and
[end](#concept-range-end){#ref-for-concept-range-end⑨ link-type="dfn"}
to ([current global
object](https://html.spec.whatwg.org/multipage/webappapis.html#current-global-object){#ref-for-current-global-object④
link-type="dfn"}'s [associated
`Document`](https://html.spec.whatwg.org/multipage/nav-history-apis.html#concept-document-window){#ref-for-concept-document-window⑤
link-type="dfn"}, 0).

------------------------------------------------------------------------

`container`{.variable} = `range`{.variable} . [`commonAncestorContainer`{.idl}](#dom-range-commonancestorcontainer){#ref-for-dom-range-commonancestorcontainer② link-type="idl"}
:   Returns the [node](#concept-node){#ref-for-concept-node①⑧⑨
    link-type="dfn"}, furthest away from the
    [document](#concept-document){#ref-for-concept-document⑤⑤
    link-type="dfn"}, that is an
    [ancestor](#concept-tree-ancestor){#ref-for-concept-tree-ancestor⑧
    link-type="dfn"} of both `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node②⑦
    link-type="dfn"} and [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node②⑤
    link-type="dfn"}.

The [`commonAncestorContainer`]{#dom-range-commonancestorcontainer .dfn
.dfn-paneled .idl-code dfn-for="Range" dfn-type="attribute" export=""}
getter steps are:

1.  Let `container`{.variable} be [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node②⑧
    link-type="dfn"}.
2.  While `container`{.variable} is not an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①①
    link-type="dfn"} of [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node②⑥
    link-type="dfn"}, let `container`{.variable} be
    `container`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent④③
    link-type="dfn"}.
3.  Return `container`{.variable}.

------------------------------------------------------------------------

To [set the start or end]{#concept-range-bp-set .dfn .dfn-paneled
dfn-for="Range" dfn-type="dfn" export="" lt="set the start|set the end"}
of a `range`{.variable} to a [boundary
point](#concept-range-bp){#ref-for-concept-range-bp⑧ link-type="dfn"}
(`node`{.variable}, `offset`{.variable}), run these steps:

1.  If `node`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype②③
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦②
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧⑥
    link-type="idl"}.

2.  If `offset`{.variable} is greater than `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length①⑤
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦③
    link-type="dfn"} an
    \"[`IndexSizeError`{.idl}](https://webidl.spec.whatwg.org/#indexsizeerror){#ref-for-indexsizeerror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧⑦
    link-type="idl"}.

3.  Let `bp`{.variable} be the [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp⑨
    link-type="dfn"} (`node`{.variable}, `offset`{.variable}).

4.  

    If these steps were invoked as \"set the start\"

    :   1.  If `range`{.variable}'s
            [root](#concept-range-root){#ref-for-concept-range-root①
            link-type="dfn"} is not equal to `node`{.variable}'s
            [root](#concept-tree-root){#ref-for-concept-tree-root④⑥
            link-type="dfn"}, or if `bp`{.variable} is
            [after](#concept-range-bp-after){#ref-for-concept-range-bp-after⑤
            link-type="dfn"} the `range`{.variable}'s
            [end](#concept-range-end){#ref-for-concept-range-end①⓪
            link-type="dfn"}, set `range`{.variable}'s
            [end](#concept-range-end){#ref-for-concept-range-end①①
            link-type="dfn"} to `bp`{.variable}.
        2.  Set `range`{.variable}'s
            [start](#concept-range-start){#ref-for-concept-range-start①⓪
            link-type="dfn"} to `bp`{.variable}.

    If these steps were invoked as \"set the end\"

    :   1.  If `range`{.variable}'s
            [root](#concept-range-root){#ref-for-concept-range-root②
            link-type="dfn"} is not equal to `node`{.variable}'s
            [root](#concept-tree-root){#ref-for-concept-tree-root④⑦
            link-type="dfn"}, or if `bp`{.variable} is
            [before](#concept-range-bp-before){#ref-for-concept-range-bp-before⑥
            link-type="dfn"} the `range`{.variable}'s
            [start](#concept-range-start){#ref-for-concept-range-start①①
            link-type="dfn"}, set `range`{.variable}'s
            [start](#concept-range-start){#ref-for-concept-range-start①②
            link-type="dfn"} to `bp`{.variable}.
        2.  Set `range`{.variable}'s
            [end](#concept-range-end){#ref-for-concept-range-end①②
            link-type="dfn"} to `bp`{.variable}.

The
[`setStart(``node`{.variable}`, ``offset`{.variable}`)`]{#dom-range-setstart
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are to [set the
start](#concept-range-bp-set){#ref-for-concept-range-bp-set
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④⑧
link-type="dfn"} to [boundary
point](#concept-range-bp){#ref-for-concept-range-bp①⓪ link-type="dfn"}
(`node`{.variable}, `offset`{.variable}).

The
[`setEnd(``node`{.variable}`, ``offset`{.variable}`)`]{#dom-range-setend
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are to [set the
end](#concept-range-bp-set){#ref-for-concept-range-bp-set①
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②④⑨
link-type="dfn"} to [boundary
point](#concept-range-bp){#ref-for-concept-range-bp①① link-type="dfn"}
(`node`{.variable}, `offset`{.variable}).

The [`setStartBefore(``node`{.variable}`)`]{#dom-range-setstartbefore
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent④④
    link-type="dfn"}.
2.  If `parent`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦④
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧⑧
    link-type="idl"}.
3.  [Set the
    start](#concept-range-bp-set){#ref-for-concept-range-bp-set②
    link-type="dfn"} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤⓪
    link-type="dfn"} to [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①②
    link-type="dfn"} (`parent`{.variable}, `node`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index①③
    link-type="dfn"}).

The [`setStartAfter(``node`{.variable}`)`]{#dom-range-setstartafter .dfn
.dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent④⑤
    link-type="dfn"}.

2.  If `parent`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦⑤
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror③
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑧⑨
    link-type="idl"}.

3.  [Set the
    start](#concept-range-bp-set){#ref-for-concept-range-bp-set③
    link-type="dfn"} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤①
    link-type="dfn"} to [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①③
    link-type="dfn"} (`parent`{.variable}, `node`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index①④
    link-type="dfn"} plus 1).

The [`setEndBefore(``node`{.variable}`)`]{#dom-range-setendbefore .dfn
.dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent④⑥
    link-type="dfn"}.
2.  If `parent`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦⑥
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨⓪
    link-type="idl"}.
3.  [Set the end](#concept-range-bp-set){#ref-for-concept-range-bp-set④
    link-type="dfn"} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤②
    link-type="dfn"} to [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①④
    link-type="dfn"} (`parent`{.variable}, `node`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index①⑤
    link-type="dfn"}).

The [`setEndAfter(``node`{.variable}`)`]{#dom-range-setendafter .dfn
.dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent④⑦
    link-type="dfn"}.

2.  If `parent`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦⑦
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨①
    link-type="idl"}.

3.  [Set the end](#concept-range-bp-set){#ref-for-concept-range-bp-set⑤
    link-type="dfn"} of
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤③
    link-type="dfn"} to [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①⑤
    link-type="dfn"} (`parent`{.variable}, `node`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index①⑥
    link-type="dfn"} plus 1).

The [`collapse(``toStart`{.variable}`)`]{#dom-range-collapse .dfn
.dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""
lt="collapse(toStart)|collapse()"} method steps are to, if
`toStart`{.variable} is true, set
[end](#concept-range-end){#ref-for-concept-range-end①③ link-type="dfn"}
to [start](#concept-range-start){#ref-for-concept-range-start①③
link-type="dfn"}; otherwise set
[start](#concept-range-start){#ref-for-concept-range-start①④
link-type="dfn"} to
[end](#concept-range-end){#ref-for-concept-range-end①④ link-type="dfn"}.

To [select]{#concept-range-select .dfn .dfn-paneled dfn-for="range"
dfn-type="dfn" export=""} a
[node](#concept-node){#ref-for-concept-node①⑨⓪ link-type="dfn"}
`node`{.variable} within a
[range](#concept-range){#ref-for-concept-range①② link-type="dfn"}
`range`{.variable}, run these steps:

1.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent④⑧
    link-type="dfn"}.

2.  If `parent`{.variable} is null, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦⑧
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨②
    link-type="idl"}.

3.  Let `index`{.variable} be `node`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index①⑦
    link-type="dfn"}.

4.  Set `range`{.variable}'s
    [start](#concept-range-start){#ref-for-concept-range-start①⑤
    link-type="dfn"} to [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①⑥
    link-type="dfn"} (`parent`{.variable}, `index`{.variable}).

5.  Set `range`{.variable}'s
    [end](#concept-range-end){#ref-for-concept-range-end①⑤
    link-type="dfn"} to [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①⑦
    link-type="dfn"} (`parent`{.variable}, `index`{.variable} plus 1).

The [`selectNode(``node`{.variable}`)`]{#dom-range-selectnode .dfn
.dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are to
[select](#concept-range-select){#ref-for-concept-range-select
link-type="dfn"} `node`{.variable} within
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤④
link-type="dfn"}.

The
[`selectNodeContents(``node`{.variable}`)`]{#dom-range-selectnodecontents
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  If `node`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype②④
    link-type="dfn"},
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑦⑨
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror⑦
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨③
    link-type="idl"}.

2.  Let `length`{.variable} be the
    [length](#concept-node-length){#ref-for-concept-node-length①⑥
    link-type="dfn"} of `node`{.variable}.

3.  Set [start](#concept-range-start){#ref-for-concept-range-start①⑥
    link-type="dfn"} to the [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①⑧
    link-type="dfn"} (`node`{.variable}, 0).

4.  Set [end](#concept-range-end){#ref-for-concept-range-end①⑥
    link-type="dfn"} to the [boundary
    point](#concept-range-bp){#ref-for-concept-range-bp①⑨
    link-type="dfn"} (`node`{.variable}, `length`{.variable}).

------------------------------------------------------------------------

The
[`compareBoundaryPoints(``how`{.variable}`, ``sourceRange`{.variable}`)`]{#dom-range-compareboundarypoints
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  If `how`{.variable} is not one of

    - [`START_TO_START`{.idl}](#dom-range-start_to_start){#ref-for-dom-range-start_to_start
      link-type="idl"},
    - [`START_TO_END`{.idl}](#dom-range-start_to_end){#ref-for-dom-range-start_to_end
      link-type="idl"},
    - [`END_TO_END`{.idl}](#dom-range-end_to_end){#ref-for-dom-range-end_to_end
      link-type="idl"}, and
    - [`END_TO_START`{.idl}](#dom-range-end_to_start){#ref-for-dom-range-end_to_start
      link-type="idl"},

    then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧⓪
    link-type="dfn"} a
    \"[`NotSupportedError`{.idl}](https://webidl.spec.whatwg.org/#notsupportederror){#ref-for-notsupportederror②⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨④
    link-type="idl"}.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤⑤
    link-type="dfn"}'s
    [root](#concept-range-root){#ref-for-concept-range-root③
    link-type="dfn"} is not the same as `sourceRange`{.variable}'s
    [root](#concept-range-root){#ref-for-concept-range-root④
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧①
    link-type="dfn"} a
    \"[`WrongDocumentError`{.idl}](https://webidl.spec.whatwg.org/#wrongdocumenterror){#ref-for-wrongdocumenterror
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨⑤
    link-type="idl"}.

3.  If `how`{.variable} is:

    [`START_TO_START`{.idl}](#dom-range-start_to_start){#ref-for-dom-range-start_to_start① link-type="idl"}:
    :   Let `this point`{.variable} be
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤⑥
        link-type="dfn"}'s
        [start](#concept-range-start){#ref-for-concept-range-start①⑦
        link-type="dfn"}. Let `other point`{.variable} be
        `sourceRange`{.variable}'s
        [start](#concept-range-start){#ref-for-concept-range-start①⑧
        link-type="dfn"}.

    [`START_TO_END`{.idl}](#dom-range-start_to_end){#ref-for-dom-range-start_to_end① link-type="idl"}:
    :   Let `this point`{.variable} be
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤⑦
        link-type="dfn"}'s
        [end](#concept-range-end){#ref-for-concept-range-end①⑦
        link-type="dfn"}. Let `other point`{.variable} be
        `sourceRange`{.variable}'s
        [start](#concept-range-start){#ref-for-concept-range-start①⑨
        link-type="dfn"}.

    [`END_TO_END`{.idl}](#dom-range-end_to_end){#ref-for-dom-range-end_to_end① link-type="idl"}:
    :   Let `this point`{.variable} be
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤⑧
        link-type="dfn"}'s
        [end](#concept-range-end){#ref-for-concept-range-end①⑧
        link-type="dfn"}. Let `other point`{.variable} be
        `sourceRange`{.variable}'s
        [end](#concept-range-end){#ref-for-concept-range-end①⑨
        link-type="dfn"}.

    [`END_TO_START`{.idl}](#dom-range-end_to_start){#ref-for-dom-range-end_to_start① link-type="idl"}:
    :   Let `this point`{.variable} be
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑤⑨
        link-type="dfn"}'s
        [start](#concept-range-start){#ref-for-concept-range-start②⓪
        link-type="dfn"}. Let `other point`{.variable} be
        `sourceRange`{.variable}'s
        [end](#concept-range-end){#ref-for-concept-range-end②⓪
        link-type="dfn"}.

4.  If the
    [position](#concept-range-bp-position){#ref-for-concept-range-bp-position①
    link-type="dfn"} of `this point`{.variable} relative to
    `other point`{.variable} is

    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before⑦ link-type="dfn"}
    :   Return −1.

    [equal](#concept-range-bp-equal){#ref-for-concept-range-bp-equal② link-type="dfn"}
    :   Return 0.

    [after](#concept-range-bp-after){#ref-for-concept-range-bp-after⑥ link-type="dfn"}
    :   Return 1.

The [`deleteContents()`]{#dom-range-deletecontents .dfn .dfn-paneled
.idl-code dfn-for="Range" dfn-type="method" export=""} method steps are:

1.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥⓪
    link-type="dfn"} is
    [collapsed](#range-collapsed){#ref-for-range-collapsed②
    link-type="dfn"}, then return.

2.  Let `original start node`{.variable},
    `original start offset`{.variable}, `original end node`{.variable},
    and `original end offset`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥①
    link-type="dfn"}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node②⑨
    link-type="dfn"}, [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②①
    link-type="dfn"}, [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node②⑦
    link-type="dfn"}, and [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset②①
    link-type="dfn"}, respectively.

3.  If `original start node`{.variable} is
    `original end node`{.variable} and it is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②③
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑨①
    link-type="dfn"}, then [replace
    data](#concept-cd-replace){#ref-for-concept-cd-replace①②
    link-type="dfn"} with node `original start node`{.variable}, offset
    `original start offset`{.variable}, count
    `original end offset`{.variable} minus
    `original start offset`{.variable}, and data the empty string, and
    then return.

4.  Let `nodes to remove`{.variable} be a list of all the
    [nodes](#concept-node){#ref-for-concept-node①⑨② link-type="dfn"}
    that are [contained](#contained){#ref-for-contained①①
    link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥②
    link-type="dfn"}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order②⑧
    link-type="dfn"}, omitting any
    [node](#concept-node){#ref-for-concept-node①⑨③ link-type="dfn"}
    whose [parent](#concept-tree-parent){#ref-for-concept-tree-parent④⑨
    link-type="dfn"} is also
    [contained](#contained){#ref-for-contained①② link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥③
    link-type="dfn"}.

5.  If `original start node`{.variable} is an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①②
    link-type="dfn"} of `original end node`{.variable}, set
    `new node`{.variable} to `original start node`{.variable} and
    `new offset`{.variable} to `original start offset`{.variable}.

6.  Otherwise:
    1.  Let `reference node`{.variable} equal
        `original start node`{.variable}.

    2.  While `reference node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤⓪
        link-type="dfn"} is not null and is not an [inclusive
        ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①③
        link-type="dfn"} of `original end node`{.variable}, set
        `reference node`{.variable} to its
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤①
        link-type="dfn"}.

    3.  Set `new node`{.variable} to the
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤②
        link-type="dfn"} of `reference node`{.variable}, and
        `new offset`{.variable} to one plus the
        [index](#concept-tree-index){#ref-for-concept-tree-index①⑧
        link-type="dfn"} of `reference node`{.variable}.

        If `reference node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤③
        link-type="dfn"} were null, it would be the
        [root](#concept-range-root){#ref-for-concept-range-root⑤
        link-type="dfn"} of
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥④
        link-type="dfn"}, so would be an [inclusive
        ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①④
        link-type="dfn"} of `original end node`{.variable}, and we could
        not reach this point.

7.  If `original start node`{.variable} is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②④
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑨④
    link-type="dfn"}, then [replace
    data](#concept-cd-replace){#ref-for-concept-cd-replace①③
    link-type="dfn"} with node `original start node`{.variable}, offset
    `original start offset`{.variable}, count
    `original start node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length①⑦
    link-type="dfn"} − `original start offset`{.variable}, data the
    empty string.

8.  For each `node`{.variable} in `nodes to remove`{.variable}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order②⑨
    link-type="dfn"},
    [remove](#concept-node-remove){#ref-for-concept-node-remove①④
    link-type="dfn"} `node`{.variable}.

9.  If `original end node`{.variable} is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②⑤
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑨⑤
    link-type="dfn"}, then [replace
    data](#concept-cd-replace){#ref-for-concept-cd-replace①④
    link-type="dfn"} with node `original end node`{.variable}, offset 0,
    count `original end offset`{.variable} and data the empty string.

10. Set [start](#concept-range-start){#ref-for-concept-range-start②①
    link-type="dfn"} and
    [end](#concept-range-end){#ref-for-concept-range-end②①
    link-type="dfn"} to (`new node`{.variable},
    `new offset`{.variable}).

To [extract]{#concept-range-extract .dfn .dfn-paneled
dfn-for="live range" dfn-type="dfn" export=""} a [live
range](#concept-live-range){#ref-for-concept-live-range③⑤
link-type="dfn"} `range`{.variable}, run these steps:

1.  Let `fragment`{.variable} be a new
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③⑥
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑨⑥
    link-type="dfn"} whose [node
    document](#concept-node-document){#ref-for-concept-node-document⑦②
    link-type="dfn"} is `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③⓪
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑦③
    link-type="dfn"}.

2.  If `range`{.variable} is
    [collapsed](#range-collapsed){#ref-for-range-collapsed③
    link-type="dfn"}, then return `fragment`{.variable}.

3.  Let `original start node`{.variable},
    `original start offset`{.variable}, `original end node`{.variable},
    and `original end offset`{.variable} be `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③①
    link-type="dfn"}, [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②②
    link-type="dfn"}, [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node②⑧
    link-type="dfn"}, and [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset②②
    link-type="dfn"}, respectively.

4.  If `original start node`{.variable} is
    `original end node`{.variable} and it is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②⑥
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑨⑦
    link-type="dfn"}:

    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone⑤
        link-type="dfn"} of `original start node`{.variable}.
    2.  Set the [data](#concept-cd-data){#ref-for-concept-cd-data④⓪
        link-type="dfn"} of `clone`{.variable} to the result of
        [substringing
        data](#concept-cd-substring){#ref-for-concept-cd-substring②
        link-type="dfn"} with node `original start node`{.variable},
        offset `original start offset`{.variable}, and count
        `original end offset`{.variable} minus
        `original start offset`{.variable}.
    3.  [Append](#concept-node-append){#ref-for-concept-node-append①②
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.
    4.  [Replace
        data](#concept-cd-replace){#ref-for-concept-cd-replace①⑤
        link-type="dfn"} with node `original start node`{.variable},
        offset `original start offset`{.variable}, count
        `original end offset`{.variable} minus
        `original start offset`{.variable}, and data the empty string.
    5.  Return `fragment`{.variable}.

5.  Let `common ancestor`{.variable} be
    `original start node`{.variable}.

6.  While `common ancestor`{.variable} is not an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①⑤
    link-type="dfn"} of `original end node`{.variable}, set
    `common ancestor`{.variable} to its own
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤④
    link-type="dfn"}.

7.  Let `first partially contained child`{.variable} be null.

8.  If `original start node`{.variable} is *not* an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①⑥
    link-type="dfn"} of `original end node`{.variable}, set
    `first partially contained child`{.variable} to the first
    [child](#concept-tree-child){#ref-for-concept-tree-child⑦⓪
    link-type="dfn"} of `common ancestor`{.variable} that is [partially
    contained](#partially-contained){#ref-for-partially-contained⑤
    link-type="dfn"} in `range`{.variable}.

9.  Let `last partially contained child`{.variable} be null.

10. If `original end node`{.variable} is *not* an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①⑦
    link-type="dfn"} of `original start node`{.variable}, set
    `last partially contained child`{.variable} to the last
    [child](#concept-tree-child){#ref-for-concept-tree-child⑦①
    link-type="dfn"} of `common ancestor`{.variable} that is [partially
    contained](#partially-contained){#ref-for-partially-contained⑥
    link-type="dfn"} in `range`{.variable}.

    These variable assignments do actually always make sense. For
    instance, if `original start node`{.variable} is not an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①⑧
    link-type="dfn"} of `original end node`{.variable},
    `original start node`{.variable} is itself [partially
    contained](#partially-contained){#ref-for-partially-contained⑦
    link-type="dfn"} in `range`{.variable}, and so are all its
    [ancestors](#concept-tree-ancestor){#ref-for-concept-tree-ancestor⑨
    link-type="dfn"} up until a
    [child](#concept-tree-child){#ref-for-concept-tree-child⑦②
    link-type="dfn"} of `common ancestor`{.variable}.
    `common ancestor`{.variable} cannot be
    `original start node`{.variable}, because it has to be an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor①⑨
    link-type="dfn"} of `original end node`{.variable}. The other case
    is similar. Also, notice that the two
    [children](#concept-tree-child){#ref-for-concept-tree-child⑦③
    link-type="dfn"} will never be equal if both are defined.

11. Let `contained children`{.variable} be a list of all
    [children](#concept-tree-child){#ref-for-concept-tree-child⑦④
    link-type="dfn"} of `common ancestor`{.variable} that are
    [contained](#contained){#ref-for-contained①③ link-type="dfn"} in
    `range`{.variable}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order③⓪
    link-type="dfn"}.

12. If any member of `contained children`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype②⑤
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧②
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨⑥
    link-type="idl"}.

    We do not have to worry about the first or last partially contained
    node, because a
    [doctype](#concept-doctype){#ref-for-concept-doctype②⑥
    link-type="dfn"} can never be partially contained. It cannot be a
    boundary point of a range, and it cannot be the ancestor of
    anything.

13. If `original start node`{.variable} is an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②⓪
    link-type="dfn"} of `original end node`{.variable}, set
    `new node`{.variable} to `original start node`{.variable} and
    `new offset`{.variable} to `original start offset`{.variable}.

14. Otherwise:
    1.  Let `reference node`{.variable} equal
        `original start node`{.variable}.

    2.  While `reference node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤⑤
        link-type="dfn"} is not null and is not an [inclusive
        ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②①
        link-type="dfn"} of `original end node`{.variable}, set
        `reference node`{.variable} to its
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤⑥
        link-type="dfn"}.

    3.  Set `new node`{.variable} to the
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤⑦
        link-type="dfn"} of `reference node`{.variable}, and
        `new offset`{.variable} to one plus
        `reference node`{.variable}'s
        [index](#concept-tree-index){#ref-for-concept-tree-index①⑨
        link-type="dfn"}.

        If `reference node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤⑧
        link-type="dfn"} is null, it would be the
        [root](#concept-range-root){#ref-for-concept-range-root⑥
        link-type="dfn"} of `range`{.variable}, so would be an
        [inclusive
        ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②②
        link-type="dfn"} of `original end node`{.variable}, and we could
        not reach this point.

15. If `first partially contained child`{.variable} is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②⑦
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑨⑧
    link-type="dfn"}:

    In this case, `first partially contained child`{.variable} is
    `original start node`{.variable}.

    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone⑥
        link-type="dfn"} of `original start node`{.variable}.

    2.  Set the [data](#concept-cd-data){#ref-for-concept-cd-data④①
        link-type="dfn"} of `clone`{.variable} to the result of
        [substringing
        data](#concept-cd-substring){#ref-for-concept-cd-substring③
        link-type="dfn"} with node `original start node`{.variable},
        offset `original start offset`{.variable}, and count
        `original start node`{.variable}'s
        [length](#concept-node-length){#ref-for-concept-node-length①⑧
        link-type="dfn"} − `original start offset`{.variable}.

    3.  [Append](#concept-node-append){#ref-for-concept-node-append①③
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

    4.  [Replace
        data](#concept-cd-replace){#ref-for-concept-cd-replace①⑥
        link-type="dfn"} with node `original start node`{.variable},
        offset `original start offset`{.variable}, count
        `original start node`{.variable}'s
        [length](#concept-node-length){#ref-for-concept-node-length①⑨
        link-type="dfn"} − `original start offset`{.variable}, and data
        the empty string.

16. Otherwise, if `first partially contained child`{.variable} is not
    null:
    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone⑦
        link-type="dfn"} of
        `first partially contained child`{.variable}.

    2.  [Append](#concept-node-append){#ref-for-concept-node-append①④
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

    3.  Let `subrange`{.variable} be a new [live
        range](#concept-live-range){#ref-for-concept-live-range③⑥
        link-type="dfn"} whose
        [start](#concept-range-start){#ref-for-concept-range-start②②
        link-type="dfn"} is (`original start node`{.variable},
        `original start offset`{.variable}) and whose
        [end](#concept-range-end){#ref-for-concept-range-end②②
        link-type="dfn"} is
        (`first partially contained child`{.variable},
        `first partially contained child`{.variable}'s
        [length](#concept-node-length){#ref-for-concept-node-length②⓪
        link-type="dfn"}).

    4.  Let `subfragment`{.variable} be the result of
        [extracting](#concept-range-extract){#ref-for-concept-range-extract
        link-type="dfn"} `subrange`{.variable}.

    5.  [Append](#concept-node-append){#ref-for-concept-node-append①⑤
        link-type="dfn"} `subfragment`{.variable} to `clone`{.variable}.

17. For each `contained child`{.variable} in
    `contained children`{.variable},
    [append](#concept-node-append){#ref-for-concept-node-append①⑥
    link-type="dfn"} `contained child`{.variable} to
    `fragment`{.variable}.

18. If `last partially contained child`{.variable} is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node①⑨⑨
    link-type="dfn"}:

    In this case, `last partially contained child`{.variable} is
    `original end node`{.variable}.

    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone⑧
        link-type="dfn"} of `original end node`{.variable}.
    2.  Set the [data](#concept-cd-data){#ref-for-concept-cd-data④②
        link-type="dfn"} of `clone`{.variable} to the result of
        [substringing
        data](#concept-cd-substring){#ref-for-concept-cd-substring④
        link-type="dfn"} with node `original end node`{.variable},
        offset 0, and count `original end offset`{.variable}.
    3.  [Append](#concept-node-append){#ref-for-concept-node-append①⑦
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.
    4.  [Replace
        data](#concept-cd-replace){#ref-for-concept-cd-replace①⑦
        link-type="dfn"} with node `original end node`{.variable},
        offset 0, count `original end offset`{.variable}, and data the
        empty string.

19. Otherwise, if `last partially contained child`{.variable} is not
    null:
    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone⑨
        link-type="dfn"} of `last partially contained child`{.variable}.

    2.  [Append](#concept-node-append){#ref-for-concept-node-append①⑧
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

    3.  Let `subrange`{.variable} be a new [live
        range](#concept-live-range){#ref-for-concept-live-range③⑦
        link-type="dfn"} whose
        [start](#concept-range-start){#ref-for-concept-range-start②③
        link-type="dfn"} is
        (`last partially contained child`{.variable}, 0) and whose
        [end](#concept-range-end){#ref-for-concept-range-end②③
        link-type="dfn"} is (`original end node`{.variable},
        `original end offset`{.variable}).

    4.  Let `subfragment`{.variable} be the result of
        [extracting](#concept-range-extract){#ref-for-concept-range-extract①
        link-type="dfn"} `subrange`{.variable}.

    5.  [Append](#concept-node-append){#ref-for-concept-node-append①⑨
        link-type="dfn"} `subfragment`{.variable} to `clone`{.variable}.

20. Set `range`{.variable}'s
    [start](#concept-range-start){#ref-for-concept-range-start②④
    link-type="dfn"} and
    [end](#concept-range-end){#ref-for-concept-range-end②④
    link-type="dfn"} to (`new node`{.variable},
    `new offset`{.variable}).

21. Return `fragment`{.variable}.

The [`extractContents()`]{#dom-range-extractcontents .dfn .dfn-paneled
.idl-code dfn-for="Range" dfn-type="method" export=""} method steps are
to return the result of
[extracting](#concept-range-extract){#ref-for-concept-range-extract②
link-type="dfn"}
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥⑤
link-type="dfn"}.

To [clone the contents]{#concept-range-clone .dfn .dfn-paneled
dfn-for="live range" dfn-type="dfn" export=""
lt="clone the contents|cloning the contents"} of a [live
range](#concept-live-range){#ref-for-concept-live-range③⑧
link-type="dfn"} `range`{.variable}, run these steps:

1.  Let `fragment`{.variable} be a new
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③⑦
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪⓪
    link-type="dfn"} whose [node
    document](#concept-node-document){#ref-for-concept-node-document⑦④
    link-type="dfn"} is `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③②
    link-type="dfn"}'s [node
    document](#concept-node-document){#ref-for-concept-node-document⑦⑤
    link-type="dfn"}.

2.  If `range`{.variable} is
    [collapsed](#range-collapsed){#ref-for-range-collapsed④
    link-type="dfn"}, then return `fragment`{.variable}.

3.  Let `original start node`{.variable},
    `original start offset`{.variable}, `original end node`{.variable},
    and `original end offset`{.variable} be `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③③
    link-type="dfn"}, [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②③
    link-type="dfn"}, [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node②⑨
    link-type="dfn"}, and [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset②③
    link-type="dfn"}, respectively.

4.  If `original start node`{.variable} is
    `original end node`{.variable} and it is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata②⑨
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪①
    link-type="dfn"}:

    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone①⓪
        link-type="dfn"} of `original start node`{.variable}.
    2.  Set the [data](#concept-cd-data){#ref-for-concept-cd-data④③
        link-type="dfn"} of `clone`{.variable} to the result of
        [substringing
        data](#concept-cd-substring){#ref-for-concept-cd-substring⑤
        link-type="dfn"} with node `original start node`{.variable},
        offset `original start offset`{.variable}, and count
        `original end offset`{.variable} minus
        `original start offset`{.variable}.
    3.  [Append](#concept-node-append){#ref-for-concept-node-append②⓪
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.
    4.  Return `fragment`{.variable}.

5.  Let `common ancestor`{.variable} be
    `original start node`{.variable}.

6.  While `common ancestor`{.variable} is not an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②③
    link-type="dfn"} of `original end node`{.variable}, set
    `common ancestor`{.variable} to its own
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑤⑨
    link-type="dfn"}.

7.  Let `first partially contained child`{.variable} be null.

8.  If `original start node`{.variable} is *not* an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②④
    link-type="dfn"} of `original end node`{.variable}, set
    `first partially contained child`{.variable} to the first
    [child](#concept-tree-child){#ref-for-concept-tree-child⑦⑤
    link-type="dfn"} of `common ancestor`{.variable} that is [partially
    contained](#partially-contained){#ref-for-partially-contained⑧
    link-type="dfn"} in `range`{.variable}.

9.  Let `last partially contained child`{.variable} be null.

10. If `original end node`{.variable} is *not* an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②⑤
    link-type="dfn"} of `original start node`{.variable}, set
    `last partially contained child`{.variable} to the last
    [child](#concept-tree-child){#ref-for-concept-tree-child⑦⑥
    link-type="dfn"} of `common ancestor`{.variable} that is [partially
    contained](#partially-contained){#ref-for-partially-contained⑨
    link-type="dfn"} in `range`{.variable}.

    These variable assignments do actually always make sense. For
    instance, if `original start node`{.variable} is not an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②⑥
    link-type="dfn"} of `original end node`{.variable},
    `original start node`{.variable} is itself [partially
    contained](#partially-contained){#ref-for-partially-contained①⓪
    link-type="dfn"} in `range`{.variable}, and so are all its
    [ancestors](#concept-tree-ancestor){#ref-for-concept-tree-ancestor①⓪
    link-type="dfn"} up until a
    [child](#concept-tree-child){#ref-for-concept-tree-child⑦⑦
    link-type="dfn"} of `common ancestor`{.variable}.
    `common ancestor`{.variable} cannot be
    `original start node`{.variable}, because it has to be an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②⑦
    link-type="dfn"} of `original end node`{.variable}. The other case
    is similar. Also, notice that the two
    [children](#concept-tree-child){#ref-for-concept-tree-child⑦⑧
    link-type="dfn"} will never be equal if both are defined.

11. Let `contained children`{.variable} be a list of all
    [children](#concept-tree-child){#ref-for-concept-tree-child⑦⑨
    link-type="dfn"} of `common ancestor`{.variable} that are
    [contained](#contained){#ref-for-contained①④ link-type="dfn"} in
    `range`{.variable}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order③①
    link-type="dfn"}.

12. If any member of `contained children`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype②⑦
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧③
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨⑦
    link-type="idl"}.

    We do not have to worry about the first or last partially contained
    node, because a
    [doctype](#concept-doctype){#ref-for-concept-doctype②⑧
    link-type="dfn"} can never be partially contained. It cannot be a
    boundary point of a range, and it cannot be the ancestor of
    anything.

13. If `first partially contained child`{.variable} is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata③⓪
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪②
    link-type="dfn"}:

    In this case, `first partially contained child`{.variable} is
    `original start node`{.variable}.

    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone①①
        link-type="dfn"} of `original start node`{.variable}.

    2.  Set the [data](#concept-cd-data){#ref-for-concept-cd-data④④
        link-type="dfn"} of `clone`{.variable} to the result of
        [substringing
        data](#concept-cd-substring){#ref-for-concept-cd-substring⑥
        link-type="dfn"} with node `original start node`{.variable},
        offset `original start offset`{.variable}, and count
        `original start node`{.variable}'s
        [length](#concept-node-length){#ref-for-concept-node-length②①
        link-type="dfn"} − `original start offset`{.variable}.

    3.  [Append](#concept-node-append){#ref-for-concept-node-append②①
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

14. Otherwise, if `first partially contained child`{.variable} is not
    null:
    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone①②
        link-type="dfn"} of
        `first partially contained child`{.variable}.

    2.  [Append](#concept-node-append){#ref-for-concept-node-append②②
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

    3.  Let `subrange`{.variable} be a new [live
        range](#concept-live-range){#ref-for-concept-live-range③⑨
        link-type="dfn"} whose
        [start](#concept-range-start){#ref-for-concept-range-start②⑤
        link-type="dfn"} is (`original start node`{.variable},
        `original start offset`{.variable}) and whose
        [end](#concept-range-end){#ref-for-concept-range-end②⑤
        link-type="dfn"} is
        (`first partially contained child`{.variable},
        `first partially contained child`{.variable}'s
        [length](#concept-node-length){#ref-for-concept-node-length②②
        link-type="dfn"}).

    4.  Let `subfragment`{.variable} be the result of [cloning the
        contents](#concept-range-clone){#ref-for-concept-range-clone
        link-type="dfn"} of `subrange`{.variable}.

    5.  [Append](#concept-node-append){#ref-for-concept-node-append②③
        link-type="dfn"} `subfragment`{.variable} to `clone`{.variable}.

15. For each `contained child`{.variable} in
    `contained children`{.variable}:
    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone①③
        link-type="dfn"} of `contained child`{.variable} with the *clone
        children flag* set.
    2.  [Append](#concept-node-append){#ref-for-concept-node-append②④
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

16. If `last partially contained child`{.variable} is a
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata③①
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪③
    link-type="dfn"}:

    In this case, `last partially contained child`{.variable} is
    `original end node`{.variable}.

    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone①④
        link-type="dfn"} of `original end node`{.variable}.
    2.  Set the [data](#concept-cd-data){#ref-for-concept-cd-data④⑤
        link-type="dfn"} of `clone`{.variable} to the result of
        [substringing
        data](#concept-cd-substring){#ref-for-concept-cd-substring⑦
        link-type="dfn"} with node `original end node`{.variable},
        offset 0, and count `original end offset`{.variable}.
    3.  [Append](#concept-node-append){#ref-for-concept-node-append②⑤
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

17. Otherwise, if `last partially contained child`{.variable} is not
    null:
    1.  Let `clone`{.variable} be a
        [clone](#concept-node-clone){#ref-for-concept-node-clone①⑤
        link-type="dfn"} of `last partially contained child`{.variable}.

    2.  [Append](#concept-node-append){#ref-for-concept-node-append②⑥
        link-type="dfn"} `clone`{.variable} to `fragment`{.variable}.

    3.  Let `subrange`{.variable} be a new [live
        range](#concept-live-range){#ref-for-concept-live-range④⓪
        link-type="dfn"} whose
        [start](#concept-range-start){#ref-for-concept-range-start②⑥
        link-type="dfn"} is
        (`last partially contained child`{.variable}, 0) and whose
        [end](#concept-range-end){#ref-for-concept-range-end②⑥
        link-type="dfn"} is (`original end node`{.variable},
        `original end offset`{.variable}).

    4.  Let `subfragment`{.variable} be the result of [cloning the
        contents](#concept-range-clone){#ref-for-concept-range-clone①
        link-type="dfn"} of `subrange`{.variable}.

    5.  [Append](#concept-node-append){#ref-for-concept-node-append②⑦
        link-type="dfn"} `subfragment`{.variable} to `clone`{.variable}.

18. Return `fragment`{.variable}.

The [`cloneContents()`]{#dom-range-clonecontents .dfn .dfn-paneled
.idl-code dfn-for="Range" dfn-type="method" export=""} method steps are
to return the result of [cloning the
contents](#concept-range-clone){#ref-for-concept-range-clone②
link-type="dfn"} of
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥⑥
link-type="dfn"}.

To [insert]{#concept-range-insert .dfn .dfn-paneled dfn-for="live range"
dfn-type="dfn" export=""} a
[node](#concept-node){#ref-for-concept-node②⓪④ link-type="dfn"}
`node`{.variable} into a [live
range](#concept-live-range){#ref-for-concept-live-range④①
link-type="dfn"} `range`{.variable}, run these steps:

1.  If `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③④
    link-type="dfn"} is a
    [`ProcessingInstruction`{.idl}](#processinginstruction){#ref-for-processinginstruction①⑦
    link-type="idl"} or [`Comment`{.idl}](#comment){#ref-for-comment①⑦
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪⑤
    link-type="dfn"}, is a [`Text`{.idl}](#text){#ref-for-text④⑨
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪⑥
    link-type="dfn"} whose
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥⓪
    link-type="dfn"} is null, or is `node`{.variable}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧④
    link-type="dfn"} a
    \"[`HierarchyRequestError`{.idl}](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#ref-for-hierarchyrequesterror②⑥
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨⑧
    link-type="idl"}.

2.  Let `referenceNode`{.variable} be null.

3.  If `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③⑤
    link-type="dfn"} is a [`Text`{.idl}](#text){#ref-for-text⑤⓪
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪⑦
    link-type="dfn"}, set `referenceNode`{.variable} to that
    [`Text`{.idl}](#text){#ref-for-text⑤① link-type="idl"}
    [node](#concept-node){#ref-for-concept-node②⓪⑧ link-type="dfn"}.

4.  Otherwise, set `referenceNode`{.variable} to the
    [child](#concept-tree-child){#ref-for-concept-tree-child⑧⓪
    link-type="dfn"} of [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③⑥
    link-type="dfn"} whose
    [index](#concept-tree-index){#ref-for-concept-tree-index②⓪
    link-type="dfn"} is [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②④
    link-type="dfn"}, and null if there is no such
    [child](#concept-tree-child){#ref-for-concept-tree-child⑧①
    link-type="dfn"}.

5.  Let `parent`{.variable} be `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③⑦
    link-type="dfn"} if `referenceNode`{.variable} is null, and
    `referenceNode`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥①
    link-type="dfn"} otherwise.

6.  [Ensure pre-insert
    validity](#concept-node-ensure-pre-insertion-validity){#ref-for-concept-node-ensure-pre-insertion-validity②
    link-type="dfn"} of `node`{.variable} into `parent`{.variable}
    before `referenceNode`{.variable}.

7.  If `range`{.variable}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③⑧
    link-type="dfn"} is a [`Text`{.idl}](#text){#ref-for-text⑤②
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②⓪⑨
    link-type="dfn"}, set `referenceNode`{.variable} to the result of
    [splitting](#concept-text-split){#ref-for-concept-text-split③
    link-type="dfn"} it with offset `range`{.variable}'s [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②⑤
    link-type="dfn"}.

8.  If `node`{.variable} is `referenceNode`{.variable}, set
    `referenceNode`{.variable} to its [next
    sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①⑥
    link-type="dfn"}.

9.  If `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥②
    link-type="dfn"} is non-null, then
    [remove](#concept-node-remove){#ref-for-concept-node-remove①⑤
    link-type="dfn"} `node`{.variable}.

10. Let `newOffset`{.variable} be `parent`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length②③
    link-type="dfn"} if `referenceNode`{.variable} is null; otherwise
    `referenceNode`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index②①
    link-type="dfn"}.

11. Increase `newOffset`{.variable} by `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length②④
    link-type="dfn"} if `node`{.variable} is a
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②①⓪
    link-type="dfn"}; otherwise 1.

12. [Pre-insert](#concept-node-pre-insert){#ref-for-concept-node-pre-insert①②
    link-type="dfn"} `node`{.variable} into `parent`{.variable} before
    `referenceNode`{.variable}.

13. If `range`{.variable} is
    [collapsed](#range-collapsed){#ref-for-range-collapsed⑤
    link-type="dfn"}, then set `range`{.variable}'s
    [end](#concept-range-end){#ref-for-concept-range-end②⑦
    link-type="dfn"} to (`parent`{.variable}, `newOffset`{.variable}).

The [`insertNode(``node`{.variable}`)`]{#dom-range-insertnode .dfn
.dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are to
[insert](#concept-range-insert){#ref-for-concept-range-insert
link-type="dfn"} `node`{.variable} into
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥⑦
link-type="dfn"}.

The
[`surroundContents(``newParent`{.variable}`)`]{#dom-range-surroundcontents
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  If a non-[`Text`{.idl}](#text){#ref-for-text⑤③ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node②①① link-type="dfn"} is
    [partially
    contained](#partially-contained){#ref-for-partially-contained①①
    link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥⑧
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧⑤
    link-type="dfn"} an
    \"[`InvalidStateError`{.idl}](https://webidl.spec.whatwg.org/#invalidstateerror){#ref-for-invalidstateerror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException⑨⑨
    link-type="idl"}.

2.  If `newParent`{.variable} is a
    [`Document`{.idl}](#document){#ref-for-document②① link-type="idl"},
    [`DocumentType`{.idl}](#documenttype){#ref-for-documenttype②④
    link-type="idl"}, or
    [`DocumentFragment`{.idl}](#documentfragment){#ref-for-documentfragment③⑨
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②①②
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧⑥
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror⑧
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪⓪
    link-type="idl"}.

    For historical reasons
    [`CharacterData`{.idl}](#characterdata){#ref-for-characterdata③②
    link-type="idl"} [nodes](#concept-node){#ref-for-concept-node②①③
    link-type="dfn"} are not checked here and end up throwing later on
    as a side effect.

3.  Let `fragment`{.variable} be the result of
    [extracting](#concept-range-extract){#ref-for-concept-range-extract③
    link-type="dfn"}
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑥⑨
    link-type="dfn"}.

4.  If `newParent`{.variable} has
    [children](#concept-tree-child){#ref-for-concept-tree-child⑧②
    link-type="dfn"}, then [replace
    all](#concept-node-replace-all){#ref-for-concept-node-replace-all②
    link-type="dfn"} with null within `newParent`{.variable}.

5.  [Insert](#concept-range-insert){#ref-for-concept-range-insert①
    link-type="dfn"} `newParent`{.variable} into
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦⓪
    link-type="dfn"}.

6.  [Append](#concept-node-append){#ref-for-concept-node-append②⑧
    link-type="dfn"} `fragment`{.variable} to `newParent`{.variable}.

7.  [Select](#concept-range-select){#ref-for-concept-range-select①
    link-type="dfn"} `newParent`{.variable} within
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦①
    link-type="dfn"}.

The [`cloneRange()`]{#dom-range-clonerange .dfn .dfn-paneled .idl-code
dfn-for="Range" dfn-type="method" export=""} method steps are to return
a new [live range](#concept-live-range){#ref-for-concept-live-range④②
link-type="dfn"} with the same
[start](#concept-range-start){#ref-for-concept-range-start②⑦
link-type="dfn"} and
[end](#concept-range-end){#ref-for-concept-range-end②⑧ link-type="dfn"}
as [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦②
link-type="dfn"}.

The [`detach()`]{#dom-range-detach .dfn .dfn-paneled .idl-code
dfn-for="Range" dfn-type="method" export=""} method steps are to do
nothing. [Its functionality (disabling a
[`Range`{.idl}](#range){#ref-for-range⑧ link-type="idl"} object) was
removed, but the method itself is preserved for compatibility.]{.note}

------------------------------------------------------------------------

`position`{.variable} = `range`{.variable} . [`comparePoint(node, offset)`{.idl}](#dom-range-comparepoint){#ref-for-dom-range-comparepoint① link-type="idl"}
:   Returns −1 if the point is before the range, 0 if the point is in
    the range, and 1 if the point is after the range.

`intersects`{.variable} = `range`{.variable} . [`intersectsNode(node)`{.idl}](#dom-range-intersectsnode){#ref-for-dom-range-intersectsnode① link-type="idl"}
:   Returns whether `range`{.variable} intersects `node`{.variable}.

::: impl
The
[`isPointInRange(``node`{.variable}`, ``offset`{.variable}`)`]{#dom-range-ispointinrange
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  If `node`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root④⑧
    link-type="dfn"} is different from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦③
    link-type="dfn"}'s
    [root](#concept-range-root){#ref-for-concept-range-root⑦
    link-type="dfn"}, return false.

2.  If `node`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype②⑨
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧⑦
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror⑨
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪①
    link-type="idl"}.

3.  If `offset`{.variable} is greater than `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length②⑤
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧⑧
    link-type="dfn"} an
    \"[`IndexSizeError`{.idl}](https://webidl.spec.whatwg.org/#indexsizeerror){#ref-for-indexsizeerror④
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪②
    link-type="idl"}.

4.  If (`node`{.variable}, `offset`{.variable}) is
    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before⑧
    link-type="dfn"}
    [start](#concept-range-start){#ref-for-concept-range-start②⑧
    link-type="dfn"} or
    [after](#concept-range-bp-after){#ref-for-concept-range-bp-after⑦
    link-type="dfn"}
    [end](#concept-range-end){#ref-for-concept-range-end②⑨
    link-type="dfn"}, return false.

5.  Return true.

The
[`comparePoint(``node`{.variable}`, ``offset`{.variable}`)`]{#dom-range-comparepoint
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  If `node`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root④⑨
    link-type="dfn"} is different from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦④
    link-type="dfn"}'s
    [root](#concept-range-root){#ref-for-concept-range-root⑧
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑧⑨
    link-type="dfn"} a
    \"[`WrongDocumentError`{.idl}](https://webidl.spec.whatwg.org/#wrongdocumenterror){#ref-for-wrongdocumenterror①
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪③
    link-type="idl"}.

2.  If `node`{.variable} is a
    [doctype](#concept-doctype){#ref-for-concept-doctype③⓪
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨⓪
    link-type="dfn"} an
    \"[`InvalidNodeTypeError`{.idl}](https://webidl.spec.whatwg.org/#invalidnodetypeerror){#ref-for-invalidnodetypeerror①⓪
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪④
    link-type="idl"}.

3.  If `offset`{.variable} is greater than `node`{.variable}'s
    [length](#concept-node-length){#ref-for-concept-node-length②⑥
    link-type="dfn"}, then
    [throw](https://webidl.spec.whatwg.org/#dfn-throw){#ref-for-dfn-throw⑨①
    link-type="dfn"} an
    \"[`IndexSizeError`{.idl}](https://webidl.spec.whatwg.org/#indexsizeerror){#ref-for-indexsizeerror⑤
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪⑤
    link-type="idl"}.

4.  If (`node`{.variable}, `offset`{.variable}) is
    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before⑨
    link-type="dfn"}
    [start](#concept-range-start){#ref-for-concept-range-start②⑨
    link-type="dfn"}, return −1.

5.  If (`node`{.variable}, `offset`{.variable}) is
    [after](#concept-range-bp-after){#ref-for-concept-range-bp-after⑧
    link-type="dfn"}
    [end](#concept-range-end){#ref-for-concept-range-end③⓪
    link-type="dfn"}, return 1.

6.  Return 0.

------------------------------------------------------------------------

The [`intersectsNode(``node`{.variable}`)`]{#dom-range-intersectsnode
.dfn .dfn-paneled .idl-code dfn-for="Range" dfn-type="method" export=""}
method steps are:

1.  If `node`{.variable}'s
    [root](#concept-tree-root){#ref-for-concept-tree-root⑤⓪
    link-type="dfn"} is different from
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦⑤
    link-type="dfn"}'s
    [root](#concept-range-root){#ref-for-concept-range-root⑨
    link-type="dfn"}, return false.
2.  Let `parent`{.variable} be `node`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥③
    link-type="dfn"}.
3.  If `parent`{.variable} is null, return true.
4.  Let `offset`{.variable} be `node`{.variable}'s
    [index](#concept-tree-index){#ref-for-concept-tree-index②②
    link-type="dfn"}.
5.  If (`parent`{.variable}, `offset`{.variable}) is
    [before](#concept-range-bp-before){#ref-for-concept-range-bp-before①⓪
    link-type="dfn"}
    [end](#concept-range-end){#ref-for-concept-range-end③①
    link-type="dfn"} and (`parent`{.variable}, `offset`{.variable}
    plus 1) is
    [after](#concept-range-bp-after){#ref-for-concept-range-bp-after⑨
    link-type="dfn"}
    [start](#concept-range-start){#ref-for-concept-range-start③⓪
    link-type="dfn"}, return true.
6.  Return false.
:::

------------------------------------------------------------------------

The [stringification behavior]{#dom-range-stringifier .dfn .dfn-paneled
dfn-for="Range" dfn-type="dfn" export="" lt="stringificationbehavior"}
must run these steps:

1.  Let `s`{.variable} be the empty string.

2.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦⑥
    link-type="dfn"}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node③⑨
    link-type="dfn"} is
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦⑦
    link-type="dfn"}'s [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node③⓪
    link-type="dfn"} and it is a [`Text`{.idl}](#text){#ref-for-text⑤④
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②①④
    link-type="dfn"}, then return the substring of that
    [`Text`{.idl}](#text){#ref-for-text⑤⑤ link-type="idl"}
    [node](#concept-node){#ref-for-concept-node②①⑤ link-type="dfn"}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data④⑥ link-type="dfn"}
    beginning at
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦⑧
    link-type="dfn"}'s [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②⑥
    link-type="dfn"} and ending at
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑦⑨
    link-type="dfn"}'s [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset②④
    link-type="dfn"}.

3.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧⓪
    link-type="dfn"}'s [start
    node](#concept-range-start-node){#ref-for-concept-range-start-node④⓪
    link-type="dfn"} is a [`Text`{.idl}](#text){#ref-for-text⑤⑥
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②①⑥
    link-type="dfn"}, then append the substring of that
    [node](#concept-node){#ref-for-concept-node②①⑦ link-type="dfn"}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data④⑦ link-type="dfn"}
    from [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧①
    link-type="dfn"}'s [start
    offset](#concept-range-start-offset){#ref-for-concept-range-start-offset②⑦
    link-type="dfn"} until the end to `s`{.variable}.

4.  Append the
    [concatenation](https://infra.spec.whatwg.org/#string-concatenate){#ref-for-string-concatenate⑤
    link-type="dfn"} of the
    [data](#concept-cd-data){#ref-for-concept-cd-data④⑧ link-type="dfn"}
    of all [`Text`{.idl}](#text){#ref-for-text⑤⑦ link-type="idl"}
    [nodes](#concept-node){#ref-for-concept-node②①⑧ link-type="dfn"}
    that are [contained](#contained){#ref-for-contained①⑤
    link-type="dfn"} in
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧②
    link-type="dfn"}, in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order③②
    link-type="dfn"}, to `s`{.variable}.

5.  If [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧③
    link-type="dfn"}'s [end
    node](#concept-range-end-node){#ref-for-concept-range-end-node③①
    link-type="dfn"} is a [`Text`{.idl}](#text){#ref-for-text⑤⑧
    link-type="idl"} [node](#concept-node){#ref-for-concept-node②①⑨
    link-type="dfn"}, then append the substring of that
    [node](#concept-node){#ref-for-concept-node②②⓪ link-type="dfn"}'s
    [data](#concept-cd-data){#ref-for-concept-cd-data④⑨ link-type="dfn"}
    from its start until
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧④
    link-type="dfn"}'s [end
    offset](#concept-range-end-offset){#ref-for-concept-range-end-offset②⑤
    link-type="dfn"} to `s`{.variable}.

6.  Return `s`{.variable}.

------------------------------------------------------------------------

The
[`createContextualFragment()`{.idl}](https://w3c.github.io/DOM-Parsing/#dom-range-createcontextualfragment){#ref-for-dom-range-createcontextualfragment
link-type="idl"},
[`getClientRects()`{.idl}](https://drafts.csswg.org/cssom-view-1/#dom-range-getclientrects){#ref-for-dom-range-getclientrects
link-type="idl"}, and
[`getBoundingClientRect()`{.idl}](https://drafts.csswg.org/cssom-view-1/#dom-range-getboundingclientrect){#ref-for-dom-range-getboundingclientrect
link-type="idl"} methods are defined in other specifications.
[\[DOM-Parsing\]](#biblio-dom-parsing "DOM Parsing and Serialization"){link-type="biblio"}
[\[CSSOM-VIEW\]](#biblio-cssom-view "CSSOM View Module"){link-type="biblio"}

## [6. ]{.secno}[Traversal]{.content}[](#traversal){.self-link} {#traversal .heading .settled level="6"}

[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator④
link-type="idl"} and
[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker② link-type="idl"}
objects can be used to filter and traverse
[node](#concept-node){#ref-for-concept-node②②① link-type="dfn"}
[trees](#concept-tree){#ref-for-concept-tree②④ link-type="dfn"}.

Each [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator⑤
link-type="idl"} and
[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker③ link-type="idl"}
object has an associated [active flag]{#concept-traversal-active .dfn
.dfn-paneled dfn-for="traversal" dfn-type="dfn" noexport=""} to avoid
recursive invocations. It is initially unset.

Each [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator⑥
link-type="idl"} and
[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker④ link-type="idl"}
object also has an associated [root]{#concept-traversal-root .dfn
.dfn-paneled dfn-for="traversal" dfn-type="dfn" noexport=""} (a
[node](#concept-node){#ref-for-concept-node②②② link-type="dfn"}), a
[whatToShow]{#concept-traversal-whattoshow .dfn .dfn-paneled
dfn-for="traversal" dfn-type="dfn" noexport=""} (a bitmask), and a
[filter]{#concept-traversal-filter .dfn .dfn-paneled dfn-for="traversal"
dfn-type="dfn" noexport=""} (a callback).

To [filter]{#concept-node-filter .dfn .dfn-paneled dfn-type="dfn"
noexport=""} a [node](#concept-node){#ref-for-concept-node②②③
link-type="dfn"} `node`{.variable} within a
[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator⑦
link-type="idl"} or
[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker⑤ link-type="idl"}
object `traverser`{.variable}, run these steps:

1.  If `traverser`{.variable}'s [active
    flag](#concept-traversal-active){#ref-for-concept-traversal-active
    link-type="dfn"} is set, then throw an
    \"[`InvalidStateError`{.idl}](https://webidl.spec.whatwg.org/#invalidstateerror){#ref-for-invalidstateerror②
    .idl-code link-type="exception"}\"
    [`DOMException`{.idl}](https://webidl.spec.whatwg.org/#idl-DOMException){#ref-for-idl-DOMException①⓪⑥
    link-type="idl"}.

2.  Let `n`{.variable} be `node`{.variable}'s
    [`nodeType`{.idl}](#dom-node-nodetype){#ref-for-dom-node-nodetype②
    link-type="idl"} attribute value − 1.

3.  If the `n`{.variable}^th^ bit (where 0 is the least significant bit)
    of `traverser`{.variable}'s
    [whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow②
    link-type="dfn"} is not set, then return
    [`FILTER_SKIP`{.idl}](#dom-nodefilter-filter_skip){#ref-for-dom-nodefilter-filter_skip
    link-type="idl"}.

4.  If `traverser`{.variable}'s
    [filter](#concept-traversal-filter){#ref-for-concept-traversal-filter②
    link-type="dfn"} is null, then return
    [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept
    link-type="idl"}.

5.  Set `traverser`{.variable}'s [active
    flag](#concept-traversal-active){#ref-for-concept-traversal-active①
    link-type="dfn"}.

6.  Let `result`{.variable} be the return value of [call a user object's
    operation](https://webidl.spec.whatwg.org/#call-a-user-objects-operation){#ref-for-call-a-user-objects-operation①
    link-type="dfn"} with `traverser`{.variable}'s
    [filter](#concept-traversal-filter){#ref-for-concept-traversal-filter③
    link-type="dfn"}, \"`acceptNode`\", and « `node`{.variable} ». If
    this throws an exception, then unset `traverser`{.variable}'s
    [active
    flag](#concept-traversal-active){#ref-for-concept-traversal-active②
    link-type="dfn"} and rethrow the exception.

7.  Unset `traverser`{.variable}'s [active
    flag](#concept-traversal-active){#ref-for-concept-traversal-active③
    link-type="dfn"}.

8.  Return `result`{.variable}.

### [6.1. ]{.secno}[Interface [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator⑧ link-type="idl"}]{.content}[](#interface-nodeiterator){.self-link} {#interface-nodeiterator .heading .settled level="6.1"}

``` {.idl .highlight .def}
[Exposed=Window]
interface NodeIterator {
  [SameObject] readonly attribute Node root;
  readonly attribute Node referenceNode;
  readonly attribute boolean pointerBeforeReferenceNode;
  readonly attribute unsigned long whatToShow;
  readonly attribute NodeFilter? filter;

  Node? nextNode();
  Node? previousNode();

  undefined detach();
};
```

[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator⑨
link-type="idl"} objects can be created using the
[`createNodeIterator()`{.idl}](#dom-document-createnodeiterator){#ref-for-dom-document-createnodeiterator①
link-type="idl"} method on
[`Document`{.idl}](#document){#ref-for-document②② link-type="idl"}
objects.

Each [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①⓪
link-type="idl"} object has an associated [iterator
collection]{#iterator-collection .dfn .dfn-paneled
dfn-for="NodeIterator" dfn-type="dfn" noexport=""}, which is a
[collection](#concept-collection){#ref-for-concept-collection①⑨
link-type="dfn"} rooted at the
[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①①
link-type="idl"} object's
[root](#concept-traversal-root){#ref-for-concept-traversal-root④
link-type="dfn"}, whose filter matches any
[node](#concept-node){#ref-for-concept-node②②④ link-type="dfn"}.

Each [`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①②
link-type="idl"} object also has an associated
[reference]{#nodeiterator-reference .dfn .dfn-paneled
dfn-for="NodeIterator" dfn-type="dfn" noexport=""} (a
[node](#concept-node){#ref-for-concept-node②②⑤ link-type="dfn"}) and
[pointer before reference]{#nodeiterator-pointer-before-reference .dfn
.dfn-paneled dfn-for="NodeIterator" dfn-type="dfn" noexport=""} (a
boolean).

As mentioned earlier,
[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①③
link-type="idl"} objects have an associated [active
flag](#concept-traversal-active){#ref-for-concept-traversal-active④
link-type="dfn"},
[root](#concept-traversal-root){#ref-for-concept-traversal-root⑤
link-type="dfn"},
[whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow③
link-type="dfn"}, and
[filter](#concept-traversal-filter){#ref-for-concept-traversal-filter④
link-type="dfn"} as well.

The [`NodeIterator` pre-remove steps]{#nodeiterator-pre-removing-steps
.dfn .dfn-paneled dfn-type="dfn" noexport=""} given a
`nodeIterator`{.variable} and `toBeRemovedNode`{.variable}, are as
follows:

1.  If `toBeRemovedNode`{.variable} is not an [inclusive
    ancestor](#concept-tree-inclusive-ancestor){#ref-for-concept-tree-inclusive-ancestor②⑧
    link-type="dfn"} of `nodeIterator`{.variable}'s
    [reference](#nodeiterator-reference){#ref-for-nodeiterator-reference①
    link-type="dfn"}, or `toBeRemovedNode`{.variable} is
    `nodeIterator`{.variable}'s
    [root](#concept-traversal-root){#ref-for-concept-traversal-root⑥
    link-type="dfn"}, then return.

2.  If `nodeIterator`{.variable}'s [pointer before
    reference](#nodeiterator-pointer-before-reference){#ref-for-nodeiterator-pointer-before-reference①
    link-type="dfn"} is true:

    1.  Let `next`{.variable} be `toBeRemovedNode`{.variable}'s first
        [following](#concept-tree-following){#ref-for-concept-tree-following①③
        link-type="dfn"} [node](#concept-node){#ref-for-concept-node②②⑥
        link-type="dfn"} that is an [inclusive
        descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant⑦
        link-type="dfn"} of `nodeIterator`{.variable}'s
        [root](#concept-traversal-root){#ref-for-concept-traversal-root⑦
        link-type="dfn"} and is not an [inclusive
        descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant⑧
        link-type="dfn"} of `toBeRemovedNode`{.variable}, and null if
        there is no such [node](#concept-node){#ref-for-concept-node②②⑦
        link-type="dfn"}.

    2.  If `next`{.variable} is non-null, then set
        `nodeIterator`{.variable}'s
        [reference](#nodeiterator-reference){#ref-for-nodeiterator-reference②
        link-type="dfn"} to `next`{.variable} and return.

    3.  Otherwise, set `nodeIterator`{.variable}'s [pointer before
        reference](#nodeiterator-pointer-before-reference){#ref-for-nodeiterator-pointer-before-reference②
        link-type="dfn"} to false.

        Steps are not terminated here.

3.  Set `nodeIterator`{.variable}'s
    [reference](#nodeiterator-reference){#ref-for-nodeiterator-reference③
    link-type="dfn"} to `toBeRemovedNode`{.variable}'s
    [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥④
    link-type="dfn"}, if `toBeRemovedNode`{.variable}'s [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①⓪
    link-type="dfn"} is null, and to the [inclusive
    descendant](#concept-tree-inclusive-descendant){#ref-for-concept-tree-inclusive-descendant⑨
    link-type="dfn"} of `toBeRemovedNode`{.variable}'s [previous
    sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①①
    link-type="dfn"} that appears last in [tree
    order](#concept-tree-order){#ref-for-concept-tree-order③③
    link-type="dfn"} otherwise.

------------------------------------------------------------------------

The [`root`]{#dom-nodeiterator-root .dfn .dfn-paneled .idl-code
dfn-for="NodeIterator" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧⑤
link-type="dfn"}'s
[root](#concept-traversal-root){#ref-for-concept-traversal-root⑧
link-type="dfn"}.

The [`referenceNode`]{#dom-nodeiterator-referencenode .dfn .dfn-paneled
.idl-code dfn-for="NodeIterator" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧⑥
link-type="dfn"}'s
[reference](#nodeiterator-reference){#ref-for-nodeiterator-reference④
link-type="dfn"}.

The
[`pointerBeforeReferenceNode`]{#dom-nodeiterator-pointerbeforereferencenode
.dfn .dfn-paneled .idl-code dfn-for="NodeIterator" dfn-type="attribute"
export=""} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧⑦
link-type="dfn"}'s [pointer before
reference](#nodeiterator-pointer-before-reference){#ref-for-nodeiterator-pointer-before-reference③
link-type="dfn"}.

The [`whatToShow`]{#dom-nodeiterator-whattoshow .dfn .dfn-paneled
.idl-code dfn-for="NodeIterator" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧⑧
link-type="dfn"}'s
[whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow④
link-type="dfn"}.

The [`filter`]{#dom-nodeiterator-filter .dfn .dfn-paneled .idl-code
dfn-for="NodeIterator" dfn-type="attribute" export=""} getter steps are
to return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑧⑨
link-type="dfn"}'s
[filter](#concept-traversal-filter){#ref-for-concept-traversal-filter⑤
link-type="dfn"}.

To [traverse]{#concept-nodeiterator-traverse .dfn .dfn-paneled
dfn-for="NodeIterator" dfn-type="dfn" export=""}, given a
[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①④
link-type="idl"} object `iterator`{.variable} and a direction
`direction`{.variable}, run these steps:

1.  Let `node`{.variable} be `iterator`{.variable}'s
    [reference](#nodeiterator-reference){#ref-for-nodeiterator-reference⑤
    link-type="dfn"}.

2.  Let `beforeNode`{.variable} be `iterator`{.variable}'s [pointer
    before
    reference](#nodeiterator-pointer-before-reference){#ref-for-nodeiterator-pointer-before-reference④
    link-type="dfn"}.

3.  While true:

    1.  Branch on `direction`{.variable}:

        next

        :   If `beforeNode`{.variable} is false, then set
            `node`{.variable} to the first
            [node](#concept-node){#ref-for-concept-node②②⑧
            link-type="dfn"}
            [following](#concept-tree-following){#ref-for-concept-tree-following①④
            link-type="dfn"} `node`{.variable} in
            `iterator`{.variable}'s [iterator
            collection](#iterator-collection){#ref-for-iterator-collection
            link-type="dfn"}. If there is no such
            [node](#concept-node){#ref-for-concept-node②②⑨
            link-type="dfn"}, then return null.

            If `beforeNode`{.variable} is true, then set it to false.

        previous

        :   If `beforeNode`{.variable} is true, then set
            `node`{.variable} to the first
            [node](#concept-node){#ref-for-concept-node②③⓪
            link-type="dfn"}
            [preceding](#concept-tree-preceding){#ref-for-concept-tree-preceding①②
            link-type="dfn"} `node`{.variable} in
            `iterator`{.variable}'s [iterator
            collection](#iterator-collection){#ref-for-iterator-collection①
            link-type="dfn"}. If there is no such
            [node](#concept-node){#ref-for-concept-node②③①
            link-type="dfn"}, then return null.

            If `beforeNode`{.variable} is false, then set it to true.

    2.  Let `result`{.variable} be the result of
        [filtering](#concept-node-filter){#ref-for-concept-node-filter
        link-type="dfn"} `node`{.variable} within `iterator`{.variable}.

    3.  If `result`{.variable} is
        [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept①
        link-type="idl"}, then
        [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break②
        link-type="dfn"}.

4.  Set `iterator`{.variable}'s
    [reference](#nodeiterator-reference){#ref-for-nodeiterator-reference⑥
    link-type="dfn"} to `node`{.variable}.

5.  Set `iterator`{.variable}'s [pointer before
    reference](#nodeiterator-pointer-before-reference){#ref-for-nodeiterator-pointer-before-reference⑤
    link-type="dfn"} to `beforeNode`{.variable}.

6.  Return `node`{.variable}.

The [`nextNode()`]{#dom-nodeiterator-nextnode .dfn .dfn-paneled
.idl-code dfn-for="NodeIterator" dfn-type="method" export=""} method
steps are to return the result of
[traversing](#concept-nodeiterator-traverse){#ref-for-concept-nodeiterator-traverse
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨⓪
link-type="dfn"} and next.

The [`previousNode()`]{#dom-nodeiterator-previousnode .dfn .dfn-paneled
.idl-code dfn-for="NodeIterator" dfn-type="method" export=""} method
steps are to return the result of
[traversing](#concept-nodeiterator-traverse){#ref-for-concept-nodeiterator-traverse①
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨①
link-type="dfn"} and previous.

The [`detach()`]{#dom-nodeiterator-detach .dfn .dfn-paneled .idl-code
dfn-for="NodeIterator" dfn-type="method" export=""} method steps are to
do nothing. [Its functionality (disabling a
[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①⑤
link-type="idl"} object) was removed, but the method itself is preserved
for compatibility.]{.note}

### [6.2. ]{.secno}[Interface [`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker⑥ link-type="idl"}]{.content}[](#interface-treewalker){.self-link} {#interface-treewalker .heading .settled level="6.2"}

``` {.idl .highlight .def}
[Exposed=Window]
interface TreeWalker {
  [SameObject] readonly attribute Node root;
  readonly attribute unsigned long whatToShow;
  readonly attribute NodeFilter? filter;
           attribute Node currentNode;

  Node? parentNode();
  Node? firstChild();
  Node? lastChild();
  Node? previousSibling();
  Node? nextSibling();
  Node? previousNode();
  Node? nextNode();
};
```

[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker⑦ link-type="idl"}
objects can be created using the
[`createTreeWalker()`{.idl}](#dom-document-createtreewalker){#ref-for-dom-document-createtreewalker①
link-type="idl"} method on
[`Document`{.idl}](#document){#ref-for-document②③ link-type="idl"}
objects.

Each [`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker⑧
link-type="idl"} object has an associated [current]{#treewalker-current
.dfn .dfn-paneled dfn-for="TreeWalker" dfn-type="dfn" noexport=""} (a
[node](#concept-node){#ref-for-concept-node②③② link-type="dfn"}).

As mentioned earlier
[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker⑨ link-type="idl"}
objects have an associated
[root](#concept-traversal-root){#ref-for-concept-traversal-root⑨
link-type="dfn"},
[whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow⑤
link-type="dfn"}, and
[filter](#concept-traversal-filter){#ref-for-concept-traversal-filter⑥
link-type="dfn"} as well.

The [`root`]{#dom-treewalker-root .dfn .dfn-paneled .idl-code
dfn-for="TreeWalker" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨②
link-type="dfn"}'s
[root](#concept-traversal-root){#ref-for-concept-traversal-root①⓪
link-type="dfn"}.

The [`whatToShow`]{#dom-treewalker-whattoshow .dfn .dfn-paneled
.idl-code dfn-for="TreeWalker" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨③
link-type="dfn"}'s
[whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow⑥
link-type="dfn"}.

The [`filter`]{#dom-treewalker-filter .dfn .dfn-paneled .idl-code
dfn-for="TreeWalker" dfn-type="attribute" export=""} getter steps are to
return [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨④
link-type="dfn"}'s
[filter](#concept-traversal-filter){#ref-for-concept-traversal-filter⑦
link-type="dfn"}.

The [`currentNode`]{#dom-treewalker-currentnode .dfn .dfn-paneled
.idl-code dfn-for="TreeWalker" dfn-type="attribute" export=""} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨⑤
link-type="dfn"}'s
[current](#treewalker-current){#ref-for-treewalker-current①
link-type="dfn"}.

The
[`currentNode`{.idl}](#dom-treewalker-currentnode){#ref-for-dom-treewalker-currentnode①
link-type="idl"} setter steps are to set
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨⑥
link-type="dfn"}'s
[current](#treewalker-current){#ref-for-treewalker-current②
link-type="dfn"} to the given value.

------------------------------------------------------------------------

The [`parentNode()`]{#dom-treewalker-parentnode .dfn .dfn-paneled
.idl-code dfn-for="TreeWalker" dfn-type="method" export=""} method steps
are:

1.  Let `node`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨⑦
    link-type="dfn"}'s
    [current](#treewalker-current){#ref-for-treewalker-current③
    link-type="dfn"}.

2.  While `node`{.variable} is non-null and is not
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨⑧
    link-type="dfn"}'s
    [root](#concept-traversal-root){#ref-for-concept-traversal-root①①
    link-type="dfn"}:

    1.  Set `node`{.variable} to `node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥⑤
        link-type="dfn"}.

    2.  If `node`{.variable} is non-null and
        [filtering](#concept-node-filter){#ref-for-concept-node-filter①
        link-type="dfn"} `node`{.variable} within
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this②⑨⑨
        link-type="dfn"} returns
        [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept②
        link-type="idl"}, then set
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪⓪
        link-type="dfn"}'s
        [current](#treewalker-current){#ref-for-treewalker-current④
        link-type="dfn"} to `node`{.variable} and return
        `node`{.variable}.

3.  Return null.

To [traverse children]{#concept-traverse-children .dfn .dfn-paneled
dfn-for="TreeWalker" dfn-type="dfn" noexport=""}, given a
`walker`{.variable} and `type`{.variable}, run these steps:

1.  Let `node`{.variable} be `walker`{.variable}'s
    [current](#treewalker-current){#ref-for-treewalker-current⑤
    link-type="dfn"}.

2.  Set `node`{.variable} to `node`{.variable}'s [first
    child](#concept-tree-first-child){#ref-for-concept-tree-first-child⑥
    link-type="dfn"} if `type`{.variable} is first, and
    `node`{.variable}'s [last
    child](#concept-tree-last-child){#ref-for-concept-tree-last-child⑥
    link-type="dfn"} if `type`{.variable} is last.

3.  While `node`{.variable} is non-null:

    1.  Let `result`{.variable} be the result of
        [filtering](#concept-node-filter){#ref-for-concept-node-filter②
        link-type="dfn"} `node`{.variable} within `walker`{.variable}.

    2.  If `result`{.variable} is
        [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept③
        link-type="idl"}, then set `walker`{.variable}'s
        [current](#treewalker-current){#ref-for-treewalker-current⑥
        link-type="dfn"} to `node`{.variable} and return
        `node`{.variable}.

    3.  If `result`{.variable} is
        [`FILTER_SKIP`{.idl}](#dom-nodefilter-filter_skip){#ref-for-dom-nodefilter-filter_skip①
        link-type="idl"}:

        1.  Let `child`{.variable} be `node`{.variable}'s [first
            child](#concept-tree-first-child){#ref-for-concept-tree-first-child⑦
            link-type="dfn"} if `type`{.variable} is first, and
            `node`{.variable}'s [last
            child](#concept-tree-last-child){#ref-for-concept-tree-last-child⑦
            link-type="dfn"} if `type`{.variable} is last.

        2.  If `child`{.variable} is non-null, then set
            `node`{.variable} to `child`{.variable} and
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#ref-for-iteration-continue⑤
            link-type="dfn"}.

    4.  While `node`{.variable} is non-null:

        1.  Let `sibling`{.variable} be `node`{.variable}'s [next
            sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①⑦
            link-type="dfn"} if `type`{.variable} is first, and
            `node`{.variable}'s [previous
            sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①②
            link-type="dfn"} if `type`{.variable} is last.

        2.  If `sibling`{.variable} is non-null, then set
            `node`{.variable} to `sibling`{.variable} and
            [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break③
            link-type="dfn"}.

        3.  Let `parent`{.variable} be `node`{.variable}'s
            [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥⑥
            link-type="dfn"}.

        4.  If `parent`{.variable} is null, `walker`{.variable}'s
            [root](#concept-traversal-root){#ref-for-concept-traversal-root①②
            link-type="dfn"}, or `walker`{.variable}'s
            [current](#treewalker-current){#ref-for-treewalker-current⑦
            link-type="dfn"}, then return null.

        5.  Set `node`{.variable} to `parent`{.variable}.

4.  Return null.

The [`firstChild()`]{#dom-treewalker-firstchild .dfn .dfn-paneled
.idl-code dfn-for="TreeWalker" dfn-type="method" export=""} method steps
are to [traverse
children](#concept-traverse-children){#ref-for-concept-traverse-children
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪①
link-type="dfn"} and first.

The [`lastChild()`]{#dom-treewalker-lastchild .dfn .dfn-paneled
.idl-code dfn-for="TreeWalker" dfn-type="method" export=""} method steps
are to [traverse
children](#concept-traverse-children){#ref-for-concept-traverse-children①
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪②
link-type="dfn"} and last.

To [traverse siblings]{#concept-traverse-siblings .dfn .dfn-paneled
dfn-type="dfn" noexport=""}, given a `walker`{.variable} and
`type`{.variable}, run these steps:

1.  Let `node`{.variable} be `walker`{.variable}'s
    [current](#treewalker-current){#ref-for-treewalker-current⑧
    link-type="dfn"}.

2.  If `node`{.variable} is
    [root](#concept-traversal-root){#ref-for-concept-traversal-root①③
    link-type="dfn"}, then return null.

3.  While true:

    1.  Let `sibling`{.variable} be `node`{.variable}'s [next
        sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①⑧
        link-type="dfn"} if `type`{.variable} is next, and
        `node`{.variable}'s [previous
        sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①③
        link-type="dfn"} if `type`{.variable} is previous.

    2.  While `sibling`{.variable} is non-null:

        1.  Set `node`{.variable} to `sibling`{.variable}.

        2.  Let `result`{.variable} be the result of
            [filtering](#concept-node-filter){#ref-for-concept-node-filter③
            link-type="dfn"} `node`{.variable} within
            `walker`{.variable}.

        3.  If `result`{.variable} is
            [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept④
            link-type="idl"}, then set `walker`{.variable}'s
            [current](#treewalker-current){#ref-for-treewalker-current⑨
            link-type="dfn"} to `node`{.variable} and return
            `node`{.variable}.

        4.  Set `sibling`{.variable} to `node`{.variable}'s [first
            child](#concept-tree-first-child){#ref-for-concept-tree-first-child⑧
            link-type="dfn"} if `type`{.variable} is next, and
            `node`{.variable}'s [last
            child](#concept-tree-last-child){#ref-for-concept-tree-last-child⑧
            link-type="dfn"} if `type`{.variable} is previous.

        5.  If `result`{.variable} is
            [`FILTER_REJECT`{.idl}](#dom-nodefilter-filter_reject){#ref-for-dom-nodefilter-filter_reject
            link-type="idl"} or `sibling`{.variable} is null, then set
            `sibling`{.variable} to `node`{.variable}'s [next
            sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling①⑨
            link-type="dfn"} if `type`{.variable} is next, and
            `node`{.variable}'s [previous
            sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①④
            link-type="dfn"} if `type`{.variable} is previous.

    3.  Set `node`{.variable} to `node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥⑦
        link-type="dfn"}.

    4.  If `node`{.variable} is null or `walker`{.variable}'s
        [root](#concept-traversal-root){#ref-for-concept-traversal-root①④
        link-type="dfn"}, then return null.

    5.  If the return value of
        [filtering](#concept-node-filter){#ref-for-concept-node-filter④
        link-type="dfn"} `node`{.variable} within `walker`{.variable} is
        [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept⑤
        link-type="idl"}, then return null.

The [`nextSibling()`]{#dom-treewalker-nextsibling .dfn .dfn-paneled
.idl-code dfn-for="TreeWalker" dfn-type="method" export=""} method steps
are to [traverse
siblings](#concept-traverse-siblings){#ref-for-concept-traverse-siblings
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪③
link-type="dfn"} and next.

The [`previousSibling()`]{#dom-treewalker-previoussibling .dfn
.dfn-paneled .idl-code dfn-for="TreeWalker" dfn-type="method" export=""}
method steps are to [traverse
siblings](#concept-traverse-siblings){#ref-for-concept-traverse-siblings①
link-type="dfn"} with
[this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪④
link-type="dfn"} and previous.

The [`previousNode()`]{#dom-treewalker-previousnode .dfn .dfn-paneled
.idl-code dfn-for="TreeWalker" dfn-type="method" export=""} method steps
are:

1.  Let `node`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪⑤
    link-type="dfn"}'s
    [current](#treewalker-current){#ref-for-treewalker-current①⓪
    link-type="dfn"}.

2.  While `node`{.variable} is not
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪⑥
    link-type="dfn"}'s
    [root](#concept-traversal-root){#ref-for-concept-traversal-root①⑤
    link-type="dfn"}:

    1.  Let `sibling`{.variable} be `node`{.variable}'s [previous
        sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①⑤
        link-type="dfn"}.

    2.  While `sibling`{.variable} is non-null:

        1.  Set `node`{.variable} to `sibling`{.variable}.

        2.  Let `result`{.variable} be the result of
            [filtering](#concept-node-filter){#ref-for-concept-node-filter⑤
            link-type="dfn"} `node`{.variable} within
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪⑦
            link-type="dfn"}.

        3.  While `result`{.variable} is not
            [`FILTER_REJECT`{.idl}](#dom-nodefilter-filter_reject){#ref-for-dom-nodefilter-filter_reject①
            link-type="idl"} and `node`{.variable} has a
            [child](#concept-tree-child){#ref-for-concept-tree-child⑧③
            link-type="dfn"}:

            1.  Set `node`{.variable} to `node`{.variable}'s [last
                child](#concept-tree-last-child){#ref-for-concept-tree-last-child⑨
                link-type="dfn"}.

            2.  Set `result`{.variable} to the result of
                [filtering](#concept-node-filter){#ref-for-concept-node-filter⑥
                link-type="dfn"} `node`{.variable} within
                [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪⑧
                link-type="dfn"}.

        4.  If `result`{.variable} is
            [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept⑥
            link-type="idl"}, then set
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③⓪⑨
            link-type="dfn"}'s
            [current](#treewalker-current){#ref-for-treewalker-current①①
            link-type="dfn"} to `node`{.variable} and return
            `node`{.variable}.

        5.  Set `sibling`{.variable} to `node`{.variable}'s [previous
            sibling](#concept-tree-previous-sibling){#ref-for-concept-tree-previous-sibling①⑥
            link-type="dfn"}.

    3.  If `node`{.variable} is
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①⓪
        link-type="dfn"}'s
        [root](#concept-traversal-root){#ref-for-concept-traversal-root①⑥
        link-type="dfn"} or `node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥⑧
        link-type="dfn"} is null, then return null.

    4.  Set `node`{.variable} to `node`{.variable}'s
        [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑥⑨
        link-type="dfn"}.

    5.  If the return value of
        [filtering](#concept-node-filter){#ref-for-concept-node-filter⑦
        link-type="dfn"} `node`{.variable} within
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①①
        link-type="dfn"} is
        [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept⑦
        link-type="idl"}, then set
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①②
        link-type="dfn"}'s
        [current](#treewalker-current){#ref-for-treewalker-current①②
        link-type="dfn"} to `node`{.variable} and return
        `node`{.variable}.

3.  Return null.

The [`nextNode()`]{#dom-treewalker-nextnode .dfn .dfn-paneled .idl-code
dfn-for="TreeWalker" dfn-type="method" export=""} method steps are:

1.  Let `node`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①③
    link-type="dfn"}'s
    [current](#treewalker-current){#ref-for-treewalker-current①③
    link-type="dfn"}.

2.  Let `result`{.variable} be
    [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept⑧
    link-type="idl"}.

3.  While true:

    1.  While `result`{.variable} is not
        [`FILTER_REJECT`{.idl}](#dom-nodefilter-filter_reject){#ref-for-dom-nodefilter-filter_reject②
        link-type="idl"} and `node`{.variable} has a
        [child](#concept-tree-child){#ref-for-concept-tree-child⑧④
        link-type="dfn"}:

        1.  Set `node`{.variable} to its [first
            child](#concept-tree-first-child){#ref-for-concept-tree-first-child⑨
            link-type="dfn"}.

        2.  Set `result`{.variable} to the result of
            [filtering](#concept-node-filter){#ref-for-concept-node-filter⑧
            link-type="dfn"} `node`{.variable} within
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①④
            link-type="dfn"}.

        3.  If `result`{.variable} is
            [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept⑨
            link-type="idl"}, then set
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①⑤
            link-type="dfn"}'s
            [current](#treewalker-current){#ref-for-treewalker-current①④
            link-type="dfn"} to `node`{.variable} and return
            `node`{.variable}.

    2.  Let `sibling`{.variable} be null.

    3.  Let `temporary`{.variable} be `node`{.variable}.

    4.  While `temporary`{.variable} is non-null:

        1.  If `temporary`{.variable} is
            [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①⑥
            link-type="dfn"}'s
            [root](#concept-traversal-root){#ref-for-concept-traversal-root①⑦
            link-type="dfn"}, then return null.

        2.  Set `sibling`{.variable} to `temporary`{.variable}'s [next
            sibling](#concept-tree-next-sibling){#ref-for-concept-tree-next-sibling②⓪
            link-type="dfn"}.

        3.  If `sibling`{.variable} is non-null, then set
            `node`{.variable} to `sibling`{.variable} and
            [break](https://infra.spec.whatwg.org/#iteration-break){#ref-for-iteration-break④
            link-type="dfn"}.

        4.  Set `temporary`{.variable} to `temporary`{.variable}'s
            [parent](#concept-tree-parent){#ref-for-concept-tree-parent⑦⓪
            link-type="dfn"}.

    5.  Set `result`{.variable} to the result of
        [filtering](#concept-node-filter){#ref-for-concept-node-filter⑨
        link-type="dfn"} `node`{.variable} within
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①⑦
        link-type="dfn"}.

    6.  If `result`{.variable} is
        [`FILTER_ACCEPT`{.idl}](#dom-nodefilter-filter_accept){#ref-for-dom-nodefilter-filter_accept①⓪
        link-type="idl"}, then set
        [this](https://webidl.spec.whatwg.org/#this){#ref-for-this③①⑧
        link-type="dfn"}'s
        [current](#treewalker-current){#ref-for-treewalker-current①⑤
        link-type="dfn"} to `node`{.variable} and return
        `node`{.variable}.

### [6.3. ]{.secno}[Interface [`NodeFilter`{.idl}](#callbackdef-nodefilter){#ref-for-callbackdef-nodefilter④ link-type="idl"}]{.content}[](#interface-nodefilter){.self-link} {#interface-nodefilter .heading .settled level="6.3"}

``` {.idl .highlight .def}
[Exposed=Window]
callback interface NodeFilter {
  // Constants for acceptNode()
  const unsigned short FILTER_ACCEPT = 1;
  const unsigned short FILTER_REJECT = 2;
  const unsigned short FILTER_SKIP = 3;

  // Constants for whatToShow
  const unsigned long SHOW_ALL = 0xFFFFFFFF;
  const unsigned long SHOW_ELEMENT = 0x1;
  const unsigned long SHOW_ATTRIBUTE = 0x2;
  const unsigned long SHOW_TEXT = 0x4;
  const unsigned long SHOW_CDATA_SECTION = 0x8;
  const unsigned long SHOW_ENTITY_REFERENCE = 0x10; // legacy
  const unsigned long SHOW_ENTITY = 0x20; // legacy
  const unsigned long SHOW_PROCESSING_INSTRUCTION = 0x40;
  const unsigned long SHOW_COMMENT = 0x80;
  const unsigned long SHOW_DOCUMENT = 0x100;
  const unsigned long SHOW_DOCUMENT_TYPE = 0x200;
  const unsigned long SHOW_DOCUMENT_FRAGMENT = 0x400;
  const unsigned long SHOW_NOTATION = 0x800; // legacy

  unsigned short acceptNode(Node node);
};
```

[`NodeFilter`{.idl}](#callbackdef-nodefilter){#ref-for-callbackdef-nodefilter⑤
link-type="idl"} objects can be used as
[filter](#concept-traversal-filter){#ref-for-concept-traversal-filter⑧
link-type="dfn"} for
[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①⑥
link-type="idl"} and
[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker①⓪ link-type="idl"}
objects and also provide constants for their
[whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow⑦
link-type="dfn"} bitmask. A
[`NodeFilter`{.idl}](#callbackdef-nodefilter){#ref-for-callbackdef-nodefilter⑥
link-type="idl"} object is typically implemented as a JavaScript
function.

These constants can be used as
[filter](#concept-traversal-filter){#ref-for-concept-traversal-filter⑨
link-type="dfn"} return value:

- [`FILTER_ACCEPT`]{#dom-nodefilter-filter_accept .dfn .dfn-paneled
  .idl-code dfn-for="NodeFilter" dfn-type="const" export=""} (1);
- [`FILTER_REJECT`]{#dom-nodefilter-filter_reject .dfn .dfn-paneled
  .idl-code dfn-for="NodeFilter" dfn-type="const" export=""} (2);
- [`FILTER_SKIP`]{#dom-nodefilter-filter_skip .dfn .dfn-paneled
  .idl-code dfn-for="NodeFilter" dfn-type="const" export=""} (3).

These constants can be used for
[whatToShow](#concept-traversal-whattoshow){#ref-for-concept-traversal-whattoshow⑧
link-type="dfn"}:

- [`SHOW_ALL`]{#dom-nodefilter-show_all .dfn .dfn-paneled .idl-code
  dfn-for="NodeFilter" dfn-type="const" export=""} (4294967295, FFFFFFFF
  in hexadecimal);
- [`SHOW_ELEMENT`]{#dom-nodefilter-show_element .dfn .dfn-paneled
  .idl-code dfn-for="NodeFilter" dfn-type="const" export=""} (1);
- [`SHOW_ATTRIBUTE`]{#dom-nodefilter-show_attribute .dfn .dfn-paneled
  .idl-code dfn-for="NodeFilter" dfn-type="const" export=""} (2);
- [`SHOW_TEXT`]{#dom-nodefilter-show_text .dfn .dfn-paneled .idl-code
  dfn-for="NodeFilter" dfn-type="const" export=""} (4);
- [`SHOW_CDATA_SECTION`]{#dom-nodefilter-show_cdata_section .dfn
  .dfn-paneled .idl-code dfn-for="NodeFilter" dfn-type="const"
  export=""} (8);
- [`SHOW_PROCESSING_INSTRUCTION`]{#dom-nodefilter-show_processing_instruction
  .dfn .dfn-paneled .idl-code dfn-for="NodeFilter" dfn-type="const"
  export=""} (64, 40 in hexadecimal);
- [`SHOW_COMMENT`]{#dom-nodefilter-show_comment .dfn .dfn-paneled
  .idl-code dfn-for="NodeFilter" dfn-type="const" export=""} (128, 80 in
  hexadecimal);
- [`SHOW_DOCUMENT`]{#dom-nodefilter-show_document .dfn .dfn-paneled
  .idl-code dfn-for="NodeFilter" dfn-type="const" export=""} (256, 100
  in hexadecimal);
- [`SHOW_DOCUMENT_TYPE`]{#dom-nodefilter-show_document_type .dfn
  .dfn-paneled .idl-code dfn-for="NodeFilter" dfn-type="const"
  export=""} (512, 200 in hexadecimal);
- [`SHOW_DOCUMENT_FRAGMENT`]{#dom-nodefilter-show_document_fragment .dfn
  .dfn-paneled .idl-code dfn-for="NodeFilter" dfn-type="const"
  export=""} (1024, 400 in hexadecimal).

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

## [8. ]{.secno}[XPath]{.content}[](#xpath){.self-link} {#xpath .heading .settled level="8"}

DOM Level 3 XPath defined an API for evaluating XPath 1.0 expressions.
These APIs are widely implemented, but have not been maintained. The
interface definitions are maintained here so that they can be updated
when Web IDL changes. Complete definitions of these APIs remain
necessary and such work is tracked and can be contributed to in
[whatwg/dom#67](https://github.com/whatwg/dom/issues/67).
[\[DOM-Level-3-XPath\]](#biblio-dom-level-3-xpath "Document Object Model (DOM) Level 3 XPath Specification"){link-type="biblio"}
[\[XPath\]](#biblio-xpath "XML Path Language (XPath) Version 1.0"){link-type="biblio"}
[\[WEBIDL\]](#biblio-webidl "Web IDL Standard"){link-type="biblio"}

### [8.1. ]{.secno}[Interface [`XPathResult`{.idl}](#xpathresult){#ref-for-xpathresult link-type="idl"}]{.content}[](#interface-xpathresult){.self-link} {#interface-xpathresult .heading .settled level="8.1"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XPathResult {
  const unsigned short ANY_TYPE = 0;
  const unsigned short NUMBER_TYPE = 1;
  const unsigned short STRING_TYPE = 2;
  const unsigned short BOOLEAN_TYPE = 3;
  const unsigned short UNORDERED_NODE_ITERATOR_TYPE = 4;
  const unsigned short ORDERED_NODE_ITERATOR_TYPE = 5;
  const unsigned short UNORDERED_NODE_SNAPSHOT_TYPE = 6;
  const unsigned short ORDERED_NODE_SNAPSHOT_TYPE = 7;
  const unsigned short ANY_UNORDERED_NODE_TYPE = 8;
  const unsigned short FIRST_ORDERED_NODE_TYPE = 9;

  readonly attribute unsigned short resultType;
  readonly attribute unrestricted double numberValue;
  readonly attribute DOMString stringValue;
  readonly attribute boolean booleanValue;
  readonly attribute Node? singleNodeValue;
  readonly attribute boolean invalidIteratorState;
  readonly attribute unsigned long snapshotLength;

  Node? iterateNext();
  Node? snapshotItem(unsigned long index);
};
```

### [8.2. ]{.secno}[Interface [`XPathExpression`{.idl}](#xpathexpression){#ref-for-xpathexpression link-type="idl"}]{.content}[](#interface-xpathexpression){.self-link} {#interface-xpathexpression .heading .settled level="8.2"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XPathExpression {
  // XPathResult.ANY_TYPE = 0
  XPathResult evaluate(Node contextNode, optional unsigned short type = 0, optional XPathResult? result = null);
};
```

### [8.3. ]{.secno}[Mixin [`XPathEvaluatorBase`{.idl}](#xpathevaluatorbase){#ref-for-xpathevaluatorbase link-type="idl"}]{.content}[](#mixin-xpathevaluatorbase){.self-link} {#mixin-xpathevaluatorbase .heading .settled level="8.3"}

``` {.idl .highlight .def}
callback interface XPathNSResolver {
  DOMString? lookupNamespaceURI(DOMString? prefix);
};

interface mixin XPathEvaluatorBase {
  [NewObject] XPathExpression createExpression(DOMString expression, optional XPathNSResolver? resolver = null);
  Node createNSResolver(Node nodeResolver); // legacy
  // XPathResult.ANY_TYPE = 0
  XPathResult evaluate(DOMString expression, Node contextNode, optional XPathNSResolver? resolver = null, optional unsigned short type = 0, optional XPathResult? result = null);
};
Document includes XPathEvaluatorBase;
```

The
[`createNSResolver(``nodeResolver`{.variable}`)`]{#dom-xpathevaluatorbase-creatensresolver
.dfn .dfn-paneled .idl-code dfn-for="XPathEvaluatorBase"
dfn-type="method" export=""} method steps are to return
`nodeResolver`{.variable}.

This method exists only for historical reasons.

### [8.4. ]{.secno}[Interface [`XPathEvaluator`{.idl}](#xpathevaluator){#ref-for-xpathevaluator link-type="idl"}]{.content}[](#interface-xpathevaluator){.self-link} {#interface-xpathevaluator .heading .settled level="8.4"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XPathEvaluator {
  constructor();
};

XPathEvaluator includes XPathEvaluatorBase;
```

For historical reasons you can both construct
[`XPathEvaluator`{.idl}](#xpathevaluator){#ref-for-xpathevaluator②
link-type="idl"} and access the same methods on
[`Document`{.idl}](#document){#ref-for-document②⑤ link-type="idl"}.

## [9. ]{.secno}[XSLT]{.content}[](#xslt){.self-link} {#xslt .heading .settled level="9"}

XSL Transformations (XSLT) is a language for transforming XML documents
into other XML documents. The APIs defined in this section have been
widely implemented, and are maintained here so that they can be updated
when Web IDL changes. Complete definitions of these APIs remain
necessary and such work is tracked and can be contributed to in
[whatwg/dom#181](https://github.com/whatwg/dom/issues/181).
[\[XSLT\]](#biblio-xslt "XSL Transformations (XSLT) Version 1.0"){link-type="biblio"}

### [9.1. ]{.secno}[Interface [`XSLTProcessor`{.idl}](#xsltprocessor){#ref-for-xsltprocessor link-type="idl"}]{.content}[](#interface-xsltprocessor){.self-link} {#interface-xsltprocessor .heading .settled level="9.1"}

``` {.idl .highlight .def}
[Exposed=Window]
interface XSLTProcessor {
  constructor();
  undefined importStylesheet(Node style);
  [CEReactions] DocumentFragment transformToFragment(Node source, Document output);
  [CEReactions] Document transformToDocument(Node source);
  undefined setParameter([LegacyNullToEmptyString] DOMString namespaceURI, DOMString localName, any value);
  any getParameter([LegacyNullToEmptyString] DOMString namespaceURI, DOMString localName);
  undefined removeParameter([LegacyNullToEmptyString] DOMString namespaceURI, DOMString localName);
  undefined clearParameters();
  undefined reset();
};
```

## [10. ]{.secno}[Security and privacy considerations]{.content}[](#security-and-privacy){.self-link} {#security-and-privacy .heading .settled level="10"}

There are no known security or privacy considerations for this standard.

## [11. ]{.secno}[Historical]{.content}[](#historical){.self-link} {#historical .heading .settled level="11"}

This standard used to contain several interfaces and interface members
that have been removed.

These interfaces have been removed:

- [`DOMConfiguration`]{#domconfiguration .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMError`]{#domerror .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMErrorHandler`]{#domerrorhandler .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMImplementationList`]{#domimplementationlist .dfn .dfn-paneled
  .idl-code dfn-type="interface" export=""}
- [`DOMImplementationSource`]{#domimplementationsource .dfn .dfn-paneled
  .idl-code dfn-type="interface" export=""}
- [`DOMLocator`]{#domlocator .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMObject`]{#domobject .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`DOMUserData`]{#domuserdata .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`Entity`]{#entity .dfn .dfn-paneled .idl-code dfn-type="interface"
  export=""}
- [`EntityReference`]{#entityreference .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`MutationEvent`]{#mutationevent .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`MutationNameEvent`]{#mutationnameevent .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`NameList`]{#namelist .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`Notation`]{#notation .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`RangeException`]{#rangeexception .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`TypeInfo`]{#typeinfo .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}
- [`UserDataHandler`]{#userdatahandler .dfn .dfn-paneled .idl-code
  dfn-type="interface" export=""}

And these interface members have been removed:

[`Attr`{.idl}](#attr){#ref-for-attr③⑨ link-type="idl"}

:   - [`schemaTypeInfo`]{#dom-attr-schematypeinfo .dfn .dfn-paneled
      .idl-code dfn-for="Attr" dfn-type="attribute" export=""}
    - [`isId`]{#dom-attr-isid .dfn .dfn-paneled .idl-code dfn-for="Attr"
      dfn-type="attribute" export=""}

[`Document`{.idl}](#document){#ref-for-document②⑧ link-type="idl"}

:   - [`createEntityReference()`]{#dom-document-createentityreference
      .dfn .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
      export=""}
    - [`xmlEncoding`]{#dom-document-xmlencoding .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="attribute" export=""}
    - [`xmlStandalone`]{#dom-document-xmlstandalone .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="attribute" export=""}
    - [`xmlVersion`]{#dom-document-xmlversion .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="attribute" export=""}
    - [`strictErrorChecking`]{#dom-document-stricterrorchecking .dfn
      .dfn-paneled .idl-code dfn-for="Document" dfn-type="attribute"
      export=""}
    - [`domConfig`]{#dom-document-domconfig .dfn .dfn-paneled .idl-code
      dfn-for="Document" dfn-type="attribute" export=""}
    - [`normalizeDocument()`]{#dom-document-normalizedocument .dfn
      .dfn-paneled .idl-code dfn-for="Document" dfn-type="method"
      export=""}
    - [`renameNode()`]{#dom-document-renamenode .dfn .dfn-paneled
      .idl-code dfn-for="Document" dfn-type="method" export=""}

[`DocumentType`{.idl}](#documenttype){#ref-for-documenttype②⑤ link-type="idl"}

:   - [`entities`]{#dom-documenttype-entities .dfn .dfn-paneled
      .idl-code dfn-for="DocumentType" dfn-type="attribute" export=""}
    - [`notations`]{#dom-documenttype-notations .dfn .dfn-paneled
      .idl-code dfn-for="DocumentType" dfn-type="attribute" export=""}
    - [`internalSubset`]{#dom-documenttype-internalsubset .dfn
      .dfn-paneled .idl-code dfn-for="DocumentType" dfn-type="attribute"
      export=""}

[`DOMImplementation`{.idl}](#domimplementation){#ref-for-domimplementation⑤ link-type="idl"}

:   - [`getFeature()`]{#dom-domimplementation-getfeature .dfn
      .dfn-paneled .idl-code dfn-for="DOMImplementation"
      dfn-type="method" export=""}

[`Element`{.idl}](#element){#ref-for-element⑤⑦ link-type="idl"}

:   - [`schemaTypeInfo`]{#dom-element-schematypeinfo .dfn .dfn-paneled
      .idl-code dfn-for="Element" dfn-type="attribute" export=""}
    - [`setIdAttribute()`]{#dom-element-setidattribute .dfn .dfn-paneled
      .idl-code dfn-for="Element" dfn-type="method" export=""}
    - [`setIdAttributeNS()`]{#dom-element-setidattributens .dfn
      .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
      export=""}
    - [`setIdAttributeNode()`]{#dom-element-setidattributenode .dfn
      .dfn-paneled .idl-code dfn-for="Element" dfn-type="method"
      export=""}

[`Node`{.idl}](#node){#ref-for-node①⓪⑧ link-type="idl"}

:   - [`isSupported`]{#dom-node-issupported .dfn .dfn-paneled .idl-code
      dfn-for="Node" dfn-type="attribute" export=""}
    - [`getFeature()`]{#dom-node-getfeature .dfn .dfn-paneled .idl-code
      dfn-for="Node" dfn-type="method" export=""}
    - [`getUserData()`]{#dom-node-getuserdata .dfn .dfn-paneled
      .idl-code dfn-for="Node" dfn-type="method" export=""}
    - [`setUserData()`]{#dom-node-setuserdata .dfn .dfn-paneled
      .idl-code dfn-for="Node" dfn-type="method" export=""}

[`NodeIterator`{.idl}](#nodeiterator){#ref-for-nodeiterator①⑦ link-type="idl"}

:   - [`expandEntityReferences`]{#dom-nodeiterator-expandentityreferences
      .dfn .dfn-paneled .idl-code dfn-for="NodeIterator"
      dfn-type="attribute" export=""}

[`Text`{.idl}](#text){#ref-for-text⑤⑨ link-type="idl"}

:   - [`isElementContentWhitespace`]{#dom-text-iselementcontentwhitespace
      .dfn .dfn-paneled .idl-code dfn-for="Text" dfn-type="attribute"
      export=""}
    - [`replaceWholeText()`]{#dom-text-replacewholetext .dfn
      .dfn-paneled .idl-code dfn-for="Text" dfn-type="method" export=""}

[`TreeWalker`{.idl}](#treewalker){#ref-for-treewalker①① link-type="idl"}

:   - [`expandEntityReferences`]{#dom-treewalker-expandentityreferences
      .dfn .dfn-paneled .idl-code dfn-for="TreeWalker"
      dfn-type="attribute" export=""}

## [Acknowledgments]{.content}[](#acks){.self-link} {#acks .no-num .heading .settled}

There have been a lot of people that have helped make DOM more
interoperable over the years and thereby furthered the goals of this
standard. Likewise many people have helped making this standard what it
is today.

With that, many thanks to Adam Klein, Adrian Bateman, Ahmid *snuggs*,
Alex Komoroske, Alex Russell, Alexey Shvayka, Andreas Kling, Andreu
Botella, Anthony Ramine, Arkadiusz Michalski, Arnaud Le Hors, Arun
Ranganathan, Benjamin Gruenbaum, Björn Höhrmann, Boris Zbarsky, Brandon
Payton, Brandon Slade, Brandon Wallace, Brian Kardell, C. Scott Ananian,
Cameron McCormack, Chris Dumez, Chris Paris, Chris Rebert, Cyrille Tuzi,
Dan Burzo, Daniel Clark, Daniel Glazman, Darien Maillet Valentine, Darin
Fisher, David Baron, David Bruant, David Flanagan, David Håsäther, David
Hyatt, Deepak Sherveghar, Dethe Elza, Dimitri Glazkov, Domenic Denicola,
Dominic Cooney, Dominique Hazaël-Massieux, Don Jordan, Doug Schepers,
Edgar Chen, Elisée Maurer, Elliott Sprehn, Emilio Cobos Álvarez, Eric
Bidelman, Erik Arvidsson, François Daoust, François Remy, Gary
Kacmarcik, Gavin Nicol, Giorgio Liscio, Glen Huang, Glenn Adams, Glenn
Maynard, Hajime Morrita, Harald Alvestrand, Hayato Ito, Henri Sivonen,
Hongchan Choi, Hunan Rostomyan, Ian Hickson, Igor Bukanov, Jacob Rossi,
Jake Archibald, Jake Verbaten, James Graham, James Greene, James M
Snell, James Robinson, Jeffrey Yasskin, Jens Lindström, Jeremy Davis,
Jesse McCarthy, Jinho Bang, João Eiras, Joe Kesselman, John Atkins, John
Dai, Jonas Sicking, Jonathan Kingston, Jonathan Robie, Joris van der
Wel, Joshua Bell, J. S. Choi, Jungkee Song, Justin Summerlin, Kagami
Sascha Rosylight, 呂康豪 (Kang-Hao Lu), 田村健人 (Kent TAMURA), Kevin J.
Sung, Kevin Sweeney, Kirill Topolyan, Koji Ishii, Lachlan Hunt, Lauren
Wood, Luca Casonato, Luke Zielinski, Magne Andersson, Majid Valipour,
Malte Ubl, Manish Goregaokar, Manish Tripathi, Marcos Caceres, Mark
Miller, Martijn van der Ven, Mason Freed, Mats Palmgren, Mounir Lamouri,
Michael Stramel, Michael™ Smith, Mike Champion, Mike Taylor, Mike West,
Nicolás Peña Moreno, Nidhi Jaju, Ojan Vafai, Oliver Nightingale, Olli
Pettay, Ondřej Žára, Peter Sharpe, Philip Jägenstedt, Philippe Le
Hégaret, Piers Wombwell, Pierre-Marie Dartus, prosody---Gab Vereable
Context(, Rafael Weinstein, Rakina Zata Amni, Richard Bradshaw, Rick
Byers, Rick Waldron, Robbert Broersma, Robin Berjon, Roland Steiner,
Rune [F.]{title="Fabulous"} Halvorsen, Russell Bicknell, Ruud
Steltenpool, Ryosuke Niwa, Sam Dutton, Sam Sneddon, Samuel Giles, Sanket
Joshi, Scott Haseley, Sebastian Mayr, Seo Sanghyeon, Sergey G. Grekhov,
Shiki Okasaka, Shinya Kawanaka, Simon Pieters, Simon Wülker, Stef
Busking, Steve Byrne, Stig Halvorsen, Tab Atkins, Takashi Sakamoto,
Takayoshi Kochi, Theresa O'Connor, Theodore Dubois, *timeless*, Timo
Tijhof, Tobie Langel, Tom Pixley, Travis Leithead, Trevor Rowbotham,
*triple-underscore*, Tristan Fraipont, Veli Şenol, Vidur Apparao, Warren
He, Xidorn Quan, Yash Handa, Yehuda Katz, Yoav Weiss, Yoichi Osato,
Yoshinori Sano, Yu Han, Yusuke Abe, and Zack Weinberg for being awesome!

This standard is written by [Anne van
Kesteren](https://annevankesteren.nl/){lang="nl"}
([Apple](https://www.apple.com/), <annevk@annevk.nl>) with substantial
contributions from Aryeh Gregor (<ayg@aryeh.name>) and Ms2ger
(<ms2ger@gmail.com>).

## [Intellectual property rights]{.content}[](#ipr){.self-link} {#ipr .no-num .heading .settled}

::: {fill-with="ipr"}
Part of the revision history of the integration points related to
[custom](#concept-element-custom){#ref-for-concept-element-custom⑧
link-type="dfn"} elements can be found in [the w3c/webcomponents
repository](https://github.com/w3c/webcomponents), which is available
under the [W3C Software and Document
License](https://www.w3.org/Consortium/Legal/2015/copyright-software-and-document).
:::

Copyright © WHATWG (Apple, Google, Mozilla, Microsoft). This work is
licensed under a [Creative Commons Attribution 4.0 International
License](https://creativecommons.org/licenses/by/4.0/){rel="license"}.
To the extent portions of it are incorporated into source code, such
portions in the source code are licensed under the [BSD 3-Clause
License](https://opensource.org/licenses/BSD-3-Clause){rel="license"}
instead.

This is the Living Standard. Those interested in the patent-review
version should view the [Living Standard Review
Draft](/review-drafts/2024-12/).
