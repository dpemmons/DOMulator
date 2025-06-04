HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 5 Microdata](microdata.html) --- [Table of Contents](./) --- [6.11
Drag and drop →](dnd.html)

1.  [[[6]{.secno} User
    interaction](interaction.html#editing)]{#toc-editing}
    1.  [[6.1]{.secno} The `hidden`
        attribute](interaction.html#the-hidden-attribute)
    2.  [[6.2]{.secno} Page
        visibility](interaction.html#page-visibility)
        1.  [[6.2.1]{.secno} The `VisibilityStateEntry`
            interface](interaction.html#the-visibilitystateentry-interface)
    3.  [[6.3]{.secno} Inert subtrees](interaction.html#inert-subtrees)
        1.  [[6.3.1]{.secno} Modal dialogs and inert
            subtrees](interaction.html#modal-dialogs-and-inert-subtrees)
        2.  [[6.3.2]{.secno} The `inert`
            attribute](interaction.html#the-inert-attribute)
    4.  [[6.4]{.secno} Tracking user
        activation](interaction.html#tracking-user-activation)
        1.  [[6.4.1]{.secno} Data
            model](interaction.html#user-activation-data-model)
        2.  [[6.4.2]{.secno} Processing
            model](interaction.html#user-activation-processing-model)
        3.  [[6.4.3]{.secno} APIs gated by user
            activation](interaction.html#user-activation-gated-apis)
        4.  [[6.4.4]{.secno} The `UserActivation`
            interface](interaction.html#the-useractivation-interface)
        5.  [[6.4.5]{.secno} User agent
            automation](interaction.html#user-activation-user-agent-automation)
    5.  [[6.5]{.secno} Activation behavior of
        elements](interaction.html#activation)
        1.  [[6.5.1]{.secno} The `ToggleEvent`
            interface](interaction.html#the-toggleevent-interface)
        2.  [[6.5.2]{.secno} The `CommandEvent`
            interface](interaction.html#the-commandevent-interface)
    6.  [[6.6]{.secno} Focus](interaction.html#focus)
        1.  [[6.6.1]{.secno}
            Introduction](interaction.html#introduction-8)
        2.  [[6.6.2]{.secno} Data model](interaction.html#data-model)
        3.  [[6.6.3]{.secno} The `tabindex`
            attribute](interaction.html#the-tabindex-attribute)
        4.  [[6.6.4]{.secno} Processing
            model](interaction.html#focus-processing-model)
        5.  [[6.6.5]{.secno} Sequential focus
            navigation](interaction.html#sequential-focus-navigation)
        6.  [[6.6.6]{.secno} Focus management
            APIs](interaction.html#focus-management-apis)
        7.  [[6.6.7]{.secno} The `autofocus`
            attribute](interaction.html#the-autofocus-attribute)
    7.  [[6.7]{.secno} Assigning keyboard
        shortcuts](interaction.html#assigning-keyboard-shortcuts)
        1.  [[6.7.1]{.secno}
            Introduction](interaction.html#introduction-9)
        2.  [[6.7.2]{.secno} The `accesskey`
            attribute](interaction.html#the-accesskey-attribute)
        3.  [[6.7.3]{.secno} Processing
            model](interaction.html#keyboard-shortcuts-processing-model)
    8.  [[6.8]{.secno} Editing](interaction.html#editing-2)
        1.  [[6.8.1]{.secno} Making document regions editable: The
            `contenteditable` content
            attribute](interaction.html#contenteditable)
        2.  [[6.8.2]{.secno} Making entire documents editable: the
            `designMode` getter and
            setter](interaction.html#making-entire-documents-editable:-the-designmode-idl-attribute)
        3.  [[6.8.3]{.secno} Best practices for in-page
            editors](interaction.html#best-practices-for-in-page-editors)
        4.  [[6.8.4]{.secno} Editing
            APIs](interaction.html#editing-apis)
        5.  [[6.8.5]{.secno} Spelling and grammar
            checking](interaction.html#spelling-and-grammar-checking)
        6.  [[6.8.6]{.secno} Writing
            suggestions](interaction.html#writing-suggestions)
        7.  [[6.8.7]{.secno}
            Autocapitalization](interaction.html#autocapitalization)
        8.  [[6.8.8]{.secno}
            Autocorrection](interaction.html#autocorrection)
        9.  [[6.8.9]{.secno} Input modalities: the `inputmode`
            attribute](interaction.html#input-modalities:-the-inputmode-attribute)
        10. [[6.8.10]{.secno} Input modalities: the `enterkeyhint`
            attribute](interaction.html#input-modalities:-the-enterkeyhint-attribute)
    9.  [[6.9]{.secno} Find-in-page](interaction.html#find-in-page)
        1.  [[6.9.1]{.secno}
            Introduction](interaction.html#introduction-10)
        2.  [[6.9.2]{.secno} Interaction with `details` and
            `hidden=until-found`](interaction.html#interaction-with-details-and-hidden=until-found)
        3.  [[6.9.3]{.secno} Interaction with
            selection](interaction.html#interaction-with-selection)
    10. [[6.10]{.secno} Close requests and close
        watchers](interaction.html#close-requests-and-close-watchers)
        1.  [[6.10.1]{.secno} Close
            requests](interaction.html#close-requests)
        2.  [[6.10.2]{.secno} Close watcher
            infrastructure](interaction.html#close-watcher-infrastructure)
        3.  [[6.10.3]{.secno} The `CloseWatcher`
            interface](interaction.html#the-closewatcher-interface)

## [6]{.secno} [User interaction]{.dfn}[](#editing){.self-link} {#editing}

### [6.1]{.secno} The [`hidden`{#the-hidden-attribute:attr-hidden}](#attr-hidden) attribute[](#the-hidden-attribute){.self-link}

::::::: {.mdn-anno .wrapped}
**⚠**MDN

:::: feature
[Global_attributes/hidden](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/hidden "The hidden global attribute is an enumerated attribute indicating that the browser should not render the contents of the element. For example, it can be used to hide elements of the page that can't be used until the login process has been completed.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome102+]{.chrome
.yes}

------------------------------------------------------------------------

[OperaNo]{.opera .no}[Edge102+]{.edge_blink .yes}

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
[Global_attributes/hidden](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/hidden "The hidden global attribute is an enumerated attribute indicating that the browser should not render the contents of the element. For example, it can be used to hide elements of the page that can't be used until the login process has been completed.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

All [HTML
elements](infrastructure.html#html-elements){#the-hidden-attribute:html-elements}
may have the [`hidden`]{#attr-hidden .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute set. The
[`hidden`{#the-hidden-attribute:attr-hidden-2}](#attr-hidden) attribute
is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-hidden-attribute:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`hidden`]{#attr-hidden-hidden .dfn dfn-for="html-global/hidden"
dfn-type="attr-value"}

[Hidden]{#attr-hidden-hidden-state .dfn}

Will not be rendered.

(the empty string)

[`until-found`]{#attr-hidden-until-found .dfn
dfn-for="html-global/hidden" dfn-type="attr-value"}

[Hidden Until Found]{#attr-hidden-until-found-state .dfn}

Will not be rendered, but content inside will be accessible to
[find-in-page](#find-in-page-2){#the-hidden-attribute:find-in-page-2}
and [fragment
navigation](browsing-the-web.html#navigate-fragid){#the-hidden-attribute:navigate-fragid}.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* is the [Not
Hidden]{#attr-hidden-not-hidden-state .dfn} state, and its *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* is the
[Hidden](#attr-hidden-hidden-state){#the-hidden-attribute:attr-hidden-hidden-state}
state.

When an element has the
[`hidden`{#the-hidden-attribute:attr-hidden-3}](#attr-hidden) attribute
in the
[Hidden](#attr-hidden-hidden-state){#the-hidden-attribute:attr-hidden-hidden-state-2}
state, it indicates that the element is not yet, or is no longer,
directly relevant to the page\'s current state, or that it is being used
to declare content to be reused by other parts of the page as opposed to
being directly accessed by the user. User agents should not render
elements that are in the
[Hidden](#attr-hidden-hidden-state){#the-hidden-attribute:attr-hidden-hidden-state-3}
state.

The requirement for user agents not to render elements that are in the
[Hidden](#attr-hidden-hidden-state){#the-hidden-attribute:attr-hidden-hidden-state-4}
state can be implemented indirectly through the style layer. For
example, a web browser could implement these requirements [using the
rules suggested in the Rendering section](rendering.html#hiddenCSS).

When an element has the
[`hidden`{#the-hidden-attribute:attr-hidden-4}](#attr-hidden) attribute
in the [Hidden Until
Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state}
state, it indicates that the element is hidden like the
[Hidden](#attr-hidden-hidden-state){#the-hidden-attribute:attr-hidden-hidden-state-5}
state but the content inside the element will be accessible to
[find-in-page](#find-in-page-2){#the-hidden-attribute:find-in-page-2-2}
and [fragment
navigation](browsing-the-web.html#navigate-fragid){#the-hidden-attribute:navigate-fragid-2}.
When these features attempt to scroll to a target which is in the
element\'s subtree, the user agent will remove the
[`hidden`{#the-hidden-attribute:attr-hidden-5}](#attr-hidden) attribute
in order to reveal the content before scrolling to it by running the
[ancestor hidden-until-found revealing
algorithm](#ancestor-hidden-until-found-revealing-algorithm){#the-hidden-attribute:ancestor-hidden-until-found-revealing-algorithm}
on the target node.

Web browsers will use \'content-visibility: hidden\' instead of
\'display: none\' when the
[`hidden`{#the-hidden-attribute:attr-hidden-6}](#attr-hidden) attribute
is in the [Hidden Until
Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-2}
state, as specified in the [Rendering
section](rendering.html#hiddenCSS).

::: note
Because this attribute is typically implemented using CSS, it\'s also
possible to override it using CSS. For instance, a rule that applies
\'display: block\' to all elements will cancel the effects of the
[Hidden](#attr-hidden-hidden-state){#the-hidden-attribute:attr-hidden-hidden-state-6}
state. Authors therefore have to take care when writing their style
sheets to make sure that the attribute is still styled as expected. In
addition, legacy user agents which don\'t support the [Hidden Until
Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-3}
state will have \'display: none\' instead of \'content-visibility:
hidden\', so authors are encouraged to make sure that their style sheets
don\'t change the \'display\' or \'content-visibility\' properties of
[Hidden Until
Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-4}
elements.

Since elements with the
[`hidden`{#the-hidden-attribute:attr-hidden-7}](#attr-hidden) attribute
in the [Hidden Until
Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-5}
state use \'content-visibility: hidden\' instead of \'display: none\',
there are two caveats of the [Hidden Until
Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-6}
state that make it different from the
[Hidden](#attr-hidden-hidden-state){#the-hidden-attribute:attr-hidden-hidden-state-7}
state:

1.  The element needs to be affected by [layout
    containment](https://drafts.csswg.org/css-contain/#containment-layout){#the-hidden-attribute:layout-containment
    x-internal="layout-containment"} in order to be revealed by
    find-in-page. This means that if the element in the [Hidden Until
    Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-7}
    state has a \'display\' value of \'none\', \'contents\', or
    \'inline\', then the element will not be revealed by find-in-page.

2.  The element will still have a [generated
    box](https://drafts.csswg.org/css2/#propdef-visibility){#the-hidden-attribute:'visibility'
    x-internal="'visibility'"} when in the [Hidden Until
    Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-8}
    state, which means that borders, margin, and padding will still be
    rendered around the element.
:::

::: example
In the following skeletal example, the attribute is used to hide the web
game\'s main screen until the user logs in:

``` html
  <h1>The Example Game</h1>
  <section id="login">
   <h2>Login</h2>
   <form>
    ...
    <!-- calls login() once the user's credentials have been checked -->
   </form>
   <script>
    function login() {
      // switch screens
      document.getElementById('login').hidden = true;
      document.getElementById('game').hidden = false;
    }
   </script>
  </section>
  <section id="game" hidden>
   ...
  </section>
```
:::

The [`hidden`{#the-hidden-attribute:attr-hidden-8}](#attr-hidden)
attribute must not be used to hide content that could legitimately be
shown in another presentation. For example, it is incorrect to use
[`hidden`{#the-hidden-attribute:attr-hidden-9}](#attr-hidden) to hide
panels in a tabbed dialog, because the tabbed interface is merely a kind
of overflow presentation --- one could equally well just show all the
form controls in one big page with a scrollbar. It is similarly
incorrect to use this attribute to hide content just from one
presentation --- if something is marked
[`hidden`{#the-hidden-attribute:attr-hidden-10}](#attr-hidden), it is
hidden from all presentations, including, for instance, screen readers.

Elements that are not themselves
[`hidden`{#the-hidden-attribute:attr-hidden-11}](#attr-hidden) must not
[hyperlink](links.html#hyperlink){#the-hidden-attribute:hyperlink} to
elements that are
[`hidden`{#the-hidden-attribute:attr-hidden-12}](#attr-hidden). The
`for` attributes of
[`label`{#the-hidden-attribute:the-label-element}](forms.html#the-label-element)
and
[`output`{#the-hidden-attribute:the-output-element}](form-elements.html#the-output-element)
elements that are not themselves
[`hidden`{#the-hidden-attribute:attr-hidden-13}](#attr-hidden) must
similarly not refer to elements that are
[`hidden`{#the-hidden-attribute:attr-hidden-14}](#attr-hidden). In both
cases, such references would cause user confusion.

Elements and scripts may, however, refer to elements that are
[`hidden`{#the-hidden-attribute:attr-hidden-15}](#attr-hidden) in other
contexts.

::: example
For example, it would be incorrect to use the
[`href`{#the-hidden-attribute:attr-hyperlink-href}](links.html#attr-hyperlink-href)
attribute to link to a section marked with the
[`hidden`{#the-hidden-attribute:attr-hidden-16}](#attr-hidden)
attribute. If the content is not applicable or relevant, then there is
no reason to link to it.

It would be fine, however, to use the ARIA
[`aria-describedby`{#the-hidden-attribute:attr-aria-describedby}](https://w3c.github.io/aria/#aria-describedby){x-internal="attr-aria-describedby"}
attribute to refer to descriptions that are themselves
[`hidden`{#the-hidden-attribute:attr-hidden-17}](#attr-hidden). While
hiding the descriptions implies that they are not useful alone, they
could be written in such a way that they are useful in the specific
context of being referenced from the elements that they describe.

Similarly, a
[`canvas`{#the-hidden-attribute:the-canvas-element}](canvas.html#the-canvas-element)
element with the
[`hidden`{#the-hidden-attribute:attr-hidden-18}](#attr-hidden) attribute
could be used by a scripted graphics engine as an off-screen buffer, and
a form control could refer to a hidden
[`form`{#the-hidden-attribute:the-form-element}](forms.html#the-form-element)
element using its
[`form`{#the-hidden-attribute:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
attribute.
:::

Elements in a section hidden by the
[`hidden`{#the-hidden-attribute:attr-hidden-19}](#attr-hidden) attribute
are still active, e.g. scripts and form controls in such sections still
execute and submit respectively. Only their presentation to the user
changes.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/hidden](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/hidden "The HTMLElement property hidden reflects the value of the element's hidden attribute.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

The [`hidden`]{#dom-hidden .dfn dfn-for="HTMLElement"
dfn-type="attribute"} getter steps are:

1.  If the
    [`hidden`{#the-hidden-attribute:attr-hidden-20}](#attr-hidden)
    attribute is in the [Hidden Until
    Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-9}
    state, then return
    \"[`until-found`{#the-hidden-attribute:attr-hidden-until-found}](#attr-hidden-until-found)\".

2.  If the
    [`hidden`{#the-hidden-attribute:attr-hidden-21}](#attr-hidden)
    attribute is set, then return true.

3.  Return false.

The [`hidden`{#the-hidden-attribute:dom-hidden}](#dom-hidden) setter
steps are:

1.  If the given value is a string that is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-hidden-attribute:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for
    \"[`until-found`{#the-hidden-attribute:attr-hidden-until-found-2}](#attr-hidden-until-found)\",
    then set the
    [`hidden`{#the-hidden-attribute:attr-hidden-22}](#attr-hidden)
    attribute to
    \"[`until-found`{#the-hidden-attribute:attr-hidden-until-found-3}](#attr-hidden-until-found)\".

2.  Otherwise, if the given value is false, then remove the
    [`hidden`{#the-hidden-attribute:attr-hidden-23}](#attr-hidden)
    attribute.

3.  Otherwise, if the given value is the empty string, then remove the
    [`hidden`{#the-hidden-attribute:attr-hidden-24}](#attr-hidden)
    attribute.

4.  Otherwise, if the given value is null, then remove the
    [`hidden`{#the-hidden-attribute:attr-hidden-25}](#attr-hidden)
    attribute.

5.  Otherwise, if the given value is 0, then remove the
    [`hidden`{#the-hidden-attribute:attr-hidden-26}](#attr-hidden)
    attribute.

6.  Otherwise, if the given value is NaN, then remove the
    [`hidden`{#the-hidden-attribute:attr-hidden-27}](#attr-hidden)
    attribute.

7.  Otherwise, set the
    [`hidden`{#the-hidden-attribute:attr-hidden-28}](#attr-hidden)
    attribute to the empty string.

The [ancestor hidden-until-found revealing
algorithm]{#ancestor-hidden-until-found-revealing-algorithm .dfn} is to
run the following steps on `currentNode`{.variable}:

1.  While `currentNode`{.variable} has a parent node within the [flat
    tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-hidden-attribute:flat-tree
    x-internal="flat-tree"}:

    1.  If `currentNode`{.variable} has the
        [`hidden`{#the-hidden-attribute:attr-hidden-29}](#attr-hidden)
        attribute in the [Hidden Until
        Found](#attr-hidden-until-found-state){#the-hidden-attribute:attr-hidden-until-found-state-10}
        state, then:

        1.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#the-hidden-attribute:concept-event-fire
            x-internal="concept-event-fire"} named
            [`beforematch`{#the-hidden-attribute:event-beforematch}](indices.html#event-beforematch)
            at `currentNode`{.variable} with the
            [`bubbles`{#the-hidden-attribute:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
            attribute initialized to true.

        2.  Remove the
            [`hidden`{#the-hidden-attribute:attr-hidden-30}](#attr-hidden)
            attribute from `currentNode`{.variable}.

    2.  Set `currentNode`{.variable} to the parent node of
        `currentNode`{.variable} within the [flat
        tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-hidden-attribute:flat-tree-2
        x-internal="flat-tree"}.

### [6.2]{.secno} Page visibility[](#page-visibility){.self-link}

A [traversable
navigable](document-sequences.html#traversable-navigable){#page-visibility:traversable-navigable}\'s
[system visibility
state](document-sequences.html#system-visibility-state){#page-visibility:system-visibility-state},
including its initial value upon creation, is determined by the user
agent. It represents, for example, whether the browser window is
minimized, a browser tab is currently in the background, or a system
element such as a task switcher obscures the page.

When a user-agent determines that the [system visibility
state](document-sequences.html#system-visibility-state){#page-visibility:system-visibility-state-2}
for [traversable
navigable](document-sequences.html#traversable-navigable){#page-visibility:traversable-navigable-2}
`traversable`{.variable} has changed to `newState`{.variable}, it must
run the following steps:

1.  Let `navigables`{.variable} be the [inclusive descendant
    navigables](document-sequences.html#inclusive-descendant-navigables){#page-visibility:inclusive-descendant-navigables}
    of `traversable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#page-visibility:nav-document}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#page-visibility:list-iterate
    x-internal="list-iterate"} `navigable`{.variable} of
    `navigables`{.variable} [in what order?]{.XXX}:

    1.  Let `document`{.variable} be `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#page-visibility:nav-document-2}.

    2.  [Queue a global
        task](webappapis.html#queue-a-global-task){#page-visibility:queue-a-global-task}
        on the [user interaction task
        source](webappapis.html#user-interaction-task-source){#page-visibility:user-interaction-task-source}
        given `document`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#page-visibility:concept-relevant-global}
        to [update the visibility
        state](#update-the-visibility-state){#page-visibility:update-the-visibility-state}
        of `document`{.variable} with `newState`{.variable}.

A [`Document`{#page-visibility:document}](dom.html#document) has a
[visibility state]{#visibility-state .dfn dfn-for="Document" export=""},
which is either \"`hidden`\" or \"`visible`\", initially set to
\"`hidden`\".

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/visibilityState](https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState "The Document.visibilityState read-only property returns the visibility of the document, that is in which context this element is now visible. It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.")

Support in all current engines.

::: support
[Firefox18+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera20+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4.3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android20+]{.opera_android .yes}
:::
::::
:::::

The [`visibilityState`]{#dom-document-visibilitystate .dfn
dfn-for="Document" dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#page-visibility:this
x-internal="this"}\'s [visibility
state](#visibility-state){#page-visibility:visibility-state}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/hidden](https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden "The Document.hidden read-only property returns a Boolean value indicating if the page is considered hidden or not.")

Support in all current engines.

::: support
[Firefox18+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome33+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4.3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`hidden`]{#dom-document-hidden .dfn dfn-for="Document"
dfn-type="attribute"} getter steps are to return true if
[this](https://webidl.spec.whatwg.org/#this){#page-visibility:this-2
x-internal="this"}\'s [visibility
state](#visibility-state){#page-visibility:visibility-state-2} is
\"`hidden`\", otherwise false.

To [update the visibility state]{#update-the-visibility-state .dfn} of
[`Document`{#page-visibility:document-2}](dom.html#document)
`document`{.variable} to `visibilityState`{.variable}:

1.  If `document`{.variable}\'s [visibility
    state](#visibility-state){#page-visibility:visibility-state-3}
    equals `visibilityState`{.variable}, then return.

2.  Set `document`{.variable}\'s [visibility
    state](#visibility-state){#page-visibility:visibility-state-4} to
    `visibilityState`{.variable}.

3.  [Queue](https://w3c.github.io/performance-timeline/#queue-a-performanceentry){#page-visibility:queue-a-performance-entry
    x-internal="queue-a-performance-entry"} a new
    [`VisibilityStateEntry`{#page-visibility:visibilitystateentry}](#visibilitystateentry)
    whose [visibility
    state](#visibilitystateentry-state){#page-visibility:visibilitystateentry-state}
    is `visibilityState`{.variable} and whose
    [timestamp](#visibilitystateentry-timestamp){#page-visibility:visibilitystateentry-timestamp}
    is the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#page-visibility:current-high-resolution-time
    x-internal="current-high-resolution-time"} given
    `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#page-visibility:concept-relevant-global-2}.

4.  Run the [screen orientation change
    steps](https://w3c.github.io/screen-orientation/#dfn-screen-orientation-change-steps){#page-visibility:screen-orientation-change-steps
    x-internal="screen-orientation-change-steps"} with
    `document`{.variable}.
    [\[SCREENORIENTATION\]](references.html#refsSCREENORIENTATION)

5.  Run the [view transition page visibility change
    steps](https://drafts.csswg.org/css-view-transitions/#view-transition-page-visibility-change-steps){#page-visibility:view-transition-page-visibility-change-steps
    x-internal="view-transition-page-visibility-change-steps"} with
    `document`{.variable}.

6.  Run any [page visibility change steps]{#page-visibility-change-steps
    .dfn export=""} which may be defined in other specifications, with
    [visibility
    state](#visibility-state){#page-visibility:visibility-state-5} and
    `document`{.variable}.

    It would be better if specification authors sent a pull request to
    add calls from here into their specifications directly, instead of
    using the [page visibility change
    steps](#page-visibility-change-steps){#page-visibility:page-visibility-change-steps}
    hook, to ensure well-defined cross-specification call order. As of
    the time of this writing the following specifications are known to
    have [page visibility change
    steps](#page-visibility-change-steps){#page-visibility:page-visibility-change-steps-2},
    which will be run in an unspecified order: Device Posture API and
    Web NFC. [\[DEVICEPOSTURE\]](references.html#refsDEVICEPOSTURE)
    [\[WEBNFC\]](references.html#refsWEBNFC)

7.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#page-visibility:concept-event-fire
    x-internal="concept-event-fire"} named
    [`visibilitychange`{#page-visibility:event-visibilitychange}](indices.html#event-visibilitychange)
    at `document`{.variable}, with its
    [`bubbles`{#page-visibility:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    attribute initialized to true.

#### [6.2.1]{.secno} The [`VisibilityStateEntry`{#the-visibilitystateentry-interface:visibilitystateentry}](#visibilitystateentry) interface[](#the-visibilitystateentry-interface){.self-link}

::::: {.mdn-anno .wrapped}
**⚠**MDN

:::: feature
[VisibilityStateEntry](https://developer.mozilla.org/en-US/docs/Web/API/VisibilityStateEntry "The VisibilityStateEntry interface provides timings of page visibility state changes, i.e., when a tab changes from the foreground to the background or vice versa.")

Support in one engine only.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome115+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge115+]{.edge_blink .yes}

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

The
[`VisibilityStateEntry`{#the-visibilitystateentry-interface:visibilitystateentry-2}](#visibilitystateentry)
interface exposes visibility changes to the document, from the moment
the document becomes active.

::: example
For example, this allows JavaScript code in the page to examine
correlation between visibility changes and paint timing:

``` js
function wasHiddenBeforeFirstContentfulPaint() {
    const fcpEntry = performance.getEntriesByName("first-contentful-paint")[0];
    const visibilityStateEntries = performance.getEntriesByType("visibility-state");
    return visibilityStateEntries.some(e =>
                                            e.startTime < fcpEntry.startTime &&
                                            e.name === "hidden");
}
```
:::

Since hiding a page can cause throttling of rendering and other
user-agent operations, it is common to use visibility changes as an
indication that such throttling has occurred. However, other things
could also cause throttling in different browsers, such as long periods
of inactivity.

``` idl
[Exposed=(Window)]
interface VisibilityStateEntry : PerformanceEntry {
  readonly attribute DOMString name;                 // shadows inherited name
  readonly attribute DOMString entryType;            // shadows inherited entryType
  readonly attribute DOMHighResTimeStamp startTime;  // shadows inherited startTime
  readonly attribute unsigned long duration;         // shadows inherited duration
};
```

The
[`VisibilityStateEntry`{#the-visibilitystateentry-interface:visibilitystateentry-3}](#visibilitystateentry)
has an associated
[`DOMHighResTimeStamp`](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){#the-visibilitystateentry-interface:domhighrestimestamp
x-internal="domhighrestimestamp"}
[timestamp]{#visibilitystateentry-timestamp .dfn}.

The
[`VisibilityStateEntry`{#the-visibilitystateentry-interface:visibilitystateentry-4}](#visibilitystateentry)
has an associated \"`visible`\" or \"`hidden`\" [visibility
state]{#visibilitystateentry-state .dfn}.

The [`name`]{#visibilitystateentry-name .dfn
dfn-for="VisibilityStateEntry" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-visibilitystateentry-interface:this
x-internal="this"}\'s [visibility
state](#visibilitystateentry-state){#the-visibilitystateentry-interface:visibilitystateentry-state}.

The [`entryType`]{#visibilitystateentry-entrytype .dfn
dfn-for="VisibilityStateEntry" dfn-type="attribute"} getter steps are to
return \"`visibility-state`\".

The [`startTime`]{#visibilitystateentry-starttime .dfn
dfn-for="VisibilityStateEntry" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-visibilitystateentry-interface:this-2
x-internal="this"}\'s
[timestamp](#visibilitystateentry-timestamp){#the-visibilitystateentry-interface:visibilitystateentry-timestamp}.

The [`duration`]{#visibilitystateentry-duration .dfn
dfn-for="VisibilityStateEntry" dfn-type="attribute"} getter steps are to
return zero.

### [6.3]{.secno} Inert subtrees[](#inert-subtrees){.self-link}

See also
[`inert`{#inert-subtrees:the-inert-attribute}](#the-inert-attribute) for
an explanation of the attribute of the same name.

A node (in particular elements and text nodes) can be [inert]{#inert
.dfn}. When a node is [inert](#inert){#inert-subtrees:inert}:

- Hit-testing must act as if the
  [\'pointer-events\'](https://drafts.csswg.org/css-ui-4/#pointer-events-control){#inert-subtrees:'pointer-events'
  x-internal="'pointer-events'"} CSS property were set to \'none\'.

- Text selection functionality must act as if the
  [\'user-select\'](https://drafts.csswg.org/css-ui-4/#content-selection){#inert-subtrees:'user-select'
  x-internal="'user-select'"} CSS property were set to \'none\'.

- If it is
  [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#inert-subtrees:editable
  x-internal="editable"}, the node behaves as if it were non-editable.

- The user agent should ignore the node for the purposes of
  [find-in-page](#find-in-page-2){#inert-subtrees:find-in-page-2}.

Inert nodes generally cannot be focused, and user agents do not expose
the inert nodes to accessibility APIs or assistive technologies. Inert
nodes that are
[commands](interactive-elements.html#concept-command){#inert-subtrees:concept-command}
will become inoperable to users, in the manner described above.

User agents may allow the user to override the restrictions on
[find-in-page](#find-in-page-2){#inert-subtrees:find-in-page-2-2} and
text selection, however.

By default, a node is not [inert](#inert){#inert-subtrees:inert-2}.

#### [6.3.1]{.secno} Modal dialogs and inert subtrees[](#modal-dialogs-and-inert-subtrees){.self-link}

A
[`Document`{#modal-dialogs-and-inert-subtrees:document}](dom.html#document)
`document`{.variable} is [blocked by a modal
dialog]{#blocked-by-a-modal-dialog .dfn} `subject`{.variable} if
`subject`{.variable} is the topmost
[`dialog`{#modal-dialogs-and-inert-subtrees:the-dialog-element}](interactive-elements.html#the-dialog-element)
element in `document`{.variable}\'s [top
layer](https://drafts.csswg.org/css-position-4/#document-top-layer){#modal-dialogs-and-inert-subtrees:top-layer
x-internal="top-layer"}. While `document`{.variable} is so blocked,
every node that is
[connected](https://dom.spec.whatwg.org/#connected){#modal-dialogs-and-inert-subtrees:connected
x-internal="connected"} to `document`{.variable}, with the exception of
the `subject`{.variable} element and its [flat
tree](https://drafts.csswg.org/css-scoping/#flat-tree){#modal-dialogs-and-inert-subtrees:flat-tree
x-internal="flat-tree"} descendants, must become
[inert](#inert){#modal-dialogs-and-inert-subtrees:inert}.

`subject`{.variable} can additionally become
[inert](#inert){#modal-dialogs-and-inert-subtrees:inert-2} via the
[`inert`{#modal-dialogs-and-inert-subtrees:the-inert-attribute}](#the-inert-attribute)
attribute, but only if specified on `subject`{.variable} itself (i.e.,
`subject`{.variable} escapes inertness of ancestors);
`subject`{.variable}\'s [flat
tree](https://drafts.csswg.org/css-scoping/#flat-tree){#modal-dialogs-and-inert-subtrees:flat-tree-2
x-internal="flat-tree"} descendants can become
[inert](#inert){#modal-dialogs-and-inert-subtrees:inert-3} in a similar
fashion.

The
[`dialog`{#modal-dialogs-and-inert-subtrees:the-dialog-element-2}](interactive-elements.html#the-dialog-element)
element\'s
[`showModal()`{#modal-dialogs-and-inert-subtrees:dom-dialog-showmodal}](interactive-elements.html#dom-dialog-showmodal)
method causes this mechanism to trigger, by
[adding](https://drafts.csswg.org/css-position-4/#add-an-element-to-the-top-layer){#modal-dialogs-and-inert-subtrees:add-an-element-to-the-top-layer
x-internal="add-an-element-to-the-top-layer"} the
[`dialog`{#modal-dialogs-and-inert-subtrees:the-dialog-element-3}](interactive-elements.html#the-dialog-element)
element to its [node
document](https://dom.spec.whatwg.org/#concept-node-document){#modal-dialogs-and-inert-subtrees:node-document
x-internal="node-document"}\'s [top
layer](https://drafts.csswg.org/css-position-4/#document-top-layer){#modal-dialogs-and-inert-subtrees:top-layer-2
x-internal="top-layer"}.

#### [6.3.2]{.secno} The [`inert`]{.dfn} attribute[](#the-inert-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/inert](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inert "The inert global attribute is a Boolean attribute indicating that the browser will ignore the element. With the inert attribute, all of the element's flat tree descendants (such as modal <dialog>s) that don't otherwise escape inertness are ignored. The inert attribute also makes the browser ignore input events sent by the user, including focus-related events and events from assistive technologies.")

Support in all current engines.

::: support
[Firefox112+]{.firefox .yes}[Safari15.5+]{.safari
.yes}[Chrome102+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge102+]{.edge_blink .yes}

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

The
[`inert`{#the-inert-attribute:the-inert-attribute}](#the-inert-attribute)
attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-inert-attribute:boolean-attribute}
that indicates, by its presence, that the element and all its [flat
tree](https://drafts.csswg.org/css-scoping/#flat-tree){#the-inert-attribute:flat-tree
x-internal="flat-tree"} descendants which don\'t otherwise escape
inertness (such as modal dialogs) are to be made
[inert](#inert){#the-inert-attribute:inert} by the user agent.

An inert subtree should not contain any content or controls which are
critical to understanding or using aspects of the page which are not in
the inert state. Content in an inert subtree will not be perceivable by
all users, or interactive. Authors should not specify elements as inert
unless the content they represent are also visually obscured in some
way. In most cases, authors should not specify the
[`inert`{#the-inert-attribute:inert-2}](#inert) attribute on individual
form controls. In these instances, the
[`disabled`{#the-inert-attribute:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
attribute is probably more appropriate.

::: example
The following example shows how to mark partially loaded content,
visually obscured by a \"loading\" message, as inert.

``` html
<section aria-labelledby=s1>
  <h3 id=s1>Population by City</h3>
  <div class=container>
    <div class=loading><p>Loading...</p></div>
    <div inert>
      <form>
        <fieldset>
          <legend>Date range</legend>
          <div>
            <label for=start>Start</label>
            <input type=date id=start>
          </div>
          <div>
            <label for=end>End</label>
            <input type=date id=end>
          </div>
          <div>
            <button>Apply</button>
          </div>
        </fieldset>
      </form>
      <table>
        <caption>From 20-- to 20--</caption>
        <thead>
          <tr>
            <th>City</th>
            <th>State</th>
            <th>20-- Population</th>
            <th>20-- Population</th>
            <th>Percentage change</th>
          </tr>
        </thead>
        <tbody>
          <!-- ... -->
        </tbody>
      </table>
    </div>
  </div>
</section>
```

![Screenshot of Population by City content with an overlaid loading
message which visually obscures the form controls and data table which
have not fully rendered, and thus are in the inert
state.](/images/inert-example-loading-section.png){width="947"
height="243"}

The \"loading\" overlay obscures the inert content, making it visually
apparent that the inert content is not presently accessible. Notice that
the heading and \"loading\" text are not descendants of the element with
the [`inert`{#the-inert-attribute:inert-3}](#inert) attribute. This will
ensure this text is accessible to all users, while the inert content
cannot be interacted with by anyone.
:::

::: note
By default, there is no persistent visual indication of an element or
its subtree being inert. Appropriate visual styles for such content is
often context-dependent. For instance, an inert off-screen navigation
panel would not require a default style, as its off-screen position
visually obscures the content. Similarly, a modal
[`dialog`{#the-inert-attribute:the-dialog-element}](interactive-elements.html#the-dialog-element)
element\'s backdrop will serve as the means to visually obscure the
inert content of the web page, rather than styling the inert content
specifically.

However, for many other situations authors are strongly encouraged to
clearly mark what parts of their document are active and which are
inert, to avoid user confusion. In particular, it is worth remembering
that not all users can see all parts of a page at once; for example,
users of screen readers, users on small devices or with magnifiers, and
even users using particularly small windows might not be able to see the
active part of a page and might get frustrated if inert sections are not
obviously inert.
:::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/inert](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/inert "The HTMLElement property inert reflects the value of the element's inert attribute. It is a boolean value that, when present, makes the browser "ignore" user input events for the element, including focus events and events from assistive technologies. The browser may also ignore page search and text selection in the element. This can be useful when building UIs such as modals where you would want to "trap" the focus inside the modal when it's visible.")

Support in all current engines.

::: support
[Firefox112+]{.firefox .yes}[Safari15.5+]{.safari
.yes}[Chrome102+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge102+]{.edge_blink .yes}

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

The [`inert`]{#dom-inert .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-inert-attribute:reflect}
the content attribute of the same name.

### [6.4]{.secno} Tracking [user activation]{.dfn export=""}[](#tracking-user-activation){.self-link}

To prevent abuse of certain APIs that could be annoying to users (e.g.,
opening popups or vibrating phones), user agents allow these APIs only
when the user is actively interacting with the web page or has
interacted with the page at least once. This \"active interaction\"
state is maintained through the mechanisms defined in this section.

#### [6.4.1]{.secno} Data model[](#user-activation-data-model){.self-link} {#user-activation-data-model}

For the purpose of tracking user activation, each
[`Window`{#user-activation-data-model:window}](nav-history-apis.html#window)
`W`{.variable} has two relevant values:

- A [last activation timestamp]{#last-activation-timestamp .dfn}, which
  is either a
  [`DOMHighResTimeStamp`{#user-activation-data-model:domhighrestimestamp}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"},
  positive infinity (indicating that `W`{.variable} has never been
  activated), or negative infinity (indicating that the activation has
  been
  [consumed](#consume-user-activation){#user-activation-data-model:consume-user-activation}).
  Initially positive infinity.

- A [last history-action activation
  timestamp]{#last-history-action-activation-timestamp .dfn}, which is
  either a
  [`DOMHighResTimeStamp`{#user-activation-data-model:domhighrestimestamp-2}](https://w3c.github.io/hr-time/#dom-domhighrestimestamp){x-internal="domhighrestimestamp"}
  or positive infinity, initially positive infinity.

A user agent also defines a [transient activation
duration]{#transient-activation-duration .dfn}, which is a constant
number indicating how long a user activation is available for certain
[user activation-gated APIs](#user-activation-gated-apis) (e.g., for
opening popups).

The [transient activation
duration](#transient-activation-duration){#user-activation-data-model:transient-activation-duration}
is expected be at most a few seconds, so that the user can possibly
perceive the link between an interaction with the page and the page
calling the activation-gated API.

We then have the following boolean user activation states for
`W`{.variable}:

[Sticky activation]{#sticky-activation .dfn export=""}

:   When the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#user-activation-data-model:current-high-resolution-time
    x-internal="current-high-resolution-time"} given `W`{.variable} is
    greater than or equal to the [last activation
    timestamp](#last-activation-timestamp){#user-activation-data-model:last-activation-timestamp}
    in `W`{.variable}, `W`{.variable} is said to have [sticky
    activation](#sticky-activation){#user-activation-data-model:sticky-activation}.

    This is `W`{.variable}\'s historical activation state, indicating
    whether the user has ever interacted in `W`{.variable}. It starts
    false, then changes to true (and never changes back to false) when
    `W`{.variable} gets the very first [activation
    notification](#activation-notification){#user-activation-data-model:activation-notification}.

[Transient activation]{#transient-activation .dfn export=""}

:   When the [current high resolution
    time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#user-activation-data-model:current-high-resolution-time-2
    x-internal="current-high-resolution-time"} given `W`{.variable} is
    greater than or equal to the [last activation
    timestamp](#last-activation-timestamp){#user-activation-data-model:last-activation-timestamp-2}
    in `W`{.variable}, and less than the [last activation
    timestamp](#last-activation-timestamp){#user-activation-data-model:last-activation-timestamp-3}
    in `W`{.variable} plus the [transient activation
    duration](#transient-activation-duration){#user-activation-data-model:transient-activation-duration-2},
    then `W`{.variable} is said to have [transient
    activation](#transient-activation){#user-activation-data-model:transient-activation}.

    This is `W`{.variable}\'s current activation state, indicating
    whether the user has interacted in `W`{.variable} recently. This
    starts with a false value, and remains true for a limited time after
    every [activation
    notification](#activation-notification){#user-activation-data-model:activation-notification-2}
    `W`{.variable} gets.

    The [transient
    activation](#transient-activation){#user-activation-data-model:transient-activation-2}
    state is considered [expired]{#activation-expiry .dfn} if it becomes
    false because the [transient activation
    duration](#transient-activation-duration){#user-activation-data-model:transient-activation-duration-3}
    time has elapsed since the last user activation. Note that it can
    become false even before the expiry time through an [activation
    consumption](#consume-user-activation){#user-activation-data-model:consume-user-activation-2}.

[History-action activation]{#history-action-activation .dfn}

:   When the [last history-action activation
    timestamp](#last-history-action-activation-timestamp){#user-activation-data-model:last-history-action-activation-timestamp}
    of `W`{.variable} is not equal to the [last activation
    timestamp](#last-activation-timestamp){#user-activation-data-model:last-activation-timestamp-4}
    of `W`{.variable}, then `W`{.variable} is said to have
    [history-action
    activation](#history-action-activation){#user-activation-data-model:history-action-activation}.

    This is a special variant of user activation, used to allow access
    to certain session history APIs which, if used too frequently, would
    make it harder for the user to traverse back using [browser
    UI](document-lifecycle.html#nav-traversal-ui). It starts with a
    false value, and becomes true whenever the user interacts with
    `W`{.variable}, but is reset to false through [history-action
    activation
    consumption](#consume-history-action-user-activation){#user-activation-data-model:consume-history-action-user-activation}.
    This ensures such APIs cannot be used multiple times in a row
    without an intervening user activation. But unlike [transient
    activation](#transient-activation){#user-activation-data-model:transient-activation-3},
    there is no time limit within which such APIs must be used.

The [last activation
timestamp](#last-activation-timestamp){#user-activation-data-model:last-activation-timestamp-5}
and [last history-action activation
timestamp](#last-history-action-activation-timestamp){#user-activation-data-model:last-history-action-activation-timestamp-2}
are retained even after the
[`Document`{#user-activation-data-model:document}](dom.html#document)
changes its [fully
active](document-sequences.html#fully-active){#user-activation-data-model:fully-active}
status (e.g., after navigating away from a
[`Document`{#user-activation-data-model:document-2}](dom.html#document),
or navigating to a cached
[`Document`{#user-activation-data-model:document-3}](dom.html#document)).
This means [sticky
activation](#sticky-activation){#user-activation-data-model:sticky-activation-2}
state spans multiple navigations as long as the same
[`Document`{#user-activation-data-model:document-4}](dom.html#document)
gets reused. For the transient activation state, the original
[expiry](#activation-expiry){#user-activation-data-model:activation-expiry}
time remains unchanged (i.e., the state still expires within the
[transient activation
duration](#transient-activation-duration){#user-activation-data-model:transient-activation-duration-4}
limit from the original [activation triggering input
event](#activation-triggering-input-event){#user-activation-data-model:activation-triggering-input-event}).
It is important to consider this when deciding whether to base certain
things off [sticky
activation](#sticky-activation){#user-activation-data-model:sticky-activation-3}
or [transient
activation](#transient-activation){#user-activation-data-model:transient-activation-4}.

#### [6.4.2]{.secno} Processing model[](#user-activation-processing-model){.self-link} {#user-activation-processing-model}

When a user interaction causes firing of an [activation triggering input
event](#activation-triggering-input-event){#user-activation-processing-model:activation-triggering-input-event}
in a
[`Document`{#user-activation-processing-model:document}](dom.html#document)
`document`{.variable}, the user agent must perform the following
[activation notification]{#activation-notification .dfn} steps *before*
[dispatching](https://dom.spec.whatwg.org/#concept-event-dispatch){#user-activation-processing-model:concept-event-dispatch
x-internal="concept-event-dispatch"} the event:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#user-activation-processing-model:assert
    x-internal="assert"}: `document`{.variable} is [fully
    active](document-sequences.html#fully-active){#user-activation-processing-model:fully-active}.

2.  Let `windows`{.variable} be « `document`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#user-activation-processing-model:concept-relevant-global}
    ».

3.  [Extend](https://infra.spec.whatwg.org/#list-extend){#user-activation-processing-model:list-extend
    x-internal="list-extend"} `windows`{.variable} with the [active
    window](document-sequences.html#nav-window){#user-activation-processing-model:nav-window}
    of each of `document`{.variable}\'s [ancestor
    navigables](document-sequences.html#ancestor-navigables){#user-activation-processing-model:ancestor-navigables}.

4.  [Extend](https://infra.spec.whatwg.org/#list-extend){#user-activation-processing-model:list-extend-2
    x-internal="list-extend"} `windows`{.variable} with the [active
    window](document-sequences.html#nav-window){#user-activation-processing-model:nav-window-2}
    of each of `document`{.variable}\'s [descendant
    navigables](document-sequences.html#descendant-navigables){#user-activation-processing-model:descendant-navigables},
    filtered to include only those
    [navigables](document-sequences.html#navigable){#user-activation-processing-model:navigable}
    whose [active
    document](document-sequences.html#nav-document){#user-activation-processing-model:nav-document}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#user-activation-processing-model:concept-document-origin
    x-internal="concept-document-origin"} is [same
    origin](browsers.html#same-origin){#user-activation-processing-model:same-origin}
    with `document`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#user-activation-processing-model:concept-document-origin-2
    x-internal="concept-document-origin"}.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#user-activation-processing-model:list-iterate
    x-internal="list-iterate"} `window`{.variable} in
    `windows`{.variable}:

    1.  Set `window`{.variable}\'s [last activation
        timestamp](#last-activation-timestamp){#user-activation-processing-model:last-activation-timestamp}
        to the [current high resolution
        time](https://w3c.github.io/hr-time/#dfn-current-high-resolution-time){#user-activation-processing-model:current-high-resolution-time
        x-internal="current-high-resolution-time"}.

    2.  [Notify the close watcher manager about user
        activation](#notify-the-close-watcher-manager-about-user-activation){#user-activation-processing-model:notify-the-close-watcher-manager-about-user-activation}
        given `window`{.variable}.

An [activation triggering input
event]{#activation-triggering-input-event .dfn} is any event whose
[`isTrusted`{#user-activation-processing-model:dom-event-istrusted}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
attribute is true and whose
[`type`{#user-activation-processing-model:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
is one of:

- \"[`keydown`{#user-activation-processing-model:event-keydown}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}\",
  provided the key is neither the [Esc]{.kbd} key nor a shortcut key
  reserved by the user agent;

- \"[`mousedown`{#user-activation-processing-model:event-mousedown}](https://w3c.github.io/uievents/#event-type-mousedown){x-internal="event-mousedown"}\";

- \"[`pointerdown`{#user-activation-processing-model:event-pointerdown}](https://w3c.github.io/pointerevents/#the-pointerdown-event){x-internal="event-pointerdown"}\",
  provided the event\'s
  [`pointerType`{#user-activation-processing-model:pointertype}](https://w3c.github.io/pointerevents/#dom-pointerevent-pointertype){x-internal="pointertype"}
  is \"`mouse`\";

- \"[`pointerup`{#user-activation-processing-model:event-pointerup}](https://w3c.github.io/pointerevents/#the-pointerup-event){x-internal="event-pointerup"}\",
  provided the event\'s
  [`pointerType`{#user-activation-processing-model:pointertype-2}](https://w3c.github.io/pointerevents/#dom-pointerevent-pointertype){x-internal="pointertype"}
  is not \"`mouse`\"; or

- \"[`touchend`{#user-activation-processing-model:event-touchend}](https://w3c.github.io/touch-events/#event-touchend){x-internal="event-touchend"}\".

[Activation consuming
APIs](#activation-consuming-api){#user-activation-processing-model:activation-consuming-api}
defined in this and other specifications can [consume user
activation]{#consume-user-activation .dfn export=""} by performing the
following steps, given a
[`Window`{#user-activation-processing-model:window}](nav-history-apis.html#window)
`W`{.variable}:

1.  If `W`{.variable}\'s
    [navigable](nav-history-apis.html#window-navigable){#user-activation-processing-model:window-navigable}
    is null, then return.

2.  Let `top`{.variable} be `W`{.variable}\'s
    [navigable](nav-history-apis.html#window-navigable){#user-activation-processing-model:window-navigable-2}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#user-activation-processing-model:nav-top}.

3.  Let `navigables`{.variable} be the [inclusive descendant
    navigables](document-sequences.html#inclusive-descendant-navigables){#user-activation-processing-model:inclusive-descendant-navigables}
    of `top`{.variable}\'s [active
    document](document-sequences.html#nav-document){#user-activation-processing-model:nav-document-2}.

4.  Let `windows`{.variable} be the list of
    [`Window`{#user-activation-processing-model:window-2}](nav-history-apis.html#window)
    objects constructed by taking the [active
    window](document-sequences.html#nav-window){#user-activation-processing-model:nav-window-3}
    of each item in `navigables`{.variable}.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#user-activation-processing-model:list-iterate-2
    x-internal="list-iterate"} `window`{.variable} in
    `windows`{.variable}, if `window`{.variable}\'s [last activation
    timestamp](#last-activation-timestamp){#user-activation-processing-model:last-activation-timestamp-2}
    is not positive infinity, then set `window`{.variable}\'s [last
    activation
    timestamp](#last-activation-timestamp){#user-activation-processing-model:last-activation-timestamp-3}
    to negative infinity.

[History-action activation-consuming
APIs](#history-action-activation-consuming-api){#user-activation-processing-model:history-action-activation-consuming-api}
can [consume history-action user
activation]{#consume-history-action-user-activation .dfn} by performing
the following steps, given a
[`Window`{#user-activation-processing-model:window-3}](nav-history-apis.html#window)
`W`{.variable}:

1.  If `W`{.variable}\'s
    [navigable](nav-history-apis.html#window-navigable){#user-activation-processing-model:window-navigable-3}
    is null, then return.

2.  Let `top`{.variable} be `W`{.variable}\'s
    [navigable](nav-history-apis.html#window-navigable){#user-activation-processing-model:window-navigable-4}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#user-activation-processing-model:nav-top-2}.

3.  Let `navigables`{.variable} be the [inclusive descendant
    navigables](document-sequences.html#inclusive-descendant-navigables){#user-activation-processing-model:inclusive-descendant-navigables-2}
    of `top`{.variable}\'s [active
    document](document-sequences.html#nav-document){#user-activation-processing-model:nav-document-3}.

4.  Let `windows`{.variable} be the list of
    [`Window`{#user-activation-processing-model:window-4}](nav-history-apis.html#window)
    objects constructed by taking the [active
    window](document-sequences.html#nav-window){#user-activation-processing-model:nav-window-4}
    of each item in `navigables`{.variable}.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#user-activation-processing-model:list-iterate-3
    x-internal="list-iterate"} `window`{.variable} in
    `windows`{.variable}, set `window`{.variable}\'s [last
    history-action activation
    timestamp](#last-history-action-activation-timestamp){#user-activation-processing-model:last-history-action-activation-timestamp}
    to `window`{.variable}\'s [last activation
    timestamp](#last-activation-timestamp){#user-activation-processing-model:last-activation-timestamp-4}.

Note the asymmetry in the sets of [browsing
contexts](document-sequences.html#browsing-context){#user-activation-processing-model:browsing-context}
in the page that are affected by an [activation
notification](#activation-notification){#user-activation-processing-model:activation-notification}
vs an [activation
consumption](#consume-user-activation){#user-activation-processing-model:consume-user-activation}:
an activation consumption changes (to false) the [transient
activation](#transient-activation){#user-activation-processing-model:transient-activation}
states for all browsing contexts in the page, but an activation
notification changes (to true) the states for a subset of those browsing
contexts. The exhaustive nature of consumption here is deliberate: it
prevents malicious sites from making multiple calls to an [activation
consuming
API](#activation-consuming-api){#user-activation-processing-model:activation-consuming-api-2}
from a single user activation (possibly by exploiting a deep hierarchy
of
[`iframe`{#user-activation-processing-model:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s).

#### [6.4.3]{.secno} APIs gated by user activation[](#user-activation-gated-apis){.self-link} {#user-activation-gated-apis}

APIs that are dependent on user activation are classified into different
levels:

[Sticky activation-gated APIs]{#sticky-activation-gated-api .dfn lt="sticky activation-gated API" export=""}
:   These APIs require the [sticky
    activation](#sticky-activation){#user-activation-gated-apis:sticky-activation}
    state to be true, so they are blocked until the very first user
    activation.

[Transient activation-gated APIs]{#transient-activation-gated-api .dfn lt="transient activation-gated
   API" export=""}
:   These APIs require the [transient
    activation](#transient-activation){#user-activation-gated-apis:transient-activation}
    state to be true, but they don\'t
    [consume](#consume-user-activation){#user-activation-gated-apis:consume-user-activation}
    it, so multiple calls are allowed per user activation until the
    transient state
    [expires](#activation-expiry){#user-activation-gated-apis:activation-expiry}.

[Transient activation-consuming APIs]{#activation-consuming-api .dfn lt="transient activation-consuming API" export=""}
:   These APIs require the [transient
    activation](#transient-activation){#user-activation-gated-apis:transient-activation-2}
    state to be true, and they [consume user
    activation](#consume-user-activation){#user-activation-gated-apis:consume-user-activation-2}
    in each call to prevent multiple calls per user activation.

[History-action activation-consuming APIs]{#history-action-activation-consuming-api .dfn lt="history-action activation-consuming API" export=""}

:   These APIs require the [history-action
    activation](#history-action-activation){#user-activation-gated-apis:history-action-activation}
    state to be true, and they [consume history-action user
    activation](#consume-history-action-user-activation){#user-activation-gated-apis:consume-history-action-user-activation}
    in each call to prevent multiple calls per user activation.

#### [6.4.4]{.secno} The [`UserActivation`{#the-useractivation-interface:useractivation}](#useractivation) interface[](#the-useractivation-interface){.self-link}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[UserActivation](https://developer.mozilla.org/en-US/docs/Web/API/UserActivation "The UserActivation interface allows querying information about a window's user activation state.")

::: support
[FirefoxNo]{.firefox .no}[Safari16.4+]{.safari .yes}[Chrome72+]{.chrome
.yes}

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

Each
[`Window`{#the-useractivation-interface:window}](nav-history-apis.html#window)
has an [associated `UserActivation`]{#associated-useractivation .dfn},
which is a
[`UserActivation`{#the-useractivation-interface:useractivation-2}](#useractivation)
object. Upon creation of the
[`Window`{#the-useractivation-interface:window-2}](nav-history-apis.html#window)
object, its [associated
`UserActivation`](#associated-useractivation){#the-useractivation-interface:associated-useractivation}
must be set to a
[new](https://webidl.spec.whatwg.org/#new){#the-useractivation-interface:new
x-internal="new"}
[`UserActivation`{#the-useractivation-interface:useractivation-3}](#useractivation)
object created in the
[`Window`{#the-useractivation-interface:window-3}](nav-history-apis.html#window)
object\'s [relevant
realm](webappapis.html#concept-relevant-realm){#the-useractivation-interface:concept-relevant-realm}.

``` idl
[Exposed=Window]
interface UserActivation {
  readonly attribute boolean hasBeenActive;
  readonly attribute boolean isActive;
};

partial interface Navigator {
  [SameObject] readonly attribute UserActivation userActivation;
};
```

[`navigator`](system-state.html#dom-navigator){#the-useractivation-interface:dom-navigator}`.`[`userActivation`](#dom-navigator-useractivation){#dom-navigator-useractivation-dev}`.`[`hasBeenActive`](#dom-useractivation-hasbeenactive){#dom-useractivation-hasbeenactive-dev}

:   Returns whether the window has [sticky
    activation](#sticky-activation){#the-useractivation-interface:sticky-activation}.

[`navigator`](system-state.html#dom-navigator){#the-useractivation-interface:dom-navigator-2}`.`[`userActivation`](#dom-navigator-useractivation){#the-useractivation-interface:dom-navigator-useractivation-2}`.`[`isActive`](#dom-useractivation-isactive){#dom-useractivation-isactive-dev}

:   Returns whether the window has [transient
    activation](#transient-activation){#the-useractivation-interface:transient-activation}.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Navigator/userActivation](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/userActivation "The read-only userActivation property of the Navigator interface returns a UserActivation object which contains information about the current window's user activation state.")

::: support
[FirefoxNo]{.firefox .no}[Safari16.4+]{.safari .yes}[Chrome72+]{.chrome
.yes}

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

The [`userActivation`]{#dom-navigator-useractivation .dfn
dfn-for="Navigator" dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-useractivation-interface:this
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#the-useractivation-interface:concept-relevant-global}\'s
[associated
`UserActivation`](#associated-useractivation){#the-useractivation-interface:associated-useractivation-2}.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[UserActivation/hasBeenActive](https://developer.mozilla.org/en-US/docs/Web/API/UserActivation/hasBeenActive "The read-only hasBeenActive property of the UserActivation interface indicates whether the current window has sticky user activation (see sticky activation).")

::: support
[FirefoxNo]{.firefox .no}[Safari16.4+]{.safari .yes}[Chrome72+]{.chrome
.yes}

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

The [`hasBeenActive`]{#dom-useractivation-hasbeenactive .dfn
dfn-for="UserActivation" dfn-type="attribute"} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#the-useractivation-interface:this-2
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#the-useractivation-interface:concept-relevant-global-2}
has [sticky
activation](#sticky-activation){#the-useractivation-interface:sticky-activation-2},
and false otherwise.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[UserActivation/hasBeenActive](https://developer.mozilla.org/en-US/docs/Web/API/UserActivation/hasBeenActive "The read-only hasBeenActive property of the UserActivation interface indicates whether the current window has sticky user activation (see sticky activation).")

::: support
[FirefoxNo]{.firefox .no}[Safari16.4+]{.safari .yes}[Chrome72+]{.chrome
.yes}

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

The [`isActive`]{#dom-useractivation-isactive .dfn
dfn-for="UserActivation" dfn-type="attribute"} getter steps are to
return true if
[this](https://webidl.spec.whatwg.org/#this){#the-useractivation-interface:this-3
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#the-useractivation-interface:concept-relevant-global-3}
has [transient
activation](#transient-activation){#the-useractivation-interface:transient-activation-2},
and false otherwise.

#### [6.4.5]{.secno} User agent automation[](#user-activation-user-agent-automation){.self-link} {#user-activation-user-agent-automation}

For the purposes of user-agent automation and application testing, this
specification defines the following [extension
command](https://w3c.github.io/webdriver/#dfn-extension-commands){#user-activation-user-agent-automation:extension-command
x-internal="extension-command"} for the Web Driver specification. It is
optional for a user agent to support the following [extension
command](https://w3c.github.io/webdriver/#dfn-extension-commands){#user-activation-user-agent-automation:extension-command-2
x-internal="extension-command"}.
[\[WEBDRIVER\]](references.html#refsWEBDRIVER)

HTTP Method

URI Template

\``POST`\`

`/session/{``session id`{.variable}`}/window/consume-user-activation`

The [remote end
steps](https://w3c.github.io/webdriver/#dfn-remote-end-steps){#user-activation-user-agent-automation:remote-end-steps
x-internal="remote-end-steps"} are:

1.  Let `window`{.variable} be the [current browsing
    context](https://w3c.github.io/webdriver/#dfn-current-browsing-context){#user-activation-user-agent-automation:webdriver-current-browsing-context
    x-internal="webdriver-current-browsing-context"}\'s [active
    window](document-sequences.html#active-window){#user-activation-user-agent-automation:active-window}.

2.  Let `consume`{.variable} be true if `window`{.variable} has
    [transient
    activation](#transient-activation){#user-activation-user-agent-automation:transient-activation};
    otherwise false.

3.  If `consume`{.variable} is true, then [consume user
    activation](#consume-user-activation){#user-activation-user-agent-automation:consume-user-activation}
    of `window`{.variable}.

4.  Return
    [success](https://w3c.github.io/webdriver/#dfn-success){#user-activation-user-agent-automation:success-value
    x-internal="success-value"} with data `consume`{.variable}.

### [6.5]{.secno} Activation behavior of elements[](#activation){.self-link} {#activation}

Certain elements in HTML have an [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#activation:activation-behaviour
x-internal="activation-behaviour"}, which means that the user can
activate them. This is always caused by a
[`click`{#activation:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event.

The user agent should allow the user to manually trigger elements that
have an [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#activation:activation-behaviour-2
x-internal="activation-behaviour"}, for instance using keyboard or voice
input, or through mouse clicks. When the user triggers an element with a
defined [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#activation:activation-behaviour-3
x-internal="activation-behaviour"} in a manner other than clicking it,
the default action of the interaction event must be to [fire a `click`
event](webappapis.html#fire-a-click-event){#activation:fire-a-click-event}
at the element.

`element`{.variable}`.`[`click`](#dom-click){#dom-click-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/click](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click "The HTMLElement.click() method simulates a mouse click on an element.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet1.0+]{.samsunginternet_android .yes}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Acts as if the element was clicked.

Each element has an associated [click in progress
flag]{#click-in-progress-flag .dfn}, which is initially unset.

The [`click()`]{#dom-click .dfn dfn-for="HTMLElement" dfn-type="method"}
method must run the following steps:

1.  If this element is a form control that is
    [disabled](form-control-infrastructure.html#concept-fe-disabled){#activation:concept-fe-disabled},
    then return.

2.  If this element\'s [click in progress
    flag](#click-in-progress-flag){#activation:click-in-progress-flag}
    is set, then return.

3.  Set this element\'s [click in progress
    flag](#click-in-progress-flag){#activation:click-in-progress-flag-2}.

4.  [Fire a synthetic pointer
    event](webappapis.html#fire-a-synthetic-pointer-event){#activation:fire-a-synthetic-pointer-event}
    named
    [`click`{#activation:event-click-2}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
    at this element, with the *not trusted flag* set.

5.  Unset this element\'s [click in progress
    flag](#click-in-progress-flag){#activation:click-in-progress-flag-3}.

#### [6.5.1]{.secno} The [`ToggleEvent`{#the-toggleevent-interface:toggleevent}](#toggleevent) interface[](#the-toggleevent-interface){.self-link}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ToggleEvent/ToggleEvent](https://developer.mozilla.org/en-US/docs/Web/API/ToggleEvent/ToggleEvent "The ToggleEvent() constructor creates a new ToggleEvent object.")

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
[ToggleEvent](https://developer.mozilla.org/en-US/docs/Web/API/ToggleEvent "The ToggleEvent interface represents an event notifying the user when a popover element's state toggles between showing and hidden.")

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

``` idl
[Exposed=Window]
interface ToggleEvent : Event {
  constructor(DOMString type, optional ToggleEventInit eventInitDict = {});
  readonly attribute DOMString oldState;
  readonly attribute DOMString newState;
  readonly attribute Element source;
};

dictionary ToggleEventInit : EventInit {
  DOMString oldState = "";
  DOMString newState = "";
};
```

`event`{.variable}`.`[`oldState`](#dom-toggleevent-oldstate){#dom-toggleevent-oldstate-dev}

:   Set to \"`closed`\" when transitioning from closed to open, or set
    to \"`open`\" when transitioning from open to closed.

`event`{.variable}`.`[`newState`](#dom-toggleevent-newstate){#dom-toggleevent-newstate-dev}

:   Set to \"`open`\" when transitioning from closed to open, or set to
    \"`closed`\" when transitioning from open to closed.

`event`{.variable}`.`[`source`](#dom-toggleevent-source){#dom-toggleevent-source-dev}

:   Set to the element which initiated the toggle, which can be set up
    with the
    [`popovertarget`{#the-toggleevent-interface:attr-popovertarget}](popover.html#attr-popovertarget)
    and
    [`commandfor`{#the-toggleevent-interface:attr-button-commandfor}](form-elements.html#attr-button-commandfor)
    attributes. If there is no source element, then it is set to null.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ToggleEvent/oldState](https://developer.mozilla.org/en-US/docs/Web/API/ToggleEvent/oldState "The oldState read-only property of the ToggleEvent interface is a string representing the state the element is transitioning from.")

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
[ToggleEvent/newState](https://developer.mozilla.org/en-US/docs/Web/API/ToggleEvent/newState "The newState read-only property of the ToggleEvent interface is a string representing the state the element is transitioning to.")

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

The [`oldState`]{#dom-toggleevent-oldstate .dfn dfn-for="ToggleEvent"
dfn-type="attribute"} and [`newState`]{#dom-toggleevent-newstate .dfn
dfn-for="ToggleEvent" dfn-type="attribute"} attributes must return the
values they are initialized to.

The [`source`]{#dom-toggleevent-source .dfn dfn-for="ToggleEvent"
dfn-type="attribute"} getter steps are to return the result of
[retargeting](https://dom.spec.whatwg.org/#retarget){#the-toggleevent-interface:dom-retarget
x-internal="dom-retarget"}
[`source`{#the-toggleevent-interface:dom-toggleevent-source-2}](#dom-toggleevent-source)
against
[this](https://webidl.spec.whatwg.org/#this){#the-toggleevent-interface:this
x-internal="this"}\'s
[`currentTarget`{#the-toggleevent-interface:dom-event-currenttarget}](https://dom.spec.whatwg.org/#dom-event-currenttarget){x-internal="dom-event-currenttarget"}.

[DOM standard issue #1328](https://github.com/whatwg/dom/issues/1328)
tracks how to better standardize associated event data in a way which
makes sense on Events. Currently an event attribute initialized to a
value cannot also have a getter, and so an internal slot (or map of
additional fields) is required to properly specify this.

A [toggle task tracker]{#toggle-task-tracker .dfn} is a
[struct](https://infra.spec.whatwg.org/#struct){#the-toggleevent-interface:struct
x-internal="struct"} which has:

[task]{#toggle-task-task .dfn}
:   A
    [task](webappapis.html#concept-task){#the-toggleevent-interface:concept-task}
    which fires a
    [`ToggleEvent`{#the-toggleevent-interface:toggleevent-2}](#toggleevent).

[old state]{#toggle-task-old-state .dfn}
:   A string which represents the
    [task](#toggle-task-task){#the-toggleevent-interface:toggle-task-task}\'s
    event\'s value for the
    [`oldState`{#the-toggleevent-interface:dom-toggleevent-oldstate-2}](#dom-toggleevent-oldstate)
    attribute.

#### [6.5.2]{.secno} The [`CommandEvent`{#the-commandevent-interface:commandevent}](#commandevent) interface[](#the-commandevent-interface){.self-link}

``` idl
[Exposed=Window]
interface CommandEvent : Event {
  constructor(DOMString type, optional CommandEventInit eventInitDict = {});
  readonly attribute Element? source;
  readonly attribute DOMString command;
};

dictionary CommandEventInit : EventInit {
  Element? source = null;
  DOMString command = "";
};
```

`event`{.variable}`.`[`command`](#dom-commandevent-command){#dom-commandevent-command-dev}

:   Returns what action the element can take.

`event`{.variable}`.`[`source`](#dom-commandevent-source){#dom-commandevent-source-dev}

:   Returns the
    [`Element`{#the-commandevent-interface:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    that was interacted with in order to cause this event.

The [`command`]{#dom-commandevent-command .dfn dfn-for="CommandEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to.

The [`source`]{#dom-commandevent-source .dfn dfn-for="CommandEvent"
dfn-type="attribute"} getter steps are to return the result of
[retargeting](https://dom.spec.whatwg.org/#retarget){#the-commandevent-interface:dom-retarget
x-internal="dom-retarget"}
[`source`{#the-commandevent-interface:dom-commandevent-source-2}](#dom-commandevent-source)
against
[this](https://webidl.spec.whatwg.org/#this){#the-commandevent-interface:this
x-internal="this"}\'s
[`currentTarget`{#the-commandevent-interface:dom-event-currenttarget}](https://dom.spec.whatwg.org/#dom-event-currenttarget){x-internal="dom-event-currenttarget"}.

[DOM standard issue #1328](https://github.com/whatwg/dom/issues/1328)
tracks how to better standardize associated event data in a way which
makes sense on Events. Currently an event attribute initialized to a
value cannot also have a getter, and so an internal slot (or map of
additional fields) is required to properly specify this.

### [6.6]{.secno} Focus[](#focus){.self-link}

#### [6.6.1]{.secno} Introduction[](#introduction-8){.self-link} {#introduction-8}

*This section is non-normative.*

An HTML user interface typically consists of multiple interactive
widgets, such as form controls, scrollable regions, links, dialog boxes,
browser tabs, and so forth. These widgets form a hierarchy, with some
(e.g. browser tabs, dialog boxes) containing others (e.g. links, form
controls).

When interacting with an interface using a keyboard, key input is
channeled from the system, through the hierarchy of interactive widgets,
to an active widget, which is said to be
[focused](#focused){#introduction-8:focused}.

::: example
Consider an HTML application running in a browser tab running in a
graphical environment. Suppose this application had a page with some
text controls and links, and was currently showing a modal dialog, which
itself had a text control and a button.

The hierarchy of focusable widgets, in this scenario, would include the
browser window, which would have, amongst its children, the browser tab
containing the HTML application. The tab itself would have as its
children the various links and text controls, as well as the dialog. The
dialog itself would have as its children the text control and the
button.

![](/images/focus-tree.png){width="800" height="450"}

If the widget with [focus](#focused){#introduction-8:focused-2} in this
example was the text control in the dialog box, then key input would be
channeled from the graphical system to ① the web browser, then to ② the
tab, then to ③ the dialog, and finally to ④ the text control.
:::

Keyboard *events* are always targeted at this
[focused](#focused){#introduction-8:focused-3} element.

#### [6.6.2]{.secno} Data model[](#data-model){.self-link} {#data-model}

A [top-level
traversable](document-sequences.html#top-level-traversable){#data-model:top-level-traversable}
has [system focus]{#system-focus .dfn dfn-for="top-level
  traversable" export=""} when it can receive keyboard input channeled
from the operating system, possibly targeted at one of its [active
document](document-sequences.html#nav-document){#data-model:nav-document}\'s
[descendant
navigables](document-sequences.html#descendant-navigables){#data-model:descendant-navigables}.

A [top-level
traversable](document-sequences.html#top-level-traversable){#data-model:top-level-traversable-2}
has [user attention]{#user-attention .dfn
dfn-for="top-level traversable"} when its [system visibility
state](document-sequences.html#system-visibility-state){#data-model:system-visibility-state}
is \"`visible`\", and it either has [system
focus](#system-focus){#data-model:system-focus} or user agent widgets
directly related to it can receive keyboard input channeled from the
operating system.

User attention is lost when a browser window loses focus, whereas system
focus might also be lost to other system widgets in the browser window
such as a location bar.

A [`Document`{#data-model:document}](dom.html#document) `d`{.variable}
is a [fully active descendant of a top-level traversable with user
attention]{#fully-active-descendant-of-a-top-level-traversable-with-user-attention
.dfn dfn-for="Document" export=""} when `d`{.variable} is [fully
active](document-sequences.html#fully-active){#data-model:fully-active}
and `d`{.variable}\'s [node
navigable](document-sequences.html#node-navigable){#data-model:node-navigable}\'s
[top-level
traversable](document-sequences.html#nav-top){#data-model:nav-top} has
[user attention](#user-attention){#data-model:user-attention}.

The term [focusable area]{#focusable-area .dfn} is used to refer to
regions of the interface that can further become the target of such
keyboard input. Focusable areas can be elements, parts of elements, or
other regions managed by the user agent.

Each [focusable area](#focusable-area){#data-model:focusable-area} has a
[DOM anchor]{#dom-anchor .dfn}, which is a
[`Node`{#data-model:node}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"}
object that represents the position of the [focusable
area](#focusable-area){#data-model:focusable-area-2} in the DOM. (When
the [focusable area](#focusable-area){#data-model:focusable-area-3} is
itself a
[`Node`{#data-model:node-2}](https://dom.spec.whatwg.org/#interface-node){x-internal="node"},
it is its own [DOM anchor](#dom-anchor){#data-model:dom-anchor}.) The
[DOM anchor](#dom-anchor){#data-model:dom-anchor-2} is used in some APIs
as a substitute for the [focusable
area](#focusable-area){#data-model:focusable-area-4} when there is no
other DOM object to represent the [focusable
area](#focusable-area){#data-model:focusable-area-5}.

The following table describes what objects can be [focusable
areas](#focusable-area){#data-model:focusable-area-6}. The cells in the
left column describe objects that can be [focusable
areas](#focusable-area){#data-model:focusable-area-7}; the cells in the
right column describe the [DOM
anchors](#dom-anchor){#data-model:dom-anchor-3} for those elements. (The
cells that span both columns are non-normative examples.)

[Focusable area](#focusable-area){#data-model:focusable-area-8}

[DOM anchor](#dom-anchor){#data-model:dom-anchor-4}

Examples

Elements that meet all the following criteria:

- the element\'s [tabindex
  value](#tabindex-value){#data-model:tabindex-value} is non-null, or
  the element is determined by the user agent to be focusable;
- the element is either not a [shadow
  host](https://dom.spec.whatwg.org/#element-shadow-host){#data-model:shadow-host
  x-internal="shadow-host"}, or has a [shadow
  root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#data-model:concept-element-shadow-root
  x-internal="concept-element-shadow-root"} whose [delegates
  focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus){#data-model:delegates-focus
  x-internal="delegates-focus"} is false;
- the element is not [actually
  disabled](semantics-other.html#concept-element-disabled){#data-model:concept-element-disabled};
- the element is not [inert](#inert){#data-model:inert};
- the element is either [being
  rendered](rendering.html#being-rendered){#data-model:being-rendered},
  [delegating its rendering to its
  children](rendering.html#delegating-its-rendering-to-its-children){#data-model:delegating-its-rendering-to-its-children},
  or [being used as relevant canvas fallback
  content](canvas.html#being-used-as-relevant-canvas-fallback-content){#data-model:being-used-as-relevant-canvas-fallback-content}.

The element itself.

[`iframe`{#data-model:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
[`dialog`{#data-model:the-dialog-element}](interactive-elements.html#the-dialog-element),
[`<input type=text>`{#data-model:text-(type=text)-state-and-search-state-(type=search)}](input.html#text-(type=text)-state-and-search-state-(type=search)),
sometimes
[`<a href="">`{#data-model:the-a-element}](text-level-semantics.html#the-a-element)
(depending on platform conventions).

The shapes of
[`area`{#data-model:the-area-element}](image-maps.html#the-area-element)
elements in an [image
map](image-maps.html#image-map){#data-model:image-map} associated with
an
[`img`{#data-model:the-img-element}](embedded-content.html#the-img-element)
element that is [being
rendered](rendering.html#being-rendered){#data-model:being-rendered-2}
and is not [inert](#inert){#data-model:inert-2}.

The
[`img`{#data-model:the-img-element-2}](embedded-content.html#the-img-element)
element.

::: example
In the following example, the
[`area`{#data-model:the-area-element-2}](image-maps.html#the-area-element)
element creates two shapes, one on each image. The [DOM
anchor](#dom-anchor){#data-model:dom-anchor-5} of the first shape is the
first
[`img`{#data-model:the-img-element-3}](embedded-content.html#the-img-element)
element, and the [DOM anchor](#dom-anchor){#data-model:dom-anchor-6} of
the second shape is the second
[`img`{#data-model:the-img-element-4}](embedded-content.html#the-img-element)
element.

``` html
<map id=wallmap><area alt="Enter Door" coords="10,10,100,200" href="door.html"></map>
...
<img src="images/innerwall.jpeg" alt="There is a white wall here, with a door." usemap="#wallmap">
...
<img src="images/outerwall.jpeg" alt="There is a red wall here, with a door." usemap="#wallmap">
```
:::

The user-agent provided subwidgets of elements that are [being
rendered](rendering.html#being-rendered){#data-model:being-rendered-3}
and are not [actually
disabled](semantics-other.html#concept-element-disabled){#data-model:concept-element-disabled-2}
or [inert](#inert){#data-model:inert-3}.

The element for which the [focusable
area](#focusable-area){#data-model:focusable-area-9} is a subwidget.

The [controls in the user
interface](media.html#expose-a-user-interface-to-the-user){#data-model:expose-a-user-interface-to-the-user}
for a
[`video`{#data-model:the-video-element}](media.html#the-video-element)
element, the up and down buttons in a spin-control version of
[`<input type=number>`{#data-model:number-state-(type=number)}](input.html#number-state-(type=number)),
the part of a
[`details`{#data-model:the-details-element}](interactive-elements.html#the-details-element)
element\'s rendering that enables the element to be opened or closed
using keyboard input.

The scrollable regions of elements that are [being
rendered](rendering.html#being-rendered){#data-model:being-rendered-4}
and are not [inert](#inert){#data-model:inert-4}.

The element for which the box that the scrollable region scrolls was
created.

The CSS
[\'overflow\'](https://drafts.csswg.org/css-overflow/#propdef-overflow){#data-model:'overflow'
x-internal="'overflow'"} property\'s \'scroll\' value typically creates
a scrollable region.

The
[viewport](https://drafts.csswg.org/css2/#viewport){#data-model:viewport
x-internal="viewport"} of a
[`Document`{#data-model:document-2}](dom.html#document) that has a
non-null [browsing
context](document-sequences.html#concept-document-bc){#data-model:concept-document-bc}
and is not [inert](#inert){#data-model:inert-5}.

The [`Document`{#data-model:document-3}](dom.html#document) for which
the
[viewport](https://drafts.csswg.org/css2/#viewport){#data-model:viewport-2
x-internal="viewport"} was created.

The contents of an
[`iframe`{#data-model:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element).

Any other element or part of an element determined by the user agent to
be a focusable area, especially to aid with accessibility or to better
match platform conventions.

The element.

A user agent could make all list item bullets [sequentially
focusable](#sequentially-focusable){#data-model:sequentially-focusable},
so that a user can more easily navigate lists.

Similarly, a user agent could make all elements with
[`title`{#data-model:attr-title}](dom.html#attr-title) attributes
[sequentially
focusable](#sequentially-focusable){#data-model:sequentially-focusable-2},
so that their advisory information can be accessed.

A [navigable
container](document-sequences.html#navigable-container){#data-model:navigable-container}
(e.g. an
[`iframe`{#data-model:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element))
is a [focusable area](#focusable-area){#data-model:focusable-area-10},
but key events routed to a [navigable
container](document-sequences.html#navigable-container){#data-model:navigable-container-2}
get immediately routed to its [content
navigable](document-sequences.html#content-navigable){#data-model:content-navigable}\'s
[active
document](document-sequences.html#nav-document){#data-model:nav-document-2}.
Similarly, in sequential focus navigation a [navigable
container](document-sequences.html#navigable-container){#data-model:navigable-container-3}
essentially acts merely as a placeholder for its [content
navigable](document-sequences.html#content-navigable){#data-model:content-navigable-2}\'s
[active
document](document-sequences.html#nav-document){#data-model:nav-document-3}.

------------------------------------------------------------------------

One [focusable area](#focusable-area){#data-model:focusable-area-11} in
each [`Document`{#data-model:document-4}](dom.html#document) is
designated the [focused area of the
document]{#focused-area-of-the-document .dfn}. Which control is so
designated changes over time, based on algorithms in this specification.

Even if a document is not [fully
active](document-sequences.html#fully-active){#data-model:fully-active-2}
and not shown to the user, it can still have a [focused area of the
document](#focused-area-of-the-document){#data-model:focused-area-of-the-document}.
If a document\'s [fully
active](document-sequences.html#fully-active){#data-model:fully-active-3}
state changes, its [focused area of the
document](#focused-area-of-the-document){#data-model:focused-area-of-the-document-2}
will stay the same.

The [currently focused area of a top-level
traversable]{#currently-focused-area-of-a-top-level-traversable .dfn}
`traversable`{.variable} is the [focusable
area](#focusable-area){#data-model:focusable-area-12}-or-null returned
by this algorithm:

1.  If `traversable`{.variable} does not have [system
    focus](#system-focus){#data-model:system-focus-2}, then return null.

2.  Let `candidate`{.variable} be `traversable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#data-model:nav-document-4}.

3.  While `candidate`{.variable}\'s [focused
    area](#focused-area-of-the-document){#data-model:focused-area-of-the-document-3}
    is a [navigable
    container](document-sequences.html#navigable-container){#data-model:navigable-container-4}
    with a non-null [content
    navigable](document-sequences.html#content-navigable){#data-model:content-navigable-3}:
    set `candidate`{.variable} to the [active
    document](document-sequences.html#nav-document){#data-model:nav-document-5}
    of that [navigable
    container](document-sequences.html#navigable-container){#data-model:navigable-container-5}\'s
    [content
    navigable](document-sequences.html#content-navigable){#data-model:content-navigable-4}.

4.  If `candidate`{.variable}\'s [focused
    area](#focused-area-of-the-document){#data-model:focused-area-of-the-document-4}
    is non-null, set `candidate`{.variable} to `candidate`{.variable}\'s
    [focused
    area](#focused-area-of-the-document){#data-model:focused-area-of-the-document-5}.

5.  Return `candidate`{.variable}.

The [current focus chain of a top-level
traversable]{#current-focus-chain-of-a-top-level-traversable .dfn}
`traversable`{.variable} is the [focus
chain](#focus-chain){#data-model:focus-chain} of the [currently focused
area](#currently-focused-area-of-a-top-level-traversable){#data-model:currently-focused-area-of-a-top-level-traversable}
of `traversable`{.variable}, if `traversable`{.variable} is non-null, or
an empty list otherwise.

An element that is the [DOM
anchor](#dom-anchor){#data-model:dom-anchor-7} of a [focusable
area](#focusable-area){#data-model:focusable-area-13} is said to [gain
focus]{#gains-focus .dfn} when that [focusable
area](#focusable-area){#data-model:focusable-area-14} becomes the
[currently focused area of a top-level
traversable](#currently-focused-area-of-a-top-level-traversable){#data-model:currently-focused-area-of-a-top-level-traversable-2}.
When an element is the [DOM
anchor](#dom-anchor){#data-model:dom-anchor-8} of a [focusable
area](#focusable-area){#data-model:focusable-area-15} of the [currently
focused area of a top-level
traversable](#currently-focused-area-of-a-top-level-traversable){#data-model:currently-focused-area-of-a-top-level-traversable-3},
it is [focused]{#focused .dfn}.

The [focus chain]{#focus-chain .dfn} of a [focusable
area](#focusable-area){#data-model:focusable-area-16}
`subject`{.variable} is the ordered list constructed as follows:

1.  Let `output`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#data-model:list
    x-internal="list"}.

2.  Let `currentObject`{.variable} be `subject`{.variable}.

3.  While true:

    1.  [Append](https://infra.spec.whatwg.org/#list-append){#data-model:list-append
        x-internal="list-append"} `currentObject`{.variable} to
        `output`{.variable}.

    2.  If `currentObject`{.variable} is an
        [`area`{#data-model:the-area-element-3}](image-maps.html#the-area-element)
        element\'s shape, then
        [append](https://infra.spec.whatwg.org/#list-append){#data-model:list-append-2
        x-internal="list-append"} that
        [`area`{#data-model:the-area-element-4}](image-maps.html#the-area-element)
        element to `output`{.variable}.

        Otherwise, if `currentObject`{.variable}\'s [DOM
        anchor](#dom-anchor){#data-model:dom-anchor-9} is an element
        that is not `currentObject`{.variable} itself, then
        [append](https://infra.spec.whatwg.org/#list-append){#data-model:list-append-3
        x-internal="list-append"} `currentObject`{.variable}\'s [DOM
        anchor](#dom-anchor){#data-model:dom-anchor-10} to
        `output`{.variable}.

    3.  If `currentObject`{.variable} is a [focusable
        area](#focusable-area){#data-model:focusable-area-17}, then set
        `currentObject`{.variable} to `currentObject`{.variable}\'s [DOM
        anchor](#dom-anchor){#data-model:dom-anchor-11}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#data-model:node-document
        x-internal="node-document"}.

        Otherwise, if `currentObject`{.variable} is a
        [`Document`{#data-model:document-5}](dom.html#document) whose
        [node
        navigable](document-sequences.html#node-navigable){#data-model:node-navigable-2}\'s
        [parent](document-sequences.html#nav-parent){#data-model:nav-parent}
        is non-null, then set `currentObject`{.variable} to
        `currentObject`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#data-model:node-navigable-3}\'s
        [parent](document-sequences.html#nav-parent){#data-model:nav-parent-2}.

        Otherwise,
        [break](https://infra.spec.whatwg.org/#iteration-break){#data-model:break
        x-internal="break"}.

4.  Return `output`{.variable}.

    The chain starts with `subject`{.variable} and (if
    `subject`{.variable} is or can be the [currently focused area of a
    top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#data-model:currently-focused-area-of-a-top-level-traversable-4})
    continues up the focus hierarchy up to the
    [`Document`{#data-model:document-6}](dom.html#document) of the
    [top-level
    traversable](document-sequences.html#top-level-traversable){#data-model:top-level-traversable-3}.

All elements that are [focusable
areas](#focusable-area){#data-model:focusable-area-18} are said to be
[focusable]{#focusable .dfn}.

There are two special types of focusability for [focusable
areas](#focusable-area){#data-model:focusable-area-19}:

- A [focusable area](#focusable-area){#data-model:focusable-area-20} is
  said to be [sequentially focusable]{#sequentially-focusable .dfn} if
  it is included in its
  [`Document`{#data-model:document-7}](dom.html#document)\'s [sequential
  focus navigation
  order](#sequential-focus-navigation-order){#data-model:sequential-focus-navigation-order}
  and the user agent determines that it is sequentially focusable.

- A [focusable area](#focusable-area){#data-model:focusable-area-21} is
  said to be [click focusable]{#click-focusable .dfn} if the user agent
  determines that it is click focusable. User agents should consider
  focusable areas with non-null [tabindex
  values](#tabindex-value){#data-model:tabindex-value-2} to be click
  focusable.

Elements which are not [focusable](#focusable){#data-model:focusable}
are not [focusable
areas](#focusable-area){#data-model:focusable-area-22}, and thus not
[sequentially
focusable](#sequentially-focusable){#data-model:sequentially-focusable-3}
and not [click
focusable](#click-focusable){#data-model:click-focusable}.

::: note
Being [focusable](#focusable){#data-model:focusable-2} is a statement
about whether an element can be focused programmatically, e.g. via the
[`focus()`{#data-model:dom-focus}](#dom-focus) method or
[`autofocus`{#data-model:attr-fe-autofocus}](#attr-fe-autofocus)
attribute. In contrast, [sequentially
focusable](#sequentially-focusable){#data-model:sequentially-focusable-4}
and [click focusable](#click-focusable){#data-model:click-focusable-2}
govern how the user agent responds to user interaction: respectively, to
[sequential focus
navigation](#sequential-focus-navigation){#data-model:sequential-focus-navigation}
and as [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#data-model:activation-behaviour
x-internal="activation-behaviour"}.

The user agent might determine that an element is not [sequentially
focusable](#sequentially-focusable){#data-model:sequentially-focusable-5}
even if it is [focusable](#focusable){#data-model:focusable-3} and is
included in its
[`Document`{#data-model:document-8}](dom.html#document)\'s [sequential
focus navigation
order](#sequential-focus-navigation-order){#data-model:sequential-focus-navigation-order-2},
according to user preferences. For example, macOS users can set the user
agent to skip non-form control elements, or can skip links when doing
[sequential focus
navigation](#sequential-focus-navigation){#data-model:sequential-focus-navigation-2}
with just the [Tab]{.kbd} key (as opposed to using both the
[Option]{.kbd} and [Tab]{.kbd} keys).

Similarly, the user agent might determine that an element is not [click
focusable](#click-focusable){#data-model:click-focusable-3} even if it
is [focusable](#focusable){#data-model:focusable-4}. For example, in
some user agents, clicking on a non-editable form control does not focus
it, i.e. the user agent has determined that such controls are not click
focusable.

Thus, an element can be
[focusable](#focusable){#data-model:focusable-5}, but neither
[sequentially
focusable](#sequentially-focusable){#data-model:sequentially-focusable-6}
nor [click focusable](#click-focusable){#data-model:click-focusable-4}.
For example, in some user agents, a non-editable form-control with a
negative-integer [tabindex
value](#tabindex-value){#data-model:tabindex-value-3} would not be
focusable via user interaction, only via programmatic APIs.
:::

When a user [activates](#activation) a [click
focusable](#click-focusable){#data-model:click-focusable-5} [focusable
area](#focusable-area){#data-model:focusable-area-23}, the user agent
must run the [focusing
steps](#focusing-steps){#data-model:focusing-steps} on the [focusable
area](#focusable-area){#data-model:focusable-area-24} with
`focus trigger`{.variable} set to \"`click`\".

Note that focusing is not an [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#data-model:activation-behaviour-2
x-internal="activation-behaviour"}, i.e. calling the
[`click()`{#data-model:dom-click}](#dom-click) method on an element or
dispatching a synthetic
[`click`{#data-model:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event on it won\'t cause the element to get focused.

------------------------------------------------------------------------

A node is a [focus navigation scope owner]{#focus-navigation-scope-owner
.dfn} if it is a
[`Document`{#data-model:document-9}](dom.html#document), a [shadow
host](https://dom.spec.whatwg.org/#element-shadow-host){#data-model:shadow-host-2
x-internal="shadow-host"}, a
[slot](scripting.html#the-slot-element){#data-model:the-slot-element},
or an element which is the [popover
invoker](popover.html#popover-invoker){#data-model:popover-invoker} of
an element in the [popover showing
state](popover.html#popover-showing-state){#data-model:popover-showing-state}.

Each [focus navigation scope
owner](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner}
has a [focus navigation scope]{#focus-navigation-scope .dfn}, which is a
list of elements. Its contents are determined as follows:

Every element `element`{.variable} has an [associated focus navigation
owner]{#associated-focus-navigation-owner .dfn}, which is either null or
a [focus navigation scope
owner](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner-2}.
It is determined by the following algorithm:

1.  If `element`{.variable}\'s parent is null, then return null.

2.  If `element`{.variable}\'s parent is a [shadow
    host](https://dom.spec.whatwg.org/#element-shadow-host){#data-model:shadow-host-3
    x-internal="shadow-host"}, then return `element`{.variable}\'s
    [assigned
    slot](https://dom.spec.whatwg.org/#slotable-assigned-slot){#data-model:assigned-slot
    x-internal="assigned-slot"}.

3.  If `element`{.variable}\'s parent is a [shadow
    root](https://dom.spec.whatwg.org/#concept-shadow-root){#data-model:shadow-root
    x-internal="shadow-root"}, then return the parent\'s
    [host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#data-model:concept-documentfragment-host
    x-internal="concept-documentfragment-host"}.

4.  If `element`{.variable}\'s parent is the [document
    element](https://dom.spec.whatwg.org/#document-element){#data-model:document-element
    x-internal="document-element"}, then return the parent\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#data-model:node-document-2
    x-internal="node-document"}.

5.  If `element`{.variable} is in the [popover showing
    state](popover.html#popover-showing-state){#data-model:popover-showing-state-2}
    and has a [popover
    invoker](popover.html#popover-invoker){#data-model:popover-invoker-2}
    set, then return `element`{.variable}\'s [popover
    invoker](popover.html#popover-invoker){#data-model:popover-invoker-3}.

6.  Return `element`{.variable}\'s parent\'s [associated focus
    navigation
    owner](#associated-focus-navigation-owner){#data-model:associated-focus-navigation-owner}.

Then, the contents of a given [focus navigation scope
owner](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner-3}
`owner`{.variable}\'s [focus navigation
scope](#focus-navigation-scope){#data-model:focus-navigation-scope} are
all elements whose [associated focus navigation
owner](#associated-focus-navigation-owner){#data-model:associated-focus-navigation-owner-2}
is `owner`{.variable}.

The order of elements within a [focus navigation
scope](#focus-navigation-scope){#data-model:focus-navigation-scope-2}
does not impact any of the algorithms in this specification. Ordering
only becomes important for the [tabindex-ordered focus navigation
scope](#tabindex-ordered-focus-navigation-scope){#data-model:tabindex-ordered-focus-navigation-scope}
and [flattened tabindex-ordered focus navigation
scope](#flattened-tabindex-ordered-focus-navigation-scope){#data-model:flattened-tabindex-ordered-focus-navigation-scope}
concepts defined below.

A [tabindex-ordered focus navigation
scope]{#tabindex-ordered-focus-navigation-scope .dfn} is a list of
[focusable areas](#focusable-area){#data-model:focusable-area-25} and
[focus navigation scope
owners](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner-4}.
Every [focus navigation scope
owner](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner-5}
`owner`{.variable} has [tabindex-ordered focus navigation
scope](#tabindex-ordered-focus-navigation-scope){#data-model:tabindex-ordered-focus-navigation-scope-2},
whose contents are determined as follows:

- It contains all elements in `owner`{.variable}\'s [focus navigation
  scope](#focus-navigation-scope){#data-model:focus-navigation-scope-3}
  that are themselves [focus navigation scope
  owners](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner-6},
  except the elements whose [tabindex
  value](#tabindex-value){#data-model:tabindex-value-4} is a negative
  integer.

- It contains all of the [focusable
  areas](#focusable-area){#data-model:focusable-area-26} whose [DOM
  anchor](#dom-anchor){#data-model:dom-anchor-12} is an element in
  `owner`{.variable}\'s [focus navigation
  scope](#focus-navigation-scope){#data-model:focus-navigation-scope-4},
  except the [focusable
  areas](#focusable-area){#data-model:focusable-area-27} whose [tabindex
  value](#tabindex-value){#data-model:tabindex-value-5} is a negative
  integer.

The order within a [tabindex-ordered focus navigation
scope](#tabindex-ordered-focus-navigation-scope){#data-model:tabindex-ordered-focus-navigation-scope-3}
is determined by each element\'s [tabindex
value](#tabindex-value){#data-model:tabindex-value-6}, as described in
the section below.

The rules there do not give a precise ordering, as they are composed
mostly of \"should\" statements and relative orderings.

A [flattened tabindex-ordered focus navigation
scope]{#flattened-tabindex-ordered-focus-navigation-scope .dfn} is a
list of [focusable
areas](#focusable-area){#data-model:focusable-area-28}. Every [focus
navigation scope
owner](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner-7}
`owner`{.variable} owns a distinct [flattened tabindex-ordered focus
navigation
scope](#flattened-tabindex-ordered-focus-navigation-scope){#data-model:flattened-tabindex-ordered-focus-navigation-scope-2},
whose contents are determined by the following algorithm:

1.  Let `result`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#data-model:list-clone
    x-internal="list-clone"} of `owner`{.variable}\'s [tabindex-ordered
    focus navigation
    scope](#tabindex-ordered-focus-navigation-scope){#data-model:tabindex-ordered-focus-navigation-scope-4}.

2.  For each `item`{.variable} of `result`{.variable}:

    1.  If `item`{.variable} is not a [focus navigation scope
        owner](#focus-navigation-scope-owner){#data-model:focus-navigation-scope-owner-8},
        then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#data-model:continue
        x-internal="continue"}.

    2.  If `item`{.variable} is not a [focusable
        area](#focusable-area){#data-model:focusable-area-29}, then
        replace `item`{.variable} with all of the items in
        `item`{.variable}\'s [flattened tabindex-ordered focus
        navigation
        scope](#flattened-tabindex-ordered-focus-navigation-scope){#data-model:flattened-tabindex-ordered-focus-navigation-scope-3}.

    3.  Otherwise, insert the contents of `item`{.variable}\'s
        [flattened tabindex-ordered focus navigation
        scope](#flattened-tabindex-ordered-focus-navigation-scope){#data-model:flattened-tabindex-ordered-focus-navigation-scope-4}
        after `item`{.variable}.

#### [6.6.3]{.secno} The [`tabindex`{#the-tabindex-attribute:attr-tabindex}](#attr-tabindex) attribute[](#the-tabindex-attribute){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/tabindex](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/tabindex "The tabindex global attribute allows developers to make HTML elements focusable, allow or prevent them from being sequentially focusable (usually with the Tab key, hence the name) and determine their relative ordering for sequential focus navigation.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`tabindex`]{#attr-tabindex .dfn dfn-for="htmlsvg-global"
dfn-type="element-attr"} content attribute allows authors to make an
element and regions that have the element as its [DOM
anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor} be [focusable
areas](#focusable-area){#the-tabindex-attribute:focusable-area}, allow
or prevent them from being [sequentially
focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable},
and determine their relative ordering for [sequential focus
navigation](#sequential-focus-navigation){#the-tabindex-attribute:sequential-focus-navigation}.

The name \"tab index\" comes from the common use of the [Tab]{.kbd} key
to navigate through the focusable elements. The term \"tabbing\" refers
to moving forward through [sequentially
focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-2}
[focusable
areas](#focusable-area){#the-tabindex-attribute:focusable-area-2}.

The
[`tabindex`{#the-tabindex-attribute:attr-tabindex-2}](#attr-tabindex)
attribute, if specified, must have a value that is a [valid
integer](common-microsyntaxes.html#valid-integer){#the-tabindex-attribute:valid-integer}.
Positive numbers specify the relative position of the element\'s
[focusable
areas](#focusable-area){#the-tabindex-attribute:focusable-area-3} in the
[sequential focus navigation
order](#sequential-focus-navigation-order){#the-tabindex-attribute:sequential-focus-navigation-order},
and negative numbers indicate that the control is not [sequentially
focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-3}.

Developers should use caution when using values other than 0 or −1 for
their
[`tabindex`{#the-tabindex-attribute:attr-tabindex-3}](#attr-tabindex)
attributes as this is complicated to do correctly.

::: note
The following provides a non-normative summary of the behaviors of the
possible
[`tabindex`{#the-tabindex-attribute:attr-tabindex-4}](#attr-tabindex)
attribute values. The below processing model gives the more precise
rules.

omitted (or non-integer values)
:   The user agent will decide whether the element is
    [focusable](#focusable){#the-tabindex-attribute:focusable}, and if
    it is, whether it is [sequentially
    focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-4}
    or [click
    focusable](#click-focusable){#the-tabindex-attribute:click-focusable}
    (or both).

−1 (or other negative integer values)
:   Causes the element to be
    [focusable](#focusable){#the-tabindex-attribute:focusable-2}, and
    indicates that the author would prefer the element to be [click
    focusable](#click-focusable){#the-tabindex-attribute:click-focusable-2}
    but not [sequentially
    focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-5}.
    The user agent might ignore this preference for click and sequential
    focusability, e.g., for specific element types according to platform
    conventions, or for keyboard-only users.

0
:   Causes the element to be
    [focusable](#focusable){#the-tabindex-attribute:focusable-3}, and
    indicates that the author would prefer the element to be both [click
    focusable](#click-focusable){#the-tabindex-attribute:click-focusable-3}
    and [sequentially
    focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-6}.
    The user agent might ignore this preference for click and sequential
    focusability.

positive integer values
:   Behaves the same as 0, but in addition creates a relative ordering
    within a [tabindex-ordered focus navigation
    scope](#tabindex-ordered-focus-navigation-scope){#the-tabindex-attribute:tabindex-ordered-focus-navigation-scope},
    so that elements with higher
    [`tabindex`{#the-tabindex-attribute:attr-tabindex-5}](#attr-tabindex)
    attribute value come later.

Note that the
[`tabindex`{#the-tabindex-attribute:attr-tabindex-6}](#attr-tabindex)
attribute cannot be used to make an element non-focusable. The only way
a page author can do that is by
[disabling](semantics-other.html#concept-element-disabled){#the-tabindex-attribute:concept-element-disabled}
the element, or making it
[inert](#inert){#the-tabindex-attribute:inert}.
:::

------------------------------------------------------------------------

The [tabindex value]{#tabindex-value .dfn} of an element is the value of
its
[`tabindex`{#the-tabindex-attribute:attr-tabindex-7}](#attr-tabindex)
attribute, parsed using the [rules for parsing
integers](common-microsyntaxes.html#rules-for-parsing-integers){#the-tabindex-attribute:rules-for-parsing-integers}.
If parsing fails or the attribute is not specified, then the [tabindex
value](#tabindex-value){#the-tabindex-attribute:tabindex-value} is null.

The [tabindex
value](#tabindex-value){#the-tabindex-attribute:tabindex-value-2} of a
[focusable
area](#focusable-area){#the-tabindex-attribute:focusable-area-4} is the
[tabindex
value](#tabindex-value){#the-tabindex-attribute:tabindex-value-3} of its
[DOM anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-2}.

The [tabindex
value](#tabindex-value){#the-tabindex-attribute:tabindex-value-4} of an
element must be interpreted as follows:

If the value is null

:   The user agent should follow platform conventions to determine if
    the element should be considered as a [focusable
    area](#focusable-area){#the-tabindex-attribute:focusable-area-5} and
    if so, whether the element and any [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-6}
    that have the element as their [DOM
    anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-3} are
    [sequentially
    focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-7},
    and if so, what their relative position in their [tabindex-ordered
    focus navigation
    scope](#tabindex-ordered-focus-navigation-scope){#the-tabindex-attribute:tabindex-ordered-focus-navigation-scope-2}
    is to be. If the element is a [focus navigation scope
    owner](#focus-navigation-scope-owner){#the-tabindex-attribute:focus-navigation-scope-owner},
    it must be included in its [tabindex-ordered focus navigation
    scope](#tabindex-ordered-focus-navigation-scope){#the-tabindex-attribute:tabindex-ordered-focus-navigation-scope-3}
    even if it is not a [focusable
    area](#focusable-area){#the-tabindex-attribute:focusable-area-7}.

    The relative ordering within a [tabindex-ordered focus navigation
    scope](#tabindex-ordered-focus-navigation-scope){#the-tabindex-attribute:tabindex-ordered-focus-navigation-scope-4}
    for elements and [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-8}
    that belong to the same [focus navigation
    scope](#focus-navigation-scope){#the-tabindex-attribute:focus-navigation-scope}
    and whose [tabindex
    value](#tabindex-value){#the-tabindex-attribute:tabindex-value-5} is
    null should be in [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#the-tabindex-attribute:shadow-including-tree-order
    x-internal="shadow-including-tree-order"}.

    Modulo platform conventions, it is suggested that the following
    elements should be considered as [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-9}
    and be [sequentially
    focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-8}:

    - [`a`{#the-tabindex-attribute:the-a-element}](text-level-semantics.html#the-a-element)
      elements that have an
      [`href`{#the-tabindex-attribute:attr-hyperlink-href}](links.html#attr-hyperlink-href)
      attribute
    - [`button`{#the-tabindex-attribute:the-button-element}](form-elements.html#the-button-element)
      elements
    - [`input`{#the-tabindex-attribute:the-input-element}](input.html#the-input-element)
      elements whose
      [`type`{#the-tabindex-attribute:attr-input-type}](input.html#attr-input-type)
      attribute are not in the
      [Hidden](input.html#hidden-state-(type=hidden)){#the-tabindex-attribute:hidden-state-(type=hidden)}
      state
    - [`select`{#the-tabindex-attribute:the-select-element}](form-elements.html#the-select-element)
      elements
    - [`textarea`{#the-tabindex-attribute:the-textarea-element}](form-elements.html#the-textarea-element)
      elements
    - [`summary`{#the-tabindex-attribute:the-summary-element}](interactive-elements.html#the-summary-element)
      elements that are the first
      [`summary`{#the-tabindex-attribute:the-summary-element-2}](interactive-elements.html#the-summary-element)
      element child of a
      [`details`{#the-tabindex-attribute:the-details-element}](interactive-elements.html#the-details-element)
      element
    - Elements with a
      [`draggable`{#the-tabindex-attribute:attr-draggable}](dnd.html#attr-draggable)
      attribute set, if that would enable the user agent to allow the
      user to begin drag operations for those elements without the use
      of a pointing device
    - [Editing
      hosts](#editing-host){#the-tabindex-attribute:editing-host}
    - [Navigable
      containers](document-sequences.html#navigable-container){#the-tabindex-attribute:navigable-container}

If the value is a negative integer

:   The user agent must consider the element as a [focusable
    area](#focusable-area){#the-tabindex-attribute:focusable-area-10},
    but should omit the element from any [tabindex-ordered focus
    navigation
    scope](#tabindex-ordered-focus-navigation-scope){#the-tabindex-attribute:tabindex-ordered-focus-navigation-scope-5}.

    One valid reason to ignore the requirement that sequential focus
    navigation not allow the author to lead to the element would be if
    the user\'s only mechanism for moving the focus is sequential focus
    navigation. For instance, a keyboard-only user would be unable to
    click on a text control with a negative
    [`tabindex`{#the-tabindex-attribute:attr-tabindex-8}](#attr-tabindex),
    so that user\'s user agent would be well justified in allowing the
    user to tab to the control regardless.

If the value is a zero

:   The user agent must allow the element to be considered as a
    [focusable
    area](#focusable-area){#the-tabindex-attribute:focusable-area-11}
    and should allow the element and any [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-12}
    that have the element as their [DOM
    anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-4} to be
    [sequentially
    focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-9}.

    The relative ordering within a [tabindex-ordered focus navigation
    scope](#tabindex-ordered-focus-navigation-scope){#the-tabindex-attribute:tabindex-ordered-focus-navigation-scope-6}
    for elements and [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-13}
    that belong to the same [focus navigation
    scope](#focus-navigation-scope){#the-tabindex-attribute:focus-navigation-scope-2}
    and whose [tabindex
    value](#tabindex-value){#the-tabindex-attribute:tabindex-value-6} is
    zero should be in [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#the-tabindex-attribute:shadow-including-tree-order-2
    x-internal="shadow-including-tree-order"}.

If the value is greater than zero

:   The user agent must allow the element to be considered as a
    [focusable
    area](#focusable-area){#the-tabindex-attribute:focusable-area-14}
    and should allow the element and any [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-15}
    that have the element as their [DOM
    anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-5} to be
    [sequentially
    focusable](#sequentially-focusable){#the-tabindex-attribute:sequentially-focusable-10},
    and should place the element --- referenced as
    `candidate`{.variable} below --- and the aforementioned [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-16}
    in the [tabindex-ordered focus navigation
    scope](#tabindex-ordered-focus-navigation-scope){#the-tabindex-attribute:tabindex-ordered-focus-navigation-scope-7}
    where the element is a part of so that, relative to other elements
    and [focusable
    areas](#focusable-area){#the-tabindex-attribute:focusable-area-17}
    that belong to the same [focus navigation
    scope](#focus-navigation-scope){#the-tabindex-attribute:focus-navigation-scope-3},
    they are:

    - before any [focusable
      area](#focusable-area){#the-tabindex-attribute:focusable-area-18}
      whose [DOM
      anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-6} is an
      element whose
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-9}](#attr-tabindex)
      attribute has been omitted or whose value, when parsed, returns an
      error,
    - before any [focusable
      area](#focusable-area){#the-tabindex-attribute:focusable-area-19}
      whose [DOM
      anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-7} is an
      element whose
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-10}](#attr-tabindex)
      attribute has a value less than or equal to zero,
    - after any [focusable
      area](#focusable-area){#the-tabindex-attribute:focusable-area-20}
      whose [DOM
      anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-8} is an
      element whose
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-11}](#attr-tabindex)
      attribute has a value greater than zero but less than the value of
      the
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-12}](#attr-tabindex)
      attribute on `candidate`{.variable},
    - after any [focusable
      area](#focusable-area){#the-tabindex-attribute:focusable-area-21}
      whose [DOM
      anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-9} is an
      element whose
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-13}](#attr-tabindex)
      attribute has a value equal to the value of the
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-14}](#attr-tabindex)
      attribute on `candidate`{.variable} but that is located earlier
      than `candidate`{.variable} in [shadow-including tree
      order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#the-tabindex-attribute:shadow-including-tree-order-3
      x-internal="shadow-including-tree-order"},
    - before any [focusable
      area](#focusable-area){#the-tabindex-attribute:focusable-area-22}
      whose [DOM
      anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-10} is an
      element whose
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-15}](#attr-tabindex)
      attribute has a value equal to the value of the
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-16}](#attr-tabindex)
      attribute on `candidate`{.variable} but that is located later than
      `candidate`{.variable} in [shadow-including tree
      order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#the-tabindex-attribute:shadow-including-tree-order-4
      x-internal="shadow-including-tree-order"}, and
    - before any [focusable
      area](#focusable-area){#the-tabindex-attribute:focusable-area-23}
      whose [DOM
      anchor](#dom-anchor){#the-tabindex-attribute:dom-anchor-11} is an
      element whose
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-17}](#attr-tabindex)
      attribute has a value greater than the value of the
      [`tabindex`{#the-tabindex-attribute:attr-tabindex-18}](#attr-tabindex)
      attribute on `candidate`{.variable}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/tabIndex](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/tabIndex "The tabIndex property of the HTMLElement interface represents the tab order of the current element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet Explorer[🔰
5.5+]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`tabIndex`]{#dom-tabindex .dfn dfn-for="HTMLOrSVGElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-tabindex-attribute:reflect}
the value of the
[`tabindex`{#the-tabindex-attribute:attr-tabindex-19}](#attr-tabindex)
content attribute. The [default
value](common-dom-interfaces.html#default-value){#the-tabindex-attribute:default-value}
is 0 if the element is an
[`a`{#the-tabindex-attribute:the-a-element-2}](text-level-semantics.html#the-a-element),
[`area`{#the-tabindex-attribute:the-area-element}](image-maps.html#the-area-element),
[`button`{#the-tabindex-attribute:the-button-element-2}](form-elements.html#the-button-element),
[`frame`{#the-tabindex-attribute:frame}](obsolete.html#frame),
[`iframe`{#the-tabindex-attribute:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
[`input`{#the-tabindex-attribute:the-input-element-2}](input.html#the-input-element),
[`object`{#the-tabindex-attribute:the-object-element}](iframe-embed-object.html#the-object-element),
[`select`{#the-tabindex-attribute:the-select-element-2}](form-elements.html#the-select-element),
[`textarea`{#the-tabindex-attribute:the-textarea-element-2}](form-elements.html#the-textarea-element),
or [SVG
`a`](https://svgwg.org/svg2-draft/linking.html#AElement){#the-tabindex-attribute:svg-a
x-internal="svg-a"} element, or is a
[`summary`{#the-tabindex-attribute:the-summary-element-3}](interactive-elements.html#the-summary-element)
element that is a [summary for its parent
details](interactive-elements.html#summary-for-its-parent-details){#the-tabindex-attribute:summary-for-its-parent-details}.
The [default
value](common-dom-interfaces.html#default-value){#the-tabindex-attribute:default-value-2}
is −1 otherwise.

The varying default value based on element type is a historical
artifact.

#### [6.6.4]{.secno} []{#processing-model-5}Processing model[](#focus-processing-model){.self-link} {#focus-processing-model}

To [get the focusable area]{#get-the-focusable-area .dfn} for a
`focus target`{.variable} that is either an element that is not a
[focusable
area](#focusable-area){#focus-processing-model:focusable-area}, or is a
[navigable](document-sequences.html#navigable){#focus-processing-model:navigable},
given an optional string `focus trigger`{.variable} (default
\"`other`\"), run the first matching set of steps from the following
list:

If `focus target`{.variable} is an [`area`{#focus-processing-model:the-area-element}](image-maps.html#the-area-element) element with one or more shapes that are [focusable areas](#focusable-area){#focus-processing-model:focusable-area-2}
:   Return the shape corresponding to the first
    [`img`{#focus-processing-model:the-img-element}](embedded-content.html#the-img-element)
    element in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#focus-processing-model:tree-order
    x-internal="tree-order"} that uses the image map to which the
    [`area`{#focus-processing-model:the-area-element-2}](image-maps.html#the-area-element)
    element belongs.

If `focus target`{.variable} is an element with one or more scrollable regions that are [focusable areas](#focusable-area){#focus-processing-model:focusable-area-3}
:   Return the element\'s first scrollable region, according to a
    pre-order, depth-first traversal of the [flat
    tree](https://drafts.csswg.org/css-scoping/#flat-tree){#focus-processing-model:flat-tree
    x-internal="flat-tree"}.
    [\[CSSSCOPING\]](references.html#refsCSSSCOPING)

If `focus target`{.variable} is the [document element](https://dom.spec.whatwg.org/#document-element){#focus-processing-model:document-element x-internal="document-element"} of its [`Document`{#focus-processing-model:document}](dom.html#document)
:   Return the
    [`Document`{#focus-processing-model:document-2}](dom.html#document)\'s
    [viewport](https://drafts.csswg.org/css2/#viewport){#focus-processing-model:viewport
    x-internal="viewport"}.

If `focus target`{.variable} is a [navigable](document-sequences.html#navigable){#focus-processing-model:navigable-2}
:   Return the
    [navigable](document-sequences.html#navigable){#focus-processing-model:navigable-3}\'s
    [active
    document](document-sequences.html#nav-document){#focus-processing-model:nav-document}.

If `focus target`{.variable} is a [navigable container](document-sequences.html#navigable-container){#focus-processing-model:navigable-container} with a non-null [content navigable](document-sequences.html#content-navigable){#focus-processing-model:content-navigable}
:   Return the [navigable
    container](document-sequences.html#navigable-container){#focus-processing-model:navigable-container-2}\'s
    [content
    navigable](document-sequences.html#content-navigable){#focus-processing-model:content-navigable-2}\'s
    [active
    document](document-sequences.html#nav-document){#focus-processing-model:nav-document-2}.

If `focus target`{.variable} is a [shadow host](https://dom.spec.whatwg.org/#element-shadow-host){#focus-processing-model:shadow-host x-internal="shadow-host"} whose [shadow root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#focus-processing-model:concept-element-shadow-root x-internal="concept-element-shadow-root"}\'s [delegates focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus){#focus-processing-model:delegates-focus x-internal="delegates-focus"} is true

:   1.  Let `focusedElement`{.variable} be the [currently focused area
        of a top-level
        traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable}\'s
        [DOM anchor](#dom-anchor){#focus-processing-model:dom-anchor}.

    2.  If `focus target`{.variable} is a [shadow-including inclusive
        ancestor](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-ancestor){#focus-processing-model:shadow-including-inclusive-ancestor
        x-internal="shadow-including-inclusive-ancestor"} of
        `focusedElement`{.variable}, then return
        `focusedElement`{.variable}.

    3.  Return the [focus
        delegate](#focus-delegate){#focus-processing-model:focus-delegate}
        for `focus target`{.variable} given `focus trigger`{.variable}.

    For [sequential
    focusability](#sequentially-focusable){#focus-processing-model:sequentially-focusable},
    the handling of [shadow
    hosts](https://dom.spec.whatwg.org/#element-shadow-host){#focus-processing-model:shadow-host-2
    x-internal="shadow-host"} and [delegates
    focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus){#focus-processing-model:delegates-focus-2
    x-internal="delegates-focus"} is done when constructing the
    [sequential focus navigation
    order](#sequential-focus-navigation-order){#focus-processing-model:sequential-focus-navigation-order}.
    That is, the [focusing
    steps](#focusing-steps){#focus-processing-model:focusing-steps} will
    never be called on such [shadow
    hosts](https://dom.spec.whatwg.org/#element-shadow-host){#focus-processing-model:shadow-host-3
    x-internal="shadow-host"} as part of sequential focus navigation.

Otherwise

:   Return null.

The [focus delegate]{#focus-delegate .dfn} for a
`focusTarget`{.variable}, given an optional string
`focusTrigger`{.variable} (default \"`other`\"), is given by the
following steps:

1.  If `focusTarget`{.variable} is a [shadow
    host](https://dom.spec.whatwg.org/#element-shadow-host){#focus-processing-model:shadow-host-4
    x-internal="shadow-host"} and its [shadow
    root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#focus-processing-model:concept-element-shadow-root-2
    x-internal="concept-element-shadow-root"}\'s [delegates
    focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus){#focus-processing-model:delegates-focus-3
    x-internal="delegates-focus"} is false, then return null.

2.  Let `whereToLook`{.variable} be `focusTarget`{.variable}.

3.  If `whereToLook`{.variable} is a [shadow
    host](https://dom.spec.whatwg.org/#element-shadow-host){#focus-processing-model:shadow-host-5
    x-internal="shadow-host"}, then set `whereToLook`{.variable} to
    `whereToLook`{.variable}\'s [shadow
    root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#focus-processing-model:concept-element-shadow-root-3
    x-internal="concept-element-shadow-root"}.

4.  Let `autofocusDelegate`{.variable} be the [autofocus
    delegate](#autofocus-delegate){#focus-processing-model:autofocus-delegate}
    for `whereToLook`{.variable} given `focusTrigger`{.variable}.

5.  If `autofocusDelegate`{.variable} is not null, then return
    `autofocusDelegate`{.variable}.

6.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#focus-processing-model:list-iterate
    x-internal="list-iterate"} `descendant`{.variable} of
    `whereToLook`{.variable}\'s
    [descendants](https://dom.spec.whatwg.org/#concept-tree-descendant){#focus-processing-model:descendant
    x-internal="descendant"}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#focus-processing-model:tree-order-2
    x-internal="tree-order"}:

    1.  Let `focusableArea`{.variable} be null.

    2.  If `focusTarget`{.variable} is a
        [`dialog`{#focus-processing-model:the-dialog-element}](interactive-elements.html#the-dialog-element)
        element and `descendant`{.variable} is [sequentially
        focusable](#sequentially-focusable){#focus-processing-model:sequentially-focusable-2},
        then set `focusableArea`{.variable} to `descendant`{.variable}.

    3.  Otherwise, if `focusTarget`{.variable} is not a
        [`dialog`{#focus-processing-model:the-dialog-element-2}](interactive-elements.html#the-dialog-element)
        and `descendant`{.variable} is a [focusable
        area](#focusable-area){#focus-processing-model:focusable-area-4},
        set `focusableArea`{.variable} to `descendant`{.variable}.

    4.  Otherwise, set `focusableArea`{.variable} to the result of
        [getting the focusable
        area](#get-the-focusable-area){#focus-processing-model:get-the-focusable-area}
        for `descendant`{.variable} given `focusTrigger`{.variable}.

        This step can end up recursing, i.e., the [get the focusable
        area](#get-the-focusable-area){#focus-processing-model:get-the-focusable-area-2}
        steps might return the [focus
        delegate](#focus-delegate){#focus-processing-model:focus-delegate-2}
        of `descendant`{.variable}.

    5.  If `focusableArea`{.variable} is not null, then return
        `focusableArea`{.variable}.

    It\'s important that we are *not* looking at the [shadow-including
    descendants](https://dom.spec.whatwg.org/#concept-shadow-including-descendant){#focus-processing-model:shadow-including-descendant
    x-internal="shadow-including-descendant"} here, but instead only at
    the
    [descendants](https://dom.spec.whatwg.org/#concept-tree-descendant){#focus-processing-model:descendant-2
    x-internal="descendant"}. [Shadow
    hosts](https://dom.spec.whatwg.org/#element-shadow-host){#focus-processing-model:shadow-host-6
    x-internal="shadow-host"} are instead handled by the recursive case
    mentioned above.

7.  Return null.

The above algorithm essentially returns the first suitable [focusable
area](#focusable-area){#focus-processing-model:focusable-area-5} where
the path between its [DOM
anchor](#dom-anchor){#focus-processing-model:dom-anchor-2} and
`focusTarget`{.variable} delegates focus at any shadow tree boundaries.

The [autofocus delegate]{#autofocus-delegate .dfn} for a
`focus target`{.variable} given a `focus trigger`{.variable} is given by
the following steps:

1.  For each
    [descendant](https://dom.spec.whatwg.org/#concept-tree-descendant){#focus-processing-model:descendant-3
    x-internal="descendant"} `descendant`{.variable} of
    `focus target`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#focus-processing-model:tree-order-3
    x-internal="tree-order"}:

    1.  If `descendant`{.variable} does not have an
        [`autofocus`{#focus-processing-model:attr-fe-autofocus}](#attr-fe-autofocus)
        content attribute, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#focus-processing-model:continue
        x-internal="continue"}.

    2.  Let `focusable area`{.variable} be `descendant`{.variable}, if
        `descendant`{.variable} is a [focusable
        area](#focusable-area){#focus-processing-model:focusable-area-6};
        otherwise let `focusable area`{.variable} be the result of
        [getting the focusable
        area](#get-the-focusable-area){#focus-processing-model:get-the-focusable-area-3}
        for `descendant`{.variable} given `focus trigger`{.variable}.

    3.  If `focusable area`{.variable} is null, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#focus-processing-model:continue-2
        x-internal="continue"}.

    4.  If `focusable area`{.variable} is not [click
        focusable](#click-focusable){#focus-processing-model:click-focusable}
        and `focus trigger`{.variable} is \"`click`\", then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#focus-processing-model:continue-3
        x-internal="continue"}.

    5.  Return `focusable area`{.variable}.

2.  Return null.

The [focusing steps]{#focusing-steps .dfn export=""} for an object
`new focus target`{.variable} that is either a [focusable
area](#focusable-area){#focus-processing-model:focusable-area-7}, or an
element that is not a [focusable
area](#focusable-area){#focus-processing-model:focusable-area-8}, or a
[navigable](document-sequences.html#navigable){#focus-processing-model:navigable-4},
are as follows. They can optionally be run with a
`fallback target`{.variable} and a string `focus trigger`{.variable}.

1.  If `new focus target`{.variable} is not a [focusable
    area](#focusable-area){#focus-processing-model:focusable-area-9},
    then set `new focus target`{.variable} to the result of [getting the
    focusable
    area](#get-the-focusable-area){#focus-processing-model:get-the-focusable-area-4}
    for `new focus target`{.variable}, given `focus trigger`{.variable}
    if it was passed.

2.  If `new focus target`{.variable} is null, then:

    1.  If no `fallback target`{.variable} was specified, then return.

    2.  Otherwise, set `new focus target`{.variable} to the
        `fallback target`{.variable}.

3.  If `new focus target`{.variable} is a [navigable
    container](document-sequences.html#navigable-container){#focus-processing-model:navigable-container-3}
    with non-null [content
    navigable](document-sequences.html#content-navigable){#focus-processing-model:content-navigable-3},
    then set `new focus target`{.variable} to the [content
    navigable](document-sequences.html#content-navigable){#focus-processing-model:content-navigable-4}\'s
    [active
    document](document-sequences.html#nav-document){#focus-processing-model:nav-document-3}.

4.  If `new focus target`{.variable} is a [focusable
    area](#focusable-area){#focus-processing-model:focusable-area-10}
    and its [DOM
    anchor](#dom-anchor){#focus-processing-model:dom-anchor-3} is
    [inert](#inert){#focus-processing-model:inert}, then return.

5.  If `new focus target`{.variable} is the [currently focused area of a
    top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-2},
    then return.

6.  Let `old chain`{.variable} be the [current focus chain of the
    top-level
    traversable](#current-focus-chain-of-a-top-level-traversable){#focus-processing-model:current-focus-chain-of-a-top-level-traversable}
    in which `new focus target`{.variable} finds itself.

7.  Let `new chain`{.variable} be the [focus
    chain](#focus-chain){#focus-processing-model:focus-chain} of
    `new focus target`{.variable}.

8.  Run the [focus update
    steps](#focus-update-steps){#focus-processing-model:focus-update-steps}
    with `old chain`{.variable}, `new chain`{.variable}, and
    `new focus target`{.variable} respectively.

User agents must
[immediately](infrastructure.html#immediately){#focus-processing-model:immediately}
run the [focusing
steps](#focusing-steps){#focus-processing-model:focusing-steps-2} for a
[focusable
area](#focusable-area){#focus-processing-model:focusable-area-11} or
[navigable](document-sequences.html#navigable){#focus-processing-model:navigable-5}
`candidate`{.variable} whenever the user attempts to move the focus to
`candidate`{.variable}.

The [unfocusing steps]{#unfocusing-steps .dfn} for an object
`old focus target`{.variable} that is either a [focusable
area](#focusable-area){#focus-processing-model:focusable-area-12} or an
element that is not a [focusable
area](#focusable-area){#focus-processing-model:focusable-area-13} are as
follows:

1.  If `old focus target`{.variable} is a [shadow
    host](https://dom.spec.whatwg.org/#element-shadow-host){#focus-processing-model:shadow-host-7
    x-internal="shadow-host"} whose [shadow
    root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#focus-processing-model:concept-element-shadow-root-4
    x-internal="concept-element-shadow-root"}\'s [delegates
    focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus){#focus-processing-model:delegates-focus-4
    x-internal="delegates-focus"} is true, and
    `old focus target`{.variable}\'s [shadow
    root](https://dom.spec.whatwg.org/#concept-element-shadow-root){#focus-processing-model:concept-element-shadow-root-5
    x-internal="concept-element-shadow-root"} is a [shadow-including
    inclusive
    ancestor](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-ancestor){#focus-processing-model:shadow-including-inclusive-ancestor-2
    x-internal="shadow-including-inclusive-ancestor"} of the [currently
    focused area of a top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-3}\'s
    [DOM anchor](#dom-anchor){#focus-processing-model:dom-anchor-4},
    then set `old focus target`{.variable} to that [currently focused
    area of a top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-4}.

2.  If `old focus target`{.variable} is
    [inert](#inert){#focus-processing-model:inert-2}, then return.

3.  If `old focus target`{.variable} is an
    [`area`{#focus-processing-model:the-area-element-3}](image-maps.html#the-area-element)
    element and one of its shapes is the [currently focused area of a
    top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-5},
    or, if `old focus target`{.variable} is an element with one or more
    scrollable regions, and one of them is the [currently focused area
    of a top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-6},
    then let `old focus target`{.variable} be that [currently focused
    area of a top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-7}.

4.  Let `old chain`{.variable} be the [current focus chain of the
    top-level
    traversable](#current-focus-chain-of-a-top-level-traversable){#focus-processing-model:current-focus-chain-of-a-top-level-traversable-2}
    in which `old focus target`{.variable} finds itself.

5.  If `old focus target`{.variable} is not one of the entries in
    `old chain`{.variable}, then return.

6.  If `old focus target`{.variable} is not a [focusable
    area](#focusable-area){#focus-processing-model:focusable-area-14},
    then return.

7.  Let `topDocument`{.variable} be `old chain`{.variable}\'s last
    entry.

8.  If `topDocument`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#focus-processing-model:node-navigable}
    has [system
    focus](#system-focus){#focus-processing-model:system-focus}, then
    run the [focusing
    steps](#focusing-steps){#focus-processing-model:focusing-steps-3}
    for `topDocument`{.variable}\'s
    [viewport](https://drafts.csswg.org/css2/#viewport){#focus-processing-model:viewport-2
    x-internal="viewport"}.

    Otherwise, apply any relevant platform-specific conventions for
    removing [system
    focus](#system-focus){#focus-processing-model:system-focus-2} from
    `topDocument`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#focus-processing-model:node-navigable-2},
    and run the [focus update
    steps](#focus-update-steps){#focus-processing-model:focus-update-steps-2}
    given `old chain`{.variable}, an empty list, and null.

The [unfocusing
steps](#unfocusing-steps){#focus-processing-model:unfocusing-steps} do
not always result in the focus changing, even when applied to the
[currently focused area of a top-level
traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-8}.
For example, if the [currently focused area of a top-level
traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-9}
is a
[viewport](https://drafts.csswg.org/css2/#viewport){#focus-processing-model:viewport-3
x-internal="viewport"}, then it will usually keep its focus regardless
until another [focusable
area](#focusable-area){#focus-processing-model:focusable-area-15} is
explicitly focused with the [focusing
steps](#focusing-steps){#focus-processing-model:focusing-steps-4}.

------------------------------------------------------------------------

The [focus update steps]{#focus-update-steps .dfn}, given an
`old chain`{.variable}, a `new chain`{.variable}, and a
`new focus target`{.variable} respectively, are as follows:

1.  If the last entry in `old chain`{.variable} and the last entry in
    `new chain`{.variable} are the same, pop the last entry from
    `old chain`{.variable} and the last entry from
    `new chain`{.variable} and redo this step.

2.  For each entry `entry`{.variable} in `old chain`{.variable}, in
    order, run these substeps:

    1.  If `entry`{.variable} is an
        [`input`{#focus-processing-model:the-input-element}](input.html#the-input-element)
        element, and the
        [`change`{#focus-processing-model:event-change}](indices.html#event-change)
        event
        [applies](input.html#concept-input-apply){#focus-processing-model:concept-input-apply}
        to the element, and the element does not have a defined
        [activation
        behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#focus-processing-model:activation-behaviour
        x-internal="activation-behaviour"}, and the user has changed the
        element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#focus-processing-model:concept-fe-value}
        or its list of [selected
        files](input.html#concept-input-type-file-selected){#focus-processing-model:concept-input-type-file-selected}
        while the control was focused without committing that change
        (such that it is different to what it was when the control was
        first focused), then:

        1.  Set `entry`{.variable}\'s [user
            validity](form-control-infrastructure.html#user-validity){#focus-processing-model:user-validity}
            to true.

        2.  ::: {#unfocus-causes-change-event}
            [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#focus-processing-model:concept-event-fire
            x-internal="concept-event-fire"} named
            [`change`{#focus-processing-model:event-change-2}](indices.html#event-change)
            at the element, with the
            [`bubbles`{#focus-processing-model:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
            attribute initialized to true.
            :::

    2.  If `entry`{.variable} is an element, let
        `blur event target`{.variable} be `entry`{.variable}.

        If `entry`{.variable} is a
        [`Document`{#focus-processing-model:document-3}](dom.html#document)
        object, let `blur event target`{.variable} be that
        [`Document`{#focus-processing-model:document-4}](dom.html#document)
        object\'s [relevant global
        object](webappapis.html#concept-relevant-global){#focus-processing-model:concept-relevant-global}.

        Otherwise, let `blur event target`{.variable} be null.

    3.  If `entry`{.variable} is the last entry in
        `old chain`{.variable}, and `entry`{.variable} is an
        [`Element`{#focus-processing-model:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
        and the last entry in `new chain`{.variable} is also an
        [`Element`{#focus-processing-model:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
        then let `related blur target`{.variable} be the last entry in
        `new chain`{.variable}. Otherwise, let
        `related blur target`{.variable} be null.

    4.  If `blur event target`{.variable} is not null, [fire a focus
        event](#fire-a-focus-event){#focus-processing-model:fire-a-focus-event}
        named
        [`blur`{#focus-processing-model:event-blur}](indices.html#event-blur)
        at `blur event target`{.variable}, with
        `related blur target`{.variable} as the related target.

        In some cases, e.g., if `entry`{.variable} is an
        [`area`{#focus-processing-model:the-area-element-4}](image-maps.html#the-area-element)
        element\'s shape, a scrollable region, or a
        [viewport](https://drafts.csswg.org/css2/#viewport){#focus-processing-model:viewport-4
        x-internal="viewport"}, no event is fired.

3.  Apply any relevant platform-specific conventions for focusing
    `new focus target`{.variable}. (For example, some platforms select
    the contents of a text control when that control is focused.)

4.  For each entry `entry`{.variable} in `new chain`{.variable}, in
    reverse order, run these substeps:

    1.  If `entry`{.variable} is a [focusable
        area](#focusable-area){#focus-processing-model:focusable-area-16},
        and the [focused area of the
        document](#focused-area-of-the-document){#focus-processing-model:focused-area-of-the-document}
        is not `entry`{.variable}:

        1.  Set `document`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#focus-processing-model:concept-relevant-global-2}\'s
            [navigation
            API](nav-history-apis.html#window-navigation-api){#focus-processing-model:window-navigation-api}\'s
            [focus changed during ongoing
            navigation](nav-history-apis.html#focus-changed-during-ongoing-navigation){#focus-processing-model:focus-changed-during-ongoing-navigation}
            to true.

        2.  Designate `entry`{.variable} as the [focused area of the
            document](#focused-area-of-the-document){#focus-processing-model:focused-area-of-the-document-2}.

    2.  If `entry`{.variable} is an element, let
        `focus event target`{.variable} be `entry`{.variable}.

        If `entry`{.variable} is a
        [`Document`{#focus-processing-model:document-5}](dom.html#document)
        object, let `focus event target`{.variable} be that
        [`Document`{#focus-processing-model:document-6}](dom.html#document)
        object\'s [relevant global
        object](webappapis.html#concept-relevant-global){#focus-processing-model:concept-relevant-global-3}.

        Otherwise, let `focus event target`{.variable} be null.

    3.  If `entry`{.variable} is the last entry in
        `new chain`{.variable}, and `entry`{.variable} is an
        [`Element`{#focus-processing-model:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
        and the last entry in `old chain`{.variable} is also an
        [`Element`{#focus-processing-model:element-4}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"},
        then let `related focus target`{.variable} be the last entry in
        `old chain`{.variable}. Otherwise, let
        `related focus target`{.variable} be null.

    4.  If `focus event target`{.variable} is not null, [fire a focus
        event](#fire-a-focus-event){#focus-processing-model:fire-a-focus-event-2}
        named
        [`focus`{#focus-processing-model:event-focus}](indices.html#event-focus)
        at `focus event target`{.variable}, with
        `related focus target`{.variable} as the related target.

        In some cases, e.g. if `entry`{.variable} is an
        [`area`{#focus-processing-model:the-area-element-5}](image-maps.html#the-area-element)
        element\'s shape, a scrollable region, or a
        [viewport](https://drafts.csswg.org/css2/#viewport){#focus-processing-model:viewport-5
        x-internal="viewport"}, no event is fired.

To [fire a focus event]{#fire-a-focus-event .dfn} named `e`{.variable}
at an element `t`{.variable} with a given related target `r`{.variable},
[fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#focus-processing-model:concept-event-fire-2
x-internal="concept-event-fire"} named `e`{.variable} at `t`{.variable},
using
[`FocusEvent`{#focus-processing-model:focusevent}](https://w3c.github.io/uievents/#focusevent){x-internal="focusevent"},
with the
[`relatedTarget`{#focus-processing-model:dom-focusevent-relatedtarget}](https://w3c.github.io/uievents/#dom-focusevent-relatedtarget){x-internal="dom-focusevent-relatedtarget"}
attribute initialized to `r`{.variable}, the
[`view`{#focus-processing-model:dom-uievent-view}](https://w3c.github.io/uievents/#dom-uievent-view){x-internal="dom-uievent-view"}
attribute initialized to `t`{.variable}\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#focus-processing-model:node-document
x-internal="node-document"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#focus-processing-model:concept-relevant-global-4},
and the [composed
flag](https://dom.spec.whatwg.org/#composed-flag){#focus-processing-model:composed-flag
x-internal="composed-flag"} set.

------------------------------------------------------------------------

When a key event is to be routed in a [top-level
traversable](document-sequences.html#top-level-traversable){#focus-processing-model:top-level-traversable},
the user agent must run the following steps:

1.  Let `target area`{.variable} be the [currently focused area of the
    top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#focus-processing-model:currently-focused-area-of-a-top-level-traversable-10}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#focus-processing-model:assert
    x-internal="assert"}: `target area`{.variable} is not null, since
    key events are only routed to [top-level
    traversables](document-sequences.html#top-level-traversable){#focus-processing-model:top-level-traversable-2}
    that have [system
    focus](#system-focus){#focus-processing-model:system-focus-3}.
    Therefore, `target area`{.variable} is a [focusable
    area](#focusable-area){#focus-processing-model:focusable-area-17}.

3.  Let `target node`{.variable} be `target area`{.variable}\'s [DOM
    anchor](#dom-anchor){#focus-processing-model:dom-anchor-5}.

4.  If `target node`{.variable} is a
    [`Document`{#focus-processing-model:document-7}](dom.html#document)
    that has a [body
    element](dom.html#the-body-element-2){#focus-processing-model:the-body-element-2},
    then let `target node`{.variable} be [the body
    element](dom.html#the-body-element-2){#focus-processing-model:the-body-element-2-2}
    of that
    [`Document`{#focus-processing-model:document-8}](dom.html#document).

    Otherwise, if `target node`{.variable} is a
    [`Document`{#focus-processing-model:document-9}](dom.html#document)
    object that has a non-null [document
    element](https://dom.spec.whatwg.org/#document-element){#focus-processing-model:document-element-2
    x-internal="document-element"}, then let `target node`{.variable} be
    that [document
    element](https://dom.spec.whatwg.org/#document-element){#focus-processing-model:document-element-3
    x-internal="document-element"}.

5.  If `target node`{.variable} is not
    [inert](#inert){#focus-processing-model:inert-3}, then:

    1.  Let `canHandle`{.variable} be the result of
        [dispatching](https://dom.spec.whatwg.org/#concept-event-dispatch){#focus-processing-model:concept-event-dispatch
        x-internal="concept-event-dispatch"} the key event at
        `target node`{.variable}.

    2.  If `canHandle`{.variable} is true, then let
        `target area`{.variable} handle the key event. This might
        include [firing a `click`
        event](webappapis.html#fire-a-click-event){#focus-processing-model:fire-a-click-event}
        at `target node`{.variable}.

------------------------------------------------------------------------

The [has focus steps]{#has-focus-steps .dfn export=""}, given a
[`Document`{#focus-processing-model:document-10}](dom.html#document)
object `target`{.variable}, are as follows:

1.  If `target`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#focus-processing-model:node-navigable-3}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#focus-processing-model:nav-top}
    does not have [system
    focus](#system-focus){#focus-processing-model:system-focus-4}, then
    return false.

2.  Let `candidate`{.variable} be `target`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#focus-processing-model:node-navigable-4}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#focus-processing-model:nav-top-2}\'s
    [active
    document](document-sequences.html#active-document){#focus-processing-model:active-document}.

3.  While true:

    1.  If `candidate`{.variable} is `target`{.variable}, then return
        true.

    2.  If the [focused
        area](#focused-area-of-the-document){#focus-processing-model:focused-area-of-the-document-3}
        of `candidate`{.variable} is a [navigable
        container](document-sequences.html#navigable-container){#focus-processing-model:navigable-container-4}
        with a non-null [content
        navigable](document-sequences.html#content-navigable){#focus-processing-model:content-navigable-5},
        then set `candidate`{.variable} to the [active
        document](document-sequences.html#nav-document){#focus-processing-model:nav-document-4}
        of that [navigable
        container](document-sequences.html#navigable-container){#focus-processing-model:navigable-container-5}\'s
        [content
        navigable](document-sequences.html#content-navigable){#focus-processing-model:content-navigable-6}.

    3.  Otherwise, return false.

#### [6.6.5]{.secno} [Sequential focus navigation]{.dfn}[](#sequential-focus-navigation){.self-link}

Each
[`Document`{#sequential-focus-navigation:document}](dom.html#document)
has a [sequential focus navigation
order]{#sequential-focus-navigation-order .dfn}, which orders some or
all of the [focusable
areas](#focusable-area){#sequential-focus-navigation:focusable-area} in
the
[`Document`{#sequential-focus-navigation:document-2}](dom.html#document)
relative to each other. Its contents and ordering are given by the
[flattened tabindex-ordered focus navigation
scope](#flattened-tabindex-ordered-focus-navigation-scope){#sequential-focus-navigation:flattened-tabindex-ordered-focus-navigation-scope}
of the
[`Document`{#sequential-focus-navigation:document-3}](dom.html#document).

Per the rules defining the [flattened tabindex-ordered focus navigation
scope](#flattened-tabindex-ordered-focus-navigation-scope){#sequential-focus-navigation:flattened-tabindex-ordered-focus-navigation-scope-2},
the ordering is not necessarily related to the [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#sequential-focus-navigation:tree-order
x-internal="tree-order"} of the
[`Document`{#sequential-focus-navigation:document-4}](dom.html#document).

If a [focusable
area](#focusable-area){#sequential-focus-navigation:focusable-area-2} is
omitted from the [sequential focus navigation
order](#sequential-focus-navigation-order){#sequential-focus-navigation:sequential-focus-navigation-order}
of its
[`Document`{#sequential-focus-navigation:document-5}](dom.html#document),
then it is unreachable via [sequential focus
navigation](#sequential-focus-navigation){#sequential-focus-navigation:sequential-focus-navigation}.

There can also be a [sequential focus navigation starting
point]{#sequential-focus-navigation-starting-point .dfn}. It is
initially unset. The user agent may set it when the user indicates that
it should be moved.

For example, the user agent could set it to the position of the user\'s
click if the user clicks on the document contents.

User agents are required to set the [sequential focus navigation
starting
point](#sequential-focus-navigation-starting-point){#sequential-focus-navigation:sequential-focus-navigation-starting-point}
to the [target
element](browsing-the-web.html#target-element){#sequential-focus-navigation:target-element}
when [navigating to a
fragment](browsing-the-web.html#navigate-fragid){#sequential-focus-navigation:navigate-fragid}.

A [sequential focus direction]{#sequential-focus-direction .dfn} is one
of two possible values: \"[`forward`]{#sequential-focus-forward .dfn}\",
or \"[`backward`]{#sequential-focus-backward .dfn}\". They are used in
the below algorithms to describe the direction in which sequential focus
travels at the user\'s request.

A [selection mechanism]{#selection-mechanism .dfn} is one of two
possible values: \"[`DOM`]{#selection-mechanism-dom .dfn}\", or
\"[`sequential`]{#selection-mechanism-sequential .dfn}\". They are used
to describe how the [sequential navigation search
algorithm](#sequential-navigation-search-algorithm){#sequential-focus-navigation:sequential-navigation-search-algorithm}
finds the [focusable
area](#focusable-area){#sequential-focus-navigation:focusable-area-3} it
returns.

When the user requests that focus move from the [currently focused area
of a top-level
traversable](#currently-focused-area-of-a-top-level-traversable){#sequential-focus-navigation:currently-focused-area-of-a-top-level-traversable}
to the next or previous [focusable
area](#focusable-area){#sequential-focus-navigation:focusable-area-4}
(e.g., as the default action of pressing the [[tab]{.kbd}]{.kbd} key),
or when the user requests that focus sequentially move to a [top-level
traversable](document-sequences.html#top-level-traversable){#sequential-focus-navigation:top-level-traversable}
in the first place (e.g., from the browser\'s location bar), the user
agent must use the following algorithm:

1.  Let `starting point`{.variable} be the [currently focused area of a
    top-level
    traversable](#currently-focused-area-of-a-top-level-traversable){#sequential-focus-navigation:currently-focused-area-of-a-top-level-traversable-2},
    if the user requested to move focus sequentially from there, or else
    the [top-level
    traversable](document-sequences.html#top-level-traversable){#sequential-focus-navigation:top-level-traversable-2}
    itself, if the user instead requested to move focus from outside the
    [top-level
    traversable](document-sequences.html#top-level-traversable){#sequential-focus-navigation:top-level-traversable-3}.

2.  If there is a [sequential focus navigation starting
    point](#sequential-focus-navigation-starting-point){#sequential-focus-navigation:sequential-focus-navigation-starting-point-2}
    defined and it is inside `starting point`{.variable}, then let
    `starting point`{.variable} be the [sequential focus navigation
    starting
    point](#sequential-focus-navigation-starting-point){#sequential-focus-navigation:sequential-focus-navigation-starting-point-3}
    instead.

3.  Let `direction`{.variable} be
    \"[`forward`{#sequential-focus-navigation:sequential-focus-forward}](#sequential-focus-forward)\"
    if the user requested the next control, and
    \"[`backward`{#sequential-focus-navigation:sequential-focus-backward}](#sequential-focus-backward)\"
    if the user requested the previous control.

    Typically, pressing [[tab]{.kbd}]{.kbd} requests the next control,
    and pressing [[shift]{.kbd} + [tab]{.kbd}]{.kbd} requests the
    previous control.

4.  *Loop*: Let `selection mechanism`{.variable} be
    \"[`sequential`{#sequential-focus-navigation:selection-mechanism-sequential}](#selection-mechanism-sequential)\"
    if `starting point`{.variable} is a
    [navigable](document-sequences.html#navigable){#sequential-focus-navigation:navigable}
    or if `starting point`{.variable} is in its
    [`Document`{#sequential-focus-navigation:document-6}](dom.html#document)\'s
    [sequential focus navigation
    order](#sequential-focus-navigation-order){#sequential-focus-navigation:sequential-focus-navigation-order-2}.

    Otherwise, `starting point`{.variable} is not in its
    [`Document`{#sequential-focus-navigation:document-7}](dom.html#document)\'s
    [sequential focus navigation
    order](#sequential-focus-navigation-order){#sequential-focus-navigation:sequential-focus-navigation-order-3};
    let `selection mechanism`{.variable} be
    \"[`DOM`{#sequential-focus-navigation:selection-mechanism-dom}](#selection-mechanism-dom)\".

5.  Let `candidate`{.variable} be the result of running the [sequential
    navigation search
    algorithm](#sequential-navigation-search-algorithm){#sequential-focus-navigation:sequential-navigation-search-algorithm-2}
    with `starting point`{.variable}, `direction`{.variable}, and
    `selection mechanism`{.variable}.

6.  If `candidate`{.variable} is not null, then run the [focusing
    steps](#focusing-steps){#sequential-focus-navigation:focusing-steps}
    for `candidate`{.variable} and return.

7.  Otherwise, unset the [sequential focus navigation starting
    point](#sequential-focus-navigation-starting-point){#sequential-focus-navigation:sequential-focus-navigation-starting-point-4}.

8.  If `starting point`{.variable} is a [top-level
    traversable](document-sequences.html#top-level-traversable){#sequential-focus-navigation:top-level-traversable-4},
    or a [focusable
    area](#focusable-area){#sequential-focus-navigation:focusable-area-5}
    in the [top-level
    traversable](document-sequences.html#top-level-traversable){#sequential-focus-navigation:top-level-traversable-5},
    the user agent should transfer focus to its own controls
    appropriately (if any), honouring `direction`{.variable}, and then
    return.

    For example, if `direction`{.variable} is *backward*, then the last
    [sequentially
    focusable](#sequentially-focusable){#sequential-focus-navigation:sequentially-focusable}
    control before the browser\'s rendering area would be the control to
    focus.

    If the user agent has no [sequentially
    focusable](#sequentially-focusable){#sequential-focus-navigation:sequentially-focusable-2}
    controls --- a kiosk-mode browser, for instance --- then the user
    agent may instead restart these steps with the
    `starting point`{.variable} being the [top-level
    traversable](document-sequences.html#top-level-traversable){#sequential-focus-navigation:top-level-traversable-6}
    itself.

9.  Otherwise, `starting point`{.variable} is a [focusable
    area](#focusable-area){#sequential-focus-navigation:focusable-area-6}
    in a [child
    navigable](document-sequences.html#child-navigable){#sequential-focus-navigation:child-navigable}.
    Set `starting point`{.variable} to that [child
    navigable](document-sequences.html#child-navigable){#sequential-focus-navigation:child-navigable-2}\'s
    [parent](document-sequences.html#nav-parent){#sequential-focus-navigation:nav-parent}
    and return to the step labeled *loop*.

The [sequential navigation search
algorithm]{#sequential-navigation-search-algorithm .dfn}, given a
[focusable
area](#focusable-area){#sequential-focus-navigation:focusable-area-7}
`starting point`{.variable}, [sequential focus
direction](#sequential-focus-direction){#sequential-focus-navigation:sequential-focus-direction}
`direction`{.variable}, and [selection
mechanism](#selection-mechanism){#sequential-focus-navigation:selection-mechanism}
`selection mechanism`{.variable}, consists of the following steps. They
return a [focusable
area](#focusable-area){#sequential-focus-navigation:focusable-area-8}-or-null.

1.  Pick the appropriate cell from the following table, and follow the
    instructions in that cell.

    The appropriate cell is the one that is from the column whose header
    describes `direction`{.variable} and from the first row whose header
    describes `starting point`{.variable} and
    `selection mechanism`{.variable}.

    `direction`{.variable} is
    \"[`forward`{#sequential-focus-navigation:sequential-focus-forward-2}](#sequential-focus-forward)\"

    `direction`{.variable} is
    \"[`backward`{#sequential-focus-navigation:sequential-focus-backward-2}](#sequential-focus-backward)\"

    `starting point`{.variable} is a
    [navigable](document-sequences.html#navigable){#sequential-focus-navigation:navigable-2}

    Let `candidate`{.variable} be the first [suitable sequentially
    focusable
    area](#suitable-sequentially-focusable-area){#sequential-focus-navigation:suitable-sequentially-focusable-area}
    in `starting point`{.variable}\'s [active
    document](document-sequences.html#nav-document){#sequential-focus-navigation:nav-document},
    if any; or else null

    Let `candidate`{.variable} be the last [suitable sequentially
    focusable
    area](#suitable-sequentially-focusable-area){#sequential-focus-navigation:suitable-sequentially-focusable-area-2}
    in `starting point`{.variable}\'s [active
    document](document-sequences.html#nav-document){#sequential-focus-navigation:nav-document-2},
    if any; or else null

    `selection mechanism`{.variable} is
    \"[`DOM`{#sequential-focus-navigation:selection-mechanism-dom-2}](#selection-mechanism-dom)\"

    Let `candidate`{.variable} be the [suitable sequentially focusable
    area](#suitable-sequentially-focusable-area){#sequential-focus-navigation:suitable-sequentially-focusable-area-3},
    that appears nearest after `starting point`{.variable} in
    `starting point`{.variable}\'s
    [`Document`{#sequential-focus-navigation:document-8}](dom.html#document),
    in [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#sequential-focus-navigation:shadow-including-tree-order
    x-internal="shadow-including-tree-order"}, if any; or else null

    In this case, `starting point`{.variable} does not necessarily
    belong to its
    [`Document`{#sequential-focus-navigation:document-9}](dom.html#document)\'s
    [sequential focus navigation
    order](#sequential-focus-navigation-order){#sequential-focus-navigation:sequential-focus-navigation-order-4},
    so we\'ll select the
    [suitable](#suitable-sequentially-focusable-area){#sequential-focus-navigation:suitable-sequentially-focusable-area-4}
    [item](https://infra.spec.whatwg.org/#list-item){#sequential-focus-navigation:list-item
    x-internal="list-item"} from that list that comes *after*
    `starting point`{.variable} in [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#sequential-focus-navigation:shadow-including-tree-order-2
    x-internal="shadow-including-tree-order"}.

    Let `candidate`{.variable} be the [suitable sequentially focusable
    area](#suitable-sequentially-focusable-area){#sequential-focus-navigation:suitable-sequentially-focusable-area-5},
    that appears nearest before `starting point`{.variable} in
    `starting point`{.variable}\'s
    [`Document`{#sequential-focus-navigation:document-10}](dom.html#document),
    in [shadow-including tree
    order](https://dom.spec.whatwg.org/#concept-shadow-including-tree-order){#sequential-focus-navigation:shadow-including-tree-order-3
    x-internal="shadow-including-tree-order"}, if any; or else null

    `selection mechanism`{.variable} is
    \"[`sequential`{#sequential-focus-navigation:selection-mechanism-sequential-2}](#selection-mechanism-sequential)\"

    Let `candidate`{.variable} be the first [suitable sequentially
    focusable
    area](#suitable-sequentially-focusable-area){#sequential-focus-navigation:suitable-sequentially-focusable-area-6}
    after `starting point`{.variable}, in `starting point`{.variable}\'s
    [`Document`{#sequential-focus-navigation:document-11}](dom.html#document)\'s
    [sequential focus navigation
    order](#sequential-focus-navigation-order){#sequential-focus-navigation:sequential-focus-navigation-order-5},
    if any; or else null

    Let `candidate`{.variable} be the last [suitable sequentially
    focusable
    area](#suitable-sequentially-focusable-area){#sequential-focus-navigation:suitable-sequentially-focusable-area-7}
    before `starting point`{.variable}, in
    `starting point`{.variable}\'s
    [`Document`{#sequential-focus-navigation:document-12}](dom.html#document)\'s
    [sequential focus navigation
    order](#sequential-focus-navigation-order){#sequential-focus-navigation:sequential-focus-navigation-order-6},
    if any; or else null

    A [suitable sequentially focusable
    area]{#suitable-sequentially-focusable-area .dfn} is a [focusable
    area](#focusable-area){#sequential-focus-navigation:focusable-area-9}
    whose [DOM
    anchor](#dom-anchor){#sequential-focus-navigation:dom-anchor} is not
    [inert](#inert){#sequential-focus-navigation:inert} and is
    [sequentially
    focusable](#sequentially-focusable){#sequential-focus-navigation:sequentially-focusable-3}.

2.  If `candidate`{.variable} is a [navigable
    container](document-sequences.html#navigable-container){#sequential-focus-navigation:navigable-container}
    with a non-null [content
    navigable](document-sequences.html#content-navigable){#sequential-focus-navigation:content-navigable},
    then:

    1.  Let `recursive candidate`{.variable} be the result of running
        the [sequential navigation search
        algorithm](#sequential-navigation-search-algorithm){#sequential-focus-navigation:sequential-navigation-search-algorithm-3}
        with `candidate`{.variable}\'s [content
        navigable](document-sequences.html#content-navigable){#sequential-focus-navigation:content-navigable-2},
        `direction`{.variable}, and
        \"[`sequential`{#sequential-focus-navigation:selection-mechanism-sequential-3}](#selection-mechanism-sequential)\".

    2.  If `recursive candidate`{.variable} is null, then return the
        result of running the [sequential navigation search
        algorithm](#sequential-navigation-search-algorithm){#sequential-focus-navigation:sequential-navigation-search-algorithm-4}
        with `candidate`{.variable}, `direction`{.variable}, and
        `selection mechanism`{.variable}.

    3.  Otherwise, set `candidate`{.variable} to
        `recursive candidate`{.variable}.

3.  Return `candidate`{.variable}.

#### [6.6.6]{.secno} Focus management APIs[](#focus-management-apis){.self-link}

``` idl
dictionary FocusOptions {
  boolean preventScroll = false;
  boolean focusVisible;
};
```

`documentOrShadowRoot`{.variable}`.`[`activeElement`](#dom-documentorshadowroot-activeelement){#dom-documentorshadowroot-activeelement-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/activeElement](https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement "The activeElement read-only property of the Document interface returns the Element within the DOM that currently has focus.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[ShadowRoot/activeElement](https://developer.mozilla.org/en-US/docs/Web/API/ShadowRoot/activeElement "The activeElement read-only property of the ShadowRoot interface returns the element within the shadow tree that has focus.")

Support in all current engines.

::: support
[Firefox63+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome53+]{.chrome
.yes}

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
:::::::

Returns the deepest element in `documentOrShadowRoot`{.variable} through
which or to which key events are being routed. This is, roughly
speaking, the focused element in the document.

For the purposes of this API, when a [child
navigable](document-sequences.html#child-navigable){#focus-management-apis:child-navigable}
is focused, its
[container](document-sequences.html#nav-container){#focus-management-apis:nav-container}
is [focused](#bc-focus-ergo-bcc-focus) within its
[parent](document-sequences.html#nav-parent){#focus-management-apis:nav-parent}\'s
[active
document](document-sequences.html#nav-document){#focus-management-apis:nav-document}.
For example, if the user moves the focus to a text control in an
[`iframe`{#focus-management-apis:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
the
[`iframe`{#focus-management-apis:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
is the element returned by the
[`activeElement`{#focus-management-apis:dom-documentorshadowroot-activeelement}](#dom-documentorshadowroot-activeelement)
API in the
[`iframe`{#focus-management-apis:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element)\'s
[node
document](https://dom.spec.whatwg.org/#concept-node-document){#focus-management-apis:node-document
x-internal="node-document"}.

Similarly, when the focused element is in a different [node
tree](https://dom.spec.whatwg.org/#concept-node-tree){#focus-management-apis:node-tree
x-internal="node-tree"} than `documentOrShadowRoot`{.variable}, the
element returned will be the
[host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#focus-management-apis:concept-documentfragment-host
x-internal="concept-documentfragment-host"} that\'s located in the same
[node
tree](https://dom.spec.whatwg.org/#concept-node-tree){#focus-management-apis:node-tree-2
x-internal="node-tree"} as `documentOrShadowRoot`{.variable} if
`documentOrShadowRoot`{.variable} is a [shadow-including inclusive
ancestor](https://dom.spec.whatwg.org/#concept-shadow-including-inclusive-ancestor){#focus-management-apis:shadow-including-inclusive-ancestor
x-internal="shadow-including-inclusive-ancestor"} of the focused
element, and null if not.

`document`{.variable}`.`[`hasFocus`](#dom-document-hasfocus){#dom-document-hasfocus-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/hasFocus](https://developer.mozilla.org/en-US/docs/Web/API/Document/hasFocus "The hasFocus() method of the Document interface returns a boolean value indicating whether the document or any element inside the document has focus. This method can be used to determine whether the active element in a document has focus.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns true if key events are being routed through or to
`document`{.variable}; otherwise, returns false. Roughly speaking, this
corresponds to `document`{.variable}, or a document nested inside
`document`{.variable}, being focused.

`window`{.variable}`.`[`focus`](#dom-window-focus){#dom-window-focus-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/focus](https://developer.mozilla.org/en-US/docs/Web/API/Window/focus "Makes a request to bring the window to the front. It may fail due to user settings and the window isn't guaranteed to be frontmost before this method returns.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android18+]{.chrome_android .yes}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Moves the focus to `window`{.variable}\'s
[navigable](nav-history-apis.html#window-navigable){#focus-management-apis:window-navigable},
if any.

`element`{.variable}`.`[`focus`](#dom-focus){#dom-focus-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/focus](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus "The HTMLElement.focus() method sets focus on the specified element, if it can be focused. The focused element is the element that will receive keyboard and similar events by default.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

`element`{.variable}`.`[`focus`](#dom-focus){#focus-management-apis:dom-focus}`({ `[`preventScroll`](#dom-focusoptions-preventscroll){#focus-management-apis:dom-focusoptions-preventscroll-2}`, `[`focusVisible`](#dom-focusoptions-focusvisible){#focus-management-apis:dom-focusoptions-focusvisible-2}` })`

Moves the focus to `element`{.variable}.

If `element`{.variable} is a [navigable
container](document-sequences.html#navigable-container){#focus-management-apis:navigable-container},
moves the focus to its [content
navigable](document-sequences.html#content-navigable){#focus-management-apis:content-navigable}
instead.

By default, this method also scrolls `element`{.variable} into view.
Providing the
[`preventScroll`{#focus-management-apis:dom-focusoptions-preventscroll-3}](#dom-focusoptions-preventscroll)
option and setting it to true prevents this behavior.

By default, user agents use
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#focus-management-apis:implementation-defined
x-internal="implementation-defined"} heuristics to determine whether to
[indicate
focus](https://drafts.csswg.org/selectors/#indicate-focus){#focus-management-apis:indicate-focus
x-internal="indicate-focus"} via a focus ring. Providing the
[`focusVisible`{#focus-management-apis:dom-focusoptions-focusvisible-3}](#dom-focusoptions-focusvisible)
option and setting it to true will ensure the focus ring is always
visible.

`element`{.variable}`.`[`blur`](#dom-blur){#dom-blur-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/blur](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur "The HTMLElement.blur() method removes keyboard focus from the current element.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Moves the focus to the
[viewport](https://drafts.csswg.org/css2/#viewport){#focus-management-apis:viewport
x-internal="viewport"}. Use of this method is discouraged; if you want
to focus the
[viewport](https://drafts.csswg.org/css2/#viewport){#focus-management-apis:viewport-2
x-internal="viewport"}, call the
[`focus()`{#focus-management-apis:dom-focus-2}](#dom-focus) method on
the [`Document`{#focus-management-apis:document}](dom.html#document)\'s
[document
element](https://dom.spec.whatwg.org/#document-element){#focus-management-apis:document-element
x-internal="document-element"}.

Do not use this method to hide the focus ring if you find the focus ring
unsightly. Instead, use the
[`:focus-visible`{#focus-management-apis::focus-visible}](https://drafts.csswg.org/selectors/#the-focus-visible-pseudo){x-internal=":focus-visible"}
pseudo-class to override the
[\'outline\'](https://drafts.csswg.org/css-ui/#outline){#focus-management-apis:'outline'
x-internal="'outline'"} property, and provide a different way to show
what element is focused. Be aware that if an alternative focusing style
isn\'t made available, the page will be significantly less usable for
people who primarily navigate pages using a keyboard, or those with
reduced vision who use focus outlines to help them navigate the page.

::: example
For example, to hide the outline from
[`textarea`{#focus-management-apis:the-textarea-element}](form-elements.html#the-textarea-element)
elements and instead use a yellow background to indicate focus, you
could use:

``` css
textarea:focus-visible { outline: none; background: yellow; color: black; }
```
:::

The
[`DocumentOrShadowRoot`{#focus-management-apis:documentorshadowroot}](dom.html#documentorshadowroot)
[`activeElement`]{#dom-documentorshadowroot-activeelement .dfn
dfn-for="DocumentOrShadowRoot" dfn-type="attribute"} getter steps are:

1.  Let `candidate`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#focus-management-apis:node-document-2
    x-internal="node-document"}\'s [focused
    area](#focused-area-of-the-document){#focus-management-apis:focused-area-of-the-document}\'s
    [DOM anchor](#dom-anchor){#focus-management-apis:dom-anchor}.

2.  Set `candidate`{.variable} to the result of
    [retargeting](https://dom.spec.whatwg.org/#retarget){#focus-management-apis:dom-retarget
    x-internal="dom-retarget"} `candidate`{.variable} against
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-2
    x-internal="this"}.

3.  If `candidate`{.variable}\'s
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#focus-management-apis:root
    x-internal="root"} is not
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-3
    x-internal="this"}, then return null.

4.  If `candidate`{.variable} is not a
    [`Document`{#focus-management-apis:document-2}](dom.html#document)
    object, then return `candidate`{.variable}.

5.  If `candidate`{.variable} has a [body
    element](dom.html#the-body-element-2){#focus-management-apis:the-body-element-2},
    then return that [body
    element](dom.html#the-body-element-2){#focus-management-apis:the-body-element-2-2}.

6.  If `candidate`{.variable}\'s [document
    element](https://dom.spec.whatwg.org/#document-element){#focus-management-apis:document-element-2
    x-internal="document-element"} is non-null, then return that
    [document
    element](https://dom.spec.whatwg.org/#document-element){#focus-management-apis:document-element-3
    x-internal="document-element"}.

7.  Return null.

The [`Document`{#focus-management-apis:document-3}](dom.html#document)
[`hasFocus()`]{#dom-document-hasfocus .dfn dfn-for="Document"
dfn-type="method"} method steps are to return the result of running the
[has focus
steps](#has-focus-steps){#focus-management-apis:has-focus-steps} given
[this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-4
x-internal="this"}.

------------------------------------------------------------------------

The
[`Window`{#focus-management-apis:window}](nav-history-apis.html#window)
[`focus()`]{#dom-window-focus .dfn dfn-for="Window" dfn-type="method"}
method steps are:

1.  Let `current`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-5
    x-internal="this"}\'s
    [navigable](nav-history-apis.html#window-navigable){#focus-management-apis:window-navigable-2}.

2.  If `current`{.variable} is null, then return.

3.  If the [allow focus
    steps](#allow-focus-steps){#focus-management-apis:allow-focus-steps}
    given `current`{.variable}\'s [active
    document](document-sequences.html#nav-document){#focus-management-apis:nav-document-2}
    return false, then return.

4.  Run the [focusing
    steps](#focusing-steps){#focus-management-apis:focusing-steps} with
    `current`{.variable}.

5.  If `current`{.variable} is a [top-level
    traversable](document-sequences.html#top-level-traversable){#focus-management-apis:top-level-traversable},
    user agents are encouraged to trigger some sort of notification to
    indicate to the user that the page is attempting to gain focus.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/blur](https://developer.mozilla.org/en-US/docs/Web/API/Window/blur "Shifts focus away from the window.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The
[`Window`{#focus-management-apis:window-2}](nav-history-apis.html#window)
[`blur()`]{#dom-window-blur .dfn dfn-for="Window" dfn-type="method"}
method steps are to do nothing.

Historically, the
[`focus()`{#focus-management-apis:dom-window-focus}](#dom-window-focus)
and [`blur()`{#focus-management-apis:dom-window-blur}](#dom-window-blur)
methods actually affected the system-level focus of the system widget
(e.g., tab or window) that contained the
[navigable](document-sequences.html#navigable){#focus-management-apis:navigable},
but hostile sites widely abuse this behavior to the user\'s detriment.

------------------------------------------------------------------------

The
[`HTMLOrSVGElement`{#focus-management-apis:htmlorsvgelement}](dom.html#htmlorsvgelement)
[`focus(``options`{.variable}`)`]{#dom-focus .dfn
dfn-for="HTMLOrSVGElement" dfn-type="method"} method steps are:

1.  If the [allow focus
    steps](#allow-focus-steps){#focus-management-apis:allow-focus-steps-2}
    given
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-6
    x-internal="this"}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#focus-management-apis:node-document-3
    x-internal="node-document"} return false, then return.

2.  Run the [focusing
    steps](#focusing-steps){#focus-management-apis:focusing-steps-2} for
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-7
    x-internal="this"}.

3.  If
    `options`{.variable}\[\"[`focusVisible`]{#dom-focusoptions-focusvisible
    .dfn dfn-for="FocusOptions" dfn-type="dict-member"}\"\] is true, or
    does not
    [exist](https://infra.spec.whatwg.org/#map-exists){#focus-management-apis:map-exists
    x-internal="map-exists"} but in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#focus-management-apis:implementation-defined-2
    x-internal="implementation-defined"} way the user agent determines
    it would be best to do so, then [indicate
    focus](https://drafts.csswg.org/selectors/#indicate-focus){#focus-management-apis:indicate-focus-2
    x-internal="indicate-focus"}.

4.  If
    `options`{.variable}\[\"[`preventScroll`]{#dom-focusoptions-preventscroll
    .dfn dfn-for="FocusOptions" dfn-type="dict-member"}\"\] is false,
    then [scroll a target into
    view](https://drafts.csswg.org/cssom-view/#scroll-a-target-into-view){#focus-management-apis:scroll-a-target-into-view
    x-internal="scroll-a-target-into-view"} given
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-8
    x-internal="this"}, \"`auto`\", \"`center`\", and \"`center`\".

The
[`HTMLOrSVGElement`{#focus-management-apis:htmlorsvgelement-2}](dom.html#htmlorsvgelement)
[`blur()`]{#dom-blur .dfn dfn-for="HTMLOrSVGElement" dfn-type="method"}
method steps are:

1.  The user agent should run the [unfocusing
    steps](#unfocusing-steps){#focus-management-apis:unfocusing-steps}
    given
    [this](https://webidl.spec.whatwg.org/#this){#focus-management-apis:this-9
    x-internal="this"}.

    User agents may instead selectively or uniformly do nothing, for
    usability reasons.

For example, if the
[`blur()`{#focus-management-apis:dom-blur}](#dom-blur) method is
unwisely being used to remove the focus ring for aesthetics reasons, the
page would become unusable by keyboard users. Ignoring calls to this
method would thus allow keyboard users to interact with the page.

------------------------------------------------------------------------

The [allow focus steps]{#allow-focus-steps .dfn export=""}, given a
[`Document`{#focus-management-apis:document-4}](dom.html#document)
object `target`{.variable}, are:

1.  If `target`{.variable} is [allowed to
    use](iframe-embed-object.html#allowed-to-use){#focus-management-apis:allowed-to-use}
    the
    \"[`focus-without-user-activation`{#focus-management-apis:focus-without-user-activation-feature}](infrastructure.html#focus-without-user-activation-feature)\"
    feature, then return true.

2.  If `target`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#focus-management-apis:concept-relevant-global}
    has [transient
    activation](#transient-activation){#focus-management-apis:transient-activation},
    then return true.

3.  Return false.

#### [6.6.7]{.secno} []{#autofocusing-a-form-control:-the-autofocus-attribute}The [`autofocus`{#the-autofocus-attribute:attr-fe-autofocus}](#attr-fe-autofocus) attribute[](#the-autofocus-attribute){.self-link}

The [`autofocus`]{#attr-fe-autofocus .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute allows the author to indicate
that an element is to be focused as soon as the page is loaded, allowing
the user to just start typing without having to manually focus the main
element.

When the
[`autofocus`{#the-autofocus-attribute:attr-fe-autofocus-2}](#attr-fe-autofocus)
attribute is specified on an element inside
[`dialog`{#the-autofocus-attribute:the-dialog-element}](interactive-elements.html#the-dialog-element)
elements or [HTML
elements](infrastructure.html#html-elements){#the-autofocus-attribute:html-elements}
whose
[`popover`{#the-autofocus-attribute:attr-popover}](popover.html#attr-popover)
attribute is set, then it will be focused when the dialog or popover
becomes shown.

The
[`autofocus`{#the-autofocus-attribute:attr-fe-autofocus-3}](#attr-fe-autofocus)
attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-autofocus-attribute:boolean-attribute}.

To find the [nearest ancestor autofocus scoping root
element]{#nearest-ancestor-autofocus-scoping-root-element .dfn} given an
[`Element`{#the-autofocus-attribute:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
`element`{.variable}:

1.  If `element`{.variable} is a
    [`dialog`{#the-autofocus-attribute:the-dialog-element-2}](interactive-elements.html#the-dialog-element)
    element, then return `element`{.variable}.

2.  If `element`{.variable}\'s
    [`popover`{#the-autofocus-attribute:attr-popover-2}](popover.html#attr-popover)
    attribute is not in the [No Popover
    State](popover.html#attr-popover-none-state){#the-autofocus-attribute:attr-popover-none-state},
    then return `element`{.variable}.

3.  Let `ancestor`{.variable} be `element`{.variable}.

4.  While `ancestor`{.variable} has a [parent
    element](https://dom.spec.whatwg.org/#parent-element){#the-autofocus-attribute:parent-element
    x-internal="parent-element"}:

    1.  Set `ancestor`{.variable} to `ancestor`{.variable}\'s [parent
        element](https://dom.spec.whatwg.org/#parent-element){#the-autofocus-attribute:parent-element-2
        x-internal="parent-element"}.

    2.  If `ancestor`{.variable} is a
        [`dialog`{#the-autofocus-attribute:the-dialog-element-3}](interactive-elements.html#the-dialog-element)
        element, then return `ancestor`{.variable}.

    3.  If `ancestor`{.variable}\'s
        [`popover`{#the-autofocus-attribute:attr-popover-3}](popover.html#attr-popover)
        attribute is not in the [No Popover
        State](popover.html#attr-popover-none-state){#the-autofocus-attribute:attr-popover-none-state-2},
        then return `ancestor`{.variable}.

5.  Return `ancestor`{.variable}.

There must not be two elements with the same [nearest ancestor autofocus
scoping root
element](#nearest-ancestor-autofocus-scoping-root-element){#the-autofocus-attribute:nearest-ancestor-autofocus-scoping-root-element}
that both have the
[`autofocus`{#the-autofocus-attribute:attr-fe-autofocus-4}](#attr-fe-autofocus)
attribute specified.

Each [`Document`{#the-autofocus-attribute:document}](dom.html#document)
has an [autofocus candidates]{#autofocus-candidates .dfn}
[list](https://infra.spec.whatwg.org/#list){#the-autofocus-attribute:list
x-internal="list"}, initially empty.

Each
[`Document`{#the-autofocus-attribute:document-2}](dom.html#document) has
an [autofocus processed flag]{#autofocus-processed-flag .dfn} boolean,
initially false.

When an element with the
[`autofocus`{#the-autofocus-attribute:attr-fe-autofocus-5}](#attr-fe-autofocus)
attribute specified is [inserted into a
document](infrastructure.html#insert-an-element-into-a-document){#the-autofocus-attribute:insert-an-element-into-a-document},
run the following steps:

1.  If the user has indicated (for example, by starting to type in a
    form control) that they do not wish focus to be changed, then
    optionally return.

2.  Let `target`{.variable} be the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-autofocus-attribute:node-document
    x-internal="node-document"}.

3.  If `target`{.variable} is not [fully
    active](document-sequences.html#fully-active){#the-autofocus-attribute:fully-active},
    then return.

4.  If `target`{.variable}\'s [active sandboxing flag
    set](browsers.html#active-sandboxing-flag-set){#the-autofocus-attribute:active-sandboxing-flag-set}
    has the [sandboxed automatic features browsing context
    flag](browsers.html#sandboxed-automatic-features-browsing-context-flag){#the-autofocus-attribute:sandboxed-automatic-features-browsing-context-flag},
    then return.

5.  If the [allow focus
    steps](#allow-focus-steps){#the-autofocus-attribute:allow-focus-steps}
    given `target`{.variable} return false, then return.

6.  Let `topDocument`{.variable} be `target`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-autofocus-attribute:node-navigable}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#the-autofocus-attribute:nav-top}\'s
    [active
    document](document-sequences.html#nav-document){#the-autofocus-attribute:nav-document}.

7.  If `topDocument`{.variable}\'s [autofocus processed
    flag](#autofocus-processed-flag){#the-autofocus-attribute:autofocus-processed-flag}
    is false, then
    [remove](https://infra.spec.whatwg.org/#list-remove){#the-autofocus-attribute:list-remove
    x-internal="list-remove"} the element from
    `topDocument`{.variable}\'s [autofocus
    candidates](#autofocus-candidates){#the-autofocus-attribute:autofocus-candidates},
    and
    [append](https://infra.spec.whatwg.org/#list-append){#the-autofocus-attribute:list-append
    x-internal="list-append"} the element to `topDocument`{.variable}\'s
    [autofocus
    candidates](#autofocus-candidates){#the-autofocus-attribute:autofocus-candidates-2}.

We do not check if an element is a [focusable
area](#focusable-area){#the-autofocus-attribute:focusable-area} before
storing it in the [autofocus
candidates](#autofocus-candidates){#the-autofocus-attribute:autofocus-candidates-3}
list, because even if it is not a focusable area when it is inserted, it
could become one by the time [flush autofocus
candidates](#flush-autofocus-candidates){#the-autofocus-attribute:flush-autofocus-candidates}
sees it.

To [flush autofocus candidates]{#flush-autofocus-candidates .dfn} for a
document `topDocument`{.variable}, run these steps:

1.  If `topDocument`{.variable}\'s [autofocus processed
    flag](#autofocus-processed-flag){#the-autofocus-attribute:autofocus-processed-flag-2}
    is true, then return.

2.  Let `candidates`{.variable} be `topDocument`{.variable}\'s
    [autofocus
    candidates](#autofocus-candidates){#the-autofocus-attribute:autofocus-candidates-4}.

3.  If `candidates`{.variable} [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#the-autofocus-attribute:list-is-empty
    x-internal="list-is-empty"}, then return.

4.  If `topDocument`{.variable}\'s [focused
    area](#focused-area-of-the-document){#the-autofocus-attribute:focused-area-of-the-document}
    is not `topDocument`{.variable} itself, or `topDocument`{.variable}
    has non-null [target
    element](browsing-the-web.html#target-element){#the-autofocus-attribute:target-element},
    then:

    1.  [Empty](https://infra.spec.whatwg.org/#list-empty){#the-autofocus-attribute:list-empty
        x-internal="list-empty"} `candidates`{.variable}.

    2.  Set `topDocument`{.variable}\'s [autofocus processed
        flag](#autofocus-processed-flag){#the-autofocus-attribute:autofocus-processed-flag-3}
        to true.

    3.  Return.

5.  While `candidates`{.variable} is not
    [empty](https://infra.spec.whatwg.org/#list-is-empty){#the-autofocus-attribute:list-is-empty-2
    x-internal="list-is-empty"}:

    1.  Let `element`{.variable} be `candidates`{.variable}\[0\].

    2.  Let `doc`{.variable} be `element`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#the-autofocus-attribute:node-document-2
        x-internal="node-document"}.

    3.  If `doc`{.variable} is not [fully
        active](document-sequences.html#fully-active){#the-autofocus-attribute:fully-active-2},
        then
        [remove](https://infra.spec.whatwg.org/#list-remove){#the-autofocus-attribute:list-remove-2
        x-internal="list-remove"} `element`{.variable} from
        `candidates`{.variable}, and
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#the-autofocus-attribute:continue
        x-internal="continue"}.

    4.  If `doc`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#the-autofocus-attribute:node-navigable-2}\'s
        [top-level
        traversable](document-sequences.html#nav-top){#the-autofocus-attribute:nav-top-2}
        is not the same as `topDocument`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#the-autofocus-attribute:node-navigable-3},
        then
        [remove](https://infra.spec.whatwg.org/#list-remove){#the-autofocus-attribute:list-remove-3
        x-internal="list-remove"} `element`{.variable} from
        `candidates`{.variable}, and
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#the-autofocus-attribute:continue-2
        x-internal="continue"}.

    5.  If `doc`{.variable}\'s [script-blocking style sheet
        set](semantics.html#script-blocking-style-sheet-set){#the-autofocus-attribute:script-blocking-style-sheet-set}
        is not
        [empty](https://infra.spec.whatwg.org/#list-is-empty){#the-autofocus-attribute:list-is-empty-3
        x-internal="list-is-empty"}, then return.

        In this case, `element`{.variable} is the currently-best
        candidate, but `doc`{.variable} is not ready for autofocusing.
        We\'ll try again next time [flush autofocus
        candidates](#flush-autofocus-candidates){#the-autofocus-attribute:flush-autofocus-candidates-2}
        is called.

    6.  [Remove](https://infra.spec.whatwg.org/#list-remove){#the-autofocus-attribute:list-remove-4
        x-internal="list-remove"} `element`{.variable} from
        `candidates`{.variable}.

    7.  Let `inclusiveAncestorDocuments`{.variable} be a
        [list](https://infra.spec.whatwg.org/#list){#the-autofocus-attribute:list-2
        x-internal="list"} consisting of the [active
        document](document-sequences.html#nav-document){#the-autofocus-attribute:nav-document-2}
        of `doc`{.variable}\'s [inclusive ancestor
        navigables](document-sequences.html#inclusive-ancestor-navigables){#the-autofocus-attribute:inclusive-ancestor-navigables}.

    8.  If any
        [`Document`{#the-autofocus-attribute:document-3}](dom.html#document)
        in `inclusiveAncestorDocuments`{.variable} has non-null [target
        element](browsing-the-web.html#target-element){#the-autofocus-attribute:target-element-2},
        then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#the-autofocus-attribute:continue-3
        x-internal="continue"}.

    9.  Let `target`{.variable} be `element`{.variable}.

    10. If `target`{.variable} is not a [focusable
        area](#focusable-area){#the-autofocus-attribute:focusable-area-2},
        then set `target`{.variable} to the result of [getting the
        focusable
        area](#get-the-focusable-area){#the-autofocus-attribute:get-the-focusable-area}
        for `target`{.variable}.

        [Autofocus
        candidates](#autofocus-candidates){#the-autofocus-attribute:autofocus-candidates-5}
        can
        [contain](https://infra.spec.whatwg.org/#list-contain){#the-autofocus-attribute:list-contains
        x-internal="list-contains"} elements which are not [focusable
        areas](#focusable-area){#the-autofocus-attribute:focusable-area-3}.
        In addition to the special cases handled in the [get the
        focusable
        area](#get-the-focusable-area){#the-autofocus-attribute:get-the-focusable-area-2}
        algorithm, this can happen because a non-[focusable
        area](#focusable-area){#the-autofocus-attribute:focusable-area-4}
        element with an
        [`autofocus`{#the-autofocus-attribute:attr-fe-autofocus-6}](#attr-fe-autofocus)
        attribute was [inserted into a
        document](infrastructure.html#insert-an-element-into-a-document){#the-autofocus-attribute:insert-an-element-into-a-document-2}
        and it never became focusable, or because the element was
        focusable but its status changed while it was stored in
        [autofocus
        candidates](#autofocus-candidates){#the-autofocus-attribute:autofocus-candidates-6}.

    11. If `target`{.variable} is not null, then:

        1.  [Empty](https://infra.spec.whatwg.org/#list-empty){#the-autofocus-attribute:list-empty-2
            x-internal="list-empty"} `candidates`{.variable}.

        2.  Set `topDocument`{.variable}\'s [autofocus processed
            flag](#autofocus-processed-flag){#the-autofocus-attribute:autofocus-processed-flag-4}
            to true.

        3.  Run the [focusing
            steps](#focusing-steps){#the-autofocus-attribute:focusing-steps}
            for `target`{.variable}.

This handles the automatic focusing during document load. The
[`show()`{#the-autofocus-attribute:dom-dialog-show}](interactive-elements.html#dom-dialog-show)
and
[`showModal()`{#the-autofocus-attribute:dom-dialog-showmodal}](interactive-elements.html#dom-dialog-showmodal)
methods of
[`dialog`{#the-autofocus-attribute:the-dialog-element-4}](interactive-elements.html#the-dialog-element)
elements also processes the
[`autofocus`{#the-autofocus-attribute:attr-fe-autofocus-7}](#attr-fe-autofocus)
attribute.

Focusing the element does not imply that the user agent has to focus the
browser window if it has lost focus.

::::: {.mdn-anno .wrapped .before}
**⚠**MDN

:::: feature
[Global_attributes/autofocus](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autofocus "The autofocus global attribute is a Boolean attribute indicating that an element should be focused on page load, or when the <dialog> that it is part of is displayed.")

Support in one engine only.

::: support
[Firefox[🔰 1+]{title="Partial implementation."}]{.firefox
.yes}[Safari[🔰 4+]{title="Partial implementation."}]{.safari
.yes}[Chrome79+]{.chrome .yes}

------------------------------------------------------------------------

[Opera66+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet Explorer[🔰
10+]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android79+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android57+]{.opera_android .yes}
:::
::::
:::::

The [`autofocus`]{#dom-fe-autofocus .dfn dfn-for="HTMLOrSVGElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-autofocus-attribute:reflect}
the content attribute of the same name.

::: example
In the following snippet, the text control would be focused when the
document was loaded.

``` html
<input maxlength="256" name="q" value="" autofocus>
<input type="submit" value="Search">
```
:::

::: example
The
[`autofocus`{#the-autofocus-attribute:attr-fe-autofocus-8}](#attr-fe-autofocus)
attribute applies to all elements, not just to form controls. This
allows examples such as the following:

``` html
<div contenteditable autofocus>Edit <strong>me!</strong><div>
```
:::

### [6.7]{.secno} Assigning keyboard shortcuts[](#assigning-keyboard-shortcuts){.self-link}

#### [6.7.1]{.secno} Introduction[](#introduction-9){.self-link} {#introduction-9}

*This section is non-normative.*

Each element that can be activated or focused can be assigned a single
key combination to activate it, using the
[`accesskey`{#introduction-9:the-accesskey-attribute}](#the-accesskey-attribute)
attribute.

The exact shortcut is determined by the user agent, based on information
about the user\'s keyboard, what keyboard shortcuts already exist on the
platform, and what other shortcuts have been specified on the page,
using the information provided in the
[`accesskey`{#introduction-9:the-accesskey-attribute-2}](#the-accesskey-attribute)
attribute as a guide.

In order to ensure that a relevant keyboard shortcut is available on a
wide variety of input devices, the author can provide a number of
alternatives in the
[`accesskey`{#introduction-9:the-accesskey-attribute-3}](#the-accesskey-attribute)
attribute.

Each alternative consists of a single character, such as a letter or
digit.

User agents can provide users with a list of the keyboard shortcuts, but
authors are encouraged to do so also. The
[`accessKeyLabel`{#introduction-9:dom-accesskeylabel}](#dom-accesskeylabel)
IDL attribute returns a string representing the actual key combination
assigned by the user agent.

::: example
In this example, an author has provided a button that can be invoked
using a shortcut key. To support full keyboards, the author has provided
\"C\" as a possible key. To support devices equipped only with numeric
keypads, the author has provided \"1\" as another possible key.

``` html
<input type=button value=Collect onclick="collect()"
       accesskey="C 1" id=c>
```
:::

::: example
To tell the user what the shortcut key is, the author has this script
here opted to explicitly add the key combination to the button\'s label:

``` js
function addShortcutKeyLabel(button) {
  if (button.accessKeyLabel != '')
    button.value += ' (' + button.accessKeyLabel + ')';
}
addShortcutKeyLabel(document.getElementById('c'));
```

Browsers on different platforms will show different labels, even for the
same key combination, based on the convention prevalent on that
platform. For example, if the key combination is the Control key, the
Shift key, and the letter C, a Windows browser might display
\"`Ctrl+Shift+C`{.sample}\", whereas a Mac browser might display
\"`^⇧C`{.sample}\", while an Emacs browser might just display
\"`C-C`{.sample}\". Similarly, if the key combination is the Alt key and
the Escape key, Windows might use \"`Alt+Esc`{.sample}\", Mac might use
\"`⌥⎋`{.sample}\", and an Emacs browser might use \"`M-ESC`{.sample}\"
or \"`ESC ESC`{.sample}\".

In general, therefore, it is unwise to attempt to parse the value
returned from the
[`accessKeyLabel`{#introduction-9:dom-accesskeylabel-2}](#dom-accesskeylabel)
IDL attribute.
:::

#### [6.7.2]{.secno} The [`accesskey`]{.dfn dfn-for="html-global" dfn-type="element-attr"} attribute[](#the-accesskey-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Global_attributes/accesskey](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/accesskey "The accesskey global attribute provides a hint for generating a keyboard shortcut for the current element. The attribute value must consist of a single printable character (which includes accented and other characters that can be generated by the keyboard).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

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
elements](infrastructure.html#html-elements){#the-accesskey-attribute:html-elements}
may have the
[`accesskey`{#the-accesskey-attribute:the-accesskey-attribute}](#the-accesskey-attribute)
content attribute set. The
[`accesskey`{#the-accesskey-attribute:the-accesskey-attribute-2}](#the-accesskey-attribute)
attribute\'s value is used by the user agent as a guide for creating a
keyboard shortcut that activates or focuses the element.

If specified, the value must be an [ordered set of unique
space-separated
tokens](common-microsyntaxes.html#ordered-set-of-unique-space-separated-tokens){#the-accesskey-attribute:ordered-set-of-unique-space-separated-tokens}
none of which are [identical
to](https://infra.spec.whatwg.org/#string-is){#the-accesskey-attribute:identical-to
x-internal="identical-to"} another token and each of which must be
exactly one code point in length.

::: example
In the following example, a variety of links are given with access keys
so that keyboard users familiar with the site can more quickly navigate
to the relevant pages:

``` html
<nav>
 <p>
  <a title="Consortium Activities" accesskey="A" href="/Consortium/activities">Activities</a> |
  <a title="Technical Reports and Recommendations" accesskey="T" href="/TR/">Technical Reports</a> |
  <a title="Alphabetical Site Index" accesskey="S" href="/Consortium/siteindex">Site Index</a> |
  <a title="About This Site" accesskey="B" href="/Consortium/">About Consortium</a> |
  <a title="Contact Consortium" accesskey="C" href="/Consortium/contact">Contact</a>
 </p>
</nav>
```
:::

::: example
In the following example, the search field is given two possible access
keys, \"s\" and \"0\" (in that order). A user agent on a device with a
full keyboard might pick [[Ctrl]{.kbd} + [Alt]{.kbd} + [S]{.kbd}]{.kbd}
as the shortcut key, while a user agent on a small device with just a
numeric keypad might pick just the plain unadorned key
[[0]{.kbd}]{.kbd}:

``` html
<form action="/search">
 <label>Search: <input type="search" name="q" accesskey="s 0"></label>
 <input type="submit">
</form>
```
:::

::: example
In the following example, a button has possible access keys described. A
script then tries to update the button\'s label to advertise the key
combination the user agent selected.

``` html
<input type=submit accesskey="N @ 1" value="Compose">
...
<script>
 function labelButton(button) {
   if (button.accessKeyLabel)
     button.value += ' (' + button.accessKeyLabel + ')';
 }
 var inputs = document.getElementsByTagName('input');
 for (var i = 0; i < inputs.length; i += 1) {
   if (inputs[i].type == "submit")
     labelButton(inputs[i]);
 }
</script>
```

On one user agent, the button\'s label might become
\"`Compose (⌘N)`{.sample}\". On another, it might become
\"`Compose (Alt+⇧+1)`{.sample}\". If the user agent doesn\'t assign a
key, it will be just \"`Compose`{.sample}\". The exact string depends on
what the [assigned access
key](#assigned-access-key){#the-accesskey-attribute:assigned-access-key}
is, and on how the user agent represents that key combination.
:::

#### [6.7.3]{.secno} []{#processing-model-6}Processing model[](#keyboard-shortcuts-processing-model){.self-link} {#keyboard-shortcuts-processing-model}

An element\'s [assigned access key]{#assigned-access-key .dfn} is a key
combination derived from the element\'s
[`accesskey`{#keyboard-shortcuts-processing-model:the-accesskey-attribute}](#the-accesskey-attribute)
content attribute. Initially, an element must not have an [assigned
access
key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key}.

Whenever an element\'s
[`accesskey`{#keyboard-shortcuts-processing-model:the-accesskey-attribute-2}](#the-accesskey-attribute)
attribute is set, changed, or removed, the user agent must update the
element\'s [assigned access
key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key-2}
by running the following steps:

1.  If the element has no
    [`accesskey`{#keyboard-shortcuts-processing-model:the-accesskey-attribute-3}](#the-accesskey-attribute)
    attribute, then skip to the *fallback* step below.

2.  Otherwise, [split the attribute\'s value on ASCII
    whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#keyboard-shortcuts-processing-model:split-a-string-on-spaces
    x-internal="split-a-string-on-spaces"}, and let `keys`{.variable} be
    the resulting tokens.

3.  For each value in `keys`{.variable} in turn, in the order the tokens
    appeared in the attribute\'s value, run the following substeps:

    1.  If the value is not a string exactly one code point in length,
        then skip the remainder of these steps for this value.

    2.  If the value does not correspond to a key on the system\'s
        keyboard, then skip the remainder of these steps for this value.

    3.  [![(This is a tracking
        vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
        crossorigin=""
        height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#keyboard-shortcuts-processing-model:tracking-vector
        .tracking-vector x-internal="tracking-vector"} If the user agent
        can find a mix of zero or more modifier keys that, combined with
        the key that corresponds to the value given in the attribute,
        can be used as the access key, then the user agent may assign
        that combination of keys as the element\'s [assigned access
        key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key-3}
        and return.

4.  *Fallback*: Optionally, the user agent may assign a key combination
    of its choosing as the element\'s [assigned access
    key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key-4}
    and then return.

5.  If this step is reached, the element has no [assigned access
    key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key-5}.

Once a user agent has selected and assigned an access key for an
element, the user agent should not change the element\'s [assigned
access
key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key-6}
unless the
[`accesskey`{#keyboard-shortcuts-processing-model:the-accesskey-attribute-4}](#the-accesskey-attribute)
content attribute is changed or the element is moved to another
[`Document`{#keyboard-shortcuts-processing-model:document}](dom.html#document).

When the user presses the key combination corresponding to the [assigned
access
key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key-7}
for an element, if the element [defines a
command](interactive-elements.html#concept-command){#keyboard-shortcuts-processing-model:concept-command},
the command\'s [Hidden
State](interactive-elements.html#command-facet-hiddenstate){#keyboard-shortcuts-processing-model:command-facet-hiddenstate}
facet is false (visible), the command\'s [Disabled
State](interactive-elements.html#command-facet-disabledstate){#keyboard-shortcuts-processing-model:command-facet-disabledstate}
facet is also false (enabled), the element is [in a
document](https://dom.spec.whatwg.org/#in-a-document){#keyboard-shortcuts-processing-model:in-a-document
x-internal="in-a-document"} that has a non-null [browsing
context](document-sequences.html#concept-document-bc){#keyboard-shortcuts-processing-model:concept-document-bc},
and neither the element nor any of its ancestors has a
[`hidden`{#keyboard-shortcuts-processing-model:attr-hidden}](#attr-hidden)
attribute specified, then the user agent must trigger the
[Action](interactive-elements.html#command-facet-action){#keyboard-shortcuts-processing-model:command-facet-action}
of the command.

User agents [might
expose](interactive-elements.html#expose-commands-in-ui) elements that
have an
[`accesskey`{#keyboard-shortcuts-processing-model:the-accesskey-attribute-5}](#the-accesskey-attribute)
attribute in other ways as well, e.g. in a menu displayed in response to
a specific key combination.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/accessKey](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/accessKey "The HTMLElement.accessKey property sets the keystroke which a user can press to jump to a given element.")

Support in all current engines.

::: support
[Firefox5+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome17+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`accessKey`]{#dom-accesskey .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#keyboard-shortcuts-processing-model:reflect}
the
[`accesskey`{#keyboard-shortcuts-processing-model:the-accesskey-attribute-6}](#the-accesskey-attribute)
content attribute.

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLElement/accessKeyLabel](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/accessKeyLabel "The HTMLElement.accessKeyLabel read-only property returns a string containing the element's browser-assigned access key (if any); otherwise it returns an empty string.")

::: support
[Firefox8+]{.firefox .yes}[Safari14+]{.safari .yes}[ChromeNo]{.chrome
.no}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeNo]{.edge_blink .no}

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

The [`accessKeyLabel`]{#dom-accesskeylabel .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute must return a string that represents
the element\'s [assigned access
key](#assigned-access-key){#keyboard-shortcuts-processing-model:assigned-access-key-8},
if any. If the element does not have one, then the IDL attribute must
return the empty string.

### [6.8]{.secno} Editing[](#editing-2){.self-link} {#editing-2}

#### [6.8.1]{.secno} Making document regions editable: The [`contenteditable`{#contenteditable:attr-contenteditable}](#attr-contenteditable) content attribute[](#contenteditable){.self-link} {#contenteditable}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HTMLElement/contentEditable](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/contentEditable "The contentEditable property of the HTMLElement interface specifies whether or not the element is editable.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
interface mixin ElementContentEditable {
  [CEReactions] attribute DOMString contentEditable;
  [CEReactions] attribute DOMString enterKeyHint;
  readonly attribute boolean isContentEditable;
  [CEReactions] attribute DOMString inputMode;
};
```

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/contenteditable](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/contenteditable "The contenteditable global attribute is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its widget to allow editing.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`contenteditable`]{#attr-contenteditable .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#contenteditable:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`true`]{#attr-contenteditable-true .dfn
dfn-for="html-global/contenteditable" dfn-type="attr-value"}

[True]{#attr-contenteditable-true-state .dfn}

The element is editable.

(the empty string)

[`false`]{#attr-contenteditable-false .dfn
dfn-for="html-global/contenteditable" dfn-type="attr-value"}

[False]{#attr-contenteditable-false-state .dfn}

The element is not editable.

[`plaintext-only`]{#attr-contenteditable-plaintextonly .dfn
dfn-for="html-global/contenteditable" dfn-type="attr-value"}

[Plaintext-Only]{#attr-contenteditable-plaintextonly-state .dfn}

Only the element\'s raw text content is editable; rich formatting is
disabled.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Inherit]{#attr-contenteditable-inherit-state .dfn} state. The
inherit state indicates that the element is editable (or not) based on
the parent element\'s state.

::: example
For example, consider a page that has a
[`form`{#contenteditable:the-form-element}](forms.html#the-form-element)
and a
[`textarea`{#contenteditable:the-textarea-element}](form-elements.html#the-textarea-element)
to publish a new article, where the user is expected to write the
article using HTML:

``` html
<form method=POST>
 <fieldset>
  <legend>New article</legend>
  <textarea name=article>&lt;p>Hello world.&lt;/p></textarea>
 </fieldset>
 <p><button>Publish</button></p>
</form>
```

When scripting is enabled, the
[`textarea`{#contenteditable:the-textarea-element-2}](form-elements.html#the-textarea-element)
element could be replaced with a rich text control instead, using the
[`contenteditable`{#contenteditable:attr-contenteditable-2}](#attr-contenteditable)
attribute:

``` html
<form method=POST>
 <fieldset>
  <legend>New article</legend>
  <textarea id=textarea name=article>&lt;p>Hello world.&lt;/p></textarea>
  <div id=div style="white-space: pre-wrap" hidden><p>Hello world.</p></div>
  <script>
   let textarea = document.getElementById("textarea");
   let div = document.getElementById("div");
   textarea.hidden = true;
   div.hidden = false;
   div.contentEditable = "true";
   div.oninput = (e) => {
     textarea.value = div.innerHTML;
   };
  </script>
 </fieldset>
 <p><button>Publish</button></p>
</form>
```

Features to enable, e.g., inserting links, can be implemented using the
[`document.execCommand()`{#contenteditable:execCommand}](https://w3c.github.io/editing/docs/execCommand/#execcommand%28%29){x-internal="execCommand"}
API, or using
[`Selection`{#contenteditable:selection}](https://w3c.github.io/selection-api/#selection-interface){x-internal="selection"}
APIs and other DOM APIs.
[\[EXECCOMMAND\]](references.html#refsEXECCOMMAND)
[\[SELECTION\]](references.html#refsSELECTION)
[\[DOM\]](references.html#refsDOM)
:::

::: example
The
[`contenteditable`{#contenteditable:attr-contenteditable-3}](#attr-contenteditable)
attribute can also be used to great effect:

``` html
<!doctype html>
<html lang=en>
<title>Live CSS editing!</title>
<style style=white-space:pre contenteditable>
html { margin:.2em; font-size:2em; color:lime; background:purple }
head, title, style { display:block }
body { display:none }
</style>
```
:::

`element`{.variable}`.`[`contentEditable`](#dom-contenteditable){#dom-contenteditable-dev}` [ = ``value`{.variable}` ]`

Returns \"`true`\", \"`plaintext-only`\", \"`false`\", or \"`inherit`\",
based on the state of the
[`contenteditable`{#contenteditable:attr-contenteditable-4}](#attr-contenteditable)
attribute.

Can be set, to change that state.

Throws a
[\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#contenteditable:syntaxerror
x-internal="syntaxerror"}
[`DOMException`{#contenteditable:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the new value isn\'t one of those strings.

`element`{.variable}`.`[`isContentEditable`](#dom-iscontenteditable){#dom-iscontenteditable-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/isContentEditable](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/isContentEditable "The HTMLElement.isContentEditable read-only property returns a boolean value that is true if the contents of the element are editable; otherwise it returns false.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns true if the element is editable; otherwise, returns false.

The [`contentEditable`]{#dom-contenteditable .dfn
dfn-for="ElementContentEditable" dfn-type="attribute"} IDL attribute, on
getting, must return the string \"`true`\" if the content attribute is
set to the
[True](#attr-contenteditable-true-state){#contenteditable:attr-contenteditable-true-state}
state, \"`plaintext-only`\" if the content attribute is set to the
[Plaintext-Only](#attr-contenteditable-plaintextonly-state){#contenteditable:attr-contenteditable-plaintextonly-state}
state, \"`false`\" if the content attribute is set to the
[False](#attr-contenteditable-false-state){#contenteditable:attr-contenteditable-false-state}
state, and \"`inherit`\" otherwise. On setting, if the new value is an
[ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#contenteditable:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string \"`inherit`\",
then the content attribute must be removed, if the new value is an
[ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#contenteditable:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for the string \"`true`\",
then the content attribute must be set to the string \"`true`\", if the
new value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#contenteditable:ascii-case-insensitive-3
x-internal="ascii-case-insensitive"} match for the string
\"`plaintext-only`\", then the content attribute must be set to the
string \"`plaintext-only`\", if the new value is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#contenteditable:ascii-case-insensitive-4
x-internal="ascii-case-insensitive"} match for the string \"`false`\",
then the content attribute must be set to the string \"`false`\", and
otherwise the attribute setter must throw a
[\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#contenteditable:syntaxerror-2
x-internal="syntaxerror"}
[`DOMException`{#contenteditable:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

The [`isContentEditable`]{#dom-iscontenteditable .dfn
dfn-for="ElementContentEditable" dfn-type="attribute"} IDL attribute, on
getting, must return true if the element is either an [editing
host](#editing-host){#contenteditable:editing-host} or
[editable](https://w3c.github.io/editing/docs/execCommand/#editable){#contenteditable:editable
x-internal="editable"}, and false otherwise.

#### [6.8.2]{.secno} Making entire documents editable: the [`designMode`{#making-entire-documents-editable:-the-designmode-idl-attribute:designMode}](#designMode) getter and setter[](#making-entire-documents-editable:-the-designmode-idl-attribute){.self-link} {#making-entire-documents-editable:-the-designmode-idl-attribute}

`document`{.variable}`.`[`designMode`](#designMode){#dom-document-designmode-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/designMode](https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode "document.designMode controls whether the entire document is editable. Valid values are "on" and "off". According to the specification, this property is meant to default to "off". Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit". Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.2+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns \"`on`\" if the document is editable, and \"`off`\" if it
isn\'t.

Can be set, to change the document\'s current state. This focuses the
document and resets the selection in that document.

[`Document`{#making-entire-documents-editable:-the-designmode-idl-attribute:document}](dom.html#document)
objects have an associated [design mode enabled]{#design-mode-enabled
.dfn}, which is a boolean. It is initially false.

The [`designMode`]{#designMode .dfn dfn-for="Document"
dfn-type="attribute"} getter steps are to return \"`on`\" if
[this](https://webidl.spec.whatwg.org/#this){#making-entire-documents-editable:-the-designmode-idl-attribute:this
x-internal="this"}\'s [design mode
enabled](#design-mode-enabled){#making-entire-documents-editable:-the-designmode-idl-attribute:design-mode-enabled}
is true; otherwise \"`off`\".

The
[`designMode`{#making-entire-documents-editable:-the-designmode-idl-attribute:designMode-2}](#designMode)
setter steps are:

1.  Let `value`{.variable} be the given value, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#making-entire-documents-editable:-the-designmode-idl-attribute:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

2.  If `value`{.variable} is \"`on`\" and
    [this](https://webidl.spec.whatwg.org/#this){#making-entire-documents-editable:-the-designmode-idl-attribute:this-2
    x-internal="this"}\'s [design mode
    enabled](#design-mode-enabled){#making-entire-documents-editable:-the-designmode-idl-attribute:design-mode-enabled-2}
    is false, then:

    1.  Set
        [this](https://webidl.spec.whatwg.org/#this){#making-entire-documents-editable:-the-designmode-idl-attribute:this-3
        x-internal="this"}\'s [design mode
        enabled](#design-mode-enabled){#making-entire-documents-editable:-the-designmode-idl-attribute:design-mode-enabled-3}
        to true.

    2.  Reset
        [this](https://webidl.spec.whatwg.org/#this){#making-entire-documents-editable:-the-designmode-idl-attribute:this-4
        x-internal="this"}\'s [active
        range](https://w3c.github.io/editing/docs/execCommand/#active-range){#making-entire-documents-editable:-the-designmode-idl-attribute:active-range
        x-internal="active-range"}\'s start and end boundary points to
        be at the start of
        [this](https://webidl.spec.whatwg.org/#this){#making-entire-documents-editable:-the-designmode-idl-attribute:this-5
        x-internal="this"}.

    3.  Run the [focusing
        steps](#focusing-steps){#making-entire-documents-editable:-the-designmode-idl-attribute:focusing-steps}
        for
        [this](https://webidl.spec.whatwg.org/#this){#making-entire-documents-editable:-the-designmode-idl-attribute:this-6
        x-internal="this"}\'s [document
        element](https://dom.spec.whatwg.org/#document-element){#making-entire-documents-editable:-the-designmode-idl-attribute:document-element
        x-internal="document-element"}, if non-null.

3.  If `value`{.variable} is \"`off`\", then set
    [this](https://webidl.spec.whatwg.org/#this){#making-entire-documents-editable:-the-designmode-idl-attribute:this-7
    x-internal="this"}\'s [design mode
    enabled](#design-mode-enabled){#making-entire-documents-editable:-the-designmode-idl-attribute:design-mode-enabled-4}
    to false.

#### [6.8.3]{.secno} Best practices for in-page editors[](#best-practices-for-in-page-editors){.self-link}

Authors are encouraged to set the
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#best-practices-for-in-page-editors:'white-space'
x-internal="'white-space'"} property on [editing
hosts](#editing-host){#best-practices-for-in-page-editors:editing-host}
and on markup that was originally created through these editing
mechanisms to the value \'pre-wrap\'. Default HTML whitespace handling
is not well suited to WYSIWYG editing, and line wrapping will not work
correctly in some corner cases if
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#best-practices-for-in-page-editors:'white-space'-2
x-internal="'white-space'"} is left at its default value.

::: example
As an example of problems that occur if the default \'normal\' value is
used instead, consider the case of the user typing
\"[yellow␣␣ball]{.kbd}\", with two spaces (here represented by \"␣\")
between the words. With the editing rules in place for the default value
of
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#best-practices-for-in-page-editors:'white-space'-3
x-internal="'white-space'"} (\'normal\'), the resulting markup will
either consist of \"`yellow&nbsp; ball`{.sample}\" or
\"`yellow &nbsp;ball`{.sample}\"; i.e., there will be a non-breaking
space between the two words in addition to the regular space. This is
necessary because the \'normal\' value for
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#best-practices-for-in-page-editors:'white-space'-4
x-internal="'white-space'"} requires adjacent regular spaces to be
collapsed together.

In the former case, \"`yellow⍽`{.sample}\" might wrap to the next line
(\"⍽\" being used here to represent a non-breaking space) even though
\"`yellow`{.sample}\" alone might fit at the end of the line; in the
latter case, \"`⍽ball`{.sample}\", if wrapped to the start of the line,
would have visible indentation from the non-breaking space.

When
[\'white-space\'](https://drafts.csswg.org/css-text/#white-space-property){#best-practices-for-in-page-editors:'white-space'-5
x-internal="'white-space'"} is set to \'pre-wrap\', however, the editing
rules will instead simply put two regular spaces between the words, and
should the two words be split at the end of a line, the spaces would be
neatly removed from the rendering.
:::

#### [6.8.4]{.secno} Editing APIs[](#editing-apis){.self-link}

An [editing host]{#editing-host .dfn export=""} is either an [HTML
element](infrastructure.html#html-elements){#editing-apis:html-elements}
with its
[`contenteditable`{#editing-apis:attr-contenteditable}](#attr-contenteditable)
attribute in the *true* state or *plaintext-only* state, or a
[child](https://dom.spec.whatwg.org/#concept-tree-child){#editing-apis:concept-tree-child
x-internal="concept-tree-child"} [HTML
element](infrastructure.html#html-elements){#editing-apis:html-elements-2}
of a [`Document`{#editing-apis:document}](dom.html#document) whose
[design mode
enabled](#design-mode-enabled){#editing-apis:design-mode-enabled} is
true.

The definition of the terms [[active
range](https://w3c.github.io/editing/docs/execCommand/#active-range)]{#active-range
.dfn}, [[editing host
of](https://w3c.github.io/editing/docs/execCommand/#editing-host-of)]{#editing-host-of
.dfn}, and
[[editable](https://w3c.github.io/editing/docs/execCommand/#editable)]{#editable
.dfn}, the user interface requirements of elements that are [editing
hosts](#editing-host){#editing-apis:editing-host} or
[editable](https://w3c.github.io/editing/docs/execCommand/#editable){#editing-apis:editable
x-internal="editable"}, the
[[`execCommand()`](https://w3c.github.io/editing/docs/execCommand/#execcommand%28%29)]{#execCommand
.dfn},
[[`queryCommandEnabled()`](https://w3c.github.io/editing/docs/execCommand/#querycommandenabled%28%29)]{#dom-document-querycommandenabled
.dfn},
[[`queryCommandIndeterm()`](https://w3c.github.io/editing/docs/execCommand/#querycommandindeterm%28%29)]{#dom-document-querycommandindeterm
.dfn},
[[`queryCommandState()`](https://w3c.github.io/editing/docs/execCommand/#querycommandstate%28%29)]{#dom-document-querycommandstate
.dfn},
[[`queryCommandSupported()`](https://w3c.github.io/editing/docs/execCommand/#querycommandsupported%28%29)]{#dom-document-querycommandsupported
.dfn}, and
[[`queryCommandValue()`](https://w3c.github.io/editing/docs/execCommand/#querycommandvalue%28%29)]{#dom-document-querycommandvalue
.dfn} methods, text selections, and the [[delete the
selection](https://w3c.github.io/editing/docs/execCommand/#delete-the-selection)]{#delete-the-selection
.dfn} algorithm are defined in execCommand.
[\[EXECCOMMAND\]](references.html#refsEXECCOMMAND)

#### [6.8.5]{.secno} Spelling and grammar checking[](#spelling-and-grammar-checking){.self-link}

User agents can support the checking of spelling and grammar of editable
text, either in form controls (such as the value of
[`textarea`{#spelling-and-grammar-checking:the-textarea-element}](form-elements.html#the-textarea-element)
elements), or in elements in an [editing
host](#editing-host){#spelling-and-grammar-checking:editing-host} (e.g.
using
[`contenteditable`{#spelling-and-grammar-checking:attr-contenteditable}](#attr-contenteditable)).

For each element, user agents must establish a [default
behavior]{#concept-spellcheck-default .dfn}, either through defaults or
through preferences expressed by the user. There are three possible
default behaviors for each element:

[true-by-default]{#concept-spellcheck-default-true .dfn}
:   The element will be checked for spelling and grammar if its contents
    are editable and spellchecking is not explicitly disabled through
    the
    [`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck}](#attr-spellcheck)
    attribute.

[false-by-default]{#concept-spellcheck-default-false .dfn}
:   The element will never be checked for spelling and grammar unless
    spellchecking is explicitly enabled through the
    [`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-2}](#attr-spellcheck)
    attribute.

[inherit-by-default]{#concept-spellcheck-default-inherit .dfn}
:   The element\'s default behavior is the same as its parent
    element\'s. Elements that have no parent element cannot have this as
    their default behavior.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/spellcheck](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/spellcheck "The spellcheck global attribute is an enumerated attribute that defines whether the element may be checked for spelling errors.")

Support in all current engines.

::: support
[FirefoxYes]{.firefox .yes}[SafariYes]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[OperaYes]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android57+]{.firefox_android .yes}[Safari iOS9.3+]{.safari_ios
.yes}[Chrome Android47+]{.chrome_android .yes}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android37+]{.opera_android .yes}
:::
::::
:::::

The [`spellcheck`]{#attr-spellcheck .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#spelling-and-grammar-checking:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`true`]{#attr-spellcheck-true .dfn dfn-for="html-global/spellcheck"
dfn-type="attr-value"}

[True]{#attr-spellcheck-true-state .dfn}

Spelling and grammar will be checked.

(the empty string)

[`false`]{#attr-spellcheck-false .dfn dfn-for="html-global/spellcheck"
dfn-type="attr-value"}

[False]{#attr-spellcheck-false-state .dfn}

Spelling and grammar will not be checked.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Default]{#attr-spellcheck-default-state .dfn} state. The
default state indicates that the element is to act according to a
default behavior, possibly based on the parent element\'s own
[`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-3}](#attr-spellcheck)
state, as defined below.

------------------------------------------------------------------------

`element`{.variable}`.`[`spellcheck`](#dom-spellcheck){#dom-spellcheck-dev}` [ = ``value`{.variable}` ]`

:   Returns true if the element is to have its spelling and grammar
    checked; otherwise, returns false.

    Can be set, to override the default and set the
    [`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-4}](#attr-spellcheck)
    content attribute.

The [`spellcheck`]{#dom-spellcheck .dfn dfn-for="HTMLElement"
dfn-type="attribute"} IDL attribute, on getting, must return true if the
element\'s
[`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-5}](#attr-spellcheck)
content attribute is in the
[True](#attr-spellcheck-true-state){#spelling-and-grammar-checking:attr-spellcheck-true-state}
state, or if the element\'s
[`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-6}](#attr-spellcheck)
content attribute is in the
[Default](#attr-spellcheck-default-state){#spelling-and-grammar-checking:attr-spellcheck-default-state}
state and the element\'s [default
behavior](#concept-spellcheck-default){#spelling-and-grammar-checking:concept-spellcheck-default}
is
[true-by-default](#concept-spellcheck-default-true){#spelling-and-grammar-checking:concept-spellcheck-default-true},
or if the element\'s
[`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-7}](#attr-spellcheck)
content attribute is in the
[Default](#attr-spellcheck-default-state){#spelling-and-grammar-checking:attr-spellcheck-default-state-2}
state and the element\'s [default
behavior](#concept-spellcheck-default){#spelling-and-grammar-checking:concept-spellcheck-default-2}
is
[inherit-by-default](#concept-spellcheck-default-inherit){#spelling-and-grammar-checking:concept-spellcheck-default-inherit}
and the element\'s parent element\'s
[`spellcheck`{#spelling-and-grammar-checking:dom-spellcheck}](#dom-spellcheck)
IDL attribute would return true; otherwise, if none of those conditions
applies, then the attribute must instead return false.

The
[`spellcheck`{#spelling-and-grammar-checking:dom-spellcheck-2}](#dom-spellcheck)
IDL attribute is not affected by user preferences that override the
[`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-8}](#attr-spellcheck)
content attribute, and therefore might not reflect the actual
spellchecking state.

On setting, if the new value is true, then the element\'s
[`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-9}](#attr-spellcheck)
content attribute must be set to \"`true`\", otherwise it must be set to
\"`false`\".

------------------------------------------------------------------------

User agents should only consider the following pieces of text as
checkable for the purposes of this feature:

- The
  [value](form-control-infrastructure.html#concept-fe-value){#spelling-and-grammar-checking:concept-fe-value}
  of
  [`input`{#spelling-and-grammar-checking:the-input-element}](input.html#the-input-element)
  elements whose
  [`type`{#spelling-and-grammar-checking:attr-input-type}](input.html#attr-input-type)
  attributes are in the
  [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#spelling-and-grammar-checking:text-(type=text)-state-and-search-state-(type=search)},
  [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#spelling-and-grammar-checking:text-(type=text)-state-and-search-state-(type=search)-2},
  [URL](input.html#url-state-(type=url)){#spelling-and-grammar-checking:url-state-(type=url)},
  or
  [Email](input.html#email-state-(type=email)){#spelling-and-grammar-checking:email-state-(type=email)}
  states and that are
  *[mutable](form-control-infrastructure.html#concept-fe-mutable)* (i.e.
  that do not have the
  [`readonly`{#spelling-and-grammar-checking:attr-input-readonly}](input.html#attr-input-readonly)
  attribute specified and that are not
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#spelling-and-grammar-checking:concept-fe-disabled}).
- The
  [value](form-control-infrastructure.html#concept-fe-value){#spelling-and-grammar-checking:concept-fe-value-2}
  of
  [`textarea`{#spelling-and-grammar-checking:the-textarea-element-2}](form-elements.html#the-textarea-element)
  elements that do not have a
  [`readonly`{#spelling-and-grammar-checking:attr-textarea-readonly}](form-elements.html#attr-textarea-readonly)
  attribute and that are not
  [disabled](form-control-infrastructure.html#concept-fe-disabled){#spelling-and-grammar-checking:concept-fe-disabled-2}.
- Text in
  [`Text`{#spelling-and-grammar-checking:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
  nodes that are children of [editing
  hosts](#editing-host){#spelling-and-grammar-checking:editing-host-2}
  or
  [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#spelling-and-grammar-checking:editable
  x-internal="editable"} elements.
- Text in attributes of
  [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#spelling-and-grammar-checking:editable-2
  x-internal="editable"} elements.

For text that is part of a
[`Text`{#spelling-and-grammar-checking:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node, the element with which the text is associated is the element that
is the immediate parent of the first character of the word, sentence, or
other piece of text. For text in attributes, it is the attribute\'s
element. For the values of
[`input`{#spelling-and-grammar-checking:the-input-element-2}](input.html#the-input-element)
and
[`textarea`{#spelling-and-grammar-checking:the-textarea-element-3}](form-elements.html#the-textarea-element)
elements, it is the element itself.

To determine if a word, sentence, or other piece of text in an
applicable element (as defined above) is to have spelling- and
grammar-checking enabled, the UA must use the following algorithm:

1.  If the user has disabled the checking for this text, then the
    checking is disabled.
2.  Otherwise, if the user has forced the checking for this text to
    always be enabled, then the checking is enabled.
3.  Otherwise, if the element with which the text is associated has a
    [`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-10}](#attr-spellcheck)
    content attribute, then: if that attribute is in the
    [True](#attr-spellcheck-true-state){#spelling-and-grammar-checking:attr-spellcheck-true-state-2}
    state, then checking is enabled; otherwise, if that attribute is in
    the
    [False](#attr-spellcheck-false-state){#spelling-and-grammar-checking:attr-spellcheck-false-state}
    state, then checking is disabled.
4.  Otherwise, if there is an ancestor element with a
    [`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-11}](#attr-spellcheck)
    content attribute that is not in the
    [Default](#attr-spellcheck-default-state){#spelling-and-grammar-checking:attr-spellcheck-default-state-3}
    state, then: if the nearest such ancestor\'s
    [`spellcheck`{#spelling-and-grammar-checking:attr-spellcheck-12}](#attr-spellcheck)
    content attribute is in the
    [True](#attr-spellcheck-true-state){#spelling-and-grammar-checking:attr-spellcheck-true-state-3}
    state, then checking is enabled; otherwise, checking is disabled.
5.  Otherwise, if the element\'s [default
    behavior](#concept-spellcheck-default){#spelling-and-grammar-checking:concept-spellcheck-default-3}
    is
    [true-by-default](#concept-spellcheck-default-true){#spelling-and-grammar-checking:concept-spellcheck-default-true-2},
    then checking is enabled.
6.  Otherwise, if the element\'s [default
    behavior](#concept-spellcheck-default){#spelling-and-grammar-checking:concept-spellcheck-default-4}
    is
    [false-by-default](#concept-spellcheck-default-false){#spelling-and-grammar-checking:concept-spellcheck-default-false},
    then checking is disabled.
7.  Otherwise, if the element\'s parent element has *its* checking
    enabled, then checking is enabled.
8.  Otherwise, checking is disabled.

If the checking is enabled for a word/sentence/text, the user agent
should indicate spelling and grammar errors in that text. User agents
should take into account the other semantics given in the document when
suggesting spelling and grammar corrections. User agents may use the
language of the element to determine what spelling and grammar rules to
use, or may use the user\'s preferred language settings. UAs should use
[`input`{#spelling-and-grammar-checking:the-input-element-3}](input.html#the-input-element)
element attributes such as
[`pattern`{#spelling-and-grammar-checking:attr-input-pattern}](input.html#attr-input-pattern)
to ensure that the resulting value is valid, where possible.

If checking is disabled, the user agent should not indicate spelling or
grammar errors for that text.

::: example
The element with ID \"a\" in the following example would be the one used
to determine if the word \"Hello\" is checked for spelling errors. In
this example, it would not be.

``` html
<div contenteditable="true">
 <span spellcheck="false" id="a">Hell</span><em>o!</em>
</div>
```

The element with ID \"b\" in the following example would have checking
enabled (the leading space character in the attribute\'s value on the
[`input`{#spelling-and-grammar-checking:the-input-element-4}](input.html#the-input-element)
element causes the attribute to be ignored, so the ancestor\'s value is
used instead, regardless of the default).

``` bad
<p spellcheck="true">
 <label>Name: <input spellcheck=" false" id="b"></label>
</p>
```
:::

This specification does not define the user interface for spelling and
grammar checkers. A user agent could offer on-demand checking, could
perform continuous checking while the checking is enabled, or could use
other interfaces.

#### [6.8.6]{.secno} Writing suggestions[](#writing-suggestions){.self-link}

User agents offer writing suggestions as users type into editable
regions, either in form controls (e.g., the
[`textarea`{#writing-suggestions:the-textarea-element}](form-elements.html#the-textarea-element)
element) or in elements in an [editing
host](#editing-host){#writing-suggestions:editing-host}.

The [`writingsuggestions`]{#attr-writingsuggestions .dfn
dfn-for="html-global" dfn-type="element-attr"} content attribute is an
[enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#writing-suggestions:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`true`]{#attr-writingsuggestions-true .dfn
dfn-for="html-global/writingsuggestions" dfn-type="attr-value"}

[True]{#attr-writingsuggestions-true-state .dfn}

Writing suggestions should be offered on this element.

(the empty string)

[`false`]{#attr-writingsuggestions-false .dfn
dfn-for="html-global/writingsuggestions" dfn-type="attr-value"}

[False]{#attr-writingsuggestions-false-state .dfn}

Writing suggestions should not be offered on this element.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* is the
[Default]{#attr-writingsuggestions-default-state .dfn} state. The
default state indicates that the element is to act according to a
default behavior, possibly based on the parent element\'s own
[`writingsuggestions`{#writing-suggestions:attr-writingsuggestions}](#attr-writingsuggestions)
state, as defined below.

The attribute\'s *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* is the
[True](#attr-writingsuggestions-true-state){#writing-suggestions:attr-writingsuggestions-true-state}
state.

`element`{.variable}`.`[`writingSuggestions`](#dom-writingsuggestions){#dom-writingsuggestions-dev}` [ = ``value`{.variable}` ]`

:   Returns \"`true`\" if the user agent is to offer writing suggestions
    under the scope of the element; otherwise, returns \"`false`\".

    Can be set, to override the default and set the
    ` `{#writing-suggestions:attr-writingsuggestions-2}[`writingsuggestions`{#writing-suggestions:attr-writingsuggestions-2}](#attr-writingsuggestions)
    content attribute.

The [computed writing suggestions
value]{#computed-writing-suggestions-value .dfn} of a given
`element`{.variable} is determined by running the following steps:

1.  If `element`{.variable}\'s
    [`writingsuggestions`{#writing-suggestions:attr-writingsuggestions-3}](#attr-writingsuggestions)
    content attribute is in the
    [False](#attr-writingsuggestions-false-state){#writing-suggestions:attr-writingsuggestions-false-state}
    state, return \"`false`\".

2.  If `element`{.variable}\'s
    [`writingsuggestions`{#writing-suggestions:attr-writingsuggestions-4}](#attr-writingsuggestions)
    content attribute is in the
    [Default](#attr-writingsuggestions-default-state){#writing-suggestions:attr-writingsuggestions-default-state}
    state, `element`{.variable} has a parent element, and the [computed
    writing suggestions
    value](#computed-writing-suggestions-value){#writing-suggestions:computed-writing-suggestions-value}
    of `element`{.variable}\'s parent element is \"`false`\", then
    return \"`false`\".

3.  Return \"`true`\".

The [`writingSuggestions`]{#dom-writingsuggestions .dfn
dfn-for="HTMLElement" dfn-type="attribute"} getter steps are:

1.  Return
    [this](https://webidl.spec.whatwg.org/#this){#writing-suggestions:this
    x-internal="this"}\'s [computed writing suggestions
    value](#computed-writing-suggestions-value){#writing-suggestions:computed-writing-suggestions-value-2}.

The
[`writingSuggestions`{#writing-suggestions:dom-writingsuggestions}](#dom-writingsuggestions)
IDL attribute is not affected by user preferences that override the
[`writingsuggestions`{#writing-suggestions:attr-writingsuggestions-5}](#attr-writingsuggestions)
content attribute, and therefore might not reflect the actual writing
suggestions state.

The
[`writingSuggestions`{#writing-suggestions:dom-writingsuggestions-2}](#dom-writingsuggestions)
setter steps are:

1.  Set
    [this](https://webidl.spec.whatwg.org/#this){#writing-suggestions:this-2
    x-internal="this"}\'s
    [`writingsuggestions`{#writing-suggestions:attr-writingsuggestions-6}](#attr-writingsuggestions)
    content attribute to the given value.

------------------------------------------------------------------------

User agents should only offer suggestions within an element\'s scope if
the result of running the following algorithm given `element`{.variable}
returns true:

1.  If the user has disabled writing suggestions, then return false.

2.  If none of the following conditions are true:

    - `element`{.variable} is an
      [`input`{#writing-suggestions:the-input-element}](input.html#the-input-element)
      element whose
      [`type`{#writing-suggestions:attr-input-type}](input.html#attr-input-type)
      attribute is in either the
      [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#writing-suggestions:text-(type=text)-state-and-search-state-(type=search)},
      [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#writing-suggestions:text-(type=text)-state-and-search-state-(type=search)-2},
      [Telephone](input.html#telephone-state-(type=tel)){#writing-suggestions:telephone-state-(type=tel)},
      [URL](input.html#url-state-(type=url)){#writing-suggestions:url-state-(type=url)},
      or
      [Email](input.html#email-state-(type=email)){#writing-suggestions:email-state-(type=email)}
      state and is
      *[mutable](form-control-infrastructure.html#concept-fe-mutable)*;

    - `element`{.variable} is a
      [`textarea`{#writing-suggestions:the-textarea-element-2}](form-elements.html#the-textarea-element)
      element that is
      *[mutable](form-control-infrastructure.html#concept-fe-mutable)*;
      or

    - `element`{.variable} is an [editing
      host](#editing-host){#writing-suggestions:editing-host-2} or is
      [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#writing-suggestions:editable
      x-internal="editable"},

    then return false.

3.  If `element`{.variable} has an [inclusive
    ancestor](https://dom.spec.whatwg.org/#concept-tree-inclusive-ancestor){#writing-suggestions:inclusive-ancestor
    x-internal="inclusive-ancestor"} with a
    [`writingsuggestions`{#writing-suggestions:attr-writingsuggestions-7}](#attr-writingsuggestions)
    content attribute that\'s not in the
    [Default](#attr-writingsuggestions-default-state){#writing-suggestions:attr-writingsuggestions-default-state-2}
    and the nearest such ancestor\'s
    [`writingsuggestions`{#writing-suggestions:attr-writingsuggestions-8}](#attr-writingsuggestions)
    content attribute is in the
    [False](#attr-writingsuggestions-false-state){#writing-suggestions:attr-writingsuggestions-false-state-2}
    state, then return false.

4.  Otherwise, return true.

This specification does not define the user interface for writing
suggestions. A user agent could offer on-demand suggestions, continuous
suggestions as the user types, inline suggestions, autofill-like
suggestions in a popup, or could use other interfaces.

#### [6.8.7]{.secno} Autocapitalization[](#autocapitalization){.self-link}

Some methods of entering text, for example virtual keyboards on mobile
devices, and also voice input, often assist users by automatically
capitalizing the first letter of sentences (when composing text in a
language with this convention). A virtual keyboard that implements
autocapitalization might automatically switch to showing uppercase
letters (but allow the user to toggle it back to lowercase) when a
letter that should be autocapitalized is about to be typed. Other types
of input, for example voice input, may perform autocapitalization in a
way that does not give users an option to intervene first. The
[`autocapitalize`{#autocapitalization:attr-autocapitalize}](#attr-autocapitalize)
attribute allows authors to control such behavior.

The
[`autocapitalize`{#autocapitalization:attr-autocapitalize-2}](#attr-autocapitalize)
attribute, as typically implemented, does not affect behavior when
typing on a physical keyboard. (For this reason, as well as the ability
for users to override the autocapitalization behavior in some cases or
edit the text after initial input, the attribute must not be relied on
for any sort of input validation.)

The
[`autocapitalize`{#autocapitalization:attr-autocapitalize-3}](#attr-autocapitalize)
attribute can be used on an [editing
host](#editing-host){#autocapitalization:editing-host} to control
autocapitalization behavior for the hosted editable region, on an
[`input`{#autocapitalization:the-input-element}](input.html#the-input-element)
or
[`textarea`{#autocapitalization:the-textarea-element}](form-elements.html#the-textarea-element)
element to control the behavior for inputting text into that element, or
on a
[`form`{#autocapitalization:the-form-element}](forms.html#the-form-element)
element to control the default behavior for all
[autocapitalize-and-autocorrect inheriting
elements](forms.html#category-autocapitalize){#autocapitalization:category-autocapitalize}
associated with the
[`form`{#autocapitalization:the-form-element-2}](forms.html#the-form-element)
element.

The
[`autocapitalize`{#autocapitalization:attr-autocapitalize-4}](#attr-autocapitalize)
attribute never causes autocapitalization to be enabled for
[`input`{#autocapitalization:the-input-element-2}](input.html#the-input-element)
elements whose
[`type`{#autocapitalization:attr-input-type}](input.html#attr-input-type)
attribute is in one of the
[URL](input.html#url-state-(type=url)){#autocapitalization:url-state-(type=url)},
[Email](input.html#email-state-(type=email)){#autocapitalization:email-state-(type=email)},
or
[Password](input.html#password-state-(type=password)){#autocapitalization:password-state-(type=password)}
states. (This behavior is included in the [used autocapitalization
hint](#used-autocapitalization-hint){#autocapitalization:used-autocapitalization-hint}
algorithm below.)

The autocapitalization processing model is based on selecting among five
[autocapitalization hints]{#autocapitalization-hint .dfn}, defined as
follows:

[default]{#autocap-hint-default .dfn}
:   The user agent and input method should make their own determination
    of whether or not to enable autocapitalization.

[none]{#autocap-hint-none .dfn}
:   No autocapitalization should be applied (all letters should default
    to lowercase).

[sentences]{#autocap-hint-sentences .dfn}
:   The first letter of each sentence should default to a capital
    letter; all other letters should default to lowercase.

[words]{#autocap-hint-words .dfn}
:   The first letter of each word should default to a capital letter;
    all other letters should default to lowercase.

[characters]{#autocap-hint-characters .dfn}

:   All letters should default to uppercase.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/autocapitalize](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autocapitalize "The autocapitalize global attribute is an enumerated attribute that controls whether and how text input is automatically capitalized as it is entered/edited by the user.")

Support in all current engines.

::: support
[Firefox111+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome43+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet Explorer?]{.ie .unknown}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`autocapitalize`]{#attr-autocapitalize .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#autocapitalization:enumerated-attribute}
whose states are the possible [autocapitalization
hints](#autocapitalization-hint){#autocapitalization:autocapitalization-hint}.
The [autocapitalization
hint](#autocapitalization-hint){#autocapitalization:autocapitalization-hint-2}
specified by the attribute\'s state combines with other considerations
to form the [used autocapitalization
hint](#used-autocapitalization-hint){#autocapitalization:used-autocapitalization-hint-2},
which informs the behavior of the user agent. The keywords for this
attribute and their state mappings are as follows:

Keyword

State

[`off`]{#attr-autocapitalize-off .dfn
dfn-for="html-global/autocapitalize" dfn-type="attr-value"}

[none](#autocap-hint-none){#autocapitalization:autocap-hint-none}

[`none`]{#attr-autocapitalize-none .dfn
dfn-for="html-global/autocapitalize" dfn-type="attr-value"}

[`on`]{#attr-autocapitalize-on .dfn dfn-for="html-global/autocapitalize"
dfn-type="attr-value"}

[sentences](#autocap-hint-sentences){#autocapitalization:autocap-hint-sentences}

[`sentences`]{#attr-autocapitalize-sentences .dfn
dfn-for="html-global/autocapitalize" dfn-type="attr-value"}

[`words`]{#attr-autocapitalize-words .dfn
dfn-for="html-global/autocapitalize" dfn-type="attr-value"}

[words](#autocap-hint-words){#autocapitalization:autocap-hint-words}

[`characters`]{#attr-autocapitalize-characters .dfn
dfn-for="html-global/autocapitalize" dfn-type="attr-value"}

[characters](#autocap-hint-characters){#autocapitalization:autocap-hint-characters}

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* is the
[default](#autocap-hint-default){#autocapitalization:autocap-hint-default}
state, and its *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* is the
[sentences](#autocap-hint-sentences){#autocapitalization:autocap-hint-sentences-2}
state.

`element`{.variable}`.`[`autocapitalize`](#dom-autocapitalize){#dom-autocapitalize-dev}` [ = ``value`{.variable}` ]`

:   Returns the current autocapitalization state for the element, or an
    empty string if it hasn\'t been set. Note that for
    [`input`{#autocapitalization:the-input-element-3}](input.html#the-input-element)
    and
    [`textarea`{#autocapitalization:the-textarea-element-2}](form-elements.html#the-textarea-element)
    elements that inherit their state from a
    [`form`{#autocapitalization:the-form-element-3}](forms.html#the-form-element)
    element, this will return the autocapitalization state of the
    [`form`{#autocapitalization:the-form-element-4}](forms.html#the-form-element)
    element, but for an element in an editable region, this will not
    return the autocapitalization state of the editing host (unless this
    element is, in fact, the [editing
    host](#editing-host){#autocapitalization:editing-host-2}).

    Can be set, to set the
    [`autocapitalize`{#autocapitalization:attr-autocapitalize-5}](#attr-autocapitalize)
    content attribute (and thereby change the autocapitalization
    behavior for the element).

To compute the [own autocapitalization
hint]{#own-autocapitalization-hint .dfn} of an element
`element`{.variable}, run the following steps:

1.  If the
    [`autocapitalize`{#autocapitalization:attr-autocapitalize-6}](#attr-autocapitalize)
    content attribute is present on `element`{.variable}, and its value
    is not the empty string, return the state of the attribute.

2.  If `element`{.variable} is an [autocapitalize-and-autocorrect
    inheriting
    element](forms.html#category-autocapitalize){#autocapitalization:category-autocapitalize-2}
    and has a non-null [form
    owner](form-control-infrastructure.html#form-owner){#autocapitalization:form-owner},
    return the [own autocapitalization
    hint](#own-autocapitalization-hint){#autocapitalization:own-autocapitalization-hint}
    of `element`{.variable}\'s [form
    owner](form-control-infrastructure.html#form-owner){#autocapitalization:form-owner-2}.

3.  Return
    [default](#autocap-hint-default){#autocapitalization:autocap-hint-default-2}.

The [`autocapitalize`]{#dom-autocapitalize .dfn dfn-for="HTMLElement"
dfn-type="attribute"} getter steps are to:

1.  Let `state`{.variable} be the [own autocapitalization
    hint](#own-autocapitalization-hint){#autocapitalization:own-autocapitalization-hint-2}
    of
    [this](https://webidl.spec.whatwg.org/#this){#autocapitalization:this
    x-internal="this"}.

2.  If `state`{.variable} is
    [default](#autocap-hint-default){#autocapitalization:autocap-hint-default-3},
    then return the empty string.

3.  If `state`{.variable} is
    [none](#autocap-hint-none){#autocapitalization:autocap-hint-none-2},
    then return
    \"[`none`{#autocapitalization:attr-autocapitalize-none}](#attr-autocapitalize-none)\".

4.  If `state`{.variable} is
    [sentences](#autocap-hint-sentences){#autocapitalization:autocap-hint-sentences-3},
    then return
    \"[`sentences`{#autocapitalization:attr-autocapitalize-sentences}](#attr-autocapitalize-sentences)\".

5.  Return the keyword value corresponding to `state`{.variable}.

The
[`autocapitalize`{#autocapitalization:dom-autocapitalize}](#dom-autocapitalize)
setter steps are to set the
[`autocapitalize`{#autocapitalization:attr-autocapitalize-7}](#attr-autocapitalize)
content attribute to the given value.

------------------------------------------------------------------------

User agents that support customizable autocapitalization behavior for a
text input method and wish to allow web developers to control this
functionality should, during text input into an element, compute the
[used autocapitalization hint]{#used-autocapitalization-hint .dfn} for
the element. This will be an [autocapitalization
hint](#autocapitalization-hint){#autocapitalization:autocapitalization-hint-3}
that describes the recommended autocapitalization behavior for text
input into the element.

User agents or input methods may choose to ignore or override the [used
autocapitalization
hint](#used-autocapitalization-hint){#autocapitalization:used-autocapitalization-hint-3}
in certain circumstances.

The [used autocapitalization
hint](#used-autocapitalization-hint){#autocapitalization:used-autocapitalization-hint-4}
for an element `element`{.variable} is computed using the following
algorithm:

1.  If `element`{.variable} is an
    [`input`{#autocapitalization:the-input-element-4}](input.html#the-input-element)
    element whose
    [`type`{#autocapitalization:attr-input-type-2}](input.html#attr-input-type)
    attribute is in one of the
    [URL](input.html#url-state-(type=url)){#autocapitalization:url-state-(type=url)-2},
    [Email](input.html#email-state-(type=email)){#autocapitalization:email-state-(type=email)-2},
    or
    [Password](input.html#password-state-(type=password)){#autocapitalization:password-state-(type=password)-2}
    states, then return
    [default](#autocap-hint-default){#autocapitalization:autocap-hint-default-4}.

2.  If `element`{.variable} is an
    [`input`{#autocapitalization:the-input-element-5}](input.html#the-input-element)
    element or a
    [`textarea`{#autocapitalization:the-textarea-element-3}](form-elements.html#the-textarea-element)
    element, then return `element`{.variable}\'s [own autocapitalization
    hint](#own-autocapitalization-hint){#autocapitalization:own-autocapitalization-hint-3}.

3.  If `element`{.variable} is an [editing
    host](#editing-host){#autocapitalization:editing-host-3} or an
    [editable](https://w3c.github.io/editing/docs/execCommand/#editable){#autocapitalization:editable
    x-internal="editable"} element, then return the [own
    autocapitalization
    hint](#own-autocapitalization-hint){#autocapitalization:own-autocapitalization-hint-4}
    of the [editing host
    of](https://w3c.github.io/editing/docs/execCommand/#editing-host-of){#autocapitalization:editing-host-of
    x-internal="editing-host-of"} `element`{.variable}.

4.  [Assert](https://infra.spec.whatwg.org/#assert){#autocapitalization:assert
    x-internal="assert"}: this step is never reached, since text input
    only occurs in elements that meet one of the above criteria.

#### [6.8.8]{.secno} Autocorrection[](#autocorrection){.self-link}

Some methods of entering text assist users by automatically correcting
misspelled words while typing, a process also known as autocorrection.
User agents can support autocorrection of editable text, either in form
controls (such as the value of
[`textarea`{#autocorrection:the-textarea-element}](form-elements.html#the-textarea-element)
elements), or in elements in an [editing
host](#editing-host){#autocorrection:editing-host} (e.g., using
[`contenteditable`{#autocorrection:attr-contenteditable}](#attr-contenteditable)).
Autocorrection may be accompanied by user interfaces indicating that
text is about to be autocorrected or has been autocorrected, and is
commonly performed when inserting punctuation characters, spaces, or new
paragraphs after misspelled words. The
[`autocorrect`{#autocorrection:attr-autocorrect}](#attr-autocorrect)
attribute allows authors to control such behavior.

The
[`autocorrect`{#autocorrection:attr-autocorrect-2}](#attr-autocorrect)
attribute can be used on an editing host to control autocorrection
behavior for the hosted editable region, on an
[`input`{#autocorrection:the-input-element}](input.html#the-input-element)
or
[`textarea`{#autocorrection:the-textarea-element-2}](form-elements.html#the-textarea-element)
element to control the behavior when inserting text into that element,
or on a
[`form`{#autocorrection:the-form-element}](forms.html#the-form-element)
element to control the default behavior for all
[autocapitalize-and-autocorrect inheriting
elements](forms.html#category-autocapitalize){#autocorrection:category-autocapitalize}
associated with the
[`form`{#autocorrection:the-form-element-2}](forms.html#the-form-element)
element.

The
[`autocorrect`{#autocorrection:attr-autocorrect-3}](#attr-autocorrect)
attribute never causes autocorrection to be enabled for
[`input`{#autocorrection:the-input-element-2}](input.html#the-input-element)
elements whose
[`type`{#autocorrection:attr-input-type}](input.html#attr-input-type)
attribute is in one of the
[URL](input.html#url-state-(type=url)){#autocorrection:url-state-(type=url)},
[Email](input.html#email-state-(type=email)){#autocorrection:email-state-(type=email)},
or
[Password](input.html#password-state-(type=password)){#autocorrection:password-state-(type=password)}
states. (This behavior is included in the [used autocorrection
state](#used-autocorrection-state){#autocorrection:used-autocorrection-state}
algorithm below.)

The [`autocorrect`]{#attr-autocorrect .dfn dfn-for="html-global"
dfn-type="element-attr"} attribute is an enumerated attribute with the
following keywords and states:

Keyword

State

Brief description

[`on`]{#attr-autocorrect-on .dfn dfn-for="html-global/autocorrect"
dfn-type="attr-value"}

[on]{#concept-autocorrection-on .dfn}

The user agent is permitted to automatically correct spelling errors
while the user types. Whether spelling is automatically corrected while
typing left is for the user agent to decide, and may depend on the
element as well as the user\'s preferences.

(the empty string)

[`off`]{#attr-autocorrect-off .dfn dfn-for="html-global/autocorrect"
dfn-type="attr-value"}

[off]{#concept-autocorrection-off .dfn}

The user agent is not allowed to automatically correct spelling while
the user types.

The attribute\'s *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* and *[missing
value default](common-microsyntaxes.html#missing-value-default)* are
both the
[on](#concept-autocorrection-on){#autocorrection:concept-autocorrection-on}
state.

The [`autocorrect`]{#dom-autocorrect .dfn dfn-for="HTMLElement"
dfn-type="attribute"} getter steps are: return true if the element\'s
[used autocorrection
state](#used-autocorrection-state){#autocorrection:used-autocorrection-state-2}
is
[on](#concept-autocorrection-on){#autocorrection:concept-autocorrection-on-2}
and false if the element\'s [used autocorrection
state](#used-autocorrection-state){#autocorrection:used-autocorrection-state-3}
is
[off](#concept-autocorrection-off){#autocorrection:concept-autocorrection-off}.
The setter steps are: if the given value is true, then the element\'s
[`autocorrect`{#autocorrection:attr-autocorrect-4}](#attr-autocorrect)
attribute must be set to \"`on`\"; otherwise it must be set to
\"`off`\".

To compute the [used autocorrection state]{#used-autocorrection-state
.dfn} of an element `element`{.variable}, run these steps:

1.  If `element`{.variable} is an
    [`input`{#autocorrection:the-input-element-3}](input.html#the-input-element)
    element whose
    [`type`{#autocorrection:attr-input-type-2}](input.html#attr-input-type)
    attribute is in one of the
    [URL](input.html#url-state-(type=url)){#autocorrection:url-state-(type=url)-2},
    [Email](input.html#email-state-(type=email)){#autocorrection:email-state-(type=email)-2},
    or
    [Password](input.html#password-state-(type=password)){#autocorrection:password-state-(type=password)-2}
    states, then return
    [off](#concept-autocorrection-off){#autocorrection:concept-autocorrection-off-2}.

2.  If the
    [`autocorrect`{#autocorrection:attr-autocorrect-5}](#attr-autocorrect)
    content attribute is present on `element`{.variable}, then return
    the state of the attribute.

3.  If `element`{.variable} is an [autocapitalize-and-autocorrect
    inheriting
    element](forms.html#category-autocapitalize){#autocorrection:category-autocapitalize-2}
    and has a non-null [form
    owner](form-control-infrastructure.html#form-owner){#autocorrection:form-owner},
    then return the state of `element`{.variable}\'s [form
    owner](form-control-infrastructure.html#form-owner){#autocorrection:form-owner-2}\'s
    [`autocorrect`{#autocorrection:attr-autocorrect-6}](#attr-autocorrect)
    attribute.

4.  Return
    [on](#concept-autocorrection-on){#autocorrection:concept-autocorrection-on-3}.

`element`{.variable} . [`autocorrect`{#dom-autocorrect-dev}](#dom-autocorrect)
:   Returns the autocorrection behavior of the element. Note that for
    [autocapitalize-and-autocorrect inheriting
    elements](forms.html#category-autocapitalize){#autocorrection:category-autocapitalize-3}
    that inherit their state from a
    [`form`{#autocorrection:the-form-element-3}](forms.html#the-form-element)
    element, this will return the autocorrection behavior of the
    [`form`{#autocorrection:the-form-element-4}](forms.html#the-form-element)
    element, but for an element in an editable region, this will not
    return the autocorrection behavior of the [editing
    host](#editing-host){#autocorrection:editing-host-2} (unless this
    element is, in fact, the [editing
    host](#editing-host){#autocorrection:editing-host-3}).

`element`{.variable} . [`autocorrect`{#autocorrection:dom-autocorrect}](#dom-autocorrect) = `value`{.variable}

:   Updates the
    [`autocorrect`{#autocorrection:attr-autocorrect-7}](#attr-autocorrect)
    content attribute (and thereby changes the autocorrection behavior
    of the element).

::: example
The
[`input`{#autocorrection:the-input-element-4}](input.html#the-input-element)
element in the following example would not allow autocorrection, since
it does not have an
[`autocorrect`{#autocorrection:attr-autocorrect-8}](#attr-autocorrect)
content attribute and therefore inherits from the
[`form`{#autocorrection:the-form-element-5}](forms.html#the-form-element)
element, which has an attribute of
\"[`off`{#autocorrection:attr-autocorrect-off}](#attr-autocorrect-off)\".
However, the
[`textarea`{#autocorrection:the-textarea-element-3}](form-elements.html#the-textarea-element)
element would allow autocorrection, since it has an
[`autocorrect`{#autocorrection:attr-autocorrect-9}](#attr-autocorrect)
content attribute with a value of
\"[`on`{#autocorrection:attr-autocorrect-on}](#attr-autocorrect-on)\".

``` html
<form autocorrect="off">
 <input type="search">
 <textarea autocorrect="on"></textarea>
</form>
```
:::

#### [6.8.9]{.secno} Input modalities: the [`inputmode`{#input-modalities:-the-inputmode-attribute:attr-inputmode}](#attr-inputmode) attribute[](#input-modalities:-the-inputmode-attribute){.self-link} {#input-modalities:-the-inputmode-attribute}

User agents can support the
[`inputmode`{#input-modalities:-the-inputmode-attribute:attr-inputmode-2}](#attr-inputmode)
attribute on form controls (such as the value of
[`textarea`{#input-modalities:-the-inputmode-attribute:the-textarea-element}](form-elements.html#the-textarea-element)
elements), or in elements in an [editing
host](#editing-host){#input-modalities:-the-inputmode-attribute:editing-host}
(e.g., using
[`contenteditable`{#input-modalities:-the-inputmode-attribute:attr-contenteditable}](#attr-contenteditable)).

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/inputmode](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inputmode "The inputmode global attribute is an enumerated attribute that hints at the type of data that might be entered by the user while editing the element or its contents. This allows a browser to display an appropriate virtual keyboard.")

Support in all current engines.

::: support
[Firefox95+]{.firefox .yes}[SafariNo]{.safari .no}[Chrome66+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android79+]{.firefox_android .yes}[Safari iOS12.2+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`inputmode`]{#attr-inputmode .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#input-modalities:-the-inputmode-attribute:enumerated-attribute}
that specifies what kind of input mechanism would be most helpful for
users entering content.

Keyword

Description

[`none`]{#attr-inputmode-keyword-none .dfn
dfn-for="html-global/inputmode" dfn-type="attr-value"}

The user agent should not display a virtual keyboard. This keyword is
useful for content that renders its own keyboard control.

[`text`]{#attr-inputmode-keyword-text .dfn
dfn-for="html-global/inputmode" dfn-type="attr-value"}

The user agent should display a virtual keyboard capable of text input
in the user\'s locale.

[`tel`]{#attr-inputmode-keyword-tel .dfn dfn-for="html-global/inputmode"
dfn-type="attr-value"}

The user agent should display a virtual keyboard capable of telephone
number input. This should including keys for the digits 0 to 9, the
\"#\" character, and the \"\*\" character. In some locales, this can
also include alphabetic mnemonic labels (e.g., in the US, the key
labeled \"2\" is historically also labeled with the letters A, B, and
C).

[`url`]{#attr-inputmode-keyword-url .dfn dfn-for="html-global/inputmode"
dfn-type="attr-value"}

The user agent should display a virtual keyboard capable of text input
in the user\'s locale, with keys for aiding in the input of
[URLs](https://url.spec.whatwg.org/#concept-url){#input-modalities:-the-inputmode-attribute:url
x-internal="url"}, such as that for the \"/\" and \".\" characters and
for quick input of strings commonly found in domain names such as
\"www.\" or \".com\".

[`email`]{#attr-inputmode-keyword-email .dfn
dfn-for="html-global/inputmode" dfn-type="attr-value"}

The user agent should display a virtual keyboard capable of text input
in the user\'s locale, with keys for aiding in the input of email
addresses, such as that for the \"@\" character and the \".\" character.

[`numeric`]{#attr-inputmode-keyword-numeric .dfn
dfn-for="html-global/inputmode" dfn-type="attr-value"}

The user agent should display a virtual keyboard capable of numeric
input. This keyword is useful for PIN entry.

[`decimal`]{#attr-inputmode-keyword-decimal .dfn
dfn-for="html-global/inputmode" dfn-type="attr-value"}

The user agent should display a virtual keyboard capable of fractional
numeric input. Numeric keys and the format separator for the locale
should be shown.

[`search`]{#attr-inputmode-keyword-search .dfn
dfn-for="html-global/inputmode" dfn-type="attr-value"}

The user agent should display a virtual keyboard optimized for search.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/inputMode](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/inputMode "The HTMLElement property inputMode reflects the value of the element's inputmode attribute.")

Support in all current engines.

::: support
[Firefox95+]{.firefox .yes}[Safari12.1+]{.safari
.yes}[Chrome66+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android79+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`inputMode`]{#dom-inputmode .dfn dfn-for="ElementContentEditable"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#input-modalities:-the-inputmode-attribute:reflect}
the
[`inputmode`{#input-modalities:-the-inputmode-attribute:attr-inputmode-3}](#attr-inputmode)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#input-modalities:-the-inputmode-attribute:limited-to-only-known-values}.

When
[`inputmode`{#input-modalities:-the-inputmode-attribute:attr-inputmode-4}](#attr-inputmode)
is unspecified (or is in a state not supported by the user agent), the
user agent should determine the default virtual keyboard to be shown.
Contextual information such as the input
[`type`{#input-modalities:-the-inputmode-attribute:attr-input-type}](input.html#attr-input-type)
or
[`pattern`{#input-modalities:-the-inputmode-attribute:attr-input-pattern}](input.html#attr-input-pattern)
attributes should be used to determine which type of virtual keyboard
should be presented to the user.

#### [6.8.10]{.secno} Input modalities: the [`enterkeyhint`{#input-modalities:-the-enterkeyhint-attribute:attr-enterkeyhint}](#attr-enterkeyhint) attribute[](#input-modalities:-the-enterkeyhint-attribute){.self-link} {#input-modalities:-the-enterkeyhint-attribute}

User agents can support the
[`enterkeyhint`{#input-modalities:-the-enterkeyhint-attribute:attr-enterkeyhint-2}](#attr-enterkeyhint)
attribute on form controls (such as the value of
[`textarea`{#input-modalities:-the-enterkeyhint-attribute:the-textarea-element}](form-elements.html#the-textarea-element)
elements), or in elements in an [editing
host](#editing-host){#input-modalities:-the-enterkeyhint-attribute:editing-host}
(e.g., using
[`contenteditable`{#input-modalities:-the-enterkeyhint-attribute:attr-contenteditable}](#attr-contenteditable)).

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Global_attributes/enterkeyhint](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/enterkeyhint "The enterkeyhint global attribute is an enumerated attribute defining what action label (or icon) to present for the enter key on virtual keyboards.")

Support in all current engines.

::: support
[Firefox94+]{.firefox .yes}[Safari13.1+]{.safari
.yes}[Chrome77+]{.chrome .yes}

------------------------------------------------------------------------

[Opera66+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android57+]{.opera_android .yes}
:::
::::
:::::

The [`enterkeyhint`]{#attr-enterkeyhint .dfn dfn-for="html-global"
dfn-type="element-attr"} content attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#input-modalities:-the-enterkeyhint-attribute:enumerated-attribute}
that specifies what action label (or icon) to present for the enter key
on virtual keyboards. This allows authors to customize the presentation
of the enter key in order to make it more helpful for users.

Keyword

Description

[`enter`]{#attr-enterkeyhint-keyword-enter .dfn
dfn-for="html-global/enterkeyhint" dfn-type="attr-value"}

The user agent should present a cue for the operation \'enter\',
typically inserting a new line.

[`done`]{#attr-enterkeyhint-keyword-done .dfn
dfn-for="html-global/enterkeyhint" dfn-type="attr-value"}

The user agent should present a cue for the operation \'done\',
typically meaning there is nothing more to input and the input method
editor (IME) will be closed.

[`go`]{#attr-enterkeyhint-keyword-go .dfn
dfn-for="html-global/enterkeyhint" dfn-type="attr-value"}

The user agent should present a cue for the operation \'go\', typically
meaning to take the user to the target of the text they typed.

[`next`]{#attr-enterkeyhint-keyword-next .dfn
dfn-for="html-global/enterkeyhint" dfn-type="attr-value"}

The user agent should present a cue for the operation \'next\',
typically taking the user to the next field that will accept text.

[`previous`]{#attr-enterkeyhint-keyword-previous .dfn
dfn-for="html-global/enterkeyhint" dfn-type="attr-value"}

The user agent should present a cue for the operation \'previous\',
typically taking the user to the previous field that will accept text.

[`search`]{#attr-enterkeyhint-keyword-search .dfn
dfn-for="html-global/enterkeyhint" dfn-type="attr-value"}

The user agent should present a cue for the operation \'search\',
typically taking the user to the results of searching for the text they
have typed.

[`send`]{#attr-enterkeyhint-keyword-send .dfn
dfn-for="html-global/enterkeyhint" dfn-type="attr-value"}

The user agent should present a cue for the operation \'send\',
typically delivering the text to its target.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLElement/enterKeyHint](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/enterKeyHint "The enterKeyHint property is an enumerated property defining what action label (or icon) to present for the enter key on virtual keyboards. It reflects the enterkeyhint HTML global attribute and is an enumerated property, only accepting the following values as a string:")

Support in all current engines.

::: support
[Firefox94+]{.firefox .yes}[Safari13.1+]{.safari
.yes}[Chrome77+]{.chrome .yes}

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

The [`enterKeyHint`]{#dom-enterkeyhint .dfn
dfn-for="ElementContentEditable" dfn-type="attribute"} IDL attribute
must
[reflect](common-dom-interfaces.html#reflect){#input-modalities:-the-enterkeyhint-attribute:reflect}
the
[`enterkeyhint`{#input-modalities:-the-enterkeyhint-attribute:attr-enterkeyhint-3}](#attr-enterkeyhint)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#input-modalities:-the-enterkeyhint-attribute:limited-to-only-known-values}.

When
[`enterkeyhint`{#input-modalities:-the-enterkeyhint-attribute:attr-enterkeyhint-4}](#attr-enterkeyhint)
is unspecified (or is in a state not supported by the user agent), the
user agent should determine the default action label (or icon) to
present. Contextual information such as the
[`inputmode`{#input-modalities:-the-enterkeyhint-attribute:attr-inputmode}](#attr-inputmode),
[`type`{#input-modalities:-the-enterkeyhint-attribute:attr-input-type}](input.html#attr-input-type),
or
[`pattern`{#input-modalities:-the-enterkeyhint-attribute:attr-input-pattern}](input.html#attr-input-pattern)
attributes should be used to determine which action label (or icon) to
present on the virtual keyboard.

### [6.9]{.secno} Find-in-page[](#find-in-page){.self-link}

#### [6.9.1]{.secno} Introduction[](#introduction-10){.self-link} {#introduction-10}

This section defines [find-in-page]{#find-in-page-2 .dfn} --- a common
user-agent mechanism which allows users to search through the contents
of the page for particular information.

Access to the
[find-in-page](#find-in-page-2){#introduction-10:find-in-page-2} feature
is provided via a [find-in-page interface]{#find-in-page-interface
.dfn}. This is a user-agent provided user interface, which allows the
user to specify input and the parameters of the search. This interface
can appear as a result of a shortcut or a menu selection.

A combination of text input and settings in the [find-in-page
interface](#find-in-page-interface){#introduction-10:find-in-page-interface}
represents the user [query]{#fip-query .dfn}. This typically includes
the text that the user wants to search for, as well as optional settings
(e.g., the ability to restrict the search to whole words only).

The user-agent processes page contents for a given
[query](#fip-query){#introduction-10:fip-query}, and identifies zero or
more [matches]{#fip-matches .dfn}, which are content ranges that satisfy
the user [query](#fip-query){#introduction-10:fip-query-2}.

One of the [matches](#fip-matches){#introduction-10:fip-matches} is
identified to the user as the [active match]{#fip-active-match .dfn}. It
is highlighted and scrolled into view. The user can navigate through the
[matches](#fip-matches){#introduction-10:fip-matches-2} by advancing the
[active match](#fip-active-match){#introduction-10:fip-active-match}
using the [find-in-page
interface](#find-in-page-interface){#introduction-10:find-in-page-interface-2}.

[Issue #3539](https://github.com/whatwg/html/issues/3539) tracks
standardizing how
[find-in-page](#find-in-page-2){#introduction-10:find-in-page-2-2}
underlies the currently-unspecified `window.find()` API.

#### [6.9.2]{.secno} Interaction with [`details`{#interaction-with-details-and-hidden=until-found:the-details-element}](interactive-elements.html#the-details-element) and [`hidden=until-found`{#interaction-with-details-and-hidden=until-found:attr-hidden-until-found-state}](#attr-hidden-until-found-state)[](#interaction-with-details-and-hidden=until-found){.self-link} {#interaction-with-details-and-hidden=until-found}

When find-in-page begins searching for matches, all
[`details`{#interaction-with-details-and-hidden=until-found:the-details-element-2}](interactive-elements.html#the-details-element)
elements in the page which do not have their
[`open`{#interaction-with-details-and-hidden=until-found:attr-details-open}](interactive-elements.html#attr-details-open)
attribute set should have the [skipped
contents](https://drafts.csswg.org/css-contain/#skips-its-contents){#interaction-with-details-and-hidden=until-found:skips-its-contents
x-internal="skips-its-contents"} of their second slot become accessible,
without modifying the
[`open`{#interaction-with-details-and-hidden=until-found:attr-details-open-2}](interactive-elements.html#attr-details-open)
attribute, in order to make find-in-page able to search through it.
Similarly, all HTML elements with the
[`hidden`{#interaction-with-details-and-hidden=until-found:attr-hidden}](#attr-hidden)
attribute in the [Hidden Until
Found](#attr-hidden-until-found-state){#interaction-with-details-and-hidden=until-found:attr-hidden-until-found-state-2}
state should have their [skipped
contents](https://drafts.csswg.org/css-contain/#skips-its-contents){#interaction-with-details-and-hidden=until-found:skips-its-contents-2
x-internal="skips-its-contents"} become accessible without modifying the
[`hidden`{#interaction-with-details-and-hidden=until-found:attr-hidden-2}](#attr-hidden)
attribute in order to make find-in-page able to search through them.
After find-in-page finishes searching for matches, the
[`details`{#interaction-with-details-and-hidden=until-found:the-details-element-3}](interactive-elements.html#the-details-element)
elements and the elements with the
[`hidden`{#interaction-with-details-and-hidden=until-found:attr-hidden-3}](#attr-hidden)
attribute in the [Hidden Until
Found](#attr-hidden-until-found-state){#interaction-with-details-and-hidden=until-found:attr-hidden-until-found-state-3}
state should have their contents become skipped again. This entire
process must happen synchronously (and so is not observable to users or
to author code). [\[CSSCONTAIN\]](references.html#refsCSSCONTAIN)

When find-in-page chooses a new [active
match](#fip-active-match){#interaction-with-details-and-hidden=until-found:fip-active-match},
perform the following steps:

1.  Let `node`{.variable} be the first node in the [active
    match](#fip-active-match){#interaction-with-details-and-hidden=until-found:fip-active-match-2}.

2.  [Queue a global
    task](webappapis.html#queue-a-global-task){#interaction-with-details-and-hidden=until-found:queue-a-global-task}
    on the [user interaction task
    source](webappapis.html#user-interaction-task-source){#interaction-with-details-and-hidden=until-found:user-interaction-task-source}
    given `node`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#interaction-with-details-and-hidden=until-found:concept-relevant-global}
    to run the following steps:

    1.  Run the [ancestor details revealing
        algorithm](interactive-elements.html#ancestor-details-revealing-algorithm){#interaction-with-details-and-hidden=until-found:ancestor-details-revealing-algorithm}
        on `node`{.variable}.

    2.  Run the [ancestor hidden-until-found revealing
        algorithm](#ancestor-hidden-until-found-revealing-algorithm){#interaction-with-details-and-hidden=until-found:ancestor-hidden-until-found-revealing-algorithm}
        on `node`{.variable}.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#interaction-with-details-and-hidden=until-found:tracking-vector
.tracking-vector x-internal="tracking-vector"} When find-in-page
auto-expands a
[`details`{#interaction-with-details-and-hidden=until-found:the-details-element-4}](interactive-elements.html#the-details-element)
element like this, it will fire a
[`toggle`{#interaction-with-details-and-hidden=until-found:event-toggle}](indices.html#event-toggle)
event. As with the separate
[`scroll`{#interaction-with-details-and-hidden=until-found:event-scroll}](https://drafts.csswg.org/cssom-view/#eventdef-document-scroll){x-internal="event-scroll"}
event that find-in-page fires, this event could be used by the page to
discover what the user is typing into the find-in-page dialog. If the
page creates a tiny scrollable area with the current search term and
every possible next character the user could type separated by a gap,
and observes which one the browser scrolls to, it can add that character
to the search term and update the scrollable area to incrementally build
the search term. By wrapping each possible next match in a closed
[`details`{#interaction-with-details-and-hidden=until-found:the-details-element-5}](interactive-elements.html#the-details-element)
element, the page could listen to
[`toggle`{#interaction-with-details-and-hidden=until-found:event-toggle-2}](indices.html#event-toggle)
events instead of
[`scroll`{#interaction-with-details-and-hidden=until-found:event-scroll-2}](https://drafts.csswg.org/cssom-view/#eventdef-document-scroll){x-internal="event-scroll"}
events. This attack could be addressed for both events by not acting on
every character the user types into the find-in-page dialog.

#### [6.9.3]{.secno} Interaction with selection[](#interaction-with-selection){.self-link}

The find-in-page process is invoked in the context of a document, and
may have an effect on the
[selection](https://w3c.github.io/selection-api/#dfn-selection){#interaction-with-selection:document-selection
x-internal="document-selection"} of that document. Specifically, the
range that defines the [active
match](#fip-active-match){#interaction-with-selection:fip-active-match}
can dictate the current selection. These selection updates, however, can
happen at different times during the find-in-page process (e.g. upon the
[find-in-page
interface](#find-in-page-interface){#interaction-with-selection:find-in-page-interface}
dismissal or upon a change in the [active
match](#fip-active-match){#interaction-with-selection:fip-active-match-2}
range).

### [6.10]{.secno} Close requests and close watchers[](#close-requests-and-close-watchers){.self-link}

#### [6.10.1]{.secno} Close requests[](#close-requests){.self-link}

In an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#close-requests:implementation-defined
x-internal="implementation-defined"} (and likely device-specific)
manner, a user can send a [close request]{#close-request .dfn export=""}
to the user agent. This indicates that the user wishes to close
something that is currently being shown on the screen, such as a
popover, menu, dialog, picker, or display mode.

::: example
Some example close requests are:

- The [Esc]{.kbd} key on desktop platforms.

- The back button or gesture on certain mobile platforms such as
  Android.

- Any assistive technology\'s dismiss gesture, such as iOS VoiceOver\'s
  two-finger scrub \"z\" gesture.

- A game controller\'s canonical \"back\" button, such as the circle
  button on a DualShock gamepad.
:::

Whenever the user agent receives a potential close request targeted at a
[`Document`{#close-requests:document}](dom.html#document)
`document`{.variable}, it must [queue a global
task](webappapis.html#queue-a-global-task){#close-requests:queue-a-global-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#close-requests:user-interaction-task-source}
given `document`{.variable}\'s [relevant global
object](webappapis.html#concept-relevant-global){#close-requests:concept-relevant-global}
to perform the following [close request steps]{#close-request-steps .dfn
export=""}:

1.  If `document`{.variable}\'s [fullscreen
    element](https://fullscreen.spec.whatwg.org/#fullscreen-element){#close-requests:fullscreen-element
    x-internal="fullscreen-element"} is not null, then:

    1.  [Fully exit
        fullscreen](https://fullscreen.spec.whatwg.org/#fully-exit-fullscreen){#close-requests:fully-exit-fullscreen
        x-internal="fully-exit-fullscreen"} given
        `document`{.variable}\'s [node
        navigable](document-sequences.html#node-navigable){#close-requests:node-navigable}\'s
        [top-level
        traversable](document-sequences.html#nav-top){#close-requests:nav-top}\'s
        [active
        document](document-sequences.html#nav-document){#close-requests:nav-document}.

    2.  Return.

    This does *not* fire any relevant event, such as
    [`keydown`{#close-requests:event-keydown}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"};
    it only causes
    [`fullscreenchange`{#close-requests:event-fullscreenchange}](https://fullscreen.spec.whatwg.org/#eventdef-document-fullscreenchange){x-internal="event-fullscreenchange"}
    to be eventually fired.

2.  Optionally, skip to the step labeled [alternative
    processing](#close-request-fallback).

    For example, if the user agent detects user frustration at repeated
    close request interception by the current web page, it might take
    this path.

3.  Fire any relevant events, per UI Events or other relevant
    specifications. [\[UIEVENTS\]](references.html#refsUIEVENTS)

    An example of a relevant event in the UI Events model would be the
    [`keydown`{#close-requests:event-keydown-2}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}
    event that UI Events
    [suggests](https://w3c.github.io/uievents/#events-keyboard-event-order)
    firing when the user presses the [Esc]{.kbd} key on their keyboard.
    On most platforms with keyboards, this is treated as a [close
    request](#close-request){#close-requests:close-request}, and so
    would trigger these [close request
    steps](#close-request-steps){#close-requests:close-request-steps}.

    An example of relevant events that are outside of the model given in
    UI Events would be assistive technology synthesizing an [Esc]{.kbd}
    [`keydown`{#close-requests:event-keydown-3}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}
    event when the user sends a [close
    request](#close-request){#close-requests:close-request-2} by using a
    dismiss gesture.

4.  Let `event`{.variable} be null if no such events are fired, or the
    [`Event`{#close-requests:event}](https://dom.spec.whatwg.org/#interface-event){x-internal="event"}
    object representing one of the fired events otherwise. If multiple
    events are fired, which one is chosen is
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#close-requests:implementation-defined-2
    x-internal="implementation-defined"}.

5.  If `event`{.variable} is not null, and its [canceled
    flag](https://dom.spec.whatwg.org/#canceled-flag){#close-requests:canceled-flag
    x-internal="canceled-flag"} is set, then return.

6.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#close-requests:fully-active},
    then return.

    This step is necessary because, if `event`{.variable} is not null,
    then an event listener might have caused `document`{.variable} to no
    longer be [fully
    active](document-sequences.html#fully-active){#close-requests:fully-active-2}.

7.  Let `closedSomething`{.variable} be the result of [processing close
    watchers](#process-close-watchers){#close-requests:process-close-watchers}
    on `document`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#close-requests:concept-relevant-global-2}.

8.  If `closedSomething`{.variable} is true, then return.

9.  ::: {#close-request-fallback}
    *Alternative processing*: Otherwise, there was nothing watching for
    a [close request](#close-request){#close-requests:close-request-3}.
    The user agent may instead interpret this interaction as some other
    action, instead of interpreting it as a close request.
    :::

On platforms where pressing the [Esc]{.kbd} key is interpreted as a
[close request](#close-request){#close-requests:close-request-4}, the
user agent must interpret the key being pressed *down* as the close
request, instead of the key being released. Thus, in the above
algorithm, the \"relevant events\" that are fired must be the single
[`keydown`{#close-requests:event-keydown-4}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}
event.

On platforms where [Esc]{.kbd} is the [close
request](#close-request){#close-requests:close-request-5}, the user
agent will first fire an appropriately-initialized
[`keydown`{#close-requests:event-keydown-5}](https://w3c.github.io/uievents/#event-type-keydown){x-internal="event-keydown"}
event. If the web developer cancels the event by calling
[`preventDefault()`{#close-requests:dom-event-preventdefault}](https://dom.spec.whatwg.org/#dom-event-preventdefault){x-internal="dom-event-preventdefault"},
then nothing further happens. But if the event fires without being
canceled, then the user agent proceeds to [process close
watchers](#process-close-watchers){#close-requests:process-close-watchers-2}.

On platforms where a back button is a potential [close
request](#close-request){#close-requests:close-request-6}, no event is
involved, so when the back button is pressed, the user agent proceeds
directly to [process close
watchers](#process-close-watchers){#close-requests:process-close-watchers-3}.
If there is an
[active](#close-watcher-active){#close-requests:close-watcher-active}
[close watcher](#close-watcher){#close-requests:close-watcher}, then
that will get triggered. If there is not, then the user agent can
interpret the back button press in another way, for example as a request
to [traverse the history by a
delta](browsing-the-web.html#traverse-the-history-by-a-delta){#close-requests:traverse-the-history-by-a-delta}
of −1.

#### [6.10.2]{.secno} Close watcher infrastructure[](#close-watcher-infrastructure){.self-link}

Each
[`Window`{#close-watcher-infrastructure:window}](nav-history-apis.html#window)
has a [close watcher manager]{#close-watcher-manager .dfn}, which is a
[struct](https://infra.spec.whatwg.org/#struct){#close-watcher-infrastructure:struct
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#close-watcher-infrastructure:struct-item
x-internal="struct-item"}:

- [Groups]{#close-watcher-manager-groups .dfn}, a
  [list](https://infra.spec.whatwg.org/#list){#close-watcher-infrastructure:list
  x-internal="list"} of
  [lists](https://infra.spec.whatwg.org/#list){#close-watcher-infrastructure:list-2
  x-internal="list"} of [close
  watchers](#close-watcher){#close-watcher-infrastructure:close-watcher},
  initially empty.

- [Allowed number of
  groups]{#close-watcher-manager-allowed-number-of-groups .dfn}, a
  number, initially 1.

- [Next user interaction allows a new group]{#close-watcher-manager-next
  .dfn}, a boolean, initially true.

Most of the complexity of the [close watcher
manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager}
comes from anti-abuse protections designed to prevent developers from
disabling users\' history traversal abilities, for platforms where a
[close
request](#close-request){#close-watcher-infrastructure:close-request}\'s
[fallback action](#close-request-fallback) is the main mechanism of
history traversal. In particular:

The grouping of [close
watchers](#close-watcher){#close-watcher-infrastructure:close-watcher-2}
is designed so that if multiple close watchers are created without
[history-action
activation](#history-action-activation){#close-watcher-infrastructure:history-action-activation},
they are grouped together, so that a user-triggered [close
request](#close-request){#close-watcher-infrastructure:close-request-2}
will close all of the close watchers in a group. This ensures that web
developers can\'t intercept an unlimited number of close requests by
creating close watchers; instead they can create a number equal to at
most 1 + the number of times the [user activates the
page](#tracking-user-activation){#close-watcher-infrastructure:tracking-user-activation}.

The [next user interaction allows a new
group](#close-watcher-manager-next){#close-watcher-infrastructure:close-watcher-manager-next}
boolean encourages web developers to create [close
watchers](#close-watcher){#close-watcher-infrastructure:close-watcher-3}
in a way that is tied to individual [user
activations](#tracking-user-activation){#close-watcher-infrastructure:tracking-user-activation-2}.
Without it, each user activation would increase the [allowed number of
groups](#close-watcher-manager-allowed-number-of-groups){#close-watcher-infrastructure:close-watcher-manager-allowed-number-of-groups},
even if the web developer isn\'t \"using\" those user activations to
create close watchers. In short:

- Allowed: user interaction; create a close watcher in its own group;
  user interaction; create a close watcher in a second independent
  group.

- Disallowed: user interaction; user interaction; create a close watcher
  in its own group; create a close watcher in a second independent
  group.

- Allowed: user interaction; user interaction; create a close watcher in
  its own group; create a close watcher grouped with the previous one.

This protection is *not* important for upholding our desired invariant
of creating at most (1 + the number of times the [user activates the
page](#tracking-user-activation){#close-watcher-infrastructure:tracking-user-activation-3})
groups. A determined abuser will just create one close watcher per user
interaction, \"banking\" them for future abuse. But this system causes
more predictable behavior for the normal case, and encourages
non-abusive developers to create close watchers directly in response to
user interactions.

To [notify the close watcher manager about user
activation]{#notify-the-close-watcher-manager-about-user-activation
.dfn} given a
[`Window`{#close-watcher-infrastructure:window-2}](nav-history-apis.html#window)
`window`{.variable}:

1.  Let `manager`{.variable} be `window`{.variable}\'s [close watcher
    manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-2}.

2.  If `manager`{.variable}\'s [next user interaction allows a new
    group](#close-watcher-manager-next){#close-watcher-infrastructure:close-watcher-manager-next-2}
    is true, then increment `manager`{.variable}\'s [allowed number of
    groups](#close-watcher-manager-allowed-number-of-groups){#close-watcher-infrastructure:close-watcher-manager-allowed-number-of-groups-2}.

3.  Set `manager`{.variable}\'s [next user interaction allows a new
    group](#close-watcher-manager-next){#close-watcher-infrastructure:close-watcher-manager-next-3}
    to false.

------------------------------------------------------------------------

A [close watcher]{#close-watcher .dfn export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#close-watcher-infrastructure:struct-2
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#close-watcher-infrastructure:struct-item-2
x-internal="struct-item"}:

- A [window]{#close-watcher-window .dfn}, a
  [`Window`{#close-watcher-infrastructure:window-3}](nav-history-apis.html#window).

- A [cancel action]{#close-watcher-cancel-action .dfn}, an algorithm
  accepting a boolean argument and returning a boolean. The argument
  indicates whether or not the cancel action algorithm can prevent the
  close request from proceeding via the algorithm\'s return value. If
  the boolean argument is true, then the algorithm can return either
  true to indicate that the caller will proceed to the [close
  action](#close-watcher-close-action){#close-watcher-infrastructure:close-watcher-close-action},
  or false to indicate that the caller will bail out. If the argument is
  false, then the return value is always false. This algorithm can never
  throw an exception.

- A [close action]{#close-watcher-close-action .dfn}, an algorithm
  accepting no arguments and returning nothing. This algorithm can never
  throw an exception.

- An [is running cancel action]{#close-watcher-is-running-cancel .dfn}
  boolean.

- A [get enabled state]{#close-watcher-get-enabled-state .dfn}, an
  algorithm accepting no arguments and returning a boolean. This
  algorithm can never throw an exception.

A [close
watcher](#close-watcher){#close-watcher-infrastructure:close-watcher-4}
`closeWatcher`{.variable} [is active]{#close-watcher-active .dfn} if
`closeWatcher`{.variable}\'s
[window](#close-watcher-window){#close-watcher-infrastructure:close-watcher-window}\'s
[close watcher
manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-3}
[contains](https://infra.spec.whatwg.org/#list-contain){#close-watcher-infrastructure:list-contains
x-internal="list-contains"} any list which
[contains](https://infra.spec.whatwg.org/#list-contain){#close-watcher-infrastructure:list-contains-2
x-internal="list-contains"} `closeWatcher`{.variable}.

------------------------------------------------------------------------

To [establish a close watcher]{#establish-a-close-watcher .dfn} given a
[`Window`{#close-watcher-infrastructure:window-4}](nav-history-apis.html#window)
`window`{.variable}, a list of steps
[`cancelAction`{.variable}]{#create-close-watcher-cancelaction .dfn}, a
list of steps
[`closeAction`{.variable}]{#create-close-watcher-closeaction .dfn}, and
an algorithm that returns a boolean
[`getEnabledState`{.variable}]{#create-close-watcher-getenabledstate
.dfn}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#close-watcher-infrastructure:assert
    x-internal="assert"}: `window`{.variable}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#close-watcher-infrastructure:concept-document-window}
    is [fully
    active](document-sequences.html#fully-active){#close-watcher-infrastructure:fully-active}.

2.  Let `closeWatcher`{.variable} be a new [close
    watcher](#close-watcher){#close-watcher-infrastructure:close-watcher-5},
    with

    [window](#close-watcher-window){#close-watcher-infrastructure:close-watcher-window-2}
    :   `window`{.variable}

    [cancel action](#close-watcher-cancel-action){#close-watcher-infrastructure:close-watcher-cancel-action}
    :   `cancelAction`{.variable}

    [close action](#close-watcher-close-action){#close-watcher-infrastructure:close-watcher-close-action-2}
    :   `closeAction`{.variable}

    [is running cancel action](#close-watcher-is-running-cancel){#close-watcher-infrastructure:close-watcher-is-running-cancel}
    :   false

    [get enabled state](#close-watcher-get-enabled-state){#close-watcher-infrastructure:close-watcher-get-enabled-state}
    :   `getEnabledState`{.variable}

3.  Let `manager`{.variable} be `window`{.variable}\'s [close watcher
    manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-4}.

4.  If `manager`{.variable}\'s
    [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#close-watcher-infrastructure:list-size
    x-internal="list-size"} is less than `manager`{.variable}\'s
    [allowed number of
    groups](#close-watcher-manager-allowed-number-of-groups){#close-watcher-infrastructure:close-watcher-manager-allowed-number-of-groups-3},
    then
    [append](https://infra.spec.whatwg.org/#list-append){#close-watcher-infrastructure:list-append
    x-internal="list-append"} « `closeWatcher`{.variable} » to
    `manager`{.variable}\'s
    [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-2}.

5.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#close-watcher-infrastructure:assert-2
        x-internal="assert"}: `manager`{.variable}\'s
        [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-3}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#close-watcher-infrastructure:list-size-2
        x-internal="list-size"} is at least 1 in this branch, since
        `manager`{.variable}\'s [allowed number of
        groups](#close-watcher-manager-allowed-number-of-groups){#close-watcher-infrastructure:close-watcher-manager-allowed-number-of-groups-4}
        is always at least 1.

    2.  [Append](https://infra.spec.whatwg.org/#list-append){#close-watcher-infrastructure:list-append-2
        x-internal="list-append"} `closeWatcher`{.variable} to
        `manager`{.variable}\'s
        [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-4}\'s
        last
        [item](https://infra.spec.whatwg.org/#list-item){#close-watcher-infrastructure:list-item
        x-internal="list-item"}.

6.  Set `manager`{.variable}\'s [next user interaction allows a new
    group](#close-watcher-manager-next){#close-watcher-infrastructure:close-watcher-manager-next-4}
    to true.

7.  Return `closeWatcher`{.variable}.

To [request to close]{#close-watcher-request-close .dfn} a [close
watcher](#close-watcher){#close-watcher-infrastructure:close-watcher-6}
`closeWatcher`{.variable} with boolean
`requireHistoryActionActivation`{.variable}:

1.  If `closeWatcher`{.variable} [is not
    active](#close-watcher-active){#close-watcher-infrastructure:close-watcher-active},
    then return true.

2.  If the result of running `closeWatcher`{.variable}\'s [get enabled
    state](#close-watcher-get-enabled-state){#close-watcher-infrastructure:close-watcher-get-enabled-state-2}
    is false, then return true.

3.  If `closeWatcher`{.variable}\'s [is running cancel
    action](#close-watcher-is-running-cancel){#close-watcher-infrastructure:close-watcher-is-running-cancel-2}
    is true, then return true.

4.  Let `window`{.variable} be `closeWatcher`{.variable}\'s
    [window](#close-watcher-window){#close-watcher-infrastructure:close-watcher-window-3}.

5.  If `window`{.variable}\'s [associated
    `Document`](nav-history-apis.html#concept-document-window){#close-watcher-infrastructure:concept-document-window-2}
    is not [fully
    active](document-sequences.html#fully-active){#close-watcher-infrastructure:fully-active-2},
    then return true.

6.  Let `canPreventClose`{.variable} be true if
    `requireHistoryActionActivation`{.variable} is false, or if
    `window`{.variable}\'s [close watcher
    manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-5}\'s
    [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-5}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#close-watcher-infrastructure:list-size-3
    x-internal="list-size"} is less than `window`{.variable}\'s [close
    watcher
    manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-6}\'s
    [allowed number of
    groups](#close-watcher-manager-allowed-number-of-groups){#close-watcher-infrastructure:close-watcher-manager-allowed-number-of-groups-5},
    and `window`{.variable} has [history-action
    activation](#history-action-activation){#close-watcher-infrastructure:history-action-activation-2};
    otherwise false.

7.  Set `closeWatcher`{.variable}\'s [is running cancel
    action](#close-watcher-is-running-cancel){#close-watcher-infrastructure:close-watcher-is-running-cancel-3}
    to true.

8.  Let `shouldContinue`{.variable} be the result of running
    `closeWatcher`{.variable}\'s [cancel
    action](#close-watcher-cancel-action){#close-watcher-infrastructure:close-watcher-cancel-action-2}
    given `canPreventClose`{.variable}.

9.  Set `closeWatcher`{.variable}\'s [is running cancel
    action](#close-watcher-is-running-cancel){#close-watcher-infrastructure:close-watcher-is-running-cancel-4}
    to false.

10. If `shouldContinue`{.variable} is false, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#close-watcher-infrastructure:assert-3
        x-internal="assert"}: `canPreventClose`{.variable} is true.

    2.  [Consume history-action user
        activation](#consume-history-action-user-activation){#close-watcher-infrastructure:consume-history-action-user-activation}
        given `window`{.variable}.

    3.  Return false.

    Note that since these substeps [consume history-action user
    activation](#consume-history-action-user-activation){#close-watcher-infrastructure:consume-history-action-user-activation-2},
    [requesting to
    close](#close-watcher-request-close){#close-watcher-infrastructure:close-watcher-request-close}
    a [close
    watcher](#close-watcher){#close-watcher-infrastructure:close-watcher-7}
    twice without any intervening [user
    activation](#tracking-user-activation){#close-watcher-infrastructure:tracking-user-activation-4}
    will result in `canPreventClose`{.variable} being false the second
    time.

11. [Close](#close-watcher-close){#close-watcher-infrastructure:close-watcher-close}
    `closeWatcher`{.variable}.

12. Return true.

To [close]{#close-watcher-close .dfn} a [close
watcher](#close-watcher){#close-watcher-infrastructure:close-watcher-8}
`closeWatcher`{.variable}:

1.  If `closeWatcher`{.variable} [is not
    active](#close-watcher-active){#close-watcher-infrastructure:close-watcher-active-2},
    then return.

2.  If the result of running `closeWatcher`{.variable}\'s [get enabled
    state](#close-watcher-get-enabled-state){#close-watcher-infrastructure:close-watcher-get-enabled-state-3}
    is false, then return.

3.  If `closeWatcher`{.variable}\'s
    [window](#close-watcher-window){#close-watcher-infrastructure:close-watcher-window-4}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#close-watcher-infrastructure:concept-document-window-3}
    is not [fully
    active](document-sequences.html#fully-active){#close-watcher-infrastructure:fully-active-3},
    then return.

4.  [Destroy](#close-watcher-destroy){#close-watcher-infrastructure:close-watcher-destroy}
    `closeWatcher`{.variable}.

5.  Run `closeWatcher`{.variable}\'s [close
    action](#close-watcher-close-action){#close-watcher-infrastructure:close-watcher-close-action-3}.

To [destroy]{#close-watcher-destroy .dfn} a [close
watcher](#close-watcher){#close-watcher-infrastructure:close-watcher-9}
`closeWatcher`{.variable}:

1.  Let `manager`{.variable} be `closeWatcher`{.variable}\'s
    [window](#close-watcher-window){#close-watcher-infrastructure:close-watcher-window-5}\'s
    [close watcher
    manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-7}.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#close-watcher-infrastructure:list-iterate
    x-internal="list-iterate"} `group`{.variable} of
    `manager`{.variable}\'s
    [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-6}:
    [remove](https://infra.spec.whatwg.org/#list-remove){#close-watcher-infrastructure:list-remove
    x-internal="list-remove"} `closeWatcher`{.variable} from
    `group`{.variable}.

3.  [Remove](https://infra.spec.whatwg.org/#list-remove){#close-watcher-infrastructure:list-remove-2
    x-internal="list-remove"} any item from `manager`{.variable}\'s
    [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-7}
    that [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#close-watcher-infrastructure:list-is-empty
    x-internal="list-is-empty"}.

------------------------------------------------------------------------

To [process close watchers]{#process-close-watchers .dfn} given a
[`Window`{#close-watcher-infrastructure:window-5}](nav-history-apis.html#window)
`window`{.variable}:

1.  Let `processedACloseWatcher`{.variable} be false.

2.  If `window`{.variable}\'s [close watcher
    manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-8}\'s
    [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-8}
    is not empty:

    1.  Let `group`{.variable} be the last
        [item](https://infra.spec.whatwg.org/#list-item){#close-watcher-infrastructure:list-item-2
        x-internal="list-item"} in `window`{.variable}\'s [close watcher
        manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-9}\'s
        [groups](#close-watcher-manager-groups){#close-watcher-infrastructure:close-watcher-manager-groups-9}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#close-watcher-infrastructure:list-iterate-2
        x-internal="list-iterate"} `closeWatcher`{.variable} of
        `group`{.variable}, in reverse order:

        1.  If the result of running `closeWatcher`{.variable}\'s [get
            enabled
            state](#close-watcher-get-enabled-state){#close-watcher-infrastructure:close-watcher-get-enabled-state-4}
            is true, set `processedACloseWatcher`{.variable} to true.

        2.  Let `shouldProceed`{.variable} be the result of [requesting
            to
            close](#close-watcher-request-close){#close-watcher-infrastructure:close-watcher-request-close-2}
            `closeWatcher`{.variable} with true.

        3.  If `shouldProceed`{.variable} is false, then
            [break](https://infra.spec.whatwg.org/#iteration-break){#close-watcher-infrastructure:break
            x-internal="break"}.

3.  If `window`{.variable}\'s [close watcher
    manager](#close-watcher-manager){#close-watcher-infrastructure:close-watcher-manager-10}\'s
    [allowed number of
    groups](#close-watcher-manager-allowed-number-of-groups){#close-watcher-infrastructure:close-watcher-manager-allowed-number-of-groups-6}
    is greater than 1, decrement it by 1.

4.  Return `processedACloseWatcher`{.variable}.

#### [6.10.3]{.secno} The [`CloseWatcher`{#the-closewatcher-interface:closewatcher}](#closewatcher) interface[](#the-closewatcher-interface){.self-link}

``` idl
[Exposed=Window]
interface CloseWatcher : EventTarget {
  constructor(optional CloseWatcherOptions options = {});

  undefined requestClose();
  undefined close();
  undefined destroy();

  attribute EventHandler oncancel;
  attribute EventHandler onclose;
};

dictionary CloseWatcherOptions {
  AbortSignal signal;
};
```

`watcher`{.variable}` = new `[`CloseWatcher`](#dom-closewatcher){#dom-closewatcher-dev}`()`\
`watcher`{.variable}` = new `[`CloseWatcher`](#dom-closewatcher){#the-closewatcher-interface:dom-closewatcher-2}`({ `[`signal`](#dom-closewatcheroptions-signal){#dom-closewatcheroptions-signal-dev}` })`

:   Creates a new
    [`CloseWatcher`{#the-closewatcher-interface:closewatcher-2}](#closewatcher)
    instance.

    If the
    [`signal`{#the-closewatcher-interface:dom-closewatcheroptions-signal}](#dom-closewatcheroptions-signal)
    option is provided, then `watcher`{.variable} can be destroyed (as
    if by
    [`watcher.destroy()`{#the-closewatcher-interface:dom-closewatcher-destroy-2}](#dom-closewatcher-destroy))
    by aborting the given
    [`AbortSignal`{#the-closewatcher-interface:abortsignal-2}](https://dom.spec.whatwg.org/#abortsignal){x-internal="abortsignal"}.

    If any [close
    watcher](#close-watcher){#the-closewatcher-interface:close-watcher}
    is already active, and the
    [`Window`{#the-closewatcher-interface:window}](nav-history-apis.html#window)
    does not have [history-action
    activation](#history-action-activation){#the-closewatcher-interface:history-action-activation},
    then the resulting
    [`CloseWatcher`{#the-closewatcher-interface:closewatcher-3}](#closewatcher)
    will be closed together with that already-active [close
    watcher](#close-watcher){#the-closewatcher-interface:close-watcher-2}
    in response to any [close
    request](#close-request){#the-closewatcher-interface:close-request}.
    (This already-active [close
    watcher](#close-watcher){#the-closewatcher-interface:close-watcher-3}
    does not necessarily have to be a
    [`CloseWatcher`{#the-closewatcher-interface:closewatcher-4}](#closewatcher)
    object; it could be a modal
    [`dialog`{#the-closewatcher-interface:the-dialog-element}](interactive-elements.html#the-dialog-element)
    element, or a popover generated by an element with the
    [`popover`{#the-closewatcher-interface:attr-popover}](popover.html#attr-popover)
    attribute.)

`watcher`{.variable}`.`[`requestClose`](#dom-closewatcher-requestclose){#dom-closewatcher-requestclose-dev}`()`

:   Acts as if a [close
    request](#close-request){#the-closewatcher-interface:close-request-2}
    was sent targeting `watcher`{.variable}, by first firing a
    [`cancel`{#the-closewatcher-interface:event-cancel}](indices.html#event-cancel)
    event, and if that event is not canceled with
    [`preventDefault()`{#the-closewatcher-interface:dom-event-preventdefault}](https://dom.spec.whatwg.org/#dom-event-preventdefault){x-internal="dom-event-preventdefault"},
    proceeding to fire a
    [`close`{#the-closewatcher-interface:event-close}](indices.html#event-close)
    event before deactivating the close watcher as if
    [`watcher.destroy()`{#the-closewatcher-interface:dom-closewatcher-destroy-3}](#dom-closewatcher-destroy)
    was called.

    This is a helper utility that can be used to consolidate cancelation
    and closing logic into the
    [`cancel`{#the-closewatcher-interface:event-cancel-2}](indices.html#event-cancel)
    and
    [`close`{#the-closewatcher-interface:event-close-2}](indices.html#event-close)
    event handlers, by having all non-[close
    request](#close-request){#the-closewatcher-interface:close-request-3}
    closing affordances call this method.

`watcher`{.variable}`.`[`close`](#dom-closewatcher-close){#dom-closewatcher-close-dev}`()`

:   Immediately fires the
    [`close`{#the-closewatcher-interface:event-close-3}](indices.html#event-close)
    event, and then deactivates the close watcher as if
    [`watcher.destroy()`{#the-closewatcher-interface:dom-closewatcher-destroy-4}](#dom-closewatcher-destroy)
    was called.

    This is a helper utility that can be used trigger the closing logic
    into the
    [`close`{#the-closewatcher-interface:event-close-4}](indices.html#event-close)
    event handler, skipping any logic in the
    [`cancel`{#the-closewatcher-interface:event-cancel-3}](indices.html#event-cancel)
    event handler.

`watcher`{.variable}`.`[`destroy`](#dom-closewatcher-destroy){#dom-closewatcher-destroy-dev}`()`

:   Deactivates `watcher`{.variable}, so that it will no longer receive
    [`close`{#the-closewatcher-interface:event-close-5}](indices.html#event-close)
    events and so that new independent
    [`CloseWatcher`{#the-closewatcher-interface:closewatcher-5}](#closewatcher)
    instances can be constructed.

    This is intended to be called if the relevant UI element is torn
    down in some other way than being closed.

Each
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-6}](#closewatcher)
instance has an [internal close watcher]{#internal-close-watcher .dfn},
which is a [close
watcher](#close-watcher){#the-closewatcher-interface:close-watcher-4}.

The [`new CloseWatcher(``options`{.variable}`)`]{#dom-closewatcher .dfn}
constructor steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-closewatcher-interface:concept-relevant-global}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-closewatcher-interface:concept-document-window}
    is not [fully
    active](document-sequences.html#fully-active){#the-closewatcher-interface:fully-active},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-closewatcher-interface:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-closewatcher-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `closeWatcher`{.variable} be the result of [establishing a close
    watcher](#establish-a-close-watcher){#the-closewatcher-interface:establish-a-close-watcher}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-closewatcher-interface:concept-relevant-global-2},
    with:

    - *[cancelAction](#create-close-watcher-cancelaction)* given
      `canPreventClose`{.variable} being to return the result of [firing
      an
      event](https://dom.spec.whatwg.org/#concept-event-fire){#the-closewatcher-interface:concept-event-fire
      x-internal="concept-event-fire"} named
      [`cancel`{#the-closewatcher-interface:event-cancel-4}](indices.html#event-cancel)
      at
      [this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this-3
      x-internal="this"}, with the
      [`cancelable`{#the-closewatcher-interface:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
      attribute initialized to `canPreventClose`{.variable}.

    - *[closeAction](#create-close-watcher-closeaction)* being to [fire
      an
      event](https://dom.spec.whatwg.org/#concept-event-fire){#the-closewatcher-interface:concept-event-fire-2
      x-internal="concept-event-fire"} named
      [`close`{#the-closewatcher-interface:event-close-6}](indices.html#event-close)
      at
      [this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this-4
      x-internal="this"}.

    - *[getEnabledState](#close-watcher-get-enabled-state)* being to
      return true.

3.  If
    `options`{.variable}\[\"[`signal`{#the-closewatcher-interface:dom-closewatcheroptions-signal-2}](#dom-closewatcheroptions-signal)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-closewatcher-interface:map-exists
    x-internal="map-exists"}, then:

    1.  If
        `options`{.variable}\[\"[`signal`{#the-closewatcher-interface:dom-closewatcheroptions-signal-3}](#dom-closewatcheroptions-signal)\"\]
        is
        [aborted](https://dom.spec.whatwg.org/#abortsignal-aborted){#the-closewatcher-interface:abortsignal-aborted
        x-internal="abortsignal-aborted"}, then
        [destroy](#close-watcher-destroy){#the-closewatcher-interface:close-watcher-destroy}
        `closeWatcher`{.variable}.

    2.  [Add](https://dom.spec.whatwg.org/#abortsignal-add){#the-closewatcher-interface:abortsignal-add
        x-internal="abortsignal-add"} the following steps to
        `options`{.variable}\[\"[`signal`{#the-closewatcher-interface:dom-closewatcheroptions-signal-4}](#dom-closewatcheroptions-signal)\"\]:

        1.  [Destroy](#close-watcher-destroy){#the-closewatcher-interface:close-watcher-destroy-2}
            `closeWatcher`{.variable}.

4.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this-5
    x-internal="this"}\'s [internal close
    watcher](#internal-close-watcher){#the-closewatcher-interface:internal-close-watcher}
    to `closeWatcher`{.variable}.

The [`requestClose()`]{#dom-closewatcher-requestclose .dfn
dfn-for="CloseWatcher" dfn-type="method"} method steps are to [request
to
close](#close-watcher-request-close){#the-closewatcher-interface:close-watcher-request-close}
[this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this-6
x-internal="this"}\'s [internal close
watcher](#internal-close-watcher){#the-closewatcher-interface:internal-close-watcher-2}
with false.

The [`close()`]{#dom-closewatcher-close .dfn dfn-for="CloseWatcher"
dfn-type="method"} method steps are to
[close](#close-watcher-close){#the-closewatcher-interface:close-watcher-close}
[this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this-7
x-internal="this"}\'s [internal close
watcher](#internal-close-watcher){#the-closewatcher-interface:internal-close-watcher-3}.

The [`destroy()`]{#dom-closewatcher-destroy .dfn dfn-for="CloseWatcher"
dfn-type="method"} method steps are to
[destroy](#close-watcher-destroy){#the-closewatcher-interface:close-watcher-destroy-3}
[this](https://webidl.spec.whatwg.org/#this){#the-closewatcher-interface:this-8
x-internal="this"}\'s [internal close
watcher](#internal-close-watcher){#the-closewatcher-interface:internal-close-watcher-4}.

The following are the [event
handlers](webappapis.html#event-handlers){#the-closewatcher-interface:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#the-closewatcher-interface:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-closewatcher-interface:event-handler-idl-attributes},
by all objects implementing the
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-7}](#closewatcher)
interface:

[Event
handler](webappapis.html#event-handlers){#the-closewatcher-interface:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#the-closewatcher-interface:event-handler-event-type-2}

[`oncancel`]{#handler-closewatcher-oncancel .dfn dfn-for="CloseWatcher"
dfn-type="attribute"}

[`cancel`{#the-closewatcher-interface:event-cancel-5}](indices.html#event-cancel)

[`onclose`]{#handler-closewatcher-onclose .dfn dfn-for="CloseWatcher"
dfn-type="attribute"}

[`close`{#the-closewatcher-interface:event-close-7}](indices.html#event-close)

::: {#example-CloseWatcher-basic .example}
If one wanted to implement a custom picker control, which closed itself
on a user-provided [close
request](#close-request){#the-closewatcher-interface:close-request-4} as
well as when a close button is pressed, the following code shows how one
would use the
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-8}](#closewatcher)
API to process close requests:

``` js
const watcher = new CloseWatcher();
const picker = setUpAndShowPickerDOMElement();

let chosenValue = null;

watcher.onclose = () => {
  chosenValue = picker.querySelector('input').value;
  picker.remove();
};

picker.querySelector('.close-button').onclick = () => watcher.requestClose();
```

Note how the logic to gather the chosen value is centralized in the
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-9}](#closewatcher)
object\'s
[`close`{#the-closewatcher-interface:event-close-8}](indices.html#event-close)
event handler, with the
[`click`{#the-closewatcher-interface:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event handler for the close button delegating to that logic by calling
[`requestClose()`{#the-closewatcher-interface:dom-closewatcher-requestclose-2}](#dom-closewatcher-requestclose).
:::

::: {#example-CloseWatcher-cancel .example}
The
[`cancel`{#the-closewatcher-interface:event-cancel-6}](indices.html#event-cancel)
event on
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-10}](#closewatcher)
objects can be used to prevent the
[`close`{#the-closewatcher-interface:event-close-9}](indices.html#event-close)
event from firing, and the
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-11}](#closewatcher)
from being destroying. A typical use case is as follows:

``` js
watcher.oncancel = async (e) => {
  if (hasUnsavedData && e.cancelable) {
    e.preventDefault();

    const userReallyWantsToClose = await askForConfirmation("Are you sure you want to close?");
    if (userReallyWantsToClose) {
      hasUnsavedData = false;
      watcher.close();
    }
  }
};
```

For abuse prevention purposes, this event is only
[`cancelable`{#the-closewatcher-interface:dom-event-cancelable-2}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
if the page has [history-action
activation](#history-action-activation){#the-closewatcher-interface:history-action-activation-2},
which will be lost after any given [close
request](#close-request){#the-closewatcher-interface:close-request-5}.
This ensures that if the user sends a close request twice in a row
without any intervening user activation, the request definitely
succeeds; the second request ignores any
[`cancel`{#the-closewatcher-interface:event-cancel-7}](indices.html#event-cancel)
event handler\'s attempt to call
[`preventDefault()`{#the-closewatcher-interface:dom-event-preventdefault-2}](https://dom.spec.whatwg.org/#dom-event-preventdefault){x-internal="dom-event-preventdefault"}
and proceeds to close the
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-12}](#closewatcher).
:::

Combined, the above two examples show how
[`requestClose()`{#the-closewatcher-interface:dom-closewatcher-requestclose-3}](#dom-closewatcher-requestclose)
and
[`close()`{#the-closewatcher-interface:dom-closewatcher-close-2}](#dom-closewatcher-close)
differ. Because we used
[`requestClose()`{#the-closewatcher-interface:dom-closewatcher-requestclose-4}](#dom-closewatcher-requestclose)
in the
[`click`{#the-closewatcher-interface:event-click-2}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event handler for the close button, clicking that button will trigger
the
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-13}](#closewatcher)\'s
[`cancel`{#the-closewatcher-interface:event-cancel-8}](indices.html#event-cancel)
event, and thus potentially ask the user for confirmation if there is
unsaved data. If we had used
[`close()`{#the-closewatcher-interface:dom-closewatcher-close-3}](#dom-closewatcher-close),
then this check would be skipped. Sometimes that is appropriate, but
usually
[`requestClose()`{#the-closewatcher-interface:dom-closewatcher-requestclose-5}](#dom-closewatcher-requestclose)
is the better option for user-triggered close requests.

::: {#example-CloseWatcher-grouping .example}
In addition to the [user
activation](#tracking-user-activation){#the-closewatcher-interface:tracking-user-activation}
restrictions for
[`cancel`{#the-closewatcher-interface:event-cancel-9}](indices.html#event-cancel)
events, there is a more subtle form of user activation gating for
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-14}](#closewatcher)
construction. If one creates more than one
[`CloseWatcher`{#the-closewatcher-interface:closewatcher-15}](#closewatcher)
without user activation, then the newly-created one will get grouped
together with the most-recently-created [close
watcher](#close-watcher){#the-closewatcher-interface:close-watcher-5},
so that a single [close
request](#close-request){#the-closewatcher-interface:close-request-6}
will close them both:

``` js
window.onload = () => {
  // This will work as normal: it is the first close watcher created without user activation.
  (new CloseWatcher()).onclose = () => { /* ... */ };
};

button1.onclick = () => {
  // This will work as normal: the button click counts as user activation.
  (new CloseWatcher()).onclose = () => { /* ... */ };
};

button2.onclick = () => {
  // These will be grouped together, and both will close in response to a single close request.
  (new CloseWatcher()).onclose = () => { /* ... */ };
  (new CloseWatcher()).onclose = () => { /* ... */ };
};
```

This means that calling
[`destroy()`{#the-closewatcher-interface:dom-closewatcher-destroy-5}](#dom-closewatcher-destroy),
[`close()`{#the-closewatcher-interface:dom-closewatcher-close-4}](#dom-closewatcher-close),
or
[`requestClose()`{#the-closewatcher-interface:dom-closewatcher-requestclose-6}](#dom-closewatcher-requestclose)
properly is important. Doing so is the only way to get back the \"free\"
ungrouped close watcher slot. Such close watchers created without user
activation are useful for cases like session inactivity timeout dialogs
or urgent notifications of server-triggered events, which are not
generated in response to user activation.
:::

[← 5 Microdata](microdata.html) --- [Table of Contents](./) --- [6.11
Drag and drop →](dnd.html)
