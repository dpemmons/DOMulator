HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.10.6 The button element](form-elements.html) --- [Table of
Contents](./) --- [4.11 Interactive elements
→](interactive-elements.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.10.17]{.secno} Form control
            infrastructure](form-control-infrastructure.html#form-control-infrastructure)
            1.  [[4.10.17.1]{.secno} A form control\'s
                value](form-control-infrastructure.html#a-form-control's-value)
            2.  [[4.10.17.2]{.secno}
                Mutability](form-control-infrastructure.html#mutability)
            3.  [[4.10.17.3]{.secno} Association of controls and
                forms](form-control-infrastructure.html#association-of-controls-and-forms)
        2.  [[4.10.18]{.secno} Attributes common to form
            controls](form-control-infrastructure.html#attributes-common-to-form-controls)
            1.  [[4.10.18.1]{.secno} Naming form controls: the `name`
                attribute](form-control-infrastructure.html#naming-form-controls:-the-name-attribute)
            2.  [[4.10.18.2]{.secno} Submitting element directionality:
                the `dirname`
                attribute](form-control-infrastructure.html#submitting-element-directionality:-the-dirname-attribute)
            3.  [[4.10.18.3]{.secno} Limiting user input length: the
                `maxlength`
                attribute](form-control-infrastructure.html#limiting-user-input-length:-the-maxlength-attribute)
            4.  [[4.10.18.4]{.secno} Setting minimum input length
                requirements: the `minlength`
                attribute](form-control-infrastructure.html#setting-minimum-input-length-requirements:-the-minlength-attribute)
            5.  [[4.10.18.5]{.secno} Enabling and disabling form
                controls: the `disabled`
                attribute](form-control-infrastructure.html#enabling-and-disabling-form-controls:-the-disabled-attribute)
            6.  [[4.10.18.6]{.secno} Form submission
                attributes](form-control-infrastructure.html#form-submission-attributes)
            7.  [[4.10.18.7]{.secno}
                Autofill](form-control-infrastructure.html#autofill)
                1.  [[4.10.18.7.1]{.secno} Autofilling form controls:
                    the `autocomplete`
                    attribute](form-control-infrastructure.html#autofilling-form-controls:-the-autocomplete-attribute)
                2.  [[4.10.18.7.2]{.secno} Processing
                    model](form-control-infrastructure.html#autofill-processing-model)
        3.  [[4.10.19]{.secno} APIs for the text control
            selections](form-control-infrastructure.html#textFieldSelection)
        4.  [[4.10.20]{.secno}
            Constraints](form-control-infrastructure.html#constraints)
            1.  [[4.10.20.1]{.secno}
                Definitions](form-control-infrastructure.html#definitions)
            2.  [[4.10.20.2]{.secno} Constraint
                validation](form-control-infrastructure.html#constraint-validation)
            3.  [[4.10.20.3]{.secno} The constraint validation
                API](form-control-infrastructure.html#the-constraint-validation-api)
            4.  [[4.10.20.4]{.secno}
                Security](form-control-infrastructure.html#security-forms)
        5.  [[4.10.21]{.secno} Form
            submission](form-control-infrastructure.html#form-submission-2)
            1.  [[4.10.21.1]{.secno}
                Introduction](form-control-infrastructure.html#introduction-5)
            2.  [[4.10.21.2]{.secno} Implicit
                submission](form-control-infrastructure.html#implicit-submission)
            3.  [[4.10.21.3]{.secno} Form submission
                algorithm](form-control-infrastructure.html#form-submission-algorithm)
            4.  [[4.10.21.4]{.secno} Constructing the entry
                list](form-control-infrastructure.html#constructing-form-data-set)
            5.  [[4.10.21.5]{.secno} Selecting a form submission
                encoding](form-control-infrastructure.html#selecting-a-form-submission-encoding)
            6.  [[4.10.21.6]{.secno} Converting an entry list to a list
                of name-value
                pairs](form-control-infrastructure.html#converting-an-entry-list-to-a-list-of-name-value-pairs)
            7.  [[4.10.21.7]{.secno} URL-encoded form
                data](form-control-infrastructure.html#url-encoded-form-data)
            8.  [[4.10.21.8]{.secno} Multipart form
                data](form-control-infrastructure.html#multipart-form-data)
            9.  [[4.10.21.9]{.secno} Plain text form
                data](form-control-infrastructure.html#plain-text-form-data)
            10. [[4.10.21.10]{.secno} The `SubmitEvent`
                interface](form-control-infrastructure.html#the-submitevent-interface)
            11. [[4.10.21.11]{.secno} The `FormDataEvent`
                interface](form-control-infrastructure.html#the-formdataevent-interface)
        6.  [[4.10.22]{.secno} Resetting a
            form](form-control-infrastructure.html#resetting-a-form)
    :::

#### [4.10.17]{.secno} Form control infrastructure[](#form-control-infrastructure){.self-link}

##### [4.10.17.1]{.secno} A form control\'s value[](#a-form-control's-value){.self-link} {#a-form-control's-value}

Most form controls have a [value]{#concept-fe-value .dfn} and a
[checkedness]{#concept-fe-checked .dfn}. (The latter is only used by
[`input`{#a-form-control's-value:the-input-element}](input.html#the-input-element)
elements.) These are used to describe how the user interacts with the
control.

A control\'s
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value} is
its internal state. As such, it might not match the user\'s current
input.

For instance, if a user enters the word \"[three]{.kbd}\" into [a
numeric
field](input.html#number-state-(type=number)){#a-form-control's-value:number-state-(type=number)}
that expects digits, the user\'s input would be the string \"three\" but
the control\'s
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-2}
would remain unchanged. Or, if a user enters the email address
\"[  awesome@example.com]{.kbd}\" (with leading whitespace) into [an
email
field](input.html#email-state-(type=email)){#a-form-control's-value:email-state-(type=email)},
the user\'s input would be the string \"  awesome@example.com\" but the
browser\'s UI for email fields might translate that into a
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-3}
of \"`awesome@example.com`\" (without the leading whitespace).

[]{#concept-textarea-dirty}[`input`{#a-form-control's-value:the-input-element-2}](input.html#the-input-element)
and
[`textarea`{#a-form-control's-value:the-textarea-element}](form-elements.html#the-textarea-element)
elements have a [dirty value flag]{#concept-fe-dirty .dfn}. This is used
to track the interaction between the
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-4}
and default value. If it is false,
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-5}
mirrors the default value. If it is true, the default value is ignored.

Some form controls also have an [optional
value]{#concept-fe-optional-value .dfn} this largely mirrors the
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-6}
but doesn\'t normalize to an empty string. [This ought to be used
sparingly, you generally want
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-7}]{.note}.

[`input`{#a-form-control's-value:the-input-element-3}](input.html#the-input-element),
[`textarea`{#a-form-control's-value:the-textarea-element-2}](form-elements.html#the-textarea-element),
and
[`select`{#a-form-control's-value:the-select-element}](form-elements.html#the-select-element)
elements have a [user validity]{#user-validity .dfn} boolean. It is
initially set to false.

To define the behavior of constraint validation in the face of the
[`input`{#a-form-control's-value:the-input-element-4}](input.html#the-input-element)
element\'s
[`multiple`{#a-form-control's-value:attr-input-multiple}](input.html#attr-input-multiple)
attribute,
[`input`{#a-form-control's-value:the-input-element-5}](input.html#the-input-element)
elements can also have separately defined [value*s*]{#concept-fe-values
.dfn}.

To define the behavior of the
[`maxlength`{#a-form-control's-value:attr-fe-maxlength}](#attr-fe-maxlength)
and
[`minlength`{#a-form-control's-value:attr-fe-minlength}](#attr-fe-minlength)
attributes, as well as other APIs specific to the
[`textarea`{#a-form-control's-value:the-textarea-element-3}](form-elements.html#the-textarea-element)
element, all form control with a
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-8}
also have an algorithm for obtaining an
[]{#concept-textarea-api-value}[API value]{#concept-fe-api-value .dfn}.
By default this algorithm is to simply return the control\'s
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-9}.

The
[`select`{#a-form-control's-value:the-select-element-2}](form-elements.html#the-select-element)
element does not have a
[value](#concept-fe-value){#a-form-control's-value:concept-fe-value-10};
the
[selectedness](form-elements.html#concept-option-selectedness){#a-form-control's-value:concept-option-selectedness}
of its
[`option`{#a-form-control's-value:the-option-element}](form-elements.html#the-option-element)
elements is what is used instead.

##### [4.10.17.2]{.secno} Mutability[](#mutability){.self-link}

A form control can be designated as [*mutable*]{#concept-fe-mutable
.dfn}.

This determines (by means of definitions and requirements in this
specification that rely on whether an element is so designated) whether
or not the user can modify the
[value](#concept-fe-value){#mutability:concept-fe-value} or
[checkedness](#concept-fe-checked){#mutability:concept-fe-checked} of a
form control, or whether or not a control can be automatically
prefilled.

##### [4.10.17.3]{.secno} Association of controls and forms[](#association-of-controls-and-forms){.self-link}

A [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element}
can have a relationship with a
[`form`{#association-of-controls-and-forms:the-form-element}](forms.html#the-form-element)
element, which is called the element\'s [form owner]{#form-owner .dfn
dfn-for="button,fieldset,input,object,output,select,textarea,img,form-associated custom element"
export=""}. If a [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-2}
is not associated with a
[`form`{#association-of-controls-and-forms:the-form-element-2}](forms.html#the-form-element)
element, its [form
owner](#form-owner){#association-of-controls-and-forms:form-owner} is
said to be null.

A [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-3}
has an associated [parser inserted flag]{#parser-inserted-flag .dfn}.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#form](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#form "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

:::: feature
[Attributes#attr-form](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes#attr-form "Elements in HTML have attributes; these are additional values that configure the elements or adjust their behavior in various ways to meet the criteria the users want.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari≤4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera≤12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer≤6+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android≤12.1+]{.opera_android .yes}
:::
::::
:::::::

A [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-4}
is, by default, associated with its nearest ancestor
[`form`{#association-of-controls-and-forms:the-form-element-3}](forms.html#the-form-element)
element (as described below), but, if it is
[listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed},
may have a [`form`]{#attr-fae-form .dfn
dfn-for="button,fieldset,input,object,output,select,textarea"
dfn-type="element-attr"} attribute specified to override this.

This feature allows authors to work around the lack of support for
nested
[`form`{#association-of-controls-and-forms:the-form-element-4}](forms.html#the-form-element)
elements.

If a
[listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed-2}
[form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-5}
has a
[`form`{#association-of-controls-and-forms:attr-fae-form}](#attr-fae-form)
attribute specified, then that attribute\'s value must be the
[ID](https://dom.spec.whatwg.org/#concept-id){#association-of-controls-and-forms:concept-id
x-internal="concept-id"} of a
[`form`{#association-of-controls-and-forms:the-form-element-5}](forms.html#the-form-element)
element in the element\'s
[tree](https://dom.spec.whatwg.org/#concept-tree){#association-of-controls-and-forms:tree
x-internal="tree"}.

The rules in this section are complicated by the fact that although
conforming documents or
[trees](https://dom.spec.whatwg.org/#concept-tree){#association-of-controls-and-forms:tree-2
x-internal="tree"} will never contain nested
[`form`{#association-of-controls-and-forms:the-form-element-6}](forms.html#the-form-element)
elements, it is quite possible (e.g., using a script that performs DOM
manipulation) to generate
[trees](https://dom.spec.whatwg.org/#concept-tree){#association-of-controls-and-forms:tree-3
x-internal="tree"} that have such nested elements. They are also
complicated by rules in the HTML parser that, for historical reasons,
can result in a [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-6}
being associated with a
[`form`{#association-of-controls-and-forms:the-form-element-7}](forms.html#the-form-element)
element that is not its ancestor.

When a [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-7}
is created, its [form
owner](#form-owner){#association-of-controls-and-forms:form-owner-2}
must be initialized to null (no owner).

When a [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-8}
is to be [associated]{#concept-form-association .dfn} with a form, its
[form
owner](#form-owner){#association-of-controls-and-forms:form-owner-3}
must be set to that form.

When a
[listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed-3}
[form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-9}\'s
[`form`{#association-of-controls-and-forms:attr-fae-form-2}](#attr-fae-form)
attribute is set, changed, or removed, then the user agent must [reset
the form
owner](#reset-the-form-owner){#association-of-controls-and-forms:reset-the-form-owner}
of that element.

When a
[listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed-4}
[form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-10}
has a
[`form`{#association-of-controls-and-forms:attr-fae-form-3}](#attr-fae-form)
attribute and the
[ID](https://dom.spec.whatwg.org/#concept-id){#association-of-controls-and-forms:concept-id-2
x-internal="concept-id"} of any of the elements in the
[tree](https://dom.spec.whatwg.org/#concept-tree){#association-of-controls-and-forms:tree-4
x-internal="tree"} changes, then the user agent must [reset the form
owner](#reset-the-form-owner){#association-of-controls-and-forms:reset-the-form-owner-2}
of that [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-11}.

When a
[listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed-5}
[form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-12}
has a
[`form`{#association-of-controls-and-forms:attr-fae-form-4}](#attr-fae-form)
attribute and an element with an
[ID](https://dom.spec.whatwg.org/#concept-id){#association-of-controls-and-forms:concept-id-3
x-internal="concept-id"} is [inserted
into](infrastructure.html#insert-an-element-into-a-document){#association-of-controls-and-forms:insert-an-element-into-a-document}
or [removed
from](infrastructure.html#remove-an-element-from-a-document){#association-of-controls-and-forms:remove-an-element-from-a-document}
the
[`Document`{#association-of-controls-and-forms:document}](dom.html#document),
or its [HTML element moving
steps](infrastructure.html#html-element-moving-steps){#association-of-controls-and-forms:html-element-moving-steps}
are run, then the user agent must [reset the form
owner](#reset-the-form-owner){#association-of-controls-and-forms:reset-the-form-owner-3}
of that [form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-13}.

The form owner is also reset by the [HTML element insertion
steps](infrastructure.html#html-element-insertion-steps){#association-of-controls-and-forms:html-element-insertion-steps},
[HTML element removing
steps](infrastructure.html#html-element-removing-steps){#association-of-controls-and-forms:html-element-removing-steps},
and [HTML element moving
steps](infrastructure.html#html-element-moving-steps){#association-of-controls-and-forms:html-element-moving-steps-2}.

To [reset the form owner]{#reset-the-form-owner .dfn} of a
[form-associated
element](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-14}
`element`{.variable}:

1.  Unset `element`{.variable}\'s [parser inserted
    flag](#parser-inserted-flag){#association-of-controls-and-forms:parser-inserted-flag}.

2.  If all of the following are true:

    - `element`{.variable}\'s [form
      owner](#form-owner){#association-of-controls-and-forms:form-owner-4}
      is not null;

    - `element`{.variable} is not
      [listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed-6}
      or its
      [`form`{#association-of-controls-and-forms:attr-fae-form-5}](#attr-fae-form)
      content attribute is not present; and

    - `element`{.variable}\'s [form
      owner](#form-owner){#association-of-controls-and-forms:form-owner-5}
      is its nearest
      [`form`{#association-of-controls-and-forms:the-form-element-8}](forms.html#the-form-element)
      element ancestor after the change to the ancestor chain,

    then return.

3.  Set `element`{.variable}\'s [form
    owner](#form-owner){#association-of-controls-and-forms:form-owner-6}
    to null.

4.  If `element`{.variable} is
    [listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed-7},
    has a
    [`form`{#association-of-controls-and-forms:attr-fae-form-6}](#attr-fae-form)
    content attribute, and is
    [connected](https://dom.spec.whatwg.org/#connected){#association-of-controls-and-forms:connected
    x-internal="connected"}, then:

    1.  If the first element in `element`{.variable}\'s
        [tree](https://dom.spec.whatwg.org/#concept-tree){#association-of-controls-and-forms:tree-5
        x-internal="tree"}, in [tree
        order](https://dom.spec.whatwg.org/#concept-tree-order){#association-of-controls-and-forms:tree-order
        x-internal="tree-order"}, to have an
        [ID](https://dom.spec.whatwg.org/#concept-id){#association-of-controls-and-forms:concept-id-4
        x-internal="concept-id"} that is [identical
        to](https://infra.spec.whatwg.org/#string-is){#association-of-controls-and-forms:identical-to
        x-internal="identical-to"} `element`{.variable}\'s
        [`form`{#association-of-controls-and-forms:attr-fae-form-7}](#attr-fae-form)
        content attribute\'s value, is a
        [`form`{#association-of-controls-and-forms:the-form-element-9}](forms.html#the-form-element)
        element, then
        [associate](#concept-form-association){#association-of-controls-and-forms:concept-form-association}
        the `element`{.variable} with that
        [`form`{#association-of-controls-and-forms:the-form-element-10}](forms.html#the-form-element)
        element.

5.  Otherwise, if `element`{.variable} has an ancestor
    [`form`{#association-of-controls-and-forms:the-form-element-11}](forms.html#the-form-element)
    element, then
    [associate](#concept-form-association){#association-of-controls-and-forms:concept-form-association-2}
    `element`{.variable} with the nearest such ancestor
    [`form`{#association-of-controls-and-forms:the-form-element-12}](forms.html#the-form-element)
    element.

::: example
In the following non-conforming snippet

``` bad
...
 <form id="a">
  <div id="b"></div>
 </form>
 <script>
  document.getElementById('b').innerHTML =
     '<table><tr><td></form><form id="c"><input id="d"></table>' +
     '<input id="e">';
 </script>
...
```

the [form
owner](#form-owner){#association-of-controls-and-forms:form-owner-7} of
\"d\" would be the inner nested form \"c\", while the [form
owner](#form-owner){#association-of-controls-and-forms:form-owner-8} of
\"e\" would be the outer form \"a\".

This happens as follows: First, the \"e\" node gets associated with
\"c\" in the [HTML
parser](parsing.html#html-parser){#association-of-controls-and-forms:html-parser}.
Then, the
[`innerHTML`{#association-of-controls-and-forms:dom-element-innerhtml}](dynamic-markup-insertion.html#dom-element-innerhtml)
algorithm moves the nodes from the temporary document to the \"b\"
element. At this point, the nodes see their ancestor chain change, and
thus all the \"magic\" associations done by the parser are reset to
normal ancestor associations.

This example is a non-conforming document, though, as it is a violation
of the content models to nest
[`form`{#association-of-controls-and-forms:the-form-element-13}](forms.html#the-form-element)
elements, and there is a [parse
error](parsing.html#parse-errors){#association-of-controls-and-forms:parse-errors}
for the `</form>` tag.
:::

`element`{.variable}`.`[`form`](#dom-fae-form){#dom-fae-form-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/form](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/form "The form read-only property of the HTMLObjectElement interface returns a HTMLFormElement representing the object element's form owner, or null if there isn't one.")

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

:::: feature
[HTMLSelectElement/form](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/form "The HTMLSelectElement.form read-only property returns a HTMLFormElement representing the form that this element is associated with. If the element is not associated with a <form> element, then it returns null.")

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
:::::::

Returns the element\'s [form
owner](#form-owner){#association-of-controls-and-forms:form-owner-9}.

Returns null if there isn\'t one.

[Listed](forms.html#category-listed){#association-of-controls-and-forms:category-listed-8}
[form-associated
elements](forms.html#form-associated-element){#association-of-controls-and-forms:form-associated-element-15}
except for [form-associated custom
elements](custom-elements.html#form-associated-custom-element){#association-of-controls-and-forms:form-associated-custom-element}
have a [`form`]{#dom-fae-form .dfn
dfn-for="HTMLObjectElement,HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="attribute"} IDL attribute, which, on getting, must return the
element\'s [form
owner](#form-owner){#association-of-controls-and-forms:form-owner-10},
or null if there isn\'t one.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/form](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/form "The form read-only property of the ElementInternals interface returns the HTMLFormElement associated with this element.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
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

[Form-associated custom
elements](custom-elements.html#form-associated-custom-element){#association-of-controls-and-forms:form-associated-custom-element-2}
don\'t have
[`form`{#association-of-controls-and-forms:dom-fae-form}](#dom-fae-form)
IDL attribute. Instead, their
[`ElementInternals`{#association-of-controls-and-forms:elementinternals}](custom-elements.html#elementinternals)
object has a [`form`]{#dom-elementinternals-form .dfn
dfn-for="ElementInternals" dfn-type="attribute"} IDL attribute. On
getting, it must throw a
[\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#association-of-controls-and-forms:notsupportederror
x-internal="notsupportederror"}
[`DOMException`{#association-of-controls-and-forms:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the [target
element](custom-elements.html#internals-target){#association-of-controls-and-forms:internals-target}
is not a [form-associated custom
element](custom-elements.html#form-associated-custom-element){#association-of-controls-and-forms:form-associated-custom-element-3}.
Otherwise, it must return the element\'s [form
owner](#form-owner){#association-of-controls-and-forms:form-owner-11},
or null if there isn\'t one.

#### [4.10.18]{.secno} Attributes common to form controls[](#attributes-common-to-form-controls){.self-link}

##### [4.10.18.1]{.secno} Naming form controls: the [`name`{#naming-form-controls:-the-name-attribute:attr-fe-name}](#attr-fe-name) attribute[](#naming-form-controls:-the-name-attribute){.self-link} {#naming-form-controls:-the-name-attribute}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#name](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#name "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

The [`name`]{#attr-fe-name .dfn
dfn-for="button,fieldset,input,output,select,textarea"
dfn-type="element-attr"} content attribute gives the name of the form
control, as used in [form
submission](#form-submission-2){#naming-form-controls:-the-name-attribute:form-submission-2}
and in the
[`form`{#naming-form-controls:-the-name-attribute:the-form-element}](forms.html#the-form-element)
element\'s
[`elements`{#naming-form-controls:-the-name-attribute:dom-form-elements}](forms.html#dom-form-elements)
object. If the attribute is specified, its value must not be the empty
string or `isindex`.

A number of user agents historically implemented special support for
first-in-form text controls with the name `isindex`, and this
specification previously defined related user agent requirements for it.
However, some user agents subsequently dropped that special support, and
the related requirements were removed from this specification. So, to
avoid problematic reinterpretations in legacy user agents, the name
`isindex` is no longer allowed.

Other than `isindex`, any non-empty value for
[`name`{#naming-form-controls:-the-name-attribute:attr-form-name}](forms.html#attr-form-name)
is allowed. An [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#naming-form-controls:-the-name-attribute:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the name
[`_charset_`]{#attr-fe-name-charset .dfn} is special: if used as the
name of a
[Hidden](input.html#hidden-state-(type=hidden)){#naming-form-controls:-the-name-attribute:hidden-state-(type=hidden)}
control with no
[`value`{#naming-form-controls:-the-name-attribute:attr-input-value}](input.html#attr-input-value)
attribute, then during submission the
[`value`{#naming-form-controls:-the-name-attribute:attr-input-value-2}](input.html#attr-input-value)
attribute is automatically given a value consisting of the submission
character encoding.

The [`name`]{#dom-fe-name .dfn
dfn-for="HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#naming-form-controls:-the-name-attribute:reflect}
the
[`name`{#naming-form-controls:-the-name-attribute:attr-fe-name-2}](#attr-fe-name)
content attribute.

::: note
DOM clobbering is a common cause of security issues. Avoid using the
names of built-in form properties with the
[`name`{#naming-form-controls:-the-name-attribute:attr-fe-name-3}](#attr-fe-name)
content attribute.

In this example, the
[`input`{#naming-form-controls:-the-name-attribute:the-input-element}](input.html#the-input-element)
element overrides the built-in
[`method`{#naming-form-controls:-the-name-attribute:attr-fs-method}](#attr-fs-method)
property:

``` js
let form = document.createElement("form");
let input = document.createElement("input");
form.appendChild(input);

form.method;           // => "get"
input.name = "method"; // DOM clobbering occurs here
form.method === input; // => true
```

Since the input name takes precedence over built-in form properties, the
JavaScript reference `form.method` will point to the
[`input`{#naming-form-controls:-the-name-attribute:the-input-element-2}](input.html#the-input-element)
element named \"method\" instead of the built-in
[`method`{#naming-form-controls:-the-name-attribute:attr-fs-method-2}](#attr-fs-method)
property.
:::

##### [4.10.18.2]{.secno} Submitting element directionality: the [`dirname`{#submitting-element-directionality:-the-dirname-attribute:attr-fe-dirname}](#attr-fe-dirname) attribute[](#submitting-element-directionality:-the-dirname-attribute){.self-link} {#submitting-element-directionality:-the-dirname-attribute}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#dirname](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#dirname "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox116+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome17+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`dirname`]{#attr-fe-dirname .dfn dfn-for="input,textarea"
dfn-type="element-attr"} attribute on a form control element enables the
submission of [the
directionality](dom.html#the-directionality){#submitting-element-directionality:-the-dirname-attribute:the-directionality}
of the element, and gives the name of the control that contains this
value during [form
submission](#form-submission-2){#submitting-element-directionality:-the-dirname-attribute:form-submission-2}.
If such an attribute is specified, its value must not be the empty
string.

::: example
In this example, a form contains a text control and a submission button:

``` html
<form action="addcomment.cgi" method=post>
 <p><label>Comment: <input type=text name="comment" dirname="comment.dir" required></label></p>
 <p><button name="mode" type=submit value="add">Post Comment</button></p>
</form>
```

When the user submits the form, the user agent includes three fields,
one called \"comment\", one called \"comment.dir\", and one called
\"mode\"; so if the user types \"Hello\", the submission body might be
something like:

    comment=Hello&comment.dir=ltr&mode=add

If the user manually switches to a right-to-left writing direction and
enters \"[مرحبا]{dir="rtl" lang="ar"}\", the submission body might be
something like:

    comment=%D9%85%D8%B1%D8%AD%D8%A8%D8%A7&comment.dir=rtl&mode=add
:::

##### [4.10.18.3]{.secno} Limiting user input length: the [`maxlength`{#limiting-user-input-length:-the-maxlength-attribute:attr-fe-maxlength}](#attr-fe-maxlength) attribute[](#limiting-user-input-length:-the-maxlength-attribute){.self-link} {#limiting-user-input-length:-the-maxlength-attribute}

A [form control `maxlength` attribute]{#attr-fe-maxlength .dfn
dfn-for="input,textarea" dfn-type="element-attr" lt="maxlength"},
controlled by the [dirty value
flag](#concept-fe-dirty){#limiting-user-input-length:-the-maxlength-attribute:concept-fe-dirty},
declares a limit on the number of characters a user can input. The
number of characters is measured using
[length](https://infra.spec.whatwg.org/#string-length){#limiting-user-input-length:-the-maxlength-attribute:length
x-internal="length"} and, in the case of
[`textarea`{#limiting-user-input-length:-the-maxlength-attribute:the-textarea-element}](form-elements.html#the-textarea-element)
elements, with all newlines normalized to a single character (as opposed
to CRLF pairs).

If an element has its [form control `maxlength`
attribute](#attr-fe-maxlength){#limiting-user-input-length:-the-maxlength-attribute:attr-fe-maxlength-2}
specified, the attribute\'s value must be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#limiting-user-input-length:-the-maxlength-attribute:valid-non-negative-integer}.
If the attribute is specified and applying the [rules for parsing
non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#limiting-user-input-length:-the-maxlength-attribute:rules-for-parsing-non-negative-integers}
to its value results in a number, then that number is the element\'s
[maximum allowed value length]{#maximum-allowed-value-length .dfn}. If
the attribute is omitted or parsing its value results in an error, then
there is no [maximum allowed value
length](#maximum-allowed-value-length){#limiting-user-input-length:-the-maxlength-attribute:maximum-allowed-value-length}.

**Constraint validation**: If an element has a [maximum allowed value
length](#maximum-allowed-value-length){#limiting-user-input-length:-the-maxlength-attribute:maximum-allowed-value-length-2},
its [dirty value
flag](#concept-fe-dirty){#limiting-user-input-length:-the-maxlength-attribute:concept-fe-dirty-2}
is true, its
[value](#concept-fe-value){#limiting-user-input-length:-the-maxlength-attribute:concept-fe-value}
was last changed by a user edit (as opposed to a change made by a
script), and the
[length](https://infra.spec.whatwg.org/#string-length){#limiting-user-input-length:-the-maxlength-attribute:length-2
x-internal="length"} of the element\'s [API
value](#concept-fe-api-value){#limiting-user-input-length:-the-maxlength-attribute:concept-fe-api-value}
is greater than the element\'s [maximum allowed value
length](#maximum-allowed-value-length){#limiting-user-input-length:-the-maxlength-attribute:maximum-allowed-value-length-3},
then the element is [suffering from being too
long](#suffering-from-being-too-long){#limiting-user-input-length:-the-maxlength-attribute:suffering-from-being-too-long}.

User agents may prevent the user from causing the element\'s [API
value](#concept-fe-api-value){#limiting-user-input-length:-the-maxlength-attribute:concept-fe-api-value-2}
to be set to a value whose
[length](https://infra.spec.whatwg.org/#string-length){#limiting-user-input-length:-the-maxlength-attribute:length-3
x-internal="length"} is greater than the element\'s [maximum allowed
value
length](#maximum-allowed-value-length){#limiting-user-input-length:-the-maxlength-attribute:maximum-allowed-value-length-4}.

In the case of
[`textarea`{#limiting-user-input-length:-the-maxlength-attribute:the-textarea-element-2}](form-elements.html#the-textarea-element)
elements, the [API
value](#concept-fe-api-value){#limiting-user-input-length:-the-maxlength-attribute:concept-fe-api-value-3}
and
[value](#concept-fe-value){#limiting-user-input-length:-the-maxlength-attribute:concept-fe-value-2}
differ. In particular, [newline
normalization](https://infra.spec.whatwg.org/#normalize-newlines){#limiting-user-input-length:-the-maxlength-attribute:normalize-newlines
x-internal="normalize-newlines"} is applied before the [maximum allowed
value
length](#maximum-allowed-value-length){#limiting-user-input-length:-the-maxlength-attribute:maximum-allowed-value-length-5}
is checked (whereas the [textarea wrapping
transformation](form-elements.html#textarea-wrapping-transformation){#limiting-user-input-length:-the-maxlength-attribute:textarea-wrapping-transformation}
is not applied).

##### [4.10.18.4]{.secno} Setting minimum input length requirements: the [`minlength`{#setting-minimum-input-length-requirements:-the-minlength-attribute:attr-fe-minlength}](#attr-fe-minlength) attribute[](#setting-minimum-input-length-requirements:-the-minlength-attribute){.self-link} {#setting-minimum-input-length-requirements:-the-minlength-attribute}

A [form control `minlength` attribute]{#attr-fe-minlength .dfn
dfn-for="input,textarea" dfn-type="element-attr" lt="minlength"},
controlled by the [dirty value
flag](#concept-fe-dirty){#setting-minimum-input-length-requirements:-the-minlength-attribute:concept-fe-dirty},
declares a lower bound on the number of characters a user can input. The
\"number of characters\" is measured using
[length](https://infra.spec.whatwg.org/#string-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:length
x-internal="length"} and, in the case of
[`textarea`{#setting-minimum-input-length-requirements:-the-minlength-attribute:the-textarea-element}](form-elements.html#the-textarea-element)
elements, with all newlines normalized to a single character (as opposed
to CRLF pairs).

The
[`minlength`{#setting-minimum-input-length-requirements:-the-minlength-attribute:attr-fe-minlength-2}](#attr-fe-minlength)
attribute does not imply the `required` attribute. If the form control
has no `required` attribute, then the value can still be omitted; the
[`minlength`{#setting-minimum-input-length-requirements:-the-minlength-attribute:attr-fe-minlength-3}](#attr-fe-minlength)
attribute only kicks in once the user has entered a value at all. If the
empty string is not allowed, then the `required` attribute also needs to
be set.

If an element has its [form control `minlength`
attribute](#attr-fe-minlength){#setting-minimum-input-length-requirements:-the-minlength-attribute:attr-fe-minlength-4}
specified, the attribute\'s value must be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#setting-minimum-input-length-requirements:-the-minlength-attribute:valid-non-negative-integer}.
If the attribute is specified and applying the [rules for parsing
non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#setting-minimum-input-length-requirements:-the-minlength-attribute:rules-for-parsing-non-negative-integers}
to its value results in a number, then that number is the element\'s
[minimum allowed value length]{#minimum-allowed-value-length .dfn}. If
the attribute is omitted or parsing its value results in an error, then
there is no [minimum allowed value
length](#minimum-allowed-value-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:minimum-allowed-value-length}.

If an element has both a [maximum allowed value
length](#maximum-allowed-value-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:maximum-allowed-value-length}
and a [minimum allowed value
length](#minimum-allowed-value-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:minimum-allowed-value-length-2},
the [minimum allowed value
length](#minimum-allowed-value-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:minimum-allowed-value-length-3}
must be smaller than or equal to the [maximum allowed value
length](#maximum-allowed-value-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:maximum-allowed-value-length-2}.

**Constraint validation**: If an element has a [minimum allowed value
length](#minimum-allowed-value-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:minimum-allowed-value-length-4},
its [dirty value
flag](#concept-fe-dirty){#setting-minimum-input-length-requirements:-the-minlength-attribute:concept-fe-dirty-2}
is true, its
[value](#concept-fe-value){#setting-minimum-input-length-requirements:-the-minlength-attribute:concept-fe-value}
was last changed by a user edit (as opposed to a change made by a
script), its
[value](#concept-fe-value){#setting-minimum-input-length-requirements:-the-minlength-attribute:concept-fe-value-2}
is not the empty string, and the
[length](https://infra.spec.whatwg.org/#string-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:length-2
x-internal="length"} of the element\'s [API
value](#concept-fe-api-value){#setting-minimum-input-length-requirements:-the-minlength-attribute:concept-fe-api-value}
is less than the element\'s [minimum allowed value
length](#minimum-allowed-value-length){#setting-minimum-input-length-requirements:-the-minlength-attribute:minimum-allowed-value-length-5},
then the element is [suffering from being too
short](#suffering-from-being-too-short){#setting-minimum-input-length-requirements:-the-minlength-attribute:suffering-from-being-too-short}.

::: example
In this example, there are four text controls. The first is required,
and has to be at least 5 characters long. The other three are optional,
but if the user fills one in, the user has to enter at least 10
characters.

``` html
<form action="/events/menu.cgi" method="post">
 <p><label>Name of Event: <input required minlength=5 maxlength=50 name=event></label></p>
 <p><label>Describe what you would like for breakfast, if anything:
    <textarea name="breakfast" minlength="10"></textarea></label></p>
 <p><label>Describe what you would like for lunch, if anything:
    <textarea name="lunch" minlength="10"></textarea></label></p>
 <p><label>Describe what you would like for dinner, if anything:
    <textarea name="dinner" minlength="10"></textarea></label></p>
 <p><input type=submit value="Submit Request"></p>
</form>
```
:::

##### [4.10.18.5]{.secno} Enabling and disabling form controls: the [`disabled`{#enabling-and-disabling-form-controls:-the-disabled-attribute:attr-fe-disabled}](#attr-fe-disabled) attribute[](#enabling-and-disabling-form-controls:-the-disabled-attribute){.self-link} {#enabling-and-disabling-form-controls:-the-disabled-attribute}

::::::::::::: {.mdn-anno .wrapped .before}
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

:::: feature
[Attributes/disabled](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled "The Boolean disabled attribute, when present, makes the element not mutable, focusable, or even submitted with the form. The user can neither edit nor focus on the control, nor its form control descendants.")

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

:::: feature
[Attributes/disabled](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled "The Boolean disabled attribute, when present, makes the element not mutable, focusable, or even submitted with the form. The user can neither edit nor focus on the control, nor its form control descendants.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

:::: feature
[Attributes/disabled](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled "The Boolean disabled attribute, when present, makes the element not mutable, focusable, or even submitted with the form. The user can neither edit nor focus on the control, nor its form control descendants.")

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

:::: feature
[Attributes/disabled](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled "The Boolean disabled attribute, when present, makes the element not mutable, focusable, or even submitted with the form. The user can neither edit nor focus on the control, nor its form control descendants.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
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
:::::::::::::

The [`disabled`]{#attr-fe-disabled .dfn
dfn-for="button,fieldset,input,object,output,select,textarea"
dfn-type="element-attr"} content attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#enabling-and-disabling-form-controls:-the-disabled-attribute:boolean-attribute}.

The
[`disabled`{#enabling-and-disabling-form-controls:-the-disabled-attribute:attr-option-disabled}](form-elements.html#attr-option-disabled)
attribute for
[`option`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-option-element}](form-elements.html#the-option-element)
elements and the
[`disabled`{#enabling-and-disabling-form-controls:-the-disabled-attribute:attr-optgroup-disabled}](form-elements.html#attr-optgroup-disabled)
attribute for
[`optgroup`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-optgroup-element}](form-elements.html#the-optgroup-element)
elements are defined separately.

A form control is [disabled]{#concept-fe-disabled .dfn} if any of the
following are true:

- the element is a
  [`button`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-button-element}](form-elements.html#the-button-element),
  [`input`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-input-element}](input.html#the-input-element),
  [`select`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-select-element}](form-elements.html#the-select-element),
  [`textarea`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-textarea-element}](form-elements.html#the-textarea-element),
  or [form-associated custom
  element](custom-elements.html#form-associated-custom-element){#enabling-and-disabling-form-controls:-the-disabled-attribute:form-associated-custom-element},
  and the
  [`disabled`{#enabling-and-disabling-form-controls:-the-disabled-attribute:attr-fe-disabled-2}](#attr-fe-disabled)
  attribute is specified on this element (regardless of its value); or

- the element is a descendant of a
  [`fieldset`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-fieldset-element}](form-elements.html#the-fieldset-element)
  element whose
  [`disabled`{#enabling-and-disabling-form-controls:-the-disabled-attribute:attr-fieldset-disabled}](form-elements.html#attr-fieldset-disabled)
  attribute is specified, and is *not* a descendant of that
  [`fieldset`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-fieldset-element-2}](form-elements.html#the-fieldset-element)
  element\'s first
  [`legend`{#enabling-and-disabling-form-controls:-the-disabled-attribute:the-legend-element}](form-elements.html#the-legend-element)
  element child, if any.

A form control that is
[disabled](#concept-fe-disabled){#enabling-and-disabling-form-controls:-the-disabled-attribute:concept-fe-disabled}
must prevent any
[`click`{#enabling-and-disabling-form-controls:-the-disabled-attribute:event-click}](https://w3c.github.io/uievents/#event-type-click){x-internal="event-click"}
events that are
[queued](webappapis.html#queue-a-task){#enabling-and-disabling-form-controls:-the-disabled-attribute:queue-a-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#enabling-and-disabling-form-controls:-the-disabled-attribute:user-interaction-task-source}
from being dispatched on the element.

**Constraint validation**: If an element is
[disabled](#concept-fe-disabled){#enabling-and-disabling-form-controls:-the-disabled-attribute:concept-fe-disabled-2},
it is [barred from constraint
validation](#barred-from-constraint-validation){#enabling-and-disabling-form-controls:-the-disabled-attribute:barred-from-constraint-validation}.

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLButtonElement/disabled](https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement/disabled "The HTMLButtonElement.disabled property indicates whether the control is disabled, meaning that it does not accept any clicks.")

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

:::: feature
[HTMLSelectElement/disabled](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/disabled "The HTMLSelectElement.disabled property is a boolean value that reflects the disabled HTML attribute, which indicates whether the control is disabled. If it is disabled, it does not accept clicks. A disabled element is unusable and un-clickable.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
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
:::::::

The [`disabled`]{#dom-fe-disabled .dfn
dfn-for="HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#enabling-and-disabling-form-controls:-the-disabled-attribute:reflect}
the
[`disabled`{#enabling-and-disabling-form-controls:-the-disabled-attribute:attr-fe-disabled-3}](#attr-fe-disabled)
content attribute.

##### [4.10.18.6]{.secno} Form submission attributes[](#form-submission-attributes){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/form#Attributes_for_form_submission](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form#Attributes_for_form_submission "The <form> HTML element represents a document section containing interactive controls for submitting information.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari10.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[Attributes for form submission]{#attributes-for-form-submission .dfn}
can be specified both on
[`form`{#form-submission-attributes:the-form-element}](forms.html#the-form-element)
elements and on [submit
buttons](forms.html#concept-submit-button){#form-submission-attributes:concept-submit-button}
(elements that represent buttons that submit forms, e.g. an
[`input`{#form-submission-attributes:the-input-element}](input.html#the-input-element)
element whose
[`type`{#form-submission-attributes:attr-input-type}](input.html#attr-input-type)
attribute is in the [Submit
Button](input.html#submit-button-state-(type=submit)){#form-submission-attributes:submit-button-state-(type=submit)}
state).

The [attributes for form
submission](#attributes-for-form-submission){#form-submission-attributes:attributes-for-form-submission}
that may be specified on
[`form`{#form-submission-attributes:the-form-element-2}](forms.html#the-form-element)
elements are
[`action`{#form-submission-attributes:attr-fs-action}](#attr-fs-action),
[`enctype`{#form-submission-attributes:attr-fs-enctype}](#attr-fs-enctype),
[`method`{#form-submission-attributes:attr-fs-method}](#attr-fs-method),
[`novalidate`{#form-submission-attributes:attr-fs-novalidate}](#attr-fs-novalidate),
and
[`target`{#form-submission-attributes:attr-fs-target}](#attr-fs-target).

The corresponding [attributes for form
submission](#attributes-for-form-submission){#form-submission-attributes:attributes-for-form-submission-2}
that may be specified on [submit
buttons](forms.html#concept-submit-button){#form-submission-attributes:concept-submit-button-2}
are
[`formaction`{#form-submission-attributes:attr-fs-formaction}](#attr-fs-formaction),
[`formenctype`{#form-submission-attributes:attr-fs-formenctype}](#attr-fs-formenctype),
[`formmethod`{#form-submission-attributes:attr-fs-formmethod}](#attr-fs-formmethod),
[`formnovalidate`{#form-submission-attributes:attr-fs-formnovalidate}](#attr-fs-formnovalidate),
and
[`formtarget`{#form-submission-attributes:attr-fs-formtarget}](#attr-fs-formtarget).
When omitted, they default to the values given on the corresponding
attributes on the
[`form`{#form-submission-attributes:the-form-element-3}](forms.html#the-form-element)
element.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#formaction](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#formaction "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`action`]{#attr-fs-action .dfn dfn-for="form,button"
dfn-type="element-attr"} and [`formaction`]{#attr-fs-formaction .dfn
dfn-for="form,button" dfn-type="element-attr"} content attributes, if
specified, must have a value that is a [valid non-empty URL potentially
surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#form-submission-attributes:valid-non-empty-url-potentially-surrounded-by-spaces}.

The [action]{#concept-fs-action .dfn} of an element is the value of the
element\'s
[`formaction`{#form-submission-attributes:attr-fs-formaction-2}](#attr-fs-formaction)
attribute, if the element is a [submit
button](forms.html#concept-submit-button){#form-submission-attributes:concept-submit-button-3}
and has such an attribute, or the value of its [form
owner](#form-owner){#form-submission-attributes:form-owner}\'s
[`action`{#form-submission-attributes:attr-fs-action-2}](#attr-fs-action)
attribute, if *it* has one, or else the empty string.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#formmethod](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#formmethod "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`method`]{#attr-fs-method .dfn dfn-for="form,button"
dfn-type="element-attr"} and [`formmethod`]{#attr-fs-formmethod .dfn
dfn-for="form,button" dfn-type="element-attr"} content attributes are
[enumerated
attributes](common-microsyntaxes.html#enumerated-attribute){#form-submission-attributes:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`get`]{#attr-fs-method-get-keyword .dfn
dfn-for="form/method,button/formmethod,input/formmethod"
dfn-type="attr-value"}

[GET]{#attr-fs-method-get .dfn}

Indicates the
[`form`{#form-submission-attributes:the-form-element-4}](forms.html#the-form-element)
will use the HTTP GET method.

[`post`]{#attr-fs-method-post-keyword .dfn
dfn-for="form/method,button/formmethod,input/formmethod"
dfn-type="attr-value"}

[POST]{#attr-fs-method-post .dfn}

Indicates the
[`form`{#form-submission-attributes:the-form-element-5}](forms.html#the-form-element)
will use the HTTP POST method.

[`dialog`]{#attr-fs-method-dialog-keyword .dfn
dfn-for="form/method,button/formmethod,input/formmethod"
dfn-type="attr-value"}

[Dialog]{#attr-fs-method-dialog .dfn}

Indicates the
[`form`{#form-submission-attributes:the-form-element-6}](forms.html#the-form-element)
is intended to close the
[`dialog`{#form-submission-attributes:the-dialog-element}](interactive-elements.html#the-dialog-element)
box in which the form finds itself, if any, and otherwise not submit.

The
[`method`{#form-submission-attributes:attr-fs-method-2}](#attr-fs-method)
attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[GET](#attr-fs-method-get){#form-submission-attributes:attr-fs-method-get}
state.

The
[`formmethod`{#form-submission-attributes:attr-fs-formmethod-2}](#attr-fs-formmethod)
attribute has no *[missing value
default](common-microsyntaxes.html#missing-value-default)*, and its
*[invalid value
default](common-microsyntaxes.html#invalid-value-default)* is the
[GET](#attr-fs-method-get){#form-submission-attributes:attr-fs-method-get-2}
state.

The [method]{#concept-fs-method .dfn} of an element is one of those
states. If the element is a [submit
button](forms.html#concept-submit-button){#form-submission-attributes:concept-submit-button-4}
and has a
[`formmethod`{#form-submission-attributes:attr-fs-formmethod-3}](#attr-fs-formmethod)
attribute, then the element\'s
[method](#concept-fs-method){#form-submission-attributes:concept-fs-method}
is that attribute\'s state; otherwise, it is the [form
owner](#form-owner){#form-submission-attributes:form-owner-2}\'s
[`method`{#form-submission-attributes:attr-fs-method-3}](#attr-fs-method)
attribute\'s state.

::: example
Here the
[`method`{#form-submission-attributes:attr-fs-method-4}](#attr-fs-method)
attribute is used to explicitly specify the default value,
\"[`get`{#form-submission-attributes:attr-fs-method-get-keyword}](#attr-fs-method-get-keyword)\",
so that the search query is submitted in the URL:

``` html
<form method="get" action="/search.cgi">
 <p><label>Search terms: <input type=search name=q></label></p>
 <p><input type=submit></p>
</form>
```
:::

::: example
On the other hand, here the
[`method`{#form-submission-attributes:attr-fs-method-5}](#attr-fs-method)
attribute is used to specify the value
\"[`post`{#form-submission-attributes:attr-fs-method-post-keyword}](#attr-fs-method-post-keyword)\",
so that the user\'s message is submitted in the HTTP request\'s body:

``` html
<form method="post" action="/post-message.cgi">
 <p><label>Message: <input type=text name=m></label></p>
 <p><input type=submit value="Submit message"></p>
</form>
```
:::

::: example
In this example, a
[`form`{#form-submission-attributes:the-form-element-7}](forms.html#the-form-element)
is used with a
[`dialog`{#form-submission-attributes:the-dialog-element-2}](interactive-elements.html#the-dialog-element).
The
[`method`{#form-submission-attributes:attr-fs-method-6}](#attr-fs-method)
attribute\'s
\"[`dialog`{#form-submission-attributes:attr-fs-method-dialog-keyword}](#attr-fs-method-dialog-keyword)\"
keyword is used to have the dialog automatically close when the form is
submitted.

``` {lang="en-GB"}
<dialog id="ship">
 <form method=dialog>
  <p>A ship has arrived in the harbour.</p>
  <button type=submit value="board">Board the ship</button>
  <button type=submit value="call">Call to the captain</button>
 </form>
</dialog>
<script>
 var ship = document.getElementById('ship');
 ship.showModal();
 ship.onclose = function (event) {
   if (ship.returnValue == 'board') {
     // ...
   } else {
     // ...
   }
 };
</script>
```
:::

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#formenctype](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#formenctype "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`enctype`]{#attr-fs-enctype .dfn dfn-for="form,button"
dfn-type="element-attr"} and [`formenctype`]{#attr-fs-formenctype .dfn
dfn-for="form,button" dfn-type="element-attr"} content attributes are
[enumerated
attributes](common-microsyntaxes.html#enumerated-attribute){#form-submission-attributes:enumerated-attribute-2}
with the following keywords and states:

- The
  \"[`application/x-www-form-urlencoded`]{#attr-fs-enctype-urlencoded
  .dfn dfn-for="form/enctype,button/enctype" dfn-type="attr-value"}\"
  keyword and corresponding state.
- The \"[`multipart/form-data`]{#attr-fs-enctype-formdata .dfn
  dfn-for="form/enctype,button/enctype" dfn-type="attr-value"}\" keyword
  and corresponding state.
- The \"[`text/plain`]{#attr-fs-enctype-text .dfn
  dfn-for="form/enctype,button/enctype" dfn-type="attr-value"}\" keyword
  and corresponding state.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[`application/x-www-form-urlencoded`{#form-submission-attributes:attr-fs-enctype-urlencoded}](#attr-fs-enctype-urlencoded)
state.

The
[`formenctype`{#form-submission-attributes:attr-fs-formenctype-2}](#attr-fs-formenctype)
attribute has no *[missing value
default](common-microsyntaxes.html#missing-value-default)*, and its
*[invalid value
default](common-microsyntaxes.html#invalid-value-default)* is the
[`application/x-www-form-urlencoded`{#form-submission-attributes:attr-fs-enctype-urlencoded-2}](#attr-fs-enctype-urlencoded)
state.

The [enctype]{#concept-fs-enctype .dfn} of an element is one of those
three states. If the element is a [submit
button](forms.html#concept-submit-button){#form-submission-attributes:concept-submit-button-5}
and has a
[`formenctype`{#form-submission-attributes:attr-fs-formenctype-3}](#attr-fs-formenctype)
attribute, then the element\'s
[enctype](#concept-fs-enctype){#form-submission-attributes:concept-fs-enctype}
is that attribute\'s state; otherwise, it is the [form
owner](#form-owner){#form-submission-attributes:form-owner-3}\'s
[`enctype`{#form-submission-attributes:attr-fs-enctype-2}](#attr-fs-enctype)
attribute\'s state.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#formtarget](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#formtarget "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome9+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`target`]{#attr-fs-target .dfn dfn-for="form,button"
dfn-type="element-attr"} and [`formtarget`]{#attr-fs-formtarget .dfn
dfn-for="form,button" dfn-type="element-attr"} content attributes, if
specified, must have values that are [valid navigable target names or
keywords](document-sequences.html#valid-navigable-target-name-or-keyword){#form-submission-attributes:valid-navigable-target-name-or-keyword}.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#formnovalidate](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#formnovalidate "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`novalidate`]{#attr-fs-novalidate .dfn dfn-for="form,button"
dfn-type="element-attr"} and [`formnovalidate`]{#attr-fs-formnovalidate
.dfn dfn-for="form,button" dfn-type="element-attr"} content attributes
are [boolean
attributes](common-microsyntaxes.html#boolean-attribute){#form-submission-attributes:boolean-attribute}.
If present, they indicate that the form is not to be validated during
submission.

The [no-validate state]{#concept-fs-novalidate .dfn} of an element is
true if the element is a [submit
button](forms.html#concept-submit-button){#form-submission-attributes:concept-submit-button-6}
and the element\'s
[`formnovalidate`{#form-submission-attributes:attr-fs-formnovalidate-2}](#attr-fs-formnovalidate)
attribute is present, or if the element\'s [form
owner](#form-owner){#form-submission-attributes:form-owner-4}\'s
[`novalidate`{#form-submission-attributes:attr-fs-novalidate-2}](#attr-fs-novalidate)
attribute is present, and false otherwise.

::: example
This attribute is useful to include \"save\" buttons on forms that have
validation constraints, to allow users to save their progress even
though they haven\'t fully entered the data in the form. The following
example shows a simple form that has two required fields. There are
three buttons: one to submit the form, which requires both fields to be
filled in; one to save the form so that the user can come back and fill
it in later; and one to cancel the form altogether.

``` html
<form action="editor.cgi" method="post">
 <p><label>Name: <input required name=fn></label></p>
 <p><label>Essay: <textarea required name=essay></textarea></label></p>
 <p><input type=submit name=submit value="Submit essay"></p>
 <p><input type=submit formnovalidate name=save value="Save essay"></p>
 <p><input type=submit formnovalidate name=cancel value="Cancel"></p>
</form>
```
:::

------------------------------------------------------------------------

::::::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/action](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/action "The HTMLFormElement.action property represents the action of the <form> element.")

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

:::: feature
[HTMLFormElement/target](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/target "The target property of the HTMLFormElement interface represents the target of the form's action (i.e., the frame in which to render its output).")

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

:::: feature
[HTMLFormElement/method](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/method "The HTMLFormElement.method property represents the HTTP method used to submit the <form>.")

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

:::: feature
[HTMLFormElement/enctype](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/enctype "The HTMLFormElement.enctype property is the MIME type of content that is used to submit the form to the server. Possible values are:")

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

:::: feature
[HTMLFormElement/encoding](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/encoding "The HTMLFormElement.encoding property is an alternative name for the enctype element on the DOM HTMLFormElement object.")

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
:::::::::::::

The [`action`]{#dom-fs-action .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect}
the content attribute of the same name, except that on getting, when the
content attribute is missing or its value is the empty string, the
element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#form-submission-attributes:node-document
x-internal="node-document"}\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#form-submission-attributes:the-document's-address
x-internal="the-document's-address"} must be returned instead. The
[`target`]{#dom-fs-target .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-2}
the content attribute of the same name. The [`method`]{#dom-fs-method
.dfn dfn-for="HTMLFormElement" dfn-type="attribute"} and
[`enctype`]{#dom-fs-enctype .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-3}
the respective content attributes of the same name, [limited to only
known
values](common-dom-interfaces.html#limited-to-only-known-values){#form-submission-attributes:limited-to-only-known-values}.
The [`encoding`]{#dom-fs-encoding .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-4}
the
[`enctype`{#form-submission-attributes:attr-fs-enctype-3}](#attr-fs-enctype)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#form-submission-attributes:limited-to-only-known-values-2}.
The [`noValidate`]{#dom-fs-novalidate .dfn dfn-for="HTMLFormElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-5}
the
[`novalidate`{#form-submission-attributes:attr-fs-novalidate-3}](#attr-fs-novalidate)
content attribute. The [`formAction`]{#dom-fs-formaction .dfn
dfn-for="HTMLButtonElement,HTMLInputElement" dfn-type="attribute"} IDL
attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-6}
the
[`formaction`{#form-submission-attributes:attr-fs-formaction-3}](#attr-fs-formaction)
content attribute, except that on getting, when the content attribute is
missing or its value is the empty string, the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#form-submission-attributes:node-document-2
x-internal="node-document"}\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#form-submission-attributes:the-document's-address-2
x-internal="the-document's-address"} must be returned instead. The
[`formEnctype`]{#dom-fs-formenctype .dfn
dfn-for="HTMLButtonElement,HTMLInputElement" dfn-type="attribute"} IDL
attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-7}
the
[`formenctype`{#form-submission-attributes:attr-fs-formenctype-4}](#attr-fs-formenctype)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#form-submission-attributes:limited-to-only-known-values-3}.
The [`formMethod`]{#dom-fs-formmethod .dfn
dfn-for="HTMLButtonElement,HTMLInputElement" dfn-type="attribute"} IDL
attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-8}
the
[`formmethod`{#form-submission-attributes:attr-fs-formmethod-4}](#attr-fs-formmethod)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#form-submission-attributes:limited-to-only-known-values-4}.
The [`formNoValidate`]{#dom-fs-formnovalidate .dfn
dfn-for="HTMLButtonElement,HTMLInputElement" dfn-type="attribute"} IDL
attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-9}
the
[`formnovalidate`{#form-submission-attributes:attr-fs-formnovalidate-3}](#attr-fs-formnovalidate)
content attribute. The [`formTarget`]{#dom-fs-formtarget .dfn
dfn-for="HTMLButtonElement,HTMLInputElement" dfn-type="attribute"} IDL
attribute must
[reflect](common-dom-interfaces.html#reflect){#form-submission-attributes:reflect-10}
the
[`formtarget`{#form-submission-attributes:attr-fs-formtarget-2}](#attr-fs-formtarget)
content attribute.

##### [4.10.18.7]{.secno} Autofill[](#autofill){.self-link}

###### [4.10.18.7.1]{.secno} Autofilling form controls: the [`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete}](#attr-fe-autocomplete) attribute[](#autofilling-form-controls:-the-autocomplete-attribute){.self-link} {#autofilling-form-controls:-the-autocomplete-attribute}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/autocomplete](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete "The HTML autocomplete attribute lets web developers specify what if any permission the user agent has to provide automated assistance in filling out form field values, as well as guidance to the browser as to the type of information expected in the field.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome14+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)No]{.edge .no}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

User agents sometimes have features for helping users fill forms in, for
example prefilling the user\'s address based on earlier user input. The
[`autocomplete`]{#attr-fe-autocomplete .dfn
dfn-for="button,fieldset,input,object,output,select,textarea"
dfn-type="element-attr"} content attribute can be used to hint to the
user agent how to, or indeed whether to, provide such a feature.

There are two ways this attribute is used. When wearing the [autofill
expectation mantle]{#autofill-expectation-mantle .dfn}, the
[`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-2}](#attr-fe-autocomplete)
attribute describes what input is expected from users. When wearing the
[autofill anchor mantle]{#autofill-anchor-mantle .dfn}, the
[`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-3}](#attr-fe-autocomplete)
attribute describes the meaning of the given value.

On an
[`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element}](input.html#the-input-element)
element whose
[`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type}](input.html#attr-input-type)
attribute is in the
[Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)}
state, the
[`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-4}](#attr-fe-autocomplete)
attribute wears the [autofill anchor
mantle](#autofill-anchor-mantle){#autofilling-form-controls:-the-autocomplete-attribute:autofill-anchor-mantle}.
In all other cases, it wears the [autofill expectation
mantle](#autofill-expectation-mantle){#autofilling-form-controls:-the-autocomplete-attribute:autofill-expectation-mantle}.

When wearing the [autofill expectation
mantle](#autofill-expectation-mantle){#autofilling-form-controls:-the-autocomplete-attribute:autofill-expectation-mantle-2},
the
[`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-5}](#attr-fe-autocomplete)
attribute, if specified, must have a value that is an ordered [set of
space-separated
tokens](common-microsyntaxes.html#set-of-space-separated-tokens){#autofilling-form-controls:-the-autocomplete-attribute:set-of-space-separated-tokens}
consisting of either a single token that is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string
\"[`off`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-off}](#attr-fe-autocomplete-off)\",
or a single token that is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} match for the string
\"[`on`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-on}](#attr-fe-autocomplete-on)\",
or [autofill detail
tokens](#autofill-detail-tokens){#autofilling-form-controls:-the-autocomplete-attribute:autofill-detail-tokens}.

When wearing the [autofill anchor
mantle](#autofill-anchor-mantle){#autofilling-form-controls:-the-autocomplete-attribute:autofill-anchor-mantle-2},
the
[`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-6}](#attr-fe-autocomplete)
attribute, if specified, must have a value that is an ordered [set of
space-separated
tokens](common-microsyntaxes.html#set-of-space-separated-tokens){#autofilling-form-controls:-the-autocomplete-attribute:set-of-space-separated-tokens-2}
consisting of just [autofill detail
tokens](#autofill-detail-tokens){#autofilling-form-controls:-the-autocomplete-attribute:autofill-detail-tokens-2}
(i.e. the
\"[`on`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-on-2}](#attr-fe-autocomplete-on)\"
and
\"[`off`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-off-2}](#attr-fe-autocomplete-off)\"
keywords are not allowed).

[Autofill detail tokens]{#autofill-detail-tokens .dfn} are the
following, in the order given below:

1.  Optionally, a token whose first eight characters are an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive-3
    x-internal="ascii-case-insensitive"} match for the string
    \"[`section-`]{#attr-fe-autocomplete-section .dfn
    dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
    dfn-type="attr-value"}\", meaning that the field belongs to the
    named group.

    ::: example
    For example, if there are two shipping addresses in the form, then
    they could be marked up as:

    ``` html
    <fieldset>
     <legend>Ship the blue gift to...</legend>
     <p> <label> Address:     <textarea name=ba autocomplete="section-blue shipping street-address"></textarea> </label>
     <p> <label> City:        <input name=bc autocomplete="section-blue shipping address-level2"> </label>
     <p> <label> Postal Code: <input name=bp autocomplete="section-blue shipping postal-code"> </label>
    </fieldset>
    <fieldset>
     <legend>Ship the red gift to...</legend>
     <p> <label> Address:     <textarea name=ra autocomplete="section-red shipping street-address"></textarea> </label>
     <p> <label> City:        <input name=rc autocomplete="section-red shipping address-level2"> </label>
     <p> <label> Postal Code: <input name=rp autocomplete="section-red shipping postal-code"> </label>
    </fieldset>
    ```
    :::

2.  Optionally, a token that is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive-4
    x-internal="ascii-case-insensitive"} match for one of the following
    strings:

    - \"[`shipping`]{#attr-fe-autocomplete-shipping .dfn
      dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
      dfn-type="attr-value"}\", meaning the field is part of the
      shipping address or contact information
    - \"[`billing`]{#attr-fe-autocomplete-billing .dfn
      dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
      dfn-type="attr-value"}\", meaning the field is part of the billing
      address or contact information

3.  Either of the following two options:

    - A token that is an [ASCII
      case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive-5
      x-internal="ascii-case-insensitive"} match for one of the
      following [autofill
      field](#autofill-field){#autofilling-form-controls:-the-autocomplete-attribute:autofill-field}
      names, excluding those that are [inappropriate for the
      control](#inappropriate-for-the-control){#autofilling-form-controls:-the-autocomplete-attribute:inappropriate-for-the-control}:

      - \"[`name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-name}](#attr-fe-autocomplete-name)\"
      - \"[`honorific-prefix`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-honorific-prefix}](#attr-fe-autocomplete-honorific-prefix)\"
      - \"[`given-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-given-name}](#attr-fe-autocomplete-given-name)\"
      - \"[`additional-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-additional-name}](#attr-fe-autocomplete-additional-name)\"
      - \"[`family-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-family-name}](#attr-fe-autocomplete-family-name)\"
      - \"[`honorific-suffix`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-honorific-suffix}](#attr-fe-autocomplete-honorific-suffix)\"
      - \"[`nickname`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-nickname}](#attr-fe-autocomplete-nickname)\"
      - \"[`username`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-username}](#attr-fe-autocomplete-username)\"
      - \"[`new-password`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-new-password}](#attr-fe-autocomplete-new-password)\"
      - \"[`current-password`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-current-password}](#attr-fe-autocomplete-current-password)\"
      - \"[`one-time-code`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-one-time-code}](#attr-fe-autocomplete-one-time-code)\"
      - \"[`organization-title`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-organization-title}](#attr-fe-autocomplete-organization-title)\"
      - \"[`organization`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-organization}](#attr-fe-autocomplete-organization)\"
      - \"[`street-address`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-street-address}](#attr-fe-autocomplete-street-address)\"
      - \"[`address-line1`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-line1}](#attr-fe-autocomplete-address-line1)\"
      - \"[`address-line2`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-line2}](#attr-fe-autocomplete-address-line2)\"
      - \"[`address-line3`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-line3}](#attr-fe-autocomplete-address-line3)\"
      - \"[`address-level4`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level4}](#attr-fe-autocomplete-address-level4)\"
      - \"[`address-level3`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level3}](#attr-fe-autocomplete-address-level3)\"
      - \"[`address-level2`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level2}](#attr-fe-autocomplete-address-level2)\"
      - \"[`address-level1`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level1}](#attr-fe-autocomplete-address-level1)\"
      - \"[`country`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-country}](#attr-fe-autocomplete-country)\"
      - \"[`country-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-country-name}](#attr-fe-autocomplete-country-name)\"
      - \"[`postal-code`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-postal-code}](#attr-fe-autocomplete-postal-code)\"
      - \"[`cc-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-name}](#attr-fe-autocomplete-cc-name)\"
      - \"[`cc-given-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-given-name}](#attr-fe-autocomplete-cc-given-name)\"
      - \"[`cc-additional-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-additional-name}](#attr-fe-autocomplete-cc-additional-name)\"
      - \"[`cc-family-name`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-family-name}](#attr-fe-autocomplete-cc-family-name)\"
      - \"[`cc-number`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-number}](#attr-fe-autocomplete-cc-number)\"
      - \"[`cc-exp`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-exp}](#attr-fe-autocomplete-cc-exp)\"
      - \"[`cc-exp-month`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-exp-month}](#attr-fe-autocomplete-cc-exp-month)\"
      - \"[`cc-exp-year`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-exp-year}](#attr-fe-autocomplete-cc-exp-year)\"
      - \"[`cc-csc`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-csc}](#attr-fe-autocomplete-cc-csc)\"
      - \"[`cc-type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-type}](#attr-fe-autocomplete-cc-type)\"
      - \"[`transaction-currency`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-transaction-currency}](#attr-fe-autocomplete-transaction-currency)\"
      - \"[`transaction-amount`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-transaction-amount}](#attr-fe-autocomplete-transaction-amount)\"
      - \"[`language`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-language}](#attr-fe-autocomplete-language)\"
      - \"[`bday`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-bday}](#attr-fe-autocomplete-bday)\"
      - \"[`bday-day`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-bday-day}](#attr-fe-autocomplete-bday-day)\"
      - \"[`bday-month`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-bday-month}](#attr-fe-autocomplete-bday-month)\"
      - \"[`bday-year`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-bday-year}](#attr-fe-autocomplete-bday-year)\"
      - \"[`sex`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-sex}](#attr-fe-autocomplete-sex)\"
      - \"[`url`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-url}](#attr-fe-autocomplete-url)\"
      - \"[`photo`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-photo}](#attr-fe-autocomplete-photo)\"

      (See the table below for descriptions of these values.)

    - The following, in the given order:

      1.  Optionally, a token that is an [ASCII
          case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive-6
          x-internal="ascii-case-insensitive"} match for one of the
          following strings:

          - \"[`home`]{#attr-fe-autocomplete-home .dfn
            dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
            dfn-type="attr-value"}\", meaning the field is for
            contacting someone at their residence
          - \"[`work`]{#attr-fe-autocomplete-work .dfn
            dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
            dfn-type="attr-value"}\", meaning the field is for
            contacting someone at their workplace
          - \"[`mobile`]{#attr-fe-autocomplete-mobile .dfn
            dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
            dfn-type="attr-value"}\", meaning the field is for
            contacting someone regardless of location
          - \"[`fax`]{#attr-fe-autocomplete-fax .dfn
            dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
            dfn-type="attr-value"}\", meaning the field describes a fax
            machine\'s contact details
          - \"[`pager`]{#attr-fe-autocomplete-pager .dfn
            dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
            dfn-type="attr-value"}\", meaning the field describes a
            pager\'s or beeper\'s contact details

      2.  A token that is an [ASCII
          case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive-7
          x-internal="ascii-case-insensitive"} match for one of the
          following [autofill
          field](#autofill-field){#autofilling-form-controls:-the-autocomplete-attribute:autofill-field-2}
          names, excluding those that are [inappropriate for the
          control](#inappropriate-for-the-control){#autofilling-form-controls:-the-autocomplete-attribute:inappropriate-for-the-control-2}:

          - \"[`tel`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel}](#attr-fe-autocomplete-tel)\"
          - \"[`tel-country-code`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel-country-code}](#attr-fe-autocomplete-tel-country-code)\"
          - \"[`tel-national`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel-national}](#attr-fe-autocomplete-tel-national)\"
          - \"[`tel-area-code`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel-area-code}](#attr-fe-autocomplete-tel-area-code)\"
          - \"[`tel-local`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel-local}](#attr-fe-autocomplete-tel-local)\"
          - \"[`tel-local-prefix`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel-local-prefix}](#attr-fe-autocomplete-tel-local-prefix)\"
          - \"[`tel-local-suffix`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel-local-suffix}](#attr-fe-autocomplete-tel-local-suffix)\"
          - \"[`tel-extension`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-tel-extension}](#attr-fe-autocomplete-tel-extension)\"
          - \"[`email`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-email}](#attr-fe-autocomplete-email)\"
          - \"[`impp`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-impp}](#attr-fe-autocomplete-impp)\"

          (See the table below for descriptions of these values.)

4.  Optionally, a token that is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofilling-form-controls:-the-autocomplete-attribute:ascii-case-insensitive-8
    x-internal="ascii-case-insensitive"} match for the string
    \"[`webauthn`]{#attr-fe-autocomplete-webauthn .dfn}\", meaning the
    user agent should show [public key
    credentials](https://w3c.github.io/webauthn/#public-key-credential){#autofilling-form-controls:-the-autocomplete-attribute:public-key-credential
    x-internal="public-key-credential"} available via
    [`conditional`{#autofilling-form-controls:-the-autocomplete-attribute:conditional-mediation}](https://w3c.github.io/webappsec-credential-management/#dom-credentialmediationrequirement-conditional){x-internal="conditional-mediation"}
    mediation when the user interacts with the form control.
    [`webauthn`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-webauthn}](#attr-fe-autocomplete-webauthn)
    is only valid for
    [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-2}](input.html#the-input-element)
    and
    [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element}](form-elements.html#the-textarea-element)
    elements.

As noted earlier, the meaning of the attribute and its keywords depends
on the mantle that the attribute is wearing.

When wearing the [autofill expectation mantle](#autofill-expectation-mantle){#autofilling-form-controls:-the-autocomplete-attribute:autofill-expectation-mantle-3}\...

:   The \"[`off`]{#attr-fe-autocomplete-off .dfn
    dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
    dfn-type="attr-value"}\" keyword indicates either that the
    control\'s input data is particularly sensitive (for example the
    activation code for a nuclear weapon); or that it is a value that
    will never be reused (for example a one-time-key for a bank login)
    and the user will therefore have to explicitly enter the data each
    time, instead of being able to rely on the UA to prefill the value
    for them; or that the document provides its own autocomplete
    mechanism and does not want the user agent to provide autocompletion
    values.

    The \"[`on`]{#attr-fe-autocomplete-on .dfn
    dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
    dfn-type="attr-value"}\" keyword indicates that the user agent is
    allowed to provide the user with autocompletion values, but does not
    provide any further information about what kind of data the user
    might be expected to enter. User agents would have to use heuristics
    to decide what autocompletion values to suggest.

    The [autofill
    field](#autofill-field){#autofilling-form-controls:-the-autocomplete-attribute:autofill-field-3}
    listed above indicate that the user agent is allowed to provide the
    user with autocompletion values, and specifies what kind of value is
    expected. The meaning of each such keyword is described in the table
    below.

    If the
    [`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-7}](#attr-fe-autocomplete)
    attribute is omitted, the default value corresponding to the state
    of the element\'s [form
    owner](#form-owner){#autofilling-form-controls:-the-autocomplete-attribute:form-owner}\'s
    [`autocomplete`{#autofilling-form-controls:-the-autocomplete-attribute:attr-form-autocomplete}](forms.html#attr-form-autocomplete)
    attribute is used instead (either
    \"[`on`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-on-3}](#attr-fe-autocomplete-on)\"
    or
    \"[`off`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-off-3}](#attr-fe-autocomplete-off)\").
    If there is no [form
    owner](#form-owner){#autofilling-form-controls:-the-autocomplete-attribute:form-owner-2},
    then the value
    \"[`on`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-on-4}](#attr-fe-autocomplete-on)\"
    is used.

When wearing the [autofill anchor mantle](#autofill-anchor-mantle){#autofilling-form-controls:-the-autocomplete-attribute:autofill-anchor-mantle-3}\...

:   The [autofill
    field](#autofill-field){#autofilling-form-controls:-the-autocomplete-attribute:autofill-field-4}
    listed above indicate that the value of the particular kind of value
    specified is that value provided for this element. The meaning of
    each such keyword is described in the table below.

    ::: example
    In this example the page has explicitly specified the currency and
    amount of the transaction. The form requests a credit card and other
    billing details. The user agent could use this information to
    suggest a credit card that it knows has sufficient balance and that
    supports the relevant currency.

    ``` html
    <form method=post action="step2.cgi">
     <input type=hidden autocomplete=transaction-currency value="CHF">
     <input type=hidden autocomplete=transaction-amount value="15.00">
     <p><label>Credit card number: <input type=text inputmode=numeric autocomplete=cc-number></label>
     <p><label>Expiry Date: <input type=month autocomplete=cc-exp></label>
     <p><input type=submit value="Continue...">
    </form>
    ```
    :::

The [autofill field]{#autofill-field .dfn} keywords relate to each other
as described in the table below. Each field name listed on a row of this
table corresponds to the meaning given in the cell for that row in the
column labeled \"Meaning\". Some fields correspond to subparts of other
fields; for example, a credit card expiry date can be expressed as one
field giving both the month and year of expiry
(\"[`cc-exp`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-exp-2}](#attr-fe-autocomplete-cc-exp)\"),
or as two fields, one giving the month
(\"[`cc-exp-month`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-exp-month-2}](#attr-fe-autocomplete-cc-exp-month)\")
and one the year
(\"[`cc-exp-year`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-cc-exp-year-2}](#attr-fe-autocomplete-cc-exp-year)\").
In such cases, the names of the broader fields cover multiple rows, in
which the narrower fields are defined.

Generally, authors are encouraged to use the broader fields rather than
the narrower fields, as the narrower fields tend to expose Western
biases. For example, while it is common in some Western cultures to have
a given name and a family name, in that order (and thus often referred
to as a *first name* and a *surname*), many cultures put the family name
first and the given name second, and many others simply have one name (a
*mononym*). Having a single field is therefore more flexible.

Some fields are only appropriate for certain form controls. An [autofill
field](#autofill-field){#autofilling-form-controls:-the-autocomplete-attribute:autofill-field-5}
name is [inappropriate for a control]{#inappropriate-for-the-control
.dfn} if the control does not belong to the group listed for that
[autofill
field](#autofill-field){#autofilling-form-controls:-the-autocomplete-attribute:autofill-field-6}
in the fifth column of the first row describing that [autofill
field](#autofill-field){#autofilling-form-controls:-the-autocomplete-attribute:autofill-field-7}
in the table below. What controls fall into each group is described
below the table.

Field name

Meaning

Canonical Format

Canonical Format Example

Control group

\"[`name`]{#attr-fe-autocomplete-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Full name

Free-form text, no newlines

Sir Timothy John Berners-Lee, OM, KBE, FRS, FREng, FRSA

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text}

\"[`honorific-prefix`]{#attr-fe-autocomplete-honorific-prefix .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Prefix or title (e.g. \"Mr.\", \"Ms.\", \"Dr.\",
\"[M^lle^]{lang="fr"}\")

Free-form text, no newlines

Sir

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-2}

\"[`given-name`]{#attr-fe-autocomplete-given-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Given name (in some Western cultures, also known as the *first name*)

Free-form text, no newlines

Timothy

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-3}

\"[`additional-name`]{#attr-fe-autocomplete-additional-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Additional names (in some Western cultures, also known as *middle
names*, forenames other than the first name)

Free-form text, no newlines

John

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-4}

\"[`family-name`]{#attr-fe-autocomplete-family-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Family name (in some Western cultures, also known as the *last name* or
*surname*)

Free-form text, no newlines

Berners-Lee

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-5}

\"[`honorific-suffix`]{#attr-fe-autocomplete-honorific-suffix .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Suffix (e.g. \"Jr.\", \"B.Sc.\", \"MBASW\", \"II\")

Free-form text, no newlines

OM, KBE, FRS, FREng, FRSA

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-6}

\"[`nickname`]{#attr-fe-autocomplete-nickname .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Nickname, screen name, handle: a typically short name used instead of
the full name

Free-form text, no newlines

Tim

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-7}

\"[`organization-title`]{#attr-fe-autocomplete-organization-title .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Job title (e.g. \"Software Engineer\", \"Senior Vice President\",
\"Deputy Managing Director\")

Free-form text, no newlines

Professor

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-8}

\"[`username`]{#attr-fe-autocomplete-username .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

A username

Free-form text, no newlines

timbl

[Username](#control-group-username){#autofilling-form-controls:-the-autocomplete-attribute:control-group-username}

\"[`new-password`]{#attr-fe-autocomplete-new-password .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

A new password (e.g. when creating an account or changing a password)

Free-form text, no newlines

GUMFXbadyrS3

[Password](#control-group-password){#autofilling-form-controls:-the-autocomplete-attribute:control-group-password}

\"[`current-password`]{#attr-fe-autocomplete-current-password .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

The current password for the account identified by the
[`username`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-username-2}](#attr-fe-autocomplete-username)
field (e.g. when logging in)

Free-form text, no newlines

qwerty

[Password](#control-group-password){#autofilling-form-controls:-the-autocomplete-attribute:control-group-password-2}

\"[`one-time-code`]{#attr-fe-autocomplete-one-time-code .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

One-time code used for verifying user identity

Free-form text, no newlines

123456

[Password](#control-group-password){#autofilling-form-controls:-the-autocomplete-attribute:control-group-password-3}

\"[`organization`]{#attr-fe-autocomplete-organization .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Company name corresponding to the person, address, or contact
information in the other fields associated with this field

Free-form text, no newlines

World Wide Web Consortium

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-9}

\"[`street-address`]{#attr-fe-autocomplete-street-address .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Street address (multiple lines, newlines preserved)

Free-form text

32 Vassar Street\
MIT Room 32-G524

[Multiline](#control-group-multiline){#autofilling-form-controls:-the-autocomplete-attribute:control-group-multiline}

\"[`address-line1`]{#attr-fe-autocomplete-address-line1 .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Street address (one line per field)

Free-form text, no newlines

32 Vassar Street

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-10}

\"[`address-line2`]{#attr-fe-autocomplete-address-line2 .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Free-form text, no newlines

MIT Room 32-G524

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-11}

\"[`address-line3`]{#attr-fe-autocomplete-address-line3 .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Free-form text, no newlines

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-12}

\"[`address-level4`]{#attr-fe-autocomplete-address-level4 .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

The most fine-grained [administrative level](#more-on-address-levels),
in addresses with four administrative levels

Free-form text, no newlines

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-13}

\"[`address-level3`]{#attr-fe-autocomplete-address-level3 .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

The [third administrative level](#more-on-address-levels), in addresses
with three or more administrative levels

Free-form text, no newlines

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-14}

\"[`address-level2`]{#attr-fe-autocomplete-address-level2 .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

The [second administrative level](#more-on-address-levels), in addresses
with two or more administrative levels; in the countries with two
administrative levels, this would typically be the city, town, village,
or other locality within which the relevant street address is found

Free-form text, no newlines

Cambridge

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-15}

\"[`address-level1`]{#attr-fe-autocomplete-address-level1 .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

The broadest [administrative level](#more-on-address-levels) in the
address, i.e. the province within which the locality is found; for
example, in the US, this would be the state; in Switzerland it would be
the canton; in the UK, the post town

Free-form text, no newlines

MA

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-16}

\"[`country`]{#attr-fe-autocomplete-country .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Country code

Valid [ISO 3166-1-alpha-2 country
code](https://www.iso.org/iso-3166-country-codes.html)
[\[ISO3166\]](references.html#refsISO3166)

US

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-17}

\"[`country-name`]{#attr-fe-autocomplete-country-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Country name

Free-form text, no newlines; [derived from `country` in some
cases](#autofill-country)

US

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-18}

\"[`postal-code`]{#attr-fe-autocomplete-postal-code .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Postal code, post code, ZIP code, CEDEX code (if CEDEX, append
\"CEDEX\", and the *arrondissement*, if relevant, to the
[`address-level2`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level2-2}](#attr-fe-autocomplete-address-level2)
field)

Free-form text, no newlines

02139

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-19}

\"[`cc-name`]{#attr-fe-autocomplete-cc-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Full name as given on the payment instrument

Free-form text, no newlines

Tim Berners-Lee

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-20}

\"[`cc-given-name`]{#attr-fe-autocomplete-cc-given-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Given name as given on the payment instrument (in some Western cultures,
also known as the *first name*)

Free-form text, no newlines

Tim

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-21}

\"[`cc-additional-name`]{#attr-fe-autocomplete-cc-additional-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Additional names given on the payment instrument (in some Western
cultures, also known as *middle names*, forenames other than the first
name)

Free-form text, no newlines

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-22}

\"[`cc-family-name`]{#attr-fe-autocomplete-cc-family-name .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Family name given on the payment instrument (in some Western cultures,
also known as the *last name* or *surname*)

Free-form text, no newlines

Berners-Lee

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-23}

\"[`cc-number`]{#attr-fe-autocomplete-cc-number .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Code identifying the payment instrument (e.g. the credit card number)

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits
x-internal="ascii-digits"}

4114360123456785

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-24}

\"[`cc-exp`]{#attr-fe-autocomplete-cc-exp .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Expiration date of the payment instrument

[Valid month
string](common-microsyntaxes.html#valid-month-string){#autofilling-form-controls:-the-autocomplete-attribute:valid-month-string}

2014-12

[Month](#control-group-month){#autofilling-form-controls:-the-autocomplete-attribute:control-group-month}

\"[`cc-exp-month`]{#attr-fe-autocomplete-cc-exp-month .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Month component of the expiration date of the payment instrument

[Valid
integer](common-microsyntaxes.html#valid-integer){#autofilling-form-controls:-the-autocomplete-attribute:valid-integer}
in the range 1..12

12

[Numeric](#control-group-numeric){#autofilling-form-controls:-the-autocomplete-attribute:control-group-numeric}

\"[`cc-exp-year`]{#attr-fe-autocomplete-cc-exp-year .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Year component of the expiration date of the payment instrument

[Valid
integer](common-microsyntaxes.html#valid-integer){#autofilling-form-controls:-the-autocomplete-attribute:valid-integer-2}
greater than zero

2014

[Numeric](#control-group-numeric){#autofilling-form-controls:-the-autocomplete-attribute:control-group-numeric-2}

\"[`cc-csc`]{#attr-fe-autocomplete-cc-csc .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Security code for the payment instrument (also known as the card
security code (CSC), card validation code (CVC), card verification value
(CVV), signature panel code (SPC), credit card ID (CCID), etc.)

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-2
x-internal="ascii-digits"}

419

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-25}

\"[`cc-type`]{#attr-fe-autocomplete-cc-type .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Type of payment instrument

Free-form text, no newlines

Visa

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-26}

\"[`transaction-currency`]{#attr-fe-autocomplete-transaction-currency
.dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

The currency that the user would prefer the transaction to use

ISO 4217 currency code [\[ISO4217\]](references.html#refsISO4217)

GBP

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-27}

\"[`transaction-amount`]{#attr-fe-autocomplete-transaction-amount .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

The amount that the user would like for the transaction (e.g. when
entering a bid or sale price)

[Valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#autofilling-form-controls:-the-autocomplete-attribute:valid-floating-point-number}

401.00

[Numeric](#control-group-numeric){#autofilling-form-controls:-the-autocomplete-attribute:control-group-numeric-3}

\"[`language`]{#attr-fe-autocomplete-language .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Preferred language

Valid BCP 47 language tag [\[BCP47\]](references.html#refsBCP47)

en

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-28}

\"[`bday`]{#attr-fe-autocomplete-bday .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Birthday

[Valid date
string](common-microsyntaxes.html#valid-date-string){#autofilling-form-controls:-the-autocomplete-attribute:valid-date-string}

1955-06-08

[Date](#control-group-date){#autofilling-form-controls:-the-autocomplete-attribute:control-group-date}

\"[`bday-day`]{#attr-fe-autocomplete-bday-day .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Day component of birthday

[Valid
integer](common-microsyntaxes.html#valid-integer){#autofilling-form-controls:-the-autocomplete-attribute:valid-integer-3}
in the range 1..31

8

[Numeric](#control-group-numeric){#autofilling-form-controls:-the-autocomplete-attribute:control-group-numeric-4}

\"[`bday-month`]{#attr-fe-autocomplete-bday-month .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Month component of birthday

[Valid
integer](common-microsyntaxes.html#valid-integer){#autofilling-form-controls:-the-autocomplete-attribute:valid-integer-4}
in the range 1..12

6

[Numeric](#control-group-numeric){#autofilling-form-controls:-the-autocomplete-attribute:control-group-numeric-5}

\"[`bday-year`]{#attr-fe-autocomplete-bday-year .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Year component of birthday

[Valid
integer](common-microsyntaxes.html#valid-integer){#autofilling-form-controls:-the-autocomplete-attribute:valid-integer-5}
greater than zero

1955

[Numeric](#control-group-numeric){#autofilling-form-controls:-the-autocomplete-attribute:control-group-numeric-6}

\"[`sex`]{#attr-fe-autocomplete-sex .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Gender identity (e.g. Female, Fa\'afafine)

Free-form text, no newlines

Male

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-29}

\"[`url`]{#attr-fe-autocomplete-url .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Home page or other web page corresponding to the company, person,
address, or contact information in the other fields associated with this
field

[Valid URL
string](https://url.spec.whatwg.org/#valid-url-string){#autofilling-form-controls:-the-autocomplete-attribute:valid-url-string
x-internal="valid-url-string"}

https://www.w3.org/People/Berners-Lee/

[URL](#control-group-url){#autofilling-form-controls:-the-autocomplete-attribute:control-group-url}

\"[`photo`]{#attr-fe-autocomplete-photo .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Photograph, icon, or other image corresponding to the company, person,
address, or contact information in the other fields associated with this
field

[Valid URL
string](https://url.spec.whatwg.org/#valid-url-string){#autofilling-form-controls:-the-autocomplete-attribute:valid-url-string-2
x-internal="valid-url-string"}

https://www.w3.org/Press/Stock/Berners-Lee/2001-europaeum-eighth.jpg

[URL](#control-group-url){#autofilling-form-controls:-the-autocomplete-attribute:control-group-url-2}

\"[`tel`]{#attr-fe-autocomplete-tel .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Full telephone number, including country code

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-3
x-internal="ascii-digits"} and U+0020 SPACE characters, prefixed by a
U+002B PLUS SIGN character (+)

+1 617 253 5702

[Tel](#control-group-tel){#autofilling-form-controls:-the-autocomplete-attribute:control-group-tel}

\"[`tel-country-code`]{#attr-fe-autocomplete-tel-country-code .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Country code component of the telephone number

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-4
x-internal="ascii-digits"} prefixed by a U+002B PLUS SIGN character (+)

+1

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-30}

\"[`tel-national`]{#attr-fe-autocomplete-tel-national .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Telephone number without the county code component, with a
country-internal prefix applied if applicable

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-5
x-internal="ascii-digits"} and U+0020 SPACE characters

617 253 5702

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-31}

\"[`tel-area-code`]{#attr-fe-autocomplete-tel-area-code .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Area code component of the telephone number, with a country-internal
prefix applied if applicable

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-6
x-internal="ascii-digits"}

617

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-32}

\"[`tel-local`]{#attr-fe-autocomplete-tel-local .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Telephone number without the country code and area code components

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-7
x-internal="ascii-digits"}

2535702

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-33}

\"[`tel-local-prefix`]{#attr-fe-autocomplete-tel-local-prefix .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

First part of the component of the telephone number that follows the
area code, when that component is split into two components

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-8
x-internal="ascii-digits"}

253

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-34}

\"[`tel-local-suffix`]{#attr-fe-autocomplete-tel-local-suffix .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Second part of the component of the telephone number that follows the
area code, when that component is split into two components

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-9
x-internal="ascii-digits"}

5702

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-35}

\"[`tel-extension`]{#attr-fe-autocomplete-tel-extension .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Telephone number internal extension code

[ASCII
digits](https://infra.spec.whatwg.org/#ascii-digit){#autofilling-form-controls:-the-autocomplete-attribute:ascii-digits-10
x-internal="ascii-digits"}

1000

[Text](#control-group-text){#autofilling-form-controls:-the-autocomplete-attribute:control-group-text-36}

\"[`email`]{#attr-fe-autocomplete-email .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

Email address

[Valid email
address](input.html#valid-e-mail-address){#autofilling-form-controls:-the-autocomplete-attribute:valid-e-mail-address}

timbl@w3.org

[Username](#control-group-username){#autofilling-form-controls:-the-autocomplete-attribute:control-group-username-2}

\"[`impp`]{#attr-fe-autocomplete-impp .dfn
dfn-for="button/autocomplete,fieldset/autocomplete,input/autocomplete,object/autocomplete,output/autocomplete,select/autocomplete,textarea/autocomplete"
dfn-type="attr-value"}\"

URL representing an instant messaging protocol endpoint (for example,
\"`aim:goim?screenname=example`\" or \"`xmpp:fred@example.net`\")

[Valid URL
string](https://url.spec.whatwg.org/#valid-url-string){#autofilling-form-controls:-the-autocomplete-attribute:valid-url-string-3
x-internal="valid-url-string"}

irc://example.org/timbl,isuser

[URL](#control-group-url){#autofilling-form-controls:-the-autocomplete-attribute:control-group-url-3}

The groups correspond to controls as follows:

[Text]{#control-group-text .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-3}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-2}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-2}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-4}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-3}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-5}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-4}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-2}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-2}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element}](form-elements.html#the-select-element)
    elements

[Multiline]{#control-group-multiline .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-6}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-5}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-3}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-3}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-2}](form-elements.html#the-select-element)
    elements

[Password]{#control-group-password .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-7}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-6}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-4}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-8}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-7}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-3}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-9}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-8}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-4}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-10}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-9}](input.html#attr-input-type)
    attribute in the
    [Password](input.html#password-state-(type=password)){#autofilling-form-controls:-the-autocomplete-attribute:password-state-(type=password)}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-4}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-3}](form-elements.html#the-select-element)
    elements

[URL]{#control-group-url .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-11}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-10}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-5}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-12}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-11}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-5}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-13}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-12}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-6}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-14}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-13}](input.html#attr-input-type)
    attribute in the
    [URL](input.html#url-state-(type=url)){#autofilling-form-controls:-the-autocomplete-attribute:url-state-(type=url)}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-5}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-4}](form-elements.html#the-select-element)
    elements

[Username]{#control-group-username .dfn}[]{#control-group-e-mail}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-15}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-14}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-6}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-16}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-15}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-7}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-17}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-16}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-8}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-18}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-17}](input.html#attr-input-type)
    attribute in the
    [Email](input.html#email-state-(type=email)){#autofilling-form-controls:-the-autocomplete-attribute:email-state-(type=email)}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-6}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-5}](form-elements.html#the-select-element)
    elements

[Tel]{#control-group-tel .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-19}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-18}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-7}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-20}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-19}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-9}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-21}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-20}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-10}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-22}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-21}](input.html#attr-input-type)
    attribute in the
    [Telephone](input.html#telephone-state-(type=tel)){#autofilling-form-controls:-the-autocomplete-attribute:telephone-state-(type=tel)}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-7}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-6}](form-elements.html#the-select-element)
    elements

[Numeric]{#control-group-numeric .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-23}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-22}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-8}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-24}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-23}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-11}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-25}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-24}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-12}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-26}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-25}](input.html#attr-input-type)
    attribute in the
    [Number](input.html#number-state-(type=number)){#autofilling-form-controls:-the-autocomplete-attribute:number-state-(type=number)}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-8}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-7}](form-elements.html#the-select-element)
    elements

[Month]{#control-group-month .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-27}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-26}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-9}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-28}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-27}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-13}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-29}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-28}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-14}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-30}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-29}](input.html#attr-input-type)
    attribute in the
    [Month](input.html#month-state-(type=month)){#autofilling-form-controls:-the-autocomplete-attribute:month-state-(type=month)}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-9}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-8}](form-elements.html#the-select-element)
    elements

[Date]{#control-group-date .dfn}
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-31}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-30}](input.html#attr-input-type)
    attribute in the
    [Hidden](input.html#hidden-state-(type=hidden)){#autofilling-form-controls:-the-autocomplete-attribute:hidden-state-(type=hidden)-10}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-32}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-31}](input.html#attr-input-type)
    attribute in the
    [Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-15}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-33}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-32}](input.html#attr-input-type)
    attribute in the
    [Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofilling-form-controls:-the-autocomplete-attribute:text-(type=text)-state-and-search-state-(type=search)-16}
    state
:   [`input`{#autofilling-form-controls:-the-autocomplete-attribute:the-input-element-34}](input.html#the-input-element)
    elements with a
    [`type`{#autofilling-form-controls:-the-autocomplete-attribute:attr-input-type-33}](input.html#attr-input-type)
    attribute in the
    [Date](input.html#date-state-(type=date)){#autofilling-form-controls:-the-autocomplete-attribute:date-state-(type=date)}
    state
:   [`textarea`{#autofilling-form-controls:-the-autocomplete-attribute:the-textarea-element-10}](form-elements.html#the-textarea-element)
    elements
:   [`select`{#autofilling-form-controls:-the-autocomplete-attribute:the-select-element-9}](form-elements.html#the-select-element)
    elements

**Address levels**: The
\"[`address-level1`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level1-2}](#attr-fe-autocomplete-address-level1)\"
--
\"[`address-level4`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level4-2}](#attr-fe-autocomplete-address-level4)\"
fields are used to describe the locality of the street address.
Different locales have different numbers of levels. For example, the US
uses two levels (state and town), the UK uses one or two depending on
the address (the post town, and in some cases the locality), and China
can use three (province, city, district). The
\"[`address-level1`{#autofilling-form-controls:-the-autocomplete-attribute:attr-fe-autocomplete-address-level1-3}](#attr-fe-autocomplete-address-level1)\"
field represents the widest administrative division. Different locales
order the fields in different ways; for example, in the US the town
(level 2) precedes the state (level 1); while in Japan the prefecture
(level 1) precedes the city (level 2) which precedes the district (level
3). Authors are encouraged to provide forms that are presented in a way
that matches the country\'s conventions (hiding, showing, and
rearranging fields accordingly as the user changes the country).

###### [4.10.18.7.2]{.secno} []{#processing-model-3}Processing model[](#autofill-processing-model){.self-link} {#autofill-processing-model}

Each
[`input`{#autofill-processing-model:the-input-element}](input.html#the-input-element)
element to which the
[`autocomplete`{#autofill-processing-model:attr-fe-autocomplete}](#attr-fe-autocomplete)
attribute
[applies](input.html#concept-input-apply){#autofill-processing-model:concept-input-apply},
each
[`select`{#autofill-processing-model:the-select-element}](form-elements.html#the-select-element)
element, and each
[`textarea`{#autofill-processing-model:the-textarea-element}](form-elements.html#the-textarea-element)
element, has an [autofill hint set]{#autofill-hint-set .dfn}, an
[autofill scope]{#autofill-scope .dfn}, an [autofill field
name]{#autofill-field-name .dfn}, a [non-autofill credential
type]{#non-autofill-credential-type .dfn}, and an [IDL-exposed autofill
value]{#idl-exposed-autofill-value .dfn}.

The [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name}
specifies the specific kind of data expected in the field, e.g.
\"[`street-address`{#autofill-processing-model:attr-fe-autocomplete-street-address}](#attr-fe-autocomplete-street-address)\"
or
\"[`cc-exp`{#autofill-processing-model:attr-fe-autocomplete-cc-exp}](#attr-fe-autocomplete-cc-exp)\".

The [autofill hint
set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set}
identifies what address or contact information type the user agent is to
look at, e.g.
\"[`shipping`{#autofill-processing-model:attr-fe-autocomplete-shipping}](#attr-fe-autocomplete-shipping)
[`fax`{#autofill-processing-model:attr-fe-autocomplete-fax}](#attr-fe-autocomplete-fax)\"
or
\"[`billing`{#autofill-processing-model:attr-fe-autocomplete-billing}](#attr-fe-autocomplete-billing)\".

The [non-autofill credential
type](#non-autofill-credential-type){#autofill-processing-model:non-autofill-credential-type}
identifies a type of
[credential](https://w3c.github.io/webappsec-credential-management/#credential){#autofill-processing-model:credman-credential
x-internal="credman-credential"} that may be offered by the user agent
when the user interacts with the field alongside other [autofill
field](#autofill-field){#autofill-processing-model:autofill-field}
values. If this value is \"`webauthn`\" instead of null, selecting a
credential of that type will resolve a pending
[`conditional`{#autofill-processing-model:conditional-mediation}](https://w3c.github.io/webappsec-credential-management/#dom-credentialmediationrequirement-conditional){x-internal="conditional-mediation"}
mediation
[`navigator.credentials.get()`{#autofill-processing-model:navigator.credentials.get()}](https://w3c.github.io/webappsec-credential-management/#dom-credentialscontainer-get){x-internal="navigator.credentials.get()"}
request, instead of autofilling the field.

::: example
For example, a sign-in page could instruct the user agent to either
autofill a saved password, or show a [public key
credential](https://w3c.github.io/webauthn/#public-key-credential){#autofill-processing-model:public-key-credential
x-internal="public-key-credential"} that will resolve a pending
[`navigator.credentials.get()`{#autofill-processing-model:navigator.credentials.get()-2}](https://w3c.github.io/webappsec-credential-management/#dom-credentialscontainer-get){x-internal="navigator.credentials.get()"}
request. A user can select either to sign-in.

``` html
<input name=password type=password autocomplete="current-password webauthn">
```
:::

The [autofill
scope](#autofill-scope){#autofill-processing-model:autofill-scope}
identifies the group of fields whose information concerns the same
subject, and consists of the [autofill hint
set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-2}
with, if applicable, the \"`section-*`\" prefix, e.g. \"`billing`\",
\"`section-parent shipping`\", or \"`section-child shipping home`\".

These values are defined as the result of running the following
algorithm:

1.  If the element has no
    [`autocomplete`{#autofill-processing-model:attr-fe-autocomplete-2}](#attr-fe-autocomplete)
    attribute, then jump to the step labeled *default*.

2.  Let `tokens`{.variable} be the result of [splitting the attribute\'s
    value on ASCII
    whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#autofill-processing-model:split-a-string-on-spaces
    x-internal="split-a-string-on-spaces"}.

3.  If `tokens`{.variable} is empty, then jump to the step labeled
    *default*.

4.  Let `index`{.variable} be the index of the last token in
    `tokens`{.variable}.

5.  Let `field`{.variable} be the `index`{.variable}th token in
    `tokens`{.variable}.

6.  Set the `category`{.variable}, `maximum tokens`{.variable} pair to
    the result of [determining a field\'s
    category](#determine-a-field's-category){#autofill-processing-model:determine-a-field's-category}
    given `field`{.variable}.

7.  If `category`{.variable} is null, then jump to the step labeled
    *default*.

8.  If the number of tokens in `tokens`{.variable} is greater than
    `maximum tokens`{.variable}, then jump to the step labeled
    *default*.

9.  If `category`{.variable} is Off or Automatic but the element\'s
    [`autocomplete`{#autofill-processing-model:attr-fe-autocomplete-3}](#attr-fe-autocomplete)
    attribute is wearing the [autofill anchor
    mantle](#autofill-anchor-mantle){#autofill-processing-model:autofill-anchor-mantle},
    then jump to the step labeled *default*.

10. If `category`{.variable} is Off, let the element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-2}
    be the string \"`off`\", let its [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-3}
    be empty, and let its [IDL-exposed autofill
    value](#idl-exposed-autofill-value){#autofill-processing-model:idl-exposed-autofill-value}
    be the string \"`off`\". Then, return.

11. If `category`{.variable} is Automatic, let the element\'s [autofill
    field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-3}
    be the string \"`on`\", let its [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-4}
    be empty, and let its [IDL-exposed autofill
    value](#idl-exposed-autofill-value){#autofill-processing-model:idl-exposed-autofill-value-2}
    be the string \"`on`\". Then, return.

12. Let `scope tokens`{.variable} be an empty list.

13. Let `hint tokens`{.variable} be an empty set.

14. Let `credential type`{.variable} be null.

15. Let `IDL value`{.variable} have the same value as
    `field`{.variable}.

16. If `category`{.variable} is Credential and the `index`{.variable}th
    token in `tokens`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofill-processing-model:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for
    \"[`webauthn`{#autofill-processing-model:attr-fe-autocomplete-webauthn}](#attr-fe-autocomplete-webauthn)\",
    then run the substeps that follow:

    1.  Set `credential type`{.variable} to \"`webauthn`\".

    2.  If the `index`{.variable}th token in `tokens`{.variable} is the
        first entry, then skip to the step labeled *done*.

    3.  Decrement `index`{.variable} by one.

    4.  Set the `category`{.variable}, `maximum tokens`{.variable} pair
        to the result of [determining a field\'s
        category](#determine-a-field's-category){#autofill-processing-model:determine-a-field's-category-2}
        given the `index`{.variable}th token in `tokens`{.variable}.

    5.  If `category`{.variable} is not Normal and `category`{.variable}
        is not Contact, then jump to the step labeled *default*.

    6.  If `index`{.variable} is greater than
        `maximum tokens`{.variable} minus one (i.e. if the number of
        remaining tokens is greater than `maximum tokens`{.variable}),
        then jump to the step labeled *default*.

    7.  Set `IDL value`{.variable} to the concatenation of the
        `index`{.variable}th token in `tokens`{.variable}, a U+0020
        SPACE character, and the previous value of
        `IDL value`{.variable}.

17. If the `index`{.variable}th token in `tokens`{.variable} is the
    first entry, then skip to the step labeled *done*.

18. Decrement `index`{.variable} by one.

19. If `category`{.variable} is Contact and the `index`{.variable}th
    token in `tokens`{.variable} is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofill-processing-model:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for one of the strings in
    the following list, then run the substeps that follow:

    - \"[`home`{#autofill-processing-model:attr-fe-autocomplete-home}](#attr-fe-autocomplete-home)\"
    - \"[`work`{#autofill-processing-model:attr-fe-autocomplete-work}](#attr-fe-autocomplete-work)\"
    - \"[`mobile`{#autofill-processing-model:attr-fe-autocomplete-mobile}](#attr-fe-autocomplete-mobile)\"
    - \"[`fax`{#autofill-processing-model:attr-fe-autocomplete-fax-2}](#attr-fe-autocomplete-fax)\"
    - \"[`pager`{#autofill-processing-model:attr-fe-autocomplete-pager}](#attr-fe-autocomplete-pager)\"

    The substeps are:

    1.  Let `contact`{.variable} be the matching string from the list
        above.

    2.  Insert `contact`{.variable} at the start of
        `scope tokens`{.variable}.

    3.  Add `contact`{.variable} to `hint tokens`{.variable}.

    4.  Let `IDL value`{.variable} be the concatenation of
        `contact`{.variable}, a U+0020 SPACE character, and the previous
        value of `IDL value`{.variable}.

    5.  If the `index`{.variable}th entry in `tokens`{.variable} is the
        first entry, then skip to the step labeled *done*.

    6.  Decrement `index`{.variable} by one.

20. If the `index`{.variable}th token in `tokens`{.variable} is an
    [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofill-processing-model:ascii-case-insensitive-3
    x-internal="ascii-case-insensitive"} match for one of the strings in
    the following list, then run the substeps that follow:

    - \"[`shipping`{#autofill-processing-model:attr-fe-autocomplete-shipping-2}](#attr-fe-autocomplete-shipping)\"
    - \"[`billing`{#autofill-processing-model:attr-fe-autocomplete-billing-2}](#attr-fe-autocomplete-billing)\"

    The substeps are:

    1.  Let `mode`{.variable} be the matching string from the list
        above.

    2.  Insert `mode`{.variable} at the start of
        `scope tokens`{.variable}.

    3.  Add `mode`{.variable} to `hint tokens`{.variable}.

    4.  Let `IDL value`{.variable} be the concatenation of
        `mode`{.variable}, a U+0020 SPACE character, and the previous
        value of `IDL value`{.variable}.

    5.  If the `index`{.variable}th entry in `tokens`{.variable} is the
        first entry, then skip to the step labeled *done*.

    6.  Decrement `index`{.variable} by one.

21. If the `index`{.variable}th entry in `tokens`{.variable} is not the
    first entry, then jump to the step labeled *default*.

22. If the first eight characters of the `index`{.variable}th token in
    `tokens`{.variable} are not an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofill-processing-model:ascii-case-insensitive-4
    x-internal="ascii-case-insensitive"} match for the string
    \"[`section-`{#autofill-processing-model:attr-fe-autocomplete-section}](#attr-fe-autocomplete-section)\",
    then jump to the step labeled *default*.

23. Let `section`{.variable} be the `index`{.variable}th token in
    `tokens`{.variable}, [converted to ASCII
    lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#autofill-processing-model:converted-to-ascii-lowercase
    x-internal="converted-to-ascii-lowercase"}.

24. Insert `section`{.variable} at the start of
    `scope tokens`{.variable}.

25. Let `IDL value`{.variable} be the concatenation of
    `section`{.variable}, a U+0020 SPACE character, and the previous
    value of `IDL value`{.variable}.

26. *Done*: Let the element\'s [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-5}
    be `hint tokens`{.variable}.

27. Let the element\'s [non-autofill credential
    type](#non-autofill-credential-type){#autofill-processing-model:non-autofill-credential-type-2}
    be `credential type`{.variable}.

28. Let the element\'s [autofill
    scope](#autofill-scope){#autofill-processing-model:autofill-scope-2}
    be `scope tokens`{.variable}.

29. Let the element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-4}
    be `field`{.variable}.

30. Let the element\'s [IDL-exposed autofill
    value](#idl-exposed-autofill-value){#autofill-processing-model:idl-exposed-autofill-value-3}
    be `IDL value`{.variable}.

31. Return.

32. *Default*: Let the element\'s [IDL-exposed autofill
    value](#idl-exposed-autofill-value){#autofill-processing-model:idl-exposed-autofill-value-4}
    be the empty string, and its [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-6}
    and [autofill
    scope](#autofill-scope){#autofill-processing-model:autofill-scope-3}
    be empty.

33. If the element\'s
    [`autocomplete`{#autofill-processing-model:attr-fe-autocomplete-4}](#attr-fe-autocomplete)
    attribute is wearing the [autofill anchor
    mantle](#autofill-anchor-mantle){#autofill-processing-model:autofill-anchor-mantle-2},
    then let the element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-5}
    be the empty string and return.

34. Let `form`{.variable} be the element\'s [form
    owner](#form-owner){#autofill-processing-model:form-owner}, if any,
    or null otherwise.

35. If `form`{.variable} is not null and `form`{.variable}\'s
    [`autocomplete`{#autofill-processing-model:attr-form-autocomplete}](forms.html#attr-form-autocomplete)
    attribute is in the
    [Off](forms.html#attr-form-autocomplete-off-state){#autofill-processing-model:attr-form-autocomplete-off-state}
    state, then let the element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-6}
    be
    \"[`off`{#autofill-processing-model:attr-fe-autocomplete-off}](#attr-fe-autocomplete-off)\".

    Otherwise, let the element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-7}
    be
    \"[`on`{#autofill-processing-model:attr-fe-autocomplete-on}](#attr-fe-autocomplete-on)\".

To [determine a field\'s category]{#determine-a-field's-category .dfn},
given `field`{.variable}:

1.  If the `field`{.variable} is not an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#autofill-processing-model:ascii-case-insensitive-5
    x-internal="ascii-case-insensitive"} match for one of the tokens
    given in the first column of the following table, return the pair
    (null, null).

    Token

    Maximum number of tokens

    Category

    \"[`off`{#autofill-processing-model:attr-fe-autocomplete-off-2}](#attr-fe-autocomplete-off)\"

    1

    Off

    \"[`on`{#autofill-processing-model:attr-fe-autocomplete-on-2}](#attr-fe-autocomplete-on)\"

    1

    Automatic

    \"[`name`{#autofill-processing-model:attr-fe-autocomplete-name}](#attr-fe-autocomplete-name)\"

    3

    Normal

    \"[`honorific-prefix`{#autofill-processing-model:attr-fe-autocomplete-honorific-prefix}](#attr-fe-autocomplete-honorific-prefix)\"

    3

    Normal

    \"[`given-name`{#autofill-processing-model:attr-fe-autocomplete-given-name}](#attr-fe-autocomplete-given-name)\"

    3

    Normal

    \"[`additional-name`{#autofill-processing-model:attr-fe-autocomplete-additional-name}](#attr-fe-autocomplete-additional-name)\"

    3

    Normal

    \"[`family-name`{#autofill-processing-model:attr-fe-autocomplete-family-name}](#attr-fe-autocomplete-family-name)\"

    3

    Normal

    \"[`honorific-suffix`{#autofill-processing-model:attr-fe-autocomplete-honorific-suffix}](#attr-fe-autocomplete-honorific-suffix)\"

    3

    Normal

    \"[`nickname`{#autofill-processing-model:attr-fe-autocomplete-nickname}](#attr-fe-autocomplete-nickname)\"

    3

    Normal

    \"[`organization-title`{#autofill-processing-model:attr-fe-autocomplete-organization-title}](#attr-fe-autocomplete-organization-title)\"

    3

    Normal

    \"[`username`{#autofill-processing-model:attr-fe-autocomplete-username}](#attr-fe-autocomplete-username)\"

    3

    Normal

    \"[`new-password`{#autofill-processing-model:attr-fe-autocomplete-new-password}](#attr-fe-autocomplete-new-password)\"

    3

    Normal

    \"[`current-password`{#autofill-processing-model:attr-fe-autocomplete-current-password}](#attr-fe-autocomplete-current-password)\"

    3

    Normal

    \"[`one-time-code`{#autofill-processing-model:attr-fe-autocomplete-one-time-code}](#attr-fe-autocomplete-one-time-code)\"

    3

    Normal

    \"[`organization`{#autofill-processing-model:attr-fe-autocomplete-organization}](#attr-fe-autocomplete-organization)\"

    3

    Normal

    \"[`street-address`{#autofill-processing-model:attr-fe-autocomplete-street-address-2}](#attr-fe-autocomplete-street-address)\"

    3

    Normal

    \"[`address-line1`{#autofill-processing-model:attr-fe-autocomplete-address-line1}](#attr-fe-autocomplete-address-line1)\"

    3

    Normal

    \"[`address-line2`{#autofill-processing-model:attr-fe-autocomplete-address-line2}](#attr-fe-autocomplete-address-line2)\"

    3

    Normal

    \"[`address-line3`{#autofill-processing-model:attr-fe-autocomplete-address-line3}](#attr-fe-autocomplete-address-line3)\"

    3

    Normal

    \"[`address-level4`{#autofill-processing-model:attr-fe-autocomplete-address-level4}](#attr-fe-autocomplete-address-level4)\"

    3

    Normal

    \"[`address-level3`{#autofill-processing-model:attr-fe-autocomplete-address-level3}](#attr-fe-autocomplete-address-level3)\"

    3

    Normal

    \"[`address-level2`{#autofill-processing-model:attr-fe-autocomplete-address-level2}](#attr-fe-autocomplete-address-level2)\"

    3

    Normal

    \"[`address-level1`{#autofill-processing-model:attr-fe-autocomplete-address-level1}](#attr-fe-autocomplete-address-level1)\"

    3

    Normal

    \"[`country`{#autofill-processing-model:attr-fe-autocomplete-country}](#attr-fe-autocomplete-country)\"

    3

    Normal

    \"[`country-name`{#autofill-processing-model:attr-fe-autocomplete-country-name}](#attr-fe-autocomplete-country-name)\"

    3

    Normal

    \"[`postal-code`{#autofill-processing-model:attr-fe-autocomplete-postal-code}](#attr-fe-autocomplete-postal-code)\"

    3

    Normal

    \"[`cc-name`{#autofill-processing-model:attr-fe-autocomplete-cc-name}](#attr-fe-autocomplete-cc-name)\"

    3

    Normal

    \"[`cc-given-name`{#autofill-processing-model:attr-fe-autocomplete-cc-given-name}](#attr-fe-autocomplete-cc-given-name)\"

    3

    Normal

    \"[`cc-additional-name`{#autofill-processing-model:attr-fe-autocomplete-cc-additional-name}](#attr-fe-autocomplete-cc-additional-name)\"

    3

    Normal

    \"[`cc-family-name`{#autofill-processing-model:attr-fe-autocomplete-cc-family-name}](#attr-fe-autocomplete-cc-family-name)\"

    3

    Normal

    \"[`cc-number`{#autofill-processing-model:attr-fe-autocomplete-cc-number}](#attr-fe-autocomplete-cc-number)\"

    3

    Normal

    \"[`cc-exp`{#autofill-processing-model:attr-fe-autocomplete-cc-exp-2}](#attr-fe-autocomplete-cc-exp)\"

    3

    Normal

    \"[`cc-exp-month`{#autofill-processing-model:attr-fe-autocomplete-cc-exp-month}](#attr-fe-autocomplete-cc-exp-month)\"

    3

    Normal

    \"[`cc-exp-year`{#autofill-processing-model:attr-fe-autocomplete-cc-exp-year}](#attr-fe-autocomplete-cc-exp-year)\"

    3

    Normal

    \"[`cc-csc`{#autofill-processing-model:attr-fe-autocomplete-cc-csc}](#attr-fe-autocomplete-cc-csc)\"

    3

    Normal

    \"[`cc-type`{#autofill-processing-model:attr-fe-autocomplete-cc-type}](#attr-fe-autocomplete-cc-type)\"

    3

    Normal

    \"[`transaction-currency`{#autofill-processing-model:attr-fe-autocomplete-transaction-currency}](#attr-fe-autocomplete-transaction-currency)\"

    3

    Normal

    \"[`transaction-amount`{#autofill-processing-model:attr-fe-autocomplete-transaction-amount}](#attr-fe-autocomplete-transaction-amount)\"

    3

    Normal

    \"[`language`{#autofill-processing-model:attr-fe-autocomplete-language}](#attr-fe-autocomplete-language)\"

    3

    Normal

    \"[`bday`{#autofill-processing-model:attr-fe-autocomplete-bday}](#attr-fe-autocomplete-bday)\"

    3

    Normal

    \"[`bday-day`{#autofill-processing-model:attr-fe-autocomplete-bday-day}](#attr-fe-autocomplete-bday-day)\"

    3

    Normal

    \"[`bday-month`{#autofill-processing-model:attr-fe-autocomplete-bday-month}](#attr-fe-autocomplete-bday-month)\"

    3

    Normal

    \"[`bday-year`{#autofill-processing-model:attr-fe-autocomplete-bday-year}](#attr-fe-autocomplete-bday-year)\"

    3

    Normal

    \"[`sex`{#autofill-processing-model:attr-fe-autocomplete-sex}](#attr-fe-autocomplete-sex)\"

    3

    Normal

    \"[`url`{#autofill-processing-model:attr-fe-autocomplete-url}](#attr-fe-autocomplete-url)\"

    3

    Normal

    \"[`photo`{#autofill-processing-model:attr-fe-autocomplete-photo}](#attr-fe-autocomplete-photo)\"

    3

    Normal

    \"[`tel`{#autofill-processing-model:attr-fe-autocomplete-tel}](#attr-fe-autocomplete-tel)\"

    4

    Contact

    \"[`tel-country-code`{#autofill-processing-model:attr-fe-autocomplete-tel-country-code}](#attr-fe-autocomplete-tel-country-code)\"

    4

    Contact

    \"[`tel-national`{#autofill-processing-model:attr-fe-autocomplete-tel-national}](#attr-fe-autocomplete-tel-national)\"

    4

    Contact

    \"[`tel-area-code`{#autofill-processing-model:attr-fe-autocomplete-tel-area-code}](#attr-fe-autocomplete-tel-area-code)\"

    4

    Contact

    \"[`tel-local`{#autofill-processing-model:attr-fe-autocomplete-tel-local}](#attr-fe-autocomplete-tel-local)\"

    4

    Contact

    \"[`tel-local-prefix`{#autofill-processing-model:attr-fe-autocomplete-tel-local-prefix}](#attr-fe-autocomplete-tel-local-prefix)\"

    4

    Contact

    \"[`tel-local-suffix`{#autofill-processing-model:attr-fe-autocomplete-tel-local-suffix}](#attr-fe-autocomplete-tel-local-suffix)\"

    4

    Contact

    \"[`tel-extension`{#autofill-processing-model:attr-fe-autocomplete-tel-extension}](#attr-fe-autocomplete-tel-extension)\"

    4

    Contact

    \"[`email`{#autofill-processing-model:attr-fe-autocomplete-email}](#attr-fe-autocomplete-email)\"

    4

    Contact

    \"[`impp`{#autofill-processing-model:attr-fe-autocomplete-impp}](#attr-fe-autocomplete-impp)\"

    4

    Contact

    \"[`webauthn`{#autofill-processing-model:attr-fe-autocomplete-webauthn-2}](#attr-fe-autocomplete-webauthn)\"

    5

    Credential

2.  Otherwise, let `maximum tokens`{.variable} and `category`{.variable}
    be the values of the cells in the second and third columns of that
    row respectively.

3.  Return the pair (`category`{.variable},
    `maximum tokens`{.variable}).

------------------------------------------------------------------------

For the purposes of autofill, a [control\'s data]{#control's-data .dfn}
depends on the kind of control:

An [`input`{#autofill-processing-model:the-input-element-2}](input.html#the-input-element) element with its [`type`{#autofill-processing-model:attr-input-type}](input.html#attr-input-type) attribute in the [Email](input.html#email-state-(type=email)){#autofill-processing-model:email-state-(type=email)} state and with the [`multiple`{#autofill-processing-model:attr-input-multiple}](input.html#attr-input-multiple) attribute specified
:   The element\'s
    [value*s*](#concept-fe-values){#autofill-processing-model:concept-fe-values}.

Any other [`input`{#autofill-processing-model:the-input-element-3}](input.html#the-input-element) element\
A [`textarea`{#autofill-processing-model:the-textarea-element-2}](form-elements.html#the-textarea-element) element
:   The element\'s
    [value](#concept-fe-value){#autofill-processing-model:concept-fe-value}.

A [`select`{#autofill-processing-model:the-select-element-2}](form-elements.html#the-select-element) element with its [`multiple`{#autofill-processing-model:attr-select-multiple}](form-elements.html#attr-select-multiple) attribute specified
:   The
    [`option`{#autofill-processing-model:the-option-element}](form-elements.html#the-option-element)
    elements in the
    [`select`{#autofill-processing-model:the-select-element-3}](form-elements.html#the-select-element)
    element\'s [list of
    options](form-elements.html#concept-select-option-list){#autofill-processing-model:concept-select-option-list}
    that have their
    [selectedness](form-elements.html#concept-option-selectedness){#autofill-processing-model:concept-option-selectedness}
    set to true.

Any other [`select`{#autofill-processing-model:the-select-element-4}](form-elements.html#the-select-element) element
:   The
    [`option`{#autofill-processing-model:the-option-element-2}](form-elements.html#the-option-element)
    element in the
    [`select`{#autofill-processing-model:the-select-element-5}](form-elements.html#the-select-element)
    element\'s [list of
    options](form-elements.html#concept-select-option-list){#autofill-processing-model:concept-select-option-list-2}
    that has its
    [selectedness](form-elements.html#concept-option-selectedness){#autofill-processing-model:concept-option-selectedness-2}
    set to true.

------------------------------------------------------------------------

How to process the [autofill hint
set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-7},
[autofill
scope](#autofill-scope){#autofill-processing-model:autofill-scope-4},
and [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-8}
depends on the mantle that the
[`autocomplete`{#autofill-processing-model:attr-fe-autocomplete-5}](#attr-fe-autocomplete)
attribute is wearing.

When wearing the [autofill expectation mantle](#autofill-expectation-mantle){#autofill-processing-model:autofill-expectation-mantle}\...

:   When an element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-9}
    is
    \"[`off`{#autofill-processing-model:attr-fe-autocomplete-off-3}](#attr-fe-autocomplete-off)\",
    the user agent should not remember the [control\'s
    data](#control's-data){#autofill-processing-model:control's-data},
    and should not offer past values to the user.

    In addition, when an element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-10}
    is
    \"[`off`{#autofill-processing-model:attr-fe-autocomplete-off-4}](#attr-fe-autocomplete-off)\",
    [values are reset](browsing-the-web.html#history-autocomplete) when
    [reactivating a
    document](browsing-the-web.html#reactivate-a-document){#autofill-processing-model:reactivate-a-document}.

    ::: example
    Banks frequently do not want UAs to prefill login information:

    ``` html
    <p><label>Account: <input type="text" name="ac" autocomplete="off"></label></p>
    <p><label>PIN: <input type="password" name="pin" autocomplete="off"></label></p>
    ```
    :::

    When an element\'s [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-11}
    is *not*
    \"[`off`{#autofill-processing-model:attr-fe-autocomplete-off-5}](#attr-fe-autocomplete-off)\",
    the user agent may store the [control\'s
    data](#control's-data){#autofill-processing-model:control's-data-2},
    and may offer previously stored values to the user.

    ::: example
    For example, suppose a user visits a page with this control:

    ``` html
    <select name="country">
     <option>Afghanistan
     <option>Albania
     <option>Algeria
     <option>Andorra
     <option>Angola
     <option>Antigua and Barbuda
     <option>Argentina
     <option>Armenia
     <!-- ... -->
     <option>Yemen
     <option>Zambia
     <option>Zimbabwe
    </select>
    ```

    This might render as follows:

    ![A drop-down control with a long alphabetical list of
    countries.](/images/select-country-1.png)

    Suppose that on the first visit to this page, the user selects
    \"Zambia\". On the second visit, the user agent could duplicate the
    entry for Zambia at the top of the list, so that the interface
    instead looks like this:

    ![The same drop-down control with the alphabetical list of
    countries, but with Zambia as an entry at the
    top.](/images/select-country-2.png)
    :::

    When the [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-12}
    is
    \"[`on`{#autofill-processing-model:attr-fe-autocomplete-on-3}](#attr-fe-autocomplete-on)\",
    the user agent should attempt to use heuristics to determine the
    most appropriate values to offer the user, e.g. based on the
    element\'s
    [`name`{#autofill-processing-model:attr-fe-name}](#attr-fe-name)
    value, the position of the element in its
    [tree](https://dom.spec.whatwg.org/#concept-tree){#autofill-processing-model:tree
    x-internal="tree"}, what other fields exist in the form, and so
    forth.

    When the [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-13}
    is one of the names of the [autofill
    fields](#autofill-field){#autofill-processing-model:autofill-field-2}
    described above, the user agent should provide suggestions that
    match the meaning of the field name as given in the table earlier in
    this section. The [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-8}
    should be used to select amongst multiple possible suggestions.

    For example, if a user once entered one address into fields that
    used the
    \"[`shipping`{#autofill-processing-model:attr-fe-autocomplete-shipping-3}](#attr-fe-autocomplete-shipping)\"
    keyword, and another address into fields that used the
    \"[`billing`{#autofill-processing-model:attr-fe-autocomplete-billing-3}](#attr-fe-autocomplete-billing)\"
    keyword, then in subsequent forms only the first address would be
    suggested for form controls whose [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-9}
    contains the keyword
    \"[`shipping`{#autofill-processing-model:attr-fe-autocomplete-shipping-4}](#attr-fe-autocomplete-shipping)\".
    Both addresses might be suggested, however, for address-related form
    controls whose [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-10}
    does not contain either keyword.

When wearing the [autofill anchor mantle](#autofill-anchor-mantle){#autofill-processing-model:autofill-anchor-mantle-3}\...

:   When the [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-14}
    is not the empty string, then the user agent must act as if the user
    had specified the [control\'s
    data](#control's-data){#autofill-processing-model:control's-data-3}
    for the given [autofill hint
    set](#autofill-hint-set){#autofill-processing-model:autofill-hint-set-11},
    [autofill
    scope](#autofill-scope){#autofill-processing-model:autofill-scope-5},
    and [autofill field
    name](#autofill-field-name){#autofill-processing-model:autofill-field-name-15}
    combination.

When the user agent [autofills form controls]{#concept-fe-autofill
.dfn}, elements with the same [form
owner](#form-owner){#autofill-processing-model:form-owner-2} and the
same [autofill
scope](#autofill-scope){#autofill-processing-model:autofill-scope-6}
must use data relating to the same person, address, payment instrument,
and contact details. [When a user agent autofills
\"[`country`{#autofill-processing-model:attr-fe-autocomplete-country-2}](#attr-fe-autocomplete-country)\"
and
\"[`country-name`{#autofill-processing-model:attr-fe-autocomplete-country-name-2}](#attr-fe-autocomplete-country-name)\"
fields with the same [form
owner](#form-owner){#autofill-processing-model:form-owner-3} and
[autofill
scope](#autofill-scope){#autofill-processing-model:autofill-scope-7},
and the user agent has a value for the
[`country`{#autofill-processing-model:attr-fe-autocomplete-country-3}](#attr-fe-autocomplete-country)\"
field(s), then the
\"[`country-name`{#autofill-processing-model:attr-fe-autocomplete-country-name-3}](#attr-fe-autocomplete-country-name)\"
field(s) must be filled using a human-readable name for the same
country.]{#autofill-country} When a user agent fills in multiple fields
at once, all fields with the same [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-16},
[form owner](#form-owner){#autofill-processing-model:form-owner-4}, and
[autofill
scope](#autofill-scope){#autofill-processing-model:autofill-scope-8}
must be filled with the same value.

Suppose a user agent knows of two phone numbers, +1 555 123 1234 and +1
555 666 7777. It would not be conforming for the user agent to fill a
field with `autocomplete="shipping tel-local-prefix"` with the value
\"123\" and another field in the same form with
`autocomplete="shipping tel-local-suffix"` with the value \"7777\". The
only valid prefilled values given the aforementioned information would
be \"123\" and \"1234\", or \"666\" and \"7777\", respectively.

Similarly, if a form for some reason contained both a
\"[`cc-exp`{#autofill-processing-model:attr-fe-autocomplete-cc-exp-3}](#attr-fe-autocomplete-cc-exp)\"
field and a
\"[`cc-exp-month`{#autofill-processing-model:attr-fe-autocomplete-cc-exp-month-2}](#attr-fe-autocomplete-cc-exp-month)\"
field, and the user agent prefilled the form, then the month component
of the former would have to match the latter.

::: example
This requirement interacts with the [autofill anchor
mantle](#autofill-anchor-mantle){#autofill-processing-model:autofill-anchor-mantle-4}
also. Consider the following markup snippet:

``` html
<form>
 <input type=hidden autocomplete="nickname" value="TreePlate">
 <input type=text autocomplete="nickname">
</form>
```

The only value that a conforming user agent could suggest in the text
control is \"TreePlate\", the value given by the hidden
[`input`{#autofill-processing-model:the-input-element-4}](input.html#the-input-element)
element.
:::

The \"`section-*`\" tokens in the [autofill
scope](#autofill-scope){#autofill-processing-model:autofill-scope-9} are
opaque; user agents must not attempt to derive meaning from the precise
values of these tokens.

For example, it would not be conforming if the user agent decided that
it should offer the address it knows to be the user\'s daughter\'s
address for \"`section-child`\" and the addresses it knows to be the
user\'s spouses\' addresses for \"`section-spouse`\".

The autocompletion mechanism must be implemented by the user agent
acting as if the user had modified the [control\'s
data](#control's-data){#autofill-processing-model:control's-data-4}, and
must be done at a time where the element is
*[mutable](#concept-fe-mutable)* (e.g. just after the element has been
inserted into the document, or when the user agent [stops
parsing](parsing.html#stop-parsing){#autofill-processing-model:stop-parsing}).
User agents must only prefill controls using values that the user could
have entered.

For example, if a
[`select`{#autofill-processing-model:the-select-element-6}](form-elements.html#the-select-element)
element only has
[`option`{#autofill-processing-model:the-option-element-3}](form-elements.html#the-option-element)
elements with values \"Steve\" and \"Rebecca\", \"Jay\", and \"Bob\",
and has an [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-17}
\"[`given-name`{#autofill-processing-model:attr-fe-autocomplete-given-name-2}](#attr-fe-autocomplete-given-name)\",
but the user agent\'s only idea for what to prefill the field with is
\"Evan\", then the user agent cannot prefill the field. It would not be
conforming to somehow set the
[`select`{#autofill-processing-model:the-select-element-7}](form-elements.html#the-select-element)
element to the value \"Evan\", since the user could not have done so
themselves.

A user agent prefilling a form control must not discriminate between
form controls that are [in a document
tree](https://dom.spec.whatwg.org/#in-a-document-tree){#autofill-processing-model:in-a-document-tree
x-internal="in-a-document-tree"} and those that are
[connected](https://dom.spec.whatwg.org/#connected){#autofill-processing-model:connected
x-internal="connected"}; that is, it is not conforming to make the
decision on whether or not to autofill based on whether the element\'s
[root](https://dom.spec.whatwg.org/#concept-tree-root){#autofill-processing-model:root
x-internal="root"} is a [shadow
root](https://dom.spec.whatwg.org/#concept-shadow-root){#autofill-processing-model:shadow-root
x-internal="shadow-root"} versus a
[`Document`{#autofill-processing-model:document}](dom.html#document).

A user agent prefilling a form control\'s
[value](#concept-fe-value){#autofill-processing-model:concept-fe-value-2}
must not cause that control to [suffer from a type
mismatch](#suffering-from-a-type-mismatch){#autofill-processing-model:suffering-from-a-type-mismatch},
[suffer from being too
long](#suffering-from-being-too-long){#autofill-processing-model:suffering-from-being-too-long},
[suffer from being too
short](#suffering-from-being-too-short){#autofill-processing-model:suffering-from-being-too-short},
[suffer from an
underflow](#suffering-from-an-underflow){#autofill-processing-model:suffering-from-an-underflow},
[suffer from an
overflow](#suffering-from-an-overflow){#autofill-processing-model:suffering-from-an-overflow},
or [suffer from a step
mismatch](#suffering-from-a-step-mismatch){#autofill-processing-model:suffering-from-a-step-mismatch}.
A user agent prefilling a form control\'s
[value](#concept-fe-value){#autofill-processing-model:concept-fe-value-3}
must not cause that control to [suffer from a pattern
mismatch](#suffering-from-a-pattern-mismatch){#autofill-processing-model:suffering-from-a-pattern-mismatch}
either. Where possible given the control\'s constraints, user agents
must use the format given as canonical in the aforementioned table.
Where it\'s not possible for the canonical format to be used, user
agents should use heuristics to attempt to convert values so that they
can be used.

::: example
For example, if the user agent knows that the user\'s middle name is
\"Ines\", and attempts to prefill a form control that looks like this:

``` html
<input name=middle-initial maxlength=1 autocomplete="additional-name">
```

\...then the user agent could convert \"Ines\" to \"I\" and prefill it
that way.
:::

::: example
A more elaborate example would be with month values. If the user agent
knows that the user\'s birthday is the 27th of July 2012, then it might
try to prefill all of the following controls with slightly different
values, all driven from this information:

``` html
<input name=b type=month autocomplete="bday">
```

2012-07

The day is dropped since the
[Month](input.html#month-state-(type=month)){#autofill-processing-model:month-state-(type=month)}
state only accepts a month/year combination. (Note that this example is
non-conforming, because the [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-18}
[`bday`{#autofill-processing-model:attr-fe-autocomplete-bday-2}](#attr-fe-autocomplete-bday)
is not allowed with the
[Month](input.html#month-state-(type=month)){#autofill-processing-model:month-state-(type=month)-2}
state.)

``` html
<select name=c autocomplete="bday">
 <option>Jan
 <option>Feb
 ...
 <option>Jul
 <option>Aug
 ...
</select>
```

July

The user agent picks the month from the listed options, either by
noticing there are twelve options and picking the 7th, or by recognizing
that one of the strings (three characters \"Jul\" followed by a newline
and a space) is a close match for the name of the month (July) in one of
the user agent\'s supported languages, or through some other similar
mechanism.

``` html
<input name=a type=number min=1 max=12 autocomplete="bday-month">
```

7

User agent converts \"July\" to a month number in the range 1..12, like
the field.

``` html
<input name=a type=number min=0 max=11 autocomplete="bday-month">
```

6

User agent converts \"July\" to a month number in the range 0..11, like
the field.

``` html
<input name=a type=number min=1 max=11 autocomplete="bday-month">
```

User agent doesn\'t fill in the field, since it can\'t make a good guess
as to what the form expects.
:::

A user agent may allow the user to override an element\'s [autofill
field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-19},
e.g. to change it from
\"[`off`{#autofill-processing-model:attr-fe-autocomplete-off-6}](#attr-fe-autocomplete-off)\"
to
\"[`on`{#autofill-processing-model:attr-fe-autocomplete-on-4}](#attr-fe-autocomplete-on)\"
to allow values to be remembered and prefilled despite the page
author\'s objections, or to always
\"[`off`{#autofill-processing-model:attr-fe-autocomplete-off-7}](#attr-fe-autocomplete-off)\",
never remembering values.

More specifically, user agents may in particular consider replacing the
[autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-20}
of form controls that match the description given in the first column of
the following table, when their [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-21}
is either
\"[`on`{#autofill-processing-model:attr-fe-autocomplete-on-5}](#attr-fe-autocomplete-on)\"
or
\"[`off`{#autofill-processing-model:attr-fe-autocomplete-off-8}](#attr-fe-autocomplete-off)\",
with the value given in the second cell of that row. If this table is
used, the replacements must be done in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#autofill-processing-model:tree-order
x-internal="tree-order"}, since all but the first row references the
[autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-22}
of earlier elements. When the descriptions below refer to form controls
being preceded or followed by others, they mean in the list of [listed
elements](forms.html#category-listed){#autofill-processing-model:category-listed}
that share the same [form
owner](#form-owner){#autofill-processing-model:form-owner-5}.

Form control

New [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-23}

an
[`input`{#autofill-processing-model:the-input-element-5}](input.html#the-input-element)
element whose
[`type`{#autofill-processing-model:attr-input-type-2}](input.html#attr-input-type)
attribute is in the
[Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#autofill-processing-model:text-(type=text)-state-and-search-state-(type=search)}
state that is followed by an
[`input`{#autofill-processing-model:the-input-element-6}](input.html#the-input-element)
element whose
[`type`{#autofill-processing-model:attr-input-type-3}](input.html#attr-input-type)
attribute is in the
[Password](input.html#password-state-(type=password)){#autofill-processing-model:password-state-(type=password)}
state

\"[`username`{#autofill-processing-model:attr-fe-autocomplete-username-2}](#attr-fe-autocomplete-username)\"

an
[`input`{#autofill-processing-model:the-input-element-7}](input.html#the-input-element)
element whose
[`type`{#autofill-processing-model:attr-input-type-4}](input.html#attr-input-type)
attribute is in the
[Password](input.html#password-state-(type=password)){#autofill-processing-model:password-state-(type=password)-2}
state that is preceded by an
[`input`{#autofill-processing-model:the-input-element-8}](input.html#the-input-element)
element whose [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-24}
is
\"[`username`{#autofill-processing-model:attr-fe-autocomplete-username-3}](#attr-fe-autocomplete-username)\"

\"[`current-password`{#autofill-processing-model:attr-fe-autocomplete-current-password-2}](#attr-fe-autocomplete-current-password)\"

an
[`input`{#autofill-processing-model:the-input-element-9}](input.html#the-input-element)
element whose
[`type`{#autofill-processing-model:attr-input-type-5}](input.html#attr-input-type)
attribute is in the
[Password](input.html#password-state-(type=password)){#autofill-processing-model:password-state-(type=password)-3}
state that is preceded by an
[`input`{#autofill-processing-model:the-input-element-10}](input.html#the-input-element)
element whose [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-25}
is
\"[`current-password`{#autofill-processing-model:attr-fe-autocomplete-current-password-3}](#attr-fe-autocomplete-current-password)\"

\"[`new-password`{#autofill-processing-model:attr-fe-autocomplete-new-password-2}](#attr-fe-autocomplete-new-password)\"

an
[`input`{#autofill-processing-model:the-input-element-11}](input.html#the-input-element)
element whose
[`type`{#autofill-processing-model:attr-input-type-6}](input.html#attr-input-type)
attribute is in the
[Password](input.html#password-state-(type=password)){#autofill-processing-model:password-state-(type=password)-4}
state that is preceded by an
[`input`{#autofill-processing-model:the-input-element-12}](input.html#the-input-element)
element whose [autofill field
name](#autofill-field-name){#autofill-processing-model:autofill-field-name-26}
is
\"[`new-password`{#autofill-processing-model:attr-fe-autocomplete-new-password-3}](#attr-fe-autocomplete-new-password)\"

\"[`new-password`{#autofill-processing-model:attr-fe-autocomplete-new-password-4}](#attr-fe-autocomplete-new-password)\"

The [`autocomplete`]{#dom-fe-autocomplete .dfn
dfn-for="HTMLInputElement,HTMLSelectElement,HTMLTextAreaElement"
dfn-type="attribute"} IDL attribute, on getting, must return the
element\'s [IDL-exposed autofill
value](#idl-exposed-autofill-value){#autofill-processing-model:idl-exposed-autofill-value-5},
and on setting, must
[reflect](common-dom-interfaces.html#reflect){#autofill-processing-model:reflect}
the content attribute of the same name.

#### [4.10.19]{.secno} APIs for the text control selections[](#textFieldSelection){.self-link} {#textFieldSelection}

The
[`input`{#textFieldSelection:the-input-element}](input.html#the-input-element)
and
[`textarea`{#textFieldSelection:the-textarea-element}](form-elements.html#the-textarea-element)
elements define several attributes and methods for handling their
selection. Their shared algorithms are defined here.

`element`{.variable}`.`[`select`](#dom-textarea/input-select){#dom-textarea/input-select-dev}`()`

Selects everything in the text control.

`element`{.variable}`.`[`selectionStart`](#dom-textarea/input-selectionstart){#dom-textarea/input-selectionstart-dev}` [ = ``value`{.variable}` ]`

Returns the offset to the start of the selection.

Can be set, to change the start of the selection.

`element`{.variable}`.`[`selectionEnd`](#dom-textarea/input-selectionend){#dom-textarea/input-selectionend-dev}` [ = ``value`{.variable}` ]`

Returns the offset to the end of the selection.

Can be set, to change the end of the selection.

`element`{.variable}`.`[`selectionDirection`](#dom-textarea/input-selectiondirection){#dom-textarea/input-selectiondirection-dev}` [ = ``value`{.variable}` ]`

Returns the current direction of the selection.

Can be set, to change the direction of the selection.

The possible values are \"`forward`\", \"`backward`\", and \"`none`\".

`element`{.variable}`.`[`setSelectionRange`](#dom-textarea/input-setselectionrange){#dom-textarea/input-setselectionrange-dev}`(``start`{.variable}`, ``end`{.variable}` [, ``direction`{.variable}`])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/setSelectionRange](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/setSelectionRange "The HTMLInputElement.setSelectionRange() method sets the start and end positions of the current text selection in an <input> or <textarea> element.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Changes the selection to cover the given substring in the given
direction. If the direction is omitted, it will be reset to be the
platform default (none or forward).

`element`{.variable}`.`[`setRangeText`](#dom-textarea/input-setrangetext){#dom-textarea/input-setrangetext-dev}`(``replacement`{.variable}` [, ``start`{.variable}`, ``end`{.variable}` [, ``selectionMode`{.variable}` ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/setRangeText](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/setRangeText "The HTMLInputElement.setRangeText() method replaces a range of text in an <input> or <textarea> element with a new string.")

Support in all current engines.

::: support
[Firefox27+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome24+]{.chrome
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

Replaces a range of text with the new text. If the `start`{.variable}
and `end`{.variable} arguments are not provided, the range is assumed to
be the selection.

The final argument determines how the selection will be set after the
text has been replaced. The possible values are:

\"[`select`{#dom-selectionmode-select-dev}](#dom-selectionmode-select)\"
:   Selects the newly inserted text.

\"[`start`{#dom-selectionmode-start-dev}](#dom-selectionmode-start)\"
:   Moves the selection to just before the inserted text.

\"[`end`{#dom-selectionmode-end-dev}](#dom-selectionmode-end)\"
:   Moves the selection to just after the selected text.

\"[`preserve`{#dom-selectionmode-preserve-dev}](#dom-selectionmode-preserve)\"
:   Attempts to preserve the selection. This is the default.

All
[`input`{#textFieldSelection:the-input-element-2}](input.html#the-input-element)
elements to which these APIs
[apply](input.html#concept-input-apply){#textFieldSelection:concept-input-apply},
and all
[`textarea`{#textFieldSelection:the-textarea-element-2}](form-elements.html#the-textarea-element)
elements, have either a [selection]{#concept-textarea/input-selection
.dfn} or a [text entry cursor position]{#concept-textarea/input-cursor
.dfn} at all times (even for elements that are not [being
rendered](rendering.html#being-rendered){#textFieldSelection:being-rendered}),
measured in offsets into the [code
units](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit
x-internal="code-unit"} of the control\'s [relevant
value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value}.
The initial state must consist of a [text entry
cursor](#concept-textarea/input-cursor){#textFieldSelection:concept-textarea/input-cursor}
at the beginning of the control.

For
[`input`{#textFieldSelection:the-input-element-3}](input.html#the-input-element)
elements, these APIs must operate on the element\'s
[value](#concept-fe-value){#textFieldSelection:concept-fe-value}. For
[`textarea`{#textFieldSelection:the-textarea-element-3}](form-elements.html#the-textarea-element)
elements, these APIs must operate on the element\'s [API
value](#concept-fe-api-value){#textFieldSelection:concept-fe-api-value}.
In the below algorithms, we call the value string being operated on the
[relevant value]{#concept-textarea/input-relevant-value .dfn}.

::: example
The use of [API
value](#concept-fe-api-value){#textFieldSelection:concept-fe-api-value-2}
instead of [raw
value](form-elements.html#concept-textarea-raw-value){#textFieldSelection:concept-textarea-raw-value}
for
[`textarea`{#textFieldSelection:the-textarea-element-4}](form-elements.html#the-textarea-element)
elements means that U+000D (CR) characters are normalized away. For
example,

``` html
<textarea id="demo"></textarea>
<script>
 demo.value = "A\r\nB";
 demo.setRangeText("replaced", 0, 2);
 assert(demo.value === "replacedB");
</script>
```

If we had operated on the [raw
value](form-elements.html#concept-textarea-raw-value){#textFieldSelection:concept-textarea-raw-value-2}
of \"`A\r\nB`\", then we would have replaced the characters \"`A\r`\",
ending up with a result of \"`replaced\nB`\". But since we used the [API
value](#concept-fe-api-value){#textFieldSelection:concept-fe-api-value-3}
of \"`A\nB`\", we replaced the characters \"`A\n`\", giving
\"`replacedB`\".
:::

Characters with no visible rendering, such as U+200D ZERO WIDTH JOINER,
still count as characters. Thus, for instance, the selection can include
just an invisible character, and the text insertion cursor can be placed
to one side or another of such a character.

Whenever the [relevant
value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-2}
changes for an element to which these APIs apply, run these steps:

1.  If the element has a
    [selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection}:

    1.  If the start of the selection is now past the end of the
        [relevant
        value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-3},
        set it to the end of the [relevant
        value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-4}.

    2.  If the end of the selection is now past the end of the [relevant
        value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-5},
        set it to the end of the [relevant
        value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-6}.

    3.  If the user agent does not support empty selection, and both the
        start and end of the selection are now pointing to the end of
        the [relevant
        value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-7},
        then instead set the element\'s [text entry cursor
        position](#concept-textarea/input-cursor){#textFieldSelection:concept-textarea/input-cursor-2}
        to the end of the [relevant
        value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-8},
        removing any selection.

2.  Otherwise, the element must have a [text entry cursor
    position](#concept-textarea/input-cursor){#textFieldSelection:concept-textarea/input-cursor-3}
    position. If it is now past the end of the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-9},
    set it to the end of the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-10}.

In some cases where the [relevant
value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-11}
changes, other parts of the specification will also modify the [text
entry cursor
position](#concept-textarea/input-cursor){#textFieldSelection:concept-textarea/input-cursor-4},
beyond just the clamping steps above. For example, see the
[`value`{#textFieldSelection:dom-textarea-value}](form-elements.html#dom-textarea-value)
setter for
[`textarea`{#textFieldSelection:the-textarea-element-5}](form-elements.html#the-textarea-element).

Where possible, user interface features for changing the [text
selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-2}
in
[`input`{#textFieldSelection:the-input-element-4}](input.html#the-input-element)
and
[`textarea`{#textFieldSelection:the-textarea-element-6}](form-elements.html#the-textarea-element)
elements must be implemented using the [set the selection
range](#set-the-selection-range){#textFieldSelection:set-the-selection-range}
algorithm so that, e.g., all the same events fire.

The
[selections](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-3}
of
[`input`{#textFieldSelection:the-input-element-5}](input.html#the-input-element)
and
[`textarea`{#textFieldSelection:the-textarea-element-7}](form-elements.html#the-textarea-element)
elements have a [selection direction]{#selection-direction .dfn}, which
is either \"`forward`\", \"`backward`\", or \"`none`\". The exact
meaning of the selection direction depends on the platform. This
direction is set when the user manipulates the selection. The initial
[selection
direction](#selection-direction){#textFieldSelection:selection-direction}
must be \"`none`\" if the platform supports that direction, or
\"`forward`\" otherwise.

To [set the selection direction]{#set-the-selection-direction .dfn} of
an element to a given direction, update the element\'s [selection
direction](#selection-direction){#textFieldSelection:selection-direction-2}
to the given direction, unless the direction is \"`none`\" and the
platform does not support that direction; in that case, update the
element\'s [selection
direction](#selection-direction){#textFieldSelection:selection-direction-3}
to \"`forward`\".

::: note
On Windows, the direction indicates the position of the caret relative
to the selection: a \"`forward`\" selection has the caret at the end of
the selection and a \"`backward`\" selection has the caret at the start
of the selection. Windows has no \"`none`\" direction.

On Mac, the direction indicates which end of the selection is affected
when the user adjusts the size of the selection using the arrow keys
with the Shift modifier: the \"`forward`\" direction means the end of
the selection is modified, and the \"`backward`\" direction means the
start of the selection is modified. The \"`none`\" direction is the
default on Mac, it indicates that no particular direction has yet been
selected. The user sets the direction implicitly when first adjusting
the selection, based on which directional arrow key was used.
:::

:::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/select](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/select "The HTMLInputElement.select() method selects all the text in a <textarea> element or in an <input> element that includes a text field.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

::: feature
[HTMLInputElement/select](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/select "The HTMLInputElement.select() method selects all the text in a <textarea> element or in an <input> element that includes a text field.")
:::
::::::

The [`select()`]{#dom-textarea/input-select .dfn
dfn-for="HTMLInputElement,HTMLTextAreaElement" dfn-type="method"}
method, when invoked, must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-6}](input.html#the-input-element)
    element, and either
    [`select()`{#textFieldSelection:dom-textarea/input-select}](#dom-textarea/input-select)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply} to
    this element or the corresponding control has no selectable text,
    return.

    For instance, in a user agent where
    [`<input type=color>`{#textFieldSelection:color-state-(type=color)}](input.html#color-state-(type=color))
    is rendered as a color well with a picker, as opposed to a text
    control accepting a hexadecimal color code, there would be no
    selectable text, and thus calls to the method are ignored.

2.  [Set the selection
    range](#set-the-selection-range){#textFieldSelection:set-the-selection-range-2}
    with 0 and infinity.

The [`selectionStart`]{#dom-textarea/input-selectionstart .dfn
dfn-for="HTMLInputElement,HTMLTextAreaElement" dfn-type="attribute"}
attribute\'s getter must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-7}](input.html#the-input-element)
    element, and
    [`selectionStart`{#textFieldSelection:dom-textarea/input-selectionstart}](#dom-textarea/input-selectionstart)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-2}
    to this element, return null.

2.  If there is no
    [selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-4},
    return the [code
    unit](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit-2
    x-internal="code-unit"} offset within the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-12}
    to the character that immediately follows the [text entry
    cursor](#concept-textarea/input-cursor){#textFieldSelection:concept-textarea/input-cursor-5}.

3.  Return the [code
    unit](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit-3
    x-internal="code-unit"} offset within the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-13}
    to the character that immediately follows the start of the
    [selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-5}.

The
[`selectionStart`{#textFieldSelection:dom-textarea/input-selectionstart-2}](#dom-textarea/input-selectionstart)
attribute\'s setter must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-8}](input.html#the-input-element)
    element, and
    [`selectionStart`{#textFieldSelection:dom-textarea/input-selectionstart-3}](#dom-textarea/input-selectionstart)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-3}
    to this element, throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#textFieldSelection:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#textFieldSelection:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `end`{.variable} be the value of this element\'s
    [`selectionEnd`{#textFieldSelection:dom-textarea/input-selectionend}](#dom-textarea/input-selectionend)
    attribute.

3.  If `end`{.variable} is less than the given value, set
    `end`{.variable} to the given value.

4.  [Set the selection
    range](#set-the-selection-range){#textFieldSelection:set-the-selection-range-3}
    with the given value, `end`{.variable}, and the value of this
    element\'s
    [`selectionDirection`{#textFieldSelection:dom-textarea/input-selectiondirection}](#dom-textarea/input-selectiondirection)
    attribute.

The [`selectionEnd`]{#dom-textarea/input-selectionend .dfn
dfn-for="HTMLInputElement,HTMLTextAreaElement" dfn-type="attribute"}
attribute\'s getter must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-9}](input.html#the-input-element)
    element, and
    [`selectionEnd`{#textFieldSelection:dom-textarea/input-selectionend-2}](#dom-textarea/input-selectionend)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-4}
    to this element, return null.

2.  If there is no
    [selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-6},
    return the [code
    unit](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit-4
    x-internal="code-unit"} offset within the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-14}
    to the character that immediately follows the [text entry
    cursor](#concept-textarea/input-cursor){#textFieldSelection:concept-textarea/input-cursor-6}.

3.  Return the [code
    unit](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit-5
    x-internal="code-unit"} offset within the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-15}
    to the character that immediately follows the end of the
    [selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-7}.

The
[`selectionEnd`{#textFieldSelection:dom-textarea/input-selectionend-3}](#dom-textarea/input-selectionend)
attribute\'s setter must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-10}](input.html#the-input-element)
    element, and
    [`selectionEnd`{#textFieldSelection:dom-textarea/input-selectionend-4}](#dom-textarea/input-selectionend)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-5}
    to this element, throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#textFieldSelection:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#textFieldSelection:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  [Set the selection
    range](#set-the-selection-range){#textFieldSelection:set-the-selection-range-4}
    with the value of this element\'s
    [`selectionStart`{#textFieldSelection:dom-textarea/input-selectionstart-4}](#dom-textarea/input-selectionstart)
    attribute, the given value, and the value of this element\'s
    [`selectionDirection`{#textFieldSelection:dom-textarea/input-selectiondirection-2}](#dom-textarea/input-selectiondirection)
    attribute.

The [`selectionDirection`]{#dom-textarea/input-selectiondirection .dfn
dfn-for="HTMLInputElement,HTMLTextAreaElement" dfn-type="attribute"}
attribute\'s getter must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-11}](input.html#the-input-element)
    element, and
    [`selectionDirection`{#textFieldSelection:dom-textarea/input-selectiondirection-3}](#dom-textarea/input-selectiondirection)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-6}
    to this element, return null.

2.  Return this element\'s [selection
    direction](#selection-direction){#textFieldSelection:selection-direction-4}.

The
[`selectionDirection`{#textFieldSelection:dom-textarea/input-selectiondirection-4}](#dom-textarea/input-selectiondirection)
attribute\'s setter must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-12}](input.html#the-input-element)
    element, and
    [`selectionDirection`{#textFieldSelection:dom-textarea/input-selectiondirection-5}](#dom-textarea/input-selectiondirection)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-7}
    to this element, throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#textFieldSelection:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#textFieldSelection:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  [Set the selection
    range](#set-the-selection-range){#textFieldSelection:set-the-selection-range-5}
    with the value of this element\'s
    [`selectionStart`{#textFieldSelection:dom-textarea/input-selectionstart-5}](#dom-textarea/input-selectionstart)
    attribute, the value of this element\'s
    [`selectionEnd`{#textFieldSelection:dom-textarea/input-selectionend-5}](#dom-textarea/input-selectionend)
    attribute, and the given value.

The
[`setSelectionRange(``start`{.variable}`, ``end`{.variable}`, ``direction`{.variable}`)`]{#dom-textarea/input-setselectionrange
.dfn dfn-for="HTMLInputElement,HTMLTextAreaElement" dfn-type="method"}
method, when invoked, must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-13}](input.html#the-input-element)
    element, and
    [`setSelectionRange()`{#textFieldSelection:dom-textarea/input-setselectionrange}](#dom-textarea/input-setselectionrange)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-8}
    to this element, throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#textFieldSelection:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#textFieldSelection:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  [Set the selection
    range](#set-the-selection-range){#textFieldSelection:set-the-selection-range-6}
    with `start`{.variable}, `end`{.variable}, and
    `direction`{.variable}.

To [set the selection range]{#set-the-selection-range .dfn} with an
integer or null `start`{.variable}, an integer or null or the special
value infinity `end`{.variable}, and optionally a string
`direction`{.variable}, run the following steps:

1.  If `start`{.variable} is null, let `start`{.variable} be zero.

2.  If `end`{.variable} is null, let `end`{.variable} be zero.

3.  Set the
    [selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-8}
    of the text control to the sequence of [code
    units](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit-6
    x-internal="code-unit"} within the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-16}
    starting with the code unit at the `start`{.variable}th position (in
    logical order) and ending with the code unit at the
    (`end`{.variable}-1)th position. Arguments greater than the
    [length](https://infra.spec.whatwg.org/#string-length){#textFieldSelection:length
    x-internal="length"} of the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-17}
    of the text control (including the special value infinity) must be
    treated as pointing at the end of the text control. If
    `end`{.variable} is less than or equal to `start`{.variable}, then
    the start of the selection and the end of the selection must both be
    placed immediately before the character with offset
    `end`{.variable}. In UAs where there is no concept of an empty
    selection, this must set the cursor to be just before the character
    with offset `end`{.variable}.

4.  If `direction`{.variable} is not [identical
    to](https://infra.spec.whatwg.org/#string-is){#textFieldSelection:identical-to
    x-internal="identical-to"} either \"`backward`\" or \"`forward`\",
    or if the `direction`{.variable} argument was not given, set
    `direction`{.variable} to \"`none`\".

5.  [Set the selection
    direction](#set-the-selection-direction){#textFieldSelection:set-the-selection-direction}
    of the text control to `direction`{.variable}.

6.  If the previous steps caused the
    [selection](#concept-textarea/input-selection){#textFieldSelection:concept-textarea/input-selection-9}
    of the text control to be modified (in either extent or
    [direction](#selection-direction){#textFieldSelection:selection-direction-5}),
    then [queue an element
    task](webappapis.html#queue-an-element-task){#textFieldSelection:queue-an-element-task}
    on the [user interaction task
    source](webappapis.html#user-interaction-task-source){#textFieldSelection:user-interaction-task-source}
    given the element to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#textFieldSelection:concept-event-fire
    x-internal="concept-event-fire"} named
    [`select`{#textFieldSelection:event-select}](indices.html#event-select)
    at the element, with the
    [`bubbles`{#textFieldSelection:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    attribute initialized to true.

The
[`setRangeText(``replacement`{.variable}`, ``start`{.variable}`, ``end`{.variable}`, ``selectMode`{.variable}`)`]{#dom-textarea/input-setrangetext
.dfn dfn-for="HTMLInputElement,HTMLTextAreaElement" dfn-type="method"}
method, when invoked, must run the following steps:

1.  If this element is an
    [`input`{#textFieldSelection:the-input-element-14}](input.html#the-input-element)
    element, and
    [`setRangeText()`{#textFieldSelection:dom-textarea/input-setrangetext}](#dom-textarea/input-setrangetext)
    [does not
    apply](input.html#do-not-apply){#textFieldSelection:do-not-apply-9}
    to this element, throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#textFieldSelection:invalidstateerror-5
    x-internal="invalidstateerror"}
    [`DOMException`{#textFieldSelection:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Set this element\'s [dirty value
    flag](#concept-fe-dirty){#textFieldSelection:concept-fe-dirty} to
    true.

3.  If the method has only one argument, then let `start`{.variable} and
    `end`{.variable} have the values of the
    [`selectionStart`{#textFieldSelection:dom-textarea/input-selectionstart-6}](#dom-textarea/input-selectionstart)
    attribute and the
    [`selectionEnd`{#textFieldSelection:dom-textarea/input-selectionend-6}](#dom-textarea/input-selectionend)
    attribute respectively.

    Otherwise, let `start`{.variable}, `end`{.variable} have the values
    of the second and third arguments respectively.

4.  If `start`{.variable} is greater than `end`{.variable}, then throw
    an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#textFieldSelection:indexsizeerror
    x-internal="indexsizeerror"}
    [`DOMException`{#textFieldSelection:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  If `start`{.variable} is greater than the
    [length](https://infra.spec.whatwg.org/#string-length){#textFieldSelection:length-2
    x-internal="length"} of the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-18}
    of the text control, then set it to the
    [length](https://infra.spec.whatwg.org/#string-length){#textFieldSelection:length-3
    x-internal="length"} of the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-19}
    of the text control.

6.  If `end`{.variable} is greater than the
    [length](https://infra.spec.whatwg.org/#string-length){#textFieldSelection:length-4
    x-internal="length"} of the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-20}
    of the text control, then set it to the
    [length](https://infra.spec.whatwg.org/#string-length){#textFieldSelection:length-5
    x-internal="length"} of the [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-21}
    of the text control.

7.  Let `selection start`{.variable} be the current value of the
    [`selectionStart`{#textFieldSelection:dom-textarea/input-selectionstart-7}](#dom-textarea/input-selectionstart)
    attribute.

8.  Let `selection end`{.variable} be the current value of the
    [`selectionEnd`{#textFieldSelection:dom-textarea/input-selectionend-7}](#dom-textarea/input-selectionend)
    attribute.

9.  If `start`{.variable} is less than `end`{.variable}, delete the
    sequence of [code
    units](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit-7
    x-internal="code-unit"} within the element\'s [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-22}
    starting with the code unit at the `start`{.variable}th position and
    ending with the code unit at the (`end`{.variable}-1)th position.

10. Insert the value of the first argument into the text of the
    [relevant
    value](#concept-textarea/input-relevant-value){#textFieldSelection:concept-textarea/input-relevant-value-23}
    of the text control, immediately before the `start`{.variable}th
    [code
    unit](https://infra.spec.whatwg.org/#code-unit){#textFieldSelection:code-unit-8
    x-internal="code-unit"}.

11. Let `new length`{.variable} be the
    [length](https://infra.spec.whatwg.org/#string-length){#textFieldSelection:length-6
    x-internal="length"} of the value of the first argument.

12. Let `new end`{.variable} be the sum of `start`{.variable} and
    `new length`{.variable}.

13. Run the appropriate set of substeps from the following list:

    If the fourth argument\'s value is \"[`select`]{#dom-selectionmode-select .dfn dfn-for="SelectionMode" dfn-type="enum-value"}\"

    :   Let `selection start`{.variable} be `start`{.variable}.

        Let `selection end`{.variable} be `new end`{.variable}.

    If the fourth argument\'s value is \"[`start`]{#dom-selectionmode-start .dfn dfn-for="SelectionMode" dfn-type="enum-value"}\"

    :   Let `selection start`{.variable} and `selection end`{.variable}
        be `start`{.variable}.

    If the fourth argument\'s value is \"[`end`]{#dom-selectionmode-end .dfn dfn-for="SelectionMode" dfn-type="enum-value"}\"

    :   Let `selection start`{.variable} and `selection end`{.variable}
        be `new end`{.variable}.

    If the fourth argument\'s value is \"[`preserve`]{#dom-selectionmode-preserve .dfn dfn-for="SelectionMode" dfn-type="enum-value"}\"\
    If the method has only one argument

    :   1.  Let `old length`{.variable} be `end`{.variable} minus
            `start`{.variable}.

        2.  Let `delta`{.variable} be `new length`{.variable} minus
            `old length`{.variable}.

        3.  If `selection start`{.variable} is greater than
            `end`{.variable}, then increment it by `delta`{.variable}.
            (If `delta`{.variable} is negative, i.e. the new text is
            shorter than the old text, then this will *decrease* the
            value of `selection start`{.variable}.)

            Otherwise: if `selection start`{.variable} is greater than
            `start`{.variable}, then set it to `start`{.variable}. (This
            snaps the start of the selection to the start of the new
            text if it was in the middle of the text that it replaced.)

        4.  If `selection end`{.variable} is greater than
            `end`{.variable}, then increment it by `delta`{.variable} in
            the same way.

            Otherwise: if `selection end`{.variable} is greater than
            `start`{.variable}, then set it to `new end`{.variable}.
            (This snaps the end of the selection to the end of the new
            text if it was in the middle of the text that it replaced.)

14. [Set the selection
    range](#set-the-selection-range){#textFieldSelection:set-the-selection-range-7}
    with `selection start`{.variable} and `selection end`{.variable}.

The
[`setRangeText()`{#textFieldSelection:dom-textarea/input-setrangetext-2}](#dom-textarea/input-setrangetext)
method uses the following enumeration:

``` idl
enum SelectionMode {
  "select",
  "start",
  "end",
  "preserve" // default
};
```

------------------------------------------------------------------------

::: example
To obtain the currently selected text, the following JavaScript
suffices:

``` js
var selectionText = control.value.substring(control.selectionStart, control.selectionEnd);
```

\...where `control`{.variable} is the
[`input`{#textFieldSelection:the-input-element-15}](input.html#the-input-element)
or
[`textarea`{#textFieldSelection:the-textarea-element-8}](form-elements.html#the-textarea-element)
element.
:::

::: example
To add some text at the start of a text control, while maintaining the
text selection, the three attributes must be preserved:

``` js
var oldStart = control.selectionStart;
var oldEnd = control.selectionEnd;
var oldDirection = control.selectionDirection;
var prefix = "http://";
control.value = prefix + control.value;
control.setSelectionRange(oldStart + prefix.length, oldEnd + prefix.length, oldDirection);
```

\...where `control`{.variable} is the
[`input`{#textFieldSelection:the-input-element-16}](input.html#the-input-element)
or
[`textarea`{#textFieldSelection:the-textarea-element-9}](form-elements.html#the-textarea-element)
element.
:::

#### [4.10.20]{.secno} Constraints[](#constraints){.self-link}

##### [4.10.20.1]{.secno} Definitions[](#definitions){.self-link}

A [submittable
element](forms.html#category-submit){#definitions:category-submit} is a
[candidate for constraint
validation]{#candidate-for-constraint-validation .dfn} except when a
condition has [barred the element from constraint
validation]{#barred-from-constraint-validation .dfn}. (For example, an
element is [barred from constraint
validation](#barred-from-constraint-validation){#definitions:barred-from-constraint-validation}
if it has a
[`datalist`{#definitions:the-datalist-element}](form-elements.html#the-datalist-element)
element ancestor.)

An element can have a [custom validity error
message]{#custom-validity-error-message .dfn} defined. Initially, an
element must have its [custom validity error
message](#custom-validity-error-message){#definitions:custom-validity-error-message}
set to the empty string. When its value is not the empty string, the
element is [suffering from a custom
error](#suffering-from-a-custom-error){#definitions:suffering-from-a-custom-error}.
It can be set using the
[`setCustomValidity()`{#definitions:dom-cva-setcustomvalidity}](#dom-cva-setcustomvalidity)
method, except for [form-associated custom
elements](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element}.
[Form-associated custom
elements](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-2}
can have a [custom validity error
message](#custom-validity-error-message){#definitions:custom-validity-error-message-2}
set via their
[`ElementInternals`{#definitions:elementinternals}](custom-elements.html#elementinternals)
object\'s
[`setValidity()`{#definitions:dom-elementinternals-setvalidity}](custom-elements.html#dom-elementinternals-setvalidity)
method. The user agent should use the [custom validity error
message](#custom-validity-error-message){#definitions:custom-validity-error-message-3}
when alerting the user to the problem with the control.

An element can be constrained in various ways. The following is the list
of [validity states]{#validity-states .dfn} that a form control can be
in, making the control invalid for the purposes of constraint
validation. (The definitions below are non-normative; other parts of
this specification define more precisely when each state applies or does
not.)

[Suffering from being missing]{#suffering-from-being-missing .dfn}

:   When a control has no
    [value](#concept-fe-value){#definitions:concept-fe-value} but has a
    `required` attribute
    ([`input`{#definitions:the-input-element}](input.html#the-input-element)
    [`required`{#definitions:attr-input-required}](input.html#attr-input-required),
    [`textarea`{#definitions:the-textarea-element}](form-elements.html#the-textarea-element)
    [`required`{#definitions:attr-textarea-required}](form-elements.html#attr-textarea-required));
    or, more complicated rules for
    [`select`{#definitions:the-select-element}](form-elements.html#the-select-element)
    elements and controls in [radio button
    groups](input.html#radio-button-group){#definitions:radio-button-group},
    as specified in their sections.

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-2}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `valueMissing` flag to true for a [form-associated
    custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-3}.

[Suffering from a type mismatch]{#suffering-from-a-type-mismatch .dfn}

:   When a control that allows arbitrary user input has a
    [value](#concept-fe-value){#definitions:concept-fe-value-2} that is
    not in the correct syntax
    ([Email](input.html#email-state-(type=email)){#definitions:email-state-(type=email)},
    [URL](input.html#url-state-(type=url)){#definitions:url-state-(type=url)}).

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-3}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `typeMismatch` flag to true for a [form-associated
    custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-4}.

[Suffering from a pattern mismatch]{#suffering-from-a-pattern-mismatch .dfn}

:   When a control has a
    [value](#concept-fe-value){#definitions:concept-fe-value-3} that
    doesn\'t satisfy the
    [`pattern`{#definitions:attr-input-pattern}](input.html#attr-input-pattern)
    attribute.

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-4}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `patternMismatch` flag to true for a [form-associated
    custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-5}.

[Suffering from being too long]{#suffering-from-being-too-long .dfn}

:   When a control has a
    [value](#concept-fe-value){#definitions:concept-fe-value-4} that is
    too long for the [form control `maxlength`
    attribute](#attr-fe-maxlength){#definitions:attr-fe-maxlength}
    ([`input`{#definitions:the-input-element-2}](input.html#the-input-element)
    [`maxlength`{#definitions:attr-input-maxlength}](input.html#attr-input-maxlength),
    [`textarea`{#definitions:the-textarea-element-2}](form-elements.html#the-textarea-element)
    [`maxlength`{#definitions:attr-textarea-maxlength}](form-elements.html#attr-textarea-maxlength)).

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-5}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `tooLong` flag to true for a [form-associated custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-6}.

[Suffering from being too short]{#suffering-from-being-too-short .dfn}

:   When a control has a
    [value](#concept-fe-value){#definitions:concept-fe-value-5} that is
    too short for the [form control `minlength`
    attribute](#attr-fe-minlength){#definitions:attr-fe-minlength}
    ([`input`{#definitions:the-input-element-3}](input.html#the-input-element)
    [`minlength`{#definitions:attr-input-minlength}](input.html#attr-input-minlength),
    [`textarea`{#definitions:the-textarea-element-3}](form-elements.html#the-textarea-element)
    [`minlength`{#definitions:attr-textarea-minlength}](form-elements.html#attr-textarea-minlength)).

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-6}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `tooShort` flag to true for a [form-associated custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-7}.

[Suffering from an underflow]{#suffering-from-an-underflow .dfn}

:   When a control has a
    [value](#concept-fe-value){#definitions:concept-fe-value-6} that is
    not the empty string and is too low for the
    [`min`{#definitions:attr-input-min}](input.html#attr-input-min)
    attribute.

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-7}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `rangeUnderflow` flag to true for a [form-associated
    custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-8}.

[Suffering from an overflow]{#suffering-from-an-overflow .dfn}

:   When a control has a
    [value](#concept-fe-value){#definitions:concept-fe-value-7} that is
    not the empty string and is too high for the
    [`max`{#definitions:attr-input-max}](input.html#attr-input-max)
    attribute.

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-8}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `rangeOverflow` flag to true for a [form-associated
    custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-9}.

[Suffering from a step mismatch]{#suffering-from-a-step-mismatch .dfn}

:   When a control has a
    [value](#concept-fe-value){#definitions:concept-fe-value-8} that
    doesn\'t fit the rules given by the
    [`step`{#definitions:attr-input-step}](input.html#attr-input-step)
    attribute.

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-9}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `stepMismatch` flag to true for a [form-associated
    custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-10}.

[Suffering from bad input]{#suffering-from-bad-input .dfn}

:   When a control has incomplete input and the user agent does not
    think the user ought to be able to submit the form in its current
    state.

    When the
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-10}](custom-elements.html#dom-elementinternals-setvalidity)
    method sets `badInput` flag to true for a [form-associated custom
    element](custom-elements.html#form-associated-custom-element){#definitions:form-associated-custom-element-11}.

[Suffering from a custom error]{#suffering-from-a-custom-error .dfn}

:   When a control\'s [custom validity error
    message](#custom-validity-error-message){#definitions:custom-validity-error-message-4}
    (as set by the element\'s
    [`setCustomValidity()`{#definitions:dom-cva-setcustomvalidity-2}](#dom-cva-setcustomvalidity)
    method or
    [`ElementInternals`{#definitions:elementinternals-2}](custom-elements.html#elementinternals)\'s
    [`setValidity()`{#definitions:dom-elementinternals-setvalidity-11}](custom-elements.html#dom-elementinternals-setvalidity)
    method) is not the empty string.

An element can still suffer from these states even when the element is
[disabled](#concept-fe-disabled){#definitions:concept-fe-disabled}; thus
these states can be represented in the DOM even if validating the form
during submission wouldn\'t indicate a problem to the user.

An element [satisfies its constraints]{#concept-fv-valid .dfn} if it is
not suffering from any of the above [validity
states](#validity-states){#definitions:validity-states}.

##### [4.10.20.2]{.secno} Constraint validation[](#constraint-validation){.self-link}

When the user agent is required to [statically validate the
constraints]{#statically-validate-the-constraints .dfn} of
[`form`{#constraint-validation:the-form-element}](forms.html#the-form-element)
element `form`{.variable}, it must run the following steps, which return
either a *positive* result (all the controls in the form are valid) or a
*negative* result (there are invalid controls) along with a (possibly
empty) list of elements that are invalid and for which no script has
claimed responsibility:

1.  Let `controls`{.variable} be a list of all the [submittable
    elements](forms.html#category-submit){#constraint-validation:category-submit}
    whose [form owner](#form-owner){#constraint-validation:form-owner}
    is `form`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#constraint-validation:tree-order
    x-internal="tree-order"}.

2.  Let `invalid controls`{.variable} be an initially empty list of
    elements.

3.  For each element `field`{.variable} in `controls`{.variable}, in
    [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#constraint-validation:tree-order-2
    x-internal="tree-order"}:

    1.  If `field`{.variable} is not a [candidate for constraint
        validation](#candidate-for-constraint-validation){#constraint-validation:candidate-for-constraint-validation},
        then move on to the next element.

    2.  Otherwise, if `field`{.variable} [satisfies its
        constraints](#concept-fv-valid){#constraint-validation:concept-fv-valid},
        then move on to the next element.

    3.  Otherwise, add `field`{.variable} to
        `invalid controls`{.variable}.

4.  If `invalid controls`{.variable} is empty, then return a *positive*
    result.

5.  Let `unhandled invalid controls`{.variable} be an initially empty
    list of elements.

6.  For each element `field`{.variable} in
    `invalid controls`{.variable}, if any, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#constraint-validation:tree-order-3
    x-internal="tree-order"}:

    1.  Let `notCanceled`{.variable} be the result of [firing an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#constraint-validation:concept-event-fire
        x-internal="concept-event-fire"} named
        [`invalid`{#constraint-validation:event-invalid}](indices.html#event-invalid)
        at `field`{.variable}, with the
        [`cancelable`{#constraint-validation:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
        attribute initialized to true.

    2.  If `notCanceled`{.variable} is true, then add `field`{.variable}
        to `unhandled invalid controls`{.variable}.

7.  Return a *negative* result with the list of elements in the
    `unhandled invalid controls`{.variable} list.

If a user agent is to [interactively validate the
constraints]{#interactively-validate-the-constraints .dfn} of
[`form`{#constraint-validation:the-form-element-2}](forms.html#the-form-element)
element `form`{.variable}, then the user agent must run the following
steps:

1.  [Statically validate the
    constraints](#statically-validate-the-constraints){#constraint-validation:statically-validate-the-constraints}
    of `form`{.variable}, and let
    `unhandled invalid controls`{.variable} be the list of elements
    returned if the result was *negative*.

2.  If the result was *positive*, then return that result.

3.  Report the problems with the constraints of at least one of the
    elements given in `unhandled invalid controls`{.variable} to the
    user.

    - User agents may focus one of those elements in the process, by
      running the [focusing
      steps](interaction.html#focusing-steps){#constraint-validation:focusing-steps}
      for that element, and may change the scrolling position of the
      document, or perform some other action that brings the element to
      the user\'s attention. For elements that are [form-associated
      custom
      elements](custom-elements.html#form-associated-custom-element){#constraint-validation:form-associated-custom-element},
      user agents should use their [validation
      anchor](custom-elements.html#face-validation-anchor){#constraint-validation:face-validation-anchor}
      instead, for the purposes of these actions.

    - User agents may report more than one constraint violation.

    - User agents may coalesce related constraint violation reports if
      appropriate (e.g. if multiple radio buttons in a
      [group](input.html#radio-button-group){#constraint-validation:radio-button-group}
      are marked as required, only one error need be reported).

    - If one of the controls is not [being
      rendered](rendering.html#being-rendered){#constraint-validation:being-rendered}
      (e.g. it has the
      [`hidden`{#constraint-validation:attr-hidden}](interaction.html#attr-hidden)
      attribute set), then user agents may report a script error.

4.  Return a *negative* result.

##### [4.10.20.3]{.secno} The [constraint validation API]{.dfn}[](#the-constraint-validation-api){.self-link}

`element`{.variable}`.`[`willValidate`](#dom-cva-willvalidate){#dom-cva-willvalidate-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/willValidate](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/willValidate "The willValidate read-only property of the HTMLObjectElement interface returns a boolean value that indicates whether the element is a candidate for constraint validation. Always false for HTMLObjectElement objects.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns true if the element will be validated when the form is
submitted; false otherwise.

`element`{.variable}`.`[`setCustomValidity`](#dom-cva-setcustomvalidity){#dom-cva-setcustomvalidity-dev}`(``message`{.variable}`)`

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/setCustomValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/setCustomValidity "The setCustomValidity() method of the HTMLObjectElement interface sets a custom validity message for the element.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLSelectElement/setCustomValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/setCustomValidity "The HTMLSelectElement.setCustomValidity() method sets the custom validity message for the selection element to the specified message. Use the empty string to indicate that the element does not have a custom validity error.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

Sets a custom error, so that the element would fail to validate. The
given message is the message to be shown to the user when reporting the
problem to the user.

If the argument is the empty string, clears the custom error.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity}`.`[`valueMissing`](#dom-validitystate-valuemissing){#dom-validitystate-valuemissing-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/valueMissing](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/valueMissing "The read-only valueMissing property of a ValidityState object indicates if a required control, such as an <input>, <select>, or <textarea>, has an empty value.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns true if the element has no value but is a required field; false
otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-2}`.`[`typeMismatch`](#dom-validitystate-typemismatch){#dom-validitystate-typemismatch-dev}

Returns true if the element\'s value is not in the correct syntax; false
otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-3}`.`[`patternMismatch`](#dom-validitystate-patternmismatch){#dom-validitystate-patternmismatch-dev}

Returns true if the element\'s value doesn\'t match the provided
pattern; false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-4}`.`[`tooLong`](#dom-validitystate-toolong){#dom-validitystate-toolong-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/tooLong](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/tooLong "The read-only tooLong property of a ValidityState object indicates if the value of an <input> or <textarea>, after having been edited by the user, exceeds the maximum code-unit length established by the element's maxlength attribute.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android64+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns true if the element\'s value is longer than the provided maximum
length; false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-5}`.`[`tooShort`](#dom-validitystate-tooshort){#dom-validitystate-tooshort-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/tooShort](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/tooShort "The read-only tooShort property of a ValidityState object indicates if the value of an <input>, <button>, <select>, <output>, <fieldset> or <textarea>, after having been edited by the user, is less than the minimum code-unit length established by the element's minlength attribute.")

Support in all current engines.

::: support
[Firefox51+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome40+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android64+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android67+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns true if the element\'s value, if it is not the empty string, is
shorter than the provided minimum length; false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-6}`.`[`rangeUnderflow`](#dom-validitystate-rangeunderflow){#dom-validitystate-rangeunderflow-dev}

Returns true if the element\'s value is lower than the provided minimum;
false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-7}`.`[`rangeOverflow`](#dom-validitystate-rangeoverflow){#dom-validitystate-rangeoverflow-dev}

Returns true if the element\'s value is higher than the provided
maximum; false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-8}`.`[`stepMismatch`](#dom-validitystate-stepmismatch){#dom-validitystate-stepmismatch-dev}

Returns true if the element\'s value doesn\'t fit the rules given by the
[`step`{#the-constraint-validation-api:attr-input-step}](input.html#attr-input-step)
attribute; false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-9}`.`[`badInput`](#dom-validitystate-badinput){#dom-validitystate-badinput-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/badInput](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/badInput "The read-only badInput property of a ValidityState object indicates if the user has provided input that the browser is unable to convert. For example, if you have a number input element whose content is a string.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome25+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android64+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns true if the user has provided input in the user interface that
the user agent is unable to convert to a value; false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-10}`.`[`customError`](#dom-validitystate-customerror){#dom-validitystate-customerror-dev}

Returns true if the element has a custom error; false otherwise.

`element`{.variable}`.`[`validity`](#dom-cva-validity){#the-constraint-validation-api:dom-cva-validity-11}`.`[`valid`](#dom-validitystate-valid){#dom-validitystate-valid-dev}

Returns true if the element\'s value has no validity problems; false
otherwise.

`valid`{.variable}` = ``element`{.variable}`.`[`checkValidity`](#dom-cva-checkvalidity){#dom-cva-checkvalidity-dev}`()`

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/checkValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/checkValidity "The HTMLInputElement.checkValidity() method returns a boolean value which indicates validity of the value of the element. If the value is invalid, this method also fires the invalid event on the element.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLObjectElement/checkValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/checkValidity "The checkValidity() method of the HTMLObjectElement interface returns a boolean value that always is true, because object objects are never candidates for constraint validation.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLSelectElement/checkValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSelectElement/checkValidity "The HTMLSelectElement.checkValidity() method checks whether the element has any constraints and whether it satisfies them. If the element fails its constraints, the browser fires a cancelable invalid event at the element, and then returns false.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera9+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::::::

Returns true if the element\'s value has no validity problems; false
otherwise. Fires an
[`invalid`{#the-constraint-validation-api:event-invalid}](indices.html#event-invalid)
event at the element in the latter case.

`valid`{.variable}` = ``element`{.variable}`.`[`reportValidity`](#dom-cva-reportvalidity){#dom-cva-reportvalidity-dev}`()`

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLFormElement/reportValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/reportValidity "The HTMLFormElement.reportValidity() method returns true if the element's child controls satisfy their validation constraints. When false is returned, cancelable invalid events are fired for each invalid child and validation problems are reported to the user.")

Support in all current engines.

::: support
[Firefox49+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome40+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HTMLInputElement/reportValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/reportValidity "The reportValidity() method of the HTMLInputElement interface performs the same validity checking steps as the checkValidity() method. If the value is invalid, this method also fires the invalid event on the element, and (if the event isn't canceled) reports the problem to the user.")

Support in all current engines.

::: support
[Firefox49+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome40+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android64+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns true if the element\'s value has no validity problems;
otherwise, returns false, fires an
[`invalid`{#the-constraint-validation-api:event-invalid-2}](indices.html#event-invalid)
event at the element, and (if the event isn\'t canceled) reports the
problem to the user.

`element`{.variable}`.`[`validationMessage`](#dom-cva-validationmessage){#dom-cva-validationmessage-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/validationMessage](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/validationMessage "The validationMessage read-only property of the HTMLObjectElement interface returns a string representing a localized message that describes the validation constraints that the control does not satisfy (if any). This is the empty string if the control is not a candidate for constraint validation (willValidate is false), or it satisfies its constraints.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the error message that would be shown to the user if the element
was to be checked for validity.

The [`willValidate`]{#dom-cva-willvalidate .dfn
dfn-for="HTMLObjectElement,HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="attribute"} attribute\'s getter must return true, if this
element is a [candidate for constraint
validation](#candidate-for-constraint-validation){#the-constraint-validation-api:candidate-for-constraint-validation},
and false otherwise (i.e., false if any conditions are [barring it from
constraint
validation](#barred-from-constraint-validation){#the-constraint-validation-api:barred-from-constraint-validation}).

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/willValidate](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/willValidate "The willValidate read-only property of the ElementInternals interface returns true if the element is a submittable element that is a candidate for constraint validation.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
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

The [`willValidate`]{#dom-elementinternals-willvalidate .dfn
dfn-for="ElementInternals" dfn-type="attribute"} attribute of
[`ElementInternals`{#the-constraint-validation-api:elementinternals}](custom-elements.html#elementinternals)
interface, on getting, must throw a
[\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-constraint-validation-api:notsupportederror
x-internal="notsupportederror"}
[`DOMException`{#the-constraint-validation-api:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the [target
element](custom-elements.html#internals-target){#the-constraint-validation-api:internals-target}
is not a [form-associated custom
element](custom-elements.html#form-associated-custom-element){#the-constraint-validation-api:form-associated-custom-element}.
Otherwise, it must return true if the [target
element](custom-elements.html#internals-target){#the-constraint-validation-api:internals-target-2}
is a [candidate for constraint
validation](#candidate-for-constraint-validation){#the-constraint-validation-api:candidate-for-constraint-validation-2},
and false otherwise.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/setCustomValidity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/setCustomValidity "The HTMLInputElement.setCustomValidity() method sets a custom validity message for the element.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The
[`setCustomValidity(``error`{.variable}`)`]{#dom-cva-setcustomvalidity
.dfn
dfn-for="HTMLObjectElement,HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="method"} method steps are:

1.  Set `error`{.variable} to the result of [normalizing
    newlines](https://infra.spec.whatwg.org/#normalize-newlines){#the-constraint-validation-api:normalize-newlines
    x-internal="normalize-newlines"} given `error`{.variable}.

2.  Set the [custom validity error
    message](#custom-validity-error-message){#the-constraint-validation-api:custom-validity-error-message}
    to `error`{.variable}.

::: example
In the following example, a script checks the value of a form control
each time it is edited, and whenever it is not a valid value, uses the
[`setCustomValidity()`{#the-constraint-validation-api:dom-cva-setcustomvalidity}](#dom-cva-setcustomvalidity)
method to set an appropriate message.

``` html
<label>Feeling: <input name=f type="text" oninput="check(this)"></label>
<script>
 function check(input) {
   if (input.value == "good" ||
       input.value == "fine" ||
       input.value == "tired") {
     input.setCustomValidity('"' + input.value + '" is not a feeling.');
   } else {
     // input is fine -- reset the error message
     input.setCustomValidity('');
   }
 }
</script>
```
:::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLObjectElement/validity](https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/validity "The validity read-only property of the HTMLObjectElement interface returns a ValidityState with the validity states that this element is in.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome10+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`validity`]{#dom-cva-validity .dfn
dfn-for="HTMLObjectElement,HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="attribute"} attribute\'s getter must return a
[`ValidityState`{#the-constraint-validation-api:validitystate}](#validitystate)
object that represents the [validity
states](#validity-states){#the-constraint-validation-api:validity-states}
of this element. This object is
[live](infrastructure.html#live){#the-constraint-validation-api:live}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/validity](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/validity "The validity read-only property of the ElementInternals interface returns a ValidityState object which represents the different validity states the element can be in, with respect to constraint validation.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
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

The [`validity`]{#dom-elementinternals-validity .dfn
dfn-for="ElementInternals" dfn-type="attribute"} attribute of
[`ElementInternals`{#the-constraint-validation-api:elementinternals-2}](custom-elements.html#elementinternals)
interface, on getting, must throw a
[\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-constraint-validation-api:notsupportederror-2
x-internal="notsupportederror"}
[`DOMException`{#the-constraint-validation-api:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the [target
element](custom-elements.html#internals-target){#the-constraint-validation-api:internals-target-3}
is not a [form-associated custom
element](custom-elements.html#form-associated-custom-element){#the-constraint-validation-api:form-associated-custom-element-2}.
Otherwise, it must return a
[`ValidityState`{#the-constraint-validation-api:validitystate-2}](#validitystate)
object that represents the [validity
states](#validity-states){#the-constraint-validation-api:validity-states-2}
of the [target
element](custom-elements.html#internals-target){#the-constraint-validation-api:internals-target-4}.
This object is
[live](infrastructure.html#live){#the-constraint-validation-api:live-2}.

``` idl
[Exposed=Window]
interface ValidityState {
  readonly attribute boolean valueMissing;
  readonly attribute boolean typeMismatch;
  readonly attribute boolean patternMismatch;
  readonly attribute boolean tooLong;
  readonly attribute boolean tooShort;
  readonly attribute boolean rangeUnderflow;
  readonly attribute boolean rangeOverflow;
  readonly attribute boolean stepMismatch;
  readonly attribute boolean badInput;
  readonly attribute boolean customError;
  readonly attribute boolean valid;
};
```

A
[`ValidityState`{#the-constraint-validation-api:validitystate-3}](#validitystate)
object has the following attributes. On getting, they must return true
if the corresponding condition given in the following list is true, and
false otherwise.

[`valueMissing`]{#dom-validitystate-valuemissing .dfn
dfn-for="ValidityState" dfn-type="attribute"}

The control is [suffering from being
missing](#suffering-from-being-missing){#the-constraint-validation-api:suffering-from-being-missing}.

[`typeMismatch`]{#dom-validitystate-typemismatch .dfn
dfn-for="ValidityState" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/typeMismatch](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/typeMismatch "The read-only typeMismatch property of a ValidityState object indicates if the value of an <input>, after having been edited by the user, does not conform to the constraints set by the element's type attribute.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The control is [suffering from a type
mismatch](#suffering-from-a-type-mismatch){#the-constraint-validation-api:suffering-from-a-type-mismatch}.

[`patternMismatch`]{#dom-validitystate-patternmismatch .dfn
dfn-for="ValidityState" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/patternMismatch](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/patternMismatch "The read-only patternMismatch property of a ValidityState object indicates if the value of an <input>, after having been edited by the user, does not conform to the constraints set by the element's pattern attribute.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The control is [suffering from a pattern
mismatch](#suffering-from-a-pattern-mismatch){#the-constraint-validation-api:suffering-from-a-pattern-mismatch}.

[`tooLong`]{#dom-validitystate-toolong .dfn dfn-for="ValidityState"
dfn-type="attribute"}

The control is [suffering from being too
long](#suffering-from-being-too-long){#the-constraint-validation-api:suffering-from-being-too-long}.

[`tooShort`]{#dom-validitystate-tooshort .dfn dfn-for="ValidityState"
dfn-type="attribute"}

The control is [suffering from being too
short](#suffering-from-being-too-short){#the-constraint-validation-api:suffering-from-being-too-short}.

[`rangeUnderflow`]{#dom-validitystate-rangeunderflow .dfn
dfn-for="ValidityState" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/rangeUnderflow](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/rangeUnderflow "The read-only rangeUnderflow property of a ValidityState object indicates if the value of an <input>, after having been edited by the user, does not conform to the constraints set by the element's min attribute.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The control is [suffering from an
underflow](#suffering-from-an-underflow){#the-constraint-validation-api:suffering-from-an-underflow}.

[`rangeOverflow`]{#dom-validitystate-rangeoverflow .dfn
dfn-for="ValidityState" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/rangeOverflow](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/rangeOverflow "The read-only rangeOverflow property of a ValidityState object indicates if the value of an <input>, after having been edited by the user, does not conform to the constraints set by the element's max attribute.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The control is [suffering from an
overflow](#suffering-from-an-overflow){#the-constraint-validation-api:suffering-from-an-overflow}.

[`stepMismatch`]{#dom-validitystate-stepmismatch .dfn
dfn-for="ValidityState" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ValidityState/stepMismatch](https://developer.mozilla.org/en-US/docs/Web/API/ValidityState/stepMismatch "The read-only stepMismatch property of a ValidityState object indicates if the value of an <input>, after having been edited by the user, does not conform to the constraints set by the element's step attribute.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The control is [suffering from a step
mismatch](#suffering-from-a-step-mismatch){#the-constraint-validation-api:suffering-from-a-step-mismatch}.

[`badInput`]{#dom-validitystate-badinput .dfn dfn-for="ValidityState"
dfn-type="attribute"}

The control is [suffering from bad
input](#suffering-from-bad-input){#the-constraint-validation-api:suffering-from-bad-input}.

[`customError`]{#dom-validitystate-customerror .dfn
dfn-for="ValidityState" dfn-type="attribute"}

The control is [suffering from a custom
error](#suffering-from-a-custom-error){#the-constraint-validation-api:suffering-from-a-custom-error}.

[`valid`]{#dom-validitystate-valid .dfn dfn-for="ValidityState"
dfn-type="attribute"}

None of the other conditions are true.

The [check validity steps]{#check-validity-steps .dfn} for an element
`element`{.variable} are:

1.  If `element`{.variable} is a [candidate for constraint
    validation](#candidate-for-constraint-validation){#the-constraint-validation-api:candidate-for-constraint-validation-3}
    and does not [satisfy its
    constraints](#concept-fv-valid){#the-constraint-validation-api:concept-fv-valid},
    then:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-constraint-validation-api:concept-event-fire
        x-internal="concept-event-fire"} named
        [`invalid`{#the-constraint-validation-api:event-invalid-3}](indices.html#event-invalid)
        at `element`{.variable}, with the
        [`cancelable`{#the-constraint-validation-api:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
        attribute initialized to true (though canceling has no effect).

    2.  Return false.

2.  Return true.

The [`checkValidity()`]{#dom-cva-checkvalidity .dfn
dfn-for="HTMLObjectElement,HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="method"} method, when invoked, must run the [check validity
steps](#check-validity-steps){#the-constraint-validation-api:check-validity-steps}
on this element.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/checkValidity](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/checkValidity "The checkValidity() method of the ElementInternals interface checks if the element meets any constraint validation rules applied to it.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
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

The [`checkValidity()`]{#dom-elementinternals-checkvalidity .dfn
dfn-for="ElementInternals" dfn-type="method"} method of the
[`ElementInternals`{#the-constraint-validation-api:elementinternals-3}](custom-elements.html#elementinternals)
interface must run these steps:

1.  Let `element`{.variable} be this
    [`ElementInternals`{#the-constraint-validation-api:elementinternals-4}](custom-elements.html#elementinternals)\'s
    [target
    element](custom-elements.html#internals-target){#the-constraint-validation-api:internals-target-5}.

2.  If `element`{.variable} is not a [form-associated custom
    element](custom-elements.html#form-associated-custom-element){#the-constraint-validation-api:form-associated-custom-element-3},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-constraint-validation-api:notsupportederror-3
    x-internal="notsupportederror"}
    [`DOMException`{#the-constraint-validation-api:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Run the [check validity
    steps](#check-validity-steps){#the-constraint-validation-api:check-validity-steps-2}
    on `element`{.variable}.

The [report validity steps]{#report-validity-steps .dfn} for an element
`element`{.variable} are:

1.  If `element`{.variable} is a [candidate for constraint
    validation](#candidate-for-constraint-validation){#the-constraint-validation-api:candidate-for-constraint-validation-4}
    and does not [satisfy its
    constraints](#concept-fv-valid){#the-constraint-validation-api:concept-fv-valid-2},
    then:

    1.  Let `report`{.variable} be the result of [firing an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#the-constraint-validation-api:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`invalid`{#the-constraint-validation-api:event-invalid-4}](indices.html#event-invalid)
        at `element`{.variable}, with the
        [`cancelable`{#the-constraint-validation-api:dom-event-cancelable-2}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
        attribute initialized to true.

    2.  If `report`{.variable} is true, then report the problems with
        the constraints of this element to the user. When reporting the
        problem with the constraints to the user, the user agent may run
        the [focusing
        steps](interaction.html#focusing-steps){#the-constraint-validation-api:focusing-steps}
        for `element`{.variable}, and may change the scrolling position
        of the document, or perform some other action that brings
        `element`{.variable} to the user\'s attention. User agents may
        report more than one constraint violation, if
        `element`{.variable} suffers from multiple problems at once.

    3.  Return false.

2.  Return true.

The [`reportValidity()`]{#dom-cva-reportvalidity .dfn
dfn-for="HTMLObjectElement,HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="method"} method, when invoked, must run the [report validity
steps](#report-validity-steps){#the-constraint-validation-api:report-validity-steps}
on this element.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ElementInternals/reportValidity](https://developer.mozilla.org/en-US/docs/Web/API/ElementInternals/reportValidity "The reportValidity() method of the ElementInternals interface checks if the element meets any constraint validation rules applied to it.")

Support in all current engines.

::: support
[Firefox98+]{.firefox .yes}[Safari16.4+]{.safari
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

The [`reportValidity()`]{#dom-elementinternals-reportvalidity .dfn
dfn-for="ElementInternals" dfn-type="method"} method of the
[`ElementInternals`{#the-constraint-validation-api:elementinternals-5}](custom-elements.html#elementinternals)
interface must run these steps:

1.  Let `element`{.variable} be this
    [`ElementInternals`{#the-constraint-validation-api:elementinternals-6}](custom-elements.html#elementinternals)\'s
    [target
    element](custom-elements.html#internals-target){#the-constraint-validation-api:internals-target-6}.

2.  If `element`{.variable} is not a [form-associated custom
    element](custom-elements.html#form-associated-custom-element){#the-constraint-validation-api:form-associated-custom-element-4},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#the-constraint-validation-api:notsupportederror-4
    x-internal="notsupportederror"}
    [`DOMException`{#the-constraint-validation-api:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Run the [report validity
    steps](#report-validity-steps){#the-constraint-validation-api:report-validity-steps-2}
    on `element`{.variable}.

The [`validationMessage`]{#dom-cva-validationmessage .dfn
dfn-for="HTMLObjectElement,HTMLInputElement,HTMLButtonElement,HTMLSelectElement,HTMLTextAreaElement,HTMLOutputElement,HTMLFieldSetElement"
dfn-type="attribute"} attribute\'s getter must run these steps:

1.  If this element is not a [candidate for constraint
    validation](#candidate-for-constraint-validation){#the-constraint-validation-api:candidate-for-constraint-validation-5}
    or if this element [satisfies its
    constraints](#concept-fv-valid){#the-constraint-validation-api:concept-fv-valid-3},
    then return the empty string.

2.  Return a suitably localized message that the user agent would show
    the user if this were the only form control with a validity
    constraint problem. If the user agent would not actually show a
    textual message in such a situation (e.g., it would show a graphical
    cue instead), then return a suitably localized message that
    expresses (one or more of) the validity constraint(s) that the
    control does not satisfy. If the element is a [candidate for
    constraint
    validation](#candidate-for-constraint-validation){#the-constraint-validation-api:candidate-for-constraint-validation-6}
    and is [suffering from a custom
    error](#suffering-from-a-custom-error){#the-constraint-validation-api:suffering-from-a-custom-error-2},
    then the [custom validity error
    message](#custom-validity-error-message){#the-constraint-validation-api:custom-validity-error-message-2}
    should be present in the return value.

##### [4.10.20.4]{.secno} Security[](#security-forms){.self-link} {#security-forms}

Servers should not rely on client-side validation. Client-side
validation can be intentionally bypassed by hostile users, and
unintentionally bypassed by users of older user agents or automated
tools that do not implement these features. The constraint validation
features are only intended to improve the user experience, not to
provide any kind of security mechanism.

#### [4.10.21]{.secno} [Form submission]{#form-submission .dfn}[](#form-submission-2){.self-link} {#form-submission-2}

##### [4.10.21.1]{.secno} Introduction[](#introduction-5){.self-link} {#introduction-5}

*This section is non-normative.*

When a form is submitted, the data in the form is converted into the
structure specified by the
[enctype](#concept-fs-enctype){#introduction-5:concept-fs-enctype}, and
then sent to the destination specified by the
[action](#concept-fs-action){#introduction-5:concept-fs-action} using
the given
[method](#concept-fs-method){#introduction-5:concept-fs-method}.

For example, take the following form:

``` html
<form action="/find.cgi" method=get>
 <input type=text name=t>
 <input type=search name=q>
 <input type=submit>
</form>
```

If the user types in \"cats\" in the first field and \"fur\" in the
second, and then hits the submit button, then the user agent will load
`/find.cgi?t=cats&q=fur`.

On the other hand, consider this form:

``` html
<form action="/find.cgi" method=post enctype="multipart/form-data">
 <input type=text name=t>
 <input type=search name=q>
 <input type=submit>
</form>
```

Given the same user input, the result on submission is quite different:
the user agent instead does an HTTP POST to the given URL, with as the
entity body something like the following text:

    ------kYFrd4jNJEgCervE
    Content-Disposition: form-data; name="t"

    cats
    ------kYFrd4jNJEgCervE
    Content-Disposition: form-data; name="q"

    fur
    ------kYFrd4jNJEgCervE--

##### [4.10.21.2]{.secno} Implicit submission[](#implicit-submission){.self-link}

A
[`form`{#implicit-submission:the-form-element}](forms.html#the-form-element)
element\'s [default button]{#default-button .dfn} is the first [submit
button](forms.html#concept-submit-button){#implicit-submission:concept-submit-button}
in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#implicit-submission:tree-order
x-internal="tree-order"} whose [form
owner](#form-owner){#implicit-submission:form-owner} is that
[`form`{#implicit-submission:the-form-element-2}](forms.html#the-form-element)
element.

If the user agent supports letting the user submit a form implicitly
(for example, on some platforms hitting the \"enter\" key while a text
control is
[focused](interaction.html#focused){#implicit-submission:focused}
implicitly submits the form), then doing so for a form, whose [default
button](#default-button){#implicit-submission:default-button} has
[activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#implicit-submission:activation-behaviour
x-internal="activation-behaviour"} and is not
[disabled](#concept-fe-disabled){#implicit-submission:concept-fe-disabled},
must cause the user agent to [fire a `click`
event](webappapis.html#fire-a-click-event){#implicit-submission:fire-a-click-event}
at that [default
button](#default-button){#implicit-submission:default-button-2}.

There are pages on the web that are only usable if there is a way to
implicitly submit forms, so user agents are strongly encouraged to
support this.

If the form has no [submit
button](forms.html#concept-submit-button){#implicit-submission:concept-submit-button-2},
then the implicit submission mechanism must perform the following steps:

1.  If the form has more than one [field that blocks implicit
    submission](#field-that-blocks-implicit-submission){#implicit-submission:field-that-blocks-implicit-submission},
    then return.

2.  [Submit](#concept-form-submit){#implicit-submission:concept-form-submit}
    the
    [`form`{#implicit-submission:the-form-element-3}](forms.html#the-form-element)
    element from the
    [`form`{#implicit-submission:the-form-element-4}](forms.html#the-form-element)
    element itself with *[userInvolvement](#submit-user-involvement)*
    set to
    \"[`activation`{#implicit-submission:uni-activation}](browsing-the-web.html#uni-activation)\".

For the purpose of the previous paragraph, an element is a [field that
blocks implicit submission]{#field-that-blocks-implicit-submission .dfn}
of a
[`form`{#implicit-submission:the-form-element-5}](forms.html#the-form-element)
element if it is an
[`input`{#implicit-submission:the-input-element}](input.html#the-input-element)
element whose [form
owner](#form-owner){#implicit-submission:form-owner-2} is that
[`form`{#implicit-submission:the-form-element-6}](forms.html#the-form-element)
element and whose
[`type`{#implicit-submission:attr-input-type}](input.html#attr-input-type)
attribute is in one of the following states:
[Text](input.html#text-(type=text)-state-and-search-state-(type=search)){#implicit-submission:text-(type=text)-state-and-search-state-(type=search)},
[Search](input.html#text-(type=text)-state-and-search-state-(type=search)){#implicit-submission:text-(type=text)-state-and-search-state-(type=search)-2},
[Telephone](input.html#telephone-state-(type=tel)){#implicit-submission:telephone-state-(type=tel)},
[URL](input.html#url-state-(type=url)){#implicit-submission:url-state-(type=url)},
[Email](input.html#email-state-(type=email)){#implicit-submission:email-state-(type=email)},
[Password](input.html#password-state-(type=password)){#implicit-submission:password-state-(type=password)},
[Date](input.html#date-state-(type=date)){#implicit-submission:date-state-(type=date)},
[Month](input.html#month-state-(type=month)){#implicit-submission:month-state-(type=month)},
[Week](input.html#week-state-(type=week)){#implicit-submission:week-state-(type=week)},
[Time](input.html#time-state-(type=time)){#implicit-submission:time-state-(type=time)},
[Local Date and
Time](input.html#local-date-and-time-state-(type=datetime-local)){#implicit-submission:local-date-and-time-state-(type=datetime-local)},
[Number](input.html#number-state-(type=number)){#implicit-submission:number-state-(type=number)}

##### [4.10.21.3]{.secno} Form submission algorithm[](#form-submission-algorithm){.self-link}

Each
[`form`{#form-submission-algorithm:the-form-element}](forms.html#the-form-element)
element has a [constructing entry list]{#constructing-entry-list .dfn}
boolean, initially false.

Each
[`form`{#form-submission-algorithm:the-form-element-2}](forms.html#the-form-element)
element has a [firing submission events]{#firing-submission-events .dfn}
boolean, initially false.

To [submit]{#concept-form-submit .dfn} a
[`form`{#form-submission-algorithm:the-form-element-3}](forms.html#the-form-element)
element `form`{.variable} from an element `submitter`{.variable}
(typically a button), given an optional boolean
[`submitted from `{.variable}`submit()`` method`{.variable}]{#submit-subbmitted-from-method
.dfn} (default false) and an optional [user navigation
involvement](browsing-the-web.html#user-navigation-involvement){#form-submission-algorithm:user-navigation-involvement}
[`userInvolvement`{.variable}]{#submit-user-involvement .dfn} (default
\"[`none`{#form-submission-algorithm:uni-none}](browsing-the-web.html#uni-none)\"):

1.  If `form`{.variable} [cannot
    navigate](links.html#cannot-navigate){#form-submission-algorithm:cannot-navigate},
    then return.

2.  If `form`{.variable}\'s [constructing entry
    list](#constructing-entry-list){#form-submission-algorithm:constructing-entry-list}
    is true, then return.

3.  Let `form document`{.variable} be `form`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#form-submission-algorithm:node-document
    x-internal="node-document"}.

4.  [If `form document`{.variable}\'s [active sandboxing flag
    set](browsers.html#active-sandboxing-flag-set){#form-submission-algorithm:active-sandboxing-flag-set}
    has its [sandboxed forms browsing context
    flag](browsers.html#sandboxed-forms-browsing-context-flag){#form-submission-algorithm:sandboxed-forms-browsing-context-flag}
    set, then return.]{#sandboxSubmitBlocked}

5.  If
    `submitted from `{.variable}[`submit()`{#form-submission-algorithm:dom-form-submit}](forms.html#dom-form-submit)` method`{.variable}
    is false, then:

    1.  If `form`{.variable}\'s [firing submission
        events](#firing-submission-events){#form-submission-algorithm:firing-submission-events}
        is true, then return.

    2.  Set `form`{.variable}\'s [firing submission
        events](#firing-submission-events){#form-submission-algorithm:firing-submission-events-2}
        to true.

    3.  For each element `field`{.variable} in the list of [submittable
        elements](forms.html#category-submit){#form-submission-algorithm:category-submit}
        whose [form
        owner](#form-owner){#form-submission-algorithm:form-owner} is
        `form`{.variable}, set `field`{.variable}\'s [user
        validity](#user-validity){#form-submission-algorithm:user-validity}
        to true.

    4.  If the `submitter`{.variable} element\'s [no-validate
        state](#concept-fs-novalidate){#form-submission-algorithm:concept-fs-novalidate}
        is false, then [interactively validate the
        constraints](#interactively-validate-the-constraints){#form-submission-algorithm:interactively-validate-the-constraints}
        of `form`{.variable} and examine the result. If the result is
        negative (i.e., the constraint validation concluded that there
        were invalid fields and probably informed the user of this),
        then:

        1.  Set `form`{.variable}\'s [firing submission
            events](#firing-submission-events){#form-submission-algorithm:firing-submission-events-3}
            to false.

        2.  Return.

    5.  Let `submitterButton`{.variable} be null if
        `submitter`{.variable} is `form`{.variable}. Otherwise, let
        `submitterButton`{.variable} be `submitter`{.variable}.

    6.  Let `shouldContinue`{.variable} be the result of [firing an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#form-submission-algorithm:concept-event-fire
        x-internal="concept-event-fire"} named
        [`submit`{#form-submission-algorithm:event-submit}](indices.html#event-submit)
        at `form`{.variable} using
        [`SubmitEvent`{#form-submission-algorithm:submitevent}](#submitevent),
        with the
        [`submitter`{#form-submission-algorithm:dom-submitevent-submitter}](#dom-submitevent-submitter)
        attribute initialized to `submitterButton`{.variable}, the
        [`bubbles`{#form-submission-algorithm:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
        attribute initialized to true, and the
        [`cancelable`{#form-submission-algorithm:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
        attribute initialized to true.

    7.  Set `form`{.variable}\'s [firing submission
        events](#firing-submission-events){#form-submission-algorithm:firing-submission-events-4}
        to false.

    8.  If `shouldContinue`{.variable} is false, then return.

    9.  If `form`{.variable} [cannot
        navigate](links.html#cannot-navigate){#form-submission-algorithm:cannot-navigate-2},
        then return.

        [Cannot
        navigate](links.html#cannot-navigate){#form-submission-algorithm:cannot-navigate-3}
        is run again as dispatching the
        [`submit`{#form-submission-algorithm:event-submit-2}](indices.html#event-submit)
        event could have changed the outcome.

6.  Let `encoding`{.variable} be the result of [picking an encoding for
    the
    form](#picking-an-encoding-for-the-form){#form-submission-algorithm:picking-an-encoding-for-the-form}.

7.  Let `entry list`{.variable} be the result of [constructing the entry
    list](#constructing-the-form-data-set){#form-submission-algorithm:constructing-the-form-data-set}
    with `form`{.variable}, `submitter`{.variable}, and
    `encoding`{.variable}.

8.  [Assert](https://infra.spec.whatwg.org/#assert){#form-submission-algorithm:assert
    x-internal="assert"}: `entry list`{.variable} is not null.

9.  If `form`{.variable} [cannot
    navigate](links.html#cannot-navigate){#form-submission-algorithm:cannot-navigate-4},
    then return.

    [Cannot
    navigate](links.html#cannot-navigate){#form-submission-algorithm:cannot-navigate-5}
    is run again as dispatching the
    [`formdata`{#form-submission-algorithm:event-formdata}](indices.html#event-formdata)
    event in [constructing the entry
    list](#constructing-the-form-data-set){#form-submission-algorithm:constructing-the-form-data-set-2}
    could have changed the outcome.

10. Let `method`{.variable} be the `submitter`{.variable} element\'s
    [method](#concept-fs-method){#form-submission-algorithm:concept-fs-method}.

11. If `method`{.variable} is
    [dialog](#attr-fs-method-dialog){#form-submission-algorithm:attr-fs-method-dialog},
    then:

    1.  If `form`{.variable} does not have an ancestor
        [`dialog`{#form-submission-algorithm:the-dialog-element}](interactive-elements.html#the-dialog-element)
        element, then return.

    2.  Let `subject`{.variable} be `form`{.variable}\'s nearest
        ancestor
        [`dialog`{#form-submission-algorithm:the-dialog-element-2}](interactive-elements.html#the-dialog-element)
        element.

    3.  Let `result`{.variable} be null.

    4.  If `submitter`{.variable} is an
        [`input`{#form-submission-algorithm:the-input-element}](input.html#the-input-element)
        element whose
        [`type`{#form-submission-algorithm:attr-input-type}](input.html#attr-input-type)
        attribute is in the [Image
        Button](input.html#image-button-state-(type=image)){#form-submission-algorithm:image-button-state-(type=image)}
        state, then:

        1.  Let (`x`{.variable}, `y`{.variable}) be the [selected
            coordinate](input.html#concept-input-type-image-coordinate){#form-submission-algorithm:concept-input-type-image-coordinate}.

        2.  Set `result`{.variable} to the concatenation of
            `x`{.variable}, \"`,`\", and `y`{.variable}.

    5.  Otherwise, if `submitter`{.variable} is a [submit
        button](forms.html#concept-submit-button){#form-submission-algorithm:concept-submit-button},
        then set `result`{.variable} to `submitter`{.variable}\'s
        [optional
        value](#concept-fe-optional-value){#form-submission-algorithm:concept-fe-optional-value}.

    6.  [Close the
        dialog](interactive-elements.html#close-the-dialog){#form-submission-algorithm:close-the-dialog}
        `subject`{.variable} with `result`{.variable} and null.

    7.  Return.

12. Let `action`{.variable} be the `submitter`{.variable} element\'s
    [action](#concept-fs-action){#form-submission-algorithm:concept-fs-action}.

13. If `action`{.variable} is the empty string, let `action`{.variable}
    be the
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#form-submission-algorithm:the-document's-address
    x-internal="the-document's-address"} of the
    `form document`{.variable}.

14. Let `parsed action`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#form-submission-algorithm:encoding-parsing-a-url}
    given `action`{.variable}, relative to `submitter`{.variable}\'s
    [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#form-submission-algorithm:node-document-2
    x-internal="node-document"}.

15. If `parsed action`{.variable} is failure, then return.

16. Let `scheme`{.variable} be the
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#form-submission-algorithm:concept-url-scheme
    x-internal="concept-url-scheme"} of `parsed action`{.variable}.

17. Let `enctype`{.variable} be the `submitter`{.variable} element\'s
    [enctype](#concept-fs-enctype){#form-submission-algorithm:concept-fs-enctype}.

18. Let `formTarget`{.variable} be null.

19. If the `submitter`{.variable} element is a [submit
    button](forms.html#concept-submit-button){#form-submission-algorithm:concept-submit-button-2}
    and it has a
    [`formtarget`{#form-submission-algorithm:attr-fs-formtarget}](#attr-fs-formtarget)
    attribute, then set `formTarget`{.variable} to the
    [`formtarget`{#form-submission-algorithm:attr-fs-formtarget-2}](#attr-fs-formtarget)
    attribute value.

20. Let `target`{.variable} be the result of [getting an element\'s
    target](semantics.html#get-an-element's-target){#form-submission-algorithm:get-an-element's-target}
    given `submitter`{.variable}\'s [form
    owner](#form-owner){#form-submission-algorithm:form-owner-2} and
    `formTarget`{.variable}.

21. Let `noopener`{.variable} be the result of [getting an element\'s
    noopener](links.html#get-an-element's-noopener){#form-submission-algorithm:get-an-element's-noopener}
    with `form`{.variable}, `parsed action`{.variable}, and
    `target`{.variable}.

22. Let `targetNavigable`{.variable} be the first return value of
    applying [the rules for choosing a
    navigable](document-sequences.html#the-rules-for-choosing-a-navigable){#form-submission-algorithm:the-rules-for-choosing-a-navigable}
    given `target`{.variable}, `form`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#form-submission-algorithm:node-navigable},
    and `noopener`{.variable}.

23. If `targetNavigable`{.variable} is null, then return.

24. Let `historyHandling`{.variable} be
    \"[`auto`{#form-submission-algorithm:navigationhistorybehavior-auto}](browsing-the-web.html#navigationhistorybehavior-auto)\".

25. If `form document`{.variable} equals `targetNavigable`{.variable}\'s
    [active
    document](document-sequences.html#nav-document){#form-submission-algorithm:nav-document},
    and `form document`{.variable} has not yet [completely
    loaded](document-lifecycle.html#completely-loaded){#form-submission-algorithm:completely-loaded},
    then set `historyHandling`{.variable} to
    \"[`replace`{#form-submission-algorithm:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

26. Select the appropriate row in the table below based on
    `scheme`{.variable} as given by the first cell of each row. Then,
    select the appropriate cell on that row based on `method`{.variable}
    as given in the first cell of each column. Then, jump to the steps
    named in that cell and defined below the table.

    [GET](#attr-fs-method-get){#form-submission-algorithm:attr-fs-method-get}

    [POST](#attr-fs-method-post){#form-submission-algorithm:attr-fs-method-post}

    `http`

    [Mutate action
    URL](#submit-mutate-action){#form-submission-algorithm:submit-mutate-action}

    [Submit as entity
    body](#submit-body){#form-submission-algorithm:submit-body}

    `https`

    [Mutate action
    URL](#submit-mutate-action){#form-submission-algorithm:submit-mutate-action-2}

    [Submit as entity
    body](#submit-body){#form-submission-algorithm:submit-body-2}

    `ftp`

    [Get action
    URL](#submit-get-action){#form-submission-algorithm:submit-get-action}

    [Get action
    URL](#submit-get-action){#form-submission-algorithm:submit-get-action-2}

    `javascript`

    [Get action
    URL](#submit-get-action){#form-submission-algorithm:submit-get-action-3}

    [Get action
    URL](#submit-get-action){#form-submission-algorithm:submit-get-action-4}

    `data`

    [Mutate action
    URL](#submit-mutate-action){#form-submission-algorithm:submit-mutate-action-3}

    [Get action
    URL](#submit-get-action){#form-submission-algorithm:submit-get-action-5}

    `mailto`

    [Mail with
    headers](#submit-mailto-headers){#form-submission-algorithm:submit-mailto-headers}

    [Mail as
    body](#submit-mailto-body){#form-submission-algorithm:submit-mailto-body}

    If `scheme`{.variable} is not one of those listed in this table,
    then the behavior is not defined by this specification. User agents
    should, in the absence of another specification defining this, act
    in a manner analogous to that defined in this specification for
    similar schemes.

    Each
    [`form`{#form-submission-algorithm:the-form-element-4}](forms.html#the-form-element)
    element has a [planned navigation]{#planned-navigation .dfn}, which
    is either null or a
    [task](webappapis.html#concept-task){#form-submission-algorithm:concept-task};
    when the
    [`form`{#form-submission-algorithm:the-form-element-5}](forms.html#the-form-element)
    is first created, its [planned
    navigation](#planned-navigation){#form-submission-algorithm:planned-navigation}
    must be set to null. In the behaviors described below, when the user
    agent is required to [plan to navigate]{#plan-to-navigate .dfn} to a
    [URL](https://url.spec.whatwg.org/#concept-url){#form-submission-algorithm:url
    x-internal="url"} `url`{.variable} given an optional [POST
    resource](browsing-the-web.html#post-resource){#form-submission-algorithm:post-resource}-or-null
    `postResource`{.variable} (default null), it must run the following
    steps:

    1.  Let `referrerPolicy`{.variable} be the empty string.

    2.  If the
        [`form`{#form-submission-algorithm:the-form-element-6}](forms.html#the-form-element)
        element\'s [link types](links.html#linkTypes) include the
        [`noreferrer`{#form-submission-algorithm:link-type-noreferrer}](links.html#link-type-noreferrer)
        keyword, then set `referrerPolicy`{.variable} to
        \"`no-referrer`\".

    3.  If the
        [`form`{#form-submission-algorithm:the-form-element-7}](forms.html#the-form-element)
        has a non-null [planned
        navigation](#planned-navigation){#form-submission-algorithm:planned-navigation-2},
        remove it from its [task
        queue](webappapis.html#task-queue){#form-submission-algorithm:task-queue}.

    4.  [Queue an element
        task](webappapis.html#queue-an-element-task){#form-submission-algorithm:queue-an-element-task}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#form-submission-algorithm:dom-manipulation-task-source}
        given the
        [`form`{#form-submission-algorithm:the-form-element-8}](forms.html#the-form-element)
        element and the following steps:

        1.  Set the
            [`form`{#form-submission-algorithm:the-form-element-9}](forms.html#the-form-element)\'s
            [planned
            navigation](#planned-navigation){#form-submission-algorithm:planned-navigation-3}
            to null.

        2.  [Navigate](browsing-the-web.html#navigate){#form-submission-algorithm:navigate}
            `targetNavigable`{.variable} to `url`{.variable} using the
            [`form`{#form-submission-algorithm:the-form-element-10}](forms.html#the-form-element)
            element\'s [node
            document](https://dom.spec.whatwg.org/#concept-node-document){#form-submission-algorithm:node-document-3
            x-internal="node-document"}, with
            *[historyHandling](browsing-the-web.html#navigation-hh)* set
            to `historyHandling`{.variable},
            *[userInvolvement](browsing-the-web.html#navigation-user-involvement)*
            set to `userInvolvement`{.variable},
            *[sourceElement](browsing-the-web.html#navigation-source-element)*
            set to `submitter`{.variable},
            *[referrerPolicy](browsing-the-web.html#navigation-referrer-policy)*
            set to `referrerPolicy`{.variable},
            *[documentResource](browsing-the-web.html#navigation-resource)*
            set to `postResource`{.variable}, and
            *[formDataEntryList](browsing-the-web.html#navigation-form-data-entry-list)*
            set to `entry list`{.variable}.

    5.  Set the
        [`form`{#form-submission-algorithm:the-form-element-11}](forms.html#the-form-element)\'s
        [planned
        navigation](#planned-navigation){#form-submission-algorithm:planned-navigation-4}
        to the just-queued
        [task](webappapis.html#concept-task){#form-submission-algorithm:concept-task-2}.

    The behaviors are as follows:

    [Mutate action URL]{#submit-mutate-action .dfn}

    :   Let `pairs`{.variable} be the result of [converting to a list of
        name-value
        pairs](#convert-to-a-list-of-name-value-pairs){#form-submission-algorithm:convert-to-a-list-of-name-value-pairs}
        with `entry list`{.variable}.

        Let `query`{.variable} be the result of running the
        [`application/x-www-form-urlencoded`
        serializer](https://url.spec.whatwg.org/#concept-urlencoded-serializer){#form-submission-algorithm:application/x-www-form-urlencoded-serializer
        x-internal="application/x-www-form-urlencoded-serializer"} with
        `pairs`{.variable} and `encoding`{.variable}.

        Set `parsed action`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#form-submission-algorithm:concept-url-query
        x-internal="concept-url-query"} component to `query`{.variable}.

        [Plan to
        navigate](#plan-to-navigate){#form-submission-algorithm:plan-to-navigate}
        to `parsed action`{.variable}.

    [Submit as entity body]{#submit-body .dfn}

    :   [Assert](https://infra.spec.whatwg.org/#assert){#form-submission-algorithm:assert-2
        x-internal="assert"}: `method`{.variable} is
        [POST](#attr-fs-method-post){#form-submission-algorithm:attr-fs-method-post-2}.

        Switch on `enctype`{.variable}:

        [`application/x-www-form-urlencoded`{#form-submission-algorithm:attr-fs-enctype-urlencoded}](#attr-fs-enctype-urlencoded)

        :   Let `pairs`{.variable} be the result of [converting to a
            list of name-value
            pairs](#convert-to-a-list-of-name-value-pairs){#form-submission-algorithm:convert-to-a-list-of-name-value-pairs-2}
            with `entry list`{.variable}.

            Let `body`{.variable} be the result of running the
            [`application/x-www-form-urlencoded`
            serializer](https://url.spec.whatwg.org/#concept-urlencoded-serializer){#form-submission-algorithm:application/x-www-form-urlencoded-serializer-2
            x-internal="application/x-www-form-urlencoded-serializer"}
            with `pairs`{.variable} and `encoding`{.variable}.

            Set `body`{.variable} to the result of
            [encoding](https://encoding.spec.whatwg.org/#utf-8-encode){#form-submission-algorithm:utf-8-encode
            x-internal="utf-8-encode"} `body`{.variable}.

            Let `mimeType`{.variable} be
            \`[`application/x-www-form-urlencoded`{#form-submission-algorithm:application/x-www-form-urlencoded}](https://url.spec.whatwg.org/#concept-urlencoded){x-internal="application/x-www-form-urlencoded"}\`.

        [`multipart/form-data`{#form-submission-algorithm:attr-fs-enctype-formdata}](#attr-fs-enctype-formdata)

        :   Let `body`{.variable} be the result of running the
            [`multipart/form-data` encoding
            algorithm](#multipart/form-data-encoding-algorithm){#form-submission-algorithm:multipart/form-data-encoding-algorithm}
            with `entry list`{.variable} and `encoding`{.variable}.

            Let `mimeType`{.variable} be the [isomorphic
            encoding](https://infra.spec.whatwg.org/#isomorphic-encode){#form-submission-algorithm:isomorphic-encode
            x-internal="isomorphic-encode"} of the concatenation of
            \"`multipart/form-data; boundary=`\" and the
            [`multipart/form-data` boundary
            string](#multipart/form-data-boundary-string){#form-submission-algorithm:multipart/form-data-boundary-string}
            generated by the [`multipart/form-data` encoding
            algorithm](#multipart/form-data-encoding-algorithm){#form-submission-algorithm:multipart/form-data-encoding-algorithm-2}.

        [`text/plain`{#form-submission-algorithm:attr-fs-enctype-text}](#attr-fs-enctype-text)

        :   Let `pairs`{.variable} be the result of [converting to a
            list of name-value
            pairs](#convert-to-a-list-of-name-value-pairs){#form-submission-algorithm:convert-to-a-list-of-name-value-pairs-3}
            with `entry list`{.variable}.

            Let `body`{.variable} be the result of running the
            [`text/plain` encoding
            algorithm](#text/plain-encoding-algorithm){#form-submission-algorithm:text/plain-encoding-algorithm}
            with `pairs`{.variable}.

            Set `body`{.variable} to the result of
            [encoding](https://encoding.spec.whatwg.org/#encode){#form-submission-algorithm:encode
            x-internal="encode"} `body`{.variable} using
            `encoding`{.variable}.

            Let `mimeType`{.variable} be
            \`[`text/plain`{#form-submission-algorithm:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}\`.

        [Plan to
        navigate](#plan-to-navigate){#form-submission-algorithm:plan-to-navigate-2}
        to `parsed action`{.variable} given a [POST
        resource](browsing-the-web.html#post-resource){#form-submission-algorithm:post-resource-2}
        whose [request
        body](browsing-the-web.html#post-resource-request-body){#form-submission-algorithm:post-resource-request-body}
        is `body`{.variable} and [request
        content-type](browsing-the-web.html#post-resource-request-content-type){#form-submission-algorithm:post-resource-request-content-type}
        is `mimeType`{.variable}.

    [Get action URL]{#submit-get-action .dfn}

    :   [Plan to
        navigate](#plan-to-navigate){#form-submission-algorithm:plan-to-navigate-3}
        to `parsed action`{.variable}.

        `entry list`{.variable} is discarded.

    [Mail with headers]{#submit-mailto-headers .dfn}

    :   Let `pairs`{.variable} be the result of [converting to a list of
        name-value
        pairs](#convert-to-a-list-of-name-value-pairs){#form-submission-algorithm:convert-to-a-list-of-name-value-pairs-4}
        with `entry list`{.variable}.

        Let `headers`{.variable} be the result of running the
        [`application/x-www-form-urlencoded`
        serializer](https://url.spec.whatwg.org/#concept-urlencoded-serializer){#form-submission-algorithm:application/x-www-form-urlencoded-serializer-3
        x-internal="application/x-www-form-urlencoded-serializer"} with
        `pairs`{.variable} and `encoding`{.variable}.

        Replace occurrences of U+002B PLUS SIGN characters (+) in
        `headers`{.variable} with the string \"`%20`\".

        Set `parsed action`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#form-submission-algorithm:concept-url-query-2
        x-internal="concept-url-query"} to `headers`{.variable}.

        [Plan to
        navigate](#plan-to-navigate){#form-submission-algorithm:plan-to-navigate-4}
        to `parsed action`{.variable}.

    [Mail as body]{#submit-mailto-body .dfn}

    :   Let `pairs`{.variable} be the result of [converting to a list of
        name-value
        pairs](#convert-to-a-list-of-name-value-pairs){#form-submission-algorithm:convert-to-a-list-of-name-value-pairs-5}
        with `entry list`{.variable}.

        Switch on `enctype`{.variable}:

        [`text/plain`{#form-submission-algorithm:attr-fs-enctype-text-2}](#attr-fs-enctype-text)

        :   Let `body`{.variable} be the result of running the
            [`text/plain` encoding
            algorithm](#text/plain-encoding-algorithm){#form-submission-algorithm:text/plain-encoding-algorithm-2}
            with `pairs`{.variable}.

            Set `body`{.variable} to the result of running [UTF-8
            percent-encode](https://url.spec.whatwg.org/#string-utf-8-percent-encode){#form-submission-algorithm:utf-8-percent-encode
            x-internal="utf-8-percent-encode"} on `body`{.variable}
            using the [default encode
            set](https://url.spec.whatwg.org/#default-encode-set){#form-submission-algorithm:default-encode-set
            x-internal="default-encode-set"}.
            [\[URL\]](references.html#refsURL)

        Otherwise

        :   Let `body`{.variable} be the result of running the
            [`application/x-www-form-urlencoded`
            serializer](https://url.spec.whatwg.org/#concept-urlencoded-serializer){#form-submission-algorithm:application/x-www-form-urlencoded-serializer-4
            x-internal="application/x-www-form-urlencoded-serializer"}
            with `pairs`{.variable} and `encoding`{.variable}.

        If `parsed action`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#form-submission-algorithm:concept-url-query-3
        x-internal="concept-url-query"} is null, then set it to the
        empty string.

        If `parsed action`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#form-submission-algorithm:concept-url-query-4
        x-internal="concept-url-query"} is not the empty string, then
        append a single U+0026 AMPERSAND character (&) to it.

        Append \"`body=`\" to `parsed action`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#form-submission-algorithm:concept-url-query-5
        x-internal="concept-url-query"}.

        Append `body`{.variable} to `parsed action`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#form-submission-algorithm:concept-url-query-6
        x-internal="concept-url-query"}.

        [Plan to
        navigate](#plan-to-navigate){#form-submission-algorithm:plan-to-navigate-5}
        to `parsed action`{.variable}.

##### [4.10.21.4]{.secno} Constructing the entry list[](#constructing-form-data-set){.self-link} {#constructing-form-data-set}

An [entry list]{#entry-list .dfn export=""} is a
[list](https://infra.spec.whatwg.org/#list){#constructing-form-data-set:list
x-internal="list"} of
[entries](#form-entry){#constructing-form-data-set:form-entry},
typically representing the contents of a form. An [entry]{#form-entry
.dfn dfn-for="entry list" export=""} is a tuple consisting of a
[name]{#form-entry-name .dfn dfn-for="entry list/entry" export=""} (a
[scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#constructing-form-data-set:scalar-value-string
x-internal="scalar-value-string"}) and a [value]{#form-entry-value .dfn
dfn-for="entry list/entry" export=""} (either a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#constructing-form-data-set:scalar-value-string-2
x-internal="scalar-value-string"} or a
[`File`{#constructing-form-data-set:file}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
object).

To [create an entry]{#create-an-entry .dfn dfn-for="entry list"
export=""} given a string `name`{.variable}, a string or
[`Blob`{#constructing-form-data-set:blob}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
object `value`{.variable}, and optionally a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#constructing-form-data-set:scalar-value-string-3
x-internal="scalar-value-string"} `filename`{.variable}:

1.  Set `name`{.variable} to the result of
    [converting](https://infra.spec.whatwg.org/#javascript-string-convert){#constructing-form-data-set:convert
    x-internal="convert"} `name`{.variable} into a [scalar value
    string](https://infra.spec.whatwg.org/#scalar-value-string){#constructing-form-data-set:scalar-value-string-4
    x-internal="scalar-value-string"}.

2.  If `value`{.variable} is a string, then set `value`{.variable} to
    the result of
    [converting](https://infra.spec.whatwg.org/#javascript-string-convert){#constructing-form-data-set:convert-2
    x-internal="convert"} `value`{.variable} into a [scalar value
    string](https://infra.spec.whatwg.org/#scalar-value-string){#constructing-form-data-set:scalar-value-string-5
    x-internal="scalar-value-string"}.

3.  Otherwise:

    1.  If `value`{.variable} is not a
        [`File`{#constructing-form-data-set:file-2}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
        object, then set `value`{.variable} to a new
        [`File`{#constructing-form-data-set:file-3}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
        object, representing the same bytes, whose
        [`name`{#constructing-form-data-set:dom-file-name}](https://w3c.github.io/FileAPI/#dfn-name){x-internal="dom-file-name"}
        attribute value is \"`blob`\".

    2.  If `filename`{.variable} is given, then set `value`{.variable}
        to a new
        [`File`{#constructing-form-data-set:file-4}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
        object, representing the same bytes, whose
        [`name`{#constructing-form-data-set:dom-file-name-2}](https://w3c.github.io/FileAPI/#dfn-name){x-internal="dom-file-name"}
        attribute is `filename`{.variable}.

    These operations will create a new
    [`File`{#constructing-form-data-set:file-5}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
    object if either `filename`{.variable} is given or the passed
    [`Blob`{#constructing-form-data-set:blob-2}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
    is not a
    [`File`{#constructing-form-data-set:file-6}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
    object. In those cases, the identity of the passed
    [`Blob`{#constructing-form-data-set:blob-3}](https://w3c.github.io/FileAPI/#dfn-Blob){x-internal="blob"}
    object is not kept.

4.  Return an
    [entry](#form-entry){#constructing-form-data-set:form-entry-2} whose
    [name](#form-entry-name){#constructing-form-data-set:form-entry-name}
    is `name`{.variable} and whose
    [value](#form-entry-value){#constructing-form-data-set:form-entry-value}
    is `value`{.variable}.

To [construct the entry list]{#constructing-the-form-data-set .dfn
lt="constructing the entry list" export=""} given a `form`{.variable},
an optional `submitter`{.variable} (default null), and an optional
`encoding`{.variable} (default
[UTF-8](https://encoding.spec.whatwg.org/#utf-8){#constructing-form-data-set:utf-8
x-internal="utf-8"}):

1.  If `form`{.variable}\'s [constructing entry
    list](#constructing-entry-list){#constructing-form-data-set:constructing-entry-list}
    is true, then return null.

2.  Set `form`{.variable}\'s [constructing entry
    list](#constructing-entry-list){#constructing-form-data-set:constructing-entry-list-2}
    to true.

3.  Let `controls`{.variable} be a list of all the [submittable
    elements](forms.html#category-submit){#constructing-form-data-set:category-submit}
    whose [form
    owner](#form-owner){#constructing-form-data-set:form-owner} is
    `form`{.variable}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#constructing-form-data-set:tree-order
    x-internal="tree-order"}.

4.  Let `entry list`{.variable} be a new empty [entry
    list](#entry-list){#constructing-form-data-set:entry-list}.

5.  For each element `field`{.variable} in `controls`{.variable}, in
    [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#constructing-form-data-set:tree-order-2
    x-internal="tree-order"}:

    1.  If any of the following are true:

        - `field`{.variable} has a
          [`datalist`{#constructing-form-data-set:the-datalist-element}](form-elements.html#the-datalist-element)
          element ancestor;

        - `field`{.variable} is
          [disabled](#concept-fe-disabled){#constructing-form-data-set:concept-fe-disabled};

        - `field`{.variable} is a
          [button](forms.html#concept-button){#constructing-form-data-set:concept-button}
          but it is not `submitter`{.variable};

        - `field`{.variable} is an
          [`input`{#constructing-form-data-set:the-input-element}](input.html#the-input-element)
          element whose
          [`type`{#constructing-form-data-set:attr-input-type}](input.html#attr-input-type)
          attribute is in the
          [Checkbox](input.html#checkbox-state-(type=checkbox)){#constructing-form-data-set:checkbox-state-(type=checkbox)}
          state and whose
          [checkedness](#concept-fe-checked){#constructing-form-data-set:concept-fe-checked}
          is false; or

        - `field`{.variable} is an
          [`input`{#constructing-form-data-set:the-input-element-2}](input.html#the-input-element)
          element whose
          [`type`{#constructing-form-data-set:attr-input-type-2}](input.html#attr-input-type)
          attribute is in the [Radio
          Button](input.html#radio-button-state-(type=radio)){#constructing-form-data-set:radio-button-state-(type=radio)}
          state and whose
          [checkedness](#concept-fe-checked){#constructing-form-data-set:concept-fe-checked-2}
          is false,

        then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#constructing-form-data-set:continue
        x-internal="continue"}.

    2.  If the `field`{.variable} element is an
        [`input`{#constructing-form-data-set:the-input-element-3}](input.html#the-input-element)
        element whose
        [`type`{#constructing-form-data-set:attr-input-type-3}](input.html#attr-input-type)
        attribute is in the [Image
        Button](input.html#image-button-state-(type=image)){#constructing-form-data-set:image-button-state-(type=image)}
        state, then:

        1.  If the `field`{.variable} element is not
            `submitter`{.variable}, then
            [continue](https://infra.spec.whatwg.org/#iteration-continue){#constructing-form-data-set:continue-2
            x-internal="continue"}.

        2.  If the `field`{.variable} element has a
            [`name`{#constructing-form-data-set:attr-fe-name}](#attr-fe-name)
            attribute specified and its value is not the empty string,
            let `name`{.variable} be that value followed by U+002E (.).
            Otherwise, let `name`{.variable} be the empty string.

        3.  Let `name`{.variable}~`x`{.variable}~ be the concatenation
            of `name`{.variable} and U+0078 (x).

        4.  Let `name`{.variable}~`y`{.variable}~ be the concatenation
            of `name`{.variable} and U+0079 (y).

        5.  Let (`x`{.variable}, `y`{.variable}) be the [selected
            coordinate](input.html#concept-input-type-image-coordinate){#constructing-form-data-set:concept-input-type-image-coordinate}.

        6.  [Create an
            entry](#create-an-entry){#constructing-form-data-set:create-an-entry}
            with `name`{.variable}~`x`{.variable}~ and `x`{.variable},
            and
            [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append
            x-internal="list-append"} it to `entry list`{.variable}.

        7.  [Create an
            entry](#create-an-entry){#constructing-form-data-set:create-an-entry-2}
            with `name`{.variable}~`y`{.variable}~ and `y`{.variable},
            and
            [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-2
            x-internal="list-append"} it to `entry list`{.variable}.

        8.  [Continue](https://infra.spec.whatwg.org/#iteration-continue){#constructing-form-data-set:continue-3
            x-internal="continue"}.

    3.  If the `field`{.variable} is a [form-associated custom
        element](custom-elements.html#form-associated-custom-element){#constructing-form-data-set:form-associated-custom-element},
        then perform the [entry construction
        algorithm](custom-elements.html#face-entry-construction){#constructing-form-data-set:face-entry-construction}
        given `field`{.variable} and `entry list`{.variable}, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#constructing-form-data-set:continue-4
        x-internal="continue"}.

    4.  If either the `field`{.variable} element does not have a
        [`name`{#constructing-form-data-set:attr-fe-name-2}](#attr-fe-name)
        attribute specified, or its
        [`name`{#constructing-form-data-set:attr-fe-name-3}](#attr-fe-name)
        attribute\'s value is the empty string, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#constructing-form-data-set:continue-5
        x-internal="continue"}.

    5.  Let `name`{.variable} be the value of the `field`{.variable}
        element\'s
        [`name`{#constructing-form-data-set:attr-fe-name-4}](#attr-fe-name)
        attribute.

    6.  If the `field`{.variable} element is a
        [`select`{#constructing-form-data-set:the-select-element}](form-elements.html#the-select-element)
        element, then for each
        [`option`{#constructing-form-data-set:the-option-element}](form-elements.html#the-option-element)
        element in the
        [`select`{#constructing-form-data-set:the-select-element-2}](form-elements.html#the-select-element)
        element\'s [list of
        options](form-elements.html#concept-select-option-list){#constructing-form-data-set:concept-select-option-list}
        whose
        [selectedness](form-elements.html#concept-option-selectedness){#constructing-form-data-set:concept-option-selectedness}
        is true and that is not
        [disabled](form-elements.html#concept-option-disabled){#constructing-form-data-set:concept-option-disabled},
        [create an
        entry](#create-an-entry){#constructing-form-data-set:create-an-entry-3}
        with `name`{.variable} and the
        [value](form-elements.html#concept-option-value){#constructing-form-data-set:concept-option-value}
        of the
        [`option`{#constructing-form-data-set:the-option-element-2}](form-elements.html#the-option-element)
        element, and
        [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-3
        x-internal="list-append"} it to `entry list`{.variable}.

    7.  Otherwise, if the `field`{.variable} element is an
        [`input`{#constructing-form-data-set:the-input-element-4}](input.html#the-input-element)
        element whose
        [`type`{#constructing-form-data-set:attr-input-type-4}](input.html#attr-input-type)
        attribute is in the
        [Checkbox](input.html#checkbox-state-(type=checkbox)){#constructing-form-data-set:checkbox-state-(type=checkbox)-2}
        state or the [Radio
        Button](input.html#radio-button-state-(type=radio)){#constructing-form-data-set:radio-button-state-(type=radio)-2}
        state, then:

        1.  If the `field`{.variable} element has a
            [`value`{#constructing-form-data-set:attr-input-value}](input.html#attr-input-value)
            attribute specified, then let `value`{.variable} be the
            value of that attribute; otherwise, let `value`{.variable}
            be the string \"`on`\".

        2.  [Create an
            entry](#create-an-entry){#constructing-form-data-set:create-an-entry-4}
            with `name`{.variable} and `value`{.variable}, and
            [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-4
            x-internal="list-append"} it to `entry list`{.variable}.

    8.  Otherwise, if the `field`{.variable} element is an
        [`input`{#constructing-form-data-set:the-input-element-5}](input.html#the-input-element)
        element whose
        [`type`{#constructing-form-data-set:attr-input-type-5}](input.html#attr-input-type)
        attribute is in the [File
        Upload](input.html#file-upload-state-(type=file)){#constructing-form-data-set:file-upload-state-(type=file)}
        state, then:

        1.  If there are no [selected
            files](input.html#concept-input-type-file-selected){#constructing-form-data-set:concept-input-type-file-selected},
            then [create an
            entry](#create-an-entry){#constructing-form-data-set:create-an-entry-5}
            with `name`{.variable} and a new
            [`File`{#constructing-form-data-set:file-7}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
            object with an empty name,
            [`application/octet-stream`{#constructing-form-data-set:application/octet-stream}](https://www.rfc-editor.org/rfc/rfc2046#section-4.5.1){x-internal="application/octet-stream"}
            as type, and an empty body, and
            [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-5
            x-internal="list-append"} it to `entry list`{.variable}.

        2.  Otherwise, for each file in [selected
            files](input.html#concept-input-type-file-selected){#constructing-form-data-set:concept-input-type-file-selected-2},
            [create an
            entry](#create-an-entry){#constructing-form-data-set:create-an-entry-6}
            with `name`{.variable} and a
            [`File`{#constructing-form-data-set:file-8}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
            object representing the file, and
            [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-6
            x-internal="list-append"} it to `entry list`{.variable}.

    9.  Otherwise, if the `field`{.variable} element is an
        [`input`{#constructing-form-data-set:the-input-element-6}](input.html#the-input-element)
        element whose
        [`type`{#constructing-form-data-set:attr-input-type-6}](input.html#attr-input-type)
        attribute is in the
        [Hidden](input.html#hidden-state-(type=hidden)){#constructing-form-data-set:hidden-state-(type=hidden)}
        state and `name`{.variable} is an [ASCII
        case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#constructing-form-data-set:ascii-case-insensitive
        x-internal="ascii-case-insensitive"} match for
        \"[`_charset_`{#constructing-form-data-set:attr-fe-name-charset}](#attr-fe-name-charset)\":

        1.  Let `charset`{.variable} be the
            [name](https://encoding.spec.whatwg.org/#name){#constructing-form-data-set:encoding-name
            x-internal="encoding-name"} of `encoding`{.variable}.

        2.  [Create an
            entry](#create-an-entry){#constructing-form-data-set:create-an-entry-7}
            with `name`{.variable} and `charset`{.variable}, and
            [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-7
            x-internal="list-append"} it to `entry list`{.variable}.

    10. Otherwise, [create an
        entry](#create-an-entry){#constructing-form-data-set:create-an-entry-8}
        with `name`{.variable} and the
        [value](#concept-fe-value){#constructing-form-data-set:concept-fe-value}
        of the `field`{.variable} element, and
        [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-8
        x-internal="list-append"} it to `entry list`{.variable}.

    11. If the element has a
        [`dirname`{#constructing-form-data-set:attr-fe-dirname}](#attr-fe-dirname)
        attribute, that attribute\'s value is not the empty string, and
        the element is an [auto-directionality form-associated
        element](dom.html#auto-directionality-form-associated-elements){#constructing-form-data-set:auto-directionality-form-associated-elements}:

        1.  Let `dirname`{.variable} be the value of the element\'s
            [`dirname`{#constructing-form-data-set:attr-fe-dirname-2}](#attr-fe-dirname)
            attribute.

        2.  Let `dir`{.variable} be the string \"`ltr`\" if [the
            directionality](dom.html#the-directionality){#constructing-form-data-set:the-directionality}
            of the element is
            \'[ltr](dom.html#concept-ltr){#constructing-form-data-set:concept-ltr}\',
            and \"`rtl`\" otherwise (i.e., when [the
            directionality](dom.html#the-directionality){#constructing-form-data-set:the-directionality-2}
            of the element is
            \'[rtl](dom.html#concept-rtl){#constructing-form-data-set:concept-rtl}\').

        3.  [Create an
            entry](#create-an-entry){#constructing-form-data-set:create-an-entry-9}
            with `dirname`{.variable} and `dir`{.variable}, and
            [append](https://infra.spec.whatwg.org/#list-append){#constructing-form-data-set:list-append-9
            x-internal="list-append"} it to `entry list`{.variable}.

6.  Let `form data`{.variable} be a new
    [`FormData`{#constructing-form-data-set:formdata}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
    object associated with `entry list`{.variable}.

7.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#constructing-form-data-set:concept-event-fire
    x-internal="concept-event-fire"} named
    [`formdata`{#constructing-form-data-set:event-formdata}](indices.html#event-formdata)
    at `form`{.variable} using
    [`FormDataEvent`{#constructing-form-data-set:formdataevent}](#formdataevent),
    with the
    [`formData`{#constructing-form-data-set:dom-formdataevent-formdata}](#dom-formdataevent-formdata)
    attribute initialized to `form data`{.variable} and the
    [`bubbles`{#constructing-form-data-set:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    attribute initialized to true.

8.  Set `form`{.variable}\'s [constructing entry
    list](#constructing-entry-list){#constructing-form-data-set:constructing-entry-list-3}
    to false.

9.  Return a
    [clone](https://infra.spec.whatwg.org/#list-clone){#constructing-form-data-set:list-clone
    x-internal="list-clone"} of `entry list`{.variable}.

##### [4.10.21.5]{.secno} Selecting a form submission encoding[](#selecting-a-form-submission-encoding){.self-link}

If the user agent is to [pick an encoding for a
form]{#picking-an-encoding-for-the-form .dfn}, it must run the following
steps:

1.  Let `encoding`{.variable} be the [document\'s character
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#selecting-a-form-submission-encoding:document's-character-encoding
    x-internal="document's-character-encoding"}.

2.  If the
    [`form`{#selecting-a-form-submission-encoding:the-form-element}](forms.html#the-form-element)
    element has an
    [`accept-charset`{#selecting-a-form-submission-encoding:attr-form-accept-charset}](forms.html#attr-form-accept-charset)
    attribute, set `encoding`{.variable} to the return value of running
    these substeps:

    1.  Let `input`{.variable} be the value of the
        [`form`{#selecting-a-form-submission-encoding:the-form-element-2}](forms.html#the-form-element)
        element\'s
        [`accept-charset`{#selecting-a-form-submission-encoding:attr-form-accept-charset-2}](forms.html#attr-form-accept-charset)
        attribute.

    2.  Let `candidate encoding labels`{.variable} be the result of
        [splitting `input`{.variable} on ASCII
        whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#selecting-a-form-submission-encoding:split-a-string-on-spaces
        x-internal="split-a-string-on-spaces"}.

    3.  Let `candidate encodings`{.variable} be an empty list of
        [character
        encodings](https://encoding.spec.whatwg.org/#encoding){#selecting-a-form-submission-encoding:encoding
        x-internal="encoding"}.

    4.  For each token in `candidate encoding labels`{.variable} in turn
        (in the order in which they were found in `input`{.variable}),
        [get an
        encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#selecting-a-form-submission-encoding:getting-an-encoding
        x-internal="getting-an-encoding"} for the token and, if this
        does not result in failure, append the
        [encoding](https://encoding.spec.whatwg.org/#encoding){#selecting-a-form-submission-encoding:encoding-2
        x-internal="encoding"} to `candidate encodings`{.variable}.

    5.  If `candidate encodings`{.variable} is empty, return
        [UTF-8](https://encoding.spec.whatwg.org/#utf-8){#selecting-a-form-submission-encoding:utf-8
        x-internal="utf-8"}.

    6.  Return the first encoding in `candidate encodings`{.variable}.

3.  Return the result of [getting an output
    encoding](https://encoding.spec.whatwg.org/#get-an-output-encoding){#selecting-a-form-submission-encoding:get-an-output-encoding
    x-internal="get-an-output-encoding"} from `encoding`{.variable}.

##### [4.10.21.6]{.secno} Converting an entry list to a list of name-value pairs[](#converting-an-entry-list-to-a-list-of-name-value-pairs){.self-link}

The
[`application/x-www-form-urlencoded`{#converting-an-entry-list-to-a-list-of-name-value-pairs:application/x-www-form-urlencoded}](https://url.spec.whatwg.org/#concept-urlencoded){x-internal="application/x-www-form-urlencoded"}
and
[`text/plain`{#converting-an-entry-list-to-a-list-of-name-value-pairs:text/plain-encoding-algorithm}](#text/plain-encoding-algorithm)
encoding algorithms take a list of name-value pairs, where the values
must be strings, rather than an [entry
list](#entry-list){#converting-an-entry-list-to-a-list-of-name-value-pairs:entry-list}
where the value can be a
[`File`{#converting-an-entry-list-to-a-list-of-name-value-pairs:file}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}.
The following algorithm performs the conversion.

To [convert to a list of name-value
pairs]{#convert-to-a-list-of-name-value-pairs .dfn} an [entry
list](#entry-list){#converting-an-entry-list-to-a-list-of-name-value-pairs:entry-list-2}
`entry list`{.variable}, run these steps:

1.  Let `list`{.variable} be an empty
    [list](https://infra.spec.whatwg.org/#list){#converting-an-entry-list-to-a-list-of-name-value-pairs:list
    x-internal="list"} of name-value pairs.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#converting-an-entry-list-to-a-list-of-name-value-pairs:list-iterate
    x-internal="list-iterate"} `entry`{.variable} of
    `entry list`{.variable}:

    1.  Let `name`{.variable} be `entry`{.variable}\'s
        [name](#form-entry-name){#converting-an-entry-list-to-a-list-of-name-value-pairs:form-entry-name},
        with every occurrence of U+000D (CR) not followed by U+000A
        (LF), and every occurrence of U+000A (LF) not preceded by U+000D
        (CR), replaced by a string consisting of U+000D (CR) and U+000A
        (LF).

    2.  If `entry`{.variable}\'s
        [value](#form-entry-value){#converting-an-entry-list-to-a-list-of-name-value-pairs:form-entry-value}
        is a
        [`File`{#converting-an-entry-list-to-a-list-of-name-value-pairs:file-2}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
        object, then let `value`{.variable} be `entry`{.variable}\'s
        [value](#form-entry-value){#converting-an-entry-list-to-a-list-of-name-value-pairs:form-entry-value-2}\'s
        [`name`{#converting-an-entry-list-to-a-list-of-name-value-pairs:dom-file-name}](https://w3c.github.io/FileAPI/#dfn-name){x-internal="dom-file-name"}.
        Otherwise, let `value`{.variable} be `entry`{.variable}\'s
        [value](#form-entry-value){#converting-an-entry-list-to-a-list-of-name-value-pairs:form-entry-value-3}.

    3.  Replace every occurrence of U+000D (CR) not followed by U+000A
        (LF), and every occurrence of U+000A (LF) not preceded by U+000D
        (CR), in `value`{.variable}, by a string consisting of U+000D
        (CR) and U+000A (LF).

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#converting-an-entry-list-to-a-list-of-name-value-pairs:list-append
        x-internal="list-append"} to `list`{.variable} a new name-value
        pair whose name is `name`{.variable} and whose value is
        `value`{.variable}.

3.  Return `list`{.variable}.

##### [4.10.21.7]{.secno} URL-encoded form data[](#url-encoded-form-data){.self-link}

[]{#application/x-www-form-urlencoded-encoding-algorithm}See URL for
details on
[`application/x-www-form-urlencoded`{#url-encoded-form-data:application/x-www-form-urlencoded}](https://url.spec.whatwg.org/#concept-urlencoded){x-internal="application/x-www-form-urlencoded"}.
[\[URL\]](references.html#refsURL)

##### [4.10.21.8]{.secno} Multipart form data[](#multipart-form-data){.self-link}

The [`multipart/form-data` encoding
algorithm]{#multipart/form-data-encoding-algorithm .dfn export=""},
given an [entry list](#entry-list){#multipart-form-data:entry-list}
`entry list`{.variable} and an
[encoding](https://encoding.spec.whatwg.org/#encoding){#multipart-form-data:encoding
x-internal="encoding"} `encoding`{.variable}, is as follows:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#multipart-form-data:list-iterate
    x-internal="list-iterate"} `entry`{.variable} of
    `entry list`{.variable}:

    1.  Replace every occurrence of U+000D (CR) not followed by U+000A
        (LF), and every occurrence of U+000A (LF) not preceded by U+000D
        (CR), in `entry`{.variable}\'s
        [name](#form-entry-name){#multipart-form-data:form-entry-name},
        by a string consisting of a U+000D (CR) and U+000A (LF).

    2.  If `entry`{.variable}\'s
        [value](#form-entry-value){#multipart-form-data:form-entry-value}
        is not a
        [`File`{#multipart-form-data:file}](https://w3c.github.io/FileAPI/#dfn-file){x-internal="file"}
        object, then replace every occurrence of U+000D (CR) not
        followed by U+000A (LF), and every occurrence of U+000A (LF) not
        preceded by U+000D (CR), in `entry`{.variable}\'s
        [value](#form-entry-value){#multipart-form-data:form-entry-value-2},
        by a string consisting of a U+000D (CR) and U+000A (LF).

2.  Return the byte sequence resulting from encoding the
    `entry list`{.variable} using the rules described by RFC 7578,
    Returning Values from Forms: `multipart/form-data`, given the
    following conditions: [\[RFC7578\]](references.html#refsRFC7578)

    - Each [entry](#form-entry){#multipart-form-data:form-entry} in
      `entry list`{.variable} is a *field*, the
      [name](#form-entry-name){#multipart-form-data:form-entry-name-2}
      of the entry is the *field name* and the
      [value](#form-entry-value){#multipart-form-data:form-entry-value-3}
      of the entry is the *field value*.

    - The order of parts must be the same as the order of fields in
      `entry list`{.variable}. Multiple entries with the same name must
      be treated as distinct fields.

    - Field names, field values for non-file fields, and filenames for
      file fields, in the generated
      [`multipart/form-data`{#multipart-form-data:multipart/form-data}](indices.html#multipart/form-data)
      resource must be set to the result of
      [encoding](https://encoding.spec.whatwg.org/#encode){#multipart-form-data:encode
      x-internal="encode"} the corresponding entry\'s name or value with
      `encoding`{.variable}, converted to a byte sequence.

    - For field names and filenames for file fields, the result of the
      encoding in the previous bullet point must be escaped by replacing
      any 0x0A (LF) bytes with the byte sequence \``%0A`\`, 0x0D (CR)
      with \``%0D`\` and 0x22 (\") with \``%22`\`. The user agent must
      not perform any other escapes.

    - The parts of the generated
      [`multipart/form-data`{#multipart-form-data:multipart/form-data-2}](indices.html#multipart/form-data)
      resource that correspond to non-file fields must not have a
      \`[`Content-Type`{#multipart-form-data:content-type}](urls-and-fetching.html#content-type)\`
      header specified.

    - The boundary used by the user agent in generating the return value
      of this algorithm is the [`multipart/form-data` boundary
      string]{#multipart/form-data-boundary-string .dfn export=""}.
      (This value is used to generate the MIME type of the form
      submission payload generated by this algorithm.)

For details on how to interpret
[`multipart/form-data`{#multipart-form-data:multipart/form-data-3}](indices.html#multipart/form-data)
payloads, see RFC 7578. [\[RFC7578\]](references.html#refsRFC7578)

##### [4.10.21.9]{.secno} Plain text form data[](#plain-text-form-data){.self-link}

The [`text/plain` encoding algorithm]{#text/plain-encoding-algorithm
.dfn}, given a list of name-value pairs `pairs`{.variable}, is as
follows:

1.  Let `result`{.variable} be the empty string.

2.  For each `pair`{.variable} in `pairs`{.variable}:

    1.  Append `pair`{.variable}\'s name to `result`{.variable}.

    2.  Append a single U+003D EQUALS SIGN character (=) to
        `result`{.variable}.

    3.  Append `pair`{.variable}\'s value to `result`{.variable}.

    4.  Append a U+000D CARRIAGE RETURN (CR) U+000A LINE FEED (LF)
        character pair to `result`{.variable}.

3.  Return `result`{.variable}.

Payloads using the
[`text/plain`{#plain-text-form-data:text/plain}](https://www.rfc-editor.org/rfc/rfc2046#section-4.1.3){x-internal="text/plain"}
format are intended to be human readable. They are not reliably
interpretable by computer, as the format is ambiguous (for example,
there is no way to distinguish a literal newline in a value from the
newline at the end of the value).

##### [4.10.21.10]{.secno} The [`SubmitEvent`{#the-submitevent-interface:submitevent}](#submitevent) interface[](#the-submitevent-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[SubmitEvent](https://developer.mozilla.org/en-US/docs/Web/API/SubmitEvent "The SubmitEvent interface defines the object used to represent an HTML form's submit event. This event is fired at the <form> when the form's submit action is invoked.")

Support in all current engines.

::: support
[Firefox75+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome81+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge81+]{.edge_blink .yes}

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
[SubmitEvent/SubmitEvent](https://developer.mozilla.org/en-US/docs/Web/API/SubmitEvent/SubmitEvent "The SubmitEvent() constructor creates and returns a new SubmitEvent object, which is used to represent a submit event fired at an HTML form.")

Support in all current engines.

::: support
[Firefox75+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome81+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge81+]{.edge_blink .yes}

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

``` idl
[Exposed=Window]
interface SubmitEvent : Event {
  constructor(DOMString type, optional SubmitEventInit eventInitDict = {});

  readonly attribute HTMLElement? submitter;
};

dictionary SubmitEventInit : EventInit {
  HTMLElement? submitter = null;
};
```

`event`{.variable}`.`[`submitter`](#dom-submitevent-submitter){#the-submitevent-interface:dom-submitevent-submitter-2}

:   Returns the element representing the [submit
    button](forms.html#concept-submit-button){#the-submitevent-interface:concept-submit-button}
    that triggered the [form
    submission](#form-submission-2){#the-submitevent-interface:form-submission-2},
    or null if the submission was not triggered by a button.

The [`submitter`]{#dom-submitevent-submitter .dfn dfn-for="SubmitEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to.

##### [4.10.21.11]{.secno} The [`FormDataEvent`{#the-formdataevent-interface:formdataevent}](#formdataevent) interface[](#the-formdataevent-interface){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[FormDataEvent/FormDataEvent](https://developer.mozilla.org/en-US/docs/Web/API/FormDataEvent/FormDataEvent "The FormDataEvent() constructor creates a new FormDataEvent object.")

Support in all current engines.

::: support
[Firefox72+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome77+]{.chrome
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

:::: feature
[FormDataEvent](https://developer.mozilla.org/en-US/docs/Web/API/FormDataEvent "The FormDataEvent interface represents a formdata event — such an event is fired on an HTMLFormElement object after the entry list representing the form's data is constructed. This happens when the form is submitted, but can also be triggered by the invocation of a FormData() constructor.")

Support in all current engines.

::: support
[Firefox72+]{.firefox .yes}[Safari15+]{.safari .yes}[Chrome77+]{.chrome
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

``` idl
[Exposed=Window]
interface FormDataEvent : Event {
  constructor(DOMString type, FormDataEventInit eventInitDict);

  readonly attribute FormData formData;
};

dictionary FormDataEventInit : EventInit {
  required FormData formData;
};
```

`event`{.variable}`.`[`formData`](#dom-formdataevent-formdata){#the-formdataevent-interface:dom-formdataevent-formdata-2}

:   Returns a
    [`FormData`{#the-formdataevent-interface:formdata-3}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
    object representing names and values of elements associated to the
    target
    [`form`{#the-formdataevent-interface:the-form-element}](forms.html#the-form-element).
    Operations on the
    [`FormData`{#the-formdataevent-interface:formdata-4}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
    object will affect form data to be submitted.

The [`formData`]{#dom-formdataevent-formdata .dfn
dfn-for="FormDataEvent" dfn-type="attribute"} attribute must return the
value it was initialized to. It represents a
[`FormData`{#the-formdataevent-interface:formdata-5}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
object associated to the [entry
list](#entry-list){#the-formdataevent-interface:entry-list} that is
[constructed](#constructing-the-form-data-set){#the-formdataevent-interface:constructing-the-form-data-set}
when the
[`form`{#the-formdataevent-interface:the-form-element-2}](forms.html#the-form-element)
is submitted.

#### [4.10.22]{.secno} Resetting a form[](#resetting-a-form){.self-link}

When a
[`form`{#resetting-a-form:the-form-element}](forms.html#the-form-element)
element `form`{.variable} is [reset]{#concept-form-reset .dfn}, run
these steps:

1.  Let `reset`{.variable} be the result of [firing an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#resetting-a-form:concept-event-fire
    x-internal="concept-event-fire"} named
    [`reset`{#resetting-a-form:event-reset}](indices.html#event-reset)
    at `form`{.variable}, with the
    [`bubbles`{#resetting-a-form:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    and
    [`cancelable`{#resetting-a-form:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    attributes initialized to true.

2.  If `reset`{.variable} is true, then invoke the [reset
    algorithm](#concept-form-reset-control){#resetting-a-form:concept-form-reset-control}
    of each [resettable
    element](forms.html#category-reset){#resetting-a-form:category-reset}
    whose [form owner](#form-owner){#resetting-a-form:form-owner} is
    `form`{.variable}.

Each [resettable
element](forms.html#category-reset){#resetting-a-form:category-reset-2}
defines its own [reset algorithm]{#concept-form-reset-control .dfn}.
Changes made to form controls as part of these algorithms do not count
as changes caused by the user (and thus, e.g., do not cause
[`input`{#resetting-a-form:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
events to fire).

[← 4.10.6 The button element](form-elements.html) --- [Table of
Contents](./) --- [4.11 Interactive elements
→](interactive-elements.html)
