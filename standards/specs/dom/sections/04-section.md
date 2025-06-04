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

