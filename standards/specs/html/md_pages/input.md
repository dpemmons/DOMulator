HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.10 Forms](forms.html) --- [Table of Contents](./) --- [4.10.6 The
button element →](form-elements.html)

1.  ::: {#toc-semantics}
    1.  1.  [[4.10.5]{.secno} The `input`
            element](input.html#the-input-element)
            1.  [[4.10.5.1]{.secno} States of the `type`
                attribute](input.html#states-of-the-type-attribute)
                1.  [[4.10.5.1.1]{.secno} Hidden state
                    (`type=hidden`)](input.html#hidden-state-(type=hidden))
                2.  [[4.10.5.1.2]{.secno} Text (`type=text`) state and
                    Search state
                    (`type=search`)](input.html#text-(type=text)-state-and-search-state-(type=search))
                3.  [[4.10.5.1.3]{.secno} Telephone state
                    (`type=tel`)](input.html#telephone-state-(type=tel))
                4.  [[4.10.5.1.4]{.secno} URL state
                    (`type=url`)](input.html#url-state-(type=url))
                5.  [[4.10.5.1.5]{.secno} Email state
                    (`type=email`)](input.html#email-state-(type=email))
                6.  [[4.10.5.1.6]{.secno} Password state
                    (`type=password`)](input.html#password-state-(type=password))
                7.  [[4.10.5.1.7]{.secno} Date state
                    (`type=date`)](input.html#date-state-(type=date))
                8.  [[4.10.5.1.8]{.secno} Month state
                    (`type=month`)](input.html#month-state-(type=month))
                9.  [[4.10.5.1.9]{.secno} Week state
                    (`type=week`)](input.html#week-state-(type=week))
                10. [[4.10.5.1.10]{.secno} Time state
                    (`type=time`)](input.html#time-state-(type=time))
                11. [[4.10.5.1.11]{.secno} Local Date and Time state
                    (`type=datetime-local`)](input.html#local-date-and-time-state-(type=datetime-local))
                12. [[4.10.5.1.12]{.secno} Number state
                    (`type=number`)](input.html#number-state-(type=number))
                13. [[4.10.5.1.13]{.secno} Range state
                    (`type=range`)](input.html#range-state-(type=range))
                14. [[4.10.5.1.14]{.secno} Color state
                    (`type=color`)](input.html#color-state-(type=color))
                15. [[4.10.5.1.15]{.secno} Checkbox state
                    (`type=checkbox`)](input.html#checkbox-state-(type=checkbox))
                16. [[4.10.5.1.16]{.secno} Radio Button state
                    (`type=radio`)](input.html#radio-button-state-(type=radio))
                17. [[4.10.5.1.17]{.secno} File Upload state
                    (`type=file`)](input.html#file-upload-state-(type=file))
                18. [[4.10.5.1.18]{.secno} Submit Button state
                    (`type=submit`)](input.html#submit-button-state-(type=submit))
                19. [[4.10.5.1.19]{.secno} Image Button state
                    (`type=image`)](input.html#image-button-state-(type=image))
                20. [[4.10.5.1.20]{.secno} Reset Button state
                    (`type=reset`)](input.html#reset-button-state-(type=reset))
                21. [[4.10.5.1.21]{.secno} Button state
                    (`type=button`)](input.html#button-state-(type=button))
            2.  [[4.10.5.2]{.secno} Implementation notes regarding
                localization of form
                controls](input.html#input-impl-notes)
            3.  [[4.10.5.3]{.secno} Common `input` element
                attributes](input.html#common-input-element-attributes)
                1.  [[4.10.5.3.1]{.secno} The `maxlength` and
                    `minlength`
                    attributes](input.html#the-maxlength-and-minlength-attributes)
                2.  [[4.10.5.3.2]{.secno} The `size`
                    attribute](input.html#the-size-attribute)
                3.  [[4.10.5.3.3]{.secno} The `readonly`
                    attribute](input.html#the-readonly-attribute)
                4.  [[4.10.5.3.4]{.secno} The `required`
                    attribute](input.html#the-required-attribute)
                5.  [[4.10.5.3.5]{.secno} The `multiple`
                    attribute](input.html#the-multiple-attribute)
                6.  [[4.10.5.3.6]{.secno} The `pattern`
                    attribute](input.html#the-pattern-attribute)
                7.  [[4.10.5.3.7]{.secno} The `min` and `max`
                    attributes](input.html#the-min-and-max-attributes)
                8.  [[4.10.5.3.8]{.secno} The `step`
                    attribute](input.html#the-step-attribute)
                9.  [[4.10.5.3.9]{.secno} The `list`
                    attribute](input.html#the-list-attribute)
                10. [[4.10.5.3.10]{.secno} The `placeholder`
                    attribute](input.html#the-placeholder-attribute)
            4.  [[4.10.5.4]{.secno} Common `input` element
                APIs](input.html#common-input-element-apis)
            5.  [[4.10.5.5]{.secno} Common event
                behaviors](input.html#common-input-element-events)
    :::

#### [4.10.5]{.secno} The [`input`]{.dfn dfn-type="element"} element[](#the-input-element){.self-link}

:::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android1+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

::: feature
[Element/input](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")
:::
::::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement "The HTMLInputElement interface provides special properties and methods for manipulating the options, layout, and presentation of <input> elements.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera8+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-input-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-input-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-input-element:phrasing-content-2}.
:   If the
    [`type`{#the-input-element:attr-input-type}](#attr-input-type)
    attribute is *not* in the
    [Hidden](#hidden-state-(type=hidden)){#the-input-element:hidden-state-(type=hidden)}
    state: [Interactive
    content](dom.html#interactive-content-2){#the-input-element:interactive-content-2}.
:   If the
    [`type`{#the-input-element:attr-input-type-2}](#attr-input-type)
    attribute is *not* in the
    [Hidden](#hidden-state-(type=hidden)){#the-input-element:hidden-state-(type=hidden)-2}
    state:
    [Listed](forms.html#category-listed){#the-input-element:category-listed},
    [labelable](forms.html#category-label){#the-input-element:category-label},
    [submittable](forms.html#category-submit){#the-input-element:category-submit},
    [resettable](forms.html#category-reset){#the-input-element:category-reset},
    and [autocapitalize-and-autocorrect
    inheriting](forms.html#category-autocapitalize){#the-input-element:category-autocapitalize}
    [form-associated
    element](forms.html#form-associated-element){#the-input-element:form-associated-element}.
:   If the
    [`type`{#the-input-element:attr-input-type-3}](#attr-input-type)
    attribute is in the
    [Hidden](#hidden-state-(type=hidden)){#the-input-element:hidden-state-(type=hidden)-3}
    state:
    [Listed](forms.html#category-listed){#the-input-element:category-listed-2},
    [submittable](forms.html#category-submit){#the-input-element:category-submit-2},
    [resettable](forms.html#category-reset){#the-input-element:category-reset-2},
    and [autocapitalize-and-autocorrect
    inheriting](forms.html#category-autocapitalize){#the-input-element:category-autocapitalize-2}
    [form-associated
    element](forms.html#form-associated-element){#the-input-element:form-associated-element-2}.
:   If the
    [`type`{#the-input-element:attr-input-type-4}](#attr-input-type)
    attribute is *not* in the
    [Hidden](#hidden-state-(type=hidden)){#the-input-element:hidden-state-(type=hidden)-4}
    state: [Palpable
    content](dom.html#palpable-content-2){#the-input-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-input-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-input-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-input-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-input-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-input-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-input-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-input-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-input-element:global-attributes}
:   [`accept`{#the-input-element:attr-input-accept}](#attr-input-accept)
    --- Hint for expected file type in [file upload
    controls](#file-upload-state-(type=file)){#the-input-element:file-upload-state-(type=file)}
:   [`alpha`{#the-input-element:attr-input-alpha}](#attr-input-alpha)
    --- Allow the color\'s alpha component to be set
:   [`alt`{#the-input-element:attr-input-alt}](#attr-input-alt) ---
    Replacement text for use when images are not available
:   [`autocomplete`{#the-input-element:attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete)
    --- Hint for form autofill feature
:   [`checked`{#the-input-element:attr-input-checked}](#attr-input-checked)
    --- Whether the control is checked
:   [`colorspace`{#the-input-element:attr-input-colorspace}](#attr-input-colorspace)
    --- The color space of the serialized color
:   [`dirname`{#the-input-element:attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname)
    --- Name of form control to use for sending the element\'s
    [directionality](dom.html#the-directionality){#the-input-element:the-directionality}
    in [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2}
:   [`disabled`{#the-input-element:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
    --- Whether the form control is disabled
:   [`form`{#the-input-element:attr-fae-form}](form-control-infrastructure.html#attr-fae-form)
    --- Associates the element with a
    [`form`{#the-input-element:the-form-element}](forms.html#the-form-element)
    element
:   [`formaction`{#the-input-element:attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction)
    ---
    [URL](https://url.spec.whatwg.org/#concept-url){#the-input-element:url
    x-internal="url"} to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2-2}
:   [`formenctype`{#the-input-element:attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype)
    --- [Entry
    list](form-control-infrastructure.html#entry-list){#the-input-element:entry-list}
    encoding type to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2-3}
:   [`formmethod`{#the-input-element:attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod)
    --- Variant to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2-4}
:   [`formnovalidate`{#the-input-element:attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate)
    --- Bypass form control validation for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2-5}
:   [`formtarget`{#the-input-element:attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget)
    ---
    [Navigable](document-sequences.html#navigable){#the-input-element:navigable}
    for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2-6}
:   [`height`{#the-input-element:attr-dim-height}](embedded-content-other.html#attr-dim-height)
    --- Vertical dimension
:   [`list`{#the-input-element:attr-input-list}](#attr-input-list) ---
    List of autocomplete options
:   [`max`{#the-input-element:attr-input-max}](#attr-input-max) ---
    Maximum value
:   [`maxlength`{#the-input-element:attr-input-maxlength}](#attr-input-maxlength)
    --- Maximum
    [length](https://infra.spec.whatwg.org/#string-length){#the-input-element:length
    x-internal="length"} of value
:   [`min`{#the-input-element:attr-input-min}](#attr-input-min) ---
    Minimum value
:   [`minlength`{#the-input-element:attr-input-minlength}](#attr-input-minlength)
    --- Minimum
    [length](https://infra.spec.whatwg.org/#string-length){#the-input-element:length-2
    x-internal="length"} of value
:   [`multiple`{#the-input-element:attr-input-multiple}](#attr-input-multiple)
    --- Whether to allow multiple values
:   [`name`{#the-input-element:attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
    --- Name of the element to use for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2-7}
    and in the
    [`form.elements`{#the-input-element:dom-form-elements}](forms.html#dom-form-elements)
    API
:   [`pattern`{#the-input-element:attr-input-pattern}](#attr-input-pattern)
    --- Pattern to be matched by the form control\'s value
:   [`placeholder`{#the-input-element:attr-input-placeholder}](#attr-input-placeholder)
    --- User-visible label to be placed within the form control
:   [`popovertarget`{#the-input-element:attr-popovertarget}](popover.html#attr-popovertarget)
    --- Targets a popover element to toggle, show, or hide
:   [`popovertargetaction`{#the-input-element:attr-popovertargetaction}](popover.html#attr-popovertargetaction)
    --- Indicates whether a targeted popover element is to be toggled,
    shown, or hidden
:   [`readonly`{#the-input-element:attr-input-readonly}](#attr-input-readonly)
    --- Whether to allow the value to be edited by the user
:   [`required`{#the-input-element:attr-input-required}](#attr-input-required)
    --- Whether the control is required for [form
    submission](form-control-infrastructure.html#form-submission-2){#the-input-element:form-submission-2-8}
:   [`size`{#the-input-element:attr-input-size}](#attr-input-size) ---
    Size of the control
:   [`src`{#the-input-element:attr-input-src}](#attr-input-src) ---
    Address of the resource
:   [`step`{#the-input-element:attr-input-step}](#attr-input-step) ---
    Granularity to be matched by the form control\'s value
:   [`type`{#the-input-element:attr-input-type-5}](#attr-input-type) ---
    Type of form control
:   [`value`{#the-input-element:attr-input-value}](#attr-input-value)
    --- Value of the form control
:   [`width`{#the-input-element:attr-dim-width}](embedded-content-other.html#attr-dim-width)
    --- Horizontal dimension
:   Also, the
    [`title`{#the-input-element:attr-input-title}](#attr-input-title)
    attribute [has special
    semantics](#attr-input-title){#the-input-element:attr-input-title-2}
    on this element: Description of pattern (when used with
    [`pattern`{#the-input-element:attr-input-pattern-2}](#attr-input-pattern)
    attribute)

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-input-element:concept-element-accessibility-considerations}:
:   [`type`{#the-input-element:attr-input-type-6}](#attr-input-type)
    attribute in the
    [Hidden](#hidden-state-(type=hidden)){#the-input-element:hidden-state-(type=hidden)-5}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-hidden); [for
    implementers](https://w3c.github.io/html-aam/#el-input-hidden).
:   [`type`{#the-input-element:attr-input-type-7}](#attr-input-type)
    attribute in the
    [Text](#text-(type=text)-state-and-search-state-(type=search)){#the-input-element:text-(type=text)-state-and-search-state-(type=search)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-text); [for
    implementers](https://w3c.github.io/html-aam/#el-input-text).
:   [`type`{#the-input-element:attr-input-type-8}](#attr-input-type)
    attribute in the
    [Search](#text-(type=text)-state-and-search-state-(type=search)){#the-input-element:text-(type=text)-state-and-search-state-(type=search)-2}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-search); [for
    implementers](https://w3c.github.io/html-aam/#el-input-search).
:   [`type`{#the-input-element:attr-input-type-9}](#attr-input-type)
    attribute in the
    [Telephone](#telephone-state-(type=tel)){#the-input-element:telephone-state-(type=tel)}
    state: [for authors](https://w3c.github.io/html-aria/#el-input-tel);
    [for implementers](https://w3c.github.io/html-aam/#el-input-tel).
:   [`type`{#the-input-element:attr-input-type-10}](#attr-input-type)
    attribute in the
    [URL](#url-state-(type=url)){#the-input-element:url-state-(type=url)}
    state: [for authors](https://w3c.github.io/html-aria/#el-input-url);
    [for implementers](https://w3c.github.io/html-aam/#el-input-url).
:   [`type`{#the-input-element:attr-input-type-11}](#attr-input-type)
    attribute in the
    [Email](#email-state-(type=email)){#the-input-element:email-state-(type=email)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-email); [for
    implementers](https://w3c.github.io/html-aam/#el-input-email).
:   [`type`{#the-input-element:attr-input-type-12}](#attr-input-type)
    attribute in the
    [Password](#password-state-(type=password)){#the-input-element:password-state-(type=password)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-password); [for
    implementers](https://w3c.github.io/html-aam/#el-input-password).
:   [`type`{#the-input-element:attr-input-type-13}](#attr-input-type)
    attribute in the
    [Date](#date-state-(type=date)){#the-input-element:date-state-(type=date)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-date); [for
    implementers](https://w3c.github.io/html-aam/#el-input-date).
:   [`type`{#the-input-element:attr-input-type-14}](#attr-input-type)
    attribute in the
    [Month](#month-state-(type=month)){#the-input-element:month-state-(type=month)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-month); [for
    implementers](https://w3c.github.io/html-aam/#el-input-month).
:   [`type`{#the-input-element:attr-input-type-15}](#attr-input-type)
    attribute in the
    [Week](#week-state-(type=week)){#the-input-element:week-state-(type=week)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-week); [for
    implementers](https://w3c.github.io/html-aam/#el-input-week).
:   [`type`{#the-input-element:attr-input-type-16}](#attr-input-type)
    attribute in the
    [Time](#time-state-(type=time)){#the-input-element:time-state-(type=time)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-time); [for
    implementers](https://w3c.github.io/html-aam/#el-input-time).
:   [`type`{#the-input-element:attr-input-type-17}](#attr-input-type)
    attribute in the [Local Date and
    Time](#local-date-and-time-state-(type=datetime-local)){#the-input-element:local-date-and-time-state-(type=datetime-local)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-datetime-local);
    [for
    implementers](https://w3c.github.io/html-aam/#el-input-datetime-local).
:   [`type`{#the-input-element:attr-input-type-18}](#attr-input-type)
    attribute in the
    [Number](#number-state-(type=number)){#the-input-element:number-state-(type=number)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-number); [for
    implementers](https://w3c.github.io/html-aam/#el-input-number).
:   [`type`{#the-input-element:attr-input-type-19}](#attr-input-type)
    attribute in the
    [Range](#range-state-(type=range)){#the-input-element:range-state-(type=range)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-range); [for
    implementers](https://w3c.github.io/html-aam/#el-input-range).
:   [`type`{#the-input-element:attr-input-type-20}](#attr-input-type)
    attribute in the
    [Color](#color-state-(type=color)){#the-input-element:color-state-(type=color)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-color); [for
    implementers](https://w3c.github.io/html-aam/#el-input-color).
:   [`type`{#the-input-element:attr-input-type-21}](#attr-input-type)
    attribute in the
    [Checkbox](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-checkbox); [for
    implementers](https://w3c.github.io/html-aam/#el-input-checkbox).
:   [`type`{#the-input-element:attr-input-type-22}](#attr-input-type)
    attribute in the [Radio
    Button](#radio-button-state-(type=radio)){#the-input-element:radio-button-state-(type=radio)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-radio); [for
    implementers](https://w3c.github.io/html-aam/#el-input-radio).
:   [`type`{#the-input-element:attr-input-type-23}](#attr-input-type)
    attribute in the [File
    Upload](#file-upload-state-(type=file)){#the-input-element:file-upload-state-(type=file)-2}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-file); [for
    implementers](https://w3c.github.io/html-aam/#el-input-file).
:   [`type`{#the-input-element:attr-input-type-24}](#attr-input-type)
    attribute in the [Submit
    Button](#submit-button-state-(type=submit)){#the-input-element:submit-button-state-(type=submit)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-submit); [for
    implementers](https://w3c.github.io/html-aam/#el-input-submit).
:   [`type`{#the-input-element:attr-input-type-25}](#attr-input-type)
    attribute in the [Image
    Button](#image-button-state-(type=image)){#the-input-element:image-button-state-(type=image)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-image); [for
    implementers](https://w3c.github.io/html-aam/#el-input-image).
:   [`type`{#the-input-element:attr-input-type-26}](#attr-input-type)
    attribute in the [Reset
    Button](#reset-button-state-(type=reset)){#the-input-element:reset-button-state-(type=reset)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-reset); [for
    implementers](https://w3c.github.io/html-aam/#el-input-reset).
:   [`type`{#the-input-element:attr-input-type-27}](#attr-input-type)
    attribute in the
    [Button](#button-state-(type=button)){#the-input-element:button-state-(type=button)}
    state: [for
    authors](https://w3c.github.io/html-aria/#el-input-button); [for
    implementers](https://w3c.github.io/html-aam/#el-input-button).

[DOM interface](dom.html#concept-element-dom){#the-input-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLInputElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString accept;
      [CEReactions] attribute boolean alpha;
      [CEReactions] attribute DOMString alt;
      [CEReactions] attribute DOMString autocomplete;
      [CEReactions] attribute boolean defaultChecked;
      attribute boolean checked;
      [CEReactions] attribute DOMString colorSpace;
      [CEReactions] attribute DOMString dirName;
      [CEReactions] attribute boolean disabled;
      readonly attribute HTMLFormElement? form;
      attribute FileList? files;
      [CEReactions] attribute USVString formAction;
      [CEReactions] attribute DOMString formEnctype;
      [CEReactions] attribute DOMString formMethod;
      [CEReactions] attribute boolean formNoValidate;
      [CEReactions] attribute DOMString formTarget;
      [CEReactions] attribute unsigned long height;
      attribute boolean indeterminate;
      readonly attribute HTMLDataListElement? list;
      [CEReactions] attribute DOMString max;
      [CEReactions] attribute long maxLength;
      [CEReactions] attribute DOMString min;
      [CEReactions] attribute long minLength;
      [CEReactions] attribute boolean multiple;
      [CEReactions] attribute DOMString name;
      [CEReactions] attribute DOMString pattern;
      [CEReactions] attribute DOMString placeholder;
      [CEReactions] attribute boolean readOnly;
      [CEReactions] attribute boolean required;
      [CEReactions] attribute unsigned long size;
      [CEReactions] attribute USVString src;
      [CEReactions] attribute DOMString step;
      [CEReactions] attribute DOMString type;
      [CEReactions] attribute DOMString defaultValue;
      [CEReactions] attribute [LegacyNullToEmptyString] DOMString value;
      attribute object? valueAsDate;
      attribute unrestricted double valueAsNumber;
      [CEReactions] attribute unsigned long width;

      undefined stepUp(optional long n = 1);
      undefined stepDown(optional long n = 1);

      readonly attribute boolean willValidate;
      readonly attribute ValidityState validity;
      readonly attribute DOMString validationMessage;
      boolean checkValidity();
      boolean reportValidity();
      undefined setCustomValidity(DOMString error);

      readonly attribute NodeList? labels;

      undefined select();
      attribute unsigned long? selectionStart;
      attribute unsigned long? selectionEnd;
      attribute DOMString? selectionDirection;
      undefined setRangeText(DOMString replacement);
      undefined setRangeText(DOMString replacement, unsigned long start, unsigned long end, optional SelectionMode selectionMode = "preserve");
      undefined setSelectionRange(unsigned long start, unsigned long end, optional DOMString direction);

      undefined showPicker();

      // also has obsolete members
    };
    HTMLInputElement includes PopoverInvokerElement;
    ```

The [`input`{#the-input-element:the-input-element}](#the-input-element)
element [represents](dom.html#represents){#the-input-element:represents}
a typed data field, usually with a form control to allow the user to
edit the data.

The [`type`]{#attr-input-type .dfn dfn-for="input"
dfn-type="element-attr"} attribute controls the data type (and
associated control) of the element. It is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-input-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Data type

Control type

[`hidden`]{#attr-input-type-hidden-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Hidden](#hidden-state-(type=hidden)){#the-input-element:hidden-state-(type=hidden)-6}

An arbitrary string

n/a

[`text`]{#attr-input-type-text-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Text](#text-(type=text)-state-and-search-state-(type=search)){#the-input-element:text-(type=text)-state-and-search-state-(type=search)-3}

Text with no line breaks

A text control

[`search`]{#attr-input-type-search-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Search](#text-(type=text)-state-and-search-state-(type=search)){#the-input-element:text-(type=text)-state-and-search-state-(type=search)-4}

Text with no line breaks

Search control

[`tel`]{#attr-input-type-tel-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Telephone](#telephone-state-(type=tel)){#the-input-element:telephone-state-(type=tel)-2}

Text with no line breaks

A text control

[`url`]{#attr-input-type-url-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[URL](#url-state-(type=url)){#the-input-element:url-state-(type=url)-2}

An absolute URL

A text control

[`email`]{#attr-input-type-email-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Email](#email-state-(type=email)){#the-input-element:email-state-(type=email)-2}

An email address or list of email addresses

A text control

[`password`]{#attr-input-type-password-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Password](#password-state-(type=password)){#the-input-element:password-state-(type=password)-2}

Text with no line breaks (sensitive information)

A text control that obscures data entry

[`date`]{#attr-input-type-date-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Date](#date-state-(type=date)){#the-input-element:date-state-(type=date)-2}

A date (year, month, day) with no time zone

A date control

[`month`]{#attr-input-type-month-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Month](#month-state-(type=month)){#the-input-element:month-state-(type=month)-2}

A date consisting of a year and a month with no time zone

A month control

[`week`]{#attr-input-type-week-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Week](#week-state-(type=week)){#the-input-element:week-state-(type=week)-2}

A date consisting of a week-year number and a week number with no time
zone

A week control

[`time`]{#attr-input-type-time-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Time](#time-state-(type=time)){#the-input-element:time-state-(type=time)-2}

A time (hour, minute, seconds, fractional seconds) with no time zone

A time control

[`datetime-local`]{#attr-input-type-datetime-local-keyword .dfn
dfn-for="input/type" dfn-type="attr-value"}

[Local Date and
Time](#local-date-and-time-state-(type=datetime-local)){#the-input-element:local-date-and-time-state-(type=datetime-local)-2}

A date and time (year, month, day, hour, minute, second, fraction of a
second) with no time zone

A date and time control

[`number`]{#attr-input-type-number-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Number](#number-state-(type=number)){#the-input-element:number-state-(type=number)-2}

A numerical value

A text control or spinner control

[`range`]{#attr-input-type-range-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Range](#range-state-(type=range)){#the-input-element:range-state-(type=range)-2}

A numerical value, with the extra semantic that the exact value is not
important

A slider control or similar

[`color`]{#attr-input-type-color-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Color](#color-state-(type=color)){#the-input-element:color-state-(type=color)-2}

An sRGB color with 8-bit red, green, and blue components

A color picker

[`checkbox`]{#attr-input-type-checkbox-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Checkbox](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)-2}

A set of zero or more values from a predefined list

A checkbox

[`radio`]{#attr-input-type-radio-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Radio
Button](#radio-button-state-(type=radio)){#the-input-element:radio-button-state-(type=radio)-2}

An enumerated value

A radio button

[`file`]{#attr-input-type-file-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[File
Upload](#file-upload-state-(type=file)){#the-input-element:file-upload-state-(type=file)-3}

Zero or more files each with a [MIME
type](https://mimesniff.spec.whatwg.org/#mime-type){#the-input-element:mime-type
x-internal="mime-type"} and optionally a filename

A label and a button

[`submit`]{#attr-input-type-submit-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Submit
Button](#submit-button-state-(type=submit)){#the-input-element:submit-button-state-(type=submit)-2}

An enumerated value, with the extra semantic that it must be the last
value selected and initiates form submission

A button

[`image`]{#attr-input-type-image-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Image
Button](#image-button-state-(type=image)){#the-input-element:image-button-state-(type=image)-2}

A coordinate, relative to a particular image\'s size, with the extra
semantic that it must be the last value selected and initiates form
submission

Either a clickable image, or a button

[`reset`]{#attr-input-type-reset-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Reset
Button](#reset-button-state-(type=reset)){#the-input-element:reset-button-state-(type=reset)-2}

n/a

A button

[`button`]{#attr-input-type-button-keyword .dfn dfn-for="input/type"
dfn-type="attr-value"}

[Button](#button-state-(type=button)){#the-input-element:button-state-(type=button)-2}

n/a

A button

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the
[Text](#text-(type=text)-state-and-search-state-(type=search)){#the-input-element:text-(type=text)-state-and-search-state-(type=search)-5}
state.

Which of the
[`accept`{#the-input-element:attr-input-accept-2}](#attr-input-accept),
[`alpha`{#the-input-element:attr-input-alpha-2}](#attr-input-alpha),
[`alt`{#the-input-element:attr-input-alt-2}](#attr-input-alt),
[`autocomplete`{#the-input-element:attr-fe-autocomplete-2}](form-control-infrastructure.html#attr-fe-autocomplete),
[`checked`{#the-input-element:attr-input-checked-2}](#attr-input-checked),
[`colorspace`{#the-input-element:attr-input-colorspace-2}](#attr-input-colorspace),
[`dirname`{#the-input-element:attr-fe-dirname-2}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#the-input-element:attr-fs-formaction-2}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#the-input-element:attr-fs-formenctype-2}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#the-input-element:attr-fs-formmethod-2}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#the-input-element:attr-fs-formnovalidate-2}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#the-input-element:attr-fs-formtarget-2}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#the-input-element:attr-dim-height-2}](embedded-content-other.html#attr-dim-height),
[`list`{#the-input-element:attr-input-list-2}](#attr-input-list),
[`max`{#the-input-element:attr-input-max-2}](#attr-input-max),
[`maxlength`{#the-input-element:attr-input-maxlength-2}](#attr-input-maxlength),
[`min`{#the-input-element:attr-input-min-2}](#attr-input-min),
[`minlength`{#the-input-element:attr-input-minlength-2}](#attr-input-minlength),
[`multiple`{#the-input-element:attr-input-multiple-2}](#attr-input-multiple),
[`pattern`{#the-input-element:attr-input-pattern-3}](#attr-input-pattern),
[`placeholder`{#the-input-element:attr-input-placeholder-2}](#attr-input-placeholder),
[`readonly`{#the-input-element:attr-input-readonly-2}](#attr-input-readonly),
[`required`{#the-input-element:attr-input-required-2}](#attr-input-required),
[`size`{#the-input-element:attr-input-size-2}](#attr-input-size),
[`src`{#the-input-element:attr-input-src-2}](#attr-input-src),
[`step`{#the-input-element:attr-input-step-2}](#attr-input-step), and
[`width`{#the-input-element:attr-dim-width-2}](embedded-content-other.html#attr-dim-width)
content attributes, the
[`checked`{#the-input-element:dom-input-checked-2}](#dom-input-checked),
[`files`{#the-input-element:dom-input-files-2}](#dom-input-files),
[`valueAsDate`{#the-input-element:dom-input-valueasdate-2}](#dom-input-valueasdate),
[`valueAsNumber`{#the-input-element:dom-input-valueasnumber-2}](#dom-input-valueasnumber),
and [`list`{#the-input-element:dom-input-list-2}](#dom-input-list) IDL
attributes, the
[`select()`{#the-input-element:dom-textarea/input-select-2}](form-control-infrastructure.html#dom-textarea/input-select)
method, the
[`selectionStart`{#the-input-element:dom-textarea/input-selectionstart-2}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#the-input-element:dom-textarea/input-selectionend-2}](form-control-infrastructure.html#dom-textarea/input-selectionend),
and
[`selectionDirection`{#the-input-element:dom-textarea/input-selectiondirection-2}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
IDL attributes, the
[`setRangeText()`{#the-input-element:dom-textarea/input-setrangetext-3}](form-control-infrastructure.html#dom-textarea/input-setrangetext)
and
[`setSelectionRange()`{#the-input-element:dom-textarea/input-setselectionrange-2}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods, the
[`stepUp()`{#the-input-element:dom-input-stepup-2}](#dom-input-stepup)
and
[`stepDown()`{#the-input-element:dom-input-stepdown-2}](#dom-input-stepdown)
methods, and the
[`input`{#the-input-element:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#the-input-element:event-change}](indices.html#event-change)
events [apply]{#concept-input-apply .dfn} to an
[`input`{#the-input-element:the-input-element-2}](#the-input-element)
element depends on the state of its
[`type`{#the-input-element:attr-input-type-28}](#attr-input-type)
attribute. The subsections that define each type also clearly define in
normative \"bookkeeping\" sections which of these feature apply, and
which [do not apply]{#do-not-apply .dfn}, to each type. The behavior of
these features depends on whether they apply or not, as defined in their
various sections (q.v. for [content
attributes](#common-input-element-attributes), for
[APIs](#common-input-element-apis), for
[events](#common-input-element-events)).

The following table is non-normative and summarizes which of those
content attributes, IDL attributes, methods, and events
[apply](#concept-input-apply){#the-input-element:concept-input-apply} to
each state:

[Hidden](#hidden-state-(type=hidden)){#the-input-element:hidden-state-(type=hidden)-7}

[Text](#text-(type=text)-state-and-search-state-(type=search)){#the-input-element:text-(type=text)-state-and-search-state-(type=search)-6},
[Search](#text-(type=text)-state-and-search-state-(type=search)){#the-input-element:text-(type=text)-state-and-search-state-(type=search)-7}

[Telephone](#telephone-state-(type=tel)){#the-input-element:telephone-state-(type=tel)-3},
[URL](#url-state-(type=url)){#the-input-element:url-state-(type=url)-3}

[Email](#email-state-(type=email)){#the-input-element:email-state-(type=email)-3}

[Password](#password-state-(type=password)){#the-input-element:password-state-(type=password)-3}

[Date](#date-state-(type=date)){#the-input-element:date-state-(type=date)-3},
[Month](#month-state-(type=month)){#the-input-element:month-state-(type=month)-3},
[Week](#week-state-(type=week)){#the-input-element:week-state-(type=week)-3},
[Time](#time-state-(type=time)){#the-input-element:time-state-(type=time)-3}

[Local Date and
Time](#local-date-and-time-state-(type=datetime-local)){#the-input-element:local-date-and-time-state-(type=datetime-local)-3}

[Number](#number-state-(type=number)){#the-input-element:number-state-(type=number)-3}

[Range](#range-state-(type=range)){#the-input-element:range-state-(type=range)-3}

[Color](#color-state-(type=color)){#the-input-element:color-state-(type=color)-3}

[Checkbox](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)-3},
[Radio
Button](#radio-button-state-(type=radio)){#the-input-element:radio-button-state-(type=radio)-3}

[File
Upload](#file-upload-state-(type=file)){#the-input-element:file-upload-state-(type=file)-4}

[Submit
Button](#submit-button-state-(type=submit)){#the-input-element:submit-button-state-(type=submit)-3}

[Image
Button](#image-button-state-(type=image)){#the-input-element:image-button-state-(type=image)-3}

[Reset
Button](#reset-button-state-(type=reset)){#the-input-element:reset-button-state-(type=reset)-3},
[Button](#button-state-(type=button)){#the-input-element:button-state-(type=button)-3}

Content attributes

[`accept`{#the-input-element:attr-input-accept-3}](#attr-input-accept)

·

·

·

·

·

·

·

·

·

·

·

Yes

·

·

·

[`alpha`{#the-input-element:attr-input-alpha-3}](#attr-input-alpha)

·

·

·

·

·

·

·

·

·

Yes

·

·

·

·

·

[`alt`{#the-input-element:attr-input-alt-3}](#attr-input-alt)

·

·

·

·

·

·

·

·

·

·

·

·

·

Yes

·

[`autocomplete`{#the-input-element:attr-fe-autocomplete-3}](form-control-infrastructure.html#attr-fe-autocomplete)

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

·

·

·

·

·

[`checked`{#the-input-element:attr-input-checked-3}](#attr-input-checked)

·

·

·

·

·

·

·

·

·

·

Yes

·

·

·

·

[`colorspace`{#the-input-element:attr-input-colorspace-3}](#attr-input-colorspace)

·

·

·

·

·

·

·

·

·

Yes

·

·

·

·

·

[`dirname`{#the-input-element:attr-fe-dirname-3}](form-control-infrastructure.html#attr-fe-dirname)

Yes

Yes

Yes

Yes

Yes

·

·

·

·

·

·

·

Yes

·

·

[`formaction`{#the-input-element:attr-fs-formaction-3}](form-control-infrastructure.html#attr-fs-formaction)

·

·

·

·

·

·

·

·

·

·

·

·

Yes

Yes

·

[`formenctype`{#the-input-element:attr-fs-formenctype-3}](form-control-infrastructure.html#attr-fs-formenctype)

·

·

·

·

·

·

·

·

·

·

·

·

Yes

Yes

·

[`formmethod`{#the-input-element:attr-fs-formmethod-3}](form-control-infrastructure.html#attr-fs-formmethod)

·

·

·

·

·

·

·

·

·

·

·

·

Yes

Yes

·

[`formnovalidate`{#the-input-element:attr-fs-formnovalidate-3}](form-control-infrastructure.html#attr-fs-formnovalidate)

·

·

·

·

·

·

·

·

·

·

·

·

Yes

Yes

·

[`formtarget`{#the-input-element:attr-fs-formtarget-3}](form-control-infrastructure.html#attr-fs-formtarget)

·

·

·

·

·

·

·

·

·

·

·

·

Yes

Yes

·

[`height`{#the-input-element:attr-dim-height-3}](embedded-content-other.html#attr-dim-height)

·

·

·

·

·

·

·

·

·

·

·

·

·

Yes

·

[`list`{#the-input-element:attr-input-list-3}](#attr-input-list)

·

Yes

Yes

Yes

·

Yes

Yes

Yes

Yes

Yes

·

·

·

·

·

[`max`{#the-input-element:attr-input-max-3}](#attr-input-max)

·

·

·

·

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

[`maxlength`{#the-input-element:attr-input-maxlength-3}](#attr-input-maxlength)

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

·

·

·

·

[`min`{#the-input-element:attr-input-min-3}](#attr-input-min)

·

·

·

·

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

[`minlength`{#the-input-element:attr-input-minlength-3}](#attr-input-minlength)

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

·

·

·

·

[`multiple`{#the-input-element:attr-input-multiple-3}](#attr-input-multiple)

·

·

·

Yes

·

·

·

·

·

·

·

Yes

·

·

·

[`pattern`{#the-input-element:attr-input-pattern-4}](#attr-input-pattern)

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

·

·

·

·

[`placeholder`{#the-input-element:attr-input-placeholder-3}](#attr-input-placeholder)

·

Yes

Yes

Yes

Yes

·

·

Yes

·

·

·

·

·

·

·

[`popovertarget`{#the-input-element:attr-popovertarget-2}](popover.html#attr-popovertarget)

·

·

·

·

·

·

·

·

·

·

·

·

Yes

Yes

Yes

[`popovertargetaction`{#the-input-element:attr-popovertargetaction-2}](popover.html#attr-popovertargetaction)

·

·

·

·

·

·

·

·

·

·

·

·

Yes

Yes

Yes

[`readonly`{#the-input-element:attr-input-readonly-3}](#attr-input-readonly)

·

Yes

Yes

Yes

Yes

Yes

Yes

Yes

·

·

·

·

·

·

·

[`required`{#the-input-element:attr-input-required-3}](#attr-input-required)

·

Yes

Yes

Yes

Yes

Yes

Yes

Yes

·

·

Yes

Yes

·

·

·

[`size`{#the-input-element:attr-input-size-3}](#attr-input-size)

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

·

·

·

·

[`src`{#the-input-element:attr-input-src-3}](#attr-input-src)

·

·

·

·

·

·

·

·

·

·

·

·

·

Yes

·

[`step`{#the-input-element:attr-input-step-3}](#attr-input-step)

·

·

·

·

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

[`width`{#the-input-element:attr-dim-width-3}](embedded-content-other.html#attr-dim-width)

·

·

·

·

·

·

·

·

·

·

·

·

·

Yes

·

IDL attributes and methods

[`checked`{#the-input-element:dom-input-checked-3}](#dom-input-checked)

·

·

·

·

·

·

·

·

·

·

Yes

·

·

·

·

[`files`{#the-input-element:dom-input-files-3}](#dom-input-files)

·

·

·

·

·

·

·

·

·

·

·

Yes

·

·

·

[`value`{#the-input-element:dom-input-value-2}](#dom-input-value)

[default](#dom-input-value-default){#the-input-element:dom-input-value-default}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-2}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-3}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-4}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-5}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-6}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-7}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-8}

[value](#dom-input-value-value){#the-input-element:dom-input-value-value-9}

[default/on](#dom-input-value-default-on){#the-input-element:dom-input-value-default-on}

[filename](#dom-input-value-filename){#the-input-element:dom-input-value-filename}

[default](#dom-input-value-default){#the-input-element:dom-input-value-default-2}

[default](#dom-input-value-default){#the-input-element:dom-input-value-default-3}

[default](#dom-input-value-default){#the-input-element:dom-input-value-default-4}

[`valueAsDate`{#the-input-element:dom-input-valueasdate-3}](#dom-input-valueasdate)

·

·

·

·

·

Yes

·

·

·

·

·

·

·

·

·

[`valueAsNumber`{#the-input-element:dom-input-valueasnumber-3}](#dom-input-valueasnumber)

·

·

·

·

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

[`list`{#the-input-element:dom-input-list-3}](#dom-input-list)

·

Yes

Yes

Yes

·

Yes

Yes

Yes

Yes

Yes

·

·

·

·

·

[`select()`{#the-input-element:dom-textarea/input-select-3}](form-control-infrastructure.html#dom-textarea/input-select)

·

Yes

Yes

Yes†

Yes

Yes†

Yes†

Yes†

·

Yes†

·

Yes†

·

·

·

[`selectionStart`{#the-input-element:dom-textarea/input-selectionstart-3}](form-control-infrastructure.html#dom-textarea/input-selectionstart)

·

Yes

Yes

·

Yes

·

·

·

·

·

·

·

·

·

·

[`selectionEnd`{#the-input-element:dom-textarea/input-selectionend-3}](form-control-infrastructure.html#dom-textarea/input-selectionend)

·

Yes

Yes

·

Yes

·

·

·

·

·

·

·

·

·

·

[`selectionDirection`{#the-input-element:dom-textarea/input-selectiondirection-3}](form-control-infrastructure.html#dom-textarea/input-selectiondirection)

·

Yes

Yes

·

Yes

·

·

·

·

·

·

·

·

·

·

[`setRangeText()`{#the-input-element:dom-textarea/input-setrangetext-4}](form-control-infrastructure.html#dom-textarea/input-setrangetext)

·

Yes

Yes

·

Yes

·

·

·

·

·

·

·

·

·

·

[`setSelectionRange()`{#the-input-element:dom-textarea/input-setselectionrange-3}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)

·

Yes

Yes

·

Yes

·

·

·

·

·

·

·

·

·

·

[`stepDown()`{#the-input-element:dom-input-stepdown-3}](#dom-input-stepdown)

·

·

·

·

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

[`stepUp()`{#the-input-element:dom-input-stepup-3}](#dom-input-stepup)

·

·

·

·

·

Yes

Yes

Yes

Yes

·

·

·

·

·

·

Events

[`input`{#the-input-element:event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
event

·

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

·

·

·

[`change`{#the-input-element:event-change-2}](indices.html#event-change)
event

·

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

·

·

·

[† If the control has no selectable text, the
[`select()`{#the-input-element:dom-textarea/input-select-4}](form-control-infrastructure.html#dom-textarea/input-select)
method results in a no-op, with no
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-input-element:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#the-input-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.]{.small}

Some states of the
[`type`{#the-input-element:attr-input-type-29}](#attr-input-type)
attribute define a [value sanitization
algorithm]{#value-sanitization-algorithm .dfn}.

Each
[`input`{#the-input-element:the-input-element-3}](#the-input-element)
element has a
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value},
which is exposed by the
[`value`{#the-input-element:dom-input-value-3}](#dom-input-value) IDL
attribute. Some states define an [algorithm to convert a string to a
number]{#concept-input-value-string-number .dfn}, an [algorithm to
convert a number to a string]{#concept-input-value-number-string .dfn},
an [algorithm to convert a string to a `Date`
object]{#concept-input-value-string-date .dfn}, and an [algorithm to
convert a `Date` object to a string]{#concept-input-value-date-string
.dfn}, which are used by
[`max`{#the-input-element:attr-input-max-4}](#attr-input-max),
[`min`{#the-input-element:attr-input-min-4}](#attr-input-min),
[`step`{#the-input-element:attr-input-step-4}](#attr-input-step),
[`valueAsDate`{#the-input-element:dom-input-valueasdate-4}](#dom-input-valueasdate),
[`valueAsNumber`{#the-input-element:dom-input-valueasnumber-4}](#dom-input-valueasnumber),
and
[`stepUp()`{#the-input-element:dom-input-stepup-4}](#dom-input-stepup).

An [`input`{#the-input-element:the-input-element-4}](#the-input-element)
element\'s [dirty value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-input-element:concept-fe-dirty}
must be set to true whenever the user interacts with the control in a
way that changes the
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-2}.
(It is also set to true when the value is programmatically changed, as
described in the definition of the
[`value`{#the-input-element:dom-input-value-4}](#dom-input-value) IDL
attribute.)

The [`value`]{#attr-input-value .dfn dfn-for="input"
dfn-type="element-attr"} content attribute gives the default
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-3}
of the
[`input`{#the-input-element:the-input-element-5}](#the-input-element)
element. When the
[`value`{#the-input-element:attr-input-value-2}](#attr-input-value)
content attribute is added, set, or removed, if the control\'s [dirty
value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-input-element:concept-fe-dirty-2}
is false, the user agent must set the
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-4}
of the element to the value of the
[`value`{#the-input-element:attr-input-value-3}](#attr-input-value)
content attribute, if there is one, or the empty string otherwise, and
then run the current [value sanitization
algorithm](#value-sanitization-algorithm){#the-input-element:value-sanitization-algorithm},
if one is defined.

Each
[`input`{#the-input-element:the-input-element-6}](#the-input-element)
element has a
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked},
which is exposed by the
[`checked`{#the-input-element:dom-input-checked-4}](#dom-input-checked)
IDL attribute.

Each
[`input`{#the-input-element:the-input-element-7}](#the-input-element)
element has a boolean [dirty checkedness
flag]{#concept-input-checked-dirty-flag .dfn}. When it is true, the
element is said to have a [*dirty
checkedness*]{#concept-input-checked-dirty .dfn}. The [dirty checkedness
flag](#concept-input-checked-dirty-flag){#the-input-element:concept-input-checked-dirty-flag}
must be initially set to false when the element is created, and must be
set to true whenever the user interacts with the control in a way that
changes the
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-2}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#checked](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#checked "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

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

The [`checked`]{#attr-input-checked .dfn dfn-for="input"
dfn-type="element-attr"} content attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-input-element:boolean-attribute}
that gives the default
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-3}
of the
[`input`{#the-input-element:the-input-element-8}](#the-input-element)
element. When the
[`checked`{#the-input-element:attr-input-checked-4}](#attr-input-checked)
content attribute is added, if the control does not have *[dirty
checkedness](#concept-input-checked-dirty)*, the user agent must set the
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-4}
of the element to true; when the
[`checked`{#the-input-element:attr-input-checked-5}](#attr-input-checked)
content attribute is removed, if the control does not have *[dirty
checkedness](#concept-input-checked-dirty)*, the user agent must set the
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-5}
of the element to false.

The [reset
algorithm](form-control-infrastructure.html#concept-form-reset-control){#the-input-element:concept-form-reset-control}
for
[`input`{#the-input-element:the-input-element-9}](#the-input-element)
elements is to set its [user
validity](form-control-infrastructure.html#user-validity){#the-input-element:user-validity},
[dirty value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-input-element:concept-fe-dirty-3},
and [dirty checkedness
flag](#concept-input-checked-dirty-flag){#the-input-element:concept-input-checked-dirty-flag-2}
back to false, set the
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-5}
of the element to the value of the
[`value`{#the-input-element:attr-input-value-4}](#attr-input-value)
content attribute, if there is one, or the empty string otherwise, set
the
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-6}
of the element to true if the element has a
[`checked`{#the-input-element:attr-input-checked-6}](#attr-input-checked)
content attribute and false if it does not, empty the list of [selected
files](#concept-input-type-file-selected){#the-input-element:concept-input-type-file-selected},
and then invoke the [value sanitization
algorithm](#value-sanitization-algorithm){#the-input-element:value-sanitization-algorithm-2},
if the [`type`{#the-input-element:attr-input-type-30}](#attr-input-type)
attribute\'s current state defines one.

Each
[`input`{#the-input-element:the-input-element-10}](#the-input-element)
element can be
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*. Except
where otherwise specified, an
[`input`{#the-input-element:the-input-element-11}](#the-input-element)
element is always
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*.
Similarly, except where otherwise specified, the user agent should not
allow the user to modify the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-6}
or
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-7}.

When an
[`input`{#the-input-element:the-input-element-12}](#the-input-element)
element is
[disabled](form-control-infrastructure.html#concept-fe-disabled){#the-input-element:concept-fe-disabled},
it is not
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*.

The
[`readonly`{#the-input-element:attr-input-readonly-4}](#attr-input-readonly)
attribute can also in some cases (e.g. for the
[Date](#date-state-(type=date)){#the-input-element:date-state-(type=date)-4}
state, but not the
[Checkbox](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)-4}
state) stop an
[`input`{#the-input-element:the-input-element-13}](#the-input-element)
element from being
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*.

The
[`input`{#the-input-element:the-input-element-14}](#the-input-element)
element can [support a picker]{#input-support-picker .dfn}. A picker is
a user interface element that allows the end user to choose a value.
Whether an
[`input`{#the-input-element:the-input-element-15}](#the-input-element)
element supports a picker depends on the
[`type`{#the-input-element:attr-input-type-31}](#attr-input-type)
attribute state and
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-input-element:implementation-defined
x-internal="implementation-defined"} behavior. An
[`input`{#the-input-element:the-input-element-16}](#the-input-element)
element must support a picker when its
[`type`{#the-input-element:attr-input-type-32}](#attr-input-type)
attribute is in the [File
Upload](#file-upload-state-(type=file)){#the-input-element:file-upload-state-(type=file)-5}
state.

::: note
As of the time of this writing, typical browser implementations show
such picker UI for:

- [`input`{#the-input-element:the-input-element-17}](#the-input-element)
  elements whose
  [`type`{#the-input-element:attr-input-type-33}](#attr-input-type)
  attributes are in the
  [Date](#date-state-(type=date)){#the-input-element:date-state-(type=date)-5},
  [Month](#month-state-(type=month)){#the-input-element:month-state-(type=month)-4},
  [Week](#week-state-(type=week)){#the-input-element:week-state-(type=week)-4},
  [Time](#time-state-(type=time)){#the-input-element:time-state-(type=time)-4},
  [Local Date and
  Time](#local-date-and-time-state-(type=datetime-local)){#the-input-element:local-date-and-time-state-(type=datetime-local)-4},
  and
  [Color](#color-state-(type=color)){#the-input-element:color-state-(type=color)-4}
  states;

- [`input`{#the-input-element:the-input-element-18}](#the-input-element)
  elements in various states that have a [suggestions source
  element](#concept-input-list){#the-input-element:concept-input-list};

- [`input`{#the-input-element:the-input-element-19}](#the-input-element)
  elements whose
  [`type`{#the-input-element:attr-input-type-34}](#attr-input-type)
  attribute is in the [File
  Upload](#file-upload-state-(type=file)){#the-input-element:file-upload-state-(type=file)-6}
  state; and

- [`select`{#the-input-element:the-select-element}](form-elements.html#the-select-element)
  elements.
:::

The [cloning
steps](https://dom.spec.whatwg.org/#concept-node-clone-ext){#the-input-element:concept-node-clone-ext
x-internal="concept-node-clone-ext"} for
[`input`{#the-input-element:the-input-element-20}](#the-input-element)
elements given `node`{.variable}, `copy`{.variable}, and
`subtree`{.variable} are to propagate the
[value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-7},
[dirty value
flag](form-control-infrastructure.html#concept-fe-dirty){#the-input-element:concept-fe-dirty-4},
[checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-8},
and [dirty checkedness
flag](#concept-input-checked-dirty-flag){#the-input-element:concept-input-checked-dirty-flag-3}
from `node`{.variable} to `copy`{.variable}.

The [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-input-element:activation-behaviour
x-internal="activation-behaviour"} for
[`input`{#the-input-element:the-input-element-21}](#the-input-element)
elements `element`{.variable}, given `event`{.variable}, are these
steps:

1.  If `element`{.variable} is not
    *[mutable](form-control-infrastructure.html#concept-fe-mutable)*,
    and `element`{.variable}\'s
    [`type`{#the-input-element:attr-input-type-35}](#attr-input-type)
    attribute is neither in the
    [Checkbox](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)-5}
    nor in the
    [Radio](#radio-button-state-(type=radio)){#the-input-element:radio-button-state-(type=radio)-4}
    state, then return.

2.  Run `element`{.variable}\'s [input activation
    behavior]{#input-activation-behavior .dfn}, if any, and do nothing
    otherwise.

3.  If `element`{.variable} has a [form
    owner](form-control-infrastructure.html#form-owner){#the-input-element:form-owner}
    and `element`{.variable}\'s
    [`type`{#the-input-element:attr-input-type-36}](#attr-input-type)
    attribute is not in the
    [Button](#button-state-(type=button)){#the-input-element:button-state-(type=button)-4}
    state, then return.

4.  Run the [popover target attribute activation
    behavior](popover.html#popover-target-attribute-activation-behavior){#the-input-element:popover-target-attribute-activation-behavior}
    given `element`{.variable} and `event`{.variable}\'s
    [target](https://dom.spec.whatwg.org/#concept-event-target){#the-input-element:concept-event-target
    x-internal="concept-event-target"}.

Recall that an element\'s [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-input-element:activation-behaviour-2
x-internal="activation-behaviour"} runs for both user-initiated
activations and for synthetic activations (e.g., via `el.click()`). User
agents might also have behaviors for a given control --- not specified
here --- that are triggered only by true user-initiated activations. A
common choice is to [show the picker, if
applicable](#show-the-picker,-if-applicable){#the-input-element:show-the-picker,-if-applicable},
for the control. In contrast, the [input activation
behavior](#input-activation-behavior){#the-input-element:input-activation-behavior}
only shows pickers for the special historical cases of the [File
Upload](#file-upload-state-(type=file)){#the-input-element:file-upload-state-(type=file)-7}
and
[Color](#color-state-(type=color)){#the-input-element:color-state-(type=color)-5}
states.

The [legacy-pre-activation
behavior](https://dom.spec.whatwg.org/#eventtarget-legacy-pre-activation-behavior){#the-input-element:legacy-pre-activation-behavior
x-internal="legacy-pre-activation-behavior"} for
[`input`{#the-input-element:the-input-element-22}](#the-input-element)
elements are these steps:

1.  If this element\'s
    [`type`{#the-input-element:attr-input-type-37}](#attr-input-type)
    attribute is in the [Checkbox
    state](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)-6},
    then set this element\'s
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-9}
    to its opposite value (i.e. true if it is false, false if it is
    true) and set this element\'s
    [`indeterminate`{#the-input-element:dom-input-indeterminate-2}](#dom-input-indeterminate)
    IDL attribute to false.

2.  If this element\'s
    [`type`{#the-input-element:attr-input-type-38}](#attr-input-type)
    attribute is in the [Radio Button
    state](#radio-button-state-(type=radio)){#the-input-element:radio-button-state-(type=radio)-5},
    then get a reference to the element in this element\'s [radio button
    group](#radio-button-group){#the-input-element:radio-button-group}
    that has its
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-10}
    set to true, if any, and then set this element\'s
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-11}
    to true.

The [legacy-canceled-activation
behavior](https://dom.spec.whatwg.org/#eventtarget-legacy-canceled-activation-behavior){#the-input-element:legacy-canceled-activation-behavior
x-internal="legacy-canceled-activation-behavior"} for
[`input`{#the-input-element:the-input-element-23}](#the-input-element)
elements are these steps:

1.  If the element\'s
    [`type`{#the-input-element:attr-input-type-39}](#attr-input-type)
    attribute is in the [Checkbox
    state](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)-7},
    then set the element\'s
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-12}
    and the element\'s
    [`indeterminate`{#the-input-element:dom-input-indeterminate-3}](#dom-input-indeterminate)
    IDL attribute back to the values they had before the
    [legacy-pre-activation
    behavior](https://dom.spec.whatwg.org/#eventtarget-legacy-pre-activation-behavior){#the-input-element:legacy-pre-activation-behavior-2
    x-internal="legacy-pre-activation-behavior"} was run.

2.  If this element\'s
    [`type`{#the-input-element:attr-input-type-40}](#attr-input-type)
    attribute is in the [Radio Button
    state](#radio-button-state-(type=radio)){#the-input-element:radio-button-state-(type=radio)-6},
    then if the element to which a reference was obtained in the
    [legacy-pre-activation
    behavior](https://dom.spec.whatwg.org/#eventtarget-legacy-pre-activation-behavior){#the-input-element:legacy-pre-activation-behavior-3
    x-internal="legacy-pre-activation-behavior"}, if any, is still in
    what is now this element\'s [radio button
    group](#radio-button-group){#the-input-element:radio-button-group-2},
    if it still has one, and if so, set that element\'s
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-13}
    to true; or else, if there was no such element, or that element is
    no longer in this element\'s [radio button
    group](#radio-button-group){#the-input-element:radio-button-group-3},
    or if this element no longer has a [radio button
    group](#radio-button-group){#the-input-element:radio-button-group-4},
    set this element\'s
    [checkedness](form-control-infrastructure.html#concept-fe-checked){#the-input-element:concept-fe-checked-14}
    to false.

------------------------------------------------------------------------

When an
[`input`{#the-input-element:the-input-element-24}](#the-input-element)
element is first created, the element\'s rendering and behavior must be
set to the rendering and behavior defined for the
[`type`{#the-input-element:attr-input-type-41}](#attr-input-type)
attribute\'s state, and the [value sanitization
algorithm](#value-sanitization-algorithm){#the-input-element:value-sanitization-algorithm-3},
if one is defined for the
[`type`{#the-input-element:attr-input-type-42}](#attr-input-type)
attribute\'s state, must be invoked.

::: {#input-type-change}
When an
[`input`{#the-input-element:the-input-element-25}](#the-input-element)
element\'s
[`type`{#the-input-element:attr-input-type-43}](#attr-input-type)
attribute changes state, the user agent must run the following steps:

1.  If the previous state of the element\'s
    [`type`{#the-input-element:attr-input-type-44}](#attr-input-type)
    attribute put the
    [`value`{#the-input-element:dom-input-value-5}](#dom-input-value)
    IDL attribute in the *[value](#dom-input-value-value)* mode, and the
    element\'s
    [value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-8}
    is not the empty string, and the new state of the element\'s
    [`type`{#the-input-element:attr-input-type-45}](#attr-input-type)
    attribute puts the
    [`value`{#the-input-element:dom-input-value-6}](#dom-input-value)
    IDL attribute in either the *[default](#dom-input-value-default)*
    mode or the *[default/on](#dom-input-value-default-on)* mode, then
    set the element\'s
    [`value`{#the-input-element:attr-input-value-5}](#attr-input-value)
    content attribute to the element\'s
    [value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-9}.

2.  Otherwise, if the previous state of the element\'s
    [`type`{#the-input-element:attr-input-type-46}](#attr-input-type)
    attribute put the
    [`value`{#the-input-element:dom-input-value-7}](#dom-input-value)
    IDL attribute in any mode other than the
    *[value](#dom-input-value-value)* mode, and the new state of the
    element\'s
    [`type`{#the-input-element:attr-input-type-47}](#attr-input-type)
    attribute puts the
    [`value`{#the-input-element:dom-input-value-8}](#dom-input-value)
    IDL attribute in the *[value](#dom-input-value-value)* mode, then
    set the
    [value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-10}
    of the element to the value of the
    [`value`{#the-input-element:attr-input-value-6}](#attr-input-value)
    content attribute, if there is one, or the empty string otherwise,
    and then set the control\'s [dirty value
    flag](form-control-infrastructure.html#concept-fe-dirty){#the-input-element:concept-fe-dirty-5}
    to false.

3.  Otherwise, if the previous state of the element\'s
    [`type`{#the-input-element:attr-input-type-48}](#attr-input-type)
    attribute put the
    [`value`{#the-input-element:dom-input-value-9}](#dom-input-value)
    IDL attribute in any mode other than the
    *[filename](#dom-input-value-filename)* mode, and the new state of
    the element\'s
    [`type`{#the-input-element:attr-input-type-49}](#attr-input-type)
    attribute puts the
    [`value`{#the-input-element:dom-input-value-10}](#dom-input-value)
    IDL attribute in the *[filename](#dom-input-value-filename)* mode,
    then set the
    [value](form-control-infrastructure.html#concept-fe-value){#the-input-element:concept-fe-value-11}
    of the element to the empty string.

4.  Update the element\'s rendering and behavior to the new state\'s.

5.  [Signal a type change]{#signal-a-type-change .dfn} for the element.
    (The [Radio
    Button](#radio-button-state-(type=radio)){#the-input-element:radio-button-state-(type=radio)-7}
    state uses this, in particular.)

6.  Invoke the [value sanitization
    algorithm](#value-sanitization-algorithm){#the-input-element:value-sanitization-algorithm-4},
    if one is defined for the
    [`type`{#the-input-element:attr-input-type-50}](#attr-input-type)
    attribute\'s new state.

7.  Let `previouslySelectable`{.variable} be true if
    [`setRangeText()`{#the-input-element:dom-textarea/input-setrangetext-5}](form-control-infrastructure.html#dom-textarea/input-setrangetext)
    previously
    [applied](#concept-input-apply){#the-input-element:concept-input-apply-2}
    to the element, and false otherwise.

8.  Let `nowSelectable`{.variable} be true if
    [`setRangeText()`{#the-input-element:dom-textarea/input-setrangetext-6}](form-control-infrastructure.html#dom-textarea/input-setrangetext)
    now
    [applies](#concept-input-apply){#the-input-element:concept-input-apply-3}
    to the element, and false otherwise.

9.  If `previouslySelectable`{.variable} is false and
    `nowSelectable`{.variable} is true, set the element\'s [text entry
    cursor
    position](form-control-infrastructure.html#concept-textarea/input-cursor){#the-input-element:concept-textarea/input-cursor}
    to the beginning of the text control, and [set its selection
    direction](form-control-infrastructure.html#set-the-selection-direction){#the-input-element:set-the-selection-direction}
    to \"`none`\".
:::

------------------------------------------------------------------------

The
[`name`{#the-input-element:attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
attribute represents the element\'s name. The
[`dirname`{#the-input-element:attr-fe-dirname-4}](form-control-infrastructure.html#attr-fe-dirname)
attribute controls how the element\'s
[directionality](dom.html#the-directionality){#the-input-element:the-directionality-2}
is submitted. The
[`disabled`{#the-input-element:attr-fe-disabled-2}](form-control-infrastructure.html#attr-fe-disabled)
attribute is used to make the control non-interactive and to prevent its
value from being submitted. The
[`form`{#the-input-element:attr-fae-form-2}](form-control-infrastructure.html#attr-fae-form)
attribute is used to explicitly associate the
[`input`{#the-input-element:the-input-element-26}](#the-input-element)
element with its [form
owner](form-control-infrastructure.html#form-owner){#the-input-element:form-owner-2}.
The
[`autocomplete`{#the-input-element:attr-fe-autocomplete-4}](form-control-infrastructure.html#attr-fe-autocomplete)
attribute controls how the user agent provides autofill behavior.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement#indeterminate](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement#indeterminate "The HTMLInputElement interface provides special properties and methods for manipulating the options, layout, and presentation of <input> elements.")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera≤12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android≤12.1+]{.opera_android .yes}

------------------------------------------------------------------------

[[caniuse.com
table](https://caniuse.com/#feat=indeterminate-checkbox "indeterminate checkbox")]{.caniuse}
:::
::::
:::::

The [`indeterminate`]{#dom-input-indeterminate .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"} IDL attribute must
initially be set to false. On getting, it must return the last value it
was set to. On setting, it must be set to the new value. It has no
effect except for changing the appearance of
[checkbox](#checkbox-state-(type=checkbox)){#the-input-element:checkbox-state-(type=checkbox)-8}
controls.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/multiple](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/multiple "The HTMLInputElement.multiple property indicates if an input can have more than one value. Firefox currently only supports multiple for <input type="file">.")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
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

The [`accept`]{#dom-input-accept .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"}, [`alpha`]{#dom-input-alpha .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"}, [`alt`]{#dom-input-alt
.dfn dfn-for="HTMLInputElement" dfn-type="attribute"},
[`max`]{#dom-input-max .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"}, [`min`]{#dom-input-min .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"},
[`multiple`]{#dom-input-multiple .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"}, [`pattern`]{#dom-input-pattern .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"},
[`placeholder`]{#dom-input-placeholder .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"}, [`required`]{#dom-input-required .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"},
[`size`]{#dom-input-size .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"}, [`src`]{#dom-input-src .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"}, and
[`step`]{#dom-input-step .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attributes must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect}
the respective content attributes of the same name. The
[`dirName`]{#dom-input-dirname .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-2}
the
[`dirname`{#the-input-element:attr-fe-dirname-5}](form-control-infrastructure.html#attr-fe-dirname)
content attribute. The [`readOnly`]{#dom-input-readonly .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-3}
the
[`readonly`{#the-input-element:attr-input-readonly-5}](#attr-input-readonly)
content attribute. The [`defaultChecked`]{#dom-input-defaultchecked .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-4}
the
[`checked`{#the-input-element:attr-input-checked-7}](#attr-input-checked)
content attribute. The [`defaultValue`]{#dom-input-defaultvalue .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-5}
the [`value`{#the-input-element:attr-input-value-7}](#attr-input-value)
content attribute.

The [`colorSpace`]{#dom-input-colorspace .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-6}
the
[`colorspace`{#the-input-element:attr-input-colorspace-4}](#attr-input-colorspace)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-input-element:limited-to-only-known-values}.
The [`type`]{#dom-input-type .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-7}
the respective content attribute of the same name, [limited to only
known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-input-element:limited-to-only-known-values-2}.
The [`maxLength`]{#dom-input-maxlength .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-8}
the
[`maxlength`{#the-input-element:attr-input-maxlength-4}](#attr-input-maxlength)
content attribute, [limited to only non-negative
numbers](common-dom-interfaces.html#limited-to-only-non-negative-numbers){#the-input-element:limited-to-only-non-negative-numbers}.
The [`minLength`]{#dom-input-minlength .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-input-element:reflect-9}
the
[`minlength`{#the-input-element:attr-input-minlength-4}](#attr-input-minlength)
content attribute, [limited to only non-negative
numbers](common-dom-interfaces.html#limited-to-only-non-negative-numbers){#the-input-element:limited-to-only-non-negative-numbers-2}.

The IDL attributes [`width`]{#dom-input-width .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"} and
[`height`]{#dom-input-height .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} must return the rendered width and height of the
image, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-input-element:'px'
x-internal="'px'"}, if an image is [being
rendered](rendering.html#being-rendered){#the-input-element:being-rendered};
or else the [natural width and
height](https://drafts.csswg.org/css-images/#natural-dimensions){#the-input-element:natural-dimensions
x-internal="natural-dimensions"} of the image, in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#the-input-element:'px'-2
x-internal="'px'"}, if an image is *[available](#input-img-available)*
but not [being
rendered](rendering.html#being-rendered){#the-input-element:being-rendered-2};
or else 0, if no image is *[available](#input-img-available)*. When the
[`input`{#the-input-element:the-input-element-27}](#the-input-element)
element\'s
[`type`{#the-input-element:attr-input-type-51}](#attr-input-type)
attribute is not in the [Image
Button](#image-button-state-(type=image)){#the-input-element:image-button-state-(type=image)-4}
state, then no image is *[available](#input-img-available)*.
[\[CSS\]](references.html#refsCSS)

On setting, they must act as if they
[reflected](common-dom-interfaces.html#reflect){#the-input-element:reflect-10}
the respective content attributes of the same name.

The
[`willValidate`{#the-input-element:dom-cva-willvalidate-2}](form-control-infrastructure.html#dom-cva-willvalidate),
[`validity`{#the-input-element:dom-cva-validity-2}](form-control-infrastructure.html#dom-cva-validity),
and
[`validationMessage`{#the-input-element:dom-cva-validationmessage-2}](form-control-infrastructure.html#dom-cva-validationmessage)
IDL attributes, and the
[`checkValidity()`{#the-input-element:dom-cva-checkvalidity-2}](form-control-infrastructure.html#dom-cva-checkvalidity),
[`reportValidity()`{#the-input-element:dom-cva-reportvalidity-2}](form-control-infrastructure.html#dom-cva-reportvalidity),
and
[`setCustomValidity()`{#the-input-element:dom-cva-setcustomvalidity-2}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
methods, are part of the [constraint validation
API](form-control-infrastructure.html#the-constraint-validation-api){#the-input-element:the-constraint-validation-api}.
The
[`labels`{#the-input-element:dom-lfe-labels-2}](forms.html#dom-lfe-labels)
IDL attribute provides a list of the element\'s
[`label`{#the-input-element:the-label-element}](forms.html#the-label-element)s.
The
[`select()`{#the-input-element:dom-textarea/input-select-5}](form-control-infrastructure.html#dom-textarea/input-select),
[`selectionStart`{#the-input-element:dom-textarea/input-selectionstart-4}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#the-input-element:dom-textarea/input-selectionend-4}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#the-input-element:dom-textarea/input-selectiondirection-4}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`setRangeText()`{#the-input-element:dom-textarea/input-setrangetext-7}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#the-input-element:dom-textarea/input-setselectionrange-4}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods and IDL attributes expose the element\'s text selection. The
[`disabled`{#the-input-element:dom-fe-disabled-2}](form-control-infrastructure.html#dom-fe-disabled),
[`form`{#the-input-element:dom-fae-form-2}](form-control-infrastructure.html#dom-fae-form),
and
[`name`{#the-input-element:dom-fe-name-2}](form-control-infrastructure.html#dom-fe-name)
IDL attributes are part of the element\'s forms API.

##### [4.10.5.1]{.secno} States of the [`type`{#states-of-the-type-attribute:attr-input-type}](#attr-input-type) attribute[](#states-of-the-type-attribute){.self-link}

###### [4.10.5.1.1]{.secno} [Hidden]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=hidden`)[](#hidden-state-(type=hidden)){.self-link} {#hidden-state-(type=hidden)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/hidden](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/hidden "<input> elements of type hidden let web developers include data that cannot be seen or modified by users when a form is submitted. For example, the ID of the content that is currently being ordered or edited, or a unique security token. Hidden inputs are completely invisible in the rendered page, and there is no way to make it visible in the page's content.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera2+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

When an
[`input`{#hidden-state-(type=hidden):the-input-element}](#the-input-element)
element\'s
[`type`{#hidden-state-(type=hidden):attr-input-type}](#attr-input-type)
attribute is in the
[Hidden](#hidden-state-(type=hidden)){#hidden-state-(type=hidden):hidden-state-(type=hidden)}
state, the rules in this section apply.

The
[`input`{#hidden-state-(type=hidden):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#hidden-state-(type=hidden):represents}
a value that is not intended to be examined or manipulated by the user.

**Constraint validation**: If an
[`input`{#hidden-state-(type=hidden):the-input-element-3}](#the-input-element)
element\'s
[`type`{#hidden-state-(type=hidden):attr-input-type-2}](#attr-input-type)
attribute is in the
[Hidden](#hidden-state-(type=hidden)){#hidden-state-(type=hidden):hidden-state-(type=hidden)-2}
state, it is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#hidden-state-(type=hidden):barred-from-constraint-validation}.

If the
[`name`{#hidden-state-(type=hidden):attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
attribute is present and has a value that is an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#hidden-state-(type=hidden):ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for
\"[`_charset_`{#hidden-state-(type=hidden):attr-fe-name-charset}](form-control-infrastructure.html#attr-fe-name-charset)\",
then the element\'s
[`value`{#hidden-state-(type=hidden):attr-input-value}](#attr-input-value)
attribute must be omitted.

::: bookkeeping
The
[`autocomplete`{#hidden-state-(type=hidden):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete)
and
[`dirname`{#hidden-state-(type=hidden):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname)
content attributes
[apply](#concept-input-apply){#hidden-state-(type=hidden):concept-input-apply}
to this element.

The
[`value`{#hidden-state-(type=hidden):dom-input-value}](#dom-input-value)
IDL attribute
[applies](#concept-input-apply){#hidden-state-(type=hidden):concept-input-apply-2}
to this element and is in mode
[default](#dom-input-value-default){#hidden-state-(type=hidden):dom-input-value-default}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#hidden-state-(type=hidden):do-not-apply} to the
element:
[`accept`{#hidden-state-(type=hidden):attr-input-accept}](#attr-input-accept),
[`alpha`{#hidden-state-(type=hidden):attr-input-alpha}](#attr-input-alpha),
[`alt`{#hidden-state-(type=hidden):attr-input-alt}](#attr-input-alt),
[`checked`{#hidden-state-(type=hidden):attr-input-checked}](#attr-input-checked),
[`colorspace`{#hidden-state-(type=hidden):attr-input-colorspace}](#attr-input-colorspace),
[`formaction`{#hidden-state-(type=hidden):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#hidden-state-(type=hidden):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#hidden-state-(type=hidden):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#hidden-state-(type=hidden):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#hidden-state-(type=hidden):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#hidden-state-(type=hidden):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#hidden-state-(type=hidden):attr-input-list}](#attr-input-list),
[`max`{#hidden-state-(type=hidden):attr-input-max}](#attr-input-max),
[`maxlength`{#hidden-state-(type=hidden):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#hidden-state-(type=hidden):attr-input-min}](#attr-input-min),
[`minlength`{#hidden-state-(type=hidden):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#hidden-state-(type=hidden):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#hidden-state-(type=hidden):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#hidden-state-(type=hidden):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#hidden-state-(type=hidden):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#hidden-state-(type=hidden):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`readonly`{#hidden-state-(type=hidden):attr-input-readonly}](#attr-input-readonly),
[`required`{#hidden-state-(type=hidden):attr-input-required}](#attr-input-required),
[`size`{#hidden-state-(type=hidden):attr-input-size}](#attr-input-size),
[`src`{#hidden-state-(type=hidden):attr-input-src}](#attr-input-src),
[`step`{#hidden-state-(type=hidden):attr-input-step}](#attr-input-step),
and
[`width`{#hidden-state-(type=hidden):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#hidden-state-(type=hidden):do-not-apply-2} to the
element:
[`checked`{#hidden-state-(type=hidden):dom-input-checked}](#dom-input-checked),
[`files`{#hidden-state-(type=hidden):dom-input-files}](#dom-input-files),
[`list`{#hidden-state-(type=hidden):dom-input-list}](#dom-input-list),
[`selectionStart`{#hidden-state-(type=hidden):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#hidden-state-(type=hidden):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#hidden-state-(type=hidden):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#hidden-state-(type=hidden):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#hidden-state-(type=hidden):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#hidden-state-(type=hidden):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#hidden-state-(type=hidden):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#hidden-state-(type=hidden):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#hidden-state-(type=hidden):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#hidden-state-(type=hidden):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`input`{#hidden-state-(type=hidden):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#hidden-state-(type=hidden):event-change}](indices.html#event-change)
events [do not
apply](#do-not-apply){#hidden-state-(type=hidden):do-not-apply-3}.
:::

###### [4.10.5.1.2]{.secno} [Text]{.dfn dfn-for="input" dfn-type="element-state"} (`type=text`) state and [Search]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=search`)[](#text-(type=text)-state-and-search-state-(type=search)){.self-link} {#text-(type=text)-state-and-search-state-(type=search)}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/search](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/search "<input> elements of type search are text fields designed for the user to enter search queries into. These are functionally identical to text inputs, but may be styled differently by the user agent.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

:::: feature
[Element/input/text](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/text "<input> elements of type text create basic single-line text fields.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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
:::::::

When an
[`input`{#text-(type=text)-state-and-search-state-(type=search):the-input-element}](#the-input-element)
element\'s
[`type`{#text-(type=text)-state-and-search-state-(type=search):attr-input-type}](#attr-input-type)
attribute is in the
[Text](#text-(type=text)-state-and-search-state-(type=search)){#text-(type=text)-state-and-search-state-(type=search):text-(type=text)-state-and-search-state-(type=search)}
state or the
[Search](#text-(type=text)-state-and-search-state-(type=search)){#text-(type=text)-state-and-search-state-(type=search):text-(type=text)-state-and-search-state-(type=search)-2}
state, the rules in this section apply.

The
[`input`{#text-(type=text)-state-and-search-state-(type=search):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#text-(type=text)-state-and-search-state-(type=search):represents}
a one line plain text edit control for the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#text-(type=text)-state-and-search-state-(type=search):concept-fe-value}.

The difference between the
[Text](#text-(type=text)-state-and-search-state-(type=search)){#text-(type=text)-state-and-search-state-(type=search):text-(type=text)-state-and-search-state-(type=search)-3}
state and the
[Search](#text-(type=text)-state-and-search-state-(type=search)){#text-(type=text)-state-and-search-state-(type=search):text-(type=text)-state-and-search-state-(type=search)-4}
state is primarily stylistic: on platforms where search controls are
distinguished from regular text controls, the
[Search](#text-(type=text)-state-and-search-state-(type=search)){#text-(type=text)-state-and-search-state-(type=search):text-(type=text)-state-and-search-state-(type=search)-5}
state might result in an appearance consistent with the platform\'s
search controls rather than appearing like a regular text control.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, its
[value](form-control-infrastructure.html#concept-fe-value){#text-(type=text)-state-and-search-state-(type=search):concept-fe-value-2}
should be editable by the user. User agents must not allow users to
insert U+000A LINE FEED (LF) or U+000D CARRIAGE RETURN (CR) characters
into the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#text-(type=text)-state-and-search-state-(type=search):concept-fe-value-3}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the writing direction of the
element, setting it either to a left-to-right writing direction or a
right-to-left writing direction. If the user does so, the user agent
must then run the following steps:

1.  Set the element\'s
    [`dir`{#text-(type=text)-state-and-search-state-(type=search):attr-dir}](dom.html#attr-dir)
    attribute to
    \"[`ltr`{#text-(type=text)-state-and-search-state-(type=search):attr-dir-ltr}](dom.html#attr-dir-ltr)\"
    if the user selected a left-to-right writing direction, and
    \"[`rtl`{#text-(type=text)-state-and-search-state-(type=search):attr-dir-rtl}](dom.html#attr-dir-rtl)\"
    if the user selected a right-to-left writing direction.

2.  [Queue an element
    task](webappapis.html#queue-an-element-task){#text-(type=text)-state-and-search-state-(type=search):queue-an-element-task}
    on the [user interaction task
    source](webappapis.html#user-interaction-task-source){#text-(type=text)-state-and-search-state-(type=search):user-interaction-task-source}
    given the element to [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#text-(type=text)-state-and-search-state-(type=search):concept-event-fire
    x-internal="concept-event-fire"} named
    [`input`{#text-(type=text)-state-and-search-state-(type=search):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    at the element, with the
    [`bubbles`{#text-(type=text)-state-and-search-state-(type=search):dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    and
    [`composed`{#text-(type=text)-state-and-search-state-(type=search):dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
    attributes initialized to true.

The
[`value`{#text-(type=text)-state-and-search-state-(type=search):attr-input-value}](#attr-input-value)
attribute, if specified, must have a value that contains no U+000A LINE
FEED (LF) or U+000D CARRIAGE RETURN (CR) characters.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#text-(type=text)-state-and-search-state-(type=search):value-sanitization-algorithm}
is as follows**: [Strip
newlines](https://infra.spec.whatwg.org/#strip-newlines){#text-(type=text)-state-and-search-state-(type=search):strip-newlines
x-internal="strip-newlines"} from the
[value](form-control-infrastructure.html#concept-fe-value){#text-(type=text)-state-and-search-state-(type=search):concept-fe-value-4}.

::: bookkeeping
The following common
[`input`{#text-(type=text)-state-and-search-state-(type=search):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#text-(type=text)-state-and-search-state-(type=search):concept-input-apply}
to the element:
[`autocomplete`{#text-(type=text)-state-and-search-state-(type=search):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`dirname`{#text-(type=text)-state-and-search-state-(type=search):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`list`{#text-(type=text)-state-and-search-state-(type=search):attr-input-list}](#attr-input-list),
[`maxlength`{#text-(type=text)-state-and-search-state-(type=search):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#text-(type=text)-state-and-search-state-(type=search):attr-input-minlength}](#attr-input-minlength),
[`pattern`{#text-(type=text)-state-and-search-state-(type=search):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#text-(type=text)-state-and-search-state-(type=search):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#text-(type=text)-state-and-search-state-(type=search):attr-input-readonly}](#attr-input-readonly),
[`required`{#text-(type=text)-state-and-search-state-(type=search):attr-input-required}](#attr-input-required),
and
[`size`{#text-(type=text)-state-and-search-state-(type=search):attr-input-size}](#attr-input-size)
content attributes;
[`list`{#text-(type=text)-state-and-search-state-(type=search):dom-input-list}](#dom-input-list),
[`selectionStart`{#text-(type=text)-state-and-search-state-(type=search):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#text-(type=text)-state-and-search-state-(type=search):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#text-(type=text)-state-and-search-state-(type=search):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
and
[`value`{#text-(type=text)-state-and-search-state-(type=search):dom-input-value}](#dom-input-value)
IDL attributes;
[`select()`{#text-(type=text)-state-and-search-state-(type=search):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#text-(type=text)-state-and-search-state-(type=search):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#text-(type=text)-state-and-search-state-(type=search):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.

The
[`value`{#text-(type=text)-state-and-search-state-(type=search):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#text-(type=text)-state-and-search-state-(type=search):dom-input-value-value}.

The
[`input`{#text-(type=text)-state-and-search-state-(type=search):event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#text-(type=text)-state-and-search-state-(type=search):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#text-(type=text)-state-and-search-state-(type=search):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#text-(type=text)-state-and-search-state-(type=search):do-not-apply}
to the element:
[`accept`{#text-(type=text)-state-and-search-state-(type=search):attr-input-accept}](#attr-input-accept),
[`alpha`{#text-(type=text)-state-and-search-state-(type=search):attr-input-alpha}](#attr-input-alpha),
[`alt`{#text-(type=text)-state-and-search-state-(type=search):attr-input-alt}](#attr-input-alt),
[`checked`{#text-(type=text)-state-and-search-state-(type=search):attr-input-checked}](#attr-input-checked),
[`colorspace`{#text-(type=text)-state-and-search-state-(type=search):attr-input-colorspace}](#attr-input-colorspace),
[`formaction`{#text-(type=text)-state-and-search-state-(type=search):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#text-(type=text)-state-and-search-state-(type=search):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#text-(type=text)-state-and-search-state-(type=search):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#text-(type=text)-state-and-search-state-(type=search):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#text-(type=text)-state-and-search-state-(type=search):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#text-(type=text)-state-and-search-state-(type=search):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`max`{#text-(type=text)-state-and-search-state-(type=search):attr-input-max}](#attr-input-max),
[`min`{#text-(type=text)-state-and-search-state-(type=search):attr-input-min}](#attr-input-min),
[`multiple`{#text-(type=text)-state-and-search-state-(type=search):attr-input-multiple}](#attr-input-multiple),
[`popovertarget`{#text-(type=text)-state-and-search-state-(type=search):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#text-(type=text)-state-and-search-state-(type=search):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`src`{#text-(type=text)-state-and-search-state-(type=search):attr-input-src}](#attr-input-src),
[`step`{#text-(type=text)-state-and-search-state-(type=search):attr-input-step}](#attr-input-step),
and
[`width`{#text-(type=text)-state-and-search-state-(type=search):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#text-(type=text)-state-and-search-state-(type=search):do-not-apply-2}
to the element:
[`checked`{#text-(type=text)-state-and-search-state-(type=search):dom-input-checked}](#dom-input-checked),
[`files`{#text-(type=text)-state-and-search-state-(type=search):dom-input-files}](#dom-input-files),
[`valueAsDate`{#text-(type=text)-state-and-search-state-(type=search):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#text-(type=text)-state-and-search-state-(type=search):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`stepDown()`{#text-(type=text)-state-and-search-state-(type=search):dom-input-stepdown}](#dom-input-stepdown)
and
[`stepUp()`{#text-(type=text)-state-and-search-state-(type=search):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.3]{.secno} [Telephone]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=tel`)[](#telephone-state-(type=tel)){.self-link} {#telephone-state-(type=tel)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/tel](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/tel "<input> elements of type tel are used to let the user enter and edit a telephone number. Unlike <input type="email"> and <input type="url">, the input value is not automatically validated to a particular format before the form can be submitted, because formats for telephone numbers vary so much around the world.")

Support in all current engines.

::: support
[FirefoxYes]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android18+]{.chrome_android .yes}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#telephone-state-(type=tel):the-input-element}](#the-input-element)
element\'s
[`type`{#telephone-state-(type=tel):attr-input-type}](#attr-input-type)
attribute is in the
[Telephone](#telephone-state-(type=tel)){#telephone-state-(type=tel):telephone-state-(type=tel)}
state, the rules in this section apply.

The
[`input`{#telephone-state-(type=tel):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#telephone-state-(type=tel):represents}
a control for editing a telephone number given in the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#telephone-state-(type=tel):concept-fe-value}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, its
[value](form-control-infrastructure.html#concept-fe-value){#telephone-state-(type=tel):concept-fe-value-2}
should be editable by the user. User agents may change the spacing and,
with care, the punctuation of
[values](form-control-infrastructure.html#concept-fe-value){#telephone-state-(type=tel):concept-fe-value-3}
that the user enters. User agents must not allow users to insert U+000A
LINE FEED (LF) or U+000D CARRIAGE RETURN (CR) characters into the
element\'s
[value](form-control-infrastructure.html#concept-fe-value){#telephone-state-(type=tel):concept-fe-value-4}.

The
[`value`{#telephone-state-(type=tel):attr-input-value}](#attr-input-value)
attribute, if specified, must have a value that contains no U+000A LINE
FEED (LF) or U+000D CARRIAGE RETURN (CR) characters.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#telephone-state-(type=tel):value-sanitization-algorithm}
is as follows**: [Strip
newlines](https://infra.spec.whatwg.org/#strip-newlines){#telephone-state-(type=tel):strip-newlines
x-internal="strip-newlines"} from the
[value](form-control-infrastructure.html#concept-fe-value){#telephone-state-(type=tel):concept-fe-value-5}.

Unlike the
[URL](#url-state-(type=url)){#telephone-state-(type=tel):url-state-(type=url)}
and
[Email](#email-state-(type=email)){#telephone-state-(type=tel):email-state-(type=email)}
types, the
[Telephone](#telephone-state-(type=tel)){#telephone-state-(type=tel):telephone-state-(type=tel)-2}
type does not enforce a particular syntax. This is intentional; in
practice, telephone number fields tend to be free-form fields, because
there are a wide variety of valid phone numbers. Systems that need to
enforce a particular format are encouraged to use the
[`pattern`{#telephone-state-(type=tel):attr-input-pattern}](#attr-input-pattern)
attribute or the
[`setCustomValidity()`{#telephone-state-(type=tel):dom-cva-setcustomvalidity}](form-control-infrastructure.html#dom-cva-setcustomvalidity)
method to hook into the client-side validation mechanism.

::: bookkeeping
The following common
[`input`{#telephone-state-(type=tel):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#telephone-state-(type=tel):concept-input-apply}
to the element:
[`autocomplete`{#telephone-state-(type=tel):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`dirname`{#telephone-state-(type=tel):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`list`{#telephone-state-(type=tel):attr-input-list}](#attr-input-list),
[`maxlength`{#telephone-state-(type=tel):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#telephone-state-(type=tel):attr-input-minlength}](#attr-input-minlength),
[`pattern`{#telephone-state-(type=tel):attr-input-pattern-2}](#attr-input-pattern),
[`placeholder`{#telephone-state-(type=tel):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#telephone-state-(type=tel):attr-input-readonly}](#attr-input-readonly),
[`required`{#telephone-state-(type=tel):attr-input-required}](#attr-input-required),
and
[`size`{#telephone-state-(type=tel):attr-input-size}](#attr-input-size)
content attributes;
[`list`{#telephone-state-(type=tel):dom-input-list}](#dom-input-list),
[`selectionStart`{#telephone-state-(type=tel):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#telephone-state-(type=tel):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#telephone-state-(type=tel):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
and
[`value`{#telephone-state-(type=tel):dom-input-value}](#dom-input-value)
IDL attributes;
[`select()`{#telephone-state-(type=tel):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#telephone-state-(type=tel):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#telephone-state-(type=tel):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.

The
[`value`{#telephone-state-(type=tel):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#telephone-state-(type=tel):dom-input-value-value}.

The
[`input`{#telephone-state-(type=tel):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#telephone-state-(type=tel):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#telephone-state-(type=tel):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#telephone-state-(type=tel):do-not-apply} to the
element:
[`accept`{#telephone-state-(type=tel):attr-input-accept}](#attr-input-accept),
[`alpha`{#telephone-state-(type=tel):attr-input-alpha}](#attr-input-alpha),
[`alt`{#telephone-state-(type=tel):attr-input-alt}](#attr-input-alt),
[`checked`{#telephone-state-(type=tel):attr-input-checked}](#attr-input-checked),
[`colorspace`{#telephone-state-(type=tel):attr-input-colorspace}](#attr-input-colorspace),
[`formaction`{#telephone-state-(type=tel):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#telephone-state-(type=tel):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#telephone-state-(type=tel):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#telephone-state-(type=tel):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#telephone-state-(type=tel):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#telephone-state-(type=tel):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`max`{#telephone-state-(type=tel):attr-input-max}](#attr-input-max),
[`min`{#telephone-state-(type=tel):attr-input-min}](#attr-input-min),
[`multiple`{#telephone-state-(type=tel):attr-input-multiple}](#attr-input-multiple),
[`popovertarget`{#telephone-state-(type=tel):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#telephone-state-(type=tel):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`src`{#telephone-state-(type=tel):attr-input-src}](#attr-input-src),
[`step`{#telephone-state-(type=tel):attr-input-step}](#attr-input-step),
and
[`width`{#telephone-state-(type=tel):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#telephone-state-(type=tel):do-not-apply-2} to the
element:
[`checked`{#telephone-state-(type=tel):dom-input-checked}](#dom-input-checked),
[`files`{#telephone-state-(type=tel):dom-input-files}](#dom-input-files),
[`valueAsDate`{#telephone-state-(type=tel):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#telephone-state-(type=tel):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`stepDown()`{#telephone-state-(type=tel):dom-input-stepdown}](#dom-input-stepdown)
and
[`stepUp()`{#telephone-state-(type=tel):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.4]{.secno} [URL]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=url`)[](#url-state-(type=url)){.self-link} {#url-state-(type=url)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/url](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/url "<input> elements of type url are used to let the user enter and edit a URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

When an
[`input`{#url-state-(type=url):the-input-element}](#the-input-element)
element\'s
[`type`{#url-state-(type=url):attr-input-type}](#attr-input-type)
attribute is in the
[URL](#url-state-(type=url)){#url-state-(type=url):url-state-(type=url)}
state, the rules in this section apply.

The
[`input`{#url-state-(type=url):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#url-state-(type=url):represents} a
control for editing a single [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#url-state-(type=url):absolute-url
x-internal="absolute-url"} given in the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the URL represented by its
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-2}.
User agents may allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-3}
to a string that is not a
[valid](https://url.spec.whatwg.org/#valid-url-string){#url-state-(type=url):valid-url-string
x-internal="valid-url-string"} [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#url-state-(type=url):absolute-url-2
x-internal="absolute-url"}, but may also or instead automatically escape
characters entered by the user so that the
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-4}
is always a
[valid](https://url.spec.whatwg.org/#valid-url-string){#url-state-(type=url):valid-url-string-2
x-internal="valid-url-string"} [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#url-state-(type=url):absolute-url-3
x-internal="absolute-url"} (even if that isn\'t the actual value seen
and edited by the user in the interface). User agents should allow the
user to set the
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-5}
to the empty string. User agents must not allow users to insert U+000A
LINE FEED (LF) or U+000D CARRIAGE RETURN (CR) characters into the
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-6}.

The [`value`{#url-state-(type=url):attr-input-value}](#attr-input-value)
attribute, if specified and not empty, must have a value that is a
[valid URL potentially surrounded by
spaces](urls-and-fetching.html#valid-url-potentially-surrounded-by-spaces){#url-state-(type=url):valid-url-potentially-surrounded-by-spaces}
that is also an [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#url-state-(type=url):absolute-url-4
x-internal="absolute-url"}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#url-state-(type=url):value-sanitization-algorithm}
is as follows**: [Strip
newlines](https://infra.spec.whatwg.org/#strip-newlines){#url-state-(type=url):strip-newlines
x-internal="strip-newlines"} from the
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-7},
then [strip leading and trailing ASCII
whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#url-state-(type=url):strip-leading-and-trailing-ascii-whitespace
x-internal="strip-leading-and-trailing-ascii-whitespace"} from the
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-8}.

**Constraint validation**: While the
[value](form-control-infrastructure.html#concept-fe-value){#url-state-(type=url):concept-fe-value-9}
of the element is neither the empty string nor a
[valid](https://url.spec.whatwg.org/#valid-url-string){#url-state-(type=url):valid-url-string-3
x-internal="valid-url-string"} [absolute
URL](https://url.spec.whatwg.org/#syntax-url-absolute){#url-state-(type=url):absolute-url-5
x-internal="absolute-url"}, the element is [suffering from a type
mismatch](form-control-infrastructure.html#suffering-from-a-type-mismatch){#url-state-(type=url):suffering-from-a-type-mismatch}.

::: bookkeeping
The following common
[`input`{#url-state-(type=url):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#url-state-(type=url):concept-input-apply}
to the element:
[`autocomplete`{#url-state-(type=url):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`dirname`{#url-state-(type=url):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`list`{#url-state-(type=url):attr-input-list}](#attr-input-list),
[`maxlength`{#url-state-(type=url):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#url-state-(type=url):attr-input-minlength}](#attr-input-minlength),
[`pattern`{#url-state-(type=url):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#url-state-(type=url):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#url-state-(type=url):attr-input-readonly}](#attr-input-readonly),
[`required`{#url-state-(type=url):attr-input-required}](#attr-input-required),
and [`size`{#url-state-(type=url):attr-input-size}](#attr-input-size)
content attributes;
[`list`{#url-state-(type=url):dom-input-list}](#dom-input-list),
[`selectionStart`{#url-state-(type=url):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#url-state-(type=url):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#url-state-(type=url):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
and [`value`{#url-state-(type=url):dom-input-value}](#dom-input-value)
IDL attributes;
[`select()`{#url-state-(type=url):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#url-state-(type=url):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#url-state-(type=url):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.

The [`value`{#url-state-(type=url):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#url-state-(type=url):dom-input-value-value}.

The
[`input`{#url-state-(type=url):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#url-state-(type=url):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#url-state-(type=url):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#url-state-(type=url):do-not-apply} to the
element:
[`accept`{#url-state-(type=url):attr-input-accept}](#attr-input-accept),
[`alpha`{#url-state-(type=url):attr-input-alpha}](#attr-input-alpha),
[`alt`{#url-state-(type=url):attr-input-alt}](#attr-input-alt),
[`checked`{#url-state-(type=url):attr-input-checked}](#attr-input-checked),
[`colorspace`{#url-state-(type=url):attr-input-colorspace}](#attr-input-colorspace),
[`formaction`{#url-state-(type=url):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#url-state-(type=url):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#url-state-(type=url):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#url-state-(type=url):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#url-state-(type=url):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#url-state-(type=url):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`max`{#url-state-(type=url):attr-input-max}](#attr-input-max),
[`min`{#url-state-(type=url):attr-input-min}](#attr-input-min),
[`multiple`{#url-state-(type=url):attr-input-multiple}](#attr-input-multiple),
[`popovertarget`{#url-state-(type=url):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#url-state-(type=url):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`src`{#url-state-(type=url):attr-input-src}](#attr-input-src),
[`step`{#url-state-(type=url):attr-input-step}](#attr-input-step), and
[`width`{#url-state-(type=url):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#url-state-(type=url):do-not-apply-2} to the
element:
[`checked`{#url-state-(type=url):dom-input-checked}](#dom-input-checked),
[`files`{#url-state-(type=url):dom-input-files}](#dom-input-files),
[`valueAsDate`{#url-state-(type=url):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#url-state-(type=url):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`stepDown()`{#url-state-(type=url):dom-input-stepdown}](#dom-input-stepdown)
and
[`stepUp()`{#url-state-(type=url):dom-input-stepup}](#dom-input-stepup)
methods.
:::

::: example
If a document contained the following markup:

``` html
<input type="url" name="location" list="urls">
<datalist id="urls">
 <option label="MIME: Format of Internet Message Bodies" value="https://www.rfc-editor.org/rfc/rfc2045">
 <option label="HTML" value="https://html.spec.whatwg.org/">
 <option label="DOM" value="https://dom.spec.whatwg.org/">
 <option label="Fullscreen" value="https://fullscreen.spec.whatwg.org/">
 <option label="Media Session" value="https://mediasession.spec.whatwg.org/">
 <option label="The Single UNIX Specification, Version 3" value="http://www.unix.org/version3/">
</datalist>
```

\...and the user had typed \"[spec.w]{.kbd}\", and the user agent had
also found that the user had visited
`https://url.spec.whatwg.org/#url-parsing` and
`https://streams.spec.whatwg.org/` in the recent past, then the
rendering might look like this:

![A text box with an icon on the left followed by the text \"spec.w\"
and a cursor, with a drop down button on the right hand side; with,
below, a drop down box containing a list of six URLs on the left, with
the first four having grayed out labels on the right; and a scrollbar to
the right of the drop down box, indicating further values are
available.](/images/sample-url.svg){.darkmode-aware width="480"
height="150"}

The first four URLs in this sample consist of the four URLs in the
author-specified list that match the text the user has entered, sorted
in some
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#url-state-(type=url):implementation-defined
x-internal="implementation-defined"} manner (maybe by how frequently the
user refers to those URLs). Note how the UA is using the knowledge that
the values are URLs to allow the user to omit the scheme part and
perform intelligent matching on the domain name.

The last two URLs (and probably many more, given the scrollbar\'s
indications of more values being available) are the matches from the
user agent\'s session history data. This data is not made available to
the page DOM. In this particular case, the UA has no titles to provide
for those values.
:::

###### [4.10.5.1.5]{.secno} [Email]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=email`)[]{#e-mail-state-(type=email)}[](#email-state-(type=email)){.self-link} {#email-state-(type=email)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/email](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/email "<input> elements of type email are used to let the user enter and edit an email address, or, if the multiple attribute is specified, a list of email addresses.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#email-state-(type=email):the-input-element}](#the-input-element)
element\'s
[`type`{#email-state-(type=email):attr-input-type}](#attr-input-type)
attribute is in the
[Email](#email-state-(type=email)){#email-state-(type=email):email-state-(type=email)}
state, the rules in this section apply.

How the
[Email](#email-state-(type=email)){#email-state-(type=email):email-state-(type=email)-2}
state operates depends on whether the
[`multiple`{#email-state-(type=email):attr-input-multiple}](#attr-input-multiple)
attribute is specified or not.

When the [`multiple`{#email-state-(type=email):attr-input-multiple-2}](#attr-input-multiple) attribute is not specified on the element

:   The
    [`input`{#email-state-(type=email):the-input-element-2}](#the-input-element)
    element
    [represents](dom.html#represents){#email-state-(type=email):represents}
    a control for editing an email address given in the element\'s
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value}.

    If the element is
    *[mutable](form-control-infrastructure.html#concept-fe-mutable)*,
    the user agent should allow the user to change the email address
    represented by its
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-2}.
    User agents may allow the user to set the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-3}
    to a string that is not a [valid email
    address](#valid-e-mail-address){#email-state-(type=email):valid-e-mail-address}.
    The user agent should act in a manner consistent with expecting the
    user to provide a single email address. User agents should allow the
    user to set the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-4}
    to the empty string. User agents must not allow users to insert
    U+000A LINE FEED (LF) or U+000D CARRIAGE RETURN (CR) characters into
    the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-5}.
    User agents may transform the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-6}
    for display and editing; in particular, user agents should convert
    punycode in the domain labels of the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-7}
    to IDN in the display and vice versa.

    **Constraint validation**: While the user interface is representing
    input that the user agent cannot convert to punycode, the control is
    [suffering from bad
    input](form-control-infrastructure.html#suffering-from-bad-input){#email-state-(type=email):suffering-from-bad-input}.

    The
    [`value`{#email-state-(type=email):attr-input-value}](#attr-input-value)
    attribute, if specified and not empty, must have a value that is a
    single [valid email
    address](#valid-e-mail-address){#email-state-(type=email):valid-e-mail-address-2}.

    **The [value sanitization
    algorithm](#value-sanitization-algorithm){#email-state-(type=email):value-sanitization-algorithm}
    is as follows**: [Strip
    newlines](https://infra.spec.whatwg.org/#strip-newlines){#email-state-(type=email):strip-newlines
    x-internal="strip-newlines"} from the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-8},
    then [strip leading and trailing ASCII
    whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#email-state-(type=email):strip-leading-and-trailing-ascii-whitespace
    x-internal="strip-leading-and-trailing-ascii-whitespace"} from the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-9}.

    **Constraint validation**: While the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-10}
    of the element is neither the empty string nor a single [valid email
    address](#valid-e-mail-address){#email-state-(type=email):valid-e-mail-address-3},
    the element is [suffering from a type
    mismatch](form-control-infrastructure.html#suffering-from-a-type-mismatch){#email-state-(type=email):suffering-from-a-type-mismatch}.

When the [`multiple`{#email-state-(type=email):attr-input-multiple-3}](#attr-input-multiple) attribute *is* specified on the element

:   The
    [`input`{#email-state-(type=email):the-input-element-3}](#the-input-element)
    element
    [represents](dom.html#represents){#email-state-(type=email):represents-2}
    a control for adding, removing, and editing the email addresses
    given in the element\'s
    [value*s*](form-control-infrastructure.html#concept-fe-values){#email-state-(type=email):concept-fe-values}.

    If the element is
    *[mutable](form-control-infrastructure.html#concept-fe-mutable)*,
    the user agent should allow the user to add, remove, and edit the
    email addresses represented by its
    [values](form-control-infrastructure.html#concept-fe-values){#email-state-(type=email):concept-fe-values-2}.
    User agents may allow the user to set any individual value in the
    list of
    [value*s*](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-11}
    to a string that is not a [valid email
    address](#valid-e-mail-address){#email-state-(type=email):valid-e-mail-address-4},
    but must not allow users to set any individual value to a string
    containing U+002C COMMA (,), U+000A LINE FEED (LF), or U+000D
    CARRIAGE RETURN (CR) characters. User agents should allow the user
    to remove all the addresses in the element\'s
    [values](form-control-infrastructure.html#concept-fe-values){#email-state-(type=email):concept-fe-values-3}.
    User agents may transform the
    [values](form-control-infrastructure.html#concept-fe-values){#email-state-(type=email):concept-fe-values-4}
    for display and editing; in particular, user agents should convert
    punycode in the domain labels of the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-12}
    to IDN in the display and vice versa.

    **Constraint validation**: While the user interface describes a
    situation where an individual value contains a U+002C COMMA (,) or
    is representing input that the user agent cannot convert to
    punycode, the control is [suffering from bad
    input](form-control-infrastructure.html#suffering-from-bad-input){#email-state-(type=email):suffering-from-bad-input-2}.

    Whenever the user changes the element\'s
    [value*s*](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-13},
    the user agent must run the following steps:

    1.  Let `latest values`{.variable} be a copy of the element\'s
        [value*s*](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-14}.

    2.  [Strip leading and trailing ASCII
        whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#email-state-(type=email):strip-leading-and-trailing-ascii-whitespace-2
        x-internal="strip-leading-and-trailing-ascii-whitespace"} from
        each value in `latest values`{.variable}.

    3.  Let the element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-15}
        be the result of concatenating all the values in
        `latest values`{.variable}, separating each value from the next
        by a single U+002C COMMA character (,), maintaining the list\'s
        order.

    The
    [`value`{#email-state-(type=email):attr-input-value-2}](#attr-input-value)
    attribute, if specified, must have a value that is a [valid email
    address
    list](#valid-e-mail-address-list){#email-state-(type=email):valid-e-mail-address-list}.

    **The [value sanitization
    algorithm](#value-sanitization-algorithm){#email-state-(type=email):value-sanitization-algorithm-2}
    is as follows**:

    1.  [Split on
        commas](https://infra.spec.whatwg.org/#split-on-commas){#email-state-(type=email):split-a-string-on-commas
        x-internal="split-a-string-on-commas"} the element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-16},
        [strip leading and trailing ASCII
        whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#email-state-(type=email):strip-leading-and-trailing-ascii-whitespace-3
        x-internal="strip-leading-and-trailing-ascii-whitespace"} from
        each resulting token, if any, and let the element\'s
        [values](form-control-infrastructure.html#concept-fe-values){#email-state-(type=email):concept-fe-values-5}
        be the (possibly empty) resulting list of (possibly empty)
        tokens, maintaining the original order.

    2.  Set the element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-17}
        to the result of concatenating the element\'s
        [values](form-control-infrastructure.html#concept-fe-values){#email-state-(type=email):concept-fe-values-6},
        separating each value from the next by a single U+002C COMMA
        character (,), maintaining the list\'s order.

    **Constraint validation**: While the
    [value](form-control-infrastructure.html#concept-fe-value){#email-state-(type=email):concept-fe-value-18}
    of the element is not a [valid email address
    list](#valid-e-mail-address-list){#email-state-(type=email):valid-e-mail-address-list-2},
    the element is [suffering from a type
    mismatch](form-control-infrastructure.html#suffering-from-a-type-mismatch){#email-state-(type=email):suffering-from-a-type-mismatch-2}.

When the
[`multiple`{#email-state-(type=email):attr-input-multiple-4}](#attr-input-multiple)
attribute is set or removed, the user agent must run the [value
sanitization
algorithm](#value-sanitization-algorithm){#email-state-(type=email):value-sanitization-algorithm-3}.

A [valid email address]{#valid-e-mail-address .dfn} is a string that
matches the `email` production of the following ABNF, the character set
for which is Unicode. This ABNF implements the extensions described in
RFC 1123. [\[ABNF\]](references.html#refsABNF)
[\[RFC5322\]](references.html#refsRFC5322)
[\[RFC1034\]](references.html#refsRFC1034)
[\[RFC1123\]](references.html#refsRFC1123)

``` abnf
email         = 1*( atext / "." ) "@" label *( "." label )
label         = let-dig [ [ ldh-str ] let-dig ]  ; limited to a length of 63 characters by RFC 1034 section 3.5
atext         = < as defined in RFC 5322 section 3.2.3 >
let-dig       = < as defined in RFC 1034 section 3.5 >
ldh-str       = < as defined in RFC 1034 section 3.5 >
```

This requirement is a [willful
violation](https://infra.spec.whatwg.org/#willful-violation){#email-state-(type=email):willful-violation
x-internal="willful-violation"} of RFC 5322, which defines a syntax for
email addresses that is simultaneously too strict (before the \"@\"
character), too vague (after the \"@\" character), and too lax (allowing
comments, whitespace characters, and quoted strings in manners
unfamiliar to most users) to be of practical use here.

::: note
The following JavaScript- and Perl-compatible regular expression is an
implementation of the above definition.

    /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/
:::

A [valid email address list]{#valid-e-mail-address-list .dfn} is a [set
of comma-separated
tokens](common-microsyntaxes.html#set-of-comma-separated-tokens){#email-state-(type=email):set-of-comma-separated-tokens},
where each token is itself a [valid email
address](#valid-e-mail-address){#email-state-(type=email):valid-e-mail-address-5}.
To obtain the list of tokens from a [valid email address
list](#valid-e-mail-address-list){#email-state-(type=email):valid-e-mail-address-list-3},
an implementation must [split the string on
commas](https://infra.spec.whatwg.org/#split-on-commas){#email-state-(type=email):split-a-string-on-commas-2
x-internal="split-a-string-on-commas"}.

::: bookkeeping
The following common
[`input`{#email-state-(type=email):the-input-element-4}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#email-state-(type=email):concept-input-apply}
to the element:
[`autocomplete`{#email-state-(type=email):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`dirname`{#email-state-(type=email):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`list`{#email-state-(type=email):attr-input-list}](#attr-input-list),
[`maxlength`{#email-state-(type=email):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#email-state-(type=email):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#email-state-(type=email):attr-input-multiple-5}](#attr-input-multiple),
[`pattern`{#email-state-(type=email):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#email-state-(type=email):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#email-state-(type=email):attr-input-readonly}](#attr-input-readonly),
[`required`{#email-state-(type=email):attr-input-required}](#attr-input-required),
and
[`size`{#email-state-(type=email):attr-input-size}](#attr-input-size)
content attributes;
[`list`{#email-state-(type=email):dom-input-list}](#dom-input-list) and
[`value`{#email-state-(type=email):dom-input-value}](#dom-input-value)
IDL attributes;
[`select()`{#email-state-(type=email):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select)
method.

The
[`value`{#email-state-(type=email):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#email-state-(type=email):dom-input-value-value}.

The
[`input`{#email-state-(type=email):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#email-state-(type=email):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#email-state-(type=email):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#email-state-(type=email):do-not-apply} to the
element:
[`accept`{#email-state-(type=email):attr-input-accept}](#attr-input-accept),
[`alpha`{#email-state-(type=email):attr-input-alpha}](#attr-input-alpha),
[`alt`{#email-state-(type=email):attr-input-alt}](#attr-input-alt),
[`checked`{#email-state-(type=email):attr-input-checked}](#attr-input-checked),
[`colorspace`{#email-state-(type=email):attr-input-colorspace}](#attr-input-colorspace),
[`formaction`{#email-state-(type=email):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#email-state-(type=email):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#email-state-(type=email):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#email-state-(type=email):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#email-state-(type=email):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#email-state-(type=email):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`max`{#email-state-(type=email):attr-input-max}](#attr-input-max),
[`min`{#email-state-(type=email):attr-input-min}](#attr-input-min),
[`popovertarget`{#email-state-(type=email):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#email-state-(type=email):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`src`{#email-state-(type=email):attr-input-src}](#attr-input-src),
[`step`{#email-state-(type=email):attr-input-step}](#attr-input-step),
and
[`width`{#email-state-(type=email):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#email-state-(type=email):do-not-apply-2} to the
element:
[`checked`{#email-state-(type=email):dom-input-checked}](#dom-input-checked),
[`files`{#email-state-(type=email):dom-input-files}](#dom-input-files),
[`selectionStart`{#email-state-(type=email):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#email-state-(type=email):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#email-state-(type=email):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#email-state-(type=email):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#email-state-(type=email):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`setRangeText()`{#email-state-(type=email):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#email-state-(type=email):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#email-state-(type=email):dom-input-stepdown}](#dom-input-stepdown)
and
[`stepUp()`{#email-state-(type=email):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.6]{.secno} [Password]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=password`)[](#password-state-(type=password)){.self-link} {#password-state-(type=password)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/password](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/password "<input> elements of type password provide a way for the user to securely enter a password.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera2+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer2+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

When an
[`input`{#password-state-(type=password):the-input-element}](#the-input-element)
element\'s
[`type`{#password-state-(type=password):attr-input-type}](#attr-input-type)
attribute is in the
[Password](#password-state-(type=password)){#password-state-(type=password):password-state-(type=password)}
state, the rules in this section apply.

The
[`input`{#password-state-(type=password):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#password-state-(type=password):represents}
a one line plain text edit control for the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#password-state-(type=password):concept-fe-value}.
The user agent should obscure the value so that people other than the
user cannot see it.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, its
[value](form-control-infrastructure.html#concept-fe-value){#password-state-(type=password):concept-fe-value-2}
should be editable by the user. User agents must not allow users to
insert U+000A LINE FEED (LF) or U+000D CARRIAGE RETURN (CR) characters
into the
[value](form-control-infrastructure.html#concept-fe-value){#password-state-(type=password):concept-fe-value-3}.

The
[`value`{#password-state-(type=password):attr-input-value}](#attr-input-value)
attribute, if specified, must have a value that contains no U+000A LINE
FEED (LF) or U+000D CARRIAGE RETURN (CR) characters.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#password-state-(type=password):value-sanitization-algorithm}
is as follows**: [Strip
newlines](https://infra.spec.whatwg.org/#strip-newlines){#password-state-(type=password):strip-newlines
x-internal="strip-newlines"} from the
[value](form-control-infrastructure.html#concept-fe-value){#password-state-(type=password):concept-fe-value-4}.

::: bookkeeping
The following common
[`input`{#password-state-(type=password):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#password-state-(type=password):concept-input-apply}
to the element:
[`autocomplete`{#password-state-(type=password):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`dirname`{#password-state-(type=password):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`maxlength`{#password-state-(type=password):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#password-state-(type=password):attr-input-minlength}](#attr-input-minlength),
[`pattern`{#password-state-(type=password):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#password-state-(type=password):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#password-state-(type=password):attr-input-readonly}](#attr-input-readonly),
[`required`{#password-state-(type=password):attr-input-required}](#attr-input-required),
and
[`size`{#password-state-(type=password):attr-input-size}](#attr-input-size)
content attributes;
[`selectionStart`{#password-state-(type=password):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#password-state-(type=password):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#password-state-(type=password):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
and
[`value`{#password-state-(type=password):dom-input-value}](#dom-input-value)
IDL attributes;
[`select()`{#password-state-(type=password):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#password-state-(type=password):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#password-state-(type=password):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.

The
[`value`{#password-state-(type=password):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#password-state-(type=password):dom-input-value-value}.

The
[`input`{#password-state-(type=password):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#password-state-(type=password):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#password-state-(type=password):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#password-state-(type=password):do-not-apply} to
the element:
[`accept`{#password-state-(type=password):attr-input-accept}](#attr-input-accept),
[`alpha`{#password-state-(type=password):attr-input-alpha}](#attr-input-alpha),
[`alt`{#password-state-(type=password):attr-input-alt}](#attr-input-alt),
[`checked`{#password-state-(type=password):attr-input-checked}](#attr-input-checked),
[`colorspace`{#password-state-(type=password):attr-input-colorspace}](#attr-input-colorspace),
[`formaction`{#password-state-(type=password):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#password-state-(type=password):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#password-state-(type=password):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#password-state-(type=password):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#password-state-(type=password):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#password-state-(type=password):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#password-state-(type=password):attr-input-list}](#attr-input-list),
[`max`{#password-state-(type=password):attr-input-max}](#attr-input-max),
[`min`{#password-state-(type=password):attr-input-min}](#attr-input-min),
[`multiple`{#password-state-(type=password):attr-input-multiple}](#attr-input-multiple),
[`popovertarget`{#password-state-(type=password):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#password-state-(type=password):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`src`{#password-state-(type=password):attr-input-src}](#attr-input-src),
[`step`{#password-state-(type=password):attr-input-step}](#attr-input-step),
and
[`width`{#password-state-(type=password):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#password-state-(type=password):do-not-apply-2} to
the element:
[`checked`{#password-state-(type=password):dom-input-checked}](#dom-input-checked),
[`files`{#password-state-(type=password):dom-input-files}](#dom-input-files),
[`list`{#password-state-(type=password):dom-input-list}](#dom-input-list),
[`valueAsDate`{#password-state-(type=password):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#password-state-(type=password):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`stepDown()`{#password-state-(type=password):dom-input-stepdown}](#dom-input-stepdown)
and
[`stepUp()`{#password-state-(type=password):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.7]{.secno} [Date]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=date`)[](#date-state-(type=date)){.self-link} {#date-state-(type=date)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/date](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/date "<input> elements of type="date" create input fields that let the user enter a date, either with a textbox that validates the input or a special date picker interface.")

Support in all current engines.

::: support
[Firefox57+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome20+]{.chrome .yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#date-state-(type=date):the-input-element}](#the-input-element)
element\'s
[`type`{#date-state-(type=date):attr-input-type}](#attr-input-type)
attribute is in the
[Date](#date-state-(type=date)){#date-state-(type=date):date-state-(type=date)}
state, the rules in this section apply.

The
[`input`{#date-state-(type=date):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#date-state-(type=date):represents} a
control for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#date-state-(type=date):concept-fe-value}
to a string representing a specific
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date-2}
represented by its
[value](form-control-infrastructure.html#concept-fe-value){#date-state-(type=date):concept-fe-value-2},
as obtained by [parsing a
date](common-microsyntaxes.html#parse-a-date-string){#date-state-(type=date):parse-a-date-string}
from it. User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#date-state-(type=date):concept-fe-value-3}
to a non-empty string that is not a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string}.
If the user agent provides a user interface for selecting a
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date-3},
then the
[value](form-control-infrastructure.html#concept-fe-value){#date-state-(type=date):concept-fe-value-4}
must be set to a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-2}
representing the user\'s selection. User agents should allow the user to
set the
[value](form-control-infrastructure.html#concept-fe-value){#date-state-(type=date):concept-fe-value-5}
to the empty string.

**Constraint validation**: While the user interface describes input that
the user agent cannot convert to a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-3},
the control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#date-state-(type=date):suffering-from-bad-input}.

See the [introduction section](forms.html#input-author-notes) for a
discussion of the difference between the input format and submission
format for date, time, and number form controls, and the [implementation
notes](#input-impl-notes) regarding localization of form controls.

The
[`value`{#date-state-(type=date):attr-input-value}](#attr-input-value)
attribute, if specified and not empty, must have a value that is a
[valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-4}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#date-state-(type=date):value-sanitization-algorithm}
is as follows**: If the
[value](form-control-infrastructure.html#concept-fe-value){#date-state-(type=date):concept-fe-value-6}
of the element is not a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-5},
then set it to the empty string instead.

The [`min`{#date-state-(type=date):attr-input-min}](#attr-input-min)
attribute, if specified, must have a value that is a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-6}.
The [`max`{#date-state-(type=date):attr-input-max}](#attr-input-max)
attribute, if specified, must have a value that is a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-7}.

The [`step`{#date-state-(type=date):attr-input-step}](#attr-input-step)
attribute is expressed in days. The [step scale
factor](#concept-input-step-scale){#date-state-(type=date):concept-input-step-scale}
is 86,400,000 (which converts the days to milliseconds, as used in the
other algorithms). The [default
step](#concept-input-step-default){#date-state-(type=date):concept-input-step-default}
is 1 day.

When the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#date-state-(type=date):suffering-from-a-step-mismatch},
the user agent may round the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#date-state-(type=date):concept-fe-value-7}
to the nearest
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date-4}
for which the element would not [suffer from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#date-state-(type=date):suffering-from-a-step-mismatch-2}.

**The [algorithm to convert a string to a
number](#concept-input-value-string-number){#date-state-(type=date):concept-input-value-string-number},
given a string `input`{.variable}, is as follows**: If [parsing a
date](common-microsyntaxes.html#parse-a-date-string){#date-state-(type=date):parse-a-date-string-2}
from `input`{.variable} results in an error, then return an error;
otherwise, return the number of milliseconds elapsed from midnight UTC
on the morning of 1970-01-01 (the time represented by the value
\"`1970-01-01T00:00:00.0Z`\") to midnight UTC on the morning of the
parsed
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date-5},
ignoring leap seconds.

**The [algorithm to convert a number to a
string](#concept-input-value-number-string){#date-state-(type=date):concept-input-value-number-string},
given a number `input`{.variable}, is as follows**: Return a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-8}
that represents the
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date-6}
that, in UTC, is current `input`{.variable} milliseconds after midnight
UTC on the morning of 1970-01-01 (the time represented by the value
\"`1970-01-01T00:00:00.0Z`\").

**The [algorithm to convert a string to a `Date`
object](#concept-input-value-string-date){#date-state-(type=date):concept-input-value-string-date},
given a string `input`{.variable}, is as follows**: If [parsing a
date](common-microsyntaxes.html#parse-a-date-string){#date-state-(type=date):parse-a-date-string-3}
from `input`{.variable} results in an error, then return an error;
otherwise, return [a new `Date`
object](infrastructure.html#create-a-date-object){#date-state-(type=date):create-a-date-object}
representing midnight UTC on the morning of the parsed
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date-7}.

**The [algorithm to convert a `Date` object to a
string](#concept-input-value-date-string){#date-state-(type=date):concept-input-value-date-string},
given a
[`Date`{#date-state-(type=date):date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object `input`{.variable}, is as follows**: Return a [valid date
string](common-microsyntaxes.html#valid-date-string){#date-state-(type=date):valid-date-string-9}
that represents the
[date](common-microsyntaxes.html#concept-date){#date-state-(type=date):concept-date-8}
current at the time represented by `input`{.variable} in the UTC time
zone.

::: {#only-contemporary-times .note}
The
[Date](#date-state-(type=date)){#date-state-(type=date):date-state-(type=date)-2}
state (and other date- and time-related states described in subsequent
sections) is not intended for the entry of values for which a precise
date and time relative to the contemporary calendar cannot be
established. For example, it would be inappropriate for the entry of
times like \"one millisecond after the big bang\", \"the early part of
the Jurassic period\", or \"a winter around 250 BCE\".

For the input of dates before the introduction of the Gregorian
calendar, authors are encouraged to not use the
[Date](#date-state-(type=date)){#date-state-(type=date):date-state-(type=date)-3}
state (and the other date- and time-related states described in
subsequent sections), as user agents are not required to support
converting dates and times from earlier periods to the Gregorian
calendar, and asking users to do so manually puts an undue burden on
users. (This is complicated by the manner in which the Gregorian
calendar was phased in, which occurred at different times in different
countries, ranging from partway through the 16th century all the way to
early in the 20th.) Instead, authors are encouraged to provide
fine-grained input controls using the
[`select`{#date-state-(type=date):the-select-element}](form-elements.html#the-select-element)
element and
[`input`{#date-state-(type=date):the-input-element-3}](#the-input-element)
elements with the
[Number](#number-state-(type=number)){#date-state-(type=date):number-state-(type=number)}
state.
:::

::: bookkeeping
The following common
[`input`{#date-state-(type=date):the-input-element-4}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#date-state-(type=date):concept-input-apply}
to the element:
[`autocomplete`{#date-state-(type=date):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`list`{#date-state-(type=date):attr-input-list}](#attr-input-list),
[`max`{#date-state-(type=date):attr-input-max-2}](#attr-input-max),
[`min`{#date-state-(type=date):attr-input-min-2}](#attr-input-min),
[`readonly`{#date-state-(type=date):attr-input-readonly}](#attr-input-readonly),
[`required`{#date-state-(type=date):attr-input-required}](#attr-input-required),
and
[`step`{#date-state-(type=date):attr-input-step-2}](#attr-input-step)
content attributes;
[`list`{#date-state-(type=date):dom-input-list}](#dom-input-list),
[`value`{#date-state-(type=date):dom-input-value}](#dom-input-value),
[`valueAsDate`{#date-state-(type=date):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#date-state-(type=date):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#date-state-(type=date):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`stepDown()`{#date-state-(type=date):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#date-state-(type=date):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`value`{#date-state-(type=date):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#date-state-(type=date):dom-input-value-value}.

The
[`input`{#date-state-(type=date):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#date-state-(type=date):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#date-state-(type=date):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#date-state-(type=date):do-not-apply} to the
element:
[`accept`{#date-state-(type=date):attr-input-accept}](#attr-input-accept),
[`alpha`{#date-state-(type=date):attr-input-alpha}](#attr-input-alpha),
[`alt`{#date-state-(type=date):attr-input-alt}](#attr-input-alt),
[`checked`{#date-state-(type=date):attr-input-checked}](#attr-input-checked),
[`colorspace`{#date-state-(type=date):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#date-state-(type=date):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#date-state-(type=date):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#date-state-(type=date):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#date-state-(type=date):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#date-state-(type=date):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#date-state-(type=date):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#date-state-(type=date):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`maxlength`{#date-state-(type=date):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#date-state-(type=date):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#date-state-(type=date):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#date-state-(type=date):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#date-state-(type=date):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#date-state-(type=date):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#date-state-(type=date):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`size`{#date-state-(type=date):attr-input-size}](#attr-input-size),
[`src`{#date-state-(type=date):attr-input-src}](#attr-input-src), and
[`width`{#date-state-(type=date):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#date-state-(type=date):do-not-apply-2} to the
element:
[`checked`{#date-state-(type=date):dom-input-checked}](#dom-input-checked),
[`selectionStart`{#date-state-(type=date):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#date-state-(type=date):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
and
[`selectionDirection`{#date-state-(type=date):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection)
IDL attributes;
[`setRangeText()`{#date-state-(type=date):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#date-state-(type=date):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.
:::

###### [4.10.5.1.8]{.secno} [Month]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=month`)[](#month-state-(type=month)){.self-link} {#month-state-(type=month)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/month](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/month "<input> elements of type month create input fields that let the user enter a month and year allowing a month and year to be easily entered. The value is a string whose value is in the format "YYYY-MM", where YYYY is the four-digit year and MM is the month number.")

Support in all current engines.

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome20+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android18+]{.firefox_android .yes}[Safari iOSYes]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

When an
[`input`{#month-state-(type=month):the-input-element}](#the-input-element)
element\'s
[`type`{#month-state-(type=month):attr-input-type}](#attr-input-type)
attribute is in the
[Month](#month-state-(type=month)){#month-state-(type=month):month-state-(type=month)}
state, the rules in this section apply.

The
[`input`{#month-state-(type=month):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#month-state-(type=month):represents}
a control for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#month-state-(type=month):concept-fe-value}
to a string representing a specific
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month-2}
represented by its
[value](form-control-infrastructure.html#concept-fe-value){#month-state-(type=month):concept-fe-value-2},
as obtained by [parsing a
month](common-microsyntaxes.html#parse-a-month-string){#month-state-(type=month):parse-a-month-string}
from it. User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#month-state-(type=month):concept-fe-value-3}
to a non-empty string that is not a [valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string}.
If the user agent provides a user interface for selecting a
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month-3},
then the
[value](form-control-infrastructure.html#concept-fe-value){#month-state-(type=month):concept-fe-value-4}
must be set to a [valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-2}
representing the user\'s selection. User agents should allow the user to
set the
[value](form-control-infrastructure.html#concept-fe-value){#month-state-(type=month):concept-fe-value-5}
to the empty string.

**Constraint validation**: While the user interface describes input that
the user agent cannot convert to a [valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-3},
the control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#month-state-(type=month):suffering-from-bad-input}.

See the [introduction section](forms.html#input-author-notes) for a
discussion of the difference between the input format and submission
format for date, time, and number form controls, and the [implementation
notes](#input-impl-notes) regarding localization of form controls.

The
[`value`{#month-state-(type=month):attr-input-value}](#attr-input-value)
attribute, if specified and not empty, must have a value that is a
[valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-4}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#month-state-(type=month):value-sanitization-algorithm}
is as follows**: If the
[value](form-control-infrastructure.html#concept-fe-value){#month-state-(type=month):concept-fe-value-6}
of the element is not a [valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-5},
then set it to the empty string instead.

The [`min`{#month-state-(type=month):attr-input-min}](#attr-input-min)
attribute, if specified, must have a value that is a [valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-6}.
The [`max`{#month-state-(type=month):attr-input-max}](#attr-input-max)
attribute, if specified, must have a value that is a [valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-7}.

The
[`step`{#month-state-(type=month):attr-input-step}](#attr-input-step)
attribute is expressed in months. The [step scale
factor](#concept-input-step-scale){#month-state-(type=month):concept-input-step-scale}
is 1 (there is no conversion needed as the algorithms use months). The
[default
step](#concept-input-step-default){#month-state-(type=month):concept-input-step-default}
is 1 month.

When the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#month-state-(type=month):suffering-from-a-step-mismatch},
the user agent may round the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#month-state-(type=month):concept-fe-value-7}
to the nearest
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month-4}
for which the element would not [suffer from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#month-state-(type=month):suffering-from-a-step-mismatch-2}.

**The [algorithm to convert a string to a
number](#concept-input-value-string-number){#month-state-(type=month):concept-input-value-string-number},
given a string `input`{.variable}, is as follows**: If [parsing a
month](common-microsyntaxes.html#parse-a-month-string){#month-state-(type=month):parse-a-month-string-2}
from `input`{.variable} results in an error, then return an error;
otherwise, return the number of months between January 1970 and the
parsed
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month-5}.

**The [algorithm to convert a number to a
string](#concept-input-value-number-string){#month-state-(type=month):concept-input-value-number-string},
given a number `input`{.variable}, is as follows**: Return a [valid
month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-8}
that represents the
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month-6}
that has `input`{.variable} months between it and January 1970.

**The [algorithm to convert a string to a `Date`
object](#concept-input-value-string-date){#month-state-(type=month):concept-input-value-string-date},
given a string `input`{.variable}, is as follows**: If [parsing a
month](common-microsyntaxes.html#parse-a-month-string){#month-state-(type=month):parse-a-month-string-3}
from `input`{.variable} results in an error, then return an error;
otherwise, return [a new `Date`
object](infrastructure.html#create-a-date-object){#month-state-(type=month):create-a-date-object}
representing midnight UTC on the morning of the first day of the parsed
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month-7}.

**The [algorithm to convert a `Date` object to a
string](#concept-input-value-date-string){#month-state-(type=month):concept-input-value-date-string},
given a
[`Date`{#month-state-(type=month):date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object `input`{.variable}, is as follows**: Return a [valid month
string](common-microsyntaxes.html#valid-month-string){#month-state-(type=month):valid-month-string-9}
that represents the
[month](common-microsyntaxes.html#concept-month){#month-state-(type=month):concept-month-8}
current at the time represented by `input`{.variable} in the UTC time
zone.

::: bookkeeping
The following common
[`input`{#month-state-(type=month):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#month-state-(type=month):concept-input-apply}
to the element:
[`autocomplete`{#month-state-(type=month):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`list`{#month-state-(type=month):attr-input-list}](#attr-input-list),
[`max`{#month-state-(type=month):attr-input-max-2}](#attr-input-max),
[`min`{#month-state-(type=month):attr-input-min-2}](#attr-input-min),
[`readonly`{#month-state-(type=month):attr-input-readonly}](#attr-input-readonly),
[`required`{#month-state-(type=month):attr-input-required}](#attr-input-required),
and
[`step`{#month-state-(type=month):attr-input-step-2}](#attr-input-step)
content attributes;
[`list`{#month-state-(type=month):dom-input-list}](#dom-input-list),
[`value`{#month-state-(type=month):dom-input-value}](#dom-input-value),
[`valueAsDate`{#month-state-(type=month):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#month-state-(type=month):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#month-state-(type=month):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`stepDown()`{#month-state-(type=month):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#month-state-(type=month):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`value`{#month-state-(type=month):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#month-state-(type=month):dom-input-value-value}.

The
[`input`{#month-state-(type=month):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#month-state-(type=month):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#month-state-(type=month):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#month-state-(type=month):do-not-apply} to the
element:
[`accept`{#month-state-(type=month):attr-input-accept}](#attr-input-accept),
[`alpha`{#month-state-(type=month):attr-input-alpha}](#attr-input-alpha),
[`alt`{#month-state-(type=month):attr-input-alt}](#attr-input-alt),
[`checked`{#month-state-(type=month):attr-input-checked}](#attr-input-checked),
[`colorspace`{#month-state-(type=month):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#month-state-(type=month):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#month-state-(type=month):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#month-state-(type=month):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#month-state-(type=month):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#month-state-(type=month):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#month-state-(type=month):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#month-state-(type=month):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`maxlength`{#month-state-(type=month):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#month-state-(type=month):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#month-state-(type=month):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#month-state-(type=month):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#month-state-(type=month):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#month-state-(type=month):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#month-state-(type=month):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`size`{#month-state-(type=month):attr-input-size}](#attr-input-size),
[`src`{#month-state-(type=month):attr-input-src}](#attr-input-src), and
[`width`{#month-state-(type=month):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#month-state-(type=month):do-not-apply-2} to the
element:
[`checked`{#month-state-(type=month):dom-input-checked}](#dom-input-checked),
[`files`{#month-state-(type=month):dom-input-files}](#dom-input-files),
[`selectionStart`{#month-state-(type=month):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#month-state-(type=month):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
and
[`selectionDirection`{#month-state-(type=month):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection)
IDL attributes;
[`setRangeText()`{#month-state-(type=month):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#month-state-(type=month):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.
:::

###### [4.10.5.1.9]{.secno} [Week]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=week`)[](#week-state-(type=week)){.self-link} {#week-state-(type=week)}

::::: {.mdn-anno .wrapped}
MDN

:::: feature
[Element/input/week](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/week "<input> elements of type week create input fields allowing easy entry of a year plus the ISO 8601 week number during that year (i.e., week 1 to 52 or 53).")

::: support
[FirefoxNo]{.firefox .no}[SafariNo]{.safari .no}[Chrome20+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android18+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

When an
[`input`{#week-state-(type=week):the-input-element}](#the-input-element)
element\'s
[`type`{#week-state-(type=week):attr-input-type}](#attr-input-type)
attribute is in the
[Week](#week-state-(type=week)){#week-state-(type=week):week-state-(type=week)}
state, the rules in this section apply.

The
[`input`{#week-state-(type=week):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#week-state-(type=week):represents} a
control for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#week-state-(type=week):concept-fe-value}
to a string representing a specific
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week-2}
represented by its
[value](form-control-infrastructure.html#concept-fe-value){#week-state-(type=week):concept-fe-value-2},
as obtained by [parsing a
week](common-microsyntaxes.html#parse-a-week-string){#week-state-(type=week):parse-a-week-string}
from it. User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#week-state-(type=week):concept-fe-value-3}
to a non-empty string that is not a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string}.
If the user agent provides a user interface for selecting a
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week-3},
then the
[value](form-control-infrastructure.html#concept-fe-value){#week-state-(type=week):concept-fe-value-4}
must be set to a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-2}
representing the user\'s selection. User agents should allow the user to
set the
[value](form-control-infrastructure.html#concept-fe-value){#week-state-(type=week):concept-fe-value-5}
to the empty string.

**Constraint validation**: While the user interface describes input that
the user agent cannot convert to a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-3},
the control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#week-state-(type=week):suffering-from-bad-input}.

See the [introduction section](forms.html#input-author-notes) for a
discussion of the difference between the input format and submission
format for date, time, and number form controls, and the [implementation
notes](#input-impl-notes) regarding localization of form controls.

The
[`value`{#week-state-(type=week):attr-input-value}](#attr-input-value)
attribute, if specified and not empty, must have a value that is a
[valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-4}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#week-state-(type=week):value-sanitization-algorithm}
is as follows**: If the
[value](form-control-infrastructure.html#concept-fe-value){#week-state-(type=week):concept-fe-value-6}
of the element is not a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-5},
then set it to the empty string instead.

The [`min`{#week-state-(type=week):attr-input-min}](#attr-input-min)
attribute, if specified, must have a value that is a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-6}.
The [`max`{#week-state-(type=week):attr-input-max}](#attr-input-max)
attribute, if specified, must have a value that is a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-7}.

The [`step`{#week-state-(type=week):attr-input-step}](#attr-input-step)
attribute is expressed in weeks. The [step scale
factor](#concept-input-step-scale){#week-state-(type=week):concept-input-step-scale}
is 604,800,000 (which converts the weeks to milliseconds, as used in the
other algorithms). The [default
step](#concept-input-step-default){#week-state-(type=week):concept-input-step-default}
is 1 week. The [default step
base](#concept-input-step-default-base){#week-state-(type=week):concept-input-step-default-base}
is −259,200,000 (the start of week 1970-W01).

When the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#week-state-(type=week):suffering-from-a-step-mismatch},
the user agent may round the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#week-state-(type=week):concept-fe-value-7}
to the nearest
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week-4}
for which the element would not [suffer from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#week-state-(type=week):suffering-from-a-step-mismatch-2}.

**The [algorithm to convert a string to a
number](#concept-input-value-string-number){#week-state-(type=week):concept-input-value-string-number},
given a string `input`{.variable}, is as follows**: If [parsing a week
string](common-microsyntaxes.html#parse-a-week-string){#week-state-(type=week):parse-a-week-string-2}
from `input`{.variable} results in an error, then return an error;
otherwise, return the number of milliseconds elapsed from midnight UTC
on the morning of 1970-01-01 (the time represented by the value
\"`1970-01-01T00:00:00.0Z`\") to midnight UTC on the morning of the
Monday of the parsed
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week-5},
ignoring leap seconds.

**The [algorithm to convert a number to a
string](#concept-input-value-number-string){#week-state-(type=week):concept-input-value-number-string},
given a number `input`{.variable}, is as follows**: Return a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-8}
that represents the
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week-6}
that, in UTC, is current `input`{.variable} milliseconds after midnight
UTC on the morning of 1970-01-01 (the time represented by the value
\"`1970-01-01T00:00:00.0Z`\").

**The [algorithm to convert a string to a `Date`
object](#concept-input-value-string-date){#week-state-(type=week):concept-input-value-string-date},
given a string `input`{.variable}, is as follows**: If [parsing a
week](common-microsyntaxes.html#parse-a-week-string){#week-state-(type=week):parse-a-week-string-3}
from `input`{.variable} results in an error, then return an error;
otherwise, return [a new `Date`
object](infrastructure.html#create-a-date-object){#week-state-(type=week):create-a-date-object}
representing midnight UTC on the morning of the Monday of the parsed
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week-7}.

**The [algorithm to convert a `Date` object to a
string](#concept-input-value-date-string){#week-state-(type=week):concept-input-value-date-string},
given a
[`Date`{#week-state-(type=week):date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object `input`{.variable}, is as follows**: Return a [valid week
string](common-microsyntaxes.html#valid-week-string){#week-state-(type=week):valid-week-string-9}
that represents the
[week](common-microsyntaxes.html#concept-week){#week-state-(type=week):concept-week-8}
current at the time represented by `input`{.variable} in the UTC time
zone.

::: bookkeeping
The following common
[`input`{#week-state-(type=week):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#week-state-(type=week):concept-input-apply}
to the element:
[`autocomplete`{#week-state-(type=week):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`list`{#week-state-(type=week):attr-input-list}](#attr-input-list),
[`max`{#week-state-(type=week):attr-input-max-2}](#attr-input-max),
[`min`{#week-state-(type=week):attr-input-min-2}](#attr-input-min),
[`readonly`{#week-state-(type=week):attr-input-readonly}](#attr-input-readonly),
[`required`{#week-state-(type=week):attr-input-required}](#attr-input-required),
and
[`step`{#week-state-(type=week):attr-input-step-2}](#attr-input-step)
content attributes;
[`list`{#week-state-(type=week):dom-input-list}](#dom-input-list),
[`value`{#week-state-(type=week):dom-input-value}](#dom-input-value),
[`valueAsDate`{#week-state-(type=week):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#week-state-(type=week):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#week-state-(type=week):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`stepDown()`{#week-state-(type=week):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#week-state-(type=week):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`value`{#week-state-(type=week):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#week-state-(type=week):dom-input-value-value}.

The
[`input`{#week-state-(type=week):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#week-state-(type=week):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#week-state-(type=week):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#week-state-(type=week):do-not-apply} to the
element:
[`accept`{#week-state-(type=week):attr-input-accept}](#attr-input-accept),
[`alpha`{#week-state-(type=week):attr-input-alpha}](#attr-input-alpha),
[`alt`{#week-state-(type=week):attr-input-alt}](#attr-input-alt),
[`checked`{#week-state-(type=week):attr-input-checked}](#attr-input-checked),
[`colorspace`{#week-state-(type=week):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#week-state-(type=week):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#week-state-(type=week):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#week-state-(type=week):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#week-state-(type=week):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#week-state-(type=week):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#week-state-(type=week):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#week-state-(type=week):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`maxlength`{#week-state-(type=week):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#week-state-(type=week):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#week-state-(type=week):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#week-state-(type=week):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#week-state-(type=week):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#week-state-(type=week):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#week-state-(type=week):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`size`{#week-state-(type=week):attr-input-size}](#attr-input-size),
[`src`{#week-state-(type=week):attr-input-src}](#attr-input-src), and
[`width`{#week-state-(type=week):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#week-state-(type=week):do-not-apply-2} to the
element:
[`checked`{#week-state-(type=week):dom-input-checked}](#dom-input-checked),
[`files`{#week-state-(type=week):dom-input-files}](#dom-input-files),
[`selectionStart`{#week-state-(type=week):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#week-state-(type=week):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
and
[`selectionDirection`{#week-state-(type=week):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection)
IDL attributes;
[`setRangeText()`{#week-state-(type=week):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#week-state-(type=week):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.
:::

###### [4.10.5.1.10]{.secno} [Time]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=time`)[](#time-state-(type=time)){.self-link} {#time-state-(type=time)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/time](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time "<input> elements of type time create input fields designed to let the user easily enter a time (hours and minutes, and optionally seconds).")

Support in all current engines.

::: support
[Firefox57+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome20+]{.chrome .yes}

------------------------------------------------------------------------

[Opera10+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#time-state-(type=time):the-input-element}](#the-input-element)
element\'s
[`type`{#time-state-(type=time):attr-input-type}](#attr-input-type)
attribute is in the
[Time](#time-state-(type=time)){#time-state-(type=time):time-state-(type=time)}
state, the rules in this section apply.

The
[`input`{#time-state-(type=time):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#time-state-(type=time):represents} a
control for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#time-state-(type=time):concept-fe-value}
to a string representing a specific
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time-2}
represented by its
[value](form-control-infrastructure.html#concept-fe-value){#time-state-(type=time):concept-fe-value-2},
as obtained by [parsing a
time](common-microsyntaxes.html#parse-a-time-string){#time-state-(type=time):parse-a-time-string}
from it. User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#time-state-(type=time):concept-fe-value-3}
to a non-empty string that is not a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string}.
If the user agent provides a user interface for selecting a
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time-3},
then the
[value](form-control-infrastructure.html#concept-fe-value){#time-state-(type=time):concept-fe-value-4}
must be set to a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-2}
representing the user\'s selection. User agents should allow the user to
set the
[value](form-control-infrastructure.html#concept-fe-value){#time-state-(type=time):concept-fe-value-5}
to the empty string.

**Constraint validation**: While the user interface describes input that
the user agent cannot convert to a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-3},
the control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#time-state-(type=time):suffering-from-bad-input}.

See the [introduction section](forms.html#input-author-notes) for a
discussion of the difference between the input format and submission
format for date, time, and number form controls, and the [implementation
notes](#input-impl-notes) regarding localization of form controls.

The
[`value`{#time-state-(type=time):attr-input-value}](#attr-input-value)
attribute, if specified and not empty, must have a value that is a
[valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-4}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#time-state-(type=time):value-sanitization-algorithm}
is as follows**: If the
[value](form-control-infrastructure.html#concept-fe-value){#time-state-(type=time):concept-fe-value-6}
of the element is not a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-5},
then set it to the empty string instead.

The form control [has a periodic
domain](#has-a-periodic-domain){#time-state-(type=time):has-a-periodic-domain}.

The [`min`{#time-state-(type=time):attr-input-min}](#attr-input-min)
attribute, if specified, must have a value that is a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-6}.
The [`max`{#time-state-(type=time):attr-input-max}](#attr-input-max)
attribute, if specified, must have a value that is a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-7}.

The [`step`{#time-state-(type=time):attr-input-step}](#attr-input-step)
attribute is expressed in seconds. The [step scale
factor](#concept-input-step-scale){#time-state-(type=time):concept-input-step-scale}
is 1000 (which converts the seconds to milliseconds, as used in the
other algorithms). The [default
step](#concept-input-step-default){#time-state-(type=time):concept-input-step-default}
is 60 seconds.

When the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#time-state-(type=time):suffering-from-a-step-mismatch},
the user agent may round the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#time-state-(type=time):concept-fe-value-7}
to the nearest
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time-4}
for which the element would not [suffer from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#time-state-(type=time):suffering-from-a-step-mismatch-2}.

**The [algorithm to convert a string to a
number](#concept-input-value-string-number){#time-state-(type=time):concept-input-value-string-number},
given a string `input`{.variable}, is as follows**: If [parsing a
time](common-microsyntaxes.html#parse-a-time-string){#time-state-(type=time):parse-a-time-string-2}
from `input`{.variable} results in an error, then return an error;
otherwise, return the number of milliseconds elapsed from midnight to
the parsed
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time-5}
on a day with no time changes.

**The [algorithm to convert a number to a
string](#concept-input-value-number-string){#time-state-(type=time):concept-input-value-number-string},
given a number `input`{.variable}, is as follows**: Return a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-8}
that represents the
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time-6}
that is `input`{.variable} milliseconds after midnight on a day with no
time changes.

**The [algorithm to convert a string to a `Date`
object](#concept-input-value-string-date){#time-state-(type=time):concept-input-value-string-date},
given a string `input`{.variable}, is as follows**: If [parsing a
time](common-microsyntaxes.html#parse-a-time-string){#time-state-(type=time):parse-a-time-string-3}
from `input`{.variable} results in an error, then return an error;
otherwise, return [a new `Date`
object](infrastructure.html#create-a-date-object){#time-state-(type=time):create-a-date-object}
representing the parsed
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time-7}
in UTC on 1970-01-01.

**The [algorithm to convert a `Date` object to a
string](#concept-input-value-date-string){#time-state-(type=time):concept-input-value-date-string},
given a
[`Date`{#time-state-(type=time):date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object `input`{.variable}, is as follows**: Return a [valid time
string](common-microsyntaxes.html#valid-time-string){#time-state-(type=time):valid-time-string-9}
that represents the UTC
[time](common-microsyntaxes.html#concept-time){#time-state-(type=time):concept-time-8}
component that is represented by `input`{.variable}.

::: bookkeeping
The following common
[`input`{#time-state-(type=time):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#time-state-(type=time):concept-input-apply}
to the element:
[`autocomplete`{#time-state-(type=time):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`list`{#time-state-(type=time):attr-input-list}](#attr-input-list),
[`max`{#time-state-(type=time):attr-input-max-2}](#attr-input-max),
[`min`{#time-state-(type=time):attr-input-min-2}](#attr-input-min),
[`readonly`{#time-state-(type=time):attr-input-readonly}](#attr-input-readonly),
[`required`{#time-state-(type=time):attr-input-required}](#attr-input-required),
and
[`step`{#time-state-(type=time):attr-input-step-2}](#attr-input-step)
content attributes;
[`list`{#time-state-(type=time):dom-input-list}](#dom-input-list),
[`value`{#time-state-(type=time):dom-input-value}](#dom-input-value),
[`valueAsDate`{#time-state-(type=time):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#time-state-(type=time):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#time-state-(type=time):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`stepDown()`{#time-state-(type=time):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#time-state-(type=time):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`value`{#time-state-(type=time):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#time-state-(type=time):dom-input-value-value}.

The
[`input`{#time-state-(type=time):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#time-state-(type=time):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#time-state-(type=time):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#time-state-(type=time):do-not-apply} to the
element:
[`accept`{#time-state-(type=time):attr-input-accept}](#attr-input-accept),
[`alpha`{#time-state-(type=time):attr-input-alpha}](#attr-input-alpha),
[`alt`{#time-state-(type=time):attr-input-alt}](#attr-input-alt),
[`checked`{#time-state-(type=time):attr-input-checked}](#attr-input-checked),
[`colorspace`{#time-state-(type=time):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#time-state-(type=time):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#time-state-(type=time):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#time-state-(type=time):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#time-state-(type=time):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#time-state-(type=time):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#time-state-(type=time):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#time-state-(type=time):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`maxlength`{#time-state-(type=time):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#time-state-(type=time):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#time-state-(type=time):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#time-state-(type=time):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#time-state-(type=time):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#time-state-(type=time):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#time-state-(type=time):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`size`{#time-state-(type=time):attr-input-size}](#attr-input-size),
[`src`{#time-state-(type=time):attr-input-src}](#attr-input-src), and
[`width`{#time-state-(type=time):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#time-state-(type=time):do-not-apply-2} to the
element:
[`checked`{#time-state-(type=time):dom-input-checked}](#dom-input-checked),
[`files`{#time-state-(type=time):dom-input-files}](#dom-input-files),
[`selectionStart`{#time-state-(type=time):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#time-state-(type=time):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
and
[`selectionDirection`{#time-state-(type=time):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection)
IDL attributes;
[`setRangeText()`{#time-state-(type=time):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#time-state-(type=time):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.
:::

###### [4.10.5.1.11]{.secno} [Local Date and Time]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=datetime-local`)[](#local-date-and-time-state-(type=datetime-local)){.self-link} {#local-date-and-time-state-(type=datetime-local)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/datetime-local](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local "<input> elements of type datetime-local create input controls that let the user easily enter both a date and a time, including the year, month, and day as well as the time in hours and minutes.")

Support in all current engines.

::: support
[Firefox93+]{.firefox .yes}[Safari14.1+]{.safari
.yes}[Chrome20+]{.chrome .yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#local-date-and-time-state-(type=datetime-local):the-input-element}](#the-input-element)
element\'s
[`type`{#local-date-and-time-state-(type=datetime-local):attr-input-type}](#attr-input-type)
attribute is in the [Local Date and
Time](#local-date-and-time-state-(type=datetime-local)){#local-date-and-time-state-(type=datetime-local):local-date-and-time-state-(type=datetime-local)}
state, the rules in this section apply.

The
[`input`{#local-date-and-time-state-(type=datetime-local):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#local-date-and-time-state-(type=datetime-local):represents}
a control for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#local-date-and-time-state-(type=datetime-local):concept-fe-value}
to a string representing a [local date and
time](common-microsyntaxes.html#concept-datetime-local){#local-date-and-time-state-(type=datetime-local):concept-datetime-local},
with no time-zone offset information.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the [date and
time](common-microsyntaxes.html#concept-datetime-local){#local-date-and-time-state-(type=datetime-local):concept-datetime-local-2}
represented by its
[value](form-control-infrastructure.html#concept-fe-value){#local-date-and-time-state-(type=datetime-local):concept-fe-value-2},
as obtained by [parsing a date and
time](common-microsyntaxes.html#parse-a-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):parse-a-local-date-and-time-string}
from it. User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#local-date-and-time-state-(type=datetime-local):concept-fe-value-3}
to a non-empty string that is not a [valid normalized local date and
time
string](common-microsyntaxes.html#valid-normalised-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-normalised-local-date-and-time-string}.
If the user agent provides a user interface for selecting a [local date
and
time](common-microsyntaxes.html#concept-datetime-local){#local-date-and-time-state-(type=datetime-local):concept-datetime-local-3},
then the
[value](form-control-infrastructure.html#concept-fe-value){#local-date-and-time-state-(type=datetime-local):concept-fe-value-4}
must be set to a [valid normalized local date and time
string](common-microsyntaxes.html#valid-normalised-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-normalised-local-date-and-time-string-2}
representing the user\'s selection. User agents should allow the user to
set the
[value](form-control-infrastructure.html#concept-fe-value){#local-date-and-time-state-(type=datetime-local):concept-fe-value-5}
to the empty string.

**Constraint validation**: While the user interface describes input that
the user agent cannot convert to a [valid normalized local date and time
string](common-microsyntaxes.html#valid-normalised-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-normalised-local-date-and-time-string-3},
the control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#local-date-and-time-state-(type=datetime-local):suffering-from-bad-input}.

See the [introduction section](forms.html#input-author-notes) for a
discussion of the difference between the input format and submission
format for date, time, and number form controls, and the [implementation
notes](#input-impl-notes) regarding localization of form controls.

The
[`value`{#local-date-and-time-state-(type=datetime-local):attr-input-value}](#attr-input-value)
attribute, if specified and not empty, must have a value that is a
[valid local date and time
string](common-microsyntaxes.html#valid-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-local-date-and-time-string}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#local-date-and-time-state-(type=datetime-local):value-sanitization-algorithm}
is as follows**: If the
[value](form-control-infrastructure.html#concept-fe-value){#local-date-and-time-state-(type=datetime-local):concept-fe-value-6}
of the element is a [valid local date and time
string](common-microsyntaxes.html#valid-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-local-date-and-time-string-2},
then set it to a [valid normalized local date and time
string](common-microsyntaxes.html#valid-normalised-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-normalised-local-date-and-time-string-4}
representing the same date and time; otherwise, set it to the empty
string instead.

The
[`min`{#local-date-and-time-state-(type=datetime-local):attr-input-min}](#attr-input-min)
attribute, if specified, must have a value that is a [valid local date
and time
string](common-microsyntaxes.html#valid-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-local-date-and-time-string-3}.
The
[`max`{#local-date-and-time-state-(type=datetime-local):attr-input-max}](#attr-input-max)
attribute, if specified, must have a value that is a [valid local date
and time
string](common-microsyntaxes.html#valid-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-local-date-and-time-string-4}.

The
[`step`{#local-date-and-time-state-(type=datetime-local):attr-input-step}](#attr-input-step)
attribute is expressed in seconds. The [step scale
factor](#concept-input-step-scale){#local-date-and-time-state-(type=datetime-local):concept-input-step-scale}
is 1000 (which converts the seconds to milliseconds, as used in the
other algorithms). The [default
step](#concept-input-step-default){#local-date-and-time-state-(type=datetime-local):concept-input-step-default}
is 60 seconds.

When the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#local-date-and-time-state-(type=datetime-local):suffering-from-a-step-mismatch},
the user agent may round the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#local-date-and-time-state-(type=datetime-local):concept-fe-value-7}
to the nearest [local date and
time](common-microsyntaxes.html#concept-datetime-local){#local-date-and-time-state-(type=datetime-local):concept-datetime-local-4}
for which the element would not [suffer from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#local-date-and-time-state-(type=datetime-local):suffering-from-a-step-mismatch-2}.

**The [algorithm to convert a string to a
number](#concept-input-value-string-number){#local-date-and-time-state-(type=datetime-local):concept-input-value-string-number},
given a string `input`{.variable}, is as follows**: If [parsing a date
and
time](common-microsyntaxes.html#parse-a-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):parse-a-local-date-and-time-string-2}
from `input`{.variable} results in an error, then return an error;
otherwise, return the number of milliseconds elapsed from midnight on
the morning of 1970-01-01 (the time represented by the value
\"`1970-01-01T00:00:00.0`\") to the parsed [local date and
time](common-microsyntaxes.html#concept-datetime-local){#local-date-and-time-state-(type=datetime-local):concept-datetime-local-5},
ignoring leap seconds.

**The [algorithm to convert a number to a
string](#concept-input-value-number-string){#local-date-and-time-state-(type=datetime-local):concept-input-value-number-string},
given a number `input`{.variable}, is as follows**: Return a [valid
normalized local date and time
string](common-microsyntaxes.html#valid-normalised-local-date-and-time-string){#local-date-and-time-state-(type=datetime-local):valid-normalised-local-date-and-time-string-5}
that represents the date and time that is `input`{.variable}
milliseconds after midnight on the morning of 1970-01-01 (the time
represented by the value \"`1970-01-01T00:00:00.0`\").

See [the note on historical dates](#only-contemporary-times) in the
[Date](#date-state-(type=date)){#local-date-and-time-state-(type=datetime-local):date-state-(type=date)}
state section.

::: bookkeeping
The following common
[`input`{#local-date-and-time-state-(type=datetime-local):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#local-date-and-time-state-(type=datetime-local):concept-input-apply}
to the element:
[`autocomplete`{#local-date-and-time-state-(type=datetime-local):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`list`{#local-date-and-time-state-(type=datetime-local):attr-input-list}](#attr-input-list),
[`max`{#local-date-and-time-state-(type=datetime-local):attr-input-max-2}](#attr-input-max),
[`min`{#local-date-and-time-state-(type=datetime-local):attr-input-min-2}](#attr-input-min),
[`readonly`{#local-date-and-time-state-(type=datetime-local):attr-input-readonly}](#attr-input-readonly),
[`required`{#local-date-and-time-state-(type=datetime-local):attr-input-required}](#attr-input-required),
and
[`step`{#local-date-and-time-state-(type=datetime-local):attr-input-step-2}](#attr-input-step)
content attributes;
[`list`{#local-date-and-time-state-(type=datetime-local):dom-input-list}](#dom-input-list),
[`value`{#local-date-and-time-state-(type=datetime-local):dom-input-value}](#dom-input-value),
and
[`valueAsNumber`{#local-date-and-time-state-(type=datetime-local):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#local-date-and-time-state-(type=datetime-local):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`stepDown()`{#local-date-and-time-state-(type=datetime-local):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#local-date-and-time-state-(type=datetime-local):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`value`{#local-date-and-time-state-(type=datetime-local):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#local-date-and-time-state-(type=datetime-local):dom-input-value-value}.

The
[`input`{#local-date-and-time-state-(type=datetime-local):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#local-date-and-time-state-(type=datetime-local):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#local-date-and-time-state-(type=datetime-local):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#local-date-and-time-state-(type=datetime-local):do-not-apply}
to the element:
[`accept`{#local-date-and-time-state-(type=datetime-local):attr-input-accept}](#attr-input-accept),
[`alpha`{#local-date-and-time-state-(type=datetime-local):attr-input-alpha}](#attr-input-alpha),
[`alt`{#local-date-and-time-state-(type=datetime-local):attr-input-alt}](#attr-input-alt),
[`checked`{#local-date-and-time-state-(type=datetime-local):attr-input-checked}](#attr-input-checked),
[`colorspace`{#local-date-and-time-state-(type=datetime-local):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#local-date-and-time-state-(type=datetime-local):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#local-date-and-time-state-(type=datetime-local):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#local-date-and-time-state-(type=datetime-local):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#local-date-and-time-state-(type=datetime-local):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#local-date-and-time-state-(type=datetime-local):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#local-date-and-time-state-(type=datetime-local):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#local-date-and-time-state-(type=datetime-local):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`maxlength`{#local-date-and-time-state-(type=datetime-local):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#local-date-and-time-state-(type=datetime-local):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#local-date-and-time-state-(type=datetime-local):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#local-date-and-time-state-(type=datetime-local):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#local-date-and-time-state-(type=datetime-local):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#local-date-and-time-state-(type=datetime-local):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#local-date-and-time-state-(type=datetime-local):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`size`{#local-date-and-time-state-(type=datetime-local):attr-input-size}](#attr-input-size),
[`src`{#local-date-and-time-state-(type=datetime-local):attr-input-src}](#attr-input-src),
and
[`width`{#local-date-and-time-state-(type=datetime-local):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#local-date-and-time-state-(type=datetime-local):do-not-apply-2}
to the element:
[`checked`{#local-date-and-time-state-(type=datetime-local):dom-input-checked}](#dom-input-checked),
[`files`{#local-date-and-time-state-(type=datetime-local):dom-input-files}](#dom-input-files),
[`selectionStart`{#local-date-and-time-state-(type=datetime-local):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#local-date-and-time-state-(type=datetime-local):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#local-date-and-time-state-(type=datetime-local):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
and
[`valueAsDate`{#local-date-and-time-state-(type=datetime-local):dom-input-valueasdate}](#dom-input-valueasdate)
IDL attributes;
[`setRangeText()`{#local-date-and-time-state-(type=datetime-local):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#local-date-and-time-state-(type=datetime-local):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.
:::

::: example
The following example shows part of a flight booking application. The
application uses an
[`input`{#local-date-and-time-state-(type=datetime-local):the-input-element-4}](#the-input-element)
element with its
[`type`{#local-date-and-time-state-(type=datetime-local):attr-input-type-2}](#attr-input-type)
attribute set to
[`datetime-local`{#local-date-and-time-state-(type=datetime-local):local-date-and-time-state-(type=datetime-local)-2}](#local-date-and-time-state-(type=datetime-local)),
and it then interprets the given date and time in the time zone of the
selected airport.

``` html
<fieldset>
 <legend>Destination</legend>
 <p><label>Airport: <input type=text name=to list=airports></label></p>
 <p><label>Departure time: <input type=datetime-local name=totime step=3600></label></p>
</fieldset>
<datalist id=airports>
 <option value=ATL label="Atlanta">
 <option value=MEM label="Memphis">
 <option value=LHR label="London Heathrow">
 <option value=LAX label="Los Angeles">
 <option value=FRA label="Frankfurt">
</datalist>
```
:::

###### [4.10.5.1.12]{.secno} [Number]{#number-state .dfn dfn-for="input" dfn-type="element-state"} state (`type=number`)[](#number-state-(type=number)){.self-link} {#number-state-(type=number)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/number](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/number "<input> elements of type number are used to let the user enter a number. They include built-in validation to reject non-numerical entries.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome7+]{.chrome
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

When an
[`input`{#number-state-(type=number):the-input-element}](#the-input-element)
element\'s
[`type`{#number-state-(type=number):attr-input-type}](#attr-input-type)
attribute is in the
[Number](#number-state-(type=number)){#number-state-(type=number):number-state-(type=number)}
state, the rules in this section apply.

The
[`input`{#number-state-(type=number):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#number-state-(type=number):represents}
a control for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#number-state-(type=number):concept-fe-value}
to a string representing a number.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the number represented by its
[value](form-control-infrastructure.html#concept-fe-value){#number-state-(type=number):concept-fe-value-2},
as obtained from applying the [rules for parsing floating-point number
values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#number-state-(type=number):rules-for-parsing-floating-point-number-values}
to it. User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#number-state-(type=number):concept-fe-value-3}
to a non-empty string that is not a [valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#number-state-(type=number):valid-floating-point-number}.
If the user agent provides a user interface for selecting a number, then
the
[value](form-control-infrastructure.html#concept-fe-value){#number-state-(type=number):concept-fe-value-4}
must be set to the [best representation of the number representing the
user\'s selection as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#number-state-(type=number):best-representation-of-the-number-as-a-floating-point-number}.
User agents should allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#number-state-(type=number):concept-fe-value-5}
to the empty string.

**Constraint validation**: While the user interface describes input that
the user agent cannot convert to a [valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#number-state-(type=number):valid-floating-point-number-2},
the control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#number-state-(type=number):suffering-from-bad-input}.

This specification does not define what user interface user agents are
to use; user agent vendors are encouraged to consider what would best
serve their users\' needs. For example, a user agent in Persian or
Arabic markets might support Persian and Arabic numeric input
(converting it to the format required for submission as described
above). Similarly, a user agent designed for Romans might display the
value in Roman numerals rather than in decimal; or (more realistically)
a user agent designed for the French market might display the value with
apostrophes between thousands and commas before the decimals, and allow
the user to enter a value in that manner, internally converting it to
the submission format described above.

The
[`value`{#number-state-(type=number):attr-input-value}](#attr-input-value)
attribute, if specified and not empty, must have a value that is a
[valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#number-state-(type=number):valid-floating-point-number-3}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#number-state-(type=number):value-sanitization-algorithm}
is as follows**: If the
[value](form-control-infrastructure.html#concept-fe-value){#number-state-(type=number):concept-fe-value-6}
of the element is not a [valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#number-state-(type=number):valid-floating-point-number-4},
then set it to the empty string instead.

The [`min`{#number-state-(type=number):attr-input-min}](#attr-input-min)
attribute, if specified, must have a value that is a [valid
floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#number-state-(type=number):valid-floating-point-number-5}.
The [`max`{#number-state-(type=number):attr-input-max}](#attr-input-max)
attribute, if specified, must have a value that is a [valid
floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#number-state-(type=number):valid-floating-point-number-6}.

The [step scale
factor](#concept-input-step-scale){#number-state-(type=number):concept-input-step-scale}
is 1. The [default
step](#concept-input-step-default){#number-state-(type=number):concept-input-step-default}
is 1 (allowing only integers to be selected by the user, unless the
[step
base](#concept-input-min-zero){#number-state-(type=number):concept-input-min-zero}
has a non-integer value).

When the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#number-state-(type=number):suffering-from-a-step-mismatch},
the user agent may round the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#number-state-(type=number):concept-fe-value-7}
to the nearest number for which the element would not [suffer from a
step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#number-state-(type=number):suffering-from-a-step-mismatch-2}.
If there are two such numbers, user agents are encouraged to pick the
one nearest positive infinity.

**The [algorithm to convert a string to a
number](#concept-input-value-string-number){#number-state-(type=number):concept-input-value-string-number},
given a string `input`{.variable}, is as follows**: If applying the
[rules for parsing floating-point number
values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#number-state-(type=number):rules-for-parsing-floating-point-number-values-2}
to `input`{.variable} results in an error, then return an error;
otherwise, return the resulting number.

**The [algorithm to convert a number to a
string](#concept-input-value-number-string){#number-state-(type=number):concept-input-value-number-string},
given a number `input`{.variable}, is as follows**: Return a [valid
floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#number-state-(type=number):valid-floating-point-number-7}
that represents `input`{.variable}.

::: bookkeeping
The following common
[`input`{#number-state-(type=number):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#number-state-(type=number):concept-input-apply}
to the element:
[`autocomplete`{#number-state-(type=number):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`list`{#number-state-(type=number):attr-input-list}](#attr-input-list),
[`max`{#number-state-(type=number):attr-input-max-2}](#attr-input-max),
[`min`{#number-state-(type=number):attr-input-min-2}](#attr-input-min),
[`placeholder`{#number-state-(type=number):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#number-state-(type=number):attr-input-readonly}](#attr-input-readonly),
[`required`{#number-state-(type=number):attr-input-required}](#attr-input-required),
and
[`step`{#number-state-(type=number):attr-input-step}](#attr-input-step)
content attributes;
[`list`{#number-state-(type=number):dom-input-list}](#dom-input-list),
[`value`{#number-state-(type=number):dom-input-value}](#dom-input-value),
and
[`valueAsNumber`{#number-state-(type=number):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#number-state-(type=number):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`stepDown()`{#number-state-(type=number):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#number-state-(type=number):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`value`{#number-state-(type=number):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#number-state-(type=number):dom-input-value-value}.

The
[`input`{#number-state-(type=number):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#number-state-(type=number):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#number-state-(type=number):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#number-state-(type=number):do-not-apply} to the
element:
[`accept`{#number-state-(type=number):attr-input-accept}](#attr-input-accept),
[`alpha`{#number-state-(type=number):attr-input-alpha}](#attr-input-alpha),
[`alt`{#number-state-(type=number):attr-input-alt}](#attr-input-alt),
[`checked`{#number-state-(type=number):attr-input-checked}](#attr-input-checked),
[`colorspace`{#number-state-(type=number):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#number-state-(type=number):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#number-state-(type=number):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#number-state-(type=number):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#number-state-(type=number):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#number-state-(type=number):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#number-state-(type=number):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#number-state-(type=number):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`maxlength`{#number-state-(type=number):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#number-state-(type=number):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#number-state-(type=number):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#number-state-(type=number):attr-input-pattern}](#attr-input-pattern),
[`popovertarget`{#number-state-(type=number):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#number-state-(type=number):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`size`{#number-state-(type=number):attr-input-size}](#attr-input-size),
[`src`{#number-state-(type=number):attr-input-src}](#attr-input-src),
and
[`width`{#number-state-(type=number):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#number-state-(type=number):do-not-apply-2} to the
element:
[`checked`{#number-state-(type=number):dom-input-checked}](#dom-input-checked),
[`files`{#number-state-(type=number):dom-input-files}](#dom-input-files),
[`selectionStart`{#number-state-(type=number):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#number-state-(type=number):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#number-state-(type=number):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
and
[`valueAsDate`{#number-state-(type=number):dom-input-valueasdate}](#dom-input-valueasdate)
IDL attributes;
[`setRangeText()`{#number-state-(type=number):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#number-state-(type=number):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.
:::

::: example
Here is an example of using a numeric input control:

``` html
<label>How much do you want to charge? $<input type=number min=0 step=0.01 name=price></label>
```

As described above, a user agent might support numeric input in the
user\'s local format, converting it to the format required for
submission as described above. This might include handling grouping
separators (as in \"872,000,000,000\") and various decimal separators
(such as \"3,99\" vs \"3.99\") or using local digits (such as those in
Arabic, Devanagari, Persian, and Thai).
:::

The `type=number` state is not appropriate for input that happens to
only consist of numbers but isn\'t strictly speaking a number. For
example, it would be inappropriate for credit card numbers or US postal
codes. A simple way of determining whether to use `type=number` is to
consider whether it would make sense for the input control to have a
spinbox interface (e.g. with \"up\" and \"down\" arrows). Getting a
credit card number wrong by 1 in the last digit isn\'t a minor mistake,
it\'s as wrong as getting every digit incorrect. So it would not make
sense for the user to select a credit card number using \"up\" and
\"down\" buttons. When a spinbox interface is not appropriate,
`type=text` is probably the right choice (possibly with an
[`inputmode`{#number-state-(type=number):attr-inputmode}](interaction.html#attr-inputmode)
or
[`pattern`{#number-state-(type=number):attr-input-pattern-2}](#attr-input-pattern)
attribute).

###### [4.10.5.1.13]{.secno} [Range]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=range`)[](#range-state-(type=range)){.self-link} {#range-state-(type=range)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/range](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/range "<input> elements of type range let the user specify a numeric value which must be no less than a given value, and no more than another given value. The precise value, however, is not considered important. This is typically represented using a slider or dial control rather than a text entry box like the number input type.")

Support in all current engines.

::: support
[Firefox23+]{.firefox .yes}[Safari3.1+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android52+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android57+]{.chrome_android .yes}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#range-state-(type=range):the-input-element}](#the-input-element)
element\'s
[`type`{#range-state-(type=range):attr-input-type}](#attr-input-type)
attribute is in the
[Range](#range-state-(type=range)){#range-state-(type=range):range-state-(type=range)}
state, the rules in this section apply.

The
[`input`{#range-state-(type=range):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#range-state-(type=range):represents}
a control for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value}
to a string representing a number, but with the caveat that the exact
value is not important, letting UAs provide a simpler interface than
they do for the
[Number](#number-state-(type=number)){#range-state-(type=range):number-state-(type=number)}
state.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the number represented by its
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-2},
as obtained from applying the [rules for parsing floating-point number
values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#range-state-(type=range):rules-for-parsing-floating-point-number-values}
to it. User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-3}
to a string that is not a [valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#range-state-(type=range):valid-floating-point-number}.
If the user agent provides a user interface for selecting a number, then
the
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-4}
must be set to a [best representation of the number representing the
user\'s selection as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#range-state-(type=range):best-representation-of-the-number-as-a-floating-point-number}.
User agents must not allow the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-5}
to the empty string.

**Constraint validation**: While the user interface describes input that
the user agent cannot convert to a [valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#range-state-(type=range):valid-floating-point-number-2},
the control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#range-state-(type=range):suffering-from-bad-input}.

The
[`value`{#range-state-(type=range):attr-input-value}](#attr-input-value)
attribute, if specified, must have a value that is a [valid
floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#range-state-(type=range):valid-floating-point-number-3}.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#range-state-(type=range):value-sanitization-algorithm}
is as follows**: If the
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-6}
of the element is not a [valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#range-state-(type=range):valid-floating-point-number-4},
then set it to the [best representation, as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#range-state-(type=range):best-representation-of-the-number-as-a-floating-point-number-2},
of the [default
value](#concept-input-value-default-range){#range-state-(type=range):concept-input-value-default-range}.

The [default value]{#concept-input-value-default-range .dfn} is the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min}
plus half the difference between the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min-2}
and the
[maximum](#concept-input-max){#range-state-(type=range):concept-input-max},
unless the
[maximum](#concept-input-max){#range-state-(type=range):concept-input-max-2}
is less than the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min-3},
in which case the [default
value](#concept-input-value-default-range){#range-state-(type=range):concept-input-value-default-range-2}
is the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min-4}.

When the element is [suffering from an
underflow](form-control-infrastructure.html#suffering-from-an-underflow){#range-state-(type=range):suffering-from-an-underflow},
the user agent must set the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-7}
to the [best representation, as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#range-state-(type=range):best-representation-of-the-number-as-a-floating-point-number-3},
of the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min-5}.

When the element is [suffering from an
overflow](form-control-infrastructure.html#suffering-from-an-overflow){#range-state-(type=range):suffering-from-an-overflow},
if the
[maximum](#concept-input-max){#range-state-(type=range):concept-input-max-3}
is not less than the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min-6},
the user agent must set the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-8}
to a [valid floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#range-state-(type=range):valid-floating-point-number-5}
that represents the
[maximum](#concept-input-max){#range-state-(type=range):concept-input-max-4}.

When the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#range-state-(type=range):suffering-from-a-step-mismatch},
the user agent must round the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#range-state-(type=range):concept-fe-value-9}
to the nearest number for which the element would not [suffer from a
step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#range-state-(type=range):suffering-from-a-step-mismatch-2},
and which is greater than or equal to the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min-7},
and, if the
[maximum](#concept-input-max){#range-state-(type=range):concept-input-max-5}
is not less than the
[minimum](#concept-input-min){#range-state-(type=range):concept-input-min-8},
which is less than or equal to the
[maximum](#concept-input-max){#range-state-(type=range):concept-input-max-6},
if there is a number that matches these constraints. If two numbers
match these constraints, then user agents must use the one nearest to
positive infinity.

For example, the markup
`<input type="range" min=0 max=100 step=20 value=50>` results in a range
control whose initial value is 60.

::: example
Here is an example of a range control using an autocomplete list with
the
[`list`{#range-state-(type=range):attr-input-list}](#attr-input-list)
attribute. This could be useful if there are values along the full range
of the control that are especially important, such as preconfigured
light levels or typical speed limits in a range control used as a speed
control. The following markup fragment:

``` html
<input type="range" min="-100" max="100" value="0" step="10" name="power" list="powers">
<datalist id="powers">
<option value="0">
<option value="-30">
<option value="30">
 <option value="++50">
</datalist>
```

\...with the following style sheet applied:

``` css
input { writing-mode: vertical-lr; height: 75px; width: 49px; background: #D5CCBB; color: black; }
```

\...might render as:

![A vertical slider control whose primary color is black and whose
background color is beige, with the slider having five tick marks, one
long one at each extremity, and three short ones clustered around the
midpoint.](/images/sample-range.png){width="49" height="75"}

Note how the UA determined the orientation of the control from the ratio
of the style-sheet-specified height and width properties. The colors
were similarly derived from the style sheet. The tick marks, however,
were derived from the markup. In particular, the
[`step`{#range-state-(type=range):attr-input-step}](#attr-input-step)
attribute has not affected the placement of tick marks, the UA deciding
to only use the author-specified completion values and then adding
longer tick marks at the extremes.

Note also how the invalid value `++50` was ignored.
:::

::: example
For another example, consider the following markup fragment:

``` html
<input name=x type=range min=100 max=700 step=9.09090909 value=509.090909>
```

A user agent could display in a variety of ways, for instance:

![As a dial.](/images/sample-range-2a.png){width="231" height="57"}

Or, alternatively, for instance:

![As a long horizontal slider with tick
marks.](/images/sample-range-2b.png){width="445" height="56"}

The user agent could pick which one to display based on the dimensions
given in the style sheet. This would allow it to maintain the same
resolution for the tick marks, despite the differences in width.
:::

::: example
Finally, here is an example of a range control with two labeled values:

``` html
<input type="range" name="a" list="a-values">
<datalist id="a-values">
<option value="10" label="Low">
<option value="90" label="High">
</datalist>
```

With styles that make the control draw vertically, it might look as
follows:

![A vertical slider control with two tick marks, one near the top
labeled \'High\', and one near the bottom labeled
\'Low\'.](/images/sample-range-labels.png){width="103" height="164"}
:::

In this state, the range and step constraints are enforced even during
user input, and there is no way to set the value to the empty string.

The [`min`{#range-state-(type=range):attr-input-min}](#attr-input-min)
attribute, if specified, must have a value that is a [valid
floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#range-state-(type=range):valid-floating-point-number-6}.
The [default
minimum](#concept-input-min-default){#range-state-(type=range):concept-input-min-default}
is 0. The
[`max`{#range-state-(type=range):attr-input-max}](#attr-input-max)
attribute, if specified, must have a value that is a [valid
floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#range-state-(type=range):valid-floating-point-number-7}.
The [default
maximum](#concept-input-max-default){#range-state-(type=range):concept-input-max-default}
is 100.

The [step scale
factor](#concept-input-step-scale){#range-state-(type=range):concept-input-step-scale}
is 1. The [default
step](#concept-input-step-default){#range-state-(type=range):concept-input-step-default}
is 1 (allowing only integers, unless the
[`min`{#range-state-(type=range):attr-input-min-2}](#attr-input-min)
attribute has a non-integer value).

**The [algorithm to convert a string to a
number](#concept-input-value-string-number){#range-state-(type=range):concept-input-value-string-number},
given a string `input`{.variable}, is as follows**: If applying the
[rules for parsing floating-point number
values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#range-state-(type=range):rules-for-parsing-floating-point-number-values-2}
to `input`{.variable} results in an error, then return an error;
otherwise, return the resulting number.

**The [algorithm to convert a number to a
string](#concept-input-value-number-string){#range-state-(type=range):concept-input-value-number-string},
given a number `input`{.variable}, is as follows**: Return the [best
representation, as a floating-point
number](common-microsyntaxes.html#best-representation-of-the-number-as-a-floating-point-number){#range-state-(type=range):best-representation-of-the-number-as-a-floating-point-number-4},
of `input`{.variable}.

::: bookkeeping
The following common
[`input`{#range-state-(type=range):the-input-element-3}](#the-input-element)
element content attributes, IDL attributes, and methods
[apply](#concept-input-apply){#range-state-(type=range):concept-input-apply}
to the element:
[`autocomplete`{#range-state-(type=range):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`list`{#range-state-(type=range):attr-input-list-2}](#attr-input-list),
[`max`{#range-state-(type=range):attr-input-max-2}](#attr-input-max),
[`min`{#range-state-(type=range):attr-input-min-3}](#attr-input-min),
and
[`step`{#range-state-(type=range):attr-input-step-2}](#attr-input-step)
content attributes;
[`list`{#range-state-(type=range):dom-input-list}](#dom-input-list),
[`value`{#range-state-(type=range):dom-input-value}](#dom-input-value),
and
[`valueAsNumber`{#range-state-(type=range):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`stepDown()`{#range-state-(type=range):dom-input-stepdown}](#dom-input-stepdown)
and
[`stepUp()`{#range-state-(type=range):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`value`{#range-state-(type=range):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#range-state-(type=range):dom-input-value-value}.

The
[`input`{#range-state-(type=range):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#range-state-(type=range):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#range-state-(type=range):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#range-state-(type=range):do-not-apply} to the
element:
[`accept`{#range-state-(type=range):attr-input-accept}](#attr-input-accept),
[`alpha`{#range-state-(type=range):attr-input-alpha}](#attr-input-alpha),
[`alt`{#range-state-(type=range):attr-input-alt}](#attr-input-alt),
[`checked`{#range-state-(type=range):attr-input-checked}](#attr-input-checked),
[`colorspace`{#range-state-(type=range):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#range-state-(type=range):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#range-state-(type=range):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#range-state-(type=range):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#range-state-(type=range):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#range-state-(type=range):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#range-state-(type=range):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#range-state-(type=range):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`maxlength`{#range-state-(type=range):attr-input-maxlength}](#attr-input-maxlength),
[`minlength`{#range-state-(type=range):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#range-state-(type=range):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#range-state-(type=range):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#range-state-(type=range):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#range-state-(type=range):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#range-state-(type=range):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`readonly`{#range-state-(type=range):attr-input-readonly}](#attr-input-readonly),
[`required`{#range-state-(type=range):attr-input-required}](#attr-input-required),
[`size`{#range-state-(type=range):attr-input-size}](#attr-input-size),
[`src`{#range-state-(type=range):attr-input-src}](#attr-input-src), and
[`width`{#range-state-(type=range):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#range-state-(type=range):do-not-apply-2} to the
element:
[`checked`{#range-state-(type=range):dom-input-checked}](#dom-input-checked),
[`files`{#range-state-(type=range):dom-input-files}](#dom-input-files),
[`selectionStart`{#range-state-(type=range):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#range-state-(type=range):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#range-state-(type=range):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
and
[`valueAsDate`{#range-state-(type=range):dom-input-valueasdate}](#dom-input-valueasdate)
IDL attributes;
[`select()`{#range-state-(type=range):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#range-state-(type=range):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
and
[`setSelectionRange()`{#range-state-(type=range):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange)
methods.
:::

###### [4.10.5.1.14]{.secno} [Color]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=color`)[](#color-state-(type=color)){.self-link} {#color-state-(type=color)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/color](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/color "<input> elements of type color provide a user interface element that lets a user specify a color, either by using a visual color picker interface or by entering the color into a text field in #rrggbb hexadecimal format.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari12.1+]{.safari
.yes}[Chrome20+]{.chrome .yes}

------------------------------------------------------------------------

[Opera12+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android[🔰
27+]{title="Partial implementation."}]{.firefox_android .yes}[Safari
iOS?]{.safari_ios .unknown}[Chrome Android?]{.chrome_android
.unknown}[WebView Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#color-state-(type=color):the-input-element}](#the-input-element)
element\'s
[`type`{#color-state-(type=color):attr-input-type}](#attr-input-type)
attribute is in the
[Color](#color-state-(type=color)){#color-state-(type=color):color-state-(type=color)}
state, the rules in this section apply.

The
[`input`{#color-state-(type=color):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#color-state-(type=color):represents}
a color well control, for setting the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#color-state-(type=color):concept-fe-value}
to a string representing the serialization of a CSS color.

In this state, there is always a CSS color picked, and there is no way
for the end user to set the value to the empty string.

The [`alpha`]{#attr-input-alpha .dfn dfn-for="input"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#color-state-(type=color):boolean-attribute}.
If present, it indicates the CSS color\'s alpha component can be
manipulated by the end user and does not have to be fully opaque.

The [`colorspace`]{#attr-input-colorspace .dfn dfn-for="input"
dfn-type="element-attr"} attribute indicates the color space of the
serialized CSS color. It also hints at the desired user interface for
selecting a CSS color. It is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#color-state-(type=color):enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`limited-srgb`]{#attr-input-colorspace-limited-srgb .dfn
dfn-for="input/colorspace" dfn-type="attr-value"}

[Limited sRGB]{#attr-input-colorspace-limited-srgb-state .dfn}

The CSS color is converted to the
[\'srgb\'](https://drafts.csswg.org/css-color/#valdef-color-srgb){#color-state-(type=color):'srgb'
x-internal="'srgb'"} color space and limited to 8-bits per component,
e.g., \"`#123456`\" or \"`color(srgb 0 1 0 / 0.5)`\".

[`display-p3`]{#attr-input-colorspace-display-p3 .dfn
dfn-for="input/colorspace" dfn-type="attr-value"}

[Display P3]{#attr-input-colorspace-display-p3-state .dfn}

The CSS color is converted to the
[\'display-p3\'](https://drafts.csswg.org/css-color/#valdef-color-display-p3){#color-state-(type=color):'display-p3'
x-internal="'display-p3'"} color space, e.g.,
\"`color(display-p3 1.84 -0.19 0.72 / 0.6)`\".

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Limited
sRGB](#attr-input-colorspace-limited-srgb-state){#color-state-(type=color):attr-input-colorspace-limited-srgb-state}
state.

Whenever the element\'s
[`alpha`{#color-state-(type=color):attr-input-alpha}](#attr-input-alpha)
or
[`colorspace`{#color-state-(type=color):attr-input-colorspace}](#attr-input-colorspace)
attributes are changed, the user agent must run [update a color well
control
color](#update-a-color-well-control-color){#color-state-(type=color):update-a-color-well-control-color}
given the element.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the color represented by its
[value](form-control-infrastructure.html#concept-fe-value){#color-state-(type=color):concept-fe-value-2},
as obtained from
[parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#color-state-(type=color):parsed-as-a-css-color-value
x-internal="parsed-as-a-css-color-value"} it. User agents must not allow
the user to set the
[value](form-control-infrastructure.html#concept-fe-value){#color-state-(type=color):concept-fe-value-3}
to a string that running [update a color well control
color](#update-a-color-well-control-color){#color-state-(type=color):update-a-color-well-control-color-2}
for the element would not set it to. If the user agent provides a user
interface for selecting a CSS color, then the
[value](form-control-infrastructure.html#concept-fe-value){#color-state-(type=color):concept-fe-value-4}
must be set to the result of [serializing a color well control
color](#serialize-a-color-well-control-color){#color-state-(type=color):serialize-a-color-well-control-color}
given the element and the end user\'s selection.

The [input activation
behavior](#input-activation-behavior){#color-state-(type=color):input-activation-behavior}
for such an element `element`{.variable} is to [show the picker, if
applicable](#show-the-picker,-if-applicable){#color-state-(type=color):show-the-picker,-if-applicable},
for `element`{.variable}.

**Constraint validation**: While the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#color-state-(type=color):concept-fe-value-5}
is not the empty string and
[parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#color-state-(type=color):parsed-as-a-css-color-value-2
x-internal="parsed-as-a-css-color-value"} it returns failure, the
control is [suffering from bad
input](form-control-infrastructure.html#suffering-from-bad-input){#color-state-(type=color):suffering-from-bad-input}.

The
[`value`{#color-state-(type=color):attr-input-value}](#attr-input-value)
attribute, if specified and not the empty string, must have a value that
is a CSS color.

**The [value sanitization
algorithm](#value-sanitization-algorithm){#color-state-(type=color):value-sanitization-algorithm}
is as follows**: Run [update a color well control
color](#update-a-color-well-control-color){#color-state-(type=color):update-a-color-well-control-color-3}
for the element.

To [update a color well control
color]{#update-a-color-well-control-color .dfn}, given an element
`element`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#color-state-(type=color):assert
    x-internal="assert"}: `element`{.variable} is an
    [`input`{#color-state-(type=color):the-input-element-3}](#the-input-element)
    element whose
    [`type`{#color-state-(type=color):attr-input-type-2}](#attr-input-type)
    attribute is in the
    [Color](#color-state-(type=color)){#color-state-(type=color):color-state-(type=color)-2}
    state.

2.  Let `color`{.variable} be the result of
    [parsing](https://drafts.csswg.org/css-color/#parse-a-css-color-value){#color-state-(type=color):parsed-as-a-css-color-value-3
    x-internal="parsed-as-a-css-color-value"} `element`{.variable}\'s
    [value](form-control-infrastructure.html#concept-fe-value){#color-state-(type=color):concept-fe-value-6}.

3.  If `color`{.variable} is failure, then set `color`{.variable} to
    [opaque
    black](https://drafts.csswg.org/css-color/#opaque-black){#color-state-(type=color):opaque-black
    x-internal="opaque-black"}.

4.  Set `element`{.variable}\'s
    [value](form-control-infrastructure.html#concept-fe-value){#color-state-(type=color):concept-fe-value-7}
    to the result of [serializing a color well control
    color](#serialize-a-color-well-control-color){#color-state-(type=color):serialize-a-color-well-control-color-2}
    given `element`{.variable} and `color`{.variable}.

To [serialize a color well control
color]{#serialize-a-color-well-control-color .dfn}, given an element
`element`{.variable} and a CSS color `color`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#color-state-(type=color):assert-2
    x-internal="assert"}: `element`{.variable} is an
    [`input`{#color-state-(type=color):the-input-element-4}](#the-input-element)
    element whose
    [`type`{#color-state-(type=color):attr-input-type-3}](#attr-input-type)
    attribute is in the
    [Color](#color-state-(type=color)){#color-state-(type=color):color-state-(type=color)-3}
    state.

2.  Let `htmlCompatible`{.variable} be false.

3.  If `element`{.variable}\'s
    [`alpha`{#color-state-(type=color):attr-input-alpha-2}](#attr-input-alpha)
    attribute is not specified, then set `color`{.variable}\'s alpha
    component to be fully opaque.

4.  If `element`{.variable}\'s
    [`colorspace`{#color-state-(type=color):attr-input-colorspace-2}](#attr-input-colorspace)
    attribute is in the [Limited
    sRGB](#attr-input-colorspace-limited-srgb-state){#color-state-(type=color):attr-input-colorspace-limited-srgb-state-2}
    state:

    1.  Set `color`{.variable} to `color`{.variable}
        [converted](https://drafts.csswg.org/css-color/#color-conversion){#color-state-(type=color):converting-colors
        x-internal="converting-colors"} to the
        [\'srgb\'](https://drafts.csswg.org/css-color/#valdef-color-srgb){#color-state-(type=color):'srgb'-2
        x-internal="'srgb'"} color space.

    2.  Round each of `color`{.variable}\'s components so they are in
        the range 0 to 255, inclusive. Components are to be [rounded
        towards
        +∞](https://drafts.csswg.org/css-values-4/#combine-integers).

    3.  If `element`{.variable}\'s
        [`alpha`{#color-state-(type=color):attr-input-alpha-3}](#attr-input-alpha)
        attribute is not specified, then set `htmlCompatible`{.variable}
        to true.

        This intentionally uses a different serialization path for
        compatibility with an earlier version of the color well control.

    4.  Otherwise, set `color`{.variable} to `color`{.variable}
        converted using the
        [\'color()\'](https://drafts.csswg.org/css-color/#color-function){#color-state-(type=color):color-function
        x-internal="color-function"} function.

5.  Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#color-state-(type=color):assert-3
        x-internal="assert"}: `element`{.variable}\'s
        [`colorspace`{#color-state-(type=color):attr-input-colorspace-3}](#attr-input-colorspace)
        attribute is in the [Display
        P3](#attr-input-colorspace-display-p3-state){#color-state-(type=color):attr-input-colorspace-display-p3-state}
        state.

    2.  Set `color`{.variable} to `color`{.variable}
        [converted](https://drafts.csswg.org/css-color/#color-conversion){#color-state-(type=color):converting-colors-2
        x-internal="converting-colors"} to the
        [\'display-p3\'](https://drafts.csswg.org/css-color/#valdef-color-display-p3){#color-state-(type=color):'display-p3'-2
        x-internal="'display-p3'"} color space.

6.  Return the result of
    [serializing](https://drafts.csswg.org/css-color/#serializing-color-values){#color-state-(type=color):serialisation-of-a-color
    x-internal="serialisation-of-a-color"} `color`{.variable}. If
    `htmlCompatible`{.variable} is true, then do so with
    [HTML-compatible serialization
    requested](https://drafts.csswg.org/css-color/#color-serialization-html-compatible-serialization-is-requested){#color-state-(type=color):html-compatible-serialization-is-requested
    x-internal="html-compatible-serialization-is-requested"}.

::: bookkeeping
The following common
[`input`{#color-state-(type=color):the-input-element-5}](#the-input-element)
element content attributes and IDL attributes
[apply](#concept-input-apply){#color-state-(type=color):concept-input-apply}
to the element:
[`alpha`{#color-state-(type=color):attr-input-alpha-4}](#attr-input-alpha),
[`autocomplete`{#color-state-(type=color):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`colorspace`{#color-state-(type=color):attr-input-colorspace-4}](#attr-input-colorspace),
and
[`list`{#color-state-(type=color):attr-input-list}](#attr-input-list)
content attributes;
[`list`{#color-state-(type=color):dom-input-list}](#dom-input-list) and
[`value`{#color-state-(type=color):dom-input-value}](#dom-input-value)
IDL attributes;
[`select()`{#color-state-(type=color):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select)
method.

The
[`value`{#color-state-(type=color):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[value](#dom-input-value-value){#color-state-(type=color):dom-input-value-value}.

The
[`input`{#color-state-(type=color):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#color-state-(type=color):event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#color-state-(type=color):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#color-state-(type=color):do-not-apply} to the
element:
[`accept`{#color-state-(type=color):attr-input-accept}](#attr-input-accept),
[`alt`{#color-state-(type=color):attr-input-alt}](#attr-input-alt),
[`checked`{#color-state-(type=color):attr-input-checked}](#attr-input-checked),
[`dirname`{#color-state-(type=color):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#color-state-(type=color):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#color-state-(type=color):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#color-state-(type=color):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#color-state-(type=color):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#color-state-(type=color):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#color-state-(type=color):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`max`{#color-state-(type=color):attr-input-max}](#attr-input-max),
[`maxlength`{#color-state-(type=color):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#color-state-(type=color):attr-input-min}](#attr-input-min),
[`minlength`{#color-state-(type=color):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#color-state-(type=color):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#color-state-(type=color):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#color-state-(type=color):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#color-state-(type=color):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#color-state-(type=color):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`readonly`{#color-state-(type=color):attr-input-readonly}](#attr-input-readonly),
[`required`{#color-state-(type=color):attr-input-required}](#attr-input-required),
[`size`{#color-state-(type=color):attr-input-size}](#attr-input-size),
[`src`{#color-state-(type=color):attr-input-src}](#attr-input-src),
[`step`{#color-state-(type=color):attr-input-step}](#attr-input-step),
and
[`width`{#color-state-(type=color):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#color-state-(type=color):do-not-apply-2} to the
element:
[`checked`{#color-state-(type=color):dom-input-checked}](#dom-input-checked),
[`files`{#color-state-(type=color):dom-input-files}](#dom-input-files),
[`selectionStart`{#color-state-(type=color):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#color-state-(type=color):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#color-state-(type=color):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#color-state-(type=color):dom-input-valueasdate}](#dom-input-valueasdate)
and,
[`valueAsNumber`{#color-state-(type=color):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`setRangeText()`{#color-state-(type=color):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#color-state-(type=color):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#color-state-(type=color):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#color-state-(type=color):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.15]{.secno} [Checkbox]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=checkbox`)[](#checkbox-state-(type=checkbox)){.self-link} {#checkbox-state-(type=checkbox)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/checkbox](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox "<input> elements of type checkbox are rendered by default as boxes that are checked (ticked) when activated, like you might see in an official government paper form. The exact appearance depends upon the operating system configuration under which the browser is running. Generally this is a square but it may have rounded corners. A checkbox allows you to select single values for submission in a form (or not).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

When an
[`input`{#checkbox-state-(type=checkbox):the-input-element}](#the-input-element)
element\'s
[`type`{#checkbox-state-(type=checkbox):attr-input-type}](#attr-input-type)
attribute is in the
[Checkbox](#checkbox-state-(type=checkbox)){#checkbox-state-(type=checkbox):checkbox-state-(type=checkbox)}
state, the rules in this section apply.

The
[`input`{#checkbox-state-(type=checkbox):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#checkbox-state-(type=checkbox):represents}
a two-state control that represents the element\'s
[checkedness](form-control-infrastructure.html#concept-fe-checked){#checkbox-state-(type=checkbox):concept-fe-checked}
state. If the element\'s
[checkedness](form-control-infrastructure.html#concept-fe-checked){#checkbox-state-(type=checkbox):concept-fe-checked-2}
state is true, the control represents a positive selection, and if it is
false, a negative selection. If the element\'s
[`indeterminate`{#checkbox-state-(type=checkbox):dom-input-indeterminate}](#dom-input-indeterminate)
IDL attribute is set to true, then the control\'s selection should be
obscured as if the control was in a third, indeterminate, state.

The control is never a true tri-state control, even if the element\'s
[`indeterminate`{#checkbox-state-(type=checkbox):dom-input-indeterminate-2}](#dom-input-indeterminate)
IDL attribute is set to true. The
[`indeterminate`{#checkbox-state-(type=checkbox):dom-input-indeterminate-3}](#dom-input-indeterminate)
IDL attribute only gives the appearance of a third state.

The [input activation
behavior](#input-activation-behavior){#checkbox-state-(type=checkbox):input-activation-behavior}
is to run the following steps:

1.  If the element is not
    [connected](https://dom.spec.whatwg.org/#connected){#checkbox-state-(type=checkbox):connected
    x-internal="connected"}, then return.

2.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#checkbox-state-(type=checkbox):concept-event-fire
    x-internal="concept-event-fire"} named
    [`input`{#checkbox-state-(type=checkbox):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    at the element with the
    [`bubbles`{#checkbox-state-(type=checkbox):dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    and
    [`composed`{#checkbox-state-(type=checkbox):dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
    attributes initialized to true.

3.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#checkbox-state-(type=checkbox):concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`change`{#checkbox-state-(type=checkbox):event-change}](indices.html#event-change)
    at the element with the
    [`bubbles`{#checkbox-state-(type=checkbox):dom-event-bubbles-2}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    attribute initialized to true.

**Constraint validation**: If the element is
*[required](#concept-input-required)* and its
[checkedness](form-control-infrastructure.html#concept-fe-checked){#checkbox-state-(type=checkbox):concept-fe-checked-3}
is false, then the element is [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#checkbox-state-(type=checkbox):suffering-from-being-missing}.

`input`{.variable}`.`[`indeterminate`](#dom-input-indeterminate){#dom-input-indeterminate-dev}` [ = ``value`{.variable}` ]`

:   When set, overrides the rendering of
    [checkbox](#checkbox-state-(type=checkbox)){#checkbox-state-(type=checkbox):checkbox-state-(type=checkbox)-2}
    controls so that the current value is not visible.

::: bookkeeping
The following common
[`input`{#checkbox-state-(type=checkbox):the-input-element-3}](#the-input-element)
element content attributes and IDL attributes
[apply](#concept-input-apply){#checkbox-state-(type=checkbox):concept-input-apply}
to the element:
[`checked`{#checkbox-state-(type=checkbox):attr-input-checked}](#attr-input-checked),
and
[`required`{#checkbox-state-(type=checkbox):attr-input-required}](#attr-input-required)
content attributes;
[`checked`{#checkbox-state-(type=checkbox):dom-input-checked}](#dom-input-checked)
and
[`value`{#checkbox-state-(type=checkbox):dom-input-value}](#dom-input-value)
IDL attributes.

The
[`value`{#checkbox-state-(type=checkbox):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[default/on](#dom-input-value-default-on){#checkbox-state-(type=checkbox):dom-input-value-default-on}.

The
[`input`{#checkbox-state-(type=checkbox):event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#checkbox-state-(type=checkbox):event-change-2}](indices.html#event-change)
events
[apply](#concept-input-apply){#checkbox-state-(type=checkbox):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#checkbox-state-(type=checkbox):do-not-apply} to
the element:
[`accept`{#checkbox-state-(type=checkbox):attr-input-accept}](#attr-input-accept),
[`alpha`{#checkbox-state-(type=checkbox):attr-input-alpha}](#attr-input-alpha),
[`alt`{#checkbox-state-(type=checkbox):attr-input-alt}](#attr-input-alt),
[`autocomplete`{#checkbox-state-(type=checkbox):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`colorspace`{#checkbox-state-(type=checkbox):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#checkbox-state-(type=checkbox):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#checkbox-state-(type=checkbox):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#checkbox-state-(type=checkbox):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#checkbox-state-(type=checkbox):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#checkbox-state-(type=checkbox):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#checkbox-state-(type=checkbox):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#checkbox-state-(type=checkbox):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#checkbox-state-(type=checkbox):attr-input-list}](#attr-input-list),
[`max`{#checkbox-state-(type=checkbox):attr-input-max}](#attr-input-max),
[`maxlength`{#checkbox-state-(type=checkbox):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#checkbox-state-(type=checkbox):attr-input-min}](#attr-input-min),
[`minlength`{#checkbox-state-(type=checkbox):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#checkbox-state-(type=checkbox):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#checkbox-state-(type=checkbox):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#checkbox-state-(type=checkbox):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#checkbox-state-(type=checkbox):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#checkbox-state-(type=checkbox):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`readonly`{#checkbox-state-(type=checkbox):attr-input-readonly}](#attr-input-readonly),
[`size`{#checkbox-state-(type=checkbox):attr-input-size}](#attr-input-size),
[`src`{#checkbox-state-(type=checkbox):attr-input-src}](#attr-input-src),
[`step`{#checkbox-state-(type=checkbox):attr-input-step}](#attr-input-step),
and
[`width`{#checkbox-state-(type=checkbox):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#checkbox-state-(type=checkbox):do-not-apply-2} to
the element:
[`files`{#checkbox-state-(type=checkbox):dom-input-files}](#dom-input-files),
[`list`{#checkbox-state-(type=checkbox):dom-input-list}](#dom-input-list),
[`selectionStart`{#checkbox-state-(type=checkbox):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#checkbox-state-(type=checkbox):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#checkbox-state-(type=checkbox):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#checkbox-state-(type=checkbox):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#checkbox-state-(type=checkbox):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#checkbox-state-(type=checkbox):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#checkbox-state-(type=checkbox):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#checkbox-state-(type=checkbox):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#checkbox-state-(type=checkbox):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#checkbox-state-(type=checkbox):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.16]{.secno} [Radio Button]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=radio`)[](#radio-button-state-(type=radio)){.self-link} {#radio-button-state-(type=radio)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/radio](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/radio "<input> elements of type radio are generally used in radio groups—collections of radio buttons describing a set of related options.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

When an
[`input`{#radio-button-state-(type=radio):the-input-element}](#the-input-element)
element\'s
[`type`{#radio-button-state-(type=radio):attr-input-type}](#attr-input-type)
attribute is in the [Radio
Button](#radio-button-state-(type=radio)){#radio-button-state-(type=radio):radio-button-state-(type=radio)}
state, the rules in this section apply.

The
[`input`{#radio-button-state-(type=radio):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#radio-button-state-(type=radio):represents}
a control that, when used in conjunction with other
[`input`{#radio-button-state-(type=radio):the-input-element-3}](#the-input-element)
elements, forms a *[radio button group](#radio-button-group)* in which
only one control can have its
[checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked}
state set to true. If the element\'s
[checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked-2}
state is true, the control represents the selected control in the group,
and if it is false, it indicates a control in the group that is not
selected.

The [*radio button group*]{#radio-button-group .dfn} that contains an
[`input`{#radio-button-state-(type=radio):the-input-element-4}](#the-input-element)
element `a`{.variable} also contains all the other
[`input`{#radio-button-state-(type=radio):the-input-element-5}](#the-input-element)
elements `b`{.variable} that fulfill all of the following conditions:

- The
  [`input`{#radio-button-state-(type=radio):the-input-element-6}](#the-input-element)
  element `b`{.variable}\'s
  [`type`{#radio-button-state-(type=radio):attr-input-type-2}](#attr-input-type)
  attribute is in the [Radio
  Button](#radio-button-state-(type=radio)){#radio-button-state-(type=radio):radio-button-state-(type=radio)-2}
  state.
- Either `a`{.variable} and `b`{.variable} have the same [form
  owner](form-control-infrastructure.html#form-owner){#radio-button-state-(type=radio):form-owner},
  or they both have no [form
  owner](form-control-infrastructure.html#form-owner){#radio-button-state-(type=radio):form-owner-2}.
- Both `a`{.variable} and `b`{.variable} are in the same
  [tree](https://dom.spec.whatwg.org/#concept-tree){#radio-button-state-(type=radio):tree
  x-internal="tree"}.
- They both have a
  [`name`{#radio-button-state-(type=radio):attr-fe-name}](form-control-infrastructure.html#attr-fe-name)
  attribute, their
  [`name`{#radio-button-state-(type=radio):attr-fe-name-2}](form-control-infrastructure.html#attr-fe-name)
  attributes are not empty, and the value of `a`{.variable}\'s
  [`name`{#radio-button-state-(type=radio):attr-fe-name-3}](form-control-infrastructure.html#attr-fe-name)
  attribute equals the value of `b`{.variable}\'s
  [`name`{#radio-button-state-(type=radio):attr-fe-name-4}](form-control-infrastructure.html#attr-fe-name)
  attribute.

A
[tree](https://dom.spec.whatwg.org/#concept-tree){#radio-button-state-(type=radio):tree-2
x-internal="tree"} must not contain an
[`input`{#radio-button-state-(type=radio):the-input-element-7}](#the-input-element)
element whose *[radio button group](#radio-button-group)* contains only
that element.

When any of the following phenomena occur, if the element\'s
[checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked-3}
state is true after the occurrence, the
[checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked-4}
state of all the other elements in the same *[radio button
group](#radio-button-group)* must be set to false:

- The element\'s
  [checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked-5}
  state is set to true (for whatever reason).
- The element\'s
  [`name`{#radio-button-state-(type=radio):attr-fe-name-5}](form-control-infrastructure.html#attr-fe-name)
  attribute is set, changed, or removed.
- The element\'s [form
  owner](form-control-infrastructure.html#form-owner){#radio-button-state-(type=radio):form-owner-3}
  changes.
- [A type change is
  signalled](#signal-a-type-change){#radio-button-state-(type=radio):signal-a-type-change}
  for the element.
- The element [becomes
  connected](infrastructure.html#becomes-connected){#radio-button-state-(type=radio):becomes-connected}.

The [input activation
behavior](#input-activation-behavior){#radio-button-state-(type=radio):input-activation-behavior}
is to run the following steps:

1.  If the element is not
    [connected](https://dom.spec.whatwg.org/#connected){#radio-button-state-(type=radio):connected
    x-internal="connected"}, then return.

2.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#radio-button-state-(type=radio):concept-event-fire
    x-internal="concept-event-fire"} named
    [`input`{#radio-button-state-(type=radio):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    at the element with the
    [`bubbles`{#radio-button-state-(type=radio):dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    and
    [`composed`{#radio-button-state-(type=radio):dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
    attributes initialized to true.

3.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#radio-button-state-(type=radio):concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`change`{#radio-button-state-(type=radio):event-change}](indices.html#event-change)
    at the element with the
    [`bubbles`{#radio-button-state-(type=radio):dom-event-bubbles-2}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
    attribute initialized to true.

**Constraint validation**: If an element in the *[radio button
group](#radio-button-group)* is *[required](#concept-input-required)*,
and all of the
[`input`{#radio-button-state-(type=radio):the-input-element-8}](#the-input-element)
elements in the *[radio button group](#radio-button-group)* have a
[checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked-6}
that is false, then the element is [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#radio-button-state-(type=radio):suffering-from-being-missing}.

::: example
The following example, for some reason, has specified that puppers are
both
[required](#concept-input-required){#radio-button-state-(type=radio):concept-input-required-2}
and
[disabled](form-control-infrastructure.html#concept-fe-disabled){#radio-button-state-(type=radio):concept-fe-disabled}:

``` html
<form>
 <p><label><input type="radio" name="dog-type" value="pupper" required disabled> Pupper</label>
 <p><label><input type="radio" name="dog-type" value="doggo"> Doggo</label>
 <p><button>Make your choice</button>
</form>
```

If the user tries to submit this form without first selecting \"Doggo\",
then *both*
[`input`{#radio-button-state-(type=radio):the-input-element-9}](#the-input-element)
elements will be [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#radio-button-state-(type=radio):suffering-from-being-missing-2},
since an element in the [radio button
group](#radio-button-group){#radio-button-state-(type=radio):radio-button-group-6}
is
[required](#concept-input-required){#radio-button-state-(type=radio):concept-input-required-3}
(viz. the first element), and both of the elements in the radio button
group have a false
[checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked-7}.

On the other hand, if the user selects \"Doggo\" and then submits the
form, then neither
[`input`{#radio-button-state-(type=radio):the-input-element-10}](#the-input-element)
element will be [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#radio-button-state-(type=radio):suffering-from-being-missing-3},
since while one of them is
[required](#concept-input-required){#radio-button-state-(type=radio):concept-input-required-4},
not all of them have a false
[checkedness](form-control-infrastructure.html#concept-fe-checked){#radio-button-state-(type=radio):concept-fe-checked-8}.
:::

If none of the radio buttons in a [radio button
group](#radio-button-group){#radio-button-state-(type=radio):radio-button-group-7}
are checked, then they will all be initially unchecked in the interface,
until such time as one of them is checked (either by the user or by
script).

::: bookkeeping
The following common
[`input`{#radio-button-state-(type=radio):the-input-element-11}](#the-input-element)
element content attributes and IDL attributes
[apply](#concept-input-apply){#radio-button-state-(type=radio):concept-input-apply}
to the element:
[`checked`{#radio-button-state-(type=radio):attr-input-checked}](#attr-input-checked)
and
[`required`{#radio-button-state-(type=radio):attr-input-required}](#attr-input-required)
content attributes;
[`checked`{#radio-button-state-(type=radio):dom-input-checked}](#dom-input-checked)
and
[`value`{#radio-button-state-(type=radio):dom-input-value}](#dom-input-value)
IDL attributes.

The
[`value`{#radio-button-state-(type=radio):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[default/on](#dom-input-value-default-on){#radio-button-state-(type=radio):dom-input-value-default-on}.

The
[`input`{#radio-button-state-(type=radio):event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#radio-button-state-(type=radio):event-change-2}](indices.html#event-change)
events
[apply](#concept-input-apply){#radio-button-state-(type=radio):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#radio-button-state-(type=radio):do-not-apply} to
the element:
[`accept`{#radio-button-state-(type=radio):attr-input-accept}](#attr-input-accept),
[`alpha`{#radio-button-state-(type=radio):attr-input-alpha}](#attr-input-alpha),
[`alt`{#radio-button-state-(type=radio):attr-input-alt}](#attr-input-alt),
[`autocomplete`{#radio-button-state-(type=radio):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`colorspace`{#radio-button-state-(type=radio):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#radio-button-state-(type=radio):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#radio-button-state-(type=radio):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#radio-button-state-(type=radio):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#radio-button-state-(type=radio):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#radio-button-state-(type=radio):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#radio-button-state-(type=radio):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#radio-button-state-(type=radio):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#radio-button-state-(type=radio):attr-input-list}](#attr-input-list),
[`max`{#radio-button-state-(type=radio):attr-input-max}](#attr-input-max),
[`maxlength`{#radio-button-state-(type=radio):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#radio-button-state-(type=radio):attr-input-min}](#attr-input-min),
[`minlength`{#radio-button-state-(type=radio):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#radio-button-state-(type=radio):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#radio-button-state-(type=radio):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#radio-button-state-(type=radio):attr-input-placeholder}](#attr-input-placeholder),
[`popovertarget`{#radio-button-state-(type=radio):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#radio-button-state-(type=radio):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`readonly`{#radio-button-state-(type=radio):attr-input-readonly}](#attr-input-readonly),
[`size`{#radio-button-state-(type=radio):attr-input-size}](#attr-input-size),
[`src`{#radio-button-state-(type=radio):attr-input-src}](#attr-input-src),
[`step`{#radio-button-state-(type=radio):attr-input-step}](#attr-input-step),
and
[`width`{#radio-button-state-(type=radio):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#radio-button-state-(type=radio):do-not-apply-2}
to the element:
[`files`{#radio-button-state-(type=radio):dom-input-files}](#dom-input-files),
[`list`{#radio-button-state-(type=radio):dom-input-list}](#dom-input-list),
[`selectionStart`{#radio-button-state-(type=radio):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#radio-button-state-(type=radio):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#radio-button-state-(type=radio):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#radio-button-state-(type=radio):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#radio-button-state-(type=radio):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#radio-button-state-(type=radio):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#radio-button-state-(type=radio):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#radio-button-state-(type=radio):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#radio-button-state-(type=radio):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#radio-button-state-(type=radio):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.17]{.secno} [File Upload]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=file`)[](#file-upload-state-(type=file)){.self-link} {#file-upload-state-(type=file)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/file](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/file "<input> elements with type="file" let the user choose one or more files from their device storage. Once chosen, the files can be uploaded to a server using form submission, or manipulated using JavaScript code and the File API.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

When an
[`input`{#file-upload-state-(type=file):the-input-element}](#the-input-element)
element\'s
[`type`{#file-upload-state-(type=file):attr-input-type}](#attr-input-type)
attribute is in the [File
Upload](#file-upload-state-(type=file)){#file-upload-state-(type=file):file-upload-state-(type=file)}
state, the rules in this section apply.

The
[`input`{#file-upload-state-(type=file):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#file-upload-state-(type=file):represents}
a list of [selected files]{#concept-input-type-file-selected .dfn}, each
file consisting of a filename, a file type, and a file body (the
contents of the file).

Filenames must not contain [path
components](#concept-input-file-path){#file-upload-state-(type=file):concept-input-file-path},
even in the case that a user has selected an entire directory hierarchy
or multiple files with the same name from different directories. [Path
components]{#concept-input-file-path .dfn}, for the purposes of the
[File
Upload](#file-upload-state-(type=file)){#file-upload-state-(type=file):file-upload-state-(type=file)-2}
state, are those parts of filenames that are separated by U+005C REVERSE
SOLIDUS character (\\) characters.

Unless the
[`multiple`{#file-upload-state-(type=file):attr-input-multiple}](#attr-input-multiple)
attribute is set, there must be no more than one file in the list of
[selected
files](#concept-input-type-file-selected){#file-upload-state-(type=file):concept-input-type-file-selected}.

The [input activation
behavior](#input-activation-behavior){#file-upload-state-(type=file):input-activation-behavior}
for such an element `element`{.variable} is to [show the picker, if
applicable](#show-the-picker,-if-applicable){#file-upload-state-(type=file):show-the-picker,-if-applicable},
for `element`{.variable}.

If the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to change the files on the list in
other ways also, e.g., adding or removing files by drag-and-drop. When
the user does so, the user agent must [update the file
selection](#update-the-file-selection){#file-upload-state-(type=file):update-the-file-selection}
for the element.

If the element is not
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent must not allow the user to change the element\'s selection.

To [update the file selection]{#update-the-file-selection .dfn} for an
element `element`{.variable}:

1.  [Queue an element
    task](webappapis.html#queue-an-element-task){#file-upload-state-(type=file):queue-an-element-task}
    on the [user interaction task
    source](webappapis.html#user-interaction-task-source){#file-upload-state-(type=file):user-interaction-task-source}
    given `element`{.variable} and the following steps:

    1.  Update `element`{.variable}\'s [selected
        files](#concept-input-type-file-selected){#file-upload-state-(type=file):concept-input-type-file-selected-2}
        so that it represents the user\'s selection.

    2.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#file-upload-state-(type=file):concept-event-fire
        x-internal="concept-event-fire"} named
        [`input`{#file-upload-state-(type=file):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
        at the
        [`input`{#file-upload-state-(type=file):the-input-element-3}](#the-input-element)
        element, with the
        [`bubbles`{#file-upload-state-(type=file):dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
        and
        [`composed`{#file-upload-state-(type=file):dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
        attributes initialized to true.

    3.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#file-upload-state-(type=file):concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`change`{#file-upload-state-(type=file):event-change}](indices.html#event-change)
        at the
        [`input`{#file-upload-state-(type=file):the-input-element-4}](#the-input-element)
        element, with the
        [`bubbles`{#file-upload-state-(type=file):dom-event-bubbles-2}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
        attribute initialized to true.

**Constraint validation**: If the element is
*[required](#concept-input-required)* and the list of [selected
files](#concept-input-type-file-selected){#file-upload-state-(type=file):concept-input-type-file-selected-3}
is empty, then the element is [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#file-upload-state-(type=file):suffering-from-being-missing}.

------------------------------------------------------------------------

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/accept](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/accept "The accept attribute takes as its value a comma-separated list of one or more file types, or unique file type specifiers, describing which file types to allow.")

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

:::: feature
[Element/input#attr-accept](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#attr-accept "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox37+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome26+]{.chrome .yes}

------------------------------------------------------------------------

[Opera15+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android37+]{.firefox_android .yes}[Safari iOS11.3+]{.safari_ios
.yes}[Chrome Android26+]{.chrome_android .yes}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet1.5+]{.samsunginternet_android .yes}[Opera
Android15+]{.opera_android .yes}
:::
::::
:::::::

The [`accept`]{#attr-input-accept .dfn dfn-for="input"
dfn-type="element-attr"} attribute may be specified to provide user
agents with a hint of what file types will be accepted.

If specified, the attribute must consist of a [set of comma-separated
tokens](common-microsyntaxes.html#set-of-comma-separated-tokens){#file-upload-state-(type=file):set-of-comma-separated-tokens},
each of which must be an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#file-upload-state-(type=file):ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for one of the following:

The string \"`audio/*`\"
:   Indicates that sound files are accepted.

The string \"`video/*`\"
:   Indicates that video files are accepted.

The string \"`image/*`\"
:   Indicates that image files are accepted.

A [valid MIME type string with no parameters](https://mimesniff.spec.whatwg.org/#valid-mime-type-with-no-parameters){#file-upload-state-(type=file):valid-mime-type-string-with-no-parameters x-internal="valid-mime-type-string-with-no-parameters"}
:   Indicates that files of the specified type are accepted.

A string whose first character is a U+002E FULL STOP character (.)
:   Indicates that files with the specified file extension are accepted.

The tokens must not be [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#file-upload-state-(type=file):ascii-case-insensitive-2
x-internal="ascii-case-insensitive"} matches for any of the other tokens
(i.e. duplicates are not allowed). To obtain the list of tokens from the
attribute, the user agent must [split the attribute value on
commas](https://infra.spec.whatwg.org/#split-on-commas){#file-upload-state-(type=file):split-a-string-on-commas
x-internal="split-a-string-on-commas"}.

User agents may use the value of this attribute to display a more
appropriate user interface than a generic file picker. For instance,
given the value `image/*`, a user agent could offer the user the option
of using a local camera or selecting a photograph from their photo
collection; given the value `audio/*`, a user agent could offer the user
the option of recording a clip using a headset microphone.

User agents should prevent the user from selecting files that are not
accepted by one (or more) of these tokens.

Authors are encouraged to specify both any MIME types and any
corresponding extensions when looking for data in a specific format.

::: example
For example, consider an application that converts Microsoft Word
documents to Open Document Format files. Since Microsoft Word documents
are described with a wide variety of MIME types and extensions, the site
can list several, as follows:

``` html
<input type="file" accept=".doc,.docx,.xml,application/msword,application/vnd.openxmlformats-officedocument.wordprocessingml.document">
```

On platforms that only use file extensions to describe file types, the
extensions listed here can be used to filter the allowed documents,
while the MIME types can be used with the system\'s type registration
table (mapping MIME types to extensions used by the system), if any, to
determine any other extensions to allow. Similarly, on a system that
does not have filenames or extensions but labels documents with MIME
types internally, the MIME types can be used to pick the allowed files,
while the extensions can be used if the system has an extension
registration table that maps known extensions to MIME types used by the
system.
:::

Extensions tend to be ambiguous (e.g. there are an untold number of
formats that use the \"`.dat`\" extension, and users can typically quite
easily rename their files to have a \"`.doc`\" extension even if they
are not Microsoft Word documents), and MIME types tend to be unreliable
(e.g. many formats have no formally registered types, and many formats
are in practice labeled using a number of different MIME types). Authors
are reminded that, as usual, data received from a client should be
treated with caution, as it may not be in an expected format even if the
user is not hostile and the user agent fully obeyed the
[`accept`{#file-upload-state-(type=file):attr-input-accept}](#attr-input-accept)
attribute\'s requirements.

:::: {.mdn-anno .wrapped .before}
MDN

::: feature
[Element/input/file](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/file "<input> elements with type="file" let the user choose one or more files from their device storage. Once chosen, the files can be uploaded to a server using form submission, or manipulated using JavaScript code and the File API.")
:::
::::

::: {#fakepath-srsly .example}
For historical reasons, the
[`value`{#file-upload-state-(type=file):dom-input-value}](#dom-input-value)
IDL attribute prefixes the filename with the string \"`C:\fakepath\`\".
Some legacy user agents actually included the full path (which was a
security vulnerability). As a result of this, obtaining the filename
from the
[`value`{#file-upload-state-(type=file):dom-input-value-2}](#dom-input-value)
IDL attribute in a backwards-compatible way is non-trivial. The
following function extracts the filename in a suitably compatible
manner:

``` js
function extractFilename(path) {
  if (path.substr(0, 12) == "C:\\fakepath\\")
    return path.substr(12); // modern browser
  var x;
  x = path.lastIndexOf('/');
  if (x >= 0) // Unix-based path
    return path.substr(x+1);
  x = path.lastIndexOf('\\');
  if (x >= 0) // Windows-based path
    return path.substr(x+1);
  return path; // just the filename
}
```

This can be used as follows:

``` html
<p><input type=file name=image onchange="updateFilename(this.value)"></p>
<p>The name of the file you picked is: <span id="filename">(none)</span></p>
<script>
 function updateFilename(path) {
   var name = extractFilename(path);
   document.getElementById('filename').textContent = name;
 }
</script>
```
:::

------------------------------------------------------------------------

::: bookkeeping
The following common
[`input`{#file-upload-state-(type=file):the-input-element-5}](#the-input-element)
element content attributes and IDL attributes
[apply](#concept-input-apply){#file-upload-state-(type=file):concept-input-apply}
to the element:
[`accept`{#file-upload-state-(type=file):attr-input-accept-2}](#attr-input-accept),
[`multiple`{#file-upload-state-(type=file):attr-input-multiple-2}](#attr-input-multiple),
and
[`required`{#file-upload-state-(type=file):attr-input-required}](#attr-input-required)
content attributes;
[`files`{#file-upload-state-(type=file):dom-input-files}](#dom-input-files)
and
[`value`{#file-upload-state-(type=file):dom-input-value-3}](#dom-input-value)
IDL attributes;
[`select()`{#file-upload-state-(type=file):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select)
method.

The
[`value`{#file-upload-state-(type=file):dom-input-value-4}](#dom-input-value)
IDL attribute is in mode
[filename](#dom-input-value-filename){#file-upload-state-(type=file):dom-input-value-filename}.

The
[`input`{#file-upload-state-(type=file):event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#file-upload-state-(type=file):event-change-2}](indices.html#event-change)
events
[apply](#concept-input-apply){#file-upload-state-(type=file):concept-input-apply-2}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#file-upload-state-(type=file):do-not-apply} to
the element:
[`alpha`{#file-upload-state-(type=file):attr-input-alpha}](#attr-input-alpha),
[`alt`{#file-upload-state-(type=file):attr-input-alt}](#attr-input-alt),
[`autocomplete`{#file-upload-state-(type=file):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`checked`{#file-upload-state-(type=file):attr-input-checked}](#attr-input-checked),
[`colorspace`{#file-upload-state-(type=file):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#file-upload-state-(type=file):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#file-upload-state-(type=file):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#file-upload-state-(type=file):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#file-upload-state-(type=file):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#file-upload-state-(type=file):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#file-upload-state-(type=file):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#file-upload-state-(type=file):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#file-upload-state-(type=file):attr-input-list}](#attr-input-list),
[`max`{#file-upload-state-(type=file):attr-input-max}](#attr-input-max),
[`maxlength`{#file-upload-state-(type=file):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#file-upload-state-(type=file):attr-input-min}](#attr-input-min),
[`minlength`{#file-upload-state-(type=file):attr-input-minlength}](#attr-input-minlength),
[`pattern`{#file-upload-state-(type=file):attr-input-pattern}](#attr-input-pattern),
[`popovertarget`{#file-upload-state-(type=file):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#file-upload-state-(type=file):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`placeholder`{#file-upload-state-(type=file):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#file-upload-state-(type=file):attr-input-readonly}](#attr-input-readonly),
[`size`{#file-upload-state-(type=file):attr-input-size}](#attr-input-size),
[`src`{#file-upload-state-(type=file):attr-input-src}](#attr-input-src),
[`step`{#file-upload-state-(type=file):attr-input-step}](#attr-input-step),
and
[`width`{#file-upload-state-(type=file):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The element\'s
[`value`{#file-upload-state-(type=file):attr-input-value}](#attr-input-value)
attribute must be omitted.

The following IDL attributes and methods [do not
apply](#do-not-apply){#file-upload-state-(type=file):do-not-apply-2} to
the element:
[`checked`{#file-upload-state-(type=file):dom-input-checked}](#dom-input-checked),
[`list`{#file-upload-state-(type=file):dom-input-list}](#dom-input-list),
[`selectionStart`{#file-upload-state-(type=file):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#file-upload-state-(type=file):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#file-upload-state-(type=file):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#file-upload-state-(type=file):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#file-upload-state-(type=file):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`setRangeText()`{#file-upload-state-(type=file):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#file-upload-state-(type=file):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#file-upload-state-(type=file):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#file-upload-state-(type=file):dom-input-stepup}](#dom-input-stepup)
methods.
:::

###### [4.10.5.1.18]{.secno} [Submit Button]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=submit`)[](#submit-button-state-(type=submit)){.self-link} {#submit-button-state-(type=submit)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/submit](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/submit "<input> elements of type submit are rendered as buttons. When the click event occurs (typically because the user clicked the button), the user agent attempts to submit the form to the server.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

When an
[`input`{#submit-button-state-(type=submit):the-input-element}](#the-input-element)
element\'s
[`type`{#submit-button-state-(type=submit):attr-input-type}](#attr-input-type)
attribute is in the [Submit
Button](#submit-button-state-(type=submit)){#submit-button-state-(type=submit):submit-button-state-(type=submit)}
state, the rules in this section apply.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#submit-button-state-(type=submit):tracking-vector
.tracking-vector x-internal="tracking-vector"} The
[`input`{#submit-button-state-(type=submit):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#submit-button-state-(type=submit):represents}
a button that, when activated, submits the form. If the element has a
[`value`{#submit-button-state-(type=submit):attr-input-value}](#attr-input-value)
attribute, the button\'s label must be the value of that attribute;
otherwise, it must be an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#submit-button-state-(type=submit):implementation-defined
x-internal="implementation-defined"} string that means \"Submit\" or
some such. The element is a
[button](forms.html#concept-button){#submit-button-state-(type=submit):concept-button},
specifically a [submit
button](forms.html#concept-submit-button){#submit-button-state-(type=submit):concept-submit-button}.

Since the default label is
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#submit-button-state-(type=submit):implementation-defined-2
x-internal="implementation-defined"}, and the width of the button
typically depends on the button\'s label, the button\'s width can leak a
few bits of fingerprintable information. These bits are likely to be
strongly correlated to the identity of the user agent and the user\'s
locale.

The element\'s [optional
value](form-control-infrastructure.html#concept-fe-optional-value){#submit-button-state-(type=submit):concept-fe-optional-value}
is the value of the element\'s
[`value`{#submit-button-state-(type=submit):attr-button-value}](form-elements.html#attr-button-value)
attribute, if there is one; otherwise null.

The element\'s [input activation
behavior](#input-activation-behavior){#submit-button-state-(type=submit):input-activation-behavior}
given `event`{.variable} is as follows:

1.  If the element does not have a [form
    owner](form-control-infrastructure.html#form-owner){#submit-button-state-(type=submit):form-owner},
    then return.

2.  If the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#submit-button-state-(type=submit):node-document
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#submit-button-state-(type=submit):fully-active},
    then return.

3.  [Submit](form-control-infrastructure.html#concept-form-submit){#submit-button-state-(type=submit):concept-form-submit}
    the element\'s [form
    owner](form-control-infrastructure.html#form-owner){#submit-button-state-(type=submit):form-owner-2}
    from the element with
    *[userInvolvement](form-control-infrastructure.html#submit-user-involvement)*
    set to `event`{.variable}\'s [user navigation
    involvement](browsing-the-web.html#event-uni){#submit-button-state-(type=submit):event-uni}.

The
[`formaction`{#submit-button-state-(type=submit):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#submit-button-state-(type=submit):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#submit-button-state-(type=submit):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#submit-button-state-(type=submit):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
and
[`formtarget`{#submit-button-state-(type=submit):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget)
attributes are [attributes for form
submission](form-control-infrastructure.html#attributes-for-form-submission){#submit-button-state-(type=submit):attributes-for-form-submission}.

The
[`formnovalidate`{#submit-button-state-(type=submit):attr-fs-formnovalidate-2}](form-control-infrastructure.html#attr-fs-formnovalidate)
attribute can be used to make submit buttons that do not trigger the
constraint validation.

::: bookkeeping
The following common
[`input`{#submit-button-state-(type=submit):the-input-element-3}](#the-input-element)
element content attributes and IDL attributes
[apply](#concept-input-apply){#submit-button-state-(type=submit):concept-input-apply}
to the element:
[`dirname`{#submit-button-state-(type=submit):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#submit-button-state-(type=submit):attr-fs-formaction-2}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#submit-button-state-(type=submit):attr-fs-formenctype-2}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#submit-button-state-(type=submit):attr-fs-formmethod-2}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#submit-button-state-(type=submit):attr-fs-formnovalidate-3}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#submit-button-state-(type=submit):attr-fs-formtarget-2}](form-control-infrastructure.html#attr-fs-formtarget),
[`popovertarget`{#submit-button-state-(type=submit):attr-popovertarget}](popover.html#attr-popovertarget),
and
[`popovertargetaction`{#submit-button-state-(type=submit):attr-popovertargetaction}](popover.html#attr-popovertargetaction)
content attributes;
[`value`{#submit-button-state-(type=submit):dom-input-value}](#dom-input-value)
IDL attribute.

The
[`value`{#submit-button-state-(type=submit):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[default](#dom-input-value-default){#submit-button-state-(type=submit):dom-input-value-default}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#submit-button-state-(type=submit):do-not-apply}
to the element:
[`accept`{#submit-button-state-(type=submit):attr-input-accept}](#attr-input-accept),
[`alpha`{#submit-button-state-(type=submit):attr-input-alpha}](#attr-input-alpha),
[`alt`{#submit-button-state-(type=submit):attr-input-alt}](#attr-input-alt),
[`autocomplete`{#submit-button-state-(type=submit):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`checked`{#submit-button-state-(type=submit):attr-input-checked}](#attr-input-checked),
[`colorspace`{#submit-button-state-(type=submit):attr-input-colorspace}](#attr-input-colorspace),
[`height`{#submit-button-state-(type=submit):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#submit-button-state-(type=submit):attr-input-list}](#attr-input-list),
[`max`{#submit-button-state-(type=submit):attr-input-max}](#attr-input-max),
[`maxlength`{#submit-button-state-(type=submit):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#submit-button-state-(type=submit):attr-input-min}](#attr-input-min),
[`minlength`{#submit-button-state-(type=submit):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#submit-button-state-(type=submit):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#submit-button-state-(type=submit):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#submit-button-state-(type=submit):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#submit-button-state-(type=submit):attr-input-readonly}](#attr-input-readonly),
[`required`{#submit-button-state-(type=submit):attr-input-required}](#attr-input-required),
[`size`{#submit-button-state-(type=submit):attr-input-size}](#attr-input-size),
[`src`{#submit-button-state-(type=submit):attr-input-src}](#attr-input-src),
[`step`{#submit-button-state-(type=submit):attr-input-step}](#attr-input-step),
and
[`width`{#submit-button-state-(type=submit):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#submit-button-state-(type=submit):do-not-apply-2}
to the element:
[`checked`{#submit-button-state-(type=submit):dom-input-checked}](#dom-input-checked),
[`files`{#submit-button-state-(type=submit):dom-input-files}](#dom-input-files),
[`list`{#submit-button-state-(type=submit):dom-input-list}](#dom-input-list),
[`selectionStart`{#submit-button-state-(type=submit):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#submit-button-state-(type=submit):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#submit-button-state-(type=submit):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#submit-button-state-(type=submit):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#submit-button-state-(type=submit):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#submit-button-state-(type=submit):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#submit-button-state-(type=submit):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#submit-button-state-(type=submit):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#submit-button-state-(type=submit):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#submit-button-state-(type=submit):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`input`{#submit-button-state-(type=submit):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#submit-button-state-(type=submit):event-change}](indices.html#event-change)
events [do not
apply](#do-not-apply){#submit-button-state-(type=submit):do-not-apply-3}.
:::

###### [4.10.5.1.19]{.secno} [Image Button]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=image`)[](#image-button-state-(type=image)){.self-link} {#image-button-state-(type=image)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/image](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/image "<input> elements of type image are used to create graphical submit buttons, i.e. submit buttons that take the form of an image rather than text.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

When an
[`input`{#image-button-state-(type=image):the-input-element}](#the-input-element)
element\'s
[`type`{#image-button-state-(type=image):attr-input-type}](#attr-input-type)
attribute is in the [Image
Button](#image-button-state-(type=image)){#image-button-state-(type=image):image-button-state-(type=image)}
state, the rules in this section apply.

The
[`input`{#image-button-state-(type=image):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#image-button-state-(type=image):represents}
either an image from which a user can select a coordinate and submit the
form, or alternatively a button from which the user can submit the form.
The element is a
[button](forms.html#concept-button){#image-button-state-(type=image):concept-button},
specifically a [submit
button](forms.html#concept-submit-button){#image-button-state-(type=image):concept-submit-button}.

The coordinate is sent to the server [during form
submission](form-control-infrastructure.html#constructing-the-form-data-set){#image-button-state-(type=image):constructing-the-form-data-set}
by sending two entries for the element, derived from the name of the
control but with \"`.x`\" and \"`.y`\" appended to the name with the
`x`{.variable} and `y`{.variable} components of the coordinate
respectively.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#src](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#src "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

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

The image is given by the [`src`]{#attr-input-src .dfn dfn-for="input"
dfn-type="element-attr"} attribute. The
[`src`{#image-button-state-(type=image):attr-input-src}](#attr-input-src)
attribute must be present, and must contain a [valid non-empty URL
potentially surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#image-button-state-(type=image):valid-non-empty-url-potentially-surrounded-by-spaces}
referencing a non-interactive, optionally animated, image resource that
is neither paged nor scripted.

When any of the these events occur

- the
  [`input`{#image-button-state-(type=image):the-input-element-3}](#the-input-element)
  element\'s
  [`type`{#image-button-state-(type=image):attr-input-type-2}](#attr-input-type)
  attribute is first set to the [Image
  Button](#image-button-state-(type=image)){#image-button-state-(type=image):image-button-state-(type=image)-2}
  state (possibly when the element is first created), and the
  [`src`{#image-button-state-(type=image):attr-input-src-2}](#attr-input-src)
  attribute is present
- the
  [`input`{#image-button-state-(type=image):the-input-element-4}](#the-input-element)
  element\'s
  [`type`{#image-button-state-(type=image):attr-input-type-3}](#attr-input-type)
  attribute is changed back to the [Image
  Button](#image-button-state-(type=image)){#image-button-state-(type=image):image-button-state-(type=image)-3}
  state, and the
  [`src`{#image-button-state-(type=image):attr-input-src-3}](#attr-input-src)
  attribute is present, and its value has changed since the last time
  the
  [`type`{#image-button-state-(type=image):attr-input-type-4}](#attr-input-type)
  attribute was in the [Image
  Button](#image-button-state-(type=image)){#image-button-state-(type=image):image-button-state-(type=image)-4}
  state
- the
  [`input`{#image-button-state-(type=image):the-input-element-5}](#the-input-element)
  element\'s
  [`type`{#image-button-state-(type=image):attr-input-type-5}](#attr-input-type)
  attribute is in the [Image
  Button](#image-button-state-(type=image)){#image-button-state-(type=image):image-button-state-(type=image)-5}
  state, and the
  [`src`{#image-button-state-(type=image):attr-input-src-4}](#attr-input-src)
  attribute is set or changed

then unless the user agent cannot support images, or its support for
images has been disabled, or the user agent only fetches images on
demand, or the
[`src`{#image-button-state-(type=image):attr-input-src-5}](#attr-input-src)
attribute\'s value is the empty string, run these steps:

1.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#image-button-state-(type=image):encoding-parsing-a-url}
    given the
    [`src`{#image-button-state-(type=image):attr-input-src-6}](#attr-input-src)
    attribute\'s value, relative to the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#image-button-state-(type=image):node-document
    x-internal="node-document"}.

2.  If `url`{.variable} is failure, then return.

3.  Let `request`{.variable} be a new
    [request](https://fetch.spec.whatwg.org/#concept-request){#image-button-state-(type=image):concept-request
    x-internal="concept-request"} whose
    [URL](https://fetch.spec.whatwg.org/#concept-request-url){#image-button-state-(type=image):concept-request-url
    x-internal="concept-request-url"} is `url`{.variable},
    [client](https://fetch.spec.whatwg.org/#concept-request-client){#image-button-state-(type=image):concept-request-client
    x-internal="concept-request-client"} is the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#image-button-state-(type=image):node-document-2
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#image-button-state-(type=image):relevant-settings-object},
    [destination](https://fetch.spec.whatwg.org/#concept-request-destination){#image-button-state-(type=image):concept-request-destination
    x-internal="concept-request-destination"} is \"`image`\", [initiator
    type](https://fetch.spec.whatwg.org/#request-initiator-type){#image-button-state-(type=image):concept-request-initiator-type
    x-internal="concept-request-initiator-type"} is \"`input`\",
    [credentials
    mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#image-button-state-(type=image):concept-request-credentials-mode
    x-internal="concept-request-credentials-mode"} is \"`include`\", and
    whose [use-URL-credentials
    flag](https://fetch.spec.whatwg.org/#concept-request-use-url-credentials-flag){#image-button-state-(type=image):use-url-credentials-flag
    x-internal="use-url-credentials-flag"} is set.

4.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#image-button-state-(type=image):concept-fetch
    x-internal="concept-fetch"} `request`{.variable}, with
    *[processResponseEndOfBody](https://fetch.spec.whatwg.org/#fetch-processresponseendofbody){x-internal="processresponseendofbody"}*
    set to the following steps given
    [response](https://fetch.spec.whatwg.org/#concept-response){#image-button-state-(type=image):concept-response
    x-internal="concept-response"} `response`{.variable}:

    1.  If the download was successful and the image is
        *[available](#input-img-available)*, [queue an element
        task](webappapis.html#queue-an-element-task){#image-button-state-(type=image):queue-an-element-task}
        on the [user interaction task
        source](webappapis.html#user-interaction-task-source){#image-button-state-(type=image):user-interaction-task-source}
        given the
        [`input`{#image-button-state-(type=image):the-input-element-6}](#the-input-element)
        element to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#image-button-state-(type=image):concept-event-fire
        x-internal="concept-event-fire"} named
        [`load`{#image-button-state-(type=image):event-load}](indices.html#event-load)
        at the
        [`input`{#image-button-state-(type=image):the-input-element-7}](#the-input-element)
        element.

    2.  Otherwise, if the fetching process fails without a response from
        the remote server, or completes but the image is not a valid or
        supported image, then [queue an element
        task](webappapis.html#queue-an-element-task){#image-button-state-(type=image):queue-an-element-task-2}
        on the [user interaction task
        source](webappapis.html#user-interaction-task-source){#image-button-state-(type=image):user-interaction-task-source-2}
        given the
        [`input`{#image-button-state-(type=image):the-input-element-8}](#the-input-element)
        element to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#image-button-state-(type=image):concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`error`{#image-button-state-(type=image):event-error}](indices.html#event-error)
        on the
        [`input`{#image-button-state-(type=image):the-input-element-9}](#the-input-element)
        element.

Fetching the image must [delay the load
event](parsing.html#delay-the-load-event){#image-button-state-(type=image):delay-the-load-event}
of the element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#image-button-state-(type=image):node-document-3
x-internal="node-document"} until the
[task](webappapis.html#concept-task){#image-button-state-(type=image):concept-task}
that is
[queued](webappapis.html#queue-a-task){#image-button-state-(type=image):queue-a-task}
by the [networking task
source](webappapis.html#networking-task-source){#image-button-state-(type=image):networking-task-source}
once the resource has been fetched (defined below) has been run.

If the image was successfully obtained, with no network errors, and the
image\'s type is a supported image type, and the image is a valid image
of that type, then the image is said to be
[*available*]{#input-img-available .dfn}. If this is true before the
image is completely downloaded, each
[task](webappapis.html#concept-task){#image-button-state-(type=image):concept-task-2}
that is
[queued](webappapis.html#queue-a-task){#image-button-state-(type=image):queue-a-task-2}
by the [networking task
source](webappapis.html#networking-task-source){#image-button-state-(type=image):networking-task-source-2}
while the image is being fetched must update the presentation of the
image appropriately.

The user agent should apply the [image sniffing
rules](https://mimesniff.spec.whatwg.org/#rules-for-sniffing-images-specifically){#image-button-state-(type=image):content-type-sniffing:-image
x-internal="content-type-sniffing:-image"} to determine the type of the
image, with the image\'s [associated Content-Type
headers](urls-and-fetching.html#content-type){#image-button-state-(type=image):content-type}
giving the `official type`{.variable}. If these rules are not applied,
then the type of the image must be the type given by the image\'s
[associated Content-Type
headers](urls-and-fetching.html#content-type){#image-button-state-(type=image):content-type-2}.

User agents must not support non-image resources with the
[`input`{#image-button-state-(type=image):the-input-element-10}](#the-input-element)
element. User agents must not run executable code embedded in the image
resource. User agents must only display the first page of a multipage
resource. User agents must not allow the resource to act in an
interactive fashion, but should honor any animation in the resource.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#alt](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#alt "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

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

The [`alt`]{#attr-input-alt .dfn dfn-for="input"
dfn-type="element-attr"} attribute provides the textual label for the
button for users and user agents who cannot use the image. The
[`alt`{#image-button-state-(type=image):attr-input-alt}](#attr-input-alt)
attribute must be present, and must contain a non-empty string giving
the label that would be appropriate for an equivalent button if the
image was unavailable.

The
[`input`{#image-button-state-(type=image):the-input-element-11}](#the-input-element)
element supports [dimension
attributes](embedded-content-other.html#dimension-attributes){#image-button-state-(type=image):dimension-attributes}.

------------------------------------------------------------------------

If the
[`src`{#image-button-state-(type=image):attr-input-src-7}](#attr-input-src)
attribute is set, and the image is *[available](#input-img-available)*
and the user agent is configured to display that image, then the element
[represents](dom.html#represents){#image-button-state-(type=image):represents-2}
a control for selecting a
[coordinate](#concept-input-type-image-coordinate){#image-button-state-(type=image):concept-input-type-image-coordinate}
from the image specified by the
[`src`{#image-button-state-(type=image):attr-input-src-8}](#attr-input-src)
attribute. In that case, if the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, the
user agent should allow the user to select this
[coordinate](#concept-input-type-image-coordinate){#image-button-state-(type=image):concept-input-type-image-coordinate-2}.

Otherwise, the element
[represents](dom.html#represents){#image-button-state-(type=image):represents-3}
a submit button whose label is given by the value of the
[`alt`{#image-button-state-(type=image):attr-input-alt-2}](#attr-input-alt)
attribute.

The element\'s [input activation
behavior](#input-activation-behavior){#image-button-state-(type=image):input-activation-behavior}
given `event`{.variable} is as follows:

1.  If the element does not have a [form
    owner](form-control-infrastructure.html#form-owner){#image-button-state-(type=image):form-owner},
    then return.

2.  If the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#image-button-state-(type=image):node-document-4
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#image-button-state-(type=image):fully-active},
    then return.

3.  If the user activated the control while explicitly selecting a
    coordinate, then set the element\'s [selected
    coordinate](#concept-input-type-image-coordinate){#image-button-state-(type=image):concept-input-type-image-coordinate-3}
    to that coordinate.

    This is only possible under the conditions outlined above, when the
    element
    [represents](dom.html#represents){#image-button-state-(type=image):represents-4}
    a control for selecting such a coordinate. Even then, the user might
    activate the control without explicitly selecting a coordinate.

4.  [Submit](form-control-infrastructure.html#concept-form-submit){#image-button-state-(type=image):concept-form-submit}
    the element\'s [form
    owner](form-control-infrastructure.html#form-owner){#image-button-state-(type=image):form-owner-2}
    from the element with
    *[userInvolvement](form-control-infrastructure.html#submit-user-involvement)*
    set to `event`{.variable}\'s [user navigation
    involvement](browsing-the-web.html#event-uni){#image-button-state-(type=image):event-uni}.

The element\'s [selected
coordinate]{#concept-input-type-image-coordinate .dfn} consists of an
`x`{.variable}-component and a `y`{.variable}-component. It is initially
(0, 0). The coordinates represent the position relative to the edge of
the element\'s image, with the coordinate space having the positive
`x`{.variable} direction to the right, and the positive `y`{.variable}
direction downwards.

The `x`{.variable}-component must be a [valid
integer](common-microsyntaxes.html#valid-integer){#image-button-state-(type=image):valid-integer}
representing a number `x`{.variable} in the range
−(`border`{.variable}~`left`{.variable}~+`padding`{.variable}~`left`{.variable}~)
≤ `x`{.variable} ≤
`width`{.variable}+`border`{.variable}~`right`{.variable}~+`padding`{.variable}~`right`{.variable}~,
where `width`{.variable} is the rendered width of the image,
`border`{.variable}~`left`{.variable}~ is the width of the border on the
left of the image, `padding`{.variable}~`left`{.variable}~ is the width
of the padding on the left of the image,
`border`{.variable}~`right`{.variable}~ is the width of the border on
the right of the image, and `padding`{.variable}~`right`{.variable}~ is
the width of the padding on the right of the image, with all dimensions
given in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#image-button-state-(type=image):'px'
x-internal="'px'"}.

The `y`{.variable}-component must be a [valid
integer](common-microsyntaxes.html#valid-integer){#image-button-state-(type=image):valid-integer-2}
representing a number `y`{.variable} in the range
−(`border`{.variable}~`top`{.variable}~+`padding`{.variable}~`top`{.variable}~)
≤ `y`{.variable} ≤
`height`{.variable}+`border`{.variable}~`bottom`{.variable}~+`padding`{.variable}~`bottom`{.variable}~,
where `height`{.variable} is the rendered height of the image,
`border`{.variable}~`top`{.variable}~ is the width of the border above
the image, `padding`{.variable}~`top`{.variable}~ is the width of the
padding above the image, `border`{.variable}~`bottom`{.variable}~ is the
width of the border below the image, and
`padding`{.variable}~`bottom`{.variable}~ is the width of the padding
below the image, with all dimensions given in [CSS
pixels](https://drafts.csswg.org/css-values/#px){#image-button-state-(type=image):'px'-2
x-internal="'px'"}.

Where a border or padding is missing, its width is zero [CSS
pixels](https://drafts.csswg.org/css-values/#px){#image-button-state-(type=image):'px'-3
x-internal="'px'"}.

------------------------------------------------------------------------

The
[`formaction`{#image-button-state-(type=image):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#image-button-state-(type=image):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#image-button-state-(type=image):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#image-button-state-(type=image):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
and
[`formtarget`{#image-button-state-(type=image):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget)
attributes are [attributes for form
submission](form-control-infrastructure.html#attributes-for-form-submission){#image-button-state-(type=image):attributes-for-form-submission}.

`image`{.variable}`.`[`width`](#dom-input-width){#dom-input-width-dev}` [ = ``value`{.variable}` ]`\
`image`{.variable}`.`[`height`](#dom-input-height){#dom-input-height-dev}` [ = ``value`{.variable}` ]`

:   These attributes return the actual rendered dimensions of the image,
    or 0 if the dimensions are not known.

    They can be set, to change the corresponding content attributes.

::: bookkeeping
The following common
[`input`{#image-button-state-(type=image):the-input-element-12}](#the-input-element)
element content attributes and IDL attributes
[apply](#concept-input-apply){#image-button-state-(type=image):concept-input-apply}
to the element:
[`alt`{#image-button-state-(type=image):attr-input-alt-3}](#attr-input-alt),
[`formaction`{#image-button-state-(type=image):attr-fs-formaction-2}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#image-button-state-(type=image):attr-fs-formenctype-2}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#image-button-state-(type=image):attr-fs-formmethod-2}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#image-button-state-(type=image):attr-fs-formnovalidate-2}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#image-button-state-(type=image):attr-fs-formtarget-2}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#image-button-state-(type=image):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`popovertarget`{#image-button-state-(type=image):attr-popovertarget}](popover.html#attr-popovertarget),
[`popovertargetaction`{#image-button-state-(type=image):attr-popovertargetaction}](popover.html#attr-popovertargetaction),
[`src`{#image-button-state-(type=image):attr-input-src-9}](#attr-input-src),
and
[`width`{#image-button-state-(type=image):attr-dim-width}](embedded-content-other.html#attr-dim-width)
content attributes;
[`value`{#image-button-state-(type=image):dom-input-value}](#dom-input-value)
IDL attribute.

The
[`value`{#image-button-state-(type=image):dom-input-value-2}](#dom-input-value)
IDL attribute is in mode
[default](#dom-input-value-default){#image-button-state-(type=image):dom-input-value-default}.

The following content attributes must not be specified and [do not
apply](#do-not-apply){#image-button-state-(type=image):do-not-apply} to
the element:
[`accept`{#image-button-state-(type=image):attr-input-accept}](#attr-input-accept),
[`alpha`{#image-button-state-(type=image):attr-input-alpha}](#attr-input-alpha),
[`autocomplete`{#image-button-state-(type=image):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`checked`{#image-button-state-(type=image):attr-input-checked}](#attr-input-checked),
[`colorspace`{#image-button-state-(type=image):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#image-button-state-(type=image):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`list`{#image-button-state-(type=image):attr-input-list}](#attr-input-list),
[`max`{#image-button-state-(type=image):attr-input-max}](#attr-input-max),
[`maxlength`{#image-button-state-(type=image):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#image-button-state-(type=image):attr-input-min}](#attr-input-min),
[`minlength`{#image-button-state-(type=image):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#image-button-state-(type=image):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#image-button-state-(type=image):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#image-button-state-(type=image):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#image-button-state-(type=image):attr-input-readonly}](#attr-input-readonly),
[`required`{#image-button-state-(type=image):attr-input-required}](#attr-input-required),
[`size`{#image-button-state-(type=image):attr-input-size}](#attr-input-size),
and
[`step`{#image-button-state-(type=image):attr-input-step}](#attr-input-step).

The element\'s
[`value`{#image-button-state-(type=image):attr-input-value}](#attr-input-value)
attribute must be omitted.

The following IDL attributes and methods [do not
apply](#do-not-apply){#image-button-state-(type=image):do-not-apply-2}
to the element:
[`checked`{#image-button-state-(type=image):dom-input-checked}](#dom-input-checked),
[`files`{#image-button-state-(type=image):dom-input-files}](#dom-input-files),
[`list`{#image-button-state-(type=image):dom-input-list}](#dom-input-list),
[`selectionStart`{#image-button-state-(type=image):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#image-button-state-(type=image):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#image-button-state-(type=image):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#image-button-state-(type=image):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#image-button-state-(type=image):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#image-button-state-(type=image):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#image-button-state-(type=image):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#image-button-state-(type=image):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#image-button-state-(type=image):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#image-button-state-(type=image):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`input`{#image-button-state-(type=image):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#image-button-state-(type=image):event-change}](indices.html#event-change)
events [do not
apply](#do-not-apply){#image-button-state-(type=image):do-not-apply-3}.
:::

Many aspects of this state\'s behavior are similar to the behavior of
the
[`img`{#image-button-state-(type=image):the-img-element}](embedded-content.html#the-img-element)
element. Readers are encouraged to read that section, where many of the
same requirements are described in more detail.

::: example
Take the following form:

``` html
<form action="process.cgi">
 <input type=image src=map.png name=where alt="Show location list">
</form>
```

If the user clicked on the image at coordinate (127,40) then the URL
used to submit the form would be
\"`process.cgi?where.x=127&where.y=40`\".

(In this example, it\'s assumed that for users who don\'t see the map,
and who instead just see a button labeled \"Show location list\",
clicking the button will cause the server to show a list of locations to
pick from instead of the map.)
:::

###### [4.10.5.1.20]{.secno} [Reset Button]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=reset`)[](#reset-button-state-(type=reset)){.self-link} {#reset-button-state-(type=reset)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/reset](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/reset "<input> elements of type reset are rendered as buttons, with a default click event handler that resets all inputs in the form to their initial values.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

When an
[`input`{#reset-button-state-(type=reset):the-input-element}](#the-input-element)
element\'s
[`type`{#reset-button-state-(type=reset):attr-input-type}](#attr-input-type)
attribute is in the [Reset
Button](#reset-button-state-(type=reset)){#reset-button-state-(type=reset):reset-button-state-(type=reset)}
state, the rules in this section apply.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#reset-button-state-(type=reset):tracking-vector
.tracking-vector x-internal="tracking-vector"} The
[`input`{#reset-button-state-(type=reset):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#reset-button-state-(type=reset):represents}
a button that, when activated, resets the form. If the element has a
[`value`{#reset-button-state-(type=reset):attr-input-value}](#attr-input-value)
attribute, the button\'s label must be the value of that attribute;
otherwise, it must be an
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#reset-button-state-(type=reset):implementation-defined
x-internal="implementation-defined"} string that means \"Reset\" or some
such. The element is a
[button](forms.html#concept-button){#reset-button-state-(type=reset):concept-button}.

Since the default label is
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#reset-button-state-(type=reset):implementation-defined-2
x-internal="implementation-defined"}, and the width of the button
typically depends on the button\'s label, the button\'s width can leak a
few bits of fingerprintable information. These bits are likely to be
strongly correlated to the identity of the user agent and the user\'s
locale.

The element\'s [input activation
behavior](#input-activation-behavior){#reset-button-state-(type=reset):input-activation-behavior}
is as follows:

1.  If the element does not have a [form
    owner](form-control-infrastructure.html#form-owner){#reset-button-state-(type=reset):form-owner},
    then return.

2.  If the element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#reset-button-state-(type=reset):node-document
    x-internal="node-document"} is not [fully
    active](document-sequences.html#fully-active){#reset-button-state-(type=reset):fully-active},
    then return.

3.  [Reset](form-control-infrastructure.html#concept-form-reset){#reset-button-state-(type=reset):concept-form-reset}
    the [form
    owner](form-control-infrastructure.html#form-owner){#reset-button-state-(type=reset):form-owner-2}
    from the element.

**Constraint validation**: The element is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#reset-button-state-(type=reset):barred-from-constraint-validation}.

::: bookkeeping
The
[`value`{#reset-button-state-(type=reset):dom-input-value}](#dom-input-value)
IDL attribute
[applies](#concept-input-apply){#reset-button-state-(type=reset):concept-input-apply}
to this element and is in mode
[default](#dom-input-value-default){#reset-button-state-(type=reset):dom-input-value-default}.

The following common
[`input`{#reset-button-state-(type=reset):the-input-element-3}](#the-input-element)
element content attributes
[apply](#concept-input-apply){#reset-button-state-(type=reset):concept-input-apply-2}
to the element:
[`popovertarget`{#reset-button-state-(type=reset):attr-popovertarget}](popover.html#attr-popovertarget)
and
[`popovertargetaction`{#reset-button-state-(type=reset):attr-popovertargetaction}](popover.html#attr-popovertargetaction).

The following content attributes must not be specified and [do not
apply](#do-not-apply){#reset-button-state-(type=reset):do-not-apply} to
the element:
[`accept`{#reset-button-state-(type=reset):attr-input-accept}](#attr-input-accept),
[`alpha`{#reset-button-state-(type=reset):attr-input-alpha}](#attr-input-alpha),
[`alt`{#reset-button-state-(type=reset):attr-input-alt}](#attr-input-alt),
[`autocomplete`{#reset-button-state-(type=reset):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`checked`{#reset-button-state-(type=reset):attr-input-checked}](#attr-input-checked),
[`colorspace`{#reset-button-state-(type=reset):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#reset-button-state-(type=reset):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#reset-button-state-(type=reset):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#reset-button-state-(type=reset):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#reset-button-state-(type=reset):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#reset-button-state-(type=reset):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#reset-button-state-(type=reset):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#reset-button-state-(type=reset):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#reset-button-state-(type=reset):attr-input-list}](#attr-input-list),
[`max`{#reset-button-state-(type=reset):attr-input-max}](#attr-input-max),
[`maxlength`{#reset-button-state-(type=reset):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#reset-button-state-(type=reset):attr-input-min}](#attr-input-min),
[`minlength`{#reset-button-state-(type=reset):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#reset-button-state-(type=reset):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#reset-button-state-(type=reset):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#reset-button-state-(type=reset):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#reset-button-state-(type=reset):attr-input-readonly}](#attr-input-readonly),
[`required`{#reset-button-state-(type=reset):attr-input-required}](#attr-input-required),
[`size`{#reset-button-state-(type=reset):attr-input-size}](#attr-input-size),
[`src`{#reset-button-state-(type=reset):attr-input-src}](#attr-input-src),
[`step`{#reset-button-state-(type=reset):attr-input-step}](#attr-input-step),
and
[`width`{#reset-button-state-(type=reset):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#reset-button-state-(type=reset):do-not-apply-2}
to the element:
[`checked`{#reset-button-state-(type=reset):dom-input-checked}](#dom-input-checked),
[`files`{#reset-button-state-(type=reset):dom-input-files}](#dom-input-files),
[`list`{#reset-button-state-(type=reset):dom-input-list}](#dom-input-list),
[`selectionStart`{#reset-button-state-(type=reset):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#reset-button-state-(type=reset):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#reset-button-state-(type=reset):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#reset-button-state-(type=reset):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#reset-button-state-(type=reset):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#reset-button-state-(type=reset):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#reset-button-state-(type=reset):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#reset-button-state-(type=reset):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#reset-button-state-(type=reset):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#reset-button-state-(type=reset):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`input`{#reset-button-state-(type=reset):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#reset-button-state-(type=reset):event-change}](indices.html#event-change)
events [do not
apply](#do-not-apply){#reset-button-state-(type=reset):do-not-apply-3}.
:::

###### [4.10.5.1.21]{.secno} [Button]{.dfn dfn-for="input" dfn-type="element-state"} state (`type=button`)[](#button-state-(type=button)){.self-link} {#button-state-(type=button)}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/input/button](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/button "<input> elements of type button are rendered as simple push buttons, which can be programmed to control custom functionality anywhere on a webpage as required when assigned an event handler function (typically for the click event).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

When an
[`input`{#button-state-(type=button):the-input-element}](#the-input-element)
element\'s
[`type`{#button-state-(type=button):attr-input-type}](#attr-input-type)
attribute is in the
[Button](#button-state-(type=button)){#button-state-(type=button):button-state-(type=button)}
state, the rules in this section apply.

The
[`input`{#button-state-(type=button):the-input-element-2}](#the-input-element)
element
[represents](dom.html#represents){#button-state-(type=button):represents}
a button with no default behavior. A label for the button must be
provided in the
[`value`{#button-state-(type=button):attr-input-value}](#attr-input-value)
attribute, though it may be the empty string. If the element has a
[`value`{#button-state-(type=button):attr-input-value-2}](#attr-input-value)
attribute, the button\'s label must be the value of that attribute;
otherwise, it must be the empty string. The element is a
[button](forms.html#concept-button){#button-state-(type=button):concept-button}.

The element has no [input activation
behavior](#input-activation-behavior){#button-state-(type=button):input-activation-behavior}.

**Constraint validation**: The element is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#button-state-(type=button):barred-from-constraint-validation}.

::: bookkeeping
The
[`value`{#button-state-(type=button):dom-input-value}](#dom-input-value)
IDL attribute
[applies](#concept-input-apply){#button-state-(type=button):concept-input-apply}
to this element and is in mode
[default](#dom-input-value-default){#button-state-(type=button):dom-input-value-default}.

The following common
[`input`{#button-state-(type=button):the-input-element-3}](#the-input-element)
element content attributes
[apply](#concept-input-apply){#button-state-(type=button):concept-input-apply-2}
to the element:
[`popovertarget`{#button-state-(type=button):attr-popovertarget}](popover.html#attr-popovertarget)
and
[`popovertargetaction`{#button-state-(type=button):attr-popovertargetaction}](popover.html#attr-popovertargetaction).

The following content attributes must not be specified and [do not
apply](#do-not-apply){#button-state-(type=button):do-not-apply} to the
element:
[`accept`{#button-state-(type=button):attr-input-accept}](#attr-input-accept),
[`alpha`{#button-state-(type=button):attr-input-alpha}](#attr-input-alpha),
[`alt`{#button-state-(type=button):attr-input-alt}](#attr-input-alt),
[`autocomplete`{#button-state-(type=button):attr-fe-autocomplete}](form-control-infrastructure.html#attr-fe-autocomplete),
[`checked`{#button-state-(type=button):attr-input-checked}](#attr-input-checked),
[`colorspace`{#button-state-(type=button):attr-input-colorspace}](#attr-input-colorspace),
[`dirname`{#button-state-(type=button):attr-fe-dirname}](form-control-infrastructure.html#attr-fe-dirname),
[`formaction`{#button-state-(type=button):attr-fs-formaction}](form-control-infrastructure.html#attr-fs-formaction),
[`formenctype`{#button-state-(type=button):attr-fs-formenctype}](form-control-infrastructure.html#attr-fs-formenctype),
[`formmethod`{#button-state-(type=button):attr-fs-formmethod}](form-control-infrastructure.html#attr-fs-formmethod),
[`formnovalidate`{#button-state-(type=button):attr-fs-formnovalidate}](form-control-infrastructure.html#attr-fs-formnovalidate),
[`formtarget`{#button-state-(type=button):attr-fs-formtarget}](form-control-infrastructure.html#attr-fs-formtarget),
[`height`{#button-state-(type=button):attr-dim-height}](embedded-content-other.html#attr-dim-height),
[`list`{#button-state-(type=button):attr-input-list}](#attr-input-list),
[`max`{#button-state-(type=button):attr-input-max}](#attr-input-max),
[`maxlength`{#button-state-(type=button):attr-input-maxlength}](#attr-input-maxlength),
[`min`{#button-state-(type=button):attr-input-min}](#attr-input-min),
[`minlength`{#button-state-(type=button):attr-input-minlength}](#attr-input-minlength),
[`multiple`{#button-state-(type=button):attr-input-multiple}](#attr-input-multiple),
[`pattern`{#button-state-(type=button):attr-input-pattern}](#attr-input-pattern),
[`placeholder`{#button-state-(type=button):attr-input-placeholder}](#attr-input-placeholder),
[`readonly`{#button-state-(type=button):attr-input-readonly}](#attr-input-readonly),
[`required`{#button-state-(type=button):attr-input-required}](#attr-input-required),
[`size`{#button-state-(type=button):attr-input-size}](#attr-input-size),
[`src`{#button-state-(type=button):attr-input-src}](#attr-input-src),
[`step`{#button-state-(type=button):attr-input-step}](#attr-input-step),
and
[`width`{#button-state-(type=button):attr-dim-width}](embedded-content-other.html#attr-dim-width).

The following IDL attributes and methods [do not
apply](#do-not-apply){#button-state-(type=button):do-not-apply-2} to the
element:
[`checked`{#button-state-(type=button):dom-input-checked}](#dom-input-checked),
[`files`{#button-state-(type=button):dom-input-files}](#dom-input-files),
[`list`{#button-state-(type=button):dom-input-list}](#dom-input-list),
[`selectionStart`{#button-state-(type=button):dom-textarea/input-selectionstart}](form-control-infrastructure.html#dom-textarea/input-selectionstart),
[`selectionEnd`{#button-state-(type=button):dom-textarea/input-selectionend}](form-control-infrastructure.html#dom-textarea/input-selectionend),
[`selectionDirection`{#button-state-(type=button):dom-textarea/input-selectiondirection}](form-control-infrastructure.html#dom-textarea/input-selectiondirection),
[`valueAsDate`{#button-state-(type=button):dom-input-valueasdate}](#dom-input-valueasdate),
and
[`valueAsNumber`{#button-state-(type=button):dom-input-valueasnumber}](#dom-input-valueasnumber)
IDL attributes;
[`select()`{#button-state-(type=button):dom-textarea/input-select}](form-control-infrastructure.html#dom-textarea/input-select),
[`setRangeText()`{#button-state-(type=button):dom-textarea/input-setrangetext}](form-control-infrastructure.html#dom-textarea/input-setrangetext),
[`setSelectionRange()`{#button-state-(type=button):dom-textarea/input-setselectionrange}](form-control-infrastructure.html#dom-textarea/input-setselectionrange),
[`stepDown()`{#button-state-(type=button):dom-input-stepdown}](#dom-input-stepdown),
and
[`stepUp()`{#button-state-(type=button):dom-input-stepup}](#dom-input-stepup)
methods.

The
[`input`{#button-state-(type=button):event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#button-state-(type=button):event-change}](indices.html#event-change)
events [do not
apply](#do-not-apply){#button-state-(type=button):do-not-apply-3}.
:::

##### [4.10.5.2]{.secno} Implementation notes regarding localization of form controls[](#input-impl-notes){.self-link} {#input-impl-notes}

*This section is non-normative.*

The formats shown to the user in date, time, and number controls is
independent of the format used for form submission.

Browsers are encouraged to use user interfaces that present dates,
times, and numbers according to the conventions of either the locale
implied by the
[`input`{#input-impl-notes:the-input-element}](#the-input-element)
element\'s [language](dom.html#language){#input-impl-notes:language} or
the user\'s preferred locale. Using the page\'s locale will ensure
consistency with page-provided data.

For example, it would be confusing to users if an American English page
claimed that a Cirque De Soleil show was going to be showing on 02/03,
but their browser, configured to use the British English locale, only
showed the date 03/02 in the ticket purchase date picker. Using the
page\'s locale would at least ensure that the date was presented in the
same format everywhere. (There\'s still a risk that the user would end
up arriving a month late, of course, but there\'s only so much that can
be done about such cultural differences\...)

##### [4.10.5.3]{.secno} Common [`input`{#common-input-element-attributes:the-input-element}](#the-input-element) element attributes[](#common-input-element-attributes){.self-link}

These attributes only
[apply](#concept-input-apply){#common-input-element-attributes:concept-input-apply}
to an
[`input`{#common-input-element-attributes:the-input-element-2}](#the-input-element)
element if its
[`type`{#common-input-element-attributes:attr-input-type}](#attr-input-type)
attribute is in a state whose definition declares that the attribute
[applies](#concept-input-apply){#common-input-element-attributes:concept-input-apply-2}.
When an attribute [doesn\'t
apply](#do-not-apply){#common-input-element-attributes:do-not-apply} to
an
[`input`{#common-input-element-attributes:the-input-element-3}](#the-input-element)
element, user agents must
[ignore](infrastructure.html#ignore){#common-input-element-attributes:ignore}
the attribute, regardless of the requirements and definitions below.

###### [4.10.5.3.1]{.secno} The [`maxlength`{#the-maxlength-and-minlength-attributes:attr-input-maxlength}](#attr-input-maxlength) and [`minlength`{#the-maxlength-and-minlength-attributes:attr-input-minlength}](#attr-input-minlength) attributes[](#the-maxlength-and-minlength-attributes){.self-link}

::::::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Attributes/maxlength](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/maxlength "The maxlength attribute defines the maximum string length that the user can enter into an <input> or <textarea>. The attribute must have an integer value of 0 or higher.")

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
[Attributes/minlength](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/minlength "The minlength attribute defines the minimum string length that the user can enter into an <input> or <textarea>. The attribute must have an integer value of 0 or higher.")

Support in all current engines.

::: support
[Firefox51+]{.firefox .yes}[Safari10.1+]{.safari
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
[Attributes/maxlength](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/maxlength "The maxlength attribute defines the maximum string length that the user can enter into an <input> or <textarea>. The attribute must have an integer value of 0 or higher.")

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

:::: feature
[Attributes/minlength](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/minlength "The minlength attribute defines the minimum string length that the user can enter into an <input> or <textarea>. The attribute must have an integer value of 0 or higher.")

Support in all current engines.

::: support
[Firefox51+]{.firefox .yes}[Safari10.1+]{.safari
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
:::::::::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#attr-maxlength](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#attr-maxlength "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera15+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android5+]{.firefox_android .yes}[Safari iOS9+]{.safari_ios
.yes}[Chrome Android18+]{.chrome_android .yes}[WebView
Android4.4+]{.webview_android .yes}[Samsung
Internet1.5+]{.samsunginternet_android .yes}[Opera
Android15+]{.opera_android .yes}

------------------------------------------------------------------------

[[caniuse.com
table](https://caniuse.com/#feat=maxlength "maxlength attribute for input and textarea elements")]{.caniuse}
:::
::::
:::::

The [`maxlength`]{#attr-input-maxlength .dfn dfn-for="input"
dfn-type="element-attr"} attribute, when it
[applies](#concept-input-apply){#the-maxlength-and-minlength-attributes:concept-input-apply},
is a [form control `maxlength`
attribute](form-control-infrastructure.html#attr-fe-maxlength){#the-maxlength-and-minlength-attributes:attr-fe-maxlength}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#attr-minlength](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#attr-minlength "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox51+]{.firefox .yes}[Safari10.1+]{.safari
.yes}[Chrome40+]{.chrome .yes}

------------------------------------------------------------------------

[Opera27+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android51+]{.firefox_android .yes}[Safari iOS10.3+]{.safari_ios
.yes}[Chrome Android40+]{.chrome_android .yes}[WebView
Android40+]{.webview_android .yes}[Samsung
Internet4.0+]{.samsunginternet_android .yes}[Opera
Android27+]{.opera_android .yes}

------------------------------------------------------------------------

[[caniuse.com
table](https://caniuse.com/#feat=input-minlength "Minimum length attribute for input fields")]{.caniuse}
:::
::::
:::::

The [`minlength`]{#attr-input-minlength .dfn dfn-for="input"
dfn-type="element-attr"} attribute, when it
[applies](#concept-input-apply){#the-maxlength-and-minlength-attributes:concept-input-apply-2},
is a [form control `minlength`
attribute](form-control-infrastructure.html#attr-fe-minlength){#the-maxlength-and-minlength-attributes:attr-fe-minlength}.

If the
[`input`{#the-maxlength-and-minlength-attributes:the-input-element}](#the-input-element)
element has a [maximum allowed value
length](form-control-infrastructure.html#maximum-allowed-value-length){#the-maxlength-and-minlength-attributes:maximum-allowed-value-length},
then the
[length](https://infra.spec.whatwg.org/#string-length){#the-maxlength-and-minlength-attributes:length
x-internal="length"} of the value of the element\'s
[`value`{#the-maxlength-and-minlength-attributes:attr-input-value}](#attr-input-value)
attribute must be less than or equal to the element\'s [maximum allowed
value
length](form-control-infrastructure.html#maximum-allowed-value-length){#the-maxlength-and-minlength-attributes:maximum-allowed-value-length-2}.

::: example
The following extract shows how a messaging client\'s text entry could
be arbitrarily restricted to a fixed number of characters, thus forcing
any conversation through this medium to be terse and discouraging
intelligent discourse.

``` html
<label>What are you doing? <input name=status maxlength=140></label>
```
:::

::: example
Here, a password is given a minimum length:

``` html
<p><label>Username: <input name=u required></label>
<p><label>Password: <input name=p required minlength=12></label>
```
:::

###### [4.10.5.3.2]{.secno} The [`size`{#the-size-attribute:attr-input-size}](#attr-input-size) attribute[](#the-size-attribute){.self-link}

The [`size`]{#attr-input-size .dfn dfn-for="input"
dfn-type="element-attr"} attribute gives the number of characters that,
in a visual rendering, the user agent is to allow the user to see while
editing the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-size-attribute:concept-fe-value}.

The [`size`{#the-size-attribute:attr-input-size-2}](#attr-input-size)
attribute, if specified, must have a value that is a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#the-size-attribute:valid-non-negative-integer}
greater than zero.

If the attribute is present, then its value must be parsed using the
[rules for parsing non-negative
integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#the-size-attribute:rules-for-parsing-non-negative-integers},
and if the result is a number greater than zero, then the user agent
should ensure that at least that many characters are visible.

The [`size`{#the-size-attribute:dom-input-size}](#dom-input-size) IDL
attribute is [limited to only positive
numbers](common-dom-interfaces.html#limited-to-only-non-negative-numbers-greater-than-zero){#the-size-attribute:limited-to-only-non-negative-numbers-greater-than-zero}
and has a [default
value](common-dom-interfaces.html#default-value){#the-size-attribute:default-value}
of 20.

###### [4.10.5.3.3]{.secno} The [`readonly`{#the-readonly-attribute:attr-input-readonly}](#attr-input-readonly) attribute[](#the-readonly-attribute){.self-link}

::::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Attributes/readonly](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/readonly "The Boolean readonly attribute, when present, makes the element not mutable, meaning the user can not edit the control.")

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
[Attributes/readonly](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/readonly "The Boolean readonly attribute, when present, makes the element not mutable, meaning the user can not edit the control.")

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

:::: feature
[Attributes/readonly](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/readonly "The Boolean readonly attribute, when present, makes the element not mutable, meaning the user can not edit the control.")

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
:::::::::

The [`readonly`]{#attr-input-readonly .dfn dfn-for="input"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-readonly-attribute:boolean-attribute}
that controls whether or not the user can edit the form control. When
specified, the element is not
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*.

**Constraint validation**: If the
[`readonly`{#the-readonly-attribute:attr-input-readonly-2}](#attr-input-readonly)
attribute is specified on an
[`input`{#the-readonly-attribute:the-input-element}](#the-input-element)
element, the element is [barred from constraint
validation](form-control-infrastructure.html#barred-from-constraint-validation){#the-readonly-attribute:barred-from-constraint-validation}.

::: note
The difference between
[`disabled`{#the-readonly-attribute:attr-fe-disabled}](form-control-infrastructure.html#attr-fe-disabled)
and
[`readonly`{#the-readonly-attribute:attr-input-readonly-3}](#attr-input-readonly)
is that read-only controls can still function, whereas disabled controls
generally do not function as controls until they are enabled. This is
spelled out in more detail elsewhere in this specification with
normative requirements that refer to the
[disabled](form-control-infrastructure.html#concept-fe-disabled){#the-readonly-attribute:concept-fe-disabled}
concept (for example, the element\'s [activation
behavior](https://dom.spec.whatwg.org/#eventtarget-activation-behavior){#the-readonly-attribute:activation-behaviour
x-internal="activation-behaviour"}, whether or not it is a [focusable
area](interaction.html#focusable-area){#the-readonly-attribute:focusable-area},
or when [constructing the entry
list](form-control-infrastructure.html#constructing-the-form-data-set){#the-readonly-attribute:constructing-the-form-data-set}).
Any other behavior related to user interaction with disabled controls,
such as whether text can be selected or copied, is not defined in this
standard.

Only text controls can be made read-only, since for other controls (such
as checkboxes and buttons) there is no useful distinction between being
read-only and being disabled, so the
[`readonly`{#the-readonly-attribute:attr-input-readonly-4}](#attr-input-readonly)
attribute [does not
apply](#do-not-apply){#the-readonly-attribute:do-not-apply}.
:::

::: example
In the following example, the existing product identifiers cannot be
modified, but they are still displayed as part of the form, for
consistency with the row representing a new product (where the
identifier is not yet filled in).

``` html
<form action="products.cgi" method="post" enctype="multipart/form-data">
 <table>
  <tr> <th> Product ID <th> Product name <th> Price <th> Action
  <tr>
   <td> <input readonly="readonly" name="1.pid" value="H412">
   <td> <input required="required" name="1.pname" value="Floor lamp Ulke">
   <td> $<input required="required" type="number" min="0" step="0.01" name="1.pprice" value="49.99">
   <td> <button formnovalidate="formnovalidate" name="action" value="delete:1">Delete</button>
  <tr>
   <td> <input readonly="readonly" name="2.pid" value="FG28">
   <td> <input required="required" name="2.pname" value="Table lamp Ulke">
   <td> $<input required="required" type="number" min="0" step="0.01" name="2.pprice" value="24.99">
   <td> <button formnovalidate="formnovalidate" name="action" value="delete:2">Delete</button>
  <tr>
   <td> <input required="required" name="3.pid" value="" pattern="[A-Z0-9]+">
   <td> <input required="required" name="3.pname" value="">
   <td> $<input required="required" type="number" min="0" step="0.01" name="3.pprice" value="">
   <td> <button formnovalidate="formnovalidate" name="action" value="delete:3">Delete</button>
 </table>
 <p> <button formnovalidate="formnovalidate" name="action" value="add">Add</button> </p>
 <p> <button name="action" value="update">Save</button> </p>
</form>
```
:::

###### [4.10.5.3.4]{.secno} The [`required`{#the-required-attribute:attr-input-required}](#attr-input-required) attribute[](#the-required-attribute){.self-link}

The [`required`]{#attr-input-required .dfn dfn-for="input"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-required-attribute:boolean-attribute}.
When specified, the element is [*required*]{#concept-input-required
.dfn}.

**Constraint validation**: If the element is
*[required](#concept-input-required)*, and its
[`value`{#the-required-attribute:dom-input-value}](#dom-input-value) IDL
attribute
[applies](#concept-input-apply){#the-required-attribute:concept-input-apply}
and is in the mode
[value](#dom-input-value-value){#the-required-attribute:dom-input-value-value},
and the element is
*[mutable](form-control-infrastructure.html#concept-fe-mutable)*, and
the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-required-attribute:concept-fe-value}
is the empty string, then the element is [suffering from being
missing](form-control-infrastructure.html#suffering-from-being-missing){#the-required-attribute:suffering-from-being-missing}.

::: example
The following form has two required fields, one for an email address and
one for a password. It also has a third field that is only considered
valid if the user types the same password in the password field and this
third field.

``` html
<h1>Create new account</h1>
<form action="/newaccount" method=post
      oninput="up2.setCustomValidity(up2.value != up.value ? 'Passwords do not match.' : '')">
 <p>
  <label for="username">Email address:</label>
  <input id="username" type=email required name=un>
 <p>
  <label for="password1">Password:</label>
  <input id="password1" type=password required name=up>
 <p>
  <label for="password2">Confirm password:</label>
  <input id="password2" type=password name=up2>
 <p>
  <input type=submit value="Create account">
</form>
```
:::

::: example
For radio buttons, the
[`required`{#the-required-attribute:attr-input-required-2}](#attr-input-required)
attribute is satisfied if any of the radio buttons in the
[group](#radio-button-group){#the-required-attribute:radio-button-group}
is selected. Thus, in the following example, any of the radio buttons
can be checked, not just the one marked as required:

``` html
<fieldset>
 <legend>Did the movie pass the Bechdel test?</legend>
 <p><label><input type="radio" name="bechdel" value="no-characters"> No, there are not even two female characters in the movie. </label>
 <p><label><input type="radio" name="bechdel" value="no-names"> No, the female characters never talk to each other. </label>
 <p><label><input type="radio" name="bechdel" value="no-topic"> No, when female characters talk to each other it's always about a male character. </label>
 <p><label><input type="radio" name="bechdel" value="yes" required> Yes. </label>
 <p><label><input type="radio" name="bechdel" value="unknown"> I don't know. </label>
</fieldset>
```

To avoid confusion as to whether a [radio button
group](#radio-button-group){#the-required-attribute:radio-button-group-2}
is required or not, authors are encouraged to specify the attribute on
all the radio buttons in a group. Indeed, in general, authors are
encouraged to avoid having radio button groups that do not have any
initially checked controls in the first place, as this is a state that
the user cannot return to, and is therefore generally considered a poor
user interface.
:::

###### [4.10.5.3.5]{.secno} The [`multiple`{#the-multiple-attribute:attr-input-multiple}](#attr-input-multiple) attribute[](#the-multiple-attribute){.self-link}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/multiple](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/multiple "The Boolean multiple attribute, if set, means the form control accepts one or more values. Valid for the email and file input types and the <select>, the manner by which the user opts for multiple values depends on the form control.")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
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
[Element/input#attr-multiple](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#attr-multiple "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera≤12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android≤37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android≤12.1+]{.opera_android .yes}
:::
::::
:::::::

The [`multiple`]{#attr-input-multiple .dfn dfn-for="input"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-multiple-attribute:boolean-attribute}
that indicates whether the user is to be allowed to specify more than
one value.

::: example
The following extract shows how an email client\'s \"To\" field could
accept multiple email addresses.

``` html
<label>To: <input type=email multiple name=to></label>
```

If the user had, amongst many friends in their user contacts database,
two friends \"Spider-Man\" (with address \"spider@parker.example.net\")
and \"Scarlet Witch\" (with address \"scarlet@avengers.example.net\"),
then, after the user has typed \"[s]{.kbd}\", the user agent might
suggest these two email addresses to the user.

![](/images/sample-email-1.svg){.darkmode-aware width="330"
height="100"}

The page could also link in the user\'s contacts database from the site:

``` html
<label>To: <input type=email multiple name=to list=contacts></label>
...
<datalist id="contacts">
 <option value="hedral@damowmow.com">
 <option value="pillar@example.com">
 <option value="astrophy@cute.example">
 <option value="astronomy@science.example.org">
</datalist>
```

Suppose the user had entered \"[bob@example.net]{.kbd}\" into this text
control, and then started typing a second email address starting with
\"[s]{.kbd}\". The user agent might show both the two friends mentioned
earlier, as well as the \"astrophy\" and \"astronomy\" values given in
the
[`datalist`{#the-multiple-attribute:the-datalist-element}](form-elements.html#the-datalist-element)
element.

![](/images/sample-email-2.svg){.darkmode-aware width="330"
height="140"}
:::

::: example
The following extract shows how an email client\'s \"Attachments\" field
could accept multiple files for upload.

``` html
<label>Attachments: <input type=file multiple name=att></label>
```
:::

###### [4.10.5.3.6]{.secno} The [`pattern`{#the-pattern-attribute:attr-input-pattern}](#attr-input-pattern) attribute[](#the-pattern-attribute){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Attributes/pattern](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/pattern "The pattern attribute specifies a regular expression the form control's value should match. If a non-null value doesn't conform to the constraints set by the pattern value, the ValidityState object's read-only patternMismatch property will be true.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera≤12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android≤37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android≤12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/pattern](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/pattern "The pattern attribute specifies a regular expression the form control's value should match. If a non-null value doesn't conform to the constraints set by the pattern value, the ValidityState object's read-only patternMismatch property will be true.")

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

The [`pattern`]{#attr-input-pattern .dfn dfn-for="input"
dfn-type="element-attr"} attribute specifies a regular expression
against which the control\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-pattern-attribute:concept-fe-value},
or, when the
[`multiple`{#the-pattern-attribute:attr-input-multiple}](#attr-input-multiple)
attribute
[applies](#concept-input-apply){#the-pattern-attribute:concept-input-apply}
and is set, the control\'s
[values](form-control-infrastructure.html#concept-fe-values){#the-pattern-attribute:concept-fe-values},
are to be checked.

If specified, the attribute\'s value must match the JavaScript
*[Pattern](https://tc39.es/ecma262/#prod-Pattern){x-internal="js-prod-pattern"}*~`[+UnicodeSetsMode, +N]`~
production.

The [compiled pattern regular
expression]{#compiled-pattern-regular-expression .dfn} of an
[`input`{#the-pattern-attribute:the-input-element}](#the-input-element)
element, if it exists, is a JavaScript
[`RegExp`{#the-pattern-attribute:regexp}](https://tc39.es/ecma262/#sec-regexp-regular-expression-objects){x-internal="regexp"}
object. It is determined as follows:

1.  If the element does not have a
    [`pattern`{#the-pattern-attribute:attr-input-pattern-2}](#attr-input-pattern)
    attribute specified, then return nothing. The element has no
    [compiled pattern regular
    expression](#compiled-pattern-regular-expression){#the-pattern-attribute:compiled-pattern-regular-expression}.

2.  Let `pattern`{.variable} be the value of the
    [`pattern`{#the-pattern-attribute:attr-input-pattern-3}](#attr-input-pattern)
    attribute of the element.

3.  Let `regexpCompletion`{.variable} be
    [RegExpCreate](https://tc39.es/ecma262/#sec-regexpcreate){#the-pattern-attribute:regexpcreate
    x-internal="regexpcreate"}(`pattern`{.variable}, \"`v`\").

4.  If `regexpCompletion`{.variable} is an [abrupt
    completion](https://tc39.es/ecma262/#sec-completion-record-specification-type){#the-pattern-attribute:completion-record
    x-internal="completion-record"}, then return nothing. The element
    has no [compiled pattern regular
    expression](#compiled-pattern-regular-expression){#the-pattern-attribute:compiled-pattern-regular-expression-2}.

    User agents are encouraged to log this error in a developer console,
    to aid debugging.

5.  Let `anchoredPattern`{.variable} be the string \"`^(?:`\", followed
    by `pattern`{.variable}, followed by \"`)$`\".

6.  Return !
    [RegExpCreate](https://tc39.es/ecma262/#sec-regexpcreate){#the-pattern-attribute:regexpcreate-2
    x-internal="regexpcreate"}(`anchoredPattern`{.variable}, \"`v`\").

The reasoning behind these steps, instead of just using the value of the
[`pattern`{#the-pattern-attribute:attr-input-pattern-4}](#attr-input-pattern)
attribute directly, is twofold. First, we want to ensure that when
matched against a string, the regular expression\'s start is anchored to
the start of the string and its end to the end of the string. Second, we
want to ensure that the regular expression is valid in standalone form,
instead of only becoming valid after being surrounded by the \"`^(?:`\"
and \"`)$`\" anchors.

A
[`RegExp`{#the-pattern-attribute:regexp-2}](https://tc39.es/ecma262/#sec-regexp-regular-expression-objects){x-internal="regexp"}
object `regexp`{.variable} [matches]{#regexp-match-a-string .dfn} a
string `input`{.variable}, if !
[RegExpBuiltinExec](https://tc39.es/ecma262/#sec-regexpbuiltinexec){#the-pattern-attribute:regexpbuiltinexec
x-internal="regexpbuiltinexec"}(`regexp`{.variable}, `input`{.variable})
is not null.

**Constraint validation**: If the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-pattern-attribute:concept-fe-value-2}
is not the empty string, and either the element\'s
[`multiple`{#the-pattern-attribute:attr-input-multiple-2}](#attr-input-multiple)
attribute is not specified or it [does not
apply](#do-not-apply){#the-pattern-attribute:do-not-apply} to the
[`input`{#the-pattern-attribute:the-input-element-2}](#the-input-element)
element given its
[`type`{#the-pattern-attribute:attr-input-type}](#attr-input-type)
attribute\'s current state, and the element has a [compiled pattern
regular
expression](#compiled-pattern-regular-expression){#the-pattern-attribute:compiled-pattern-regular-expression-3}
but that regular expression does not
[match](#regexp-match-a-string){#the-pattern-attribute:regexp-match-a-string}
the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-pattern-attribute:concept-fe-value-3},
then the element is [suffering from a pattern
mismatch](form-control-infrastructure.html#suffering-from-a-pattern-mismatch){#the-pattern-attribute:suffering-from-a-pattern-mismatch}.

**Constraint validation**: If the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-pattern-attribute:concept-fe-value-4}
is not the empty string, and the element\'s
[`multiple`{#the-pattern-attribute:attr-input-multiple-3}](#attr-input-multiple)
attribute is specified and
[applies](#concept-input-apply){#the-pattern-attribute:concept-input-apply-2}
to the
[`input`{#the-pattern-attribute:the-input-element-3}](#the-input-element)
element, and the element has a [compiled pattern regular
expression](#compiled-pattern-regular-expression){#the-pattern-attribute:compiled-pattern-regular-expression-4}
but that regular expression does not
[match](#regexp-match-a-string){#the-pattern-attribute:regexp-match-a-string-2}
each of the element\'s
[values](form-control-infrastructure.html#concept-fe-values){#the-pattern-attribute:concept-fe-values-2},
then the element is [suffering from a pattern
mismatch](form-control-infrastructure.html#suffering-from-a-pattern-mismatch){#the-pattern-attribute:suffering-from-a-pattern-mismatch-2}.

When an
[`input`{#the-pattern-attribute:the-input-element-4}](#the-input-element)
element has a
[`pattern`{#the-pattern-attribute:attr-input-pattern-5}](#attr-input-pattern)
attribute specified, authors should include a
[`title`]{#attr-input-title .dfn dfn-for="input"
dfn-type="element-attr"} attribute to give a description of the pattern.
User agents may use the contents of this attribute, if it is present,
when informing the user that the pattern is not matched, or at any other
suitable time, such as in a tooltip or read out by assistive technology
when the control [gains
focus](interaction.html#gains-focus){#the-pattern-attribute:gains-focus}.

::: example
For example, the following snippet:

``` html
<label> Part number:
 <input pattern="[0-9][A-Z]{3}" name="part"
        title="A part number is a digit followed by three uppercase letters."/>
</label>
```

\...could cause the UA to display an alert such as:

    A part number is a digit followed by three uppercase letters.
    You cannot submit this form when the field is incorrect.
:::

When a control has a
[`pattern`{#the-pattern-attribute:attr-input-pattern-6}](#attr-input-pattern)
attribute, the
[`title`{#the-pattern-attribute:attr-input-title}](#attr-input-title)
attribute, if used, must describe the pattern. Additional information
could also be included, so long as it assists the user in filling in the
control. Otherwise, assistive technology would be impaired.

For instance, if the title attribute contained the caption of the
control, assistive technology could end up saying something like
`The text you have entered does not match the required pattern. Birthday`{.sample},
which is not useful.

UAs may still show the
[`title`{#the-pattern-attribute:attr-title}](dom.html#attr-title) in
non-error situations (for example, as a tooltip when hovering over the
control), so authors should be careful not to word
[`title`{#the-pattern-attribute:attr-input-title-2}](#attr-input-title)s
as if an error has necessarily occurred.

###### [4.10.5.3.7]{.secno} The [`min`{#the-min-and-max-attributes:attr-input-min}](#attr-input-min) and [`max`{#the-min-and-max-attributes:attr-input-max}](#attr-input-max) attributes[](#the-min-and-max-attributes){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Attributes/max](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/max "The max attribute defines the maximum value that is acceptable and valid for the input containing the attribute. If the value of the element is greater than this, the element fails validation. This value must be greater than or equal to the value of the min attribute. If the max attribute is present but is not specified or is invalid, no max value is applied. If the max attribute is valid and a non-empty value is greater than the maximum allowed by the max attribute, constraint validation will prevent form submission.")

Support in all current engines.

::: support
[Firefox16+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
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
[Attributes/min](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/min "The min attribute defines the minimum value that is acceptable and valid for the input containing the attribute. If the value of the element is less than this, the element fails validation. This value must be less than or equal to the value of the max attribute.")

Support in all current engines.

::: support
[Firefox16+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
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

Some form controls can have explicit constraints applied limiting the
allowed range of values that the user can provide. Normally, such a
range would be linear and continuous. A form control can [have a
periodic domain]{#has-a-periodic-domain .dfn}, however, in which case
the form control\'s broadest possible range is finite, and authors can
specify explicit ranges within it that span the boundaries.

Specifically, the broadest range of a
[`type=time`{#the-min-and-max-attributes:time-state-(type=time)}](#time-state-(type=time))
control is midnight to midnight (24 hours), and authors can set both
continuous linear ranges (such as 9pm to 11pm) and discontinuous ranges
spanning midnight (such as 11pm to 1am).

The [`min`]{#attr-input-min .dfn dfn-for="input"
dfn-type="element-attr"} and [`max`]{#attr-input-max .dfn
dfn-for="input" dfn-type="element-attr"} attributes indicate the allowed
range of values for the element.

Their syntax is defined by the section that defines the
[`type`{#the-min-and-max-attributes:attr-input-type}](#attr-input-type)
attribute\'s current state.

If the element has a
[`min`{#the-min-and-max-attributes:attr-input-min-2}](#attr-input-min)
attribute, and the result of applying the [algorithm to convert a string
to a
number](#concept-input-value-string-number){#the-min-and-max-attributes:concept-input-value-string-number}
to the value of the
[`min`{#the-min-and-max-attributes:attr-input-min-3}](#attr-input-min)
attribute is a number, then that number is the element\'s
[minimum]{#concept-input-min .dfn}; otherwise, if the
[`type`{#the-min-and-max-attributes:attr-input-type-2}](#attr-input-type)
attribute\'s current state defines a [default
minimum]{#concept-input-min-default .dfn}, then that is the
[minimum](#concept-input-min){#the-min-and-max-attributes:concept-input-min};
otherwise, the element has no
[minimum](#concept-input-min){#the-min-and-max-attributes:concept-input-min-2}.

The
[`min`{#the-min-and-max-attributes:attr-input-min-4}](#attr-input-min)
attribute also defines the [step
base](#concept-input-min-zero){#the-min-and-max-attributes:concept-input-min-zero}.

If the element has a
[`max`{#the-min-and-max-attributes:attr-input-max-2}](#attr-input-max)
attribute, and the result of applying the [algorithm to convert a string
to a
number](#concept-input-value-string-number){#the-min-and-max-attributes:concept-input-value-string-number-2}
to the value of the
[`max`{#the-min-and-max-attributes:attr-input-max-3}](#attr-input-max)
attribute is a number, then that number is the element\'s
[maximum]{#concept-input-max .dfn}; otherwise, if the
[`type`{#the-min-and-max-attributes:attr-input-type-3}](#attr-input-type)
attribute\'s current state defines a [default
maximum]{#concept-input-max-default .dfn}, then that is the
[maximum](#concept-input-max){#the-min-and-max-attributes:concept-input-max};
otherwise, the element has no
[maximum](#concept-input-max){#the-min-and-max-attributes:concept-input-max-2}.

If the element does not [have a periodic
domain](#has-a-periodic-domain){#the-min-and-max-attributes:has-a-periodic-domain},
the
[`max`{#the-min-and-max-attributes:attr-input-max-4}](#attr-input-max)
attribute\'s value (the
[maximum](#concept-input-max){#the-min-and-max-attributes:concept-input-max-3})
must not be less than the
[`min`{#the-min-and-max-attributes:attr-input-min-5}](#attr-input-min)
attribute\'s value (its
[minimum](#concept-input-min){#the-min-and-max-attributes:concept-input-min-3}).

If an element that does not [have a periodic
domain](#has-a-periodic-domain){#the-min-and-max-attributes:has-a-periodic-domain-2}
has a
[maximum](#attr-input-max){#the-min-and-max-attributes:attr-input-max-5}
that is less than its
[minimum](#attr-input-min){#the-min-and-max-attributes:attr-input-min-6},
then so long as the element has a
[value](form-control-infrastructure.html#concept-fe-value){#the-min-and-max-attributes:concept-fe-value},
it will either be [suffering from an
underflow](form-control-infrastructure.html#suffering-from-an-underflow){#the-min-and-max-attributes:suffering-from-an-underflow}
or [suffering from an
overflow](form-control-infrastructure.html#suffering-from-an-overflow){#the-min-and-max-attributes:suffering-from-an-overflow}.

An element [has a reversed range]{#has-a-reversed-range .dfn} if it [has
a periodic
domain](#has-a-periodic-domain){#the-min-and-max-attributes:has-a-periodic-domain-3}
and its
[maximum](#concept-input-max){#the-min-and-max-attributes:concept-input-max-4}
is less than its
[minimum](#concept-input-min){#the-min-and-max-attributes:concept-input-min-4}.

An element [has range limitations]{#have-range-limitations .dfn} if it
has a defined
[minimum](#concept-input-min){#the-min-and-max-attributes:concept-input-min-5}
or a defined
[maximum](#concept-input-max){#the-min-and-max-attributes:concept-input-max-5}.

**Constraint validation**: When the element has a
[minimum](#attr-input-min){#the-min-and-max-attributes:attr-input-min-7}
and does not [have a reversed
range](#has-a-reversed-range){#the-min-and-max-attributes:has-a-reversed-range},
and the result of applying the [algorithm to convert a string to a
number](#concept-input-value-string-number){#the-min-and-max-attributes:concept-input-value-string-number-3}
to the string given by the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-min-and-max-attributes:concept-fe-value-2}
is a number, and the number obtained from that algorithm is less than
the
[minimum](#attr-input-min){#the-min-and-max-attributes:attr-input-min-8},
the element is [suffering from an
underflow](form-control-infrastructure.html#suffering-from-an-underflow){#the-min-and-max-attributes:suffering-from-an-underflow-2}.

**Constraint validation**: When the element has a
[maximum](#attr-input-max){#the-min-and-max-attributes:attr-input-max-6}
and does not [have a reversed
range](#has-a-reversed-range){#the-min-and-max-attributes:has-a-reversed-range-2},
and the result of applying the [algorithm to convert a string to a
number](#concept-input-value-string-number){#the-min-and-max-attributes:concept-input-value-string-number-4}
to the string given by the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-min-and-max-attributes:concept-fe-value-3}
is a number, and the number obtained from that algorithm is more than
the
[maximum](#attr-input-max){#the-min-and-max-attributes:attr-input-max-7},
the element is [suffering from an
overflow](form-control-infrastructure.html#suffering-from-an-overflow){#the-min-and-max-attributes:suffering-from-an-overflow-2}.

**Constraint validation**: When an element [has a reversed
range](#has-a-reversed-range){#the-min-and-max-attributes:has-a-reversed-range-3},
and the result of applying the [algorithm to convert a string to a
number](#concept-input-value-string-number){#the-min-and-max-attributes:concept-input-value-string-number-5}
to the string given by the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-min-and-max-attributes:concept-fe-value-4}
is a number, and the number obtained from that algorithm is more than
the
[maximum](#attr-input-max){#the-min-and-max-attributes:attr-input-max-8}
*and* less than the
[minimum](#attr-input-min){#the-min-and-max-attributes:attr-input-min-9},
the element is simultaneously [suffering from an
underflow](form-control-infrastructure.html#suffering-from-an-underflow){#the-min-and-max-attributes:suffering-from-an-underflow-3}
and [suffering from an
overflow](form-control-infrastructure.html#suffering-from-an-overflow){#the-min-and-max-attributes:suffering-from-an-overflow-3}.

::: example
The following date control limits input to dates that are before the
1980s:

``` html
<input name=bday type=date max="1979-12-31">
```
:::

::: example
The following number control limits input to whole numbers greater than
zero:

``` html
<input name=quantity required="" type="number" min="1" value="1">
```
:::

::: example
The following time control limits input to those minutes that occur
between 9pm and 6am, defaulting to midnight:

``` html
<input name="sleepStart" type=time min="21:00" max="06:00" step="60" value="00:00">
```
:::

###### [4.10.5.3.8]{.secno} The [`step`{#the-step-attribute:attr-input-step}](#attr-input-step) attribute[](#the-step-attribute){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Attributes/step](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/step "The step attribute is a number that specifies the granularity that the value must adhere to or the keyword any. It is valid for the numeric input types, including the date, month, week, time, datetime-local, number and range types.")

Support in all current engines.

::: support
[Firefox16+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
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

The [`step`]{#attr-input-step .dfn dfn-for="input"
dfn-type="element-attr"} attribute indicates the granularity that is
expected (and required) of the
[value](form-control-infrastructure.html#concept-fe-value){#the-step-attribute:concept-fe-value}
or
[values](form-control-infrastructure.html#concept-fe-values){#the-step-attribute:concept-fe-values},
by limiting the allowed values. The section that defines the
[`type`{#the-step-attribute:attr-input-type}](#attr-input-type)
attribute\'s current state also defines the [default
step]{#concept-input-step-default .dfn}, the [step scale
factor]{#concept-input-step-scale .dfn}, and in some cases the [default
step base]{#concept-input-step-default-base .dfn}, which are used in
processing the attribute as described below.

The [`step`{#the-step-attribute:attr-input-step-2}](#attr-input-step)
attribute, if specified, must either have a value that is a [valid
floating-point
number](common-microsyntaxes.html#valid-floating-point-number){#the-step-attribute:valid-floating-point-number}
that
[parses](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#the-step-attribute:rules-for-parsing-floating-point-number-values}
to a number that is greater than zero, or must have a value that is an
[ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-step-attribute:ascii-case-insensitive
x-internal="ascii-case-insensitive"} match for the string \"`any`\".

The attribute provides the [allowed value step]{#concept-input-step
.dfn} for the element, as follows:

1.  If the attribute does not
    [apply](#concept-input-apply){#the-step-attribute:concept-input-apply},
    then there is no [allowed value
    step](#concept-input-step){#the-step-attribute:concept-input-step}.

2.  Otherwise, if the attribute is absent, then the [allowed value
    step](#concept-input-step){#the-step-attribute:concept-input-step-2}
    is the [default
    step](#concept-input-step-default){#the-step-attribute:concept-input-step-default}
    multiplied by the [step scale
    factor](#concept-input-step-scale){#the-step-attribute:concept-input-step-scale}.

3.  Otherwise, if the attribute\'s value is an [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-step-attribute:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for the string \"`any`\",
    then there is no [allowed value
    step](#concept-input-step){#the-step-attribute:concept-input-step-3}.

4.  Otherwise, if the [rules for parsing floating-point number
    values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#the-step-attribute:rules-for-parsing-floating-point-number-values-2},
    when they are applied to the attribute\'s value, return an error,
    zero, or a number less than zero, then the [allowed value
    step](#concept-input-step){#the-step-attribute:concept-input-step-4}
    is the [default
    step](#concept-input-step-default){#the-step-attribute:concept-input-step-default-2}
    multiplied by the [step scale
    factor](#concept-input-step-scale){#the-step-attribute:concept-input-step-scale-2}.

5.  Otherwise, the [allowed value
    step](#concept-input-step){#the-step-attribute:concept-input-step-5}
    is the number returned by the [rules for parsing floating-point
    number
    values](common-microsyntaxes.html#rules-for-parsing-floating-point-number-values){#the-step-attribute:rules-for-parsing-floating-point-number-values-3}
    when they are applied to the attribute\'s value, multiplied by the
    [step scale
    factor](#concept-input-step-scale){#the-step-attribute:concept-input-step-scale-3}.

The [step base]{#concept-input-min-zero .dfn} is the value returned by
the following algorithm:

1.  If the element has a
    [`min`{#the-step-attribute:attr-input-min}](#attr-input-min) content
    attribute, and the result of applying the [algorithm to convert a
    string to a
    number](#concept-input-value-string-number){#the-step-attribute:concept-input-value-string-number}
    to the value of the
    [`min`{#the-step-attribute:attr-input-min-2}](#attr-input-min)
    content attribute is not an error, then return that result.

2.  If the element has a
    [`value`{#the-step-attribute:attr-input-value}](#attr-input-value)
    content attribute, and the result of applying the [algorithm to
    convert a string to a
    number](#concept-input-value-string-number){#the-step-attribute:concept-input-value-string-number-2}
    to the value of the
    [`value`{#the-step-attribute:attr-input-value-2}](#attr-input-value)
    content attribute is not an error, then return that result.

3.  If a [default step
    base](#concept-input-step-default-base){#the-step-attribute:concept-input-step-default-base}
    is defined for this element given its
    [`type`{#the-step-attribute:attr-input-type-2}](#attr-input-type)
    attribute\'s state, then return it.

4.  Return zero.

**Constraint validation**: When the element has an [allowed value
step](#concept-input-step){#the-step-attribute:concept-input-step-6},
and the result of applying the [algorithm to convert a string to a
number](#concept-input-value-string-number){#the-step-attribute:concept-input-value-string-number-3}
to the string given by the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-step-attribute:concept-fe-value-2}
is a number, and that number subtracted from the [step
base](#concept-input-min-zero){#the-step-attribute:concept-input-min-zero}
is not an integral multiple of the [allowed value
step](#concept-input-step){#the-step-attribute:concept-input-step-7},
the element is [suffering from a step
mismatch](form-control-infrastructure.html#suffering-from-a-step-mismatch){#the-step-attribute:suffering-from-a-step-mismatch}.

::: example
The following range control only accepts values in the range 0..1, and
allows 256 steps in that range:

``` html
<input name=opacity type=range min=0 max=1 step=0.00392156863>
```
:::

::: example
The following control allows any time in the day to be selected, with
any accuracy (e.g. thousandth-of-a-second accuracy or more):

``` html
<input name=favtime type=time step=any>
```

Normally, time controls are limited to an accuracy of one minute.
:::

###### [4.10.5.3.9]{.secno} The [`list`{#the-list-attribute:attr-input-list}](#attr-input-list) attribute[](#the-list-attribute){.self-link}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#list](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#list "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

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

The [`list`]{#attr-input-list .dfn dfn-for="input"
dfn-type="element-attr"} attribute is used to identify an element that
lists predefined options suggested to the user.

If present, its value must be the
[ID](https://dom.spec.whatwg.org/#concept-id){#the-list-attribute:concept-id
x-internal="concept-id"} of a
[`datalist`{#the-list-attribute:the-datalist-element}](form-elements.html#the-datalist-element)
element in the same
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-list-attribute:tree
x-internal="tree"}.

The [suggestions source element]{#concept-input-list .dfn} is the first
element in the
[tree](https://dom.spec.whatwg.org/#concept-tree){#the-list-attribute:tree-2
x-internal="tree"} in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-list-attribute:tree-order
x-internal="tree-order"} to have an
[ID](https://dom.spec.whatwg.org/#concept-id){#the-list-attribute:concept-id-2
x-internal="concept-id"} equal to the value of the
[`list`{#the-list-attribute:attr-input-list-2}](#attr-input-list)
attribute, if that element is a
[`datalist`{#the-list-attribute:the-datalist-element-2}](form-elements.html#the-datalist-element)
element. If there is no
[`list`{#the-list-attribute:attr-input-list-3}](#attr-input-list)
attribute, or if there is no element with that
[ID](https://dom.spec.whatwg.org/#concept-id){#the-list-attribute:concept-id-3
x-internal="concept-id"}, or if the first element with that
[ID](https://dom.spec.whatwg.org/#concept-id){#the-list-attribute:concept-id-4
x-internal="concept-id"} is not a
[`datalist`{#the-list-attribute:the-datalist-element-3}](form-elements.html#the-datalist-element)
element, then there is no [suggestions source
element](#concept-input-list){#the-list-attribute:concept-input-list}.

If there is a [suggestions source
element](#concept-input-list){#the-list-attribute:concept-input-list-2},
then, when the user agent is allowing the user to edit the
[`input`{#the-list-attribute:the-input-element}](#the-input-element)
element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-list-attribute:concept-fe-value},
the user agent should offer the suggestions represented by the
[suggestions source
element](#concept-input-list){#the-list-attribute:concept-input-list-3}
to the user in a manner suitable for the type of control used. If
appropriate, the user agent should use the suggestion\'s
[label](form-elements.html#concept-option-label){#the-list-attribute:concept-option-label}
and
[value](form-elements.html#concept-option-value){#the-list-attribute:concept-option-value}
to identify the suggestion to the user.

User agents are encouraged to filter the suggestions represented by the
[suggestions source
element](#concept-input-list){#the-list-attribute:concept-input-list-4}
when the number of suggestions is large, including only the most
relevant ones (e.g. based on the user\'s input so far). No precise
threshold is defined, but capping the list at four to seven values is
reasonable. If filtering based on the user\'s input, user agents should
search within both the
[label](form-elements.html#concept-option-label){#the-list-attribute:concept-option-label-2}
and
[value](form-elements.html#concept-option-value){#the-list-attribute:concept-option-value-2}
of the suggestions for matches. User agents should consider how input
variations affect the matching process. Unicode normalization should be
applied so that different underlying Unicode code point sequences,
caused by different keyboard- or input-specific mechanisms, do not
interfere with the matching process. Case variations should be ignored,
which may require language-specific case mapping. For examples of these,
see Character Model for the World Wide Web: String Matching. User agents
may also provide other matching features: for illustration, a few
examples include matching different forms of kana to each other (or to
kanji), ignoring accents, or applying spelling correction.
[\[CHARMODNORM\]](references.html#refsCHARMODNORM)

::: example
This text field allows you to choose a type of JavaScript function.

``` html
<input type="text" list="function-types">
<datalist id="function-types">
  <option value="function">function</option>
  <option value="async function">async function</option>
  <option value="function*">generator function</option>
  <option value="=>">arrow function</option>
  <option value="async =>">async arrow function</option>
  <option value="async function*">async generator function</option>
</datalist>
```

For user agents that follow the above suggestions, both the
[label](form-elements.html#concept-option-label){#the-list-attribute:concept-option-label-3}
and
[value](form-elements.html#concept-option-value){#the-list-attribute:concept-option-value-3}
would be shown:

![A text box with a drop down button on the right hand side; with,
below, a drop down box containing a list of the six values the left and
the six labels on the
right.](/images/sample-datalist.svg){.darkmode-aware width="280"
height="150"}

Then, typing \"[arrow]{.kbd}\" or \"[=\>]{.kbd}\" would filter the list
to the entries with labels \"arrow function\" and \"async arrow
function\". Typing \"[generator]{.kbd}\" or \"[\*]{.kbd}\" would filter
the list to the entries with labels \"generator function\" and \"async
generator function\".
:::

As always, user agents are free to make user interface decisions which
are appropriate for their particular requirements and for the user\'s
particular circumstances. However, this has historically been an area of
confusion for implementers, web developers, and users alike, so we\'ve
given some \"should\" suggestions above.

How user selections of suggestions are handled depends on whether the
element is a control accepting a single value only, or whether it
accepts multiple values:

If the element does not have a [`multiple`{#the-list-attribute:attr-input-multiple}](#attr-input-multiple) attribute specified or if the [`multiple`{#the-list-attribute:attr-input-multiple-2}](#attr-input-multiple) attribute [does not apply](#do-not-apply){#the-list-attribute:do-not-apply}

:   When the user selects a suggestion, the
    [`input`{#the-list-attribute:the-input-element-2}](#the-input-element)
    element\'s
    [value](form-control-infrastructure.html#concept-fe-value){#the-list-attribute:concept-fe-value-2}
    must be set to the selected suggestion\'s
    [value](form-elements.html#concept-option-value){#the-list-attribute:concept-option-value-4},
    as if the user had written that value themself.

If the element\'s [`type`{#the-list-attribute:attr-input-type}](#attr-input-type) attribute is in the [Email](#email-state-(type=email)){#the-list-attribute:email-state-(type=email)} state and the element has a [`multiple`{#the-list-attribute:attr-input-multiple-3}](#attr-input-multiple) attribute specified

:   When the user selects a suggestion, the user agent must either add a
    new entry to the
    [`input`{#the-list-attribute:the-input-element-3}](#the-input-element)
    element\'s
    [values](form-control-infrastructure.html#concept-fe-values){#the-list-attribute:concept-fe-values},
    whose value is the selected suggestion\'s
    [value](form-elements.html#concept-option-value){#the-list-attribute:concept-option-value-5},
    or change an existing entry in the
    [`input`{#the-list-attribute:the-input-element-4}](#the-input-element)
    element\'s
    [values](form-control-infrastructure.html#concept-fe-values){#the-list-attribute:concept-fe-values-2}
    to have the value given by the selected suggestion\'s
    [value](form-elements.html#concept-option-value){#the-list-attribute:concept-option-value-6},
    as if the user had themself added an entry with that value, or
    edited an existing entry to be that value. Which behavior is to be
    applied depends on the user interface in an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-list-attribute:implementation-defined
    x-internal="implementation-defined"} manner.

------------------------------------------------------------------------

If the [`list`{#the-list-attribute:attr-input-list-4}](#attr-input-list)
attribute [does not
apply](#do-not-apply){#the-list-attribute:do-not-apply-2}, there is no
[suggestions source
element](#concept-input-list){#the-list-attribute:concept-input-list-5}.

::: example
This URL field offers some suggestions.

``` html
<label>Homepage: <input name=hp type=url list=hpurls></label>
<datalist id=hpurls>
 <option value="https://www.google.com/" label="Google">
 <option value="https://www.reddit.com/" label="Reddit">
</datalist>
```

Other URLs from the user\'s history might show also; this is up to the
user agent.
:::

::: example
This example demonstrates how to design a form that uses the
autocompletion list feature while still degrading usefully in legacy
user agents.

If the autocompletion list is merely an aid, and is not important to the
content, then simply using a
[`datalist`{#the-list-attribute:the-datalist-element-4}](form-elements.html#the-datalist-element)
element with children
[`option`{#the-list-attribute:the-option-element}](form-elements.html#the-option-element)
elements is enough. To prevent the values from being rendered in legacy
user agents, they need to be placed inside the
[`value`{#the-list-attribute:attr-option-value}](form-elements.html#attr-option-value)
attribute instead of inline.

``` html
<p>
 <label>
  Enter a breed:
  <input type="text" name="breed" list="breeds">
  <datalist id="breeds">
   <option value="Abyssinian">
   <option value="Alpaca">
   <!-- ... -->
  </datalist>
 </label>
</p>
```

However, if the values need to be shown in legacy UAs, then fallback
content can be placed inside the
[`datalist`{#the-list-attribute:the-datalist-element-5}](form-elements.html#the-datalist-element)
element, as follows:

``` html
<p>
 <label>
  Enter a breed:
  <input type="text" name="breed" list="breeds">
 </label>
 <datalist id="breeds">
  <label>
   or select one from the list:
   <select name="breed">
    <option value=""> (none selected)
    <option>Abyssinian
    <option>Alpaca
    <!-- ... -->
   </select>
  </label>
 </datalist>
</p>
```

The fallback content will only be shown in UAs that don\'t support
[`datalist`{#the-list-attribute:the-datalist-element-6}](form-elements.html#the-datalist-element).
The options, on the other hand, will be detected by all UAs, even though
they are not children of the
[`datalist`{#the-list-attribute:the-datalist-element-7}](form-elements.html#the-datalist-element)
element.

Note that if an
[`option`{#the-list-attribute:the-option-element-2}](form-elements.html#the-option-element)
element used in a
[`datalist`{#the-list-attribute:the-datalist-element-8}](form-elements.html#the-datalist-element)
is
[`selected`{#the-list-attribute:attr-option-selected}](form-elements.html#attr-option-selected),
it will be selected by default by legacy UAs (because it affects the
[`select`{#the-list-attribute:the-select-element}](form-elements.html#the-select-element)
element), but it will not have any effect on the
[`input`{#the-list-attribute:the-input-element-5}](#the-input-element)
element in UAs that support
[`datalist`{#the-list-attribute:the-datalist-element-9}](form-elements.html#the-datalist-element).
:::

###### [4.10.5.3.10]{.secno} The [`placeholder`{#the-placeholder-attribute:attr-input-placeholder}](#attr-input-placeholder) attribute[](#the-placeholder-attribute){.self-link}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/input#placeholder](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#placeholder "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
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
[Element/input#attr-placeholder](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#attr-placeholder "The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The <input> element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome3+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera≤12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android≤37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android≤12.1+]{.opera_android .yes}
:::
::::
:::::::

The [`placeholder`]{#attr-input-placeholder .dfn dfn-for="input"
dfn-type="element-attr"} attribute represents a *short* hint (a word or
short phrase) intended to aid the user with data entry when the control
has no value. A hint could be a sample value or a brief description of
the expected format. The attribute, if specified, must have a value that
contains no U+000A LINE FEED (LF) or U+000D CARRIAGE RETURN (CR)
characters.

The
[`placeholder`{#the-placeholder-attribute:attr-input-placeholder-2}](#attr-input-placeholder)
attribute should not be used as an alternative to a
[`label`{#the-placeholder-attribute:the-label-element}](forms.html#the-label-element).
For a longer hint or other advisory text, the
[`title`{#the-placeholder-attribute:attr-title}](dom.html#attr-title)
attribute is more appropriate.

These mechanisms are very similar but subtly different: the hint given
by the control\'s
[`label`{#the-placeholder-attribute:the-label-element-2}](forms.html#the-label-element)
is shown at all times; the short hint given in the
[`placeholder`{#the-placeholder-attribute:attr-input-placeholder-3}](#attr-input-placeholder)
attribute is shown before the user enters a value; and the hint in the
[`title`{#the-placeholder-attribute:attr-title-2}](dom.html#attr-title)
attribute is shown when the user requests further help.

User agents should present this hint to the user, after having [stripped
newlines](https://infra.spec.whatwg.org/#strip-newlines){#the-placeholder-attribute:strip-newlines
x-internal="strip-newlines"} from it, when the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#the-placeholder-attribute:concept-fe-value}
is the empty string, especially if the control is not
[focused](interaction.html#focused){#the-placeholder-attribute:focused}.

If a user agent normally doesn\'t show this hint to the user when the
control is
[focused](interaction.html#focused){#the-placeholder-attribute:focused-2},
then the user agent should nonetheless show the hint for the control if
it was focused as a result of the
[`autofocus`{#the-placeholder-attribute:attr-fe-autofocus}](interaction.html#attr-fe-autofocus)
attribute, since in that case the user will not have had an opportunity
to examine the control before focusing it.

::: example
Here is an example of a mail configuration user interface that uses the
[`placeholder`{#the-placeholder-attribute:attr-input-placeholder-4}](#attr-input-placeholder)
attribute:

``` html
<fieldset>
 <legend>Mail Account</legend>
 <p><label>Name: <input type="text" name="fullname" placeholder="John Ratzenberger"></label></p>
 <p><label>Address: <input type="email" name="address" placeholder="john@example.net"></label></p>
 <p><label>Password: <input type="password" name="password"></label></p>
 <p><label>Description: <input type="text" name="desc" placeholder="My Email Account"></label></p>
</fieldset>
```
:::

::: example
In situations where the control\'s content has one directionality but
the placeholder needs to have a different directionality, Unicode\'s
bidirectional-algorithm formatting characters can be used in the
attribute value:

``` html
<input name=t1 type=tel placeholder="&#x202B; رقم الهاتف 1 &#x202E;">
<input name=t2 type=tel placeholder="&#x202B; رقم الهاتف 2 &#x202E;">
```

For slightly more clarity, here\'s the same example using numeric
character references instead of inline Arabic:

``` html
<input name=t1 type=tel placeholder="&#x202B;&#1585;&#1602;&#1605; &#1575;&#1604;&#1607;&#1575;&#1578;&#1601; 1&#x202E;">
<input name=t2 type=tel placeholder="&#x202B;&#1585;&#1602;&#1605; &#1575;&#1604;&#1607;&#1575;&#1578;&#1601; 2&#x202E;">
```
:::

##### [4.10.5.4]{.secno} Common [`input`{#common-input-element-apis:the-input-element}](#the-input-element) element APIs[](#common-input-element-apis){.self-link}

`input`{.variable}`.`[`value`](#dom-input-value){#dom-input-value-dev}` [ = ``value`{.variable}` ]`

Returns the current
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value}
of the form control.

Can be set, to change the value.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror
x-internal="invalidstateerror"}
[`DOMException`{#common-input-element-apis:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if it is set to any value other than the empty string when the control
is a file upload control.

`input`{.variable}`.`[`checked`](#dom-input-checked){#dom-input-checked-dev}` [ = ``value`{.variable}` ]`

Returns the current
[checkedness](form-control-infrastructure.html#concept-fe-checked){#common-input-element-apis:concept-fe-checked}
of the form control.

Can be set, to change the
[checkedness](form-control-infrastructure.html#concept-fe-checked){#common-input-element-apis:concept-fe-checked-2}.

`input`{.variable}`.`[`files`](#dom-input-files){#dom-input-files-dev}` [ = ``files`{.variable}` ]`

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[FileList](https://developer.mozilla.org/en-US/docs/Web/API/FileList "An object of this type is returned by the files property of the HTML <input> element; this lets you access the list of files selected with the <input type="file"> element. It's also used for a list of files dropped into web content when using the drag and drop API; see the DataTransfer object for details on this usage.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.1+]{.opera_android .yes}
:::
::::

:::: feature
[HTMLInputElement/files](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/files "The HTMLInputElement.files property allows you to access the FileList selected with the <input type="file"> element.")

Support in all current engines.

::: support
[Firefox3+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer[🔰
10+]{title="Partial implementation."}]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::::

Returns a
[`FileList`{#common-input-element-apis:filelist}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"}
object listing the [selected
files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected}
of the form control.

Returns null if the control isn\'t a file control.

Can be set to a
[`FileList`{#common-input-element-apis:filelist-2}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"}
object to change the [selected
files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-2}
of the form control. For instance, as the result of a drag-and-drop
operation.

`input`{.variable}`.`[`valueAsDate`](#dom-input-valueasdate){#dom-input-valueasdate-dev}` [ = ``value`{.variable}` ]`

Returns a
[`Date`{#common-input-element-apis:date}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object representing the form control\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-2},
if applicable; otherwise, returns null.

Can be set, to change the value.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-2
x-internal="invalidstateerror"}
[`DOMException`{#common-input-element-apis:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the control isn\'t date- or time-based.

`input`{.variable}`.`[`valueAsNumber`](#dom-input-valueasnumber){#dom-input-valueasnumber-dev}` [ = ``value`{.variable}` ]`

Returns a number representing the form control\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-3},
if applicable; otherwise, returns NaN.

Can be set, to change the value. Setting this to NaN will set the
underlying value to the empty string.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-3
x-internal="invalidstateerror"}
[`DOMException`{#common-input-element-apis:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the control is neither date- or time-based nor numeric.

`input`{.variable}`.`[`stepUp`](#dom-input-stepup){#dom-input-stepup-dev}`([ ``n`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLInputElement/stepUp](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/stepUp "The HTMLInputElement.stepUp() method increments the value of a numeric type of <input> element by the value of the step attribute, or the default step value if the step attribute is not explicitly set. The method, when invoked, increments the value by (step * n), where n defaults to 1 if not specified, and step defaults to the default value for step if not specified.")

::: support
[Firefox[🔰 16+]{title="Partial implementation."}]{.firefox
.yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome .yes}

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

`input`{.variable}`.`[`stepDown`](#dom-input-stepdown){#dom-input-stepdown-dev}`([ ``n`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[HTMLInputElement/stepDown](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/stepDown "The HTMLInputElement.stepDown([n]) method decrements the value of a numeric type of <input> element by the value of the step attribute or up to n multiples of the step attribute if a number is passed as the parameter.")

::: support
[Firefox[🔰 16+]{title="Partial implementation."}]{.firefox
.yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome .yes}

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

Changes the form control\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-4}
by the value given in the
[`step`{#common-input-element-apis:attr-input-step}](#attr-input-step)
attribute, multiplied by `n`{.variable}. The default value for
`n`{.variable} is 1.

Throws
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-4
x-internal="invalidstateerror"}
[`DOMException`{#common-input-element-apis:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if the control is neither date- or time-based nor numeric, or if the
[`step`{#common-input-element-apis:attr-input-step-2}](#attr-input-step)
attribute\'s value is \"`any`\".

`input`{.variable}`.`[`list`](#dom-input-list){#dom-input-list-dev}

Returns the
[`datalist`{#common-input-element-apis:the-datalist-element}](form-elements.html#the-datalist-element)
element indicated by the
[`list`{#common-input-element-apis:attr-input-list}](#attr-input-list)
attribute.

`input`{.variable}`.`[`showPicker`](#dom-input-showpicker){#dom-input-showpicker-dev}`()`

Shows any applicable picker UI for `input`{.variable}, so that the user
can select a value.

If `input`{.variable} does not [support a
picker](#input-support-picker){#common-input-element-apis:input-support-picker},
this method does nothing.

Throws an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-5
x-internal="invalidstateerror"}
[`DOMException`{#common-input-element-apis:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `input`{.variable} is not
[mutable](form-control-infrastructure.html#concept-fe-mutable){#common-input-element-apis:concept-fe-mutable}.

Throws a
[\"`NotAllowedError`\"](https://webidl.spec.whatwg.org/#notallowederror){#common-input-element-apis:notallowederror
x-internal="notallowederror"}
[`DOMException`{#common-input-element-apis:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if called without [transient user
activation](interaction.html#transient-activation){#common-input-element-apis:transient-activation}.

Throws a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#common-input-element-apis:securityerror
x-internal="securityerror"}
[`DOMException`{#common-input-element-apis:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
if `input`{.variable} is inside a cross-origin
[`iframe`{#common-input-element-apis:the-iframe-element}](iframe-embed-object.html#the-iframe-element),
unless `input`{.variable} is in the [File
Upload](#file-upload-state-(type=file)){#common-input-element-apis:file-upload-state-(type=file)}
or
[Color](#color-state-(type=color)){#common-input-element-apis:color-state-(type=color)}
states.

The [`value`]{#dom-input-value .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute allows scripts to manipulate the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-5}
of an
[`input`{#common-input-element-apis:the-input-element-2}](#the-input-element)
element. The attribute is in one of the following modes, which define
its behavior:

[value]{#dom-input-value-value .dfn}

:   On getting, return the current
    [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-6}
    of the element.

    On setting:

    1.  Let `oldValue`{.variable} be the element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-7}.

    2.  Set the element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-8}
        to the new value.

    3.  Set the element\'s [dirty value
        flag](form-control-infrastructure.html#concept-fe-dirty){#common-input-element-apis:concept-fe-dirty}
        to true.

    4.  Invoke the [value sanitization
        algorithm](#value-sanitization-algorithm){#common-input-element-apis:value-sanitization-algorithm},
        if the element\'s
        [`type`{#common-input-element-apis:attr-input-type}](#attr-input-type)
        attribute\'s current state defines one.

    5.  If the element\'s
        [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-9}
        (after applying the [value sanitization
        algorithm](#value-sanitization-algorithm){#common-input-element-apis:value-sanitization-algorithm-2})
        is different from `oldValue`{.variable}, and the element has a
        [text entry cursor
        position](form-control-infrastructure.html#concept-textarea/input-cursor){#common-input-element-apis:concept-textarea/input-cursor},
        move the [text entry cursor
        position](form-control-infrastructure.html#concept-textarea/input-cursor){#common-input-element-apis:concept-textarea/input-cursor-2}
        to the end of the text control, unselecting any selected text
        and [resetting the selection
        direction](form-control-infrastructure.html#set-the-selection-direction){#common-input-element-apis:set-the-selection-direction}
        to \"`none`\".

[default]{#dom-input-value-default .dfn}

:   On getting, if the element has a
    [`value`{#common-input-element-apis:attr-input-value}](#attr-input-value)
    content attribute, return that attribute\'s value; otherwise, return
    the empty string.

    On setting, set the value of the element\'s
    [`value`{#common-input-element-apis:attr-input-value-2}](#attr-input-value)
    content attribute to the new value.

[default/on]{#dom-input-value-default-on .dfn}

:   On getting, if the element has a
    [`value`{#common-input-element-apis:attr-input-value-3}](#attr-input-value)
    content attribute, return that attribute\'s value; otherwise, return
    the string \"`on`\".

    On setting, set the value of the element\'s
    [`value`{#common-input-element-apis:attr-input-value-4}](#attr-input-value)
    content attribute to the new value.

[filename]{#dom-input-value-filename .dfn}

:   On getting, return the string \"`C:\fakepath\`\" followed by the
    name of the first file in the list of [selected
    files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-3},
    if any, or the empty string if the list is empty.

    On setting, if the new value is the empty string, empty the list of
    [selected
    files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-4};
    otherwise, throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-6
    x-internal="invalidstateerror"}
    [`DOMException`{#common-input-element-apis:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    This \"fakepath\" requirement is a sad accident of history. See [the
    example in the File Upload state section](#fakepath-srsly) for more
    information.

    Since [path
    components](#concept-input-file-path){#common-input-element-apis:concept-input-file-path}
    are not permitted in filenames in the list of [selected
    files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-5},
    the \"\\fakepath\\\" cannot be mistaken for a path component.

------------------------------------------------------------------------

The [`checked`]{#dom-input-checked .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute allows scripts to manipulate the
[checkedness](form-control-infrastructure.html#concept-fe-checked){#common-input-element-apis:concept-fe-checked-3}
of an
[`input`{#common-input-element-apis:the-input-element-3}](#the-input-element)
element. On getting, it must return the current
[checkedness](form-control-infrastructure.html#concept-fe-checked){#common-input-element-apis:concept-fe-checked-4}
of the element; and on setting, it must set the element\'s
[checkedness](form-control-infrastructure.html#concept-fe-checked){#common-input-element-apis:concept-fe-checked-5}
to the new value and set the element\'s [dirty checkedness
flag](#concept-input-checked-dirty-flag){#common-input-element-apis:concept-input-checked-dirty-flag}
to true.

------------------------------------------------------------------------

The [`files`]{#dom-input-files .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute allows scripts to access the
element\'s [selected
files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-6}.

On getting, if the IDL attribute
[applies](#concept-input-apply){#common-input-element-apis:concept-input-apply},
it must return a
[`FileList`{#common-input-element-apis:filelist-3}](https://w3c.github.io/FileAPI/#filelist-section){x-internal="filelist"}
object that represents the current [selected
files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-7}.
The same object must be returned until the list of [selected
files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-8}
changes. If the IDL attribute [does not
apply](#do-not-apply){#common-input-element-apis:do-not-apply}, then it
must instead return null. [\[FILEAPI\]](references.html#refsFILEAPI)

On setting, it must run these steps:

1.  If the IDL attribute [does not
    apply](#do-not-apply){#common-input-element-apis:do-not-apply-2} or
    the given value is null, then return.

2.  Replace the element\'s [selected
    files](#concept-input-type-file-selected){#common-input-element-apis:concept-input-type-file-selected-9}
    with the given value.

------------------------------------------------------------------------

The [`valueAsDate`]{#dom-input-valueasdate .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"} IDL attribute
represents the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-10}
of the element, interpreted as a date.

On getting, if the
[`valueAsDate`{#common-input-element-apis:dom-input-valueasdate}](#dom-input-valueasdate)
attribute [does not
apply](#do-not-apply){#common-input-element-apis:do-not-apply-3}, as
defined for the
[`input`{#common-input-element-apis:the-input-element-4}](#the-input-element)
element\'s
[`type`{#common-input-element-apis:attr-input-type-2}](#attr-input-type)
attribute\'s current state, then return null. Otherwise, run the
[algorithm to convert a string to a `Date`
object](#concept-input-value-string-date){#common-input-element-apis:concept-input-value-string-date}
defined for that state to the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-11};
if the algorithm returned a
[`Date`{#common-input-element-apis:date-2}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object, then return it, otherwise, return null.

On setting, if the
[`valueAsDate`{#common-input-element-apis:dom-input-valueasdate-2}](#dom-input-valueasdate)
attribute [does not
apply](#do-not-apply){#common-input-element-apis:do-not-apply-4}, as
defined for the
[`input`{#common-input-element-apis:the-input-element-5}](#the-input-element)
element\'s
[`type`{#common-input-element-apis:attr-input-type-3}](#attr-input-type)
attribute\'s current state, then throw an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-7
x-internal="invalidstateerror"}
[`DOMException`{#common-input-element-apis:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"};
otherwise, if the new value is not null and not a
[`Date`{#common-input-element-apis:date-3}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object throw a
[`TypeError`{#common-input-element-apis:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
exception; otherwise, if the new value is null or a
[`Date`{#common-input-element-apis:date-4}](https://tc39.es/ecma262/#sec-date-objects){x-internal="date"}
object representing the NaN time value, then set the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-12}
of the element to the empty string; otherwise, run the [algorithm to
convert a `Date` object to a
string](#concept-input-value-date-string){#common-input-element-apis:concept-input-value-date-string},
as defined for that state, on the new value, and set the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-13}
of the element to the resulting string.

------------------------------------------------------------------------

The [`valueAsNumber`]{#dom-input-valueasnumber .dfn
dfn-for="HTMLInputElement" dfn-type="attribute"} IDL attribute
represents the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-14}
of the element, interpreted as a number.

On getting, if the
[`valueAsNumber`{#common-input-element-apis:dom-input-valueasnumber}](#dom-input-valueasnumber)
attribute [does not
apply](#do-not-apply){#common-input-element-apis:do-not-apply-5}, as
defined for the
[`input`{#common-input-element-apis:the-input-element-6}](#the-input-element)
element\'s
[`type`{#common-input-element-apis:attr-input-type-4}](#attr-input-type)
attribute\'s current state, then return a Not-a-Number (NaN) value.
Otherwise, run the [algorithm to convert a string to a
number](#concept-input-value-string-number){#common-input-element-apis:concept-input-value-string-number}
defined for that state to the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-15};
if the algorithm returned a number, then return it, otherwise, return a
Not-a-Number (NaN) value.

On setting, if the new value is infinite, then throw a
[`TypeError`{#common-input-element-apis:typeerror-2}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
exception. Otherwise, if the
[`valueAsNumber`{#common-input-element-apis:dom-input-valueasnumber-2}](#dom-input-valueasnumber)
attribute [does not
apply](#do-not-apply){#common-input-element-apis:do-not-apply-6}, as
defined for the
[`input`{#common-input-element-apis:the-input-element-7}](#the-input-element)
element\'s
[`type`{#common-input-element-apis:attr-input-type-5}](#attr-input-type)
attribute\'s current state, then throw an
[\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-8
x-internal="invalidstateerror"}
[`DOMException`{#common-input-element-apis:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
Otherwise, if the new value is a Not-a-Number (NaN) value, then set the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-16}
of the element to the empty string. Otherwise, run the [algorithm to
convert a number to a
string](#concept-input-value-number-string){#common-input-element-apis:concept-input-value-number-string},
as defined for that state, on the new value, and set the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-17}
of the element to the resulting string.

------------------------------------------------------------------------

The [`stepDown(``n`{.variable}`)`]{#dom-input-stepdown .dfn
dfn-for="HTMLInputElement" dfn-type="method"} and
[`stepUp(``n`{.variable}`)`]{#dom-input-stepup .dfn
dfn-for="HTMLInputElement" dfn-type="method"} methods, when invoked,
must run the following algorithm:

1.  If the
    [`stepDown()`{#common-input-element-apis:dom-input-stepdown}](#dom-input-stepdown)
    and
    [`stepUp()`{#common-input-element-apis:dom-input-stepup}](#dom-input-stepup)
    methods [do not
    apply](#do-not-apply){#common-input-element-apis:do-not-apply-7}, as
    defined for the
    [`input`{#common-input-element-apis:the-input-element-8}](#the-input-element)
    element\'s
    [`type`{#common-input-element-apis:attr-input-type-6}](#attr-input-type)
    attribute\'s current state, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-9
    x-internal="invalidstateerror"}
    [`DOMException`{#common-input-element-apis:domexception-11}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If the element has no [allowed value
    step](#concept-input-step){#common-input-element-apis:concept-input-step},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-10
    x-internal="invalidstateerror"}
    [`DOMException`{#common-input-element-apis:domexception-12}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If the element has a
    [minimum](#concept-input-min){#common-input-element-apis:concept-input-min}
    and a
    [maximum](#concept-input-max){#common-input-element-apis:concept-input-max}
    and the
    [minimum](#concept-input-min){#common-input-element-apis:concept-input-min-2}
    is greater than the
    [maximum](#concept-input-max){#common-input-element-apis:concept-input-max-2},
    then return.

4.  If the element has a
    [minimum](#concept-input-min){#common-input-element-apis:concept-input-min-3}
    and a
    [maximum](#concept-input-max){#common-input-element-apis:concept-input-max-3}
    and there is no value greater than or equal to the element\'s
    [minimum](#concept-input-min){#common-input-element-apis:concept-input-min-4}
    and less than or equal to the element\'s
    [maximum](#concept-input-max){#common-input-element-apis:concept-input-max-4}
    that, when subtracted from the [step
    base](#concept-input-min-zero){#common-input-element-apis:concept-input-min-zero},
    is an integral multiple of the [allowed value
    step](#concept-input-step){#common-input-element-apis:concept-input-step-2},
    then return.

5.  If applying the [algorithm to convert a string to a
    number](#concept-input-value-string-number){#common-input-element-apis:concept-input-value-string-number-2}
    to the string given by the element\'s
    [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-18}
    does not result in an error, then let `value`{.variable} be the
    result of that algorithm. Otherwise, let `value`{.variable} be zero.

6.  Let `valueBeforeStepping`{.variable} be `value`{.variable}.

7.  If `value`{.variable} subtracted from the [step
    base](#concept-input-min-zero){#common-input-element-apis:concept-input-min-zero-2}
    is not an integral multiple of the [allowed value
    step](#concept-input-step){#common-input-element-apis:concept-input-step-3},
    then set `value`{.variable} to the nearest value that, when
    subtracted from the [step
    base](#concept-input-min-zero){#common-input-element-apis:concept-input-min-zero-3},
    is an integral multiple of the [allowed value
    step](#concept-input-step){#common-input-element-apis:concept-input-step-4},
    and that is less than `value`{.variable} if the method invoked was
    the
    [`stepDown()`{#common-input-element-apis:dom-input-stepdown-2}](#dom-input-stepdown)
    method, and more than `value`{.variable} otherwise.

    Otherwise (`value`{.variable} subtracted from the [step
    base](#concept-input-min-zero){#common-input-element-apis:concept-input-min-zero-4}
    is an integral multiple of the [allowed value
    step](#concept-input-step){#common-input-element-apis:concept-input-step-5}):

    1.  Let `n`{.variable} be the argument.

    2.  Let `delta`{.variable} be the [allowed value
        step](#concept-input-step){#common-input-element-apis:concept-input-step-6}
        multiplied by `n`{.variable}.

    3.  If the method invoked was the
        [`stepDown()`{#common-input-element-apis:dom-input-stepdown-3}](#dom-input-stepdown)
        method, negate `delta`{.variable}.

    4.  Let `value`{.variable} be the result of adding
        `delta`{.variable} to `value`{.variable}.

8.  If the element has a
    [minimum](#concept-input-min){#common-input-element-apis:concept-input-min-5},
    and `value`{.variable} is less than that
    [minimum](#concept-input-min){#common-input-element-apis:concept-input-min-6},
    then set `value`{.variable} to the smallest value that, when
    subtracted from the [step
    base](#concept-input-min-zero){#common-input-element-apis:concept-input-min-zero-5},
    is an integral multiple of the [allowed value
    step](#concept-input-step){#common-input-element-apis:concept-input-step-7},
    and that is more than or equal to that
    [minimum](#concept-input-min){#common-input-element-apis:concept-input-min-7}.

9.  If the element has a
    [maximum](#concept-input-max){#common-input-element-apis:concept-input-max-5},
    and `value`{.variable} is greater than that
    [maximum](#concept-input-max){#common-input-element-apis:concept-input-max-6},
    then set `value`{.variable} to the largest value that, when
    subtracted from the [step
    base](#concept-input-min-zero){#common-input-element-apis:concept-input-min-zero-6},
    is an integral multiple of the [allowed value
    step](#concept-input-step){#common-input-element-apis:concept-input-step-8},
    and that is less than or equal to that
    [maximum](#concept-input-max){#common-input-element-apis:concept-input-max-7}.

10. If either the method invoked was the
    [`stepDown()`{#common-input-element-apis:dom-input-stepdown-4}](#dom-input-stepdown)
    method and `value`{.variable} is greater than
    `valueBeforeStepping`{.variable}, or the method invoked was the
    [`stepUp()`{#common-input-element-apis:dom-input-stepup-2}](#dom-input-stepup)
    method and `value`{.variable} is less than
    `valueBeforeStepping`{.variable}, then return.

    ::: example
    This ensures that invoking the
    [`stepUp()`{#common-input-element-apis:dom-input-stepup-3}](#dom-input-stepup)
    method on the
    [`input`{#common-input-element-apis:the-input-element-9}](#the-input-element)
    element in the following example does not change the
    [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-19}
    of that element:

    ``` html
    <input type=number value=1 max=0>
    ```
    :::

11. Let `value as string`{.variable} be the result of running the
    [algorithm to convert a number to a
    string](#concept-input-value-number-string){#common-input-element-apis:concept-input-value-number-string-2},
    as defined for the
    [`input`{#common-input-element-apis:the-input-element-10}](#the-input-element)
    element\'s
    [`type`{#common-input-element-apis:attr-input-type-7}](#attr-input-type)
    attribute\'s current state, on `value`{.variable}.

12. Set the
    [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-20}
    of the element to `value as string`{.variable}.

------------------------------------------------------------------------

The [`list`]{#dom-input-list .dfn dfn-for="HTMLInputElement"
dfn-type="attribute"} IDL attribute must return the current [suggestions
source
element](#concept-input-list){#common-input-element-apis:concept-input-list},
if any, or null otherwise.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLInputElement/showPicker](https://developer.mozilla.org/en-US/docs/Web/API/HTMLInputElement/showPicker "The HTMLInputElement.showPicker() method displays the browser picker for an input element.")

Support in all current engines.

::: support
[Firefox101+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome99+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge99+]{.edge_blink .yes}

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
[`HTMLInputElement`{#common-input-element-apis:htmlinputelement}](#htmlinputelement)
[`showPicker()`]{#dom-input-showpicker .dfn dfn-for="HTMLInputElement"
dfn-type="method"} and
[`HTMLSelectElement`{#common-input-element-apis:htmlselectelement}](form-elements.html#htmlselectelement)
[`showPicker()`]{#dom-select-showpicker .dfn dfn-for="HTMLSelectElement"
dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this
    x-internal="this"} is not
    [mutable](form-control-infrastructure.html#concept-fe-mutable){#common-input-element-apis:concept-fe-mutable-2},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#common-input-element-apis:invalidstateerror-11
    x-internal="invalidstateerror"}
    [`DOMException`{#common-input-element-apis:domexception-13}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-2
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#common-input-element-apis:relevant-settings-object}\'s
    [origin](webappapis.html#concept-settings-object-origin){#common-input-element-apis:concept-settings-object-origin}
    is not [same
    origin](browsers.html#same-origin){#common-input-element-apis:same-origin}
    with
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-3
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#common-input-element-apis:relevant-settings-object-2}\'s
    [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#common-input-element-apis:concept-environment-top-level-origin},
    and
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-4
    x-internal="this"} is a
    [`select`{#common-input-element-apis:the-select-element}](form-elements.html#the-select-element)
    element, or
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-5
    x-internal="this"}\'s
    [`type`{#common-input-element-apis:attr-input-type-8}](#attr-input-type)
    attribute is not in the [File
    Upload](#file-upload-state-(type=file)){#common-input-element-apis:file-upload-state-(type=file)-2}
    state or
    [Color](#color-state-(type=color)){#common-input-element-apis:color-state-(type=color)-2}
    state, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#common-input-element-apis:securityerror-2
    x-internal="securityerror"}
    [`DOMException`{#common-input-element-apis:domexception-14}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    [File](#file-upload-state-(type=file)){#common-input-element-apis:file-upload-state-(type=file)-3}
    and
    [Color](#color-state-(type=color)){#common-input-element-apis:color-state-(type=color)-3}
    inputs are exempted from this check for historical reason: their
    [input activation
    behavior](#input-activation-behavior){#common-input-element-apis:input-activation-behavior}
    also shows their pickers, and has never been guarded by an origin
    check.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-6
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#common-input-element-apis:concept-relevant-global}
    does not have [transient
    activation](interaction.html#transient-activation){#common-input-element-apis:transient-activation-2},
    then throw a
    [\"`NotAllowedError`\"](https://webidl.spec.whatwg.org/#notallowederror){#common-input-element-apis:notallowederror-2
    x-internal="notallowederror"}
    [`DOMException`{#common-input-element-apis:domexception-15}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  If
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-7
    x-internal="this"} is a
    [`select`{#common-input-element-apis:the-select-element-2}](form-elements.html#the-select-element)
    element, and
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-8
    x-internal="this"} is not [being
    rendered](rendering.html#being-rendered){#common-input-element-apis:being-rendered},
    then throw a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#common-input-element-apis:notsupportederror
    x-internal="notsupportederror"}
    [`DOMException`{#common-input-element-apis:domexception-16}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  [Show the picker, if
    applicable](#show-the-picker,-if-applicable){#common-input-element-apis:show-the-picker,-if-applicable},
    for
    [this](https://webidl.spec.whatwg.org/#this){#common-input-element-apis:this-9
    x-internal="this"}.

To [show the picker, if applicable]{#show-the-picker,-if-applicable
.dfn} for an
[`input`{#common-input-element-apis:the-input-element-11}](#the-input-element)
or
[`select`{#common-input-element-apis:the-select-element-3}](form-elements.html#the-select-element)
element `element`{.variable}:

1.  If `element`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#common-input-element-apis:concept-relevant-global-2}
    does not have [transient
    activation](interaction.html#transient-activation){#common-input-element-apis:transient-activation-3},
    then return.

2.  If `element`{.variable} is not
    *[mutable](form-control-infrastructure.html#concept-fe-mutable)*,
    then return.

3.  [Consume user
    activation](interaction.html#consume-user-activation){#common-input-element-apis:consume-user-activation}
    given `element`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#common-input-element-apis:concept-relevant-global-3}.

4.  If `element`{.variable} does not [support a
    picker](#input-support-picker){#common-input-element-apis:input-support-picker-2},
    then return.

5.  If `element`{.variable} is an
    [`input`{#common-input-element-apis:the-input-element-12}](#the-input-element)
    element and `element`{.variable}\'s
    [`type`{#common-input-element-apis:attr-input-type-9}](#attr-input-type)
    attribute is in the [File
    Upload](#file-upload-state-(type=file)){#common-input-element-apis:file-upload-state-(type=file)-4}
    state, then run these steps [in
    parallel](infrastructure.html#in-parallel){#common-input-element-apis:in-parallel}:

    1.  Optionally, wait until any prior execution of this algorithm has
        terminated.

    2.  Let `dismissed`{.variable} be the result of [WebDriver BiDi file
        dialog
        opened](https://w3c.github.io/webdriver-bidi/#webdriver-bidi-file-dialog-opened){#common-input-element-apis:webdriver-bidi-file-dialog-opened
        x-internal="webdriver-bidi-file-dialog-opened"} with
        `element`{.variable}.

    3.  If `dismissed`{.variable} is false:

        1.  Display a prompt to the user requesting that the user
            specify some files. If the
            [`multiple`{#common-input-element-apis:attr-input-multiple}](#attr-input-multiple)
            attribute is not set on `element`{.variable}, there must be
            no more than one file selected; otherwise, any number may be
            selected. Files can be from the filesystem or created on the
            fly, e.g., a picture taken from a camera connected to the
            user\'s device.

        2.  Wait for the user to have made their selection.

    4.  If `dismissed`{.variable} is true or if the user dismissed the
        prompt without changing their selection, then [queue an element
        task](webappapis.html#queue-an-element-task){#common-input-element-apis:queue-an-element-task}
        on the [user interaction task
        source](webappapis.html#user-interaction-task-source){#common-input-element-apis:user-interaction-task-source}
        given `element`{.variable} to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-apis:concept-event-fire
        x-internal="concept-event-fire"} named
        [`cancel`{#common-input-element-apis:event-cancel}](indices.html#event-cancel)
        at `element`{.variable}, with the
        [`bubbles`{#common-input-element-apis:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
        attribute initialized to true.

    5.  Otherwise, [update the file
        selection](#update-the-file-selection){#common-input-element-apis:update-the-file-selection}
        for `element`{.variable}.

    As with all user interface specifications, user agents have a good
    deal of freedom in how they interpret these requirements. The above
    text implies that a user either dismisses the prompt or changes
    their selection; exactly one of these will be true. But the mapping
    of these possibilities to specific user interface elements is not
    mandated by the standard. For example, a user agent might interpret
    clicking the \"Cancel\" button when files were previously selected
    as a change of selection to select zero files, thus firing
    [`input`{#common-input-element-apis:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    and
    [`change`{#common-input-element-apis:event-change}](indices.html#event-change).
    Or it might interpret such a click as a dismissal that leaves the
    selection unchanged, thus firing
    [`cancel`{#common-input-element-apis:event-cancel-2}](indices.html#event-cancel).
    Similarly, it\'s up to the user agent whether re-selecting the same
    files as were previously selected counts as a dismissal, or as a
    change of selection.

6.  Otherwise, the user agent should show the relevant user interface
    for selecting a value for `element`{.variable}, in the way it
    normally would when the user interacts with the control.

    When showing such a user interface, it must respect the requirements
    stated in the relevant parts of the specification for how
    `element`{.variable} behaves given its
    [`type`{#common-input-element-apis:attr-input-type-10}](#attr-input-type)
    attribute state. (For example, various sections describe
    restrictions on the resulting
    [value](form-control-infrastructure.html#concept-fe-value){#common-input-element-apis:concept-fe-value-21}
    string.)

    This step can have side effects, such as closing other pickers that
    were previously shown by this algorithm. (If this closes a file
    selection picker, then per the above that will lead to firing either
    [`input`{#common-input-element-apis:event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
    and
    [`change`{#common-input-element-apis:event-change-2}](indices.html#event-change)
    events, or a
    [`cancel`{#common-input-element-apis:event-cancel-3}](indices.html#event-cancel)
    event.)

##### [4.10.5.5]{.secno} Common event behaviors[](#common-input-element-events){.self-link} {#common-input-element-events}

When the
[`input`{#common-input-element-events:event-input}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
and
[`change`{#common-input-element-events:event-change}](indices.html#event-change)
events
[apply](#concept-input-apply){#common-input-element-events:concept-input-apply}
(which is the case for all
[`input`{#common-input-element-events:the-input-element}](#the-input-element)
controls other than
[buttons](forms.html#concept-button){#common-input-element-events:concept-button}
and those with the
[`type`{#common-input-element-events:attr-input-type}](#attr-input-type)
attribute in the
[Hidden](#hidden-state-(type=hidden)){#common-input-element-events:hidden-state-(type=hidden)}
state), the events are fired to indicate that the user has interacted
with the control. The
[`input`{#common-input-element-events:event-input-2}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
event fires whenever the user has modified the data of the control. The
[`change`{#common-input-element-events:event-change-2}](indices.html#event-change)
event fires when the value is committed, if that makes sense for the
control, or else when the control [loses
focus](interaction.html#unfocus-causes-change-event). In all cases, the
[`input`{#common-input-element-events:event-input-3}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
event comes before the corresponding
[`change`{#common-input-element-events:event-change-3}](indices.html#event-change)
event (if any).

When an
[`input`{#common-input-element-events:the-input-element-2}](#the-input-element)
element has a defined [input activation
behavior](#input-activation-behavior){#common-input-element-events:input-activation-behavior},
the rules for dispatching these events, if they
[apply](#concept-input-apply){#common-input-element-events:concept-input-apply-2},
are given in the section above that defines the
[`type`{#common-input-element-events:attr-input-type-2}](#attr-input-type)
attribute\'s state. (This is the case for all
[`input`{#common-input-element-events:the-input-element-3}](#the-input-element)
controls with the
[`type`{#common-input-element-events:attr-input-type-3}](#attr-input-type)
attribute in the
[Checkbox](#checkbox-state-(type=checkbox)){#common-input-element-events:checkbox-state-(type=checkbox)}
state, the [Radio
Button](#radio-button-state-(type=radio)){#common-input-element-events:radio-button-state-(type=radio)}
state, or the [File
Upload](#file-upload-state-(type=file)){#common-input-element-events:file-upload-state-(type=file)}
state.)

For
[`input`{#common-input-element-events:the-input-element-4}](#the-input-element)
elements without a defined [input activation
behavior](#input-activation-behavior){#common-input-element-events:input-activation-behavior-2},
but to which these events
[apply](#concept-input-apply){#common-input-element-events:concept-input-apply-3},
and for which the user interface involves both interactive manipulation
and an explicit commit action, then when the user changes the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value},
the user agent must [queue an element
task](webappapis.html#queue-an-element-task){#common-input-element-events:queue-an-element-task}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#common-input-element-events:user-interaction-task-source}
given the
[`input`{#common-input-element-events:the-input-element-5}](#the-input-element)
element to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-events:concept-event-fire
x-internal="concept-event-fire"} named
[`input`{#common-input-element-events:event-input-4}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
at the
[`input`{#common-input-element-events:the-input-element-6}](#the-input-element)
element, with the
[`bubbles`{#common-input-element-events:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
and
[`composed`{#common-input-element-events:dom-event-composed}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
attributes initialized to true, and any time the user commits the
change, the user agent must [queue an element
task](webappapis.html#queue-an-element-task){#common-input-element-events:queue-an-element-task-2}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#common-input-element-events:user-interaction-task-source-2}
given the
[`input`{#common-input-element-events:the-input-element-7}](#the-input-element)
element to set its [user
validity](form-control-infrastructure.html#user-validity){#common-input-element-events:user-validity}
to true and [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-events:concept-event-fire-2
x-internal="concept-event-fire"} named
[`change`{#common-input-element-events:event-change-4}](indices.html#event-change)
at the
[`input`{#common-input-element-events:the-input-element-8}](#the-input-element)
element, with the
[`bubbles`{#common-input-element-events:dom-event-bubbles-2}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
attribute initialized to true.

An example of a user interface involving both interactive manipulation
and a commit action would be a
[Range](#range-state-(type=range)){#common-input-element-events:range-state-(type=range)}
controls that use a slider, when manipulated using a pointing device.
While the user is dragging the control\'s knob,
[`input`{#common-input-element-events:event-input-5}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
events would fire whenever the position changed, whereas the
[`change`{#common-input-element-events:event-change-5}](indices.html#event-change)
event would only fire when the user let go of the knob, committing to a
specific value.

For
[`input`{#common-input-element-events:the-input-element-9}](#the-input-element)
elements without a defined [input activation
behavior](#input-activation-behavior){#common-input-element-events:input-activation-behavior-3},
but to which these events
[apply](#concept-input-apply){#common-input-element-events:concept-input-apply-4},
and for which the user interface involves an explicit commit action but
no intermediate manipulation, then any time the user commits a change to
the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value-2},
the user agent must [queue an element
task](webappapis.html#queue-an-element-task){#common-input-element-events:queue-an-element-task-3}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#common-input-element-events:user-interaction-task-source-3}
given the
[`input`{#common-input-element-events:the-input-element-10}](#the-input-element)
element to first [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-events:concept-event-fire-3
x-internal="concept-event-fire"} named
[`input`{#common-input-element-events:event-input-6}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
at the
[`input`{#common-input-element-events:the-input-element-11}](#the-input-element)
element, with the
[`bubbles`{#common-input-element-events:dom-event-bubbles-3}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
and
[`composed`{#common-input-element-events:dom-event-composed-2}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
attributes initialized to true, and then [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-events:concept-event-fire-4
x-internal="concept-event-fire"} named
[`change`{#common-input-element-events:event-change-6}](indices.html#event-change)
at the
[`input`{#common-input-element-events:the-input-element-12}](#the-input-element)
element, with the
[`bubbles`{#common-input-element-events:dom-event-bubbles-4}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
attribute initialized to true.

An example of a user interface with a commit action would be a
[Color](#color-state-(type=color)){#common-input-element-events:color-state-(type=color)}
control that consists of a single button that brings up a color wheel:
if the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value-3}
only changes when the dialog is closed, then that would be the explicit
commit action. On the other hand, if manipulating the control changes
the color interactively, then there might be no commit action.

Another example of a user interface with a commit action would be a
[Date](#date-state-(type=date)){#common-input-element-events:date-state-(type=date)}
control that allows both text-based user input and user selection from a
drop-down calendar: while text input might not have an explicit commit
step, selecting a date from the drop down calendar and then dismissing
the drop down would be a commit action.

For
[`input`{#common-input-element-events:the-input-element-13}](#the-input-element)
elements without a defined [input activation
behavior](#input-activation-behavior){#common-input-element-events:input-activation-behavior-4},
but to which these events
[apply](#concept-input-apply){#common-input-element-events:concept-input-apply-5},
any time the user causes the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value-4}
to change without an explicit commit action, the user agent must [queue
an element
task](webappapis.html#queue-an-element-task){#common-input-element-events:queue-an-element-task-4}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#common-input-element-events:user-interaction-task-source-4}
given the
[`input`{#common-input-element-events:the-input-element-14}](#the-input-element)
element to [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-events:concept-event-fire-5
x-internal="concept-event-fire"} named
[`input`{#common-input-element-events:event-input-7}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
at the
[`input`{#common-input-element-events:the-input-element-15}](#the-input-element)
element, with the
[`bubbles`{#common-input-element-events:dom-event-bubbles-5}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
and
[`composed`{#common-input-element-events:dom-event-composed-3}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
attributes initialized to true. The corresponding
[`change`{#common-input-element-events:event-change-7}](indices.html#event-change)
event, if any, will be fired when the control [loses
focus](interaction.html#unfocus-causes-change-event).

Examples of a user changing the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value-5}
would include the user typing into a text control, pasting a new value
into the control, or undoing an edit in that control. Some user
interactions do not cause changes to the value, e.g., hitting the
\"delete\" key in an empty text control, or replacing some text in the
control with text from the clipboard that happens to be exactly the same
text.

A
[Range](#range-state-(type=range)){#common-input-element-events:range-state-(type=range)-2}
control in the form of a slider that the user has
[focused](interaction.html#focused){#common-input-element-events:focused}
and is interacting with using a keyboard would be another example of the
user changing the element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value-6}
without a commit step.

In the case of
[tasks](webappapis.html#concept-task){#common-input-element-events:concept-task}
that just fire an
[`input`{#common-input-element-events:event-input-8}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
event, user agents may wait for a suitable break in the user\'s
interaction before
[queuing](webappapis.html#queue-an-element-task){#common-input-element-events:queue-an-element-task-5}
the tasks; for example, a user agent could wait for the user to have not
hit a key for 100ms, so as to only fire the event when the user pauses,
instead of continuously for each keystroke.

When the user agent is to change an
[`input`{#common-input-element-events:the-input-element-16}](#the-input-element)
element\'s
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value-7}
on behalf of the user (e.g. as part of a form prefilling feature), the
user agent must [queue an element
task](webappapis.html#queue-an-element-task){#common-input-element-events:queue-an-element-task-6}
on the [user interaction task
source](webappapis.html#user-interaction-task-source){#common-input-element-events:user-interaction-task-source-5}
given the
[`input`{#common-input-element-events:the-input-element-17}](#the-input-element)
element to first update the
[value](form-control-infrastructure.html#concept-fe-value){#common-input-element-events:concept-fe-value-8}
accordingly, then [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-events:concept-event-fire-6
x-internal="concept-event-fire"} named
[`input`{#common-input-element-events:event-input-9}](https://w3c.github.io/uievents/#event-type-input){x-internal="event-input"}
at the
[`input`{#common-input-element-events:the-input-element-18}](#the-input-element)
element, with the
[`bubbles`{#common-input-element-events:dom-event-bubbles-6}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
and
[`composed`{#common-input-element-events:dom-event-composed-4}](https://dom.spec.whatwg.org/#dom-event-composed){x-internal="dom-event-composed"}
attributes initialized to true, then [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#common-input-element-events:concept-event-fire-7
x-internal="concept-event-fire"} named
[`change`{#common-input-element-events:event-change-8}](indices.html#event-change)
at the
[`input`{#common-input-element-events:the-input-element-19}](#the-input-element)
element, with the
[`bubbles`{#common-input-element-events:dom-event-bubbles-7}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
attribute initialized to true.

These events are not fired in response to changes made to the values of
form controls by scripts. (This is to make it easier to update the
values of form controls in response to the user manipulating the
controls, without having to then filter out the script\'s own changes to
avoid an infinite loop.)

These events are also not fired when the browser changes the values of
form controls as part of [state restoration during
navigation](browsing-the-web.html#restore-persisted-state){#common-input-element-events:restore-persisted-state}.

[← 4.10 Forms](forms.html) --- [Table of Contents](./) --- [4.10.6 The
button element →](form-elements.html)
