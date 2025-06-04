HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 7.2 APIs related to navigation and session
history](nav-history-apis.html) --- [Table of Contents](./) --- [7.4
Navigation and session history →](browsing-the-web.html)

1.  ::: {#toc-browsers}
    1.  [[7.3]{.secno} Infrastructure for sequences of
        documents](document-sequences.html#infrastructure-for-sequences-of-documents)
        1.  [[7.3.1]{.secno}
            Navigables](document-sequences.html#navigables)
            1.  [[7.3.1.1]{.secno} Traversable
                navigables](document-sequences.html#traversable-navigables)
            2.  [[7.3.1.2]{.secno} Top-level
                traversables](document-sequences.html#top-level-traversables)
            3.  [[7.3.1.3]{.secno} Child
                navigables](document-sequences.html#child-navigables)
            4.  [[7.3.1.4]{.secno} Jake
                diagrams](document-sequences.html#jake-diagrams)
            5.  [[7.3.1.5]{.secno} Related navigable
                collections](document-sequences.html#related-navigable-collections)
            6.  [[7.3.1.6]{.secno} Navigable
                destruction](document-sequences.html#garbage-collection-and-browsing-contexts)
            7.  [[7.3.1.7]{.secno} Navigable target
                names](document-sequences.html#navigable-target-names)
        2.  [[7.3.2]{.secno} Browsing
            contexts](document-sequences.html#windows)
            1.  [[7.3.2.1]{.secno} Creating browsing
                contexts](document-sequences.html#creating-browsing-contexts)
            2.  [[7.3.2.2]{.secno} Related browsing
                contexts](document-sequences.html#nested-browsing-contexts)
            3.  [[7.3.2.3]{.secno} Groupings of browsing
                contexts](document-sequences.html#groupings-of-browsing-contexts)
        3.  [[7.3.3]{.secno} Fully active
            documents](document-sequences.html#fully-active-documents)
    :::

### [7.3]{.secno} Infrastructure for sequences of documents[](#infrastructure-for-sequences-of-documents){.self-link}

This standard contains several related concepts for grouping sequences
of documents. As a brief, non-normative summary:

- [Navigables](#navigable){#infrastructure-for-sequences-of-documents:navigable}
  are a user-facing representation of a sequence of documents, i.e.,
  they represent something that can be navigated between documents.
  Typical examples are tabs or windows in a web browser, or
  [`iframe`{#infrastructure-for-sequences-of-documents:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s,
  or
  [`frame`{#infrastructure-for-sequences-of-documents:frame}](obsolete.html#frame)s
  in a
  [`frameset`{#infrastructure-for-sequences-of-documents:frameset}](obsolete.html#frameset).

- [Traversable
  navigables](#traversable-navigable){#infrastructure-for-sequences-of-documents:traversable-navigable}
  are a special type of navigable which control the session history of
  themselves and of their descendant navigables. That is, in addition to
  their own series of documents, they represent a tree of further series
  of documents, plus the ability to linearly traverse back and forward
  through a flattened view of this tree.

- [Browsing
  contexts](#browsing-context){#infrastructure-for-sequences-of-documents:browsing-context}
  are a developer-facing representation of a series of documents. They
  correspond 1:1 with
  [`WindowProxy`{#infrastructure-for-sequences-of-documents:windowproxy}](nav-history-apis.html#windowproxy)
  objects. Each navigable can present a series of browsing contexts,
  with
  [switches](browsers.html#browsing-context-group-switches-due-to-cross-origin-opener-policy)
  between those browsing contexts occuring under certain well-defined
  circumstances.

Most of this standard works in the language of navigables, but certain
APIs expose the existence of browsing context switches, and so some
parts of the standard need to work in terms of browsing contexts.

#### [7.3.1]{.secno} Navigables[](#navigables){.self-link}

A [navigable]{#navigable .dfn export=""} presents a
[`Document`{#navigables:document}](dom.html#document) to the user via
its [active session history
entry](#nav-active-history-entry){#navigables:nav-active-history-entry}.
Each navigable has:

- An [id]{#nav-id .dfn}, a [new unique internal
  value](common-microsyntaxes.html#new-unique-internal-value){#navigables:new-unique-internal-value}.

- A [parent]{#nav-parent .dfn dfn-for="navigable" export=""}, a
  [navigable](#navigable){#navigables:navigable} or null.

- A [current session history entry]{#nav-current-history-entry .dfn}, a
  [session history
  entry](browsing-the-web.html#session-history-entry){#navigables:session-history-entry}.

  This can only be modified within the [session history traversal
  queue](#tn-session-history-traversal-queue){#navigables:tn-session-history-traversal-queue}
  of the parent [traversable
  navigable](#traversable-navigable){#navigables:traversable-navigable}.

- An [active session history entry]{#nav-active-history-entry .dfn}, a
  [session history
  entry](browsing-the-web.html#session-history-entry){#navigables:session-history-entry-2}.

  This can only be modified from the event loop of the [active session
  history
  entry](#nav-active-history-entry){#navigables:nav-active-history-entry-2}\'s
  [document](browsing-the-web.html#she-document){#navigables:she-document}.

- An [is closing]{#is-closing .dfn} boolean, initially false.

  This is only ever set to true for [top-level traversable
  navigables](#top-level-traversable){#navigables:top-level-traversable}.

- An [is delaying `load` events]{#delaying-load-events-mode .dfn}
  boolean, initially false.

  This is only ever set to true in cases where the navigable\'s
  [parent](#nav-parent){#navigables:nav-parent} is non-null.

The [current session history
entry](#nav-current-history-entry){#navigables:nav-current-history-entry}
and the [active session history
entry](#nav-active-history-entry){#navigables:nav-active-history-entry-3}
are usually the same, but they get out of sync when:

- Synchronous navigations are performed. This causes the [active session
  history
  entry](#nav-active-history-entry){#navigables:nav-active-history-entry-4}
  to temporarily step ahead of the [current session history
  entry](#nav-current-history-entry){#navigables:nav-current-history-entry-2}.

- A non-displayable, non-error response is received when [applying the
  history
  step](browsing-the-web.html#apply-the-history-step){#navigables:apply-the-history-step}.
  This updates the [current session history
  entry](#nav-current-history-entry){#navigables:nav-current-history-entry-3}
  but leaves the [active session history
  entry](#nav-active-history-entry){#navigables:nav-active-history-entry-5}
  as-is.

------------------------------------------------------------------------

A [navigable](#navigable){#navigables:navigable-2}\'s [active
document]{#nav-document .dfn dfn-for="navigable" export=""} is its
[active session history
entry](#nav-active-history-entry){#navigables:nav-active-history-entry-6}\'s
[document](browsing-the-web.html#she-document){#navigables:she-document-2}.

This can be safely read from within the [session history traversal
queue](#tn-session-history-traversal-queue){#navigables:tn-session-history-traversal-queue-2}
of the navigable\'s [top-level
traversable](#nav-top){#navigables:nav-top}. Although a
[navigable](#navigable){#navigables:navigable-3}\'s [active history
entry](#nav-active-history-entry){#navigables:nav-active-history-entry-7}
can change synchronously, the new entry will always have the same
[`Document`{#navigables:document-2}](dom.html#document).

A [navigable](#navigable){#navigables:navigable-4}\'s [active browsing
context]{#nav-bc .dfn dfn-for="navigable" export=""} is its [active
document](#nav-document){#navigables:nav-document}\'s [browsing
context](#concept-document-bc){#navigables:concept-document-bc}. If this
[navigable](#navigable){#navigables:navigable-5} is a [traversable
navigable](#traversable-navigable){#navigables:traversable-navigable-2},
then its [active browsing context](#nav-bc){#navigables:nav-bc} will be
a [top-level browsing
context](#top-level-browsing-context){#navigables:top-level-browsing-context}.

A [navigable](#navigable){#navigables:navigable-6}\'s [active
`WindowProxy`]{#nav-wp .dfn} is its [active browsing
context](#nav-bc){#navigables:nav-bc-2}\'s associated
[`WindowProxy`{#navigables:windowproxy}](nav-history-apis.html#windowproxy).

A [navigable](#navigable){#navigables:navigable-7}\'s [active
window]{#nav-window .dfn dfn-for="navigable" export=""} is its [active
`WindowProxy`](#nav-wp){#navigables:nav-wp}\'s
[\[\[Window\]\]](nav-history-apis.html#concept-windowproxy-window){#navigables:concept-windowproxy-window}.

This will always equal the navigable\'s [active
document](#nav-document){#navigables:nav-document-2}\'s [relevant global
object](webappapis.html#concept-relevant-global){#navigables:concept-relevant-global};
this is kept in sync by the [make
active](browsing-the-web.html#make-active){#navigables:make-active}
algorithm.

A [navigable](#navigable){#navigables:navigable-8}\'s [target
name]{#nav-target .dfn} is its [active session history
entry](#nav-active-history-entry){#navigables:nav-active-history-entry-8}\'s
[document
state](browsing-the-web.html#she-document-state){#navigables:she-document-state}\'s
[navigable target
name](browsing-the-web.html#document-state-nav-target-name){#navigables:document-state-nav-target-name}.

------------------------------------------------------------------------

To get the [node navigable]{#node-navigable .dfn export=""} of a
[node](https://dom.spec.whatwg.org/#interface-node){#navigables:node
x-internal="node"} `node`{.variable}, return the
[navigable](#navigable){#navigables:navigable-9} whose [active
document](#nav-document){#navigables:nav-document-3} is
`node`{.variable}\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#navigables:node-document
x-internal="node-document"}, or null if there is no such
[navigable](#navigable){#navigables:navigable-10}.

------------------------------------------------------------------------

To [initialize the navigable]{#initialize-the-navigable .dfn}
[navigable](#navigable){#navigables:navigable-11}
`navigable`{.variable}, given a [document
state](browsing-the-web.html#document-state-2){#navigables:document-state-2}
`documentState`{.variable} and an optional
[navigable](#navigable){#navigables:navigable-12}-or-null
`parent`{.variable} (default null):

1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigables:assert
    x-internal="assert"}: `documentState`{.variable}\'s
    [document](browsing-the-web.html#document-state-document){#navigables:document-state-document}
    is non-null.

2.  Let `entry`{.variable} be a new [session history
    entry](browsing-the-web.html#session-history-entry){#navigables:session-history-entry-3},
    with

    [URL](browsing-the-web.html#she-url){#navigables:she-url}
    :   `documentState`{.variable}\'s
        [document](browsing-the-web.html#document-state-document){#navigables:document-state-document-2}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#navigables:the-document's-address
        x-internal="the-document's-address"}

    [document state](browsing-the-web.html#she-document-state){#navigables:she-document-state-2}
    :   `documentState`{.variable}

    The caller of this algorithm is responsible for initializing
    `entry`{.variable}\'s
    [step](browsing-the-web.html#she-step){#navigables:she-step}; it
    will be left as \"`pending`\" until that is complete.

3.  Set `navigable`{.variable}\'s [current session history
    entry](#nav-current-history-entry){#navigables:nav-current-history-entry-4}
    to `entry`{.variable}.

4.  Set `navigable`{.variable}\'s [active session history
    entry](#nav-active-history-entry){#navigables:nav-active-history-entry-9}
    to `entry`{.variable}.

5.  Set `navigable`{.variable}\'s
    [parent](#nav-parent){#navigables:nav-parent-2} to
    `parent`{.variable}.

##### [7.3.1.1]{.secno} Traversable navigables[](#traversable-navigables){.self-link}

A [traversable navigable]{#traversable-navigable .dfn export=""} is a
[navigable](#navigable){#traversable-navigables:navigable} that also
controls which [session history
entry](browsing-the-web.html#session-history-entry){#traversable-navigables:session-history-entry}
should be the [current session history
entry](#nav-current-history-entry){#traversable-navigables:nav-current-history-entry}
and [active session history
entry](#nav-active-history-entry){#traversable-navigables:nav-active-history-entry}
for itself and its descendant
[navigables](#navigable){#traversable-navigables:navigable-2}.

In addition to the properties of a
[navigable](#navigable){#traversable-navigables:navigable-3}, a
[traversable
navigable](#traversable-navigable){#traversable-navigables:traversable-navigable}
has:

- A [current session history step]{#tn-current-session-history-step .dfn
  export=""}, a number, initially 0.

- [Session history entries]{#tn-session-history-entries .dfn}, a
  [list](https://infra.spec.whatwg.org/#list){#traversable-navigables:list
  x-internal="list"} of [session history
  entries](browsing-the-web.html#session-history-entry){#traversable-navigables:session-history-entry-2},
  initially a new
  [list](https://infra.spec.whatwg.org/#list){#traversable-navigables:list-2
  x-internal="list"}.

- A [session history traversal
  queue]{#tn-session-history-traversal-queue .dfn}, a [session history
  traversal parallel
  queue](browsing-the-web.html#session-history-traversal-parallel-queue){#traversable-navigables:session-history-traversal-parallel-queue},
  the result of [starting a new session history traversal parallel
  queue](browsing-the-web.html#starting-a-new-session-history-traversal-parallel-queue){#traversable-navigables:starting-a-new-session-history-traversal-parallel-queue}.

- A [running nested apply history
  step]{#tn-running-nested-apply-history-step .dfn} boolean, initially
  false.

- A [system visibility state]{#system-visibility-state .dfn}, which is
  either \"`hidden`\" or \"`visible`\".

  See the [page visibility](interaction.html#page-visibility) section
  for the requirements on this item.

- An [is created by web content]{#is-created-by-web-content .dfn}
  boolean, initially false.

To get the [traversable navigable]{#nav-traversable .dfn
dfn-for="navigable" export=""} of a
[navigable](#navigable){#traversable-navigables:navigable-4}
`inputNavigable`{.variable}:

1.  Let `navigable`{.variable} be `inputNavigable`{.variable}.

2.  While `navigable`{.variable} is not a [traversable
    navigable](#traversable-navigable){#traversable-navigables:traversable-navigable-2},
    set `navigable`{.variable} to `navigable`{.variable}\'s
    [parent](#nav-parent){#traversable-navigables:nav-parent}.

3.  Return `navigable`{.variable}.

##### [7.3.1.2]{.secno} Top-level traversables[](#top-level-traversables){.self-link}

A [top-level traversable]{#top-level-traversable .dfn export=""} is a
[traversable
navigable](#traversable-navigable){#top-level-traversables:traversable-navigable}
with a null [parent](#nav-parent){#top-level-traversables:nav-parent}.

Currently, all [traversable
navigables](#traversable-navigable){#top-level-traversables:traversable-navigable-2}
are [top-level
traversables](#top-level-traversable){#top-level-traversables:top-level-traversable}.
Future proposals envision introducing non-top-level traversables.

A user agent holds a [top-level traversable
set]{#top-level-traversable-set .dfn dfn-for="user agent" export=""} (a
[set](https://infra.spec.whatwg.org/#ordered-set){#top-level-traversables:set
x-internal="set"} of [top-level
traversables](#top-level-traversable){#top-level-traversables:top-level-traversable-2}).
These are typically presented to the user in the form of browser windows
or browser tabs.

To get the [top-level traversable]{#nav-top .dfn dfn-for="navigable"
export=""} of a
[navigable](#navigable){#top-level-traversables:navigable}
`inputNavigable`{.variable}:

1.  Let `navigable`{.variable} be `inputNavigable`{.variable}.

2.  While `navigable`{.variable}\'s
    [parent](#nav-parent){#top-level-traversables:nav-parent-2} is not
    null, set `navigable`{.variable} to `navigable`{.variable}\'s
    [parent](#nav-parent){#top-level-traversables:nav-parent-3}.

3.  Return `navigable`{.variable}.

To [create a new top-level
traversable]{#creating-a-new-top-level-traversable .dfn} given a
[browsing
context](#browsing-context){#top-level-traversables:browsing-context}-or-null
`opener`{.variable}, a string `targetName`{.variable}, and an optional
[navigable](#navigable){#top-level-traversables:navigable-2}
`openerNavigableForWebDriver`{.variable}:

1.  Let `document`{.variable} be null.

2.  If `opener`{.variable} is null, then set `document`{.variable} to
    the second return value of [creating a new top-level browsing
    context and
    document](#creating-a-new-top-level-browsing-context){#top-level-traversables:creating-a-new-top-level-browsing-context}.

3.  Otherwise, set `document`{.variable} to the second return value of
    [creating a new auxiliary browsing context and
    document](#creating-a-new-auxiliary-browsing-context){#top-level-traversables:creating-a-new-auxiliary-browsing-context}
    given `opener`{.variable}.

4.  Let `documentState`{.variable} be a new [document
    state](browsing-the-web.html#document-state-2){#top-level-traversables:document-state-2},
    with

    [document](browsing-the-web.html#document-state-document){#top-level-traversables:document-state-document}
    :   `document`{.variable}

    [initiator origin](browsing-the-web.html#document-state-initiator-origin){#top-level-traversables:document-state-initiator-origin}
    :   null if `opener`{.variable} is null; otherwise,
        `document`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#top-level-traversables:concept-document-origin
        x-internal="concept-document-origin"}

    [origin](browsing-the-web.html#document-state-origin){#top-level-traversables:document-state-origin}
    :   `document`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#top-level-traversables:concept-document-origin-2
        x-internal="concept-document-origin"}

    [navigable target name](browsing-the-web.html#document-state-nav-target-name){#top-level-traversables:document-state-nav-target-name}
    :   `targetName`{.variable}

    [about base URL](browsing-the-web.html#document-state-about-base-url){#top-level-traversables:document-state-about-base-url}
    :   `document`{.variable}\'s [about base
        URL](dom.html#concept-document-about-base-url){#top-level-traversables:concept-document-about-base-url}

5.  Let `traversable`{.variable} be a new [traversable
    navigable](#traversable-navigable){#top-level-traversables:traversable-navigable-3}.

6.  [Initialize the
    navigable](#initialize-the-navigable){#top-level-traversables:initialize-the-navigable}
    `traversable`{.variable} given `documentState`{.variable}.

7.  Let `initialHistoryEntry`{.variable} be `traversable`{.variable}\'s
    [active session history
    entry](#nav-active-history-entry){#top-level-traversables:nav-active-history-entry}.

8.  Set `initialHistoryEntry`{.variable}\'s
    [step](browsing-the-web.html#she-step){#top-level-traversables:she-step}
    to 0.

9.  [Append](https://infra.spec.whatwg.org/#list-append){#top-level-traversables:list-append
    x-internal="list-append"} `initialHistoryEntry`{.variable} to
    `traversable`{.variable}\'s [session history
    entries](#tn-session-history-entries){#top-level-traversables:tn-session-history-entries}.

10. [If `opener`{.variable} is non-null, then [legacy-clone a
    traversable storage
    shed](https://storage.spec.whatwg.org/#legacy-clone-a-traversable-storage-shed){#top-level-traversables:legacy-clone-a-traversable-storage-shed
    x-internal="legacy-clone-a-traversable-storage-shed"} given
    `opener`{.variable}\'s [top-level
    traversable](#bc-traversable){#top-level-traversables:bc-traversable}
    and `traversable`{.variable}.
    [\[STORAGE\]](references.html#refsSTORAGE)]{#copy-session-storage}

11. [Append](https://infra.spec.whatwg.org/#list-append){#top-level-traversables:list-append-2
    x-internal="list-append"} `traversable`{.variable} to the user
    agent\'s [top-level traversable
    set](#top-level-traversable-set){#top-level-traversables:top-level-traversable-set}.

12. Invoke [WebDriver BiDi navigable
    created](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigable-created){#top-level-traversables:webdriver-bidi-navigable-created
    x-internal="webdriver-bidi-navigable-created"} with
    `traversable`{.variable} and
    `openerNavigableForWebDriver`{.variable}.

13. Return `traversable`{.variable}.

To [create a fresh top-level
traversable]{#create-a-fresh-top-level-traversable .dfn export=""} given
a
[URL](https://url.spec.whatwg.org/#concept-url){#top-level-traversables:url
x-internal="url"} `initialNavigationURL`{.variable} and an optional
[POST
resource](browsing-the-web.html#post-resource){#top-level-traversables:post-resource}-or-null
`initialNavigationPostResource`{.variable} (default null):

1.  Let `traversable`{.variable} be the result of [creating a new
    top-level
    traversable](#creating-a-new-top-level-traversable){#top-level-traversables:creating-a-new-top-level-traversable}
    given null and the empty string.

2.  [Navigate](browsing-the-web.html#navigate){#top-level-traversables:navigate}
    `traversable`{.variable} to `initialNavigationURL`{.variable} using
    `traversable`{.variable}\'s [active
    document](#nav-document){#top-level-traversables:nav-document}, with
    *[documentResource](browsing-the-web.html#navigation-resource)* set
    to `initialNavigationPostResource`{.variable}.

    We treat these initial navigations as `traversable`{.variable}
    navigating itself, which will ensure all relevant security checks
    pass.

3.  Return `traversable`{.variable}.

##### [7.3.1.3]{.secno} Child navigables[](#child-navigables){.self-link}

Certain elements (for example,
[`iframe`{#child-navigables:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
elements) can present a
[navigable](#navigable){#child-navigables:navigable} to the user. These
elements are called [navigable containers]{#navigable-container .dfn
lt="navigable container" export=""}.

Each [navigable
container](#navigable-container){#child-navigables:navigable-container}
has a [content navigable]{#content-navigable .dfn
dfn-for="navigable container" export=""}, which is either a
[navigable](#navigable){#child-navigables:navigable-2} or null. It is
initially null.

The [container]{#nav-container .dfn dfn-for="navigable" export=""} of a
[navigable](#navigable){#child-navigables:navigable-3}
`navigable`{.variable} is the [navigable
container](#navigable-container){#child-navigables:navigable-container-2}
whose [content
navigable](#content-navigable){#child-navigables:content-navigable} is
`navigable`{.variable}, or null if there is no such element.

The [container document]{#nav-container-document .dfn
dfn-for="navigable" export=""} of a
[navigable](#navigable){#child-navigables:navigable-4}
`navigable`{.variable} is the result of running these steps:

1.  If `navigable`{.variable}\'s
    [container](#nav-container){#child-navigables:nav-container} is
    null, then return null.

2.  Return `navigable`{.variable}\'s
    [container](#nav-container){#child-navigables:nav-container-2}\'s
    [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#child-navigables:node-document
    x-internal="node-document"}.

    This is equal to `navigable`{.variable}\'s
    [container](#nav-container){#child-navigables:nav-container-3}\'s
    [shadow-including
    root](https://dom.spec.whatwg.org/#concept-shadow-including-root){#child-navigables:shadow-including-root
    x-internal="shadow-including-root"} as `navigable`{.variable}\'s
    [container](#nav-container){#child-navigables:nav-container-4} has
    to be
    [connected](https://dom.spec.whatwg.org/#connected){#child-navigables:connected
    x-internal="connected"}.

The [container document]{#doc-container-document .dfn} of a
[`Document`{#child-navigables:document}](dom.html#document)
`document`{.variable} is the result of running these steps:

1.  If `document`{.variable}\'s [node
    navigable](#node-navigable){#child-navigables:node-navigable} is
    null, then return null.

2.  Return `document`{.variable}\'s [node
    navigable](#node-navigable){#child-navigables:node-navigable-2}\'s
    [container
    document](#nav-container-document){#child-navigables:nav-container-document}.

A [navigable](#navigable){#child-navigables:navigable-5}
`navigable`{.variable} is a [child navigable]{#child-navigable .dfn
export=""} of another navigable `potentialParent`{.variable} when
`navigable`{.variable}\'s
[parent](#nav-parent){#child-navigables:nav-parent} is
`potentialParent`{.variable}. We can also just say that a
[navigable](#navigable){#child-navigables:navigable-6} \"is a [child
navigable](#child-navigable){#child-navigables:child-navigable}\", which
means that its [parent](#nav-parent){#child-navigables:nav-parent-2} is
non-null.

All [child
navigables](#child-navigable){#child-navigables:child-navigable-2} are
the [content
navigable](#content-navigable){#child-navigables:content-navigable-2} of
their [container](#nav-container){#child-navigables:nav-container-5}.

The [content document]{#concept-bcc-content-document .dfn} of a
[navigable
container](#navigable-container){#child-navigables:navigable-container-3}
`container`{.variable} is the result of running these steps:

1.  If `container`{.variable}\'s [content
    navigable](#content-navigable){#child-navigables:content-navigable-3}
    is null, then return null.

2.  Let `document`{.variable} be `container`{.variable}\'s [content
    navigable](#content-navigable){#child-navigables:content-navigable-4}\'s
    [active document](#nav-document){#child-navigables:nav-document}.

3.  If `document`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#child-navigables:concept-document-origin
    x-internal="concept-document-origin"} and `container`{.variable}\'s
    [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#child-navigables:node-document-2
    x-internal="node-document"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#child-navigables:concept-document-origin-2
    x-internal="concept-document-origin"} are not [same
    origin-domain](browsers.html#same-origin-domain){#child-navigables:same-origin-domain},
    then return null.

4.  Return `document`{.variable}.

The [content window]{#content-window .dfn} of a [navigable
container](#navigable-container){#child-navigables:navigable-container-4}
`container`{.variable} is the result of running these steps:

1.  If `container`{.variable}\'s [content
    navigable](#content-navigable){#child-navigables:content-navigable-5}
    is null, then return null.

2.  Return `container`{.variable}\'s [content
    navigable](#content-navigable){#child-navigables:content-navigable-6}\'s
    [active `WindowProxy`](#nav-wp){#child-navigables:nav-wp}\'s object.

------------------------------------------------------------------------

To [create a new child navigable]{#create-a-new-child-navigable .dfn},
given an element `element`{.variable}:

1.  Let `parentNavigable`{.variable} be `element`{.variable}\'s [node
    navigable](#node-navigable){#child-navigables:node-navigable-3}.

2.  Let `group`{.variable} be `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#child-navigables:node-document-3
    x-internal="node-document"}\'s [browsing
    context](#concept-document-bc){#child-navigables:concept-document-bc}\'s
    [top-level browsing context](#bc-tlbc){#child-navigables:bc-tlbc}\'s
    [group](#tlbc-group){#child-navigables:tlbc-group}.

3.  Let `browsingContext`{.variable} and `document`{.variable} be the
    result of [creating a new browsing context and
    document](#creating-a-new-browsing-context){#child-navigables:creating-a-new-browsing-context}
    given `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#child-navigables:node-document-4
    x-internal="node-document"}, `element`{.variable}, and
    `group`{.variable}.

4.  Let `targetName`{.variable} be null.

5.  If `element`{.variable} has a `name` content attribute, then set
    `targetName`{.variable} to the value of that attribute.

6.  Let `documentState`{.variable} be a new [document
    state](browsing-the-web.html#document-state-2){#child-navigables:document-state-2},
    with

    [document](browsing-the-web.html#document-state-document){#child-navigables:document-state-document}
    :   `document`{.variable}

    [initiator origin](browsing-the-web.html#document-state-initiator-origin){#child-navigables:document-state-initiator-origin}
    :   `document`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#child-navigables:concept-document-origin-3
        x-internal="concept-document-origin"}

    [origin](browsing-the-web.html#document-state-origin){#child-navigables:document-state-origin}
    :   `document`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#child-navigables:concept-document-origin-4
        x-internal="concept-document-origin"}

    [navigable target name](browsing-the-web.html#document-state-nav-target-name){#child-navigables:document-state-nav-target-name}
    :   `targetName`{.variable}

    [about base URL](browsing-the-web.html#document-state-about-base-url){#child-navigables:document-state-about-base-url}
    :   `document`{.variable}\'s [about base
        URL](dom.html#concept-document-about-base-url){#child-navigables:concept-document-about-base-url}

7.  Let `navigable`{.variable} be a new
    [navigable](#navigable){#child-navigables:navigable-7}.

8.  [Initialize the
    navigable](#initialize-the-navigable){#child-navigables:initialize-the-navigable}
    `navigable`{.variable} given `documentState`{.variable} and
    `parentNavigable`{.variable}.

9.  Set `element`{.variable}\'s [content
    navigable](#content-navigable){#child-navigables:content-navigable-7}
    to `navigable`{.variable}.

10. Let `historyEntry`{.variable} be `navigable`{.variable}\'s [active
    session history
    entry](#nav-active-history-entry){#child-navigables:nav-active-history-entry}.

11. Let `traversable`{.variable} be `parentNavigable`{.variable}\'s
    [traversable
    navigable](#nav-traversable){#child-navigables:nav-traversable}.

12. [Append the following session history traversal
    steps](browsing-the-web.html#tn-append-session-history-traversal-steps){#child-navigables:tn-append-session-history-traversal-steps}
    to `traversable`{.variable}:

    1.  Let `parentDocState`{.variable} be
        `parentNavigable`{.variable}\'s [active session history
        entry](#nav-active-history-entry){#child-navigables:nav-active-history-entry-2}\'s
        [document
        state](browsing-the-web.html#she-document-state){#child-navigables:she-document-state}.

    2.  Let `parentNavigableEntries`{.variable} be the result of
        [getting session history
        entries](browsing-the-web.html#getting-session-history-entries){#child-navigables:getting-session-history-entries}
        for `parentNavigable`{.variable}.

    3.  Let `targetStepSHE`{.variable} be the first [session history
        entry](browsing-the-web.html#session-history-entry){#child-navigables:session-history-entry}
        in `parentNavigableEntries`{.variable} whose [document
        state](browsing-the-web.html#she-document-state){#child-navigables:she-document-state-2}
        equals `parentDocState`{.variable}.

    4.  Set `historyEntry`{.variable}\'s
        [step](browsing-the-web.html#she-step){#child-navigables:she-step}
        to `targetStepSHE`{.variable}\'s
        [step](browsing-the-web.html#she-step){#child-navigables:she-step-2}.

    5.  Let `nestedHistory`{.variable} be a new [nested
        history](browsing-the-web.html#nested-history){#child-navigables:nested-history}
        whose
        [id](browsing-the-web.html#nested-history-id){#child-navigables:nested-history-id}
        is `navigable`{.variable}\'s
        [id](#nav-id){#child-navigables:nav-id} and [entries
        list](browsing-the-web.html#nested-history-entries){#child-navigables:nested-history-entries}
        is « `historyEntry`{.variable} ».

    6.  [Append](https://infra.spec.whatwg.org/#list-append){#child-navigables:list-append
        x-internal="list-append"} `nestedHistory`{.variable} to
        `parentDocState`{.variable}\'s [nested
        histories](browsing-the-web.html#document-state-nested-histories){#child-navigables:document-state-nested-histories}.

    7.  [Update for navigable
        creation/destruction](browsing-the-web.html#update-for-navigable-creation/destruction){#child-navigables:update-for-navigable-creation/destruction}
        given `traversable`{.variable}.

13. Invoke [WebDriver BiDi navigable
    created](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigable-created){#child-navigables:webdriver-bidi-navigable-created
    x-internal="webdriver-bidi-navigable-created"} with
    `traversable`{.variable}.

##### [7.3.1.4]{.secno} Jake diagrams[](#jake-diagrams){.self-link}

A useful method for visualizing sequences of documents, and in
particular [navigables](#navigable){#jake-diagrams:navigable} and their
[session history
entries](browsing-the-web.html#session-history-entry){#jake-diagrams:session-history-entry},
is the [Jake diagram]{#jake-diagram .dfn}. A typical Jake diagram is the
following:

0

1

2

3

4

`top`

/t-a

/t-a#foo

/t-b

`frames[0]`

/i-0-a

/i-0-b

`frames[1]`

/i-1-a

/i-1-b

Here, each numbered column denotes a possible value for the
traversable\'s [session history
step](#tn-current-session-history-step){#jake-diagrams:tn-current-session-history-step}.
Each labeled row depicts a
[navigable](#navigable){#jake-diagrams:navigable-2}, as it transitions
between different URLs and documents. The first, labeled `top`, being
the [top-level
traversable](#top-level-traversable){#jake-diagrams:top-level-traversable},
and the others being [child
navigables](#child-navigable){#jake-diagrams:child-navigable}. The
documents are given by the background color of each cell, with a new
background color indicating a new document in that
[navigable](#navigable){#jake-diagrams:navigable-3}. The URLs are given
by the text content of the cells; usually they are given as [relative
URLs](https://url.spec.whatwg.org/#syntax-url-relative){#jake-diagrams:relative-url
x-internal="relative-url"} for brevity, unless a cross-origin case is
specifically under investigation. A given navigable might not exist at a
given step, in which case the corresponding cells are empty. The
bold-italic step number depicts the [current session history
step](#tn-current-session-history-step){#jake-diagrams:tn-current-session-history-step-2}
of the traversable, and all cells with bold-italic URLs represent the
[current session history
entry](#nav-current-history-entry){#jake-diagrams:nav-current-history-entry}
for that row\'s navigable.

Thus, the above Jake diagram depicts the following sequence of events:

0.  A [top-level
    traversable](#top-level-traversable){#jake-diagrams:top-level-traversable-2}
    is created, starting a the URL `/t-a`, with two [child
    navigables](#child-navigable){#jake-diagrams:child-navigable-2}
    starting at `/i-0-a` and `/i-1-a` respectively.

1.  The first child navigable is
    [navigated](browsing-the-web.html#navigate){#jake-diagrams:navigate}
    to another document, with URL `/i-0-b`.

2.  The second child navigable is
    [navigated](browsing-the-web.html#navigate){#jake-diagrams:navigate-2}
    to another document, with URL `/i-1-b`.

3.  The top-level traversable is
    [navigated](browsing-the-web.html#navigate){#jake-diagrams:navigate-3}
    to the *same* document, updating its URL to `/t-a#foo`.

4.  The top-level traversable is
    [navigated](browsing-the-web.html#navigate){#jake-diagrams:navigate-4}
    to another document, with URL `/t-b`. (Notice how this document, of
    course, does not carry over the old document\'s child navigables.)

5.  The traversable was [traversed by a
    delta](browsing-the-web.html#traverse-the-history-by-a-delta){#jake-diagrams:traverse-the-history-by-a-delta}
    of −3, back to step 1.

[Jake diagrams](#jake-diagram){#jake-diagrams:jake-diagram} are a
powerful tool for visualizing the interactions of multiple navigables,
navigations, and traversals. They cannot capture every possible
interaction --- for example, they only work with a single level of
nesting --- but we will have ocassion to use them to illustrate several
complex situations throughout this standard.

[Jake diagrams](#jake-diagram){#jake-diagrams:jake-diagram-2} are named
after their creator, the inimitable Jake Archibald.

##### [7.3.1.5]{.secno} Related navigable collections[](#related-navigable-collections){.self-link}

It is often helpful in this standard\'s algorithms to look at
collections of
[navigables](#navigable){#related-navigable-collections:navigable}
starting at a given
[`Document`{#related-navigable-collections:document}](dom.html#document).
This section contains a curated set of algorithms for collecting those
navigables.

The return values of these algorithms are ordered so that parents
appears before their children. Callers rely on this ordering.

Starting with a
[`Document`{#related-navigable-collections:document-2}](dom.html#document),
rather than a
[navigable](#navigable){#related-navigable-collections:navigable-2}, is
generally better because it makes the caller cognizant of whether they
are starting with a [fully
active](#fully-active){#related-navigable-collections:fully-active}
[`Document`{#related-navigable-collections:document-3}](dom.html#document)
or not. Although non-[fully
active](#fully-active){#related-navigable-collections:fully-active-2}
[`Document`{#related-navigable-collections:document-4}](dom.html#document)s
do have ancestor and descendant navigables, they often behave as if they
don\'t (e.g., in the
[`window.parent`{#related-navigable-collections:dom-parent}](nav-history-apis.html#dom-parent)
getter).

The [ancestor navigables]{#ancestor-navigables .dfn dfn-for="Document"
export=""} of a
[`Document`{#related-navigable-collections:document-5}](dom.html#document)
`document`{.variable} are given by these steps:

1.  Let `navigable`{.variable} be `document`{.variable}\'s [node
    navigable](#node-navigable){#related-navigable-collections:node-navigable}\'s
    [parent](#nav-parent){#related-navigable-collections:nav-parent}.

2.  Let `ancestors`{.variable} be an empty list.

3.  While `navigable`{.variable} is not null:

    1.  [Prepend](https://infra.spec.whatwg.org/#list-prepend){#related-navigable-collections:list-prepend
        x-internal="list-prepend"} `navigable`{.variable} to
        `ancestors`{.variable}.

    2.  Set `navigable`{.variable} to `navigable`{.variable}\'s
        [parent](#nav-parent){#related-navigable-collections:nav-parent-2}.

4.  Return `ancestors`{.variable}.

The [inclusive ancestor navigables]{#inclusive-ancestor-navigables .dfn
dfn-for="Document" export=""} of a
[`Document`{#related-navigable-collections:document-6}](dom.html#document)
`document`{.variable} are given by these steps:

1.  Let `navigables`{.variable} be `document`{.variable}\'s [ancestor
    navigables](#ancestor-navigables){#related-navigable-collections:ancestor-navigables}.

2.  [Append](https://infra.spec.whatwg.org/#list-append){#related-navigable-collections:list-append
    x-internal="list-append"} `document`{.variable}\'s [node
    navigable](#node-navigable){#related-navigable-collections:node-navigable-2}
    to `navigables`{.variable}.

3.  Return `navigables`{.variable}.

The [descendant navigables]{#descendant-navigables .dfn
dfn-for="Document" export=""} of a
[`Document`{#related-navigable-collections:document-7}](dom.html#document)
`document`{.variable} are given by these steps:

1.  Let `navigables`{.variable} be new
    [list](https://infra.spec.whatwg.org/#list){#related-navigable-collections:list
    x-internal="list"}.

2.  Let `navigableContainers`{.variable} be a
    [list](https://infra.spec.whatwg.org/#list){#related-navigable-collections:list-2
    x-internal="list"} of all [shadow-including
    descendants](https://dom.spec.whatwg.org/#concept-shadow-including-descendant){#related-navigable-collections:shadow-including-descendant
    x-internal="shadow-including-descendant"} of `document`{.variable}
    that are [navigable
    containers](#navigable-container){#related-navigable-collections:navigable-container},
    in [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#related-navigable-collections:shadow-including-tree-order
    x-internal="shadow-including-tree-order"}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#related-navigable-collections:list-iterate
    x-internal="list-iterate"} `navigableContainer`{.variable} of
    `navigableContainers`{.variable}:

    1.  If `navigableContainer`{.variable}\'s [content
        navigable](#content-navigable){#related-navigable-collections:content-navigable}
        is null, then continue.

    2.  [Extend](https://infra.spec.whatwg.org/#list-extend){#related-navigable-collections:list-extend
        x-internal="list-extend"} `navigables`{.variable} with
        `navigableContainer`{.variable}\'s [content
        navigable](#content-navigable){#related-navigable-collections:content-navigable-2}\'s
        [active
        document](#nav-document){#related-navigable-collections:nav-document}\'s
        [inclusive descendant
        navigables](#inclusive-descendant-navigables){#related-navigable-collections:inclusive-descendant-navigables}.

4.  Return `navigables`{.variable}.

The [inclusive descendant navigables]{#inclusive-descendant-navigables
.dfn dfn-for="Document" export=""} of a
[`Document`{#related-navigable-collections:document-8}](dom.html#document)
`document`{.variable} are given by these steps:

1.  Let `navigables`{.variable} be « `document`{.variable}\'s [node
    navigable](#node-navigable){#related-navigable-collections:node-navigable-3}
    ».

2.  [Extend](https://infra.spec.whatwg.org/#list-extend){#related-navigable-collections:list-extend-2
    x-internal="list-extend"} `navigables`{.variable} with
    `document`{.variable}\'s [descendant
    navigables](#descendant-navigables){#related-navigable-collections:descendant-navigables}.

3.  Return `navigables`{.variable}.

These descendant-collecting algorithms are described as looking at the
DOM tree of descendant
[`Document`{#related-navigable-collections:document-9}](dom.html#document)
objects. In reality, this is often not feasible since the DOM tree can
be in another process from the caller of the algorithm. Instead,
implementations generally replicate the appropriate trees across
processes.

The [document-tree child navigables]{#document-tree-child-navigables
.dfn} of a
[`Document`{#related-navigable-collections:document-10}](dom.html#document)
`document`{.variable} are given by these steps:

1.  If `document`{.variable}\'s [node
    navigable](#node-navigable){#related-navigable-collections:node-navigable-4}
    is null, then return the empty list.

2.  Let `navigables`{.variable} be new
    [list](https://infra.spec.whatwg.org/#list){#related-navigable-collections:list-3
    x-internal="list"}.

3.  Let `navigableContainers`{.variable} be a
    [list](https://infra.spec.whatwg.org/#list){#related-navigable-collections:list-4
    x-internal="list"} of all
    [descendants](https://dom.spec.whatwg.org/#concept-tree-descendant){#related-navigable-collections:descendant
    x-internal="descendant"} of `document`{.variable} that are
    [navigable
    containers](#navigable-container){#related-navigable-collections:navigable-container-2},
    in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#related-navigable-collections:tree-order
    x-internal="tree-order"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#related-navigable-collections:list-iterate-2
    x-internal="list-iterate"} `navigableContainer`{.variable} of
    `navigableContainers`{.variable}:

    1.  If `navigableContainer`{.variable}\'s [content
        navigable](#content-navigable){#related-navigable-collections:content-navigable-3}
        is null, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#related-navigable-collections:continue
        x-internal="continue"}.

    2.  [Append](https://infra.spec.whatwg.org/#list-append){#related-navigable-collections:list-append-2
        x-internal="list-append"} `navigableContainer`{.variable}\'s
        [content
        navigable](#content-navigable){#related-navigable-collections:content-navigable-4}
        to `navigables`{.variable}.

5.  Return `navigables`{.variable}.

##### [7.3.1.6]{.secno} Navigable destruction[](#garbage-collection-and-browsing-contexts){.self-link} {#garbage-collection-and-browsing-contexts}

To [destroy a child navigable]{#destroy-a-child-navigable .dfn} given a
[navigable
container](#navigable-container){#garbage-collection-and-browsing-contexts:navigable-container}
`container`{.variable}:

1.  Let `navigable`{.variable} be `container`{.variable}\'s [content
    navigable](#content-navigable){#garbage-collection-and-browsing-contexts:content-navigable}.

2.  If `navigable`{.variable} is null, then return.

3.  Set `container`{.variable}\'s [content
    navigable](#content-navigable){#garbage-collection-and-browsing-contexts:content-navigable-2}
    to null.

4.  [Inform the navigation API about child navigable
    destruction](nav-history-apis.html#inform-the-navigation-api-about-child-navigable-destruction){#garbage-collection-and-browsing-contexts:inform-the-navigation-api-about-child-navigable-destruction}
    given `navigable`{.variable}.

5.  [Destroy a document and its
    descendants](document-lifecycle.html#destroy-a-document-and-its-descendants){#garbage-collection-and-browsing-contexts:destroy-a-document-and-its-descendants}
    given `navigable`{.variable}\'s [active
    document](#nav-document){#garbage-collection-and-browsing-contexts:nav-document}.

6.  Let `parentDocState`{.variable} be `container`{.variable}\'s [node
    navigable](#node-navigable){#garbage-collection-and-browsing-contexts:node-navigable}\'s
    [active session history
    entry](#nav-active-history-entry){#garbage-collection-and-browsing-contexts:nav-active-history-entry}\'s
    [document
    state](browsing-the-web.html#she-document-state){#garbage-collection-and-browsing-contexts:she-document-state}.

7.  [Remove](https://infra.spec.whatwg.org/#list-remove){#garbage-collection-and-browsing-contexts:list-remove
    x-internal="list-remove"} the [nested
    history](browsing-the-web.html#nested-history){#garbage-collection-and-browsing-contexts:nested-history}
    from `parentDocState`{.variable}\'s [nested
    histories](browsing-the-web.html#document-state-nested-histories){#garbage-collection-and-browsing-contexts:document-state-nested-histories}
    whose
    [id](browsing-the-web.html#nested-history-id){#garbage-collection-and-browsing-contexts:nested-history-id}
    equals `navigable`{.variable}\'s
    [id](#nav-id){#garbage-collection-and-browsing-contexts:nav-id}.

8.  Let `traversable`{.variable} be `container`{.variable}\'s [node
    navigable](#node-navigable){#garbage-collection-and-browsing-contexts:node-navigable-2}\'s
    [traversable
    navigable](#nav-traversable){#garbage-collection-and-browsing-contexts:nav-traversable}.

9.  [Append the following session history traversal
    steps](browsing-the-web.html#tn-append-session-history-traversal-steps){#garbage-collection-and-browsing-contexts:tn-append-session-history-traversal-steps}
    to `traversable`{.variable}:

    1.  [Update for navigable
        creation/destruction](browsing-the-web.html#update-for-navigable-creation/destruction){#garbage-collection-and-browsing-contexts:update-for-navigable-creation/destruction}
        given `traversable`{.variable}.

10. Invoke [WebDriver BiDi navigable
    destroyed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigable-destroyed){#garbage-collection-and-browsing-contexts:webdriver-bidi-navigable-destroyed
    x-internal="webdriver-bidi-navigable-destroyed"} with
    `navigable`{.variable}.

To [destroy]{#destroy-a-top-level-traversable .dfn} a [top-level
traversable](#top-level-traversable){#garbage-collection-and-browsing-contexts:top-level-traversable}
`traversable`{.variable}:

1.  Let `browsingContext`{.variable} be `traversable`{.variable}\'s
    [active browsing
    context](#nav-bc){#garbage-collection-and-browsing-contexts:nav-bc}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#garbage-collection-and-browsing-contexts:list-iterate
    x-internal="list-iterate"} `historyEntry`{.variable} in
    `traversable`{.variable}\'s [session history
    entries](#tn-session-history-entries){#garbage-collection-and-browsing-contexts:tn-session-history-entries}
    [in what order?]{.XXX}:

    1.  Let `document`{.variable} be `historyEntry`{.variable}\'s
        [document](browsing-the-web.html#she-document){#garbage-collection-and-browsing-contexts:she-document}.

    2.  If `document`{.variable} is not null, then [destroy a document
        and its
        descendants](document-lifecycle.html#destroy-a-document-and-its-descendants){#garbage-collection-and-browsing-contexts:destroy-a-document-and-its-descendants-2}
        given `document`{.variable}.

3.  [Remove](#bcg-remove){#garbage-collection-and-browsing-contexts:bcg-remove}
    `browsingContext`{.variable}.

4.  Remove `traversable`{.variable} from the user interface (e.g., close
    or hide its tab in a tabbed browser).

5.  [Remove](https://infra.spec.whatwg.org/#list-remove){#garbage-collection-and-browsing-contexts:list-remove-2
    x-internal="list-remove"} `traversable`{.variable} from the user
    agent\'s [top-level traversable
    set](#top-level-traversable-set){#garbage-collection-and-browsing-contexts:top-level-traversable-set}.

6.  Invoke [WebDriver BiDi navigable
    destroyed](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-navigable-destroyed){#garbage-collection-and-browsing-contexts:webdriver-bidi-navigable-destroyed-2
    x-internal="webdriver-bidi-navigable-destroyed"} with
    `traversable`{.variable}.

User agents may [destroy a top-level
traversable](#destroy-a-top-level-traversable){#garbage-collection-and-browsing-contexts:destroy-a-top-level-traversable}
at any time (typically, [in response to user
requests](document-lifecycle.html#nav-traversal-ui)).

[]{#close-a-browsing-context}To [close]{#close-a-top-level-traversable
.dfn dfn-for="top-level traversable" export=""} a [top-level
traversable](#top-level-traversable){#garbage-collection-and-browsing-contexts:top-level-traversable-2}
`traversable`{.variable}:

1.  If `traversable`{.variable}\'s [is
    closing](#is-closing){#garbage-collection-and-browsing-contexts:is-closing}
    is true, then return.

2.  [Definitely
    close](#definitely-close-a-top-level-traversable){#garbage-collection-and-browsing-contexts:definitely-close-a-top-level-traversable}
    `traversable`{.variable}.

To [definitely close]{#definitely-close-a-top-level-traversable .dfn} a
[top-level
traversable](#top-level-traversable){#garbage-collection-and-browsing-contexts:top-level-traversable-3}
`traversable`{.variable}:

1.  Let `toUnload`{.variable} be `traversable`{.variable}\'s [active
    document](#nav-document){#garbage-collection-and-browsing-contexts:nav-document-2}\'s
    [inclusive descendant
    navigables](#inclusive-descendant-navigables){#garbage-collection-and-browsing-contexts:inclusive-descendant-navigables}.

2.  If the result of [checking if unloading is
    canceled](browsing-the-web.html#checking-if-unloading-is-canceled){#garbage-collection-and-browsing-contexts:checking-if-unloading-is-canceled}
    for `toUnload`{.variable} is true, then return.

3.  [Append the following session history traversal
    steps](browsing-the-web.html#tn-append-session-history-traversal-steps){#garbage-collection-and-browsing-contexts:tn-append-session-history-traversal-steps-2}
    to `traversable`{.variable}:

    1.  Let `afterAllUnloads`{.variable} be an algorithm step which
        [destroys](#destroy-a-top-level-traversable){#garbage-collection-and-browsing-contexts:destroy-a-top-level-traversable-2}
        `traversable`{.variable}.

    2.  [Unload a document and its
        descendants](document-lifecycle.html#unload-a-document-and-its-descendants){#garbage-collection-and-browsing-contexts:unload-a-document-and-its-descendants}
        given `traversable`{.variable}\'s [active
        document](#nav-document){#garbage-collection-and-browsing-contexts:nav-document-3},
        null, and `afterAllUnloads`{.variable}.

The
[close](#close-a-top-level-traversable){#garbage-collection-and-browsing-contexts:close-a-top-level-traversable}
vs. [definitely
close](#definitely-close-a-top-level-traversable){#garbage-collection-and-browsing-contexts:definitely-close-a-top-level-traversable-2}
separation allows other specifications to call
[close](#close-a-top-level-traversable){#garbage-collection-and-browsing-contexts:close-a-top-level-traversable-2}
and have it be a no-op if the top-level traversable is already closing
due to JavaScript code calling
[`window.close()`{#garbage-collection-and-browsing-contexts:dom-window-close}](nav-history-apis.html#dom-window-close).

##### [7.3.1.7]{.secno} []{#browsing-context-names}Navigable target names[](#navigable-target-names){.self-link}

[Navigables](#navigable){#navigable-target-names:navigable} can be given
[target names](#nav-target){#navigable-target-names:nav-target}, which
are strings allowing certain APIs (such as
[`window.open()`{#navigable-target-names:dom-open}](nav-history-apis.html#dom-open)
or the
[`a`{#navigable-target-names:the-a-element}](text-level-semantics.html#the-a-element)
element\'s
[`target`{#navigable-target-names:attr-hyperlink-target}](links.html#attr-hyperlink-target)
attribute) to target
[navigations](browsing-the-web.html#navigate){#navigable-target-names:navigate}
at that navigable.

A [valid navigable target name]{#valid-navigable-target-name .dfn} is
any string with at least one character that does not contain both an
[ASCII tab or
newline](https://infra.spec.whatwg.org/#ascii-tab-or-newline){#navigable-target-names:ascii-tab-or-newline
x-internal="ascii-tab-or-newline"} and a U+003C (\<), and it does not
start with a U+005F (\_). (Names starting with a U+005F (\_) are
reserved for special keywords.)

A [valid navigable target name or
keyword]{#valid-navigable-target-name-or-keyword .dfn} is any string
that is either a [valid navigable target
name](#valid-navigable-target-name){#navigable-target-names:valid-navigable-target-name}
or that is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#navigable-target-names:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for one of: `_blank`,
`_self`, `_parent`, or `_top`.

These values have different meanings based on whether the page is
sandboxed or not, as summarized in the following (non-normative) table.
In this table, \"current\" means the
[navigable](#navigable){#navigable-target-names:navigable-2} that the
link or script is in, \"parent\" means the
[parent](#nav-parent){#navigable-target-names:nav-parent} of the
[navigable](#navigable){#navigable-target-names:navigable-3} that the
link or script is in, \"top\" means the [top-level
traversable](#nav-top){#navigable-target-names:nav-top} of the
[navigable](#navigable){#navigable-target-names:navigable-4} that the
link or script is in, \"new\" means a new [traversable
navigable](#traversable-navigable){#navigable-target-names:traversable-navigable}
with a null [parent](#nav-parent){#navigable-target-names:nav-parent-2}
(which may use an [auxiliary browsing
context](#auxiliary-browsing-context){#navigable-target-names:auxiliary-browsing-context},
subject to various user preferences and user agent policies), \"none\"
means that nothing will happen, and \"maybe new\" means the same as
\"new\" if the
\"[`allow-popups`{#navigable-target-names:attr-iframe-sandbox-allow-popups}](browsers.html#attr-iframe-sandbox-allow-popups)\"
keyword is also specified on the
[`sandbox`{#navigable-target-names:attr-iframe-sandbox}](iframe-embed-object.html#attr-iframe-sandbox)
attribute (or if the user overrode the sandboxing), and the same as
\"none\" otherwise.

Keyword

Ordinary effect

Effect in an
[`iframe`{#navigable-target-names:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
with\...

`sandbox=""`

`sandbox="allow-top-navigation"`

none specified, for links and form submissions

current

current

current

empty string

current

current

current

`_blank`

new

maybe new

maybe new

`_self`

current

current

current

`_parent` if there isn\'t a parent

current

current

current

`_parent` if parent is also top

parent/top

none

parent/top

`_parent` if there is one and it\'s not top

parent

none

none

`_top` if top is current

current

current

current

`_top` if top is not current

top

none

top

name that doesn\'t exist

new

maybe new

maybe new

name that exists and is a descendant

specified descendant

specified descendant

specified descendant

name that exists and is current

current

current

current

name that exists and is an ancestor that is top

specified ancestor

none

specified ancestor/top

name that exists and is an ancestor that is not top

specified ancestor

none

none

other name that exists with common top

specified

none

none

name that exists with different top, if
[familiar](#familiar-with){#navigable-target-names:familiar-with} and
[one permitted sandboxed
navigator](browsers.html#one-permitted-sandboxed-navigator){#navigable-target-names:one-permitted-sandboxed-navigator}

specified

specified

specified

name that exists with different top, if
[familiar](#familiar-with){#navigable-target-names:familiar-with-2} but
not [one permitted sandboxed
navigator](browsers.html#one-permitted-sandboxed-navigator){#navigable-target-names:one-permitted-sandboxed-navigator-2}

specified

none

none

name that exists with different top, not
[familiar](#familiar-with){#navigable-target-names:familiar-with-3}

new

maybe new

maybe new

[Most of the restrictions on sandboxed browsing contexts are applied by
other algorithms, e.g. the
[navigation](browsing-the-web.html#navigate){#navigable-target-names:navigate-2}
algorithm, not [the rules for choosing a
navigable](#the-rules-for-choosing-a-navigable){#navigable-target-names:the-rules-for-choosing-a-navigable}
given below.]{.small}

------------------------------------------------------------------------

To [find a navigable by target name]{#find-a-navigable-by-target-name
.dfn} given a string `name`{.variable} and a
[navigable](#navigable){#navigable-target-names:navigable-5}
`currentNavigable`{.variable}:

1.  Let `currentDocument`{.variable} be `currentNavigable`{.variable}\'s
    [active
    document](#nav-document){#navigable-target-names:nav-document}.

2.  Let `sourceSnapshotParams`{.variable} be the result of [snapshotting
    source snapshot
    params](browsing-the-web.html#snapshotting-source-snapshot-params){#navigable-target-names:snapshotting-source-snapshot-params}
    given `currentDocument`{.variable}.

3.  Let `subtreesToSearch`{.variable} be an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#navigable-target-names:implementation-defined
    x-internal="implementation-defined"} choice of one of the following:

    - « `currentNavigable`{.variable}\'s [traversable
      navigable](#nav-traversable){#navigable-target-names:nav-traversable},
      `currentNavigable`{.variable} »

    - the [inclusive ancestor
      navigables](#inclusive-ancestor-navigables){#navigable-target-names:inclusive-ancestor-navigables}
      of `currentDocument`{.variable}

    [Issue #10848](https://github.com/whatwg/html/issues/10848) tracks
    settling on one of these two possibilities, to achieve
    interoperability.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#navigable-target-names:list-iterate
    x-internal="list-iterate"} `subtreeToSearch`{.variable} of
    `subtreesToSearch`{.variable}, in reverse order:

    1.  Let `documentToSearch`{.variable} be
        `subtreeToSearch`{.variable}\'s [active
        document](#nav-document){#navigable-target-names:nav-document-2}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#navigable-target-names:list-iterate-2
        x-internal="list-iterate"} `navigable`{.variable} of the
        [inclusive descendant
        navigables](#inclusive-descendant-navigables){#navigable-target-names:inclusive-descendant-navigables}
        of `documentToSearch`{.variable}:

        1.  If `currentNavigable`{.variable} is not [allowed by
            sandboxing to
            navigate](browsing-the-web.html#allowed-to-navigate){#navigable-target-names:allowed-to-navigate}
            `navigable`{.variable} given
            `sourceSnapshotParams`{.variable}, then optionally
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#navigable-target-names:continue
            x-internal="continue"}.

            [Issue #10849](https://github.com/whatwg/html/issues/10849)
            tracks making this check required, to achieve
            interoperability.

        2.  If `navigable`{.variable}\'s [target
            name](#nav-target){#navigable-target-names:nav-target-2} is
            `name`{.variable}, then return `navigable`{.variable}.

5.  Let `currentTopLevelBrowsingContext`{.variable} be
    `currentNavigable`{.variable}\'s [active browsing
    context](#nav-bc){#navigable-target-names:nav-bc}\'s [top-level
    browsing context](#bc-tlbc){#navigable-target-names:bc-tlbc}.

6.  Let `group`{.variable} be
    `currentTopLevelBrowsingContext`{.variable}\'s
    [group](#tlbc-group){#navigable-target-names:tlbc-group}.

7.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#navigable-target-names:list-iterate-3
    x-internal="list-iterate"} `topLevelBrowsingContext`{.variable} of
    `group`{.variable}\'s [browsing context
    set](#browsing-context-set){#navigable-target-names:browsing-context-set},
    in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#navigable-target-names:implementation-defined-2
    x-internal="implementation-defined"} order (the user agent should
    pick a consistent ordering, such as the most recently opened, most
    recently focused, or more closely related):

    [Issue #10850](https://github.com/whatwg/html/issues/10850) tracks
    picking a specific ordering, to achieve interoperability.

    1.  If `currentTopLevelBrowsingContext`{.variable} is
        `topLevelBrowsingContext`{.variable}, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#navigable-target-names:continue-2
        x-internal="continue"}.

    2.  Let `documentToSearch`{.variable} be
        `topLevelBrowsingContext`{.variable}\'s [active
        document](#active-document){#navigable-target-names:active-document}.

    3.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#navigable-target-names:list-iterate-4
        x-internal="list-iterate"} `navigable`{.variable} of the
        [inclusive descendant
        navigables](#inclusive-descendant-navigables){#navigable-target-names:inclusive-descendant-navigables-2}
        of `documentToSearch`{.variable}:

        1.  If `currentNavigable`{.variable}\'s [active browsing
            context](#nav-bc){#navigable-target-names:nav-bc-2} is not
            [familiar
            with](#familiar-with){#navigable-target-names:familiar-with-4}
            `navigable`{.variable}\'s [active browsing
            context](#nav-bc){#navigable-target-names:nav-bc-3}, then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#navigable-target-names:continue-3
            x-internal="continue"}.

        2.  If `currentNavigable`{.variable} is not [allowed by
            sandboxing to
            navigate](browsing-the-web.html#allowed-to-navigate){#navigable-target-names:allowed-to-navigate-2}
            `navigable`{.variable} given
            `sourceSnapshotParams`{.variable}, then optionally
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#navigable-target-names:continue-4
            x-internal="continue"}.

            [Issue #10849](https://github.com/whatwg/html/issues/10849)
            tracks making this check required, to achieve
            interoperability.

        3.  If `navigable`{.variable}\'s [target
            name](#nav-target){#navigable-target-names:nav-target-3} is
            `name`{.variable}, then return `navigable`{.variable}.

8.  Return null.

[The rules for choosing a navigable]{#the-rules-for-choosing-a-navigable
.dfn}, given a string `name`{.variable}, a
[navigable](#navigable){#navigable-target-names:navigable-6}
`currentNavigable`{.variable}, and a boolean `noopener`{.variable} are
as follows:

1.  Let `chosen`{.variable} be null.

2.  Let `windowType`{.variable} be \"`existing or none`\".

3.  Let `sandboxingFlagSet`{.variable} be
    `currentNavigable`{.variable}\'s [active
    document](#nav-document){#navigable-target-names:nav-document-3}\'s
    [active sandboxing flag
    set](browsers.html#active-sandboxing-flag-set){#navigable-target-names:active-sandboxing-flag-set}.

4.  If `name`{.variable} is the empty string or an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#navigable-target-names:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for \"`_self`\", then set
    `chosen`{.variable} to `currentNavigable`{.variable}.

5.  Otherwise, if `name`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#navigable-target-names:ascii-case-insensitive-3
    x-internal="ascii-case-insensitive"} match for \"`_parent`\", set
    `chosen`{.variable} to `currentNavigable`{.variable}\'s
    [parent](#nav-parent){#navigable-target-names:nav-parent-3}, if any,
    and `currentNavigable`{.variable} otherwise.

6.  Otherwise, if `name`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#navigable-target-names:ascii-case-insensitive-4
    x-internal="ascii-case-insensitive"} match for \"`_top`\", set
    `chosen`{.variable} to `currentNavigable`{.variable}\'s [traversable
    navigable](#nav-traversable){#navigable-target-names:nav-traversable-2}.

7.  Otherwise, if `name`{.variable} is not an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#navigable-target-names:ascii-case-insensitive-5
    x-internal="ascii-case-insensitive"} match for \"`_blank`\" and
    `noopener`{.variable} is false, then set `chosen`{.variable} to the
    result of [finding a navigable by target
    name](#find-a-navigable-by-target-name){#navigable-target-names:find-a-navigable-by-target-name}
    given `name`{.variable} and `currentNavigable`{.variable}.

8.  If `chosen`{.variable} is null, then a new [top-level
    traversable](#top-level-traversable){#navigable-target-names:top-level-traversable}
    is being requested, and what happens depends on the user agent\'s
    configuration and abilities --- it is determined by the rules given
    for the first applicable option from the following list:

    If `currentNavigable`{.variable}\'s [active window](#nav-window){#navigable-target-names:nav-window} does not have [transient activation](interaction.html#transient-activation){#navigable-target-names:transient-activation} and the user agent has been configured to not show popups (i.e., the user agent has a \"popup blocker\" enabled)
    :   The user agent may inform the user that a popup has been
        blocked.

    If `sandboxingFlagSet`{.variable} has the [sandboxed auxiliary navigation browsing context flag](browsers.html#sandboxed-auxiliary-navigation-browsing-context-flag){#navigable-target-names:sandboxed-auxiliary-navigation-browsing-context-flag} set
    :   The user agent may report to a developer console that a popup
        has been blocked.

    If the user agent has been configured such that in this instance it will create a new [top-level traversable](#top-level-traversable){#navigable-target-names:top-level-traversable-2}

    :   1.  [Consume user
            activation](interaction.html#consume-user-activation){#navigable-target-names:consume-user-activation}
            of `currentNavigable`{.variable}\'s [active
            window](#nav-window){#navigable-target-names:nav-window-2}.

        2.  Set `windowType`{.variable} to \"`new and unrestricted`\".

        3.  Let `currentDocument`{.variable} be
            `currentNavigable`{.variable}\'s [active
            document](#nav-document){#navigable-target-names:nav-document-4}.

        4.  If `currentDocument`{.variable}\'s [opener
            policy](dom.html#concept-document-coop){#navigable-target-names:concept-document-coop}\'s
            [value](browsers.html#coop-struct-value){#navigable-target-names:coop-struct-value}
            is
            \"[`same-origin`{#navigable-target-names:coop-same-origin}](browsers.html#coop-same-origin)\"
            or
            \"[`same-origin-plus-COEP`{#navigable-target-names:coop-same-origin-plus-coep}](browsers.html#coop-same-origin-plus-coep)\",
            and `currentDocument`{.variable}\'s
            [origin](https://dom.spec.whatwg.org/#concept-document-origin){#navigable-target-names:concept-document-origin
            x-internal="concept-document-origin"} is not [same
            origin](browsers.html#same-origin){#navigable-target-names:same-origin}
            with `currentDocument`{.variable}\'s [relevant settings
            object](webappapis.html#relevant-settings-object){#navigable-target-names:relevant-settings-object}\'s
            [top-level
            origin](webappapis.html#concept-environment-top-level-origin){#navigable-target-names:concept-environment-top-level-origin},
            then:

            1.  Set `noopener`{.variable} to true.

            2.  Set `name`{.variable} to \"`_blank`\".

            3.  Set `windowType`{.variable} to \"`new with no opener`\".

            In the presence of an [opener
            policy](browsers.html#cross-origin-opener-policy){#navigable-target-names:cross-origin-opener-policy},
            nested documents that are cross-origin with their top-level
            browsing context\'s active document always set
            `noopener`{.variable} to true.

        5.  Let `targetName`{.variable} be the empty string.

        6.  If `name`{.variable} is not an [ASCII
            case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#navigable-target-names:ascii-case-insensitive-6
            x-internal="ascii-case-insensitive"} match for \"`_blank`\",
            then set `targetName`{.variable} to `name`{.variable}.

        7.  [If `noopener`{.variable} is true, then set
            `chosen`{.variable} to the result of [creating a new
            top-level
            traversable](#creating-a-new-top-level-traversable){#navigable-target-names:creating-a-new-top-level-traversable}
            given null, `targetName`{.variable}, and
            `currentNavigable`{.variable}.]{#noopener}

        8.  Otherwise:

            1.  Set `chosen`{.variable} to the result of [creating a new
                top-level
                traversable](#creating-a-new-top-level-traversable){#navigable-target-names:creating-a-new-top-level-traversable-2}
                given `currentNavigable`{.variable}\'s [active browsing
                context](#nav-bc){#navigable-target-names:nav-bc-4},
                `targetName`{.variable}, and
                `currentNavigable`{.variable}.

            2.  If `sandboxingFlagSet`{.variable}\'s [sandboxed
                navigation browsing context
                flag](browsers.html#sandboxed-navigation-browsing-context-flag){#navigable-target-names:sandboxed-navigation-browsing-context-flag}
                is set, then set `chosen`{.variable}\'s [active browsing
                context](#nav-bc){#navigable-target-names:nav-bc-5}\'s
                [one permitted sandboxed
                navigator](browsers.html#one-permitted-sandboxed-navigator){#navigable-target-names:one-permitted-sandboxed-navigator-3}
                to `currentNavigable`{.variable}\'s [active browsing
                context](#nav-bc){#navigable-target-names:nav-bc-6}.

        9.  If `sandboxingFlagSet`{.variable}\'s [sandbox propagates to
            auxiliary browsing contexts
            flag](browsers.html#sandbox-propagates-to-auxiliary-browsing-contexts-flag){#navigable-target-names:sandbox-propagates-to-auxiliary-browsing-contexts-flag}
            is set, then all the flags that are set in
            `sandboxingFlagSet`{.variable} must be set in
            `chosen`{.variable}\'s [active browsing
            context](#nav-bc){#navigable-target-names:nav-bc-7}\'s
            [popup sandboxing flag
            set](browsers.html#popup-sandboxing-flag-set){#navigable-target-names:popup-sandboxing-flag-set}.

        10. Set `chosen`{.variable}\'s [is created by web
            content](#is-created-by-web-content){#navigable-target-names:is-created-by-web-content}
            to true.

        If the newly created
        [navigable](#navigable){#navigable-target-names:navigable-7}
        `chosen`{.variable} is immediately
        [navigated](browsing-the-web.html#navigate){#navigable-target-names:navigate-3},
        then the navigation will be done as a
        \"[`replace`{#navigable-target-names:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\"
        navigation.

    If the user agent has been configured such that in this instance it will choose `currentNavigable`{.variable}
    :   Set `chosen`{.variable} to `currentNavigable`{.variable}.

    If the user agent has been configured such that in this instance it will not find a navigable

    :   Do nothing.

    User agents are encouraged to provide a way for users to configure
    the user agent to always choose `currentNavigable`{.variable}.

9.  Return `chosen`{.variable} and `windowType`{.variable}.

#### [7.3.2]{.secno} Browsing contexts[](#windows){.self-link} {#windows}

A [browsing context]{#browsing-context .dfn export=""} is a programmatic
representation of a series of documents, multiple of which can live
within a single [navigable](#navigable){#windows:navigable}. Each
[browsing context](#browsing-context){#windows:browsing-context} has a
corresponding
[`WindowProxy`{#windows:windowproxy}](nav-history-apis.html#windowproxy)
object, as well as the following:

- An [opener browsing context]{#opener-browsing-context .dfn export=""},
  a [browsing context](#browsing-context){#windows:browsing-context-2}
  or null, initially null.

- An [opener origin at creation]{#opener-origin-at-creation .dfn}, an
  [origin](browsers.html#concept-origin){#windows:concept-origin} or
  null, initially null.

- An [is popup]{#is-popup .dfn} boolean, initially false.

  The only mandatory impact in this specification of [is
  popup](#is-popup){#windows:is-popup} is on the
  [`visible`{#windows:dom-barprop-visible}](nav-history-apis.html#dom-barprop-visible)
  getter of the relevant
  [`BarProp`{#windows:barprop}](nav-history-apis.html#barprop) objects.
  However, user agents might also use it for [user interface
  considerations](document-lifecycle.html#nav-traversal-ui).

- An [is auxiliary]{#is-auxiliary .dfn} boolean, initially false.

- An [initial URL]{#browsing-context-initial-url .dfn}, a
  [URL](https://url.spec.whatwg.org/#concept-url){#windows:url
  x-internal="url"} or null, initially null.

- A [virtual browsing context group
  ID]{#virtual-browsing-context-group-id .dfn} integer, initially 0.
  This is used by [opener policy
  reporting](browsers.html#coop-struct-report-only-value){#windows:coop-struct-report-only-value},
  to keep track of the browsing context group switches that would have
  happened if the report-only policy had been enforced.

A [browsing context](#browsing-context){#windows:browsing-context-3}\'s
[active window]{#active-window .dfn} is its
[`WindowProxy`{#windows:windowproxy-2}](nav-history-apis.html#windowproxy)
object\'s
[\[\[Window\]\]](nav-history-apis.html#concept-windowproxy-window){#windows:concept-windowproxy-window}
internal slot value. A [browsing
context](#browsing-context){#windows:browsing-context-4}\'s [active
document]{#active-document .dfn} is its [active
window](#active-window){#windows:active-window}\'s [associated
`Document`](nav-history-apis.html#concept-document-window){#windows:concept-document-window}.

A [browsing context](#browsing-context){#windows:browsing-context-5}\'s
[top-level traversable]{#bc-traversable .dfn dfn-for="browsing context"
export=""} is its [active
document](#active-document){#windows:active-document}\'s [node
navigable](#node-navigable){#windows:node-navigable}\'s [top-level
traversable](#nav-top){#windows:nav-top}.

A [browsing context](#browsing-context){#windows:browsing-context-6}
whose [is auxiliary](#is-auxiliary){#windows:is-auxiliary} is true is
known as an [auxiliary browsing context]{#auxiliary-browsing-context
.dfn export=""}. Auxiliary browsing contexts are always [top-level
browsing
contexts](#top-level-browsing-context){#windows:top-level-browsing-context}.

It\'s unclear whether a separate [is
auxiliary](#is-auxiliary){#windows:is-auxiliary-2} concept is necessary.
In [issue #5680](https://github.com/whatwg/html/issues/5680), it is
indicated that we may be able to simplify this by using whether or not
the [opener browsing
context](#opener-browsing-context){#windows:opener-browsing-context} is
null.

Modern specifications should avoid using the [browsing
context](#browsing-context){#windows:browsing-context-7} concept in most
cases, unless they are dealing with the subtleties of [browsing context
group
switches](browsers.html#browsing-context-group-switches-due-to-cross-origin-opener-policy)
and [agent cluster
allocation](#agent-cluster-map){#windows:agent-cluster-map}. Instead,
the [`Document`{#windows:document}](dom.html#document) and
[navigable](#navigable){#windows:navigable-2} concepts are usually more
appropriate.

------------------------------------------------------------------------

A [`Document`\'s browsing context]{#concept-document-bc .dfn
dfn-for="Document" lt="browsing context" export=""} is a [browsing
context](#browsing-context){#windows:browsing-context-8} or null,
initially null.

A [`Document`{#windows:document-2}](dom.html#document) does not
necessarily have a non-null [browsing
context](#concept-document-bc){#windows:concept-document-bc}. In
particular, data mining tools are likely to never instantiate browsing
contexts. A [`Document`{#windows:document-3}](dom.html#document) created
using an API such as
[`createDocument()`{#windows:dom-domimplementation-createdocument}](https://dom.spec.whatwg.org/#dom-domimplementation-createdocument){x-internal="dom-domimplementation-createdocument"}
never has a non-null [browsing
context](#concept-document-bc){#windows:concept-document-bc-2}. And the
[`Document`{#windows:document-4}](dom.html#document) originally created
for an
[`iframe`{#windows:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
element, which has since been [removed from the
document](infrastructure.html#remove-an-element-from-a-document){#windows:remove-an-element-from-a-document},
has no associated browsing context, since that browsing context was
[nulled
out](document-lifecycle.html#destroy-a-document){#windows:destroy-a-document}.

In general, there is a 1-to-1 mapping from the
[`Window`{#windows:window}](nav-history-apis.html#window) object to the
[`Document`{#windows:document-5}](dom.html#document) object, as long as
the [`Document`{#windows:document-6}](dom.html#document) object has a
non-null [browsing
context](#concept-document-bc){#windows:concept-document-bc-3}. There is
one exception. A
[`Window`{#windows:window-2}](nav-history-apis.html#window) can be
reused for the presentation of a second
[`Document`{#windows:document-7}](dom.html#document) in the same
[browsing context](#browsing-context){#windows:browsing-context-9}, such
that the mapping is then 1-to-2. This occurs when a [browsing
context](#browsing-context){#windows:browsing-context-10} is
[navigated](browsing-the-web.html#navigate){#windows:navigate} from the
[initial
`about:blank`](dom.html#is-initial-about:blank){#windows:is-initial-about:blank}
[`Document`{#windows:document-8}](dom.html#document) to another, which
will be done with
[replacement](browsing-the-web.html#navigationhistorybehavior-replace){#windows:navigationhistorybehavior-replace}.

##### [7.3.2.1]{.secno} Creating browsing contexts[](#creating-browsing-contexts){.self-link}

To [create a new browsing context and
document]{#creating-a-new-browsing-context .dfn}, given null or a
[`Document`{#creating-browsing-contexts:document}](dom.html#document)
object `creator`{.variable}, null or an element `embedder`{.variable},
and a [browsing context
group](#browsing-context-group){#creating-browsing-contexts:browsing-context-group}
`group`{.variable}:

1.  Let `browsingContext`{.variable} be a new [browsing
    context](#browsing-context){#creating-browsing-contexts:browsing-context}.

2.  Let `unsafeContextCreationTime`{.variable} be the [unsafe shared
    current
    time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#creating-browsing-contexts:unsafe-shared-current-time
    x-internal="unsafe-shared-current-time"}.

3.  Let `creatorOrigin`{.variable} be null.

4.  Let `creatorBaseURL`{.variable} be null.

5.  ::: {#creator-browsing-context}
    If `creator`{.variable} is non-null, then:

    1.  Set `creatorOrigin`{.variable} to `creator`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#creating-browsing-contexts:concept-document-origin
        x-internal="concept-document-origin"}.

    2.  Set `creatorBaseURL`{.variable} to `creator`{.variable}\'s
        [document base
        URL](urls-and-fetching.html#document-base-url){#creating-browsing-contexts:document-base-url}.

    3.  Set `browsingContext`{.variable}\'s [virtual browsing context
        group
        ID](#virtual-browsing-context-group-id){#creating-browsing-contexts:virtual-browsing-context-group-id}
        to `creator`{.variable}\'s [browsing
        context](#concept-document-bc){#creating-browsing-contexts:concept-document-bc}\'s
        [top-level browsing
        context](#bc-tlbc){#creating-browsing-contexts:bc-tlbc}\'s
        [virtual browsing context group
        ID](#virtual-browsing-context-group-id){#creating-browsing-contexts:virtual-browsing-context-group-id-2}.
    :::

6.  Let `sandboxFlags`{.variable} be the result of [determining the
    creation sandboxing
    flags](browsers.html#determining-the-creation-sandboxing-flags){#creating-browsing-contexts:determining-the-creation-sandboxing-flags}
    given `browsingContext`{.variable} and `embedder`{.variable}.

7.  [Let `origin`{.variable} be the result of [determining the
    origin](#determining-the-origin){#creating-browsing-contexts:determining-the-origin}
    given
    [`about:blank`{#creating-browsing-contexts:about:blank}](infrastructure.html#about:blank),
    `sandboxFlags`{.variable}, and
    `creatorOrigin`{.variable}.]{#about-blank-origin}

8.  Let `permissionsPolicy`{.variable} be the result of [creating a
    permissions
    policy](https://w3c.github.io/webappsec-feature-policy/#create-for-navigable){#creating-browsing-contexts:creating-a-permissions-policy
    x-internal="creating-a-permissions-policy"} given
    `embedder`{.variable} and `origin`{.variable}.
    [\[PERMISSIONSPOLICY\]](references.html#refsPERMISSIONSPOLICY)

9.  Let `agent`{.variable} be the result of [obtaining a similar-origin
    window
    agent](webappapis.html#obtain-similar-origin-window-agent){#creating-browsing-contexts:obtain-similar-origin-window-agent}
    given `origin`{.variable}, `group`{.variable}, and false.

10. Let `realm execution context`{.variable} be the result of [creating
    a new
    realm](webappapis.html#creating-a-new-javascript-realm){#creating-browsing-contexts:creating-a-new-javascript-realm}
    given `agent`{.variable} and the following customizations:

    - For the global object, create a new
      [`Window`{#creating-browsing-contexts:window}](nav-history-apis.html#window)
      object.

    - For the global **this** binding, use
      `browsingContext`{.variable}\'s
      [`WindowProxy`{#creating-browsing-contexts:windowproxy}](nav-history-apis.html#windowproxy)
      object.

11. Let `topLevelCreationURL`{.variable} be
    [`about:blank`{#creating-browsing-contexts:about:blank-2}](infrastructure.html#about:blank)
    if `embedder`{.variable} is null; otherwise `embedder`{.variable}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#creating-browsing-contexts:relevant-settings-object}\'s
    [top-level creation
    URL](webappapis.html#concept-environment-top-level-creation-url){#creating-browsing-contexts:concept-environment-top-level-creation-url}.

12. Let `topLevelOrigin`{.variable} be `origin`{.variable} if
    `embedder`{.variable} is null; otherwise `embedder`{.variable}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#creating-browsing-contexts:relevant-settings-object-2}\'s
    [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#creating-browsing-contexts:concept-environment-top-level-origin}.

13. [Set up a window environment settings
    object](nav-history-apis.html#set-up-a-window-environment-settings-object){#creating-browsing-contexts:set-up-a-window-environment-settings-object}
    with
    [`about:blank`{#creating-browsing-contexts:about:blank-3}](infrastructure.html#about:blank),
    `realm execution context`{.variable}, null,
    `topLevelCreationURL`{.variable}, and `topLevelOrigin`{.variable}.

14. Let `loadTimingInfo`{.variable} be a new [document load timing
    info](dom.html#document-load-timing-info){#creating-browsing-contexts:document-load-timing-info}
    with its [navigation start
    time](dom.html#navigation-start-time){#creating-browsing-contexts:navigation-start-time}
    set to the result of calling [coarsen
    time](https://w3c.github.io/hr-time/#dfn-coarsen-time){#creating-browsing-contexts:coarsen-time
    x-internal="coarsen-time"} with
    `unsafeContextCreationTime`{.variable} and the new [environment
    settings
    object](webappapis.html#environment-settings-object){#creating-browsing-contexts:environment-settings-object}\'s
    [cross-origin isolated
    capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#creating-browsing-contexts:concept-settings-object-cross-origin-isolated-capability}.

15. Let `document`{.variable} be a new
    [`Document`{#creating-browsing-contexts:document-2}](dom.html#document),
    with:

    [type](https://dom.spec.whatwg.org/#concept-document-type){#creating-browsing-contexts:concept-document-type x-internal="concept-document-type"}
    :   \"`html`\"

    [content type](https://dom.spec.whatwg.org/#concept-document-content-type){#creating-browsing-contexts:concept-document-content-type x-internal="concept-document-content-type"}
    :   \"[`text/html`{#creating-browsing-contexts:text/html}](iana.html#text/html)\"

    [mode](https://dom.spec.whatwg.org/#concept-document-mode){#creating-browsing-contexts:concept-document-mode x-internal="concept-document-mode"}
    :   \"`quirks`\"

    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#creating-browsing-contexts:concept-document-origin-2 x-internal="concept-document-origin"}
    :   `origin`{.variable}

    [browsing context](#concept-document-bc){#creating-browsing-contexts:concept-document-bc-2}
    :   `browsingContext`{.variable}

    [permissions policy](dom.html#concept-document-permissions-policy){#creating-browsing-contexts:concept-document-permissions-policy}
    :   `permissionsPolicy`{.variable}

    [active sandboxing flag set](browsers.html#active-sandboxing-flag-set){#creating-browsing-contexts:active-sandboxing-flag-set}
    :   `sandboxFlags`{.variable}

    [load timing info](dom.html#load-timing-info){#creating-browsing-contexts:load-timing-info}
    :   `loadTimingInfo`{.variable}

    [is initial `about:blank`](dom.html#is-initial-about:blank){#creating-browsing-contexts:is-initial-about:blank}
    :   true

    [about base URL](dom.html#concept-document-about-base-url){#creating-browsing-contexts:concept-document-about-base-url}
    :   `creatorBaseURL`{.variable}

    [allow declarative shadow roots](https://dom.spec.whatwg.org/#concept-document-allow-declarative-shadow-roots){#creating-browsing-contexts:concept-document-allow-declarative-shadow-roots x-internal="concept-document-allow-declarative-shadow-roots"}
    :   true

    [custom element registry](https://dom.spec.whatwg.org/#document-custom-element-registry){#creating-browsing-contexts:document-custom-element-registry x-internal="document-custom-element-registry"}
    :   a new
        [`CustomElementRegistry`{#creating-browsing-contexts:customelementregistry}](custom-elements.html#customelementregistry)
        object

16. If `creator`{.variable} is non-null, then:

    1.  Set `document`{.variable}\'s
        [referrer](dom.html#the-document's-referrer){#creating-browsing-contexts:the-document's-referrer}
        to the
        [serialization](https://url.spec.whatwg.org/#concept-url-serializer){#creating-browsing-contexts:concept-url-serializer
        x-internal="concept-url-serializer"} of `creator`{.variable}\'s
        [URL](https://dom.spec.whatwg.org/#concept-document-url){#creating-browsing-contexts:the-document's-address
        x-internal="the-document's-address"}.

    2.  Set `document`{.variable}\'s [policy
        container](dom.html#concept-document-policy-container){#creating-browsing-contexts:concept-document-policy-container}
        to a
        [clone](browsers.html#clone-a-policy-container){#creating-browsing-contexts:clone-a-policy-container}
        of `creator`{.variable}\'s [policy
        container](dom.html#concept-document-policy-container){#creating-browsing-contexts:concept-document-policy-container-2}.

    3.  If `creator`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#creating-browsing-contexts:concept-document-origin-3
        x-internal="concept-document-origin"} is [same
        origin](browsers.html#same-origin){#creating-browsing-contexts:same-origin}
        with `creator`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#creating-browsing-contexts:relevant-settings-object-3}\'s
        [top-level
        origin](webappapis.html#concept-environment-top-level-origin){#creating-browsing-contexts:concept-environment-top-level-origin-2},
        then set `document`{.variable}\'s [opener
        policy](dom.html#concept-document-coop){#creating-browsing-contexts:concept-document-coop}
        to `creator`{.variable}\'s [browsing
        context](#concept-document-bc){#creating-browsing-contexts:concept-document-bc-3}\'s
        [top-level browsing
        context](#bc-tlbc){#creating-browsing-contexts:bc-tlbc-2}\'s
        [active
        document](#active-document){#creating-browsing-contexts:active-document}\'s
        [opener
        policy](dom.html#concept-document-coop){#creating-browsing-contexts:concept-document-coop-2}.

17. [Assert](https://infra.spec.whatwg.org/#assert){#creating-browsing-contexts:assert
    x-internal="assert"}: `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#creating-browsing-contexts:the-document's-address-2
    x-internal="the-document's-address"} and `document`{.variable}\'s
    [relevant settings
    object](webappapis.html#relevant-settings-object){#creating-browsing-contexts:relevant-settings-object-4}\'s
    [creation
    URL](webappapis.html#concept-environment-creation-url){#creating-browsing-contexts:concept-environment-creation-url}
    are
    [`about:blank`{#creating-browsing-contexts:about:blank-4}](infrastructure.html#about:blank).

18. Mark `document`{.variable} as [ready for post-load
    tasks](parsing.html#ready-for-post-load-tasks){#creating-browsing-contexts:ready-for-post-load-tasks}.

19. [Populate with
    `html`/`head`/`body`](document-lifecycle.html#populate-with-html/head/body){#creating-browsing-contexts:populate-with-html/head/body}
    given `document`{.variable}.

20. [Make
    active](browsing-the-web.html#make-active){#creating-browsing-contexts:make-active}
    `document`{.variable}.

21. [Completely finish
    loading](document-lifecycle.html#completely-finish-loading){#creating-browsing-contexts:completely-finish-loading}
    `document`{.variable}.

22. Return `browsingContext`{.variable} and `document`{.variable}.

To [create a new top-level browsing context and
document]{#creating-a-new-top-level-browsing-context .dfn}:

1.  Let `group`{.variable} and `document`{.variable} be the result of
    [creating a new browsing context group and
    document](#creating-a-new-browsing-context-group-and-document){#creating-browsing-contexts:creating-a-new-browsing-context-group-and-document}.

2.  Return `group`{.variable}\'s [browsing context
    set](#browsing-context-set){#creating-browsing-contexts:browsing-context-set}\[0\]
    and `document`{.variable}.

To [create a new auxiliary browsing context and
document]{#creating-a-new-auxiliary-browsing-context .dfn}, given a
[browsing
context](#browsing-context){#creating-browsing-contexts:browsing-context-2}
`opener`{.variable}:

1.  Let `openerTopLevelBrowsingContext`{.variable} be
    `opener`{.variable}\'s [top-level
    traversable](#bc-traversable){#creating-browsing-contexts:bc-traversable}\'s
    [active browsing
    context](#nav-bc){#creating-browsing-contexts:nav-bc}.

2.  Let `group`{.variable} be
    `openerTopLevelBrowsingContext`{.variable}\'s
    [group](#tlbc-group){#creating-browsing-contexts:tlbc-group}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#creating-browsing-contexts:assert-2
    x-internal="assert"}: `group`{.variable} is non-null, as
    [navigating](browsing-the-web.html#navigate){#creating-browsing-contexts:navigate}
    invokes this directly.

4.  Let `browsingContext`{.variable} and `document`{.variable} be the
    result of [creating a new browsing context and
    document](#creating-a-new-browsing-context){#creating-browsing-contexts:creating-a-new-browsing-context}
    with `opener`{.variable}\'s [active
    document](#nav-document){#creating-browsing-contexts:nav-document},
    null, and `group`{.variable}.

5.  Set `browsingContext`{.variable}\'s [is
    auxiliary](#is-auxiliary){#creating-browsing-contexts:is-auxiliary}
    to true.

6.  [Append](#bcg-append){#creating-browsing-contexts:bcg-append}
    `browsingContext`{.variable} to `group`{.variable}.

7.  Set `browsingContext`{.variable}\'s [opener browsing
    context](#opener-browsing-context){#creating-browsing-contexts:opener-browsing-context}
    to `opener`{.variable}.

8.  Set `browsingContext`{.variable}\'s [virtual browsing context group
    ID](#virtual-browsing-context-group-id){#creating-browsing-contexts:virtual-browsing-context-group-id-3}
    to `openerTopLevelBrowsingContext`{.variable}\'s [virtual browsing
    context group
    ID](#virtual-browsing-context-group-id){#creating-browsing-contexts:virtual-browsing-context-group-id-4}.

9.  Set `browsingContext`{.variable}\'s [opener origin at
    creation](#opener-origin-at-creation){#creating-browsing-contexts:opener-origin-at-creation}
    to `opener`{.variable}\'s [active
    document](#nav-document){#creating-browsing-contexts:nav-document-2}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#creating-browsing-contexts:concept-document-origin-4
    x-internal="concept-document-origin"}.

10. Return `browsingContext`{.variable} and `document`{.variable}.

To [determine the origin]{#determining-the-origin .dfn}, given a
[URL](https://url.spec.whatwg.org/#concept-url){#creating-browsing-contexts:url
x-internal="url"} `url`{.variable}, a [sandboxing flag
set](browsers.html#sandboxing-flag-set){#creating-browsing-contexts:sandboxing-flag-set}
`sandboxFlags`{.variable}, and an
[origin](browsers.html#concept-origin){#creating-browsing-contexts:concept-origin}-or-null
`sourceOrigin`{.variable}:

1.  [If `sandboxFlags`{.variable} has its [sandboxed origin browsing
    context
    flag](browsers.html#sandboxed-origin-browsing-context-flag){#creating-browsing-contexts:sandboxed-origin-browsing-context-flag}
    set, then return a new [opaque
    origin](browsers.html#concept-origin-opaque){#creating-browsing-contexts:concept-origin-opaque}.]{#sandboxOrigin}

2.  If `url`{.variable} is null, then return a new [opaque
    origin](browsers.html#concept-origin-opaque){#creating-browsing-contexts:concept-origin-opaque-2}.

3.  If `url`{.variable} is
    [`about:srcdoc`{#creating-browsing-contexts:about:srcdoc}](urls-and-fetching.html#about:srcdoc),
    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#creating-browsing-contexts:assert-3
        x-internal="assert"}: `sourceOrigin`{.variable} is non-null.

    2.  Return `sourceOrigin`{.variable}.

4.  If `url`{.variable} [matches
    `about:blank`](urls-and-fetching.html#matches-about:blank){#creating-browsing-contexts:matches-about:blank}
    and `sourceOrigin`{.variable} is non-null, then return
    `sourceOrigin`{.variable}.

5.  Return `url`{.variable}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#creating-browsing-contexts:concept-url-origin
    x-internal="concept-url-origin"}.

The cases that return `sourceOrigin`{.variable} result in two
[`Document`{#creating-browsing-contexts:document-3}](dom.html#document)s
that end up with the same underlying
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#creating-browsing-contexts:concept-document-origin-5
x-internal="concept-document-origin"}, meaning that
[`document.domain`{#creating-browsing-contexts:dom-document-domain}](browsers.html#dom-document-domain)
affects both.

##### [7.3.2.2]{.secno} Related browsing contexts[](#nested-browsing-contexts){.self-link} {#nested-browsing-contexts}

A [browsing
context](#browsing-context){#nested-browsing-contexts:browsing-context}
`potentialDescendant`{.variable} is said to be an
[ancestor]{#ancestor-browsing-context .dfn} of a browsing context
`potentialAncestor`{.variable} if the following algorithm returns true:

1.  Let `potentialDescendantDocument`{.variable} be
    `potentialDescendant`{.variable}\'s [active
    document](#active-document){#nested-browsing-contexts:active-document}.

2.  If `potentialDescendantDocument`{.variable} is not [fully
    active](#fully-active){#nested-browsing-contexts:fully-active}, then
    return false.

3.  Let `ancestorBCs`{.variable} be the list obtained by taking the
    [browsing
    context](#concept-document-bc){#nested-browsing-contexts:concept-document-bc}
    of the [active
    document](#nav-document){#nested-browsing-contexts:nav-document} of
    each member of `potentialDescendantDocument`{.variable}\'s [ancestor
    navigables](#ancestor-navigables){#nested-browsing-contexts:ancestor-navigables}.

4.  If `ancestorBCs`{.variable}
    [contains](https://infra.spec.whatwg.org/#list-contain){#nested-browsing-contexts:list-contains
    x-internal="list-contains"} `potentialAncestor`{.variable}, then
    return true.

5.  Return false.

A [top-level browsing context]{#top-level-browsing-context .dfn
export=""} is a [browsing
context](#browsing-context){#nested-browsing-contexts:browsing-context-2}
whose [active
document](#active-document){#nested-browsing-contexts:active-document-2}\'s
[node
navigable](#node-navigable){#nested-browsing-contexts:node-navigable} is
a [traversable
navigable](#traversable-navigable){#nested-browsing-contexts:traversable-navigable}.

It is *not* required to be a [top-level
traversable](#top-level-traversable){#nested-browsing-contexts:top-level-traversable}.

The [top-level browsing context]{#bc-tlbc .dfn} of a [browsing
context](#browsing-context){#nested-browsing-contexts:browsing-context-3}
`start`{.variable} is the result of the following algorithm:

1.  If `start`{.variable}\'s [active
    document](#active-document){#nested-browsing-contexts:active-document-3}
    is not [fully
    active](#fully-active){#nested-browsing-contexts:fully-active-2},
    then return null.

2.  Let `navigable`{.variable} be `start`{.variable}\'s [active
    document](#active-document){#nested-browsing-contexts:active-document-4}\'s
    [node
    navigable](#node-navigable){#nested-browsing-contexts:node-navigable-2}.

3.  While `navigable`{.variable}\'s
    [parent](#nav-parent){#nested-browsing-contexts:nav-parent} is not
    null, set `navigable`{.variable} to `navigable`{.variable}\'s
    [parent](#nav-parent){#nested-browsing-contexts:nav-parent-2}.

4.  Return `navigable`{.variable}\'s [active browsing
    context](#nav-bc){#nested-browsing-contexts:nav-bc}.

::: {#warning-avoid-related-bc-terms .warning}
The terms [ancestor browsing
context](#ancestor-browsing-context){#nested-browsing-contexts:ancestor-browsing-context}
and [top-level browsing
context](#top-level-browsing-context){#nested-browsing-contexts:top-level-browsing-context}
are rarely useful, since [browsing
contexts](#browsing-context){#nested-browsing-contexts:browsing-context-4}
in general are [usually the inappropriate specification concept to
use](#warning-avoid-using-bcs). Note in particular that when a [browsing
context](#browsing-context){#nested-browsing-contexts:browsing-context-5}\'s
[active
document](#active-document){#nested-browsing-contexts:active-document-5}
is not [fully
active](#fully-active){#nested-browsing-contexts:fully-active-3}, it
never counts as an ancestor or top-level browsing context, and as such
these concepts are not useful when
[bfcache](browsing-the-web.html#note-bfcache) is in play.

Instead, use concepts such as the [ancestor
navigables](#ancestor-navigables){#nested-browsing-contexts:ancestor-navigables-2}
collection, the [parent
navigable](#nav-parent){#nested-browsing-contexts:nav-parent-3}, or a
navigable\'s [top-level
traversable](#nav-traversable){#nested-browsing-contexts:nav-traversable}.
:::

------------------------------------------------------------------------

[]{#security-nav}A [browsing
context](#browsing-context){#nested-browsing-contexts:browsing-context-6}
`A`{.variable} is [familiar with]{#familiar-with .dfn} a second
[browsing
context](#browsing-context){#nested-browsing-contexts:browsing-context-7}
`B`{.variable} if the following algorithm returns true:

1.  If `A`{.variable}\'s [active
    document](#active-document){#nested-browsing-contexts:active-document-6}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#nested-browsing-contexts:concept-document-origin
    x-internal="concept-document-origin"} is [same
    origin](browsers.html#same-origin){#nested-browsing-contexts:same-origin}
    with `B`{.variable}\'s [active
    document](#active-document){#nested-browsing-contexts:active-document-7}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#nested-browsing-contexts:concept-document-origin-2
    x-internal="concept-document-origin"}, then return true.

2.  If `A`{.variable}\'s [top-level browsing
    context](#bc-tlbc){#nested-browsing-contexts:bc-tlbc} is
    `B`{.variable}, then return true.

3.  If `B`{.variable} is an [auxiliary browsing
    context](#auxiliary-browsing-context){#nested-browsing-contexts:auxiliary-browsing-context}
    and `A`{.variable} is [familiar
    with](#familiar-with){#nested-browsing-contexts:familiar-with}
    `B`{.variable}\'s [opener browsing
    context](#opener-browsing-context){#nested-browsing-contexts:opener-browsing-context},
    then return true.

4.  If there exists an [ancestor browsing
    context](#ancestor-browsing-context){#nested-browsing-contexts:ancestor-browsing-context-2}
    of `B`{.variable} whose [active
    document](#active-document){#nested-browsing-contexts:active-document-8}
    has the
    [same](browsers.html#same-origin){#nested-browsing-contexts:same-origin-2}
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#nested-browsing-contexts:concept-document-origin-3
    x-internal="concept-document-origin"} as the [active
    document](#active-document){#nested-browsing-contexts:active-document-9}
    of `A`{.variable}, then return true.

    This includes the case where `A`{.variable} is an [ancestor browsing
    context](#ancestor-browsing-context){#nested-browsing-contexts:ancestor-browsing-context-3}
    of `B`{.variable}.

5.  Return false.

##### [7.3.2.3]{.secno} Groupings of browsing contexts[](#groupings-of-browsing-contexts){.self-link}

A [top-level browsing
context](#top-level-browsing-context){#groupings-of-browsing-contexts:top-level-browsing-context}
has an associated [group]{#tlbc-group .dfn} (null or a [browsing context
group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group}).
It is initially null.

A user agent holds a [browsing context group
set]{#browsing-context-group-set .dfn} (a
[set](https://infra.spec.whatwg.org/#ordered-set){#groupings-of-browsing-contexts:set
x-internal="set"} of [browsing context
groups](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-2}).

A [browsing context group]{#browsing-context-group .dfn} holds a
[browsing context set]{#browsing-context-set .dfn} (a
[set](https://infra.spec.whatwg.org/#ordered-set){#groupings-of-browsing-contexts:set-2
x-internal="set"} of [top-level browsing
contexts](#top-level-browsing-context){#groupings-of-browsing-contexts:top-level-browsing-context-2}).

A [top-level browsing
context](#top-level-browsing-context){#groupings-of-browsing-contexts:top-level-browsing-context-3}
is added to the
[group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-3}
when the group is
[created](#creating-a-new-browsing-context-group-and-document){#groupings-of-browsing-contexts:creating-a-new-browsing-context-group-and-document}.
All subsequent [top-level browsing
contexts](#top-level-browsing-context){#groupings-of-browsing-contexts:top-level-browsing-context-4}
added to the
[group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-4}
will be [auxiliary browsing
contexts](#auxiliary-browsing-context){#groupings-of-browsing-contexts:auxiliary-browsing-context}.

A [browsing context
group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-5}
has an associated [agent cluster map]{#agent-cluster-map .dfn} (a weak
[map](https://infra.spec.whatwg.org/#ordered-map){#groupings-of-browsing-contexts:ordered-map
x-internal="ordered-map"} of [agent cluster
keys](webappapis.html#agent-cluster-key){#groupings-of-browsing-contexts:agent-cluster-key}
to [agent
clusters](https://tc39.es/ecma262/#sec-agent-clusters){#groupings-of-browsing-contexts:agent-cluster
x-internal="agent-cluster"}). User agents are responsible for collecting
agent clusters when it is deemed that nothing can access them anymore.

A [browsing context
group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-6}
has an associated [historical agent cluster key
map]{#historical-agent-cluster-key-map .dfn}, which is a
[map](https://infra.spec.whatwg.org/#ordered-map){#groupings-of-browsing-contexts:ordered-map-2
x-internal="ordered-map"} of
[origins](browsers.html#concept-origin){#groupings-of-browsing-contexts:concept-origin}
to [agent cluster
keys](webappapis.html#agent-cluster-key){#groupings-of-browsing-contexts:agent-cluster-key-2}.
This map is used to ensure the consistency of the [origin-keyed agent
clusters](browsers.html#origin-keyed-agent-clusters) feature by
recording what agent cluster keys were previously used for a given
origin.

The [historical agent cluster key
map](#historical-agent-cluster-key-map){#groupings-of-browsing-contexts:historical-agent-cluster-key-map}
only ever gains entries over the lifetime of the browsing context group.

A [browsing context
group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-7}
has a [cross-origin isolation mode]{#bcg-cross-origin-isolation .dfn},
which is a [cross-origin isolation
mode](#cross-origin-isolation-mode){#groupings-of-browsing-contexts:cross-origin-isolation-mode}.
It is initially
\"[`none`{#groupings-of-browsing-contexts:cross-origin-isolation-none}](#cross-origin-isolation-none)\".

A [cross-origin isolation mode]{#cross-origin-isolation-mode .dfn} is
one of three possible values: \"[`none`]{#cross-origin-isolation-none
.dfn}\", \"[`logical`]{#cross-origin-isolation-logical .dfn}\", or
\"[`concrete`]{#cross-origin-isolation-concrete .dfn}\".

::: note
\"[`logical`{#groupings-of-browsing-contexts:cross-origin-isolation-logical}](#cross-origin-isolation-logical)\"
and
\"[`concrete`{#groupings-of-browsing-contexts:cross-origin-isolation-concrete}](#cross-origin-isolation-concrete)\"
are similar. They are both used for [browsing context
groups](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-8}
where:

- every top-level
  [`Document`{#groupings-of-browsing-contexts:document}](dom.html#document)
  has
  \`[`Cross-Origin-Opener-Policy`](browsers.html#cross-origin-opener-policy-2){#groupings-of-browsing-contexts:cross-origin-opener-policy-2}`: `[`same-origin`](browsers.html#coop-same-origin){#groupings-of-browsing-contexts:coop-same-origin}\`,
  and

- every
  [`Document`{#groupings-of-browsing-contexts:document-2}](dom.html#document)
  has a
  \`[`Cross-Origin-Embedder-Policy`{#groupings-of-browsing-contexts:cross-origin-embedder-policy}](browsers.html#cross-origin-embedder-policy)\`
  header whose value is [compatible with cross-origin
  isolation](browsers.html#compatible-with-cross-origin-isolation){#groupings-of-browsing-contexts:compatible-with-cross-origin-isolation}.

On some platforms, it is difficult to provide the security properties
required to grant safe access to the APIs gated by the [cross-origin
isolated
capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#groupings-of-browsing-contexts:concept-settings-object-cross-origin-isolated-capability}.
As a result, only
\"[`concrete`{#groupings-of-browsing-contexts:cross-origin-isolation-concrete-2}](#cross-origin-isolation-concrete)\"
can grant access that capability.
\"[`logical`{#groupings-of-browsing-contexts:cross-origin-isolation-logical-2}](#cross-origin-isolation-logical)\"
is used on platform not supporting this capability, where various
restrictions imposed by cross-origin isolation will still apply, but the
capability is not granted.
:::

To [create a new browsing context group and
document]{#creating-a-new-browsing-context-group-and-document .dfn}:

1.  Let `group`{.variable} be a new [browsing context
    group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-9}.

2.  [Append](https://infra.spec.whatwg.org/#set-append){#groupings-of-browsing-contexts:set-append
    x-internal="set-append"} `group`{.variable} to the user agent\'s
    [browsing context group
    set](#browsing-context-group-set){#groupings-of-browsing-contexts:browsing-context-group-set}.

3.  Let `browsingContext`{.variable} and `document`{.variable} be the
    result of [creating a new browsing context and
    document](#creating-a-new-browsing-context){#groupings-of-browsing-contexts:creating-a-new-browsing-context}
    with null, null, and `group`{.variable}.

4.  [Append](#bcg-append){#groupings-of-browsing-contexts:bcg-append}
    `browsingContext`{.variable} to `group`{.variable}.

5.  Return `group`{.variable} and `document`{.variable}.

To [append]{#bcg-append .dfn} a [top-level browsing
context](#top-level-browsing-context){#groupings-of-browsing-contexts:top-level-browsing-context-5}
`browsingContext`{.variable} to a [browsing context
group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-10}
`group`{.variable}:

1.  [Append](https://infra.spec.whatwg.org/#set-append){#groupings-of-browsing-contexts:set-append-2
    x-internal="set-append"} `browsingContext`{.variable} to
    `group`{.variable}\'s [browsing context
    set](#browsing-context-set){#groupings-of-browsing-contexts:browsing-context-set}.

2.  Set `browsingContext`{.variable}\'s
    [group](#tlbc-group){#groupings-of-browsing-contexts:tlbc-group} to
    `group`{.variable}.

To [remove]{#bcg-remove .dfn} a [top-level browsing
context](#top-level-browsing-context){#groupings-of-browsing-contexts:top-level-browsing-context-6}
`browsingContext`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#groupings-of-browsing-contexts:assert
    x-internal="assert"}: `browsingContext`{.variable}\'s
    [group](#tlbc-group){#groupings-of-browsing-contexts:tlbc-group-2}
    is non-null.

2.  Let `group`{.variable} be `browsingContext`{.variable}\'s
    [group](#tlbc-group){#groupings-of-browsing-contexts:tlbc-group-3}.

3.  Set `browsingContext`{.variable}\'s
    [group](#tlbc-group){#groupings-of-browsing-contexts:tlbc-group-4}
    to null.

4.  [Remove](https://infra.spec.whatwg.org/#list-remove){#groupings-of-browsing-contexts:list-remove
    x-internal="list-remove"} `browsingContext`{.variable} from
    `group`{.variable}\'s [browsing context
    set](#browsing-context-set){#groupings-of-browsing-contexts:browsing-context-set-2}.

5.  If `group`{.variable}\'s [browsing context
    set](#browsing-context-set){#groupings-of-browsing-contexts:browsing-context-set-3}
    [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#groupings-of-browsing-contexts:list-is-empty
    x-internal="list-is-empty"}, then
    [remove](https://infra.spec.whatwg.org/#list-remove){#groupings-of-browsing-contexts:list-remove-2
    x-internal="list-remove"} `group`{.variable} from the user agent\'s
    [browsing context group
    set](#browsing-context-group-set){#groupings-of-browsing-contexts:browsing-context-group-set-2}.

[Append](#bcg-append){#groupings-of-browsing-contexts:bcg-append-2} and
[remove](#bcg-remove){#groupings-of-browsing-contexts:bcg-remove} are
primitive operations that help define the lifetime of a [browsing
context
group](#browsing-context-group){#groupings-of-browsing-contexts:browsing-context-group-11}.
They are called by higher-level creation and destruction operations for
[`Document`{#groupings-of-browsing-contexts:document-3}](dom.html#document)s
and [browsing
contexts](#browsing-context){#groupings-of-browsing-contexts:browsing-context}.

When there are no
[`Document`{#groupings-of-browsing-contexts:document-4}](dom.html#document)
objects whose [browsing
context](#concept-document-bc){#groupings-of-browsing-contexts:concept-document-bc}
equals a given [browsing
context](#browsing-context){#groupings-of-browsing-contexts:browsing-context-2}
(i.e., all such
[`Document`{#groupings-of-browsing-contexts:document-5}](dom.html#document)s
have been
[destroyed](document-lifecycle.html#destroy-a-document){#groupings-of-browsing-contexts:destroy-a-document}),
and that [browsing
context](#browsing-context){#groupings-of-browsing-contexts:browsing-context-3}\'s
[`WindowProxy`{#groupings-of-browsing-contexts:windowproxy}](nav-history-apis.html#windowproxy)
is eligible for garbage collection, then the [browsing
context](#browsing-context){#groupings-of-browsing-contexts:browsing-context-4}
will never be accessed again. If it is a [top-level browsing
context](#top-level-browsing-context){#groupings-of-browsing-contexts:top-level-browsing-context-7},
then at this point the user agent must
[remove](#bcg-remove){#groupings-of-browsing-contexts:bcg-remove-2} it.

#### [7.3.3]{.secno} Fully active documents[](#fully-active-documents){.self-link}

A [`Document`{#fully-active-documents:document}](dom.html#document)
`d`{.variable} is said to be [fully active]{#fully-active .dfn
dfn-for="Document" export=""} when `d`{.variable} is the [active
document](#nav-document){#fully-active-documents:nav-document} of a
[navigable](#navigable){#fully-active-documents:navigable}
`navigable`{.variable}, and either `navigable`{.variable} is a
[top-level
traversable](#top-level-traversable){#fully-active-documents:top-level-traversable}
or `navigable`{.variable}\'s [container
document](#nav-container-document){#fully-active-documents:nav-container-document}
is [fully active](#fully-active){#fully-active-documents:fully-active}.

Because they are associated with an element, [child
navigables](#child-navigable){#fully-active-documents:child-navigable}
are always tied to a specific
[`Document`{#fully-active-documents:document-2}](dom.html#document),
their [container
document](#nav-container-document){#fully-active-documents:nav-container-document-2},
in their [parent
navigable](#nav-parent){#fully-active-documents:nav-parent}. User agents
must not allow the user to interact with [child
navigables](#child-navigable){#fully-active-documents:child-navigable-2}
whose [container
documents](#nav-container-document){#fully-active-documents:nav-container-document-3}
are not themselves [fully
active](#fully-active){#fully-active-documents:fully-active-2}.

::: example
The following example illustrates how a
[`Document`{#fully-active-documents:document-3}](dom.html#document) can
be the [active
document](#nav-document){#fully-active-documents:nav-document-2} of its
[node
navigable](#node-navigable){#fully-active-documents:node-navigable},
while not being [fully
active](#fully-active){#fully-active-documents:fully-active-3}. Here
`a.html` is loaded into a browser window, `b-1.html` starts out loaded
into an
[`iframe`{#fully-active-documents:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
as shown, and `b-2.html` and `c.html` are omitted (they can simply be an
empty document).

``` html
<!-- a.html -->
<!DOCTYPE html>
<html lang="en">
<title>Navigable A</title>

<iframe src="b-1.html"></iframe>
<button onclick="frames[0].location.href = 'b-2.html'">Click me</button>

<!-- b-1.html -->
<!DOCTYPE html>
<html lang="en">
<title>Navigable B</title>

<iframe src="c.html"></iframe>
```

At this point, the documents given by `a.html`, `b-1.html`, and `c.html`
are all the [active
documents](#nav-document){#fully-active-documents:nav-document-3} of
their respective [node
navigables](#node-navigable){#fully-active-documents:node-navigable-2}.
They are also all [fully
active](#fully-active){#fully-active-documents:fully-active-4}.

After clicking on the
[`button`{#fully-active-documents:the-button-element}](form-elements.html#the-button-element),
and thus loading a new
[`Document`{#fully-active-documents:document-4}](dom.html#document) from
`b-2.html` into navigable B, we have the following results:

- The `a.html`
  [`Document`{#fully-active-documents:document-5}](dom.html#document)
  remains both the [active
  document](#nav-document){#fully-active-documents:nav-document-4} of
  navigable A, and [fully
  active](#fully-active){#fully-active-documents:fully-active-5}.

- The `b-1.html`
  [`Document`{#fully-active-documents:document-6}](dom.html#document) is
  now *not* the [active
  document](#nav-document){#fully-active-documents:nav-document-5} of
  navigable B. As such it is also not [fully
  active](#fully-active){#fully-active-documents:fully-active-6}.

- The new `b-2.html`
  [`Document`{#fully-active-documents:document-7}](dom.html#document) is
  now the [active
  document](#nav-document){#fully-active-documents:nav-document-6} of
  navigable B, and is also [fully
  active](#fully-active){#fully-active-documents:fully-active-7}.

- The `c.html`
  [`Document`{#fully-active-documents:document-8}](dom.html#document) is
  still the [active
  document](#nav-document){#fully-active-documents:nav-document-7} of
  navigable C. However, since C\'s [container
  document](#nav-container-document){#fully-active-documents:nav-container-document-4}
  is the `b-1.html`
  [`Document`{#fully-active-documents:document-9}](dom.html#document),
  which is itself not [fully
  active](#fully-active){#fully-active-documents:fully-active-8}, this
  means the `c.html`
  [`Document`{#fully-active-documents:document-10}](dom.html#document)
  is now not [fully
  active](#fully-active){#fully-active-documents:fully-active-9}.
:::

[← 7.2 APIs related to navigation and session
history](nav-history-apis.html) --- [Table of Contents](./) --- [7.4
Navigation and session history →](browsing-the-web.html)
