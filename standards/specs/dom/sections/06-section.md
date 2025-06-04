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

