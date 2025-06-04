HTML Standard

::: {#head .head .with-buttons}
[![WHATWG](https://resources.whatwg.org/logo.svg){.darkmode-aware
width="100" crossorigin="" height="100"}](https://whatwg.org/){.logo}

# HTML {#html .allcaps}

Living Standard --- Last Updated [2 June 2025]{.pubdate}
:::

[← 7 Loading web pages](browsers.html) --- [Table of Contents](./) ---
[7.3 Infrastructure for sequences of documents
→](document-sequences.html)

1.  ::: {#toc-browsers}
    1.  [[7.2]{.secno} APIs related to navigation and session
        history](nav-history-apis.html#nav-traversal-apis)
        1.  [[7.2.1]{.secno} Security infrastructure for `Window`,
            `WindowProxy`, and `Location`
            objects](nav-history-apis.html#cross-origin-objects)
            1.  [[7.2.1.1]{.secno} Integration with
                IDL](nav-history-apis.html#integration-with-idl)
            2.  [[7.2.1.2]{.secno} Shared internal slot:
                \[\[CrossOriginPropertyDescriptorMap\]\]](nav-history-apis.html#shared-internal-slot:-crossoriginpropertydescriptormap)
            3.  [[7.2.1.3]{.secno} Shared abstract
                operations](nav-history-apis.html#shared-abstract-operations)
                1.  [[7.2.1.3.1]{.secno} CrossOriginProperties (
                    `O`{.variable}
                    )](nav-history-apis.html#crossoriginproperties-(-o-))
                2.  [[7.2.1.3.2]{.secno} CrossOriginPropertyFallback (
                    `P`{.variable}
                    )](nav-history-apis.html#crossoriginpropertyfallback-(-p-))
                3.  [[7.2.1.3.3]{.secno} IsPlatformObjectSameOrigin (
                    `O`{.variable}
                    )](nav-history-apis.html#isplatformobjectsameorigin-(-o-))
                4.  [[7.2.1.3.4]{.secno} CrossOriginGetOwnPropertyHelper
                    ( `O`{.variable}, `P`{.variable}
                    )](nav-history-apis.html#crossorigingetownpropertyhelper-(-o,-p-))
                5.  [[7.2.1.3.5]{.secno} CrossOriginGet (
                    `O`{.variable}, `P`{.variable},
                    `Receiver`{.variable}
                    )](nav-history-apis.html#crossoriginget-(-o,-p,-receiver-))
                6.  [[7.2.1.3.6]{.secno} CrossOriginSet (
                    `O`{.variable}, `P`{.variable}, `V`{.variable},
                    `Receiver`{.variable}
                    )](nav-history-apis.html#crossoriginset-(-o,-p,-v,-receiver-))
                7.  [[7.2.1.3.7]{.secno} CrossOriginOwnPropertyKeys (
                    `O`{.variable}
                    )](nav-history-apis.html#crossoriginownpropertykeys-(-o-))
        2.  [[7.2.2]{.secno} The `Window`
            object](nav-history-apis.html#the-window-object)
            1.  [[7.2.2.1]{.secno} Opening and closing
                windows](nav-history-apis.html#apis-for-creating-and-navigating-browsing-contexts-by-name)
            2.  [[7.2.2.2]{.secno} Indexed access on the `Window`
                object](nav-history-apis.html#accessing-other-browsing-contexts)
            3.  [[7.2.2.3]{.secno} Named access on the `Window`
                object](nav-history-apis.html#named-access-on-the-window-object)
            4.  [[7.2.2.4]{.secno} Accessing related
                windows](nav-history-apis.html#navigating-nested-browsing-contexts-in-the-dom)
            5.  [[7.2.2.5]{.secno} Historical browser interface element
                APIs](nav-history-apis.html#browser-interface-elements)
            6.  [[7.2.2.6]{.secno} Script settings for `Window`
                objects](nav-history-apis.html#script-settings-for-window-objects)
        3.  [[7.2.3]{.secno} The `WindowProxy` exotic
            object](nav-history-apis.html#the-windowproxy-exotic-object)
            1.  [[7.2.3.1]{.secno} \[\[GetPrototypeOf\]\] (
                )](nav-history-apis.html#windowproxy-getprototypeof)
            2.  [[7.2.3.2]{.secno} \[\[SetPrototypeOf\]\] (
                `V`{.variable}
                )](nav-history-apis.html#windowproxy-setprototypeof)
            3.  [[7.2.3.3]{.secno} \[\[IsExtensible\]\] (
                )](nav-history-apis.html#windowproxy-isextensible)
            4.  [[7.2.3.4]{.secno} \[\[PreventExtensions\]\] (
                )](nav-history-apis.html#windowproxy-preventextensions)
            5.  [[7.2.3.5]{.secno} \[\[GetOwnProperty\]\] (
                `P`{.variable}
                )](nav-history-apis.html#windowproxy-getownproperty)
            6.  [[7.2.3.6]{.secno} \[\[DefineOwnProperty\]\] (
                `P`{.variable}, `Desc`{.variable}
                )](nav-history-apis.html#windowproxy-defineownproperty)
            7.  [[7.2.3.7]{.secno} \[\[Get\]\] ( `P`{.variable},
                `Receiver`{.variable}
                )](nav-history-apis.html#windowproxy-get)
            8.  [[7.2.3.8]{.secno} \[\[Set\]\] ( `P`{.variable},
                `V`{.variable}, `Receiver`{.variable}
                )](nav-history-apis.html#windowproxy-set)
            9.  [[7.2.3.9]{.secno} \[\[Delete\]\] ( `P`{.variable}
                )](nav-history-apis.html#windowproxy-delete)
            10. [[7.2.3.10]{.secno} \[\[OwnPropertyKeys\]\] (
                )](nav-history-apis.html#windowproxy-ownpropertykeys)
        4.  [[7.2.4]{.secno} The `Location`
            interface](nav-history-apis.html#the-location-interface)
            1.  [[7.2.4.1]{.secno} \[\[GetPrototypeOf\]\] (
                )](nav-history-apis.html#location-getprototypeof)
            2.  [[7.2.4.2]{.secno} \[\[SetPrototypeOf\]\] (
                `V`{.variable}
                )](nav-history-apis.html#location-setprototypeof)
            3.  [[7.2.4.3]{.secno} \[\[IsExtensible\]\] (
                )](nav-history-apis.html#location-isextensible)
            4.  [[7.2.4.4]{.secno} \[\[PreventExtensions\]\] (
                )](nav-history-apis.html#location-preventextensions)
            5.  [[7.2.4.5]{.secno} \[\[GetOwnProperty\]\] (
                `P`{.variable}
                )](nav-history-apis.html#location-getownproperty)
            6.  [[7.2.4.6]{.secno} \[\[DefineOwnProperty\]\] (
                `P`{.variable}, `Desc`{.variable}
                )](nav-history-apis.html#location-defineownproperty)
            7.  [[7.2.4.7]{.secno} \[\[Get\]\] ( `P`{.variable},
                `Receiver`{.variable}
                )](nav-history-apis.html#location-get)
            8.  [[7.2.4.8]{.secno} \[\[Set\]\] ( `P`{.variable},
                `V`{.variable}, `Receiver`{.variable}
                )](nav-history-apis.html#location-set)
            9.  [[7.2.4.9]{.secno} \[\[Delete\]\] ( `P`{.variable}
                )](nav-history-apis.html#location-delete)
            10. [[7.2.4.10]{.secno} \[\[OwnPropertyKeys\]\] (
                )](nav-history-apis.html#location-ownpropertykeys)
        5.  [[7.2.5]{.secno} The `History`
            interface](nav-history-apis.html#the-history-interface)
        6.  [[7.2.6]{.secno} The navigation
            API](nav-history-apis.html#navigation-api)
            1.  [[7.2.6.1]{.secno}
                Introduction](nav-history-apis.html#navigation-api-intro)
            2.  [[7.2.6.2]{.secno} The `Navigation`
                interface](nav-history-apis.html#navigation-interface)
            3.  [[7.2.6.3]{.secno} Core
                infrastructure](nav-history-apis.html#navigation-api-core)
            4.  [[7.2.6.4]{.secno} Initializing and updating the entry
                list](nav-history-apis.html#navigation-api-entry-updates)
            5.  [[7.2.6.5]{.secno} The `NavigationHistoryEntry`
                interface](nav-history-apis.html#the-navigationhistoryentry-interface)
            6.  [[7.2.6.6]{.secno} The history entry
                list](nav-history-apis.html#the-history-entry-list)
            7.  [[7.2.6.7]{.secno} Initiating
                navigations](nav-history-apis.html#navigation-api-initiating-navigations)
            8.  [[7.2.6.8]{.secno} Ongoing navigation
                tracking](nav-history-apis.html#ongoing-navigation-tracking)
            9.  [[7.2.6.9]{.secno} The `NavigationActivation`
                interface](nav-history-apis.html#navigation-activation-interface)
            10. [[7.2.6.10]{.secno} The `navigate`
                event](nav-history-apis.html#the-navigate-event)
                1.  [[7.2.6.10.1]{.secno} The `NavigateEvent`
                    interface](nav-history-apis.html#the-navigateevent-interface)
                2.  [[7.2.6.10.2]{.secno} The `NavigationDestination`
                    interface](nav-history-apis.html#the-navigationdestination-interface)
                3.  [[7.2.6.10.3]{.secno} Firing the
                    event](nav-history-apis.html#navigate-event-firing)
                4.  [[7.2.6.10.4]{.secno} Scroll and focus
                    behavior](nav-history-apis.html#navigate-event-scroll-focus)
        7.  [[7.2.7]{.secno} Event
            interfaces](nav-history-apis.html#nav-traversal-event-interfaces)
            1.  [[7.2.7.1]{.secno} The
                `NavigationCurrentEntryChangeEvent`
                interface](nav-history-apis.html#the-navigationcurrententrychangeevent-interface)
            2.  [[7.2.7.2]{.secno} The `PopStateEvent`
                interface](nav-history-apis.html#the-popstateevent-interface)
            3.  [[7.2.7.3]{.secno} The `HashChangeEvent`
                interface](nav-history-apis.html#the-hashchangeevent-interface)
            4.  [[7.2.7.4]{.secno} The `PageSwapEvent`
                interface](nav-history-apis.html#the-pageswapevent-interface)
            5.  [[7.2.7.5]{.secno} The `PageRevealEvent`
                interface](nav-history-apis.html#the-pagerevealevent-interface)
            6.  [[7.2.7.6]{.secno} The `PageTransitionEvent`
                interface](nav-history-apis.html#the-pagetransitionevent-interface)
            7.  [[7.2.7.7]{.secno} The `BeforeUnloadEvent`
                interface](nav-history-apis.html#the-beforeunloadevent-interface)
        8.  [[7.2.8]{.secno} The `NotRestoredReasons`
            interface](nav-history-apis.html#the-notrestoredreasons-interface)
    :::

### [7.2]{.secno} APIs related to navigation and session history[](#nav-traversal-apis){.self-link} {#nav-traversal-apis}

#### [7.2.1]{.secno} Security infrastructure for [`Window`{#cross-origin-objects:window}](#window), [`WindowProxy`{#cross-origin-objects:windowproxy}](#windowproxy), and [`Location`{#cross-origin-objects:location}](#location) objects[](#cross-origin-objects){.self-link} {#cross-origin-objects}

Although typically objects cannot be accessed across
[origins](browsers.html#concept-origin){#cross-origin-objects:concept-origin},
the web platform would not be true to itself if it did not have some
legacy exceptions to that rule that the web depends upon.

This section uses the terminology and typographic conventions from the
JavaScript specification.
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

##### [7.2.1.1]{.secno} Integration with IDL[](#integration-with-idl){.self-link}

When [perform a security
check](https://webidl.spec.whatwg.org/#dfn-perform-a-security-check){#integration-with-idl:perform-a-security-check
x-internal="perform-a-security-check"} is invoked, with a
`platformObject`{.variable}, `identifier`{.variable}, and
`type`{.variable}, run these steps:

1.  If `platformObject`{.variable} is not a
    [`Window`{#integration-with-idl:window}](#window) or
    [`Location`{#integration-with-idl:location}](#location) object, then
    return.

2.  For each `e`{.variable} of
    [CrossOriginProperties](#crossoriginproperties-(-o-)){#integration-with-idl:crossoriginproperties-(-o-)}(`platformObject`{.variable}):

    1.  If
        [SameValue](https://tc39.es/ecma262/#sec-samevalue){#integration-with-idl:samevalue
        x-internal="samevalue"}(`e`{.variable}.\[\[Property\]\],
        `identifier`{.variable}) is true, then:

        1.  If `type`{.variable} is \"`method`\" and `e`{.variable} has
            neither \[\[NeedsGet\]\] nor \[\[NeedsSet\]\], then return.

        2.  Otherwise, if `type`{.variable} is \"`getter`\" and
            `e`{.variable}.\[\[NeedsGet\]\] is true, then return.

        3.  Otherwise, if `type`{.variable} is \"`setter`\" and
            `e`{.variable}.\[\[NeedsSet\]\] is true, then return.

3.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#integration-with-idl:isplatformobjectsameorigin-(-o-)}(`platformObject`{.variable})
    is false, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#integration-with-idl:securityerror
    x-internal="securityerror"}
    [`DOMException`{#integration-with-idl:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

##### [7.2.1.2]{.secno} Shared internal slot: \[\[CrossOriginPropertyDescriptorMap\]\][](#shared-internal-slot:-crossoriginpropertydescriptormap){.self-link} {#shared-internal-slot:-crossoriginpropertydescriptormap}

[`Window`{#shared-internal-slot:-crossoriginpropertydescriptormap:window}](#window)
and
[`Location`{#shared-internal-slot:-crossoriginpropertydescriptormap:location}](#location)
objects both have a
[\[\[CrossOriginPropertyDescriptorMap\]\]]{#crossoriginpropertydescriptormap
.dfn} internal slot, whose value is initially an empty map.

The
[\[\[CrossOriginPropertyDescriptorMap\]\]](#crossoriginpropertydescriptormap){#shared-internal-slot:-crossoriginpropertydescriptormap:crossoriginpropertydescriptormap}
internal slot contains a map with entries whose keys are
(`currentGlobal`{.variable}, `objectGlobal`{.variable},
`propertyKey`{.variable})-tuples and values are property descriptors, as
a memoization of what is visible to scripts when
`currentGlobal`{.variable} inspects a
[`Window`{#shared-internal-slot:-crossoriginpropertydescriptormap:window-2}](#window)
or
[`Location`{#shared-internal-slot:-crossoriginpropertydescriptormap:location-2}](#location)
object from `objectGlobal`{.variable}. It is filled lazily by
[CrossOriginGetOwnPropertyHelper](#crossorigingetownpropertyhelper-(-o,-p-)){#shared-internal-slot:-crossoriginpropertydescriptormap:crossorigingetownpropertyhelper-(-o,-p-)},
which consults it on future lookups.

User agents should allow a value held in the map to be garbage collected
along with its corresponding key when nothing holds a reference to any
part of the value. That is, as long as garbage collection is not
observable.

For example, with
`const href = Object.getOwnPropertyDescriptor(crossOriginLocation, "href").set`
the value and its corresponding key in the map cannot be garbage
collected as that would be observable.

User agents may have an optimization whereby they remove key-value pairs
from the map when
[`document.domain`{#shared-internal-slot:-crossoriginpropertydescriptormap:dom-document-domain}](browsers.html#dom-document-domain)
is set. This is not observable as
[`document.domain`{#shared-internal-slot:-crossoriginpropertydescriptormap:dom-document-domain-2}](browsers.html#dom-document-domain)
cannot revisit an earlier value.

For example, setting
[`document.domain`{#shared-internal-slot:-crossoriginpropertydescriptormap:dom-document-domain-3}](browsers.html#dom-document-domain)
to \"`example.com`\" on www.example.com means user agents can remove all
key-value pairs from the map where part of the key is www.example.com,
as that can never be part of the
[origin](browsers.html#concept-origin){#shared-internal-slot:-crossoriginpropertydescriptormap:concept-origin}
again and therefore the corresponding value could never be retrieved
from the map.

##### [7.2.1.3]{.secno} Shared abstract operations[](#shared-abstract-operations){.self-link}

###### [7.2.1.3.1]{.secno} [CrossOriginProperties]{.dfn} ( `O`{.variable} )[](#crossoriginproperties-(-o-)){.self-link} {#crossoriginproperties-(-o-)}

1.  [Assert](https://infra.spec.whatwg.org/#assert){#crossoriginproperties-(-o-):assert
    x-internal="assert"}: `O`{.variable} is a
    [`Location`{#crossoriginproperties-(-o-):location}](#location) or
    [`Window`{#crossoriginproperties-(-o-):window}](#window) object.

2.  If `O`{.variable} is a
    [`Location`{#crossoriginproperties-(-o-):location-2}](#location)
    object, then return « { \[\[Property\]\]: \"`href`\",
    \[\[NeedsGet\]\]: false, \[\[NeedsSet\]\]: true }, {
    \[\[Property\]\]: \"`replace`\" } ».

3.  Return « { \[\[Property\]\]: \"`window`\", \[\[NeedsGet\]\]: true,
    \[\[NeedsSet\]\]: false }, { \[\[Property\]\]: \"`self`\",
    \[\[NeedsGet\]\]: true, \[\[NeedsSet\]\]: false }, {
    \[\[Property\]\]: \"`location`\", \[\[NeedsGet\]\]: true,
    \[\[NeedsSet\]\]: true }, { \[\[Property\]\]: \"`close`\" }, {
    \[\[Property\]\]: \"`closed`\", \[\[NeedsGet\]\]: true,
    \[\[NeedsSet\]\]: false }, { \[\[Property\]\]: \"`focus`\" }, {
    \[\[Property\]\]: \"`blur`\" }, { \[\[Property\]\]: \"`frames`\",
    \[\[NeedsGet\]\]: true, \[\[NeedsSet\]\]: false }, {
    \[\[Property\]\]: \"`length`\", \[\[NeedsGet\]\]: true,
    \[\[NeedsSet\]\]: false }, { \[\[Property\]\]: \"`top`\",
    \[\[NeedsGet\]\]: true, \[\[NeedsSet\]\]: false }, {
    \[\[Property\]\]: \"`opener`\", \[\[NeedsGet\]\]: true,
    \[\[NeedsSet\]\]: false }, { \[\[Property\]\]: \"`parent`\",
    \[\[NeedsGet\]\]: true, \[\[NeedsSet\]\]: false }, {
    \[\[Property\]\]: \"`postMessage`\" } ».

This abstract operation does not return a [Completion
Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#crossoriginproperties-(-o-):completion-record
x-internal="completion-record"}.

Indexed properties do not need to be safelisted in this algorithm, as
they are handled directly by the
[`WindowProxy`{#crossoriginproperties-(-o-):windowproxy}](#windowproxy)
object.

A JavaScript property name `P`{.variable} is a [cross-origin accessible
window property name]{#cross-origin-accessible-window-property-name
.dfn} if it is \"`window`\", \"`self`\", \"`location`\", \"`close`\",
\"`closed`\", \"`focus`\", \"`blur`\", \"`frames`\", \"`length`\",
\"`top`\", \"`opener`\", \"`parent`\", \"`postMessage`\", or an [array
index property
name](https://webidl.spec.whatwg.org/#dfn-array-index-property-name){#crossoriginproperties-(-o-):array-index-property-name
x-internal="array-index-property-name"}.

###### [7.2.1.3.2]{.secno} [CrossOriginPropertyFallback]{.dfn} ( `P`{.variable} )[](#crossoriginpropertyfallback-(-p-)){.self-link} {#crossoriginpropertyfallback-(-p-)}

1.  If `P`{.variable} is \"`then`\",
    [%Symbol.toStringTag%](infrastructure.html#symbol.tostringtag){#crossoriginpropertyfallback-(-p-):symbol.tostringtag},
    [%Symbol.hasInstance%](infrastructure.html#symbol.hasinstance){#crossoriginpropertyfallback-(-p-):symbol.hasinstance},
    or
    [%Symbol.isConcatSpreadable%](infrastructure.html#symbol.isconcatspreadable){#crossoriginpropertyfallback-(-p-):symbol.isconcatspreadable},
    then return
    [PropertyDescriptor](https://tc39.es/ecma262/#sec-property-descriptor-specification-type){#crossoriginpropertyfallback-(-p-):propertydescriptor
    x-internal="propertydescriptor"}{ \[\[Value\]\]: undefined,
    \[\[Writable\]\]: false, \[\[Enumerable\]\]: false,
    \[\[Configurable\]\]: true }.

2.  Throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#crossoriginpropertyfallback-(-p-):securityerror
    x-internal="securityerror"}
    [`DOMException`{#crossoriginpropertyfallback-(-p-):domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

###### [7.2.1.3.3]{.secno} [IsPlatformObjectSameOrigin]{.dfn} ( `O`{.variable} )[](#isplatformobjectsameorigin-(-o-)){.self-link} {#isplatformobjectsameorigin-(-o-)}

1.  Return true if the [current settings
    object](webappapis.html#current-settings-object){#isplatformobjectsameorigin-(-o-):current-settings-object}\'s
    [origin](webappapis.html#concept-settings-object-origin){#isplatformobjectsameorigin-(-o-):concept-settings-object-origin}
    is [same
    origin-domain](browsers.html#same-origin-domain){#isplatformobjectsameorigin-(-o-):same-origin-domain}
    with `O`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#isplatformobjectsameorigin-(-o-):relevant-settings-object}\'s
    [origin](webappapis.html#concept-settings-object-origin){#isplatformobjectsameorigin-(-o-):concept-settings-object-origin-2},
    and false otherwise.

This abstract operation does not return a [Completion
Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#isplatformobjectsameorigin-(-o-):completion-record
x-internal="completion-record"}.

Here the [current settings
object](webappapis.html#current-settings-object){#isplatformobjectsameorigin-(-o-):current-settings-object-2}
roughly corresponds to the \"caller\", because this check occurs before
the [execution
context](https://tc39.es/ecma262/#sec-execution-contexts){#isplatformobjectsameorigin-(-o-):javascript-execution-context
x-internal="javascript-execution-context"} for the getter/setter/method
in question makes its way onto the [JavaScript execution context
stack](https://tc39.es/ecma262/#execution-context-stack){#isplatformobjectsameorigin-(-o-):javascript-execution-context-stack
x-internal="javascript-execution-context-stack"}. For example, in the
code `w.document`, this step is invoked before the
[`document`{#isplatformobjectsameorigin-(-o-):dom-document-2}](#dom-document-2)
getter is reached as part of the [\[\[Get\]\]](#windowproxy-get)
algorithm for the
[`WindowProxy`{#isplatformobjectsameorigin-(-o-):windowproxy}](#windowproxy)
`w`{.variable}.

###### [7.2.1.3.4]{.secno} [CrossOriginGetOwnPropertyHelper]{.dfn} ( `O`{.variable}, `P`{.variable} )[](#crossorigingetownpropertyhelper-(-o,-p-)){.self-link} {#crossorigingetownpropertyhelper-(-o,-p-)}

If this abstract operation returns undefined and there is no custom
behavior, the caller needs to throw a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#crossorigingetownpropertyhelper-(-o,-p-):securityerror
x-internal="securityerror"}
[`DOMException`{#crossorigingetownpropertyhelper-(-o,-p-):domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
In practice this is handled by the caller calling
[CrossOriginPropertyFallback](#crossoriginpropertyfallback-(-p-)){#crossorigingetownpropertyhelper-(-o,-p-):crossoriginpropertyfallback-(-p-)}.

1.  Let `crossOriginKey`{.variable} be a tuple consisting of the
    [current settings
    object](webappapis.html#current-settings-object){#crossorigingetownpropertyhelper-(-o,-p-):current-settings-object},
    `O`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#crossorigingetownpropertyhelper-(-o,-p-):relevant-settings-object},
    and `P`{.variable}.

2.  For each `e`{.variable} of
    [CrossOriginProperties](#crossoriginproperties-(-o-)){#crossorigingetownpropertyhelper-(-o,-p-):crossoriginproperties-(-o-)}(`O`{.variable}):

    1.  If
        [SameValue](https://tc39.es/ecma262/#sec-samevalue){#crossorigingetownpropertyhelper-(-o,-p-):samevalue
        x-internal="samevalue"}(`e`{.variable}.\[\[Property\]\],
        `P`{.variable}) is true, then:

        1.  If the value of the
            [\[\[CrossOriginPropertyDescriptorMap\]\]](#crossoriginpropertydescriptormap){#crossorigingetownpropertyhelper-(-o,-p-):crossoriginpropertydescriptormap}
            internal slot of `O`{.variable} contains an entry whose key
            is `crossOriginKey`{.variable}, then return that entry\'s
            value.

        2.  Let `originalDesc`{.variable} be
            [OrdinaryGetOwnProperty](https://tc39.es/ecma262/#sec-ordinarygetownproperty){#crossorigingetownpropertyhelper-(-o,-p-):ordinarygetownproperty
            x-internal="ordinarygetownproperty"}(`O`{.variable},
            `P`{.variable}).

        3.  Let `crossOriginDesc`{.variable} be undefined.

        4.  If `e`{.variable}.\[\[NeedsGet\]\] and
            `e`{.variable}.\[\[NeedsSet\]\] are absent, then:

            1.  Let `value`{.variable} be
                `originalDesc`{.variable}.\[\[Value\]\].

            2.  If
                [IsCallable](https://tc39.es/ecma262/#sec-iscallable){#crossorigingetownpropertyhelper-(-o,-p-):iscallable
                x-internal="iscallable"}(`value`{.variable}) is true,
                then set `value`{.variable} to an anonymous built-in
                function, created in the [current
                realm](https://tc39.es/ecma262/#current-realm){#crossorigingetownpropertyhelper-(-o,-p-):current-realm
                x-internal="current-realm"}, that performs the same
                steps as the IDL operation `P`{.variable} on object
                `O`{.variable}.

            3.  Set `crossOriginDesc`{.variable} to
                [PropertyDescriptor](https://tc39.es/ecma262/#sec-property-descriptor-specification-type){#crossorigingetownpropertyhelper-(-o,-p-):propertydescriptor
                x-internal="propertydescriptor"}{ \[\[Value\]\]:
                `value`{.variable}, \[\[Enumerable\]\]: false,
                \[\[Writable\]\]: false, \[\[Configurable\]\]: true }.

        5.  Otherwise:

            1.  Let `crossOriginGet`{.variable} be undefined.

            2.  If `e`{.variable}.\[\[NeedsGet\]\] is true, then set
                `crossOriginGet`{.variable} to an anonymous built-in
                function, created in the [current
                realm](https://tc39.es/ecma262/#current-realm){#crossorigingetownpropertyhelper-(-o,-p-):current-realm-2
                x-internal="current-realm"}, that performs the same
                steps as the getter of the IDL attribute `P`{.variable}
                on object `O`{.variable}.

            3.  Let `crossOriginSet`{.variable} be undefined.

            4.  If `e`{.variable}.\[\[NeedsSet\]\] is true, then set
                `crossOriginSet`{.variable} to an anonymous built-in
                function, created in the [current
                realm](https://tc39.es/ecma262/#current-realm){#crossorigingetownpropertyhelper-(-o,-p-):current-realm-3
                x-internal="current-realm"}, that performs the same
                steps as the setter of the IDL attribute `P`{.variable}
                on object `O`{.variable}.

            5.  Set `crossOriginDesc`{.variable} to
                [PropertyDescriptor](https://tc39.es/ecma262/#sec-property-descriptor-specification-type){#crossorigingetownpropertyhelper-(-o,-p-):propertydescriptor-2
                x-internal="propertydescriptor"}{ \[\[Get\]\]:
                `crossOriginGet`{.variable}, \[\[Set\]\]:
                `crossOriginSet`{.variable}, \[\[Enumerable\]\]: false,
                \[\[Configurable\]\]: true }.

        6.  Create an entry in the value of the
            [\[\[CrossOriginPropertyDescriptorMap\]\]](#crossoriginpropertydescriptormap){#crossorigingetownpropertyhelper-(-o,-p-):crossoriginpropertydescriptormap-2}
            internal slot of `O`{.variable} with key
            `crossOriginKey`{.variable} and value
            `crossOriginDesc`{.variable}.

        7.  Return `crossOriginDesc`{.variable}.

3.  Return undefined.

This abstract operation does not return a [Completion
Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#crossorigingetownpropertyhelper-(-o,-p-):completion-record
x-internal="completion-record"}.

The reason that the property descriptors produced here are configurable
is to preserve the [invariants of the essential internal
methods](https://tc39.es/ecma262/#sec-invariants-of-the-essential-internal-methods){#crossorigingetownpropertyhelper-(-o,-p-):invariants-of-the-essential-internal-methods
x-internal="invariants-of-the-essential-internal-methods"} required by
the JavaScript specification. In particular, since the value of the
property can change as a consequence of navigation, it is required that
the property be configurable. (However, see [tc39/ecma262 issue
#672](https://github.com/tc39/ecma262/issues/672) and references to it
elsewhere in this specification for cases where we are not able to
preserve these invariants, for compatibility with existing web content.)
[\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

The reason the property descriptors are non-enumerable, despite this
mismatching the same-origin behavior, is for compatibility with existing
web content. See [issue
#3183](https://github.com/whatwg/html/issues/3183) for details.

###### [7.2.1.3.5]{.secno} [CrossOriginGet]{.dfn} ( `O`{.variable}, `P`{.variable}, `Receiver`{.variable} )[](#crossoriginget-(-o,-p,-receiver-)){.self-link} {#crossoriginget-(-o,-p,-receiver-)}

1.  Let `desc`{.variable} be ?
    `O`{.variable}.\[\[GetOwnProperty\]\](`P`{.variable}).

2.  [Assert](https://infra.spec.whatwg.org/#assert){#crossoriginget-(-o,-p,-receiver-):assert
    x-internal="assert"}: `desc`{.variable} is not undefined.

3.  If
    [IsDataDescriptor](https://tc39.es/ecma262/#sec-isdatadescriptor){#crossoriginget-(-o,-p,-receiver-):isdatadescriptor
    x-internal="isdatadescriptor"}(`desc`{.variable}) is true, then
    return `desc`{.variable}.\[\[Value\]\].

4.  [Assert](https://infra.spec.whatwg.org/#assert){#crossoriginget-(-o,-p,-receiver-):assert-2
    x-internal="assert"}:
    [IsAccessorDescriptor](https://tc39.es/ecma262/#sec-isaccessordescriptor){#crossoriginget-(-o,-p,-receiver-):isaccessordescriptor
    x-internal="isaccessordescriptor"}(`desc`{.variable}) is true.

5.  Let `getter`{.variable} be `desc`{.variable}.\[\[Get\]\].

6.  If `getter`{.variable} is undefined, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#crossoriginget-(-o,-p,-receiver-):securityerror
    x-internal="securityerror"}
    [`DOMException`{#crossoriginget-(-o,-p,-receiver-):domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

7.  Return ?
    [Call](https://tc39.es/ecma262/#sec-call){#crossoriginget-(-o,-p,-receiver-):call
    x-internal="call"}(`getter`{.variable}, `Receiver`{.variable}).

###### [7.2.1.3.6]{.secno} [CrossOriginSet]{.dfn} ( `O`{.variable}, `P`{.variable}, `V`{.variable}, `Receiver`{.variable} )[](#crossoriginset-(-o,-p,-v,-receiver-)){.self-link} {#crossoriginset-(-o,-p,-v,-receiver-)}

1.  Let `desc`{.variable} be ?
    `O`{.variable}.\[\[GetOwnProperty\]\](`P`{.variable}).

2.  [Assert](https://infra.spec.whatwg.org/#assert){#crossoriginset-(-o,-p,-v,-receiver-):assert
    x-internal="assert"}: `desc`{.variable} is not undefined.

3.  If `desc`{.variable}.\[\[Set\]\] is present and its value is not
    undefined, then:

    1.  Perform ?
        [Call](https://tc39.es/ecma262/#sec-call){#crossoriginset-(-o,-p,-v,-receiver-):call
        x-internal="call"}(`desc`{.variable}.\[\[Set\]\],
        `Receiver`{.variable}, « `V`{.variable} »).

    2.  Return true.

4.  Throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#crossoriginset-(-o,-p,-v,-receiver-):securityerror
    x-internal="securityerror"}
    [`DOMException`{#crossoriginset-(-o,-p,-v,-receiver-):domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

###### [7.2.1.3.7]{.secno} [CrossOriginOwnPropertyKeys]{.dfn} ( `O`{.variable} )[](#crossoriginownpropertykeys-(-o-)){.self-link} {#crossoriginownpropertykeys-(-o-)}

1.  Let `keys`{.variable} be a new empty
    [List](https://tc39.es/ecma262/#sec-list-and-record-specification-type){#crossoriginownpropertykeys-(-o-):js-list
    x-internal="js-list"}.

2.  For each `e`{.variable} of
    [CrossOriginProperties](#crossoriginproperties-(-o-)){#crossoriginownpropertykeys-(-o-):crossoriginproperties-(-o-)}(`O`{.variable}),
    [append](https://infra.spec.whatwg.org/#list-append){#crossoriginownpropertykeys-(-o-):list-append
    x-internal="list-append"} `e`{.variable}.\[\[Property\]\] to
    `keys`{.variable}.

3.  Return the concatenation of `keys`{.variable} and « \"`then`\",
    [%Symbol.toStringTag%](infrastructure.html#symbol.tostringtag){#crossoriginownpropertykeys-(-o-):symbol.tostringtag},
    [%Symbol.hasInstance%](infrastructure.html#symbol.hasinstance){#crossoriginownpropertykeys-(-o-):symbol.hasinstance},
    [%Symbol.isConcatSpreadable%](infrastructure.html#symbol.isconcatspreadable){#crossoriginownpropertykeys-(-o-):symbol.isconcatspreadable}
    ».

This abstract operation does not return a [Completion
Record](https://tc39.es/ecma262/#sec-completion-record-specification-type){#crossoriginownpropertykeys-(-o-):completion-record
x-internal="completion-record"}.

#### [7.2.2]{.secno} The [`Window`{#the-window-object:window}](#window) object[](#the-window-object){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Window](https://developer.mozilla.org/en-US/docs/Web/API/Window "The Window interface represents a window containing a DOM document; the document property points to the DOM document loaded in that window.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Global=Window,
 Exposed=Window,
 LegacyUnenumerableNamedProperties]
interface Window : EventTarget {
  // the current browsing context
  [LegacyUnforgeable] readonly attribute WindowProxy window;
  [Replaceable] readonly attribute WindowProxy self;
  [LegacyUnforgeable] readonly attribute Document document;
  attribute DOMString name; 
  [PutForwards=href, LegacyUnforgeable] readonly attribute Location location;
  readonly attribute History history;
  readonly attribute Navigation navigation;
  readonly attribute CustomElementRegistry customElements;
  [Replaceable] readonly attribute BarProp locationbar;
  [Replaceable] readonly attribute BarProp menubar;
  [Replaceable] readonly attribute BarProp personalbar;
  [Replaceable] readonly attribute BarProp scrollbars;
  [Replaceable] readonly attribute BarProp statusbar;
  [Replaceable] readonly attribute BarProp toolbar;
  attribute DOMString status;
  undefined close();
  readonly attribute boolean closed;
  undefined stop();
  undefined focus();
  undefined blur();

  // other browsing contexts
  [Replaceable] readonly attribute WindowProxy frames;
  [Replaceable] readonly attribute unsigned long length;
  [LegacyUnforgeable] readonly attribute WindowProxy? top;
  attribute any opener;
  [Replaceable] readonly attribute WindowProxy? parent;
  readonly attribute Element? frameElement;
  WindowProxy? open(optional USVString url = "", optional DOMString target = "_blank", optional [LegacyNullToEmptyString] DOMString features = "");

  // Since this is the global object, the IDL named getter adds a NamedPropertiesObject exotic
  // object on the prototype chain. Indeed, this does not make the global object an exotic object.
  // Indexed access is taken care of by the WindowProxy exotic object.
  getter object (DOMString name);

  // the user agent
  readonly attribute Navigator navigator;
  [Replaceable] readonly attribute Navigator clientInformation; // legacy alias of .navigator
  readonly attribute boolean originAgentCluster;

  // user prompts
  undefined alert();
  undefined alert(DOMString message);
  boolean confirm(optional DOMString message = "");
  DOMString? prompt(optional DOMString message = "", optional DOMString default = "");
  undefined print();

  undefined postMessage(any message, USVString targetOrigin, optional sequence<object> transfer = []);
  undefined postMessage(any message, optional WindowPostMessageOptions options = {});

  // also has obsolete members
};
Window includes GlobalEventHandlers;
Window includes WindowEventHandlers;

dictionary WindowPostMessageOptions : StructuredSerializeOptions {
  USVString targetOrigin = "/";
};
```

`window`{.variable}`.`[`window`](#dom-window){#dom-window-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/window](https://developer.mozilla.org/en-US/docs/Web/API/Window/window "The window property of a Window object points to the window object itself.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`window`{.variable}`.`[`frames`](#dom-frames){#dom-frames-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/frames](https://developer.mozilla.org/en-US/docs/Web/API/Window/frames "Returns the window itself, which is an array-like object, listing the direct sub-frames of the current window.")

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

`window`{.variable}`.`[`self`](#dom-self){#dom-self-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/self](https://developer.mozilla.org/en-US/docs/Web/API/Window/self "The Window.self read-only property returns the window itself, as a WindowProxy. It can be used with dot notation on a window object (that is, window.self) or standalone (self). The advantage of the standalone notation is that a similar notation exists for non-window contexts, such as in Web Workers. By using self, you can refer to the global scope in a way that will work not only in a window context (self will resolve to window.self) but also in a worker context (self will then resolve to WorkerGlobalScope.self).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

These attributes all return `window`{.variable}.

`window`{.variable}`.`[`document`](#dom-document-2){#dom-document-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/document](https://developer.mozilla.org/en-US/docs/Web/API/Window/document "window.document returns a reference to the document contained in the window.")

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

Returns the
[`Document`{#the-window-object:document-2}](dom.html#document)
associated with `window`{.variable}.

`document`{.variable}`.`[`defaultView`](#dom-document-defaultview){#dom-document-defaultview-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Document/defaultView](https://developer.mozilla.org/en-US/docs/Web/API/Document/defaultView "In browsers, document.defaultView returns the window object associated with a document, or null if none is available.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Window`{#the-window-object:window-4}](#window) associated
with `document`{.variable}, if there is one, or null otherwise.

The [`Window`{#the-window-object:window-5}](#window) object has an
[associated `Document`]{#concept-document-window .dfn export=""}, which
is a [`Document`{#the-window-object:document-3}](dom.html#document)
object. It is set when the
[`Window`{#the-window-object:window-6}](#window) object is created, and
only ever changed during
[navigation](browsing-the-web.html#navigate){#the-window-object:navigate}
from the [initial
`about:blank`](dom.html#is-initial-about:blank){#the-window-object:is-initial-about:blank}
[`Document`{#the-window-object:document-4}](dom.html#document).

A [`Window`{#the-window-object:window-7}](#window)\'s [browsing
context]{#window-bc .dfn dfn-for="Window" export=""} is its [associated
`Document`](#concept-document-window){#the-window-object:concept-document-window}\'s
[browsing
context](document-sequences.html#concept-document-bc){#the-window-object:concept-document-bc}.
[It is either null or a [browsing
context](document-sequences.html#browsing-context){#the-window-object:browsing-context}.]{.note}

A [`Window`{#the-window-object:window-8}](#window)\'s
[navigable]{#window-navigable .dfn dfn-for="Window" export=""} is the
[navigable](document-sequences.html#navigable){#the-window-object:navigable}
whose [active
document](document-sequences.html#nav-document){#the-window-object:nav-document}
is the [`Window`{#the-window-object:window-9}](#window)\'s [associated
`Document`](#concept-document-window){#the-window-object:concept-document-window-2}\'s,
or null if there is no such
[navigable](document-sequences.html#navigable){#the-window-object:navigable-2}.

The [`window`]{#dom-window .dfn dfn-for="Window" dfn-type="attribute"},
[`frames`]{#dom-frames .dfn dfn-for="Window" dfn-type="attribute"}, and
[`self`]{#dom-self .dfn dfn-for="Window" dfn-type="attribute"} getter
steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-window-object:this
x-internal="this"}\'s [relevant
realm](webappapis.html#concept-relevant-realm){#the-window-object:concept-relevant-realm}.\[\[GlobalEnv\]\].\[\[GlobalThisValue\]\].

The [`document`]{#dom-document-2 .dfn dfn-for="Window"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-window-object:this-2
x-internal="this"}\'s [associated
`Document`](#concept-document-window){#the-window-object:concept-document-window-3}.

The [`Document`{#the-window-object:document-5}](dom.html#document)
object associated with a
[`Window`{#the-window-object:window-10}](#window) object can change in
exactly one case: when the
[navigate](browsing-the-web.html#navigate){#the-window-object:navigate-2}
algorithm [creates a new `Document`
object](document-lifecycle.html#initialise-the-document-object){#the-window-object:initialise-the-document-object}
for the first page loaded in a [browsing
context](document-sequences.html#browsing-context){#the-window-object:browsing-context-2}.
In that specific case, the
[`Window`{#the-window-object:window-11}](#window) object of the [initial
`about:blank`](dom.html#is-initial-about:blank){#the-window-object:is-initial-about:blank-2}
page is reused and gets a new
[`Document`{#the-window-object:document-6}](dom.html#document) object.

The [`defaultView`]{#dom-document-defaultview .dfn dfn-for="Document"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-window-object:this-3
    x-internal="this"}\'s [browsing
    context](document-sequences.html#concept-document-bc){#the-window-object:concept-document-bc-2}
    is null, then return null.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-window-object:this-4
    x-internal="this"}\'s [browsing
    context](document-sequences.html#concept-document-bc){#the-window-object:concept-document-bc-3}\'s
    [`WindowProxy`{#the-window-object:windowproxy-8}](#windowproxy)
    object.

------------------------------------------------------------------------

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HTMLDocument](https://developer.mozilla.org/en-US/docs/Web/API/HTMLDocument "For historical reasons, Window objects have a window.HTMLDocument property whose value is the Document interface. So you can think of HTMLDocument as an alias for Document, and you can find documentation for HTMLDocument members under the documentation for the Document interface.")

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

For historical reasons,
[`Window`{#the-window-object:window-12}](#window) objects must also have
a writable, configurable, non-enumerable property named
[`HTMLDocument`]{#htmldocument .dfn} whose value is the
[`Document`{#the-window-object:document-7}](dom.html#document)
[interface
object](https://webidl.spec.whatwg.org/#dfn-interface-object){#the-window-object:interface-object
x-internal="interface-object"}.

##### [7.2.2.1]{.secno} Opening and closing windows[](#apis-for-creating-and-navigating-browsing-contexts-by-name){.self-link} {#apis-for-creating-and-navigating-browsing-contexts-by-name}

`window`{.variable}` = ``window`{.variable}`.`[`open`](#dom-open){#dom-open-dev}`([ ``url`{.variable}` [, ``target`{.variable}` [, ``features`{.variable}` ] ] ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/open](https://developer.mozilla.org/en-US/docs/Web/API/Window/open "The open() method of the Window interface loads a specified resource into a new or existing browsing context (that is, a tab, a window, or an iframe) under a specified name.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Opens a window to show `url`{.variable} (defaults to
\"[`about:blank`{#apis-for-creating-and-navigating-browsing-contexts-by-name:about:blank}](infrastructure.html#about:blank)\"),
and returns it. `target`{.variable} (defaults to \"`_blank`\") gives the
name of the new window. If a window already exists with that name, it is
reused. The `features`{.variable} argument can contain a [set of
comma-separated
tokens](common-microsyntaxes.html#set-of-comma-separated-tokens){#apis-for-creating-and-navigating-browsing-contexts-by-name:set-of-comma-separated-tokens}:

\"`noopener`\"\
\"`noreferrer`\"
:   These behave equivalently to the
    [`noopener`{#apis-for-creating-and-navigating-browsing-contexts-by-name:link-type-noopener}](links.html#link-type-noopener)
    and
    [`noreferrer`{#apis-for-creating-and-navigating-browsing-contexts-by-name:link-type-noreferrer}](links.html#link-type-noreferrer)
    link types on
    [hyperlinks](links.html#hyperlink){#apis-for-creating-and-navigating-browsing-contexts-by-name:hyperlink}.

\"`popup`\"

:   Encourages user agents to provide a minimal web browser user
    interface for the new window. (Impacts the
    [`visible`{#apis-for-creating-and-navigating-browsing-contexts-by-name:dom-barprop-visible}](#dom-barprop-visible)
    getter on all
    [`BarProp`{#apis-for-creating-and-navigating-browsing-contexts-by-name:barprop}](#barprop)
    objects as well.)

``` example
globalThis.open("https://email.example/message/CAOOOkFcWW97r8yg=SsWg7GgCmp4suVX9o85y8BvNRqMjuc5PXg", undefined, "noopener,popup");
```

`window`{.variable}`.`[`name`](#dom-name){#dom-name-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/name](https://developer.mozilla.org/en-US/docs/Web/API/Window/name "The Window.name property gets/sets the name of the window's browsing context.")

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

Returns the name of the window.

Can be set, to change the name.

`window`{.variable}`.`[`close`](#dom-window-close){#dom-window-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/close](https://developer.mozilla.org/en-US/docs/Web/API/Window/close "The Window.close() method closes the current window, or the window on which it was called.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Closes the window.

`window`{.variable}`.`[`closed`](#dom-window-closed){#dom-window-closed-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/closed](https://developer.mozilla.org/en-US/docs/Web/API/Window/closed "The Window.closed read-only property indicates whether the referenced window is closed or not.")

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

Returns true if the window has been closed, false otherwise.

`window`{.variable}`.`[`stop`](#dom-window-stop){#dom-window-stop-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/stop](https://developer.mozilla.org/en-US/docs/Web/API/Window/stop "The window.stop() stops further resource loading in the current browsing context, equivalent to the stop button in the browser.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Cancels the document load.

To [get noopener for window open]{#get-noopener-for-window-open .dfn},
given a
[`Document`{#apis-for-creating-and-navigating-browsing-contexts-by-name:document}](dom.html#document)
`sourceDocument`{.variable}, an [ordered
map](https://infra.spec.whatwg.org/#ordered-map){#apis-for-creating-and-navigating-browsing-contexts-by-name:ordered-map
x-internal="ordered-map"} `tokenizedFeatures`{.variable}, and a [URL
record](https://url.spec.whatwg.org/#concept-url){#apis-for-creating-and-navigating-browsing-contexts-by-name:url-record
x-internal="url-record"}-or-null `url`{.variable}, perform the following
steps. They return a boolean.

1.  If `url`{.variable} is not null and `url`{.variable}\'s [blob URL
    entry](https://url.spec.whatwg.org/#concept-url-blob-entry){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-url-blob-entry
    x-internal="concept-url-blob-entry"} is not null:

    1.  Let `blobOrigin`{.variable} be `url`{.variable}\'s [blob URL
        entry](https://url.spec.whatwg.org/#concept-url-blob-entry){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-url-blob-entry-2
        x-internal="concept-url-blob-entry"}\'s
        [environment](https://w3c.github.io/FileAPI/#blob-url-entry-environment){#apis-for-creating-and-navigating-browsing-contexts-by-name:blob-url-entry-environment
        x-internal="blob-url-entry-environment"}\'s
        [origin](webappapis.html#concept-settings-object-origin){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-settings-object-origin}.

    2.  Let `topLevelOrigin`{.variable} be
        `sourceDocument`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#apis-for-creating-and-navigating-browsing-contexts-by-name:relevant-settings-object}\'s
        [top-level
        origin](webappapis.html#concept-environment-top-level-origin){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-environment-top-level-origin}.

    3.  If `blobOrigin`{.variable} is not [same
        site](browsers.html#same-site){#apis-for-creating-and-navigating-browsing-contexts-by-name:same-site}
        with `topLevelOrigin`{.variable}, then return true.

2.  Let `noopener`{.variable} be false.

3.  If `tokenizedFeatures`{.variable}\[\"`noopener`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#apis-for-creating-and-navigating-browsing-contexts-by-name:map-exists
    x-internal="map-exists"}, then set `noopener`{.variable} to the
    result of [parsing `tokenizedFeatures`{.variable}\[\"`noopener`\"\]
    as a boolean
    feature](#concept-window-open-features-parse-boolean){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-window-open-features-parse-boolean}.

4.  Return `noopener`{.variable}.

The [window open steps]{#window-open-steps .dfn}, given a string
`url`{.variable}, a string `target`{.variable}, and a string
`features`{.variable}, are as follows:

1.  If the [event
    loop](webappapis.html#event-loop){#apis-for-creating-and-navigating-browsing-contexts-by-name:event-loop}\'s
    [termination nesting
    level](document-lifecycle.html#termination-nesting-level){#apis-for-creating-and-navigating-browsing-contexts-by-name:termination-nesting-level}
    is nonzero, then return null.

2.  Let `sourceDocument`{.variable} be the [entry global
    object](webappapis.html#entry-global-object){#apis-for-creating-and-navigating-browsing-contexts-by-name:entry-global-object}\'s
    [associated
    `Document`](#concept-document-window){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-document-window}.

3.  Let `urlRecord`{.variable} be null.

4.  If `url`{.variable} is not the empty string, then:

    1.  Set `urlRecord`{.variable} to the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#apis-for-creating-and-navigating-browsing-contexts-by-name:encoding-parsing-a-url}
        given `url`{.variable}, relative to `sourceDocument`{.variable}.

    2.  If `urlRecord`{.variable} is failure, then throw a
        [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#apis-for-creating-and-navigating-browsing-contexts-by-name:syntaxerror
        x-internal="syntaxerror"}
        [`DOMException`{#apis-for-creating-and-navigating-browsing-contexts-by-name:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  If `target`{.variable} is the empty string, then set
    `target`{.variable} to \"`_blank`\".

6.  Let `tokenizedFeatures`{.variable} be the result of
    [tokenizing](#concept-window-open-features-tokenize){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-window-open-features-tokenize}
    `features`{.variable}.

7.  Let `noreferrer`{.variable} be false.

8.  If `tokenizedFeatures`{.variable}\[\"`noreferrer`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#apis-for-creating-and-navigating-browsing-contexts-by-name:map-exists-2
    x-internal="map-exists"}, then set `noreferrer`{.variable} to the
    result of [parsing
    `tokenizedFeatures`{.variable}\[\"`noreferrer`\"\] as a boolean
    feature](#concept-window-open-features-parse-boolean){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-window-open-features-parse-boolean-2}.

9.  Let `noopener`{.variable} be the result of [getting noopener for
    window
    open](#get-noopener-for-window-open){#apis-for-creating-and-navigating-browsing-contexts-by-name:get-noopener-for-window-open}
    with `sourceDocument`{.variable}, `tokenizedFeatures`{.variable},
    and `urlRecord`{.variable}.

10. [Remove](https://infra.spec.whatwg.org/#map-remove){#apis-for-creating-and-navigating-browsing-contexts-by-name:map-remove
    x-internal="map-remove"}
    `tokenizedFeatures`{.variable}\[\"`noopener`\"\] and
    `tokenizedFeatures`{.variable}\[\"`noreferrer`\"\].

11. Let `referrerPolicy`{.variable} be the empty string.

12. If `noreferrer`{.variable} is true, then set `noopener`{.variable}
    to true and set `referrerPolicy`{.variable} to \"`no-referrer`\".

13. Let `targetNavigable`{.variable} and `windowType`{.variable} be the
    result of applying [the rules for choosing a
    navigable](document-sequences.html#the-rules-for-choosing-a-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:the-rules-for-choosing-a-navigable}
    given `target`{.variable}, `sourceDocument`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:node-navigable},
    and `noopener`{.variable}.

    If there is a user agent that supports control-clicking a link to
    open it in a new tab, and the user control-clicks on an element
    whose
    [`onclick`{#apis-for-creating-and-navigating-browsing-contexts-by-name:handler-onclick}](webappapis.html#handler-onclick)
    handler uses the
    [`window.open()`{#apis-for-creating-and-navigating-browsing-contexts-by-name:dom-open}](#dom-open)
    API to open a page in an
    [`iframe`{#apis-for-creating-and-navigating-browsing-contexts-by-name:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
    element, the user agent could override the selection of the target
    browsing context to instead target a new tab.

14. If `targetNavigable`{.variable} is null, then return null.

15. If `windowType`{.variable} is either \"`new and unrestricted`\" or
    \"`new with no opener`\", then:

    1.  Set `targetNavigable`{.variable}\'s [active browsing
        context](document-sequences.html#nav-bc){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-bc}\'s
        [is
        popup](document-sequences.html#is-popup){#apis-for-creating-and-navigating-browsing-contexts-by-name:is-popup}
        to the result of [checking if a popup window is
        requested](#popup-window-is-requested){#apis-for-creating-and-navigating-browsing-contexts-by-name:popup-window-is-requested},
        given `tokenizedFeatures`{.variable}.

    2.  [Set up browsing context
        features](https://drafts.csswg.org/cssom-view/#set-up-browsing-context-features){#apis-for-creating-and-navigating-browsing-contexts-by-name:set-up-browsing-context-features
        x-internal="set-up-browsing-context-features"} for
        `targetNavigable`{.variable}\'s [active browsing
        context](document-sequences.html#nav-bc){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-bc-2}
        given `tokenizedFeatures`{.variable}.
        [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    3.  If `urlRecord`{.variable} is null, then set
        `urlRecord`{.variable} to a [URL
        record](https://url.spec.whatwg.org/#concept-url){#apis-for-creating-and-navigating-browsing-contexts-by-name:url-record-2
        x-internal="url-record"} representing
        [`about:blank`{#apis-for-creating-and-navigating-browsing-contexts-by-name:about:blank-2}](infrastructure.html#about:blank).

    4.  If `urlRecord`{.variable} [matches
        `about:blank`](urls-and-fetching.html#matches-about:blank){#apis-for-creating-and-navigating-browsing-contexts-by-name:matches-about:blank},
        then perform the [URL and history update
        steps](browsing-the-web.html#url-and-history-update-steps){#apis-for-creating-and-navigating-browsing-contexts-by-name:url-and-history-update-steps}
        given `targetNavigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-document}
        and `urlRecord`{.variable}.

        This is necessary in case `url`{.variable} is something like
        `about:blank?foo`. If `url`{.variable} is just plain
        `about:blank`, this will do nothing.

    5.  Otherwise,
        [navigate](browsing-the-web.html#navigate){#apis-for-creating-and-navigating-browsing-contexts-by-name:navigate}
        `targetNavigable`{.variable} to `urlRecord`{.variable} using
        `sourceDocument`{.variable}, with
        *[referrerPolicy](browsing-the-web.html#navigation-referrer-policy)*
        set to `referrerPolicy`{.variable} and
        *[exceptionsEnabled](browsing-the-web.html#exceptions-enabled){#apis-for-creating-and-navigating-browsing-contexts-by-name:exceptions-enabled}*
        set to true.

16. Otherwise:

    1.  If `urlRecord`{.variable} is not null, then
        [navigate](browsing-the-web.html#navigate){#apis-for-creating-and-navigating-browsing-contexts-by-name:navigate-2}
        `targetNavigable`{.variable} to `urlRecord`{.variable} using
        `sourceDocument`{.variable}, with
        *[referrerPolicy](browsing-the-web.html#navigation-referrer-policy)*
        set to `referrerPolicy`{.variable} and
        *[exceptionsEnabled](browsing-the-web.html#exceptions-enabled){#apis-for-creating-and-navigating-browsing-contexts-by-name:exceptions-enabled-2}*
        set to true.

    2.  If `noopener`{.variable} is false, then set
        `targetNavigable`{.variable}\'s [active browsing
        context](document-sequences.html#nav-bc){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-bc-3}\'s
        [opener browsing
        context](document-sequences.html#opener-browsing-context){#apis-for-creating-and-navigating-browsing-contexts-by-name:opener-browsing-context}
        to `sourceDocument`{.variable}\'s [browsing
        context](document-sequences.html#concept-document-bc){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-document-bc}.

17. If `noopener`{.variable} is true or `windowType`{.variable} is
    \"`new with no opener`\", then return null.

18. Return `targetNavigable`{.variable}\'s [active
    `WindowProxy`](document-sequences.html#nav-wp){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-wp}.

The
[`open(``url`{.variable}`, ``target`{.variable}`, ``features`{.variable}`)`]{#dom-open
.dfn dfn-for="Window" dfn-type="method"} method steps are to run the
[window open
steps](#window-open-steps){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-open-steps}
with `url`{.variable}, `target`{.variable}, and `features`{.variable}.

The method provides a mechanism for
[navigating](browsing-the-web.html#navigate){#apis-for-creating-and-navigating-browsing-contexts-by-name:navigate-3}
an existing [browsing
context](document-sequences.html#browsing-context){#apis-for-creating-and-navigating-browsing-contexts-by-name:browsing-context}
or opening and navigating an [auxiliary browsing
context](document-sequences.html#auxiliary-browsing-context){#apis-for-creating-and-navigating-browsing-contexts-by-name:auxiliary-browsing-context}.

------------------------------------------------------------------------

To [tokenize the `features`{.variable}
argument]{#concept-window-open-features-tokenize .dfn}:

1.  Let `tokenizedFeatures`{.variable} be a new [ordered
    map](https://infra.spec.whatwg.org/#ordered-map){#apis-for-creating-and-navigating-browsing-contexts-by-name:ordered-map-2
    x-internal="ordered-map"}.

2.  Let `position`{.variable} point at the first code point of
    `features`{.variable}.

3.  [While](https://infra.spec.whatwg.org/#iteration-while){#apis-for-creating-and-navigating-browsing-contexts-by-name:while
    x-internal="while"} `position`{.variable} is not past the end of
    `features`{.variable}:

    1.  Let `name`{.variable} be the empty string.

    2.  Let `value`{.variable} be the empty string.

    3.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#apis-for-creating-and-navigating-browsing-contexts-by-name:collect-a-sequence-of-code-points
        x-internal="collect-a-sequence-of-code-points"} that are
        [feature
        separators](#feature-separator){#apis-for-creating-and-navigating-browsing-contexts-by-name:feature-separator}
        from `features`{.variable} given `position`{.variable}. This
        skips past leading separators before the name.

    4.  [Collect a sequence of code
        points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#apis-for-creating-and-navigating-browsing-contexts-by-name:collect-a-sequence-of-code-points-2
        x-internal="collect-a-sequence-of-code-points"} that are not
        [feature
        separators](#feature-separator){#apis-for-creating-and-navigating-browsing-contexts-by-name:feature-separator-2}
        from `features`{.variable} given `position`{.variable}. Set
        `name`{.variable} to the collected characters, [converted to
        ASCII
        lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#apis-for-creating-and-navigating-browsing-contexts-by-name:converted-to-ascii-lowercase
        x-internal="converted-to-ascii-lowercase"}.

    5.  Set `name`{.variable} to the result of [normalizing the feature
        name](#normalizing-the-feature-name){#apis-for-creating-and-navigating-browsing-contexts-by-name:normalizing-the-feature-name}
        `name`{.variable}.

    6.  [While](https://infra.spec.whatwg.org/#iteration-while){#apis-for-creating-and-navigating-browsing-contexts-by-name:while-2
        x-internal="while"} `position`{.variable} is not past the end of
        `features`{.variable} and the code point at
        `position`{.variable} in `features`{.variable} is not U+003D
        (=):

        1.  If the code point at `position`{.variable} in
            `features`{.variable} is U+002C (,), or if it is not a
            [feature
            separator](#feature-separator){#apis-for-creating-and-navigating-browsing-contexts-by-name:feature-separator-3},
            then
            [break](https://infra.spec.whatwg.org/#iteration-break){#apis-for-creating-and-navigating-browsing-contexts-by-name:break
            x-internal="break"}.

        2.  Advance `position`{.variable} by 1.

        This skips to the first U+003D (=) but does not skip past a
        U+002C (,) or a non-separator.

    7.  If the code point at `position`{.variable} in
        `features`{.variable} is a [feature
        separator](#feature-separator){#apis-for-creating-and-navigating-browsing-contexts-by-name:feature-separator-4}:

        1.  While `position`{.variable} is not past the end of
            `features`{.variable} and the code point at
            `position`{.variable} in `features`{.variable} is a [feature
            separator](#feature-separator){#apis-for-creating-and-navigating-browsing-contexts-by-name:feature-separator-5}:

            1.  If the code point at `position`{.variable} in
                `features`{.variable} is U+002C (,), then
                [break](https://infra.spec.whatwg.org/#iteration-break){#apis-for-creating-and-navigating-browsing-contexts-by-name:break-2
                x-internal="break"}.

            2.  Advance `position`{.variable} by 1.

            This skips to the first non-separator but does not skip past
            a U+002C (,).

        2.  [Collect a sequence of code
            points](https://infra.spec.whatwg.org/#collect-a-sequence-of-code-points){#apis-for-creating-and-navigating-browsing-contexts-by-name:collect-a-sequence-of-code-points-3
            x-internal="collect-a-sequence-of-code-points"} that are not
            [feature
            separators](#feature-separator){#apis-for-creating-and-navigating-browsing-contexts-by-name:feature-separator-6}
            code points from `features`{.variable} given
            `position`{.variable}. Set `value`{.variable} to the
            collected code points, [converted to ASCII
            lowercase](https://infra.spec.whatwg.org/#ascii-lowercase){#apis-for-creating-and-navigating-browsing-contexts-by-name:converted-to-ascii-lowercase-2
            x-internal="converted-to-ascii-lowercase"}.

    8.  If `name`{.variable} is not the empty string, then set
        `tokenizedFeatures`{.variable}\[`name`{.variable}\] to
        `value`{.variable}.

4.  Return `tokenizedFeatures`{.variable}.

To [check if a window feature is set]{#window-feature-is-set .dfn},
given `tokenizedFeatures`{.variable}, `featureName`{.variable}, and
`defaultValue`{.variable}:

1.  If `tokenizedFeatures`{.variable}\[`featureName`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#apis-for-creating-and-navigating-browsing-contexts-by-name:map-exists-3
    x-internal="map-exists"}, then return the result of [parsing
    `tokenizedFeatures`{.variable}\[`featureName`{.variable}\] as a
    boolean
    feature](#concept-window-open-features-parse-boolean){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-window-open-features-parse-boolean-3}.

2.  Return `defaultValue`{.variable}.

To [check if a popup window is requested]{#popup-window-is-requested
.dfn}, given `tokenizedFeatures`{.variable}:

1.  If `tokenizedFeatures`{.variable} is
    [empty](https://infra.spec.whatwg.org/#map-is-empty){#apis-for-creating-and-navigating-browsing-contexts-by-name:map-empty
    x-internal="map-empty"}, then return false.

2.  If `tokenizedFeatures`{.variable}\[\"`popup`\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#apis-for-creating-and-navigating-browsing-contexts-by-name:map-exists-4
    x-internal="map-exists"}, then return the result of [parsing
    `tokenizedFeatures`{.variable}\[\"`popup`\"\] as a boolean
    feature](#concept-window-open-features-parse-boolean){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-window-open-features-parse-boolean-4}.

3.  Let `location`{.variable} be the result of [checking if a window
    feature is
    set](#window-feature-is-set){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-feature-is-set},
    given `tokenizedFeatures`{.variable}, \"`location`\", and false.

4.  Let `toolbar`{.variable} be the result of [checking if a window
    feature is
    set](#window-feature-is-set){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-feature-is-set-2},
    given `tokenizedFeatures`{.variable}, \"`toolbar`\", and false.

5.  If `location`{.variable} and `toolbar`{.variable} are both false,
    then return true.

6.  Let `menubar`{.variable} be the result of [checking if a window
    feature is
    set](#window-feature-is-set){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-feature-is-set-3},
    given `tokenizedFeatures`{.variable}, \"`menubar`\", and false.

7.  If `menubar`{.variable} is false, then return true.

8.  Let `resizable`{.variable} be the result of [checking if a window
    feature is
    set](#window-feature-is-set){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-feature-is-set-4},
    given `tokenizedFeatures`{.variable}, \"`resizable`\", and true.

9.  If `resizable`{.variable} is false, then return true.

10. Let `scrollbars`{.variable} be the result of [checking if a window
    feature is
    set](#window-feature-is-set){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-feature-is-set-5},
    given `tokenizedFeatures`{.variable}, \"`scrollbars`\", and false.

11. If `scrollbars`{.variable} is false, then return true.

12. Let `status`{.variable} be the result of [checking if a window
    feature is
    set](#window-feature-is-set){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-feature-is-set-6},
    given `tokenizedFeatures`{.variable}, \"`status`\", and false.

13. If `status`{.variable} is false, then return true.

14. Return false.

A code point is a [feature separator]{#feature-separator .dfn} if it is
[ASCII
whitespace](https://infra.spec.whatwg.org/#ascii-whitespace){#apis-for-creating-and-navigating-browsing-contexts-by-name:space-characters
x-internal="space-characters"}, U+003D (=), or U+002C (,).

For legacy reasons, there are some aliases of some feature names. To
[normalize a feature name]{#normalizing-the-feature-name .dfn}
`name`{.variable}, switch on `name`{.variable}:

\"`screenx`\"
:   Return \"`left`\".

\"`screeny`\"
:   Return \"`top`\".

\"`innerwidth`\"
:   Return \"`width`\".

\"`innerheight`\"
:   Return \"`height`\".

Anything else
:   Return `name`{.variable}.

To [parse a boolean feature]{#concept-window-open-features-parse-boolean
.dfn} given a string `value`{.variable}:

1.  If `value`{.variable} is the empty string, then return true.

2.  If `value`{.variable}
    [is](https://infra.spec.whatwg.org/#string-is){#apis-for-creating-and-navigating-browsing-contexts-by-name:is
    x-internal="is"} \"`yes`\", then return true.

3.  If `value`{.variable}
    [is](https://infra.spec.whatwg.org/#string-is){#apis-for-creating-and-navigating-browsing-contexts-by-name:is-2
    x-internal="is"} \"`true`\", then return true.

4.  Let `parsed`{.variable} be the result of [parsing `value`{.variable}
    as an
    integer](common-microsyntaxes.html#rules-for-parsing-integers){#apis-for-creating-and-navigating-browsing-contexts-by-name:rules-for-parsing-integers}.

5.  If `parsed`{.variable} is an error, then set it to 0.

6.  Return false if `parsed`{.variable} is 0, and true otherwise.

------------------------------------------------------------------------

The [`name`]{#dom-name .dfn dfn-for="Window" dfn-type="attribute"}
getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this
    x-internal="this"}\'s
    [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable}
    is null, then return the empty string.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this-2
    x-internal="this"}\'s
    [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable-2}\'s
    [target
    name](document-sequences.html#nav-target){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-target}.

The
[`name`{#apis-for-creating-and-navigating-browsing-contexts-by-name:dom-name}](#dom-name)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this-3
    x-internal="this"}\'s
    [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable-3}
    is null, then return.

2.  Set
    [this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this-4
    x-internal="this"}\'s
    [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable-4}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-active-history-entry}\'s
    [document
    state](browsing-the-web.html#she-document-state){#apis-for-creating-and-navigating-browsing-contexts-by-name:she-document-state}\'s
    [navigable target
    name](browsing-the-web.html#document-state-nav-target-name){#apis-for-creating-and-navigating-browsing-contexts-by-name:document-state-nav-target-name}
    to the given value.

The name [gets reset](browsing-the-web.html#resetBCName) when the
navigable is
[navigated](browsing-the-web.html#navigate){#apis-for-creating-and-navigating-browsing-contexts-by-name:navigate-4}
to another
[origin](browsers.html#concept-origin){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-origin}.

------------------------------------------------------------------------

The [`close()`]{#dom-window-close .dfn dfn-for="Window"
dfn-type="method"} method steps are:

1.  Let `thisTraversable`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this-5
    x-internal="this"}\'s
    [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable-5}.

2.  If `thisTraversable`{.variable} is not a [top-level
    traversable](document-sequences.html#top-level-traversable){#apis-for-creating-and-navigating-browsing-contexts-by-name:top-level-traversable},
    then return.

3.  If `thisTraversable`{.variable}\'s [is
    closing](document-sequences.html#is-closing){#apis-for-creating-and-navigating-browsing-contexts-by-name:is-closing}
    is true, then return.

4.  Let `browsingContext`{.variable} be `thisTraversable`{.variable}\'s
    [active browsing
    context](document-sequences.html#nav-bc){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-bc-4}.

5.  Let `sourceSnapshotParams`{.variable} be the result of [snapshotting
    source snapshot
    params](browsing-the-web.html#snapshotting-source-snapshot-params){#apis-for-creating-and-navigating-browsing-contexts-by-name:snapshotting-source-snapshot-params}
    given `thisTraversable`{.variable}\'s [active
    document](document-sequences.html#nav-document){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-document-2}.

6.  If all the following are true:

    - `thisTraversable`{.variable} is
      [script-closable](#script-closable){#apis-for-creating-and-navigating-browsing-contexts-by-name:script-closable};

    - the [incumbent global
      object](webappapis.html#concept-incumbent-global){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-incumbent-global}\'s
      [browsing
      context](#window-bc){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-bc}
      is [familiar
      with](document-sequences.html#familiar-with){#apis-for-creating-and-navigating-browsing-contexts-by-name:familiar-with}
      `browsingContext`{.variable}; and

    - ::: {#sandboxClose}
      the [incumbent global
      object](webappapis.html#concept-incumbent-global){#apis-for-creating-and-navigating-browsing-contexts-by-name:concept-incumbent-global-2}\'s
      [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable-6}
      is [allowed by sandboxing to
      navigate](browsing-the-web.html#allowed-to-navigate){#apis-for-creating-and-navigating-browsing-contexts-by-name:allowed-to-navigate}
      `thisTraversable`{.variable}, given
      `sourceSnapshotParams`{.variable},
      :::

    then:

    1.  Set `thisTraversable`{.variable}\'s [is
        closing](document-sequences.html#is-closing){#apis-for-creating-and-navigating-browsing-contexts-by-name:is-closing-2}
        to true.

    2.  [Queue a
        task](webappapis.html#queue-a-task){#apis-for-creating-and-navigating-browsing-contexts-by-name:queue-a-task}
        on the [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#apis-for-creating-and-navigating-browsing-contexts-by-name:dom-manipulation-task-source}
        to [definitely
        close](document-sequences.html#definitely-close-a-top-level-traversable){#apis-for-creating-and-navigating-browsing-contexts-by-name:definitely-close-a-top-level-traversable}
        `thisTraversable`{.variable}.

A
[navigable](document-sequences.html#navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:navigable}
is [script-closable]{#script-closable .dfn} if it is a [top-level
traversable](document-sequences.html#top-level-traversable){#apis-for-creating-and-navigating-browsing-contexts-by-name:top-level-traversable-2},
and any of the following are true:

- its [is created by web
  content](document-sequences.html#is-created-by-web-content){#apis-for-creating-and-navigating-browsing-contexts-by-name:is-created-by-web-content}
  is true; or
- its [session history
  entries](document-sequences.html#tn-session-history-entries){#apis-for-creating-and-navigating-browsing-contexts-by-name:tn-session-history-entries}\'s
  [size](https://infra.spec.whatwg.org/#list-size){#apis-for-creating-and-navigating-browsing-contexts-by-name:list-size
  x-internal="list-size"} is 1.

The [`closed`]{#dom-window-closed .dfn dfn-for="Window"
dfn-type="attribute"} getter steps are to return true if
[this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this-6
x-internal="this"}\'s [browsing
context](#window-bc){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-bc-2}
is null or its [is
closing](document-sequences.html#is-closing){#apis-for-creating-and-navigating-browsing-contexts-by-name:is-closing-3}
is true; otherwise false.

The [`stop()`]{#dom-window-stop .dfn dfn-for="Window" dfn-type="method"}
method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this-7
    x-internal="this"}\'s
    [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable-7}
    is null, then return.

2.  [Stop
    loading](document-lifecycle.html#nav-stop){#apis-for-creating-and-navigating-browsing-contexts-by-name:nav-stop}
    [this](https://webidl.spec.whatwg.org/#this){#apis-for-creating-and-navigating-browsing-contexts-by-name:this-8
    x-internal="this"}\'s
    [navigable](#window-navigable){#apis-for-creating-and-navigating-browsing-contexts-by-name:window-navigable-8}.

##### [7.2.2.2]{.secno} Indexed access on the [`Window`{#accessing-other-browsing-contexts:window}](#window) object[](#accessing-other-browsing-contexts){.self-link} {#accessing-other-browsing-contexts}

`window`{.variable}`.`[`length`](#dom-length){#dom-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/length](https://developer.mozilla.org/en-US/docs/Web/API/Window/length "Returns the number of frames (either <frame> or <iframe> elements) in the window.")

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

Returns the number of [document-tree child
navigables](document-sequences.html#document-tree-child-navigables){#accessing-other-browsing-contexts:document-tree-child-navigables}.

`window`{.variable}`[``index`{.variable}`]`

Returns the
[`WindowProxy`{#accessing-other-browsing-contexts:windowproxy}](#windowproxy)
corresponding to the indicated [document-tree child
navigables](document-sequences.html#document-tree-child-navigables){#accessing-other-browsing-contexts:document-tree-child-navigables-2}.

The [`length`]{#dom-length .dfn dfn-for="Window" dfn-type="attribute"}
getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#accessing-other-browsing-contexts:this
x-internal="this"}\'s [associated
`Document`](#concept-document-window){#accessing-other-browsing-contexts:concept-document-window}\'s
[document-tree child
navigables](document-sequences.html#document-tree-child-navigables){#accessing-other-browsing-contexts:document-tree-child-navigables-3}\'s
[size](https://infra.spec.whatwg.org/#list-size){#accessing-other-browsing-contexts:list-size
x-internal="list-size"}.

Indexed access to [document-tree child
navigables](document-sequences.html#document-tree-child-navigables){#accessing-other-browsing-contexts:document-tree-child-navigables-4}
is defined through the
[\[\[GetOwnProperty\]\]](#windowproxy-getownproperty) internal method of
the
[`WindowProxy`{#accessing-other-browsing-contexts:windowproxy-2}](#windowproxy)
object.

##### [7.2.2.3]{.secno} Named access on the [`Window`{#named-access-on-the-window-object:window}](#window) object[](#named-access-on-the-window-object){.self-link}

`window`{.variable}`[``name`{.variable}`]`

:   Returns the indicated element or collection of elements.

    As a general rule, relying on this will lead to brittle code. Which
    IDs end up mapping to this API can vary over time, as new features
    are added to the web platform, for example. Instead of this, use
    `document.getElementById()` or `document.querySelector()`.

[]{#document-tree-child-browsing-context-name-property-set}The
[document-tree child navigable target name property
set]{#document-tree-child-navigable-target-name-property-set .dfn} of a
[`Window`{#named-access-on-the-window-object:window-2}](#window) object
`window`{.variable} is the return value of running these steps:

1.  Let `children`{.variable} be the [document-tree child
    navigables](document-sequences.html#document-tree-child-navigables){#named-access-on-the-window-object:document-tree-child-navigables}
    of `window`{.variable}\'s [associated
    `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window}.

2.  Let `firstNamedChildren`{.variable} be an empty [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#named-access-on-the-window-object:set
    x-internal="set"}.

3.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#named-access-on-the-window-object:list-iterate
    x-internal="list-iterate"} `navigable`{.variable} of
    `children`{.variable}:

    1.  Let `name`{.variable} be `navigable`{.variable}\'s [target
        name](document-sequences.html#nav-target){#named-access-on-the-window-object:nav-target}.

    2.  If `name`{.variable} is the empty string, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#named-access-on-the-window-object:continue
        x-internal="continue"}.

    3.  If `firstNamedChildren`{.variable}
        [contains](https://infra.spec.whatwg.org/#list-contain){#named-access-on-the-window-object:list-contains
        x-internal="list-contains"} a
        [navigable](document-sequences.html#navigable){#named-access-on-the-window-object:navigable}
        whose [target
        name](document-sequences.html#nav-target){#named-access-on-the-window-object:nav-target-2}
        is `name`{.variable}, then
        [continue](https://infra.spec.whatwg.org/#iteration-continue){#named-access-on-the-window-object:continue-2
        x-internal="continue"}.

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#named-access-on-the-window-object:list-append
        x-internal="list-append"} `navigable`{.variable} to
        `firstNamedChildren`{.variable}.

4.  Let `names`{.variable} be an empty [ordered
    set](https://infra.spec.whatwg.org/#ordered-set){#named-access-on-the-window-object:set-2
    x-internal="set"}.

5.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#named-access-on-the-window-object:list-iterate-2
    x-internal="list-iterate"} `navigable`{.variable} of
    `firstNamedChildren`{.variable}:

    1.  Let `name`{.variable} be `navigable`{.variable}\'s [target
        name](document-sequences.html#nav-target){#named-access-on-the-window-object:nav-target-3}.

    2.  If `navigable`{.variable}\'s [active
        document](document-sequences.html#nav-document){#named-access-on-the-window-object:nav-document}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#named-access-on-the-window-object:concept-document-origin
        x-internal="concept-document-origin"} is [same
        origin](browsers.html#same-origin){#named-access-on-the-window-object:same-origin}
        with `window`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#named-access-on-the-window-object:relevant-settings-object}\'s
        [origin](webappapis.html#concept-settings-object-origin){#named-access-on-the-window-object:concept-settings-object-origin},
        then
        [append](https://infra.spec.whatwg.org/#list-append){#named-access-on-the-window-object:list-append-2
        x-internal="list-append"} `name`{.variable} to
        `names`{.variable}.

6.  Return `names`{.variable}.

::: example
The two seperate iterations mean that in the following example, hosted
on `https://example.org/`, assuming `https://elsewhere.example/` sets
[`window.name`{#named-access-on-the-window-object:dom-name}](#dom-name)
to \"`spices`\", evaluating `window.spices` after everything has loaded
will yield undefined:

``` html
<iframe src=https://elsewhere.example.com/></iframe>
<iframe name=spices></iframe>
```
:::

The [`Window`{#named-access-on-the-window-object:window-3}](#window)
object [supports named
properties](https://webidl.spec.whatwg.org/#dfn-support-named-properties){#named-access-on-the-window-object:support-named-properties
x-internal="support-named-properties"}. The [supported property
names](https://webidl.spec.whatwg.org/#dfn-supported-property-names){#named-access-on-the-window-object:supported-property-names
x-internal="supported-property-names"} of a
[`Window`{#named-access-on-the-window-object:window-4}](#window) object
`window`{.variable} at any moment consist of the following, in [tree
order](https://dom.spec.whatwg.org/#concept-tree-order){#named-access-on-the-window-object:tree-order
x-internal="tree-order"} according to the element that contributed them,
ignoring later duplicates:

- `window`{.variable}\'s [document-tree child navigable target name
  property
  set](#document-tree-child-navigable-target-name-property-set){#named-access-on-the-window-object:document-tree-child-navigable-target-name-property-set};

- the value of the `name` content attribute for all
  [`embed`{#named-access-on-the-window-object:the-embed-element}](iframe-embed-object.html#the-embed-element),
  [`form`{#named-access-on-the-window-object:the-form-element}](forms.html#the-form-element),
  [`img`{#named-access-on-the-window-object:the-img-element}](embedded-content.html#the-img-element),
  and
  [`object`{#named-access-on-the-window-object:the-object-element}](iframe-embed-object.html#the-object-element)
  elements that have a non-empty `name` content attribute and are [in a
  document
  tree](https://dom.spec.whatwg.org/#in-a-document-tree){#named-access-on-the-window-object:in-a-document-tree
  x-internal="in-a-document-tree"} with `window`{.variable}\'s
  [associated
  `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window-2}
  as their
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#named-access-on-the-window-object:root
  x-internal="root"}; and

- the value of the
  [`id`{#named-access-on-the-window-object:the-id-attribute}](dom.html#the-id-attribute)
  content attribute for all [HTML
  elements](infrastructure.html#html-elements){#named-access-on-the-window-object:html-elements}
  that have a non-empty
  [`id`{#named-access-on-the-window-object:the-id-attribute-2}](dom.html#the-id-attribute)
  content attribute and are [in a document
  tree](https://dom.spec.whatwg.org/#in-a-document-tree){#named-access-on-the-window-object:in-a-document-tree-2
  x-internal="in-a-document-tree"} with `window`{.variable}\'s
  [associated
  `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window-3}
  as their
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#named-access-on-the-window-object:root-2
  x-internal="root"}.

To [determine the value of a named
property](https://webidl.spec.whatwg.org/#dfn-determine-the-value-of-a-named-property){#named-access-on-the-window-object:determine-the-value-of-a-named-property
x-internal="determine-the-value-of-a-named-property"} `name`{.variable}
in a [`Window`{#named-access-on-the-window-object:window-5}](#window)
object `window`{.variable}, the user agent must return the value
obtained using the following steps:

1.  Let `objects`{.variable} be the list of [named
    objects](#dom-window-nameditem-filter){#named-access-on-the-window-object:dom-window-nameditem-filter}
    of `window`{.variable} with the name `name`{.variable}.

    There will be at least one such object, since the algorithm would
    otherwise not have been [invoked by Web
    IDL](https://webidl.spec.whatwg.org/#named-properties-object-getownproperty){#named-access-on-the-window-object:named-properties-object-getownproperty
    x-internal="named-properties-object-getownproperty"}.

2.  If `objects`{.variable} contains a
    [navigable](document-sequences.html#navigable){#named-access-on-the-window-object:navigable-2},
    then:

    1.  Let `container`{.variable} be the first [navigable
        container](document-sequences.html#navigable-container){#named-access-on-the-window-object:navigable-container}
        in `window`{.variable}\'s [associated
        `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window-4}\'s
        [descendants](https://dom.spec.whatwg.org/#concept-tree-descendant){#named-access-on-the-window-object:descendant
        x-internal="descendant"} whose [content
        navigable](document-sequences.html#content-navigable){#named-access-on-the-window-object:content-navigable}
        is in `objects`{.variable}.

    2.  Return `container`{.variable}\'s [content
        navigable](document-sequences.html#content-navigable){#named-access-on-the-window-object:content-navigable-2}\'s
        [active
        `WindowProxy`](document-sequences.html#nav-wp){#named-access-on-the-window-object:nav-wp}.

3.  Otherwise, if `objects`{.variable} has only one element, return that
    element.

4.  Otherwise, return an
    [`HTMLCollection`{#named-access-on-the-window-object:htmlcollection}](https://dom.spec.whatwg.org/#interface-htmlcollection){x-internal="htmlcollection"}
    rooted at `window`{.variable}\'s [associated
    `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window-5},
    whose filter matches only [named
    objects](#dom-window-nameditem-filter){#named-access-on-the-window-object:dom-window-nameditem-filter-2}
    of `window`{.variable} with the name `name`{.variable}. (By
    definition, these will all be elements.)

[Named objects]{#dom-window-nameditem-filter .dfn} of
[`Window`{#named-access-on-the-window-object:window-6}](#window) object
`window`{.variable} with the name `name`{.variable}, for the purposes of
the above algorithm, consist of the following:

- [document-tree child
  navigables](document-sequences.html#document-tree-child-navigables){#named-access-on-the-window-object:document-tree-child-navigables-2}
  of `window`{.variable}\'s [associated
  `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window-6}
  whose [target
  name](document-sequences.html#nav-target){#named-access-on-the-window-object:nav-target-4}
  is `name`{.variable};

- [`embed`{#named-access-on-the-window-object:the-embed-element-2}](iframe-embed-object.html#the-embed-element),
  [`form`{#named-access-on-the-window-object:the-form-element-2}](forms.html#the-form-element),
  [`img`{#named-access-on-the-window-object:the-img-element-2}](embedded-content.html#the-img-element),
  or
  [`object`{#named-access-on-the-window-object:the-object-element-2}](iframe-embed-object.html#the-object-element)
  elements that have a `name` content attribute whose value is
  `name`{.variable} and are [in a document
  tree](https://dom.spec.whatwg.org/#in-a-document-tree){#named-access-on-the-window-object:in-a-document-tree-3
  x-internal="in-a-document-tree"} with `window`{.variable}\'s
  [associated
  `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window-7}
  as their
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#named-access-on-the-window-object:root-3
  x-internal="root"}; and

- [HTML
  elements](infrastructure.html#html-elements){#named-access-on-the-window-object:html-elements-2}
  that have an
  [`id`{#named-access-on-the-window-object:the-id-attribute-3}](dom.html#the-id-attribute)
  content attribute whose value is `name`{.variable} and are [in a
  document
  tree](https://dom.spec.whatwg.org/#in-a-document-tree){#named-access-on-the-window-object:in-a-document-tree-4
  x-internal="in-a-document-tree"} with `window`{.variable}\'s
  [associated
  `Document`](#concept-document-window){#named-access-on-the-window-object:concept-document-window-8}
  as their
  [root](https://dom.spec.whatwg.org/#concept-tree-root){#named-access-on-the-window-object:root-4
  x-internal="root"}.

Since the
[`Window`{#named-access-on-the-window-object:window-7}](#window)
interface has the
[`[Global]`{#named-access-on-the-window-object:global}](https://webidl.spec.whatwg.org/#Global){x-internal="global"}
extended attribute, its named properties follow the rules for [named
properties
objects](https://webidl.spec.whatwg.org/#dfn-named-properties-object){#named-access-on-the-window-object:named-properties-object
x-internal="named-properties-object"} rather than [legacy platform
objects](https://webidl.spec.whatwg.org/#dfn-legacy-platform-object){#named-access-on-the-window-object:legacy-platform-object
x-internal="legacy-platform-object"}.

##### [7.2.2.4]{.secno} Accessing related windows[](#navigating-nested-browsing-contexts-in-the-dom){.self-link} {#navigating-nested-browsing-contexts-in-the-dom}

`window`{.variable}`.`[`top`](#dom-top){#dom-top-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/top](https://developer.mozilla.org/en-US/docs/Web/API/Window/top "Returns a reference to the topmost window in the window hierarchy.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android4+]{.firefox_android .yes}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[`WindowProxy`{#navigating-nested-browsing-contexts-in-the-dom:windowproxy}](#windowproxy)
for the [top-level
traversable](document-sequences.html#top-level-traversable){#navigating-nested-browsing-contexts-in-the-dom:top-level-traversable}.

`window`{.variable}`.`[`opener`](#dom-opener){#dom-opener-dev}` [ = ``value`{.variable}` ]`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/opener](https://developer.mozilla.org/en-US/docs/Web/API/Window/opener "The Window interface's opener property returns a reference to the window that opened the window, either with open(), or by navigating a link with a target attribute.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[`WindowProxy`{#navigating-nested-browsing-contexts-in-the-dom:windowproxy-2}](#windowproxy)
for the [opener browsing
context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context}.

Returns null if there isn\'t one or if it has been set to null.

Can be set to null.

`window`{.variable}`.`[`parent`](#dom-parent){#dom-parent-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/parent](https://developer.mozilla.org/en-US/docs/Web/API/Window/parent "The Window.parent property is a reference to the parent of the current window or subframe.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1.3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer9+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[`WindowProxy`{#navigating-nested-browsing-contexts-in-the-dom:windowproxy-3}](#windowproxy)
for the [parent
navigable](document-sequences.html#nav-parent){#navigating-nested-browsing-contexts-in-the-dom:nav-parent}.

`window`{.variable}`.`[`frameElement`](#dom-frameelement){#dom-frameelement-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/frameElement](https://developer.mozilla.org/en-US/docs/Web/API/Window/frameElement "The Window.frameElement property returns the element (such as <iframe> or <object>) in which the window is embedded.")

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

Returns the [navigable
container](document-sequences.html#navigable-container){#navigating-nested-browsing-contexts-in-the-dom:navigable-container}
element.

Returns null if there isn\'t one, and in cross-origin situations.

The [`top`]{#dom-top .dfn dfn-for="Window" dfn-type="attribute"} getter
steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this
    x-internal="this"}\'s
    [navigable](#window-navigable){#navigating-nested-browsing-contexts-in-the-dom:window-navigable}
    is null, then return null.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this-2
    x-internal="this"}\'s
    [navigable](#window-navigable){#navigating-nested-browsing-contexts-in-the-dom:window-navigable-2}\'s
    [top-level
    traversable](document-sequences.html#nav-top){#navigating-nested-browsing-contexts-in-the-dom:nav-top}\'s
    [active
    `WindowProxy`](document-sequences.html#nav-wp){#navigating-nested-browsing-contexts-in-the-dom:nav-wp}.

The [`opener`]{#dom-opener .dfn dfn-for="Window" dfn-type="attribute"}
getter steps are:

1.  Let `current`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this-3
    x-internal="this"}\'s [browsing
    context](#window-bc){#navigating-nested-browsing-contexts-in-the-dom:window-bc}.

2.  If `current`{.variable} is null, then return null.

3.  If `current`{.variable}\'s [opener browsing
    context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context-2}
    is null, then return null.

4.  Return `current`{.variable}\'s [opener browsing
    context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context-3}\'s
    [`WindowProxy`{#navigating-nested-browsing-contexts-in-the-dom:windowproxy-4}](#windowproxy)
    object.

The
[`opener`{#navigating-nested-browsing-contexts-in-the-dom:dom-opener}](#dom-opener)
setter steps are:

1.  If the given value is null and
    [this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this-4
    x-internal="this"}\'s [browsing
    context](#window-bc){#navigating-nested-browsing-contexts-in-the-dom:window-bc-2}
    is non-null, then set
    [this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this-5
    x-internal="this"}\'s [browsing
    context](#window-bc){#navigating-nested-browsing-contexts-in-the-dom:window-bc-3}\'s
    [opener browsing
    context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context-4}
    to null.

2.  If the given value is non-null, then perform ?
    [DefinePropertyOrThrow](https://tc39.es/ecma262/#sec-definepropertyorthrow){#navigating-nested-browsing-contexts-in-the-dom:definepropertyorthrow
    x-internal="definepropertyorthrow"}([this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this-6
    x-internal="this"}, \"`opener`\", { \[\[Value\]\]: the given value,
    \[\[Writable\]\]: true, \[\[Enumerable\]\]: true,
    \[\[Configurable\]\]: true }).

::: note
Setting
[`window.opener`{#navigating-nested-browsing-contexts-in-the-dom:dom-opener-2}](#dom-opener)
to null clears the [opener browsing
context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context-5}
reference. In practice, this prevents future scripts from accessing
their [opener browsing
context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context-6}\'s
[`Window`{#navigating-nested-browsing-contexts-in-the-dom:window}](#window)
object.

By default, scripts can access their [opener browsing
context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context-7}\'s
[`Window`{#navigating-nested-browsing-contexts-in-the-dom:window-2}](#window)
object through the
[`window.opener`{#navigating-nested-browsing-contexts-in-the-dom:dom-opener-3}](#dom-opener)
getter. E.g., a script can set `window.opener.location`, causing the
[opener browsing
context](document-sequences.html#opener-browsing-context){#navigating-nested-browsing-contexts-in-the-dom:opener-browsing-context-8}
to navigate.
:::

The [`parent`]{#dom-parent .dfn dfn-for="Window" dfn-type="attribute"}
getter steps are:

1.  Let `navigable`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this-7
    x-internal="this"}\'s
    [navigable](#window-navigable){#navigating-nested-browsing-contexts-in-the-dom:window-navigable-3}.

2.  If `navigable`{.variable} is null, then return null.

3.  If `navigable`{.variable}\'s
    [parent](document-sequences.html#nav-parent){#navigating-nested-browsing-contexts-in-the-dom:nav-parent-2}
    is not null, then set `navigable`{.variable} to
    `navigable`{.variable}\'s
    [parent](document-sequences.html#nav-parent){#navigating-nested-browsing-contexts-in-the-dom:nav-parent-3}.

4.  Return `navigable`{.variable}\'s [active
    `WindowProxy`](document-sequences.html#nav-wp){#navigating-nested-browsing-contexts-in-the-dom:nav-wp-2}.

The [`frameElement`]{#dom-frameelement .dfn dfn-for="Window"
dfn-type="attribute"} getter steps are:

1.  Let `current`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#navigating-nested-browsing-contexts-in-the-dom:this-8
    x-internal="this"}\'s [node
    navigable](document-sequences.html#node-navigable){#navigating-nested-browsing-contexts-in-the-dom:node-navigable}.

2.  If `current`{.variable} is null, then return null.

3.  Let `container`{.variable} be `current`{.variable}\'s
    [container](document-sequences.html#nav-container){#navigating-nested-browsing-contexts-in-the-dom:nav-container}.

4.  If `container`{.variable} is null, then return null.

5.  If `container`{.variable}\'s [node
    document](https://dom.spec.whatwg.org/#concept-node-document){#navigating-nested-browsing-contexts-in-the-dom:node-document
    x-internal="node-document"}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#navigating-nested-browsing-contexts-in-the-dom:concept-document-origin
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#navigating-nested-browsing-contexts-in-the-dom:same-origin-domain}
    with the [current settings
    object](webappapis.html#current-settings-object){#navigating-nested-browsing-contexts-in-the-dom:current-settings-object}\'s
    [origin](webappapis.html#concept-settings-object-origin){#navigating-nested-browsing-contexts-in-the-dom:concept-settings-object-origin},
    then return null.

6.  Return `container`{.variable}.

::: example
An example of when these properties can return null is as follows:

``` html
<!DOCTYPE html>
<iframe></iframe>

<script>
"use strict";
const element = document.querySelector("iframe");
const iframeWindow = element.contentWindow;
element.remove();

console.assert(iframeWindow.top === null);
console.assert(iframeWindow.parent === null);
console.assert(iframeWindow.frameElement === null);
</script>
```

Here the [browsing
context](document-sequences.html#browsing-context){#navigating-nested-browsing-contexts-in-the-dom:browsing-context}
corresponding to `iframeWindow` was [nulled
out](document-lifecycle.html#destroy-a-document){#navigating-nested-browsing-contexts-in-the-dom:destroy-a-document}
when `element` was removed from the document.
:::

##### [7.2.2.5]{.secno} Historical browser interface element APIs[](#browser-interface-elements){.self-link} {#browser-interface-elements}

For historical reasons, the
[`Window`{#browser-interface-elements:window}](#window) interface had
some properties that represented the visibility of certain web browser
interface elements.

For privacy and interoperability reasons, those properties now return
values that represent whether the
[`Window`{#browser-interface-elements:window-2}](#window)\'s [browsing
context](#window-bc){#browser-interface-elements:window-bc}\'s [is
popup](document-sequences.html#is-popup){#browser-interface-elements:is-popup}
property is true or false.

Each interface element is represented by a
[`BarProp`{#browser-interface-elements:barprop}](#barprop) object:

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BarProp](https://developer.mozilla.org/en-US/docs/Web/API/BarProp "The BarProp interface of the Document Object Model represents the web browser user interface elements that are exposed to scripts in web pages. Each of the following interface elements are represented by a BarProp object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

``` idl
[Exposed=Window]
interface BarProp {
  readonly attribute boolean visible;
};
```

`window`{.variable}`.`[`locationbar`](#dom-window-locationbar){#dom-window-locationbar-dev}`.`[`visible`](#dom-barprop-visible){#dom-barprop-visible-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/locationbar](https://developer.mozilla.org/en-US/docs/Web/API/Window/locationbar "Returns the locationbar object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`window`{.variable}`.`[`menubar`](#dom-window-menubar){#dom-window-menubar-dev}`.`[`visible`](#dom-barprop-visible){#browser-interface-elements:dom-barprop-visible-2}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/menubar](https://developer.mozilla.org/en-US/docs/Web/API/Window/menubar "Returns the menubar object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`window`{.variable}`.`[`personalbar`](#dom-window-personalbar){#dom-window-personalbar-dev}`.`[`visible`](#dom-barprop-visible){#browser-interface-elements:dom-barprop-visible-3}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/personalbar](https://developer.mozilla.org/en-US/docs/Web/API/Window/personalbar "Returns the personalbar object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`window`{.variable}`.`[`scrollbars`](#dom-window-scrollbars){#dom-window-scrollbars-dev}`.`[`visible`](#dom-barprop-visible){#browser-interface-elements:dom-barprop-visible-4}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/scrollbars](https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollbars "Returns the scrollbars object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

`window`{.variable}`.`[`statusbar`](#dom-window-statusbar){#dom-window-statusbar-dev}`.`[`visible`](#dom-barprop-visible){#browser-interface-elements:dom-barprop-visible-5}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/statusbar](https://developer.mozilla.org/en-US/docs/Web/API/Window/statusbar "Returns the statusbar object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

`window`{.variable}`.`[`toolbar`](#dom-window-toolbar){#dom-window-toolbar-dev}`.`[`visible`](#dom-barprop-visible){#browser-interface-elements:dom-barprop-visible-6}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Window/toolbar](https://developer.mozilla.org/en-US/docs/Web/API/Window/toolbar "Returns the toolbar object.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns true if the
[`Window`{#browser-interface-elements:window-3}](#window) is not a
popup; otherwise, returns false.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[BarProp/visible](https://developer.mozilla.org/en-US/docs/Web/API/BarProp/visible "The visible read-only property of the BarProp interface returns true if the user interface element it represents is visible.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS1+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`visible`]{#dom-barprop-visible .dfn dfn-for="BarProp"
dfn-type="attribute"} getter steps are:

1.  Let `browsingContext`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#browser-interface-elements:this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#browser-interface-elements:concept-relevant-global}\'s
    [browsing
    context](#window-bc){#browser-interface-elements:window-bc-2}.

2.  If `browsingContext`{.variable} is null, then return true.

3.  Return the negation of `browsingContext`{.variable}\'s [top-level
    browsing
    context](document-sequences.html#bc-tlbc){#browser-interface-elements:bc-tlbc}\'s
    [is
    popup](document-sequences.html#is-popup){#browser-interface-elements:is-popup-2}.

The following
[`BarProp`{#browser-interface-elements:barprop-2}](#barprop) objects
must exist for each
[`Window`{#browser-interface-elements:window-4}](#window) object:

[The location bar `BarProp` object]{#the-location-bar-barprop-object .dfn}
:   Historically represented the user interface element that contains a
    control that displays the browser\'s location bar.

[The menu bar `BarProp` object]{#the-menu-bar-barprop-object .dfn}
:   Historically represented the user interface element that contains a
    list of commands in menu form, or some similar interface concept.

[The personal bar `BarProp` object]{#the-personal-bar-barprop-object .dfn}
:   Historically represented the user interface element that contains
    links to the user\'s favorite pages, or some similar interface
    concept.

[The scrollbar `BarProp` object]{#the-scrollbar-barprop-object .dfn}
:   Historically represented the user interface element that contains a
    scrolling mechanism, or some similar interface concept.

[The status bar `BarProp` object]{#the-status-bar-barprop-object .dfn}
:   Historically represented a user interface element found immediately
    below or after the document, as appropriate for the user\'s media,
    which typically provides information about ongoing network activity
    or information about elements that the user\'s pointing device is
    currently indicating.

[The toolbar `BarProp` object]{#the-toolbar-barprop-object .dfn}
:   Historically represented the user interface element found
    immediately above or before the document, as appropriate for the
    user\'s media, which typically provides [session history
    traversal](browsing-the-web.html#apply-the-traverse-history-step){#browser-interface-elements:apply-the-traverse-history-step}
    controls (back and forward buttons, reload buttons, etc.).

The [`locationbar`]{#dom-window-locationbar .dfn dfn-for="Window"
dfn-type="attribute"} attribute must return [the location bar `BarProp`
object](#the-location-bar-barprop-object){#browser-interface-elements:the-location-bar-barprop-object}.

The [`menubar`]{#dom-window-menubar .dfn dfn-for="Window"
dfn-type="attribute"} attribute must return [the menu bar `BarProp`
object](#the-menu-bar-barprop-object){#browser-interface-elements:the-menu-bar-barprop-object}.

The [`personalbar`]{#dom-window-personalbar .dfn dfn-for="Window"
dfn-type="attribute"} attribute must return [the personal bar `BarProp`
object](#the-personal-bar-barprop-object){#browser-interface-elements:the-personal-bar-barprop-object}.

The [`scrollbars`]{#dom-window-scrollbars .dfn dfn-for="Window"
dfn-type="attribute"} attribute must return [the scrollbar `BarProp`
object](#the-scrollbar-barprop-object){#browser-interface-elements:the-scrollbar-barprop-object}.

The [`statusbar`]{#dom-window-statusbar .dfn dfn-for="Window"
dfn-type="attribute"} attribute must return [the status bar `BarProp`
object](#the-status-bar-barprop-object){#browser-interface-elements:the-status-bar-barprop-object}.

The [`toolbar`]{#dom-window-toolbar .dfn dfn-for="Window"
dfn-type="attribute"} attribute must return [the toolbar `BarProp`
object](#the-toolbar-barprop-object){#browser-interface-elements:the-toolbar-barprop-object}.

------------------------------------------------------------------------

For historical reasons, the [`status`]{#dom-window-status .dfn
dfn-for="Window" dfn-type="attribute"} attribute on the
[`Window`{#browser-interface-elements:window-5}](#window) object must,
on getting, return the last string it was set to, and on setting, must
set itself to the new value. When the
[`Window`{#browser-interface-elements:window-6}](#window) object is
created, the attribute must be set to the empty string. It does not do
anything else.

##### [7.2.2.6]{.secno} Script settings for [`Window`{#script-settings-for-window-objects:window}](#window) objects[](#script-settings-for-window-objects){.self-link}

To [set up a window environment settings
object]{#set-up-a-window-environment-settings-object .dfn}, given a
[URL](https://url.spec.whatwg.org/#concept-url){#script-settings-for-window-objects:url
x-internal="url"} `creationURL`{.variable}, a [JavaScript execution
context](https://tc39.es/ecma262/#sec-execution-contexts){#script-settings-for-window-objects:javascript-execution-context
x-internal="javascript-execution-context"}
`execution context`{.variable}, null or an
[environment](webappapis.html#environment){#script-settings-for-window-objects:environment}
`reservedEnvironment`{.variable}, a
[URL](https://url.spec.whatwg.org/#concept-url){#script-settings-for-window-objects:url-2
x-internal="url"} `topLevelCreationURL`{.variable}, and an
[origin](browsers.html#concept-origin){#script-settings-for-window-objects:concept-origin}
`topLevelOrigin`{.variable}, run these steps:

1.  Let `realm`{.variable} be the value of
    `execution context`{.variable}\'s Realm component.

2.  Let `window`{.variable} be `realm`{.variable}\'s [global
    object](webappapis.html#concept-realm-global){#script-settings-for-window-objects:concept-realm-global}.

3.  Let `settings object`{.variable} be a new [environment settings
    object](webappapis.html#environment-settings-object){#script-settings-for-window-objects:environment-settings-object}
    whose algorithms are defined as follows:

    The [realm execution context](webappapis.html#realm-execution-context){#script-settings-for-window-objects:realm-execution-context}

    :   Return `execution context`{.variable}.

    The [module map](webappapis.html#concept-settings-object-module-map){#script-settings-for-window-objects:concept-settings-object-module-map}

    :   Return the [module
        map](dom.html#concept-document-module-map){#script-settings-for-window-objects:concept-document-module-map}
        of `window`{.variable}\'s [associated
        `Document`](#concept-document-window){#script-settings-for-window-objects:concept-document-window}.

    The [API base URL](webappapis.html#api-base-url){#script-settings-for-window-objects:api-base-url}

    :   Return the current [base
        URL](urls-and-fetching.html#document-base-url){#script-settings-for-window-objects:document-base-url}
        of `window`{.variable}\'s [associated
        `Document`](#concept-document-window){#script-settings-for-window-objects:concept-document-window-2}.

    The [origin](webappapis.html#concept-settings-object-origin){#script-settings-for-window-objects:concept-settings-object-origin}

    :   Return the
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#script-settings-for-window-objects:concept-document-origin
        x-internal="concept-document-origin"} of `window`{.variable}\'s
        [associated
        `Document`](#concept-document-window){#script-settings-for-window-objects:concept-document-window-3}.

    The [policy container](webappapis.html#concept-settings-object-policy-container){#script-settings-for-window-objects:concept-settings-object-policy-container}

    :   Return the [policy
        container](dom.html#concept-document-policy-container){#script-settings-for-window-objects:concept-document-policy-container}
        of `window`{.variable}\'s [associated
        `Document`](#concept-document-window){#script-settings-for-window-objects:concept-document-window-4}.

    The [cross-origin isolated capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#script-settings-for-window-objects:concept-settings-object-cross-origin-isolated-capability}

    :   Return true if both of the following hold, and false otherwise:

        - `realm`{.variable}\'s [agent
          cluster](https://tc39.es/ecma262/#sec-agent-clusters){#script-settings-for-window-objects:agent-cluster
          x-internal="agent-cluster"}\'s [cross-origin-isolation
          mode](webappapis.html#agent-cluster-cross-origin-isolation){#script-settings-for-window-objects:agent-cluster-cross-origin-isolation}
          is
          \"[`concrete`{#script-settings-for-window-objects:cross-origin-isolation-concrete}](document-sequences.html#cross-origin-isolation-concrete)\",
          and

        - `window`{.variable}\'s [associated
          `Document`](#concept-document-window){#script-settings-for-window-objects:concept-document-window-5}
          is [allowed to
          use](iframe-embed-object.html#allowed-to-use){#script-settings-for-window-objects:allowed-to-use}
          the
          \"[`cross-origin-isolated`{#script-settings-for-window-objects:cross-origin-isolated-feature}](infrastructure.html#cross-origin-isolated-feature)\"
          feature.

    The [time origin](webappapis.html#concept-settings-object-time-origin){#script-settings-for-window-objects:concept-settings-object-time-origin}

    :   Return `window`{.variable}\'s [associated
        `Document`](#concept-document-window){#script-settings-for-window-objects:concept-document-window-6}\'s
        [load timing
        info](dom.html#load-timing-info){#script-settings-for-window-objects:load-timing-info}\'s
        [navigation start
        time](dom.html#navigation-start-time){#script-settings-for-window-objects:navigation-start-time}.

4.  If `reservedEnvironment`{.variable} is non-null, then:

    1.  Set `settings object`{.variable}\'s
        [id](webappapis.html#concept-environment-id){#script-settings-for-window-objects:concept-environment-id}
        to `reservedEnvironment`{.variable}\'s
        [id](webappapis.html#concept-environment-id){#script-settings-for-window-objects:concept-environment-id-2},
        [target browsing
        context](webappapis.html#concept-environment-target-browsing-context){#script-settings-for-window-objects:concept-environment-target-browsing-context}
        to `reservedEnvironment`{.variable}\'s [target browsing
        context](webappapis.html#concept-environment-target-browsing-context){#script-settings-for-window-objects:concept-environment-target-browsing-context-2},
        and [active service
        worker](webappapis.html#concept-environment-active-service-worker){#script-settings-for-window-objects:concept-environment-active-service-worker}
        to `reservedEnvironment`{.variable}\'s [active service
        worker](webappapis.html#concept-environment-active-service-worker){#script-settings-for-window-objects:concept-environment-active-service-worker-2}.

    2.  Set `reservedEnvironment`{.variable}\'s
        [id](webappapis.html#concept-environment-id){#script-settings-for-window-objects:concept-environment-id-3}
        to the empty string.

        The identity of the reserved environment is considered to be
        fully transferred to the created [environment settings
        object](webappapis.html#environment-settings-object){#script-settings-for-window-objects:environment-settings-object-2}.
        The reserved environment is not searchable by the
        [environment](webappapis.html#environment){#script-settings-for-window-objects:environment-2}'s
        [id](webappapis.html#concept-environment-id){#script-settings-for-window-objects:concept-environment-id-4}
        from this point on.

5.  Otherwise, set `settings object`{.variable}\'s
    [id](webappapis.html#concept-environment-id){#script-settings-for-window-objects:concept-environment-id-5}
    to a new unique opaque string, `settings object`{.variable}\'s
    [target browsing
    context](webappapis.html#concept-environment-target-browsing-context){#script-settings-for-window-objects:concept-environment-target-browsing-context-3}
    to null, and `settings object`{.variable}\'s [active service
    worker](webappapis.html#concept-environment-active-service-worker){#script-settings-for-window-objects:concept-environment-active-service-worker-3}
    to null.

6.  Set `settings object`{.variable}\'s [creation
    URL](webappapis.html#concept-environment-creation-url){#script-settings-for-window-objects:concept-environment-creation-url}
    to `creationURL`{.variable}, `settings object`{.variable}\'s
    [top-level creation
    URL](webappapis.html#concept-environment-top-level-creation-url){#script-settings-for-window-objects:concept-environment-top-level-creation-url}
    to `topLevelCreationURL`{.variable}, and
    `settings object`{.variable}\'s [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#script-settings-for-window-objects:concept-environment-top-level-origin}
    to `topLevelOrigin`{.variable}.

7.  Set `realm`{.variable}\'s \[\[HostDefined\]\] field to
    `settings object`{.variable}.

#### [7.2.3]{.secno} The [`WindowProxy`{#the-windowproxy-exotic-object:windowproxy}](#windowproxy) exotic object[](#the-windowproxy-exotic-object){.self-link}

A [`WindowProxy`]{#windowproxy .dfn dfn-type="interface"} is an exotic
object that wraps a
[`Window`{#the-windowproxy-exotic-object:window}](#window) ordinary
object, indirecting most operations through to the wrapped object. Each
[browsing
context](document-sequences.html#browsing-context){#the-windowproxy-exotic-object:browsing-context}
has an associated
[`WindowProxy`{#the-windowproxy-exotic-object:windowproxy-2}](#windowproxy)
object. When the [browsing
context](document-sequences.html#browsing-context){#the-windowproxy-exotic-object:browsing-context-2}
is
[navigated](browsing-the-web.html#navigate){#the-windowproxy-exotic-object:navigate},
the [`Window`{#the-windowproxy-exotic-object:window-2}](#window) object
wrapped by the [browsing
context](document-sequences.html#browsing-context){#the-windowproxy-exotic-object:browsing-context-3}\'s
associated
[`WindowProxy`{#the-windowproxy-exotic-object:windowproxy-3}](#windowproxy)
object is changed.

The
[`WindowProxy`{#the-windowproxy-exotic-object:windowproxy-4}](#windowproxy)
exotic object must use the ordinary internal methods except where it is
explicitly specified otherwise below.

There is no
[`WindowProxy`{#the-windowproxy-exotic-object:windowproxy-5}](#windowproxy)
[interface
object](https://webidl.spec.whatwg.org/#dfn-interface-object){#the-windowproxy-exotic-object:interface-object
x-internal="interface-object"}.

Every
[`WindowProxy`{#the-windowproxy-exotic-object:windowproxy-6}](#windowproxy)
object has a [\[\[Window\]\]]{#concept-windowproxy-window .dfn
export=""} internal slot representing the wrapped
[`Window`{#the-windowproxy-exotic-object:window-3}](#window) object.

Although
[`WindowProxy`{#the-windowproxy-exotic-object:windowproxy-7}](#windowproxy)
is named as a \"proxy\", it does not do polymorphic dispatch on its
target\'s internal methods as a real proxy would, due to a desire to
reuse machinery between
[`WindowProxy`{#the-windowproxy-exotic-object:windowproxy-8}](#windowproxy)
and [`Location`{#the-windowproxy-exotic-object:location}](#location)
objects. As long as the
[`Window`{#the-windowproxy-exotic-object:window-4}](#window) object
remains an ordinary object this is unobservable and can be implemented
either way.

##### [7.2.3.1]{.secno} \[\[GetPrototypeOf\]\] ( )[](#windowproxy-getprototypeof){.self-link} {#windowproxy-getprototypeof}

1.  Let `W`{.variable} be the value of the
    [\[\[Window\]\]](#concept-windowproxy-window){#windowproxy-getprototypeof:concept-windowproxy-window}
    internal slot of **this**.

2.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-getprototypeof:isplatformobjectsameorigin-(-o-)}(`W`{.variable})
    is true, then return !
    [OrdinaryGetPrototypeOf](https://tc39.es/ecma262/#sec-ordinarygetprototypeof){#windowproxy-getprototypeof:ordinarygetprototypeof
    x-internal="ordinarygetprototypeof"}(`W`{.variable}).

3.  Return null.

##### [7.2.3.2]{.secno} \[\[SetPrototypeOf\]\] ( `V`{.variable} )[](#windowproxy-setprototypeof){.self-link} {#windowproxy-setprototypeof}

1.  Return !
    [SetImmutablePrototype](https://tc39.es/ecma262/#sec-set-immutable-prototype){#windowproxy-setprototypeof:setimmutableprototype
    x-internal="setimmutableprototype"}(**this**, `V`{.variable}).

##### [7.2.3.3]{.secno} \[\[IsExtensible\]\] ( )[](#windowproxy-isextensible){.self-link} {#windowproxy-isextensible}

1.  Return true.

##### [7.2.3.4]{.secno} \[\[PreventExtensions\]\] ( )[](#windowproxy-preventextensions){.self-link} {#windowproxy-preventextensions}

1.  Return false.

##### [7.2.3.5]{.secno} \[\[GetOwnProperty\]\] ( `P`{.variable} )[](#windowproxy-getownproperty){.self-link} {#windowproxy-getownproperty}

1.  Let `W`{.variable} be the value of the
    [\[\[Window\]\]](#concept-windowproxy-window){#windowproxy-getownproperty:concept-windowproxy-window}
    internal slot of **this**.

2.  If `P`{.variable} is an [array index property
    name](https://webidl.spec.whatwg.org/#dfn-array-index-property-name){#windowproxy-getownproperty:array-index-property-name
    x-internal="array-index-property-name"}, then:

    1.  Let `index`{.variable} be !
        [ToUint32](https://tc39.es/ecma262/#sec-touint32){#windowproxy-getownproperty:touint32
        x-internal="touint32"}(`P`{.variable}).

    2.  Let `children`{.variable} be the [document-tree child
        navigables](document-sequences.html#document-tree-child-navigables){#windowproxy-getownproperty:document-tree-child-navigables}
        of `W`{.variable}\'s [associated
        `Document`](#concept-document-window){#windowproxy-getownproperty:concept-document-window}.

    3.  Let `value`{.variable} be undefined.

    4.  If `index`{.variable} is less than `children`{.variable}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#windowproxy-getownproperty:list-size
        x-internal="list-size"}, then:

        1.  [Sort](https://infra.spec.whatwg.org/#list-sort-in-ascending-order){#windowproxy-getownproperty:list-sort
            x-internal="list-sort"} `children`{.variable} in ascending
            order, with `navigableA`{.variable} being less than
            `navigableB`{.variable} if `navigableA`{.variable}\'s
            [container](document-sequences.html#nav-container){#windowproxy-getownproperty:nav-container}
            was inserted into `W`{.variable}\'s [associated
            `Document`](#concept-document-window){#windowproxy-getownproperty:concept-document-window-2}
            earlier than `navigableB`{.variable}\'s
            [container](document-sequences.html#nav-container){#windowproxy-getownproperty:nav-container-2}
            was.

        2.  Set `value`{.variable} to
            `children`{.variable}\[`index`{.variable}\]\'s [active
            `WindowProxy`](document-sequences.html#nav-wp){#windowproxy-getownproperty:nav-wp}.

    5.  If `value`{.variable} is undefined, then:

        1.  If
            [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-getownproperty:isplatformobjectsameorigin-(-o-)}(`W`{.variable})
            is true, then return undefined.

        2.  Throw a
            [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#windowproxy-getownproperty:securityerror
            x-internal="securityerror"}
            [`DOMException`{#windowproxy-getownproperty:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    6.  Return
        [PropertyDescriptor](https://tc39.es/ecma262/#sec-property-descriptor-specification-type){#windowproxy-getownproperty:propertydescriptor
        x-internal="propertydescriptor"}{ \[\[Value\]\]:
        `value`{.variable}, \[\[Writable\]\]: false, \[\[Enumerable\]\]:
        true, \[\[Configurable\]\]: true }.

3.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-getownproperty:isplatformobjectsameorigin-(-o-)-2}(`W`{.variable})
    is true, then return !
    [OrdinaryGetOwnProperty](https://tc39.es/ecma262/#sec-ordinarygetownproperty){#windowproxy-getownproperty:ordinarygetownproperty
    x-internal="ordinarygetownproperty"}(`W`{.variable},
    `P`{.variable}).

    This is a [willful
    violation](https://infra.spec.whatwg.org/#willful-violation){#windowproxy-getownproperty:willful-violation
    x-internal="willful-violation"} of the JavaScript specification\'s
    [invariants of the essential internal
    methods](https://tc39.es/ecma262/#sec-invariants-of-the-essential-internal-methods){#windowproxy-getownproperty:invariants-of-the-essential-internal-methods
    x-internal="invariants-of-the-essential-internal-methods"} to
    maintain compatibility with existing web content. See [tc39/ecma262
    issue #672](https://github.com/tc39/ecma262/issues/672) for more
    information. [\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

4.  Let `property`{.variable} be
    [CrossOriginGetOwnPropertyHelper](#crossorigingetownpropertyhelper-(-o,-p-)){#windowproxy-getownproperty:crossorigingetownpropertyhelper-(-o,-p-)}(`W`{.variable},
    `P`{.variable}).

5.  If `property`{.variable} is not undefined, then return
    `property`{.variable}.

6.  If `property`{.variable} is undefined and `P`{.variable} is in
    `W`{.variable}\'s [document-tree child navigable target name
    property
    set](#document-tree-child-navigable-target-name-property-set){#windowproxy-getownproperty:document-tree-child-navigable-target-name-property-set},
    then:

    1.  Let `value`{.variable} be the [active
        `WindowProxy`](document-sequences.html#nav-wp){#windowproxy-getownproperty:nav-wp-2}
        of the [named
        object](#dom-window-nameditem-filter){#windowproxy-getownproperty:dom-window-nameditem-filter}
        of `W`{.variable} with the name `P`{.variable}.

    2.  Return
        [PropertyDescriptor](https://tc39.es/ecma262/#sec-property-descriptor-specification-type){#windowproxy-getownproperty:propertydescriptor-2
        x-internal="propertydescriptor"}{ \[\[Value\]\]:
        `value`{.variable}, \[\[Enumerable\]\]: false, \[\[Writable\]\]:
        false, \[\[Configurable\]\]: true }.

        The reason the property descriptors are non-enumerable, despite
        this mismatching the same-origin behavior, is for compatibility
        with existing web content. See [issue
        #3183](https://github.com/whatwg/html/issues/3183) for details.

7.  Return ?
    [CrossOriginPropertyFallback](#crossoriginpropertyfallback-(-p-)){#windowproxy-getownproperty:crossoriginpropertyfallback-(-p-)}(`P`{.variable}).

##### [7.2.3.6]{.secno} \[\[DefineOwnProperty\]\] ( `P`{.variable}, `Desc`{.variable} )[](#windowproxy-defineownproperty){.self-link} {#windowproxy-defineownproperty}

1.  Let `W`{.variable} be the value of the
    [\[\[Window\]\]](#concept-windowproxy-window){#windowproxy-defineownproperty:concept-windowproxy-window}
    internal slot of **this**.

2.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-defineownproperty:isplatformobjectsameorigin-(-o-)}(`W`{.variable})
    is true, then:

    1.  If `P`{.variable} is an [array index property
        name](https://webidl.spec.whatwg.org/#dfn-array-index-property-name){#windowproxy-defineownproperty:array-index-property-name
        x-internal="array-index-property-name"}, return false.

    2.  Return ?
        [OrdinaryDefineOwnProperty](https://tc39.es/ecma262/#sec-ordinarydefineownproperty){#windowproxy-defineownproperty:ordinarydefineownproperty
        x-internal="ordinarydefineownproperty"}(`W`{.variable},
        `P`{.variable}, `Desc`{.variable}).

        This is a [willful
        violation](https://infra.spec.whatwg.org/#willful-violation){#windowproxy-defineownproperty:willful-violation
        x-internal="willful-violation"} of the JavaScript
        specification\'s [invariants of the essential internal
        methods](https://tc39.es/ecma262/#sec-invariants-of-the-essential-internal-methods){#windowproxy-defineownproperty:invariants-of-the-essential-internal-methods
        x-internal="invariants-of-the-essential-internal-methods"} to
        maintain compatibility with existing web content. See
        [tc39/ecma262 issue
        #672](https://github.com/tc39/ecma262/issues/672) for more
        information. [\[JAVASCRIPT\]](references.html#refsJAVASCRIPT)

3.  Throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#windowproxy-defineownproperty:securityerror
    x-internal="securityerror"}
    [`DOMException`{#windowproxy-defineownproperty:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

##### [7.2.3.7]{.secno} \[\[Get\]\] ( `P`{.variable}, `Receiver`{.variable} )[](#windowproxy-get){.self-link} {#windowproxy-get}

1.  Let `W`{.variable} be the value of the
    [\[\[Window\]\]](#concept-windowproxy-window){#windowproxy-get:concept-windowproxy-window}
    internal slot of **this**.

2.  [Check if an access between two browsing contexts should be
    reported](browsers.html#coop-check-access-report){#windowproxy-get:coop-check-access-report},
    given the [current global
    object](webappapis.html#current-global-object){#windowproxy-get:current-global-object}\'s
    [browsing context](#window-bc){#windowproxy-get:window-bc},
    `W`{.variable}\'s [browsing
    context](#window-bc){#windowproxy-get:window-bc-2}, `P`{.variable},
    and the [current settings
    object](webappapis.html#current-settings-object){#windowproxy-get:current-settings-object}.

3.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-get:isplatformobjectsameorigin-(-o-)}(`W`{.variable})
    is true, then return ?
    [OrdinaryGet](https://tc39.es/ecma262/#sec-ordinaryget){#windowproxy-get:ordinaryget
    x-internal="ordinaryget"}(**this**, `P`{.variable},
    `Receiver`{.variable}).

4.  Return ?
    [CrossOriginGet](#crossoriginget-(-o,-p,-receiver-)){#windowproxy-get:crossoriginget-(-o,-p,-receiver-)}(**this**,
    `P`{.variable}, `Receiver`{.variable}).

**this** is passed rather than `W`{.variable} as
[OrdinaryGet](https://tc39.es/ecma262/#sec-ordinaryget){#windowproxy-get:ordinaryget-2
x-internal="ordinaryget"} and
[CrossOriginGet](#crossoriginget-(-o,-p,-receiver-)){#windowproxy-get:crossoriginget-(-o,-p,-receiver-)-2}
will invoke the [\[\[GetOwnProperty\]\]](#windowproxy-getownproperty)
internal method.

##### [7.2.3.8]{.secno} \[\[Set\]\] ( `P`{.variable}, `V`{.variable}, `Receiver`{.variable} )[](#windowproxy-set){.self-link} {#windowproxy-set}

1.  Let `W`{.variable} be the value of the
    [\[\[Window\]\]](#concept-windowproxy-window){#windowproxy-set:concept-windowproxy-window}
    internal slot of **this**.

2.  [Check if an access between two browsing contexts should be
    reported](browsers.html#coop-check-access-report){#windowproxy-set:coop-check-access-report},
    given the [current global
    object](webappapis.html#current-global-object){#windowproxy-set:current-global-object}\'s
    [browsing
    context](document-sequences.html#browsing-context){#windowproxy-set:browsing-context},
    `W`{.variable}\'s [browsing
    context](document-sequences.html#browsing-context){#windowproxy-set:browsing-context-2},
    `P`{.variable}, and the [current settings
    object](webappapis.html#current-settings-object){#windowproxy-set:current-settings-object}.

3.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-set:isplatformobjectsameorigin-(-o-)}(`W`{.variable})
    is true, then:

    1.  If `P`{.variable} is an [array index property
        name](https://webidl.spec.whatwg.org/#dfn-array-index-property-name){#windowproxy-set:array-index-property-name
        x-internal="array-index-property-name"}, then return false.

    2.  Return ?
        [OrdinarySet](https://tc39.es/ecma262/#sec-ordinaryset){#windowproxy-set:ordinaryset
        x-internal="ordinaryset"}(`W`{.variable}, `P`{.variable},
        `V`{.variable}, `Receiver`{.variable}).

4.  Return ?
    [CrossOriginSet](#crossoriginset-(-o,-p,-v,-receiver-)){#windowproxy-set:crossoriginset-(-o,-p,-v,-receiver-)}(**this**,
    `P`{.variable}, `V`{.variable}, `Receiver`{.variable}).

    **this** is passed rather than `W`{.variable} as
    [CrossOriginSet](#crossoriginset-(-o,-p,-v,-receiver-)){#windowproxy-set:crossoriginset-(-o,-p,-v,-receiver-)-2}
    will invoke the
    [\[\[GetOwnProperty\]\]](#windowproxy-getownproperty) internal
    method.

##### [7.2.3.9]{.secno} \[\[Delete\]\] ( `P`{.variable} )[](#windowproxy-delete){.self-link} {#windowproxy-delete}

1.  Let `W`{.variable} be the value of the
    [\[\[Window\]\]](#concept-windowproxy-window){#windowproxy-delete:concept-windowproxy-window}
    internal slot of **this**.

2.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-delete:isplatformobjectsameorigin-(-o-)}(`W`{.variable})
    is true, then:

    1.  If `P`{.variable} is an [array index property
        name](https://webidl.spec.whatwg.org/#dfn-array-index-property-name){#windowproxy-delete:array-index-property-name
        x-internal="array-index-property-name"}, then:

        1.  Let `desc`{.variable} be !
            **this**.\[\[GetOwnProperty\]\](`P`{.variable}).

        2.  If `desc`{.variable} is undefined, then return true.

        3.  Return false.

    2.  Return ?
        [OrdinaryDelete](https://tc39.es/ecma262/#sec-ordinarydelete){#windowproxy-delete:ordinarydelete
        x-internal="ordinarydelete"}(`W`{.variable}, `P`{.variable}).

3.  Throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#windowproxy-delete:securityerror
    x-internal="securityerror"}
    [`DOMException`{#windowproxy-delete:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

##### [7.2.3.10]{.secno} \[\[OwnPropertyKeys\]\] ( )[](#windowproxy-ownpropertykeys){.self-link} {#windowproxy-ownpropertykeys}

1.  Let `W`{.variable} be the value of the
    [\[\[Window\]\]](#concept-windowproxy-window){#windowproxy-ownpropertykeys:concept-windowproxy-window}
    internal slot of **this**.

2.  Let `maxProperties`{.variable} be `W`{.variable}\'s [associated
    `Document`](#concept-document-window){#windowproxy-ownpropertykeys:concept-document-window}\'s
    [document-tree child
    navigables](document-sequences.html#document-tree-child-navigables){#windowproxy-ownpropertykeys:document-tree-child-navigables}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#windowproxy-ownpropertykeys:list-size
    x-internal="list-size"}.

3.  Let `keys`{.variable} be [the
    range](https://infra.spec.whatwg.org/#the-exclusive-range){#windowproxy-ownpropertykeys:exclusive-range
    x-internal="exclusive-range"} 0 to `maxProperties`{.variable},
    exclusive.

4.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#windowproxy-ownpropertykeys:isplatformobjectsameorigin-(-o-)}(`W`{.variable})
    is true, then return the concatenation of `keys`{.variable} and
    [OrdinaryOwnPropertyKeys](https://tc39.es/ecma262/#sec-ordinaryownpropertykeys){#windowproxy-ownpropertykeys:ordinaryownpropertykeys
    x-internal="ordinaryownpropertykeys"}(`W`{.variable}).

5.  Return the concatenation of `keys`{.variable} and !
    [CrossOriginOwnPropertyKeys](#crossoriginownpropertykeys-(-o-)){#windowproxy-ownpropertykeys:crossoriginownpropertykeys-(-o-)}(`W`{.variable}).

#### [7.2.4]{.secno} The [`Location`{#the-location-interface:location}](#location) interface[](#the-location-interface){.self-link}

::::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Document/location](https://developer.mozilla.org/en-US/docs/Web/API/Document/location "The Document.location read-only property returns a Location object, which contains information about the URL of the document and provides methods for changing that URL and loading another URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

:::: feature
[Location](https://developer.mozilla.org/en-US/docs/Web/API/Location "The Location interface represents the location (URL) of the object it is linked to. Changes done on it are reflected on the object it relates to. Both the Document and Window interface have such a linked Location, accessible via Document.location and Window.location respectively.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[Window/location](https://developer.mozilla.org/en-US/docs/Web/API/Window/location "The Window.location read-only property returns a Location object with information about the current location of the document.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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
:::::::::

Each [`Window`{#the-location-interface:window}](#window) object is
associated with a unique instance of a
[`Location`{#the-location-interface:location-2}](#location) object,
allocated when the [`Window`{#the-location-interface:window-2}](#window)
object is created.

The [`Location`{#the-location-interface:location-3}](#location) exotic
object is defined through a mishmash of IDL, invocation of JavaScript
internal methods post-creation, and overridden JavaScript internal
methods. Coupled with its scary security policy, please take extra care
while implementing this excrescence.

To create a [`Location`{#the-location-interface:location-4}](#location)
object, run these steps:

1.  Let `location`{.variable} be a new
    [`Location`{#the-location-interface:location-5}](#location)
    [platform
    object](https://webidl.spec.whatwg.org/#dfn-platform-object){#the-location-interface:platform-object
    x-internal="platform-object"}.

2.  Let `valueOf`{.variable} be `location`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#the-location-interface:concept-relevant-realm}.\[\[Intrinsics\]\].\[\[[%Object.prototype.valueOf%](https://tc39.es/ecma262/#sec-object.prototype.valueof){#the-location-interface:object.prototype.valueof
    x-internal="object.prototype.valueof"}\]\].

3.  Perform !
    `location`{.variable}.\[\[DefineOwnProperty\]\](\"`valueOf`\", {
    \[\[Value\]\]: `valueOf`{.variable}, \[\[Writable\]\]: false,
    \[\[Enumerable\]\]: false, \[\[Configurable\]\]: false }).

4.  Perform !
    `location`{.variable}.\[\[DefineOwnProperty\]\]([%Symbol.toPrimitive%](infrastructure.html#symbol.toprimitive){#the-location-interface:symbol.toprimitive},
    { \[\[Value\]\]: undefined, \[\[Writable\]\]: false,
    \[\[Enumerable\]\]: false, \[\[Configurable\]\]: false }).

5.  Set the value of the
    [\[\[DefaultProperties\]\]](#defaultproperties){#the-location-interface:defaultproperties}
    internal slot of `location`{.variable} to
    `location`{.variable}.\[\[OwnPropertyKeys\]\]().

6.  Return `location`{.variable}.

The addition of `valueOf` and
[%Symbol.toPrimitive%](infrastructure.html#symbol.toprimitive){#the-location-interface:symbol.toprimitive-2}
own data properties, as well as the fact that all of
[`Location`{#the-location-interface:location-6}](#location)\'s IDL
attributes are marked
`[`[`LegacyUnforgeable`](https://webidl.spec.whatwg.org/#LegacyUnforgeable){#the-location-interface:legacyunforgeable
x-internal="legacyunforgeable"}`]`, is required by legacy code that
consulted the
[`Location`{#the-location-interface:location-7}](#location) interface,
or stringified it, to determine the [document
URL](https://dom.spec.whatwg.org/#concept-document-url){#the-location-interface:the-document's-address
x-internal="the-document's-address"}, and then used it in a
security-sensitive way. In particular, the `valueOf`,
[%Symbol.toPrimitive%](infrastructure.html#symbol.toprimitive){#the-location-interface:symbol.toprimitive-3},
and
`[`[`LegacyUnforgeable`](https://webidl.spec.whatwg.org/#LegacyUnforgeable){#the-location-interface:legacyunforgeable-2
x-internal="legacyunforgeable"}`]` stringifier mitigations ensure that
code such as `foo[location] = bar` or `location + ""` cannot be
misdirected.

`document`{.variable}`.`[`location`](#dom-document-location){#dom-document-location-dev}` [ = ``value`{.variable}` ]`\
`window`{.variable}`.`[`location`](#dom-location){#dom-location-dev}` [ = ``value`{.variable}` ]`

:   Returns a
    [`Location`{#the-location-interface:location-8}](#location) object
    with the current page\'s location.

    Can be set, to navigate to another page.

The [`Document`{#the-location-interface:document}](dom.html#document)
object\'s [`location`]{#dom-document-location .dfn dfn-for="Document"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this
x-internal="this"}\'s [relevant global
object](webappapis.html#concept-relevant-global){#the-location-interface:concept-relevant-global}\'s
[`Location`{#the-location-interface:location-9}](#location) object, if
[this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-2
x-internal="this"} is [fully
active](document-sequences.html#fully-active){#the-location-interface:fully-active},
and null otherwise.

The [`Window`{#the-location-interface:window-3}](#window) object\'s
[`location`]{#dom-location .dfn dfn-for="Window" dfn-type="attribute"}
getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-3
x-internal="this"}\'s
[`Location`{#the-location-interface:location-10}](#location) object.

[`Location`{#the-location-interface:location-11}](#location) objects
provide a representation of the
[URL](https://dom.spec.whatwg.org/#concept-document-url){#the-location-interface:the-document's-address-2
x-internal="the-document's-address"} of their associated
[`Document`{#the-location-interface:document-2}](dom.html#document), as
well as methods for
[navigating](browsing-the-web.html#navigate){#the-location-interface:navigate}
and
[reloading](browsing-the-web.html#reload){#the-location-interface:reload}
the associated
[navigable](document-sequences.html#navigable){#the-location-interface:navigable}.

``` idl
[Exposed=Window]
interface Location { // but see also additional creation steps and overridden internal methods
  [LegacyUnforgeable] stringifier attribute USVString href;
  [LegacyUnforgeable] readonly attribute USVString origin;
  [LegacyUnforgeable] attribute USVString protocol;
  [LegacyUnforgeable] attribute USVString host;
  [LegacyUnforgeable] attribute USVString hostname;
  [LegacyUnforgeable] attribute USVString port;
  [LegacyUnforgeable] attribute USVString pathname;
  [LegacyUnforgeable] attribute USVString search;
  [LegacyUnforgeable] attribute USVString hash;

  [LegacyUnforgeable] undefined assign(USVString url);
  [LegacyUnforgeable] undefined replace(USVString url);
  [LegacyUnforgeable] undefined reload();

  [LegacyUnforgeable, SameObject] readonly attribute DOMStringList ancestorOrigins;
};
```

`location`{.variable}`.``toString()`

`location`{.variable}`.`[`href`](#dom-location-href){#dom-location-href-dev}

::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/href](https://developer.mozilla.org/en-US/docs/Web/API/Location/href "The href property of the Location interface is a stringifier that returns a string containing the whole URL, and allows the href to be updated.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::

:::: feature
[Location/toString](https://developer.mozilla.org/en-US/docs/Web/API/Location/toString "The toString() stringifier method of the Location interface returns a string containing the whole URL. It is a read-only version of Location.href.")

Support in all current engines.

::: support
[Firefox22+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome52+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the [`Location`{#the-location-interface:location-12}](#location)
object\'s URL.

Can be set, to navigate to the given URL.

`location`{.variable}`.`[`origin`](#dom-location-origin){#dom-location-origin-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/origin](https://developer.mozilla.org/en-US/docs/Web/API/Location/origin "The origin read-only property of the Location interface is a string containing the Unicode serialization of the origin of the represented URL.")

Support in all current engines.

::: support
[Firefox21+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android3+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-13}](#location)
object\'s URL\'s origin.

`location`{.variable}`.`[`protocol`](#dom-location-protocol){#dom-location-protocol-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/protocol](https://developer.mozilla.org/en-US/docs/Web/API/Location/protocol "The protocol property of the Location interface is a string representing the protocol scheme of the URL, including the final ':'.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-14}](#location)
object\'s URL\'s scheme.

Can be set, to navigate to the same URL with a changed scheme.

`location`{.variable}`.`[`host`](#dom-location-host){#dom-location-host-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/host](https://developer.mozilla.org/en-US/docs/Web/API/Location/host "The host property of the Location interface is a string containing the host, that is the hostname, and then, if the port of the URL is nonempty, a ':', and the port of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-15}](#location)
object\'s URL\'s host and port (if different from the default port for
the scheme).

Can be set, to navigate to the same URL with a changed host and port.

`location`{.variable}`.`[`hostname`](#dom-location-hostname){#dom-location-hostname-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/hostname](https://developer.mozilla.org/en-US/docs/Web/API/Location/hostname "The hostname property of the Location interface is a string containing the domain of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-16}](#location)
object\'s URL\'s host.

Can be set, to navigate to the same URL with a changed host.

`location`{.variable}`.`[`port`](#dom-location-port){#dom-location-port-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/port](https://developer.mozilla.org/en-US/docs/Web/API/Location/port "The port property of the Location interface is a string containing the port number of the URL. If the URL does not contain an explicit port number, it will be set to ''.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-17}](#location)
object\'s URL\'s port.

Can be set, to navigate to the same URL with a changed port.

`location`{.variable}`.`[`pathname`](#dom-location-pathname){#dom-location-pathname-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/pathname](https://developer.mozilla.org/en-US/docs/Web/API/Location/pathname "The pathname property of the Location interface is a string containing the path of the URL for the location. If there is no path, pathname will be empty: otherwise, pathname contains an initial '/' followed by the path of the URL, not including the query string or fragment.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-18}](#location)
object\'s URL\'s path.

Can be set, to navigate to the same URL with a changed path.

`location`{.variable}`.`[`search`](#dom-location-search){#dom-location-search-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/search](https://developer.mozilla.org/en-US/docs/Web/API/Location/search "The search property of the Location interface is a search string, also called a query string; that is, a string containing a '?' followed by the parameters of the URL.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-19}](#location)
object\'s URL\'s query (includes leading \"`?`\" if non-empty).

Can be set, to navigate to the same URL with a changed query (ignores
leading \"`?`\").

`location`{.variable}`.`[`hash`](#dom-location-hash){#dom-location-hash-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/hash](https://developer.mozilla.org/en-US/docs/Web/API/Location/hash "The hash property of the Location interface returns a string containing a '#' followed by the fragment identifier of the URL — the ID on the page that the URL is trying to target.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer3+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the [`Location`{#the-location-interface:location-20}](#location)
object\'s URL\'s fragment (includes leading \"`#`\" if non-empty).

Can be set, to navigate to the same URL with a changed fragment (ignores
leading \"`#`\").

`location`{.variable}`.`[`assign`](#dom-location-assign){#dom-location-assign-dev}`(``url`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/assign](https://developer.mozilla.org/en-US/docs/Web/API/Location/assign "The Location.assign() method causes the window to load and display the document at the URL specified. After the navigation occurs, the user can navigate back to the page that called Location.assign() by pressing the "back" button.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari3+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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

Navigates to the given URL.

`location`{.variable}`.`[`replace`](#dom-location-replace){#dom-location-replace-dev}`(``url`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/replace](https://developer.mozilla.org/en-US/docs/Web/API/Location/replace "The replace() method of the Location interface replaces the current resource with the one at the provided URL. The difference from the assign() method is that after using replace() the current page will not be saved in session History, meaning the user won't be able to use the back button to navigate to it.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Removes the current page from the session history and navigates to the
given URL.

`location`{.variable}`.`[`reload`](#dom-location-reload){#dom-location-reload-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Location/reload](https://developer.mozilla.org/en-US/docs/Web/API/Location/reload "The location.reload() method reloads the current URL, like the Refresh button.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer5.5+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::
:::::

Reloads the current page.

`location`{.variable}`.`[`ancestorOrigins`](#dom-location-ancestororigins){#dom-location-ancestororigins-dev}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Location/ancestorOrigins](https://developer.mozilla.org/en-US/docs/Web/API/Location/ancestorOrigins "The ancestorOrigins read-only property of the Location interface is a static DOMStringList containing, in reverse order, the origins of all ancestor browsing contexts of the document associated with the given Location object.")

::: support
[FirefoxNo]{.firefox .no}[Safari6+]{.safari .yes}[Chrome20+]{.chrome
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

Returns a
[`DOMStringList`{#the-location-interface:domstringlist-2}](common-dom-interfaces.html#domstringlist)
object listing the origins of the [ancestor
navigables](document-sequences.html#ancestor-navigables){#the-location-interface:ancestor-navigables}\'
[active
documents](document-sequences.html#nav-document){#the-location-interface:nav-document}.

A [`Location`{#the-location-interface:location-21}](#location) object
has an associated [relevant `Document`]{#relevant-document .dfn}, which
is its [relevant global
object](webappapis.html#concept-relevant-global){#the-location-interface:concept-relevant-global-2}\'s
[browsing context](#window-bc){#the-location-interface:window-bc}\'s
[active
document](document-sequences.html#active-document){#the-location-interface:active-document},
if this [`Location`{#the-location-interface:location-22}](#location)
object\'s [relevant global
object](webappapis.html#concept-relevant-global){#the-location-interface:concept-relevant-global-3}\'s
[browsing context](#window-bc){#the-location-interface:window-bc-2} is
non-null, and null otherwise.

A [`Location`{#the-location-interface:location-23}](#location) object
has an associated [url]{#concept-location-url .dfn}, which is this
[`Location`{#the-location-interface:location-24}](#location) object\'s
[relevant
`Document`](#relevant-document){#the-location-interface:relevant-document}\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#the-location-interface:the-document's-address-3
x-internal="the-document's-address"}, if this
[`Location`{#the-location-interface:location-25}](#location) object\'s
[relevant
`Document`](#relevant-document){#the-location-interface:relevant-document-2}
is non-null, and
[`about:blank`{#the-location-interface:about:blank}](infrastructure.html#about:blank)
otherwise.

A [`Location`{#the-location-interface:location-26}](#location) object
has an associated [ancestor origins
list]{#concept-location-ancestor-origins-list .dfn}. When a
[`Location`{#the-location-interface:location-27}](#location) object is
created, its [ancestor origins
list](#concept-location-ancestor-origins-list){#the-location-interface:concept-location-ancestor-origins-list}
must be set to a
[`DOMStringList`{#the-location-interface:domstringlist-3}](common-dom-interfaces.html#domstringlist)
object whose associated list is the
[list](https://infra.spec.whatwg.org/#list){#the-location-interface:list
x-internal="list"} of strings that the following steps would produce:

1.  Let `output`{.variable} be a new
    [list](https://infra.spec.whatwg.org/#list){#the-location-interface:list-2
    x-internal="list"} of strings.

2.  Let `current`{.variable} be the
    [`Location`{#the-location-interface:location-28}](#location)
    object\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-3}.

3.  While `current`{.variable}\'s [container
    document](document-sequences.html#doc-container-document){#the-location-interface:doc-container-document}
    is non-null:

    1.  Set `current`{.variable} to `current`{.variable}\'s [container
        document](document-sequences.html#doc-container-document){#the-location-interface:doc-container-document-2}.

    2.  [Append](https://infra.spec.whatwg.org/#list-append){#the-location-interface:list-append
        x-internal="list-append"} the
        [serialization](browsers.html#ascii-serialisation-of-an-origin){#the-location-interface:ascii-serialisation-of-an-origin}
        of `current`{.variable}\'s
        [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin
        x-internal="concept-document-origin"} to `output`{.variable}.

4.  Return `output`{.variable}.

To [`Location`-object navigate]{#location-object-navigate .dfn} a
[`Location`{#the-location-interface:location-29}](#location) object
`location`{.variable} to a
[URL](https://url.spec.whatwg.org/#concept-url){#the-location-interface:url
x-internal="url"} `url`{.variable}, optionally given a
[`NavigationHistoryBehavior`{#the-location-interface:navigationhistorybehavior}](#navigationhistorybehavior)
`historyHandling`{.variable} (default
\"[`auto`{#the-location-interface:navigationhistorybehavior-auto}](browsing-the-web.html#navigationhistorybehavior-auto)\"):

1.  Let `navigable`{.variable} be `location`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#the-location-interface:concept-relevant-global-4}\'s
    [navigable](#window-navigable){#the-location-interface:window-navigable}.

2.  Let `sourceDocument`{.variable} be the [incumbent global
    object](webappapis.html#concept-incumbent-global){#the-location-interface:concept-incumbent-global}\'s
    [associated
    `Document`](#concept-document-window){#the-location-interface:concept-document-window}.

3.  If `location`{.variable}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-4}
    is not yet [completely
    loaded](document-lifecycle.html#completely-loaded){#the-location-interface:completely-loaded},
    and the [incumbent global
    object](webappapis.html#concept-incumbent-global){#the-location-interface:concept-incumbent-global-2}
    does not have [transient
    activation](interaction.html#transient-activation){#the-location-interface:transient-activation},
    then set `historyHandling`{.variable} to
    \"[`replace`{#the-location-interface:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

4.  [Navigate](browsing-the-web.html#navigate){#the-location-interface:navigate-2}
    `navigable`{.variable} to `url`{.variable} using
    `sourceDocument`{.variable}, with
    *[exceptionsEnabled](browsing-the-web.html#exceptions-enabled){#the-location-interface:exceptions-enabled}*
    set to true and
    *[historyHandling](browsing-the-web.html#navigation-hh)* set to
    `historyHandling`{.variable}.

The [`href`]{#dom-location-href .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-4
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-5}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-2
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-5
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#the-location-interface:concept-url-serializer
    x-internal="concept-url-serializer"}.

The
[`href`{#the-location-interface:dom-location-href-2}](#dom-location-href)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-6
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-6}
    is null, then return.

2.  Let `url`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#the-location-interface:encoding-parsing-a-url}
    given the given value, relative to the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-2}.

3.  If `url`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-location-interface:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#the-location-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-7
    x-internal="this"} to `url`{.variable}.

The
[`href`{#the-location-interface:dom-location-href-3}](#dom-location-href)
setter intentionally has no security check.

The [`origin`]{#dom-location-origin .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-8
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-7}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-3
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-2}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-3}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-2},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-2
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return the
    [serialization](browsers.html#ascii-serialisation-of-an-origin){#the-location-interface:ascii-serialisation-of-an-origin-2}
    of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-9
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-2}\'s
    [origin](https://url.spec.whatwg.org/#concept-url-origin){#the-location-interface:concept-url-origin
    x-internal="concept-url-origin"}.

The [`protocol`]{#dom-location-protocol .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-10
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-8}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-4
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-3}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-4}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-3},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-3
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-11
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-3}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#the-location-interface:concept-url-scheme
    x-internal="concept-url-scheme"}, followed by \"`:`\".

The
[`protocol`{#the-location-interface:dom-location-protocol-2}](#dom-location-protocol)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-12
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-9}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-13
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-10}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-5
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-4}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-5}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-4},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-4
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `copyURL`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-14
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-4}.

4.  Let `possibleFailure`{.variable} be the result of [basic URL
    parsing](https://url.spec.whatwg.org/#concept-basic-url-parser){#the-location-interface:basic-url-parser
    x-internal="basic-url-parser"} the given value, followed by \"`:`\",
    with `copyURL`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#the-location-interface:basic-url-parser-url
    x-internal="basic-url-parser-url"} and [scheme start
    state](https://url.spec.whatwg.org/#scheme-start-state){#the-location-interface:scheme-start-state
    x-internal="scheme-start-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#the-location-interface:basic-url-parser-state-override
    x-internal="basic-url-parser-state-override"}.

    Because the URL parser ignores multiple consecutive colons,
    providing a value of \"`https:`\" (or even \"`https::::`\") is the
    same as providing a value of \"`https`\".

5.  If `possibleFailure`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-location-interface:syntaxerror-2
    x-internal="syntaxerror"}
    [`DOMException`{#the-location-interface:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  If `copyURL`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#the-location-interface:concept-url-scheme-2
    x-internal="concept-url-scheme"} is not an [HTTP(S)
    scheme](https://fetch.spec.whatwg.org/#http-scheme){#the-location-interface:http(s)-scheme
    x-internal="http(s)-scheme"}, then terminate these steps.

7.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-2}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-15
    x-internal="this"} to `copyURL`{.variable}.

The [`host`]{#dom-location-host .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-16
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-11}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-6
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-5}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-6}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-5},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-5
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-17
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-5}.

3.  If `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#the-location-interface:concept-url-host
    x-internal="concept-url-host"} is null, return the empty string.

4.  If `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#the-location-interface:concept-url-port
    x-internal="concept-url-port"} is null, return `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#the-location-interface:concept-url-host-2
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#the-location-interface:host-serializer
    x-internal="host-serializer"}.

5.  Return `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#the-location-interface:concept-url-host-3
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#the-location-interface:host-serializer-2
    x-internal="host-serializer"}, followed by \"`:`\" and
    `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#the-location-interface:concept-url-port-2
    x-internal="concept-url-port"},
    [serialized](https://url.spec.whatwg.org/#serialize-an-integer){#the-location-interface:serialize-an-integer
    x-internal="serialize-an-integer"}.

The
[`host`{#the-location-interface:dom-location-host-2}](#dom-location-host)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-18
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-12}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-19
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-13}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-7
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-6}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-7}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-6},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-6
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `copyURL`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-20
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-6}.

4.  If `copyURL`{.variable} has an [opaque
    path](https://url.spec.whatwg.org/#url-opaque-path){#the-location-interface:opaque-path
    x-internal="opaque-path"}, then return.

5.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#the-location-interface:basic-url-parser-2
    x-internal="basic-url-parser"} the given value, with
    `copyURL`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#the-location-interface:basic-url-parser-url-2
    x-internal="basic-url-parser-url"} and [host
    state](https://url.spec.whatwg.org/#host-state){#the-location-interface:host-state
    x-internal="host-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#the-location-interface:basic-url-parser-state-override-2
    x-internal="basic-url-parser-state-override"}.

6.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-3}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-21
    x-internal="this"} to `copyURL`{.variable}.

The [`hostname`]{#dom-location-hostname .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-22
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-14}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-8
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-7}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-8}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-7},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-7
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-23
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-7}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#the-location-interface:concept-url-host-4
    x-internal="concept-url-host"} is null, return the empty string.

3.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-24
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-8}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#the-location-interface:concept-url-host-5
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#the-location-interface:host-serializer-3
    x-internal="host-serializer"}.

The
[`hostname`{#the-location-interface:dom-location-hostname-2}](#dom-location-hostname)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-25
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-15}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-26
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-16}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-9
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-8}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-9}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-8},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-8
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `copyURL`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-27
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-9}.

4.  If `copyURL`{.variable} has an [opaque
    path](https://url.spec.whatwg.org/#url-opaque-path){#the-location-interface:opaque-path-2
    x-internal="opaque-path"}, then return.

5.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#the-location-interface:basic-url-parser-3
    x-internal="basic-url-parser"} the given value, with
    `copyURL`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#the-location-interface:basic-url-parser-url-3
    x-internal="basic-url-parser-url"} and [hostname
    state](https://url.spec.whatwg.org/#hostname-state){#the-location-interface:hostname-state
    x-internal="hostname-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#the-location-interface:basic-url-parser-state-override-3
    x-internal="basic-url-parser-state-override"}.

6.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-4}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-28
    x-internal="this"} to `copyURL`{.variable}.

The [`port`]{#dom-location-port .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-29
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-17}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-10
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-9}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-10}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-9},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-9
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-11}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-30
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-10}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#the-location-interface:concept-url-port-3
    x-internal="concept-url-port"} is null, return the empty string.

3.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-31
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-11}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#the-location-interface:concept-url-port-4
    x-internal="concept-url-port"},
    [serialized](https://url.spec.whatwg.org/#serialize-an-integer){#the-location-interface:serialize-an-integer-2
    x-internal="serialize-an-integer"}.

The
[`port`{#the-location-interface:dom-location-port-2}](#dom-location-port)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-32
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-18}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-33
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-19}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-11
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-10}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-11}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-10},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-10
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-12}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `copyURL`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-34
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-12}.

4.  If `copyURL`{.variable} [cannot have a
    username/password/port](https://url.spec.whatwg.org/#cannot-have-a-username-password-port){#the-location-interface:cannot-have-a-username/password/port
    x-internal="cannot-have-a-username/password/port"}, then return.

5.  If the given value is the empty string, then set
    `copyURL`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#the-location-interface:concept-url-port-5
    x-internal="concept-url-port"} to null.

6.  Otherwise, [basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#the-location-interface:basic-url-parser-4
    x-internal="basic-url-parser"} the given value, with
    `copyURL`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#the-location-interface:basic-url-parser-url-4
    x-internal="basic-url-parser-url"} and [port
    state](https://url.spec.whatwg.org/#port-state){#the-location-interface:port-state
    x-internal="port-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#the-location-interface:basic-url-parser-state-override-4
    x-internal="basic-url-parser-state-override"}.

7.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-5}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-35
    x-internal="this"} to `copyURL`{.variable}.

The [`pathname`]{#dom-location-pathname .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-36
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-20}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-12
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-11}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-12}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-11},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-11
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-13}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return the result of [URL path
    serializing](https://url.spec.whatwg.org/#url-path-serializer){#the-location-interface:url-path-serializer
    x-internal="url-path-serializer"} this
    [`Location`{#the-location-interface:location-30}](#location)
    object\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-13}.

The
[`pathname`{#the-location-interface:dom-location-pathname-2}](#dom-location-pathname)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-37
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-21}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-38
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-22}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-13
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-12}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-13}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-12},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-12
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-14}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `copyURL`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-39
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-14}.

4.  If `copyURL`{.variable} has an [opaque
    path](https://url.spec.whatwg.org/#url-opaque-path){#the-location-interface:opaque-path-3
    x-internal="opaque-path"}, then return.

5.  Set `copyURL`{.variable}\'s
    [path](https://url.spec.whatwg.org/#concept-url-path){#the-location-interface:concept-url-path
    x-internal="concept-url-path"} to the empty list.

6.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#the-location-interface:basic-url-parser-5
    x-internal="basic-url-parser"} the given value, with
    `copyURL`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#the-location-interface:basic-url-parser-url-5
    x-internal="basic-url-parser-url"} and [path start
    state](https://url.spec.whatwg.org/#path-start-state){#the-location-interface:path-start-state
    x-internal="path-start-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#the-location-interface:basic-url-parser-state-override-5
    x-internal="basic-url-parser-state-override"}.

7.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-6}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-40
    x-internal="this"} to `copyURL`{.variable}.

The [`search`]{#dom-location-search .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-41
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-23}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-14
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-13}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-14}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-13},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-13
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-15}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-42
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-15}\'s
    [query](https://url.spec.whatwg.org/#concept-url-query){#the-location-interface:concept-url-query
    x-internal="concept-url-query"} is either null or the empty string,
    return the empty string.

3.  Return \"`?`\", followed by
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-43
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-16}\'s
    [query](https://url.spec.whatwg.org/#concept-url-query){#the-location-interface:concept-url-query-2
    x-internal="concept-url-query"}.

The
[`search`{#the-location-interface:dom-location-search-2}](#dom-location-search)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-44
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-24}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-45
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-25}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-15
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-14}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-15}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-14},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-14
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-16}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `copyURL`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-46
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-17}.

4.  If the given value is the empty string, set `copyURL`{.variable}\'s
    [query](https://url.spec.whatwg.org/#concept-url-query){#the-location-interface:concept-url-query-3
    x-internal="concept-url-query"} to null.

5.  Otherwise, run these substeps:

    1.  Let `input`{.variable} be the given value with a single leading
        \"`?`\" removed, if any.

    2.  Set `copyURL`{.variable}\'s
        [query](https://url.spec.whatwg.org/#concept-url-query){#the-location-interface:concept-url-query-4
        x-internal="concept-url-query"} to the empty string.

    3.  [Basic URL
        parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#the-location-interface:basic-url-parser-6
        x-internal="basic-url-parser"} `input`{.variable}, with null,
        the [relevant
        `Document`](#relevant-document){#the-location-interface:relevant-document-26}\'s
        [document\'s character
        encoding](https://dom.spec.whatwg.org/#concept-document-encoding){#the-location-interface:document's-character-encoding
        x-internal="document's-character-encoding"},
        `copyURL`{.variable} as
        [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#the-location-interface:basic-url-parser-url-6
        x-internal="basic-url-parser-url"}, and [query
        state](https://url.spec.whatwg.org/#query-state){#the-location-interface:query-state
        x-internal="query-state"} as [*state
        override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#the-location-interface:basic-url-parser-state-override-6
        x-internal="basic-url-parser-state-override"}.

6.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-7}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-47
    x-internal="this"} to `copyURL`{.variable}.

The [`hash`]{#dom-location-hash .dfn dfn-for="Location"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-48
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-27}
    is non-null and its
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-16
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-15}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-16}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-15},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-15
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-17}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-49
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-18}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-location-interface:concept-url-fragment
    x-internal="concept-url-fragment"} is either null or the empty
    string, return the empty string.

3.  Return \"`#`\", followed by
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-50
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-19}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-location-interface:concept-url-fragment-2
    x-internal="concept-url-fragment"}.

The
[`hash`{#the-location-interface:dom-location-hash-2}](#dom-location-hash)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-51
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-28}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-52
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-29}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-17
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-16}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-17}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-16},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-16
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-18}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `copyURL`{.variable} be a copy of
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-53
    x-internal="this"}\'s
    [url](#concept-location-url){#the-location-interface:concept-location-url-20}.

4.  Let `thisURLFragment`{.variable} be `copyURL`{.variable}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-location-interface:concept-url-fragment-3
    x-internal="concept-url-fragment"} if it is non-null; otherwise the
    empty string.

5.  Let `input`{.variable} be the given value with a single leading
    \"`#`\" removed, if any.

6.  Set `copyURL`{.variable}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-location-interface:concept-url-fragment-4
    x-internal="concept-url-fragment"} to the empty string.

7.  [Basic URL
    parse](https://url.spec.whatwg.org/#concept-basic-url-parser){#the-location-interface:basic-url-parser-7
    x-internal="basic-url-parser"} `input`{.variable}, with
    `copyURL`{.variable} as
    [*url*](https://url.spec.whatwg.org/#basic-url-parser-url){#the-location-interface:basic-url-parser-url-7
    x-internal="basic-url-parser-url"} and [fragment
    state](https://url.spec.whatwg.org/#fragment-state){#the-location-interface:fragment-state
    x-internal="fragment-state"} as [*state
    override*](https://url.spec.whatwg.org/#basic-url-parser-state-override){#the-location-interface:basic-url-parser-state-override-7
    x-internal="basic-url-parser-state-override"}.

8.  If `copyURL`{.variable}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-location-interface:concept-url-fragment-5
    x-internal="concept-url-fragment"} is `thisURLFragment`{.variable},
    then return.

    This bailout is necessary for compatibility with deployed content,
    which [redundantly sets `location.hash` on
    scroll](https://bugzilla.mozilla.org/show_bug.cgi?id=1733797#c2). It
    does not apply to other mechanisms of fragment navigation, such as
    the
    [`location.href`{#the-location-interface:dom-location-href-4}](#dom-location-href)
    setter or
    [`location.assign()`{#the-location-interface:dom-location-assign-2}](#dom-location-assign).

9.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-8}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-54
    x-internal="this"} to `copyURL`{.variable}.

Unlike the equivalent API for the
[`a`{#the-location-interface:the-a-element}](text-level-semantics.html#the-a-element)
and
[`area`{#the-location-interface:the-area-element}](image-maps.html#the-area-element)
elements, the
[`hash`{#the-location-interface:dom-location-hash-3}](#dom-location-hash)
setter does not special case the empty string, to remain compatible with
deployed scripts.

------------------------------------------------------------------------

The [`assign(``url`{.variable}`)`]{#dom-location-assign .dfn
dfn-for="Location" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-55
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-30}
    is null, then return.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-56
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-31}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-18
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-17}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-18}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-17},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-17
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-19}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `urlRecord`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#the-location-interface:encoding-parsing-a-url-2}
    given `url`{.variable}, relative to the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-19}.

4.  If `urlRecord`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-location-interface:syntaxerror-3
    x-internal="syntaxerror"}
    [`DOMException`{#the-location-interface:domexception-20}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-9}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-57
    x-internal="this"} to `urlRecord`{.variable}.

The [`replace(``url`{.variable}`)`]{#dom-location-replace .dfn
dfn-for="Location" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-58
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-32}
    is null, then return.

2.  Let `urlRecord`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#the-location-interface:encoding-parsing-a-url-3}
    given `url`{.variable}, relative to the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-20}.

3.  If `urlRecord`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#the-location-interface:syntaxerror-4
    x-internal="syntaxerror"}
    [`DOMException`{#the-location-interface:domexception-21}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  [`Location`-object
    navigate](#location-object-navigate){#the-location-interface:location-object-navigate-10}
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-59
    x-internal="this"} to `urlRecord`{.variable} given
    \"[`replace`{#the-location-interface:navigationhistorybehavior-replace-2}](browsing-the-web.html#navigationhistorybehavior-replace)\".

The
[`replace()`{#the-location-interface:dom-location-replace-2}](#dom-location-replace)
method intentionally has no security check.

The [`reload()`]{#dom-location-reload .dfn dfn-for="Location"
dfn-type="method"} method steps are:

1.  Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-60
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-33}.

2.  If `document`{.variable} is null, then return.

3.  If `document`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-19
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-18}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-21}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-18},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-18
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-22}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  [Reload](browsing-the-web.html#reload){#the-location-interface:reload-2}
    `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-location-interface:node-navigable}.

------------------------------------------------------------------------

The [`ancestorOrigins`]{#dom-location-ancestororigins .dfn
dfn-for="Location" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-61
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-34}
    is null, then return an empty
    [list](https://infra.spec.whatwg.org/#list){#the-location-interface:list-3
    x-internal="list"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-62
    x-internal="this"}\'s [relevant
    `Document`](#relevant-document){#the-location-interface:relevant-document-35}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-location-interface:concept-document-origin-20
    x-internal="concept-document-origin"} is not [same
    origin-domain](browsers.html#same-origin-domain){#the-location-interface:same-origin-domain-19}
    with the [entry settings
    object](webappapis.html#entry-settings-object){#the-location-interface:entry-settings-object-22}\'s
    [origin](webappapis.html#concept-settings-object-origin){#the-location-interface:concept-settings-object-origin-19},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-location-interface:securityerror-19
    x-internal="securityerror"}
    [`DOMException`{#the-location-interface:domexception-23}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Otherwise, return
    [this](https://webidl.spec.whatwg.org/#this){#the-location-interface:this-63
    x-internal="this"}\'s [ancestor origins
    list](#concept-location-ancestor-origins-list){#the-location-interface:concept-location-ancestor-origins-list-2}.

The details of how the
[`ancestorOrigins`{#the-location-interface:dom-location-ancestororigins-2}](#dom-location-ancestororigins)
attribute works are still controversial and might change. See [issue
#1918](https://github.com/whatwg/html/issues/1918) for more information.

------------------------------------------------------------------------

As explained earlier, the
[`Location`{#the-location-interface:location-31}](#location) exotic
object requires additional logic beyond IDL for security purposes. The
[`Location`{#the-location-interface:location-32}](#location) object must
use the ordinary internal methods except where it is explicitly
specified otherwise below.

Also, every [`Location`{#the-location-interface:location-33}](#location)
object has a [\[\[DefaultProperties\]\]]{#defaultproperties .dfn}
internal slot representing its own properties at time of its creation.

##### [7.2.4.1]{.secno} \[\[GetPrototypeOf\]\] ( )[](#location-getprototypeof){.self-link} {#location-getprototypeof}

1.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#location-getprototypeof:isplatformobjectsameorigin-(-o-)}(**this**)
    is true, then return !
    [OrdinaryGetPrototypeOf](https://tc39.es/ecma262/#sec-ordinarygetprototypeof){#location-getprototypeof:ordinarygetprototypeof
    x-internal="ordinarygetprototypeof"}(**this**).

2.  Return null.

##### [7.2.4.2]{.secno} \[\[SetPrototypeOf\]\] ( `V`{.variable} )[](#location-setprototypeof){.self-link} {#location-setprototypeof}

1.  Return !
    [SetImmutablePrototype](https://tc39.es/ecma262/#sec-set-immutable-prototype){#location-setprototypeof:setimmutableprototype
    x-internal="setimmutableprototype"}(**this**, `V`{.variable}).

##### [7.2.4.3]{.secno} \[\[IsExtensible\]\] ( )[](#location-isextensible){.self-link} {#location-isextensible}

1.  Return true.

##### [7.2.4.4]{.secno} \[\[PreventExtensions\]\] ( )[](#location-preventextensions){.self-link} {#location-preventextensions}

1.  Return false.

##### [7.2.4.5]{.secno} \[\[GetOwnProperty\]\] ( `P`{.variable} )[](#location-getownproperty){.self-link} {#location-getownproperty}

1.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#location-getownproperty:isplatformobjectsameorigin-(-o-)}(**this**)
    is true, then:

    1.  Let `desc`{.variable} be
        [OrdinaryGetOwnProperty](https://tc39.es/ecma262/#sec-ordinarygetownproperty){#location-getownproperty:ordinarygetownproperty
        x-internal="ordinarygetownproperty"}(**this**, `P`{.variable}).

    2.  If the value of the
        [\[\[DefaultProperties\]\]](#defaultproperties){#location-getownproperty:defaultproperties}
        internal slot of **this** contains `P`{.variable}, then set
        `desc`{.variable}.\[\[Configurable\]\] to true.

    3.  Return `desc`{.variable}.

2.  Let `property`{.variable} be
    [CrossOriginGetOwnPropertyHelper](#crossorigingetownpropertyhelper-(-o,-p-)){#location-getownproperty:crossorigingetownpropertyhelper-(-o,-p-)}(**this**,
    `P`{.variable}).

3.  If `property`{.variable} is not undefined, then return
    `property`{.variable}.

4.  Return ?
    [CrossOriginPropertyFallback](#crossoriginpropertyfallback-(-p-)){#location-getownproperty:crossoriginpropertyfallback-(-p-)}(`P`{.variable}).

##### [7.2.4.6]{.secno} \[\[DefineOwnProperty\]\] ( `P`{.variable}, `Desc`{.variable} )[](#location-defineownproperty){.self-link} {#location-defineownproperty}

1.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#location-defineownproperty:isplatformobjectsameorigin-(-o-)}(**this**)
    is true, then:

    1.  If the value of the
        [\[\[DefaultProperties\]\]](#defaultproperties){#location-defineownproperty:defaultproperties}
        internal slot of **this** contains `P`{.variable}, then return
        false.

    2.  Return ?
        [OrdinaryDefineOwnProperty](https://tc39.es/ecma262/#sec-ordinarydefineownproperty){#location-defineownproperty:ordinarydefineownproperty
        x-internal="ordinarydefineownproperty"}(**this**,
        `P`{.variable}, `Desc`{.variable}).

2.  Throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#location-defineownproperty:securityerror
    x-internal="securityerror"}
    [`DOMException`{#location-defineownproperty:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

##### [7.2.4.7]{.secno} \[\[Get\]\] ( `P`{.variable}, `Receiver`{.variable} )[](#location-get){.self-link} {#location-get}

1.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#location-get:isplatformobjectsameorigin-(-o-)}(**this**)
    is true, then return ?
    [OrdinaryGet](https://tc39.es/ecma262/#sec-ordinaryget){#location-get:ordinaryget
    x-internal="ordinaryget"}(**this**, `P`{.variable},
    `Receiver`{.variable}).

2.  Return ?
    [CrossOriginGet](#crossoriginget-(-o,-p,-receiver-)){#location-get:crossoriginget-(-o,-p,-receiver-)}(**this**,
    `P`{.variable}, `Receiver`{.variable}).

##### [7.2.4.8]{.secno} \[\[Set\]\] ( `P`{.variable}, `V`{.variable}, `Receiver`{.variable} )[](#location-set){.self-link} {#location-set}

1.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#location-set:isplatformobjectsameorigin-(-o-)}(**this**)
    is true, then return ?
    [OrdinarySet](https://tc39.es/ecma262/#sec-ordinaryset){#location-set:ordinaryset
    x-internal="ordinaryset"}(**this**, `P`{.variable}, `V`{.variable},
    `Receiver`{.variable}).

2.  Return ?
    [CrossOriginSet](#crossoriginset-(-o,-p,-v,-receiver-)){#location-set:crossoriginset-(-o,-p,-v,-receiver-)}(**this**,
    `P`{.variable}, `V`{.variable}, `Receiver`{.variable}).

##### [7.2.4.9]{.secno} \[\[Delete\]\] ( `P`{.variable} )[](#location-delete){.self-link} {#location-delete}

1.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#location-delete:isplatformobjectsameorigin-(-o-)}(**this**)
    is true, then return ?
    [OrdinaryDelete](https://tc39.es/ecma262/#sec-ordinarydelete){#location-delete:ordinarydelete
    x-internal="ordinarydelete"}(**this**, `P`{.variable}).

2.  Throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#location-delete:securityerror
    x-internal="securityerror"}
    [`DOMException`{#location-delete:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

##### [7.2.4.10]{.secno} \[\[OwnPropertyKeys\]\] ( )[](#location-ownpropertykeys){.self-link} {#location-ownpropertykeys}

1.  If
    [IsPlatformObjectSameOrigin](#isplatformobjectsameorigin-(-o-)){#location-ownpropertykeys:isplatformobjectsameorigin-(-o-)}(**this**)
    is true, then return
    [OrdinaryOwnPropertyKeys](https://tc39.es/ecma262/#sec-ordinaryownpropertykeys){#location-ownpropertykeys:ordinaryownpropertykeys
    x-internal="ordinaryownpropertykeys"}(**this**).

2.  Return
    [CrossOriginOwnPropertyKeys](#crossoriginownpropertykeys-(-o-)){#location-ownpropertykeys:crossoriginownpropertykeys-(-o-)}(**this**).

#### [7.2.5]{.secno} The [`History`{#the-history-interface:history-3}](#history-3) interface[](#the-history-interface){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[History](https://developer.mozilla.org/en-US/docs/Web/API/History "The History interface allows manipulation of the browser session history, that is the pages visited in the tab or frame that the current page is loaded in.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android10.1+]{.opera_android .yes}
:::
::::

:::: feature
[Window/history](https://developer.mozilla.org/en-US/docs/Web/API/Window/history "The Window.history read-only property returns a reference to the History object, which provides an interface for manipulating the browser session history (pages visited in the tab or frame that the current page is loaded in).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera3+]{.opera .yes}[Edge79+]{.edge_blink .yes}

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
:::::::

``` idl
enum ScrollRestoration { "auto", "manual" };

[Exposed=Window]
interface History {
  readonly attribute unsigned long length;
  attribute ScrollRestoration scrollRestoration;
  readonly attribute any state;
  undefined go(optional long delta = 0);
  undefined back();
  undefined forward();
  undefined pushState(any data, DOMString unused, optional USVString? url = null);
  undefined replaceState(any data, DOMString unused, optional USVString? url = null);
};
```

[`history`](#dom-history){#dom-history-dev}`.`[`length`](#dom-history-length){#dom-history-length-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/length](https://developer.mozilla.org/en-US/docs/Web/API/History/length "The History.length read-only property returns an integer representing the number of elements in the session history, including the currently loaded page.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

Returns the number of overall [session history
entries](document-sequences.html#tn-session-history-entries){#the-history-interface:tn-session-history-entries}
for the current [traversable
navigable](document-sequences.html#traversable-navigable){#the-history-interface:traversable-navigable}.

[`history`](#dom-history){#the-history-interface:dom-history}`.`[`scrollRestoration`](#dom-history-scroll-restoration){#dom-history-scroll-restoration-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/scrollRestoration](https://developer.mozilla.org/en-US/docs/Web/API/History/scrollRestoration "The scrollRestoration property of History interface allows web applications to explicitly set default scroll restoration behavior on history navigation.")

Support in all current engines.

::: support
[Firefox46+]{.firefox .yes}[Safari11+]{.safari .yes}[Chrome46+]{.chrome
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

Returns the [scroll restoration
mode](browsing-the-web.html#she-scroll-restoration-mode){#the-history-interface:she-scroll-restoration-mode}
of the [active session history
entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry}.

[`history`](#dom-history){#the-history-interface:dom-history-2}`.`[`scrollRestoration`](#dom-history-scroll-restoration){#the-history-interface:dom-history-scroll-restoration-2}` = ``value`{.variable}

Set the [scroll restoration
mode](browsing-the-web.html#she-scroll-restoration-mode){#the-history-interface:she-scroll-restoration-mode-2}
of the [active session history
entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry-2}
to `value`{.variable}.

[`history`](#dom-history){#the-history-interface:dom-history-3}`.`[`state`](#dom-history-state){#dom-history-state-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/state](https://developer.mozilla.org/en-US/docs/Web/API/History/state "The History.state property returns a value representing the state at the top of the history stack. This is a way to look at the state without having to wait for a popstate event.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome19+]{.chrome
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

Returns the [classic history API
state](browsing-the-web.html#she-classic-history-api-state){#the-history-interface:she-classic-history-api-state}
of the [active session history
entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry-3},
deserialized into a JavaScript value.

[`history`](#dom-history){#the-history-interface:dom-history-4}`.`[`go`](#dom-history-go){#the-history-interface:dom-history-go-2}`()`

Reloads the current page.

[`history`](#dom-history){#the-history-interface:dom-history-5}`.`[`go`](#dom-history-go){#dom-history-go-dev}`(``delta`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/go](https://developer.mozilla.org/en-US/docs/Web/API/History/go "The History.go() method loads a specific page from the session history. You can use it to move forwards and backwards through the history depending on the value of a parameter.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

Goes back or forward the specified number of steps in the overall
[session history
entries](document-sequences.html#tn-session-history-entries){#the-history-interface:tn-session-history-entries-2}
list for the current [traversable
navigable](document-sequences.html#traversable-navigable){#the-history-interface:traversable-navigable-2}.

A zero delta will reload the current page.

If the delta is out of range, does nothing.

[`history`](#dom-history){#the-history-interface:dom-history-6}`.`[`back`](#dom-history-back){#dom-history-back-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/back](https://developer.mozilla.org/en-US/docs/Web/API/History/back "The History.back() method causes the browser to move back one page in the session history.")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

Goes back one step in the overall [session history
entries](document-sequences.html#tn-session-history-entries){#the-history-interface:tn-session-history-entries-3}
list for the current [traversable
navigable](document-sequences.html#traversable-navigable){#the-history-interface:traversable-navigable-3}.

If there is no previous page, does nothing.

[`history`](#dom-history){#the-history-interface:dom-history-7}`.`[`forward`](#dom-history-forward){#dom-history-forward-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/forward](https://developer.mozilla.org/en-US/docs/Web/API/History/forward "The History.forward() method causes the browser to move forward one page in the session history. It has the same effect as calling history.go(1).")

Support in all current engines.

::: support
[Firefox1+]{.firefox .yes}[Safari1+]{.safari .yes}[Chrome1+]{.chrome
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

Goes forward one step in the overall [session history
entries](document-sequences.html#tn-session-history-entries){#the-history-interface:tn-session-history-entries-4}
list for the current [traversable
navigable](document-sequences.html#traversable-navigable){#the-history-interface:traversable-navigable-4}.

If there is no next page, does nothing.

[`history`](#dom-history){#the-history-interface:dom-history-8}`.`[`pushState`](#dom-history-pushstate){#dom-history-pushstate-dev}`(``data`{.variable}`, "")`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/pushState](https://developer.mozilla.org/en-US/docs/Web/API/History/pushState "In an HTML document, the history.pushState() method adds an entry to the browser's session history stack.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.5+]{.opera_android .yes}
:::
::::
:::::

Adds a new entry into session history with its [classic history API
state](browsing-the-web.html#she-classic-history-api-state){#the-history-interface:she-classic-history-api-state-2}
set to a serialization of `data`{.variable}. The [active history
entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry-4}\'s
[URL](browsing-the-web.html#she-url){#the-history-interface:she-url}
will be copied over and used for the new entry\'s URL.

(The second parameter exists for historical reasons, and cannot be
omitted; passing the empty string is traditional.)

[`history`](#dom-history){#the-history-interface:dom-history-9}`.`[`pushState`](#dom-history-pushstate){#the-history-interface:dom-history-pushstate-2}`(``data`{.variable}`, "", ``url`{.variable}`)`

Adds a new entry into session history with its [classic history API
state](browsing-the-web.html#she-classic-history-api-state){#the-history-interface:she-classic-history-api-state-3}
set to a serialization of `data`{.variable}, and with its
[URL](browsing-the-web.html#she-url){#the-history-interface:she-url-2}
set to `url`{.variable}.

If the current
[`Document`{#the-history-interface:document}](dom.html#document) [cannot
have its URL
rewritten](#can-have-its-url-rewritten){#the-history-interface:can-have-its-url-rewritten}
to `url`{.variable}, a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror
x-internal="securityerror"}
[`DOMException`{#the-history-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
will be thrown.

(The second parameter exists for historical reasons, and cannot be
omitted; passing the empty string is traditional.)

[`history`](#dom-history){#the-history-interface:dom-history-10}`.`[`replaceState`](#dom-history-replacestate){#dom-history-replacestate-dev}`(``data`{.variable}`, "")`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[History/replaceState](https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState "The History.replaceState() method modifies the current history entry, replacing it with the state object and URL passed in the method parameters. This method is particularly useful when you want to update the state object or URL of the current history entry in response to some user action.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11.5+]{.opera_android .yes}
:::
::::
:::::

Updates the [classic history API
state](browsing-the-web.html#she-classic-history-api-state){#the-history-interface:she-classic-history-api-state-4}
of the [active session history
entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry-5}
to a structured clone of `data`{.variable}.

(The second parameter exists for historical reasons, and cannot be
omitted; passing the empty string is traditional.)

[`history`](#dom-history){#the-history-interface:dom-history-11}`.`[`replaceState`](#dom-history-replacestate){#the-history-interface:dom-history-replacestate-2}`(``data`{.variable}`, "", ``url`{.variable}`)`

Updates the [classic history API
state](browsing-the-web.html#she-classic-history-api-state){#the-history-interface:she-classic-history-api-state-5}
of the [active session history
entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry-6}
to a structured clone of `data`{.variable}, and its
[URL](browsing-the-web.html#she-url){#the-history-interface:she-url-3}
to `url`{.variable}.

If the current
[`Document`{#the-history-interface:document-2}](dom.html#document)
[cannot have its URL
rewritten](#can-have-its-url-rewritten){#the-history-interface:can-have-its-url-rewritten-2}
to `url`{.variable}, a
[\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-2
x-internal="securityerror"}
[`DOMException`{#the-history-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
will be thrown.

(The second parameter exists for historical reasons, and cannot be
omitted; passing the empty string is traditional.)

A [`Document`{#the-history-interface:document-3}](dom.html#document) has
a [history object]{#doc-history .dfn}, a
[`History`{#the-history-interface:history-3-2}](#history-3) object.

The [`history`]{#dom-history .dfn dfn-for="Window" dfn-type="attribute"}
getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this
x-internal="this"}\'s [associated
`Document`](#concept-document-window){#the-history-interface:concept-document-window}\'s
[history object](#doc-history){#the-history-interface:doc-history}.

------------------------------------------------------------------------

Each [`History`{#the-history-interface:history-3-3}](#history-3) object
has [state]{#concept-history-state .dfn}, initially null.

Each [`History`{#the-history-interface:history-3-4}](#history-3) object
has a [length]{#concept-history-length .dfn}, a non-negative integer,
initially 0.

Each [`History`{#the-history-interface:history-3-5}](#history-3) object
has an [index]{#concept-history-index .dfn}, a non-negative integer,
initially 0.

Although the
[index](#concept-history-index){#the-history-interface:concept-history-index}
is not directly exposed, it can be inferred from changes to the
[length](#concept-history-length){#the-history-interface:concept-history-length}
during synchronous navigations. In fact, that is what it\'s used for.

The [`length`]{#dom-history-length .dfn dfn-for="History"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global}\'s
    [associated
    `Document`](#concept-document-window){#the-history-interface:concept-document-window-2}
    is not [fully
    active](document-sequences.html#fully-active){#the-history-interface:fully-active},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-3
    x-internal="securityerror"}
    [`DOMException`{#the-history-interface:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-3
    x-internal="this"}\'s
    [length](#concept-history-length){#the-history-interface:concept-history-length-2}.

The [`scrollRestoration`]{#dom-history-scroll-restoration .dfn
dfn-for="History" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-4
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-2}\'s
    [associated
    `Document`](#concept-document-window){#the-history-interface:concept-document-window-3}
    is not [fully
    active](document-sequences.html#fully-active){#the-history-interface:fully-active-2},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-4
    x-internal="securityerror"}
    [`DOMException`{#the-history-interface:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-5
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-3}\'s
    [navigable](#window-navigable){#the-history-interface:window-navigable}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry-7}\'s
    [scroll restoration
    mode](browsing-the-web.html#she-scroll-restoration-mode){#the-history-interface:she-scroll-restoration-mode-3}.

The
[`scrollRestoration`{#the-history-interface:dom-history-scroll-restoration-3}](#dom-history-scroll-restoration)
setter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-6
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-4}\'s
    [associated
    `Document`](#concept-document-window){#the-history-interface:concept-document-window-4}
    is not [fully
    active](document-sequences.html#fully-active){#the-history-interface:fully-active-3},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-5
    x-internal="securityerror"}
    [`DOMException`{#the-history-interface:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-7
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-5}\'s
    [navigable](#window-navigable){#the-history-interface:window-navigable-2}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#the-history-interface:nav-active-history-entry-8}\'s
    [scroll restoration
    mode](browsing-the-web.html#she-scroll-restoration-mode){#the-history-interface:she-scroll-restoration-mode-4}
    to the given value.

The [`state`]{#dom-history-state .dfn dfn-for="History"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-8
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-6}\'s
    [associated
    `Document`](#concept-document-window){#the-history-interface:concept-document-window-5}
    is not [fully
    active](document-sequences.html#fully-active){#the-history-interface:fully-active-4},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-6
    x-internal="securityerror"}
    [`DOMException`{#the-history-interface:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-9
    x-internal="this"}\'s
    [state](#concept-history-state){#the-history-interface:concept-history-state}.

The [`go(``delta`{.variable}`)`]{#dom-history-go .dfn dfn-for="History"
dfn-type="method"} method steps are to [delta
traverse](#delta-traverse){#the-history-interface:delta-traverse}
[this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-10
x-internal="this"} given `delta`{.variable}.

The [`back()`]{#dom-history-back .dfn dfn-for="History"
dfn-type="method"} method steps are to [delta
traverse](#delta-traverse){#the-history-interface:delta-traverse-2}
[this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-11
x-internal="this"} given −1.

The [`forward()`]{#dom-history-forward .dfn dfn-for="History"
dfn-type="method"} method steps are to [delta
traverse](#delta-traverse){#the-history-interface:delta-traverse-3}
[this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-12
x-internal="this"} given +1.

To [delta traverse]{#delta-traverse .dfn} a
[`History`{#the-history-interface:history-3-6}](#history-3) object
`history`{.variable} given an integer `delta`{.variable}:

1.  Let `document`{.variable} be `history`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-7}\'s
    [associated
    `Document`](#concept-document-window){#the-history-interface:concept-document-window-6}.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#the-history-interface:fully-active-5},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-7
    x-internal="securityerror"}
    [`DOMException`{#the-history-interface:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `delta`{.variable} is 0, then
    [reload](browsing-the-web.html#reload){#the-history-interface:reload}
    `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-history-interface:node-navigable},
    and return.

4.  [Traverse the history by a
    delta](browsing-the-web.html#traverse-the-history-by-a-delta){#the-history-interface:traverse-the-history-by-a-delta}
    given `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#the-history-interface:node-navigable-2}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#the-history-interface:nav-traversable},
    `delta`{.variable}, and with
    *[sourceDocument](browsing-the-web.html#traverse-sourcedocument)*
    set to `document`{.variable}.

The
[`pushState(``data`{.variable}`, ``unused`{.variable}`, ``url`{.variable}`)`]{#dom-history-pushstate
.dfn dfn-for="History" dfn-type="method"} method steps are to run the
[shared history push/replace state
steps](#shared-history-push/replace-state-steps){#the-history-interface:shared-history-push/replace-state-steps}
given
[this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-13
x-internal="this"}, `data`{.variable}, `url`{.variable}, and
\"[`push`{#the-history-interface:navigationhistorybehavior-push}](browsing-the-web.html#navigationhistorybehavior-push)\".

The
[`replaceState(``data`{.variable}`, ``unused`{.variable}`, ``url`{.variable}`)`]{#dom-history-replacestate
.dfn dfn-for="History" dfn-type="method"} method steps are to run the
[shared history push/replace state
steps](#shared-history-push/replace-state-steps){#the-history-interface:shared-history-push/replace-state-steps-2}
given
[this](https://webidl.spec.whatwg.org/#this){#the-history-interface:this-14
x-internal="this"}, `data`{.variable}, `url`{.variable}, and
\"[`replace`{#the-history-interface:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\".

The [shared history push/replace state
steps]{#shared-history-push/replace-state-steps .dfn}, given a
[`History`{#the-history-interface:history-3-7}](#history-3)
`history`{.variable}, a value `data`{.variable}, a [scalar value
string](https://infra.spec.whatwg.org/#scalar-value-string){#the-history-interface:scalar-value-string
x-internal="scalar-value-string"}-or-null `url`{.variable}, and a
[history handling
behavior](browsing-the-web.html#history-handling-behavior){#the-history-interface:history-handling-behavior}
`historyHandling`{.variable}, are:

1.  Let `document`{.variable} be `history`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-8}\'s
    [associated
    `Document`](#concept-document-window){#the-history-interface:concept-document-window-7}.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#the-history-interface:fully-active-6},
    then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-8
    x-internal="securityerror"}
    [`DOMException`{#the-history-interface:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Optionally, throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-9
    x-internal="securityerror"}
    [`DOMException`{#the-history-interface:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.
    (For example, the user agent might disallow calls to these methods
    that are invoked on a timer, or from event listeners that are not
    triggered in response to a clear user action, or that are invoked in
    rapid succession.)

4.  Let `serializedData`{.variable} be
    [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#the-history-interface:structuredserializeforstorage}(`data`{.variable}).
    Rethrow any exceptions.

5.  Let `newURL`{.variable} be `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-history-interface:the-document's-address
    x-internal="the-document's-address"}.

6.  If `url`{.variable} is not null or the empty string, then:

    1.  Set `newURL`{.variable} to the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#the-history-interface:encoding-parsing-a-url}
        given `url`{.variable}, relative to the [relevant settings
        object](webappapis.html#relevant-settings-object){#the-history-interface:relevant-settings-object}
        of `history`{.variable}.

    2.  If `newURL`{.variable} is failure, then throw a
        [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-10
        x-internal="securityerror"}
        [`DOMException`{#the-history-interface:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  If `document`{.variable} [cannot have its URL
        rewritten](#can-have-its-url-rewritten){#the-history-interface:can-have-its-url-rewritten-3}
        to `newURL`{.variable}, then throw a
        [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-history-interface:securityerror-11
        x-internal="securityerror"}
        [`DOMException`{#the-history-interface:domexception-11}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    The special case for the empty string here is historical, and leads
    to different resulting URLs when comparing code such as
    `location.href = ""` (which performs URL parsing on the empty
    string) versus `history.pushState(null, "", "")` (which bypasses
    it).

7.  Let `navigation`{.variable} be `history`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#the-history-interface:concept-relevant-global-9}\'s
    [navigation
    API](#window-navigation-api){#the-history-interface:window-navigation-api}.

8.  Let `continue`{.variable} be the result of [firing a
    push/replace/reload `navigate`
    event](#fire-a-push/replace/reload-navigate-event){#the-history-interface:fire-a-push/replace/reload-navigate-event}
    at `navigation`{.variable} with
    *[navigationType](#fire-navigate-prr-navigationtype)* set to
    `historyHandling`{.variable},
    *[isSameDocument](#fire-navigate-prr-issamedocument)* set to true,
    *[destinationURL](#fire-navigate-prr-destinationurl)* set to
    `newURL`{.variable}, and
    *[classicHistoryAPIState](#fire-navigate-prr-classichistoryapistate)*
    set to `serializedData`{.variable}.

9.  If `continue`{.variable} is false, then return.

10. Run the [URL and history update
    steps](browsing-the-web.html#url-and-history-update-steps){#the-history-interface:url-and-history-update-steps}
    given `document`{.variable} and `newURL`{.variable}, with
    *[serializedData](browsing-the-web.html#uhus-serializeddata)* set to
    `serializedData`{.variable} and
    *[historyHandling](browsing-the-web.html#uhus-historyhandling)* set
    to `historyHandling`{.variable}.

User agents may limit the number of state objects added to the session
history per page. If a page hits the
[implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-history-interface:implementation-defined
x-internal="implementation-defined"} limit, user agents must remove the
entry immediately after the first entry for that
[`Document`{#the-history-interface:document-4}](dom.html#document)
object in the session history after having added the new entry. (Thus
the state history acts as a FIFO buffer for eviction, but as a LIFO
buffer for navigation.)

A [`Document`{#the-history-interface:document-5}](dom.html#document)
`document`{.variable} [can have its URL
rewritten]{#can-have-its-url-rewritten .dfn} to a
[URL](https://url.spec.whatwg.org/#concept-url){#the-history-interface:url
x-internal="url"} `targetURL`{.variable} if the following algorithm
returns true:

1.  Let `documentURL`{.variable} be `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-history-interface:the-document's-address-2
    x-internal="the-document's-address"}.

2.  If `targetURL`{.variable} and `documentURL`{.variable} differ in
    their
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#the-history-interface:concept-url-scheme
    x-internal="concept-url-scheme"},
    [username](https://url.spec.whatwg.org/#concept-url-username){#the-history-interface:concept-url-username
    x-internal="concept-url-username"},
    [password](https://url.spec.whatwg.org/#concept-url-password){#the-history-interface:concept-url-password
    x-internal="concept-url-password"},
    [host](https://url.spec.whatwg.org/#concept-url-host){#the-history-interface:concept-url-host
    x-internal="concept-url-host"}, or
    [port](https://url.spec.whatwg.org/#concept-url-port){#the-history-interface:concept-url-port
    x-internal="concept-url-port"} components, then return false.

3.  If `targetURL`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#the-history-interface:concept-url-scheme-2
    x-internal="concept-url-scheme"} is an [HTTP(S)
    scheme](https://fetch.spec.whatwg.org/#http-scheme){#the-history-interface:http(s)-scheme
    x-internal="http(s)-scheme"}, then return true.

    Differences in
    [path](https://url.spec.whatwg.org/#concept-url-path){#the-history-interface:concept-url-path
    x-internal="concept-url-path"},
    [query](https://url.spec.whatwg.org/#concept-url-query){#the-history-interface:concept-url-query
    x-internal="concept-url-query"}, and
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-history-interface:concept-url-fragment
    x-internal="concept-url-fragment"} are allowed for
    [`http:`{#the-history-interface:http-protocol}](https://httpwg.org/specs/rfc7230.html#http.uri){x-internal="http-protocol"}
    and
    [`https:`{#the-history-interface:https-protocol}](https://httpwg.org/specs/rfc7230.html#https.uri){x-internal="https-protocol"}
    URLs.

4.  If `targetURL`{.variable}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#the-history-interface:concept-url-scheme-3
    x-internal="concept-url-scheme"} is \"`file`\", then:

    1.  If `targetURL`{.variable} and `documentURL`{.variable} differ in
        their
        [path](https://url.spec.whatwg.org/#concept-url-path){#the-history-interface:concept-url-path-2
        x-internal="concept-url-path"} component, then return false.

    2.  Return true.

    Differences in
    [query](https://url.spec.whatwg.org/#concept-url-query){#the-history-interface:concept-url-query-2
    x-internal="concept-url-query"} and
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-history-interface:concept-url-fragment-2
    x-internal="concept-url-fragment"} are allowed for `file:` URLs.

5.  If `targetURL`{.variable} and `documentURL`{.variable} differ in
    their
    [path](https://url.spec.whatwg.org/#concept-url-path){#the-history-interface:concept-url-path-3
    x-internal="concept-url-path"} component or
    [query](https://url.spec.whatwg.org/#concept-url-query){#the-history-interface:concept-url-query-3
    x-internal="concept-url-query"} components, then return false.

    Only differences in
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-history-interface:concept-url-fragment-3
    x-internal="concept-url-fragment"} are allowed for other types of
    URLs.

6.  Return true.

::: example
`document`{.variable}\'s
[URL](https://dom.spec.whatwg.org/#concept-document-url){#the-history-interface:the-document's-address-3
x-internal="the-document's-address"}

`targetURL`{.variable}

[can have its URL
rewritten](#can-have-its-url-rewritten){#can-have-its-url-rewritten-dev}

`https://example.com/home`

`https://example.com/home#about`

✅

`https://example.com/home`

`https://example.com/home?page=shop`

✅

`https://example.com/home`

`https://example.com/shop`

✅

`https://example.com/home`

`https://user:pass@example.com/home`

❌

`https://example.com/home`

`http://example.com/home`

❌

`file:///path/to/x`

`file:///path/to/x#hash`

✅

`file:///path/to/x`

`file:///path/to/x?search`

✅

`file:///path/to/x`

`file:///path/to/y`

❌

`about:blank`

`about:blank#hash`

✅

`about:blank`

`about:blank?search`

❌

`about:blank`

`about:srcdoc`

❌

`data:text/html,foo`

`data:text/html,foo#hash`

✅

`data:text/html,foo`

`data:text/html,foo?search`

❌

`data:text/html,foo`

`data:text/html,bar`

❌

`data:text/html,foo`

`data:bar`

❌

`blob:https://example.com/77becafe-657b-4fdc-8bd3-e83aaa5e8f43`

`blob:https://example.com/77becafe-657b-4fdc-8bd3-e83aaa5e8f43#hash`

✅

`blob:https://example.com/77becafe-657b-4fdc-8bd3-e83aaa5e8f43`

`blob:https://example.com/77becafe-657b-4fdc-8bd3-e83aaa5e8f43?search`

❌

`blob:https://example.com/77becafe-657b-4fdc-8bd3-e83aaa5e8f43`

`blob:https://example.com/anything`

❌

`blob:https://example.com/77becafe-657b-4fdc-8bd3-e83aaa5e8f43`

`blob:path`

❌

Note how only the
[URL](https://dom.spec.whatwg.org/#concept-document-url){#the-history-interface:the-document's-address-4
x-internal="the-document's-address"} of the
[`Document`{#the-history-interface:document-6}](dom.html#document)
matters, and not its
[origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-history-interface:concept-document-origin
x-internal="concept-document-origin"}. They can mismatch in cases like
[`about:blank`{#the-history-interface:about:blank}](infrastructure.html#about:blank)
[`Document`{#the-history-interface:document-7}](dom.html#document)s with
inherited origins, in sandboxed
[`iframe`{#the-history-interface:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s,
or when the
[`document.domain`{#the-history-interface:dom-document-domain}](browsers.html#dom-document-domain)
setter has been used.
:::

::: example
Consider a game where the user can navigate along a line, such that the
user is always at some coordinate, and such that the user can bookmark
the page corresponding to a particular coordinate, to return to it
later.

A static page implementing the x=5 position in such a game could look
like the following:

``` html
<!DOCTYPE HTML>
<!-- this is https://example.com/line?x=5 -->
<html lang="en">
<title>Line Game - 5</title>
<p>You are at coordinate 5 on the line.</p>
<p>
 <a href="?x=6">Advance to 6</a> or
 <a href="?x=4">retreat to 4</a>?
</p>
```

The problem with such a system is that each time the user clicks, the
whole page has to be reloaded. Here instead is another way of doing it,
using script:

``` html
<!DOCTYPE HTML>
<!-- this starts off as https://example.com/line?x=5 -->
<html lang="en">
<title>Line Game - 5</title>
<p>You are at coordinate <span id="coord">5</span> on the line.</p>
<p>
 <a href="?x=6" onclick="go(1); return false;">Advance to 6</a> or
 <a href="?x=4" onclick="go(-1); return false;">retreat to 4</a>?
</p>
<script>
 var currentPage = 5; // prefilled by server
 function go(d) {
   setupPage(currentPage + d);
   history.pushState(currentPage, "", '?x=' + currentPage);
 }
 onpopstate = function(event) {
   setupPage(event.state);
 }
 function setupPage(page) {
   currentPage = page;
   document.title = 'Line Game - ' + currentPage;
   document.getElementById('coord').textContent = currentPage;
   document.links[0].href = '?x=' + (currentPage+1);
   document.links[0].textContent = 'Advance to ' + (currentPage+1);
   document.links[1].href = '?x=' + (currentPage-1);
   document.links[1].textContent = 'retreat to ' + (currentPage-1);
 }
</script>
```

In systems without script, this still works like the previous example.
However, users that *do* have script support can now navigate much
faster, since there is no network access for the same experience.
Furthermore, contrary to the experience the user would have with just a
naïve script-based approach, bookmarking and navigating the session
history still work.

In the example above, the `data`{.variable} argument to the
[`pushState()`{#the-history-interface:dom-history-pushstate-3}](#dom-history-pushstate)
method is the same information as would be sent to the server, but in a
more convenient form, so that the script doesn\'t have to parse the URL
each time the user navigates.
:::

::: example
Most applications want to use the same [scroll restoration
mode](browsing-the-web.html#scroll-restoration-mode){#the-history-interface:scroll-restoration-mode}
value for all of their history entries. To achieve this they can set the
[`scrollRestoration`{#the-history-interface:dom-history-scroll-restoration-4}](#dom-history-scroll-restoration)
attribute as soon as possible (e.g., in the first
[`script`{#the-history-interface:the-script-element}](scripting.html#the-script-element)
element in the document\'s
[`head`{#the-history-interface:the-head-element}](semantics.html#the-head-element)
element) to ensure that any entry added to the history session gets the
desired scroll restoration mode.

``` html
<head>
  <script>
       if ('scrollRestoration' in history)
            history.scrollRestoration = 'manual';
  </script>
</head>
   
```
:::

#### [7.2.6]{.secno} The navigation API[](#navigation-api){.self-link} {#navigation-api}

##### [7.2.6.1]{.secno} Introduction[](#navigation-api-intro){.self-link} {#navigation-api-intro}

*This section is non-normative.*

The navigation API, provided by the global
[`navigation`{#navigation-api-intro:dom-navigation}](#dom-navigation)
property, provides a modern and web application-focused way of managing
navigations and history entries. It is a successor to the classic
[`location`{#navigation-api-intro:dom-location}](#dom-location) and
[`history`{#navigation-api-intro:dom-history}](#dom-history) APIs.

One ability the API provides is inspecting [session history
entries](browsing-the-web.html#session-history-entry){#navigation-api-intro:session-history-entry}.
For example, the following will display the entries\' URLs in an ordered
list:

``` js
const ol = document.createElement("ol");
ol.start = 0; // so that the list items' ordinal values match up with the entry indices

for (const entry of navigation.entries()) {
  const li = document.createElement("li");

  if (entry.index < navigation.currentEntry.index) {
    li.className = "backward";
  } else if (entry.index > navigation.currentEntry.index) {
    li.className = "forward";
  } else {
    li.className = "current";
  }

  li.textContent = entry.url;
  ol.append(li);
}
```

The
[`navigation.entries()`{#navigation-api-intro:dom-navigation-entries}](#dom-navigation-entries)
array contains
[`NavigationHistoryEntry`{#navigation-api-intro:navigationhistoryentry}](#navigationhistoryentry)
instances, which have other useful properties in addition to the
[`url`{#navigation-api-intro:dom-navigationhistoryentry-url}](#dom-navigationhistoryentry-url)
and
[`index`{#navigation-api-intro:dom-navigationhistoryentry-index}](#dom-navigationhistoryentry-index)
properties shown here. Note that the array only contains
[`NavigationHistoryEntry`{#navigation-api-intro:navigationhistoryentry-2}](#navigationhistoryentry)
objects that represent the current
[navigable](document-sequences.html#navigable){#navigation-api-intro:navigable},
and thus its contents are not impacted by navigations inside [navigable
containers](document-sequences.html#navigable-container){#navigation-api-intro:navigable-container}
such as
[`iframe`{#navigation-api-intro:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s,
or by navigations of the [parent
navigable](document-sequences.html#nav-parent){#navigation-api-intro:nav-parent}
in cases where the navigation API is itself being used inside an
[`iframe`{#navigation-api-intro:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element).
Additionally, it only contains
[`NavigationHistoryEntry`{#navigation-api-intro:navigationhistoryentry-3}](#navigationhistoryentry)
objects representing
same-[origin](browsers.html#concept-origin){#navigation-api-intro:concept-origin}
[session history
entries](browsing-the-web.html#session-history-entry){#navigation-api-intro:session-history-entry-2},
meaning that if the user has visited other origins before or after the
current one, there will not be corresponding
[`NavigationHistoryEntry`{#navigation-api-intro:navigationhistoryentry-4}](#navigationhistoryentry)s.

------------------------------------------------------------------------

The navigation API can also be used to navigate, reload, or traverse
through the history:

``` html
<button onclick="navigation.reload()">Reload</button>

<input type="url" id="navigationURL">
<button onclick="navigation.navigate(navigationURL.value)">Navigate</button>

<button id="backButton" onclick="navigation.back()">Back</button>
<button id="forwardButton" onclick="navigation.forward()">Forward</button>

<select id="traversalDestinations"></select>
<button id="goButton" onclick="navigation.traverseTo(traversalDestinations.value)">Traverse To</button>

<script>
backButton.disabled = !navigation.canGoBack;
forwardButton.disabled = !navigation.canGoForward;

for (const entry of navigation.entries()) {
  traversalDestinations.append(new Option(entry.url, entry.key));
}
</script>
```

Note that traversals are again limited to
same-[origin](browsers.html#concept-origin){#navigation-api-intro:concept-origin-2}
destinations, meaning that, for example,
[`navigation.canGoBack`{#navigation-api-intro:dom-navigation-cangoback}](#dom-navigation-cangoback)
will be false if the previous [session history
entry](browsing-the-web.html#session-history-entry){#navigation-api-intro:session-history-entry-3}
is for a page from another origin.

------------------------------------------------------------------------

The most powerful part of the navigation API is the
[`navigate`{#navigation-api-intro:event-navigate}](indices.html#event-navigate)
event, which fires whenever almost any navigation or traversal occurs in
the current
[navigable](document-sequences.html#navigable){#navigation-api-intro:navigable-2}:

``` js
navigation.onnavigate = event => {
  console.log(event.navigationType); // "push", "replace", "reload", or "traverse"
  console.log(event.destination.url);
  console.log(event.userInitiated);
  // ... and other useful properties
};
```

(The event will not fire for [location bar-initiated
navigations](document-lifecycle.html#nav-traversal-ui), or navigations
initiated from other windows, when the destination of the navigation is
a new document.)

Much of the time, the event\'s
[`cancelable`{#navigation-api-intro:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
property will be true, meaning this event can be canceled using
[`preventDefault()`{#navigation-api-intro:dom-event-preventdefault}](https://dom.spec.whatwg.org/#dom-event-preventdefault){x-internal="dom-event-preventdefault"}:

``` js
navigation.onnavigate = event => {
  if (event.cancelable && isDisallowedURL(event.destination.url)) {
    alert(`Please don't go to ${event.destination.url}!`);
    event.preventDefault();
  }
};
```

The
[`cancelable`{#navigation-api-intro:dom-event-cancelable-2}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
property will be false for some
\"[`traverse`{#navigation-api-intro:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\"
navigations, such as those taking place inside [child
navigables](document-sequences.html#child-navigable){#navigation-api-intro:child-navigable},
those crossing to new origins, or when the user attempts to traverse
again shortly after a previous call to
[`preventDefault()`{#navigation-api-intro:dom-event-preventdefault-2}](https://dom.spec.whatwg.org/#dom-event-preventdefault){x-internal="dom-event-preventdefault"}
prevented them from doing so.

The
[`NavigateEvent`{#navigation-api-intro:navigateevent}](#navigateevent)\'s
[`intercept()`{#navigation-api-intro:dom-navigateevent-intercept}](#dom-navigateevent-intercept)
method allows intercepting a navigation and converting it into a
same-document navigation:

``` js
navigation.addEventListener("navigate", e => {
  // Some navigations, e.g. cross-origin navigations, we cannot intercept.
  // Let the browser handle those normally.
  if (!e.canIntercept) {
    return;
  }

  // Similarly, don't intercept fragment navigations or downloads.
  if (e.hashChange || e.downloadRequest !== null) {
    return;
  }

  const url = new URL(event.destination.url);

  if (url.pathname.startsWith("/articles/")) {
    e.intercept({
      async handler() {
        // The URL has already changed, so show a placeholder while
        // fetching the new content, such as a spinner or loading page.
        renderArticlePagePlaceholder();

        // Fetch the new content and display when ready.
        const articleContent = await getArticleContent(url.pathname, { signal: e.signal });
        renderArticlePage(articleContent);
      }
    });
  }
});
```

Note that the
[`handler`{#navigation-api-intro:dom-navigationinterceptoptions-handler}](#dom-navigationinterceptoptions-handler)
function can return a promise to represent the asynchronous progress,
and success or failure, of the navigation. While the promise is still
pending, browser UI can treat the navigation as ongoing (e.g., by
presenting a loading spinner). Other parts of the navigation API are
also sensitive to these promises, such as the return value of
[`navigation.navigate()`{#navigation-api-intro:dom-navigation-navigate}](#dom-navigation-navigate):

``` js
const { committed, finished } = await navigation.navigate("/articles/the-navigation-api-is-cool");

// The committed promise will fulfill once the URL has changed, which happens
// immediately (as long as the NavigateEvent wasn't canceled).
await committed;

// The finished promise will fulfill once the Promise returned by handler() has
// fulfilled, which happens once the article is downloaded and rendered. (Or,
// it will reject, if handler() fails along the way).
await finished;
```

##### [7.2.6.2]{.secno} The [`Navigation`{#navigation-interface:navigation}](#navigation) interface[](#navigation-interface){.self-link} {#navigation-interface}

``` idl
[Exposed=Window]
interface Navigation : EventTarget {
  sequence<NavigationHistoryEntry> entries();
  readonly attribute NavigationHistoryEntry? currentEntry;
  undefined updateCurrentEntry(NavigationUpdateCurrentEntryOptions options);
  readonly attribute NavigationTransition? transition;
  readonly attribute NavigationActivation? activation;

  readonly attribute boolean canGoBack;
  readonly attribute boolean canGoForward;

  NavigationResult navigate(USVString url, optional NavigationNavigateOptions options = {});
  NavigationResult reload(optional NavigationReloadOptions options = {});

  NavigationResult traverseTo(DOMString key, optional NavigationOptions options = {});
  NavigationResult back(optional NavigationOptions options = {});
  NavigationResult forward(optional NavigationOptions options = {});

  attribute EventHandler onnavigate;
  attribute EventHandler onnavigatesuccess;
  attribute EventHandler onnavigateerror;
  attribute EventHandler oncurrententrychange;
};

dictionary NavigationUpdateCurrentEntryOptions {
  required any state;
};

dictionary NavigationOptions {
  any info;
};

dictionary NavigationNavigateOptions : NavigationOptions {
  any state;
  NavigationHistoryBehavior history = "auto";
};

dictionary NavigationReloadOptions : NavigationOptions {
  any state;
};

dictionary NavigationResult {
  Promise<NavigationHistoryEntry> committed;
  Promise<NavigationHistoryEntry> finished;
};

enum NavigationHistoryBehavior {
  "auto",
  "push",
  "replace"
};
```

Each [`Window`{#navigation-interface:window}](#window) has an associated
[navigation API]{#window-navigation-api .dfn}, which is a
[`Navigation`{#navigation-interface:navigation-2}](#navigation) object.
Upon creation of the [`Window`{#navigation-interface:window-2}](#window)
object, its [navigation
API](#window-navigation-api){#navigation-interface:window-navigation-api}
must be set to a
[new](https://webidl.spec.whatwg.org/#new){#navigation-interface:new
x-internal="new"}
[`Navigation`{#navigation-interface:navigation-3}](#navigation) object
created in the [`Window`{#navigation-interface:window-3}](#window)
object\'s [relevant
realm](webappapis.html#concept-relevant-realm){#navigation-interface:concept-relevant-realm}.

The [`navigation`]{#dom-navigation .dfn dfn-for="Window"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#navigation-interface:this
x-internal="this"}\'s [navigation
API](#window-navigation-api){#navigation-interface:window-navigation-api-2}.

The following are the [event
handlers](webappapis.html#event-handlers){#navigation-interface:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#navigation-interface:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#navigation-interface:event-handler-idl-attributes},
by all objects implementing the
[`Navigation`{#navigation-interface:navigation-4}](#navigation)
interface:

[Event
handler](webappapis.html#event-handlers){#navigation-interface:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#navigation-interface:event-handler-event-type-2}

[`onnavigate`]{#handler-navigation-onnavigate .dfn dfn-for="Navigation"
dfn-type="attribute"}

[`navigate`{#navigation-interface:event-navigate}](indices.html#event-navigate)

[`onnavigatesuccess`]{#handler-navigation-onnavigatesuccess .dfn
dfn-for="Navigation" dfn-type="attribute"}

[`navigatesuccess`{#navigation-interface:event-navigatesuccess}](indices.html#event-navigatesuccess)

[`onnavigateerror`]{#handler-navigation-onnavigateerror .dfn
dfn-for="Navigation" dfn-type="attribute"}

[`navigateerror`{#navigation-interface:event-navigateerror}](indices.html#event-navigateerror)

[`oncurrententrychange`]{#handler-navigation-oncurrententrychange .dfn
dfn-for="Navigation" dfn-type="attribute"}

[`currententrychange`{#navigation-interface:event-currententrychange}](indices.html#event-currententrychange)

##### [7.2.6.3]{.secno} Core infrastructure[](#navigation-api-core){.self-link} {#navigation-api-core}

Each [`Navigation`{#navigation-api-core:navigation}](#navigation) has an
associated [entry list]{#navigation-entry-list .dfn}, a
[list](https://infra.spec.whatwg.org/#list){#navigation-api-core:list
x-internal="list"} of
[`NavigationHistoryEntry`{#navigation-api-core:navigationhistoryentry}](#navigationhistoryentry)
objects, initially empty.

Each [`Navigation`{#navigation-api-core:navigation-2}](#navigation) has
an associated [current entry index]{#navigation-current-entry-index
.dfn}, an integer, initially −1.

The [current entry]{#navigation-current-entry .dfn} of a
[`Navigation`{#navigation-api-core:navigation-3}](#navigation)
`navigation`{.variable} is the result of running the following steps:

1.  If `navigation`{.variable} [has entries and events
    disabled](#has-entries-and-events-disabled){#navigation-api-core:has-entries-and-events-disabled},
    then return null.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#navigation-api-core:assert
    x-internal="assert"}: `navigation`{.variable}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-core:navigation-current-entry-index}
    is not −1.

3.  Return `navigation`{.variable}\'s [entry
    list](#navigation-entry-list){#navigation-api-core:navigation-entry-list}\[`navigation`{.variable}\'s
    [current entry
    index](#navigation-current-entry-index){#navigation-api-core:navigation-current-entry-index-2}\].

A [`Navigation`{#navigation-api-core:navigation-4}](#navigation)
`navigation`{.variable} [has entries and events
disabled]{#has-entries-and-events-disabled .dfn} if the following steps
return true:

1.  Let `document`{.variable} be `navigation`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#navigation-api-core:concept-relevant-global}\'s
    [associated
    `Document`](#concept-document-window){#navigation-api-core:concept-document-window}.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#navigation-api-core:fully-active},
    then return true.

3.  If `document`{.variable}\'s [is initial
    `about:blank`](dom.html#is-initial-about:blank){#navigation-api-core:is-initial-about:blank}
    is true, then return true.

4.  If `document`{.variable}\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#navigation-api-core:concept-document-origin
    x-internal="concept-document-origin"} is
    [opaque](browsers.html#concept-origin-opaque){#navigation-api-core:concept-origin-opaque},
    then return true.

5.  Return false.

To [get the navigation API entry
index]{#getting-the-navigation-api-entry-index .dfn} of a [session
history
entry](browsing-the-web.html#session-history-entry){#navigation-api-core:session-history-entry}
`she`{.variable} within a
[`Navigation`{#navigation-api-core:navigation-5}](#navigation)
`navigation`{.variable}:

1.  Let `index`{.variable} be 0.

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#navigation-api-core:list-iterate
    x-internal="list-iterate"} `nhe`{.variable} of
    `navigation`{.variable}\'s [entry
    list](#navigation-entry-list){#navigation-api-core:navigation-entry-list-2}:

    1.  If `nhe`{.variable}\'s [session history
        entry](#nhe-she){#navigation-api-core:nhe-she} is equal to
        `she`{.variable}, then return `index`{.variable}.

    2.  Increment `index`{.variable} by 1.

3.  Return −1.

------------------------------------------------------------------------

A key type used throughout the navigation API is the
[`NavigationType`{#navigation-api-core:navigationtype}](#navigationtype)
enumeration:

``` idl
enum NavigationType {
 "push",
 "replace",
 "reload",
 "traverse"
};
```

This captures the main web developer-visible types of \"navigations\",
which (as [noted
elsewhere](browsing-the-web.html#note-meaning-of-navigate)) do not
exactly correspond to this standard\'s singular
[navigate](browsing-the-web.html#navigate){#navigation-api-core:navigate}
algorithm. The meaning of each value is the following:

\"[`push`]{#dom-navigationtype-push .dfn dfn-for="NavigationType" dfn-type="enum-value"}\"
:   Corresponds to calls to
    [navigate](browsing-the-web.html#navigate){#navigation-api-core:navigate-2}
    where the [history handling
    behavior](browsing-the-web.html#history-handling-behavior){#navigation-api-core:history-handling-behavior}
    ends up as
    \"[`push`{#navigation-api-core:navigationhistorybehavior-push}](browsing-the-web.html#navigationhistorybehavior-push)\",
    or to
    [`history.pushState()`{#navigation-api-core:dom-history-pushstate}](#dom-history-pushstate).

\"[`replace`]{#dom-navigationtype-replace .dfn dfn-for="NavigationType" dfn-type="enum-value"}\"
:   Corresponds to calls to
    [navigate](browsing-the-web.html#navigate){#navigation-api-core:navigate-3}
    where the [history handling
    behavior](browsing-the-web.html#history-handling-behavior){#navigation-api-core:history-handling-behavior-2}
    ends up as
    \"[`replace`{#navigation-api-core:navigationhistorybehavior-replace}](browsing-the-web.html#navigationhistorybehavior-replace)\",
    or to
    [`history.replaceState()`{#navigation-api-core:dom-history-pushstate-2}](#dom-history-pushstate).

\"[`reload`]{#dom-navigationtype-reload .dfn dfn-for="NavigationType" dfn-type="enum-value"}\"
:   Corresponds to calls to
    [reload](browsing-the-web.html#reload){#navigation-api-core:reload}.

\"[`traverse`]{#dom-navigationtype-traverse .dfn dfn-for="NavigationType" dfn-type="enum-value"}\"
:   Corresponds to calls to [traverse the history by a
    delta](browsing-the-web.html#traverse-the-history-by-a-delta){#navigation-api-core:traverse-the-history-by-a-delta}.

The value space of the
[`NavigationType`{#navigation-api-core:navigationtype-2}](#navigationtype)
enumeration is a superset of the value space of the
specification-internal [history handling
behavior](browsing-the-web.html#history-handling-behavior){#navigation-api-core:history-handling-behavior-3}
type. Several parts of this standard make use of this overlap, by
passing in a [history handling
behavior](browsing-the-web.html#history-handling-behavior){#navigation-api-core:history-handling-behavior-4}
to an algorithm that expects a
[`NavigationType`{#navigation-api-core:navigationtype-3}](#navigationtype).

##### [7.2.6.4]{.secno} Initializing and updating the entry list[](#navigation-api-entry-updates){.self-link} {#navigation-api-entry-updates}

To [initialize the navigation API entries for a new
document]{#initialize-the-navigation-api-entries-for-a-new-document
.dfn} given a
[`Navigation`{#navigation-api-entry-updates:navigation}](#navigation)
`navigation`{.variable}, a
[list](https://infra.spec.whatwg.org/#list){#navigation-api-entry-updates:list
x-internal="list"} of [session history
entries](browsing-the-web.html#session-history-entry){#navigation-api-entry-updates:session-history-entry}
`newSHEs`{.variable}, and a [session history
entry](browsing-the-web.html#session-history-entry){#navigation-api-entry-updates:session-history-entry-2}
`initialSHE`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigation-api-entry-updates:assert
    x-internal="assert"}: `navigation`{.variable}\'s [entry
    list](#navigation-entry-list){#navigation-api-entry-updates:navigation-entry-list}
    [is
    empty](https://infra.spec.whatwg.org/#list-is-empty){#navigation-api-entry-updates:list-is-empty
    x-internal="list-is-empty"}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#navigation-api-entry-updates:assert-2
    x-internal="assert"}: `navigation`{.variable}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index}
    is −1.

3.  If `navigation`{.variable} [has entries and events
    disabled](#has-entries-and-events-disabled){#navigation-api-entry-updates:has-entries-and-events-disabled},
    then return.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#navigation-api-entry-updates:list-iterate
    x-internal="list-iterate"} `newSHE`{.variable} of
    `newSHEs`{.variable}:

    1.  Let `newNHE`{.variable} be a
        [new](https://webidl.spec.whatwg.org/#new){#navigation-api-entry-updates:new
        x-internal="new"}
        [`NavigationHistoryEntry`{#navigation-api-entry-updates:navigationhistoryentry}](#navigationhistoryentry)
        created in the [relevant
        realm](webappapis.html#concept-relevant-realm){#navigation-api-entry-updates:concept-relevant-realm}
        of `navigation`{.variable}.

    2.  Set `newNHE`{.variable}\'s [session history
        entry](#nhe-she){#navigation-api-entry-updates:nhe-she} to
        `newSHE`{.variable}.

    3.  [Append](https://infra.spec.whatwg.org/#list-append){#navigation-api-entry-updates:list-append
        x-internal="list-append"} `newNHE`{.variable} to
        `navigation`{.variable}\'s [entry
        list](#navigation-entry-list){#navigation-api-entry-updates:navigation-entry-list-2}.

    `newSHEs`{.variable} will have originally come from [getting session
    history entries for the navigation
    API](browsing-the-web.html#getting-session-history-entries-for-the-navigation-api){#navigation-api-entry-updates:getting-session-history-entries-for-the-navigation-api},
    and thus each `newSHE`{.variable} will be contiguous
    [same](browsers.html#same-origin){#navigation-api-entry-updates:same-origin}
    [origin](browsing-the-web.html#document-state-origin){#navigation-api-entry-updates:document-state-origin}
    with `initialSHE`{.variable}.

5.  Set `navigation`{.variable}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-2}
    to the result of [getting the navigation API entry
    index](#getting-the-navigation-api-entry-index){#navigation-api-entry-updates:getting-the-navigation-api-entry-index}
    of `initialSHE`{.variable} within `navigation`{.variable}.

To [update the navigation API entries for
reactivation]{#update-the-navigation-api-entries-for-reactivation .dfn}
given a
[`Navigation`{#navigation-api-entry-updates:navigation-2}](#navigation)
`navigation`{.variable}, a
[list](https://infra.spec.whatwg.org/#list){#navigation-api-entry-updates:list-2
x-internal="list"} of [session history
entries](browsing-the-web.html#session-history-entry){#navigation-api-entry-updates:session-history-entry-3}
`newSHEs`{.variable}, and a [session history
entry](browsing-the-web.html#session-history-entry){#navigation-api-entry-updates:session-history-entry-4}
`reactivatedSHE`{.variable}:

1.  If `navigation`{.variable} [has entries and events
    disabled](#has-entries-and-events-disabled){#navigation-api-entry-updates:has-entries-and-events-disabled-2},
    then return.

2.  Let `newNHEs`{.variable} be a new empty
    [list](https://infra.spec.whatwg.org/#list){#navigation-api-entry-updates:list-3
    x-internal="list"}.

3.  Let `oldNHEs`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#navigation-api-entry-updates:list-clone
    x-internal="list-clone"} of `navigation`{.variable}\'s [entry
    list](#navigation-entry-list){#navigation-api-entry-updates:navigation-entry-list-3}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#navigation-api-entry-updates:list-iterate-2
    x-internal="list-iterate"} `newSHE`{.variable} of
    `newSHEs`{.variable}:

    1.  Let `newNHE`{.variable} be null.

    2.  If `oldNHEs`{.variable}
        [contains](https://infra.spec.whatwg.org/#list-contain){#navigation-api-entry-updates:list-contains
        x-internal="list-contains"} a
        [`NavigationHistoryEntry`{#navigation-api-entry-updates:navigationhistoryentry-2}](#navigationhistoryentry)
        `matchingOldNHE`{.variable} whose [session history
        entry](#nhe-she){#navigation-api-entry-updates:nhe-she-2} is
        `newSHE`{.variable}, then:

        1.  Set `newNHE`{.variable} to `matchingOldNHE`{.variable}.

        2.  [Remove](https://infra.spec.whatwg.org/#list-remove){#navigation-api-entry-updates:list-remove
            x-internal="list-remove"} `matchingOldNHE`{.variable} from
            `oldNHEs`{.variable}.

    3.  Otherwise:

        1.  Set `newNHE`{.variable} to a
            [new](https://webidl.spec.whatwg.org/#new){#navigation-api-entry-updates:new-2
            x-internal="new"}
            [`NavigationHistoryEntry`{#navigation-api-entry-updates:navigationhistoryentry-3}](#navigationhistoryentry)
            created in the [relevant
            realm](webappapis.html#concept-relevant-realm){#navigation-api-entry-updates:concept-relevant-realm-2}
            of `navigation`{.variable}.

        2.  Set `newNHE`{.variable}\'s [session history
            entry](#nhe-she){#navigation-api-entry-updates:nhe-she-3} to
            `newSHE`{.variable}.

    4.  [Append](https://infra.spec.whatwg.org/#list-append){#navigation-api-entry-updates:list-append-2
        x-internal="list-append"} `newNHE`{.variable} to
        `newNHEs`{.variable}.

    `newSHEs`{.variable} will have originally come from [getting session
    history entries for the navigation
    API](browsing-the-web.html#getting-session-history-entries-for-the-navigation-api){#navigation-api-entry-updates:getting-session-history-entries-for-the-navigation-api-2},
    and thus each `newSHE`{.variable} will be contiguous
    [same](browsers.html#same-origin){#navigation-api-entry-updates:same-origin-2}
    [origin](browsing-the-web.html#document-state-origin){#navigation-api-entry-updates:document-state-origin-2}
    with `reactivatedSHE`{.variable}.

    By the end of this loop, all
    [`NavigationHistoryEntry`{#navigation-api-entry-updates:navigationhistoryentry-4}](#navigationhistoryentry)s
    that remain in `oldNHEs`{.variable} represent [session history
    entries](browsing-the-web.html#session-history-entry){#navigation-api-entry-updates:session-history-entry-5}
    which have been disposed while the
    [`Document`{#navigation-api-entry-updates:document}](dom.html#document)
    was in [bfcache](browsing-the-web.html#note-bfcache).

5.  Set `navigation`{.variable}\'s [entry
    list](form-control-infrastructure.html#entry-list){#navigation-api-entry-updates:entry-list}
    to `newNHEs`{.variable}.

6.  Set `navigation`{.variable}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-3}
    to the result of [getting the navigation API entry
    index](#getting-the-navigation-api-entry-index){#navigation-api-entry-updates:getting-the-navigation-api-entry-index-2}
    of `reactivatedSHE`{.variable} within `navigation`{.variable}.

7.  [Queue a global
    task](webappapis.html#queue-a-global-task){#navigation-api-entry-updates:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#navigation-api-entry-updates:navigation-and-traversal-task-source}
    given `navigation`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#navigation-api-entry-updates:concept-relevant-global}
    to run the following steps:

    1.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#navigation-api-entry-updates:list-iterate-3
        x-internal="list-iterate"} `disposedNHE`{.variable} of
        `oldNHEs`{.variable}:

        1.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#navigation-api-entry-updates:concept-event-fire
            x-internal="concept-event-fire"} named
            [`dispose`{#navigation-api-entry-updates:event-dispose}](indices.html#event-dispose)
            at `disposedNHE`{.variable}.

    ::: note
    We delay these steps by a task to ensure that
    [`dispose`{#navigation-api-entry-updates:event-dispose-2}](indices.html#event-dispose)
    events will fire after the
    [`pageshow`{#navigation-api-entry-updates:event-pageshow}](indices.html#event-pageshow)
    event. This ensures that
    [`pageshow`{#navigation-api-entry-updates:event-pageshow-2}](indices.html#event-pageshow)
    is the first event a page receives upon
    [reactivation](browsing-the-web.html#reactivate-a-document){#navigation-api-entry-updates:reactivate-a-document}.

    (However, the rest of this algorithm runs before the
    [`pageshow`{#navigation-api-entry-updates:event-pageshow-3}](indices.html#event-pageshow)
    event fires. This ensures that
    [`navigation.entries()`{#navigation-api-entry-updates:dom-navigation-entries}](#dom-navigation-entries)
    and
    [`navigation.currentEntry`{#navigation-api-entry-updates:dom-navigation-currententry}](#dom-navigation-currententry)
    will have correctly-updated values during any
    [`pageshow`{#navigation-api-entry-updates:event-pageshow-4}](indices.html#event-pageshow)
    event handlers.)
    :::

To [update the navigation API entries for a same-document
navigation]{#update-the-navigation-api-entries-for-a-same-document-navigation
.dfn} given a
[`Navigation`{#navigation-api-entry-updates:navigation-3}](#navigation)
`navigation`{.variable}, a [session history
entry](browsing-the-web.html#session-history-entry){#navigation-api-entry-updates:session-history-entry-6}
`destinationSHE`{.variable}, and a
[`NavigationType`{#navigation-api-entry-updates:navigationtype}](#navigationtype)
`navigationType`{.variable}:

1.  If `navigation`{.variable} [has entries and events
    disabled](#has-entries-and-events-disabled){#navigation-api-entry-updates:has-entries-and-events-disabled-3},
    then return.

2.  Let `oldCurrentNHE`{.variable} be the [current
    entry](#navigation-current-entry){#navigation-api-entry-updates:navigation-current-entry}
    of `navigation`{.variable}.

3.  Let `disposedNHEs`{.variable} be a new empty
    [list](https://infra.spec.whatwg.org/#list){#navigation-api-entry-updates:list-4
    x-internal="list"}.

4.  If `navigationType`{.variable} is
    \"[`traverse`{#navigation-api-entry-updates:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\",
    then:

    1.  Set `navigation`{.variable}\'s [current entry
        index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-4}
        to the result of [getting the navigation API entry
        index](#getting-the-navigation-api-entry-index){#navigation-api-entry-updates:getting-the-navigation-api-entry-index-3}
        of `destinationSHE`{.variable} within `navigation`{.variable}.

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#navigation-api-entry-updates:assert-3
        x-internal="assert"}: `navigation`{.variable}\'s [current entry
        index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-5}
        is not −1.

    This algorithm is only called for same-document traversals.
    Cross-document traversals will instead call either [initialize the
    navigation API entries for a new
    document](#initialize-the-navigation-api-entries-for-a-new-document){#navigation-api-entry-updates:initialize-the-navigation-api-entries-for-a-new-document}
    or [update the navigation API entries for
    reactivation](#update-the-navigation-api-entries-for-reactivation){#navigation-api-entry-updates:update-the-navigation-api-entries-for-reactivation}.

5.  Otherwise, if `navigationType`{.variable} is
    \"[`push`{#navigation-api-entry-updates:dom-navigationtype-push}](#dom-navigationtype-push)\",
    then:

    1.  Set `navigation`{.variable}\'s [current entry
        index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-6}
        to `navigation`{.variable}\'s [current entry
        index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-7} +
        1.

    2.  Let `i`{.variable} be `navigation`{.variable}\'s [current entry
        index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-8}.

    3.  [While](https://infra.spec.whatwg.org/#iteration-while){#navigation-api-entry-updates:while
        x-internal="while"} `i`{.variable} \< `navigation`{.variable}\'s
        [entry
        list](#navigation-entry-list){#navigation-api-entry-updates:navigation-entry-list-4}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#navigation-api-entry-updates:list-size
        x-internal="list-size"}:

        1.  [Append](https://infra.spec.whatwg.org/#list-append){#navigation-api-entry-updates:list-append-3
            x-internal="list-append"} `navigation`{.variable}\'s [entry
            list](#navigation-entry-list){#navigation-api-entry-updates:navigation-entry-list-5}\[`i`{.variable}\]
            to `disposedNHEs`{.variable}.

        2.  Set `i`{.variable} to `i`{.variable} + 1.

    4.  [Remove](https://infra.spec.whatwg.org/#list-remove){#navigation-api-entry-updates:list-remove-2
        x-internal="list-remove"} all items in `disposedNHEs`{.variable}
        from `navigation`{.variable}\'s [entry
        list](#navigation-entry-list){#navigation-api-entry-updates:navigation-entry-list-6}.

6.  Otherwise, if `navigationType`{.variable} is
    \"[`replace`{#navigation-api-entry-updates:dom-navigationtype-replace}](#dom-navigationtype-replace)\",
    then:

    1.  [Append](https://infra.spec.whatwg.org/#list-append){#navigation-api-entry-updates:list-append-4
        x-internal="list-append"} `oldCurrentNHE`{.variable} to
        `disposedNHEs`{.variable}.

7.  If `navigationType`{.variable} is
    \"[`push`{#navigation-api-entry-updates:dom-navigationtype-push-2}](#dom-navigationtype-push)\"
    or
    \"[`replace`{#navigation-api-entry-updates:dom-navigationtype-replace-2}](#dom-navigationtype-replace)\",
    then:

    1.  Let `newNHE`{.variable} be a
        [new](https://webidl.spec.whatwg.org/#new){#navigation-api-entry-updates:new-3
        x-internal="new"}
        [`NavigationHistoryEntry`{#navigation-api-entry-updates:navigationhistoryentry-5}](#navigationhistoryentry)
        created in the [relevant
        realm](webappapis.html#concept-relevant-realm){#navigation-api-entry-updates:concept-relevant-realm-3}
        of `navigation`{.variable}.

    2.  Set `newNHE`{.variable}\'s [session history
        entry](#nhe-she){#navigation-api-entry-updates:nhe-she-4} to
        `destinationSHE`{.variable}.

    3.  Set `navigation`{.variable}\'s [entry
        list](#navigation-entry-list){#navigation-api-entry-updates:navigation-entry-list-7}\[`navigation`{.variable}\'s
        [current entry
        index](#navigation-current-entry-index){#navigation-api-entry-updates:navigation-current-entry-index-9}\]
        to `newNHE`{.variable}.

8.  If `navigation`{.variable}\'s [ongoing API method
    tracker](#ongoing-api-method-tracker){#navigation-api-entry-updates:ongoing-api-method-tracker}
    is non-null, then [notify about the committed-to
    entry](#notify-about-the-committed-to-entry){#navigation-api-entry-updates:notify-about-the-committed-to-entry}
    given `navigation`{.variable}\'s [ongoing API method
    tracker](#ongoing-api-method-tracker){#navigation-api-entry-updates:ongoing-api-method-tracker-2}
    and the [current
    entry](#navigation-current-entry){#navigation-api-entry-updates:navigation-current-entry-2}
    of `navigation`{.variable}.

    It is important to do this before firing the
    [`dispose`{#navigation-api-entry-updates:event-dispose-3}](indices.html#event-dispose)
    or
    [`currententrychange`{#navigation-api-entry-updates:event-currententrychange}](indices.html#event-currententrychange)
    events, since event handlers could start another navigation, or
    otherwise change the value of `navigation`{.variable}\'s [ongoing
    API method
    tracker](#ongoing-api-method-tracker){#navigation-api-entry-updates:ongoing-api-method-tracker-3}.

9.  [Prepare to run
    script](webappapis.html#prepare-to-run-script){#navigation-api-entry-updates:prepare-to-run-script}
    given `navigation`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#navigation-api-entry-updates:relevant-settings-object}.

    See [the discussion for other navigation API
    events](#note-suppress-microtasks-during-navigation-events) to
    understand why we do this.

10. [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#navigation-api-entry-updates:concept-event-fire-2
    x-internal="concept-event-fire"} named
    [`currententrychange`{#navigation-api-entry-updates:event-currententrychange-2}](indices.html#event-currententrychange)
    at `navigation`{.variable} using
    [`NavigationCurrentEntryChangeEvent`{#navigation-api-entry-updates:navigationcurrententrychangeevent}](#navigationcurrententrychangeevent),
    with its
    [`navigationType`{#navigation-api-entry-updates:dom-navigationcurrententrychangeevent-navigationtype}](#dom-navigationcurrententrychangeevent-navigationtype)
    attribute initialized to `navigationType`{.variable} and its
    [`from`{#navigation-api-entry-updates:dom-navigationcurrententrychangeevent-from}](#dom-navigationcurrententrychangeevent-from)
    initialized to `oldCurrentNHE`{.variable}.

11. [For
    each](https://infra.spec.whatwg.org/#list-iterate){#navigation-api-entry-updates:list-iterate-4
    x-internal="list-iterate"} `disposedNHE`{.variable} of
    `disposedNHEs`{.variable}:

    1.  [Fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#navigation-api-entry-updates:concept-event-fire-3
        x-internal="concept-event-fire"} named
        [`dispose`{#navigation-api-entry-updates:event-dispose-4}](indices.html#event-dispose)
        at `disposedNHE`{.variable}.

12. [Clean up after running
    script](webappapis.html#clean-up-after-running-script){#navigation-api-entry-updates:clean-up-after-running-script}
    given `navigation`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#navigation-api-entry-updates:relevant-settings-object-2}.

In implementations, same-document navigations can cause [session history
entries](browsing-the-web.html#session-history-entry){#navigation-api-entry-updates:session-history-entry-7}
to be disposed by falling off the back of the session history entry
list. This is not yet handled by the above algorithm (or by any other
part of this standard). See [issue
#8620](https://github.com/whatwg/html/issues/8620) to track progress on
defining the correct behavior in such cases.

##### [7.2.6.5]{.secno} The [`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry}](#navigationhistoryentry) interface[](#the-navigationhistoryentry-interface){.self-link}

``` idl
[Exposed=Window]
interface NavigationHistoryEntry : EventTarget {
  readonly attribute USVString? url;
  readonly attribute DOMString key;
  readonly attribute DOMString id;
  readonly attribute long long index;
  readonly attribute boolean sameDocument;

  any getState();

  attribute EventHandler ondispose;
};
```

`entry`{.variable}`.`[`url`](#dom-navigationhistoryentry-url){#dom-navigationhistoryentry-url-dev}

:   The URL of this navigation history entry.

    This can return null if the entry corresponds to a different
    [`Document`{#the-navigationhistoryentry-interface:document}](dom.html#document)
    than the current one (i.e., if
    [`sameDocument`{#the-navigationhistoryentry-interface:dom-navigationhistoryentry-samedocument-2}](#dom-navigationhistoryentry-samedocument)
    is false), and that
    [`Document`{#the-navigationhistoryentry-interface:document-2}](dom.html#document)
    was fetched with a [referrer
    policy](https://w3c.github.io/webappsec-referrer-policy/#referrer-policy){#the-navigationhistoryentry-interface:referrer-policy
    x-internal="referrer-policy"} of \"`no-referrer`\" or \"`origin`\",
    since that indicates the
    [`Document`{#the-navigationhistoryentry-interface:document-3}](dom.html#document)
    in question is hiding its URL even from other same-origin pages.

`entry`{.variable}`.`[`key`](#dom-navigationhistoryentry-key){#dom-navigationhistoryentry-key-dev}

:   A user agent-generated random UUID string representing this
    navigation history entry\'s place in the navigation history list.
    This value will be reused by other
    [`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-2}](#navigationhistoryentry)
    instances that replace this one due to
    \"[`replace`{#the-navigationhistoryentry-interface:dom-navigationtype-replace}](#dom-navigationtype-replace)\"
    navigations, and will survive reloads and session restores.

    This is useful for navigating back to this entry in the navigation
    history list, using
    [`navigation.traverseTo(key)`{#the-navigationhistoryentry-interface:dom-navigation-traverseto}](#dom-navigation-traverseto).

`entry`{.variable}`.`[`id`](#dom-navigationhistoryentry-id){#dom-navigationhistoryentry-id-dev}

:   A user agent-generated random UUID string representing this specific
    navigation history entry. This value will *not* be reused by other
    [`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-3}](#navigationhistoryentry)
    instances. This value will survive reloads and session restores.

    This is useful for associating data with this navigation history
    entry using other storage APIs.

`entry`{.variable}`.`[`index`](#dom-navigationhistoryentry-index){#dom-navigationhistoryentry-index-dev}
:   The index of this
    [`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-4}](#navigationhistoryentry)
    within
    [`navigation.entries()`{#the-navigationhistoryentry-interface:dom-navigation-entries}](#dom-navigation-entries),
    or −1 if the entry is not in the navigation history entry list.

`entry`{.variable}`.`[`sameDocument`](#dom-navigationhistoryentry-samedocument){#dom-navigationhistoryentry-samedocument-dev}
:   Indicates whether or not this navigation history entry is for the
    same
    [`Document`{#the-navigationhistoryentry-interface:document-4}](dom.html#document)
    as the current one, or not. This will be true, for example, when the
    entry represents a fragment navigation or single-page app
    navigation.

`entry`{.variable}`.`[`getState`](#dom-navigationhistoryentry-getstate){#dom-navigationhistoryentry-getstate-dev}`()`

:   Returns the
    [deserialization](structured-data.html#structureddeserialize){#the-navigationhistoryentry-interface:structureddeserialize}
    of the state stored in this entry, which was added to the entry
    using
    [`navigation.navigate()`{#the-navigationhistoryentry-interface:dom-navigation-navigate}](#dom-navigation-navigate)
    or
    [`navigation.updateCurrentEntry()`{#the-navigationhistoryentry-interface:dom-navigation-updatecurrententry}](#dom-navigation-updatecurrententry).
    This state survives reloads and session restores.

    Note that in general, unless the state value is a primitive,
    `entry.getState() !== entry.getState()`{.js}, since a fresh
    deserialization is returned each time.

    This state is unrelated to the classic history API\'s
    [`history.state`{#the-navigationhistoryentry-interface:dom-history-state}](#dom-history-state).

Each
[`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-5}](#navigationhistoryentry)
has an associated [session history entry]{#nhe-she .dfn}, which is a
[session history
entry](browsing-the-web.html#session-history-entry){#the-navigationhistoryentry-interface:session-history-entry}.

The [key]{#concept-navigationhistoryentry-key .dfn} of a
[`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-6}](#navigationhistoryentry)
`nhe`{.variable} is given by the return value of the following
algorithm:

1.  If `nhe`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigationhistoryentry-interface:concept-relevant-global}\'s
    [associated
    `Document`](#concept-document-window){#the-navigationhistoryentry-interface:concept-document-window}
    is not [fully
    active](document-sequences.html#fully-active){#the-navigationhistoryentry-interface:fully-active},
    then return the empty string.

2.  Return `nhe`{.variable}\'s [session history
    entry](#nhe-she){#the-navigationhistoryentry-interface:nhe-she}\'s
    [navigation API
    key](browsing-the-web.html#she-navigation-api-key){#the-navigationhistoryentry-interface:she-navigation-api-key}.

The [ID]{#concept-navigationhistoryentry-id .dfn} of a
[`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-7}](#navigationhistoryentry)
`nhe`{.variable} is given by the return value of the following
algorithm:

1.  If `nhe`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigationhistoryentry-interface:concept-relevant-global-2}\'s
    [associated
    `Document`](#concept-document-window){#the-navigationhistoryentry-interface:concept-document-window-2}
    is not [fully
    active](document-sequences.html#fully-active){#the-navigationhistoryentry-interface:fully-active-2},
    then return the empty string.

2.  Return `nhe`{.variable}\'s [session history
    entry](#nhe-she){#the-navigationhistoryentry-interface:nhe-she-2}\'s
    [navigation API
    ID](browsing-the-web.html#she-navigation-api-id){#the-navigationhistoryentry-interface:she-navigation-api-id}.

The [index]{#concept-navigationhistoryentry-index .dfn} of a
[`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-8}](#navigationhistoryentry)
`nhe`{.variable} is given by the return value of the following
algorithm:

1.  If `nhe`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigationhistoryentry-interface:concept-relevant-global-3}\'s
    [associated
    `Document`](#concept-document-window){#the-navigationhistoryentry-interface:concept-document-window-3}
    is not [fully
    active](document-sequences.html#fully-active){#the-navigationhistoryentry-interface:fully-active-3},
    then return −1.

2.  Return the result of [getting the navigation API entry
    index](#getting-the-navigation-api-entry-index){#the-navigationhistoryentry-interface:getting-the-navigation-api-entry-index}
    of
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this
    x-internal="this"}\'s [session history
    entry](#nhe-she){#the-navigationhistoryentry-interface:nhe-she-3}
    within
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigationhistoryentry-interface:concept-relevant-global-4}\'s
    [navigation
    API](#window-navigation-api){#the-navigationhistoryentry-interface:window-navigation-api}.

The [`url`]{#dom-navigationhistoryentry-url .dfn
dfn-for="NavigationHistoryEntry" dfn-type="attribute"} getter steps are:

1.  Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-3
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigationhistoryentry-interface:concept-relevant-global-5}\'s
    [associated
    `Document`](#concept-document-window){#the-navigationhistoryentry-interface:concept-document-window-4}.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#the-navigationhistoryentry-interface:fully-active-4},
    then return the empty string.

3.  Let `she`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-4
    x-internal="this"}\'s [session history
    entry](#nhe-she){#the-navigationhistoryentry-interface:nhe-she-4}.

4.  If `she`{.variable}\'s
    [document](browsing-the-web.html#she-document){#the-navigationhistoryentry-interface:she-document}
    does not equal `document`{.variable}, and `she`{.variable}\'s
    [document
    state](browsing-the-web.html#she-document-state){#the-navigationhistoryentry-interface:she-document-state}\'s
    [request referrer
    policy](browsing-the-web.html#document-state-request-referrer-policy){#the-navigationhistoryentry-interface:document-state-request-referrer-policy}
    is \"`no-referrer`\" or \"`origin`\", then return null.

5.  Return `she`{.variable}\'s
    [URL](browsing-the-web.html#she-url){#the-navigationhistoryentry-interface:she-url},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#the-navigationhistoryentry-interface:concept-url-serializer
    x-internal="concept-url-serializer"}.

The [`key`]{#dom-navigationhistoryentry-key .dfn
dfn-for="NavigationHistoryEntry" dfn-type="attribute"} getter steps are
to return
[this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-5
x-internal="this"}\'s
[key](#concept-navigationhistoryentry-key){#the-navigationhistoryentry-interface:concept-navigationhistoryentry-key}.

The [`id`]{#dom-navigationhistoryentry-id .dfn
dfn-for="NavigationHistoryEntry" dfn-type="attribute"} getter steps are
to return
[this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-6
x-internal="this"}\'s
[ID](#concept-navigationhistoryentry-id){#the-navigationhistoryentry-interface:concept-navigationhistoryentry-id}.

The [`index`]{#dom-navigationhistoryentry-index .dfn
dfn-for="NavigationHistoryEntry" dfn-type="attribute"} getter steps are
to return
[this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-7
x-internal="this"}\'s
[index](#concept-navigationhistoryentry-index){#the-navigationhistoryentry-interface:concept-navigationhistoryentry-index}.

The [`sameDocument`]{#dom-navigationhistoryentry-samedocument .dfn
dfn-for="NavigationHistoryEntry" dfn-type="attribute"} getter steps are:

1.  Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-8
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigationhistoryentry-interface:concept-relevant-global-6}\'s
    [associated
    `Document`](#concept-document-window){#the-navigationhistoryentry-interface:concept-document-window-5}.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#the-navigationhistoryentry-interface:fully-active-5},
    then return false.

3.  Return true if
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-9
    x-internal="this"}\'s [session history
    entry](#nhe-she){#the-navigationhistoryentry-interface:nhe-she-5}\'s
    [document](browsing-the-web.html#she-document){#the-navigationhistoryentry-interface:she-document-2}
    equals `document`{.variable}, and false otherwise.

The [`getState()`]{#dom-navigationhistoryentry-getstate .dfn
dfn-for="NavigationHistoryEntry" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-10
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigationhistoryentry-interface:concept-relevant-global-7}\'s
    [associated
    `Document`](#concept-document-window){#the-navigationhistoryentry-interface:concept-document-window-6}
    is not [fully
    active](document-sequences.html#fully-active){#the-navigationhistoryentry-interface:fully-active-6},
    then return undefined.

2.  Return
    [StructuredDeserialize](structured-data.html#structureddeserialize){#the-navigationhistoryentry-interface:structureddeserialize-2}([this](https://webidl.spec.whatwg.org/#this){#the-navigationhistoryentry-interface:this-11
    x-internal="this"}\'s [session history
    entry](#nhe-she){#the-navigationhistoryentry-interface:nhe-she-6}\'s
    [navigation API
    state](browsing-the-web.html#she-navigation-api-state){#the-navigationhistoryentry-interface:she-navigation-api-state}).
    Rethrow any exceptions.

    This can in theory throw an exception, if attempting to deserialize
    a large
    [`ArrayBuffer`{#the-navigationhistoryentry-interface:idl-arraybuffer}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
    when not enough memory is available.

The following are the [event
handlers](webappapis.html#event-handlers){#the-navigationhistoryentry-interface:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#the-navigationhistoryentry-interface:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-navigationhistoryentry-interface:event-handler-idl-attributes},
by all objects implementing the
[`NavigationHistoryEntry`{#the-navigationhistoryentry-interface:navigationhistoryentry-9}](#navigationhistoryentry)
interface:

[Event
handler](webappapis.html#event-handlers){#the-navigationhistoryentry-interface:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#the-navigationhistoryentry-interface:event-handler-event-type-2}

[`ondispose`]{#handler-navigationhistoryentry-ondispose .dfn
dfn-for="NavigationHistoryEntry" dfn-type="attribute"}

[`dispose`{#the-navigationhistoryentry-interface:event-dispose}](indices.html#event-dispose)

##### [7.2.6.6]{.secno} The history entry list[](#the-history-entry-list){.self-link}

`entries`{.variable}` = `[`navigation`](#dom-navigation){#the-history-entry-list:dom-navigation}`.`[`entries()`](#dom-navigation-entries){#dom-navigation-entries-dev}

:   Returns an array of
    [`NavigationHistoryEntry`{#the-history-entry-list:navigationhistoryentry}](#navigationhistoryentry)
    instances represent the current navigation history entry list, i.e.,
    all [session history
    entries](browsing-the-web.html#session-history-entry){#the-history-entry-list:session-history-entry}
    for this
    [navigable](document-sequences.html#navigable){#the-history-entry-list:navigable}
    that are [same
    origin](browsers.html#same-origin){#the-history-entry-list:same-origin}
    and contiguous to the [current session history
    entry](document-sequences.html#nav-current-history-entry){#the-history-entry-list:nav-current-history-entry}.

[`navigation`](#dom-navigation){#the-history-entry-list:dom-navigation-2}`.`[`currentEntry`](#dom-navigation-currententry){#dom-navigation-currententry-dev}

:   Returns the
    [`NavigationHistoryEntry`{#the-history-entry-list:navigationhistoryentry-2}](#navigationhistoryentry)
    corresponding to the [current session history
    entry](document-sequences.html#nav-current-history-entry){#the-history-entry-list:nav-current-history-entry-2}.

[`navigation`](#dom-navigation){#the-history-entry-list:dom-navigation-3}`.`[`updateCurrentEntry`](#dom-navigation-updatecurrententry){#dom-navigation-updatecurrententry-dev}`({ `[`state`](#dom-navigationupdatecurrententryoptions-state){#the-history-entry-list:dom-navigationupdatecurrententryoptions-state}` })`

:   Updates the [navigation API
    state](browsing-the-web.html#she-navigation-api-state){#the-history-entry-list:she-navigation-api-state}
    of the [current session history
    entry](document-sequences.html#nav-current-history-entry){#the-history-entry-list:nav-current-history-entry-3},
    without performing a navigation like
    [`navigation.reload()`{#the-history-entry-list:dom-navigation-reload}](#dom-navigation-reload)
    would do.

    This method is best used to capture updates to the page that have
    already happened, and need to be reflected into the navigation API
    state. For cases where the state update is meant to drive a page
    update, instead use
    [`navigation.navigate()`{#the-history-entry-list:dom-navigation-navigate}](#dom-navigation-navigate)
    or
    [`navigation.reload()`{#the-history-entry-list:dom-navigation-reload-2}](#dom-navigation-reload),
    which will trigger a
    [`navigate`{#the-history-entry-list:event-navigate}](indices.html#event-navigate)
    event.

[`navigation`](#dom-navigation){#the-history-entry-list:dom-navigation-4}`.`[`canGoBack`](#dom-navigation-cangoback){#dom-navigation-cangoback-dev}

:   Returns true if the current [current session history
    entry](document-sequences.html#nav-current-history-entry){#the-history-entry-list:nav-current-history-entry-4}
    (i.e.,
    [`currentEntry`{#the-history-entry-list:dom-navigation-currententry}](#dom-navigation-currententry))
    is not the first one in the navigation history entry list (i.e., in
    [`entries()`{#the-history-entry-list:dom-navigation-entries}](#dom-navigation-entries)).
    This means that there is a previous [session history
    entry](browsing-the-web.html#session-history-entry){#the-history-entry-list:session-history-entry-2}
    for this
    [navigable](document-sequences.html#navigable){#the-history-entry-list:navigable-2},
    and its [document
    state](browsing-the-web.html#she-document-state){#the-history-entry-list:she-document-state}\'s
    [origin](browsing-the-web.html#document-state-origin){#the-history-entry-list:document-state-origin}
    is [same
    origin](browsers.html#same-origin){#the-history-entry-list:same-origin-2}
    with the current
    [`Document`{#the-history-entry-list:document}](dom.html#document)\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-history-entry-list:concept-document-origin
    x-internal="concept-document-origin"}.

[`navigation`](#dom-navigation){#the-history-entry-list:dom-navigation-5}`.`[`canGoForward`](#dom-navigation-cangoforward){#dom-navigation-cangoforward-dev}

:   Returns true if the current [current session history
    entry](document-sequences.html#nav-current-history-entry){#the-history-entry-list:nav-current-history-entry-5}
    (i.e.,
    [`currentEntry`{#the-history-entry-list:dom-navigation-currententry-2}](#dom-navigation-currententry))
    is not the last one in the navigation history entry list (i.e., in
    [`entries()`{#the-history-entry-list:dom-navigation-entries-2}](#dom-navigation-entries)).
    This means that there is a next [session history
    entry](browsing-the-web.html#session-history-entry){#the-history-entry-list:session-history-entry-3}
    for this
    [navigable](document-sequences.html#navigable){#the-history-entry-list:navigable-3},
    and its [document
    state](browsing-the-web.html#she-document-state){#the-history-entry-list:she-document-state-2}\'s
    [origin](browsing-the-web.html#document-state-origin){#the-history-entry-list:document-state-origin-2}
    is [same
    origin](browsers.html#same-origin){#the-history-entry-list:same-origin-3}
    with the current
    [`Document`{#the-history-entry-list:document-2}](dom.html#document)\'s
    [origin](https://dom.spec.whatwg.org/#concept-document-origin){#the-history-entry-list:concept-document-origin-2
    x-internal="concept-document-origin"}.

The [`entries()`]{#dom-navigation-entries .dfn dfn-for="Navigation"
dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this
    x-internal="this"} [has entries and events
    disabled](#has-entries-and-events-disabled){#the-history-entry-list:has-entries-and-events-disabled},
    then return the empty list.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-2
    x-internal="this"}\'s [entry
    list](#navigation-entry-list){#the-history-entry-list:navigation-entry-list}.

    Recall that because of Web IDL\'s sequence type conversion rules,
    this will create a new JavaScript array object on each call. That
    is,
    [`navigation.entries()`](#dom-navigation-entries){#the-history-entry-list:dom-navigation-entries-3}` !== `[`navigation.entries()`](#dom-navigation-entries){#the-history-entry-list:dom-navigation-entries-4}.

The [`currentEntry`]{#dom-navigation-currententry .dfn
dfn-for="Navigation" dfn-type="attribute"} getter steps are to return
the [current
entry](#navigation-current-entry){#the-history-entry-list:navigation-current-entry}
of
[this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-3
x-internal="this"}.

The
[`updateCurrentEntry(``options`{.variable}`)`]{#dom-navigation-updatecurrententry
.dfn dfn-for="Navigation" dfn-type="method"} method steps are:

1.  Let `current`{.variable} be the [current
    entry](#navigation-current-entry){#the-history-entry-list:navigation-current-entry-2}
    of
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-4
    x-internal="this"}.

2.  If `current`{.variable} is null, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-history-entry-list:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-history-entry-list:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `serializedState`{.variable} be
    [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#the-history-entry-list:structuredserializeforstorage}(`options`{.variable}\[\"[`state`{#the-history-entry-list:dom-navigationupdatecurrententryoptions-state-2}](#dom-navigationupdatecurrententryoptions-state)\"\]),
    rethrowing any exceptions.

4.  Set `current`{.variable}\'s [session history
    entry](#nhe-she){#the-history-entry-list:nhe-she}\'s [navigation API
    state](browsing-the-web.html#she-navigation-api-state){#the-history-entry-list:she-navigation-api-state-2}
    to `serializedState`{.variable}.

5.  [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#the-history-entry-list:concept-event-fire
    x-internal="concept-event-fire"} named
    [`currententrychange`{#the-history-entry-list:event-currententrychange}](indices.html#event-currententrychange)
    at
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-5
    x-internal="this"} using
    [`NavigationCurrentEntryChangeEvent`{#the-history-entry-list:navigationcurrententrychangeevent}](#navigationcurrententrychangeevent),
    with its
    [`navigationType`{#the-history-entry-list:dom-navigationcurrententrychangeevent-navigationtype}](#dom-navigationcurrententrychangeevent-navigationtype)
    attribute initialized to null and its
    [`from`{#the-history-entry-list:dom-navigationcurrententrychangeevent-from}](#dom-navigationcurrententrychangeevent-from)
    initialized to `current`{.variable}.

The [`canGoBack`]{#dom-navigation-cangoback .dfn dfn-for="Navigation"
dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-6
    x-internal="this"} [has entries and events
    disabled](#has-entries-and-events-disabled){#the-history-entry-list:has-entries-and-events-disabled-2},
    then return false.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#the-history-entry-list:assert
    x-internal="assert"}:
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-7
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#the-history-entry-list:navigation-current-entry-index}
    is not −1.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-8
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#the-history-entry-list:navigation-current-entry-index-2}
    is 0, then return false.

4.  Return true.

The [`canGoForward`]{#dom-navigation-cangoforward .dfn
dfn-for="Navigation" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-9
    x-internal="this"} [has entries and events
    disabled](#has-entries-and-events-disabled){#the-history-entry-list:has-entries-and-events-disabled-3},
    then return false.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#the-history-entry-list:assert-2
    x-internal="assert"}:
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-10
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#the-history-entry-list:navigation-current-entry-index-3}
    is not −1.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-11
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#the-history-entry-list:navigation-current-entry-index-4}
    is equal to
    [this](https://webidl.spec.whatwg.org/#this){#the-history-entry-list:this-12
    x-internal="this"}\'s [entry
    list](#navigation-entry-list){#the-history-entry-list:navigation-entry-list-2}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#the-history-entry-list:list-size
    x-internal="list-size"} − 1, then return false.

4.  Return true.

##### [7.2.6.7]{.secno} Initiating navigations[](#navigation-api-initiating-navigations){.self-link} {#navigation-api-initiating-navigations}

`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation}`.`[`navigate`](#dom-navigation-navigate){#dom-navigation-navigate-dev}`(``url`{.variable}`)`\
`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-2}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-2}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-2}`.`[`navigate`](#dom-navigation-navigate){#navigation-api-initiating-navigations:dom-navigation-navigate}`(``url`{.variable}`, ``options`{.variable}`)`

:   [Navigates](browsing-the-web.html#navigate){#navigation-api-initiating-navigations:navigate}
    the current page to the given `url`{.variable}. `options`{.variable}
    can contain the following values:

    - [`history`{#dom-navigationnavigateoptions-history-dev}](#dom-navigationnavigateoptions-history)
      can be set to
      \"[`replace`{#navigationhistorybehavior-replace-dev}](browsing-the-web.html#navigationhistorybehavior-replace)\"
      to replace the current session history entry, instead of pushing a
      new one.

    - [`info`{#dom-navigationoptions-info-dev}](#dom-navigationoptions-info)
      can be set to any value; it will populate the
      [`info`{#navigation-api-initiating-navigations:dom-navigateevent-info}](#dom-navigateevent-info)
      property of the corresponding
      [`NavigateEvent`{#navigation-api-initiating-navigations:navigateevent}](#navigateevent).

    - [`state`{#dom-navigationnavigateoptions-state-dev}](#dom-navigationnavigateoptions-state)
      can be set to any
      [serializable](structured-data.html#serializable-objects){#navigation-api-initiating-navigations:serializable-objects}
      value; it will populate the state retrieved by
      [`navigation.currentEntry.getState()`{#navigation-api-initiating-navigations:dom-navigationhistoryentry-getstate}](#dom-navigationhistoryentry-getstate)
      once the navigation completes, for same-document navigations. (It
      will be ignored for navigations that end up cross-document.)

    By default this will perform a full navigation (i.e., a
    cross-document navigation, unless the given URL differs only in a
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#navigation-api-initiating-navigations:concept-url-fragment
    x-internal="concept-url-fragment"} from the current one). The
    [`navigateEvent.intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept}](#dom-navigateevent-intercept)
    method can be used to convert it into a same-document navigation.

    The returned promises will behave as follows:

    - For navigations that get aborted, both promises will reject with
      an
      [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#navigation-api-initiating-navigations:aborterror
      x-internal="aborterror"}
      [`DOMException`{#navigation-api-initiating-navigations:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    - For same-document navigations created by using the
      [`navigateEvent.intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept-2}](#dom-navigateevent-intercept)
      method,
      [`committed`{#navigation-api-initiating-navigations:dom-navigationresult-committed-3}](#dom-navigationresult-committed)
      will fulfill immediately, and
      [`finished`{#navigation-api-initiating-navigations:dom-navigationresult-finished-3}](#dom-navigationresult-finished)
      will fulfill or reject according to any promsies returned by
      handlers passed to
      [`intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept-3}](#dom-navigateevent-intercept).

    - For other same-document navigations (e.g., non-intercepted
      [fragment
      navigations](browsing-the-web.html#navigate-fragid){#navigation-api-initiating-navigations:navigate-fragid}),
      both promises will fulfill immediately.

    - For cross-document navigations, or navigations that result in 204
      or 205
      [statuses](https://fetch.spec.whatwg.org/#concept-response-status){#navigation-api-initiating-navigations:concept-response-status
      x-internal="concept-response-status"} or
      \`[`Content-Disposition: attachment`{#navigation-api-initiating-navigations:http-content-disposition}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
      header fields from the server (and thus do not actually navigate),
      both promises will never settle.

    In all cases, when the returned promises fulfill, it will be with
    the
    [`NavigationHistoryEntry`{#navigation-api-initiating-navigations:navigationhistoryentry}](#navigationhistoryentry)
    that was navigated to.

`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-4}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-4}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-3}`.`[`reload`](#dom-navigation-reload){#dom-navigation-reload-dev}`(``options`{.variable}`)`

:   [Reloads](browsing-the-web.html#reload){#navigation-api-initiating-navigations:reload}
    the current page. `options`{.variable} can contain
    [`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info}](#dom-navigationoptions-info)
    and
    [`state`{#dom-navigationreloadoptions-state-dev}](#dom-navigationreloadoptions-state),
    which behave as described above.

    The default behavior of performing a from-network-or-cache reload of
    the current page can be overriden by the using the
    [`navigateEvent.intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept-4}](#dom-navigateevent-intercept)
    method. Doing so will mean this call only updates state or passes
    along the appropriate
    [`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info-2}](#dom-navigationoptions-info),
    plus performing whater actions the
    [`navigate`{#navigation-api-initiating-navigations:event-navigate}](indices.html#event-navigate)
    event handlers see fit to carry out.

    The returned promises will behave as follows:

    - If the reload is intercepted by using the
      [`navigateEvent.intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept-5}](#dom-navigateevent-intercept)
      method,
      [`committed`{#navigation-api-initiating-navigations:dom-navigationresult-committed-5}](#dom-navigationresult-committed)
      will fulfill immediately, and
      [`finished`{#navigation-api-initiating-navigations:dom-navigationresult-finished-5}](#dom-navigationresult-finished)
      will fulfill or reject according to any promsies returned by
      handlers passed to
      [`intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept-6}](#dom-navigateevent-intercept).

    - Otherwise, both promises will never settle.

`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-6}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-6}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-4}`.`[`traverseTo`](#dom-navigation-traverseto){#dom-navigation-traverseto-dev}`(``key`{.variable}`)`\
`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-7}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-7}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-5}`.`[`traverseTo`](#dom-navigation-traverseto){#navigation-api-initiating-navigations:dom-navigation-traverseto}`(``key`{.variable}`, { `[`info`](#dom-navigationoptions-info){#navigation-api-initiating-navigations:dom-navigationoptions-info-3}` })`

:   [Traverses](browsing-the-web.html#apply-the-traverse-history-step){#navigation-api-initiating-navigations:apply-the-traverse-history-step}
    to the closest [session history
    entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry}
    that matches the
    [`NavigationHistoryEntry`{#navigation-api-initiating-navigations:navigationhistoryentry-2}](#navigationhistoryentry)
    with the given `key`{.variable}.
    [`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info-4}](#dom-navigationoptions-info)
    can be set to any value; it will populate the
    [`info`{#navigation-api-initiating-navigations:dom-navigateevent-info-2}](#dom-navigateevent-info)
    property of the corresponding
    [`NavigateEvent`{#navigation-api-initiating-navigations:navigateevent-2}](#navigateevent).

    If a traversal to that [session history
    entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry-2}
    is already in progress, then this will return the promises for that
    original traversal, and
    [`info`{#navigation-api-initiating-navigations:dom-navigateevent-info-3}](#dom-navigateevent-info)
    will be ignored.

    The returned promises will behave as follows:

    - If there is no
      [`NavigationHistoryEntry`{#navigation-api-initiating-navigations:navigationhistoryentry-3}](#navigationhistoryentry)
      in
      [`navigation.entries()`{#navigation-api-initiating-navigations:dom-navigation-entries}](#dom-navigation-entries)
      whose
      [`key`{#navigation-api-initiating-navigations:dom-navigationhistoryentry-key}](#dom-navigationhistoryentry-key)
      matches `key`{.variable}, both promises will reject with an
      [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror
      x-internal="invalidstateerror"}
      [`DOMException`{#navigation-api-initiating-navigations:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    - For same-document traversals intercepted by the
      [`navigateEvent.intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept-7}](#dom-navigateevent-intercept)
      method,
      [`committed`{#navigation-api-initiating-navigations:dom-navigationresult-committed-8}](#dom-navigationresult-committed)
      will fulfill as soon as the traversal is processed and
      [`navigation.currentEntry`{#navigation-api-initiating-navigations:dom-navigation-currententry}](#dom-navigation-currententry)
      is updated, and
      [`finished`{#navigation-api-initiating-navigations:dom-navigationresult-finished-8}](#dom-navigationresult-finished)
      will fulfill or reject according to any promsies returned by the
      handlers passed to
      [`intercept()`{#navigation-api-initiating-navigations:dom-navigateevent-intercept-8}](#dom-navigateevent-intercept).

    - For non-intercepted same-document travesals, both promises will
      fulfill as soon as the traversal is processed and
      [`navigation.currentEntry`{#navigation-api-initiating-navigations:dom-navigation-currententry-2}](#dom-navigation-currententry)
      is updated.

    - For cross-document traversals, including attempted cross-document
      traversals that end up resulting in a 204 or 205
      [statuses](https://fetch.spec.whatwg.org/#concept-response-status){#navigation-api-initiating-navigations:concept-response-status-2
      x-internal="concept-response-status"} or
      \`[`Content-Disposition: attachment`{#navigation-api-initiating-navigations:http-content-disposition-2}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
      header fields from the server (and thus do not actually traverse),
      both promises will never settle.

`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-9}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-9}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-6}`.`[`back`](#dom-navigation-back){#dom-navigation-back-dev}`(``key`{.variable}`)`\
`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-10}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-10}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-7}`.`[`back`](#dom-navigation-back){#navigation-api-initiating-navigations:dom-navigation-back}`(``key`{.variable}`, { `[`info`](#dom-navigationoptions-info){#navigation-api-initiating-navigations:dom-navigationoptions-info-5}` })`

:   Traverses to the closest previous [session history
    entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry-3}
    which results in this
    [navigable](document-sequences.html#navigable){#navigation-api-initiating-navigations:navigable}
    traversing, i.e., which corresponds to a different
    [`NavigationHistoryEntry`{#navigation-api-initiating-navigations:navigationhistoryentry-4}](#navigationhistoryentry)
    and thus will cause
    [`navigation.currentEntry`{#navigation-api-initiating-navigations:dom-navigation-currententry-3}](#dom-navigation-currententry)
    to change.
    [`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info-6}](#dom-navigationoptions-info)
    can be set to any value; it will populate the
    [`info`{#navigation-api-initiating-navigations:dom-navigateevent-info-4}](#dom-navigateevent-info)
    property of the corresponding
    [`NavigateEvent`{#navigation-api-initiating-navigations:navigateevent-3}](#navigateevent).

    If a traversal to that [session history
    entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry-4}
    is already in progress, then this will return the promises for that
    original traversal, and
    [`info`{#navigation-api-initiating-navigations:dom-navigateevent-info-5}](#dom-navigateevent-info)
    will be ignored.

    The returned promises behave equivalently to those returned by
    [`traverseTo()`{#navigation-api-initiating-navigations:dom-navigation-traverseto-2}](#dom-navigation-traverseto).

`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-11}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-11}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-8}`.`[`forward`](#dom-navigation-forward){#dom-navigation-forward-dev}`(``key`{.variable}`)`\
`{ `[`committed`](#dom-navigationresult-committed){#navigation-api-initiating-navigations:dom-navigationresult-committed-12}`, `[`finished`](#dom-navigationresult-finished){#navigation-api-initiating-navigations:dom-navigationresult-finished-12}` } = `[`navigation`](#dom-navigation){#navigation-api-initiating-navigations:dom-navigation-9}`.`[`forward`](#dom-navigation-forward){#navigation-api-initiating-navigations:dom-navigation-forward}`(``key`{.variable}`, { `[`info`](#dom-navigationoptions-info){#navigation-api-initiating-navigations:dom-navigationoptions-info-7}` })`

:   Traverses to the closest forward [session history
    entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry-5}
    which results in this
    [navigable](document-sequences.html#navigable){#navigation-api-initiating-navigations:navigable-2}
    traversing, i.e., which corresponds to a different
    [`NavigationHistoryEntry`{#navigation-api-initiating-navigations:navigationhistoryentry-5}](#navigationhistoryentry)
    and thus will cause
    [`navigation.currentEntry`{#navigation-api-initiating-navigations:dom-navigation-currententry-4}](#dom-navigation-currententry)
    to change.
    [`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info-8}](#dom-navigationoptions-info)
    can be set to any value; it will populate the
    [`info`{#navigation-api-initiating-navigations:dom-navigateevent-info-6}](#dom-navigateevent-info)
    property of the corresponding
    [`NavigateEvent`{#navigation-api-initiating-navigations:navigateevent-4}](#navigateevent).

    If a traversal to that [session history
    entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry-6}
    is already in progress, then this will return the promises for that
    original traversal, and
    [`info`{#navigation-api-initiating-navigations:dom-navigateevent-info-7}](#dom-navigateevent-info)
    will be ignored.

    The returned promises behave equivalently to those returned by
    [`traverseTo()`{#navigation-api-initiating-navigations:dom-navigation-traverseto-3}](#dom-navigation-traverseto).

The
[`navigate(``url`{.variable}`, ``options`{.variable}`)`]{#dom-navigation-navigate
.dfn dfn-for="Navigation" dfn-type="method"} method steps are:

1.  Let `urlRecord`{.variable} be the result of [parsing a
    URL](urls-and-fetching.html#parse-a-url){#navigation-api-initiating-navigations:parse-a-url}
    given `url`{.variable}, relative to
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this
    x-internal="this"}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#navigation-api-initiating-navigations:relevant-settings-object}.

2.  If `urlRecord`{.variable} is failure, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result}
    for a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#navigation-api-initiating-navigations:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#navigation-api-initiating-navigations:concept-relevant-global}\'s
    [associated
    `Document`](#concept-document-window){#navigation-api-initiating-navigations:concept-document-window}.

4.  If
    `options`{.variable}\[\"[`history`{#navigation-api-initiating-navigations:dom-navigationnavigateoptions-history}](#dom-navigationnavigateoptions-history)\"\]
    is
    \"[`push`{#navigation-api-initiating-navigations:navigationhistorybehavior-push}](browsing-the-web.html#navigationhistorybehavior-push)\",
    and [the navigation must be a
    replace](browsing-the-web.html#the-navigation-must-be-a-replace){#navigation-api-initiating-navigations:the-navigation-must-be-a-replace}
    given `urlRecord`{.variable} and `document`{.variable}, then return
    an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-2}
    for a
    [\"`NotSupportedError`\"](https://webidl.spec.whatwg.org/#notsupportederror){#navigation-api-initiating-navigations:notsupportederror
    x-internal="notsupportederror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  Let `state`{.variable} be
    `options`{.variable}\[\"[`state`{#navigation-api-initiating-navigations:dom-navigationnavigateoptions-state}](#dom-navigationnavigateoptions-state)\"\],
    if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#navigation-api-initiating-navigations:map-exists
    x-internal="map-exists"}; otherwise, undefined.

6.  Let `serializedState`{.variable} be
    [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#navigation-api-initiating-navigations:structuredserializeforstorage}(`state`{.variable}).
    If this throws an exception, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-3}
    for that exception.

    It is important to perform this step early, since serialization can
    invoke web developer code, which in turn might change various things
    we check in later steps.

7.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#navigation-api-initiating-navigations:fully-active},
    then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-4}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

8.  If `document`{.variable}\'s [unload
    counter](document-lifecycle.html#unload-counter){#navigation-api-initiating-navigations:unload-counter}
    is greater than 0, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-5}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

9.  Let `info`{.variable} be
    `options`{.variable}\[\"[`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info-9}](#dom-navigationoptions-info)\"\],
    if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#navigation-api-initiating-navigations:map-exists-2
    x-internal="map-exists"}; otherwise, undefined.

10. Let `apiMethodTracker`{.variable} be the result of [maybe setting
    the upcoming non-traverse API method
    tracker](#maybe-set-the-upcoming-non-traverse-api-method-tracker){#navigation-api-initiating-navigations:maybe-set-the-upcoming-non-traverse-api-method-tracker}
    for
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-3
    x-internal="this"} given `info`{.variable} and
    `serializedState`{.variable}.

11. [Navigate](browsing-the-web.html#navigate){#navigation-api-initiating-navigations:navigate-2}
    `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#navigation-api-initiating-navigations:node-navigable}
    to `urlRecord`{.variable} using `document`{.variable}, with
    *[historyHandling](browsing-the-web.html#navigation-hh)* set to
    `options`{.variable}\[\"[`history`{#navigation-api-initiating-navigations:dom-navigationnavigateoptions-history-2}](#dom-navigationnavigateoptions-history)\"\]
    and
    *[navigationAPIState](browsing-the-web.html#navigation-navigation-api-state)*
    set to `serializedState`{.variable}.

    Unlike
    [`location.assign()`{#navigation-api-initiating-navigations:dom-location-assign}](#dom-location-assign)
    and friends, which are exposed across
    [origin-domain](browsers.html#same-origin-domain){#navigation-api-initiating-navigations:same-origin-domain}
    boundaries,
    [`navigation.navigate()`{#navigation-api-initiating-navigations:dom-navigation-navigate-2}](#dom-navigation-navigate)
    can only be accessed by code with direct synchronous access to the
    [`window.navigation`{#navigation-api-initiating-navigations:dom-navigation-10}](#dom-navigation)
    property. Thus, we avoid the complications about attributing the
    source document of the navigation, and we don\'t need to deal with
    the [allowed by sandboxing to
    navigate](browsing-the-web.html#allowed-to-navigate){#navigation-api-initiating-navigations:allowed-to-navigate}
    check and its acccompanying
    *[exceptionsEnabled](browsing-the-web.html#exceptions-enabled)*
    flag. We just treat all navigations as if they come from the
    [`Document`{#navigation-api-initiating-navigations:document}](dom.html#document)
    corresponding to this
    [`Navigation`{#navigation-api-initiating-navigations:navigation}](#navigation)
    object itself (i.e., `document`{.variable}).

12. If
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-4
    x-internal="this"}\'s [upcoming non-traverse API method
    tracker](#upcoming-non-traverse-api-method-tracker){#navigation-api-initiating-navigations:upcoming-non-traverse-api-method-tracker}
    is `apiMethodTracker`{.variable}, then:

    If the [upcoming non-traverse API method
    tracker](#upcoming-non-traverse-api-method-tracker){#navigation-api-initiating-navigations:upcoming-non-traverse-api-method-tracker-2}
    is still `apiMethodTracker`{.variable}, this means that the
    [navigate](browsing-the-web.html#navigate){#navigation-api-initiating-navigations:navigate-3}
    algorithm bailed out before ever getting to the [inner `navigate`
    event firing
    algorithm](#inner-navigate-event-firing-algorithm){#navigation-api-initiating-navigations:inner-navigate-event-firing-algorithm}
    which would [promote that upcoming API method tracker to
    ongoing](#promote-an-upcoming-api-method-tracker-to-ongoing){#navigation-api-initiating-navigations:promote-an-upcoming-api-method-tracker-to-ongoing}.

    1.  Set
        [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-5
        x-internal="this"}\'s [upcoming non-traverse API method
        tracker](#upcoming-non-traverse-api-method-tracker){#navigation-api-initiating-navigations:upcoming-non-traverse-api-method-tracker-3}
        to null.

    2.  Return an [early error
        result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-6}
        for an
        [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#navigation-api-initiating-navigations:aborterror-2
        x-internal="aborterror"}
        [`DOMException`{#navigation-api-initiating-navigations:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

13. Return a [navigation API method tracker-derived
    result](#navigation-api-method-tracker-derived-result){#navigation-api-initiating-navigations:navigation-api-method-tracker-derived-result}
    for `apiMethodTracker`{.variable}.

The [`reload(``options`{.variable}`)`]{#dom-navigation-reload .dfn
dfn-for="Navigation" dfn-type="method"} method steps are:

1.  Let `document`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-6
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#navigation-api-initiating-navigations:concept-relevant-global-2}\'s
    [associated
    `Document`](#concept-document-window){#navigation-api-initiating-navigations:concept-document-window-2}.

2.  Let `serializedState`{.variable} be
    [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#navigation-api-initiating-navigations:structuredserializeforstorage-2}(undefined).

3.  If
    `options`{.variable}\[\"[`state`{#navigation-api-initiating-navigations:dom-navigationreloadoptions-state}](#dom-navigationreloadoptions-state)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#navigation-api-initiating-navigations:map-exists-3
    x-internal="map-exists"}, then set `serializedState`{.variable} to
    [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#navigation-api-initiating-navigations:structuredserializeforstorage-3}(`options`{.variable}\[\"[`state`{#navigation-api-initiating-navigations:dom-navigationreloadoptions-state-2}](#dom-navigationreloadoptions-state)\"\]).
    If this throws an exception, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-7}
    for that exception.

    It is important to perform this step early, since serialization can
    invoke web developer code, which in turn might change various things
    we check in later steps.

4.  Otherwise:

    1.  Let `current`{.variable} be the [current
        entry](#navigation-current-entry){#navigation-api-initiating-navigations:navigation-current-entry}
        of
        [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-7
        x-internal="this"}.

    2.  If `current`{.variable} is not null, then set
        `serializedState`{.variable} to `current`{.variable}\'s [session
        history
        entry](#nhe-she){#navigation-api-initiating-navigations:nhe-she}\'s
        [navigation API
        state](browsing-the-web.html#she-navigation-api-state){#navigation-api-initiating-navigations:she-navigation-api-state}.

5.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#navigation-api-initiating-navigations:fully-active-2},
    then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-8}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  If `document`{.variable}\'s [unload
    counter](document-lifecycle.html#unload-counter){#navigation-api-initiating-navigations:unload-counter-2}
    is greater than 0, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-9}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-5
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

7.  Let `info`{.variable} be
    `options`{.variable}\[\"[`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info-10}](#dom-navigationoptions-info)\"\],
    if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#navigation-api-initiating-navigations:map-exists-4
    x-internal="map-exists"}; otherwise, undefined.

8.  Let `apiMethodTracker`{.variable} be the result of [maybe setting
    the upcoming non-traverse API method
    tracker](#maybe-set-the-upcoming-non-traverse-api-method-tracker){#navigation-api-initiating-navigations:maybe-set-the-upcoming-non-traverse-api-method-tracker-2}
    for
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-8
    x-internal="this"} given `info`{.variable} and
    `serializedState`{.variable}.

9.  [Reload](browsing-the-web.html#reload){#navigation-api-initiating-navigations:reload-2}
    `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#navigation-api-initiating-navigations:node-navigable-2}
    with
    *[navigationAPIState](browsing-the-web.html#reload-navigation-api-state)*
    set to `serializedState`{.variable}.

10. Return a [navigation API method tracker-derived
    result](#navigation-api-method-tracker-derived-result){#navigation-api-initiating-navigations:navigation-api-method-tracker-derived-result-2}
    for `apiMethodTracker`{.variable}.

The
[`traverseTo(``key`{.variable}`, ``options`{.variable}`)`]{#dom-navigation-traverseto
.dfn dfn-for="Navigation" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-9
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-initiating-navigations:navigation-current-entry-index}
    is −1, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-10}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-6
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-10}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-10
    x-internal="this"}\'s [entry
    list](#navigation-entry-list){#navigation-api-initiating-navigations:navigation-entry-list}
    does not
    [contain](https://infra.spec.whatwg.org/#list-contain){#navigation-api-initiating-navigations:list-contains
    x-internal="list-contains"} a
    [`NavigationHistoryEntry`{#navigation-api-initiating-navigations:navigationhistoryentry-6}](#navigationhistoryentry)
    whose [session history
    entry](#nhe-she){#navigation-api-initiating-navigations:nhe-she-2}\'s
    [navigation API
    key](browsing-the-web.html#she-navigation-api-key){#navigation-api-initiating-navigations:she-navigation-api-key}
    equals `key`{.variable}, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-11}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-7
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-11}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  Return the result of [performing a navigation API
    traversal](#performing-a-navigation-api-traversal){#navigation-api-initiating-navigations:performing-a-navigation-api-traversal}
    given
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-11
    x-internal="this"}, `key`{.variable}, and `options`{.variable}.

The [`back(``options`{.variable}`)`]{#dom-navigation-back .dfn
dfn-for="Navigation" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-12
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-initiating-navigations:navigation-current-entry-index-2}
    is −1 or 0, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-12}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-8
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-12}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `key`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-13
    x-internal="this"}\'s [entry
    list](#navigation-entry-list){#navigation-api-initiating-navigations:navigation-entry-list-2}\[[this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-14
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-initiating-navigations:navigation-current-entry-index-3}
    − 1\]\'s [session history
    entry](#nhe-she){#navigation-api-initiating-navigations:nhe-she-3}\'s
    [navigation API
    key](browsing-the-web.html#she-navigation-api-key){#navigation-api-initiating-navigations:she-navigation-api-key-2}.

3.  Return the result of [performing a navigation API
    traversal](#performing-a-navigation-api-traversal){#navigation-api-initiating-navigations:performing-a-navigation-api-traversal-2}
    given
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-15
    x-internal="this"}, `key`{.variable}, and `options`{.variable}.

The [`forward(``options`{.variable}`)`]{#dom-navigation-forward .dfn
dfn-for="Navigation" dfn-type="method"} method steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-16
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-initiating-navigations:navigation-current-entry-index-4}
    is −1 or is equal to
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-17
    x-internal="this"}\'s [entry
    list](#navigation-entry-list){#navigation-api-initiating-navigations:navigation-entry-list-3}\'s
    [size](https://infra.spec.whatwg.org/#list-size){#navigation-api-initiating-navigations:list-size
    x-internal="list-size"} − 1, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-13}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-9
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-13}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  Let `key`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-18
    x-internal="this"}\'s [entry
    list](#navigation-entry-list){#navigation-api-initiating-navigations:navigation-entry-list-4}\[[this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-19
    x-internal="this"}\'s [current entry
    index](#navigation-current-entry-index){#navigation-api-initiating-navigations:navigation-current-entry-index-5} +
    1\]\'s [session history
    entry](#nhe-she){#navigation-api-initiating-navigations:nhe-she-4}\'s
    [navigation API
    key](browsing-the-web.html#she-navigation-api-key){#navigation-api-initiating-navigations:she-navigation-api-key-3}.

3.  Return the result of [performing a navigation API
    traversal](#performing-a-navigation-api-traversal){#navigation-api-initiating-navigations:performing-a-navigation-api-traversal-3}
    given
    [this](https://webidl.spec.whatwg.org/#this){#navigation-api-initiating-navigations:this-20
    x-internal="this"}, `key`{.variable}, and `options`{.variable}.

To [perform a navigation API
traversal]{#performing-a-navigation-api-traversal .dfn} given a
[`Navigation`{#navigation-api-initiating-navigations:navigation-2}](#navigation)
`navigation`{.variable}, a string `key`{.variable}, and a
[`NavigationOptions`{#navigation-api-initiating-navigations:navigationoptions}](#navigationoptions)
`options`{.variable}:

1.  Let `document`{.variable} be `navigation`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#navigation-api-initiating-navigations:concept-relevant-global-3}\'s
    [associated
    `Document`](#concept-document-window){#navigation-api-initiating-navigations:concept-document-window-3}.

2.  If `document`{.variable} is not [fully
    active](document-sequences.html#fully-active){#navigation-api-initiating-navigations:fully-active-3},
    then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-14}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-10
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-14}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `document`{.variable}\'s [unload
    counter](document-lifecycle.html#unload-counter){#navigation-api-initiating-navigations:unload-counter-3}
    is greater than 0, then return an [early error
    result](#navigation-api-early-error-result){#navigation-api-initiating-navigations:navigation-api-early-error-result-15}
    for an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-11
    x-internal="invalidstateerror"}
    [`DOMException`{#navigation-api-initiating-navigations:domexception-15}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  Let `current`{.variable} be the [current
    entry](#navigation-current-entry){#navigation-api-initiating-navigations:navigation-current-entry-2}
    of `navigation`{.variable}.

5.  If `key`{.variable} equals `current`{.variable}\'s [session history
    entry](#nhe-she){#navigation-api-initiating-navigations:nhe-she-5}\'s
    [navigation API
    key](browsing-the-web.html#she-navigation-api-key){#navigation-api-initiating-navigations:she-navigation-api-key-4},
    then return «\[
    \"[`committed`{#navigation-api-initiating-navigations:dom-navigationresult-committed-13}](#dom-navigationresult-committed)\"
    → [a promise resolved
    with](https://webidl.spec.whatwg.org/#a-promise-resolved-with){#navigation-api-initiating-navigations:a-promise-resolved-with
    x-internal="a-promise-resolved-with"} `current`{.variable},
    \"[`finished`{#navigation-api-initiating-navigations:dom-navigationresult-finished-13}](#dom-navigationresult-finished)\"
    → [a promise resolved
    with](https://webidl.spec.whatwg.org/#a-promise-resolved-with){#navigation-api-initiating-navigations:a-promise-resolved-with-2
    x-internal="a-promise-resolved-with"} `current`{.variable} \]».

6.  If `navigation`{.variable}\'s [upcoming traverse API method
    trackers](#upcoming-traverse-api-method-trackers){#navigation-api-initiating-navigations:upcoming-traverse-api-method-trackers}\[`key`{.variable}\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#navigation-api-initiating-navigations:map-exists-5
    x-internal="map-exists"}, then return a [navigation API method
    tracker-derived
    result](#navigation-api-method-tracker-derived-result){#navigation-api-initiating-navigations:navigation-api-method-tracker-derived-result-3}
    for `navigation`{.variable}\'s [upcoming traverse API method
    trackers](#upcoming-traverse-api-method-trackers){#navigation-api-initiating-navigations:upcoming-traverse-api-method-trackers-2}\[`key`{.variable}\].

7.  Let `info`{.variable} be
    `options`{.variable}\[\"[`info`{#navigation-api-initiating-navigations:dom-navigationoptions-info-11}](#dom-navigationoptions-info)\"\],
    if it
    [exists](https://infra.spec.whatwg.org/#map-exists){#navigation-api-initiating-navigations:map-exists-6
    x-internal="map-exists"}; otherwise, undefined.

8.  Let `apiMethodTracker`{.variable} be the result of [adding an
    upcoming traverse API method
    tracker](#add-an-upcoming-traverse-api-method-tracker){#navigation-api-initiating-navigations:add-an-upcoming-traverse-api-method-tracker}
    for `navigation`{.variable} given `key`{.variable} and
    `info`{.variable}.

9.  Let `navigable`{.variable} be `document`{.variable}\'s [node
    navigable](document-sequences.html#node-navigable){#navigation-api-initiating-navigations:node-navigable-3}.

10. Let `traversable`{.variable} be `navigable`{.variable}\'s
    [traversable
    navigable](document-sequences.html#nav-traversable){#navigation-api-initiating-navigations:nav-traversable}.

11. Let `sourceSnapshotParams`{.variable} be the result of [snapshotting
    source snapshot
    params](browsing-the-web.html#snapshotting-source-snapshot-params){#navigation-api-initiating-navigations:snapshotting-source-snapshot-params}
    given `document`{.variable}.

12. [Append the following session history traversal
    steps](browsing-the-web.html#tn-append-session-history-traversal-steps){#navigation-api-initiating-navigations:tn-append-session-history-traversal-steps}
    to `traversable`{.variable}:

    1.  Let `navigableSHEs`{.variable} be the result of [getting session
        history
        entries](browsing-the-web.html#getting-session-history-entries){#navigation-api-initiating-navigations:getting-session-history-entries}
        given `navigable`{.variable}.

    2.  Let `targetSHE`{.variable} be the [session history
        entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry-7}
        in `navigableSHEs`{.variable} whose [navigation API
        key](browsing-the-web.html#she-navigation-api-key){#navigation-api-initiating-navigations:she-navigation-api-key-5}
        is `key`{.variable}. If no such entry exists, then:

        1.  [Queue a global
            task](webappapis.html#queue-a-global-task){#navigation-api-initiating-navigations:queue-a-global-task}
            on the [navigation and traversal task
            source](webappapis.html#navigation-and-traversal-task-source){#navigation-api-initiating-navigations:navigation-and-traversal-task-source}
            given `navigation`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#navigation-api-initiating-navigations:concept-relevant-global-4}
            to [reject the finished
            promise](#reject-the-finished-promise){#navigation-api-initiating-navigations:reject-the-finished-promise}
            for `apiMethodTracker`{.variable} with an
            [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#navigation-api-initiating-navigations:invalidstateerror-12
            x-internal="invalidstateerror"}
            [`DOMException`{#navigation-api-initiating-navigations:domexception-16}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

        2.  Abort these steps.

        This path is taken if `navigation`{.variable}\'s [entry
        list](#navigation-entry-list){#navigation-api-initiating-navigations:navigation-entry-list-5}
        was outdated compared to `navigableSHEs`{.variable}, which can
        occur for brief periods while all the relevant threads and
        processes are being synchronized in reaction to a history
        change.

    3.  If `targetSHE`{.variable} is `navigable`{.variable}\'s [active
        session history
        entry](document-sequences.html#nav-active-history-entry){#navigation-api-initiating-navigations:nav-active-history-entry},
        then abort these steps.

        This can occur if a previously
        [queued](browsing-the-web.html#tn-append-session-history-traversal-steps){#navigation-api-initiating-navigations:tn-append-session-history-traversal-steps-2}
        traversal already took us to this [session history
        entry](browsing-the-web.html#session-history-entry){#navigation-api-initiating-navigations:session-history-entry-8}.
        In that case the previous traversal will have dealt with
        `apiMethodTracker`{.variable} already.

    4.  Let `result`{.variable} be the result of [applying the traverse
        history
        step](browsing-the-web.html#apply-the-traverse-history-step){#navigation-api-initiating-navigations:apply-the-traverse-history-step-2}
        given by `targetSHE`{.variable}\'s
        [step](browsing-the-web.html#she-step){#navigation-api-initiating-navigations:she-step}
        to `traversable`{.variable}, given
        `sourceSnapshotParams`{.variable}, `navigable`{.variable}, and
        \"[`none`{#navigation-api-initiating-navigations:uni-none}](browsing-the-web.html#uni-none)\".

    5.  If `result`{.variable} is \"`canceled-by-beforeunload`\", then
        [queue a global
        task](webappapis.html#queue-a-global-task){#navigation-api-initiating-navigations:queue-a-global-task-2}
        on the [navigation and traversal task
        source](webappapis.html#navigation-and-traversal-task-source){#navigation-api-initiating-navigations:navigation-and-traversal-task-source-2}
        given `navigation`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#navigation-api-initiating-navigations:concept-relevant-global-5}
        to [reject the finished
        promise](#reject-the-finished-promise){#navigation-api-initiating-navigations:reject-the-finished-promise-2}
        for `apiMethodTracker`{.variable} with a new
        [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#navigation-api-initiating-navigations:aborterror-3
        x-internal="aborterror"}
        [`DOMException`{#navigation-api-initiating-navigations:domexception-17}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
        created in `navigation`{.variable}\'s [relevant
        realm](webappapis.html#concept-relevant-realm){#navigation-api-initiating-navigations:concept-relevant-realm}.

    6.  If `result`{.variable} is \"`initiator-disallowed`\", then
        [queue a global
        task](webappapis.html#queue-a-global-task){#navigation-api-initiating-navigations:queue-a-global-task-3}
        on the [navigation and traversal task
        source](webappapis.html#navigation-and-traversal-task-source){#navigation-api-initiating-navigations:navigation-and-traversal-task-source-3}
        given `navigation`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#navigation-api-initiating-navigations:concept-relevant-global-6}
        to [reject the finished
        promise](#reject-the-finished-promise){#navigation-api-initiating-navigations:reject-the-finished-promise-3}
        for `apiMethodTracker`{.variable} with a new
        [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#navigation-api-initiating-navigations:securityerror
        x-internal="securityerror"}
        [`DOMException`{#navigation-api-initiating-navigations:domexception-18}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
        created in `navigation`{.variable}\'s [relevant
        realm](webappapis.html#concept-relevant-realm){#navigation-api-initiating-navigations:concept-relevant-realm-2}.

        ::: note
        When `result`{.variable} is \"`canceled-by-beforeunload`\" or
        \"`initiator-disallowed`\", the
        [`navigate`{#navigation-api-initiating-navigations:event-navigate-2}](indices.html#event-navigate)
        event was never fired, [aborting the ongoing
        navigation](#abort-the-ongoing-navigation){#navigation-api-initiating-navigations:abort-the-ongoing-navigation}
        would not be correct; it would result in a
        [`navigateerror`{#navigation-api-initiating-navigations:event-navigateerror}](indices.html#event-navigateerror)
        event without a preceding
        [`navigate`{#navigation-api-initiating-navigations:event-navigate-3}](indices.html#event-navigate)
        event.

        In the \"`canceled-by-navigate`\" case,
        [`navigate`{#navigation-api-initiating-navigations:event-navigate-4}](indices.html#event-navigate)
        *is* fired, but the [inner `navigate` event firing
        algorithm](#inner-navigate-event-firing-algorithm){#navigation-api-initiating-navigations:inner-navigate-event-firing-algorithm-2}
        will take care of [aborting the ongoing
        navigation](#abort-the-ongoing-navigation){#navigation-api-initiating-navigations:abort-the-ongoing-navigation-2}.
        :::

13. Return a [navigation API method tracker-derived
    result](#navigation-api-method-tracker-derived-result){#navigation-api-initiating-navigations:navigation-api-method-tracker-derived-result-4}
    for `apiMethodTracker`{.variable}.

An [early error result]{#navigation-api-early-error-result .dfn} for an
exception `e`{.variable} is a
[`NavigationResult`{#navigation-api-initiating-navigations:navigationresult}](#navigationresult)
dictionary instance given by «\[
\"[`committed`{#navigation-api-initiating-navigations:dom-navigationresult-committed-14}](#dom-navigationresult-committed)\"
→ [a promise rejected
with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#navigation-api-initiating-navigations:a-promise-rejected-with
x-internal="a-promise-rejected-with"} `e`{.variable},
\"[`finished`{#navigation-api-initiating-navigations:dom-navigationresult-finished-14}](#dom-navigationresult-finished)\"
→ [a promise rejected
with](https://webidl.spec.whatwg.org/#a-promise-rejected-with){#navigation-api-initiating-navigations:a-promise-rejected-with-2
x-internal="a-promise-rejected-with"} `e`{.variable} \]».

A [navigation API method tracker-derived
result]{#navigation-api-method-tracker-derived-result .dfn} for a
[navigation API method
tracker](#navigation-api-method-tracker){#navigation-api-initiating-navigations:navigation-api-method-tracker}
is a
[`NavigationResult`{#navigation-api-initiating-navigations:navigationresult-2}](#navigationresult)
dictionary instance given by «\[
\"[`committed`{#navigation-api-initiating-navigations:dom-navigationresult-committed-15}](#dom-navigationresult-committed)\"
→ `apiMethodTracker`{.variable}\'s [committed
promise](#navigation-api-method-tracker-committed){#navigation-api-initiating-navigations:navigation-api-method-tracker-committed},
\"[`finished`{#navigation-api-initiating-navigations:dom-navigationresult-finished-15}](#dom-navigationresult-finished)\"
→ `apiMethodTracker`{.variable}\'s [finished
promise](#navigation-api-method-tracker-finished){#navigation-api-initiating-navigations:navigation-api-method-tracker-finished}
\]».

##### [7.2.6.8]{.secno} Ongoing navigation tracking[](#ongoing-navigation-tracking){.self-link}

During any given navigation (in the [broad sense of the
word](#navigationtype){#ongoing-navigation-tracking:navigationtype}),
the [`Navigation`{#ongoing-navigation-tracking:navigation}](#navigation)
object needs to keep track of the following:

For all navigations

State

Duration

Explanation

The
[`NavigateEvent`{#ongoing-navigation-tracking:navigateevent}](#navigateevent)

For the duration of event firing

So that if the navigation is canceled while the event is firing, we can
[cancel](https://dom.spec.whatwg.org/#canceled-flag){#ongoing-navigation-tracking:canceled-flag
x-internal="canceled-flag"} the event

The event\'s [abort
controller](#concept-navigateevent-abort-controller){#ongoing-navigation-tracking:concept-navigateevent-abort-controller}

Until all promises returned from handlers passed to
[`intercept()`{#ongoing-navigation-tracking:dom-navigateevent-intercept}](#dom-navigateevent-intercept)
have settled

So that if the navigation is canceled, we can [signal
abort](https://dom.spec.whatwg.org/#abortcontroller-signal-abort){#ongoing-navigation-tracking:signal-abort
x-internal="signal-abort"}

Whether a new element was
[focused](interaction.html#focusing-steps){#ongoing-navigation-tracking:focusing-steps}

Until all promises returned from handlers passed to
[`intercept()`{#ongoing-navigation-tracking:dom-navigateevent-intercept-2}](#dom-navigateevent-intercept)
have settled

So that if one was, focus is not
[reset](#potentially-reset-the-focus){#ongoing-navigation-tracking:potentially-reset-the-focus}

The
[`NavigationHistoryEntry`{#ongoing-navigation-tracking:navigationhistoryentry}](#navigationhistoryentry)
being navigated to

From when it is determined, until all promises returned from handlers
passed to
[`intercept()`{#ongoing-navigation-tracking:dom-navigateevent-intercept-3}](#dom-navigateevent-intercept)
have settled

So that we know what to resolve any
[`committed`{#ongoing-navigation-tracking:dom-navigationresult-committed}](#dom-navigationresult-committed)
and
[`finished`{#ongoing-navigation-tracking:dom-navigationresult-finished}](#dom-navigationresult-finished)
promises with

Any
[`finished`{#ongoing-navigation-tracking:dom-navigationresult-finished-2}](#dom-navigationresult-finished)
promise that was returned

Until all promises returned from handlers passed to
[`intercept()`{#ongoing-navigation-tracking:dom-navigateevent-intercept-4}](#dom-navigateevent-intercept)
have settled

So that we can resolve or reject it appropriately

For
non-\"[`traverse`{#ongoing-navigation-tracking:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\"
navigations

State

Duration

Explanation

Any
[`state`{#ongoing-navigation-tracking:dom-navigationnavigateoptions-state}](#dom-navigationnavigateoptions-state)

For the duration of event firing

So that we can update the current entry\'s [navigation API
state](browsing-the-web.html#she-navigation-api-state){#ongoing-navigation-tracking:she-navigation-api-state}
if the event finishes firing without being
[canceled](https://dom.spec.whatwg.org/#canceled-flag){#ongoing-navigation-tracking:canceled-flag-2
x-internal="canceled-flag"}

For
\"[`traverse`{#ongoing-navigation-tracking:dom-navigationtype-traverse-2}](#dom-navigationtype-traverse)\"
navigations

State

Duration

Explanation

Any
[`info`{#ongoing-navigation-tracking:dom-navigationoptions-info}](#dom-navigationoptions-info)

Until the task is queued to fire the
[`navigate`{#ongoing-navigation-tracking:event-navigate}](indices.html#event-navigate)
event

So that we can use it to fire the
[`navigate`{#ongoing-navigation-tracking:event-navigate-2}](indices.html#event-navigate)
after the trip through the [session history traversal
queue](document-sequences.html#tn-session-history-traversal-queue){#ongoing-navigation-tracking:tn-session-history-traversal-queue}.

Any
[`committed`{#ongoing-navigation-tracking:dom-navigationresult-committed-2}](#dom-navigationresult-committed)
promise that was returned

Until the session history is updated (inside that same task)

So that we can resolve or reject it appropriately

Whether
[`intercept()`{#ongoing-navigation-tracking:dom-navigateevent-intercept-5}](#dom-navigateevent-intercept)
was called

Until the session history is updated (inside that same task)

So that we can suppress the normal scroll restoration logic in favor of
the behavior given by the
[`scroll`{#ongoing-navigation-tracking:dom-navigationinterceptoptions-scroll}](#dom-navigationinterceptoptions-scroll)
option

We also cannot assume there is only a single navigation requested at any
given time, due to web developer code such as:

``` js
const p1 = navigation.navigate(url1).finished;
const p2 = navigation.navigate(url2).finished;
```

That is, in this scenario, we need to ensure that while navigating to
`url2`, we still have the promise `p1` around so that we can reject it.
We can\'t just get rid of any ongoing navigation promises the moment the
second call to
[`navigate()`{#ongoing-navigation-tracking:dom-navigation-navigate}](#dom-navigation-navigate)
happens.

We end up accomplishing all this by associating the following with each
[`Navigation`{#ongoing-navigation-tracking:navigation-2}](#navigation):

- [Ongoing `navigate` event]{#ongoing-navigate-event .dfn}, a
  [`NavigateEvent`{#ongoing-navigation-tracking:navigateevent-2}](#navigateevent)
  or null, initially null.

- [Focus changed during ongoing
  navigation]{#focus-changed-during-ongoing-navigation .dfn}, a boolean,
  initially false.

- [Suppress normal scroll restoration during ongoing
  navigation]{#suppress-normal-scroll-restoration-during-ongoing-navigation
  .dfn}, a boolean, initially false.

- [Ongoing API method tracker]{#ongoing-api-method-tracker .dfn}, a
  [navigation API method
  tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker}
  or null, initially null.

- [Upcoming non-traverse API method
  tracker]{#upcoming-non-traverse-api-method-tracker .dfn}, a
  [navigation API method
  tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-2}
  or null, initially null.

- [Upcoming traverse API method
  trackers]{#upcoming-traverse-api-method-trackers .dfn}, an [ordered
  map](https://infra.spec.whatwg.org/#ordered-map){#ongoing-navigation-tracking:ordered-map
  x-internal="ordered-map"} from strings to [navigation API method
  trackers](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-3},
  initially empty.

The state here that is not stored in [navigation API method
trackers](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-4}
is state which needs to be tracked even for navigations that are not
initiated via navigation API methods.

A [navigation API method tracker]{#navigation-api-method-tracker .dfn}
is a
[struct](https://infra.spec.whatwg.org/#struct){#ongoing-navigation-tracking:struct
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#ongoing-navigation-tracking:struct-item
x-internal="struct-item"}:

- A [navigation object]{#navigation-api-method-tracker-navigation .dfn},
  a
  [`Navigation`{#ongoing-navigation-tracking:navigation-3}](#navigation)

- A [key]{#navigation-api-method-tracker-key .dfn}, a string or null

- An [info]{#navigation-api-method-tracker-info .dfn}, a JavaScript
  value

- A [serialized state]{#navigation-api-method-tracker-state .dfn}, a
  [serialized
  state](browsing-the-web.html#serialized-state){#ongoing-navigation-tracking:serialized-state}
  or null

- A [committed-to
  entry]{#navigation-api-method-tracker-committed-to-entry .dfn}, a
  [`NavigationHistoryEntry`{#ongoing-navigation-tracking:navigationhistoryentry-2}](#navigationhistoryentry)
  or null

- A [committed promise]{#navigation-api-method-tracker-committed .dfn},
  a promise

- A [finished promise]{#navigation-api-method-tracker-finished .dfn}, a
  promise

All this state is then managed via the following algorithms.

To [maybe set the upcoming non-traverse API method
tracker]{#maybe-set-the-upcoming-non-traverse-api-method-tracker .dfn}
given a
[`Navigation`{#ongoing-navigation-tracking:navigation-4}](#navigation)
`navigation`{.variable}, a JavaScript value `info`{.variable}, and a
[serialized
state](browsing-the-web.html#serialized-state){#ongoing-navigation-tracking:serialized-state-2}-or-null
`serializedState`{.variable}:

1.  Let `committedPromise`{.variable} and `finishedPromise`{.variable}
    be new promises created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#ongoing-navigation-tracking:concept-relevant-realm}.

2.  [Mark as
    handled](https://webidl.spec.whatwg.org/#mark-a-promise-as-handled){#ongoing-navigation-tracking:mark-as-handled
    x-internal="mark-as-handled"} `finishedPromise`{.variable}.

    ::: {#note-mark-as-handled-navigation-api-finished .note}
    The web developer doesn't necessarily care about
    `finishedPromise`{.variable} being rejected:

    - They might only care about `committedPromise`{.variable}.

    - They could be doing multiple synchronous navigations within the
      same task, in which case all but the last will be aborted (causing
      their `finishedPromise`{.variable} to reject). This could be an
      application bug, but also could just be an emergent feature of
      disparate parts of the application overriding each others\'
      actions.

    - They might prefer to listen to other transition-failure signals
      instead of `finishedPromise`{.variable}, e.g., the
      [`navigateerror`{#ongoing-navigation-tracking:event-navigateerror}](indices.html#event-navigateerror)
      event, or the
      [`navigation.transition.finished`{#ongoing-navigation-tracking:dom-navigationtransition-finished}](#dom-navigationtransition-finished)
      promise.

    As such, we mark it as handled to ensure that it never triggers
    [`unhandledrejection`{#ongoing-navigation-tracking:event-unhandledrejection}](indices.html#event-unhandledrejection)
    events.
    :::

3.  Let `apiMethodTracker`{.variable} be a new [navigation API method
    tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-5}
    with:

    [navigation object](#navigation-api-method-tracker-navigation){#ongoing-navigation-tracking:navigation-api-method-tracker-navigation}
    :   `navigation`{.variable}

    [key](#navigation-api-method-tracker-key){#ongoing-navigation-tracking:navigation-api-method-tracker-key}
    :   null

    [info](#navigation-api-method-tracker-info){#ongoing-navigation-tracking:navigation-api-method-tracker-info}
    :   `info`{.variable}

    [serialized state](#navigation-api-method-tracker-state){#ongoing-navigation-tracking:navigation-api-method-tracker-state}
    :   `serializedState`{.variable}

    [committed-to entry](#navigation-api-method-tracker-committed-to-entry){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-to-entry}
    :   null

    [committed promise](#navigation-api-method-tracker-committed){#ongoing-navigation-tracking:navigation-api-method-tracker-committed}
    :   `committedPromise`{.variable}

    [finished promise](#navigation-api-method-tracker-finished){#ongoing-navigation-tracking:navigation-api-method-tracker-finished}
    :   `finishedPromise`{.variable}

4.  [Assert](https://infra.spec.whatwg.org/#assert){#ongoing-navigation-tracking:assert
    x-internal="assert"}: `navigation`{.variable}\'s [upcoming
    non-traverse API method
    tracker](#upcoming-non-traverse-api-method-tracker){#ongoing-navigation-tracking:upcoming-non-traverse-api-method-tracker}
    is null.

5.  ::: {#dont-always-set-upcoming-non-traverse-api-method-tracker}
    If `navigation`{.variable} does not [have entries and events
    disabled](#has-entries-and-events-disabled){#ongoing-navigation-tracking:has-entries-and-events-disabled},
    then set `navigation`{.variable}\'s [upcoming non-traverse API
    method
    tracker](#upcoming-non-traverse-api-method-tracker){#ongoing-navigation-tracking:upcoming-non-traverse-api-method-tracker-2}
    to `apiMethodTracker`{.variable}.

    If `navigation`{.variable} [has entries and events
    disabled](#has-entries-and-events-disabled){#ongoing-navigation-tracking:has-entries-and-events-disabled-2},
    then `committedPromise`{.variable} and `finishedPromise`{.variable}
    will never fulfill (since we never create a
    [`NavigationHistoryEntry`{#ongoing-navigation-tracking:navigationhistoryentry-3}](#navigationhistoryentry)
    object for such
    [`Document`{#ongoing-navigation-tracking:document}](dom.html#document)s,
    and so we have nothing to resolve them with); there is no
    [`NavigationHistoryEntry`{#ongoing-navigation-tracking:navigationhistoryentry-4}](#navigationhistoryentry)
    to apply `serializedState`{.variable} to; and there is no
    [`navigate`{#ongoing-navigation-tracking:event-navigate-3}](indices.html#event-navigate)
    event to include `info`{.variable} with. So, we don\'t need to track
    this API method call after all.
    :::

6.  Return `apiMethodTracker`{.variable}.

To [add an upcoming traverse API method
tracker]{#add-an-upcoming-traverse-api-method-tracker .dfn} given a
[`Navigation`{#ongoing-navigation-tracking:navigation-5}](#navigation)
`navigation`{.variable}, a string `destinationKey`{.variable}, and a
JavaScript value `info`{.variable}:

1.  Let `committedPromise`{.variable} and `finishedPromise`{.variable}
    be new promises created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#ongoing-navigation-tracking:concept-relevant-realm-2}.

2.  [Mark as
    handled](https://webidl.spec.whatwg.org/#mark-a-promise-as-handled){#ongoing-navigation-tracking:mark-as-handled-2
    x-internal="mark-as-handled"} `finishedPromise`{.variable}.

    See the [previous
    discussion](#note-mark-as-handled-navigation-api-finished) about why
    this is done.

3.  Let `apiMethodTracker`{.variable} be a new [navigation API method
    tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-6}
    with:

    [navigation object](#navigation-api-method-tracker-navigation){#ongoing-navigation-tracking:navigation-api-method-tracker-navigation-2}
    :   `navigation`{.variable}

    [key](#navigation-api-method-tracker-key){#ongoing-navigation-tracking:navigation-api-method-tracker-key-2}
    :   `destinationKey`{.variable}

    [info](#navigation-api-method-tracker-info){#ongoing-navigation-tracking:navigation-api-method-tracker-info-2}
    :   `info`{.variable}

    [serialized state](#navigation-api-method-tracker-state){#ongoing-navigation-tracking:navigation-api-method-tracker-state-2}
    :   null

    [committed-to entry](#navigation-api-method-tracker-committed-to-entry){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-to-entry-2}
    :   null

    [committed promise](#navigation-api-method-tracker-committed){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-2}
    :   `committedPromise`{.variable}

    [finished promise](#navigation-api-method-tracker-finished){#ongoing-navigation-tracking:navigation-api-method-tracker-finished-2}
    :   `finishedPromise`{.variable}

4.  Set `navigation`{.variable}\'s [upcoming traverse API method
    trackers](#upcoming-traverse-api-method-trackers){#ongoing-navigation-tracking:upcoming-traverse-api-method-trackers}\[`destinationKey`{.variable}\]
    to `apiMethodTracker`{.variable}.

5.  Return `apiMethodTracker`{.variable}.

To [promote an upcoming API method tracker to
ongoing]{#promote-an-upcoming-api-method-tracker-to-ongoing .dfn} given
a [`Navigation`{#ongoing-navigation-tracking:navigation-6}](#navigation)
`navigation`{.variable} and a string-or-null
`destinationKey`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ongoing-navigation-tracking:assert-2
    x-internal="assert"}: `navigation`{.variable}\'s [ongoing API method
    tracker](#ongoing-api-method-tracker){#ongoing-navigation-tracking:ongoing-api-method-tracker}
    is null.

2.  If `destinationKey`{.variable} is not null, then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#ongoing-navigation-tracking:assert-3
        x-internal="assert"}: `navigation`{.variable}\'s [upcoming
        non-traverse API method
        tracker](#upcoming-non-traverse-api-method-tracker){#ongoing-navigation-tracking:upcoming-non-traverse-api-method-tracker-3}
        is null.

    2.  If `navigation`{.variable}\'s [upcoming traverse API method
        trackers](#upcoming-traverse-api-method-trackers){#ongoing-navigation-tracking:upcoming-traverse-api-method-trackers-2}\[`destinationKey`{.variable}\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ongoing-navigation-tracking:map-exists
        x-internal="map-exists"}, then:

        1.  Set `navigation`{.variable}\'s [ongoing API method
            tracker](#ongoing-api-method-tracker){#ongoing-navigation-tracking:ongoing-api-method-tracker-2}
            to `navigation`{.variable}\'s [upcoming traverse API method
            trackers](#upcoming-traverse-api-method-trackers){#ongoing-navigation-tracking:upcoming-traverse-api-method-trackers-3}\[`destinationKey`{.variable}\].

        2.  [Remove](https://infra.spec.whatwg.org/#map-remove){#ongoing-navigation-tracking:map-remove
            x-internal="map-remove"} `navigation`{.variable}\'s
            [upcoming traverse API method
            trackers](#upcoming-traverse-api-method-trackers){#ongoing-navigation-tracking:upcoming-traverse-api-method-trackers-4}\[`destinationKey`{.variable}\].

3.  Otherwise:

    1.  Set `navigation`{.variable}\'s [ongoing API method
        tracker](#ongoing-api-method-tracker){#ongoing-navigation-tracking:ongoing-api-method-tracker-3}
        to `navigation`{.variable}\'s [upcoming non-traverse API method
        tracker](#upcoming-non-traverse-api-method-tracker){#ongoing-navigation-tracking:upcoming-non-traverse-api-method-tracker-4}.

    2.  Set `navigation`{.variable}\'s [upcoming non-traverse API method
        tracker](#upcoming-non-traverse-api-method-tracker){#ongoing-navigation-tracking:upcoming-non-traverse-api-method-tracker-5}
        to null.

To [clean up]{#navigation-api-method-tracker-clean-up .dfn} a
[navigation API method
tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-7}
`apiMethodTracker`{.variable}:

1.  Let `navigation`{.variable} be `apiMethodTracker`{.variable}\'s
    [navigation
    object](#navigation-api-method-tracker-navigation){#ongoing-navigation-tracking:navigation-api-method-tracker-navigation-3}.

2.  If `navigation`{.variable}\'s [ongoing API method
    tracker](#ongoing-api-method-tracker){#ongoing-navigation-tracking:ongoing-api-method-tracker-4}
    is `apiMethodTracker`{.variable}, then set
    `navigation`{.variable}\'s [ongoing API method
    tracker](#ongoing-api-method-tracker){#ongoing-navigation-tracking:ongoing-api-method-tracker-5}
    to null.

3.  Otherwise:

    1.  Let `key`{.variable} be `apiMethodTracker`{.variable}\'s
        [key](#navigation-api-method-tracker-key){#ongoing-navigation-tracking:navigation-api-method-tracker-key-3}.

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#ongoing-navigation-tracking:assert-4
        x-internal="assert"}: `key`{.variable} is not null.

    3.  [Assert](https://infra.spec.whatwg.org/#assert){#ongoing-navigation-tracking:assert-5
        x-internal="assert"}: `navigation`{.variable}\'s [upcoming
        traverse API method
        trackers](#upcoming-traverse-api-method-trackers){#ongoing-navigation-tracking:upcoming-traverse-api-method-trackers-5}\[`key`{.variable}\]
        [exists](https://infra.spec.whatwg.org/#map-exists){#ongoing-navigation-tracking:map-exists-2
        x-internal="map-exists"}.

    4.  [Remove](https://infra.spec.whatwg.org/#map-remove){#ongoing-navigation-tracking:map-remove-2
        x-internal="map-remove"} `navigation`{.variable}\'s [upcoming
        traverse API method
        trackers](#upcoming-traverse-api-method-trackers){#ongoing-navigation-tracking:upcoming-traverse-api-method-trackers-6}\[`key`{.variable}\].

To [notify about the committed-to
entry]{#notify-about-the-committed-to-entry .dfn} given a [navigation
API method
tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-8}
`apiMethodTracker`{.variable} and a
[`NavigationHistoryEntry`{#ongoing-navigation-tracking:navigationhistoryentry-5}](#navigationhistoryentry)
`nhe`{.variable}:

1.  Set `apiMethodTracker`{.variable}\'s [committed-to
    entry](#navigation-api-method-tracker-committed-to-entry){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-to-entry-3}
    to `nhe`{.variable}.

2.  If `apiMethodTracker`{.variable}\'s [serialized
    state](#navigation-api-method-tracker-state){#ongoing-navigation-tracking:navigation-api-method-tracker-state-3}
    is not null, then set `nhe`{.variable}\'s [session history
    entry](#nhe-she){#ongoing-navigation-tracking:nhe-she}\'s
    [navigation API
    state](browsing-the-web.html#she-navigation-api-state){#ongoing-navigation-tracking:she-navigation-api-state-2}
    to `apiMethodTracker`{.variable}\'s [serialized
    state](#navigation-api-method-tracker-state){#ongoing-navigation-tracking:navigation-api-method-tracker-state-4}.

    If it\'s null, then we\'re traversing to `nhe`{.variable} via
    [`navigation.traverseTo()`{#ongoing-navigation-tracking:dom-navigation-traverseto}](#dom-navigation-traverseto),
    which does not allow changing the state.

    At this point, `apiMethodTracker`{.variable}\'s [serialized
    state](#navigation-api-method-tracker-state){#ongoing-navigation-tracking:navigation-api-method-tracker-state-5}
    is no longer needed. Implementations might want to clear it out to
    avoid keeping it alive for the lifetime of the [navigation API
    method
    tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-9}.

3.  Resolve `apiMethodTracker`{.variable}\'s [committed
    promise](#navigation-api-method-tracker-committed){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-3}
    with `nhe`{.variable}.

    At this point, `apiMethodTracker`{.variable}\'s [committed
    promise](#navigation-api-method-tracker-committed){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-4}
    is only needed in cases where it has not yet been returned to author
    code. Implementations might want to clear it out to avoid keeping it
    alive for the lifetime of the [navigation API method
    tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-10}.

To [resolve the finished promise]{#resolve-the-finished-promise .dfn}
for a [navigation API method
tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-11}
`apiMethodTracker`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#ongoing-navigation-tracking:assert-6
    x-internal="assert"}: `apiMethodTracker`{.variable}\'s [committed-to
    entry](#navigation-api-method-tracker-committed-to-entry){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-to-entry-4}
    is not null.

2.  Resolve `apiMethodTracker`{.variable}\'s [finished
    promise](#navigation-api-method-tracker-finished){#ongoing-navigation-tracking:navigation-api-method-tracker-finished-3}
    with its [committed-to
    entry](#navigation-api-method-tracker-committed-to-entry){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-to-entry-5}.

3.  [Clean
    up](#navigation-api-method-tracker-clean-up){#ongoing-navigation-tracking:navigation-api-method-tracker-clean-up}
    `apiMethodTracker`{.variable}.

To [reject the finished promise]{#reject-the-finished-promise .dfn} for
a [navigation API method
tracker](#navigation-api-method-tracker){#ongoing-navigation-tracking:navigation-api-method-tracker-12}
`apiMethodTracker`{.variable} with a JavaScript value
`exception`{.variable}:

1.  Reject `apiMethodTracker`{.variable}\'s [committed
    promise](#navigation-api-method-tracker-committed){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-5}
    with `exception`{.variable}.

    This will do nothing if `apiMethodTracker`{.variable}\'s [committed
    promise](#navigation-api-method-tracker-committed){#ongoing-navigation-tracking:navigation-api-method-tracker-committed-6}
    was previously resolved via [notify about the committed-to
    entry](#notify-about-the-committed-to-entry){#ongoing-navigation-tracking:notify-about-the-committed-to-entry}.

2.  Reject `apiMethodTracker`{.variable}\'s [finished
    promise](#navigation-api-method-tracker-finished){#ongoing-navigation-tracking:navigation-api-method-tracker-finished-4}
    with `exception`{.variable}.

3.  [Clean
    up](#navigation-api-method-tracker-clean-up){#ongoing-navigation-tracking:navigation-api-method-tracker-clean-up-2}
    `apiMethodTracker`{.variable}.

To [abort the ongoing navigation]{#abort-the-ongoing-navigation .dfn}
given a
[`Navigation`{#ongoing-navigation-tracking:navigation-7}](#navigation)
`navigation`{.variable} and an optional
[`DOMException`{#ongoing-navigation-tracking:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
`error`{.variable}:

1.  Let `event`{.variable} be `navigation`{.variable}\'s [ongoing
    `navigate`
    event](#ongoing-navigate-event){#ongoing-navigation-tracking:ongoing-navigate-event}.

2.  [Assert](https://infra.spec.whatwg.org/#assert){#ongoing-navigation-tracking:assert-7
    x-internal="assert"}: `event`{.variable} is not null.

3.  Set `navigation`{.variable}\'s [focus changed during ongoing
    navigation](#focus-changed-during-ongoing-navigation){#ongoing-navigation-tracking:focus-changed-during-ongoing-navigation}
    to false.

4.  Set `navigation`{.variable}\'s [suppress normal scroll restoration
    during ongoing
    navigation](#suppress-normal-scroll-restoration-during-ongoing-navigation){#ongoing-navigation-tracking:suppress-normal-scroll-restoration-during-ongoing-navigation}
    to false.

5.  If `error`{.variable} was not given, then let `error`{.variable} be
    a new
    [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#ongoing-navigation-tracking:aborterror
    x-internal="aborterror"}
    [`DOMException`{#ongoing-navigation-tracking:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#ongoing-navigation-tracking:concept-relevant-realm-3}.

6.  If `event`{.variable}\'s [dispatch
    flag](https://dom.spec.whatwg.org/#dispatch-flag){#ongoing-navigation-tracking:dispatch-flag
    x-internal="dispatch-flag"} is set, then set `event`{.variable}\'s
    [canceled
    flag](https://dom.spec.whatwg.org/#canceled-flag){#ongoing-navigation-tracking:canceled-flag-3
    x-internal="canceled-flag"} to true.

7.  [Signal
    abort](https://dom.spec.whatwg.org/#abortcontroller-signal-abort){#ongoing-navigation-tracking:signal-abort-2
    x-internal="signal-abort"} on `event`{.variable}\'s [abort
    controller](#concept-navigateevent-abort-controller){#ongoing-navigation-tracking:concept-navigateevent-abort-controller-2}
    given `error`{.variable}.

8.  Set `navigation`{.variable}\'s [ongoing `navigate`
    event](#ongoing-navigate-event){#ongoing-navigation-tracking:ongoing-navigate-event-2}
    to null.

9.  Let `errorInfo`{.variable} be the result of [extracting error
    information](webappapis.html#extract-error){#ongoing-navigation-tracking:extract-error}
    from `error`{.variable}.

    For example, if this algorithm is reached because of a call to
    [`window.stop()`{#ongoing-navigation-tracking:dom-window-stop}](#dom-window-stop),
    these properties would probably end up initialized based on the line
    of script that called
    [`window.stop()`{#ongoing-navigation-tracking:dom-window-stop-2}](#dom-window-stop).
    But if it\'s because the user clicked the stop button, these
    properties would probably end up with default values like the empty
    string or 0.

10. [Fire an
    event](https://dom.spec.whatwg.org/#concept-event-fire){#ongoing-navigation-tracking:concept-event-fire
    x-internal="concept-event-fire"} named
    [`navigateerror`{#ongoing-navigation-tracking:event-navigateerror-2}](indices.html#event-navigateerror)
    at `navigation`{.variable} using
    [`ErrorEvent`{#ongoing-navigation-tracking:errorevent}](webappapis.html#errorevent),
    with additional attributes initialized according to
    `errorInfo`{.variable}.

11. If `navigation`{.variable}\'s [ongoing API method
    tracker](#ongoing-api-method-tracker){#ongoing-navigation-tracking:ongoing-api-method-tracker-6}
    is non-null, then [reject the finished
    promise](#reject-the-finished-promise){#ongoing-navigation-tracking:reject-the-finished-promise}
    for `apiMethodTracker`{.variable} with `error`{.variable}.

12. If `navigation`{.variable}\'s
    [transition](#concept-navigation-transition){#ongoing-navigation-tracking:concept-navigation-transition}
    is not null, then:

    1.  Reject `navigation`{.variable}\'s
        [transition](#concept-navigation-transition){#ongoing-navigation-tracking:concept-navigation-transition-2}\'s
        [finished
        promise](#concept-navigationtransition-finished){#ongoing-navigation-tracking:concept-navigationtransition-finished}
        with `error`{.variable}.

    2.  Set `navigation`{.variable}\'s
        [transition](#concept-navigation-transition){#ongoing-navigation-tracking:concept-navigation-transition-3}
        to null.

To [inform the navigation API about aborting
navigation]{#inform-the-navigation-api-about-aborting-navigation .dfn}
in a
[navigable](document-sequences.html#navigable){#ongoing-navigation-tracking:navigable}
`navigable`{.variable}:

1.  If this algorithm is running on `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#ongoing-navigation-tracking:nav-window}\'s
    [relevant
    agent](webappapis.html#relevant-agent){#ongoing-navigation-tracking:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#ongoing-navigation-tracking:concept-agent-event-loop},
    then continue on to the following steps. Otherwise, [queue a global
    task](webappapis.html#queue-a-global-task){#ongoing-navigation-tracking:queue-a-global-task}
    on the [navigation and traversal task
    source](webappapis.html#navigation-and-traversal-task-source){#ongoing-navigation-tracking:navigation-and-traversal-task-source}
    given `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#ongoing-navigation-tracking:nav-window-2}
    to run the following steps.

2.  Let `navigation`{.variable} be `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#ongoing-navigation-tracking:nav-window-3}\'s
    [navigation
    API](#window-navigation-api){#ongoing-navigation-tracking:window-navigation-api}.

3.  If `navigation`{.variable}\'s [ongoing `navigate`
    event](#ongoing-navigate-event){#ongoing-navigation-tracking:ongoing-navigate-event-3}
    is null, then return.

4.  [Abort the ongoing
    navigation](#abort-the-ongoing-navigation){#ongoing-navigation-tracking:abort-the-ongoing-navigation}
    given `navigation`{.variable}.

To [inform the navigation API about child navigable
destruction]{#inform-the-navigation-api-about-child-navigable-destruction
.dfn} given a
[navigable](document-sequences.html#navigable){#ongoing-navigation-tracking:navigable-2}
`navigable`{.variable}:

1.  [Inform the navigation API about aborting
    navigation](#inform-the-navigation-api-about-aborting-navigation){#ongoing-navigation-tracking:inform-the-navigation-api-about-aborting-navigation}
    in `navigable`{.variable}.

2.  Let `navigation`{.variable} be `navigable`{.variable}\'s [active
    window](document-sequences.html#nav-window){#ongoing-navigation-tracking:nav-window-4}\'s
    [navigation
    API](#window-navigation-api){#ongoing-navigation-tracking:window-navigation-api-2}.

3.  Let `traversalAPIMethodTrackers`{.variable} be a
    [clone](https://infra.spec.whatwg.org/#list-clone){#ongoing-navigation-tracking:list-clone
    x-internal="list-clone"} of `navigation`{.variable}\'s [upcoming
    traverse API method
    trackers](#upcoming-traverse-api-method-trackers){#ongoing-navigation-tracking:upcoming-traverse-api-method-trackers-7}.

4.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#ongoing-navigation-tracking:list-iterate
    x-internal="list-iterate"} `apiMethodTracker`{.variable} of
    `traversalAPIMethodTrackers`{.variable}: [reject the finished
    promise](#reject-the-finished-promise){#ongoing-navigation-tracking:reject-the-finished-promise-2}
    for `apiMethodTracker`{.variable} with a new
    [\"`AbortError`\"](https://webidl.spec.whatwg.org/#aborterror){#ongoing-navigation-tracking:aborterror-2
    x-internal="aborterror"}
    [`DOMException`{#ongoing-navigation-tracking:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#ongoing-navigation-tracking:concept-relevant-realm-4}.

------------------------------------------------------------------------

The ongoing navigation concept is most-directly exposed to web
developers through the
[`navigation.transition`{#ongoing-navigation-tracking:dom-navigation-transition}](#dom-navigation-transition)
property, which is an instance of the
[`NavigationTransition`{#ongoing-navigation-tracking:navigationtransition}](#navigationtransition)
interface:

``` idl
[Exposed=Window]
interface NavigationTransition {
  readonly attribute NavigationType navigationType;
  readonly attribute NavigationHistoryEntry from;
  readonly attribute Promise<undefined> finished;
};
```

[`navigation`](#dom-navigation){#ongoing-navigation-tracking:dom-navigation}`.`[`transition`](#dom-navigation-transition){#dom-navigation-transition-dev}

:   A
    [`NavigationTransition`{#ongoing-navigation-tracking:navigationtransition-2}](#navigationtransition)
    representing any ongoing navigation that hasn\'t yet reached the
    [`navigatesuccess`{#ongoing-navigation-tracking:event-navigatesuccess}](indices.html#event-navigatesuccess)
    or
    [`navigateerror`{#ongoing-navigation-tracking:event-navigateerror-3}](indices.html#event-navigateerror)
    stage, if one exists; or null, if there is no such transition
    ongoing.

    Since
    [`navigation.currentEntry`{#ongoing-navigation-tracking:dom-navigation-currententry}](#dom-navigation-currententry)
    (and other properties like
    [`location.href`{#ongoing-navigation-tracking:dom-location-href}](#dom-location-href))
    are updated immediately upon navigation, this
    [`navigation.transition`{#ongoing-navigation-tracking:dom-navigation-transition-2}](#dom-navigation-transition)
    property is useful for determining when such navigations are not yet
    fully settled, according to any handlers passed to
    [`navigateEvent.intercept()`{#ongoing-navigation-tracking:dom-navigateevent-intercept-6}](#dom-navigateevent-intercept).

[`navigation`](#dom-navigation){#ongoing-navigation-tracking:dom-navigation-2}`.`[`transition`](#dom-navigation-transition){#ongoing-navigation-tracking:dom-navigation-transition-3}`.`[`navigationType`](#dom-navigationtransition-navigationtype){#dom-navigationtransition-navigationtype-dev}
:   One of
    \"[`push`{#ongoing-navigation-tracking:dom-navigationtype-push}](#dom-navigationtype-push)\",
    \"[`replace`{#ongoing-navigation-tracking:dom-navigationtype-replace}](#dom-navigationtype-replace)\",
    \"[`reload`{#ongoing-navigation-tracking:dom-navigationtype-reload}](#dom-navigationtype-reload)\",
    or
    \"[`traverse`{#ongoing-navigation-tracking:dom-navigationtype-traverse-3}](#dom-navigationtype-traverse)\",
    indicating what type of navigation this transition is for.

[`navigation`](#dom-navigation){#ongoing-navigation-tracking:dom-navigation-3}`.`[`transition`](#dom-navigation-transition){#ongoing-navigation-tracking:dom-navigation-transition-4}`.`[`from`](#dom-navigationtransition-from){#dom-navigationtransition-from-dev}
:   The
    [`NavigationHistoryEntry`{#ongoing-navigation-tracking:navigationhistoryentry-7}](#navigationhistoryentry)
    from which the transition is coming. This can be useful to compare
    against
    [`navigation.currentEntry`{#ongoing-navigation-tracking:dom-navigation-currententry-2}](#dom-navigation-currententry).

[`navigation`](#dom-navigation){#ongoing-navigation-tracking:dom-navigation-4}`.`[`transition`](#dom-navigation-transition){#ongoing-navigation-tracking:dom-navigation-transition-5}`.`[`finished`](#dom-navigationtransition-finished){#dom-navigationtransition-finished-dev}

:   A promise which fulfills at the same time as the
    [`navigatesuccess`{#ongoing-navigation-tracking:event-navigatesuccess-2}](indices.html#event-navigatesuccess)
    fires, or rejects at the same time the
    [`navigateerror`{#ongoing-navigation-tracking:event-navigateerror-4}](indices.html#event-navigateerror)
    event fires.

Each
[`Navigation`{#ongoing-navigation-tracking:navigation-8}](#navigation)
has a [transition]{#concept-navigation-transition .dfn}, which is a
[`NavigationTransition`{#ongoing-navigation-tracking:navigationtransition-3}](#navigationtransition)
or null, initially null.

The [`transition`]{#dom-navigation-transition .dfn dfn-for="Navigation"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#ongoing-navigation-tracking:this
x-internal="this"}\'s
[transition](#concept-navigation-transition){#ongoing-navigation-tracking:concept-navigation-transition-4}.

Each
[`NavigationTransition`{#ongoing-navigation-tracking:navigationtransition-4}](#navigationtransition)
has an associated [navigation
type]{#concept-navigationtransition-navigationtype .dfn}, which is a
[`NavigationType`{#ongoing-navigation-tracking:navigationtype-3}](#navigationtype).

Each
[`NavigationTransition`{#ongoing-navigation-tracking:navigationtransition-5}](#navigationtransition)
has an associated [from entry]{#concept-navigationtransition-from .dfn},
which is a
[`NavigationHistoryEntry`{#ongoing-navigation-tracking:navigationhistoryentry-8}](#navigationhistoryentry).

Each
[`NavigationTransition`{#ongoing-navigation-tracking:navigationtransition-6}](#navigationtransition)
has an associated [finished
promise]{#concept-navigationtransition-finished .dfn}, which is a
promise.

The [`navigationType`]{#dom-navigationtransition-navigationtype .dfn
dfn-for="NavigationTransition" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#ongoing-navigation-tracking:this-2
x-internal="this"}\'s [navigation
type](#concept-navigationtransition-navigationtype){#ongoing-navigation-tracking:concept-navigationtransition-navigationtype}.

The [`from`]{#dom-navigationtransition-from .dfn
dfn-for="NavigationTransition" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#ongoing-navigation-tracking:this-3
x-internal="this"}\'s [from
entry](#concept-navigationtransition-from){#ongoing-navigation-tracking:concept-navigationtransition-from}.

The [`finished`]{#dom-navigationtransition-finished .dfn
dfn-for="NavigationTransition" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#ongoing-navigation-tracking:this-4
x-internal="this"}\'s [finished
promise](#concept-navigationtransition-finished){#ongoing-navigation-tracking:concept-navigationtransition-finished-2}.

##### [7.2.6.9]{.secno} The [`NavigationActivation`{#navigation-activation-interface:navigationactivation}](#navigationactivation) interface[](#navigation-activation-interface){.self-link} {#navigation-activation-interface}

``` idl
[Exposed=Window]
interface NavigationActivation {
  readonly attribute NavigationHistoryEntry? from;
  readonly attribute NavigationHistoryEntry entry;
  readonly attribute NavigationType navigationType;
};
```

[`navigation`](#dom-navigation){#navigation-activation-interface:dom-navigation}`.`[`activation`](#dom-navigation-activation){#dom-navigation-activation-dev}

:   A
    [`NavigationActivation`{#navigation-activation-interface:navigationactivation-2}](#navigationactivation)
    containing information about the most recent cross-document
    navigation, the navigation that \"activated\" this
    [`Document`{#navigation-activation-interface:document}](dom.html#document).

    While
    [`navigation.currentEntry`{#navigation-activation-interface:dom-navigation-currententry}](#dom-navigation-currententry)
    and the
    [`Document`{#navigation-activation-interface:document-2}](dom.html#document)\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#navigation-activation-interface:the-document's-address
    x-internal="the-document's-address"} can be updated regularly due to
    same-document navigations,
    [`navigation.activation`{#navigation-activation-interface:dom-navigation-activation}](#dom-navigation-activation)
    stays constant, and its properties are only updated if the
    [`Document`{#navigation-activation-interface:document-3}](dom.html#document)
    is
    [reactivated](browsing-the-web.html#reactivate-a-document){#navigation-activation-interface:reactivate-a-document}
    from history.

[`navigation`](#dom-navigation){#navigation-activation-interface:dom-navigation-2}`.`[`activation`](#dom-navigation-activation){#navigation-activation-interface:dom-navigation-activation-2}`.`[`entry`](#dom-navigationactivation-entry){#dom-navigationactivation-entry-dev}
:   A
    [`NavigationHistoryEntry`{#navigation-activation-interface:navigationhistoryentry-3}](#navigationhistoryentry),
    equivalent to the value of the
    [`navigation.currentEntry`{#navigation-activation-interface:dom-navigation-currententry-2}](#dom-navigation-currententry)
    property at the moment the
    [`Document`{#navigation-activation-interface:document-4}](dom.html#document)
    was activated.

[`navigation`](#dom-navigation){#navigation-activation-interface:dom-navigation-3}`.`[`activation`](#dom-navigation-activation){#navigation-activation-interface:dom-navigation-activation-3}`.`[`from`](#dom-navigationactivation-from){#dom-navigationactivation-from-dev}

:   A
    [`NavigationHistoryEntry`{#navigation-activation-interface:navigationhistoryentry-4}](#navigationhistoryentry),
    representing the
    [`Document`{#navigation-activation-interface:document-5}](dom.html#document)
    that was active right before the current
    [`Document`{#navigation-activation-interface:document-6}](dom.html#document).
    This will have a value null in case the previous
    [`Document`{#navigation-activation-interface:document-7}](dom.html#document)
    was not [same
    origin](browsers.html#same-origin){#navigation-activation-interface:same-origin}
    with this one or if it was the [initial
    `about:blank`](dom.html#is-initial-about:blank){#navigation-activation-interface:is-initial-about:blank}
    [`Document`{#navigation-activation-interface:document-8}](dom.html#document).

    There are some cases in which either the
    [`from`{#navigation-activation-interface:dom-navigationactivation-from-2}](#dom-navigationactivation-from)
    or
    [`entry`{#navigation-activation-interface:dom-navigationactivation-entry-2}](#dom-navigationactivation-entry)
    [`NavigationHistoryEntry`{#navigation-activation-interface:navigationhistoryentry-5}](#navigationhistoryentry)
    objects would not be viable targets for the
    [`traverseTo()`{#navigation-activation-interface:dom-navigation-traverseto}](#dom-navigation-traverseto)
    method, as they might not be retained in history. For example, the
    [`Document`{#navigation-activation-interface:document-9}](dom.html#document)
    can be activated using
    [`location.replace()`{#navigation-activation-interface:dom-location-replace}](#dom-location-replace)
    or its initial entry could be replaced by
    [`history.replaceState()`{#navigation-activation-interface:dom-history-replacestate}](#dom-history-replacestate).
    However, those entries\'
    [`url`{#navigation-activation-interface:dom-navigationhistoryentry-url}](#dom-navigationhistoryentry-url)
    property and
    [`getState()`{#navigation-activation-interface:dom-navigationhistoryentry-getstate}](#dom-navigationhistoryentry-getstate)
    method are still accessible.

[`navigation`](#dom-navigation){#navigation-activation-interface:dom-navigation-4}`.`[`activation`](#dom-navigation-activation){#navigation-activation-interface:dom-navigation-activation-4}`.`[`navigationType`](#dom-navigationactivation-navigationtype){#dom-navigationactivation-navigationtype-dev}

:   One of
    \"[`push`{#navigation-activation-interface:dom-navigationtype-push}](#dom-navigationtype-push)\",
    \"[`replace`{#navigation-activation-interface:dom-navigationtype-replace}](#dom-navigationtype-replace)\",
    \"[`reload`{#navigation-activation-interface:dom-navigationtype-reload}](#dom-navigationtype-reload)\",
    or
    \"[`traverse`{#navigation-activation-interface:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\",
    indicating what type of navigation activated this
    [`Document`{#navigation-activation-interface:document-10}](dom.html#document).

Each
[`Navigation`{#navigation-activation-interface:navigation}](#navigation)
has an associated [activation]{#navigation-activation .dfn}, which is
null or a
[`NavigationActivation`{#navigation-activation-interface:navigationactivation-3}](#navigationactivation)
object, initially null.

Each
[`NavigationActivation`{#navigation-activation-interface:navigationactivation-4}](#navigationactivation)
has:

- [old entry]{#nav-activation-old-entry .dfn}, null or a
  [`NavigationHistoryEntry`{#navigation-activation-interface:navigationhistoryentry-6}](#navigationhistoryentry).

- [new entry]{#nav-activation-new-entry .dfn}, null or a
  [`NavigationHistoryEntry`{#navigation-activation-interface:navigationhistoryentry-7}](#navigationhistoryentry).

- [navigation type]{#nav-activation-navigation-type .dfn}, a
  [`NavigationType`{#navigation-activation-interface:navigationtype-2}](#navigationtype).

The [`activation`]{#dom-navigation-activation .dfn dfn-for="Navigation"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#navigation-activation-interface:this
x-internal="this"}\'s
[activation](#navigation-activation){#navigation-activation-interface:navigation-activation}.

The [`from`]{#dom-navigationactivation-from .dfn
dfn-for="NavigationActivation" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#navigation-activation-interface:this-2
x-internal="this"}\'s [old
entry](#nav-activation-old-entry){#navigation-activation-interface:nav-activation-old-entry}.

The [`entry`]{#dom-navigationactivation-entry .dfn
dfn-for="NavigationActivation" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#navigation-activation-interface:this-3
x-internal="this"}\'s [new
entry](#nav-activation-new-entry){#navigation-activation-interface:nav-activation-new-entry}.

The [`navigationType`]{#dom-navigationactivation-navigationtype .dfn
dfn-for="NavigationActivation" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#navigation-activation-interface:this-4
x-internal="this"}\'s [navigation
type](#nav-activation-navigation-type){#navigation-activation-interface:nav-activation-navigation-type}.

##### [7.2.6.10]{.secno} The [`navigate`{#the-navigate-event:event-navigate}](indices.html#event-navigate) event[](#the-navigate-event){.self-link}

A major feature of the navigation API is the
[`navigate`{#the-navigate-event:event-navigate-2}](indices.html#event-navigate)
event. This event is fired on any navigation (in the [broad sense of the
word](#navigationtype){#the-navigate-event:navigationtype}), allowing
web developers to monitor such outgoing navigations. In many cases, the
event is
[`cancelable`{#the-navigate-event:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"},
which allows preventing the navigation from happening. And in others,
the navigation can be intercepted and replaced with a same-document
navigation by using the
[`intercept()`{#the-navigate-event:dom-navigateevent-intercept}](#dom-navigateevent-intercept)
method of the
[`NavigateEvent`{#the-navigate-event:navigateevent}](#navigateevent)
class.

###### [7.2.6.10.1]{.secno} The [`NavigateEvent`{#the-navigateevent-interface:navigateevent}](#navigateevent) interface[](#the-navigateevent-interface){.self-link}

``` idl
[Exposed=Window]
interface NavigateEvent : Event {
  constructor(DOMString type, NavigateEventInit eventInitDict);

  readonly attribute NavigationType navigationType;
  readonly attribute NavigationDestination destination;
  readonly attribute boolean canIntercept;
  readonly attribute boolean userInitiated;
  readonly attribute boolean hashChange;
  readonly attribute AbortSignal signal;
  readonly attribute FormData? formData;
  readonly attribute DOMString? downloadRequest;
  readonly attribute any info;
  readonly attribute boolean hasUAVisualTransition;
  readonly attribute Element? sourceElement;

  undefined intercept(optional NavigationInterceptOptions options = {});
  undefined scroll();
};

dictionary NavigateEventInit : EventInit {
  NavigationType navigationType = "push";
  required NavigationDestination destination;
  boolean canIntercept = false;
  boolean userInitiated = false;
  boolean hashChange = false;
  required AbortSignal signal;
  FormData? formData = null;
  DOMString? downloadRequest = null;
  any info;
  boolean hasUAVisualTransition = false;
  Element? sourceElement = null;
};

dictionary NavigationInterceptOptions {
  NavigationInterceptHandler handler;
  NavigationFocusReset focusReset;
  NavigationScrollBehavior scroll;
};

enum NavigationFocusReset {
  "after-transition",
  "manual"
};

enum NavigationScrollBehavior {
  "after-transition",
  "manual"
};

callback NavigationInterceptHandler = Promise<undefined> ();
```

`event`{.variable}`.`[`navigationType`](#dom-navigateevent-navigationtype){#dom-navigateevent-navigationtype-dev}
:   One of
    \"[`push`{#the-navigateevent-interface:dom-navigationtype-push-2}](#dom-navigationtype-push)\",
    \"[`replace`{#the-navigateevent-interface:dom-navigationtype-replace}](#dom-navigationtype-replace)\",
    \"[`reload`{#the-navigateevent-interface:dom-navigationtype-reload}](#dom-navigationtype-reload)\",
    or
    \"[`traverse`{#the-navigateevent-interface:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\",
    indicating what type of navigation this is.

`event`{.variable}`.`[`destination`](#dom-navigateevent-destination){#dom-navigateevent-destination-dev}
:   A
    [`NavigationDestination`{#the-navigateevent-interface:navigationdestination-3}](#navigationdestination)
    representing the destination of the navigation.

`event`{.variable}`.`[`canIntercept`](#dom-navigateevent-canintercept){#dom-navigateevent-canintercept-dev}

:   True if
    [`intercept()`{#the-navigateevent-interface:dom-navigateevent-intercept-2}](#dom-navigateevent-intercept)
    can be called to intercept this navigation and convert it into a
    same-document navigation, replacing its usual behavior; false
    otherwise.

    Generally speaking, this will be true whenever the current
    [`Document`{#the-navigateevent-interface:document}](dom.html#document)
    [can have its URL
    rewritten](#can-have-its-url-rewritten){#the-navigateevent-interface:can-have-its-url-rewritten}
    to the destination URL, except for in the case of cross-document
    \"[`traverse`{#the-navigateevent-interface:dom-navigationtype-traverse-2}](#dom-navigationtype-traverse)\"
    navigations, where it will always be false.

`event`{.variable}`.`[`userInitiated`](#dom-navigateevent-userinitiated){#dom-navigateevent-userinitiated-dev}
:   True if this navigation was due to a user clicking on an
    [`a`{#the-navigateevent-interface:the-a-element}](text-level-semantics.html#the-a-element)
    element, submitting a
    [`form`{#the-navigateevent-interface:the-form-element}](forms.html#the-form-element)
    element, or using the [browser
    UI](document-lifecycle.html#nav-traversal-ui) to navigate; false
    otherwise.

`event`{.variable}`.`[`hashChange`](#dom-navigateevent-hashchange){#dom-navigateevent-hashchange-dev}
:   True for a [fragment
    navigation](browsing-the-web.html#navigate-fragid){#the-navigateevent-interface:navigate-fragid};
    false otherwise.

`event`{.variable}`.`[`signal`](#dom-navigateevent-signal){#dom-navigateevent-signal-dev}

:   An
    [`AbortSignal`{#the-navigateevent-interface:abortsignal-3}](https://dom.spec.whatwg.org/#abortsignal){x-internal="abortsignal"}
    which will become aborted if the navigation gets canceled, e.g., by
    the user pressing their browser\'s \"Stop\" button, or by another
    navigation interrupting this one.

    The expected pattern is for developers to pass this along to any
    async operations, such as
    [`fetch()`{#the-navigateevent-interface:fetch()}](https://fetch.spec.whatwg.org/#dom-global-fetch){x-internal="fetch()"},
    which they perform as part of handling this navigation.

`event`{.variable}`.`[`formData`](#dom-navigateevent-formdata){#dom-navigateevent-formdata-dev}

:   The
    [`FormData`{#the-navigateevent-interface:formdata-3}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
    representing the submitted form entries for this navigation, if this
    navigation is a
    \"[`push`{#the-navigateevent-interface:dom-navigationtype-push-3}](#dom-navigationtype-push)\"
    or
    \"[`replace`{#the-navigateevent-interface:dom-navigationtype-replace-2}](#dom-navigationtype-replace)\"
    navigation representing a POST [form
    submission](form-control-infrastructure.html#concept-form-submit){#the-navigateevent-interface:concept-form-submit};
    null otherwise.

    (Notably, this will be null even for
    \"[`reload`{#the-navigateevent-interface:dom-navigationtype-reload-2}](#dom-navigationtype-reload)\"
    or
    \"[`traverse`{#the-navigateevent-interface:dom-navigationtype-traverse-3}](#dom-navigationtype-traverse)\"
    navigations that are revisiting a [session history
    entry](browsing-the-web.html#session-history-entry){#the-navigateevent-interface:session-history-entry}
    that was originally created from a form submission.)

`event`{.variable}`.`[`downloadRequest`](#dom-navigateevent-downloadrequest){#dom-navigateevent-downloadrequest-dev}

:   Represents whether or not this navigation was requested to be a
    download, by using an
    [`a`{#the-navigateevent-interface:the-a-element-2}](text-level-semantics.html#the-a-element)
    or
    [`area`{#the-navigateevent-interface:the-area-element}](image-maps.html#the-area-element)
    element\'s
    [`download`{#the-navigateevent-interface:attr-hyperlink-download}](links.html#attr-hyperlink-download)
    attribute:

    - If a download was not requested, then this property is null.

    - If a download was requested, returns the filename that was
      supplied as the
      [`download`{#the-navigateevent-interface:attr-hyperlink-download-2}](links.html#attr-hyperlink-download)
      attribute\'s value. (This could be the empty string.)

    Note that a download being requested does not always mean that a
    download will happen: for example, a download might be blocked by
    browser security policies, or end up being treated as a
    \"[`push`{#the-navigateevent-interface:navigationhistorybehavior-push}](browsing-the-web.html#navigationhistorybehavior-push)\"
    navigation for [unspecified
    reasons](https://github.com/whatwg/html/issues/7718){.XXX}.

    Similarly, a navigation might end up being a download even if it was
    not requested to be one, due to the destination server responding
    with a
    \`[`Content-Disposition: attachment`{#the-navigateevent-interface:http-content-disposition}](https://httpwg.org/specs/rfc6266.html){x-internal="http-content-disposition"}\`
    header.

    Finally, note that the
    [`navigate`{#the-navigateevent-interface:event-navigate}](indices.html#event-navigate)
    event will not fire at all for downloads initiated using browser UI
    affordances, e.g., those created by right-clicking and choosing to
    save the target of a link.

`event`{.variable}`.`[`info`](#dom-navigateevent-info){#dom-navigateevent-info-dev}
:   An arbitrary JavaScript value passed via one of the [navigation API
    methods](#navigation-api-initiating-navigations) which initiated
    this navigation, or undefined if the navigation was initiated by the
    user or by a different API.

`event`{.variable}`.`[`hasUAVisualTransition`](#dom-navigateevent-hasuavisualtransition){#dom-navigateevent-hasuavisualtransition-dev}
:   Returns true if the user agent performed a visual transition for
    this navigation before dispatching this event. If true, the best
    user experience will be given if the author synchronously updates
    the DOM to the post-navigation state.

`event`{.variable}`.`[`sourceElement`](#dom-navigateevent-sourceelement){#dom-navigateevent-sourceelement-dev}
:   Returns the
    [`Element`{#the-navigateevent-interface:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}
    responsible for this navigation. This can be an
    [`a`{#the-navigateevent-interface:the-a-element-3}](text-level-semantics.html#the-a-element)or
    [`area`{#the-navigateevent-interface:the-area-element-2}](image-maps.html#the-area-element)
    element, a [submit
    button](forms.html#concept-submit-button){#the-navigateevent-interface:concept-submit-button},
    or a submitted
    [`form`{#the-navigateevent-interface:the-form-element-2}](forms.html#the-form-element)
    element.

`event`{.variable}`.`[`intercept`](#dom-navigateevent-intercept){#dom-navigateevent-intercept-dev}`({ `[`handler`](#dom-navigationinterceptoptions-handler){#dom-navigationinterceptoptions-handler-dev}`, `[`focusReset`](#dom-navigationinterceptoptions-focusreset){#dom-navigationinterceptoptions-focusreset-dev}`, `[`scroll`](#dom-navigationinterceptoptions-scroll){#dom-navigationinterceptoptions-scroll-dev}` })`

:   Intercepts this navigation, preventing its normal handling and
    instead converting it into a same-document navigation of the same
    type to the destination URL.

    The
    [`handler`{#the-navigateevent-interface:dom-navigationinterceptoptions-handler}](#dom-navigationinterceptoptions-handler)
    option can be a function that returns a promise. The handler
    function will run after the
    [`navigate`{#the-navigateevent-interface:event-navigate-2}](indices.html#event-navigate)
    event has finished firing, and the
    [`navigation.currentEntry`{#the-navigateevent-interface:dom-navigation-currententry}](#dom-navigation-currententry)
    property has been synchronously updated. This returned promise is
    used to signal the duration, and success or failure, of the
    navigation. After it settles, the browser signals to the user (e.g.,
    via a loading spinner UI, or assistive technology) that the
    navigation is finished. Additionally, it fires
    [`navigatesuccess`{#the-navigateevent-interface:event-navigatesuccess}](indices.html#event-navigatesuccess)
    or
    [`navigateerror`{#the-navigateevent-interface:event-navigateerror}](indices.html#event-navigateerror)
    events as appropriate, which other parts of the web application can
    respond to.

    By default, using this method will cause focus to reset when any
    handlers\' returned promises settle. Focus will be reset to the
    first element with the
    [`autofocus`{#the-navigateevent-interface:attr-fe-autofocus}](interaction.html#attr-fe-autofocus)
    attribute set, or [the body
    element](dom.html#the-body-element-2){#the-navigateevent-interface:the-body-element-2}
    if the attribute isn\'t present. The
    [`focusReset`{#the-navigateevent-interface:dom-navigationinterceptoptions-focusreset}](#dom-navigationinterceptoptions-focusreset)
    option can be set to
    \"[`manual`{#the-navigateevent-interface:dom-navigationfocusreset-manual}](#dom-navigationfocusreset-manual)\"
    to avoid this behavior.

    By default, using this method will delay the browser\'s scroll
    restoration logic for
    \"[`traverse`{#the-navigateevent-interface:dom-navigationtype-traverse-4}](#dom-navigationtype-traverse)\"
    or
    \"[`reload`{#the-navigateevent-interface:dom-navigationtype-reload-3}](#dom-navigationtype-reload)\"
    navigations, or its scroll-reset/scroll-to-a-fragment logic for
    \"[`push`{#the-navigateevent-interface:dom-navigationtype-push-4}](#dom-navigationtype-push)\"
    or
    \"[`replace`{#the-navigateevent-interface:dom-navigationtype-replace-3}](#dom-navigationtype-replace)\"
    navigations, until any handlers\' returned promises settle. The
    [`scroll`{#the-navigateevent-interface:dom-navigationinterceptoptions-scroll}](#dom-navigationinterceptoptions-scroll)
    option can be set to
    \"[`manual`{#the-navigateevent-interface:dom-navigationscrollbehavior-manual}](#dom-navigationscrollbehavior-manual)\"
    to turn off any browser-driven scroll behavior entirely for this
    navigation, or
    [`scroll()`{#the-navigateevent-interface:dom-navigateevent-scroll-2}](#dom-navigateevent-scroll)
    can be called before the promise settles to trigger this behavior
    early.

    This method will throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-navigateevent-interface:securityerror
    x-internal="securityerror"}
    [`DOMException`{#the-navigateevent-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if
    [`canIntercept`{#the-navigateevent-interface:dom-navigateevent-canintercept-2}](#dom-navigateevent-canintercept)
    is false, or if
    [`isTrusted`{#the-navigateevent-interface:dom-event-istrusted}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
    is false. It will throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-navigateevent-interface:invalidstateerror
    x-internal="invalidstateerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-2}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}
    if not called synchronously, during event dispatch.

`event`{.variable}`.`[`scroll`](#dom-navigateevent-scroll){#dom-navigateevent-scroll-dev}`()`

:   For
    \"[`traverse`{#the-navigateevent-interface:dom-navigationtype-traverse-5}](#dom-navigationtype-traverse)\"
    or
    \"[`reload`{#the-navigateevent-interface:dom-navigationtype-reload-4}](#dom-navigationtype-reload)\"
    navigations, restores the scroll position using the browser\'s usual
    scroll restoration logic.

    For
    \"[`push`{#the-navigateevent-interface:dom-navigationtype-push-5}](#dom-navigationtype-push)\"
    or
    \"[`replace`{#the-navigateevent-interface:dom-navigationtype-replace-4}](#dom-navigationtype-replace)\"
    navigations, either resets the scroll position to the top of the
    document or scrolls to the
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#the-navigateevent-interface:concept-url-fragment
    x-internal="concept-url-fragment"} specified by
    [`destination.url`{#the-navigateevent-interface:dom-navigationdestination-url}](#dom-navigationdestination-url)
    if there is one.

    If called more than once, or called after automatic post-transition
    scroll processing has happened due to the
    [`scroll`{#the-navigateevent-interface:dom-navigationinterceptoptions-scroll-2}](#dom-navigationinterceptoptions-scroll)
    option being left as
    \"[`after-transition`{#the-navigateevent-interface:dom-navigationscrollbehavior-after-transition}](#dom-navigationscrollbehavior-after-transition)\",
    or called before the navigation has committed, this method will
    throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-navigateevent-interface:invalidstateerror-2
    x-internal="invalidstateerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-3}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

Each
[`NavigateEvent`{#the-navigateevent-interface:navigateevent-2}](#navigateevent)
has an [interception state]{#concept-navigateevent-interception-state
.dfn}, which is either \"`none`\", \"`intercepted`\", \"`committed`\",
\"`scrolled`\", or \"`finished`\", initially \"`none`\".

Each
[`NavigateEvent`{#the-navigateevent-interface:navigateevent-3}](#navigateevent)
has a [navigation handler
list]{#concept-navigateevent-navigation-handler-list .dfn}, a
[list](https://infra.spec.whatwg.org/#list){#the-navigateevent-interface:list
x-internal="list"} of
[`NavigationInterceptHandler`{#the-navigateevent-interface:navigationintercepthandler-2}](#navigationintercepthandler)
callbacks, initially empty.

Each
[`NavigateEvent`{#the-navigateevent-interface:navigateevent-4}](#navigateevent)
has a [focus reset behavior]{#concept-navigateevent-focusreset .dfn}, a
[`NavigationFocusReset`{#the-navigateevent-interface:navigationfocusreset-2}](#navigationfocusreset)-or-null,
initially null.

Each
[`NavigateEvent`{#the-navigateevent-interface:navigateevent-5}](#navigateevent)
has a [scroll behavior]{#concept-navigateevent-scroll .dfn}, a
[`NavigationScrollBehavior`{#the-navigateevent-interface:navigationscrollbehavior-2}](#navigationscrollbehavior)-or-null,
initially null.

Each
[`NavigateEvent`{#the-navigateevent-interface:navigateevent-6}](#navigateevent)
has an [abort controller]{#concept-navigateevent-abort-controller .dfn},
an
[`AbortController`{#the-navigateevent-interface:abortcontroller}](https://dom.spec.whatwg.org/#abortcontroller){x-internal="abortcontroller"}-or-null,
initially null.

Each
[`NavigateEvent`{#the-navigateevent-interface:navigateevent-7}](#navigateevent)
has a [classic history API
state]{#concept-navigateevent-classic-history-api-state .dfn}, a
[serialized
state](browsing-the-web.html#serialized-state){#the-navigateevent-interface:serialized-state}
or null. It is only used in some cases where the event\'s
[`navigationType`{#the-navigateevent-interface:dom-navigateevent-navigationtype-2}](#dom-navigateevent-navigationtype)
is
\"[`push`{#the-navigateevent-interface:dom-navigationtype-push-6}](#dom-navigationtype-push)\"
or
\"[`replace`{#the-navigateevent-interface:dom-navigationtype-replace-5}](#dom-navigationtype-replace)\",
and is set appropriately when the event is
[fired](https://dom.spec.whatwg.org/#concept-event-fire){#the-navigateevent-interface:concept-event-fire
x-internal="concept-event-fire"}.

The [`navigationType`]{#dom-navigateevent-navigationtype .dfn
dfn-for="NavigateEvent" dfn-type="attribute"},
[`destination`]{#dom-navigateevent-destination .dfn
dfn-for="NavigateEvent" dfn-type="attribute"},
[`canIntercept`]{#dom-navigateevent-canintercept .dfn
dfn-for="NavigateEvent" dfn-type="attribute"},
[`userInitiated`]{#dom-navigateevent-userinitiated .dfn
dfn-for="NavigateEvent" dfn-type="attribute"},
[`hashChange`]{#dom-navigateevent-hashchange .dfn
dfn-for="NavigateEvent" dfn-type="attribute"},
[`signal`]{#dom-navigateevent-signal .dfn dfn-for="NavigateEvent"
dfn-type="attribute"}, [`formData`]{#dom-navigateevent-formdata .dfn
dfn-for="NavigateEvent" dfn-type="attribute"},
[`downloadRequest`]{#dom-navigateevent-downloadrequest .dfn
dfn-for="NavigateEvent" dfn-type="attribute"},
[`info`]{#dom-navigateevent-info .dfn dfn-for="NavigateEvent"
dfn-type="attribute"},
[`hasUAVisualTransition`]{#dom-navigateevent-hasuavisualtransition .dfn
dfn-for="NavigateEvent" dfn-type="attribute"}, and
[`sourceElement`]{#dom-navigateevent-sourceelement .dfn
dfn-for="NavigateEvent" dfn-type="attribute"} attributes must return the
values they are initialized to.

The [`intercept(``options`{.variable}`)`]{#dom-navigateevent-intercept
.dfn dfn-for="NavigateEvent" dfn-type="method"} method steps are:

1.  [Perform shared
    checks](#navigateevent-perform-shared-checks){#the-navigateevent-interface:navigateevent-perform-shared-checks}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this
    x-internal="this"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-2
    x-internal="this"}\'s
    [`canIntercept`{#the-navigateevent-interface:dom-navigateevent-canintercept-3}](#dom-navigateevent-canintercept)
    attribute was initialized to false, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-navigateevent-interface:securityerror-2
    x-internal="securityerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-4}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-3
    x-internal="this"}\'s [dispatch
    flag](https://dom.spec.whatwg.org/#dispatch-flag){#the-navigateevent-interface:dispatch-flag
    x-internal="dispatch-flag"} is unset, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-navigateevent-interface:invalidstateerror-3
    x-internal="invalidstateerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-5}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

4.  [Assert](https://infra.spec.whatwg.org/#assert){#the-navigateevent-interface:assert
    x-internal="assert"}:
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-4
    x-internal="this"}\'s [interception
    state](#concept-navigateevent-interception-state){#the-navigateevent-interface:concept-navigateevent-interception-state}
    is either \"`none`\" or \"`intercepted`\".

5.  Set
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-5
    x-internal="this"}\'s [interception
    state](#concept-navigateevent-interception-state){#the-navigateevent-interface:concept-navigateevent-interception-state-2}
    to \"`intercepted`\".

6.  If
    `options`{.variable}\[\"[`handler`{#the-navigateevent-interface:dom-navigationinterceptoptions-handler-2}](#dom-navigationinterceptoptions-handler)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-navigateevent-interface:map-exists
    x-internal="map-exists"}, then
    [append](https://infra.spec.whatwg.org/#list-append){#the-navigateevent-interface:list-append
    x-internal="list-append"} it to
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-6
    x-internal="this"}\'s [navigation handler
    list](#concept-navigateevent-navigation-handler-list){#the-navigateevent-interface:concept-navigateevent-navigation-handler-list}.

7.  If
    `options`{.variable}\[\"[`focusReset`{#the-navigateevent-interface:dom-navigationinterceptoptions-focusreset-2}](#dom-navigationinterceptoptions-focusreset)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-navigateevent-interface:map-exists-2
    x-internal="map-exists"}, then:

    1.  If
        [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-7
        x-internal="this"}\'s [focus reset
        behavior](#concept-navigateevent-focusreset){#the-navigateevent-interface:concept-navigateevent-focusreset}
        is not null, and it is not equal to
        `options`{.variable}\[\"[`focusReset`{#the-navigateevent-interface:dom-navigationinterceptoptions-focusreset-3}](#dom-navigationinterceptoptions-focusreset)\"\],
        then the user agent may [report a warning to the
        console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#the-navigateevent-interface:report-a-warning-to-the-console
        x-internal="report-a-warning-to-the-console"} indicating that
        the
        [`focusReset`{#the-navigateevent-interface:dom-navigationinterceptoptions-focusreset-4}](#dom-navigationinterceptoptions-focusreset)
        option for a previous call to
        [`intercept()`{#the-navigateevent-interface:dom-navigateevent-intercept-3}](#dom-navigateevent-intercept)
        was overridden by this new value, and the previous value will be
        ignored.

    2.  Set
        [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-8
        x-internal="this"}\'s [focus reset
        behavior](#concept-navigateevent-focusreset){#the-navigateevent-interface:concept-navigateevent-focusreset-2}
        to
        `options`{.variable}\[\"[`focusReset`{#the-navigateevent-interface:dom-navigationinterceptoptions-focusreset-5}](#dom-navigationinterceptoptions-focusreset)\"\].

8.  If
    `options`{.variable}\[\"[`scroll`{#the-navigateevent-interface:dom-navigationinterceptoptions-scroll-3}](#dom-navigationinterceptoptions-scroll)\"\]
    [exists](https://infra.spec.whatwg.org/#map-exists){#the-navigateevent-interface:map-exists-3
    x-internal="map-exists"}, then:

    1.  If
        [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-9
        x-internal="this"}\'s [scroll
        behavior](#concept-navigateevent-scroll){#the-navigateevent-interface:concept-navigateevent-scroll}
        is not null, and it is not equal to
        `options`{.variable}\[\"[`scroll`{#the-navigateevent-interface:dom-navigationinterceptoptions-scroll-4}](#dom-navigationinterceptoptions-scroll)\"\],
        then the user agent may [report a warning to the
        console](https://console.spec.whatwg.org/#report-a-warning-to-the-console){#the-navigateevent-interface:report-a-warning-to-the-console-2
        x-internal="report-a-warning-to-the-console"} indicating that
        the
        [`scroll`{#the-navigateevent-interface:dom-navigationinterceptoptions-scroll-5}](#dom-navigationinterceptoptions-scroll)
        option for a previous call to
        [`intercept()`{#the-navigateevent-interface:dom-navigateevent-intercept-4}](#dom-navigateevent-intercept)
        was overridden by this new value, and the previous value will be
        ignored.

    2.  Set
        [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-10
        x-internal="this"}\'s [scroll
        behavior](#concept-navigateevent-scroll){#the-navigateevent-interface:concept-navigateevent-scroll-2}
        to
        `options`{.variable}\[\"[`scroll`{#the-navigateevent-interface:dom-navigationinterceptoptions-scroll-6}](#dom-navigationinterceptoptions-scroll)\"\].

The [`scroll()`]{#dom-navigateevent-scroll .dfn dfn-for="NavigateEvent"
dfn-type="method"} method steps are:

1.  [Perform shared
    checks](#navigateevent-perform-shared-checks){#the-navigateevent-interface:navigateevent-perform-shared-checks-2}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-11
    x-internal="this"}.

2.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-12
    x-internal="this"}\'s [interception
    state](#concept-navigateevent-interception-state){#the-navigateevent-interface:concept-navigateevent-interception-state-3}
    is not \"`committed`\", then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-navigateevent-interface:invalidstateerror-4
    x-internal="invalidstateerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-6}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  [Process scroll
    behavior](#process-scroll-behavior){#the-navigateevent-interface:process-scroll-behavior}
    given
    [this](https://webidl.spec.whatwg.org/#this){#the-navigateevent-interface:this-13
    x-internal="this"}.

To [perform shared checks]{#navigateevent-perform-shared-checks .dfn}
for a
[`NavigateEvent`{#the-navigateevent-interface:navigateevent-8}](#navigateevent)
`event`{.variable}:

1.  If `event`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#the-navigateevent-interface:concept-relevant-global}\'s
    [associated
    `Document`](#concept-document-window){#the-navigateevent-interface:concept-document-window}
    is not [fully
    active](document-sequences.html#fully-active){#the-navigateevent-interface:fully-active},
    then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-navigateevent-interface:invalidstateerror-5
    x-internal="invalidstateerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-7}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

2.  If `event`{.variable}\'s
    [`isTrusted`{#the-navigateevent-interface:dom-event-istrusted-2}](https://dom.spec.whatwg.org/#dom-event-istrusted){x-internal="dom-event-istrusted"}
    attribute was initialized to false, then throw a
    [\"`SecurityError`\"](https://webidl.spec.whatwg.org/#securityerror){#the-navigateevent-interface:securityerror-3
    x-internal="securityerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-8}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

3.  If `event`{.variable}\'s [canceled
    flag](https://dom.spec.whatwg.org/#canceled-flag){#the-navigateevent-interface:canceled-flag
    x-internal="canceled-flag"} is set, then throw an
    [\"`InvalidStateError`\"](https://webidl.spec.whatwg.org/#invalidstateerror){#the-navigateevent-interface:invalidstateerror-6
    x-internal="invalidstateerror"}
    [`DOMException`{#the-navigateevent-interface:domexception-9}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

###### [7.2.6.10.2]{.secno} The [`NavigationDestination`{#the-navigationdestination-interface:navigationdestination}](#navigationdestination) interface[](#the-navigationdestination-interface){.self-link}

``` idl
[Exposed=Window]
interface NavigationDestination {
  readonly attribute USVString url;
  readonly attribute DOMString key;
  readonly attribute DOMString id;
  readonly attribute long long index;
  readonly attribute boolean sameDocument;

  any getState();
};
```

`event`{.variable}`.`[`destination`](#dom-navigateevent-destination){#the-navigationdestination-interface:dom-navigateevent-destination}`.`[`url`](#dom-navigationdestination-url){#the-navigationdestination-interface:dom-navigationdestination-url-2}
:   The URL being navigated to.

`event`{.variable}`.`[`destination`](#dom-navigateevent-destination){#the-navigationdestination-interface:dom-navigateevent-destination-2}`.`[`key`](#dom-navigationdestination-key){#the-navigationdestination-interface:dom-navigationdestination-key-2}
:   The value of the
    [`key`{#the-navigationdestination-interface:dom-navigationhistoryentry-key}](#dom-navigationhistoryentry-key)
    property of the destination
    [`NavigationHistoryEntry`{#the-navigationdestination-interface:navigationhistoryentry}](#navigationhistoryentry),
    if this is a
    \"[`traverse`{#the-navigationdestination-interface:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\"
    navigation, or the empty string otherwise.

`event`{.variable}`.`[`destination`](#dom-navigateevent-destination){#the-navigationdestination-interface:dom-navigateevent-destination-3}`.`[`id`](#dom-navigationdestination-id){#the-navigationdestination-interface:dom-navigationdestination-id-2}
:   The value of the
    [`id`{#the-navigationdestination-interface:dom-navigationhistoryentry-id}](#dom-navigationhistoryentry-id)
    property of the destination
    [`NavigationHistoryEntry`{#the-navigationdestination-interface:navigationhistoryentry-2}](#navigationhistoryentry),
    if this is a
    \"[`traverse`{#the-navigationdestination-interface:dom-navigationtype-traverse-2}](#dom-navigationtype-traverse)\"
    navigation, or the empty string otherwise.

`event`{.variable}`.`[`destination`](#dom-navigateevent-destination){#the-navigationdestination-interface:dom-navigateevent-destination-4}`.`[`index`](#dom-navigationdestination-index){#the-navigationdestination-interface:dom-navigationdestination-index-2}
:   The value of the
    [`index`{#the-navigationdestination-interface:dom-navigationhistoryentry-index}](#dom-navigationhistoryentry-index)
    property of the destination
    [`NavigationHistoryEntry`{#the-navigationdestination-interface:navigationhistoryentry-3}](#navigationhistoryentry),
    if this is a
    \"[`traverse`{#the-navigationdestination-interface:dom-navigationtype-traverse-3}](#dom-navigationtype-traverse)\"
    navigation, or −1 otherwise.

`event`{.variable}`.`[`destination`](#dom-navigateevent-destination){#the-navigationdestination-interface:dom-navigateevent-destination-5}`.`[`sameDocument`](#dom-navigationdestination-samedocument){#the-navigationdestination-interface:dom-navigationdestination-samedocument-2}

:   Indicates whether or not this navigation is to the same
    [`Document`{#the-navigationdestination-interface:document}](dom.html#document)
    as the current one, or not. This will be true, for example, in the
    case of fragment navigations or
    [`history.pushState()`{#the-navigationdestination-interface:dom-history-pushstate}](#dom-history-pushstate)
    navigations.

    Note that this property indicates the original nature of the
    navigation. If a cross-document navigation is converted into a
    same-document navigation using
    [`navigateEvent.intercept()`{#the-navigationdestination-interface:dom-navigateevent-intercept}](#dom-navigateevent-intercept),
    that will not change the value of this property.

`event`{.variable}`.`[`destination`](#dom-navigateevent-destination){#the-navigationdestination-interface:dom-navigateevent-destination-6}`.`[`getState`](#dom-navigationdestination-getstate){#the-navigationdestination-interface:dom-navigationdestination-getstate-2}`()`

:   For
    \"[`traverse`{#the-navigationdestination-interface:dom-navigationtype-traverse-4}](#dom-navigationtype-traverse)\"
    navigations, returns the
    [deserialization](structured-data.html#structureddeserialize){#the-navigationdestination-interface:structureddeserialize}
    of the state stored in the destination [session history
    entry](browsing-the-web.html#session-history-entry){#the-navigationdestination-interface:session-history-entry}.

    For
    \"[`push`{#the-navigationdestination-interface:dom-navigationtype-push}](#dom-navigationtype-push)\"
    or
    \"[`replace`{#the-navigationdestination-interface:dom-navigationtype-replace}](#dom-navigationtype-replace)\"
    navigations, returns the
    [deserialization](structured-data.html#structureddeserialize){#the-navigationdestination-interface:structureddeserialize-2}
    of the state passed to
    [`navigation.navigate()`{#the-navigationdestination-interface:dom-navigation-navigate}](#dom-navigation-navigate),
    if the navigation was initiated by that method, or undefined it if
    it wasn\'t.

    For
    \"[`reload`{#the-navigationdestination-interface:dom-navigationtype-reload}](#dom-navigationtype-reload)\"
    navigations, returns the
    [deserialization](structured-data.html#structureddeserialize){#the-navigationdestination-interface:structureddeserialize-3}
    of the state passed to
    [`navigation.reload()`{#the-navigationdestination-interface:dom-navigation-reload}](#dom-navigation-reload),
    if the reload was initiated by that method, or undefined it if it
    wasn\'t.

Each
[`NavigationDestination`{#the-navigationdestination-interface:navigationdestination-2}](#navigationdestination)
has a [URL]{#concept-navigationdestination-url .dfn}, which is a
[URL](https://url.spec.whatwg.org/#concept-url){#the-navigationdestination-interface:url
x-internal="url"}.

Each
[`NavigationDestination`{#the-navigationdestination-interface:navigationdestination-3}](#navigationdestination)
has an [entry]{#concept-navigationdestination-entry .dfn}, which is a
[`NavigationHistoryEntry`{#the-navigationdestination-interface:navigationhistoryentry-4}](#navigationhistoryentry)
or null.

It will be non-null if and only if the
[`NavigationDestination`{#the-navigationdestination-interface:navigationdestination-4}](#navigationdestination)
corresponds to a
\"[`traverse`{#the-navigationdestination-interface:dom-navigationtype-traverse-5}](#dom-navigationtype-traverse)\"
navigation.

Each
[`NavigationDestination`{#the-navigationdestination-interface:navigationdestination-5}](#navigationdestination)
has a [state]{#concept-navigationdestination-state .dfn}, which is a
[serialized
state](browsing-the-web.html#serialized-state){#the-navigationdestination-interface:serialized-state}.

Each
[`NavigationDestination`{#the-navigationdestination-interface:navigationdestination-6}](#navigationdestination)
has an [is same document]{#concept-navigationdestination-samedocument
.dfn}, which is a boolean.

------------------------------------------------------------------------

The [`url`]{#dom-navigationdestination-url .dfn
dfn-for="NavigationDestination" dfn-type="attribute"} getter steps are
to return
[this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this
x-internal="this"}\'s
[URL](#concept-navigationdestination-url){#the-navigationdestination-interface:concept-navigationdestination-url},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#the-navigationdestination-interface:concept-url-serializer
x-internal="concept-url-serializer"}.

The [`key`]{#dom-navigationdestination-key .dfn
dfn-for="NavigationDestination" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-2
    x-internal="this"}\'s
    [entry](#concept-navigationdestination-entry){#the-navigationdestination-interface:concept-navigationdestination-entry}
    is null, then return the empty string.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-3
    x-internal="this"}\'s
    [entry](#concept-navigationdestination-entry){#the-navigationdestination-interface:concept-navigationdestination-entry-2}\'s
    [key](#concept-navigationhistoryentry-key){#the-navigationdestination-interface:concept-navigationhistoryentry-key}.

The [`id`]{#dom-navigationdestination-id .dfn
dfn-for="NavigationDestination" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-4
    x-internal="this"}\'s
    [entry](#concept-navigationdestination-entry){#the-navigationdestination-interface:concept-navigationdestination-entry-3}
    is null, then return the empty string.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-5
    x-internal="this"}\'s
    [entry](#concept-navigationdestination-entry){#the-navigationdestination-interface:concept-navigationdestination-entry-4}\'s
    [ID](#concept-navigationhistoryentry-id){#the-navigationdestination-interface:concept-navigationhistoryentry-id}.

The [`index`]{#dom-navigationdestination-index .dfn
dfn-for="NavigationDestination" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-6
    x-internal="this"}\'s
    [entry](#concept-navigationdestination-entry){#the-navigationdestination-interface:concept-navigationdestination-entry-5}
    is null, then return −1.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-7
    x-internal="this"}\'s
    [entry](#concept-navigationdestination-entry){#the-navigationdestination-interface:concept-navigationdestination-entry-6}\'s
    [index](#concept-navigationhistoryentry-index){#the-navigationdestination-interface:concept-navigationhistoryentry-index}.

The [`sameDocument`]{#dom-navigationdestination-samedocument .dfn
dfn-for="NavigationDestination" dfn-type="attribute"} getter steps are
to return
[this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-8
x-internal="this"}\'s [is same
document](#concept-navigationdestination-samedocument){#the-navigationdestination-interface:concept-navigationdestination-samedocument}.

The [`getState()`]{#dom-navigationdestination-getstate .dfn
dfn-for="NavigationDestination" dfn-type="method"} method steps are to
return
[StructuredDeserialize](structured-data.html#structureddeserialize){#the-navigationdestination-interface:structureddeserialize-4}([this](https://webidl.spec.whatwg.org/#this){#the-navigationdestination-interface:this-9
x-internal="this"}\'s
[state](#concept-navigationdestination-state){#the-navigationdestination-interface:concept-navigationdestination-state}).

###### [7.2.6.10.3]{.secno} Firing the event[](#navigate-event-firing){.self-link} {#navigate-event-firing}

Other parts of the standard fire the
[`navigate`{#navigate-event-firing:event-navigate}](indices.html#event-navigate)
event, through a series of wrapper algorithms given in this section.

To [fire a traverse `navigate` event]{#fire-a-traverse-navigate-event
.dfn} at a
[`Navigation`{#navigate-event-firing:navigation}](#navigation)
`navigation`{.variable} given a [session history
entry](browsing-the-web.html#session-history-entry){#navigate-event-firing:session-history-entry}
[`destinationSHE`{.variable}]{#fire-navigate-traverse-destinationshe
.dfn} and an optional [user navigation
involvement](browsing-the-web.html#user-navigation-involvement){#navigate-event-firing:user-navigation-involvement}
[`userInvolvement`{.variable}]{#fire-navigate-traverse-userinvolvement
.dfn} (default
\"[`none`{#navigate-event-firing:uni-none}](browsing-the-web.html#uni-none)\"):

1.  Let `event`{.variable} be the result of [creating an
    event](https://dom.spec.whatwg.org/#concept-event-create){#navigate-event-firing:creating-an-event
    x-internal="creating-an-event"} given
    [`NavigateEvent`{#navigate-event-firing:navigateevent}](#navigateevent),
    in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm}.

2.  Set `event`{.variable}\'s [classic history API
    state](#concept-navigateevent-classic-history-api-state){#navigate-event-firing:concept-navigateevent-classic-history-api-state}
    to null.

3.  Let `destination`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#navigate-event-firing:new
    x-internal="new"}
    [`NavigationDestination`{#navigate-event-firing:navigationdestination}](#navigationdestination)
    created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-2}.

4.  Set `destination`{.variable}\'s
    [URL](#concept-navigationdestination-url){#navigate-event-firing:concept-navigationdestination-url}
    to `destinationSHE`{.variable}\'s
    [URL](browsing-the-web.html#she-url){#navigate-event-firing:she-url}.

5.  Let `destinationNHE`{.variable} be the
    [`NavigationHistoryEntry`{#navigate-event-firing:navigationhistoryentry}](#navigationhistoryentry)
    in `navigation`{.variable}\'s [entry
    list](#navigation-entry-list){#navigate-event-firing:navigation-entry-list}
    whose [session history
    entry](#nhe-she){#navigate-event-firing:nhe-she} is
    `destinationSHE`{.variable}, or null if no such
    [`NavigationHistoryEntry`{#navigate-event-firing:navigationhistoryentry-2}](#navigationhistoryentry)
    exists.

6.  If `destinationNHE`{.variable} is non-null, then:

    1.  Set `destination`{.variable}\'s
        [entry](#concept-navigationdestination-entry){#navigate-event-firing:concept-navigationdestination-entry}
        to `destinationNHE`{.variable}.

    2.  Set `destination`{.variable}\'s
        [state](#concept-navigationdestination-state){#navigate-event-firing:concept-navigationdestination-state}
        to `destinationSHE`{.variable}\'s [navigation API
        state](browsing-the-web.html#she-navigation-api-state){#navigate-event-firing:she-navigation-api-state}.

7.  Otherwise,

    1.  Set `destination`{.variable}\'s
        [entry](#concept-navigationdestination-entry){#navigate-event-firing:concept-navigationdestination-entry-2}
        to null.

    2.  Set `destination`{.variable}\'s
        [state](#concept-navigationdestination-state){#navigate-event-firing:concept-navigationdestination-state-2}
        to
        [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#navigate-event-firing:structuredserializeforstorage}(null).

8.  Set `destination`{.variable}\'s [is same
    document](#concept-navigationdestination-samedocument){#navigate-event-firing:concept-navigationdestination-samedocument}
    to true if `destinationSHE`{.variable}\'s
    [document](browsing-the-web.html#she-document){#navigate-event-firing:she-document}
    is equal to `navigation`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#navigate-event-firing:concept-relevant-global}\'s
    [associated
    `Document`](#concept-document-window){#navigate-event-firing:concept-document-window};
    otherwise false.

9.  Return the result of performing the [inner `navigate` event firing
    algorithm](#inner-navigate-event-firing-algorithm){#navigate-event-firing:inner-navigate-event-firing-algorithm}
    given `navigation`{.variable},
    \"[`traverse`{#navigate-event-firing:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\",
    `event`{.variable}, `destination`{.variable},
    `userInvolvement`{.variable}, null, null, and null.

To [fire a push/replace/reload `navigate`
event]{#fire-a-push/replace/reload-navigate-event .dfn} at a
[`Navigation`{#navigate-event-firing:navigation-2}](#navigation)
`navigation`{.variable} given a
[`NavigationType`{#navigate-event-firing:navigationtype}](#navigationtype)
[`navigationType`{.variable}]{#fire-navigate-prr-navigationtype .dfn}, a
[URL](https://url.spec.whatwg.org/#concept-url){#navigate-event-firing:url
x-internal="url"}
[`destinationURL`{.variable}]{#fire-navigate-prr-destinationurl .dfn}, a
boolean [`isSameDocument`{.variable}]{#fire-navigate-prr-issamedocument
.dfn}, an optional [user navigation
involvement](browsing-the-web.html#user-navigation-involvement){#navigate-event-firing:user-navigation-involvement-2}
[`userInvolvement`{.variable}]{#fire-navigate-prr-userinvolvement .dfn}
(default
\"[`none`{#navigate-event-firing:uni-none-2}](browsing-the-web.html#uni-none)\"),
an optional
[`Element`{#navigate-event-firing:element}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}-or-null
[`sourceElement`{.variable}]{#fire-navigate-prr-sourceelement .dfn}
(default null), an optional [entry
list](form-control-infrastructure.html#entry-list){#navigate-event-firing:entry-list}-or-null
[`formDataEntryList`{.variable}]{#fire-navigate-prr-formdataentrylist
.dfn} (default null), an optional [serialized
state](browsing-the-web.html#serialized-state){#navigate-event-firing:serialized-state}
[`navigationAPIState`{.variable}]{#fire-navigate-prr-navigationapistate
.dfn} (default
[StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#navigate-event-firing:structuredserializeforstorage-2}(null)),
and an optional [serialized
state](browsing-the-web.html#serialized-state){#navigate-event-firing:serialized-state-2}-or-null
[`classicHistoryAPIState`{.variable}]{#fire-navigate-prr-classichistoryapistate
.dfn} (default null):

1.  Let `event`{.variable} be the result of [creating an
    event](https://dom.spec.whatwg.org/#concept-event-create){#navigate-event-firing:creating-an-event-2
    x-internal="creating-an-event"} given
    [`NavigateEvent`{#navigate-event-firing:navigateevent-2}](#navigateevent),
    in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-3}.

2.  Set `event`{.variable}\'s [classic history API
    state](#concept-navigateevent-classic-history-api-state){#navigate-event-firing:concept-navigateevent-classic-history-api-state-2}
    to `classicHistoryAPIState`{.variable}.

3.  Let `destination`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#navigate-event-firing:new-2
    x-internal="new"}
    [`NavigationDestination`{#navigate-event-firing:navigationdestination-2}](#navigationdestination)
    created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-4}.

4.  Set `destination`{.variable}\'s
    [URL](#concept-navigationdestination-url){#navigate-event-firing:concept-navigationdestination-url-2}
    to `destinationURL`{.variable}.

5.  Set `destination`{.variable}\'s
    [entry](#concept-navigationdestination-entry){#navigate-event-firing:concept-navigationdestination-entry-3}
    to null.

6.  Set `destination`{.variable}\'s
    [state](#concept-navigationdestination-state){#navigate-event-firing:concept-navigationdestination-state-3}
    to `navigationAPIState`{.variable}.

7.  Set `destination`{.variable}\'s [is same
    document](#concept-navigationdestination-samedocument){#navigate-event-firing:concept-navigationdestination-samedocument-2}
    to `isSameDocument`{.variable}.

8.  Return the result of performing the [inner `navigate` event firing
    algorithm](#inner-navigate-event-firing-algorithm){#navigate-event-firing:inner-navigate-event-firing-algorithm-2}
    given `navigation`{.variable}, `navigationType`{.variable},
    `event`{.variable}, `destination`{.variable},
    `userInvolvement`{.variable}, `sourceElement`{.variable},
    `formDataEntryList`{.variable}, and null.

To [fire a download request `navigate`
event]{#fire-a-download-request-navigate-event .dfn} at a
[`Navigation`{#navigate-event-firing:navigation-3}](#navigation)
`navigation`{.variable} given a
[URL](https://url.spec.whatwg.org/#concept-url){#navigate-event-firing:url-2
x-internal="url"}
[`destinationURL`{.variable}]{#fire-navigate-download-destinationurl
.dfn}, a [user navigation
involvement](browsing-the-web.html#user-navigation-involvement){#navigate-event-firing:user-navigation-involvement-3}
[`userInvolvement`{.variable}]{#fire-navigate-download-userinvolvement
.dfn}, an
[`Element`{#navigate-event-firing:element-2}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}-or-null
[`sourceElement`{.variable}]{#fire-navigate-download-sourceelement
.dfn}, and a string
[`filename`{.variable}]{#fire-navigate-download-filename .dfn}:

1.  Let `event`{.variable} be the result of [creating an
    event](https://dom.spec.whatwg.org/#concept-event-create){#navigate-event-firing:creating-an-event-3
    x-internal="creating-an-event"} given
    [`NavigateEvent`{#navigate-event-firing:navigateevent-3}](#navigateevent),
    in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-5}.

2.  Set `event`{.variable}\'s [classic history API
    state](#concept-navigateevent-classic-history-api-state){#navigate-event-firing:concept-navigateevent-classic-history-api-state-3}
    to null.

3.  Let `destination`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#navigate-event-firing:new-3
    x-internal="new"}
    [`NavigationDestination`{#navigate-event-firing:navigationdestination-3}](#navigationdestination)
    created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-6}.

4.  Set `destination`{.variable}\'s
    [URL](#concept-navigationdestination-url){#navigate-event-firing:concept-navigationdestination-url-3}
    to `destinationURL`{.variable}.

5.  Set `destination`{.variable}\'s
    [entry](#concept-navigationdestination-entry){#navigate-event-firing:concept-navigationdestination-entry-4}
    to null.

6.  Set `destination`{.variable}\'s
    [state](#concept-navigationdestination-state){#navigate-event-firing:concept-navigationdestination-state-4}
    to
    [StructuredSerializeForStorage](structured-data.html#structuredserializeforstorage){#navigate-event-firing:structuredserializeforstorage-3}(null).

7.  Set `destination`{.variable}\'s [is same
    document](#concept-navigationdestination-samedocument){#navigate-event-firing:concept-navigationdestination-samedocument-3}
    to false.

8.  Return the result of performing the [inner `navigate` event firing
    algorithm](#inner-navigate-event-firing-algorithm){#navigate-event-firing:inner-navigate-event-firing-algorithm-3}
    given `navigation`{.variable},
    \"[`push`{#navigate-event-firing:dom-navigationtype-push}](#dom-navigationtype-push)\",
    `event`{.variable}, `destination`{.variable},
    `userInvolvement`{.variable}, `sourceElement`{.variable}, null, and
    `filename`{.variable}.

The [inner `navigate` event firing
algorithm]{#inner-navigate-event-firing-algorithm .dfn} consists of the
following steps, given a
[`Navigation`{#navigate-event-firing:navigation-4}](#navigation)
`navigation`{.variable}, a
[`NavigationType`{#navigate-event-firing:navigationtype-2}](#navigationtype)
`navigationType`{.variable}, a
[`NavigateEvent`{#navigate-event-firing:navigateevent-4}](#navigateevent)
`event`{.variable}, a
[`NavigationDestination`{#navigate-event-firing:navigationdestination-4}](#navigationdestination)
`destination`{.variable}, a [user navigation
involvement](browsing-the-web.html#user-navigation-involvement){#navigate-event-firing:user-navigation-involvement-4}
`userInvolvement`{.variable}, an
[`Element`{#navigate-event-firing:element-3}](https://dom.spec.whatwg.org/#interface-element){x-internal="element"}-or-null
`sourceElement`{.variable}, an [entry
list](form-control-infrastructure.html#entry-list){#navigate-event-firing:entry-list-2}-or-null
`formDataEntryList`{.variable}, and a string-or-null
`downloadRequestFilename`{.variable}:

1.  If `navigation`{.variable} [has entries and events
    disabled](#has-entries-and-events-disabled){#navigate-event-firing:has-entries-and-events-disabled},
    then:

    1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert
        x-internal="assert"}: `navigation`{.variable}\'s [ongoing API
        method
        tracker](#ongoing-api-method-tracker){#navigate-event-firing:ongoing-api-method-tracker}
        is null.

    2.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert-2
        x-internal="assert"}: `navigation`{.variable}\'s [upcoming
        non-traverse API method
        tracker](#upcoming-non-traverse-api-method-tracker){#navigate-event-firing:upcoming-non-traverse-api-method-tracker}
        is null.

    3.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert-3
        x-internal="assert"}: `navigation`{.variable}\'s [upcoming
        traverse API method
        trackers](#upcoming-traverse-api-method-trackers){#navigate-event-firing:upcoming-traverse-api-method-trackers}
        [is
        empty](https://infra.spec.whatwg.org/#list-is-empty){#navigate-event-firing:list-is-empty
        x-internal="list-is-empty"}.

    4.  Return true.

    These assertions holds because
    [`traverseTo()`{#navigate-event-firing:dom-navigation-traverseto}](#dom-navigation-traverseto),
    [`back()`{#navigate-event-firing:dom-navigation-back}](#dom-navigation-back),
    and
    [`forward()`{#navigate-event-firing:dom-navigation-forward}](#dom-navigation-forward)
    will immediately fail when entries and events are disabled (since
    there are no entries to traverse to), and if our starting point is
    instead
    [`navigate()`{#navigate-event-firing:dom-navigation-navigate}](#dom-navigation-navigate)
    or
    [`reload()`{#navigate-event-firing:dom-navigation-reload}](#dom-navigation-reload),
    then we
    [avoided](#dont-always-set-upcoming-non-traverse-api-method-tracker)
    setting the [upcoming non-traverse API method
    tracker](#upcoming-non-traverse-api-method-tracker){#navigate-event-firing:upcoming-non-traverse-api-method-tracker-2}
    in the first place.

2.  Let `destinationKey`{.variable} be null.

3.  If `destination`{.variable}\'s
    [entry](#concept-navigationdestination-entry){#navigate-event-firing:concept-navigationdestination-entry-5}
    is non-null, then set `destinationKey`{.variable} to
    `destination`{.variable}\'s
    [entry](#concept-navigationdestination-entry){#navigate-event-firing:concept-navigationdestination-entry-6}\'s
    [key](#concept-navigationhistoryentry-key){#navigate-event-firing:concept-navigationhistoryentry-key}.

4.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert-4
    x-internal="assert"}: `destinationKey`{.variable} is not the empty
    string.

5.  [Promote an upcoming API method tracker to
    ongoing](#promote-an-upcoming-api-method-tracker-to-ongoing){#navigate-event-firing:promote-an-upcoming-api-method-tracker-to-ongoing}
    given `navigation`{.variable} and `destinationKey`{.variable}.

6.  Let `apiMethodTracker`{.variable} be `navigation`{.variable}\'s
    [ongoing API method
    tracker](#ongoing-api-method-tracker){#navigate-event-firing:ongoing-api-method-tracker-2}.

7.  Let `navigable`{.variable} be `navigation`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#navigate-event-firing:concept-relevant-global-2}\'s
    [navigable](#window-navigable){#navigate-event-firing:window-navigable}.

8.  Let `document`{.variable} be `navigation`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#navigate-event-firing:concept-relevant-global-3}\'s
    [associated
    `Document`](#concept-document-window){#navigate-event-firing:concept-document-window-2}.

9.  If `document`{.variable} [can have its URL
    rewritten](#can-have-its-url-rewritten){#navigate-event-firing:can-have-its-url-rewritten}
    to `destination`{.variable}\'s
    [URL](#concept-navigationdestination-url){#navigate-event-firing:concept-navigationdestination-url-4},
    and either `destination`{.variable}\'s [is same
    document](#concept-navigationdestination-samedocument){#navigate-event-firing:concept-navigationdestination-samedocument-4}
    is true or `navigationType`{.variable} is not
    \"[`traverse`{#navigate-event-firing:dom-navigationtype-traverse-2}](#dom-navigationtype-traverse)\",
    then initialize `event`{.variable}\'s
    [`canIntercept`{#navigate-event-firing:dom-navigateevent-canintercept}](#dom-navigateevent-canintercept)
    to true. Otherwise, initialize it to false.

10. ::: {#navigate-event-traverse-can-be-canceled}
    Let `traverseCanBeCanceled`{.variable} be true if all of the
    following are true:

    - `navigable`{.variable} is a [top-level
      traversable](document-sequences.html#top-level-traversable){#navigate-event-firing:top-level-traversable};

    - `destination`{.variable}\'s [is same
      document](#concept-navigationdestination-samedocument){#navigate-event-firing:concept-navigationdestination-samedocument-5}
      is true; and

    - either `userInvolvement`{.variable} is not
      \"[`browser UI`{#navigate-event-firing:uni-browser-ui}](browsing-the-web.html#uni-browser-ui)\",
      or `navigation`{.variable}\'s [relevant global
      object](webappapis.html#concept-relevant-global){#navigate-event-firing:concept-relevant-global-4}
      has [history-action
      activation](interaction.html#history-action-activation){#navigate-event-firing:history-action-activation}.

    Otherwise, let it be false.
    :::

11. If either:

    - `navigationType`{.variable} is not
      \"[`traverse`{#navigate-event-firing:dom-navigationtype-traverse-3}](#dom-navigationtype-traverse)\";
      or

    - `traverseCanBeCanceled`{.variable} is true,

    then initialize `event`{.variable}\'s
    [`cancelable`{#navigate-event-firing:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
    to true. Otherwise, initialize it to false.

12. Initialize `event`{.variable}\'s
    [`type`{#navigate-event-firing:dom-event-type}](https://dom.spec.whatwg.org/#dom-event-type){x-internal="dom-event-type"}
    to
    \"[`navigate`{#navigate-event-firing:event-navigate-2}](indices.html#event-navigate)\".

13. Initialize `event`{.variable}\'s
    [`navigationType`{#navigate-event-firing:dom-navigateevent-navigationtype}](#dom-navigateevent-navigationtype)
    to `navigationType`{.variable}.

14. Initialize `event`{.variable}\'s
    [`destination`{#navigate-event-firing:dom-navigateevent-destination}](#dom-navigateevent-destination)
    to `destination`{.variable}.

15. Initialize `event`{.variable}\'s
    [`downloadRequest`{#navigate-event-firing:dom-navigateevent-downloadrequest}](#dom-navigateevent-downloadrequest)
    to `downloadRequestFilename`{.variable}.

16. If `apiMethodTracker`{.variable} is not null, then initialize
    `event`{.variable}\'s
    [`info`{#navigate-event-firing:dom-navigateevent-info}](#dom-navigateevent-info)
    to `apiMethodTracker`{.variable}\'s
    [info](#navigation-api-method-tracker-info){#navigate-event-firing:navigation-api-method-tracker-info}.
    Otherwise, initialize it to undefined.

    At this point `apiMethodTracker`{.variable}\'s
    [info](#navigation-api-method-tracker-info){#navigate-event-firing:navigation-api-method-tracker-info-2}
    is no longer needed and can be nulled out instead of keeping it
    alive for the lifetime of the [navigation API method
    tracker](#navigation-api-method-tracker){#navigate-event-firing:navigation-api-method-tracker}.

17. Initialize `event`{.variable}\'s
    [`hasUAVisualTransition`{#navigate-event-firing:dom-navigateevent-hasuavisualtransition}](#dom-navigateevent-hasuavisualtransition)
    to true if a visual transition, to display a cached rendered state
    of the `document`{.variable}\'s [latest
    entry](browsing-the-web.html#latest-entry){#navigate-event-firing:latest-entry},
    was done by the user agent. Otherwise, initialize it to false.

18. Initialize `event`{.variable}\'s
    [`sourceElement`{#navigate-event-firing:dom-navigateevent-sourceelement}](#dom-navigateevent-sourceelement)
    to `sourceElement`{.variable}.

19. Set `event`{.variable}\'s [abort
    controller](#concept-navigateevent-abort-controller){#navigate-event-firing:concept-navigateevent-abort-controller}
    to a
    [new](https://webidl.spec.whatwg.org/#new){#navigate-event-firing:new-4
    x-internal="new"}
    [`AbortController`{#navigate-event-firing:abortcontroller}](https://dom.spec.whatwg.org/#abortcontroller){x-internal="abortcontroller"}
    created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-7}.

20. Initialize `event`{.variable}\'s
    [`signal`{#navigate-event-firing:dom-navigateevent-signal}](#dom-navigateevent-signal)
    to `event`{.variable}\'s [abort
    controller](#concept-navigateevent-abort-controller){#navigate-event-firing:concept-navigateevent-abort-controller-2}\'s
    [signal](https://dom.spec.whatwg.org/#abortcontroller-signal){#navigate-event-firing:concept-abortcontroller-signal
    x-internal="concept-abortcontroller-signal"}.

21. Let `currentURL`{.variable} be `document`{.variable}\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#navigate-event-firing:the-document's-address
    x-internal="the-document's-address"}.

22. If all of the following are true:

    - `event`{.variable}\'s [classic history API
      state](#concept-navigateevent-classic-history-api-state){#navigate-event-firing:concept-navigateevent-classic-history-api-state-4}
      is null;

    - `destination`{.variable}\'s [is same
      document](#concept-navigationdestination-samedocument){#navigate-event-firing:concept-navigationdestination-samedocument-6}
      is true;

    - `destination`{.variable}\'s
      [URL](#concept-navigationdestination-url){#navigate-event-firing:concept-navigationdestination-url-5}
      [equals](https://url.spec.whatwg.org/#concept-url-equals){#navigate-event-firing:concept-url-equals
      x-internal="concept-url-equals"} `currentURL`{.variable} with
      [*exclude
      fragments*](https://url.spec.whatwg.org/#url-equals-exclude-fragments){#navigate-event-firing:url-equals-exclude-fragments
      x-internal="url-equals-exclude-fragments"} set to true; and

    - `destination`{.variable}\'s
      [URL](#concept-navigationdestination-url){#navigate-event-firing:concept-navigationdestination-url-6}\'s
      [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#navigate-event-firing:concept-url-fragment
      x-internal="concept-url-fragment"} is not [identical
      to](https://infra.spec.whatwg.org/#string-is){#navigate-event-firing:identical-to
      x-internal="identical-to"} `currentURL`{.variable}\'s
      [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#navigate-event-firing:concept-url-fragment-2
      x-internal="concept-url-fragment"},

    then initialize `event`{.variable}\'s
    [`hashChange`{#navigate-event-firing:dom-navigateevent-hashchange}](#dom-navigateevent-hashchange)
    to true. Otherwise, initialize it to false.

    The first condition here means that
    [`hashChange`{#navigate-event-firing:dom-navigateevent-hashchange-2}](#dom-navigateevent-hashchange)
    will be true for [fragment
    navigations](browsing-the-web.html#navigate-fragid){#navigate-event-firing:navigate-fragid},
    but false for cases like
    `history.pushState(undefined, "", "#fragment")`.

23. If `userInvolvement`{.variable} is not
    \"[`none`{#navigate-event-firing:uni-none-3}](browsing-the-web.html#uni-none)\",
    then initialize `event`{.variable}\'s
    [`userInitiated`{#navigate-event-firing:dom-navigateevent-userinitiated}](#dom-navigateevent-userinitiated)
    to true. Otherwise, initialize it to false.

24. If `formDataEntryList`{.variable} is not null, then initialize
    `event`{.variable}\'s
    [`formData`{#navigate-event-firing:dom-navigateevent-formdata}](#dom-navigateevent-formdata)
    to a
    [new](https://webidl.spec.whatwg.org/#new){#navigate-event-firing:new-5
    x-internal="new"}
    [`FormData`{#navigate-event-firing:formdata}](https://xhr.spec.whatwg.org/#formdata){x-internal="formdata"}
    created in `navigation`{.variable}\'s [relevant
    realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-8},
    associated to `formDataEntryList`{.variable}. Otherwise, initialize
    it to null.

25. [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert-5
    x-internal="assert"}: `navigation`{.variable}\'s [ongoing `navigate`
    event](#ongoing-navigate-event){#navigate-event-firing:ongoing-navigate-event}
    is null.

26. Set `navigation`{.variable}\'s [ongoing `navigate`
    event](#ongoing-navigate-event){#navigate-event-firing:ongoing-navigate-event-2}
    to `event`{.variable}.

27. Set `navigation`{.variable}\'s [focus changed during ongoing
    navigation](#focus-changed-during-ongoing-navigation){#navigate-event-firing:focus-changed-during-ongoing-navigation}
    to false.

28. Set `navigation`{.variable}\'s [suppress normal scroll restoration
    during ongoing
    navigation](#suppress-normal-scroll-restoration-during-ongoing-navigation){#navigate-event-firing:suppress-normal-scroll-restoration-during-ongoing-navigation}
    to false.

29. Let `dispatchResult`{.variable} be the result of
    [dispatching](https://dom.spec.whatwg.org/#concept-event-dispatch){#navigate-event-firing:concept-event-dispatch
    x-internal="concept-event-dispatch"} `event`{.variable} at
    `navigation`{.variable}.

30. If `dispatchResult`{.variable} is false:

    1.  If `navigationType`{.variable} is
        \"[`traverse`{#navigate-event-firing:dom-navigationtype-traverse-4}](#dom-navigationtype-traverse)\",
        then [consume history-action user
        activation](interaction.html#consume-history-action-user-activation){#navigate-event-firing:consume-history-action-user-activation}
        given `navigation`{.variable}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#navigate-event-firing:concept-relevant-global-5}.

    2.  If `event`{.variable}\'s [abort
        controller](#concept-navigateevent-abort-controller){#navigate-event-firing:concept-navigateevent-abort-controller-3}\'s
        [signal](https://dom.spec.whatwg.org/#abortcontroller-signal){#navigate-event-firing:concept-abortcontroller-signal-2
        x-internal="concept-abortcontroller-signal"} is not
        [aborted](https://dom.spec.whatwg.org/#abortsignal-aborted){#navigate-event-firing:abortsignal-aborted
        x-internal="abortsignal-aborted"}, then [abort the ongoing
        navigation](#abort-the-ongoing-navigation){#navigate-event-firing:abort-the-ongoing-navigation}
        given `navigation`{.variable}.

    3.  Return false.

31. Let `endResultIsSameDocument`{.variable} be true if
    `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-firing:concept-navigateevent-interception-state}
    is not \"`none`\" or `event`{.variable}\'s
    [`destination`{#navigate-event-firing:dom-navigateevent-destination-2}](#dom-navigateevent-destination)\'s
    [is same
    document](#concept-navigationdestination-samedocument){#navigate-event-firing:concept-navigationdestination-samedocument-7}
    is true.

32. [Prepare to run
    script](webappapis.html#prepare-to-run-script){#navigate-event-firing:prepare-to-run-script}
    given `navigation`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#navigate-event-firing:relevant-settings-object}.

    ::: {#note-suppress-microtasks-during-navigation-events .note}
    This is done to avoid the [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#navigate-event-firing:javascript-execution-context-stack
    x-internal="javascript-execution-context-stack"} becoming empty
    right after any
    [`currententrychange`{#navigate-event-firing:event-currententrychange}](indices.html#event-currententrychange)
    event handlers run as a result of the [URL and history update
    steps](browsing-the-web.html#url-and-history-update-steps){#navigate-event-firing:url-and-history-update-steps}
    that could soon happen. If the stack were to become empty at that
    time, then it would immediately [perform a microtask
    checkpoint](webappapis.html#perform-a-microtask-checkpoint){#navigate-event-firing:perform-a-microtask-checkpoint},
    causing various promise fulfillment handlers to run interleaved with
    the event handlers and before any handlers passed to
    [`navigateEvent.intercept()`{#navigate-event-firing:dom-navigateevent-intercept}](#dom-navigateevent-intercept).
    This is undesirable since it means promise handler ordering vs.
    [`currententrychange`{#navigate-event-firing:event-currententrychange-2}](indices.html#event-currententrychange)
    event handler ordering vs.
    [`intercept()`{#navigate-event-firing:dom-navigateevent-intercept-2}](#dom-navigateevent-intercept)
    handler ordering would be dependent on whether the navigation is
    happening with an empty [JavaScript execution context
    stack](https://tc39.es/ecma262/#execution-context-stack){#navigate-event-firing:javascript-execution-context-stack-2
    x-internal="javascript-execution-context-stack"} (e.g., because the
    navigation was user-initiated) or with a nonempty one (e.g., because
    the navigation was caused by a JavaScript API call).

    By inserting an otherwise-unnecessary [JavaScript execution
    context](https://tc39.es/ecma262/#sec-execution-contexts){#navigate-event-firing:javascript-execution-context
    x-internal="javascript-execution-context"} onto the stack in this
    step, we essentially suppress the [perform a microtask
    checkpoint](webappapis.html#perform-a-microtask-checkpoint){#navigate-event-firing:perform-a-microtask-checkpoint-2}
    algorithm until later, thus ensuring that the sequence is always:
    [`currententrychange`{#navigate-event-firing:event-currententrychange-3}](indices.html#event-currententrychange)
    event handlers, then
    [`intercept()`{#navigate-event-firing:dom-navigateevent-intercept-3}](#dom-navigateevent-intercept)
    handlers, then promise handlers.
    :::

33. If `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-firing:concept-navigateevent-interception-state-2}
    is not \"`none`\":

    1.  Set `event`{.variable}\'s [interception
        state](#concept-navigateevent-interception-state){#navigate-event-firing:concept-navigateevent-interception-state-3}
        to \"`committed`\".

    2.  Let `fromNHE`{.variable} be the [current
        entry](#navigation-current-entry){#navigate-event-firing:navigation-current-entry}
        of `navigation`{.variable}.

    3.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert-6
        x-internal="assert"}: `fromNHE`{.variable} is not null.

    4.  Set `navigation`{.variable}\'s
        [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition}
        to a
        [new](https://webidl.spec.whatwg.org/#new){#navigate-event-firing:new-6
        x-internal="new"}
        [`NavigationTransition`{#navigate-event-firing:navigationtransition}](#navigationtransition)
        created in `navigation`{.variable}\'s [relevant
        realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-9},
        whose [navigation
        type](#concept-navigationtransition-navigationtype){#navigate-event-firing:concept-navigationtransition-navigationtype}
        is `navigationType`{.variable}, whose [from
        entry](#concept-navigationtransition-from){#navigate-event-firing:concept-navigationtransition-from}
        is `fromNHE`{.variable}, and whose [finished
        promise](#concept-navigationtransition-finished){#navigate-event-firing:concept-navigationtransition-finished}
        is a new promise created in `navigation`{.variable}\'s [relevant
        realm](webappapis.html#concept-relevant-realm){#navigate-event-firing:concept-relevant-realm-10}.

    5.  [Mark as
        handled](https://webidl.spec.whatwg.org/#mark-a-promise-as-handled){#navigate-event-firing:mark-as-handled
        x-internal="mark-as-handled"} `navigation`{.variable}\'s
        [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition-2}\'s
        [finished
        promise](#concept-navigationtransition-finished){#navigate-event-firing:concept-navigationtransition-finished-2}.

        See the [discussion about other finished
        promises](#note-mark-as-handled-navigation-api-finished) to
        understand why this is done.

    6.  If `navigationType`{.variable} is
        \"[`traverse`{#navigate-event-firing:dom-navigationtype-traverse-5}](#dom-navigationtype-traverse)\",
        then set `navigation`{.variable}\'s [suppress normal scroll
        restoration during ongoing
        navigation](#suppress-normal-scroll-restoration-during-ongoing-navigation){#navigate-event-firing:suppress-normal-scroll-restoration-during-ongoing-navigation-2}
        to true.

        If `event`{.variable}\'s [scroll
        behavior](#concept-navigateevent-scroll){#navigate-event-firing:concept-navigateevent-scroll}
        was set to
        \"[`after-transition`{#navigate-event-firing:dom-navigationscrollbehavior-after-transition}](#dom-navigationscrollbehavior-after-transition)\",
        then scroll restoration will happen as part of
        [finishing](#navigateevent-finish){#navigate-event-firing:navigateevent-finish}
        the relevant
        [`NavigateEvent`{#navigate-event-firing:navigateevent-5}](#navigateevent).
        Otherwise, there will be no scroll restoration. That is, no
        navigation which is intercepted by
        [`intercept()`{#navigate-event-firing:dom-navigateevent-intercept-4}](#dom-navigateevent-intercept)
        goes through the normal scroll restoration process; scroll
        restoration for such navigations is either done manually, by the
        web developer, or is done after the transition.

    7.  If `navigationType`{.variable} is
        \"[`push`{#navigate-event-firing:dom-navigationtype-push-2}](#dom-navigationtype-push)\"
        or
        \"[`replace`{#navigate-event-firing:dom-navigationtype-replace}](#dom-navigationtype-replace)\",
        then run the [URL and history update
        steps](browsing-the-web.html#url-and-history-update-steps){#navigate-event-firing:url-and-history-update-steps-2}
        given `document`{.variable} and `event`{.variable}\'s
        [`destination`{#navigate-event-firing:dom-navigateevent-destination-3}](#dom-navigateevent-destination)\'s
        [URL](#dom-navigationdestination-url){#navigate-event-firing:dom-navigationdestination-url},
        with
        *[serializedData](browsing-the-web.html#uhus-serializeddata)*
        set to `event`{.variable}\'s [classic history API
        state](#concept-navigateevent-classic-history-api-state){#navigate-event-firing:concept-navigateevent-classic-history-api-state-5}
        and
        *[historyHandling](browsing-the-web.html#uhus-historyhandling)*
        set to `navigationType`{.variable}.

    8.  Otherwise, if `navigationType`{.variable} is
        \"[`reload`{#navigate-event-firing:dom-navigationtype-reload}](#dom-navigationtype-reload)\",
        then [update the navigation API entries for a same-document
        navigation](#update-the-navigation-api-entries-for-a-same-document-navigation){#navigate-event-firing:update-the-navigation-api-entries-for-a-same-document-navigation}
        given `navigation`{.variable}, `navigable`{.variable}\'s [active
        session history
        entry](document-sequences.html#nav-active-history-entry){#navigate-event-firing:nav-active-history-entry},
        and
        \"[`reload`{#navigate-event-firing:dom-navigationtype-reload-2}](#dom-navigationtype-reload)\".

        If `navigationType`{.variable} is
        \"[`traverse`{#navigate-event-firing:dom-navigationtype-traverse-6}](#dom-navigationtype-traverse)\",
        then this event firing is happening as part of [the traversal
        process](browsing-the-web.html#apply-the-traverse-history-step){#navigate-event-firing:apply-the-traverse-history-step},
        and that process will take care of performing the appropriate
        session history entry updates.

34. If `endResultIsSameDocument`{.variable} is true:

    1.  Let `promisesList`{.variable} be an empty
        [list](https://infra.spec.whatwg.org/#list){#navigate-event-firing:list
        x-internal="list"}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#navigate-event-firing:list-iterate
        x-internal="list-iterate"} `handler`{.variable} of
        `event`{.variable}\'s [navigation handler
        list](#concept-navigateevent-navigation-handler-list){#navigate-event-firing:concept-navigateevent-navigation-handler-list}:

        1.  [Append](https://infra.spec.whatwg.org/#list-append){#navigate-event-firing:list-append
            x-internal="list-append"} the result of
            [invoking](https://webidl.spec.whatwg.org/#invoke-a-callback-function){#navigate-event-firing:es-invoking-callback-functions
            x-internal="es-invoking-callback-functions"}
            `handler`{.variable} with an empty arguments list to
            `promisesList`{.variable}.

    3.  If `promisesList`{.variable}\'s
        [size](https://infra.spec.whatwg.org/#list-size){#navigate-event-firing:list-size
        x-internal="list-size"} is 0, then set `promisesList`{.variable}
        to « [a promise resolved
        with](https://webidl.spec.whatwg.org/#a-promise-resolved-with){#navigate-event-firing:a-promise-resolved-with
        x-internal="a-promise-resolved-with"} undefined ».

        There is a subtle timing difference between how [waiting for
        all](https://webidl.spec.whatwg.org/#wait-for-all){#navigate-event-firing:wait-for-all
        x-internal="wait-for-all"} schedules its success and failure
        steps when given zero promises versus ≥1 promises. For most uses
        of [waiting for
        all](https://webidl.spec.whatwg.org/#wait-for-all){#navigate-event-firing:wait-for-all-2
        x-internal="wait-for-all"}, this does not matter. However, with
        this API, there are so many events and promise handlers which
        could fire around the same time that the difference is pretty
        easily observable: it can cause the event/promise handler
        sequence to vary. (Some of the events and promises involved
        include:
        [`navigatesuccess`{#navigate-event-firing:event-navigatesuccess}](indices.html#event-navigatesuccess)
        /
        [`navigateerror`{#navigate-event-firing:event-navigateerror}](indices.html#event-navigateerror),
        [`currententrychange`{#navigate-event-firing:event-currententrychange-4}](indices.html#event-currententrychange),
        [`dispose`{#navigate-event-firing:event-dispose}](indices.html#event-dispose),
        `apiMethodTracker`{.variable}\'s promises, and the
        [`navigation.transition.finished`{#navigate-event-firing:dom-navigationtransition-finished}](#dom-navigationtransition-finished)
        promise.)

    4.  [Wait for
        all](https://webidl.spec.whatwg.org/#wait-for-all){#navigate-event-firing:wait-for-all-3
        x-internal="wait-for-all"} of `promisesList`{.variable}, with
        the following success steps:

        1.  If `event`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#navigate-event-firing:concept-relevant-global-6}\'s
            [associated
            `Document`](#concept-document-window){#navigate-event-firing:concept-document-window-3}
            is not [fully
            active](document-sequences.html#fully-active){#navigate-event-firing:fully-active},
            then abort these steps.

        2.  If `event`{.variable}\'s [abort
            controller](#concept-navigateevent-abort-controller){#navigate-event-firing:concept-navigateevent-abort-controller-4}\'s
            [signal](https://dom.spec.whatwg.org/#abortcontroller-signal){#navigate-event-firing:concept-abortcontroller-signal-3
            x-internal="concept-abortcontroller-signal"} is
            [aborted](https://dom.spec.whatwg.org/#abortsignal-aborted){#navigate-event-firing:abortsignal-aborted-2
            x-internal="abortsignal-aborted"}, then abort these steps.

        3.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert-7
            x-internal="assert"}: `event`{.variable} equals
            `navigation`{.variable}\'s [ongoing `navigate`
            event](#ongoing-navigate-event){#navigate-event-firing:ongoing-navigate-event-3}.

        4.  Set `navigation`{.variable}\'s [ongoing `navigate`
            event](#ongoing-navigate-event){#navigate-event-firing:ongoing-navigate-event-4}
            to null.

        5.  [Finish](#navigateevent-finish){#navigate-event-firing:navigateevent-finish-2}
            `event`{.variable} given true.

        6.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#navigate-event-firing:concept-event-fire
            x-internal="concept-event-fire"} named
            [`navigatesuccess`{#navigate-event-firing:event-navigatesuccess-2}](indices.html#event-navigatesuccess)
            at `navigation`{.variable}.

        7.  If `apiMethodTracker`{.variable} is non-null, then [resolve
            the finished
            promise](#resolve-the-finished-promise){#navigate-event-firing:resolve-the-finished-promise}
            for `apiMethodTracker`{.variable}.

        8.  If `navigation`{.variable}\'s
            [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition-3}
            is not null, then resolve `navigation`{.variable}\'s
            [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition-4}\'s
            [finished
            promise](#concept-navigationtransition-finished){#navigate-event-firing:concept-navigationtransition-finished-3}
            with undefined.

        9.  Set `navigation`{.variable}\'s
            [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition-5}
            to null.

        and the following failure steps given reason
        `rejectionReason`{.variable}:

        1.  If `event`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#navigate-event-firing:concept-relevant-global-7}\'s
            [associated
            `Document`](#concept-document-window){#navigate-event-firing:concept-document-window-4}
            is not [fully
            active](document-sequences.html#fully-active){#navigate-event-firing:fully-active-2},
            then abort these steps.

        2.  If `event`{.variable}\'s [abort
            controller](#concept-navigateevent-abort-controller){#navigate-event-firing:concept-navigateevent-abort-controller-5}\'s
            [signal](https://dom.spec.whatwg.org/#abortcontroller-signal){#navigate-event-firing:concept-abortcontroller-signal-4
            x-internal="concept-abortcontroller-signal"} is
            [aborted](https://dom.spec.whatwg.org/#abortsignal-aborted){#navigate-event-firing:abortsignal-aborted-3
            x-internal="abortsignal-aborted"}, then abort these steps.

        3.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-firing:assert-8
            x-internal="assert"}: `event`{.variable} equals
            `navigation`{.variable}\'s [ongoing `navigate`
            event](#ongoing-navigate-event){#navigate-event-firing:ongoing-navigate-event-5}.

        4.  Set `navigation`{.variable}\'s [ongoing `navigate`
            event](#ongoing-navigate-event){#navigate-event-firing:ongoing-navigate-event-6}
            to null.

        5.  [Finish](#navigateevent-finish){#navigate-event-firing:navigateevent-finish-3}
            `event`{.variable} given false.

        6.  Let `errorInfo`{.variable} be the result of [extracting
            error
            information](webappapis.html#extract-error){#navigate-event-firing:extract-error}
            from `rejectionReason`{.variable}.

        7.  [Fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#navigate-event-firing:concept-event-fire-2
            x-internal="concept-event-fire"} named
            [`navigateerror`{#navigate-event-firing:event-navigateerror-2}](indices.html#event-navigateerror)
            at `navigation`{.variable} using
            [`ErrorEvent`{#navigate-event-firing:errorevent}](webappapis.html#errorevent),
            with additional attributes initialized according to
            `errorInfo`{.variable}.

        8.  If `apiMethodTracker`{.variable} is non-null, then [reject
            the finished
            promise](#reject-the-finished-promise){#navigate-event-firing:reject-the-finished-promise}
            for `apiMethodTracker`{.variable} with
            `rejectionReason`{.variable}.

        9.  If `navigation`{.variable}\'s
            [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition-6}
            is not null, then reject `navigation`{.variable}\'s
            [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition-7}\'s
            [finished
            promise](#concept-navigationtransition-finished){#navigate-event-firing:concept-navigationtransition-finished-4}
            with `rejectionReason`{.variable}.

        10. Set `navigation`{.variable}\'s
            [transition](#concept-navigation-transition){#navigate-event-firing:concept-navigation-transition-8}
            to null.

35. Otherwise, if `apiMethodTracker`{.variable} is non-null, then [clean
    up](#navigation-api-method-tracker-clean-up){#navigate-event-firing:navigation-api-method-tracker-clean-up}
    `apiMethodTracker`{.variable}.

36. [Clean up after running
    script](webappapis.html#clean-up-after-running-script){#navigate-event-firing:clean-up-after-running-script}
    given `navigation`{.variable}\'s [relevant settings
    object](webappapis.html#relevant-settings-object){#navigate-event-firing:relevant-settings-object-2}.

    Per the [previous
    note](#note-suppress-microtasks-during-navigation-events), this
    stops suppressing any potential promise handler microtasks, causing
    them to run at this point or later.

37. If `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-firing:concept-navigateevent-interception-state-4}
    is \"`none`\", then return true.

38. Return false.

###### [7.2.6.10.4]{.secno} Scroll and focus behavior[](#navigate-event-scroll-focus){.self-link} {#navigate-event-scroll-focus}

By calling
[`navigateEvent.intercept()`{#navigate-event-scroll-focus:dom-navigateevent-intercept}](#dom-navigateevent-intercept),
web developers can suppress the normal scroll and focus behavior for
same-document navigations, instead invoking cross-document
navigation-like behavior at a later time. The algorithms in this section
are called at those appropriate later points.

To [finish]{#navigateevent-finish .dfn} a
[`NavigateEvent`{#navigate-event-scroll-focus:navigateevent}](#navigateevent)
`event`{.variable}, given a boolean `didFulfill`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-scroll-focus:assert
    x-internal="assert"}: `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state}
    is not \"`intercepted`\" or \"`finished`\".

2.  If `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state-2}
    is \"`none`\", then return.

3.  [Potentially reset the
    focus](#potentially-reset-the-focus){#navigate-event-scroll-focus:potentially-reset-the-focus}
    given `event`{.variable}.

4.  If `didFulfill`{.variable} is true, then [potentially process scroll
    behavior](#potentially-process-scroll-behavior){#navigate-event-scroll-focus:potentially-process-scroll-behavior}
    given `event`{.variable}.

5.  Set `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state-3}
    to \"`finished`\".

To [potentially reset the focus]{#potentially-reset-the-focus .dfn}
given a
[`NavigateEvent`{#navigate-event-scroll-focus:navigateevent-2}](#navigateevent)
`event`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-scroll-focus:assert-2
    x-internal="assert"}: `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state-4}
    is \"`committed`\" or \"`scrolled`\".

2.  Let `navigation`{.variable} be `event`{.variable}\'s [relevant
    global
    object](webappapis.html#concept-relevant-global){#navigate-event-scroll-focus:concept-relevant-global}\'s
    [navigation
    API](#window-navigation-api){#navigate-event-scroll-focus:window-navigation-api}.

3.  Let `focusChanged`{.variable} be `navigation`{.variable}\'s [focus
    changed during ongoing
    navigation](#focus-changed-during-ongoing-navigation){#navigate-event-scroll-focus:focus-changed-during-ongoing-navigation}.

4.  Set `navigation`{.variable}\'s [focus changed during ongoing
    navigation](#focus-changed-during-ongoing-navigation){#navigate-event-scroll-focus:focus-changed-during-ongoing-navigation-2}
    to false.

5.  If `focusChanged`{.variable} is true, then return.

6.  If `event`{.variable}\'s [focus reset
    behavior](#concept-navigateevent-focusreset){#navigate-event-scroll-focus:concept-navigateevent-focusreset}
    is
    \"[`manual`{#navigate-event-scroll-focus:dom-navigationfocusreset-manual}](#dom-navigationfocusreset-manual)\",
    then return.

    If it was left as null, then we treat that as
    \"[`after-transition`{#navigate-event-scroll-focus:dom-navigationfocusreset-after-transition}](#dom-navigationfocusreset-after-transition)\",
    and continue onward.

7.  Let `document`{.variable} be `event`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#navigate-event-scroll-focus:concept-relevant-global-2}\'s
    [associated
    `Document`](#concept-document-window){#navigate-event-scroll-focus:concept-document-window}.

8.  Let `focusTarget`{.variable} be the [autofocus
    delegate](interaction.html#autofocus-delegate){#navigate-event-scroll-focus:autofocus-delegate}
    for `document`{.variable}.

9.  If `focusTarget`{.variable} is null, then set
    `focusTarget`{.variable} to `document`{.variable}\'s [body
    element](dom.html#the-body-element-2){#navigate-event-scroll-focus:the-body-element-2}.

10. If `focusTarget`{.variable} is null, then set
    `focusTarget`{.variable} to `document`{.variable}\'s [document
    element](https://dom.spec.whatwg.org/#document-element){#navigate-event-scroll-focus:document-element
    x-internal="document-element"}.

11. Run the [focusing
    steps](interaction.html#focusing-steps){#navigate-event-scroll-focus:focusing-steps}
    for `focusTarget`{.variable}, with `document`{.variable}\'s
    [viewport](https://drafts.csswg.org/css2/#viewport){#navigate-event-scroll-focus:viewport
    x-internal="viewport"} as the `fallback target`{.variable}.

12. Move the [sequential focus navigation starting
    point](interaction.html#sequential-focus-navigation-starting-point){#navigate-event-scroll-focus:sequential-focus-navigation-starting-point}
    to `focusTarget`{.variable}.

To [potentially process scroll
behavior]{#potentially-process-scroll-behavior .dfn} given a
[`NavigateEvent`{#navigate-event-scroll-focus:navigateevent-3}](#navigateevent)
`event`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-scroll-focus:assert-3
    x-internal="assert"}: `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state-5}
    is \"`committed`\" or \"`scrolled`\".

2.  If `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state-6}
    is \"`scrolled`\", then return.

3.  If `event`{.variable}\'s [scroll
    behavior](#concept-navigateevent-scroll){#navigate-event-scroll-focus:concept-navigateevent-scroll}
    is
    \"[`manual`{#navigate-event-scroll-focus:dom-navigationscrollbehavior-manual}](#dom-navigationscrollbehavior-manual)\",
    then return.

    If it was left as null, then we treat that as
    \"[`after-transition`{#navigate-event-scroll-focus:dom-navigationscrollbehavior-after-transition}](#dom-navigationscrollbehavior-after-transition)\",
    and continue onward.

4.  [Process scroll
    behavior](#process-scroll-behavior){#navigate-event-scroll-focus:process-scroll-behavior}
    given `event`{.variable}.

To [process scroll behavior]{#process-scroll-behavior .dfn} given a
[`NavigateEvent`{#navigate-event-scroll-focus:navigateevent-4}](#navigateevent)
`event`{.variable}:

1.  [Assert](https://infra.spec.whatwg.org/#assert){#navigate-event-scroll-focus:assert-4
    x-internal="assert"}: `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state-7}
    is \"`committed`\".

2.  Set `event`{.variable}\'s [interception
    state](#concept-navigateevent-interception-state){#navigate-event-scroll-focus:concept-navigateevent-interception-state-8}
    to \"`scrolled`\".

3.  If `event`{.variable}\'s
    [`navigationType`{#navigate-event-scroll-focus:dom-navigateevent-navigationtype}](#dom-navigateevent-navigationtype)
    was initialized to
    \"[`traverse`{#navigate-event-scroll-focus:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\"
    or
    \"[`reload`{#navigate-event-scroll-focus:dom-navigationtype-reload}](#dom-navigationtype-reload)\",
    then [restore scroll position
    data](browsing-the-web.html#restore-scroll-position-data){#navigate-event-scroll-focus:restore-scroll-position-data}
    given `event`{.variable}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#navigate-event-scroll-focus:concept-relevant-global-3}\'s
    [navigable](#window-navigable){#navigate-event-scroll-focus:window-navigable}\'s
    [active session history
    entry](document-sequences.html#nav-active-history-entry){#navigate-event-scroll-focus:nav-active-history-entry}.

4.  Otherwise:

    1.  Let `document`{.variable} be `event`{.variable}\'s [relevant
        global
        object](webappapis.html#concept-relevant-global){#navigate-event-scroll-focus:concept-relevant-global-4}\'s
        [associated
        `Document`](#concept-document-window){#navigate-event-scroll-focus:concept-document-window-2}.

    2.  If `document`{.variable}\'s [indicated
        part](browsing-the-web.html#the-indicated-part-of-the-document){#navigate-event-scroll-focus:the-indicated-part-of-the-document}
        is null, then [scroll to the beginning of the
        document](https://drafts.csswg.org/cssom-view/#scroll-to-the-beginning-of-the-document){#navigate-event-scroll-focus:scroll-to-the-beginning-of-the-document
        x-internal="scroll-to-the-beginning-of-the-document"} given
        `document`{.variable}.
        [\[CSSOMVIEW\]](references.html#refsCSSOMVIEW)

    3.  Otherwise, [scroll to the
        fragment](browsing-the-web.html#scroll-to-the-fragment-identifier){#navigate-event-scroll-focus:scroll-to-the-fragment-identifier}
        given `document`{.variable}.

#### [7.2.7]{.secno} Event interfaces[](#nav-traversal-event-interfaces){.self-link} {#nav-traversal-event-interfaces}

The
[`NavigateEvent`{#nav-traversal-event-interfaces:navigateevent}](#navigateevent)
interface has [its own dedicated section](#the-navigate-event), due to
its complexity.

##### [7.2.7.1]{.secno} The [`NavigationCurrentEntryChangeEvent`{#the-navigationcurrententrychangeevent-interface:navigationcurrententrychangeevent}](#navigationcurrententrychangeevent) interface[](#the-navigationcurrententrychangeevent-interface){.self-link}

``` idl
[Exposed=Window]
interface NavigationCurrentEntryChangeEvent : Event {
  constructor(DOMString type, NavigationCurrentEntryChangeEventInit eventInitDict);

  readonly attribute NavigationType? navigationType;
  readonly attribute NavigationHistoryEntry from;
};

dictionary NavigationCurrentEntryChangeEventInit : EventInit {
  NavigationType? navigationType = null;
  required NavigationHistoryEntry from;
};
```

`event`{.variable}`.`[`navigationType`](#dom-navigationcurrententrychangeevent-navigationtype){#dom-navigationcurrententrychangeevent-navigationtype-dev}
:   Returns the type of navigation which caused the current entry to
    change, or null if the change is due to
    [`navigation.updateCurrentEntry()`{#the-navigationcurrententrychangeevent-interface:dom-navigation-updatecurrententry}](#dom-navigation-updatecurrententry).

`event`{.variable}`.`[`from`](#dom-navigationcurrententrychangeevent-from){#dom-navigationcurrententrychangeevent-from-dev}

:   Returns the previous value of
    [`navigation.currentEntry`{#the-navigationcurrententrychangeevent-interface:dom-navigation-currententry}](#dom-navigation-currententry),
    before the current entry changed.

    If
    [`navigationType`{#the-navigationcurrententrychangeevent-interface:dom-navigationcurrententrychangeevent-navigationtype-2}](#dom-navigationcurrententrychangeevent-navigationtype)
    is null or
    \"[`reload`{#the-navigationcurrententrychangeevent-interface:dom-navigationtype-reload}](#dom-navigationtype-reload)\",
    then this value will be the same as
    [`navigation.currentEntry`{#the-navigationcurrententrychangeevent-interface:dom-navigation-currententry-2}](#dom-navigation-currententry).
    In that case, the event signifies that the contents of the entry
    changed, even if we did not move to a new entry or replace the
    current one.

The
[`navigationType`]{#dom-navigationcurrententrychangeevent-navigationtype
.dfn dfn-for="NavigationCurrentEntryChangeEvent" dfn-type="attribute"}
and [`from`]{#dom-navigationcurrententrychangeevent-from .dfn
dfn-for="NavigationCurrentEntryChangeEvent" dfn-type="attribute"}
attributes must return the values they were initialized to.

##### [7.2.7.2]{.secno} The [`PopStateEvent`{#the-popstateevent-interface:popstateevent}](#popstateevent) interface[](#the-popstateevent-interface){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[PopStateEvent/PopStateEvent](https://developer.mozilla.org/en-US/docs/Web/API/PopStateEvent/PopStateEvent "The PopStateEvent() constructor creates a new PopStateEvent object.")

Support in all current engines.

::: support
[Firefox11+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome16+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[PopStateEvent](https://developer.mozilla.org/en-US/docs/Web/API/PopStateEvent "PopStateEvent is an interface for the popstate event.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome4+]{.chrome
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
:::::::

``` idl
[Exposed=Window]
interface PopStateEvent : Event {
  constructor(DOMString type, optional PopStateEventInit eventInitDict = {});

  readonly attribute any state;
  readonly attribute boolean hasUAVisualTransition;
};

dictionary PopStateEventInit : EventInit {
  any state = null;
  boolean hasUAVisualTransition = false;
};
```

`event`{.variable}`.`[`state`](#dom-popstateevent-state){#dom-popstateevent-state-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[PopStateEvent/state](https://developer.mozilla.org/en-US/docs/Web/API/PopStateEvent/state "The state read-only property of the PopStateEvent interface represents the state stored when the event was created.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome4+]{.chrome
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

Returns a copy of the information that was provided to
[`pushState()`{#the-popstateevent-interface:dom-history-pushstate}](#dom-history-pushstate)
or
[`replaceState()`{#the-popstateevent-interface:dom-history-replacestate}](#dom-history-replacestate).

`event`{.variable}`.`[`hasUAVisualTransition`](#dom-popstateevent-hasuavisualtransition){#dom-popstateevent-hasuavisualtransition-dev}

Returns true if the user agent performed a visual transition for this
navigation before dispatching this event. If true, the best user
experience will be given if the author synchronously updates the DOM to
the post-navigation state.

The [`state`]{#dom-popstateevent-state .dfn dfn-for="PopStateEvent"
dfn-type="attribute"} attribute must return the value it was initialized
to. It represents the context information for the event, or null, if the
state represented is the initial state of the
[`Document`{#the-popstateevent-interface:document}](dom.html#document).

The [` hasUAVisualTransition`]{#dom-popstateevent-hasuavisualtransition
.dfn dfn-for="PopStateEvent" dfn-type="attribute"} attribute must return
the value it was initialized to.

##### [7.2.7.3]{.secno} The [`HashChangeEvent`{#the-hashchangeevent-interface:hashchangeevent}](#hashchangeevent) interface[](#the-hashchangeevent-interface){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[HashChangeEvent/HashChangeEvent](https://developer.mozilla.org/en-US/docs/Web/API/HashChangeEvent/HashChangeEvent "The HashChangeEvent() constructor creates a new HashChangeEvent object, that is used by the hashchange event fired at the window object when the fragment of the URL changes.")

Support in all current engines.

::: support
[Firefox11+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome16+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[HashChangeEvent](https://developer.mozilla.org/en-US/docs/Web/API/HashChangeEvent "The HashChangeEvent interface represents events that fire when the fragment identifier of the URL has changed.")

Support in all current engines.

::: support
[Firefox3.6+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer8+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::::

``` idl
[Exposed=Window]
interface HashChangeEvent : Event {
  constructor(DOMString type, optional HashChangeEventInit eventInitDict = {});

  readonly attribute USVString oldURL;
  readonly attribute USVString newURL;
};

dictionary HashChangeEventInit : EventInit {
  USVString oldURL = "";
  USVString newURL = "";
};
```

`event`{.variable}`.`[`oldURL`](#dom-hashchangeevent-oldurl){#dom-hashchangeevent-oldurl-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HashChangeEvent/oldURL](https://developer.mozilla.org/en-US/docs/Web/API/HashChangeEvent/oldURL "The oldURL read-only property of the HashChangeEvent interface returns the previous URL from which the window was navigated.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[URL](https://url.spec.whatwg.org/#concept-url){#the-hashchangeevent-interface:url
x-internal="url"} of the [session history
entry](browsing-the-web.html#session-history-entry){#the-hashchangeevent-interface:session-history-entry}
that was previously current.

`event`{.variable}`.`[`newURL`](#dom-hashchangeevent-newurl){#dom-hashchangeevent-newurl-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[HashChangeEvent/newURL](https://developer.mozilla.org/en-US/docs/Web/API/HashChangeEvent/newURL "The newURL read-only property of the HashChangeEvent interface returns the new URL to which the window is navigating.")

Support in all current engines.

::: support
[Firefox6+]{.firefox .yes}[Safari5.1+]{.safari .yes}[Chrome8+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

Returns the
[URL](https://url.spec.whatwg.org/#concept-url){#the-hashchangeevent-interface:url-2
x-internal="url"} of the [session history
entry](browsing-the-web.html#session-history-entry){#the-hashchangeevent-interface:session-history-entry-2}
that is now current.

The [`oldURL`]{#dom-hashchangeevent-oldurl .dfn
dfn-for="HashChangeEvent" dfn-type="attribute"} attribute must return
the value it was initialized to. It represents context information for
the event, specifically the URL of the [session history
entry](browsing-the-web.html#session-history-entry){#the-hashchangeevent-interface:session-history-entry-3}
that was traversed from.

The [`newURL`]{#dom-hashchangeevent-newurl .dfn
dfn-for="HashChangeEvent" dfn-type="attribute"} attribute must return
the value it was initialized to. It represents context information for
the event, specifically the URL of the [session history
entry](browsing-the-web.html#session-history-entry){#the-hashchangeevent-interface:session-history-entry-4}
that was traversed to.

##### [7.2.7.4]{.secno} The [`PageSwapEvent`{#the-pageswapevent-interface:pageswapevent}](#pageswapevent) interface[](#the-pageswapevent-interface){.self-link}

``` idl
[Exposed=Window]
interface PageSwapEvent : Event {
  constructor(DOMString type, optional PageSwapEventInit eventInitDict = {});
  readonly attribute NavigationActivation? activation;
  readonly attribute ViewTransition? viewTransition;
};

dictionary PageSwapEventInit : EventInit {
  NavigationActivation? activation = null;
  ViewTransition? viewTransition = null;
};
```

`event`{.variable}`.`[`activation`](#dom-pageswapevent-activation){#dom-pageswapevent-activation-dev}
:   A
    [`NavigationActivation`{#the-pageswapevent-interface:navigationactivation-3}](#navigationactivation)
    object representing the destination and type of the cross-document
    navigation. This would be null for cross-origin navigations.

`event`{.variable}`.`[`activation`](#dom-pageswapevent-activation){#the-pageswapevent-interface:dom-pageswapevent-activation-2}`.`[`entry`](#dom-navigationactivation-entry){#the-pageswapevent-interface:dom-navigationactivation-entry}
:   A
    [`NavigationHistoryEntry`{#the-pageswapevent-interface:navigationhistoryentry}](#navigationhistoryentry),
    representing the
    [`Document`{#the-pageswapevent-interface:document}](dom.html#document)
    that is about to become active.

`event`{.variable}`.`[`activation`](#dom-pageswapevent-activation){#the-pageswapevent-interface:dom-pageswapevent-activation-3}`.`[`from`](#dom-navigationactivation-from){#the-pageswapevent-interface:dom-navigationactivation-from}
:   A
    [`NavigationHistoryEntry`{#the-pageswapevent-interface:navigationhistoryentry-2}](#navigationhistoryentry),
    equivalent to the value of the
    [`navigation.currentEntry`{#the-pageswapevent-interface:dom-navigation-currententry}](#dom-navigation-currententry)
    property at the moment the event is fired.

`event`{.variable}`.`[`activation`](#dom-pageswapevent-activation){#the-pageswapevent-interface:dom-pageswapevent-activation-4}`.`[`navigationType`](#dom-navigationactivation-navigationtype){#the-pageswapevent-interface:dom-navigationactivation-navigationtype}
:   One of
    \"[`push`{#the-pageswapevent-interface:dom-navigationtype-push}](#dom-navigationtype-push)\",
    \"[`replace`{#the-pageswapevent-interface:dom-navigationtype-replace}](#dom-navigationtype-replace)\",
    \"[`reload`{#the-pageswapevent-interface:dom-navigationtype-reload}](#dom-navigationtype-reload)\",
    or
    \"[`traverse`{#the-pageswapevent-interface:dom-navigationtype-traverse}](#dom-navigationtype-traverse)\",
    indicating what type of navigation that is about to result in a page
    swap.

`event`{.variable}`.`[`viewTransition`](#dom-pageswapevent-viewtransition){#dom-pageswapevent-viewtransition-dev}

:   Returns the
    [`ViewTransition`{#the-pageswapevent-interface:viewtransition-3}](https://drafts.csswg.org/css-view-transitions/#viewtransition){x-internal="viewtransition"}
    object that represents an outbound cross-document view transition,
    if such transition is active when the event is fired. Otherwise,
    returns null.

The [`activation`]{#dom-pageswapevent-activation .dfn
dfn-for="PageSwapEvent" dfn-type="attribute"} and
[`viewTransition`]{#dom-pageswapevent-viewtransition .dfn
dfn-for="PageSwapEvent" dfn-type="attribute"} attributes must return the
values they were initialized to.

##### [7.2.7.5]{.secno} The [`PageRevealEvent`{#the-pagerevealevent-interface:pagerevealevent}](#pagerevealevent) interface[](#the-pagerevealevent-interface){.self-link}

``` idl
[Exposed=Window]
interface PageRevealEvent : Event {
  constructor(DOMString type, optional PageRevealEventInit eventInitDict = {});
  readonly attribute ViewTransition? viewTransition;
};

dictionary PageRevealEventInit : EventInit {
  ViewTransition? viewTransition = null;
};
```

`event`{.variable}`.`[`viewTransition`](#dom-pagerevealevent-viewtransition){#dom-pagerevealevent-viewtransition-dev}

:   Returns the
    [`ViewTransition`{#the-pagerevealevent-interface:viewtransition-3}](https://drafts.csswg.org/css-view-transitions/#viewtransition){x-internal="viewtransition"}
    object that represents an inbound cross-document view transition, if
    such transition is active when the event is fired. Otherwise,
    returns null.

The [`viewTransition`]{#dom-pagerevealevent-viewtransition .dfn
dfn-for="PageRevealEvent" dfn-type="attribute"} attribute must return
the value it was initialized to.

##### [7.2.7.6]{.secno} The [`PageTransitionEvent`{#the-pagetransitionevent-interface:pagetransitionevent}](#pagetransitionevent) interface[](#the-pagetransitionevent-interface){.self-link}

::::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[PageTransitionEvent/PageTransitionEvent](https://developer.mozilla.org/en-US/docs/Web/API/PageTransitionEvent/PageTransitionEvent "The PageTransitionEvent() constructor creates a new PageTransitionEvent object, that is used by the pageshow or pagehide events, fired at the window object when a page is loaded or unloaded.")

Support in all current engines.

::: support
[Firefox11+]{.firefox .yes}[Safari6+]{.safari .yes}[Chrome16+]{.chrome
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
[PageTransitionEvent](https://developer.mozilla.org/en-US/docs/Web/API/PageTransitionEvent "The PageTransitionEvent event object is available inside handler functions for the pageshow and pagehide events, fired when a document is being loaded or unloaded.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

``` idl
[Exposed=Window]
interface PageTransitionEvent : Event {
  constructor(DOMString type, optional PageTransitionEventInit eventInitDict = {});

  readonly attribute boolean persisted;
};

dictionary PageTransitionEventInit : EventInit {
  boolean persisted = false;
};
```

`event`{.variable}`.`[`persisted`](#dom-pagetransitionevent-persisted){#dom-pagetransitionevent-persisted-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[PageTransitionEvent/persisted](https://developer.mozilla.org/en-US/docs/Web/API/PageTransitionEvent/persisted "The persisted read-only property indicates if a webpage is loading from a cache.")

Support in all current engines.

::: support
[Firefox11+]{.firefox .yes}[Safari5+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer11]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS4+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

For the
[`pageshow`{#the-pagetransitionevent-interface:event-pageshow}](indices.html#event-pageshow)
event, returns false if the page is newly being loaded (and the
[`load`{#the-pagetransitionevent-interface:event-load}](indices.html#event-load)
event will fire). Otherwise, returns true.

For the
[`pagehide`{#the-pagetransitionevent-interface:event-pagehide}](indices.html#event-pagehide)
event, returns false if the page is going away for the last time.
Otherwise, returns true, meaning that the page might be reused if the
user navigates back to this page (if the
[`Document`{#the-pagetransitionevent-interface:document}](dom.html#document)\'s
*[salvageable](document-lifecycle.html#concept-document-salvageable)*
state stays true).

Things that can cause the page to be unsalvageable include:

- The user agent decided to not keep the
  [`Document`{#the-pagetransitionevent-interface:document-2}](dom.html#document)
  alive in a [session history
  entry](browsing-the-web.html#session-history-entry){#the-pagetransitionevent-interface:session-history-entry}
  after
  [unload](document-lifecycle.html#unload-a-document){#the-pagetransitionevent-interface:unload-a-document}

- Having
  [`iframe`{#the-pagetransitionevent-interface:the-iframe-element}](iframe-embed-object.html#the-iframe-element)s
  that are not
  *[salvageable](document-lifecycle.html#concept-document-salvageable)*

- Active
  [`WebSocket`{#the-pagetransitionevent-interface:websocket-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
  objects

- [Aborting a
  `Document`](document-lifecycle.html#abort-a-document){#the-pagetransitionevent-interface:abort-a-document}

The [`persisted`]{#dom-pagetransitionevent-persisted .dfn
dfn-for="PageTransitionEvent" dfn-type="attribute"} attribute must
return the value it was initialized to. It represents the context
information for the event.

To [fire a page transition event]{#fire-a-page-transition-event .dfn}
named `eventName`{.variable} at a
[`Window`{#the-pagetransitionevent-interface:window}](#window)
`window`{.variable} with a boolean `persisted`{.variable}, [fire an
event](https://dom.spec.whatwg.org/#concept-event-fire){#the-pagetransitionevent-interface:concept-event-fire
x-internal="concept-event-fire"} named `eventName`{.variable} at
`window`{.variable}, using
[`PageTransitionEvent`{#the-pagetransitionevent-interface:pagetransitionevent-2}](#pagetransitionevent),
with the
[`persisted`{#the-pagetransitionevent-interface:dom-pagetransitionevent-persisted-2}](#dom-pagetransitionevent-persisted)
attribute initialized to `persisted`{.variable}, the
[`cancelable`{#the-pagetransitionevent-interface:dom-event-cancelable}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
attribute initialized to true, the
[`bubbles`{#the-pagetransitionevent-interface:dom-event-bubbles}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
attribute initialized to true, and *legacy target override flag* set.

The values for
[`cancelable`{#the-pagetransitionevent-interface:dom-event-cancelable-2}](https://dom.spec.whatwg.org/#dom-event-cancelable){x-internal="dom-event-cancelable"}
and
[`bubbles`{#the-pagetransitionevent-interface:dom-event-bubbles-2}](https://dom.spec.whatwg.org/#dom-event-bubbles){x-internal="dom-event-bubbles"}
don\'t make any sense, since canceling the event does nothing and it\'s
not possible to bubble past the
[`Window`{#the-pagetransitionevent-interface:window-2}](#window) object.
They are set to true for historical reasons.

##### [7.2.7.7]{.secno} The [`BeforeUnloadEvent`{#the-beforeunloadevent-interface:beforeunloadevent}](#beforeunloadevent) interface[](#the-beforeunloadevent-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[BeforeUnloadEvent](https://developer.mozilla.org/en-US/docs/Web/API/BeforeUnloadEvent "BeforeUnloadEvent is an interface for the beforeunload event.")

Support in all current engines.

::: support
[Firefox1.5+]{.firefox .yes}[Safari7+]{.safari .yes}[Chrome30+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet Explorer4+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet3.0+]{.samsunginternet_android .yes}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

``` idl
[Exposed=Window]
interface BeforeUnloadEvent : Event {
  attribute DOMString returnValue;
};
```

There are no
[`BeforeUnloadEvent`{#the-beforeunloadevent-interface:beforeunloadevent-2}](#beforeunloadevent)-specific
initialization methods.

The
[`BeforeUnloadEvent`{#the-beforeunloadevent-interface:beforeunloadevent-3}](#beforeunloadevent)
interface is a legacy interface which allows [checking if unloading is
canceled](browsing-the-web.html#checking-if-unloading-is-canceled){#the-beforeunloadevent-interface:checking-if-unloading-is-canceled}
to be controlled not only by canceling the event, but by setting the
[`returnValue`{#the-beforeunloadevent-interface:dom-beforeunloadevent-returnvalue-2}](#dom-beforeunloadevent-returnvalue)
attribute to a value besides the empty string. Authors should use the
[`preventDefault()`{#the-beforeunloadevent-interface:dom-event-preventdefault}](https://dom.spec.whatwg.org/#dom-event-preventdefault){x-internal="dom-event-preventdefault"}
method, or other means of canceling events, instead of using
[`returnValue`{#the-beforeunloadevent-interface:dom-beforeunloadevent-returnvalue-3}](#dom-beforeunloadevent-returnvalue).

The [`returnValue`]{#dom-beforeunloadevent-returnvalue .dfn
dfn-for="BeforeUnloadEvent" dfn-type="attribute"} attribute controls the
process of [checking if unloading is
canceled](browsing-the-web.html#checking-if-unloading-is-canceled){#the-beforeunloadevent-interface:checking-if-unloading-is-canceled-2}.
When the event is created, the attribute must be set to the empty
string. On getting, it must return the last value it was set to. On
setting, the attribute must be set to the new value.

This attribute is a `DOMString` only for historical reasons. Any value
besides the empty string will be treated as a request to ask the user
for confirmation.

#### [7.2.8]{.secno} The [`NotRestoredReasons`{#the-notrestoredreasons-interface:notrestoredreasons}](#notrestoredreasons) interface[](#the-notrestoredreasons-interface){.self-link}

``` idl
[Exposed=Window]
interface NotRestoredReasonDetails {
  readonly attribute DOMString reason;
  [Default] object toJSON();
};

[Exposed=Window]
interface NotRestoredReasons {
  readonly attribute USVString? src;
  readonly attribute DOMString? id;
  readonly attribute DOMString? name;
  readonly attribute USVString? url;
  readonly attribute FrozenArray<NotRestoredReasonDetails>? reasons;
  readonly attribute FrozenArray<NotRestoredReasons>? children;
  [Default] object toJSON();
};
```

`notRestoredReasonDetails`{.variable}`.`[`reason`](#dom-not-restored-reason-details-reason){#dom-not-restored-reason-details-reason-dev}
:   Returns a string that explains the reason that prevented the
    document from [being served from back/forward
    cache](browsing-the-web.html#note-bfcache). See the definition of
    [bfcache blocking
    details](dom.html#concept-document-bfcache-blocking-details){#the-notrestoredreasons-interface:concept-document-bfcache-blocking-details}
    for the possible string values.

`notRestoredReasons`{.variable}`.`[`src`](#dom-not-restored-reasons-src){#dom-not-restored-reasons-src-dev}
:   Returns the
    [`src`{#the-notrestoredreasons-interface:attr-iframe-src}](iframe-embed-object.html#attr-iframe-src)
    attribute of the document\'s [node
    navigable](document-sequences.html#node-navigable){#the-notrestoredreasons-interface:node-navigable}\'s
    [container](document-sequences.html#nav-container){#the-notrestoredreasons-interface:nav-container}
    if it is an
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
    element. This can be null if not set or if it is not an
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-2}](iframe-embed-object.html#the-iframe-element)
    element.

`notRestoredReasons`{.variable}`.`[`id`](#dom-not-restored-reasons-id){#dom-not-restored-reasons-id-dev}
:   Returns the
    [`id`{#the-notrestoredreasons-interface:the-id-attribute}](dom.html#the-id-attribute)
    attribute of the document\'s [node
    navigable](document-sequences.html#node-navigable){#the-notrestoredreasons-interface:node-navigable-2}\'s
    [container](document-sequences.html#nav-container){#the-notrestoredreasons-interface:nav-container-2}
    if it is an
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-3}](iframe-embed-object.html#the-iframe-element)
    element. This can be null if not set or if it is not an
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-4}](iframe-embed-object.html#the-iframe-element)
    element.

`notRestoredReasons`{.variable}`.`[`name`](#dom-not-restored-reasons-name){#dom-not-restored-reasons-name-dev}
:   Returns the
    [`name`{#the-notrestoredreasons-interface:attr-iframe-name}](iframe-embed-object.html#attr-iframe-name)
    attribute of the document\'s [node
    navigable](document-sequences.html#node-navigable){#the-notrestoredreasons-interface:node-navigable-3}\'s
    [container](document-sequences.html#nav-container){#the-notrestoredreasons-interface:nav-container-3}
    if it is an
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-5}](iframe-embed-object.html#the-iframe-element)
    element. This can be null if not set or if it is not an
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-6}](iframe-embed-object.html#the-iframe-element)
    element.

`notRestoredReasons`{.variable}`.`[`url`](#dom-not-restored-reasons-url){#dom-not-restored-reasons-url-dev}
:   Returns the document\'s
    [URL](https://dom.spec.whatwg.org/#concept-document-url){#the-notrestoredreasons-interface:the-document's-address
    x-internal="the-document's-address"}, or null if the document is in
    a cross-origin
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-7}](iframe-embed-object.html#the-iframe-element).
    This is reported in addition to
    [`src`{#the-notrestoredreasons-interface:dom-not-restored-reasons-src-2}](#dom-not-restored-reasons-src)
    because it is possible
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-8}](iframe-embed-object.html#the-iframe-element)
    navigated since the original
    [`src`{#the-notrestoredreasons-interface:attr-iframe-src-2}](iframe-embed-object.html#attr-iframe-src)
    was set.

`notRestoredReasons`{.variable}`.`[`reasons`](#dom-not-restored-reasons-reasons){#dom-not-restored-reasons-reasons-dev}
:   Returns an array of
    [`NotRestoredReasonDetails`{#the-notrestoredreasons-interface:notrestoredreasondetails-2}](#notrestoredreasondetails)
    for the document. This is null if the document is in a cross-origin
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-9}](iframe-embed-object.html#the-iframe-element).

`notRestoredReasons`{.variable}`.`[`children`](#dom-not-restored-reasons-children){#dom-not-restored-reasons-children-dev}

:   Returns an array of
    [`NotRestoredReasons`{#the-notrestoredreasons-interface:notrestoredreasons-3}](#notrestoredreasons)
    that are for the document's children. This is null if the document
    is in a cross-origin
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-10}](iframe-embed-object.html#the-iframe-element).

A
[`NotRestoredReasonDetails`{#the-notrestoredreasons-interface:notrestoredreasondetails-3}](#notrestoredreasondetails)
object has a [backing
struct]{#concept-not-restored-reason-details-backing-struct .dfn}, a
[not restored reason
details](#nrr-details-struct){#the-notrestoredreasons-interface:nrr-details-struct}
or null, initially null.

The [`reason`]{#dom-not-restored-reason-details-reason .dfn
dfn-for="NotRestoredReasonDetails" dfn-type="attribute"} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this
x-internal="this"}\'s [backing
struct](#concept-not-restored-reason-details-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reason-details-backing-struct}\'s
[reason](#nrr-details-reason){#the-notrestoredreasons-interface:nrr-details-reason}.

To [create a `NotRestoredReasonDetails`
object]{#create-a-notrestoredreasondetails-object .dfn} given a [not
restored reason
details](#nrr-details-struct){#the-notrestoredreasons-interface:nrr-details-struct-2}
`backingStruct`{.variable} and a
[realm](https://tc39.es/ecma262/#sec-code-realms){#the-notrestoredreasons-interface:realm
x-internal="realm"} `realm`{.variable}:

1.  Let `notRestoredReasonDetails`{.variable} be a new
    [`NotRestoredReasonDetails`{#the-notrestoredreasons-interface:notrestoredreasondetails-4}](#notrestoredreasondetails)
    object created in `realm`{.variable}.

2.  Set `notRestoredReasonDetails`{.variable}\'s [backing
    struct](#concept-not-restored-reason-details-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reason-details-backing-struct-2}
    to `backingStruct`{.variable}.

3.  Return `notRestoredReasonDetails`{.variable}.

A [not restored reason details]{#nrr-details-struct .dfn export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#the-notrestoredreasons-interface:struct
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#the-notrestoredreasons-interface:struct-item
x-internal="struct-item"}:

- [reason]{#nrr-details-reason .dfn}, a string, initially empty.

The
[reason](#nrr-details-reason){#the-notrestoredreasons-interface:nrr-details-reason-2}
is a string that represents the reason that prevented the page from
being restored from [back/forward
cache](browsing-the-web.html#note-bfcache). The string is one of the
following:

\"[`fetch`]{#blocking-fetch .dfn}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document},
    a fetch initiated by this
    [`Document`{#the-notrestoredreasons-interface:document}](dom.html#document)
    was still ongoing and was canceled, so the page was not in a state
    that could be stored in the [back/forward
    cache](browsing-the-web.html#note-bfcache).

\"[`navigation-failure`]{#blocking-navigation-failure .dfn}\"
:   The original navigation that created this
    [`Document`{#the-notrestoredreasons-interface:document-2}](dom.html#document)
    errored, so storing the resulting error document in the
    [back/forward cache](browsing-the-web.html#note-bfcache) was
    prevented.

\"[`parser-aborted`]{#blocking-parser-aborted .dfn}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-3}](dom.html#document)
    never finished its initial HTML parsing, so storing the unfinished
    document in the [back/forward
    cache](browsing-the-web.html#note-bfcache) was prevented.

\"[`websocket`]{#blocking-websocket .dfn}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-2},
    an open
    [`WebSocket`{#the-notrestoredreasons-interface:websocket-2}](https://websockets.spec.whatwg.org/#websocket){x-internal="websocket-2"}
    connection was shut down, so the page was not in a state that could
    be stored in the [back/forward
    cache](browsing-the-web.html#note-bfcache).
    [\[WEBSOCKETS\]](references.html#refsWEBSOCKETS)

\"[`lock`]{#blocking-weblock .dfn export=""}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-3},
    held
    [locks](https://w3c.github.io/web-locks/#lock-concept){#the-notrestoredreasons-interface:locks
    x-internal="locks"} and [lock
    requests](https://w3c.github.io/web-locks/#lock-request){#the-notrestoredreasons-interface:lock-requests
    x-internal="lock-requests"} were terminated, so the page was not in
    a state that could be stored in the [back/forward
    cache](browsing-the-web.html#note-bfcache).
    [\[WEBLOCKS\]](references.html#refsWEBLOCKS)

\"[`masked`]{#blocking-masked .dfn export=""}\"
:   This
    [`Document`{#the-notrestoredreasons-interface:document-4}](dom.html#document)
    has children that are in a cross-origin
    [`iframe`{#the-notrestoredreasons-interface:the-iframe-element-11}](iframe-embed-object.html#the-iframe-element),
    and they prevented [back/forward
    cache](browsing-the-web.html#note-bfcache); or this
    [`Document`{#the-notrestoredreasons-interface:document-5}](dom.html#document)
    could not be [back/forward
    cached](browsing-the-web.html#note-bfcache) for user agent-specific
    reasons, and the user agent has chosen not to use one of the more
    specific reasons from the list of [user-agent specific blocking
    reasons](#ua-specific-blocking-reasons){#the-notrestoredreasons-interface:ua-specific-blocking-reasons}.

In addition to the list above, a user agent might choose to expose a
reason that prevented the page from being restored from [back/forward
cache](note-bfcache) for [user-agent specific blocking
reasons]{#ua-specific-blocking-reasons .dfn}. These are one of the
following strings:

\"[`audio-capture`]{#blocking-audio-capture .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-6}](dom.html#document)
    requested audio capture permission by using Media Capture and
    Streams\'s
    [`getUserMedia()`{#the-notrestoredreasons-interface:getusermedia()}](https://w3c.github.io/mediacapture-main/getusermedia.html#dom-mediadevices-getusermedia){x-internal="getusermedia()"}
    with audio. [\[MEDIASTREAM\]](references.html#refsMEDIASTREAM)

\"[`background-work`]{#blocking-background-work .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-7}](dom.html#document)
    requested background work by calling
    [`SyncManager`{#the-notrestoredreasons-interface:syncmanager}](https://wicg.github.io/background-sync/spec/#syncmanager){x-internal="syncmanager"}\'s
    [`register()`{#the-notrestoredreasons-interface:dom-syncmanager-register}](https://wicg.github.io/background-sync/spec/#dom-syncmanager-register){x-internal="dom-syncmanager-register"}
    method,
    [`PeriodicSyncManager`{#the-notrestoredreasons-interface:periodicsyncmanager}](https://wicg.github.io/periodic-background-sync/#periodicsyncmanager){x-internal="periodicsyncmanager"}\'s
    [`register()`{#the-notrestoredreasons-interface:dom-periodicsyncmanager-register}](https://wicg.github.io/periodic-background-sync/#dom-periodicsyncmanager-register){x-internal="dom-periodicsyncmanager-register"}
    method, or
    [`BackgroundFetchManager`{#the-notrestoredreasons-interface:backgroundfetchmanager}](https://wicg.github.io/background-fetch/#backgroundfetchmanager){x-internal="backgroundfetchmanager"}\'s
    [`fetch()`{#the-notrestoredreasons-interface:dom-backgroundfetchmanager-fetch}](https://wicg.github.io/background-fetch/#dom-backgroundfetchmanager-fetch){x-internal="dom-backgroundfetchmanager-fetch"}
    method.

\"[`broadcastchannel-message`]{#blocking-broadcastchannel-message .dfn export=""}\"
:   While the page was stored in [back/forward
    cache](browsing-the-web.html#note-bfcache), a
    [`BroadcastChannel`{#the-notrestoredreasons-interface:broadcastchannel}](web-messaging.html#broadcastchannel)
    connection on the page received a message and
    [`message`{#the-notrestoredreasons-interface:event-message}](indices.html#event-message)
    event was fired.

\"[`idbversionchangeevent`]{#blocking-idb-event .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-8}](dom.html#document)
    had a pending
    [`IDBVersionChangeEvent`{#the-notrestoredreasons-interface:idbversionchangeevent}](https://w3c.github.io/IndexedDB/#idbversionchangeevent){x-internal="idbversionchangeevent"}
    while
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-4}.
    [\[INDEXEDDB\]](references.html#refsINDEXEDDB)

\"[`idledetector`]{#blocking-idledetector .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-9}](dom.html#document)
    had an active
    [`IdleDetector`{#the-notrestoredreasons-interface:idledetector}](https://wicg.github.io/idle-detection/#idledetector){x-internal="idledetector"}
    while
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-5}.

\"[`keyboardlock`]{#blocking-keyboard-lock .dfn export=""}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-6},
    keyboard lock was still active because
    [`Keyboard`{#the-notrestoredreasons-interface:keyboard}](https://wicg.github.io/keyboard-lock/#keyboard){x-internal="keyboard"}\'s
    [`lock()`{#the-notrestoredreasons-interface:dom-keyboard-lock}](https://wicg.github.io/keyboard-lock/#dom-keyboard-lock){x-internal="dom-keyboard-lock"}
    method was called.

\"[`mediastream`]{#blocking-mediastream .dfn export=""}\"
:   A
    [`MediaStreamTrack`{#the-notrestoredreasons-interface:mediastreamtrack}](https://w3c.github.io/mediacapture-main/getusermedia.html#mediastreamtrack){x-internal="mediastreamtrack"}
    was in the [live
    state](https://w3c.github.io/mediacapture-main/getusermedia.html#idl-def-MediaStreamTrackState.live){#the-notrestoredreasons-interface:live-state
    x-internal="live-state"} upon
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-7}.
    [\[MEDIASTREAM\]](references.html#refsMEDIASTREAM)

\"[`midi`]{#blocking-midi .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-10}](dom.html#document)
    requested a MIDI permission by calling
    [`navigator.requestMIDIAccess()`{#the-notrestoredreasons-interface:dom-navigator-requestmidiaccess}](https://webaudio.github.io/web-midi-api/#dom-navigator-requestmidiaccess){x-internal="dom-navigator-requestmidiaccess"}.

\"[`modals`]{#blocking-modals .dfn export=""}\"
:   [User prompts](timers-and-user-prompts.html#user-prompts) were shown
    while
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-8}.

\"[`navigating`]{#blocking-navigating .dfn export=""}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-9},
    loading was still ongoing, and so the
    [`Document`{#the-notrestoredreasons-interface:document-11}](dom.html#document)
    was not in a state that could be stored in [back/forward
    cache](browsing-the-web.html#note-bfcache).

\"[`navigation-canceled`]{#blocking-navigation-canceled .dfn export=""}\"
:   The navigation request was canceled by calling
    [`window.stop()`{#the-notrestoredreasons-interface:dom-window-stop}](#dom-window-stop)
    and the page was not in a state to be stored in [back/forward
    cache](note-bfcache).

\"[`non-trivial-browsing-context-group`]{#blocking-non-trivial-bcg .dfn export=""}\"
:   The [browsing context
    group](document-sequences.html#browsing-context-group){#the-notrestoredreasons-interface:browsing-context-group}
    of this
    [`Document`{#the-notrestoredreasons-interface:document-12}](dom.html#document)
    had more than one [top-level browsing
    context](document-sequences.html#top-level-browsing-context){#the-notrestoredreasons-interface:top-level-browsing-context}.

\"[`otpcredential`]{#blocking-otp .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-13}](dom.html#document)
    created an
    [`OTPCredential`{#the-notrestoredreasons-interface:otpcredential}](https://wicg.github.io/web-otp/#otpcredential){x-internal="otpcredential"}.

\"[`outstanding-network-request`]{#blocking-outstanding-network-request .dfn export=""}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-10},
    the
    [`Document`{#the-notrestoredreasons-interface:document-14}](dom.html#document)
    had outstanding network requests and was not in a state that could
    be stored in [back/forward
    cache](browsing-the-web.html#note-bfcache).

\"[`paymentrequest`]{#blocking-payment .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-15}](dom.html#document)
    had an active
    [`PaymentRequest`{#the-notrestoredreasons-interface:paymentrequest}](https://w3c.github.io/payment-request/#dom-paymentrequest){x-internal="paymentrequest"}
    while
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-11}.
    [\[PAYMENTREQUEST\]](references.html#refsPAYMENTREQUEST)

\"[`pictureinpicturewindow`]{#blocking-picture-in-picture .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-16}](dom.html#document)
    had an active
    [`PictureInPictureWindow`{#the-notrestoredreasons-interface:pictureinpicturewindow}](https://w3c.github.io/picture-in-picture/#pictureinpicturewindow){x-internal="pictureinpicturewindow"}
    while
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-12}.
    [\[PICTUREINPICTURE\]](references.html#refsPICTUREINPICTURE)

\"[`plugins`]{#blocking-plugins .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-17}](dom.html#document)
    contained
    [plugins](infrastructure.html#plugin){#the-notrestoredreasons-interface:plugin}.

\"[`request-method-not-get`]{#blocking-method-not-get .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-18}](dom.html#document)
    was created from an HTTP request whose
    [method](https://fetch.spec.whatwg.org/#concept-request-method){#the-notrestoredreasons-interface:concept-request-method
    x-internal="concept-request-method"} was not \``GET`\`.
    [\[FETCH\]](references.html#refsFETCH)

\"[`response-auth-required`]{#blocking-auth-required .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-19}](dom.html#document)
    was created from an HTTP response that required HTTP authentication.

\"[`response-cache-control-no-store`]{#blocking-ccns .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-20}](dom.html#document)
    was created from an HTTP response whose
    \`[`Cache-Control`{#the-notrestoredreasons-interface:http-cache-control}](https://httpwg.org/specs/rfc7234.html#header.cache-control){x-internal="http-cache-control"}\`
    header included the \"`no-store`\" token.
    [\[HTTP\]](references.html#refsHTTP)

\"[`response-cache-control-no-cache`]{#blocking-ccnc .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-21}](dom.html#document)
    was created from an HTTP response whose
    \`[`Cache-Control`{#the-notrestoredreasons-interface:http-cache-control-2}](https://httpwg.org/specs/rfc7234.html#header.cache-control){x-internal="http-cache-control"}\`
    header included the \"`no-cache`\" token.
    [\[HTTP\]](references.html#refsHTTP)

\"[`response-keep-alive`]{#blocking-keepalive .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-22}](dom.html#document)
    was created from an HTTP response that contained a \``Keep-Alive`\`
    header.

\"[`response-scheme-not-http-or-https`]{#blocking-not-http-https .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-23}](dom.html#document)
    was created from a response whose
    [URL](https://fetch.spec.whatwg.org/#concept-response-url){#the-notrestoredreasons-interface:concept-response-url
    x-internal="concept-response-url"}\'s
    [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#the-notrestoredreasons-interface:concept-url-scheme
    x-internal="concept-url-scheme"} was not an [HTTP(S)
    scheme](https://fetch.spec.whatwg.org/#http-scheme){#the-notrestoredreasons-interface:http(s)-scheme
    x-internal="http(s)-scheme"}. [\[FETCH\]](references.html#refsFETCH)

\"[`response-status-not-ok`]{#blocking-response-not-ok .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-24}](dom.html#document)
    was created from an HTTP response whose
    [status](https://fetch.spec.whatwg.org/#concept-response-status){#the-notrestoredreasons-interface:concept-response-status
    x-internal="concept-response-status"} was not an [ok
    status](https://fetch.spec.whatwg.org/#ok-status){#the-notrestoredreasons-interface:ok-status
    x-internal="ok-status"}. [\[FETCH\]](references.html#refsFETCH)

\"[`rtc`]{#blocking-rtc .dfn export=""}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-13},
    a
    [`RTCPeerConnection`{#the-notrestoredreasons-interface:rtcpeerconnection}](https://w3c.github.io/webrtc-pc/#dom-rtcpeerconnection){x-internal="rtcpeerconnection"}
    or
    [`RTCDataChannel`{#the-notrestoredreasons-interface:rtcdatachannel}](https://w3c.github.io/webrtc-pc/#dom-rtcdatachannel){x-internal="rtcdatachannel"}
    was shut down, so the page was not in a state that could be stored
    in the [back/forward cache](browsing-the-web.html#note-bfcache).
    [\[WEBRTC\]](references.html#refsWEBRTC)

\"[`sensors`]{#blocking-sensors .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-25}](dom.html#document)
    [requested sensor
    access](https://w3c.github.io/sensors/#request-sensor-access){#the-notrestoredreasons-interface:request-sensor-access
    x-internal="request-sensor-access"}.

\"[`serviceworker-added`]{#sw-added .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-26}](dom.html#document)\'s
    [service worker
    client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client){#the-notrestoredreasons-interface:serviceworkercontainer-service-worker-client
    x-internal="serviceworkercontainer-service-worker-client"} started
    to be
    [controlled](https://w3c.github.io/ServiceWorker/#dfn-control){#the-notrestoredreasons-interface:dfn-control
    x-internal="dfn-control"} by a
    [`ServiceWorker`{#the-notrestoredreasons-interface:serviceworker}](https://w3c.github.io/ServiceWorker/#serviceworker){x-internal="serviceworker"}
    while the page was in [back/forward
    cache](browsing-the-web.html#note-bfcache).
    [\[SW\]](references.html#refsSW)

\"[`serviceworker-claimed`]{#sw-claimed .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-27}](dom.html#document)\'s
    [service worker
    client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client){#the-notrestoredreasons-interface:serviceworkercontainer-service-worker-client-2
    x-internal="serviceworkercontainer-service-worker-client"}\'s
    [active service
    worker](webappapis.html#concept-environment-active-service-worker){#the-notrestoredreasons-interface:concept-environment-active-service-worker}
    was claimed while the page was in [back/forward
    cache](browsing-the-web.html#note-bfcache).
    [\[SW\]](references.html#refsSW)

\"[`serviceworker-postmessage`]{#blocking-sw-message .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-28}](dom.html#document)\'s
    [service worker
    client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client){#the-notrestoredreasons-interface:serviceworkercontainer-service-worker-client-3
    x-internal="serviceworkercontainer-service-worker-client"}\'s
    [active service
    worker](webappapis.html#concept-environment-active-service-worker){#the-notrestoredreasons-interface:concept-environment-active-service-worker-2}
    received a message while the page was in [back/forward
    cache](browsing-the-web.html#note-bfcache).
    [\[SW\]](references.html#refsSW)

\"[`serviceworker-version-activated`]{#blocking-sw-activation .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-29}](dom.html#document)\'s
    [service worker
    client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client){#the-notrestoredreasons-interface:serviceworkercontainer-service-worker-client-4
    x-internal="serviceworkercontainer-service-worker-client"}\'s
    [active service
    worker](webappapis.html#concept-environment-active-service-worker){#the-notrestoredreasons-interface:concept-environment-active-service-worker-3}\'s
    version was activated while the page was in [back/forward
    cache](browsing-the-web.html#note-bfcache).
    [\[SW\]](references.html#refsSW)

\"[`serviceworker-unregistered`]{#blocking-sw-unregistered .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-30}](dom.html#document)\'s
    [service worker
    client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client){#the-notrestoredreasons-interface:serviceworkercontainer-service-worker-client-5
    x-internal="serviceworkercontainer-service-worker-client"}\'s
    [active service
    worker](webappapis.html#concept-environment-active-service-worker){#the-notrestoredreasons-interface:concept-environment-active-service-worker-4}\'s
    [service worker
    registration](https://w3c.github.io/service-workers/#dfn-service-worker-registration){#the-notrestoredreasons-interface:service-worker-registration
    x-internal="service-worker-registration"} was
    [unregistered](https://w3c.github.io/service-workers/#navigator-service-worker-unregister){#the-notrestoredreasons-interface:service-worker-unregister
    x-internal="service-worker-unregister"} while the page was in
    [back/forward cache](browsing-the-web.html#note-bfcache).
    [\[SW\]](references.html#refsSW)

\"[`sharedworker`]{#blocking-sharedworker .dfn export=""}\"
:   This
    [`Document`{#the-notrestoredreasons-interface:document-31}](dom.html#document)
    was in the [owner
    set](workers.html#concept-WorkerGlobalScope-owner-set){#the-notrestoredreasons-interface:concept-WorkerGlobalScope-owner-set}
    of a
    [`SharedWorkerGlobalScope`{#the-notrestoredreasons-interface:sharedworkerglobalscope}](workers.html#sharedworkerglobalscope).

\"[`smartcardconnection`]{#blocking-smartcard .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-32}](dom.html#document)
    had an active
    [`SmartCardConnection`{#the-notrestoredreasons-interface:smartcardconnection}](https://wicg.github.io/web-smart-card/#dom-smartcardconnection){x-internal="smartcardconnection"}
    while
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-14}.

\"[`speechrecognition`]{#blocking-speech-reco .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-33}](dom.html#document)
    had an active
    [`SpeechRecognition`{#the-notrestoredreasons-interface:speechrecognition}](https://wicg.github.io/speech-api/#speechrecognition){x-internal="speechrecognition"}
    while
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-15}.

\"[`storageaccess`]{#blocking-storage .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-34}](dom.html#document)
    requested storage access permission by using the Storage Access API.

\"[`unload-listener`]{#blocking-unload .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-35}](dom.html#document)
    registered an [event
    listener](https://dom.spec.whatwg.org/#concept-event-listener){#the-notrestoredreasons-interface:event-listener
    x-internal="event-listener"} for the
    [`unload`{#the-notrestoredreasons-interface:event-unload}](indices.html#event-unload)
    event.

\"[`video-capture`]{#blocking-video-capture .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-36}](dom.html#document)
    requested video capture permission by using Media Capture and
    Streams\'s
    [`getUserMedia()`{#the-notrestoredreasons-interface:getusermedia()-2}](https://w3c.github.io/mediacapture-main/getusermedia.html#dom-mediadevices-getusermedia){x-internal="getusermedia()"}
    with video. [\[MEDIASTREAM\]](references.html#refsMEDIASTREAM)

\"[`webhid`]{#blocking-webhid .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-37}](dom.html#document)
    called the WebHID API\'s
    [`requestDevice()`{#the-notrestoredreasons-interface:request-device}](https://wicg.github.io/webhid/#requestdevice-method){x-internal="request-device"}
    method.

\"[`webshare`]{#blocking-webshare .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-38}](dom.html#document)
    used the Web Share API\'s
    [`navigator.share()`{#the-notrestoredreasons-interface:dom-navigator-share}](https://w3c.github.io/web-share/#share-method){x-internal="dom-navigator-share"}
    method.

\"[`webtransport`]{#blocking-webtransport .dfn export=""}\"
:   While
    [unloading](document-lifecycle.html#unload-a-document){#the-notrestoredreasons-interface:unload-a-document-16},
    an open
    [`WebTransport`{#the-notrestoredreasons-interface:webtransport}](https://w3c.github.io/webtransport/#webtransport){x-internal="webtransport"}
    connection was shut down, so the page was not in a state that could
    be stored in the [back/forward
    cache](browsing-the-web.html#note-bfcache).
    [\[WEBTRANSPORT\]](references.html#refsWEBTRANSPORT)

\"[`webxrdevice`]{#blocking-webxr .dfn export=""}\"
:   The
    [`Document`{#the-notrestoredreasons-interface:document-39}](dom.html#document)
    created a
    [`XRSystem`{#the-notrestoredreasons-interface:xrsystem}](https://immersive-web.github.io/webxr/#xrsystem){x-internal="xrsystem"}.

------------------------------------------------------------------------

A
[`NotRestoredReasons`{#the-notrestoredreasons-interface:notrestoredreasons-4}](#notrestoredreasons)
object has a [backing
struct]{#concept-not-restored-reasons-backing-struct .dfn}, a [not
restored
reasons](#nrr-struct){#the-notrestoredreasons-interface:nrr-struct} or
null, initially null.

A
[`NotRestoredReasons`{#the-notrestoredreasons-interface:notrestoredreasons-5}](#notrestoredreasons)
object has a [reasons array]{#concept-not-restored-reasons-reasons
.dfn}, a
`FrozenArray<`[`NotRestoredReasonDetails`{#the-notrestoredreasons-interface:notrestoredreasondetails-5}](#notrestoredreasondetails)`>`
or null, initially null.

A
[`NotRestoredReasons`{#the-notrestoredreasons-interface:notrestoredreasons-6}](#notrestoredreasons)
object has a [children array]{#concept-not-restored-reasons-children
.dfn}, a
`FrozenArray<`[`NotRestoredReasons`](#notrestoredreasons){#the-notrestoredreasons-interface:notrestoredreasons-7}`>`
or null, initially null.

The [`src`]{#dom-not-restored-reasons-src .dfn
dfn-for="NotRestoredReasons" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this-2
x-internal="this"}\'s [backing
struct](#concept-not-restored-reasons-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reasons-backing-struct}\'s
[src](#nrr-src){#the-notrestoredreasons-interface:nrr-src}.

The [`id`]{#dom-not-restored-reasons-id .dfn
dfn-for="NotRestoredReasons" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this-3
x-internal="this"}\'s [backing
struct](#concept-not-restored-reasons-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reasons-backing-struct-2}\'s
[id](#nrr-id){#the-notrestoredreasons-interface:nrr-id}.

The [`name`]{#dom-not-restored-reasons-name .dfn
dfn-for="NotRestoredReasons" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this-4
x-internal="this"}\'s [backing
struct](#concept-not-restored-reasons-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reasons-backing-struct-3}\'s
[name](#nrr-name){#the-notrestoredreasons-interface:nrr-name}.

The [`url`]{#dom-not-restored-reasons-url .dfn
dfn-for="NotRestoredReasons" dfn-type="attribute"} getter steps are:

1.  If
    [this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this-5
    x-internal="this"}\'s [backing
    struct](#concept-not-restored-reasons-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reasons-backing-struct-4}\'s
    [URL](#nrr-url){#the-notrestoredreasons-interface:nrr-url} is null,
    then return null.

2.  Return
    [this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this-6
    x-internal="this"}\'s [backing
    struct](#concept-not-restored-reasons-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reasons-backing-struct-5}\'s
    [URL](#nrr-url){#the-notrestoredreasons-interface:nrr-url-2},
    [serialized](https://url.spec.whatwg.org/#concept-url-serializer){#the-notrestoredreasons-interface:concept-url-serializer
    x-internal="concept-url-serializer"}.

The [`reasons`]{#dom-not-restored-reasons-reasons .dfn
dfn-for="NotRestoredReasons" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this-7
x-internal="this"}\'s [reasons
array](#concept-not-restored-reasons-reasons){#the-notrestoredreasons-interface:concept-not-restored-reasons-reasons}.

The [`children`]{#dom-not-restored-reasons-children .dfn
dfn-for="NotRestoredReasons" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#the-notrestoredreasons-interface:this-8
x-internal="this"}\'s [children
array](#concept-not-restored-reasons-children){#the-notrestoredreasons-interface:concept-not-restored-reasons-children}.

To [create a `NotRestoredReasons`
object]{#create-a-notrestoredreasons-object .dfn export=""} given a [not
restored
reasons](#nrr-struct){#the-notrestoredreasons-interface:nrr-struct-2}
`backingStruct`{.variable} and a
[realm](https://tc39.es/ecma262/#sec-code-realms){#the-notrestoredreasons-interface:realm-2
x-internal="realm"} `realm`{.variable}:

1.  Let `notRestoredReasons`{.variable} be a new
    [`NotRestoredReasons`{#the-notrestoredreasons-interface:notrestoredreasons-8}](#notrestoredreasons)
    object created in `realm`{.variable}.

2.  Set `notRestoredReasons`{.variable}\'s [backing
    struct](#concept-not-restored-reasons-backing-struct){#the-notrestoredreasons-interface:concept-not-restored-reasons-backing-struct-6}
    to `backingStruct`{.variable}.

3.  If `backingStruct`{.variable}\'s
    [reasons](#nrr-reasons){#the-notrestoredreasons-interface:nrr-reasons}
    is null, set `notRestoredReasons`{.variable}\'s [reasons
    array](#concept-not-restored-reasons-reasons){#the-notrestoredreasons-interface:concept-not-restored-reasons-reasons-2}
    to null.

4.  Otherwise:

    1.  Let `reasonsArray`{.variable} be an empty
        [list](https://infra.spec.whatwg.org/#list){#the-notrestoredreasons-interface:list
        x-internal="list"}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#the-notrestoredreasons-interface:list-iterate
        x-internal="list-iterate"} `reason`{.variable} of
        `backingStruct`{.variable}\'s
        [reasons](#nrr-reasons){#the-notrestoredreasons-interface:nrr-reasons-2}:

        1.  [Create a `NotRestoredReasonDetails`
            object](#create-a-notrestoredreasondetails-object){#the-notrestoredreasons-interface:create-a-notrestoredreasondetails-object}
            given `reason`{.variable} and `realm`{.variable}, and
            [append](https://infra.spec.whatwg.org/#list-append){#the-notrestoredreasons-interface:list-append
            x-internal="list-append"} it to `reasonsArray`{.variable}.

    3.  Set `notRestoredReasons`{.variable}\'s [reasons
        array](#concept-not-restored-reasons-reasons){#the-notrestoredreasons-interface:concept-not-restored-reasons-reasons-3}
        to the result of [creating a frozen
        array](https://webidl.spec.whatwg.org/#dfn-create-frozen-array){#the-notrestoredreasons-interface:creating-a-frozen-array
        x-internal="creating-a-frozen-array"} given
        `reasonsArray`{.variable}.

5.  If `backingStruct`{.variable}\'s
    [children](#nrr-children){#the-notrestoredreasons-interface:nrr-children}
    is null, set `notRestoredReasons`{.variable}\'s [children
    array](#concept-not-restored-reasons-children){#the-notrestoredreasons-interface:concept-not-restored-reasons-children-2}
    to null.

6.  Otherwise:

    1.  Let `childrenArray`{.variable} be an empty
        [list](https://infra.spec.whatwg.org/#list){#the-notrestoredreasons-interface:list-2
        x-internal="list"}.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#the-notrestoredreasons-interface:list-iterate-2
        x-internal="list-iterate"} `child`{.variable} of
        `backingStruct`{.variable}\'s
        [children](#nrr-children){#the-notrestoredreasons-interface:nrr-children-2}:

        1.  [Create a `NotRestoredReasons`
            object](#create-a-notrestoredreasons-object){#the-notrestoredreasons-interface:create-a-notrestoredreasons-object}
            given `child`{.variable} and `realm`{.variable} and
            [append](https://infra.spec.whatwg.org/#list-append){#the-notrestoredreasons-interface:list-append-2
            x-internal="list-append"} it to `childrenArray`{.variable}.

    3.  Set `notRestoredReasons`{.variable}\'s [children
        array](#concept-not-restored-reasons-children){#the-notrestoredreasons-interface:concept-not-restored-reasons-children-3}
        to the result of [creating a frozen
        array](https://webidl.spec.whatwg.org/#dfn-create-frozen-array){#the-notrestoredreasons-interface:creating-a-frozen-array-2
        x-internal="creating-a-frozen-array"} given
        `childrenArray`{.variable}.

7.  Return `notRestoredReasons`{.variable}.

A [not restored reasons]{#nrr-struct .dfn export=""} is a
[struct](https://infra.spec.whatwg.org/#struct){#the-notrestoredreasons-interface:struct-2
x-internal="struct"} with the following
[items](https://infra.spec.whatwg.org/#struct-item){#the-notrestoredreasons-interface:struct-item-2
x-internal="struct-item"}:

- [src]{#nrr-src .dfn}, a [scalar value
  string](https://infra.spec.whatwg.org/#scalar-value-string){#the-notrestoredreasons-interface:scalar-value-string
  x-internal="scalar-value-string"} or null, initially null.

- [id]{#nrr-id .dfn}, a string or null, initially null.

- [name]{#nrr-name .dfn}, a string or null, initially null.

- [url]{#nrr-url .dfn}, a
  [URL](https://url.spec.whatwg.org/#concept-url){#the-notrestoredreasons-interface:url
  x-internal="url"} or null, initially null.

- [reasons]{#nrr-reasons .dfn}, null or a
  [list](https://infra.spec.whatwg.org/#list){#the-notrestoredreasons-interface:list-3
  x-internal="list"} of [not restored reason
  details](#nrr-details-struct){#the-notrestoredreasons-interface:nrr-details-struct-3},
  initially null.

- [children]{#nrr-children .dfn}, null or a
  [list](https://infra.spec.whatwg.org/#list){#the-notrestoredreasons-interface:list-4
  x-internal="list"} of [not restored
  reasons](#nrr-struct){#the-notrestoredreasons-interface:nrr-struct-3},
  initially null.

A [`Document`\'s not restored reasons]{#concept-document-nrr .dfn
dfn-for="Document" lt="not restored reasons" export=""} is its [node
navigable](document-sequences.html#node-navigable){#the-notrestoredreasons-interface:node-navigable-4}\'s
[active session history
entry](document-sequences.html#nav-active-history-entry){#the-notrestoredreasons-interface:nav-active-history-entry}\'s
[document
state](browsing-the-web.html#she-document-state){#the-notrestoredreasons-interface:she-document-state}\'s
[not restored
reasons](browsing-the-web.html#document-state-not-restored-reasons){#the-notrestoredreasons-interface:document-state-not-restored-reasons},
if
[`Document`{#the-notrestoredreasons-interface:document-40}](dom.html#document)\'s
[node
navigable](document-sequences.html#node-navigable){#the-notrestoredreasons-interface:node-navigable-5}
is a [top-level
traversable](document-sequences.html#top-level-traversable){#the-notrestoredreasons-interface:top-level-traversable};
otherwise null.

[← 7 Loading web pages](browsers.html) --- [Table of Contents](./) ---
[7.3 Infrastructure for sequences of documents
→](document-sequences.html)
