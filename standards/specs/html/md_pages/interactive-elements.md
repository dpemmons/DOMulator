HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.10.17 Form control
infrastructure](form-control-infrastructure.html) --- [Table of
Contents](./) --- [4.12 Scripting →](scripting.html)

1.  ::: {#toc-semantics}
    1.  [[4.11]{.secno} Interactive
        elements](interactive-elements.html#interactive-elements)
        1.  [[4.11.1]{.secno} The `details`
            element](interactive-elements.html#the-details-element)
        2.  [[4.11.2]{.secno} The `summary`
            element](interactive-elements.html#the-summary-element)
        3.  [[4.11.3]{.secno}
            Commands](interactive-elements.html#commands)
            1.  [[4.11.3.1]{.secno}
                Facets](interactive-elements.html#facets-2)
            2.  [[4.11.3.2]{.secno} Using the `a` element to define a
                command](interactive-elements.html#using-the-a-element-to-define-a-command)
            3.  [[4.11.3.3]{.secno} Using the `button` element to define
                a
                command](interactive-elements.html#using-the-button-element-to-define-a-command)
            4.  [[4.11.3.4]{.secno} Using the `input` element to define
                a
                command](interactive-elements.html#using-the-input-element-to-define-a-command)
            5.  [[4.11.3.5]{.secno} Using the `option` element to define
                a
                command](interactive-elements.html#using-the-option-element-to-define-a-command)
            6.  [[4.11.3.6]{.secno} Using the `accesskey` attribute on a
                `legend` element to define a
                command](interactive-elements.html#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command)
            7.  [[4.11.3.7]{.secno} Using the `accesskey` attribute to
                define a command on other
                elements](interactive-elements.html#using-the-accesskey-attribute-to-define-a-command-on-other-elements)
        4.  [[4.11.4]{.secno} The `dialog`
            element](interactive-elements.html#the-dialog-element)
        5.  [[4.11.5]{.secno} Dialog light
            dismiss](interactive-elements.html#dialog-light-dismiss)
    :::

### [4.11]{.secno} Interactive elements[](#interactive-elements){.self-link}

#### [4.11.1]{.secno} The [`details`]{.dfn dfn-type="element"} element[](#the-details-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/details](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details "The <details> HTML element creates a disclosure widget in which information is visible only when the widget is toggled into an "open" state. A summary or label must be provided using the <summary> element.")

Support in all current engines.

::: support
[Firefox49+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome12+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android49+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDetailsElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDetailsElement "The HTMLDetailsElement interface provides special properties (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating <details> elements.")

Support in all current engines.

::: support
[Firefox49+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-details-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-details-element:flow-content-2}.
:   [Interactive
    content](dom.html#interactive-content-2){#the-details-element:interactive-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-details-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-details-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-details-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-details-element:concept-element-content-model}:
:   One
    [`summary`{#the-details-element:the-summary-element}](#the-summary-element)
    element followed by [flow
    content](dom.html#flow-content-2){#the-details-element:flow-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-details-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-details-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-details-element:global-attributes}
:   [`name`{#the-details-element:attr-details-name}](#attr-details-name)
    --- Name of group of mutually-exclusive
    [`details`{#the-details-element:the-details-element}](#the-details-element)
    elements
:   [`open`{#the-details-element:attr-details-open}](#attr-details-open)
    --- Whether the details are visible

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-details-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-details).
:   [For implementers](https://w3c.github.io/html-aam/#el-details).

[DOM interface](dom.html#concept-element-dom){#the-details-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLDetailsElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString name;
      [CEReactions] attribute boolean open;
    };
    ```

The
[`details`{#the-details-element:the-details-element-2}](#the-details-element)
element
[represents](dom.html#represents){#the-details-element:represents} a
disclosure widget from which the user can obtain additional information
or controls.

As with all HTML elements, it is not conforming to use the
[`details`{#the-details-element:the-details-element-3}](#the-details-element)
element when attempting to represent another type of control. For
example, tab widgets and menu widgets are not disclosure widgets, so
abusing the
[`details`{#the-details-element:the-details-element-4}](#the-details-element)
element to implement these patterns is incorrect.

The
[`details`{#the-details-element:the-details-element-5}](#the-details-element)
element is not appropriate for footnotes. Please see [the section on
footnotes](semantics-other.html#footnotes) for details on how to mark up
footnotes.

The first
[`summary`{#the-details-element:the-summary-element-2}](#the-summary-element)
element child of the element, if any,
[represents](dom.html#represents){#the-details-element:represents-2} the
summary or legend of the details. If there is no child
[`summary`{#the-details-element:the-summary-element-3}](#the-summary-element)
element, the user agent should provide its own legend (e.g.
\"Details\").

The rest of the element\'s contents
[represents](dom.html#represents){#the-details-element:represents-3} the
additional information or controls.

The [`name`]{#attr-details-name .dfn dfn-for="details"
dfn-type="element-attr"} content attribute gives the name of the group
of related
[`details`{#the-details-element:the-details-element-6}](#the-details-element)
elements that the element is a member of. Opening one member of this
group causes other members of the group to close. If the attribute is
specified, its value must not be the empty string.

Before using this feature, authors should consider whether this grouping
of related
[`details`{#the-details-element:the-details-element-7}](#the-details-element)
elements into an exclusive accordion is helpful or harmful to users.
While using an exclusive accordion can reduce the maximum amount of
space that a set of content can occupy, it can also frustrate users who
have to open many items to find what they want or users who want to look
at the contents of multiple items at the same time.

A document must not contain more than one
[`details`{#the-details-element:the-details-element-8}](#the-details-element)
element in the same [details name
group](#details-name-group){#the-details-element:details-name-group}
that has the
[`open`{#the-details-element:attr-details-open-2}](#attr-details-open)
attribute present. Authors must not use script to add
[`details`{#the-details-element:the-details-element-9}](#the-details-element)
elements to a document in a way that would cause a [details name
group](#details-name-group){#the-details-element:details-name-group-2}
to have more than one
[`details`{#the-details-element:the-details-element-10}](#the-details-element)
element with the
[`open`{#the-details-element:attr-details-open-3}](#attr-details-open)
attribute present.

The group of elements that is created by a common
[`name`{#the-details-element:attr-details-name-2}](#attr-details-name)
attribute is exclusive, meaning that at most one of the
[`details`{#the-details-element:the-details-element-11}](#the-details-element)
elements can be open at once. While this exclusivity is enforced by user
agents, the resulting enforcement immediately changes the
[`open`{#the-details-element:attr-details-open-4}](#attr-details-open)
attributes in the markup. This requirement on authors forbids such
misleading markup.

A document must not contain a
[`details`{#the-details-element:the-details-element-12}](#the-details-element)
element that is a descendant of another
[`details`{#the-details-element:the-details-element-13}](#the-details-element)
element in the same [details name
group](#details-name-group){#the-details-element:details-name-group-3}.

Documents that use the
[`name`{#the-details-element:attr-details-name-3}](#attr-details-name)
attribute to group multiple related
[`details`{#the-details-element:the-details-element-14}](#the-details-element)
elements should keep those related elements together in a containing
element (such as a
[`section`{#the-details-element:the-section-element}](sections.html#the-section-element)
element or
[`article`{#the-details-element:the-article-element}](sections.html#the-article-element)
element). When it makes sense for the group to be introduced with a
heading, authors should put that heading in a
[heading](sections.html#concept-heading){#the-details-element:concept-heading}
element at the start of the containing element.

Visually and programmatically grouping related elements together can be
important for accessible user experiences. This can help users
understand the relationship between such elements. When related elements
are in disparate sections of a web page rather than being grouped, the
elements\' relationships to each other can be less discoverable or
understandable.

The [`open`]{#attr-details-open .dfn dfn-for="details"
dfn-type="element-attr"} content attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-details-element:boolean-attribute}.
If present, it indicates that both the summary and the additional
information is to be shown to the user. If the attribute is absent, only
the summary is to be shown.

When the element is created, if the attribute is absent, the additional
information should be hidden; if the attribute is present, that
information should be shown. Subsequently, if the attribute is removed,
then the information should be hidden; if the attribute is added, the
information should be shown.

The user agent should allow the user to request that the additional
information be shown or hidden. To honor a request for the details to be
shown, the user agent must
[set](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-details-element:concept-element-attributes-set-value
x-internal="concept-element-attributes-set-value"} the
[`open`{#the-details-element:attr-details-open-5}](#attr-details-open)
attribute on the element to the empty string. To honor a request for the
information to be hidden, the user agent must
[remove](https://dom.spec.whatwg.org/#concept-element-attributes-remove){#the-details-element:concept-element-attributes-remove
x-internal="concept-element-attributes-remove"} the
[`open`{#the-details-element:attr-details-open-6}](#attr-details-open)
attribute from the element.

This ability to request that additional information be shown or hidden
may simply be the [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-details-element:activation-behaviour
x-internal="activation-behaviour"} of the appropriate
[`summary`{#the-details-element:the-summary-element-4}](#the-summary-element)
element, in the case such an element exists. However, if no such element
exists, user agents can still provide this ability through some other
user interface affordance.

The [details name group]{#details-name-group .dfn} that contains a
[`details`{#the-details-element:the-details-element-15}](#the-details-element)
element `a`{.variable} also contains all the other
[`details`{#the-details-element:the-details-element-16}](#the-details-element)
elements `b`{.variable} that fulfill all of the following conditions:

- Both `a`{.variable} and `b`{.variable} are in the same
  [tree](https://dom.spec.whatwg.org/#concept-tree){#the-details-element:tree
  x-internal="tree"}.
- They both have a
  [`name`{#the-details-element:attr-details-name-4}](#attr-details-name)
  attribute, their
  [`name`{#the-details-element:attr-details-name-5}](#attr-details-name)
  attributes are not the empty string, and the value of
  `a`{.variable}\'s
  [`name`{#the-details-element:attr-details-name-6}](#attr-details-name)
  attribute equals the value of `b`{.variable}\'s
  [`name`{#the-details-element:attr-details-name-7}](#attr-details-name)
  attribute.

Every
[`details`{#the-details-element:the-details-element-17}](#the-details-element)
element has a [details toggle task tracker]{#details-toggle-task-tracker
.dfn}, which is a [toggle task
tracker](interaction.html#toggle-task-tracker){#the-details-element:toggle-task-tracker}
or null, initially null.

The following [attribute change
steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#the-details-element:concept-element-attributes-change-ext
x-internal="concept-element-attributes-change-ext"}, given
`element`{.variable}, `localName`{.variable}, `oldValue`{.variable},
`value`{.variable}, and `namespace`{.variable}, are used for all
[`details`{#the-details-element:the-details-element-18}](#the-details-element)
elements:

1.  If `namespace`{.variable} is not null, then return.

2.  If `localName`{.variable} is
    [`name`{#the-details-element:attr-details-name-8}](#attr-details-name),
    then [ensure details exclusivity by closing the given element if
    needed](#ensure-details-exclusivity-by-closing-the-given-element-if-needed){#the-details-element:ensure-details-exclusivity-by-closing-the-given-element-if-needed}
    given `element`{.variable}.

3.  If `localName`{.variable} is
    [`open`{#the-details-element:attr-details-open-7}](#attr-details-open),
    then:

    1.  If one of `oldValue`{.variable} or `value`{.variable} is null
        and the other is not null, run the following steps, which are
        known as the [details notification task
        steps]{#details-notification-task-steps .dfn}, for this
        [`details`{#the-details-element:the-details-element-19}](#the-details-element)
        element:

        When the
        [`open`{#the-details-element:attr-details-open-8}](#attr-details-open)
        attribute is toggled several times in succession, the resulting
        tasks essentially get coalesced so that only one event is fired.

        1.  If `oldValue`{.variable} is null, [queue a details toggle
            event
            task](#queue-a-details-toggle-event-task){#the-details-element:queue-a-details-toggle-event-task}
            given the
            [`details`{#the-details-element:the-details-element-20}](#the-details-element)
            element, \"`closed`\", and \"`open`\".

        2.  Otherwise, [queue a details toggle event
            task](#queue-a-details-toggle-event-task){#the-details-element:queue-a-details-toggle-event-task-2}
            given the
            [`details`{#the-details-element:the-details-element-21}](#the-details-element)
            element, \"`open`\", and \"`closed`\".

    2.  If `oldValue`{.variable} is null and `value`{.variable} is not
        null, then [ensure details exclusivity by closing other elements
        if
        needed](#ensure-details-exclusivity-by-closing-other-elements-if-needed){#the-details-element:ensure-details-exclusivity-by-closing-other-elements-if-needed}
        given `element`{.variable}.

The
[`details`{#the-details-element:the-details-element-22}](#the-details-element)
[HTML element insertion
steps](infrastructure.html#html-element-insertion-steps){#the-details-element:html-element-insertion-steps},
given `insertedNode`{.variable}, are:

1.  [Ensure details exclusivity by closing the given element if
    needed](#ensure-details-exclusivity-by-closing-the-given-element-if-needed){#the-details-element:ensure-details-exclusivity-by-closing-the-given-element-if-needed-2}
    given `insertedNode`{.variable}.

To be clear, these attribute change and insertion steps also run when an
attribute or element is inserted via the parser.

To [queue a details toggle event
task]{#queue-a-details-toggle-event-task .dfn} given a
[`details`{#the-details-element:the-details-element-23}](#the-details-element)
element `element`{.variable}, a string `oldState`{.variable}, and a
string `newState`{.variable}:

1.  If `element`{.variable}\'s [details toggle task
    tracker](#details-toggle-task-tracker){#the-details-element:details-toggle-task-tracker}
    is not null, then:

    1.  Set `oldState`{.variable} to `element`{.variable}\'s [details
        toggle task
        tracker](#details-toggle-task-tracker){#the-details-element:details-toggle-task-tracker-2}\'s
        [old
        state](interaction.html#toggle-task-old-state){#the-details-element:toggle-task-old-state}.

    2.  Remove `element`{.variable}\'s [details toggle task
        tracker](#details-toggle-task-tracker){#the-details-element:details-toggle-task-tracker-3}\'s
        [task](interaction.html#toggle-task-task){#the-details-element:toggle-task-task}
        from its [task
        queue](webappapis.html#task-queue){#the-details-element:task-queue}.

    3.  Set `element`{.variable}\'s [details toggle task
        tracker](#details-toggle-task-tracker){#the-details-element:details-toggle-task-tracker-4}
        to null.

2.  [Queue an element
    task](webappapis.html#queue-an-element-task){#the-details-element:queue-an-element-task}
    given the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#the-details-element:dom-manipulation-task-source}
    and `element`{.variable} to run the following steps:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-details-element:concept-event-fire
        x-internal="concept-event-fire"} named
        [`toggle`{#the-details-element:event-toggle}](indices.html#event-toggle)
        at `element`{.variable}, using
        [`ToggleEvent`{#the-details-element:toggleevent}](interaction.html#toggleevent),
        with the
        [`oldState`{#the-details-element:dom-toggleevent-oldstate}](interaction.html#dom-toggleevent-oldstate)
        attribute initialized to `oldState`{.variable} and the
        [`newState`{#the-details-element:dom-toggleevent-newstate}](interaction.html#dom-toggleevent-newstate)
        attribute initialized to `newState`{.variable}.

    2.  Set `element`{.variable}\'s [details toggle task
        tracker](#details-toggle-task-tracker){#the-details-element:details-toggle-task-tracker-5}
        to null.

3.  Set `element`{.variable}\'s [details toggle task
    tracker](#details-toggle-task-tracker){#the-details-element:details-toggle-task-tracker-6}
    to a struct with
    [task](interaction.html#toggle-task-task){#the-details-element:toggle-task-task-2}
    set to the just-queued
    [task](webappapis.html#concept-task){#the-details-element:concept-task}
    and [old
    state](interaction.html#toggle-task-old-state){#the-details-element:toggle-task-old-state-2}
    set to `oldState`{.variable}.

To [ensure details exclusivity by closing other elements if
needed]{#ensure-details-exclusivity-by-closing-other-elements-if-needed
.dfn} given a
[`details`{#the-details-element:the-details-element-24}](#the-details-element)
element `element`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-details-element:assert
    x-internal="assert"}: `element`{.variable} has an
    [`open`{#the-details-element:attr-details-open-9}](#attr-details-open)
    attribute.

2.  If `element`{.variable} does not have a
    [`name`{#the-details-element:attr-details-name-9}](#attr-details-name)
    attribute, or its
    [`name`{#the-details-element:attr-details-name-10}](#attr-details-name)
    attribute is the empty string, then return.

3.  Let `groupMembers`{.variable} be a list of elements, containing all
    elements in `element`{.variable}\'s [details name
    group](#details-name-group){#the-details-element:details-name-group-4}
    except for `element`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-details-element:tree-order
    x-internal="tree-order"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-details-element:list-iterate
    x-internal="list-iterate"} element `otherElement`{.variable} of
    `groupMembers`{.variable}:

    1.  If the
        [`open`{#the-details-element:attr-details-open-10}](#attr-details-open)
        attribute is set on `otherElement`{.variable}, then:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-details-element:assert-2
            x-internal="assert"}: `otherElement`{.variable} is the only
            element in `groupMembers`{.variable} that has the
            [`open`{#the-details-element:attr-details-open-11}](#attr-details-open)
            attribute set.

        2.  [Remove](https://dom.spec.whatwg.org/#concept-element-attributes-remove){#the-details-element:concept-element-attributes-remove-2
            x-internal="concept-element-attributes-remove"} the
            [`open`{#the-details-element:attr-details-open-12}](#attr-details-open)
            attribute on `otherElement`{.variable}.

        3.  [Break](https://infra.spec.whatwg.org/#iteration-break){#the-details-element:break
            x-internal="break"}.

To [ensure details exclusivity by closing the given element if
needed]{#ensure-details-exclusivity-by-closing-the-given-element-if-needed
.dfn} given a
[`details`{#the-details-element:the-details-element-25}](#the-details-element)
element `element`{.variable}:

1.  If `element`{.variable} does not have an
    [`open`{#the-details-element:attr-details-open-13}](#attr-details-open)
    attribute, then return.

2.  If `element`{.variable} does not have a
    [`name`{#the-details-element:attr-details-name-11}](#attr-details-name)
    attribute, or its
    [`name`{#the-details-element:attr-details-name-12}](#attr-details-name)
    attribute is the empty string, then return.

3.  Let `groupMembers`{.variable} be a list of elements, containing all
    elements in `element`{.variable}\'s [details name
    group](#details-name-group){#the-details-element:details-name-group-5}
    except for `element`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-details-element:tree-order-2
    x-internal="tree-order"}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-details-element:list-iterate-2
    x-internal="list-iterate"} element `otherElement`{.variable} of
    `groupMembers`{.variable}:

    1.  If the
        [`open`{#the-details-element:attr-details-open-14}](#attr-details-open)
        attribute is set on `otherElement`{.variable}, then:

        1.  [Remove](https://dom.spec.whatwg.org/#concept-element-attributes-remove){#the-details-element:concept-element-attributes-remove-3
            x-internal="concept-element-attributes-remove"} the
            [`open`{#the-details-element:attr-details-open-15}](#attr-details-open)
            attribute on `element`{.variable}.

        2.  [Break](https://infra.spec.whatwg.org/#iteration-break){#the-details-element:break-2
            x-internal="break"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDetailsElement/open](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDetailsElement/open "The open property of the HTMLDetailsElement interface is a boolean value reflecting the open HTML attribute, indicating whether the <details>'s contents (not counting the <summary>) is to be shown to the user.")

Support in all current engines.

::: support
[Firefox49+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`name`]{#dom-details-name .dfn dfn-for="HTMLDetailsElement"
dfn-type="attribute"} and [`open`]{#dom-details-open .dfn
dfn-for="HTMLDetailsElement" dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-details-element:reflect}
the respective content attributes of the same name.

The [ancestor details revealing
algorithm]{#ancestor-details-revealing-algorithm .dfn} is to run the
following steps on `currentNode`{.variable}:

1.  While `currentNode`{.variable} has a parent node within the [flat
    tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-details-element:flat-tree
    x-internal="flat-tree"}:

    1.  If `currentNode`{.variable} is slotted into the second slot of a
        [`details`{#the-details-element:the-details-element-26}](#the-details-element)
        element:

        1.  Set `currentNode`{.variable} to the
            [`details`{#the-details-element:the-details-element-27}](#the-details-element)
            element which `currentNode`{.variable} is slotted into.

        2.  If the
            [`open`{#the-details-element:attr-details-open-16}](#attr-details-open)
            attribute is not set on `currentNode`{.variable}, then
            [set](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-details-element:concept-element-attributes-set-value-2
            x-internal="concept-element-attributes-set-value"} the
            [`open`{#the-details-element:attr-details-open-17}](#attr-details-open)
            attribute on `currentNode`{.variable} to the empty string.

    2.  Otherwise, set `currentNode`{.variable} to the parent node of
        `currentNode`{.variable} within the [flat
        tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-details-element:flat-tree-2
        x-internal="flat-tree"}.

::: example
The following example shows the
[`details`{#the-details-element:the-details-element-28}](#the-details-element)
element being used to hide technical details in a progress report.

``` html
<section class="progress window">
 <h1>Copying "Really Achieving Your Childhood Dreams"</h1>
 <details>
  <summary>Copying... <progress max="375505392" value="97543282"></progress> 25%</summary>
  <dl>
   <dt>Transfer rate:</dt> <dd>452KB/s</dd>
   <dt>Local filename:</dt> <dd>/home/rpausch/raycd.m4v</dd>
   <dt>Remote filename:</dt> <dd>/var/www/lectures/raycd.m4v</dd>
   <dt>Duration:</dt> <dd>01:16:27</dd>
   <dt>Color profile:</dt> <dd>SD (6-1-6)</dd>
   <dt>Dimensions:</dt> <dd>320×240</dd>
  </dl>
 </details>
</section>
```
:::

::: example
The following shows how a
[`details`{#the-details-element:the-details-element-29}](#the-details-element)
element can be used to hide some controls by default:

``` html
<details>
 <summary><label for=fn>Name & Extension:</label></summary>
 <p><input type=text id=fn name=fn value="Pillar Magazine.pdf">
 <p><label><input type=checkbox name=ext checked> Hide extension</label>
</details>
```

One could use this in conjunction with other
[`details`{#the-details-element:the-details-element-30}](#the-details-element)
in a list to allow the user to collapse a set of fields down to a small
set of headings, with the ability to open each one.

![](/images/sample-details-1.png){width="345"
height="611"}![](/images/sample-details-2.png){width="345" height="666"}

In these examples, the summary really just summarizes what the controls
can change, and not the actual values, which is less than ideal.
:::

::: {#example-details-exclusive-accordion .example}
The following example shows the
[`name`{#the-details-element:attr-details-name-13}](#attr-details-name)
attribute of the
[`details`{#the-details-element:the-details-element-31}](#the-details-element)
element being used to create an exclusive accordion, a set of
[`details`{#the-details-element:the-details-element-32}](#the-details-element)
elements where a user action to open one
[`details`{#the-details-element:the-details-element-33}](#the-details-element)
element causes any open
[`details`{#the-details-element:the-details-element-34}](#the-details-element)
to close.

``` html
<section class="characteristics">
 <details name="frame-characteristics">
  <summary>Material</summary>
  The picture frame is made of solid oak wood.
 </details>
 <details name="frame-characteristics">
  <summary>Size</summary>
  The picture frame fits a photo 40cm tall and 30cm wide.
  The frame is 45cm tall, 35cm wide, and 2cm thick.
 </details>
 <details name="frame-characteristics">
  <summary>Color</summary>
  The picture frame is available in its natural wood
  color, or with black stain.
 </details>
</section>
```
:::

::: {#example-details-exclusive-accordion-setting-open .example}
The following example shows what happens when the
[`open`{#the-details-element:attr-details-open-18}](#attr-details-open)
attribute is set on a
[`details`{#the-details-element:the-details-element-35}](#the-details-element)
element that is part of a set of elements using the
[`name`{#the-details-element:attr-details-name-14}](#attr-details-name)
attribute to create an exclusive accordion.

Given the initial markup:

``` html
<section class="characteristics">
 <details name="frame-characteristics" id="d1" open>...</details>
 <details name="frame-characteristics" id="d2">...</details>
 <details name="frame-characteristics" id="d3">...</details>
</section>
```

and the script:

``` js
document.getElementById("d2").setAttribute("open", "");
```

then the resulting tree after the script executes will be equivalent to
the markup:

``` html
<section class="characteristics">
 <details name="frame-characteristics" id="d1">...</details>
 <details name="frame-characteristics" id="d2" open>...</details>
 <details name="frame-characteristics" id="d3">...</details>
</section>
```

because setting the
[`open`{#the-details-element:attr-details-open-19}](#attr-details-open)
attribute on `d2` removes it from `d1`.

The same happens when the user activates the
[`summary`{#the-details-element:the-summary-element-5}](#the-summary-element)
element inside of `d2`.
:::

::: example
Because the
[`open`{#the-details-element:attr-details-open-20}](#attr-details-open)
attribute is added and removed automatically as the user interacts with
the control, it can be used in CSS to style the element differently
based on its state. Here, a style sheet is used to animate the color of
the summary when the element is opened or closed:

``` html
<style>
 details > summary { transition: color 1s; color: black; }
 details[open] > summary { color: red; }
</style>
<details>
 <summary>Automated Status: Operational</summary>
 <p>Velocity: 12m/s</p>
 <p>Direction: North</p>
</details>
```
:::

#### [4.11.2]{.secno} The [`summary`]{.dfn dfn-type="element"} element[](#the-summary-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/summary](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary "The <summary> HTML element specifies a summary, caption, or legend for a <details> element's disclosure box. Clicking the <summary> element toggles the state of the parent <details> element open and closed.")

Support in all current engines.

::: support
[Firefox49+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome12+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-summary-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-summary-element:concept-element-contexts}:
:   As the [first
    child](https://dom.spec.whatwg.org/#concept-tree-first-child){#the-summary-element:first-child
    x-internal="first-child"} of a
    [`details`{#the-summary-element:the-details-element}](#the-details-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-summary-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-summary-element:phrasing-content-2},
    optionally intermixed with [heading
    content](dom.html#heading-content-2){#the-summary-element:heading-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-summary-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-summary-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-summary-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-summary-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-summary).
:   [For implementers](https://w3c.github.io/html-aam/#el-summary).

[DOM interface](dom.html#concept-element-dom){#the-summary-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-summary-element:htmlelement}](dom.html#htmlelement).

The
[`summary`{#the-summary-element:the-summary-element}](#the-summary-element)
element
[represents](dom.html#represents){#the-summary-element:represents} a
summary, caption, or legend for the rest of the contents of the
[`summary`{#the-summary-element:the-summary-element-2}](#the-summary-element)
element\'s parent
[`details`{#the-summary-element:the-details-element-2}](#the-details-element)
element, if any.

A
[`summary`{#the-summary-element:the-summary-element-3}](#the-summary-element)
element is a [summary for its parent
details]{#summary-for-its-parent-details .dfn} if the following
algorithm returns true:

1.  If this
    [`summary`{#the-summary-element:the-summary-element-4}](#the-summary-element)
    element has no parent, then return false.

2.  Let `parent`{.variable} be this
    [`summary`{#the-summary-element:the-summary-element-5}](#the-summary-element)
    element\'s parent.

3.  If `parent`{.variable} is not a
    [`details`{#the-summary-element:the-details-element-3}](#the-details-element)
    element, then return false.

4.  If `parent`{.variable}\'s first
    [`summary`{#the-summary-element:the-summary-element-6}](#the-summary-element)
    element child is not this
    [`summary`{#the-summary-element:the-summary-element-7}](#the-summary-element)
    element, then return false.

5.  Return true.

The [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-summary-element:activation-behaviour
x-internal="activation-behaviour"} of
[`summary`{#the-summary-element:the-summary-element-8}](#the-summary-element)
elements is to run the following steps:

1.  If this
    [`summary`{#the-summary-element:the-summary-element-9}](#the-summary-element)
    element is not the [summary for its parent
    details](#summary-for-its-parent-details){#the-summary-element:summary-for-its-parent-details},
    then return.

2.  Let `parent`{.variable} be this
    [`summary`{#the-summary-element:the-summary-element-10}](#the-summary-element)
    element\'s parent.

3.  If the
    [`open`{#the-summary-element:attr-details-open}](#attr-details-open)
    attribute is present on `parent`{.variable}, then
    [remove](https://dom.spec.whatwg.org/#concept-element-attributes-remove){#the-summary-element:concept-element-attributes-remove
    x-internal="concept-element-attributes-remove"} it. Otherwise,
    [set](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-summary-element:concept-element-attributes-set-value
    x-internal="concept-element-attributes-set-value"}
    `parent`{.variable}\'s
    [`open`{#the-summary-element:attr-details-open-2}](#attr-details-open)
    attribute to the empty string.

    This will then run the [details notification task
    steps](#details-notification-task-steps){#the-summary-element:details-notification-task-steps}.

#### [4.11.3]{.secno} Commands[](#commands){.self-link}

##### [4.11.3.1]{.secno} Facets[](#facets-2){.self-link} {#facets-2}

A [command]{#concept-command .dfn} is the abstraction behind menu items,
buttons, and links. Once a command is defined, other parts of the
interface can refer to the same command, allowing many access points to
a single feature to share facets such as the [Disabled
State](#command-facet-disabledstate){#facets-2:command-facet-disabledstate}.

Commands are defined to have the following [facets]{#concept-facet
.dfn}:

[Label]{#command-facet-label .dfn}
:   The name of the command as seen by the user.

[Access Key]{#command-facet-accesskey .dfn}
:   A key combination selected by the user agent that triggers the
    command. A command might not have an Access Key.

[Hidden State]{#command-facet-hiddenstate .dfn}
:   Whether the command is hidden or not (basically, whether it should
    be shown in menus).

[Disabled State]{#command-facet-disabledstate .dfn}
:   Whether the command is relevant and can be triggered or not.

[Action]{#command-facet-action .dfn}
:   The actual effect that triggering the command will have. This could
    be a scripted event handler, a
    [URL](https://url.spec.whatwg.org/#concept-url){#facets-2:url
    x-internal="url"} to which to
    [navigate](browsing-the-web.html#navigate){#facets-2:navigate}, or a
    form submission.

User agents may expose the
[commands](#concept-command){#facets-2:concept-command} that match the
following criteria:

- The [Hidden
  State](#command-facet-hiddenstate){#facets-2:command-facet-hiddenstate}
  facet is false (visible)

- The element is [in a
  document](https://dom.spec.whatwg.org/#in-a-document){#facets-2:in-a-document
  x-internal="in-a-document"} with a non-null [browsing
  context](document-sequences.html#concept-document-bc){#facets-2:concept-document-bc}.

- Neither the element nor any of its ancestors has a
  [`hidden`{#facets-2:attr-hidden}](interaction.html#attr-hidden)
  attribute specified.

User agents are encouraged to do this especially for commands that have
[Access
Keys](#command-facet-accesskey){#facets-2:command-facet-accesskey}, as a
way to advertise those keys to the user.

For example, such commands could be listed in the user agent\'s menu
bar.

##### [4.11.3.2]{.secno} [Using the `a` element to define a command]{.dfn}[](#using-the-a-element-to-define-a-command){.self-link}

An
[`a`{#using-the-a-element-to-define-a-command:the-a-element}](text-level-semantics.html#the-a-element)
element with an
[`href`{#using-the-a-element-to-define-a-command:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attribute [defines a
command](#concept-command){#using-the-a-element-to-define-a-command:concept-command}.

The
[Label](#command-facet-label){#using-the-a-element-to-define-a-command:command-facet-label}
of the command is the element\'s [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#using-the-a-element-to-define-a-command:descendant-text-content
x-internal="descendant-text-content"}.

The [Access
Key](#command-facet-accesskey){#using-the-a-element-to-define-a-command:command-facet-accesskey}
of the command is the element\'s [assigned access
key](interaction.html#assigned-access-key){#using-the-a-element-to-define-a-command:assigned-access-key},
if any.

The [Hidden
State](#command-facet-hiddenstate){#using-the-a-element-to-define-a-command:command-facet-hiddenstate}
of the command is true (hidden) if the element has a
[`hidden`{#using-the-a-element-to-define-a-command:attr-hidden}](interaction.html#attr-hidden)
attribute, and false otherwise.

The [Disabled
State](#command-facet-disabledstate){#using-the-a-element-to-define-a-command:command-facet-disabledstate}
facet of the command is true if the element or one of its ancestors is
[inert](interaction.html#inert){#using-the-a-element-to-define-a-command:inert},
and false otherwise.

The
[Action](#command-facet-action){#using-the-a-element-to-define-a-command:command-facet-action}
of the command is to [fire a `click`
event](webappapis.html#fire-a-click-event){#using-the-a-element-to-define-a-command:fire-a-click-event}
at the element.

##### [4.11.3.3]{.secno} [Using the `button` element to define a command]{.dfn}[](#using-the-button-element-to-define-a-command){.self-link}

A
[`button`{#using-the-button-element-to-define-a-command:the-button-element}](form-elements.html#the-button-element)
element always [defines a
command](#concept-command){#using-the-button-element-to-define-a-command:concept-command}.

The
[Label](#command-facet-label){#using-the-button-element-to-define-a-command:command-facet-label},
[Access
Key](#command-facet-accesskey){#using-the-button-element-to-define-a-command:command-facet-accesskey},
[Hidden
State](#command-facet-hiddenstate){#using-the-button-element-to-define-a-command:command-facet-hiddenstate},
and
[Action](#command-facet-action){#using-the-button-element-to-define-a-command:command-facet-action}
facets of the command are determined [as for `a`
elements](#using-the-a-element-to-define-a-command){#using-the-button-element-to-define-a-command:using-the-a-element-to-define-a-command}
(see the previous section).

The [Disabled
State](#command-facet-disabledstate){#using-the-button-element-to-define-a-command:command-facet-disabledstate}
of the command is true if the element or one of its ancestors is
[inert](interaction.html#inert){#using-the-button-element-to-define-a-command:inert},
or if the element\'s
[disabled](form-control-infrastructure.html#concept-fe-disabled){#using-the-button-element-to-define-a-command:concept-fe-disabled}
state is set, and false otherwise.

##### [4.11.3.4]{.secno} [Using the `input` element to define a command]{.dfn}[](#using-the-input-element-to-define-a-command){.self-link}

An
[`input`{#using-the-input-element-to-define-a-command:the-input-element}](input.html#the-input-element)
element whose
[`type`{#using-the-input-element-to-define-a-command:attr-input-type}](input.html#attr-input-type)
attribute is in one of the [Submit
Button](input.html#submit-button-state-(type=submit)){#using-the-input-element-to-define-a-command:submit-button-state-(type=submit)},
[Reset
Button](input.html#reset-button-state-(type=reset)){#using-the-input-element-to-define-a-command:reset-button-state-(type=reset)},
[Image
Button](input.html#image-button-state-(type=image)){#using-the-input-element-to-define-a-command:image-button-state-(type=image)},
[Button](input.html#button-state-(type=button)){#using-the-input-element-to-define-a-command:button-state-(type=button)},
[Radio
Button](input.html#radio-button-state-(type=radio)){#using-the-input-element-to-define-a-command:radio-button-state-(type=radio)},
or
[Checkbox](input.html#checkbox-state-(type=checkbox)){#using-the-input-element-to-define-a-command:checkbox-state-(type=checkbox)}
states [defines a
command](#concept-command){#using-the-input-element-to-define-a-command:concept-command}.

The
[Label](#command-facet-label){#using-the-input-element-to-define-a-command:command-facet-label}
of the command is determined as follows:

- If the
  [`type`{#using-the-input-element-to-define-a-command:attr-input-type-2}](input.html#attr-input-type)
  attribute is in one of the [Submit
  Button](input.html#submit-button-state-(type=submit)){#using-the-input-element-to-define-a-command:submit-button-state-(type=submit)-2},
  [Reset
  Button](input.html#reset-button-state-(type=reset)){#using-the-input-element-to-define-a-command:reset-button-state-(type=reset)-2},
  [Image
  Button](input.html#image-button-state-(type=image)){#using-the-input-element-to-define-a-command:image-button-state-(type=image)-2},
  or
  [Button](input.html#button-state-(type=button)){#using-the-input-element-to-define-a-command:button-state-(type=button)-2}
  states, then the
  [Label](#command-facet-label){#using-the-input-element-to-define-a-command:command-facet-label-2}
  is the string given by the
  [`value`{#using-the-input-element-to-define-a-command:attr-input-value}](input.html#attr-input-value)
  attribute, if any, and a UA-dependent, locale-dependent value that the
  UA uses to label the button itself if the attribute is absent.

- Otherwise, if the element is a [labeled
  control](forms.html#labeled-control){#using-the-input-element-to-define-a-command:labeled-control},
  then the
  [Label](#command-facet-label){#using-the-input-element-to-define-a-command:command-facet-label-3}
  is the [descendant text
  content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#using-the-input-element-to-define-a-command:descendant-text-content
  x-internal="descendant-text-content"} of the first
  [`label`{#using-the-input-element-to-define-a-command:the-label-element}](forms.html#the-label-element)
  element in [tree
  order](https://dom.spec.whatwg.org/#concept-tree-order){#using-the-input-element-to-define-a-command:tree-order
  x-internal="tree-order"} whose [labeled
  control](forms.html#labeled-control){#using-the-input-element-to-define-a-command:labeled-control-2}
  is the element in question. (In JavaScript terms, this is given by
  `element`{.variable}`.labels[0].textContent`.)

- Otherwise, if the
  [`value`{#using-the-input-element-to-define-a-command:attr-input-value-2}](input.html#attr-input-value)
  attribute is present, then the
  [Label](#command-facet-label){#using-the-input-element-to-define-a-command:command-facet-label-4}
  is the value of that attribute.

- Otherwise, the
  [Label](#command-facet-label){#using-the-input-element-to-define-a-command:command-facet-label-5}
  is the empty string.

::: note
Even though the
[`value`{#using-the-input-element-to-define-a-command:attr-input-value-3}](input.html#attr-input-value)
attribute on
[`input`{#using-the-input-element-to-define-a-command:the-input-element-2}](input.html#the-input-element)
elements in the [Image
Button](input.html#image-button-state-(type=image)){#using-the-input-element-to-define-a-command:image-button-state-(type=image)-3}
state is non-conformant, the attribute can still contribute to the
[Label](#command-facet-label){#using-the-input-element-to-define-a-command:command-facet-label-6}
determination, if it is present and the Image Button\'s
[`alt`{#using-the-input-element-to-define-a-command:attr-input-alt}](input.html#attr-input-alt)
attribute is missing.
:::

The [Access
Key](#command-facet-accesskey){#using-the-input-element-to-define-a-command:command-facet-accesskey}
of the command is the element\'s [assigned access
key](interaction.html#assigned-access-key){#using-the-input-element-to-define-a-command:assigned-access-key},
if any.

The [Hidden
State](#command-facet-hiddenstate){#using-the-input-element-to-define-a-command:command-facet-hiddenstate}
of the command is true (hidden) if the element has a
[`hidden`{#using-the-input-element-to-define-a-command:attr-hidden}](interaction.html#attr-hidden)
attribute, and false otherwise.

The [Disabled
State](#command-facet-disabledstate){#using-the-input-element-to-define-a-command:command-facet-disabledstate}
of the command is true if the element or one of its ancestors is
[inert](interaction.html#inert){#using-the-input-element-to-define-a-command:inert},
or if the element\'s
[disabled](form-control-infrastructure.html#concept-fe-disabled){#using-the-input-element-to-define-a-command:concept-fe-disabled}
state is set, and false otherwise.

The
[Action](#command-facet-action){#using-the-input-element-to-define-a-command:command-facet-action}
of the command is to [fire a `click`
event](webappapis.html#fire-a-click-event){#using-the-input-element-to-define-a-command:fire-a-click-event}
at the element.

##### [4.11.3.5]{.secno} [Using the `option` element to define a command]{.dfn}[](#using-the-option-element-to-define-a-command){.self-link}

An
[`option`{#using-the-option-element-to-define-a-command:the-option-element}](form-elements.html#the-option-element)
element with an ancestor
[`select`{#using-the-option-element-to-define-a-command:the-select-element}](form-elements.html#the-select-element)
element and either no
[`value`{#using-the-option-element-to-define-a-command:attr-option-value}](form-elements.html#attr-option-value)
attribute or a
[`value`{#using-the-option-element-to-define-a-command:attr-option-value-2}](form-elements.html#attr-option-value)
attribute that is not the empty string [defines a
command](#concept-command){#using-the-option-element-to-define-a-command:concept-command}.

The
[Label](#command-facet-label){#using-the-option-element-to-define-a-command:command-facet-label}
of the command is the value of the
[`option`{#using-the-option-element-to-define-a-command:the-option-element-2}](form-elements.html#the-option-element)
element\'s
[`label`{#using-the-option-element-to-define-a-command:attr-option-label}](form-elements.html#attr-option-label)
attribute, if there is one, or else the
[`option`{#using-the-option-element-to-define-a-command:the-option-element-3}](form-elements.html#the-option-element)
element\'s [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#using-the-option-element-to-define-a-command:descendant-text-content
x-internal="descendant-text-content"}, with [ASCII whitespace stripped
and
collapsed](https://infra.spec.whatwg.org/#strip-and-collapse-ascii-whitespace){#using-the-option-element-to-define-a-command:strip-and-collapse-ascii-whitespace
x-internal="strip-and-collapse-ascii-whitespace"}.

The [Access
Key](#command-facet-accesskey){#using-the-option-element-to-define-a-command:command-facet-accesskey}
of the command is the element\'s [assigned access
key](interaction.html#assigned-access-key){#using-the-option-element-to-define-a-command:assigned-access-key},
if any.

The [Hidden
State](#command-facet-hiddenstate){#using-the-option-element-to-define-a-command:command-facet-hiddenstate}
of the command is true (hidden) if the element has a
[`hidden`{#using-the-option-element-to-define-a-command:attr-hidden}](interaction.html#attr-hidden)
attribute, and false otherwise.

The [Disabled
State](#command-facet-disabledstate){#using-the-option-element-to-define-a-command:command-facet-disabledstate}
of the command is true if the element is
[disabled](form-elements.html#concept-option-disabled){#using-the-option-element-to-define-a-command:concept-option-disabled},
or if its nearest ancestor
[`select`{#using-the-option-element-to-define-a-command:the-select-element-2}](form-elements.html#the-select-element)
element is
[disabled](form-control-infrastructure.html#concept-fe-disabled){#using-the-option-element-to-define-a-command:concept-fe-disabled},
or if it or one of its ancestors is
[inert](interaction.html#inert){#using-the-option-element-to-define-a-command:inert},
and false otherwise.

If the
[`option`{#using-the-option-element-to-define-a-command:the-option-element-4}](form-elements.html#the-option-element)\'s
nearest ancestor
[`select`{#using-the-option-element-to-define-a-command:the-select-element-3}](form-elements.html#the-select-element)
element has a
[`multiple`{#using-the-option-element-to-define-a-command:attr-select-multiple}](form-elements.html#attr-select-multiple)
attribute, the
[Action](#command-facet-action){#using-the-option-element-to-define-a-command:command-facet-action}
of the command is to
[toggle](form-elements.html#concept-select-toggle){#using-the-option-element-to-define-a-command:concept-select-toggle}
the
[`option`{#using-the-option-element-to-define-a-command:the-option-element-5}](form-elements.html#the-option-element)
element. Otherwise, the
[Action](#command-facet-action){#using-the-option-element-to-define-a-command:command-facet-action-2}
is to
[pick](form-elements.html#concept-select-pick){#using-the-option-element-to-define-a-command:concept-select-pick}
the
[`option`{#using-the-option-element-to-define-a-command:the-option-element-6}](form-elements.html#the-option-element)
element.

##### [4.11.3.6]{.secno} [Using the `accesskey` attribute on a `legend` element to define a command]{.dfn}[](#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command){.self-link}

A
[`legend`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-legend-element}](form-elements.html#the-legend-element)
element [defines a
command](#concept-command){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:concept-command}
if all of the following are true:

- It has an [assigned access
  key](interaction.html#assigned-access-key){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:assigned-access-key}.

- It is a child of a
  [`fieldset`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-fieldset-element}](form-elements.html#the-fieldset-element)
  element.

- Its parent has a descendant that [defines a
  command](#concept-command){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:concept-command-2}
  that is neither a
  [`label`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-label-element}](forms.html#the-label-element)
  element nor a
  [`legend`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-legend-element-2}](form-elements.html#the-legend-element)
  element. This element, if it exists, is [the `legend` element\'s
  `accesskey` delegatee]{#the-legend-element's-accesskey-delegatee
  .dfn}.

The
[Label](#command-facet-label){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:command-facet-label}
of the command is the element\'s [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:descendant-text-content
x-internal="descendant-text-content"}.

The [Access
Key](#command-facet-accesskey){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:command-facet-accesskey}
of the command is the element\'s [assigned access
key](interaction.html#assigned-access-key){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:assigned-access-key-2}.

The [Hidden
State](#command-facet-hiddenstate){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:command-facet-hiddenstate},
[Disabled
State](#command-facet-disabledstate){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:command-facet-disabledstate},
and
[Action](#command-facet-action){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:command-facet-action}
facets of the command are the same as the respective facets of [the
`legend` element\'s `accesskey`
delegatee](#the-legend-element's-accesskey-delegatee){#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-legend-element's-accesskey-delegatee}.

::: example
In this example, the
[`legend`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-legend-element-3}](form-elements.html#the-legend-element)
element specifies an
[`accesskey`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-accesskey-attribute}](interaction.html#the-accesskey-attribute),
which, when activated, will delegate to the
[`input`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-input-element}](input.html#the-input-element)
element inside the
[`legend`{#using-the-accesskey-attribute-on-a-legend-element-to-define-a-command:the-legend-element-4}](form-elements.html#the-legend-element)
element.

``` html
<fieldset>
 <legend accesskey=p>
  <label>I want <input name=pizza type=number step=1 value=1 min=0>
   pizza(s) with these toppings</label>
 </legend>
 <label><input name=pizza-cheese type=checkbox checked> Cheese</label>
 <label><input name=pizza-ham type=checkbox checked> Ham</label>
 <label><input name=pizza-pineapple type=checkbox> Pineapple</label>
</fieldset>
```
:::

##### [4.11.3.7]{.secno} [Using the `accesskey` attribute to define a command on other elements]{.dfn}[](#using-the-accesskey-attribute-to-define-a-command-on-other-elements){.self-link}

An element that has an [assigned access
key](interaction.html#assigned-access-key){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:assigned-access-key}
[defines a
command](#concept-command){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:concept-command}.

If one of the earlier sections that define elements that [define
commands](#concept-command){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:concept-command-2}
define that this element [defines a
command](#concept-command){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:concept-command-3},
then that section applies to this element, and this section does not.
Otherwise, this section applies to that element.

The
[Label](#command-facet-label){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:command-facet-label}
of the command depends on the element. If the element is a [labeled
control](forms.html#labeled-control){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:labeled-control},
the [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:descendant-text-content
x-internal="descendant-text-content"} of the first
[`label`{#using-the-accesskey-attribute-to-define-a-command-on-other-elements:the-label-element}](forms.html#the-label-element)
element in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:tree-order
x-internal="tree-order"} whose [labeled
control](forms.html#labeled-control){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:labeled-control-2}
is the element in question is the
[Label](#command-facet-label){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:command-facet-label-2}
(in JavaScript terms, this is given by
`element`{.variable}`.labels[0].textContent`). Otherwise, the
[Label](#command-facet-label){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:command-facet-label-3}
is the element\'s [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:descendant-text-content-2
x-internal="descendant-text-content"}.

The [Access
Key](#command-facet-accesskey){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:command-facet-accesskey}
of the command is the element\'s [assigned access
key](interaction.html#assigned-access-key){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:assigned-access-key-2}.

The [Hidden
State](#command-facet-hiddenstate){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:command-facet-hiddenstate}
of the command is true (hidden) if the element has a
[`hidden`{#using-the-accesskey-attribute-to-define-a-command-on-other-elements:attr-hidden}](interaction.html#attr-hidden)
attribute, and false otherwise.

The [Disabled
State](#command-facet-disabledstate){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:command-facet-disabledstate}
of the command is true if the element or one of its ancestors is
[inert](interaction.html#inert){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:inert},
and false otherwise.

The
[Action](#command-facet-action){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:command-facet-action}
of the command is to run the following steps:

1.  Run the [focusing
    steps](interaction.html#focusing-steps){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:focusing-steps}
    for the element.
2.  [Fire a `click`
    event](webappapis.html#fire-a-click-event){#using-the-accesskey-attribute-to-define-a-command-on-other-elements:fire-a-click-event}
    at the element.

#### [4.11.4]{.secno} The [`dialog`]{.dfn dfn-type="element"} element[](#the-dialog-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/dialog](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog "The <dialog> HTML element represents a dialog box or other interactive component, such as a dismissible alert, inspector, or subwindow.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

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
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDialogElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement "The HTMLDialogElement interface provides methods to manipulate <dialog> elements. It inherits properties and methods from the HTMLElement interface.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

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
:::::

[Categories](dom.html#concept-element-categories){#the-dialog-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-dialog-element:flow-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-dialog-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-dialog-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-dialog-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-dialog-element:flow-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-dialog-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-dialog-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-dialog-element:global-attributes}
:   [`closedby`{#the-dialog-element:attr-dialog-closedby}](#attr-dialog-closedby)
    --- Which user actions will close the dialog
:   [`open`{#the-dialog-element:attr-dialog-open}](#attr-dialog-open)
    --- Whether the dialog box is showing

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-dialog-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-dialog).
:   [For implementers](https://w3c.github.io/html-aam/#el-dialog).

[DOM interface](dom.html#concept-element-dom){#the-dialog-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLDialogElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute boolean open;
      attribute DOMString returnValue;
      [CEReactions] attribute DOMString closedBy;
      [CEReactions] undefined show();
      [CEReactions] undefined showModal();
      [CEReactions] undefined close(optional DOMString returnValue);
      [CEReactions] undefined requestClose(optional DOMString returnValue);
    };
    ```

The
[`dialog`{#the-dialog-element:the-dialog-element}](#the-dialog-element)
element represents a transitory part of an application, in the form of a
small window (\"dialog box\"), which the user interacts with to perform
a task or gather information. Once the user is done, the dialog can be
automatically closed by the application, or manually closed by the user.

Especially for modal dialogs, which are a familiar pattern across all
types of applications, authors should work to ensure that dialogs in
their web applications behave in a way that is familiar to users of
non-web applications.

As with all HTML elements, it is not conforming to use the
[`dialog`{#the-dialog-element:the-dialog-element-2}](#the-dialog-element)
element when attempting to represent another type of control. For
example, context menus, tooltips, and popup listboxes are not dialog
boxes, so abusing the
[`dialog`{#the-dialog-element:the-dialog-element-3}](#the-dialog-element)
element to implement these patterns is incorrect.

An important part of user-facing dialog behavior is the placement of
initial focus. The [dialog focusing
steps](#dialog-focusing-steps){#the-dialog-element:dialog-focusing-steps}
attempt to pick a good candidate for initial focus when a dialog is
shown, but might not be a substitute for authors carefully thinking
through the correct choice to match user expectations for a specific
dialog. As such, authors should use the
[`autofocus`{#the-dialog-element:attr-fe-autofocus}](interaction.html#attr-fe-autofocus)
attribute on the descendant element of the dialog that the user is
expected to immediately interact with after the dialog opens. If there
is no such element, then authors should use the
[`autofocus`{#the-dialog-element:attr-fe-autofocus-2}](interaction.html#attr-fe-autofocus)
attribute on the
[`dialog`{#the-dialog-element:the-dialog-element-4}](#the-dialog-element)
element itself.

::: example
In the following example, a dialog is used for editing the details of a
product in an inventory management web application.

``` html
<dialog>
  <label>Product Number <input type="text" readonly></label>
  <label>Product Name <input type="text" autofocus></label>
</dialog>
```

If the
[`autofocus`{#the-dialog-element:attr-fe-autofocus-3}](interaction.html#attr-fe-autofocus)
attribute was not present, the Product Number field would have been
focused by the dialog focusing steps. Although that is reasonable
behavior, the author determined that the more relevant field to focus
was the Product Name field, as the Product Number field is readonly and
expects no user input. So, the author used autofocus to override the
default.

Even if the author wants to focus the Product Number field by default,
they are best off explicitly specifying that by using autofocus on that
[`input`{#the-dialog-element:the-input-element}](input.html#the-input-element)
element. This makes the intent obvious to future readers of the code,
and ensures the code stays robust in the face of future updates. (For
example, if another developer added a close button, and positioned it in
the node tree before the Product Number field).
:::

Another important aspect of user behavior is whether dialogs are
scrollable or not. In some cases, overflow (and thus scrollability)
cannot be avoided, e.g., when it is caused by the user\'s high text zoom
settings. But in general, scrollable dialogs are not expected by users.
Adding large text nodes directly to dialog elements is particularly bad
as this is likely to cause the dialog element itself to overflow.
Authors are best off avoiding them.

::: example
The following terms of service dialog respects the above suggestions.

``` html
<dialog style="height: 80vh;">
  <div style="overflow: auto; height: 60vh;" autofocus>
    <p>By placing an order via this Web site on the first day of the fourth month of the year
    2010 Anno Domini, you agree to grant Us a non-transferable option to claim, for now and for
    ever more, your immortal soul.</p>
    <p>Should We wish to exercise this option, you agree to surrender your immortal soul,
    and any claim you may have on it, within 5 (five) working days of receiving written
    notification from  this site or one of its duly authorized minions.</p>
    <!-- ... etc., with many more <p> elements ... -->
  </div>
  <form method="dialog">
    <button type="submit" value="agree">Agree</button>
    <button type="submit" value="disagree">Disagree</button>
  </form>
</dialog>
```

Note how the [dialog focusing
steps](#dialog-focusing-steps){#the-dialog-element:dialog-focusing-steps-2}
would have picked the scrollable
[`div`{#the-dialog-element:the-div-element}](grouping-content.html#the-div-element)
element by default, but similarly to the previous example, we have
placed
[`autofocus`{#the-dialog-element:attr-fe-autofocus-4}](interaction.html#attr-fe-autofocus)
on the
[`div`{#the-dialog-element:the-div-element-2}](grouping-content.html#the-div-element)
so as to be more explicit and robust against future changes.

In contrast, if the
[`p`{#the-dialog-element:the-p-element}](grouping-content.html#the-p-element)
elements expressing the terms of service did not have such a wrapper
[`div`{#the-dialog-element:the-div-element-3}](grouping-content.html#the-div-element)
element, then the
[`dialog`{#the-dialog-element:the-dialog-element-5}](#the-dialog-element)
itself would become scrollable, violating the above advice. Furthermore,
in the absence of any
[`autofocus`{#the-dialog-element:attr-fe-autofocus-5}](interaction.html#attr-fe-autofocus)
attribute, such a markup pattern would have violated the above advice
and tripped up the [dialog focusing
steps](#dialog-focusing-steps){#the-dialog-element:dialog-focusing-steps-3}\'s
default behavior, and caused focus to jump to the Agree
[`button`{#the-dialog-element:the-button-element}](form-elements.html#the-button-element),
which is a bad user experience.
:::

::: example
This dialog box has some small print. The
[`strong`{#the-dialog-element:the-strong-element}](text-level-semantics.html#the-strong-element)
element is used to draw the user\'s attention to the more important
part.

``` html
<dialog>
 <h1>Add to Wallet</h1>
 <p><strong><label for=amt>How many gold coins do you want to add to your wallet?</label></strong></p>
 <p><input id=amt name=amt type=number min=0 step=0.01 value=100></p>
 <p><small>You add coins at your own risk.</small></p>
 <p><label><input name=round type=checkbox> Only add perfectly round coins</label></p>
 <p><input type=button onclick="submit()" value="Add Coins"></p>
</dialog>
```
:::

------------------------------------------------------------------------

The [`open`]{#attr-dialog-open .dfn dfn-for="dialog"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-dialog-element:boolean-attribute}.
When specified, it indicates that the
[`dialog`{#the-dialog-element:the-dialog-element-6}](#the-dialog-element)
element is active and that the user can interact with it.

The [`closedby`]{#attr-dialog-closedby .dfn dfn-for="dialog"
dfn-type="element-attr"} content attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-dialog-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`any`]{#attr-dialog-closedby-any .dfn dfn-for="dialog/closedby"
dfn-type="attr-value"}

[Any]{#attr-dialog-closedby-any-state .dfn}

[Close
requests](interaction.html#close-request){#the-dialog-element:close-request}
or clicks outside close the dialog.

[`closerequest`]{#attr-dialog-closedby-closerequest .dfn
dfn-for="dialog/closedby" dfn-type="attr-value"}

[Close Request]{#attr-dialog-closedby-closerequest-state .dfn}

[Close
requests](interaction.html#close-request){#the-dialog-element:close-request-2}
close the dialog.

[`none`]{#attr-dialog-closedby-none .dfn dfn-for="dialog/closedby"
dfn-type="attr-value"}

[None]{#attr-dialog-closedby-none-state .dfn}

No user actions automatically close the dialog.

The
[`closedby`{#the-dialog-element:attr-dialog-closedby-2}](#attr-dialog-closedby)
attribute\'s *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* and *[missing
value default](common-microsyntaxes.html#missing-value-default)* are
both the [Auto]{#attr-dialog-closedby-auto-state .dfn} state.

The
[Auto](#attr-dialog-closedby-auto-state){#the-dialog-element:attr-dialog-closedby-auto-state}
state behaves as [Close
Request](#attr-dialog-closedby-closerequest-state){#the-dialog-element:attr-dialog-closedby-closerequest-state}
state when the
[`dialog`{#the-dialog-element:the-dialog-element-7}](#the-dialog-element)
was shown using its
[`showModal()`{#the-dialog-element:dom-dialog-showmodal-2}](#dom-dialog-showmodal)
method; otherwise the
[None](#attr-dialog-closedby-none-state){#the-dialog-element:attr-dialog-closedby-none-state}
state.

A
[`dialog`{#the-dialog-element:the-dialog-element-8}](#the-dialog-element)
element without an
[`open`{#the-dialog-element:attr-dialog-open-2}](#attr-dialog-open)
attribute specified should not be shown to the user. This requirement
may be implemented indirectly through the style layer. For example, user
agents that [support the suggested default
rendering](infrastructure.html#renderingUA) implement this requirement
using the CSS rules described in the [Rendering
section](rendering.html#rendering).

::: {#note-dialog-remove-open-attribute .note}
Removing the
[`open`{#the-dialog-element:attr-dialog-open-3}](#attr-dialog-open)
attribute will usually hide the dialog. However, doing so has a number
of strange additional consequences:

- The
  [`close`{#the-dialog-element:event-close}](indices.html#event-close)
  event will not be fired.

- The
  [`close()`{#the-dialog-element:dom-dialog-close-2}](#dom-dialog-close)
  method, and any [close
  requests](interaction.html#close-request){#the-dialog-element:close-request-3},
  will no longer be able to close the dialog.

- If the dialog was shown using its
  [`showModal()`{#the-dialog-element:dom-dialog-showmodal-3}](#dom-dialog-showmodal)
  method, the
  [`Document`{#the-dialog-element:document}](dom.html#document) will
  still be
  [blocked](interaction.html#blocked-by-a-modal-dialog){#the-dialog-element:blocked-by-a-modal-dialog}.

For these reasons, it is generally better to never remove the
[`open`{#the-dialog-element:attr-dialog-open-4}](#attr-dialog-open)
attribute manually. Instead, use the
[`requestClose()`{#the-dialog-element:dom-dialog-requestclose-2}](#dom-dialog-requestclose)
or
[`close()`{#the-dialog-element:dom-dialog-close-3}](#dom-dialog-close)
methods to close the dialog, or the
[`hidden`{#the-dialog-element:attr-hidden}](interaction.html#attr-hidden)
attribute to hide it.
:::

The
[`tabindex`{#the-dialog-element:attr-tabindex}](interaction.html#attr-tabindex)
attribute must not be specified on
[`dialog`{#the-dialog-element:the-dialog-element-9}](#the-dialog-element)
elements.

------------------------------------------------------------------------

`dialog`{.variable}`.`[`show`](#dom-dialog-show){#dom-dialog-show-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDialogElement/show](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/show "The show() method of the HTMLDialogElement interface displays the dialog modelessly, i.e. still allowing interaction with content outside of the dialog.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

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
:::::

Displays the
[`dialog`{#the-dialog-element:the-dialog-element-10}](#the-dialog-element)
element.

`dialog`{.variable}`.`[`showModal`](#dom-dialog-showmodal){#dom-dialog-showmodal-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDialogElement/showModal](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/showModal "The showModal() method of the HTMLDialogElement interface displays the dialog as a modal, over the top of any other dialogs that might be present. It displays in the top layer, along with a ::backdrop pseudo-element. Interaction outside the dialog is blocked and the content outside it is rendered inert.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

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
:::::

Displays the
[`dialog`{#the-dialog-element:the-dialog-element-11}](#the-dialog-element)
element and makes it the top-most modal dialog.

This method honors the
[`autofocus`{#the-dialog-element:attr-fe-autofocus-6}](interaction.html#attr-fe-autofocus)
attribute.

`dialog`{.variable}`.`[`close`](#dom-dialog-close){#dom-dialog-close-dev}`([ ``result`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDialogElement/close](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/close "The close() method of the HTMLDialogElement interface closes the <dialog>. An optional string may be passed as an argument, updating the returnValue of the dialog.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

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
:::::

Closes the
[`dialog`{#the-dialog-element:the-dialog-element-12}](#the-dialog-element)
element.

The argument, if provided, provides a return value.

`dialog`{.variable}`.`[`requestClose`](#dom-dialog-requestclose){#dom-dialog-requestclose-dev}`([ ``result`{.variable}` ])`

Acts as if a [close
request](interaction.html#close-request){#the-dialog-element:close-request-4}
was sent targeting `dialog`{.variable}, by first firing a
[`cancel`{#the-dialog-element:event-cancel}](indices.html#event-cancel)
event, and if that event is not canceled with
[`preventDefault()`{#the-dialog-element:dom-event-preventdefault}](https://dom.spec.whatwg.org/#dom-event-preventdefault){x-internal="dom-event-preventdefault"},
proceeding to close the `dialog`{.variable} in the same way as the
[`close()`{#the-dialog-element:dom-dialog-close-4}](#dom-dialog-close)
method (including firing a
[`close`{#the-dialog-element:event-close-2}](indices.html#event-close)
event).

This is a helper utility that can be used to consolidate cancelation and
closing logic into the
[`cancel`{#the-dialog-element:event-cancel-2}](indices.html#event-cancel)
and
[`close`{#the-dialog-element:event-close-3}](indices.html#event-close)
event handlers, by having all non-[close
request](interaction.html#close-request){#the-dialog-element:close-request-5}
closing affordances call this method.

Note that this method ignores the
[`closedby`{#the-dialog-element:attr-dialog-closedby-3}](#attr-dialog-closedby)
attribute: that is, even if
[`closedby`{#the-dialog-element:attr-dialog-closedby-4}](#attr-dialog-closedby)
is set to
\"[`none`{#the-dialog-element:attr-dialog-closedby-none}](#attr-dialog-closedby-none)\",
the same behavior will apply.

The argument, if provided, provides a return value.

`dialog`{.variable}`.`[`returnValue`](#dom-dialog-returnvalue){#dom-dialog-returnvalue-dev}` [ = ``result`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDialogElement/returnValue](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/returnValue "The returnValue property of the HTMLDialogElement interface gets or sets the return value for the <dialog>, usually to indicate which button the user pressed to close it.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

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
:::::

Returns the
[`dialog`{#the-dialog-element:the-dialog-element-13}](#the-dialog-element)\'s
return value.

Can be set, to update the return value.

The [`show()`]{#dom-dialog-show .dfn dfn-for="HTMLDialogElement"
dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this
    x-internal="this"} has an
    [`open`{#the-dialog-element:attr-dialog-open-5}](#attr-dialog-open)
    attribute and [is modal](#is-modal){#the-dialog-element:is-modal} of
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-2
    x-internal="this"} is false, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-3
    x-internal="this"} has an
    [`open`{#the-dialog-element:attr-dialog-open-6}](#attr-dialog-open)
    attribute, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-dialog-element:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-dialog-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If the result of [firing an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-dialog-element:concept-event-fire
    x-internal="concept-event-fire"} named
    [`beforetoggle`{#the-dialog-element:event-beforetoggle}](indices.html#event-beforetoggle),
    using
    [`ToggleEvent`{#the-dialog-element:toggleevent}](interaction.html#toggleevent),
    with the
    [`cancelable`{#the-dialog-element:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    attribute initialized to true, the
    [`oldState`{#the-dialog-element:dom-toggleevent-oldstate}](interaction.html#dom-toggleevent-oldstate)
    attribute initialized to \"`closed`\", and the
    [`newState`{#the-dialog-element:dom-toggleevent-newstate}](interaction.html#dom-toggleevent-newstate)
    attribute initialized to \"`open`\" at
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-4
    x-internal="this"} is false, then return.

4.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-5
    x-internal="this"} has an
    [`open`{#the-dialog-element:attr-dialog-open-7}](#attr-dialog-open)
    attribute, then return.

5.  [Queue a dialog toggle event
    task](#queue-a-dialog-toggle-event-task){#the-dialog-element:queue-a-dialog-toggle-event-task}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-6
    x-internal="this"}, \"`closed`\", \"`open`\", and null.

6.  Add an
    [`open`{#the-dialog-element:attr-dialog-open-8}](#attr-dialog-open)
    attribute to
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-7
    x-internal="this"}, whose value is the empty string.

7.  [Assert](https://infra.spec.whatwg.org/#assert){#the-dialog-element:assert
    x-internal="assert"}:
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-8
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document
    x-internal="node-document"}\'s [open dialogs
    list](dom.html#open-dialogs-list){#the-dialog-element:open-dialogs-list}
    does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#the-dialog-element:list-contains
    x-internal="list-contains"}
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-9
    x-internal="this"}.

8.  Add
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-10
    x-internal="this"} to
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-11
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-2
    x-internal="node-document"}\'s [open dialogs
    list](dom.html#open-dialogs-list){#the-dialog-element:open-dialogs-list-2}.

9.  [Set the dialog close
    watcher](#set-the-dialog-close-watcher){#the-dialog-element:set-the-dialog-close-watcher}
    with
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-12
    x-internal="this"}.

10. Set
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-13
    x-internal="this"}\'s [previously focused
    element](#previously-focused-element){#the-dialog-element:previously-focused-element}
    to the
    [focused](interaction.html#focused){#the-dialog-element:focused}
    element.

11. Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-14
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-3
    x-internal="node-document"}.

12. Let `hideUntil`{.variable} be the result of running [topmost popover
    ancestor](popover.html#topmost-popover-ancestor){#the-dialog-element:topmost-popover-ancestor}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-15
    x-internal="this"}, `document`{.variable}\'s [showing hint popover
    list](popover.html#showing-hint-popover-list){#the-dialog-element:showing-hint-popover-list},
    null, and false.

13. If `hideUntil`{.variable} is null, then set `hideUntil`{.variable}
    to the result of running [topmost popover
    ancestor](popover.html#topmost-popover-ancestor){#the-dialog-element:topmost-popover-ancestor-2}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-16
    x-internal="this"}, `document`{.variable}\'s [showing auto popover
    list](popover.html#auto-popover-list){#the-dialog-element:auto-popover-list},
    null, and false.

14. If `hideUntil`{.variable} is null, then set `hideUntil`{.variable}
    to `document`{.variable}.

15. Run [hide all popovers
    until](popover.html#hide-all-popovers-until){#the-dialog-element:hide-all-popovers-until}
    given `hideUntil`{.variable}, false, and true.

16. Run the [dialog focusing
    steps](#dialog-focusing-steps){#the-dialog-element:dialog-focusing-steps-4}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-17
    x-internal="this"}.

The [`showModal()`]{#dom-dialog-showmodal .dfn
dfn-for="HTMLDialogElement" dfn-type="method"} method steps are to [show
a modal
dialog](#show-a-modal-dialog){#the-dialog-element:show-a-modal-dialog}
given
[this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-18
x-internal="this"} and null.

The [`close(``returnValue`{.variable}`)`]{#dom-dialog-close .dfn
dfn-for="HTMLDialogElement" dfn-type="method"} method steps are:

1.  If `returnValue`{.variable} is not given, then set it to null.

2.  [Close the
    dialog](#close-the-dialog){#the-dialog-element:close-the-dialog}
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-19
    x-internal="this"} with `returnValue`{.variable} and null.

The
[`requestClose(``returnValue`{.variable}`)`]{#dom-dialog-requestclose
.dfn dfn-for="HTMLDialogElement" dfn-type="method"} method steps are:

1.  If `returnValue`{.variable} is not given, then set it to null.

2.  [Request to close the
    dialog](#dialog-request-close){#the-dialog-element:dialog-request-close}
    [this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-20
    x-internal="this"} with `returnValue`{.variable} and null.

::: {#note-dialog-method-names .note}
We use show/close as the verbs for
[`dialog`{#the-dialog-element:the-dialog-element-14}](#the-dialog-element)
elements, as opposed to verb pairs that are more commonly thought of as
antonyms such as show/hide or open/close, due to the following
constraints:

- Hiding a dialog is different from closing one. Closing a dialog gives
  it a return value, fires an event, unblocks the page for other
  dialogs, and so on. Whereas hiding a dialog is a purely visual
  property, and is something you can already do with the
  [`hidden`{#the-dialog-element:attr-hidden-2}](interaction.html#attr-hidden)
  attribute or by removing the
  [`open`{#the-dialog-element:attr-dialog-open-9}](#attr-dialog-open)
  attribute. (See also the [note
  above](#note-dialog-remove-open-attribute) about removing the
  [`open`{#the-dialog-element:attr-dialog-open-10}](#attr-dialog-open)
  attribute, and how hiding the dialog in that way is generally not
  desired.)

- Showing a dialog is different from opening one. Opening a dialog
  consists of creating and showing that dialog (similar to how
  [`window.open()`{#the-dialog-element:dom-open}](nav-history-apis.html#dom-open)
  both creates and shows a new window). Whereas showing the dialog is
  the process of taking a
  [`dialog`{#the-dialog-element:the-dialog-element-15}](#the-dialog-element)
  element that is already in the DOM, and making it interactive and
  visible to the user.

- If we were to have a `dialog.open()` method despite the above, it
  would conflict with the
  [`dialog.open`{#the-dialog-element:dom-dialog-open-2}](#dom-dialog-open)
  property.

Furthermore, a
[survey](https://lists.whatwg.org/pipermail/whatwg-whatwg.org/2013-December/041799.html)
of many other UI frameworks contemporary to the original design of the
[`dialog`{#the-dialog-element:the-dialog-element-16}](#the-dialog-element)
element made it clear that the show/close verb pair was reasonably
common.

In summary, it turns out that the implications of certain verbs, and how
they are used in technology contexts, mean that paired actions such as
showing and closing a dialog are not always expressible as antonyms.
:::

The [`returnValue`]{#dom-dialog-returnvalue .dfn
dfn-for="HTMLDialogElement" dfn-type="attribute"} IDL attribute, on
getting, must return the last value to which it was set. On setting, it
must be set to the new value. When the element is created, it must be
set to the empty string.

The [`closedBy`]{#dom-dialog-closedby .dfn dfn-for="HTMLDialogElement"
dfn-type="attribute"} getter steps are to return the keyword
corresponding to the [computed closed-by
state](#computed-closed-by-state){#the-dialog-element:computed-closed-by-state}
given
[this](https://webidl.spec.whatwg.org/#this){#the-dialog-element:this-21
x-internal="this"}.

The
[`closedBy`{#the-dialog-element:dom-dialog-closedby-2}](#dom-dialog-closedby)
setter steps are to set the
[`closedby`{#the-dialog-element:attr-dialog-closedby-5}](#attr-dialog-closedby)
content attribute to the given value.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDialogElement/open](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/open "The open property of the HTMLDialogElement interface is a boolean value reflecting the open HTML attribute, indicating whether the <dialog> is available for interaction.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari15.4+]{.safari
.yes}[Chrome37+]{.chrome .yes}

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
:::::

The [`open`]{#dom-dialog-open .dfn dfn-for="HTMLDialogElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-dialog-element:reflect}
the [`open`{#the-dialog-element:attr-dialog-open-11}](#attr-dialog-open)
content attribute.

------------------------------------------------------------------------

Each [`Document`{#the-dialog-element:document-2}](dom.html#document) has
a [dialog pointerdown target]{#dialog-pointerdown-target .dfn}, which is
an [HTML dialog
element](#htmldialogelement){#the-dialog-element:htmldialogelement} or
null, initially null.

Each [HTML
element](infrastructure.html#html-elements){#the-dialog-element:html-elements}
has a [previously focused element]{#previously-focused-element .dfn}
which is null or an element, and it is initially null. When
[`showModal()`{#the-dialog-element:dom-dialog-showmodal-4}](#dom-dialog-showmodal)
and [`show()`{#the-dialog-element:dom-dialog-show-2}](#dom-dialog-show)
are called, this element is set to the currently
[focused](interaction.html#focused){#the-dialog-element:focused-2}
element before running the [dialog focusing
steps](#dialog-focusing-steps){#the-dialog-element:dialog-focusing-steps-5}.
Elements with the
[`popover`{#the-dialog-element:attr-popover}](popover.html#attr-popover)
attribute set this element to the currently
[focused](interaction.html#focused){#the-dialog-element:focused-3}
element during the [show popover
algorithm](popover.html#show-popover){#the-dialog-element:show-popover}.

Each
[`dialog`{#the-dialog-element:the-dialog-element-17}](#the-dialog-element)
element has a [dialog toggle task tracker]{#dialog-toggle-task-tracker
.dfn}, which is a [toggle task
tracker](interaction.html#toggle-task-tracker){#the-dialog-element:toggle-task-tracker}
or null, initially null.

Each
[`dialog`{#the-dialog-element:the-dialog-element-18}](#the-dialog-element)
element has a [close watcher]{#dialog-close-watcher .dfn}, which is a
[close
watcher](interaction.html#close-watcher){#the-dialog-element:close-watcher}
or null, initially null.

Each
[`dialog`{#the-dialog-element:the-dialog-element-19}](#the-dialog-element)
element has a [request close return value]{#request-close-return-value
.dfn}, which is a string or null, initially null.

Each
[`dialog`{#the-dialog-element:the-dialog-element-20}](#the-dialog-element)
element has a [request close source
element]{#request-close-source-element .dfn}, which is an
[`Element`{#the-dialog-element:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or null, initially null.

Each
[`dialog`{#the-dialog-element:the-dialog-element-21}](#the-dialog-element)
element has an [enable close watcher for request
close]{#enable-close-watcher-for-requestclose() .dfn} boolean, initially
false.

[`dialog`{#the-dialog-element:the-dialog-element-22}](#the-dialog-element)\'s
[enable close watcher for request
close](#enable-close-watcher-for-requestclose()){#the-dialog-element:enable-close-watcher-for-requestclose()}
is used to force enable the dialog\'s [close
watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher}
while [requesting to
close](#dialog-request-close){#the-dialog-element:dialog-request-close-2}
it. A dialog whose [computed closed-by
state](#computed-closed-by-state){#the-dialog-element:computed-closed-by-state-2}
is the
[None](#attr-dialog-closedby-none-state){#the-dialog-element:attr-dialog-closedby-none-state-2}
state would otherwise have a disabled [close
watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-2}.

Each
[`dialog`{#the-dialog-element:the-dialog-element-23}](#the-dialog-element)
element has an [is modal]{#is-modal .dfn} boolean, initially false.

------------------------------------------------------------------------

The
[`dialog`{#the-dialog-element:the-dialog-element-24}](#the-dialog-element)
[HTML element removing
steps](infrastructure.html#html-element-removing-steps){#the-dialog-element:html-element-removing-steps},
given `removedNode`{.variable} and `oldParent`{.variable}, are:

1.  If `removedNode`{.variable}\'s [close
    watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-3}
    is not null, then:

    1.  [Destroy](interaction.html#close-watcher-destroy){#the-dialog-element:close-watcher-destroy}
        `removedNode`{.variable}\'s [close
        watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-4}.

    2.  Set `removedNode`{.variable}\'s [close
        watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-5}
        to null.

2.  If `removedNode`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-4
    x-internal="node-document"}\'s [top
    layer](https://drafts.csswg.org/css-position-4/#document-top-layer){#the-dialog-element:top-layer
    x-internal="top-layer"}
    [contains](https://infra.spec.whatwg.org/#list-contain){#the-dialog-element:list-contains-2
    x-internal="list-contains"} `removedNode`{.variable}, then [remove
    an element from the top layer
    immediately](https://drafts.csswg.org/css-position-4/#remove-an-element-from-the-top-layer-immediately){#the-dialog-element:remove-an-element-from-the-top-layer-immediately
    x-internal="remove-an-element-from-the-top-layer-immediately"} given
    `removedNode`{.variable}.

3.  Set [is modal](#is-modal){#the-dialog-element:is-modal-2} of
    `removedNode`{.variable} to false.

4.  [Remove](https://infra.spec.whatwg.org/#list-remove){#the-dialog-element:list-remove
    x-internal="list-remove"} `removedNode`{.variable} from
    `removedNode`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-5
    x-internal="node-document"}\'s [open dialogs
    list](dom.html#open-dialogs-list){#the-dialog-element:open-dialogs-list-3}.

To [show a modal dialog]{#show-a-modal-dialog .dfn} given a
[`dialog`{#the-dialog-element:the-dialog-element-25}](#the-dialog-element)
element `subject`{.variable} and an
[`Element`{#the-dialog-element:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or null `source`{.variable}:

1.  If `subject`{.variable} has an
    [`open`{#the-dialog-element:attr-dialog-open-12}](#attr-dialog-open)
    attribute and [is modal](#is-modal){#the-dialog-element:is-modal-3}
    of `subject`{.variable} is true, then return.

2.  If `subject`{.variable} has an
    [`open`{#the-dialog-element:attr-dialog-open-13}](#attr-dialog-open)
    attribute, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-dialog-element:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#the-dialog-element:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-6
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#the-dialog-element:fully-active},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-dialog-element:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#the-dialog-element:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  If `subject`{.variable} is not
    [connected](https://dom.spec.whatwg.org/#connected){#the-dialog-element:connected
    x-internal="connected"}, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-dialog-element:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#the-dialog-element:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  If `subject`{.variable} is in the [popover showing
    state](popover.html#popover-showing-state){#the-dialog-element:popover-showing-state},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-dialog-element:invalidstateerror-5
    x-internal="invalidstateerror"}
    [`DOMException`{#the-dialog-element:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  If the result of [firing an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-dialog-element:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`beforetoggle`{#the-dialog-element:event-beforetoggle-2}](indices.html#event-beforetoggle),
    using
    [`ToggleEvent`{#the-dialog-element:toggleevent-2}](interaction.html#toggleevent),
    with the
    [`cancelable`{#the-dialog-element:dom-event-cancelable-2}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    attribute initialized to true, the
    [`oldState`{#the-dialog-element:dom-toggleevent-oldstate-2}](interaction.html#dom-toggleevent-oldstate)
    attribute initialized to \"`closed`\", the
    [`newState`{#the-dialog-element:dom-toggleevent-newstate-2}](interaction.html#dom-toggleevent-newstate)
    attribute initialized to \"`open`\", and the
    [`source`{#the-dialog-element:dom-toggleevent-source}](interaction.html#dom-toggleevent-source)
    attribute initialized to `source`{.variable} at `subject`{.variable}
    is false, then return.

7.  If `subject`{.variable} has an
    [`open`{#the-dialog-element:attr-dialog-open-14}](#attr-dialog-open)
    attribute, then return.

8.  If `subject`{.variable} is not
    [connected](https://dom.spec.whatwg.org/#connected){#the-dialog-element:connected-2
    x-internal="connected"}, then return.

9.  If `subject`{.variable} is in the [popover showing
    state](popover.html#popover-showing-state){#the-dialog-element:popover-showing-state-2},
    then return.

10. [Queue a dialog toggle event
    task](#queue-a-dialog-toggle-event-task){#the-dialog-element:queue-a-dialog-toggle-event-task-2}
    given `subject`{.variable}, \"`closed`\", \"`open`\", and
    `source`{.variable}.

11. Add an
    [`open`{#the-dialog-element:attr-dialog-open-15}](#attr-dialog-open)
    attribute to `subject`{.variable}, whose value is the empty string.

12. Set [is modal](#is-modal){#the-dialog-element:is-modal-4} of
    `subject`{.variable} to true.

13. [Assert](https://infra.spec.whatwg.org/#assert){#the-dialog-element:assert-2
    x-internal="assert"}: `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-7
    x-internal="node-document"}\'s [open dialogs
    list](dom.html#open-dialogs-list){#the-dialog-element:open-dialogs-list-4}
    does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#the-dialog-element:list-contains-3
    x-internal="list-contains"} `subject`{.variable}.

14. Add `subject`{.variable} to `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-8
    x-internal="node-document"}\'s [open dialogs
    list](dom.html#open-dialogs-list){#the-dialog-element:open-dialogs-list-5}.

15. Set `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-9
    x-internal="node-document"} to be [blocked by the modal
    dialog](interaction.html#blocked-by-a-modal-dialog){#the-dialog-element:blocked-by-a-modal-dialog-2}
    `subject`{.variable}.

    This will cause the [focused area of the
    document](interaction.html#focused-area-of-the-document){#the-dialog-element:focused-area-of-the-document}
    to become [inert](interaction.html#inert){#the-dialog-element:inert}
    (unless that currently focused area is a [shadow-including
    descendant](https://dom.spec.whatwg.org/#concept-shadow-including-descendant){#the-dialog-element:shadow-including-descendant
    x-internal="shadow-including-descendant"} of `subject`{.variable}).
    In such cases, the [focused area of the
    document](interaction.html#focused-area-of-the-document){#the-dialog-element:focused-area-of-the-document-2}
    will soon be [reset](webappapis.html#focus-fixup-rule) to the
    [viewport](https://drafts.csswg.org/css2/#viewport){#the-dialog-element:viewport
    x-internal="viewport"}. In a couple steps we will attempt to find a
    better candidate to focus.

16. If `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-10
    x-internal="node-document"}\'s [top
    layer](https://drafts.csswg.org/css-position-4/#document-top-layer){#the-dialog-element:top-layer-2
    x-internal="top-layer"} does not already
    [contain](https://infra.spec.whatwg.org/#list-contain){#the-dialog-element:list-contains-4
    x-internal="list-contains"} `subject`{.variable}, then [add an
    element to the top
    layer](https://drafts.csswg.org/css-position-4/#add-an-element-to-the-top-layer){#the-dialog-element:add-an-element-to-the-top-layer
    x-internal="add-an-element-to-the-top-layer"} given
    `subject`{.variable}.

17. [Set the dialog close
    watcher](#set-the-dialog-close-watcher){#the-dialog-element:set-the-dialog-close-watcher-2}
    with `subject`{.variable}.

18. Set `subject`{.variable}\'s [previously focused
    element](#previously-focused-element){#the-dialog-element:previously-focused-element-2}
    to the
    [focused](interaction.html#focused){#the-dialog-element:focused-4}
    element.

19. Let `document`{.variable} be `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-11
    x-internal="node-document"}.

20. Let `hideUntil`{.variable} be the result of running [topmost popover
    ancestor](popover.html#topmost-popover-ancestor){#the-dialog-element:topmost-popover-ancestor-3}
    given `subject`{.variable}, `document`{.variable}\'s [showing hint
    popover
    list](popover.html#showing-hint-popover-list){#the-dialog-element:showing-hint-popover-list-2},
    null, and false.

21. If `hideUntil`{.variable} is null, then set `hideUntil`{.variable}
    to the result of running [topmost popover
    ancestor](popover.html#topmost-popover-ancestor){#the-dialog-element:topmost-popover-ancestor-4}
    given `subject`{.variable}, `document`{.variable}\'s [showing auto
    popover
    list](popover.html#auto-popover-list){#the-dialog-element:auto-popover-list-2},
    null, and false.

22. If `hideUntil`{.variable} is null, then set `hideUntil`{.variable}
    to `document`{.variable}.

23. Run [hide all popovers
    until](popover.html#hide-all-popovers-until){#the-dialog-element:hide-all-popovers-until-2}
    given `hideUntil`{.variable}, false, and true.

24. Run the [dialog focusing
    steps](#dialog-focusing-steps){#the-dialog-element:dialog-focusing-steps-6}
    given `subject`{.variable}.

To [set the dialog close watcher]{#set-the-dialog-close-watcher .dfn},
given a
[`dialog`{#the-dialog-element:the-dialog-element-26}](#the-dialog-element)
element `dialog`{.variable}:

1.  Set `dialog`{.variable}\'s [close
    watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-6}
    to the result of [establishing a close
    watcher](interaction.html#establish-a-close-watcher){#the-dialog-element:establish-a-close-watcher}
    given `dialog`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-dialog-element:concept-relevant-global},
    with:

    - *[cancelAction](interaction.html#create-close-watcher-cancelaction)*
      given `canPreventClose`{.variable} being to return the result of
      [firing an
      event](https://dom.spec.whatwg.org/#concept-event-fire){#the-dialog-element:concept-event-fire-3
      x-internal="concept-event-fire"} named
      [`cancel`{#the-dialog-element:event-cancel-3}](indices.html#event-cancel)
      at `dialog`{.variable}, with the
      [`cancelable`{#the-dialog-element:dom-event-cancelable-3}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
      attribute initialized to `canPreventClose`{.variable}.

    - *[closeAction](interaction.html#create-close-watcher-closeaction)*
      being to [close the
      dialog](#close-the-dialog){#the-dialog-element:close-the-dialog-2}
      given `dialog`{.variable}, `dialog`{.variable}\'s [request close
      return
      value](#request-close-return-value){#the-dialog-element:request-close-return-value},
      and `dialog`{.variable}\'s [request close source
      element](#request-close-source-element){#the-dialog-element:request-close-source-element}.

    - *[getEnabledState](interaction.html#create-close-watcher-getenabledstate)*
      being to return true if `dialog`{.variable}\'s [enable close
      watcher for request
      close](#enable-close-watcher-for-requestclose()){#the-dialog-element:enable-close-watcher-for-requestclose()-2}
      is true or `dialog`{.variable}\'s [computed closed-by
      state](#computed-closed-by-state){#the-dialog-element:computed-closed-by-state-3}
      is not
      [None](#attr-dialog-closedby-none-state){#the-dialog-element:attr-dialog-closedby-none-state-3};
      otherwise false.

The [is valid invoker command
steps](form-elements.html#is-valid-invoker-command-steps){#the-dialog-element:is-valid-invoker-command-steps}
for
[`dialog`{#the-dialog-element:the-dialog-element-27}](#the-dialog-element)
elements, given a
[`command`{#the-dialog-element:attr-button-command}](form-elements.html#attr-button-command)
attribute `command`{.variable}, are:

1.  If `command`{.variable} is in the
    [Close](form-elements.html#attr-button-command-close-state){#the-dialog-element:attr-button-command-close-state}
    state, the [Request
    Close](form-elements.html#attr-button-command-request-close-state){#the-dialog-element:attr-button-command-request-close-state}
    state, or the [Show
    Modal](form-elements.html#attr-button-command-show-modal-state){#the-dialog-element:attr-button-command-show-modal-state}
    state, then return true.

2.  Return false.

The [invoker command
steps](form-elements.html#invoker-command-steps){#the-dialog-element:invoker-command-steps}
for
[`dialog`{#the-dialog-element:the-dialog-element-28}](#the-dialog-element)
elements, given an element `element`{.variable}, an element
`invoker`{.variable}, and a
[`command`{#the-dialog-element:attr-button-command-2}](form-elements.html#attr-button-command)
attribute `command`{.variable}, are:

1.  If `element`{.variable} is in the [popover
    showing](popover.html#popover-showing-state){#the-dialog-element:popover-showing-state-3}
    state, then return.

2.  If `command`{.variable} is in the
    [Close](form-elements.html#attr-button-command-close-state){#the-dialog-element:attr-button-command-close-state-2}
    state and `element`{.variable} has an
    [`open`{#the-dialog-element:attr-dialog-open-16}](#attr-dialog-open)
    attribute, then [close the
    dialog](#close-the-dialog){#the-dialog-element:close-the-dialog-3}
    `element`{.variable} with `invoker`{.variable}\'s [optional
    value](form-control-infrastructure.html#concept-fe-optional-value){#the-dialog-element:concept-fe-optional-value}
    and `invoker`{.variable}.

3.  If `command`{.variable} is in the [Request
    Close](form-elements.html#attr-button-command-request-close-state){#the-dialog-element:attr-button-command-request-close-state-2}
    state and `element`{.variable} has an
    [`open`{#the-dialog-element:attr-dialog-open-17}](#attr-dialog-open)
    attribute, then [request to close the
    dialog](#dialog-request-close){#the-dialog-element:dialog-request-close-3}
    `element`{.variable} with `invoker`{.variable}\'s [optional
    value](form-control-infrastructure.html#concept-fe-optional-value){#the-dialog-element:concept-fe-optional-value-2}
    and `invoker`{.variable}.

4.  If `command`{.variable} is the [Show
    Modal](form-elements.html#attr-button-command-show-modal-state){#the-dialog-element:attr-button-command-show-modal-state-2}
    state and `element`{.variable} does not have an
    [`open`{#the-dialog-element:attr-dialog-open-18}](#attr-dialog-open)
    attribute, then [show a modal
    dialog](#show-a-modal-dialog){#the-dialog-element:show-a-modal-dialog-2}
    given `element`{.variable} and `invoker`{.variable}.

::: example
The following buttons use
[`commandfor`{#the-dialog-element:attr-button-commandfor}](form-elements.html#attr-button-commandfor)
to open and close a \"confirm\"
[`dialog`{#the-dialog-element:the-dialog-element-29}](#the-dialog-element)
as modal when activated:

``` html
<button type=button
        commandfor="the-dialog"
        command="show-modal">
 Delete
</button>
<dialog id="the-dialog">
 <form action="/delete" method="POST">
  <button type="submit">
   Delete
  </button>
  <button commandfor="the-dialog"
          command="close">
   Cancel
  </button>
 </form>
</dialog>
    
```
:::

When a
[`dialog`{#the-dialog-element:the-dialog-element-30}](#the-dialog-element)
element `subject`{.variable} is to be [closed]{#close-the-dialog .dfn},
with null or a string `result`{.variable} and an
[`Element`{#the-dialog-element:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or null `source`{.variable}, run these steps:

1.  If `subject`{.variable} does not have an
    [`open`{#the-dialog-element:attr-dialog-open-19}](#attr-dialog-open)
    attribute, then return.

2.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-dialog-element:concept-event-fire-4
    x-internal="concept-event-fire"} named
    [`beforetoggle`{#the-dialog-element:event-beforetoggle-3}](indices.html#event-beforetoggle),
    using
    [`ToggleEvent`{#the-dialog-element:toggleevent-3}](interaction.html#toggleevent),
    with the
    [`oldState`{#the-dialog-element:dom-toggleevent-oldstate-3}](interaction.html#dom-toggleevent-oldstate)
    attribute initialized to \"`open`\", the
    [`newState`{#the-dialog-element:dom-toggleevent-newstate-3}](interaction.html#dom-toggleevent-newstate)
    attribute initialized to \"`closed`\", and the
    [`source`{#the-dialog-element:dom-toggleevent-source-2}](interaction.html#dom-toggleevent-source)
    attribute initialized to `source`{.variable} at
    `subject`{.variable}.

3.  If `subject`{.variable} does not have an
    [`open`{#the-dialog-element:attr-dialog-open-20}](#attr-dialog-open)
    attribute, then return.

4.  [Queue a dialog toggle event
    task](#queue-a-dialog-toggle-event-task){#the-dialog-element:queue-a-dialog-toggle-event-task-3}
    given `subject`{.variable}, \"`open`\", \"`closed`\", and
    `source`{.variable}.

5.  Remove `subject`{.variable}\'s
    [`open`{#the-dialog-element:attr-dialog-open-21}](#attr-dialog-open)
    attribute.

6.  If [is modal](#is-modal){#the-dialog-element:is-modal-5} of
    `subject`{.variable} is true, then [request an element to be removed
    from the top
    layer](https://drafts.csswg.org/css-position-4/#request-an-element-to-be-removed-from-the-top-layer){#the-dialog-element:request-an-element-to-be-removed-from-the-top-layer
    x-internal="request-an-element-to-be-removed-from-the-top-layer"}
    given `subject`{.variable}.

7.  Let `wasModal`{.variable} be the value of `subject`{.variable}\'s
    [is modal](#is-modal){#the-dialog-element:is-modal-6} flag.

8.  Set [is modal](#is-modal){#the-dialog-element:is-modal-7} of
    `subject`{.variable} to false.

9.  [Remove](https://infra.spec.whatwg.org/#list-remove){#the-dialog-element:list-remove-2
    x-internal="list-remove"} `subject`{.variable} from
    `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-12
    x-internal="node-document"}\'s [open dialogs
    list](dom.html#open-dialogs-list){#the-dialog-element:open-dialogs-list-6}.

10. If `result`{.variable} is not null, then set `subject`{.variable}\'s
    [`returnValue`{#the-dialog-element:dom-dialog-returnvalue-2}](#dom-dialog-returnvalue)
    attribute to `result`{.variable}.

11. Set `subject`{.variable}\'s [request close return
    value](#request-close-return-value){#the-dialog-element:request-close-return-value-2}
    to null.

12. Set `subject`{.variable}\'s [request close source
    element](#request-close-source-element){#the-dialog-element:request-close-source-element-2}
    to null.

13. If `subject`{.variable}\'s [previously focused
    element](#previously-focused-element){#the-dialog-element:previously-focused-element-3}
    is not null, then:

    1.  Let `element`{.variable} be `subject`{.variable}\'s [previously
        focused
        element](#previously-focused-element){#the-dialog-element:previously-focused-element-4}.

    2.  Set `subject`{.variable}\'s [previously focused
        element](#previously-focused-element){#the-dialog-element:previously-focused-element-5}
        to null.

    3.  If `subject`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-13
        x-internal="node-document"}\'s [focused area of the
        document](interaction.html#focused-area-of-the-document){#the-dialog-element:focused-area-of-the-document-3}\'s
        [DOM
        anchor](interaction.html#dom-anchor){#the-dialog-element:dom-anchor}
        is a [shadow-including inclusive
        descendant](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-descendant){#the-dialog-element:shadow-including-inclusive-descendant
        x-internal="shadow-including-inclusive-descendant"} of
        `subject`{.variable}, or `wasModal`{.variable} is true, then run
        the [focusing
        steps](interaction.html#focusing-steps){#the-dialog-element:focusing-steps}
        for `element`{.variable}; the viewport should not be scrolled by
        doing this step.

14. [Queue an element
    task](webappapis.html#queue-an-element-task){#the-dialog-element:queue-an-element-task}
    on the [user interaction task
    source](webappapis.html#user-interaction-task-source){#the-dialog-element:user-interaction-task-source}
    given the `subject`{.variable} element to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-dialog-element:concept-event-fire-5
    x-internal="concept-event-fire"} named
    [`close`{#the-dialog-element:event-close-4}](indices.html#event-close)
    at `subject`{.variable}.

15. If `subject`{.variable}\'s [close
    watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-7}
    is not null, then:

    1.  [Destroy](interaction.html#close-watcher-destroy){#the-dialog-element:close-watcher-destroy-2}
        `subject`{.variable}\'s [close
        watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-8}.

    2.  Set `subject`{.variable}\'s [close
        watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-9}
        to null.

To [request to close]{#dialog-request-close .dfn}
[`dialog`{#the-dialog-element:the-dialog-element-31}](#the-dialog-element)
element `subject`{.variable}, given null or a string
`returnValue`{.variable} and null or an
[`Element`{#the-dialog-element:element-4}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
`source`{.variable}:

1.  If `subject`{.variable} does not have an
    [`open`{#the-dialog-element:attr-dialog-open-22}](#attr-dialog-open)
    attribute, then return.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#the-dialog-element:assert-3
    x-internal="assert"}: `subject`{.variable}\'s [close
    watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-10}
    is not null.

3.  Set `subject`{.variable}\'s [enable close watcher for request
    close](#enable-close-watcher-for-requestclose()){#the-dialog-element:enable-close-watcher-for-requestclose()-3}
    to true.

4.  Set `subject`{.variable}\'s [request close return
    value](#request-close-return-value){#the-dialog-element:request-close-return-value-3}
    to `returnValue`{.variable}.

5.  Set `subject`{.variable}\'s [request close source
    element](#request-close-source-element){#the-dialog-element:request-close-source-element-3}
    to `source`{.variable}.

6.  [Request to
    close](interaction.html#close-watcher-request-close){#the-dialog-element:close-watcher-request-close}
    `subject`{.variable}\'s [close
    watcher](#dialog-close-watcher){#the-dialog-element:dialog-close-watcher-11}
    with false.

7.  Set `subject`{.variable}\'s [enable close watcher for request
    close](#enable-close-watcher-for-requestclose()){#the-dialog-element:enable-close-watcher-for-requestclose()-4}
    to false.

To [queue a dialog toggle event task]{#queue-a-dialog-toggle-event-task
.dfn} given a
[`dialog`{#the-dialog-element:the-dialog-element-32}](#the-dialog-element)
element `element`{.variable}, a string `oldState`{.variable}, a string
`newState`{.variable}, and an
[`Element`{#the-dialog-element:element-5}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or null `source`{.variable}:

1.  If `element`{.variable}\'s [dialog toggle task
    tracker](#dialog-toggle-task-tracker){#the-dialog-element:dialog-toggle-task-tracker}
    is not null, then:

    1.  Set `oldState`{.variable} to `element`{.variable}\'s [dialog
        toggle task
        tracker](#dialog-toggle-task-tracker){#the-dialog-element:dialog-toggle-task-tracker-2}\'s
        [old
        state](interaction.html#toggle-task-old-state){#the-dialog-element:toggle-task-old-state}.

    2.  Remove `element`{.variable}\'s [dialog toggle task
        tracker](#dialog-toggle-task-tracker){#the-dialog-element:dialog-toggle-task-tracker-3}\'s
        [task](interaction.html#toggle-task-task){#the-dialog-element:toggle-task-task}
        from its [task
        queue](webappapis.html#task-queue){#the-dialog-element:task-queue}.

    3.  Set `element`{.variable}\'s [dialog toggle task
        tracker](#dialog-toggle-task-tracker){#the-dialog-element:dialog-toggle-task-tracker-4}
        to null.

2.  [Queue an element
    task](webappapis.html#queue-an-element-task){#the-dialog-element:queue-an-element-task-2}
    given the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#the-dialog-element:dom-manipulation-task-source}
    and `element`{.variable} to run the following steps:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-dialog-element:concept-event-fire-6
        x-internal="concept-event-fire"} named
        [`toggle`{#the-dialog-element:event-toggle}](indices.html#event-toggle)
        at `element`{.variable}, using
        [`ToggleEvent`{#the-dialog-element:toggleevent-4}](interaction.html#toggleevent),
        with the
        [`oldState`{#the-dialog-element:dom-toggleevent-oldstate-4}](interaction.html#dom-toggleevent-oldstate)
        attribute initialized to `oldState`{.variable}, the
        [`newState`{#the-dialog-element:dom-toggleevent-newstate-4}](interaction.html#dom-toggleevent-newstate)
        attribute initialized to `newState`{.variable}, and the
        [`source`{#the-dialog-element:dom-toggleevent-source-3}](interaction.html#dom-toggleevent-source)
        attribute initialized to `source`{.variable}.

    2.  Set `element`{.variable}\'s [dialog toggle task
        tracker](#dialog-toggle-task-tracker){#the-dialog-element:dialog-toggle-task-tracker-5}
        to null.

3.  Set `element`{.variable}\'s [dialog toggle task
    tracker](#dialog-toggle-task-tracker){#the-dialog-element:dialog-toggle-task-tracker-6}
    to a struct with
    [task](interaction.html#toggle-task-task){#the-dialog-element:toggle-task-task-2}
    set to the just-queued
    [task](webappapis.html#concept-task){#the-dialog-element:concept-task}
    and [old
    state](interaction.html#toggle-task-old-state){#the-dialog-element:toggle-task-old-state-2}
    set to `oldState`{.variable}.

To retrieve a dialog\'s [computed closed-by
state]{#computed-closed-by-state .dfn}, given a
[`dialog`{#the-dialog-element:the-dialog-element-33}](#the-dialog-element)
`dialog`{.variable}:

1.  If the state of `dialog`{.variable}\'s
    [`closedby`{#the-dialog-element:attr-dialog-closedby-6}](#attr-dialog-closedby)
    attribute is
    [Auto](#attr-dialog-closedby-auto-state){#the-dialog-element:attr-dialog-closedby-auto-state-2}:

    1.  If `dialog`{.variable}\'s [is
        modal](#is-modal){#the-dialog-element:is-modal-8} is true, then
        return [Close
        Request](#attr-dialog-closedby-closerequest-state){#the-dialog-element:attr-dialog-closedby-closerequest-state-2}.

    2.  Return
        [None](#attr-dialog-closedby-none-state){#the-dialog-element:attr-dialog-closedby-none-state-4}.

2.  Return the state of `dialog`{.variable}\'s
    [`closedby`{#the-dialog-element:attr-dialog-closedby-7}](#attr-dialog-closedby)
    attribute.

The [dialog focusing steps]{#dialog-focusing-steps .dfn}, given a
[`dialog`{#the-dialog-element:the-dialog-element-34}](#the-dialog-element)
element `subject`{.variable}, are as follows:

1.  If the [allow focus
    steps](interaction.html#allow-focus-steps){#the-dialog-element:allow-focus-steps}
    given `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-14
    x-internal="node-document"} return false, then return.

2.  Let `control`{.variable} be null.

3.  If `subject`{.variable} has the
    [`autofocus`{#the-dialog-element:attr-fe-autofocus-7}](interaction.html#attr-fe-autofocus)
    attribute, then set `control`{.variable} to `subject`{.variable}.

4.  If `control`{.variable} is null, then set `control`{.variable} to
    the [focus
    delegate](interaction.html#focus-delegate){#the-dialog-element:focus-delegate}
    of `subject`{.variable}.

5.  If `control`{.variable} is null, then set `control`{.variable} to
    `subject`{.variable}.

6.  Run the [focusing
    steps](interaction.html#focusing-steps){#the-dialog-element:focusing-steps-2}
    for `control`{.variable}.

    If `control`{.variable} is not
    [focusable](interaction.html#focusable){#the-dialog-element:focusable},
    this will do nothing. This would only happen if subject had no focus
    delegate, and the user agent decided that
    [`dialog`{#the-dialog-element:the-dialog-element-35}](#the-dialog-element)
    elements were not generally focusable. In that case, any [earlier
    modifications](#note-dialog-plus-focus-fixup) to the [focused area
    of the
    document](interaction.html#focused-area-of-the-document){#the-dialog-element:focused-area-of-the-document-4}
    will apply.

7.  Let `topDocument`{.variable} be `control`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-dialog-element:node-navigable}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#the-dialog-element:nav-top}\'s
    [active
    document](document-sequences.html#nav-document){#the-dialog-element:nav-document}.

8.  If `control`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-dialog-element:node-document-15
    x-internal="node-document"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-dialog-element:concept-document-origin
    x-internal="concept-document-origin"} is not the
    [same](browsers.html#same-origin){#the-dialog-element:same-origin}
    as the
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-dialog-element:concept-document-origin-2
    x-internal="concept-document-origin"} of `topDocument`{.variable},
    then return.

9.  [Empty](https://infra.spec.whatwg.org/#list-empty){#the-dialog-element:list-empty
    x-internal="list-empty"} `topDocument`{.variable}\'s [autofocus
    candidates](interaction.html#autofocus-candidates){#the-dialog-element:autofocus-candidates}.

10. Set `topDocument`{.variable}\'s [autofocus processed
    flag](interaction.html#autofocus-processed-flag){#the-dialog-element:autofocus-processed-flag}
    to true.

#### [4.11.5]{.secno} [Dialog light dismiss]{.dfn}[](#dialog-light-dismiss){.self-link}

\"Light dismiss\" means that clicking outside of a
[`dialog`{#dialog-light-dismiss:the-dialog-element}](#the-dialog-element)
element whose
[`closedby`{#dialog-light-dismiss:attr-dialog-closedby}](#attr-dialog-closedby)
attribute is in the
[Any](#attr-dialog-closedby-any-state){#dialog-light-dismiss:attr-dialog-closedby-any-state}
state will close the
[`dialog`{#dialog-light-dismiss:the-dialog-element-2}](#the-dialog-element)
element. This is in addition to how such
[`dialog`{#dialog-light-dismiss:the-dialog-element-3}](#the-dialog-element)s
respond to [close
requests](interaction.html#close-request){#dialog-light-dismiss:close-request}.

To [light dismiss open dialogs]{#light-dismiss-open-dialogs .dfn}, given
a
[`PointerEvent`{#dialog-light-dismiss:pointerevent}](https://w3c.github.io/pointerevents/#pointerevent-interface){x-internal="pointerevent"}
`event`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#dialog-light-dismiss:assert
    x-internal="assert"}: `event`{.variable}\'s
    [`isTrusted`{#dialog-light-dismiss:dom-event-istrusted}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
    attribute is true.

2.  Let `document`{.variable} be `event`{.variable}\'s
    [target](https://dom.spec.whatwg.org/#concept-event-target){#dialog-light-dismiss:concept-event-target
    x-internal="concept-event-target"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#dialog-light-dismiss:node-document
    x-internal="node-document"}.

3.  If `document`{.variable}\'s [open dialogs
    list](dom.html#open-dialogs-list){#dialog-light-dismiss:open-dialogs-list}
    is
    [empty](https://infra.spec.whatwg.org/#list-is-empty){#dialog-light-dismiss:list-is-empty
    x-internal="list-is-empty"}, then return.

4.  Let `ancestor`{.variable} be the result of running [nearest clicked
    dialog](#nearest-clicked-dialog){#dialog-light-dismiss:nearest-clicked-dialog}
    given `event`{.variable}.

5.  If `event`{.variable}\'s
    [`type`{#dialog-light-dismiss:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    is
    \"[`pointerdown`{#dialog-light-dismiss:event-pointerdown}](https://w3c.github.io/pointerevents/#the-pointerdown-event){x-internal="event-pointerdown"}\",
    then set `document`{.variable}\'s [dialog pointerdown
    target](#dialog-pointerdown-target){#dialog-light-dismiss:dialog-pointerdown-target}
    to `ancestor`{.variable}.

6.  If `event`{.variable}\'s
    [`type`{#dialog-light-dismiss:dom-event-type-2}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    is
    \"[`pointerup`{#dialog-light-dismiss:event-pointerup}](https://w3c.github.io/pointerevents/#the-pointerup-event){x-internal="event-pointerup"}\",
    then:

    1.  Let `sameTarget`{.variable} be true if `ancestor`{.variable} is
        `document`{.variable}\'s [dialog pointerdown
        target](#dialog-pointerdown-target){#dialog-light-dismiss:dialog-pointerdown-target-2}.

    2.  Set `document`{.variable}\'s [dialog pointerdown
        target](#dialog-pointerdown-target){#dialog-light-dismiss:dialog-pointerdown-target-3}
        to null.

    3.  If `sameTarget`{.variable} is false, then return.

    4.  Let `topmostDialog`{.variable} be the last element of
        `document`{.variable}\'s [open dialogs
        list](dom.html#open-dialogs-list){#dialog-light-dismiss:open-dialogs-list-2}.

    5.  If `ancestor`{.variable} is `topmostDialog`{.variable}, then
        return.

    6.  If `topmostDialog`{.variable}\'s [computed closed-by
        state](#computed-closed-by-state){#dialog-light-dismiss:computed-closed-by-state}
        is not
        [Any](#attr-dialog-closedby-any-state){#dialog-light-dismiss:attr-dialog-closedby-any-state-2},
        then return.

    7.  [Assert](https://infra.spec.whatwg.org/#assert){#dialog-light-dismiss:assert-2
        x-internal="assert"}: `topmostDialog`{.variable}\'s [close
        watcher](#dialog-close-watcher){#dialog-light-dismiss:dialog-close-watcher}
        is not null.

    8.  [Request to
        close](interaction.html#close-watcher-request-close){#dialog-light-dismiss:close-watcher-request-close}
        `topmostDialog`{.variable}\'s [close
        watcher](#dialog-close-watcher){#dialog-light-dismiss:dialog-close-watcher-2}
        with false.

To [run light dismiss activities]{#run-light-dismiss-activities .dfn},
given a
[`PointerEvent`{#dialog-light-dismiss:pointerevent-2}](https://w3c.github.io/pointerevents/#pointerevent-interface){x-internal="pointerevent"}
`event`{.variable}:

1.  Run [light dismiss open
    popovers](popover.html#light-dismiss-open-popovers){#dialog-light-dismiss:light-dismiss-open-popovers}
    with `event`{.variable}.

2.  Run [light dismiss open
    dialogs](#light-dismiss-open-dialogs){#dialog-light-dismiss:light-dismiss-open-dialogs}
    with `event`{.variable}.

[Run light dismiss
activities](#run-light-dismiss-activities){#dialog-light-dismiss:run-light-dismiss-activities}
will be called by the [Pointer Events
spec](https://github.com/w3c/pointerevents/pull/460) when the user
clicks or touches anywhere on the page.

To find the [nearest clicked dialog]{#nearest-clicked-dialog .dfn},
given a
[`PointerEvent`{#dialog-light-dismiss:pointerevent-3}](https://w3c.github.io/pointerevents/#pointerevent-interface){x-internal="pointerevent"}
`event`{.variable}:

1.  Let `target`{.variable} be `event`{.variable}\'s
    [target](https://dom.spec.whatwg.org/#concept-event-target){#dialog-light-dismiss:concept-event-target-2
    x-internal="concept-event-target"}.

2.  If `target`{.variable} is a
    [`dialog`{#dialog-light-dismiss:the-dialog-element-4}](#the-dialog-element)
    element, `target`{.variable} has an
    [`open`{#dialog-light-dismiss:attr-dialog-open}](#attr-dialog-open)
    attribute, `target`{.variable}\'s [is
    modal](#is-modal){#dialog-light-dismiss:is-modal} is true, and
    `event`{.variable}\'s
    [`clientX`{#dialog-light-dismiss:mouseevent-clientx}](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clientx){x-internal="mouseevent-clientx"}
    and
    [`clientY`{#dialog-light-dismiss:mouseevent-clienty}](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clienty){x-internal="mouseevent-clienty"}
    are outside the bounds of `target`{.variable}, then return null.

    The check for
    [`clientX`{#dialog-light-dismiss:mouseevent-clientx-2}](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clientx){x-internal="mouseevent-clientx"}
    and
    [`clientY`{#dialog-light-dismiss:mouseevent-clienty-2}](https://drafts.csswg.org/cssom-view/#dom-mouseevent-clienty){x-internal="mouseevent-clienty"}
    is because a pointer event that hits the `::backdrop` pseudo element
    of a dialog will result in `event`{.variable} having a target of the
    dialog element itself.

3.  Let `currentNode`{.variable} be `target`{.variable}.

4.  While `currentNode`{.variable} is not null:

    1.  If `currentNode`{.variable} is a
        [`dialog`{#dialog-light-dismiss:the-dialog-element-5}](#the-dialog-element)
        element and `currentNode`{.variable} has an
        [`open`{#dialog-light-dismiss:attr-dialog-open-2}](#attr-dialog-open)
        attribute, then return `currentNode`{.variable}.

    2.  Set `currentNode`{.variable} to `currentNode`{.variable}\'s
        parent in the [flat
        tree](https://drafts.csswg.org/css-scoping/#flat-tree){#dialog-light-dismiss:flat-tree
        x-internal="flat-tree"}.

5.  Return null.

[← 4.10.17 Form control
infrastructure](form-control-infrastructure.html) --- [Table of
Contents](./) --- [4.12 Scripting →](scripting.html)
