HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.8.15 MathML](embedded-content-other.html) --- [Table of
Contents](./) --- [4.10 Forms →](forms.html)

1.  ::: {#toc-semantics}
    1.  [[4.9]{.secno} Tabular data](tables.html#tables)
        1.  [[4.9.1]{.secno} The `table`
            element](tables.html#the-table-element)
            1.  [[4.9.1.1]{.secno} Techniques for describing
                tables](tables.html#table-descriptions-techniques)
            2.  [[4.9.1.2]{.secno} Techniques for table
                design](tables.html#table-layout-techniques)
        2.  [[4.9.2]{.secno} The `caption`
            element](tables.html#the-caption-element)
        3.  [[4.9.3]{.secno} The `colgroup`
            element](tables.html#the-colgroup-element)
        4.  [[4.9.4]{.secno} The `col`
            element](tables.html#the-col-element)
        5.  [[4.9.5]{.secno} The `tbody`
            element](tables.html#the-tbody-element)
        6.  [[4.9.6]{.secno} The `thead`
            element](tables.html#the-thead-element)
        7.  [[4.9.7]{.secno} The `tfoot`
            element](tables.html#the-tfoot-element)
        8.  [[4.9.8]{.secno} The `tr`
            element](tables.html#the-tr-element)
        9.  [[4.9.9]{.secno} The `td`
            element](tables.html#the-td-element)
        10. [[4.9.10]{.secno} The `th`
            element](tables.html#the-th-element)
        11. [[4.9.11]{.secno} Attributes common to `td` and `th`
            elements](tables.html#attributes-common-to-td-and-th-elements)
        12. [[4.9.12]{.secno} Processing
            model](tables.html#table-processing-model)
            1.  [[4.9.12.1]{.secno} Forming a
                table](tables.html#forming-a-table)
            2.  [[4.9.12.2]{.secno} Forming relationships between data
                cells and header
                cells](tables.html#header-and-data-cell-semantics)
        13. [[4.9.13]{.secno} Examples](tables.html#table-examples)
    :::

### [4.9]{.secno} Tabular data[](#tables){.self-link} {#tables}

#### [4.9.1]{.secno} The [`table`]{.dfn dfn-type="element"} element[](#the-table-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/table](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table "The <table> HTML element represents tabular data — that is, information presented in a two-dimensional table comprised of rows and columns of cells containing data.")

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
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement "The HTMLTableElement interface provides special properties and methods (beyond the regular HTMLElement object interface it also has available to it by inheritance) for manipulating the layout and presentation of tables in an HTML document.")

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

[Categories](dom.html#concept-element-categories){#the-table-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-table-element:flow-content-2}.
:   [Palpable
    content](dom.html#palpable-content-2){#the-table-element:palpable-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-table-element:concept-element-contexts}:
:   Where [flow
    content](dom.html#flow-content-2){#the-table-element:flow-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-table-element:concept-element-content-model}:
:   In this order: optionally a
    [`caption`{#the-table-element:the-caption-element}](#the-caption-element)
    element, followed by zero or more
    [`colgroup`{#the-table-element:the-colgroup-element}](#the-colgroup-element)
    elements, followed optionally by a
    [`thead`{#the-table-element:the-thead-element}](#the-thead-element)
    element, followed by either zero or more
    [`tbody`{#the-table-element:the-tbody-element}](#the-tbody-element)
    elements or one or more
    [`tr`{#the-table-element:the-tr-element}](#the-tr-element) elements,
    followed optionally by a
    [`tfoot`{#the-table-element:the-tfoot-element}](#the-tfoot-element)
    element, optionally intermixed with one or more [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-table-element:script-supporting-elements-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-table-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-table-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-table-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-table-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-table).
:   [For implementers](https://w3c.github.io/html-aam/#el-table).

[DOM interface](dom.html#concept-element-dom){#the-table-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTableElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute HTMLTableCaptionElement? caption;
      HTMLTableCaptionElement createCaption();
      [CEReactions] undefined deleteCaption();

      [CEReactions] attribute HTMLTableSectionElement? tHead;
      HTMLTableSectionElement createTHead();
      [CEReactions] undefined deleteTHead();

      [CEReactions] attribute HTMLTableSectionElement? tFoot;
      HTMLTableSectionElement createTFoot();
      [CEReactions] undefined deleteTFoot();

      [SameObject] readonly attribute HTMLCollection tBodies;
      HTMLTableSectionElement createTBody();

      [SameObject] readonly attribute HTMLCollection rows;
      HTMLTableRowElement insertRow(optional long index = -1);
      [CEReactions] undefined deleteRow(long index);

      // also has obsolete members
    };
    ```

The [`table`{#the-table-element:the-table-element}](#the-table-element)
element [represents](dom.html#represents){#the-table-element:represents}
data with more than one dimension, in the form of a
[table](#concept-table){#the-table-element:concept-table}.

The
[`table`{#the-table-element:the-table-element-2}](#the-table-element)
element takes part in the [table
model](#table-model){#the-table-element:table-model}. Tables have rows,
columns, and cells given by their descendants. The rows and columns form
a grid; a table\'s cells must completely cover that grid without
overlap.

Precise rules for determining whether this conformance requirement is
met are described in the description of the [table
model](#table-model){#the-table-element:table-model-2}.

Authors are encouraged to provide information describing how to
interpret complex tables. Guidance on how to [provide such
information](#table-descriptions-techniques) is given below.

Tables must not be used as layout aids. Historically, some web authors
have misused tables in HTML as a way to control their page layout. This
usage is non-conforming, because tools attempting to extract tabular
data from such documents would obtain very confusing results. In
particular, users of accessibility tools like screen readers are likely
to find it very difficult to navigate pages with tables used for layout.

There are a variety of alternatives to using HTML tables for layout,
such as CSS grid layout, CSS flexible box layout (\"flexbox\"), CSS
multi-column layout, CSS positioning, and the CSS table model.
[\[CSS\]](references.html#refsCSS)

------------------------------------------------------------------------

Tables can be complicated to understand and navigate. To help users with
this, user agents should clearly delineate cells in a table from each
other, unless the user agent has classified the table as a
(non-conforming) layout table.

Authors and implementers are encouraged to consider using some of the
[table design techniques](#table-layout-techniques) described below to
make tables easier to navigate for users.

User agents, especially those that do table analysis on arbitrary
content, are encouraged to find heuristics to determine which tables
actually contain data and which are merely being used for layout. This
specification does not define a precise heuristic, but the following are
suggested as possible indicators:

Feature

Indication

The use of the
[`role`{#the-table-element:attr-aria-role}](infrastructure.html#attr-aria-role)
attribute with the value
[`presentation`{#the-table-element:attr-aria-role-presentation}](https://w3c.github.io/aria/#presentation){x-internal="attr-aria-role-presentation"}

Probably a layout table

The use of the non-conforming
[`border`{#the-table-element:attr-table-border}](obsolete.html#attr-table-border)
attribute with the non-conforming value 0

Probably a layout table

The use of the non-conforming
[`cellspacing`{#the-table-element:attr-table-cellspacing}](obsolete.html#attr-table-cellspacing)
and
[`cellpadding`{#the-table-element:attr-table-cellpadding}](obsolete.html#attr-table-cellpadding)
attributes with the value 0

Probably a layout table

The use of
[`caption`{#the-table-element:the-caption-element-2}](#the-caption-element),
[`thead`{#the-table-element:the-thead-element-2}](#the-thead-element),
or [`th`{#the-table-element:the-th-element}](#the-th-element) elements

Probably a non-layout table

The use of the
[`headers`{#the-table-element:attr-tdth-headers}](#attr-tdth-headers)
and [`scope`{#the-table-element:attr-th-scope}](#attr-th-scope)
attributes

Probably a non-layout table

The use of the non-conforming
[`border`{#the-table-element:attr-table-border-2}](obsolete.html#attr-table-border)
attribute with a value other than 0

Probably a non-layout table

Explicit visible borders set using CSS

Probably a non-layout table

The use of the
[`summary`{#the-table-element:attr-table-summary}](obsolete.html#attr-table-summary)
attribute

Not a good indicator (both layout and non-layout tables have
historically been given this attribute)

It is quite possible that the above suggestions are wrong. Implementers
are urged to provide feedback elaborating on their experiences with
trying to create a layout table detection heuristic.

If a
[`table`{#the-table-element:the-table-element-3}](#the-table-element)
element has a (non-conforming)
[`summary`{#the-table-element:attr-table-summary-2}](obsolete.html#attr-table-summary)
attribute, and the user agent has not classified the table as a layout
table, the user agent may report the contents of that attribute to the
user.

------------------------------------------------------------------------

`table`{.variable}`.`[`caption`](#dom-table-caption){#dom-table-caption-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/caption](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/caption "The HTMLTableElement.caption property represents the table caption. If no caption element is associated with the table, this property is null.")

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

Returns the table\'s
[`caption`{#the-table-element:the-caption-element-3}](#the-caption-element)
element.

Can be set, to replace the
[`caption`{#the-table-element:the-caption-element-4}](#the-caption-element)
element.

`caption`{.variable}` = ``table`{.variable}`.`[`createCaption`](#dom-table-createcaption){#dom-table-createcaption-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/createCaption](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/createCaption "The HTMLTableElement.createCaption() method returns the <caption> element associated with a given <table>. If no <caption> element exists on the table, this method creates it, and then returns it.")

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

Ensures the table has a
[`caption`{#the-table-element:the-caption-element-5}](#the-caption-element)
element, and returns it.

`table`{.variable}`.`[`deleteCaption`](#dom-table-deletecaption){#dom-table-deletecaption-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/deleteCaption](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/deleteCaption "The HTMLTableElement.deleteCaption() method removes the <caption> element from a given <table>. If there is no <caption> element associated with the table, this method does nothing.")

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

Ensures the table does not have a
[`caption`{#the-table-element:the-caption-element-6}](#the-caption-element)
element.

`table`{.variable}`.`[`tHead`](#dom-table-thead){#dom-table-thead-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/tHead](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/tHead "The HTMLTableElement.tHead represents the <thead> element of a <table>. Its value will be null if there is no such element.")

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

Returns the table\'s
[`thead`{#the-table-element:the-thead-element-3}](#the-thead-element)
element.

Can be set, to replace the
[`thead`{#the-table-element:the-thead-element-4}](#the-thead-element)
element. If the new value is not a
[`thead`{#the-table-element:the-thead-element-5}](#the-thead-element)
element, throws a
[\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#the-table-element:hierarchyrequesterror
x-internal="hierarchyrequesterror"}
[`DOMException`{#the-table-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

`thead`{.variable}` = ``table`{.variable}`.`[`createTHead`](#dom-table-createthead){#dom-table-createthead-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/createTHead](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/createTHead "The createTHead() method of HTMLTableElement objects returns the <thead> element associated with a given <table>. If no header exists in the table, this method creates it, and then returns it.")

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

Ensures the table has a
[`thead`{#the-table-element:the-thead-element-6}](#the-thead-element)
element, and returns it.

`table`{.variable}`.`[`deleteTHead`](#dom-table-deletethead){#dom-table-deletethead-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/deleteTHead](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/deleteTHead "The HTMLTableElement.deleteTHead() removes the <thead> element from a given <table>.")

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

Ensures the table does not have a
[`thead`{#the-table-element:the-thead-element-7}](#the-thead-element)
element.

`table`{.variable}`.`[`tFoot`](#dom-table-tfoot){#dom-table-tfoot-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/tFoot](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/tFoot "The HTMLTableElement.tFoot property represents the <tfoot> element of a <table>. Its value will be null if there is no such element.")

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

Returns the table\'s
[`tfoot`{#the-table-element:the-tfoot-element-2}](#the-tfoot-element)
element.

Can be set, to replace the
[`tfoot`{#the-table-element:the-tfoot-element-3}](#the-tfoot-element)
element. If the new value is not a
[`tfoot`{#the-table-element:the-tfoot-element-4}](#the-tfoot-element)
element, throws a
[\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#the-table-element:hierarchyrequesterror-2
x-internal="hierarchyrequesterror"}
[`DOMException`{#the-table-element:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

`tfoot`{.variable}` = ``table`{.variable}`.`[`createTFoot`](#dom-table-createtfoot){#dom-table-createtfoot-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/createTFoot](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/createTFoot "The createTFoot() method of HTMLTableElement objects returns the <tfoot> element associated with a given <table>. If no footer exists in the table, this method creates it, and then returns it.")

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

Ensures the table has a
[`tfoot`{#the-table-element:the-tfoot-element-5}](#the-tfoot-element)
element, and returns it.

`table`{.variable}`.`[`deleteTFoot`](#dom-table-deletetfoot){#dom-table-deletetfoot-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/deleteTFoot](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/deleteTFoot "The HTMLTableElement.deleteTFoot() method removes the <tfoot> element from a given <table>.")

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

Ensures the table does not have a
[`tfoot`{#the-table-element:the-tfoot-element-6}](#the-tfoot-element)
element.

`table`{.variable}`.`[`tBodies`](#dom-table-tbodies){#dom-table-tbodies-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/tBodies](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/tBodies "The HTMLTableElement.tBodies read-only property returns a live HTMLCollection of the bodies in a <table>.")

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
[`HTMLCollection`{#the-table-element:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the
[`tbody`{#the-table-element:the-tbody-element-2}](#the-tbody-element)
elements of the table.

`tbody`{.variable}` = ``table`{.variable}`.`[`createTBody`](#dom-table-createtbody){#dom-table-createtbody-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/createTBody](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/createTBody "The createTBody() method of HTMLTableElement objects creates and returns a new <tbody> element associated with a given <table>.")

Support in all current engines.

::: support
[Firefox25+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome20+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Creates a
[`tbody`{#the-table-element:the-tbody-element-3}](#the-tbody-element)
element, inserts it into the table, and returns it.

`table`{.variable}`.`[`rows`](#dom-table-rows){#dom-table-rows-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/rows](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/rows "The read-only HTMLTableElement property rows returns a live HTMLCollection of all the rows in the table, including the rows contained within any <thead>, <tfoot>, and <tbody> elements.")

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
[`HTMLCollection`{#the-table-element:htmlcollection-4}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the [`tr`{#the-table-element:the-tr-element-2}](#the-tr-element)
elements of the table.

`tr`{.variable}` = ``table`{.variable}`.`[`insertRow`](#dom-table-insertrow){#dom-table-insertrow-dev}`([ ``index`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/insertRow](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/insertRow "The HTMLTableElement.insertRow() method inserts a new row (<tr>) in a given <table>, and returns a reference to the new row.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS3+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Creates a [`tr`{#the-table-element:the-tr-element-3}](#the-tr-element)
element, along with a
[`tbody`{#the-table-element:the-tbody-element-4}](#the-tbody-element) if
required, inserts them into the table at the position given by the
argument, and returns the
[`tr`{#the-table-element:the-tr-element-4}](#the-tr-element).

The position is relative to the rows in the table. The index −1, which
is the default if the argument is omitted, is equivalent to inserting at
the end of the table.

If the given position is less than −1 or greater than the number of
rows, throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-table-element:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#the-table-element:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

`table`{.variable}`.`[`deleteRow`](#dom-table-deleterow){#dom-table-deleterow-dev}`(``index`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableElement/deleteRow](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableElement/deleteRow "The HTMLTableElement.deleteRow() method removes a specific row (<tr>) from a given <table>.")

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

Removes the [`tr`{#the-table-element:the-tr-element-5}](#the-tr-element)
element with the given position in the table.

The position is relative to the rows in the table. The index −1 is
equivalent to deleting the last row of the table.

If the given position is less than −1 or greater than the index of the
last row, or if there are no rows, throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-table-element:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#the-table-element:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

In all of the following attribute and method definitions, when an
element is to be [table-created]{#table-created .dfn}, that means to
[create an
element](https://dom.spec.whatwg.org/#concept-create-element){#the-table-element:create-an-element
x-internal="create-an-element"} given the
[`table`{#the-table-element:the-table-element-4}](#the-table-element)
element\'s [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-table-element:node-document
x-internal="node-document"}, the given local name, and the [HTML
namespace](https://infra.spec.whatwg.org/#html-namespace){#the-table-element:html-namespace-2
x-internal="html-namespace-2"}.

The [`caption`]{#dom-table-caption .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"} IDL attribute must return, on getting, the first
[`caption`{#the-table-element:the-caption-element-7}](#the-caption-element)
element child of the
[`table`{#the-table-element:the-table-element-5}](#the-table-element)
element, if any, or null otherwise. On setting, the first
[`caption`{#the-table-element:the-caption-element-8}](#the-caption-element)
element child of the
[`table`{#the-table-element:the-table-element-6}](#the-table-element)
element, if any, must be removed, and the new value, if not null, must
be inserted as the first node of the
[`table`{#the-table-element:the-table-element-7}](#the-table-element)
element.

The [`createCaption()`]{#dom-table-createcaption .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method must return the
first
[`caption`{#the-table-element:the-caption-element-9}](#the-caption-element)
element child of the
[`table`{#the-table-element:the-table-element-8}](#the-table-element)
element, if any; otherwise a new
[`caption`{#the-table-element:the-caption-element-10}](#the-caption-element)
element must be
[table-created](#table-created){#the-table-element:table-created},
inserted as the first node of the
[`table`{#the-table-element:the-table-element-9}](#the-table-element)
element, and then returned.

The [`deleteCaption()`]{#dom-table-deletecaption .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method must remove the
first
[`caption`{#the-table-element:the-caption-element-11}](#the-caption-element)
element child of the
[`table`{#the-table-element:the-table-element-10}](#the-table-element)
element, if any.

The [`tHead`]{#dom-table-thead .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"} IDL attribute must return, on getting, the first
[`thead`{#the-table-element:the-thead-element-8}](#the-thead-element)
element child of the
[`table`{#the-table-element:the-table-element-11}](#the-table-element)
element, if any, or null otherwise. On setting, if the new value is null
or a
[`thead`{#the-table-element:the-thead-element-9}](#the-thead-element)
element, the first
[`thead`{#the-table-element:the-thead-element-10}](#the-thead-element)
element child of the
[`table`{#the-table-element:the-table-element-12}](#the-table-element)
element, if any, must be removed, and the new value, if not null, must
be inserted immediately before the first element in the
[`table`{#the-table-element:the-table-element-13}](#the-table-element)
element that is neither a
[`caption`{#the-table-element:the-caption-element-12}](#the-caption-element)
element nor a
[`colgroup`{#the-table-element:the-colgroup-element-2}](#the-colgroup-element)
element, if any, or at the end of the table if there are no such
elements. If the new value is neither null nor a
[`thead`{#the-table-element:the-thead-element-11}](#the-thead-element)
element, then a
[\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#the-table-element:hierarchyrequesterror-3
x-internal="hierarchyrequesterror"}
[`DOMException`{#the-table-element:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
must be thrown instead.

The [`createTHead()`]{#dom-table-createthead .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method must return the
first
[`thead`{#the-table-element:the-thead-element-12}](#the-thead-element)
element child of the
[`table`{#the-table-element:the-table-element-14}](#the-table-element)
element, if any; otherwise a new
[`thead`{#the-table-element:the-thead-element-13}](#the-thead-element)
element must be
[table-created](#table-created){#the-table-element:table-created-2} and
inserted immediately before the first element in the
[`table`{#the-table-element:the-table-element-15}](#the-table-element)
element that is neither a
[`caption`{#the-table-element:the-caption-element-13}](#the-caption-element)
element nor a
[`colgroup`{#the-table-element:the-colgroup-element-3}](#the-colgroup-element)
element, if any, or at the end of the table if there are no such
elements, and then that new element must be returned.

The [`deleteTHead()`]{#dom-table-deletethead .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method must remove the
first
[`thead`{#the-table-element:the-thead-element-14}](#the-thead-element)
element child of the
[`table`{#the-table-element:the-table-element-16}](#the-table-element)
element, if any.

The [`tFoot`]{#dom-table-tfoot .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"} IDL attribute must return, on getting, the first
[`tfoot`{#the-table-element:the-tfoot-element-7}](#the-tfoot-element)
element child of the
[`table`{#the-table-element:the-table-element-17}](#the-table-element)
element, if any, or null otherwise. On setting, if the new value is null
or a
[`tfoot`{#the-table-element:the-tfoot-element-8}](#the-tfoot-element)
element, the first
[`tfoot`{#the-table-element:the-tfoot-element-9}](#the-tfoot-element)
element child of the
[`table`{#the-table-element:the-table-element-18}](#the-table-element)
element, if any, must be removed, and the new value, if not null, must
be inserted at the end of the table. If the new value is neither null
nor a
[`tfoot`{#the-table-element:the-tfoot-element-10}](#the-tfoot-element)
element, then a
[\"`HierarchyRequestError`\"](https://webidl.spec.whatwg.org/#hierarchyrequesterror){#the-table-element:hierarchyrequesterror-4
x-internal="hierarchyrequesterror"}
[`DOMException`{#the-table-element:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
must be thrown instead.

The [`createTFoot()`]{#dom-table-createtfoot .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method must return the
first
[`tfoot`{#the-table-element:the-tfoot-element-11}](#the-tfoot-element)
element child of the
[`table`{#the-table-element:the-table-element-19}](#the-table-element)
element, if any; otherwise a new
[`tfoot`{#the-table-element:the-tfoot-element-12}](#the-tfoot-element)
element must be
[table-created](#table-created){#the-table-element:table-created-3} and
inserted at the end of the table, and then that new element must be
returned.

The [`deleteTFoot()`]{#dom-table-deletetfoot .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method must remove the
first
[`tfoot`{#the-table-element:the-tfoot-element-13}](#the-tfoot-element)
element child of the
[`table`{#the-table-element:the-table-element-20}](#the-table-element)
element, if any.

The [`tBodies`]{#dom-table-tbodies .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#the-table-element:htmlcollection-5}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`table`{#the-table-element:the-table-element-21}](#the-table-element)
node, whose filter matches only
[`tbody`{#the-table-element:the-tbody-element-5}](#the-tbody-element)
elements that are children of the
[`table`{#the-table-element:the-table-element-22}](#the-table-element)
element.

The [`createTBody()`]{#dom-table-createtbody .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method must
[table-create](#table-created){#the-table-element:table-created-4} a new
[`tbody`{#the-table-element:the-tbody-element-6}](#the-tbody-element)
element, insert it immediately after the last
[`tbody`{#the-table-element:the-tbody-element-7}](#the-tbody-element)
element child in the
[`table`{#the-table-element:the-table-element-23}](#the-table-element)
element, if any, or at the end of the
[`table`{#the-table-element:the-table-element-24}](#the-table-element)
element if the
[`table`{#the-table-element:the-table-element-25}](#the-table-element)
element has no
[`tbody`{#the-table-element:the-tbody-element-8}](#the-tbody-element)
element children, and then must return the new
[`tbody`{#the-table-element:the-tbody-element-9}](#the-tbody-element)
element.

The [`rows`]{#dom-table-rows .dfn dfn-for="HTMLTableElement"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#the-table-element:htmlcollection-6}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at the
[`table`{#the-table-element:the-table-element-26}](#the-table-element)
node, whose filter matches only
[`tr`{#the-table-element:the-tr-element-6}](#the-tr-element) elements
that are either children of the
[`table`{#the-table-element:the-table-element-27}](#the-table-element)
element, or children of
[`thead`{#the-table-element:the-thead-element-15}](#the-thead-element),
[`tbody`{#the-table-element:the-tbody-element-10}](#the-tbody-element),
or
[`tfoot`{#the-table-element:the-tfoot-element-14}](#the-tfoot-element)
elements that are themselves children of the
[`table`{#the-table-element:the-table-element-28}](#the-table-element)
element. The elements in the collection must be ordered such that those
elements whose parent is a
[`thead`{#the-table-element:the-thead-element-16}](#the-thead-element)
are included first, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-table-element:tree-order
x-internal="tree-order"}, followed by those elements whose parent is
either a
[`table`{#the-table-element:the-table-element-29}](#the-table-element)
or
[`tbody`{#the-table-element:the-tbody-element-11}](#the-tbody-element)
element, again in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-table-element:tree-order-2
x-internal="tree-order"}, followed finally by those elements whose
parent is a
[`tfoot`{#the-table-element:the-tfoot-element-15}](#the-tfoot-element)
element, still in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#the-table-element:tree-order-3
x-internal="tree-order"}.

The behavior of the
[`insertRow(``index`{.variable}`)`]{#dom-table-insertrow .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method depends on the
state of the table. When it is called, the method must act as required
by the first item in the following list of conditions that describes the
state of the table and the `index`{.variable} argument:

If `index`{.variable} is less than −1 or greater than the number of elements in [`rows`{#the-table-element:dom-table-rows-2}](#dom-table-rows) collection:
:   The method must throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-table-element:indexsizeerror-3
    x-internal="indexsizeerror"}
    [`DOMException`{#the-table-element:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

If the [`rows`{#the-table-element:dom-table-rows-3}](#dom-table-rows) collection has zero elements in it, and the [`table`{#the-table-element:the-table-element-30}](#the-table-element) has no [`tbody`{#the-table-element:the-tbody-element-12}](#the-tbody-element) elements in it:
:   The method must
    [table-create](#table-created){#the-table-element:table-created-5} a
    [`tbody`{#the-table-element:the-tbody-element-13}](#the-tbody-element)
    element, then
    [table-create](#table-created){#the-table-element:table-created-6} a
    [`tr`{#the-table-element:the-tr-element-7}](#the-tr-element)
    element, then append the
    [`tr`{#the-table-element:the-tr-element-8}](#the-tr-element) element
    to the
    [`tbody`{#the-table-element:the-tbody-element-14}](#the-tbody-element)
    element, then append the
    [`tbody`{#the-table-element:the-tbody-element-15}](#the-tbody-element)
    element to the
    [`table`{#the-table-element:the-table-element-31}](#the-table-element)
    element, and finally return the
    [`tr`{#the-table-element:the-tr-element-9}](#the-tr-element)
    element.

If the [`rows`{#the-table-element:dom-table-rows-4}](#dom-table-rows) collection has zero elements in it:
:   The method must
    [table-create](#table-created){#the-table-element:table-created-7} a
    [`tr`{#the-table-element:the-tr-element-10}](#the-tr-element)
    element, append it to the last
    [`tbody`{#the-table-element:the-tbody-element-16}](#the-tbody-element)
    element in the table, and return the
    [`tr`{#the-table-element:the-tr-element-11}](#the-tr-element)
    element.

If `index`{.variable} is −1 or equal to the number of items in [`rows`{#the-table-element:dom-table-rows-5}](#dom-table-rows) collection:
:   The method must
    [table-create](#table-created){#the-table-element:table-created-8} a
    [`tr`{#the-table-element:the-tr-element-12}](#the-tr-element)
    element, and append it to the parent of the last
    [`tr`{#the-table-element:the-tr-element-13}](#the-tr-element)
    element in the
    [`rows`{#the-table-element:dom-table-rows-6}](#dom-table-rows)
    collection. Then, the newly created
    [`tr`{#the-table-element:the-tr-element-14}](#the-tr-element)
    element must be returned.

Otherwise:
:   The method must
    [table-create](#table-created){#the-table-element:table-created-9} a
    [`tr`{#the-table-element:the-tr-element-15}](#the-tr-element)
    element, insert it immediately before the `index`{.variable}th
    [`tr`{#the-table-element:the-tr-element-16}](#the-tr-element)
    element in the
    [`rows`{#the-table-element:dom-table-rows-7}](#dom-table-rows)
    collection, in the same parent, and finally must return the newly
    created
    [`tr`{#the-table-element:the-tr-element-17}](#the-tr-element)
    element.

When the [`deleteRow(``index`{.variable}`)`]{#dom-table-deleterow .dfn
dfn-for="HTMLTableElement" dfn-type="method"} method is called, the user
agent must run the following steps:

1.  If `index`{.variable} is less than −1 or greater than or equal to
    the number of elements in the
    [`rows`{#the-table-element:dom-table-rows-8}](#dom-table-rows)
    collection, then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-table-element:indexsizeerror-4
    x-internal="indexsizeerror"}
    [`DOMException`{#the-table-element:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If `index`{.variable} is −1, then
    [remove](https://dom.spec.whatwg.org/#concept-node-remove){#the-table-element:concept-node-remove
    x-internal="concept-node-remove"} the last element in the
    [`rows`{#the-table-element:dom-table-rows-9}](#dom-table-rows)
    collection from its parent, or do nothing if the
    [`rows`{#the-table-element:dom-table-rows-10}](#dom-table-rows)
    collection is empty.

3.  Otherwise,
    [remove](https://dom.spec.whatwg.org/#concept-node-remove){#the-table-element:concept-node-remove-2
    x-internal="concept-node-remove"} the `index`{.variable}th element
    in the
    [`rows`{#the-table-element:dom-table-rows-11}](#dom-table-rows)
    collection from its parent.

::: example
Here is an example of a table being used to mark up a Sudoku puzzle.
Observe the lack of headers, which are not necessary in such a table.

``` html
<style>
 #sudoku { border-collapse: collapse; border: solid thick; }
 #sudoku colgroup, table#sudoku tbody { border: solid medium; }
 #sudoku td { border: solid thin; height: 1.4em; width: 1.4em; text-align: center; padding: 0; }
</style>
<h1>Today's Sudoku</h1>
<table id="sudoku">
 <colgroup><col><col><col>
 <colgroup><col><col><col>
 <colgroup><col><col><col>
 <tbody>
  <tr> <td> 1 <td>   <td> 3 <td> 6 <td>   <td> 4 <td> 7 <td>   <td> 9
  <tr> <td>   <td> 2 <td>   <td>   <td> 9 <td>   <td>   <td> 1 <td>
  <tr> <td> 7 <td>   <td>   <td>   <td>   <td>   <td>   <td>   <td> 6
 <tbody>
  <tr> <td> 2 <td>   <td> 4 <td>   <td> 3 <td>   <td> 9 <td>   <td> 8
  <tr> <td>   <td>   <td>   <td>   <td>   <td>   <td>   <td>   <td>
  <tr> <td> 5 <td>   <td>   <td> 9 <td>   <td> 7 <td>   <td>   <td> 1
 <tbody>
  <tr> <td> 6 <td>   <td>   <td>   <td> 5 <td>   <td>   <td>   <td> 2
  <tr> <td>   <td>   <td>   <td>   <td> 7 <td>   <td>   <td>   <td>
  <tr> <td> 9 <td>   <td>   <td> 8 <td>   <td> 2 <td>   <td>   <td> 5
</table>
```
:::

##### [4.9.1.1]{.secno} Techniques for describing tables[](#table-descriptions-techniques){.self-link} {#table-descriptions-techniques}

For tables that consist of more than just a grid of cells with headers
in the first row and headers in the first column, and for any table in
general where the reader might have difficulty understanding the
content, authors should include explanatory information introducing the
table. This information is useful for all users, but is especially
useful for users who cannot see the table, e.g. users of screen readers.

Such explanatory information should introduce the purpose of the table,
outline its basic cell structure, highlight any trends or patterns, and
generally teach the user how to use the table.

For instance, the following table:

Characteristics with positive and negative sides

Negative

Characteristic

Positive

Sad

Mood

Happy

Failing

Grade

Passing

\...might benefit from a description explaining the way the table is
laid out, something like \"Characteristics are given in the second
column, with the negative side in the left column and the positive side
in the right column\".

There are a variety of ways to include this information, such as:

In prose, surrounding the table

:   ::: example
    ``` html
    <p>In the following table, characteristics are given in the second
    column, with the negative side in the left column and the positive
    side in the right column.</p>
    <table>
     <caption>Characteristics with positive and negative sides</caption>
     <thead>
      <tr>
       <th id="n"> Negative
       <th> Characteristic
       <th> Positive
     <tbody>
      <tr>
       <td headers="n r1"> Sad
       <th id="r1"> Mood
       <td> Happy
      <tr>
       <td headers="n r2"> Failing
       <th id="r2"> Grade
       <td> Passing
    </table>
    ```
    :::

In the table\'s [`caption`{#table-descriptions-techniques:the-caption-element}](#the-caption-element)

:   ::: example
    ``` html
    <table>
     <caption>
      <strong>Characteristics with positive and negative sides.</strong>
      <p>Characteristics are given in the second column, with the
      negative side in the left column and the positive side in the right
      column.</p>
     </caption>
     <thead>
      <tr>
       <th id="n"> Negative
       <th> Characteristic
       <th> Positive
     <tbody>
      <tr>
       <td headers="n r1"> Sad
       <th id="r1"> Mood
       <td> Happy
      <tr>
       <td headers="n r2"> Failing
       <th id="r2"> Grade
       <td> Passing
    </table>
    ```
    :::

In the table\'s [`caption`{#table-descriptions-techniques:the-caption-element-2}](#the-caption-element), in a [`details`{#table-descriptions-techniques:the-details-element}](interactive-elements.html#the-details-element) element

:   ::: example
    ``` html
    <table>
     <caption>
      <strong>Characteristics with positive and negative sides.</strong>
      <details>
       <summary>Help</summary>
       <p>Characteristics are given in the second column, with the
       negative side in the left column and the positive side in the right
       column.</p>
      </details>
     </caption>
     <thead>
      <tr>
       <th id="n"> Negative
       <th> Characteristic
       <th> Positive
     <tbody>
      <tr>
       <td headers="n r1"> Sad
       <th id="r1"> Mood
       <td> Happy
      <tr>
       <td headers="n r2"> Failing
       <th id="r2"> Grade
       <td> Passing
    </table>
    ```
    :::

Next to the table, in the same [`figure`{#table-descriptions-techniques:the-figure-element}](grouping-content.html#the-figure-element)

:   ::: example
    ``` html
    <figure>
     <figcaption>Characteristics with positive and negative sides</figcaption>
     <p>Characteristics are given in the second column, with the
     negative side in the left column and the positive side in the right
     column.</p>
     <table>
      <thead>
       <tr>
        <th id="n"> Negative
        <th> Characteristic
        <th> Positive
      <tbody>
       <tr>
        <td headers="n r1"> Sad
        <th id="r1"> Mood
        <td> Happy
       <tr>
        <td headers="n r2"> Failing
        <th id="r2"> Grade
        <td> Passing
     </table>
    </figure>
    ```
    :::

Next to the table, in a [`figure`{#table-descriptions-techniques:the-figure-element-2}](grouping-content.html#the-figure-element)\'s [`figcaption`{#table-descriptions-techniques:the-figcaption-element}](grouping-content.html#the-figcaption-element)

:   ::: example
    ``` html
    <figure>
     <figcaption>
      <strong>Characteristics with positive and negative sides</strong>
      <p>Characteristics are given in the second column, with the
      negative side in the left column and the positive side in the right
      column.</p>
     </figcaption>
     <table>
      <thead>
       <tr>
        <th id="n"> Negative
        <th> Characteristic
        <th> Positive
      <tbody>
       <tr>
        <td headers="n r1"> Sad
        <th id="r1"> Mood
        <td> Happy
       <tr>
        <td headers="n r2"> Failing
        <th id="r2"> Grade
        <td> Passing
     </table>
    </figure>
    ```
    :::

Authors may also use other techniques, or combinations of the above
techniques, as appropriate.

The best option, of course, rather than writing a description explaining
the way the table is laid out, is to adjust the table such that no
explanation is needed.

::: example
In the case of the table used in the examples above, a simple
rearrangement of the table so that the headers are on the top and left
sides removes the need for an explanation as well as removing the need
for the use of
[`headers`{#table-descriptions-techniques:attr-tdth-headers}](#attr-tdth-headers)
attributes:

``` html
<table>
 <caption>Characteristics with positive and negative sides</caption>
 <thead>
  <tr>
   <th> Characteristic
   <th> Negative
   <th> Positive
 <tbody>
  <tr>
   <th> Mood
   <td> Sad
   <td> Happy
  <tr>
   <th> Grade
   <td> Failing
   <td> Passing
</table>
```
:::

##### [4.9.1.2]{.secno} Techniques for table design[](#table-layout-techniques){.self-link} {#table-layout-techniques}

Good table design is key to making tables more readable and usable.

In visual media, providing column and row borders and alternating row
backgrounds can be very effective to make complicated tables more
readable.

For tables with large volumes of numeric content, using monospaced fonts
can help users see patterns, especially in situations where a user agent
does not render the borders. (Unfortunately, for historical reasons, not
rendering borders on tables is a common default.)

In speech media, table cells can be distinguished by reporting the
corresponding headers before reading the cell\'s contents, and by
allowing users to navigate the table in a grid fashion, rather than
serializing the entire contents of the table in source order.

Authors are encouraged to use CSS to achieve these effects.

User agents are encouraged to render tables using these techniques
whenever the page does not use CSS and the table is not classified as a
layout table.

#### [4.9.2]{.secno} The [`caption`]{.dfn dfn-type="element"} element[](#the-caption-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/caption](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption "The <caption> HTML element specifies the caption (or title) of a table.")

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
[HTMLTableCaptionElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableCaptionElement "The HTMLTableCaptionElement interface provides special properties (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating table <caption> elements.")

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

[Categories](dom.html#concept-element-categories){#the-caption-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-caption-element:concept-element-contexts}:
:   As the first element child of a
    [`table`{#the-caption-element:the-table-element}](#the-table-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-caption-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-caption-element:flow-content-2},
    but with no descendant
    [`table`{#the-caption-element:the-table-element-2}](#the-table-element)
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-caption-element:concept-element-tag-omission}:
:   A
    [`caption`{#the-caption-element:the-caption-element}](#the-caption-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-caption-element:syntax-end-tag}
    can be omitted if the
    [`caption`{#the-caption-element:the-caption-element-2}](#the-caption-element)
    element is not immediately followed by [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-caption-element:space-characters
    x-internal="space-characters"} or a
    [comment](syntax.html#syntax-comments){#the-caption-element:syntax-comments}.

[Content attributes](dom.html#concept-element-attributes){#the-caption-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-caption-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-caption-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-caption).
:   [For implementers](https://w3c.github.io/html-aam/#el-caption).

[DOM interface](dom.html#concept-element-dom){#the-caption-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTableCaptionElement : HTMLElement {
      [HTMLConstructor] constructor();

      // also has obsolete members
    };
    ```

The
[`caption`{#the-caption-element:the-caption-element-3}](#the-caption-element)
element
[represents](dom.html#represents){#the-caption-element:represents} the
title of the
[`table`{#the-caption-element:the-table-element-3}](#the-table-element)
that is its parent, if it has a parent and that is a
[`table`{#the-caption-element:the-table-element-4}](#the-table-element)
element.

The
[`caption`{#the-caption-element:the-caption-element-4}](#the-caption-element)
element takes part in the [table
model](#table-model){#the-caption-element:table-model}.

When a
[`table`{#the-caption-element:the-table-element-5}](#the-table-element)
element is the only content in a
[`figure`{#the-caption-element:the-figure-element}](grouping-content.html#the-figure-element)
element other than the
[`figcaption`{#the-caption-element:the-figcaption-element}](grouping-content.html#the-figcaption-element),
the
[`caption`{#the-caption-element:the-caption-element-5}](#the-caption-element)
element should be omitted in favor of the
[`figcaption`{#the-caption-element:the-figcaption-element-2}](grouping-content.html#the-figcaption-element).

A caption can introduce context for a table, making it significantly
easier to understand.

::: example
Consider, for instance, the following table:

1

2

3

4

5

6

1

2

3

4

5

6

7

2

3

4

5

6

7

8

3

4

5

6

7

8

9

4

5

6

7

8

9

10

5

6

7

8

9

10

11

6

7

8

9

10

11

12

In the abstract, this table is not clear. However, with a caption giving
the table\'s number (for
[reference](dom.html#referenced){#the-caption-element:referenced} in the
main prose) and explaining its use, it makes more sense:

``` html
<caption>
<p>Table 1.
<p>This table shows the total score obtained from rolling two
six-sided dice. The first row represents the value of the first die,
the first column the value of the second die. The total is given in
the cell that corresponds to the values of the two dice.
</caption>
```

This provides the user with more context:

Table 1.

This table shows the total score obtained from rolling two six-sided
dice. The first row represents the value of the first die, the first
column the value of the second die. The total is given in the cell that
corresponds to the values of the two dice.

1

2

3

4

5

6

1

2

3

4

5

6

7

2

3

4

5

6

7

8

3

4

5

6

7

8

9

4

5

6

7

8

9

10

5

6

7

8

9

10

11

6

7

8

9

10

11

12
:::

#### [4.9.3]{.secno} The [`colgroup`]{.dfn dfn-type="element"} element[](#the-colgroup-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/colgroup](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup "The <colgroup> HTML element defines a group of columns within a table.")

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
[HTMLTableColElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableColElement "The HTMLTableColElement interface provides properties for manipulating single or grouped table column elements.")

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

[Categories](dom.html#concept-element-categories){#the-colgroup-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-colgroup-element:concept-element-contexts}:
:   As a child of a
    [`table`{#the-colgroup-element:the-table-element}](#the-table-element)
    element, after any
    [`caption`{#the-colgroup-element:the-caption-element}](#the-caption-element)
    elements and before any
    [`thead`{#the-colgroup-element:the-thead-element}](#the-thead-element),
    [`tbody`{#the-colgroup-element:the-tbody-element}](#the-tbody-element),
    [`tfoot`{#the-colgroup-element:the-tfoot-element}](#the-tfoot-element),
    and [`tr`{#the-colgroup-element:the-tr-element}](#the-tr-element)
    elements.

[Content model](dom.html#concept-element-content-model){#the-colgroup-element:concept-element-content-model}:
:   If the
    [`span`{#the-colgroup-element:attr-colgroup-span}](#attr-colgroup-span)
    attribute is present:
    [Nothing](dom.html#concept-content-nothing){#the-colgroup-element:concept-content-nothing}.
:   If the
    [`span`{#the-colgroup-element:attr-colgroup-span-2}](#attr-colgroup-span)
    attribute is absent: Zero or more
    [`col`{#the-colgroup-element:the-col-element}](#the-col-element) and
    [`template`{#the-colgroup-element:the-template-element}](scripting.html#the-template-element)
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-colgroup-element:concept-element-tag-omission}:
:   A
    [`colgroup`{#the-colgroup-element:the-colgroup-element}](#the-colgroup-element)
    element\'s [start
    tag](syntax.html#syntax-start-tag){#the-colgroup-element:syntax-start-tag}
    can be omitted if the first thing inside the
    [`colgroup`{#the-colgroup-element:the-colgroup-element-2}](#the-colgroup-element)
    element is a
    [`col`{#the-colgroup-element:the-col-element-2}](#the-col-element)
    element, and if the element is not immediately preceded by another
    [`colgroup`{#the-colgroup-element:the-colgroup-element-3}](#the-colgroup-element)
    element whose [end
    tag](syntax.html#syntax-end-tag){#the-colgroup-element:syntax-end-tag}
    has been omitted. (It can\'t be omitted if the element is empty.)
:   A
    [`colgroup`{#the-colgroup-element:the-colgroup-element-4}](#the-colgroup-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-colgroup-element:syntax-end-tag-2}
    can be omitted if the
    [`colgroup`{#the-colgroup-element:the-colgroup-element-5}](#the-colgroup-element)
    element is not immediately followed by [ASCII
    whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#the-colgroup-element:space-characters
    x-internal="space-characters"} or a
    [comment](syntax.html#syntax-comments){#the-colgroup-element:syntax-comments}.

[Content attributes](dom.html#concept-element-attributes){#the-colgroup-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-colgroup-element:global-attributes}
:   [`span`{#the-colgroup-element:attr-colgroup-span-3}](#attr-colgroup-span)
    --- Number of columns spanned by the element

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-colgroup-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-colgroup).
:   [For implementers](https://w3c.github.io/html-aam/#el-colgroup).

[DOM interface](dom.html#concept-element-dom){#the-colgroup-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTableColElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute unsigned long span;

      // also has obsolete members
    };
    ```

The
[`colgroup`{#the-colgroup-element:the-colgroup-element-6}](#the-colgroup-element)
element
[represents](dom.html#represents){#the-colgroup-element:represents} a
[group](#concept-column-group){#the-colgroup-element:concept-column-group}
of one or more
[columns](#concept-column){#the-colgroup-element:concept-column} in the
[`table`{#the-colgroup-element:the-table-element-2}](#the-table-element)
that is its parent, if it has a parent and that is a
[`table`{#the-colgroup-element:the-table-element-3}](#the-table-element)
element.

If the
[`colgroup`{#the-colgroup-element:the-colgroup-element-7}](#the-colgroup-element)
element contains no
[`col`{#the-colgroup-element:the-col-element-3}](#the-col-element)
elements, then the element may have a [`span`]{#attr-colgroup-span .dfn
dfn-for="colgroup" dfn-type="element-attr"} content attribute specified,
whose value must be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#the-colgroup-element:valid-non-negative-integer}
greater than zero and less than or equal to 1000.

The
[`colgroup`{#the-colgroup-element:the-colgroup-element-8}](#the-colgroup-element)
element and its
[`span`{#the-colgroup-element:attr-colgroup-span-4}](#attr-colgroup-span)
attribute take part in the [table
model](#table-model){#the-colgroup-element:table-model}.

The [`span`]{#dom-colgroup-span .dfn dfn-for="HTMLTableColElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-colgroup-element:reflect}
the content attribute of the same name. It is [clamped to the
range](common-dom-interfaces.html#clamped-to-the-range){#the-colgroup-element:clamped-to-the-range}
\[1, 1000\], and its [default
value](common-dom-interfaces.html#default-value){#the-colgroup-element:default-value}
is 1.

#### [4.9.4]{.secno} The [`col`]{.dfn dfn-type="element"} element[](#the-col-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/col](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col "The <col> HTML element defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.")

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

[Categories](dom.html#concept-element-categories){#the-col-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-col-element:concept-element-contexts}:
:   As a child of a
    [`colgroup`{#the-col-element:the-colgroup-element}](#the-colgroup-element)
    element that doesn\'t have a
    [`span`{#the-col-element:attr-col-span}](#attr-col-span) attribute.

[Content model](dom.html#concept-element-content-model){#the-col-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-col-element:concept-content-nothing}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-col-element:concept-element-tag-omission}:
:   No [end
    tag](syntax.html#syntax-end-tag){#the-col-element:syntax-end-tag}.

[Content attributes](dom.html#concept-element-attributes){#the-col-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-col-element:global-attributes}
:   [`span`{#the-col-element:attr-col-span-2}](#attr-col-span) ---
    Number of columns spanned by the element

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-col-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-col).
:   [For implementers](https://w3c.github.io/html-aam/#el-col).

[DOM interface](dom.html#concept-element-dom){#the-col-element:concept-element-dom}:
:   Uses
    [`HTMLTableColElement`{#the-col-element:htmltablecolelement}](#htmltablecolelement),
    as defined for
    [`colgroup`{#the-col-element:the-colgroup-element-2}](#the-colgroup-element)
    elements.

If a [`col`{#the-col-element:the-col-element}](#the-col-element) element
has a parent and that is a
[`colgroup`{#the-col-element:the-colgroup-element-3}](#the-colgroup-element)
element that itself has a parent that is a
[`table`{#the-col-element:the-table-element}](#the-table-element)
element, then the
[`col`{#the-col-element:the-col-element-2}](#the-col-element) element
[represents](dom.html#represents){#the-col-element:represents} one or
more [columns](#concept-column){#the-col-element:concept-column} in the
[column
group](#concept-column-group){#the-col-element:concept-column-group}
represented by that
[`colgroup`{#the-col-element:the-colgroup-element-4}](#the-colgroup-element).

The element may have a [`span`]{#attr-col-span .dfn dfn-for="col"
dfn-type="element-attr"} content attribute specified, whose value must
be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#the-col-element:valid-non-negative-integer}
greater than zero and less than or equal to 1000.

The [`col`{#the-col-element:the-col-element-3}](#the-col-element)
element and its
[`span`{#the-col-element:attr-col-span-3}](#attr-col-span) attribute
take part in the [table
model](#table-model){#the-col-element:table-model}.

#### [4.9.5]{.secno} The [`tbody`]{.dfn dfn-type="element"} element[](#the-tbody-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/tbody](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody "The <tbody> HTML element encapsulates a set of table rows (<tr> elements), indicating that they comprise the body of the table (<table>).")

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
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableSectionElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableSectionElement "The HTMLTableSectionElement interface provides special properties and methods (beyond the HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of sections, that is headers, footers and bodies (<thead>, <tfoot>, and <tbody>, respectively) in an HTML table.")

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

[Categories](dom.html#concept-element-categories){#the-tbody-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-tbody-element:concept-element-contexts}:
:   As a child of a
    [`table`{#the-tbody-element:the-table-element}](#the-table-element)
    element, after any
    [`caption`{#the-tbody-element:the-caption-element}](#the-caption-element),
    [`colgroup`{#the-tbody-element:the-colgroup-element}](#the-colgroup-element),
    and
    [`thead`{#the-tbody-element:the-thead-element}](#the-thead-element)
    elements, but only if there are no
    [`tr`{#the-tbody-element:the-tr-element}](#the-tr-element) elements
    that are children of the
    [`table`{#the-tbody-element:the-table-element-2}](#the-table-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-tbody-element:concept-element-content-model}:
:   Zero or more
    [`tr`{#the-tbody-element:the-tr-element-2}](#the-tr-element) and
    [script-supporting](dom.html#script-supporting-elements-2){#the-tbody-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-tbody-element:concept-element-tag-omission}:
:   A
    [`tbody`{#the-tbody-element:the-tbody-element}](#the-tbody-element)
    element\'s [start
    tag](syntax.html#syntax-start-tag){#the-tbody-element:syntax-start-tag}
    can be omitted if the first thing inside the
    [`tbody`{#the-tbody-element:the-tbody-element-2}](#the-tbody-element)
    element is a
    [`tr`{#the-tbody-element:the-tr-element-3}](#the-tr-element)
    element, and if the element is not immediately preceded by a
    [`tbody`{#the-tbody-element:the-tbody-element-3}](#the-tbody-element),
    [`thead`{#the-tbody-element:the-thead-element-2}](#the-thead-element),
    or
    [`tfoot`{#the-tbody-element:the-tfoot-element}](#the-tfoot-element)
    element whose [end
    tag](syntax.html#syntax-end-tag){#the-tbody-element:syntax-end-tag}
    has been omitted. (It can\'t be omitted if the element is empty.)
:   A
    [`tbody`{#the-tbody-element:the-tbody-element-4}](#the-tbody-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-tbody-element:syntax-end-tag-2}
    can be omitted if the
    [`tbody`{#the-tbody-element:the-tbody-element-5}](#the-tbody-element)
    element is immediately followed by a
    [`tbody`{#the-tbody-element:the-tbody-element-6}](#the-tbody-element)
    or
    [`tfoot`{#the-tbody-element:the-tfoot-element-2}](#the-tfoot-element)
    element, or if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-tbody-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-tbody-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-tbody-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-tbody).
:   [For implementers](https://w3c.github.io/html-aam/#el-tbody).

[DOM interface](dom.html#concept-element-dom){#the-tbody-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTableSectionElement : HTMLElement {
      [HTMLConstructor] constructor();

      [SameObject] readonly attribute HTMLCollection rows;
      HTMLTableRowElement insertRow(optional long index = -1);
      [CEReactions] undefined deleteRow(long index);

      // also has obsolete members
    };
    ```

    The
    [`HTMLTableSectionElement`{#the-tbody-element:htmltablesectionelement}](#htmltablesectionelement)
    interface is also used for
    [`thead`{#the-tbody-element:the-thead-element-3}](#the-thead-element)
    and
    [`tfoot`{#the-tbody-element:the-tfoot-element-3}](#the-tfoot-element)
    elements.

The
[`tbody`{#the-tbody-element:the-tbody-element-7}](#the-tbody-element)
element [represents](dom.html#represents){#the-tbody-element:represents}
a [block](#concept-row-group){#the-tbody-element:concept-row-group} of
[rows](#concept-row){#the-tbody-element:concept-row} that consist of a
body of data for the parent
[`table`{#the-tbody-element:the-table-element-3}](#the-table-element)
element, if the
[`tbody`{#the-tbody-element:the-tbody-element-8}](#the-tbody-element)
element has a parent and it is a
[`table`{#the-tbody-element:the-table-element-4}](#the-table-element).

The
[`tbody`{#the-tbody-element:the-tbody-element-9}](#the-tbody-element)
element takes part in the [table
model](#table-model){#the-tbody-element:table-model}.

`tbody`{.variable}`.`[`rows`](#dom-tbody-rows){#dom-tbody-rows-dev}

:   Returns an
    [`HTMLCollection`{#the-tbody-element:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
    of the [`tr`{#the-tbody-element:the-tr-element-4}](#the-tr-element)
    elements of the table section.

`tr`{.variable}` = ``tbody`{.variable}`.`[`insertRow`](#dom-tbody-insertrow){#dom-tbody-insertrow-dev}`([ ``index`{.variable}` ])`

:   Creates a
    [`tr`{#the-tbody-element:the-tr-element-5}](#the-tr-element)
    element, inserts it into the table section at the position given by
    the argument, and returns the
    [`tr`{#the-tbody-element:the-tr-element-6}](#the-tr-element).

    The position is relative to the rows in the table section. The index
    −1, which is the default if the argument is omitted, is equivalent
    to inserting at the end of the table section.

    If the given position is less than −1 or greater than the number of
    rows, throws an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tbody-element:indexsizeerror
    x-internal="indexsizeerror"}
    [`DOMException`{#the-tbody-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

`tbody`{.variable}`.`[`deleteRow`](#dom-tbody-deleterow){#dom-tbody-deleterow-dev}`(``index`{.variable}`)`

:   Removes the
    [`tr`{#the-tbody-element:the-tr-element-7}](#the-tr-element) element
    with the given position in the table section.

    The position is relative to the rows in the table section. The index
    −1 is equivalent to deleting the last row of the table section.

    If the given position is less than −1 or greater than the index of
    the last row, or if there are no rows, throws an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tbody-element:indexsizeerror-2
    x-internal="indexsizeerror"}
    [`DOMException`{#the-tbody-element:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

The [`rows`]{#dom-tbody-rows .dfn dfn-for="HTMLTableSectionElement"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#the-tbody-element:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at this element, whose filter matches only
[`tr`{#the-tbody-element:the-tr-element-8}](#the-tr-element) elements
that are children of this element.

The [`insertRow(``index`{.variable}`)`]{#dom-tbody-insertrow .dfn
dfn-for="HTMLTableSectionElement" dfn-type="method"} method must act as
follows:

1.  If `index`{.variable} is less than −1 or greater than the number of
    elements in the
    [`rows`{#the-tbody-element:dom-tbody-rows-2}](#dom-tbody-rows)
    collection, throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tbody-element:indexsizeerror-3
    x-internal="indexsizeerror"}
    [`DOMException`{#the-tbody-element:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `table row`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-tbody-element:create-an-element
    x-internal="create-an-element"} given this element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-tbody-element:node-document
    x-internal="node-document"}, \"`tr`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-tbody-element:html-namespace-2
    x-internal="html-namespace-2"}.

3.  If `index`{.variable} is −1 or equal to the number of items in the
    [`rows`{#the-tbody-element:dom-tbody-rows-3}](#dom-tbody-rows)
    collection, then
    [append](https://dom.spec.whatwg.org/#concept-node-append){#the-tbody-element:concept-node-append
    x-internal="concept-node-append"} `table row`{.variable} to this
    element.

4.  Otherwise,
    [insert](https://dom.spec.whatwg.org/#concept-node-insert){#the-tbody-element:concept-node-insert
    x-internal="concept-node-insert"} `table row`{.variable} as a child
    of this element, immediately before the `index`{.variable}th
    [`tr`{#the-tbody-element:the-tr-element-9}](#the-tr-element) element
    in the
    [`rows`{#the-tbody-element:dom-tbody-rows-4}](#dom-tbody-rows)
    collection.

5.  Return `table row`{.variable}.

The [`deleteRow(``index`{.variable}`)`]{#dom-tbody-deleterow .dfn
dfn-for="HTMLTableSectionElement" dfn-type="method"} method must, when
invoked, act as follows:

1.  If `index`{.variable} is less than −1 or greater than or equal to
    the number of elements in the
    [`rows`{#the-tbody-element:dom-tbody-rows-5}](#dom-tbody-rows)
    collection, then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tbody-element:indexsizeerror-4
    x-internal="indexsizeerror"}
    [`DOMException`{#the-tbody-element:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If `index`{.variable} is −1, then
    [remove](https://dom.spec.whatwg.org/#concept-node-remove){#the-tbody-element:concept-node-remove
    x-internal="concept-node-remove"} the last element in the
    [`rows`{#the-tbody-element:dom-tbody-rows-6}](#dom-tbody-rows)
    collection from this element, or do nothing if the
    [`rows`{#the-tbody-element:dom-tbody-rows-7}](#dom-tbody-rows)
    collection is empty.

3.  Otherwise,
    [remove](https://dom.spec.whatwg.org/#concept-node-remove){#the-tbody-element:concept-node-remove-2
    x-internal="concept-node-remove"} the `index`{.variable}th element
    in the
    [`rows`{#the-tbody-element:dom-tbody-rows-8}](#dom-tbody-rows)
    collection from this element.

#### [4.9.6]{.secno} The [`thead`]{.dfn dfn-type="element"} element[](#the-thead-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/thead](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead "The <thead> HTML element defines a set of rows defining the head of the columns of the table.")

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
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-thead-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-thead-element:concept-element-contexts}:
:   As a child of a
    [`table`{#the-thead-element:the-table-element}](#the-table-element)
    element, after any
    [`caption`{#the-thead-element:the-caption-element}](#the-caption-element),
    and
    [`colgroup`{#the-thead-element:the-colgroup-element}](#the-colgroup-element)
    elements and before any
    [`tbody`{#the-thead-element:the-tbody-element}](#the-tbody-element),
    [`tfoot`{#the-thead-element:the-tfoot-element}](#the-tfoot-element),
    and [`tr`{#the-thead-element:the-tr-element}](#the-tr-element)
    elements, but only if there are no other
    [`thead`{#the-thead-element:the-thead-element}](#the-thead-element)
    elements that are children of the
    [`table`{#the-thead-element:the-table-element-2}](#the-table-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-thead-element:concept-element-content-model}:
:   Zero or more
    [`tr`{#the-thead-element:the-tr-element-2}](#the-tr-element) and
    [script-supporting](dom.html#script-supporting-elements-2){#the-thead-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-thead-element:concept-element-tag-omission}:
:   A
    [`thead`{#the-thead-element:the-thead-element-2}](#the-thead-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-thead-element:syntax-end-tag}
    can be omitted if the
    [`thead`{#the-thead-element:the-thead-element-3}](#the-thead-element)
    element is immediately followed by a
    [`tbody`{#the-thead-element:the-tbody-element-2}](#the-tbody-element)
    or
    [`tfoot`{#the-thead-element:the-tfoot-element-2}](#the-tfoot-element)
    element.

[Content attributes](dom.html#concept-element-attributes){#the-thead-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-thead-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-thead-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-thead).
:   [For implementers](https://w3c.github.io/html-aam/#el-thead).

[DOM interface](dom.html#concept-element-dom){#the-thead-element:concept-element-dom}:
:   Uses
    [`HTMLTableSectionElement`{#the-thead-element:htmltablesectionelement}](#htmltablesectionelement),
    as defined for
    [`tbody`{#the-thead-element:the-tbody-element-3}](#the-tbody-element)
    elements.

The
[`thead`{#the-thead-element:the-thead-element-4}](#the-thead-element)
element [represents](dom.html#represents){#the-thead-element:represents}
the [block](#concept-row-group){#the-thead-element:concept-row-group} of
[rows](#concept-row){#the-thead-element:concept-row} that consist of the
column labels (headers) and any ancillary non-header cells for the
parent
[`table`{#the-thead-element:the-table-element-3}](#the-table-element)
element, if the
[`thead`{#the-thead-element:the-thead-element-5}](#the-thead-element)
element has a parent and it is a
[`table`{#the-thead-element:the-table-element-4}](#the-table-element).

The
[`thead`{#the-thead-element:the-thead-element-6}](#the-thead-element)
element takes part in the [table
model](#table-model){#the-thead-element:table-model}.

::: example
This example shows a
[`thead`{#the-thead-element:the-thead-element-7}](#the-thead-element)
element being used. Notice the use of both
[`th`{#the-thead-element:the-th-element}](#the-th-element) and
[`td`{#the-thead-element:the-td-element}](#the-td-element) elements in
the
[`thead`{#the-thead-element:the-thead-element-8}](#the-thead-element)
element: the first row is the headers, and the second row is an
explanation of how to fill in the table.

``` html
<table>
 <caption> School auction sign-up sheet </caption>
 <thead>
  <tr>
   <th><label for=e1>Name</label>
   <th><label for=e2>Product</label>
   <th><label for=e3>Picture</label>
   <th><label for=e4>Price</label>
  <tr>
   <td>Your name here
   <td>What are you selling?
   <td>Link to a picture
   <td>Your reserve price
 <tbody>
  <tr>
   <td>Ms Danus
   <td>Doughnuts
   <td><img src="https://example.com/mydoughnuts.png" title="Doughnuts from Ms Danus">
   <td>$45
  <tr>
   <td><input id=e1 type=text name=who required form=f>
   <td><input id=e2 type=text name=what required form=f>
   <td><input id=e3 type=url name=pic form=f>
   <td><input id=e4 type=number step=0.01 min=0 value=0 required form=f>
</table>
<form id=f action="/auction.cgi">
 <input type=button name=add value="Submit">
</form>
```
:::

#### [4.9.7]{.secno} The [`tfoot`]{.dfn dfn-type="element"} element[](#the-tfoot-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/tfoot](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot "The <tfoot> HTML element defines a set of rows summarizing the columns of the table.")

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
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-tfoot-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-tfoot-element:concept-element-contexts}:
:   As a child of a
    [`table`{#the-tfoot-element:the-table-element}](#the-table-element)
    element, after any
    [`caption`{#the-tfoot-element:the-caption-element}](#the-caption-element),
    [`colgroup`{#the-tfoot-element:the-colgroup-element}](#the-colgroup-element),
    [`thead`{#the-tfoot-element:the-thead-element}](#the-thead-element),
    [`tbody`{#the-tfoot-element:the-tbody-element}](#the-tbody-element),
    and [`tr`{#the-tfoot-element:the-tr-element}](#the-tr-element)
    elements, but only if there are no other
    [`tfoot`{#the-tfoot-element:the-tfoot-element}](#the-tfoot-element)
    elements that are children of the
    [`table`{#the-tfoot-element:the-table-element-2}](#the-table-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-tfoot-element:concept-element-content-model}:
:   Zero or more
    [`tr`{#the-tfoot-element:the-tr-element-2}](#the-tr-element) and
    [script-supporting](dom.html#script-supporting-elements-2){#the-tfoot-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-tfoot-element:concept-element-tag-omission}:
:   A
    [`tfoot`{#the-tfoot-element:the-tfoot-element-2}](#the-tfoot-element)
    element\'s [end
    tag](syntax.html#syntax-end-tag){#the-tfoot-element:syntax-end-tag}
    can be omitted if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-tfoot-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-tfoot-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-tfoot-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-tfoot).
:   [For implementers](https://w3c.github.io/html-aam/#el-tfoot).

[DOM interface](dom.html#concept-element-dom){#the-tfoot-element:concept-element-dom}:
:   Uses
    [`HTMLTableSectionElement`{#the-tfoot-element:htmltablesectionelement}](#htmltablesectionelement),
    as defined for
    [`tbody`{#the-tfoot-element:the-tbody-element-2}](#the-tbody-element)
    elements.

The
[`tfoot`{#the-tfoot-element:the-tfoot-element-3}](#the-tfoot-element)
element [represents](dom.html#represents){#the-tfoot-element:represents}
the [block](#concept-row-group){#the-tfoot-element:concept-row-group} of
[rows](#concept-row){#the-tfoot-element:concept-row} that consist of the
column summaries (footers) for the parent
[`table`{#the-tfoot-element:the-table-element-3}](#the-table-element)
element, if the
[`tfoot`{#the-tfoot-element:the-tfoot-element-4}](#the-tfoot-element)
element has a parent and it is a
[`table`{#the-tfoot-element:the-table-element-4}](#the-table-element).

The
[`tfoot`{#the-tfoot-element:the-tfoot-element-5}](#the-tfoot-element)
element takes part in the [table
model](#table-model){#the-tfoot-element:table-model}.

#### [4.9.8]{.secno} The [`tr`]{.dfn dfn-type="element"} element[](#the-tr-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/tr](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr "The <tr> HTML element defines a row of cells in a table. The row's cells can then be established using a mix of <td> (data cell) and <th> (header cell) elements.")

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
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableRowElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableRowElement "The HTMLTableRowElement interface provides special properties and methods (beyond the HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of rows in an HTML table.")

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

[Categories](dom.html#concept-element-categories){#the-tr-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-tr-element:concept-element-contexts}:
:   As a child of a
    [`thead`{#the-tr-element:the-thead-element}](#the-thead-element)
    element.
:   As a child of a
    [`tbody`{#the-tr-element:the-tbody-element}](#the-tbody-element)
    element.
:   As a child of a
    [`tfoot`{#the-tr-element:the-tfoot-element}](#the-tfoot-element)
    element.
:   As a child of a
    [`table`{#the-tr-element:the-table-element}](#the-table-element)
    element, after any
    [`caption`{#the-tr-element:the-caption-element}](#the-caption-element),
    [`colgroup`{#the-tr-element:the-colgroup-element}](#the-colgroup-element),
    and
    [`thead`{#the-tr-element:the-thead-element-2}](#the-thead-element)
    elements, but only if there are no
    [`tbody`{#the-tr-element:the-tbody-element-2}](#the-tbody-element)
    elements that are children of the
    [`table`{#the-tr-element:the-table-element-2}](#the-table-element)
    element.

[Content model](dom.html#concept-element-content-model){#the-tr-element:concept-element-content-model}:
:   Zero or more
    [`td`{#the-tr-element:the-td-element}](#the-td-element),
    [`th`{#the-tr-element:the-th-element}](#the-th-element), and
    [script-supporting](dom.html#script-supporting-elements-2){#the-tr-element:script-supporting-elements-2}
    elements.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-tr-element:concept-element-tag-omission}:
:   A [`tr`{#the-tr-element:the-tr-element}](#the-tr-element) element\'s
    [end
    tag](syntax.html#syntax-end-tag){#the-tr-element:syntax-end-tag} can
    be omitted if the
    [`tr`{#the-tr-element:the-tr-element-2}](#the-tr-element) element is
    immediately followed by another
    [`tr`{#the-tr-element:the-tr-element-3}](#the-tr-element) element,
    or if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-tr-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-tr-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-tr-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-tr).
:   [For implementers](https://w3c.github.io/html-aam/#el-tr).

[DOM interface](dom.html#concept-element-dom){#the-tr-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTableRowElement : HTMLElement {
      [HTMLConstructor] constructor();

      readonly attribute long rowIndex;
      readonly attribute long sectionRowIndex;
      [SameObject] readonly attribute HTMLCollection cells;
      HTMLTableCellElement insertCell(optional long index = -1);
      [CEReactions] undefined deleteCell(long index);

      // also has obsolete members
    };
    ```

The [`tr`{#the-tr-element:the-tr-element-4}](#the-tr-element) element
[represents](dom.html#represents){#the-tr-element:represents} a
[row](#concept-row){#the-tr-element:concept-row} of
[cells](#concept-cell){#the-tr-element:concept-cell} in a
[table](#concept-table){#the-tr-element:concept-table}.

The [`tr`{#the-tr-element:the-tr-element-5}](#the-tr-element) element
takes part in the [table
model](#table-model){#the-tr-element:table-model}.

`tr`{.variable}`.`[`rowIndex`](#dom-tr-rowindex){#dom-tr-rowindex-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableRowElement/rowIndex](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableRowElement/rowIndex "The HTMLTableRowElement.rowIndex read-only property represents the position of a row in relation to the whole <table>.")

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

Returns the position of the row in the table\'s
[`rows`{#the-tr-element:dom-table-rows}](#dom-table-rows) list.

Returns −1 if the element isn\'t in a table.

`tr`{.variable}`.`[`sectionRowIndex`](#dom-tr-sectionrowindex){#dom-tr-sectionrowindex-dev}

Returns the position of the row in the table section\'s
[`rows`{#the-tr-element:dom-tbody-rows}](#dom-tbody-rows) list.

Returns −1 if the element isn\'t in a table section.

`tr`{.variable}`.`[`cells`](#dom-tr-cells){#dom-tr-cells-dev}

Returns an
[`HTMLCollection`{#the-tr-element:htmlcollection-2}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
of the [`td`{#the-tr-element:the-td-element-2}](#the-td-element) and
[`th`{#the-tr-element:the-th-element-2}](#the-th-element) elements of
the row.

`cell`{.variable}` = ``tr`{.variable}`.`[`insertCell`](#dom-tr-insertcell){#dom-tr-insertcell-dev}`([ ``index`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableRowElement/insertCell](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableRowElement/insertCell "The HTMLTableRowElement.insertCell() method inserts a new cell (<td>) into a table row (<tr>) and returns a reference to the cell.")

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

Creates a [`td`{#the-tr-element:the-td-element-3}](#the-td-element)
element, inserts it into the table row at the position given by the
argument, and returns the
[`td`{#the-tr-element:the-td-element-4}](#the-td-element).

The position is relative to the cells in the row. The index −1, which is
the default if the argument is omitted, is equivalent to inserting at
the end of the row.

If the given position is less than −1 or greater than the number of
cells, throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tr-element:indexsizeerror
x-internal="indexsizeerror"}
[`DOMException`{#the-tr-element:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

`tr`{.variable}`.`[`deleteCell`](#dom-tr-deletecell){#dom-tr-deletecell-dev}`(``index`{.variable}`)`

Removes the [`td`{#the-tr-element:the-td-element-5}](#the-td-element) or
[`th`{#the-tr-element:the-th-element-3}](#the-th-element) element with
the given position in the row.

The position is relative to the cells in the row. The index −1 is
equivalent to deleting the last cell of the row.

If the given position is less than −1 or greater than the index of the
last cell, or if there are no cells, throws an
[\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tr-element:indexsizeerror-2
x-internal="indexsizeerror"}
[`DOMException`{#the-tr-element:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

The [`rowIndex`]{#dom-tr-rowindex .dfn dfn-for="HTMLTableRowElement"
dfn-type="attribute"} attribute must, if this element has a parent
[`table`{#the-tr-element:the-table-element-3}](#the-table-element)
element, or a parent
[`tbody`{#the-tr-element:the-tbody-element-3}](#the-tbody-element),
[`thead`{#the-tr-element:the-thead-element-3}](#the-thead-element), or
[`tfoot`{#the-tr-element:the-tfoot-element-2}](#the-tfoot-element)
element and a *grandparent*
[`table`{#the-tr-element:the-table-element-4}](#the-table-element)
element, return the index of this
[`tr`{#the-tr-element:the-tr-element-6}](#the-tr-element) element in
that [`table`{#the-tr-element:the-table-element-5}](#the-table-element)
element\'s [`rows`{#the-tr-element:dom-table-rows-2}](#dom-table-rows)
collection. If there is no such
[`table`{#the-tr-element:the-table-element-6}](#the-table-element)
element, then the attribute must return −1.

The [`sectionRowIndex`]{#dom-tr-sectionrowindex .dfn
dfn-for="HTMLTableRowElement" dfn-type="attribute"} attribute must, if
this element has a parent
[`table`{#the-tr-element:the-table-element-7}](#the-table-element),
[`tbody`{#the-tr-element:the-tbody-element-4}](#the-tbody-element),
[`thead`{#the-tr-element:the-thead-element-4}](#the-thead-element), or
[`tfoot`{#the-tr-element:the-tfoot-element-3}](#the-tfoot-element)
element, return the index of the
[`tr`{#the-tr-element:the-tr-element-7}](#the-tr-element) element in the
parent element\'s `rows` collection (for tables, that\'s
[`HTMLTableElement`{#the-tr-element:htmltableelement}](#htmltableelement)\'s
[`rows`{#the-tr-element:dom-table-rows-3}](#dom-table-rows) collection;
for table sections, that\'s
[`HTMLTableSectionElement`{#the-tr-element:htmltablesectionelement}](#htmltablesectionelement)\'s
[`rows`{#the-tr-element:dom-tbody-rows-2}](#dom-tbody-rows) collection).
If there is no such parent element, then the attribute must return −1.

The [`cells`]{#dom-tr-cells .dfn dfn-for="HTMLTableRowElement"
dfn-type="attribute"} attribute must return an
[`HTMLCollection`{#the-tr-element:htmlcollection-3}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
rooted at this [`tr`{#the-tr-element:the-tr-element-8}](#the-tr-element)
element, whose filter matches only
[`td`{#the-tr-element:the-td-element-6}](#the-td-element) and
[`th`{#the-tr-element:the-th-element-4}](#the-th-element) elements that
are children of the
[`tr`{#the-tr-element:the-tr-element-9}](#the-tr-element) element.

The [`insertCell(``index`{.variable}`)`]{#dom-tr-insertcell .dfn
dfn-for="HTMLTableRowElement" dfn-type="method"} method must act as
follows:

1.  If `index`{.variable} is less than −1 or greater than the number of
    elements in the
    [`cells`{#the-tr-element:dom-tr-cells-2}](#dom-tr-cells) collection,
    then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tr-element:indexsizeerror-3
    x-internal="indexsizeerror"}
    [`DOMException`{#the-tr-element:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `table cell`{.variable} be the result of [creating an
    element](https://dom.spec.whatwg.org/#concept-create-element){#the-tr-element:create-an-element
    x-internal="create-an-element"} given this
    [`tr`{#the-tr-element:the-tr-element-10}](#the-tr-element)
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-tr-element:node-document
    x-internal="node-document"}, \"`td`\", and the [HTML
    namespace](https://infra.spec.whatwg.org/#html-namespace){#the-tr-element:html-namespace-2
    x-internal="html-namespace-2"}.

3.  If `index`{.variable} is equal to −1 or equal to the number of items
    in [`cells`{#the-tr-element:dom-tr-cells-3}](#dom-tr-cells)
    collection, then
    [append](https://dom.spec.whatwg.org/#concept-node-append){#the-tr-element:concept-node-append
    x-internal="concept-node-append"} `table cell`{.variable} to this
    [`tr`{#the-tr-element:the-tr-element-11}](#the-tr-element) element.

4.  Otherwise,
    [insert](https://dom.spec.whatwg.org/#concept-node-insert){#the-tr-element:concept-node-insert
    x-internal="concept-node-insert"} `table cell`{.variable} as a child
    of this [`tr`{#the-tr-element:the-tr-element-12}](#the-tr-element)
    element, immediately before the `index`{.variable}th
    [`td`{#the-tr-element:the-td-element-7}](#the-td-element) or
    [`th`{#the-tr-element:the-th-element-5}](#the-th-element) element in
    the [`cells`{#the-tr-element:dom-tr-cells-4}](#dom-tr-cells)
    collection.

5.  Return `table cell`{.variable}.

The [`deleteCell(``index`{.variable}`)`]{#dom-tr-deletecell .dfn
dfn-for="HTMLTableRowElement" dfn-type="method"} method must act as
follows:

1.  If `index`{.variable} is less than −1 or greater than or equal to
    the number of elements in the
    [`cells`{#the-tr-element:dom-tr-cells-5}](#dom-tr-cells) collection,
    then throw an
    [\"`IndexSizeError`\"](https://webidl.spec.whatwg.org/#indexsizeerror){#the-tr-element:indexsizeerror-4
    x-internal="indexsizeerror"}
    [`DOMException`{#the-tr-element:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If `index`{.variable} is −1, then
    [remove](https://dom.spec.whatwg.org/#concept-node-remove){#the-tr-element:concept-node-remove
    x-internal="concept-node-remove"} the last element in the
    [`cells`{#the-tr-element:dom-tr-cells-6}](#dom-tr-cells) collection
    from its parent, or do nothing if the
    [`cells`{#the-tr-element:dom-tr-cells-7}](#dom-tr-cells) collection
    is empty.

3.  Otherwise,
    [remove](https://dom.spec.whatwg.org/#concept-node-remove){#the-tr-element:concept-node-remove-2
    x-internal="concept-node-remove"} the `index`{.variable}th element
    in the [`cells`{#the-tr-element:dom-tr-cells-8}](#dom-tr-cells)
    collection from its parent.

#### [4.9.9]{.secno} The [`td`]{.dfn dfn-type="element"} element[](#the-td-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/td](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td "The <td> HTML element defines a cell of a table that contains data. It participates in the table model.")

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
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTableCellElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTableCellElement "The HTMLTableCellElement interface provides special properties and methods (beyond the regular HTMLElement interface it also has available to it by inheritance) for manipulating the layout and presentation of table cells, either header cells (<th>)) or data cells (<td>), in an HTML document.")

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

[Categories](dom.html#concept-element-categories){#the-td-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-td-element:concept-element-contexts}:
:   As a child of a
    [`tr`{#the-td-element:the-tr-element}](#the-tr-element) element.

[Content model](dom.html#concept-element-content-model){#the-td-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-td-element:flow-content-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-td-element:concept-element-tag-omission}:
:   A [`td`{#the-td-element:the-td-element}](#the-td-element) element\'s
    [end
    tag](syntax.html#syntax-end-tag){#the-td-element:syntax-end-tag} can
    be omitted if the
    [`td`{#the-td-element:the-td-element-2}](#the-td-element) element is
    immediately followed by a
    [`td`{#the-td-element:the-td-element-3}](#the-td-element) or
    [`th`{#the-td-element:the-th-element}](#the-th-element) element, or
    if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-td-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-td-element:global-attributes}
:   [`colspan`{#the-td-element:attr-tdth-colspan}](#attr-tdth-colspan)
    --- Number of columns that the cell is to span
:   [`rowspan`{#the-td-element:attr-tdth-rowspan}](#attr-tdth-rowspan)
    --- Number of rows that the cell is to span
:   [`headers`{#the-td-element:attr-tdth-headers}](#attr-tdth-headers)
    --- The header cells for this cell

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-td-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-td).
:   [For implementers](https://w3c.github.io/html-aam/#el-td).

[DOM interface](dom.html#concept-element-dom){#the-td-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTableCellElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute unsigned long colSpan;
      [CEReactions] attribute unsigned long rowSpan;
      [CEReactions] attribute DOMString headers;
      readonly attribute long cellIndex;

      [CEReactions] attribute DOMString scope; // only conforming for th elements
      [CEReactions] attribute DOMString abbr;  // only conforming for th elements

      // also has obsolete members
    };
    ```

    The
    [`HTMLTableCellElement`{#the-td-element:htmltablecellelement}](#htmltablecellelement)
    interface is also used for
    [`th`{#the-td-element:the-th-element-2}](#the-th-element) elements.

The [`td`{#the-td-element:the-td-element-4}](#the-td-element) element
[represents](dom.html#represents){#the-td-element:represents} a data
[cell](#concept-cell){#the-td-element:concept-cell} in a table.

The [`td`{#the-td-element:the-td-element-5}](#the-td-element) element
and its
[`colspan`{#the-td-element:attr-tdth-colspan-2}](#attr-tdth-colspan),
[`rowspan`{#the-td-element:attr-tdth-rowspan-2}](#attr-tdth-rowspan),
and [`headers`{#the-td-element:attr-tdth-headers-2}](#attr-tdth-headers)
attributes take part in the [table
model](#table-model){#the-td-element:table-model}.

User agents, especially in non-visual environments or where displaying
the table as a 2D grid is impractical, may give the user context for the
cell when rendering the contents of a cell; for instance, giving its
position in the [table
model](#table-model){#the-td-element:table-model-2}, or listing the
cell\'s header cells (as determined by the [algorithm for assigning
header
cells](#algorithm-for-assigning-header-cells){#the-td-element:algorithm-for-assigning-header-cells}).
When a cell\'s header cells are being listed, user agents may use the
value of [`abbr`{#the-td-element:attr-th-abbr}](#attr-th-abbr)
attributes on those header cells, if any, instead of the contents of the
header cells themselves.

::: example
In this example, we see a snippet of a web application consisting of a
grid of editable cells (essentially a simple spreadsheet). One of the
cells has been configured to show the sum of the cells above it. Three
have been marked as headings, which use
[`th`{#the-td-element:the-th-element-3}](#the-th-element) elements
instead of [`td`{#the-td-element:the-td-element-6}](#the-td-element)
elements. A script would attach event handlers to these elements to
maintain the total.

``` html
<table>
 <tr>
  <th><input value="Name">
  <th><input value="Paid ($)">
 <tr>
  <td><input value="Jeff">
  <td><input value="14">
 <tr>
  <td><input value="Britta">
  <td><input value="9">
 <tr>
  <td><input value="Abed">
  <td><input value="25">
 <tr>
  <td><input value="Shirley">
  <td><input value="2">
 <tr>
  <td><input value="Annie">
  <td><input value="5">
 <tr>
  <td><input value="Troy">
  <td><input value="5">
 <tr>
  <td><input value="Pierce">
  <td><input value="1000">
 <tr>
  <th><input value="Total">
  <td><output value="1060">
</table>
```
:::

#### [4.9.10]{.secno} The [`th`]{.dfn dfn-type="element"} element[](#the-th-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/th](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th "The <th> HTML element defines a cell as the header of a group of table cells. The exact nature of this group is defined by the scope and headers attributes.")

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
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-th-element:concept-element-categories}:
:   None.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-th-element:concept-element-contexts}:
:   As a child of a
    [`tr`{#the-th-element:the-tr-element}](#the-tr-element) element.

[Content model](dom.html#concept-element-content-model){#the-th-element:concept-element-content-model}:
:   [Flow
    content](dom.html#flow-content-2){#the-th-element:flow-content-2},
    but with no
    [`header`{#the-th-element:the-header-element}](sections.html#the-header-element),
    [`footer`{#the-th-element:the-footer-element}](sections.html#the-footer-element),
    [sectioning
    content](dom.html#sectioning-content-2){#the-th-element:sectioning-content-2},
    or [heading
    content](dom.html#heading-content-2){#the-th-element:heading-content-2}
    descendants.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-th-element:concept-element-tag-omission}:
:   A [`th`{#the-th-element:the-th-element}](#the-th-element) element\'s
    [end
    tag](syntax.html#syntax-end-tag){#the-th-element:syntax-end-tag} can
    be omitted if the
    [`th`{#the-th-element:the-th-element-2}](#the-th-element) element is
    immediately followed by a
    [`td`{#the-th-element:the-td-element}](#the-td-element) or
    [`th`{#the-th-element:the-th-element-3}](#the-th-element) element,
    or if there is no more content in the parent element.

[Content attributes](dom.html#concept-element-attributes){#the-th-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-th-element:global-attributes}
:   [`colspan`{#the-th-element:attr-tdth-colspan}](#attr-tdth-colspan)
    --- Number of columns that the cell is to span
:   [`rowspan`{#the-th-element:attr-tdth-rowspan}](#attr-tdth-rowspan)
    --- Number of rows that the cell is to span
:   [`headers`{#the-th-element:attr-tdth-headers}](#attr-tdth-headers)
    --- The header cells for this cell
:   [`scope`{#the-th-element:attr-th-scope}](#attr-th-scope) ---
    Specifies which cells the header cell applies to
:   [`abbr`{#the-th-element:attr-th-abbr}](#attr-th-abbr) ---
    Alternative label to use for the header cell when referencing the
    cell in other contexts

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-th-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-th).
:   [For implementers](https://w3c.github.io/html-aam/#el-th).

[DOM interface](dom.html#concept-element-dom){#the-th-element:concept-element-dom}:
:   Uses
    [`HTMLTableCellElement`{#the-th-element:htmltablecellelement}](#htmltablecellelement),
    as defined for
    [`td`{#the-th-element:the-td-element-2}](#the-td-element) elements.

The [`th`{#the-th-element:the-th-element-4}](#the-th-element) element
[represents](dom.html#represents){#the-th-element:represents} a header
[cell](#concept-cell){#the-th-element:concept-cell} in a table.

The [`th`{#the-th-element:the-th-element-5}](#the-th-element) element
may have a [`scope`]{#attr-th-scope .dfn dfn-for="th"
dfn-type="element-attr"} content attribute specified.

The [`scope`{#the-th-element:attr-th-scope-2}](#attr-th-scope) attribute
is an [enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-th-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`row`]{#attr-th-scope-row .dfn dfn-for="th/scope"
dfn-type="attr-value"}

[Row]{#attr-th-scope-row-state .dfn}

The header cell applies to some of the subsequent cells in the same
row(s).

[`col`]{#attr-th-scope-col .dfn dfn-for="th/scope"
dfn-type="attr-value"}

[Column]{#attr-th-scope-column-state .dfn}

The header cell applies to some of the subsequent cells in the same
column(s).

[`rowgroup`]{#attr-th-scope-rowgroup .dfn dfn-for="th/scope"
dfn-type="attr-value"}

[Row Group]{#attr-th-scope-rowgroup-state .dfn}

The header cell applies to all the remaining cells in the row group.

[`colgroup`]{#attr-th-scope-colgroup .dfn dfn-for="th/scope"
dfn-type="attr-value"}

[Column Group]{#attr-th-scope-colgroup-state .dfn}

The header cell applies to all the remaining cells in the column group.

The attribute\'s *[missing value
default](common-microsyntaxes.html#missing-value-default)* and *[invalid
value default](common-microsyntaxes.html#invalid-value-default)* are
both the [Auto]{#attr-th-scope-auto-state .dfn} state. (In this state
the header cell applies to a set of cells selected based on context.)

A [`th`{#the-th-element:the-th-element-6}](#the-th-element) element\'s
[`scope`{#the-th-element:attr-th-scope-3}](#attr-th-scope) attribute
must not be in the [Row
Group](#attr-th-scope-rowgroup-state){#the-th-element:attr-th-scope-rowgroup-state}
state if the element is not anchored in a [row
group](#concept-row-group){#the-th-element:concept-row-group}, nor in
the [Column
Group](#attr-th-scope-colgroup-state){#the-th-element:attr-th-scope-colgroup-state}
state if the element is not anchored in a [column
group](#concept-column-group){#the-th-element:concept-column-group}.

The [`th`{#the-th-element:the-th-element-7}](#the-th-element) element
may have an [`abbr`]{#attr-th-abbr .dfn dfn-for="th"
dfn-type="element-attr"} content attribute specified. Its value must be
an alternative label for the header cell, to be used when referencing
the cell in other contexts (e.g. when describing the header cells that
apply to a data cell). It is typically an abbreviated form of the full
header cell, but can also be an expansion, or merely a different
phrasing.

The [`th`{#the-th-element:the-th-element-8}](#the-th-element) element
and its
[`colspan`{#the-th-element:attr-tdth-colspan-2}](#attr-tdth-colspan),
[`rowspan`{#the-th-element:attr-tdth-rowspan-2}](#attr-tdth-rowspan),
[`headers`{#the-th-element:attr-tdth-headers-2}](#attr-tdth-headers),
and [`scope`{#the-th-element:attr-th-scope-4}](#attr-th-scope)
attributes take part in the [table
model](#table-model){#the-th-element:table-model}.

::: example
The following example shows how the
[`scope`{#the-th-element:attr-th-scope-5}](#attr-th-scope) attribute\'s
[`rowgroup`{#the-th-element:attr-th-scope-rowgroup}](#attr-th-scope-rowgroup)
value affects which data cells a header cell applies to.

Here is a markup fragment showing a table:

``` html
<table>
 <thead>
  <tr> <th> ID <th> Measurement <th> Average <th> Maximum
 <tbody>
  <tr> <td> <th scope=rowgroup> Cats <td> <td>
  <tr> <td> 93 <th> Legs <td> 3.5 <td> 4
  <tr> <td> 10 <th> Tails <td> 1 <td> 1
 <tbody>
  <tr> <td> <th scope=rowgroup> English speakers <td> <td>
  <tr> <td> 32 <th> Legs <td> 2.67 <td> 4
  <tr> <td> 35 <th> Tails <td> 0.33 <td> 1
</table>
```

This would result in the following table:

ID

Measurement

Average

Maximum

Cats

93

Legs

3.5

4

10

Tails

1

1

English speakers

32

Legs

2.67

4

35

Tails

0.33

1

The headers in the first row all apply directly down to the rows in
their column.

The headers with a
[`scope`{#the-th-element:attr-th-scope-6}](#attr-th-scope) attribute in
the [Row
Group](#attr-th-scope-rowgroup-state){#the-th-element:attr-th-scope-rowgroup-state-2}
state apply to all the cells in their row group other than the cells in
the first column.

The remaining headers apply just to the cells to the right of them.

![](/images/table-scope-diagram.png){width="459" height="256"}
:::

#### [4.9.11]{.secno} Attributes common to [`td`{#attributes-common-to-td-and-th-elements:the-td-element}](#the-td-element) and [`th`{#attributes-common-to-td-and-th-elements:the-th-element}](#the-th-element) elements[](#attributes-common-to-td-and-th-elements){.self-link}

The
[`td`{#attributes-common-to-td-and-th-elements:the-td-element-2}](#the-td-element)
and
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-2}](#the-th-element)
elements may have a [`colspan`]{#attr-tdth-colspan .dfn dfn-for="td,th"
dfn-type="element-attr"} content attribute specified, whose value must
be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#attributes-common-to-td-and-th-elements:valid-non-negative-integer}
greater than zero and less than or equal to 1000.

The
[`td`{#attributes-common-to-td-and-th-elements:the-td-element-3}](#the-td-element)
and
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-3}](#the-th-element)
elements may also have a [`rowspan`]{#attr-tdth-rowspan .dfn
dfn-for="td,th" dfn-type="element-attr"} content attribute specified,
whose value must be a [valid non-negative
integer](common-microsyntaxes.html#valid-non-negative-integer){#attributes-common-to-td-and-th-elements:valid-non-negative-integer-2}
less than or equal to 65534. For this attribute, the value zero means
that the cell is to span all the remaining rows in the row group.

These attributes give the number of columns and rows respectively that
the cell is to span. These attributes must not be used to overlap cells,
as described in the description of the [table
model](#table-model){#attributes-common-to-td-and-th-elements:table-model}.

------------------------------------------------------------------------

The
[`td`{#attributes-common-to-td-and-th-elements:the-td-element-4}](#the-td-element)
and
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-4}](#the-th-element)
element may have a [`headers`]{#attr-tdth-headers .dfn dfn-for="td,th"
dfn-type="element-attr"} content attribute specified. The
[`headers`{#attributes-common-to-td-and-th-elements:attr-tdth-headers}](#attr-tdth-headers)
attribute, if specified, must contain a string consisting of an
[unordered set of unique space-separated
tokens](common-microsyntaxes.html#unordered-set-of-unique-space-separated-tokens){#attributes-common-to-td-and-th-elements:unordered-set-of-unique-space-separated-tokens},
none of which are [identical
to](https://infra.spec.whatwg.org/#string-is){#attributes-common-to-td-and-th-elements:identical-to
x-internal="identical-to"} another token and each of which must have the
value of an
[ID](https://dom.spec.whatwg.org/#concept-id){#attributes-common-to-td-and-th-elements:concept-id
x-internal="concept-id"} of a
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-5}](#the-th-element)
element taking part in the same
[table](#concept-table){#attributes-common-to-td-and-th-elements:concept-table}
as the
[`td`{#attributes-common-to-td-and-th-elements:the-td-element-5}](#the-td-element)
or
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-6}](#the-th-element)
element (as defined by the [table
model](#table-model){#attributes-common-to-td-and-th-elements:table-model-2}).

A
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-7}](#the-th-element)
element with
[ID](https://dom.spec.whatwg.org/#concept-id){#attributes-common-to-td-and-th-elements:concept-id-2
x-internal="concept-id"} `id`{.variable} is said to be *directly
targeted* by all
[`td`{#attributes-common-to-td-and-th-elements:the-td-element-6}](#the-td-element)
and
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-8}](#the-th-element)
elements in the same
[table](#concept-table){#attributes-common-to-td-and-th-elements:concept-table-2}
that have
[`headers`{#attributes-common-to-td-and-th-elements:attr-tdth-headers-2}](#attr-tdth-headers)
attributes whose values include as one of their tokens the
[ID](https://dom.spec.whatwg.org/#concept-id){#attributes-common-to-td-and-th-elements:concept-id-3
x-internal="concept-id"} `id`{.variable}. A
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-9}](#the-th-element)
element `A`{.variable} is said to be *targeted* by a
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-10}](#the-th-element)
or
[`td`{#attributes-common-to-td-and-th-elements:the-td-element-7}](#the-td-element)
element `B`{.variable} if either `A`{.variable} is *directly targeted*
by `B`{.variable} or if there exists an element `C`{.variable} that is
itself *targeted* by the element `B`{.variable} and `A`{.variable} is
*directly targeted* by `C`{.variable}.

A
[`th`{#attributes-common-to-td-and-th-elements:the-th-element-11}](#the-th-element)
element must not be *targeted* by itself.

The
[`colspan`{#attributes-common-to-td-and-th-elements:attr-tdth-colspan}](#attr-tdth-colspan),
[`rowspan`{#attributes-common-to-td-and-th-elements:attr-tdth-rowspan}](#attr-tdth-rowspan),
and
[`headers`{#attributes-common-to-td-and-th-elements:attr-tdth-headers-3}](#attr-tdth-headers)
attributes take part in the [table
model](#table-model){#attributes-common-to-td-and-th-elements:table-model-3}.

------------------------------------------------------------------------

`cell`{.variable}`.`[`cellIndex`](#dom-tdth-cellindex){#dom-tdth-cellindex-dev}

:   Returns the position of the cell in the row\'s
    [`cells`{#attributes-common-to-td-and-th-elements:dom-tr-cells}](#dom-tr-cells)
    list. This does not necessarily correspond to the
    `x`{.variable}-position of the cell in the table, since earlier
    cells might cover multiple rows or columns.

    Returns −1 if the element isn\'t in a row.

The [`colSpan`]{#dom-tdth-colspan .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#attributes-common-to-td-and-th-elements:reflect}
the
[`colspan`{#attributes-common-to-td-and-th-elements:attr-tdth-colspan-2}](#attr-tdth-colspan)
content attribute. It is [clamped to the
range](common-dom-interfaces.html#clamped-to-the-range){#attributes-common-to-td-and-th-elements:clamped-to-the-range}
\[1, 1000\], and its [default
value](common-dom-interfaces.html#default-value){#attributes-common-to-td-and-th-elements:default-value}
is 1.

The [`rowSpan`]{#dom-tdth-rowspan .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#attributes-common-to-td-and-th-elements:reflect-2}
the
[`rowspan`{#attributes-common-to-td-and-th-elements:attr-tdth-rowspan-2}](#attr-tdth-rowspan)
content attribute. It is [clamped to the
range](common-dom-interfaces.html#clamped-to-the-range){#attributes-common-to-td-and-th-elements:clamped-to-the-range-2}
\[0, 65534\], and its [default
value](common-dom-interfaces.html#default-value){#attributes-common-to-td-and-th-elements:default-value-2}
is 1.

The [`headers`]{#dom-tdth-headers .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#attributes-common-to-td-and-th-elements:reflect-3}
the content attribute of the same name.

The [`cellIndex`]{#dom-tdth-cellindex .dfn
dfn-for="HTMLTableCellElement" dfn-type="attribute"} IDL attribute must,
if the element has a parent
[`tr`{#attributes-common-to-td-and-th-elements:the-tr-element}](#the-tr-element)
element, return the index of the cell\'s element in the parent
element\'s
[`cells`{#attributes-common-to-td-and-th-elements:dom-tr-cells-2}](#dom-tr-cells)
collection. If there is no such parent element, then the attribute must
return −1.

The [`scope`]{#dom-th-scope .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#attributes-common-to-td-and-th-elements:reflect-4}
the content attribute of the same name, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#attributes-common-to-td-and-th-elements:limited-to-only-known-values}.

The [`abbr`]{#dom-th-abbr .dfn dfn-for="HTMLTableCellElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#attributes-common-to-td-and-th-elements:reflect-5}
the content attribute of the same name.

#### [4.9.12]{.secno} []{#processing-model-2}Processing model[](#table-processing-model){.self-link} {#table-processing-model}

The various table elements and their content attributes together define
the [table model]{#table-model .dfn}.

A [table]{#concept-table .dfn} consists of cells aligned on a
two-dimensional grid of [slots]{#concept-slots .dfn} with coordinates
(`x`{.variable}, `y`{.variable}). The grid is finite, and is either
empty or has one or more slots. If the grid has one or more slots, then
the `x`{.variable} coordinates are always in the range
0 ≤ `x`{.variable} \< `x`{.variable}~`width`{.variable}~, and the
`y`{.variable} coordinates are always in the range
0 ≤ `y`{.variable} \< `y`{.variable}~`height`{.variable}~. If one or
both of `x`{.variable}~`width`{.variable}~ and
`y`{.variable}~`height`{.variable}~ are zero, then the table is empty
(has no slots). Tables correspond to
[`table`{#table-processing-model:the-table-element}](#the-table-element)
elements.

A [cell]{#concept-cell .dfn} is a set of slots anchored at a slot
(`cell`{.variable}~`x`{.variable}~, `cell`{.variable}~`y`{.variable}~),
and with a particular `width`{.variable} and `height`{.variable} such
that the cell covers all the slots with coordinates (`x`{.variable},
`y`{.variable}) where
`cell`{.variable}~`x`{.variable}~ ≤ `x`{.variable} \< `cell`{.variable}~`x`{.variable}~+`width`{.variable}
and
`cell`{.variable}~`y`{.variable}~ ≤ `y`{.variable} \< `cell`{.variable}~`y`{.variable}~+`height`{.variable}.
Cells can either be *data cells* or *header cells*. Data cells
correspond to
[`td`{#table-processing-model:the-td-element}](#the-td-element)
elements, and header cells correspond to
[`th`{#table-processing-model:the-th-element}](#the-th-element)
elements. Cells of both types can have zero or more associated header
cells.

It is possible, in certain error cases, for two cells to occupy the same
slot.

A [row]{#concept-row .dfn} is a complete set of slots from
`x`{.variable}=0 to `x`{.variable}=`x`{.variable}~`width`{.variable}~-1,
for a particular value of `y`{.variable}. Rows usually correspond to
[`tr`{#table-processing-model:the-tr-element}](#the-tr-element)
elements, though a [row
group](#concept-row-group){#table-processing-model:concept-row-group}
can have some implied
[rows](#concept-row){#table-processing-model:concept-row} at the end in
some cases involving
[cells](#concept-cell){#table-processing-model:concept-cell} spanning
multiple rows.

A [column]{#concept-column .dfn} is a complete set of slots from
`y`{.variable}=0 to
`y`{.variable}=`y`{.variable}~`height`{.variable}~-1, for a particular
value of `x`{.variable}. Columns can correspond to
[`col`{#table-processing-model:the-col-element}](#the-col-element)
elements. In the absence of
[`col`{#table-processing-model:the-col-element-2}](#the-col-element)
elements, columns are implied.

A [row group]{#concept-row-group .dfn} is a set of
[rows](#concept-row){#table-processing-model:concept-row-2} anchored at
a slot (0, `group`{.variable}~`y`{.variable}~) with a particular
`height`{.variable} such that the row group covers all the slots with
coordinates (`x`{.variable}, `y`{.variable}) where
0 ≤ `x`{.variable} \< `x`{.variable}~`width`{.variable}~ and
`group`{.variable}~`y`{.variable}~ ≤ `y`{.variable} \< `group`{.variable}~`y`{.variable}~+`height`{.variable}.
Row groups correspond to
[`tbody`{#table-processing-model:the-tbody-element}](#the-tbody-element),
[`thead`{#table-processing-model:the-thead-element}](#the-thead-element),
and
[`tfoot`{#table-processing-model:the-tfoot-element}](#the-tfoot-element)
elements. Not every row is necessarily in a row group.

A [column group]{#concept-column-group .dfn} is a set of
[columns](#concept-column){#table-processing-model:concept-column}
anchored at a slot (`group`{.variable}~`x`{.variable}~, 0) with a
particular `width`{.variable} such that the column group covers all the
slots with coordinates (`x`{.variable}, `y`{.variable}) where
`group`{.variable}~`x`{.variable}~ ≤ `x`{.variable} \< `group`{.variable}~`x`{.variable}~+`width`{.variable}
and 0 ≤ `y`{.variable} \< `y`{.variable}~`height`{.variable}~. Column
groups correspond to
[`colgroup`{#table-processing-model:the-colgroup-element}](#the-colgroup-element)
elements. Not every column is necessarily in a column group.

[Row
groups](#concept-row-group){#table-processing-model:concept-row-group-2}
cannot overlap each other. Similarly, [column
groups](#concept-column-group){#table-processing-model:concept-column-group}
cannot overlap each other.

A [cell](#concept-cell){#table-processing-model:concept-cell-2} cannot
cover slots that are from two or more [row
groups](#concept-row-group){#table-processing-model:concept-row-group-3}.
It is, however, possible for a cell to be in multiple [column
groups](#concept-column-group){#table-processing-model:concept-column-group-2}.
All the slots that form part of one cell are part of zero or one [row
groups](#concept-row-group){#table-processing-model:concept-row-group-4}
and zero or more [column
groups](#concept-column-group){#table-processing-model:concept-column-group-3}.

In addition to
[cells](#concept-cell){#table-processing-model:concept-cell-3},
[columns](#concept-column){#table-processing-model:concept-column-2},
[rows](#concept-row){#table-processing-model:concept-row-3}, [row
groups](#concept-row-group){#table-processing-model:concept-row-group-5},
and [column
groups](#concept-column-group){#table-processing-model:concept-column-group-4},
[tables](#concept-table){#table-processing-model:concept-table} can have
a
[`caption`{#table-processing-model:the-caption-element}](#the-caption-element)
element associated with them. This gives the table a heading, or legend.

A [table model error]{#table-model-error .dfn} is an error with the data
represented by
[`table`{#table-processing-model:the-table-element-2}](#the-table-element)
elements and their descendants. Documents must not have table model
errors.

##### [4.9.12.1]{.secno} Forming a table[](#forming-a-table){.self-link}

To determine which elements correspond to which slots in a
[table](#concept-table){#forming-a-table:concept-table} associated with
a [`table`{#forming-a-table:the-table-element}](#the-table-element)
element, to determine the dimensions of the table
(`x`{.variable}~`width`{.variable}~ and
`y`{.variable}~`height`{.variable}~), and to determine if there are any
[table model
errors](#table-model-error){#forming-a-table:table-model-error}, user
agents must use the following algorithm:

1.  Let `x`{.variable}~`width`{.variable}~ be zero.

2.  Let `y`{.variable}~`height`{.variable}~ be zero.

3.  Let
    `pending `{.variable}[`tfoot`{#forming-a-table:the-tfoot-element}](#the-tfoot-element)` elements`{.variable}
    be a list of
    [`tfoot`{#forming-a-table:the-tfoot-element-2}](#the-tfoot-element)
    elements, initially empty.

4.  Let `the table`{.variable} be the
    [table](#concept-table){#forming-a-table:concept-table-2}
    represented by the
    [`table`{#forming-a-table:the-table-element-2}](#the-table-element)
    element. The `x`{.variable}~`width`{.variable}~ and
    `y`{.variable}~`height`{.variable}~ variables give
    `the table`{.variable}\'s dimensions. `The table`{.variable} is
    initially empty.

5.  If the
    [`table`{#forming-a-table:the-table-element-3}](#the-table-element)
    element has no children elements, then return `the table`{.variable}
    (which will be empty).

6.  Associate the first
    [`caption`{#forming-a-table:the-caption-element}](#the-caption-element)
    element child of the
    [`table`{#forming-a-table:the-table-element-4}](#the-table-element)
    element with `the table`{.variable}. If there are no such children,
    then it has no associated
    [`caption`{#forming-a-table:the-caption-element-2}](#the-caption-element)
    element.

7.  Let the `current element`{.variable} be the first element child of
    the
    [`table`{#forming-a-table:the-table-element-5}](#the-table-element)
    element.

    If a step in this algorithm ever requires the
    `current element`{.variable} to be [advanced to the next child of
    the `table`]{#concept-table-advance .dfn} when there is no such next
    child, then the user agent must jump to the step labeled *end*, near
    the end of this algorithm.

8.  While the `current element`{.variable} is not one of the following
    elements,
    [advance](#concept-table-advance){#forming-a-table:concept-table-advance}
    the `current element`{.variable} to the next child of the
    [`table`{#forming-a-table:the-table-element-6}](#the-table-element):

    - [`colgroup`{#forming-a-table:the-colgroup-element}](#the-colgroup-element)
    - [`thead`{#forming-a-table:the-thead-element}](#the-thead-element)
    - [`tbody`{#forming-a-table:the-tbody-element}](#the-tbody-element)
    - [`tfoot`{#forming-a-table:the-tfoot-element-3}](#the-tfoot-element)
    - [`tr`{#forming-a-table:the-tr-element}](#the-tr-element)

9.  If the `current element`{.variable} is a
    [`colgroup`{#forming-a-table:the-colgroup-element-2}](#the-colgroup-element),
    follow these substeps:

    1.  *Column groups*: Process the `current element`{.variable}
        according to the appropriate case below:

        If the `current element`{.variable} has any [`col`{#forming-a-table:the-col-element}](#the-col-element) element children

        :   Follow these steps:

            1.  Let `x`{.variable}~`start`{.variable}~ have the value of
                `x`{.variable}~`width`{.variable}~.

            2.  Let the `current column`{.variable} be the first
                [`col`{#forming-a-table:the-col-element-2}](#the-col-element)
                element child of the
                [`colgroup`{#forming-a-table:the-colgroup-element-3}](#the-colgroup-element)
                element.

            3.  *Columns*: If the `current column`{.variable}
                [`col`{#forming-a-table:the-col-element-3}](#the-col-element)
                element has a
                [`span`{#forming-a-table:attr-col-span}](#attr-col-span)
                attribute, then parse its value using the [rules for
                parsing non-negative
                integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#forming-a-table:rules-for-parsing-non-negative-integers}.

                If the result of parsing the value is not an error or
                zero, then let `span`{.variable} be that value.

                Otherwise, if the
                [`col`{#forming-a-table:the-col-element-4}](#the-col-element)
                element has no
                [`span`{#forming-a-table:attr-col-span-2}](#attr-col-span)
                attribute, or if trying to parse the attribute\'s value
                resulted in an error or zero, then let `span`{.variable}
                be 1.

                If `span`{.variable} is greater than 1000, let it be
                1000 instead.

            4.  Increase `x`{.variable}~`width`{.variable}~ by
                `span`{.variable}.

            5.  Let the last `span`{.variable}
                [columns](#concept-column){#forming-a-table:concept-column}
                in `the table`{.variable} correspond to the
                `current column`{.variable}
                [`col`{#forming-a-table:the-col-element-5}](#the-col-element)
                element.

            6.  If `current column`{.variable} is not the last
                [`col`{#forming-a-table:the-col-element-6}](#the-col-element)
                element child of the
                [`colgroup`{#forming-a-table:the-colgroup-element-4}](#the-colgroup-element)
                element, then let the `current column`{.variable} be the
                next
                [`col`{#forming-a-table:the-col-element-7}](#the-col-element)
                element child of the
                [`colgroup`{#forming-a-table:the-colgroup-element-5}](#the-colgroup-element)
                element, and return to the step labeled *columns*.

            7.  Let all the last
                [columns](#concept-column){#forming-a-table:concept-column-2}
                in `the table`{.variable} from
                x=`x`{.variable}~`start`{.variable}~ to
                x=`x`{.variable}~`width`{.variable}~-1 form a new
                [column
                group](#concept-column-group){#forming-a-table:concept-column-group},
                anchored at the slot
                (`x`{.variable}~`start`{.variable}~, 0), with width
                `x`{.variable}~`width`{.variable}~-`x`{.variable}~`start`{.variable}~,
                corresponding to the
                [`colgroup`{#forming-a-table:the-colgroup-element-6}](#the-colgroup-element)
                element.

        If the `current element`{.variable} has no [`col`{#forming-a-table:the-col-element-8}](#the-col-element) element children

        :   1.  If the
                [`colgroup`{#forming-a-table:the-colgroup-element-7}](#the-colgroup-element)
                element has a
                [`span`{#forming-a-table:attr-colgroup-span}](#attr-colgroup-span)
                attribute, then parse its value using the [rules for
                parsing non-negative
                integers](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#forming-a-table:rules-for-parsing-non-negative-integers-2}.

                If the result of parsing the value is not an error or
                zero, then let `span`{.variable} be that value.

                Otherwise, if the
                [`colgroup`{#forming-a-table:the-colgroup-element-8}](#the-colgroup-element)
                element has no
                [`span`{#forming-a-table:attr-colgroup-span-2}](#attr-colgroup-span)
                attribute, or if trying to parse the attribute\'s value
                resulted in an error or zero, then let `span`{.variable}
                be 1.

                If `span`{.variable} is greater than 1000, let it be
                1000 instead.

            2.  Increase `x`{.variable}~`width`{.variable}~ by
                `span`{.variable}.

            3.  Let the last `span`{.variable}
                [columns](#concept-column){#forming-a-table:concept-column-3}
                in `the table`{.variable} form a new [column
                group](#concept-column-group){#forming-a-table:concept-column-group-2},
                anchored at the slot
                (`x`{.variable}~`width`{.variable}~-`span`{.variable},
                0), with width `span`{.variable}, corresponding to the
                [`colgroup`{#forming-a-table:the-colgroup-element-9}](#the-colgroup-element)
                element.

    2.  [Advance](#concept-table-advance){#forming-a-table:concept-table-advance-2}
        the `current element`{.variable} to the next child of the
        [`table`{#forming-a-table:the-table-element-7}](#the-table-element).

    3.  While the `current element`{.variable} is not one of the
        following elements,
        [advance](#concept-table-advance){#forming-a-table:concept-table-advance-3}
        the `current element`{.variable} to the next child of the
        [`table`{#forming-a-table:the-table-element-8}](#the-table-element):

        - [`colgroup`{#forming-a-table:the-colgroup-element-10}](#the-colgroup-element)
        - [`thead`{#forming-a-table:the-thead-element-2}](#the-thead-element)
        - [`tbody`{#forming-a-table:the-tbody-element-2}](#the-tbody-element)
        - [`tfoot`{#forming-a-table:the-tfoot-element-4}](#the-tfoot-element)
        - [`tr`{#forming-a-table:the-tr-element-2}](#the-tr-element)

    4.  If the `current element`{.variable} is a
        [`colgroup`{#forming-a-table:the-colgroup-element-11}](#the-colgroup-element)
        element, jump to the step labeled *column groups* above.

10. Let `y`{.variable}~`current`{.variable}~ be zero.

11. Let the `list of downward-growing cells`{.variable} be an empty
    list.

12. *Rows*: While the `current element`{.variable} is not one of the
    following elements,
    [advance](#concept-table-advance){#forming-a-table:concept-table-advance-4}
    the `current element`{.variable} to the next child of the
    [`table`{#forming-a-table:the-table-element-9}](#the-table-element):

    - [`thead`{#forming-a-table:the-thead-element-3}](#the-thead-element)
    - [`tbody`{#forming-a-table:the-tbody-element-3}](#the-tbody-element)
    - [`tfoot`{#forming-a-table:the-tfoot-element-5}](#the-tfoot-element)
    - [`tr`{#forming-a-table:the-tr-element-3}](#the-tr-element)

13. If the `current element`{.variable} is a
    [`tr`{#forming-a-table:the-tr-element-4}](#the-tr-element), then run
    the [algorithm for processing
    rows](#algorithm-for-processing-rows){#forming-a-table:algorithm-for-processing-rows},
    [advance](#concept-table-advance){#forming-a-table:concept-table-advance-5}
    the `current element`{.variable} to the next child of the
    [`table`{#forming-a-table:the-table-element-10}](#the-table-element),
    and return to the step labeled *rows*.

14. Run the [algorithm for ending a row
    group](#algorithm-for-ending-a-row-group){#forming-a-table:algorithm-for-ending-a-row-group}.

15. If the `current element`{.variable} is a
    [`tfoot`{#forming-a-table:the-tfoot-element-6}](#the-tfoot-element),
    then add that element to the list of
    `pending `{.variable}[`tfoot`{#forming-a-table:the-tfoot-element-7}](#the-tfoot-element)` elements`{.variable},
    [advance](#concept-table-advance){#forming-a-table:concept-table-advance-6}
    the `current element`{.variable} to the next child of the
    [`table`{#forming-a-table:the-table-element-11}](#the-table-element),
    and return to the step labeled *rows*.

16. The `current element`{.variable} is either a
    [`thead`{#forming-a-table:the-thead-element-4}](#the-thead-element)
    or a
    [`tbody`{#forming-a-table:the-tbody-element-4}](#the-tbody-element).

    Run the [algorithm for processing row
    groups](#algorithm-for-processing-row-groups){#forming-a-table:algorithm-for-processing-row-groups}.

17. [Advance](#concept-table-advance){#forming-a-table:concept-table-advance-7}
    the `current element`{.variable} to the next child of the
    [`table`{#forming-a-table:the-table-element-12}](#the-table-element).

18. Return to the step labeled *rows*.

19. *End*: For each
    [`tfoot`{#forming-a-table:the-tfoot-element-8}](#the-tfoot-element)
    element in the list of
    `pending `{.variable}[`tfoot`{#forming-a-table:the-tfoot-element-9}](#the-tfoot-element)` elements`{.variable},
    in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#forming-a-table:tree-order
    x-internal="tree-order"}, run the [algorithm for processing row
    groups](#algorithm-for-processing-row-groups){#forming-a-table:algorithm-for-processing-row-groups-2}.

20. If there exists a [row](#concept-row){#forming-a-table:concept-row}
    or [column](#concept-column){#forming-a-table:concept-column-4} in
    `the table`{.variable} containing only
    [slots](#concept-slots){#forming-a-table:concept-slots} that do not
    have a [cell](#concept-cell){#forming-a-table:concept-cell} anchored
    to them, then this is a [table model
    error](#table-model-error){#forming-a-table:table-model-error-2}.

21. Return `the table`{.variable}.

The [algorithm for processing row
groups]{#algorithm-for-processing-row-groups .dfn}, which is invoked by
the set of steps above for processing
[`thead`{#forming-a-table:the-thead-element-5}](#the-thead-element),
[`tbody`{#forming-a-table:the-tbody-element-5}](#the-tbody-element), and
[`tfoot`{#forming-a-table:the-tfoot-element-10}](#the-tfoot-element)
elements, is:

1.  Let `y`{.variable}~`start`{.variable}~ have the value of
    `y`{.variable}~`height`{.variable}~.

2.  For each [`tr`{#forming-a-table:the-tr-element-5}](#the-tr-element)
    element that is a child of the element being processed, in tree
    order, run the [algorithm for processing
    rows](#algorithm-for-processing-rows){#forming-a-table:algorithm-for-processing-rows-2}.

3.  If
    `y`{.variable}~`height`{.variable}~ \> `y`{.variable}~`start`{.variable}~,
    then let all the last
    [rows](#concept-row){#forming-a-table:concept-row-2} in
    `the table`{.variable} from y=`y`{.variable}~`start`{.variable}~ to
    y=`y`{.variable}~`height`{.variable}~-1 form a new [row
    group](#concept-row-group){#forming-a-table:concept-row-group},
    anchored at the slot with coordinate (0,
    `y`{.variable}~`start`{.variable}~), with height
    `y`{.variable}~`height`{.variable}~-`y`{.variable}~`start`{.variable}~,
    corresponding to the element being processed.

4.  Run the [algorithm for ending a row
    group](#algorithm-for-ending-a-row-group){#forming-a-table:algorithm-for-ending-a-row-group-2}.

The [algorithm for ending a row group]{#algorithm-for-ending-a-row-group
.dfn}, which is invoked by the set of steps above when starting and
ending a block of rows, is:

1.  While `y`{.variable}~`current`{.variable}~ is less than
    `y`{.variable}~`height`{.variable}~, follow these steps:

    1.  Run the [algorithm for growing downward-growing
        cells](#algorithm-for-growing-downward-growing-cells){#forming-a-table:algorithm-for-growing-downward-growing-cells}.

    2.  Increase `y`{.variable}~`current`{.variable}~ by 1.

2.  Empty the `list of downward-growing cells`{.variable}.

The [algorithm for processing rows]{#algorithm-for-processing-rows
.dfn}, which is invoked by the set of steps above for processing
[`tr`{#forming-a-table:the-tr-element-6}](#the-tr-element) elements, is:

1.  If `y`{.variable}~`height`{.variable}~ is equal to
    `y`{.variable}~`current`{.variable}~, then increase
    `y`{.variable}~`height`{.variable}~ by 1.
    (`y`{.variable}~`current`{.variable}~ is never *greater* than
    `y`{.variable}~`height`{.variable}~.)

2.  Let `x`{.variable}~`current`{.variable}~ be 0.

3.  Run the [algorithm for growing downward-growing
    cells](#algorithm-for-growing-downward-growing-cells){#forming-a-table:algorithm-for-growing-downward-growing-cells-2}.

4.  If the [`tr`{#forming-a-table:the-tr-element-7}](#the-tr-element)
    element being processed has no
    [`td`{#forming-a-table:the-td-element}](#the-td-element) or
    [`th`{#forming-a-table:the-th-element}](#the-th-element) element
    children, then increase `y`{.variable}~`current`{.variable}~ by 1,
    abort this set of steps, and return to the algorithm above.

5.  Let `current cell`{.variable} be the first
    [`td`{#forming-a-table:the-td-element-2}](#the-td-element) or
    [`th`{#forming-a-table:the-th-element-2}](#the-th-element) element
    child in the
    [`tr`{#forming-a-table:the-tr-element-8}](#the-tr-element) element
    being processed.

6.  *Cells*: While `x`{.variable}~`current`{.variable}~ is less than
    `x`{.variable}~`width`{.variable}~ and the slot with coordinate
    (`x`{.variable}~`current`{.variable}~,
    `y`{.variable}~`current`{.variable}~) already has a cell assigned to
    it, increase `x`{.variable}~`current`{.variable}~ by 1.

7.  If `x`{.variable}~`current`{.variable}~ is equal to
    `x`{.variable}~`width`{.variable}~, increase
    `x`{.variable}~`width`{.variable}~ by 1.
    (`x`{.variable}~`current`{.variable}~ is never *greater* than
    `x`{.variable}~`width`{.variable}~.)

8.  If the `current cell`{.variable} has a
    [`colspan`{#forming-a-table:attr-tdth-colspan}](#attr-tdth-colspan)
    attribute, then [parse that attribute\'s
    value](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#forming-a-table:rules-for-parsing-non-negative-integers-3},
    and let `colspan`{.variable} be the result.

    If parsing that value failed, or returned zero, or if the attribute
    is absent, then let `colspan`{.variable} be 1, instead.

    If `colspan`{.variable} is greater than 1000, let it be 1000
    instead.

9.  If the `current cell`{.variable} has a
    [`rowspan`{#forming-a-table:attr-tdth-rowspan}](#attr-tdth-rowspan)
    attribute, then [parse that attribute\'s
    value](common-microsyntaxes.html#rules-for-parsing-non-negative-integers){#forming-a-table:rules-for-parsing-non-negative-integers-4},
    and let `rowspan`{.variable} be the result.

    If parsing that value failed or if the attribute is absent, then let
    `rowspan`{.variable} be 1, instead.

    If `rowspan`{.variable} is greater than 65534, let it be 65534
    instead.

10. If `rowspan`{.variable} is zero and the
    [`table`{#forming-a-table:the-table-element-13}](#the-table-element)
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#forming-a-table:node-document
    x-internal="node-document"} is not set to [quirks
    mode](https://dom.spec.whatwg.org/#concept-document-quirks){#forming-a-table:quirks-mode
    x-internal="quirks-mode"}, then let `cell grows downward`{.variable}
    be true, and set `rowspan`{.variable} to 1. Otherwise, let
    `cell grows downward`{.variable} be false.

11. If
    `x`{.variable}~`width`{.variable}~ \< `x`{.variable}~`current`{.variable}~+`colspan`{.variable},
    then let `x`{.variable}~`width`{.variable}~ be
    `x`{.variable}~`current`{.variable}~+`colspan`{.variable}.

12. If
    `y`{.variable}~`height`{.variable}~ \< `y`{.variable}~`current`{.variable}~+`rowspan`{.variable},
    then let `y`{.variable}~`height`{.variable}~ be
    `y`{.variable}~`current`{.variable}~+`rowspan`{.variable}.

13. Let the slots with coordinates (`x`{.variable}, `y`{.variable}) such
    that
    `x`{.variable}~`current`{.variable}~ ≤ `x`{.variable} \< `x`{.variable}~`current`{.variable}~+`colspan`{.variable}
    and
    `y`{.variable}~`current`{.variable}~ ≤ `y`{.variable} \< `y`{.variable}~`current`{.variable}~+`rowspan`{.variable}
    be covered by a new
    [cell](#concept-cell){#forming-a-table:concept-cell-2}
    `c`{.variable}, anchored at (`x`{.variable}~`current`{.variable}~,
    `y`{.variable}~`current`{.variable}~), which has width
    `colspan`{.variable} and height `rowspan`{.variable}, corresponding
    to the `current cell`{.variable} element.

    If the `current cell`{.variable} element is a
    [`th`{#forming-a-table:the-th-element-3}](#the-th-element) element,
    let this new cell `c`{.variable} be a header cell; otherwise, let it
    be a data cell.

    To establish which header cells apply to the
    `current cell`{.variable} element, use the [algorithm for assigning
    header
    cells](#algorithm-for-assigning-header-cells){#forming-a-table:algorithm-for-assigning-header-cells}
    described in the next section.

    If any of the slots involved already had a
    [cell](#concept-cell){#forming-a-table:concept-cell-3} covering
    them, then this is a [table model
    error](#table-model-error){#forming-a-table:table-model-error-3}.
    Those slots now have two cells overlapping.

14. If `cell grows downward`{.variable} is true, then add the tuple
    {`c`{.variable}, `x`{.variable}~`current`{.variable}~,
    `colspan`{.variable}} to the
    `list of downward-growing cells`{.variable}.

15. Increase `x`{.variable}~`current`{.variable}~ by
    `colspan`{.variable}.

16. If `current cell`{.variable} is the last
    [`td`{#forming-a-table:the-td-element-3}](#the-td-element) or
    [`th`{#forming-a-table:the-th-element-4}](#the-th-element) element
    child in the
    [`tr`{#forming-a-table:the-tr-element-9}](#the-tr-element) element
    being processed, then increase `y`{.variable}~`current`{.variable}~
    by 1, abort this set of steps, and return to the algorithm above.

17. Let `current cell`{.variable} be the next
    [`td`{#forming-a-table:the-td-element-4}](#the-td-element) or
    [`th`{#forming-a-table:the-th-element-5}](#the-th-element) element
    child in the
    [`tr`{#forming-a-table:the-tr-element-10}](#the-tr-element) element
    being processed.

18. Return to the step labeled *cells*.

When the algorithms above require the user agent to run the [algorithm
for growing downward-growing
cells]{#algorithm-for-growing-downward-growing-cells .dfn}, the user
agent must, for each {`cell`{.variable},
`cell`{.variable}~`x`{.variable}~, `width`{.variable}} tuple in the
`list of downward-growing cells`{.variable}, if any, extend the
[cell](#concept-cell){#forming-a-table:concept-cell-4} `cell`{.variable}
so that it also covers the slots with coordinates (`x`{.variable},
`y`{.variable}~`current`{.variable}~), where
`cell`{.variable}~`x`{.variable}~ ≤ `x`{.variable} \< `cell`{.variable}~`x`{.variable}~+`width`{.variable}.

##### [4.9.12.2]{.secno} Forming relationships between data cells and header cells[](#header-and-data-cell-semantics){.self-link} {#header-and-data-cell-semantics}

Each cell can be assigned zero or more header cells. The [algorithm for
assigning header cells]{#algorithm-for-assigning-header-cells .dfn} to a
cell `principal cell`{.variable} is as follows.

1.  Let `header list`{.variable} be an empty list of cells.

2.  Let (`principal`{.variable}~`x`{.variable}~,
    `principal`{.variable}~`y`{.variable}~) be the coordinate of the
    slot to which the `principal cell`{.variable} is anchored.

3.  

    If the `principal cell`{.variable} has a [`headers`{#header-and-data-cell-semantics:attr-tdth-headers}](#attr-tdth-headers) attribute specified

    :   1.  Take the value of the `principal cell`{.variable}\'s
            [`headers`{#header-and-data-cell-semantics:attr-tdth-headers-2}](#attr-tdth-headers)
            attribute and [split it on ASCII
            whitespace](https://infra.spec.whatwg.org/#split-on-ascii-whitespace){#header-and-data-cell-semantics:split-a-string-on-spaces
            x-internal="split-a-string-on-spaces"}, letting
            `id list`{.variable} be the list of tokens obtained.

        2.  For each token in the `id list`{.variable}, if the first
            element in the
            [`Document`{#header-and-data-cell-semantics:document}](dom.html#document)
            with an
            [ID](https://dom.spec.whatwg.org/#concept-id){#header-and-data-cell-semantics:concept-id
            x-internal="concept-id"} equal to the token is a cell in the
            same
            [table](#concept-table){#header-and-data-cell-semantics:concept-table},
            and that cell is not the `principal cell`{.variable}, then
            add that cell to `header list`{.variable}.

    If `principal cell`{.variable} does not have a [`headers`{#header-and-data-cell-semantics:attr-tdth-headers-3}](#attr-tdth-headers) attribute specified

    :   1.  Let `principal`{.variable}~`width`{.variable}~ be the width
            of the `principal cell`{.variable}.

        2.  Let `principal`{.variable}~`height`{.variable}~ be the
            height of the `principal cell`{.variable}.

        3.  For each value of `y`{.variable} from
            `principal`{.variable}~`y`{.variable}~ to
            `principal`{.variable}~`y`{.variable}~+`principal`{.variable}~`height`{.variable}~-1,
            run the [internal algorithm for scanning and assigning
            header
            cells](#internal-algorithm-for-scanning-and-assigning-header-cells){#header-and-data-cell-semantics:internal-algorithm-for-scanning-and-assigning-header-cells},
            with the `principal cell`{.variable}, the
            `header list`{.variable}, the initial coordinate
            (`principal`{.variable}~`x`{.variable}~, `y`{.variable}),
            and the increments Δ`x`{.variable}=−1 and Δ`y`{.variable}=0.

        4.  For each value of `x`{.variable} from
            `principal`{.variable}~`x`{.variable}~ to
            `principal`{.variable}~`x`{.variable}~+`principal`{.variable}~`width`{.variable}~-1,
            run the [internal algorithm for scanning and assigning
            header
            cells](#internal-algorithm-for-scanning-and-assigning-header-cells){#header-and-data-cell-semantics:internal-algorithm-for-scanning-and-assigning-header-cells-2},
            with the `principal cell`{.variable}, the
            `header list`{.variable}, the initial coordinate
            (`x`{.variable}, `principal`{.variable}~`y`{.variable}~),
            and the increments Δ`x`{.variable}=0 and Δ`y`{.variable}=−1.

        5.  If the `principal cell`{.variable} is anchored in a [row
            group](#concept-row-group){#header-and-data-cell-semantics:concept-row-group},
            then add all header cells that are [row group
            headers](#row-group-header){#header-and-data-cell-semantics:row-group-header}
            and are anchored in the same row group with an
            `x`{.variable}-coordinate less than or equal to
            `principal`{.variable}~`x`{.variable}~+`principal`{.variable}~`width`{.variable}~-1
            and a `y`{.variable}-coordinate less than or equal to
            `principal`{.variable}~`y`{.variable}~+`principal`{.variable}~`height`{.variable}~-1
            to `header list`{.variable}.

        6.  If the `principal cell`{.variable} is anchored in a [column
            group](#concept-column-group){#header-and-data-cell-semantics:concept-column-group},
            then add all header cells that are [column group
            headers](#column-group-header){#header-and-data-cell-semantics:column-group-header}
            and are anchored in the same column group with an
            `x`{.variable}-coordinate less than or equal to
            `principal`{.variable}~`x`{.variable}~+`principal`{.variable}~`width`{.variable}~-1
            and a `y`{.variable}-coordinate less than or equal to
            `principal`{.variable}~`y`{.variable}~+`principal`{.variable}~`height`{.variable}~-1
            to `header list`{.variable}.

4.  Remove all the [empty
    cells](#empty-cell){#header-and-data-cell-semantics:empty-cell} from
    the `header list`{.variable}.

5.  Remove any duplicates from the `header list`{.variable}.

6.  Remove `principal cell`{.variable} from the `header list`{.variable}
    if it is there.

7.  Assign the headers in the `header list`{.variable} to the
    `principal cell`{.variable}.

The [internal algorithm for scanning and assigning header
cells]{#internal-algorithm-for-scanning-and-assigning-header-cells
.dfn}, given a `principal cell`{.variable}, a `header list`{.variable},
an initial coordinate (`initial`{.variable}~`x`{.variable}~,
`initial`{.variable}~`y`{.variable}~), and Δ`x`{.variable} and
Δ`y`{.variable} increments, is as follows:

1.  Let `x`{.variable} equal `initial`{.variable}~`x`{.variable}~.

2.  Let `y`{.variable} equal `initial`{.variable}~`y`{.variable}~.

3.  Let `opaque headers`{.variable} be an empty list of cells.

4.  

    If `principal cell`{.variable} is a header cell
    :   Let `in header block`{.variable} be true, and let
        `headers from current header block`{.variable} be a list of
        cells containing just the `principal cell`{.variable}.

    Otherwise

    :   Let `in header block`{.variable} be false and let
        `headers from current header block`{.variable} be an empty list
        of cells.

5.  *Loop*: Increment `x`{.variable} by Δ`x`{.variable}; increment
    `y`{.variable} by Δ`y`{.variable}.

    For each invocation of this algorithm, one of Δ`x`{.variable} and
    Δ`y`{.variable} will be −1, and the other will be 0.

6.  If either `x`{.variable} or `y`{.variable} are less than 0, then
    abort this internal algorithm.

7.  If there is no cell covering slot (`x`{.variable}, `y`{.variable}),
    or if there is more than one cell covering slot (`x`{.variable},
    `y`{.variable}), return to the substep labeled *loop*.

8.  Let `current cell`{.variable} be the cell covering slot
    (`x`{.variable}, `y`{.variable}).

9.  

    If `current cell`{.variable} is a header cell

    :   1.  Set `in header block`{.variable} to true.

        2.  Add `current cell`{.variable} to
            `headers from current header block`{.variable}.

        3.  Let `blocked`{.variable} be false.

        4.  

            If Δ`x`{.variable} is 0

            :   If there are any cells in the
                `opaque headers`{.variable} list anchored with the same
                `x`{.variable}-coordinate as the
                `current cell`{.variable}, and with the same width as
                `current cell`{.variable}, then let `blocked`{.variable}
                be true.

                If the `current cell`{.variable} is not a [column
                header](#column-header){#header-and-data-cell-semantics:column-header},
                then let `blocked`{.variable} be true.

            If Δ`y`{.variable} is 0

            :   If there are any cells in the
                `opaque headers`{.variable} list anchored with the same
                `y`{.variable}-coordinate as the
                `current cell`{.variable}, and with the same height as
                `current cell`{.variable}, then let `blocked`{.variable}
                be true.

                If the `current cell`{.variable} is not a [row
                header](#row-header){#header-and-data-cell-semantics:row-header},
                then let `blocked`{.variable} be true.

        5.  If `blocked`{.variable} is false, then add the
            `current cell`{.variable} to the `header list`{.variable}.

    If `current cell`{.variable} is a data cell and `in header block`{.variable} is true

    :   Set `in header block`{.variable} to false. Add all the cells in
        `headers from current header block`{.variable} to the
        `opaque headers`{.variable} list, and empty the
        `headers from current header block`{.variable} list.

10. Return to the step labeled *loop*.

A header cell anchored at the slot with coordinate (`x`{.variable},
`y`{.variable}) with width `width`{.variable} and height
`height`{.variable} is said to be a [column header]{#column-header .dfn}
if any of the following are true:

- the cell\'s
  [`scope`{#header-and-data-cell-semantics:attr-th-scope}](#attr-th-scope)
  attribute is in the
  [Column](#attr-th-scope-column-state){#header-and-data-cell-semantics:attr-th-scope-column-state}
  state; or

- the cell\'s
  [`scope`{#header-and-data-cell-semantics:attr-th-scope-2}](#attr-th-scope)
  attribute is in the
  [Auto](#attr-th-scope-auto-state){#header-and-data-cell-semantics:attr-th-scope-auto-state}
  state, and there are no data cells in any of the cells covering slots
  with `y`{.variable}-coordinates `y`{.variable} ..
  `y`{.variable}+`height`{.variable}-1.

A header cell anchored at the slot with coordinate (`x`{.variable},
`y`{.variable}) with width `width`{.variable} and height
`height`{.variable} is said to be a [row header]{#row-header .dfn} if
any of the following are true:

- the cell\'s
  [`scope`{#header-and-data-cell-semantics:attr-th-scope-3}](#attr-th-scope)
  attribute is in the
  [Row](#attr-th-scope-row-state){#header-and-data-cell-semantics:attr-th-scope-row-state}
  state; or

- the cell\'s
  [`scope`{#header-and-data-cell-semantics:attr-th-scope-4}](#attr-th-scope)
  attribute is in the
  [Auto](#attr-th-scope-auto-state){#header-and-data-cell-semantics:attr-th-scope-auto-state-2}
  state, the cell is not a [column
  header](#column-header){#header-and-data-cell-semantics:column-header-2},
  and there are no data cells in any of the cells covering slots with
  `x`{.variable}-coordinates `x`{.variable} ..
  `x`{.variable}+`width`{.variable}-1.

A header cell is said to be a [column group header]{#column-group-header
.dfn} if its
[`scope`{#header-and-data-cell-semantics:attr-th-scope-5}](#attr-th-scope)
attribute is in the [Column
Group](#attr-th-scope-colgroup-state){#header-and-data-cell-semantics:attr-th-scope-colgroup-state}
state.

A header cell is said to be a [row group header]{#row-group-header .dfn}
if its
[`scope`{#header-and-data-cell-semantics:attr-th-scope-6}](#attr-th-scope)
attribute is in the [Row
Group](#attr-th-scope-rowgroup-state){#header-and-data-cell-semantics:attr-th-scope-rowgroup-state}
state.

A cell is said to be an [empty cell]{#empty-cell .dfn} if it contains no
elements and its [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#header-and-data-cell-semantics:child-text-content
x-internal="child-text-content"}, if any, consists only of [ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#header-and-data-cell-semantics:space-characters
x-internal="space-characters"}.

#### [4.9.13]{.secno} Examples[](#table-examples){.self-link} {#table-examples}

*This section is non-normative.*

The following shows how one might mark up the bottom part of table 45 of
the Smithsonian physical tables, Volume 71:

``` html
<table>
 <caption>Specification values: <b>Steel</b>, <b>Castings</b>,
 Ann. A.S.T.M. A27-16, Class B;* P max. 0.06; S max. 0.05.</caption>
 <thead>
  <tr>
   <th rowspan=2>Grade.</th>
   <th rowspan=2>Yield Point.</th>
   <th colspan=2>Ultimate tensile strength</th>
   <th rowspan=2>Per cent elong. 50.8&nbsp;mm or 2&nbsp;in.</th>
   <th rowspan=2>Per cent reduct. area.</th>
  </tr>
  <tr>
   <th>kg/mm<sup>2</sup></th>
   <th>lb/in<sup>2</sup></th>
  </tr>
 </thead>
 <tbody>
  <tr>
   <td>Hard</td>
   <td>0.45 ultimate</td>
   <td>56.2</td>
   <td>80,000</td>
   <td>15</td>
   <td>20</td>
  </tr>
  <tr>
   <td>Medium</td>
   <td>0.45 ultimate</td>
   <td>49.2</td>
   <td>70,000</td>
   <td>18</td>
   <td>25</td>
  </tr>
  <tr>
   <td>Soft</td>
   <td>0.45 ultimate</td>
   <td>42.2</td>
   <td>60,000</td>
   <td>22</td>
   <td>30</td>
  </tr>
 </tbody>
</table>
```

This table could look like this:

Specification values: **Steel**, **Castings**, Ann. A.S.T.M. A27-16,
Class B;\* P max. 0.06; S max. 0.05.

Grade.

Yield Point.

Ultimate tensile strength

Per cent elong. 50.8 mm or 2 in.

Per cent reduct. area.

kg/mm^2^

lb/in^2^

Hard

0.45 ultimate

56.2

80,000

15

20

Medium

0.45 ultimate

49.2

70,000

18

25

Soft

0.45 ultimate

42.2

60,000

22

30

------------------------------------------------------------------------

The following shows how one might mark up the gross margin table on page
46 of Apple, Inc\'s 10-K filing for fiscal year 2008:

``` html
<table>
 <thead>
  <tr>
   <th>
   <th>2008
   <th>2007
   <th>2006
 <tbody>
  <tr>
   <th>Net sales
   <td>$ 32,479
   <td>$ 24,006
   <td>$ 19,315
  <tr>
   <th>Cost of sales
   <td>  21,334
   <td>  15,852
   <td>  13,717
 <tbody>
  <tr>
   <th>Gross margin
   <td>$ 11,145
   <td>$  8,154
   <td>$  5,598
 <tfoot>
  <tr>
   <th>Gross margin percentage
   <td>34.3%
   <td>34.0%
   <td>29.0%
</table>
```

This table could look like this:

2008

2007

2006

Net sales

\$ 32,479

\$ 24,006

\$ 19,315

Cost of sales

21,334

15,852

13,717

Gross margin

\$ 11,145

\$ 8,154

\$ 5,598

Gross margin percentage

34.3%

34.0%

29.0%

------------------------------------------------------------------------

The following shows how one might mark up the operating expenses table
from lower on the same page of that document:

``` html
<table>
 <colgroup> <col>
 <colgroup> <col> <col> <col>
 <thead>
  <tr> <th> <th>2008 <th>2007 <th>2006
 <tbody>
  <tr> <th scope=rowgroup> Research and development
       <td> $ 1,109 <td> $ 782 <td> $ 712
  <tr> <th scope=row> Percentage of net sales
       <td> 3.4% <td> 3.3% <td> 3.7%
 <tbody>
  <tr> <th scope=rowgroup> Selling, general, and administrative
       <td> $ 3,761 <td> $ 2,963 <td> $ 2,433
  <tr> <th scope=row> Percentage of net sales
       <td> 11.6% <td> 12.3% <td> 12.6%
</table>
```

This table could look like this:

2008

2007

2006

Research and development

\$ 1,109

\$ 782

\$ 712

Percentage of net sales

3.4%

3.3%

3.7%

Selling, general, and administrative

\$ 3,761

\$ 2,963

\$ 2,433

Percentage of net sales

11.6%

12.3%

12.6%

[← 4.8.15 MathML](embedded-content-other.html) --- [Table of
Contents](./) --- [4.10 Forms →](forms.html)
