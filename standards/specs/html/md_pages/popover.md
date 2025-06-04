HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 6.11 Drag and drop](dnd.html) --- [Table of Contents](./) --- [7
Loading web pages →](browsers.html)

1.  ::: {#toc-editing}
    1.  [[6.12]{.secno} The `popover`
        attribute](popover.html#the-popover-attribute)
        1.  [[6.12.1]{.secno} The popover target
            attributes](popover.html#the-popover-target-attributes)
        2.  [[6.12.2]{.secno} Popover light
            dismiss](popover.html#popover-light-dismiss)
    :::

### [6.12]{.secno} The [`popover`{#the-popover-attribute:attr-popover}](#attr-popover) attribute[](#the-popover-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/popover](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/popover "The popover global attribute is used to designate an element as a popover element.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safaripreview+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

All [HTML
elements](infrastructure.html#html-elements){#the-popover-attribute:html-elements}
may have the [`popover`]{#attr-popover .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute set. When specified, the
element won\'t be rendered until it becomes shown, at which point it
will be rendered on top of other page content.

::: example
The [`popover`{#the-popover-attribute:attr-popover-2}](#attr-popover)
attribute is a global attribute that allows authors flexibility to
ensure popover functionality can be applied to elements with the most
relevant semantics.

The following demonstrates how one might create a popover sub-navigation
list of links, within the global navigation for a website.

``` html
<ul>
  <li>
    <a href=...>All Products</a>
    <button popovertarget=sub-nav>
     <img src=down-arrow.png alt="Product pages">
    </button>
    <ul popover id=sub-nav>
     <li><a href=...>Shirts</a>
     <li><a href=...>Shoes</a>
     <li><a href=...>Hats etc.</a>
    </ul>
  </li>
  <!-- other list items and links here -->
</ul>
```
:::

When using
[`popover`{#the-popover-attribute:attr-popover-3}](#attr-popover) on
elements without accessibility semantics, for instance the
[`div`{#the-popover-attribute:the-div-element}](grouping-content.html#the-div-element)
element, authors should use the appropriate ARIA attributes to ensure
the popover is accessible.

::: example
The following shows the baseline markup to create a custom menu popover,
where the first menuitem will receive keyboard focus when the popover is
invoked due to the use of the `autofocus` attribute. Navigating the
menuitems with arrow keys and activation behaviors would still need
author scripting. Additional requirements for building custom menus
widgets are defined in the [WAI-ARIA
specification](https://w3c.github.io/aria/#menu).

``` html
<button popovertarget=m>Actions</button>
<div role=menu id=m popover>
  <button role=menuitem tabindex=-1 autofocus>Edit</button>
  <button role=menuitem tabindex=-1>Hide</button>
  <button role=menuitem tabindex=-1>Delete</button>
</div>
```

A popover can be useful for rendering a status message, confirming the
action performed by the user. The following demonstrates how one could
reveal a popover in an
[`output`{#the-popover-attribute:the-output-element}](form-elements.html#the-output-element)
element.

``` html
<button id=submit>Submit</button>
<p><output><span popover=manual></span></output></p>

<script>
 const sBtn = document.getElementById("submit");
 const outSpan = document.querySelector("output [popover=manual]");
 let successMessage;
 let errorMessage;

 /* define logic for determining success of action
  and determining the appropriate success or error
  messages to use */

 sBtn.addEventListener("click", ()=> {
  if ( success ) {
   outSpan.textContent = successMessage;
  }
  else {
   outSpan.textContent = errorMessage;
  }
  outSpan.showPopover();

  setTimeout(function () {
   outSpan.hidePopover();
  }, 10000);
 });
</script>
```
:::

Inserting a popover element into an
[`output`{#the-popover-attribute:the-output-element-2}](form-elements.html#the-output-element)
element will generally cause screen readers to announce the content when
it becomes visible. Depending on the complexity or frequency of the
content, this could be either useful or annoying to users of these
assistive technologies. Keep this in mind when using the
[`output`{#the-popover-attribute:the-output-element-3}](form-elements.html#the-output-element)
element or other ARIA live regions to ensure the best user experience.

The [`popover`{#the-popover-attribute:attr-popover-4}](#attr-popover)
attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-popover-attribute:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`auto`]{#attr-popover-auto .dfn dfn-for="html-global/popover"
dfn-type="attr-value"}

[Auto]{#attr-popover-auto-state .dfn}

Closes other popovers when opened; has [light
dismiss](#popover-light-dismiss){#the-popover-attribute:popover-light-dismiss}
and responds to [close
requests](interaction.html#close-request){#the-popover-attribute:close-request}.

(the empty string)

[`manual`]{#attr-popover-manual .dfn dfn-for="html-global/popover"
dfn-type="attr-value"}

[Manual]{#attr-popover-manual-state .dfn}

Does not close other popovers; does not [light
dismiss](#popover-light-dismiss){#the-popover-attribute:popover-light-dismiss-2}
or respond to [close
requests](interaction.html#close-request){#the-popover-attribute:close-request-2}.

[`hint`]{#attr-popover-hint .dfn dfn-for="html-global/popover"
dfn-type="attr-value"}

[Hint]{#attr-popover-hint-state .dfn}

Closes other hint popovers when opened, but not other auto popovers; has
[light
dismiss](#popover-light-dismiss){#the-popover-attribute:popover-light-dismiss-3}
and responds to [close
requests](interaction.html#close-request){#the-popover-attribute:close-request-3}.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* is the [No
Popover]{#attr-popover-none-state .dfn} state, and its *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* is the
[Manual](#attr-popover-manual-state){#the-popover-attribute:attr-popover-manual-state}
state.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/popover](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/popover "The popover property of the HTMLElement interface gets and sets an element's popover state via JavaScript ("auto" or "manual"), and can be used for feature detection.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

The [`popover`]{#dom-popover .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-popover-attribute:reflect}
the [popover](#attr-popover){#the-popover-attribute:attr-popover-5}
attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-popover-attribute:limited-to-only-known-values}.

Every [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-2}
has a [popover visibility state]{#popover-visibility-state .dfn
export=""}, initially
[hidden](#popover-hidden-state){#the-popover-attribute:popover-hidden-state},
with these potential values:

- [hidden]{#popover-hidden-state .dfn dfn-for="popover visibility state"
  export=""}

- [showing]{#popover-showing-state .dfn
  dfn-for="popover visibility state" export=""}

Every [`Document`{#the-popover-attribute:document}](dom.html#document)
has a [popover pointerdown target]{#popover-pointerdown-target .dfn},
which is an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-3}
or null, initially null.

Every [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-4}
has a [popover invoker]{#popover-invoker .dfn}, which is an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-5}
or null, initially set to null.

Every [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-6}
has a [popover showing or hiding]{#popover-showing-or-hiding .dfn},
which is a boolean, initially set to false.

Every [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-7}
[popover toggle task tracker]{#popover-toggle-task-tracker .dfn}, which
is a [toggle task
tracker](interaction.html#toggle-task-tracker){#the-popover-attribute:toggle-task-tracker}
or null, initially null.

Every [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-8}
has a [popover close watcher]{#popover-close-watcher .dfn}, which is a
[close
watcher](interaction.html#close-watcher){#the-popover-attribute:close-watcher}
or null, initially null.

Every [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-9}
has an [opened in popover mode]{#opened-in-popover-mode .dfn}, which is
a string or null, initially null.

The following [attribute change
steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#the-popover-attribute:concept-element-attributes-change-ext
x-internal="concept-element-attributes-change-ext"}, given
`element`{.variable}, `localName`{.variable}, `oldValue`{.variable},
`value`{.variable}, and `namespace`{.variable}, are used for all [HTML
elements](infrastructure.html#html-elements){#the-popover-attribute:html-elements-10}:

1.  If `namespace`{.variable} is not null, then return.

2.  If `localName`{.variable} is not
    [`popover`{#the-popover-attribute:attr-popover-6}](#attr-popover),
    then return.

3.  If `element`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state}
    is in the [showing
    state](#popover-showing-state){#the-popover-attribute:popover-showing-state}
    and `oldValue`{.variable} and `value`{.variable} are in different
    [states](#attr-popover){#the-popover-attribute:attr-popover-7}, then
    run the [hide popover
    algorithm](#hide-popover-algorithm){#the-popover-attribute:hide-popover-algorithm}
    given `element`{.variable}, true, true, false, and null.

`element`{.variable}`.`[`showPopover`](#dom-showpopover){#dom-showpopover-dev}`()`
:   Shows the popover `element`{.variable} by adding it to the top
    layer. If `element`{.variable}\'s
    [`popover`{#the-popover-attribute:attr-popover-8}](#attr-popover)
    attribute is in the
    [Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state}
    state, then this will also close all other
    [Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state-2}
    popovers unless they are an ancestor of `element`{.variable}
    according to the [topmost popover
    ancestor](#topmost-popover-ancestor){#the-popover-attribute:topmost-popover-ancestor}
    algorithm.

`element`{.variable}`.`[`hidePopover`](#dom-hidepopover){#dom-hidepopover-dev}`()`
:   Hides the popover `element`{.variable} by removing it from the top
    layer and applying `display: none` to it.

`element`{.variable}`.`[`togglePopover`](#dom-togglepopover){#dom-togglepopover-dev}`()`
:   If the popover `element`{.variable} is not showing, then this method
    shows it. Otherwise, this method hides it. This method returns true
    if the popover is open after calling it, otherwise false.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/showPopover](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/showPopover "The showPopover() method of the HTMLElement interface shows a popover element (i.e. one that has a valid popover attribute) by adding it to the top layer.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

The [`showPopover(``options`{.variable}`)`]{#dom-showpopover .dfn
dfn-for="HTMLElement" dfn-type="method"} method steps are:

1.  Let `invoker`{.variable} be
    `options`{.variable}\[\"[`source`{#the-popover-attribute:dom-showpopoveroptions-source}](dom.html#dom-showpopoveroptions-source)\"\]
    if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-popover-attribute:map-exists
    x-internal="map-exists"}; otherwise, null.

2.  Run [show
    popover](#show-popover){#the-popover-attribute:show-popover} given
    [this](https://webidl.spec.whatwg.org/#this){#the-popover-attribute:this
    x-internal="this"}, true, and `invoker`{.variable}.

To [show popover]{#show-popover .dfn}, given an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-11}
`element`{.variable}, a boolean `throwExceptions`{.variable}, and an
[HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-12}
or null `invoker`{.variable}:

1.  If the result of running [check popover
    validity](#check-popover-validity){#the-popover-attribute:check-popover-validity}
    given `element`{.variable}, false, `throwExceptions`{.variable}, and
    null is false, then return.

2.  Let `document`{.variable} be `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-popover-attribute:node-document
    x-internal="node-document"}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert
    x-internal="assert"}: `element`{.variable}\'s [popover
    invoker](#popover-invoker){#the-popover-attribute:popover-invoker}
    is null.

4.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-2
    x-internal="assert"}: `element`{.variable} is not in
    `document`{.variable}\'s [top
    layer](https://drafts.csswg.org/css-position-4/#document-top-layer){#the-popover-attribute:top-layer
    x-internal="top-layer"}.

5.  Let `nestedShow`{.variable} be `element`{.variable}\'s [popover
    showing or
    hiding](#popover-showing-or-hiding){#the-popover-attribute:popover-showing-or-hiding}.

6.  Let `fireEvents`{.variable} be the boolean negation of
    `nestedShow`{.variable}.

7.  Set `element`{.variable}\'s [popover showing or
    hiding](#popover-showing-or-hiding){#the-popover-attribute:popover-showing-or-hiding-2}
    to true.

8.  Let `cleanupShowingFlag`{.variable} be the following steps:

    1.  If `nestedShow`{.variable} is false, then set
        `element`{.variable}\'s [popover showing or
        hiding](#popover-showing-or-hiding){#the-popover-attribute:popover-showing-or-hiding-3}
        to false.

9.  If the result of [firing an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-popover-attribute:concept-event-fire
    x-internal="concept-event-fire"} named
    [`beforetoggle`{#the-popover-attribute:event-beforetoggle}](indices.html#event-beforetoggle),
    using
    [`ToggleEvent`{#the-popover-attribute:toggleevent}](interaction.html#toggleevent),
    with the
    [`cancelable`{#the-popover-attribute:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    attribute initialized to true, the
    [`oldState`{#the-popover-attribute:dom-toggleevent-oldstate}](interaction.html#dom-toggleevent-oldstate)
    attribute initialized to \"`closed`\", the
    [`newState`{#the-popover-attribute:dom-toggleevent-newstate}](interaction.html#dom-toggleevent-newstate)
    attribute initialized to \"`open`\" at `element`{.variable}, and the
    [`source`{#the-popover-attribute:dom-toggleevent-source}](interaction.html#dom-toggleevent-source)
    attribute initialized to `invoker`{.variable} is false, then run
    `cleanupShowingFlag`{.variable} and return.

10. If the result of running [check popover
    validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-2}
    given `element`{.variable}, false, `throwExceptions`{.variable}, and
    `document`{.variable} is false, then run
    `cleanupShowingFlag`{.variable} and return.

    [Check popover
    validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-3}
    is called again because firing the
    [`beforetoggle`{#the-popover-attribute:event-beforetoggle-2}](indices.html#event-beforetoggle)
    event could have disconnected this element or changed its
    [`popover`{#the-popover-attribute:attr-popover-9}](#attr-popover)
    attribute.

11. Let `shouldRestoreFocus`{.variable} be false.

12. Let `originalType`{.variable} be the current state of
    `element`{.variable}\'s
    [`popover`{#the-popover-attribute:attr-popover-10}](#attr-popover)
    attribute.

13. Let `stackToAppendTo`{.variable} be null.

14. Let `autoAncestor`{.variable} be the result of running the [topmost
    popover
    ancestor](#topmost-popover-ancestor){#the-popover-attribute:topmost-popover-ancestor-2}
    algorithm given `element`{.variable}, `document`{.variable}\'s
    [showing auto popover
    list](#auto-popover-list){#the-popover-attribute:auto-popover-list},
    `invoker`{.variable}, and true.

15. Let `hintAncestor`{.variable} be the result of running the [topmost
    popover
    ancestor](#topmost-popover-ancestor){#the-popover-attribute:topmost-popover-ancestor-3}
    algorithm given `element`{.variable}, `document`{.variable}\'s
    [showing hint popover
    list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list},
    `invoker`{.variable}, and true.

16. If `originalType`{.variable} is the
    [Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state-3}
    state, then:

    1.  Run [close entire popover
        list](#close-entire-popover-list){#the-popover-attribute:close-entire-popover-list}
        given `document`{.variable}\'s [showing hint popover
        list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-2},
        `shouldRestoreFocus`{.variable}, and `fireEvents`{.variable}.

    2.  Let `ancestor`{.variable} be the result of running the [topmost
        popover
        ancestor](#topmost-popover-ancestor){#the-popover-attribute:topmost-popover-ancestor-4}
        algorithm given `element`{.variable}, `document`{.variable}\'s
        [showing auto popover
        list](#auto-popover-list){#the-popover-attribute:auto-popover-list-2},
        `invoker`{.variable}, and true.

    3.  If `ancestor`{.variable} is null, then set `ancestor`{.variable}
        to `document`{.variable}.

    4.  Run [hide all popovers
        until](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until}
        given `ancestor`{.variable}, `shouldRestoreFocus`{.variable},
        and `fireEvents`{.variable}.

    5.  Set `stackToAppendTo`{.variable} to \"`auto`\".

17. If `originalType`{.variable} is the
    [Hint](#attr-popover-hint-state){#the-popover-attribute:attr-popover-hint-state}
    state, then:

    1.  If `hintAncestor`{.variable} is not null, then:

        1.  Run [hide all popovers
            until](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until-2}
            given `hintAncestor`{.variable},
            `shouldRestoreFocus`{.variable}, and
            `fireEvents`{.variable}.

        2.  Set `stackToAppendTo`{.variable} to \"`hint`\".

    2.  Otherwise:

        1.  Run [close entire popover
            list](#close-entire-popover-list){#the-popover-attribute:close-entire-popover-list-2}
            given `document`{.variable}\'s [showing hint popover
            list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-3},
            `shouldRestoreFocus`{.variable}, and
            `fireEvents`{.variable}.

        2.  If `autoAncestor`{.variable} is not null, then:

            1.  Run [hide all popovers
                until](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until-3}
                given `autoAncestor`{.variable},
                `shouldRestoreFocus`{.variable}, and
                `fireEvents`{.variable}.

            2.  Set `stackToAppendTo`{.variable} to \"`auto`\".

        3.  Otherwise, set `stackToAppendTo`{.variable} to \"`hint`\".

18. If `originalType`{.variable} is
    [Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state-4}
    or
    [Hint](#attr-popover-hint-state){#the-popover-attribute:attr-popover-hint-state-2},
    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-3
        x-internal="assert"}: `stackToAppendTo`{.variable} is not null.

    2.  If `originalType`{.variable} is not equal to the value of
        `element`{.variable}\'s
        [`popover`{#the-popover-attribute:attr-popover-11}](#attr-popover)
        attribute, then:

        1.  If `throwExceptions`{.variable} is true, then throw an
            [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-popover-attribute:invalidstateerror
            x-internal="invalidstateerror"}
            [`DOMException`{#the-popover-attribute:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        2.  Return.

    3.  If the result of running [check popover
        validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-4}
        given `element`{.variable}, false, `throwExceptions`{.variable},
        and `document`{.variable} is false, then run
        `cleanupShowingFlag`{.variable} and return.

        [Check popover
        validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-5}
        is called again because running [hide all popovers
        until](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until-4}
        above could have fired the
        [`beforetoggle`{#the-popover-attribute:event-beforetoggle-3}](indices.html#event-beforetoggle)
        event, and an event handler could have disconnected this element
        or changed its
        [`popover`{#the-popover-attribute:attr-popover-12}](#attr-popover)
        attribute.

    4.  If the result of running [topmost auto or hint
        popover](#topmost-auto-popover){#the-popover-attribute:topmost-auto-popover}
        on `document`{.variable} is null, then set
        `shouldRestoreFocus`{.variable} to true.

        This ensures that focus is returned to the previously-focused
        element only for the first popover in a stack.

    5.  If `stackToAppendTo`{.variable} is \"`auto`\":

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-4
            x-internal="assert"}: `document`{.variable}\'s [showing auto
            popover
            list](#auto-popover-list){#the-popover-attribute:auto-popover-list-3}
            does not contain `element`{.variable}.

        2.  Set `element`{.variable}\'s [opened in popover
            mode](#opened-in-popover-mode){#the-popover-attribute:opened-in-popover-mode}
            to \"`auto`\".

        Otherwise:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-5
            x-internal="assert"}: `stackToAppendTo`{.variable} is
            \"`hint`\".

        2.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-6
            x-internal="assert"}: `document`{.variable}\'s [showing hint
            popover
            list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-4}
            does not contain `element`{.variable}.

        3.  Set `element`{.variable}\'s [opened in popover
            mode](#opened-in-popover-mode){#the-popover-attribute:opened-in-popover-mode-2}
            to \"`hint`\".

    6.  ::: {#canceling-popovers}
        Set `element`{.variable}\'s [popover close
        watcher](#popover-close-watcher){#the-popover-attribute:popover-close-watcher}
        to the result of [establishing a close
        watcher](interaction.html#establish-a-close-watcher){#the-popover-attribute:establish-a-close-watcher}
        given `element`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#the-popover-attribute:concept-relevant-global},
        with:

        - *[cancelAction](interaction.html#create-close-watcher-cancelaction)*
          being to return true.

        - *[closeAction](interaction.html#create-close-watcher-closeaction)*
          being to [hide a
          popover](#hide-popover-algorithm){#the-popover-attribute:hide-popover-algorithm-2}
          given `element`{.variable}, true, true, false, and null.

        - *[getEnabledState](interaction.html#create-close-watcher-getenabledstate)*
          being to return true.
        :::

19. Set `element`{.variable}\'s [previously focused
    element](interactive-elements.html#previously-focused-element){#the-popover-attribute:previously-focused-element}
    to null.

20. Let `originallyFocusedElement`{.variable} be
    `document`{.variable}\'s [focused area of the
    document](interaction.html#focused-area-of-the-document){#the-popover-attribute:focused-area-of-the-document}\'s
    [DOM
    anchor](interaction.html#dom-anchor){#the-popover-attribute:dom-anchor}.

21. [Add an element to the top
    layer](https://drafts.csswg.org/css-position-4/#add-an-element-to-the-top-layer){#the-popover-attribute:add-an-element-to-the-top-layer
    x-internal="add-an-element-to-the-top-layer"} given
    `element`{.variable}.

22. Set `element`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-2}
    to
    [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-2}.

23. Set `element`{.variable}\'s [popover
    invoker](#popover-invoker){#the-popover-attribute:popover-invoker-2}
    to `invoker`{.variable}.

24. Set `element`{.variable}\'s [implicit anchor
    element](https://drafts.csswg.org/css-anchor-position/#implicit-anchor-element){#the-popover-attribute:implicit-anchor-element
    x-internal="implicit-anchor-element"} to `invoker`{.variable}.

25. Run the [popover focusing
    steps](#popover-focusing-steps){#the-popover-attribute:popover-focusing-steps}
    given `element`{.variable}.

26. If `shouldRestoreFocus`{.variable} is true and
    `element`{.variable}\'s
    [`popover`{#the-popover-attribute:attr-popover-13}](#attr-popover)
    attribute is not in the [No
    Popover](#attr-popover-none-state){#the-popover-attribute:attr-popover-none-state}
    state, then set `element`{.variable}\'s [previously focused
    element](interactive-elements.html#previously-focused-element){#the-popover-attribute:previously-focused-element-2}
    to `originallyFocusedElement`{.variable}.

27. [Queue a popover toggle event
    task](#queue-a-popover-toggle-event-task){#the-popover-attribute:queue-a-popover-toggle-event-task}
    given `element`{.variable}, \"`closed`\", \"`open`\", and
    `invoker`{.variable}.

28. Run `cleanupShowingFlag`{.variable}.

To [queue a popover toggle event
task]{#queue-a-popover-toggle-event-task .dfn} given an element
`element`{.variable}, a string `oldState`{.variable}, a string
`newState`{.variable}, and an
[`Element`{#the-popover-attribute:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
or null `source`{.variable}:

1.  If `element`{.variable}\'s [popover toggle task
    tracker](#popover-toggle-task-tracker){#the-popover-attribute:popover-toggle-task-tracker}
    is not null, then:

    1.  Set `oldState`{.variable} to `element`{.variable}\'s [popover
        toggle task
        tracker](#popover-toggle-task-tracker){#the-popover-attribute:popover-toggle-task-tracker-2}\'s
        [old
        state](interaction.html#toggle-task-old-state){#the-popover-attribute:toggle-task-old-state}.

    2.  Remove `element`{.variable}\'s [popover toggle task
        tracker](#popover-toggle-task-tracker){#the-popover-attribute:popover-toggle-task-tracker-3}\'s
        [task](interaction.html#toggle-task-task){#the-popover-attribute:toggle-task-task}
        from its [task
        queue](webappapis.html#task-queue){#the-popover-attribute:task-queue}.

    3.  Set `element`{.variable}\'s [popover toggle task
        tracker](#popover-toggle-task-tracker){#the-popover-attribute:popover-toggle-task-tracker-4}
        to null.

2.  [Queue an element
    task](webappapis.html#queue-an-element-task){#the-popover-attribute:queue-an-element-task}
    given the [DOM manipulation task
    source](webappapis.html#dom-manipulation-task-source){#the-popover-attribute:dom-manipulation-task-source}
    and `element`{.variable} to run the following steps:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-popover-attribute:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`toggle`{#the-popover-attribute:event-toggle}](indices.html#event-toggle)
        at `element`{.variable}, using
        [`ToggleEvent`{#the-popover-attribute:toggleevent-2}](interaction.html#toggleevent),
        with the
        [`oldState`{#the-popover-attribute:dom-toggleevent-oldstate-2}](interaction.html#dom-toggleevent-oldstate)
        attribute initialized to `oldState`{.variable}, the
        [`newState`{#the-popover-attribute:dom-toggleevent-newstate-2}](interaction.html#dom-toggleevent-newstate)
        attribute initialized to `newState`{.variable}, and the
        [`source`{#the-popover-attribute:dom-toggleevent-source-2}](interaction.html#dom-toggleevent-source)
        attribute initialized to `source`{.variable}.

    2.  Set `element`{.variable}\'s [popover toggle task
        tracker](#popover-toggle-task-tracker){#the-popover-attribute:popover-toggle-task-tracker-5}
        to null.

3.  Set `element`{.variable}\'s [popover toggle task
    tracker](#popover-toggle-task-tracker){#the-popover-attribute:popover-toggle-task-tracker-6}
    to a struct with
    [task](interaction.html#toggle-task-task){#the-popover-attribute:toggle-task-task-2}
    set to the just-queued
    [task](webappapis.html#concept-task){#the-popover-attribute:concept-task}
    and [old
    state](interaction.html#toggle-task-old-state){#the-popover-attribute:toggle-task-old-state-2}
    set to `oldState`{.variable}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/hidePopover](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/hidePopover "The hidePopover() method of the HTMLElement interface hides a popover element (i.e. one that has a valid popover attribute) by removing it from the top layer and styling it with display: none.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

The [`hidePopover()`]{#dom-hidepopover .dfn dfn-for="HTMLElement"
dfn-type="method"} method steps are:

1.  Run the [hide popover
    algorithm](#hide-popover-algorithm){#the-popover-attribute:hide-popover-algorithm-3}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-popover-attribute:this-2
    x-internal="this"}, true, true, true, and null.

To [hide a popover]{#hide-popover-algorithm .dfn} given an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-13}
`element`{.variable}, a boolean `focusPreviousElement`{.variable}, a
boolean `fireEvents`{.variable}, a boolean `throwExceptions`{.variable},
and an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-14}
or null `source`{.variable}:

1.  If the result of running [check popover
    validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-6}
    given `element`{.variable}, true, `throwExceptions`{.variable}, and
    null is false, then return.

2.  Let `document`{.variable} be `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-popover-attribute:node-document-2
    x-internal="node-document"}.

3.  Let `nestedHide`{.variable} be `element`{.variable}\'s [popover
    showing or
    hiding](#popover-showing-or-hiding){#the-popover-attribute:popover-showing-or-hiding-4}.

4.  Set `element`{.variable}\'s [popover showing or
    hiding](#popover-showing-or-hiding){#the-popover-attribute:popover-showing-or-hiding-5}
    to true.

5.  If `nestedHide`{.variable} is true, then set `fireEvents`{.variable}
    to false.

6.  Let `cleanupSteps`{.variable} be the following steps:

    1.  If `nestedHide`{.variable} is false, then set
        `element`{.variable}\'s [popover showing or
        hiding](#popover-showing-or-hiding){#the-popover-attribute:popover-showing-or-hiding-6}
        to false.

    2.  If `element`{.variable}\'s [popover close
        watcher](#popover-close-watcher){#the-popover-attribute:popover-close-watcher-2}
        is not null, then:

        1.  [Destroy](interaction.html#close-watcher-destroy){#the-popover-attribute:close-watcher-destroy}
            `element`{.variable}\'s [popover close
            watcher](#popover-close-watcher){#the-popover-attribute:popover-close-watcher-3}.

        2.  Set `element`{.variable}\'s [popover close
            watcher](#popover-close-watcher){#the-popover-attribute:popover-close-watcher-4}
            to null.

7.  If `element`{.variable}\'s [opened in popover
    mode](#opened-in-popover-mode){#the-popover-attribute:opened-in-popover-mode-3}
    is \"`auto`\" or \"`hint`\", then:

    1.  Run [hide all popovers
        until](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until-5}
        given `element`{.variable}, `focusPreviousElement`{.variable},
        and `fireEvents`{.variable}.

    2.  If the result of running [check popover
        validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-7}
        given `element`{.variable}, true, and
        `throwExceptions`{.variable} is false, then run
        `cleanupSteps`{.variable} and return.

        [Check popover
        validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-8}
        is called again because running [hide all popovers
        until](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until-6}
        could have disconnected `element`{.variable} or changed its
        [`popover`{#the-popover-attribute:attr-popover-14}](#attr-popover)
        attribute.

8.  Let `autoPopoverListContainsElement`{.variable} be true if
    `document`{.variable}\'s [showing auto popover
    list](#auto-popover-list){#the-popover-attribute:auto-popover-list-4}\'s
    last item is `element`{.variable}, otherwise false.

9.  If `fireEvents`{.variable} is true:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-popover-attribute:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`beforetoggle`{#the-popover-attribute:event-beforetoggle-4}](indices.html#event-beforetoggle),
        using
        [`ToggleEvent`{#the-popover-attribute:toggleevent-3}](interaction.html#toggleevent),
        with the
        [`oldState`{#the-popover-attribute:dom-toggleevent-oldstate-3}](interaction.html#dom-toggleevent-oldstate)
        attribute initialized to \"`open`\", the
        [`newState`{#the-popover-attribute:dom-toggleevent-newstate-3}](interaction.html#dom-toggleevent-newstate)
        attribute initialized to \"`closed`\", and the
        [`source`{#the-popover-attribute:dom-toggleevent-source-3}](interaction.html#dom-toggleevent-source)
        attribute set to `source`{.variable} at `element`{.variable}.

    2.  If `autoPopoverListContainsElement`{.variable} is true and
        `document`{.variable}\'s [showing auto popover
        list](#auto-popover-list){#the-popover-attribute:auto-popover-list-5}\'s
        last item is not `element`{.variable}, then run [hide all
        popovers
        until](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until-7}
        given `element`{.variable}, `focusPreviousElement`{.variable},
        and false.

    3.  If the result of running [check popover
        validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-9}
        given `element`{.variable}, true, `throwExceptions`{.variable},
        and null is false, then run `cleanupSteps`{.variable} and
        return.

        [Check popover
        validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-10}
        is called again because firing the
        [`beforetoggle`{#the-popover-attribute:event-beforetoggle-5}](indices.html#event-beforetoggle)
        event could have disconnected `element`{.variable} or changed
        its
        [`popover`{#the-popover-attribute:attr-popover-15}](#attr-popover)
        attribute.

    4.  [Request an element to be removed from the top
        layer](https://drafts.csswg.org/css-position-4/#request-an-element-to-be-removed-from-the-top-layer){#the-popover-attribute:request-an-element-to-be-removed-from-the-top-layer
        x-internal="request-an-element-to-be-removed-from-the-top-layer"}
        given `element`{.variable}.

    5.  Set `element`{.variable}\'s [implicit anchor
        element](https://drafts.csswg.org/css-anchor-position/#implicit-anchor-element){#the-popover-attribute:implicit-anchor-element-2
        x-internal="implicit-anchor-element"} to null.

10. Otherwise, [remove an element from the top layer
    immediately](https://drafts.csswg.org/css-position-4/#remove-an-element-from-the-top-layer-immediately){#the-popover-attribute:remove-an-element-from-the-top-layer-immediately
    x-internal="remove-an-element-from-the-top-layer-immediately"} given
    `element`{.variable}.

11. Set `element`{.variable}\'s [popover
    invoker](#popover-invoker){#the-popover-attribute:popover-invoker-3}
    to null.

12. Set `element`{.variable}\'s [opened in popover
    mode](#opened-in-popover-mode){#the-popover-attribute:opened-in-popover-mode-4}
    to null.

13. Set `element`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-3}
    to
    [hidden](#popover-hidden-state){#the-popover-attribute:popover-hidden-state-2}.

14. If `fireEvents`{.variable} is true, then [queue a popover toggle
    event
    task](#queue-a-popover-toggle-event-task){#the-popover-attribute:queue-a-popover-toggle-event-task-2}
    given `element`{.variable}, \"`open`\", \"`closed`\", and
    `source`{.variable}.

15. Let `previouslyFocusedElement`{.variable} be `element`{.variable}\'s
    [previously focused
    element](interactive-elements.html#previously-focused-element){#the-popover-attribute:previously-focused-element-3}.

16. If `previouslyFocusedElement`{.variable} is not null, then:

    1.  Set `element`{.variable}\'s [previously focused
        element](interactive-elements.html#previously-focused-element){#the-popover-attribute:previously-focused-element-4}
        to null.

    2.  If `focusPreviousElement`{.variable} is true and
        `document`{.variable}\'s [focused area of the
        document](interaction.html#focused-area-of-the-document){#the-popover-attribute:focused-area-of-the-document-2}\'s
        [DOM
        anchor](interaction.html#dom-anchor){#the-popover-attribute:dom-anchor-2}
        is a [shadow-including inclusive
        descendant](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-descendant){#the-popover-attribute:shadow-including-inclusive-descendant
        x-internal="shadow-including-inclusive-descendant"} of
        `element`{.variable}, then run the [focusing
        steps](interaction.html#focusing-steps){#the-popover-attribute:focusing-steps}
        for `previouslyFocusedElement`{.variable}; the viewport should
        not be scrolled by doing this step.

17. Run `cleanupSteps`{.variable}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/togglePopover](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/togglePopover "The togglePopover() method of the HTMLElement interface toggles a popover element (i.e. one that has a valid popover attribute) between the hidden and showing states.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

The [`togglePopover(``options`{.variable}`)`]{#dom-togglepopover .dfn
dfn-for="HTMLElement" dfn-type="method"} method steps are:

1.  Let `force`{.variable} be null.

2.  If `options`{.variable} is a boolean, set `force`{.variable} to
    `options`{.variable}.

3.  Otherwise, if
    `options`{.variable}\[\"[`force`{#the-popover-attribute:dom-togglepopoveroptions-force}](dom.html#dom-togglepopoveroptions-force)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-popover-attribute:map-exists-2
    x-internal="map-exists"}, set `force`{.variable} to
    `options`{.variable}\[\"[`force`{#the-popover-attribute:dom-togglepopoveroptions-force-2}](dom.html#dom-togglepopoveroptions-force)\"\].

4.  Let `invoker`{.variable} be
    `options`{.variable}\[\"[`source`{#the-popover-attribute:dom-showpopoveroptions-source-2}](dom.html#dom-showpopoveroptions-source)\"\]
    if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-popover-attribute:map-exists-3
    x-internal="map-exists"}; otherwise, null.

5.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-popover-attribute:this-3
    x-internal="this"}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-4}
    is
    [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-3},
    and `force`{.variable} is null or false, then run the [hide popover
    algorithm](#hide-popover-algorithm){#the-popover-attribute:hide-popover-algorithm-4}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-popover-attribute:this-4
    x-internal="this"}, true, true, true, and null.

6.  Otherwise, if `force`{.variable} is null or true, then run [show
    popover](#show-popover){#the-popover-attribute:show-popover-2} given
    [this](https://webidl.spec.whatwg.org/#this){#the-popover-attribute:this-5
    x-internal="this"}, true, and `invoker`{.variable}.

7.  Otherwise:

    1.  Let `expectedToBeShowing`{.variable} be true if
        [this](https://webidl.spec.whatwg.org/#this){#the-popover-attribute:this-6
        x-internal="this"}\'s [popover visibility
        state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-5}
        is
        [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-4};
        otherwise false.

    2.  Run [check popover
        validity](#check-popover-validity){#the-popover-attribute:check-popover-validity-11}
        given `expectedToBeShowing`{.variable}, true, and null.

8.  Return true if
    [this](https://webidl.spec.whatwg.org/#this){#the-popover-attribute:this-7
    x-internal="this"}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-6}
    is
    [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-5};
    otherwise false.

To [hide all popovers until]{#hide-all-popovers-until .dfn export=""},
given an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-15}
or [`Document`{#the-popover-attribute:document-2}](dom.html#document)
`endpoint`{.variable}, a boolean `focusPreviousElement`{.variable}, and
a boolean `fireEvents`{.variable}:

1.  If `endpoint`{.variable} is an [HTML
    element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-16}
    and `endpoint`{.variable} is not in the [popover showing
    state](#popover-showing-state){#the-popover-attribute:popover-showing-state-6},
    then return.

2.  Let `document`{.variable} be `endpoint`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-popover-attribute:node-document-3
    x-internal="node-document"}.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-7
    x-internal="assert"}: `endpoint`{.variable} is a
    [`Document`{#the-popover-attribute:document-3}](dom.html#document)
    or `endpoint`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-7}
    is
    [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-7}.

4.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-8
    x-internal="assert"}: `endpoint`{.variable} is a
    [`Document`{#the-popover-attribute:document-4}](dom.html#document)
    or `endpoint`{.variable}\'s
    [`popover`{#the-popover-attribute:attr-popover-16}](#attr-popover)
    attribute is in the
    [Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state-5}
    state or `endpoint`{.variable}\'s
    [`popover`{#the-popover-attribute:attr-popover-17}](#attr-popover)
    attribute is in the
    [Hint](#attr-popover-hint-state){#the-popover-attribute:attr-popover-hint-state-3}
    state.

5.  If `endpoint`{.variable} is a
    [`Document`{#the-popover-attribute:document-5}](dom.html#document):

    1.  Run [close entire popover
        list](#close-entire-popover-list){#the-popover-attribute:close-entire-popover-list-3}
        given `document`{.variable}\'s [showing hint popover
        list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-5},
        `focusPreviousElement`{.variable}, and `fireEvents`{.variable}.

    2.  Run [close entire popover
        list](#close-entire-popover-list){#the-popover-attribute:close-entire-popover-list-4}
        given `document`{.variable}\'s [showing auto popover
        list](#auto-popover-list){#the-popover-attribute:auto-popover-list-6},
        `focusPreviousElement`{.variable}, and `fireEvents`{.variable}.

    3.  Return.

6.  If `document`{.variable}\'s [showing hint popover
    list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-6}
    contains `endpoint`{.variable}:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-9
        x-internal="assert"}: `endpoint`{.variable}\'s
        [`popover`{#the-popover-attribute:attr-popover-18}](#attr-popover)
        attribute is in the
        [Hint](#attr-popover-hint-state){#the-popover-attribute:attr-popover-hint-state-4}
        state.

    2.  Run [hide popover stack
        until](#hide-popover-stack-until){#the-popover-attribute:hide-popover-stack-until}
        given `endpoint`{.variable}, `document`{.variable}\'s [showing
        hint popover
        list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-7},
        `focusPreviousElement`{.variable}, and `fireEvents`{.variable}.

    3.  Return.

7.  Run [close entire popover
    list](#close-entire-popover-list){#the-popover-attribute:close-entire-popover-list-5}
    given `document`{.variable}\'s [showing hint popover
    list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-8},
    `focusPreviousElement`{.variable}, and `fireEvents`{.variable}.

8.  If `document`{.variable}\'s [showing auto popover
    list](#auto-popover-list){#the-popover-attribute:auto-popover-list-7}
    does not contain `endpoint`{.variable}, then return.

9.  Run [hide popover stack
    until](#hide-popover-stack-until){#the-popover-attribute:hide-popover-stack-until-2}
    given `endpoint`{.variable}, `document`{.variable}\'s [showing auto
    popover
    list](#auto-popover-list){#the-popover-attribute:auto-popover-list-8},
    `focusPreviousElement`{.variable}, and `fireEvents`{.variable}.

To [hide popover stack until]{#hide-popover-stack-until .dfn}, given an
[HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-17}
`endpoint`{.variable}, a
[list](https://infra.spec.whatwg.org/#list){#the-popover-attribute:list
x-internal="list"} `popoverList`{.variable}, a boolean
`focusPreviousElement`{.variable}, and a boolean
`fireEvents`{.variable}:

1.  Let `repeatingHide`{.variable} be false.

2.  Perform the following steps at least once:

    1.  Let `lastToHide`{.variable} be null.

    2.  For each `popover`{.variable} in `popoverList`{.variable}:

        1.  If `popover`{.variable} is `endpoint`{.variable}, then
            [break](https://infra.spec.whatwg.org/#iteration-break){#the-popover-attribute:break
            x-internal="break"}.

        2.  Set `lastToHide`{.variable} to `popover`{.variable}.

    3.  If `lastToHide`{.variable} is null, then return.

    4.  [While](https://infra.spec.whatwg.org/#iteration-while){#the-popover-attribute:while
        x-internal="while"} `lastToHide`{.variable}\'s [popover
        visibility
        state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-8}
        is
        [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-8}:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-10
            x-internal="assert"}: `popoverList`{.variable} is not empty.

        2.  Run the [hide popover
            algorithm](#hide-popover-algorithm){#the-popover-attribute:hide-popover-algorithm-5}
            given the last item in `popoverList`{.variable},
            `focusPreviousElement`{.variable}, `fireEvents`{.variable},
            false, and null.

    5.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-11
        x-internal="assert"}: `repeatingHide`{.variable} is false or
        `popoverList`{.variable}\'s last item is `endpoint`{.variable}.

    6.  Set `repeatingHide`{.variable} to true if
        `popoverList`{.variable} contains `endpoint`{.variable} and
        `popoverList`{.variable}\'s last item is not
        `endpoint`{.variable}, otherwise false.

    7.  If `repeatingHide`{.variable} is true, then set
        `fireEvents`{.variable} to false.

    and keep performing them
    [while](https://infra.spec.whatwg.org/#iteration-while){#the-popover-attribute:while-2
    x-internal="while"} `repeatingHide`{.variable} is true.

The [hide all popovers until
algorithm](#hide-all-popovers-until){#the-popover-attribute:hide-all-popovers-until-8}
is used in several cases to hide all popovers that don\'t stay open when
something happens. For example, during light-dismiss of a popover, this
algorithm ensures that we close only the popovers that aren\'t related
to the node clicked by the user.

To find the [topmost popover ancestor]{#topmost-popover-ancestor .dfn
export=""}, given a
[`Node`{#the-popover-attribute:node}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
`newPopoverOrTopLayerElement`{.variable}, a
[list](https://infra.spec.whatwg.org/#list){#the-popover-attribute:list-2
x-internal="list"} `popoverList`{.variable}, an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-18}
or null `invoker`{.variable}, and a boolean `isPopover`{.variable},
perform the following steps. They return an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-19}
or null.

::: note
The [topmost popover
ancestor](#topmost-popover-ancestor){#the-popover-attribute:topmost-popover-ancestor-5}
algorithm will return the topmost (latest in the [showing auto popover
list](#auto-popover-list){#the-popover-attribute:auto-popover-list-9})
ancestor popover for the provided popover or top layer element. Popovers
can be related to each other in several ways, creating a tree of
popovers. There are two paths through which one popover (call it the
\"child\" popover) can have a topmost ancestor popover (call it the
\"parent\" popover):

1.  The popovers are nested within each other in the node tree. In this
    case, the descendant popover is the \"child\" and its topmost
    ancestor popover is the \"parent\".

2.  An invoking element (e.g., a
    [`button`{#the-popover-attribute:the-button-element}](form-elements.html#the-button-element))
    has a
    [`popovertarget`{#the-popover-attribute:attr-popovertarget}](#attr-popovertarget)
    attribute pointing to a popover. In this case, the popover is the
    \"child\", and the popover subtree the invoking element is in is the
    \"parent\". The invoker has to be in a popover and reference an open
    popover.

In each of the relationships formed above, the parent popover has to be
strictly earlier in the [showing auto popover
list](#auto-popover-list){#the-popover-attribute:auto-popover-list-10}
than the child popover, or it does not form a valid ancestral
relationship. This eliminates non-showing popovers and self-pointers
(e.g., a popover containing an invoking element that points back to the
containing popover), and it allows for the construction of a well-formed
tree from the (possibly cyclic) graph of connections. Only
[Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state-6}
popovers are considered.

If the provided element is a top layer element such as a
[`dialog`{#the-popover-attribute:the-dialog-element}](interactive-elements.html#the-dialog-element)
which is not showing as a popover, then [topmost popover
ancestor](#topmost-popover-ancestor){#the-popover-attribute:topmost-popover-ancestor-6}
will only look in the node tree to find the first popover.
:::

1.  If `isPopover`{.variable} is true:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-12
        x-internal="assert"}: `newPopoverOrTopLayerElement`{.variable}
        is an [HTML
        element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-20}.

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-13
        x-internal="assert"}:
        `newPopoverOrTopLayerElement`{.variable}\'s
        [`popover`{#the-popover-attribute:attr-popover-19}](#attr-popover)
        attribute is not in the [No Popover
        State](#attr-popover-none-state){#the-popover-attribute:attr-popover-none-state-2}
        or the
        [Manual](#attr-popover-manual-state){#the-popover-attribute:attr-popover-manual-state-2}
        state.

    3.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-14
        x-internal="assert"}:
        `newPopoverOrTopLayerElement`{.variable}\'s [popover visibility
        state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-9}
        is not in the [popover showing
        state](#popover-showing-state){#the-popover-attribute:popover-showing-state-9}.

2.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-15
        x-internal="assert"}: `invoker`{.variable} is null.

3.  Let `popoverPositions`{.variable} be an empty [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#the-popover-attribute:ordered-map
    x-internal="ordered-map"}.

4.  Let `index`{.variable} be 0.

5.  For each `popover`{.variable} of `popoverList`{.variable}:

    1.  [Set](https://infra.spec.whatwg.org/#map-set){#the-popover-attribute:map-set
        x-internal="map-set"}
        `popoverPositions`{.variable}\[`popover`{.variable}\] to
        `index`{.variable}.

    2.  Increment `index`{.variable} by 1.

6.  If `isPopover`{.variable} is true, then
    [set](https://infra.spec.whatwg.org/#map-set){#the-popover-attribute:map-set-2
    x-internal="map-set"}
    `popoverPositions`{.variable}\[`newPopoverOrTopLayerElement`{.variable}\]
    to `index`{.variable}.

7.  Increment `index`{.variable} by 1.

8.  Let `topmostPopoverAncestor`{.variable} be null.

9.  Let `checkAncestor`{.variable} be an algorithm which performs the
    following steps given `candidate`{.variable}:

    1.  If `candidate`{.variable} is null, then return.

    2.  Let `okNesting`{.variable} be false.

    3.  Let `candidateAncestor`{.variable} be null.

    4.  [While](https://infra.spec.whatwg.org/#iteration-while){#the-popover-attribute:while-3
        x-internal="while"} `okNesting`{.variable} is false:

        1.  Set `candidateAncestor`{.variable} to the result of running
            [nearest inclusive open
            popover](#nearest-inclusive-open-popover){#the-popover-attribute:nearest-inclusive-open-popover}
            given `candidate`{.variable}.

        2.  If `candidateAncestor`{.variable} is null or
            `popoverPositions`{.variable} does not contain
            `candidateAncestor`{.variable}, then return.

        3.  [Assert](https://infra.spec.whatwg.org/#assert){#the-popover-attribute:assert-16
            x-internal="assert"}: `candidateAncestor`{.variable}\'s
            [`popover`{#the-popover-attribute:attr-popover-20}](#attr-popover)
            attribute is not in the
            [Manual](#attr-popover-manual-state){#the-popover-attribute:attr-popover-manual-state-3}
            or
            [None](#attr-popover-none-state){#the-popover-attribute:attr-popover-none-state-3}
            state.

        4.  Set `okNesting`{.variable} to true if
            `newPopoverOrTopLayerElement`{.variable}\'s
            [`popover`{#the-popover-attribute:attr-popover-21}](#attr-popover)
            attribute is in the
            [Hint](#attr-popover-hint-state){#the-popover-attribute:attr-popover-hint-state-5}
            state or `candidateAncestor`{.variable}\'s
            [`popover`{#the-popover-attribute:attr-popover-22}](#attr-popover)
            attribute is in the
            [Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state-7}
            state.

        5.  If `okNesting`{.variable} is false, then set
            `candidate`{.variable} to `candidateAncestor`{.variable}\'s
            parent in the [flat
            tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-popover-attribute:flat-tree
            x-internal="flat-tree"}.

    5.  Let `candidatePosition`{.variable} be
        `popoverPositions`{.variable}\[`candidateAncestor`{.variable}\].

    6.  If `topmostPopoverAncestor`{.variable} is null or
        `popoverPositions`{.variable}\[`topmostPopoverAncestor`{.variable}\]
        is less than `candidatePosition`{.variable}, then set
        `topmostPopoverAncestor`{.variable} to
        `candidateAncestor`{.variable}.

10. Run `checkAncestor`{.variable} given
    `newPopoverOrTopLayerElement`{.variable}\'s parent node within the
    [flat
    tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-popover-attribute:flat-tree-2
    x-internal="flat-tree"}.

11. Run `checkAncestor`{.variable} given `invoker`{.variable}.

12. Return `topmostPopoverAncestor`{.variable}.

To find the [nearest inclusive open
popover]{#nearest-inclusive-open-popover .dfn} given a
[`Node`{#the-popover-attribute:node-2}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
`node`{.variable}, perform the following steps. They return an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-21}
or null.

1.  Let `currentNode`{.variable} be `node`{.variable}.

2.  While `currentNode`{.variable} is not null:

    1.  If `currentNode`{.variable}\'s
        [`popover`{#the-popover-attribute:attr-popover-23}](#attr-popover)
        attribute is in the
        [Auto](#attr-popover-auto-state){#the-popover-attribute:attr-popover-auto-state-8}
        state and `currentNode`{.variable}\'s [popover visibility
        state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-10}
        is
        [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-10},
        then return `currentNode`{.variable}.

    2.  Set `currentNode`{.variable} to `currentNode`{.variable}\'s
        parent in the [flat
        tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-popover-attribute:flat-tree-3
        x-internal="flat-tree"}.

3.  Return null.

To [find the topmost auto or hint popover]{#topmost-auto-popover .dfn}
given a
[`Document`{#the-popover-attribute:document-6}](dom.html#document)
`document`{.variable}, perform the following steps. They return an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-22}
or null.

1.  If `document`{.variable}\'s [showing hint popover
    list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-9}
    is not empty, then return `document`{.variable}\'s [showing hint
    popover
    list](#showing-hint-popover-list){#the-popover-attribute:showing-hint-popover-list-10}\'s
    last element.

2.  If `document`{.variable}\'s [showing auto popover
    list](#auto-popover-list){#the-popover-attribute:auto-popover-list-11}
    is not empty, then return `document`{.variable}\'s [showing auto
    popover
    list](#auto-popover-list){#the-popover-attribute:auto-popover-list-12}\'s
    last element.

3.  Return null.

To perform the [popover focusing steps]{#popover-focusing-steps .dfn}
for an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-23}
`subject`{.variable}:

1.  If the [allow focus
    steps](interaction.html#allow-focus-steps){#the-popover-attribute:allow-focus-steps}
    given `subject`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-popover-attribute:node-document-4
    x-internal="node-document"} return false, then return.

2.  If `subject`{.variable} is a
    [`dialog`{#the-popover-attribute:the-dialog-element-2}](interactive-elements.html#the-dialog-element)
    element, then run the [dialog focusing
    steps](interactive-elements.html#dialog-focusing-steps){#the-popover-attribute:dialog-focusing-steps}
    given `subject`{.variable} and return.

3.  If `subject`{.variable} has the
    [`autofocus`{#the-popover-attribute:attr-fe-autofocus}](interaction.html#attr-fe-autofocus)
    attribute, then let `control`{.variable} be `subject`{.variable}.

4.  Otherwise, let `control`{.variable} be the [autofocus
    delegate](interaction.html#autofocus-delegate){#the-popover-attribute:autofocus-delegate}
    for `subject`{.variable} given \"`other`\".

5.  If `control`{.variable} is null, then return.

6.  Run the [focusing
    steps](interaction.html#focusing-steps){#the-popover-attribute:focusing-steps-2}
    given `control`{.variable}.

7.  Let `topDocument`{.variable} be `control`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-popover-attribute:node-navigable}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#the-popover-attribute:nav-top}\'s
    [active
    document](document-sequences.html#nav-document){#the-popover-attribute:nav-document}.

8.  If `control`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-popover-attribute:node-document-5
    x-internal="node-document"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-popover-attribute:concept-document-origin
    x-internal="concept-document-origin"} is not the
    [same](browsers.html#same-origin){#the-popover-attribute:same-origin}
    as the
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-popover-attribute:concept-document-origin-2
    x-internal="concept-document-origin"} of `topDocument`{.variable},
    then return.

9.  [Empty](https://infra.spec.whatwg.org/#list-empty){#the-popover-attribute:list-empty
    x-internal="list-empty"} `topDocument`{.variable}\'s [autofocus
    candidates](interaction.html#autofocus-candidates){#the-popover-attribute:autofocus-candidates}.

10. Set `topDocument`{.variable}\'s [autofocus processed
    flag](interaction.html#autofocus-processed-flag){#the-popover-attribute:autofocus-processed-flag}
    to true.

To [check popover validity]{#check-popover-validity .dfn} for an [HTML
element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-24}
`element`{.variable} given a boolean `expectedToBeShowing`{.variable}, a
boolean `throwExceptions`{.variable}, and a
[`Document`{#the-popover-attribute:document-7}](dom.html#document) or
null `expectedDocument`{.variable}, perform the following steps. They
throw an exception or return a boolean.

1.  If `element`{.variable}\'s
    [`popover`{#the-popover-attribute:attr-popover-24}](#attr-popover)
    attribute is in the [No
    Popover](#attr-popover-none-state){#the-popover-attribute:attr-popover-none-state-4}
    state, then:

    1.  If `throwExceptions`{.variable} is true, then throw a
        [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-popover-attribute:notsupportederror
        x-internal="notsupportederror"}
        [`DOMException`{#the-popover-attribute:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  Return false.

2.  If any of the following are true:

    - `expectedToBeShowing`{.variable} is true and
      `element`{.variable}\'s [popover visibility
      state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-11}
      is not
      [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-11};
      or

    - `expectedToBeShowing`{.variable} is false and
      `element`{.variable}\'s [popover visibility
      state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-12}
      is not
      [hidden](#popover-hidden-state){#the-popover-attribute:popover-hidden-state-3},

    then return false.

3.  If any of the following are true:

    - `element`{.variable} is not
      [connected](https://dom.spec.whatwg.org/#connected){#the-popover-attribute:connected
      x-internal="connected"};

    - `element`{.variable}\'s [node
      document](https://dom.spec.whatwg.org/#concept-node-document){#the-popover-attribute:node-document-6
      x-internal="node-document"} is not [fully
      active](document-sequences.html#fully-active){#the-popover-attribute:fully-active};

    - `expectedDocument`{.variable} is not null and
      `element`{.variable}\'s [node
      document](https://dom.spec.whatwg.org/#concept-node-document){#the-popover-attribute:node-document-7
      x-internal="node-document"} is not `expectedDocument`{.variable};

    - `element`{.variable} is a
      [`dialog`{#the-popover-attribute:the-dialog-element-3}](interactive-elements.html#the-dialog-element)
      element and its [is
      modal](interactive-elements.html#is-modal){#the-popover-attribute:is-modal}
      is set to true; or

    - `element`{.variable}\'s [fullscreen
      flag](https://fullscreen.spec.whatwg.org/#fullscreen-flag){#the-popover-attribute:fullscreen-flag
      x-internal="fullscreen-flag"} is set,

    then:

    1.  If `throwExceptions`{.variable} is true, then throw an
        [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-popover-attribute:invalidstateerror-2
        x-internal="invalidstateerror"}
        [`DOMException`{#the-popover-attribute:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    2.  Return false.

4.  Return true.

To get the [showing auto popover list]{#auto-popover-list .dfn} for a
[`Document`{#the-popover-attribute:document-8}](dom.html#document)
`document`{.variable}:

1.  Let `popovers`{.variable} be « ».

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-popover-attribute:list-iterate
    x-internal="list-iterate"}
    [`Element`{#the-popover-attribute:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    `element`{.variable} in `document`{.variable}\'s [top
    layer](https://drafts.csswg.org/css-position-4/#document-top-layer){#the-popover-attribute:top-layer-2
    x-internal="top-layer"}:

    1.  If all of the following are true:

        - `element`{.variable} is an [HTML
          element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-25};

        - `element`{.variable}\'s [opened in popover
          mode](#opened-in-popover-mode){#the-popover-attribute:opened-in-popover-mode-5}
          is \"`auto`\"; and

        - `element`{.variable}\'s [popover visibility
          state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-13}
          is
          [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-12},

        then
        [append](https://infra.spec.whatwg.org/#list-append){#the-popover-attribute:list-append
        x-internal="list-append"} `element`{.variable} to
        `popovers`{.variable}.

3.  Return `popovers`{.variable}.

To get the [showing hint popover list]{#showing-hint-popover-list .dfn}
for a [`Document`{#the-popover-attribute:document-9}](dom.html#document)
`document`{.variable}:

1.  Let `popovers`{.variable} be « ».

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-popover-attribute:list-iterate-2
    x-internal="list-iterate"}
    [`Element`{#the-popover-attribute:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    `element`{.variable} in `document`{.variable}\'s [top
    layer](https://drafts.csswg.org/css-position-4/#document-top-layer){#the-popover-attribute:top-layer-3
    x-internal="top-layer"}:

    1.  If all of the following are true:

        - `element`{.variable} is an [HTML
          element](infrastructure.html#html-elements){#the-popover-attribute:html-elements-26};

        - `element`{.variable}\'s [opened in popover
          mode](#opened-in-popover-mode){#the-popover-attribute:opened-in-popover-mode-6}
          is \"`hint`\"; and

        - `element`{.variable}\'s [popover visibility
          state](#popover-visibility-state){#the-popover-attribute:popover-visibility-state-14}
          is
          [showing](#popover-showing-state){#the-popover-attribute:popover-showing-state-13},

        then
        [append](https://infra.spec.whatwg.org/#list-append){#the-popover-attribute:list-append-2
        x-internal="list-append"} `element`{.variable} to
        `popovers`{.variable}.

3.  Return `popovers`{.variable}.

To [close entire popover list]{#close-entire-popover-list .dfn} given a
[list](https://infra.spec.whatwg.org/#list){#the-popover-attribute:list-3
x-internal="list"} `popoverList`{.variable}, a boolean
`focusPreviousElement`{.variable}, and a boolean
`fireEvents`{.variable}:

1.  While `popoverList`{.variable} is not empty:

    1.  Run the [hide popover
        algorithm](#hide-popover-algorithm){#the-popover-attribute:hide-popover-algorithm-6}
        given `popoverList`{.variable}\'s last item,
        `focusPreviousElement`{.variable}, `fireEvents`{.variable},
        false, and null.

#### [6.12.1]{.secno} The popover target attributes[](#the-popover-target-attributes){.self-link}

[Buttons](forms.html#concept-button){#the-popover-target-attributes:concept-button}
may have the following content attributes:

- [`popovertarget`]{#attr-popovertarget .dfn dfn-for="html-global"
  dfn-type="element-attr"}

- [`popovertargetaction`]{#attr-popovertargetaction .dfn
  dfn-for="html-global" dfn-type="element-attr"}

If specified, the
[`popovertarget`{#the-popover-target-attributes:attr-popovertarget}](#attr-popovertarget)
attribute value must be the
[ID](https://dom.spec.whatwg.org/#concept-id){#the-popover-target-attributes:concept-id
x-internal="concept-id"} of an element with a
[`popover`{#the-popover-target-attributes:attr-popover}](#attr-popover)
attribute in the same
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-popover-target-attributes:tree
x-internal="tree"} as the
[button](forms.html#concept-button){#the-popover-target-attributes:concept-button-2}
with the
[`popovertarget`{#the-popover-target-attributes:attr-popovertarget-2}](#attr-popovertarget)
attribute.

The
[`popovertargetaction`{#the-popover-target-attributes:attr-popovertargetaction}](#attr-popovertargetaction)
attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-popover-target-attributes:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`toggle`]{#attr-popovertargetaction-toggle .dfn
dfn-for="html-global/popovertargetaction" dfn-type="attr-value"}

[Toggle]{#attr-popovertargetaction-toggle-state .dfn}

Shows or hides the targeted popover element.

[`show`]{#attr-popovertargetaction-show .dfn
dfn-for="html-global/popovertargetaction" dfn-type="attr-value"}

[Show]{#attr-popovertargetaction-show-state .dfn}

Shows the targeted popover element.

[`hide`]{#attr-popovertargetaction-hide .dfn
dfn-for="html-global/popovertargetaction" dfn-type="attr-value"}

[Hide]{#attr-popovertargetaction-hide-state .dfn}

Hides the targeted popover element.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[Toggle](#attr-popovertargetaction-toggle-state){#the-popover-target-attributes:attr-popovertargetaction-toggle-state}
state.

Whenever possible ensure the popover element is placed immediately after
its triggering element in the DOM. Doing so will help ensure that the
popover is exposed in a logical programmatic reading order for users of
assistive technology, such as screen readers.

::: example
The following shows how the
[`popovertarget`{#the-popover-target-attributes:attr-popovertarget-3}](#attr-popovertarget)
attribute in combination with the
[`popovertargetaction`{#the-popover-target-attributes:attr-popovertargetaction-2}](#attr-popovertargetaction)
attribute can be used to show and close a popover:

``` html
<button popovertarget="foo" popovertargetaction="show">
  Show a popover
</button>

<article popover="auto" id="foo">
  This is a popover article!
  <button popovertarget="foo" popovertargetaction="hide">Close</button>
</article>
```

If a
[`popovertargetaction`{#the-popover-target-attributes:attr-popovertargetaction-3}](#attr-popovertargetaction)
attribute is not specified, the default action will be to toggle the
associated popover. The following shows how only specifying the
[`popovertarget`{#the-popover-target-attributes:attr-popovertarget-4}](#attr-popovertarget)
attribute on its invoking button can toggle a manual popover between its
opened and closed states. A manual popover will not respond to [light
dismiss](#popover-light-dismiss){#the-popover-target-attributes:popover-light-dismiss}
or [close
requests](interaction.html#close-request){#the-popover-target-attributes:close-request}:

``` html
<input type="button" popovertarget="foo" value="Toggle the popover">

<div popover=manual id="foo">
  This is a popover!
</div>
```
:::

[DOM
interface](dom.html#concept-element-dom){#the-popover-target-attributes:concept-element-dom}:

``` idl
interface mixin PopoverInvokerElement {
  [CEReactions] attribute Element? popoverTargetElement;
  [CEReactions] attribute DOMString popoverTargetAction;
};
```

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLButtonElement/popoverTargetElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement/popoverTargetElement "The popoverTargetElement property of the HTMLButtonElement interface gets and sets the popover element to control via a button.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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
[HTMLInputElement/popoverTargetElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/popoverTargetElement "The popoverTargetElement property of the HTMLInputElement interface gets and sets the popover element to control via an <input> element of type="button".")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

The [`popoverTargetElement`]{#dom-popovertargetelement .dfn
dfn-for="HTMLElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-popover-target-attributes:reflect}
the
[`popovertarget`{#the-popover-target-attributes:attr-popovertarget-5}](#attr-popovertarget)
attribute.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLButtonElement/popoverTargetAction](https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement/popoverTargetAction "The popoverTargetAction property of the HTMLButtonElement interface gets and sets the action to be performed ("hide", "show", or "toggle") on a popover element being controlled by a button.")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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
[HTMLInputElement/popoverTargetAction](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/popoverTargetAction "The popoverTargetAction property of the HTMLInputElement interface gets and sets the action to be performed ("hide", "show", or "toggle") on a popover element being controlled by an <input> element of type="button".")

Support in all current engines.

::: support
[Firefox[🔰
114+]{title="Requires setting a user preference or runtime flag."}]{.firefox
.yes}[Safari17+]{.safari .yes}[Chrome114+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge114+]{.edge_blink .yes}

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

The [`popoverTargetAction`]{#dom-popovertargetaction .dfn
dfn-for="HTMLElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-popover-target-attributes:reflect-2}
the
[`popovertargetaction`{#the-popover-target-attributes:attr-popovertargetaction-4}](#attr-popovertargetaction)
attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-popover-target-attributes:limited-to-only-known-values}.

To run the [popover target attribute activation
behavior]{#popover-target-attribute-activation-behavior .dfn} given a
[`Node`{#the-popover-target-attributes:node}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
`node`{.variable} and a
[`Node`{#the-popover-target-attributes:node-2}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
`eventTarget`{.variable}:

1.  Let `popover`{.variable} be `node`{.variable}\'s [popover target
    element](#popover-target-element){#the-popover-target-attributes:popover-target-element}.

2.  If `popover`{.variable} is null, then return.

3.  If `eventTarget`{.variable} is a [shadow-including inclusive
    descendant](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-descendant){#the-popover-target-attributes:shadow-including-inclusive-descendant
    x-internal="shadow-including-inclusive-descendant"} of
    `popover`{.variable} and `popover`{.variable} is a [shadow-including
    descendant](https://dom.spec.whatwg.org/#concept-shadow-including-descendant){#the-popover-target-attributes:shadow-including-descendant
    x-internal="shadow-including-descendant"} of `node`{.variable}, then
    return.

4.  If `node`{.variable}\'s
    [`popovertargetaction`{#the-popover-target-attributes:attr-popovertargetaction-5}](#attr-popovertargetaction)
    attribute is in the
    [show](#attr-popovertargetaction-show){#the-popover-target-attributes:attr-popovertargetaction-show}
    state and `popover`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-target-attributes:popover-visibility-state}
    is
    [showing](#popover-showing-state){#the-popover-target-attributes:popover-showing-state},
    then return.

5.  If `node`{.variable}\'s
    [`popovertargetaction`{#the-popover-target-attributes:attr-popovertargetaction-6}](#attr-popovertargetaction)
    attribute is in the
    [hide](#attr-popovertargetaction-hide){#the-popover-target-attributes:attr-popovertargetaction-hide}
    state and `popover`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-target-attributes:popover-visibility-state-2}
    is
    [hidden](#popover-hidden-state){#the-popover-target-attributes:popover-hidden-state},
    then return.

6.  If `popover`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-target-attributes:popover-visibility-state-3}
    is
    [showing](#popover-showing-state){#the-popover-target-attributes:popover-showing-state-2},
    then run the [hide popover
    algorithm](#hide-popover-algorithm){#the-popover-target-attributes:hide-popover-algorithm}
    given `popover`{.variable}, true, true, false, and
    `node`{.variable}.

7.  Otherwise, if `popover`{.variable}\'s [popover visibility
    state](#popover-visibility-state){#the-popover-target-attributes:popover-visibility-state-4}
    is
    [hidden](#popover-hidden-state){#the-popover-target-attributes:popover-hidden-state-2}
    and the result of running [check popover
    validity](#check-popover-validity){#the-popover-target-attributes:check-popover-validity}
    given `popover`{.variable}, false, false, and null is true, then run
    [show
    popover](#show-popover){#the-popover-target-attributes:show-popover}
    given `popover`{.variable}, false, and `node`{.variable}.

To get the [popover target element]{#popover-target-element .dfn} given
a
[`Node`{#the-popover-target-attributes:node-3}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
`node`{.variable}, perform the following steps. They return an [HTML
element](infrastructure.html#html-elements){#the-popover-target-attributes:html-elements}
or null.

1.  If `node`{.variable} is not a
    [button](forms.html#concept-button){#the-popover-target-attributes:concept-button-3},
    then return null.

2.  If `node`{.variable} is
    [disabled](form-control-infrastructure.html#concept-fe-disabled){#the-popover-target-attributes:concept-fe-disabled},
    then return null.

3.  If `node`{.variable} has a [form
    owner](form-control-infrastructure.html#form-owner){#the-popover-target-attributes:form-owner}
    and `node`{.variable} is a [submit
    button](forms.html#concept-submit-button){#the-popover-target-attributes:concept-submit-button},
    then return null.

4.  Let `popoverElement`{.variable} be the result of running
    `node`{.variable}\'s [get the `popovertarget`-associated
    element](common-dom-interfaces.html#attr-associated-element){#the-popover-target-attributes:attr-associated-element}.

5.  If `popoverElement`{.variable} is null, then return null.

6.  If `popoverElement`{.variable}\'s
    [`popover`{#the-popover-target-attributes:attr-popover-2}](#attr-popover)
    attribute is in the [No
    Popover](#attr-popover-none-state){#the-popover-target-attributes:attr-popover-none-state}
    state, then return null.

7.  Return `popoverElement`{.variable}.

#### [6.12.2]{.secno} [Popover light dismiss]{.dfn}[](#popover-light-dismiss){.self-link}

\"Light dismiss\" means that clicking outside of a popover whose
[`popover`{#popover-light-dismiss:attr-popover}](#attr-popover)
attribute is in the
[Auto](#attr-popover-auto-state){#popover-light-dismiss:attr-popover-auto-state}
state will close the popover. This is in addition to how such popovers
respond to [close
requests](interaction.html#close-request){#popover-light-dismiss:close-request}.

To [light dismiss open popovers]{#light-dismiss-open-popovers .dfn},
given a
[`PointerEvent`{#popover-light-dismiss:pointerevent}](https://w3c.github.io/pointerevents/#pointerevent-interface){x-internal="pointerevent"}
`event`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#popover-light-dismiss:assert
    x-internal="assert"}: `event`{.variable}\'s
    [`isTrusted`{#popover-light-dismiss:dom-event-istrusted}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
    attribute is true.

2.  Let `target`{.variable} be `event`{.variable}\'s
    [target](https://dom.spec.whatwg.org/#concept-event-target){#popover-light-dismiss:concept-event-target
    x-internal="concept-event-target"}.

3.  Let `document`{.variable} be `target`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#popover-light-dismiss:node-document
    x-internal="node-document"}.

4.  Let `topmostPopover`{.variable} be the result of running [topmost
    auto
    popover](#topmost-auto-popover){#popover-light-dismiss:topmost-auto-popover}
    given `document`{.variable}.

5.  If `topmostPopover`{.variable} is null, then return.

6.  If `event`{.variable}\'s
    [`type`{#popover-light-dismiss:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    is
    \"[`pointerdown`{#popover-light-dismiss:event-pointerdown}](https://w3c.github.io/pointerevents/#the-pointerdown-event){x-internal="event-pointerdown"}\",
    then: set `document`{.variable}\'s [popover pointerdown
    target](#popover-pointerdown-target){#popover-light-dismiss:popover-pointerdown-target}
    to the result of running [topmost clicked
    popover](#topmost-clicked-popover){#popover-light-dismiss:topmost-clicked-popover}
    given `target`{.variable}.

7.  If `event`{.variable}\'s
    [`type`{#popover-light-dismiss:dom-event-type-2}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    is
    \"[`pointerup`{#popover-light-dismiss:event-pointerup}](https://w3c.github.io/pointerevents/#the-pointerup-event){x-internal="event-pointerup"}\",
    then:

    1.  Let `ancestor`{.variable} be the result of running [topmost
        clicked
        popover](#topmost-clicked-popover){#popover-light-dismiss:topmost-clicked-popover-2}
        given `target`{.variable}.

    2.  Let `sameTarget`{.variable} be true if `ancestor`{.variable} is
        `document`{.variable}\'s [popover pointerdown
        target](#popover-pointerdown-target){#popover-light-dismiss:popover-pointerdown-target-2}.

    3.  Set `document`{.variable}\'s [popover pointerdown
        target](#popover-pointerdown-target){#popover-light-dismiss:popover-pointerdown-target-3}
        to null.

    4.  If `ancestor`{.variable} is null, then set `ancestor`{.variable}
        to `document`{.variable}.

    5.  If `sameTarget`{.variable} is true, then run [hide all popovers
        until](#hide-all-popovers-until){#popover-light-dismiss:hide-all-popovers-until}
        given `ancestor`{.variable}, false, and true.

To find the [topmost clicked popover]{#topmost-clicked-popover .dfn},
given a
[`Node`{#popover-light-dismiss:node}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
`node`{.variable}:

1.  Let `clickedPopover`{.variable} be the result of running [nearest
    inclusive open
    popover](#nearest-inclusive-open-popover){#popover-light-dismiss:nearest-inclusive-open-popover}
    given `node`{.variable}.

2.  Let `invokerPopover`{.variable} be the result of running [nearest
    inclusive target popover for
    invoker](#nearest-inclusive-target-popover-for-invoker){#popover-light-dismiss:nearest-inclusive-target-popover-for-invoker}
    given `node`{.variable}.

3.  If the result of [getting the popover stack
    position](#get-the-popover-stack-position){#popover-light-dismiss:get-the-popover-stack-position}
    given `clickedPopover`{.variable} is greater than the result of
    [getting the popover stack
    position](#get-the-popover-stack-position){#popover-light-dismiss:get-the-popover-stack-position-2}
    given `invokerPopover`{.variable}, then return
    `clickedPopover`{.variable}.

4.  Return `invokerPopover`{.variable}.

To [get the popover stack position]{#get-the-popover-stack-position
.dfn}, given an [HTML
element](infrastructure.html#html-elements){#popover-light-dismiss:html-elements}
`popover`{.variable}:

1.  Let `hintList`{.variable} be `popover`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#popover-light-dismiss:node-document-2
    x-internal="node-document"}\'s [showing hint popover
    list](#showing-hint-popover-list){#popover-light-dismiss:showing-hint-popover-list}.

2.  Let `autoList`{.variable} be `popover`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#popover-light-dismiss:node-document-3
    x-internal="node-document"}\'s [showing auto popover
    list](#auto-popover-list){#popover-light-dismiss:auto-popover-list}.

3.  If `popover`{.variable} is in `hintList`{.variable}, then return the
    index of `popover`{.variable} in `hintList`{.variable} + the size of
    `autoList`{.variable} + 1.

4.  If `popover`{.variable} is in `autoList`{.variable}, then return the
    index of `popover`{.variable} in `autoList`{.variable} + 1.

5.  Return 0.

To find the [nearest inclusive target popover for
invoker]{#nearest-inclusive-target-popover-for-invoker .dfn} given a
[`Node`{#popover-light-dismiss:node-2}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
`node`{.variable}:

1.  Let `currentNode`{.variable} be `node`{.variable}.

2.  While `currentNode`{.variable} is not null:

    1.  Let `targetPopover`{.variable} be `currentNode`{.variable}\'s
        [popover target
        element](#popover-target-element){#popover-light-dismiss:popover-target-element}.

    2.  If `targetPopover`{.variable} is not null and
        `targetPopover`{.variable}\'s
        [`popover`{#popover-light-dismiss:attr-popover-2}](#attr-popover)
        attribute is in the
        [Auto](#attr-popover-auto-state){#popover-light-dismiss:attr-popover-auto-state-2}
        state and `targetPopover`{.variable}\'s [popover visibility
        state](#popover-visibility-state){#popover-light-dismiss:popover-visibility-state}
        is
        [showing](#popover-showing-state){#popover-light-dismiss:popover-showing-state},
        then return `targetPopover`{.variable}.

    3.  Set `currentNode`{.variable} to `currentNode`{.variable}\'s
        ancestor in the [flat
        tree](https://drafts.csswg.org/css-scoping/#flat-tree){#popover-light-dismiss:flat-tree
        x-internal="flat-tree"}.

[← 6.11 Drag and drop](dnd.html) --- [Table of Contents](./) --- [7
Loading web pages →](browsers.html)
