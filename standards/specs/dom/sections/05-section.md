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

