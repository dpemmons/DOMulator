HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 4.11 Interactive elements](interactive-elements.html) --- [Table of
Contents](./) --- [4.12.5 The canvas element →](canvas.html)

1.  ::: {#toc-semantics}
    1.  [[4.12]{.secno} Scripting](scripting.html#scripting-3)
        1.  [[4.12.1]{.secno} The `script`
            element](scripting.html#the-script-element)
            1.  [[4.12.1.1]{.secno} Processing
                model](scripting.html#script-processing-model)
            2.  [[4.12.1.2]{.secno} Scripting
                languages](scripting.html#scriptingLanguages)
            3.  [[4.12.1.3]{.secno} Restrictions for contents of
                `script`
                elements](scripting.html#restrictions-for-contents-of-script-elements)
            4.  [[4.12.1.4]{.secno} Inline documentation for external
                scripts](scripting.html#inline-documentation-for-external-scripts)
            5.  [[4.12.1.5]{.secno} Interaction of `script` elements and
                XSLT](scripting.html#scriptTagXSLT)
        2.  [[4.12.2]{.secno} The `noscript`
            element](scripting.html#the-noscript-element)
        3.  [[4.12.3]{.secno} The `template`
            element](scripting.html#the-template-element)
            1.  [[4.12.3.1]{.secno} Interaction of `template` elements
                with XSLT and XPath](scripting.html#template-XSLT-XPath)
        4.  [[4.12.4]{.secno} The `slot`
            element](scripting.html#the-slot-element)
    :::

### [4.12]{.secno} Scripting[](#scripting-3){.self-link} {#scripting-3}

Scripts allow authors to add interactivity to their documents.

Authors are encouraged to use declarative alternatives to scripting
where possible, as declarative mechanisms are often more maintainable,
and many users disable scripting.

::: example
For example, instead of using a script to show or hide a section to show
more details, the
[`details`{#scripting-3:the-details-element}](interactive-elements.html#the-details-element)
element could be used.
:::

Authors are also encouraged to make their applications degrade
gracefully in the absence of scripting support.

::: example
For example, if an author provides a link in a table header to
dynamically resort the table, the link could also be made to function
without scripts by requesting the sorted table from the server.
:::

#### [4.12.1]{.secno} The [`script`]{#script .dfn dfn-type="element"} element[](#the-script-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/script](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script "The <script> HTML element is used to embed executable code or data; this is typically used to embed or refer to JavaScript code. The <script> element can also be used with other languages, such as WebGL's GLSL shader programming language and JSON.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerYes]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
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
[HTMLScriptElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLScriptElement "HTML <script> elements expose the HTMLScriptElement interface, which provides special properties and methods for manipulating the behavior and execution of <script> elements (beyond the inherited HTMLElement interface).")

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

[Categories](dom.html#concept-element-categories){#the-script-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-script-element:metadata-content-2}.
:   [Flow
    content](dom.html#flow-content-2){#the-script-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-script-element:phrasing-content-2}.
:   [Script-supporting
    element](dom.html#script-supporting-elements-2){#the-script-element:script-supporting-elements-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-script-element:concept-element-contexts}:
:   Where [metadata
    content](dom.html#metadata-content-2){#the-script-element:metadata-content-2-2}
    is expected.
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-script-element:phrasing-content-2-2}
    is expected.
:   Where [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-script-element:script-supporting-elements-2-2}
    are expected.

[Content model](dom.html#concept-element-content-model){#the-script-element:concept-element-content-model}:
:   If there is no
    [`src`{#the-script-element:attr-script-src}](#attr-script-src)
    attribute, depends on the value of the
    [`type`{#the-script-element:attr-script-type}](#attr-script-type)
    attribute, but must match [script content
    restrictions](#restrictions-for-contents-of-script-elements){#the-script-element:restrictions-for-contents-of-script-elements}.
:   If there *is* a
    [`src`{#the-script-element:attr-script-src-2}](#attr-script-src)
    attribute, the element must be either empty or contain only [script
    documentation](#inline-documentation-for-external-scripts){#the-script-element:inline-documentation-for-external-scripts}
    that also matches [script content
    restrictions](#restrictions-for-contents-of-script-elements){#the-script-element:restrictions-for-contents-of-script-elements-2}.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-script-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-script-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-script-element:global-attributes}
:   [`type`{#the-script-element:attr-script-type-2}](#attr-script-type)
    --- Type of script
:   [`src`{#the-script-element:attr-script-src-3}](#attr-script-src) ---
    Address of the resource
:   [`nomodule`{#the-script-element:attr-script-nomodule}](#attr-script-nomodule)
    --- Prevents execution in user agents that support [module
    scripts](webappapis.html#module-script){#the-script-element:module-script}
:   [`async`{#the-script-element:attr-script-async}](#attr-script-async)
    --- Execute script when available, without blocking while fetching
:   [`defer`{#the-script-element:attr-script-defer}](#attr-script-defer)
    --- Defer script execution
:   [`blocking`{#the-script-element:attr-script-blocking}](#attr-script-blocking)
    --- Whether the element is [potentially
    render-blocking](urls-and-fetching.html#potentially-render-blocking){#the-script-element:potentially-render-blocking}
:   [`crossorigin`{#the-script-element:attr-script-crossorigin}](#attr-script-crossorigin)
    --- How the element handles crossorigin requests
:   [`referrerpolicy`{#the-script-element:attr-script-referrerpolicy}](#attr-script-referrerpolicy)
    --- [Referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-script-element:referrer-policy
    x-internal="referrer-policy"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-script-element:concept-fetch
    x-internal="concept-fetch"} initiated by the element
:   [`integrity`{#the-script-element:attr-script-integrity}](#attr-script-integrity)
    --- Integrity metadata used in Subresource Integrity checks
    [\[SRI\]](references.html#refsSRI)
:   [`fetchpriority`{#the-script-element:attr-script-fetchpriority}](#attr-script-fetchpriority)
    --- Sets the
    [priority](https://fetch.spec.whatwg.org/#request-priority){#the-script-element:concept-request-priority
    x-internal="concept-request-priority"} for
    [fetches](https://fetch.spec.whatwg.org/#concept-fetch){#the-script-element:concept-fetch-2
    x-internal="concept-fetch"} initiated by the element

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-script-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-script).
:   [For implementers](https://w3c.github.io/html-aam/#el-script).

[DOM interface](dom.html#concept-element-dom){#the-script-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLScriptElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString type;
      [CEReactions] attribute USVString src;
      [CEReactions] attribute boolean noModule;
      [CEReactions] attribute boolean async;
      [CEReactions] attribute boolean defer;
      [SameObject, PutForwards=value] readonly attribute DOMTokenList blocking;
      [CEReactions] attribute DOMString? crossOrigin;
      [CEReactions] attribute DOMString referrerPolicy;
      [CEReactions] attribute DOMString integrity;
      [CEReactions] attribute DOMString fetchPriority;

      [CEReactions] attribute DOMString text;

      static boolean supports(DOMString type);

      // also has obsolete members
    };
    ```

The
[`script`{#the-script-element:the-script-element}](#the-script-element)
element allows authors to include dynamic script, instructions to the
user agent, and data blocks in their documents. The element does not
[represent](dom.html#represents){#the-script-element:represents} content
for the user.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/script#attr-type](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script#attr-type "The <script> HTML element is used to embed executable code or data; this is typically used to embed or refer to JavaScript code. The <script> element can also be used with other languages, such as WebGL's GLSL shader programming language and JSON.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari≤4+]{.safari .yes}[Chrome1+]{.chrome
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

The script element has two core attributes. The
[`type`]{#attr-script-type .dfn dfn-for="script"
dfn-type="element-attr"} attribute allows customization of the type of
script represented:

- Omitting the attribute, setting it to the empty string, or setting it
  to a [JavaScript MIME type essence
  match](https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match){#the-script-element:javascript-mime-type-essence-match
  x-internal="javascript-mime-type-essence-match"} means that the script
  is a [classic
  script](webappapis.html#classic-script){#the-script-element:classic-script},
  to be interpreted according to the JavaScript
  *[Script](https://tc39.es/ecma262/#prod-Script){x-internal="js-prod-script"}*
  top-level production. Authors should omit the
  [`type`{#the-script-element:attr-script-type-3}](#attr-script-type)
  attribute instead of redundantly setting it.

- Setting the attribute to an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-script-element:ascii-case-insensitive
  x-internal="ascii-case-insensitive"} match for \"`module`\" means that
  the script is a [JavaScript module
  script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script},
  to be interpreted according to the JavaScript
  *[Module](https://tc39.es/ecma262/#prod-Module){x-internal="js-prod-module"}*
  top-level production.

- Setting the attribute to an [ASCII
  case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-script-element:ascii-case-insensitive-2
  x-internal="ascii-case-insensitive"} match for \"`importmap`\" means
  that the script is an [import
  map](webappapis.html#import-map){#the-script-element:import-map},
  containing JSON that will be used to control the behavior of [module
  specifier
  resolution](webappapis.html#resolve-a-module-specifier){#the-script-element:resolve-a-module-specifier}.

- Setting the attribute to any other value means that the script is a
  [data block]{#data-block .dfn}, which is not processed by the user
  agent, but instead by author script or other tools. Authors must use a
  [valid MIME type
  string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#the-script-element:valid-mime-type-string
  x-internal="valid-mime-type-string"} that is not a [JavaScript MIME
  type essence
  match](https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match){#the-script-element:javascript-mime-type-essence-match-2
  x-internal="javascript-mime-type-essence-match"} to denote data
  blocks.

  The requirement that [data
  blocks](#data-block){#the-script-element:data-block} must be denoted
  using a [valid MIME type
  string](https://mimesniff.spec.whatwg.org/#valid-mime-type){#the-script-element:valid-mime-type-string-2
  x-internal="valid-mime-type-string"} is in place to avoid potential
  future collisions. Values for the
  [`type`{#the-script-element:attr-script-type-4}](#attr-script-type)
  attribute that are not MIME types, like \"`module`\" or
  \"`importmap`\", are used by the standard to denote types of scripts
  which have special behavior in user agents. By using a valid MIME type
  string now, you ensure that your data block will not ever be
  reinterpreted as a different script type, even in future user agents.

The second core attribute is the [`src`]{#attr-script-src .dfn
dfn-for="script" dfn-type="element-attr"} attribute. It must only be
specified for [classic
scripts](webappapis.html#classic-script){#the-script-element:classic-script-2}
and [JavaScript module
scripts](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-2},
and denotes that instead of using the element\'s [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-script-element:child-text-content
x-internal="child-text-content"} as the script content, the script will
be fetched from the specified
[URL](https://url.spec.whatwg.org/#concept-url){#the-script-element:url
x-internal="url"}. If
[`src`{#the-script-element:attr-script-src-4}](#attr-script-src) is
specified, it must be a [valid non-empty URL potentially surrounded by
spaces](urls-and-fetching.html#valid-non-empty-url-potentially-surrounded-by-spaces){#the-script-element:valid-non-empty-url-potentially-surrounded-by-spaces}.

Which other attributes may be specified on a given
[`script`{#the-script-element:the-script-element-2}](#the-script-element)
element is determined by the following table:

[`nomodule`{#the-script-element:attr-script-nomodule-2}](#attr-script-nomodule)

[`async`{#the-script-element:attr-script-async-2}](#attr-script-async)

[`defer`{#the-script-element:attr-script-defer-2}](#attr-script-defer)

[`blocking`{#the-script-element:attr-script-blocking-2}](#attr-script-blocking)

[`crossorigin`{#the-script-element:attr-script-crossorigin-2}](#attr-script-crossorigin)

[`referrerpolicy`{#the-script-element:attr-script-referrerpolicy-2}](#attr-script-referrerpolicy)

[`integrity`{#the-script-element:attr-script-integrity-2}](#attr-script-integrity)

[`fetchpriority`{#the-script-element:attr-script-fetchpriority-2}](#attr-script-fetchpriority)

External classic scripts

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Yes

Inline classic scripts

Yes

·

·

·

Yes\*

Yes\*

·

·†

External module scripts

·

Yes

·

Yes

Yes

Yes

Yes

Yes

Inline module scripts

·

Yes

·

·

Yes\*

Yes\*

·

·†

Import maps

·

·

·

·

·

·

·

·

Data blocks

·

·

·

·

·

·

·

·

[\* Although inline scripts have no initial fetches, the
[`crossorigin`{#the-script-element:attr-script-crossorigin-3}](#attr-script-crossorigin)
and
[`referrerpolicy`{#the-script-element:attr-script-referrerpolicy-3}](#attr-script-referrerpolicy)
attribute on inline scripts affects the [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-script-element:concept-request-credentials-mode
x-internal="concept-request-credentials-mode"} and [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-script-element:referrer-policy-2
x-internal="referrer-policy"} used by module imports, including dynamic
[`import()`{#the-script-element:import()}](https://tc39.es/ecma262/#sec-import-calls){x-internal="import()"}.]{.small}

[† Unlike
[`crossorigin`{#the-script-element:attr-script-crossorigin-4}](#attr-script-crossorigin)
and
[`referrerpolicy`{#the-script-element:attr-script-referrerpolicy-4}](#attr-script-referrerpolicy),
[`fetchpriority`{#the-script-element:attr-script-fetchpriority-3}](#attr-script-fetchpriority)
does not affect module imports. See some discussion in [issue
#10276](https://github.com/whatwg/html/issues/10276).]{.small}

------------------------------------------------------------------------

The contents of inline
[`script`{#the-script-element:the-script-element-3}](#the-script-element)
elements, or the external script resource, must conform with the
requirements of the JavaScript specification\'s
*[Script](https://tc39.es/ecma262/#prod-Script){x-internal="js-prod-script"}*
or
*[Module](https://tc39.es/ecma262/#prod-Module){x-internal="js-prod-module"}*
productions, for [classic
scripts](webappapis.html#classic-script){#the-script-element:classic-script-3}
and [JavaScript module
scripts](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-3}
respectively. [\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

The contents of inline
[`script`{#the-script-element:the-script-element-4}](#the-script-element)
elements for [import
maps](webappapis.html#import-map){#the-script-element:import-map-2} must
conform with the [import map authoring
requirements](webappapis.html#import-map-authoring-requirements){#the-script-element:import-map-authoring-requirements}.

When used to include [data
blocks](#data-block){#the-script-element:data-block-2}, the data must be
embedded inline, the format of the data must be given using the
[`type`{#the-script-element:attr-script-type-5}](#attr-script-type)
attribute, and the contents of the
[`script`{#the-script-element:the-script-element-5}](#the-script-element)
element must conform to the requirements defined for the format used.

------------------------------------------------------------------------

The [`nomodule`]{#attr-script-nomodule .dfn dfn-for="script"
dfn-type="element-attr"} attribute is a [boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-script-element:boolean-attribute}
that prevents a script from being executed in user agents that support
[module
scripts](webappapis.html#module-script){#the-script-element:module-script-2}.
This allows selective execution of [module
scripts](webappapis.html#module-script){#the-script-element:module-script-3}
in modern user agents and [classic
scripts](webappapis.html#classic-script){#the-script-element:classic-script-4}
in older user agents, [as shown below](#script-nomodule-example).

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Element/script#attr-async](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script#attr-async "The <script> HTML element is used to embed executable code or data; this is typically used to embed or refer to JavaScript code. The <script> element can also be used with other languages, such as WebGL's GLSL shader programming language and JSON.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari≤4+]{.safari .yes}[Chrome1+]{.chrome
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
[Element/script#attr-defer](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script#attr-defer "The <script> HTML element is used to embed executable code or data; this is typically used to embed or refer to JavaScript code. The <script> element can also be used with other languages, such as WebGL's GLSL shader programming language and JSON.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

The [`async`]{#attr-script-async .dfn dfn-for="script"
dfn-type="element-attr"} and [`defer`]{#attr-script-defer .dfn
dfn-for="script" dfn-type="element-attr"} attributes are [boolean
attributes](common-microsyntaxes.html#boolean-attribute){#the-script-element:boolean-attribute-2}
that indicate how the script should be evaluated. There are several
possible modes that can be selected using these attributes, depending on
the script\'s type.

For external [classic
scripts](webappapis.html#classic-script){#the-script-element:classic-script-5},
if the
[`async`{#the-script-element:attr-script-async-3}](#attr-script-async)
attribute is present, then the classic script will be fetched [in
parallel](infrastructure.html#in-parallel){#the-script-element:in-parallel}
to parsing and evaluated as soon as it is available (potentially before
parsing completes). If the
[`async`{#the-script-element:attr-script-async-4}](#attr-script-async)
attribute is not present but the
[`defer`{#the-script-element:attr-script-defer-3}](#attr-script-defer)
attribute is present, then the classic script will be fetched [in
parallel](infrastructure.html#in-parallel){#the-script-element:in-parallel-2}
and evaluated when the page has finished parsing. If neither attribute
is present, then the script is fetched and evaluated immediately,
blocking parsing until these are both complete.

For [module
scripts](webappapis.html#module-script){#the-script-element:module-script-4},
if the
[`async`{#the-script-element:attr-script-async-5}](#attr-script-async)
attribute is present, then the module script and all its dependencies
will be fetched [in
parallel](infrastructure.html#in-parallel){#the-script-element:in-parallel-3}
to parsing, and the module script will be evaluated as soon as it is
available (potentially before parsing completes). Otherwise, the module
script and its dependencies will be fetched [in
parallel](infrastructure.html#in-parallel){#the-script-element:in-parallel-4}
to parsing and evaluated when the page has finished parsing. (The
[`defer`{#the-script-element:attr-script-defer-4}](#attr-script-defer)
attribute has no effect on module scripts.)

This is all summarized in the following schematic diagram:

![With \<script\>, parsing is interrupted by fetching and execution.
With \<script defer\>, fetching is parallel to parsing and execution
takes place after all parsing has finished. And with \<script async\>,
fetching is parallel to parsing but once it finishes parsing is
interrupted to execute the script. The story for \<script
type=\"module\"\> is similar to \<script defer\>, but the dependencies
will be fetched as well, and the story for \<script type=\"module\"
async\> is similar to \<script async\> with the extra dependency
fetching.](/images/asyncdefer.svg){.darkmode-aware
style="width: 80%; min-width: 820px;"}

The exact processing details for these attributes are, for mostly
historical reasons, somewhat non-trivial, involving a number of aspects
of HTML. The implementation requirements are therefore by necessity
scattered throughout the specification. The algorithms
[below](#script-processing-model) describe the core of this processing,
but these algorithms reference and are referenced by the parsing rules
for
[`script`{#the-script-element:the-script-element-6}](#the-script-element)
[start](parsing.html#scriptTag) and [end](parsing.html#scriptEndTag)
tags in HTML, [in foreign content](parsing.html#scriptForeignEndTag),
and [in XML](xhtml.html#scriptTagXML), the rules for the
[`document.write()`{#the-script-element:dom-document-write}](dynamic-markup-insertion.html#dom-document-write)
method, the handling of [scripting](webappapis.html#scripting), etc.

When inserted using the
[`document.write()`{#the-script-element:dom-document-write-2}](dynamic-markup-insertion.html#dom-document-write)
method,
[`script`{#the-script-element:the-script-element-7}](#the-script-element)
elements [usually](parsing.html#document-written-scripts-intervention)
execute (typically blocking further script execution or HTML parsing).
When inserted using the
[`innerHTML`{#the-script-element:dom-element-innerhtml}](dynamic-markup-insertion.html#dom-element-innerhtml)
and
[`outerHTML`{#the-script-element:dom-element-outerhtml}](dynamic-markup-insertion.html#dom-element-outerhtml)
attributes, they do not execute at all.

The
[`defer`{#the-script-element:attr-script-defer-5}](#attr-script-defer)
attribute may be specified even if the
[`async`{#the-script-element:attr-script-async-6}](#attr-script-async)
attribute is specified, to cause legacy web browsers that only support
[`defer`{#the-script-element:attr-script-defer-6}](#attr-script-defer)
(and not
[`async`{#the-script-element:attr-script-async-7}](#attr-script-async))
to fall back to the
[`defer`{#the-script-element:attr-script-defer-7}](#attr-script-defer)
behavior instead of the blocking behavior that is the default.

The [`blocking`]{#attr-script-blocking .dfn dfn-for="script"
dfn-type="element-attr"} attribute is a [blocking
attribute](urls-and-fetching.html#blocking-attribute){#the-script-element:blocking-attribute}.

The [`crossorigin`]{#attr-script-crossorigin .dfn dfn-for="script"
dfn-type="element-attr"} attribute is a [CORS settings
attribute](urls-and-fetching.html#cors-settings-attribute){#the-script-element:cors-settings-attribute}.
For external [classic
scripts](webappapis.html#classic-script){#the-script-element:classic-script-6},
it controls whether error information will be exposed, when the script
is obtained from other
[origins](browsers.html#concept-origin){#the-script-element:concept-origin}.
For external [module
scripts](webappapis.html#module-script){#the-script-element:module-script-5},
it controls the [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-script-element:concept-request-credentials-mode-2
x-internal="concept-request-credentials-mode"} used for the initial
fetch of the module source, if cross-origin. For both
[classic](webappapis.html#classic-script){#the-script-element:classic-script-7}
and [module
scripts](webappapis.html#module-script){#the-script-element:module-script-6},
it controls the [credentials
mode](https://fetch.spec.whatwg.org/#concept-request-credentials-mode){#the-script-element:concept-request-credentials-mode-3
x-internal="concept-request-credentials-mode"} used for cross-origin
module imports.

Unlike [classic
scripts](webappapis.html#classic-script){#the-script-element:classic-script-8},
[module
scripts](webappapis.html#module-script){#the-script-element:module-script-7}
require the use of the [CORS
protocol](https://fetch.spec.whatwg.org/#http-cors-protocol){#the-script-element:cors-protocol
x-internal="cors-protocol"} for cross-origin fetching.

The [`referrerpolicy`]{#attr-script-referrerpolicy .dfn dfn-for="script"
dfn-type="element-attr"} attribute is a [referrer policy
attribute](urls-and-fetching.html#referrer-policy-attribute){#the-script-element:referrer-policy-attribute}.
Its sets the [referrer
policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-script-element:referrer-policy-3
x-internal="referrer-policy"} used for the initial fetch of an external
script, as well as the fetching of any imported module scripts.
[\[REFERRERPOLICY\]](references.html#refsREFERRERPOLICY)

::: example
An example of a
[`script`{#the-script-element:the-script-element-8}](#the-script-element)
element\'s referrer policy being used when fetching imported scripts but
not other subresources:

``` html
<script referrerpolicy="origin">
  fetch('/api/data');    // not fetched with <script>'s referrer policy
  import('./utils.mjs'); // is fetched with <script>'s referrer policy ("origin" in this case)
</script>
```
:::

The [`integrity`]{#attr-script-integrity .dfn dfn-for="script"
dfn-type="element-attr"} attribute sets the [integrity
metadata](https://fetch.spec.whatwg.org/#concept-request-integrity-metadata){#the-script-element:concept-request-integrity-metadata
x-internal="concept-request-integrity-metadata"} used for the initial
fetch of an external script. The value must match [the requirements of
the integrity
attribute](https://w3c.github.io/webappsec-subresource-integrity/#the-integrity-attribute){#the-script-element:the-requirements-of-the-integrity-attribute
x-internal="the-requirements-of-the-integrity-attribute"}.
[\[SRI\]](references.html#refsSRI)

The [`fetchpriority`]{#attr-script-fetchpriority .dfn dfn-for="script"
dfn-type="element-attr"} attribute is a [fetch priority
attribute](urls-and-fetching.html#fetch-priority-attribute){#the-script-element:fetch-priority-attribute}.
Its sets the
[priority](https://fetch.spec.whatwg.org/#request-priority){#the-script-element:concept-request-priority-2
x-internal="concept-request-priority"} used for the initial fetch of an
external script.

Changing any of these attributes dynamically has no direct effect; these
attributes are only used at specific times described in the [processing
model](#script-processing-model).

------------------------------------------------------------------------

The IDL attributes [`src`]{#dom-script-src .dfn
dfn-for="HTMLScriptElement" dfn-type="attribute"},
[`type`]{#dom-script-type .dfn dfn-for="HTMLScriptElement"
dfn-type="attribute"}, [`defer`]{#dom-script-defer .dfn
dfn-for="HTMLScriptElement" dfn-type="attribute"},
[`integrity`]{#dom-script-integrity .dfn dfn-for="HTMLScriptElement"
dfn-type="attribute"}, and [`blocking`]{#dom-script-blocking .dfn
dfn-for="HTMLScriptElement" dfn-type="attribute"} must each
[reflect](common-dom-interfaces.html#reflect){#the-script-element:reflect}
the respective content attributes of the same name.

The [`noModule`]{#dom-script-nomodule .dfn dfn-for="HTMLScriptElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-script-element:reflect-2}
the
[`nomodule`{#the-script-element:attr-script-nomodule-3}](#attr-script-nomodule)
content attribute.

The [`crossOrigin`]{#dom-script-crossorigin .dfn
dfn-for="HTMLScriptElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-script-element:reflect-3}
the
[`crossorigin`{#the-script-element:attr-script-crossorigin-5}](#attr-script-crossorigin)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-script-element:limited-to-only-known-values}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLScriptElement/referrerPolicy](https://developer.mozilla.org/en-US/docs/Web/API/HTMLScriptElement/referrerPolicy "The referrerPolicy property of the HTMLScriptElement interface reflects the HTML referrerpolicy of the <script> element, which defines how the referrer is set when fetching the script and any scripts it imports.")

Support in all current engines.

::: support
[Firefox65+]{.firefox .yes}[Safari14+]{.safari .yes}[Chrome70+]{.chrome
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

The [`referrerPolicy`]{#dom-script-referrerpolicy .dfn
dfn-for="HTMLScriptElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-script-element:reflect-4}
the
[`referrerpolicy`{#the-script-element:attr-script-referrerpolicy-5}](#attr-script-referrerpolicy)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-script-element:limited-to-only-known-values-2}.

The [`fetchPriority`]{#dom-script-fetchpriority .dfn
dfn-for="HTMLScriptElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-script-element:reflect-5}
the
[`fetchpriority`{#the-script-element:attr-script-fetchpriority-4}](#attr-script-fetchpriority)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-script-element:limited-to-only-known-values-3}.

The [`async`]{#dom-script-async .dfn dfn-for="HTMLScriptElement"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-script-element:this
    x-internal="this"}\'s [force
    async](#script-force-async){#the-script-element:script-force-async}
    is true, then return true.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-script-element:this-2
    x-internal="this"}\'s
    [`async`{#the-script-element:attr-script-async-8}](#attr-script-async)
    content attribute is present, then return true.

3.  Return false.

The [`async`{#the-script-element:dom-script-async-2}](#dom-script-async)
setter steps are:

1.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-script-element:this-3
    x-internal="this"}\'s [force
    async](#script-force-async){#the-script-element:script-force-async-2}
    to false.

2.  If the given value is true, then set
    [this](https://webidl.spec.whatwg.org/#this){#the-script-element:this-4
    x-internal="this"}\'s
    [`async`{#the-script-element:attr-script-async-9}](#attr-script-async)
    content attribute to the empty string.

3.  Otherwise, remove
    [this](https://webidl.spec.whatwg.org/#this){#the-script-element:this-5
    x-internal="this"}\'s
    [`async`{#the-script-element:attr-script-async-10}](#attr-script-async)
    content attribute.

`script`{.variable}`.`[`text`](#dom-script-text){#dom-script-text-dev}` [ = ``value`{.variable}` ]`

:   Returns the [child text
    content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-script-element:child-text-content-2
    x-internal="child-text-content"} of the element.

`script`{.variable}`.`[`text`](#dom-script-text){#the-script-element:dom-script-text-2}` = ``value`{.variable}

:   Replaces the element\'s children with the text given by
    `value`{.variable}.

[`HTMLScriptElement`{#the-script-element:htmlscriptelement}](#htmlscriptelement)`.`[`supports`](#dom-script-supports){#dom-script-supports-dev}`(``type`{.variable}`)`

:   Returns true if the given `type`{.variable} is a script type
    supported by the user agent. The possible script types in this
    specification are \"`classic`\", \"`module`\", and \"`importmap`\",
    but others might be added in the future.

The [`text`]{#dom-script-text .dfn dfn-for="HTMLScriptElement"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-script-element:this-6
x-internal="this"}\'s [child text
content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-script-element:child-text-content-3
x-internal="child-text-content"}.

The [`text`{#the-script-element:dom-script-text-3}](#dom-script-text)
setter steps are to [string replace
all](https://dom.spec.whatwg.org/#string-replace-all){#the-script-element:string-replace-all
x-internal="string-replace-all"} with the given value within
[this](https://webidl.spec.whatwg.org/#this){#the-script-element:this-7
x-internal="this"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLScriptElement/supports_static](https://developer.mozilla.org/en-US/docs/Web/API/HTMLScriptElement/supports_static "The supports() static method of the HTMLScriptElement interface provides a simple and consistent method to feature-detect what types of scripts are supported by the user agent.")

Support in all current engines.

::: support
[Firefox94+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome96+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge96+]{.edge_blink .yes}

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

The static [`supports(``type`{.variable}`)`]{#dom-script-supports .dfn
dfn-for="HTMLScriptElement" dfn-type="method"} method steps are:

1.  If `type`{.variable}
    [is](https://infra.spec.whatwg.org/#string-is){#the-script-element:is
    x-internal="is"} \"`classic`\", then return true.

2.  If `type`{.variable}
    [is](https://infra.spec.whatwg.org/#string-is){#the-script-element:is-2
    x-internal="is"} \"`module`\", then return true.

3.  If `type`{.variable}
    [is](https://infra.spec.whatwg.org/#string-is){#the-script-element:is-3
    x-internal="is"} \"`importmap`\", then return true.

4.  Return false.

The `type`{.variable} argument has to exactly match these values; we do
not perform an [ASCII
case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#the-script-element:ascii-case-insensitive-3
x-internal="ascii-case-insensitive"} match. This is different from how
[`type`{#the-script-element:attr-script-type-6}](#attr-script-type)
content attribute values are treated, and how
[`DOMTokenList`{#the-script-element:domtokenlist-2}](https://dom.spec.whatwg.org/#interface-domtokenlist){x-internal="domtokenlist"}\'s
[`supports()`{#the-script-element:dom-domtokenlist-supports}](https://dom.spec.whatwg.org/#dom-domtokenlist-supports){x-internal="dom-domtokenlist-supports"}
method works, but it aligns with the
[`WorkerType`{#the-script-element:workertype}](workers.html#workertype)
enumeration used in the
[`Worker()`{#the-script-element:dom-worker}](workers.html#dom-worker)
constructor.

------------------------------------------------------------------------

::: example
In this example, two
[`script`{#the-script-element:the-script-element-9}](#the-script-element)
elements are used. One embeds an external [classic
script](webappapis.html#classic-script){#the-script-element:classic-script-9},
and the other includes some data as a [data
block](#data-block){#the-script-element:data-block-3}.

``` html
<script src="game-engine.js"></script>
<script type="text/x-game-map">
........U.........e
o............A....e
.....A.....AAA....e
.A..AAA...AAAAA...e
</script>
```

The data in this case might be used by the script to generate the map of
a video game. The data doesn\'t have to be used that way, though; maybe
the map data is actually embedded in other parts of the page\'s markup,
and the data block here is just used by the site\'s search engine to
help users who are looking for particular features in their game maps.
:::

::: example
The following sample shows how a
[`script`{#the-script-element:the-script-element-10}](#the-script-element)
element can be used to define a function that is then used by other
parts of the document, as part of a [classic
script](webappapis.html#classic-script){#the-script-element:classic-script-10}.
It also shows how a
[`script`{#the-script-element:the-script-element-11}](#the-script-element)
element can be used to invoke script while the document is being parsed,
in this case to initialize the form\'s output.

``` html
<script>
 function calculate(form) {
   var price = 52000;
   if (form.elements.brakes.checked)
     price += 1000;
   if (form.elements.radio.checked)
     price += 2500;
   if (form.elements.turbo.checked)
     price += 5000;
   if (form.elements.sticker.checked)
     price += 250;
   form.elements.result.value = price;
 }
</script>
<form name="pricecalc" onsubmit="return false" onchange="calculate(this)">
 <fieldset>
  <legend>Work out the price of your car</legend>
  <p>Base cost: £52000.</p>
  <p>Select additional options:</p>
  <ul>
   <li><label><input type=checkbox name=brakes> Ceramic brakes (£1000)</label></li>
   <li><label><input type=checkbox name=radio> Satellite radio (£2500)</label></li>
   <li><label><input type=checkbox name=turbo> Turbo charger (£5000)</label></li>
   <li><label><input type=checkbox name=sticker> "XZ" sticker (£250)</label></li>
  </ul>
  <p>Total: £<output name=result></output></p>
 </fieldset>
 <script>
  calculate(document.forms.pricecalc);
 </script>
</form>
```
:::

::: {#script-type-module-example-1 .example}
The following sample shows how a
[`script`{#the-script-element:the-script-element-12}](#the-script-element)
element can be used to include an external [JavaScript module
script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-4}.

``` html
<script type="module" src="app.mjs"></script>
```

This module, and all its dependencies (expressed through JavaScript
`import` statements in the source file), will be fetched. Once the
entire resulting module graph has been imported, and the document has
finished parsing, the contents of `app.mjs` will be evaluated.

Additionally, if code from another
[`script`{#the-script-element:the-script-element-13}](#the-script-element)
element in the same
[`Window`{#the-script-element:window}](nav-history-apis.html#window)
imports the module from `app.mjs` (e.g. via `import "./app.mjs";`), then
the same [JavaScript module
script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-5}
created by the former
[`script`{#the-script-element:the-script-element-14}](#the-script-element)
element will be imported.
:::

::: {#script-nomodule-example .example}
This example shows how to include a [JavaScript module
script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-6}
for modern user agents, and a [classic
script](webappapis.html#classic-script){#the-script-element:classic-script-11}
for older user agents:

``` html
<script type="module" src="app.mjs"></script>
<script nomodule defer src="classic-app-bundle.js"></script>
```

In modern user agents that support [JavaScript module
scripts](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-7},
the
[`script`{#the-script-element:the-script-element-15}](#the-script-element)
element with the
[`nomodule`{#the-script-element:attr-script-nomodule-4}](#attr-script-nomodule)
attribute will be ignored, and the
[`script`{#the-script-element:the-script-element-16}](#the-script-element)
element with a
[`type`{#the-script-element:attr-script-type-7}](#attr-script-type) of
\"`module`\" will be fetched and evaluated (as a [JavaScript module
script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-8}).
Conversely, older user agents will ignore the
[`script`{#the-script-element:the-script-element-17}](#the-script-element)
element with a
[`type`{#the-script-element:attr-script-type-8}](#attr-script-type) of
\"`module`\", as that is an unknown script type for them --- but they
will have no problem fetching and evaluating the other
[`script`{#the-script-element:the-script-element-18}](#the-script-element)
element (as a [classic
script](webappapis.html#classic-script){#the-script-element:classic-script-12}),
since they do not implement the
[`nomodule`{#the-script-element:attr-script-nomodule-5}](#attr-script-nomodule)
attribute.
:::

::: {#script-type-module-example-2 .example}
The following sample shows how a
[`script`{#the-script-element:the-script-element-19}](#the-script-element)
element can be used to write an inline [JavaScript module
script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-9}
that performs a number of substitutions on the document\'s text, in
order to make for a more interesting reading experience (e.g. on a news
site): [\[XKCD1288\]](references.html#refsXKCD1288)

``` html
<script type="module">
 import { walkAllTextNodeDescendants } from "./dom-utils.mjs";

 const substitutions = new Map([
   ["witnesses", "these dudes I know"]
   ["allegedly", "kinda probably"]
   ["new study", "Tumblr post"]
   ["rebuild", "avenge"]
   ["space", "spaaace"]
   ["Google glass", "Virtual Boy"]
   ["smartphone", "Pokédex"]
   ["electric", "atomic"]
   ["Senator", "Elf-Lord"]
   ["car", "cat"]
   ["election", "eating contest"]
   ["Congressional leaders", "river spirits"]
   ["homeland security", "Homestar Runner"]
   ["could not be reached for comment", "is guilty and everyone knows it"]
 ]);

 function substitute(textNode) {
   for (const [before, after] of substitutions.entries()) {
     textNode.data = textNode.data.replace(new RegExp(`\\b${before}\\b`, "ig"), after);
   }
 }

 walkAllTextNodeDescendants(document.body, substitute);
</script>
```

Some notable features gained by using a JavaScript module script include
the ability to import functions from other JavaScript modules, strict
mode by default, and how top-level declarations do not introduce new
properties onto the [global
object](webappapis.html#global-object){#the-script-element:global-object}.
Also note that no matter where this
[`script`{#the-script-element:the-script-element-20}](#the-script-element)
element appears in the document, it will not be evaluated until both
document parsing has complete and its dependency (`dom-utils.mjs`) has
been fetched and evaluated.
:::

::: {#json-module-script-example .example}
The following sample shows how a [JSON module
script](webappapis.html#json-module-script){#the-script-element:json-module-script}
can be imported from inside a [JavaScript module
script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-10}:

``` html
<script type="module">
 import peopleInSpace from "http://api.open-notify.org/astros.json" with { type: "json" };

 const list = document.querySelector("#people-in-space");
 for (const { craft, name } of peopleInSpace.people) {
   const li = document.createElement("li");
   li.textContent = `${name} / ${craft}`;
   list.append(li);
 }
</script>
```

MIME type checking for module scripts is strict. In order for the fetch
of the [JSON module
script](webappapis.html#json-module-script){#the-script-element:json-module-script-2}
to succeed, the HTTP response must have a [JSON MIME
type](https://mimesniff.spec.whatwg.org/#json-mime-type){#the-script-element:json-mime-type
x-internal="json-mime-type"}, for example `Content-Type: text/json`. On
the other hand, if the `with { type: "json" }` part of the statement is
omitted, it is assumed that the intent is to import a [JavaScript module
script](webappapis.html#javascript-module-script){#the-script-element:javascript-module-script-11},
and the fetch will fail if the HTTP response has a MIME type that is not
a [JavaScript MIME
type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#the-script-element:javascript-mime-type
x-internal="javascript-mime-type"}.
:::

##### [4.12.1.1]{.secno} Processing model[](#script-processing-model){.self-link} {#script-processing-model}

A
[`script`{#script-processing-model:the-script-element}](#the-script-element)
element has several associated pieces of state.

A
[`script`{#script-processing-model:the-script-element-2}](#the-script-element)
element has a [parser document]{#parser-document .dfn}, which is either
null or a
[`Document`{#script-processing-model:document}](dom.html#document),
initially null. It is set by the [HTML
parser](parsing.html#html-parser){#script-processing-model:html-parser}
and the [XML
parser](xhtml.html#xml-parser){#script-processing-model:xml-parser} on
[`script`{#script-processing-model:the-script-element-3}](#the-script-element)
elements they insert, and affects the processing of those elements.
[`script`{#script-processing-model:the-script-element-4}](#the-script-element)
elements with non-null [parser
documents](#parser-document){#script-processing-model:parser-document}
are known as [parser-inserted]{#parser-inserted .dfn dfn-for="script"
export=""}.

A
[`script`{#script-processing-model:the-script-element-5}](#the-script-element)
element has a [preparation-time document]{#preparation-time-document
.dfn}, which is either null or a
[`Document`{#script-processing-model:document-2}](dom.html#document),
initially null. It is used to prevent scripts that move between
documents during
[preparation](#prepare-the-script-element){#script-processing-model:prepare-the-script-element}
from
[executing](#execute-the-script-element){#script-processing-model:execute-the-script-element}.

A
[`script`{#script-processing-model:the-script-element-6}](#the-script-element)
element has a [force async]{#script-force-async .dfn} boolean, initially
true. It is set to false by the [HTML
parser](parsing.html#html-parser){#script-processing-model:html-parser-2}
and the [XML
parser](xhtml.html#xml-parser){#script-processing-model:xml-parser-2} on
[`script`{#script-processing-model:the-script-element-7}](#the-script-element)
elements they insert, and when the element gets an
[`async`{#script-processing-model:attr-script-async}](#attr-script-async)
content attribute added.

A
[`script`{#script-processing-model:the-script-element-8}](#the-script-element)
element has a [from an external file]{#concept-script-external .dfn}
boolean, initially false. It is determined when the script is
[prepared](#prepare-the-script-element){#script-processing-model:prepare-the-script-element-2},
based on the
[`src`{#script-processing-model:attr-script-src}](#attr-script-src)
attribute of the element at that time.

A
[`script`{#script-processing-model:the-script-element-9}](#the-script-element)
element has a [ready to be parser-executed]{#ready-to-be-parser-executed
.dfn} boolean, initially false. This is used only used for elements that
are also
[parser-inserted](#parser-inserted){#script-processing-model:parser-inserted},
to let the parser know when to execute the script.

A
[`script`{#script-processing-model:the-script-element-10}](#the-script-element)
element has an [already started]{#already-started .dfn} boolean,
initially false.

A
[`script`{#script-processing-model:the-script-element-11}](#the-script-element)
element has a [delaying the load event]{#concept-script-delay-load .dfn}
boolean, initially false.

A
[`script`{#script-processing-model:the-script-element-12}](#the-script-element)
element has a [type]{#concept-script-type .dfn}, which is either null,
\"`classic`\", \"`module`\", or \"`importmap`\", initially null. It is
determined when the element is
[prepared](#prepare-the-script-element){#script-processing-model:prepare-the-script-element-3},
based on the
[`type`{#script-processing-model:attr-script-type}](#attr-script-type)
attribute of the element at that time.

A
[`script`{#script-processing-model:the-script-element-13}](#the-script-element)
element has a [result]{#concept-script-result .dfn}, which is either
\"`uninitialized`\", null (representing an error), a
[script](webappapis.html#concept-script){#script-processing-model:concept-script},
or an [import map parse
result](webappapis.html#import-map-parse-result){#script-processing-model:import-map-parse-result}.
It is initially \"`uninitialized`\".

A
[`script`{#script-processing-model:the-script-element-14}](#the-script-element)
element has [steps to run when the result is
ready]{#steps-to-run-when-the-result-is-ready .dfn}, which are a series
of steps or null, initially null. To [mark as ready]{#mark-as-ready
.dfn} a
[`script`{#script-processing-model:the-script-element-15}](#the-script-element)
element `el`{.variable} given a
[script](webappapis.html#concept-script){#script-processing-model:concept-script-2},
[import map parse
result](webappapis.html#import-map-parse-result){#script-processing-model:import-map-parse-result-2},
or null `result`{.variable}:

1.  Set `el`{.variable}\'s
    [result](#concept-script-result){#script-processing-model:concept-script-result}
    to `result`{.variable}.

2.  If `el`{.variable}\'s [steps to run when the result is
    ready](#steps-to-run-when-the-result-is-ready){#script-processing-model:steps-to-run-when-the-result-is-ready}
    are not null, then run them.

3.  Set `el`{.variable}\'s [steps to run when the result is
    ready](#steps-to-run-when-the-result-is-ready){#script-processing-model:steps-to-run-when-the-result-is-ready-2}
    to null.

4.  Set `el`{.variable}\'s [delaying the load
    event](#concept-script-delay-load){#script-processing-model:concept-script-delay-load}
    to false.

------------------------------------------------------------------------

A
[`script`{#script-processing-model:the-script-element-16}](#the-script-element)
element `el`{.variable} is [implicitly potentially
render-blocking](urls-and-fetching.html#implicitly-potentially-render-blocking){#script-processing-model:implicitly-potentially-render-blocking}
if `el`{.variable}\'s
[type](#concept-script-type){#script-processing-model:concept-script-type}
is \"`classic`\", `el`{.variable} is
[parser-inserted](#parser-inserted){#script-processing-model:parser-inserted-2},
and `el`{.variable} does not have an
[`async`{#script-processing-model:attr-script-async-2}](#attr-script-async)
or
[`defer`{#script-processing-model:attr-script-defer}](#attr-script-defer)
attribute.

The [cloning
steps](https://dom.spec.whatwg.org/#concept-node-clone-ext){#script-processing-model:concept-node-clone-ext
x-internal="concept-node-clone-ext"} for
[`script`{#script-processing-model:the-script-element-17}](#the-script-element)
elements given `node`{.variable}, `copy`{.variable}, and
`subtree`{.variable} are to set `copy`{.variable}\'s [already
started](#already-started){#script-processing-model:already-started} to
`node`{.variable}\'s [already
started](#already-started){#script-processing-model:already-started-2}.

When an
[`async`{#script-processing-model:attr-script-async-3}](#attr-script-async)
attribute is added to a
[`script`{#script-processing-model:the-script-element-18}](#the-script-element)
element `el`{.variable}, the user agent must set `el`{.variable}\'s
[force
async](#script-force-async){#script-processing-model:script-force-async}
to false.

Whenever a
[`script`{#script-processing-model:the-script-element-19}](#the-script-element)
element `el`{.variable}\'s [delaying the load
event](#concept-script-delay-load){#script-processing-model:concept-script-delay-load-2}
is true, the user agent must [delay the load
event](parsing.html#delay-the-load-event){#script-processing-model:delay-the-load-event}
of `el`{.variable}\'s [preparation-time
document](#preparation-time-document){#script-processing-model:preparation-time-document}.

------------------------------------------------------------------------

The
[`script`{#script-processing-model:the-script-element-20}](#the-script-element)
[HTML element post-connection
steps](infrastructure.html#html-element-post-connection-steps){#script-processing-model:html-element-post-connection-steps},
given `insertedNode`{.variable}, are:

1.  If `insertedNode`{.variable} is not
    [connected](https://dom.spec.whatwg.org/#connected){#script-processing-model:connected
    x-internal="connected"}, then return.

    ::: example
    This can happen in the case where an earlier-inserted
    [`script`{#script-processing-model:the-script-element-21}](#the-script-element)
    removes a later-inserted
    [`script`{#script-processing-model:the-script-element-22}](#the-script-element).
    For instance:

    ``` html
    <script>
    const script1 = document.createElement('script');
    script1.innerText = `
      document.querySelector('#script2').remove();
    `;

    const script2 = document.createElement('script');
    script2.id = 'script2';
    script2.textContent = `console.log('script#2 running')`;

    document.body.append(script1, script2);
    </script>
    ```

    Nothing is printed to the console in this example. By the time the
    [HTML element post-connection
    steps](infrastructure.html#html-element-post-connection-steps){#script-processing-model:html-element-post-connection-steps-2}
    run for the first
    [`script`{#script-processing-model:the-script-element-23}](#the-script-element)
    that was atomically inserted by
    [`append()`{#script-processing-model:dom-node-append}](https://dom.spec.whatwg.org/#dom-node-append){x-internal="dom-node-append"},
    it can observe that the second
    [`script`{#script-processing-model:the-script-element-24}](#the-script-element)
    is already
    [connected](https://dom.spec.whatwg.org/#connected){#script-processing-model:connected-2
    x-internal="connected"} to the DOM. It removes the second
    [`script`{#script-processing-model:the-script-element-25}](#the-script-element),
    so that by the time *its* [HTML element post-connection
    steps](infrastructure.html#html-element-post-connection-steps){#script-processing-model:html-element-post-connection-steps-3}
    run, it is no longer
    [connected](https://dom.spec.whatwg.org/#connected){#script-processing-model:connected-3
    x-internal="connected"}, and does not get
    [prepared](#prepare-the-script-element){#script-processing-model:prepare-the-script-element-4}.
    :::

2.  If `insertedNode`{.variable} is
    [parser-inserted](#parser-inserted){#script-processing-model:parser-inserted-3},
    then return.

3.  [Prepare the script
    element](#prepare-the-script-element){#script-processing-model:prepare-the-script-element-5}
    given `insertedNode`{.variable}.

The
[`script`{#script-processing-model:the-script-element-26}](#the-script-element)
[children changed
steps](https://dom.spec.whatwg.org/#concept-node-children-changed-ext){#script-processing-model:children-changed-steps
x-internal="children-changed-steps"} are:

1.  Run the
    [`script`{#script-processing-model:the-script-element-27}](#the-script-element)
    [HTML element post-connection
    steps](infrastructure.html#html-element-post-connection-steps){#script-processing-model:html-element-post-connection-steps-4},
    given the
    [`script`{#script-processing-model:the-script-element-28}](#the-script-element)
    element.

::: example
This has an interesting implication on the execution order of a
[`script`{#script-processing-model:the-script-element-29}](#the-script-element)
element and any newly-inserted child
[`script`{#script-processing-model:the-script-element-30}](#the-script-element)
elements. Consider the following snippet:

``` html
<script id=outer-script></script>

<script>
  const outerScript = document.querySelector('#outer-script');

  const start = new Text('console.log(1);');
  const innerScript = document.createElement('script');
  innerScript.textContent = `console.log('inner script executing')`;
  const end = new Text('console.log(2);');

  outerScript.append(start, innerScript, end);

  // Logs:
  // 1
  // 2
  // inner script executing
</script>
```

By the time the second script block executes, the `outer-script` has
already been
[prepared](#prepare-the-script-element){#script-processing-model:prepare-the-script-element-6},
but because it is empty, it did not execute and therefore is not marked
as [already
started](#already-started){#script-processing-model:already-started-3}.
The atomic insertion of the
[`Text`{#script-processing-model:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes and nested
[`script`{#script-processing-model:the-script-element-31}](#the-script-element)
element have the following effects:

1.  All three child nodes get atomically inserted as children of
    `outer-script`; all of their [insertion
    steps](https://dom.spec.whatwg.org/#concept-node-insert-ext){#script-processing-model:concept-node-insert-ext
    x-internal="concept-node-insert-ext"} run, which have no observable
    consequences in this case.

2.  The `outer-script`\'s [children changed
    steps](https://dom.spec.whatwg.org/#concept-node-children-changed-ext){#script-processing-model:children-changed-steps-2
    x-internal="children-changed-steps"} run, which
    [prepares](#prepare-the-script-element){#script-processing-model:prepare-the-script-element-7}
    that script; because its body is now non-empty, this executes the
    contents of the two
    [`Text`{#script-processing-model:text-2}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
    nodes, in order.

3.  The
    [`script`{#script-processing-model:the-script-element-32}](#the-script-element)
    [HTML element post-connection
    steps](infrastructure.html#html-element-post-connection-steps){#script-processing-model:html-element-post-connection-steps-5}
    finally run for `innerScript`, causing its body to execute.
:::

The following [attribute change
steps](https://dom.spec.whatwg.org/#concept-element-attributes-change-ext){#script-processing-model:concept-element-attributes-change-ext
x-internal="concept-element-attributes-change-ext"}, given
`element`{.variable}, `localName`{.variable}, `oldValue`{.variable},
`value`{.variable}, and `namespace`{.variable}, are used for all
[`script`{#script-processing-model:the-script-element-33}](#the-script-element)
elements:

1.  If `namespace`{.variable} is not null, then return.

2.  If `localName`{.variable} is
    [`src`{#script-processing-model:attr-script-src-2}](#attr-script-src),
    then run the
    [`script`{#script-processing-model:the-script-element-34}](#the-script-element)
    [HTML element post-connection
    steps](infrastructure.html#html-element-post-connection-steps){#script-processing-model:html-element-post-connection-steps-6},
    given `element`{.variable}.

To [prepare the script element]{#prepare-the-script-element .dfn} given
a
[`script`{#script-processing-model:the-script-element-35}](#the-script-element)
element `el`{.variable}:

1.  If `el`{.variable}\'s [already
    started](#already-started){#script-processing-model:already-started-4}
    is true, then return.

2.  Let `parser document`{.variable} be `el`{.variable}\'s [parser
    document](#parser-document){#script-processing-model:parser-document-2}.

3.  Set `el`{.variable}\'s [parser
    document](#parser-document){#script-processing-model:parser-document-3}
    to null.

    This is done so that if parser-inserted
    [`script`{#script-processing-model:the-script-element-36}](#the-script-element)
    elements fail to run when the parser tries to run them, e.g. because
    they are empty or specify an unsupported scripting language, another
    script can later mutate them and cause them to run again.

4.  If `parser document`{.variable} is non-null and `el`{.variable} does
    not have an
    [`async`{#script-processing-model:attr-script-async-4}](#attr-script-async)
    attribute, then set `el`{.variable}\'s [force
    async](#script-force-async){#script-processing-model:script-force-async-2}
    to true.

    This is done so that if a parser-inserted
    [`script`{#script-processing-model:the-script-element-37}](#the-script-element)
    element fails to run when the parser tries to run it, but it is
    later executed after a script dynamically updates it, it will
    execute in an async fashion even if the
    [`async`{#script-processing-model:attr-script-async-5}](#attr-script-async)
    attribute isn\'t set.

5.  Let `source text`{.variable} be `el`{.variable}\'s [child text
    content](https://dom.spec.whatwg.org/#concept-child-text-content){#script-processing-model:child-text-content
    x-internal="child-text-content"}.

6.  [If `el`{.variable} has no
    [`src`{#script-processing-model:attr-script-src-3}](#attr-script-src)
    attribute, and `source text`{.variable} is the empty string, then
    return.]{#script-processing-empty}

7.  If `el`{.variable} is not
    [connected](https://dom.spec.whatwg.org/#connected){#script-processing-model:connected-4
    x-internal="connected"}, then return.

8.  ::: {#script-processing-prepare}
    If any of the following are true:

    - `el`{.variable} has a
      [`type`{#script-processing-model:attr-script-type-2}](#attr-script-type)
      attribute whose value is the empty string;

    - `el`{.variable} has no
      [`type`{#script-processing-model:attr-script-type-3}](#attr-script-type)
      attribute but it has a
      [`language`{#script-processing-model:attr-script-language}](obsolete.html#attr-script-language)
      attribute and *that* attribute\'s value is the empty string; or

    - `el`{.variable} has neither a
      [`type`{#script-processing-model:attr-script-type-4}](#attr-script-type)
      attribute nor a
      [`language`{#script-processing-model:attr-script-language-2}](obsolete.html#attr-script-language)
      attribute,

    then let `the script block's type string`{.variable} for this
    [`script`{#script-processing-model:the-script-element-38}](#the-script-element)
    element be \"`text/javascript`\".

    Otherwise, if `el`{.variable} has a
    [`type`{#script-processing-model:attr-script-type-5}](#attr-script-type)
    attribute, then let `the script block's type string`{.variable} be
    the value of that attribute with [leading and trailing ASCII
    whitespace
    stripped](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#script-processing-model:strip-leading-and-trailing-ascii-whitespace
    x-internal="strip-leading-and-trailing-ascii-whitespace"}.

    Otherwise, `el`{.variable} has a non-empty
    [`language`{#script-processing-model:attr-script-language-3}](obsolete.html#attr-script-language)
    attribute; let `the script block's type string`{.variable} be the
    concatenation of \"`text/`\" and the value of `el`{.variable}\'s
    [`language`{#script-processing-model:attr-script-language-4}](obsolete.html#attr-script-language)
    attribute.

    The
    [`language`{#script-processing-model:attr-script-language-5}](obsolete.html#attr-script-language)
    attribute is never conforming, and is always ignored if there is a
    [`type`{#script-processing-model:attr-script-type-6}](#attr-script-type)
    attribute present.
    :::

9.  If `the script block's type string`{.variable} is a [JavaScript MIME
    type essence
    match](https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match){#script-processing-model:javascript-mime-type-essence-match
    x-internal="javascript-mime-type-essence-match"}, then set
    `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-2}
    to \"`classic`\".

10. Otherwise, if `the script block's type string`{.variable} is an
    [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#script-processing-model:ascii-case-insensitive
    x-internal="ascii-case-insensitive"} match for the string
    \"`module`\", then set `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-3}
    to \"`module`\".

11. Otherwise, if `the script block's type string`{.variable} is an
    [ASCII
    case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#script-processing-model:ascii-case-insensitive-2
    x-internal="ascii-case-insensitive"} match for the string
    \"`importmap`\", then set `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-4}
    to \"`importmap`\".

12. Otherwise, return. (No script is executed, and `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-5}
    is left as null.)

13. If `parser document`{.variable} is non-null, then set
    `el`{.variable}\'s [parser
    document](#parser-document){#script-processing-model:parser-document-4}
    back to `parser document`{.variable} and set `el`{.variable}\'s
    [force
    async](#script-force-async){#script-processing-model:script-force-async-3}
    to false.

14. [Set `el`{.variable}\'s [already
    started](#already-started){#script-processing-model:already-started-5}
    to true.]{#script-processing-start}

15. Set `el`{.variable}\'s [preparation-time
    document](#preparation-time-document){#script-processing-model:preparation-time-document-2}
    to its [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#script-processing-model:node-document
    x-internal="node-document"}.

16. If `parser document`{.variable} is non-null, and
    `parser document`{.variable} is not equal to `el`{.variable}\'s
    [preparation-time
    document](#preparation-time-document){#script-processing-model:preparation-time-document-3},
    then return.

17. ::: {#script-processing-noscript}
    If [scripting is
    disabled](webappapis.html#concept-n-noscript){#script-processing-model:concept-n-noscript}
    for `el`{.variable}, then return.

    The definition of [scripting is
    disabled](webappapis.html#concept-n-noscript){#script-processing-model:concept-n-noscript-2}
    means that, amongst others, the following scripts will not execute:
    scripts in
    [`XMLHttpRequest`{#script-processing-model:xmlhttprequest}](https://xhr.spec.whatwg.org/#xmlhttprequest){x-internal="xmlhttprequest"}\'s
    [`responseXML`{#script-processing-model:dom-xmlhttprequest-responsexml}](https://xhr.spec.whatwg.org/#dom-xmlhttprequest-responsexml){x-internal="dom-xmlhttprequest-responsexml"}
    documents, scripts in
    [`DOMParser`{#script-processing-model:domparser}](dynamic-markup-insertion.html#domparser)-created
    documents, scripts in documents created by
    [`XSLTProcessor`{#script-processing-model:xsltprocessor}](infrastructure.html#xsltprocessor)\'s
    [`transformToDocument`{#script-processing-model:dom-xsltprocessor-transformtodocument}](infrastructure.html#dom-xsltprocessor-transformtodocument)
    feature, and scripts that are first inserted by a script into a
    [`Document`{#script-processing-model:document-3}](dom.html#document)
    that was created using the
    [`createDocument()`{#script-processing-model:dom-domimplementation-createdocument}](https://dom.spec.whatwg.org/#dom-domimplementation-createdocument){x-internal="dom-domimplementation-createdocument"}
    API. [\[XHR\]](references.html#refsXHR)
    [\[DOMPARSING\]](references.html#refsDOMPARSING)
    [\[XSLTP\]](references.html#refsXSLTP)
    [\[DOM\]](references.html#refsDOM)
    :::

18. If `el`{.variable} has a
    [`nomodule`{#script-processing-model:attr-script-nomodule}](#attr-script-nomodule)
    content attribute and its
    [type](#concept-script-type){#script-processing-model:concept-script-type-6}
    is \"`classic`\", then return.

    This means specifying
    [`nomodule`{#script-processing-model:attr-script-nomodule-2}](#attr-script-nomodule)
    on a [module
    script](webappapis.html#module-script){#script-processing-model:module-script}
    has no effect; the algorithm continues onward.

19. [If `el`{.variable} does not have a
    [`src`{#script-processing-model:attr-script-src-4}](#attr-script-src)
    content attribute, and the [Should element\'s inline behavior be
    blocked by Content Security
    Policy?](https://w3c.github.io/webappsec-csp/#should-block-inline){#script-processing-model:should-element's-inline-behavior-be-blocked-by-content-security-policy
    x-internal="should-element's-inline-behavior-be-blocked-by-content-security-policy"}
    algorithm returns \"`Blocked`\" when given `el`{.variable},
    \"`script`\", and `source text`{.variable}, then return.
    [\[CSP\]](references.html#refsCSP)]{#script-processing-csp}

20. ::: {#script-processing-for}
    If `el`{.variable} has an
    [`event`{#script-processing-model:attr-script-event}](obsolete.html#attr-script-event)
    attribute and a
    [`for`{#script-processing-model:attr-script-for}](obsolete.html#attr-script-for)
    attribute, and `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-7}
    is \"`classic`\", then:

    1.  Let `for`{.variable} be the value of `el`{.variable}\'s
        [`for`{#script-processing-model:attr-script-for-2}](obsolete.html#attr-script-for)
        attribute.

    2.  Let `event`{.variable} be the value of `el`{.variable}\'s
        [`event`{#script-processing-model:attr-script-event-2}](obsolete.html#attr-script-event)
        attribute.

    3.  [Strip leading and trailing ASCII
        whitespace](https://infra.spec.whatwg.org/#strip-leading-and-trailing-ascii-whitespace){#script-processing-model:strip-leading-and-trailing-ascii-whitespace-2
        x-internal="strip-leading-and-trailing-ascii-whitespace"} from
        `event`{.variable} and `for`{.variable}.

    4.  If `for`{.variable} is not an [ASCII
        case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#script-processing-model:ascii-case-insensitive-3
        x-internal="ascii-case-insensitive"} match for the string
        \"`window`\", then return.

    5.  If `event`{.variable} is not an [ASCII
        case-insensitive](https://infra.spec.whatwg.org/#ascii-case-insensitive){#script-processing-model:ascii-case-insensitive-4
        x-internal="ascii-case-insensitive"} match for either the string
        \"`onload`\" or the string \"`onload()`\", then return.
    :::

21. ::: {#script-processing-encoding}
    If `el`{.variable} has a
    [`charset`{#script-processing-model:attr-script-charset}](obsolete.html#attr-script-charset)
    attribute, then let `encoding`{.variable} be the result of [getting
    an
    encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#script-processing-model:getting-an-encoding
    x-internal="getting-an-encoding"} from the value of the
    [`charset`{#script-processing-model:attr-script-charset-2}](obsolete.html#attr-script-charset)
    attribute.

    If `el`{.variable} does not have a
    [`charset`{#script-processing-model:attr-script-charset-3}](obsolete.html#attr-script-charset)
    attribute, or if [getting an
    encoding](https://encoding.spec.whatwg.org/#concept-encoding-get){#script-processing-model:getting-an-encoding-2
    x-internal="getting-an-encoding"} failed, then let
    `encoding`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#script-processing-model:node-document-2
    x-internal="node-document"}\'s [the
    encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#script-processing-model:document's-character-encoding
    x-internal="document's-character-encoding"}.

    If `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-8}
    is \"`module`\", this encoding will be ignored.
    :::

22. Let `classic script CORS setting`{.variable} be the current state of
    `el`{.variable}\'s
    [`crossorigin`{#script-processing-model:attr-script-crossorigin}](#attr-script-crossorigin)
    content attribute.

23. Let `module script credentials mode`{.variable} be the [CORS
    settings attribute credentials
    mode](urls-and-fetching.html#cors-settings-attribute-credentials-mode){#script-processing-model:cors-settings-attribute-credentials-mode}
    for `el`{.variable}\'s
    [`crossorigin`{#script-processing-model:attr-script-crossorigin-2}](#attr-script-crossorigin)
    content attribute.

24. Let `cryptographic nonce`{.variable} be `el`{.variable}\'s
    [\[\[CryptographicNonce\]\]](urls-and-fetching.html#cryptographicnonce){#script-processing-model:cryptographicnonce}
    internal slot\'s value.

25. If `el`{.variable} has an
    [`integrity`{#script-processing-model:attr-script-integrity}](#attr-script-integrity)
    attribute, then let `integrity metadata`{.variable} be that
    attribute\'s value.

    Otherwise, let `integrity metadata`{.variable} be the empty string.

26. Let `referrer policy`{.variable} be the current state of
    `el`{.variable}\'s
    [`referrerpolicy`{#script-processing-model:attr-script-referrerpolicy}](#attr-script-referrerpolicy)
    content attribute.

27. Let `fetch priority`{.variable} be the current state of
    `el`{.variable}\'s
    [`fetchpriority`{#script-processing-model:attr-script-fetchpriority}](#attr-script-fetchpriority)
    content attribute.

28. Let `parser metadata`{.variable} be \"`parser-inserted`\" if
    `el`{.variable} is
    [parser-inserted](#parser-inserted){#script-processing-model:parser-inserted-4},
    and \"`not-parser-inserted`\" otherwise.

29. Let `options`{.variable} be a [script fetch
    options](webappapis.html#script-fetch-options){#script-processing-model:script-fetch-options}
    whose [cryptographic
    nonce](webappapis.html#concept-script-fetch-options-nonce){#script-processing-model:concept-script-fetch-options-nonce}
    is `cryptographic nonce`{.variable}, [integrity
    metadata](webappapis.html#concept-script-fetch-options-integrity){#script-processing-model:concept-script-fetch-options-integrity}
    is `integrity metadata`{.variable}, [parser
    metadata](webappapis.html#concept-script-fetch-options-parser){#script-processing-model:concept-script-fetch-options-parser}
    is `parser metadata`{.variable}, [credentials
    mode](webappapis.html#concept-script-fetch-options-credentials){#script-processing-model:concept-script-fetch-options-credentials}
    is `module script credentials mode`{.variable}, [referrer
    policy](webappapis.html#concept-script-fetch-options-referrer-policy){#script-processing-model:concept-script-fetch-options-referrer-policy}
    is `referrer policy`{.variable}, and [fetch
    priority](webappapis.html#concept-script-fetch-options-fetch-priority){#script-processing-model:concept-script-fetch-options-fetch-priority}
    is `fetch priority`{.variable}.

30. Let `settings object`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#script-processing-model:node-document-3
    x-internal="node-document"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#script-processing-model:relevant-settings-object}.

31. ::: {#script-processing-src-prepare}
    If `el`{.variable} has a
    [`src`{#script-processing-model:attr-script-src-5}](#attr-script-src)
    content attribute, then:

    1.  If `el`{.variable}\'s
        [type](#concept-script-type){#script-processing-model:concept-script-type-9}
        is \"`importmap`\", then [queue an element
        task](webappapis.html#queue-an-element-task){#script-processing-model:queue-an-element-task}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#script-processing-model:dom-manipulation-task-source}
        given `el`{.variable} to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#script-processing-model:concept-event-fire
        x-internal="concept-event-fire"} named
        [`error`{#script-processing-model:event-error}](indices.html#event-error)
        at `el`{.variable}, and return.

        External import map scripts are not currently supported. See
        [WICG/import-maps issue
        #235](https://github.com/WICG/import-maps/issues/235) for
        discussions on adding support.

    2.  Let `src`{.variable} be the value of `el`{.variable}\'s
        [`src`{#script-processing-model:attr-script-src-6}](#attr-script-src)
        attribute.

    3.  If `src`{.variable} is the empty string, then [queue an element
        task](webappapis.html#queue-an-element-task){#script-processing-model:queue-an-element-task-2}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#script-processing-model:dom-manipulation-task-source-2}
        given `el`{.variable} to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#script-processing-model:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`error`{#script-processing-model:event-error-2}](indices.html#event-error)
        at `el`{.variable}, and return.

    4.  Set `el`{.variable}\'s [from an external
        file](#concept-script-external){#script-processing-model:concept-script-external}
        to true.

    5.  Let `url`{.variable} be the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#script-processing-model:encoding-parsing-a-url}
        given `src`{.variable}, relative to `el`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#script-processing-model:node-document-4
        x-internal="node-document"}.

    6.  If `url`{.variable} is failure, then [queue an element
        task](webappapis.html#queue-an-element-task){#script-processing-model:queue-an-element-task-3}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#script-processing-model:dom-manipulation-task-source-3}
        given `el`{.variable} to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#script-processing-model:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`error`{#script-processing-model:event-error-3}](indices.html#event-error)
        at `el`{.variable}, and return.

    7.  If `el`{.variable} is [potentially
        render-blocking](urls-and-fetching.html#potentially-render-blocking){#script-processing-model:potentially-render-blocking},
        then [block
        rendering](dom.html#block-rendering){#script-processing-model:block-rendering}
        on `el`{.variable}.

    8.  Set `el`{.variable}\'s [delaying the load
        event](#concept-script-delay-load){#script-processing-model:concept-script-delay-load-3}
        to true.

    9.  If `el`{.variable} is currently
        [render-blocking](dom.html#render-blocking){#script-processing-model:render-blocking},
        then set `options`{.variable}\'s
        [render-blocking](webappapis.html#concept-script-fetch-options-render-blocking){#script-processing-model:concept-script-fetch-options-render-blocking}
        to true.

    10. Let `onComplete`{.variable} given `result`{.variable} be the
        following steps:

        1.  [Mark as
            ready](#mark-as-ready){#script-processing-model:mark-as-ready}
            `el`{.variable} given `result`{.variable}.

    11. Switch on `el`{.variable}\'s
        [type](#concept-script-type){#script-processing-model:concept-script-type-10}:

        \"`classic`\"

        :   [Fetch a classic
            script](webappapis.html#fetch-a-classic-script){#script-processing-model:fetch-a-classic-script}
            given `url`{.variable}, `settings object`{.variable},
            `options`{.variable},
            `classic script CORS setting`{.variable},
            `encoding`{.variable}, and `onComplete`{.variable}.

        \"`module`\"

        :   If `el`{.variable} does not have an
            [`integrity`{#script-processing-model:attr-script-integrity-2}](#attr-script-integrity)
            attribute, then set `options`{.variable}\'s [integrity
            metadata](webappapis.html#concept-script-fetch-options-integrity){#script-processing-model:concept-script-fetch-options-integrity-2}
            to the result of [resolving a module integrity
            metadata](webappapis.html#resolving-a-module-integrity-metadata){#script-processing-model:resolving-a-module-integrity-metadata}
            with `url`{.variable} and `settings object`{.variable}.

            [Fetch an external module script
            graph](webappapis.html#fetch-a-module-script-tree){#script-processing-model:fetch-a-module-script-tree}
            given `url`{.variable}, `settings object`{.variable},
            `options`{.variable}, and `onComplete`{.variable}.

        For performance reasons, user agents may start fetching the
        classic script or module graph (as defined above) as soon as the
        [`src`{#script-processing-model:attr-script-src-7}](#attr-script-src)
        attribute is set, instead, in the hope that `el`{.variable} will
        become connected (and that the
        [`crossorigin`{#script-processing-model:attr-script-crossorigin-3}](#attr-script-crossorigin)
        attribute won\'t change value in the meantime). Either way, once
        `el`{.variable} [becomes
        connected](infrastructure.html#becomes-connected){#script-processing-model:becomes-connected},
        the load must have started as described in this step. If the UA
        performs such prefetching, but `el`{.variable} never becomes
        connected, or the
        [`src`{#script-processing-model:attr-script-src-8}](#attr-script-src)
        attribute is dynamically changed, or the
        [`crossorigin`{#script-processing-model:attr-script-crossorigin-4}](#attr-script-crossorigin)
        attribute is dynamically changed, then the user agent will not
        execute the script so obtained, and the fetching process will
        have been effectively wasted.
    :::

32. ::: {#establish-script-block-source}
    If `el`{.variable} does not have a
    [`src`{#script-processing-model:attr-script-src-9}](#attr-script-src)
    content attribute:

    1.  Let `base URL`{.variable} be `el`{.variable}\'s [node
        document](https://dom.spec.whatwg.org/#concept-node-document){#script-processing-model:node-document-5
        x-internal="node-document"}\'s [document base
        URL](urls-and-fetching.html#document-base-url){#script-processing-model:document-base-url}.

    2.  Switch on `el`{.variable}\'s
        [type](#concept-script-type){#script-processing-model:concept-script-type-11}:

        \"`classic`\"

        :   1.  Let `script`{.variable} be the result of [creating a
                classic
                script](webappapis.html#creating-a-classic-script){#script-processing-model:creating-a-classic-script}
                using `source text`{.variable},
                `settings object`{.variable}, `base URL`{.variable}, and
                `options`{.variable}.

            2.  [Mark as
                ready](#mark-as-ready){#script-processing-model:mark-as-ready-2}
                `el`{.variable} given `script`{.variable}.

        \"`module`\"

        :   1.  Set `el`{.variable}\'s [delaying the load
                event](#concept-script-delay-load){#script-processing-model:concept-script-delay-load-4}
                to true.

            2.  If `el`{.variable} is [potentially
                render-blocking](urls-and-fetching.html#potentially-render-blocking){#script-processing-model:potentially-render-blocking-2},
                then:

                1.  [Block
                    rendering](dom.html#block-rendering){#script-processing-model:block-rendering-2}
                    on `el`{.variable}.

                2.  Set `options`{.variable}\'s
                    [render-blocking](webappapis.html#concept-script-fetch-options-render-blocking){#script-processing-model:concept-script-fetch-options-render-blocking-2}
                    to true.

            3.  [Fetch an inline module script
                graph](webappapis.html#fetch-an-inline-module-script-graph){#script-processing-model:fetch-an-inline-module-script-graph},
                given `source text`{.variable}, `base URL`{.variable},
                `settings object`{.variable}, `options`{.variable}, and
                with the following steps given `result`{.variable}:

                1.  [Queue an element
                    task](webappapis.html#queue-an-element-task){#script-processing-model:queue-an-element-task-4}
                    on the [networking task
                    source](webappapis.html#networking-task-source){#script-processing-model:networking-task-source}
                    given `el`{.variable} to perform the following
                    steps:

                    1.  [Mark as
                        ready](#mark-as-ready){#script-processing-model:mark-as-ready-3}
                        `el`{.variable} given `result`{.variable}.

                    Queueing a task here means that, even if the inline
                    module script has no dependencies or synchronously
                    results in a parse error, we won\'t proceed to
                    [execute the script
                    element](#execute-the-script-element){#script-processing-model:execute-the-script-element-2}
                    synchronously.

        \"`importmap`\"

        :   1.  Let `result`{.variable} be the result of [creating an
                import map parse
                result](webappapis.html#create-an-import-map-parse-result){#script-processing-model:create-an-import-map-parse-result}
                given `source text`{.variable} and
                `base URL`{.variable}.

            2.  [Mark as
                ready](#mark-as-ready){#script-processing-model:mark-as-ready-4}
                `el`{.variable} given `result`{.variable}.
    :::

33. If `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-12}
    is \"`classic`\" and `el`{.variable} has a
    [`src`{#script-processing-model:attr-script-src-10}](#attr-script-src)
    attribute, or `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-13}
    is \"`module`\":

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#script-processing-model:assert
        x-internal="assert"}: `el`{.variable}\'s
        [result](#concept-script-result){#script-processing-model:concept-script-result-2}
        is \"`uninitialized`\".

    2.  ::: {#script-processing-src}
        If `el`{.variable} has an
        [`async`{#script-processing-model:attr-script-async-6}](#attr-script-async)
        attribute or `el`{.variable}\'s [force
        async](#script-force-async){#script-processing-model:script-force-async-4}
        is true:

        1.  Let `scripts`{.variable} be `el`{.variable}\'s
            [preparation-time
            document](#preparation-time-document){#script-processing-model:preparation-time-document-4}\'s
            [set of scripts that will execute as soon as
            possible](#set-of-scripts-that-will-execute-as-soon-as-possible){#script-processing-model:set-of-scripts-that-will-execute-as-soon-as-possible}.

        2.  [Append](https://infra.spec.whatwg.org/#set-append){#script-processing-model:set-append
            x-internal="set-append"} `el`{.variable} to
            `scripts`{.variable}.

        3.  Set `el`{.variable}\'s [steps to run when the result is
            ready](#steps-to-run-when-the-result-is-ready){#script-processing-model:steps-to-run-when-the-result-is-ready-3}
            to the following:

            1.  [Execute the script
                element](#execute-the-script-element){#script-processing-model:execute-the-script-element-3}
                `el`{.variable}.

            2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#script-processing-model:list-remove
                x-internal="list-remove"} `el`{.variable} from
                `scripts`{.variable}.
        :::

    3.  ::: {#script-processing-src-sync}
        Otherwise, if `el`{.variable} is not
        [parser-inserted](#parser-inserted){#script-processing-model:parser-inserted-5}:

        1.  Let `scripts`{.variable} be `el`{.variable}\'s
            [preparation-time
            document](#preparation-time-document){#script-processing-model:preparation-time-document-5}\'s
            [list of scripts that will execute in order as soon as
            possible](#list-of-scripts-that-will-execute-in-order-as-soon-as-possible){#script-processing-model:list-of-scripts-that-will-execute-in-order-as-soon-as-possible}.

        2.  [Append](https://infra.spec.whatwg.org/#list-append){#script-processing-model:list-append
            x-internal="list-append"} `el`{.variable} to
            `scripts`{.variable}.

        3.  Set `el`{.variable}\'s [steps to run when the result is
            ready](#steps-to-run-when-the-result-is-ready){#script-processing-model:steps-to-run-when-the-result-is-ready-4}
            to the following:

            1.  If `scripts`{.variable}\[0\] is not `el`{.variable},
                then abort these steps.

            2.  While `scripts`{.variable} is not empty, and
                `scripts`{.variable}\[0\]\'s
                [result](#concept-script-result){#script-processing-model:concept-script-result-3}
                is not \"`uninitialized`\":

                1.  [Execute the script
                    element](#execute-the-script-element){#script-processing-model:execute-the-script-element-4}
                    `scripts`{.variable}\[0\].

                2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#script-processing-model:list-remove-2
                    x-internal="list-remove"} `scripts`{.variable}\[0\].
        :::

    4.  ::: {#script-processing-defer}
        Otherwise, if `el`{.variable} has a
        [`defer`{#script-processing-model:attr-script-defer-2}](#attr-script-defer)
        attribute or `el`{.variable}\'s
        [type](#concept-script-type){#script-processing-model:concept-script-type-14}
        is \"`module`\":

        1.  [Append](https://infra.spec.whatwg.org/#list-append){#script-processing-model:list-append-2
            x-internal="list-append"} `el`{.variable} to its [parser
            document](#parser-document){#script-processing-model:parser-document-5}\'s
            [list of scripts that will execute when the document has
            finished
            parsing](#list-of-scripts-that-will-execute-when-the-document-has-finished-parsing){#script-processing-model:list-of-scripts-that-will-execute-when-the-document-has-finished-parsing}.

        2.  Set `el`{.variable}\'s [steps to run when the result is
            ready](#steps-to-run-when-the-result-is-ready){#script-processing-model:steps-to-run-when-the-result-is-ready-5}
            to the following: set `el`{.variable}\'s [ready to be
            parser-executed](#ready-to-be-parser-executed){#script-processing-model:ready-to-be-parser-executed}
            to true. (The parser will handle executing the script.)
        :::

    5.  ::: {#script-processing-parser-inserted}
        Otherwise:

        1.  Set `el`{.variable}\'s [parser
            document](#parser-document){#script-processing-model:parser-document-6}\'s
            [pending parsing-blocking
            script](#pending-parsing-blocking-script){#script-processing-model:pending-parsing-blocking-script}
            to `el`{.variable}.

        2.  [Block
            rendering](dom.html#block-rendering){#script-processing-model:block-rendering-3}
            on `el`{.variable}.

        3.  Set `el`{.variable}\'s [steps to run when the result is
            ready](#steps-to-run-when-the-result-is-ready){#script-processing-model:steps-to-run-when-the-result-is-ready-6}
            to the following: set `el`{.variable}\'s [ready to be
            parser-executed](#ready-to-be-parser-executed){#script-processing-model:ready-to-be-parser-executed-2}
            to true. (The parser will handle executing the script.)
        :::

34. Otherwise:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#script-processing-model:assert-2
        x-internal="assert"}: `el`{.variable}\'s
        [result](#concept-script-result){#script-processing-model:concept-script-result-4}
        is *not* \"`uninitialized`\".

    2.  ::: {#script-processing-style-delayed}
        If all of the following are true:

        - `el`{.variable}\'s
          [type](#concept-script-type){#script-processing-model:concept-script-type-15}
          is \"`classic`\";
        - `el`{.variable} is
          [parser-inserted](#parser-inserted){#script-processing-model:parser-inserted-6};
        - `el`{.variable}\'s [parser
          document](#parser-document){#script-processing-model:parser-document-7}
          [has a style sheet that is blocking
          scripts](semantics.html#has-a-style-sheet-that-is-blocking-scripts){#script-processing-model:has-a-style-sheet-that-is-blocking-scripts};
          and
        - either the parser that created `el`{.variable} is an [XML
          parser](xhtml.html#xml-parser){#script-processing-model:xml-parser-3},
          or it\'s an [HTML
          parser](parsing.html#html-parser){#script-processing-model:html-parser-3}
          whose [script nesting
          level](parsing.html#script-nesting-level){#script-processing-model:script-nesting-level}
          is not greater than one,

        then:

        1.  Set `el`{.variable}\'s [parser
            document](#parser-document){#script-processing-model:parser-document-8}\'s
            [pending parsing-blocking
            script](#pending-parsing-blocking-script){#script-processing-model:pending-parsing-blocking-script-2}
            to `el`{.variable}.

        2.  Set `el`{.variable}\'s [ready to be
            parser-executed](#ready-to-be-parser-executed){#script-processing-model:ready-to-be-parser-executed-3}
            to true. (The parser will handle executing the script.)
        :::

    3.  ::: {#script-processing-inline}
        Otherwise,
        [immediately](infrastructure.html#immediately){#script-processing-model:immediately}
        [execute the script
        element](#execute-the-script-element){#script-processing-model:execute-the-script-element-5}
        `el`{.variable}, even if other scripts are already executing.
        :::

Each
[`Document`{#script-processing-model:document-4}](dom.html#document) has
a [pending parsing-blocking script]{#pending-parsing-blocking-script
.dfn}, which is a
[`script`{#script-processing-model:the-script-element-39}](#the-script-element)
element or null, initially null.

Each
[`Document`{#script-processing-model:document-5}](dom.html#document) has
a [set of scripts that will execute as soon as
possible]{#set-of-scripts-that-will-execute-as-soon-as-possible .dfn},
which is a
[set](https://infra.spec.whatwg.org/#ordered-set){#script-processing-model:set
x-internal="set"} of
[`script`{#script-processing-model:the-script-element-40}](#the-script-element)
elements, initially empty.

Each
[`Document`{#script-processing-model:document-6}](dom.html#document) has
a [list of scripts that will execute in order as soon as
possible]{#list-of-scripts-that-will-execute-in-order-as-soon-as-possible
.dfn}, which is a
[list](https://infra.spec.whatwg.org/#list){#script-processing-model:list
x-internal="list"} of
[`script`{#script-processing-model:the-script-element-41}](#the-script-element)
elements, initially empty.

Each
[`Document`{#script-processing-model:document-7}](dom.html#document) has
a [list of scripts that will execute when the document has finished
parsing]{#list-of-scripts-that-will-execute-when-the-document-has-finished-parsing
.dfn}, which is a
[list](https://infra.spec.whatwg.org/#list){#script-processing-model:list-2
x-internal="list"} of
[`script`{#script-processing-model:the-script-element-42}](#the-script-element)
elements, initially empty.

If a
[`script`{#script-processing-model:the-script-element-43}](#the-script-element)
element that blocks a parser gets moved to another
[`Document`{#script-processing-model:document-8}](dom.html#document)
before it would normally have stopped blocking that parser, it
nonetheless continues blocking that parser until the condition that
causes it to be blocking the parser no longer applies (e.g., if the
script is a [pending parsing-blocking
script](#pending-parsing-blocking-script){#script-processing-model:pending-parsing-blocking-script-3}
because the original
[`Document`{#script-processing-model:document-9}](dom.html#document)
[has a style sheet that is blocking
scripts](semantics.html#has-a-style-sheet-that-is-blocking-scripts){#script-processing-model:has-a-style-sheet-that-is-blocking-scripts-2}
when it was parsed, but then the script is moved to another
[`Document`{#script-processing-model:document-10}](dom.html#document)
before the blocking style sheet(s) loaded, the script still blocks the
parser until the style sheets are all loaded, at which time the script
executes and the parser is unblocked).

To [execute the script element]{#execute-the-script-element .dfn} given
a
[`script`{#script-processing-model:the-script-element-44}](#the-script-element)
element `el`{.variable}:

1.  Let `document`{.variable} be `el`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#script-processing-model:node-document-6
    x-internal="node-document"}.

2.  If `el`{.variable}\'s [preparation-time
    document](#preparation-time-document){#script-processing-model:preparation-time-document-6}
    is not equal to `document`{.variable}, then return.

3.  [Unblock
    rendering](dom.html#unblock-rendering){#script-processing-model:unblock-rendering}
    on `el`{.variable}.

4.  If `el`{.variable}\'s
    [result](#concept-script-result){#script-processing-model:concept-script-result-5}
    is null, then [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#script-processing-model:concept-event-fire-4
    x-internal="concept-event-fire"} named
    [`error`{#script-processing-model:event-error-4}](indices.html#event-error)
    at `el`{.variable}, and return.

5.  If `el`{.variable}\'s [from an external
    file](#concept-script-external){#script-processing-model:concept-script-external-2}
    is true, or `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-16}
    is \"`module`\", then increment `document`{.variable}\'s
    [ignore-destructive-writes
    counter](dynamic-markup-insertion.html#ignore-destructive-writes-counter){#script-processing-model:ignore-destructive-writes-counter}.

6.  Switch on `el`{.variable}\'s
    [type](#concept-script-type){#script-processing-model:concept-script-type-17}:

    \"`classic`\"

    :   1.  Let `oldCurrentScript`{.variable} be the value to which
            `document`{.variable}\'s
            [`currentScript`{#script-processing-model:dom-document-currentscript}](dom.html#dom-document-currentscript)
            object was most recently set.

        2.  If `el`{.variable}\'s
            [root](https://dom.spec.whatwg.org/#concept-tree-root){#script-processing-model:root
            x-internal="root"} is *not* a [shadow
            root](https://dom.spec.whatwg.org/#concept-shadow-root){#script-processing-model:shadow-root
            x-internal="shadow-root"}, then set `document`{.variable}\'s
            [`currentScript`{#script-processing-model:dom-document-currentscript-2}](dom.html#dom-document-currentscript)
            attribute to `el`{.variable}. Otherwise, set it to null.

            This does not use the [in a document
            tree](https://dom.spec.whatwg.org/#in-a-document-tree){#script-processing-model:in-a-document-tree
            x-internal="in-a-document-tree"} check, as `el`{.variable}
            could have been removed from the document prior to
            execution, and in that scenario
            [`currentScript`{#script-processing-model:dom-document-currentscript-3}](dom.html#dom-document-currentscript)
            still needs to point to it.

        3.  [Run the classic
            script](webappapis.html#run-a-classic-script){#script-processing-model:run-a-classic-script}
            given by `el`{.variable}\'s
            [result](#concept-script-result){#script-processing-model:concept-script-result-6}.

        4.  Set `document`{.variable}\'s
            [`currentScript`{#script-processing-model:dom-document-currentscript-4}](dom.html#dom-document-currentscript)
            attribute to `oldCurrentScript`{.variable}.

    \"`module`\"

    :   1.  [Assert](https://infra.spec.whatwg.org/#assert){#script-processing-model:assert-3
            x-internal="assert"}: `document`{.variable}\'s
            [`currentScript`{#script-processing-model:dom-document-currentscript-5}](dom.html#dom-document-currentscript)
            attribute is null.

        2.  [Run the module
            script](webappapis.html#run-a-module-script){#script-processing-model:run-a-module-script}
            given by `el`{.variable}\'s
            [result](#concept-script-result){#script-processing-model:concept-script-result-7}.

    \"`importmap`\"

    :   1.  [Register an import
            map](webappapis.html#register-an-import-map){#script-processing-model:register-an-import-map}
            given `el`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#script-processing-model:concept-relevant-global}
            and `el`{.variable}\'s
            [result](#concept-script-result){#script-processing-model:concept-script-result-8}.

7.  Decrement the [ignore-destructive-writes
    counter](dynamic-markup-insertion.html#ignore-destructive-writes-counter){#script-processing-model:ignore-destructive-writes-counter-2}
    of `document`{.variable}, if it was incremented in the earlier step.

8.  If `el`{.variable}\'s [from an external
    file](#concept-script-external){#script-processing-model:concept-script-external-3}
    is true, then [fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#script-processing-model:concept-event-fire-5
    x-internal="concept-event-fire"} named
    [`load`{#script-processing-model:event-load}](indices.html#event-load)
    at `el`{.variable}.

##### [4.12.1.2]{.secno} Scripting languages[](#scriptingLanguages){.self-link} {#scriptingLanguages}

User agents are not required to support JavaScript. This standard needs
to be updated if a language other than JavaScript comes along and gets
similar wide adoption by web browsers. Until such a time, implementing
other languages is in conflict with this standard, given the processing
model defined for the
[`script`{#scriptingLanguages:the-script-element}](#the-script-element)
element.

Servers should use
[`text/javascript`{#scriptingLanguages:text/javascript}](indices.html#text/javascript)
for JavaScript resources, in accordance with Updates to ECMAScript Media
Types. Servers should not use other [JavaScript MIME
types](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#scriptingLanguages:javascript-mime-type
x-internal="javascript-mime-type"} for JavaScript resources, and must
not use non-[JavaScript MIME
types](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#scriptingLanguages:javascript-mime-type-2
x-internal="javascript-mime-type"}.
[\[RFC9239\]](references.html#refsRFC9239)

For external JavaScript resources, MIME type parameters in
\`[`Content-Type`{#scriptingLanguages:content-type}](urls-and-fetching.html#content-type)\`
headers are generally ignored. (In some cases the \``charset`\`
parameter has an effect.) However, for the
[`script`{#scriptingLanguages:the-script-element-2}](#the-script-element)
element\'s
[`type`{#scriptingLanguages:attr-script-type}](#attr-script-type)
attribute they are significant; it uses the [JavaScript MIME type
essence
match](https://mimesniff.spec.whatwg.org/#javascript-mime-type-essence-match){#scriptingLanguages:javascript-mime-type-essence-match
x-internal="javascript-mime-type-essence-match"} concept.

For example, scripts with their
[`type`{#scriptingLanguages:attr-script-type-2}](#attr-script-type)
attribute set to \"`text/javascript; charset=utf-8`\" will not be
evaluated, even though that is a valid [JavaScript MIME
type](https://mimesniff.spec.whatwg.org/#javascript-mime-type){#scriptingLanguages:javascript-mime-type-3
x-internal="javascript-mime-type"} when parsed.

Furthermore, again for external JavaScript resources, special
considerations apply around
\`[`Content-Type`{#scriptingLanguages:content-type-2}](urls-and-fetching.html#content-type)\`
header processing as detailed in the [prepare the script
element](#prepare-the-script-element){#scriptingLanguages:prepare-the-script-element}
algorithm and Fetch. [\[FETCH\]](references.html#refsFETCH)

##### [4.12.1.3]{.secno} [Restrictions for contents of `script` elements]{.dfn}[](#restrictions-for-contents-of-script-elements){.self-link}

The easiest and safest way to avoid the rather strange restrictions
described in this section is to always escape an ASCII case-insensitive
match for \"`<!--`\" as \"`\x3C!--`\", \"`<script`\" as
\"`\x3Cscript`\", and \"`</script`\" as \"`\x3C/script`\" when these
sequences appear in literals in scripts (e.g. in strings, regular
expressions, or comments), and to avoid writing code that uses such
constructs in expressions. Doing so avoids the pitfalls that the
restrictions in this section are prone to triggering: namely, that, for
historical reasons, parsing of
[`script`{#restrictions-for-contents-of-script-elements:the-script-element}](#the-script-element)
blocks in HTML is a strange and exotic practice that acts unintuitively
in the face of these sequences.

The
[`script`{#restrictions-for-contents-of-script-elements:the-script-element-2}](#the-script-element)
element\'s [descendant text
content](https://dom.spec.whatwg.org/#concept-descendant-text-content){#restrictions-for-contents-of-script-elements:descendant-text-content
x-internal="descendant-text-content"} must match the `script` production
in the following ABNF, the character set for which is Unicode.
[\[ABNF\]](references.html#refsABNF)

``` abnf
script        = outer *( comment-open inner comment-close outer )

outer         = < any string that doesn't contain a substring that matches not-in-outer >
not-in-outer  = comment-open
inner         = < any string that doesn't contain a substring that matches not-in-inner >
not-in-inner  = comment-close / script-open

comment-open  = "<!--"
comment-close = "-->"
script-open   = "<" s c r i p t tag-end

s             =  %x0053 ; U+0053 LATIN CAPITAL LETTER S
s             =/ %x0073 ; U+0073 LATIN SMALL LETTER S
c             =  %x0043 ; U+0043 LATIN CAPITAL LETTER C
c             =/ %x0063 ; U+0063 LATIN SMALL LETTER C
r             =  %x0052 ; U+0052 LATIN CAPITAL LETTER R
r             =/ %x0072 ; U+0072 LATIN SMALL LETTER R
i             =  %x0049 ; U+0049 LATIN CAPITAL LETTER I
i             =/ %x0069 ; U+0069 LATIN SMALL LETTER I
p             =  %x0050 ; U+0050 LATIN CAPITAL LETTER P
p             =/ %x0070 ; U+0070 LATIN SMALL LETTER P
t             =  %x0054 ; U+0054 LATIN CAPITAL LETTER T
t             =/ %x0074 ; U+0074 LATIN SMALL LETTER T

tag-end       =  %x0009 ; U+0009 CHARACTER TABULATION (tab)
tag-end       =/ %x000A ; U+000A LINE FEED (LF)
tag-end       =/ %x000C ; U+000C FORM FEED (FF)
tag-end       =/ %x0020 ; U+0020 SPACE
tag-end       =/ %x002F ; U+002F SOLIDUS (/)
tag-end       =/ %x003E ; U+003E GREATER-THAN SIGN (>)
```

When a
[`script`{#restrictions-for-contents-of-script-elements:the-script-element-3}](#the-script-element)
element contains [script
documentation](#inline-documentation-for-external-scripts){#restrictions-for-contents-of-script-elements:inline-documentation-for-external-scripts},
there are further restrictions on the contents of the element, as
described in the section below.

::: example
The following script illustrates this issue. Suppose you have a script
that contains a string, as in:

``` js
const example = 'Consider this string: <!-- <script>';
console.log(example);
```

If one were to put this string directly in a
[`script`{#restrictions-for-contents-of-script-elements:the-script-element-4}](#the-script-element)
block, it would violate the restrictions above:

``` html
<script>
  const example = 'Consider this string: <!-- <script>';
  console.log(example);
</script>
```

The bigger problem, though, and the reason why it would violate those
restrictions, is that actually the script would get parsed weirdly: *the
script block above is not terminated*. That is, what looks like a
\"`</script>`\" end tag in this snippet is actually still part of the
[`script`{#restrictions-for-contents-of-script-elements:the-script-element-5}](#the-script-element)
block. The script doesn\'t execute (since it\'s not terminated); if it
somehow were to execute, as it might if the markup looked as follows, it
would fail because the script (highlighted here) is not valid
JavaScript:

``` html
<script>
  const example = 'Consider this string: <!-- <script>';
  console.log(example);
</script>
<!-- despite appearances, this is actually part of the script still! -->
<script>
 ... // this is the same script block still...
</script>
```

What is going on here is that for legacy reasons, \"`<!--`\" and
\"`<script`\" strings in
[`script`{#restrictions-for-contents-of-script-elements:the-script-element-6}](#the-script-element)
elements in HTML need to be balanced in order for the parser to consider
closing the block.

By escaping the problematic strings as mentioned at the top of this
section, the problem is avoided entirely:

``` html
<script>
  // Note: `\x3C` is an escape sequence for `<`.
  const example = 'Consider this string: \x3C!-- \x3Cscript>';
  console.log(example);
</script>
<!-- this is just a comment between script blocks -->
<script>
 ... // this is a new script block
</script>
```

It is possible for these sequences to naturally occur in script
expressions, as in the following examples:

``` js
if (x<!--y) { ... }
if ( player<script ) { ... }
```

In such cases the characters cannot be escaped, but the expressions can
be rewritten so that the sequences don\'t occur, as in:

``` js
if (x < !--y) { ... }
if (!--y > x) { ... }
if (!(--y) > x) { ... }
if (player < script) { ... }
if (script > player) { ... }
```

Doing this also avoids a different pitfall as well: for related
historical reasons, the string \"\<!\--\" in [classic
scripts](webappapis.html#classic-script){#restrictions-for-contents-of-script-elements:classic-script}
is actually treated as a line comment start, just like \"//\".
:::

##### [4.12.1.4]{.secno} [Inline documentation for external scripts]{.dfn}[](#inline-documentation-for-external-scripts){.self-link}

If a
[`script`{#inline-documentation-for-external-scripts:the-script-element}](#the-script-element)
element\'s
[`src`{#inline-documentation-for-external-scripts:attr-script-src}](#attr-script-src)
attribute is specified, then the contents of the
[`script`{#inline-documentation-for-external-scripts:the-script-element-2}](#the-script-element)
element, if any, must be such that the value of the
[`text`{#inline-documentation-for-external-scripts:dom-script-text}](#dom-script-text)
IDL attribute, which is derived from the element\'s contents, matches
the `documentation` production in the following ABNF, the character set
for which is Unicode. [\[ABNF\]](references.html#refsABNF)

``` abnf
documentation = *( *( space / tab / comment ) [ line-comment ] newline )
comment       = slash star *( not-star / star not-slash ) 1*star slash
line-comment  = slash slash *not-newline

; characters
tab           = %x0009 ; U+0009 CHARACTER TABULATION (tab)
newline       = %x000A ; U+000A LINE FEED (LF)
space         = %x0020 ; U+0020 SPACE
star          = %x002A ; U+002A ASTERISK (*)
slash         = %x002F ; U+002F SOLIDUS (/)
not-newline   = %x0000-0009 / %x000B-10FFFF
                ; a scalar value other than U+000A LINE FEED (LF)
not-star      = %x0000-0029 / %x002B-10FFFF
                ; a scalar value other than U+002A ASTERISK (*)
not-slash     = %x0000-002E / %x0030-10FFFF
                ; a scalar value other than U+002F SOLIDUS (/)
```

This corresponds to putting the contents of the element in JavaScript
comments.

This requirement is in addition to the earlier restrictions on the
syntax of contents of
[`script`{#inline-documentation-for-external-scripts:the-script-element-3}](#the-script-element)
elements.

::: example
This allows authors to include documentation, such as license
information or API information, inside their documents while still
referring to external script files. The syntax is constrained so that
authors don\'t accidentally include what looks like valid script while
also providing a
[`src`{#inline-documentation-for-external-scripts:attr-script-src-2}](#attr-script-src)
attribute.

``` html
<script src="cool-effects.js">
 // create new instances using:
 //    var e = new Effect();
 // start the effect using .play, stop using .stop:
 //    e.play();
 //    e.stop();
</script>
```
:::

##### [4.12.1.5]{.secno} Interaction of [`script`{#scriptTagXSLT:the-script-element}](#the-script-element) elements and XSLT[](#scriptTagXSLT){.self-link} {#scriptTagXSLT}

*This section is non-normative.*

This specification does not define how XSLT interacts with the
[`script`{#scriptTagXSLT:the-script-element-2}](#the-script-element)
element. However, in the absence of another specification actually
defining this, here are some guidelines for implementers, based on
existing implementations:

- When an XSLT transformation program is triggered by an
  `<?xml-stylesheet?>` processing instruction and the browser implements
  a direct-to-DOM transformation,
  [`script`{#scriptTagXSLT:the-script-element-3}](#the-script-element)
  elements created by the XSLT processor need to have its [parser
  document](#parser-document){#scriptTagXSLT:parser-document} set
  correctly, and run in document order (modulo scripts marked
  [`defer`{#scriptTagXSLT:attr-script-defer}](#attr-script-defer) or
  [`async`{#scriptTagXSLT:attr-script-async}](#attr-script-async)),
  [immediately](infrastructure.html#immediately){#scriptTagXSLT:immediately},
  as the transformation is occurring.

- The
  [`XSLTProcessor`{#scriptTagXSLT:xsltprocessor}](infrastructure.html#xsltprocessor)
  [`transformToDocument()`{#scriptTagXSLT:dom-xsltprocessor-transformtodocument}](infrastructure.html#dom-xsltprocessor-transformtodocument)
  method adds elements to a
  [`Document`{#scriptTagXSLT:document}](dom.html#document) object with a
  null [browsing
  context](document-sequences.html#concept-document-bc){#scriptTagXSLT:concept-document-bc},
  and, accordingly, any
  [`script`{#scriptTagXSLT:the-script-element-4}](#the-script-element)
  elements they create need to have their [already
  started](#already-started){#scriptTagXSLT:already-started} set to true
  in the [prepare the script
  element](#prepare-the-script-element){#scriptTagXSLT:prepare-the-script-element}
  algorithm and never get executed ([scripting is
  disabled](webappapis.html#concept-environment-noscript){#scriptTagXSLT:concept-environment-noscript}).
  Such
  [`script`{#scriptTagXSLT:the-script-element-5}](#the-script-element)
  elements still need to have their [parser
  document](#parser-document){#scriptTagXSLT:parser-document-2} set,
  though, such that their
  [`async`{#scriptTagXSLT:dom-script-async}](#dom-script-async) IDL
  attribute will return false in the absence of an
  [`async`{#scriptTagXSLT:attr-script-async-2}](#attr-script-async)
  content attribute.

- The
  [`XSLTProcessor`{#scriptTagXSLT:xsltprocessor-2}](infrastructure.html#xsltprocessor)
  [`transformToFragment()`{#scriptTagXSLT:dom-xsltprocessor-transformtofragment}](infrastructure.html#dom-xsltprocessor-transformtofragment)
  method needs to create a fragment that is equivalent to one built
  manually by creating the elements using
  [`document.createElementNS()`{#scriptTagXSLT:dom-document-createelementns}](https://dom.spec.whatwg.org/#dom-document-createelementns){x-internal="dom-document-createelementns"}.
  For instance, it needs to create
  [`script`{#scriptTagXSLT:the-script-element-6}](#the-script-element)
  elements with null [parser
  document](#parser-document){#scriptTagXSLT:parser-document-3} and with
  their [already
  started](#already-started){#scriptTagXSLT:already-started-2} set to
  false, so that they will execute when the fragment is inserted into a
  document.

The main distinction between the first two cases and the last case is
that the first two operate on
[`Document`{#scriptTagXSLT:document-2}](dom.html#document)s and the last
operates on a fragment.

#### [4.12.2]{.secno} The [`noscript`]{.dfn dfn-type="element"} element[](#the-noscript-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/noscript](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript "The <noscript> HTML element defines a section of HTML to be inserted if a script type on the page is unsupported or if scripting is currently turned off in the browser.")

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

[Categories](dom.html#concept-element-categories){#the-noscript-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-noscript-element:metadata-content-2}.
:   [Flow
    content](dom.html#flow-content-2){#the-noscript-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-noscript-element:phrasing-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-noscript-element:concept-element-contexts}:
:   In a
    [`head`{#the-noscript-element:the-head-element}](semantics.html#the-head-element)
    element of an [HTML
    document](https://dom.spec.whatwg.org/#html-document){#the-noscript-element:html-documents
    x-internal="html-documents"}, if there are no ancestor
    [`noscript`{#the-noscript-element:the-noscript-element}](#the-noscript-element)
    elements.
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-noscript-element:phrasing-content-2-2}
    is expected in [HTML
    documents](https://dom.spec.whatwg.org/#html-document){#the-noscript-element:html-documents-2
    x-internal="html-documents"}, if there are no ancestor
    [`noscript`{#the-noscript-element:the-noscript-element-2}](#the-noscript-element)
    elements.

[Content model](dom.html#concept-element-content-model){#the-noscript-element:concept-element-content-model}:
:   When [scripting is
    disabled](webappapis.html#concept-n-noscript){#the-noscript-element:concept-n-noscript},
    in a
    [`head`{#the-noscript-element:the-head-element-2}](semantics.html#the-head-element)
    element: in any order, zero or more
    [`link`{#the-noscript-element:the-link-element}](semantics.html#the-link-element)
    elements, zero or more
    [`style`{#the-noscript-element:the-style-element}](semantics.html#the-style-element)
    elements, and zero or more
    [`meta`{#the-noscript-element:the-meta-element}](semantics.html#the-meta-element)
    elements.
:   When [scripting is
    disabled](webappapis.html#concept-n-noscript){#the-noscript-element:concept-n-noscript-2},
    not in a
    [`head`{#the-noscript-element:the-head-element-3}](semantics.html#the-head-element)
    element:
    [transparent](dom.html#transparent){#the-noscript-element:transparent},
    but there must be no
    [`noscript`{#the-noscript-element:the-noscript-element-3}](#the-noscript-element)
    element descendants.
:   Otherwise: text that conforms to the requirements given in the
    prose.

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-noscript-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-noscript-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-noscript-element:global-attributes}

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-noscript-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-noscript).
:   [For implementers](https://w3c.github.io/html-aam/#el-noscript).

[DOM interface](dom.html#concept-element-dom){#the-noscript-element:concept-element-dom}:
:   Uses
    [`HTMLElement`{#the-noscript-element:htmlelement}](dom.html#htmlelement).

The
[`noscript`{#the-noscript-element:the-noscript-element-4}](#the-noscript-element)
element
[represents](dom.html#represents){#the-noscript-element:represents}
nothing if [scripting is
enabled](webappapis.html#concept-n-script){#the-noscript-element:concept-n-script},
and
[represents](dom.html#represents){#the-noscript-element:represents-2}
its children if [scripting is
disabled](webappapis.html#concept-n-noscript){#the-noscript-element:concept-n-noscript-3}.
It is used to present different markup to user agents that support
scripting and those that don\'t support scripting, by affecting how the
document is parsed.

When used in [HTML
documents](https://dom.spec.whatwg.org/#html-document){#the-noscript-element:html-documents-3
x-internal="html-documents"}, the allowed content model is as follows:

In a [`head`{#the-noscript-element:the-head-element-4}](semantics.html#the-head-element) element, if [scripting is disabled](webappapis.html#concept-n-noscript){#the-noscript-element:concept-n-noscript-4} for the [`noscript`{#the-noscript-element:the-noscript-element-5}](#the-noscript-element) element
:   The
    [`noscript`{#the-noscript-element:the-noscript-element-6}](#the-noscript-element)
    element must contain only
    [`link`{#the-noscript-element:the-link-element-2}](semantics.html#the-link-element),
    [`style`{#the-noscript-element:the-style-element-2}](semantics.html#the-style-element),
    and
    [`meta`{#the-noscript-element:the-meta-element-2}](semantics.html#the-meta-element)
    elements.

In a [`head`{#the-noscript-element:the-head-element-5}](semantics.html#the-head-element) element, if [scripting is enabled](webappapis.html#concept-n-script){#the-noscript-element:concept-n-script-2} for the [`noscript`{#the-noscript-element:the-noscript-element-7}](#the-noscript-element) element
:   The
    [`noscript`{#the-noscript-element:the-noscript-element-8}](#the-noscript-element)
    element must contain only text, except that invoking the [HTML
    fragment parsing
    algorithm](parsing.html#html-fragment-parsing-algorithm){#the-noscript-element:html-fragment-parsing-algorithm}
    with the
    [`noscript`{#the-noscript-element:the-noscript-element-9}](#the-noscript-element)
    element as the
    [`context`{#the-noscript-element:concept-frag-parse-context
    .variable}](parsing.html#concept-frag-parse-context) element and the
    text contents as the `input`{.variable} must result in a list of
    nodes that consists only of
    [`link`{#the-noscript-element:the-link-element-3}](semantics.html#the-link-element),
    [`style`{#the-noscript-element:the-style-element-3}](semantics.html#the-style-element),
    and
    [`meta`{#the-noscript-element:the-meta-element-3}](semantics.html#the-meta-element)
    elements that would be conforming if they were children of the
    [`noscript`{#the-noscript-element:the-noscript-element-10}](#the-noscript-element)
    element, and no [parse
    errors](parsing.html#parse-errors){#the-noscript-element:parse-errors}.

Outside of [`head`{#the-noscript-element:the-head-element-6}](semantics.html#the-head-element) elements, if [scripting is disabled](webappapis.html#concept-n-noscript){#the-noscript-element:concept-n-noscript-5} for the [`noscript`{#the-noscript-element:the-noscript-element-11}](#the-noscript-element) element
:   The
    [`noscript`{#the-noscript-element:the-noscript-element-12}](#the-noscript-element)
    element\'s content model is
    [transparent](dom.html#transparent){#the-noscript-element:transparent-2},
    with the additional restriction that a
    [`noscript`{#the-noscript-element:the-noscript-element-13}](#the-noscript-element)
    element must not have a
    [`noscript`{#the-noscript-element:the-noscript-element-14}](#the-noscript-element)
    element as an ancestor (that is,
    [`noscript`{#the-noscript-element:the-noscript-element-15}](#the-noscript-element)
    can\'t be nested).

Outside of [`head`{#the-noscript-element:the-head-element-7}](semantics.html#the-head-element) elements, if [scripting is enabled](webappapis.html#concept-n-script){#the-noscript-element:concept-n-script-3} for the [`noscript`{#the-noscript-element:the-noscript-element-16}](#the-noscript-element) element

:   The
    [`noscript`{#the-noscript-element:the-noscript-element-17}](#the-noscript-element)
    element must contain only text, except that the text must be such
    that running the following algorithm results in a conforming
    document with no
    [`noscript`{#the-noscript-element:the-noscript-element-18}](#the-noscript-element)
    elements and no
    [`script`{#the-noscript-element:the-script-element}](#the-script-element)
    elements, and such that no step in the algorithm throws an exception
    or causes an [HTML
    parser](parsing.html#html-parser){#the-noscript-element:html-parser}
    to flag a [parse
    error](parsing.html#parse-errors){#the-noscript-element:parse-errors-2}:

    1.  Remove every
        [`script`{#the-noscript-element:the-script-element-2}](#the-script-element)
        element from the document.
    2.  Make a list of every
        [`noscript`{#the-noscript-element:the-noscript-element-19}](#the-noscript-element)
        element in the document. For every
        [`noscript`{#the-noscript-element:the-noscript-element-20}](#the-noscript-element)
        element in that list, perform the following steps:
        1.  Let `s`{.variable} be the [child text
            content](https://dom.spec.whatwg.org/#concept-child-text-content){#the-noscript-element:child-text-content
            x-internal="child-text-content"} of the
            [`noscript`{#the-noscript-element:the-noscript-element-21}](#the-noscript-element)
            element.
        2.  Set the
            [`outerHTML`{#the-noscript-element:dom-element-outerhtml}](dynamic-markup-insertion.html#dom-element-outerhtml)
            attribute of the
            [`noscript`{#the-noscript-element:the-noscript-element-22}](#the-noscript-element)
            element to the value of `s`{.variable}. (This, as a
            side-effect, causes the
            [`noscript`{#the-noscript-element:the-noscript-element-23}](#the-noscript-element)
            element to be removed from the document.)

All these contortions are required because, for historical reasons, the
[`noscript`{#the-noscript-element:the-noscript-element-24}](#the-noscript-element)
element is handled differently by the [HTML
parser](parsing.html#html-parser){#the-noscript-element:html-parser-2}
based on whether [scripting was enabled or
not](parsing.html#scripting-flag){#the-noscript-element:scripting-flag}
when the parser was invoked.

The
[`noscript`{#the-noscript-element:the-noscript-element-25}](#the-noscript-element)
element must not be used in [XML
documents](https://dom.spec.whatwg.org/#xml-document){#the-noscript-element:xml-documents
x-internal="xml-documents"}.

The
[`noscript`{#the-noscript-element:the-noscript-element-26}](#the-noscript-element)
element is only effective in [the HTML
syntax](syntax.html#syntax){#the-noscript-element:syntax}, it has no
effect in [the XML
syntax](xhtml.html#the-xhtml-syntax){#the-noscript-element:the-xhtml-syntax}.
This is because the way it works is by essentially \"turning off\" the
parser when scripts are enabled, so that the contents of the element are
treated as pure text and not as real elements. XML does not define a
mechanism by which to do this.

The
[`noscript`{#the-noscript-element:the-noscript-element-27}](#the-noscript-element)
element has no other requirements. In particular, children of the
[`noscript`{#the-noscript-element:the-noscript-element-28}](#the-noscript-element)
element are not exempt from [form
submission](form-control-infrastructure.html#form-submission-2){#the-noscript-element:form-submission-2},
scripting, and so forth, even when [scripting is
enabled](webappapis.html#concept-n-script){#the-noscript-element:concept-n-script-4}
for the element.

::: example
In the following example, a
[`noscript`{#the-noscript-element:the-noscript-element-29}](#the-noscript-element)
element is used to provide fallback for a script.

``` html
<form action="calcSquare.php">
 <p>
  <label for=x>Number</label>:
  <input id="x" name="x" type="number">
 </p>
 <script>
  var x = document.getElementById('x');
  var output = document.createElement('p');
  output.textContent = 'Type a number; it will be squared right then!';
  x.form.appendChild(output);
  x.form.onsubmit = function () { return false; }
  x.oninput = function () {
    var v = x.valueAsNumber;
    output.textContent = v + ' squared is ' + v * v;
  };
 </script>
 <noscript>
  <input type=submit value="Calculate Square">
 </noscript>
</form>
```

When script is disabled, a button appears to do the calculation on the
server side. When script is enabled, the value is computed on-the-fly
instead.

The
[`noscript`{#the-noscript-element:the-noscript-element-30}](#the-noscript-element)
element is a blunt instrument. Sometimes, scripts might be enabled, but
for some reason the page\'s script might fail. For this reason, it\'s
generally better to avoid using
[`noscript`{#the-noscript-element:the-noscript-element-31}](#the-noscript-element),
and to instead design the script to change the page from being a
scriptless page to a scripted page on the fly, as in the next example:

``` html
<form action="calcSquare.php">
 <p>
  <label for=x>Number</label>:
  <input id="x" name="x" type="number">
 </p>
 <input id="submit" type=submit value="Calculate Square">
 <script>
  var x = document.getElementById('x');
  var output = document.createElement('p');
  output.textContent = 'Type a number; it will be squared right then!';
  x.form.appendChild(output);
  x.form.onsubmit = function () { return false; }
  x.oninput = function () {
    var v = x.valueAsNumber;
    output.textContent = v + ' squared is ' + v * v;
  };
  var submit = document.getElementById('submit');
  submit.parentNode.removeChild(submit);
 </script>
</form>
```

The above technique is also useful in [XML
documents](https://dom.spec.whatwg.org/#xml-document){#the-noscript-element:xml-documents-2
x-internal="xml-documents"}, since
[`noscript`{#the-noscript-element:the-noscript-element-32}](#the-noscript-element)
is not allowed there.
:::

#### [4.12.3]{.secno} The [`template`]{.dfn dfn-type="element"} element[](#the-template-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/template](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template "The <template> HTML element is a mechanism for holding HTML that is not to be rendered immediately when a page is loaded but may be instantiated subsequently during runtime using JavaScript.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome26+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
AndroidYes]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTemplateElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTemplateElement "The HTMLTemplateElement interface enables access to the contents of an HTML <template> element.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome26+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[Categories](dom.html#concept-element-categories){#the-template-element:concept-element-categories}:
:   [Metadata
    content](dom.html#metadata-content-2){#the-template-element:metadata-content-2}.
:   [Flow
    content](dom.html#flow-content-2){#the-template-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-template-element:phrasing-content-2}.
:   [Script-supporting
    element](dom.html#script-supporting-elements-2){#the-template-element:script-supporting-elements-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-template-element:concept-element-contexts}:
:   Where [metadata
    content](dom.html#metadata-content-2){#the-template-element:metadata-content-2-2}
    is expected.
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-template-element:phrasing-content-2-2}
    is expected.
:   Where [script-supporting
    elements](dom.html#script-supporting-elements-2){#the-template-element:script-supporting-elements-2-2}
    are expected.
:   As a child of a
    [`colgroup`{#the-template-element:the-colgroup-element}](tables.html#the-colgroup-element)
    element that doesn\'t have a
    [`span`{#the-template-element:attr-colgroup-span}](tables.html#attr-colgroup-span)
    attribute.

[Content model](dom.html#concept-element-content-model){#the-template-element:concept-element-content-model}:
:   [Nothing](dom.html#concept-content-nothing){#the-template-element:concept-content-nothing}
    (for clarification, [see example](#template-example)).

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-template-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-template-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-template-element:global-attributes}
:   [`shadowrootmode`{#the-template-element:attr-template-shadowrootmode}](#attr-template-shadowrootmode)
    --- Enables streaming declarative shadow roots
:   [`shadowrootdelegatesfocus`{#the-template-element:attr-template-shadowrootdelegatesfocus}](#attr-template-shadowrootdelegatesfocus)
    --- Sets [delegates
    focus](https://dom.spec.whatwg.org/#shadowroot-delegates-focus){#the-template-element:delegates-focus
    x-internal="delegates-focus"} on a declarative shadow root
:   [`shadowrootclonable`{#the-template-element:attr-template-shadowrootclonable}](#attr-template-shadowrootclonable)
    --- Sets
    [clonable](https://dom.spec.whatwg.org/#shadowroot-clonable){#the-template-element:clonable
    x-internal="clonable"} on a declarative shadow root
:   [`shadowrootserializable`{#the-template-element:attr-template-shadowrootserializable}](#attr-template-shadowrootserializable)
    --- Sets
    [serializable](structured-data.html#serializable){#the-template-element:serializable}
    on a declarative shadow root
:   [`shadowrootcustomelementregistry`{#the-template-element:attr-template-shadowrootcustomelementregistry}](#attr-template-shadowrootcustomelementregistry)
    --- Enables declarative shadow roots to indicate they will use a
    custom element registry

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-template-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-template).
:   [For implementers](https://w3c.github.io/html-aam/#el-template).

[DOM interface](dom.html#concept-element-dom){#the-template-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLTemplateElement : HTMLElement {
      [HTMLConstructor] constructor();

      readonly attribute DocumentFragment content;
      [CEReactions] attribute DOMString shadowRootMode;
      [CEReactions] attribute boolean shadowRootDelegatesFocus;
      [CEReactions] attribute boolean shadowRootClonable;
      [CEReactions] attribute boolean shadowRootSerializable;
      [CEReactions] attribute DOMString shadowRootCustomElementRegistry;
    };
    ```

The
[`template`{#the-template-element:the-template-element}](#the-template-element)
element is used to declare fragments of HTML that can be cloned and
inserted in the document by script.

In a rendering, the
[`template`{#the-template-element:the-template-element-2}](#the-template-element)
element
[represents](dom.html#represents){#the-template-element:represents}
nothing.

The [`shadowrootmode`]{#attr-template-shadowrootmode .dfn
dfn-for="template" dfn-type="element-attr"} content attribute is an
[enumerated
attribute](common-microsyntaxes.html#enumerated-attribute){#the-template-element:enumerated-attribute}
with the following keywords and states:

Keyword

State

Brief description

[`open`]{#attr-shadowrootmode-open .dfn
dfn-for="template/shadowrootmode" dfn-type="attr-value"}

[Open]{#attr-shadowrootmode-open-state .dfn}

The template element represents an open declarative shadow root.

[`closed`]{#attr-shadowrootmode-closed .dfn
dfn-for="template/shadowrootmode" dfn-type="attr-value"}

[Closed]{#attr-shadowrootmode-closed-state .dfn}

The template element represents a closed declarative shadow root.

The
[`shadowrootmode`{#the-template-element:attr-template-shadowrootmode-2}](#attr-template-shadowrootmode)
attribute\'s *[invalid value
default](common-microsyntaxes.html#invalid-value-default)* and *[missing
value default](common-microsyntaxes.html#missing-value-default)* are
both the [None]{#attr-shadowrootmode-none-state .dfn} state.

The [`shadowrootdelegatesfocus`]{#attr-template-shadowrootdelegatesfocus
.dfn dfn-for="template" dfn-type="element-attr"} content attribute is a
[boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-template-element:boolean-attribute}.

The [`shadowrootclonable`]{#attr-template-shadowrootclonable .dfn
dfn-for="template" dfn-type="element-attr"} content attribute is a
[boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-template-element:boolean-attribute-2}.

The [`shadowrootserializable`]{#attr-template-shadowrootserializable
.dfn dfn-for="template" dfn-type="element-attr"} content attribute is a
[boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-template-element:boolean-attribute-3}.

The
[`shadowrootcustomelementregistry`]{#attr-template-shadowrootcustomelementregistry
.dfn dfn-for="template" dfn-type="element-attr"} content attribute is a
[boolean
attribute](common-microsyntaxes.html#boolean-attribute){#the-template-element:boolean-attribute-4}.

The [template
contents](#template-contents){#the-template-element:template-contents}
of a
[`template`{#the-template-element:the-template-element-3}](#the-template-element)
element [are not children of the element
itself](syntax.html#template-syntax).

It is also possible, as a result of DOM manipulation, for a
[`template`{#the-template-element:the-template-element-4}](#the-template-element)
element to contain
[`Text`{#the-template-element:text}](https://dom.spec.whatwg.org/#interface-text){x-internal="text"}
nodes and element nodes; however, having any is a violation of the
[`template`{#the-template-element:the-template-element-5}](#the-template-element)
element\'s content model, since its content model is defined as
[nothing](dom.html#concept-content-nothing){#the-template-element:concept-content-nothing-2}.

::: {#template-example .example}
For example, consider the following document:

``` html
<!doctype html>
<html lang="en">
 <head>
  <title>Homework</title>
 <body>
  <template id="template"><p>Smile!</p></template>
  <script>
   let num = 3;
   const fragment = document.getElementById('template').content.cloneNode(true);
   while (num-- > 1) {
     fragment.firstChild.before(fragment.firstChild.cloneNode(true));
     fragment.firstChild.textContent += fragment.lastChild.textContent;
   }
   document.body.appendChild(fragment);
  </script>
</html>
```

The
[`p`{#the-template-element:the-p-element}](grouping-content.html#the-p-element)
element in the
[`template`{#the-template-element:the-template-element-6}](#the-template-element)
is *not* a child of the
[`template`{#the-template-element:the-template-element-7}](#the-template-element)
in the DOM; it is a child of the
[`DocumentFragment`{#the-template-element:documentfragment-2}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
returned by the
[`template`{#the-template-element:the-template-element-8}](#the-template-element)
element\'s
[`content`{#the-template-element:dom-template-content-2}](#dom-template-content)
IDL attribute.

If the script were to call
[`appendChild()`{#the-template-element:dom-node-appendchild}](https://dom.spec.whatwg.org/#dom-node-appendchild){x-internal="dom-node-appendchild"}
on the
[`template`{#the-template-element:the-template-element-9}](#the-template-element)
element, that would add a child to the
[`template`{#the-template-element:the-template-element-10}](#the-template-element)
element (as for any other element); however, doing so is a violation of
the
[`template`{#the-template-element:the-template-element-11}](#the-template-element)
element\'s content model.
:::

`template`{.variable}`.`[`content`](#dom-template-content){#dom-template-content-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLTemplateElement/content](https://developer.mozilla.org/en-US/docs/Web/API/HTMLTemplateElement/content "The HTMLTemplateElement.content property returns a <template> element's template contents (a DocumentFragment).")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari8+]{.safari .yes}[Chrome26+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)13+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the [template
contents](#template-contents){#the-template-element:template-contents-2}
(a
[`DocumentFragment`{#the-template-element:documentfragment-3}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}).

Each
[`template`{#the-template-element:the-template-element-12}](#the-template-element)
element has an associated
[`DocumentFragment`{#the-template-element:documentfragment-4}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
object that is its [template contents]{#template-contents .dfn}. The
[template
contents](#template-contents){#the-template-element:template-contents-3}
have [no conformance requirements](dom.html#no-browsing-context). When a
[`template`{#the-template-element:the-template-element-13}](#the-template-element)
element is created, the user agent must run the following steps to
establish the [template
contents](#template-contents){#the-template-element:template-contents-4}:

1.  Let `doc`{.variable} be the
    [`template`{#the-template-element:the-template-element-14}](#the-template-element)
    element\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-template-element:node-document
    x-internal="node-document"}\'s [appropriate template contents owner
    document](#appropriate-template-contents-owner-document){#the-template-element:appropriate-template-contents-owner-document}.

2.  Create a
    [`DocumentFragment`{#the-template-element:documentfragment-5}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    object whose [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-template-element:node-document-2
    x-internal="node-document"} is `doc`{.variable} and
    [host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#the-template-element:concept-documentfragment-host
    x-internal="concept-documentfragment-host"} is the
    [`template`{#the-template-element:the-template-element-15}](#the-template-element)
    element.

3.  Set the
    [`template`{#the-template-element:the-template-element-16}](#the-template-element)
    element\'s [template
    contents](#template-contents){#the-template-element:template-contents-5}
    to the newly created
    [`DocumentFragment`{#the-template-element:documentfragment-6}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    object.

A [`Document`{#the-template-element:document}](dom.html#document)
`doc`{.variable}\'s [appropriate template contents owner
document]{#appropriate-template-contents-owner-document .dfn} is the
[`Document`{#the-template-element:document-2}](dom.html#document)
returned by the following algorithm:

1.  If `doc`{.variable} is not a
    [`Document`{#the-template-element:document-3}](dom.html#document)
    created by this algorithm, then:

    1.  If `doc`{.variable} does not yet have an [associated inert
        template document]{#associated-inert-template-document .dfn},
        then:

        1.  Let `new doc`{.variable} be a new
            [`Document`{#the-template-element:document-4}](dom.html#document)
            (whose [browsing
            context](document-sequences.html#concept-document-bc){#the-template-element:concept-document-bc}
            is null). This is \"a
            [`Document`{#the-template-element:document-5}](dom.html#document)
            created by this algorithm\" for the purposes of the step
            above.

        2.  If `doc`{.variable} is an [HTML
            document](https://dom.spec.whatwg.org/#html-document){#the-template-element:html-documents
            x-internal="html-documents"}, mark `new doc`{.variable} as
            an [HTML
            document](https://dom.spec.whatwg.org/#html-document){#the-template-element:html-documents-2
            x-internal="html-documents"} also.

        3.  Set `doc`{.variable}\'s [associated inert template
            document](#associated-inert-template-document){#the-template-element:associated-inert-template-document}
            to `new doc`{.variable}.

    2.  Set `doc`{.variable} to `doc`{.variable}\'s [associated inert
        template
        document](#associated-inert-template-document){#the-template-element:associated-inert-template-document-2}.

    Each
    [`Document`{#the-template-element:document-6}](dom.html#document)
    not created by this algorithm thus gets a single
    [`Document`{#the-template-element:document-7}](dom.html#document) to
    act as its proxy for owning the [template
    contents](#template-contents){#the-template-element:template-contents-6}
    of all its
    [`template`{#the-template-element:the-template-element-17}](#the-template-element)
    elements, so that they aren\'t in a [browsing
    context](document-sequences.html#browsing-context){#the-template-element:browsing-context}
    and thus remain inert (e.g. scripts do not run). Meanwhile,
    [`template`{#the-template-element:the-template-element-18}](#the-template-element)
    elements inside
    [`Document`{#the-template-element:document-8}](dom.html#document)
    objects that *are* created by this algorithm just reuse the same
    [`Document`{#the-template-element:document-9}](dom.html#document)
    owner for their contents.

2.  Return `doc`{.variable}.

The [adopting
steps](https://dom.spec.whatwg.org/#concept-node-adopt-ext){#the-template-element:concept-node-adopt-ext
x-internal="concept-node-adopt-ext"} (with `node`{.variable} and
`oldDocument`{.variable} as parameters) for
[`template`{#the-template-element:the-template-element-19}](#the-template-element)
elements are the following:

1.  Let `doc`{.variable} be `node`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-template-element:node-document-3
    x-internal="node-document"}\'s [appropriate template contents owner
    document](#appropriate-template-contents-owner-document){#the-template-element:appropriate-template-contents-owner-document-2}.

    `node`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-template-element:node-document-4
    x-internal="node-document"} is the
    [`Document`{#the-template-element:document-10}](dom.html#document)
    object that `node`{.variable} was just adopted *into*.

2.  [Adopt](https://dom.spec.whatwg.org/#concept-node-adopt){#the-template-element:concept-node-adopt
    x-internal="concept-node-adopt"} `node`{.variable}\'s [template
    contents](#template-contents){#the-template-element:template-contents-7}
    (a
    [`DocumentFragment`{#the-template-element:documentfragment-7}](https://dom.spec.whatwg.org/#interface-documentfragment){x-internal="documentfragment"}
    object) into `doc`{.variable}.

The [`content`]{#dom-template-content .dfn dfn-for="HTMLTemplateElement"
dfn-type="attribute"} getter steps are to return
[`template`{#the-template-element:the-template-element-20}](#the-template-element)\'s
[template
contents](#template-contents){#the-template-element:template-contents-8},
if the [template
contents](#template-contents){#the-template-element:template-contents-9}
is not a
[`ShadowRoot`{#the-template-element:shadowroot}](https://dom.spec.whatwg.org/#interface-shadowroot){x-internal="shadowroot"}
node; otherwise null.

The [`shadowRootMode`]{#dom-template-shadowrootmode .dfn
dfn-for="HTMLTemplateElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-template-element:reflect}
the
[`shadowrootmode`{#the-template-element:attr-template-shadowrootmode-3}](#attr-template-shadowrootmode)
content attribute, [limited to only known
values](common-dom-interfaces.html#limited-to-only-known-values){#the-template-element:limited-to-only-known-values}.

The [`shadowRootDelegatesFocus`]{#dom-template-shadowrootdelegatesfocus
.dfn dfn-for="HTMLTemplateElement" dfn-type="attribute"} IDL attribute
must
[reflect](common-dom-interfaces.html#reflect){#the-template-element:reflect-2}
the
[`shadowrootdelegatesfocus`{#the-template-element:attr-template-shadowrootdelegatesfocus-2}](#attr-template-shadowrootdelegatesfocus)
content attribute.

The [`shadowRootClonable`]{#dom-template-shadowrootclonable .dfn
dfn-for="HTMLTemplateElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-template-element:reflect-3}
the
[`shadowrootclonable`{#the-template-element:attr-template-shadowrootclonable-2}](#attr-template-shadowrootclonable)
content attribute.

The [`shadowRootSerializable`]{#dom-template-shadowrootserializable .dfn
dfn-for="HTMLTemplateElement" dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-template-element:reflect-4}
the
[`shadowrootserializable`{#the-template-element:attr-template-shadowrootserializable-2}](#attr-template-shadowrootserializable)
content attribute.

The
[`shadowRootCustomElementRegistry`]{#dom-template-shadowrootcustomelementregistry
.dfn dfn-for="HTMLTemplateElement" dfn-type="attribute"} IDL attribute
must
[reflect](common-dom-interfaces.html#reflect){#the-template-element:reflect-5}
the
[`shadowrootcustomelementregistry`{#the-template-element:attr-template-shadowrootcustomelementregistry-2}](#attr-template-shadowrootcustomelementregistry)
content attribute.

The IDL attribute does intentionally not have a boolean type so it can
be extended.

------------------------------------------------------------------------

The [cloning
steps](https://dom.spec.whatwg.org/#concept-node-clone-ext){#the-template-element:concept-node-clone-ext
x-internal="concept-node-clone-ext"} for
[`template`{#the-template-element:the-template-element-21}](#the-template-element)
elements given `node`{.variable}, `copy`{.variable}, and
`subtree`{.variable} are:

1.  If `subtree`{.variable} is false, then return.

2.  For each `child`{.variable} of `node`{.variable}\'s [template
    contents](#template-contents){#the-template-element:template-contents-10}\'s
    [children](https://dom.spec.whatwg.org/#concept-tree-child){#the-template-element:concept-tree-child
    x-internal="concept-tree-child"}, in [tree
    order](https://dom.spec.whatwg.org/#concept-tree-order){#the-template-element:tree-order
    x-internal="tree-order"}: [clone a
    node](https://dom.spec.whatwg.org/#concept-node-clone){#the-template-element:concept-node-clone
    x-internal="concept-node-clone"} given `child`{.variable} with
    [*document*](https://dom.spec.whatwg.org/#clone-a-node-document){#the-template-element:concept-node-clone-document
    x-internal="concept-node-clone-document"} set to
    `copy`{.variable}\'s [template
    contents](#template-contents){#the-template-element:template-contents-11}\'s
    [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#the-template-element:node-document-5
    x-internal="node-document"},
    [*subtree*](https://dom.spec.whatwg.org/#clone-a-node-subtree){#the-template-element:concept-node-clone-subtree
    x-internal="concept-node-clone-subtree"} set to true, and
    [*parent*](https://dom.spec.whatwg.org/#clone-a-node-parent){#the-template-element:concept-node-clone-parent
    x-internal="concept-node-clone-parent"} set to `copy`{.variable}\'s
    [template
    contents](#template-contents){#the-template-element:template-contents-12}.

::: example
In this example, a script populates a table four-column with data from a
data structure, using a
[`template`{#the-template-element:the-template-element-22}](#the-template-element)
to provide the element structure instead of manually generating the
structure from markup.

``` html
<!DOCTYPE html>
<html lang='en'>
<title>Cat data</title>
<script>
 // Data is hard-coded here, but could come from the server
 var data = [
   { name: 'Pillar', color: 'Ticked Tabby', sex: 'Female (neutered)', legs: 3 },
   { name: 'Hedral', color: 'Tuxedo', sex: 'Male (neutered)', legs: 4 },
 ];
</script>
<table>
 <thead>
  <tr>
   <th>Name <th>Color <th>Sex <th>Legs
 <tbody>
  <template id="row">
   <tr><td><td><td><td>
  </template>
</table>
<script>
 var template = document.querySelector('#row');
 for (var i = 0; i < data.length; i += 1) {
   var cat = data[i];
   var clone = template.content.cloneNode(true);
   var cells = clone.querySelectorAll('td');
   cells[0].textContent = cat.name;
   cells[1].textContent = cat.color;
   cells[2].textContent = cat.sex;
   cells[3].textContent = cat.legs;
   template.parentNode.appendChild(clone);
 }
</script>
```

This example uses
[`cloneNode()`{#the-template-element:dom-node-clonenode}](https://dom.spec.whatwg.org/#dom-node-clonenode){x-internal="dom-node-clonenode"}
on the
[`template`{#the-template-element:the-template-element-23}](#the-template-element)\'s
contents; it could equivalently have used
[`document.importNode()`{#the-template-element:dom-document-importnode}](https://dom.spec.whatwg.org/#dom-document-importnode){x-internal="dom-document-importnode"},
which does the same thing. The only difference between these two APIs is
when the [node
document](https://dom.spec.whatwg.org/#concept-node-document){#the-template-element:node-document-6
x-internal="node-document"} is updated: with
[`cloneNode()`{#the-template-element:dom-node-clonenode-2}](https://dom.spec.whatwg.org/#dom-node-clonenode){x-internal="dom-node-clonenode"}
it is updated when the nodes are appended with
[`appendChild()`{#the-template-element:dom-node-appendchild-2}](https://dom.spec.whatwg.org/#dom-node-appendchild){x-internal="dom-node-appendchild"},
with
[`document.importNode()`{#the-template-element:dom-document-importnode-2}](https://dom.spec.whatwg.org/#dom-document-importnode){x-internal="dom-document-importnode"}
it is updated when the nodes are cloned.
:::

##### [4.12.3.1]{.secno} Interaction of [`template`{#template-XSLT-XPath:the-template-element}](#the-template-element) elements with XSLT and XPath[](#template-XSLT-XPath){.self-link} {#template-XSLT-XPath}

*This section is non-normative.*

This specification does not define how XSLT and XPath interact with the
[`template`{#template-XSLT-XPath:the-template-element-2}](#the-template-element)
element. However, in the absence of another specification actually
defining this, here are some guidelines for implementers, which are
intended to be consistent with other processing described in this
specification:

- An XSLT processor based on an XML parser that acts [as described in
  this
  specification](xhtml.html#xml-parser){#template-XSLT-XPath:xml-parser}
  needs to act as if
  [`template`{#template-XSLT-XPath:the-template-element-3}](#the-template-element)
  elements contain as descendants their [template
  contents](#template-contents){#template-XSLT-XPath:template-contents}
  for the purposes of the transform.

- An XSLT processor that outputs a DOM needs to ensure that nodes that
  would go into a
  [`template`{#template-XSLT-XPath:the-template-element-4}](#the-template-element)
  element are instead placed into the element\'s [template
  contents](#template-contents){#template-XSLT-XPath:template-contents-2}.

- XPath evaluation using the XPath DOM API when applied to a
  [`Document`{#template-XSLT-XPath:document}](dom.html#document) parsed
  using the [HTML
  parser](parsing.html#html-parser){#template-XSLT-XPath:html-parser} or
  the [XML
  parser](xhtml.html#xml-parser){#template-XSLT-XPath:xml-parser-2}
  described in this specification needs to ignore [template
  contents](#template-contents){#template-XSLT-XPath:template-contents-3}.

#### [4.12.4]{.secno} The [`slot`]{.dfn dfn-type="element"} element[](#the-slot-element){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Element/slot](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot "The <slot> HTML element—part of the Web Components technology suite—is a placeholder inside a web component that you can fill with your own markup, which lets you create separate DOM trees and present them together.")

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
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSlotElement](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSlotElement "The HTMLSlotElement interface of the Shadow DOM API enables access to the name and assigned nodes of an HTML <slot> element.")

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
:::::

[Categories](dom.html#concept-element-categories){#the-slot-element:concept-element-categories}:
:   [Flow
    content](dom.html#flow-content-2){#the-slot-element:flow-content-2}.
:   [Phrasing
    content](dom.html#phrasing-content-2){#the-slot-element:phrasing-content-2}.

[Contexts in which this element can be used](dom.html#concept-element-contexts){#the-slot-element:concept-element-contexts}:
:   Where [phrasing
    content](dom.html#phrasing-content-2){#the-slot-element:phrasing-content-2-2}
    is expected.

[Content model](dom.html#concept-element-content-model){#the-slot-element:concept-element-content-model}:
:   [Transparent](dom.html#transparent){#the-slot-element:transparent}

[Tag omission in text/html](dom.html#concept-element-tag-omission){#the-slot-element:concept-element-tag-omission}:
:   Neither tag is omissible.

[Content attributes](dom.html#concept-element-attributes){#the-slot-element:concept-element-attributes}:
:   [Global
    attributes](dom.html#global-attributes){#the-slot-element:global-attributes}
:   [`name`{#the-slot-element:attr-slot-name}](#attr-slot-name) --- Name
    of shadow tree slot

[Accessibility considerations](dom.html#concept-element-accessibility-considerations){#the-slot-element:concept-element-accessibility-considerations}:
:   [For authors](https://w3c.github.io/html-aria/#el-slot).
:   [For implementers](https://w3c.github.io/html-aam/#el-slot).

[DOM interface](dom.html#concept-element-dom){#the-slot-element:concept-element-dom}:

:   ``` idl
    [Exposed=Window]
    interface HTMLSlotElement : HTMLElement {
      [HTMLConstructor] constructor();

      [CEReactions] attribute DOMString name;
      sequence<Node> assignedNodes(optional AssignedNodesOptions options = {});
      sequence<Element> assignedElements(optional AssignedNodesOptions options = {});
      undefined assign((Element or Text)... nodes);
    };

    dictionary AssignedNodesOptions {
      boolean flatten = false;
    };
    ```

The [`slot`{#the-slot-element:the-slot-element}](#the-slot-element)
element defines a
[slot](https://dom.spec.whatwg.org/#concept-slot){#the-slot-element:concept-slot
x-internal="concept-slot"}. It is typically used in a [shadow
tree](https://dom.spec.whatwg.org/#concept-shadow-tree){#the-slot-element:shadow-tree
x-internal="shadow-tree"}. A
[`slot`{#the-slot-element:the-slot-element-2}](#the-slot-element)
element [represents](dom.html#represents){#the-slot-element:represents}
its [assigned
nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-slot-element:assigned-nodes
x-internal="assigned-nodes"}, if any, and its contents otherwise.

The [`name`]{#attr-slot-name .dfn dfn-for="slot"
dfn-type="element-attr"} content attribute may contain any string value.
It represents a
[slot](https://dom.spec.whatwg.org/#concept-slot){#the-slot-element:concept-slot-2
x-internal="concept-slot"}\'s
[name](https://dom.spec.whatwg.org/#slot-name){#the-slot-element:slot-name
x-internal="slot-name"}.

The [`name`{#the-slot-element:attr-slot-name-2}](#attr-slot-name)
attribute is used to [assign
slots](https://dom.spec.whatwg.org/#assign-a-slot){#the-slot-element:assign-a-slot
x-internal="assign-a-slot"} to other elements: a
[`slot`{#the-slot-element:the-slot-element-3}](#the-slot-element)
element with a
[`name`{#the-slot-element:attr-slot-name-3}](#attr-slot-name) attribute
creates a named
[slot](https://dom.spec.whatwg.org/#concept-slot){#the-slot-element:concept-slot-3
x-internal="concept-slot"} to which any element is
[assigned](https://dom.spec.whatwg.org/#assign-a-slot){#the-slot-element:assign-a-slot-2
x-internal="assign-a-slot"} if that element has a
[`slot`{#the-slot-element:attr-slot}](dom.html#attr-slot) attribute
whose value matches that
[`name`{#the-slot-element:attr-slot-name-4}](#attr-slot-name)
attribute\'s value, and the
[`slot`{#the-slot-element:the-slot-element-4}](#the-slot-element)
element is a child of the [shadow
tree](https://dom.spec.whatwg.org/#concept-shadow-tree){#the-slot-element:shadow-tree-2
x-internal="shadow-tree"} whose
[root](https://dom.spec.whatwg.org/#concept-tree-root){#the-slot-element:root
x-internal="root"}\'s
[host](https://dom.spec.whatwg.org/#concept-documentfragment-host){#the-slot-element:concept-documentfragment-host
x-internal="concept-documentfragment-host"} has that corresponding
[`slot`{#the-slot-element:attr-slot-2}](dom.html#attr-slot) attribute
value.

`slot`{.variable}`.`[`name`](#dom-slot-name){#dom-slot-name-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSlotElement/name](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSlotElement/name "The name property of the HTMLSlotElement interface returns or sets the slot name. A slot is a placeholder inside a web component that users can fill with their own markup.")

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
:::::

Can be used to get and set `slot`{.variable}\'s
[name](https://dom.spec.whatwg.org/#slot-name){#the-slot-element:slot-name-2
x-internal="slot-name"}.

`slot`{.variable}`.`[`assignedNodes`](#dom-slot-assignednodes){#dom-slot-assignednodes-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSlotElement/assignedNodes](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSlotElement/assignedNodes "The assignedNodes() method of the HTMLSlotElement interface returns a sequence of the nodes assigned to this slot.")

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
:::::

Returns `slot`{.variable}\'s [assigned
nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-slot-element:assigned-nodes-2
x-internal="assigned-nodes"}.

`slot`{.variable}`.`[`assignedNodes`](#dom-slot-assignednodes){#the-slot-element:dom-slot-assignednodes-2}`({ flatten: true })`

Returns `slot`{.variable}\'s [assigned
nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-slot-element:assigned-nodes-3
x-internal="assigned-nodes"}, if any, and `slot`{.variable}\'s children
otherwise, and does the same for any
[`slot`{#the-slot-element:the-slot-element-5}](#the-slot-element)
elements encountered therein, recursively, until there are no
[`slot`{#the-slot-element:the-slot-element-6}](#the-slot-element)
elements left.

`slot`{.variable}`.`[`assignedElements`](#dom-slot-assignedelements){#dom-slot-assignedelements-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSlotElement/assignedElements](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSlotElement/assignedElements "The assignedElements() method of the HTMLSlotElement interface returns a sequence of the elements assigned to this slot (and no other nodes).")

Support in all current engines.

::: support
[Firefox66+]{.firefox .yes}[Safari12.1+]{.safari
.yes}[Chrome65+]{.chrome .yes}

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

Returns `slot`{.variable}\'s [assigned
nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-slot-element:assigned-nodes-4
x-internal="assigned-nodes"}, limited to elements.

`slot`{.variable}`.`[`assignedElements`](#dom-slot-assignedelements){#the-slot-element:dom-slot-assignedelements-2}`({ flatten: true })`

Returns the same as
[`assignedNodes({ flatten: true })`{#the-slot-element:dom-slot-assignednodes-3}](#dom-slot-assignednodes),
limited to elements.

`slot`{.variable}`.`[`assign`](#dom-slot-assign){#the-slot-element:dom-slot-assign-2}`(...``nodes`{.variable}`)`

Sets `slot`{.variable}\'s [manually assigned
nodes](#manually-assigned-nodes){#the-slot-element:manually-assigned-nodes}
to the given `nodes`{.variable}.

The [`name`]{#dom-slot-name .dfn dfn-for="HTMLSlotElement"
dfn-type="attribute"} IDL attribute must
[reflect](common-dom-interfaces.html#reflect){#the-slot-element:reflect}
the content attribute of the same name.

The [`slot`{#the-slot-element:the-slot-element-7}](#the-slot-element)
element has [manually assigned nodes]{#manually-assigned-nodes .dfn
dfn-for="HTMLSlotElement" export=""}, which is an [ordered
set](https://infra.spec.whatwg.org/#ordered-set){#the-slot-element:set
x-internal="set"} of
[slottables](https://dom.spec.whatwg.org/#concept-slotable){#the-slot-element:slottable
x-internal="slottable"} set by
[`assign()`{#the-slot-element:dom-slot-assign-3}](#dom-slot-assign).
This set is initially empty.

The [manually assigned
nodes](#manually-assigned-nodes){#the-slot-element:manually-assigned-nodes-2}
set can be implemented using weak references to the
[slottables](https://dom.spec.whatwg.org/#concept-slotable){#the-slot-element:slottable-2
x-internal="slottable"}, because this set is not directly accessible
from script.

The [`assignedNodes(``options`{.variable}`)`]{#dom-slot-assignednodes
.dfn dfn-for="HTMLSlotElement" dfn-type="method"} method steps are:

1.  If
    `options`{.variable}\[\"[`flatten`{#the-slot-element:dom-assignednodesoptions-flatten}](#dom-assignednodesoptions-flatten)\"\]
    is false, then return
    [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this
    x-internal="this"}\'s [assigned
    nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-slot-element:assigned-nodes-5
    x-internal="assigned-nodes"}.

2.  Return the result of [finding flattened
    slottables](https://dom.spec.whatwg.org/#find-flattened-slotables){#the-slot-element:finding-flattened-slottables
    x-internal="finding-flattened-slottables"} with
    [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this-2
    x-internal="this"}.

The
[`assignedElements(``options`{.variable}`)`]{#dom-slot-assignedelements
.dfn dfn-for="HTMLSlotElement" dfn-type="method"} method steps are:

1.  If
    `options`{.variable}\[\"[`flatten`{#the-slot-element:dom-assignednodesoptions-flatten-2}](#dom-assignednodesoptions-flatten)\"\]
    is false, then return
    [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this-3
    x-internal="this"}\'s [assigned
    nodes](https://dom.spec.whatwg.org/#slot-assigned-nodes){#the-slot-element:assigned-nodes-6
    x-internal="assigned-nodes"}, filtered to contain only
    [`Element`{#the-slot-element:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    nodes.

2.  Return the result of [finding flattened
    slottables](https://dom.spec.whatwg.org/#find-flattened-slotables){#the-slot-element:finding-flattened-slottables-2
    x-internal="finding-flattened-slottables"} with
    [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this-4
    x-internal="this"}, filtered to contain only
    [`Element`{#the-slot-element:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    nodes.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLSlotElement/assign](https://developer.mozilla.org/en-US/docs/Web/API/HTMLSlotElement/assign "The assign() method of the HTMLSlotElement interface sets the slot's manually assigned nodes to an ordered set of slottables. The manually assigned nodes set is initially empty until nodes are assigned using assign().")

Support in all current engines.

::: support
[Firefox92+]{.firefox .yes}[Safari16.4+]{.safari
.yes}[Chrome86+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge86+]{.edge_blink .yes}

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

The [`assign(...``nodes`{.variable}`)`]{#dom-slot-assign .dfn
dfn-for="HTMLSlotElement" dfn-type="method"} method steps are:

1.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-slot-element:list-iterate
    x-internal="list-iterate"} `node`{.variable} of
    [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this-5
    x-internal="this"}\'s [manually assigned
    nodes](#manually-assigned-nodes){#the-slot-element:manually-assigned-nodes-3},
    set `node`{.variable}\'s [manual slot
    assignment](https://dom.spec.whatwg.org/#slottable-manual-slot-assignment){#the-slot-element:manual-slot-assignment
    x-internal="manual-slot-assignment"} to null.

2.  Let `nodesSet`{.variable} be a new [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#the-slot-element:set-2
    x-internal="set"}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#the-slot-element:list-iterate-2
    x-internal="list-iterate"} `node`{.variable} of `nodes`{.variable}:

    1.  If `node`{.variable}\'s [manual slot
        assignment](https://dom.spec.whatwg.org/#slottable-manual-slot-assignment){#the-slot-element:manual-slot-assignment-2
        x-internal="manual-slot-assignment"} refers to a
        [slot](#the-slot-element){#the-slot-element:the-slot-element-8},
        then remove `node`{.variable} from that
        [slot](#the-slot-element){#the-slot-element:the-slot-element-9}\'s
        [manually assigned
        nodes](#manually-assigned-nodes){#the-slot-element:manually-assigned-nodes-4}.

    2.  Set `node`{.variable}\'s [manual slot
        assignment](https://dom.spec.whatwg.org/#slottable-manual-slot-assignment){#the-slot-element:manual-slot-assignment-3
        x-internal="manual-slot-assignment"} to
        [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this-6
        x-internal="this"}.

    3.  [Append](https://infra.spec.whatwg.org/#set-append){#the-slot-element:set-append
        x-internal="set-append"} `node`{.variable} to
        `nodesSet`{.variable}.

4.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this-7
    x-internal="this"}\'s [manually assigned
    nodes](#manually-assigned-nodes){#the-slot-element:manually-assigned-nodes-5}
    to `nodesSet`{.variable}.

5.  Run [assign slottables for a
    tree](https://dom.spec.whatwg.org/#assign-slotables-for-a-tree){#the-slot-element:assign-slottables-for-a-tree
    x-internal="assign-slottables-for-a-tree"} for
    [this](https://webidl.spec.whatwg.org/#this){#the-slot-element:this-8
    x-internal="this"}\'s
    [root](https://dom.spec.whatwg.org/#concept-tree-root){#the-slot-element:root-2
    x-internal="root"}.

[← 4.11 Interactive elements](interactive-elements.html) --- [Table of
Contents](./) --- [4.12.5 The canvas element →](canvas.html)
