HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.10.5 The input element](input.html) --- [Table of Contents](./) ---
[4.10.17 Form control infrastructure
→](form-control-infrastructure.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.10.6]{.secno} The `button`
            element](form-elements.html#the-button-element)
        2.  [[4.10.7]{.secno} The `select`
            element](form-elements.html#the-select-element)
        3.  [[4.10.8]{.secno} The `datalist`
            element](form-elements.html#the-datalist-element)
        4.  [[4.10.9]{.secno} The `optgroup`
            element](form-elements.html#the-optgroup-element)
        5.  [[4.10.10]{.secno} The `option`
            element](form-elements.html#the-option-element)
        6.  [[4.10.11]{.secno} The `textarea`
            element](form-elements.html#the-textarea-element)
        7.  [[4.10.12]{.secno} The `output`
            element](form-elements.html#the-output-element)
        8.  [[4.10.13]{.secno} The `progress`
            element](form-elements.html#the-progress-element)
        9.  [[4.10.14]{.secno} The `meter`
            element](form-elements.html#the-meter-element)
        10. [[4.10.15]{.secno} The `fieldset`
            element](form-elements.html#the-fieldset-element)
        11. [[4.10.16]{.secno} The `legend`
            element](form-elements.html#the-legend-element)
    :::

#### [4.10.6]{.secno} The [`button`]{.dfn dfn-type="element"} element[](#the-button-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/button](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button "The <button> HTML element is an interactive element activated by a user with a mouse, keyboard, finger, voice command, or other assistive technology. Once activated, it then performs an action, such as submitting a form or opening a dialog.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera15+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLButtonElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement "The HTMLButtonElement interface provides properties and methods (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating <button> elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-button-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-button-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-button-element:phrasing-content-2}.
:   [Interactive
    content](dom.html#interactive-content-2){#the-button-element:interactive-content-2}.
:   [Listed](forms.html#category-listed){#the-button-element:category-listed},
    [labelable](forms.html#category-label){#the-button-element:category-label},
    [submittable](forms.html#category-submit){#the-button-element:category-submit},
    and [autocapitalize-and-autocorrect
    inheriting](forms.html#category-autocapitalize){#the-button-element:category-autocapitalize}
    [form-associated
    element](forms.html#form-associated-element){#the-button-element:form-associated-element}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-button-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-button-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-button-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-button-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-button-element:phrasing-content-2-3},
    but there must be no [interactive
    content](dom.html#interactive-content-2){#the-button-element:interactive-content-2-2}
    descendant and no descendant with the
    [`tabindex`{#the-button-element:attr-tabindex}](interaction.html#attr-tabindex)
    attribute specified.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-button-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-button-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-button-element:global-attributes}
:   [`command`{#the-button-element:attr-button-command}](#attr-button-command)
    --- Indicates to the targeted element which action to take.
:   [`commandfor`{#the-button-element:attr-button-commandfor}](#attr-button-commandfor)
    --- Targets another element to be invoked.
:   [`disabled`{#the-button-element:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
    --- Whether the form control is disabled
:   [`form`{#the-button-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    --- Associates the element with a
    [`form`{#the-button-element:the-form-element}](forms.html#the-form-element)
    element
:   [`formaction`{#the-button-element:attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction)
    ---
    [URL](https://url.spec.whatwg.org/#concept-url){#the-button-element:url
    x-internal="url"} to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-button-element:form-submission-2}
:   [`formenctype`{#the-button-element:attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype)
    --- [Entry
    list](form-control-infrastructure.html#entry-list){#the-button-element:entry-list}
    encoding type to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-button-element:form-submission-2-2}
:   [`formmethod`{#the-button-element:attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod)
    --- Variant to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-button-element:form-submission-2-3}
:   [`formnovalidate`{#the-button-element:attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate)
    --- Bypass form control validation for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-button-element:form-submission-2-4}
:   [`formtarget`{#the-button-element:attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget)
    ---
    [Navigable](document-sequences.html#navigable){#the-button-element:navigable}
    for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-button-element:form-submission-2-5}
:   [`name`{#the-button-element:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    --- Name of the element to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-button-element:form-submission-2-6}
    and in the
    [`form.elements`{#the-button-element:dom-form-elements}](forms.html#dom-form-elements)
    API
:   [`popovertarget`{#the-button-element:attr-popovertarget}](popover.html#attr-popovertarget)
    --- Targets a popover element to toggle, show, or hide
:   [`popovertargetaction`{#the-button-element:attr-popovertargetaction}](popover.html#attr-popovertargetaction)
    --- Indicates whether a targeted popover element is to be toggled,
    shown, or hidden
:   [`type`{#the-button-element:attr-button-type}](#attr-button-type)
    --- Type of button
:   [`value`{#the-button-element:attr-button-value}](#attr-button-value)
    --- Value to be used for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-button-element:form-submission-2-7}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-button-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-button).
:   [For implementers](https://w3c.github.io/html-aam/#el-button).

[DOM interface](dom.html#concept-element-dom){#the-button-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLButtonElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString command;
      [CEReactions] attribute Element? commandForElement;
      [CEReactions] attribute boolean disabled;
      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute USVString formAction;
      [CEReactions] attribute DOMString formEnctype;
      [CEReactions] attribute DOMString formMethod;
      [CEReactions] attribute boolean formNoValidate;
      [CEReactions] attribute DOMString formTarget;
      [CEReactions] attribute DOMString name;
      [CEReactions] attribute DOMString type;
      [CEReactions] attribute DOMString value;

      readonly attribute boolean willValidate;
      readonly attribute ValidityState validity;
      readonly attribute DOMString validationMessage;
      boolean checkValidity();
      boolean reportValidity();
      undefined setCustomValidity(DOMString error);

      readonly attribute NodeList labels;
    };
    HTMLButtonElement includes PopoverInvokerElement;
    ```

The
[`button`{#the-button-element:the-button-element}](#the-button-element)
element
[represents](dom.html#represents){#the-button-element:represents} a
button labeled by its contents.

The element is a
[button](forms.html#concept-button){#the-button-element:concept-button}.

The [`type`]{#attr-button-type .dfn dfn-for="button"
dfn-type="element-attr"} attribute controls the behavior of the button
when it is activated. It is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-button-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`submit`]{#attr-button-type-submit .dfn dfn-for="button/type"
dfn-type="attr-value"}

[Submit Button]{#attr-button-type-submit-state .dfn}

Submits the form.

[`reset`]{#attr-button-type-reset .dfn dfn-for="button/type"
dfn-type="attr-value"}

[Reset Button]{#attr-button-type-reset-state .dfn}

Resets the form.

[`button`]{#attr-button-type-button .dfn dfn-for="button/type"
dfn-type="attr-value"}

[Button]{#attr-button-type-button-state .dfn}

Does nothing.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Auto]{#attr-button-type-auto-state .dfn} state.

A
[`button`{#the-button-element:the-button-element-2}](#the-button-element)
element is said to be a [submit
button](forms.html#concept-submit-button){#the-button-element:concept-submit-button}
if any of the following are true:

- the
  [`type`{#the-button-element:attr-button-type-2}](#attr-button-type)
  attribute is in the
  [Auto](#attr-button-type-auto-state){#the-button-element:attr-button-type-auto-state}
  state and both the
  [`command`{#the-button-element:attr-button-command-2}](#attr-button-command)
  and
  [`commandfor`{#the-button-element:attr-button-commandfor-2}](#attr-button-commandfor)
  content attributes are not present; or

- the
  [`type`{#the-button-element:attr-button-type-3}](#attr-button-type)
  attribute is in the [Submit
  Button](#attr-button-type-submit-state){#the-button-element:attr-button-type-submit-state}
  state.

**Constraint validation**: If the element is not a [submit
button](forms.html#concept-submit-button){#the-button-element:concept-submit-button-2},
the element is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#the-button-element:barred-from-constraint-validation}.

If specified, the [`commandfor`]{#attr-button-commandfor .dfn
dfn-for="button" dfn-type="element-attr"} attribute value must be the
[ID](https://dom.spec.whatwg.org/#concept-id){#the-button-element:concept-id
x-internal="concept-id"} of an element in the same
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-button-element:tree
x-internal="tree"} as the
[button](forms.html#concept-button){#the-button-element:concept-button-2}
with the
[`commandfor`{#the-button-element:attr-button-commandfor-3}](#attr-button-commandfor)
attribute.

The [`command`]{#attr-button-command .dfn dfn-for="button"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-button-element:enumerated-attribute-2}
with the following keywords and states:

Keyword

State

Brief description

[`toggle-popover`]{#attr-button-command-toggle-popover .dfn
dfn-for="html-global/command" dfn-type="attr-value"}

[Toggle Popover]{#attr-button-command-toggle-popover-state .dfn}

Shows or hides the targeted
[`popover`{#the-button-element:attr-popover}](popover.html#attr-popover)
element.

[`show-popover`]{#attr-button-command-show-popover .dfn
dfn-for="html-global/command" dfn-type="attr-value"}

[Show Popover]{#attr-button-command-show-popover-state .dfn}

Shows the targeted
[`popover`{#the-button-element:attr-popover-2}](popover.html#attr-popover)
element.

[`hide-popover`]{#attr-button-command-hide-popover .dfn
dfn-for="html-global/command" dfn-type="attr-value"}

[Hide Popover]{#attr-button-command-hide-popover-state .dfn}

Hides the targeted
[`popover`{#the-button-element:attr-popover-3}](popover.html#attr-popover)
element.

[`close`]{#attr-button-command-close .dfn dfn-for="html-global/command"
dfn-type="attr-value"}

[Close]{#attr-button-command-close-state .dfn}

Closes the targeted
[`dialog`{#the-button-element:the-dialog-element}](interactive-elements.html#the-dialog-element)
element.

[`request-close`]{#attr-button-command-request-close .dfn
dfn-for="html-global/command" dfn-type="attr-value"}

[Request Close]{#attr-button-command-request-close-state .dfn}

Requests to close the targeted
[`dialog`{#the-button-element:the-dialog-element-2}](interactive-elements.html#the-dialog-element)
element.

[`show-modal`]{#attr-button-command-show-modal .dfn
dfn-for="html-global/command" dfn-type="attr-value"}

[Show Modal]{#attr-button-command-show-modal-state .dfn}

Opens the targeted
[`dialog`{#the-button-element:the-dialog-element-3}](interactive-elements.html#the-dialog-element)
element as modal.

A [custom command
keyword](#attr-button-command-custom){#the-button-element:attr-button-command-custom}

[Custom]{#attr-button-command-custom-state .dfn}

Only dispatches the
[`command`{#the-button-element:event-command}](indices.html#event-command)
event on the targeted element.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Unknown]{#attr-button-command-unknown-state .dfn} state.

A [custom command keyword]{#attr-button-command-custom .dfn} is a string
that [starts
with](https://infra.spec.whatwg.org/#string-starts-with){#the-button-element:starts-with
x-internal="starts-with"} \"`--`\".

A
[`button`{#the-button-element:the-button-element-3}](#the-button-element)
element `element`{.variable}\'s [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-button-element:activation-behaviour
x-internal="activation-behaviour"} given `event`{.variable} is:

1.  If `element`{.variable} is
    [disabled](form-control-infrastructure.html#concept-fe-disabled){#the-button-element:concept-fe-disabled},
    then return.

2.  If `element`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-button-element:node-document
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#the-button-element:fully-active},
    then return.

3.  If `element`{.variable} has a [form
    owner](form-control-infrastructure.html#form-owner){#the-button-element:form-owner}:

    1.  If `element`{.variable} is a [submit
        button](forms.html#concept-submit-button){#the-button-element:concept-submit-button-3},
        then
        [submit](form-control-infrastructure.html#concept-form-submit){#the-button-element:concept-form-submit}
        `element`{.variable}\'s [form
        owner](form-control-infrastructure.html#form-owner){#the-button-element:form-owner-2}
        from `element`{.variable} with
        *[userInvolvement](form-control-infrastructure.html#submit-user-involvement)*
        set to `event`{.variable}\'s [user navigation
        involvement](browsing-the-web.html#event-uni){#the-button-element:event-uni},
        and return.

    2.  If `element`{.variable}\'s
        [`type`{#the-button-element:attr-button-type-4}](#attr-button-type)
        attribute is in the [Reset
        Button](#attr-button-type-reset-state){#the-button-element:attr-button-type-reset-state}
        state, then
        [reset](form-control-infrastructure.html#concept-form-reset){#the-button-element:concept-form-reset}
        `element`{.variable}\'s [form
        owner](form-control-infrastructure.html#form-owner){#the-button-element:form-owner-3},
        and return.

    3.  If `element`{.variable}\'s
        [`type`{#the-button-element:attr-button-type-5}](#attr-button-type)
        attribute is in the
        [Auto](#attr-button-type-auto-state){#the-button-element:attr-button-type-auto-state-2}
        state, then return.

4.  Let `target`{.variable} be the result of running
    `element`{.variable}\'s [get the `commandfor`-associated
    element](common-dom-interfaces.html#attr-associated-element){#the-button-element:attr-associated-element}.

5.  If `target`{.variable} is not null:

    1.  Let `command`{.variable} be `element`{.variable}\'s
        [`command`{#the-button-element:attr-button-command-3}](#attr-button-command)
        attribute.

    2.  If `command`{.variable} is in the
        [Unknown](#attr-button-command-unknown-state){#the-button-element:attr-button-command-unknown-state}
        state, then return.

    3.  Let `isPopover`{.variable} be true if `target`{.variable}\'s
        [`popover`{#the-button-element:attr-popover-4}](popover.html#attr-popover)
        attribute is not in the [No
        Popover](popover.html#attr-popover-none-state){#the-button-element:attr-popover-none-state}
        state; otherwise false.

    4.  If `isPopover`{.variable} is false and `command`{.variable} is
        not in the
        [Custom](#attr-button-command-custom-state){#the-button-element:attr-button-command-custom-state}
        state:

        1.  [Assert](https://infra.spec.whatwg.org/#assert){#the-button-element:assert
            x-internal="assert"}: `target`{.variable}\'s
            [namespace](https://dom.spec.whatwg.org/#concept-element-namespace){#the-button-element:concept-element-namespace
            x-internal="concept-element-namespace"} is the [HTML
            namespace](https://infra.spec.whatwg.org/#html-namespace){#the-button-element:html-namespace-2
            x-internal="html-namespace-2"}.

        2.  If this standard does not define [is valid invoker command
            steps](#is-valid-invoker-command-steps){#the-button-element:is-valid-invoker-command-steps}
            for `target`{.variable}\'s [local
            name](https://dom.spec.whatwg.org/#concept-element-local-name){#the-button-element:concept-element-local-name
            x-internal="concept-element-local-name"}, then return.

        3.  Otherwise, if the result of running `target`{.variable}\'s
            corresponding [is valid invoker command
            steps](#is-valid-invoker-command-steps){#the-button-element:is-valid-invoker-command-steps-2}
            given `command`{.variable} is false, then return.

    5.  Let `continue`{.variable} be the result of [firing an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-button-element:concept-event-fire
        x-internal="concept-event-fire"} named
        [`command`{#the-button-element:event-command-2}](indices.html#event-command)
        at `target`{.variable}, using
        [`CommandEvent`{#the-button-element:commandevent}](interaction.html#commandevent),
        with its
        [`command`{#the-button-element:dom-commandevent-command}](interaction.html#dom-commandevent-command)
        attribute initialized to `command`{.variable}, its
        [`source`{#the-button-element:dom-commandevent-source}](interaction.html#dom-commandevent-source)
        attribute initialized to `element`{.variable}, and its
        [`cancelable`{#the-button-element:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
        and
        [`composed`{#the-button-element:dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
        attributes initialized to true.

        [DOM standard issue
        #1328](https://github.com/whatwg/dom/issues/1328) tracks how to
        better standardize associated event data in a way which makes
        sense on Events. Currently an event attribute initialized to a
        value cannot also have a getter, and so an internal slot (or map
        of additional fields) is required to properly specify this.

    6.  If `continue`{.variable} is false, then return.

    7.  If `target`{.variable} is not
        [connected](https://dom.spec.whatwg.org/#connected){#the-button-element:connected
        x-internal="connected"}, then return.

    8.  If `command`{.variable} is in the
        [Custom](#attr-button-command-custom-state){#the-button-element:attr-button-command-custom-state-2}
        state, then return.

    9.  If `command`{.variable} is in the [Hide
        Popover](#attr-button-command-hide-popover-state){#the-button-element:attr-button-command-hide-popover-state}
        state:

        1.  If the result of running [check popover
            validity](popover.html#check-popover-validity){#the-button-element:check-popover-validity}
            given `target`{.variable}, true, false, and null is true,
            then run the [hide popover
            algorithm](popover.html#hide-popover-algorithm){#the-button-element:hide-popover-algorithm}
            given `target`{.variable}, true, true, false, and
            `element`{.variable}.

    10. Otherwise, if `command`{.variable} is in the [Toggle
        Popover](#attr-button-command-toggle-popover-state){#the-button-element:attr-button-command-toggle-popover-state}
        state:

        1.  If the result of running [check popover
            validity](popover.html#check-popover-validity){#the-button-element:check-popover-validity-2}
            given `target`{.variable}, false, false, and null is true,
            then run the [show popover
            algorithm](popover.html#show-popover){#the-button-element:show-popover}
            given `target`{.variable}, false, and
            [this](https://webidl.spec.whatwg.org/#this){#the-button-element:this
            x-internal="this"}.

        2.  Otherwise, if the result of running [check popover
            validity](popover.html#check-popover-validity){#the-button-element:check-popover-validity-3}
            given `target`{.variable}, true, false, and null is true,
            then run the [hide popover
            algorithm](popover.html#hide-popover-algorithm){#the-button-element:hide-popover-algorithm-2}
            given `target`{.variable}, true, true, false, and
            `element`{.variable}.

    11. Otherwise, if `command`{.variable} is in the [Show
        Popover](#attr-button-command-show-popover-state){#the-button-element:attr-button-command-show-popover-state}
        state:

        1.  If the result of running [check popover
            validity](popover.html#check-popover-validity){#the-button-element:check-popover-validity-4}
            given `target`{.variable}, false, false, and null is true,
            then run the [show popover
            algorithm](popover.html#show-popover){#the-button-element:show-popover-2}
            given `target`{.variable}, false, and
            [this](https://webidl.spec.whatwg.org/#this){#the-button-element:this-2
            x-internal="this"}.

    12. Otherwise, if this standard defines [invoker command
        steps](#invoker-command-steps){#the-button-element:invoker-command-steps}
        for `target`{.variable}\'s [local
        name](https://dom.spec.whatwg.org/#concept-element-local-name){#the-button-element:concept-element-local-name-2
        x-internal="concept-element-local-name"}, then run the
        corresponding [invoker command
        steps](#invoker-command-steps){#the-button-element:invoker-command-steps-2}
        given `target`{.variable}, `element`{.variable}, and
        `command`{.variable}.

6.  Otherwise, run the [popover target attribute activation
    behavior](popover.html#popover-target-attribute-activation-behavior){#the-button-element:popover-target-attribute-activation-behavior}
    given `element`{.variable} and `event`{.variable}\'s
    [target](https://dom.spec.whatwg.org/#concept-event-target){#the-button-element:concept-event-target
    x-internal="concept-event-target"}.

An [HTML
element](infrastructure.html#html-elements){#the-button-element:html-elements}
can have specific [is valid invoker command
steps]{#is-valid-invoker-command-steps .dfn} and [invoker command
steps]{#invoker-command-steps .dfn} defined for the element\'s [local
name](https://dom.spec.whatwg.org/#concept-element-local-name){#the-button-element:concept-element-local-name-3
x-internal="concept-element-local-name"}.

The
[`form`{#the-button-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the
[`button`{#the-button-element:the-button-element-4}](#the-button-element)
element with its [form
owner](form-control-infrastructure.html#form-owner){#the-button-element:form-owner-4}.
The
[`name`{#the-button-element:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attribute represents the element\'s name. The
[`disabled`{#the-button-element:attr-fe-disabled-2}](form-control-infrastructure.html#attr-fe-disabled)
attribute is used to make the control non-interactive and to prevent its
value from being submitted. The
[`formaction`{#the-button-element:attr-fs-formaction-2}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#the-button-element:attr-fs-formenctype-2}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#the-button-element:attr-fs-formmethod-2}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#the-button-element:attr-fs-formnovalidate-2}](form-control-infrastructure.html#attr-fs-formnovalidate),
and
[`formtarget`{#the-button-element:attr-fs-formtarget-2}](form-control-infrastructure.html#attr-fs-formtarget)
attributes are [attributes for form
submission](form-control-infrastructure.html#attributes-for-form-submission){#the-button-element:attributes-for-form-submission}.

The
[`formnovalidate`{#the-button-element:attr-fs-formnovalidate-3}](form-control-infrastructure.html#attr-fs-formnovalidate)
attribute can be used to make submit buttons that do not trigger the
constraint validation.

The
[`formaction`{#the-button-element:attr-fs-formaction-3}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#the-button-element:attr-fs-formenctype-3}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#the-button-element:attr-fs-formmethod-3}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#the-button-element:attr-fs-formnovalidate-4}](form-control-infrastructure.html#attr-fs-formnovalidate),
and
[`formtarget`{#the-button-element:attr-fs-formtarget-3}](form-control-infrastructure.html#attr-fs-formtarget)
must not be specified if the element is not a [submit
button](forms.html#concept-submit-button){#the-button-element:concept-submit-button-4}.

The [`commandForElement`]{#dom-button-commandforelement .dfn
dfn-for="HTMLButtonElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-button-element:reflect}
the element\'s
[`commandfor`{#the-button-element:attr-button-commandfor-4}](#attr-button-commandfor)
content attribute.

The [`command`]{#dom-button-command .dfn dfn-for="HTMLButtonElement"
dfn-type="attribute"} getter steps are:

1.  Let `command`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-button-element:this-3
    x-internal="this"}\'s
    [`command`{#the-button-element:attr-button-command-4}](#attr-button-command)
    attribute.

2.  If `command`{.variable} is in the
    [Custom](#attr-button-command-custom-state){#the-button-element:attr-button-command-custom-state-3}
    state, then return `command`{.variable}\'s value.

3.  If `command`{.variable} is in the
    [Unknown](#attr-button-command-unknown-state){#the-button-element:attr-button-command-unknown-state-2}
    state, then return the empty string.

4.  Return the keyword corresponding to the value of
    `command`{.variable}.

The
[`command`{#the-button-element:dom-button-command-2}](#dom-button-command)
setter steps are to set the
[`command`{#the-button-element:attr-button-command-5}](#attr-button-command)
content attribute to the given value.

The [`value`]{#attr-button-value .dfn dfn-for="button"
dfn-type="element-attr"} attribute gives the element\'s value for the
purposes of form submission. The element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-button-element:concept-fe-value}
is the value of the element\'s
[`value`{#the-button-element:attr-button-value-2}](#attr-button-value)
attribute, if there is one; otherwise the empty string. The element\'s
[optional
value](form-control-infrastructure.html#concept-fe-optional-value){#the-button-element:concept-fe-optional-value}
is the value of the element\'s
[`value`{#the-button-element:attr-button-value-3}](#attr-button-value)
attribute, if there is one; otherwise null.

A button (and its value) is only included in the form submission if the
button itself was used to initiate the form submission.

------------------------------------------------------------------------

The [`value`]{#dom-button-value .dfn dfn-for="HTMLButtonElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-button-element:reflect-2}
the content attribute of the same name.

The [`type`]{#dom-button-type .dfn dfn-for="HTMLButtonElement"
dfn-type="attribute"} getter steps are:

1.  If `this`{.variable} is a [submit
    button](forms.html#concept-submit-button){#the-button-element:concept-submit-button-5},
    then return \"`submit`\".

2.  Let `state`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-button-element:this-4
    x-internal="this"}\'s
    [`type`{#the-button-element:attr-button-type-6}](#attr-button-type)
    attribute.

3.  [Assert](https://infra.spec.whatwg.org/#assert){#the-button-element:assert-2
    x-internal="assert"}: `state`{.variable} is not in the [Submit
    Button](#attr-button-type-submit-state){#the-button-element:attr-button-type-submit-state-2}
    state.

4.  If `state`{.variable} is in the
    [Auto](#attr-button-type-auto-state){#the-button-element:attr-button-type-auto-state-3}
    state, then return \"`button`\".

5.  Return the keyword value corresponding to `state`{.variable}.

The [`type`{#the-button-element:dom-button-type-2}](#dom-button-type)
setter steps are to set the
[`type`{#the-button-element:attr-button-type-7}](#attr-button-type)
content attribute to the given value.

The
[`willValidate`{#the-button-element:dom-cva-willvalidate-2}](form-control-infrastructure.html#dom-cva-willvalidate),
[`validity`{#the-button-element:dom-cva-validity-2}](form-control-infrastructure.html#dom-cva-validity),
and
[`validationMessage`{#the-button-element:dom-cva-validationmessage-2}](form-control-infrastructure.html#dom-cva-validationmessage)
IDL attributes, and the
[`checkValidity()`{#the-button-element:dom-cva-checkvalidity-2}](form-control-infrastructure.html#dom-cva-checkvalidity),
[`reportValidity()`{#the-button-element:dom-cva-reportvalidity-2}](form-control-infrastructure.html#dom-cva-reportvalidity),
and
[`setCustomValidity()`{#the-button-element:dom-cva-setcustomvalidity-2}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
methods, are part of the [constraint validation
API](form-control-infrastructure.html#the-constraint-validation-api){#the-button-element:the-constraint-validation-api}.
The
[`labels`{#the-button-element:dom-lfe-labels-2}](forms.html#dom-lfe-labels)
IDL attribute provides a list of the element\'s
[`label`{#the-button-element:the-label-element}](forms.html#the-label-element)s.
The
[`disabled`{#the-button-element:dom-fe-disabled-2}](form-control-infrastructure.html#dom-fe-disabled),
[`form`{#the-button-element:dom-fae-form-2}](form-control-infrastructure.html#dom-fae-form),
and
[`name`{#the-button-element:dom-fe-name-2}](form-control-infrastructure.html#dom-fe-name)
IDL attributes are part of the element\'s forms API.

::: example
The following button is labeled \"Show hint\" and pops up a dialog box
when activated:

``` html
<button type=button
        onclick="alert('This 15-20 minute piece was composed by George Gershwin.')">
 Show hint
</button>
```
:::

::: example
The following shows how
[buttons](forms.html#concept-button){#the-button-element:concept-button-3}
can use
[`commandfor`{#the-button-element:attr-button-commandfor-5}](#attr-button-commandfor)
to show and hide an element with the
[`popover`{#the-button-element:attr-popover-5}](popover.html#attr-popover)
attribute when activated:

``` html
<button type=button
        commandfor="the-popover"
        command="show-popover">
 Show menu
</button>
<div popover
     id="the-popover">
 <button commandfor="the-popover"
         command="hide-popover">
  Hide menu
 </button>
</div>
    
```
:::

::: example
The following shows how buttons can use
[`commandfor`{#the-button-element:attr-button-commandfor-6}](#attr-button-commandfor)
with a [custom command
keyword](#attr-button-command-custom){#the-button-element:attr-button-command-custom-2}
on an element, demonstrating how one could use custom commands for
unspecified behavior:

``` html
<button type=button
        commandfor="the-image"
        command="--rotate-landscape">
 Rotate Left
</button>
<button type=button
        commandfor="the-image"
        command="--rotate-portrait">
 Rotate Right
</button>
<img id="the-image"
     src="photo.jpg">
<script>
  const image = document.getElementById("the-image");
  image.addEventListener("command", (event) => {
   if ( event.command == "--rotate-landscape" ) {
    event.target.style.rotate = "-90deg"
   } else if ( event.command == "--rotate-portrait" ) {
    event.target.style.rotate = "0deg"
   }
  });
</script>
    
```
:::

#### [4.10.7]{.secno} The [`select`]{.dfn dfn-type="element"} element[](#the-select-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/select](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select "The <select> HTML element represents a control that provides a menu of options.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera2+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement "The HTMLSelectElement interface represents a <select> HTML Element. These elements also share all of the properties and methods of other HTML elements via the HTMLElement interface.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera2+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer1+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-select-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-select-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-select-element:phrasing-content-2}.
:   [Interactive
    content](dom.html#interactive-content-2){#the-select-element:interactive-content-2}.
:   [Listed](forms.html#category-listed){#the-select-element:category-listed},
    [labelable](forms.html#category-label){#the-select-element:category-label},
    [submittable](forms.html#category-submit){#the-select-element:category-submit},
    [resettable](forms.html#category-reset){#the-select-element:category-reset},
    and [autocapitalize-and-autocorrect
    inheriting](forms.html#category-autocapitalize){#the-select-element:category-autocapitalize}
    [form-associated
    element](forms.html#form-associated-element){#the-select-element:form-associated-element}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-select-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-select-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-select-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-select-element:concept-element-content-model}:
:   Zero or more
    [`option`{#the-select-element:the-option-element}](#the-option-element),
    [`optgroup`{#the-select-element:the-optgroup-element}](#the-optgroup-element),
    [`hr`{#the-select-element:the-hr-element}](grouping-content.html#the-hr-element),
    and
    [script-supporting](dom.html#script-supporting-elements-2){#the-select-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-select-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-select-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-select-element:global-attributes}
:   [`autocomplete`{#the-select-element:attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete)
    --- Hint for form autofill feature
:   [`disabled`{#the-select-element:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
    --- Whether the form control is disabled
:   [`form`{#the-select-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    --- Associates the element with a
    [`form`{#the-select-element:the-form-element}](forms.html#the-form-element)
    element
:   [`multiple`{#the-select-element:attr-select-multiple}](#attr-select-multiple)
    --- Whether to allow multiple values
:   [`name`{#the-select-element:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    --- Name of the element to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-select-element:form-submission-2}
    and in the
    [`form.elements`{#the-select-element:dom-form-elements}](forms.html#dom-form-elements)
    API
:   [`required`{#the-select-element:attr-select-required}](#attr-select-required)
    --- Whether the control is required for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-select-element:form-submission-2-2}
:   [`size`{#the-select-element:attr-select-size}](#attr-select-size)
    --- Size of the control

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-select-element:concept-element-accessibility-considerations}:
:   If the element has a
    [`multiple`{#the-select-element:attr-select-multiple-2}](#attr-select-multiple)
    attribute or a
    [`size`{#the-select-element:attr-select-size-2}](#attr-select-size)
    attribute with a value \> 1: [for
    authors](https://w3c.github.io/html-aria/#el-select-multiple-or-size-greater-1);
    [for
    implementers](https://w3c.github.io/html-aam/#el-select-listbox).
:   Otherwise: [for
    authors](https://w3c.github.io/html-aria/#el-select); [for
    implementers](https://w3c.github.io/html-aam/#el-select-combobox).

[DOM interface](dom.html#concept-element-dom){#the-select-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLSelectElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString autocomplete;
      [CEReactions] attribute boolean disabled;
      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute boolean multiple;
      [CEReactions] attribute DOMString name;
      [CEReactions] attribute boolean required;
      [CEReactions] attribute unsigned long size;

      readonly attribute DOMString type;

      [SameObject] readonly attribute HTMLOptionsCollection options;
      [CEReactions] attribute unsigned long length;
      getter HTMLOptionElement? item(unsigned long index);
      HTMLOptionElement? namedItem(DOMString name);
      [CEReactions] undefined add((HTMLOptionElement or HTMLOptGroupElement) element, optional (HTMLElement or long)? before = null);
      [CEReactions] undefined remove(); // ChildNode overload
      [CEReactions] undefined remove(long index);
      [CEReactions] setter undefined (unsigned long index, HTMLOptionElement? option);

      [SameObject] readonly attribute HTMLCollection selectedOptions;
      attribute long selectedIndex;
      attribute DOMString value;

      readonly attribute boolean willValidate;
      readonly attribute ValidityState validity;
      readonly attribute DOMString validationMessage;
      boolean checkValidity();
      boolean reportValidity();
      undefined setCustomValidity(DOMString error);

      undefined showPicker();

      readonly attribute NodeList labels;
    };
    ```

The
[`select`{#the-select-element:the-select-element}](#the-select-element)
element represents a control for selecting amongst a set of options.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/multiple](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/multiple "The Boolean multiple attribute, if set, means the form control accepts one or more values. Valid for the email and file input types and the <select>, the manner by which the user opts for multiple values depends on the form control.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

The [`multiple`]{#attr-select-multiple .dfn dfn-for="select"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-select-element:boolean-attribute}.
If the attribute is present, then the
[`select`{#the-select-element:the-select-element-2}](#the-select-element)
element
[represents](dom.html#represents){#the-select-element:represents} a
control for selecting zero or more options from the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list}.
If the attribute is absent, then the
[`select`{#the-select-element:the-select-element-3}](#the-select-element)
element
[represents](dom.html#represents){#the-select-element:represents-2} a
control for selecting a single option from the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-2}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/size](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/size "The size attribute defines the width of the <input> and the height of the <select> element. For the input, if the type attribute is text or password then it's the number of characters. This must be an integer value of 0 or higher. If no size is specified, or an invalid value is specified, the input has no size declared, and the form control will be the default width based on the user agent. If CSS targets the element with properties impacting the width, CSS takes precedence.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

The [`size`]{#attr-select-size .dfn dfn-for="select"
dfn-type="element-attr"} attribute gives the number of options to show
to the user. The
[`size`{#the-select-element:attr-select-size-3}](#attr-select-size)
attribute, if specified, must have a value that is a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#the-select-element:valid-non-negative-integer}
greater than zero.

The [display size]{#concept-select-size .dfn} of a
[`select`{#the-select-element:the-select-element-4}](#the-select-element)
element is the result of applying the [rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-select-element:rules-for-parsing-non-negative-integers}
to the value of element\'s
[`size`{#the-select-element:attr-select-size-4}](#attr-select-size)
attribute, if it has one and parsing it is successful. If applying those
rules to the attribute\'s value is not successful, or if the
[`size`{#the-select-element:attr-select-size-5}](#attr-select-size)
attribute is absent, then the element\'s [display
size](#concept-select-size){#the-select-element:concept-select-size} is
4 if the element\'s
[`multiple`{#the-select-element:attr-select-multiple-3}](#attr-select-multiple)
content attribute is present, and 1 otherwise.

The [list of options]{#concept-select-option-list .dfn} for a
[`select`{#the-select-element:the-select-element-5}](#the-select-element)
element consists of all the
[`option`{#the-select-element:the-option-element-2}](#the-option-element)
element children of the
[`select`{#the-select-element:the-select-element-6}](#the-select-element)
element, and all the
[`option`{#the-select-element:the-option-element-3}](#the-option-element)
element children of all the
[`optgroup`{#the-select-element:the-optgroup-element-2}](#the-optgroup-element)
element children of the
[`select`{#the-select-element:the-select-element-7}](#the-select-element)
element, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-select-element:tree-order
x-internal="tree-order"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/required](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/required "The Boolean required attribute, if present, indicates that the user must specify a value for the input before the owning form can be submitted.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`required`]{#attr-select-required .dfn dfn-for="select"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-select-element:boolean-attribute-2}.
When specified, the user will be required to select a value before
submitting the form.

If a
[`select`{#the-select-element:the-select-element-8}](#the-select-element)
element has a
[`required`{#the-select-element:attr-select-required-2}](#attr-select-required)
attribute specified, does not have a
[`multiple`{#the-select-element:attr-select-multiple-4}](#attr-select-multiple)
attribute specified, and has a [display
size](#concept-select-size){#the-select-element:concept-select-size-2}
of 1; and if the
[value](#concept-option-value){#the-select-element:concept-option-value}
of the first
[`option`{#the-select-element:the-option-element-4}](#the-option-element)
element in the
[`select`{#the-select-element:the-select-element-9}](#the-select-element)
element\'s [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-3}
(if any) is the empty string, and that
[`option`{#the-select-element:the-option-element-5}](#the-option-element)
element\'s parent node is the
[`select`{#the-select-element:the-select-element-10}](#the-select-element)
element (and not an
[`optgroup`{#the-select-element:the-optgroup-element-3}](#the-optgroup-element)
element), then that
[`option`{#the-select-element:the-option-element-6}](#the-option-element)
is the
[`select`{#the-select-element:the-select-element-11}](#the-select-element)
element\'s [placeholder label option]{#placeholder-label-option .dfn}.

If a
[`select`{#the-select-element:the-select-element-12}](#the-select-element)
element has a
[`required`{#the-select-element:attr-select-required-3}](#attr-select-required)
attribute specified, does not have a
[`multiple`{#the-select-element:attr-select-multiple-5}](#attr-select-multiple)
attribute specified, and has a [display
size](#concept-select-size){#the-select-element:concept-select-size-3}
of 1, then the
[`select`{#the-select-element:the-select-element-13}](#the-select-element)
element must have a [placeholder label
option](#placeholder-label-option){#the-select-element:placeholder-label-option}.

In practice, the requirement stated in the paragraph above can only
apply when a
[`select`{#the-select-element:the-select-element-14}](#the-select-element)
element does not have a
[`size`{#the-select-element:attr-select-size-6}](#attr-select-size)
attribute with a value greater than 1.

**Constraint validation**: If the element has its
[`required`{#the-select-element:attr-select-required-4}](#attr-select-required)
attribute specified, and either none of the
[`option`{#the-select-element:the-option-element-7}](#the-option-element)
elements in the
[`select`{#the-select-element:the-select-element-15}](#the-select-element)
element\'s [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-4}
have their
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness}
set to true, or the only
[`option`{#the-select-element:the-option-element-8}](#the-option-element)
element in the
[`select`{#the-select-element:the-select-element-16}](#the-select-element)
element\'s [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-5}
with its
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-2}
set to true is the [placeholder label
option](#placeholder-label-option){#the-select-element:placeholder-label-option-2},
then the element is [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#the-select-element:suffering-from-being-missing}.

If the
[`multiple`{#the-select-element:attr-select-multiple-6}](#attr-select-multiple)
attribute is absent, and the element is not
[disabled](form-control-infrastructure.html#concept-fe-disabled){#the-select-element:concept-fe-disabled},
then the user agent should allow the user to pick an
[`option`{#the-select-element:the-option-element-9}](#the-option-element)
element in its [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-6}
that is itself not
[disabled](#concept-option-disabled){#the-select-element:concept-option-disabled}.
Upon this
[`option`{#the-select-element:the-option-element-10}](#the-option-element)
element being [picked]{#concept-select-pick .dfn} (either through a
click, or through unfocusing the element after changing its value, or
through a [menu
command](interactive-elements.html#using-the-option-element-to-define-a-command){#the-select-element:using-the-option-element-to-define-a-command},
or through any other mechanism), and before the relevant user
interaction event is queued (e.g. before the
[`click`{#the-select-element:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event), the user agent must set the
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-3}
of the picked
[`option`{#the-select-element:the-option-element-11}](#the-option-element)
element to true, set its
[dirtiness](#concept-option-dirtiness){#the-select-element:concept-option-dirtiness}
to true, and then [send `select` update
notifications](#send-select-update-notifications){#the-select-element:send-select-update-notifications}.

If the
[`multiple`{#the-select-element:attr-select-multiple-7}](#attr-select-multiple)
attribute is absent, whenever an
[`option`{#the-select-element:the-option-element-12}](#the-option-element)
element in the
[`select`{#the-select-element:the-select-element-17}](#the-select-element)
element\'s [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-7}
has its
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-4}
set to true, and whenever an
[`option`{#the-select-element:the-option-element-13}](#the-option-element)
element with its
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-5}
set to true is added to the
[`select`{#the-select-element:the-select-element-18}](#the-select-element)
element\'s [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-8},
the user agent must set the
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-6}
of all the other
[`option`{#the-select-element:the-option-element-14}](#the-option-element)
elements in its [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-9}
to false.

If the
[`multiple`{#the-select-element:attr-select-multiple-8}](#attr-select-multiple)
attribute is absent and the element\'s [display
size](#concept-select-size){#the-select-element:concept-select-size-4}
is greater than 1, then the user agent should also allow the user to
request that the
[`option`{#the-select-element:the-option-element-15}](#the-option-element)
whose
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-7}
is true, if any, be unselected. Upon this request being conveyed to the
user agent, and before the relevant user interaction event is queued
(e.g. before the
[`click`{#the-select-element:event-click-2}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event), the user agent must set the
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-8}
of that
[`option`{#the-select-element:the-option-element-16}](#the-option-element)
element to false, set its
[dirtiness](#concept-option-dirtiness){#the-select-element:concept-option-dirtiness-2}
to true, and then [send `select` update
notifications](#send-select-update-notifications){#the-select-element:send-select-update-notifications-2}.

The [selectedness setting algorithm]{#selectedness-setting-algorithm
.dfn}, given a
[`select`{#the-select-element:the-select-element-19}](#the-select-element)
element `element`{.variable}, is to run the following steps:

1.  If `element`{.variable}\'s
    [`multiple`{#the-select-element:attr-select-multiple-9}](#attr-select-multiple)
    attribute is absent, and `element`{.variable}\'s [display
    size](#concept-select-size){#the-select-element:concept-select-size-5}
    is 1, and no
    [`option`{#the-select-element:the-option-element-17}](#the-option-element)
    elements in the `element`{.variable}\'s [list of
    options](#concept-select-option-list){#the-select-element:concept-select-option-list-10}
    have their
    [selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-9}
    set to true, then set the
    [selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-10}
    of the first
    [`option`{#the-select-element:the-option-element-18}](#the-option-element)
    element in the [list of
    options](#concept-select-option-list){#the-select-element:concept-select-option-list-11}
    in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-select-element:tree-order-2
    x-internal="tree-order"} that is not
    [disabled](#concept-option-disabled){#the-select-element:concept-option-disabled-2},
    if any, to true, and return.

2.  If `element`{.variable}\'s
    [`multiple`{#the-select-element:attr-select-multiple-10}](#attr-select-multiple)
    attribute is absent, and two or more
    [`option`{#the-select-element:the-option-element-19}](#the-option-element)
    elements in `element`{.variable}\'s [list of
    options](#concept-select-option-list){#the-select-element:concept-select-option-list-12}
    have their
    [selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-11}
    set to true, then set the
    [selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-12}
    of all but the last
    [`option`{#the-select-element:the-option-element-20}](#the-option-element)
    element with its
    [selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-13}
    set to true in the [list of
    options](#concept-select-option-list){#the-select-element:concept-select-option-list-13}
    in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-select-element:tree-order-3
    x-internal="tree-order"} to false.

The
[`option`{#the-select-element:the-option-element-21}](#the-option-element)
[HTML element insertion
steps](infrastructure.html#html-element-insertion-steps){#the-select-element:html-element-insertion-steps},
given `insertedNode`{.variable}, are:

1.  If `insertedNode`{.variable}\'s parent is a
    [`select`{#the-select-element:the-select-element-20}](#the-select-element)
    element, or `insertedNode`{.variable}\'s parent is an
    [`optgroup`{#the-select-element:the-optgroup-element-4}](#the-optgroup-element)
    element whose parent is a
    [`select`{#the-select-element:the-select-element-21}](#the-select-element)
    element, then run that
    [`select`{#the-select-element:the-select-element-22}](#the-select-element)
    element\'s [selectedness setting
    algorithm](#selectedness-setting-algorithm){#the-select-element:selectedness-setting-algorithm}.

The
[`option`{#the-select-element:the-option-element-22}](#the-option-element)
[HTML element removing
steps](infrastructure.html#html-element-removing-steps){#the-select-element:html-element-removing-steps},
given `removedNode`{.variable} and `oldParent`{.variable}, are:

1.  If `oldParent`{.variable} is a
    [`select`{#the-select-element:the-select-element-23}](#the-select-element)
    element, or `oldParent`{.variable} is an
    [`optgroup`{#the-select-element:the-optgroup-element-5}](#the-optgroup-element)
    element whose parent is a
    [`select`{#the-select-element:the-select-element-24}](#the-select-element)
    element, then run that
    [`select`{#the-select-element:the-select-element-25}](#the-select-element)
    element\'s [selectedness setting
    algorithm](#selectedness-setting-algorithm){#the-select-element:selectedness-setting-algorithm-2}.

The
[`option`{#the-select-element:the-option-element-23}](#the-option-element)
[HTML element moving
steps](infrastructure.html#html-element-moving-steps){#the-select-element:html-element-moving-steps},
given `movedNode`{.variable} and `oldParent`{.variable}, are:

1.  Run the
    [`option`{#the-select-element:the-option-element-24}](#the-option-element)
    [HTML element removing
    steps](infrastructure.html#html-element-removing-steps){#the-select-element:html-element-removing-steps-2}
    given `movedNode`{.variable} and `oldParent`{.variable}.

2.  Run the
    [`option`{#the-select-element:the-option-element-25}](#the-option-element)
    [HTML element insertion
    steps](infrastructure.html#html-element-insertion-steps){#the-select-element:html-element-insertion-steps-2}
    given `movedNode`{.variable}.

The
[`optgroup`{#the-select-element:the-optgroup-element-6}](#the-optgroup-element)
[HTML element removing
steps](infrastructure.html#html-element-removing-steps){#the-select-element:html-element-removing-steps-3},
given `removedNode`{.variable} and `oldParent`{.variable}, are:

1.  If `oldParent`{.variable} is a
    [`select`{#the-select-element:the-select-element-26}](#the-select-element)
    element and `removedNode`{.variable} has an
    [`option`{#the-select-element:the-option-element-26}](#the-option-element)
    child, then run `oldParent`{.variable}\'s [selectedness setting
    algorithm](#selectedness-setting-algorithm){#the-select-element:selectedness-setting-algorithm-3}.

The
[`optgroup`{#the-select-element:the-optgroup-element-7}](#the-optgroup-element)
[HTML element moving
steps](infrastructure.html#html-element-moving-steps){#the-select-element:html-element-moving-steps-2},
given `movedNode`{.variable} and `oldParent`{.variable}, are:

1.  Run the
    [`optgroup`{#the-select-element:the-optgroup-element-8}](#the-optgroup-element)
    [HTML element removing
    steps](infrastructure.html#html-element-removing-steps){#the-select-element:html-element-removing-steps-4}
    given `movedNode`{.variable} and `oldParent`{.variable}.

If an
[`option`{#the-select-element:the-option-element-27}](#the-option-element)
element in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-14}
[asks for a reset]{#ask-for-a-reset .dfn}, then run that
[`select`{#the-select-element:the-select-element-27}](#the-select-element)
element\'s [selectedness setting
algorithm](#selectedness-setting-algorithm){#the-select-element:selectedness-setting-algorithm-4}.

If the
[`multiple`{#the-select-element:attr-select-multiple-11}](#attr-select-multiple)
attribute is present, and the element is not
[disabled](form-control-infrastructure.html#concept-fe-disabled){#the-select-element:concept-fe-disabled-2},
then the user agent should allow the user to
[toggle]{#concept-select-toggle .dfn} the
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-14}
of the
[`option`{#the-select-element:the-option-element-28}](#the-option-element)
elements in its [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-15}
that are themselves not
[disabled](#concept-option-disabled){#the-select-element:concept-option-disabled-3}.
Upon such an element being
[toggled](#concept-select-toggle){#the-select-element:concept-select-toggle}
(either through a click, or through a [menu
command](interactive-elements.html#using-the-option-element-to-define-a-command){#the-select-element:using-the-option-element-to-define-a-command-2},
or any other mechanism), and before the relevant user interaction event
is queued (e.g. before a related
[`click`{#the-select-element:event-click-3}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
event), the
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-15}
of the
[`option`{#the-select-element:the-option-element-29}](#the-option-element)
element must be changed (from true to false or false to true), the
[dirtiness](#concept-option-dirtiness){#the-select-element:concept-option-dirtiness-3}
of the element must be set to true, and the user agent must [send
`select` update
notifications](#send-select-update-notifications){#the-select-element:send-select-update-notifications-3}.

When the user agent is to [send `select` update
notifications]{#send-select-update-notifications .dfn}, [queue an
element
task](webappapis.html#queue-an-element-task){#the-select-element:queue-an-element-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#the-select-element:user-interaction-task-source}
given the
[`select`{#the-select-element:the-select-element-28}](#the-select-element)
element to run these steps:

1.  Set the
    [`select`{#the-select-element:the-select-element-29}](#the-select-element)
    element\'s [user
    validity](form-control-infrastructure.html#user-validity){#the-select-element:user-validity}
    to true.

2.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-select-element:concept-event-fire
    x-internal="concept-event-fire"} named
    [`input`{#the-select-element:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    at the
    [`select`{#the-select-element:the-select-element-30}](#the-select-element)
    element, with the
    [`bubbles`{#the-select-element:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    and
    [`composed`{#the-select-element:dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
    attributes initialized to true.

3.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-select-element:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`change`{#the-select-element:event-change}](indices.html#event-change)
    at the
    [`select`{#the-select-element:the-select-element-31}](#the-select-element)
    element, with the
    [`bubbles`{#the-select-element:dom-event-bubbles-2}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    attribute initialized to true.

The [reset
algorithm](form-control-infrastructure.html#concept-form-reset-control){#the-select-element:concept-form-reset-control}
for a
[`select`{#the-select-element:the-select-element-32}](#the-select-element)
element `selectElement`{.variable} is:

1.  Set `selectElement`{.variable}\'s [user
    validity](form-control-infrastructure.html#user-validity){#the-select-element:user-validity-2}
    to false.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-select-element:list-iterate
    x-internal="list-iterate"} `optionElement`{.variable} of
    `selectElement`{.variable}\'s [list of
    options](#concept-select-option-list){#the-select-element:concept-select-option-list-16}:

    1.  If `optionElement`{.variable} has a
        [`selected`{#the-select-element:attr-option-selected}](#attr-option-selected)
        attribute, then set `optionElement`{.variable}\'s
        [selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-16}
        to true; otherwise set it to false.

    2.  Set `optionElement`{.variable}\'s
        [dirtiness](#concept-option-dirtiness){#the-select-element:concept-option-dirtiness-4}
        to false.

3.  Run the [selectedness setting
    algorithm](#selectedness-setting-algorithm){#the-select-element:selectedness-setting-algorithm-5}
    given `selectElement`{.variable}.

The
[`form`{#the-select-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the
[`select`{#the-select-element:the-select-element-33}](#the-select-element)
element with its [form
owner](form-control-infrastructure.html#form-owner){#the-select-element:form-owner}.
The
[`name`{#the-select-element:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attribute represents the element\'s name. The
[`disabled`{#the-select-element:attr-fe-disabled-2}](form-control-infrastructure.html#attr-fe-disabled)
attribute is used to make the control non-interactive and to prevent its
value from being submitted. The
[`autocomplete`{#the-select-element:attr-fe-autocomplete-2}](form-control-infrastructure.html#attr-fe-autocomplete)
attribute controls how the user agent provides autofill behavior.

A
[`select`{#the-select-element:the-select-element-34}](#the-select-element)
element that is not
[disabled](form-control-infrastructure.html#concept-fe-disabled){#the-select-element:concept-fe-disabled-3}
is *[mutable](form-control-infrastructure.html#concept-fe-mutable)*.

`select`{.variable}`.`[`type`](#dom-select-type){#dom-select-type-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/type](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/type "The HTMLSelectElement.type read-only property returns the form control's type.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera2+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer1+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns \"`select-multiple`\" if the element has a
[`multiple`{#the-select-element:attr-select-multiple-12}](#attr-select-multiple)
attribute, and \"`select-one`\" otherwise.

`select`{.variable}`.`[`options`](#dom-select-options){#dom-select-options-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/options](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/options "The HTMLSelectElement.options read-only property returns a HTMLOptionsCollection of the <option> elements contained by the <select> element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

Returns an
[`HTMLOptionsCollection`{#the-select-element:htmloptionscollection-2}](common-dom-interfaces.html#htmloptionscollection)
of the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-17}.

`select`{.variable}`.`[`length`](#dom-select-length){#dom-select-length-dev}` [ = ``value`{.variable}` ]`

Returns the number of elements in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-18}.

When set to a smaller number, truncates the number of
[`option`{#the-select-element:the-option-element-30}](#the-option-element)
elements in the
[`select`{#the-select-element:the-select-element-35}](#the-select-element).

When set to a greater number, adds new blank
[`option`{#the-select-element:the-option-element-31}](#the-option-element)
elements to the
[`select`{#the-select-element:the-select-element-36}](#the-select-element).

`element`{.variable}` = ``select`{.variable}`.`[`item`](#dom-select-item){#dom-select-item-dev}`(``index`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/item](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/item "The HTMLSelectElement.item() method returns the Element corresponding to the HTMLOptionElement whose position in the options list corresponds to the index given in the parameter, or null if there are none.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

`select`{.variable}`[``index`{.variable}`]`

Returns the item with index `index`{.variable} from the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-19}.
The items are sorted in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-select-element:tree-order-4
x-internal="tree-order"}.

`element`{.variable}` = ``select`{.variable}`.`[`namedItem`](#dom-select-nameditem){#dom-select-nameditem-dev}`(``name`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/namedItem](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/namedItem "The HTMLSelectElement.namedItem() method returns the HTMLOptionElement corresponding to the HTMLOptionElement whose name or id match the specified name, or null if no option matches.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the first item with
[ID](https://dom.spec.whatwg.org/#concept-id){#the-select-element:concept-id
x-internal="concept-id"} or
[`name`{#the-select-element:attr-option-name}](obsolete.html#attr-option-name)
`name`{.variable} from the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-20}.

Returns null if no element with that
[ID](https://dom.spec.whatwg.org/#concept-id){#the-select-element:concept-id-2
x-internal="concept-id"} could be found.

`select`{.variable}`.`[`add`](#dom-select-add){#dom-select-add-dev}`(``element`{.variable}` [, ``before`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/add](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/add "The HTMLSelectElement.add() method adds an element to the collection of option elements for this select element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

Inserts `element`{.variable} before the node given by
`before`{.variable}.

The `before`{.variable} argument can be a number, in which case
`element`{.variable} is inserted before the item with that number, or an
element from the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-21},
in which case `element`{.variable} is inserted before that element.

If `before`{.variable} is omitted, null, or a number out of range, then
`element`{.variable} will be added at the end of the list.

This method will throw a
[\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#the-select-element:hierarchyrequesterror
x-internal="hierarchyrequesterror"}
[`DOMException`{#the-select-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `element`{.variable} is an ancestor of the element into which it is
to be inserted.

`select`{.variable}`.`[`selectedOptions`](#dom-select-selectedoptions){#dom-select-selectedoptions-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/selectedOptions](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/selectedOptions "The read-only HTMLSelectElement property selectedOptions contains a list of the <option> elements contained within the <select> element that are currently selected. The list of selected options is an HTMLCollection object with one entry per currently selected option.")

Support in all current engines.

::: support
[Firefox26+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome19+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns an
[`HTMLCollection`{#the-select-element:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-22}
that are selected.

`select`{.variable}`.`[`selectedIndex`](#dom-select-selectedindex){#dom-select-selectedindex-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/selectedIndex](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/selectedIndex "The HTMLSelectElement.selectedIndex property is a long that reflects the index of the first or last selected <option> element, depending on the value of multiple. The value -1 indicates that no element is selected.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

Returns the index of the first selected item, if any, or −1 if there is
no selected item.

Can be set, to change the selection.

`select`{.variable}`.`[`value`](#dom-select-value){#dom-select-value-dev}` [ = ``value`{.variable}` ]`

Returns the
[value](#concept-option-value){#the-select-element:concept-option-value-2}
of the first selected item, if any, or the empty string if there is no
selected item.

Can be set, to change the selection.

`select`{.variable}`.`[`showPicker`](input.html#dom-select-showpicker){#dom-select-showpicker-dev}`()`

Shows any applicable picker UI for `select`{.variable}, so that the user
can select a value.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-select-element:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#the-select-element:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `select`{.variable} is not
[mutable](form-control-infrastructure.html#concept-fe-mutable){#the-select-element:concept-fe-mutable-2}.

Throws a
[\"`NotAllowedError`\"](https://webidl.spec.whatwg.org/#notallowederror){#the-select-element:notallowederror
x-internal="notallowederror"}
[`DOMException`{#the-select-element:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if called without [transient user
activation](interaction.html#transient-activation){#the-select-element:transient-activation}.

Throws a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-select-element:securityerror
x-internal="securityerror"}
[`DOMException`{#the-select-element:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `select`{.variable} is inside a cross-origin
[`iframe`{#the-select-element:the-iframe-element}](iframe-embed-object.html#the-iframe-element).

Throws a
[\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-select-element:notsupportederror
x-internal="notsupportederror"}
[`DOMException`{#the-select-element:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `select`{.variable} is not [being
rendered](rendering.html#being-rendered){#the-select-element:being-rendered}.

The [`type`]{#dom-select-type .dfn dfn-for="HTMLSelectElement"
dfn-type="attribute"} IDL attribute, on getting, must return the string
\"`select-one`\" if the
[`multiple`{#the-select-element:attr-select-multiple-13}](#attr-select-multiple)
attribute is absent, and the string \"`select-multiple`\" if the
[`multiple`{#the-select-element:attr-select-multiple-14}](#attr-select-multiple)
attribute is present.

The [`options`]{#dom-select-options .dfn dfn-for="HTMLSelectElement"
dfn-type="attribute"} IDL attribute must return an
[`HTMLOptionsCollection`{#the-select-element:htmloptionscollection-3}](common-dom-interfaces.html#htmloptionscollection)
rooted at the
[`select`{#the-select-element:the-select-element-37}](#the-select-element)
node, whose filter matches the elements in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-23}.

The
[`options`{#the-select-element:dom-select-options-2}](#dom-select-options)
collection is also mirrored on the
[`HTMLSelectElement`{#the-select-element:htmlselectelement}](#htmlselectelement)
object. The [supported property
indices](https://webidl.spec.whatwg.org/#dfn-supported-property-indices){#the-select-element:supported-property-indices
x-internal="supported-property-indices"} at any instant are the indices
supported by the object returned by the
[`options`{#the-select-element:dom-select-options-3}](#dom-select-options)
attribute at that instant.

The [`length`]{#dom-select-length .dfn dfn-for="HTMLSelectElement"
dfn-type="attribute"} IDL attribute must return the number of nodes
[represented](https://dom.spec.whatwg.org/#represented-by-the-collection){#the-select-element:represented-by-the-collection
x-internal="represented-by-the-collection"} by the
[`options`{#the-select-element:dom-select-options-4}](#dom-select-options)
collection. On setting, it must act like the attribute of the same name
on the
[`options`{#the-select-element:dom-select-options-5}](#dom-select-options)
collection.

The [`item(``index`{.variable}`)`]{#dom-select-item .dfn
dfn-for="HTMLSelectElement" dfn-type="method"} method must return the
value returned by [the method of the same
name](https://dom.spec.whatwg.org/#dom-htmlcollection-item){#the-select-element:dom-htmlcollection-item
x-internal="dom-htmlcollection-item"} on the
[`options`{#the-select-element:dom-select-options-6}](#dom-select-options)
collection, when invoked with the same argument.

The [`namedItem(``name`{.variable}`)`]{#dom-select-nameditem .dfn
dfn-for="HTMLSelectElement" dfn-type="method"} method must return the
value returned by [the method of the same
name](https://dom.spec.whatwg.org/#dom-htmlcollection-nameditem){#the-select-element:dom-htmlcollection-nameditem
x-internal="dom-htmlcollection-nameditem"} on the
[`options`{#the-select-element:dom-select-options-7}](#dom-select-options)
collection, when invoked with the same argument.

When the user agent is to [set the value of a new indexed
property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-a-new-indexed-property){#the-select-element:set-the-value-of-a-new-indexed-property
x-internal="set-the-value-of-a-new-indexed-property"} or [set the value
of an existing indexed
property](https://webidl.spec.whatwg.org/#dfn-set-the-value-of-an-existing-indexed-property){#the-select-element:set-the-value-of-an-existing-indexed-property
x-internal="set-the-value-of-an-existing-indexed-property"} for a
[`select`{#the-select-element:the-select-element-38}](#the-select-element)
element, it must instead run [the corresponding
algorithm](common-dom-interfaces.html#dom-htmloptionscollection-setter)
on the
[`select`{#the-select-element:the-select-element-39}](#the-select-element)
element\'s
[`options`{#the-select-element:dom-select-options-8}](#dom-select-options)
collection.

Similarly, the
[`add(``element`{.variable}`, ``before`{.variable}`)`]{#dom-select-add
.dfn dfn-for="HTMLSelectElement" dfn-type="method"} method must act like
its namesake method on that same
[`options`{#the-select-element:dom-select-options-9}](#dom-select-options)
collection.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSelectElement/remove](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/remove "The HTMLSelectElement.remove() method removes the element at the specified index from the options collection for this select element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

The [`remove()`]{#dom-select-remove .dfn dfn-for="HTMLSelectElement"
dfn-type="method"} method must act like its namesake method on that same
[`options`{#the-select-element:dom-select-options-10}](#dom-select-options)
collection when it has arguments, and like its namesake method on the
[`ChildNode`{#the-select-element:childnode}](https://dom.spec.whatwg.org/#interface-childnode){x-internal="childnode"}
interface implemented by the
[`HTMLSelectElement`{#the-select-element:htmlselectelement-2}](#htmlselectelement)
ancestor interface
[`Element`{#the-select-element:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
when it has no arguments.

The [`selectedOptions`]{#dom-select-selectedoptions .dfn
dfn-for="HTMLSelectElement" dfn-type="attribute"} IDL attribute must
return an
[`HTMLCollection`{#the-select-element:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`select`{#the-select-element:the-select-element-40}](#the-select-element)
node, whose filter matches the elements in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-24}
that have their
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-17}
set to true.

The [`selectedIndex`]{#dom-select-selectedindex .dfn
dfn-for="HTMLSelectElement" dfn-type="attribute"} IDL attribute, on
getting, must return the
[index](#concept-option-index){#the-select-element:concept-option-index}
of the first
[`option`{#the-select-element:the-option-element-32}](#the-option-element)
element in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-25}
in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-select-element:tree-order-5
x-internal="tree-order"} that has its
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-18}
set to true, if any. If there isn\'t one, then it must return −1.

On setting, the
[`selectedIndex`{#the-select-element:dom-select-selectedindex-2}](#dom-select-selectedindex)
attribute must set the
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-19}
of all the
[`option`{#the-select-element:the-option-element-33}](#the-option-element)
elements in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-26}
to false, and then the
[`option`{#the-select-element:the-option-element-34}](#the-option-element)
element in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-27}
whose
[index](#concept-option-index){#the-select-element:concept-option-index-2}
is the given new value, if any, must have its
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-20}
set to true and its
[dirtiness](#concept-option-dirtiness){#the-select-element:concept-option-dirtiness-5}
set to true.

This can result in no element having a
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-21}
set to true even in the case of the
[`select`{#the-select-element:the-select-element-41}](#the-select-element)
element having no
[`multiple`{#the-select-element:attr-select-multiple-15}](#attr-select-multiple)
attribute and a [display
size](#concept-select-size){#the-select-element:concept-select-size-6}
of 1.

The [`value`]{#dom-select-value .dfn dfn-for="HTMLSelectElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
[value](#concept-option-value){#the-select-element:concept-option-value-3}
of the first
[`option`{#the-select-element:the-option-element-35}](#the-option-element)
element in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-28}
in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-select-element:tree-order-6
x-internal="tree-order"} that has its
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-22}
set to true, if any. If there isn\'t one, then it must return the empty
string.

On setting, the
[`value`{#the-select-element:dom-select-value-2}](#dom-select-value)
attribute must set the
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-23}
of all the
[`option`{#the-select-element:the-option-element-36}](#the-option-element)
elements in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-29}
to false, and then the first
[`option`{#the-select-element:the-option-element-37}](#the-option-element)
element in the [list of
options](#concept-select-option-list){#the-select-element:concept-select-option-list-30},
in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-select-element:tree-order-7
x-internal="tree-order"}, whose
[value](#concept-option-value){#the-select-element:concept-option-value-4}
is equal to the given new value, if any, must have its
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-24}
set to true and its
[dirtiness](#concept-option-dirtiness){#the-select-element:concept-option-dirtiness-6}
set to true.

This can result in no element having a
[selectedness](#concept-option-selectedness){#the-select-element:concept-option-selectedness-25}
set to true even in the case of the
[`select`{#the-select-element:the-select-element-42}](#the-select-element)
element having no
[`multiple`{#the-select-element:attr-select-multiple-16}](#attr-select-multiple)
attribute and a [display
size](#concept-select-size){#the-select-element:concept-select-size-7}
of 1.

The [`multiple`]{#dom-select-multiple .dfn dfn-for="HTMLSelectElement"
dfn-type="attribute"}, [`required`]{#dom-select-required .dfn
dfn-for="HTMLSelectElement" dfn-type="attribute"}, and
[`size`]{#dom-select-size .dfn dfn-for="HTMLSelectElement"
dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-select-element:reflect}
the respective content attributes of the same name. The
[`size`{#the-select-element:dom-select-size-2}](#dom-select-size) IDL
attribute has a [default
value](common-dom-interfaces.html#default-value){#the-select-element:default-value}
of 0.

For historical reasons, the default value of the
[`size`{#the-select-element:dom-select-size-3}](#dom-select-size) IDL
attribute does not return the actual size used, which, in the absence of
the [`size`{#the-select-element:attr-select-size-7}](#attr-select-size)
content attribute, is either 1 or 4 depending on the presence of the
[`multiple`{#the-select-element:attr-select-multiple-17}](#attr-select-multiple)
attribute.

The
[`willValidate`{#the-select-element:dom-cva-willvalidate-2}](form-control-infrastructure.html#dom-cva-willvalidate),
[`validity`{#the-select-element:dom-cva-validity-2}](form-control-infrastructure.html#dom-cva-validity),
and
[`validationMessage`{#the-select-element:dom-cva-validationmessage-2}](form-control-infrastructure.html#dom-cva-validationmessage)
IDL attributes, and the
[`checkValidity()`{#the-select-element:dom-cva-checkvalidity-2}](form-control-infrastructure.html#dom-cva-checkvalidity),
[`reportValidity()`{#the-select-element:dom-cva-reportvalidity-2}](form-control-infrastructure.html#dom-cva-reportvalidity),
and
[`setCustomValidity()`{#the-select-element:dom-cva-setcustomvalidity-2}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
methods, are part of the [constraint validation
API](form-control-infrastructure.html#the-constraint-validation-api){#the-select-element:the-constraint-validation-api}.
The
[`labels`{#the-select-element:dom-lfe-labels-2}](forms.html#dom-lfe-labels)
IDL attribute provides a list of the element\'s
[`label`{#the-select-element:the-label-element}](forms.html#the-label-element)s.
The
[`disabled`{#the-select-element:dom-fe-disabled-2}](form-control-infrastructure.html#dom-fe-disabled),
[`form`{#the-select-element:dom-fae-form-2}](form-control-infrastructure.html#dom-fae-form),
and
[`name`{#the-select-element:dom-fe-name-2}](form-control-infrastructure.html#dom-fe-name)
IDL attributes are part of the element\'s forms API.

::: example
The following example shows how a
[`select`{#the-select-element:the-select-element-43}](#the-select-element)
element can be used to offer the user with a set of options from which
the user can select a single option. The default option is preselected.

``` html
<p>
 <label for="unittype">Select unit type:</label>
 <select id="unittype" name="unittype">
  <option value="1"> Miner </option>
  <option value="2"> Puffer </option>
  <option value="3" selected> Snipey </option>
  <option value="4"> Max </option>
  <option value="5"> Firebot </option>
 </select>
</p>
```

When there is no default option, a placeholder can be used instead:

``` html
<select name="unittype" required>
 <option value=""> Select unit type </option>
 <option value="1"> Miner </option>
 <option value="2"> Puffer </option>
 <option value="3"> Snipey </option>
 <option value="4"> Max </option>
 <option value="5"> Firebot </option>
</select>
```
:::

::: example
Here, the user is offered a set of options from which they can select
any number. By default, all five options are selected.

``` html
<p>
 <label for="allowedunits">Select unit types to enable on this map:</label>
 <select id="allowedunits" name="allowedunits" multiple>
  <option value="1" selected> Miner </option>
  <option value="2" selected> Puffer </option>
  <option value="3" selected> Snipey </option>
  <option value="4" selected> Max </option>
  <option value="5" selected> Firebot </option>
 </select>
</p>
```
:::

::: example
Sometimes, a user has to select one or more items. This example shows
such an interface.

``` html
<label>
 Select the songs from that you would like on your Act II Mix Tape:
 <select multiple required name="act2">
  <option value="s1">It Sucks to Be Me (Reprise)
  <option value="s2">There is Life Outside Your Apartment
  <option value="s3">The More You Ruv Someone
  <option value="s4">Schadenfreude
  <option value="s5">I Wish I Could Go Back to College
  <option value="s6">The Money Song
  <option value="s7">School for Monsters
  <option value="s8">The Money Song (Reprise)
  <option value="s9">There's a Fine, Fine Line (Reprise)
  <option value="s10">What Do You Do With a B.A. in English? (Reprise)
  <option value="s11">For Now
 </select>
</label>
```
:::

::: example
Occasionally it can be useful to have a separator:

``` html
<label>
 Select the song to play next:
 <select required name="next">
  <option value="sr">Random
  <hr>
  <option value="s1">It Sucks to Be Me (Reprise)
  <option value="s2">There is Life Outside Your Apartment
  …
```
:::

#### [4.10.8]{.secno} The [`datalist`]{.dfn dfn-type="element"} element[](#the-datalist-element){.self-link}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[Element/datalist](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist "The <datalist> HTML element contains a set of <option> elements that represent the permissible or recommended options available to choose from within other controls.")

::: support
[Firefox[🔰 4+]{title="Partial implementation."}]{.firefox
.yes}[Safari12.1+]{.safari .yes}[Chrome20+]{.chrome .yes}

------------------------------------------------------------------------

[Opera9.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android[🔰
79+]{title="Partial implementation."}]{.firefox_android .yes}[Safari
iOS?]{.safari_ios .unknown}[Chrome Android33+]{.chrome_android
.yes}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDataListElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDataListElement "The HTMLDataListElement interface provides special properties (beyond the HTMLElement object interface it also has available to it by inheritance) to manipulate <datalist> elements and their content.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari12.1+]{.safari .yes}[Chrome20+]{.chrome
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

[Categories](dom.html#concept-element-categories){#the-datalist-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-datalist-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-datalist-element:phrasing-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-datalist-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-datalist-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-datalist-element:concept-element-content-model}:
:   Either: [phrasing
    content](dom.html#phrasing-content-2){#the-datalist-element:phrasing-content-2-3}.
:   Or: Zero or more
    [`option`{#the-datalist-element:the-option-element}](#the-option-element)
    and
    [script-supporting](dom.html#script-supporting-elements-2){#the-datalist-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-datalist-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-datalist-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-datalist-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-datalist-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-datalist).
:   [For implementers](https://w3c.github.io/html-aam/#el-datalist).

[DOM interface](dom.html#concept-element-dom){#the-datalist-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLDataListElement : HTMLElement {
      [HTMLConstructor] constructor();

      [SameObject] readonly attribute HTMLCollection options;
    };
    ```

The
[`datalist`{#the-datalist-element:the-datalist-element}](#the-datalist-element)
element represents a set of
[`option`{#the-datalist-element:the-option-element-2}](#the-option-element)
elements that represent predefined options for other controls. In the
rendering, the
[`datalist`{#the-datalist-element:the-datalist-element-2}](#the-datalist-element)
element
[represents](dom.html#represents){#the-datalist-element:represents}
nothing and it, along with its children, should be hidden.

The
[`datalist`{#the-datalist-element:the-datalist-element-3}](#the-datalist-element)
element can be used in two ways. In the simplest case, the
[`datalist`{#the-datalist-element:the-datalist-element-4}](#the-datalist-element)
element has just
[`option`{#the-datalist-element:the-option-element-3}](#the-option-element)
element children.

::: example
``` html
<label>
 Animal:
 <input name=animal list=animals>
 <datalist id=animals>
  <option value="Cat">
  <option value="Dog">
 </datalist>
</label>
```
:::

In the more elaborate case, the
[`datalist`{#the-datalist-element:the-datalist-element-5}](#the-datalist-element)
element can be given contents that are to be displayed for down-level
clients that don\'t support
[`datalist`{#the-datalist-element:the-datalist-element-6}](#the-datalist-element).
In this case, the
[`option`{#the-datalist-element:the-option-element-4}](#the-option-element)
elements are provided inside a
[`select`{#the-datalist-element:the-select-element}](#the-select-element)
element inside the
[`datalist`{#the-datalist-element:the-datalist-element-7}](#the-datalist-element)
element.

::: example
``` html
<label>
 Animal:
 <input name=animal list=animals>
</label>
<datalist id=animals>
 <label>
  or select from the list:
  <select name=animal>
   <option value="">
   <option>Cat
   <option>Dog
  </select>
 </label>
</datalist>
```
:::

The
[`datalist`{#the-datalist-element:the-datalist-element-8}](#the-datalist-element)
element is hooked up to an
[`input`{#the-datalist-element:the-input-element}](input.html#the-input-element)
element using the
[`list`{#the-datalist-element:attr-input-list}](input.html#attr-input-list)
attribute on the
[`input`{#the-datalist-element:the-input-element-2}](input.html#the-input-element)
element.

Each
[`option`{#the-datalist-element:the-option-element-5}](#the-option-element)
element that is a descendant of the
[`datalist`{#the-datalist-element:the-datalist-element-9}](#the-datalist-element)
element, that is not
[disabled](#concept-option-disabled){#the-datalist-element:concept-option-disabled},
and whose
[value](#concept-option-value){#the-datalist-element:concept-option-value}
is a string that isn\'t the empty string, represents a suggestion. Each
suggestion has a
[value](#concept-option-value){#the-datalist-element:concept-option-value-2}
and a
[label](#concept-option-label){#the-datalist-element:concept-option-label}.

`datalist`{.variable}`.`[`options`](#dom-datalist-options){#dom-datalist-options-dev}

:   Returns an
    [`HTMLCollection`{#the-datalist-element:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
    of the
    [`option`{#the-datalist-element:the-option-element-6}](#the-option-element)
    elements of the
    [`datalist`{#the-datalist-element:the-datalist-element-10}](#the-datalist-element)
    element.

The [`options`]{#dom-datalist-options .dfn dfn-for="HTMLDataListElement"
dfn-type="attribute"} IDL attribute must return an
[`HTMLCollection`{#the-datalist-element:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`datalist`{#the-datalist-element:the-datalist-element-11}](#the-datalist-element)
node, whose filter matches
[`option`{#the-datalist-element:the-option-element-7}](#the-option-element)
elements.

**Constraint validation**: If an element has a
[`datalist`{#the-datalist-element:the-datalist-element-12}](#the-datalist-element)
element ancestor, it is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#the-datalist-element:barred-from-constraint-validation}.

#### [4.10.9]{.secno} The [`optgroup`]{.dfn dfn-type="element"} element[](#the-optgroup-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/optgroup](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup "The <optgroup> HTML element creates a grouping of options within a <select> element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLOptGroupElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOptGroupElement "The HTMLOptGroupElement interface provides special properties and methods (beyond the regular HTMLElement object interface they also have available to them by inheritance) for manipulating the layout and presentation of <optgroup> elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

[Categories](dom.html#concept-element-categories){#the-optgroup-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-optgroup-element:concept-element-contexts}:
:   As a child of a
    [`select`{#the-optgroup-element:the-select-element}](#the-select-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-optgroup-element:concept-element-content-model}:
:   Zero or more
    [`option`{#the-optgroup-element:the-option-element}](#the-option-element)
    and
    [script-supporting](dom.html#script-supporting-elements-2){#the-optgroup-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-optgroup-element:concept-element-tag-omission}:
:   An
    [`optgroup`{#the-optgroup-element:the-optgroup-element}](#the-optgroup-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-optgroup-element:syntax-end-tag}
    can be omitted if the
    [`optgroup`{#the-optgroup-element:the-optgroup-element-2}](#the-optgroup-element)
    element is immediately followed by another
    [`optgroup`{#the-optgroup-element:the-optgroup-element-3}](#the-optgroup-element)
    element, if it is immediately followed by an
    [`hr`{#the-optgroup-element:the-hr-element}](grouping-content.html#the-hr-element)
    element, or if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-optgroup-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-optgroup-element:global-attributes}
:   [`disabled`{#the-optgroup-element:attr-optgroup-disabled}](#attr-optgroup-disabled)
    --- Whether the form control is disabled
:   [`label`{#the-optgroup-element:attr-optgroup-label}](#attr-optgroup-label)
    --- User-visible label

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-optgroup-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-optgroup).
:   [For implementers](https://w3c.github.io/html-aam/#el-optgroup).

[DOM interface](dom.html#concept-element-dom){#the-optgroup-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLOptGroupElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute boolean disabled;
      [CEReactions] attribute DOMString label;
    };
    ```

The
[`optgroup`{#the-optgroup-element:the-optgroup-element-4}](#the-optgroup-element)
element
[represents](dom.html#represents){#the-optgroup-element:represents} a
group of
[`option`{#the-optgroup-element:the-option-element-2}](#the-option-element)
elements with a common label.

The element\'s group of
[`option`{#the-optgroup-element:the-option-element-3}](#the-option-element)
elements consists of the
[`option`{#the-optgroup-element:the-option-element-4}](#the-option-element)
elements that are children of the
[`optgroup`{#the-optgroup-element:the-optgroup-element-5}](#the-optgroup-element)
element.

When showing
[`option`{#the-optgroup-element:the-option-element-5}](#the-option-element)
elements in
[`select`{#the-optgroup-element:the-select-element-2}](#the-select-element)
elements, user agents should show the
[`option`{#the-optgroup-element:the-option-element-6}](#the-option-element)
elements of such groups as being related to each other and separate from
other
[`option`{#the-optgroup-element:the-option-element-7}](#the-option-element)
elements.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/disabled](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled "The Boolean disabled attribute, when present, makes the element not mutable, focusable, or even submitted with the form. The user can neither edit nor focus on the control, nor its form control descendants.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`disabled`]{#attr-optgroup-disabled .dfn dfn-for="optgroup"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-optgroup-element:boolean-attribute}
and can be used to
[disable](#concept-option-disabled){#the-optgroup-element:concept-option-disabled}
a group of
[`option`{#the-optgroup-element:the-option-element-8}](#the-option-element)
elements together.

The [`label`]{#attr-optgroup-label .dfn dfn-for="optgroup"
dfn-type="element-attr"} attribute must be specified. Its value gives
the name of the group, for the purposes of the user interface. User
agents should use this attribute\'s value when labeling the group of
[`option`{#the-optgroup-element:the-option-element-9}](#the-option-element)
elements in a
[`select`{#the-optgroup-element:the-select-element-3}](#the-select-element)
element.

The [`disabled`]{#dom-optgroup-disabled .dfn
dfn-for="HTMLOptGroupElement" dfn-type="attribute"} and
[`label`]{#dom-optgroup-label .dfn dfn-for="HTMLOptGroupElement"
dfn-type="attribute"} attributes must
[reflect](common-dom-interfaces.html#reflect){#the-optgroup-element:reflect}
the respective content attributes of the same name.

There is no way to select an
[`optgroup`{#the-optgroup-element:the-optgroup-element-6}](#the-optgroup-element)
element. Only
[`option`{#the-optgroup-element:the-option-element-10}](#the-option-element)
elements can be selected. An
[`optgroup`{#the-optgroup-element:the-optgroup-element-7}](#the-optgroup-element)
element merely provides a label for a group of
[`option`{#the-optgroup-element:the-option-element-11}](#the-option-element)
elements.

::: example
The following snippet shows how a set of lessons from three courses
could be offered in a
[`select`{#the-optgroup-element:the-select-element-4}](#the-select-element)
drop-down widget:

``` html
<form action="courseselector.dll" method="get">
 <p>Which course would you like to watch today?
 <p><label>Course:
  <select name="c">
   <optgroup label="8.01 Physics I: Classical Mechanics">
    <option value="8.01.1">Lecture 01: Powers of Ten
    <option value="8.01.2">Lecture 02: 1D Kinematics
    <option value="8.01.3">Lecture 03: Vectors
   <optgroup label="8.02 Electricity and Magnetism">
    <option value="8.02.1">Lecture 01: What holds our world together?
    <option value="8.02.2">Lecture 02: Electric Field
    <option value="8.02.3">Lecture 03: Electric Flux
   <optgroup label="8.03 Physics III: Vibrations and Waves">
    <option value="8.03.1">Lecture 01: Periodic Phenomenon
    <option value="8.03.2">Lecture 02: Beats
    <option value="8.03.3">Lecture 03: Forced Oscillations with Damping
  </select>
 </label>
 <p><input type=submit value="▶ Play">
</form>
```
:::

#### [4.10.10]{.secno} The [`option`]{.dfn dfn-type="element"} element[](#the-option-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/option](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option "The <option> HTML element is used to define an item contained in a <select>, an <optgroup>, or a <datalist> element. As such, <option> can represent menu items in popups and other lists of items in an HTML document.")

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

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLOptionElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOptionElement "The HTMLOptionElement interface represents <option> elements and inherits all properties and methods of the HTMLElement interface.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.2+]{.safari .yes}[Chrome1+]{.chrome
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

[Categories](dom.html#concept-element-categories){#the-option-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-option-element:concept-element-contexts}:
:   As a child of a
    [`select`{#the-option-element:the-select-element}](#the-select-element)
    element.
:   As a child of a
    [`datalist`{#the-option-element:the-datalist-element}](#the-datalist-element)
    element.
:   As a child of an
    [`optgroup`{#the-option-element:the-optgroup-element}](#the-optgroup-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-option-element:concept-element-content-model}:
:   If the element has a
    [`label`{#the-option-element:attr-option-label}](#attr-option-label)
    attribute and a
    [`value`{#the-option-element:attr-option-value}](#attr-option-value)
    attribute:
    [Nothing](dom.html#concept-content-nothing){#the-option-element:concept-content-nothing}.
:   If the element has a
    [`label`{#the-option-element:attr-option-label-2}](#attr-option-label)
    attribute but no
    [`value`{#the-option-element:attr-option-value-2}](#attr-option-value)
    attribute:
    [Text](dom.html#text-content){#the-option-element:text-content}.
:   If the element has no
    [`label`{#the-option-element:attr-option-label-3}](#attr-option-label)
    attribute and is not a child of a
    [`datalist`{#the-option-element:the-datalist-element-2}](#the-datalist-element)
    element:
    [Text](dom.html#text-content){#the-option-element:text-content-2}
    that is not [inter-element
    whitespace](dom.html#inter-element-whitespace){#the-option-element:inter-element-whitespace}.
:   If the element has no
    [`label`{#the-option-element:attr-option-label-4}](#attr-option-label)
    attribute and is a child of a
    [`datalist`{#the-option-element:the-datalist-element-3}](#the-datalist-element)
    element:
    [Text](dom.html#text-content){#the-option-element:text-content-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-option-element:concept-element-tag-omission}:
:   An
    [`option`{#the-option-element:the-option-element}](#the-option-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-option-element:syntax-end-tag}
    can be omitted if the
    [`option`{#the-option-element:the-option-element-2}](#the-option-element)
    element is immediately followed by another
    [`option`{#the-option-element:the-option-element-3}](#the-option-element)
    element, if it is immediately followed by an
    [`optgroup`{#the-option-element:the-optgroup-element-2}](#the-optgroup-element)
    element, if it is immediately followed by an
    [`hr`{#the-option-element:the-hr-element}](grouping-content.html#the-hr-element)
    element, or if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-option-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-option-element:global-attributes}
:   [`disabled`{#the-option-element:attr-option-disabled}](#attr-option-disabled)
    --- Whether the form control is disabled
:   [`label`{#the-option-element:attr-option-label-5}](#attr-option-label)
    --- User-visible label
:   [`selected`{#the-option-element:attr-option-selected}](#attr-option-selected)
    --- Whether the option is selected by default
:   [`value`{#the-option-element:attr-option-value-3}](#attr-option-value)
    --- Value to be used for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-option-element:form-submission-2}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-option-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-option).
:   [For implementers](https://w3c.github.io/html-aam/#el-option).

[DOM interface](dom.html#concept-element-dom){#the-option-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window,
     LegacyFactoryFunction=Option(optional DOMString text = "", optional DOMString value, optional boolean defaultSelected = false, optional boolean selected = false)]
    interface HTMLOptionElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute boolean disabled;
      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute DOMString label;
      [CEReactions] attribute boolean defaultSelected;
      attribute boolean selected;
      [CEReactions] attribute DOMString value;

      [CEReactions] attribute DOMString text;
      readonly attribute long index;
    };
    ```

The
[`option`{#the-option-element:the-option-element-4}](#the-option-element)
element
[represents](dom.html#represents){#the-option-element:represents} an
option in a
[`select`{#the-option-element:the-select-element-2}](#the-select-element)
element or as part of a list of suggestions in a
[`datalist`{#the-option-element:the-datalist-element-4}](#the-datalist-element)
element.

In certain circumstances described in the definition of the
[`select`{#the-option-element:the-select-element-3}](#the-select-element)
element, an
[`option`{#the-option-element:the-option-element-5}](#the-option-element)
element can be a
[`select`{#the-option-element:the-select-element-4}](#the-select-element)
element\'s [placeholder label
option](#placeholder-label-option){#the-option-element:placeholder-label-option}.
A [placeholder label
option](#placeholder-label-option){#the-option-element:placeholder-label-option-2}
does not represent an actual option, but instead represents a label for
the
[`select`{#the-option-element:the-select-element-5}](#the-select-element)
control.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/disabled](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled "The Boolean disabled attribute, when present, makes the element not mutable, focusable, or even submitted with the form. The user can neither edit nor focus on the control, nor its form control descendants.")

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

The [`disabled`]{#attr-option-disabled .dfn dfn-for="option"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-option-element:boolean-attribute}.
An
[`option`{#the-option-element:the-option-element-6}](#the-option-element)
element is [disabled]{#concept-option-disabled .dfn} if its
[`disabled`{#the-option-element:attr-option-disabled-2}](#attr-option-disabled)
attribute is present or if it is a child of an
[`optgroup`{#the-option-element:the-optgroup-element-3}](#the-optgroup-element)
element whose
[`disabled`{#the-option-element:attr-optgroup-disabled}](#attr-optgroup-disabled)
attribute is present.

An
[`option`{#the-option-element:the-option-element-7}](#the-option-element)
element that is
[disabled](#attr-option-disabled){#the-option-element:attr-option-disabled-3}
must prevent any
[`click`{#the-option-element:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
events that are
[queued](webappapis.html#queue-a-task){#the-option-element:queue-a-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#the-option-element:user-interaction-task-source}
from being dispatched on the element.

The [`label`]{#attr-option-label .dfn dfn-for="option"
dfn-type="element-attr"} attribute provides a label for element. The
[label]{#concept-option-label .dfn} of an
[`option`{#the-option-element:the-option-element-8}](#the-option-element)
element is the value of the
[`label`{#the-option-element:attr-option-label-6}](#attr-option-label)
content attribute, if there is one and its value is not the empty
string, or, otherwise, the value of the element\'s
[`text`{#the-option-element:dom-option-text-2}](#dom-option-text) IDL
attribute.

The
[`label`{#the-option-element:attr-option-label-7}](#attr-option-label)
content attribute, if specified, must not be empty.

The [`value`]{#attr-option-value .dfn dfn-for="option"
dfn-type="element-attr"} attribute provides a value for element. The
[value]{#concept-option-value .dfn} of an
[`option`{#the-option-element:the-option-element-9}](#the-option-element)
element is the value of the
[`value`{#the-option-element:attr-option-value-4}](#attr-option-value)
content attribute, if there is one, or, if there is not, the value of
the element\'s
[`text`{#the-option-element:dom-option-text-3}](#dom-option-text) IDL
attribute.

The [`selected`]{#attr-option-selected .dfn dfn-for="option"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-option-element:boolean-attribute-2}.
It represents the default
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness}
of the element.

The [dirtiness]{#concept-option-dirtiness .dfn} of an
[`option`{#the-option-element:the-option-element-10}](#the-option-element)
element is a boolean state, initially false. It controls whether adding
or removing the
[`selected`{#the-option-element:attr-option-selected-2}](#attr-option-selected)
content attribute has any effect.

The [selectedness]{#concept-option-selectedness .dfn} of an
[`option`{#the-option-element:the-option-element-11}](#the-option-element)
element is a boolean state, initially false. Except where otherwise
specified, when the element is created, its
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-2}
must be set to true if the element has a
[`selected`{#the-option-element:attr-option-selected-3}](#attr-option-selected)
attribute. Whenever an
[`option`{#the-option-element:the-option-element-12}](#the-option-element)
element\'s
[`selected`{#the-option-element:attr-option-selected-4}](#attr-option-selected)
attribute is added, if its
[dirtiness](#concept-option-dirtiness){#the-option-element:concept-option-dirtiness}
is false, its
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-3}
must be set to true. Whenever an
[`option`{#the-option-element:the-option-element-13}](#the-option-element)
element\'s
[`selected`{#the-option-element:attr-option-selected-5}](#attr-option-selected)
attribute is *removed*, if its
[dirtiness](#concept-option-dirtiness){#the-option-element:concept-option-dirtiness-2}
is false, its
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-4}
must be set to false.

The [`Option()`{#the-option-element:dom-option-2}](#dom-option)
constructor, when called with three or fewer arguments, overrides the
initial state of the
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-5}
state to always be false even if the third argument is true (implying
that a
[`selected`{#the-option-element:attr-option-selected-6}](#attr-option-selected)
attribute is to be set). The fourth argument can be used to explicitly
set the initial
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-6}
state when using the constructor.

A
[`select`{#the-option-element:the-select-element-6}](#the-select-element)
element whose
[`multiple`{#the-option-element:attr-select-multiple}](#attr-select-multiple)
attribute is not specified must not have more than one descendant
[`option`{#the-option-element:the-option-element-14}](#the-option-element)
element with its
[`selected`{#the-option-element:attr-option-selected-7}](#attr-option-selected)
attribute set.

An
[`option`{#the-option-element:the-option-element-15}](#the-option-element)
element\'s [index]{#concept-option-index .dfn} is the number of
[`option`{#the-option-element:the-option-element-16}](#the-option-element)
elements that are in the same [list of
options](#concept-select-option-list){#the-option-element:concept-select-option-list}
but that come before it in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-option-element:tree-order
x-internal="tree-order"}. If the
[`option`{#the-option-element:the-option-element-17}](#the-option-element)
element is not in a [list of
options](#concept-select-option-list){#the-option-element:concept-select-option-list-2},
then the
[`option`{#the-option-element:the-option-element-18}](#the-option-element)
element\'s
[index](#concept-option-index){#the-option-element:concept-option-index}
is zero.

`option`{.variable}`.`[`selected`](#dom-option-selected){#dom-option-selected-dev}

Returns true if the element is selected, and false otherwise.

Can be set, to override the current state of the element.

`option`{.variable}`.`[`index`](#dom-option-index){#dom-option-index-dev}

Returns the index of the element in its
[`select`{#the-option-element:the-select-element-7}](#the-select-element)
element\'s
[`options`{#the-option-element:dom-select-options}](#dom-select-options)
list.

`option`{.variable}`.`[`form`](#dom-option-form){#dom-option-form-dev}

Returns the element\'s
[`form`{#the-option-element:the-form-element}](forms.html#the-form-element)
element, if any, or null otherwise.

`option`{.variable}`.`[`text`](#dom-option-text){#dom-option-text-dev}

Same as
[`textContent`{#the-option-element:textcontent}](https://dom.spec.whatwg.org/#dom-node-textcontent){x-internal="textcontent"},
except that spaces are collapsed and
[`script`{#the-option-element:the-script-element}](scripting.html#the-script-element)
elements are skipped.

`option`{.variable}` = new `[`Option`](#dom-option){#dom-option-dev}`([ ``text`{.variable}` [, ``value`{.variable}` [, ``defaultSelected`{.variable}` [, ``selected`{.variable}` ] ] ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLOptionElement/Option](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOptionElement/Option "The Option() constructor creates a new HTMLOptionElement.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.2+]{.safari .yes}[Chrome1+]{.chrome
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

Returns a new
[`option`{#the-option-element:the-option-element-19}](#the-option-element)
element.

The `text`{.variable} argument sets the contents of the element.

The `value`{.variable} argument sets the
[`value`{#the-option-element:attr-option-value-5}](#attr-option-value)
attribute.

The `defaultSelected`{.variable} argument sets the
[`selected`{#the-option-element:attr-option-selected-8}](#attr-option-selected)
attribute.

The `selected`{.variable} argument sets whether or not the element is
selected. If it is omitted, even if the `defaultSelected`{.variable}
argument is true, the element is not selected.

The [`disabled`]{#dom-option-disabled .dfn dfn-for="HTMLOptionElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-option-element:reflect}
the content attribute of the same name. The
[`defaultSelected`]{#dom-option-defaultselected .dfn
dfn-for="HTMLOptionElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-option-element:reflect-2}
the
[`selected`{#the-option-element:attr-option-selected-9}](#attr-option-selected)
content attribute.

The [`label`]{#dom-option-label .dfn dfn-for="HTMLOptionElement"
dfn-type="attribute"} IDL attribute, on getting, if there is a
[`label`{#the-option-element:attr-option-label-8}](#attr-option-label)
content attribute, must return that attribute\'s value; otherwise, it
must return the element\'s
[label](#concept-option-label){#the-option-element:concept-option-label}.
On setting, the element\'s
[`label`{#the-option-element:attr-option-label-9}](#attr-option-label)
content attribute must be set to the new value.

The [`value`]{#dom-option-value .dfn dfn-for="HTMLOptionElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
element\'s
[value](#concept-option-value){#the-option-element:concept-option-value}.
On setting, the element\'s
[`value`{#the-option-element:attr-option-value-6}](#attr-option-value)
content attribute must be set to the new value.

The [`selected`]{#dom-option-selected .dfn dfn-for="HTMLOptionElement"
dfn-type="attribute"} IDL attribute, on getting, must return true if the
element\'s
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-7}
is true, and false otherwise. On setting, it must set the element\'s
[selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-8}
to the new value, set its
[dirtiness](#concept-option-dirtiness){#the-option-element:concept-option-dirtiness-3}
to true, and then cause the element to [ask for a
reset](#ask-for-a-reset){#the-option-element:ask-for-a-reset}.

The [`index`]{#dom-option-index .dfn dfn-for="HTMLOptionElement"
dfn-type="attribute"} IDL attribute must return the element\'s
[index](#concept-option-index){#the-option-element:concept-option-index-2}.

The [`text`]{#dom-option-text .dfn dfn-for="HTMLOptionElement"
dfn-type="attribute"} IDL attribute, on getting, must return the result
of [stripping and collapsing ASCII
whitespace](https://infra.spec.whatwg.org/#strip-and-collapse-ascii-whitespace){#the-option-element:strip-and-collapse-ascii-whitespace
x-internal="strip-and-collapse-ascii-whitespace"} from the concatenation
of
[data](https://dom.spec.whatwg.org/#concept-cd-data){#the-option-element:concept-cd-data
x-internal="concept-cd-data"} of all the
[`Text`{#the-option-element:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
node descendants of the
[`option`{#the-option-element:the-option-element-20}](#the-option-element)
element, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-option-element:tree-order-2
x-internal="tree-order"}, excluding any that are descendants of
descendants of the
[`option`{#the-option-element:the-option-element-21}](#the-option-element)
element that are themselves
[`script`{#the-option-element:the-script-element-2}](scripting.html#the-script-element)
or [SVG
`script`](https://svgwg.org/svg2-draft/interact.html#ScriptElement){#the-option-element:svg-script
x-internal="svg-script"} elements.

The [`text`{#the-option-element:dom-option-text-4}](#dom-option-text)
attribute\'s setter must [string replace
all](https://dom.spec.whatwg.org/#string-replace-all){#the-option-element:string-replace-all
x-internal="string-replace-all"} with the given value within this
element.

The [`form`]{#dom-option-form .dfn dfn-for="HTMLOptionElement"
dfn-type="attribute"} IDL attribute\'s behavior depends on whether the
[`option`{#the-option-element:the-option-element-22}](#the-option-element)
element is in a
[`select`{#the-option-element:the-select-element-8}](#the-select-element)
element or not. If the
[`option`{#the-option-element:the-option-element-23}](#the-option-element)
has a
[`select`{#the-option-element:the-select-element-9}](#the-select-element)
element as its parent, or has an
[`optgroup`{#the-option-element:the-optgroup-element-4}](#the-optgroup-element)
element as its parent and that
[`optgroup`{#the-option-element:the-optgroup-element-5}](#the-optgroup-element)
element has a
[`select`{#the-option-element:the-select-element-10}](#the-select-element)
element as its parent, then the
[`form`{#the-option-element:dom-option-form-2}](#dom-option-form) IDL
attribute must return the same value as the
[`form`{#the-option-element:dom-fae-form}](form-control-infrastructure.html#dom-fae-form)
IDL attribute on that
[`select`{#the-option-element:the-select-element-11}](#the-select-element)
element. Otherwise, it must return null.

A legacy factory function is provided for creating
[`HTMLOptionElement`{#the-option-element:htmloptionelement}](#htmloptionelement)
objects (in addition to the factory methods from DOM such as
[`createElement()`{#the-option-element:dom-document-createelement}](https://dom.spec.whatwg.org/#dom-document-createelement){x-internal="dom-document-createelement"}):
[`Option(``text`{.variable}`, ``value`{.variable}`, ``defaultSelected`{.variable}`, ``selected`{.variable}`)`]{#dom-option
.dfn dfn-for="HTMLOptionElement" dfn-type="constructor"}. When invoked,
the legacy factory function must perform the following steps:

1.  Let `document`{.variable} be the [current global
    object](webappapis.html#current-global-object){#the-option-element:current-global-object}\'s
    [associated
    `Document`](nav-history-apis.html#concept-document-window){#the-option-element:concept-document-window}.

2.  Let `option`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-option-element:create-an-element
    x-internal="create-an-element"} given `document`{.variable},
    \"`option`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-option-element:html-namespace-2
    x-internal="html-namespace-2"}.

3.  If `text`{.variable} is not the empty string, then append to
    `option`{.variable} a new
    [`Text`{#the-option-element:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    node whose data is `text`{.variable}.

4.  If `value`{.variable} is given, then [set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-option-element:concept-element-attributes-set-value
    x-internal="concept-element-attributes-set-value"} for
    `option`{.variable} using
    \"[`value`{#the-option-element:attr-option-value-7}](#attr-option-value)\"
    and `value`{.variable}.

5.  If `defaultSelected`{.variable} is true, then [set an attribute
    value](https://dom.spec.whatwg.org/#concept-element-attributes-set-value){#the-option-element:concept-element-attributes-set-value-2
    x-internal="concept-element-attributes-set-value"} for
    `option`{.variable} using
    \"[`selected`{#the-option-element:attr-option-selected-10}](#attr-option-selected)\"
    and the empty string.

6.  If `selected`{.variable} is true, then set `option`{.variable}\'s
    [selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-9}
    to true; otherwise set its
    [selectedness](#concept-option-selectedness){#the-option-element:concept-option-selectedness-10}
    to false (even if `defaultSelected`{.variable} is true).

7.  Return `option`{.variable}.

#### [4.10.11]{.secno} The [`textarea`]{.dfn dfn-type="element"} element[](#the-textarea-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/textarea](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea "The <textarea> HTML element represents a multi-line plain-text editing control, useful when you want to allow users to enter a sizeable amount of free-form text, for example a comment on a review or feedback form.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTextAreaElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTextAreaElement "The HTMLTextAreaElement interface provides special properties and methods for manipulating the layout and presentation of <textarea> elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-textarea-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-textarea-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-textarea-element:phrasing-content-2}.
:   [Interactive
    content](dom.html#interactive-content-2){#the-textarea-element:interactive-content-2}.
:   [Listed](forms.html#category-listed){#the-textarea-element:category-listed},
    [labelable](forms.html#category-label){#the-textarea-element:category-label},
    [submittable](forms.html#category-submit){#the-textarea-element:category-submit},
    [resettable](forms.html#category-reset){#the-textarea-element:category-reset},
    and [autocapitalize-and-autocorrect
    inheriting](forms.html#category-autocapitalize){#the-textarea-element:category-autocapitalize}
    [form-associated
    element](forms.html#form-associated-element){#the-textarea-element:form-associated-element}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-textarea-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-textarea-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-textarea-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-textarea-element:concept-element-content-model}:
:   [Text](dom.html#text-content){#the-textarea-element:text-content}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-textarea-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-textarea-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-textarea-element:global-attributes}
:   [`autocomplete`{#the-textarea-element:attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete)
    --- Hint for form autofill feature
:   [`cols`{#the-textarea-element:attr-textarea-cols}](#attr-textarea-cols)
    --- Maximum number of characters per line
:   [`dirname`{#the-textarea-element:attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname)
    --- Name of form control to use for sending the element\'s
    [directionality](dom.html#the-directionality){#the-textarea-element:the-directionality}
    in [form
    submission](form-control-infrastructure.html#form-submission-2){#the-textarea-element:form-submission-2}
:   [`disabled`{#the-textarea-element:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
    --- Whether the form control is disabled
:   [`form`{#the-textarea-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    --- Associates the element with a
    [`form`{#the-textarea-element:the-form-element}](forms.html#the-form-element)
    element
:   [`maxlength`{#the-textarea-element:attr-textarea-maxlength}](#attr-textarea-maxlength)
    --- Maximum
    [length](https://infra.spec.whatwg.org/#string-length){#the-textarea-element:length
    x-internal="length"} of value
:   [`minlength`{#the-textarea-element:attr-textarea-minlength}](#attr-textarea-minlength)
    --- Minimum
    [length](https://infra.spec.whatwg.org/#string-length){#the-textarea-element:length-2
    x-internal="length"} of value
:   [`name`{#the-textarea-element:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    --- Name of the element to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-textarea-element:form-submission-2-2}
    and in the
    [`form.elements`{#the-textarea-element:dom-form-elements}](forms.html#dom-form-elements)
    API
:   [`placeholder`{#the-textarea-element:attr-textarea-placeholder}](#attr-textarea-placeholder)
    --- User-visible label to be placed within the form control
:   [`readonly`{#the-textarea-element:attr-textarea-readonly}](#attr-textarea-readonly)
    --- Whether to allow the value to be edited by the user
:   [`required`{#the-textarea-element:attr-textarea-required}](#attr-textarea-required)
    --- Whether the control is required for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-textarea-element:form-submission-2-3}
:   [`rows`{#the-textarea-element:attr-textarea-rows}](#attr-textarea-rows)
    --- Number of lines to show
:   [`wrap`{#the-textarea-element:attr-textarea-wrap}](#attr-textarea-wrap)
    --- How the value of the form control is to be wrapped for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-textarea-element:form-submission-2-4}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-textarea-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-textarea).
:   [For implementers](https://w3c.github.io/html-aam/#el-textarea).

[DOM interface](dom.html#concept-element-dom){#the-textarea-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTextAreaElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString autocomplete;
      [CEReactions] attribute unsigned long cols;
      [CEReactions] attribute DOMString dirName;
      [CEReactions] attribute boolean disabled;
      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute long maxLength;
      [CEReactions] attribute long minLength;
      [CEReactions] attribute DOMString name;
      [CEReactions] attribute DOMString placeholder;
      [CEReactions] attribute boolean readOnly;
      [CEReactions] attribute boolean required;
      [CEReactions] attribute unsigned long rows;
      [CEReactions] attribute DOMString wrap;

      readonly attribute DOMString type;
      [CEReactions] attribute DOMString defaultValue;
      attribute [LegacyNullToEmptyString] DOMString value;
      readonly attribute unsigned long textLength;

      readonly attribute boolean willValidate;
      readonly attribute ValidityState validity;
      readonly attribute DOMString validationMessage;
      boolean checkValidity();
      boolean reportValidity();
      undefined setCustomValidity(DOMString error);

      readonly attribute NodeList labels;

      undefined select();
      attribute unsigned long selectionStart;
      attribute unsigned long selectionEnd;
      attribute DOMString selectionDirection;
      undefined setRangeText(DOMString replacement);
      undefined setRangeText(DOMString replacement, unsigned long start, unsigned long end, optional SelectionMode selectionMode = "preserve");
      undefined setSelectionRange(unsigned long start, unsigned long end, optional DOMString direction);
    };
    ```

The
[`textarea`{#the-textarea-element:the-textarea-element}](#the-textarea-element)
element
[represents](dom.html#represents){#the-textarea-element:represents} a
multiline plain text edit control for the element\'s [raw
value]{#concept-textarea-raw-value .dfn}. The contents of the control
represent the control\'s default value.

The [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value}
of a
[`textarea`{#the-textarea-element:the-textarea-element-2}](#the-textarea-element)
control must be initially the empty string.

This element [has rendering requirements involving the bidirectional
algorithm](dom.html#bidireq).

The [`readonly`]{#attr-textarea-readonly .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-textarea-element:boolean-attribute}
used to control whether the text can be edited by the user or not.

::: example
In this example, a text control is marked read-only because it
represents a read-only file:

``` html
Filename: <code>/etc/bash.bashrc</code>
<textarea name="buffer" readonly>
# System-wide .bashrc file for interactive bash(1) shells.

# To enable the settings / commands in this file for login shells as well,
# this file has to be sourced in /etc/profile.

# If not running interactively, don't do anything
[ -z "$PS1" ] &amp;&amp; return

...</textarea>
```
:::

**Constraint validation**: If the
[`readonly`{#the-textarea-element:attr-textarea-readonly-2}](#attr-textarea-readonly)
attribute is specified on a
[`textarea`{#the-textarea-element:the-textarea-element-3}](#the-textarea-element)
element, the element is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#the-textarea-element:barred-from-constraint-validation}.

A
[`textarea`{#the-textarea-element:the-textarea-element-4}](#the-textarea-element)
element is
[mutable](form-control-infrastructure.html#concept-fe-mutable){#the-textarea-element:concept-fe-mutable}
if it is neither
[disabled](form-control-infrastructure.html#concept-fe-disabled){#the-textarea-element:concept-fe-disabled}
nor has a
[`readonly`{#the-textarea-element:attr-textarea-readonly-3}](#attr-textarea-readonly)
attribute specified.

When a
[`textarea`{#the-textarea-element:the-textarea-element-5}](#the-textarea-element)
is
[mutable](form-control-infrastructure.html#concept-fe-mutable){#the-textarea-element:concept-fe-mutable-2},
its [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-2}
should be editable by the user: the user agent should allow the user to
edit, insert, and remove text, and to insert and remove line breaks in
the form of U+000A LINE FEED (LF) characters. Any time the user causes
the element\'s [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-3}
to change, the user agent must [queue an element
task](webappapis.html#queue-an-element-task){#the-textarea-element:queue-an-element-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#the-textarea-element:user-interaction-task-source}
given the
[`textarea`{#the-textarea-element:the-textarea-element-6}](#the-textarea-element)
element to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#the-textarea-element:concept-event-fire
x-internal="concept-event-fire"} named
[`input`{#the-textarea-element:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
at the
[`textarea`{#the-textarea-element:the-textarea-element-7}](#the-textarea-element)
element, with the
[`bubbles`{#the-textarea-element:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
and
[`composed`{#the-textarea-element:dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
attributes initialized to true. User agents may wait for a suitable
break in the user\'s interaction before queuing the task; for example, a
user agent could wait for the user to have not hit a key for 100ms, so
as to only fire the event when the user pauses, instead of continuously
for each keystroke.

A
[`textarea`{#the-textarea-element:the-textarea-element-8}](#the-textarea-element)
element\'s [dirty value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-textarea-element:concept-fe-dirty}
must be set to true whenever the user interacts with the control in a
way that changes the [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-4}.

The [cloning
steps](https://dom.spec.whatwg.org/#concept-node-clone-ext){#the-textarea-element:concept-node-clone-ext
x-internal="concept-node-clone-ext"} for
[`textarea`{#the-textarea-element:the-textarea-element-9}](#the-textarea-element)
elements given `node`{.variable}, `copy`{.variable}, and
`subtree`{.variable} are to propagate the [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-5}
and [dirty value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-textarea-element:concept-fe-dirty-2}
from `node`{.variable} to `copy`{.variable}.

The [children changed
steps](https://dom.spec.whatwg.org/#concept-node-children-changed-ext){#the-textarea-element:children-changed-steps
x-internal="children-changed-steps"} for
[`textarea`{#the-textarea-element:the-textarea-element-10}](#the-textarea-element)
elements must, if the element\'s [dirty value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-textarea-element:concept-fe-dirty-3}
is false, set the element\'s [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-6}
to its [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-textarea-element:child-text-content
x-internal="child-text-content"}.

The [reset
algorithm](form-control-infrastructure.html#concept-form-reset-control){#the-textarea-element:concept-form-reset-control}
for
[`textarea`{#the-textarea-element:the-textarea-element-11}](#the-textarea-element)
elements is to set the [user
validity](form-control-infrastructure.html#user-validity){#the-textarea-element:user-validity}
to false, the [dirty value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-textarea-element:concept-fe-dirty-4}
back to false, and the [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-7}
to its [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-textarea-element:child-text-content-2
x-internal="child-text-content"}.

When a
[`textarea`{#the-textarea-element:the-textarea-element-12}](#the-textarea-element)
element is popped off the [stack of open
elements](parsing.html#stack-of-open-elements){#the-textarea-element:stack-of-open-elements}
of an [HTML
parser](parsing.html#html-parser){#the-textarea-element:html-parser} or
[XML parser](xhtml.html#xml-parser){#the-textarea-element:xml-parser},
then the user agent must invoke the element\'s [reset
algorithm](form-control-infrastructure.html#concept-form-reset-control){#the-textarea-element:concept-form-reset-control-2}.

If the element is
[mutable](form-control-infrastructure.html#concept-fe-mutable){#the-textarea-element:concept-fe-mutable-3},
the user agent should allow the user to change the writing direction of
the element, setting it either to a left-to-right writing direction or a
right-to-left writing direction. If the user does so, the user agent
must then run the following steps:

1.  Set the element\'s
    [`dir`{#the-textarea-element:attr-dir}](dom.html#attr-dir) attribute
    to
    \"[`ltr`{#the-textarea-element:attr-dir-ltr}](dom.html#attr-dir-ltr)\"
    if the user selected a left-to-right writing direction, and
    \"[`rtl`{#the-textarea-element:attr-dir-rtl}](dom.html#attr-dir-rtl)\"
    if the user selected a right-to-left writing direction.

2.  [Queue an element
    task](webappapis.html#queue-an-element-task){#the-textarea-element:queue-an-element-task-2}
    on the [user interaction task
    source](webappapis.html#user-interaction-task-source){#the-textarea-element:user-interaction-task-source-2}
    given the
    [`textarea`{#the-textarea-element:the-textarea-element-13}](#the-textarea-element)
    element to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-textarea-element:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`input`{#the-textarea-element:event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    at the
    [`textarea`{#the-textarea-element:the-textarea-element-14}](#the-textarea-element)
    element, with the
    [`bubbles`{#the-textarea-element:dom-event-bubbles-2}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    and
    [`composed`{#the-textarea-element:dom-event-composed-2}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
    attributes initialized to true.

The [`cols`]{#attr-textarea-cols .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute specifies the expected maximum number
of characters per line. If the
[`cols`{#the-textarea-element:attr-textarea-cols-2}](#attr-textarea-cols)
attribute is specified, its value must be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#the-textarea-element:valid-non-negative-integer}
greater than zero. If applying the [rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-textarea-element:rules-for-parsing-non-negative-integers}
to the attribute\'s value results in a number greater than zero, then
the element\'s [character width]{#attr-textarea-cols-value .dfn} is that
value; otherwise, it is 20.

The user agent may use the
[`textarea`{#the-textarea-element:the-textarea-element-15}](#the-textarea-element)
element\'s [character
width](#attr-textarea-cols-value){#the-textarea-element:attr-textarea-cols-value}
as a hint to the user as to how many characters the server prefers per
line (e.g. for visual user agents by making the width of the control be
that many characters). In visual renderings, the user agent should wrap
the user\'s input in the rendering so that each line is no wider than
this number of characters.

The [`rows`]{#attr-textarea-rows .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute specifies the number of lines to
show. If the
[`rows`{#the-textarea-element:attr-textarea-rows-2}](#attr-textarea-rows)
attribute is specified, its value must be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#the-textarea-element:valid-non-negative-integer-2}
greater than zero. If applying the [rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-textarea-element:rules-for-parsing-non-negative-integers-2}
to the attribute\'s value results in a number greater than zero, then
the element\'s [character height]{#attr-textarea-rows-value .dfn} is
that value; otherwise, it is 2.

Visual user agents should set the height of the control to the number of
lines given by [character
height](#attr-textarea-rows-value){#the-textarea-element:attr-textarea-rows-value}.

The [`wrap`]{#attr-textarea-wrap .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-textarea-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`soft`]{#attr-textarea-wrap-soft .dfn dfn-for="textarea/wrap"
dfn-type="attr-value"}

[Soft]{#attr-textarea-wrap-soft-state .dfn}

Text is not to be wrapped when submitted (though can still be wrapped in
the rendering).

[`hard`]{#attr-textarea-wrap-hard .dfn dfn-for="textarea/wrap"
dfn-type="attr-value"}

[Hard]{#attr-textarea-wrap-hard-state .dfn}

Text is to have newlines added by the user agent so that the text is
wrapped when it is submitted.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[Soft](#attr-textarea-wrap-soft-state){#the-textarea-element:attr-textarea-wrap-soft-state}
state.

If the element\'s
[`wrap`{#the-textarea-element:attr-textarea-wrap-2}](#attr-textarea-wrap)
attribute is in the
[Hard](#attr-textarea-wrap-hard-state){#the-textarea-element:attr-textarea-wrap-hard-state}
state, the
[`cols`{#the-textarea-element:attr-textarea-cols-3}](#attr-textarea-cols)
attribute must be specified.

For historical reasons, the element\'s value is normalized in three
different ways for three different purposes. The [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-8}
is the value as it was originally set. It is not normalized. The [API
value](form-control-infrastructure.html#concept-fe-api-value){#the-textarea-element:concept-fe-api-value}
is the value used in the
[`value`{#the-textarea-element:dom-textarea-value-2}](#dom-textarea-value)
IDL attribute,
[`textLength`{#the-textarea-element:dom-textarea-textlength-2}](#dom-textarea-textlength)
IDL attribute, and by the
[`maxlength`{#the-textarea-element:attr-fe-maxlength}](form-control-infrastructure.html#attr-fe-maxlength)
and
[`minlength`{#the-textarea-element:attr-fe-minlength}](form-control-infrastructure.html#attr-fe-minlength)
content attributes. It is normalized so that line breaks use U+000A LINE
FEED (LF) characters. Finally, there is the
[value](form-control-infrastructure.html#concept-fe-value){#the-textarea-element:concept-fe-value},
as used in form submission and other processing models in this
specification. It is normalized as for the [API
value](form-control-infrastructure.html#concept-fe-api-value){#the-textarea-element:concept-fe-api-value-2},
and in addition, if necessary given the element\'s
[`wrap`{#the-textarea-element:attr-textarea-wrap-3}](#attr-textarea-wrap)
attribute, additional line breaks are inserted to wrap the text at the
given width.

The algorithm for obtaining the element\'s [API
value](form-control-infrastructure.html#concept-fe-api-value){#the-textarea-element:concept-fe-api-value-3}
is to return the element\'s [raw
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-9},
with [newlines
normalized](https://infra.spec.whatwg.org/#normalize-newlines){#the-textarea-element:normalize-newlines
x-internal="normalize-newlines"}.

The element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-textarea-element:concept-fe-value-2}
is defined to be the element\'s [API
value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-10}
with the [textarea wrapping
transformation](#textarea-wrapping-transformation){#the-textarea-element:textarea-wrapping-transformation}
applied. The [textarea wrapping
transformation]{#textarea-wrapping-transformation .dfn} is the following
algorithm, as applied to a string:

1.  If the element\'s
    [`wrap`{#the-textarea-element:attr-textarea-wrap-4}](#attr-textarea-wrap)
    attribute is in the
    [Hard](#attr-textarea-wrap-hard-state){#the-textarea-element:attr-textarea-wrap-hard-state-2}
    state, insert U+000A LINE FEED (LF) characters into the string using
    an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-textarea-element:implementation-defined
    x-internal="implementation-defined"} algorithm so that each line has
    no more than [character
    width](#attr-textarea-cols-value){#the-textarea-element:attr-textarea-cols-value-2}
    characters. For the purposes of this requirement, lines are
    delimited by the start of the string, the end of the string, and
    U+000A LINE FEED (LF) characters.

The [`maxlength`]{#attr-textarea-maxlength .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute is a [form control `maxlength`
attribute](form-control-infrastructure.html#attr-fe-maxlength){#the-textarea-element:attr-fe-maxlength-2}.

If the
[`textarea`{#the-textarea-element:the-textarea-element-16}](#the-textarea-element)
element has a [maximum allowed value
length](form-control-infrastructure.html#maximum-allowed-value-length){#the-textarea-element:maximum-allowed-value-length},
then the element\'s children must be such that the
[length](https://infra.spec.whatwg.org/#string-length){#the-textarea-element:length-3
x-internal="length"} of the value of the element\'s [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#the-textarea-element:descendant-text-content
x-internal="descendant-text-content"} with [newlines
normalized](https://infra.spec.whatwg.org/#normalize-newlines){#the-textarea-element:normalize-newlines-2
x-internal="normalize-newlines"} is less than or equal to the element\'s
[maximum allowed value
length](form-control-infrastructure.html#maximum-allowed-value-length){#the-textarea-element:maximum-allowed-value-length-2}.

The [`minlength`]{#attr-textarea-minlength .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute is a [form control `minlength`
attribute](form-control-infrastructure.html#attr-fe-minlength){#the-textarea-element:attr-fe-minlength-2}.

The [`required`]{#attr-textarea-required .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-textarea-element:boolean-attribute-2}.
When specified, the user will be required to enter a value before
submitting the form.

**Constraint validation**: If the element has its
[`required`{#the-textarea-element:attr-textarea-required-2}](#attr-textarea-required)
attribute specified, and the element is
[mutable](form-control-infrastructure.html#concept-fe-mutable){#the-textarea-element:concept-fe-mutable-4},
and the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-textarea-element:concept-fe-value-3}
is the empty string, then the element is [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#the-textarea-element:suffering-from-being-missing}.

The [`placeholder`]{#attr-textarea-placeholder .dfn dfn-for="textarea"
dfn-type="element-attr"} attribute represents a *short* hint (a word or
short phrase) intended to aid the user with data entry when the control
has no value. A hint could be a sample value or a brief description of
the expected format.

The
[`placeholder`{#the-textarea-element:attr-textarea-placeholder-2}](#attr-textarea-placeholder)
attribute should not be used as an alternative to a
[`label`{#the-textarea-element:the-label-element}](forms.html#the-label-element).
For a longer hint or other advisory text, the
[`title`{#the-textarea-element:attr-title}](dom.html#attr-title)
attribute is more appropriate.

These mechanisms are very similar but subtly different: the hint given
by the control\'s
[`label`{#the-textarea-element:the-label-element-2}](forms.html#the-label-element)
is shown at all times; the short hint given in the
[`placeholder`{#the-textarea-element:attr-textarea-placeholder-3}](#attr-textarea-placeholder)
attribute is shown before the user enters a value; and the hint in the
[`title`{#the-textarea-element:attr-title-2}](dom.html#attr-title)
attribute is shown when the user requests further help.

User agents should present this hint to the user when the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-textarea-element:concept-fe-value-4}
is the empty string and the control is not
[focused](interaction.html#focused){#the-textarea-element:focused} (e.g.
by displaying it inside a blank unfocused control). All U+000D CARRIAGE
RETURN U+000A LINE FEED character pairs (CRLF) in the hint, as well as
all other U+000D CARRIAGE RETURN (CR) and U+000A LINE FEED (LF)
characters in the hint, must be treated as line breaks when rendering
the hint.

If a user agent normally doesn\'t show this hint to the user when the
control is
[focused](interaction.html#focused){#the-textarea-element:focused-2},
then the user agent should nonetheless show the hint for the control if
it was focused as a result of the
[`autofocus`{#the-textarea-element:attr-fe-autofocus}](interaction.html#attr-fe-autofocus)
attribute, since in that case the user will not have had an opportunity
to examine the control before focusing it.

The
[`name`{#the-textarea-element:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attribute represents the element\'s name. The
[`dirname`{#the-textarea-element:attr-fe-dirname-2}](form-control-infrastructure.html#attr-fe-dirname)
attribute controls how the element\'s
[directionality](dom.html#the-directionality){#the-textarea-element:the-directionality-2}
is submitted. The
[`disabled`{#the-textarea-element:attr-fe-disabled-2}](form-control-infrastructure.html#attr-fe-disabled)
attribute is used to make the control non-interactive and to prevent its
value from being submitted. The
[`form`{#the-textarea-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the
[`textarea`{#the-textarea-element:the-textarea-element-17}](#the-textarea-element)
element with its [form
owner](form-control-infrastructure.html#form-owner){#the-textarea-element:form-owner}.
The
[`autocomplete`{#the-textarea-element:attr-fe-autocomplete-2}](form-control-infrastructure.html#attr-fe-autocomplete)
attribute controls how the user agent provides autofill behavior.

`textarea`{.variable}`.`[`type`](#dom-textarea-type){#dom-textarea-type-dev}
:   Returns the string \"`textarea`\".

`textarea`{.variable}`.`[`value`](#dom-textarea-value){#dom-textarea-value-dev}

:   Returns the current value of the element.

    Can be set, to change the value.

The [`cols`]{#dom-textarea-cols .dfn dfn-for="HTMLTextAreaElement"
dfn-type="attribute"}, [`placeholder`]{#dom-textarea-placeholder .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"},
[`required`]{#dom-textarea-required .dfn dfn-for="HTMLTextAreaElement"
dfn-type="attribute"}, [`rows`]{#dom-textarea-rows .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"}, and
[`wrap`]{#dom-textarea-wrap .dfn dfn-for="HTMLTextAreaElement"
dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-textarea-element:reflect}
the respective content attributes of the same name. The
[`cols`{#the-textarea-element:dom-textarea-cols-2}](#dom-textarea-cols)
and
[`rows`{#the-textarea-element:dom-textarea-rows-2}](#dom-textarea-rows)
attributes are [limited to only positive numbers with
fallback](common-dom-interfaces.html#limited-to-only-non-negative-numbers-greater-than-zero-with-fallback){#the-textarea-element:limited-to-only-non-negative-numbers-greater-than-zero-with-fallback}.
The
[`cols`{#the-textarea-element:dom-textarea-cols-3}](#dom-textarea-cols)
IDL attribute\'s [default
value](common-dom-interfaces.html#default-value){#the-textarea-element:default-value}
is 20. The
[`rows`{#the-textarea-element:dom-textarea-rows-3}](#dom-textarea-rows)
IDL attribute\'s [default
value](common-dom-interfaces.html#default-value){#the-textarea-element:default-value-2}
is 2. The [`dirName`]{#dom-textarea-dirname .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-textarea-element:reflect-2}
the
[`dirname`{#the-textarea-element:attr-fe-dirname-3}](form-control-infrastructure.html#attr-fe-dirname)
content attribute. The [`maxLength`]{#dom-textarea-maxlength .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-textarea-element:reflect-3}
the
[`maxlength`{#the-textarea-element:attr-textarea-maxlength-2}](#attr-textarea-maxlength)
content attribute, [limited to only non-negative
numbers](common-dom-interfaces.html#limited-to-only-non-negative-numbers){#the-textarea-element:limited-to-only-non-negative-numbers}.
The [`minLength`]{#dom-textarea-minlength .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-textarea-element:reflect-4}
the
[`minlength`{#the-textarea-element:attr-textarea-minlength-2}](#attr-textarea-minlength)
content attribute, [limited to only non-negative
numbers](common-dom-interfaces.html#limited-to-only-non-negative-numbers){#the-textarea-element:limited-to-only-non-negative-numbers-2}.
The [`readOnly`]{#dom-textarea-readonly .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-textarea-element:reflect-5}
the
[`readonly`{#the-textarea-element:attr-textarea-readonly-4}](#attr-textarea-readonly)
content attribute.

The [`type`]{#dom-textarea-type .dfn dfn-for="HTMLTextAreaElement"
dfn-type="attribute"} IDL attribute must return the value
\"`textarea`\".

The [`defaultValue`]{#dom-textarea-defaultvalue .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"} attribute\'s getter
must return the element\'s [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-textarea-element:child-text-content-3
x-internal="child-text-content"}.

The
[`defaultValue`{#the-textarea-element:dom-textarea-defaultvalue-2}](#dom-textarea-defaultvalue)
attribute\'s setter must [string replace
all](https://dom.spec.whatwg.org/#string-replace-all){#the-textarea-element:string-replace-all
x-internal="string-replace-all"} with the given value within this
element.

The [`value`]{#dom-textarea-value .dfn dfn-for="HTMLTextAreaElement"
dfn-type="attribute"} IDL attribute must, on getting, return the
element\'s [API
value](form-control-infrastructure.html#concept-fe-api-value){#the-textarea-element:concept-fe-api-value-4}.
On setting, it must perform the following steps:

1.  Let `oldAPIValue`{.variable} be this element\'s [API
    value](form-control-infrastructure.html#concept-fe-api-value){#the-textarea-element:concept-fe-api-value-5}.

2.  Set this element\'s [raw
    value](#concept-textarea-raw-value){#the-textarea-element:concept-textarea-raw-value-11}
    to the new value.

3.  Set this element\'s [dirty value
    flag](form-control-infrastructure.html#concept-fe-dirty){#the-textarea-element:concept-fe-dirty-5}
    to true.

4.  If the new [API
    value](form-control-infrastructure.html#concept-fe-api-value){#the-textarea-element:concept-fe-api-value-6}
    is different from `oldAPIValue`{.variable}, then move the [text
    entry cursor
    position](form-control-infrastructure.html#concept-textarea/input-cursor){#the-textarea-element:concept-textarea/input-cursor}
    to the end of the text control, unselecting any selected text and
    [resetting the selection
    direction](form-control-infrastructure.html#set-the-selection-direction){#the-textarea-element:set-the-selection-direction}
    to \"`none`\".

The [`textLength`]{#dom-textarea-textlength .dfn
dfn-for="HTMLTextAreaElement" dfn-type="attribute"} IDL attribute must
return the
[length](https://infra.spec.whatwg.org/#string-length){#the-textarea-element:length-4
x-internal="length"} of the element\'s [API
value](form-control-infrastructure.html#concept-fe-api-value){#the-textarea-element:concept-fe-api-value-7}.

The
[`willValidate`{#the-textarea-element:dom-cva-willvalidate-2}](form-control-infrastructure.html#dom-cva-willvalidate),
[`validity`{#the-textarea-element:dom-cva-validity-2}](form-control-infrastructure.html#dom-cva-validity),
and
[`validationMessage`{#the-textarea-element:dom-cva-validationmessage-2}](form-control-infrastructure.html#dom-cva-validationmessage)
IDL attributes, and the
[`checkValidity()`{#the-textarea-element:dom-cva-checkvalidity-2}](form-control-infrastructure.html#dom-cva-checkvalidity),
[`reportValidity()`{#the-textarea-element:dom-cva-reportvalidity-2}](form-control-infrastructure.html#dom-cva-reportvalidity),
and
[`setCustomValidity()`{#the-textarea-element:dom-cva-setcustomvalidity-2}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
methods, are part of the [constraint validation
API](form-control-infrastructure.html#the-constraint-validation-api){#the-textarea-element:the-constraint-validation-api}.
The
[`labels`{#the-textarea-element:dom-lfe-labels-2}](forms.html#dom-lfe-labels)
IDL attribute provides a list of the element\'s
[`label`{#the-textarea-element:the-label-element-3}](forms.html#the-label-element)s.
The
[`select()`{#the-textarea-element:dom-textarea/input-select-2}](form-control-infrastructure.html#dom-textarea/input-select),
[`selectionStart`{#the-textarea-element:dom-textarea/input-selectionstart-2}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#the-textarea-element:dom-textarea/input-selectionend-2}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#the-textarea-element:dom-textarea/input-selectiondirection-2}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`setRangeText()`{#the-textarea-element:dom-textarea/input-setrangetext-3}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#the-textarea-element:dom-textarea/input-setselectionrange-2}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods and IDL attributes expose the element\'s text selection. The
[`disabled`{#the-textarea-element:dom-fe-disabled-2}](form-control-infrastructure.html#dom-fe-disabled),
[`form`{#the-textarea-element:dom-fae-form-2}](form-control-infrastructure.html#dom-fae-form),
and
[`name`{#the-textarea-element:dom-fe-name-2}](form-control-infrastructure.html#dom-fe-name)
IDL attributes are part of the element\'s forms API.

::: example
Here is an example of a
[`textarea`{#the-textarea-element:the-textarea-element-18}](#the-textarea-element)
being used for unrestricted free-form text input in a form:

``` html
<p>If you have any comments, please let us know: <textarea cols=80 name=comments></textarea></p>
```

To specify a maximum length for the comments, one can use the
[`maxlength`{#the-textarea-element:attr-textarea-maxlength-3}](#attr-textarea-maxlength)
attribute:

``` html
<p>If you have any short comments, please let us know: <textarea cols=80 name=comments maxlength=200></textarea></p>
```

To give a default value, text can be included inside the element:

``` html
<p>If you have any comments, please let us know: <textarea cols=80 name=comments>You rock!</textarea></p>
```

You can also give a minimum length. Here, a letter needs to be filled
out by the user; a template (which is shorter than the minimum length)
is provided, but is insufficient to submit the form:

``` html
<textarea required minlength="500">Dear Madam Speaker,

Regarding your letter dated ...

...

Yours Sincerely,

...</textarea>
```

A placeholder can be given as well, to suggest the basic form to the
user, without providing an explicit template:

``` html
<textarea placeholder="Dear Francine,

They closed the parks this week, so we won't be able to
meet your there. Should we just have dinner?

Love,
Daddy"></textarea>
```

To have the browser submit [the
directionality](dom.html#the-directionality){#the-textarea-element:the-directionality-3}
of the element along with the value, the
[`dirname`{#the-textarea-element:attr-fe-dirname-4}](form-control-infrastructure.html#attr-fe-dirname)
attribute can be specified:

``` html
<p>If you have any comments, please let us know (you may use either English or Hebrew for your comments):
<textarea cols=80 name=comments dirname=comments.dir></textarea></p>
```
:::

#### [4.10.12]{.secno} The [`output`]{.dfn dfn-type="element"} element[](#the-output-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/output](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output "The <output> HTML element is a container element into which a site or app can inject the results of a calculation or the outcome of a user action.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLOutputElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOutputElement "The HTMLOutputElement interface provides properties and methods (beyond those inherited from HTMLElement) for manipulating the layout and presentation of <output> elements.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-output-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-output-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-output-element:phrasing-content-2}.
:   [Listed](forms.html#category-listed){#the-output-element:category-listed},
    [labelable](forms.html#category-label){#the-output-element:category-label},
    [resettable](forms.html#category-reset){#the-output-element:category-reset},
    and [autocapitalize-and-autocorrect
    inheriting](forms.html#category-autocapitalize){#the-output-element:category-autocapitalize}
    [form-associated
    element](forms.html#form-associated-element){#the-output-element:form-associated-element}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-output-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-output-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-output-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-output-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-output-element:phrasing-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-output-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-output-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-output-element:global-attributes}
:   [`for`{#the-output-element:attr-output-for}](#attr-output-for) ---
    Specifies controls from which the output was calculated
:   [`form`{#the-output-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    --- Associates the element with a
    [`form`{#the-output-element:the-form-element}](forms.html#the-form-element)
    element
:   [`name`{#the-output-element:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    --- Name of the element to use in the
    [`form.elements`{#the-output-element:dom-form-elements}](forms.html#dom-form-elements)
    API.

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-output-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-output).
:   [For implementers](https://w3c.github.io/html-aam/#el-output).

[DOM interface](dom.html#concept-element-dom){#the-output-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLOutputElement : HTMLElement {
      [HTMLConstructor] constructor();

      [SameObject, PutForwards=value] readonly attribute DOMTokenList htmlFor;
      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute DOMString name;

      readonly attribute DOMString type;
      [CEReactions] attribute DOMString defaultValue;
      [CEReactions] attribute DOMString value;

      readonly attribute boolean willValidate;
      readonly attribute ValidityState validity;
      readonly attribute DOMString validationMessage;
      boolean checkValidity();
      boolean reportValidity();
      undefined setCustomValidity(DOMString error);

      readonly attribute NodeList labels;
    };
    ```

The
[`output`{#the-output-element:the-output-element}](#the-output-element)
element
[represents](dom.html#represents){#the-output-element:represents} the
result of a calculation performed by the application, or the result of a
user action.

This element can be contrasted with the
[`samp`{#the-output-element:the-samp-element}](text-level-semantics.html#the-samp-element)
element, which is the appropriate element for quoting the output of
other programs run previously.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/for](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/for "The for attribute is an allowed attribute for <label> and <output>. When used on a <label> element it indicates the form element that this label describes. When used on an <output> element it allows for an explicit relationship between the elements that represent values which are used in the output.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

The [`for`]{#attr-output-for .dfn dfn-for="output"
dfn-type="element-attr"} content attribute allows an explicit
relationship to be made between the result of a calculation and the
elements that represent the values that went into the calculation or
that otherwise influenced the calculation. The
[`for`{#the-output-element:attr-output-for-2}](#attr-output-for)
attribute, if specified, must contain a string consisting of an
[unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#the-output-element:unordered-set-of-unique-space-separated-tokens},
none of which are [identical
to](https://infra.spec.whatwg.org/#string-is){#the-output-element:identical-to
x-internal="identical-to"} another token and each of which must have the
value of an
[ID](https://dom.spec.whatwg.org/#concept-id){#the-output-element:concept-id
x-internal="concept-id"} of an element in the same
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-output-element:tree
x-internal="tree"}.

The
[`form`{#the-output-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the
[`output`{#the-output-element:the-output-element-2}](#the-output-element)
element with its [form
owner](form-control-infrastructure.html#form-owner){#the-output-element:form-owner}.
The
[`name`{#the-output-element:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attribute represents the element\'s name. The
[`output`{#the-output-element:the-output-element-3}](#the-output-element)
element is associated with a form so that it can be easily
[referenced](dom.html#referenced){#the-output-element:referenced} from
the event handlers of form controls; the element\'s value itself is not
submitted when the form is submitted.

The element has a [default value
override]{#concept-output-default-value-override .dfn} (null or a
string). Initially it must be null.

The element\'s [default value]{#concept-output-default-value .dfn} is
determined by the following steps:

1.  If this element\'s [default value
    override](#concept-output-default-value-override){#the-output-element:concept-output-default-value-override}
    is non-null, then return it.

2.  Return this element\'s [descendant text
    content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#the-output-element:descendant-text-content
    x-internal="descendant-text-content"}.

The [reset
algorithm](form-control-infrastructure.html#concept-form-reset-control){#the-output-element:concept-form-reset-control}
for
[`output`{#the-output-element:the-output-element-4}](#the-output-element)
elements is to run these steps:

1.  [String replace
    all](https://dom.spec.whatwg.org/#string-replace-all){#the-output-element:string-replace-all
    x-internal="string-replace-all"} with this element\'s [default
    value](#concept-output-default-value){#the-output-element:concept-output-default-value}
    within this element.

2.  Set this element\'s [default value
    override](#concept-output-default-value-override){#the-output-element:concept-output-default-value-override-2}
    to null.

`output`{.variable}`.`[`value`](#dom-output-value){#dom-output-value-dev}` [ = ``value`{.variable}` ]`

:   Returns the element\'s current value.

    Can be set, to change the value.

`output`{.variable}`.`[`defaultValue`](#dom-output-defaultvalue){#dom-output-defaultvalue-dev}` [ = ``value`{.variable}` ]`

:   Returns the element\'s current default value.

    Can be set, to change the default value.

`output`{.variable}`.`[`type`](#dom-output-type){#dom-output-type-dev}

:   Returns the string \"`output`\".

The [`value`]{#dom-output-value .dfn dfn-for="HTMLOutputElement"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-output-element:this
x-internal="this"}\'s [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#the-output-element:descendant-text-content-2
x-internal="descendant-text-content"}.

The [`value`{#the-output-element:dom-output-value-2}](#dom-output-value)
setter steps are:

1.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-output-element:this-2
    x-internal="this"}\'s [default value
    override](#concept-output-default-value-override){#the-output-element:concept-output-default-value-override-3}
    to its [default
    value](#concept-output-default-value){#the-output-element:concept-output-default-value-2}.

2.  [String replace
    all](https://dom.spec.whatwg.org/#string-replace-all){#the-output-element:string-replace-all-2
    x-internal="string-replace-all"} with the given value within
    [this](https://webidl.spec.whatwg.org/#this){#the-output-element:this-3
    x-internal="this"}.

The [`defaultValue`]{#dom-output-defaultvalue .dfn
dfn-for="HTMLOutputElement" dfn-type="attribute"} getter steps are to
return the result of running
[this](https://webidl.spec.whatwg.org/#this){#the-output-element:this-4
x-internal="this"}\'s [default
value](#concept-output-default-value){#the-output-element:concept-output-default-value-3}.

The
[`defaultValue`{#the-output-element:dom-output-defaultvalue-2}](#dom-output-defaultvalue)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-output-element:this-5
    x-internal="this"}\'s [default value
    override](#concept-output-default-value-override){#the-output-element:concept-output-default-value-override-4}
    is null, then [string replace
    all](https://dom.spec.whatwg.org/#string-replace-all){#the-output-element:string-replace-all-3
    x-internal="string-replace-all"} with the given value within
    [this](https://webidl.spec.whatwg.org/#this){#the-output-element:this-6
    x-internal="this"} and return.

2.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-output-element:this-7
    x-internal="this"}\'s [default value
    override](#concept-output-default-value-override){#the-output-element:concept-output-default-value-override-5}
    to the given value.

The [`type`]{#dom-output-type .dfn dfn-for="HTMLOutputElement"
dfn-type="attribute"} getter steps are to return \"`output`\".

The [`htmlFor`]{#dom-output-htmlfor .dfn dfn-for="HTMLOutputElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-output-element:reflect}
the [`for`{#the-output-element:attr-output-for-3}](#attr-output-for)
content attribute.

The
[`willValidate`{#the-output-element:dom-cva-willvalidate-2}](form-control-infrastructure.html#dom-cva-willvalidate),
[`validity`{#the-output-element:dom-cva-validity-2}](form-control-infrastructure.html#dom-cva-validity),
and
[`validationMessage`{#the-output-element:dom-cva-validationmessage-2}](form-control-infrastructure.html#dom-cva-validationmessage)
IDL attributes, and the
[`checkValidity()`{#the-output-element:dom-cva-checkvalidity-2}](form-control-infrastructure.html#dom-cva-checkvalidity),
[`reportValidity()`{#the-output-element:dom-cva-reportvalidity-2}](form-control-infrastructure.html#dom-cva-reportvalidity),
and
[`setCustomValidity()`{#the-output-element:dom-cva-setcustomvalidity-2}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
methods, are part of the [constraint validation
API](form-control-infrastructure.html#the-constraint-validation-api){#the-output-element:the-constraint-validation-api}.
The
[`labels`{#the-output-element:dom-lfe-labels-2}](forms.html#dom-lfe-labels)
IDL attribute provides a list of the element\'s
[`label`{#the-output-element:the-label-element}](forms.html#the-label-element)s.
The
[`form`{#the-output-element:dom-fae-form-2}](form-control-infrastructure.html#dom-fae-form)
and
[`name`{#the-output-element:dom-fe-name-2}](form-control-infrastructure.html#dom-fe-name)
IDL attributes are part of the element\'s forms API.

::: example
A simple calculator could use
[`output`{#the-output-element:the-output-element-5}](#the-output-element)
for its display of calculated results:

``` html
<form onsubmit="return false" oninput="o.value = a.valueAsNumber + b.valueAsNumber">
 <input id=a type=number step=any> +
 <input id=b type=number step=any> =
 <output id=o for="a b"></output>
</form>
```
:::

::: example
In this example, an
[`output`{#the-output-element:the-output-element-6}](#the-output-element)
element is used to report the results of a calculation performed by a
remote server, as they come in:

``` html
<output id="result"></output>
<script>
 var primeSource = new WebSocket('ws://primes.example.net/');
 primeSource.onmessage = function (event) {
   document.getElementById('result').value = event.data;
 }
</script>
```
:::

#### [4.10.13]{.secno} The [`progress`]{.dfn dfn-type="element"} element[](#the-progress-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/progress](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress "The <progress> HTML element displays an indicator showing the completion progress of a task, typically displayed as a progress bar.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLProgressElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLProgressElement "The HTMLProgressElement interface provides special properties and methods (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of <progress> elements.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-progress-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-progress-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-progress-element:phrasing-content-2}.
:   [Labelable
    element](forms.html#category-label){#the-progress-element:category-label}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-progress-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-progress-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-progress-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-progress-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-progress-element:phrasing-content-2-3},
    but there must be no
    [`progress`{#the-progress-element:the-progress-element}](#the-progress-element)
    element descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-progress-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-progress-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-progress-element:global-attributes}
:   [`value`{#the-progress-element:attr-progress-value}](#attr-progress-value)
    --- Current value of the element
:   [`max`{#the-progress-element:attr-progress-max}](#attr-progress-max)
    --- Upper bound of range

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-progress-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-progress).
:   [For implementers](https://w3c.github.io/html-aam/#el-progress).

[DOM interface](dom.html#concept-element-dom){#the-progress-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLProgressElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute double value;
      [CEReactions] attribute double max;
      readonly attribute double position;
      readonly attribute NodeList labels;
    };
    ```

The
[`progress`{#the-progress-element:the-progress-element-2}](#the-progress-element)
element
[represents](dom.html#represents){#the-progress-element:represents} the
completion progress of a task. The progress is either indeterminate,
indicating that progress is being made but that it is not clear how much
more work remains to be done before the task is complete (e.g. because
the task is waiting for a remote host to respond), or the progress is a
number in the range zero to a maximum, giving the fraction of work that
has so far been completed.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/max](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/max "The max attribute defines the maximum value that is acceptable and valid for the input containing the attribute. If the value of the element is greater than this, the element fails validation. This value must be greater than or equal to the value of the min attribute. If the max attribute is present but is not specified or is invalid, no max value is applied. If the max attribute is valid and a non-empty value is greater than the maximum allowed by the max attribute, constraint validation will prevent form submission.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS7+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

There are two attributes that determine the current task completion
represented by the element. The [`value`]{#attr-progress-value .dfn
dfn-for="progress" dfn-type="element-attr"} attribute specifies how much
of the task has been completed, and the [`max`]{#attr-progress-max .dfn
dfn-for="progress" dfn-type="element-attr"} attribute specifies how much
work the task requires in total. The units are arbitrary and not
specified.

To make a determinate progress bar, add a
[`value`{#the-progress-element:attr-progress-value-2}](#attr-progress-value)
attribute with the current progress (either a number from 0.0 to 1.0,
or, if the
[`max`{#the-progress-element:attr-progress-max-2}](#attr-progress-max)
attribute is specified, a number from 0 to the value of the
[`max`{#the-progress-element:attr-progress-max-3}](#attr-progress-max)
attribute). To make an indeterminate progress bar, remove the
[`value`{#the-progress-element:attr-progress-value-3}](#attr-progress-value)
attribute.

Authors are encouraged to also include the current value and the maximum
value inline as text inside the element, so that the progress is made
available to users of legacy user agents.

::: example
Here is a snippet of a web application that shows the progress of some
automated task:

``` html
<section>
 <h2>Task Progress</h2>
 <p>Progress: <progress id=p max=100><span>0</span>%</progress></p>
 <script>
  var progressBar = document.getElementById('p');
  function updateProgress(newValue) {
    progressBar.value = newValue;
    progressBar.getElementsByTagName('span')[0].textContent = newValue;
  }
 </script>
</section>
```

(The `updateProgress()` method in this example would be called by some
other code on the page to update the actual progress bar as the task
progressed.)
:::

The
[`value`{#the-progress-element:attr-progress-value-4}](#attr-progress-value)
and
[`max`{#the-progress-element:attr-progress-max-4}](#attr-progress-max)
attributes, when present, must have values that are [valid
floating-point
numbers](common-microsyntaxes.html#valid-floating-point-number){#the-progress-element:valid-floating-point-number}.
The
[`value`{#the-progress-element:attr-progress-value-5}](#attr-progress-value)
attribute, if present, must have a value greater than or equal to zero,
and less than or equal to the value of the
[`max`{#the-progress-element:attr-progress-max-5}](#attr-progress-max)
attribute, if present, or 1.0, otherwise. The
[`max`{#the-progress-element:attr-progress-max-6}](#attr-progress-max)
attribute, if present, must have a value greater than zero.

The
[`progress`{#the-progress-element:the-progress-element-3}](#the-progress-element)
element is the wrong element to use for something that is just a gauge,
as opposed to task progress. For instance, indicating disk space usage
using
[`progress`{#the-progress-element:the-progress-element-4}](#the-progress-element)
would be inappropriate. Instead, the
[`meter`{#the-progress-element:the-meter-element}](#the-meter-element)
element is available for such use cases.

**User agent requirements**: If the
[`value`{#the-progress-element:attr-progress-value-6}](#attr-progress-value)
attribute is omitted, then the progress bar is an indeterminate progress
bar. Otherwise, it is a determinate progress bar.

If the progress bar is a determinate progress bar and the element has a
[`max`{#the-progress-element:attr-progress-max-7}](#attr-progress-max)
attribute, the user agent must parse the
[`max`{#the-progress-element:attr-progress-max-8}](#attr-progress-max)
attribute\'s value according to the [rules for parsing floating-point
number
values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#the-progress-element:rules-for-parsing-floating-point-number-values}.
If this does not result in an error, and if the parsed value is greater
than zero, then the [maximum value]{#concept-progress-maximum .dfn} of
the progress bar is that value. Otherwise, if the element has no
[`max`{#the-progress-element:attr-progress-max-9}](#attr-progress-max)
attribute, or if it has one but parsing it resulted in an error, or if
the parsed value was less than or equal to zero, then the [maximum
value](#concept-progress-maximum){#the-progress-element:concept-progress-maximum}
of the progress bar is 1.0.

If the progress bar is a determinate progress bar, user agents must
parse the
[`value`{#the-progress-element:attr-progress-value-7}](#attr-progress-value)
attribute\'s value according to the [rules for parsing floating-point
number
values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#the-progress-element:rules-for-parsing-floating-point-number-values-2}.
If this does not result in an error and the parsed value is greater than
zero, then the [value]{#concept-progress-value .dfn} of the progress bar
is that parsed value. Otherwise, if parsing the
[`value`{#the-progress-element:attr-progress-value-8}](#attr-progress-value)
attribute\'s value resulted in an error or a number less than or equal
to zero, then the
[value](#concept-progress-value){#the-progress-element:concept-progress-value}
of the progress bar is zero.

If the progress bar is a determinate progress bar, then the [current
value]{#concept-progress-current-value .dfn} is the [maximum
value](#concept-progress-maximum){#the-progress-element:concept-progress-maximum-2},
if
[value](#concept-progress-value){#the-progress-element:concept-progress-value-2}
is greater than the [maximum
value](#concept-progress-maximum){#the-progress-element:concept-progress-maximum-3},
and
[value](#concept-progress-value){#the-progress-element:concept-progress-value-3}
otherwise.

**UA requirements for showing the progress bar**: When representing a
[`progress`{#the-progress-element:the-progress-element-5}](#the-progress-element)
element to the user, the UA should indicate whether it is a determinate
or indeterminate progress bar, and in the former case, should indicate
the relative position of the [current
value](#concept-progress-current-value){#the-progress-element:concept-progress-current-value}
relative to the [maximum
value](#concept-progress-maximum){#the-progress-element:concept-progress-maximum-4}.

`progress`{.variable}`.`[`position`](#dom-progress-position){#dom-progress-position-dev}

:   For a determinate progress bar (one with known current and maximum
    values), returns the result of dividing the current value by the
    maximum value.

    For an indeterminate progress bar, returns −1.

If the progress bar is an indeterminate progress bar, then the
[`position`]{#dom-progress-position .dfn dfn-for="HTMLProgressElement"
dfn-type="attribute"} IDL attribute must return −1. Otherwise, it must
return the result of dividing the [current
value](#concept-progress-current-value){#the-progress-element:concept-progress-current-value-2}
by the [maximum
value](#concept-progress-maximum){#the-progress-element:concept-progress-maximum-5}.

If the progress bar is an indeterminate progress bar, then the
[`value`]{#dom-progress-value .dfn dfn-for="HTMLProgressElement"
dfn-type="attribute"} IDL attribute, on getting, must return 0.
Otherwise, it must return the [current
value](#concept-progress-value){#the-progress-element:concept-progress-value-4}.
On setting, the given value must be converted to the [best
representation of the number as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#the-progress-element:best-representation-of-the-number-as-a-floating-point-number}
and then the
[`value`{#the-progress-element:dom-progress-value-2}](#dom-progress-value)
content attribute must be set to that string.

Setting the
[`value`{#the-progress-element:dom-progress-value-3}](#dom-progress-value)
IDL attribute to itself when the corresponding content attribute is
absent would change the progress bar from an indeterminate progress bar
to a determinate progress bar with no progress.

The [`max`]{#dom-progress-max .dfn dfn-for="HTMLProgressElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-progress-element:reflect}
the content attribute of the same name, [limited to only positive
numbers](common-dom-interfaces.html#limited-to-only-non-negative-numbers-greater-than-zero){#the-progress-element:limited-to-only-non-negative-numbers-greater-than-zero}.
The [default
value](common-dom-interfaces.html#default-value){#the-progress-element:default-value}
for [`max`{#the-progress-element:dom-progress-max-2}](#dom-progress-max)
is 1.0.

The
[`labels`{#the-progress-element:dom-lfe-labels-2}](forms.html#dom-lfe-labels)
IDL attribute provides a list of the element\'s
[`label`{#the-progress-element:the-label-element}](forms.html#the-label-element)s.

#### [4.10.14]{.secno} The [`meter`]{.dfn dfn-type="element"} element[](#the-meter-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/meter](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter "The <meter> HTML element represents either a scalar value within a known range or a fractional value.")

Support in all current engines.

::: support
[Firefox16+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS10.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView AndroidNo]{.webview_android .no}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLMeterElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMeterElement "The HTML <meter> elements expose the HTMLMeterElement interface, which provides special properties and methods (beyond the HTMLElement object interface they also have available to them by inheritance) for manipulating the layout and presentation of <meter> elements.")

Support in all current engines.

::: support
[Firefox16+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-meter-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-meter-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-meter-element:phrasing-content-2}.
:   [Labelable
    element](forms.html#category-label){#the-meter-element:category-label}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-meter-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-meter-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-meter-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-meter-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-meter-element:phrasing-content-2-3},
    but there must be no
    [`meter`{#the-meter-element:the-meter-element}](#the-meter-element)
    element descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-meter-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-meter-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-meter-element:global-attributes}
:   [`value`{#the-meter-element:attr-meter-value}](#attr-meter-value)
    --- Current value of the element
:   [`min`{#the-meter-element:attr-meter-min}](#attr-meter-min) ---
    Lower bound of range
:   [`max`{#the-meter-element:attr-meter-max}](#attr-meter-max) ---
    Upper bound of range
:   [`low`{#the-meter-element:attr-meter-low}](#attr-meter-low) --- High
    limit of low range
:   [`high`{#the-meter-element:attr-meter-high}](#attr-meter-high) ---
    Low limit of high range
:   [`optimum`{#the-meter-element:attr-meter-optimum}](#attr-meter-optimum)
    --- Optimum value in gauge

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-meter-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-meter).
:   [For implementers](https://w3c.github.io/html-aam/#el-meter).

[DOM interface](dom.html#concept-element-dom){#the-meter-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLMeterElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute double value;
      [CEReactions] attribute double min;
      [CEReactions] attribute double max;
      [CEReactions] attribute double low;
      [CEReactions] attribute double high;
      [CEReactions] attribute double optimum;
      readonly attribute NodeList labels;
    };
    ```

The
[`meter`{#the-meter-element:the-meter-element-2}](#the-meter-element)
element [represents](dom.html#represents){#the-meter-element:represents}
a scalar measurement within a known range, or a fractional value; for
example disk usage, the relevance of a query result, or the fraction of
a voting population to have selected a particular candidate.

This is also known as a gauge.

The
[`meter`{#the-meter-element:the-meter-element-3}](#the-meter-element)
element should not be used to indicate progress (as in a progress bar).
For that role, HTML provides a separate
[`progress`{#the-meter-element:the-progress-element}](#the-progress-element)
element.

The
[`meter`{#the-meter-element:the-meter-element-4}](#the-meter-element)
element also does not represent a scalar value of arbitrary range ---
for example, it would be wrong to use this to report a weight, or
height, unless there is a known maximum value.

There are six attributes that determine the semantics of the gauge
represented by the element.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/max](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/max "The max attribute defines the maximum value that is acceptable and valid for the input containing the attribute. If the value of the element is greater than this, the element fails validation. This value must be greater than or equal to the value of the min attribute. If the max attribute is present but is not specified or is invalid, no max value is applied. If the max attribute is valid and a non-empty value is greater than the maximum allowed by the max attribute, constraint validation will prevent form submission.")

Support in all current engines.

::: support
[Firefox16+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS10.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView AndroidNo]{.webview_android .no}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::

:::: feature
[Attributes/min](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/min "The min attribute defines the minimum value that is acceptable and valid for the input containing the attribute. If the value of the element is less than this, the element fails validation. This value must be less than or equal to the value of the max attribute.")

Support in all current engines.

::: support
[Firefox16+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome6+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari
iOS10.3+]{.safari_ios .yes}[Chrome Android?]{.chrome_android
.unknown}[WebView AndroidNo]{.webview_android .no}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::::

The [`min`]{#attr-meter-min .dfn dfn-for="meter"
dfn-type="element-attr"} attribute specifies the lower bound of the
range, and the [`max`]{#attr-meter-max .dfn dfn-for="meter"
dfn-type="element-attr"} attribute specifies the upper bound. The
[`value`]{#attr-meter-value .dfn dfn-for="meter"
dfn-type="element-attr"} attribute specifies the value to have the gauge
indicate as the \"measured\" value.

The other three attributes can be used to segment the gauge\'s range
into \"low\", \"medium\", and \"high\" parts, and to indicate which part
of the gauge is the \"optimum\" part. The [`low`]{#attr-meter-low .dfn
dfn-for="meter" dfn-type="element-attr"} attribute specifies the range
that is considered to be the \"low\" part, and the
[`high`]{#attr-meter-high .dfn dfn-for="meter" dfn-type="element-attr"}
attribute specifies the range that is considered to be the \"high\"
part. The [`optimum`]{#attr-meter-optimum .dfn dfn-for="meter"
dfn-type="element-attr"} attribute gives the position that is
\"optimum\"; if that is higher than the \"high\" value then this
indicates that the higher the value, the better; if it\'s lower than the
\"low\" mark then it indicates that lower values are better, and
naturally if it is in between then it indicates that neither high nor
low values are good.

**Authoring requirements**: The
[`value`{#the-meter-element:attr-meter-value-2}](#attr-meter-value)
attribute must be specified. The
[`value`{#the-meter-element:attr-meter-value-3}](#attr-meter-value),
[`min`{#the-meter-element:attr-meter-min-2}](#attr-meter-min),
[`low`{#the-meter-element:attr-meter-low-2}](#attr-meter-low),
[`high`{#the-meter-element:attr-meter-high-2}](#attr-meter-high),
[`max`{#the-meter-element:attr-meter-max-2}](#attr-meter-max), and
[`optimum`{#the-meter-element:attr-meter-optimum-2}](#attr-meter-optimum)
attributes, when present, must have values that are [valid
floating-point
numbers](common-microsyntaxes.html#valid-floating-point-number){#the-meter-element:valid-floating-point-number}.

In addition, the attributes\' values are further constrained:

Let `value`{.variable} be the
[`value`{#the-meter-element:attr-meter-value-4}](#attr-meter-value)
attribute\'s number.

If the [`min`{#the-meter-element:attr-meter-min-3}](#attr-meter-min)
attribute is specified, then let `minimum`{.variable} be that
attribute\'s value; otherwise, let it be zero.

If the [`max`{#the-meter-element:attr-meter-max-3}](#attr-meter-max)
attribute is specified, then let `maximum`{.variable} be that
attribute\'s value; otherwise, let it be 1.0.

The following inequalities must hold, as applicable:

- `minimum`{.variable} ≤ `value`{.variable} ≤ `maximum`{.variable}

- `minimum`{.variable} ≤
  [`low`{#the-meter-element:attr-meter-low-3}](#attr-meter-low) ≤
  `maximum`{.variable} (if
  [`low`{#the-meter-element:attr-meter-low-4}](#attr-meter-low) is
  specified)

- `minimum`{.variable} ≤
  [`high`{#the-meter-element:attr-meter-high-3}](#attr-meter-high) ≤
  `maximum`{.variable} (if
  [`high`{#the-meter-element:attr-meter-high-4}](#attr-meter-high) is
  specified)

- `minimum`{.variable} ≤
  [`optimum`{#the-meter-element:attr-meter-optimum-3}](#attr-meter-optimum)
  ≤ `maximum`{.variable} (if
  [`optimum`{#the-meter-element:attr-meter-optimum-4}](#attr-meter-optimum)
  is specified)

- [`low`{#the-meter-element:attr-meter-low-5}](#attr-meter-low) ≤
  [`high`{#the-meter-element:attr-meter-high-5}](#attr-meter-high) (if
  both [`low`{#the-meter-element:attr-meter-low-6}](#attr-meter-low) and
  [`high`{#the-meter-element:attr-meter-high-6}](#attr-meter-high) are
  specified)

If no minimum or maximum is specified, then the range is assumed to be
0..1, and the value thus has to be within that range.

Authors are encouraged to include a textual representation of the
gauge\'s state in the element\'s contents, for users of user agents that
do not support the
[`meter`{#the-meter-element:the-meter-element-5}](#the-meter-element)
element.

When used with
[microdata](microdata.html#microdata){#the-meter-element:microdata}, the
[`meter`{#the-meter-element:the-meter-element-6}](#the-meter-element)
element\'s
[`value`{#the-meter-element:attr-meter-value-5}](#attr-meter-value)
attribute provides the element\'s machine-readable value.

::: example
The following examples show three gauges that would all be
three-quarters full:

``` html
Storage space usage: <meter value=6 max=8>6 blocks used (out of 8 total)</meter>
```

``` html
Voter turnout: <meter value=0.75><img alt="75%" src="graph75.png"></meter>
```

``` html
Tickets sold: <meter min="0" max="100" value="75"></meter>
```

The following example is incorrect use of the element, because it
doesn\'t give a range (and since the default maximum is 1, both of the
gauges would end up looking maxed out):

``` bad
<p>The grapefruit pie had a radius of <meter value=12>12cm</meter>
and a height of <meter value=2>2cm</meter>.</p> <!-- BAD! -->
```

Instead, one would either not include the meter element, or use the
meter element with a defined range to give the dimensions in context
compared to other pies:

``` html
<p>The grapefruit pie had a radius of 12cm and a height of
2cm.</p>
<dl>
 <dt>Radius: <dd> <meter min=0 max=20 value=12>12cm</meter>
 <dt>Height: <dd> <meter min=0 max=10 value=2>2cm</meter>
</dl>
```
:::

There is no explicit way to specify units in the
[`meter`{#the-meter-element:the-meter-element-7}](#the-meter-element)
element, but the units may be specified in the
[`title`{#the-meter-element:attr-title}](dom.html#attr-title) attribute
in free-form text.

::: example
The example above could be extended to mention the units:

``` html
<dl>
 <dt>Radius: <dd> <meter min=0 max=20 value=12 title="centimeters">12cm</meter>
 <dt>Height: <dd> <meter min=0 max=10 value=2 title="centimeters">2cm</meter>
</dl>
```
:::

**User agent requirements**: User agents must parse the
[`min`{#the-meter-element:attr-meter-min-4}](#attr-meter-min),
[`max`{#the-meter-element:attr-meter-max-4}](#attr-meter-max),
[`value`{#the-meter-element:attr-meter-value-6}](#attr-meter-value),
[`low`{#the-meter-element:attr-meter-low-7}](#attr-meter-low),
[`high`{#the-meter-element:attr-meter-high-7}](#attr-meter-high), and
[`optimum`{#the-meter-element:attr-meter-optimum-5}](#attr-meter-optimum)
attributes using the [rules for parsing floating-point number
values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#the-meter-element:rules-for-parsing-floating-point-number-values}.

User agents must then use all these numbers to obtain values for six
points on the gauge, as follows. (The order in which these are evaluated
is important, as some of the values refer to earlier ones.)

The [minimum value]{#concept-meter-minimum .dfn}

:   If the [`min`{#the-meter-element:attr-meter-min-5}](#attr-meter-min)
    attribute is specified and a value could be parsed out of it, then
    the minimum value is that value. Otherwise, the minimum value is
    zero.

The [maximum value]{#concept-meter-maximum .dfn}

:   If the [`max`{#the-meter-element:attr-meter-max-5}](#attr-meter-max)
    attribute is specified and a value could be parsed out of it, then
    the candidate maximum value is that value. Otherwise, the candidate
    maximum value is 1.0.

    If the candidate maximum value is greater than or equal to the
    minimum value, then the maximum value is the candidate maximum
    value. Otherwise, the maximum value is the same as the minimum
    value.

The [actual value]{#concept-meter-actual .dfn}

:   If the
    [`value`{#the-meter-element:attr-meter-value-7}](#attr-meter-value)
    attribute is specified and a value could be parsed out of it, then
    that value is the candidate actual value. Otherwise, the candidate
    actual value is zero.

    If the candidate actual value is less than the minimum value, then
    the actual value is the minimum value.

    Otherwise, if the candidate actual value is greater than the maximum
    value, then the actual value is the maximum value.

    Otherwise, the actual value is the candidate actual value.

The [low boundary]{#concept-meter-low .dfn}

:   If the [`low`{#the-meter-element:attr-meter-low-8}](#attr-meter-low)
    attribute is specified and a value could be parsed out of it, then
    the candidate low boundary is that value. Otherwise, the candidate
    low boundary is the same as the minimum value.

    If the candidate low boundary is less than the minimum value, then
    the low boundary is the minimum value.

    Otherwise, if the candidate low boundary is greater than the maximum
    value, then the low boundary is the maximum value.

    Otherwise, the low boundary is the candidate low boundary.

The [high boundary]{#concept-meter-high .dfn}

:   If the
    [`high`{#the-meter-element:attr-meter-high-8}](#attr-meter-high)
    attribute is specified and a value could be parsed out of it, then
    the candidate high boundary is that value. Otherwise, the candidate
    high boundary is the same as the maximum value.

    If the candidate high boundary is less than the low boundary, then
    the high boundary is the low boundary.

    Otherwise, if the candidate high boundary is greater than the
    maximum value, then the high boundary is the maximum value.

    Otherwise, the high boundary is the candidate high boundary.

The [optimum point]{#concept-meter-optimum .dfn}

:   If the
    [`optimum`{#the-meter-element:attr-meter-optimum-6}](#attr-meter-optimum)
    attribute is specified and a value could be parsed out of it, then
    the candidate optimum point is that value. Otherwise, the candidate
    optimum point is the midpoint between the minimum value and the
    maximum value.

    If the candidate optimum point is less than the minimum value, then
    the optimum point is the minimum value.

    Otherwise, if the candidate optimum point is greater than the
    maximum value, then the optimum point is the maximum value.

    Otherwise, the optimum point is the candidate optimum point.

All of which will result in the following inequalities all being true:

- minimum value ≤ actual value ≤ maximum value

- minimum value ≤ low boundary ≤ high boundary ≤ maximum value

- minimum value ≤ optimum point ≤ maximum value

**UA requirements for regions of the gauge**: If the optimum point is
equal to the low boundary or the high boundary, or anywhere in between
them, then the region between the low and high boundaries of the gauge
must be treated as the optimum region, and the low and high parts, if
any, must be treated as suboptimal. Otherwise, if the optimum point is
less than the low boundary, then the region between the minimum value
and the low boundary must be treated as the optimum region, the region
from the low boundary up to the high boundary must be treated as a
suboptimal region, and the remaining region must be treated as an even
less good region. Finally, if the optimum point is higher than the high
boundary, then the situation is reversed; the region between the high
boundary and the maximum value must be treated as the optimum region,
the region from the high boundary down to the low boundary must be
treated as a suboptimal region, and the remaining region must be treated
as an even less good region.

**UA requirements for showing the gauge**: When representing a
[`meter`{#the-meter-element:the-meter-element-8}](#the-meter-element)
element to the user, the UA should indicate the relative position of the
actual value to the minimum and maximum values, and the relationship
between the actual value and the three regions of the gauge.

::: example
The following markup:

``` html
<h3>Suggested groups</h3>
<menu>
 <li><a href="?cmd=hsg" onclick="hideSuggestedGroups()">Hide suggested groups</a></li>
</menu>
<ul>
 <li>
  <p><a href="/group/comp.infosystems.www.authoring.stylesheets/view">comp.infosystems.www.authoring.stylesheets</a> -
     <a href="/group/comp.infosystems.www.authoring.stylesheets/subscribe">join</a></p>
  <p>Group description: <strong>Layout/presentation on the WWW.</strong></p>
  <p><meter value="0.5">Moderate activity,</meter> Usenet, 618 subscribers</p>
 </li>
 <li>
  <p><a href="/group/netscape.public.mozilla.xpinstall/view">netscape.public.mozilla.xpinstall</a> -
     <a href="/group/netscape.public.mozilla.xpinstall/subscribe">join</a></p>
  <p>Group description: <strong>Mozilla XPInstall discussion.</strong></p>
  <p><meter value="0.25">Low activity,</meter> Usenet, 22 subscribers</p>
 </li>
 <li>
  <p><a href="/group/mozilla.dev.general/view">mozilla.dev.general</a> -
     <a href="/group/mozilla.dev.general/subscribe">join</a></p>
  <p><meter value="0.25">Low activity,</meter> Usenet, 66 subscribers</p>
 </li>
</ul>
```

Might be rendered as follows:

![With the \<meter\> elements rendered as inline green bars of varying
lengths.](/images/sample-meter.png){width="332" height="178"}
:::

User agents may combine the value of the
[`title`{#the-meter-element:attr-title-2}](dom.html#attr-title)
attribute and the other attributes to provide context-sensitive help or
inline text detailing the actual values.

::: example
For example, the following snippet:

``` html
<meter min=0 max=60 value=23.2 title=seconds></meter>
```

\...might cause the user agent to display a gauge with a tooltip saying
\"Value: 23.2 out of 60.\" on one line and \"seconds\" on a second line.
:::

The [`value`]{#dom-meter-value .dfn dfn-for="HTMLMeterElement"
dfn-type="attribute"} IDL attribute, on getting, must return the [actual
value](#concept-meter-actual){#the-meter-element:concept-meter-actual}.
On setting, the given value must be converted to the [best
representation of the number as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#the-meter-element:best-representation-of-the-number-as-a-floating-point-number}
and then the
[`value`{#the-meter-element:attr-meter-value-8}](#attr-meter-value)
content attribute must be set to that string.

The [`min`]{#dom-meter-min .dfn dfn-for="HTMLMeterElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
[minimum
value](#concept-meter-minimum){#the-meter-element:concept-meter-minimum}.
On setting, the given value must be converted to the [best
representation of the number as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#the-meter-element:best-representation-of-the-number-as-a-floating-point-number-2}
and then the
[`min`{#the-meter-element:attr-meter-min-6}](#attr-meter-min) content
attribute must be set to that string.

The [`max`]{#dom-meter-max .dfn dfn-for="HTMLMeterElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
[maximum
value](#concept-meter-maximum){#the-meter-element:concept-meter-maximum}.
On setting, the given value must be converted to the [best
representation of the number as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#the-meter-element:best-representation-of-the-number-as-a-floating-point-number-3}
and then the
[`max`{#the-meter-element:attr-meter-max-6}](#attr-meter-max) content
attribute must be set to that string.

The [`low`]{#dom-meter-low .dfn dfn-for="HTMLMeterElement"
dfn-type="attribute"} IDL attribute, on getting, must return the [low
boundary](#concept-meter-low){#the-meter-element:concept-meter-low}. On
setting, the given value must be converted to the [best representation
of the number as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#the-meter-element:best-representation-of-the-number-as-a-floating-point-number-4}
and then the
[`low`{#the-meter-element:attr-meter-low-9}](#attr-meter-low) content
attribute must be set to that string.

The [`high`]{#dom-meter-high .dfn dfn-for="HTMLMeterElement"
dfn-type="attribute"} IDL attribute, on getting, must return the [high
boundary](#concept-meter-high){#the-meter-element:concept-meter-high}.
On setting, the given value must be converted to the [best
representation of the number as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#the-meter-element:best-representation-of-the-number-as-a-floating-point-number-5}
and then the
[`high`{#the-meter-element:attr-meter-high-9}](#attr-meter-high) content
attribute must be set to that string.

The [`optimum`]{#dom-meter-optimum .dfn dfn-for="HTMLMeterElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
[optimum
value](#concept-meter-optimum){#the-meter-element:concept-meter-optimum}.
On setting, the given value must be converted to the [best
representation of the number as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#the-meter-element:best-representation-of-the-number-as-a-floating-point-number-6}
and then the
[`optimum`{#the-meter-element:attr-meter-optimum-7}](#attr-meter-optimum)
content attribute must be set to that string.

The
[`labels`{#the-meter-element:dom-lfe-labels-2}](forms.html#dom-lfe-labels)
IDL attribute provides a list of the element\'s
[`label`{#the-meter-element:the-label-element}](forms.html#the-label-element)s.

::: example
The following example shows how a gauge could fall back to localized or
pretty-printed text.

``` html
<p>Disk usage: <meter min=0 value=170261928 max=233257824>170 261 928 bytes used
out of 233 257 824 bytes available</meter></p>
```
:::

#### [4.10.15]{.secno} The [`fieldset`]{.dfn dfn-type="element"} element[](#the-fieldset-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/fieldset](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset "The <fieldset> HTML element is used to group several controls as well as labels (<label>) within a web form.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera15+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFieldSetElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFieldSetElement "The HTMLFieldSetElement interface provides special properties and methods (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of <fieldset> elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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

[Categories](dom.html#concept-element-categories){#the-fieldset-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-fieldset-element:flow-content-2}.
:   [Listed](forms.html#category-listed){#the-fieldset-element:category-listed}
    and [autocapitalize-and-autocorrect
    inheriting](forms.html#category-autocapitalize){#the-fieldset-element:category-autocapitalize}
    [form-associated
    element](forms.html#form-associated-element){#the-fieldset-element:form-associated-element}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-fieldset-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-fieldset-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-fieldset-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-fieldset-element:concept-element-content-model}:
:   Optionally, a
    [`legend`{#the-fieldset-element:the-legend-element}](#the-legend-element)
    element, followed by [flow
    content](dom.html#flow-content-2){#the-fieldset-element:flow-content-2-3}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-fieldset-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-fieldset-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-fieldset-element:global-attributes}
:   [`disabled`{#the-fieldset-element:attr-fieldset-disabled}](#attr-fieldset-disabled)
    --- Whether the descendant form controls, except any inside
    [`legend`{#the-fieldset-element:the-legend-element-2}](#the-legend-element),
    are disabled
:   [`form`{#the-fieldset-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    --- Associates the element with a
    [`form`{#the-fieldset-element:the-form-element}](forms.html#the-form-element)
    element
:   [`name`{#the-fieldset-element:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    --- Name of the element to use in the
    [`form.elements`{#the-fieldset-element:dom-form-elements}](forms.html#dom-form-elements)
    API.

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-fieldset-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-fieldset).
:   [For implementers](https://w3c.github.io/html-aam/#el-fieldset).

[DOM interface](dom.html#concept-element-dom){#the-fieldset-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLFieldSetElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute boolean disabled;
      readonly attribute HTMLFormElement? form;
      [CEReactions] attribute DOMString name;

      readonly attribute DOMString type;

      [SameObject] readonly attribute HTMLCollection elements;

      readonly attribute boolean willValidate;
      [SameObject] readonly attribute ValidityState validity;
      readonly attribute DOMString validationMessage;
      boolean checkValidity();
      boolean reportValidity();
      undefined setCustomValidity(DOMString error);
    };
    ```

The
[`fieldset`{#the-fieldset-element:the-fieldset-element}](#the-fieldset-element)
element
[represents](dom.html#represents){#the-fieldset-element:represents} a
set of form controls (or other content) grouped together, optionally
with a caption. The caption is given by the first
[`legend`{#the-fieldset-element:the-legend-element-3}](#the-legend-element)
element that is a child of the
[`fieldset`{#the-fieldset-element:the-fieldset-element-2}](#the-fieldset-element)
element, if any. The remainder of the descendants form the group.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/fieldset#attr-disabled](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset#attr-disabled "The <fieldset> HTML element is used to group several controls as well as labels (<label>) within a web form.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome20+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

The [`disabled`]{#attr-fieldset-disabled .dfn dfn-for="fieldset"
dfn-type="element-attr"} attribute, when specified, causes all the form
control descendants of the
[`fieldset`{#the-fieldset-element:the-fieldset-element-3}](#the-fieldset-element)
element, excluding those that are descendants of the
[`fieldset`{#the-fieldset-element:the-fieldset-element-4}](#the-fieldset-element)
element\'s first
[`legend`{#the-fieldset-element:the-legend-element-4}](#the-legend-element)
element child, if any, to be
[disabled](form-control-infrastructure.html#concept-fe-disabled){#the-fieldset-element:concept-fe-disabled}.

A
[`fieldset`{#the-fieldset-element:the-fieldset-element-5}](#the-fieldset-element)
element is a [disabled fieldset]{#concept-fieldset-disabled .dfn} if it
matches any of the following conditions:

- Its
  [`disabled`{#the-fieldset-element:attr-fieldset-disabled-2}](#attr-fieldset-disabled)
  attribute is specified
- It is a descendant of another
  [`fieldset`{#the-fieldset-element:the-fieldset-element-6}](#the-fieldset-element)
  element whose
  [`disabled`{#the-fieldset-element:attr-fieldset-disabled-3}](#attr-fieldset-disabled)
  attribute is specified, and is *not* a descendant of that
  [`fieldset`{#the-fieldset-element:the-fieldset-element-7}](#the-fieldset-element)
  element\'s first
  [`legend`{#the-fieldset-element:the-legend-element-5}](#the-legend-element)
  element child, if any.

The
[`form`{#the-fieldset-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the
[`fieldset`{#the-fieldset-element:the-fieldset-element-8}](#the-fieldset-element)
element with its [form
owner](form-control-infrastructure.html#form-owner){#the-fieldset-element:form-owner}.
The
[`name`{#the-fieldset-element:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attribute represents the element\'s name.

`fieldset`{.variable}`.`[`type`](#dom-fieldset-type){#dom-fieldset-type-dev}
:   Returns the string \"fieldset\".

`fieldset`{.variable}`.`[`elements`](#dom-fieldset-elements){#dom-fieldset-elements-dev}

:   Returns an
    [`HTMLCollection`{#the-fieldset-element:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
    of the form controls in the element.

The [`disabled`]{#dom-fieldset-disabled .dfn
dfn-for="HTMLFieldSetElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-fieldset-element:reflect}
the content attribute of the same name.

The [`type`]{#dom-fieldset-type .dfn dfn-for="HTMLFieldSetElement"
dfn-type="attribute"} IDL attribute must return the string
\"`fieldset`\".

The [`elements`]{#dom-fieldset-elements .dfn
dfn-for="HTMLFieldSetElement" dfn-type="attribute"} IDL attribute must
return an
[`HTMLCollection`{#the-fieldset-element:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`fieldset`{#the-fieldset-element:the-fieldset-element-9}](#the-fieldset-element)
element, whose filter matches [listed
elements](forms.html#category-listed){#the-fieldset-element:category-listed-2}.

The
[`willValidate`{#the-fieldset-element:dom-cva-willvalidate-2}](form-control-infrastructure.html#dom-cva-willvalidate),
[`validity`{#the-fieldset-element:dom-cva-validity-2}](form-control-infrastructure.html#dom-cva-validity),
and
[`validationMessage`{#the-fieldset-element:dom-cva-validationmessage-2}](form-control-infrastructure.html#dom-cva-validationmessage)
attributes, and the
[`checkValidity()`{#the-fieldset-element:dom-cva-checkvalidity-2}](form-control-infrastructure.html#dom-cva-checkvalidity),
[`reportValidity()`{#the-fieldset-element:dom-cva-reportvalidity-2}](form-control-infrastructure.html#dom-cva-reportvalidity),
and
[`setCustomValidity()`{#the-fieldset-element:dom-cva-setcustomvalidity-2}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
methods, are part of the [constraint validation
API](form-control-infrastructure.html#the-constraint-validation-api){#the-fieldset-element:the-constraint-validation-api}.
The
[`form`{#the-fieldset-element:dom-fae-form-2}](form-control-infrastructure.html#dom-fae-form)
and
[`name`{#the-fieldset-element:dom-fe-name-2}](form-control-infrastructure.html#dom-fe-name)
IDL attributes are part of the element\'s forms API.

::: example
This example shows a
[`fieldset`{#the-fieldset-element:the-fieldset-element-10}](#the-fieldset-element)
element being used to group a set of related controls:

``` html
<fieldset>
 <legend>Display</legend>
 <p><label><input type=radio name=c value=0 checked> Black on White</label>
 <p><label><input type=radio name=c value=1> White on Black</label>
 <p><label><input type=checkbox name=g> Use grayscale</label>
 <p><label>Enhance contrast <input type=range name=e list=contrast min=0 max=100 value=0 step=1></label>
 <datalist id=contrast>
  <option label=Normal value=0>
  <option label=Maximum value=100>
 </datalist>
</fieldset>
```
:::

::: example
The following snippet shows a fieldset with a checkbox in the legend
that controls whether or not the fieldset is enabled. The contents of
the fieldset consist of two required text controls and an optional
year/month control.

``` html
<fieldset name="clubfields" disabled>
 <legend> <label>
  <input type=checkbox name=club onchange="form.clubfields.disabled = !checked">
  Use Club Card
 </label> </legend>
 <p><label>Name on card: <input name=clubname required></label></p>
 <p><label>Card number: <input name=clubnum required pattern="[-0-9]+"></label></p>
 <p><label>Expiry date: <input name=clubexp type=month></label></p>
</fieldset>
```
:::

::: example
You can also nest
[`fieldset`{#the-fieldset-element:the-fieldset-element-11}](#the-fieldset-element)
elements. Here is an example expanding on the previous one that does so:

``` html
<fieldset name="clubfields" disabled>
 <legend> <label>
  <input type=checkbox name=club onchange="form.clubfields.disabled = !checked">
  Use Club Card
 </label> </legend>
 <p><label>Name on card: <input name=clubname required></label></p>
 <fieldset name="numfields">
  <legend> <label>
   <input type=radio checked name=clubtype onchange="form.numfields.disabled = !checked">
   My card has numbers on it
  </label> </legend>
  <p><label>Card number: <input name=clubnum required pattern="[-0-9]+"></label></p>
 </fieldset>
 <fieldset name="letfields" disabled>
  <legend> <label>
   <input type=radio name=clubtype onchange="form.letfields.disabled = !checked">
   My card has letters on it
  </label> </legend>
  <p><label>Card code: <input name=clublet required pattern="[A-Za-z]+"></label></p>
 </fieldset>
</fieldset>
```

In this example, if the outer \"Use Club Card\" checkbox is not checked,
everything inside the outer
[`fieldset`{#the-fieldset-element:the-fieldset-element-12}](#the-fieldset-element),
including the two radio buttons in the legends of the two nested
[`fieldset`{#the-fieldset-element:the-fieldset-element-13}](#the-fieldset-element)s,
will be disabled. However, if the checkbox is checked, then the radio
buttons will both be enabled and will let you select which of the two
inner
[`fieldset`{#the-fieldset-element:the-fieldset-element-14}](#the-fieldset-element)s
is to be enabled.
:::

::: example
This example shows a grouping of controls where the
[`legend`{#the-fieldset-element:the-legend-element-6}](#the-legend-element)
element both labels the grouping, and the nested heading element
surfaces the grouping in the document outline:

``` html
<fieldset>
 <legend> <h2>
  How can we best reach you?
 </h2> </legend>
 <p> <label>
 <input type=radio checked name=contact_pref>
  Phone
 </label> </p>
 <p> <label>
  <input type=radio name=contact_pref>
  Text
 </label> </p>
 <p> <label>
  <input type=radio name=contact_pref>
  Email
 </label> </p>
</fieldset>
```
:::

#### [4.10.16]{.secno} The [`legend`]{.dfn dfn-type="element"} element[](#the-legend-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/legend](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend "The <legend> HTML element represents a caption for the content of its parent <fieldset>.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLLegendElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLegendElement "The HTMLLegendElement is an interface allowing to access properties of the <legend> elements. It inherits properties and methods from the HTMLElement interface.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-legend-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-legend-element:concept-element-contexts}:
:   As the [first
    child](https://dom.spec.whatwg.org/#concept-tree-first-child){#the-legend-element:first-child
    x-internal="first-child"} of a
    [`fieldset`{#the-legend-element:the-fieldset-element}](#the-fieldset-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-legend-element:concept-element-content-model}:
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-legend-element:phrasing-content-2},
    optionally intermixed with [heading
    content](dom.html#heading-content-2){#the-legend-element:heading-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-legend-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-legend-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-legend-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-legend-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-legend).
:   [For implementers](https://w3c.github.io/html-aam/#el-legend).

[DOM interface](dom.html#concept-element-dom){#the-legend-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLLegendElement : HTMLElement {
      [HTMLConstructor] constructor();

      readonly attribute HTMLFormElement? form;

      // also has obsolete members
    };
    ```

The
[`legend`{#the-legend-element:the-legend-element}](#the-legend-element)
element
[represents](dom.html#represents){#the-legend-element:represents} a
caption for the rest of the contents of the
[`legend`{#the-legend-element:the-legend-element-2}](#the-legend-element)
element\'s parent
[`fieldset`{#the-legend-element:the-fieldset-element-2}](#the-fieldset-element)
element, if any.

`legend`{.variable}`.`[`form`](#dom-legend-form){#dom-legend-form-dev}

:   Returns the element\'s
    [`form`{#the-legend-element:the-form-element}](forms.html#the-form-element)
    element, if any, or null otherwise.

The [`form`]{#dom-legend-form .dfn dfn-for="HTMLLegendElement"
dfn-type="attribute"} IDL attribute\'s behavior depends on whether the
[`legend`{#the-legend-element:the-legend-element-3}](#the-legend-element)
element is in a
[`fieldset`{#the-legend-element:the-fieldset-element-3}](#the-fieldset-element)
element or not. If the
[`legend`{#the-legend-element:the-legend-element-4}](#the-legend-element)
has a
[`fieldset`{#the-legend-element:the-fieldset-element-4}](#the-fieldset-element)
element as its parent, then the
[`form`{#the-legend-element:dom-legend-form-2}](#dom-legend-form) IDL
attribute must return the same value as the
[`form`{#the-legend-element:dom-fae-form}](form-control-infrastructure.html#dom-fae-form)
IDL attribute on that
[`fieldset`{#the-legend-element:the-fieldset-element-5}](#the-fieldset-element)
element. Otherwise, it must return null.

[← 4.10.5 The input element](input.html) --- [Table of Contents](./) ---
[4.10.17 Form control infrastructure
→](form-control-infrastructure.html)
