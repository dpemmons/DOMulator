## [2]{.secno} Common infrastructure[](#infrastructure){.self-link} {#infrastructure}

This specification depends on Infra.
[\[INFRA\]](references.html#refsINFRA)

### [2.1]{.secno} Terminology[](#terminology){.self-link}

This specification refers to both HTML and XML attributes and IDL
attributes, often in the same context. When it is not clear which is
being referred to, they are referred to as [content attributes]{.dfn
lt="content attribute" export=""} for HTML and XML attributes, and [IDL
attributes]{.dfn} for those defined on IDL interfaces. Similarly, the
term \"properties\" is used for both JavaScript object properties and
CSS properties. When these are ambiguous they are qualified as [object
properties]{.dfn} and [CSS properties]{.dfn} respectively.

Generally, when the specification states that a feature applies to [the
HTML syntax](syntax.html#syntax){#terminology:syntax} or [the XML
syntax](xhtml.html#the-xhtml-syntax){#terminology:the-xhtml-syntax}, it
also includes the other. When a feature specifically only applies to one
of the two languages, it is called out by explicitly stating that it
does not apply to the other format, as in \"for HTML, \... (this does
not apply to XML)\".

This specification uses the term [document]{.dfn} to refer to any use of
HTML, ranging from short static documents to long essays or reports with
rich multimedia, as well as to fully-fledged interactive applications.
The term is used to refer both to
[`Document`{#terminology:document}](dom.html#document) objects and their
descendant DOM trees, and to serialized byte streams using the [HTML
syntax](syntax.html#syntax){#terminology:syntax-2} or the [XML
syntax](xhtml.html#the-xhtml-syntax){#terminology:the-xhtml-syntax-2},
depending on context.

In the context of the DOM structures, the terms [HTML
document](https://dom.spec.whatwg.org/#html-document){#terminology:html-documents
x-internal="html-documents"} and [XML
document](https://dom.spec.whatwg.org/#xml-document){#terminology:xml-documents
x-internal="xml-documents"} are used as defined in DOM, and refer
specifically to two different modes that
[`Document`{#terminology:document-2}](dom.html#document) objects can
find themselves in. [\[DOM\]](references.html#refsDOM) (Such uses are
always hyperlinked to their definition.)

In the context of byte streams, the term HTML document refers to
resources labeled as
[`text/html`{#terminology:text/html}](iana.html#text/html), and the term
XML document refers to resources labeled with an [XML MIME
type](https://mimesniff.spec.whatwg.org/#xml-mime-type){#terminology:xml-mime-type
x-internal="xml-mime-type"}.

------------------------------------------------------------------------

For simplicity, terms such as [shown]{.dfn}, [displayed]{.dfn}, and
[visible]{.dfn} might sometimes be used when referring to the way a
document is rendered to the user. These terms are not meant to imply a
visual medium; they must be considered to apply to other media in
equivalent ways.

#### [2.1.1]{.secno} Parallelism[](#parallelism){.self-link}

To run steps [in parallel]{#in-parallel .dfn export=""} means those
steps are to be run, one after another, at the same time as other logic
in the standard (e.g., at the same time as the [event
loop](webappapis.html#event-loop){#parallelism:event-loop}). This
standard does not define the precise mechanism by which this is
achieved, be it time-sharing cooperative multitasking, fibers, threads,
processes, using different hyperthreads, cores, CPUs, machines, etc. By
contrast, an operation that is to run [immediately]{#immediately .dfn}
must interrupt the currently running task, run itself, and then resume
the previously running task.

For guidance on writing specifications that leverage parallelism, see
[Dealing with the event loop from other
specifications](webappapis.html#event-loop-for-spec-authors).

To avoid race conditions between different [in
parallel](#in-parallel){#parallelism:in-parallel} algorithms that
operate on the same data, a [parallel
queue](#parallel-queue){#parallelism:parallel-queue} can be used.

A [parallel queue]{#parallel-queue .dfn export=""} represents a queue of
algorithm steps that must be run in series.

A [parallel queue](#parallel-queue){#parallelism:parallel-queue-2} has
an [algorithm queue]{#algorithm-queue .dfn} (a
[queue](https://infra.spec.whatwg.org/#queue){#parallelism:queue
x-internal="queue"}), initially empty.

To [enqueue steps]{#enqueue-the-following-steps .dfn
dfn-for="parallel queue" lt="enqueue steps|enqueue the following
  steps" export=""} to a [parallel
queue](#parallel-queue){#parallelism:parallel-queue-3},
[enqueue](https://infra.spec.whatwg.org/#queue-enqueue){#parallelism:enqueue
x-internal="enqueue"} the algorithm steps to the [parallel
queue](#parallel-queue){#parallelism:parallel-queue-4}\'s [algorithm
queue](#algorithm-queue){#parallelism:algorithm-queue}.

To [start a new parallel queue]{#starting-a-new-parallel-queue .dfn
lt="start a new parallel queue|starting a
  new parallel queue" export=""}, run the following steps:

1.  Let `parallelQueue`{.variable} be a new [parallel
    queue](#parallel-queue){#parallelism:parallel-queue-5}.

2.  Run the following steps [in
    parallel](#in-parallel){#parallelism:in-parallel-2}:

    1.  While true:

        1.  Let `steps`{.variable} be the result of
            [dequeuing](https://infra.spec.whatwg.org/#queue-dequeue){#parallelism:dequeue
            x-internal="dequeue"} from `parallelQueue`{.variable}\'s
            [algorithm
            queue](#algorithm-queue){#parallelism:algorithm-queue-2}.

        2.  If `steps`{.variable} is not nothing, then run
            `steps`{.variable}.

        3.  [Assert](https://infra.spec.whatwg.org/#assert){#parallelism:assert
            x-internal="assert"}: running `steps`{.variable} did not
            throw an exception, as steps running [in
            parallel](#in-parallel){#parallelism:in-parallel-3} are not
            allowed to throw.

        Implementations are not expected to implement this as a
        continuously running loop. Algorithms in standards are to be
        easy to understand and are not necessarily great for battery
        life or performance.

3.  Return `parallelQueue`{.variable}.

Steps running [in parallel](#in-parallel){#parallelism:in-parallel-4}
can themselves run other steps in [in
parallel](#in-parallel){#parallelism:in-parallel-5}. E.g., inside a
[parallel queue](#parallel-queue){#parallelism:parallel-queue-6} it can
be useful to run a series of steps in parallel with the queue.

::: example
Imagine a standard defined `nameList`{.variable} (a
[list](https://infra.spec.whatwg.org/#list){#parallelism:list
x-internal="list"}), along with a method to add a `name`{.variable} to
`nameList`{.variable}, unless `nameList`{.variable} already
[contains](https://infra.spec.whatwg.org/#list-contain){#parallelism:list-contains
x-internal="list-contains"} `name`{.variable}, in which case it rejects.

The following solution suffers from race conditions:

1.  Let `p`{.variable} be a new promise created in
    [this](https://webidl.spec.whatwg.org/#this){#parallelism:this
    x-internal="this"}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#parallelism:concept-relevant-realm}.

2.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#parallelism:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#parallelism:concept-relevant-global}.

3.  Run the following steps [in
    parallel](#in-parallel){#parallelism:in-parallel-6}:

    1.  If `nameList`{.variable}
        [contains](https://infra.spec.whatwg.org/#list-contain){#parallelism:list-contains-2
        x-internal="list-contains"} `name`{.variable}, then [queue a
        global
        task](webappapis.html#queue-a-global-task){#parallelism:queue-a-global-task}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#parallelism:dom-manipulation-task-source}
        given `global`{.variable} to reject `p`{.variable} with a
        [`TypeError`{#parallelism:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"},
        and abort these steps.

    2.  Do some potentially lengthy work.

    3.  [Append](https://infra.spec.whatwg.org/#list-append){#parallelism:list-append
        x-internal="list-append"} `name`{.variable} to
        `nameList`{.variable}.

    4.  [Queue a global
        task](webappapis.html#queue-a-global-task){#parallelism:queue-a-global-task-2}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#parallelism:dom-manipulation-task-source-2}
        given `global`{.variable} to resolve `p`{.variable} with
        undefined.

4.  Return `p`{.variable}.

Two invocations of the above could run simultaneously, meaning
`name`{.variable} isn\'t in `nameList`{.variable} during step 2.1, but
it *might be added* before step 2.3 runs, meaning `name`{.variable} ends
up in `nameList`{.variable} twice.

Parallel queues solve this. The standard would let
`nameListQueue`{.variable} be the result of [starting a new parallel
queue](#starting-a-new-parallel-queue){#parallelism:starting-a-new-parallel-queue},
then:

1.  Let `p`{.variable} be a new promise created in
    [this](https://webidl.spec.whatwg.org/#this){#parallelism:this-3
    x-internal="this"}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#parallelism:concept-relevant-realm-2}.

2.  Let `global`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#parallelism:this-4
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#parallelism:concept-relevant-global-2}.

3.  [[Enqueue the following
    steps](#enqueue-the-following-steps){#parallelism:enqueue-the-following-steps}
    to `nameListQueue`{.variable}:]{.mark}

    1.  If `nameList`{.variable}
        [contains](https://infra.spec.whatwg.org/#list-contain){#parallelism:list-contains-3
        x-internal="list-contains"} `name`{.variable}, then [queue a
        global
        task](webappapis.html#queue-a-global-task){#parallelism:queue-a-global-task-3}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#parallelism:dom-manipulation-task-source-3}
        given `global`{.variable} to reject `p`{.variable} with a
        [`TypeError`{#parallelism:typeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"},
        and abort these steps.

    2.  Do some potentially lengthy work.

    3.  [Append](https://infra.spec.whatwg.org/#list-append){#parallelism:list-append-2
        x-internal="list-append"} `name`{.variable} to
        `nameList`{.variable}.

    4.  [Queue a global
        task](webappapis.html#queue-a-global-task){#parallelism:queue-a-global-task-4}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#parallelism:dom-manipulation-task-source-4}
        given `global`{.variable} to resolve `p`{.variable} with
        undefined.

4.  Return `p`{.variable}.

The steps would now queue and the race is avoided.
:::

#### [2.1.2]{.secno} Resources[](#resources){.self-link}

The specification uses the term [supported]{.dfn} when referring to
whether a user agent has an implementation capable of decoding the
semantics of an external resource. A format or type is said to be
*supported* if the implementation can process an external resource of
that format or type without critical aspects of the resource being
ignored. Whether a specific resource is *supported* can depend on what
features of the resource\'s format are in use.

For example, a PNG image would be considered to be in a supported format
if its pixel data could be decoded and rendered, even if, unbeknownst to
the implementation, the image also contained animation data.

An MPEG-4 video file would not be considered to be in a supported format
if the compression format used was not supported, even if the
implementation could determine the dimensions of the movie from the
file\'s metadata.

What some specifications, in particular the HTTP specifications, refer
to as a *representation* is referred to in this specification as a
[resource]{.dfn}. [\[HTTP\]](references.html#refsHTTP)

A resource\'s [critical subresources]{#critical-subresources .dfn} are
those that the resource needs to have available to be correctly
processed. Which resources are considered critical or not is defined by
the specification that defines the resource\'s format.

For [CSS style
sheets](https://drafts.csswg.org/cssom/#css-style-sheet){#resources:css-style-sheet
x-internal="css-style-sheet"}, we tentatively define here that their
critical subresources are other style sheets imported via `@import`
rules, including those indirectly imported by other imported style
sheets.

This definition is not fully interoperable; furthermore, some user
agents seem to count resources like background images or web fonts as
critical subresources. Ideally, the CSS Working Group would define this;
see [w3c/csswg-drafts issue
#1088](https://github.com/w3c/csswg-drafts/issues/1088) to track
progress on that front.

#### [2.1.3]{.secno} XML compatibility[](#xml){.self-link} {#xml}

To ease migration from HTML to XML, user agents conforming to this
specification will place elements in HTML in the
[`http://www.w3.org/1999/xhtml`{#xml:html-namespace-2}](https://infra.spec.whatwg.org/#html-namespace){x-internal="html-namespace-2"}
namespace, at least for the purposes of the DOM and CSS. The term
\"[HTML elements]{#html-elements .dfn export=""}\" refers to any element
in that namespace, even in XML documents.

Except where otherwise stated, all elements defined or mentioned in this
specification are in the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#xml:html-namespace-2-2
x-internal="html-namespace-2"} (\"`http://www.w3.org/1999/xhtml`\"), and
all attributes defined or mentioned in this specification have no
namespace.

The term [element type]{#element-type .dfn} is used to refer to the set
of elements that have a given local name and namespace. For example,
[`button`{#xml:the-button-element}](form-elements.html#the-button-element)
elements are elements with the element type
[`button`{#xml:the-button-element-2}](form-elements.html#the-button-element),
meaning they have the local name \"`button`\" and (implicitly as defined
above) the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#xml:html-namespace-2-3
x-internal="html-namespace-2"}.

Attribute names are said to be [XML-compatible]{#xml-compatible .dfn} if
they match the
[`Name`{#xml:xml-name}](https://www.w3.org/TR/xml/#NT-Name){x-internal="xml-name"}
production defined in XML and they contain no U+003A COLON characters
(:). [\[XML\]](references.html#refsXML)

#### [2.1.4]{.secno} DOM trees[](#dom-trees){.self-link}

When it is stated that some element or attribute is [ignored]{#ignore
.dfn}, or treated as some other value, or handled as if it was something
else, this refers only to the processing of the node after it is in the
DOM. A user agent must not mutate the DOM in such situations.

A content attribute is said to [change]{.dfn} value only if its new
value is different than its previous value; setting an attribute to a
value it already has does not change it.

The term [empty]{.dfn}, when used for an attribute value,
[`Text`{#dom-trees:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node, or string, means that the
[length](https://infra.spec.whatwg.org/#string-length){#dom-trees:length
x-internal="length"} of the text is zero (i.e., not even containing
[controls](https://infra.spec.whatwg.org/#control){#dom-trees:control
x-internal="control"} or U+0020 SPACE).

An HTML element can have specific [HTML element insertion
steps]{#html-element-insertion-steps .dfn}, [HTML element
post-connection steps]{#html-element-post-connection-steps .dfn}, [HTML
element removing steps]{#html-element-removing-steps .dfn}, and [HTML
element moving steps]{#html-element-moving-steps .dfn} all defined for
the element\'s [local
name](https://dom.spec.whatwg.org/#concept-element-local-name){#dom-trees:concept-element-local-name
x-internal="concept-element-local-name"}.

The [insertion
steps](https://dom.spec.whatwg.org/#concept-node-insert-ext){#dom-trees:concept-node-insert-ext
x-internal="concept-node-insert-ext"} for the HTML Standard, given
`insertedNode`{.variable}, are defined as the following:

1.  If `insertedNode`{.variable} is an element whose
    [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#dom-trees:concept-element-namespace
    x-internal="concept-element-namespace"} is the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#dom-trees:html-namespace-2
    x-internal="html-namespace-2"}, and this standard defines [HTML
    element insertion
    steps](#html-element-insertion-steps){#dom-trees:html-element-insertion-steps}
    for `insertedNode`{.variable}\'s [local
    name](https://dom.spec.whatwg.org/#concept-element-local-name){#dom-trees:concept-element-local-name-2
    x-internal="concept-element-local-name"}, then run the corresponding
    [HTML element insertion
    steps](#html-element-insertion-steps){#dom-trees:html-element-insertion-steps-2}
    given `insertedNode`{.variable}.

2.  If `insertedNode`{.variable} is a [form-associated
    element](forms.html#form-associated-element){#dom-trees:form-associated-element}
    or the ancestor of a [form-associated
    element](forms.html#form-associated-element){#dom-trees:form-associated-element-2},
    then:

    1.  If the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-3}\'s
        [parser inserted
        flag](form-control-infrastructure.html#parser-inserted-flag){#dom-trees:parser-inserted-flag}
        is set, then return.

    2.  [Reset the form
        owner](form-control-infrastructure.html#reset-the-form-owner){#dom-trees:reset-the-form-owner}
        of the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-4}.

3.  If `insertedNode`{.variable} is an
    [`Element`{#dom-trees:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    that is not on the [stack of open
    elements](parsing.html#stack-of-open-elements){#dom-trees:stack-of-open-elements}
    of an [HTML
    parser](parsing.html#html-parser){#dom-trees:html-parser}, then
    [process internal resource
    links](links.html#process-internal-resource-links){#dom-trees:process-internal-resource-links}
    given `insertedNode`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#dom-trees:node-document
    x-internal="node-document"}.

The [post-connection
steps](https://dom.spec.whatwg.org/#concept-node-post-connection-ext){#dom-trees:concept-node-post-connection-ext
x-internal="concept-node-post-connection-ext"} for the HTML Standard,
given `insertedNode`{.variable}, are defined as the following:

1.  If `insertedNode`{.variable} is an element whose
    [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#dom-trees:concept-element-namespace-2
    x-internal="concept-element-namespace"} is the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#dom-trees:html-namespace-2-2
    x-internal="html-namespace-2"}, and this standard defines [HTML
    element post-connection
    steps](#html-element-post-connection-steps){#dom-trees:html-element-post-connection-steps}
    for `insertedNode`{.variable}\'s [local
    name](https://dom.spec.whatwg.org/#concept-element-local-name){#dom-trees:concept-element-local-name-3
    x-internal="concept-element-local-name"}, then run the corresponding
    [HTML element post-connection
    steps](#html-element-post-connection-steps){#dom-trees:html-element-post-connection-steps-2}
    given `insertedNode`{.variable}.

The [removing
steps](https://dom.spec.whatwg.org/#concept-node-remove-ext){#dom-trees:concept-node-remove-ext
x-internal="concept-node-remove-ext"} for the HTML Standard, given
`removedNode`{.variable} and `oldParent`{.variable}, are defined as the
following:

1.  Let `document`{.variable} be `removedNode`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#dom-trees:node-document-2
    x-internal="node-document"}.

2.  ::: {#node-remove-focus-fixup}
    If `document`{.variable}\'s [focused
    area](interaction.html#focused-area-of-the-document){#dom-trees:focused-area-of-the-document}
    is `removedNode`{.variable}, then set `document`{.variable}\'s
    [focused
    area](interaction.html#focused-area-of-the-document){#dom-trees:focused-area-of-the-document-2}
    to `document`{.variable}\'s
    [viewport](https://drafts.csswg.org/css2/#viewport){#dom-trees:viewport
    x-internal="viewport"}, and set `document`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#dom-trees:concept-relevant-global}\'s
    [navigation
    API](nav-history-apis.html#window-navigation-api){#dom-trees:window-navigation-api}\'s
    [focus changed during ongoing
    navigation](nav-history-apis.html#focus-changed-during-ongoing-navigation){#dom-trees:focus-changed-during-ongoing-navigation}
    to false.

    This does *not* perform the [unfocusing
    steps](interaction.html#unfocusing-steps){#dom-trees:unfocusing-steps},
    [focusing
    steps](interaction.html#focusing-steps){#dom-trees:focusing-steps},
    or [focus update
    steps](interaction.html#focus-update-steps){#dom-trees:focus-update-steps},
    and thus no [`blur`{#dom-trees:event-blur}](indices.html#event-blur)
    or [`change`{#dom-trees:event-change}](indices.html#event-change)
    events are fired.
    :::

3.  If `removedNode`{.variable} is an element whose
    [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#dom-trees:concept-element-namespace-3
    x-internal="concept-element-namespace"} is the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#dom-trees:html-namespace-2-3
    x-internal="html-namespace-2"}, and this standard defines [HTML
    element removing
    steps](#html-element-removing-steps){#dom-trees:html-element-removing-steps}
    for `removedNode`{.variable}\'s [local
    name](https://dom.spec.whatwg.org/#concept-element-local-name){#dom-trees:concept-element-local-name-4
    x-internal="concept-element-local-name"}, then run the corresponding
    [HTML element removing
    steps](#html-element-removing-steps){#dom-trees:html-element-removing-steps-2}
    given `removedNode`{.variable} and `oldParent`{.variable}.

4.  If `removedNode`{.variable} is a [form-associated
    element](forms.html#form-associated-element){#dom-trees:form-associated-element-5}
    or the ancestor of a [form-associated
    element](forms.html#form-associated-element){#dom-trees:form-associated-element-6},
    then:

    1.  If the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-7}
        has a [form
        owner](form-control-infrastructure.html#form-owner){#dom-trees:form-owner}
        and the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-8}
        and its [form
        owner](form-control-infrastructure.html#form-owner){#dom-trees:form-owner-2}
        are no longer in the same
        [tree](https://dom.spec.whatwg.org/#concept-tree){#dom-trees:tree
        x-internal="tree"}, then [reset the form
        owner](form-control-infrastructure.html#reset-the-form-owner){#dom-trees:reset-the-form-owner-2}
        of the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-9}.

5.  If `removedNode`{.variable}\'s
    [`popover`{#dom-trees:attr-popover}](popover.html#attr-popover)
    attribute is not in the [No Popover
    State](popover.html#attr-popover-none-state){#dom-trees:attr-popover-none-state},
    then run the [hide popover
    algorithm](popover.html#hide-popover-algorithm){#dom-trees:hide-popover-algorithm}
    given `removedNode`{.variable}, false, false, false, and null.

The [moving
steps](https://dom.spec.whatwg.org/#concept-node-move-ext){#dom-trees:concept-node-move-ext
x-internal="concept-node-move-ext"} for the HTML Standard, given
`movedNode`{.variable}, are defined as the following:

1.  If `movedNode`{.variable} is an element whose
    [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#dom-trees:concept-element-namespace-4
    x-internal="concept-element-namespace"} is the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#dom-trees:html-namespace-2-4
    x-internal="html-namespace-2"}, and this standard defines [HTML
    element moving
    steps](#html-element-moving-steps){#dom-trees:html-element-moving-steps}
    for `movedNode`{.variable}\'s [local
    name](https://dom.spec.whatwg.org/#concept-element-local-name){#dom-trees:concept-element-local-name-5
    x-internal="concept-element-local-name"}, then run the corresponding
    [HTML element moving
    steps](#html-element-moving-steps){#dom-trees:html-element-moving-steps-2}
    given `movedNode`{.variable}.

2.  If `movedNode`{.variable} is a [form-associated
    element](forms.html#form-associated-element){#dom-trees:form-associated-element-10}
    or the ancestor of a [form-associated
    element](forms.html#form-associated-element){#dom-trees:form-associated-element-11},
    then:

    1.  If the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-12}
        has a [form
        owner](form-control-infrastructure.html#form-owner){#dom-trees:form-owner-3}
        and the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-13}
        and its [form
        owner](form-control-infrastructure.html#form-owner){#dom-trees:form-owner-4}
        are no longer in the same
        [tree](https://dom.spec.whatwg.org/#concept-tree){#dom-trees:tree-2
        x-internal="tree"}, then [reset the form
        owner](form-control-infrastructure.html#reset-the-form-owner){#dom-trees:reset-the-form-owner-3}
        of the [form-associated
        element](forms.html#form-associated-element){#dom-trees:form-associated-element-14}.

A [node is inserted into a document]{#insert-an-element-into-a-document
.dfn lt="inserted into a document|node is inserted into a document"
export=""} when the [insertion
steps](https://dom.spec.whatwg.org/#concept-node-insert-ext){#dom-trees:concept-node-insert-ext-2
x-internal="concept-node-insert-ext"} are invoked with it as the
argument and it is now [in a document
tree](https://dom.spec.whatwg.org/#in-a-document-tree){#dom-trees:in-a-document-tree
x-internal="in-a-document-tree"}. Analogously, a [node is removed from a
document]{#remove-an-element-from-a-document .dfn lt="removed
  from a document|node is removed from a document" export=""} when the
[removing
steps](https://dom.spec.whatwg.org/#concept-node-remove-ext){#dom-trees:concept-node-remove-ext-2
x-internal="concept-node-remove-ext"} are invoked with it as the
argument and it is now no longer [in a document
tree](https://dom.spec.whatwg.org/#in-a-document-tree){#dom-trees:in-a-document-tree-2
x-internal="in-a-document-tree"}.

A node [becomes connected]{#becomes-connected .dfn export=""} when the
[insertion
steps](https://dom.spec.whatwg.org/#concept-node-insert-ext){#dom-trees:concept-node-insert-ext-3
x-internal="concept-node-insert-ext"} are invoked with it as the
argument and it is now
[connected](https://dom.spec.whatwg.org/#connected){#dom-trees:connected
x-internal="connected"}. Analogously, a node [becomes
disconnected]{#becomes-disconnected .dfn lt="become disconnected"
export=""} when the [removing
steps](https://dom.spec.whatwg.org/#concept-node-remove-ext){#dom-trees:concept-node-remove-ext-3
x-internal="concept-node-remove-ext"} are invoked with it as the
argument and it is now no longer
[connected](https://dom.spec.whatwg.org/#connected){#dom-trees:connected-2
x-internal="connected"}.

A node is [browsing-context connected]{#browsing-context-connected .dfn
export=""} when it is
[connected](https://dom.spec.whatwg.org/#connected){#dom-trees:connected-3
x-internal="connected"} and its [shadow-including
root](https://dom.spec.whatwg.org/#concept-shadow-including-root){#dom-trees:shadow-including-root
x-internal="shadow-including-root"}\'s [browsing
context](document-sequences.html#concept-document-bc){#dom-trees:concept-document-bc}
is non-null. A node [becomes browsing-context
connected]{#becomes-browsing-context-connected .dfn lt="become
  browsing-context connected" export=""} when the [insertion
steps](https://dom.spec.whatwg.org/#concept-node-insert-ext){#dom-trees:concept-node-insert-ext-4
x-internal="concept-node-insert-ext"} are invoked with it as the
argument and it is now [browsing-context
connected](#browsing-context-connected){#dom-trees:browsing-context-connected}.
A node [becomes browsing-context
disconnected]{#becomes-browsing-context-disconnected .dfn
lt="become browsing-context
  disconnected" export=""} either when the [removing
steps](https://dom.spec.whatwg.org/#concept-node-remove-ext){#dom-trees:concept-node-remove-ext-4
x-internal="concept-node-remove-ext"} are invoked with it as the
argument and it is now no longer [browsing-context
connected](#browsing-context-connected){#dom-trees:browsing-context-connected-2},
or when its [shadow-including
root](https://dom.spec.whatwg.org/#concept-shadow-including-root){#dom-trees:shadow-including-root-2
x-internal="shadow-including-root"}\'s [browsing
context](document-sequences.html#concept-document-bc){#dom-trees:concept-document-bc-2}
becomes null.

#### [2.1.5]{.secno} Scripting[](#scripting-2){.self-link} {#scripting-2}

The construction \"a `Foo` object\", where `Foo` is actually an
interface, is sometimes used instead of the more accurate \"an object
implementing the interface `Foo`\".

An IDL attribute is said to be [getting]{.dfn} when its value is being
retrieved (e.g. by author script), and is said to be [setting]{.dfn}
when a new value is assigned to it.

If a DOM object is said to be [live]{#live .dfn}, then the attributes
and methods on that object must operate on the actual underlying data,
not a snapshot of the data.

#### [2.1.6]{.secno} Plugins[](#plugins){.self-link}

The term [plugin]{#plugin .dfn} refers to an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#plugins:implementation-defined
x-internal="implementation-defined"} set of content handlers used by the
user agent that can take part in the user agent\'s rendering of a
[`Document`{#plugins:document}](dom.html#document) object, but that
neither act as [child
navigables](document-sequences.html#child-navigable){#plugins:child-navigable}
of the [`Document`{#plugins:document-2}](dom.html#document) nor
introduce any
[`Node`{#plugins:node}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
objects to the [`Document`{#plugins:document-3}](dom.html#document)\'s
DOM.

Typically such content handlers are provided by third parties, though a
user agent can also designate built-in content handlers as plugins.

A user agent must not consider the types
[`text/plain`{#plugins:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}
and
[`application/octet-stream`{#plugins:application/octet-stream}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"}
as having a registered [plugin](#plugin){#plugins:plugin}.

One example of a plugin would be a PDF viewer that is instantiated in a
[navigable](document-sequences.html#navigable){#plugins:navigable} when
the user navigates to a PDF file. This would count as a plugin
regardless of whether the party that implemented the PDF viewer
component was the same as that which implemented the user agent itself.
However, a PDF viewer application that launches separate from the user
agent (as opposed to using the same interface) is not a plugin by this
definition.

This specification does not define a mechanism for interacting with
plugins, as it is expected to be user-agent- and platform-specific. Some
UAs might opt to support a plugin mechanism such as the Netscape Plugin
API; others might use remote content converters or have built-in support
for certain types. Indeed, this specification doesn\'t require user
agents to support plugins at all. [\[NPAPI\]](references.html#refsNPAPI)

Browsers should take extreme care when interacting with external content
intended for [plugins](#plugin){#plugins:plugin-2}. When third-party
software is run with the same privileges as the user agent itself,
vulnerabilities in the third-party software become as dangerous as those
in the user agent.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#plugins:tracking-vector
.tracking-vector x-internal="tracking-vector"} Since different users
having different sets of [plugins](#plugin){#plugins:plugin-3} provides
a tracking vector that increases the chances of users being uniquely
identified, user agents are encouraged to support the exact same set of
[plugins](#plugin){#plugins:plugin-4} for each user.

#### [2.1.7]{.secno} Character encodings[](#encoding-terminology){.self-link} {#encoding-terminology}

A [[character
encoding](https://encoding.spec.whatwg.org/#encoding)]{#encoding .dfn},
or just *encoding* where that is not ambiguous, is a defined way to
convert between byte streams and Unicode strings, as defined in
Encoding. An
[encoding](https://encoding.spec.whatwg.org/#encoding){#encoding-terminology:encoding
x-internal="encoding"} has an [[encoding
name](https://encoding.spec.whatwg.org/#name)]{#encoding-name .dfn} and
one or more [[encoding
labels](https://encoding.spec.whatwg.org/#label)]{#encoding-label .dfn},
referred to as the encoding\'s *name* and *labels* in the Encoding
standard. [\[ENCODING\]](references.html#refsENCODING)

#### [2.1.8]{.secno} Conformance classes[](#conformance-classes){.self-link}

This specification describes the conformance criteria for user agents
(relevant to implementers) and documents (relevant to authors and
authoring tool implementers).

[Conforming documents]{#conforming-documents .dfn} are those that comply
with all the conformance criteria for documents. For readability, some
of these conformance requirements are phrased as conformance
requirements on authors; such requirements are implicitly requirements
on documents: by definition, all documents are assumed to have had an
author. (In some cases, that author may itself be a user agent --- such
user agents are subject to additional rules, as explained below.)

For example, if a requirement states that \"authors must not use the
`foobar` element\", it would imply that documents are not allowed to
contain elements named `foobar`.

There is no implied relationship between document conformance
requirements and implementation conformance requirements. User agents
are not free to handle non-conformant documents as they please; the
processing model described in this specification applies to
implementations regardless of the conformity of the input documents.

User agents fall into several (overlapping) categories with different
conformance requirements.

Web browsers and other interactive user agents

:   Web browsers that support [the XML
    syntax](xhtml.html#the-xhtml-syntax){#conformance-classes:the-xhtml-syntax}
    must process elements and attributes from the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#conformance-classes:html-namespace-2
    x-internal="html-namespace-2"} found in XML documents as described
    in this specification, so that users can interact with them, unless
    the semantics of those elements have been overridden by other
    specifications.

    A conforming web browser would, upon finding a
    [`script`{#conformance-classes:the-script-element}](scripting.html#the-script-element)
    element in an XML document, execute the script contained in that
    element. However, if the element is found within a transformation
    expressed in XSLT (assuming the user agent also supports XSLT), then
    the processor would instead treat the
    [`script`{#conformance-classes:the-script-element-2}](scripting.html#the-script-element)
    element as an opaque element that forms part of the transform.

    Web browsers that support [the HTML
    syntax](syntax.html#syntax){#conformance-classes:syntax} must
    process documents labeled with an [HTML MIME
    type](https://mimesniff.spec.whatwg.org/#html-mime-type){#conformance-classes:html-mime-type
    x-internal="html-mime-type"} as described in this specification, so
    that users can interact with them.

    User agents that support scripting must also be conforming
    implementations of the IDL fragments in this specification, as
    described in Web IDL. [\[WEBIDL\]](references.html#refsWEBIDL)

    Unless explicitly stated, specifications that override the semantics
    of HTML elements do not override the requirements on DOM objects
    representing those elements. For example, the
    [`script`{#conformance-classes:the-script-element-3}](scripting.html#the-script-element)
    element in the example above would still implement the
    [`HTMLScriptElement`{#conformance-classes:htmlscriptelement}](scripting.html#htmlscriptelement)
    interface.

Non-interactive presentation user agents

:   User agents that process HTML and XML documents purely to render
    non-interactive versions of them must comply to the same conformance
    criteria as web browsers, except that they are exempt from
    requirements regarding user interaction.

    Typical examples of non-interactive presentation user agents are
    printers (static UAs) and overhead displays (dynamic UAs). It is
    expected that most static non-interactive presentation user agents
    will also opt to [lack scripting support](#non-scripted).

    A non-interactive but dynamic presentation UA would still execute
    scripts, allowing forms to be dynamically submitted, and so forth.
    However, since the concept of \"focus\" is irrelevant when the user
    cannot interact with the document, the UA would not need to support
    any of the focus-related DOM APIs.

Visual user agents that support the suggested default rendering

:   User agents, whether interactive or not, may be designated (possibly
    as a user option) as supporting the suggested default rendering
    defined by this specification.

    This is not required. In particular, even user agents that do
    implement the suggested default rendering are encouraged to offer
    settings that override this default to improve the experience for
    the user, e.g. changing the color contrast, using different focus
    styles, or otherwise making the experience more accessible and
    usable to the user.

    User agents that are designated as supporting the suggested default
    rendering must, while so designated, implement the rules [the
    Rendering section](rendering.html#rendering) defines as the behavior
    that user agents are *expected* to implement.

User agents with no scripting support

:   Implementations that do not support scripting (or which have their
    scripting features disabled entirely) are exempt from supporting the
    events and DOM interfaces mentioned in this specification. For the
    parts of this specification that are defined in terms of an events
    model or in terms of the DOM, such user agents must still act as if
    events and the DOM were supported.

    Scripting can form an integral part of an application. Web browsers
    that do not support scripting, or that have scripting disabled,
    might be unable to fully convey the author\'s intent.

Conformance checkers

:   Conformance checkers must verify that a document conforms to the
    applicable conformance criteria described in this specification.
    Automated conformance checkers are exempt from detecting errors that
    require interpretation of the author\'s intent (for example, while a
    document is non-conforming if the content of a
    [`blockquote`{#conformance-classes:the-blockquote-element}](grouping-content.html#the-blockquote-element)
    element is not a quote, conformance checkers running without the
    input of human judgement do not have to check that
    [`blockquote`{#conformance-classes:the-blockquote-element-2}](grouping-content.html#the-blockquote-element)
    elements only contain quoted material).

    Conformance checkers must check that the input document conforms
    when parsed without a [browsing
    context](document-sequences.html#concept-document-bc){#conformance-classes:concept-document-bc}
    (meaning that no scripts are run, and that the parser\'s [scripting
    flag](parsing.html#scripting-flag){#conformance-classes:scripting-flag}
    is disabled), and should also check that the input document conforms
    when parsed with a [browsing
    context](document-sequences.html#concept-document-bc){#conformance-classes:concept-document-bc-2}
    in which scripts execute, and that the scripts never cause
    non-conforming states to occur other than transiently during script
    execution itself. (This is only a \"SHOULD\" and not a \"MUST\"
    requirement because it has been proven to be impossible.
    [\[COMPUTABLE\]](references.html#refsCOMPUTABLE))

    The term \"HTML validator\" can be used to refer to a conformance
    checker that itself conforms to the applicable requirements of this
    specification.

    ::: note
    XML DTDs cannot express all the conformance requirements of this
    specification. Therefore, a validating XML processor and a DTD
    cannot constitute a conformance checker. Also, since neither of the
    two authoring formats defined in this specification are applications
    of SGML, a validating SGML system cannot constitute a conformance
    checker either.

    To put it another way, there are three types of conformance
    criteria:

    1.  Criteria that can be expressed in a DTD.
    2.  Criteria that cannot be expressed by a DTD, but can still be
        checked by a machine.
    3.  Criteria that can only be checked by a human.

    A conformance checker must check for the first two. A simple
    DTD-based validator only checks for the first class of errors and is
    therefore not a conforming conformance checker according to this
    specification.
    :::

Data mining tools

:   Applications and tools that process HTML and XML documents for
    reasons other than to either render the documents or check them for
    conformance should act in accordance with the semantics of the
    documents that they process.

    A tool that generates [document
    outlines](sections.html#outline){#conformance-classes:outline} but
    increases the nesting level for each paragraph and does not increase
    the nesting level for
    [headings](sections.html#concept-heading){#conformance-classes:concept-heading}
    would not be conforming.

Authoring tools and markup generators

:   Authoring tools and markup generators must generate [conforming
    documents](#conforming-documents){#conformance-classes:conforming-documents}.
    Conformance criteria that apply to authors also apply to authoring
    tools, where appropriate.

    Authoring tools are exempt from the strict requirements of using
    elements only for their specified purpose, but only to the extent
    that authoring tools are not yet able to determine author intent.
    However, authoring tools must not automatically misuse elements or
    encourage their users to do so.

    For example, it is not conforming to use an
    [`address`{#conformance-classes:the-address-element}](sections.html#the-address-element)
    element for arbitrary contact information; that element can only be
    used for marking up contact information for its nearest
    [`article`{#conformance-classes:the-article-element}](sections.html#the-article-element)
    or
    [`body`{#conformance-classes:the-body-element}](sections.html#the-body-element)
    element ancestor. However, since an authoring tool is likely unable
    to determine the difference, an authoring tool is exempt from that
    requirement. This does not mean, though, that authoring tools can
    use
    [`address`{#conformance-classes:the-address-element-2}](sections.html#the-address-element)
    elements for any block of italics text (for instance); it just means
    that the authoring tool doesn\'t have to verify that when the user
    uses a tool for inserting contact information for an
    [`article`{#conformance-classes:the-article-element-2}](sections.html#the-article-element)
    element, that the user really is doing that and not inserting
    something else instead.

    In terms of conformance checking, an editor has to output documents
    that conform to the same extent that a conformance checker will
    verify.

    When an authoring tool is used to edit a non-conforming document, it
    may preserve the conformance errors in sections of the document that
    were not edited during the editing session (i.e. an editing tool is
    allowed to round-trip erroneous content). However, an authoring tool
    must not claim that the output is conformant if errors have been so
    preserved.

    Authoring tools are expected to come in two broad varieties: tools
    that work from structure or semantic data, and tools that work on a
    What-You-See-Is-What-You-Get media-specific editing basis (WYSIWYG).

    The former is the preferred mechanism for tools that author HTML,
    since the structure in the source information can be used to make
    informed choices regarding which HTML elements and attributes are
    most appropriate.

    However, WYSIWYG tools are legitimate. WYSIWYG tools should use
    elements they know are appropriate, and should not use elements that
    they do not know to be appropriate. This might in certain extreme
    cases mean limiting the use of flow elements to just a few elements,
    like
    [`div`{#conformance-classes:the-div-element}](grouping-content.html#the-div-element),
    [`b`{#conformance-classes:the-b-element}](text-level-semantics.html#the-b-element),
    [`i`{#conformance-classes:the-i-element}](text-level-semantics.html#the-i-element),
    and
    [`span`{#conformance-classes:the-span-element}](text-level-semantics.html#the-span-element)
    and making liberal use of the
    [`style`{#conformance-classes:attr-style}](dom.html#attr-style)
    attribute.

    All authoring tools, whether WYSIWYG or not, should make a best
    effort attempt at enabling users to create well-structured,
    semantically rich, media-independent content.

For compatibility with existing content and prior specifications, this
specification describes two authoring formats: one based on
[XML](xhtml.html#the-xhtml-syntax){#conformance-classes:the-xhtml-syntax-2},
and one using a [custom format](syntax.html#writing) inspired by SGML
(referred to as [the HTML
syntax](syntax.html#syntax){#conformance-classes:syntax-2}).
Implementations must support at least one of these two formats, although
supporting both is encouraged.

Some conformance requirements are phrased as requirements on elements,
attributes, methods or objects. Such requirements fall into two
categories: those describing content model restrictions, and those
describing implementation behavior. Those in the former category are
requirements on documents and authoring tools. Those in the second
category are requirements on user agents. Similarly, some conformance
requirements are phrased as requirements on authors; such requirements
are to be interpreted as conformance requirements on the documents that
authors produce. (In other words, this specification does not
distinguish between conformance criteria on authors and conformance
criteria on documents.)

#### [2.1.9]{.secno} Dependencies[](#dependencies){.self-link}

::: {noexport=""}
This specification relies on several other underlying specifications.

Infra

:   The following terms are defined in Infra:
    [\[INFRA\]](references.html#refsINFRA)

    - The general iteration terms
      [[while](https://infra.spec.whatwg.org/#iteration-while)]{#while
      .dfn},
      [[continue](https://infra.spec.whatwg.org/#iteration-continue)]{#continue
      .dfn}, and
      [[break](https://infra.spec.whatwg.org/#iteration-break)]{#break
      .dfn}.
    - [[Assert](https://infra.spec.whatwg.org/#assert)]{#assert .dfn}
    - [[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined)]{#implementation-defined
      .dfn}
    - [[willful
      violation](https://infra.spec.whatwg.org/#willful-violation)]{#willful-violation
      .dfn}
    - [[]{#fingerprinting-vector} [[tracking
      vector](https://infra.spec.whatwg.org/#tracking-vector)]{#tracking-vector
      .dfn}]{#fingerprint}
    - [[code
      point](https://infra.spec.whatwg.org/#code-point)]{#code-point
      .dfn} and its synonym
      [[character](https://infra.spec.whatwg.org/#code-point)]{#character
      .dfn}
    - [[surrogate](https://infra.spec.whatwg.org/#surrogate)]{#surrogate
      .dfn}
    - [[scalar
      value](https://infra.spec.whatwg.org/#scalar-value)]{#scalar-value
      .dfn}
    - [[tuple](https://infra.spec.whatwg.org/#tuple)]{#tuple .dfn}
    - [[noncharacter](https://infra.spec.whatwg.org/#noncharacter)]{#noncharacter
      .dfn}
    - [[string](https://infra.spec.whatwg.org/#string)]{#string .dfn},
      [[code unit](https://infra.spec.whatwg.org/#code-unit)]{#code-unit
      .dfn}, [[code unit
      prefix](https://infra.spec.whatwg.org/#code-unit-prefix)]{#code-unit-prefix
      .dfn}, [[code unit less
      than](https://infra.spec.whatwg.org/#code-unit-less-than)]{#code-unit-less-than
      .dfn}, [[starts
      with](https://infra.spec.whatwg.org/#string-starts-with)]{#starts-with
      .dfn}, [[ends
      with](https://infra.spec.whatwg.org/#string-ends-with)]{#ends-with
      .dfn},
      [[length](https://infra.spec.whatwg.org/#string-length)]{#length
      .dfn}, and [[code point
      length](https://infra.spec.whatwg.org/#string-code-point-length)]{#code-point-length
      .dfn}
    - [The string equality operations
      [[is](https://infra.spec.whatwg.org/#string-is)]{#is .dfn} and
      [[identical
      to](https://infra.spec.whatwg.org/#string-is)]{#identical-to
      .dfn}]{#case-sensitive}
    - [[scalar value
      string](https://infra.spec.whatwg.org/#scalar-value-string)]{#scalar-value-string
      .dfn}
    - [[convert](https://infra.spec.whatwg.org/#javascript-string-convert)]{#convert
      .dfn}
    - [[ASCII
      string](https://infra.spec.whatwg.org/#ascii-string)]{#ascii-string
      .dfn}
    - [[ASCII tab or
      newline](https://infra.spec.whatwg.org/#ascii-tab-or-newline)]{#ascii-tab-or-newline
      .dfn}
    - [[ASCII
      whitespace](https://infra.spec.whatwg.org/#ascii-whitespace)]{#space-characters
      .dfn}
    - [[control](https://infra.spec.whatwg.org/#control)]{#control .dfn}
    - [[ASCII
      digit](https://infra.spec.whatwg.org/#ascii-digit)]{#ascii-digits
      .dfn}
    - [[ASCII upper hex
      digit](https://infra.spec.whatwg.org/#ascii-upper-hex-digit)]{#uppercase-ascii-hex-digits
      .dfn}
    - [[ASCII lower hex
      digit](https://infra.spec.whatwg.org/#ascii-lower-hex-digit)]{#lowercase-ascii-hex-digits
      .dfn}
    - [[ASCII hex
      digit](https://infra.spec.whatwg.org/#ascii-hex-digit)]{#ascii-hex-digits
      .dfn}
    - [[ASCII upper
      alpha](https://infra.spec.whatwg.org/#ascii-upper-alpha)]{#uppercase-ascii-letters
      .dfn}
    - [[ASCII lower
      alpha](https://infra.spec.whatwg.org/#ascii-lower-alpha)]{#lowercase-ascii-letters
      .dfn}
    - [[ASCII
      alpha](https://infra.spec.whatwg.org/#ascii-alpha)]{#ascii-letters
      .dfn}
    - [[ASCII
      alphanumeric](https://infra.spec.whatwg.org/#ascii-alphanumeric)]{#alphanumeric-ascii-characters
      .dfn}
    - [[isomorphic
      decode](https://infra.spec.whatwg.org/#isomorphic-decode)]{#isomorphic-decode
      .dfn}
    - [[isomorphic
      encode](https://infra.spec.whatwg.org/#isomorphic-encode)]{#isomorphic-encode
      .dfn}
    - [[ASCII
      lowercase](https://infra.spec.whatwg.org/#ascii-lowercase)]{#converted-to-ascii-lowercase
      .dfn}
    - [[ASCII
      uppercase](https://infra.spec.whatwg.org/#ascii-uppercase)]{#converted-to-ascii-uppercase
      .dfn}
    - [[ASCII
      case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive)]{#ascii-case-insensitive
      .dfn}
    - [[strip
      newlines](https://infra.spec.whatwg.org/#strip-newlines)]{#strip-newlines
      .dfn}
    - [[normalize
      newlines](https://infra.spec.whatwg.org/#normalize-newlines)]{#normalize-newlines
      .dfn}
    - [[strip leading and trailing ASCII
      whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace)]{#strip-leading-and-trailing-ascii-whitespace
      .dfn}
    - [[strip and collapse ASCII
      whitespace](https://infra.spec.whatwg.org/#strip-and-collapse-ascii-whitespace)]{#strip-and-collapse-ascii-whitespace
      .dfn}
    - [[split a string on ASCII
      whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace)]{#split-a-string-on-spaces
      .dfn}
    - [[split a string on
      commas](https://infra.spec.whatwg.org/#split-on-commas)]{#split-a-string-on-commas
      .dfn}
    - [[collect a sequence of code
      points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points)]{#collect-a-sequence-of-code-points
      .dfn} and its associated [[position
      variable](https://infra.spec.whatwg.org/#string-position-variable)]{#position-variable
      .dfn}
    - [[skip ASCII
      whitespace](https://infra.spec.whatwg.org/#skip-ascii-whitespace)]{#skip-ascii-whitespace
      .dfn}
    - The [[ordered
      map](https://infra.spec.whatwg.org/#ordered-map)]{#ordered-map
      .dfn} data structure and the associated definitions for
      [[key](https://infra.spec.whatwg.org/#map-key)]{#map-key .dfn},
      [[value](https://infra.spec.whatwg.org/#map-value)]{#map-value
      .dfn},
      [[empty](https://infra.spec.whatwg.org/#map-is-empty)]{#map-empty
      .dfn},
      [[entry](https://infra.spec.whatwg.org/#map-entry)]{#map-entry
      .dfn},
      [[exists](https://infra.spec.whatwg.org/#map-exists)]{#map-exists
      .dfn}, [[getting the value of an
      entry](https://infra.spec.whatwg.org/#map-get)]{#map-get .dfn},
      [[setting the value of an
      entry](https://infra.spec.whatwg.org/#map-set)]{#map-set .dfn},
      [[removing an
      entry](https://infra.spec.whatwg.org/#map-remove)]{#map-remove
      .dfn},
      [[clear](https://infra.spec.whatwg.org/#map-clear)]{#map-clear
      .dfn}, [[getting the
      keys](https://infra.spec.whatwg.org/#map-getting-the-keys)]{#map-get-the-keys
      .dfn}, [[getting the
      values](https://infra.spec.whatwg.org/#map-getting-the-values)]{#map-get-the-values
      .dfn}, [[sorting in descending
      order](https://infra.spec.whatwg.org/#map-sort-in-descending-order)]{#map-sort-descending
      .dfn}, [[size](https://infra.spec.whatwg.org/#map-size)]{#map-size
      .dfn}, and
      [[iterate](https://infra.spec.whatwg.org/#map-iterate)]{#map-iterate
      .dfn}
    - The [[list](https://infra.spec.whatwg.org/#list)]{#list .dfn} data
      structure and the associated definitions for
      [[append](https://infra.spec.whatwg.org/#list-append)]{#list-append
      .dfn},
      [[extend](https://infra.spec.whatwg.org/#list-extend)]{#list-extend
      .dfn},
      [[prepend](https://infra.spec.whatwg.org/#list-prepend)]{#list-prepend
      .dfn},
      [[replace](https://infra.spec.whatwg.org/#list-replace)]{#list-replace
      .dfn},
      [[remove](https://infra.spec.whatwg.org/#list-remove)]{#list-remove
      .dfn},
      [[empty](https://infra.spec.whatwg.org/#list-empty)]{#list-empty
      .dfn},
      [[contains](https://infra.spec.whatwg.org/#list-contain)]{#list-contains
      .dfn},
      [[size](https://infra.spec.whatwg.org/#list-size)]{#list-size
      .dfn},
      [[indices](https://infra.spec.whatwg.org/#list-get-the-indices)]{#indices
      .dfn}, [[is
      empty](https://infra.spec.whatwg.org/#list-is-empty)]{#list-is-empty
      .dfn},
      [[item](https://infra.spec.whatwg.org/#list-item)]{#list-item
      .dfn},
      [[iterate](https://infra.spec.whatwg.org/#list-iterate)]{#list-iterate
      .dfn}, and
      [[clone](https://infra.spec.whatwg.org/#list-clone)]{#list-clone
      .dfn} [[sort in ascending
      order](https://infra.spec.whatwg.org/#list-sort-in-ascending-order)]{#list-sort
      .dfn} [[sort in descending
      order](https://infra.spec.whatwg.org/#list-sort-in-descending-order)]{#list-sort-descending
      .dfn}
    - The [[stack](https://infra.spec.whatwg.org/#stack)]{#stack .dfn}
      data structure and the associated definitions for
      [[push](https://infra.spec.whatwg.org/#stack-push)]{#stack-push
      .dfn} and
      [[pop](https://infra.spec.whatwg.org/#stack-pop)]{#stack-pop .dfn}
    - The [[queue](https://infra.spec.whatwg.org/#queue)]{#queue .dfn}
      data structure and the associated definitions for
      [[enqueue](https://infra.spec.whatwg.org/#queue-enqueue)]{#enqueue
      .dfn} and
      [[dequeue](https://infra.spec.whatwg.org/#queue-dequeue)]{#dequeue
      .dfn}
    - The [[ordered
      set](https://infra.spec.whatwg.org/#ordered-set)]{#set .dfn} data
      structure and the associated definition for
      [[append](https://infra.spec.whatwg.org/#set-append)]{#set-append
      .dfn} and
      [[union](https://infra.spec.whatwg.org/#set-union)]{#set-union
      .dfn}
    - The [[struct](https://infra.spec.whatwg.org/#struct)]{#struct
      .dfn} specification type and the associated definition for
      [[item](https://infra.spec.whatwg.org/#struct-item)]{#struct-item
      .dfn}
    - The [[byte
      sequence](https://infra.spec.whatwg.org/#byte-sequence)]{#byte-sequence
      .dfn} data structure
    - The [[forgiving-base64
      encode](https://infra.spec.whatwg.org/#forgiving-base64-encode)]{#forgiving-base64-encode
      .dfn} and [[forgiving-base64
      decode](https://infra.spec.whatwg.org/#forgiving-base64-decode)]{#forgiving-base64-decode
      .dfn} algorithms
    - [[exclusive
      range](https://infra.spec.whatwg.org/#the-exclusive-range)]{#exclusive-range
      .dfn}
    - [[parse a JSON string to an Infra
      value](https://infra.spec.whatwg.org/#parse-a-json-string-to-an-infra-value)]{#parse-a-json-string-to-an-infra-value
      .dfn}
    - [[HTML
      namespace](https://infra.spec.whatwg.org/#html-namespace)]{#html-namespace-2
      .dfn}
    - [[MathML
      namespace](https://infra.spec.whatwg.org/#mathml-namespace)]{#mathml-namespace
      .dfn}
    - [[SVG
      namespace](https://infra.spec.whatwg.org/#svg-namespace)]{#svg-namespace
      .dfn}
    - [[XLink
      namespace](https://infra.spec.whatwg.org/#xlink-namespace)]{#xlink-namespace
      .dfn}
    - [[XML
      namespace](https://infra.spec.whatwg.org/#xml-namespace)]{#xml-namespace
      .dfn}
    - [[XMLNS
      namespace](https://infra.spec.whatwg.org/#xmlns-namespace)]{#xmlns-namespace
      .dfn}

Unicode and Encoding

:   The Unicode character set is used to represent textual data, and
    Encoding defines requirements around [character
    encodings](https://encoding.spec.whatwg.org/#encoding){#dependencies:encoding
    x-internal="encoding"}. [\[UNICODE\]](references.html#refsUNICODE)

    This specification [introduces terminology](#encoding-terminology)
    based on the terms defined in those specifications, as described
    earlier.

    The following terms are used as defined in Encoding:
    [\[ENCODING\]](references.html#refsENCODING)

    - [[Getting an
      encoding](https://encoding.spec.whatwg.org/#concept-encoding-get)]{#getting-an-encoding
      .dfn}
    - [[Get an output
      encoding](https://encoding.spec.whatwg.org/#get-an-output-encoding)]{#get-an-output-encoding
      .dfn}
    - The generic
      [[decode](https://encoding.spec.whatwg.org/#decode)]{#decode .dfn}
      algorithm which takes a byte stream and an encoding and returns a
      character stream
    - The [[UTF-8
      decode](https://encoding.spec.whatwg.org/#utf-8-decode)]{#utf-8-decode
      .dfn} algorithm which takes a byte stream and returns a character
      stream, additionally stripping one leading UTF-8 Byte Order Mark
      (BOM), if any
    - The [[UTF-8 decode without
      BOM](https://encoding.spec.whatwg.org/#utf-8-decode-without-bom)]{#utf-8-decode-without-bom
      .dfn} algorithm which is identical to [UTF-8
      decode](https://encoding.spec.whatwg.org/#utf-8-decode){#dependencies:utf-8-decode
      x-internal="utf-8-decode"} except that it does not strip one
      leading UTF-8 Byte Order Mark (BOM)
    - The [[encode](https://encoding.spec.whatwg.org/#encode)]{#encode
      .dfn} algorithm which takes a character stream and an encoding and
      returns a byte stream
    - The [[UTF-8
      encode](https://encoding.spec.whatwg.org/#utf-8-encode)]{#utf-8-encode
      .dfn} algorithm which takes a character stream and returns a byte
      stream
    - The [[BOM
      sniff](https://encoding.spec.whatwg.org/#bom-sniff)]{#bom-sniff
      .dfn} algorithm which takes a byte stream and returns an encoding
      or null.

XML and related specifications

:   Implementations that support [the XML
    syntax](xhtml.html#the-xhtml-syntax){#dependencies:the-xhtml-syntax}
    for HTML must support some version of XML, as well as its
    corresponding namespaces specification, because that syntax uses an
    XML serialization with namespaces.
    [\[XML\]](references.html#refsXML)
    [\[XMLNS\]](references.html#refsXMLNS)

    Data mining tools and other user agents that perform operations on
    content without running scripts, evaluating CSS or XPath
    expressions, or otherwise exposing the resulting DOM to arbitrary
    content, may \"support namespaces\" by just asserting that their DOM
    node analogues are in certain namespaces, without actually exposing
    the namespace strings.

    In [the HTML syntax](syntax.html#syntax){#dependencies:syntax},
    namespace prefixes and namespace declarations do not have the same
    effect as in XML. For instance, the colon has no special meaning in
    HTML element names.

    ------------------------------------------------------------------------

    The attribute with the name
    [[`space`](https://www.w3.org/TR/xml/#sec-white-space)]{#attr-xml-space
    .dfn} in the [XML
    namespace](https://infra.spec.whatwg.org/#xml-namespace){#dependencies:xml-namespace
    x-internal="xml-namespace"} is defined by Extensible Markup Language
    (XML). [\[XML\]](references.html#refsXML)

    The [[`Name`](https://www.w3.org/TR/xml/#NT-Name)]{#xml-name .dfn}
    production is defined in XML. [\[XML\]](references.html#refsXML)

    This specification also references the
    [[`<?xml-stylesheet?>`](https://www.w3.org/TR/xml-stylesheet/#the-xml-stylesheet-processing-instruction)]{#xml-stylesheet
    .dfn} processing instruction, defined in Associating Style Sheets
    with XML documents. [\[XMLSSPI\]](references.html#refsXMLSSPI)

    This specification also non-normatively mentions the
    [`XSLTProcessor`]{#xsltprocessor .dfn} interface and its
    [`transformToFragment()`]{#dom-xsltprocessor-transformtofragment
    .dfn} and
    [`transformToDocument()`]{#dom-xsltprocessor-transformtodocument
    .dfn} methods. [\[XSLTP\]](references.html#refsXSLTP)

URLs

:   The following terms are defined in URL:
    [\[URL\]](references.html#refsURL)

    - [[host](https://url.spec.whatwg.org/#concept-host)]{#concept-host
      .dfn}
    - [[public
      suffix](https://url.spec.whatwg.org/#host-public-suffix)]{#public-suffix
      .dfn}
    - [[domain](https://url.spec.whatwg.org/#concept-domain)]{#concept-domain
      .dfn}
    - [[IP
      address](https://url.spec.whatwg.org/#ip-address)]{#ip-address
      .dfn}
    - [[URL](https://url.spec.whatwg.org/#concept-url)]{#url .dfn}
    - [[Origin](https://url.spec.whatwg.org/#concept-url-origin)]{#concept-url-origin
      .dfn} of URLs
    - [[Absolute
      URL](https://url.spec.whatwg.org/#syntax-url-absolute)]{#absolute-url
      .dfn}
    - [[Relative
      URL](https://url.spec.whatwg.org/#syntax-url-relative)]{#relative-url
      .dfn}
    - [[registrable
      domain](https://url.spec.whatwg.org/#host-registrable-domain)]{#registrable-domain
      .dfn}
    - The [[URL
      parser](https://url.spec.whatwg.org/#concept-url-parser)]{#url-parser
      .dfn}
    - The [[basic URL
      parser](https://url.spec.whatwg.org/#concept-basic-url-parser)]{#basic-url-parser
      .dfn} and its
      [[`url`{.variable}](https://url.spec.whatwg.org/#basic-url-parser-url)]{#basic-url-parser-url
      .dfn} and
      [[`state override`{.variable}](https://url.spec.whatwg.org/#basic-url-parser-state-override)]{#basic-url-parser-state-override
      .dfn} arguments, as well as these parser states:
      - [[scheme start
        state](https://url.spec.whatwg.org/#scheme-start-state)]{#scheme-start-state
        .dfn}
      - [[host
        state](https://url.spec.whatwg.org/#host-state)]{#host-state
        .dfn}
      - [[hostname
        state](https://url.spec.whatwg.org/#hostname-state)]{#hostname-state
        .dfn}
      - [[port
        state](https://url.spec.whatwg.org/#port-state)]{#port-state
        .dfn}
      - [[path start
        state](https://url.spec.whatwg.org/#path-start-state)]{#path-start-state
        .dfn}
      - [[query
        state](https://url.spec.whatwg.org/#query-state)]{#query-state
        .dfn}
      - [[fragment
        state](https://url.spec.whatwg.org/#fragment-state)]{#fragment-state
        .dfn}
    - [[URL
      record](https://url.spec.whatwg.org/#concept-url)]{#url-record
      .dfn}, as well as its individual components:
      - [[scheme](https://url.spec.whatwg.org/#concept-url-scheme)]{#concept-url-scheme
        .dfn}
      - [[username](https://url.spec.whatwg.org/#concept-url-username)]{#concept-url-username
        .dfn}
      - [[password](https://url.spec.whatwg.org/#concept-url-password)]{#concept-url-password
        .dfn}
      - [[host](https://url.spec.whatwg.org/#concept-url-host)]{#concept-url-host
        .dfn}
      - [[port](https://url.spec.whatwg.org/#concept-url-port)]{#concept-url-port
        .dfn}
      - [[path](https://url.spec.whatwg.org/#concept-url-path)]{#concept-url-path
        .dfn}
      - [[query](https://url.spec.whatwg.org/#concept-url-query)]{#concept-url-query
        .dfn}
      - [[fragment](https://url.spec.whatwg.org/#concept-url-fragment)]{#concept-url-fragment
        .dfn}
      - [[blob URL
        entry](https://url.spec.whatwg.org/#concept-url-blob-entry)]{#concept-url-blob-entry
        .dfn}
    - [[valid URL
      string](https://url.spec.whatwg.org/#valid-url-string)]{#valid-url-string
      .dfn}
    - The [[cannot have a
      username/password/port](https://url.spec.whatwg.org/#cannot-have-a-username-password-port)]{#cannot-have-a-username/password/port
      .dfn} concept
    - The [[opaque
      path](https://url.spec.whatwg.org/#url-opaque-path)]{#opaque-path
      .dfn} concept
    - [[URL
      serializer](https://url.spec.whatwg.org/#concept-url-serializer)]{#concept-url-serializer
      .dfn} and its [[exclude
      fragment](https://url.spec.whatwg.org/#url-serializer-exclude-fragment)]{#url-serializer-exclude-fragment
      .dfn} argument
    - [[URL path
      serializer](https://url.spec.whatwg.org/#url-path-serializer)]{#url-path-serializer
      .dfn}
    - The [[host
      parser](https://url.spec.whatwg.org/#concept-host-parser)]{#host-parser
      .dfn}
    - The [[host
      serializer](https://url.spec.whatwg.org/#concept-host-serializer)]{#host-serializer
      .dfn}
    - [[Host
      equals](https://url.spec.whatwg.org/#concept-host-equals)]{#host-equals
      .dfn}
    - [[URL
      equals](https://url.spec.whatwg.org/#concept-url-equals)]{#concept-url-equals
      .dfn} and its [[exclude
      fragments](https://url.spec.whatwg.org/#url-equals-exclude-fragments)]{#url-equals-exclude-fragments
      .dfn} argument
    - [[serialize an
      integer](https://url.spec.whatwg.org/#serialize-an-integer)]{#serialize-an-integer
      .dfn}
    - [[Default encode
      set](https://url.spec.whatwg.org/#default-encode-set)]{#default-encode-set
      .dfn}
    - [[component percent-encode
      set](https://url.spec.whatwg.org/#component-percent-encode-set)]{#component-percent-encode-set
      .dfn}
    - [[UTF-8
      percent-encode](https://url.spec.whatwg.org/#string-utf-8-percent-encode)]{#utf-8-percent-encode
      .dfn}
    - [[percent-decode](https://url.spec.whatwg.org/#string-percent-decode)]{#percent-decode
      .dfn}
    - [[set the
      username](https://url.spec.whatwg.org/#set-the-username)]{#set-the-username
      .dfn}
    - [[set the
      password](https://url.spec.whatwg.org/#set-the-password)]{#set-the-password
      .dfn}
    - The
      [[`application/x-www-form-urlencoded`](https://url.spec.whatwg.org/#concept-urlencoded)]{#application/x-www-form-urlencoded
      .dfn} format
    - The [[`application/x-www-form-urlencoded`
      serializer](https://url.spec.whatwg.org/#concept-urlencoded-serializer)]{#application/x-www-form-urlencoded-serializer
      .dfn}
    - [[is
      special](https://url.spec.whatwg.org/#is-special)]{#is-special
      .dfn}

    A number of schemes and protocols are referenced by this
    specification also:

    - The
      [[`about:`](https://www.rfc-editor.org/rfc/rfc6694#section-2)]{#about-protocol
      .dfn} scheme [\[ABOUT\]](references.html#refsABOUT)
    - The
      [[`blob:`](https://w3c.github.io/FileAPI/#DefinitionOfScheme)]{#blob-protocol
      .dfn} scheme [\[FILEAPI\]](references.html#refsFILEAPI)
    - The
      [[`data:`](https://www.rfc-editor.org/rfc/rfc2397#section-2)]{#data-protocol
      .dfn} scheme [\[RFC2397\]](references.html#refsRFC2397)
    - The
      [[`http:`](https://httpwg.org/specs/rfc7230.html#http.uri)]{#http-protocol
      .dfn} scheme [\[HTTP\]](references.html#refsHTTP)
    - The
      [[`https:`](https://httpwg.org/specs/rfc7230.html#https.uri)]{#https-protocol
      .dfn} scheme [\[HTTP\]](references.html#refsHTTP)
    - The
      [[`mailto:`](https://www.rfc-editor.org/rfc/rfc6068#section-2)]{#mailto-protocol
      .dfn} scheme [\[MAILTO\]](references.html#refsMAILTO)
    - The
      [[`sms:`](https://www.rfc-editor.org/rfc/rfc5724#section-2)]{#sms-protocol
      .dfn} scheme [\[SMS\]](references.html#refsSMS)
    - The
      [[`urn:`](https://www.rfc-editor.org/rfc/rfc2141#section-2)]{#urn-protocol
      .dfn} scheme [\[URN\]](references.html#refsURN)

    [[Media fragment
    syntax](https://www.w3.org/TR/media-frags/#media-fragment-syntax)]{#media-fragment-syntax
    .dfn} is defined in Media Fragments URI.
    [\[MEDIAFRAG\]](references.html#refsMEDIAFRAG)

HTTP and related specifications

:   The following terms are defined in the HTTP specifications:
    [\[HTTP\]](references.html#refsHTTP)

    - \`[[`Accept`](https://httpwg.org/specs/rfc7231.html#header.accept)]{#http-accept
      .dfn}\` header
    - \`[[`Accept-Language`](https://httpwg.org/specs/rfc7231.html#header.accept-language)]{#http-accept-language
      .dfn}\` header
    - \`[[`Cache-Control`](https://httpwg.org/specs/rfc7234.html#header.cache-control)]{#http-cache-control
      .dfn}\` header
    - \`[[`Content-Disposition`](https://httpwg.org/specs/rfc6266.html)]{#http-content-disposition
      .dfn}\` header
    - \`[[`Content-Language`](https://httpwg.org/specs/rfc7231.html#header.content-language)]{#http-content-language
      .dfn}\` header
    - \`[[`Content-Range`](https://httpwg.org/specs/rfc7233.html#header.content-range)]{#http-content-range
      .dfn}\` header
    - \`[[`Last-Modified`](https://httpwg.org/specs/rfc7232.html#header.last-modified)]{#http-last-modified
      .dfn}\` header
    - \`[[`Range`](https://httpwg.org/specs/rfc7233.html#header.range)]{#http-range
      .dfn}\` header
    - \`[[`Referer`](https://httpwg.org/specs/rfc7231.html#header.referer)]{#http-referer
      .dfn}\` header

    The following terms are defined in HTTP State Management Mechanism:
    [\[COOKIES\]](references.html#refsCOOKIES)

    - [[cookie-string](https://httpwg.org/specs/rfc6265.html#sane-cookie-syntax)]{#cookie-string
      .dfn}
    - [[receives a
      set-cookie-string](https://httpwg.org/specs/rfc6265.html#storage-model)]{#receives-a-set-cookie-string
      .dfn}
    - \`[[`Cookie`](https://httpwg.org/specs/rfc6265.html#cookie)]{#http-cookie
      .dfn}\` header

    The following term is defined in Web Linking:
    [\[WEBLINK\]](references.html#refsWEBLINK)

    - \`[[`Link`](https://httpwg.org/specs/rfc8288.html#header)]{#http-link
      .dfn}\` header
    - [[Parsing a \``Link`\` field
      value](https://httpwg.org/specs/rfc8288.html#parse-fv)]{#parsing-a-link-field-value
      .dfn}

    The following terms are defined in Structured Field Values for HTTP:
    [\[STRUCTURED-FIELDS\]](references.html#refsSTRUCTURED-FIELDS)

    - [[structured
      header](https://httpwg.org/specs/rfc8941.html)]{#http-structured-header
      .dfn}
    - [[boolean](https://httpwg.org/specs/rfc8941.html#boolean)]{#http-structured-header-boolean
      .dfn}
    - [[token](https://httpwg.org/specs/rfc8941.html#token)]{#http-structured-header-token
      .dfn}
    - [[parameters](https://httpwg.org/specs/rfc8941.html#param)]{#http-structured-header-parameters
      .dfn}

    The following terms are defined in MIME Sniffing:
    [\[MIMESNIFF\]](references.html#refsMIMESNIFF)

    - [[MIME
      type](https://mimesniff.spec.whatwg.org/#mime-type)]{#mime-type
      .dfn}
    - [[MIME type
      essence](https://mimesniff.spec.whatwg.org/#mime-type-essence)]{#mime-type-essence
      .dfn}
    - [[valid MIME type
      string](https://mimesniff.spec.whatwg.org/#valid-mime-type)]{#valid-mime-type-string
      .dfn}
    - [[valid MIME type string with no
      parameters](https://mimesniff.spec.whatwg.org/#valid-mime-type-with-no-parameters)]{#valid-mime-type-string-with-no-parameters
      .dfn}
    - [[HTML MIME
      type](https://mimesniff.spec.whatwg.org/#html-mime-type)]{#html-mime-type
      .dfn}
    - [[JavaScript MIME
      type](https://mimesniff.spec.whatwg.org/#javascript-mime-type)]{#javascript-mime-type
      .dfn} and [[JavaScript MIME type essence
      match](https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match)]{#javascript-mime-type-essence-match
      .dfn}
    - [[JSON MIME
      type](https://mimesniff.spec.whatwg.org/#json-mime-type)]{#json-mime-type
      .dfn}
    - [[XML MIME
      type](https://mimesniff.spec.whatwg.org/#xml-mime-type)]{#xml-mime-type
      .dfn}
    - [[image MIME
      type](https://mimesniff.spec.whatwg.org/#image-mime-type)]{#image-mime-type
      .dfn}
    - [[audio or video MIME
      type](https://mimesniff.spec.whatwg.org/#audio-or-video-mime-type)]{#audio-or-video-mime-type
      .dfn}
    - [[font MIME
      type](https://mimesniff.spec.whatwg.org/#font-mime-type)]{#font-mime-type
      .dfn}
    - [[parse a MIME
      type](https://mimesniff.spec.whatwg.org/#parse-a-mime-type)]{#parse-a-mime-type
      .dfn}
    - [[is MIME type supported by the user
      agent?](https://mimesniff.spec.whatwg.org/#supported-by-the-user-agent)]{#is-mime-type-supported-by-the-user-agent
      .dfn}

Fetch

:   The following terms are defined in Fetch:
    [\[FETCH\]](references.html#refsFETCH)

    - [[ABNF](https://fetch.spec.whatwg.org/#abnf)]{#header-abnf .dfn}
    - [`about:blank`]{#about:blank .dfn}
    - An [[HTTP(S)
      scheme](https://fetch.spec.whatwg.org/#http-scheme)]{#http(s)-scheme
      .dfn}
    - A URL which [[is
      local](https://fetch.spec.whatwg.org/#is-local)]{#is-local .dfn}
    - A [[local
      scheme](https://fetch.spec.whatwg.org/#local-scheme)]{#local-scheme
      .dfn}
    - A [[fetch
      scheme](https://fetch.spec.whatwg.org/#fetch-scheme)]{#fetch-scheme
      .dfn}
    - [[CORS
      protocol](https://fetch.spec.whatwg.org/#http-cors-protocol)]{#cors-protocol
      .dfn}
    - [[default \``User-Agent`\`
      value](https://fetch.spec.whatwg.org/#default-user-agent-value)]{#default-user-agent-value
      .dfn}
    - [[extract a MIME
      type](https://fetch.spec.whatwg.org/#concept-header-extract-mime-type)]{#extract-a-mime-type
      .dfn}
    - [[legacy extract an
      encoding](https://fetch.spec.whatwg.org/#legacy-extract-an-encoding)]{#legacy-extract-an-encoding
      .dfn}
    - [[fetch](https://fetch.spec.whatwg.org/#concept-fetch)]{#concept-fetch
      .dfn}
    - [[fetch
      controller](https://fetch.spec.whatwg.org/#fetch-controller)]{#fetch-controller
      .dfn}
    - [[process the next manual
      redirect](https://fetch.spec.whatwg.org/#fetch-controller-process-the-next-manual-redirect)]{#process-the-next-manual-redirect
      .dfn}
    - [[ok status](https://fetch.spec.whatwg.org/#ok-status)]{#ok-status
      .dfn}
    - [[navigation
      request](https://fetch.spec.whatwg.org/#navigation-request)]{#navigation-request
      .dfn}
    - [[network
      error](https://fetch.spec.whatwg.org/#concept-network-error)]{#network-error
      .dfn}
    - [[aborted network
      error](https://fetch.spec.whatwg.org/#concept-aborted-network-error)]{#aborted-network-error
      .dfn}
    - \`[[`Origin`](https://fetch.spec.whatwg.org/#http-origin)]{#http-origin
      .dfn}\` header
    - \`[[`Cross-Origin-Resource-Policy`](https://fetch.spec.whatwg.org/#http-cross-origin-resource-policy)]{#cross-origin-resource-policy
      .dfn}\` header
    - [[getting a structured field
      value](https://fetch.spec.whatwg.org/#concept-header-list-get-structured-header)]{#getting-a-structured-field-value
      .dfn}
    - [[header
      list](https://fetch.spec.whatwg.org/#concept-header-list)]{#concept-header-list
      .dfn}
    - [[set](https://fetch.spec.whatwg.org/#concept-header-list-set)]{#concept-header-list-set
      .dfn}
    - [[get, decode, and
      split](https://fetch.spec.whatwg.org/#concept-header-list-get-decode-split)]{#concept-header-list-get-decode-split
      .dfn}
    - [[abort](https://fetch.spec.whatwg.org/#fetch-controller-abort)]{#fetch-controller-abort
      .dfn}
    - [[cross-origin resource policy
      check](https://fetch.spec.whatwg.org/#cross-origin-resource-policy-check)]{#cross-origin-resource-policy-check
      .dfn}
    - the
      [[`RequestCredentials`](https://fetch.spec.whatwg.org/#requestcredentials)]{#requestcredentials
      .dfn} enumeration
    - the
      [[`RequestDestination`](https://fetch.spec.whatwg.org/#requestdestination)]{#requestdestination
      .dfn} enumeration
    - the
      [[`fetch()`](https://fetch.spec.whatwg.org/#dom-global-fetch)]{#fetch()
      .dfn} method
    - [[report
      timing](https://fetch.spec.whatwg.org/#finalize-and-report-timing)]{#report-timing
      .dfn}
    - [[serialize a response URL for
      reporting](https://fetch.spec.whatwg.org/#serialize-a-response-url-for-reporting)]{#serialize-a-response-url-for-reporting
      .dfn}
    - [[safely extracting a
      body](https://fetch.spec.whatwg.org/#bodyinit-safely-extract)]{#body-safely-extract
      .dfn}
    - [[incrementally reading a
      body](https://fetch.spec.whatwg.org/#body-incrementally-read)]{#body-incrementally-read
      .dfn}
    - [[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body)]{#processresponseconsumebody
      .dfn}
    - [[processResponseEndOfBody](https://fetch.spec.whatwg.org/#fetch-processresponseendofbody)]{#processresponseendofbody
      .dfn}
    - [[processResponse](https://fetch.spec.whatwg.org/#process-response)]{#processresponse
      .dfn}
    - [[useParallelQueue](https://fetch.spec.whatwg.org/#fetch-useparallelqueue)]{#useparallelqueue
      .dfn}
    - [[processEarlyHintsResponse](https://fetch.spec.whatwg.org/#fetch-processearlyhintsresponse)]{#processearlyhintsresponse
      .dfn}
    - [[connection
      pool](https://fetch.spec.whatwg.org/#concept-connection-pool)]{#connection-pool
      .dfn}
    - [[obtain a
      connection](https://fetch.spec.whatwg.org/#concept-connection-obtain)]{#obtain-a-connection
      .dfn}
    - [[determine the network partition
      key](https://fetch.spec.whatwg.org/#determine-the-network-partition-key)]{#determine-the-network-partition-key
      .dfn}
    - [[extract full timing
      info](https://fetch.spec.whatwg.org/#extract-full-timing-info)]{#extract-full-timing-info
      .dfn}
    - [[as a
      body](https://fetch.spec.whatwg.org/#byte-sequence-as-a-body)]{#as-a-body
      .dfn}
    - [[response body
      info](https://fetch.spec.whatwg.org/#response-body-info)]{#response-body-info
      .dfn}
    - [[resolve an
      origin](https://fetch.spec.whatwg.org/#resolve-an-origin)]{#resolve-an-origin
      .dfn}
    - [[response](https://fetch.spec.whatwg.org/#concept-response)]{#concept-response
      .dfn} and its associated:
      - [[type](https://fetch.spec.whatwg.org/#concept-response-type)]{#concept-response-type
        .dfn}
      - [[URL](https://fetch.spec.whatwg.org/#concept-response-url)]{#concept-response-url
        .dfn}
      - [[URL
        list](https://fetch.spec.whatwg.org/#concept-response-url-list)]{#concept-response-url-list
        .dfn}
      - [[status](https://fetch.spec.whatwg.org/#concept-response-status)]{#concept-response-status
        .dfn}
      - [[header
        list](https://fetch.spec.whatwg.org/#concept-response-header-list)]{#concept-response-header-list
        .dfn}
      - [[body](https://fetch.spec.whatwg.org/#concept-response-body)]{#concept-response-body
        .dfn}
      - [[body
        info](https://fetch.spec.whatwg.org/#concept-response-body-info)]{#concept-response-body-info
        .dfn}
      - [[internal
        response](https://fetch.spec.whatwg.org/#concept-internal-response)]{#concept-internal-response
        .dfn}
      - [[location
        URL](https://fetch.spec.whatwg.org/#concept-response-location-url)]{#concept-response-location-url
        .dfn}
      - [[timing
        info](https://fetch.spec.whatwg.org/#concept-response-timing-info)]{#concept-response-timing-info
        .dfn}
      - [[service worker timing
        info](https://fetch.spec.whatwg.org/#response-service-worker-timing-info)]{#concept-response-service-worker-timing-info
        .dfn}
      - [[has-cross-origin-redirects](https://fetch.spec.whatwg.org/#response-has-cross-origin-redirects)]{#concept-response-has-cross-origin-redirects
        .dfn}
      - [[timing allow
        passed](https://fetch.spec.whatwg.org/#concept-response-timing-allow-passed)]{#concept-response-timing-allow-passed
        .dfn}
      - [[extract content-range
        values](https://wicg.github.io/background-fetch/#extract-content-range-values)]{#extract-content-range-values
        .dfn}
    - [[request](https://fetch.spec.whatwg.org/#concept-request)]{#concept-request
      .dfn} and its associated:
      - [[URL](https://fetch.spec.whatwg.org/#concept-request-url)]{#concept-request-url
        .dfn}
      - [[method](https://fetch.spec.whatwg.org/#concept-request-method)]{#concept-request-method
        .dfn}
      - [[header
        list](https://fetch.spec.whatwg.org/#concept-request-header-list)]{#concept-request-header-list
        .dfn}
      - [[body](https://fetch.spec.whatwg.org/#concept-request-body)]{#concept-request-body
        .dfn}
      - [[client](https://fetch.spec.whatwg.org/#concept-request-client)]{#concept-request-client
        .dfn}
      - [[URL
        list](https://fetch.spec.whatwg.org/#concept-request-url-list)]{#concept-request-url-list
        .dfn}
      - [[current
        URL](https://fetch.spec.whatwg.org/#concept-request-current-url)]{#concept-request-current-url
        .dfn}
      - [[reserved
        client](https://fetch.spec.whatwg.org/#concept-request-reserved-client)]{#concept-request-reserved-client
        .dfn}
      - [[replaces client
        id](https://fetch.spec.whatwg.org/#concept-request-replaces-client-id)]{#concept-request-replaces-client-id
        .dfn}
      - [[initiator](https://fetch.spec.whatwg.org/#concept-request-initiator)]{#concept-request-initiator
        .dfn}
      - [[destination](https://fetch.spec.whatwg.org/#concept-request-destination)]{#concept-request-destination
        .dfn}
      - [[potential
        destination](https://fetch.spec.whatwg.org/#concept-potential-destination)]{#concept-potential-destination
        .dfn}
      - [[translating](https://fetch.spec.whatwg.org/#concept-potential-destination-translate)]{#concept-potential-destination-translate
        .dfn} a [potential
        destination](https://fetch.spec.whatwg.org/#concept-potential-destination){#dependencies:concept-potential-destination
        x-internal="concept-potential-destination"}
      - [[script-like](https://fetch.spec.whatwg.org/#request-destination-script-like)]{#concept-script-like-destination
        .dfn}
        [destinations](https://fetch.spec.whatwg.org/#concept-request-destination){#dependencies:concept-request-destination
        x-internal="concept-request-destination"}
      - [[priority](https://fetch.spec.whatwg.org/#request-priority)]{#concept-request-priority
        .dfn}
      - [[origin](https://fetch.spec.whatwg.org/#concept-request-origin)]{#concept-request-origin
        .dfn}
      - [[referrer](https://fetch.spec.whatwg.org/#concept-request-referrer)]{#concept-request-referrer
        .dfn}
      - [[synchronous
        flag](https://fetch.spec.whatwg.org/#synchronous-flag)]{#synchronous-flag
        .dfn}
      - [[mode](https://fetch.spec.whatwg.org/#concept-request-mode)]{#concept-request-mode
        .dfn}
      - [[credentials
        mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode)]{#concept-request-credentials-mode
        .dfn}
      - [[use-URL-credentials
        flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag)]{#use-url-credentials-flag
        .dfn}
      - [[unsafe-request
        flag](https://fetch.spec.whatwg.org/#unsafe-request-flag)]{#unsafe-request-flag
        .dfn}
      - [[cache
        mode](https://fetch.spec.whatwg.org/#concept-request-cache-mode)]{#concept-request-cache-mode
        .dfn}
      - [[redirect
        count](https://fetch.spec.whatwg.org/#concept-request-redirect-count)]{#concept-request-redirect-count
        .dfn}
      - [[redirect
        mode](https://fetch.spec.whatwg.org/#concept-request-redirect-mode)]{#concept-request-redirect-mode
        .dfn}
      - [[policy
        container](https://fetch.spec.whatwg.org/#concept-request-policy-container)]{#concept-request-policy-container
        .dfn}
      - [[referrer
        policy](https://fetch.spec.whatwg.org/#concept-request-referrer-policy)]{#concept-request-referrer-policy
        .dfn}
      - [[cryptographic nonce
        metadata](https://fetch.spec.whatwg.org/#concept-request-nonce-metadata)]{#concept-request-nonce-metadata
        .dfn}
      - [[integrity
        metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata)]{#concept-request-integrity-metadata
        .dfn}
      - [[parser
        metadata](https://fetch.spec.whatwg.org/#concept-request-parser-metadata)]{#concept-request-parser-metadata
        .dfn}
      - [[reload-navigation
        flag](https://fetch.spec.whatwg.org/#concept-request-reload-navigation-flag)]{#concept-request-reload-navigation-flag
        .dfn}
      - [[history-navigation
        flag](https://fetch.spec.whatwg.org/#concept-request-history-navigation-flag)]{#concept-request-history-navigation-flag
        .dfn}
      - [[user-activation](https://fetch.spec.whatwg.org/#request-user-activation)]{#concept-request-user-activation
        .dfn}
      - [[render-blocking](https://fetch.spec.whatwg.org/#request-render-blocking)]{#concept-request-render-blocking
        .dfn}
      - [[initiator
        type](https://fetch.spec.whatwg.org/#request-initiator-type)]{#concept-request-initiator-type
        .dfn}
      - [[service-workers
        mode](https://fetch.spec.whatwg.org/#request-service-workers-mode)]{#concept-request-service-workers-mode
        .dfn}
      - [[traversable for user
        prompts](https://fetch.spec.whatwg.org/#concept-request-window)]{#concept-request-traversable-for-user-prompts
        .dfn}
      - [[add a range
        header](https://fetch.spec.whatwg.org/#concept-request-add-range-header)]{#add-a-range-header
        .dfn}
    - [[fetch timing
      info](https://fetch.spec.whatwg.org/#fetch-timing-info)]{#fetch-timing-info
      .dfn} and its associated:
      - [[start
        time](https://fetch.spec.whatwg.org/#fetch-timing-info-start-time)]{#fetch-timing-info-start-time
        .dfn}
      - [[end
        time](https://fetch.spec.whatwg.org/#fetch-timing-info-end-time)]{#fetch-timing-info-end-time
        .dfn}

    The following terms are defined in Referrer Policy:
    [\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

    - [[referrer
      policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy)]{#referrer-policy
      .dfn}
    - The
      \`[[`Referrer-Policy`](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-header-dfn)]{#http-referrer-policy
      .dfn}\` HTTP header
    - The [[parse a referrer policy from a \``Referrer-Policy`\`
      header](https://w3c.github.io/webappsec-referrer-policy/#parse-referrer-policy-from-header)]{#parse-referrer-policy-header
      .dfn} algorithm
    - The
      \"[[`no-referrer`](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-no-referrer)]{#referrer-policy-no-referrer
      .dfn}\",
      \"[[`no-referrer-when-downgrade`](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-no-referrer-when-downgrade)]{#referrer-policy-no-referrer-when-downgrade
      .dfn}\",
      \"[[`origin-when-cross-origin`](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-origin-when-cross-origin)]{#referrer-policy-origin-when-cross-origin
      .dfn}\", and
      \"[[`unsafe-url`](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy-unsafe-url)]{#referrer-policy-unsafe-url
      .dfn}\" referrer policies
    - The [[default referrer
      policy](https://w3c.github.io/webappsec-referrer-policy/#default-referrer-policy)]{#default-referrer-policy
      .dfn}

    The following terms are defined in Mixed Content:
    [\[MIX\]](references.html#refsMIX)

    - [[*a priori* authenticated
      URL](https://w3c.github.io/webappsec-mixed-content/#a-priori-authenticated-url)]{#a-priori-authenticated-url
      .dfn}

    The following terms are defined in Subresource Integrity:
    [\[SRI\]](references.html#refsSRI)

    - [[parse integrity
      metadata](https://w3c.github.io/webappsec-subresource-integrity/#parse-metadata)]{#parse-integrity-metadata
      .dfn}
    - [[the requirements of the integrity
      attribute](https://w3c.github.io/webappsec-subresource-integrity/#the-integrity-attribute)]{#the-requirements-of-the-integrity-attribute
      .dfn}
    - [[get the strongest metadata from
      set](https://w3c.github.io/webappsec-subresource-integrity/#get-the-strongest-metadata)]{#get-the-strongest-metadata-from-set
      .dfn}

Paint Timing

:   The following terms are defined in Paint Timing:
    [\[PAINTTIMING\]](references.html#refsPAINTTIMING)

    - [[mark paint
      timing](https://w3c.github.io/paint-timing/#mark-paint-timing)]{#mark-paint-timing
      .dfn}

Navigation Timing

:   The following terms are defined in Navigation Timing:
    [\[NAVIGATIONTIMING\]](references.html#refsNAVIGATIONTIMING)

    - [[create the navigation timing
      entry](https://w3c.github.io/navigation-timing/#dfn-create-the-navigation-timing-entry)]{#create-the-navigation-timing-entry
      .dfn}
    - [[queue the navigation timing
      entry](https://w3c.github.io/navigation-timing/#dfn-queue-the-navigation-timing-entry)]{#queue-the-navigation-timing-entry
      .dfn}
    - [[`NavigationTimingType`](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype)]{#navigationtimingtype
      .dfn} and its
      \"[[`navigate`](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-navigate)]{#dom-navigationtimingtype-navigate
      .dfn}\",
      \"[[`reload`](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-reload)]{#dom-navigationtimingtype-reload
      .dfn}\", and
      \"[[`back_forward`](https://w3c.github.io/navigation-timing/#dom-navigationtimingtype-back_forward)]{#dom-navigationtimingtype-back_forward
      .dfn}\" values.

Resource Timing

:   The following terms are defined in Resource Timing:
    [\[RESOURCETIMING\]](references.html#refsRESOURCETIMING)

    - [[Mark resource
      timing](https://w3c.github.io/resource-timing/#dfn-mark-resource-timing)]{#mark-resource-timing
      .dfn}

Performance Timeline

:   The following terms are defined in Performance Timeline:
    [\[PERFORMANCETIMELINE\]](references.html#refsPERFORMANCETIMELINE)

    - [[`PerformanceEntry`](https://w3c.github.io/performance-timeline/#dom-performanceentry)]{#performanceentry
      .dfn} and its
      [[`name`](https://w3c.github.io/performance-timeline/#dom-performanceentry-name)]{#performanceentry-name
      .dfn},
      [[`entryType`](https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype)]{#performanceentry-entrytype
      .dfn},
      [[`startTime`](https://w3c.github.io/performance-timeline/#dom-performanceentry-starttime)]{#performanceentry-starttime
      .dfn}, and
      [[`duration`](https://w3c.github.io/performance-timeline/#dom-performanceentry-duration)]{#performanceentry-duration
      .dfn} attributes.
    - [[Queue a performance
      entry](https://w3c.github.io/performance-timeline/#queue-a-performanceentry)]{#queue-a-performance-entry
      .dfn}

Long Animation Frames

:   The following terms are defined in Long Animation Frames:
    [\[LONGANIMATIONFRAMES\]](references.html#refsLONGANIMATIONFRAMES)

    - [[record task start
      time](https://w3c.github.io/long-animation-frames/#record-task-start-time)]{#record-task-start-time
      .dfn}
    - [[record task end
      time](https://w3c.github.io/long-animation-frames/#record-task-end-time)]{#record-task-end-time
      .dfn}
    - [[record rendering
      time](https://w3c.github.io/long-animation-frames/#record-rendering-time)]{#record-rendering-time
      .dfn}
    - [[record classic script creation
      time](https://w3c.github.io/long-animation-frames/#record-classic-script-creation-time)]{#record-classic-script-creation-time
      .dfn}
    - [[record classic script execution start
      time](https://w3c.github.io/long-animation-frames/#record-classic-script-execution-start-time)]{#record-classic-script-execution-start-time
      .dfn}
    - [[record module script execution start
      time](https://w3c.github.io/long-animation-frames/#record-module-script-execution-start-time)]{#record-module-script-execution-start-time
      .dfn}
    - [[Record pause
      duration](https://w3c.github.io/long-animation-frames/#record-pause-duration)]{#record-pause-duration
      .dfn}
    - [[record timing info for timer
      handler](https://w3c.github.io/long-animation-frames/#record-timing-info-for-timer-handler)]{#record-timing-info-for-timer-handler
      .dfn}
    - [[record timing info for microtask
      checkpoint](https://w3c.github.io/long-animation-frames/#record-timing-info-for-microtask-checkpoint)]{#record-timing-info-for-microtask-checkpoint
      .dfn}

Long Tasks

:   The following terms are defined in Long Tasks:
    [\[LONGTASKS\]](references.html#refsLONGTASKS)

    - [[report long
      tasks](https://w3c.github.io/longtasks/#report-long-tasks)]{#report-long-tasks
      .dfn}

Web IDL

:   The IDL fragments in this specification must be interpreted as
    required for conforming IDL fragments, as described in Web IDL.
    [\[WEBIDL\]](references.html#refsWEBIDL)

    The following terms are defined in Web IDL:

    - [[this](https://webidl.spec.whatwg.org/#this)]{#this .dfn}
    - [[extended
      attribute](https://webidl.spec.whatwg.org/#dfn-extended-attribute)]{#extended-attribute
      .dfn}
    - [[named
      constructor](https://webidl.spec.whatwg.org/#dfn-named-constructor)]{#named-constructor
      .dfn}
    - [[constructor
      operation](https://webidl.spec.whatwg.org/#idl-constructors)]{#constructor-operation
      .dfn}
    - [[overridden constructor
      steps](https://webidl.spec.whatwg.org/#overridden-constructor-steps)]{#overridden-constructor-steps
      .dfn}
    - [[internally create a new object implementing the
      interface](https://webidl.spec.whatwg.org/#internally-create-a-new-object-implementing-the-interface)]{#internally-create-a-new-object-implementing-the-interface
      .dfn}
    - [[array index property
      name](https://webidl.spec.whatwg.org/#dfn-array-index-property-name)]{#array-index-property-name
      .dfn}
    - [[buffer source byte
      length](https://webidl.spec.whatwg.org/#buffersource-byte-length)]{#buffer-source-byte-length
      .dfn}
    - [[supports indexed
      properties](https://webidl.spec.whatwg.org/#dfn-support-indexed-properties)]{#supports-indexed-properties
      .dfn}
    - [[supported property
      indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices)]{#supported-property-indices
      .dfn}
    - [[determine the value of an indexed
      property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-an-indexed-property)]{#determine-the-value-of-an-indexed-property
      .dfn}
    - [[set the value of an existing indexed
      property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-an-existing-indexed-property)]{#set-the-value-of-an-existing-indexed-property
      .dfn}
    - [[set the value of a new indexed
      property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-a-new-indexed-property)]{#set-the-value-of-a-new-indexed-property
      .dfn}
    - [[support named
      properties](https://webidl.spec.whatwg.org/#dfn-support-named-properties)]{#support-named-properties
      .dfn}
    - [[supported property
      names](https://webidl.spec.whatwg.org/#dfn-supported-property-names)]{#supported-property-names
      .dfn}
    - [[determine the value of a named
      property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-a-named-property)]{#determine-the-value-of-a-named-property
      .dfn}
    - [[set the value of an existing named
      property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-an-existing-named-property)]{#set-the-value-of-an-existing-named-property
      .dfn}
    - [[set the value of a new named
      property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-a-new-named-property)]{#set-the-value-of-a-new-named-property
      .dfn}
    - [[delete an existing named
      property](https://webidl.spec.whatwg.org/#dfn-delete-an-existing-named-property)]{#delete-an-existing-named-property
      .dfn}
    - [[perform a security
      check](https://webidl.spec.whatwg.org/#dfn-perform-a-security-check)]{#perform-a-security-check
      .dfn}
    - [[platform
      object](https://webidl.spec.whatwg.org/#dfn-platform-object)]{#platform-object
      .dfn}
    - [[legacy platform
      object](https://webidl.spec.whatwg.org/#dfn-legacy-platform-object)]{#legacy-platform-object
      .dfn}
    - [[primary
      interface](https://webidl.spec.whatwg.org/#dfn-primary-interface)]{#primary-interface
      .dfn}
    - [[interface
      object](https://webidl.spec.whatwg.org/#dfn-interface-object)]{#interface-object
      .dfn}
    - [[named properties
      object](https://webidl.spec.whatwg.org/#dfn-named-properties-object)]{#named-properties-object
      .dfn}
    - [[include](https://webidl.spec.whatwg.org/#include)]{#include
      .dfn}
    - [[inherit](https://webidl.spec.whatwg.org/#dfn-inherit)]{#inherit
      .dfn}
    - [[interface prototype
      object](https://webidl.spec.whatwg.org/#dfn-interface-prototype-object)]{#interface-prototype-object
      .dfn}
    - [[implements](https://webidl.spec.whatwg.org/#implements)]{#implements
      .dfn}
    - [[associated
      realm](https://webidl.spec.whatwg.org/#dfn-associated-realm)]{#associated-realm
      .dfn}
    - [[\[\[Realm\]\] field of a platform
      object](https://webidl.spec.whatwg.org/#es-platform-objects)]{#realm-field-of-a-platform-object
      .dfn}
    - [[\[\[GetOwnProperty\]\] internal method of a named properties
      object](https://webidl.spec.whatwg.org/#named-properties-object-getownproperty)]{#named-properties-object-getownproperty
      .dfn}
    - [[callback
      context](https://webidl.spec.whatwg.org/#dfn-callback-context)]{#callback-context
      .dfn}
    - [[frozen
      array](https://webidl.spec.whatwg.org/#dfn-frozen-array-type)]{#frozen-array
      .dfn} and [[creating a frozen
      array](https://webidl.spec.whatwg.org/#dfn-create-frozen-array)]{#creating-a-frozen-array
      .dfn}
    - [[create a new object implementing the
      interface](https://webidl.spec.whatwg.org/#new)]{#new .dfn}
    - [[callback this
      value](https://webidl.spec.whatwg.org/#dfn-callback-this-value)]{#dfn-callback-this-value
      .dfn}
    - [[converting](https://webidl.spec.whatwg.org/#es-type-mapping)]{#concept-idl-convert
      .dfn} between Web IDL types and JS types
    - [[invoking](https://webidl.spec.whatwg.org/#invoke-a-callback-function)]{#es-invoking-callback-functions
      .dfn} and
      [[constructing](https://webidl.spec.whatwg.org/#construct-a-callback-function)]{#es-constructing-callback-functions
      .dfn} callback functions
    - [[overload resolution
      algorithm](https://webidl.spec.whatwg.org/#dfn-overload-resolution-algorithm)]{#overload-resolution-algorithm
      .dfn}
    - [[exposed](https://webidl.spec.whatwg.org/#dfn-exposed)]{#idl-exposed
      .dfn}
    - [[a promise resolved
      with](https://webidl.spec.whatwg.org/#a-promise-resolved-with)]{#a-promise-resolved-with
      .dfn}
    - [[a promise rejected
      with](https://webidl.spec.whatwg.org/#a-promise-rejected-with)]{#a-promise-rejected-with
      .dfn}
    - [[wait for
      all](https://webidl.spec.whatwg.org/#wait-for-all)]{#wait-for-all
      .dfn}
    - [[upon
      rejection](https://webidl.spec.whatwg.org/#upon-rejection)]{#upon-rejection
      .dfn}
    - [[upon
      fulfillment](https://webidl.spec.whatwg.org/#upon-fulfillment)]{#upon-fulfillment
      .dfn}
    - [[mark as
      handled](https://webidl.spec.whatwg.org/#mark-a-promise-as-handled)]{#mark-as-handled
      .dfn}
    - [[`[Global]`](https://webidl.spec.whatwg.org/#Global)]{#global
      .dfn}
    - [[`[LegacyFactoryFunction]`](https://webidl.spec.whatwg.org/#LegacyFactoryFunction)]{#legacyfactoryfunction
      .dfn}
    - [[`[LegacyLenientThis]`](https://webidl.spec.whatwg.org/#LegacyLenientThis)]{#legacylenientthis
      .dfn}
    - [[`[LegacyNullToEmptyString]`](https://webidl.spec.whatwg.org/#LegacyNullToEmptyString)]{#legacynulltoemptystring
      .dfn}
    - [[`[LegacyOverrideBuiltIns]`](https://webidl.spec.whatwg.org/#LegacyOverrideBuiltIns)]{#legacyoverridebuiltins
      .dfn}
    - [[LegacyPlatformObjectGetOwnProperty](https://webidl.spec.whatwg.org/#LegacyPlatformObjectGetOwnProperty)]{#legacyplatformobjectgetownproperty
      .dfn}
    - [[`[LegacyTreatNonObjectAsNull]`](https://webidl.spec.whatwg.org/#LegacyTreatNonObjectAsNull)]{#legacytreatnonobjectasnull
      .dfn}
    - [[`[LegacyUnenumerableNamedProperties]`](https://webidl.spec.whatwg.org/#LegacyUnenumerableNamedProperties)]{#legacyunenumerablenamedproperties
      .dfn}
    - [[`[LegacyUnforgeable]`](https://webidl.spec.whatwg.org/#LegacyUnforgeable)]{#legacyunforgeable
      .dfn}
    - [[set
      entries](https://webidl.spec.whatwg.org/#dfn-set-entries)]{#set-entries
      .dfn}

    Web IDL also defines the following types that are used in Web IDL
    fragments in this specification:

    - [[`ArrayBuffer`](https://webidl.spec.whatwg.org/#idl-ArrayBuffer)]{#idl-arraybuffer
      .dfn}
    - [[`ArrayBufferView`](https://webidl.spec.whatwg.org/#common-ArrayBufferView)]{#idl-arraybufferview
      .dfn}
    - [[`boolean`](https://webidl.spec.whatwg.org/#idl-boolean)]{#idl-boolean
      .dfn}
    - [[`DOMString`](https://webidl.spec.whatwg.org/#idl-DOMString)]{#idl-domstring
      .dfn}
    - [[`double`](https://webidl.spec.whatwg.org/#idl-double)]{#idl-double
      .dfn}
    - [[enumeration](https://webidl.spec.whatwg.org/#idl-enums)]{#idl-enumeration
      .dfn}
    - [[`Float16Array`](https://webidl.spec.whatwg.org/#idl-Float16Array)]{#idl-float16array
      .dfn}
    - [[`Function`](https://webidl.spec.whatwg.org/#common-Function)]{#idl-function
      .dfn}
    - [[`long`](https://webidl.spec.whatwg.org/#idl-long)]{#idl-long
      .dfn}
    - [[`object`](https://webidl.spec.whatwg.org/#idl-object)]{#idl-object
      .dfn}
    - [[`Promise`](https://webidl.spec.whatwg.org/#idl-promise)]{#idl-promise
      .dfn}
    - [[`Uint8ClampedArray`](https://webidl.spec.whatwg.org/#idl-Uint8ClampedArray)]{#idl-uint8clampedarray
      .dfn}
    - [[`unrestricted double`](https://webidl.spec.whatwg.org/#idl-unrestricted-double)]{#idl-unrestricted-double
      .dfn}
    - [[`unsigned long`](https://webidl.spec.whatwg.org/#idl-unsigned-long)]{#idl-unsigned-long
      .dfn}
    - [[`USVString`](https://webidl.spec.whatwg.org/#idl-USVString)]{#idl-usvstring
      .dfn}
    - [[`VoidFunction`](https://webidl.spec.whatwg.org/#VoidFunction)]{#idl-voidfunction
      .dfn}

    The term [[throw](https://webidl.spec.whatwg.org/#dfn-throw)]{#throw
    .dfn} in this specification is used as defined in Web IDL. The
    [[`DOMException`](https://webidl.spec.whatwg.org/#dfn-DOMException)]{#domexception
    .dfn} type and the following exception names are defined by Web IDL
    and used by this specification:

    - [[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror)]{#indexsizeerror
      .dfn}
    - [[\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror)]{#hierarchyrequesterror
      .dfn}
    - [[\"`InvalidCharacterError`\"](https://webidl.spec.whatwg.org/#invalidcharactererror)]{#invalidcharactererror
      .dfn}
    - [[\"`NoModificationAllowedError`\"](https://webidl.spec.whatwg.org/#nomodificationallowederror)]{#nomodificationallowederror
      .dfn}
    - [[\"`NotFoundError`\"](https://webidl.spec.whatwg.org/#notfounderror)]{#notfounderror
      .dfn}
    - [[\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror)]{#notsupportederror
      .dfn}
    - [[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror)]{#invalidstateerror
      .dfn}
    - [[\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror)]{#syntaxerror
      .dfn}
    - [[\"`InvalidAccessError`\"](https://webidl.spec.whatwg.org/#invalidaccesserror)]{#invalidaccesserror
      .dfn}
    - [[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror)]{#securityerror
      .dfn}
    - [[\"`NetworkError`\"](https://webidl.spec.whatwg.org/#networkerror)]{#networkerror
      .dfn}
    - [[\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror)]{#aborterror
      .dfn}
    - [[\"`QuotaExceededError`\"](https://webidl.spec.whatwg.org/#quotaexceedederror)]{#quotaexceedederror
      .dfn}
    - [[\"`DataCloneError`\"](https://webidl.spec.whatwg.org/#datacloneerror)]{#datacloneerror
      .dfn}
    - [[\"`EncodingError`\"](https://webidl.spec.whatwg.org/#encodingerror)]{#encodingerror
      .dfn}
    - [[\"`NotAllowedError`\"](https://webidl.spec.whatwg.org/#notallowederror)]{#notallowederror
      .dfn}

    When this specification requires a user agent to [create a `Date`
    object]{#create-a-date-object .dfn} representing a particular time
    (which could be the special value Not-a-Number), the milliseconds
    component of that time, if any, must be truncated to an integer, and
    the time value of the newly created
    [`Date`{#dependencies:date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
    object must represent the resulting truncated time.

    For instance, given the time 23045 millionths of a second after
    01:00 UTC on January 1st 2000, i.e. the time
    2000-01-01T00:00:00.023045Z, then the
    [`Date`{#dependencies:date-2}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
    object created representing that time would represent the same time
    as that created representing the time 2000-01-01T00:00:00.023Z, 45
    millionths earlier. If the given time is NaN, then the result is a
    [`Date`{#dependencies:date-3}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
    object that represents a time value NaN (indicating that the object
    does not represent a specific instant of time).

JavaScript

:   Some parts of the language described by this specification only
    support JavaScript as the underlying scripting language.
    [\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

    The term \"JavaScript\" is used to refer to ECMA-262, rather than
    the official term ECMAScript, since the term JavaScript is more
    widely known.

    The following terms are defined in the JavaScript specification and
    used in this specification:

    - [[active function
      object](https://tc39.es/ecma262/#active-function-object)]{#active-function-object
      .dfn}
    - [[agent](https://tc39.es/ecma262/#sec-agents)]{#agent .dfn} and
      [[agent
      cluster](https://tc39.es/ecma262/#sec-agent-clusters)]{#agent-cluster
      .dfn}
    - [[automatic semicolon
      insertion](https://tc39.es/ecma262/#sec-automatic-semicolon-insertion)]{#automatic-semicolon-insertion
      .dfn}
    - [[candidate
      execution](https://tc39.es/ecma262/#sec-candidate-executions)]{#candidate-execution
      .dfn}
    - The [[current
      realm](https://tc39.es/ecma262/#current-realm)]{#current-realm
      .dfn}
    - [[clamping](https://tc39.es/ecma262/#clamping)]{#clamping .dfn} a
      mathematical value
    - [[early
      error](https://tc39.es/ecma262/#early-error-rule)]{#early-error
      .dfn}
    - [[forward
      progress](https://tc39.es/ecma262/#sec-forward-progress)]{#forward-progress
      .dfn}
    - [[invariants of the essential internal
      methods](https://tc39.es/ecma262/#sec-invariants-of-the-essential-internal-methods)]{#invariants-of-the-essential-internal-methods
      .dfn}
    - [[JavaScript execution
      context](https://tc39.es/ecma262/#sec-execution-contexts)]{#javascript-execution-context
      .dfn}
    - [[JavaScript execution context
      stack](https://tc39.es/ecma262/#execution-context-stack)]{#javascript-execution-context-stack
      .dfn}
    - [[realm](https://tc39.es/ecma262/#sec-code-realms)]{#realm .dfn}
    - [[JobCallback
      Record](https://tc39.es/ecma262/#sec-jobcallback-records)]{#jobcallback-record
      .dfn}
    - [[NewTarget](https://tc39.es/ecma262/#sec-built-in-function-objects)]{#newtarget
      .dfn}
    - [[running JavaScript execution
      context](https://tc39.es/ecma262/#running-execution-context)]{#running-javascript-execution-context
      .dfn}
    - [[surrounding
      agent](https://tc39.es/ecma262/#surrounding-agent)]{#surrounding-agent
      .dfn}
    - [[abstract
      closure](https://tc39.es/ecma262/#sec-abstract-closure)]{#abstract-closure
      .dfn}
    - [[immutable prototype exotic
      object](https://tc39.es/ecma262/#immutable-prototype-exotic-object)]{#immutable-prototype-exotic-object
      .dfn}
    - [[Well-Known
      Symbols](https://tc39.es/ecma262/#sec-well-known-symbols)]{#well-known-symbols
      .dfn}, including [%Symbol.hasInstance%]{#symbol.hasinstance .dfn},
      [%Symbol.isConcatSpreadable%]{#symbol.isconcatspreadable .dfn},
      [%Symbol.toPrimitive%]{#symbol.toprimitive .dfn}, and
      [%Symbol.toStringTag%]{#symbol.tostringtag .dfn}
    - [[Well-Known Intrinsic
      Objects](https://tc39.es/ecma262/#sec-well-known-intrinsic-objects)]{#well-known-intrinsic-objects
      .dfn}, including
      [[%Array.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-array-prototype-object)]{#array.prototype
      .dfn},
      [[%Error.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-error-prototype-object)]{#error.prototype
      .dfn}, [%EvalError.prototype%]{#evalerror.prototype .dfn},
      [[%Function.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-function-prototype-object)]{#function.prototype
      .dfn},
      [[%Object.prototype%](https://tc39.es/ecma262/#sec-properties-of-the-object-prototype-object)]{#object.prototype
      .dfn},
      [[%Object.prototype.valueOf%](https://tc39.es/ecma262/#sec-object.prototype.valueof)]{#object.prototype.valueof
      .dfn}, [%RangeError.prototype%]{#rangeerror.prototype .dfn},
      [%ReferenceError.prototype%]{#referenceerror.prototype .dfn},
      [%SyntaxError.prototype%]{#syntaxerror.prototype .dfn},
      [%TypeError.prototype%]{#typeerror.prototype .dfn}, and
      [%URIError.prototype%]{#urierror.prototype .dfn}
    - The
      [[*FunctionBody*](https://tc39.es/ecma262/#prod-FunctionBody)]{#js-prod-functionbody
      .dfn} production
    - The
      [[*Module*](https://tc39.es/ecma262/#prod-Module)]{#js-prod-module
      .dfn} production
    - The
      [[*Pattern*](https://tc39.es/ecma262/#prod-Pattern)]{#js-prod-pattern
      .dfn} production
    - The
      [[*Script*](https://tc39.es/ecma262/#prod-Script)]{#js-prod-script
      .dfn} production
    - The
      [[BigInt](https://tc39.es/ecma262/#sec-ecmascript-language-types-bigint-type)]{#js-bigint
      .dfn},
      [[Boolean](https://tc39.es/ecma262/#sec-ecmascript-language-types-boolean-type)]{#js-boolean
      .dfn},
      [[Number](https://tc39.es/ecma262/#sec-ecmascript-language-types-number-type)]{#js-number
      .dfn},
      [[String](https://tc39.es/ecma262/#sec-ecmascript-language-types-string-type)]{#js-string
      .dfn},
      [[Symbol](https://tc39.es/ecma262/#sec-ecmascript-language-types-symbol-type)]{#js-symbol
      .dfn}, and
      [[Object](https://tc39.es/ecma262/#sec-object-type)]{#js-object
      .dfn} ECMAScript language types
    - The [[Completion
      Record](https://tc39.es/ecma262/#sec-completion-record-specification-type)]{#completion-record
      .dfn} specification type
    - The
      [[List](https://tc39.es/ecma262/#sec-list-and-record-specification-type)]{#js-list
      .dfn} and
      [[Record](https://tc39.es/ecma262/#sec-list-and-record-specification-type)]{#record
      .dfn} specification types
    - The [[Property
      Descriptor](https://tc39.es/ecma262/#sec-property-descriptor-specification-type)]{#propertydescriptor
      .dfn} specification type
    - The [[ModuleRequest
      Record](https://tc39.es/ecma262/#modulerequest-record)]{#modulerequest-record
      .dfn} specification type
    - The [[Script
      Record](https://tc39.es/ecma262/#sec-script-records)]{#script-record
      .dfn} specification type
    - The [[Synthetic Module
      Record](https://tc39.es/ecma262/#sec-synthetic-module-records)]{#synthetic-module-record
      .dfn} specification type
    - The [[Cyclic Module
      Record](https://tc39.es/ecma262/#sec-cyclic-module-records)]{#cyclic-module-record
      .dfn} specification type
    - The [[Source Text Module
      Record](https://tc39.es/ecma262/#sec-source-text-module-records)]{#source-text-module-record
      .dfn} specification type and its
      [[Evaluate](https://tc39.es/ecma262/#sec-moduleevaluation)]{#js-evaluate
      .dfn},
      [[Link](https://tc39.es/ecma262/#sec-moduledeclarationlinking)]{#js-link
      .dfn} and
      [[LoadRequestedModules](https://tc39.es/ecma262/#sec-LoadRequestedModules)]{#js-loadrequestedmodules
      .dfn} methods
    - The
      [[ArrayCreate](https://tc39.es/ecma262/#sec-arraycreate)]{#arraycreate
      .dfn} abstract operation
    - The [[Call](https://tc39.es/ecma262/#sec-call)]{#call .dfn}
      abstract operation
    - The
      [[ClearKeptObjects](https://tc39.es/ecma262/#sec-clear-kept-objects)]{#clearkeptobjects
      .dfn} abstract operation
    - The
      [[CleanupFinalizationRegistry](https://tc39.es/ecma262/#sec-cleanup-finalization-registry)]{#cleanupfinalizationregistry
      .dfn} abstract operation
    - The
      [[Construct](https://tc39.es/ecma262/#sec-construct)]{#construct
      .dfn} abstract operation
    - The
      [[CopyDataBlockBytes](https://tc39.es/ecma262/#sec-copydatablockbytes)]{#copydatablockbytes
      .dfn} abstract operation
    - The
      [[CreateBuiltinFunction](https://tc39.es/ecma262/#sec-createbuiltinfunction)]{#createbuiltinfunction
      .dfn} abstract operation
    - The
      [[CreateByteDataBlock](https://tc39.es/ecma262/#sec-createbytedatablock)]{#createbytedatablock
      .dfn} abstract operation
    - The
      [[CreateDataProperty](https://tc39.es/ecma262/#sec-createdataproperty)]{#createdataproperty
      .dfn} abstract operation
    - The
      [[CreateDefaultExportSyntheticModule](https://tc39.es/ecma262/#sec-create-default-export-synthetic-module)]{#createdefaultexportsyntheticmodule
      .dfn} abstract operation
    - The
      [[DefinePropertyOrThrow](https://tc39.es/ecma262/#sec-definepropertyorthrow)]{#definepropertyorthrow
      .dfn} abstract operation
    - The
      [[DetachArrayBuffer](https://tc39.es/ecma262/#sec-detacharraybuffer)]{#detacharraybuffer
      .dfn} abstract operation
    - The
      [[EnumerableOwnProperties](https://tc39.es/ecma262/#sec-enumerableownproperties)]{#enumerableownproperties
      .dfn} abstract operation
    - The
      [[FinishLoadingImportedModule](https://tc39.es/ecma262/#sec-FinishLoadingImportedModule)]{#finishloadingimportedmodule
      .dfn} abstract operation
    - The
      [[OrdinaryFunctionCreate](https://tc39.es/ecma262/#sec-ordinaryfunctioncreate)]{#js-ordinaryfunctioncreate
      .dfn} abstract operation
    - The [[Get](https://tc39.es/ecma262/#sec-get-o-p)]{#js-get .dfn}
      abstract operation
    - The
      [[GetActiveScriptOrModule](https://tc39.es/ecma262/#sec-getactivescriptormodule)]{#getactivescriptormodule
      .dfn} abstract operation
    - The
      [[GetFunctionRealm](https://tc39.es/ecma262/#sec-getfunctionrealm)]{#getfunctionrealm
      .dfn} abstract operation
    - The
      [[HasOwnProperty](https://tc39.es/ecma262/#sec-hasownproperty)]{#hasownproperty
      .dfn} abstract operation
    - The
      [[HostCallJobCallback](https://tc39.es/ecma262/#sec-hostcalljobcallback)]{#js-hostcalljobcallback
      .dfn} abstract operation
    - The
      [[HostEnqueueFinalizationRegistryCleanupJob](https://tc39.es/ecma262/#sec-host-cleanup-finalization-registry)]{#js-hostenqueuefinalizationregistrycleanupjob
      .dfn} abstract operation
    - The
      [[HostEnqueueGenericJob](https://tc39.es/ecma262/#sec-hostenqueuegenericjob)]{#js-hostenqueuegenericjob
      .dfn} abstract operation
    - The
      [[HostEnqueuePromiseJob](https://tc39.es/ecma262/#sec-hostenqueuepromisejob)]{#js-hostenqueuepromisejob
      .dfn} abstract operation
    - The
      [[HostEnqueueTimeoutJob](https://tc39.es/ecma262/#sec-hostenqueuetimeoutjob)]{#js-hostenqueuetimeoutjob
      .dfn} abstract operation
    - The
      [[HostEnsureCanAddPrivateElement](https://tc39.es/ecma262/#sec-hostensurecanaddprivateelement)]{#js-hostensurecanaddprivateelement
      .dfn} abstract operation
    - The
      [[HostGetSupportedImportAttributes](https://tc39.es/ecma262/#sec-hostgetsupportedimportattributes)]{#js-hostgetsupportedimportattributes
      .dfn} abstract operation
    - The
      [[HostLoadImportedModule](https://tc39.es/ecma262/#sec-HostLoadImportedModule)]{#js-hostloadimportedmodule
      .dfn} abstract operation
    - The
      [[HostMakeJobCallback](https://tc39.es/ecma262/#sec-hostmakejobcallback)]{#js-hostmakejobcallback
      .dfn} abstract operation
    - The
      [[HostPromiseRejectionTracker](https://tc39.es/ecma262/#sec-host-promise-rejection-tracker)]{#js-hostpromiserejectiontracker
      .dfn} abstract operation
    - The
      [[InitializeHostDefinedRealm](https://tc39.es/ecma262/#sec-initializehostdefinedrealm)]{#js-initializehostdefinedrealm
      .dfn} abstract operation
    - The
      [[IsArrayBufferViewOutOfBounds](https://tc39.es/ecma262/#sec-isarraybufferviewoutofbounds)]{#isarraybufferviewoutofbounds
      .dfn} abstract operation
    - The
      [[IsAccessorDescriptor](https://tc39.es/ecma262/#sec-isaccessordescriptor)]{#isaccessordescriptor
      .dfn} abstract operation
    - The
      [[IsCallable](https://tc39.es/ecma262/#sec-iscallable)]{#iscallable
      .dfn} abstract operation
    - The
      [[IsConstructor](https://tc39.es/ecma262/#sec-isconstructor)]{#isconstructor
      .dfn} abstract operation
    - The
      [[IsDataDescriptor](https://tc39.es/ecma262/#sec-isdatadescriptor)]{#isdatadescriptor
      .dfn} abstract operation
    - The
      [[IsDetachedBuffer](https://tc39.es/ecma262/#sec-isdetachedbuffer)]{#isdetachedbuffer
      .dfn} abstract operation
    - The
      [[IsSharedArrayBuffer](https://tc39.es/ecma262/#sec-issharedarraybuffer)]{#issharedarraybuffer
      .dfn} abstract operation
    - The
      [[NewObjectEnvironment](https://tc39.es/ecma262/#sec-newobjectenvironment)]{#js-newobjectenvironment
      .dfn} abstract operation
    - The
      [[NormalCompletion](https://tc39.es/ecma262/#sec-normalcompletion)]{#normalcompletion
      .dfn} abstract operation
    - The
      [[OrdinaryGetPrototypeOf](https://tc39.es/ecma262/#sec-ordinarygetprototypeof)]{#ordinarygetprototypeof
      .dfn} abstract operation
    - The
      [[OrdinarySetPrototypeOf](https://tc39.es/ecma262/#sec-ordinarysetprototypeof)]{#ordinarysetprototypeof
      .dfn} abstract operation
    - The
      [[OrdinaryIsExtensible](https://tc39.es/ecma262/#sec-ordinaryisextensible)]{#ordinaryisextensible
      .dfn} abstract operation
    - The
      [[OrdinaryPreventExtensions](https://tc39.es/ecma262/#sec-ordinarypreventextensions)]{#ordinarypreventextensions
      .dfn} abstract operation
    - The
      [[OrdinaryGetOwnProperty](https://tc39.es/ecma262/#sec-ordinarygetownproperty)]{#ordinarygetownproperty
      .dfn} abstract operation
    - The
      [[OrdinaryDefineOwnProperty](https://tc39.es/ecma262/#sec-ordinarydefineownproperty)]{#ordinarydefineownproperty
      .dfn} abstract operation
    - The
      [[OrdinaryGet](https://tc39.es/ecma262/#sec-ordinaryget)]{#ordinaryget
      .dfn} abstract operation
    - The
      [[OrdinarySet](https://tc39.es/ecma262/#sec-ordinaryset)]{#ordinaryset
      .dfn} abstract operation
    - The
      [[OrdinaryDelete](https://tc39.es/ecma262/#sec-ordinarydelete)]{#ordinarydelete
      .dfn} abstract operation
    - The
      [[OrdinaryOwnPropertyKeys](https://tc39.es/ecma262/#sec-ordinaryownpropertykeys)]{#ordinaryownpropertykeys
      .dfn} abstract operation
    - The
      [[OrdinaryObjectCreate](https://tc39.es/ecma262/#sec-objectcreate)]{#ordinaryobjectcreate
      .dfn} abstract operation
    - The
      [[ParseJSONModule](https://tc39.es/ecma262/#sec-parse-json-module)]{#parsejsonmodule
      .dfn} abstract operation
    - The
      [[ParseModule](https://tc39.es/ecma262/#sec-parsemodule)]{#js-parsemodule
      .dfn} abstract operation
    - The
      [[ParseScript](https://tc39.es/ecma262/#sec-parse-script)]{#js-parsescript
      .dfn} abstract operation
    - The
      [[NewPromiseReactionJob](https://tc39.es/ecma262/#sec-newpromisereactionjob)]{#newpromisereactionjob
      .dfn} abstract operation
    - The
      [[NewPromiseResolveThenableJob](https://tc39.es/ecma262/#sec-newpromiseresolvethenablejob)]{#newpromiseresolvethenablejob
      .dfn} abstract operation
    - The
      [[RegExpBuiltinExec](https://tc39.es/ecma262/#sec-regexpbuiltinexec)]{#regexpbuiltinexec
      .dfn} abstract operation
    - The
      [[RegExpCreate](https://tc39.es/ecma262/#sec-regexpcreate)]{#regexpcreate
      .dfn} abstract operation
    - The [[RunJobs](https://tc39.es/ecma262/#sec-runjobs)]{#runjobs
      .dfn} abstract operation
    - The
      [[SameValue](https://tc39.es/ecma262/#sec-samevalue)]{#samevalue
      .dfn} abstract operation
    - The
      [[ScriptEvaluation](https://tc39.es/ecma262/#sec-runtime-semantics-scriptevaluation)]{#js-scriptevaluation
      .dfn} abstract operation
    - The
      [[SetSyntheticModuleExport](https://tc39.es/ecma262/#sec-setsyntheticmoduleexport)]{#setsyntheticmoduleexport
      .dfn} abstract operation
    - The
      [[SetImmutablePrototype](https://tc39.es/ecma262/#sec-set-immutable-prototype)]{#setimmutableprototype
      .dfn} abstract operation
    - The
      [[ToBoolean](https://tc39.es/ecma262/#sec-toboolean)]{#js-toboolean
      .dfn} abstract operation
    - The [[ToString](https://tc39.es/ecma262/#sec-tostring)]{#tostring
      .dfn} abstract operation
    - The [[ToUint32](https://tc39.es/ecma262/#sec-touint32)]{#touint32
      .dfn} abstract operation
    - The
      [[TypedArrayCreate](https://tc39.es/ecma262/#typedarray-create)]{#typedarraycreate
      .dfn} abstract operation
    - The
      [[IsLooselyEqual](https://tc39.es/ecma262/#sec-islooselyequal)]{#js-abstract-equality
      .dfn} abstract operation
    - The
      [[IsStrictlyEqual](https://tc39.es/ecma262/#sec-isstrictlyequal)]{#js-strict-equality
      .dfn} abstract operation
    - The
      [[`Atomics`](https://tc39.es/ecma262/#sec-atomics-object)]{#atomics
      .dfn} object
    - The
      [[`Atomics.waitAsync`](https://tc39.es/ecma262/#sec-atomics.waitasync)]{#atomics.waitasync
      .dfn} object
    - The [[`Date`](https://tc39.es/ecma262/#sec-date-objects)]{#date
      .dfn} class
    - The
      [[`FinalizationRegistry`](https://tc39.es/ecma262/#sec-finalization-registry-objects)]{#finalizationregistry
      .dfn} class
    - The
      [[`RegExp`](https://tc39.es/ecma262/#sec-regexp-regular-expression-objects)]{#regexp
      .dfn} class
    - The
      [[`SharedArrayBuffer`](https://tc39.es/ecma262/#sec-sharedarraybuffer-objects)]{#sharedarraybuffer
      .dfn} class
    - The
      [[`SyntaxError`](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-syntaxerror)]{#syntaxerror-2
      .dfn} class
    - The
      [[`TypeError`](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror)]{#typeerror
      .dfn} class
    - The
      [[`RangeError`](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-rangeerror)]{#js-rangeerror
      .dfn} class
    - The
      [[`WeakRef`](https://tc39.es/ecma262/#sec-weak-ref-objects)]{#weakref
      .dfn} class
    - The [[`eval()`](https://tc39.es/ecma262/#sec-eval-x)]{#eval()
      .dfn} function
    - The
      [[`WeakRef.prototype.deref()`](https://tc39.es/ecma262/#sec-weak-ref.prototype.deref)]{#weakref.prototype.deref()
      .dfn} function
    - The
      [[\[\[IsHTMLDDA\]\]](https://tc39.es/ecma262/#sec-IsHTMLDDA-internal-slot)]{#ishtmldda
      .dfn} internal slot
    - [[`import()`](https://tc39.es/ecma262/#sec-import-calls)]{#import()
      .dfn}
    - [[`import.meta`](https://tc39.es/ecma262/#sec-meta-properties)]{#import.meta
      .dfn}
    - The
      [[HostGetImportMetaProperties](https://tc39.es/ecma262/#sec-hostgetimportmetaproperties)]{#js-hostgetimportmetaproperties
      .dfn} abstract operation
    - The
      [[`typeof`](https://tc39.es/ecma262/#sec-typeof-operator)]{#js-typeof
      .dfn} operator
    - The
      [[`delete`](https://tc39.es/ecma262/#sec-delete-operator)]{#delete
      .dfn} operator
    - [[The `TypedArray`{.variable}
      Constructors](https://tc39.es/ecma262/#table-49)]{#the-typedarray-constructors
      .dfn} table

    User agents that support JavaScript must also implement the Dynamic
    Code Brand Checks proposal. The following terms are defined there,
    and used in this specification:
    [\[JSDYNAMICCODEBRANDCHECKS\]](references.html#refsJSDYNAMICCODEBRANDCHECKS)

    - The
      [[HostEnsureCanCompileStrings](https://tc39.es/proposal-dynamic-code-brand-checks/#sec-hostensurecancompilestrings)]{#js-hostensurecancompilestrings
      .dfn} abstract operation
    - The
      [[HostGetCodeForEval](https://tc39.es/proposal-dynamic-code-brand-checks/#sec-hostgetcodeforeval)]{#js-hostgetcodeforeval
      .dfn} abstract operation

    User agents that support JavaScript must also implement ECMAScript
    Internationalization API. [\[JSINTL\]](references.html#refsJSINTL)

    User agents that support JavaScript must also implement the Temporal
    proposal. The following terms are defined there, and used in this
    specification: [\[JSTEMPORAL\]](references.html#refsJSTEMPORAL)

    - The
      [[HostSystemUTCEpochNanoseconds](https://tc39.es/proposal-temporal/#sec-hostsystemutcepochnanoseconds)]{#js-hostsystemutcepochnanoseconds
      .dfn} abstract operation
    - The
      [[nsMaxInstant](https://tc39.es/proposal-temporal/#eqn-nsMaxInstant)]{#nsmaxinstant
      .dfn} and
      [[nsMinInstant](https://tc39.es/proposal-temporal/#eqn-nsMinInstant)]{#nsmininstant
      .dfn} values

WebAssembly

:   The following term is defined in WebAssembly JavaScript Interface:
    [\[WASMJS\]](references.html#refsWASMJS)

    - [[`WebAssembly.Module`](https://webassembly.github.io/spec/js-api/#module)]{#webassembly.module
      .dfn}

DOM

:   The Document Object Model (DOM) is a representation --- a model ---
    of a document and its content. The DOM is not just an API; the
    conformance criteria of HTML implementations are defined, in this
    specification, in terms of operations on the DOM.
    [\[DOM\]](references.html#refsDOM)

    Implementations must support DOM and the events defined in UI
    Events, because this specification is defined in terms of the DOM,
    and some of the features are defined as extensions to the DOM
    interfaces. [\[DOM\]](references.html#refsDOM)
    [\[UIEVENTS\]](references.html#refsUIEVENTS)

    In particular, the following features are defined in DOM:
    [\[DOM\]](references.html#refsDOM)

    - [[`Attr`](https://dom.spec.whatwg.org/#interface-attr)]{#attr
      .dfn} interface
    - [[`CharacterData`](https://dom.spec.whatwg.org/#interface-characterdata)]{#characterdata
      .dfn} interface
    - [[`Comment`](https://dom.spec.whatwg.org/#interface-comment)]{#comment-2
      .dfn} interface
    - [[`DOMImplementation`](https://dom.spec.whatwg.org/#interface-domimplementation)]{#domimplementation
      .dfn} interface
    - [[`Document`](https://dom.spec.whatwg.org/#interface-document)]{#dom-document
      .dfn} interface and its
      [[`doctype`](https://dom.spec.whatwg.org/#dom-document-doctype)]{#dom-document-doctype
      .dfn} attribute
    - [[`DocumentOrShadowRoot`](https://dom.spec.whatwg.org/#documentorshadowroot)]{#dom-documentorshadowroot
      .dfn} interface
    - [[`DocumentFragment`](https://dom.spec.whatwg.org/#interface-documentfragment)]{#documentfragment
      .dfn} interface
    - [[`DocumentType`](https://dom.spec.whatwg.org/#interface-documenttype)]{#documenttype
      .dfn} interface
    - [[`ChildNode`](https://dom.spec.whatwg.org/#interface-childnode)]{#childnode
      .dfn} interface
    - [[`Element`](https://dom.spec.whatwg.org/#interface-element)]{#element
      .dfn} interface
    - [[`attachShadow()`](https://dom.spec.whatwg.org/#dom-element-attachshadow)]{#dom-element-attachshadow
      .dfn} method.
    - An element\'s [[shadow
      root](https://dom.spec.whatwg.org/#concept-element-shadow-root)]{#concept-element-shadow-root
      .dfn}
    - A [shadow
      root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#dependencies:concept-element-shadow-root
      x-internal="concept-element-shadow-root"}\'s
      [[mode](https://dom.spec.whatwg.org/#shadowroot-mode)]{#concept-shadow-root-mode
      .dfn}
    - A [shadow
      root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#dependencies:concept-element-shadow-root-2
      x-internal="concept-element-shadow-root"}\'s
      [[declarative](https://dom.spec.whatwg.org/#shadowroot-declarative)]{#concept-shadow-root-declarative
      .dfn} member
    - The [[attach a shadow
      root](https://dom.spec.whatwg.org/#concept-attach-a-shadow-root)]{#concept-attach-a-shadow-root
      .dfn} algorithm
    - The [[retargeting
      algorithm](https://dom.spec.whatwg.org/#retarget)]{#dom-retarget
      .dfn}
    - [[`Node`](https://dom.spec.whatwg.org/#interface-node)]{#node
      .dfn} interface
    - [[`NodeList`](https://dom.spec.whatwg.org/#interface-nodelist)]{#nodelist
      .dfn} interface
    - [[`ProcessingInstruction`](https://dom.spec.whatwg.org/#interface-processinginstruction)]{#processinginstruction
      .dfn} interface
    - [[`ShadowRoot`](https://dom.spec.whatwg.org/#interface-shadowroot)]{#shadowroot
      .dfn} interface
    - [[`Text`](https://dom.spec.whatwg.org/#interface-text)]{#text
      .dfn} interface
    - [[`Range`](https://dom.spec.whatwg.org/#interface-range)]{#range
      .dfn} interface
    - [[node
      document](https://dom.spec.whatwg.org/#concept-node-document)]{#node-document
      .dfn} concept
    - [[document
      type](https://dom.spec.whatwg.org/#concept-document-type)]{#concept-document-type
      .dfn} concept
    - [[host](https://dom.spec.whatwg.org/#concept-documentfragment-host)]{#concept-documentfragment-host
      .dfn} concept
    - The [[shadow
      root](https://dom.spec.whatwg.org/#concept-shadow-root)]{#shadow-root
      .dfn} concept, and its [[delegates
      focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus)]{#delegates-focus
      .dfn}, [[available to element
      internals](https://dom.spec.whatwg.org/#shadowroot-available-to-element-internals)]{#available-to-element-internals
      .dfn},
      [[clonable](https://dom.spec.whatwg.org/#shadowroot-clonable)]{#clonable
      .dfn},
      [[serializable](https://dom.spec.whatwg.org/#shadowroot-serializable)]{#shadow-serializable
      .dfn}, [[custom element
      registry](https://dom.spec.whatwg.org/#shadowroot-custom-element-registry)]{#shadow-root-custom-element-registry
      .dfn}, and [[keep custom element registry
      null](https://dom.spec.whatwg.org/#shadowroot-keep-custom-element-registry-null)]{#keep-custom-element-registry-null
      .dfn}.
    - The [[shadow
      host](https://dom.spec.whatwg.org/#element-shadow-host)]{#shadow-host
      .dfn} concept
    - [[`HTMLCollection`](https://dom.spec.whatwg.org/#interface-htmlcollection)]{#htmlcollection
      .dfn} interface, its
      [[`length`](https://dom.spec.whatwg.org/#dom-htmlcollection-length)]{#dom-htmlcollection-length
      .dfn} attribute, and its
      [[`item()`](https://dom.spec.whatwg.org/#dom-htmlcollection-item)]{#dom-htmlcollection-item
      .dfn} and
      [[`namedItem()`](https://dom.spec.whatwg.org/#dom-htmlcollection-nameditem)]{#dom-htmlcollection-nameditem
      .dfn} methods
    - The terms
      [[collection](https://dom.spec.whatwg.org/#concept-collection)]{#concept-collection
      .dfn} and [[represented by the
      collection](https://dom.spec.whatwg.org/#represented-by-the-collection)]{#represented-by-the-collection
      .dfn}
    - [[`DOMTokenList`](https://dom.spec.whatwg.org/#interface-domtokenlist)]{#domtokenlist
      .dfn} interface, and its
      [[`value`](https://dom.spec.whatwg.org/#dom-domtokenlist-value)]{#dom-domtokenlist-value
      .dfn} attribute and
      [[`supports`](https://dom.spec.whatwg.org/#dom-domtokenlist-supports)]{#dom-domtokenlist-supports
      .dfn} operation
    - [[`createDocument()`](https://dom.spec.whatwg.org/#dom-domimplementation-createdocument)]{#dom-domimplementation-createdocument
      .dfn} method
    - [[`createHTMLDocument()`](https://dom.spec.whatwg.org/#dom-domimplementation-createhtmldocument)]{#dom-domimplementation-createhtmldocument
      .dfn} method
    - [[`createElement()`](https://dom.spec.whatwg.org/#dom-document-createelement)]{#dom-document-createelement
      .dfn} method
    - [[`createElementNS()`](https://dom.spec.whatwg.org/#dom-document-createelementns)]{#dom-document-createelementns
      .dfn} method
    - [[`getElementById()`](https://dom.spec.whatwg.org/#dom-nonelementparentnode-getelementbyid)]{#dom-document-getelementbyid
      .dfn} method
    - [[`getElementsByClassName()`](https://dom.spec.whatwg.org/#dom-document-getelementsbyclassname)]{#dom-document-getelementsbyclassname
      .dfn} method
    - [[`append()`](https://dom.spec.whatwg.org/#dom-node-append)]{#dom-node-append
      .dfn} method
    - [[`appendChild()`](https://dom.spec.whatwg.org/#dom-node-appendchild)]{#dom-node-appendchild
      .dfn} method
    - [[`cloneNode()`](https://dom.spec.whatwg.org/#dom-node-clonenode)]{#dom-node-clonenode
      .dfn} method
    - [[`moveBefore()`](https://dom.spec.whatwg.org/#dom-parentnode-movebefore)]{#dom-parentnode-movebefore
      .dfn} method
    - [[`importNode()`](https://dom.spec.whatwg.org/#dom-document-importnode)]{#dom-document-importnode
      .dfn} method
    - [[`preventDefault()`](https://dom.spec.whatwg.org/#dom-event-preventdefault)]{#dom-event-preventdefault
      .dfn} method
    - [[`id`](https://dom.spec.whatwg.org/#dom-element-id)]{#dom-element-id
      .dfn} attribute
    - [[`setAttribute()`](https://dom.spec.whatwg.org/#dom-element-setattribute)]{#dom-element-setattribute
      .dfn} method
    - [[`textContent`](https://dom.spec.whatwg.org/#dom-node-textcontent)]{#textcontent
      .dfn} attribute
    - The [[tree](https://dom.spec.whatwg.org/#concept-tree)]{#tree
      .dfn}, [[shadow
      tree](https://dom.spec.whatwg.org/#concept-shadow-tree)]{#shadow-tree
      .dfn}, and [[node
      tree](https://dom.spec.whatwg.org/#concept-node-tree)]{#node-tree
      .dfn} concepts
    - The [[tree
      order](https://dom.spec.whatwg.org/#concept-tree-order)]{#tree-order
      .dfn} and [[shadow-including tree
      order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order)]{#shadow-including-tree-order
      .dfn} concepts
    - The
      [[element](https://dom.spec.whatwg.org/#concept-element)]{#concept-element
      .dfn} concept
    - The
      [[child](https://dom.spec.whatwg.org/#concept-tree-child)]{#concept-tree-child
      .dfn} concept
    - The [[root](https://dom.spec.whatwg.org/#concept-tree-root)]{#root
      .dfn} and [[shadow-including
      root](https://dom.spec.whatwg.org/#concept-shadow-including-root)]{#shadow-including-root
      .dfn} concepts
    - The [[inclusive
      ancestor](https://dom.spec.whatwg.org/#concept-tree-inclusive-ancestor)]{#inclusive-ancestor
      .dfn},
      [[descendant](https://dom.spec.whatwg.org/#concept-tree-descendant)]{#descendant
      .dfn}, [[shadow-including
      ancestor](https://dom.spec.whatwg.org/#concept-shadow-including-ancestor)]{#concept-shadow-including-ancestor
      .dfn}, [[shadow-including
      descendant](https://dom.spec.whatwg.org/#concept-shadow-including-descendant)]{#shadow-including-descendant
      .dfn}, [[shadow-including inclusive
      descendant](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-descendant)]{#shadow-including-inclusive-descendant
      .dfn}, and [[shadow-including inclusive
      ancestor](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-ancestor)]{#shadow-including-inclusive-ancestor
      .dfn} concepts
    - The [[first
      child](https://dom.spec.whatwg.org/#concept-tree-first-child)]{#first-child
      .dfn}, [[next
      sibling](https://dom.spec.whatwg.org/#concept-tree-next-sibling)]{#next-sibling
      .dfn}, [[previous
      sibling](https://dom.spec.whatwg.org/#concept-tree-previous-sibling)]{#previous-sibling
      .dfn}, and
      [[parent](https://dom.spec.whatwg.org/#concept-tree-parent)]{#parent
      .dfn} concepts
    - The [[parent
      element](https://dom.spec.whatwg.org/#parent-element)]{#parent-element
      .dfn} concept
    - The [[document
      element](https://dom.spec.whatwg.org/#document-element)]{#document-element
      .dfn} concept
    - The [[in a document
      tree](https://dom.spec.whatwg.org/#in-a-document-tree)]{#in-a-document-tree
      .dfn}, [[in a
      document](https://dom.spec.whatwg.org/#in-a-document)]{#in-a-document
      .dfn} (legacy), and
      [[connected](https://dom.spec.whatwg.org/#connected)]{#connected
      .dfn} concepts
    - The
      [[slot](https://dom.spec.whatwg.org/#concept-slot)]{#concept-slot
      .dfn} concept, and its
      [[name](https://dom.spec.whatwg.org/#slot-name)]{#slot-name .dfn}
      and [[assigned
      nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes)]{#assigned-nodes
      .dfn}
    - The [[assigned
      slot](https://dom.spec.whatwg.org/#slotable-assigned-slot)]{#assigned-slot
      .dfn} concept
    - The [[slot
      assignment](https://dom.spec.whatwg.org/#dom-shadowroot-slot-assignment)]{#slot-assignment
      .dfn} concept
    - The
      [[slottable](https://dom.spec.whatwg.org/#concept-slotable)]{#slottable
      .dfn} concept
    - The [[assign slottables for a
      tree](https://dom.spec.whatwg.org/#assign-slotables-for-a-tree)]{#assign-slottables-for-a-tree
      .dfn} algorithm
    - The
      [[`slotchange`](https://dom.spec.whatwg.org/#eventdef-htmlslotelement-slotchange)]{#event-slotchange
      .dfn} event
    - The [[inclusive
      descendant](https://dom.spec.whatwg.org/#concept-tree-inclusive-descendant)]{#inclusive-descendant
      .dfn} concept
    - The [[find flattened
      slottables](https://dom.spec.whatwg.org/#find-flattened-slotables)]{#finding-flattened-slottables
      .dfn} algorithm
    - The [[manual slot
      assignment](https://dom.spec.whatwg.org/#slottable-manual-slot-assignment)]{#manual-slot-assignment
      .dfn} concept
    - The [[assign a
      slot](https://dom.spec.whatwg.org/#assign-a-slot)]{#assign-a-slot
      .dfn} algorithm
    - The
      [[pre-insert](https://dom.spec.whatwg.org/#concept-node-pre-insert)]{#pre-insert
      .dfn},
      [[insert](https://dom.spec.whatwg.org/#concept-node-insert)]{#concept-node-insert
      .dfn},
      [[append](https://dom.spec.whatwg.org/#concept-node-append)]{#concept-node-append
      .dfn},
      [[replace](https://dom.spec.whatwg.org/#concept-node-replace)]{#concept-node-replace
      .dfn}, [[replace
      all](https://dom.spec.whatwg.org/#concept-node-replace-all)]{#concept-node-replace-all
      .dfn}, [[string replace
      all](https://dom.spec.whatwg.org/#string-replace-all)]{#string-replace-all
      .dfn},
      [[remove](https://dom.spec.whatwg.org/#concept-node-remove)]{#concept-node-remove
      .dfn}, and
      [[adopt](https://dom.spec.whatwg.org/#concept-node-adopt)]{#concept-node-adopt
      .dfn} algorithms for nodes
    - The
      [[descendant](https://dom.spec.whatwg.org/#concept-tree-descendant)]{#concept-tree-descendant
      .dfn} concept
    - The [[insertion
      steps](https://dom.spec.whatwg.org/#concept-node-insert-ext)]{#concept-node-insert-ext
      .dfn},
    - The [[post-connection
      steps](https://dom.spec.whatwg.org/#concept-node-post-connection-ext)]{#concept-node-post-connection-ext
      .dfn}, [[removing
      steps](https://dom.spec.whatwg.org/#concept-node-remove-ext)]{#concept-node-remove-ext
      .dfn}, [[moving
      steps](https://dom.spec.whatwg.org/#concept-node-move-ext)]{#concept-node-move-ext
      .dfn}, [[adopting
      steps](https://dom.spec.whatwg.org/#concept-node-adopt-ext)]{#concept-node-adopt-ext
      .dfn}, and [[children changed
      steps](https://dom.spec.whatwg.org/#concept-node-children-changed-ext)]{#children-changed-steps
      .dfn} hooks for elements
    - The
      [[change](https://dom.spec.whatwg.org/#concept-element-attributes-change)]{#concept-element-attributes-change
      .dfn},
      [[append](https://dom.spec.whatwg.org/#concept-element-attributes-append)]{#concept-element-attributes-append
      .dfn},
      [[remove](https://dom.spec.whatwg.org/#concept-element-attributes-remove)]{#concept-element-attributes-remove
      .dfn},
      [[replace](https://dom.spec.whatwg.org/#concept-element-attributes-replace)]{#concept-element-attributes-replace
      .dfn}, [[get an attribute by namespace and local
      name](https://dom.spec.whatwg.org/#concept-element-attributes-get-by-namespace)]{#concept-element-attributes-get-by-namespace
      .dfn}, [[set
      value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value)]{#concept-element-attributes-set-value
      .dfn}, and [[remove an attribute by namespace and local
      name](https://dom.spec.whatwg.org/#concept-element-attributes-remove-by-namespace)]{#concept-element-attributes-remove-by-namespace
      .dfn} algorithms for attributes
    - The [[attribute change
      steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext)]{#concept-element-attributes-change-ext
      .dfn} hook for attributes
    - The
      [[value](https://dom.spec.whatwg.org/#concept-attribute-value)]{#concept-attribute-value
      .dfn} concept for attributes
    - The [[local
      name](https://dom.spec.whatwg.org/#concept-attribute-local-name)]{#concept-attribute-local-name
      .dfn} concept for attributes
    - The [[attribute
      list](https://dom.spec.whatwg.org/#concept-element-attribute)]{#attribute-list
      .dfn} concept
    - The
      [[data](https://dom.spec.whatwg.org/#concept-cd-data)]{#concept-cd-data
      .dfn} of a
      [`CharacterData`{#dependencies:characterdata}](https://dom.spec.whatwg.org/#interface-characterdata){x-internal="characterdata"}
      node and its [[replace
      data](https://dom.spec.whatwg.org/#concept-cd-replace)]{#replace-data
      .dfn} algorithm
    - The [[child text
      content](https://dom.spec.whatwg.org/#concept-child-text-content)]{#child-text-content
      .dfn} of a node
    - The [[descendant text
      content](https://dom.spec.whatwg.org/#concept-descendant-text-content)]{#descendant-text-content
      .dfn} of a node
    - The
      [[name](https://dom.spec.whatwg.org/#concept-doctype-name)]{#concept-doctype-name
      .dfn}, [[public
      ID](https://dom.spec.whatwg.org/#concept-doctype-publicid)]{#concept-doctype-publicid
      .dfn}, and [[system
      ID](https://dom.spec.whatwg.org/#concept-doctype-systemid)]{#concept-doctype-systemid
      .dfn} of a doctype
    - [[`Event`](https://dom.spec.whatwg.org/#interface-event)]{#event
      .dfn} interface
    - [[`Event` and derived interfaces constructor
      behavior](https://dom.spec.whatwg.org/#concept-event-constructor)]{#dom-event-constructor
      .dfn}
    - [[`EventTarget`](https://dom.spec.whatwg.org/#interface-eventtarget)]{#eventtarget
      .dfn} interface
    - The [[activation
      behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior)]{#activation-behaviour
      .dfn} hook
    - The [[legacy-pre-activation
      behavior](https://dom.spec.whatwg.org/#eventtarget-legacy-pre-activation-behavior)]{#legacy-pre-activation-behavior
      .dfn} hook
    - The [[legacy-canceled-activation
      behavior](https://dom.spec.whatwg.org/#eventtarget-legacy-canceled-activation-behavior)]{#legacy-canceled-activation-behavior
      .dfn} hook
    - The [[create an
      event](https://dom.spec.whatwg.org/#concept-event-create)]{#creating-an-event
      .dfn} algorithm
    - The [[fire an
      event](https://dom.spec.whatwg.org/#concept-event-fire)]{#concept-event-fire
      .dfn} algorithm
    - The [[canceled
      flag](https://dom.spec.whatwg.org/#canceled-flag)]{#canceled-flag
      .dfn}
    - The [[dispatch
      flag](https://dom.spec.whatwg.org/#dispatch-flag)]{#dispatch-flag
      .dfn}
    - The
      [[dispatch](https://dom.spec.whatwg.org/#concept-event-dispatch)]{#concept-event-dispatch
      .dfn} algorithm
    - [[`EventInit`](https://dom.spec.whatwg.org/#dictdef-eventinit)]{#eventinit
      .dfn} dictionary type
    - [[`type`](https://dom.spec.whatwg.org/#dom-event-type)]{#dom-event-type
      .dfn} attribute
    - An event\'s
      [[target](https://dom.spec.whatwg.org/#concept-event-target)]{#concept-event-target
      .dfn}
    - [[`currentTarget`](https://dom.spec.whatwg.org/#dom-event-currenttarget)]{#dom-event-currenttarget
      .dfn} attribute
    - [[`bubbles`](https://dom.spec.whatwg.org/#dom-event-bubbles)]{#dom-event-bubbles
      .dfn} attribute
    - [[`cancelable`](https://dom.spec.whatwg.org/#dom-event-cancelable)]{#dom-event-cancelable
      .dfn} attribute
    - [[`composed`](https://dom.spec.whatwg.org/#dom-event-composed)]{#dom-event-composed
      .dfn} attribute
    - [[composed
      flag](https://dom.spec.whatwg.org/#composed-flag)]{#composed-flag
      .dfn}
    - [[`isTrusted`](https://dom.spec.whatwg.org/#dom-event-istrusted)]{#dom-event-istrusted
      .dfn} attribute
    - [[`initEvent()`](https://dom.spec.whatwg.org/#dom-event-initevent)]{#dom-event-initevent
      .dfn} method
    - [[add an event
      listener](https://dom.spec.whatwg.org/#add-an-event-listener)]{#add-an-event-listener
      .dfn}
    - [[`addEventListener()`](https://dom.spec.whatwg.org/#dom-eventtarget-addeventlistener)]{#dom-eventtarget-addeventlistener
      .dfn} method
    - The [[remove an event
      listener](https://dom.spec.whatwg.org/#remove-an-event-listener)]{#remove-an-event-listener
      .dfn} and [[remove all event
      listeners](https://dom.spec.whatwg.org/#remove-all-event-listeners)]{#remove-all-event-listeners
      .dfn} algorithms
    - [[`EventListener`](https://dom.spec.whatwg.org/#callbackdef-eventlistener)]{#dom-eventlistener
      .dfn} callback interface
    - The
      [[type](https://dom.spec.whatwg.org/#dom-event-type)]{#concept-event-type
      .dfn} of an event
    - An [[event
      listener](https://dom.spec.whatwg.org/#concept-event-listener)]{#event-listener
      .dfn} and its
      [[type](https://dom.spec.whatwg.org/#event-listener-type)]{#event-listener-type
      .dfn} and
      [[callback](https://dom.spec.whatwg.org/#event-listener-callback)]{#event-listener-callback
      .dfn}
    - The
      [[encoding](https://dom.spec.whatwg.org/#concept-document-encoding)]{#document's-character-encoding
      .dfn} (herein the *character encoding*),
      [[mode](https://dom.spec.whatwg.org/#concept-document-mode)]{#concept-document-mode
      .dfn}, [[custom element
      registry](https://dom.spec.whatwg.org/#document-custom-element-registry)]{#document-custom-element-registry
      .dfn}, [[allow declarative shadow
      roots](https://dom.spec.whatwg.org/#concept-document-allow-declarative-shadow-roots)]{#concept-document-allow-declarative-shadow-roots
      .dfn}, and [[content
      type](https://dom.spec.whatwg.org/#concept-document-content-type)]{#concept-document-content-type
      .dfn} of a [`Document`{#dependencies:document}](dom.html#document)
    - The distinction between [[XML
      documents](https://dom.spec.whatwg.org/#xml-document)]{#xml-documents
      .dfn} and [[HTML
      documents](https://dom.spec.whatwg.org/#html-document)]{#html-documents
      .dfn}
    - The terms [[quirks
      mode](https://dom.spec.whatwg.org/#concept-document-quirks)]{#quirks-mode
      .dfn}, [[limited-quirks
      mode](https://dom.spec.whatwg.org/#concept-document-limited-quirks)]{#limited-quirks-mode
      .dfn}, and [[no-quirks
      mode](https://dom.spec.whatwg.org/#concept-document-no-quirks)]{#no-quirks-mode
      .dfn}
    - The algorithm [[clone a
      node](https://dom.spec.whatwg.org/#concept-node-clone)]{#concept-node-clone
      .dfn} with its arguments
      [[`document`{.variable}](https://dom.spec.whatwg.org/#clone-a-node-document)]{#concept-node-clone-document
      .dfn},
      [[`subtree`{.variable}](https://dom.spec.whatwg.org/#clone-a-node-subtree)]{#concept-node-clone-subtree
      .dfn},
      [[`parent`{.variable}](https://dom.spec.whatwg.org/#clone-a-node-parent)]{#concept-node-clone-parent
      .dfn}, and
      [[`fallbackRegistry`{.variable}](https://dom.spec.whatwg.org/#clone-a-node-fallbackregistry)]{#concept-node-clone-fallbackregistry
      .dfn}, and the concept of [[cloning
      steps](https://dom.spec.whatwg.org/#concept-node-clone-ext)]{#concept-node-clone-ext
      .dfn}
    - The concept of [base URL change steps]{#base-url-change-steps
      .dfn} and the definition of what happens when an element is
      [affected by a base URL change]{#affected-by-a-base-url-change
      .dfn}
    - The concept of an element\'s [[unique identifier
      (ID)](https://dom.spec.whatwg.org/#concept-id)]{#concept-id .dfn}
    - The concept of an element\'s
      [[classes](https://dom.spec.whatwg.org/#concept-class)]{#concept-class
      .dfn}
    - The term [[supported
      tokens](https://dom.spec.whatwg.org/#concept-supported-tokens)]{#concept-supported-tokens
      .dfn}
    - The concept of a DOM
      [[range](https://dom.spec.whatwg.org/#concept-range)]{#concept-range
      .dfn}, and the terms [[start
      node](https://dom.spec.whatwg.org/#concept-range-start-node)]{#concept-range-start-node
      .dfn},
      [[start](https://dom.spec.whatwg.org/#concept-range-start)]{#concept-range-start
      .dfn},
      [[end](https://dom.spec.whatwg.org/#concept-range-end)]{#concept-range-end
      .dfn}, and [[boundary
      point](https://dom.spec.whatwg.org/#concept-range-bp)]{#concept-range-bp
      .dfn} as applied to ranges.
    - The [[create an
      element](https://dom.spec.whatwg.org/#concept-create-element)]{#create-an-element
      .dfn} algorithm
    - The [[element
      interface](https://dom.spec.whatwg.org/#concept-element-interface)]{#element-interface
      .dfn} concept
    - The concepts of [[custom element
      state](https://dom.spec.whatwg.org/#concept-element-custom-element-state)]{#custom-element-state
      .dfn}, and of
      [[defined](https://dom.spec.whatwg.org/#concept-element-defined)]{#concept-element-defined
      .dfn} and
      [[custom](https://dom.spec.whatwg.org/#concept-element-custom)]{#concept-element-custom
      .dfn} elements
    - An element\'s
      [[namespace](https://dom.spec.whatwg.org/#concept-element-namespace)]{#concept-element-namespace
      .dfn}, [[namespace
      prefix](https://dom.spec.whatwg.org/#concept-element-namespace-prefix)]{#concept-element-namespace-prefix
      .dfn}, [[local
      name](https://dom.spec.whatwg.org/#concept-element-local-name)]{#concept-element-local-name
      .dfn}, [[custom element
      registry](https://dom.spec.whatwg.org/#element-custom-element-registry)]{#element-custom-element-registry
      .dfn}, [[custom element
      definition](https://dom.spec.whatwg.org/#concept-element-custom-element-definition)]{#concept-element-custom-element-definition
      .dfn}, and [[`is`
      value](https://dom.spec.whatwg.org/#concept-element-is-value)]{#concept-element-is-value
      .dfn}
    - [[`MutationObserver`](https://dom.spec.whatwg.org/#mutationobserver)]{#mutationobserver
      .dfn} interface and [[mutation
      observers](https://dom.spec.whatwg.org/#mutation-observers)]{#mutation-observers
      .dfn} in general
    - [[`AbortController`](https://dom.spec.whatwg.org/#abortcontroller)]{#abortcontroller
      .dfn} and its
      [[signal](https://dom.spec.whatwg.org/#abortcontroller-signal)]{#concept-abortcontroller-signal
      .dfn}
    - [[`AbortSignal`](https://dom.spec.whatwg.org/#abortsignal)]{#abortsignal
      .dfn}
    - [[aborted](https://dom.spec.whatwg.org/#abortsignal-aborted)]{#abortsignal-aborted
      .dfn}
    - [[signal
      abort](https://dom.spec.whatwg.org/#abortcontroller-signal-abort)]{#signal-abort
      .dfn}
    - [[add](https://dom.spec.whatwg.org/#abortsignal-add)]{#abortsignal-add
      .dfn}
    - The [[get an attribute by
      name](https://dom.spec.whatwg.org/#concept-element-attributes-get-by-name)]{#get-an-attribute-by-name
      .dfn} algorithm

    The following features are defined in UI Events:
    [\[UIEVENTS\]](references.html#refsUIEVENTS)

    - The
      [[`MouseEvent`](https://w3c.github.io/uievents/#mouseevent)]{#mouseevent
      .dfn} interface
    - The
      [`MouseEvent`{#dependencies:mouseevent}](https://w3c.github.io/uievents/#mouseevent){x-internal="mouseevent"}
      interface\'s
      [[`relatedTarget`](https://w3c.github.io/uievents/#dom-mouseevent-relatedtarget)]{#dom-mouseevent-relatedtarget
      .dfn} attribute
    - [[`MouseEventInit`](https://w3c.github.io/uievents/#dictdef-mouseeventinit)]{#mouseeventinit
      .dfn} dictionary type
    - The
      [[`FocusEvent`](https://w3c.github.io/uievents/#focusevent)]{#focusevent
      .dfn} interface
    - The
      [`FocusEvent`{#dependencies:focusevent}](https://w3c.github.io/uievents/#focusevent){x-internal="focusevent"}
      interface\'s
      [[`relatedTarget`](https://w3c.github.io/uievents/#dom-focusevent-relatedtarget)]{#dom-focusevent-relatedtarget
      .dfn} attribute
    - The
      [[`UIEvent`](https://w3c.github.io/uievents/#uievent)]{#uievent
      .dfn} interface
    - The
      [`UIEvent`{#dependencies:uievent}](https://w3c.github.io/uievents/#uievent){x-internal="uievent"}
      interface\'s
      [[`view`](https://w3c.github.io/uievents/#dom-uievent-view)]{#dom-uievent-view
      .dfn} attribute
    - [[`auxclick`](https://w3c.github.io/uievents/#event-type-auxclick)]{#event-auxclick
      .dfn} event
    - [[`beforeinput`](https://w3c.github.io/uievents/#event-type-beforeinput)]{#event-beforeinput
      .dfn} event
    - [[`click`](https://w3c.github.io/uievents/#event-type-click)]{#event-click
      .dfn} event
    - [[`contextmenu`](https://w3c.github.io/uievents/#event-type-contextmenu)]{#event-contextmenu
      .dfn} event
    - [[`dblclick`](https://w3c.github.io/uievents/#event-type-dblclick)]{#event-dblclick
      .dfn} event
    - [[`input`](https://w3c.github.io/uievents/#event-type-input)]{#event-input
      .dfn} event
    - [[`mousedown`](https://w3c.github.io/uievents/#event-type-mousedown)]{#event-mousedown
      .dfn} event
    - [[`mouseenter`](https://w3c.github.io/uievents/#event-type-mouseenter)]{#event-mouseenter
      .dfn} event
    - [[`mouseleave`](https://w3c.github.io/uievents/#event-type-mouseleave)]{#event-mouseleave
      .dfn} event
    - [[`mousemove`](https://w3c.github.io/uievents/#event-type-mousemove)]{#event-mousemove
      .dfn} event
    - [[`mouseout`](https://w3c.github.io/uievents/#event-type-mouseout)]{#event-mouseout
      .dfn} event
    - [[`mouseover`](https://w3c.github.io/uievents/#event-type-mouseover)]{#event-mouseover
      .dfn} event
    - [[`mouseup`](https://w3c.github.io/uievents/#event-type-mouseup)]{#event-mouseup
      .dfn} event
    - [[`wheel`](https://w3c.github.io/uievents/#event-type-wheel)]{#event-wheel
      .dfn} event
    - [[`keydown`](https://w3c.github.io/uievents/#event-type-keydown)]{#event-keydown
      .dfn} event
    - [[`keypress`](https://w3c.github.io/uievents/#event-type-keypress)]{#event-keypress
      .dfn} event
    - [[`keyup`](https://w3c.github.io/uievents/#event-type-keyup)]{#event-keyup
      .dfn} event

    The following features are defined in Touch Events:
    [\[TOUCH\]](references.html#refsTOUCH)

    - [[`Touch`](https://w3c.github.io/touch-events/#touch-interface)]{#touch
      .dfn} interface
    - [[Touch
      point](https://w3c.github.io/touch-events/#dfn-touch-point)]{#touch-point
      .dfn} concept
    - [[`touchend`](https://w3c.github.io/touch-events/#event-touchend)]{#event-touchend
      .dfn} event

    The following features are defined in Pointer Events:
    [\[POINTEREVENTS\]](references.html#refsPOINTEREVENTS)

    - The
      [[`PointerEvent`](https://w3c.github.io/pointerevents/#pointerevent-interface)]{#pointerevent
      .dfn} interface
    - The
      [`PointerEvent`{#dependencies:pointerevent}](https://w3c.github.io/pointerevents/#pointerevent-interface){x-internal="pointerevent"}
      interface\'s
      [[`pointerType`](https://w3c.github.io/pointerevents/#dom-pointerevent-pointertype)]{#pointertype
      .dfn} attribute
    - [[fire a pointer
      event](https://w3c.github.io/pointerevents/#dfn-fire-a-pointer-event)]{#fire-a-pointer-event
      .dfn}
    - [[`pointerdown`](https://w3c.github.io/pointerevents/#the-pointerdown-event)]{#event-pointerdown
      .dfn} event
    - [[`pointerup`](https://w3c.github.io/pointerevents/#the-pointerup-event)]{#event-pointerup
      .dfn} event
    - [[`pointercancel`](https://w3c.github.io/pointerevents/#the-pointercancel-event)]{#event-pointercancel
      .dfn} event

    The following events are defined in Clipboard API and events:
    [\[CLIPBOARD-APIS\]](references.html#refsCLIPBOARD-APIS)

    - [[`copy`](https://w3c.github.io/clipboard-apis/#clipboard-event-copy)]{#event-copy
      .dfn} event
    - [[`cut`](https://w3c.github.io/clipboard-apis/#clipboard-event-cut)]{#event-cut
      .dfn} event
    - [[`paste`](https://w3c.github.io/clipboard-apis/#clipboard-event-paste)]{#event-paste
      .dfn} event

    This specification sometimes uses the term [name]{.dfn} to refer to
    the event\'s
    [type](https://dom.spec.whatwg.org/#dom-event-type){#dependencies:concept-event-type
    x-internal="concept-event-type"}; as in, \"an event named `click`\"
    or \"if the event name is `keypress`\". The terms \"name\" and
    \"type\" for events are synonymous.

    The following features are defined in DOM Parsing and Serialization:
    [\[DOMPARSING\]](references.html#refsDOMPARSING)

    - [[`XML serialization`](https://w3c.github.io/DOM-Parsing/#dfn-xml-serialization)]{#xml-serialization
      .dfn}

    The following features are defined in Selection API:
    [\[SELECTION\]](references.html#refsSELECTION)

    - [[selection](https://w3c.github.io/selection-api/#dfn-selection)]{#document-selection
      .dfn}
    - [[`Selection`](https://w3c.github.io/selection-api/#selection-interface)]{#selection
      .dfn}

    User agents are encouraged to implement the features described in
    execCommand. [\[EXECCOMMAND\]](references.html#refsEXECCOMMAND)

    The following features are defined in Fullscreen API:
    [\[FULLSCREEN\]](references.html#refsFULLSCREEN)

    - [[`requestFullscreen()`](https://fullscreen.spec.whatwg.org/#dom-element-requestfullscreen)]{#dom-element-requestfullscreen
      .dfn}
    - [[`fullscreenchange`](https://fullscreen.spec.whatwg.org/#eventdef-document-fullscreenchange)]{#event-fullscreenchange
      .dfn}
    - [[run the fullscreen
      steps](https://fullscreen.spec.whatwg.org/#run-the-fullscreen-steps)]{#run-the-fullscreen-steps
      .dfn}
    - [[fully exit
      fullscreen](https://fullscreen.spec.whatwg.org/#fully-exit-fullscreen)]{#fully-exit-fullscreen
      .dfn}
    - [[fullscreen
      element](https://fullscreen.spec.whatwg.org/#fullscreen-element)]{#fullscreen-element
      .dfn}
    - [[fullscreen
      flag](https://fullscreen.spec.whatwg.org/#fullscreen-flag)]{#fullscreen-flag
      .dfn}

    High Resolution Time provides the following features:
    [\[HRT\]](references.html#refsHRT)

    - [[current high resolution
      time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time)]{#current-high-resolution-time
      .dfn}
    - [[relative high resolution
      time](https://w3c.github.io/hr-time/#dfn-relative-high-resolution-time)]{#relative-high-resolution-time
      .dfn}
    - [[unsafe shared current
      time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time)]{#unsafe-shared-current-time
      .dfn}
    - [[shared monotonic
      clock](https://w3c.github.io/hr-time/#dfn-shared-monotonic-clock)]{#shared-monotonic-clock
      .dfn}
    - [[unsafe
      moment](https://w3c.github.io/hr-time/#dfn-unsafe-moment)]{#unsafe-moment
      .dfn}
    - [[duration
      from](https://w3c.github.io/hr-time/#dfn-duration-from)]{#duration-from
      .dfn}
    - [[coarsen
      time](https://w3c.github.io/hr-time/#dfn-coarsen-time)]{#coarsen-time
      .dfn}
    - [[current wall
      time](https://w3c.github.io/hr-time/#dfn-current-wall-time)]{#current-wall-time
      .dfn}
    - [[Unix
      epoch](https://w3c.github.io/hr-time/#dfn-unix-epoch)]{#unix-epoch
      .dfn}
    - [[`DOMHighResTimeStamp`](https://w3c.github.io/hr-time/#dom-domhighrestimestamp)]{#domhighrestimestamp
      .dfn}

File API

:   This specification uses the following features defined in File API:
    [\[FILEAPI\]](references.html#refsFILEAPI)

    - The [[`Blob`](https://w3c.github.io/FileAPI/#dfn-Blob)]{#blob
      .dfn} interface and its
      [[`type`](https://w3c.github.io/FileAPI/#dfn-type)]{#dom-blob-type
      .dfn} attribute
    - The [[`File`](https://w3c.github.io/FileAPI/#dfn-file)]{#file
      .dfn} interface and its
      [[`name`](https://w3c.github.io/FileAPI/#dfn-name)]{#dom-file-name
      .dfn} and
      [[`lastModified`](https://w3c.github.io/FileAPI/#dfn-lastModified)]{#dom-file-lastmodified
      .dfn} attributes
    - The
      [[`FileList`](https://w3c.github.io/FileAPI/#filelist-section)]{#filelist
      .dfn} interface
    - The concept of a
      [`Blob`{#dependencies:blob}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}\'s
      [[snapshot
      state](https://w3c.github.io/FileAPI/#snapshot-state)]{#snapshot-state
      .dfn}
    - The concept of [read errors]{#file-error-read .dfn}
    - [[Blob URL
      Store](https://w3c.github.io/FileAPI/#BlobURLStore)]{#blob-url-store
      .dfn}
    - [[blob URL
      entry](https://w3c.github.io/FileAPI/#blob-url-entry)]{#blob-url-entry
      .dfn} and its
      [[environment](https://w3c.github.io/FileAPI/#blob-url-entry-environment)]{#blob-url-entry-environment
      .dfn}
    - The [[obtain a blob
      object](https://w3c.github.io/FileAPI/#blob-url-obtain-object)]{#blob-url-obtain-object
      .dfn} algorithm

Indexed Database API

:   The following terms are defined in Indexed Database API:
    [\[INDEXEDDB\]](references.html#refsINDEXEDDB)

    - [[cleanup Indexed Database
      transactions](https://w3c.github.io/IndexedDB/#cleanup-indexed-database-transactions)]{#cleanup-indexed-database-transactions
      .dfn}
    - [[`IDBVersionChangeEvent`](https://w3c.github.io/IndexedDB/#idbversionchangeevent)]{#idbversionchangeevent
      .dfn}

Media Source Extensions

:   The following terms are defined in Media Source Extensions:
    [\[MEDIASOURCE\]](references.html#refsMEDIASOURCE)

    - [[`MediaSource`](https://w3c.github.io/media-source/#idl-def-mediasource)]{#mediasource
      .dfn} interface
    - [[detaching from a media
      element](https://w3c.github.io/media-source/#mediasource-detach)]{#detaching-from-a-media-element
      .dfn}

Media Capture and Streams

:   The following terms are defined in Media Capture and Streams:
    [\[MEDIASTREAM\]](references.html#refsMEDIASTREAM)

    - [[`MediaStream`](https://w3c.github.io/mediacapture-main/getusermedia.html#idl-def-mediastream)]{#mediastream
      .dfn} interface
    - [[`MediaStreamTrack`](https://w3c.github.io/mediacapture-main/getusermedia.html#mediastreamtrack)]{#mediastreamtrack
      .dfn}
    - [[live
      state](https://w3c.github.io/mediacapture-main/getusermedia.html#idl-def-MediaStreamTrackState.live)]{#live-state
      .dfn}
    - [[`getUserMedia()`](https://w3c.github.io/mediacapture-main/getusermedia.html#dom-mediadevices-getusermedia)]{#getusermedia()
      .dfn}

Reporting

:   The following terms are defined in Reporting:
    [\[REPORTING\]](references.html#refsREPORTING)

    - [[Queue a
      report](https://w3c.github.io/reporting/#queue-report)]{#queue-a-report
      .dfn}
    - [[report
      type](https://w3c.github.io/reporting/#report-type)]{#report-type
      .dfn}
    - [[visible to
      `ReportingObserver`s](https://w3c.github.io/reporting/#visible-to-reportingobservers)]{#visible-to-reportingobservers
      .dfn}

XMLHttpRequest

:   The following features and terms are defined in XMLHttpRequest:
    [\[XHR\]](references.html#refsXHR)

    - The
      [[`XMLHttpRequest`](https://xhr.spec.whatwg.org/#xmlhttprequest)]{#xmlhttprequest
      .dfn} interface, and its
      [[`responseXML`](https://xhr.spec.whatwg.org/#dom-xmlhttprequest-responsexml)]{#dom-xmlhttprequest-responsexml
      .dfn} attribute
    - The
      [[`ProgressEvent`](https://xhr.spec.whatwg.org/#interface-progressevent)]{#progressevent
      .dfn} interface, and its
      [[`lengthComputable`](https://xhr.spec.whatwg.org/#dom-progressevent-lengthcomputable)]{#dom-progressevent-lengthcomputable
      .dfn},
      [[`loaded`](https://xhr.spec.whatwg.org/#dom-progressevent-loaded)]{#dom-progressevent-loaded
      .dfn}, and
      [[`total`](https://xhr.spec.whatwg.org/#dom-progressevent-total)]{#dom-progressevent-total
      .dfn} attributes
    - The
      [[`FormData`](https://xhr.spec.whatwg.org/#formdata)]{#formdata
      .dfn} interface, and its associated [[entry
      list](https://xhr.spec.whatwg.org/#concept-formdata-entry-list)]{#formdata-entry-list
      .dfn}

Battery Status

:   The following features are defined in Battery Status API:
    [\[BATTERY\]](references.html#refsBATTERY)

    - [[`getBattery()`](https://w3c.github.io/battery/#widl-Navigator-getBattery-Promise-BatteryManager)]{#dom-navigator-getbattery
      .dfn} method

Media Queries

:   Implementations must support Media Queries. The
    [[\<media-condition\>](https://drafts.csswg.org/mediaqueries/#typedef-media-condition)]{#media-condition
    .dfn} feature is defined therein. [\[MQ\]](references.html#refsMQ)

CSS modules

:   While support for CSS as a whole is not required of implementations
    of this specification (though it is encouraged, at least for web
    browsers), some features are defined in terms of specific CSS
    requirements.

    When this specification requires that something be [[parsed
    according to a particular CSS
    grammar](https://drafts.csswg.org/css-syntax/#parse-grammar)]{#parse-something-according-to-a-css-grammar
    .dfn}, the relevant algorithm in CSS Syntax must be followed,
    including error handling rules.
    [\[CSSSYNTAX\]](references.html#refsCSSSYNTAX)

    For example, user agents are required to close all open constructs
    upon finding the end of a style sheet unexpectedly. Thus, when
    parsing the string \"`rgb(0,0,0`\" (with a missing
    close-parenthesis) for a color value, the close parenthesis is
    implied by this error handling rule, and a value is obtained (the
    color \'black\'). However, the similar construct \"`rgb(0,0,`\"
    (with both a missing parenthesis and a missing \"blue\" value)
    cannot be parsed, as closing the open construct does not result in a
    viable value.

    The following terms and features are defined in Cascading Style
    Sheets (CSS): [\[CSS\]](references.html#refsCSS)

    - [[viewport](https://drafts.csswg.org/css2/#viewport)]{#viewport
      .dfn}
    - [[line box](https://drafts.csswg.org/css2/#line-box)]{#line-box
      .dfn}
    - [[out-of-flow](https://drafts.csswg.org/css2/#out-of-flow)]{#out-of-flow
      .dfn}
    - [[in-flow](https://drafts.csswg.org/css2/#in-flow)]{#in-flow .dfn}
    - [[collapsing
      margins](https://drafts.csswg.org/css2/#collapsing-margins)]{#collapsing-margins
      .dfn}
    - [[containing
      block](https://drafts.csswg.org/css2/#containing-block-details)]{#containing-block
      .dfn}
    - [[inline
      box](https://drafts.csswg.org/css2/#inline-box)]{#inline-box .dfn}
    - [[block
      box](https://drafts.csswg.org/css2/#block-boxes%E2%91%A0)]{#block-box
      .dfn}
    - The [[\'top\'](https://drafts.csswg.org/css2/#propdef-top)]{#'top'
      .dfn},
      [[\'bottom\'](https://drafts.csswg.org/css2/#propdef-bottom)]{#'bottom'
      .dfn},
      [[\'left\'](https://drafts.csswg.org/css2/#propdef-left)]{#'left'
      .dfn}, and
      [[\'right\'](https://drafts.csswg.org/css2/#propdef-right)]{#'right'
      .dfn} properties
    - The
      [[\'float\'](https://drafts.csswg.org/css2/#float-position)]{#'float'
      .dfn} property
    - The
      [[\'clear\'](https://drafts.csswg.org/css2/#flow-control)]{#'clear'
      .dfn} property
    - The
      [[\'width\'](https://drafts.csswg.org/css2/#the-width-property)]{#'width'
      .dfn} property
    - The
      [[\'height\'](https://drafts.csswg.org/css2/#the-height-property)]{#'height'
      .dfn} property
    - The
      [[\'min-width\'](https://drafts.csswg.org/css2/#min-max-widths)]{#'min-width'
      .dfn} property
    - The
      [[\'min-height\'](https://drafts.csswg.org/css2/#min-max-heights)]{#'min-height'
      .dfn} property
    - The
      [[\'max-width\'](https://drafts.csswg.org/css2/#min-max-widths)]{#'max-width'
      .dfn} property
    - The
      [[\'max-height\'](https://drafts.csswg.org/css2/#min-max-heights)]{#'max-height'
      .dfn} property
    - The
      [[\'line-height\'](https://drafts.csswg.org/css2/#propdef-line-height)]{#'line-height'
      .dfn} property
    - The
      [[\'vertical-align\'](https://drafts.csswg.org/css2/#propdef-vertical-align)]{#'vertical-align'
      .dfn} property
    - The
      [[\'content\'](https://drafts.csswg.org/css2/#content%E2%91%A0)]{#'content'
      .dfn} property
    - The
      [[\'inline-block\'](https://drafts.csswg.org/css2/#value-def-inline-block)]{#'inline-block'
      .dfn} value of the
      [\'display\'](https://drafts.csswg.org/css2/#display-prop){#dependencies:'display'
      x-internal="'display'"} property
    - The
      [[\'visibility\'](https://drafts.csswg.org/css2/#propdef-visibility)]{#'visibility'
      .dfn} property

    The basic version of the
    [[\'display\'](https://drafts.csswg.org/css2/#display-prop)]{#'display'
    .dfn} property is defined in CSS, and the property is extended by
    other CSS modules. [\[CSS\]](references.html#refsCSS)
    [\[CSSRUBY\]](references.html#refsCSSRUBY)
    [\[CSSTABLE\]](references.html#refsCSSTABLE)

    The following terms and features are defined in CSS Box Model:
    [\[CSSBOX\]](references.html#refsCSSBOX)

    - [[content
      area](https://drafts.csswg.org/css-box/#content-area)]{#content-area
      .dfn}
    - [[content
      box](https://drafts.csswg.org/css-box/#content-box)]{#content-box
      .dfn}
    - [[border
      box](https://drafts.csswg.org/css-box/#border-box)]{#border-box
      .dfn}
    - [[margin
      box](https://drafts.csswg.org/css-box/#margin-box)]{#margin-box
      .dfn}
    - [[border
      edge](https://drafts.csswg.org/css-box/#border-edge)]{#border-edge
      .dfn}
    - [[margin
      edge](https://drafts.csswg.org/css-box/#margin-edge)]{#margin-edge
      .dfn}
    - The
      [[\'margin-top\'](https://drafts.csswg.org/css-box/#propdef-margin-top)]{#'margin-top'
      .dfn},
      [[\'margin-bottom\'](https://drafts.csswg.org/css-box/#propdef-margin-bottom)]{#'margin-bottom'
      .dfn},
      [[\'margin-left\'](https://drafts.csswg.org/css-box/#propdef-margin-left)]{#'margin-left'
      .dfn}, and
      [[\'margin-right\'](https://drafts.csswg.org/css-box/#propdef-margin-right)]{#'margin-right'
      .dfn} properties
    - The
      [[\'padding-top\'](https://drafts.csswg.org/css-box/#propdef-padding-top)]{#'padding-top'
      .dfn},
      [[\'padding-bottom\'](https://drafts.csswg.org/css-box/#propdef-padding-bottom)]{#'padding-bottom'
      .dfn},
      [[\'padding-left\'](https://drafts.csswg.org/css-box/#propdef-padding-left)]{#'padding-left'
      .dfn}, and
      [[\'padding-right\'](https://drafts.csswg.org/css-box/#propdef-padding-right)]{#'padding-right'
      .dfn} properties

    The following features are defined in CSS Logical Properties:
    [\[CSSLOGICAL\]](references.html#refsCSSLOGICAL)

    - The
      [[\'margin-block\'](https://drafts.csswg.org/css-logical/#propdef-margin-block)]{#'margin-block'
      .dfn},
      [[\'margin-block-start\'](https://drafts.csswg.org/css-logical/#propdef-margin-block-start)]{#'margin-block-start'
      .dfn},
      [[\'margin-block-end\'](https://drafts.csswg.org/css-logical/#propdef-margin-block-end)]{#'margin-block-end'
      .dfn},
      [[\'margin-inline\'](https://drafts.csswg.org/css-logical/#propdef-margin-inline)]{#'margin-inline'
      .dfn},
      [[\'margin-inline-start\'](https://drafts.csswg.org/css-logical/#propdef-margin-inline-start)]{#'margin-inline-start'
      .dfn}, and
      [[\'margin-inline-end\'](https://drafts.csswg.org/css-logical/#propdef-margin-inline-end)]{#'margin-inline-end'
      .dfn} properties
    - The
      [[\'padding-block\'](https://drafts.csswg.org/css-logical/#propdef-padding-block)]{#'padding-block'
      .dfn},
      [[\'padding-block-start\'](https://drafts.csswg.org/css-logical/#propdef-padding-block-start)]{#'padding-block-start'
      .dfn},
      [[\'padding-block-end\'](https://drafts.csswg.org/css-logical/#propdef-padding-block-end)]{#'padding-block-end'
      .dfn},
      [[\'padding-inline\'](https://drafts.csswg.org/css-logical/#propdef-padding-inline)]{#'padding-inline'
      .dfn},
      [[\'padding-inline-start\'](https://drafts.csswg.org/css-logical/#propdef-padding-inline-start)]{#'padding-inline-start'
      .dfn}, and
      [[\'padding-inline-end\'](https://drafts.csswg.org/css-logical/#propdef-padding-inline-end)]{#'padding-inline-end'
      .dfn} properties
    - The
      [[\'border-block-width\'](https://drafts.csswg.org/css-logical/#propdef-border-block-width)]{#'border-block-width'
      .dfn},
      [[\'border-block-start-width\'](https://drafts.csswg.org/css-logical/#propdef-border-block-start-width)]{#'border-block-start-width'
      .dfn},
      [[\'border-block-end-width\'](https://drafts.csswg.org/css-logical/#propdef-border-block-end-width)]{#'border-block-end-width'
      .dfn},
      [[\'border-inline-width\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-width)]{#'border-inline-width'
      .dfn},
      [[\'border-inline-start-width\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-start-width)]{#'border-inline-start-width'
      .dfn},
      [[\'border-inline-end-width\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-end-width)]{#'border-inline-end-width'
      .dfn},
      [[\'border-block-style\'](https://drafts.csswg.org/css-logical/#propdef-border-block-style)]{#'border-block-style'
      .dfn},
      [[\'border-block-start-style\'](https://drafts.csswg.org/css-logical/#propdef-border-block-start-style)]{#'border-block-start-style'
      .dfn},
      [[\'border-block-end-style\'](https://drafts.csswg.org/css-logical/#propdef-border-block-end-style)]{#'border-block-end-style'
      .dfn},
      [[\'border-inline-style\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-style)]{#'border-inline-style'
      .dfn},
      [[\'border-inline-start-style\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-start-style)]{#'border-inline-start-style'
      .dfn},
      [[\'border-inline-end-style\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-end-style)]{#'border-inline-end-style'
      .dfn},
      [[\'border-block-start-color\'](https://drafts.csswg.org/css-logical/#propdef-border-block-start-color)]{#'border-block-start-color'
      .dfn},
      [[\'border-block-end-color\'](https://drafts.csswg.org/css-logical/#propdef-border-block-end-color)]{#'border-block-end-color'
      .dfn},
      [[\'border-inline-start-color\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-start-color)]{#'border-inline-start-color'
      .dfn},
      [[\'border-inline-end-color\'](https://drafts.csswg.org/css-logical/#propdef-border-inline-end-color)]{#'border-inline-end-color'
      .dfn},
      [[\'border-start-start-radius\'](https://drafts.csswg.org/css-logical/#propdef-border-start-start-radius)]{#'border-start-start-radius'
      .dfn},
      [[\'border-start-end-radius\'](https://drafts.csswg.org/css-logical/#propdef-border-start-end-radius)]{#'border-start-end-radius'
      .dfn},
      [[\'border-end-start-radius\'](https://drafts.csswg.org/css-logical/#propdef-border-end-start-radius)]{#'border-end-start-radius'
      .dfn}, and
      [[\'border-end-end-radius\'](https://drafts.csswg.org/css-logical/#propdef-border-end-end-radius)]{#'border-end-end-radius'
      .dfn} properties
    - The
      [[\'block-size\'](https://drafts.csswg.org/css-logical/#propdef-block-size)]{#'block-size'
      .dfn} property
    - The
      [[\'inline-size\'](https://drafts.csswg.org/css-logical/#propdef-inline-size)]{#'inline-size'
      .dfn} property
    - The
      [[\'inset-block-start\'](https://drafts.csswg.org/css-logical/#propdef-inset-block-start)]{#'inset-block-start'
      .dfn} property
    - The
      [[\'inset-block-end\'](https://drafts.csswg.org/css-logical/#propdef-inset-block-end)]{#'inset-block-end'
      .dfn} property

    The following terms and features are defined in CSS Color:
    [\[CSSCOLOR\]](references.html#refsCSSCOLOR)

    - [[named
      color](https://drafts.csswg.org/css-color/#named-color)]{#named-color
      .dfn}
    - [[\<color\>](https://drafts.csswg.org/css-color/#typedef-color)]{#color
      .dfn}
    - The
      [[\'color\'](https://drafts.csswg.org/css-color/#the-color-property)]{#'color'
      .dfn} property
    - The
      [[\'currentcolor\'](https://drafts.csswg.org/css-color/#valdef-color-currentcolor)]{#currentcolor
      .dfn} value
    - [[opaque
      black](https://drafts.csswg.org/css-color/#opaque-black)]{#opaque-black
      .dfn}
    - [[transparent
      black](https://drafts.csswg.org/css-color/#transparent-black)]{#transparent-black
      .dfn}
    - [[\'srgb\'](https://drafts.csswg.org/css-color/#valdef-color-srgb)]{#'srgb'
      .dfn} color space
    - [[\'display-p3\'](https://drafts.csswg.org/css-color/#valdef-color-display-p3)]{#'display-p3'
      .dfn} color space
    - [[\'relative-colorimetric\'](https://drafts.csswg.org/css-color-5/#valdef-color-profile-rendering-intent-relative-colorimetric)]{#'relative-colorimetric'
      .dfn} rendering intent
    - [[parse a CSS \<color\>
      value](https://drafts.csswg.org/css-color/#parse-a-css-color-value)]{#parsed-as-a-css-color-value
      .dfn}
    - [[serialize a CSS \<color\>
      value](https://drafts.csswg.org/css-color/#serializing-color-values)]{#serialisation-of-a-color
      .dfn} including [[HTML-compatible serialization is
      requested](https://drafts.csswg.org/css-color/#color-serialization-html-compatible-serialization-is-requested)]{#html-compatible-serialization-is-requested
      .dfn}
    - [[Converting
      Colors](https://drafts.csswg.org/css-color/#color-conversion)]{#converting-colors
      .dfn}
    - [[\'color()\'](https://drafts.csswg.org/css-color/#color-function)]{#color-function
      .dfn}

    The following terms are defined in CSS Images:
    [\[CSSIMAGES\]](references.html#refsCSSIMAGES)

    - [[default object
      size](https://drafts.csswg.org/css-images/#default-object-size)]{#default-object-size
      .dfn}
    - [[concrete object
      size](https://drafts.csswg.org/css-images/#concrete-object-size)]{#concrete-object-size
      .dfn}
    - [[natural
      dimensions](https://drafts.csswg.org/css-images/#natural-dimensions)]{#natural-dimensions
      .dfn}
    - [[natural
      height](https://drafts.csswg.org/css-images/#natural-height)]{#natural-height
      .dfn}
    - [[natural
      width](https://drafts.csswg.org/css-images/#natural-width)]{#natural-width
      .dfn}
    - The
      [[\'image-orientation\'](https://drafts.csswg.org/css-images-3/#the-image-orientation)]{#'image-orientation'
      .dfn} property
    - [[\'conic-gradient\'](https://drafts.csswg.org/css-images-4/#funcdef-conic-gradient)]{#'conic-gradient'
      .dfn}
    - The
      [[\'object-fit\'](https://drafts.csswg.org/css-images/#the-object-fit)]{#'object-fit'
      .dfn} property

    The term [[paint
    source](https://drafts.csswg.org/css-images-4/#paint-source)]{#paint-source
    .dfn} is used as defined in CSS Images Level 4 to define the
    interaction of certain HTML elements with the CSS \'element()\'
    function. [\[CSSIMAGES4\]](references.html#refsCSSIMAGES4)

    The following features are defined in CSS Backgrounds and Borders:
    [\[CSSBG\]](references.html#refsCSSBG)

    - The
      [[\'background-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-color)]{#'background-color'
      .dfn},
      [[\'background-image\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-image)]{#'background-image'
      .dfn},
      [[\'background-repeat\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-repeat)]{#'background-repeat'
      .dfn},
      [[\'background-attachment\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-attachment)]{#'background-attachment'
      .dfn},
      [[\'background-position\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-position)]{#'background-position'
      .dfn},
      [[\'background-clip\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-clip)]{#'background-clip'
      .dfn},
      [[\'background-origin\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-origin)]{#'background-origin'
      .dfn}, and
      [[\'background-size\'](https://drafts.csswg.org/css-backgrounds/#propdef-background-size)]{#'background-size'
      .dfn} properties
    - The
      [[\'border-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-radius)]{#'border-radius'
      .dfn},
      [[\'border-top-left-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-left-radius)]{#'border-top-left-radius'
      .dfn},
      [[\'border-top-right-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-right-radius)]{#'border-top-right-radius'
      .dfn},
      [[\'border-bottom-right-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-right-radius)]{#'border-bottom-right-radius'
      .dfn},
      [[\'border-bottom-left-radius\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-left-radius)]{#'border-bottom-left-radius'
      .dfn} properties
    - The
      [[\'border-image-source\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-image-source)]{#'border-image-source'
      .dfn},
      [[\'border-image-slice\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-image-slice)]{#'border-image-slice'
      .dfn},
      [[\'border-image-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-image-width)]{#'border-image-width'
      .dfn},
      [[\'border-image-outset\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-image-outset)]{#'border-image-outset'
      .dfn}, and
      [[\'border-image-repeat\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-image-repeat)]{#'border-image-repeat'
      .dfn} properties

    CSS Backgrounds and Borders also defines the following border
    properties: [\[CSSBG\]](references.html#refsCSSBG)

    Border properties

    Top

    Bottom

    Left

    Right

    Width

    [[\'border-top-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-width)]{#'border-top-width'
    .dfn}

    [[\'border-bottom-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-width)]{#'border-bottom-width'
    .dfn}

    [[\'border-left-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-width)]{#'border-left-width'
    .dfn}

    [[\'border-right-width\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-width)]{#'border-right-width'
    .dfn}

    Style

    [[\'border-top-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-style)]{#'border-top-style'
    .dfn}

    [[\'border-bottom-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-style)]{#'border-bottom-style'
    .dfn}

    [[\'border-left-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-style)]{#'border-left-style'
    .dfn}

    [[\'border-right-style\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-style)]{#'border-right-style'
    .dfn}

    Color

    [[\'border-top-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-top-color)]{#'border-top-color'
    .dfn}

    [[\'border-bottom-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-bottom-color)]{#'border-bottom-color'
    .dfn}

    [[\'border-left-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-left-color)]{#'border-left-color'
    .dfn}

    [[\'border-right-color\'](https://drafts.csswg.org/css-backgrounds/#propdef-border-right-color)]{#'border-right-color'
    .dfn}

    The following features are defined in CSS Box Alignment:
    [\[CSSALIGN\]](references.html#refsCSSALIGN)

    - The
      [[\'align-content\'](https://drafts.csswg.org/css-align/#propdef-align-content)]{#'align-content'
      .dfn} property
    - The
      [[\'align-items\'](https://drafts.csswg.org/css-align/#propdef-align-items)]{#'align-items'
      .dfn} property
    - The
      [[\'align-self\'](https://drafts.csswg.org/css-align/#propdef-align-self)]{#'align-self'
      .dfn} property
    - The
      [[\'justify-self\'](https://drafts.csswg.org/css-align/#propdef-justify-self)]{#'justify-self'
      .dfn} property
    - The
      [[\'justify-content\'](https://drafts.csswg.org/css-align/#propdef-propdef-justify-content)]{#'justify-content'
      .dfn} property
    - The
      [[\'justify-items\'](https://drafts.csswg.org/css-align/#propdef-propdef-justify-items)]{#'justify-items'
      .dfn} property

    The following terms and features are defined in CSS Display:
    [\[CSSDISPLAY\]](references.html#refsCSSDISPLAY)

    - [[outer display
      type](https://drafts.csswg.org/css-display/#outer-display-type)]{#outer-display-type
      .dfn}
    - [[inner display
      type](https://drafts.csswg.org/css-display/#inner-display-type)]{#inner-display-type
      .dfn}
    - [[block-level](https://drafts.csswg.org/css-display/#block-level)]{#block-level
      .dfn}
    - [[block
      container](https://drafts.csswg.org/css-display/#block-container)]{#block-container
      .dfn}
    - [[formatting
      context](https://drafts.csswg.org/css-display/#formatting-context)]{#formatting-context
      .dfn}
    - [[block formatting
      context](https://drafts.csswg.org/css-display/#block-formatting-context)]{#block-formatting-context
      .dfn}
    - [[inline formatting
      context](https://drafts.csswg.org/css-display/#inline-formatting-context)]{#inline-formatting-context
      .dfn}
    - [[replaced
      element](https://drafts.csswg.org/css-display/#replaced-element)]{#replaced-element
      .dfn}
    - [[CSS
      box](https://drafts.csswg.org/css-display/#css-box)]{#css-box
      .dfn}

    The following features are defined in CSS Flexible Box Layout:
    [\[CSSFLEXBOX\]](references.html#refsCSSFLEXBOX)

    - The
      [[\'flex-direction\'](https://drafts.csswg.org/css-flexbox/#propdef-flex-direction)]{#'flex-direction'
      .dfn} property
    - The
      [[\'flex-wrap\'](https://drafts.csswg.org/css-flexbox/#propdef-flex-wrap)]{#'flex-wrap'
      .dfn} property

    The following terms and features are defined in CSS Fonts:
    [\[CSSFONTS\]](references.html#refsCSSFONTS)

    - [[first available
      font](https://drafts.csswg.org/css-fonts/#first-available-font)]{#first-available-font
      .dfn}
    - The
      [[\'font-family\'](https://drafts.csswg.org/css-fonts/#font-family-prop)]{#'font-family'
      .dfn} property
    - The
      [[\'font-weight\'](https://drafts.csswg.org/css-fonts/#font-weight-prop)]{#'font-weight'
      .dfn} property
    - The
      [[\'font-size\'](https://drafts.csswg.org/css-fonts/#font-size-prop)]{#'font-size'
      .dfn} property
    - The
      [[\'font\'](https://drafts.csswg.org/css-fonts/#font-prop)]{#'font'
      .dfn} property
    - The
      [[\'font-kerning\'](https://drafts.csswg.org/css-fonts/#propdef-font-kerning)]{#'font-kerning'
      .dfn} property
    - The
      [[\'font-stretch\'](https://drafts.csswg.org/css-fonts/#propdef-font-stretch)]{#'font-stretch'
      .dfn} property
    - The
      [[\'font-variant-caps\'](https://drafts.csswg.org/css-fonts/#propdef-font-variant-caps)]{#'font-variant-caps'
      .dfn} property
    - The
      [[\'small-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-small-caps)]{#'small-caps'
      .dfn} value
    - The
      [[\'all-small-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-all-small-caps)]{#'all-small-caps'
      .dfn} value
    - The
      [[\'petite-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-petite-caps)]{#'petite-caps'
      .dfn} value
    - The
      [[\'all-petite-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-all-petite-caps)]{#'all-petite-caps'
      .dfn} value
    - The
      [[\'unicase\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-unicase)]{#'unicase'
      .dfn} value
    - The
      [[\'titling-caps\'](https://drafts.csswg.org/css-fonts/#valdef-font-variant-caps-titling-caps)]{#'titling-caps'
      .dfn} value
    - The
      [[\'ultra-condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-ultra-condensed)]{#'ultra-condensed'
      .dfn} value
    - The
      [[\'extra-condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-extra-condensed)]{#'extra-condensed'
      .dfn} value
    - The
      [[\'condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-condensed)]{#'condensed'
      .dfn} value
    - The
      [[\'semi-condensed\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-semi-condensed)]{#'semi-condensed'
      .dfn} value
    - The
      [[\'semi-expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-semi-expanded)]{#'semi-expanded'
      .dfn} value
    - The
      [[\'expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-expanded)]{#'expanded'
      .dfn} value
    - The
      [[\'extra-expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-extra-expanded)]{#'extra-expanded'
      .dfn} value
    - The
      [[\'ultra-expanded\'](https://drafts.csswg.org/css-fonts/#valdef-font-stretch-ultra-expanded)]{#'ultra-expanded'
      .dfn} value

    The following features are defined in CSS Grid Layout:
    [\[CSSGRID\]](references.html#refsCSSGRID)

    - The
      [[\'grid-auto-columns\'](https://drafts.csswg.org/css-grid/#propdef-grid-auto-columns)]{#'grid-auto-columns'
      .dfn} property
    - The
      [[\'grid-auto-flow\'](https://drafts.csswg.org/css-grid/#propdef-grid-auto-flow)]{#'grid-auto-flow'
      .dfn} property
    - The
      [[\'grid-auto-rows\'](https://drafts.csswg.org/css-grid/#propdef-grid-auto-rows)]{#'grid-auto-rows'
      .dfn} property
    - The
      [[\'grid-column-gap\'](https://drafts.csswg.org/css-grid/#propdef-grid-column-gap)]{#'grid-column-gap'
      .dfn} property
    - The
      [[\'grid-row-gap\'](https://drafts.csswg.org/css-grid/#propdef-grid-row-gap)]{#'grid-row-gap'
      .dfn} property
    - The
      [[\'grid-template-areas\'](https://drafts.csswg.org/css-grid/#propdef-grid-template-areas)]{#'grid-template-areas'
      .dfn} property
    - The
      [[\'grid-template-columns\'](https://drafts.csswg.org/css-grid/#propdef-grid-template-columns)]{#'grid-template-columns'
      .dfn} property
    - The
      [[\'grid-template-rows\'](https://drafts.csswg.org/css-grid/#propdef-grid-template-rows)]{#'grid-template-rows'
      .dfn} property

    The following terms are defined in CSS Inline Layout:
    [\[CSSINLINE\]](references.html#refsCSSINLINE)

    - [[alphabetic
      baseline](https://drafts.csswg.org/css-inline/#alphabetic-baseline)]{#alphabetic-baseline
      .dfn}
    - [[ascent
      metric](https://drafts.csswg.org/css-inline/#ascent-metric)]{#ascent-metric
      .dfn}
    - [[descent
      metric](https://drafts.csswg.org/css-inline/#descent-metric)]{#descent-metric
      .dfn}
    - [[em-over
      baseline](https://drafts.csswg.org/css-inline/#em-over-baseline)]{#em-over-baseline
      .dfn}
    - [[em-under
      baseline](https://drafts.csswg.org/css-inline/#em-under-baseline)]{#em-under-baseline
      .dfn}
    - [[hanging
      baseline](https://drafts.csswg.org/css-inline/#hanging-baseline)]{#hanging-baseline
      .dfn}
    - [[ideographic-under
      baseline](https://drafts.csswg.org/css-inline/#ideographic-under-baseline)]{#ideographic-under-baseline
      .dfn}

    The following terms and features are defined in CSS Box Sizing:
    [\[CSSSIZING\]](references.html#refsCSSSIZING)

    - [[fit-content inline
      size](https://drafts.csswg.org/css-sizing/#fit-content-inline-size)]{#fit-content-inline-size
      .dfn}
    - [[\'aspect-ratio\'](https://drafts.csswg.org/css-sizing-4/#aspect-ratio)]{#'aspect-ratio'
      .dfn} property
    - [[intrinsic
      size](https://drafts.csswg.org/css-sizing/#intrinsic-size)]{#intrinsic-size
      .dfn}

    The following features are defined in CSS Lists and Counters.
    [\[CSSLISTS\]](references.html#refsCSSLISTS)

    - [[list
      item](https://drafts.csswg.org/css-lists/#list-item)]{#css-list-item
      .dfn}
    - The
      [[\'counter-reset\'](https://drafts.csswg.org/css-lists/#propdef-counter-reset)]{#'counter-reset'
      .dfn} property
    - The
      [[\'counter-set\'](https://drafts.csswg.org/css-lists/#propdef-counter-set)]{#'counter-set'
      .dfn} property
    - The
      [[\'list-style-type\'](https://drafts.csswg.org/css-lists/#propdef-list-style-type)]{#'list-style-type'
      .dfn} property

    The following features are defined in CSS Overflow.
    [\[CSSOVERFLOW\]](references.html#refsCSSOVERFLOW)

    - The
      [[\'overflow\'](https://drafts.csswg.org/css-overflow/#propdef-overflow)]{#'overflow'
      .dfn} property and its
      [[\'hidden\'](https://drafts.csswg.org/css-overflow/#valdef-overflow-hidden)]{#'hidden'
      .dfn} value
    - The
      [[\'text-overflow\'](https://drafts.csswg.org/css-overflow/#propdef-text-overflow)]{#'text-overflow'
      .dfn} property
    - The term [[scroll
      container](https://drafts.csswg.org/css-overflow/#scroll-container)]{#scroll-container
      .dfn}

    The following terms and features are defined in CSS Positioned
    Layout: [\[CSSPOSITION\]](references.html#refsCSSPOSITION)

    - [[absolutely-positioned](https://drafts.csswg.org/css-position/#absolute-position)]{#absolutely-positioned
      .dfn}
    - The
      [[\'position\'](https://drafts.csswg.org/css-position/#position-property)]{#'position'
      .dfn} property and its
      [[\'static\'](https://drafts.csswg.org/css-position/#valdef-position-static)]{#'static'
      .dfn} value
    - The [[top
      layer](https://drafts.csswg.org/css-position-4/#document-top-layer)]{#top-layer
      .dfn} (an [ordered
      set](https://infra.spec.whatwg.org/#ordered-set){#dependencies:set
      x-internal="set"})
    - [[add an element to the top
      layer](https://drafts.csswg.org/css-position-4/#add-an-element-to-the-top-layer)]{#add-an-element-to-the-top-layer
      .dfn}
    - [[request an element to be removed from the top
      layer](https://drafts.csswg.org/css-position-4/#request-an-element-to-be-removed-from-the-top-layer)]{#request-an-element-to-be-removed-from-the-top-layer
      .dfn}
    - [[remove an element from the top layer
      immediately](https://drafts.csswg.org/css-position-4/#remove-an-element-from-the-top-layer-immediately)]{#remove-an-element-from-the-top-layer-immediately
      .dfn}
    - [[process top layer
      removals](https://drafts.csswg.org/css-position-4/#process-top-layer-removals)]{#process-top-layer-removals
      .dfn}

    The following features are defined in CSS Multi-column Layout.
    [\[CSSMULTICOL\]](references.html#refsCSSMULTICOL)

    - The
      [[\'column-count\'](https://drafts.csswg.org/css-multicol/#propdef-column-count)]{#'column-count'
      .dfn} property
    - The
      [[\'column-fill\'](https://drafts.csswg.org/css-multicol/#propdef-column-fill)]{#'column-fill'
      .dfn} property
    - The
      [[\'column-gap\'](https://drafts.csswg.org/css-multicol/#propdef-column-gap)]{#'column-gap'
      .dfn} property
    - The
      [[\'column-rule\'](https://drafts.csswg.org/css-multicol/#propdef-column-rule)]{#'column-rule'
      .dfn} property
    - The
      [[\'column-width\'](https://drafts.csswg.org/css-multicol/#propdef-column-width)]{#'column-width'
      .dfn} property

    The
    [[\'ruby-base\'](https://drafts.csswg.org/css-ruby/#valdef-display-ruby-base)]{#'ruby-base'
    .dfn} value of the
    [\'display\'](https://drafts.csswg.org/css2/#display-prop){#dependencies:'display'-2
    x-internal="'display'"} property is defined in CSS Ruby Layout.
    [\[CSSRUBY\]](references.html#refsCSSRUBY)

    The following features are defined in CSS Table:
    [\[CSSTABLE\]](references.html#refsCSSTABLE)

    - The
      [[\'border-spacing\'](https://drafts.csswg.org/css-tables/#propdef-border-spacing)]{#'border-spacing'
      .dfn} property
    - The
      [[\'border-collapse\'](https://drafts.csswg.org/css-tables/#border-collapse-property)]{#'border-collapse'
      .dfn} property
    - The
      [[\'table-cell\'](https://drafts.csswg.org/css-tables/#table-cell)]{#'table-cell'
      .dfn},
      [[\'table-row\'](https://drafts.csswg.org/css-tables/#table-row)]{#'table-row'
      .dfn},
      [[\'table-caption\'](https://drafts.csswg.org/css-tables/#table-caption)]{#'table-caption'
      .dfn}, and
      [[\'table\'](https://drafts.csswg.org/css-tables/#table)]{#'table'
      .dfn} values of the
      [\'display\'](https://drafts.csswg.org/css2/#display-prop){#dependencies:'display'-3
      x-internal="'display'"} property

    The following features are defined in CSS Text:
    [\[CSSTEXT\]](references.html#refsCSSTEXT)

    - The [[content
      language](https://drafts.csswg.org/css-text-4/#content-language)]{#content-language
      .dfn} concept
    - The
      [[\'text-transform\'](https://drafts.csswg.org/css-text/#text-transform-property)]{#'text-transform'
      .dfn} property
    - The
      [[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property)]{#'white-space'
      .dfn} property
    - The
      [[\'text-align\'](https://drafts.csswg.org/css-text/#text-align-property)]{#'text-align'
      .dfn} property
    - The
      [[\'letter-spacing\'](https://drafts.csswg.org/css-text/#letter-spacing-property)]{#'letter-spacing'
      .dfn} property
    - The
      [[\'word-spacing\'](https://drafts.csswg.org/css-text/#propdef-word-spacing)]{#'word-spacing'
      .dfn} property

    The following features are defined in CSS Writing Modes:
    [\[CSSWM\]](references.html#refsCSSWM)

    - The
      [[\'direction\'](https://drafts.csswg.org/css-writing-modes/#direction)]{#'direction'
      .dfn} property
    - The
      [[\'unicode-bidi\'](https://drafts.csswg.org/css-writing-modes/#unicode-bidi)]{#'unicode-bidi'
      .dfn} property
    - The
      [[\'writing-mode\'](https://drafts.csswg.org/css-writing-modes/#propdef-writing-mode)]{#'writing-mode'
      .dfn} property
    - The [[block flow
      direction](https://drafts.csswg.org/css-writing-modes/#block-flow-direction)]{#block-flow-direction
      .dfn}, [[block
      axis](https://drafts.csswg.org/css-writing-modes/#block-axis)]{#block-axis
      .dfn}, [[inline
      axis](https://drafts.csswg.org/css-writing-modes/#inline-axis)]{#inline-axis
      .dfn}, [[block
      size](https://drafts.csswg.org/css-writing-modes/#block-size)]{#block-size
      .dfn}, [[inline
      size](https://drafts.csswg.org/css-writing-modes/#inline-size)]{#inline-size
      .dfn},
      [[block-start](https://drafts.csswg.org/css-writing-modes/#block-start)]{#block-start
      .dfn},
      [[block-end](https://drafts.csswg.org/css-writing-modes/#block-end)]{#block-end
      .dfn},
      [[inline-start](https://drafts.csswg.org/css-writing-modes/#inline-start)]{#inline-start
      .dfn},
      [[inline-end](https://drafts.csswg.org/css-writing-modes/#inline-end)]{#inline-end
      .dfn},
      [[line-left](https://drafts.csswg.org/css-writing-modes/#line-left)]{#line-left
      .dfn}, and
      [[line-right](https://drafts.csswg.org/css-writing-modes/#line-right)]{#line-right
      .dfn} concepts

    The following features are defined in CSS Basic User Interface:
    [\[CSSUI\]](references.html#refsCSSUI)

    - The
      [[\'outline\'](https://drafts.csswg.org/css-ui/#outline)]{#'outline'
      .dfn} property
    - The
      [[\'cursor\'](https://drafts.csswg.org/css-ui/#cursor)]{#'cursor'
      .dfn} property
    - The
      [[\'appearance\'](https://drafts.csswg.org/css-ui/#appearance-switching)]{#'appearance'
      .dfn} property, its
      [[\<compat-auto\>](https://drafts.csswg.org/css-ui/#typedef-appearance-compat-auto)]{#compat-auto
      .dfn} non-terminal value type, its
      [[\'textfield\'](https://drafts.csswg.org/css-ui/#valdef-appearance-textfield)]{#'textfield'
      .dfn} value, and its
      [[\'menulist-button\'](https://drafts.csswg.org/css-ui/#valdef-appearance-menulist-button)]{#'menulist-button'
      .dfn} value.
    - The
      [[\'field-sizing\'](https://drafts.csswg.org/css-ui/#field-sizing)]{#'field-sizing'
      .dfn} property, and its
      [[\'content\'](https://drafts.csswg.org/css-ui/#valdef-field-sizing-content)]{#field-sizing-content
      .dfn} value.
    - The concept
      [[widget](https://drafts.csswg.org/css-ui/#widget)]{#widget .dfn}
    - The concept [[native
      appearance](https://drafts.csswg.org/css-ui/#native-appearance)]{#native-appearance
      .dfn}
    - The concept [[primitive
      appearance](https://drafts.csswg.org/css-ui/#primitive-appearance)]{#primitive-appearance
      .dfn}
    - The concept [[element with default preferred
      size](https://drafts.csswg.org/css-ui/#element-with-default-preferred-size)]{#element-with-default-preferred-size
      .dfn}
    - The [[non-devolvable
      widget](https://drafts.csswg.org/css-ui/#non-devolvable)]{#non-devolvable-widget
      .dfn} and [[devolvable
      widget](https://drafts.csswg.org/css-ui/#devolvable)]{#devolvable-widget
      .dfn} classification, and the related [[devolved
      widget](https://drafts.csswg.org/css-ui/#devolved)]{#devolved-widget
      .dfn} state.
    - The
      [[\'pointer-events\'](https://drafts.csswg.org/css-ui-4/#pointer-events-control)]{#'pointer-events'
      .dfn} property
    - The
      [[\'user-select\'](https://drafts.csswg.org/css-ui-4/#content-selection)]{#'user-select'
      .dfn} property

    The algorithm to [[update animations and send
    events](https://drafts.csswg.org/web-animations-1/#update-animations-and-send-events)]{#update-animations-and-send-events
    .dfn} is defined in Web Animations.
    [\[WEBANIMATIONS\]](references.html#refsWEBANIMATIONS)

    Implementations that support scripting must support the CSS Object
    Model. The following features and terms are defined in the CSSOM
    specifications: [\[CSSOM\]](references.html#refsCSSOM)
    [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    - [[`Screen`](https://drafts.csswg.org/cssom-view/#the-screen-interface)]{#screen
      .dfn} interface
    - [[`LinkStyle`](https://drafts.csswg.org/cssom/#the-linkstyle-interface)]{#linkstyle
      .dfn} interface
    - [[`CSSStyleDeclaration`](https://drafts.csswg.org/cssom/#the-cssstyledeclaration-interface)]{#cssstyledeclaration
      .dfn} interface
    - [[`style`](https://drafts.csswg.org/cssom/#dom-elementcssinlinestyle-style)]{#dom-style
      .dfn} IDL attribute
    - [[`cssText`](https://drafts.csswg.org/cssom/#dom-cssstyledeclaration-csstext)]{#dom-cssstyledeclaration-csstext
      .dfn} attribute of
      [`CSSStyleDeclaration`{#dependencies:cssstyledeclaration}](https://drafts.csswg.org/cssom/#the-cssstyledeclaration-interface){x-internal="cssstyledeclaration"}
    - [[`StyleSheet`](https://drafts.csswg.org/cssom/#the-stylesheet-interface)]{#stylesheet
      .dfn} interface
    - [[`CSSStyleSheet`](https://drafts.csswg.org/cssom/#the-cssstylesheet-interface)]{#cssstylesheet
      .dfn} interface
    - [[create a CSS style
      sheet](https://drafts.csswg.org/cssom/#create-a-css-style-sheet)]{#create-a-css-style-sheet
      .dfn}
    - [[remove a CSS style
      sheet](https://drafts.csswg.org/cssom/#remove-a-css-style-sheet)]{#remove-a-css-style-sheet
      .dfn}
    - [[associated CSS style
      sheet](https://drafts.csswg.org/cssom/#associated-css-style-sheet)]{#associated-css-style-sheet
      .dfn}
    - [[create a constructed
      `CSSStyleSheet`](https://drafts.csswg.org/cssom/#create-a-constructed-cssstylesheet)]{#create-a-constructed-cssstylesheet
      .dfn}
    - [[synchronously replace the rules of a
      `CSSStyleSheet`](https://drafts.csswg.org/cssom/#synchronously-replace-the-rules-of-a-cssstylesheet)]{#synchronously-replace-the-rules-of-a-cssstylesheet
      .dfn}
    - [[disable a CSS style
      sheet](https://drafts.csswg.org/cssom/#disable-a-css-style-sheet)]{#disable-a-css-style-sheet
      .dfn}
    - [[CSS style
      sheets](https://drafts.csswg.org/cssom/#css-style-sheet)]{#css-style-sheet
      .dfn} and their properties:
      - [[type](https://drafts.csswg.org/cssom/#concept-css-style-sheet-type)]{#concept-css-style-sheet-type
        .dfn}
      - [[location](https://drafts.csswg.org/cssom/#concept-css-style-sheet-location)]{#concept-css-style-sheet-location
        .dfn}
      - [[parent CSS style
        sheet](https://drafts.csswg.org/cssom/#concept-css-style-sheet-parent-css-style-sheet)]{#concept-css-style-sheet-parent-css-style-sheet
        .dfn}
      - [[owner
        node](https://drafts.csswg.org/cssom/#concept-css-style-sheet-owner-node)]{#concept-css-style-sheet-owner-node
        .dfn}
      - [[owner CSS
        rule](https://drafts.csswg.org/cssom/#concept-css-style-sheet-owner-css-rule)]{#concept-css-style-sheet-owner-css-rule
        .dfn}
      - [[media](https://drafts.csswg.org/cssom/#concept-css-style-sheet-media)]{#concept-css-style-sheet-media
        .dfn}
      - [[title](https://drafts.csswg.org/cssom/#concept-css-style-sheet-title)]{#concept-css-style-sheet-title
        .dfn}
      - [[alternate
        flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-alternate-flag)]{#concept-css-style-sheet-alternate-flag
        .dfn}
      - [[disabled
        flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-disabled-flag)]{#concept-css-style-sheet-disabled-flag
        .dfn}
      - [[CSS
        rules](https://drafts.csswg.org/cssom/#concept-css-style-sheet-css-rules)]{#concept-css-style-sheet-css-rules
        .dfn}
      - [[origin-clean
        flag](https://drafts.csswg.org/cssom/#concept-css-style-sheet-origin-clean-flag)]{#concept-css-style-sheet-origin-clean-flag
        .dfn}
    - [[CSS style sheet
      set](https://drafts.csswg.org/cssom/#css-style-sheet-set)]{#css-style-sheet-set
      .dfn}
    - [[CSS style sheet set
      name](https://drafts.csswg.org/cssom/#css-style-sheet-set-name)]{#css-style-sheet-set-name
      .dfn}
    - [[preferred CSS style sheet set
      name](https://drafts.csswg.org/cssom/#preferred-css-style-sheet-set-name)]{#preferred-css-style-sheet-set-name
      .dfn}
    - [[change the preferred CSS style sheet set
      name](https://drafts.csswg.org/cssom/#change-the-preferred-css-style-sheet-set-name)]{#change-the-preferred-css-style-sheet-set-name
      .dfn}
    - [[Serializing a CSS
      value](https://drafts.csswg.org/cssom/#serialize-a-css-value)]{#serializing-a-css-value
      .dfn}
    - [[run the resize
      steps](https://drafts.csswg.org/cssom-view/#document-run-the-resize-steps)]{#run-the-resize-steps
      .dfn}
    - [[run the scroll
      steps](https://drafts.csswg.org/cssom-view/#document-run-the-scroll-steps)]{#run-the-scroll-steps
      .dfn}
    - [[evaluate media queries and report
      changes](https://drafts.csswg.org/cssom-view/#evaluate-media-queries-and-report-changes)]{#evaluate-media-queries-and-report-changes
      .dfn}
    - [[Scroll a target into
      view](https://drafts.csswg.org/cssom-view/#scroll-a-target-into-view)]{#scroll-a-target-into-view
      .dfn}
    - [[Scroll to the beginning of the
      document](https://drafts.csswg.org/cssom-view/#scroll-to-the-beginning-of-the-document)]{#scroll-to-the-beginning-of-the-document
      .dfn}
    - The
      [[`resize`](https://drafts.csswg.org/cssom-view/#eventdef-window-resize)]{#event-resize
      .dfn} event
    - The
      [[`scroll`](https://drafts.csswg.org/cssom-view/#eventdef-document-scroll)]{#event-scroll
      .dfn} event
    - The
      [[`scrollend`](https://drafts.csswg.org/cssom-view/#eventdef-document-scrollend)]{#event-scrollend
      .dfn} event
    - [[set up browsing context
      features](https://drafts.csswg.org/cssom-view/#set-up-browsing-context-features)]{#set-up-browsing-context-features
      .dfn}
    - The
      [[clientX](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clientx)]{#mouseevent-clientx
      .dfn} and
      [[clientY](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clienty)]{#mouseevent-clienty
      .dfn} extension attributes of the
      [MouseEvent](https://w3c.github.io/uievents/#mouseevent){#dependencies:mouseevent-2
      x-internal="mouseevent"} interface

    The following features and terms are defined in CSS Syntax:
    [\[CSSSYNTAX\]](references.html#refsCSSSYNTAX)

    - [[conformant style
      sheet](https://drafts.csswg.org/css-syntax/#conform-classes)]{#conformant-style-sheet
      .dfn}
    - [[parse a list of component
      values](https://drafts.csswg.org/css-syntax/#parse-a-list-of-component-values)]{#parse-a-list-of-component-values
      .dfn}
    - [[parse a comma-separated list of component
      values](https://drafts.csswg.org/css-syntax/#parse-a-comma-separated-list-of-component-values)]{#parse-a-comma-separated-list-of-component-values
      .dfn}
    - [[component
      value](https://drafts.csswg.org/css-syntax/#component-value)]{#component-value
      .dfn}
    - [[environment
      encoding](https://drafts.csswg.org/css-syntax/#environment-encoding)]{#environment-encoding
      .dfn}
    - [[\<whitespace-token\>](https://drafts.csswg.org/css-syntax/#typedef-whitespace-token)]{#whitespace-token
      .dfn}

    The following terms are defined in Selectors:
    [\[SELECTORS\]](references.html#refsSELECTORS)

    - [[type
      selector](https://drafts.csswg.org/selectors/#type-selector)]{#type-selector
      .dfn}
    - [[attribute
      selector](https://drafts.csswg.org/selectors/#attribute-selector)]{#attribute-selector
      .dfn}
    - [[pseudo-class](https://drafts.csswg.org/selectors/#pseudo-class)]{#pseudo-class
      .dfn}
    - [[`:focus-visible`](https://drafts.csswg.org/selectors/#the-focus-visible-pseudo)]{#:focus-visible
      .dfn} pseudo-class
    - [[indicate
      focus](https://drafts.csswg.org/selectors/#indicate-focus)]{#indicate-focus
      .dfn}
    - [[pseudo-element](https://drafts.csswg.org/selectors/#pseudo-element)]{#pseudo-element
      .dfn}

    The following features are defined in CSS Values and Units:
    [\[CSSVALUES\]](references.html#refsCSSVALUES)

    - [[\<length\>](https://drafts.csswg.org/css-values/#lengths)]{#length-2
      .dfn}
    - The [[\'em\'](https://drafts.csswg.org/css-values/#em)]{#'em'
      .dfn} unit
    - The [[\'ex\'](https://drafts.csswg.org/css-values/#ex)]{#'ex'
      .dfn} unit
    - The [[\'vw\'](https://drafts.csswg.org/css-values/#vw)]{#'vw'
      .dfn} unit
    - The [[\'in\'](https://drafts.csswg.org/css-values/#in)]{#'in'
      .dfn} unit
    - The [[\'px\'](https://drafts.csswg.org/css-values/#px)]{#'px'
      .dfn} unit
    - The [[\'pt\'](https://drafts.csswg.org/css-values/#pt)]{#'pt'
      .dfn} unit
    - The
      [[\'attr()\'](https://drafts.csswg.org/css-values/#funcdef-attr)]{#'attr()'
      .dfn} function
    - The [[math
      functions](https://drafts.csswg.org/css-values/#math-function)]{#math-functions
      .dfn}

    The following features are defined in CSS View Transitions:
    [\[CSSVIEWTRANSITIONS\]](references.html#refsCSSVIEWTRANSITIONS)

    - [[perform pending transition
      operations](https://drafts.csswg.org/css-view-transitions/#perform-pending-transition-operations)]{#perform-pending-transition-operations
      .dfn}
    - [[rendering suppression for view
      transitions](https://drafts.csswg.org/css-view-transitions/#document-rendering-suppression-for-view-transitions)]{#rendering-suppression-for-view-transitions
      .dfn}
    - [[activate view
      transition](https://drafts.csswg.org/css-view-transitions/#activate-view-transition)]{#activate-view-transition
      .dfn}
    - [[`ViewTransition`](https://drafts.csswg.org/css-view-transitions/#viewtransition)]{#viewtransition
      .dfn}
    - [[view transition page visibility change
      steps](https://drafts.csswg.org/css-view-transitions/#view-transition-page-visibility-change-steps)]{#view-transition-page-visibility-change-steps
      .dfn}
    - [[resolving inbound cross-document
      view-transition](https://drafts.csswg.org/css-view-transitions-2/#resolve-inbound-cross-document-view-transition)]{#resolving-inbound-cross-document-view-transition
      .dfn}
    - [[setting up a cross-document
      view-transition](https://drafts.csswg.org/css-view-transitions-2/#setup-cross-document-view-transition)]{#setting-up-a-cross-document-view-transition
      .dfn}
    - [[can navigation trigger a cross-document
      view-transition?](https://drafts.csswg.org/css-view-transitions-2/#can-navigation-trigger-a-cross-document-view-transition)]{#can-navigation-trigger-a-cross-document-view-transition
      .dfn}

    The term [[style
    attribute](https://drafts.csswg.org/css-style-attr/#style-attribute)]{#css-styling-attribute
    .dfn} is defined in CSS Style Attributes.
    [\[CSSATTR\]](references.html#refsCSSATTR)

    The following terms are defined in the CSS Cascading and
    Inheritance: [\[CSSCASCADE\]](references.html#refsCSSCASCADE)

    - [[cascaded
      value](https://drafts.csswg.org/css-cascade/#cascaded-value)]{#cascaded-value
      .dfn}
    - [[specified
      value](https://drafts.csswg.org/css-cascade/#specified-value)]{#specified-value
      .dfn}
    - [[computed
      value](https://drafts.csswg.org/css-cascade/#computed-value)]{#computed-value
      .dfn}
    - [[used
      value](https://drafts.csswg.org/css-cascade/#used-value)]{#used-value
      .dfn}
    - [[cascade
      origin](https://drafts.csswg.org/css-cascade/#origin)]{#cascade-origin
      .dfn}
    - [[Author
      Origin](https://drafts.csswg.org/css-cascade/#cascade-origin-author)]{#author-origin
      .dfn}
    - [[User
      Origin](https://drafts.csswg.org/css-cascade/#cascade-origin-user)]{#user-origin
      .dfn}
    - [[User Agent
      Origin](https://drafts.csswg.org/css-cascade/#cascade-origin-ua)]{#user-agent-origin
      .dfn}
    - [[Animation
      Origin](https://drafts.csswg.org/css-cascade/#cascade-origin-animation)]{#animation-origin
      .dfn}
    - [[Transition
      Origin](https://drafts.csswg.org/css-cascade/#cascade-origin-transition)]{#transition-origin
      .dfn}
    - [[initial
      value](https://drafts.csswg.org/css-cascade/#initial-value)]{#initial-value
      .dfn}

    The
    [`CanvasRenderingContext2D`{#dependencies:canvasrenderingcontext2d}](canvas.html#canvasrenderingcontext2d)
    object\'s use of fonts depends on the features described in the CSS
    Fonts and Font Loading specifications, including in particular
    [`FontFace`]{#fontface .dfn} objects and the [[font
    source](https://drafts.csswg.org/css-font-loading/#font-source)]{#font-source
    .dfn} concept. [\[CSSFONTS\]](references.html#refsCSSFONTS)
    [\[CSSFONTLOAD\]](references.html#refsCSSFONTLOAD)

    The following interfaces and terms are defined in Geometry
    Interfaces: [\[GEOMETRY\]](references.html#refsGEOMETRY)

    - [[`DOMMatrix`](https://drafts.fxtf.org/geometry/#dommatrix)]{#dommatrix
      .dfn} interface, and associated [[m11
      element](https://drafts.fxtf.org/geometry/#matrix-m11-element)]{#m11-element
      .dfn}, [[m12
      element](https://drafts.fxtf.org/geometry/#matrix-m12-element)]{#m12-element
      .dfn}, [[m21
      element](https://drafts.fxtf.org/geometry/#matrix-m21-element)]{#m21-element
      .dfn}, [[m22
      element](https://drafts.fxtf.org/geometry/#matrix-m22-element)]{#m22-element
      .dfn}, [[m41
      element](https://drafts.fxtf.org/geometry/#matrix-m41-element)]{#m41-element
      .dfn}, and [[m42
      element](https://drafts.fxtf.org/geometry/#matrix-m42-element)]{#m42-element
      .dfn}
    - [[`DOMMatrix2DInit`](https://drafts.fxtf.org/geometry/#dictdef-dommatrix2dinit)]{#dommatrix2dinit
      .dfn} and
      [[`DOMMatrixInit`](https://drafts.fxtf.org/geometry/#dictdef-dommatrixinit)]{#dommatrixinit
      .dfn} dictionaries
    - The [[create a `DOMMatrix` from a
      dictionary](https://drafts.fxtf.org/geometry/#create-a-dommatrix-from-the-dictionary)]{#create-a-dommatrix-from-a-dictionary
      .dfn} and [[create a `DOMMatrix` from a 2D
      dictionary](https://drafts.fxtf.org/geometry/#create-a-dommatrix-from-the-2d-dictionary)]{#create-a-dommatrix-from-a-2d-dictionary
      .dfn} algorithms for
      [`DOMMatrix2DInit`{#dependencies:dommatrix2dinit}](https://drafts.fxtf.org/geometry/#dictdef-dommatrix2dinit){x-internal="dommatrix2dinit"}
      or
      [`DOMMatrixInit`{#dependencies:dommatrixinit}](https://drafts.fxtf.org/geometry/#dictdef-dommatrixinit){x-internal="dommatrixinit"}
    - The
      [[`DOMPointInit`](https://drafts.fxtf.org/geometry/#dictdef-dompointinit)]{#dompointinit
      .dfn} dictionary, and associated
      [[x](https://drafts.fxtf.org/geometry/#dom-dompointinit-x)]{#dompointinit-x
      .dfn} and
      [[y](https://drafts.fxtf.org/geometry/#dom-dompointinit-y)]{#dompointinit-y
      .dfn} members
    - [[Matrix
      multiplication](https://drafts.fxtf.org/geometry/#matrix-multiply)]{#matrix-multiplication
      .dfn}

    The following terms are defined in the CSS Scoping:
    [\[CSSSCOPING\]](references.html#refsCSSSCOPING)

    - [[flat
      tree](https://drafts.csswg.org/css-scoping/#flat-tree)]{#flat-tree
      .dfn}

    The following terms and features are defined in CSS Color
    Adjustment: [\[CSSCOLORADJUST\]](references.html#refsCSSCOLORADJUST)

    - [[\'color-scheme\'](https://drafts.csswg.org/css-color-adjust/#color-scheme-prop)]{#'color-scheme'
      .dfn}
    - [[page\'s supported
      color-schemes](https://drafts.csswg.org/css-color-adjust/#pages-supported-color-schemes)]{#page's-supported-color-schemes
      .dfn}

    The following terms are defined in CSS Pseudo-Elements:
    [\[CSSPSEUDO\]](references.html#refsCSSPSEUDO)

    - [[\'::details-content\'](https://drafts.csswg.org/css-pseudo/#details-content-pseudo)]{#'::details-content'
      .dfn}
    - [[\'::file-selector-button\'](https://drafts.csswg.org/css-pseudo/#file-selector-button-pseudo)]{#'::file-selector-button'
      .dfn}

    The following terms are defined in CSS Containment:
    [\[CSSCONTAIN\]](references.html#refsCSSCONTAIN)

    - [[skips its
      contents](https://drafts.csswg.org/css-contain/#skips-its-contents)]{#skips-its-contents
      .dfn}
    - [[relevant to the
      user](https://drafts.csswg.org/css-contain/#relevant-to-the-user)]{#relevant-to-the-user
      .dfn}
    - [[proximity to the
      viewport](https://drafts.csswg.org/css-contain/#proximity-to-the-viewport)]{#proximity-to-the-viewport
      .dfn}
    - [[layout
      containment](https://drafts.csswg.org/css-contain/#containment-layout)]{#layout-containment
      .dfn}
    - [[\'content-visibility\'](https://drafts.csswg.org/css-contain/#content-visibility)]{#'content-visibility'
      .dfn} property
    - [[\'auto\'](https://drafts.csswg.org/css-contain/#propdef-content-visibility)]{#content-visibility-auto
      .dfn} value for
      [\'content-visibility\'](https://drafts.csswg.org/css-contain/#content-visibility){#dependencies:'content-visibility'
      x-internal="'content-visibility'"}

    The following terms are defined in CSS Anchor Positioning:
    [\[CSSANCHOR\]](references.html#refsCSSANCHOR)

    - [[implicit anchor
      element](https://drafts.csswg.org/css-anchor-position/#implicit-anchor-element)]{#implicit-anchor-element
      .dfn}

Intersection Observer

:   The following term is defined in Intersection Observer:
    [\[INTERSECTIONOBSERVER\]](references.html#refsINTERSECTIONOBSERVER)

    - [[run the update intersection observations
      steps](https://w3c.github.io/IntersectionObserver/#run-the-update-intersection-observations-steps)]{#run-the-update-intersection-observations-steps
      .dfn}
    - [[`IntersectionObserver`](https://w3c.github.io/IntersectionObserver/#intersectionobserver)]{#intersectionobserver
      .dfn}
    - [[`IntersectionObserverInit`](https://w3c.github.io/IntersectionObserver/#dictdef-intersectionobserverinit)]{#intersectionobserverinit
      .dfn}
    - [[`observe`](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserver-observe)]{#dom-intersectionobserver-observe
      .dfn}
    - [[`unobserve`](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserver-unobserve)]{#dom-intersectionobserver-unobserve
      .dfn}
    - [[`isIntersecting`](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-isintersecting)]{#dom-intersectionobserverentry-isintersecting
      .dfn}
    - [[`target`](https://w3c.github.io/IntersectionObserver/#dom-intersectionobserverentry-target)]{#dom-intersectionobserverentry-target
      .dfn}

Resize Observer

:   The following terms are defined in Resize Observer:
    [\[RESIZEOBSERVER\]](references.html#refsRESIZEOBSERVER)

    - [[gather active resize observations at
      depth](https://drafts.csswg.org/resize-observer-1/#gather-active-observations-h)]{#gather-active-resize-observations-at-depth
      .dfn}
    - [[has active resize
      observations](https://drafts.csswg.org/resize-observer-1/#has-active-observations-h)]{#has-active-resize-observations
      .dfn}
    - [[has skipped resize
      observations](https://drafts.csswg.org/resize-observer-1/#has-skipped-observations-h)]{#has-skipped-resize-observations
      .dfn}
    - [[broadcast active resize
      observations](https://drafts.csswg.org/resize-observer-1/#broadcast-resize-notifications-h)]{#broadcast-active-resize-observations
      .dfn}
    - [[deliver resize loop
      error](https://drafts.csswg.org/resize-observer-1/#deliver-resize-error)]{#deliver-resize-loop-error
      .dfn}

WebGL

:   The following interfaces are defined in the WebGL specifications:
    [\[WEBGL\]](references.html#refsWEBGL)

    - [[`WebGLRenderingContext`](https://www.khronos.org/registry/webgl/specs/latest/1.0/#WebGLRenderingContext)]{#webglrenderingcontext
      .dfn} interface
    - [[`WebGL2RenderingContext`](https://www.khronos.org/registry/webgl/specs/latest/2.0/#WebGL2RenderingContext)]{#webgl2renderingcontext
      .dfn} interface
    - [[`WebGLContextAttributes`](https://www.khronos.org/registry/webgl/specs/latest/1.0/#WebGLContextAttributes)]{#webglcontextattributes
      .dfn} dictionary

WebGPU

:   The following interfaces are defined in WebGPU:
    [\[WEBGPU\]](references.html#refsWEBGPU)

    - [[`GPUCanvasContext`](https://gpuweb.github.io/gpuweb/#canvas-context)]{#gpucanvascontext
      .dfn} interface

WebVTT

:   Implementations may support WebVTT as a text track format for
    subtitles, captions, metadata, etc., for media resources.
    [\[WEBVTT\]](references.html#refsWEBVTT)

    The following terms, used in this specification, are defined in
    WebVTT:

    - [[WebVTT
      file](https://w3c.github.io/webvtt/#webvtt-file)]{#webvtt-file
      .dfn}
    - [[WebVTT file using cue
      text](https://w3c.github.io/webvtt/#webvtt-file-using-cue-text)]{#webvtt-file-using-cue-text
      .dfn}
    - [[WebVTT file using only nested
      cues](https://w3c.github.io/webvtt/#webvtt-file-using-only-nested-cues)]{#webvtt-file-using-only-nested-cues
      .dfn}
    - [[WebVTT
      parser](https://w3c.github.io/webvtt/#webvtt-parser)]{#webvtt-parser
      .dfn}
    - The [[rules for updating the display of WebVTT text
      tracks](https://w3c.github.io/webvtt/#rules-for-updating-the-display-of-webvtt-text-tracks)]{#rules-for-updating-the-display-of-webvtt-text-tracks
      .dfn}
    - The WebVTT [[text track cue writing
      direction](https://w3c.github.io/webvtt/#webvtt-cue-writing-direction)]{#text-track-cue-writing-direction
      .dfn}
    - [[`VTTCue`](https://w3c.github.io/webvtt/#vttcue)]{#vttcue .dfn}
      interface

ARIA

:   The [`role`]{#attr-aria-role .dfn dfn-type="element-attr"} attribute
    is defined in Accessible Rich Internet Applications (ARIA), as are
    the following roles: [\[ARIA\]](references.html#refsARIA)

    - [[`button`](https://w3c.github.io/aria/#button)]{#attr-aria-role-button
      .dfn}
    - [[`presentation`](https://w3c.github.io/aria/#presentation)]{#attr-aria-role-presentation
      .dfn}

    In addition, the following [`aria-*`]{#attr-aria-* .dfn} content
    attributes are defined in ARIA: [\[ARIA\]](references.html#refsARIA)

    - [[`aria-checked`](https://w3c.github.io/aria/#aria-checked)]{#attr-aria-checked
      .dfn}
    - [[`aria-describedby`](https://w3c.github.io/aria/#aria-describedby)]{#attr-aria-describedby
      .dfn}
    - [[`aria-disabled`](https://w3c.github.io/aria/#aria-disabled)]{#attr-aria-disabled
      .dfn}
    - [[`aria-label`](https://w3c.github.io/aria/#aria-label)]{#attr-aria-label
      .dfn}

    Finally, the following terms are defined in ARIA:
    [\[ARIA\]](references.html#refsARIA)

    - [[role](https://w3c.github.io/aria/#dfn-role)]{#role .dfn}
    - [[accessible
      name](https://w3c.github.io/aria/#dfn-accessible-name)]{#concept-accessible-name
      .dfn}
    - The
      [[`ARIAMixin`](https://w3c.github.io/aria/#ARIAMixin)]{#ariamixin
      .dfn} interface, with its associated [[`ARIAMixin` getter
      steps](https://w3c.github.io/aria/#dfn-ariamixin-getter-steps)]{#ariamixin-getter-steps
      .dfn} and [[`ARIAMixin` setter
      steps](https://w3c.github.io/aria/#dfn-ariamixin-setter-steps)]{#ariamixin-setter-steps
      .dfn} hooks, and its
      [[`role`](https://w3c.github.io/aria/#idl-def-ariamixin-role)]{#dom-ariamixin-role
      .dfn} and
      [[`aria*`](https://w3c.github.io/aria/#idl-def-ariamixin-ariaactivedescendantelement)]{#dom-ariamixin-aria*
      .dfn} attributes

Content Security Policy

:   The following terms are defined in Content Security Policy:
    [\[CSP\]](references.html#refsCSP)

    - [[Content Security
      Policy](https://w3c.github.io/webappsec-csp/#content-security-policy-object)]{#content-security-policy
      .dfn}
    - [[disposition](https://w3c.github.io/webappsec-csp/#policy-disposition)]{#csp-disposition
      .dfn}
    - [[directive
      set](https://w3c.github.io/webappsec-csp/#policy-directive-set)]{#csp-directive-set
      .dfn}
    - [[Content Security Policy
      directive](https://w3c.github.io/webappsec-csp/#directives)]{#content-security-policy-directive
      .dfn}
    - [[CSP
      list](https://w3c.github.io/webappsec-csp/#csp-list)]{#concept-csp-list
      .dfn}
    - The [[Content Security Policy
      syntax](https://w3c.github.io/webappsec-csp/#grammardef-serialized-policy)]{#content-security-policy-syntax
      .dfn}
    - [[enforce the
      policy](https://w3c.github.io/webappsec-csp/#enforced)]{#enforce-the-policy
      .dfn}
    - The [[parse a serialized Content Security
      Policy](https://w3c.github.io/webappsec-csp/#parse-serialized-policy)]{#parse-a-serialized-content-security-policy
      .dfn} algorithm
    - The [[Run CSP initialization for a
      Document](https://w3c.github.io/webappsec-csp/#run-document-csp-initialization)]{#run-csp-initialization-for-a-document
      .dfn} algorithm
    - The [[Run CSP initialization for a global
      object](https://w3c.github.io/webappsec-csp/#run-global-object-csp-initialization)]{#run-csp-initialization-for-a-global-object
      .dfn} algorithm
    - The [[Should element\'s inline behavior be blocked by Content
      Security
      Policy?](https://w3c.github.io/webappsec-csp/#should-block-inline)]{#should-element's-inline-behavior-be-blocked-by-content-security-policy
      .dfn} algorithm
    - The [[Should navigation request of type be blocked by Content
      Security
      Policy?](https://w3c.github.io/webappsec-csp/#should-block-navigation-request)]{#should-navigation-request-of-type-be-blocked-by-content-security-policy
      .dfn} algorithm
    - The [[Should navigation response to navigation request of type in
      target be blocked by Content Security
      Policy?](https://w3c.github.io/webappsec-csp/#should-block-navigation-response)]{#should-navigation-response-to-navigation-request-of-type-in-target-be-blocked-by-content-security-policy
      .dfn} algorithm
    - The [[`report-uri`
      directive](https://w3c.github.io/webappsec-csp/#report-uri)]{#report-uri-directive
      .dfn}
    - The
      [[EnsureCSPDoesNotBlockStringCompilation](https://w3c.github.io/webappsec-csp/#can-compile-strings)]{#csp-ensurecspdoesnotblockstringcompilation
      .dfn} abstract operation
    - The [[Is base allowed for
      Document?](https://w3c.github.io/webappsec-csp/#allow-base-for-document)]{#is-base-allowed-for-document
      .dfn} algorithm
    - The [[`frame-ancestors`
      directive](https://w3c.github.io/webappsec-csp/#frame-ancestors)]{#frame-ancestors-directive
      .dfn}
    - The [[`sandbox`
      directive](https://w3c.github.io/webappsec-csp/#sandbox)]{#sandbox-directive
      .dfn}
    - The [[contains a header-delivered Content Security
      Policy](https://w3c.github.io/webappsec-csp/#contains-a-header-delivered-content-security-policy)]{#contains-a-header-delivered-content-security-policy
      .dfn} property.
    - The [[Parse a response\'s Content Security
      Policies](https://w3c.github.io/webappsec-csp/#parse-response-csp)]{#parse-response-csp
      .dfn} algorithm.
    - [[`SecurityPolicyViolationEvent`](https://w3c.github.io/webappsec-csp/#securitypolicyviolationevent)]{#securitypolicyviolationevent
      .dfn} interface
    - The
      [[`securitypolicyviolation`](https://w3c.github.io/webappsec-csp/#eventdef-globaleventhandlers-securitypolicyviolation)]{#event-securitypolicyviolation
      .dfn} event

Service Workers

:   The following terms are defined in Service Workers:
    [\[SW\]](references.html#refsSW)

    - [[active
      worker](https://w3c.github.io/ServiceWorker/#dfn-active-worker)]{#dfn-active-worker
      .dfn}
    - [[client message
      queue](https://w3c.github.io/ServiceWorker/#dfn-client-message-queue)]{#dfn-client-message-queue
      .dfn}
    - [[control](https://w3c.github.io/ServiceWorker/#dfn-control)]{#dfn-control
      .dfn}
    - [[handle
      fetch](https://w3c.github.io/ServiceWorker/#on-fetch-request-algorithm)]{#on-fetch-request-algorithm
      .dfn}
    - [[match service worker
      registration](https://w3c.github.io/ServiceWorker/#scope-match-algorithm)]{#scope-match-algorithm
      .dfn}
    - [[service
      worker](https://w3c.github.io/ServiceWorker/#dfn-service-worker)]{#dfn-service-worker
      .dfn}
    - [[service worker
      client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client)]{#serviceworkercontainer-service-worker-client
      .dfn}
    - [[service worker
      registration](https://w3c.github.io/service-workers/#dfn-service-worker-registration)]{#service-worker-registration
      .dfn}
    - [[`ServiceWorker`](https://w3c.github.io/ServiceWorker/#serviceworker)]{#serviceworker
      .dfn} interface
    - [[`ServiceWorkerContainer`](https://w3c.github.io/ServiceWorker/#serviceworkercontainer)]{#serviceworkercontainer
      .dfn} interface
    - [[`ServiceWorkerGlobalScope`](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope)]{#serviceworkerglobalscope
      .dfn} interface
    - [[`unregister`](https://w3c.github.io/service-workers/#navigator-service-worker-unregister)]{#service-worker-unregister
      .dfn}

Secure Contexts

:   The following algorithms are defined in Secure Contexts:
    [\[SECURE-CONTEXTS\]](references.html#refsSECURE-CONTEXTS)

    - [[Is url potentially
      trustworthy?](https://w3c.github.io/webappsec-secure-contexts/#potentially-trustworthy-url)]{#is-url-potentially-trustworthy
      .dfn}

Permissions Policy

:   The following terms are defined in Permissions Policy:
    [\[PERMISSIONSPOLICY\]](references.html#refsPERMISSIONSPOLICY)

    - [[permissions
      policy](https://w3c.github.io/webappsec-feature-policy/#permissions-policy)]{#concept-permissions-policy
      .dfn}
    - [[policy-controlled
      feature](https://w3c.github.io/webappsec-feature-policy/#policy-controlled-feature)]{#concept-policy-controlled-feature
      .dfn}
    - [[container
      policy](https://w3c.github.io/webappsec-feature-policy/#container-policy)]{#concept-container-policy
      .dfn}
    - [[serialized permissions
      policy](https://w3c.github.io/webappsec-feature-policy/#serialized-permissions-policy)]{#concept-serialized-permissions-policy
      .dfn}
    - [[default
      allowlist](https://w3c.github.io/webappsec-feature-policy/#default-allowlist)]{#concept-default-allowlist
      .dfn}
    - The [[creating a permissions
      policy](https://w3c.github.io/webappsec-feature-policy/#create-for-navigable)]{#creating-a-permissions-policy
      .dfn} algorithm
    - The [[creating a permissions policy from a
      response](https://w3c.github.io/webappsec-feature-policy/#create-from-response)]{#creating-a-permissions-policy-from-a-response
      .dfn} algorithm
    - The [[is feature enabled by policy for
      origin](https://w3c.github.io/webappsec-feature-policy/#is-feature-enabled)]{#is-feature-enabled
      .dfn} algorithm
    - The [[process permissions policy
      attributes](https://w3c.github.io/webappsec-feature-policy/#process-permissions-policy-attributes)]{#process-permissions-policy-attributes
      .dfn} algorithm

Payment Request API

:   The following feature is defined in Payment Request API:
    [\[PAYMENTREQUEST\]](references.html#refsPAYMENTREQUEST)

    - [[`PaymentRequest`](https://w3c.github.io/payment-request/#dom-paymentrequest)]{#paymentrequest
      .dfn} interface

MathML

:   While support for MathML as a whole is not required by this
    specification (though it is encouraged, at least for web browsers),
    certain features depend upon small parts of MathML being
    implemented. [\[MATHML\]](references.html#refsMATHML)

    The following features are defined in Mathematical Markup Language
    (MathML):

    - [[MathML
      `annotation-xml`](https://w3c.github.io/mathml-core/#dfn-annotation-xml)]{#mathml-annotation-xml
      .dfn} element
    - [[MathML
      `math`](https://w3c.github.io/mathml-core/#the-top-level-math-element)]{#mathml-math
      .dfn} element
    - [[MathML
      `merror`](https://w3c.github.io/mathml-core/#error-message-merror)]{#mathml-merror
      .dfn} element
    - [[MathML
      `mi`](https://w3c.github.io/mathml-core/#the-mi-element)]{#mathml-mi
      .dfn} element
    - [[MathML
      `mn`](https://w3c.github.io/mathml-core/#number-mn)]{#mathml-mn
      .dfn} element
    - [[MathML
      `mo`](https://w3c.github.io/mathml-core/#operator-fence-separator-or-accent-mo)]{#mathml-mo
      .dfn} element
    - [[MathML
      `ms`](https://w3c.github.io/mathml-core/#string-literal-ms)]{#mathml-ms
      .dfn} element
    - [[MathML
      `mtext`](https://w3c.github.io/mathml-core/#text-mtext)]{#mathml-mtext
      .dfn} element

SVG

:   While support for SVG as a whole is not required by this
    specification (though it is encouraged, at least for web browsers),
    certain features depend upon parts of SVG being implemented.

    User agents that implement SVG must implement the SVG 2
    specification, and not any earlier revisions.

    The following features are defined in the SVG 2 specification:
    [\[SVG\]](references.html#refsSVG)

    - [[`SVGElement`](https://svgwg.org/svg2-draft/types.html#InterfaceSVGElement)]{#svgelement
      .dfn} interface
    - [[`SVGImageElement`](https://svgwg.org/svg2-draft/embedded.html#InterfaceSVGImageElement)]{#svgimageelement
      .dfn} interface
    - [[`SVGScriptElement`](https://svgwg.org/svg2-draft/interact.html#InterfaceSVGScriptElement)]{#svgscriptelement
      .dfn} interface
    - [[`SVGSVGElement`](https://svgwg.org/svg2-draft/struct.html#InterfaceSVGSVGElement)]{#svgsvgelement
      .dfn} interface
    - [[SVG
      `a`](https://svgwg.org/svg2-draft/linking.html#AElement)]{#svg-a
      .dfn} element
    - [[SVG
      `desc`](https://svgwg.org/svg2-draft/struct.html#DescElement)]{#svg-desc
      .dfn} element
    - [[SVG
      `foreignObject`](https://svgwg.org/svg2-draft/embedded.html#ForeignObjectElement)]{#svg-foreignobject
      .dfn} element
    - [[SVG
      `image`](https://svgwg.org/svg2-draft/embedded.html#ImageElement)]{#svg-image
      .dfn} element
    - [[SVG
      `script`](https://svgwg.org/svg2-draft/interact.html#ScriptElement)]{#svg-script
      .dfn} element
    - [[SVG
      `svg`](https://svgwg.org/svg2-draft/struct.html#SVGElement)]{#svg-svg
      .dfn} element
    - [[SVG
      `title`](https://svgwg.org/svg2-draft/struct.html#TitleElement)]{#svg-title
      .dfn} element
    - [[SVG
      `use`](https://svgwg.org/svg2-draft/struct.html#UseElement)]{#svg-use
      .dfn} element
    - [[SVG
      `text-rendering`](https://svgwg.org/svg2-draft/painting.html#TextRendering)]{#svg-text-rendering
      .dfn} property

Filter Effects

:   The following features are defined in Filter Effects:
    [\[FILTERS\]](references.html#refsFILTERS)

    - [[\<filter-value-list\>](https://drafts.fxtf.org/filter-effects/#typedef-filter-value-list)]{#filter-value-list
      .dfn}

Compositing

:   The following features are defined in Compositing and Blending:
    [\[COMPOSITE\]](references.html#refsCOMPOSITE)

    - [[\<blend-mode\>](https://drafts.fxtf.org/compositing/#ltblendmodegt)]{#blend-mode
      .dfn}
    - [[\<composite-mode\>](https://drafts.fxtf.org/compositing/#compositemode)]{#composite-mode
      .dfn}
    - [[source-over](https://drafts.fxtf.org/compositing/#porterduffcompositingoperators_srcover)]{#gcop-source-over
      .dfn}
    - [[copy](https://drafts.fxtf.org/compositing/#porterduffcompositingoperators_src)]{#gcop-copy
      .dfn}

Cooperative Scheduling of Background Tasks

:   The following features are defined in Cooperative Scheduling of
    Background Tasks:
    [\[REQUESTIDLECALLBACK\]](references.html#refsREQUESTIDLECALLBACK)

    - [[`requestIdleCallback()`](https://w3c.github.io/requestidlecallback/#the-requestidlecallback-method)]{#requestidlecallback()
      .dfn}
    - [[start an idle period
      algorithm](https://w3c.github.io/requestidlecallback/#start-an-idle-period-algorithm)]{#start-an-idle-period-algorithm
      .dfn}

Screen Orientation

:   The following terms are defined in Screen Orientation:
    [\[SCREENORIENTATION\]](references.html#refsSCREENORIENTATION)

    - [[screen orientation change
      steps](https://w3c.github.io/screen-orientation/#dfn-screen-orientation-change-steps)]{#screen-orientation-change-steps
      .dfn}

Storage

:   The following terms are defined in Storage:
    [\[STORAGE\]](references.html#refsSTORAGE)

    - [[obtain a local storage bottle
      map](https://storage.spec.whatwg.org/#obtain-a-local-storage-bottle-map)]{#obtain-a-local-storage-bottle-map
      .dfn}
    - [[obtain a session storage bottle
      map](https://storage.spec.whatwg.org/#obtain-a-session-storage-bottle-map)]{#obtain-a-session-storage-bottle-map
      .dfn}
    - [[obtain a storage key for non-storage
      purposes](https://storage.spec.whatwg.org/#obtain-a-storage-key-for-non-storage-purposes)]{#obtain-a-storage-key-for-non-storage-purposes
      .dfn}
    - [[storage key
      equal](https://storage.spec.whatwg.org/#storage-key-equal)]{#storage-key-equal
      .dfn}
    - [[storage proxy
      map](https://storage.spec.whatwg.org/#storage-proxy-map)]{#storage-proxy-map
      .dfn}
    - [[legacy-clone a traversable storage
      shed](https://storage.spec.whatwg.org/#legacy-clone-a-traversable-storage-shed)]{#legacy-clone-a-traversable-storage-shed
      .dfn}

Web App Manifest

:   The following features are defined in Web App Manifest:
    [\[MANIFEST\]](references.html#refsMANIFEST)

    - [[application
      manifest](https://w3c.github.io/manifest/#dfn-manifest)]{#application-manifest
      .dfn}
    - [[installed web
      application](https://w3c.github.io/manifest/#dfn-installed-web-application)]{#installed-web-application
      .dfn}
    - [[process the
      manifest](https://w3c.github.io/manifest/#dfn-processing-a-manifest)]{#process-the-manifest
      .dfn}

WebAssembly JavaScript Interface: ESM Integration

:   The following terms are defined in WebAssembly JavaScript Interface:
    ESM Integration: [\[WASMESM\]](references.html#refsWASMESM)

    - [[WebAssembly Module
      Record](https://webassembly.github.io/esm-integration/js-api/index.html#webassembly-module-record)]{#webassembly-module-record
      .dfn}
    - [[parse a WebAssembly
      module](https://webassembly.github.io/esm-integration/js-api/index.html#parse-a-webassembly-module)]{#parse-a-webassembly-module
      .dfn}

WebCodecs

:   The following features are defined in WebCodecs:
    [\[WEBCODECS\]](references.html#refsWEBCODECS)

    - [[`VideoFrame`](https://w3c.github.io/webcodecs/#videoframe-interface)]{#videoframe
      .dfn} interface.
    - [[\[\[display
      width\]\]](https://w3c.github.io/webcodecs/#dom-videoframe-display-width-slot)]{#display-width
      .dfn}
    - [[\[\[display
      height\]\]](https://w3c.github.io/webcodecs/#dom-videoframe-display-height-slot)]{#display-height
      .dfn}

WebDriver

:   The following terms are defined in WebDriver:
    [\[WEBDRIVER\]](references.html#refsWEBDRIVER)

    - [[extension
      command](https://w3c.github.io/webdriver/#dfn-extension-commands)]{#extension-command
      .dfn}
    - [[remote end
      steps](https://w3c.github.io/webdriver/#dfn-remote-end-steps)]{#remote-end-steps
      .dfn}
    - [[WebDriver
      error](https://w3c.github.io/webdriver/#dfn-errors)]{#webdriver-error
      .dfn}
    - [[WebDriver error
      code](https://w3c.github.io/webdriver/#dfn-error-code)]{#webdriver-error-code
      .dfn}
    - [[invalid
      argument](https://w3c.github.io/webdriver/#dfn-invalid-argument)]{#invalid-argument
      .dfn}
    - [[getting a
      property](https://w3c.github.io/webdriver/#dfn-getting-properties)]{#getting-a-property
      .dfn}
    - [[success](https://w3c.github.io/webdriver/#dfn-success)]{#success-value
      .dfn}
    - [[WebDriver\'s security
      considerations](https://w3c.github.io/webdriver/#security)]{#webdriver's-security-considerations
      .dfn}
    - [[current browsing
      context](https://w3c.github.io/webdriver/#dfn-current-browsing-context)]{#webdriver-current-browsing-context
      .dfn}

WebDriver BiDi

:   The following terms are defined in WebDriver BiDi:
    [\[WEBDRIVERBIDI\]](references.html#refsWEBDRIVERBIDI)

    - [[WebDriver BiDi navigation
      status](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-status)]{#webdriver-bidi-navigation-status
      .dfn}
    - [[navigation status
      id](https://w3c.github.io/webdriver-bidi/#navigation-status-id)]{#navigation-status-id
      .dfn}
    - [[navigation status
      status](https://w3c.github.io/webdriver-bidi/#navigation-status-status)]{#navigation-status-status
      .dfn}
    - [[navigation status
      canceled](https://w3c.github.io/webdriver-bidi/#navigation-status-canceled)]{#navigation-status-canceled
      .dfn}
    - [[navigation status
      committed](https://w3c.github.io/webdriver-bidi/#navigation-status-committed)]{#navigation-status-committed
      .dfn}
    - [[navigation status
      pending](https://w3c.github.io/webdriver-bidi/#navigation-status-pending)]{#navigation-status-pending
      .dfn}
    - [[navigation status
      complete](https://w3c.github.io/webdriver-bidi/#navigation-status-complete)]{#navigation-status-complete
      .dfn}
    - [[navigation status
      url](https://w3c.github.io/webdriver-bidi/#navigation-status-url)]{#navigation-status-url
      .dfn}
    - [[navigation status suggested
      filename](https://w3c.github.io/webdriver-bidi/#navigation-status-suggested-filename)]{#navigation-status-suggested-filename
      .dfn}
    - [[WebDriver BiDi navigation
      aborted](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-aborted)]{#webdriver-bidi-navigation-aborted
      .dfn}
    - [[WebDriver BiDi navigation
      committed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-committed)]{#webdriver-bidi-navigation-committed
      .dfn}
    - [[WebDriver BiDi navigation
      failed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-failed)]{#webdriver-bidi-navigation-failed
      .dfn}
    - [[WebDriver BiDi navigation
      started](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigation-started)]{#webdriver-bidi-navigation-started
      .dfn}
    - [[WebDriver BiDi download
      started](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-download-started)]{#webdriver-bidi-download-started
      .dfn}
    - [[WebDriver BiDi fragment
      navigated](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-fragment-navigated)]{#webdriver-bidi-fragment-navigated
      .dfn}
    - [[WebDriver BiDi DOM content
      loaded](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-dom-content-loaded)]{#webdriver-bidi-dom-content-loaded
      .dfn}
    - [[WebDriver BiDi load
      complete](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-load-complete)]{#webdriver-bidi-load-complete
      .dfn}
    - [[WebDriver BiDi history
      updated](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-history-updated)]{#webdriver-bidi-history-updated
      .dfn}
    - [[WebDriver BiDi navigable
      created](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigable-created)]{#webdriver-bidi-navigable-created
      .dfn}
    - [[WebDriver BiDi navigable
      destroyed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigable-destroyed)]{#webdriver-bidi-navigable-destroyed
      .dfn}
    - [[WebDriver BiDi user prompt
      closed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-closed)]{#webdriver-bidi-user-prompt-closed
      .dfn}
    - [[WebDriver BiDi user prompt
      opened](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-user-prompt-opened)]{#webdriver-bidi-user-prompt-opened
      .dfn}
    - [[WebDriver BiDi file dialog
      opened](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-file-dialog-opened)]{#webdriver-bidi-file-dialog-opened
      .dfn}

Web Cryptography API

:   The following terms are defined in Web Cryptography API:
    [\[WEBCRYPTO\]](references.html#refsWEBCRYPTO)

    - [[generating a random
      UUID](https://w3c.github.io/webcrypto/#dfn-generate-a-random-uuid)]{#generating-a-random-uuid
      .dfn}

WebSockets

:   The following terms are defined in WebSockets:
    [\[WEBSOCKETS\]](references.html#refsWEBSOCKETS)

    - [[`WebSocket`](https://websockets.spec.whatwg.org/#websocket)]{#websocket-2
      .dfn}
    - [[make
      disappear](https://websockets.spec.whatwg.org/#make-disappear)]{#make-disappear
      .dfn dfn-for="WebSocket"}

WebTransport

:   The following terms are defined in WebTransport:
    [\[WEBTRANSPORT\]](references.html#refsWEBTRANSPORT)

    - [[`WebTransport`](https://w3c.github.io/webtransport/#webtransport)]{#webtransport
      .dfn}
    - [[`context cleanup steps`](https://w3c.github.io/webtransport/#context-cleanup-steps)]{#context-cleanup-steps
      .dfn}

Web Authentication: An API for accessing Public Key Credentials

:   The following terms are defined in Web Authentication: An API for
    accessing Public Key Credentials:
    [\[WEBAUTHN\]](references.html#refsWEBAUTHN)

    - [[public key
      credential](https://w3c.github.io/webauthn/#public-key-credential)]{#public-key-credential
      .dfn}

Credential Management

:   The following terms are defined in Credential Management:
    [\[CREDMAN\]](references.html#refsCREDMAN)

    - [[conditional
      mediation](https://w3c.github.io/webappsec-credential-management/#dom-credentialmediationrequirement-conditional)]{#conditional-mediation
      .dfn}
    - [[credential](https://w3c.github.io/webappsec-credential-management/#credential)]{#credman-credential
      .dfn}
    - [[`navigator.credentials.get()`](https://w3c.github.io/webappsec-credential-management/#dom-credentialscontainer-get)]{#navigator.credentials.get()
      .dfn}

Console

:   The following terms are defined in Console:
    [\[CONSOLE\]](references.html#refsCONSOLE)

    - [[report a warning to the
      console](https://console.spec.whatwg.org/#report-a-warning-to-the-console)]{#report-a-warning-to-the-console
      .dfn}

Web Locks API

:   The following terms are defined in Web Locks API:
    [\[WEBLOCKS\]](references.html#refsWEBLOCKS)

    - [[locks](https://w3c.github.io/web-locks/#lock-concept)]{#locks
      .dfn}
    - [[lock
      requests](https://w3c.github.io/web-locks/#lock-request)]{#lock-requests
      .dfn}

Trusted Types

:   This specification uses the following features defined in Trusted
    Types: [\[TRUSTED-TYPES\]](references.html#refsTRUSTED-TYPES)

    - [[`TrustedHTML`](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml)]{#tt-trustedhtml
      .dfn}
    - [[data](https://w3c.github.io/trusted-types/dist/spec/#trustedhtml-data)]{#tt-trustedhtml-data
      .dfn}
    - [[`TrustedScript`](https://w3c.github.io/trusted-types/dist/spec/#trusted-script)]{#tt-trustedscript
      .dfn}
    - [[`data`](https://w3c.github.io/trusted-types/dist/spec/#trustedscript-data)]{#tt-trustedscript-data
      .dfn}
    - [[`TrustedScriptURL`](https://w3c.github.io/trusted-types/dist/spec/#trustedscripturl)]{#tt-trustedscripturl
      .dfn}
    - [[Get Trusted Type compliant
      string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm)]{#tt-getcompliantstring
      .dfn}

WebRTC API

:   The following terms are defined in WebRTC API:
    [\[WEBRTC\]](references.html#refsWEBRTC)

    - [[`RTCDataChannel`](https://w3c.github.io/webrtc-pc/#dom-rtcdatachannel)]{#rtcdatachannel
      .dfn}
    - [[`RTCPeerConnection`](https://w3c.github.io/webrtc-pc/#dom-rtcpeerconnection)]{#rtcpeerconnection
      .dfn}

Picture-in-Picture API

:   The following terms are defined in Picture-in-Picture API:
    [\[PICTUREINPICTURE\]](references.html#refsPICTUREINPICTURE)

    - [[`PictureInPictureWindow`](https://w3c.github.io/picture-in-picture/#pictureinpicturewindow)]{#pictureinpicturewindow
      .dfn}

Idle Detection API

:   The following terms are defined in Idle Detection API:

    - [[`IdleDetector`](https://wicg.github.io/idle-detection/#idledetector)]{#idledetector
      .dfn}

Web Speech API

:   The following terms are defined in Web Speech API:

    - [[`SpeechRecognition`](https://wicg.github.io/speech-api/#speechrecognition)]{#speechrecognition
      .dfn}

WebOTP API

:   The following terms are defined in WebOTP API:

    - [[`OTPCredential`](https://wicg.github.io/web-otp/#otpcredential)]{#otpcredential
      .dfn}

Web Share API

:   The following terms are defined in Web Share API:

    - [[share()](https://w3c.github.io/web-share/#share-method)]{#dom-navigator-share
      .dfn}

Web Smart Card API

:   The following terms are defined in Web Smart Card API:

    - [[`SmartCardConnection`](https://wicg.github.io/web-smart-card/#dom-smartcardconnection)]{#smartcardconnection
      .dfn}

Web Background Synchronization

:   The following terms are defined in Web Background Synchronization:

    - [[`SyncManager`](https://wicg.github.io/background-sync/spec/#syncmanager)]{#syncmanager
      .dfn}
    - [[`register()`](https://wicg.github.io/background-sync/spec/#dom-syncmanager-register)]{#dom-syncmanager-register
      .dfn}

Web Periodic Background Synchronization

:   The following terms are defined in Web Periodic Background
    Synchronization:

    - [[`PeriodicSyncManager`](https://wicg.github.io/periodic-background-sync/#periodicsyncmanager)]{#periodicsyncmanager
      .dfn}
    - [[`register()`](https://wicg.github.io/periodic-background-sync/#dom-periodicsyncmanager-register)]{#dom-periodicsyncmanager-register
      .dfn}

Web Background Fetch

:   The following terms are defined in Background Fetch:

    - [[`BackgroundFetchManager`](https://wicg.github.io/background-fetch/#backgroundfetchmanager)]{#backgroundfetchmanager
      .dfn}
    - [[`fetch()`](https://wicg.github.io/background-fetch/#dom-backgroundfetchmanager-fetch)]{#dom-backgroundfetchmanager-fetch
      .dfn}

Keyboard Lock

:   The following terms are defined in Keyboard Lock:

    - [[`Keyboard`](https://wicg.github.io/keyboard-lock/#keyboard)]{#keyboard
      .dfn}
    - [[`lock()`](https://wicg.github.io/keyboard-lock/#dom-keyboard-lock)]{#dom-keyboard-lock
      .dfn}

Web MIDI API

:   The following terms are defined in Web MIDI API:

    - [[`requestMIDIAccess()`](https://webaudio.github.io/web-midi-api/#dom-navigator-requestmidiaccess)]{#dom-navigator-requestmidiaccess
      .dfn}

Generic Sensor API

:   The following terms are defined in Generic Sensor API:

    - [[request sensor
      access](https://w3c.github.io/sensors/#request-sensor-access)]{#request-sensor-access
      .dfn}

WebHID API

:   The following terms are defined in WebHID API:

    - [[`requestDevice`](https://wicg.github.io/webhid/#requestdevice-method)]{#request-device
      .dfn}

WebXR Device API

:   The following terms are defined in WebXR Device API:

    - [[`XRSystem`](https://immersive-web.github.io/webxr/#xrsystem)]{#xrsystem
      .dfn}

------------------------------------------------------------------------

This specification does not *require* support of any particular network
protocol, style sheet language, scripting language, or any of the DOM
specifications beyond those required in the list above. However, the
language described by this specification is biased towards CSS as the
styling language, JavaScript as the scripting language, and HTTP as the
network protocol, and several features assume that those languages and
protocols are in use.

A user agent that implements the HTTP protocol must implement HTTP State
Management Mechanism (Cookies) as well.
[\[HTTP\]](references.html#refsHTTP)
[\[COOKIES\]](references.html#refsCOOKIES)

This specification might have certain additional requirements on
character encodings, image formats, audio formats, and video formats in
the respective sections.
:::

#### [2.1.10]{.secno} Extensibility[](#extensibility-2){.self-link} {#extensibility-2}

Vendor-specific proprietary user agent extensions to this specification
are strongly discouraged. Documents must not use such extensions, as
doing so reduces interoperability and fragments the user base, allowing
only users of specific user agents to access the content in question.

All extensions must be defined so that the use of extensions neither
contradicts nor causes the non-conformance of functionality defined in
the specification.

::: example
For example, while strongly discouraged from doing so, an implementation
could add a new IDL attribute \"`typeTime`\" to a control that returned
the time it took the user to select the current value of a control
(say). On the other hand, defining a new control that appears in a
form\'s
[`elements`{#extensibility-2:dom-form-elements}](forms.html#dom-form-elements)
array would be in violation of the above requirement, as it would
violate the definition of
[`elements`{#extensibility-2:dom-form-elements-2}](forms.html#dom-form-elements)
given in this specification.
:::

------------------------------------------------------------------------

When vendor-neutral extensions to this specification are needed, either
this specification can be updated accordingly, or an extension
specification can be written that overrides the requirements in this
specification. When someone applying this specification to their
activities decides that they will recognize the requirements of such an
extension specification, it becomes an [applicable
specification]{#other-applicable-specifications .dfn} for the purposes
of conformance requirements in this specification.

Someone could write a specification that defines any arbitrary byte
stream as conforming, and then claim that their random junk is
conforming. However, that does not mean that their random junk actually
is conforming for everyone\'s purposes: if someone else decides that
that specification does not apply to their work, then they can quite
legitimately say that the aforementioned random junk is just that, junk,
and not conforming at all. As far as conformance goes, what matters in a
particular community is what that community *agrees* is applicable.

------------------------------------------------------------------------

User agents must treat elements and attributes that they do not
understand as semantically neutral; leaving them in the DOM (for DOM
processors), and styling them according to CSS (for CSS processors), but
not inferring any meaning from them.

When support for a feature is disabled (e.g. as an emergency measure to
mitigate a security problem, or to aid in development, or for
performance reasons), user agents must act as if they had no support for
the feature whatsoever, and as if the feature was not mentioned in this
specification. For example, if a particular feature is accessed via an
attribute in a Web IDL interface, the attribute itself would be omitted
from the objects that implement that interface --- leaving the attribute
on the object but making it return null or throw an exception is
insufficient.

#### [2.1.11]{.secno} Interactions with XPath and XSLT[](#interactions-with-xpath-and-xslt){.self-link}

Implementations of XPath 1.0 that operate on [HTML
documents](https://dom.spec.whatwg.org/#html-document){#interactions-with-xpath-and-xslt:html-documents
x-internal="html-documents"} parsed or created in the manners described
in this specification (e.g. as part of the `document.evaluate()` API)
must act as if the following edit was applied to the XPath 1.0
specification.

First, remove this paragraph:

> A [QName](https://www.w3.org/TR/REC-xml-names/#NT-QName) in the node
> test is expanded into an
> [expanded-name](https://www.w3.org/TR/1999/REC-xpath-19991116/#dt-expanded-name)
> using the namespace declarations from the expression context. This is
> the same way expansion is done for element type names in start and
> end-tags except that the default namespace declared with `xmlns` is
> not used: if the
> [QName](https://www.w3.org/TR/REC-xml-names/#NT-QName) does not have a
> prefix, then the namespace URI is null (this is the same way attribute
> names are expanded). It is an error if the
> [QName](https://www.w3.org/TR/REC-xml-names/#NT-QName) has a prefix
> for which there is no namespace declaration in the expression context.

Then, insert in its place the following:

> A QName in the node test is expanded into an expanded-name using the
> namespace declarations from the expression context. If the QName has a
> prefix, then there must be a namespace declaration for this prefix in
> the expression context, and the corresponding namespace URI is the one
> that is associated with this prefix. It is an error if the QName has a
> prefix for which there is no namespace declaration in the expression
> context.
>
> If the QName has no prefix and the principal node type of the axis is
> element, then the default element namespace is used. Otherwise, if the
> QName has no prefix, the namespace URI is null. The default element
> namespace is a member of the context for the XPath expression. The
> value of the default element namespace when executing an XPath
> expression through the DOM3 XPath API is determined in the following
> way:
>
> 1.  If the context node is from an HTML DOM, the default element
>     namespace is \"http://www.w3.org/1999/xhtml\".
> 2.  Otherwise, the default element namespace URI is null.
>
> This is equivalent to adding the default element namespace feature of
> XPath 2.0 to XPath 1.0, and using the HTML namespace as the default
> element namespace for HTML documents. It is motivated by the desire to
> have implementations be compatible with legacy HTML content while
> still supporting the changes that this specification introduces to
> HTML regarding the namespace used for HTML elements, and by the desire
> to use XPath 1.0 rather than XPath 2.0.

This change is a [willful
violation](https://infra.spec.whatwg.org/#willful-violation){#interactions-with-xpath-and-xslt:willful-violation
x-internal="willful-violation"} of the XPath 1.0 specification,
motivated by desire to have implementations be compatible with legacy
content while still supporting the changes that this specification
introduces to HTML regarding which namespace is used for HTML elements.
[\[XPATH10\]](references.html#refsXPATH10)

------------------------------------------------------------------------

XSLT 1.0 processors outputting to a DOM when the output method is
\"html\" (either explicitly or via the defaulting rule in XSLT 1.0) are
affected as follows:

If the transformation program outputs an element in no namespace, the
processor must, prior to constructing the corresponding DOM element
node, change the namespace of the element to the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#interactions-with-xpath-and-xslt:html-namespace-2
x-internal="html-namespace-2"},
[ASCII-lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#interactions-with-xpath-and-xslt:converted-to-ascii-lowercase
x-internal="converted-to-ascii-lowercase"} the element\'s local name,
and
[ASCII-lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#interactions-with-xpath-and-xslt:converted-to-ascii-lowercase-2
x-internal="converted-to-ascii-lowercase"} the names of any
non-namespaced attributes on the element.

This requirement is a [willful
violation](https://infra.spec.whatwg.org/#willful-violation){#interactions-with-xpath-and-xslt:willful-violation-2
x-internal="willful-violation"} of the XSLT 1.0 specification, required
because this specification changes the namespaces and case-sensitivity
rules of HTML in a manner that would otherwise be incompatible with
DOM-based XSLT transformations. (Processors that serialize the output
are unaffected.) [\[XSLT10\]](references.html#refsXSLT10)

------------------------------------------------------------------------

This specification does not specify precisely how XSLT processing
interacts with the [HTML
parser](parsing.html#html-parser){#interactions-with-xpath-and-xslt:html-parser}
infrastructure (for example, whether an XSLT processor acts as if it
puts any elements into a [stack of open
elements](parsing.html#stack-of-open-elements){#interactions-with-xpath-and-xslt:stack-of-open-elements}).
However, XSLT processors must [stop
parsing](parsing.html#stop-parsing){#interactions-with-xpath-and-xslt:stop-parsing}
if they successfully complete, and must [update the current document
readiness](dom.html#update-the-current-document-readiness){#interactions-with-xpath-and-xslt:update-the-current-document-readiness}
first to \"`interactive`\" and then to \"`complete`\" if they are
aborted.

------------------------------------------------------------------------

This specification does not specify how XSLT interacts with the
[navigation](browsing-the-web.html#navigate){#interactions-with-xpath-and-xslt:navigate}
algorithm, how it fits in with the [event
loop](webappapis.html#event-loop){#interactions-with-xpath-and-xslt:event-loop},
nor how error pages are to be handled (e.g. whether XSLT errors are to
replace an incremental XSLT output, or are rendered inline, etc.).

There are also additional non-normative comments regarding the
interaction of XSLT and HTML [in the `script` element
section](scripting.html#scriptTagXSLT), and of XSLT, XPath, and HTML [in
the `template` element section](scripting.html#template-XSLT-XPath).

### [2.2]{.secno} Policy-controlled features[](#policy-controlled-features){.self-link}

::::: {.mdn-anno .wrapped}
**⚠**MDN

:::: feature
[Headers/Permissions-Policy/document-domain](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/document-domain "The HTTP Permissions-Policy header document-domain directive controls whether the current document is allowed to set document.domain.")

Support in one engine only.

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[SafariNo]{.safari .no}[Chrome[🔰
88+]{title="Requires setting a user preference or runtime flag."}]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge[🔰
88+]{title="Requires setting a user preference or runtime flag."}]{.edge_blink
.yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

This document defines the following [policy-controlled
features](https://w3c.github.io/webappsec-feature-policy/#policy-controlled-feature){#policy-controlled-features:concept-policy-controlled-feature
x-internal="concept-policy-controlled-feature"}:

::::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Headers/Feature-Policy/autoplay](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Feature-Policy/autoplay "The HTTP Permissions-Policy header autoplay directive controls whether the current document is allowed to autoplay media requested through the HTMLMediaElement interface.")

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[SafariNo]{.safari .no}[Chrome64+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[Headers/Permissions-Policy/autoplay](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/autoplay "The HTTP Permissions-Policy header autoplay directive controls whether the current document is allowed to autoplay media requested through the HTMLMediaElement interface.")

Support in one engine only.

::: support
[Firefox[🔰
74+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[SafariNo]{.safari .no}[Chrome88+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge88+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

- \"[`autoplay`]{#autoplay-feature .dfn}\", which has a [default
  allowlist](https://w3c.github.io/webappsec-feature-policy/#default-allowlist){#policy-controlled-features:concept-default-allowlist
  x-internal="concept-default-allowlist"} of `'self'`.
- \"[`cross-origin-isolated`]{#cross-origin-isolated-feature .dfn}\",
  which has a [default
  allowlist](https://w3c.github.io/webappsec-feature-policy/#default-allowlist){#policy-controlled-features:concept-default-allowlist-2
  x-internal="concept-default-allowlist"} of `'self'`.
- \"[`focus-without-user-activation`]{#focus-without-user-activation-feature
  .dfn}\", which has a [default
  allowlist](https://w3c.github.io/webappsec-feature-policy/#default-allowlist){#policy-controlled-features:concept-default-allowlist-3
  x-internal="concept-default-allowlist"} of `'self'`.

[← 1 Introduction](introduction.html) --- [Table of Contents](./) ---
[2.3 Common microsyntaxes →](common-microsyntaxes.html)
